# ---- Build stage ----
FROM golang:1.26-alpine AS build

WORKDIR /go/src/backend

# Copy only dependency manifests first — enables Docker layer caching.
# `go mod download` only re-runs when go.mod/go.sum actually change.
COPY go.mod go.sum ./
RUN go mod download

# Now copy the rest of the source
COPY . .

# Build the API binary specifically from cmd/api, not repo root
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -a -installsuffix cgo -o app ./cmd/api

# Build the seed binary too, if you want it in the image for one-off admin seeding
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -a -installsuffix cgo -o seed ./cmd/seed


# ---- Final stage ----
FROM alpine:latest AS release

WORKDIR /app

RUN apk -U upgrade \
  && apk add --no-cache dumb-init ca-certificates

COPY --from=build /go/src/backend/app .
COPY --from=build /go/src/backend/seed .

# Needed if you run migrations at startup or the seed script reads SQL files
COPY migrations ./migrations

RUN chmod +x /app/app /app/seed

EXPOSE 3000

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["./app"]