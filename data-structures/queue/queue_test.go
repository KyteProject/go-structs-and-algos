package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := New()

	t.Run("Check if queue is empty", func(t *testing.T) {
		got := q.IsEmpty()
		want := true
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Test Enqueue and Peek", func(t *testing.T) {
		q.Enqueue(56)
		q.Enqueue(124)
		q.Enqueue(3)
		q.Enqueue(673)

		want := 56
		got := q.Peek()
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Test queue size", func(t *testing.T) {
		want := 4
		got := q.Len()
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Test Dequeue", func(t *testing.T) {
		q.Dequeue()

		want := 124
		got := q.Dequeue()
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Test Clear", func(t *testing.T) {
		q.Clear()

		want := 0
		got := q.Len()
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
