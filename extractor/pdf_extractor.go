package extractor

import (
    "bytes"
    "fmt"
    "os/exec"
)

func ExtractTextFromPDF(pdfPath string) (string, error) {

    cmd := exec.Command("pdftotext", pdfPath, "-")

    var out bytes.Buffer
    var stderr bytes.Buffer

    cmd.Stdout = &out
    cmd.Stderr = &stderr

    err := cmd.Run()
    if err != nil {
        return "", fmt.Errorf(stderr.String())
    }

    return out.String(), nil
}