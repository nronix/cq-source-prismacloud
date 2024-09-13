package client

import (
	"context"
	"fmt"
	"github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger
	Spec   Spec

	Account  PrismaAccount
	accounts *PrismaAccounts
}

func (c *Client) ID() string {
	// TODO: Change to either your plugin name or a unique dynamic identifier
	return fmt.Sprintf("PRISMACLOUD:%s", c.Account)
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s *Spec) (Client, error) {
	// TODO: Add your client initialization here
	var p_acounts = PrismaAccounts{}
	for _, acc := range s.PRISMA {
		client := &prismacloud.Client{}
		client.Protocol = "https"
		client.Url = acc.ENDPOINT

		client.Username = acc.API_KEY
		client.Password = acc.API_SECRET

		if err := client.Initialize(""); err != nil {
			logger.Err(err)
		}

		p_acounts.items = append(p_acounts.items, PrismaAccount{acc.ACCOUNT, acc.ALERT_POLICIES, acc.INV_QUERY_STRING, client})
	}

	c := Client{
		logger:   logger,
		Spec:     *s,
		accounts: &p_acounts,
	}

	return c, nil
}

func (c *Client) WithAccount(account PrismaAccount) *Client {
	newC := *c
	newC.logger = c.logger.With().Str("account", account.Account).Logger()
	newC.Account = account
	return &newC
}
