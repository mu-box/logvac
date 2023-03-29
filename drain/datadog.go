package drain

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/DataDog/datadog-agent/pkg/logs/client/tcp"
	"github.com/DataDog/datadog-agent/pkg/logs/config"
	logvac "github.com/mu-box/logvac/core"
)

// Datadog drain implements the publisher interface for publishing logs to datadog.
type Datadog struct {
	ID      string // app ID or name
	Key     string // datadog api key
	manager *tcp.ConnectionManager
	conn    net.Conn
}

// NewDatadogClient creates a new mist publisher
func NewDatadogClient(id, key string) (*Datadog, error) {
	// emulate the server config
	config := config.Endpoint{
		APIKey: key,
		Host:   "intake.logs.datadoghq.com",
		Port:   10516,
		UseSSL: true,
	}

	// initialize a connection manager
	manager := tcp.NewConnectionManager(config)

	// establish a connection
	conn, err := manager.NewConnection(context.Background())

	return &Datadog{
		ID:      id,
		Key:     key,
		manager: manager,
		conn:    conn,
	}, err
}

// Init initializes a connection to mist
func (p Datadog) Init() error {

	// add drain
	logvac.AddDrain("datadog", p.Publish)

	return nil
}

// Publish utilizes mist's Publish to "drain" a log message
func (p *Datadog) Publish(msg logvac.Message) {
	// keep track of the attempts
	msg.PubTries++

	// re-establish the connection if it's been closed
	if p.conn == nil {
		p.conn, _ = p.manager.NewConnection(context.Background()) // doesn't block (don't goroutine call to Publish)
	}

	// generate the payload for this entry
	payload := formatDataDogMessage(msg, p.ID, p.Key)

	// send the payload
	_, err := p.conn.Write(payload)
	if err != nil {
		// it's possible the connection is bad, so let's close it
		p.manager.CloseConnection(p.conn)
		p.conn = nil

		// let's try to send it again (at least 3 times)
		if msg.PubTries <= 3 {
			// give it a sec, networks are fickle
			time.Sleep(2 * time.Second)
			// retry!
			p.Publish(msg)
		}
	}
}

// Close closes the connection to datadog.
func (p *Datadog) Close() error {
	p.manager.CloseConnection(p.conn)
	return nil
}

// format the syslog message, prefixed with the datadog api key
func formatDataDogMessage(msg logvac.Message, id, key string) []byte {
	// format the date in the proper syslog format
	date := msg.Time.Format(time.RFC3339)

	// prefix the app id to the hostname identifier
	hostname := fmt.Sprintf("%s.%s", id, msg.Id)

	// extract the first tag as the tag (ie: nginx[access])
	tag := msg.Tag[0]

	// the final message
	message := fmt.Sprintf("%s <%d>1 %s %s %s - - - %s\n",
		key, msg.Priority, date, hostname, tag, msg.Content)

	// return the message as a byte array
	return []byte(message)
}
