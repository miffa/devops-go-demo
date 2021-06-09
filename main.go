package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	Header = "京"
)

var (
	Second = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N"}
	Last   = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	stopCh = make(chan struct{}, 1)
)

func RandInt64(min, max int64) int64 {
	if min >= max || max == 0 {
		return max

	}
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}

func GenLicensePlate() string {
	return Header + Second[RandInt64(0, 14)] + " " + Last[RandInt64(0, 9)] + Last[RandInt64(0, 9)] + Last[RandInt64(0, 9)] + Last[RandInt64(0, 9)] + Last[RandInt64(0, 9)]
}

func main() {

	fmt.Println("Welcome to our park:", GenLicensePlate())
	go InitSignal()
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Welcome to our park:", GenLicensePlate())

		case <-stopCh:
			fmt.Println("Our park is out of service")
			return
		}
	}
}

func InitSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2)
	for {
		<-c
		fmt.Println("Our park get a signal")
		close(stopCh)
		return
	}
}
