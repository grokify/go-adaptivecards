package demo

import (
	"encoding/json"
	"io/fs"
	"os"
)

// WriteJsonFiles is a helper that marshals data two rote both
// minimized and indented JSON files.
func WriteJsonFiles(filename, filenameIndent, prefix, indent string, data interface{}, perm fs.FileMode) error {
	if len(filename) > 0 {
		bytesMin, err := json.Marshal(data)
		if err != nil {
			return err
		}
		err = os.WriteFile(filename, bytesMin, perm)
		if err != nil {
			return err
		}
	}
	if len(filenameIndent) > 0 {
		bytesIndent, err := json.MarshalIndent(data, prefix, indent)
		if err != nil {
			return err
		}
		err = os.WriteFile(filenameIndent, bytesIndent, perm)
		if err != nil {
			return err
		}
	}
	return nil
}
