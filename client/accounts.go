package client

import prismacloud "github.com/paloaltonetworks/prisma-cloud-go"

type PrismaAccount struct {
	Account          string
	ALERT_POLICIES   []string
	INV_QUERY_STRING string
	PsClient         *prismacloud.Client
}

type PrismaAccounts struct {
	items []PrismaAccount
}
