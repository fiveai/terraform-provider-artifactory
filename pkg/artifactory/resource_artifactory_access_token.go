package artifactory

import (
	"context"
	"github.com/atlassian/go-artifactory/v2/artifactory"
	"github.com/atlassian/go-artifactory/v2/artifactory/v2"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceArtifactoryAccessToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessTokenCreate,
		Read:   resourceAccessTokenRead,
		Delete: resourceAccessTokenDelete,
		Exists: resourceAccessTokenExists,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"groups": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
			},
			"token": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAccessTokenRead(d *schema.ResourceData, m interface{}) error {
	// There isn't an API exposed for this purpose - see https://www.jfrog.com/jira/browse/RTFACT-19071
	return nil
}

func resourceAccessTokenCreate(s *schema.ResourceData, m interface{}) error {
	c := m.(*artifactory.Artifactory)
	d := &ResourceData{s}

	options :=  v2.AccessTokenOptions{
		Username: *d.getStringRef("username"),
		Groups: *d.getSetRef("groups"),
		ExpiresIn: 0,
	}

	token, err := c.V2.Token.CreateAccessToken(context.Background(), &options)

	if err != nil {
		return err
	}

	d.SetId(options.Username)

	err = d.Set("token", token.Token)

	if err != nil {
		return err
	}

	return resourceAccessTokenRead(s, m)
}

func resourceAccessTokenDelete(s *schema.ResourceData, m interface{}) error {
	c := m.(*artifactory.Artifactory)
	d := &ResourceData{s}

	_, err := c.V2.Token.RevokeAccessToken(context.Background(), *d.getStringRef("token"))

	if err != nil {
		return nil
	}

	return err
}

func resourceAccessTokenExists(d *schema.ResourceData, m interface{}) (bool, error) {
	// There isn't an API exposed for this purpose - see https://www.jfrog.com/jira/browse/RTFACT-19071
	return true, nil
}
