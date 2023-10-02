# Build stage
FROM golang:1.21.1 AS BuildStage

ENV CGO_ENABLED=1
ENV GO111MODULE=on

ENV APP_HOME /go/src/app
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"

RUN apt-get update && apt-get install -y make gcc

COPY . .

RUN go mod download

RUN make build

FROM golang:1.21.1 AS DeployStage

WORKDIR /app

RUN apt-get update
RUN apt-get install -y curl libssl-dev build-essential glibc-source gcc


COPY --from=BuildStage /go/src/app/articles .
COPY --from=BuildStage /go/src/app/docker-entrypoint.sh .

RUN chmod +x docker-entrypoint.sh
RUN chmod +x articles


ENTRYPOINT ["sh", "/app/docker-entrypoint.sh"]


