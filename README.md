# EN
**GenGoPortScanner** is a high-performance, multi-threaded port scanner written in Go. It is designed to scan ports on a given target quickly and efficiently, providing detailed information about open ports.

## Features

- High-performance multi-threaded scanning
- Supports CIDR notation and files with multiple targets
- Outputs results in various formats (JSON, CSV, Qubo)
- Customizable timeout and retry settings

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/geniuszlyy/GenGoPortScanner.git
    cd GenGoPortScanner
    ```

2. **Initialize Go modules**:
    ```sh
    go mod tidy
    ```

## Usage
```sh
go run .\cmd\gengoscanner\main.go -t <target> -p <ports> -c <threads> --timeout <timeout> -o <output> -f <format>
```

![image](https://github.com/user-attachments/assets/028a2325-094c-44ad-96f7-6d6ba74f2def)


## Parameters
- `-t, --target`: Target CIDR or file with targets (required)
- `-p, --ports`: Ports or port ranges to scan (required)
- `-c, --threads`: Number of concurrent threads (required)
- `--timeout`: Timeout in milliseconds (required)
- `-r, --retries`: Number of ping retries (default: 1)
- `-o, --output`: Output file (default: target.GenGoPortScanner.txt)
- `-f, --format`: Output format (qubo/json/csv) (default: qubo)

## Examples
1. **Scan a single IP with default ports**:
    ```bash
    go run .\cmd\gengoscanner\main.go -t 192.168.1.1 -p 80,443 -c 10 --timeout 500 -o result.txt -f json
    ```
2. **Scan a range of IPs with all ports**:
    ```bash
    go run .\cmd\gengoscanner\main.go -t 192.168.1.0/24 -p 1-65535 -c 100 --timeout 500 -o result.txt -f json
    ```
3. **Scan targets from a file**:
    ```bash
    go run .\cmd\gengoscanner\main.go -t targets.txt -p 80,443 -c 50 --timeout 300 -o output.csv -f csv
    ```
## Output Formats
### JSON
```json
{
    "target": "192.168.1.1",
    "version": "1.16.5",
    "players": "5/100",
    "description": "Minecraft server"
}
```
### CSV
```csv
"192.168.1.1","1.16.5","5/100","Minecraft server"
```
### qubo
```scss
(192.168.1.1)(5/100)(1.16.5)(Minecraft server)
```

# RU
**GenGoPortScanner** — это высокопроизводительный многопоточный сканер портов, написанный на Go. Он предназначен для быстрого и эффективного сканирования портов на заданной цели, предоставляя подробную информацию об открытых портах.

## Особенности
- Высокопроизводительное многопоточное сканирование
- Поддержка CIDR нотации и файлов с несколькими целями
- Вывод результатов в различных форматах (JSON, CSV, Qubo)
- Настраиваемые параметры тайм-аута и количества попыток

## Установка
1. **Клонируйте репозиторий**:
    ```sh
    git clone https://github.com/geniuszlyy/GenGoPortScanner.git
    cd GenGoPortScanner
    ```

2. **Инициализируйте модули Go**:
    ```sh
    go mod tidy
    ```

## Использование
```sh
go run .\cmd\gengoscanner\main.go -t <цель> -p <порты> -c <потоки> --timeout <тайм-аут> -o <файл вывода> -f <формат>
```

![image](https://github.com/user-attachments/assets/5ab6d4a9-9de4-4ff7-ad32-8045b5117b94)


## Параметры
- `-t, --target`: Целевой CIDR или файл с целями (обязательно)
- `-p, --ports`: Порты или диапазоны портов для сканирования (обязательно)
- `-c, --threads`: Количество потоков (обязательно)
- `--timeout`: Таймаут в миллисекундах (обязательно)
- `-r, --retries`: Количество попыток пинга (по умолчанию: 1)
- `-o, --output`: Файл вывода (по умолчанию: target.GenGoPortScanner.txt)
- `-f, --format`: Формат вывода (qubo/json/csv) (по умолчанию: qubo)

## Примеры
1. **Сканирование одного IP с портами по умолчанию**:
    ```bash
    go run .\cmd\gengoscanner\main.go -t 192.168.1.1 -p 80,443 -c 10 --timeout 500 -o result.txt -f json
    ```
2. **Сканирование диапазона IP с всеми портами**:
    ```bash
    go run .\cmd\gengoscanner\main.go -t 192.168.1.0/24 -p 1-65535 -c 100 --timeout 500 -o result.txt -f json
    ```
3. **Сканирование целей из файла**:
    ```bash
    go run .\cmd\gengoscanner\main.go -t targets.txt -p 80,443 -c 50 --timeout 300 -o output.csv -f csv
    ```

## Форматы вывода
### JSON
```json
{
    "target": "192.168.1.1",
    "version": "1.16.5",
    "players": "5/100",
    "description": "Minecraft server"
}
```
### CSV
```csv
"192.168.1.1","1.16.5","5/100","Minecraft server"
```
### qubo
```scss
(192.168.1.1)(5/100)(1.16.5)(Minecraft server)
```
