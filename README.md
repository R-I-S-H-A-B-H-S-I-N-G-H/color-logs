
# color-logs 🖍️

A simple, blazingly fast, and lightweight command-line tool to color-code log levels (`DEBUG`, `INFO`, `ERROR`, `WARN`) for better readability in your terminal.

Written in Go, `colo` is a single binary designed to be piped with your existing logging commands, adding color without sacrificing performance.

## Demo

**Before `colo`:**

```log
2025-09-17 23:40:15 INFO: Server successfully started on port 8080.
2025-09-17 23:40:16 DEBUG: Checking database connection...
2025-09-17 23:40:17 DEBUG: Database connection established.
2025-09-17 23:40:18 WARN: Configuration value 'timeout' is deprecated. Use 'connection_timeout' instead.
2025-09-17 23:40:19 ERROR: Failed to process request 1A45C: user not found.
```

**After `colo` (`cat your.log | colo`):**

Your terminal output will be transformed, with each log level getting its own distinct, readable background color, making it easy to spot important events at a glance.

  - `INFO` lines get a **green** background.
  - `DEBUG` lines get a **blue** background.
  - `WARN` and `WARNING` lines get a **yellow** background.
  - `ERROR` lines get a **red** background.

## Features

  - **Blazingly Fast:** Built with Go, it can process millions of log lines per second with minimal CPU and memory overhead.
  - **Simple & Lightweight:** It's a single, dependency-free binary. Just download and run.
  - **Easy to Use:** Integrates seamlessly into your workflow using standard Unix pipes.
  - **Highlights Standard Log Levels:** Recognizes and highlights `DEBUG`, `INFO`, `WARN`, `WARNING`, and `ERROR`.

## Installation

### For Linux Users (Recommended)

You can install `colo` with a single command. It will download the binary, make it executable, and move it to `/usr/local/bin` so it's available system-wide.

```bash
curl -L https://github.com/R-I-S-H-A-B-H-S-I-N-G-H/color-logs/raw/main/dist/colo -o colo && \
chmod +x colo && \
sudo mv colo /usr/local/bin/
```

### From Source

If you have Go installed, you can build it from the source.

```bash
# 1. Clone the repository
git clone https://github.com/R-I-S-H-A-B-H-S-I-N-G-H/color-logs.git
cd color-logs

# 2. Build the binary
go build -o colo .

# 3. (Optional) Move it to your system's PATH
sudo mv colo /usr/local/bin/
```

## Usage

The `colo` tool is designed to be used with pipes (`|`). Simply pipe the output of any command that produces log text into `colo`.

### Tailing a log file

```bash
tail -f /var/log/application.log | colo
```

### Viewing a whole log file

```bash
cat server.log | colo
```

### Docker container logs

```bash
docker logs -f my_container | colo
```

### Journalctl (for systemd services)

```bash
journalctl -fu my_service.service | colo
```

### Kubernetes pod logs

```bash
kubectl logs -f my_pod | colo
```

## Uninstallation

To remove `colo` from your system, simply delete the binary:

```bash
sudo rm /usr/local/bin/colo
```

## License

This project is licensed under the MIT License.
