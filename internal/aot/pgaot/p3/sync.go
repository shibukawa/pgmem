package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_SyncRepGetCandidateStandbys(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v73 int64
	_ = v73
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[0]))
	v10 = F_palloc_mul(m, int32(48), v9)
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
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v10
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[1]))
	if v16 == int32(0) {
		v111 = v2
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v111
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[0]))
	if int32(0) < v20 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = v2
	v28 = v2
	goto L8
L6:
	;
	v92 = v16
	v94 = v2
	goto L7
L7:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+8)))
	if v97 != 0 {
		v111 = v94
		goto L3
	} else {
		goto L19
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = v29 + v26*int32(48)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[2]))
	v37 = v34 + v28*int32(96)
	v39 = v37 + int32(88)
	v42 = base.AtomicRmwXchg32(m, v37, int32(164), int32(1))
	if v42 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[1]))
	v92 = v90
	v94 = v83
	goto L7
L10:
	;
	F_s_lock(m, v37+int32(164), int32(_a_F_SyncRepGetCandidateStandbys_0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v39)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v51
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v39)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v53
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v39)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+24)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v57
	v59 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v39)+76)), uint32(v59))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if base.B2i32(v62 == v59)|base.B2i32(base.Ui32(v50-int32(5)) < base.Ui32(int32(-2))) != 0 {
		v83 = v26
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v85 = v28 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[0]))
	if v85 < v87 {
		v26 = v83
		v28 = v85
		goto L8
	} else {
		goto L18
	}
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	if v70 == int32(0) {
		v83 = v26
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
	if v73 == int64(0) {
		v83 = v26
		goto L14
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v28
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[3]))
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+40)) = uint8(base.B2i32(v39 == v78))
	v83 = v26 + int32(1)
	goto L14
L18:
	;
	goto L9
L19:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v94 <= v98 {
		v111 = v94
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_pg_qsort(m, v100, v94, int32(48), int32(1102))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[1]))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v111 = v107
	goto L3
}
func F_SyncScanShmemRequest(m *base.Module, l0 int32) {
	var v6 int32
	_ = v6
	Fn14222(m, l0, int32(_a_F_SyncScanShmemRequest_0), int64(488), int32(_a_F_SyncScanShmemRequest_1))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
