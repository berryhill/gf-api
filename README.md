# Gf-Api

### -------------
## IMPLEMENTING SEEING DATA FOR DYNAMIC PORTIONS OF THE APP.. STAND BY IF TRYING TO UP
### -------------

Remove sudo if use OSX

### API SPEC

[API Spec](./docs/api-spec.md)

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

# Tasks

### Completed
+ Dockerize

### WIP
+ Continuous Integration / Continuous Deployment
+ Scrapers x5
+ Item Model (to be composed by the Product Model)
+ Setup/Seed DB

### TODO
+ Tests
+ Logging
+ Cache
+ Prod Deploy

