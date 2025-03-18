# Choose whatever you want, version >= 1.16
FROM golang:1.24-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

RUN apk --update add --no-cache ca-certificates ffmpeg opus python3
RUN wget --no-check-certificate https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -O /usr/local/bin/yt-dlp
RUN chmod a+rx /usr/local/bin/yt-dlp
RUN ln -sf /usr/bin/python3 /usr/bin/python
RUN yt-dlp --version

CMD ["air", "-c", ".air.toml"]