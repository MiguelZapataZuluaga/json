package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"os"
)

func Tree() {

    
	root := "1_.txt"
	jsonData := make(map[string]interface{})

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error walking directory: ", err)
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("Error reading file: ", err)
				return err
			}
			str_content := string(data)
			jsonData[info.Name()] = str_content
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking directory: ", err)
		return
	}

	//jsonStr, err := json.MarshalIndent(jsonData, "", "  ")
	jsonStr, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error converting to JSON: ", err)
		return
	}

	fmt.Println(string(jsonStr))
}

func main() {
	Tree()
}
