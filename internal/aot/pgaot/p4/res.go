package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResOwnerPrintCatCache(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v12 | v13<<(uint(int32(16))%32)
	v30 = F_psprintf(m, int32(_a_F_ResOwnerPrintCatCache_0), v10)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int32(0)
	} else {
		m.G0 = v10 + int32(32)
		return v30
	}
}
func F_ResOwnerPrintDSM(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v7
	v10 = F_psprintf(m, int32(_a_F_ResOwnerPrintDSM_0), v5)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(16)
		return v10
	}
}
func F_ResOwnerReleaseBufferIO(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferIO[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(_a_F_ResOwnerReleaseBufferIO_0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(_a_F_ResOwnerReleaseBufferIO_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_ResOwnerReleaseBufferIO_2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v24 = v11 + l0<<(uint(int32(6))%32)
	v26 = v24 - int32(40)
	v27 = int32(_a_F_ResOwnerReleaseBufferIO_3)
	v29 = base.AtomicRmwOr32(m, v26, v2, v27)
	if v29&v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L4
L2:
	;
	v46 = v29
	goto L3
L3:
	;
	v54 = int32(_a_F_ResOwnerReleaseBufferIO_4)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferIO[1]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))+8))
	if v57 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	F_perform_spin_delay(m, v8+int32(8))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v46 = v43
	goto L3
L6:
	;
	return
L7:
	;
	v41 = int32(_a_F_ResOwnerReleaseBufferIO_3)
	v43 = base.AtomicRmwOr32(m, v26, int32(0), v41)
	if v43&v41 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	if v46&int32(16777216) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L10:
	;
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferIO[1])) = v72
	goto L10
L12:
	;
	if int32(999) < v55 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v55 < int32(11) {
		goto L10
	} else {
		goto L19
	}
L15:
	;
	v62 = int32(900)
	if v62 <= v55 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v65 = v62
	goto L18
L17:
	;
	v65 = v55
	goto L18
L18:
	;
	v72 = v65 + int32(100)
	goto L11
L19:
	;
	v72 = v55 - int32(1)
	goto L11
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(_a_F_ResOwnerReleaseBufferIO_0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(_a_F_ResOwnerReleaseBufferIO_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_ResOwnerReleaseBufferIO_2)
	v139 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v139
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v143 = int32(_a_F_ResOwnerReleaseBufferIO_3)
	v145 = base.AtomicRmwOr32(m, v26, v139, v143)
	if v145&v143 != 0 {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v46 & int32(-20971521)
	goto L20
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v46 & int32(-4194305)
	if v46&int32(134217728) == int32(0) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v90 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	if v90 == int32(0) {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(_a_F_ResOwnerReleaseBufferIO_5))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(48))))
	v101 = v8 + int32(8)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(60))))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-64))))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(56))))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(52))))
	F_GetRelationPath(m, v101, v104, v107, v110, int32(-1), v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v101
	F_errmsg(m, int32(_a_F_ResOwnerReleaseBufferIO_6), v8)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	F_errdetail(m, int32(_a_F_ResOwnerReleaseBufferIO_7), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_ResOwnerReleaseBufferIO_2), int32(_a_F_ResOwnerReleaseBufferIO_8), int32(_a_F_ResOwnerReleaseBufferIO_9))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	goto L20
L32:
	;
	goto L35
L33:
	;
	v162 = v145
	goto L34
L34:
	;
	v170 = int32(_a_F_ResOwnerReleaseBufferIO_4)
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferIO[1]))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))+8))
	if v173 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	F_perform_spin_delay(m, v8+int32(8))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L37
	}
L36:
	;
	v162 = v159
	goto L34
L37:
	;
	v157 = int32(_a_F_ResOwnerReleaseBufferIO_3)
	v159 = base.AtomicRmwOr32(m, v26, int32(0), v157)
	if v159&v157 != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v162&int32(-205520897) | int32(134217728)
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferIO[2]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(44))))
	F_ConditionVariableBroadcast(m, v196+v199<<(uint(int32(4))%32))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L6
	} else {
		goto L50
	}
L40:
	;
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferIO[1])) = v188
	goto L40
L42:
	;
	if int32(999) < v171 {
		goto L40
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if v171 < int32(11) {
		goto L40
	} else {
		goto L49
	}
L45:
	;
	v178 = int32(900)
	if v178 <= v171 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v181 = v178
	goto L48
L47:
	;
	v181 = v171
	goto L48
L48:
	;
	v188 = v181 + int32(100)
	goto L41
L49:
	;
	v188 = v171 - int32(1)
	goto L41
L50:
	;
	m.G0 = v8 + int32(80)
	return
}
func F_ResOwnerReleaseCatCacheList(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v4 = int32(1)
	v5 = v3 - v4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v5
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if base.B2i32(v7 != v4)|v5 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		F_CatCacheRemoveCList(m, v13, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_ResOwnerReleaseDSM(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	F_dsm_detach(m, l0)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_ResOwnerReleaseFile(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseFile[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v3+l0*int32(48))+8)) = int32(0)
	F_FileClose(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
func F_ResOwnerReleaseSnapshot(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_UnregisterSnapshotNoOwner(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
