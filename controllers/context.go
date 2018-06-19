package controllers

import (
	r "gopkg.in/gorethink/gorethink.v4"
	"squeezecnn/common"
)

type Context struct {
	RethinkSession *r.Session
}

func (c *Context) Close() {
	c.RethinkSession.Close()
}
var context *Context

func NewContext() {
	session := common.GetSession()
	context = &Context{
		RethinkSession: session,
	}
}

func GetContext() *Context {
	return context
}