package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"blekksprut.net/esper"
)

func main() {
	c := flag.String("c", "/etc/esper.conf", "config file")
	s := flag.String("s", "", "symbols (animals/cards)")

	flag.Parse()

	f, err := os.Open(*c)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var symbols []string
	switch *s {
	case "animals", "":
		symbols = esper.Animals
	case "cards":
		symbols = esper.Cards
	default:
		log.Fatal("unknown symbols")
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "=>") {
			line = strings.TrimSpace(strings.TrimPrefix(line, "=>"))

			url, name, found := strings.Cut(line, " ")
			if !found {
				name = url
			}
			err := esper.Try(url, name)
			var first, second string
			if err != nil {
				i := rand.Intn(len(symbols))
				n := rand.Intn(len(symbols)-1) + 1
				j := (i + n) % len(symbols)
				first, second = symbols[i], symbols[j]
			} else {
				first = symbols[rand.Intn(len(symbols))]
				second = first
			}

			fmt.Printf("=> %s %s%s %s\n", url, first, second, name)
		} else {
			fmt.Println(line)
		}
	}

	now := time.Now()
	fmt.Printf("\n%s\n", name, now.Format("2006.01.02 15:04"))
}
