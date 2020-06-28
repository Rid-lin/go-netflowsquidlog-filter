package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var year, inFile, outFile, configFile string

// Config ..
type Config struct {
	ProcessingDirection string   `json:"processing_direction"`
	UserFinder          string   `json:"user_finder"`
	SubNet              []string `json:"subnets"`
	IgnorList           []string `json:"ignor"`
}

var config *Config

// var cfg *Config
var err error

func init() {
	flag.StringVar(&configFile, "c", "./go-netflowsquidlog-filter.json", "configuration file")
	flag.StringVar(&inFile, "in", "./1.log", "Temp log file for filtering")
	flag.StringVar(&outFile, "out", "./access.log", "Log file for further processing by the analyzer")
	flag.Parse()
}

func main() {
	err := config.loadConfigFromFile(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}

	file, err := os.Open(inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	err2 := fullFileHandling(scanner)
	if err2 != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func (config *Config) loadConfigFromFile(configFile string) error {
	plan, _ := ioutil.ReadFile(configFile)
	// var data interface{}
	err := json.Unmarshal(plan, &config)
	if err != nil {
		return err
	}
	return nil
}

func fullFileHandling(scanner *bufio.Scanner) error {
	:TO_SCAN
	for scanner.Scan() {
		line := scanner.Text()
		for _, ignorItem := range IgnorList{
		if strings.Contains(line, ignorItem)
	}
		valueArray := strings.Fields(line)

		// TODO тут должна находится функция обработки строк лог-фала
		fmt.Println(valueArray)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
