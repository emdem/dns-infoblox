package dns-infoblox

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	infoblox "github.com/emdem/go-infoblox"
)

func resourceDNSInfobloxRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSInfobloxRecordCreate,
		Read:   resourceDNSInfobloxRecordRead,
		Update: resourceDNSInfobloxRecordUpdate,
		Delete: resourceDNSInfobloxRecordDelete,

		Schema: map[string]*schema.Schema{

		}
	}
}

func resourceDNSInfobloxRecordCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}


func resourceDNSInfobloxRecordRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}


func resourceDNSInfobloxRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}


func resourceDNSInfobloxRecordDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
