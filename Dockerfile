# этап сборки (build stage)
FROM node:9.11.1-alpine as build-stage
WORKDIR /public/view
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# этап production (production-stage)
FROM nginx:1.13.12-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

FROM golang:onbuild
EXPOSE 8080