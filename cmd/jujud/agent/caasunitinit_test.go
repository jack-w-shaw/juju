// Copyright 2019 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package agent

import (
	"bytes"
	"math/rand"
	"net"
	"os"
	"runtime"
	"sync"
	"time"

	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/cmd/cmdtesting"
	"github.com/juju/juju/juju/sockets"
	coretesting "github.com/juju/juju/testing"
	"github.com/juju/juju/worker/uniter/runner/jujuc"
	"github.com/juju/testing"
)

type CAASUnitInitSuite struct {
	coretesting.BaseSuite

	rootDir string
}

var _ = gc.Suite(&CAASUnitInitSuite{})

func (s *CAASUnitInitSuite) SetUpTest(c *gc.C) {
	if runtime.GOOS == "windows" {
		c.Skip("unsupported")
	}
	s.BaseSuite.SetUpTest(c)
}

func (s *CAASUnitInitSuite) newCommand(c *gc.C, st *testing.Stub) *CAASUnitInitCommand {
	cmd := NewCAASUnitInitCommand()
	cmd.copyFunc = func(src, dst string) error {
		st.AddCall("Copy", src, dst)
		return st.NextErr()
	}
	cmd.symlinkFunc = func(src, dst string) error {
		st.AddCall("Symlink", src, dst)
		return st.NextErr()
	}
	cmd.removeAllFunc = func(path string) error {
		st.AddCall("RemoveAll", path)
		return st.NextErr()
	}
	cmd.mkdirAllFunc = func(path string, mode os.FileMode) error {
		st.AddCall("MkdirAll", path, mode)
		return st.NextErr()
	}
	return cmd
}

func (s *CAASUnitInitSuite) checkCommand(c *gc.C, cmd *CAASUnitInitCommand, args []string,
	unit string, operatorFile string,
	operatorCACertFile string, charmDir string) []testing.StubCall {
	ctx, err := cmdtesting.RunCommand(c, cmd, args...)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(ctx, gc.NotNil)

	toolsPath := "/var/lib/juju/tools/" + unit
	agentPath := "/var/lib/juju/agents/" + unit

	calls := []testing.StubCall{
		{FuncName: "RemoveAll", Args: []interface{}{toolsPath}},
		{FuncName: "MkdirAll", Args: []interface{}{toolsPath, os.FileMode(0775)}},
		{FuncName: "RemoveAll", Args: []interface{}{agentPath}},
		{FuncName: "MkdirAll", Args: []interface{}{agentPath, os.FileMode(0775)}},
		{FuncName: "Symlink", Args: []interface{}{"/var/lib/juju/tools/jujud", toolsPath + "/jujud"}},
	}
	for _, cmdName := range jujuc.CommandNames() {
		_ = cmdName
		calls = append(calls,
			testing.StubCall{FuncName: "Symlink", Args: []interface{}{"/var/lib/juju/tools/jujud", toolsPath + "/" + cmdName}})
	}

	calls = append(calls,
		testing.StubCall{FuncName: "Copy", Args: []interface{}{operatorFile, agentPath + "/operator-client.yaml"}},
		testing.StubCall{FuncName: "Copy", Args: []interface{}{operatorCACertFile, agentPath + "/ca.crt"}},
		testing.StubCall{FuncName: "Copy", Args: []interface{}{charmDir, agentPath + "/charm"}})

	return calls
}

func (s *CAASUnitInitSuite) TestInitUnit(c *gc.C) {
	args := []string{"--unit", "unit-wow-0",
		"--operator-file", "operator/file/path",
		"--operator-ca-cert-file", "operator/cert/file/path",
		"--charm-dir", "charm/dir"}
	st := &testing.Stub{}
	cmd := s.newCommand(c, st)
	calls := s.checkCommand(c, cmd, args, "unit-wow-0",
		"operator/file/path", "operator/cert/file/path", "charm/dir")
	st.CheckCalls(c, calls)
}

func (s *CAASUnitInitSuite) TestInitUnitWaitSend(c *gc.C) {
	socketName := "@" + string(rand.Int63())
	listening := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		st := &testing.Stub{}
		cmd := s.newCommand(c, st)
		cmd.socketName = socketName
		cmd.listenFunc = func(s sockets.Socket) (net.Listener, error) {
			l, err := sockets.Listen(s)
			close(listening)
			return l, err
		}
		calls := s.checkCommand(c, cmd, []string{"--wait"}, "unit-wow-0",
			"operator/file/path", "operator/cert/file/path", "charm/dir")
		st.CheckCalls(c, calls)
	}()

	select {
	case <-listening:
	case <-time.After(coretesting.LongWait):
		c.Fatal("failed to listen")
	}

	stdErr := &bytes.Buffer{}
	args := []string{"--send", "--unit", "unit-wow-0",
		"--operator-file", "operator/file/path",
		"--operator-ca-cert-file", "operator/cert/file/path",
		"--charm-dir", "charm/dir"}
	st := &testing.Stub{}
	cmd := s.newCommand(c, st)
	cmd.stdErr = stdErr
	cmd.socketName = socketName
	ctx, err := cmdtesting.RunCommand(c, cmd, args...)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(ctx, gc.NotNil)
	c.Assert(stdErr.Bytes(), gc.Not(gc.HasLen), 0)

	wg.Wait()
}
