FROM golang:onbuild
FROM golang:latest 
RUN mkdir /app 
COPY buildresult /app/
WORKDIR /app
CMD ["/app/main"]