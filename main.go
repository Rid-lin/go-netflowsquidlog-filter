package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// Config ..
type Config struct {
	ProcessingDirection string   `json:"processing_direction"`
	UserFinder          string   `json:"user_finder"`
	SubNets             []string `json:"subnets"`
	IgnorList           []string `json:"ignor"`
}

var (
	inFile, configFile string
	config             *Config
)

func init() {
	flag.StringVar(&configFile, "c", "go-netflowsquidlog-filter.json", "configuration file")
	flag.StringVar(&inFile, "in", "", "Temp log file for filtering")
	// flag.StringVar(&outFile, "out", "", "Log file for further processing by the analyzer")
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
	err2 := config.fullFileHandling(scanner)
	if err2 != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func (cfg *Config) loadConfigFromFile(configFile string) error {
	file, _ := os.Open(configFile)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *Config) fullFileHandling(scanner *bufio.Scanner) error {
	for scanner.Scan() { // Проходим по всему файлу\экрану до конца
		line := scanner.Text() // получем текст из линии
		line = cfg.removeIgnoringLine(line)
		if line == "" {
			continue
		}
		output := cfg.logFileFiltering(line)
		if output != "" {
			fmt.Println(output)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Получает на вход строку в виде лога Squid по-умолчанию от программы
// Проверяет принадлежность адреса источника указанным в конфиге подсетям
// Если принадлежит, то возвращает строку не меняя её
// Если не пренадлежит, проверяет в каком направлении считается трафик, если в оба (значение "both" параметра processing_direction)
// 		то проверяет адрес получателя на напринадежность подсетям указанным в конфиге,
// 		если принадлежит,
// 			то меняет местами адреса получателя и источника
// 		если не принадлежит,
// 			то данная строка исключается (возвращается пустое значение)

func (cfg *Config) logFileFiltering(line string) string {
	valueArray := strings.Fields(line) // разбиваем на поля через пробел
	if len(valueArray) == 0 {          // проверяем длину строки, чтобы убедиться что строка нормально распарсилась\её формат
		return "" // если это не так то возвращаем ничего
	}

	srcIP := valueArray[2]
	srcPortStr := valueArray[3]
	destIPPort := valueArray[6]
	srcPort := strings.Split(strings.Split(srcPortStr, ":")[1], "/")[0]

	for _, subNet := range config.SubNets {
		ok, err := checkIP(subNet, srcIP)
		if err != nil { // если ошибка, то следующая строка
			return "" //  то возвращаем ничего

		}

		if !ok { // если адрес не принадлежит необходимой подсети
			if config.ProcessingDirection == "both" { // если трафик считается в оба направления,
				destIP := strings.Split(destIPPort, ":")[0]
				destPort := strings.Split(destIPPort, ":")[1] //то проверяем адрес назначения
				ok, err := checkIP(subNet, destIP)
				if !ok || err != nil { // если адрес назначения не входит в проверяемую подсеть или проверка вызвала ошибку,
					continue // то переходим к следующей подсети
				}
				//если адрес добрался сюда, значит он входит в подсеть и необходимо поменять адрес назначения и источника
				newSrcPortStr := strings.Split(srcPortStr, ":")[0] + "_REVERSED:" + destPort + "/" + strings.Split(srcPortStr, "/")[1]
				line = fmt.Sprintf("%v %6v %v %v %v %v %v%v %v %v %v", valueArray[0], valueArray[1], destIP, newSrcPortStr, valueArray[4], valueArray[5], srcIP, srcPort, valueArray[7], valueArray[8], valueArray[9])

				return line
			}
			return ""

		}
		return line

	}

	return ""
}

// Получает на вход строку в виде лога Squid по-умолчанию
// Фильтрует от лишних записей по вхождению строк из списка в конфиге
func (cfg *Config) removeIgnoringLine(line string) string {
	for _, ignorItem := range cfg.IgnorList { //проходим по списку исключения,
		if strings.Contains(line, ignorItem) { //если линия содержит хотя бы один объект из списка,
			return "" // то мы её игнорируем и возвращаем ничего

		}
	}
	return line

}

func checkIP(subnet, ip string) (bool, error) {
	maskSubnetTmpl := inetAton(net.ParseIP("255.255.255.255")) // переод маски /32 в int64
	ipInt64 := inetAton(net.ParseIP(ip))

	maskSubnetArray := strings.Split(subnet, "/")            // разбиваю входные данные на подсеть и маску
	subnetInt64 := inetAton(net.ParseIP(maskSubnetArray[0])) // подсеть в int64
	// maskSubnetStr := strings.Split(subnet, "/")[1] //
	maskSubnet, err := strconv.Atoi(maskSubnetArray[1]) // маска в виде Int для проведения битового сдвига
	if err != nil {
		return false, err
	}
	maskSubnetBytes := maskSubnetTmpl << (32 - maskSubnet) // сдиваю маску /32 на оставшееся количество бит после маски
	if subnetInt64 == (ipInt64 & maskSubnetBytes) {        // Проверка на хождение в подсеть IP-адреса
		return true, nil
	}
	return false, nil
}

// Convert uint to net.
// https://groups.google.com/forum/#!topic/golang-nuts/v4eJ5HK3stI
// func inetNtoa(ipnr int64) net.IP {
// 	var bytes [4]byte
// 	bytes[0] = byte(ipnr & 0xFF)
// 	bytes[1] = byte((ipnr >> 8) & 0xFF)
// 	bytes[2] = byte((ipnr >> 16) & 0xFF)
// 	bytes[3] = byte((ipnr >> 24) & 0xFF)

// 	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
// }

// Convert net.IP to int64
// https://groups.google.com/forum/#!topic/golang-nuts/v4eJ5HK3stI
func inetAton(ipnr net.IP) int64 {
	bits := strings.Split(ipnr.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}
