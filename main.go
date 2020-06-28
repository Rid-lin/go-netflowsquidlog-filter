package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sasbury/mini"
)

var year, inFile, outFile, configFile string

// Config ..
type Config struct {
	ProcessingDirection string   `ini:"processing_direction"`
	UserFinder          string   `ini:"user_finder"`
	SubNet              []string `ini:"subnets"`
	IgnorList           []string `ini:"ignor"`
}

var config *mini.Config
var cfg *Config
var err error

func init() {
	flag.StringVar(&configFile, "c", "go-netflowsquidlog-filter.ini", "configuration file")
	flag.StringVar(&inFile, "in", "", "Temp log file for filtering")
	flag.StringVar(&outFile, "out", "/var/log/squid/access2.log", "Log file for further processing by the analyzer")
	flag.Parse()
}

func main() {
	config, err = mini.LoadConfiguration(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	} else {
		cfg.ProcessingDirection = config.String("processing_direction", "both")
		cfg.UserFinder = config.String("user_finder", "")
		cfg.SubNet = config.Strings("subnets")
		cfg.IgnorList = config.Strings("ignor")
	}

	file, err := os.Open(inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fullFileHandling(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func fullFileHandling(scanner *bufio.Scanner) error {
	for scanner.Scan() {
		line := scanner.Text()
		valueArray := strings.Fields(line)

		// TODO тут должна находится функция обработки строк лог-фала
		fmt.Println(valueArray)
	}

	return nil
}
