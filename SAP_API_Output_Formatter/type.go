package sap_api_output_formatter

type MemoActivity struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	MemoActivityCode  string `json:"memo_activity_code"`
	Deleted       bool   `json:"deleted"`
}

type ActivityCollection struct {
	ObjectID                     string `json:"ObjectID"`
	ETag                 		 string `json:"ETag"`
	ID             				 string `json:"ID"`
	ProcessingTypeCode           string `json:"ProcessingTypeCode"`
	PriorityCode                 string `json:"PriorityCode"`
	ReportedDate                 string `json:"ReportedDate"`
	ReportedDateTime             string `json:"ReportedDateTime"`
	SubjectName                  string `json:"SubjectName"`
	UUID                   		 string `json:"UUID"`
	TypeCode              		 string `json:"TypeCode"`
	LifeCycleStatusCode          string `json:"LifeCycleStatusCode"`
	InitiatorCode                string `json:"InitiatorCode"`
	CreationDate            	 string `json:"CreationDate"`
	EntityLastChangedOn        	 string `json:"EntityLastChangedOn"`
}
