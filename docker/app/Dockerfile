FROM golang:1.15.2-alpine3.12 as build

WORKDIR /go/src/account-management/root

COPY ./root .

# RUN export AWS_ACCESS_KEY_ID=$ACCESS_KEY
# RUN export AWS_SECRET_ACCESS_KEY=$SECRET_KEY
# RUN export AWS_DEFAULT_REGION=$REGION

ENV GO111MODULE=on
ENV CGO_ENABLED=0

# cosmtrek/airのインストール
RUN go get -u github.com/cosmtrek/air
RUN go get golang.org/x/tools/gopls

RUN go mod download

CMD air -c air.toml

# FROM alpine:3.11

# WORKDIR /app

# COPY --from=build /usr/local/go/src/account-management .

# RUN set -x && \
#   addgroup go && \
#   adduser -D -G go go && \
#   chown -R go:go /app/root


# CMD ["./root"]
