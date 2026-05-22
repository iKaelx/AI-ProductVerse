package agents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Generate3DModel(
	apiKey string,
	imageURL string,
) (string, error) {

	// =====================================
	// CREATE TASK
	// =====================================

	payload := map[string]interface{}{
		"image_url": imageURL,
		"texture":   true,
		"pbr":       true,
	}

	jsonData, _ := json.Marshal(payload)

	req, err := http.NewRequest(
		"POST",
		"https://api.3daistudio.com/v1/3d-models/tripo/image-to-3d/",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return "", err
	}

	req.Header.Set(
		"Authorization",
		"Bearer "+apiKey,
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("========== TRIPO RESPONSE ==========")
	fmt.Println(string(body))

	var result map[string]interface{}

	json.Unmarshal(body, &result)

	taskID, ok := result["task_id"].(string)

	if !ok {

		return "", fmt.Errorf(
			"failed to create task",
		)
	}

	fmt.Println("✅ TASK ID:", taskID)

	// =====================================
	// POLLING
	// =====================================

	for {

		time.Sleep(10 * time.Second)

		req, _ := http.NewRequest(
			"GET",
			"https://api.3daistudio.com/v1/tasks/"+taskID+"/",
			nil,
		)

		req.Header.Set(
			"Authorization",
			"Bearer "+apiKey,
		)

		resp, err := client.Do(req)

		if err != nil {
			return "", err
		}

		body, _ := io.ReadAll(resp.Body)

		fmt.Println("========== STATUS ==========")
		fmt.Println(string(body))

		var task map[string]interface{}

		json.Unmarshal(body, &task)

		status, _ := task["status"].(string)

		fmt.Println("⏳ STATUS:", status)

		if status == "success" {

			output := task["output"].(map[string]interface{})

			modelURL := output["model"].(string)

			return modelURL, nil
		}

		if status == "failed" {

			return "", fmt.Errorf(
				"3D generation failed",
			)
		}
	}
}