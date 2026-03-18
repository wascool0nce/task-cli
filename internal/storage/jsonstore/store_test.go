package jsonstore

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStoreLoad_FileNotExists(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "not_found.json")

	store := New(path)

	data, err := store.Load()

	if err != nil {
		t.Fatalf("Load() unexpected error: %v", err)
	}

	if data == nil {
		t.Fatalf("Load() returned nil slice, want empty slice")
	}

	if len(data) != 0 {
		t.Errorf("len(Load()) = %d, want 0", len(data))
	}
}

func TestStore_Load_InvalidJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")

	err := os.WriteFile(path, []byte(`{invalid json}`), 0644)

	if err != err {
		t.Fatalf("failed to write invalid json: %v", err)
	}

	store := New(path)

	data, err := store.Load()

	if err == nil {
		t.Fatalf("Load() unexpected error got nil")
	}

	if data != nil {
		t.Errorf("Load = %#v, want nil", data)
	}
}
