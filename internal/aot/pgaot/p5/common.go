package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_common_entry_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	v6 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.F64_lt(v6, v7) != 0 {
		v10 = int32(-1)
	} else {
		v10 = base.F64_gt(v6, v7)
	}
	return v10
}
func F_common_prefix_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = int32(0)
	goto L2
L1:
	;
	return v87
L2:
	;
	v20 = int32(0)
	if v10 == v20 {
		v30 = v20
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v71 = int32(0)
	if v10 != 0 {
		goto L26
	} else {
		goto L27
	}
L4:
	;
	if v9 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v24 <= v13 {
		v30 = int32(0)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v30 = v26 + v13<<(uint(int32(2))%32)
	goto L4
L7:
	;
	goto L3
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v33 <= v13 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if v30 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v40 = v37 + v13<<(uint(int32(2))%32)
	if v40 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if base.Ui32(v46) < base.Ui32(v44) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(-1)
L13:
	;
	goto L14
L14:
	;
	v50 = int32(1)
	if base.Ui32(v44) < base.Ui32(v46) {
		v87 = v50
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	if base.Ui32(v53) < base.Ui32(v52) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(-1)
L17:
	;
	goto L18
L18:
	;
	if base.Ui32(v52) < base.Ui32(v53) {
		v87 = v50
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+17)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+17)))
	if v59 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v13 = v13 + int32(1)
	goto L2
L21:
	;
	if v58&int32(1) != 0 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v58&int32(1) != 0 {
		v87 = v50
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v87 = int32(-1)
	goto L1
L25:
	;
	goto L20
L26:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v75 = v74
	goto L28
L27:
	;
	v75 = v71
	goto L28
L28:
	;
	if v9 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v77 = v76
	goto L31
L30:
	;
	v77 = v71
	goto L31
L31:
	;
	if v77 < v75 {
		v87 = int32(-1)
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v10 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v80 = v79
	goto L35
L34:
	;
	v80 = v71
	goto L35
L35:
	;
	if v9 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v82 = v81
	goto L38
L37:
	;
	v82 = v71
	goto L38
L38:
	;
	v87 = base.B2i32(v80 < v82)
	goto L1
}
