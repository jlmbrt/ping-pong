FROM golang as builder
WORKDIR /build
COPY ./*.go ./go.mod /build/
RUN CGO_ENABLED=0 go build -o app .

CMD ["/build/app"]

FROM scratch
COPY --from=builder /build/app /app
CMD ["/app"]