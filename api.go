package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "os/user"
  "path/filepath"
  "strings"
)

var ConfigFile = GetHomePath(".webbynode")

func UserHome() string {
  usr, err := user.Current()
  if err != nil {
    log.Fatal(err)
  }
  return usr.HomeDir
}

func GetHomePath(file string) string {
  return filepath.Join(UserHome(), file)
}

func ReadConfig(file string) []string {
  b, err := ioutil.ReadFile(file)
  if err != nil {
    panic(err)
  }
  return strings.Split(string(b), "\n")
}

func GetCredentials(inCfg *WebbynodeCfg, overwrite bool) *WebbynodeCfg {
  var config *WebbynodeCfg
  if inCfg == nil {
    config = &WebbynodeCfg{}
  } else {
    config = inCfg
  }
  config.configFile = GetHomePath(".webbynode")

  if config.Exists() {
    config.Load()
  }

  if !config.Exists() || overwrite {
    if overwrite || config.system == "" {
      system, err := Ask("What's the end point you're using - manager or manager2? ")
      if err != nil {
        panic(err)
      }
      config.system = system
    }
    if overwrite || config.email == "" {
      email, err := Ask("Login email: ")
      if err != nil {
        panic(err)
      }
      config.email = email
    }
    if overwrite || config.token == "" {
      token, err := Ask("API token:   ")
      if err != nil {
        panic(err)
      }
      config.token = token
    }
    config.Save()
  }

  return config
}

func (cfg *WebbynodeCfg) Load() (bool, error) {
  if !cfg.Exists() {
    return false, nil
  }

  for _, line := range ReadConfig(cfg.configFile) {
    if strings.TrimSpace(line) != "" {
      parts := strings.SplitN(line, "=", 2)
      switch parts[0] {
      case "email":
        cfg.email = parts[1]

      case "token":
        cfg.token = parts[1]

      case "system":
        cfg.system = parts[1]

      case "aws_key":
        cfg.awsKey = parts[1]

      case "aws_secret":
        cfg.awsSecret = parts[1]
      }
    }
  }
  return true, nil
}

func (cfg *WebbynodeCfg) Save() error {
  var lines []string
  if cfg.email != "" {
    lines = append(lines, fmt.Sprintf("email=%s", cfg.email))
  }
  if cfg.token != "" {
    lines = append(lines, fmt.Sprintf("token=%s", cfg.token))
  }
  if cfg.system != "" {
    lines = append(lines, fmt.Sprintf("system=%s", cfg.system))
  }
  if cfg.awsKey != "" {
    lines = append(lines, fmt.Sprintf("aws_key=%s", cfg.awsKey))
  }
  if cfg.awsSecret != "" {
    lines = append(lines, fmt.Sprintf("aws_secret=%s", cfg.awsSecret))
  }
  lines = append(lines, "")
  contents := []byte(strings.Join(lines, "\n"))
  return ioutil.WriteFile(cfg.configFile, contents, 0644)
}

func (cfg *WebbynodeCfg) Exists() bool {
  return FileExists(cfg.configFile)
}

type WebbynodeCfg struct {
  configFile string
  email      string
  token      string
  system     string
  awsKey     string
  awsSecret  string
}
