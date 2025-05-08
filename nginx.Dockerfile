# front
FROM node:slim AS build-front

WORKDIR /app
COPY front/package*.json ./
RUN npm install
COPY front/ .
RUN npm run build

# nginx
FROM nginx:stable

COPY back/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build-front /app/dist /usr/share/nginx/html
