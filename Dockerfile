FROM alpine

COPY go-qbot setting.yml /bot/

CMD cd /bot && ./mirai-ws-bot