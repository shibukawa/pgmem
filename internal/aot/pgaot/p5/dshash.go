package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dshash_attach(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v7 = F_palloc(m, int32(44))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = v12
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+12)) = v14
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+20)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = l3
		v19 = F_dsa_get_address(m, l0, l2)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v19
			return v7
		}
	}
}
func F_dshash_delete_key(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = m.T0[v9].(func(*base.Module, int32, int32, int32) int32)(m, l1, v7, v8)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v16 = int32(base.Ui32(v10) >> (uint(int32(25)) % 32))
	v23 = F_LWLockAcquire(m, v14+v16*int32(20)+int32(8), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+2572))
	if v25 == v27 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v48 = v38 + int32(base.Ui32(v10)>>(uint(int32(32)-v39)%32))<<(uint(int32(2))%32)
	goto L10
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v38 = v29
	v39 = v25
	goto L4
L6:
	;
	goto L7
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+2576))
	v32 = F_dsa_get_address(m, v30, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v32
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+2572))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v36
	v38 = v32
	v39 = v36
	goto L4
L9:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_LWLockRelease(m, v82+v16*int32(20)+int32(8))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L17
	}
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v52 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	F_dsa_free(m, v66, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L16
	}
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v56 = F_dsa_get_address(m, v55, v52)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v63 = m.T0[v62].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v56+int32(8), v60, v61)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v63 != 0 {
		v48 = v56
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v65
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v76 = v71 + v16*int32(20) + int32(24)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v77 - int32(1)
	goto L9
L17:
	;
	return base.B2i32(v52 != int32(0))
}
func F_dshash_get_hash_table_handle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	return v3
}
func F_dshash_memhash(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	v4 = F_hash_bytes(m, l0, l1)
	return v4
}
