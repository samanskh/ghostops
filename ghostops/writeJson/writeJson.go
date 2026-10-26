package writejson

import (
	"encoding/json"
	"fmt"
	"os"
	// "path/filepath"

	"ghostops/analyze"
)

func WriteJson(stats analyze.Stats, path string) error {
	// path := filepath.Join(outDir, fileName)

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("there was a problem with the path and file name you gave")
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")

	if err := enc.Encode(stats); err != nil {
		return fmt.Errorf("encode stats: %w", err)
	}
	return nil
}