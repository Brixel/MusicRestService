FROM golang:1.10-windowsservercore-1709 
RUN mkdir /app 
COPY buildresult /app/
WORKDIR /app
CMD ["/app/main"]