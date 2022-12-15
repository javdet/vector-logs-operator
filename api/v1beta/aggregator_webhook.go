package v1beta

import (
	"fmt"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *VectorAggregator) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-logging-vlo-io-v1beta-vectoraggregator,mutating=true,failurePolicy=fail,sideEffects=None,groups=logging.vlo.io,resources=vectoraggregators,verbs=create;update,versions={v1,v1beta},name=mvectoraggregator.kb.io,admissionReviewVersions={v1,v1beta}

var _ webhook.Defaulter = &VectorAggregator{}

func (r *VectorAggregator) Default() {
	log.Info("default", "name", r.Name)
	if r.Spec.Image == "" {
		r.Spec.Image = "timberio/vector"
	}
	if r.Spec.Tag == "" {
		r.Spec.Tag = "0.23.0-distroless-libc"
	}
	if r.Spec.Resources.Limit.Cpu == "" {
		r.Spec.Resources.Limit.Cpu = "1000m"
	}
	if r.Spec.Resources.Limit.Memory == "" {
		r.Spec.Resources.Limit.Memory = "1Gi"
	}
	if r.Spec.Resources.Requests.Cpu == "" {
		r.Spec.Resources.Requests.Cpu = "100m"
	}
	if r.Spec.Resources.Requests.Memory == "" {
		r.Spec.Resources.Requests.Memory = "256Mi"
	}
	if r.Spec.LogLevel == "" {
		r.Spec.LogLevel = "info"
	}
	if r.Spec.MetricsScrapeInterval == 0 {
		r.Spec.MetricsScrapeInterval = 15
	}

	for id, source := range r.Spec.Sources {
		if source.S3.Name == "" {
			source.S3.Name = fmt.Sprint("s3_", id)
		}
		if source.S3.Compression == "" {
			source.S3.Compression = "none"
		}
		if source.Kafka.Name == "" {
			source.Kafka.Name = fmt.Sprint("kafka_", id)
		}
		if source.Vector.Name == "" {
			source.Vector.Name = fmt.Sprint("vector_", id)
		}
		if source.Vector.Port == 0 {
			source.Vector.Port = 9000
		}
		if source.Vector.Version == "" {
			source.Vector.Version = "2"
		}
		if source.AMQP.Name == "" {
			source.AMQP.Name = fmt.Sprint("amqp_", id)
		}
	}

	for id, sinks := range r.Spec.Sinks {
		if sinks.S3.Name == "" {
			sinks.S3.Name = fmt.Sprint("s3_", id)
		}
		if sinks.S3.Compression == "" {
			sinks.S3.Compression = "gzip"
		}
		if sinks.S3.ContentType == "" {
			sinks.S3.ContentType = "text/x-log"
		}
		if sinks.S3.KeyPrefix == "" {
			sinks.S3.KeyPrefix = "date=%F/"
		}
		if sinks.S3.Encoding == "" {
			sinks.S3.Encoding = "text"
		}
		if sinks.Console.Name == "" {
			sinks.Console.Name = fmt.Sprint("console_", id)
		}
		if sinks.Console.Encoding == "" {
			sinks.Console.Encoding = "text"
		}
		if sinks.Console.Target == "" {
			sinks.Console.Target = "stdout"
		}
		if sinks.File.Name == "" {
			sinks.File.Name = fmt.Sprint("file_", id)
		}
		if sinks.File.Compression == "" {
			sinks.File.Compression = "none"
		}
		if sinks.File.Encoding == "" {
			sinks.File.Encoding = "text"
		}
		if sinks.Elasticsearch.Name == "" {
			sinks.Elasticsearch.Name = fmt.Sprint("elasticsearch_", id)
		}
		if sinks.Elasticsearch.Compression == "" {
			sinks.Elasticsearch.Compression = "none"
		}
		if sinks.Elasticsearch.Mode == "" {
			sinks.Elasticsearch.Mode = "bulk"
		}
		if sinks.HTTP.Encoding == "" {
			sinks.HTTP.Encoding = "text"
		}
		if sinks.HTTP.Name == "" {
			sinks.HTTP.Name = fmt.Sprint("http_", id)
		}
		if sinks.HTTP.Compression == "" {
			sinks.HTTP.Compression = "none"
		}
		if sinks.HTTP.Method == "" {
			sinks.HTTP.Method = "POST"
		}
		if sinks.Kafka.Name == "" {
			sinks.Kafka.Name = fmt.Sprint("kafka_", id)
		}
		if sinks.Kafka.Compression == "" {
			sinks.Kafka.Compression = "none"
		}
		if sinks.Loki.Name == "" {
			sinks.Loki.Name = fmt.Sprint("loki_", id)
		}
		if sinks.Loki.Compression == "" {
			sinks.Loki.Compression = "snappy"
		}
		if sinks.Loki.Encoding == "" {
			sinks.Loki.Encoding = "text"
		}
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-logging-vlo-io-v1beta-vectoraggregator,mutating=false,failurePolicy=fail,sideEffects=None,groups=logging.vlo.io,resources=vectoraggregators,verbs=create;update,versions={v1,v1beta},name=vvectoraggregator.kb.io,admissionReviewVersions={v1,v1beta}

var _ webhook.Validator = &VectorAggregator{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *VectorAggregator) ValidateCreate() error {
	log.Info("validate create", "name", r.Name)
	return r.validateVectorAggregator()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *VectorAggregator) ValidateUpdate(old runtime.Object) error {
	log.Info("validate update", "name", r.Name)

	return r.validateVectorAggregator()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *VectorAggregator) ValidateDelete() error {
	log.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (r *VectorAggregator) validateVectorAggregator() error {
	var allErrs field.ErrorList
	if err := r.validateVectorAggregatorName(); err != nil {
		allErrs = append(allErrs, err)
	}
	//todo: validation Spec
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "vlo.io", Kind: "VectorAgent"},
		r.Name, allErrs)
}

func (r *VectorAggregator) validateVectorAggregatorName() *field.Error {
	if len(r.ObjectMeta.Name) > validation.DNS1123LabelMaxLength {
		return field.Invalid(field.NewPath("metadata").Child("name"), r.Name, "must be no more than 63 characters")
	}
	return nil
}
