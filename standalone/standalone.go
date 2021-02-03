package standalone

import (
	"github.com/allez-chauffe/marcel/api"
	"github.com/allez-chauffe/marcel/backoffice"
	"github.com/allez-chauffe/marcel/frontend"
	"github.com/allez-chauffe/marcel/module"
)

// Module creates the standalone module.
func Module() *module.Module {
	return &module.Module{
		Name: "Standalone",
		SubModules: []*module.Module{
			api.Module(),
			backoffice.Module(),
			frontend.Module(),
		},
	}
}
