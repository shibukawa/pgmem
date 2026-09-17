package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_px_find_digest(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v7 = int32(0)
	goto L2
L1:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_px_find_digest[0]))
	F_ResourceOwnerEnlarge(m, v64)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v10 = v7 << (uint(int32(3)) % 32)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_px_find_digest[1])))
	v16 = v13
	v17 = l0
	goto L5
L3:
	;
	return int32(-2)
L4:
	;
	if v54 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L5:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v20 == v21 {
		v43 = v20
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v54 = int32(0)
	goto L4
L7:
	;
	v45 = int32(1)
	if v43 != 0 {
		v16 = v16 + v45
		v17 = v17 + v45
		goto L5
	} else {
		goto L16
	}
L8:
	;
	if base.Ui32((v20-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v31 = v20 | int32(32)
	goto L11
L10:
	;
	v31 = v20
	goto L11
L11:
	;
	if base.Ui32((v21-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v40 = v21 | int32(32)
	goto L14
L13:
	;
	v40 = v21
	goto L14
L14:
	;
	if v31 == v40 {
		v43 = v31
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v54 = v31 - v40
	goto L4
L16:
	;
	goto L6
L17:
	;
	v58 = v7 + int32(1)
	if v58 != int32(22) {
		v7 = v58
		goto L2
	} else {
		goto L18
	}
L18:
	;
	goto L3
L19:
	;
	return int32(0)
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_px_find_digest[2])))
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_px_find_digest[3]))
	v73 = F_MemoryContextAlloc(m, v71, int32(16))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v75 = m.Env.Pgmem_hash_create(m, v69)
	mBase = m.M
	if v75 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_pfree(m, v73)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L19
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v82 = m.Env.Pgmem_hash_info(m, v75)
	mBase = m.M
	if v82 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	return int32(-2)
L26:
	;
	m.Env.Pgmem_hash_free(m, v75)
	mBase = m.M
	F_pfree(m, v73)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L19
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = int32(base.Ui32(v82) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v82 & int32(_a_F_px_find_digest_0)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_px_find_digest[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v98
	F_ResourceOwnerRemember(m, v98, v73, int32(_a_F_px_find_digest_1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L19
	} else {
		goto L30
	}
L29:
	;
	return int32(-8)
L30:
	;
	v104 = F_palloc(m, int32(28))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+24)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v104)+20)) = int32(_a_F_px_find_digest_2)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+16)) = int32(_a_F_px_find_digest_3)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = int32(_a_F_px_find_digest_4)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+8)) = int32(_a_F_px_find_digest_5)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = int32(_a_F_px_find_digest_6)
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = int32(_a_F_px_find_digest_7)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v104
	return int32(0)
}
