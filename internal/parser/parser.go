package parser

import (
	"creepy/internal/models"
	"creepy/pkg/utils"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type LinkType string
type parserFunction func(url string, linkType LinkType, page playwright.Page) (*models.Property, error)

const (
	Linktype_DivarBuy  LinkType = "Divar_Buy"
	Linktype_DivarRent LinkType = "Divar_Rent"
)

var parserFunctions = map[string]map[string]parserFunction{
	"Divar": {
		"Buy":  parseDivarBuy,
		"Rent": parseDivarRent,
	},
	"Sheypoor": {
		"Buy":  nil,
		"Rent": nil,
	},
}

func ParseProperties(urls []string, linkType LinkType) ([]*models.Property, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, fmt.Errorf("خطا در اجرای Playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		return nil, fmt.Errorf("خطا در اجرای مرورگر: %v", err)
	}
	defer browser.Close()

	var wg sync.WaitGroup
	results := make([]*models.Property, len(urls))
	errors := make(chan error, len(urls))

	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			page, err := browser.NewPage()
			if err != nil {
				errors <- fmt.Errorf("خطا در ایجاد صفحه: %v", err)
				return
			}
			defer page.Close()

			parts := strings.Split(string(linkType), "_")
			platform := parts[0]
			propertyType := parts[1]
			property, err := parserFunctions[platform][propertyType](url, linkType, page)
			if err != nil {
				errors <- err
				return
			}
			results[i] = property
		}(i, url)
	}

	wg.Wait()
	close(errors)

	for err := range errors {
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}

func parseDivarBuy(url string, linkType LinkType, page playwright.Page) (*models.Property, error) {
	parts := strings.Split(string(linkType), "_")
	platform := parts[0]
	propertyType := parts[1]
	selectors := Selectors[platform][propertyType]

	if _, err := page.Goto(url); err != nil {
		return nil, fmt.Errorf("error in openning page: %v", err)
	}

	resultProperty := models.Property{}

	resultProperty.Type = propertyType

	resultProperty.Source = platform
	resultProperty.URL = url

	for field, selector := range selectors {
		element, _ := page.QuerySelector(selector)
		if element == nil {
			continue
		}

		content, err := element.InnerHTML()
		if err != nil {
			fmt.Printf("Error getting inner HTML for field %s: %v\n", field, err)
			continue
		}

		content = utils.ReplacePersianDigits(content)
		if err == nil {
			switch field {
			case "Title":
				resultProperty.Title = content
			case "Description":
				resultProperty.Description = content
			case "BuyPrice":
				num, err := strconv.ParseUint(regexp.MustCompile(`[٬ تومان]`).ReplaceAllString(content, ""), 10, 64)
				if err == nil {
					resultProperty.BuyPrice = num
				}
			case "Area":
				num, err := strconv.ParseUint(content, 10, 64)
				if err == nil {
					resultProperty.Area = num
				}
			case "Rooms":
				num, err := strconv.ParseUint(content, 10, 0)
				if err == nil {
					resultProperty.Rooms = uint(num)
				}
			case "City":
				resultProperty.City = strings.Split(strings.Split(content, "،")[0], " در ")[1]
				resultProperty.District = strings.Split(content, "، ")[1]
			case "BuildYear":
				num, err := strconv.ParseUint(content, 10, 64)
				if err == nil {
					resultProperty.BuildYear = uint(num)
				}
			case "Floor":
				floor, err := handlerFloor(content)
				if err != nil {
					continue
				}
				resultProperty.Floor = floor
			case "HasElevator":
				resultProperty.HasElevator = !strings.Contains(content, "ندارد")
			case "HasStorage":
				resultProperty.HasStorage = !strings.Contains(content, "ندارد")
			case "HasParking":
				resultProperty.HasParking = !strings.Contains(content, "ندارد")
			case "Images":
				src, _ := element.GetAttribute("src")
				resultProperty.Image = src
			}
		}
	}
	//printStructFields(resultProperty)
	return &resultProperty, nil
}

func parseDivarRent(url string, linkType LinkType, page playwright.Page) (*models.Property, error) {
	parts := strings.Split(string(linkType), "_")
	platform := parts[0]
	propertyType := parts[1]
	selectors := Selectors[platform][propertyType]

	if _, err := page.Goto(url); err != nil {
		return nil, fmt.Errorf("error in openning page: %v", err)
	}

	resultProperty := models.Property{}

	resultProperty.Type = propertyType
	resultProperty.Source = platform
	resultProperty.URL = url

	isConvertable := "RentTypeNonConvertable"
	for field, selector := range selectors {
		element, _ := page.QuerySelector(selector)
		if element == nil {
			continue
		}

		content, err := element.InnerHTML()
		if err != nil {
			fmt.Printf("Error getting inner HTML for field %s: %v\n", field, err)
			continue
		}

		content = utils.ReplacePersianDigits(content)

		switch field {
		case "Title":
			resultProperty.Title = content
		case "Description":
			resultProperty.Description = content
		case "Area":
			num, err := strconv.ParseUint(content, 10, 64)
			if err == nil {
				resultProperty.Area = num
			}
		case "Rooms":
			num, err := strconv.ParseUint(content, 10, 0)
			if err == nil {
				resultProperty.Rooms = uint(num)
			}
		case "City":
			resultProperty.City = strings.Split(strings.Split(content, "،")[0], " در ")[1]
			resultProperty.District = strings.Split(content, "، ")[1]
		case "BuildYear":
			num, err := strconv.ParseUint(content, 10, 64)
			if err == nil {
				resultProperty.BuildYear = uint(num)
			}
		case "Images":
			src, _ := element.GetAttribute("src")
			resultProperty.Image = src
		case "Convertable":
			if content == "ودیعه و اجارهٔ این ملک قابل تبدیل است." {
				isConvertable = "RentTypeConvertable"
			}
		}
	}
	// New Selectors
	selectors = Selectors[platform][isConvertable]

	if isConvertable == "RentTypeConvertable" {

		for field, selector := range selectors {
			element, _ := page.QuerySelector(selector)
			if element == nil {
				continue
			}
			content, err := element.InnerHTML()
			if err != nil {
				fmt.Printf("Error getting inner HTML for field %s: %v\n", field, err)
				continue
			}
			content = utils.ReplacePersianDigits(content)

			switch field {
			case "HasElevator":
				resultProperty.HasElevator = !strings.Contains(content, "ندارد")
			case "HasStorage":
				resultProperty.HasStorage = !strings.Contains(content, "ندارد")
			case "HasParking":
				resultProperty.HasParking = !strings.Contains(content, "ندارد")
			case "Floor":
				floor, err := handlerFloor(content)
				if err != nil {
					continue
				}
				resultProperty.Floor = floor

			case "RahnPriceMax":
				price, err := handlerConvertablePrice(content)
				if err != nil {
					continue
				}
				resultProperty.RahnPriceMax = price
			case "RahnPriceMin":
				price, err := handlerConvertablePrice(content)
				if err != nil {
					continue
				}
				resultProperty.RahnPriceMin = price
			case "RentPriceMin":
				price, err := handlerConvertablePrice(content)
				if err != nil {
					continue
				}
				resultProperty.RentPriceMin = price
			case "RentPriceMin2":
				price, err := handlerConvertablePrice(content)
				if err != nil {
					continue
				}
				resultProperty.RentPriceMin = price
			case "RentPriceMax":
				price, err := handlerConvertablePrice(content)
				if err != nil {
					continue
				}
				resultProperty.RentPriceMax = uint64(price)
			}
		}
	}
	printStructFields(resultProperty)
	return &resultProperty, nil
}
