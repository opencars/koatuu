package parser

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/opencars/koatuu/pkg/koatuu"
)

func Parse(path string) (*koatuu.Level, error) {

	// Check ext. "json"
	if filepath.Ext(path) != ".json" {
		return nil, errors.New("extension is not valid")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var level koatuu.Level
	if json.Unmarshal(data, &level) != nil {
		return nil, err
	}

	defer file.Close()

	return &level, nil
}
