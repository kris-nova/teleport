package services

import (
	"time"

	"github.com/gravitational/log"
	"github.com/gravitational/teleport/backend"
	"github.com/gravitational/trace"
)

type ProvisioningService struct {
	backend backend.Backend
}

func NewProvisioningService(backend backend.Backend) *ProvisioningService {
	return &ProvisioningService{backend}
}

// Tokens are provisioning tokens for the auth server
func (s *ProvisioningService) UpsertToken(token, fqdn string, ttl time.Duration) error {
	err := s.backend.UpsertVal([]string{"tokens"}, token, []byte(fqdn), ttl)
	if err != nil {
		log.Errorf(err.Error())
		return trace.Wrap(err)
	}
	return nil

}
func (s *ProvisioningService) GetToken(token string) (string, error) {
	fqdn, err := s.backend.GetVal([]string{"tokens"}, token)
	if err != nil {
		log.Errorf(err.Error())
		return "", convertErr(err)
	}
	return string(fqdn), nil
}
func (s *ProvisioningService) DeleteToken(token string) error {
	err := s.backend.DeleteKey([]string{"tokens"}, token)
	if err != nil {
		log.Errorf(err.Error())
	}
	return convertErr(err)
}