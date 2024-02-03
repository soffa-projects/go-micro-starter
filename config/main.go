package config

import (
	"embed"
)

//go:embed db/* locales/*
var RootFS embed.FS
