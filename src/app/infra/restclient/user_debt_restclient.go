package restclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rapicreds-backend/src/app/domain"
	"rapicreds-backend/src/app/infra/apierror"
)

const baseURL = "https://api.bcra.gob.ar/centraldedeudores/v1.0/Deudas/%d"

type IUserDebtRestClient interface {
	GetUserDebt(document uint64) (*domain.UserDebt, *apierror.ApiError)
}

type BaseUserDebtRestClient struct {
}

func NewUserDebtRestClient() IUserDebtRestClient {
	return &BaseUserDebtRestClient{}
}

func (b BaseUserDebtRestClient) GetUserDebt(document uint64) (*domain.UserDebt, *apierror.ApiError) {
	url := fmt.Sprintf(baseURL, document)

	resp, err := http.Get(url)
	if err != nil {
		err := apierror.NewCustomErrorWithStatus("BaseUserDebtRestClient - Error calling http", err, http.StatusInternalServerError)
		apierror.LogCustomError(err)
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println(url)

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := apierror.NewCustomErrorWithStatus("BaseUserDebtRestClient - reading response body", err, http.StatusInternalServerError)
		apierror.LogCustomError(err)
		return nil, err
	}

	var result *domain.UserDebt
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		err := apierror.NewCustomErrorWithStatus("BaseUserDebtRestClient - Error converting JSON", err, http.StatusInternalServerError)
		apierror.LogCustomError(err)
		return nil, err
	}

	return result, nil
}
