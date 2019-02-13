package mixer

import (
	"fmt"

	istiov1beta1 "github.com/banzaicloud/istio-operator/pkg/apis/operator/v1beta1"
	"github.com/banzaicloud/istio-operator/pkg/k8sutil"
	"github.com/banzaicloud/istio-operator/pkg/resources"
	"github.com/go-logr/logr"
	"github.com/goph/emperror"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	serviceAccountName     = "istio-mixer-service-account"
	clusterRoleName        = "istio-mixer-cluster-role"
	clusterRoleBindingName = "istio-mixer-cluster-role-binding"
	configMapName          = "istio-statsd-prom-bridge"
)

var mixerLabels = map[string]string{
	"app": "mixer",
}

var labelSelector = map[string]string{
	"istio": "mixer",
}

type Reconciler struct {
	resources.Reconciler
	dynamic dynamic.Interface
}

type MixerResource func(t string, owner *istiov1beta1.Config) runtime.Object

func rr(t string, tResources []MixerResource) []resources.Resource {
	resources := make([]resources.Resource, 0)
	for _, r := range tResources {
		resources = append(resources, func(owner *istiov1beta1.Config) runtime.Object {
			return r(t, owner)
		})
	}
	return resources
}

func New(client client.Client, dc dynamic.Interface, istio *istiov1beta1.Config) *Reconciler {
	return &Reconciler{
		Reconciler: resources.Reconciler{
			Client: client,
			Owner:  istio,
		},
		dynamic: dc,
	}
}

func (r *Reconciler) Reconcile(log logr.Logger) error {
	res := []resources.Resource{
		r.serviceAccount,
		r.clusterRole,
		r.clusterRoleBinding,
		r.configMap,
	}
	mResources := []MixerResource{
		r.deployment,
		r.service,
		r.horizontalPodAutoscaler,
		r.destinationRule,
	}
	res = append(res, rr("policy", mResources)...)
	res = append(res, rr("telemetry", mResources)...)
	for _, res := range res {
		o := res(r.Owner)
		err := k8sutil.Reconcile(log, r.Client, o)
		if err != nil {
			return emperror.WrapWith(err, "failed to reconcile resource", "resource", o.GetObjectKind().GroupVersionKind())
		}
	}
	dcrs := r.dynamicCustomResources(r.Owner)
	for _, res := range dcrs {
		err := res.Reconcile(log, r.dynamic)
		if err != nil {
			return emperror.WrapWith(err, "failed to reconcile dynamic resource", "resource", res.Gvr.Resource, "name", res.Name)
		}
	}
	return nil
}

func deploymentName(t string) string {
	return fmt.Sprintf("istio-%s", t)
}

func serviceName(t string) string {
	return fmt.Sprintf("istio-%s", t)
}

func hpaName(t string) string {
	return fmt.Sprintf("istio-%s-autoscaler", t)
}

func destinationRuleName(t string) string {
	return fmt.Sprintf("istio-%s", t)
}

func appLabel(t string) map[string]string {
	return map[string]string{
		"app": t,
	}
}

func mixerTypeLabel(t string) map[string]string {
	return map[string]string{
		"istio-mixer-type": t,
	}
}
