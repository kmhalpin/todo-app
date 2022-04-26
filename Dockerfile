FROM node:14-alpine3.14 as node
WORKDIR /fe
COPY web .
RUN npm i
RUN npm run build

FROM golang:1.17-alpine3.14 as go
WORKDIR /be
COPY . .
COPY --from=node /fe/build ./web/build
RUN apk update && apk upgrade
RUN apk add --no-cache gcc musl-dev
RUN go build -o bin/server ./cmd/server

FROM alpine:3.14
WORKDIR /app
COPY --from=node /fe/build ./web/build
COPY --from=go /be/configs ./configs
COPY --from=go /be/bin/server .
COPY --from=go /be/.env .
RUN pass=$(echo date +%s | sha256sum | base64 | head -c 32; echo | mkpasswd) && \
    echo "root:${pass}" | chpasswd
RUN addgroup -g 1002 -S app \
    && adduser -u 1002 -S -D -G app app \
    && chown -R app:app /app
USER app
CMD ["./server"]