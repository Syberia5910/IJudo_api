package config

import (
    "io/ioutil"
    
    "gopkg.in/yaml.v2"
  )

type Cfg struct {
    BaseCategory []struct {
        Name string    `yaml:"name"`
        Genre string `yaml:"genre"`
        Variance int `yaml:"variance"`
        Gap int `yaml:"gap"`
    } `yaml:"base_category"`
}

func LoadConfig(path string, cfg *Cfg) error {
    // Read the YAML file
    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        return err
    }

    // Unmarshal the YAML into the config struct
    err = yaml.Unmarshal(yamlFile, cfg)
    if err != nil {
        return err
    }

    return nil
}