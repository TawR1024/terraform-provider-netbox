package netbox

import (
	"log"

	"github.com/netbox-community/go-netbox/netbox/client/tenancy"

	"github.com/netbox-community/go-netbox/netbox/client/dcim"
)

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
