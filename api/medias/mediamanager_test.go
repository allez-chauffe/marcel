package medias

import (
	"testing"

	"github.com/Zenika/MARCEL/api/clients"
	"github.com/Zenika/MARCEL/api/plugins"
)

func TestNewManager(t *testing.T) {
	m := NewManager(plugins.NewService().GetManager(), clients.NewService(), MEDIAS_CONFIG_PATH, MEDIAS_CONFIG_FILENAME)

	if m.Config == nil {
		t.Error("Configuration should not be nil")
	}

	if m.Config.LastID != 0 {
		t.Error("LastID for a new Media Manager configuration should be 0")
	}
}

func TestManager_GetPortNumberForPlugin(t *testing.T) {
	m := NewManager(plugins.NewService().GetManager(), clients.NewService(), MEDIAS_CONFIG_PATH, MEDIAS_CONFIG_FILENAME)

	freePort := m.Config.NextFreePortNumber //8100

	newPort := m.GetPortNumberForPlugin() //8100
	if newPort != freePort {
		t.Error("GetPortNumberForPlugin should first the NextFreePortNumber")
	}

	if len(m.Config.PortsPool) != 0 {
		t.Error("GetPortNumberForPlugin should not append the port into the pool")
	}

	newPort2 := m.GetPortNumberForPlugin() //8101
	if newPort2 == newPort {
		t.Error("GetPortNumberForPlugin should generate a new port number")
	}

	m.Config.PortsPool = append(m.Config.PortsPool, 8100)
	newPort = m.GetPortNumberForPlugin() //8100
	if newPort != 8100 {
		t.Error("GetPortNumberForPlugin should use number from the pool first")
	}
	if len(m.Config.PortsPool) != 0 {
		t.Error("GetPortNumberForPlugin should pop the port into the pool")
	}
}
