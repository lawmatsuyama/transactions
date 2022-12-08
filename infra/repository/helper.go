package repository

import "go.mongodb.org/mongo-driver/mongo"

func bulkInsertModel[T any](objects []T) (models []mongo.WriteModel) {
	for object := range objects {
		insertModel := mongo.NewInsertOneModel()
		insertModel.SetDocument(object)
		models = append(models, insertModel)
	}
	return
}
