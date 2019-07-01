package utils

import (
	"gopkg.in/ini.v1"
	"jmanager/common/logs"
	"errors"
)

type iniParser struct {
	confReaders map[string]*ini.File
}

var IniParser = iniParser{}

func (this *iniParser) getReader(fileName string) (*ini.File, error) {
	reader, ok := this.confReaders[fileName]
	if ok {
		return reader, nil
	}

	nowReader, err := ini.Load(fileName)
	if err != nil {
		logs.Log.LogError(err)
		return nil, err
	}

	this.confReaders[fileName] = nowReader
	return nowReader, nil
}

func (this *iniParser) GetString(fileName string, section string, key string) (string, error) {
	reader, err := this.getReader(fileName)
	if err != nil {
		return "", err
	}

	s := reader.Section(section)
	if s == nil {
		return "", errors.New("块配置不存在")
	}

	return s.Key(key).String(), nil
}

func (this *iniParser) GetInt64(fileName string, section string, key string) (int64, error) {
	reader, err := this.getReader(fileName)
	if err != nil {
		return 0, err
	}

	s := reader.Section(section)
	if s == nil {
		return 0, errors.New("块配置不存在")
	}

	return s.Key(key).Int64()
}

func (this *iniParser) GetFloat64(fileName string, section string, key string) (float64, error) {
	reader, err := this.getReader(fileName)
	if err != nil {
		return 0, err
	}

	s := reader.Section(section)
	if s == nil {
		return 0, errors.New("块配置不存在")
	}

	return s.Key(key).Float64()
}
