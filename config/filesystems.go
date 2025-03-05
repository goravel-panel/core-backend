package config

import (
	cosFacades "github.com/goravel/cos/facades"
	ossFacades "github.com/goravel/oss/facades"

	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/path"
)

func init() {
	config := facades.Config()
	config.Add("filesystems", map[string]any{
		// Default Filesystem Disk
		//
		// Here you may specify the default filesystem disk that should be used
		// by the framework. The "local" disk, as well as a variety of cloud
		// based disks are available to your application. Just store away!
		"default": config.Env("FILESYSTEM_DISK", "local"),

		// Filesystem Disks
		//
		// Here you may configure as many filesystem "disks" as you wish, and you
		// may even configure multiple disks of the same driver. Defaults have
		// been set up for each driver as an example of the required values.
		//
		// Supported Drivers: "local", "s3", "oss", "cos", "minio", "custom"
		"disks": map[string]any{
			"local": map[string]any{
				"driver": "local",
				"root":   path.Storage("app"),
			},
			"public": map[string]any{
				"driver": "local",
				"root":   path.Public("static"),
				"url":    config.Env("APP_URL", "").(string) + "/static",
			},
			"oss": map[string]any{
				"driver":   "custom",
				"key":      config.Env("ALIYUN_ACCESS_KEY_ID"),
				"secret":   config.Env("ALIYUN_ACCESS_KEY_SECRET"),
				"bucket":   config.Env("ALIYUN_BUCKET"),
				"url":      config.Env("ALIYUN_URL"),
				"endpoint": config.Env("ALIYUN_ENDPOINT"),
				"via": func() (filesystem.Driver, error) {
					return ossFacades.Oss("oss"), nil
				},
			},
			"cos": map[string]any{
				"driver": "custom",
				"key":    config.Env("TENCENT_ACCESS_KEY_ID"),
				"secret": config.Env("TENCENT_ACCESS_KEY_SECRET"),
				"url":    config.Env("TENCENT_URL"),
				"via": func() (filesystem.Driver, error) {
					return cosFacades.Cos("cos"), nil
				},
			},
		},
	})
}
