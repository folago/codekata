package main

import (
	"flag"
	"log"
	"time"

	"github.com/folago/codekata/chains"
)

var words, begin, end string

func init() {
	flag.StringVar(&begin, "begin", "", "the word to begin the chain")
	flag.StringVar(&end, "end", "", "the word to end the chain")
	flag.StringVar(&words, "words", "", "the file containing the words")
}

func main() {
	flag.Parse()
	if begin == "" || end == "" || words == "" {
		log.Fatalln("Please provode all the arguments")
	}
	if len(begin) != len(end) {
		log.Fatalln("begin and end words must have the same length")
	}

	words, err := chains.ReadWords(words, len(begin))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Read word list.")
	graph := chains.BuildGraph(words)
	log.Println("Built graph.")
	t0 := time.Now()
	path := chains.Path(graph, begin, end)
	tt := time.Now().Sub(t0)
	if len(path) == 0 {
		log.Printf("No path found between %s and %s.\nTime elapsed in searchig %v.\n", begin, end, tt)

	}
	log.Printf("Path between %s and %s:\n\t%v.\nTime elapsed in searchig %v.\n", begin, end, path, tt)
}
