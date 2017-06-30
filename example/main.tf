provider "javascript" {

}

resource "javascript_script" "s" {
  script = "test = 'one';"
  values = [
    {
      name = "test"
      type = "string"
    }
  ]
}