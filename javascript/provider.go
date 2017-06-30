package javascript

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"javascript_script": resourceJavascriptScript(),
		},
		ConfigureFunc: provideConfigure,
	}
}

func provideConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{}

	return config.Vm()
}

func validateJsType(v interface{}, k string) (ws []string, errors []error) {
	if containsString(JS_TYPES, v.(string)) == false {
		errors = append(errors, fmt.Errorf("Valid value types are %q, %q, %q, and %q", JS_BOOL, JS_FLOAT, JS_INT, JS_STRING))
	}
	return
}
