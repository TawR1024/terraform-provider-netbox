package netbox

import (
	"log"
	"strings"

	"github.com/go-openapi/swag"

	"github.com/netbox-community/go-netbox/netbox/client/tenancy"

	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/netbox/client/virtualization"
	//log "github.com/sirupsen/logrus"
)

var status = map[string]int64{
	"Active":  1,
	"Planned": 2,
	"Staged":  3,
}

var protocol = map[string]int64{
	"TCP": 6,
	"UDP": 17,
}

func (c *ProviderNetboxClient) GetSiteID(siteName *string) *int64 {
	params := dcim.NewDcimSitesListParams()
	siteSlug := siteSlug(*siteName)
	params.WithSlug(swag.String(siteSlug))
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
		log.Print("[DEBUG] Cant get cluster id ", err)
	}
	return &res.Payload.Results[0].ID
}

func (c *ProviderNetboxClient) GetDeviceTypeId(deviceTypeName *string) *int64 {
	params := dcim.NewDcimDeviceTypesListParams()
	slug := toslug2(*deviceTypeName)
	params.WithSlug(&slug)
	res, err := c.netboxClient.Dcim.DcimDeviceTypesList(params, nil)
	if err != nil {
		log.Print("[DEBUG] Cant get device type id ", err)
	}
	log.Printf("typeID: %d", res.Payload.Results[0].ID)
	return &res.Payload.Results[0].ID

}

func (c *ProviderNetboxClient) GetDeviceRoleId(deviceRoleName *string) *int64 {
	params := dcim.NewDcimDeviceRolesListParams()
	params.WithName(deviceRoleName)
	res, err := c.netboxClient.Dcim.DcimDeviceRolesList(params, nil)
	if err != nil {
		log.Print("[DEBUG] Cant get device role id ", err)
	}

	return &res.Payload.Results[0].ID
}

func (c *ProviderNetboxClient) GetRackId(rackName *string, site *string) *int64 {
	params := dcim.NewDcimRacksListParams()
	params.WithName(rackName)
	siteSlug := siteSlug(*site)
	params.WithSite(&siteSlug)
	res, err := c.netboxClient.Dcim.DcimRacksList(params, nil)
	if err != nil {
		log.Print("[DEBUG] Cant Get Rack ID", err)
	}
	return &res.Payload.Results[0].ID
}

func toslug(str string) string {
	return strings.ToLower(strings.ReplaceAll(str, ".", "-"))
}

func toslug2(str string) string {
	return strings.ToLower(strings.ReplaceAll(str, "-", "_"))
}

func siteSlug(str string) string {
	return strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(str, ".", "_"), "-", "_"))
}

func (c *ProviderNetboxClient) GetInterfaceID(deviceName, interfaceName *string) (interfaceId *int64) {
	params := dcim.NewDcimInterfacesListParams()
	params.WithDevice(deviceName)
	params.WithName(interfaceName)

	res, err := c.netboxClient.Dcim.DcimInterfacesList(params, nil)
	if err != nil {
		log.Print("Cant Get Inteface ID ", err)
	}

	return &res.Payload.Results[0].ID
}

func (c *ProviderNetboxClient) GetVRFID(vrfName *string) (vrfID *int64) {
	params := ipam.NewIPAMVrfsListParams()
	params.WithName(vrfName)
	res, err := c.netboxClient.IPAM.IPAMVrfsList(params, nil)
	if err != nil {
		log.Print("Can't find VRF", err)
	}
	return &res.Payload.Results[0].ID
}

func (c *ProviderNetboxClient) GetRackRoleID(roleName *string) (roleId *int64) {
	params := dcim.NewDcimRackRolesListParams()
	params.WithName(roleName)
	res, err := c.netboxClient.Dcim.DcimRackRolesList(params, nil)
	if err != nil {
		log.Print("Can't find rack role", err)
	}
	return &res.Payload.Results[0].ID
}
