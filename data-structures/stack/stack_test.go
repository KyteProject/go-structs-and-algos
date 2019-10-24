package stack

import (
	"testing"
)

// Test new, push, pop, peek, isEmpty, depth

func TestStack(t *testing.T) {
	s := New()

	t.Run("Check if stack is empty", func(t *testing.T) {
		got := s.IsEmpty()
		want := true
		if want != got {
			t.Errorf("got %v want %v given, %v", got, want, s)
		}
	})

	t.Run("Push items onto stack, check correct order", func(t *testing.T) {
		s.Push('a')
		s.Push('b')
		s.Push('c')

		want := 'c'
		got := s.stack[0]
		if want != got {
			t.Errorf("got %v want %v given, %v", got, want, s.stack)
		}

		want = 'b'
		got = s.stack[1]
		if want != got {
			t.Errorf("got %v want %v given, %v", got, want, s.stack)
		}

		want = 'a'
		got = s.stack[2]
		if want != got {
			t.Errorf("got %v want %v given, %v", got, want, s.stack)
		}
	})

	t.Run("Peek top of stack", func(t *testing.T) {
		want := 'c'
		got := s.Peek()
		if want != got && s.depth != 3 {
			t.Errorf("got %v want %v given, %v", got, want, s.stack)
		}
	})

	t.Run("Pop items off stack", func(t *testing.T) {

		want := 'c'
		got := s.Pop()
		if want != got {
			t.Errorf("got %v want %v given, %v", got, want, s.stack)
		}

		want = 'b'
		got = s.Pop()
		if want != got {
			t.Errorf("got %v want %v given, %v", got, want, s.stack)
		}

		want = 'a'
		got = s.Pop()
		if want != got {
			t.Errorf("got %v want %v given, %v", got, want, s.stack)
		}
	})
}
