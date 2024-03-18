FROM golang:1.21.7
WORKDIR /app
COPY go.mod go.sum fro-part-sharing-firebase-adminsdk-2q8ix-24d2cddd19.json ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
EXPOSE 8080
CMD ["/docker-gs-ping"]

