FROM scratch

WORKDIR /

# Copy the pre-built binary and config
COPY app .
COPY .env.example .env

EXPOSE 8080

ENTRYPOINT ["./app"]