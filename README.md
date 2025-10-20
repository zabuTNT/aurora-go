# aurora-go
Aurora inverter communication server written in Go.

This server fetches live data from PowerOne/Aurora PVI series (and compatible) inverters using the Aurora protocol and exposes them over HTTP in plain text, JSON and XML.

Based on Daniele De Santis' PHP InverterPowerMeterLITE project: http://www.desantix.it/index.php?page=show_articles&cmd=show_article&id=233

## Requirements

- Go 1.22+
- Network reachability to the inverter TCP interface (default `1470`)

## Build

Using Go modules:

```bash
go build -o aurora-go
```

Or run directly:

```bash
go run . -r 192.168.0.190 -p 1470 -s 8100
```

## Usage

Run with defaults:

```bash
./aurora-go
```

Flags:

- `-r` Inverter IP (default: `192.168.0.190`)
- `-p` Inverter TCP port (default: `1470`)
- `-s` HTTP server port (default: `8100`)

Example:

```bash
./aurora-go -r 192.168.1.133 -s 80
```

On start you’ll see:

```
Inverter IP:PORT : 192.168.1.133:1470
Simple Data URL : http://localhost:80/
Json Data URL   : http://localhost:80/json/
XML Data URL    : http://localhost:80/xml/
```

## HTTP Endpoints

- `/` Plain text summary
- `/json/` JSON payload
- `/xml/` XML payload
- `/health/` Health probe: performs a lightweight status query and returns `OK` or `ERROR` (HTTP 503)

## Connection robustness

The TCP client uses timeouts and deadlines for write/read, reads exactly 8 bytes per Aurora frame, and retries with exponential backoff. If the inverter goes to standby (e.g., at night), requests will fail fast with `ERROR` and will automatically recover as soon as the inverter is back, without restarting the process.

## Troubleshooting

- Cannot connect / timeouts: verify IP/port, firewall rules, and that the inverter’s TCP interface is reachable.
- Frequent `ERROR` at night: expected when the inverter is in standby; it will recover automatically in the morning.
- Health check: `curl -i http://localhost:8100/health/` to quickly verify connectivity.
