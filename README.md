# SecretString

SecretString provides an encoding.TextUnmarshaler interface around a string
that will retrieve a secret stored in a project's 
[Google Secret Manager](https://cloud.google.com/secret-manager/docs). It is
meant to be used with Google Cloud Functions and 
[Kelsey Hightower's envconfig library](https://github.com/kelseyhightower/envconfig).

## Requirements

The following is required:
* The Go code that uses this needs a `GCP_PROJECT` environment variable set
* It is a bit of a hack, but set the `default` tag to the name of the secret (see example below)

## Instructions
To use this in your project, pull it down:

```bash
go get -u github.com/brockwood/secretstring
```

### Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/brockwood/secretstring"
    "github.com/kelseyhightower/envconfig"
)

type Config struct {
    SuperSecret secretstring.SecretString `default:"SecretName"`
    GCPProject  string `split_words:"true"`
}

func main() {
    var c Config
    err := envconfig.Process("", &c)
    if err != nil {
        log.Fatal(err.Error())
    }
    format := "SuperSecret: %s\nGCP_PROJECT: %s\n"
    _, err = fmt.Printf(format, c.SuperSecret, c.GCPProject)
    if err != nil {
        log.Fatal(err.Error())
    }
}
```

The output:

```bash
Windows4Ever:testenv rockwoodson$ GCP_PROJECT=my-project go run main.go
SuperSecret: Drink your Ovaltine
GCP_PROJECT: my-project
Windows4Ever:testenv rockwoodson$
```

#### Note
The GCPProject struct member above is not required for the SecretString decoder
to work and is only there as additional debugging output.