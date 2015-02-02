// Copyright (C) 2014 Jakob Borg and Contributors (see the CONTRIBUTORS file).
// All rights reserved. Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package discosrv

type address struct {
	ip   []byte
	port uint16
	seen int64 // epoch seconds
}

type addressList struct {
	addresses []address
}
