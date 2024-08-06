FROM golang:1.22-alpine

WORKDIR /

ARG SERVICE
ARG ENV
ARG PORT

ENV ENV=$ENV
ENV port=$PORT

COPY /pkg/swagger-ui /public/swagger-ui
COPY $SERVICE/swagger.json swagger.json
COPY $SERVICE/$ENV.ini $ENV.ini

COPY $SERVICE/bin /app
RUN chmod +x /app

EXPOSE 18000

CMD ["/app", "serve"]