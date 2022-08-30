FROM golang:1.18

#create dir
RUN mkdir /AEsir

# add all local files to DDapp
ADD . /AEsir

WORKDIR /AEsir

COPY go.* ./

RUN go mod download && go mod verify

#RUN go run app.go
RUN go build -o app .

EXPOSE 8080

CMD ["/AEsir/app"]