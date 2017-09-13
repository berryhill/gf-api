# Gf-Api

Remove sudo if using OSX

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

## Build / Seed

```
sudo docker-compose -f docker-compose.seed.yml up --build
```
Project will be up at this point.. but if you shutdown, to re-up, continue to 'Run'

## Run
```
$ sudo docker-compose up
```


## Frontend
Find the front end [here](https://github.com/berryhill/gf-frontend)

# Tasks

### Completed
+ **Dockerize**
+ **Setup/Seed DB**
+ **Pagination**

### WIP
+ Continuous Integration / Continuous Deployment
    + *Kubernetes*
    + *Travis CI*
+ Scrapers x5
    + **backcountry**
    + **cabelas**
    + trouts
    +
    +
+ Item Model
    + **Change product to item model**
    + *Refactor product model to compose item models*
    + Implement algorithms to choose best item
+ CMS
    + **CMS admin user .jwt architecture**
    + *restricted CMS endpoints*
+ Filters
    + Endpoints
    + Query params

### TODO
+ Tests
+ Logging
+ Better search
+ Cache
+ Prod Deploy

