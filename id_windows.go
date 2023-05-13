//go:build windows

package id

import "golang.org/x/sys/windows/registry"

func id() (string, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Cryptography`, registry.QUERY_VALUE|registry.WOW64_64KEY)
	if err != nil {
		return "", err
	}
	defer key.Close()

	s, _, err := key.GetStringValue("MachineGuid")
	return s, err
}
