package netbox

import (
	"log"
	"strconv"

	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func resourceNetboxService() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCreate,
		Read:   resourceServiceRead,
		Update: resourceServiceUpdate,
		Delete: resourceServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "service port",
			},
			"protocol": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "TCP or UDP",
			},
			"device_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "DeviceID where service is runnig",
			},
			"description": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Descending units",
			},
		},
	}
}

func resourceServiceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient)
	netboxService := new(models.WritableService)
	name := d.Get("name").(string)
	netboxService.Name = &name
	netboxService.Port = swag.Int64(int64(d.Get("port").(int)))
	netboxService.Protocol = swag.Int64(protocol[d.Get("protocol").(string)])
	netboxService.Description = d.Get("description").(string)
	netboxService.Device = swag.Int64(d.Get("device_id").(int64))
	params := ipam.NewIPAMServicesCreateParams()
	params.WithData(netboxService)

	res, err := c.netboxClient.IPAM.IPAMServicesCreate(params, nil)
	if err != nil {
		log.Print(err)
		return err
	}

	log.Print("Service ID is: ", res.Payload.ID)
	d.SetId(strconv.FormatInt(res.Payload.ID, 10))
	return resourceServiceRead(d, m)
}

func resourceServiceRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient).netboxClient
	params := ipam.NewIPAMServicesListParams()
	//todo: Search by name or ID if provided
	switch {
	case d.Id() != "":
		params.WithID(swag.String(d.Id()))
	case d.Get("name") != nil:
		name := d.Get("name").(string)
		params.WithName(&name)
	}
	res, err := c.IPAM.IPAMServicesList(params, nil)
	if err != nil {
		log.Print("[DEBUG] Cant read Rack info resourceRackRead() ", err)
	}
	d.Set("name", res.Payload.Results[0].Name)
	d.Set("port", res.Payload.Results[0].Port)
	d.Set("protocol", res.Payload.Results[0].Protocol.Label)
	d.Set("device_id", res.Payload.Results[0].Device.ID)
	d.Set("description", res.Payload.Results[0].Description)

	return nil
}
func resourceServiceUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient)
	netboxService := new(models.WritableService)
	name := d.Get("name").(string)
	netboxService.Name = &name
	netboxService.Port = swag.Int64(int64(d.Get("port").(int)))
	netboxService.Protocol = swag.Int64(protocol[d.Get("protocol").(string)])
	netboxService.Description = d.Get("description").(string)
	netboxService.Device = swag.Int64(d.Get("device_id").(int64))

	params := ipam.NewIPAMServicesUpdateParams()
	serviceID, err := strconv.Atoi(d.Id())
	if err != nil {
		log.Print("string converting failed")
	}
	params.WithID(int64(serviceID))
	params.WithData(netboxService)
	_, err = c.netboxClient.IPAM.IPAMServicesUpdate(params, nil)
	if err != nil {
		log.Print("[DEBUG] Update Service failed\n", err)
	} else {
		log.Print("Updated...")
	}

	return resourceRackRead(d, m)
}

func resourceServiceDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient).netboxClient
	params := ipam.NewIPAMServicesDeleteParams()
	rackID, err := strconv.Atoi(d.Id())
	if err != nil {
		log.Print("string converting failed")
	}
	params.WithID(int64(rackID))
	_, err = c.IPAM.IPAMServicesDelete(params, nil)
	if err != nil {
		log.Print("[DEBUG] Delete service failed\n", err)
	}

	return nil
}
