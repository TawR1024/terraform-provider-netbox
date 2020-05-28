package netbox

import (
	"log"

	"github.com/netbox-community/go-netbox/netbox/client/tenancy"

	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/netbox-community/go-netbox/netbox/client/virtualization"
)

var status = map[string]int64{
	"Active":  1,
	"Planned": 2,
}

func (c *ProviderNetboxClient) GetSiteID(siteName *string) *int64 {
	params := dcim.NewDcimSitesListParams()
	params.WithName(siteName)
	log.Print("[DEBUG] Cant get site id ", *siteName)
	res, err := c.netboxClient.Dcim.DcimSitesList(params, nil)
	if err != nil {
		log.Print("[DEBUG] Cant get site id ", err)
	}
	return &res.Payload.Results[0].ID
}

func (c *ProviderNetboxClient) GetTenantId(tenantName *string) *int64 {
	params := tenancy.NewTenancyTenantsListParams()
	params.WithName(tenantName)
	res, err := c.netboxClient.Tenancy.TenancyTenantsList(params, nil)
	if err != nil {
		log.Print("[DEBUG] Cant get tenant id ", err)
	}
	return &res.Payload.Results[0].ID
}

func (c *ProviderNetboxClient) GetClusterID(clusterName *string) *int64 {
	params := virtualization.NewVirtualizationClustersListParams()
	params.WithName(clusterName)
	res, err := c.netboxClient.Virtualization.VirtualizationClustersList(params, nil)
	if err != nil {
		log.Print("[DEBUG] Cant get tenant id ", err)
	}
	return &res.Payload.Results[0].ID
}

func (c *ProviderNetboxClient) GetDeviceTypeId(deviceTypeName *string) *int64 {
	return nil
}

func (c *ProviderNetboxClient) GetDeviceRoleId(deviceRoleName *string) *int64 {
	return nil
}

func (c *ProviderNetboxClient) GetRackId(rackName *string, siteID *int64) *int64 {
	return nil
}
