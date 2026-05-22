package main

import (
	"fmt"
	"os"
	"strings"

	"ai-agent-system/agents"
	"ai-agent-system/config"
	"ai-agent-system/extractor"
	"ai-agent-system/renderer"
	"ai-agent-system/uploader"
)

func main() {

	// =====================================
	// LOAD ENV
	// =====================================
	config.LoadEnv()

	// =====================================
	// API KEYS
	// =====================================
	groqKey := config.GetEnv("GROQ_API_KEY")

	geminiKey := config.GetEnv("GEMINI_API_KEY")

	tripoKey := config.GetEnv("TRIPO_API_KEY")

	cloudName := config.GetEnv(
		"CLOUDINARY_CLOUD_NAME",
	)

	cloudKey := config.GetEnv(
		"CLOUDINARY_API_KEY",
	)

	cloudSecret := config.GetEnv(
		"CLOUDINARY_API_SECRET",
	)

	// =====================================
	// VALIDATION
	// =====================================
	if groqKey == "" {

		fmt.Println("❌ GROQ_API_KEY missing")
		return
	}

	if geminiKey == "" {

		fmt.Println("❌ GEMINI_API_KEY missing")
		return
	}

	if tripoKey == "" {

		fmt.Println("❌ TRIPO_API_KEY missing")
		return
	}

	if cloudName == "" ||
		cloudKey == "" ||
		cloudSecret == "" {

		fmt.Println("❌ Cloudinary keys missing")
		return
	}

	// =====================================
	// INPUT FILE
	// =====================================
	filePath := "uploads/product.pdf"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {

		fmt.Println("❌ File not found:", filePath)
		return
	}

	// =====================================
	// EXTRACT TEXT
	// =====================================
	var extractedText string

	if strings.HasSuffix(
		strings.ToLower(filePath),
		".pdf",
	) {

		text, err := extractor.ExtractTextFromPDF(
			filePath,
		)

		if err != nil {

			fmt.Println(
				"❌ PDF Extraction Error:",
				err,
			)

			return
		}

		extractedText = text

	} else {

		fmt.Println("❌ Unsupported file type")
		return
	}

	// =====================================
	// SHOW EXTRACTED TEXT
	// =====================================
	fmt.Println(
		"\n========== EXTRACTED TEXT ==========",
	)

	fmt.Println(extractedText)

	// =====================================
	// AGENT 1 -> ANALYSIS
	// =====================================
	analysis, err := agents.AnalyzeProduct(
		groqKey,
		extractedText,
	)

	if err != nil {

		fmt.Println(
			"❌ Analysis Error:",
			err,
		)

		return
	}

	fmt.Println(
		"\n========== PRODUCT ANALYSIS ==========",
	)

	fmt.Println(analysis)

	// =====================================
	// AGENT 2 -> GENERATE HTML
	// =====================================
	html, err := agents.GenerateProductCardHTML(
		geminiKey,
		analysis,
	)

	if err != nil {

		fmt.Println(
			"❌ HTML Generation Error:",
			err,
		)

		return
	}

	// =====================================
	// CREATE OUTPUT DIRECTORY
	// =====================================
	err = os.MkdirAll(
		"outputs",
		os.ModePerm,
	)

	if err != nil {

		fmt.Println(
			"❌ Failed creating outputs folder:",
			err,
		)

		return
	}

	// =====================================
	// SAVE HTML
	// =====================================
	htmlPath := "outputs/product_card.html"

	err = os.WriteFile(
		htmlPath,
		[]byte(html),
		0644,
	)

	if err != nil {

		fmt.Println(
			"❌ HTML Save Error:",
			err,
		)

		return
	}

	fmt.Println(
		"\n✅ HTML SAVED:",
		htmlPath,
	)

	// =====================================
	// AGENT 3 -> HTML TO PNG
	// =====================================
	imagePath := "outputs/product_card.png"

	err = renderer.RenderHTMLToPNG(
		html,
		imagePath,
	)

	if err != nil {

		fmt.Println(
			"❌ PNG Render Error:",
			err,
		)

		return
	}

	fmt.Println(
		"✅ PNG SAVED:",
		imagePath,
	)

	// =====================================
	// AGENT 4 -> CLOUDINARY UPLOAD
	// =====================================
	imageURL, err := uploader.UploadImage(
		cloudName,
		cloudKey,
		cloudSecret,
		imagePath,
	)

	if err != nil {

		fmt.Println(
			"❌ Cloudinary Upload Error:",
			err,
		)

		return
	}

	fmt.Println(
		"\n✅ PUBLIC IMAGE URL:",
	)

	fmt.Println(imageURL)

	// =====================================
	// AGENT 5 -> TRIPO 3D MODEL
	// =====================================
	modelURL, err := agents.Generate3DModel(
		tripoKey,
		imageURL,
	)

	if err != nil {

		fmt.Println(
			"❌ 3D Generation Error:",
			err,
		)

		return
	}

	// =====================================
	// FINAL RESULT
	// =====================================
	fmt.Println(
		"\n========== SUCCESS ==========",
	)

	fmt.Println(
		"\n✅ HTML FILE:",
	)

	fmt.Println(htmlPath)

	fmt.Println(
		"\n✅ PNG IMAGE:",
	)

	fmt.Println(imagePath)

	fmt.Println(
		"\n✅ CLOUDINARY URL:",
	)

	fmt.Println(imageURL)

	fmt.Println(
		"\n✅ 3D MODEL URL:",
	)

	fmt.Println(modelURL)
}