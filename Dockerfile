FROM golang:1.10
RUN mkdir /app 
COPY buildresult /app/
WORKDIR /app
ENTRYPOINT ["go-wrapper", "run", "app/main.go"]