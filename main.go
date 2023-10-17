package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"golang.design/x/clipboard"
)

type TimeLog struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("mittagctl start")
		fmt.Println("mittagctl end")
		fmt.Println("mittagctl status")
		fmt.Println("mittagctl status start")
		fmt.Println("mittagctl status end")
		return
	}

	currentTime := time.Now().Local()

	if os.Args[1] == "start" {
		timeLog := TimeLog{Start: currentTime}
		timeLogJSON, err := json.Marshal(timeLog)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.WriteFile("/tmp/mittagctl.json", timeLogJSON, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if os.Args[1] == "end" {
		timeLogJSON, err := os.ReadFile("/tmp/mittagctl.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		var timeLog TimeLog
		err = json.Unmarshal(timeLogJSON, &timeLog)
		if err != nil {
			fmt.Println(err)
			return
		}
		timeLog.End = currentTime
		timeLogJSON, err = json.Marshal(timeLog)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.WriteFile("/tmp/mittagctl.json", timeLogJSON, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	timeLogJSON, err := os.ReadFile("/tmp/mittagctl.json")
	if err != nil {
		log.Fatalln(err)
		return
	}
	var timeLog TimeLog
	err = json.Unmarshal(timeLogJSON, &timeLog)
	if err != nil {
		fmt.Println(err)
		return
	}
	if os.Args[1] == "status" && len(os.Args) == 3 {

		if os.Args[2] == "start" {
			fmt.Println(timeLog.Start.Format("Mon 02 Jan 15:04:05"))
			sapTime := timeLog.Start.Format("15:04")
			if err != nil {
				log.Fatalln(err)
			}
			clipboard.Write(clipboard.FmtText, []byte(sapTime))
		} else if os.Args[2] == "end" {
			fmt.Println(timeLog.End.Format("Mon 02 Jan 15:04:05"))
			sapTime := timeLog.End.Format("15:04")
			if err != nil {
				log.Fatalln(err)
			}
			clipboard.Write(clipboard.FmtText, []byte(sapTime))
		}
	} else if os.Args[1] == "status" {
		fmt.Println("Start: ", timeLog.Start.Format("Mon 02 Jan 15:04:05"))
		fmt.Println("End: ", timeLog.End.Format("Mon 02 Jan 15:04:05"))
	} else if os.Args[1] != "start" && os.Args[1] != "end" && os.Args[1] != "status" {
		fmt.Println("Usage:")
		fmt.Println("mittagctl start")
		fmt.Println("mittagctl end")
		fmt.Println("mittagctl status")
		fmt.Println("mittagctl status start")
		fmt.Println("mittagctl status end")
	}
}
