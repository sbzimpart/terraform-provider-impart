# Create a new specification
resource "impart_spec" "example" {
  name        = "spec_example"
  source_file = "${path.module}/spec.yaml"
}
