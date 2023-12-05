---
title: "Steampipe Table: ipinfo_asn - Query IPinfo ASN using SQL"
description: "Allows users to query Autonomous System Numbers (ASNs) in IPinfo, providing insights into network routing and traffic patterns."
---

# Table: ipinfo_asn - Query IPinfo ASN using SQL

Autonomous System Numbers (ASNs) are a crucial part of the internet's routing infrastructure. They are used by internet service providers and large organizations to manage their network routes and traffic. With ASNs, users can understand how internet traffic is routed and identify the networks responsible for specific internet activities.

## Table Usage Guide

The `ipinfo_asn` table provides detailed insights into Autonomous System Numbers (ASNs) within IPinfo. As a network engineer or IT professional, explore ASN-specific details through this table, including the network prefix, name, country, and associated metadata. Utilize it to uncover information about network routing, such as the routes used by an internet service provider or a large organization, and the traffic patterns of these networks.

**Important Notes**
- This table requires a paid plan.
- Because ASNs cannot be listed in full, the `ipinfo_asn` table requires the `asn` field to be specified in all queries, defining the ASN to lookup.

## Examples

### Info about an ASN
Gain insights into the specifics of a particular Autonomous System Number (ASN) including its name, country of origin, and the number of IP addresses it has. This is useful for network administrators and cybersecurity professionals to understand the scale and geographical location of a specific ASN.

```sql
select
  name,
  country_name,
  num_ips
from
  ipinfo_asn
where
  asn = 'AS7922';
```

### IPv4 prefixes managed by an ASN
Explore which IPv4 prefixes are managed by a specific Autonomous System Number (ASN). This can be useful to understand the scope of networks managed by an ASN and potentially identify patterns or anomalies.

```sql
select
  p ->> 'netblock' as netblock,
  p ->> 'name' as name
from
  ipinfo_asn,
  jsonb_array_elements(prefixes) as p
where
  asn = 'AS7922';
```