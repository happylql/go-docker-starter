FROM golang 

WORKDIR $GOPATH/src/godocker 

ADD . $GOPATH/src/godocker 

RUN go build main.go 

EXPOSE 9000 

ENTRYPOINT ["./main"]