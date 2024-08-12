package parse

import (
	"bufio"
	"net"
	"os"
	"strconv"
	"strings"

	"geniuszly.GenGoPortScanner/internal/portscanner/validate"
	"github.com/projectdiscovery/gologger"
)

// Увеличивает IP-адрес
func inc(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		if ip[i]++; ip[i] > 0 {
			break
		}
	}
}

// Парсит целевой хост или файл с хостами
func ParseTarget(target string) []string {
	var ips []string

	if _, err := os.Stat(target); err == nil {
		file, err := os.Open(target)
		if err != nil {
			gologger.Fatal().Msg("Не удалось открыть файл с целями")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ips = append(ips, ParseTarget(scanner.Text())...)
		}
		return ips
	}

	if strings.Contains(target, "/") {
		ip, ipnet, err := net.ParseCIDR(target)
		if err != nil {
			gologger.Fatal().Msg("Ошибка при разборе CIDR: " + err.Error())
		}
		for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
			ips = append(ips, ip.String())
		}
	} else {
		ips = append(ips, target)
	}

	return ips
}

// Парсит порты из строки
func ParsePorts(port string) []int {
	var ports []int

	for _, portRange := range strings.Split(port, ",") {
		if strings.Contains(portRange, "-") {
			portRangeSplit := strings.Split(portRange, "-")
			start, err := strconv.Atoi(portRangeSplit[0])
			if err != nil {
				gologger.Fatal().Msg("Ошибка при разборе диапазона портов: " + err.Error())
			}
			end, err := strconv.Atoi(portRangeSplit[1])
			if err != nil {
				gologger.Fatal().Msg("Ошибка при разборе диапазона портов: " + err.Error())
			}
			for i := start; i <= end; i++ {
				if validate.ValidatePort(i) {
					ports = append(ports, i)
				}
			}
		} else {
			portInt, err := strconv.Atoi(portRange)
			if err != nil {
				gologger.Fatal().Msg("Ошибка при разборе диапазона портов: " + err.Error())
			}
			if validate.ValidatePort(portInt) {
				ports = append(ports, portInt)
			}
		}
	}

	return ports
}
