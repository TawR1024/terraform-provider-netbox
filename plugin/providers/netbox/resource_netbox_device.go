package netbox

import (
	"log"
	"strconv"

	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
)

func resourceNetboxDevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeviceCreate,
		Read:   resourceDeviceRead,
		Update: resourceDeviceUpdate,
		Delete: resourceDeviceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Device name",
			},
			"site": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "netbox site",
			},
			"tenant": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "netbox tenant",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "device type",
			},
			"rack": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "rack name",
			},
			"position": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "device position in rack",
			},
			"face": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "front or rear",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "device status",
			},
			"serial": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "device serial number",
			},
		},
	}
}

func resourceDeviceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient)
	netboxDevice := new(models.WritableDeviceWithConfigContext)
	name := d.Get("name").(string)
	netboxDevice.Name = &name
	netboxDevice.Site = c.GetSiteID(swag.String(d.Get("site").(string)))
	netboxDevice.Tenant = c.GetTenantId(swag.String(d.Get("tenant").(string)))
	netboxDevice.DeviceType = c.GetDeviceTypeId(swag.String(d.Get("type").(string)))
	netboxDevice.DeviceRole = c.GetDeviceRoleId(swag.String(d.Get("role").(string)))
	netboxDevice.Rack = c.GetRackId(swag.String(d.Get("rack").(string)), netboxDevice.Site)
	netboxDevice.Position = swag.Int64(int64(d.Get("position").(int)))
	netboxDevice.Face = swag.Int64(int64(d.Get("face").(int)))
	netboxDevice.Status = status[d.Get("status").(string)]
	netboxDevice.Serial = d.Get("setial").(string)
	netboxDevice.Tags = []string{}

	params := dcim.NewDcimDevicesCreateParams()
	params.WithData(netboxDevice)

	res, err := c.netboxClient.Dcim.DcimDevicesCreate(params, nil)
	if err != nil {
		log.Print(err)
		return err
	}

	log.Print("Device ID is: ", res.Payload.ID)
	d.SetId(strconv.FormatInt(res.Payload.ID, 10))
	return resourceRackRead(d, m)
}

func resourceDeviceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDeviceUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceDeviceDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
