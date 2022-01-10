package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/olelarssen/provider/models"
	log "github.com/olelarssen/provider/service/logger"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Client struct {
	Ctx        context.Context
	logger     log.Logger
	HTTPClient *http.Client
	mutex      sync.Mutex
}

func NewClient(ctx context.Context, logger log.Logger) *Client {
	return &Client{
		Ctx:    ctx,
		logger: logger,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}
		return fmt.Errorf("status code: %d", res.StatusCode)
	}
	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetError(err error) *models.Error {
	message := fmt.Sprintf("%v", err)
	var code int64 = 500
	e := true
	return &models.Error{Message: &message, Code: &code, Error: &e}
}
