FROM iron/go:dev

ARG DATABASE=dev

WORKDIR /app

# Build API
ENV SRC_DIR=/go/src/github.com/berryhill/gf-api/api/
ADD . $SRC_DIR
ENV DATABASE ${DATABASE}
RUN cd $SRC_DIR; go build -o api -tags=${DATABASE}; cp api /app/

ENTRYPOINT ["./api"]
