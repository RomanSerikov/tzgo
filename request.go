package tzgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// TezosNode -
type TezosNode struct {
	baseURL          string
	client           http.Client
	outOfSyncTimeout time.Duration
}

// NewTezosNode -
func NewTezosNode(baseURL string, timeout time.Duration) *TezosNode {
	return &TezosNode{
		baseURL:          baseURL,
		outOfSyncTimeout: -3 * time.Minute,
		client: http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

// SetTimeout - default is 10 sec
func (t *TezosNode) SetTimeout(timeout time.Duration) {
	t.client.Timeout = timeout
}

func (t *TezosNode) get(endpoint string, response interface{}) error {
	uri := fmt.Sprintf("%s%s", t.baseURL, endpoint)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return err
	}

	resp, err := t.client.Do(req)
	if err != nil {
		return fmt.Errorf("[GET]: %s error: %v", req.URL.String(), err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return fmt.Errorf("json.Decode error: %v", err)
	}
	return nil
}

func (t *TezosNode) isOutOfSync(header Header) bool {
	ts := time.Now().UTC().Add(t.outOfSyncTimeout)
	return header.Timestamp.UTC().Before(ts)
}
