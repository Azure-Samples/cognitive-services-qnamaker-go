#
# BUILD IMAGE
# $ docker build --no-cache -t qnamaker-go-create-kb .
#
# RUN CODE
#
# WINDOWS BASH COMMAND 
# $ winpty docker run -it --rm --name qnamaker-go-create-kb qnamaker-go-create-kb
#
# NON-WINDOWS
# $ docker run -it --rm --name qnamaker-go-create-kb qnamaker-go-create-kb

FROM golang:latest
WORKDIR /go/src/qnamaker
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

RUN go build create-new-knowledge-base.go

RUN ls

ENTRYPOINT ["./create-new-knowledge-base"]
