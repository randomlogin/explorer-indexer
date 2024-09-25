package node

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"strings"

	"net/http"
)

type Client struct {
	httpclient *http.Client
	origin     string
	username   string
	password   string
}

func NewClient(origin, username, password string) *Client {
	return &Client{&http.Client{}, origin, username, password}
}

func (client *Client) do(ctx context.Context, method string, path string, body interface{}, target interface{}) error {
	var reader io.Reader = nil
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reader = bytes.NewBuffer(buf)
	}
	req, err := http.NewRequestWithContext(ctx, method, client.origin+"/"+path, reader)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	req.SetBasicAuth(client.username, client.password)
	// log.Printf("%+v\n", req)
	res, err := client.httpclient.Do(req)
	if err != nil {
		return err
	}
	// log.Printf("%+v\n", res)
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}

func (client *Client) rest(ctx context.Context, method string, path []string, body interface{}, target interface{}) error {
	p := strings.Join(path, "/")
	return client.do(ctx, method, p, body, target)
}

type rpcBody struct {
	Method  string        `json:"method"`
	Params  []interface{} `json:"params,omitempty"`
	JsonRPC string        `json:"jsonrpc"`
	Id      int           `json:"id"`
}

type RpcResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	ID      int             `json:"id"`
	Error   *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (client *Client) Rpc(ctx context.Context, method string, params []interface{}, target interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	b := rpcBody{method, params, "2.0", 1337}
	// log.Printf("%+v", b)
	// t := rpcResult{}
	t := RpcResponse{}
	if err := client.do(ctx, "POST", "", &b, &t); err != nil {
		return err
	}
	if t.Error != nil {
		return fmt.Errorf("spaces node: %v", t.Error.Message)
	}

	return json.Unmarshal(t.Result, target)
}