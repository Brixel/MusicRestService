FROM golang:alpine
RUN mkdir /app 
COPY server /app/
WORKDIR /app
ENTRYPOINT ["server"]