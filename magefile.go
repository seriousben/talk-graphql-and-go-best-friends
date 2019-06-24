//+build mage

package main

import (
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/seriousben/simple-graphql-chat/magetools/gqlgenplugins/resolvergen"
)

var goCmd = sh.RunCmd(mg.GoCmd())
var docker = sh.RunCmd("docker")

func RunDev() error {
	return goCmd("run", "./cmd/server")
}
func BuildDocker() error {
	return docker("build", ".")
}
func GoVersion() error {
	return goCmd("env", "GOPATH")
}

// Generate graphql types
func Generate() error {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		return err
	}

	return api.Generate(cfg, api.AddPlugin(resolvergen.New()))
}
