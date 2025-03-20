# Use Node.js official image
FROM node:18-bullseye

# Install FFmpeg
RUN apt-get update && apt-get install -y ffmpeg && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /app

# Copy package.json and install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy bot files
COPY . .

RUN npm run build

# Compile TypeScript to JavaScript

# Start the bot
CMD ["node", "dist/index.js"]
