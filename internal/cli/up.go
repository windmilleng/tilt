package cli

import (
	"context"
	"errors"
	"github.com/spf13/cobra"
	"github.com/windmilleng/tilt/internal/tiltd/tiltd_client"
	"github.com/windmilleng/tilt/internal/tiltd/tiltd_server"
	"github.com/windmilleng/tilt/internal/tiltfile"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type upCmd struct{}

func (c *upCmd) register() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up <servicename>",
		Short: "stand up a service",
		Args:  cobra.ExactArgs(1),
	}

	return cmd
}

func foo(args []string) error {
	ctx := context.Background()
	proc, err := tiltd_server.RunDaemon(ctx)
	if err != nil {
		return err
	}
	defer proc.Kill()

	dCli, err := tiltd_client.NewDaemonClient()
	if err != nil {
		return err
	}

	tf, err := tiltfile.Load("Tiltfile")
	if err != nil {
		return err
	}

	serviceName := args[0]
	service, err := tf.GetServiceConfig(serviceName)
	if err != nil {
		return err
	}

	err = dCli.CreateService(ctx, *service)
	s, ok := status.FromError(err)
	if ok && s.Code() == codes.Unknown {
		return errors.New(s.Message())
	}

	return nil
}

func (c *upCmd) run(args []string) error {
	err := foo(args)
	if err != nil {
		// log the error and exit so that cobra doesn't print usage
		log.Fatalf("Error: %v", err)
	}

	return nil
}
