package main

import (
	"fmt"
	"os"
	"os/signal"
	"word_stat/mapdata"
	"word_stat/words"
)

func main() {
	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt, os.Kill)
	nd := mapdata.NewDobj()
	para := "This command is useful in order to modify a node's view of the cluster configuration. Specifically it assigns a set of hash slots to the node receiving the command. If the command is successful, the node will map the specified hash slots to itself, and will start broadcasting"
	words.WordFromPara(para, nd)
	//nd.Write("string")
	//go words.WordsFromPara
	//go words.WordsFromFile
	//Check if proceeding done
	s := <-cs
	fmt.Println("\nSignal received:", s)
	fmt.Println(nd.Data)
}
