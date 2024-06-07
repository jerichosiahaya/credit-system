FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk update --no-cache && apk add --no-cache tzdata
RUN apk add --no-cache ca-certificates 
RUN apk add --no-cache git

WORKDIR /build

COPY .env .env

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -o credit-system . && \mv credit-system /usr/bin

FROM scratch

COPY --from=builder /usr/bin/credit-system /usr/bin/credit-system
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Jakarta /usr/share/zoneinfo/Asia/Jakarta

EXPOSE 8080

CMD ["/usr/bin/credit-system"]