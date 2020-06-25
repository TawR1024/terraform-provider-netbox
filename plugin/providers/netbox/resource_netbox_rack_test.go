package netbox

import (
	"testing"

	acctest2 "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func TestAccNetboxRack(t *testing.T) {
	var rack models.WritableRack
	rackName := acctest2.RandString(3)

}
