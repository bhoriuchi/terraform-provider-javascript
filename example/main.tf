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