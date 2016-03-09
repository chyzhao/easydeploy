package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("prease use cammand \"deploy create -f [filepath]\" to deploy your apps")
		os.Exit(1)
	} else if os.Args[1] != "create" || os.Args[2] != "-f" {
		fmt.Println("parameters error")
		fmt.Println("prease use cammand \"deploy create -f [filepath]\" to deploy your apps")
		os.Exit(1)
	} else {
		filePath := os.Args[3] + "/EASYDEPLOY"
		yamlFilePaths, err := readFile(filePath)
		fmt.Printf("yaml file length %d\n", len(yamlFilePaths))
		if err != nil {
			fmt.Println("can not open file " + filePath + ", does it exist?")
			os.Exit(1)
		}
		for _, yamlFilePath := range yamlFilePaths {
			fmt.Println("yamlFilePath" + yamlFilePath)
			//cmd := exec.Command("kubectl", "create", "-f", yamlFilePath)
			//if err := cmd.Run(); err != nil {
			//fmt.Println("exec error")
			//os.Exit(1)
			//}
		}
		fmt.Println("succeed!")
	}
}

func readFile(filePath string) ([]string, error) {
	result := make([]string, 0, 10)
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				return result, nil
			}
			return nil, err
		} else {
			result = append(result, line)
		}
	}
}
