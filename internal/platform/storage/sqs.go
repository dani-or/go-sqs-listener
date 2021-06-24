package sqs

import (
	"log"
	"fmt"
	"os"
	"nequi.com/poc-sqs-listener/internal/domain"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
	//"errors"
)

type S3Repository struct {
	client *sqs.SQS
}

type LineRecord struct {
	Cuenta string
	CifId string
	TipoCRedito string
}

func NewS3Repository() *S3Repository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sqs.New(sess)
	return &S3Repository{
		client : svc,
	}
}

//requiere la variable de entorno 
//export NEQUI_QUEUE_NAME=https://sqs.us-east-1.amazonaws.com/177333342796/sqs-lambda-customer-service-create-ticket-comment-qa
func (r *S3Repository) GetTransactions() (transaction.Transaction, error) {
	transaction, error := transaction.NewTransaction(500,1, 2, 3, "debentura" )
	
	urlResult := &sqs.GetQueueUrlOutput{}
	
	urlResult.SetQueueUrl(os.Getenv("NEQUI_QUEUE_NAME"));
	queueURL := urlResult.QueueUrl;
	msgResult, err := r.client.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: aws.Int64(10),
		//VisibilityTimeout:   timeout,
	})
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	if(msgResult != nil && msgResult.Messages != nil && msgResult.Messages[0] != nil){
		fmt.Println(*msgResult.Messages[0].Body)
	}else{
		fmt.Println("Not messages")
	}
	

	return transaction, error
}