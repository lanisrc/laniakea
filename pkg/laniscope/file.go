package laniscope

import "os"

// Load reads and decodes one Laniscope from filename.
func Load(filename string) (Laniscope, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Laniscope{}, err
	}
	return Decode(data)
}

// Save encodes and writes one Laniscope to filename.
func (s Laniscope) Save(filename string) error {
	data, err := s.Encode()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0o644)
}
