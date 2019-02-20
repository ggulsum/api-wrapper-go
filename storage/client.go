package storage

import (
	"api-wrapper-go"
	"net/http"
)

// Client is used to invoke /Entity and entity-related APIs.
type Client struct {
	B   uiza.Backend
	Key string
}

const (
	baseURL = "api/public/v3/media/storage"
)

// Add Storage API
func Add(params *uiza.StorageAddParams) (*uiza.StorageIdData, error) {
	return getC().Add(params)
}

// Add Storage API
func (c Client) Add(params *uiza.StorageAddParams) (*uiza.StorageIdData, error) {
	storageAdd := &uiza.StorageIdData{}

	err := c.B.Call(http.MethodPost, baseURL, c.Key, params, storageAdd)
	return storageAdd, err
}

// Get Backend Client
func getC() Client {
	return Client{uiza.GetBackend(uiza.APIBackend), uiza.Key}
}