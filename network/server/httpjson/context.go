// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package httpjson

import (
	m3dbcontext "github.com/m3db/m3db/context"

	apachethrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/uber/tchannel-go/thrift"
	"golang.org/x/net/context"
)

const (
	contextKey = "m3dbcontext"
)

// NewDefaultContextFn returns a function that will create M3DB contexts per request
func NewDefaultContextFn(contextPool m3dbcontext.Pool) ContextFn {
	return func(ctx context.Context, method string, headers map[string]string) thrift.Context {
		ctxWithValue := context.WithValue(ctx, interface{}(contextKey), contextPool.Get())
		return thrift.WithHeaders(ctxWithValue, headers)
	}
}

// DefaulPostResponseFn will close M3DB contexts per request
func DefaulPostResponseFn(ctx context.Context, method string, response apachethrift.TStruct) {
	value := ctx.Value(contextKey)
	inner := value.(m3dbcontext.Context)
	inner.Close()
}
