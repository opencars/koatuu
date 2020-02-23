package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/opencars/koatuu/pkg/model"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./config/config.toml", "Path to the configuration file")

	flag.Parse()

	path := flag.Arg(0)

	// Check ext. "json"
	if filepath.Ext(path) != ".json" {
		log.Fatal("extension is not valid")
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var head model.Head
	if json.Unmarshal(data, &head) != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for _, lvl1 := range head.Level1 {
		// TODO: Insert into level1.
		for _, lvl2 := range lvl1.Level2 {
			// TODO: Insert into level2.
			for _, lv3 := range lvl2.Level3 {
				// TODO: Insert into level3.
				for _, lv4 := range lv3.Level4 {
					// TODO: Insert into level4.
					fmt.Println(lv4)
				}
			}
		}
	}

}
