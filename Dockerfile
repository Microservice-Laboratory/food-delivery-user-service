FROM golang:1.25.0-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/ ./cmd
COPY configs/ ./configs
COPY internal/ ./internal

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o food-delivery-user-service ./cmd/main.go

FROM gcr.io/distroless/static

WORKDIR /app

COPY --from=build /src/food-delivery-user-service .

ENTRYPOINT ["./food-delivery-user-service"]