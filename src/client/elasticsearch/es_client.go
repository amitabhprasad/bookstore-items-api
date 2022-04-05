package elasticsearch

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/amitabhprasad/bookstore-util-go/logger"
	"github.com/olivere/elastic"
)

var (
	Client esClientInterface = &esClient{}
)

var (
	es_host     = os.Getenv("es_host")
	es_password = os.Getenv("es_password")
)

type esClientInterface interface {
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	setClient(*elastic.Client)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
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
	logger.Info("Connecting to url " + es_host + "using password " + es_password)
	c, err := elastic.NewClient(
		elastic.SetHttpClient(client),
		elastic.SetURL(es_host),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetBasicAuth("elastic", es_password), //"F0=HQFdS=kOx3Y0nilFT"),
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
func (c *esClient) Index(index string, itemType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		BodyJson(doc).
		Type(itemType).
		Do(ctx)

	if err != nil {
		logger.Error(
			fmt.Sprintf("error when trying to index document in elasticsearch %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Get(index string, docType string, itemID string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().
		Index(index).
		Type(docType).
		Id(itemID).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to get id %s", itemID), err)
		return nil, err
	}
	if !result.Found {
		return nil, nil
	}
	return result, nil
}

func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	if err := c.client.Search(index).Query(query).Validate(); err != nil {
		fmt.Println("Error ", err)
		return nil, nil
	}
	result, err := c.client.Search(index).Query(query).RestTotalHitsAsInt(true).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to search document in index %s ", index), err)
		return nil, err
	}
	return result, err
}
