package tater_test

import (
	"io"
	"testing"

	"github.com/mikepartelow/logzo/pkg/tater"
)

var _ io.Writer = &tater.Tater{}

func Test(t *testing.T) {
	t.Errorf("FIXME")
}
