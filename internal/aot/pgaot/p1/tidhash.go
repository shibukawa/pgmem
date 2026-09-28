package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tidhash_estimate_space(m *base.Module, l0 float64) int32 {
	var v4 int32
	_ = v4
	var v6 float64
	_ = v6
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v33 int32
	_ = v33
	v4 = int32(-1)
	v6 = base.F64_div(l0, float64(0.9))
	if base.F64_ge(v6, float64(4.294967296e+09)) != 0 {
		v33 = v4
	} else {
		v9 = int64(2)
		v10 = base.I64_trunc_sat_f64_u(v6)
		if base.Ui64(v10) <= base.Ui64(v9) {
			v13 = v9
		} else {
			v13 = v10
		}
		v14 = int64(1)
		if v13&(v13-v14) == int64(0) {
			v24 = v13
		} else {
			v24 = v14 << (uint(int64(64)-base.I64_clz(v13)) % 64)
		}
		v26 = v24 << (uint(int64(3)) % 64)
		if base.Ui64(int64(2147483646)) < base.Ui64(v26) {
			v33 = v4
		} else {
			v33 = base.I32_wrap_i64(v26) + int32(32)
		}
	}
	return v33
}
func F_tidhash_start_iterate(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v6 = int32(-1)
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v7 == int64(0) {
		v29 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v32)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(0)
	goto L3
L3:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v12<<(uint(int32(3))%32))+6)))
	if v20 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v29 = v6
	goto L1
L5:
	;
	v29 = v12
	goto L1
L6:
	;
	goto L7
L7:
	;
	v24 = v12 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v24)) < base.Ui64(v7) {
		v12 = v24
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L4
}
