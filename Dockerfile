FROM golang:1.22-bookworm as build-go
WORKDIR /app

ENV CGO_ENABLED=1

COPY server/go.mod server/go.sum ./
RUN go mod download

COPY server/ .
RUN go build -o /babyfood

FROM debian:bookworm-slim
WORKDIR /app

COPY server/db ./db
COPY app/dist static
COPY --from=build-go /babyfood /app/babyfood

ENV PORT "8080"
ENV PLAYGROUND_ENABLED "false"
ENV STATIC_FILES_PATH "./static"
ENV COOKIE_SECURE "false"
ENV TOKEN "test"
ENV SHARE_URL "http://localhost:8080"
EXPOSE 8080

USER nobody

ENTRYPOINT [ "/app/babyfood" ]