package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_show_tcp_keepalives_count(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	v1 = int32(0)
	v2 = m.G0
	v3 = int32(16)
	v4 = v2 - v3
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_show_tcp_keepalives_count[0]))
	v10 = m.G0
	v12 = v10 - v3
	m.G0 = v12
	if v7 == v1 {
		v27 = v1
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)))
		if v16 == int32(1) {
			v27 = v1
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+408))
			if v19 != 0 {
				v27 = v19
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+392))
				if v20 != 0 {
					v27 = v20
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(4)
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(392))))
					v27 = v26
				}
			}
		}
	}
	m.G0 = v12 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = v27
	v35 = F_pg_snprintf(m, int32(_a_F_show_tcp_keepalives_count_0), int32(16), int32(_a_F_show_tcp_keepalives_count_1), v4)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		return int32(0)
	} else {
		m.G0 = v4 + int32(16)
		return int32(_a_F_show_tcp_keepalives_count_0)
	}
}
