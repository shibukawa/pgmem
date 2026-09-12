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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
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
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v271 int32
	_ = v271
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
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
	if v25 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v271 == int32(0) {
		goto L1
	} else {
		goto L59
	}
L4:
	;
	v271 = v7
	goto L3
L5:
	;
	goto L6
L6:
	;
	v29 = l5 + int32(28)
	v39 = v7
	v42 = v7
	goto L7
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v42<<(uint(int32(2))%32))))
	v52 = int32(0)
	if l3 == v52 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v271 = v250
	goto L3
L9:
	;
	v259 = v42 + int32(1)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v259 < v260 {
		v39 = v250
		v42 = v259
		goto L7
	} else {
		goto L58
	}
L10:
	;
	if v90 == int32(0) {
		v250 = v39
		goto L9
	} else {
		goto L23
	}
L11:
	;
	v90 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v58 <= int32(0) {
		v83 = v52
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v90 = v83
	goto L10
L15:
	;
	v61 = int32(0)
	if v61 < v58 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v64 = v58
	goto L18
L17:
	;
	v64 = v61
	goto L18
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v67 = int32(0)
	goto L19
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v65+v67<<(uint(int32(2))%32))))
	v76 = base.B2i32(v75 == v51)
	if v75 == v51 {
		v83 = v76
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v83 = v76
	goto L14
L21:
	;
	v78 = v67 + int32(1)
	if v78 != v64 {
		v67 = v78
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v94 = int32(0)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v94 < v95 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v105 = v95
	v106 = v94
	goto L27
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v51
	v156 = F_list_make1_impl(m, int32(472), v20+int32(12))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L32
	} else {
		goto L35
	}
L27:
	;
	v116 = v106 << (uint(int32(2)) % 32)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116+v117)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+56))
	if v120 == v51 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L26
L29:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122+v116)))
	F_BuildSpeculativeIndexInfo(m, v119, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v128 = v105
	goto L31
L31:
	;
	v130 = v106 + int32(1)
	if v130 < v128 {
		v105 = v128
		v106 = v130
		goto L27
	} else {
		goto L34
	}
L32:
	;
	return
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v128 = v127
	goto L31
L34:
	;
	goto L28
L35:
	;
	v158 = F_ExecCheckIndexConstraints(m, l0, l5, l1, v20+int32(42), v29, v156)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	if v158 != 0 {
		v250 = v39
		goto L9
	} else {
		goto L37
	}
L37:
	;
	goto L39
L38:
	;
	v227 = F_palloc0(m, int32(24))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L32
	} else {
		goto L55
	}
L39:
	;
	v178 = F_table_slot_create(m, v93, int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L32
	} else {
		goto L41
	}
L40:
	;
	if v178 == int32(0) {
		v250 = v39
		goto L9
	} else {
		goto L53
	}
L41:
	;
	v180 = F_GetLatestSnapshot(m)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L32
	} else {
		goto L42
	}
L42:
	;
	F_PushActiveSnapshot(m, v180)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L32
	} else {
		goto L43
	}
L43:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _consts[104]))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	goto L44
L44:
	;
	v190 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L32
	} else {
		goto L45
	}
L45:
	;
	v193 = int32(0)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v93)+188))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+104))
	v199 = m.T0[v198].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v93, v20+int32(42), v188, v178, v190, int32(1), v193, v193, v20+int32(20))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L32
	} else {
		goto L46
	}
L46:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L32
	} else {
		goto L47
	}
L47:
	;
	v205 = F_should_refetch_tuple(m, v199, v20+int32(20))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L32
	} else {
		goto L48
	}
L48:
	;
	if v205 == int32(0) {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v51
	v216 = F_list_make1_impl(m, int32(472), v20+int32(8))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L32
	} else {
		goto L50
	}
L50:
	;
	v218 = F_ExecCheckIndexConstraints(m, l0, l5, l1, v20+int32(42), v29, v216)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L32
	} else {
		goto L51
	}
L51:
	;
	if v218 == int32(0) {
		goto L39
	} else {
		goto L52
	}
L52:
	;
	goto L40
L53:
	;
	F_ExecDropSingleTupleTableSlot(m, v178)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L32
	} else {
		goto L54
	}
L54:
	;
	v250 = v39
	goto L9
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227)+4)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v227))) = v178
	v237 = F_GetTupleTransactionInfo(m, v178, v227+int32(8), v227+int32(12), v227+int32(16))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L32
	} else {
		goto L56
	}
L56:
	;
	v239 = F_lappend(m, v39, v227)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L32
	} else {
		goto L57
	}
L57:
	;
	v250 = v239
	goto L9
L58:
	;
	goto L8
L59:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	if int32(1) < v283 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v286 = int32(6)
	goto L62
L61:
	;
	v286 = l2
	goto L62
L62:
	;
	F_ReportApplyConflict(m, l1, l0, int32(21), v286, l4, l5, v271)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L32
	} else {
		goto L63
	}
L63:
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if base.Ui32(int32(19)) < base.Ui32(v10) {
		F_CleanupSubTransaction(m)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	} else {
		if v10 != int32(12) {
			if int32(1)<<(uint(v10)%32)&int32(1011711) == int32(0) {
				F_CleanupSubTransaction(m)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
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
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v25<<(uint(int32(2))%32))+uint32(_consts[175])))
						v35 = v34
					} else {
						v35 = int32(531628)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v35
					F_errmsg_internal(m, int32(183340), v6)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						F_errfinish(m, int32(482378), int32(4833), int32(252684))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
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
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				F_CleanupSubTransaction(m)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			}
		}
	}
}
