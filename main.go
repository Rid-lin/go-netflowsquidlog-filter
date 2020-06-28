package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/subpop/go-ini"
)

var year, inFile, outFile, configFile string

// Config ..
type Config struct {
	ProcessingDirection string   `ini:"processing_direction"`
	UserFinder          string   `ini:"user_finder"`
	SubNet              []string `ini:"subnets"`
	IgnorList           []string `ini:"ignor"`
}

var config *Config

// var cfg *Config
var err error

func init() {
	flag.StringVar(&configFile, "c", "./go-netflowsquidlog-filter.ini", "configuration file")
	flag.StringVar(&inFile, "in", "./1.log", "Temp log file for filtering")
	flag.StringVar(&outFile, "out", "./access.log", "Log file for further processing by the analyzer")
	flag.Parse()
}

func main() {
	var data []byte

	f, err := os.Open(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}

	defer f.Close()
	_, err2 := bufio.NewReader(f).Read(data)
	if err2 != nil {
		fmt.Println(err)
		os.Exit(500)
	}

	if err := ini.Unmarshal(data, &config); err != nil {
		fmt.Println(err)
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
