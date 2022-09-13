package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/dhowden/tag"
)

func main() {
	var dirFlag = flag.String("p", ".", "The directory to scan for files in.")
	flag.Parse()
	items, err := ioutil.ReadDir(*dirFlag)
	if err != nil {
		log.Fatal(err)
	}

	var albums map[string]bool = make(map[string]bool)

	for _, item := range items {
		if item.IsDir() {
			continue
		}
		fullPath := *dirFlag + "/" + item.Name()
		file, err := os.Open(fullPath)
		if err != nil {
			log.Fatal(err)
		}
		f, err := tag.ReadFrom(file)
		if err != nil {
			continue
		}
		albums[f.Album()] = true
		album := f.Album()
		if album == "" {
			album = "misc"
		}
		err = os.Mkdir(*dirFlag+"/"+album, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		err = os.Rename(fullPath, *dirFlag+"/"+album+"/"+item.Name())
		if err != nil {
			log.Fatal(err)
		}
	}
}
