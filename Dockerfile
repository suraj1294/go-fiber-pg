# Build UI
FROM --platform=$BUILDPLATFORM node:18.15.0-bullseye-slim AS build-ui
RUN apt update && apt install -y git
COPY /client /src/
RUN cd /src && yarn install
RUN cd /src && yarn run build

# Build server
FROM golang:1.21-alpine AS build-go
COPY . /src/
RUN rm -rf /src/client
RUN cd /src && go build cmd/api/main.go

FROM alpine:3
RUN apk update && apk add ca-certificates && apk upgrade
COPY --from=build-ui /src/dist/ /client/dist/
COPY --from=build-go /src/main /bin/main
COPY --from=build-go /src/.env /bin/.env
EXPOSE 8080

ENTRYPOINT ["/bin/main"]