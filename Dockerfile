FROM golang:1.25-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG VERSION=dev
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -X main.version=${VERSION}" -o /out/tf-edge-node .

FROM gcr.io/distroless/static-debian12
COPY --from=builder /out/tf-edge-node /app/tf-edge-node
EXPOSE 8090
ENTRYPOINT ["/app/tf-edge-node"]
