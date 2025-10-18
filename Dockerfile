# Build
FROM golang:1.23.11-alpine3.22 AS artifacts

# disable cgo to avoid gcc requirement bug
ENV CGO_ENABLED=0

WORKDIR /app

# dependencies
COPY go.mod ./
RUN go mod download

# copy app
COPY . ./

# build
RUN go build -o argocd-example cmd/argocd-example/main.go

# Release
FROM golang:1.23.11-alpine3.22 AS release

# disable cgo to avoid gcc requirement bug
ENV CGO_ENABLED=0

RUN apk --no-cache add tini ca-certificates

WORKDIR /app

COPY --from=artifacts /app/argocd-example ./bin/argocd-example

ENTRYPOINT ["tini", "-g", "--"]
CMD ["/app/bin/argocd-example"]
