#docker image build -t new .
#docker run --network host -d new
FROM golang:alpine AS build
WORKDIR /go/src/myapp
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/myapp cmd/api/main.go
FROM scratch
ENV NEQUI_QUEUE_NAME=https://sqs.us-east-1.amazonaws.com/177333342796/sqs-lambda-customer-service-create-ticket-comment-qa
COPY --from=build /go/bin/myapp /go/bin/myapp
EXPOSE 8080
ENTRYPOINT ["/go/bin/myapp"]