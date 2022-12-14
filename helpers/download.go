package helpers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Download(fullURLFile string, filepath string) {
	// Create blank file
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d\n", filepath, size)
}
