# The base go-image
FROM golang:1.16.7-bullseye

# Create a directory for the app
RUN mkdir /app
 
# Copy all files from the current directory to the app directory
COPY . /app
 
# Set working directory
WORKDIR /app

RUN go mod download
# Run command as described:
# go build will build an executable file named server in the current directory
RUN go build -o main .
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]