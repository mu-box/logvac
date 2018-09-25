package drain

import (
	"fmt"
	"io"
	"net"

	"github.com/nanopack/logvac/config"
	"github.com/nanopack/logvac/core"
)

// Papertrail drain implements the publisher interface for publishing logs to papertrail.
type Papertrail struct {
	ID		string				 // the app id or name
	Conn 	io.WriteCloser // connection to forward logs through
}

// NewPapertrailClient creates a new mist publisher
func NewPapertrailClient(uri, id string) (*Papertrail, error) {
	addr, err := net.ResolveUDPAddr("udp", uri)
	if err != nil {
		return nil, fmt.Errorf("Failed to resolve papertrail address - %s", err.Error())
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial papertrail - %s", err.Error())
	}
	
	config.Log.Info("Connection to papertrail endpoint established")

	return &Papertrail{Conn: conn, ID: id}, nil
}

// Init initializes a connection to mist
func (p *Papertrail) Init() error {

	// add drain
	logvac.AddDrain("papertrail", p.Publish)

	return nil
}

// Publish utilizes mist's Publish to "drain" a log message
func (p *Papertrail) Publish(msg logvac.Message) {
	config.Log.Info("%v", msg)
	p.Conn.Write(msg.Raw)
}

// Close closes the connection to papertrail.
func (p *Papertrail) Close() error {
	if p.Conn == nil {
		return nil
	}
	return p.Conn.Close()
}
