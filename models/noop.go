package models

import (
	v1 "k8s.io/api/core/v1"
)

const NotifierNoop = "noop"

type Noop struct {
	Calls int
}

func (t *Noop) NotifyPodTermination(pod v1.Pod) error {
	t.Calls++
	return nil
}
