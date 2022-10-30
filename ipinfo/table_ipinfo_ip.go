package ipinfo

import (
	"context"
	"net"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableIPInfoIP(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "ipinfo_ip",
		Description: "Information about an IP address.",
		List: &plugin.ListConfig{
			Hydrate: listIP,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "ip", Require: plugin.Optional, CacheMatch: "exact"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "ip", Type: proto.ColumnType_INET, Transform: transform.FromField("IP").Transform(ipToString), Description: "IP address to query."},
			{Name: "anycast", Type: proto.ColumnType_BOOL, Description: "True if this is an anycast IP."},
			{Name: "asn", Type: proto.ColumnType_JSON, Description: "Autonomous System Number (ASN) of the administrator for the IP, e.g. AS13335"},
			{Name: "bogon", Type: proto.ColumnType_BOOL, Description: "True if the IP is from a special use, non-public range. See https://ipinfo.io/bogon"},
			{Name: "city", Type: proto.ColumnType_STRING, Description: "City of the location, e.g. Los Angeles"},
			{Name: "country", Type: proto.ColumnType_STRING, Description: "Country code of the location, e.g. US"},
			{Name: "country_name", Type: proto.ColumnType_STRING, Description: "Country name of the location, e.g. United States"},
			{Name: "hostname", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "is_eu", Type: proto.ColumnType_BOOL, Description: "True if the IP address is located in the European Union (EU)."},
			{Name: "location", Type: proto.ColumnType_STRING, Description: "Geocode location, e.g. 34.0522,-118.2437"},
			{Name: "org", Type: proto.ColumnType_STRING, Description: "Name of the organization that manages the IP, e.g. AS13335 Cloudflare, Inc."},
			{Name: "postal", Type: proto.ColumnType_STRING, Description: "Postal code of the location, e.g. 90076"},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "Region of the location, e.g. California"},
			{Name: "timezone", Type: proto.ColumnType_STRING, Description: "Timezone of the location, e.g. America/Los_Angeles"},
			// Other columns
			{Name: "abuse", Type: proto.ColumnType_JSON, Description: "Abuse contact information. Paid plans only."},
			{Name: "carrier", Type: proto.ColumnType_JSON, Description: "Carrier contact information. Paid plans only."},
			{Name: "company", Type: proto.ColumnType_JSON, Description: "Company information. Paid plans only."},
			{Name: "domains", Type: proto.ColumnType_JSON, Description: "Domains associated with the IP. Pain plans only."},
			{Name: "privacy", Type: proto.ColumnType_JSON, Description: "Privacy data (e.g. VPC). Paid plans only."},
		},
	}
}

func listIP(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("ipinfo_ip.listIP", "connection_error", err)
		return nil, err
	}

	ipStr := ""
	if d.KeyColumnQuals["ip"] != nil {
		ipStr = d.KeyColumnQuals["ip"].GetInetValue().GetAddr()
		plugin.Logger(ctx).Debug("ipinfo_ip.listIP", "ipStr", ipStr)
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		plugin.Logger(ctx).Warn("ipinfo_ip.listIP", "invalid_ip", ip, "status", "defaulting to IP of caller")
	}

	details, err := conn.GetIPInfo(ip)
	plugin.Logger(ctx).Debug("ipinfo_ip.listIP", "details", details)
	if err != nil {
		plugin.Logger(ctx).Error("ipinfo_asn.listAsn", "query_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, details)

	return nil, nil
}
