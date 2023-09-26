FROM golang:1.20 AS builder

WORKDIR /src

COPY /server/. .
RUN go mod download

#COPY . .
RUN make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
        ca-certificates  \
        netbase \
        ffmpeg \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app
COPY --from=builder /src/entrypoint.sh /app/entrypoint.sh
COPY --from=builder /src/configs/config.yaml /app/configs/config.yaml
COPY --from=builder /src/.well-known /app/.well-known/

WORKDIR /app

RUN chmod +x ./entrypoint.sh

EXPOSE 8001
EXPOSE 9001

CMD ["./entrypoint.sh"]
