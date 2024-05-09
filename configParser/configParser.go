package configParser

import (
	"bytes"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

type ConfigParser struct {
	Endpoints []SimpleHandler `yaml:"endpoints"`
}

type SimpleHandler struct {
	Endpoint string `yaml:"endpoint"`
	Content  string `yaml:"content"`
}

func (cP ConfigParser) String() string {
	var buffer bytes.Buffer
	for _, handler := range cP.Endpoints {
		buffer.WriteString(fmt.Sprintf("{endpoint: %s, content: %s},\n", handler.Endpoint, handler.Content))
	}
	return buffer.String()
}

func NewConfigParser(reader io.Reader) (ConfigParser, error) {
	buf := make([]byte, 1024)
	n, err := reader.Read(buf)
	if err != nil {
		return ConfigParser{}, err
	}
	var config ConfigParser
	if err := yaml.Unmarshal(buf[:n], &config); err != nil {
		return ConfigParser{}, err
	}
	return config, nil
}
