FROM golang:latest AS builder

WORKDIR /app
COPY . .

CMD ["cat", "/app/main.go"]

RUN go mod tidy
RUN GOARCH=amd64 GOOS=linux go build -o chat-bot .

# Build the final image
FROM amd64/alpine:latest

# Install necessary dependencies
RUN apk add --no-cache libc6-compat

ENV AUTH=<AUTH-TOKEN>
ENV CLIENT_ID=<CLIENT-ID>
ENV BOT_USERNAME=proofoffaceit
ENV STREAMER_USERNAME=Olesha
ENV COOLDOWN=10
ENV FACEIT_ID=<FACEIT-ID>
ENV FACEIT_API=<FACEIT-API>
ENV LANG=<LANG RU|EN>

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/chat-bot .

# Verify if the binary exists in the final image
RUN ls -l /root/

CMD ["./chat-bot"]

