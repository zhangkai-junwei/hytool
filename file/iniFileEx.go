package file

import "gopkg.in/ini.v1"

type IniCtl struct {
	cfg      *ini.File
	filePath string
}

func (m *IniCtl) LoadIniFile(filePath string) (err error) {
	m.filePath = filePath
	m.cfg, err = ini.Load(filePath)
	return
}

func (m *IniCtl) GetValueAsString(section, key string) string {
	return m.cfg.Section(section).Key(key).String()
}

func (m *IniCtl) GetValueAsInt(section, key string) (int, error) {
	return m.cfg.Section(section).Key(key).Int()
}

func (m *IniCtl) GetValueAsBool(section, key string) (bool, error) {
	return m.cfg.Section(section).Key(key).Bool()
}

func (m *IniCtl) SetValue(section, key, value string) error {
	m.cfg.Section(section).Key(key).SetValue(value)
	return m.cfg.SaveTo(m.filePath)
}
