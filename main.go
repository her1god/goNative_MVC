package main

import (
	"go_native/config"
	"go_native/errorhandles/errorhandle"
)

func main() {
	config.ConnectDB()

	errorhandle.ErrorHandle()
}
