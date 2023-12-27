package misc

import "fmt"

/*
The question was in Python
For a map that supports the following operations:
- get(self, k) -> v or keyerror
- put(self, k, v)
- delete(self, k)

Implement (including the above methods):
- take_snapshot(self) -> returns snapshot ID (int)

Enhance:
- get(self, k, snapshot_id=None) -> v from snapshot

Aim is for the testcases to pass

Constraints:
- To save space, the snapshots should not contain what has not changed or is unchanged
- All operations should be sub-linear to the size of the keys in the map

Assume to begin with you have 10 million keys in the map and 100 snapshots
only 1% of the values change between snapshots
*/

/*
Implementation:
In addition to our map, we also have an additional map,
which as its keys stores snapshot IDs and value a delta
at that snapshot from the previous one.

To lookup a value at each snapshot we traverse back
snapshots (since we only store the delta i.e. "changes" at
any given snapshot).
*/

// T is comparable, comparable implements == and != only
// we need that for map key finding
type SnapshotMap[T, V comparable] struct {
	Map   map[T]V
	Snaps []map[int]map[T]V
	ID    int
	Diff  map[int]map[T]V
}

func NewSnapshotMap[T, V comparable]() *SnapshotMap[T, V] {
	return &SnapshotMap[T, V]{
		Map: map[T]V{},
		Snaps: []map[int]map[T]V{
			nil,
		},
		ID: 1,
		Diff: map[int]map[T]V{
			0: {},
			1: {},
		},
	}
}

func (s *SnapshotMap[T, V]) TakeSnapshot() int {
	retVal := s.ID
	s.Snaps = append(s.Snaps, s.Diff)
	s.Diff = map[int]map[T]V{
		0: {},
		1: {},
	}
	s.ID++
	return retVal
}

func (s *SnapshotMap[T, V]) Get(k T, snapID int) (V, error) {
	var val V // How else to return the default value for a type we don't know?
	var err error
	var found bool
	if snapID == 0 {
		if val, found = s.Map[k]; !found {
			err = fmt.Errorf("key not found")
		}
	} else {
		lookupID := snapID
		for !found && lookupID > 0 {
			if _, deleted := s.Snaps[lookupID][0][k]; deleted {
				break
			} else {
				val, found = s.Snaps[lookupID][1][k]
			}
			lookupID--
		}
		if !found {
			err = fmt.Errorf("key not found")
		}
	}
	return val, err
}

func (s *SnapshotMap[T, V]) Put(k T, v V) {
	if val, found := s.Map[k]; found && val == v {
		return
	}
	s.Map[k], s.Diff[1][k] = v, v
}

func (s *SnapshotMap[T, V]) Delete(k T) {
	if val, found := s.Map[k]; found {
		delete(s.Map, k)
		s.Diff[0][k] = val
	}
}
