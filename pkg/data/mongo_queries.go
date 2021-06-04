package data

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (wl *MongoWordList) randomWord(ctx context.Context, l int16) (w string, err error) {
	cur, err := wl.words.Aggregate(ctx, selectPipe(l))
	if err != nil {
		return "", fmt.Errorf("failed to aggregate from the db: %w", err)
	}

	defer func() {
		err = cur.Close(ctx)
	}()

	var res []bson.M
	if err = cur.All(ctx, &res); err != nil {
		return "", fmt.Errorf("failed to decode result of the database query: %w", err)
	}

	return res[0]["word"].(string), nil
}

func selectPipe(l int16) mongo.Pipeline {
	matchStage := bson.D{{
		Key: "$match", Value: bson.D{
			{
				Key: "word",
				Value: bson.M{
					"$exists": true,
				},
			},
			{
				Key: "$expr",
				Value: bson.M{
					"$eq": bson.A{
						bson.M{
							"$strLenCP": "$word",
						},
						l,
					},
				},
			},
		},
	}}

	selectStage := bson.D{{
		Key: "$sample", Value: bson.M{
			"size": 1,
		},
	}}

	projectStage := bson.D{{
		Key: "$project", Value: bson.M{
			"_id":  0,
			"word": 1,
		},
	}}

	return mongo.Pipeline{matchStage, selectStage, projectStage}
}
