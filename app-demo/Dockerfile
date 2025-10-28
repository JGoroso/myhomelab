# syntax=docker/dockerfile:1.7

# Build stage (multi-arch friendly)
FROM --platform=$BUILDPLATFORM golang:1.23 AS build
WORKDIR /src

# 1) Deps cache
COPY go.mod go.sum ./
RUN go mod download

# 2) Copy the rest of the code
COPY . .

# 3) Build binary
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -trimpath -ldflags="-s -w" -o /out/app .

# Final minimalist image
FROM gcr.io/distroless/static-debian12
# User no root
USER 65532:65532
COPY --from=build /out/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
