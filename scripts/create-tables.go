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

	CreateUsersTable(ddb)

}

// CreateUsersTable creates the Users table with composite keys and a GSI on entityType
func CreateUsersTable(ddb *dynamodb.DynamoDB) {
	tableName := "Users"

	// Check if table already exists
	_, err := ddb.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})
	if err == nil {
		fmt.Println("Users table already exists")
		return
	}

	input := &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),

		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("pk"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("sk"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("entityType"),
				AttributeType: aws.String("S"),
			},
		},

		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("pk"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("sk"),
				KeyType:       aws.String("RANGE"),
			},
		},

		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("EntityIndex"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("entityType"),
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

	_, err = ddb.CreateTable(input)
	if err != nil {
		fmt.Println(" Error creating NoteTaking table:", err)
		return
	}

	fmt.Println("NoteTaking table created successfully")
}
