package main

import (
	"bufio"
	"checker/internal/checker"
	"sync"
	"os"
	"math/rand"
)

const threads = 100

var (
	proxies []string
)

func main() {
	invites, err := os.Open("invites.txt")
	if err != nil {
		panic(err)
	}
	defer invites.Close()

	proxyFile, err := os.Open("proxies.txt")
	if err != nil {
		panic(err)
	}
	defer proxyFile.Close()

	scanner := bufio.NewScanner(proxyFile)
	for scanner.Scan() {
		proxy := scanner.Text()
		proxies = append(proxies, proxy)
	}

	scanner = bufio.NewScanner(invites)
    var wg sync.WaitGroup
    inviteCh := make(chan string, threads)

    for i := 0; i < threads; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for invite := range inviteCh {
                checker.Check(invite, proxies[rand.Intn(len(proxies))])
            }
        }()
    }

    for scanner.Scan() {
        invite := scanner.Text()
        inviteCh <- invite
    }

    close(inviteCh)
    wg.Wait()
}