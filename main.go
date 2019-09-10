package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

func detailsBanner(duration time.Duration) string {
	// Mon Jan 2 15:04:05 MST 2006
	timeFormat := "15:04"
	now := time.Now()

	return fmt.Sprintf("\n============\n"+
		"Start: %s\n"+
		"End:   %s\n"+
		"Duration: %02d\n"+
		"============\n", now.Format(timeFormat), now.Add(duration).Format(timeFormat), int(duration.Minutes()))
}

func startupBanner(standTime time.Duration, seatTime time.Duration) string {
	return fmt.Sprintf("\n==============\n"+
		"Stand Time: %02d\n"+
		"Seat Time:  %02d\n"+
		"==============", int(standTime.Minutes()), int(seatTime.Minutes()))
}

func sleep(duration time.Duration) {
	cycles := int(duration.Minutes())

	for i := 0; i < cycles; i++ {
		fmt.Printf("%s\n", strings.Repeat("* ", cycles-i))
		time.Sleep(time.Minute)
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
		standSound := exec.Command("say", "Stand Up")
		seatSound := exec.Command("say", "Seat Down")
		standNotification := exec.Command("osascript", "-e", "display notification \"Stand Up\" with title \"Standing Desk Notifier\"")
		seatNotification := exec.Command("osascript", "-e", "display notification \"Seat Down\" with title \"Standing Desk Notifier\"")

		fmt.Println(standBanner, detailsBanner(standTime))
		standSound.Run()
		standNotification.Run()
		sleep(standTime)

		fmt.Println(seatBanner, detailsBanner(seatTime))
		seatSound.Run()
		seatNotification.Run()
		sleep(seatTime)
	}
}
