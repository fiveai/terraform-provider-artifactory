module github.com/fiveai/terraform-provider-artifactory

require (
	github.com/atlassian/go-artifactory/v2 v2.1.1
	github.com/atlassian/terraform-provider-artifactory v1.5.0
	github.com/hashicorp/terraform v0.11.13
)

replace github.com/atlassian/go-artifactory/v2 => github.com/fiveai/go-artifactory/v2 v2.3.0

replace github.com/atlassian/terraform-provider-artifactory => ./
