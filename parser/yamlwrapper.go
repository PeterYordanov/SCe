package parser

import (
	"gopkg.in/yaml.v3"
)

type YamlWrapper[T any] struct {
	data T
}

func NewYamlWrapper[T any]() *YamlWrapper[T] {
	return &YamlWrapper[T]{}
}

// ToString marshals the internal data of the wrapper into a YAML string.
func (y *YamlWrapper[T]) ToString() (string, error) {
	out, err := yaml.Marshal(y.data)
	if err != nil {
		return "Marshalling YAML Unsuccessful", err
	}
	return string(out), nil
}

// Parse unmarshals a YAML string into the internal data of the wrapper.
func (y *YamlWrapper[T]) Parse(input string) error {
	return yaml.Unmarshal([]byte(input), &y.data)
}
