package ole

import (
	"syscall"
)

var (
	modcombase     = syscall.NewLazyDLL("combase.dll")
	modkernel32, _ = syscall.LoadDLL("kernel32.dll")
	modole32, _    = syscall.LoadDLL("ole32.dll")
	modoleaut32, _ = syscall.LoadDLL("oleaut32.dll")
	modmsvcrt, _   = syscall.LoadDLL("msvcrt.dll")
	moduser32, _   = syscall.LoadDLL("user32.dll")

	IID_NULL                      = &GUID{0x00000000, 0x0000, 0x0000, [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}
	IID_IUnknown                  = &GUID{0x00000000, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	IID_IDispatch                 = &GUID{0x00020400, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	IID_IConnectionPointContainer = &GUID{0xB196B284, 0xBAB4, 0x101A, [8]byte{0xB6, 0x9C, 0x00, 0xAA, 0x00, 0x34, 0x1D, 0x07}}
	IID_IConnectionPoint          = &GUID{0xB196B286, 0xBAB4, 0x101A, [8]byte{0xB6, 0x9C, 0x00, 0xAA, 0x00, 0x34, 0x1D, 0x07}}
	IID_IInspectable              = &GUID{0xaf86e2e0, 0xb12d, 0x4c6a, [8]byte{0x9c, 0x5a, 0xd7, 0xaa, 0x65, 0x10, 0x1e, 0x90}}
	IID_IProvideClassInfo         = &GUID{0xb196b283, 0xbab4, 0x101a, [8]byte{0xB6, 0x9C, 0x00, 0xAA, 0x00, 0x34, 0x1D, 0x07}}
)

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}
