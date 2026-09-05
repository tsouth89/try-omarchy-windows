package main

import (
	"syscall"
	"unsafe"
)

// hostTimeZoneKey reads the Windows time zone key name, for example
// "Eastern Standard Time", which the CLDR table maps to an IANA zone.
func hostTimeZoneKey() string {
	path, _ := syscall.UTF16PtrFromString(`SYSTEM\CurrentControlSet\Control\TimeZoneInformation`)
	var key syscall.Handle
	if syscall.RegOpenKeyEx(syscall.HKEY_LOCAL_MACHINE, path, 0, syscall.KEY_READ, &key) != nil {
		return ""
	}
	defer syscall.RegCloseKey(key)
	return registryString(key, "TimeZoneKeyName")
}

// hostKeyboardLayoutID reads the user's default input language as the
// eight-digit keyboard layout identifier at the head of the Preload list.
func hostKeyboardLayoutID() string {
	path, _ := syscall.UTF16PtrFromString(`Keyboard Layout\Preload`)
	var key syscall.Handle
	if syscall.RegOpenKeyEx(syscall.HKEY_CURRENT_USER, path, 0, syscall.KEY_READ, &key) != nil {
		return ""
	}
	defer syscall.RegCloseKey(key)
	return registryString(key, "1")
}

var procGetUserDefaultLocaleName = kernel32.NewProc("GetUserDefaultLocaleName")

// hostLocaleName reads the user's Windows display locale, for example "de-DE".
func hostLocaleName() string {
	buf := make([]uint16, 85)
	n, _, _ := procGetUserDefaultLocaleName.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if n == 0 {
		return ""
	}
	return syscall.UTF16ToString(buf)
}

// hostLocale resolves what the guest should follow, honoring the explicit
// overrides: "" follows Windows, "keep" leaves the guest alone, anything else
// is used as given.
func hostLocale(zoneOverride, keyboardOverride, localeOverride string) (zone, layout, variant, locale string) {
	switch localeOverride {
	case "":
		locale = posixLocaleForWindows(hostLocaleName())
	case "keep":
	default:
		locale = localeOverride
	}
	switch zoneOverride {
	case "":
		zone = ianaZoneForWindows(hostTimeZoneKey())
	case "keep":
	default:
		zone = zoneOverride
	}
	switch keyboardOverride {
	case "":
		layout, variant = xkbForKLID(hostKeyboardLayoutID())
	case "keep":
	default:
		layout, variant = splitKeyboardSpec(keyboardOverride)
	}
	return zone, layout, variant, locale
}
