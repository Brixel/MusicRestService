FROM golang:1.10
RUN mkdir /app 
COPY buildresult /app/
RUN go-wrapper download
RUN go-wrapper install
WORKDIR /app
ENTRYPOINT ["go-wrapper", "run", "app/main.go"]