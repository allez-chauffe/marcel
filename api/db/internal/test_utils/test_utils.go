package test_utils

import (
	"os"

	"github.com/allez-chauffe/marcel/api/db"
	"github.com/allez-chauffe/marcel/config"
	"github.com/sirupsen/logrus"
)

func DatabaseTest(test func()) {
	os.Remove("/tmp-clients-test.db")
	config.SetDefault(config.New())
	config.Default().SetLogLevel(logrus.WarnLevel)
	config.Default().API().SetDBFile("/tmp/marcel-clients-test.db")
	db.Open()
	defer db.Close()
	defer os.Remove("/tmp/marcel-clients-test.db")
	test()
}

type Runner func(task func(tx *db.Tx) error) error
