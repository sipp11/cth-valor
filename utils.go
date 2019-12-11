package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
		// os.Exit(1)
	}
}

func ReadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	CheckError(err)
	return data
}

func HTTPClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func GetJsonIn(root string) []string {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if ext == ".json" {
			files = append(files, path)
		}
		return nil
	})
	CheckError(err)
	return files
}
