//go:build linux

package id

func id() (string, error) {
	return fallback(
		read("/var/lib/dbus/machine-id"),
		read("/etc/machine-id"),
		read("/sys/class/dmi/id/product_uuid"),
	)
}
