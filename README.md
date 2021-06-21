# GO-Feed-Me

## An Example micro frontends with Golang

Writeup <https://blog.joe-davidson.co.uk/posts/micro-frontends-with-go-1/>

Demo <https://go-feed-me.joe-davidson.co.uk/>

## Start

Docker-compose version 3+ is required 

```shell
cd docker
docker-compose up
```

You should see the site at <http://127.0.0.1/>

## Troubleshooting

View the logs

```shell
docker-compose logs browse
docker-compose logs content
docker-compose logs random
docker-compose logs basket
docker-compose logs details
docker-compose logs container
docker-compose logs proxy
```