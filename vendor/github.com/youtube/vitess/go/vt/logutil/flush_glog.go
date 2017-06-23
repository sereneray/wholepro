// Copyright 2013, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logutil

import (
	log "github.com/golang/glog"
)

func init() {
	OnFlush(log.Flush)
}
