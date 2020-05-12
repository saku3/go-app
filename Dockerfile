FROM golang:latest

ENV APP_ROOT /app

WORKDIR $APP_ROOT
ADD ./app $APP_ROOT

CMD ["go", "run", "main.go"]
