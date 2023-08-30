package main

import "github.com/amtoaer/daed-sub/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
