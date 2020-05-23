provider "netbox" {
  token = ""
  url   = ""
}

resource "netbox_virtual_machine" "terraformTestVM" {
  name    = "terraformTestVM"
  cluster = 8
  cores   = 1
  ram     = 1024
  disk    = 56
  status  = 1
}
