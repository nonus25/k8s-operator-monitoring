package monitoring

import (
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	ruleName                     = "monitor-operator-rules"
	alertRuleGroup               = "monitor.rules"
	deploymentSizeUndesiredAlert = "MonitorDeploymentSizeUndesired"
	operatorDownAlert            = "MonitorOperatorDown"
	operatorUpTotalRecordingRule = "monitor_operator_up_total"
)

// NewPrometheusRule creates new PrometheusRule(CR) for the operator to have alerts and recording rules
func NewPrometheusRule(namespace string) *monitoringv1.PrometheusRule {
	return &monitoringv1.PrometheusRule{
		TypeMeta: metav1.TypeMeta{
			APIVersion: monitoringv1.SchemeGroupVersion.String(),
			Kind:       "PrometheusRule",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      ruleName,
			Namespace: namespace,
		},
		Spec: *NewPrometheusRuleSpec(),
	}
}

// NewPrometheusRuleSpec creates PrometheusRuleSpec for alerts and recording rules
func NewPrometheusRuleSpec() *monitoringv1.PrometheusRuleSpec {
	return &monitoringv1.PrometheusRuleSpec{
		Groups: []monitoringv1.RuleGroup{{
			Name: alertRuleGroup,
			Rules: []monitoringv1.Rule{
				createDeploymentSizeUndesiredAlertRule(),
				createOperatorDownAlertRule(),
				createOperatorUpTotalRecordingRule(),
			},
		}},
	}
}

// createDeploymentSizeUndesiredAlertRule creates MonitorDeploymentSizeUndesired alert rule
func createDeploymentSizeUndesiredAlertRule() monitoringv1.Rule {
	return monitoringv1.Rule{
		Alert: deploymentSizeUndesiredAlert,
		Expr:  intstr.FromString("increase(monitor_deployment_size_undesired_count_total[5m]) >= 3"),
		Annotations: map[string]string{
			"description": "Monitor-sample deployment size was not as desired more than 3 times in the last 5 minutes.",
		},
		Labels: map[string]string{
			"severity": "warning",
		},
	}
}

// createOperatorDownAlertRule creates MonitorOperatorDown alert rule
func createOperatorDownAlertRule() monitoringv1.Rule {
	return monitoringv1.Rule{
		Alert: operatorDownAlert,
		Expr:  intstr.FromString("monitor_operator_up_total == 0"),
		Annotations: map[string]string{
			"description": "No running monitor-operator pods were detected in the last 5 min.",
		},
		For: "5m",
		Labels: map[string]string{
			"severity": "critical",
		},
	}
}

// createOperatorUpTotalRecordingRule creates monitor_operator_up_total recording rule
func createOperatorUpTotalRecordingRule() monitoringv1.Rule {
	return monitoringv1.Rule{
		Record: operatorUpTotalRecordingRule,
		Expr:   intstr.FromString("sum(up{pod=~'k8s-operator-monitoring-controller-manager-.*'} or vector(0))"),
	}
}
