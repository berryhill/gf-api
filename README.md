# Gf-Api
Remove sudo if use OSX

## Clone

```
$ git clone git@github.com:berryhill/gf-api.git
```

## Install Dependencies

```
$ cd api
$ glide install
$ cd ..
```

If you don't have glide, click [here](https://github.com/Masterminds/glide#install) for installation instructions

## Docker Setup, Build, and Run

Install Docker Glide Image

```
$ sudo docker run --rm -it -v $PWD:/go/src/github.com/treeder/dockergo -w /go/src/github.com/treeder/dockergo treeder/glide init
# Say No to the question it asks..
```

```
$ sudo docker run --rm -it -v $PWD:/go/src/github.com/treeder/dockergo -w /go/src/github.com/treeder/dockergo treeder/glide update
```

Build / Run

```
$ sudo docker-compose up --build
```

## Frontend
Find the front end [here](https://github.com/berryhill/gf-frontend)
