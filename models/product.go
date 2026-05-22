package models

type Product struct {
    Name        string   `json:"name"`
    Brand       string   `json:"brand"`
    Price       string   `json:"price"`
    Description string   `json:"description"`
    Features    []string `json:"features"`
    Specifications map[string]string `json:"specifications"`
    Images      []string `json:"images"`
}