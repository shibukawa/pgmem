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
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v15 | v16<<(uint(int32(16))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
	v30 = F_psprintf(m, int32(467904), v10)
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
	v10 = F_psprintf(m, int32(41194), v5)
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
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
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = int32(229240)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(495087)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = int64(0)
	v23 = v10 + l0<<(uint(int32(6))%32)
	v25 = v23 - int32(40)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v26 | v27
	if v26&v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L4
L2:
	;
	v48 = v26
	goto L3
L3:
	;
	v53 = int32(4122060)
	v54 = *(*int32)(unsafe.Add(mBase, _consts[423]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(16))+8))
	if v56 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	F_perform_spin_delay(m, v7+int32(16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v48 = v40
	goto L3
L6:
	;
	return
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v41 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v40 | v41
	if v40&v41 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	if v48&int32(16777216) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L10:
	;
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[423])) = v71
	goto L10
L12:
	;
	if int32(999) < v54 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v54 < int32(11) {
		goto L10
	} else {
		goto L19
	}
L15:
	;
	v61 = int32(900)
	if v61 <= v54 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v64 = v61
	goto L18
L17:
	;
	v64 = v54
	goto L18
L18:
	;
	v71 = v64 + int32(100)
	goto L11
L19:
	;
	v71 = v54 - int32(1)
	goto L11
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+108)) = int32(229240)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+104)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = int32(495087)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+88)) = int64(0)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v144 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v143 | v144
	if v143&v144 != 0 {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v48 & int32(-20971521)
	goto L20
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v48 & int32(-4194305)
	if v48&int32(134217728) == int32(0) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v89 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	if v89 == int32(0) {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(786949))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(48))))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(60))))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(-64))))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(56))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(52))))
	F_GetRelationPath(m, v7+int32(16), v103, v106, v109, int32(-1), v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v7 + int32(16)
	F_errmsg(m, int32(186514), v7)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	F_errdetail(m, int32(580189), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(495087), int32(6182), int32(527038))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
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
	v165 = v143
	goto L34
L34:
	;
	v170 = int32(4122060)
	v171 = *(*int32)(unsafe.Add(mBase, _consts[423]))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(88))+8))
	if v173 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	F_perform_spin_delay(m, v7+int32(88))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L37
	}
L36:
	;
	v165 = v157
	goto L34
L37:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v158 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v157 | v158
	if v157&v158 != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v165&int32(-205520897) | int32(134217728)
	v196 = *(*int32)(unsafe.Add(mBase, _consts[914]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(44))))
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
	*(*int32)(unsafe.Add(mBase, _consts[423])) = v188
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
	m.G0 = v7 + int32(112)
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v4 = int32(1)
	v5 = v3 - v4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v5
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v7 != v4 {
		return
	} else {
		if v5 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			F_CatCacheRemoveCList(m, v10, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		}
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[953]))
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
