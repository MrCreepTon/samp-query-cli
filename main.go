package main

import (
	"context"
	"fmt"
	sampquery "github.com/Southclaws/go-samp-query"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: %s [IP:Port] <-r to see rules>\n", args[0])
		return
	}
	showRules := false
	host := ""
	for _, arg := range args[1:] {
		switch arg {
		case "-r":
			showRules = true
		default:
			host = arg
		}
	}
	if !strings.Contains(host, ":") {
		host += ":7777"
	}
	ctx := context.Background()
	query, err := sampquery.NewQuery(host)
	if err != nil {
		fmt.Printf("Error! %s\n", err)
		return
	}
	info, err := query.GetInfo(ctx, true)
	if err != nil {
		fmt.Printf("Error! %s\n", err)
		return
	}
	ping, err := query.GetPing(ctx)
	if err != nil {
		fmt.Printf("Error! %s\n", err)
		return
	}
	if showRules {
		fmt.Println(">> Basic Info <<")
	}
	fmt.Printf("Hostname: %s\n", info.Hostname)
	fmt.Printf("Address: %s\n", host)
	fmt.Printf("Ping: %d ms\n", ping.Milliseconds())
	fmt.Printf("Players: %d / %d\n", info.Players, info.MaxPlayers)
	fmt.Printf("Gamemode: %s\n", info.Gamemode)
	if showRules {
		rules, err := query.GetRules(ctx)
		if err != nil {
			fmt.Printf("Error! %s\n", err)
			return
		}
		fmt.Println(">> Rules <<")
		for ruleName, ruleValue := range rules {
			fmt.Printf("%s: %s\n", ruleName, ruleValue)
		}
	}
}
