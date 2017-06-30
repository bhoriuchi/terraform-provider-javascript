package javascript

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
	"log"
	"time"
)

const (
	CREATE = "create"
	READ   = "read"
	UPDATE = "update"
	DELETE = "delete"
)

func resourceJavascriptScript() *schema.Resource {
	return &schema.Resource{
		Create: resourceJavascriptScriptCreate,
		Read:   resourceJavascriptScriptRead,
		Update: resourceJavascriptScriptUpdate,
		Delete: resourceJavascriptScriptDelete,

		Schema: map[string]*schema.Schema{
			"script": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "JavaScript source to run",
			},
			"context": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceJavascriptScriptCreate(d *schema.ResourceData, meta interface{}) error {
	return runScript(d, CREATE)
}

func resourceJavascriptScriptRead(d *schema.ResourceData, meta interface{}) error {
	return runScript(d, READ)
}

func resourceJavascriptScriptUpdate(d *schema.ResourceData, meta interface{}) error {
	return runScript(d, UPDATE)
}

func resourceJavascriptScriptDelete(d *schema.ResourceData, meta interface{}) error {
	return runScript(d, DELETE)
}

func runScript(d *schema.ResourceData, operation string) error {
	vm := otto.New()
	vm.Set("operation", operation)
	script := d.Get("script").(string)
	d.SetId(time.Now().UTC().String())

	if context, ok := d.GetOk("context"); ok {
		vm.Set("context", context)

		if result, err := vm.Run(script + "\nresult = context;"); err == nil {
			log.Printf("[javascript-provider] Result: %v+", result)

			if res, err := result.Export(); err == nil {
				ctx := res.(map[string]interface{})

				// there seems to be an issue with terraform TypeMaps that only allow
				// the map to contain a single type. So in order to homogenize the the
				// map, convert everything to a string
				for k, v := range ctx {
					ctx[k] = fmt.Sprintf("%v", v)
				}

				if err := d.Set("context", ctx); err != nil {
					log.Printf("[infoblox-provider] Failed to set context: %v+", err)
				}
			} else {
				log.Printf("[javascript-provider] Failed to export context: %v+", err)
			}
		} else {
			return fmt.Errorf("Script run error: %v+", err)
		}
	} else {
		if _, err := vm.Run(script); err != nil {
			return fmt.Errorf("Script run error: %v+", err)
		}
	}

	return nil
}
