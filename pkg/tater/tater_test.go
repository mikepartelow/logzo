package tater_test

import (
	"io"
	"os"
	"path"
	"testing"

	"github.com/mikepartelow/logzo/pkg/tater"
	"github.com/stretchr/testify/assert"
)

var _ io.Writer = &tater.Tater{}

func TestCreatesDirectories(t *testing.T) {
	filename := path.Join(t.TempDir(), "foo", "bar", "baz", "blof.txt")

	ttr := tater.New(filename)
	assert.NotNil(t, ttr)

	n, err := ttr.Write([]byte("spam"))
	assert.NoError(t, err)
	assert.Equal(t, 4, n)

	assert.FileExists(t, filename)
}

func TestRotate(t *testing.T) {
	filename := path.Join(t.TempDir(), "foo", "bar", "baz", "blof.txt")

	ttr := tater.New(filename)
	_, _ = ttr.Write([]byte("spam"))

	contents, err := os.ReadFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, "spam", string(contents))

	assert.NoError(t, ttr.Rotate())

	contents, err = os.ReadFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, "", string(contents))

	contents1, err := os.ReadFile(filename + ".1")
	assert.NoError(t, err)
	assert.Equal(t, "spam", string(contents1))
}

func TestItDoesntOverwriteWithZeroBytes(t *testing.T) {
	filename := path.Join(t.TempDir(), "foo", "bar", "baz", "blof.txt")

	ttr := tater.New(filename)
	_, _ = ttr.Write([]byte("spam"))
	_ = ttr.Rotate()

	_, _ = ttr.Write([]byte(""))

	assert.ErrorIs(t, ttr.Rotate(), tater.ErrWontErase)

	contents1, _ := os.ReadFile(filename + ".1")
	assert.Equal(t, "spam", string(contents1))
}
