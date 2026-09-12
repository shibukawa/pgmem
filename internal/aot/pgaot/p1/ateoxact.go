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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[751])))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[750]))
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
	v95 = *(*int32)(unsafe.Add(mBase, _consts[752]))
	if l0 == int32(0) {
		v119 = v95
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v23 = int32(1)
	v26 = v16
	v27 = v20
	goto L7
L5:
	;
	goto L6
L6:
	;
	v84 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[751])) = uint8(v84)
	goto L3
L7:
	;
	v31 = v23 * int32(48)
	v32 = v27 + v31
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+4)))
	if v33&int32(3) == int32(0) {
		v70 = v26
		v71 = v27
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v73 = v23 + int32(1)
	if base.Ui32(v73) < base.Ui32(v70) {
		v23 = v73
		v26 = v70
		v27 = v71
		goto L7
	} else {
		goto L21
	}
L10:
	;
	if v33&int32(2) == int32(0) {
		v70 = v26
		v71 = v27
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	if v42 == int32(0) {
		v70 = v26
		v71 = v27
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v47 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if v47 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v50+v31)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v52
	F_errmsg_internal(m, int32(255730), v11+int32(16))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_FileClose(m, v23)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L20
	}
L18:
	;
	F_errfinish(m, int32(497922), int32(3302), int32(164768))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[750]))
	v69 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v70 = v67
	v71 = v69
	goto L9
L21:
	;
	goto L8
L22:
	;
	if int32(0) < v119 {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	if v95 <= int32(0) {
		v119 = v95
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v102 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L25
	}
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[752]))
	if v102 == int32(0) {
		v119 = v105
		goto L22
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v105
	F_errmsg_internal(m, int32(255662), v11)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(497922), int32(3314), int32(164768))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _consts[752]))
	v119 = v118
	goto L22
L29:
	;
	goto L32
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[758])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[759])) = int32(0)
	m.G0 = v11 + int32(32)
	return
L32:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _consts[753]))
	v133 = F_FreeDesc(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L13
	} else {
		goto L34
	}
L33:
	;
	goto L31
L34:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[752]))
	if int32(0) < v136 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
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
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
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
	*(*int32)(unsafe.Add(mBase, _consts[1187])) = int32(0)
	m.G0 = v5 + int32(16)
	return
L2:
	;
	v9 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[1187]))
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
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v14<<(uint(int32(2))%32))+uint32(_consts[1188])))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v24
	F_errmsg_internal(m, int32(238605), v5)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v35 = v14 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, _consts[1187]))
	if v35 < v37 {
		v14 = v35
		goto L4
	} else {
		goto L13
	}
L11:
	;
	F_errfinish(m, int32(496116), int32(1933), int32(166819))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
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
	var v26 int32
	_ = v26
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
	*(*int32)(unsafe.Add(mBase, _consts[663])) = int32(0)
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _consts[663]))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v16 = F_LWLockAcquire(m, v12+int32(5504), int32(1))
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
	v19 = *(*int32)(unsafe.Add(mBase, _consts[663]))
	if v19 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v72+int32(5504))
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
	v26 = v2
	goto L9
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v26<<(uint(int32(2))%32))))
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
	v64 = v26 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v64 < v65 {
		v26 = v64
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
