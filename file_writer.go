package main

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func WriteToJson(server string, filename string, txt []byte) {
	today := time.Now()
	path := filepath.Join("./data", server, today.Format("2006-01-02"))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.FileMode.Perm(0755))
	}
	fp := filepath.Join(path, fmt.Sprintf("%s.json", filename))
	// write to file too
	f, err := os.Create(fp)
	if err == nil {
		f.Write(txt)
		defer f.Close()
	}
}

func MediaDownload(server string, day string, what string, targetUrl string) bool {
	path := filepath.Join("./data", server, day, what)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.FileMode.Perm(0755))
	}
	client := HTTPClient()
	if targetUrl == "" {
		return false
	}
	// fmt.Println(targetUrl)
	if targetUrl[0] == '/' {
		targetUrl = fmt.Sprintf("https:%s", targetUrl)
	}
	// Check if we already have it
	fileName := buildFileName(targetUrl)
	filePath := filepath.Join(path, fileName)
	// if we found, then don't re-download
	if _, err := os.Stat(filePath); err == nil {
		return false
	}

	// initiate download
	resp, err := client.Get(targetUrl)
	CheckError(err)
	defer resp.Body.Close()

	file, err := os.Create(filePath)
	CheckError(err)

	_, err = io.Copy(file, resp.Body) // return size & err
	// fmt.Printf("%s - size %d", filePath, size)
	CheckError(err)

	defer file.Close()
	return true
}

func buildFileName(targetUrl string) string {
	fileUrl, err := url.Parse(targetUrl)
	CheckError(err)
	path := fileUrl.Path
	segments := strings.Split(path, "/")
	return segments[len(segments)-1]
}
