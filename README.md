go-modpacksch
=============

go-modpacksch is a Go client library for accessing the [modpacks.ch API].

## Usage

```go
import "github.com/jamiemansfield/go-modpacksch/modpacksch"
```

Construct a new modpacks.ch client, then use the various services on the
client to access different parts of the modpacks.ch API. For example:

```go
client := modpacksch.NewClient(nil)

// Get the FTB Revelation pack (id = 35)
pack, err := client.Packs.GetPack(35)
```

## License

This library is distributed under the MIT license, found in the [LICENSE.txt]
file.

[modpacks.ch API]: https://modpacksch.docs.apiary.io/
[LICENSE.txt]: ./LICENSE.txt
