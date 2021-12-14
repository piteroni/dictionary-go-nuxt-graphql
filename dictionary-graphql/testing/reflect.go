package testing

func Int(v int) *int {
	return &v
}

func String(v string) *string {
	return &v
}

func UInt(v uint) *uint {
	return &v
}

func Bytes(v []byte) *[]byte {
	return &v
}
