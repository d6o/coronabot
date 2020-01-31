FROM golang:1.13 AS builder

ENV DIR /coronabot

COPY . ${DIR}
WORKDIR ${DIR}

RUN CGO_ENABLED=0 go build -o coronabot ./cmd/coronabot/main.go

RUN chmod +X /coronabot/coronabot

FROM scratch

COPY --from=builder /coronabot/coronabot /coronabot

# Copy the certs from the builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/coronabot"]
