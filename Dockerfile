# === Build Stage ===
FROM golang:1.25-bookworm AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download

# ===== Dev Image =====
FROM builder AS dev

RUN go install github.com/cosmtrek/air@v1.43.0

COPY . .

CMD [ "air", "-c", ".air.toml" ]

# ===== Production Image =====
FROM builder AS production


COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -tags netgo -ldflags="-s -w" -o api-golang .

# ========= Runtime Image =========
FROM gcr.io/distroless/base-debian12:nonroot

ENV GIN_MODE=release

WORKDIR /

COPY --from=production /build/api-golang .

EXPOSE 8000

CMD [ "/api-golang" ]
