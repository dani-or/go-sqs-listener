#docker image build -t new .
#docker run --network host -d new
FROM golang:alpine AS build
WORKDIR /go/src/myapp
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/myapp cmd/main.go
FROM scratch
ADD resources/test.crt /etc/ssl/certs/
ENV NEQUI_QUEUE_NAME="https://sqs.us-east-1.amazonaws.com/851560454673/sqs-lambda-customer-service-create-ticket-comment-qa"

#ENV AWS_ACCESS_KEY_ID="AKIA4MRHVLIIW7W3CUPY"
#ENV AWS_SECRET_ACCESS_KEY="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"
#ENV AWS_DEFAULT_REGION="us-east-1"
COPY --from=build /go/bin/myapp /go/bin/myapp
EXPOSE 8080
ENTRYPOINT ["/go/bin/myapp"]