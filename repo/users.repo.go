package repo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/yash-gkmit/NOTE-TAKER/models"
)

type UserRepo struct {
	Client    *dynamodb.DynamoDB
	GSI       string
	TableName string
}

func NewUserRepo(ddb *dynamodb.DynamoDB) *UserRepo {
	return &UserRepo{
		Client:    ddb,
		GSI:       "UserEmailIndex",
		TableName: "Users",
	}
}

func (r *UserRepo) CreateUser(user *models.UserModel) error {

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

func (r *UserRepo) GetUserByEmail(email string) (*models.UserModel, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String(r.TableName),
		IndexName:              aws.String(r.GSI),
		KeyConditionExpression: aws.String("email = :email"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":email": {S: aws.String(email)},
		},
		Limit: aws.Int64(1),
	}

	result, err := r.Client.Query(input)
	if err != nil {
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, nil
	}

	var user models.UserModel
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) GetUserById(userId string) (*models.UserModel, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(userId),
			},
		},
	}

	result, err := r.Client.GetItem(input)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from DynamoDB: %w", err)
	}

	if result.Item == nil {
		return nil, fmt.Errorf("user not found")
	}

	var user models.UserModel
	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %w", err)
	}

	return &user, nil
}

// GetAllUsers fetches all users from DynamoDB.
func (r *UserRepo) GetAllUsers() ([]*models.UserModel, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Users"),
	}

	result, err := r.Client.Scan(input)
	if err != nil {
		return nil, err
	}

	var users []*models.UserModel
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users) //converts the DynamoDB items into a slice of UserModel structs.
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepo) DeleteUser(userId string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(userId),
			},
		},
		ConditionExpression: aws.String("attribute_exists(userId)"), // ✅ check for userId instead
	}

	_, err := r.Client.DeleteItem(input)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
