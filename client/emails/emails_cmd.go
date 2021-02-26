// Code generated by go-swagger; DO NOT EDIT.

package emails

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
		Use:   "emails",
		Short: "Client for emails",
	}

	command.AddCommand(createEmail(ctx, client))
	command.AddCommand(deleteEmail(ctx, client))
	command.AddCommand(findEmailById(ctx, client))
	command.AddCommand(updateEmail(ctx, client))

	return &command
}

func createEmail(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &CreateEmailParams{}

	command := &cobra.Command{
		Use: "CreateEmail",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.CreateEmail(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// types.CreateEmailInput

	return command
}

func deleteEmail(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &DeleteEmailParams{}

	command := &cobra.Command{
		Use: "DeleteEmail",
		RunE: func(cmd *cobra.Command, args []string) error {

			_, err := client.DeleteEmail(parameters, nil)
			return err

		},
	}

	parameters.Context = ctx

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	return command
}

func findEmailById(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &FindEmailByIDParams{}

	command := &cobra.Command{
		Use: "FindEmailByID",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.FindEmailByID(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// strfmt.UUID

	command.MarkPersistentFlagRequired("id")

	// string

	parameters.Include = command.PersistentFlags().String("include", "", "related attributes to include")

	return command
}

func updateEmail(ctx context.Context, client ClientService) *cobra.Command {
	parameters := &UpdateEmailParams{}

	command := &cobra.Command{
		Use: "UpdateEmail",
		RunE: func(cmd *cobra.Command, args []string) error {

			result, err := client.UpdateEmail(parameters, nil)
			if err != nil {
				return err
			}
			return printPayload(result.Payload)

		},
	}

	parameters.Context = ctx

	// types.UpdateEmailInput

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