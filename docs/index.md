---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/ipinfo.svg"
brand_color: "#0095E5"
display_name: "ipinfo.io"
short_name: "ipinfo"
description: "Steampipe plugin to query IP address information from ipinfo.io."
og_description: "Query ipinfo.io with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/ipinfo-social-graphic.png"
---

# ipinfo.io + Steampipe

[ipinfo.io](https://ipinfo.com) is an API for IP address information (e.g. location).

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

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
+---------+---------+---------------+
| ip      | city    | country_name  |
+---------+---------+---------------+
| 1.2.3.4 | Chicago | United States |
+---------+---------+---------------+
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

## Documentation

- **[Table definitions & examples →](/plugins/turbot/ipinfo/tables)**

## Get started

### Install

Download and install the latest ipinfo.io plugin:

```bash
steampipe plugin install ipinfo
```

### Configuration

Installing the latest ipinfo plugin will create a config file (`~/.steampipe/config/ipinfo.spc`) with a single connection named `ipinfo`:

```hcl
connection "ipinfo" {
  plugin = "ipinfo"

  # Optional: Set your access token
  # No token is needed for basic info requests. For
  # higher limits and more data a token is required.
  # token = "4692efda8b2b56"
}
```

- `token` - Optional access token from ipinfo.io.

Environment variables are also available as an alternate configuration method:
* `IPINFO_TOKEN`

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-ipinfo
- Community: [Slack Channel](https://steampipe.io/community/join)
