package main

import (
	"docker-golang/cmd"
)

func main() {
	server := cmd.Server{}
	server.Start()
}
