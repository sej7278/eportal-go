# eportal-go

A golang application that demonstrates how to use the [ePortal API](https://docs.tuxcare.com/eportal-api/)

## Running from go

You can run the source script directly from go, for example:

```bash
go run main.go --servers
```

## Building binaries

The included Makefile can build various binaries, subject to your go environment setup. The easiest invocation is:

```bash
make clean all
```

If you want to return in JSON format, append the `--json` flag and optionally pipe to `jq` like so:

```bash
eportal-go --users --json | jq
```

Which returns something like:

```json
{
  "result": [
    {
      "description": null,
      "id": 1,
      "readonly": false,
      "username": "admin"
    },
    {
      "description": "",
      "id": 2,
      "readonly": false,
      "username": "api"
    },
    {
      "description": "Readonly user",
      "id": 3,
      "readonly": true,
      "username": "user"
    }
  ]
}
```

## Setting up credentials

You should create a ~/.eportal.ini file with the following syntax:

```ini
username = api
password = Password123
url = https://eportal.example.com
```

You should give the file some level of protection using POSIX permissions, as the contents are not encrypted:

```bash
chmod 0600 ~/.eportal.ini
```

## Supported queries

As of the current release, eportal-go supports the following queries (API endpoints):

* users `--users`
* patchsets (defaults to main kernel feed) `--patchsets`
* keys `--keys`
* servers `--servers`
* feeds `--feeds`

They can be combined like so:

```bash
eportal-go --users --feeds
```

Which returns something like:

```text
FEEDS:
  Name: main
  Auto: false
  Channel: default
  DeployAfter: 0

  Name: all
  Auto: false
  Channel: default
  DeployAfter: 0

USERS:
  1: admin
  2: api
  3: user, Readonly user (readonly)
```
