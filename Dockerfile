FROM golang:alpine

LABEL maintainer="ylqjgm <admin@usebsd.com>" \
	description="Telegram Messenger MTProto zero-configuration proxy server."

RUN apk --no-cache --update add git \
	&& go get github.com/ylqjgm/MTProxy \
	&& cd /go/src/github.com/ylqjgm/MTProxy \
	&& go build -o MTProxy \
	&& cp MTProxy /MTProxy

VOLUME /data
EXPOSE 8822

ENTRYPOINT ["/MTProxy"]