package netbox

import (
	"fmt"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client/virtualization"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func resourceVM() *schema.Resource {
	return &schema.Resource{
		Create: resourceVMCreate,
		Read:   resourceVMRead,
		Update: resourceVMUpdate,
		Delete: resourceVMDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"cores": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
			},
			"ram": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
			},
			"disk": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
			},
			"status": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
			},
		},
	}
}

func resourceVMCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	var netboxVM models.WritableVirtualMachineWithConfigContext
	name := d.Get("name").(string)
	netboxVM.Name = &name
	netboxVM.Tags = []string{}
	netboxVM.Cluster = swag.Int64(int64(d.Get("cluster").(int)))
	//netboxVM.Status = d.Get("status").(int)

	netboxVM.Vcpus = swag.Int64(int64(d.Get("cores").(int)))
	netboxVM.Memory = swag.Int64(d.Get("ram").(int64))

	netboxVM.Disk = swag.Int64(d.Get("disk").(int64))

	params := virtualization.NewVirtualizationVirtualMachinesCreateParams()
	params.WithData(&netboxVM)

	res, err := c.NetboxClient.Virtualization.VirtualizationVirtualMachinesCreate(params, nil)
	if err != nil {
		fmt.Print(err)
		return err
	}

	fmt.Print("VM ID is: ", res.Payload.ID)
	return nil
}

func resourceVMRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceVMUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceVMRead(d, m)
}

func resourceVMDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
