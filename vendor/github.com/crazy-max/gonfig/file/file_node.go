package file

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/crazy-max/gonfig/parser"
	"gopkg.in/yaml.v3"
)

// decodeFileToNode decodes the configuration in filePath in a tree of untyped nodes.
// If filters is not empty, it skips any configuration element whose name is not among filters.
func decodeFileToNode(filePath string, filters ...string) (*parser.Node, error) {
	content, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})

	switch strings.ToLower(filepath.Ext(filePath)) {
	case ".toml":
		err = toml.Unmarshal(content, &data)
		if err != nil {
			return nil, err
		}

	case ".yml", ".yaml", ".json":
		err = yaml.Unmarshal(content, data)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("unsupported file extension: %s", filePath)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("no configuration found in file: %s", filePath)
	}

	node, err := decodeRawToNode(data, filters...)
	if err != nil {
		return nil, err
	}

	if len(node.Children) == 0 {
		return nil, fmt.Errorf("no valid configuration found in file: %s", filePath)
	}

	return node, nil
}

func getRootFieldNames(element interface{}) []string {
	if element == nil {
		return nil
	}

	rootType := reflect.TypeOf(element)

	return getFieldNames(rootType)
}

func getFieldNames(rootType reflect.Type) []string {
	var names []string

	if rootType.Kind() == reflect.Pointer {
		rootType = rootType.Elem()
	}

	if rootType.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < rootType.NumField(); i++ {
		field := rootType.Field(i)

		if !parser.IsExported(field) {
			continue
		}

		if field.Anonymous &&
			(field.Type.Kind() == reflect.Pointer && field.Type.Elem().Kind() == reflect.Struct || field.Type.Kind() == reflect.Struct) {
			names = append(names, getFieldNames(field.Type)...)
			continue
		}

		names = append(names, field.Name)
	}

	return names
}
