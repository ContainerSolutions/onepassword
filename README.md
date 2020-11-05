# 1Password Client
Thin wrapper around the 1Password CLI for use in Golang.

## Usage
First [install the 1Password CLI](https://support.1password.com/command-line/).

Import the package, create a client, and retrieve an item.
```go
import (
    "os"
    "log"

    op "github.com/ContainerSolutions/onepassword"
)

func main() {
    password := os.GetEnv("OP_PASSWORD")
    secretKey := os.GetEnv("OP_SECRET_KEY")

    client, err := op.NewClient("op", "subdomain", "test@subdomain.com", password, secretKey)
    if err != nil {
        log.Println(err.Error())
    }

    item, err := client.GetItem(op.VaultName("test-vault"), op.ItemName("test-item"))
    if err != nil {
        log.Println(err.Error())
    }
}
```
