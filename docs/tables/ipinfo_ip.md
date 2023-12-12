---
title: "Steampipe Table: ipinfo_ip - Query IPinfo IP Addresses using SQL"
description: "Allows users to query IP Addresses in IPinfo, specifically providing insights into the geolocation, ISP, and other related information of an IP address."
---

# Table: ipinfo_ip - Query IPinfo IP Addresses using SQL

IPinfo is a comprehensive IP data provider, offering data on IP addresses including geolocation, ASN, ISP, and more. This service allows you to gain insights into where your users are coming from and how your app's performance varies by region, ISP, and other factors. It's an invaluable resource for businesses that need to understand their online traffic and optimize their user experience.

## Table Usage Guide

The `ipinfo_ip` table provides insights into IP addresses data within IPinfo. As a network administrator or cybersecurity analyst, explore IP-specific details through this table, including geolocation, ISP, and associated metadata. Utilize it to uncover information about IP addresses, such as their geographical location, the ISP providing the service, and other pertinent details.

## Examples

### Info for your IP address
Discover the geographical details associated with your IP address, such as city, region, postal code, and country. This can be useful for understanding the location data linked to your online activities.

```sql+postgres
select
  ip,
  city,
  region,
  postal,
  country_name
from
  ipinfo_ip;
```

```sql+sqlite
select
  ip,
  city,
  region,
  postal,
  country_name
from
  ipinfo_ip;
```

### Info for a specific IP address
Explore which city, region, and country a specific IP address is associated with. This can be useful for understanding the geographic distribution of your users or for detecting potentially suspicious activity.

```sql+postgres
select
  ip,
  city,
  region,
  postal,
  country_name
from
  ipinfo_ip
where
  ip = '1.1.1.1';
```

```sql+sqlite
select
  ip,
  city,
  region,
  postal,
  country_name
from
  ipinfo_ip
where
  ip = '1.1.1.1';
```

### Info for a collection of IP addresses
Explore the geographical locations associated with a set of IP addresses. This query is useful for understanding the distribution of users or network requests across different regions.

```sql+postgres
select
  ip,
  city,
  region,
  postal,
  country_name
from
  ipinfo_ip
where
  ip in ('1.1.1.1', '8.8.8.8');
```

```sql+sqlite
select
  ip,
  city,
  region,
  postal,
  country_name
from
  ipinfo_ip
where
  ip in ('1.1.1.1', '8.8.8.8');
```

### Info for IP addresses in a CSV file
This query is useful for identifying geographical information related to a set of IP addresses. It can be used to gain insights into where certain IP addresses are based, including details about their city, region, and country, which can be beneficial for understanding user demographics or tracking potential security threats.
Assume a CSV file called `my_ips.csv` made available by the CSV plugin:
```
address,provider
1.1.1.1,Cloudflare
8.8.8.8,Google
```

Then it can be joined with the `ipinfo_ip` table to gather information:


```sql+postgres
-- Query IPs from the CSV first, to force the planner
with ips as (
  select
    address::inet as ip
  from
    my_ips
  order by
    ip
)
select
  ip,
  city,
  region,
  postal,
  country_name
from
  ipinfo_ip
where
  ipinfo_ip.ip in (select ip from ips);
```

```sql+sqlite
Error: SQLite does not support CIDR operations.
```

Unfortunately, this simpler query _does not work_ because the Postgres planner
tries to scan the ipinfo_ip table:
```sql+postgres
select
  ip,
  city,
  region,
  postal,
  country_name
from
  ipinfo_ip
where
  ipinfo_ip.ip in (select address::inet from my_ips);
```

```sql+sqlite
Error: SQLite does not support CIDR operations.
```

Unfortunately, this simpler query _does not work_ because the Postgres planner
tries to scan the ipinfo_ip table:
```sql+postgres
select
  ip,
  city,
  region,
  postal,
  country_name
from
  csv.my_ips as i,
  ipinfo_ip
where
  ipinfo_ip.ip = i.address;
```

```sql+sqlite
select
  ip,
  city,
  region,
  postal,
  country_name
from
  csv_my_ips as i,
  ipinfo_ip
where
  ipinfo_ip.ip = i.address;
```