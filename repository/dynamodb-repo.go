package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/vishnunanduz/go-rest-api/entity"
)

type dynamoDBRepo struct {
	tabble string
}

// NewPostrepository

func NewDynamoDBRepository() PostRepo {
	return &dynamoDBRepo{
		tabble: "posts",
	}
}

func createDynamoDBClient() *dynamodb.DynamoDB {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client and return
	return dynamodb.New(sess)

}

func (dm *dynamoDBRepo) Save(post *entity.Post) (*entity.Post, error) {
	dmClient := createDynamoDBClient()

	attributeVal, err := dynamodbattribute.MarshalMap(post)
	if err != nil {
		return nil, err
	}

	// Create item in table
	input := &dynamodb.PutItemInput{
		Item:      attributeVal,
		TableName: aws.String(dm.tabble),
	}

	_, err = dmClient.PutItem(input)
	if err != nil {
		return nil, err
	}

	return post, nil

}

func (dm *dynamoDBRepo) FindAll() ([]entity.Post, error) {
	dmClient := createDynamoDBClient()

	params := &dynamodb.ScanInput{
		TableName: aws.String(dm.tabble),
	}

	result, err := dmClient.Scan(params)
	if err != nil {
		return nil, err
	}

	var posts []entity.Post
	for _, i := range result.Items {
		post := entity.Post{}
		err = dynamodbattribute.UnmarshalMap(i, &post)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}
	return posts, nil

}

func (dm *dynamoDBRepo) FindByID(id string) (*entity.Post, error) {
	// Get a new DynamoDB client
	dynamoClient := createDynamoDBClient()

	result, err := dynamoClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(dm.tabble),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	post := entity.Post{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &post)
	if err != nil {
		panic(err)
	}
	return &post, nil
}

// Delete: TODO
func (repo *dynamoDBRepo) Delete(post *entity.Post) error {
	return nil
}
