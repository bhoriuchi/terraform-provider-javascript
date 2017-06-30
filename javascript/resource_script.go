package javascript

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/robertkrimen/otto"
	"time"
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
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "JSON formatted string that is used as the context",
				Default:     "{}",
			},
			"values": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Variable names to set as values",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Variable name to return",
						},
						"type": &schema.Schema{
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateJsType,
							Description:  "Type to coerce variable to",
						},
					},
				},
			},
		},
	}
}

func resourceJavascriptScriptCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceJavascriptScriptRead(d, meta)
}

func resourceJavascriptScriptUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceJavascriptScriptRead(d, meta)
}

func resourceJavascriptScriptDelete(d *schema.ResourceData, meta interface{}) error {
	return resourceJavascriptScriptRead(d, meta)
}

func resourceJavascriptScriptRead(d *schema.ResourceData, meta interface{}) error {
	vm := meta.(*otto.Otto)
	script := d.Get("script")
	// context := d.Get("context")
	d.SetId(time.Now().UTC().String())

	if result, err := vm.Run(script); err == nil {
		fmt.Printf("[javascript-provider] Result: %v+", result)

		// set values if they exist
		if values, ok := d.GetOk("values"); ok {
			for _, v := range values.([]interface{}) {
				vmap := v.(map[string]interface{})

				name := vmap["name"].(string)
				if value, err := vm.Get(name); err == nil {
					switch t := vmap["type"].(string); t {
					case JS_BOOL:
						if val, err := value.ToBoolean(); err == nil {
							if err := d.Set(name, val); err != nil {
								return fmt.Errorf("Failed to set field %q to %q; %v+", name, val, err)
							}
						} else {
							fmt.Printf("[javascript-provider] Failed to convert value %q to %q", value, t)
						}
					case JS_FLOAT:
						if val, err := value.ToFloat(); err == nil {
							if err := d.Set(name, val); err != nil {
								return fmt.Errorf("Failed to set field %q to %q; %v+", name, val, err)
							}
						} else {
							fmt.Printf("[javascript-provider] Failed to convert value %q to %q", value, t)
						}
					case JS_INT:
						if val, err := value.ToInteger(); err == nil {
							if err := d.Set(name, val); err != nil {
								return fmt.Errorf("Failed to set field %q to %q; %v+", name, val, err)
							}
						} else {
							fmt.Printf("[javascript-provider] Failed to convert value %q to %q", value, t)
						}
					case JS_STRING:
						if val, err := value.ToString(); err == nil {
							if err := d.Set(name, val); err != nil {
								return fmt.Errorf("Failed to set field %q to %q; %v+", name, val, err)
							}
						} else {
							fmt.Printf("[javascript-provider] Failed to convert value %q to %q", value, t)
						}
					}
				} else {
					fmt.Printf("[javascript-provider] Failed to get value %q from vm", name)
				}
			}
		}
	} else {
		return fmt.Errorf("Script run error: %v+", err)
	}

	return nil
}
