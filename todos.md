# Todos/Questions for Sentences

1. Create a VSCode snippet for a Test func.
2. Add proper DNS friendly name (e.g. Route53) to template.yaml.
3. Add deploy target to Makefile
4. The sentences package should really be moved out into a separate git repo/go
   package. Otherwise, we'll end up with coupling.
5. Mock ParseText via a test interface in main_test.go. Otherwise this is an
   integration test, not a unit test.
6. Troubleshoot why API Gateway is slow.
