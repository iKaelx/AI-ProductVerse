# AI-ProductVerse

AI-powered multi-agent system that transforms product documents into:

- AI product analysis
- Modern marketing product cards
- High-quality rendered product images
- AI-generated 3D models

Built with Go, Groq, Gemini, Playwright, Cloudinary, and Tripo AI.

---

# Features

## Multi-Agent AI Architecture

### Agent 1 — Product Analyzer
Extracts and analyzes product information from PDFs using Groq LLM.

### Agent 2 — Product Card Generator
Generates a modern HTML marketing product card using Gemini AI.

### Agent 3 — Renderer
Converts generated HTML into a high-quality PNG using Playwright.

### Agent 4 — Cloud Uploader
Uploads generated assets to Cloudinary.

### Agent 5 — 3D Model Generator
Creates AI-generated 3D models from product images using Tripo AI.

---

# Architecture

```text
PDF
 ↓
Text Extraction
 ↓
Groq Analysis
 ↓
Gemini HTML Generation
 ↓
Playwright PNG Rendering
 ↓
Cloudinary Upload
 ↓
Tripo AI
 ↓
GLB / OBJ / FBX 3D Model
```

---

# Tech Stack

| Technology | Usage |
|---|---|
| Go | Backend |
| Groq | Product Analysis |
| Gemini | HTML/UI Generation |
| Playwright | HTML Rendering |
| Cloudinary | Image Hosting |
| Tripo AI | 3D Generation |
| PDF Extraction | OCR/Text Parsing |

---

# Project Structure

```text
ai-productverse/
│
├── agents/
│   ├── groq_agent.go
│   ├── gemini_agent.go
│   └── model3d_agent.go
│
├── config/
│
├── extractor/
│
├── renderer/
│
├── uploader/
│
├── uploads/
│   └── product.pdf
│
├── outputs/
│   ├── product_card.html
│   └── product_card.png
│
├── .env
├── go.mod
└── main.go
```

---

# Installation

## Clone Repository

```bash
git clone https://github.com/YOUR_USERNAME/AI-ProductVerse.git

cd AI-ProductVerse
```

---

# Install Dependencies

```bash
go get github.com/playwright-community/playwright-go

go get github.com/cloudinary/cloudinary-go/v2

go get github.com/google/generative-ai-go/genai

go get google.golang.org/api/option
```

---

# Install Playwright

```bash
go run github.com/playwright-community/playwright-go/cmd/playwright install
```

---

# Linux Dependencies

```bash
sudo npx playwright install-deps
```

If npx is missing:

```bash
sudo apt install nodejs npm -y
```

---

# Environment Variables

Create `.env`

```env
GROQ_API_KEY=xxxx

GEMINI_API_KEY=xxxx

TRIPO_API_KEY=xxxx

CLOUDINARY_CLOUD_NAME=xxxx
CLOUDINARY_API_KEY=xxxx
CLOUDINARY_API_SECRET=xxxx
```

---

# APIs

## Groq

https://console.groq.com/

## Gemini

https://aistudio.google.com/

## Cloudinary

https://cloudinary.com/

## Tripo AI

https://www.3daistudio.com/

---

# Usage

Put your PDF inside:

```text
uploads/product.pdf
```

Run:

```bash
go run .
```

---

# Output

Generated files:

```text
outputs/product_card.html

outputs/product_card.png
```

Console output:

```text
✅ Product Analysis

✅ HTML Generated

✅ PNG Generated

✅ Cloudinary Upload Complete

✅ 3D Model Generated
```

---

# Example Workflow

1. Upload PDF product documentation
2. AI extracts product details
3. AI creates marketing product card
4. HTML rendered into PNG
5. Image uploaded to cloud
6. AI generates 3D model
7. Final GLB/OBJ returned

---

# Future Improvements

- REST API
- Next.js frontend
- Drag & drop UI
- AI video ads
- AI voice assistant
- Shopify integration
- AR viewer
- Blender automation
- Vector database memory
- Multi-modal agents

---

# Screenshots

Add generated examples here.

---

# License

MIT License

---

# Author

Raouf Dahmani

AI Engineer & Multi-Agent Systems Developer
