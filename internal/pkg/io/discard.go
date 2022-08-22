package io

import "io"

func Discard(src io.Reader) error {
	if _, err := io.Copy(io.Discard, src); err != nil {
		return err
	}
	if closer, is := src.(io.Closer); is {
		return closer.Close()
	}
	return nil
}

func WillClose(closer io.Closer, handle func(error)) {
	handle(closer.Close())
}

func WillDiscard(src io.Reader, handle func(error)) {
	handle(Discard(src))
}
