package aws_secrets_manager

import (
	"fmt"
        "context"
        "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go-v2/config"
	"os"
)

func Auth(profile string) *session.Session {

        ctx := context.Background()
       
        shared_config, err := config.LoadSharedConfigProfile(ctx, profile)

	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,

                Config: aws.Config {
                  Region: aws.String(shared_config.Region),
                },
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}

	return sess
}
