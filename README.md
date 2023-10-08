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

If you want to return in JSON format, append the `--json` flag and optionally pipe to jq:

```bash
eportal-go --users --json | jq
```

## Setting up credentials

You should create a ~/eportal.ini file with the following syntax:

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

They can be combined e.g.

```bash
eportal-go --users --feeds
```
