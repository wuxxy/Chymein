package System

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"server/internal/Common"
)

type File struct {
	Path    string
	Name    string
	Purpose string
	Content []byte
}

var (
	Files           []File
	ErrFileNotFound = errors.New("file not found")
)

// Open reads file content and fills the File struct
func (f *File) Open(path string) (*File, error) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, ErrFileNotFound
		}
		Common.Warn("Error occurred while trying to open file: "+path, err, "File")
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			Common.Warn("Error occurred while trying to close file: "+path, err, "File")
			return
		}
	}(file)

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	*f = File{
		Path:    path,
		Name:    filepath.Base(path),
		Content: content,
	}
	Files = append(Files, *f)
	return f, nil
}

// Create writes a new file with content and metadata
func (f *File) Create(path string, content []byte, purpose string) (*File, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	if _, err := file.Write(content); err != nil {
		return nil, err
	}

	*f = File{
		Path:    path,
		Name:    filepath.Base(path),
		Purpose: purpose,
		Content: content,
	}
	Files = append(Files, *f)
	return f, nil
}

// ParseJSON unmarshals File content into the provided struct
func (f *File) ParseJSON(target interface{}) error {
	if len(f.Content) == 0 {
		return errors.New("no file content to parse")
	}
	return json.Unmarshal(f.Content, target)
}
func (f *File) AddLine(line string) error {
	file, err := os.OpenFile(f.Path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	formatted := fmt.Sprintf("%s\n", line)

	if _, err := file.WriteString(formatted); err != nil {
		return err
	}

	// Update in-memory content
	f.Content = append(f.Content, []byte(formatted)...)
	return nil
}

// Exists returns true if the file physically exists on disk
func (f *File) Exists() bool {
	_, err := os.Stat(f.Path)
	return err == nil
}
