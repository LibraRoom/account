package repository

import (
	"account/features/permissions"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PermissionsModel struct {
	Code     string    `bson:"_id"`
	Name     string    `bson:"name"`
	CreateAt time.Time `bson:"create_at"`
	UpdateAt time.Time `bson:"update_at"`
	DeleteAt time.Time `bson:"delete_at"`
}

type PermissionsQuery struct {
	db         *mongo.Database
	collection string
}

func New(client *mongo.Database, collection string) permissions.Repository {
	return &PermissionsQuery{
		db:         client,
		collection: collection,
	}
}

// GetAllPermissions implements permissions.Repository.
func (pq *PermissionsQuery) GetAllPermissions() ([]permissions.Permissions, error) {
	cursor, err := pq.db.Collection(pq.collection).Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var permission []PermissionsModel
	if err := cursor.All(context.TODO(), &permission); err != nil {
		return nil, err
	}
	var result []permissions.Permissions
	for _, s := range permission {
		result = append(result, permissions.Permissions{
			Code: s.Code,
			Name: s.Name,
		})
	}
	return result, nil
}
