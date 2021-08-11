// Copyright (c) 2021, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package siftool

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_command_getNew(t *testing.T) {
	tests := []struct {
		name string
		opts commandOpts
	}{
		{
			name: "OK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tf, err := ioutil.TempFile("", "sif-test-*")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tf.Name())
			tf.Close()

			c := &command{opts: tt.opts}

			cmd := c.getNew()

			runCommand(t, cmd, []string{tf.Name()})
		})
	}
}
