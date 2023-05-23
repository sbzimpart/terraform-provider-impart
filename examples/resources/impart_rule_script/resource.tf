
# Create a new rule script
resource "impart_rule_script" "example" {
  name        = "example"
  disabled    = false
  description = "Rule description"
  source_file = "${path.module}/rule.js"
}
