# ---------------------------------------------------------------------------------------------- Golang
# Building the go binary
FROM golang:1.25 AS builder
WORKDIR /github.com/jnnkrdb/orasul

# copy the code files
COPY ./ /github.com/jnnkrdb/orasul
RUN ls -lah

# set env vars
ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux

# START BUILD
RUN go mod download && go build -o /orasul ./bin/orasul/main.go

# ---------------------------------------------------------------------------------------------- Final Alpine
FROM alpine:3.22
WORKDIR /

# install neccessary binaries
RUN apk add --no-cache --update openssl

# Copy the Directory Contents
COPY orasul/ /opt/orasul

# create user with home dir
RUN addgroup -S orasul && adduser -S orasul -H -h /opt/orasul/home -s /bin/sh -G orasul -u 3453

# Copy Binary and Frontend Files
COPY --from=builder /orasul /usr/local/bin/orasul

# Set the user for the config and the operator binaries
RUN chmod 744 /usr/local/bin/orasul &&\
    chmod 744 -R /opt/orasul &&\
    chown orasul:orasul /usr/local/bin/orasul &&\
    chown orasul:orasul -R /opt/orasul

USER orasul:orasul

# set the entrypoints
ENTRYPOINT ["/opt/orasul/entrypoint.sh"]
CMD [ "orasul" ]