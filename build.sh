CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w --extldflags "-static -fpic"' -o ./go-qbot
docker build -t go-qbot:latest --network=host .