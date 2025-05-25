package repo

import (
	"errors"
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
		GSI:       "UserNotesIndex",
		TableName: "Notes",
	}
}

var NoteNotFound = errors.New("Notes not found")

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

func (r *NoteRepo) GetAllNote(userId string) ([]*models.NoteModel, error) {

	input := &dynamodb.QueryInput{
		TableName:              aws.String(r.TableName),
		IndexName:              aws.String(r.GSI),
		KeyConditionExpression: aws.String("userId = :userId"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":userId": {
				S: aws.String(userId),
			},
		},
	}

	result, err := r.Client.Query(input)
	if err != nil {
		return nil, fmt.Errorf("failed to query Notes: %w", err)
	}

	var notes []*models.NoteModel
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &notes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal query result: %w", err)
	}

	return notes, nil
}

func (r *NoteRepo) GetNoteById(userId, noteId string) (*models.NoteModel, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"noteId": {
				S: aws.String(noteId),
			},
			"userId": {
				S: aws.String(userId),
			},
		},
	}

	result, err := r.Client.GetItem(input)
	if err != nil {
		return nil, fmt.Errorf("failed to get note: %w", err)
	}

	if result.Item == nil {
		return nil, NoteNotFound
	}

	var note models.NoteModel
	err = dynamodbattribute.UnmarshalMap(result.Item, &note)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal result: %w", err)
	}

	return &note, nil
}
