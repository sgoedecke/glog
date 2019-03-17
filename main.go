package main

import (
  "fmt"
  "log"
  "github.com/appleboy/easyssh-proxy"
  "github.com/kevinburke/ssh_config"
  "time"
  "os"
  "os/user"
)

func main() {
    u, _ := user.Current()
    user := u.Username
    log.Print("Executing as " + user)

    host := os.Args[1]
    targetServer := ssh_config.Get(host, "HostName")
    bastionNickname := ssh_config.Get(host, "ProxyJump")
    bastionServer := ssh_config.Get(bastionNickname, "HostName")

	ssh := &easyssh.MakeConfig{
		User:    user,
		Server:  targetServer,
		Port:    "22",
		Proxy: easyssh.DefaultConfig{
			User:    user,
			Server:  bastionServer,
			Port:    "22",
		},
	}

  hostname, _, _, err := ssh.Run("hostname")
  if err != nil {
    log.Fatal("Failed to run: " + err.Error())
  }
  log.Print("Executing on " + hostname)


  datePath := time.Now().Format("2006/01/02")

  searchGlob := os.Args[2]
  searchKey := os.Args[3]

  if searchGlob == "" {
    log.Fatal("You need to pass a glob")
  }
  if searchKey == "" {
    log.Fatal("You need to pass a search key")
  }

  cmd := fmt.Sprintf("cat /storage/logs/hosts/%s/%s | grep %s", datePath, searchGlob, searchKey)
  log.Print(cmd)
  stdout, stderr, _, err := ssh.Run(cmd)
  if err != nil {
    log.Print("Stderr: " + stderr + stdout)
    log.Fatal("Failed to run: " + err.Error())
  }

  log.Print(stdout)
}
