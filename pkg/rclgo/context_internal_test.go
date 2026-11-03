package rclgo

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// fakeResource closes like a real resource: it removes itself from its store
// and may close another resource it owns.
type fakeResource struct {
	rosID
	store  *rosResourceStore
	owned  *fakeResource
	closes atomic.Int32
}

func (f *fakeResource) Close() error {
	if f.closes.Add(1) > 1 {
		return closeErr("fake resource")
	}
	f.store.removeResource(f)
	if f.owned != nil {
		// Like WaitSet.Close, tolerate the owned resource being closed first.
		var closeError closeError
		if err := f.owned.Close(); err != nil && !errors.As(err, &closeError) {
			return err
		}
	}
	return nil
}

func newFakeResource(store *rosResourceStore) *fakeResource {
	f := &fakeResource{store: store}
	store.addResource(f)
	return f
}

func TestResourceStoreCloseSkipsResourcesClosedByOthers(t *testing.T) {
	var store rosResourceStore
	owners := make([]*fakeResource, 20)
	for i := range owners {
		owners[i] = newFakeResource(&store)
		owners[i].owned = newFakeResource(&store)
	}
	// Without the skip, the store re-closes owned resources after their owner
	// did and reports the double close.
	require.NoError(t, store.Close())
	for _, o := range owners {
		assert.Equal(t, int32(1), o.closes.Load())
		assert.Positive(t, o.owned.closes.Load())
	}
	assert.Empty(t, store.snapshot())
}

func TestResourceStoreCloseConcurrentWithAdd(t *testing.T) {
	// Run with -race: Close must not iterate the map while another goroutine
	// adds to it.
	var store rosResourceStore
	for range 100 {
		newFakeResource(&store)
	}
	var wg sync.WaitGroup
	wg.Go(func() {
		for range 1000 {
			newFakeResource(&store)
		}
	})
	require.NoError(t, store.Close())
	wg.Wait()
}
