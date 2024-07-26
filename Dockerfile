# Use the official Golang image as a build stage
FROM golang:1.22.4-alpine3.20 AS builder

WORKDIR /app

# Install necessary dependencies (uncomment if needed)
# RUN apk add --no-cache make git

# Uncomment if you need to verify git installation
# RUN git version

# Copy the source code into the container
COPY . .

# Build the Go application without version control information
RUN go build -buildvcs=false

# Use a smaller Alpine Linux image for the final stage
FROM alpine:3.14

# Install necessary packages
RUN apk add --update tzdata \
    && cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime \
    && echo "Asia/Jakarta" > /etc/timezone \
    && apk del tzdata \
    && apk add --no-cache bash

# Environment variables
ENV TZ=Asia/Jakarta
ENV USER=be-go
ENV UID=101
ENV GID=101
ENV HOME=/home/$USER
ENV APP_NAME=be-fismed

# Add a group and a user with the specified UID and GID
RUN addgroup -g "$GID" -S "$USER" \
    && adduser --disabled-password -G "$USER" -h "$HOME" -u "$UID" "$USER"

WORKDIR ${HOME}

# Change ownership of the home directory
RUN chown -R ${UID}:${GID} ${HOME}

# Switch to the non-root user
USER ${USER}

# Create a log directory
RUN mkdir log

# Copy the built application and entrypoint script from the builder stage
COPY --from=builder /app/backend_project_fismed ${HOME}/${APP_NAME}
COPY --from=builder /app/docker-entrypoint.sh ${HOME}/

# Make the entrypoint script executable
USER root
RUN chmod +x ${HOME}/docker-entrypoint.sh \
    && chown -R ${UID}:${GID} ${HOME}

# Set the user back to non-root
USER ${USER}

# Expose the port the application runs on
EXPOSE 8080

# Set the working directory
WORKDIR ${HOME}

# Run the entrypoint script and the application
CMD ["./docker-entrypoint.sh", "./be-fismed"]
