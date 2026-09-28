package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BloomNewBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(32)
	return v77
L2:
	;
	return int32(0)
L3:
	;
	if v9 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v9
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v7)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v67
	v69 = int32(8)
	v71 = int32(0)
	v74 = F_ExtendBufferedRel(m, v7+v69, v71, v71, v69)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L24
	}
L7:
	;
	v19 = F_ReadBuffer(m, l0, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v21 = F_ConditionalLockBuffer(m, v19)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v21 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v19 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	F_ReleaseBuffer(m, v19)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L21
	}
L14:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+14)))
	if v41 == int32(0) {
		v77 = v19
		goto L1
	} else {
		goto L18
	}
L15:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_BloomNewBuffer[0]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+(v19^int32(-1))<<(uint(int32(2))%32))))
	v40 = v32
	goto L14
L16:
	;
	goto L17
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_BloomNewBuffer[1]))
	v40 = v34 + v19<<(uint(int32(13))%32) + int32(-8192)
	goto L14
L18:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v44)+2)))
	if v46&int32(2) != 0 {
		v77 = v19
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_UnlockBuffer(m, v19)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	goto L13
L21:
	;
	v54 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	if v54 != int32(-1) {
		v16 = v54
		goto L7
	} else {
		goto L23
	}
L23:
	;
	goto L8
L24:
	;
	v77 = v74
	goto L1
}
func F_bloom_get_procinfo(m *base.Module, l0 int32, l1 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14233(m, l0, l1, int32(_a_F_bloom_get_procinfo_0), int32(741), int32(_a_F_bloom_get_procinfo_1))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
