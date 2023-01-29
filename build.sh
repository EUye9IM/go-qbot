go build -o ./go-qbot
docker build -t go-qbot:latest --network=host .