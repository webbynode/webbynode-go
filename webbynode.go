package main

import (
  "flag"
  "fmt"
  "log"
  "os"
)

func main() {
  if err := ParseCommands(flag.Args()...); err != nil {
    log.Fatal(err)
    os.Exit(-1)
  }
}

func (cli *WebbynodeCli) CmdHelp(args ...string) error {
  help := fmt.Sprintf("Usage: webbynode [OPTIONS] COMMAND [arg...]\n\nCommands:\n")
  // \n  -H=[tcp://%s:%d]: tcp://host:port to bind/connect to or unix://path/to/socker to use\n\nA self-sufficient runtime for linux containers.\n\nCommands:\n", DEFAULTHTTPHOST, DEFAULTHTTPPORT)
  for _, command := range [][2]string{
    {"accounts",        "Manages multiple Webbynode accounts"},
    {"add_backup",      "Configures automatic nightly backups for the current application"},
    {"add_key",         "Adds your ssh public key to your Webby, making deployments easy"},
    {"addons",          "Manages you application's add-ons"},
    {"alias",           "Manages a list of alias for your common used remote commands"},
    {"apps",            "Lists all apps installed in your Webby"},
    {"authorize_root",  "Adds your ssh public key to your Webby's root user"},
    {"change_dns",      "Changes the DNS entry for this application"},
    {"config",          "Adds or changes your Webbynode API credentials"},
    {"console",         "Opens a Rails 3 console session"},
    {"database",        "Manages your application database"},
    {"delete",          "Removes the current application from your Webby"},
    {"dns_aliases",     "Changes the DNS aliases for this application"},
    {"docs",            "Opens Webbynode Documentation in your browser"},
    {"help",            "Guess what? You're on it!"},
    {"init",            "Prepares the application on current folder for deployment"},
    {"logs",            "Tails a your Rails application logs"},
    {"open",            "Opens the current application in your browser"},
    {"push",            "Sends pending changes on the current application to your Webby"},
    {"remote",          "Execute commands on your Webby for the current application"},
    {"restart",         "Reboots your Webby"},
    {"settings",        "Manages application settings"},
    {"ssh",             "Log into your Webby via SSH"},
    {"start",           "Starts your Webby, when it's off"},
    {"stop",            "Shuts down your Webby"},
    {"tasks",           "Manages tasks executed before or after you push your changes"},
    {"user",            "Manages Webbynode Trial user"},
    {"version",         "Displays current version of Webbynode Gem"},
    {"webbies",         "Lists the Webbies you currently own"},
  } {
    help += fmt.Sprintf("    %-15.15s%s\n", command[0], command[1])
  }
  fmt.Println(help)
  return nil
}

