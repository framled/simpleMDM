# API GO for Simple MDM

## Usage

```go
package main
import mdm "github.com/framled/simpleMDM/v1/device"

func main() {
    mdm := mdm.NewMDM(apiKey, testServer.URL)
    devices, err := mdm.findAll(0, 0)
}
```