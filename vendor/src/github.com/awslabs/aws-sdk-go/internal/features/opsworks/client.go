package opsworks

import (
	"github.com/awslabs/aws-sdk-go/internal/features/shared"
	"github.com/awslabs/aws-sdk-go/service/opsworks"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@opsworks", func() {
		World["client"] = opsworks.New(nil)
	})
}
