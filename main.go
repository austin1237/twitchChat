package main

import (
  "fmt"
  "crypto/tls"
  "time"
  "flag"
  "os"
  irc "github.com/fluffle/goirc/client"
)
// Variables used for command line parameters
var (
  Username string
	Channel string
	Token    string
  Duration int
)

var currentChatCount int
var lastChatCount int

func init() {

	flag.StringVar(&Channel, "c", "", "twitch channel to moniter")
	flag.StringVar(&Token, "t", "", "twitch oauth token")
  flag.StringVar(&Username, "u", "", "your twitch username")
  flag.IntVar(&Duration, "d", 5, "# of seconds the chat count will be calculated")
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

func checkChatCount(){
  time.Sleep(5 * time.Second)
  if (lastChatCount != 0 && currentChatCount != 0){
    increase := currentChatCount - lastChatCount
    changePercentage := (float64(increase) / float64(lastChatCount)) * 100
    fmt.Println("Percentage change was")
    fmt.Printf("%.6f\n", changePercentage)
  }
  fmt.Printf("current %d last %d\n", currentChatCount, lastChatCount)
  lastChatCount = currentChatCount
  currentChatCount = 0
  go checkChatCount()
}

func printMessage(conn *irc.Conn, line *irc.Line) {
  // Will print e.g.: "#go-nuts: <fluffle> I am saying something"
  currentChatCount++
  fmt.Printf("Message")
  fmt.Printf("%s: <%s> %s\n", line.Target(), line.Nick, line.Text())
}

func printAction(conn *irc.Conn, line *irc.Line) {
// Will print e.g.: "#go-nuts: * fluffle does something"
fmt.Printf("Action")
fmt.Printf("%s: * %s %s\n", line.Target(), line.Nick, line.Text())
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
  c.HandleFunc(irc.PRIVMSG, printMessage)
  c.HandleFunc(irc.ACTION, printAction)
  c.HandleFunc(irc.CONNECTED,
        func(conn *irc.Conn, line *irc.Line) {
            fmt.Println("Connected")
            conn.Join("#" + Channel)
            fmt.Println("Connected to channel")
            go checkChatCount()
        })
// Tell client to connect.
  if err := c.Connect(); err != nil {
      fmt.Printf("Connection error: %s\n", err.Error())
  }

  <- quit
}
