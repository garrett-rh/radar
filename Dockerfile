FROM golang:1.19 as build

WORKDIR /go/radar

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o radar
ENTRYPOINT [ "/bin/bash" ]

FROM scratch as runner
COPY --from=0 /go/radar/radar .
ENTRYPOINT [ "./radar" ]
CMD [ "-h" ]
