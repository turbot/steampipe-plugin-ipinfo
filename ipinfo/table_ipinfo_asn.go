package ipinfo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableIPInfoAsn(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "ipinfo_asn",
		Description: "Details about an Autonomous System Number (ASN).",
		List: &plugin.ListConfig{
			Hydrate: listAsn,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "asn"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "asn", Type: proto.ColumnType_STRING, Description: "Autonomous System Number (ASN) of the administrator for the IP, e.g. AS13335"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the ASN."},
			{Name: "country_name", Type: proto.ColumnType_STRING, Description: "Country name of the ASN, e.g. United States"},
			{Name: "num_ips", Type: proto.ColumnType_INT, Description: "Number of IPs in the ASN, e.g. 71224576"},
			// Other columns
			{Name: "allocated", Type: proto.ColumnType_STRING, Description: "Date when the ASN was allocated, e.g. 1997-02-14"},
			{Name: "country", Type: proto.ColumnType_STRING, Description: "Country code of the ASN, e.g. US"},
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "Domain of the ASN, e.g. comcast.com"},
			{Name: "downstreams", Type: proto.ColumnType_JSON, Description: "Downstream information for the ASN."},
			{Name: "peers", Type: proto.ColumnType_JSON, Description: "Peers for the ASN."},
			{Name: "prefixes", Type: proto.ColumnType_JSON, Description: "List of objects representing IPv4 ranges for the ASN."},
			{Name: "prefixes6", Type: proto.ColumnType_JSON, Transform: transform.FromField("Prefixes6"), Description: "List of objects representing IPv6 ranges for the ASN."},
			{Name: "registry", Type: proto.ColumnType_STRING, Description: "Registry for the ASN, e.g. arin"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the ASN, e.g. isp"},
			{Name: "upstreams", Type: proto.ColumnType_JSON, Description: "Upstream information for the ASN."},
		},
	}
}

func listAsn(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("ipinfo_asn.listAsn", "connection_error", err)
		return nil, err
	}

	asn := d.EqualsQuals["asn"].GetStringValue()
	plugin.Logger(ctx).Debug("ipinfo_asn.listAsn", "asn", asn)

	details, err := conn.GetASNDetails(asn)
	plugin.Logger(ctx).Debug("ipinfo_asn.listAsn", "details", details)
	if err != nil {
		plugin.Logger(ctx).Error("ipinfo_asn.listAsn", "query_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, details)

	return nil, nil
}
