package db

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/yash-gkmit/NOTE-TAKER/config"
)

func CreateDynamodbTables(configI *config.Config) {
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   &configI.AwsConfig.Region,
			Endpoint: &configI.DynamoEndpoint,
			Credentials: credentials.NewStaticCredentials(
				configI.AwsConfig.AccessKey,
				configI.AwsConfig.SecretKey,
				"",
			),
		},
	}))

	ddb := dynamodb.New(awsSession)

	createUsersTable(ddb)
	createNotesTable(ddb)

}

// CreateUsersTable creates the Users table with composite keys and a GSI on entityType
func createUsersTable(ddb *dynamodb.DynamoDB) {
	input := &dynamodb.CreateTableInput{
		TableName: aws.String("Users"),

		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("userId"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("email"),
				AttributeType: aws.String("S"),
			},
		},

		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("userId"),
				KeyType:       aws.String("HASH"),
			},
		},

		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("UserEmailIndex"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("email"),
						KeyType:       aws.String("HASH"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(5),
					WriteCapacityUnits: aws.Int64(5),
				},
			},
		},

		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := ddb.CreateTable(input)
	if err != nil {
		fmt.Println("Error creating Users table:", err)
		return
	}

	fmt.Println("Users table created successfully!")
}

// createNotesTable creates the Notes table with composite keys and a GSI on entityType
func createNotesTable(ddb *dynamodb.DynamoDB) {
	input := &dynamodb.CreateTableInput{
		TableName: aws.String("Notes"),

		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("noteId"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("userId"),
				AttributeType: aws.String("S"),
			},
		},

		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("noteId"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("userId"),
				KeyType:       aws.String("RANGE"),
			},
		},

		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("UserNotesIndex"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("userId"),
						KeyType:       aws.String("HASH"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(5),
					WriteCapacityUnits: aws.Int64(5),
				},
			},
		},

		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := ddb.CreateTable(input)
	if err != nil {
		fmt.Println("Error creating Notes table:", err)
		return
	}

	fmt.Println("Notes table created successfully!")
}
