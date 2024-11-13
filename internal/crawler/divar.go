package crawler

import (
	"creepy/internal/parser"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/tebeka/selenium"
)

// Constants for Selenium paths and settings
const (
	//seleniumPath    = "/home/selenium-server-4.26.0.jar"
	geckoDriverPath = "./pkg/utils/firefoxdriver/geckodriver.exe"
	port            = 4444
)

// CrawlDivar crawls a specified category type with a maximum link count and returns a list of links.
func crawlDivar(linkType parser.LinkType, maxLinks int) []string {
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

	// Loop through categories and start a goroutine for each
	for category, label := range categories {
		wg.Add(1)
		go func(category, label string) {
			defer wg.Done()
			links, _ := crawlCategory(category, label, maxLinks/len(categories))
			results <- links
		}(category, label)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(results)

	// Collect results from channel into a single list
	for result := range results {
		crawledLinks = append(crawledLinks, result...)
	}

	return crawledLinks
}
func customURLEncode(input string) string {
	// ابتدا URL رو URL-encode می‌کنیم
	encoded := url.QueryEscape(input)
	// بعد از آن / و : رو به حالت اصلی خودشون برمی‌گردونیم
	encoded = strings.ReplaceAll(encoded, "%2F", "/")
	encoded = strings.ReplaceAll(encoded, "%3A", ":")
	return encoded
}

// crawlCategory performs web crawling for a specific category and returns a list of links crawled
func crawlCategory(category, label string, maxLinks int) ([]string, error) {
	// راه‌اندازی مستقیم GeckoDriver بدون نیاز به Selenium Server
	service, err := selenium.NewGeckoDriverService(geckoDriverPath, port)
	if err != nil {
		return nil, fmt.Errorf("Error starting the GeckoDriver service: %v", err)
	}
	defer service.Stop()

	// ایجاد یک جلسه WebDriver برای فایرفاکس
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		return nil, fmt.Errorf("Error creating new remote session: %v", err)
	}
	defer wd.Quit()

	// Define URL based on the category and navigate to it
	url := fmt.Sprintf("https://divar.ir/s/iran/%s", category)
	if err := wd.Get(url); err != nil {
		return nil, fmt.Errorf("Error navigating to page: %v", err)
	}

	//time.Sleep(2 * time.Second) // Wait for the page to load

	// Variables for tracking visited links and session state
	visitedLinks := make(map[string]bool)
	links := []string{}
	count := 0
	ticker := time.NewTicker(500 * time.Millisecond)   // Ticker to scroll the page periodically
	noNewLinksTimer := time.NewTicker(5 * time.Second) // Timer for detecting if new links are found
	defer ticker.Stop()
	defer noNewLinksTimer.Stop()

	lastLinkCount := 0          // Track last link count to detect no new links
	refreshCount := 0           // Track page refresh count
	moveToNextCategory := false // Flag to control when to move to the next category

	for count < maxLinks && !moveToNextCategory {
		select {
		case <-ticker.C:
			// Scroll the page down to load more content
			_, err := wd.ExecuteScript("window.scrollBy(0, window.innerHeight / 2);", nil)
			if err != nil {
				return links, fmt.Errorf("Error during scrolling: %v", err)
			}
			time.Sleep(200 * time.Millisecond)

			// Check if "Show More" button is available and click it
			moreButton, err := wd.FindElement(selenium.ByXPATH, "//*[contains(text(), 'آگهی‌های بیشتر')]")
			if err == nil {
				_, err = wd.ExecuteScript("arguments[0].click();", []interface{}{moreButton})
				if err != nil {
					return links, fmt.Errorf("Error clicking 'Show More' button: %v", err)
				}
				time.Sleep(500 * time.Millisecond) // Wait for new content to load
			}

			// Find ad boxes on the page
			adBoxes, err := wd.FindElements(selenium.ByCSSSelector, ".kt-post-card")
			if err != nil {
				continue // Skip if there was an error finding ad boxes
			}

			// Loop through each ad box to extract link
			for _, adBox := range adBoxes {
				var encodedLink string
				// Try up to 3 times to retrieve the ad link
				for retry := 0; retry < 3; retry++ {
					linkElement, err := adBox.FindElement(selenium.ByCSSSelector, "a[href^='/v/']")
					if err != nil {
						continue // Retry if link is not found
					}
					href, err := linkElement.GetAttribute("href")
					if err != nil {
						time.Sleep(500 * time.Millisecond)
						continue
					}
					encodedLink = customURLEncode("https://divar.ir" + href)

					break
				}

				// Skip if the link is empty, already visited, or max links are reached
				if encodedLink == "" || visitedLinks[encodedLink] || count >= maxLinks {
					continue
				}

				// Add link to list and mark as visited
				visitedLinks[encodedLink] = true
				links = append(links, encodedLink)
				count++
			}

		case <-noNewLinksTimer.C:
			// Check if new links are found; if not, refresh the page or move to next category
			if count == lastLinkCount {
				if refreshCount >= 1 { // If already refreshed once, move to next category
					moveToNextCategory = true
					break
				} else {
					wd.Refresh() // Refresh the page if no new links were found
					time.Sleep(2 * time.Second)
					refreshCount++
				}
			} else {
				lastLinkCount = count // Update last link count if new links are found
				refreshCount = 0      // Reset refresh counter
			}
		}
	}

	return links, nil // Return list of links crawled if successful
}
