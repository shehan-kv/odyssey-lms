package web

import (
	"embed"
)

//go:embed all:ui/build
var WebUiFS embed.FS
