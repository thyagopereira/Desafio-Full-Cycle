FROM golang:1.18 as compiler

WORKDIR /go
COPY . .

RUN go build -o /bin/main ./code/main.go

FROM scratch

WORKDIR bin
COPY --from=compiler /bin/main /bin/main  
ENTRYPOINT [ "/bin/main" ]