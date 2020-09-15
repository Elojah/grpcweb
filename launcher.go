package grpcweb

import (
	"sync"

	"github.com/elojah/services"
)

// Namespaces maps configs used for grpcweb service with config file namespaces.
type Namespaces struct {
	GRPCWeb services.Namespace
}

// Launcher represents a grpcweb launcher.
type Launcher struct {
	*services.Configs
	ns Namespaces

	s *Service
	m sync.Mutex
}

// NewLauncher returns a new grpcweb Launcher.
func (s *Service) NewLauncher(ns Namespaces, nsRead ...services.Namespace) *Launcher {
	return &Launcher{
		Configs: services.NewConfigs(nsRead...),
		s:       s,
		ns:      ns,
	}
}

// Up starts the grpcweb service with new configs.
func (l *Launcher) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	sconfig := Config{}
	if err := sconfig.Dial(configs[l.ns.GRPCWeb]); err != nil {
		return err
	}

	return l.s.Dial(sconfig)
}

// Down stops the grpcweb service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return l.s.Close()
}
