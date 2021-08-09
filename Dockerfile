#docker image build -t new .
#docker run --network host -d new
FROM golang:alpine AS build
WORKDIR /go/src/myapp
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/myapp cmd/main.go
FROM scratch
ADD resources/test.crt /etc/ssl/certs/
ENV NEQUI_QUEUE_NAME="https://sqs.us-east-1.amazonaws.com/851560454673/sqs-lambda-customer-service-create-ticket-comment-qa"

COPY --from=build /go/bin/myapp /go/bin/myapp
EXPOSE 8080
ENTRYPOINT ["/go/bin/myapp"]