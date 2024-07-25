//go:generate go run github.com/vektra/mockery/v2
package outbox

import (
	"context"
	"database/sql"

	"github.com/tink-crypto/tink-go/v2/tink"
)

type AEADFunc func(ob Outbox[any]) (tink.AEAD, error)

type Relay func(ctx context.Context, obs []Outbox[Serialized]) error

type Manager interface {
	StoreOutbox(ctx context.Context, tx *sql.Tx, ob Outbox[any]) (err error)
	StoreOutboxEncrypted(ctx context.Context, tx *sql.Tx, ob Outbox[any]) (err error)

	ObserveOutboxes(ctx context.Context, relay Relay) (err error)
}
