provider "javascript" {

}

resource "javascript_script" "s" {
  script = "context.test1 = 'ok'; context.test2 = 1; context.test3 = true"
  context = {
    test1 = "himom"
    test2 = 0
  }
}