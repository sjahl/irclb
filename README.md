# A very terrible idea -- ircLB

Inspired by [ircdns](https://github.com/purpleidea/ircdns), I thought I'd try a lo-fi service discovery mechanism for adding and removing load balancer backends.

## Testing

I've included a Dockerfile and a mostly untouched ngircd.conf file. This lets you build a docker image with ngircd, which you can connect to for testing.

```
docker build -f ./Dockerfile <image name>
docker run -d -p 6667:6667 <image name>
```
