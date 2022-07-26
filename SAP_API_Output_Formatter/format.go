package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-memo-activity-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToActivityCollection(raw []byte, l *logger.Logger) ([]ActivityCollection, error) {
	pm := &responses.ActivityCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ActivityCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	activityCollection := make([]ActivityCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		activityCollection = append(activityCollection, ActivityCollection{
			ObjectID:                     data.ObjectID,
			ETag:                         data.ETag,
			ID:                           data.ID,
			ProcessingTypeCode:           data.ProcessingTypeCode,
			PriorityCode:                 data.PriorityCode,
			ReportedDate:                 data.ReportedDate,
			ReportedDateTime:             data.ReportedDateTime,
			SubjectName:                  data.SubjectName,
			UUID:                         data.UUID,
			TypeCode:                     data.TypeCode,
			LifeCycleStatusCode:          data.LifeCycleStatusCode,
			InitiatorCode:                data.InitiatorCode,
			CreationDate:                 data.CreationDate,
			EntityLastChangedOn:          data.EntityLastChangedOn,
		})
	}

	return activityCollection, nil
}