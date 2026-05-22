package agents

import (
	"context"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerateProductCardHTML(
	apiKey string,
	productData string,
) (string, error) {

	ctx := context.Background()

	client, err := genai.NewClient(
		ctx,
		option.WithAPIKey(apiKey),
	)

	if err != nil {
		return "", err
	}

	defer client.Close()

	model := client.GenerativeModel(
		"gemini-2.5-flash",
	)

	prompt := `
Create ONLY a professional ecommerce product card HTML page.

RULES:
- Return ONLY HTML
- Include CSS inside <style>
- Dark premium UI
- Responsive design
- Modern typography
- Product features section
- Technical specifications
- Marketing section
- Professional ecommerce look
- Glassmorphism
- Beautiful spacing	
- NO markdown
- NO explanations

PRODUCT:
` + productData

	resp, err := model.GenerateContent(
		ctx,
		genai.Text(prompt),
	)

	if err != nil {
		return "", err
	}

	text := resp.Candidates[0].
		Content.Parts[0].(genai.Text)

	return string(text), nil
}