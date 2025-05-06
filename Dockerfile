

FROM debian:bookworm-slim

WORKDIR /app

COPY service .
COPY ./config/config.yml /app/config/config.yml

RUN chmod +x ./service

EXPOSE 8082

CMD ["./service"]

