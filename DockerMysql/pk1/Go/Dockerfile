FROM golang:1.15

#RUN mkdir /Aesir

#ADD . /Aesir

#WORKDIR /Aesir

WORKDIR /go/src/app
COPY . .

#COPY go.* ./

RUN go mod download && go mod verify

#Used for recompiling program after file edits
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main

EXPOSE 8088

#CMD ["/Aesir/main"]