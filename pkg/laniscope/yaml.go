package laniscope

import (
	"bytes"
	"gopkg.in/yaml.v3"
)

// Decode decodes one Laniscope YAML document from data.
func Decode(data []byte) (Laniscope, error) {
	var scope Laniscope
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)

	if err := decoder.Decode(&scope); err != nil {
		return Laniscope{}, err
	}

	if err := scope.Validate(); err != nil {
		return Laniscope{}, err
	}

	return scope, nil
}

// Encode encodes a valid Laniscope as one YAML document.
func (s Laniscope) Encode() ([]byte, error) {
	if err := s.Validate(); err != nil {
		return nil, err
	}

	var data bytes.Buffer
	encoder := yaml.NewEncoder(&data)
	encoder.SetIndent(2)

	if err := encoder.Encode(s); err != nil {
		return nil, err
	}
	if err := encoder.Close(); err != nil {
		return nil, err
	}

	return data.Bytes(), nil
}
