# Use oven/bun as base image
FROM oven/bun as base

ARG VITE_WEATHER_API_KEY
ARG VITE_NINJA_API_KEY

ENV VITE_API_BASE_URL=$VITE_API_BASE_URL
ENV VITE_API_KEY=$VITE_API_KEY

WORKDIR /usr/src/app

# Install dependencies into temp directory
# This will cache them and speed up future builds
FROM base AS install

RUN mkdir -p /temp/dev

COPY ./frontend/package.json ./frontend/bun.lockb /temp/dev/
RUN cd /temp/dev && bun install --frozen-lockfile

# Install with --production (exclude devDependencies)
RUN mkdir -p /temp/prod
COPY ./frontend/package.json ./frontend/bun.lockb /temp/prod/
RUN cd /temp/prod && bun install --frozen-lockfile --production

# Copy node_modules from temp directory
# Then copy all (non-ignored) project files into the image
FROM install AS prerelease
COPY --from=install /temp/dev/node_modules node_modules
COPY frontend .

RUN bun --bun run vite build

FROM prerelease AS release

# Copy built files
COPY --from=prerelease /usr/src/app/dist dist
# Build server using golang
FROM golang:1.22.2-alpine as build

WORKDIR /go/src/app

COPY backend .

RUN go mod tidy

RUN go build -ldflags="-s -w" -o app

# Final image
FROM alpine:3.14

WORKDIR /app

# Copy server binary
COPY --from=build /go/src/app/app ./app

RUN mkdir static

RUN chmod +x ./app

# Copy frontend/dist into backend's static folder
COPY --from=release /usr/src/app/dist ./static

# Expose port
EXPOSE 3000

RUN cd /app

# Run server
CMD ["./app"]
