package utils

import (
	"log"
)

// InitLogger initializes the logger
func InitLogger() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}