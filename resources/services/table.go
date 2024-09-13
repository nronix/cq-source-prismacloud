package services

import (
	"context"
	//"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/nronix/cq-source-prismacloud/client"
	"github.com/nronix/cq-source-prismacloud/internal/alerts"
	"github.com/nronix/cq-source-prismacloud/internal/ec2"
)

func Ec2InventoryTable() *schema.Table {
	return &schema.Table{
		Name:      "prismacloud_ec2_inventory",
		Resolver:  fetchEc2InventoryTable,
		Transform: transformers.TransformWithStruct(&ec2.Ec2Inventory{}),
	}
}

func AlertsTable() *schema.Table {
	return &schema.Table{
		Name:      "prismacloud_alerts",
		Resolver:  fetchAlertsTables,
		Transform: transformers.TransformWithStruct(&alerts.Alerts{}),
		Multiplex: client.AccountMultiplex,
	}
}

func fetchEc2InventoryTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	response, err := ec2.GetEc2Inventry(cl)
	if err != nil {
		return err
	}
	res <- response
	return nil
}

func fetchAlertsTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	//account := cl.Account
	response, err := alerts.GetAlerts(cl)
	if err != nil {
		return err
	}
	res <- response
	return nil
}
