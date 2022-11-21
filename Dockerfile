###
### Build Stage
###

# The base go-image
FROM golang:1.19-alpine as build-env

# Create a directory for the app
RUN mkdir /app

# Copy all files from the current directory to the app directory
COPY . /app

# Set working directory
WORKDIR /app

# Run command as described:
# go build will build a 64bit Linux executable binary file named server in the current directory
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o stockyard .

###
### Run stage.
###

FROM alpine:3.14

# Copy required binary executable into this image.
COPY --from=build-env /app/stockyard .

EXPOSE 8000

# Run the server executable
CMD [ "/stockyard", "rpc_serve" ]

# BUILD
# (Run following from project root directory!)
# docker build  -f ./docker/app/Dockerfile --rm -t registry.digitalocean.com/bci/stockyardend-app:latest -f Dockerfile .

# EXECUTE
# docker run -d -p 8000:8000 stockyard-app

# UPLOAD
# docker push registry.digitalocean.com/bci/stockyardend-app:latest
