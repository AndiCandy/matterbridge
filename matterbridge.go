package main

import (
	"crypto/tls"
	"github.com/42wim/matterbridge/matterhook"
	"github.com/thoj/go-ircevent"
	"log"
	"strconv"
	"strings"
	"time"
)

type Bridge struct {
	i *irc.Connection
	m *matterhook.Client
	*Config
}

func NewBridge(name string, config *Config) *Bridge {
	b := &Bridge{}
	b.Config = config
	b.m = matterhook.New(b.Config.Mattermost.URL,
		matterhook.Config{Port: b.Config.Mattermost.Port, Token: b.Config.Mattermost.Token})
	b.i = b.createIRC(name)
	go b.handleMatter()
	return b
}

func (b *Bridge) createIRC(name string) *irc.Connection {
	i := irc.IRC(b.Config.IRC.Nick, b.Config.IRC.Nick)
	i.UseTLS = b.Config.IRC.UseTLS
	i.TLSConfig = &tls.Config{InsecureSkipVerify: b.Config.IRC.SkipTLSVerify}
	i.Connect(b.Config.IRC.Server + ":" + strconv.Itoa(b.Config.IRC.Port))
	time.Sleep(time.Second)
	log.Println("Joining", b.Config.IRC.Channel, "as", b.Config.IRC.Nick)
	i.Join(b.Config.IRC.Channel)
	i.AddCallback("PRIVMSG", b.handlePrivMsg)
	i.AddCallback("CTCP_ACTION", b.handlePrivMsg)
	if b.Config.Mattermost.ShowJoinPart {
		i.AddCallback("JOIN", b.handleJoinPart)
		i.AddCallback("PART", b.handleJoinPart)
	}
	i.AddCallback("353", b.handleOther)
	return i
}

func (b *Bridge) handlePrivMsg(event *irc.Event) {
	msg := ""
	if event.Code == "CTCP_ACTION" {
		msg = event.Nick + " "
	}
	msg += event.Message()
	b.Send("irc-"+event.Nick, msg)
}

func (b *Bridge) handleJoinPart(event *irc.Event) {
	b.Send(b.Config.IRC.Nick, "irc-"+event.Nick+" "+strings.ToLower(event.Code)+"s "+event.Message())
}

func (b *Bridge) handleOther(event *irc.Event) {
	switch event.Code {
	case "353":
		b.Send(b.Config.IRC.Nick, event.Message()+" currently on IRC")
	}
}

func (b *Bridge) Send(nick string, message string) error {
	matterMessage := matterhook.OMessage{IconURL: b.Config.Mattermost.IconURL}
	matterMessage.UserName = nick
	matterMessage.Text = message
	b.m.Send(matterMessage)
	return nil
}

func (b *Bridge) handleMatter() {
	for {
		message := b.m.Receive()
		switch message.Text {
		case "!users":
			log.Println("received !users from", message.UserName)
			b.i.SendRaw("NAMES " + b.Config.IRC.Channel)
		}
		b.i.Privmsg(b.Config.IRC.Channel, message.UserName+": "+message.Text)
	}
}

func main() {
	NewBridge("matterbot", NewConfig("matterbridge.conf"))
	select {}
}