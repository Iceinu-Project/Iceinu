package config

import (
	"encoding/json"
	"github.com/Iceinu-Project/iceinu/log"
	"os"

	"github.com/pelletier/go-toml"
)

// 将结构体转换为 map
func structToMap(v interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	return m, err
}

// 读取 TOML 文件到 map
func readTomlFileToMap(filename string) (map[string]interface{}, error) {
	tree, err := toml.LoadFile(filename)
	if err != nil {
		return nil, err
	}
	return tree.ToMap(), nil
}

// 将 map 写入 TOML 文件
func writeMapToTomlFile(m map[string]interface{}, filename string) error {
	tree, err := toml.TreeFromMap(m)
	if err != nil {
		return err
	}
	tomlString := tree.String()
	return os.WriteFile(filename, []byte(tomlString), 0644)
}

// 备份文件
func backupFile(filename string) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	backupFilename := filename + ".backup.toml"
	return os.WriteFile(backupFilename, input, 0644)
}

// 合并两个 map，并检测是否有缺失的键
func mergeMaps(defaultMap, fileMap map[string]interface{}) (map[string]interface{}, bool) {
	mergedMap := make(map[string]interface{})
	missingKeys := false

	// 复制 defaultMap 的所有键值对到 mergedMap
	for key, defVal := range defaultMap {
		mergedMap[key] = defVal
	}

	// 用 fileMap 的值覆盖 mergedMap，并检测缺失的键
	for key, fileVal := range fileMap {
		if defVal, ok := defaultMap[key]; ok {
			defSubMap, defIsMap := defVal.(map[string]interface{})
			fileSubMap, fileIsMap := fileVal.(map[string]interface{})
			if defIsMap && fileIsMap {
				subMergedMap, subMissing := mergeMaps(defSubMap, fileSubMap)
				mergedMap[key] = subMergedMap
				if subMissing {
					missingKeys = true
				}
			} else {
				mergedMap[key] = fileVal
			}
		} else {
			mergedMap[key] = fileVal
		}
	}

	// 检查是否有缺失的键
	for key := range defaultMap {
		if _, ok := fileMap[key]; !ok {
			missingKeys = true
			break
		}
	}

	return mergedMap, missingKeys
}

// ProcessConfig 处理配置文件，需要传入预先配置了默认值的结构体和配置文件名
func ProcessConfig(cfg interface{}, filename string) error {
	defaultMap, err := structToMap(cfg)
	if err != nil {
		return err
	}

	fileMap, err := readTomlFileToMap(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// 文件不存在，生成新的配置文件
			log.Warnf("配置文件 %s 不存在，将生成新的配置文件", filename)
			err = writeMapToTomlFile(defaultMap, filename)
			if err != nil {
				return err
			}
		} else {
			return err
		}
		// 文件已生成，使用默认配置
		return nil
	}

	mergedMap, missingKeys := mergeMaps(defaultMap, fileMap)
	if missingKeys {
		// 备份原始配置文件
		log.Warnf("配置文件 %s 缺失了键，将自动备份原始配置文件并写入新的配置文件", filename)
		err = backupFile(filename)
		if err != nil {
			return err
		}
		// 写入更新后的配置文件
		err = writeMapToTomlFile(mergedMap, filename)
		if err != nil {
			return err
		}
	}

	// 将 TOML 文件解析到配置结构体
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = toml.Unmarshal(data, cfg)
	return err
}
