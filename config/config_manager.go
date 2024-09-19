package config

import (
	"github.com/pelletier/go-toml"
	"gtihub.com/Iceinu-Project/iceinu/logger"
	"os"
	"path/filepath"
	"reflect"
	"sync"
)

// Iceinu的配置文件处理器，除了iceinu自带的配置文件之外也负责处理适配器和插件的配置文件设置
// 这配置文件管理器设计一大半是ChatGPT帮我优化的，o1模型真的是太斯巴拉西了
// 总之是实现了动态的配置文件注册加载解析修正等等功能

// ConfManager 管理配置文件的结构体
type ConfManager struct {
	configs       map[string]interface{}
	defaults      map[string]interface{}
	mutex         sync.RWMutex
	configChanged bool // 标志位，指示配置文件是否有生成或修改
}

// manager 是全局的配置管理器实例
var manager = NewManager()

// NewManager 创建一个新的配置管理器
func NewManager() *ConfManager {
	logger.Debugf("正在初始化配置文件管理器...")
	return &ConfManager{
		configs:  make(map[string]interface{}),
		defaults: make(map[string]interface{}),
	}
}

// InitConfig 注册一个配置文件和对应的结构体
func (cm *ConfManager) InitConfig(filename string, configStruct interface{}) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	cm.configs[filename] = configStruct

	// 保存默认配置的深拷贝
	defaultConfig := deepCopy(configStruct)
	cm.defaults[filename] = defaultConfig

	logger.Debugf("已注册配置文件：%s", filename)
}

// LoadConfigs 加载并解析所有已注册的配置文件，如果不存在则自动生成
// 如果已有的配置文件缺少字段，则补全并写回，并在覆盖前备份原始文件
func (cm *ConfManager) LoadConfigs() error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		logger.Errorf("获取工作目录失败：%v", err)
		return err
	}

	// 遍历已注册的配置文件
	for filename, configStruct := range cm.configs {
		fullPath := filepath.Join(dir, filename)

		// 检查文件是否存在
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			// 如果文件不存在，生成默认配置文件
			data, err := toml.Marshal(configStruct)
			if err != nil {
				logger.Errorf("序列化默认配置失败：%v", err)
				return err
			}

			err = os.WriteFile(fullPath, data, 0644)
			if err != nil {
				logger.Errorf("写入配置文件失败：%v", err)
				return err
			}

			logger.Infof("已生成配置文件：%s", fullPath)
			cm.configChanged = true // 设置标志位
		} else {
			// 如果文件存在，读取并解析配置文件
			data, err := os.ReadFile(fullPath)
			if err != nil {
				logger.Errorf("读取配置文件失败：%v", err)
				return err
			}

			// 创建一个新的配置实例，用于加载文件内容
			newConfig := deepCopy(cm.defaults[filename])

			err = toml.Unmarshal(data, newConfig)
			if err != nil {
				logger.Errorf("解析配置文件失败：%v", err)
				return err
			}

			// 比较新配置和默认配置，补全缺失的字段
			changed := mergeConfig(newConfig, cm.defaults[filename])

			// 将补全后的配置赋值回 configs
			cm.configs[filename] = newConfig

			if changed {
				// 在覆盖前备份原始配置文件
				backupPath := fullPath + ".backup.toml"
				err = backupFile(fullPath, backupPath)
				if err != nil {
					logger.Errorf("备份配置文件失败：%v", err)
					return err
				}

				// 将完整的配置（包含默认值和新解析的值）写回配置文件
				data, err = toml.Marshal(newConfig)
				if err != nil {
					logger.Errorf("序列化配置文件失败：%v", err)
					return err
				}

				err = os.WriteFile(fullPath, data, 0644)
				if err != nil {
					logger.Errorf("写入配置文件失败：%v", err)
					return err
				}

				logger.Infof("配置文件已更新并备份：%s", fullPath)
				cm.configChanged = true // 设置标志位
			} else {
				logger.Debugf("配置文件已加载，无需更新：%s", fullPath)
			}
		}
	}

	// 如果配置文件有变动，提示用户并退出程序
	if cm.configChanged {
		logger.Warnf("配置文件已生成或更新，请检查配置文件后重新启动程序。")
		os.Exit(0)
	}

	return nil
}

// Get 获取指定名称的配置
func (cm *ConfManager) Get(filename string) interface{} {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	return cm.configs[filename]
}

// GetManager 获取配置管理器实例
func GetManager() *ConfManager {
	return manager
}

// deepCopy 深拷贝一个配置结构体
func deepCopy(src interface{}) interface{} {
	// 将 src 序列化为 TOML
	data, err := toml.Marshal(src)
	if err != nil {
		logger.Errorf("深拷贝失败：%v", err)
		return nil
	}

	// 创建一个新的与 src 类型相同的实例
	dst := reflect.New(reflect.TypeOf(src).Elem()).Interface()

	// 将 TOML 反序列化到 dst
	err = toml.Unmarshal(data, dst)
	if err != nil {
		logger.Errorf("深拷贝失败：%v", err)
		return nil
	}

	return dst
}

// mergeConfig 递归地将 src 中的非零值合并到 dst 中，返回是否有修改
func mergeConfig(dst, src interface{}) bool {
	dstVal := reflect.ValueOf(dst).Elem()
	srcVal := reflect.ValueOf(src).Elem()

	changed := false

	for i := 0; i < dstVal.NumField(); i++ {
		dstField := dstVal.Field(i)
		srcField := srcVal.Field(i)

		switch dstField.Kind() {
		case reflect.Struct:
			if mergeConfig(dstField.Addr().Interface(), srcField.Addr().Interface()) {
				changed = true
			}
		case reflect.Slice, reflect.Map:
			if dstField.IsNil() && !srcField.IsNil() {
				dstField.Set(srcField)
				changed = true
			}
		default:
			// 如果 dstField 是零值，则使用 srcField 的值
			if isZeroValue(dstField) && !isZeroValue(srcField) {
				dstField.Set(srcField)
				changed = true
			}
		}
	}

	return changed
}

// isZeroValue 判断一个值是否是零值
func isZeroValue(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

// backupFile 备份原始配置文件
func backupFile(originalPath, backupPath string) error {
	// 读取原始文件内容
	data, err := os.ReadFile(originalPath)
	if err != nil {
		return err
	}

	// 写入备份文件
	err = os.WriteFile(backupPath, data, 0644)
	if err != nil {
		return err
	}

	logger.Infof("已备份配置文件：%s", backupPath)
	return nil
}
