package config

import (
	"net"

	"github.com/labstack/echo/v4"
)

type ServerConfig struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Debug   bool   `yaml:"debug"`
	URL     string `yaml:"url"`
	RootDir string `yaml:"rootdir"`
	// ProxyMode is either "direct" (default), "xff", "real-ip"
	ProxyMode string `yaml:"proxy-mode"`
	// TrustedProxyCIDR is a list of CIDR ranges to trust
	TrustedProxyCIDR []string `yaml:"trusted-proxy-cidr"`
	// TrustedProxyLoopback is a flag to trust loopback addresses
	TrustedProxyLoopback bool `yaml:"trusted-proxy-loopback"`
	// TrustedProxyLocalLink is a flag to trust local link addresses like 169.254.*.*
	TrustedProxyLocallink bool `yaml:"trusted-proxy-local-link"`
	// TrustedProxyPrivateNet is a flag to trust private network addresses like 192.168.*.* or 10.*.*.*
	TrustedProxyPrivateNet bool `yaml:"trusted-proxy-private-net"`
}

var trustedProxyOptions []echo.TrustOption = nil

func (s *ServerConfig) GetTrustedProxyOptions() []echo.TrustOption {
	if trustedProxyOptions == nil {
		srvOptions := s
		trustOptions := make([]echo.TrustOption, 0, len(srvOptions.TrustedProxyCIDR))
		for _, cidr := range srvOptions.TrustedProxyCIDR {
			if _, parsedCIDR, err := net.ParseCIDR(cidr); err == nil {
				trustOptions = append(trustOptions, echo.TrustIPRange(parsedCIDR))
			}
		}
		trustOptions = append(trustOptions, echo.TrustLoopback(srvOptions.TrustedProxyLoopback))
		trustOptions = append(trustOptions, echo.TrustLinkLocal(srvOptions.TrustedProxyLocallink))
		trustOptions = append(trustOptions, echo.TrustPrivateNet(srvOptions.TrustedProxyPrivateNet))
		trustedProxyOptions = trustOptions
	}
	return trustedProxyOptions
}
