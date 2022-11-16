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

func (r *VectorAgentPipeline) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-logging-vlo-io-v1beta-vectoragentpipeline,mutating=true,failurePolicy=fail,sideEffects=None,groups=logging.vlo.io,resources=vectoragentpipelines,verbs=create;update,versions={v1,v1beta},name=mvectoragentpipeline.kb.io,admissionReviewVersions={v1,v1beta}

var _ webhook.Defaulter = &VectorAgentPipeline{}

func (r *VectorAgentPipeline) Default() {
	log.Info("default", "name", r.Name)
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
		if sinks.Vector.Name == "" {
			sinks.Vector.Name = fmt.Sprint("vector_", id)
		}
		if sinks.Vector.Compression == "" {
			sinks.Vector.Compression = "false"
		}
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-logging-vlo-io-v1beta-vectoragentpipeline,mutating=false,failurePolicy=fail,sideEffects=None,groups=logging.vlo.io,resources=vectoragentpipelines,verbs=create;update,versions={v1,v1beta},name=vvectoragentpipeline.kb.io,admissionReviewVersions={v1,v1beta}

var _ webhook.Validator = &VectorAgentPipeline{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *VectorAgentPipeline) ValidateCreate() error {
	log.Info("validate create", "name", r.Name)
	return r.validateVectorAgent()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *VectorAgentPipeline) ValidateUpdate(old runtime.Object) error {
	log.Info("validate update", "name", r.Name)

	return r.validateVectorAgent()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *VectorAgentPipeline) ValidateDelete() error {
	log.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (r *VectorAgentPipeline) validateVectorAgent() error {
	var allErrs field.ErrorList
	if err := r.validateVectorAgentName(); err != nil {
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

func (r *VectorAgentPipeline) validateVectorAgentName() *field.Error {
	if len(r.ObjectMeta.Name) > validation.DNS1123LabelMaxLength {
		return field.Invalid(field.NewPath("metadata").Child("name"), r.Name, "must be no more than 63 characters")
	}
	return nil
}
