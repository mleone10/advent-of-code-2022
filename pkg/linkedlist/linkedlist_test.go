package linkedlist_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/pkg/linkedlist"
)

func TestNewNode(t *testing.T) {
	n := linkedlist.NewNode(1)
	assert.Equal(t, n.Next(), nil)
	assert.Equal(t, n.Prev(), nil)
	assert.Equal(t, n.Value(), 1)
}

func TestLinkNext(t *testing.T) {
	n := linkedlist.NewNode(1)
	m := linkedlist.NewNode(2)
	n.LinkNext(m)
	assert.Equal(t, n.Next(), m)
	assert.Equal(t, n.Prev(), nil)
	assert.Equal(t, m.Prev(), n)
	assert.Equal(t, m.Next(), nil)
}

func TestLinkPrev(t *testing.T) {
	n := linkedlist.NewNode(1)
	m := linkedlist.NewNode(2)
	n.LinkPrev(m)
	assert.Equal(t, n.Prev(), m)
	assert.Equal(t, n.Next(), nil)
	assert.Equal(t, m.Next(), n)
	assert.Equal(t, m.Prev(), nil)
}

func TestLength(t *testing.T) {
	head := linkedlist.NewNode(0)
	n := head
	for i := 0; i < 10; i++ {
		m := linkedlist.NewNode(i)
		n.LinkNext(m)
		n = m
	}
	assert.Equal(t, head.Length(), 11)
}
