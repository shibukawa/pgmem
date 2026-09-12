package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_px_find_digest(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	v10 = int32(556880)
	v11 = l0
	goto L3
L1:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	F_ResourceOwnerEnlarge(m, v114)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L37
	} else {
		goto L38
	}
L2:
	;
	if v48 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v14 == v15 {
		v37 = v14
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v48 = int32(0)
	goto L2
L5:
	;
	v39 = int32(1)
	if v37 != 0 {
		v10 = v10 + v39
		v11 = v11 + v39
		goto L3
	} else {
		goto L14
	}
L6:
	;
	if base.Ui32((v14-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v25 = v14 | int32(32)
	goto L9
L8:
	;
	v25 = v14
	goto L9
L9:
	;
	if base.Ui32((v15-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v34 = v15 | int32(32)
	goto L12
L11:
	;
	v34 = v15
	goto L12
L12:
	;
	if v25 == v34 {
		v37 = v25
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v48 = v25 - v34
	goto L2
L14:
	;
	goto L4
L15:
	;
	v112 = int32(4393744)
	goto L1
L16:
	;
	goto L17
L17:
	;
	v55 = int32(0)
	goto L18
L18:
	;
	v58 = v55 + int32(1)
	if v58 == int32(22) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v112 = v64 + int32(4393744)
	goto L1
L20:
	;
	return int32(-2)
L21:
	;
	goto L22
L22:
	;
	v64 = v58 << (uint(int32(3)) % 32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+uint32(_consts[1452])))
	v70 = v67
	v71 = l0
	goto L24
L23:
	;
	if v108 != 0 {
		v55 = v58
		goto L18
	} else {
		goto L36
	}
L24:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v74 == v75 {
		v97 = v74
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v108 = int32(0)
	goto L23
L26:
	;
	v99 = int32(1)
	if v97 != 0 {
		v70 = v70 + v99
		v71 = v71 + v99
		goto L24
	} else {
		goto L35
	}
L27:
	;
	if base.Ui32((v74-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v85 = v74 | int32(32)
	goto L30
L29:
	;
	v85 = v74
	goto L30
L30:
	;
	if base.Ui32((v75-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v94 = v75 | int32(32)
	goto L33
L32:
	;
	v94 = v75
	goto L33
L33:
	;
	if v85 == v94 {
		v97 = v85
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v108 = v85 - v94
	goto L23
L35:
	;
	goto L25
L36:
	;
	goto L19
L37:
	;
	return int32(0)
L38:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v121 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	v123 = F_MemoryContextAlloc(m, v121, int32(16))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v125 = m.Env.Pgmem_hash_create(m, v119)
	mBase = m.M
	if v125 <= int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_pfree(m, v123)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L37
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v132 = m.Env.Pgmem_hash_info(m, v125)
	mBase = m.M
	if v132 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	return int32(-2)
L44:
	;
	m.Env.Pgmem_hash_free(m, v125)
	mBase = m.M
	F_pfree(m, v123)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L37
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v123)+8)) = int32(base.Ui32(v132) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+4)) = v132 & int32(65535)
	v148 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+12)) = v148
	F_ResourceOwnerRemember(m, v148, v123, int32(4394168))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L37
	} else {
		goto L48
	}
L47:
	;
	return int32(-8)
L48:
	;
	v154 = F_palloc(m, int32(28))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L37
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+24)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v154)+20)) = int32(6830)
	*(*int32)(unsafe.Add(mBase, uint32(v154)+16)) = int32(6831)
	*(*int32)(unsafe.Add(mBase, uint32(v154)+12)) = int32(6832)
	*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = int32(6833)
	*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = int32(6834)
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = int32(6835)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v154
	return int32(0)
}
