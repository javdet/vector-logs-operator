# VectorAggregator
CRD defines base configuration for vector aggregator StatefulSet and pipeline.

## VectorAggregatorSpec
`image (string, optional)` Default: 'timberio/vector'
`tag (string, optional)` Default: '0.23.0-distroless-libc'
`metricsScrapeInterval (int, optional)` Default: 15
`internalLogs (bool)`
`logLevel (string, optional)` Default: info
`resources (VectorAgentSpecResources, optional)`
`podAnnotations (map[string]string, optional)`
`sources ([]VectorAggregatorSources, optional)` sources config
`transforms (string, optional)`
`sinks ([]VectorAggregatorSinks, optional)`

## VectorAggregatorSpecResources
`limits (VectorAgentSpecResourcesLimits, optional)`
`requests (VectorAgentSpecResourcesRequests, optional)`

## VectorAggregatorSpecRequestsLimits
`cpu (string, optional)` Default: '100m'
`memory (string, optional)` Default '256Mi'

## VectorAggregatorSpecResourcesRequests
`cpu (string, optional)` Default: '1000m'
`memory (string, optional)` Default '1Gi'

## VectorAggregatorSources
`kafka (VectorPipelineSourcesKafka, optional)` [kafka](https://vector.dev/docs/reference/configuration/sources/kafka/) source configuration
`vector (VectorPipelineSourcesVector, optional)` [vector](https://vector.dev/docs/reference/configuration/sources/vector/) source configuration
`amqp (VectorPipelineSourcesAMQP, optional)` [amqp](https://vector.dev/docs/reference/configuration/sources/amqp/) source configuration
`s3 (VectorPipelineSourcesS3, optional)` [s3](https://vector.dev/docs/reference/configuration/sources/aws_s3/) source configuration

## VectorAggregatorSourcesKafka
`name (string)` name of source
`bootstrapServers (string)` list of kafka hosts
`keyField (string, optional)` partition key
`topics ([]string, optional)` list of topics
`decoding (string, optional)`
`groupID (string)` consumer group
`sasl (VectorPipelineSinksKafkaSasl)`
`autoOffsetReset (string)`

## VectorPipelineSourcesVector
`name (string)` name of source
`host (string)` listen address
`port (int)` listen port
`version (string, optional)`

## VectorPipelineSourcesAMQP
`name (string)` name of source
`connection (string)` connection URI
`groupID (string)`
`offsetKey (string, optional)`

## VectorPipelineSourcesS3
`name (string)` name of source
`region (string)`
`sqs (string)`
`compression (string, optional)`
`endpoint (string, optional)`

## VectorAggregatorSinks
`s3 (VectorPipelineSinksS3, optional)"` s3 sink configuration
`console (VectorPipelineSinksConsole, optional)` console sink configuration
`file (VectorPipelineSinksFile, optional)` file sink configuration
`elasticsearch (VectorPipelineSinksElasticsearch, optional)"` elasticsearch sink configuration
`http (VectorPipelineSinksHTTP, optional)` http sink configuration
`kafka (VectorPipelineSinksKafka, optional)` kafka sink configuration
`loki (VectorPipelineSinksLoki, optional)` loki sink configuration
`vector (VectorPipelineSinksVector, optional)` vector sink configuration