package sshtunnel

import (
	"crayontool-go/pkg/constant"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
	"sync"
)

type event string

const (
	eventStop event = "stop"
)

const (
	copyBufferSize = 1 << 6
)

type Client struct {
	Config      *Config
	netListener net.Listener
	eventChan   chan event
}

type connManager struct {
	localConn  net.Conn
	sshConn    *ssh.Client
	remoteConn net.Conn
}

func NewClient(config *Config) Client {
	return Client{
		Config:    config,
		eventChan: make(chan event, 1),
	}
}

func (c *Client) initConnManager(localConn net.Conn) (*connManager, error) {
	sshClient, err := ssh.Dial(constant.TCP, c.Config.SSHTunnel.GetAddrStr(), &c.Config.sshConfig)
	if err != nil {
		fmt.Printf("ssh dial error, err: %v\n", err)
		return nil, err
	}
	remoteConn, err := sshClient.Dial(constant.TCP, c.Config.Remote.GetAddrStr())
	if err != nil {
		fmt.Printf("remote dial error, err: %v\n", err)
		sshClient.Close()
		return nil, err
	}
	return &connManager{
		localConn:  localConn,
		sshConn:    sshClient,
		remoteConn: remoteConn,
	}, nil
}

func (m *connManager) close() {
	if m.sshConn != nil {
		m.sshConn.Close()
	}
	if m.remoteConn != nil {
		m.remoteConn.Close()
	}
	if m.localConn != nil {
		m.localConn.Close()
	}
}

func (m *connManager) dataTransfer() {
	wg := sync.WaitGroup{}
	c1 := make(chan struct{}, 1)
	c2 := make(chan struct{}, 1)
	// 有一个出现异常就退出
	// 将localConn.Reader复制到remoteConn.Writer
	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			//case <-c1:
			//	return
			default:
				_, err := io.CopyN(m.remoteConn, m.localConn, copyBufferSize)
				if err != nil {
					fmt.Printf("io copy (local -> remote) error: %v\n", err)
					c2 <- struct{}{}
					return
				}
			}
		}
	}()
	// remoteConn.Reader复制到localConn.Writer
	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			//case <-c2:
			//	return
			default:
				_, err := io.CopyN(m.localConn, m.remoteConn, copyBufferSize)
				if err != nil {
					fmt.Printf("io copy (remote -> local) error: %v\n", err)
					c1 <- struct{}{}
					return
				}
			}
		}
	}()
	wg.Wait()
	//m.close()
}

func (c *Client) listenSignal() {
	go func() {
		select {
		case signal := <-c.eventChan:
			switch signal {
			case eventStop:
				c.doStop()
				return
			}
		}
	}()
}

func (c *Client) initListener() error {
	listener, err := net.Listen(constant.TCP, c.Config.Local.GetAddrStr())
	if err != nil {
		return err
	}
	c.netListener = listener
	return nil
}

func (c *Client) acceptConn() {
	for {
		localConn, err := c.netListener.Accept()
		if err != nil {
			fmt.Printf("accept error, err: %v\n", err)
			break
		}
		manager, err := c.initConnManager(localConn)
		if err != nil {
			fmt.Printf("init conn error, err: %v\n", err)
			continue
		}
		go manager.dataTransfer()
	}
}

func (c *Client) sendEventSignal(e event) {
	c.eventChan <- e
}

func (c *Client) doStop() {
	c.netListener.Close()
}

func (c *Client) Start() error {
	err := c.initListener()
	if err != nil {
		return err
	}
	c.listenSignal()
	c.acceptConn()
	return nil
}

func (c *Client) Stop() {
	c.sendEventSignal(eventStop)
}
