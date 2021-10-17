package main

import (
	. "golang_in_memory_key_value/server"
)

//go:generate swagger generate spec
func main() {
	NewServer().Run("localhost:8080")
}
