package main

import (
	"fmt"
	"os"
	"path/filepath"
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
