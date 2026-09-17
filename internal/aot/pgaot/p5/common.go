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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
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
	if v10 != 0 {
		goto L25
	} else {
		goto L26
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
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if base.B2i32(v30 == int32(0))|base.B2i32(v35 <= v13) != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v38 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38+v13<<(uint(int32(2))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if base.Ui32(v47) < base.Ui32(v42) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(-1)
L12:
	;
	goto L13
L13:
	;
	v51 = int32(1)
	if base.Ui32(v42) < base.Ui32(v47) {
		v87 = v51
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	if base.Ui32(v54) < base.Ui32(v53) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(-1)
L16:
	;
	goto L17
L17:
	;
	if base.Ui32(v53) < base.Ui32(v54) {
		v87 = v51
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+17)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+17)))
	if v60 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v13 = v13 + int32(1)
	goto L2
L20:
	;
	if v59&int32(1) != 0 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v59&int32(1) != 0 {
		v87 = v51
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v87 = int32(-1)
	goto L1
L24:
	;
	goto L19
L25:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v75 = v73
	goto L27
L26:
	;
	v75 = int32(0)
	goto L27
L27:
	;
	if v9 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v78 = v76
	goto L30
L29:
	;
	v78 = int32(0)
	goto L30
L30:
	;
	if v78 < v75 {
		v87 = int32(-1)
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v10 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v82 = v80
	goto L34
L33:
	;
	v82 = int32(0)
	goto L34
L34:
	;
	if v9 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v85 = v83
	goto L37
L36:
	;
	v85 = int32(0)
	goto L37
L37:
	;
	v87 = base.B2i32(v82 < v85)
	goto L1
}
