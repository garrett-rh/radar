FROM golang:1.19 as base

WORKDIR /go/radar

COPY . .

FROM base as build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o radar

FROM base as test
RUN go test ./...

FROM scratch as runner
COPY --from=1 /go/radar/radar .
ENTRYPOINT [ "./radar" ]
CMD [ "-h" ]
