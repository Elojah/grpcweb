package grpcweb

import "github.com/elojah/game_02/pkg/errors"

// Config is redis structure config.
type Config struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`

	ConnectionTimeout uint `json:"connection_timeout"`
	NumStreamWorkers  uint `json:"num_stream_workers"`
	MaxMsgSize        uint `json:"max_msg_size"`

	Origin map[string]struct{} `json:"origin"`
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.ErrEmptyNamespace{}
	}

	cCert, ok := fconf["cert"]
	if !ok {
		return errors.ErrMissingKey{Key: "cert"}
	}

	if c.Cert, ok = cCert.(string); !ok {
		return errors.ErrInvalidType{
			Key:    "cert",
			Expect: "string",
			Value:  cCert,
		}
	}

	cKey, ok := fconf["key"]
	if !ok {
		return errors.ErrMissingKey{Key: "key"}
	}

	if c.Key, ok = cKey.(string); !ok {
		return errors.ErrInvalidType{
			Key:    "key",
			Expect: "string",
			Value:  cKey,
		}
	}

	cConnectionTimeout, ok := fconf["connection_timeout"]
	if !ok {
		return errors.ErrMissingKey{Key: "connection_timeout"}
	}

	f, ok := cConnectionTimeout.(float64)
	if !ok {
		return errors.ErrInvalidType{
			Key:    "connection_timeout",
			Expect: "number",
			Value:  cConnectionTimeout,
		}
	}
	c.ConnectionTimeout = uint(f)

	cNumStreamWorkers, ok := fconf["num_stream_workers"]
	if !ok {
		return errors.ErrMissingKey{Key: "num_stream_workers"}
	}

	f, ok = cNumStreamWorkers.(float64)
	if !ok {
		return errors.ErrInvalidType{
			Key:    "num_stream_workers",
			Expect: "number",
			Value:  cNumStreamWorkers,
		}
	}
	c.NumStreamWorkers = uint(f)

	cMaxMsgSize, ok := fconf["max_msg_size"]
	if !ok {
		return errors.ErrMissingKey{Key: "max_msg_size"}
	}

	f, ok = cMaxMsgSize.(float64)
	if !ok {
		return errors.ErrInvalidType{
			Key:    "max_msg_size",
			Expect: "number",
			Value:  cMaxMsgSize,
		}
	}
	c.MaxMsgSize = uint(f)

	cOrigin, ok := fconf["max_msg_size"]
	if !ok {
		return errors.ErrMissingKey{Key: "max_msg_size"}
	}

	m, ok := cOrigin.(map[string]struct{})
	if !ok {
		return errors.ErrInvalidType{
			Key:    "max_msg_size",
			Expect: "map[string]struct{}",
			Value:  cOrigin,
		}
	}
	c.Origin = m

	return nil
}
