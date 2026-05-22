package renderer

import (
	"os"
	"path/filepath"

	"github.com/playwright-community/playwright-go"
)

func RenderHTMLToPNG(
	html string,
	outputPath string,
) error {

	// Create outputs directory
	os.MkdirAll("outputs", os.ModePerm)

	// Temporary HTML file
	tempHTML := filepath.Join(
		"temp",
		"product_card.html",
	)

	os.MkdirAll("temp", os.ModePerm)

	// Save HTML
	err := os.WriteFile(
		tempHTML,
		[]byte(html),
		0644,
	)

	if err != nil {
		return err
	}

	// Install playwright
	err = playwright.Install()

	if err != nil {
		return err
	}

	// Start playwright
	pw, err := playwright.Run()

	if err != nil {
		return err
	}

	defer pw.Stop()

	// Launch browser
	browser, err := pw.Chromium.Launch()

	if err != nil {
		return err
	}

	defer browser.Close()

	// Create page
	page, err := browser.NewPage()

	if err != nil {
		return err
	}

	absPath, _ := filepath.Abs(tempHTML)

	// Open HTML
	_, err = page.Goto(
		"file://" + absPath,
	)

	if err != nil {
		return err
	}

	// Screenshot
	_, err = page.Screenshot(
		playwright.PageScreenshotOptions{
			Path: playwright.String(outputPath),
			FullPage: playwright.Bool(true),
		},
	)

	if err != nil {
		return err
	}

	return nil
}