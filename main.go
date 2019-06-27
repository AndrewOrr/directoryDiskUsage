package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please specify a directory.")
		return
	}
	var err error
	var sb strings.Builder
	sb.WriteString("{\n\t")
	sb.WriteString(`"files":`)
	sb.WriteString("\n\t[\n")
	for _, dir := range args {
		err = findDirPath(dir, &sb)
		if err != nil {
			fmt.Println(fmt.Errorf("Failed to read file information", err).Error())
		}
	}
	sb.WriteString("\t],\n}")
	fmt.Println(sb.String())

}

func findDirPath(dir string, sb *strings.Builder) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return err
	}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		sb.WriteString("\t ")
		sb.WriteString(fmt.Sprintf(`{"%s", %d},`, path, int(info.Size())))
		sb.WriteString("\n")
		return nil
	})
	return err
}
