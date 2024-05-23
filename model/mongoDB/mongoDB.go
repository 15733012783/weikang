package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCli struct {
	*mongo.Client
	Url        string
	Database   string
	Collection string
}

func InitMongoDB(db ...string) *MongoCli {
	return &MongoCli{
		Url:        db[0],
		Database:   db[1],
		Collection: db[2],
	}
}

func (m *MongoCli) withClient(ctx context.Context, fun func(cli *mongo.Collection) error) error {
	clientOptions := options.Client().ApplyURI(m.Url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)
	coll := client.Database(m.Database).Collection(m.Collection)
	return fun(coll)
}

func (m *MongoCli) InstallOne(ctx context.Context, doc interface{}) error {
	return m.withClient(ctx, func(cli *mongo.Collection) error {
		_, err := cli.InsertOne(ctx, doc)
		return err
	})
}

func (m *MongoCli) InstallAny(ctx context.Context, doc ...interface{}) error {
	return m.withClient(ctx, func(cli *mongo.Collection) error {
		_, err := cli.InsertMany(ctx, doc)
		return err
	})
}

func (m *MongoCli) FindOne(ctx context.Context, filter interface{}) (doc interface{}, err error) {
	err = m.withClient(ctx, func(cli *mongo.Collection) error {
		return cli.FindOne(ctx, filter).Decode(&doc)
	})
	if err != nil {
		return nil, err
	}
	return
}

func (m *MongoCli) FindAny(ctx context.Context, filter interface{}) (docs []bson.D, errs error) {
	errs = m.withClient(ctx, func(cli *mongo.Collection) error {
		find, err := cli.Find(ctx, filter)
		if err != nil {
			return err
		}
		defer find.Close(ctx)
		for find.Next(ctx) {
			var doc bson.D
			err = find.Decode(&doc)
			if err != nil {
				return err
			}
			docs = append(docs, doc)
		}
		return nil
	})
	if errs != nil {
		return nil, errs
	}
	return
}

func (m *MongoCli) DeleteOne(ctx context.Context, filter interface{}) error {
	return m.withClient(ctx, func(cli *mongo.Collection) error {
		_, err := cli.DeleteOne(ctx, filter)
		return err
	})
}

func (m *MongoCli) DeleteByID(ctx context.Context, id int64) error {
	return m.withClient(ctx, func(cli *mongo.Collection) error {
		_, err := cli.DeleteOne(ctx, bson.D{{"_id", id}})
		return err
	})
}

func (m *MongoCli) UpdateOne(ctx context.Context, filter, data interface{}) error {
	return m.withClient(ctx, func(cli *mongo.Collection) error {
		_, err := cli.UpdateOne(ctx, filter, data)
		return err
	})
}

func (m *MongoCli) UpdateByID(ctx context.Context, id int64, data interface{}) error {
	return m.withClient(ctx, func(cli *mongo.Collection) error {
		_, err := cli.UpdateByID(ctx, bson.D{{"_id", id}}, data)
		return err
	})
}
