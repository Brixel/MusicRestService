FROM golang:alpine
COPY server .
ENV PORT 3000
EXPOSE 3000
ENTRYPOINT ["server"]