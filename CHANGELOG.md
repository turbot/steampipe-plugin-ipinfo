## v0.2.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#9](https://github.com/turbot/steampipe-plugin-ipinfo/pull/9))

## v0.2.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#7](https://github.com/turbot/steampipe-plugin-ipinfo/pull/7))
- Recompiled plugin with Go version `1.21`. ([#7](https://github.com/turbot/steampipe-plugin-ipinfo/pull/7))

## v0.1.0 [2023-03-23]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#3](https://github.com/turbot/steampipe-plugin-ipinfo/pull/3))

## v0.0.2 [2022-11-10]

_Bug fixes_

- Fixed link to ipinfo.io in index document.

## v0.0.1 [2022-11-02]

_What's new?_

- New tables added
  - [ipinfo_asn](https://hub.steampipe.io/plugins/turbot/ipinfo/tables/ipinfo_asn)
  - [ipinfo_ip](https://hub.steampipe.io/plugins/turbot/ipinfo/tables/ipinfo_ip)
