package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	config := Config{}
	file, err := getConfigFilePath()
	if err != nil {
		return config
	}
	body, err := os.ReadFile(file)
	if err != nil {
		return config
	}
	if err := json.Unmarshal(body, &config); err != nil {
		return config
	}
	return config

}

func SetCurrentUser(user string) {
	cfg := Read()
	cfg.CurrentUserName = user
	write(cfg)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := homeDir + "/" + configFileName
	return fullPath, nil
}

func write(cfg Config) error {
	cfgData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err2 := f.Write(cfgData)

	if err2 != nil {
		return err2
	}

	return nil
}
