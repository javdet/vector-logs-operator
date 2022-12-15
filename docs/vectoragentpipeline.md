# VectorAgentPipeline
CRD defines custom transform and sink sections for vector agent

## VectorAgentPipelineSpec
`transforms (string, optional)` [Transform rules](https://vector.dev/docs/reference/configuration/transforms/)
`sinks ([]VectorPipelineSinks)`
`selector (map[string]string)` VectorAgent selector

## VectorPipelineSinks
`s3 (VectorPipelineSinksS3, optional)"` [s3](https://vector.dev/docs/reference/configuration/sinks/aws_s3/) sink configuration
`console (VectorPipelineSinksConsole, optional)` [console](https://vector.dev/docs/reference/configuration/sinks/console/) sink configuration
`file (VectorPipelineSinksFile, optional)` [file](https://vector.dev/docs/reference/configuration/sinks/file/) sink configuration
`elasticsearch (VectorPipelineSinksElasticsearch, optional)"` [elasticsearch](https://vector.dev/docs/reference/configuration/sinks/elasticsearch/) sink configuration
`http (VectorPipelineSinksHTTP, optional)` [http](https://vector.dev/docs/reference/configuration/sinks/http/) sink configuration
`kafka (VectorPipelineSinksKafka, optional)` [kafka](https://vector.dev/docs/reference/configuration/sinks/kafka/) sink configuration
`loki (VectorPipelineSinksLoki, optional)` [loki](https://vector.dev/docs/reference/configuration/sinks/loki/) sink configuration
`vector (VectorPipelineSinksVector, optional)` [vector](https://vector.dev/docs/reference/configuration/sinks/vector/) sink configuration

## VectorPipelineSinksS3
`name (string)` name of sink
`inputs ([]string)` list of inputs. Need to match to any sources or transforms
`bucket (string, optional)` bucket name
`region (string)` region name
`acl (string)` bucket ACL
`compression (string)` compression type
`contentType (string, optional)` Content-type
`encoding (string, optional)` encoding format
`endpoint (string, optional)` custom s3 endpoint
`keyPrefix (string, optional)` a prefix to apply to all object key names.
`serverSideEncryption (string, optional)` Server-side Encryption algorithm used when storing these objects.
`secret (VectorPipelineSinksS3Secret)` secret name with credentials

## VectorPipelineSinksS3Secret
`name (string)` secret name

## VectorPipelineSinksConsole
`name (string)` name of sink
`inputs ([]string)` list of inputs. Need to match to any sources or transforms
`target (string, optional)` standard stream to write to
`encoding (string, optional)` encoding format

## VectorPipelineSinksFile
`name (string)` name of sink
`inputs ([]string)` list of inputs. Need to match to any sources or transforms
`compression (string, optional)` compression type
`path (string)` file path
`encoding (string, optional)` encoding format

## VectorPipelineSinksElasticsearch
`name (string)` name of sink
`inputs ([]string)` list of inputs. Need to match to any sources or transforms
`compression (string, optional)` compression type
`endpoints ([]string)` list of elasticsearch endpoints
`pipeline (string, optional)`
`index (string, optional)` index name
`mode (string, optional)` write mode
`secret (VectorPipelineSinksElasticsearchSecret, optional)` secret name with credentials
`idKey (string, optional)`
`tlsCA (string, optional)`

## VectorPipelineSinksHTTP
`name (string)` name of sink
`inputs ([]string)` list of inputs. Need to match to any sources or transforms
`compression (string, optional)` compression type
`uri (string)` HTTP URI
`encoding (string, optional)` encoding format
`secret (VectorPipelineSinksHTTPSecret, optional)` secret name with credentials
`method (string, optional)` HTTP method
`tlsCA (string, optional)`

## VectorPipelineSinksKafka
`name (string)` name of sink
`inputs ([]string)` list of inputs. Need to match to any sources or transforms
`bootstrapServers (string)` list kafka servers comma separated
`keyField (string, optional)` partition key
`topic (string)` topic name
`compression (string, optional)` compression type
`encoding (string, optional)` encoding format
`sasl (VectorPipelineSinksKafkaSasl, optional)` auth configuration

## VectorPipelineSinksKafkaSasl
`mechanism (string)` auth mechanism
`secret (VectorPipelineSinksKafkaSaslSecret)` secret name with credentials

## VectorPipelineSinksKafkaSaslSecret
`name (string)` name of sink
`usernameKey (string)` username
`passwordKey (string)` password

## VectorPipelineSinksLoki
`name (string)` name of sink
`inputs ([]string)` list of inputs. Need to match to any sources or transforms
`endpoint (string)` Loki endpoint
`labels (map[string]string, optional)` array of labels
`compression (string, optional)` compression type
`encoding (string, optional)` encoding format
`exceptFields (string, optional)`
`tenantId (string, optional)` tenant id

## VectorPipelineSinksVector
`name (string)` name of sink
`inputs ([]string)` list of inputs. Need to match to any sources or transforms
`compression (string, optional)` compression type
`address (string)` vector address and port
