# Choose Ubuntu as the base image
FROM golang:alpine

# Set the working directory
WORKDIR /app

# Install necessary packages
RUN apk update && \
    apk add -y --no-cache \
    nodejs \
    npm \
    ffmpeg \
    python3 \
    yt-dlp \
    build-base

# Create and move into the encoder directory
WORKDIR /app/encode

# Copy package.json and package-lock.json first (for caching dependencies)
COPY encoder/package*.json ./

# Install NPM dependencies inside /encode
RUN npm install

# Go back to /app for Go application
WORKDIR /app

# Copy everything else **after** installing dependencies
COPY . .

# Install the air package (for live-reloading Go projects)
RUN go install github.com/air-verse/air@latest

# Set the entry point to activate the virtual environment and run air
ENTRYPOINT ["air"]
