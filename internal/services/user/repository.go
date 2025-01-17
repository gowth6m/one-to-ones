package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"one-to-one/internal/db"
)

type UserRepository interface {
	CreateUser(c context.Context, user User) (User, error)
	GetAllUsers(c context.Context) ([]User, error)
	GetUserByID(c context.Context, id primitive.ObjectID) (*User, error)
	GetUserByEmail(c context.Context, email string) (*User, error)
	
	AddReportee(c context.Context, userID primitive.ObjectID, reporteeID primitive.ObjectID) error
	RemoveReportee(c context.Context, userID primitive.ObjectID, reporteeID primitive.ObjectID) error
	AddReportsTo(c context.Context, userID primitive.ObjectID, reportsToID primitive.ObjectID) error
}

type repositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepository() UserRepository {
	collection := db.Client.Database(db.DATABASE_NAME).Collection(db.COLLECTION_USER)
	return &repositoryImpl{collection: collection}
}

func (r *repositoryImpl) CreateUser(c context.Context, user User) (User, error) {
	filter := bson.M{"$or": []bson.M{{"email": user.Email}}}

	var existingUser User
	err := r.collection.FindOne(c, filter).Decode(&existingUser)
	if err == nil {
		return User{}, errors.New("a user with this email already exists")
	} else if err != mongo.ErrNoDocuments {
		return User{}, err
	}

	result, err := r.collection.InsertOne(c, user)
	if err != nil {
		return User{}, err
	}

	_, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return User{}, errors.New("failed to convert inserted ID to ObjectID")
	}

	return user, nil
}

func (r *repositoryImpl) GetAllUsers(c context.Context) ([]User, error) {
	cursor, err := r.collection.Find(c, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var users []User
	if err := cursor.All(c, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repositoryImpl) GetUserByID(c context.Context, id primitive.ObjectID) (*User, error) {
	filter := bson.M{"_id": id}

	var user User
	err := r.collection.FindOne(c, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repositoryImpl) GetUserByEmail(c context.Context, email string) (*User, error) {
	filter := bson.M{"email": email}

	var user User
	err := r.collection.FindOne(c, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repositoryImpl) AddReportee(c context.Context, userID primitive.ObjectID, reporteeID primitive.ObjectID) error {
	filter := bson.M{"_id": userID}
	update := bson.M{"$addToSet": bson.M{"reportees": reporteeID}}

	_, err := r.collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *repositoryImpl) RemoveReportee(c context.Context, userID primitive.ObjectID, reporteeID primitive.ObjectID) error {
	filter := bson.M{"_id": userID}
	update := bson.M{"$pull": bson.M{"reportees": reporteeID}}

	_, err := r.collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *repositoryImpl) AddReportsTo(c context.Context, userID primitive.ObjectID, reportsToID primitive.ObjectID) error {
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{"reportsTo": reportsToID}}

	_, err := r.collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	return nil
}
