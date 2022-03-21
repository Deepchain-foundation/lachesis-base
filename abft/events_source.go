package abft

import (
	"github.com/Deepchain-foundation/lachesis-base/hash"
	"github.com/Deepchain-foundation/lachesis-base/inter/dag"
)

// EventSource is a callback for getting events from an external storage.
type EventSource interface {
	HasEvent(hash.Event) bool
	GetEvent(hash.Event) dag.Event
}
