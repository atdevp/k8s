package config

import (
	"sync"

	"k8s.io/apimachinery/pkg/util/sets"
)

// SourceReadyFn is func.
type SourceReadyFn func(SourceSeen sets.String) bool

// SourceReady track the set of configured sources seen by kubelet.
type SourceReady interface {
	AddSource(source string)
	AllReady() bool
}

// NewSourceReady .
func NewSourceReady(sourceReadyFn SourceReadyFn) SourceReady {
	return &SourcesImpl{
		sourcesSeen:   sets.NewString(),
		sourceReadyFn: sourceReadyFn,
	}
}

// SourcesImpl struc
type SourcesImpl struct {
	lock          sync.RWMutex
	sourcesSeen   sets.String
	sourceReadyFn SourceReadyFn
}

// AddSource specified source to the set of sources managed.
func (s *SourcesImpl) AddSource(source string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.sourcesSeen.Insert(source)
}

// AllReady return true if each configured source is ready.
func (s *SourcesImpl) AllReady() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.sourceReadyFn(s.sourcesSeen)
}
