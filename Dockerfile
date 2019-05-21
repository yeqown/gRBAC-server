FROM alpine

# build
COPY package/bin/grbac-server /usr/bin

EXPOSE 8080
EXPOSE 8081

CMD ["/usr/bin/grbac-server"]