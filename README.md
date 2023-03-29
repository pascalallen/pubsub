# pubsub

[![Go Reference](https://pkg.go.dev/badge/github.com/pascalallen/pubsub.svg)](https://pkg.go.dev/github.com/pascalallen/pubsub)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pascalallen/pubsub)
[![Go Report Card](https://goreportcard.com/badge/github.com/pascalallen/pubsub)](https://goreportcard.com/report/github.com/pascalallen/pubsub)
![GitHub Workflow Status (with branch)](https://img.shields.io/github/actions/workflow/status/pascalallen/pubsub/go.yml?branch=main)
![GitHub](https://img.shields.io/github/license/pascalallen/pubsub)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/pascalallen/pubsub)

pubsub is a Go module that offers a concurrent pub/sub service leveraging goroutines and channels.

## Installation

Use the Go CLI tool [go](https://go.dev/dl/) to install hmac.

```bash
go get github.com/pascalallen/pubsub
```

## Usage

```go
...

import (
    "fmt"
    "github.com/pascalallen/pubsub"
)

...

c := make(chan pubsub.Message)
t := pubsub.CreateTopic("my-topic", c)
m := pubsub.CreateMessage([]byte("hello, jupiter!"))
t.AddMessage(*m)

s := pubsub.CreateSubscriber("my-subscription", *t, c)

go t.Publish()

for m := range s.Channel {
	fmt.Printf("Message received for subscriber %s: %s\n", s.Name, m)
}

...
```

## Testing

Run tests and create coverage profile:

```bash
go test -covermode=count -coverprofile=coverage.out
```

View test coverage profile:

```bash
go tool cover -html=coverage.out
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](LICENSE)
