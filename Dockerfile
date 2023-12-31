# Build UI
FROM --platform=$BUILDPLATFORM node:18.15.0-bullseye-slim AS build-ui
WORKDIR /src
COPY /client .
RUN yarn install
RUN yarn run build


# Build server
FROM golang:1.21-alpine AS build-go
WORKDIR /app
COPY . .
RUN rm -rf /client
RUN go build -o server cmd/api/main.go 


FROM alpine:3
RUN apk update && apk add ca-certificates && apk upgrade
COPY --from=build-ui /src/dist /client/dist
COPY --from=build-go /app/server /bin/server
EXPOSE 8080

ENTRYPOINT ["/bin/server"]