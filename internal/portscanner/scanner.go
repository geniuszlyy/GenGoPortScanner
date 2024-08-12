package portscanner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"regexp"
	"sync"
	"time"

	"geniuszly.GenGoPortScanner/internal/portscanner/utils"
	"github.com/fatih/color"
	"github.com/projectdiscovery/gologger"
)

type PortScanner struct {
	Hosts        []string
	Ports        []int
	Timeout      int
	OutputFormat string
	OutputFile   string
	Retries      int
	Workers      int
	active       int
	mu           sync.Mutex
}

func NewPortScanner(hosts []string, ports []int, timeout, retries, workers int, outputFile, outputFormat string) *PortScanner {
	return &PortScanner{
		Hosts:        hosts,
		Ports:        ports,
		Timeout:      timeout,
		Retries:      retries,
		Workers:      workers,
		OutputFile:   outputFile,
		OutputFormat: outputFormat,
	}
}

// ping проверяет доступность порта
func (ps *PortScanner) ping(target string) (string, error) {
	conn, err := net.DialTimeout("tcp", target, time.Duration(ps.Timeout)*time.Millisecond)
	if err != nil {
		return "", err
	}

	if _, err := conn.Write([]byte("\x07\x00/\x01_\x00\x01\x01\x01\x00")); err != nil {
		return "", err
	}

	totalLength, err := utils.ReadVarint(conn)
	if err != nil {
		return "", err
	}

	bufTotal := bytes.NewBuffer(nil)
	if _, err = io.CopyN(bufTotal, conn, int64(totalLength)); err != nil {
		return "", err
	}

	packetID, err := utils.ReadVarint(bufTotal)
	if err != nil || uint32(packetID) != uint32(0x00) {
		return "", err
	}

	length, err := utils.ReadVarint(bufTotal)
	if err != nil {
		return "", err
	}

	bufData := make([]byte, length)
	max, err := bufTotal.Read(bufData)
	if err != nil {
		return "", err
	}

	defer conn.Close()
	return string(bufData[:max]), nil
}

// scanTarget сканирует указанный хост и порт
func (ps *PortScanner) scanTarget(host string, port int) {
	target := fmt.Sprintf("%s:%d", host, port)
	var rawData string
	var err error

	for i := 0; i < ps.Retries; i++ {
		rawData, err = ps.ping(target)
		if err != nil {
			if i == ps.Retries-1 {
				gologger.Info().Msgf("Порт %d на хосте %s недоступен", port, host)
				return
			}
		} else {
			break
		}
	}

	data := &utils.Response{}
	if err = json.Unmarshal([]byte(rawData), data); err != nil {
		return
	}

	var rawMOTD string
	results := &utils.ResponseMOTD{}
	if err = json.Unmarshal([]byte(rawData), results); err != nil {
		var result map[string]interface{}
		json.Unmarshal([]byte(rawData), &result)
		rawMOTD = fmt.Sprintf("%s", result["description"])
	} else {
		rawMOTD = results.Description.Text
	}

	motd := regexp.MustCompile(`§[a-fl-ork0-9]|\n`).ReplaceAllString(rawMOTD, "")
	motd = regexp.MustCompile(`\ +|\t`).ReplaceAllString(motd, " ")

	outputResult := utils.OutputResult{
		Target:      target,
		Version:     data.Version.Name,
		Players:     fmt.Sprintf("%d/%d", data.Players.Online, data.Players.Max),
		Description: motd,
	}

	outputStr := utils.FormatQubo(outputResult)
	color.Cyan("%s\n", outputStr)

	ps.mu.Lock()
	defer ps.mu.Unlock()

	f, err := os.OpenFile(ps.OutputFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	if ps.OutputFormat == "csv" {
		outputStr = utils.FormatCSV(outputResult)
	} else if ps.OutputFormat == "json" {
		outputStr = utils.FormatJSON(outputResult)
	}

	f.WriteString(fmt.Sprintf("%s\n", outputStr))
	gologger.Info().Msgf("Порт %d на хосте %s открыт", port, host)
}

// Scan запускает процесс сканирования
func (ps *PortScanner) Scan() {
	var wg sync.WaitGroup
	for _, host := range ps.Hosts {
		for _, port := range ps.Ports {
			for {
				ps.mu.Lock()
				if ps.active < ps.Workers {
					ps.active++
					ps.mu.Unlock()

					wg.Add(1)
					go func(host string, port int) {
						defer wg.Done()
						ps.scanTarget(host, port)
						ps.mu.Lock()
						ps.active--
						ps.mu.Unlock()
					}(host, port)
					break
				}
				ps.mu.Unlock()
				time.Sleep(10 * time.Millisecond) // избегаем "busy waiting"
			}
		}
	}
	wg.Wait()
	gologger.Info().Msg("Сканирование завершено")
}
