package utility

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	entity "milestone2/entities"
)

func CreateInvoice(product entity.InvoiceRequest, user entity.User) (*entity.Invoice, string, error) {
	apiKey := os.Getenv("XENDIT_API_KEY")
	apiUrl := "https://api.xendit.co/v2/invoices"

	randomString := GenerateRandomString(10)
	externalId := fmt.Sprintf("HCKTV8-%s", randomString)

	bodyRequest := map[string]interface{}{
		"external_id":      externalId,
		"amount":           product.Price,
		"description":      product.Description,
		"invoice_duration": 86400,
		"customer": map[string]interface{}{
			"given_names": user.FullName,
			"email":       user.Email,
		},
		"currency": "IDR",
		"items": []interface{}{
			map[string]interface{}{
				"name":     product.Name,
				"quantity": 1,
				"price":    product.Price,
			},
		},
	}

	reqBody, err := json.Marshal(bodyRequest)
	if err != nil {
		return nil, "", err
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, "", err
	}

	request.SetBasicAuth(apiKey, "")
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, "", err
	}

	defer response.Body.Close()

	var resInvoice entity.Invoice
	if err := json.NewDecoder(response.Body).Decode(&resInvoice); err != nil {
		return nil, "", err
	}

	return &resInvoice, externalId, nil

}
