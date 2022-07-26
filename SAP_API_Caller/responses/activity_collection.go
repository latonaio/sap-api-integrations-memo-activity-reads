package responses

type ActivityCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                 string `json:"ObjectID"`
			ETag                     string `json:"ETag"`
			ID                       string `json:"ID"`
			ProcessingTypeCode       string `json:"ProcessingTypeCode"`
			PriorityCode             string `json:"PriorityCode"`
			ReportedDate             string `json:"ReportedDate"`
			ReportedDateTime         string `json:"ReportedDateTime"`
			SubjectName              string `json:"SubjectName"`
			UUID                     string `json:"UUID"`
			TypeCode                 string `json:"TypeCode"`
			LifeCycleStatusCode      string `json:"LifeCycleStatusCode"`
			InitiatorCode            string `json:"InitiatorCode"`
			CreationDate             string `json:"CreationDate"`
			EntityLastChangedOn      string `json:"EntityLastChangedOn"`
		} `json:"results"`
	} `json:"d"`
}
