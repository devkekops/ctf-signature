FROM golang:1.18-alpine AS build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY ./ ./

RUN CGO_ENABLED=0 go build -o /ctf-signature cmd/main.go

FROM golang:1.18-alpine

ENV APP_HOME /go/src/ctf-signature
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY static/ static/
COPY --from=build /ctf-signature $APP_HOME

EXPOSE 80

ENTRYPOINT ["./ctf-signature"]