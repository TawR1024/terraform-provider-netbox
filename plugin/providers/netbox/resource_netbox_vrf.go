package netbox

import (
	"log"
	"strconv"

	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func resourceNetboxVRF() *schema.Resource {
	return &schema.Resource{
		Create: resourceVRFCreate,
		Read:   resourceVRFRead,
		Update: resourceVRFUpdate,
		Delete: resourceVRFDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "VRF name",
			},
			"tenant": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "tenant  name",
			},
		},
	}
}

func resourceVRFCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient)
	netboxVRF := new(models.WritableVRF)
	params := ipam.NewIPAMVrfsCreateParams()
	name := d.Get("name").(string)
	netboxVRF.Name = &name
	netboxVRF.Tenant = c.GetTenantId(swag.String(d.Get("tenant").(string)))

	params.WithData(netboxVRF)

	res, err := c.netboxClient.IPAM.IPAMVrfsCreate(params, nil)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Print("Interface ID is: ", res.Payload.ID)
	d.SetId(strconv.FormatInt(res.Payload.ID, 10))
	return resourceVMRead(d, m)
}

func resourceVRFRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient).netboxClient
	params := ipam.NewIPAMVrfsListParams()
	switch {
	case d.Id() != "":
		params.WithName(swag.String(d.Id()))
	case d.Get("name") != nil:
		name := d.Get("name").(string)
		params.WithName(&name)
	}
	res, err := c.IPAM.IPAMVrfsList(params, nil)
	if err != nil {
		log.Print("[DEBUG] Cant read Rack info resourceDeviceRead() ", err)
	}
	d.Set("name", res.Payload.Results[0].Name)
	d.Set("tenant", res.Payload.Results[0].Tenant.Name)
	return nil
}
func resourceVRFUpdate(data *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceVRFDelete(data *schema.ResourceData, m interface{}) error {
	return nil
}
