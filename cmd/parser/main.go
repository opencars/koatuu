package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/opencars/koatuu/pkg/config"
	"github.com/opencars/koatuu/pkg/domain/model"
	"github.com/opencars/koatuu/pkg/store/sqlstore"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./config/config.yaml", "Path to the configuration file")

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

	data, err := io.ReadAll(file)
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
			level3, err := model.NewThirdLevel(level.ThirdLevel, level.Name)
			if err != nil {
				log.Fatal(err)
			}

			if err := store.Level3().Create(level3); err != nil {
				log.Fatal(err)
			}

			continue
		}

		if level.SecondLevel != "" {
			level2, err := model.NewSecondLevel(level.SecondLevel, level.Name)
			if err != nil {
				log.Fatal(err)
			}

			if err := store.Level2().Create(level2); err != nil {
				log.Fatal(err)
			}

			continue
		}

		if err := store.Level1().Create(level); err != nil {
			log.Fatal(err)
		}

	}
}
