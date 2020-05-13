package main

import (
"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
"github.com/netbox-community/go-netbox/netbox/client/virtualization"
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
		},
	}
}

func resourceVMCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	cluster := int64(d.Get("cluster").(int))
	params:= virtualization.NewVirtualizationVirtualMachinesCreateParams()
	params.Data.Name = &name
	params.Data.Cluster = cluster


	return resourceVMRead(d, m)
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