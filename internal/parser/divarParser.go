package parser

import (
	"creepy/internal/models"
	"creepy/pkg/utils"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"regexp"
	"strconv"
	"strings"
)

func parseDivarBuy(url string, linkType LinkType, page playwright.Page) (*models.Property, error) {
	parts := strings.Split(string(linkType), "_")
	platform := parts[0]
	propertyType := parts[1]
	selectors := Selectors[platform][propertyType]

	if _, err := page.Goto(url); err != nil {
		return nil, fmt.Errorf("error in openning page: %v", err)
	}

	resultProperty := models.Property{}

	resultProperty.DealingType = propertyType

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
				num, err := handlerArea(content)
				if err != nil {
					continue
				}
				resultProperty.Area = num
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
	printStructFields(resultProperty)
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

	resultProperty.DealingType = propertyType
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
			num, err := handlerArea(content)
			if err != nil {
				continue
			}
			resultProperty.Area = num
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

		case "RahnPriceNonConvertable":
			price, err := hanldeNonConvertablePrices(content)
			if err != nil {
				continue
			}
			resultProperty.RahnPriceMin = price
			resultProperty.RahnPriceMax = price

		case "RentPriceNonConvertable":
			price, err := hanldeNonConvertablePrices(content)
			if err != nil {
				continue
			}
			resultProperty.RentPriceMin = price
			resultProperty.RentPriceMax = price
			if price == 0 {
				resultProperty.DealingType = "Rahn"
			}

		}
	}
	printStructFields(resultProperty)
	return &resultProperty, nil
}
