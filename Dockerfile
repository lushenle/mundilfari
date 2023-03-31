# syntax=docker/dockerfile:1
# Build the binary
FROM --platform=$BUILDPLATFORM golang:1.20 as builder

WORKDIR /workspace

# Install upx for compress binary file
RUN apt update && apt install -y upx curl

# Copy the go source
COPY . .

#ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GO111MODULE=on

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Build and compression
ARG TARGETARCH

RUN GOARCH=$TARGETARCH go build -a -installsuffix cgo -ldflags="-s -w" -o bin/server main.go \
    && upx bin/server

# Add migrate tool
#RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
RUN cd bin && wget -q -O - https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-${TARGETARCH}.tar.gz | tar xz

# build server
FROM alpine:3.17.2
WORKDIR /

COPY --from=builder /workspace/bin .

COPY db/migration /migration
COPY app.env start.sh /

ENV GIN_MODE=release

EXPOSE 8080/tcp

CMD ["/server"]
ENTRYPOINT ["/start.sh"]
