package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"
)

type Secret struct {
	Filename string            `json:"filename"`
	Keys     map[string]string `json:"keys"`
}

func main() {
	var secrets []Secret

	secretTemplate := flag.String("secrets", "", "secrets for replacing")

	err:=json.Unmarshal([]byte(*secretTemplate), &secrets)

	if err!=nil{
		log.Fatal(err)
	}

	for _, v := range secrets {
		err=PasteSecret(v.Filename, v.Keys)
		if err!=nil{
			log.Fatal(err)
		}
	}

}

func PasteSecret(filename string, keys map[string]string) error {

	file, err := os.ReadFile(filename)
	text := string(file)

	if err != nil {
		return err
	}

	for key, value := range keys {
		text = strings.ReplaceAll(text, key, value)
	}

	err = os.WriteFile(filename, []byte(text), os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
