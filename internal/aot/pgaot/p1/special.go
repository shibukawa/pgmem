package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DecodeSpecial(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	v9 = *(*int32)(unsafe.Add(mBase, _consts[1067]))
	if v9 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1067])) = v140
	v147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v140)+11)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v148
	return v147
L2:
	;
	goto L7
L3:
	;
	goto L4
L4:
	;
	v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	v63 = int32(1617440)
	v65 = int32(1618576)
	goto L20
L5:
	;
	if v46-v47 == int32(0) {
		v140 = v9
		goto L1
	} else {
		goto L19
	}
L7:
	;
	goto L8
L8:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v17 = l0
	v18 = v9
	v19 = int32(10)
	v20 = v16
	goto L13
L10:
	;
	v42 = v9
	v46 = int32(0)
	goto L11
L11:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	goto L5
L12:
	;
	v42 = v37
	v46 = v39
	goto L11
L13:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v20 != v22 {
		v37 = v18
		v39 = v20
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v37 = v31
	v39 = int32(0)
	goto L12
L15:
	;
	if v22 == int32(0) {
		v37 = v18
		v39 = v20
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v27 = v19 - int32(1)
	if v27 == int32(0) {
		v37 = v18
		v39 = v20
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v30 = int32(1)
	v31 = v18 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v32 != 0 {
		v17 = v17 + v30
		v18 = v31
		v19 = v27
		v20 = v32
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	goto L4
L20:
	;
	v72 = v63 + (v65-v63)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v73 = int32(*(*int8)(unsafe.Add(mBase, uint32(v72))))
	v74 = v57 - v73
	if v74 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return int32(31)
L22:
	;
	goto L27
L23:
	;
	v124 = v74
	goto L24
L24:
	;
	v128 = base.B2i32(v124 < int32(0))
	if v124 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L25:
	;
	if v115 == int32(0) {
		v140 = v72
		goto L1
	} else {
		goto L39
	}
L27:
	;
	goto L28
L28:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v83 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v84 = l0
	v85 = v72
	v86 = int32(10)
	v87 = v83
	goto L33
L30:
	;
	v109 = v72
	v113 = int32(0)
	goto L31
L31:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v115 = v113 - v114
	goto L25
L32:
	;
	v109 = v104
	v113 = v106
	goto L31
L33:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v87 != v89 {
		v104 = v85
		v106 = v87
		goto L32
	} else {
		goto L35
	}
L34:
	;
	v104 = v98
	v106 = int32(0)
	goto L32
L35:
	;
	if v89 == int32(0) {
		v104 = v85
		v106 = v87
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v94 = v86 - int32(1)
	if v94 == int32(0) {
		v104 = v85
		v106 = v87
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v97 = int32(1)
	v98 = v85 + v97
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	if v99 != 0 {
		v84 = v84 + v97
		v85 = v98
		v86 = v94
		v87 = v99
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v124 = v115
	goto L24
L40:
	;
	v129 = v72 - int32(16)
	goto L42
L41:
	;
	v129 = v65
	goto L42
L42:
	;
	if v124 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v132 = v63
	goto L45
L44:
	;
	v132 = v72 + int32(16)
	goto L45
L45:
	;
	if base.Ui32(v132) <= base.Ui32(v129) {
		v63 = v132
		v65 = v129
		goto L20
	} else {
		goto L46
	}
L46:
	;
	goto L21
}
func F_EncodeSpecialDate(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	if l0 != int32(2147483647) {
		if l0 == int32(-2147483648) {
			v8 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1055])))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v8)
			v11 = *(*int64)(unsafe.Add(mBase, _consts[1056]))
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v11
			return
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(340426), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_errfinish(m, int32(475273), int32(309), int32(340447))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v27 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1057])))
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v27)
		v30 = *(*int64)(unsafe.Add(mBase, _consts[1058]))
		*(*int64)(unsafe.Add(mBase, uint32(l1))) = v30
		return
	}
}
