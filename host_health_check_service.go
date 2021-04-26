package audiusclient

import (
	"sync"
	"time"
)

type hostHealthCheckResult struct {
	Host        string
	CompletedAt time.Time
	Duration    time.Duration
	Err         error
}

type HostHealthCheckService struct {

	// A mutex to use while performing updates on the service.
	mu sync.Mutex

	// The configuration for the service.
	config hostHealthCheckServiceConfig

	// The fetcher to use when healtch checking a host.
	healthCheckFetcher HostHealthCheckFetcher

	// A mapping of hosts to the result the last time they were health checked.
	hostHealthCheckResultMap map[string]hostHealthCheckResult
}

func NewHostHealthCheckService(healthCheckFetcher HostHealthCheckFetcher) *HostHealthCheckService {
	hostHealthCheckServiceConfig := newHostHealthCheckServiceConfig()
	return &HostHealthCheckService{
		config:             hostHealthCheckServiceConfig,
		healthCheckFetcher: healthCheckFetcher,
	}
}

func (s *HostHealthCheckService) HealthCheckHosts(hosts []string) map[string]hostHealthCheckResult {
	s.mu.Lock()
	defer s.mu.Unlock()

	// First check to see which hosts need to be health checked again:
	resultMap := map[string]hostHealthCheckResult{}
	var hostsToTest []string
	for _, host := range hosts {
		result, ok := s.hostHealthCheckResultMap[host]
		if ok && time.Since(result.CompletedAt) < s.config.HostHealthCheckResultTTL {
			// If the host was healthchecked recently enough, just use the last result.
			resultMap[host] = result
		} else {
			// Otherwise, we need to test this host.
			hostsToTest = append(hostsToTest, host)
		}
	}

	// If all of the hosts were health checked recently enough, just return the resultMap.
	if len(resultMap) == len(hosts) {
		return resultMap
	}

	// Health check the hosts:
	results := make([]hostHealthCheckResult, len(hostsToTest))
	sem := make(chan hostHealthCheckResult, len(hostsToTest))
	for index, host := range hostsToTest {
		go func(index int, host string) {
			duration, err := s.healthCheckFetcher.FetchHostHealthCheck(host)
			res := hostHealthCheckResult{
				Host:        host,
				CompletedAt: time.Now(),
				Duration:    duration,
				Err:         err,
			}
			results[index] = res
			sem <- res
		}(index, host)
	}
	for i := 0; i < len(hostsToTest); i++ {
		<-sem
	}

	// Combine the results into the result map:
	for _, result := range results {
		resultMap[result.Host] = result
	}

	return resultMap
}
