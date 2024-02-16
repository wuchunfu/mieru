// Copyright (C) 2022  mieru authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package util

import (
	"context"
	"net"
	"time"
)

// SetReadTimeout set read deadline to the connection.
// It cancels the deadline if the timeout is 0 or negative.
func SetReadTimeout(conn net.Conn, timeout time.Duration) {
	if timeout > 0 {
		conn.SetReadDeadline(time.Now().Add(timeout))
	} else {
		conn.SetReadDeadline(ZeroTime())
	}
}

// WaitForClose blocks the go routine. It returns when the peer closes the connection.
// In the meanwhile, everything send by the peer is discarded.
func WaitForClose(conn net.Conn) {
	b := make([]byte, 64)
	for {
		_, err := conn.Read(b)
		if err != nil {
			return
		}
	}
}

// SendReceive sends a request to the connection and returns the response.
// The maxinum size of response is 4096 bytes.
func SendReceive(ctx context.Context, conn net.Conn, req []byte) (resp []byte, err error) {
	_, err = conn.Write(req)
	if err != nil {
		return
	}

	resp = make([]byte, 4096)
	var n int
	n, err = conn.Read(resp)
	resp = resp[:n]
	return
}
