package parser

import (
	"creepy/internal/models"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"strings"
	"sync"
)

type LinkType string
type parserFunction func(url string, linkType LinkType, page playwright.Page) (*models.Property, error)

const (
	Linktype_DivarBuy    LinkType = "Divar_Buy"
	Linktype_DivarRent   LinkType = "Divar_Rent"
	Linktype_SheypoorBuy LinkType = "Sheypoor_Buy"
)

var parserFunctions = map[string]map[string]parserFunction{
	"Divar": {
		"Buy":  parseDivarBuy,
		"Rent": parseDivarRent,
	},
	"Sheypoor": {
		"Buy":  parseSheypoorBuy,
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
			f := parserFunctions[platform][propertyType]
			property, err := f(url, linkType, page)
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
