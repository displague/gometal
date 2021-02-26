// Code generated by go-swagger; DO NOT EDIT.

package v_l_a_ns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Command(ctx context.Context, client ClientService) *cobra.Command {
	command := cobra.Command{
		Use:   "v_l_a_ns",
		Short: "Client for v_l_a_ns",
	}

	command.AddCommand(deleteVirtualNetwork(ctx, client))
	command.AddCommand(getVirtualNetwork(ctx, client))

	return &command
}

func deleteVirtualNetwork(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &DeleteVirtualNetworkParams{}

	command := &cobra.Command{
		Use: "DeleteVirtualNetwork",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.DeleteVirtualNetwork(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	return command
}

func getVirtualNetwork(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &GetVirtualNetworkParams{}

	command := &cobra.Command{
		Use: "GetVirtualNetwork",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.GetVirtualNetwork(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	return command
}

func printPayload(value interface{}) error {
	bytes, err := json.MarshalIndent(value, "", "  ")

	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(bytes)
	return err
}

type dateTimeValue struct {
	v *strfmt.DateTime
}

func (dtv *dateTimeValue) String() string {
	return time.Time(*dtv.v).Format(strfmt.ISO8601LocalTime)
}

func (dtv *dateTimeValue) Set(flag string) error {
	return dtv.v.UnmarshalText([]byte(flag))
}

func (dtv *dateTimeValue) Type() string {
	return "dateTime"
}

func newDateTimeValue(v *strfmt.DateTime) pflag.Value {
	return &dateTimeValue{v}
}