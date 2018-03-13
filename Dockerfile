FROM golang:onbuild
FROM golang:latest 
RUN mkdir /app 
COPY application /app/
WORKDIR /app
CMD ["/app/main"]