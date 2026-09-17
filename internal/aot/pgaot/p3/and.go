package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckAndReportConflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(48)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v22 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 + int32(48)
	return
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if int32(0) < v25 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v29 = l5 + int32(28)
	v41 = v7
	v46 = v7
	goto L6
L4:
	;
	v269 = v7
	goto L5
L5:
	;
	if v269 == int32(0) {
		goto L1
	} else {
		goto L58
	}
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v46<<(uint(int32(2))%32))))
	v52 = int32(0)
	if l3 == v52 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v269 = v248
	goto L5
L8:
	;
	v255 = v46 + int32(1)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v255 < v256 {
		v41 = v248
		v46 = v255
		goto L6
	} else {
		goto L57
	}
L9:
	;
	if v90 == int32(0) {
		v248 = v41
		goto L8
	} else {
		goto L22
	}
L10:
	;
	v90 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v58 <= int32(0) {
		v84 = v52
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v90 = v84
	goto L9
L14:
	;
	v61 = int32(0)
	if v61 < v58 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v64 = v58
	goto L17
L16:
	;
	v64 = v61
	goto L17
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v67 = int32(0)
	goto L18
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v65+v67<<(uint(int32(2))%32))))
	v76 = base.B2i32(v75 == v51)
	if v75 == v51 {
		v84 = v76
		goto L13
	} else {
		goto L20
	}
L19:
	;
	v84 = v76
	goto L13
L20:
	;
	v78 = v67 + int32(1)
	if v78 != v64 {
		v67 = v78
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v94 = int32(0)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v94 < v95 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v106 = v94
	v108 = v95
	goto L26
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v51
	v156 = F_list_make1_impl(m, int32(472), v20+int32(12))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L31
	} else {
		goto L34
	}
L26:
	;
	v116 = v106 << (uint(int32(2)) % 32)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116+v117)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+56))
	if v120 == v51 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L25
L28:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122+v116)))
	F_BuildSpeculativeIndexInfo(m, v119, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v128 = v108
	goto L30
L30:
	;
	v130 = v106 + int32(1)
	if v130 < v128 {
		v106 = v130
		v108 = v128
		goto L26
	} else {
		goto L33
	}
L31:
	;
	return
L32:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v128 = v127
	goto L30
L33:
	;
	goto L27
L34:
	;
	v158 = F_ExecCheckIndexConstraints(m, l0, l5, l1, v20+int32(42), v29, v156)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	if v158 != 0 {
		v248 = v41
		goto L8
	} else {
		goto L36
	}
L36:
	;
	goto L38
L37:
	;
	v223 = F_palloc0(m, int32(24))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L31
	} else {
		goto L54
	}
L38:
	;
	v178 = F_table_slot_create(m, v93, int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L31
	} else {
		goto L40
	}
L39:
	;
	if v178 == int32(0) {
		v248 = v41
		goto L8
	} else {
		goto L52
	}
L40:
	;
	v180 = F_GetLatestSnapshot(m)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L31
	} else {
		goto L41
	}
L41:
	;
	F_PushActiveSnapshot(m, v180)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	v185 = v20 + int32(42)
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_CheckAndReportConflict[0]))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	goto L43
L43:
	;
	v190 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L31
	} else {
		goto L44
	}
L44:
	;
	v193 = int32(0)
	v196 = v20 + int32(20)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v93)+188))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+104))
	v199 = m.T0[v198].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v93, v185, v188, v178, v190, int32(1), v193, v193, v196)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L31
	} else {
		goto L45
	}
L45:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L31
	} else {
		goto L46
	}
L46:
	;
	v203 = F_should_refetch_tuple(m, v199, v196)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L31
	} else {
		goto L47
	}
L47:
	;
	if v203 == int32(0) {
		goto L37
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v51
	v212 = F_list_make1_impl(m, int32(472), v20+int32(8))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L31
	} else {
		goto L49
	}
L49:
	;
	v214 = F_ExecCheckIndexConstraints(m, l0, l5, l1, v185, v29, v212)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L31
	} else {
		goto L50
	}
L50:
	;
	if v214 == int32(0) {
		goto L38
	} else {
		goto L51
	}
L51:
	;
	goto L39
L52:
	;
	F_ExecDropSingleTupleTableSlot(m, v178)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L31
	} else {
		goto L53
	}
L53:
	;
	v248 = v41
	goto L8
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+4)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v178
	v233 = F_GetTupleTransactionInfo(m, v178, v223+int32(8), v223+int32(12), v223+int32(16))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L31
	} else {
		goto L55
	}
L55:
	;
	v235 = F_lappend(m, v41, v223)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L31
	} else {
		goto L56
	}
L56:
	;
	v248 = v235
	goto L8
L57:
	;
	goto L7
L58:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if int32(1) < v279 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v282 = int32(6)
	goto L61
L60:
	;
	v282 = l2
	goto L61
L61:
	;
	F_ReportApplyConflict(m, l1, l0, int32(21), v282, l4, l5, v269)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L31
	} else {
		goto L62
	}
L62:
	;
	goto L1
}
func F_RollbackAndReleaseCurrentSubTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_RollbackAndReleaseCurrentSubTransaction[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if base.Ui32(int32(19)) < base.Ui32(v10) {
		F_CleanupSubTransaction(m)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	} else {
		if v10 != int32(12) {
			if int32(1)<<(uint(v10)%32)&int32(_a_F_RollbackAndReleaseCurrentSubTransaction_0) == int32(0) {
				F_CleanupSubTransaction(m)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			} else {
				F_errstart_cold(m, int32(22), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					if base.Ui32(v25) <= base.Ui32(int32(19)) {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v25<<(uint(int32(2))%32))+uint32(_c_F_RollbackAndReleaseCurrentSubTransaction[1])))
						v32 = v30
					} else {
						v32 = int32(_a_F_RollbackAndReleaseCurrentSubTransaction_1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v32
					F_errmsg_internal(m, int32(_a_F_RollbackAndReleaseCurrentSubTransaction_2), v6)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_RollbackAndReleaseCurrentSubTransaction_3), int32(_a_F_RollbackAndReleaseCurrentSubTransaction_4), int32(_a_F_RollbackAndReleaseCurrentSubTransaction_5))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			F_AbortSubTransaction(m)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				F_CleanupSubTransaction(m)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			}
		}
	}
}
