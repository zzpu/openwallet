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
	"archive/tar"
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"

	"docker.io/go-docker/api/types"
)

// CopyFromContainer copy file from container to local filesystem
//
//	src := "wallet.dat"  // 备份来源，全节点中的文件名 (MainDataPath + '/' + src)
//	dst := "2018.....wallet.dat" // 备份目标，自设
//	src/dst: filename
func (w *WalletnodeManager) CopyFromContainer(symbol, src, dst string) error {

	var buf bytes.Buffer

	if err := loadConfig(symbol); err != nil {
		return err
	}

	// Init docker client
	c, err := getDockerClient(symbol)
	if err != nil {
		return err
	}

	cname, err := getCName(symbol)
	if err != nil {
		return err
	}

	// dataDir, err := WNConfig.getDataDir()
	// if err != nil {
	// 	return err
	// }
	// src = filepath.Join(dataDir, src)

	// API Return: CopyFromContainer -> (io.ReadCloser, types.ContainerPathStat, error)
	fp, _, err := c.CopyFromContainer(context.Background(), cname, src)
	if err != nil {
		return err
	}
	defer fp.Close()

	tw := tar.NewReader(fp)
	for {
		// Copy file from container return within archive/tar
		_, err := tw.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		// log.Printf("Contents of %s:\n", hdr.Name)

		if _, err := buf.ReadFrom(tw); err != nil {
			log.Fatal(err)
			return err
		}
	}

	if err = ioutil.WriteFile(dst, buf.Bytes(), 0600); err != nil {
		return err
	}

	return nil
}

// CopyToContainer copy file to container from local filesystem
//
// Define:
//	src: filename
//	dst: path
// Example:
//	src := "/tmp/2018......wallet.dat"  // 恢复来源，用户提供
//	dst := "wallet.dat" // 恢复目标的文件名 (MainDataPath + '/' + dst)
func (w *WalletnodeManager) CopyToContainer(symbol, src, dst string) error {

	var content io.Reader

	if err := loadConfig(symbol); err != nil {
		return err
	}

	// Init docker client
	c, err := getDockerClient(symbol)
	if err != nil {
		return err
	}

	cname, err := getCName(symbol)
	if err != nil {
		return err
	}

	// dataDir, err := WNConfig.getDataDir()
	// if err != nil {
	// 	return err
	// }
	// dst = filepath.Join(dataDir, dst)

	// Return: ioutil.ReadFile() -> ([]byte, err)
	if dat, err := ioutil.ReadFile(src); err != nil {
		log.Println(err)
		return err
	} else {
		// Copy file into container within archive/tar
		var buf bytes.Buffer
		tw := tar.NewWriter(&buf)
		tw.WriteHeader(&tar.Header{
			Name: filepath.Base(src), //file.Name,
			Mode: 0600,
			Size: int64(len(dat)), //int64(len(file.Body)),
		})
		tw.Write([]byte(dat))
		tw.Close()

		// Transform tar to []byte as Reader for Docker API
		content = bytes.NewReader(buf.Bytes())
	}

	// API Params: (ctx context.Context, container, path string, content io.Reader, options types.CopyToContainerOptions)
	if err := c.CopyToContainer(context.Background(), cname, dst, content, types.CopyToContainerOptions{AllowOverwriteDirWithFile: false}); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
