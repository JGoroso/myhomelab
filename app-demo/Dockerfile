# syntax=docker/dockerfile:1.7

FROM --platform=$BUILDPLATFORM golang:1.25.3 AS build
WORKDIR /src

COPY app-demo/go.* ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download || true

COPY app-demo/ ./app-demo
WORKDIR /src/app-demo

ARG TARGETOS TARGETARCH
ENV CGO_ENABLED=0
RUN GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} \
    go build -trimpath -ldflags="-s -w" -o /out/app .

FROM gcr.io/distroless/static-debian12
USER 65532:65532
COPY --from=build /out/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
