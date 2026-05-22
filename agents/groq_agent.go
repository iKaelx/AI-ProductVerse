package agents

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type GroqResponse struct {
    Choices []struct {
        Message struct {
            Content string `json:"content"`
        } `json:"message"`
    } `json:"choices"`
}

func AnalyzeProduct(apiKey string, content string) (string, error) {

    url := "https://api.groq.com/openai/v1/chat/completions"

    payload := map[string]interface{}{
        "model": "llama-3.3-70b-versatile",
        "messages": []map[string]string{
            {
                "role": "system",
                "content": "You are an expert product analyzer. Extract all information about the product.",
            },
            {
                "role": "user",
                "content": content,
            },
        },
        "temperature": 0.3,
    }

    jsonData, err := json.Marshal(payload)
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return "", err
    }

    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}

    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    fmt.Println("========== GROQ RAW RESPONSE ==========")
    fmt.Println(string(body))

    var result GroqResponse

    err = json.Unmarshal(body, &result)
    if err != nil {
        return "", err
    }

    if len(result.Choices) == 0 {
        return "", fmt.Errorf("Groq returned no choices")
    }

    return result.Choices[0].Message.Content, nil
}