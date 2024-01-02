FROM golang:1.21-alpine@sha256:4db4aac30880b978cae5445dd4a706215249ad4f43d28bd7cdf7906e9be8dd6b AS builder

RUN apk add --no-cache git sqlite-dev bash gcc libc-dev

WORKDIR /go/src/rye

COPY . .

RUN mkdir -p /out && \
    go build -x -tags "b_tiny" -o /out/rye .

FROM alpine AS bin
COPY --from=builder /out/rye /
ENTRYPOINT ["/rye"]
