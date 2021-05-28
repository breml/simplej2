package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar/v2"
	"github.com/flosch/pongo2"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) != 4 {
		log.Fatal("usage: simplej2 <value.json> <src> <destDir>")
	}

	valueFilename := os.Args[1]

	src, err := filepath.Abs(os.Args[2])
	if err != nil {
		log.Panic(err)
	}

	destDir, err := filepath.Abs(os.Args[3])
	if err != nil {
		log.Panic(err)
	}

	b, err := os.ReadFile(valueFilename)
	if err != nil {
		log.Panic(err)
	}

	values := pongo2.Context{}
	err = json.Unmarshal(b, &values)
	if err != nil {
		log.Panic(err)
	}

	oldWD, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	defer os.Chdir(oldWD)

	files := []string{filepath.Base(src)}
	err = os.Chdir(filepath.Dir(src))
	if err != nil {
		log.Panic(err)
	}

	fi, err := os.Stat(src)
	if err != nil {
		log.Panic(err)
	}
	if fi.IsDir() {
		err = os.Chdir(src)
		if err != nil {
			log.Panic(err)
		}

		files, err = doublestar.Glob(filepath.Join("**", "*.j2"))
		if err != nil {
			log.Panic(err)
		}
	}

	for _, file := range files {
		fi, err := os.Stat(file)
		if err != nil {
			log.Panic(err)
		}
		if fi.IsDir() {
			continue
		}

		tpl, err := pongo2.DefaultSet.FromFile(file)
		if err != nil {
			log.Panic(err)
		}

		out, err := tpl.ExecuteBytes(values)
		if err != nil {
			log.Panic(err)
		}

		err = os.MkdirAll(filepath.Join(destDir, filepath.Dir(file)), 0700)
		if err != nil {
			log.Panic(err)
		}

		targetFile := strings.TrimSuffix(file, ".j2")
		err = os.WriteFile(filepath.Join(destDir, targetFile), out, 0600)
		if err != nil {
			log.Panic(err)
		}
	}
}
