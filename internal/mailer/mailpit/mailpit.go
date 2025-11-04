package mailpit

import (
	"context"
	"fmt"

	"github.com/DevKayoS/journey-go/internal/pgstore"
	"github.com/google/uuid"
	"github.com/wneessen/go-mail"
)

type store interface {
	GetTrip(ctx context.Context, tripID uuid.UUID) (pgstore.Trip, error)
}

type Mailpit struct {
	store store
}

func NewMailpit() Mailpit {
	return Mailpit{}
}

// TODO: adicionar contexto com timeout para possivel gargalo?
func (mp Mailpit) SendConfirmTripEmailToTripOwner(tripID uuid.UUID) error {
	ctx := context.Background()
	trip, err := mp.store.GetTrip(ctx, tripID)
	if err != nil {
		return fmt.Errorf("mailpit: failed to get trip for SendConfirmTripEmailToTripOwner: %w", err.Error())
	}

	// TODO: terminar servico de envio de email
	msg := mail.NewMsg()
}
