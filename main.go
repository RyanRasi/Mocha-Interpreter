package main

import (
	"./repl"
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	currentTime := time.Now()
	currentTimeSplit := strings.Split(currentTime.String(), " ")
	dateSplit := strings.Split(currentTimeSplit[0], "-")

	var daySuffixes = map[string]string{
		"01": "st",
		"02": "nd",
		"03": "rd", //
		"04": "th",
		"05": "th",
		"06": "th",
		"07": "th",
		"08": "th",
		"09": "th",
		"10": "th",
		"11": "th",
		"12": "th",
		"13": "th",
		"14": "th",
		"15": "th",
		"16": "th",
		"17": "th",
		"18": "th",
		"19": "th",
		"20": "th",
		"21": "st", //
		"22": "nd",
		"23": "rd", //
		"24": "th",
		"25": "th",
		"26": "th",
		"27": "th",
		"28": "th",
		"29": "th",
		"30": "th",
		"31": "st", //
	}
	for k, v := range daySuffixes {
		if dateSplit[2] == k {
			dateSplit[2] = dateSplit[2] + v
		}
	}

	fmt.Printf("Hello %s! This is the .Mocha programming language!\n",
		strings.ToUpper(user.Username))

	fmt.Printf("The current date is:   %s! The %s of %s %s\nThe time is currently: %s\n\n",
		currentTime.Format("Monday"), dateSplit[2], currentTime.Format("January"), dateSplit[0], currentTime.Format("3:04:05 PM"))

	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

//Page 38
