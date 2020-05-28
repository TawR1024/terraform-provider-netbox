package netbox

import (
	"log"

	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client/virtualization"
	"github.com/netbox-community/go-netbox/netbox/models"

	"strconv"
)

func resourceNetboxVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Create: resourceVMCreate,
		Read:   resourceVMRead,
		Update: resourceVMUpdate,
		Delete: resourceVMDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ram": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"disk": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVMCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient)
	netboxVM := new(models.WritableVirtualMachineWithConfigContext)
	name := d.Get("name").(string)

	netboxVM.Name = &name
	netboxVM.Tags = []string{}
	netboxVM.Cluster = c.GetClusterID(swag.String(d.Get("cluster").(string)))
	netboxVM.Tenant = c.GetTenantId(swag.String(d.Get("tenant").(string)))
	netboxVM.Status = int64(d.Get("status").(int))
	netboxVM.Vcpus = swag.Int64(int64(d.Get("cores").(int)))
	netboxVM.Memory = swag.Int64(int64(d.Get("ram").(int)))
	netboxVM.Disk = swag.Int64(int64(d.Get("disk").(int)))

	params := virtualization.NewVirtualizationVirtualMachinesCreateParams()
	params.WithData(netboxVM)

	res, err := c.netboxClient.Virtualization.VirtualizationVirtualMachinesCreate(params, nil)
	if err != nil {
		log.Print(err)
		return err
	}

	log.Print("VM ID is: ", res.Payload.ID)
	d.SetId(strconv.FormatInt(res.Payload.ID, 10))
	return resourceVMRead(d, m)
}

func resourceVMRead(d *schema.ResourceData, m interface{}) error { //todo: Bug -- method returns info about first result in response.
	c := m.(*ProviderNetboxClient).netboxClient
	params := virtualization.NewVirtualizationVirtualMachinesListParams()
	//todo: Search by name or ID if provided
	switch {
	case d.Id() != "":
		params.WithID(swag.String(d.Id()))
	case d.Get("name") != nil:
		name := d.Get("name").(string)
		params.WithName(&name)
	}
	res, err := c.Virtualization.VirtualizationVirtualMachinesList(params, nil)
	if err != nil {
		log.Print("[DEBUG] Cant read VM info resourceVMRead() ", err)
	}
	d.Set("name", res.Payload.Results[0].Name)
	d.Set("cluster", res.Payload.Results[0].Cluster.Name)
	d.Set("tenant", res.Payload.Results[0].Tenant.Name)
	d.Set("cores", res.Payload.Results[0].Vcpus)
	d.Set("ram", res.Payload.Results[0].Memory)
	d.Set("disk", res.Payload.Results[0].Disk)
	d.Set("tags", res.Payload.Results[0].Tags)
	d.Set("status", res.Payload.Results[0].Status.Value)

	return nil
}

func resourceVMUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient)
	netboxVM := new(models.WritableVirtualMachineWithConfigContext)
	name := swag.String(d.Get("name").(string))

	netboxVM.Name = name
	netboxVM.Cluster = c.GetClusterID(swag.String(d.Get("cluster").(string)))
	netboxVM.Tenant = c.GetTenantId(swag.String(d.Get("tenant").(string)))
	netboxVM.Disk = swag.Int64(int64(d.Get("disk").(int)))
	netboxVM.Vcpus = swag.Int64(int64(d.Get("cores").(int)))
	netboxVM.Memory = swag.Int64(int64(d.Get("ram").(int)))
	netboxVM.Tags = []string{}

	params := virtualization.NewVirtualizationVirtualMachinesPartialUpdateParams()
	vmId, err := strconv.Atoi(d.Id())
	if err != nil {
		log.Print("string converting failed")
	}
	params.WithID(int64(vmId))
	params.WithData(netboxVM)
	_, err = c.netboxClient.Virtualization.VirtualizationVirtualMachinesPartialUpdate(params, nil)
	if err != nil {
		log.Print("[DEBUG] Update VM failed\n", err)
	} else {
		log.Print("Updated...")
	}

	return resourceVMRead(d, m)
}

func resourceVMDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*ProviderNetboxClient).netboxClient
	params := virtualization.NewVirtualizationVirtualMachinesDeleteParams()
	vmId, err := strconv.Atoi(d.Id())
	if err != nil {
		log.Print("string converting failed")
	}
	params.WithID(int64(vmId))
	_, err = c.Virtualization.VirtualizationVirtualMachinesDelete(params, nil)
	if err != nil {
		log.Print("[DEBUG] Delete VM failed\n", err)
	}

	return nil
}
