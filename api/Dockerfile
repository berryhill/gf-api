FROM iron/go:dev

WORKDIR /app

ENV SRC_DIR=/go/src/github.com/berryhill/gf-api/
ADD . $SRC_DIR
RUN cd $SRC_DIR; go build -o api; cp api /app/

ENTRYPOINT ["./api"]
