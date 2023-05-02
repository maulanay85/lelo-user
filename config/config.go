package config

import (
	"io/ioutil"
	entity "lelo-user/entity"

	"gopkg.in/yaml.v2"
)

var ConfigData entity.Config
var CredentialData entity.Credential

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

