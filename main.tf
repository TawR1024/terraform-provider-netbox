provider "netbox" {
  url   = "netbox.url"
  token = "token1234"
}

resource "netbox_rack" "testRack2" {
  name       = "testRack"
  facility   = "facilityID"
  site       = "TestSite"
  tenant     = "VPC"
  width      = 19
  height     = 46
  desc_units = true
}

resource "netbox_virtual_machine" "terraformTestVM" {
  name    = "terraformTestVM"
  cluster = "testCluster"
  tenant  = "VPC"
  cores   = 1
  ram     = 1024
  disk    = 56
  status  = 1
}

resource "netbox_device" "cmp1_test" {
  name     = "cmp1"
  site     = "TestSite"
  tenant   = "VPC"
  rack     = "TestRack"
  position = 1
  type     = "Generic-1U"
  role     = "cmp"
  status   = "Planned"
}

resource "netbox_interface" "eth1" {
  name   = "eth1"
  device = "${netbox_device.cmp1_test.id}"
}
