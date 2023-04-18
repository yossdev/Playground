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

	var cmdArgs string
	if len(os.Args) > 1 {
		cmdArgs = os.Args[1]
	}

	const root string = "D:/Code/Playground"
	const cwd string = "go/copy-file"
	const dir string = "dst"

	if cmdArgs == "clean" {
		pathDir := fmt.Sprintf("%s/%s/%s", root, cwd, dir)
		cleanUp(pathDir)
	} else {
		const count = 10_000
		src := fmt.Sprintf("%s/common/src/test.md", root)
		copyFile(count, src, dir)
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

func copyFile(count int, src, dir string) {
	start := time.Now()

	var sumNBytes int64

	file, _ := os.Open(src)
	defer file.Close()

	var fn string
	var dst *os.File
	var nBytes int64

	for i := 0; i < count; i++ {
		fn = fmt.Sprintf("./%s/test%d.md", dir, i)
		dst, _ = os.Create(fn)
		defer dst.Close()

		file.Seek(0, io.SeekStart)
		nBytes, _ = io.Copy(dst, file)

		sumNBytes += nBytes
	}

	fmt.Printf("Time took in go copyFile: %s\n", time.Since(start))
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
