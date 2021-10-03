package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestItem_Next(t *testing.T) {
	first := Item{value: 1}
	second := Item{prev: &first, value: 2}
	first.next = &second
	require.Equal(t, &second, first.Next())
}

func TestItem_Prev(t *testing.T) {
	first := Item{value: 1}
	second := Item{prev: &first, value: 2}
	first.next = &second
	require.Equal(t, &first, second.Prev())
}

func TestItem_Value(t *testing.T) {
	item := Item{value: 1}
	require.Equal(t, 1, item.Value())
}

func TestList_First(t *testing.T) {
	item := Item{value: 1}
	list := List{
		first: &item,
	}
	require.Equal(t, &item, list.First())
}

func TestList_Last(t *testing.T) {
	item := Item{value: 1}
	list := List{
		last: &item,
	}
	require.Equal(t, &item, list.Last())
}

func TestList_Len(t *testing.T) {
	list := List{
		length: 1,
	}
	require.Equal(t, 1, list.Len())
}

func TestList_PushBack(t *testing.T) {
	list := List{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	require.Equal(t, 3, list.Len())
	require.Equal(t, 1, list.First().Value())
	require.Equal(t, 3, list.Last().Value())
}

func TestList_PushFront(t *testing.T) {
	list := List{}
	list.PushFront(3)
	list.PushFront(2)
	list.PushFront(1)
	require.Equal(t, 3, list.Len())
	require.Equal(t, 1, list.First().Value())
	require.Equal(t, 3, list.Last().Value())
}

func TestList_Remove(t *testing.T) {
	list := List{}
	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)
	list.Remove(list.First())
	list.Remove(list.Last())
	require.Equal(t, 1, list.Len())
	require.Equal(t, 2, list.First().Value())
	require.Equal(t, 2, list.Last().Value())
}
