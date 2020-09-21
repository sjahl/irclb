# A very terrible idea -- ircLB

Inspired by [ircdns](https://github.com/purpleidea/ircdns), I thought I'd try a lo-fi service discovery mechanism for adding and removing load balancer backends.

Still very much a work in progress that I hack on when I'm bored.

## Current Status

Right now, the daemon just sits in the specified IRC channel, and responds to messages that begin with "add", and quote the message in the reply. More to come later!

## Testing

I've included a Dockerfile and a mostly untouched ngircd.conf file. This lets you build a docker image with ngircd, which you can connect to for testing.

```
docker build . -t <image name>
docker run -d -p 6667:6667 <image name>
```
