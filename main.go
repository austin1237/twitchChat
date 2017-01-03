package main

import (
  "fmt"
  "crypto/tls"
  "flag"
  "os"
  "github.com/austin1237/twitchChat/chatMonitor"
  irc "github.com/fluffle/goirc/client"
)
// Variables used for command line parameters
var (
  Username string
	Channel string
	Token    string
  Debug bool
)

func init() {

	flag.StringVar(&Channel, "c", "", "twitch channel to moniter")
	flag.StringVar(&Token, "t", "", "twitch oauth token")
  flag.StringVar(&Username, "u", "", "your twitch username")
  flag.BoolVar(&Debug, "debug", false, "boolean show debug logs")
	flag.Parse()


	if Token == ""{
		fmt.Println("twitch chat oauth token was not provided.")
		os.Exit(1)
	}

  if Channel == ""{
    fmt.Println("Twitch channel was not provided")
    os.Exit(1)
  }

  if Username == ""{
    fmt.Println("Twitch username was not provided")
    os.Exit(1)
  }

}


func joinChannel(conn *irc.Conn, line *irc.Line){
  fmt.Println("Connected")
  conn.Join("#" + Channel)
  fmt.Println("Connected to channel")
  go chatMonitor.StartMonitoring()
}

func chatMessageRecevied(conn *irc.Conn, line *irc.Line) {
  chatMonitor.AddToCount()
}

func main(){
  // Or, create a config and fiddle with it first:
  cfg := irc.NewConfig(Username)
  cfg.SSL = true
  cfg.SSLConfig = &tls.Config{ServerName: "irc.chat.twitch.tv"}
  cfg.Server = "irc.chat.twitch.tv:443"
  cfg.Pass = "oauth:" + Token
  c := irc.Client(cfg)
  quit := make(chan bool)
  c.HandleFunc(irc.PRIVMSG, chatMessageRecevied)
  c.HandleFunc(irc.CONNECTED, joinChannel)
// Tell client to connect.
  if err := c.Connect(); err != nil {
      fmt.Printf("Connection error: %s\n", err.Error())
  }

  <- quit
}
