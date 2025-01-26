#the build stage
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo app.go


#the run stage
FROM scratch
WORKDIR /app
COPY --from=builder /app/app .
COPY static ./static
EXPOSE 80
CMD ["./app"]
