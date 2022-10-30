# Table: ipinfo_ip

Get location and other information about an IP address.

## Examples

### Info for your IP address

```sql
select
  ip,
  city,
  region,
  postal,
  country_name
from
  ipinfo_ip
```

### Info for a specific IP address

```sql
select
  ip,
  city,
  region,
  postal,
  country_name
from
  ipinfo_ip
where
  ip = '1.1.1.1'
```

### Info for a collection of IP addresses

```sql
select
  ip,
  city,
  region,
  postal,
  country_name
from
  ipinfo_ip
where
  ip in ('1.1.1.1', '8.8.8.8')
```

### Info for IP addresses in a CSV file

Assume a CSV file called `my_ips.csv` made available by the CSV plugin:
```
address,provider
1.1.1.1,Cloudflare
8.8.8.8,Google
```

Then it can be joined with the `ipinfo_ip` table to gather information:

```sql
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
  ipinfo_ip.ip in (select ip from ips)
```

Unfortunately, this simpler query _does not work_ because the Postgres planner
tries to scan the ipinfo_ip table:
```sql
select
  ip,
  city,
  region,
  postal,
  country_name
from
  ipinfo_ip
where
  ipinfo_ip.ip in (select address::inet from my_ips)
```

Unfortunately, this simpler query _does not work_ because the Postgres planner
tries to scan the ipinfo_ip table:
```sql
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
  ipinfo_ip.ip = i.address
```
