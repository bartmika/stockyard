# The base go-image
FROM golang:1.19

COPY . /go/src/github.com/bartmika/stockyard
WORKDIR /go/src/github.com/bartmika/stockyard

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy only `.go` files, if you want all files to be copied then replace `with `COPY . .` for the code below.
COPY . .

# Install our third-party application for hot-reloading capability.
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "install", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -polling -log-prefix=false -build="go build ." -command="./stockyard rpc_serve" -directory="./"

# BUILD
# docker build --rm -t stockyard-dev -f dev.Dockerfile .

# EXECUTE
# docker run -d -p 8000:8000 stockyard-dev
