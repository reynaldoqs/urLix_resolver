package mongodb

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/reynaldoqs/urLix_resolver/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoDataBase struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func NewRechargesRepository(mongoUrl, mongoDB string, duration int) *mongoDataBase {

	mClient := newMongoClient(mongoUrl, duration)

	repo := &mongoDataBase{
		client:   mClient,
		database: mongoDB,
		timeout:  time.Duration(duration) * time.Second,
	}
	return repo

}

func newMongoClient(mongoURL string, mongoTimeout int) *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()

	mongoclient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		err := errors.Wrap(err, "mongodb.newMongoClient")
		log.Fatalln(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		err := errors.Wrap(err, "mongodb.newMongoClient")
		log.Fatalln(err)
	}
	return mongoclient
}

func (mdb *mongoDataBase) GetAll() ([]*domain.Recharge, error) {

	collection := mdb.client.Database(mdb.database).Collection("recharges")
	findOptions := options.Find()

	var results []*domain.Recharge

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		err := errors.Wrap(err, "rechargesrepo.GetAll")
		return nil, err
	}

	for cur.Next(context.TODO()) {

		var elem domain.Recharge
		err := cur.Decode(&elem)
		if err != nil {
			err := errors.Wrap(err, "rechargesrepo.GetAll")
			return nil, err
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		err := errors.Wrap(err, "rechargesrepo.GetAll")
		return nil, err
	}

	cur.Close(context.TODO())
	return results, nil
}
func (mdb *mongoDataBase) Save(recharge *domain.Recharge) error {
	ctx, cancel := context.WithTimeout(context.Background(), mdb.timeout)
	defer cancel()

	collection := mdb.client.Database(mdb.database).Collection("recharges")

	result, err := collection.InsertOne(
		ctx,
		bson.M{
			"ID":          recharge.ID,
			"phoneNumber": recharge.PhoneNumber,
			"company":     recharge.Company,
			"cardNumber":  recharge.CardNumber,
			"status":      recharge.Status,
			"mount":       recharge.Mount,
			"idResolver":  recharge.IDResolver,
			"createdAt":   recharge.CreatedAt,
			"resolvedAt":  recharge.ResolvedAt,
		},
	)

	if err != nil {
		err := errors.Wrap(err, "rechargesrepo.Save")
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		recharge.ID = oid.Hex()
	}

	return err
}
