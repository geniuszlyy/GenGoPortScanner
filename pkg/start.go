package pkg

import (
	"fmt"
	"os"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/projectdiscovery/gologger"

	"geniuszly.GenGoPortScanner/internal/portscanner"
	"geniuszly.GenGoPortScanner/internal/portscanner/parse"
)

const banner = `
    __
 __/  \__       GenGoPortScanner   Сканер портов. 🐯
/  \__/  \__
\__/  \__/  \   Версия  1.0.0
   \__/  \__/   Автор   @geniuszly

`

// Start запускает выполнение программы
func Start() {
	fmt.Printf("%s", banner)

	parser := argparse.NewParser("GenGoPortScanner", "Сканер портов. 🐯")
	target := parser.String("t", "target", &argparse.Options{Required: true, Help: "Целевой CIDR или файл с целями"})
	portRange := parser.String("p", "ports", &argparse.Options{Required: true, Help: "Диапазоны портов для сканирования"})
	threads := parser.Int("c", "threads", &argparse.Options{Required: true, Help: "Количество потоков"})
	timeout := parser.Int("", "timeout", &argparse.Options{Required: true, Help: "Таймаут в миллисекундах"})
	retries := parser.Int("r", "retries", &argparse.Options{Required: false, Help: "Количество попыток пинга", Default: 1})
	outputFile := parser.String("o", "output", &argparse.Options{Required: false, Help: "Файл вывода", Default: nil})
	outputFmt := parser.String("f", "format", &argparse.Options{Required: false, Help: "Формат вывода (qubo/json/csv)", Default: "qubo"})

	if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	// Проверка параметров
	if *target == "" || *portRange == "" || *threads <= 0 || *timeout <= 0 {
		fmt.Print(parser.Usage(nil))
		return
	}

	hosts := parse.ParseTarget(*target)
	ports := parse.ParsePorts(*portRange)

	var output string
	if *outputFile == "" {
		output = fmt.Sprintf("%s.GenGoPortScanner.txt", strings.ReplaceAll(*target, "/", "_"))
	} else {
		output = *outputFile
	}

	f, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		gologger.Fatal().Msg("Не удалось открыть файл вывода: " + err.Error())
	}
	f.Close()

	gologger.Info().Msg("Файл вывода установлен на '" + output + "'")

	scanner := portscanner.NewPortScanner(hosts, ports, *timeout, *retries, *threads, output, *outputFmt)
	scanner.Scan()

	gologger.Info().Msg("Сканирование завершено")
}
