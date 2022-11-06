package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Test test
}

type test struct {
	test map[string][]string
}

func main() {
	var config Config
	if _, err := toml.DecodeFile("example.toml", &config); err != nil {
		fmt.Printf("err int decode file")
	}

	for k, v := range config.Test.test {
		fmt.Printf("k: %s, v: %+v\n", k, v)
	}
}
