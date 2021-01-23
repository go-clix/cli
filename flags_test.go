package cli

import (
	"reflect"
	"testing"
)

func TestStripFlags(t *testing.T) {
	tests := []struct {
		input  []string
		output []string
	}{
		{
			[]string{"foo", "bar"},
			[]string{"foo", "bar"},
		},
		{
			[]string{"foo", "--str", "-s"},
			[]string{"foo"},
		},
		{
			[]string{"-s", "foo", "--str", "bar"},
			[]string{},
		},
		{
			[]string{"-i10", "echo"},
			[]string{"echo"},
		},
		{
			[]string{"-i=10", "echo"},
			[]string{"echo"},
		},
		{
			[]string{"--int=100", "echo"},
			[]string{"echo"},
		},
		{
			[]string{"-ib", "echo", "-sfoo", "baz"},
			[]string{"echo", "baz"},
		},
		{
			[]string{"-i=baz", "bar", "-i", "foo", "blah"},
			[]string{"bar", "blah"},
		},
		{
			[]string{"--int=baz", "-sbar", "-i", "foo", "blah"},
			[]string{"blah"},
		},
		{
			[]string{"--bool", "bar", "-i", "foo", "blah"},
			[]string{"bar", "blah"},
		},
		{
			[]string{"-b", "bar", "-i", "foo", "blah"},
			[]string{"bar", "blah"},
		},
		{
			[]string{"--persist", "bar"},
			[]string{"bar"},
		},
		{
			[]string{"-p", "bar"},
			[]string{"bar"},
		},
	}

	c := &Command{Use: "c"}
	c.Flags().BoolP("persist", "p", false, "")
	c.Flags().IntP("int", "i", -1, "")
	c.Flags().StringP("str", "s", "", "")
	c.Flags().BoolP("bool", "b", false, "")

	for _, test := range tests {
		got := stripFlags(test.input, c)
		if !reflect.DeepEqual(test.output, got) {
			t.Fatalf("want %+v but got %+v", got, test.output)
		}
	}
}

func TestPersistentFlags(t *testing.T) {
	parent := &Command{}
	parent.PersistentFlags().String("persistent", "", "")
	parent.Flags().String("non-persistent", "", "")
	child := &Command{}
	parent.AddCommand(child)
	if child.PersistentFlags().Lookup("persistent") == nil {
		t.Error("expected persistent flag to be passed to child")
	}
	if child.Flags().Lookup("non-persistent") != nil {
		t.Error("expected non-persistent flag to not be passed to child")
	}
}
