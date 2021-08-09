aws ecr create-repository --repository-name node-with-ddd
aws ecr create-repository --repository-name node-sqs-listener
aws ecr create-repository --repository-name node-poc-reading-files
aws ecr create-repository --repository-name go-with-ddd
aws ecr create-repository --repository-name go-sqs-listener
aws ecr create-repository --repository-name go-poc-reading-files

aws dynamodb create-table \
    --table-name credit-customer-product-qa \
    --attribute-definitions \
    AttributeName=customerId,AttributeType=S \
    AttributeName=debenture,AttributeType=S \
    --key-schema \
    AttributeName=customerId,KeyType=HASH \
    AttributeName=debenture,KeyType=RANGE \
    --billing-mode=PAY_PER_REQUEST


aws s3api create-bucket --bucket nequi-s3-select-tmp --region us-east-1 

aws sqs create-queue --queue-name sqs-lambda-customer-service-create-ticket-comment-qa


