package service

import (
	"SAI/pkg/conf"
	"os"
)

const (
	serverGroup = "server"

	nameKey = "name"
	portKey = "port"
)

var (
	name, _ = conf.GetString(serverGroup, nameKey, "")
	port, _ = conf.GetInt32(serverGroup, portKey, 0)

	Name = name
	Port = port

	AppMode = func() string {
		env := os.Getenv("env")
		if env == "" {
			env = "dev"
		}
		return env
	}()
)
