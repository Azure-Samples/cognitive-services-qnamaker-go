#
# BUILD IMAGE
# $ docker build --no-cache -t qnamaker-go-get-answer .
#
# RUN CODE
#
# WINDOWS BASH COMMAND 
# $ winpty docker run -it --rm --name qnamaker-go-get-answer qnamaker-go-get-answer
#
# NON-WINDOWS
# $ docker run -it --rm --name qnamaker-go-get-answer qnamaker-go-get-answer

FROM golang:latest
WORKDIR /go/src/qnamaker
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

RUN go build get-answer.go

RUN ls

ENTRYPOINT ["./get-answer"]
