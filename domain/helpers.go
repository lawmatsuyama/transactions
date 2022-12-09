package domain

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// ReadJSON read json from file and set object
func ReadJSON(file string, object interface{}) error {
	b, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		return err
	}

	return json.Unmarshal(b, object)

}

func CreateJSON(file string, object interface{}) (err error) {
	b, err := json.MarshalIndent(object, "", "\t")
	if err != nil {
		return
	}

	return os.WriteFile(file, b, 0644)
}

func Compare(t *testing.T, description string, exp, got any, opts ...cmp.Option) {
	d := cmp.Diff(exp, got, opts...)
	if len(d) > 0 {
		t.Fatalf("test [%s] compare description [%s] mismatch (-want +got):\n%s", t.Name(), description, d)
	}
}
