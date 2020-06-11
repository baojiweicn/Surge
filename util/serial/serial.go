// Package serial implements a wrapper around serial for fake
// purpose.
//
// It's very simple to works with a fake or real serial port.
//
// Example
//
// Here is a simple example for realworld serial port operations:
//
//	options := serial.Options{ PortName: "/dev/ttyS3", Fake: false }
//	port, err := serial.Open(options)
//	// handle err
//	// Write
//	port.Write([]byte("Some bytes"))
//	// Read
//	go func() {
//		b := make([]byte, 40)
//		n, err := port.Read()
//		// handle err
//		// handle bytes read in b
//	}()
//
// Here is a simple example to fake a serial port remote peer:
//
//	remote := serial.NewFakeRemotePeer()
//	options := serial.Options{
//		PortName: "/dev/ttyS3",
//		Fake: true,
//		FakeRemotePeer: remote,
//	}
//	port, err := serial.Open(options)
//	// handle err
//	// Write
//	port.Write([]byte("Ping"))
//	// Read
//	go func() {
//		b := make([]byte, 40)
//		n, err := port.Read()
//		// handle err
//		// handle bytes read in b
//	}()
//
//	// Remote end read and writes fakings.
//	go func() {
//		b := make([]byte, 40)
//		n, err := remote.Read()
//		// handle err
//		// We received the data writes from `port`.
//	}()
//
//	remote.Write([]byte("Pong"))
//
package serial

import (
	"io"

	serial "github.com/baojiweicn/Surge/go-serial/serial"
)

// Alias
const (
	PARITY_NONE = serial.PARITY_NONE
	PARITY_ODD  = serial.PARITY_ODD
	PARITY_EVEN = serial.PARITY_EVEN
)

// Options is a wrapper around serial.OpenOptions.
type Options struct {
	// Whether to use fake serial.
	Fake bool
	// The remote fake peer.
	FakeRemotePeer *FakeRemotePeer

	// The name of the port, e.g. "/dev/tty.usbserial-A8008HlV".
	PortName string

	// The baud rate for the port.
	BaudRate uint

	// The number of data bits per frame. Legal values are 5, 6, 7, and 8.
	DataBits uint

	// The number of stop bits per frame. Legal values are 1 and 2.
	StopBits uint

	// The type of parity bits to use for the connection. Currently parity errors
	// are simply ignored; that is, bytes are delivered to the user no matter
	// whether they were received with a parity error or not.
	ParityMode serial.ParityMode

	// Enable RTS/CTS (hardware) flow control.
	RTSCTSFlowControl     bool
	InterCharacterTimeout uint
	MinimumReadSize       uint
}

// Opens a serial port with given options.
func Open(options Options) (io.ReadWriteCloser, error) {
	if !options.Fake {
		return serial.Open(serial.OpenOptions{
			PortName:              options.PortName,
			BaudRate:              options.BaudRate,
			DataBits:              options.DataBits,
			StopBits:              options.StopBits,
			ParityMode:            options.ParityMode,
			RTSCTSFlowControl:     options.RTSCTSFlowControl,
			InterCharacterTimeout: options.InterCharacterTimeout,
			MinimumReadSize:       options.MinimumReadSize,
		})
	}
	return &FakePort{remote: options.FakeRemotePeer}, nil
}

// FakePort is a fake serial port handle implementation.
// Which implements the interface io.ReadWriteCloser.
type FakePort struct {
	remote *FakeRemotePeer
}

// Read implements io.Reader.
func (r *FakePort) Read(b []byte) (int, error) {
	return r.remote.readOut(b)
}

// Write implements io.Writer.
func (r *FakePort) Write(b []byte) (int, error) {
	return r.remote.writeIn(b)
}

// Close implements io.Closer.
func (r *FakePort) Close() error { return nil }

type FakeRemotePeer struct {
	// Buffer In.
	ib chan byte
	// Buffer Out.
	ob chan byte
}

// NewFakeRemotePeer returns a new FakeRemotePeer.
func NewFakeRemotePeer() *FakeRemotePeer {
	return &FakeRemotePeer{
		ib: make(chan byte, 4*1024),
		ob: make(chan byte, 4*1024),
	}
}

// writeIn writes given bytes b to the input buffer of this peer.
func (r *FakeRemotePeer) writeIn(b []byte) (int, error) {
	for _, v := range b {
		r.ib <- v
	}
	return len(b), nil
}

// Read reads bytes into given bytes b from remote serial port.
func (r *FakeRemotePeer) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = <-r.ib
	}
	return len(b), nil
}

// readOut reads given bytes into b to from the output buffer of this peer.
func (r *FakeRemotePeer) readOut(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = <-r.ob
	}
	return len(b), nil
}

// Write writes given bytes b to remote serial port.
func (r *FakeRemotePeer) Write(b []byte) (int, error) {
	for _, v := range b {
		r.ob <- v
	}
	return len(b), nil
}
