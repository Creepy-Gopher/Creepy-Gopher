package scraper

import (
	"creepy/internal/property"
	"creepy/internal/utils/converter"
	"fmt"
	"github.com/google/uuid"
	"github.com/playwright-community/playwright-go"
	"regexp"
	"strconv"
	"strings"
)

type LinkType string

const (
	LinkType_Divar_Buy  LinkType = "Divar_Buy"
	LinkType_Divar_Rent LinkType = "Divar_Rent"
)

func ScrapePropertyData(url string, linkType LinkType) (*property.Property, error) {
	parts := strings.Split(string(linkType), "-")
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
		Latitude:     0,
		Longitude:    0,
		Source:       "",
		URL:          "",
		Images:       nil,
	}
	if propertyType != "" {
		resultProperty.Type = propertyType
	}

	for field, selector := range selectors {
		element, err := page.QuerySelector(selector)
		content, _ := element.InnerText()
		content = converter.ReplacePersianDigits(content)
		if err == nil {
			switch field {
			case "Title":
				resultProperty.Title = content
			case "Description":
				resultProperty.Description = content
			case "BuyPrice":
				num, err := strconv.ParseUint(regexp.MustCompile(`[٬ تومان]`).ReplaceAllString(content, ""), 10, 64)
				if err != nil {
					resultProperty.BuyPrice = num
				}
			case "Area":
				num, err := strconv.ParseUint(content, 10, 64)
				if err != nil {
					resultProperty.Area = num
				}
			case "Rooms":
				num, err := strconv.ParseUint(content, 10, 0)
				if err != nil {
					resultProperty.Rooms = uint(num)
				}
			case "City":
				//TODO: Fix &zwnj
				resultProperty.City = content
			case "BuildYear":
				num, err := strconv.ParseUint(content, 10, 64)
				if err != nil {
					resultProperty.BuyPrice = num
				}
			case "Floor":
				parts := strings.Split(string(content), " از ")
				num, err := strconv.ParseUint(parts[0], 10, 0)
				if err != nil {
					resultProperty.Rooms = uint(num)
				}
			case "HasElevator":
				resultProperty.HasElevator = !strings.Contains(content, "ندارد")
			case "HasStorage":
				resultProperty.HasStorage = !strings.Contains(content, "ندارد")
			case "Images":
				src, _ := element.GetAttribute("src")
				resultProperty.Images = []string{src}
			}
		}

	}
	fmt.Println(resultProperty)
	return &resultProperty, nil
}
