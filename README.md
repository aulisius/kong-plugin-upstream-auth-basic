# Kong Upstream HTTP Basic Authentication (upstream-auth-basic)

[Kong](https://getkong.org) Plugin to add [HTTP Basic Authentication](https://tools.ietf.org/html/rfc2617#section-2) to the upstream request header.

Compatible with Kong 2.0

## Installation

1. The [Golang](https://golang.org) package manager must be installed.
2. [Kong](https://getkong.org) must be [Installed](https://getkong.org/install/) and you must be familiar with using and configuring Kong.
3. Install the module upstream-auth-basic.
```
go build -buildmode=plugin github.com/aulisius/kong-plugin-upstream-auth-basic
```
4. Add the custom plugin while starting Kong
```
-e "KONG_PLUGINS=bundled,upstream-auth-basic"
```
5. Restart kong

## Plugin Fields
The plugin accepts the following fields.

|Name    |Type|Required|Description                                                             |
|--------|----|--------|------------------------------------------------------------------------|
|username|text|true    |The username to send in the Authorization header to the upstream service|
|password|text|false   |The password to send in the Authorization header to the upstream service|
