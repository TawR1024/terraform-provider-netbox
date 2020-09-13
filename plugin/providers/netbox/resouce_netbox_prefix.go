package netbox

import (
	"log"
	"strconv"

	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
)

func resourceNetboxPrefix() *schema.Resource {
	return &schema.Resource{
		Create: resourcePrefixCreate,
		Read:   resourcePrefixRead,
		Update: resourcePrefixUpdate,
		Delete: resourcePrefixDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"prefix": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Prefix with netmask",
			},
			"site": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "netbox site",
			},
			"tenant": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "netbox tenant",
			},
			"vrf": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Global",
				Description: "vrf",
			},
		},
	}
}

func resourcePrefixCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient)
	netboxPrefix := new(models.WritablePrefix)
	prefix := d.Get("prefix").(string)
	netboxPrefix.Prefix = &prefix //todo: validate prefix string
	netboxPrefix.Site = c.GetSiteID(swag.String(d.Get("site").(string)))
	netboxPrefix.Tenant = c.GetTenantId(swag.String(d.Get("tenant").(string)))
	netboxPrefix.Vrf = c.GetVRFID(swag.String(d.Get("vrf").(string)))
	params := ipam.NewIPAMPrefixesCreateParams()
	params.WithData(netboxPrefix)
	netboxPrefix.Tags = []string{}
	res, err := c.netboxClient.IPAM.IPAMPrefixesCreate(params, nil)
	if err != nil {
		log.Print(err)
		return err
	}

	log.Print("Prefix ID is: ", res.Payload.ID)
	d.SetId(strconv.FormatInt(res.Payload.ID, 10))
	return resourcePrefixRead(d, m)
}

func resourcePrefixRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient).netboxClient
	params := ipam.NewIPAMPrefixesListParams()
	switch {
	case d.Id() != "":
		params.WithIDIn(swag.String(d.Id()))
	case d.Get("prefix") != nil:
		prefix := d.Get("prefix").(string)
		params.WithPrefix(swag.String(prefix))
	}
	res, err := c.IPAM.IPAMPrefixesList(params, nil)
	if err != nil {
		log.Print("[DEBUG] Cant read Rack info resourceDeviceRead() ", err)
	}
	d.Set("prefix", res.Payload.Results[0].Prefix)
	d.Set("site", res.Payload.Results[0].Site.Name)
	d.Set("tenant", res.Payload.Results[0].Tenant.Name)
	d.Set("vrf", res.Payload.Results[0].Vrf.Name)

	return nil
}

func resourcePrefixUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient)
	netboxPrefix := new(models.WritablePrefix)
	prefix := d.Get("prefix").(string)

	netboxPrefix.Prefix = &prefix
	netboxPrefix.Site = c.GetSiteID(swag.String(d.Get("site").(string)))
	netboxPrefix.Tenant = c.GetTenantId(swag.String(d.Get("tenant").(string)))
	netboxPrefix.Vrf = c.GetVRFID(swag.String(d.Get("vrf").(string)))

	netboxPrefix.Tags = []string{}

	params := ipam.NewIPAMPrefixesUpdateParams()
	prefixID, err := strconv.Atoi(d.Id())
	if err != nil {
		log.Print("string converting failed")
	}
	params.WithID(int64(prefixID))
	params.WithData(netboxPrefix)
	_, err = c.netboxClient.IPAM.IPAMPrefixesUpdate(params, nil)
	if err != nil {
		log.Print("[DEBUG] Update Rack failed\n", err)
	} else {
		log.Print("Updated...")
	}

	return resourcePrefixRead(d, m)
}
func resourcePrefixDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient).netboxClient
	params := ipam.NewIPAMPrefixesDeleteParams()
	prefixID, err := strconv.Atoi(d.Id())
	if err != nil {
		log.Print("string converting failed")
	}
	params.WithID(int64(prefixID))
	_, err = c.IPAM.IPAMPrefixesDelete(params, nil)
	if err != nil {
		log.Print("[DEBUG] Delete Prefix failed\n", err)
	}

	return nil
}
