FROM golang:1.17 AS build

WORKDIR /api
COPY . .
RUN go mod tidy
RUN go build -o server ./cmd/

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /api/server /server

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT ["/server"]