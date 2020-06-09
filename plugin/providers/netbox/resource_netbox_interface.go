package netbox

import (
	"log"
	"strconv"

	"github.com/netbox-community/go-netbox/netbox/client/dcim"

	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func resourceNetboxInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceCreate,
		Read:   resourceInterfaceRead,
		Update: resourceInterfaceUpdate,
		Delete: resourceInterfaceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1200,
			},
		},
	}
}

func resourceInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient)
	netboxInterface := new(models.WritableDeviceInterface)
	name := d.Get("name").(string)

	netboxInterface.Name = &name
	netboxInterface.Type = int64(d.Get("type").(int))
	netboxInterface.Tags = []string{}
	netboxInterface.Device = swag.Int64(int64(d.Get("device").(int)))

	params := dcim.NewDcimInterfacesCreateParams()
	params.WithData(netboxInterface)
	res, err := c.netboxClient.Dcim.DcimInterfacesCreate(params, nil)
	if err != nil {
		log.Print(err)
		return err
	}

	log.Print("Interface ID is: ", res.Payload.ID)
	d.SetId(strconv.FormatInt(res.Payload.ID, 10))
	return resourceInterfaceRead(d, m)
}

func resourceInterfaceRead(d *schema.ResourceData, m interface{}) error { //todo: Bug -- method returns info about first result in response.
	c := m.(*ProviderNetboxClient).netboxClient
	params := dcim.NewDcimInterfacesReadParams()
	//todo: Search by name or ID if provided
	interfaceID, _ := strconv.Atoi(d.Id())
	params.WithID(int64(interfaceID))

	res, err := c.Dcim.DcimInterfacesRead(params, nil)
	if err != nil {
		log.Print("[DEBUG] Cant read Interface info resourceInterfaceRead() ", err)
	}
	d.Set("name", res.Payload.ID)
	d.Set("type", res.Payload.Type)
	d.Set("device", res.Payload.Device.ID)

	return nil
}

func resourceInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
