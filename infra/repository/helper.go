package repository

import (
	"github.com/lawmatsuyama/transactions/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

func bulkInsertModel[T domain.IDSetter](objects []T) (models []mongo.WriteModel) {
	for _, object := range objects {
		insertModel := mongo.NewInsertOneModel()
		object.SetID()
		insertModel.SetDocument(object)
		models = append(models, insertModel)
	}
	return
}
