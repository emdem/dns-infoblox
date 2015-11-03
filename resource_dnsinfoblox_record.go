package dnsinfoblox

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	infoblox "github.com/fanatic/go-infoblox"
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
	client := meta.(*infoblox.Client)
	rec, err := client.FindRecordA(d.Get("name").(string), d.Id())
	if err != nil {
		return fmt.Errorf("Couldn't find infoblox record: %s", err)
	}


	return nil
}


func resourceDNSInfobloxRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}


func resourceDNSInfobloxRecordDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
