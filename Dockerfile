#FROM registry.tor.ph/go/hiveon-api:base as build-deps
FROM golang:alpine as build-deps
RUN apk add git curl
ADD . ./src/idp
WORKDIR ./src/idp
RUN mv conf/config.dev.yaml conf/config.yaml
RUN go get -u github.com/gobuffalo/packr/packr && \
    go get ./... && \
    packr build


FROM golang:alpine3.8
#FROM build-deps
RUN mkdir -p /app/conf 
WORKDIR /app
COPY --from=build-deps /go/src/idp/idp /app
COPY --from=build-deps /go/src/idp/conf/. /app/conf
#COPY views  /app/views
ENV build-number=${CI_PIPELINE_ID:-latest}
EXPOSE 3000
CMD ["./idp"]

