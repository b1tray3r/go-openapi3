package embed

import "embed"

//go:embed swagger/*
var StaticFiles embed.FS
