package v1beta

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var log = logf.Log.WithName("vectoragent-resource")

func (r *VectorAgent) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-logging-vlo-io-v1-vectoragent,mutating=true,failurePolicy=fail,sideEffects=None,groups=logging.vlo.io,resources=vectoragents,verbs=create;update,versions={v1,v1beta},name=mvectoragent.kb.io,admissionReviewVersions={v1,v1beta}

var _ webhook.Defaulter = &VectorAgent{}

func (r *VectorAgent) Default() {
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
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-logging-vlo-io-v1-vectoragent,mutating=false,failurePolicy=fail,sideEffects=None,groups=logging.vlo.io,resources=vectoragents,verbs=create;update,versions={v1,v1beta},name=vvectoragents.kb.io,admissionReviewVersions={v1,v1beta}

var _ webhook.Validator = &VectorAgent{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *VectorAgent) ValidateCreate() error {
	log.Info("validate create", "name", r.Name)
	return r.validateVectorAgent()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *VectorAgent) ValidateUpdate(old runtime.Object) error {
	log.Info("validate update", "name", r.Name)

	return r.validateVectorAgent()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *VectorAgent) ValidateDelete() error {
	log.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (r *VectorAgent) validateVectorAgent() error {
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

func (r *VectorAgent) validateVectorAgentName() *field.Error {
	if len(r.ObjectMeta.Name) > validation.DNS1123LabelMaxLength {
		return field.Invalid(field.NewPath("metadata").Child("name"), r.Name, "must be no more than 63 characters")
	}
	return nil
}
