FROM golang:1.19-alpine

WORKDIR /app

COPY . .
RUN go mod download

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

EXPOSE 8080

CMD ["app"]