# Terraform Javascript Provider

The JavaScript provider is used to perform javascript operations on data
obtained during Terraform runs. For example a resource has a property that is
a JSON formatted string, and you need the value from one of the nested properties
you can use the JavaScript provider to parse the JSON string and return the value.

The provider uses [`otto`](https://github.com/robertkrimen/otto) to run JavaScript
code and includes the underscore library.

**Example Usage**

```
variable "data" {
  type = "string"
  default = "{\"foo\": \"bar\"}"
}

provider "javascript" {}

resource "javascript_script" "s" {
  script = "context.foo = JSON.parse(context.json_string).foo"
  context = {
    json_string = "${var.data}"
  }
}

output val {
  value = "${javascript_script.s.context.foo}"
}

// val = "bar"
```

---

**Argument Reference**

The JavaScript provider currently takes no arguments

## Resources

* [javascript_script]()

### javascript_script

Runs a piece of javascript code and returns an updated context object. The context
object can be set with initial values an can be updated by the script by referencing
it with `context.*`. The global variable `operation` is also set and available to
the script. This allows for specific code to be run based on the operation type.
Please note that all values in context are converted to strings due to the way
terraform handles `TypeMap` values

**Argument Reference**

* `script` - (Required) The javascript to run
* `context` - (Optional) Optional `TypeMap` to make available in the script

### Building

If the builds in the current release do not include the OS you require, you can download the
source and create a custom distribution.

**Requirements**

* go tools installed
* docker installed
* source cloned to `$GOPATH/src/github.com/bhoriuchi/terraform-provider-javascript`

**Steps**

The following will set the build environment and use a docker instance to build the requested
distribution packages.

1. cd to `$GOPATH/src/github.com/bhoriuchi/terraform-provider-javascript`
2. Set the environment variable `GOX_OS_ARCH` to the desired os/architecture combination. See [gox usage](https://github.com/mitchellh/gox#usage)
3. run `make dist`

### Credits

* Build scripts and project outline copied from [https://github.com/prudhvitella/terraform-provider-infoblox](https://github.com/prudhvitella/terraform-provider-infoblox)