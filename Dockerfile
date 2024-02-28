# build golang app
FROM golang:1.22 as build

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags "-s -w" -o ironmanager .

# build chrome headless-shell
FROM chromedp/headless-shell:latest

WORKDIR /app
# copy golang app to headless shell img
COPY --from=build /app/ ./

ENTRYPOINT [ "/app/ironmanager" ]