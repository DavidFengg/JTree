package conf

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

//Conf is an object created by the conf.yaml file
type Conf struct {
	collections []string `yaml:"collections"`
}

//GetConf fills the conf struct
func (c *Conf) GetConf() *Conf {

	yamlFile, err := ioutil.ReadFile("./conf/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
