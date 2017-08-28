FROM alpine:3.4

RUN apk add --no-cache ca-certificates

ADD slack_bot slack_bot
RUN chmod +x slack_bot

CMD ["./slack_bot"]

