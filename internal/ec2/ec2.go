package ec2

import (
	"github.com/nronix/cq-source-prismacloud/client"
	"github.com/paloaltonetworks/prisma-cloud-go/rql/search"
	"github.com/paloaltonetworks/prisma-cloud-go/timerange"
	"time"
)

type Ec2Inventory struct {
	Tags []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"tags"`
	State struct {
		Code int    `json:"code"`
		Name string `json:"name"`
	} `json:"state"`
	VpcId     string        `json:"vpcId"`
	ImageId   string        `json:"imageId"`
	KeyName   string        `json:"keyName"`
	Licenses  []interface{} `json:"licenses"`
	SubnetId  string        `json:"subnetId"`
	Placement struct {
		Tenancy          string `json:"tenancy"`
		GroupName        string `json:"groupName"`
		AvailabilityZone string `json:"availabilityZone"`
	} `json:"placement"`
	CpuOptions struct {
		CoreCount      int `json:"coreCount"`
		ThreadsPerCore int `json:"threadsPerCore"`
	} `json:"cpuOptions"`
	EnaSupport bool      `json:"enaSupport"`
	Hypervisor string    `json:"hypervisor"`
	InstanceId string    `json:"instanceId"`
	LaunchTime time.Time `json:"launchTime"`
	Monitoring struct {
		State string `json:"state"`
	} `json:"monitoring"`
	ClientToken    string        `json:"clientToken"`
	Architecture   string        `json:"architecture"`
	EbsOptimized   bool          `json:"ebsOptimized"`
	InstanceType   string        `json:"instanceType"`
	ProductCodes   []interface{} `json:"productCodes"`
	StatusEvents   []interface{} `json:"statusEvents"`
	PublicDnsName  string        `json:"publicDnsName"`
	AmiLaunchIndex int           `json:"amiLaunchIndex"`
	EnclaveOptions struct {
		Enabled bool `json:"enabled"`
	} `json:"enclaveOptions"`
	PrivateDnsName string `json:"privateDnsName"`
	RootDeviceName string `json:"rootDeviceName"`
	RootDeviceType string `json:"rootDeviceType"`
	SecurityGroups []struct {
		GroupId   string `json:"groupId"`
		GroupName string `json:"groupName"`
	} `json:"securityGroups"`
	MetadataOptions struct {
		State                   string `json:"state"`
		HttpTokens              string `json:"httpTokens"`
		HttpEndpoint            string `json:"httpEndpoint"`
		HttpPutResponseHopLimit int    `json:"httpPutResponseHopLimit"`
	} `json:"metadataOptions"`
	SourceDestCheck   bool   `json:"sourceDestCheck"`
	PrivateIpAddress  string `json:"privateIpAddress"`
	NetworkInterfaces []struct {
		VpcId  string `json:"vpcId"`
		Groups []struct {
			GroupId   string `json:"groupId"`
			GroupName string `json:"groupName"`
		} `json:"groups"`
		Status     string `json:"status"`
		OwnerId    string `json:"ownerId"`
		SubnetId   string `json:"subnetId"`
		Attachment struct {
			Status              string    `json:"status"`
			AttachTime          time.Time `json:"attachTime"`
			DeviceIndex         int       `json:"deviceIndex"`
			AttachmentId        string    `json:"attachmentId"`
			NetworkCardIndex    int       `json:"networkCardIndex"`
			DeleteOnTermination bool      `json:"deleteOnTermination"`
		} `json:"attachment"`
		MacAddress         string        `json:"macAddress"`
		Description        string        `json:"description"`
		InterfaceType      string        `json:"interfaceType"`
		Ipv6Addresses      []interface{} `json:"ipv6Addresses"`
		PrivateDnsName     string        `json:"privateDnsName"`
		SourceDestCheck    bool          `json:"sourceDestCheck"`
		PrivateIpAddress   string        `json:"privateIpAddress"`
		NetworkInterfaceId string        `json:"networkInterfaceId"`
		PrivateIpAddresses []struct {
			Primary          bool   `json:"primary"`
			PrivateDnsName   string `json:"privateDnsName"`
			PrivateIpAddress string `json:"privateIpAddress"`
		} `json:"privateIpAddresses"`
	} `json:"networkInterfaces"`
	HibernationOptions struct {
		Configured bool `json:"configured"`
	} `json:"hibernationOptions"`
	IamInstanceProfile struct {
		Id  string `json:"id"`
		Arn string `json:"arn"`
	} `json:"iamInstanceProfile"`
	VirtualizationType  string `json:"virtualizationType"`
	BlockDeviceMappings []struct {
		Ebs struct {
			Status              string    `json:"status"`
			VolumeId            string    `json:"volumeId"`
			AttachTime          time.Time `json:"attachTime"`
			DeleteOnTermination bool      `json:"deleteOnTermination"`
		} `json:"ebs"`
		DeviceName string `json:"deviceName"`
	} `json:"blockDeviceMappings"`
	DisableApiTermination            bool          `json:"disableApiTermination"`
	StateTransitionReason            string        `json:"stateTransitionReason"`
	ElasticGpuAssociations           []interface{} `json:"elasticGpuAssociations"`
	CapacityReservationSpecification struct {
		CapacityReservationPreference string `json:"capacityReservationPreference"`
	} `json:"capacityReservationSpecification"`
	ElasticInferenceAcceleratorAssociations []interface{} `json:"elasticInferenceAcceleratorAssociations"`
}

type ConfigRequest struct {
	Id               string              `json:"id,omitempty"`
	TimeRange        timerange.TimeRange `json:"timeRange"`
	Query            string              `json:"query"`
	Limit            int                 `json:"limit,omitempty"`
	WithResourceJson bool                `json:"withResourceJson,omitempty"`
	SkipResult       bool                `json:"skipResult,omitempty"`
}

type ConfigResponse struct {
	GroupBy     []string   `json:"groupBy"`
	Id          string     `json:"id"`
	CloudType   string     `json:"cloudType"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	SearchType  string     `json:"searchType"`
	Data        ConfigData `json:"data"`
}

type ConfigData struct {
	Items []ConfigItem `json:"items"`
}

type ConfigItem struct {
	StateId                  string       `json:"stateId"`
	Name                     string       `json:"name"`
	Url                      string       `json:"url"`
	AccountId                string       `json:"accountId"`
	AccountName              string       `json:"accountName"`
	AccountGroupName         string       `json:"accountGroupName"`
	CloudType                string       `json:"cloudType"`
	RegionId                 string       `json:"regionId"`
	RegionName               string       `json:"regionName"`
	Service                  string       `json:"service"`
	ResourceType             string       `json:"resourceType"`
	InsertTs                 int          `json:"insertTs"`
	Deleted                  bool         `json:"deleted"`
	VpcId                    string       `json:"vpcId"`
	VpnName                  string       `json:"vpcName"`
	RiskGrade                string       `json:"riskGrade"`
	HasNetwork               bool         `json:"hasNetwork"`
	HasAlert                 bool         `json:"hasAlert"`
	HasExternalFinding       bool         `json:"hasExternalFinding"`
	HasExternalIntegration   bool         `json:"hasExternalIntegration"`
	AllowDrillDown           bool         `json:"allowDrillDown"`
	HasExtFindingRiskFactors bool         `json:"hasExtFindingRiskFactors"`
	Data                     Ec2Inventory `json:"data"`
}

type Exception struct {
	MessageCode string `json:"messageCode"`
}

func GetEc2Inventry(client *client.Client) ([]Ec2Inventory, error) {
	var query = ""
	if client.Account.INV_QUERY_STRING != "" {
		query = client.Account.INV_QUERY_STRING
	} else {
		query = "config from cloud.resource where api.name = 'aws-ec2-describe-instances'"
	}

	var c_req = search.ConfigRequest{
		TimeRange: timerange.TimeRange{
			Type: timerange.TypeRelative,
			Value: timerange.Relative{
				Amount: 1,
				Unit:   timerange.Day,
			},
		},
		WithResourceJson: true,
		Query:            query,
	}

	// Sanity check the time range.
	if err := c_req.TimeRange.SetType(); err != nil {
		return nil, err
	}
	var resp ConfigResponse
	var BaseSuffix = []string{"search"}
	var ConfigSuffix = []string{"config"}
	path := make([]string, 0, len(BaseSuffix)+len(ConfigSuffix))
	path = append(path, BaseSuffix...)
	path = append(path, ConfigSuffix...)
	_, err := client.Account.PsClient.Communicate("POST", path, nil, c_req, &resp)

	if err != nil {
		return nil, err
	}

	comics := []Ec2Inventory{}
	for _, elm := range resp.Data.Items {

		comics = append(comics, elm.Data)
	}

	return comics, nil
}
