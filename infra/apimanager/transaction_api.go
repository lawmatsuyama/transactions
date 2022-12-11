package apimanager

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/lawmatsuyama/transactions/domain"
	"github.com/sirupsen/logrus"
)

// TransactionAPI represents an API for transaction
type TransactionAPI struct {
	UseCase domain.TransactionUseCase
}

// NewTransactionAPI returns a new TransactionAPI
func NewTransactionAPI(useCase domain.TransactionUseCase) TransactionAPI {
	return TransactionAPI{
		UseCase: useCase,
	}
}

// Save godoc
//
//	@Summary		API to save transactions in the application.
//	@Description	Receives transactions data, registed it in application and finish notifying other applications.
//	@Tags			transaction
//	@Accept			json
//	@Produce		json
//	@Param			transactions_save_request			body		TransactionsSaveRequest								true	"Transactions Save Request"
//	@Success		200				{object}	apimanager.GenericResponse[string]
//	@Failure		400				{object}	apimanager.GenericResponse[[]TransactionSaveResponse]
//	@Failure		404				{object}	apimanager.GenericResponse[[]TransactionSaveResponse]
//	@Router			/v1/save [post]
func (api TransactionAPI) Save(w http.ResponseWriter, r *http.Request) {
	var trsReq TransactionsSaveRequest
	err := json.NewDecoder(r.Body).Decode(&trsReq)
	if err != nil {
		handleResponse[*string](w, r, nil, domain.ErrInvalidTransaction)
		return
	}

	ctx := context.Background()
	trsResult, err := api.UseCase.Save(ctx, trsReq.UserID, trsReq.ToTransactions(domain.Now()))
	if err != nil {
		handleResponse(w, r, FromTransactionSaveResult(trsResult), err)
		return
	}

	handleResponse(w, r, "Save transactions successfully", err)
}

// Get godoc
//
//	@Summary		API to get transactions in the application.
//	@Description	List transactions by giving filter
//	@Tags			transaction
//	@Accept			json
//	@Produce		json
//	@Param			transactions_get_request			body		TransactionsGetRequest								true	"Transactions Get Request"
//	@Success		200				{object}	apimanager.GenericResponse[domain.TransactionsPaging]
//	@Failure		400				{object}	apimanager.GenericResponse[string]
//	@Failure		404				{object}	apimanager.GenericResponse[string]
//	@Router			/v1/get [post]
func (api TransactionAPI) Get(w http.ResponseWriter, r *http.Request) {
	var trsReq TransactionsGetRequest
	err := json.NewDecoder(r.Body).Decode(&trsReq)
	if err != nil {
		handleResponse[*string](w, r, nil, domain.ErrInvalidTransaction)
		return
	}

	ctx := context.Background()
	trsPage, err := api.UseCase.Get(ctx, trsReq.ToTransactionsFilter())
	if err != nil {
		handleResponse[*string](w, r, nil, err)
		return
	}

	response := FromTransactionPaging(trsPage)
	handleResponse(w, r, response, err)
}

func handleResponse[T any](w http.ResponseWriter, r *http.Request, in T, err error) {
	var errStr string
	statusCode := http.StatusOK
	if err != nil {
		errTr := domain.ErrorToErrorTransaction(err)
		errStr = errTr.Error()
		statusCode = errTr.Status()
	}

	genRes := GenericResponse[T]{
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
