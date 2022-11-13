# Vector logging operator
The Logging operator automates the deployment and configuration of a Kubernetes logging. Operator uses vector instances for collecting, aggregating and storing logs into different systems.

## Work scheme
![](docs/img/vector-operator.png)

![](docs/img/vector-operator2.png)

Ресурсы
* VectorAgent - описание инсталяции daemonset агентов с базовой конфигурацией
* VectorAgregator - описание statefulset агрегаторов с базовой конфигурацией
* VectorAgentPipeline - описание sinks
* VectorAgregatorPipeline - описание пайплайна для агрегатора

## How does it work?
* VectorAgent - defines parameters for run daemonset vector-agent instances
* VectorAgregator - defines parameters for run statefulset vector-agregator instances
* VectorAgentPipeline - defines transforms and sinks sections for vector-agent
* VectorAgregatorPipeline - define sources, transforms and sinks sections for vector-agregator
![](docs/img/vector-operator3.png)


## Requirements
* Kubernetes >= 1.18
* cert-manager >= 1.5
* helm >= 3.0.0

