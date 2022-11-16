# VectorAgent
CRD defines base configuration for vector agents DaemonSet.

## VectorAgentSpec
`image (string, optional)` Default: 'timberio/vector'
`tag (string, optional)` Default: '0.23.0-distroless-libc'
`metricsScrapeInterval (int, optional)` Default: 15
`internalLogs (bool)`
`logLevel (string, optional)` Default: info
`resources (VectorAgentSpecResources, optional)`
`podAnnotations (map[string]string, optional)`

## VectorAgentSpecResources
`limits (VectorAgentSpecResourcesLimits, optional)`
`requests (VectorAgentSpecResourcesRequests, optional)`

## M2LogstashSpecRequestsLimits
`cpu (string, optional)` Default: '100m'
`memory (string, optional)` Default '256Mi'

## VectorAgentSpecResourcesRequests
`cpu (string, optional)` Default: '1000m'
`memory (string, optional)` Default '1Gi'
