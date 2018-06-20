https://travis-ci.org/AmirSoleimani/IRValidator.svg?branch=master

# Installation

Just use go get.

    go get github.com/AmirSoleimani/IRValidator

And then just import the package into your own code.

```go
import (
	"github.com/AmirSoleimani/IRValidator"
)
```

# Usage


```go
package main

import (
	"fmt"
	iv "github.com/AmirSoleimani/IRValidator"
)

func main() {
	// CardNumber check
	if err := iv.validate.CardNumber("6221061049447982"); err != nil {
		fmt.Println(err)
	}

	// NationalID check
	if err := iv.validate.NationalID("1111111111"); err != nil {
		fmt.Println(err)
	}

	// IBAN check
	if err := iv.validate.IBAN("IR340570030280001175105001"); err != nil {
		fmt.Println(err)
	}
}

```