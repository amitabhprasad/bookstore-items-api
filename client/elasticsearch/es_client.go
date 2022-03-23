package elasticsearch

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/amitabhprasad/bookstore-util-go/logger"
	"github.com/olivere/elastic"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	Index(string, interface{}) (*elastic.IndexResponse, error)
	setClient(*elastic.Client)
}
type esClient struct {
	client *elastic.Client
}

func Init() {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}
	log := logger.GetLogger()
	c, err := elastic.NewClient(
		elastic.SetHttpClient(client),
		elastic.SetURL("https://9.30.161.130:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetBasicAuth("elastic", "F0=HQFdS=kOx3Y0nilFT"),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
		// elastic.SetHeaders(http.Header{
		// 	"X-Caller-Id": []string{"..."},
		// }),
	)
	if err != nil {
		fmt.Println("Error during init ", err)
		panic(err)
	}
	Client.setClient(c)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}
func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		BodyJson(doc).
		Type("_doc").
		Do(ctx)

	if err != nil {
		logger.Error(
			fmt.Sprintf("error when trying to index document in elasticsearch %s", index), err)
		return nil, err
	}
	return result, nil
}
