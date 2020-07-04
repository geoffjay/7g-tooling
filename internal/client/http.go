package client

//import (
//	"net/http"
//	"time"
//
//	gcontext "github.com/geoffjay/7g-tooling/internal/context"
//
//	"github.com/gojektech/heimdall/v6/httpclient"
//)
//
//type SevenGeese struct {
//	client http.Client
//	config *gcontext.Config
//}
//
//func NewSevenGeese(config *gcontext.Config) *SevenGeese {
//	timeout := 1000 * time.Millisecond
//	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
//	httpclient.NewClient(httpclient.WithHTTPClient(&SevenGeese{client: http.DefaultClient}))
//	return &SevenGeese{
//		client: client,
//		config: config,
//	}
//}
//
//func (c *SevenGeese) Do(request *http.Request) (*http.Response, error) {
//	request.SetBasicAuth("username", "passwd")
//	return c.client.Do(request)
//}
