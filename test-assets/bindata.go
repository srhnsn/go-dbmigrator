package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _test_migrations_2015_01_01_120000_create_table_role_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x8e\x31\xcb\xc2\x30\x10\x86\xf7\xfc\x8a\x1b\x13\xf8\x86\x4f\xc1\x49\x1c\xd2\xf6\xa8\x87\xed\x55\xe3\x45\xec\x64\x04\x3b\x88\x9a\x82\x9b\xff\xde\x68\x2b\x0e\xde\x72\x70\xef\xfb\x1c\x4f\xee\xd0\x0a\x82\xd8\xac\x42\x08\xf7\xfe\xda\x05\xd0\x0a\xd2\x84\xf3\x29\x00\xb1\xe8\xc9\xbf\x01\xcf\x5b\x2a\x19\x0b\xe0\x46\x80\x7d\x55\x81\xf5\xd2\x1c\x88\x13\x5f\x23\xcb\xdf\x80\xc4\xe3\x2d\xf1\x3b\xeb\xf2\xa5\x75\x7a\x96\xc0\x4f\x7f\x2c\x5c\xba\xc7\x37\x9f\xfe\xe4\x6b\x47\xb5\x75\x2d\xac\xb0\x05\xfd\x12\x30\xc3\xdd\x33\x6d\x3c\x26\x9b\x02\xf7\xe3\x13\xfd\x5e\x46\x19\x85\x5c\x12\xe3\x82\x62\xec\x8b\x6c\xae\x9e\x01\x00\x00\xff\xff\xd3\x42\x64\x0e\xd2\x00\x00\x00")

func test_migrations_2015_01_01_120000_create_table_role_sql_bytes() ([]byte, error) {
	return bindata_read(
		_test_migrations_2015_01_01_120000_create_table_role_sql,
		"test-migrations/2015-01-01 120000 create table role.sql",
	)
}

func test_migrations_2015_01_01_120000_create_table_role_sql() (*asset, error) {
	bytes, err := test_migrations_2015_01_01_120000_create_table_role_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test-migrations/2015-01-01 120000 create table role.sql", size: 210, mode: os.FileMode(438), modTime: time.Unix(1426941689, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _test_migrations_2015_01_01_120100_create_table_user_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\xd0\xcd\x4a\xc4\x30\x10\xc0\xf1\x7b\x9e\x62\x8e\x09\x78\xd8\x45\x3c\x89\x87\xec\x76\xa8\xc1\x76\xaa\x31\x11\x7b\x6a\x02\xad\x50\xe8\x87\x34\x8a\xaf\x6f\x4a\x95\x56\xea\x61\x73\x0a\xe1\xf7\x27\xcc\x9c\x35\x4a\x83\x60\xe4\x29\x43\x70\x9f\xa1\x99\x1c\x70\x06\xf1\xb8\xb6\x76\xa0\xc8\xf0\xe3\x41\x80\xa5\x67\x95\x12\x26\x40\x85\x01\xb2\x59\x06\xd2\x9a\xa2\x52\x14\xfb\x1c\xc9\x5c\x2d\xc9\xdc\x0f\xbe\x6f\x1c\xbc\x48\x7d\xbe\x97\x9a\x5f\xc7\xf8\xb7\xf9\x41\xef\x3e\x84\xaf\x71\xaa\x57\x74\x3c\xec\x55\xd3\xfb\xb6\xdb\x90\x9b\x3d\x79\x6b\xa7\xf0\x51\xfd\xfd\xef\x1f\xd6\xf9\x0b\xd4\xa3\x56\xb9\xd4\x25\x3c\x60\x09\x7c\x1e\x5d\x2c\xef\x96\xd4\x93\xc5\xb8\x87\x04\x5f\xb7\xf3\xf1\xf5\x2e\x98\x60\x48\xa9\x22\xbc\x53\xc3\x30\x26\xa7\x5b\xf6\x1d\x00\x00\xff\xff\x1e\xbc\x85\x0e\x56\x01\x00\x00")

func test_migrations_2015_01_01_120100_create_table_user_sql_bytes() ([]byte, error) {
	return bindata_read(
		_test_migrations_2015_01_01_120100_create_table_user_sql,
		"test-migrations/2015-01-01 120100 create table user.sql",
	)
}

func test_migrations_2015_01_01_120100_create_table_user_sql() (*asset, error) {
	bytes, err := test_migrations_2015_01_01_120100_create_table_user_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test-migrations/2015-01-01 120100 create table user.sql", size: 342, mode: os.FileMode(438), modTime: time.Unix(1426941715, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _test_migrations_2015_01_01_120200_create_table_user_to_role_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x90\x31\x4f\xc4\x20\x18\x86\x77\x7e\xc5\x37\x5e\x13\x07\x9d\x8d\x03\x07\x6f\x2f\xe4\xc8\x87\x01\x9a\xe8\x54\x06\x6f\x30\x31\xd7\xa4\xea\xff\x17\x4a\xad\xed\xa0\x96\x8d\x97\x17\x9e\xe7\x43\x79\xc8\x08\x8a\xf2\x68\x41\xe9\xf3\xfd\x32\xf6\x1f\x43\x3f\x0e\x6f\x97\x44\x07\x41\x79\xd5\xf0\xf5\x25\x91\xe1\x78\xb8\xbb\x6d\xa8\xe3\x60\x4e\x0c\x4d\xec\x22\x71\x67\xed\x4d\x2d\x96\x5b\x7b\x8a\x86\x35\x9e\x28\xb5\xe7\x7e\xcd\x9b\x36\x19\xba\xf0\x9a\x3f\xdb\xb3\xe2\x02\x9d\xdb\xca\x71\x88\x5e\x66\x83\xdf\x00\xad\xf3\xc8\x56\x74\xc6\xf3\x1a\x46\x1e\x2d\x3c\x58\x21\xd4\x91\xcb\xe3\xd3\x81\x63\xea\x1e\x75\xf9\x25\x25\x83\x92\x1a\x25\xd1\xb0\xf8\x49\xfe\x67\x57\xdd\x2d\x7b\x51\xdf\xb0\xbf\x07\xdb\xcb\x16\x8d\x00\x9f\x0c\xe3\xc1\x5c\xaf\x83\x3e\xde\x8b\xaf\x00\x00\x00\xff\xff\x3f\xe0\x59\x53\xd4\x01\x00\x00")

func test_migrations_2015_01_01_120200_create_table_user_to_role_sql_bytes() ([]byte, error) {
	return bindata_read(
		_test_migrations_2015_01_01_120200_create_table_user_to_role_sql,
		"test-migrations/2015-01-01 120200 create table user_to_role.sql",
	)
}

func test_migrations_2015_01_01_120200_create_table_user_to_role_sql() (*asset, error) {
	bytes, err := test_migrations_2015_01_01_120200_create_table_user_to_role_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test-migrations/2015-01-01 120200 create table user_to_role.sql", size: 468, mode: os.FileMode(438), modTime: time.Unix(1426941725, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _test_migrations_2015_01_01_130000_insert_into_role_default_data_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xf2\xf4\x0b\x76\x0d\x0a\x51\xf0\xf4\x0b\xf1\xe7\x52\x00\x82\xa2\xfc\x9c\x54\x05\x0d\x85\xbc\xc4\xdc\x54\x05\x1d\x85\x84\xec\xd4\xca\x04\x05\x4d\xae\x30\x47\x9f\x50\xd7\x60\xb0\x02\x0d\x05\x25\xc7\x94\xdc\xcc\xbc\xcc\xe2\x92\xa2\xc4\x92\xfc\x22\x25\xa0\x32\xa5\x44\x90\x88\x92\x82\xa6\x82\x0e\x4c\x8d\x6f\x6a\x5e\x29\x58\x2a\x17\xcc\x40\x92\x71\x4b\x4d\x4d\x49\x4a\x4c\xce\x56\x02\x4a\xa6\xc1\xd8\x0a\x9a\xd6\x5c\x80\x00\x00\x00\xff\xff\x69\x2b\xc7\x15\x8b\x00\x00\x00")

func test_migrations_2015_01_01_130000_insert_into_role_default_data_sql_bytes() ([]byte, error) {
	return bindata_read(
		_test_migrations_2015_01_01_130000_insert_into_role_default_data_sql,
		"test-migrations/2015-01-01 130000 insert into role default data.sql",
	)
}

func test_migrations_2015_01_01_130000_insert_into_role_default_data_sql() (*asset, error) {
	bytes, err := test_migrations_2015_01_01_130000_insert_into_role_default_data_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test-migrations/2015-01-01 130000 insert into role default data.sql", size: 139, mode: os.FileMode(438), modTime: time.Unix(1426941604, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _test_migrations_2015_01_01_130100_insert_into_user_admin_user_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x34\x8d\x51\x4f\xc2\x30\x10\xc7\xdf\xf9\x14\x7d\x58\x32\x48\x10\xe2\x18\xcc\xc4\x17\x2b\x36\x01\x9d\x1b\x3a\x34\xfa\xe4\x4e\x56\x68\xb3\xdd\xda\xac\x35\x96\x6f\x2f\x45\x76\x2f\xf7\xfb\xff\x72\xb9\xff\x3a\x2b\xd8\xeb\x96\xac\xb3\x6d\x4e\xca\x1f\xc3\xbb\x92\x0c\x4b\x59\x95\xe3\xff\xd4\x02\x72\xcf\x1a\x8c\xf9\x55\xdd\xd9\x73\x04\xd9\x78\xd8\xcb\xce\xd8\xaf\xfe\xa4\x81\x3e\x8c\xc8\x3b\x4d\xdf\x58\x31\x20\xa7\x19\x5e\x8f\x49\x08\x15\xca\x36\x3c\x41\xa0\xbf\xeb\x6a\x1f\x5d\x19\x01\xd1\x7c\x11\x24\x37\x8b\x79\x40\xe5\x92\xb1\xe7\x34\x9e\x39\xfb\x00\x34\x12\x76\xf5\x99\xab\x97\x60\x9a\xb2\xf6\x5e\x47\xb3\xe3\x2a\xa9\xd1\x31\x79\x9c\xe6\x9b\x83\x60\x4f\x1f\xb1\xe0\x71\x76\x28\x68\xb5\xc1\x7a\xe2\xe8\x32\xd9\xf9\xcf\xdc\x01\xea\x86\xdf\x5d\xf6\x64\xa7\xd0\xeb\x47\x25\xce\xc5\x05\x4a\x2b\xc2\xd1\xed\xe0\x2f\x00\x00\xff\xff\x35\x12\xc6\x8b\xf1\x00\x00\x00")

func test_migrations_2015_01_01_130100_insert_into_user_admin_user_sql_bytes() ([]byte, error) {
	return bindata_read(
		_test_migrations_2015_01_01_130100_insert_into_user_admin_user_sql,
		"test-migrations/2015-01-01 130100 insert into user admin user.sql",
	)
}

func test_migrations_2015_01_01_130100_insert_into_user_admin_user_sql() (*asset, error) {
	bytes, err := test_migrations_2015_01_01_130100_insert_into_user_admin_user_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test-migrations/2015-01-01 130100 insert into user admin user.sql", size: 241, mode: os.FileMode(438), modTime: time.Unix(1426941651, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _test_migrations_2015_01_01_130200_insert_into_user_to_role_admin_user_roles_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xf2\xf4\x0b\x76\x0d\x0a\x51\xf0\xf4\x0b\xf1\x57\x48\x28\x2d\x4e\x2d\x8a\x2f\xc9\x8f\x2f\xca\xcf\x49\x4d\x50\xd0\x80\xf0\x33\x53\x12\x74\x14\x12\x40\x42\x20\xa6\xa6\x42\x98\xa3\x4f\xa8\x6b\x30\x97\x02\x10\x68\x18\xea\x28\x18\x6a\x5a\x73\x01\x02\x00\x00\xff\xff\x66\xbf\x03\xf7\x45\x00\x00\x00")

func test_migrations_2015_01_01_130200_insert_into_user_to_role_admin_user_roles_sql_bytes() ([]byte, error) {
	return bindata_read(
		_test_migrations_2015_01_01_130200_insert_into_user_to_role_admin_user_roles_sql,
		"test-migrations/2015-01-01 130200 insert into user_to_role admin user roles.sql",
	)
}

func test_migrations_2015_01_01_130200_insert_into_user_to_role_admin_user_roles_sql() (*asset, error) {
	bytes, err := test_migrations_2015_01_01_130200_insert_into_user_to_role_admin_user_roles_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test-migrations/2015-01-01 130200 insert into user_to_role admin user roles.sql", size: 69, mode: os.FileMode(438), modTime: time.Unix(1403475657, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"test-migrations/2015-01-01 120000 create table role.sql": test_migrations_2015_01_01_120000_create_table_role_sql,
	"test-migrations/2015-01-01 120100 create table user.sql": test_migrations_2015_01_01_120100_create_table_user_sql,
	"test-migrations/2015-01-01 120200 create table user_to_role.sql": test_migrations_2015_01_01_120200_create_table_user_to_role_sql,
	"test-migrations/2015-01-01 130000 insert into role default data.sql": test_migrations_2015_01_01_130000_insert_into_role_default_data_sql,
	"test-migrations/2015-01-01 130100 insert into user admin user.sql": test_migrations_2015_01_01_130100_insert_into_user_admin_user_sql,
	"test-migrations/2015-01-01 130200 insert into user_to_role admin user roles.sql": test_migrations_2015_01_01_130200_insert_into_user_to_role_admin_user_roles_sql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"test-migrations": &_bintree_t{nil, map[string]*_bintree_t{
		"2015-01-01 120000 create table role.sql": &_bintree_t{test_migrations_2015_01_01_120000_create_table_role_sql, map[string]*_bintree_t{
		}},
		"2015-01-01 120100 create table user.sql": &_bintree_t{test_migrations_2015_01_01_120100_create_table_user_sql, map[string]*_bintree_t{
		}},
		"2015-01-01 120200 create table user_to_role.sql": &_bintree_t{test_migrations_2015_01_01_120200_create_table_user_to_role_sql, map[string]*_bintree_t{
		}},
		"2015-01-01 130000 insert into role default data.sql": &_bintree_t{test_migrations_2015_01_01_130000_insert_into_role_default_data_sql, map[string]*_bintree_t{
		}},
		"2015-01-01 130100 insert into user admin user.sql": &_bintree_t{test_migrations_2015_01_01_130100_insert_into_user_admin_user_sql, map[string]*_bintree_t{
		}},
		"2015-01-01 130200 insert into user_to_role admin user roles.sql": &_bintree_t{test_migrations_2015_01_01_130200_insert_into_user_to_role_admin_user_roles_sql, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

