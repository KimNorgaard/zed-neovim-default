package vim

import (
	"encoding/json"
	"fmt"
	"io"
)

type Palette map[string]string

// LoadPalette reads a palette from the given reader.
func LoadPalette(r io.Reader) (Palette, error) {
	palette := make(Palette)

	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read palette: %w", err)
	}

	if err = json.Unmarshal(data, &palette); err != nil {
		return nil, err
	}

	return palette, nil
}
