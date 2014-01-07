# Diffbot API Client Library for Go (stub)

> This is not a final version. Please don't try to use this library yet.

## Usage

```go
import (
	"fmt"
	"github.com/attilaolah/diffbot"
)

const APIKey = "my-secret-api-key"

func main() {
	a, err := diffbot.Article(APIKey, "http://attilaolah.eu")
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}
```
