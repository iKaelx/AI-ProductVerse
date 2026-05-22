package extractor

import (
    "bytes"
    "os/exec"
)

func ExtractTextFromImage(imagePath string) (string, error) {

    cmd := exec.Command("tesseract", imagePath, "stdout")

    var out bytes.Buffer
    cmd.Stdout = &out

    err := cmd.Run()
    if err != nil {
        return "", err
    }

    return out.String(), nil
}