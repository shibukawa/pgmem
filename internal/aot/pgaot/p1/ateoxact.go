package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOXact_Files(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AtEOXact_Files[0])))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Files[1]))
	if base.Ui32(int32(2)) <= base.Ui32(v16) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v95 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Files[2]))
	if base.B2i32(l0 == v95)|base.B2i32(v98 <= v95) != 0 {
		v121 = v98
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Files[3]))
	v23 = int32(1)
	v25 = v16
	v26 = v20
	goto L7
L5:
	;
	goto L6
L6:
	;
	v85 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_AtEOXact_Files[0])) = uint8(v85)
	goto L3
L7:
	;
	v31 = v23 * int32(48)
	v32 = v26 + v31
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+4)))
	v36 = int32(0)
	if base.B2i32(v33&int32(3) == v36)|base.B2i32(v33&int32(2) == v36) != 0 {
		v71 = v25
		v72 = v26
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v74 = v23 + int32(1)
	if base.Ui32(v74) < base.Ui32(v71) {
		v23 = v74
		v25 = v71
		v26 = v72
		goto L7
	} else {
		goto L20
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	if v43 == int32(0) {
		v71 = v25
		v72 = v26
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v48 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Files[3]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v51+v31)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v53
	F_errmsg_internal(m, int32(_a_F_AtEOXact_Files_0), v11+int32(16))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_FileClose(m, v23)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	F_errfinish(m, int32(_a_F_AtEOXact_Files_1), int32(3302), int32(_a_F_AtEOXact_Files_2))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Files[1]))
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Files[3]))
	v71 = v68
	v72 = v70
	goto L9
L20:
	;
	goto L8
L21:
	;
	if int32(0) < v121 {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v104 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Files[2]))
	if v104 == int32(0) {
		v121 = v107
		goto L21
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v107
	F_errmsg_internal(m, int32(_a_F_AtEOXact_Files_3), v11)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_AtEOXact_Files_1), int32(3314), int32(_a_F_AtEOXact_Files_2))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Files[2]))
	v121 = v120
	goto L21
L27:
	;
	goto L30
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Files[4])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Files[5])) = int32(0)
	m.G0 = v11 + int32(32)
	return
L30:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Files[6]))
	v134 = F_FreeDesc(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L12
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Files[2]))
	if int32(0) < v137 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
}
func F_AtEOXact_HashTables(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_HashTables[0])) = int32(0)
	m.G0 = v5 + int32(16)
	return
L2:
	;
	v9 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_HashTables[0]))
	if v11 <= v9 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = v9
	goto L4
L4:
	;
	v18 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v14<<(uint(int32(2))%32))+uint32(_c_F_AtEOXact_HashTables[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v22
	F_errmsg_internal(m, int32(_a_F_AtEOXact_HashTables_0), v5)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v33 = v14 + int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_HashTables[0]))
	if v33 < v35 {
		v14 = v33
		goto L4
	} else {
		goto L13
	}
L11:
	;
	F_errfinish(m, int32(_a_F_AtEOXact_HashTables_1), int32(1933), int32(_a_F_AtEOXact_HashTables_2))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	goto L5
}
func F_AtEOXact_LogicalRepWorkers(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LogicalRepWorkers[0])) = int32(0)
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LogicalRepWorkers[0]))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LogicalRepWorkers[1]))
	v16 = F_LWLockAcquire(m, v12+int32(_a_F_AtEOXact_LogicalRepWorkers_0), int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LogicalRepWorkers[0]))
	if v19 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LogicalRepWorkers[1]))
	F_LWLockRelease(m, v72+int32(_a_F_AtEOXact_LogicalRepWorkers_0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L20
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v22 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v28 = v2
	goto L9
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v28<<(uint(int32(2))%32))))
	v36 = F_logicalrep_workers_find(m, v33, int32(1), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	v64 = v28 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v64 < v65 {
		v28 = v64
		goto L9
	} else {
		goto L19
	}
L12:
	;
	if v36 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v40 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v41 <= v40 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v44 = v40
	goto L15
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v44<<(uint(int32(2))%32))))
	F_logicalrep_worker_wakeup_ptr(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	goto L11
L17:
	;
	v56 = v44 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v56 < v57 {
		v44 = v56
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	goto L10
L20:
	;
	goto L1
}
