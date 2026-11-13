package ipinfo

import (
	"context"
	"net"
	"os"

	"github.com/ipinfo/go/v2/ipinfo"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/transform"
)

func connect(ctx context.Context, d *plugin.QueryData) (*ipinfo.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "ipinfo"
	if cachedData, ok := d.ConnectionCache.Get(ctx, cacheKey); ok {
		return cachedData.(*ipinfo.Client), nil
	}

	// Default to the env var settings
	token := os.Getenv("IPINFO_TOKEN")

	// Prefer config settings
	ipinfoConfig := GetConfig(d.Connection)
	if ipinfoConfig.Token != nil {
		token = *ipinfoConfig.Token
	}

	conn := ipinfo.NewClient(nil, nil, token)

	// Save to cache
	_ = d.ConnectionCache.Set(ctx, cacheKey, conn)

	return conn, nil
}

func ipToString(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Debug("ipToString", "d.Value", d.Value)
	if d.Value == nil {
		return nil, nil
	}
	ip := d.Value.(net.IP)
	plugin.Logger(ctx).Debug("ipToString", "ip", ip)
	return ip.String(), nil
}
