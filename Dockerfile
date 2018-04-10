FROM golang:alpine
COPY server .
ENV PORT 3000
ENV PATH server:$PATH
RUN chmod +x server
EXPOSE 3000
CMD  ["server"]