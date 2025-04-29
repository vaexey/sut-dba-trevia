# backend
FROM golang:1.24-alpine AS build-back

WORKDIR /app
COPY back/ .
RUN go mod download
RUN CGO_ENABLED=0 go build -o server

# final image
FROM alpine

WORKDIR /app
COPY --from=build-back /app/server .

EXPOSE 7777
CMD ["/app/server"]