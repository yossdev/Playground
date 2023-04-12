package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	start := time.Now()

	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	dir := "dst"

	if cmd == "clean" {
		pathDir := fmt.Sprintf("./%s", dir)
		cleanUp(pathDir)
	} else {
		count := 10000
		createFile(count, dir)
	}

	fmt.Printf("Time took in go: %s\n", time.Since(start))
}

func cleanUp(path string) {
	size, _ := dirSize(path)
	os.RemoveAll(path)

	if err := os.Mkdir(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total deleted: %d bytes\n", size)
}

func createFile(count int, dir string) {
	start := time.Now()

	var sumNBytes int64

	src, _ := os.Open("./src/test.md")
	defer src.Close()

	var fn string
	var dst *os.File
	var nBytes int64

	for i := 0; i < count; i++ {
		fn = fmt.Sprintf("./%s/test%d.md", dir, i)
		dst, _ = os.Create(fn)
		defer dst.Close()

		nBytes, _ = io.Copy(dst, src)

		sumNBytes += nBytes
	}

	fmt.Printf("Time took in go createFile: %s\n", time.Since(start))
	fmt.Printf("Total written: %d bytes\n", sumNBytes)
}

func dirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
