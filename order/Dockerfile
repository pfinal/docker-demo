FROM golang:1.11-alpine3.8

COPY ./myapp /go/src/myapp

RUN cd /go/src/myapp && go build && go install

EXPOSE 8080

CMD ["myapp"]
