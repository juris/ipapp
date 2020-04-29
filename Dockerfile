FROM golang:1.14.2-alpine

WORKDIR /ipapp
COPY main.go /ipapp
RUN apk add --no-cache git upx \
    && go get github.com/gorilla/handlers \
    && go get github.com/julienschmidt/httprouter \
    && go build -o ipapp -ldflags="-s -w" \
    && goupx ipapp

FROM alpine:3.11
COPY --from=0 /ipapp/ipapp /usr/local/bin/ipapp

ENTRYPOINT [ "/usr/local/bin/ipapp" ]
