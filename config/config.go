package config

import (
  "os"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "bytes"
)

type Config struct {
  Values []struct {
    ProcessPath string `json:"process"`
    X int `json:"x"`
    Y int `json:"y"`
    Width int `json:"width"`
    Height int `json:"height"`
  } `json:"values"`
}

func New() *Config {
  return &Config{}
}

func (aConfig *Config) ImportFromJsonPath(aPath string) (*Config, error) {
  
  tFileStat, tError := os.Stat(aPath)
  if tFileStat.IsDir() || tError != nil {
    return nil, fmt.Errorf("Invalid JSON filepath: %s", aPath)
  }
  
  tRawBytes, tError := ioutil.ReadFile(aPath)
  if tError != nil {
    return nil, fmt.Errorf("Cannot read JSON file: %s", aPath)
  }
  
  var tConfig Config
  tError = json.NewDecoder(bytes.NewReader(tRawBytes)).Decode(&tConfig)
  if tError != nil {
    return nil, fmt.Errorf("Error on parse JSON: \n error = %v \n", tError)
  }
  
  return &tConfig, nil
}