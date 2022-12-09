# VectorAgentPipeline
CRD defines custom transform and sink sections for vector agent

## VectorAgentPipelineSpec
`transforms (string, optional)`
`sinks ([]VectorPipelineSinks)`
`selector (map[string]string)` VectorAgent selector

##VectorPipelineSinks
`s3 (VectorPipelineSinksS3, optional)"`
`console (VectorPipelineSinksConsole, optional)`
`file (VectorPipelineSinksFile, optional)`
`elasticsearch (VectorPipelineSinksElasticsearch, optional)"`
`http (VectorPipelineSinksHTTP, optional)`
`kafka (VectorPipelineSinksKafka, optional)`
`loki (VectorPipelineSinksLoki, optional)`
`vector (VectorPipelineSinksVector, optional)`

