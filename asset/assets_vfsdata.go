package asset

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

var Assets = func() http.FileSystem {
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{fs["/static"].(os.FileInfo), fs["/templates"].(os.FileInfo)}
	fs["/static"].(*vfsgen۰DirInfo).entries = []os.FileInfo{fs["/static/favicon.ico"].(os.FileInfo), fs["/static/index.html"].(os.FileInfo), fs["/static/lib"].(os.FileInfo), fs["/static/script.js"].(os.FileInfo)}
	fs["/static/lib"].(*vfsgen۰DirInfo).entries = []os.FileInfo{fs["/static/lib/bootstrap-4.0.0-alpha.6-dist"].(os.FileInfo), fs["/static/lib/font-awesome-4.7.0"].(os.FileInfo)}
	fs["/static/lib/bootstrap-4.0.0-alpha.6-dist"].(*vfsgen۰DirInfo).entries = []os.FileInfo{fs["/static/lib/bootstrap-4.0.0-alpha.6-dist/css"].(os.FileInfo)}
	fs["/static/lib/bootstrap-4.0.0-alpha.6-dist/css"].(*vfsgen۰DirInfo).entries = []os.FileInfo{fs["/static/lib/bootstrap-4.0.0-alpha.6-dist/css/bootstrap.min.css"].(os.FileInfo), fs["/static/lib/bootstrap-4.0.0-alpha.6-dist/css/bootstrap.min.css.map"].(os.FileInfo)}
	fs["/static/lib/font-awesome-4.7.0"].(*vfsgen۰DirInfo).entries = []os.FileInfo{fs["/static/lib/font-awesome-4.7.0/css"].(os.FileInfo), fs["/static/lib/font-awesome-4.7.0/fonts"].(os.FileInfo)}
	fs["/static/lib/font-awesome-4.7.0/css"].(*vfsgen۰DirInfo).entries = []os.FileInfo{fs["/static/lib/font-awesome-4.7.0/css/font-awesome.css"].(os.FileInfo), fs["/static/lib/font-awesome-4.7.0/css/font-awesome.min.css"].(os.FileInfo)}
	fs["/static/lib/font-awesome-4.7.0/fonts"].(*vfsgen۰DirInfo).entries = []os.FileInfo{fs["/static/lib/font-awesome-4.7.0/fonts/FontAwesome.otf"].(os.FileInfo), fs["/static/lib/font-awesome-4.7.0/fonts/fontawesome-webfont.eot"].(os.FileInfo), fs["/static/lib/font-awesome-4.7.0/fonts/fontawesome-webfont.svg"].(os.FileInfo), fs["/static/lib/font-awesome-4.7.0/fonts/fontawesome-webfont.ttf"].(os.FileInfo), fs["/static/lib/font-awesome-4.7.0/fonts/fontawesome-webfont.woff"].(os.FileInfo), fs["/static/lib/font-awesome-4.7.0/fonts/fontawesome-webfont.woff2"].(os.FileInfo)}
	fs["/templates"].(*vfsgen۰DirInfo).entries = []os.FileInfo{fs["/templates/default.tmpl"].(os.FileInfo)}
	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}
	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{vfsgen۰CompressedFileInfo: f, gr: gr}, nil
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{vfsgen۰FileInfo: f, Reader: bytes.NewReader(f.content)}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{vfsgen۰DirInfo: f}, nil
	default:
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

type vfsgen۰CompressedFileInfo struct {
	name			string
	modTime			time.Time
	compressedContent	[]byte
	uncompressedSize	int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f, nil
}
func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.compressedContent
}
func (f *vfsgen۰CompressedFileInfo) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.name
}
func (f *vfsgen۰CompressedFileInfo) Size() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.uncompressedSize
}
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0444
}
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.modTime
}
func (f *vfsgen۰CompressedFileInfo) IsDir() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (f *vfsgen۰CompressedFileInfo) Sys() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}

type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr	*gzip.Reader
	grPos	int64
	seekPos	int64
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if f.grPos > f.seekPos {
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.gr.Close()
}

type vfsgen۰FileInfo struct {
	name	string
	modTime	time.Time
	content	[]byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f, nil
}
func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (f *vfsgen۰FileInfo) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.name
}
func (f *vfsgen۰FileInfo) Size() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return int64(len(f.content))
}
func (f *vfsgen۰FileInfo) Mode() os.FileMode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0444
}
func (f *vfsgen۰FileInfo) ModTime() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.modTime
}
func (f *vfsgen۰FileInfo) IsDir() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (f *vfsgen۰FileInfo) Sys() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}

type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}

type vfsgen۰DirInfo struct {
	name	string
	modTime	time.Time
	entries	[]os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return d, nil
}
func (d *vfsgen۰DirInfo) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return d.name
}
func (d *vfsgen۰DirInfo) Size() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0
}
func (d *vfsgen۰DirInfo) Mode() os.FileMode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0755 | os.ModeDir
}
func (d *vfsgen۰DirInfo) ModTime() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return d.modTime
}
func (d *vfsgen۰DirInfo) IsDir() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (d *vfsgen۰DirInfo) Sys() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}

type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos	int
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}
func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}