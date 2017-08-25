# Docker Fun

## Install Dependencies

```
glide install
```

## Docker Setup, Build, and Run

Install Docker Glide Image

```
docker run --rm -it -v $PWD:/go/src/github.com/treeder/dockergo -w /go/src/github.com/treeder/dockergo treeder/glide init
# Say No to the question it asks..
```

```
docker run --rm -it -v $PWD:/go/src/github.com/treeder/dockergo -w /go/src/github.com/treeder/dockergo treeder/glide update
```

Build

```
sudo docker build -t berryhill/docker-fun .
```

Run

```
sudo docker run --rm -p 8080:8080 berryhill/hello-docker
```
