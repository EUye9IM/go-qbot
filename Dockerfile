FROM alpine

COPY go-qbot setting.yml /bot/

CMD cd /bot && ./go-qbot