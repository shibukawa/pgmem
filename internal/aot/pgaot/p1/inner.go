package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inner_isparent(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if base.Ui32(v7) < base.Ui32(v6) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	if v6 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(1)
L5:
	;
	goto L6
L6:
	;
	v15 = int32(8)
	v19 = l0 + v15
	v20 = l1 + v15
	v21 = v6
	goto L7
L7:
	;
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19))))
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20))))
	if v24 != v25 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	return int32(1)
L9:
	;
	return int32(0)
L10:
	;
	goto L11
L11:
	;
	v29 = int32(2)
	v30 = v19 + v29
	v32 = v20 + v29
	if base.Ui32(int32(4)) <= base.Ui32(v24) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v94 != 0 {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	v94 = int32(0)
	goto L12
L14:
	;
	v68 = v63
	v69 = v64
	v70 = v65
	goto L24
L15:
	;
	if (v30|v32)&int32(3) != 0 {
		v63 = v30
		v64 = v32
		v65 = v24
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v56 = v30
	v57 = v32
	v58 = v24
	goto L17
L17:
	;
	if v58 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v40 = v30
	v41 = v32
	v42 = v24
	goto L19
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v45 != v46 {
		v63 = v40
		v64 = v41
		v65 = v42
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v56 = v51
	v57 = v49
	v58 = v53
	goto L17
L21:
	;
	v48 = int32(4)
	v49 = v41 + v48
	v51 = v40 + v48
	v53 = v42 - v48
	if base.Ui32(int32(3)) < base.Ui32(v53) {
		v40 = v51
		v41 = v49
		v42 = v53
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v63 = v56
	v64 = v57
	v65 = v58
	goto L14
L24:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v73 == v74 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v94 = v73 - v74
	goto L12
L26:
	;
	v76 = int32(1)
	v81 = v70 - v76
	if v81 != 0 {
		v68 = v68 + v76
		v69 = v69 + v76
		v70 = v81
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	goto L13
L30:
	;
	return int32(0)
L31:
	;
	goto L32
L32:
	;
	v97 = int32(9)
	v99 = int32(_a_F_inner_isparent_0)
	v107 = int32(1)
	if v107 < v21 {
		v19 = v19 + (v24+v97)&v99
		v20 = v20 + (v25+v97)&v99
		v21 = v21 - v107
		goto L7
	} else {
		goto L33
	}
L33:
	;
	goto L8
}
