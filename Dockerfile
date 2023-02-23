FROM golang:1.18.1 AS builder
RUN go env -w GOPROXY=direct
WORKDIR /backend
COPY . .
RUN go mod download
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o api .

FROM gcr.io/distroless/static:nonroot
COPY --from=builder --chown=nonroot:nonroot /backend ./backend
EXPOSE 8090
ENTRYPOINT [ "./main"]