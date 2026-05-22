package agents

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type ProductCardGenerator struct {
	APIKey string
}

// GenerateProductCardImage creates a professional product card image
func (p *ProductCardGenerator) GenerateProductCardImage(productData string) (string, error) {

	ctx := context.Background()

	// Create Gemini client
	client, err := genai.NewClient(
		ctx,
		option.WithAPIKey(p.APIKey),
	)

	if err != nil {
		return "", fmt.Errorf("failed to create client: %v", err)
	}

	defer client.Close()

	// IMPORTANT:
	// Use image generation model
	model := client.GenerativeModel("gemini-2.5-flash-image")

	// Better image generation quality
	model.SetTemperature(0.8)

	prompt := `
Create a PREMIUM modern ecommerce product card.

STYLE:
- Professional
- Clean UI
- Apple style design
- Dark luxury background
- High quality product showcase
- Modern typography
- Gradient effects
- Glassmorphism
- Tech aesthetic

INCLUDE:
1. Product title
2. Marketing slogan
3. Main features section
4. Technical specifications
5. Product image area
6. Pricing section
7. Premium modern layout
8. Beautiful spacing
9. Icons for features

OUTPUT:
Generate ONE complete professional product card image.

PRODUCT DATA:
` + productData

	// Generate content
	resp, err := model.GenerateContent(
		ctx,
		genai.Text(prompt),
	)

	if err != nil {
		return "", fmt.Errorf("generation failed: %v", err)
	}

	// Save generated image
	for _, candidate := range resp.Candidates {

		if candidate.Content == nil {
			continue
		}

		for _, part := range candidate.Content.Parts {

			// Image data
			if blob, ok := part.(genai.Blob); ok {

				fileName := "generated_product_card.png"

				err := os.WriteFile(
					fileName,
					blob.Data,
					0644,
				)

				if err != nil {
					return "", fmt.Errorf("failed to save image: %v", err)
				}

				return fileName, nil
			}

			// Alternative image format
			if data, ok := part.(genai.FileData); ok {

				fmt.Println("Generated file URI:", data.URI)
			}
		}
	}

	return "", fmt.Errorf("no image generated")
}