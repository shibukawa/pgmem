package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_KnownAssignedXidsRemove(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v19 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg_internal(m, int32(_a_F_KnownAssignedXidsRemove_0), v15)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemove[0]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	v35 = v33 - int32(1)
	if v35 < v32 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_errfinish(m, int32(_a_F_KnownAssignedXidsRemove_1), int32(_a_F_KnownAssignedXidsRemove_2), int32(_a_F_KnownAssignedXidsRemove_3))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	m.G0 = v15 + int32(16)
	return
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemove[1]))
	v46 = v35
	v47 = v32
	goto L10
L10:
	;
	v53 = v46 + v47
	v54 = int32(2)
	v55 = base.I32_div_s(v53, v54)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v38+v55<<(uint(v54)%32))))
	if v59 != l0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v53 < int32(-1) {
		goto L8
	} else {
		goto L25
	}
L12:
	;
	if base.B2i32(base.B2i32(base.Ui32(l0) < base.Ui32(int32(3))) == int32(0))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(v59)) != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	goto L11
L15:
	;
	v72 = int32(base.Ui32(l0-v59) >> (uint(int32(31)) % 32))
	goto L17
L16:
	;
	v72 = base.B2i32(base.Ui32(l0) < base.Ui32(v59))
	goto L17
L17:
	;
	if v72 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v73 = v47
	goto L20
L19:
	;
	v73 = v55 + int32(1)
	goto L20
L20:
	;
	if v72 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v76 = v55 - int32(1)
	goto L23
L22:
	;
	v76 = v46
	goto L23
L23:
	;
	if v73 <= v76 {
		v46 = v76
		v47 = v73
		goto L10
	} else {
		goto L24
	}
L24:
	;
	goto L8
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemove[2]))
	v82 = v81 + v55
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v83 != int32(1) {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	v86 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v82))) = uint8(v86)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v88 - int32(1)
	if v32 != v55 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsRemove[2]))
	v96 = v32
	goto L29
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v117
	goto L8
L29:
	;
	v108 = v96 + int32(1)
	if v108 < v33 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v114 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v114
	v117 = v114
	goto L28
L31:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v108))))
	if v111 == int32(0) {
		v96 = v108
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	goto L30
L34:
	;
	v117 = v108
	goto L28
}
