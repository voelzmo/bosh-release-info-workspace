package output

import (
	"os"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

func Fail(err error, logger boshlog.Logger) {
	logger.Error("main", err.Error())
	os.Exit(1)
}
