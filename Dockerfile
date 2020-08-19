FROM golang:1.14.4-alpine AS compiler

WORKDIR /app/src

COPY . .

RUN go build -o gogithubprofiler

FROM alpine

COPY --from=compiler /app/src/gogithubprofiler /gogithubprofiler

ENTRYPOINT ["/gogithubprofiler"]
