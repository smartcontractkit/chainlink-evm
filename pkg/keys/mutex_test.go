package keys

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMutex_LockUnlock(t *testing.T) {
	rm := &Mutex{}

	err := rm.TryLock(TXMv1)
	require.NoError(t, err)

	err = rm.Unlock(TXMv1)
	require.NoError(t, err)
}

func TestMutex_LockByDifferentServiceType(t *testing.T) {
	rm := &Mutex{}

	err := rm.TryLock(TXMv1)
	require.NoError(t, err)

	err = rm.TryLock(TXMv2)
	require.Error(t, err)
	require.Equal(t, "resource is locked by another service type", err.Error())
}

func TestMutex_UnlockWithoutLock(t *testing.T) {
	rm := &Mutex{}

	err := rm.Unlock(TXMv1)
	require.Error(t, err)
	require.Equal(t, "no active lock", err.Error())

	require.NoError(t, rm.TryLock(TXMv1))
	err = rm.Unlock(TXMv2)
	require.Error(t, err)
	require.Equal(t, "no active lock for this service type", err.Error())
}

func TestMutex_MultipleLocks(t *testing.T) {
	rm := &Mutex{}

	err := rm.TryLock(TXMv1)
	require.NoError(t, err)

	err = rm.TryLock(TXMv1)
	require.NoError(t, err)

	err = rm.Unlock(TXMv1)
	require.NoError(t, err)

	err = rm.Unlock(TXMv1)
	require.NoError(t, err)
}

func TestIsLocked_WhenResourceIsLockedByServiceType(t *testing.T) {
	rm := &Mutex{serviceType: TXMv1, count: 1}

	locked := rm.IsLocked(TXMv1)
	require.True(t, locked)
}

func TestIsLocked_WhenResourceIsNotLockedByServiceType(t *testing.T) {
	rm := &Mutex{}

	locked := rm.IsLocked(TXMv1)
	require.False(t, locked)
}

func TestIsLocked_WhenResourceIsLockedByDifferentServiceType(t *testing.T) {
	rm := &Mutex{serviceType: TXMv2, count: 1}

	locked := rm.IsLocked(TXMv1)
	require.False(t, locked)
}
