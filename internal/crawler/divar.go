package crawler

import (
	"creepy/internal/parser"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/tebeka/selenium"
)

// Constants for Selenium paths and settings
const (
	seleniumPath    = "/home/selenium-server-4.26.0.jar"
	geckoDriverPath = "/home/geckodriver"
	port            = 4444
	maxRetries      = 10 // حداکثر تعداد تلاش‌ها قبل از رفرش صفحه
)

// CrawlDivar crawls a specified category type with a maximum link count and returns a list of links.
func CrawlDivar(linkType parser.LinkType, maxLinks int) []string {
	var categories map[string]string
	parts := strings.Split(string(linkType), "_")
	propertyType := parts[1]

	// Set categories based on property type
	switch propertyType {
	case "Buy":
		categories = map[string]string{
			"buy-apartment": "خرید",
			"buy-villa":     "خرید",
		}
	case "Rent":
		categories = map[string]string{
			"rent-apartment": "رهن/اجاره",
			"rent-villa":     "رهن/اجاره",
		}
	}

	var wg sync.WaitGroup
	results := make(chan []string, len(categories))
	var crawledLinks []string

	for category, label := range categories {
		wg.Add(1)
		go func(category, label string) {
			defer wg.Done()
			links, _ := crawlCategory(category, label, maxLinks/len(categories))
			results <- links
		}(category, label)
	}

	wg.Wait()
	close(results)

	for result := range results {
		crawledLinks = append(crawledLinks, result...)
	}

	return crawledLinks
}

func customURLEncode(input string) string {
	encoded := url.QueryEscape(input)
	encoded = strings.ReplaceAll(encoded, "%2F", "/")
	encoded = strings.ReplaceAll(encoded, "%3A", ":")
	return encoded
}

func waitForPageLoad(wd selenium.WebDriver) error {
	retries := 0
	for {
		state, err := wd.ExecuteScript("return document.readyState;", nil)
		if err != nil {
			return fmt.Errorf("error checking page load state: %v", err)
		}

		if state == "complete" {
			return nil
		}

		retries++
		if retries >= maxRetries {
			// اگر بعد از 10 تلاش صفحه کامل لود نشد، رفرش می‌کنیم
			fmt.Println("Page did not load completely after multiple attempts. Refreshing...")
			if err := wd.Refresh(); err != nil {
				return fmt.Errorf("error refreshing the page: %v", err)
			}
			retries = 0 // شمارشگر را ریست می‌کنیم
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func crawlCategory(category, label string, maxLinks int) ([]string, error) {
	opts := []selenium.ServiceOption{
		selenium.GeckoDriver(geckoDriverPath),
	}
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		return nil, fmt.Errorf("Error starting the Selenium service: %v", err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{
		"browserName": "firefox",
		"moz:firefoxOptions": map[string]interface{}{
			"args": []string{"-headless"},
		},
	}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		return nil, fmt.Errorf("Error creating new remote session: %v", err)
	}
	defer wd.Quit()

	url := fmt.Sprintf("https://divar.ir/s/iran/%s", category)
	if err := wd.Get(url); err != nil {
		return nil, fmt.Errorf("Error navigating to page: %v", err)
	}

	if err := waitForPageLoad(wd); err != nil {
		return nil, err
	}

	visitedLinks := make(map[string]bool)
	links := []string{}
	count := 0

	for count < maxLinks {
		_, err := wd.ExecuteScript("window.scrollBy(0, window.innerHeight);", nil)
		if err != nil {
			fmt.Println("Scrolling failed, refreshing page...")
			if err := wd.Refresh(); err != nil {
				return links, fmt.Errorf("Error refreshing the page: %v", err)
			}
			if err := waitForPageLoad(wd); err != nil {
				return links, err
			}
			continue
		}

		time.Sleep(2 * time.Second)

		moreButton, err := wd.FindElement(selenium.ByXPATH, "//*[contains(text(), 'آگهی‌های بیشتر')]")
		if err == nil {
			fmt.Println("Found 'Show More' button, clicking to load more ads...")
			_, err = wd.ExecuteScript("arguments[0].click();", []interface{}{moreButton})
			if err != nil {
				return links, fmt.Errorf("Error clicking 'Show More' button: %v", err)
			}
			time.Sleep(1 * time.Second)
			if err := waitForPageLoad(wd); err != nil {
				return links, err
			}
		}

		pageSource, err := wd.PageSource()
		if err != nil {
			continue
		}

		re := regexp.MustCompile(`href="(/v/[^"]+)"`)
		matches := re.FindAllStringSubmatch(pageSource, -1)

		for _, match := range matches {
			fullLink := customURLEncode("https://divar.ir" + match[1])

			if fullLink == "" || visitedLinks[fullLink] {
				continue
			}

			visitedLinks[fullLink] = true
			links = append(links, fullLink)
			count++
			//fmt.Printf("New link found: %s\n", fullLink)
			if count >= maxLinks {
				break
			}
		}

		//fmt.Printf("Number of links collected so far: %d\n", len(links))
	}

	return links, nil
}
