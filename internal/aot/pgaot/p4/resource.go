package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResourceOwnerForget(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	var v73 int64
	_ = v73
	var v76 int64
	_ = v76
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	v2 = l1
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v16 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(32)
	return
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v100))) = int64(0)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
	*(*int32)(unsafe.Add(mBase, uint32(v170+v90<<(uint(int32(4))%32))+8)) = int32(0)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v176 - int32(1)
	goto L1
L3:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L27
	} else {
		goto L31
	}
L6:
	;
	v145 = v21 + v19<<(uint(int32(4))%32) - int32(16)
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v145)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+8)) = v146
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v145)))
	*(*int64)(unsafe.Add(mBase, uint32(v37))) = v148
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v152 = v150 - int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v152)
	goto L1
L7:
	;
	v21 = l0 + int32(24)
	v25 = v19
	goto L10
L8:
	;
	goto L9
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v55 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v34 = v25 - int32(1)
	v37 = v21 + v34<<(uint(int32(4))%32)
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
	if v2 == v38 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	if v40 == l2 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if base.Ui32(int32(1)) < base.Ui32(v25) {
		v25 = v34
		goto L10
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	goto L11
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L27
	} else {
		goto L28
	}
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+540))
	if v58 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v62 = v58 - int32(1)
	v64 = int64(33)
	v68 = (int64(base.Ui64(v2)>>(uint(v64)%64)) ^ v2) * int64(-49064778989728563)
	v73 = (int64(base.Ui64(v68)>>(uint(v64)%64)) ^ v68) * int64(-4265267296055464877)
	v76 = int64(base.Ui64(v73)>>(uint(v64)%64)) ^ v73
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
	v90 = v62 & base.I32_wrap_i64(base.I64_extend_i32_u(l2)+int64(base.Ui64(v76)>>(uint(int64(7))%64))+int64(367372515)^v76)
	v95 = int32(0)
	goto L20
L20:
	;
	v100 = v85 + v90<<(uint(int32(4))%32)
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v100)))
	if v2 == v101 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L17
L22:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	if v103 == l2 {
		goto L2
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v105 = int32(1)
	v109 = v95 + v105
	if v109 != v58 {
		v90 = (v90 + v105) & v62
		v95 = v109
		goto L20
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	goto L21
L27:
	;
	return
L28:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v127
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+20)) = uint32(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v126
	F_errmsg_internal(m, int32(_a_F_ResourceOwnerForget_0), v14+int32(16))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_ResourceOwnerForget_1), int32(629), int32(_a_F_ResourceOwnerForget_2))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v158
	F_errmsg_internal(m, int32(_a_F_ResourceOwnerForget_3), v14)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_ResourceOwnerForget_1), int32(579), int32(_a_F_ResourceOwnerForget_2))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ResourceOwnerRemember(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	if base.Ui32(int32(32)) <= base.Ui32(v6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_ResourceOwnerRemember_0), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ResourceOwnerRemember_1), int32(549), int32(_a_F_ResourceOwnerRemember_2))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v24 = l0 + v6<<(uint(int32(4))%32)
		*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = l2
		*(*int64)(unsafe.Add(mBase, uint32(v24)+24)) = l1
		v28 = v6 + int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v28)
		return
	}
}
