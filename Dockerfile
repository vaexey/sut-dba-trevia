# front
FROM node:23-alpine AS build-front

WORKDIR /app
COPY front/package*.json ./
RUN npm install
COPY front/ .
RUN npm run build

# backend
FROM golang:1.23-alpine AS build-back

WORKDIR /app
COPY back/ .
RUN go mod download
RUN CGO_ENABLED=0 go build -o server

# final image
FROM alpine

WORKDIR /app
COPY --from=build-back /app/server .
COPY --from=build-front /app/dist ./static

EXPOSE 7777
CMD ["/app/server"]