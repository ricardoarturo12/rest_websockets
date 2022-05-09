ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine AS builder

# no hay proxy
RUN go env -w GOPROXY=direct
# instalar git
RUN apk add --no-cache git
# certificados de seguridad
RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /src/

COPY ./go.mod ./go.sum ./

# instala las dependencias
RUN go mod download

COPY ./ ./

# para no compilar c++
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /rest-ws

# encargada de ejecutar la aplicación
FROM scratch AS runner

# copiar los certificados que se descargaron
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY .env ./

COPY --from=builder /rest-ws /rest-ws

EXPOSE 5050

# correr go mod tidy

ENTRYPOINT [ "/rest-ws" ]