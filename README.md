# swappable-logging
Swapping the logging library for golang

## Build
```sh
go mod tidy
go build -o service $PWD/cmd/service
```

## Usage
```sh
./service -h
Usage of ./service:
  -h, --help            Prints this help content.
  -l, --logger string   Required logger. logrus/zap. (default "logrus")

./service -l logrus
#./service -l zap
#./service -l random
```
