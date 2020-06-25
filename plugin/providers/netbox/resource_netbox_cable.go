package netbox

import (
	"log"
	"strconv"

	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func resourceNetboxCable() *schema.Resource {
	return &schema.Resource{
		Create: resourceCableCreate,
		Read:   resourceCableRead,
		Update: resourceCableUpdate,
		Delete: resourceCableDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			//"termination_a_type": &schema.Schema{
			//	Type:        schema.TypeString,
			//	Required:    true,
			//	Description: "Side a type",
			//},
			"interface_a_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Interface A name",
			},
			"device_a_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Interface A name",
			},
			//"termination_b_type": &schema.Schema{
			//	Type:        schema.TypeString,
			//	Required:    true,
			//	Description: "Device name",
			//},
			"interface_b_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Device name",
			},
			"device_b_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Device name",
			},
		},
	}
}

func resourceCableCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient)
	netboxCable := new(models.WritableCable)
	params := dcim.NewDcimCablesCreateParams()
	sideAInterfaceId := c.GetInterfaceID(swag.String(d.Get("device_a_name").(string)), swag.String(d.Get("interface_a_name").(string)))
	sideBInterfaceId := c.GetInterfaceID(swag.String(d.Get("device_b_name").(string)), swag.String(d.Get("interface_b_name").(string)))

	netboxCable.TerminationAID = sideAInterfaceId
	netboxCable.TerminationAType = swag.String("dcim.interface")
	netboxCable.TerminationBType = swag.String("dcim.interface")
	netboxCable.TerminationBID = sideBInterfaceId
	params.WithData(netboxCable)

	res, err := c.netboxClient.Dcim.DcimCablesCreate(params, nil)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Print("Interface ID is: ", res.Payload.ID)
	d.SetId(strconv.FormatInt(res.Payload.ID, 10))
	return nil
}

func resourceCableRead(d *schema.ResourceData, m interface{}) error {
	//c := m.(*ProviderNetboxClient).netboxClient
	//params := dcim.NewDcimCablesReadParams()
	////todo: Search by name or ID if provided
	//interfaceID, _ := strconv.Atoi(d.Id())
	//params.WithID(int64(interfaceID))
	//
	//res, err := c.Dcim.DcimCablesRead(params, nil)
	//if err != nil {
	//	log.Print("[DEBUG] Cant read Interface info resourceInterfaceRead() ", err)
	//}
	//d.Set("device_a_name", res.Payload.TerminationA.Name)
	////d.Set("interface_a_name", res.Payload.TerminationA.)
	//d.Set("device_b_name", res.Payload.Type)
	//d.Set("interface_b_name", res.Payload.Device.ID)

	//return nil
	return nil
}
func resourceCableUpdate(data *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCableDelete(data *schema.ResourceData, m interface{}) error {
	return nil
}
