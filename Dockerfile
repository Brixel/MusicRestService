FROM golang:1.10
RUN mkdir /app 
COPY server /app/
WORKDIR /app
ENTRYPOINT ["server"]