package config

import (
	"fmt"
	"io/ioutil"
	entity "lelo-user/entity"

	"gopkg.in/yaml.v2"
)

var ConfigData entity.Config
var CredentialData entity.Credential

func ReadConfiguration(path string) error {

	configFile, err := ioutil.ReadFile(
		fmt.Sprintf("%s/config.yaml", path),
	)
	if err != nil {
		return err
	}
	credentialFile, err := ioutil.ReadFile(
		fmt.Sprintf("%s/credential.yaml", path),
	)
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
