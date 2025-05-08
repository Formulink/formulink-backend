

FROM debian:bookworm-slim

WORKDIR /app

COPY service .
COPY ./config/config.yml /app/config/config.yml


RUN chmod +x ./service
RUN apt-get update && apt-get install -y ca-certificates
RUN update-ca-certificates



EXPOSE 8082

CMD ["./service"]

