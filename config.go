package grpcweb

import "github.com/elojah/services"

// Config is grpcweb structure config.
type Config struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`

	ConnectionTimeout uint `json:"connection_timeout"`
	NumStreamWorkers  uint `json:"num_stream_workers"`
	MaxRecvMsgSize    uint `json:"max_recv_msg_size"`

	Origin map[string]struct{} `json:"origin"`
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return services.ErrEmptyNamespace{}
	}

	cCert, ok := fconf["cert"]
	if !ok {
		return services.ErrMissingKey{Key: "cert"}
	}

	if c.Cert, ok = cCert.(string); !ok {
		return services.ErrInvalidType{
			Key:    "cert",
			Expect: "string",
			Value:  cCert,
		}
	}

	cKey, ok := fconf["key"]
	if !ok {
		return services.ErrMissingKey{Key: "key"}
	}

	if c.Key, ok = cKey.(string); !ok {
		return services.ErrInvalidType{
			Key:    "key",
			Expect: "string",
			Value:  cKey,
		}
	}

	cConnectionTimeout, ok := fconf["connection_timeout"]
	if !ok {
		return services.ErrMissingKey{Key: "connection_timeout"}
	}

	f, ok := cConnectionTimeout.(float64)
	if !ok {
		return services.ErrInvalidType{
			Key:    "connection_timeout",
			Expect: "number",
			Value:  cConnectionTimeout,
		}
	}
	c.ConnectionTimeout = uint(f)

	cNumStreamWorkers, ok := fconf["num_stream_workers"]
	if !ok {
		return services.ErrMissingKey{Key: "num_stream_workers"}
	}

	f, ok = cNumStreamWorkers.(float64)
	if !ok {
		return services.ErrInvalidType{
			Key:    "num_stream_workers",
			Expect: "number",
			Value:  cNumStreamWorkers,
		}
	}
	c.NumStreamWorkers = uint(f)

	cMaxMsgSize, ok := fconf["max_recv_msg_size"]
	if !ok {
		return services.ErrMissingKey{Key: "max_recv_msg_size"}
	}

	f, ok = cMaxMsgSize.(float64)
	if !ok {
		return services.ErrInvalidType{
			Key:    "max_recv_msg_size",
			Expect: "number",
			Value:  cMaxMsgSize,
		}
	}
	c.MaxRecvMsgSize = uint(f)

	cOrigin, ok := fconf["origin"]
	if !ok {
		return services.ErrMissingKey{Key: "origin"}
	}

	m, ok := cOrigin.(map[string]struct{})
	if !ok {
		return services.ErrInvalidType{
			Key:    "origin",
			Expect: "map[string]struct{}",
			Value:  cOrigin,
		}
	}
	c.Origin = m

	return nil
}
