# Use a minimal base image to reduce image size
FROM golang:latest

# Install additional packages (e.g., /bin/sh and jq)
RUN apt-get update && apt-get install -y \
    jq \
    dash \
    && rm -rf /var/lib/apt/lists/*

# Install additional tools (e.g., goreleaser)
RUN go install github.com/goreleaser/goreleaser@latest

RUN mkdir /opt/guac

# Set the working directory
WORKDIR /guac

# Copy the source code into the container
COPY . .

# Build the Go application
RUN make build

WORKDIR /guac/bin

# Copy the guacgql binary from your local directory into the container
RUN cp ./guacgql     /opt/guac/
RUN cp ./guaccollect /opt/guac/
RUN cp ./guaccsub    /opt/guac/
RUN cp ./guacingest  /opt/guac/
RUN cp ./guacone  /opt/guac/

# Make the guacgql etc binary executable (if necessary)
RUN chmod +x /opt/guac/guacgql
RUN chmod +x /opt/guac/guaccollect
RUN chmod +x /opt/guac/guaccsub
RUN chmod +x /opt/guac/guacingest
RUN chmod +x /opt/guac/guacone

# Expose a port if your application listens on a specific port
EXPOSE 8080
