// Copyright © 2019 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/duration"
	"knative.dev/client/pkg/printers"
	"knative.dev/pkg/apis"
)

func Age(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return duration.ShortHumanDuration(time.Now().Sub(t))
}

// Color the type of the conditions
func formatConditionType(condition apis.Condition) string {
	return string(condition.Type)
}

// Status in ASCII format
func formatStatus(c apis.Condition) string {
	switch c.Status {
	case corev1.ConditionTrue:
		return "++"
	case corev1.ConditionFalse:
		switch c.Severity {
		case apis.ConditionSeverityError:
			return "!!"
		case apis.ConditionSeverityWarning:
			return " W"
		case apis.ConditionSeverityInfo:
			return " I"
		default:
			return " !"
		}
	default:
		return "??"
	}
}

// Used for conditions table to do own formatting for the table,
// as the tabbed writer doesn't work nicely with colors
func getMaxTypeLen(conditions []apis.Condition) int {
	max := 0
	for _, condition := range conditions {
		if len(condition.Type) > max {
			max = len(condition.Type)
		}
	}
	return max
}

// Sort conditions: Ready first, followed by error, then Warning, then Info
func sortConditions(conditions []apis.Condition) []apis.Condition {
	// Don't change the orig slice
	ret := make([]apis.Condition, 0, len(conditions))
	for _, c := range conditions {
		ret = append(ret, c)
	}
	sort.SliceStable(ret, func(i, j int) bool {
		ic := &ret[i]
		jc := &ret[j]
		// Ready first
		if ic.Type == apis.ConditionReady {
			return jc.Type != apis.ConditionReady
		}
		// Among conditions of the same Severity, sort by Type
		if ic.Severity == jc.Severity {
			return ic.Type < jc.Type
		}
		// Error < Warning < Info
		switch ic.Severity {
		case apis.ConditionSeverityError:
			return true
		case apis.ConditionSeverityWarning:
			return jc.Severity == apis.ConditionSeverityInfo
		case apis.ConditionSeverityInfo:
			return false
		default:
			return false
		}
		return false
	})
	return ret
}

// Print out a table with conditions. Use green for 'ok', and red for 'nok' if color is enabled
func WriteConditions(dw printers.PrefixWriter, conditions []apis.Condition, printMessage bool) {
	dw.WriteColsLn(printers.Level0, "Conditions:")
	conditions = sortConditions(conditions)
	maxLen := getMaxTypeLen(conditions)
	formatHeader := "%-2s %-" + strconv.Itoa(maxLen) + "s %6s %-s\n"
	formatRow := "%-2s %-" + strconv.Itoa(maxLen) + "s %6s %-s\n"
	dw.Write(printers.Level1, formatHeader, "OK", "TYPE", "AGE", "REASON")
	for _, condition := range conditions {
		ok := formatStatus(condition)
		reason := condition.Reason
		if printMessage && reason != "" {
			reason = fmt.Sprintf("%s (%s)", reason, condition.Message)
		}
		dw.Write(printers.Level1, formatRow, ok, formatConditionType(condition), Age(condition.LastTransitionTime.Inner.Time), reason)
	}
}
