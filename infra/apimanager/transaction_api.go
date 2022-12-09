package apimanager

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/lawmatsuyama/transactions/domain"
	"github.com/sirupsen/logrus"
)

var Now = time.Now()

type TransactionAPI struct {
	UseCase domain.TransactionUseCase
}

func NewTransactionAPI(useCase domain.TransactionUseCase) TransactionAPI {
	return TransactionAPI{
		UseCase: useCase,
	}
}

func (api TransactionAPI) Save(w http.ResponseWriter, r *http.Request) {
	var trsReq TransactionsSaveRequest
	err := json.NewDecoder(r.Body).Decode(&trsReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	trsResult, err := api.UseCase.Save(ctx, trsReq.UserID, trsReq.ToTransactions(Now))
	var trsResponse []TransactionSaveResponse
	if err != nil {
		trsResponse = FromTransactionValidateResult(trsResult)
	}
	handleResponse(w, r, trsResponse, err)

}

func handleResponse(w http.ResponseWriter, r *http.Request, in any, err error) {
	var errStr string
	statusCode := http.StatusOK
	if err != nil {
		errTr := domain.ErrorTransactionToError(err)
		errStr = errTr.Error()
		statusCode = errTr.Status()
	}

	genRes := GenericResponse{
		Error:  errStr,
		Result: in,
	}

	res, err := json.Marshal(genRes)
	if err != nil {
		logrus.WithError(err).Error("couldnt marshal the response to json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := w.Write(res); err != nil {
		logrus.WithError(err).Error("couldnt send response to writer")
		http.Error(w, domain.ErrUnknow.Error(), http.StatusBadRequest)
	}
}
