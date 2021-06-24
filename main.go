package main

import (
	"fmt"
	"github.com/pborman/getopt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {

	configFilePath := getopt.StringLong("config-file", 'c', "/etc/ccb/circuitbreaker.yaml", "Configuration for circuitbreaker")
	getopt.Parse()
	println(*configFilePath)
	var config CcbConfig
	configFile, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		log.Errorf("%s  에 설정파일이 존재하지 않습니다", *configFilePath)
	}
	err = yaml.Unmarshal(configFile, &config)
	fmt.Println(config)
	//_ = NewStorage("./", "ccb")
	InitDB("./ccb")


}

