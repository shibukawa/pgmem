package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dshash_create(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_palloc(m, int32(44))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = F_dsa_allocate_extended(m, l0, int32(2580), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = v23
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+12)) = v25
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+20)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = l2
	v30 = F_dsa_get_address(m, l0, v20)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v20
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(1979673120)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+2568)) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+2568))
	v46 = int32(0)
	goto L5
L5:
	;
	v55 = v40 + int32(8) + v46*int32(20)
	F_LWLockInitialize(m, v55, v43)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+2572)) = int32(7)
	v69 = F_dsa_allocate_extended(m, l0, int32(512), int32(6))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = int32(0)
	v61 = v46 + int32(1)
	if v61 != int32(128) {
		v46 = v61
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+2576)) = v69
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+2576))
	if v74 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_dsa_free(m, l0, v20)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v100 = F_dsa_get_address(m, l0, v74)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L19
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(_a_F_dshash_create_0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errmsg(m, int32(_a_F_dshash_create_1), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(512)
	v93 = F_errdetail(m, int32(_a_F_dshash_create_2), v11)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_dshash_create_3), int32(259), int32(_a_F_dshash_create_4))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v100
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+2572))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v104
	m.G0 = v11 + int32(16)
	return v14
}
func F_dshash_seq_init(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	v3 = l2
	v4 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v4
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(-1)
	return
}
