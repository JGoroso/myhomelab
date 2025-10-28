# syntax=docker/dockerfile:1.7
FROM --platform=$BUILDPLATFORM golang:1.23 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY app-demo ./app-demo
WORKDIR /src/app-demo
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
  go build -trimpath -ldflags="-s -w" -o /out/app .

FROM gcr.io/distroless/static-debian12
USER 65532:65532
COPY --from=build /out/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
