# migration

# Usage

### Create a new migration

```go
package main

import (
	"github.com/dnnyjns/migration"
)

func main() {
  migration.Create("first_migration")
}
```

### Run migrations

```go
package main

import (
	_ "github.com/username/repo/migrations"
	"github.com/dnnyjns/migration"
)

func main() {
  db := setupDB()
  defer db.Close()
  migration.Run(db)
}
```
