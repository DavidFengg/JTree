package conf

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type conf struct {
	collections []string `yaml:"collections"`
}

func (c *conf) GetConf() *conf {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
