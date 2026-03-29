package repository

import (
	"context"
	"thoropa/internal/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type LinkRepository struct {
	client *dynamodb.Client
	table  string
}

func NewLinkRepository(client *dynamodb.Client) *LinkRepository {
	return &LinkRepository{
		client: client,
		table:  "links",
	}
}

func (r *LinkRepository) Create(ctx context.Context, link *model.Link) error {
	item, err := attributevalue.MarshalMap(link)

	if err != nil {
		return err
	}

	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &r.table,
		Item:      item,
	})

	return err
}

func (r *LinkRepository) FindByID(ctx context.Context, id string) (*model.Link, error) {
	out, err := r.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(r.table),
		KeyConditionExpression: aws.String("id = :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberS{Value: id},
		},
		Limit: aws.Int32(1),
	})

	if err != nil {
		return nil, err
	}

	if len(out.Items) == 0 {
		return nil, nil
	}

	var link model.Link

	err = attributevalue.UnmarshalMap(out.Items[0], &link)
	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (r *LinkRepository) FindByIP(ctx context.Context, ip string) ([]*model.Link, error) {
	var links []*model.Link

	var startKey map[string]types.AttributeValue
	for {
		out, err := r.client.Scan(ctx, &dynamodb.ScanInput{
			TableName:                 aws.String(r.table),
			FilterExpression:          aws.String("ip = :ip"),
			ExpressionAttributeValues: map[string]types.AttributeValue{":ip": &types.AttributeValueMemberS{Value: ip}},
			ExclusiveStartKey:         startKey,
		})
		if err != nil {
			return nil, err
		}

		var page []*model.Link
		err = attributevalue.UnmarshalListOfMaps(out.Items, &page)
		if err != nil {
			return nil, err
		}

		links = append(links, page...)

		if len(out.LastEvaluatedKey) == 0 {
			break
		}
		startKey = out.LastEvaluatedKey
	}

	return links, nil
}
