# windhager-biowint2-client-go

Windhager BioWinTouch 2 goes cloud native with Go.

## Accessing the UI

The UI is served by the Windhager pellet appliance built-in MES inifinity controller. Behind the UI, a REST API serves all values that are avaialable through the ui or the touch panel. More information about the mes infinity controller can be found here: [https://www.windhager.com/int_en/products/control/mes-infinity/](https://www.windhager.com/int_en/products/control/mes-infinity/)

## Installation

go-github is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/google/go-github/v39
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/google/go-github/v39/github"
```

and run `go get` without parameters.

Finally, to use the top-of-trunk version of this repo, use the following command:

```bash
go get github.com/google/go-github/v39@master
```

## Usage

```go
import "github.com/lukibahr/windhager-biowin2t-client-go" // with go modules enabled (GO111MODULE=on or outside GOPATH)
```

Construct a new `NewWindhagerClient` client, then use the various functions on the client to
access different parts of the Windhager BioWin2T API. Make sure to pass the required Url, Username and Password to the client. For example:

```go
client := github.NewWindhagerClient("url", "username", "password")
```

Some API methods have optional parameters that can be passed. For example:

NOTE: Using the [context](https://godoc.org/context) package, one can easily
pass cancelation signals and deadlines to various services of the client for
handling a request. In case there is no context available, then `context.Background()`
can be used as a starting point.

## Curling the API

The api endpoint is available under `http://<your-host>/api/1.0/lookup/<OID>`. To add the digest authentication, use the `--digest` parameter like the following: `curl http://<your-host>/api/1.0/lookup/<OID> --digest -u "$USERNAME:$PASSWORD"`

Sample response looks like the following:

```bash
{
    "OID": "/1/60/0/12/101/0",
    "groupNr": 12,
    "maxValue": "14.0",
    "memberNr": 101,
    "minValue": "6.0",
    "name": "12-101",
    "step": "0.1",
    "stepId": 0,
    "subtypeId": -1,
    "timestamp": "2021-09-07 11:23:03",
    "typeId": 15,
    "unit": "kg",
    "unitId": 45,
    "value": "6.0",
    "writeProt": false
}
```
