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

package bitcoin

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestDecimalShit(t *testing.T) {
	num, _ := decimal.NewFromString("0.00005")
	num2 := num.Shift(-1)
	t.Logf("balance: %v\n", num2)
}

func orderHash(origins []string, addr string, start int) []string {
	fmt.Printf("find addr: %v\n", addr)
	fmt.Printf("origins: %v\n", origins)
	newHashs := make([]string, start)
	copy(newHashs, origins[:start])
	end := 0
	for i := start; i < len(origins); i++ {
		txAddr := origins[i]
		if txAddr == addr {
			newHashs = append(newHashs, txAddr)
			end = i
			break
		}
	}

	fmt.Printf("head: %v\n", newHashs)
	fmt.Printf("front: %v\n", origins[start:end])
	fmt.Printf("behind: %v\n", origins[end+1:])

	newHashs = append(newHashs, origins[start:end]...)
	newHashs = append(newHashs, origins[end+1:]...)
	return newHashs
}

func TestOrderHash(t *testing.T) {
	origins := []string{
		"c", "a", "b", "a", "b", "c", "b", "c", "a",
	}

	confused := []string{
		"b", "c", "a", "a", "c", "b", "b", "c", "a",
	}

	for i, w := range origins {
		confused = orderHash(confused, w, i)
	}

	fmt.Println(confused)
}
