package netbox

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client/virtualization"
	"github.com/netbox-community/go-netbox/netbox/models"
	"log"
	"strconv"
)

func resourceVM() *schema.Resource {
	return &schema.Resource{
		Create: resourceVMCreate,
		Read:   resourceVMRead,
		Update: resourceVMUpdate,
		Delete: resourceVMDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vm_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: false,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"cluster": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: false,
			},
			"cores": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: false,
			},
			"ram": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: false,
			},
			"disk": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: false,
			},
			"status": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: false,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
		},
	}
}

func resourceVMCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient).client
	var netboxVM models.WritableVirtualMachineWithConfigContext
	name := d.Get("name").(string)
	netboxVM.Name = &name
	netboxVM.Tags = []string{}
	netboxVM.Cluster = swag.Int64(int64(d.Get("cluster").(int)))
	netboxVM.Status = int64(d.Get("status").(int))

	netboxVM.Vcpus = swag.Int64(int64(d.Get("cores").(int)))
	netboxVM.Memory = swag.Int64(int64(d.Get("ram").(int)))
	netboxVM.Disk = swag.Int64(int64(d.Get("disk").(int)))
	params := virtualization.NewVirtualizationVirtualMachinesCreateParams()
	params.WithData(&netboxVM)
	res, err := c.Virtualization.VirtualizationVirtualMachinesCreate(params, nil)
	if err != nil {
		log.Print(err)
		return err
	}

	log.Print("VM ID is: ", res.Payload.ID)
	d.SetId(strconv.FormatInt(res.Payload.ID,10))
	d.Set("vm_id", res.Payload.ID)
	return resourceVMRead(d,m)
}

func resourceVMRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient).client
	name := d.Get("name").(string)
	params :=virtualization.NewVirtualizationVirtualMachinesListParams()
	params.WithName(&name)
	res, err := c.Virtualization.VirtualizationVirtualMachinesList(params,nil)
	if err != nil{
		log.Print("[DEBUG] Cant read VM info resourceVMRead() ",err)
	}
	d.Set("vm_id", res.Payload.Results[0].ID)
	d.Set("name", res.Payload.Results[0].Name)
	d.Set("cluster", res.Payload.Results[0].Cluster)
	d.Set("cores", res.Payload.Results[0].Vcpus)
	d.Set("ram", res.Payload.Results[0].Memory)
	d.Set("disk", res.Payload.Results[0].Disk)
	d.Set("tags", res.Payload.Results[0].Tags)

	return nil
}

func resourceVMUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient).client
	data := new(models.WritableVirtualMachineWithConfigContext)
	name := swag.String(d.Get("name").(string))

	data.Name = name
	data.Cluster = swag.Int64(int64(d.Get("cluster").(int)))
	data.Disk = swag.Int64(int64(d.Get("disk").(int)))
	data.Vcpus = swag.Int64(int64(d.Get("cores").(int)))
	data.Memory = swag.Int64(int64(d.Get("ram").(int)))
	data.Tags = []string{}

	params := virtualization.NewVirtualizationVirtualMachinesPartialUpdateParams()
	vmId:= d.Get("vm_id").(int)
	params.WithID(int64(vmId))
	params.WithData(data)
	_, err := c.Virtualization.VirtualizationVirtualMachinesPartialUpdate(params, nil)
	if err != nil{
		log.Print("[DEBUG] Update VM failed\n", err)
	}else{
		log.Print("Updated...")
	}

	return resourceVMRead(d,m)
}

func resourceVMDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient).client
	params := virtualization.NewVirtualizationVirtualMachinesDeleteParams()
	vmId:= d.Get("vm_id").(int)
	params.WithID(int64(vmId))
	_, err := c.Virtualization.VirtualizationVirtualMachinesDelete(params, nil)
	if err != nil{
		log.Print("[DEBUG] Delete VM failed\n", err)
	}

	return nil
}
