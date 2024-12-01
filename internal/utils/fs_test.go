package utils

import "testing"

func TestFileExists(t *testing.T) {
	exists := FileExists("fs.go")
	if !exists {
		t.Errorf("Expected file exists, received %v", exists)
	}

	exists = FileExists("someotherfile")
	if exists {
		t.Errorf("Expected file does not exist, received %v", exists)
	}

	exists = FileExists("./")
	if exists {
		t.Errorf("Expected file does not exist (since it is a directory), received %v", exists)
	}
}

func TestDirExists(t *testing.T) {
	exists := DirectoryExists("./")
	if !exists {
		t.Errorf("Expected dir exists, received %v", exists)
	}

	exists = DirectoryExists("./somedir")
	if exists {
		t.Errorf("Expected dir does not exist, received %v", exists)
	}

	exists = DirectoryExists("fs.go")
	if exists {
		t.Errorf("Expected dir does not exist (since it is a file), received %v", exists)
	}
}
