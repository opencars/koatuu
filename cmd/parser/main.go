package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/opencars/koatuu/pkg/config"
	"github.com/opencars/koatuu/pkg/model"
	"github.com/opencars/koatuu/pkg/store/sqlstore"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./config/config.toml", "Path to the configuration file")

	flag.Parse()

	path := flag.Arg(0)

	settings, err := config.New(configPath)
	if err != nil {
		log.Fatal(err)
	}

	store, err := sqlstore.New(&settings.DB)
	if err != nil {
		log.Fatal(err)
	}

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

	var levels []model.Level
	if err := json.Unmarshal(data, &levels); err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for i, level := range levels {
		fmt.Println(i, level)

		if level.ForthLevel != "" {
			if err := store.Level4().Create(level); err != nil {
				log.Fatal(err)
			}

			continue
		}

		if level.ThirdLevel != "" {
			if err := store.Level3().Create(level); err != nil {
				log.Fatal(err)
			}
			continue
		}

		if level.SecondLevel != "" {
			if err := store.Level2().Create(level); err != nil {
				log.Fatal(err)
			}
			continue
		}

		if err := store.Level1().Create(level); err != nil {
			log.Fatal(err)
		}

	}
}
