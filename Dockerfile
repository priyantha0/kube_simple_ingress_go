FROM golang:1.19-bullseye

RUN apt-get update \
  && apt-get install -y --no-install-recommends --no-install-suggests \
  build-essential

WORKDIR /app

COPY ./go.mod ./
RUN go mod download
copy *.go ./

RUN go build -o ./hellogo

EXPOSE 8085
CMD ["./hellogo"]

