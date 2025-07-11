package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func ExportFile(data any, fileName string) error {

	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return err
	}

	err = os.WriteFile(fileName, prettyJSON, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	fmt.Printf("File written to %s\n", fileName)
	return nil
}

func ExportPlainText(data []string, fileName string) error {
	var err error
	content := strings.Join(data, "\n")

	err = os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	fmt.Printf("File written to %s\n", fileName)
	return nil
}
