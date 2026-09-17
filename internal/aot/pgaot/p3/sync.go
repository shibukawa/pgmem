package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
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
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
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
		v113 = v2
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v113
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
	v28 = v2
	goto L8
L6:
	;
	v94 = v17
	v96 = v2
	goto L7
L7:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+8)))
	if v99 != 0 {
		v113 = v96
		goto L3
	} else {
		goto L19
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[2]))
	v34 = v31 + v28*int32(96)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+164))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+164)) = int32(1)
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[1]))
	v94 = v92
	v96 = v85
	goto L7
L10:
	;
	F_s_lock(m, v34+int32(164), int32(_a_F_SyncRepGetCandidateStandbys_0), int32(779), int32(_a_F_SyncRepGetCandidateStandbys_1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v48 = v36 + v27*int32(48)
	v50 = v34 + int32(88)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v50)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v48)+8)) = v54
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v50)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v48)+16)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v50)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v48)+24)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v50)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = v60
	v62 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+76)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if base.B2i32(v64 == v62)|base.B2i32(base.Ui32(v53-int32(5)) < base.Ui32(int32(-2))) != 0 {
		v85 = v27
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v87 = v28 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[0]))
	if v87 < v89 {
		v27 = v85
		v28 = v87
		goto L8
	} else {
		goto L18
	}
L15:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v48)+32))
	if v72 == int32(0) {
		v85 = v27
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v48)+16))
	if v75 == int64(0) {
		v85 = v27
		goto L14
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+36)) = v28
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[3]))
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+40)) = uint8(base.B2i32(v50 == v80))
	v85 = v27 + int32(1)
	goto L14
L18:
	;
	goto L9
L19:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v96 <= v100 {
		v113 = v96
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_pg_qsort(m, v102, v96, int32(48), int32(1028))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepGetCandidateStandbys[1]))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v113 = v109
	goto L3
}
