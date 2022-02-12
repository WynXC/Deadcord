package util

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	ColorReset = "\033[0m"
	Red        = "\033[31m"
	Green      = "\033[32m"
	Yellow     = "\033[33m"
	Blue       = "\033[34m"
	Purple     = "\033[35m"
	Cyan       = "\033[36m"
	White      = "\033[37m"
	Black      = "\u001b[30;1m"
)

func GetQuote() string {
	quotes := [12]string{
		"This ain't hacking.",
		"Lagging Discord since last month.",
		"What are you here for?",
		"Great to see you again.",
		"I am sentient.",
		"Tokens not included.",
		"R.I.P Groovy & Rhythm o7.",
		"We built an entire GUI before switching to a terminal window.",
		"Built to protest against Discord.",
		"Deadcord never dies. Its already dead.",
		"Some assembly required.",
	}

	random_quote := quotes[rand.Intn(len(quotes))]

	return random_quote
}

func GetTimestamp() string {
	current_time := time.Now()
	return current_time.Format("15:04:05")
}

func WriteToConsole(status string, mode int) {
	switch mode {
	case 0:
		fmt.Println(White + "[ INFO ] " + "[ " + GetTimestamp() + " ] " + status + ColorReset)
	case 1:
		fmt.Println(Yellow + "[ WARNING ] " + "[ " + GetTimestamp() + " ] " + status + ColorReset)
	case 2:
		fmt.Println(Purple + "[ SUCCESS ] " + "[ " + GetTimestamp() + " ] " + status + ColorReset)
	case 3:
		fmt.Println(Red + "[ ERROR ] " + "[ " + GetTimestamp() + " ] " + status + ColorReset)
	}
}

func Sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func NumberSliceCounts(arr []int) map[int]int {
	dict := make(map[int]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}

	return dict
}

func Contains(elements []string, value string) bool {
	for _, search := range elements {
		if value == search {
			return true
		}
	}
	return false
}

func RemoveFromSlice(slice []string, index int) []string {
	slice[index] = slice[len(slice)-1]
	slice[len(slice)-1] = ""
	slice = slice[:len(slice)-1]

	return slice
}

func AllParameters(parameters []string) bool {
	needed_paramters := len(parameters)
	parameters_filled := 0

	for _, parameter := range parameters {
		if len(parameter) > 0 {
			parameters_filled++
		}
	}

	if needed_paramters == parameters_filled {
		return true
	} else {
		return false
	}
}
