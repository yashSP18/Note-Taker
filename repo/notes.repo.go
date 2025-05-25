package repo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/yash-gkmit/NOTE-TAKER/models"
)

type NoteRepo struct {
	Client    *dynamodb.DynamoDB
	GSI       string
	TableName string
}

func NewNoteRepo(ddb *dynamodb.DynamoDB) *NoteRepo {
	return &NoteRepo{
		Client:    ddb,
		GSI:       "UserIndex",
		TableName: "Notes",
	}
}

func (r *NoteRepo) CreateNote(user *models.NoteModel) error {

	data, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user to map: %v", err)
	}

	_, err = r.Client.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(r.TableName),
		Item:      data,
	})
	if err != nil {
		return fmt.Errorf("failed to create user in DynamoDB: %v", err)
	}

	return nil
}
