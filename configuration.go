package main

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

type DatabaseConfiguration struct {
    Username string `yaml:username`
    Password string `yaml:password`
    DatabaseName string `database`
}

type HttpConfiguration struct {
    Port string `yaml:port`
}

type Configuration struct {
    Database DatabaseConfiguration `yaml:database`
    Http HttpConfiguration `yaml:http`
}

func LoadConfiguration(pathToConfigFile string) (Configuration) {
    data, err := ioutil.ReadFile(pathToConfigFile)

    if (err != nil) {
        panic(err)
    }

    var configuration = Configuration{}

    err = yaml.Unmarshal(data, &configuration)

    if (err != nil) {
        panic(err)
    }

    return configuration
}