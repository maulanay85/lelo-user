package config

import (
	"io/ioutil"
	config "lelo-user/entity/config"

	"gopkg.in/yaml.v2"
)

var ConfigData config.Config
var CredentialData config.Credential

func ReadConfiguration() error {

	configFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return err
	}
	credentialFile, err := ioutil.ReadFile("./credential.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configFile, &ConfigData)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(credentialFile, &CredentialData)

	if err != nil {
		return err
	}
	return nil
}
