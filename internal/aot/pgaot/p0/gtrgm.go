package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtrgm_options(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(281935), int32(159184), int32(12), int32(1), int32(2024))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_gtrgm_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == v2 {
		v29 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v29&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if v16 == int32(0) {
		v29 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v19 != int32(7) {
		v29 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v22 != int32(17) {
		v29 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+24)))
	v29 = v25 ^ int32(1)
	goto L2
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = F_get_fn_opclass_options(m, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v38 = int32(12)
	goto L9
L9:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+4)))
	if v39&int32(2) != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	return int32(0)
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v38 = v37
	goto L9
L12:
	;
	return v8
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v129)
	goto L12
L14:
	;
	v43 = v39 & int32(4)
	if v43 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v80 = int32(2)
	v82 = int32(5)
	v83 = int32(base.Ui32(v79)>>(uint(v80)%32)) - v82
	v84 = int32(3)
	v85 = base.I32_div_u_s(v83, v84)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v92 = base.I32_div_u_s(int32(base.Ui32(v86)>>(uint(v80)%32))-v82, v84)
	if v85 != v92 {
		goto L29
	} else {
		goto L30
	}
L17:
	;
	if v43 != 0 {
		v129 = v2
		goto L13
	} else {
		goto L20
	}
L18:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)))
	if v46&int32(4) == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v129 = int32(1)
	goto L13
L20:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)))
	if v52&int32(4) != 0 {
		v129 = v2
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v55 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v55)
	v57 = int32(0)
	if v38 <= v57 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	v60 = int32(5)
	v64 = v57
	goto L23
L23:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+(v10+v60)))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+(v9+v60)))))
	if v72 == v74 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v129 = v2
	goto L13
L25:
	;
	v77 = v64 + int32(1)
	if v38 != v77 {
		v64 = v77
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	goto L12
L29:
	;
	v129 = v2
	goto L13
L30:
	;
	goto L31
L31:
	;
	v94 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v94)
	if base.Ui32(v83) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	v98 = int32(5)
	v102 = int32(1)
	if base.Ui32(v85) <= base.Ui32(v102) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v105 = v102
	goto L35
L34:
	;
	v105 = v85
	goto L35
L35:
	;
	v107 = int32(0)
	goto L36
L36:
	;
	v115 = v107 * int32(3)
	v119 = *(*int32)(unsafe.Add(mBase, _consts[1097]))
	v120 = m.T0[v119].(func(*base.Module, int32, int32) int32)(m, v10+v98+v115, v115+(v9+v98))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L10
	} else {
		goto L38
	}
L37:
	;
	goto L12
L38:
	;
	if v120 != 0 {
		v129 = v2
		goto L13
	} else {
		goto L39
	}
L39:
	;
	v123 = v107 + int32(1)
	if v105 != v123 {
		v107 = v123
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
}
