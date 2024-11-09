package scraper

import (
	"creepy/internal/property"
	"creepy/internal/utils/converter"
	"fmt"
	"github.com/google/uuid"
	"github.com/playwright-community/playwright-go"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type LinkType string

const (
	LinkType_Divar_Buy  LinkType = "Divar_Buy"
	LinkType_Divar_Rent LinkType = "Divar_Rent"
)

func PrintStructFields(s interface{}) {
	val := reflect.ValueOf(s)
	typ := val.Type()

	if typ.Kind() != reflect.Struct {
		fmt.Println("Provided value is not a struct")
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		fmt.Printf("%s: %v\n", field.Name, value)
	}
}

func ScrapePropertyData(url string, linkType LinkType) (*property.Property, error) {
	parts := strings.Split(string(linkType), "_")
	platform := parts[0]
	propertyType := parts[1]
	selectors := Selectors[platform][propertyType]

	pw, err := playwright.Run()
	defer pw.Stop()
	if err != nil {
		return nil, fmt.Errorf("خطا در اجرای Playwright: %v", err)
	}

	browser, err := pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	defer browser.Close()
	if err != nil {
		return nil, fmt.Errorf("خطا در اجرای مرورگر: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		return nil, fmt.Errorf("خطا در ایجاد صفحه: %v", err)
	}
	if _, err := page.Goto(url); err != nil {
		return nil, fmt.Errorf("خطا در باز کردن صفحه: %v", err)
	}

	resultProperty := property.Property{
		ID:           uuid.UUID{},
		Title:        "",
		Description:  "",
		BuyPrice:     0,
		RentPrice:    0,
		RentPriceMin: 0,
		RentPriceMax: 0,
		RahnPriceMin: 0,
		RahnPriceMax: 0,
		Area:         0,
		Rooms:        0,
		DealingType:  "",
		Type:         "",
		City:         "",
		District:     "",
		Address:      "",
		BuildYear:    0,
		Floor:        0,
		HasElevator:  false,
		HasStorage:   false,
		HasParking:   false,
		Latitude:     0,
		Longitude:    0,
		Source:       "",
		URL:          "",
		Images:       nil,
	}

	if propertyType != "" {
		resultProperty.Type = propertyType
	} else {
		//TODO: Sheypoor
	}

	resultProperty.Source = platform
	resultProperty.URL = url

	for field, selector := range selectors {
		element, _ := page.QuerySelector(selector)
		content, err := element.InnerHTML()
		content = converter.ReplacePersianDigits(content)
		//fmt.Println(field)
		//fmt.Println(content)
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
				parts := strings.Split(string(content), " از ")
				num, err := strconv.ParseUint(parts[0], 10, 0)
				if err == nil {
					resultProperty.Floor = uint(num)
				}
			case "HasElevator":
				resultProperty.HasElevator = !strings.Contains(content, "ندارد")
			case "HasStorage":
				resultProperty.HasStorage = !strings.Contains(content, "ندارد")
			case "HasParking":
				resultProperty.HasParking = !strings.Contains(content, "ندارد")
			case "Images":
				src, _ := element.GetAttribute("src")
				resultProperty.Images = []string{src}
			}
		}

	}
	PrintStructFields(resultProperty)
	return &resultProperty, nil
}
