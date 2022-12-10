package domain

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// ReadJSON read json from file and set object
func ReadJSON(t *testing.T, file string, object any) {
	b, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		t.Fatal("failed to read file")
	}

	err = json.Unmarshal(b, object)
	if err != nil {
		t.Fatal("failed to unmarshal object")
	}
}

// Read read json from file and return bytes
func Read(t *testing.T, file string) []byte {
	b, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		t.Fatal("failed to read file")
	}
	return b
}

// CreateJSON creates a json file using the given object
func CreateJSON(t *testing.T, file string, object any) {
	b, err := json.MarshalIndent(object, "", "\t")
	if err != nil {
		t.Fatal("failed to create file")
	}

	/* #nosec */
	err = os.WriteFile(file, b, 0644)
	if err != nil {
		t.Fatal("failed to write json file")
	}
}

// Compare it compares two objects and check if they are equals
func Compare(t *testing.T, description string, exp, got any, opts ...cmp.Option) {
	d := cmp.Diff(exp, got, opts...)
	if len(d) > 0 {
		t.Fatalf("test [%s] compare description [%s] mismatch (-want +got):\n%s", t.Name(), description, d)
	}
}
