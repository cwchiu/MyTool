package server

import (
    ftpserver "github.com/fclairamb/ftpserver/server"
	"github.com/spf13/cobra"
    "errors"
    "crypto/tls"
    "os"
    "io"
    "log"
    "time"
	"bytes"
	"io/ioutil"
	"net/http"
)

// MainDriver defines a very basic serverftp driver
type MainDriver struct {
	baseDir   string
    port int
    username string
    password string
	tlsConfig *tls.Config
}

func (driver *MainDriver) WelcomeUser(cc ftpserver.ClientContext) (string, error) {
	return "Welcome on Arick's FTP Server", nil
}

func (driver *MainDriver) AuthUser(cc ftpserver.ClientContext, user, pass string) (ftpserver.ClientHandlingDriver, error) {
	if user == driver.username && pass == driver.password {
        return driver, nil
	}
    
    return nil, errors.New("Bad username or password")
}

func (driver *MainDriver) GetSettings() *ftpserver.Settings {
	var config ftpserver.Settings
    config.MaxConnections = 10
    config.ListenHost = "0.0.0.0"
    config.ListenPort = driver.port

	return &config
}

// GetTLSConfig returns a TLS Certificate to use
func (driver *MainDriver) GetTLSConfig() (*tls.Config, error) {
	if driver.tlsConfig == nil {
		log.Println("Loading certificate")
		if cert, err := tls.LoadX509KeyPair("sample/certs/mycert.crt", "sample/certs/mycert.key"); err == nil {
			driver.tlsConfig = &tls.Config{
				NextProtos:   []string{"ftp"},
				Certificates: []tls.Certificate{cert},
			}
		} else {
			return nil, err
		}
	}
	return driver.tlsConfig, nil
}

// ChangeDirectory changes the current working directory
func (driver *MainDriver) ChangeDirectory(cc ftpserver.ClientContext, directory string) error {
	if directory == "/debug" {
		cc.SetDebug(!cc.Debug())
		return nil
	} else if directory == "/virtual" {
		return nil
	}
	_, err := os.Stat(driver.baseDir + directory)
	return err
}

// MakeDirectory creates a directory
func (driver *MainDriver) MakeDirectory(cc ftpserver.ClientContext, directory string) error {
	return os.Mkdir(driver.baseDir+directory, 0777)
}

// ListFiles lists the files of a directory
func (driver *MainDriver) ListFiles(cc ftpserver.ClientContext) ([]os.FileInfo, error) {

	if cc.Path() == "/virtual" {
		files := make([]os.FileInfo, 0)
		files = append(files,
			virtualFileInfo{
				name: "localpath.txt",
				mode: os.FileMode(0666),
				size: 1024,
			},
			virtualFileInfo{
				name: "file2.txt",
				mode: os.FileMode(0666),
				size: 2048,
			},
		)
		return files, nil
	}

	path := driver.baseDir + cc.Path()

	files, err := ioutil.ReadDir(path)

	// We add a virtual dir
	if cc.Path() == "/" && err == nil {
		files = append(files, virtualFileInfo{
			name: "virtual",
			mode: os.FileMode(0666) | os.ModeDir,
			size: 4096,
		})
	}

	return files, err
}

// UserLeft is called when the user disconnects, even if he never authenticated
func (driver *MainDriver) UserLeft(cc ftpserver.ClientContext) {

}

// OpenFile opens a file in 3 possible modes: read, write, appending write (use appropriate flags)
func (driver *MainDriver) OpenFile(cc ftpserver.ClientContext, path string, flag int) (ftpserver.FileStream, error) {

	if path == "/virtual/localpath.txt" {
		return &virtualFile{content: []byte(driver.baseDir)}, nil
	}

	path = driver.baseDir + path

	// If we are writing and we are not in append mode, we should remove the file
	if (flag & os.O_WRONLY) != 0 {
		flag |= os.O_CREATE
		if (flag & os.O_APPEND) == 0 {
			os.Remove(path)
		}
	}

	return os.OpenFile(path, flag, 0666)
}

// GetFileInfo gets some info around a file or a directory
func (driver *MainDriver) GetFileInfo(cc ftpserver.ClientContext, path string) (os.FileInfo, error) {
	path = driver.baseDir + path

	return os.Stat(path)
}

// CanAllocate gives the approval to allocate some data
func (driver *MainDriver) CanAllocate(cc ftpserver.ClientContext, size int) (bool, error) {
	return true, nil
}

// ChmodFile changes the attributes of the file
func (driver *MainDriver) ChmodFile(cc ftpserver.ClientContext, path string, mode os.FileMode) error {
	path = driver.baseDir + path

	return os.Chmod(path, mode)
}

// DeleteFile deletes a file or a directory
func (driver *MainDriver) DeleteFile(cc ftpserver.ClientContext, path string) error {
	path = driver.baseDir + path

	return os.Remove(path)
}

// RenameFile renames a file or a directory
func (driver *MainDriver) RenameFile(cc ftpserver.ClientContext, from, to string) error {
	from = driver.baseDir + from
	to = driver.baseDir + to

	return os.Rename(from, to)
}


func NewMyDriver(port int, username string, password string) *MainDriver {
	driver := &MainDriver{
		baseDir: ".",
        username: username,
        password: password,
        port: port,
	}
	os.MkdirAll(driver.baseDir, 0777)
	return driver
}

type virtualFile struct {
	content    []byte // Content of the file
	readOffset int    // Reading offset
}

func (f *virtualFile) Close() error {
	return nil
}

func (f *virtualFile) Read(buffer []byte) (int, error) {
	n := copy(buffer, f.content[f.readOffset:])
	f.readOffset += n
	if n == 0 {
		return 0, io.EOF
	}

	return n, nil
}

func (f *virtualFile) Seek(n int64, w int) (int64, error) {
	return 0, nil
}

func (f *virtualFile) Write(buffer []byte) (int, error) {
	return 0, nil
}

type virtualFileInfo struct {
	name string
	size int64
	mode os.FileMode
}

func (f virtualFileInfo) Name() string {
	return f.name
}

func (f virtualFileInfo) Size() int64 {
	return f.size
}

func (f virtualFileInfo) Mode() os.FileMode {
	return f.mode
}

func (f virtualFileInfo) IsDir() bool {
	return f.mode.IsDir()
}

func (f virtualFileInfo) ModTime() time.Time {
	return time.Now().UTC()
}

func (f virtualFileInfo) Sys() interface{} {
	return nil
}

func externalIP() (string, error) {
	// If you need to take a bet, amazon is about as reliable & sustainable a service as you can get
	rsp, err := http.Get("http://checkip.amazonaws.com")
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes.TrimSpace(buf)), nil
}


func SetupFtpCommand(rootCmd *cobra.Command) {
	var port int
    var username string
    var password string
	cmd := &cobra.Command{
		Use:   "ftp",
		Short: "ftp Server",
		Run: func(cmd *cobra.Command, args []string) {
            if username == "" {
                username = "guest"
            }
            
            if password == "" {
                password = "1234"
            }
            fserver := ftpserver.NewFtpServer(NewMyDriver(port, username, password))
            log.Printf("Username=%s, Password=%s\n", username, password)
            err := fserver.ListenAndServe()
            if err != nil {
                panic(err)
            }
		},
	}
	cmd.Flags().IntVarP(&port, "port", "p", 21, "listen port")
	cmd.Flags().StringVarP(&username, "username", "u", "arick", "username")
	cmd.Flags().StringVarP(&password, "password", "w", "1234", "password")
	rootCmd.AddCommand(cmd)
}
