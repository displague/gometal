// Code generated by go-swagger; DO NOT EDIT.

package userdata

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
		Use:   "userdata",
		Short: "Client for userdata",
	}

	command.AddCommand(validateUserdata(ctx, client))

	return &command
}

func validateUserdata(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &ValidateUserdataParams{}

	command := &cobra.Command{
		Use: "ValidateUserdata",
		RunE: func(cmd *cobra.Command, args []string) error {

			_, err := client.ValidateUserdata(parameters, nil)
			return err

		},
	}

	parameters.Context = ctx

	// string

	parameters.Userdata = command.PersistentFlags().String("userdata", "", "Userdata to validate")

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