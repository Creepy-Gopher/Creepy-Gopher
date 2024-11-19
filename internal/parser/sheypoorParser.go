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

func parseSheypoorBuy(url string, linkType LinkType, page playwright.Page) (*models.Property, error) {
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
				num, err := strconv.ParseUint(regexp.MustCompile(`[,]`).ReplaceAllString(content, ""), 10, 64)
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
				resultProperty.City = strings.Split(content, "، ")[0]
				resultProperty.District = strings.Split(content, "، ")[2]
			case "BuildYear":
				num, err := strconv.ParseUint(strings.Split(content, " ")[0], 10, 64)
				if err == nil {
					resultProperty.BuildYear = uint(num)
				}
			case "Floor":
				continue
			case "HasElevator":
				resultProperty.HasElevator = strings.Contains(content, "دارد")
			case "HasStorage":
				resultProperty.HasStorage = strings.Contains(content, "دارد")
			case "HasParking":
				resultProperty.HasParking = strings.Contains(content, "دارد")
			case "Images":
				src, _ := element.GetAttribute("src")
				resultProperty.Image = src
			}
		}
	}
	printStructFields(resultProperty)
	return &resultProperty, nil
}
