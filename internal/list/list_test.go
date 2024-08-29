package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPushRemove(t *testing.T) {
	ll := NewList[int]()
	length := 10

	for i := 1; i <= length; i++ {
		ll.InsertNewNode(i)
		require.Equal(t, ll.Len(), i)
	}

	for i := length; i <= 1; i++ {
		ll.Remove(ll.GetLastNode())
		require.Equal(t, ll.Len(), i)
	}
}

func TestInsertNode(t *testing.T) {
	ll := NewList[int]()
	e := ll.InsertNewNode(0)
	ll.InsertNewNode(1)
	require.Equal(t, e, ll.GetLastNode())
}

func TestInit(t *testing.T) {
	ll := NewList[int]()

	ll.InsertNewNode(1)
	require.Equal(t, ll.Len(), 1)

	ll.Init()
	require.Equal(t, ll.Len(), 0)
}
