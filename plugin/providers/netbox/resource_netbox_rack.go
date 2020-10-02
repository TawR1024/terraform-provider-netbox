package netbox

import (
	"log"
	"strconv"

	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
)

func resourceNetboxRack() *schema.Resource {
	return &schema.Resource{
		Create: resourceRackCreate,
		Read:   resourceRackRead,
		Update: resourceRackUpdate,
		Delete: resourceRackDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"site": &schema.Schema{
				Type:     schema.TypeString, // todo: use site name
				Required: true,
			},
			"tenant": &schema.Schema{
				Type:     schema.TypeString, // todo: use tenant name
				Required: true,
			},
			"height": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Rack height in units",
			},
			"width": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Rack width in inches",
			},
			"facility": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Data-center name",
			},
			"desc_units": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Descending units",
			},
			"role": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Rack role project.zone",
			},
		},
	}
}

func resourceRackCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient)
	netboxRack := new(models.WritableRack)
	name := d.Get("name").(string)
	netboxRack.Name = &name
	netboxRack.Site = c.GetSiteID(swag.String(d.Get("site").(string)))
	netboxRack.Tenant = c.GetTenantId(swag.String(d.Get("tenant").(string)))
	netboxRack.UHeight = int64(d.Get("height").(int))
	netboxRack.Width = int64(d.Get("width").(int))
	netboxRack.Tags = []string{}
	facility := d.Get("facility").(string)
	netboxRack.FacilityID = &facility
	netboxRack.DescUnits = d.Get("desc_units").(bool)

	params := dcim.NewDcimRacksCreateParams()
	params.WithData(netboxRack)

	res, err := c.netboxClient.Dcim.DcimRacksCreate(params, nil)
	if err != nil {
		log.Print(err)
		return err
	}

	log.Print("Rack ID is: ", res.Payload.ID)
	d.SetId(strconv.FormatInt(res.Payload.ID, 10))
	return resourceRackRead(d, m)
}

func resourceRackRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient).netboxClient
	params := dcim.NewDcimRacksListParams()
	//todo: Search by name or ID if provided
	switch {
	case d.Id() != "":
		params.WithID(swag.String(d.Id()))
	case d.Get("name") != nil:
		name := d.Get("name").(string)
		params.WithName(&name)
	}
	res, err := c.Dcim.DcimRacksList(params, nil)
	if err != nil {
		log.Print("[DEBUG] Cant read Rack info resourceRackRead() ", err)
	}
	d.Set("name", res.Payload.Results[0].Name)
	d.Set("site", res.Payload.Results[0].Site.Name)
	d.Set("tenant", res.Payload.Results[0].Tenant.Name)
	d.Set("height", res.Payload.Results[0].UHeight)
	d.Set("width", res.Payload.Results[0].Width.Value)
	d.Set("facility", res.Payload.Results[0].FacilityID)
	d.Set("desc_units", res.Payload.Results[0].DescUnits)

	return nil
}
func resourceRackUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient)
	netboxRack := new(models.WritableRack)
	name := d.Get("name").(string)

	netboxRack.Name = &name
	netboxRack.Site = c.GetSiteID(swag.String(d.Get("site").(string)))
	netboxRack.Tenant = c.GetTenantId(swag.String(d.Get("tenant").(string)))
	netboxRack.UHeight = int64(d.Get("height").(int))
	netboxRack.Width = int64(d.Get("width").(int))
	netboxRack.Tags = []string{}
	facility := d.Get("facility").(string)
	netboxRack.FacilityID = &facility
	netboxRack.DescUnits = d.Get("desc_units").(bool)

	params := dcim.NewDcimRacksPartialUpdateParams()
	rackID, err := strconv.Atoi(d.Id())
	if err != nil {
		log.Print("string converting failed")
	}
	params.WithID(int64(rackID))
	params.WithData(netboxRack)
	_, err = c.netboxClient.Dcim.DcimRacksPartialUpdate(params, nil)
	if err != nil {
		log.Print("[DEBUG] Update Rack failed\n", err)
	} else {
		log.Print("Updated...")
	}

	return resourceRackRead(d, m)
}

func resourceRackDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient).netboxClient
	params := dcim.NewDcimRacksDeleteParams()
	rackID, err := strconv.Atoi(d.Id())
	if err != nil {
		log.Print("string converting failed")
	}
	params.WithID(int64(rackID))
	_, err = c.Dcim.DcimRacksDelete(params, nil)
	if err != nil {
		log.Print("[DEBUG] Delete Rack failed\n", err)
	}

	return nil
}
