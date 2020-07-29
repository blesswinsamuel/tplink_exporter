# build stage
FROM --platform=$BUILDPLATFORM golang:1.14-stretch AS build-env

ADD . /src
WORKDIR /src

ENV CGO_ENABLED=0
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o dnsmasq_exporter

# final stage
FROM scratch
WORKDIR /app
COPY --from=build-env /src/dnsmasq_exporter /app/
ENTRYPOINT ["/app/dnsmasq_exporter"]
