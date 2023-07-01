package tater

import (
	"errors"
	"os"
	"path/filepath"
)

var ErrWontErase = errors.New("won't overwrite with zero byte")

type Tater struct {
	filename string
	file     *os.File
}

func New(filename string) *Tater {
	t := &Tater{
		filename: filename,
	}

	if err := t.Rotate(); err != nil {
		panic(err)
	}

	return t
}

// Write implements io.Writer
func (t *Tater) Write(b []byte) (n int, err error) {
	return t.file.Write(b)
}

func (t *Tater) Rotate() error {
	if t.file != nil {
		if err := t.file.Close(); err != nil {
			return err
		}
		t.file = nil
	}

	if info, err := os.Stat(t.filename); err == nil {
		if info.Size() == 0 {
			return ErrWontErase
		}
		// file exists, rotate it
		if err := os.Rename(t.filename, t.filename+".1"); err != nil {
			return err
		}
	}

	_ = os.MkdirAll(filepath.Dir(t.filename), 0750)

	var err error

	t.file, err = os.Create(t.filename)

	return err
}
