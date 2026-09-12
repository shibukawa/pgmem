package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResourceOwnerForget(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v141 int64
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v15 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return
L2:
	;
	v161 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, uint32(v163+v94<<(uint(int32(3))%32))+4)) = v161
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v169 - int32(1)
	goto L1
L3:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	if v18 != 0 {
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
	v150 = m.ExcPending
	if v150 != 0 {
		goto L27
	} else {
		goto L31
	}
L6:
	;
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v18<<(uint(int32(3))%32)+v20-int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v141
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v145 = v143 - int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v145)
	goto L1
L7:
	;
	v20 = l0 + int32(24)
	v24 = v18
	goto L10
L8:
	;
	goto L9
L9:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v52 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v32 = v24 - int32(1)
	v35 = v20 + v32<<(uint(int32(3))%32)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if l1 == v36 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v38 == l2 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if base.Ui32(int32(1)) < base.Ui32(v24) {
		v24 = v32
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
	v120 = m.ExcPending
	if v120 != 0 {
		goto L27
	} else {
		goto L28
	}
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v55 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v60 = int32(16)
	v64 = (int32(base.Ui32(l1)>>(uint(v60)%32)) ^ l1) * int32(-2048144789)
	v69 = (int32(base.Ui32(v64)>>(uint(int32(13))%32)) ^ v64) * int32(-1028477387)
	v72 = int32(base.Ui32(v69)>>(uint(v60)%32)) ^ v69
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v87 = l2 + v72<<(uint(int32(6))%32) + int32(base.Ui32(v72)>>(uint(int32(2))%32)) - int32(1640531527) ^ v72
	v92 = int32(0)
	goto L20
L20:
	;
	v94 = v87 & (v55 - int32(1))
	v97 = v82 + v94<<(uint(int32(3))%32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if l1 == v98 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L17
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if v100 == l2 {
		goto L2
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v102 = int32(1)
	v105 = v92 + v102
	if v105 != v55 {
		v87 = v94 + v102
		v92 = v105
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
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v121
	F_errmsg_internal(m, int32(171882), v13+int32(16))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(472174), int32(622), int32(100457))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
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
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v151
	F_errmsg_internal(m, int32(423786), v13)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(472174), int32(572), int32(100457))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
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
func F_ResourceOwnerRelease(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v6 int32
	_ = v6
	F_ResourceOwnerReleaseInternal(m, l0, l1, l2, l3)
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_ResourceOwnerRemember(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
			F_errmsg_internal(m, int32(288664), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errfinish(m, int32(472174), int32(542), int32(217410))
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
		v24 = l0 + v6<<(uint(int32(3))%32)
		*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = l1
		v28 = v6 + int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v28)
		return
	}
}
