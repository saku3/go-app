FROM golang:latest

ENV APP_ROOT /app

WORKDIR $APP_ROOT
ADD . $APP_ROOT

CMD ["go", "run", "main.go"]
