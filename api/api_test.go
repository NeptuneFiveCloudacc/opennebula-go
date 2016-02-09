package api

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type S struct {
	cm map[string]string
}

var _ = check.Suite(&S{})

func (s *S) SetUpSuite(c *check.C) {
	cm := make(map[string]string)
	cm[ENDPOINT] = "http://103.56.92.4:2633/RPC2"
	cm[USERID] = "oneadmin"
	cm[PASSWORD] = "yourWuOtHij3"
	s.cm = cm
}


func (s *S) TestCreateClient(c *check.C) {
	_, error := NewClient(s.cm)
	c.Assert(error, check.IsNil)
}

func (s *S) TestCall(c *check.C) {
	c1, err := NewClient(s.cm)
	c.Assert(err, check.IsNil)
	args := []interface{}{c1.Key, -2, 3, 3}
	_, err = c1.Call("one.templatepool.info", args)
	c.Assert(err, check.IsNil)
}
