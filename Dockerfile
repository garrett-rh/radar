FROM golang:1.19 as build

WORKDIR /go/sonar

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o sonar
ENTRYPOINT [ "/bin/bash" ]

FROM scratch as runner
COPY --from=0 /go/sonar/sonar .
ENTRYPOINT [ "./sonar" ]
CMD [ "-h" ]