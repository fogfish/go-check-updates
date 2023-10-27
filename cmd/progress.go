//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/go-check-updates
//

package cmd

import (
	"time"

	"github.com/schollz/progressbar/v3"
)

func progress() *progressbar.ProgressBar {
	ch := make(chan struct{})

	bar := progressbar.NewOptions(-1,
		progressbar.OptionSpinnerType(14),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionOnCompletion(
			func() {
				ch <- struct{}{}
			},
		),
	)

	go func() {
		for {
			select {
			case <-ch:
				close(ch)
				return
			case <-time.After(40 * time.Millisecond):
				bar.Add(1)
			}
		}
	}()

	return bar
}
