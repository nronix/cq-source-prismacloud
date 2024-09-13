package client

type Prisma struct {
	ENDPOINT         string
	ACCOUNT          string
	API_KEY          string
	API_SECRET       string
	ALERT_POLICIES   []string
	INV_QUERY_STRING string
}
type Spec struct {
	// plugin spec goes here
	PRISMA []Prisma
}
