/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package common

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeToTimestamp(t *testing.T) {

	count := 60
	year := 2018
	month := 2

	for i := 0; i < count; i++ {
		t := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
		timestamp := t.Unix()
		fmt.Println(timestamp)

		month++
		if month > 12 {
			month = 1
			year++
		}
	}
}
