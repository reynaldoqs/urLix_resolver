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
		err := errors.Wrap(err, "mongodb.GetAll")
		return nil, err
	}

	for cur.Next(context.TODO()) {

		var elem domain.Recharge
		err := cur.Decode(&elem)
		if err != nil {
			err := errors.Wrap(err, "mongodb.GetAll")
			return nil, err
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		err := errors.Wrap(err, "mongodb.GetAll")
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
		err := errors.Wrap(err, "mongodb.Save")
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		recharge.ID = oid.Hex()
	}

	return err
}

// ReportRepository implementations
func (mdb *mongoDataBase) SaveR(report *domain.RechargeReport) error {
	ctx, cancel := context.WithTimeout(context.Background(), mdb.timeout)
	defer cancel()

	collection := mdb.client.Database(mdb.database).Collection("rechargeReports")

	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"idRecharge":    report.IDRecharge,
			"farmerNumber":  report.FarmerNumber,
			"successful":    report.Successful,
			"codeResponses": report.CodeResponses,
		},
	)

	if err != nil {
		err := errors.Wrap(err, "mongodb.SaveR")
		return err
	}

	return err
}
func (mdb *mongoDataBase) SaveA(report *domain.AdminMsgReport) error {
	ctx, cancel := context.WithTimeout(context.Background(), mdb.timeout)
	defer cancel()

	collection := mdb.client.Database(mdb.database).Collection("adminReports")

	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"idMessage":     report.IDMessage,
			"farmerNumber":  report.FarmerNumber,
			"successful":    report.Successful,
			"codeResponses": report.CodeResponses,
		},
	)

	if err != nil {
		err := errors.Wrap(err, "mongodb.SaveA")
		return err
	}

	return err
}

// Ussd Repository implementation

func (mdb *mongoDataBase) GetUssdActions() ([]*domain.UssdAction, error) {
	collection := mdb.client.Database(mdb.database).Collection("ussdActions")
	findOptions := options.Find()

	var results []*domain.UssdAction

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		err := errors.Wrap(err, "mongodb.GetUssdActions")
		return nil, err
	}

	for cur.Next(context.TODO()) {

		var elem domain.UssdAction
		err := cur.Decode(&elem)
		if err != nil {
			err := errors.Wrap(err, "mongodb.GetUssdActions")
			return nil, err
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		err := errors.Wrap(err, "mongodb.GetUssdActions")
		return nil, err
	}

	cur.Close(context.TODO())
	return results, nil
}

func (mdb *mongoDataBase) GetByAction(action string) (*domain.UssdAction, error) {
	collection := mdb.client.Database(mdb.database).Collection("ussdActions")
	findOptions := options.FindOne()

	var result *domain.UssdAction

	res := collection.FindOne(context.TODO(), bson.M{"action": action}, findOptions)

	err := res.Decode(&result)
	if err != nil {
		err := errors.Wrap(err, "mongodb.GetByAction")
		return nil, err
	}
	return result, nil

}
func (mdb *mongoDataBase) SaveUssd(ussd *domain.UssdAction) error {
	ctx, cancel := context.WithTimeout(context.Background(), mdb.timeout)
	defer cancel()

	collection := mdb.client.Database(mdb.database).Collection("ussdActions")

	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"ussdSteps": ussd.UssdSteps,
			"action":    ussd.Action,
		},
	)

	if err != nil {
		err := errors.Wrap(err, "mongodb.SaveUssd")
		return err
	}

	return err

}
func (mdb *mongoDataBase) UpdateUssd(ussd *domain.UssdAction) error {
	collection := mdb.client.Database(mdb.database).Collection("ussdActions")
	filter := bson.D{{"action", ussd.Action}}

	updateResult := collection.FindOneAndUpdate(context.TODO(), filter, bson.D{
		{"$set", bson.D{
			{"ussdSteps", ussd.UssdSteps},
		}},
	})

	if err := updateResult.Err(); err != nil {
		err := errors.Wrap(err, "mongodb.SaveUssd")
		return err
	}
	return nil

}
