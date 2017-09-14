FROM iron/go:dev

ARG DATABASE=dev

WORKDIR /app

## Install glide
#RUN mkdir /tools
#WORKDIR /tools
#RUN wget https://github.com/Masterminds/glide/releases/download/0.10.2/glide-0.10.2-linux-386.tar.gz
#RUN tar -zxvf glide-0.10.2-linux-386.tar.gz
#RUN mv linux-386/ glide/
#ENV PATH /tools/glide:$PATH

# Build API
ENV SRC_DIR=/go/src/github.com/berryhill/gf-api/api/
ADD . $SRC_DIR
ENV DATABASE ${DATABASE}
RUN cd $SRC_DIR; go build -o api -tags=${DATABASE}; cp api /app/

ENTRYPOINT ["./api"]
