package alerts

import (
	"github.com/nronix/cq-source-prismacloud/client"
	"github.com/paloaltonetworks/prisma-cloud-go/alert"
	"github.com/paloaltonetworks/prisma-cloud-go/timerange"
)

type Exception struct {
	MessageCode string `json:"messageCode"`
}

type Alerts struct {
	Id                 string                   `json:"id"`
	Status             string                   `json:"status"`
	FirstSeen          int                      `json:"firstSeen"`
	LastSeen           int                      `json:"lastSeen"`
	AlertTime          int                      `json:"alertTime"`
	EventOccurred      int                      `json:"eventOccurred"`
	TriggeredBy        string                   `json:"triggeredBy"`
	AlertCount         int                      `json:"alertCount"`
	PolicyId           string                   `json:"policyId"`
	AccountId          string                   `json:"accountId"`
	ResourceId         string                   `json:"resourceId"`
	History            []alert.History          `json:"history"`
	Policy             alert.Policy             `json:"policy"`
	Risk               alert.RiskDetail         `json:"riskDetail"`
	Resource           alert.Resource           `json:"resource"`
	InvestigateOptions alert.InvestigateOptions `json:"investigateOptions"`
}

func GetAlerts(client *client.Client) ([]Alerts, error) {

	var filters []alert.Filter
	for _, acc := range client.Account.ALERT_POLICIES {
		filters = append(filters, alert.Filter{Name: "policy.name", Operator: "=", Value: acc})
	}
	filters = append(filters, alert.Filter{Name: "alert.status",
		Operator: "=",
		Value:    "open"})
	filters = append(filters, alert.Filter{Name: "timeRange.type",
		Operator: "=",
		Value:    "ALERT_OPENED"})

	var a_req = alert.Request{
		TimeRange: timerange.TimeRange{
			Type:  timerange.TypeToNow,
			Value: timerange.Epoch,
		},
		Filters: filters,
	}

	// Sanity check the time range.
	if err := a_req.TimeRange.SetType(); err != nil {
		return nil, err
	}
	var resp alert.Response
	var BaseSuffix = []string{"v2"}
	var ConfigSuffix = []string{"alert"}
	path := make([]string, 0, len(BaseSuffix)+len(ConfigSuffix))
	path = append(path, BaseSuffix...)
	path = append(path, ConfigSuffix...)
	_, err := client.Account.PsClient.Communicate("POST", path, nil, a_req, &resp)

	if err != nil {
		return nil, err
	}

	comics := []Alerts{}
	for _, elm := range resp.Data {

		comics = append(comics, Alerts{
			Id:            elm.Id,
			FirstSeen:     elm.FirstSeen,
			Status:        elm.Status,
			LastSeen:      elm.FirstSeen,
			AlertTime:     elm.AlertTime,
			EventOccurred: elm.EventOccurred,
			TriggeredBy:   elm.TriggeredBy,
			AlertCount:    elm.AlertCount,
			PolicyId:      elm.Policy.Id,
			AccountId:     elm.Resource.AccountId,
			ResourceId:    elm.Resource.Name,
			Policy:        elm.Policy,
			Resource:      elm.Resource,
		})
	}

	return comics, nil
}
