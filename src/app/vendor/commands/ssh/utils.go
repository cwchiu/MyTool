package ssh

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"gopkg.in/cheggaaa/pb.v1"
	"io"
	"net"
	"os"
	"path"
	"time"
)

func connect(user, password, host string, port int) (*ssh.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	return sshClient, nil
}

func newftpclient(user, password, host string, port int) (*sftp.Client, error) {
	sshClient, err := connect(user, password, host, port)
	if err != nil {
		return nil, err
	}
	// create sftp client
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return nil, err
	}

	return sftpClient, nil
}

func upload(sftpClient *sftp.Client, local_fn string, remote_dir string) error {
	srcFile, err := os.Open(local_fn)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	remoteFileName := path.Base(local_fn)
	dstFile, err := sftpClient.Create(path.Join(remote_dir, remoteFileName))
	if err != nil {
		return err
	}
	defer dstFile.Close()
	finfo, err := srcFile.Stat()
	if err != nil {
		return err
	}
	bar := pb.New64(finfo.Size()).SetUnits(pb.U_BYTES)
	bar.Start()
	buf_size := 1024
	buf := make([]byte, buf_size)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
		bar.Add(buf_size)
		// bar.Increment()
	}

	return nil
}

func download(sftpClient *sftp.Client, remote_fn string, local_dir string) error {
	srcFile, err := sftpClient.Open(remote_fn)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	finfo, err := srcFile.Stat()
	if err != nil {
		return err
	}

	localFileName := path.Base(remote_fn)
	dstFile, err := os.Create(path.Join(local_dir, localFileName))
	if err != nil {
		return err
	}
	defer dstFile.Close()

	bar := pb.New64(finfo.Size()).SetUnits(pb.U_BYTES)
	bar.Start()
	// writer := bar.NewProxyWriter(dstFile)
	reader := bar.NewProxyReader(srcFile)
	_, err = io.Copy(dstFile, reader)
	if err != nil {
		return err
	}
	// if _, err = srcFile.WriteTo(writer); err != nil {
	// if _, err = srcFile.WriteTo(dstFile); err != nil {
	// return err
	// }

	return nil
}
