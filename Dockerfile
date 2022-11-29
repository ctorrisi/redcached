# syntax=docker/dockerfile:1

FROM golang:1.19-bullseye AS build

WORKDIR /build

COPY . ./

RUN make

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /build/redcached /redcached

USER nonroot:nonroot

ENTRYPOINT ["/redcached"]
