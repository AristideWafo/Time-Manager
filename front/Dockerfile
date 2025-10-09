FROM node:18 AS build-stage
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run build --prod

# FROM nginx:alpine
# COPY --from=build-stage /app/dist /usr/share/nginx/html