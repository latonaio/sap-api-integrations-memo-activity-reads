package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-memo-activity-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetMemoActivity(iD, subjectName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ActivityCollectionByID":
			func() {
				c.ActivityCollectionByID(iD)
				wg.Done()
			}()
		case "ActivityCollectionBySubjectName":
			func() {
				c.ActivityCollectionBySubjectName(subjectName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) ActivityCollectionByID(iD string) {
	activityCollectionData, err := c.callMemoActivitySrvAPIRequirementActivityCollection1("ActivityCollection", iD)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(activityCollectionData)

}

func (c *SAPAPICaller) callMemoActivitySrvAPIRequirementActivityCollection1(api, iD string) ([]sap_api_output_formatter.ActivityCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithActivityCollectionByID(req, iD)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToActivityCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ActivityCollectionBySubjectName(subjectName string) {
	activityCollectionData, err := c.callMemoActivitySrvAPIRequirementActivityCollection2("ActivityCollection", subjectName)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(activityCollectionData)

}

func (c *SAPAPICaller) callMemoActivitySrvAPIRequirementActivityCollection2(api, subjectName string) ([]sap_api_output_formatter.ActivityCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithActivityCollectionBySubjectName(req, subjectName)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToActivityCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithActivityCollectionByID(req *http.Request, iD string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ID eq '%s'", iD))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithActivityCollectionBySubjectName(req *http.Request, subjectName string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("substringof('%s', SubjectName)", subjectName))
	req.URL.RawQuery = params.Encode()
}