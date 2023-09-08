package main

import (
	"github.com/Shackelford-Arden/BookBrowser/models"
	"github.com/Shackelford-Arden/BookBrowser/pkg/cli"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"

	// These are required to initialize/Register each supported type
	_ "github.com/Shackelford-Arden/BookBrowser/pkg/formats/epub"
	_ "github.com/Shackelford-Arden/BookBrowser/pkg/formats/mobi"
	_ "github.com/Shackelford-Arden/BookBrowser/pkg/formats/pdf"
)

func main() {

	var config models.Config
	configErr := envconfig.Process("bookbrowser", &config)
	if configErr != nil {
		log.Fatal(configErr.Error())
	}

	app := cli.StartCLI()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

//func checkUpdate() {
//	resp, err := http.Get("https://api.github.com/repos/Shackelford-Arden/BookBrowser/releases/latest")
//	if err != nil {
//		return
//	}
//	defer resp.Body.Close()
//
//	buf, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return
//	}
//
//	if resp.StatusCode != 200 {
//		return
//	}
//
//	var obj struct {
//		URL string `json:"html_url"`
//		Tag string `json:"tag_name"`
//	}
//	if json.Unmarshal(buf, &obj) != nil {
//		return
//	}
//	if curversion != "dev" {
//		if !strings.HasPrefix(curversion, obj.Tag) {
//			log.Printf("Running version %s. Latest version is %s: %s\n", curversion, obj.Tag, obj.URL)
//		}
//	}
//}
