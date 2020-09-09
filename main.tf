provider "netbox" {
  //  url   = "netbox.cloud.selectel.org"
  url   = ""
  token = ""
}

resource "netbox_rack" "testRack2" {
  name       = "testRack"
  facility   = "dubrovka-test"
  site       = "ru-9.os.selectel.org"
  tenant     = "VPC"
  width      = 19
  height     = 46
  desc_units = true
}
//resource "netbox_device" "img1" {
//  name = "img1.ru-9.os.selectel.org"
//  site = "ru-9.os.selectel.org"
//  tenant = "VPC"
//  rack = "testRack"
//  position = 31
//  type = "Generic-1U"
//  role = "img"
//  status = "Planned"
//  lifecycle {
//    ignore_changes = [
//      status,
//    ]
//  }
//}
//
//
//resource "netbox_interface" "img1eth1" {
//  name   = "eth1"
//  device = "${netbox_device.img1.id}"
//}
//
//resource "netbox_interface" "img1eth2" {
//  name   = "eth2"
//  device = "${netbox_device.img1.id}"
//}
//
//resource "netbox_interface" "img1eth3" {
//  name   = "eth3"
//  device = "${netbox_device.img1.id}"
//}
//resource "netbox_interface" "img1eth4" {
//  name   = "eth4"
//  device = "${netbox_device.img1.id}"
//}
//
//resource "netbox_interface" "img1eth5" {
//  name   = "eth5"
//  device = "${netbox_device.img1.id}"
//}
//
//resource "netbox_interface" "img1ipmi" {
//  name   = "ipmi"
//  device = "${netbox_device.img1.id}"
//}
//
//
//
//resource "netbox_device" "img2" {
//  name     = "img2.ru-9a.os.selectel.org"
//  site     = "ru-9.os.selectel.org"
//  tenant   = "VPC"
//  rack     = "68"
//  position = 5
//  type     = "Generic-4U"
//  role     = "img"
//  status   = "Planned"
//}
//
//resource "netbox_interface" "img2eth0" {
//  name   = "eth0"
//  device = "${netbox_device.img2.id}"
//}
//
//resource "netbox_interface" "img2eth1" {
//  name   = "eth1"
//  device = "${netbox_device.img2.id}"
//}
//
//resource "netbox_interface" "img2eth2" {
//  name   = "eth2"
//  device = "${netbox_device.img2.id}"
//}
//
//resource "netbox_interface" "img2eth3" {
//  name   = "eth3"
//  device = "${netbox_device.img2.id}"
//}
//resource "netbox_interface" "img2eth4" {
//  name   = "eth4"
//  device = "${netbox_device.img2.id}"
//}
//
//resource "netbox_interface" "img2eth5" {
//  name   = "eth5"
//  device = "${netbox_device.img2.id}"
//}
//
//resource "netbox_interface" "img2ipmi" {
//  name   = "ipmi"
//  device = "${netbox_device.img2.id}"
//}
//
//# === img3
//
//resource "netbox_device" "img3" {
//  name     = "img3.ru-9a.os.selectel.org"
//  site     = "ru-9.os.selectel.org"
//  tenant   = "VPC"
//  rack     = "69"
//  position = 5
//  type     = "Generic-4U"
//  role     = "img"
//  status   = "Planned"
//}
//
//resource "netbox_interface" "img3eth0" {
//  name   = "eth0"
//  device = "${netbox_device.img3.id}"
//}
//
//resource "netbox_interface" "img3eth1" {
//  name   = "eth1"
//  device = "${netbox_device.img3.id}"
//}
//
//resource "netbox_interface" "img3eth2" {
//  name   = "eth2"
//  device = "${netbox_device.img3.id}"
//}
//
//resource "netbox_interface" "img3eth3" {
//  name   = "eth3"
//  device = "${netbox_device.img3.id}"
//}
//resource "netbox_interface" "img3eth4" {
//  name   = "eth4"
//  device = "${netbox_device.img3.id}"
//}
//
//resource "netbox_interface" "img3eth5" {
//  name   = "eth5"
//  device = "${netbox_device.img3.id}"
//}
//
//resource "netbox_interface" "img3ipmi" {
//  name   = "ipmi"
//  device = "${netbox_device.img3.id}"
//}
//
//// img4
//
//resource "netbox_device" "img4" {
//  name     = "img4.ru-9a.os.selectel.org"
//  site     = "ru-9.os.selectel.org"
//  tenant   = "VPC"
//  rack     = "70"
//  position = 5
//  type     = "Generic-4U"
//  role     = "img"
//  status   = "Planned"
//}
//
//resource "netbox_interface" "img4eth0" {
//  name   = "eth0"
//  device = "${netbox_device.img4.id}"
//}
//
//resource "netbox_interface" "img4eth1" {
//  name   = "eth1"
//  device = "${netbox_device.img4.id}"
//}
//
//resource "netbox_interface" "img4eth2" {
//  name   = "eth2"
//  device = "${netbox_device.img4.id}"
//}
//
//resource "netbox_interface" "img4eth3" {
//  name   = "eth3"
//  device = "${netbox_device.img4.id}"
//}
//resource "netbox_interface" "img4eth4" {
//  name   = "eth4"
//  device = "${netbox_device.img4.id}"
//}
//
//resource "netbox_interface" "img4eth5" {
//  name   = "eth5"
//  device = "${netbox_device.img4.id}"
//}
//
//resource "netbox_interface" "img4ipmi" {
//  name   = "ipmi"
//  device = "${netbox_device.img4.id}"
//}
//
//# --------------------------------- Connetions
//// img1
//resource "netbox_cable" "img1eth0cable" {
//  device_a_name    = "${netbox_device.img1.name}"
//  interface_a_name = "eth0"
//  device_b_name    = "rvc2-vpc-s1-mgsw02-02"
//  interface_b_name = "Ethernet 5"
//}
//resource "netbox_cable" "img1eth1cable" {
//  device_a_name    = "${netbox_device.img1.name}"
//  interface_a_name = "eth1"
//  device_b_name    = "rvc2-vpc-s1-mgsw02-01"
//  interface_b_name = "Ethernet 5"
//}
//resource "netbox_cable" "img1eth2cable" {
//  device_a_name    = "${netbox_device.img1.name}"
//  interface_a_name = "eth2"
//  device_b_name    = "rvc2-vpc-s1-netsw02-02"
//  interface_b_name = "TenGigabitEthernet 2/0/4"
//}
//resource "netbox_cable" "img1eth3cable" {
//  device_a_name    = "${netbox_device.img1.name}"
//  interface_a_name = "eth3"
//  device_b_name    = "rvc2-vpc-s1-sansw01-02"
//  interface_b_name = "Ethernet 4"
//}
//resource "netbox_cable" "img1eth4cable" {
//  device_a_name    = "${netbox_device.img1.name}"
//  interface_a_name = "eth4"
//  device_b_name    = "rvc2-vpc-s1-netsw02-01"
//  interface_b_name = "TenGigabitEthernet 1/0/4"
//}
//resource "netbox_cable" "img1eth5cable" {
//  device_a_name    = "${netbox_device.img1.name}"
//  interface_a_name = "eth5"
//  device_b_name    = "rvc2-vpc-s1-sansw01-01"
//  interface_b_name = "Ethernet 4"
//}
//resource "netbox_cable" "img1ipmicable" {
//  device_a_name    = "${netbox_device.img1.name}"
//  interface_a_name = "ipmi"
//  device_b_name    = "rvc2-vpc-ipmi-02"
//  interface_b_name = "port 2:5"
//}
//
//// img2
//resource "netbox_cable" "img2eth0cable" {
//  device_a_name    = "${netbox_device.img2.name}"
//  interface_a_name = "eth0"
//  device_b_name    = "rvc2-vpc-s1-mgsw02-01"
//  interface_b_name = "Ethernet 6"
//}
//resource "netbox_cable" "img2eth1cable" {
//  device_a_name    = "${netbox_device.img2.name}"
//  interface_a_name = "eth1"
//  device_b_name    = "rvc2-vpc-s1-mgsw02-02"
//  interface_b_name = "Ethernet 6"
//}
//resource "netbox_cable" "img2eth2cable" {
//  device_a_name    = "${netbox_device.img2.name}"
//  interface_a_name = "eth2"
//  device_b_name    = "rvc2-vpc-s1-netsw02-01"
//  interface_b_name = "TenGigabitEthernet 1/0/5"
//}
//resource "netbox_cable" "img2eth3cable" {
//  device_a_name    = "${netbox_device.img2.name}"
//  interface_a_name = "eth3"
//  device_b_name    = "rvc2-vpc-s1-sansw01-01"
//  interface_b_name = "Ethernet 5"
//}
//resource "netbox_cable" "img2eth4cable" {
//  device_a_name    = "${netbox_device.img2.name}"
//  interface_a_name = "eth4"
//  device_b_name    = "rvc2-vpc-s1-netsw02-02"
//  interface_b_name = "TenGigabitEthernet 2/0/5"
//}
//resource "netbox_cable" "img2eth5cable" {
//  device_a_name    = "${netbox_device.img2.name}"
//  interface_a_name = "eth5"
//  device_b_name    = "rvc2-vpc-s1-sansw01-02"
//  interface_b_name = "Ethernet 5"
//}
//resource "netbox_cable" "img2ipmicable" {
//  device_a_name    = "${netbox_device.img2.name}"
//  interface_a_name = "ipmi"
//  device_b_name    = "rvc2-vpc-ipmi-02"
//  interface_b_name = "port 2:6"
//}


//resource "netbox_vrf" "testVrf" {
//  name   = "testVrf"
//  tenant = "VPC"
//}