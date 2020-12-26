package mongo

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const AutoIncIdName = "auto_inc_id"

type AutoIncId struct {
	Id  string `bson:"_id" json:"id"`
	Num int64  `bson:"n" json:"n"`
}

// GetIncId 获取
func GetIncId(ctx context.Context, db *mongo.Database, id string) (int64, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	res := db.Collection(AutoIncIdName).
		FindOneAndUpdate(
			ctx,
			bson.M{"_id": id},
			bson.M{"$inc": bson.M{"n": int64(1)}},
			options.
				FindOneAndUpdate().
				SetUpsert(true).
				SetReturnDocument(options.After),
		)
	if res.Err() != nil {
		return 0, res.Err()
	}

	var incId = new(AutoIncId)
	if err := res.Decode(&incId); err != nil {
		return 0, err
	} else if incId.Num == 0 {
		return 0, errors.New("invalid inc id")
	}

	return incId.Num, nil
}
