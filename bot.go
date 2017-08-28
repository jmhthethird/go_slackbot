package main

import (
    "fmt"
    "os"
    "strings"
    // "github.com/alexflint/go-arg"
    "github.com/nlopes/slack"
)

// xoxb-228144520231-R1GvxePgK9SwUT3YsmIgXLTH

var acceptedCommands = []string{
    "hey",
    "help",
    "stahp",
}

type module interface {
    
}

type response struct {
    rtm      *slack.RTM
    msg      *slack.MessageEvent
    response string
}

func respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string) {
    text := msg.Text
    text = strings.TrimPrefix(text, prefix)
    text = strings.TrimSpace(text)
    text = strings.ToLower(text)
    textslice := strings.Split(text, " ")
    fmt.Println(textslice)
    command, rc := isStringInSlice(textslice[0], acceptedCommands)
    if rc {
        fmt.Println("Command accepted")
        message := whatToSay(command)
        rtm.SendMessage(rtm.NewOutgoingMessage(message, msg.Channel))
    }else {
        fmt.Println(command,"is not a valid command.")
        usage()
        return
    }
}

func whatToSay(command string) string {
    if command == "hey" {
        return "suuuuuup"
    }else if command == "help" {
        usage()
        return ""
    }else if command == "stahp" {
        return `
┓┏┓┏┓┃
┛┗┛┗┛┃＼○／
┓┏┓┏┓┃   /     STAHP
┛┗┛┗┛┃ノ)
┓┏┓┏┓┃         PUSHING
┛┗┛┗┛┃
┓┏┓┏┓┃         UNTESTED
┛┗┛┗┛┃
┓┏┓┏┓┃         CODE
┃┃┃┃┃┃
`
    }
    return ""
    
}

func isStringInSlice(aString string, aSlice []string) (string, bool) {
    for _, value := range aSlice {
        if value == aString {
            return value, true
        }
    }
    return aString, false
}

func usage() {
    fmt.Println("Usage info")
}




func main() {
    // commandSent := "hey"
    // fmt.Println("Command sent:",commandSent)
    // processCommand(commandSent)
    // fmt.Println("Acceptable Commands:",acceptedCommands)

    token := os.Getenv("SLACK_TOKEN")
    api := slack.New(token)
    rtm := api.NewRTM()
    go rtm.ManageConnection()

Loop:
    for {
        select {
        case msg := <-rtm.IncomingEvents:
            fmt.Print("\nEvent Received: ")
            switch ev := msg.Data.(type) {
            case *slack.ConnectedEvent:
                fmt.Println("Connection counter:", ev.ConnectionCount)

            case *slack.MessageEvent:
                fmt.Printf("\n\tMessage: %v\n", ev)
                info := rtm.GetInfo()
                prefix := fmt.Sprintf("<@%s> ", info.User.ID)

                if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
                    respond(rtm, ev, prefix)
                }

            case *slack.RTMError:
                fmt.Printf("Error %s\n", ev.Error())

            case *slack.InvalidAuthEvent:
                fmt.Printf("Invalid credentials")
                break Loop

            default:
                // Do nothing.
            }
        }
    }
}