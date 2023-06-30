package tater

import "os"

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

	if _, err := os.Stat(t.filename); err == nil {
		// file exists, rotate it
		if err := os.Rename(t.filename, t.filename+".1"); err != nil {
			return err
		}
	}

	var err error

	t.file, err = os.Create(t.filename)

	return err
}
