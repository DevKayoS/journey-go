package mailpit

import (
	"context"
	"fmt"
	"time"

	"github.com/DevKayoS/journey-go/internal/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wneessen/go-mail"
)

type store interface {
	GetTrip(ctx context.Context, tripID uuid.UUID) (pgstore.Trip, error)
}

type Mailpit struct {
	store store
}

func NewMailpit(pool *pgxpool.Pool) Mailpit {
	return Mailpit{
		store: pgstore.New(pool),
	}
}

func (mp Mailpit) SendConfirmTripEmailToTripOwner(tripID uuid.UUID) error {
	ctx := context.Background()
	trip, err := mp.store.GetTrip(ctx, tripID)
	if err != nil {
		return fmt.Errorf("mailpit: failed to get trip for SendConfirmTripEmailToTripOwner: %w", err.Error())
	}

	msg := mail.NewMsg()

	if err := msg.From("mailpit@journey.com"); err != nil {
		return fmt.Errorf("mailpit: failed to set From in email SendConfirmTripEmailToTripOwner: %w", err.Error())
	}

	if err := msg.To(trip.OwnerEmail); err != nil {
		return fmt.Errorf("mailpit: failed to set To in email SendConfirmTripEmailToTripOwner: %w", err.Error())
	}

	msg.Subject("Confirme sua viagem")

	msg.SetBodyString(mail.TypeTextPlain, fmt.Sprintf(`
		Olá, %s! 🌴✈️

		Temos uma notícia incrível para você! Sua viagem para %s, esta agendada para começar no dia %s

		Você foi **convidado para participar de uma viagem especial**! 🧳✨  
		Prepare-se para viver momentos inesquecíveis, relaxar e se divertir em grande estilo.

		Em breve, entraremos em contato com mais detalhes sobre o destino, datas e tudo o que você precisa saber para embarcar nessa experiência única.

		Enquanto isso, já pode começar a sonhar com essa nova aventura! 😄

		Atenciosamente,  
		Equipe de Viagens 🌍
		`,
		trip.OwnerName,
		trip.Destination,
		trip.StartsAt.Time.Format(time.DateOnly),
	))

	client, err := mail.NewClient(
		"localhost",
		mail.WithTLSPortPolicy(mail.NoTLS),
		mail.WithPort(1025),
	)
	if err != nil {
		return fmt.Errorf("mailpit: failed to create client SendConfirmTripEmailToTripOwner: %w", err.Error())
	}

	if err := client.DialAndSend(msg); err != nil {
		return fmt.Errorf("mailpit: failed to create client SendConfirmTripEmailToTripOwner: %w", err.Error())
	}

	return nil
}
