package config

import (
	"fmt"
	"io"
	"j4f/core/json"
	"j4f/core/toml"
	"j4f/core/yaml"
	"os"
	"path/filepath"
	"strings"
)

var decoders = []decoder{
	{toml.GetExt(), toml.Decode},
	{yaml.GetExt(), yaml.Decode},
	{json.GetExt(), json.Decode},
}

func ParseFile(file string, out interface{}) error {
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err == nil {
		return parseFile(f, out)
	}

	if filepath.IsAbs(file) {
		return err
	}

	file = filepath.Join(`config`, file)
	f, err = os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	return parseFile(f, out)
}

func parseFile(file *os.File, out interface{}) error {
	return encode(file, filepath.Ext(file.Name()), out)
}

type decodeHandler func(r io.Reader, out interface{}) error
type decoder struct {
	exts       []string
	decodeFunc decodeHandler
}

func encode(r io.Reader, ext string, out interface{}) error {
	h := getEncoder(ext)
	if h == nil {
		return fmt.Errorf("no support encode config with %s extension.", ext)
	}
	return h(r, out)
}

func getEncoder(ext string) decodeHandler {
	for _, e := range decoders {
		if fitExt(ext, e.exts) {
			return e.decodeFunc
		}
	}
	return nil
}

func fitExt(ext string, fits []string) bool {
	ext = strings.TrimLeft(ext, ".")

	for _, fit := range fits {
		if fit == ext {
			return true
		}
	}
	return false
}
