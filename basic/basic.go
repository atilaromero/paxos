package basicpaxos

import (
	"fmt"
)

// Comm is passed to the Proposer. It takes care of the communication tasks in the protocol.
type Comm interface {
	// SendPrepare is a method that should construct and send the prepare message to all acceptors. The proposer does not care about the details of that implementation, it can be done using a REST API, grpc, or any other communication strategy.
	SendPrepare(v interface{})
}

// Proposer implements the proposer role of the basic paxos protocol. It must receive a Client that knows how to communicate with the acceptors. To use it, fill the Comm variable and call the Start method.
type Proposer struct {
	started bool
	v       interface{}
	Comm    Comm
}

// Start initiates the basic paxos protocol. It expects v, a value that can be the final choice of the protocol, if the other peers agree. The next step is to call Prepare. If called more than once, it returns an error.
func (p *Proposer) Start(v interface{}) error {
	if p.started {
		return fmt.Errorf("Start already called")
	}
	p.started = true
	p.v = v
	p.Prepare()
	return nil
}

// Prepare is the fist communication stage of the basic paxos protocol.
func (p *Proposer) Prepare() {
	p.Comm.SendPrepare(p.v)
}
