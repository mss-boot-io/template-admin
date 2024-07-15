FROM alpine:latest

WORKDIR /app

COPY application .

ENTRYPOINT [ "/app/application" ]

CMD [ "server" ]
