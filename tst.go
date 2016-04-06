package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
	"word_stat/mapdata"
	"word_stat/words"
)

func main() {
	filename := ""
	if len(os.Args) == 2 {
		filename = os.Args[1]
	} else {
		filename = "test.txt"
	}
	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt, os.Kill)
	nd := mapdata.NewDobj()
	//nd.Write("string")
	//go words.WordsFromPara
	go words.WordsFromFile(filename, nd)
	//Check if proceeding done
	l1 := 0
	l2 := 0
	for {
		time.Sleep(300000000)
		l2 = len(nd.Data)
		fmt.Println("Map's size:", l2, "; Queue's size:", len(nd.Queue))
		if l2 == l1 {
			fmt.Println("It looks done now. It looks!!")
			break
		}
		l1 = l2
	}
	s := <-cs
	fmt.Println("\nSignal received:", s)
	fmt.Println("Final map's size:", len(nd.Data))
	totalWordsCount := 0
	for _, v := range nd.Data {
		totalWordsCount += v
	}
	fmt.Println("Total Words Count:", totalWordsCount)
}
