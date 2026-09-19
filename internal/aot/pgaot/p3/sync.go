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
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[0]))
	v11 = F_palloc(m, v8*int32(48))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v11
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[1]))
	if v17 == int32(0) {
		v114 = v2
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v114
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[0]))
	if int32(0) < v21 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v27 = v2
	v29 = v2
	goto L8
L6:
	;
	v95 = v17
	v97 = v2
	goto L7
L7:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+8)))
	if v100 != 0 {
		v114 = v97
		goto L3
	} else {
		goto L19
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = v30 + v27*int32(48)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[2]))
	v38 = v35 + v29*int32(96)
	v40 = v38 + int32(88)
	v43 = base.AtomicRmwXchg32(m, v38, int32(164), int32(1))
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[1]))
	v95 = v93
	v97 = v86
	goto L7
L10:
	;
	F_s_lock(m, v38+int32(164), int32(_a_F_SyncRepGetCandidateStandbys_0), int32(779), int32(_a_F_SyncRepGetCandidateStandbys_1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v40)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v54
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v40)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v40)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+24)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v40)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v60
	v62 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v40)+76)), uint32(v62))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if base.B2i32(v65 == v62)|base.B2i32(base.Ui32(v53-int32(5)) < base.Ui32(int32(-2))) != 0 {
		v86 = v27
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v88 = v29 + int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[0]))
	if v88 < v90 {
		v27 = v86
		v29 = v88
		goto L8
	} else {
		goto L18
	}
L15:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
	if v73 == int32(0) {
		v86 = v27
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v33)+16))
	if v76 == int64(0) {
		v86 = v27
		goto L14
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+36)) = v29
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[3]))
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+40)) = uint8(base.B2i32(v40 == v81))
	v86 = v27 + int32(1)
	goto L14
L18:
	;
	goto L9
L19:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v97 <= v101 {
		v114 = v97
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_pg_qsort(m, v103, v97, int32(48), int32(1028))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[1]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v114 = v110
	goto L3
}
