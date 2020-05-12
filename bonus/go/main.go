package main

import (
	"Poll/poll"
	"os"
	"sync"
)

func main() {
	args := os.Args
	err := poll.CheckArgv(args[1:])
	if err != nil {
		poll.PrintHelp(args[0], err)
	}
	p, err := poll.InitPoll(args[1:])
	if err != nil {
		poll.PrintHelp(args[0], err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go p.CalculateVariance(&wg)
	wg.Wait()
	p.Print()
}
