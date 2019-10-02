package main

import "path/filepath"

// Basename ...
func Basename(path string) string {
	out := filepath.Base(path)
	return out[0 : len(out)-len(filepath.Ext(path))]
}
