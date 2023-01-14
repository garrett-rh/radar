FROM golang:1.19 as build

WORKDIR /go/rummage

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o rummage

FROM scratch as runner
COPY --from=0 /go/rummage/rummage .
ENTRYPOINT [ "./rummage" ]
CMD [ "-h" ]