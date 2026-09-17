package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SPI_datumTransfer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_datumTransfer[0]))
	if v6 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_SPI_datumTransfer_0), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_SPI_datumTransfer_1), int32(1367), int32(_a_F_SPI_datumTransfer_2))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v24 = int32(_a_F_SPI_datumTransfer_3)
		v25 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_datumTransfer[1]))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
		*(*int32)(unsafe.Add(mBase, _c_F_SPI_datumTransfer[1])) = v27
		v30 = F_datumTransfer(m, l0, int32(0), l1)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_SPI_datumTransfer[1])) = v25
			return v30
		}
	}
}
func F_SPI_execute_plan_with_paramlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v3 = l2
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(-6)
	if base.B2i32(l0 == v5)|base.B2i32(l3 < v5) != 0 {
		v63 = v12
		m.G0 = v10 + int32(32)
		return v63
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v18 != int32(569278163) {
			v63 = v12
			m.G0 = v10 + int32(32)
			return v63
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_plan_with_paramlist[0]))
			if v22 == int32(0) {
				v63 = int32(-4)
				m.G0 = v10 + int32(32)
				return v63
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_plan_with_paramlist[1]))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_plan_with_paramlist[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v28
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
				*(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_plan_with_paramlist[2])) = v33
				v35 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = base.I64_extend_i32_u(l3)
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v3)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
				v45 = int32(0)
				v48 = F__SPI_execute_plan(m, l0, v10+int32(8), v45, v45, int32(1))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_plan_with_paramlist[0]))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
					*(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_plan_with_paramlist[2])) = v55
					*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = int32(0)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
					F_MemoryContextReset(m, v59)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v63 = v48
						m.G0 = v10 + int32(32)
						return v63
					}
				}
			}
		}
	}
}
func F__SPI_commit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int64
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(192)
	m.G0 = v10
	v15 = v2
	v16 = v2
	v17 = v2
	v18 = v2
	v19 = int32(-1)
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v19 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v235 = int32(m.ExcTag)
	v236 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v235 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[0]))
	v25 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[1]))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+40)))
	if v26 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v119 = v15
	v120 = v16
	v121 = v17
	v122 = v18
	goto L9
L9:
	;
	if v122 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	v61 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[2]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
	goto L17
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errcode(m, int32(1282))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errmsg(m, int32(_a_F__SPI_commit_0), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errfinish(m, int32(_a_F__SPI_commit_1), int32(241), int32(_a_F__SPI_commit_2))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	goto L3
L17:
	;
	if int32(1) < v62 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if l0 != 0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errcode(m, int32(1282))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errmsg(m, int32(_a_F__SPI_commit_3), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	F_errfinish(m, int32(_a_F__SPI_commit_1), int32(256), int32(_a_F__SPI_commit_2))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	goto L3
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v23
	v97 = v10 + int32(172)
	v99 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v99
	v102 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__SPI_commit[4])))
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)) = uint8(v102)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__SPI_commit[5])))
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)) = uint8(v105)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[6]))
	v111 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[7]))
	goto L29
L28:
	;
	goto L27
L29:
	;
	v113 = v10 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v10 + int32(12)
	goto L32
L30:
	;
	v119 = v23
	v120 = v111
	v121 = v109
	v122 = int32(0)
	goto L9
L32:
	;
	goto L30
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[7])) = v10 + int32(16)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[1]))
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v130)+41)) = uint8(v131)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_HoldPinnedPortals(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L6
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[6])) = v121
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[7])) = v120
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[0])) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	v189 = F_CopyErrorData(m)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L44
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_ForgetPortalSnapshots(m)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_CommitTransactionCommand(m)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_StartTransactionCommand(m)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	if l0 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	v157 = v10 + int32(172)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[3])) = v159
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+4)))
	*(*uint8)(unsafe.Add(mBase, _c_F__SPI_commit[4])) = uint8(v162)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+5)))
	*(*uint8)(unsafe.Add(mBase, _c_F__SPI_commit[5])) = uint8(v165)
	goto L43
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[0])) = v119
	v170 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[1]))
	v171 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v170)+41)) = uint8(v171)
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[6])) = v121
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[7])) = v120
	m.G0 = v10 + int32(192)
	return
L43:
	;
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_FlushErrorState(m)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_StartTransactionCommand(m)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	if l0 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	v210 = v10 + int32(172)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[3])) = v212
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+4)))
	*(*uint8)(unsafe.Add(mBase, _c_F__SPI_commit[4])) = uint8(v215)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+5)))
	*(*uint8)(unsafe.Add(mBase, _c_F__SPI_commit[5])) = uint8(v218)
	goto L51
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[0])) = v119
	v223 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_commit[1]))
	v224 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+41)) = uint8(v224)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v119
	F_ReThrowError(m, v189)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	goto L5
L53:
	;
	v240 = int32(v236)
	m.G0 = v10
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	if v10+int32(12) == v246 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	m.ExcPending = 1
	goto L62
L55:
	;
	if v250 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	v250 = v248
	goto L58
L57:
	;
	v250 = int32(0)
	goto L58
L58:
	;
	goto L55
L59:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v10)+188))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v10)+184))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v10)+180))
	v15 = v251
	v16 = v252
	v17 = v253
	v18 = v242
	v19 = v250
	goto L1
L60:
	;
	goto L61
L61:
	;
	F___wasm_longjmp(m, v243, v242)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	return
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__SPI_make_plan_non_temp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	v10 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_make_plan_non_temp[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v16 = F_AllocSetContextCreateInternal(m, v11, int32(_a_F__SPI_make_plan_non_temp_0), int32(0), int32(1024), int32(_a_F__SPI_make_plan_non_temp_1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = int32(_a_F__SPI_make_plan_non_temp_2)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_make_plan_non_temp[1]))
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_make_plan_non_temp[1])) = v16
	v25 = F_palloc0(m, int32(40))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(569278163)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v34
	if int32(0) < v34 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v58 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v40 = F_palloc(m, v34<<(uint(int32(2))%32))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = int32(0)
	goto L4
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v40
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v45 = v43 << (uint(int32(2)) % 32)
	if v45 == int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	base.MemoryCopy(m, v40, v48, v45)
	goto L4
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_make_plan_non_temp[1])) = v21
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	return v25
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v61 <= int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v68 = int32(0)
	goto L13
L13:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v68<<(uint(int32(2))%32))))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+94)))
	if v78 != int32(1) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L10
L15:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v178 = F_lappend(m, v177, v77)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L64
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L61
	}
L17:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+92)))
	if v81 == int32(1) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L58
	}
L20:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v77)+56))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	if v88 != v11 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v77)+88))
	if v117 != 0 {
		goto L38
	} else {
		goto L39
	}
L22:
	;
	if v88 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	if v11 != 0 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v84)+28))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84)+24))
	if v93 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v92 == int32(0) {
		goto L25
	} else {
		goto L31
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+28)) = v92
	goto L27
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = v92
	goto L27
L31:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v84)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = v98
	goto L25
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+16)) = v11
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v84)+28)) = v105
	if v105 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v84)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+16)) = int32(0)
	goto L24
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v84
	goto L37
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v84
	goto L21
L38:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+32))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	if v122 != v11 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L40
L40:
	;
	goto L15
L41:
	;
	goto L40
L42:
	;
	if v122 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	goto L41
L45:
	;
	if v11 != 0 {
		goto L52
	} else {
		goto L53
	}
L46:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v118)+28))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v118)+24))
	if v127 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v126 == int32(0) {
		goto L45
	} else {
		goto L51
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+28)) = v126
	goto L47
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = v126
	goto L47
L51:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v118)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+24)) = v132
	goto L45
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v118)+16)) = v11
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+28)) = v139
	if v139 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v118)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v118)+16)) = int32(0)
	goto L44
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v118
	goto L57
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v118
	goto L41
L58:
	;
	F_errmsg_internal(m, int32(_a_F__SPI_make_plan_non_temp_3), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F__SPI_make_plan_non_temp_4), int32(1619), int32(_a_F__SPI_make_plan_non_temp_5))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errmsg_internal(m, int32(_a_F__SPI_make_plan_non_temp_6), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F__SPI_make_plan_non_temp_4), int32(1621), int32(_a_F__SPI_make_plan_non_temp_5))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v178
	v182 = v68 + int32(1)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v182 < v183 {
		v68 = v182
		goto L13
	} else {
		goto L65
	}
L65:
	;
	goto L14
}
