package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResOwnerPrintBufferIO(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13836(m, l0, int32(_a_F_ResOwnerPrintBufferIO_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_ResOwnerPrintCatCacheList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
	v18 = F_psprintf(m, int32(_a_F_ResOwnerPrintCatCacheList_0), v7)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v18
	}
}
func F_ResOwnerReleasePGMEMDigest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	m.Env.Pgmem_hash_free(m, v5)
	mBase = m.M
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 != 0 {
		F_ResourceOwnerForget(m, v7, l0, int32(_a_F_ResOwnerReleasePGMEMDigest_0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F__equalResTarget(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v53
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v41 = F_equal(m, v39, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	if v6 == int32(0) {
		v53 = v3
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v6 == v7 {
		goto L2
	} else {
		goto L15
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v12 == int32(0))|base.B2i32(v12 != v15) != 0 {
		v33 = v12
		v34 = v15
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v34 != 0 {
		v53 = v3
		goto L1
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v18 = v7
	v19 = v6
	goto L10
L10:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v23
		v34 = v22
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v33 = v23
	v34 = v22
	goto L8
L12:
	;
	v26 = int32(1)
	if v23 == v22 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L2
L15:
	;
	return int32(0)
L16:
	;
	return int32(0)
L17:
	;
	if v41 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	goto L20
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v51 = F_equal(m, v49, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v53 = v51
	goto L1
}
