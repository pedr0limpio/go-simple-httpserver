package main

import (
	"github.com/pedr0limpio/go-simple-httpserver/httputils"
)

// TestePrint : funcao print
func TestePrint() string {
	return "Por enquanto..."
}

func main() {
	httputils.ServeHTTPS()
	httputils.ServeHTTP()
}
