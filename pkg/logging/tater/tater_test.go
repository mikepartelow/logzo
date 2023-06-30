package tater_test

import (
	"io"

	"github.com/mikepartelow/logzo/pkg/logging/tater"
)

var _ io.Writer = &tater.Tater{}
