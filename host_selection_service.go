package audiusclient

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

type HostSelectionService struct {

	// A mutex to use while performing updates on the service.
	mu sync.Mutex

	// The name of the application issuing requests to audius.
	appName string

	// The configuration for the service.
	config hostSelectionServiceConfig

	// The fetcher to use when updating the host list.
	fetcher hostFetcher

	// The fetcher to use when healtch checking a host.
	healthCheckFetcher hostHealthCheckFetcher

	// The list of possible hosts to choose from.
	// An update to the host list is only considered successful if one or more hosts was found.
	hostList []string

	// The time when the host list was last updated successfully (if at all).
	hostListUpdatedAt *time.Time

	// A mapping of unhealthy hosts (a subset of the host list) to the time they were marked unhealthy.
	unhealthyHostMap map[string]time.Time

	// The most recently selected host (among the host list).
	// If the host is an empty string, it is unset.
	selectedHost string

	// The time when the selected host was last updated successfully (if at all).
	selectedHostUpdatedAt *time.Time
}

func NewHostSelectionService(
	appName string,
) *HostSelectionService {
	hostSelectionServiceConfig := newHostSelectionServiceConfig()
	hostFetcher := newDiscoveryHostFetcher(appName)
	hostHealthCheckFetcher := newDiscoveryHostHealthCheckFetcher(appName)
	return &HostSelectionService{
		appName:            appName,
		config:             hostSelectionServiceConfig,
		fetcher:            hostFetcher,
		healthCheckFetcher: hostHealthCheckFetcher,
	}
}

func (s *HostSelectionService) getHostList() ([]string, error) {

	// Check if the host list has been fetched recently enough - if so just short circuit and return it.
	if len(s.hostList) != 0 && s.hostListUpdatedAt != nil && time.Since(*s.hostListUpdatedAt) < s.config.HostListTTL {
		return s.hostList, nil
	}

	// We need to re-fetch the host list.
	hosts, err := s.fetcher.FetchHosts()
	if err != nil {
		return nil, err
	}
	if len(hosts) == 0 {
		return nil, errors.New("fetched hosts were empty")
	}

	// Save the new host list:
	s.hostList = hosts
	t := time.Now()
	s.hostListUpdatedAt = &t
	s.unhealthyHostMap = map[string]time.Time{}
	s.selectedHost = ""
	s.selectedHostUpdatedAt = nil

	return hosts, nil
}

type hostHealthCheckResult struct {
	Duration time.Duration
	Err      error
}

func (s *HostSelectionService) GetSelectedHost() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the selected host has been fetched recently enough - if so just short circuit and return it.
	if s.selectedHost != "" && s.selectedHostUpdatedAt != nil && time.Since(*s.selectedHostUpdatedAt) < s.config.SelectedHostTTL {
		return s.selectedHost, nil
	}

	// We need to re-evaluate the hosts to select a new host.
	hosts, err := s.getHostList()
	if err != nil {
		return "", err
	}

	// Filter out the unhealthy hosts:
	filteredHosts := []string{}
	for _, host := range hosts {
		unhealthyAt, unhealthy := s.unhealthyHostMap[host]
		if unhealthy && time.Since(unhealthyAt) < s.config.UnhealthyHostTTL {
			// This host is still considered unhealthy - filter it out.
		} else {
			filteredHosts = append(filteredHosts, host)
		}
	}

	// If the filtered hosts are empty, short circuit and return.
	if len(filteredHosts) == 0 {
		return "", errors.New("all hosts are currently unhealthy")
	}

	// Shuffle the filtered hosts:
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(filteredHosts), func(i, j int) { filteredHosts[i], filteredHosts[j] = filteredHosts[j], filteredHosts[i] })

	// Pick a random subset of the hosts to test:
	var hostsToTest []string
	if len(filteredHosts) < s.config.MaximumConcurrentRequests {
		hostsToTest = filteredHosts
	} else {
		hostsToTest = filteredHosts[:s.config.MaximumConcurrentRequests]
	}

	// Health check the hosts:
	results := make([]hostHealthCheckResult, len(hostsToTest))
	sem := make(chan hostHealthCheckResult, len(hostsToTest))
	for index, host := range hostsToTest {
		go func(index int, host string) {
			duration, err := s.healthCheckFetcher.FetchHostHealthCheck(host)
			res := hostHealthCheckResult{
				Duration: duration,
				Err:      err,
			}
			results[index] = res
			sem <- res
		}(index, host)
	}
	for i := 0; i < len(hostsToTest); i++ {
		<-sem
	}

	// Select the best host, and mark the hosts that failed their health checks as unhealthy:
	var minimumDuration time.Duration
	var selectedHost string
	for index, healthCheckResult := range results {
		host := hostsToTest[index]
		if healthCheckResult.Err != nil {
			s.unhealthyHostMap[host] = time.Now()
		} else if selectedHost == "" || healthCheckResult.Duration < minimumDuration {
			minimumDuration = healthCheckResult.Duration
			selectedHost = host
		}
	}
	if selectedHost == "" {
		return "", errors.New("all tested hosts failed their health checks")
	}
	s.selectedHost = selectedHost

	return selectedHost, nil
}
