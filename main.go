package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const defaultStandTime = 10
const defaultSeatTime = 50

const standBanner = `
 $$$$$$\ $$$$$$$$\  $$$$$$\  $$\   $$\ $$$$$$$\        $$\   $$\ $$$$$$$\  
$$  __$$\\__$$  __|$$  __$$\ $$$\  $$ |$$  __$$\       $$ |  $$ |$$  __$$\ 
$$ /  \__|  $$ |   $$ /  $$ |$$$$\ $$ |$$ |  $$ |      $$ |  $$ |$$ |  $$ |
\$$$$$$\    $$ |   $$$$$$$$ |$$ $$\$$ |$$ |  $$ |      $$ |  $$ |$$$$$$$  |
 \____$$\   $$ |   $$  __$$ |$$ \$$$$ |$$ |  $$ |      $$ |  $$ |$$  ____/ 
$$\   $$ |  $$ |   $$ |  $$ |$$ |\$$$ |$$ |  $$ |      $$ |  $$ |$$ |      
\$$$$$$  |  $$ |   $$ |  $$ |$$ | \$$ |$$$$$$$  |      \$$$$$$  |$$ |      
 \______/   \__|   \__|  \__|\__|  \__|\_______/        \______/ \__|      
 `

const seatBanner = `
 $$$$$$\  $$$$$$$$\  $$$$$$\ $$$$$$$$\       $$$$$$$\   $$$$$$\  $$\      $$\ $$\   $$\ 
$$  __$$\ $$  _____|$$  __$$\\__$$  __|      $$  __$$\ $$  __$$\ $$ | $\  $$ |$$$\  $$ |
$$ /  \__|$$ |      $$ /  $$ |  $$ |         $$ |  $$ |$$ /  $$ |$$ |$$$\ $$ |$$$$\ $$ |
\$$$$$$\  $$$$$\    $$$$$$$$ |  $$ |         $$ |  $$ |$$ |  $$ |$$ $$ $$\$$ |$$ $$\$$ |
 \____$$\ $$  __|   $$  __$$ |  $$ |         $$ |  $$ |$$ |  $$ |$$$$  _$$$$ |$$ \$$$$ |
$$\   $$ |$$ |      $$ |  $$ |  $$ |         $$ |  $$ |$$ |  $$ |$$$  / \$$$ |$$ |\$$$ |
\$$$$$$  |$$$$$$$$\ $$ |  $$ |  $$ |         $$$$$$$  | $$$$$$  |$$  /   \$$ |$$ | \$$ |
 \______/ \________|\__|  \__|  \__|         \_______/  \______/ \__/     \__|\__|  \__|
 `

var standSound = exec.Command("say", "Stand Up")
var seatSound = exec.Command("say", "Seat Down")

var standNotification = exec.Command("osascript", "-e", "display notification \"Stand Up\" with title \"Standing Desk Notifier\"")
var seatNotification = exec.Command("osascript", "-e", "display notification \"Seat Down\" with title \"Standing Desk Notifier\"")

func detailsBanner(duration time.Duration) string {
	// Mon Jan 2 15:04:05 MST 2006
	timeFormat := "15:04"
	now := time.Now()

	return fmt.Sprintf("\n===============\n"+
		"Start: %s\n"+
		"End:   %s\n"+
		"Duration: %s\n"+
		"===============", now.Format(timeFormat), now.Add(duration).Format(timeFormat), duration)
}

func startupBanner(standTime time.Duration, seatTime time.Duration) string {
	return fmt.Sprintf("\n=================\nStand Time: %s\nSeat Time:  %s\n=================", standTime, seatTime)
}

func progressBar() {
	fmt.Println("")

	for i := 0; i < 8; i++ {
		fmt.Println("* * * * * * * *")
		time.Sleep(time.Second)
	}
}

func main() {
	var standArg int
	var seatArg int

	switch len(os.Args) {
	case 1:
		standArg = defaultStandTime
		seatArg = defaultSeatTime
	case 2:
		standArg, _ = strconv.Atoi(os.Args[1])
	case 3:
		standArg, _ = strconv.Atoi(os.Args[1])
		seatArg, _ = strconv.Atoi(os.Args[2])
	}

	standTime := time.Minute * time.Duration(standArg)
	seatTime := time.Minute * time.Duration(seatArg)

	fmt.Println(startupBanner(standTime, seatTime))

	for {
		fmt.Println(standBanner, detailsBanner(standTime))
		standSound.Run()
		standNotification.Run()
		progressBar()
		time.Sleep(standTime)

		fmt.Println(seatBanner, detailsBanner(seatTime))
		seatSound.Run()
		seatNotification.Run()
		progressBar()
		time.Sleep(seatTime)
	}
}
