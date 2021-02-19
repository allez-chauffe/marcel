package standalone

import (
	"github.com/allez-chauffe/marcel/api"
	"github.com/allez-chauffe/marcel/pkg/backoffice"
	"github.com/allez-chauffe/marcel/pkg/frontend"
	"github.com/allez-chauffe/marcel/pkg/module"
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
