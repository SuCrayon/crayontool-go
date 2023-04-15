package sshtunnel

import (
	"github.com/SuCrayon/crayontool-go/pkg/types"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

const (
	// defaultIP 默认绑定机器上的所有可用IP
	defaultIP = "0.0.0.0"
	// defaultPort 默认0号端口，由操作系统自动分配可用端口
	defaultPort = 0
	// defaultTimeout 默认10秒超时时间
	defaultTimeout = 10 * time.Second
)

type Config struct {
	sshConfig ssh.ClientConfig
	SSHTunnel types.Addr
	Remote    types.Addr
	Local     types.Addr
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) SetSSHTunnelIP(ip string) *Config {
	c.SSHTunnel.IP = ip
	return c
}

func (c *Config) SetSSHTunnelPort(port uint16) *Config {
	c.SSHTunnel.Port = port
	return c
}

func (c *Config) SetLocalIP(ip string) *Config {
	c.Local.IP = ip
	return c
}

func (c *Config) SetLocalPort(port uint16) *Config {
	c.Local.Port = port
	return c
}

func (c *Config) SetRemoteIP(ip string) *Config {
	c.Remote.IP = ip
	return c
}

func (c *Config) SetRemotePort(port uint16) *Config {
	c.Remote.Port = port
	return c
}

func (c *Config) SetUser(user string) *Config {
	c.sshConfig.User = user
	return c
}

func (c *Config) SetPassword(password string) *Config {
	c.sshConfig.Auth = []ssh.AuthMethod{
		ssh.Password(password),
	}
	return c
}

func (c *Config) SetTimeout(timeout time.Duration) *Config {
	c.sshConfig.Timeout = timeout
	return c
}

func (c *Config) SetHostKeyCallback(hostKeyCallback ssh.HostKeyCallback) *Config {
	c.sshConfig.HostKeyCallback = hostKeyCallback
	return c
}

func (c *Config) SetBannerCallback(bannerCallback ssh.BannerCallback) *Config {
	c.sshConfig.BannerCallback = bannerCallback
	return c
}

func (c *Config) SetClientVersion(clientVersion string) *Config {
	c.sshConfig.ClientVersion = clientVersion
	return c
}

func (c *Config) SetHostKeyAlgorithms(hostKeyAlgorithms []string) *Config {
	c.sshConfig.HostKeyAlgorithms = hostKeyAlgorithms
	return c
}

func NewDefaultConfig() *Config {
	return NewConfig().SetLocalIP(defaultIP).SetLocalPort(defaultPort).SetHostKeyCallback(
		func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	).SetTimeout(defaultTimeout)
}
