package vault

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/vault/api"
)

func databaseSecretBackendRotateRootResource() *schema.Resource {
	return &schema.Resource{
		Read: func(d *schema.ResourceData, meta interface{}) error {
			return nil
		},
		Delete: func(d *schema.ResourceData, meta interface{}) error {
			return nil
		},
		Exists: func(d *schema.ResourceData, meta interface{}) (bool, error) {
			return false, nil
		},

		Create: databaseSecretBackendConnectionRotateRootPassword,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"db_connection": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the database connection.",
			},
			"backend": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Unique name of the Vault mount to configure.",
				// standardise on no beginning or trailing slashes
				StateFunc: func(v interface{}) string {
					return strings.Trim(v.(string), "/")
				},
			},
		},
	}
}

func databaseSecretBackendConnectionRotateRootPassword(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	d.Set("rotate_password_on_change", false)

	backend := d.Get("backend").(string)
	name := d.Get("db_connection").(string)
	path := databaseSecretBackendRotatePath(backend, name)
	d.SetId(path)
	_, err := client.RawRequest(client.NewRequest("POST", path))
	if err != nil {
		return err
	}

	d.Set("rotate_password_on_change", true)
	return nil
}

func databaseSecretBackendRotatePath(backend, name string) string {
	return "/v1/" + strings.Trim(backend, "/") + "/rotate-root/" + strings.Trim(name, "/")
}
