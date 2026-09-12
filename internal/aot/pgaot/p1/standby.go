package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_StandbyLockTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[796])) = int32(1)
	return
}
func F_StandbyRecoverPreparedTransactions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v1 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v8 = F_LWLockAcquire(m, v4+int32(2304), v1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(0) < v12 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v15 = v11
	v16 = v1
	goto L6
L4:
	;
	goto L5
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v39+int32(2304))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L14
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15+v16<<(uint(int32(2))%32))+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v20)+16))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+45)))
	v26 = F_ProcessTwoPhaseBuffer(m, v21, v22, v23, int32(1), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_pfree(m, v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v31 = v16 + int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v31 < v34 {
		v15 = v33
		v16 = v31
		goto L6
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	goto L7
L14:
	;
	return
}
func F_standby_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+48)))
	v16 = v14 & int32(240)
	switch v16 - int32(16) {
	case 0:
		goto L3
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L1
	case 16:
		goto L2
	default:
		goto L4
	}
L1:
	;
	m.G0 = v10 + int32(96)
	return
L2:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)))
	F_standby_desc_invalidations(m, l0, v144, v13+int32(16), v147, v148, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L9
	} else {
		goto L32
	}
L3:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v46
	F_appendStringInfo(m, l0, int32(54765), v10+int32(80))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L9
	} else {
		goto L12
	}
L4:
	;
	if v16 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v19 <= int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v27 = int32(0)
	goto L7
L7:
	;
	v34 = v13 + int32(4) + v27*int32(12)
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v34)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v35
	F_appendStringInfo(m, l0, int32(731233), v10)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L1
L9:
	;
	return
L10:
	;
	v43 = v27 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v43 < v44 {
		v27 = v43
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v57 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)))
	if v100 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v57
	F_appendStringInfo(m, l0, int32(547049), v10-int32(-64))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v66 <= int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v74 = int32(0)
	goto L17
L17:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(24)+v74<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v82
	F_appendStringInfo(m, l0, int32(59414), v10+int32(48))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	goto L13
L19:
	;
	v90 = v74 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v90 < v91 {
		v74 = v90
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	F_appendStringInfoString(m, l0, int32(440032))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L9
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v106 <= int32(0) {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v106
	F_appendStringInfo(m, l0, int32(547022), v10+int32(32))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v115 <= int32(0) {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v123 = int32(0)
	goto L28
L28:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(24)+(v128+v123)<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v133
	F_appendStringInfo(m, l0, int32(59414), v10+int32(16))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L30
	}
L29:
	;
	goto L1
L30:
	;
	v141 = v123 + int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v141 < v142 {
		v123 = v141
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	goto L1
}
func F_standby_priority_comparator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v5 != v6 {
		return v5 - v6
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
		return v10 - v11
	}
}
