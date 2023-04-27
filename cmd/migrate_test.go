package cmd_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/dnitsch/git-local-util/cmd"
)

func Test_migrate_command_successfully_supplied_all_args(t *testing.T) {
	b := new(bytes.Buffer)

	cmd := cmd.GluCmd

	cmd.SetArgs([]string{"migrate", "-d", "../test", "-f", "bar", "-r", "foo"})
	cmd.SetErr(b)
	cmd.Execute()
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	// 	//
	if len(out) > 0 {
		t.Errorf(`%s
	got: %v
	wanted: ""`, "expected empty buffer", string(out))
	}
}

func Test_verbose(t *testing.T) {
	b := new(bytes.Buffer)
	stdout := new(bytes.Buffer)

	cmd := cmd.GluCmd

	cmd.SetArgs([]string{"migrate", "-d", "../test", "-f", "bar", "-r", "foo", "-v"})
	cmd.SetErr(b)
	cmd.SetOut(stdout)
	cmd.Execute()
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	if len(out) > 0 {
		t.Errorf(`%s
	got: %v
	wanted: ""`, "expected empty buffer", string(out))
	}
}
