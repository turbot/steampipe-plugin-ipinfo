![image](https://hub.steampipe.io/images/plugins/turbot/ipinfo-social-graphic.png)

# ipinfo.io Plugin for Steampipe

Use SQL to IP address information using [ipinfo.io](https://ipinfo.io).

- **[Get started →](https://hub.steampipe.io/plugins/turbot/ipinfo)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/ipinfo/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-ipinfo/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install ipinfo
```

Configure the server address in `~/.steampipe/config/ipinfo.spc`:

```hcl
connection "ipinfo" {
  plugin = "ipinfo"
  # Optional token for paid plans
  # token = "4692efda8b2b56"
}
```

Run steampipe:

```shell
steampipe query
```

Query your own IP:

```sql
select
  ip,
  city,
  country_name
from
  ipinfo_ip;
```

```
+---------+-------------+---------------+
| ip      | city        | country_name  |
+---------+-------------+---------------+
| 1.1.1.1 | Los Angeles | United States |
+---------+-------------+---------------+
```

Or query any IP:

```sql
select
  ip,
  city,
  country_name
from
  ipinfo_ip
where
  ip = '8.8.8.8';
```

```
+---------+---------------+---------------+
| ip      | city          | country_name  |
+---------+---------------+---------------+
| 8.8.8.8 | Mountain View | United States |
+---------+---------------+---------------+
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-ipinfo.git
cd steampipe-plugin-ipinfo
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/ipinfo.spc
```

Try it!

```
steampipe query
> .inspect ipinfo
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-ipinfo/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-ipinfo/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [ipinfo.io Plugin](https://github.com/turbot/steampipe-plugin-ipinfo/labels/help%20wanted)
