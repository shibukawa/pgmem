package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_sts_initialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
	v13 = F_strlen(m, l4)
	mBase = m.M
	if base.Ui32(v13) < base.Ui32(int32(64)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = l0 + int32(12)
	if (l4^v17)&int32(3) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L30
	} else {
		goto L34
	}
L4:
	;
	if int32(0) < l1 {
		goto L25
	} else {
		goto L26
	}
L5:
	;
	goto L4
L6:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v72))) = uint8(v71)
	if v71&int32(255) == int32(0) {
		goto L5
	} else {
		goto L21
	}
L7:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v70 = l4
	v71 = v23
	v72 = v17
	goto L6
L8:
	;
	goto L9
L9:
	;
	if l4&int32(3) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v27 = l4
	v29 = v17
	goto L13
L11:
	;
	v41 = l4
	v43 = v17
	goto L12
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v48 = int32(-2139062144)
	if (int32(16843008)-v45|v45)&v48 != v48 {
		v70 = v41
		v71 = v45
		v72 = v43
		goto L6
	} else {
		goto L17
	}
L13:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v30)
	if v30 == int32(0) {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v41 = v37
	v43 = v35
	goto L12
L15:
	;
	v34 = int32(1)
	v35 = v29 + v34
	v37 = v27 + v34
	if v37&int32(3) != 0 {
		v27 = v37
		v29 = v35
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v53 = v41
	v54 = v45
	v55 = v43
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v54
	v57 = int32(4)
	v58 = v55 + v57
	v60 = v53 + v57
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v65 = int32(-2139062144)
	if (int32(16843008)-v62|v62)&v65 == v65 {
		v53 = v60
		v54 = v62
		v55 = v58
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v70 = v60
	v71 = v62
	v72 = v58
	goto L6
L20:
	;
	goto L19
L21:
	;
	v79 = v70
	v81 = v72
	goto L22
L22:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)) = uint8(v82)
	v84 = int32(1)
	if v82 != 0 {
		v79 = v79 + v84
		v81 = v81 + v84
		goto L22
	} else {
		goto L24
	}
L23:
	;
	goto L5
L24:
	;
	goto L23
L25:
	;
	v101 = int32(0)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v126 = F_palloc0(m, int32(64))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L30
	} else {
		goto L33
	}
L28:
	;
	v105 = l0 + int32(76) + v101*int32(28)
	F_LWLockInitialize(m, v105, int32(79))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L27
L30:
	;
	return int32(0)
L31:
	;
	v111 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v105)+24)) = uint8(v111)
	*(*int64)(unsafe.Add(mBase, uint32(v105)+16)) = int64(0)
	v116 = v101 + int32(1)
	if v116 != l1 {
		v101 = v116
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = l2
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_sts_initialize[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+12)) = v132
	return v126
L34:
	;
	F_errmsg_internal(m, int32(_a_F_sts_initialize_0), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_sts_initialize_1), int32(142), int32(_a_F_sts_initialize_2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
