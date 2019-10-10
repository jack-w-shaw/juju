// Copyright 2012, 2013 Canonical Ltd.
// Copyright 2014 Cloudbase Solutions SRL
// Licensed under the AGPLv3, see LICENCE file for details.

package jujuc_test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/juju/cmd"
	"github.com/juju/cmd/cmdtesting"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/worker/uniter/runner/jujuc"
	"github.com/juju/juju/worker/uniter/runner/jujuc/jujuctesting"
)

type RelationGetSuite struct {
	relationSuite
}

var _ = gc.Suite(&RelationGetSuite{})

func (s *RelationGetSuite) newHookContext(relid int, remote string, app string) (jujuc.Context, *relationInfo) {
	hctx, info := s.relationSuite.newHookContext(relid, remote, app)
	info.rels[0].Units["u/0"]["private-address"] = "foo: bar\n"
	info.rels[1].SetRelated("m/0", jujuctesting.Settings{"pew": "pew\npew\n"})
	info.rels[1].SetRelated("u/1", jujuctesting.Settings{"value": "12345"})
	return hctx, info
}

var relationGetTests = []struct {
	summary     string
	relid       int
	unit        string
	args        []string
	code        int
	out         string
	key         string
	application bool
}{
	{
		summary: "no default relation",
		relid:   -1,
		code:    2,
		out:     `no relation id specified`,
	}, {
		summary: "explicit relation, not known",
		relid:   -1,
		code:    2,
		args:    []string{"-r", "burble:123"},
		out:     `invalid value "burble:123" for option -r: relation not found`,
	}, {
		summary: "default relation, no unit chosen",
		relid:   1,
		code:    2,
		out:     `no unit id specified`,
	}, {
		summary: "explicit relation, no unit chosen",
		relid:   -1,
		code:    2,
		args:    []string{"-r", "burble:1"},
		out:     `no unit id specified`,
	}, {
		summary: "missing key",
		relid:   1,
		unit:    "m/0",
		args:    []string{"ker-plunk"},
	}, {
		summary: "missing unit",
		relid:   1,
		unit:    "bad/0",
		code:    1,
		out:     `unknown unit bad/0`,
	}, {
		summary: "all keys with explicit non-member",
		relid:   1,
		args:    []string{"-", "u/1"},
		out:     `value: "12345"`,
	}, {
		summary: "specific key with implicit member",
		relid:   1,
		unit:    "m/0",
		args:    []string{"pew"},
		out:     "pew\npew\n",
	}, {
		summary: "specific key with explicit member",
		relid:   1,
		args:    []string{"pew", "m/0"},
		out:     "pew\npew\n",
	}, {
		summary: "specific key with explicit non-member",
		relid:   1,
		args:    []string{"value", "u/1"},
		out:     "12345",
	}, {
		summary: "specific key with explicit local",
		relid:   0,
		args:    []string{"private-address", "u/0"},
		out:     "foo: bar\n",
	}, {
		summary: "all keys with implicit member",
		relid:   1,
		unit:    "m/0",
		out:     "pew: |\n  pew\n  pew",
	}, {
		summary: "all keys with explicit member",
		relid:   1,
		args:    []string{"-", "m/0"},
		out:     "pew: |\n  pew\n  pew",
	}, {
		summary: "all keys with explicit local",
		relid:   0,
		args:    []string{"-", "u/0"},
		out:     "private-address: |\n  foo: bar",
	}, {
		summary: "explicit smart formatting 1",
		relid:   1,
		unit:    "m/0",
		args:    []string{"--format", "smart"},
		out:     "pew: |\n  pew\n  pew",
	}, {
		summary: "explicit smart formatting 2",
		relid:   1,
		unit:    "m/0",
		args:    []string{"pew", "--format", "smart"},
		out:     "pew\npew\n",
	}, {
		summary: "explicit smart formatting 3",
		relid:   1,
		args:    []string{"value", "u/1", "--format", "smart"},
		out:     "12345",
	}, {
		summary: "explicit smart formatting 4",
		relid:   1,
		args:    []string{"missing", "u/1", "--format", "smart"},
		out:     "",
	},
}

func (s *RelationGetSuite) TestRelationGet(c *gc.C) {
	for i, t := range relationGetTests {
		c.Logf("test %d: %s", i, t.summary)
		hctx, _ := s.newHookContext(t.relid, t.unit, "")
		com, err := jujuc.NewCommand(hctx, cmdString("relation-get"))
		c.Assert(err, jc.ErrorIsNil)
		ctx := cmdtesting.Context(c)
		code := cmd.Main(jujuc.NewJujucCommandWrappedForTest(com), ctx, t.args)
		c.Check(code, gc.Equals, t.code)
		if code == 0 {
			c.Check(bufferString(ctx.Stderr), gc.Equals, "")
			expect := t.out
			if len(expect) > 0 {
				expect += "\n"
			}
			c.Check(bufferString(ctx.Stdout), gc.Equals, expect)
		} else {
			c.Check(bufferString(ctx.Stdout), gc.Equals, "")
			expect := fmt.Sprintf(`(.|\n)*ERROR %s\n`, t.out)
			c.Check(bufferString(ctx.Stderr), gc.Matches, expect)
		}
	}
}

var relationGetFormatTests = []struct {
	summary string
	relid   int
	unit    string
	args    []string
	out     interface{}
}{
	{
		summary: "formatting 1",
		relid:   1,
		unit:    "m/0",
		out:     map[string]interface{}{"pew": "pew\npew\n"},
	}, {
		summary: "formatting 2",
		relid:   1,
		unit:    "m/0",
		args:    []string{"pew"},
		out:     "pew\npew\n",
	}, {
		summary: "formatting 3",
		relid:   1,
		args:    []string{"value", "u/1"},
		out:     "12345",
	}, {
		summary: "formatting 4",
		relid:   1,
		args:    []string{"missing", "u/1"},
		out:     nil,
	},
}

func (s *RelationGetSuite) TestRelationGetFormat(c *gc.C) {
	testFormat := func(format string, checker gc.Checker) {
		for i, t := range relationGetFormatTests {
			c.Logf("test %d: %s %s", i, format, t.summary)
			hctx, _ := s.newHookContext(t.relid, t.unit, "")
			com, err := jujuc.NewCommand(hctx, cmdString("relation-get"))
			c.Assert(err, jc.ErrorIsNil)
			ctx := cmdtesting.Context(c)
			args := append(t.args, "--format", format)
			code := cmd.Main(jujuc.NewJujucCommandWrappedForTest(com), ctx, args)
			c.Check(code, gc.Equals, 0)
			c.Check(bufferString(ctx.Stderr), gc.Equals, "")
			stdout := bufferString(ctx.Stdout)
			c.Check(stdout, checker, t.out)
		}
	}
	testFormat("yaml", jc.YAMLEquals)
	testFormat("json", jc.JSONEquals)
}

var helpTemplate = `
Usage: %s

Summary:
get relation settings

Options:
--app  (= false)
    Get the relation data for the overall application, not just a unit
--format  (= smart)
    Specify output format (json|smart|yaml)
-o, --output (= "")
    Specify an output file
-r, --relation  (= %s)
    Specify a relation by id

Details:
relation-get prints the value of a unit's relation setting, specified by key.
If no key is given, or if the key is "-", all keys and values will be printed.

A unit can see its own settings by calling "relation-get - MYUNIT", this will include
any changes that have been made with "relation-set".

When reading remote relation data, a charm can call relation-get --app - to get
the data for the application data bag that is set by the remote applications
leader.
%s`[1:]

var relationGetHelpTests = []struct {
	summary string
	relid   int
	unit    string
	usage   string
	rel     string
}{
	{
		summary: "no default relation",
		relid:   -1,
		usage:   "relation-get [options] <key> <unit id>",
	}, {
		summary: "no default unit",
		relid:   1,
		usage:   "relation-get [options] <key> <unit id>",
		rel:     "peer1:1",
	}, {
		summary: "default unit",
		relid:   1,
		unit:    "any/1",
		usage:   `relation-get [options] [<key> [<unit id>]]`,
		rel:     "peer1:1",
	},
}

func (s *RelationGetSuite) TestHelp(c *gc.C) {
	for i, t := range relationGetHelpTests {
		c.Logf("test %d", i)
		hctx, _ := s.newHookContext(t.relid, t.unit, "")
		com, err := jujuc.NewCommand(hctx, cmdString("relation-get"))
		c.Assert(err, jc.ErrorIsNil)
		ctx := cmdtesting.Context(c)
		code := cmd.Main(jujuc.NewJujucCommandWrappedForTest(com), ctx, []string{"--help"})
		c.Assert(code, gc.Equals, 0)
		unitHelp := ""
		if t.unit != "" {
			unitHelp = fmt.Sprintf("Current default unit id is %q.\n", t.unit)
		}
		expect := fmt.Sprintf(helpTemplate, t.usage, t.rel, unitHelp)
		c.Assert(bufferString(ctx.Stdout), gc.Equals, expect)
		c.Assert(bufferString(ctx.Stderr), gc.Equals, "")
	}
}

func (s *RelationGetSuite) TestOutputPath(c *gc.C) {
	hctx, _ := s.newHookContext(1, "m/0", "")
	com, err := jujuc.NewCommand(hctx, cmdString("relation-get"))
	c.Assert(err, jc.ErrorIsNil)
	ctx := cmdtesting.Context(c)
	code := cmd.Main(jujuc.NewJujucCommandWrappedForTest(com), ctx, []string{"--output", "some-file", "pew"})
	c.Assert(code, gc.Equals, 0)
	c.Assert(bufferString(ctx.Stderr), gc.Equals, "")
	c.Assert(bufferString(ctx.Stdout), gc.Equals, "")
	content, err := ioutil.ReadFile(filepath.Join(ctx.Dir, "some-file"))
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(string(content), gc.Equals, "pew\npew\n\n")
}

type relationGetInitTest struct {
	summary     string
	ctxrelid    int
	ctxunit     string
	ctxapp      string
	args        []string
	content     string
	err         string
	relid       int
	key         string
	unit        string
	application bool
}

func (t relationGetInitTest) log(c *gc.C, i int) {
	var summary string
	if t.summary != "" {
		summary = " - " + t.summary
	}
	c.Logf("test %d%s", i, summary)
}

func (t relationGetInitTest) init(c *gc.C, s *RelationGetSuite) (cmd.Command, []string) {
	args := make([]string, len(t.args))
	copy(args, t.args)

	hctx, _ := s.newHookContext(t.ctxrelid, t.ctxunit, t.ctxapp)
	com, err := jujuc.NewCommand(hctx, cmdString("relation-get"))
	c.Assert(err, jc.ErrorIsNil)

	return com, args
}

func (t relationGetInitTest) check(c *gc.C, com cmd.Command, err error) {
	if t.err == "" {
		if !c.Check(err, jc.ErrorIsNil) {
			return
		}

		rset := com.(*jujuc.RelationGetCommand)
		c.Check(rset.RelationId, gc.Equals, t.relid)
		c.Check(rset.Key, gc.Equals, t.key)
		c.Check(rset.UnitName, gc.Equals, t.unit)
		c.Check(rset.Application, gc.Equals, t.application)
	} else {
		c.Check(err, gc.ErrorMatches, t.err)
	}
}

var relationGetInitTests = []relationGetInitTest{
	{
		summary:  "no relation id",
		ctxrelid: -1,
		err:      `no relation id specified`,
	}, {
		summary:  "invalid relation id",
		ctxrelid: -1,
		args:     []string{"-r", "one"},
		err:      `invalid value "one" for option -r: invalid relation id`,
	}, {
		summary:  "invalid relation id with builtin context relation id",
		ctxrelid: 1,
		args:     []string{"-r", "one"},
		err:      `invalid value "one" for option -r: invalid relation id`,
	}, {
		summary:  "relation not found",
		ctxrelid: -1,
		args:     []string{"-r", "2"},
		err:      `invalid value "2" for option -r: relation not found`,
	}, {
		summary:  "-r overrides context relation id",
		ctxrelid: 1,
		ctxunit:  "u/0",
		unit:     "u/0",
		args:     []string{"-r", "ignored:0"},
		relid:    0,
	}, {
		summary:  "key=value for relation-get (maybe should be invalid?)",
		ctxrelid: 1,
		relid:    1,
		ctxunit:  "u/0",
		unit:     "u/0",
		args:     []string{"key=value"},
		key:      "key=value",
	}, {
		summary:  "key supplied",
		ctxrelid: 1,
		relid:    1,
		ctxunit:  "u/0",
		unit:     "u/0",
		args:     []string{"key"},
		key:      "key",
	}, {
		summary: "magic key supplied",
		ctxunit: "u/0",
		unit:    "u/0",
		args:    []string{"-"},
		key:     "",
	}, {
		summary: "override ctxunit with explicit unit",
		ctxunit: "u/0",
		args:    []string{"key", "u/1"},
		key:     "key",
		unit:    "u/1",
	}, {
		summary: "magic key with unit",
		ctxunit: "u/0",
		args:    []string{"-", "u/1"},
		key:     "",
		unit:    "u/1",
	}, {
		summary:     "supply --app",
		ctxunit:     "u/0",
		unit:        "u/0",
		args:        []string{"--app"},
		application: true,
	}, {
		summary:     "supply --app and app name",
		ctxunit:     "u/0",
		unit:        "u",
		args:        []string{"--app", "-", "u"},
		application: true,
	}, {
		summary: "app name but no context unit name",
		ctxunit: "",
		ctxapp:  "u",
		unit:    "u",
	}, {
		/// 		ctxrelid: 0,
		/// 		args:     []string{"-r", "1", "foo=bar"},
		/// 		relid:    1,
		/// 		settings: map[string]string{"foo": "bar"},
		/// 	}, {
		/// 		ctxrelid: 1,
		/// 		args:     []string{"foo=123", "bar=true", "baz=4.5", "qux="},
		/// 		relid:    1,
		/// 		settings: map[string]string{"foo": "123", "bar": "true", "baz": "4.5", "qux": ""},
		/// 	}, {
		/// 		summary:  "file with a valid setting",
		/// 		args:     []string{"--file", "spam"},
		/// 		content:  "{foo: bar}",
		/// 		settings: map[string]string{"foo": "bar"},
		/// 	}, {
		/// 		summary:  "file with multiple settings on a line",
		/// 		args:     []string{"--file", "spam"},
		/// 		content:  "{foo: bar, spam: eggs}",
		/// 		settings: map[string]string{"foo": "bar", "spam": "eggs"},
		/// 	}, {
		/// 		summary:  "file with multiple lines",
		/// 		args:     []string{"--file", "spam"},
		/// 		content:  "{\n  foo: bar,\n  spam: eggs\n}",
		/// 		settings: map[string]string{"foo": "bar", "spam": "eggs"},
		/// 	}, {
		/// 		summary:  "an empty file",
		/// 		args:     []string{"--file", "spam"},
		/// 		content:  "",
		/// 		settings: map[string]string{},
		/// 	}, {
		/// 		summary:  "an empty map",
		/// 		args:     []string{"--file", "spam"},
		/// 		content:  "{}",
		/// 		settings: map[string]string{},
		/// 	}, {
		/// 		summary: "accidental same format as command-line",
		/// 		args:    []string{"--file", "spam"},
		/// 		content: "foo=bar ham=eggs good=bad",
		/// 		err:     "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `foo=bar...` into map.*",
		/// 	}, {
		/// 		summary: "scalar instead of map",
		/// 		args:    []string{"--file", "spam"},
		/// 		content: "haha",
		/// 		err:     "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `haha` into map.*",
		/// 	}, {
		/// 		summary: "sequence instead of map",
		/// 		args:    []string{"--file", "spam"},
		/// 		content: "[haha]",
		/// 		err:     "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!seq into map.*",
		/// 	}, {
		/// 		summary:  "multiple maps",
		/// 		args:     []string{"--file", "spam"},
		/// 		content:  "{a: b}\n{c: d}",
		/// 		settings: map[string]string{"a": "b"},
		/// 	}, {
		/// 		summary:  "value with a space",
		/// 		args:     []string{"--file", "spam"},
		/// 		content:  "{foo: 'bar baz'}",
		/// 		settings: map[string]string{"foo": "bar baz"},
		/// 	}, {
		/// 		summary:  "value with an equal sign",
		/// 		args:     []string{"--file", "spam"},
		/// 		content:  "{foo: foo=bar, base64: YmFzZTY0IGV4YW1wbGU=}",
		/// 		settings: map[string]string{"foo": "foo=bar", "base64": "YmFzZTY0IGV4YW1wbGU="},
		/// 	}, {
		/// 		summary:  "values with brackets",
		/// 		args:     []string{"--file", "spam"},
		/// 		content:  "{foo: '[x]', bar: '{y}'}",
		/// 		settings: map[string]string{"foo": "[x]", "bar": "{y}"},
		/// 	}, {
		/// 		summary:  "a messy file",
		/// 		args:     []string{"--file", "spam"},
		/// 		content:  "\n {  \n # a comment \n\n  \nfoo: bar,  \nham: eggs,\n\n  good: bad,\nup: down, left: right\n}\n",
		/// 		settings: map[string]string{"foo": "bar", "ham": "eggs", "good": "bad", "up": "down", "left": "right"},
		/// 	}, {
		/// 		summary:  "file + settings",
		/// 		args:     []string{"--file", "spam", "foo=bar"},
		/// 		content:  "{ham: eggs}",
		/// 		settings: map[string]string{"ham": "eggs", "foo": "bar"},
		/// 	}, {
		/// 		summary:  "file overridden by settings",
		/// 		args:     []string{"--file", "spam", "foo=bar"},
		/// 		content:  "{foo: baz}",
		/// 		settings: map[string]string{"foo": "bar"},
		/// 	}, {
		/// 		summary:  "read from stdin",
		/// 		args:     []string{"--file", "-"},
		/// 		content:  "{foo: bar}",
		/// 		settings: map[string]string{"foo": "bar"},
		/// 	}, {
		/// 		summary:     "pass --app",
		/// 		args:        []string{"--app", "baz=qux"},
		/// 		settings:    map[string]string{"baz": "qux"},
		/// 		application: true,
		ctxunit: "u/0",
		unit:    "u/0",
	},
}

func (s *RelationGetSuite) TestInit(c *gc.C) {
	for i, t := range relationGetInitTests {
		t.log(c, i)
		com, args := t.init(c, s)

		err := cmdtesting.InitCommand(com, args)
		t.check(c, com, err)
	}
}
