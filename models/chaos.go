package models

import (
	"regexp"
	"time"

	log "github.com/sirupsen/logrus"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
)

type Chaos struct {
	// a kubernetes client object
	Client kubernetes.Interface
	// a label selector which restricts the pods to choose from
	Labels labels.Selector
	// an annotation selector which restricts the pods to choose from
	Annotations labels.Selector
	// a kind label selector which restricts the kinds to choose from
	Kinds labels.Selector
	// a namespace selector which restricts the pods to choose from
	Namespaces labels.Selector
	// a namespace label selector which restricts the namespaces to choose from
	NamespaceLabels labels.Selector
	// a regular expression for pod names to include
	IncludedPodNames *regexp.Regexp
	// a regular expression for pod names to exclude
	ExcludedPodNames *regexp.Regexp
	// a list of weekdays when termination is suspended
	ExcludedWeekdays []time.Weekday
	// a list of time periods of a day when termination is suspended
	ExcludedTimesOfDay []TimePeriod
	// a list of days of a year when termination is suspended
	ExcludedDaysOfYear []time.Time
	// the timezone to apply when detecting the current weekday
	Timezone *time.Location
	// minimum age of pods to consider
	MinimumAge time.Duration
	// an instance of logrus.StdLogger to write log messages to
	Logger log.FieldLogger
	// a terminator that terminates victim pods
	Terminator Terminator
	// dry run will not allow any pod terminations
	DryRun bool
	// grace period to terminate the pods
	GracePeriod time.Duration
	// event recorder allows to publish events to Kubernetes
	EventRecorder record.EventRecorder
	// a function to retrieve the current time
	Now func() time.Time

	MaxKill int
	// chaos events notifier
	Notifier Notifier
	// namespace scope for the Kubernetes client
	ClientNamespaceScope string
}