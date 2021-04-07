package policy_rule_count

import (
	"github.com/kyverno/kyverno/pkg/metrics"
)

type PolicyRuleCountMetricChangeType string

const (
	PolicyRuleCreated PolicyRuleCountMetricChangeType = "created"
	PolicyRuleDeleted PolicyRuleCountMetricChangeType = "deleted"
)

type PromMetrics metrics.PromMetrics
