package setting

import (
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/goravel/framework/support/path"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"path/filepath"
	"sync"
)

type commonConf struct {
	HttpPort  string
	RPCPort   string
	Cluster   bool
	CryptoKey string
}

var CommonSetting = &commonConf{}

type etcdConf struct {
	Endpoints []string
}

var EtcdSetting = &etcdConf{}

type global struct {
	LocalHost      string //本机内网IP
	ServerList     map[string]string
	ServerListLock sync.RWMutex
}

var GlobalSetting = &global{}

var cfg *ini.File

func getConfigPath() string {
	var defaultConfigFile string
	path := path.Base()
	defaultConfigFile = filepath.Join(path, "packages", "goravel-socket", "config", "app.ini")
	return defaultConfigFile
}

func Setup() {

	defaultConfigFile := getConfigPath()

	// 命令行参数
	configFile := flag.String("c", defaultConfigFile, "Path to the configuration file")
	flag.Parse()
	var err error
	cfg, err = ini.Load(*configFile)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"错误": "ini.Load",
		}).Infof(fmt.Sprintf("加载配置文件失败：%s", err))
		//log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
		//加载默认数据
		Default()
	}
	// 打印所有配置项
	log.Println("[CONFIG] Loaded configuration:")
	for _, section := range cfg.Sections() {
		if len(section.Keys()) == 0 {
			log.Printf("  [%s] (empty section)", section.Name())
			continue
		}
		log.Printf("  [%s]", section.Name())
		for _, key := range section.Keys() {
			log.Printf("    %s = %v", key.Name(), key.Value())
			if key.Name() == "Cluster" {
				CommonSetting.Cluster = key.MustBool()
			}
			if key.Name() == "HttpPort" {
				CommonSetting.HttpPort = key.MustString("")
			}
			if key.Name() == "RPCPort" {
				CommonSetting.RPCPort = key.MustString("")
			}
		}
	}

	// 重点检查CommonSetting是否加载成功
	if cfg.Section("common").HasKey("CryptoKey") {
		log.Println("[CONFIG] CryptoKey found in configuration")
		CommonSetting.CryptoKey = cfg.Section("common").Key("CryptoKey").String()
	} else {
		log.Println("[CONFIG] WARNING: CryptoKey NOT found in configuration")
	}
	mapTo("common", cfg)
	mapTo("etcd", EtcdSetting)

	GlobalSetting = &global{
		LocalHost:  getIntranetIp(),
		ServerList: make(map[string]string),
	}
}

func Default() {
	CommonSetting = &commonConf{
		HttpPort:  "6000",
		RPCPort:   "7000",
		Cluster:   false,
		CryptoKey: "Adba723b7fe06819",
	}

	GlobalSetting = &global{
		LocalHost:  getIntranetIp(),
		ServerList: make(map[string]string),
	}
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

// 获取本机内网IP
func getIntranetIp() string {
	addrs, _ := net.InterfaceAddrs()

	for _, addr := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
