package operations

import (
	"github.com/behavioral-ai/collective/repository"
	"github.com/behavioral-ai/core/messaging"
)

const (
	NamespaceName = "collective:agent/operations/interact"
)

var (
	agents = messaging.NewExchange()
	agent  *agentT
)

type agentT struct{}

func init() {
	repository.RegisterConstructor(NamespaceName, func() messaging.Agent {
		return newAgent()
	})

}

func newAgent() *agentT {
	agent = new(agentT)
	return agent
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}

	if m.Event() == messaging.ShutdownEvent {
		agents.Broadcast(m)
	}
}

/*
func (a *agentT) configure(m *messaging.Message) {
	//ur := messaging.messaging.ConfigMapContent(m)
	//if cfg == nil {
	//	messaging.Reply(m, messaging.ConfigEmptyStatusError(a), a.Uri())
	//}
	// configure
	//messaging.Reply(m, messaging.StatusOK(), a.Uri())
	agents.Broadcast(m)
}


*/
