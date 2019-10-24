package list

import "testing"

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

		want := 'e'
		got := l.head.data
		if want != got {
			t.Errorf("Head: got %v want %v", got, want)
		}

		want = 'd'
		got = l.head.next.data
		if want != got {
			t.Errorf("Next: got %v want %v", got, want)
		}

		want = 'a'
		got = l.tail.data
		if want != got {
			t.Errorf("Tail: got %v want %v", got, want)
		}
	})

	t.Run("Add items to end of list", func(t *testing.T) {
		l.AddLast('f')
		l.AddLast('g')

		want = 'g'
		got = l.tail.data
		if want != got {
			t.Errorf("Tail: got %v want %v", got, want)
		}
	})
}
