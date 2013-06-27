package main

import (
  "bufio"
  "fmt"
  "path/filepath"
  "io/ioutil"
  "os"
  "os/user"
  "log"
  "strings"
)

var ConfigFile = getHomePath(".webbynode")

func getHomePath(file string) string {
  usr, err := user.Current()
  if err != nil {
    log.Fatal(err)
  }
  return filepath.Join(usr.HomeDir, file)
}

func readConfig(file string) []string {
  b, err := ioutil.ReadFile(file)
  if err != nil {
    panic(err)
  }
  return strings.Split(string(b), "\n")
}

func GetCredentials(inCfg *WebbynodeCfg, overwrite bool) *WebbynodeCfg {
  var config *WebbynodeCfg
  if inCfg != nil { config = inCfg }
  config.configFile = getHomePath(".webbynode")

  if config.Exists() {
    config.Load()
  }

  if !config.Exists() || overwrite {
    if overwrite || config.system == "" {
      system, err := getInput("What's the end point you're using - manager or manager2? ")
      if err != nil {
        panic(err)
      }
      config.system = system
    }
    if overwrite || config.email == "" {
      email, err := getInput("Login email: ")
      if err != nil {
        panic(err)
      }
      config.email = email
    }
    if overwrite || config.token == "" {
      token, err := getInput("API token:   ")
      if err != nil {
        panic(err)
      }
      config.token = token
    }
    config.Save()
  }

  return config
}

func getInput(prompt string) (string, error) {
  fmt.Print(prompt)
  reader := bufio.NewReader(os.Stdin)
  value, err := reader.ReadString('\n')
  return strings.TrimSpace(value), err
}

func (cfg *WebbynodeCfg) Load() (bool, error) {
  if !cfg.Exists() {
    return false, nil
  }

  for _, line := range readConfig(cfg.configFile) {
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
  contents := []byte(strings.Join(lines, "\n"))
  return ioutil.WriteFile(cfg.configFile, contents, 0644)
}

func (cfg *WebbynodeCfg) Exists() bool {
  if _, err := os.Stat(cfg.configFile); err == nil {
    return true
  }
  return false
}

type WebbynodeCfg struct {
  configFile  string
  email       string
  token       string
  system      string
  awsKey      string
  awsSecret   string
}

