---
layout: "vault"
page_title: "Vault: vault_database_secret_backend_rotate_root resource"
sidebar_current: "docs-vault-resource-database-secret-backend-rotate-root"
description: |-
  Rotates the root password for a database secret backend connection in Vault.
---

# vault\_database\_secret\_backend\_rotate_root

Rotates the root password for a database secret backend connection in Vault.
This resource does not create anything. It is "created" every time the plan is executed. 

~> **Important** You will not be able to retrieve the password again. Use separate credentials for vault. See
[Official Vault Guide: Database Root Credential Rotation ](https://www.vaultproject.io/guides/secret-mgmt/db-root-rotation.html)
for more details.

## Example Usage

```hcl
resource "vault_database_secret_backend_connection" "mysql" {
  backend = "database"
  name    = "mysql"

  allowed_roles = [
    "*",
  ]

  mysql {
    connection_url = "{{username}}:{{password}}@tcp(host:port)/"
    username       = "username"
    password       = "password"
  }
}

resource "vault_database_secret_backend_rotate_root" "mysql" {
  backend       = "database"
  db_connection = "${vault_database_secret_backend_connection.mydb.name}"
}
```

## Argument Reference

The following arguments are supported:

* `backend` - (Required) The unique name of the Vault mount to configure.

* `db_connection` - (Required) The unique name of the database connection to use.

## Attributes Reference

No additional attributes are exported by this resource.
