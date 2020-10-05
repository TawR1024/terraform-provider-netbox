package netbox

import (
	"log"
	"strconv"

	"github.com/netbox-community/go-netbox/netbox/client/dcim"

	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func resourceNetboxSite() *schema.Resource {
	return &schema.Resource{
		Create: resourceSiteCreate,
		Read:   resourceSiteRead,
		Update: resourceSiteUpdate,
		Delete: resourceSiteDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Site name",
			},
			"tenant": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Tenant name",
			},
			"slug": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "site name slug",
			},
			"facility": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "tenant  name",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Site description",
			},
		},
	}
}

func resourceSiteCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient)
	netboxSite := new(models.WritableSite)
	params := dcim.NewDcimSitesCreateParams()

	name := d.Get("name").(string)
	netboxSite.Name = &name
	netboxSite.Tenant = c.GetTenantId(swag.String(d.Get("tenant").(string)))
	netboxSite.Slug = swag.String(d.Get("slug").(string))
	netboxSite.Facility = d.Get("facility").(string)
	netboxSite.Description = d.Get("description").(string)
	netboxSite.Tags = []string{}

	params.WithData(netboxSite)
	res, err := c.netboxClient.Dcim.DcimSitesCreate(params, nil)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Print("Site ID is: ", res.Payload.ID)
	d.SetId(strconv.FormatInt(res.Payload.ID, 10))
	return resourceSiteRead(d, m)
}

func resourceSiteRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient).netboxClient
	params := dcim.NewDcimSitesListParams()
	switch {
	case d.Id() != "":
		params.WithIDIn(swag.String(d.Id()))
	case d.Get("name") != nil:
		name := d.Get("name").(string)
		params.WithName(&name)
	}
	res, err := c.Dcim.DcimSitesList(params, nil)
	if err != nil {
		log.Print("[DEBUG] Cant read Vrf info resourceVRFRead() ", err)
	}
	d.Set("name", res.Payload.Results[0].Name)
	d.Set("tenant", res.Payload.Results[0].Tenant.Name)
	d.Set("slug", res.Payload.Results[0].Slug)
	d.Set("description", res.Payload.Results[0].Description)
	d.Set("facility", res.Payload.Results[0].Facility)
	return nil
}
func resourceSiteUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient)
	netboxSite := new(models.WritableSite)

	name := d.Get("name").(string)
	netboxSite.Name = &name
	netboxSite.Tenant = c.GetTenantId(swag.String(d.Get("tenant").(string)))
	netboxSite.Slug = swag.String(d.Get("slug").(string))
	netboxSite.Facility = d.Get("facility").(string)
	netboxSite.Description = d.Get("description").(string)
	netboxSite.Tags = []string{}

	params := dcim.NewDcimSitesUpdateParams()
	vrfID, err := strconv.Atoi(d.Id())
	if err != nil {
		log.Print("string converting failed")
		return err
	}
	params.WithID(int64(vrfID))
	params.WithData(netboxSite)
	_, err = c.netboxClient.Dcim.DcimSitesUpdate(params, nil)
	if err != nil {
		log.Print("[DEBUG] Update site failed\n", err)
		return err
	} else {
		log.Print("Updated...")
	}

	return resourceSiteRead(d, m)
}

func resourceSiteDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient).netboxClient
	params := dcim.NewDcimSitesDeleteParams()
	siteID, err := strconv.Atoi(d.Id())
	if err != nil {
		log.Print("string converting failed")
	}
	params.WithID(int64(siteID))
	_, err = c.Dcim.DcimSitesDelete(params, nil)
	if err != nil {
		log.Print("[DEBUG] Delete site failed\n", err)
	}

	return nil
}
