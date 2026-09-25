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
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
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
	v74 = int32(0)
	v77 = base.AtomicRmwOr32(m, v74, int32(_a_F_ResOwnerReleaseBufferIO_5), v74)
	if v46&int32(16777216) == v74 {
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
	v143 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v143
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v147 = int32(_a_F_ResOwnerReleaseBufferIO_3)
	v149 = base.AtomicRmwOr32(m, v26, v143, v147)
	if v149&v147 != 0 {
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
	v94 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	if v94 == int32(0) {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(_a_F_ResOwnerReleaseBufferIO_6))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(48))))
	v105 = v8 + int32(8)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(60))))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-64))))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(56))))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(52))))
	F_GetRelationPath(m, v105, v108, v111, v114, int32(-1), v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v105
	F_errmsg(m, int32(_a_F_ResOwnerReleaseBufferIO_7), v8)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	F_errdetail(m, int32(_a_F_ResOwnerReleaseBufferIO_8), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_ResOwnerReleaseBufferIO_2), int32(_a_F_ResOwnerReleaseBufferIO_9), int32(_a_F_ResOwnerReleaseBufferIO_10))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
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
	v166 = v149
	goto L34
L34:
	;
	v174 = int32(_a_F_ResOwnerReleaseBufferIO_4)
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferIO[1]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))+8))
	if v177 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	F_perform_spin_delay(m, v8+int32(8))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L37
	}
L36:
	;
	v166 = v163
	goto L34
L37:
	;
	v161 = int32(_a_F_ResOwnerReleaseBufferIO_3)
	v163 = base.AtomicRmwOr32(m, v26, int32(0), v161)
	if v163&v161 != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v194 = int32(0)
	v197 = base.AtomicRmwOr32(m, v194, int32(_a_F_ResOwnerReleaseBufferIO_5), v194)
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v166&int32(-205520897) | int32(134217728)
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferIO[2]))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v24-int32(44))))
	F_ConditionVariableBroadcast(m, v204+v207<<(uint(int32(4))%32))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L6
	} else {
		goto L50
	}
L40:
	;
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferIO[1])) = v192
	goto L40
L42:
	;
	if int32(999) < v175 {
		goto L40
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if v175 < int32(11) {
		goto L40
	} else {
		goto L49
	}
L45:
	;
	v182 = int32(900)
	if v182 <= v175 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v185 = v182
	goto L48
L47:
	;
	v185 = v175
	goto L48
L48:
	;
	v192 = v185 + int32(100)
	goto L41
L49:
	;
	v192 = v175 - int32(1)
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
