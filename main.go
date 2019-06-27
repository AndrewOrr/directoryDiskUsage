package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fileStats, err := findDirPath("C:/Go")
	if err != nil {
		fmt.Errorf("Failed to read file information", err)
	}
	js, err := json.MarshalIndent(fileStats, "", "    ")
	if err != nil {
		fmt.Errorf("Failed to create json interpretation of map", err)
	}
	fmt.Println(string(js))

}

func findDirPath(dir string) (map[string]int64, error) {
	fileStats := make(map[string]int64)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, err
	}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fileStats[path] = info.Size()
		return nil
	})
	return fileStats, err
}
