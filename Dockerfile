# ---------- FRONTEND ----------
FROM node:23-alpine AS build-front

WORKDIR /app
COPY front/package*.json ./
RUN npm install
COPY front/ .
RUN npm run build

# ---------- BACKEND ----------
FROM golang:1.23 AS build-back

WORKDIR /app
COPY back/ .
RUN go mod download
RUN go build -o server

# ---------- FINAL IMAGE ----------
FROM alpine

WORKDIR /app
COPY --from=build-back /app/server .
COPY --from=build-front /app/dist ./static

EXPOSE 6969
CMD ["./server"]
