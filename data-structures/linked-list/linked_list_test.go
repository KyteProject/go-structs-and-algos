package list

import (
	"testing"
)

func TestLinkedList(t *testing.T) {

	l := NewList()

	t.Run("Check if list exists and is empty", func(t *testing.T) {
		got := l.IsEmpty()
		want := true
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Add items to beginning of list", func(t *testing.T) {
		l.Add('a')
		l.Add('b')
		l.Add('c')
		l.Add('d')
		l.Add('e')
	})

	t.Run("Test Head", func(t *testing.T) {
		want := 'e'
		got := l.PeekFirst()
		if want != got {
			t.Errorf("Head: got %v want %v", got, want)
		}
	})

	t.Run("Test Next", func(t *testing.T) {
		want := 'd'
		got := l.PeekNext()
		if want != got {
			t.Errorf("Next: got %v want %v", got, want)
		}
	})

	t.Run("Test Tail", func(t *testing.T) {
		want := 'a'
		got := l.PeekLast()
		if want != got {
			t.Errorf("Tail: got %v want %v", got, want)
		}
	})

	t.Run("Add items to end of list", func(t *testing.T) {
		l.AddLast('f')
		l.AddLast('g')

		want := 'g'
		got := l.PeekLast()
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Test removing first node", func(t *testing.T) {
		want := l.PeekFirst()
		got, _ := l.RemoveFirst()
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}

		want = 'd'
		got = l.PeekFirst()
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Test removing the last node", func(t *testing.T) {
		want := l.PeekLast()
		got, _ := l.RemoveLast()
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}

		want = 'f'
		got = l.PeekLast()
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Test removing by Value", func(t *testing.T) {
		want := 'b'
		got, _ := l.RemoveByValue('b')
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Test clearing list", func(t *testing.T) {
		l.Clear()

		want := true
		got := l.IsEmpty()
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
