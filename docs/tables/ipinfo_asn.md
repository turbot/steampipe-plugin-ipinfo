# Table: ipinfo_asn

Get detailed information about an Autonomous System Number (ASN) provider.

**NOTES**:
* This table requires a paid plan.
* Because ASNs cannot be listed in full, the `ipinfo_asn` table
  requires the `asn` field to be specified in all queries, defining the ASN
  to lookup.

## Examples

### Info about an ASN

```sql
select
  name,
  country_name,
  num_ips
from
  ipinfo_asn
where
  asn = 'AS7922'
```

### IPv4 prefixes managed by an ASN

```sql
select
  p ->> 'netblock' as netblock,
  p ->> 'name' as name
from
  ipinfo_asn,
  jsonb_array_elements(prefixes) as p
where
  asn = 'AS7922'
```
