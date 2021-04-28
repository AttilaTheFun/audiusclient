package audiusclient

import (
	"errors"
	"log"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"
)

type HostSelectionService struct {

	// A mutex to use while performing updates on the service.
	mu sync.Mutex

	// The name of the application issuing requests to Audius.
	appName string

	// The configuration for the service.
	config hostSelectionServiceConfig

	// The fetcher to use when updating the host list.
	fetcher hostFetcher

	// The service to use for health checking hosts.
	healthCheckService *HostHealthCheckService

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
	selectionServiceConfig := newHostSelectionServiceConfig()
	fetcher := NewDiscoveryHostFetcher(appName)
	healthCheckFetcher := NewDiscoveryHostHealthCheckFetcher(appName)
	healthCheckService := NewHostHealthCheckService(healthCheckFetcher)
	return &HostSelectionService{
		appName:            appName,
		config:             selectionServiceConfig,
		fetcher:            fetcher,
		healthCheckService: healthCheckService,
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

func (s *HostSelectionService) GetSelectedHost() (string, error) {
	s.mu.Lock()

	startTime := time.Now()
	log.Printf("Started host selection at: %v", startTime)
	log.Println()
	defer func() {
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Completed host selection with host: %v at: %v, duration: %v", s.selectedHost, endTime, duration)
		log.Println()

		s.mu.Unlock()
	}()

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
	resultMap := s.healthCheckService.HealthCheckHosts(hostsToTest)

	// Mark the unhealthy hosts and put the healthy hosts in an array for sorting:
	var healthyHosts []hostHealthCheckResult
	for host, healthCheckResult := range resultMap {
		if healthCheckResult.Err != nil {
			s.unhealthyHostMap[host] = time.Now()
		} else {
			healthyHosts = append(healthyHosts, healthCheckResult)
		}
	}

	// If the healthy hosts are empty, short circuit and return.
	if len(healthyHosts) == 0 {
		return "", errors.New("all hosts are currently unhealthy")
	}

	// Sort the healthy hosts in order of preference.
	sort.Slice(healthyHosts, func(i, j int) bool {
		firstHostResult := healthyHosts[i]
		secondHostResult := healthyHosts[j]

		// We prefer Audius hosts above all others:
		isFirstHostAudius := strings.HasSuffix(firstHostResult.Host, "audius.co")
		isSecondHostAudius := strings.HasSuffix(secondHostResult.Host, "audius.co")
		if isFirstHostAudius && !isSecondHostAudius {
			return true
		}
		if isSecondHostAudius && !isFirstHostAudius {
			return false
		}

		// We prefer staked hosts over others:
		isFirstHostStaked := strings.HasSuffix(firstHostResult.Host, "staked.cloud")
		isSecondHostStaked := strings.HasSuffix(secondHostResult.Host, "staked.cloud")
		if isFirstHostStaked && !isSecondHostStaked {
			return true
		}
		if isSecondHostStaked && !isFirstHostStaked {
			return false
		}

		// We prefer faster hosts last:
		return firstHostResult.Duration < secondHostResult.Duration
	})

	// Select the first host in order of preference:
	selectedHost := healthyHosts[0].Host
	s.selectedHost = selectedHost
	t := time.Now()
	s.selectedHostUpdatedAt = &t

	return selectedHost, nil
}
