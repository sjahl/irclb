FROM alpine:latest

EXPOSE 6667/tcp

RUN apk update && apk add ngircd

COPY ngircd.conf /etc/ngircd/ngircd.conf

ENTRYPOINT /usr/sbin/ngircd -n
