package model

import (
	"fmt"
)

type AppConfig struct {
	Hosts []Host
	Tftp  struct {
		Root string
	}
	Configuration struct {
		Directory string
	}
	Server struct {
		Port int
	}
}

func (c *AppConfig) String() string {
	return fmt.Sprintf("%+v", *c)
}