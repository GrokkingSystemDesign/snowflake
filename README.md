# snowflake
## Install
Set up your Go environment and run

`go get github.com/GrokkingSystemDesign/snowflake`

## Usage
```go
package main

import (
	"fmt"

	"github.com/GrokkingSystemDesign/snowflake"
)

func main() {
	m, err := snowflake.NewMachine(0)
	if err != nil {
		panic(err)
	}
	id, _ := m.Generate()
	fmt.Printf("Int64 ID: %d\n", id.Int64())
	fmt.Printf("String ID: %s\n", id.String())
	fmt.Printf("Base64 ID: %s\n", id.Base64())
}
```

