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

package walletnode

import (
	"errors"
)

// RestartWalletnode restart walletnode
func (w *WalletnodeManager) RestartWalletnode(symbol string) error {

	return errors.New("Function closed! Use stop/start, please")

	// if err := loadConfig(symbol); err != nil {
	// 	return err
	// }

	// // Init docker client
	// c, err := getDockerClient(symbol)
	// if err != nil {
	// 	return err
	// }
	// // Action within client
	// cName, err := getCName(symbol) // container name
	// if err != nil {
	// 	return err
	// }
	// err = c.ContainerRestart(context.Background(), cName, nil)
	// if err != nil {
	// 	return err
	// }
	// return nil
}
