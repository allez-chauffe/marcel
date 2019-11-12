package standalone

import (
	"github.com/Zenika/marcel/api"
	"github.com/Zenika/marcel/backoffice"
	"github.com/Zenika/marcel/frontend"
	"github.com/Zenika/marcel/module"
)

func Module() module.Module {
	return module.Module{
		Name: "Standalone",
		SubModules: []module.Module{
			api.Module(),
			backoffice.Module(),
			frontend.Module(),
		},
	}
}
