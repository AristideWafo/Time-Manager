FROM node:22.13-alpine AS build-stage

WORKDIR /app

COPY package*.json ./

RUN npm ci --only=production

COPY . .

RUN npm run build --prod

FROM node:22.13-alpine

RUN npm install -g serve

RUN addgroup -g 1001 -S msc1 && \
    adduser -u 1001 -S g26 -G msc1

WORKDIR /app

COPY --from=build-stage /app/dist/ ./dist

RUN chown -R g26:msc1 /app

USER g26

EXPOSE 4200

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:4200/ || exit 1

# Start the application with serve
CMD ["serve", "-s", "dist", "-l", "4200"]