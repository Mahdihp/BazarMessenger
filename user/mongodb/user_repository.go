package mongodb

import (
	"BazarMessenger/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mongodbUserRepo struct {
	Collection *mongo.Collection
}

func NewMongodbUserRepository(db *mongo.Collection) domain.UserRepository {
	return &mongodbUserRepo{
		Collection: db,
	}
}
func (this *mongodbUserRepo) GetById(filter interface{}, context context.Context) (*domain.User, error) {
	var user *domain.User
	cur, err := this.Collection.Find(context, filter)
	if err != nil {
		return user, err
	}
	if err := cur.Err(); err != nil {
		return user, err
	}
	cur.Close(context)

	if len(user.ID) == 0 {
		return user, mongo.ErrNoDocuments
	}

	return user, nil
}

func (this *mongodbUserRepo) Create(user *domain.User, context context.Context) error {
	_, err := this.Collection.InsertOne(context, user)
	return err
}

func (this *mongodbUserRepo) Update(user *domain.User, context context.Context) error {
	t := &domain.User{}
	filter := bson.D{primitive.E{Key: "id", Value: user}}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "updated_at", Value: time.Now()},
	}}}

	return this.Collection.FindOneAndUpdate(context, filter, update).Decode(t)
}

func (this *mongodbUserRepo) GetByUsername(filter interface{}, context context.Context) (*domain.User, error) {
	var user *domain.User
	cur, err := this.Collection.Find(context, filter)
	if err != nil {
		return user, err
	}
	if err := cur.Err(); err != nil {
		return user, err
	}
	cur.Close(context)

	if len(user.ID) == 0 {
		return user, mongo.ErrNoDocuments
	}

	return user, nil
}
