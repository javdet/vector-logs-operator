# Vector logging operator
The Logging operator automates the deployment and configuration of a Kubernetes logging. Operator uses vector instances for collecting, aggregating and storing logs into different systems.

## Work scheme
![](docs/img/vector-operator.png)

![](docs/img/vector-operator2.png)

Ресурсы
* VectorAgent - описание инсталяции daemonset агентов с базовой конфигурацией
* VectorAggregator - описание statefulset агрегаторов с базовой конфигурацией
* VectorAgentPipeline - описание sinks и transforms

## How does it work?
* VectorAgent - defines parameters for run daemonset vector-agent instances
* VectorAggregator - defines parameters for run statefulset vector-agregator instances
* VectorAgentPipeline - defines transforms and sinks sections for vector-agent
![](docs/img/vector-operator3.png)

Labels
Для сбора логов только с определенных неймспейсов не обходимо задать label `vlo.io/logs: "true"`. В таком случае будет
создан transform с названием namespaces, который может дальше использоваться, как input.

## Requirements
* Kubernetes >= 1.18
* cert-manager >= 1.5
* helm >= 3.0.0

