package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Key struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Secret struct {
	Filename string `json:"filename"`
	Keys     []Key  `json:"keys"`
}

func main() {
	var secrets []Secret

	secretTemplate := flag.String("secrets", "", "Replace secrets with given keys")
	flag.Parse()

	err := json.Unmarshal([]byte(*secretTemplate), &secrets)

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range secrets {
		err = PasteSecret(v.Filename, v.Keys)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func PasteSecret(filename string, keys []Key) error {
	filename, _ = filepath.Abs(filename)
	file, err := os.ReadFile(filename)
	text := string(file)

	if err != nil {
		return err
	}

	for _, key := range keys {
		text = strings.ReplaceAll(text, fmt.Sprintf("{{%s}}", key.Key), key.Value)
	}

	err = os.WriteFile(filename, []byte(text), os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
