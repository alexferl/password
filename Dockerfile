FROM golang:1.12.7-alpine as builder
MAINTAINER Alexandre Ferland <aferlandqc@gmail.com>

ENV GO111MODULE=on

WORKDIR /build

RUN apk add --no-cache git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN git describe --tags > /tmp/git_tag 2> /dev/null; if [ $? -ne 0 ]; then git branch --show-current > /tmp/git_tag; fi

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags \
    "-X 'password/cmd.version=$(cat /tmp/git_tag)' \
    -X 'password/cmd.commit=$(git rev-parse --verify HEAD)' \
    -X 'password/cmd.date=$(date +%Y-%m-%dT%T%z)'" \
    && ./password version

FROM scratch
COPY --from=builder /build/password /password

ENTRYPOINT ["/password"]
