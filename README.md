# Golang stress test

A tool for doing stress testing using golang

## Config

```bash
go mod tidy
```

## Usage

The application use three mandatory parameters:
- `--url` or `-u`: url where stress tests will be performed
- `--requests` or `-r`: number of total requests that will be performed
- `--concurrency` or `-c`: number of concurrently requests

### Running

```bash
go run main.go --url "https://httpbin.org/get" --requests 1000 --concurrency 100
```
#### Or
```bash
docker run container_name --url "https://httpbin.org/get" --requests 100 --concurrency 10
```