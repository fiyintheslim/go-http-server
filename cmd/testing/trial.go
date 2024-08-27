package main

import (
	"fmt"

	"github.com/fiyintheslim/go-http-server/internal/utils"
)

func main() {
	res := utils.HashString("hello there")
	fmt.Println(res)
}