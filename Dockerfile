# Build UI
FROM --platform=$BUILDPLATFORM node:18.15.0-bullseye-slim AS build-ui
WORKDIR /app
COPY /client .
RUN yarn install
RUN yarn run build

ENV PORT = 8080 

# Build server
FROM golang:1.21-alpine AS build-go
WORKDIR /app
COPY . /app/
RUN rm -rf /src/client
RUN go build -o server cmd/api/main.go 



FROM alpine:3
WORKDIR /app
RUN apk update && apk add ca-certificates && apk upgrade
COPY --from=build-ui /app/dist/ /client/dist/
COPY --from=build-go /app/server /bin/server
EXPOSE 8080

ENTRYPOINT ["/bin/server"]