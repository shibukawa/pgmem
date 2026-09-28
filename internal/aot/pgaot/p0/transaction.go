package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BeginTransactionBlock(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_BeginTransactionBlock[0]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	switch v9 {
	case 0, 2, 6, 8, 9, 10, 11, 13, 14, 16, 17, 18, 19:
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
			if base.Ui32(v34) <= base.Ui32(int32(19)) {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v34<<(uint(int32(2))%32))+uint32(_c_F_BeginTransactionBlock[1])))
				v41 = v39
			} else {
				v41 = int32(_a_F_BeginTransactionBlock_0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v41
			F_errmsg_internal(m, int32(_a_F_BeginTransactionBlock_1), v5)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_BeginTransactionBlock_2), int32(4028), int32(_a_F_BeginTransactionBlock_3))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 1:
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(2)
		m.G0 = v5 + int32(16)
		return
	case 3, 5, 7, 12, 15:
		v14 = F_errstart(m, int32(19), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			if v14 == int32(0) {
				m.G0 = v5 + int32(16)
				return
			} else {
				F_errcode(m, int32(16777538))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_BeginTransactionBlock_4), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_BeginTransactionBlock_2), int32(4010), int32(_a_F_BeginTransactionBlock_3))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							m.G0 = v5 + int32(16)
							return
						}
					}
				}
			}
		}
	case 4:
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(2)
		m.G0 = v5 + int32(16)
		return
	default:
		m.G0 = v5 + int32(16)
		return
	}
}
func F_CommitTransactionCommand(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int64
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	goto L2
L1:
	;
	m.G0 = v10 + int32(32)
	return
L2:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[0])))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[1])))
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[2]))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[3]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	switch v27 {
	case 0, 5:
		goto L18
	case 1:
		goto L4
	case 2:
		goto L17
	case 3, 4, 12:
		goto L16
	case 6:
		goto L15
	default:
		goto L1
	case 8:
		goto L14
	case 9:
		goto L13
	case 10:
		goto L12
	case 11:
		goto L11
	case 13:
		goto L10
	case 14:
		goto L9
	case 16:
		goto L5
	case 17:
		goto L6
	case 18:
		goto L8
	case 19:
		goto L7
	}
L3:
	;
	F_CommitTransaction(m)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L19
	} else {
		goto L71
	}
L4:
	;
	goto L3
L5:
	;
	F_CleanupSubTransaction(m)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L19
	} else {
		goto L70
	}
L6:
	;
	F_AbortSubTransaction(m)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L19
	} else {
		goto L69
	}
L7:
	;
	v213 = *(*int64)(unsafe.Add(mBase, uint32(v26)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = int32(0)
	F_CleanupSubTransaction(m)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L19
	} else {
		goto L66
	}
L8:
	;
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v26)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = int32(0)
	F_AbortSubTransaction(m)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L19
	} else {
		goto L62
	}
L9:
	;
	goto L44
L10:
	;
	goto L40
L11:
	;
	F_StartSubTransaction(m)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L19
	} else {
		goto L39
	}
L12:
	;
	F_PrepareTransaction(m)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L19
	} else {
		goto L38
	}
L13:
	;
	F_AbortTransaction(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L19
	} else {
		goto L34
	}
L14:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L31
	}
L15:
	;
	F_CommitTransaction(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L19
	} else {
		goto L28
	}
L16:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L19
	} else {
		goto L27
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(3)
	goto L1
L18:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if base.Ui32(v32) <= base.Ui32(int32(19)) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v39
	F_errmsg_internal(m, int32(_a_F_CommitTransactionCommand_0), v10)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L19
	} else {
		goto L25
	}
L22:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_c_F_CommitTransactionCommand[4])))
	v39 = v37
	goto L24
L23:
	;
	v39 = int32(_a_F_CommitTransactionCommand_1)
	goto L24
L24:
	;
	goto L21
L25:
	;
	F_errfinish(m, int32(_a_F_CommitTransactionCommand_2), int32(3247), int32(_a_F_CommitTransactionCommand_3))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	goto L1
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(0)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+77)))
	if v57 != int32(1) {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_StartTransaction(m)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	v62 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+77)) = uint8(v62)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(3)
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[1])) = uint8(v22)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[2])) = v24
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[0])) = uint8(v20)
	goto L1
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(0)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+77)))
	if v76 != int32(1) {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_StartTransaction(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	v81 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+77)) = uint8(v81)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(3)
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[1])) = uint8(v22)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[2])) = v24
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[0])) = uint8(v20)
	goto L1
L34:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(0)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+77)))
	if v97 != int32(1) {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_StartTransaction(m)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	v102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+77)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(3)
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[1])) = uint8(v22)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[2])) = v24
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[0])) = uint8(v20)
	goto L1
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(0)
	goto L1
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(12)
	goto L1
L40:
	;
	F_CommitSubTransaction(m)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L19
	} else {
		goto L42
	}
L41:
	;
	goto L1
L42:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[3]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+24))
	if v131 == int32(13) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	F_CommitSubTransaction(m)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L19
	} else {
		goto L46
	}
L45:
	;
	switch v145 - int32(6) {
	case 0:
		goto L50
	default:
		goto L48
	case 4:
		goto L49
	}
L46:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[3]))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+24))
	if v145 == int32(14) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L19
	} else {
		goto L55
	}
L49:
	;
	F_PrepareTransaction(m)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L19
	} else {
		goto L54
	}
L50:
	;
	F_CommitTransaction(m)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L19
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+24)) = int32(0)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+77)))
	if v154 != int32(1) {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_StartTransaction(m)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L19
	} else {
		goto L53
	}
L53:
	;
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+77)) = uint8(v159)
	*(*int32)(unsafe.Add(mBase, uint32(v144)+24)) = int32(3)
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[1])) = uint8(v22)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[2])) = v24
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[0])) = uint8(v20)
	goto L1
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+24)) = int32(0)
	goto L1
L55:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v144)+24))
	if base.Ui32(v177) <= base.Ui32(int32(19)) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v184
	F_errmsg_internal(m, int32(_a_F_CommitTransactionCommand_0), v10+int32(16))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L19
	} else {
		goto L60
	}
L57:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v177<<(uint(int32(2))%32))+uint32(_c_F_CommitTransactionCommand[4])))
	v184 = v182
	goto L59
L58:
	;
	v184 = int32(_a_F_CommitTransactionCommand_1)
	goto L59
L59:
	;
	goto L56
L60:
	;
	F_errfinish(m, int32(_a_F_CommitTransactionCommand_2), int32(3413), int32(_a_F_CommitTransactionCommand_3))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L19
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_CleanupSubTransaction(m)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L19
	} else {
		goto L63
	}
L63:
	;
	F_DefineSavepoint(m, int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L19
	} else {
		goto L64
	}
L64:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v207)+12)) = v196
	F_StartSubTransaction(m)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L19
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+24)) = int32(12)
	goto L1
L66:
	;
	F_DefineSavepoint(m, int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L19
	} else {
		goto L67
	}
L67:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransactionCommand[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v222)+12)) = v213
	F_StartSubTransaction(m)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L19
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+24)) = int32(12)
	goto L1
L69:
	;
	goto L5
L70:
	;
	goto L2
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(0)
	goto L1
}
func F_GetTransactionSnapshot(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int64
	_ = v286
	var v288 int64
	_ = v288
	var v290 int64
	_ = v290
	var v292 int64
	_ = v292
	var v294 int64
	_ = v294
	var v296 int64
	_ = v296
	var v298 int64
	_ = v298
	var v300 int64
	_ = v300
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[0]))
	if v7 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L10
	} else {
		goto L122
	}
L2:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[1])))
	if v11 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v437 = v7
	goto L4
L4:
	;
	return v437
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[2]))
	if v15 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[3]))
	if int32(2) <= v376 {
		goto L105
	} else {
		goto L106
	}
L8:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[4]))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+72))
	if v69 != 0 {
		goto L23
	} else {
		goto L24
	}
L9:
	;
	F_pairingheap_remove(m, int32(_a_F_GetTransactionSnapshot_0), v15+int32(52))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[2])) = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[5]))
	if v29 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[6]))
	if v31 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v35 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[7])) = v35
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v35
	goto L8
L14:
	;
	goto L15
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[8]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+52))
	v44 = int32(3)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[6]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47-int32(48))))
	if base.B2i32(base.Ui32(v43) < base.Ui32(v44))|base.B2i32(base.Ui32(v50) < base.Ui32(v44)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[7])) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v42)+52)) = v50
	goto L8
L17:
	;
	if v43-v50 < int32(0) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if base.Ui32(v50) <= base.Ui32(v43) {
		goto L8
	} else {
		goto L21
	}
L20:
	;
	goto L8
L21:
	;
	goto L16
L22:
	;
	if v72&int32(1) != 0 {
		goto L1
	} else {
		goto L26
	}
L23:
	;
	v72 = int32(1)
	goto L25
L24:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+76)))
	v72 = v71
	goto L25
L25:
	;
	goto L22
L26:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[3]))
	if int32(2) <= v76 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v372 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[1])) = uint8(v372)
	return v366
L28:
	;
	if v76 == int32(3) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v363 = F_GetSnapshotData(m, int32(_a_F_GetTransactionSnapshot_1))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L10
	} else {
		goto L104
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[9])) = v266
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[10]))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v266)+24))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v266)+16))
	v273 = int32(2)
	v275 = int32(72)
	v280 = v271<<(uint(v273)%32) + v275
	if int32(0) < v270 {
		goto L84
	} else {
		goto L85
	}
L32:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[11])))
	if v84 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	goto L34
L34:
	;
	v259 = F_GetSnapshotData(m, int32(_a_F_GetTransactionSnapshot_1))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L10
	} else {
		goto L83
	}
L35:
	;
	v266 = v255
	goto L31
L36:
	;
	if v94 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[12]))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+308))
	v92 = base.B2i32(v90 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[11])) = uint8(v92)
	v94 = v92
	goto L39
L38:
	;
	v94 = int32(0)
	goto L39
L39:
	;
	goto L36
L40:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[13])))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[14])))
	v104 = F_GetSerializableTransactionSnapshotInt(m, int32(_a_F_GetTransactionSnapshot_1), int32(0), int32(-1))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L10
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L10
	} else {
		goto L77
	}
L43:
	;
	v106 = int32(1)
	if base.B2i32(v100&v106 == int32(0))|base.B2i32(v98 != v106) != 0 {
		v226 = v104
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v255 = v226
	goto L35
L45:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[15]))
	if v114 == int32(0) {
		v226 = v104
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v119 = v104
	goto L47
L47:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[16]))
	v127 = F_LWLockAcquire(m, v123+int32(3584), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L10
	} else {
		goto L49
	}
L48:
	;
	v226 = v220
	goto L44
L49:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[15]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+108))
	v133 = v131 | int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(v130)+108)) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+92))
	if base.B2i32(v135 == int32(0))|base.B2i32(v135 == v130+int32(88)) != 0 {
		v175 = v130
		v176 = v133
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+108)) = v176 & int32(-65)
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[16]))
	F_LWLockRelease(m, v184+int32(3584))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L10
	} else {
		goto L62
	}
L51:
	;
	v142 = v130
	goto L52
L52:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v142)+108))
	if v147&int32(256) != 0 {
		v175 = v142
		v176 = v147
		goto L50
	} else {
		goto L54
	}
L53:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v167)+108))
	v175 = v167
	v176 = v174
	goto L50
L54:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[16]))
	F_LWLockRelease(m, v151+int32(3584))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	F_ProcWaitForSignal(m, int32(134217780))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[16]))
	v164 = F_LWLockAcquire(m, v160+int32(3584), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L10
	} else {
		goto L57
	}
L57:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[15]))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+92))
	if v168 != v167+int32(88) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v173 = v168
	goto L60
L59:
	;
	v173 = int32(0)
	goto L60
L60:
	;
	if v173 != 0 {
		v142 = v167
		goto L52
	} else {
		goto L61
	}
L61:
	;
	goto L53
L62:
	;
	if v176&int32(256) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_ReleasePredicateLocks(m, int32(0), int32(1))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L10
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v199 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L10
	} else {
		goto L67
	}
L66:
	;
	v255 = v119
	goto L35
L67:
	;
	if v199 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L10
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v213 = int32(0)
	F_ReleasePredicateLocks(m, v213, v213)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L10
	} else {
		goto L74
	}
L71:
	;
	F_errmsg_internal(m, int32(_a_F_GetTransactionSnapshot_2), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_GetTransactionSnapshot_3), int32(1534), int32(_a_F_GetTransactionSnapshot_4))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L10
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	v220 = F_GetSerializableTransactionSnapshotInt(m, int32(_a_F_GetTransactionSnapshot_1), int32(0), int32(-1))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[15]))
	if v223 != 0 {
		v119 = v220
		goto L47
	} else {
		goto L76
	}
L76:
	;
	goto L48
L77:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L10
	} else {
		goto L78
	}
L78:
	;
	F_errmsg(m, int32(_a_F_GetTransactionSnapshot_5), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L10
	} else {
		goto L79
	}
L79:
	;
	v242 = F_errdetail(m, int32(_a_F_GetTransactionSnapshot_6), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L10
	} else {
		goto L80
	}
L80:
	;
	F_errhint(m, int32(_a_F_GetTransactionSnapshot_7), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L10
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_GetTransactionSnapshot_3), int32(1626), int32(_a_F_GetTransactionSnapshot_8))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L10
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	v266 = v259
	goto L31
L84:
	;
	v283 = (v270+v271)<<(uint(v273)%32) + v275
	goto L86
L85:
	;
	v283 = v280
	goto L86
L86:
	;
	v284 = F_MemoryContextAlloc(m, v269, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L10
	} else {
		goto L87
	}
L87:
	;
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v266)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v284)+48)) = v286
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v266)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v284)+40)) = v288
	v290 = *(*int64)(unsafe.Add(mBase, uint32(v266)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v284)+24)) = v290
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v266)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v284)+56)) = v292
	v294 = *(*int64)(unsafe.Add(mBase, uint32(v266)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v284)+32)) = v294
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v266)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v284)+16)) = v296
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v266)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v284)+8)) = v298
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v266)))
	*(*int64)(unsafe.Add(mBase, uint32(v284))) = v300
	*(*int64)(unsafe.Add(mBase, uint32(v284)+64)) = int64(0)
	v304 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v284)+48)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v284)+44)) = v304
	v308 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v284)+30)) = uint8(v308)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v266)+16))
	if v310 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v266)+24))
	if v326 <= int32(0) {
		goto L94
	} else {
		goto L95
	}
L89:
	;
	v312 = v284 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v284)+12)) = v312
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v266)+16))
	v316 = v314 << (uint(int32(2)) % 32)
	if v316 == int32(0) {
		goto L88
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+12)) = int32(0)
	goto L88
L92:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	base.MemoryCopy(m, v312, v319, v316)
	goto L88
L93:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[17])) = v284
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[9])) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v284)+48)) = v347
	F_pairingheap_add(m, int32(_a_F_GetTransactionSnapshot_0), v284+int32(52))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L10
	} else {
		goto L103
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+20)) = int32(0)
	v347 = int32(1)
	goto L93
L95:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+28)))
	if v329 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+29)))
	if v332 != int32(1) {
		goto L94
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v335 = v284 + v280
	*(*int32)(unsafe.Add(mBase, uint32(v284)+20)) = v335
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v266)+24))
	v339 = v337 << (uint(int32(2)) % 32)
	if v339 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L98
L100:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	base.MemoryCopy(m, v335, v340, v339)
	goto L102
L101:
	;
	goto L102
L102:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v284)+48))
	v347 = v342 + int32(1)
	goto L93
L103:
	;
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[9]))
	v366 = v360
	goto L27
L104:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[9])) = v363
	v366 = v363
	goto L27
L105:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[9]))
	return v380
L106:
	;
	goto L107
L107:
	;
	v383 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[2]))
	if v383 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v434 = F_GetSnapshotData(m, int32(_a_F_GetTransactionSnapshot_1))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L10
	} else {
		goto L121
	}
L109:
	;
	F_pairingheap_remove(m, int32(_a_F_GetTransactionSnapshot_0), v383+int32(52))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L10
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[2])) = int32(0)
	v395 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[5]))
	if v395 != 0 {
		goto L108
	} else {
		goto L111
	}
L111:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[6]))
	if v397 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v401 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[7])) = v401
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v404)+52)) = v401
	goto L108
L113:
	;
	goto L114
L114:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[8]))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)+52))
	v410 = int32(3)
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[6]))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v413-int32(48))))
	if base.B2i32(base.Ui32(v409) < base.Ui32(v410))|base.B2i32(base.Ui32(v416) < base.Ui32(v410)) == int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[7])) = v416
	*(*int32)(unsafe.Add(mBase, uint32(v408)+52)) = v416
	goto L108
L116:
	;
	if v409-v416 < int32(0) {
		goto L115
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	if base.Ui32(v416) <= base.Ui32(v409) {
		goto L108
	} else {
		goto L120
	}
L119:
	;
	goto L108
L120:
	;
	goto L115
L121:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[9])) = v434
	v437 = v434
	goto L4
L122:
	;
	F_errmsg_internal(m, int32(_a_F_GetTransactionSnapshot_9), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L10
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_GetTransactionSnapshot_10), int32(307), int32(_a_F_GetTransactionSnapshot_11))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L10
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_IsTransactionOrTransactionBlock(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_IsTransactionOrTransactionBlock[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+24))
	return base.B2i32(v3 != int32(0))
}
func F_PushTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_PushTransaction[0]))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_PushTransaction[1]))
	v11 = F_MemoryContextAllocZero(m, v9, int32(88))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = int32(1)
		v14 = int32(_a_F_PushTransaction_0)
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_PushTransaction[2]))
		v18 = v16 + v13
		*(*int32)(unsafe.Add(mBase, _c_F_PushTransaction[2])) = v18
		if v18 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v7
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(0)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
			v25 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v24 + v25
			v29 = int32(_a_F_PushTransaction_1)
			v31 = *(*int32)(unsafe.Add(mBase, _c_F_PushTransaction[3]))
			v33 = v31 + v25
			*(*int32)(unsafe.Add(mBase, _c_F_PushTransaction[3])) = v33
			*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v33
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v11)+20)) = int64(47244640256)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v36
			v45 = *(*int32)(unsafe.Add(mBase, _c_F_PushTransaction[4]))
			*(*int32)(unsafe.Add(mBase, uint32(v11+int32(60)))) = v45
			v48 = *(*int32)(unsafe.Add(mBase, _c_F_PushTransaction[5]))
			*(*int32)(unsafe.Add(mBase, uint32(v11-int32(-64)))) = v48
			v51 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PushTransaction[6])))
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+68)) = uint8(v51)
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+69)))
			v54 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v54
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+69)) = uint8(v53)
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v7)+72))
			if v57 == v54 {
				v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+76)))
				v61 = v60
			} else {
				v61 = v13
			}
			v62 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+78)) = uint8(v62)
			v65 = v61 & int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+76)) = uint8(v65)
			*(*int32)(unsafe.Add(mBase, _c_F_PushTransaction[0])) = v11
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_PushTransaction[2])) = v16
			F_pfree(m, v11)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_PushTransaction_2), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_PushTransaction_3), int32(_a_F_PushTransaction_4), int32(_a_F_PushTransaction_5))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_SetTransactionSnapshot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v248 int64
	_ = v248
	var v250 int64
	_ = v250
	var v252 int64
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[0]))
	if v14 == int32(0) {
		v65 = F_GetSnapshotData(m, int32(_a_F_SetTransactionSnapshot_0))
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v65
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v68
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v70
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v72
			if v72 == int32(0) {
			} else {
				v77 = v72 << (uint(int32(2)) % 32)
				if v77 == int32(0) {
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					base.MemoryCopy(m, v80, v81, v77)
				}
			}
			v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v65)+24)) = v84
			if v84 <= int32(0) {
			} else {
				v89 = v84 << (uint(int32(2)) % 32)
				if v89 == int32(0) {
				} else {
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					base.MemoryCopy(m, v92, v93, v89)
				}
			}
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
			*(*uint8)(unsafe.Add(mBase, uint32(v65)+28)) = uint8(v96)
			v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
			*(*int64)(unsafe.Add(mBase, uint32(v65)+64)) = int64(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v65)+29)) = uint8(v98)
			v102 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
			if l3 != 0 {
				v104 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[2]))
				v108 = F_LWLockAcquire(m, v104+int32(512), int32(0))
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return
				} else {
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
					v111 = int32(2)
					v116 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					v118 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[3]))
					v124 = base.B2i32(base.Ui32(v111) < base.Ui32(v110)) & base.B2i32(base.Ui32(v111) < base.Ui32(v102)) & base.B2i32(v116 == v118) & base.B2i32(v110-v102 <= int32(0))
					if v124 != 0 {
						*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[4])) = v102
						v128 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[5]))
						*(*int32)(unsafe.Add(mBase, uint32(v128)+52)) = v102
						v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+36)))
						v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+36)))
						v136 = v130&int32(6) | v133&int32(-7)
						*(*uint8)(unsafe.Add(mBase, uint32(v128)+36)) = uint8(v136)
						v139 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[6]))
						v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
						v141 = *(*int32)(unsafe.Add(mBase, uint32(v128)+32))
						*(*uint8)(unsafe.Add(mBase, uint32(v140+v141))) = uint8(v136)
					} else {
					}
					v147 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[2]))
					F_LWLockRelease(m, v147+int32(512))
					mBase = m.M
					v151 = m.ExcPending
					if v151 != 0 {
						return
					} else {
						if v124 != 0 {
							v180 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[7]))
							if int32(2) <= v180 {
								if v180 == int32(3) {
									v186 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
									v188 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[8]))
									if v188 < int32(0) {
										v192 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[9])))
										if v192 == int32(1) {
											v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[10])))
											if v196&int32(1) != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v204 = m.ExcPending
												if v204 != 0 {
													return
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v207 = m.ExcPending
													if v207 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_SetTransactionSnapshot_1), int32(0))
														mBase = m.M
														v211 = m.ExcPending
														if v211 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_SetTransactionSnapshot_2), int32(1677), int32(_a_F_SetTransactionSnapshot_3))
															mBase = m.M
															v216 = m.ExcPending
															if v216 != 0 {
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
												v199 = F_GetSerializableTransactionSnapshotInt(m, v186, l1, l2)
												mBase = m.M
												v200 = m.ExcPending
												if v200 != 0 {
													return
												} else {
													v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
													v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
													v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
													v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
													v225 = int32(2)
													v227 = int32(72)
													v232 = v223<<(uint(v225)%32) + v227
													if int32(0) < v222 {
														v235 = (v222+v223)<<(uint(v225)%32) + v227
													} else {
														v235 = v232
													}
													v236 = F_MemoryContextAlloc(m, v219, v235)
													mBase = m.M
													v237 = m.ExcPending
													if v237 != 0 {
														return
													} else {
														v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
														v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
														v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
														v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
														v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
														v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
														v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
														v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
														*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
														*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
														v256 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
														*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
														v260 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
														v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
														if v262 != 0 {
															v264 = v236 + int32(72)
															*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
															v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
															v268 = v266 << (uint(int32(2)) % 32)
															if v268 == int32(0) {
															} else {
																v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
																base.MemoryCopy(m, v264, v271, v268)
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
														}
														v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
														if v278 <= int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
															v300 = int32(1)
														} else {
															v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
															if v281 == int32(1) {
																v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
																if v284 != int32(1) {
																	*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
																	v300 = int32(1)
																} else {
																	v287 = v236 + v232
																	*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
																	v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
																	v291 = v289 << (uint(int32(2)) % 32)
																	if v291 != 0 {
																		v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																		base.MemoryCopy(m, v287, v292, v291)
																	} else {
																	}
																	v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
																	v300 = v294 + int32(1)
																}
															} else {
																v287 = v236 + v232
																*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
																v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
																v291 = v289 << (uint(int32(2)) % 32)
																if v291 != 0 {
																	v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																	base.MemoryCopy(m, v287, v292, v291)
																} else {
																}
																v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
																v300 = v294 + int32(1)
															}
														}
														*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
														*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
														*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
														F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
														mBase = m.M
														v310 = m.ExcPending
														if v310 != 0 {
															return
														} else {
															v317 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
															m.G0 = v11 + int32(16)
															return
														}
													}
												}
											}
										} else {
											v199 = F_GetSerializableTransactionSnapshotInt(m, v186, l1, l2)
											mBase = m.M
											v200 = m.ExcPending
											if v200 != 0 {
												return
											} else {
												v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
												v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
												v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
												v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
												v225 = int32(2)
												v227 = int32(72)
												v232 = v223<<(uint(v225)%32) + v227
												if int32(0) < v222 {
													v235 = (v222+v223)<<(uint(v225)%32) + v227
												} else {
													v235 = v232
												}
												v236 = F_MemoryContextAlloc(m, v219, v235)
												mBase = m.M
												v237 = m.ExcPending
												if v237 != 0 {
													return
												} else {
													v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
													v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
													v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
													v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
													v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
													v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
													v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
													v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
													*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
													*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
													v256 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
													*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
													v260 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
													v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
													if v262 != 0 {
														v264 = v236 + int32(72)
														*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
														v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
														v268 = v266 << (uint(int32(2)) % 32)
														if v268 == int32(0) {
														} else {
															v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
															base.MemoryCopy(m, v264, v271, v268)
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
													}
													v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
													if v278 <= int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
														v300 = int32(1)
													} else {
														v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
														if v281 == int32(1) {
															v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
															if v284 != int32(1) {
																*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
																v300 = int32(1)
															} else {
																v287 = v236 + v232
																*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
																v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
																v291 = v289 << (uint(int32(2)) % 32)
																if v291 != 0 {
																	v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																	base.MemoryCopy(m, v287, v292, v291)
																} else {
																}
																v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
																v300 = v294 + int32(1)
															}
														} else {
															v287 = v236 + v232
															*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
															v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
															v291 = v289 << (uint(int32(2)) % 32)
															if v291 != 0 {
																v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																base.MemoryCopy(m, v287, v292, v291)
															} else {
															}
															v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
															v300 = v294 + int32(1)
														}
													}
													*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
													*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
													*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
													F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
													mBase = m.M
													v310 = m.ExcPending
													if v310 != 0 {
														return
													} else {
														v317 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
														m.G0 = v11 + int32(16)
														return
													}
												}
											}
										}
									} else {
										v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
										v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
										v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
										v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
										v225 = int32(2)
										v227 = int32(72)
										v232 = v223<<(uint(v225)%32) + v227
										if int32(0) < v222 {
											v235 = (v222+v223)<<(uint(v225)%32) + v227
										} else {
											v235 = v232
										}
										v236 = F_MemoryContextAlloc(m, v219, v235)
										mBase = m.M
										v237 = m.ExcPending
										if v237 != 0 {
											return
										} else {
											v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
											v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
											v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
											v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
											v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
											v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
											v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
											v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
											*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
											*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
											v256 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
											*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
											v260 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
											v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
											if v262 != 0 {
												v264 = v236 + int32(72)
												*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
												v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
												v268 = v266 << (uint(int32(2)) % 32)
												if v268 == int32(0) {
												} else {
													v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
													base.MemoryCopy(m, v264, v271, v268)
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
											}
											v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
											if v278 <= int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
												v300 = int32(1)
											} else {
												v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
												if v281 == int32(1) {
													v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
													if v284 != int32(1) {
														*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
														v300 = int32(1)
													} else {
														v287 = v236 + v232
														*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
														v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
														v291 = v289 << (uint(int32(2)) % 32)
														if v291 != 0 {
															v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
															base.MemoryCopy(m, v287, v292, v291)
														} else {
														}
														v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
														v300 = v294 + int32(1)
													}
												} else {
													v287 = v236 + v232
													*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
													v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
													v291 = v289 << (uint(int32(2)) % 32)
													if v291 != 0 {
														v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
														base.MemoryCopy(m, v287, v292, v291)
													} else {
													}
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
													v300 = v294 + int32(1)
												}
											}
											*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
											*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
											*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
											F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
											mBase = m.M
											v310 = m.ExcPending
											if v310 != 0 {
												return
											} else {
												v317 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
												m.G0 = v11 + int32(16)
												return
											}
										}
									}
								} else {
									v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
									v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
									v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
									v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
									v225 = int32(2)
									v227 = int32(72)
									v232 = v223<<(uint(v225)%32) + v227
									if int32(0) < v222 {
										v235 = (v222+v223)<<(uint(v225)%32) + v227
									} else {
										v235 = v232
									}
									v236 = F_MemoryContextAlloc(m, v219, v235)
									mBase = m.M
									v237 = m.ExcPending
									if v237 != 0 {
										return
									} else {
										v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
										v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
										v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
										v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
										v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
										v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
										v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
										v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
										*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
										*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
										v256 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
										*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
										v260 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
										v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
										if v262 != 0 {
											v264 = v236 + int32(72)
											*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
											v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
											v268 = v266 << (uint(int32(2)) % 32)
											if v268 == int32(0) {
											} else {
												v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
												base.MemoryCopy(m, v264, v271, v268)
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
										}
										v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
										if v278 <= int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
											v300 = int32(1)
										} else {
											v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
											if v281 == int32(1) {
												v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
												if v284 != int32(1) {
													*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
													v300 = int32(1)
												} else {
													v287 = v236 + v232
													*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
													v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
													v291 = v289 << (uint(int32(2)) % 32)
													if v291 != 0 {
														v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
														base.MemoryCopy(m, v287, v292, v291)
													} else {
													}
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
													v300 = v294 + int32(1)
												}
											} else {
												v287 = v236 + v232
												*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
												v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
												v291 = v289 << (uint(int32(2)) % 32)
												if v291 != 0 {
													v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
													base.MemoryCopy(m, v287, v292, v291)
												} else {
												}
												v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
												v300 = v294 + int32(1)
											}
										}
										*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
										*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
										*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
										F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
										mBase = m.M
										v310 = m.ExcPending
										if v310 != 0 {
											return
										} else {
											v317 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
											m.G0 = v11 + int32(16)
											return
										}
									}
								}
							} else {
								v317 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
								m.G0 = v11 + int32(16)
								return
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
								return
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v158 = m.ExcPending
								if v158 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_SetTransactionSnapshot_5), int32(0))
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										v165 = F_errdetail(m, int32(_a_F_SetTransactionSnapshot_6), int32(0))
										mBase = m.M
										v166 = m.ExcPending
										if v166 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_SetTransactionSnapshot_7), int32(570), int32(_a_F_SetTransactionSnapshot_8))
											mBase = m.M
											v171 = m.ExcPending
											if v171 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v172 = F_ProcArrayInstallImportedXmin(m, v102, l1)
				mBase = m.M
				v173 = m.ExcPending
				if v173 != 0 {
					return
				} else {
					if v172 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v325 = m.ExcPending
						if v325 != 0 {
							return
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v328 = m.ExcPending
							if v328 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_SetTransactionSnapshot_5), int32(0))
								mBase = m.M
								v332 = m.ExcPending
								if v332 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
									v335 = F_errdetail(m, int32(_a_F_SetTransactionSnapshot_9), v11)
									mBase = m.M
									v336 = m.ExcPending
									if v336 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SetTransactionSnapshot_7), int32(577), int32(_a_F_SetTransactionSnapshot_8))
										mBase = m.M
										v341 = m.ExcPending
										if v341 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					} else {
						v180 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[7]))
						if int32(2) <= v180 {
							if v180 == int32(3) {
								v186 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
								v188 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[8]))
								if v188 < int32(0) {
									v192 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[9])))
									if v192 == int32(1) {
										v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[10])))
										if v196&int32(1) != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v204 = m.ExcPending
											if v204 != 0 {
												return
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v207 = m.ExcPending
												if v207 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_SetTransactionSnapshot_1), int32(0))
													mBase = m.M
													v211 = m.ExcPending
													if v211 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_SetTransactionSnapshot_2), int32(1677), int32(_a_F_SetTransactionSnapshot_3))
														mBase = m.M
														v216 = m.ExcPending
														if v216 != 0 {
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
											v199 = F_GetSerializableTransactionSnapshotInt(m, v186, l1, l2)
											mBase = m.M
											v200 = m.ExcPending
											if v200 != 0 {
												return
											} else {
												v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
												v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
												v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
												v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
												v225 = int32(2)
												v227 = int32(72)
												v232 = v223<<(uint(v225)%32) + v227
												if int32(0) < v222 {
													v235 = (v222+v223)<<(uint(v225)%32) + v227
												} else {
													v235 = v232
												}
												v236 = F_MemoryContextAlloc(m, v219, v235)
												mBase = m.M
												v237 = m.ExcPending
												if v237 != 0 {
													return
												} else {
													v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
													v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
													v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
													v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
													v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
													v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
													v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
													v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
													*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
													*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
													v256 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
													*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
													v260 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
													v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
													if v262 != 0 {
														v264 = v236 + int32(72)
														*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
														v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
														v268 = v266 << (uint(int32(2)) % 32)
														if v268 == int32(0) {
														} else {
															v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
															base.MemoryCopy(m, v264, v271, v268)
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
													}
													v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
													if v278 <= int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
														v300 = int32(1)
													} else {
														v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
														if v281 == int32(1) {
															v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
															if v284 != int32(1) {
																*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
																v300 = int32(1)
															} else {
																v287 = v236 + v232
																*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
																v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
																v291 = v289 << (uint(int32(2)) % 32)
																if v291 != 0 {
																	v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																	base.MemoryCopy(m, v287, v292, v291)
																} else {
																}
																v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
																v300 = v294 + int32(1)
															}
														} else {
															v287 = v236 + v232
															*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
															v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
															v291 = v289 << (uint(int32(2)) % 32)
															if v291 != 0 {
																v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																base.MemoryCopy(m, v287, v292, v291)
															} else {
															}
															v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
															v300 = v294 + int32(1)
														}
													}
													*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
													*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
													*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
													F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
													mBase = m.M
													v310 = m.ExcPending
													if v310 != 0 {
														return
													} else {
														v317 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
														m.G0 = v11 + int32(16)
														return
													}
												}
											}
										}
									} else {
										v199 = F_GetSerializableTransactionSnapshotInt(m, v186, l1, l2)
										mBase = m.M
										v200 = m.ExcPending
										if v200 != 0 {
											return
										} else {
											v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
											v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
											v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
											v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
											v225 = int32(2)
											v227 = int32(72)
											v232 = v223<<(uint(v225)%32) + v227
											if int32(0) < v222 {
												v235 = (v222+v223)<<(uint(v225)%32) + v227
											} else {
												v235 = v232
											}
											v236 = F_MemoryContextAlloc(m, v219, v235)
											mBase = m.M
											v237 = m.ExcPending
											if v237 != 0 {
												return
											} else {
												v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
												*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
												v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
												*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
												v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
												*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
												v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
												*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
												v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
												*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
												v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
												*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
												v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
												*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
												v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
												*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
												*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
												v256 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
												*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
												v260 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
												v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
												if v262 != 0 {
													v264 = v236 + int32(72)
													*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
													v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
													v268 = v266 << (uint(int32(2)) % 32)
													if v268 == int32(0) {
													} else {
														v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
														base.MemoryCopy(m, v264, v271, v268)
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
												}
												v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
												if v278 <= int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
													v300 = int32(1)
												} else {
													v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
													if v281 == int32(1) {
														v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
														if v284 != int32(1) {
															*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
															v300 = int32(1)
														} else {
															v287 = v236 + v232
															*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
															v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
															v291 = v289 << (uint(int32(2)) % 32)
															if v291 != 0 {
																v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																base.MemoryCopy(m, v287, v292, v291)
															} else {
															}
															v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
															v300 = v294 + int32(1)
														}
													} else {
														v287 = v236 + v232
														*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
														v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
														v291 = v289 << (uint(int32(2)) % 32)
														if v291 != 0 {
															v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
															base.MemoryCopy(m, v287, v292, v291)
														} else {
														}
														v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
														v300 = v294 + int32(1)
													}
												}
												*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
												*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
												*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
												F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
												mBase = m.M
												v310 = m.ExcPending
												if v310 != 0 {
													return
												} else {
													v317 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
													m.G0 = v11 + int32(16)
													return
												}
											}
										}
									}
								} else {
									v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
									v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
									v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
									v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
									v225 = int32(2)
									v227 = int32(72)
									v232 = v223<<(uint(v225)%32) + v227
									if int32(0) < v222 {
										v235 = (v222+v223)<<(uint(v225)%32) + v227
									} else {
										v235 = v232
									}
									v236 = F_MemoryContextAlloc(m, v219, v235)
									mBase = m.M
									v237 = m.ExcPending
									if v237 != 0 {
										return
									} else {
										v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
										v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
										v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
										v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
										v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
										v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
										v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
										v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
										*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
										*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
										v256 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
										*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
										v260 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
										v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
										if v262 != 0 {
											v264 = v236 + int32(72)
											*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
											v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
											v268 = v266 << (uint(int32(2)) % 32)
											if v268 == int32(0) {
											} else {
												v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
												base.MemoryCopy(m, v264, v271, v268)
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
										}
										v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
										if v278 <= int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
											v300 = int32(1)
										} else {
											v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
											if v281 == int32(1) {
												v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
												if v284 != int32(1) {
													*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
													v300 = int32(1)
												} else {
													v287 = v236 + v232
													*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
													v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
													v291 = v289 << (uint(int32(2)) % 32)
													if v291 != 0 {
														v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
														base.MemoryCopy(m, v287, v292, v291)
													} else {
													}
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
													v300 = v294 + int32(1)
												}
											} else {
												v287 = v236 + v232
												*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
												v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
												v291 = v289 << (uint(int32(2)) % 32)
												if v291 != 0 {
													v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
													base.MemoryCopy(m, v287, v292, v291)
												} else {
												}
												v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
												v300 = v294 + int32(1)
											}
										}
										*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
										*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
										*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
										F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
										mBase = m.M
										v310 = m.ExcPending
										if v310 != 0 {
											return
										} else {
											v317 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
											m.G0 = v11 + int32(16)
											return
										}
									}
								}
							} else {
								v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
								v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
								v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
								v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
								v225 = int32(2)
								v227 = int32(72)
								v232 = v223<<(uint(v225)%32) + v227
								if int32(0) < v222 {
									v235 = (v222+v223)<<(uint(v225)%32) + v227
								} else {
									v235 = v232
								}
								v236 = F_MemoryContextAlloc(m, v219, v235)
								mBase = m.M
								v237 = m.ExcPending
								if v237 != 0 {
									return
								} else {
									v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
									*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
									v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
									*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
									v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
									*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
									v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
									*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
									v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
									*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
									v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
									*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
									v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
									v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
									*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
									*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
									v256 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
									*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
									v260 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
									v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
									if v262 != 0 {
										v264 = v236 + int32(72)
										*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
										v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
										v268 = v266 << (uint(int32(2)) % 32)
										if v268 == int32(0) {
										} else {
											v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
											base.MemoryCopy(m, v264, v271, v268)
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
									}
									v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
									if v278 <= int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
										v300 = int32(1)
									} else {
										v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
										if v281 == int32(1) {
											v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
											if v284 != int32(1) {
												*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
												v300 = int32(1)
											} else {
												v287 = v236 + v232
												*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
												v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
												v291 = v289 << (uint(int32(2)) % 32)
												if v291 != 0 {
													v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
													base.MemoryCopy(m, v287, v292, v291)
												} else {
												}
												v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
												v300 = v294 + int32(1)
											}
										} else {
											v287 = v236 + v232
											*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
											v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
											v291 = v289 << (uint(int32(2)) % 32)
											if v291 != 0 {
												v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
												base.MemoryCopy(m, v287, v292, v291)
											} else {
											}
											v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
											v300 = v294 + int32(1)
										}
									}
									*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
									*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
									*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
									F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
									mBase = m.M
									v310 = m.ExcPending
									if v310 != 0 {
										return
									} else {
										v317 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
										m.G0 = v11 + int32(16)
										return
									}
								}
							}
						} else {
							v317 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
							m.G0 = v11 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		F_pairingheap_remove(m, int32(_a_F_SetTransactionSnapshot_4), v14+int32(52))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[0])) = int32(0)
			v26 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[14]))
			if v26 != 0 {
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[15]))
				if v28 == int32(0) {
					v32 = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[4])) = v32
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[5]))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+52)) = v32
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[5]))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
					v41 = int32(3)
					v44 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[15]))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v44-int32(48))))
					if base.B2i32(base.Ui32(v40) < base.Ui32(v41))|base.B2i32(base.Ui32(v47) < base.Ui32(v41)) == int32(0) {
						if v40-v47 < int32(0) {
							*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[4])) = v47
							*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v47
						} else {
						}
					} else {
						if base.Ui32(v47) <= base.Ui32(v40) {
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[4])) = v47
							*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v47
						}
					}
				}
			}
			v65 = F_GetSnapshotData(m, int32(_a_F_SetTransactionSnapshot_0))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v65
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v68
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v70
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v72
				if v72 == int32(0) {
				} else {
					v77 = v72 << (uint(int32(2)) % 32)
					if v77 == int32(0) {
					} else {
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						base.MemoryCopy(m, v80, v81, v77)
					}
				}
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v65)+24)) = v84
				if v84 <= int32(0) {
				} else {
					v89 = v84 << (uint(int32(2)) % 32)
					if v89 == int32(0) {
					} else {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						base.MemoryCopy(m, v92, v93, v89)
					}
				}
				v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
				*(*uint8)(unsafe.Add(mBase, uint32(v65)+28)) = uint8(v96)
				v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
				*(*int64)(unsafe.Add(mBase, uint32(v65)+64)) = int64(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v65)+29)) = uint8(v98)
				v102 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
				if l3 != 0 {
					v104 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[2]))
					v108 = F_LWLockAcquire(m, v104+int32(512), int32(0))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return
					} else {
						v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
						v111 = int32(2)
						v116 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
						v118 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[3]))
						v124 = base.B2i32(base.Ui32(v111) < base.Ui32(v110)) & base.B2i32(base.Ui32(v111) < base.Ui32(v102)) & base.B2i32(v116 == v118) & base.B2i32(v110-v102 <= int32(0))
						if v124 != 0 {
							*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[4])) = v102
							v128 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[5]))
							*(*int32)(unsafe.Add(mBase, uint32(v128)+52)) = v102
							v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+36)))
							v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+36)))
							v136 = v130&int32(6) | v133&int32(-7)
							*(*uint8)(unsafe.Add(mBase, uint32(v128)+36)) = uint8(v136)
							v139 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[6]))
							v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
							v141 = *(*int32)(unsafe.Add(mBase, uint32(v128)+32))
							*(*uint8)(unsafe.Add(mBase, uint32(v140+v141))) = uint8(v136)
						} else {
						}
						v147 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[2]))
						F_LWLockRelease(m, v147+int32(512))
						mBase = m.M
						v151 = m.ExcPending
						if v151 != 0 {
							return
						} else {
							if v124 != 0 {
								v180 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[7]))
								if int32(2) <= v180 {
									if v180 == int32(3) {
										v186 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
										v188 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[8]))
										if v188 < int32(0) {
											v192 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[9])))
											if v192 == int32(1) {
												v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[10])))
												if v196&int32(1) != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v204 = m.ExcPending
													if v204 != 0 {
														return
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v207 = m.ExcPending
														if v207 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_SetTransactionSnapshot_1), int32(0))
															mBase = m.M
															v211 = m.ExcPending
															if v211 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_SetTransactionSnapshot_2), int32(1677), int32(_a_F_SetTransactionSnapshot_3))
																mBase = m.M
																v216 = m.ExcPending
																if v216 != 0 {
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
													v199 = F_GetSerializableTransactionSnapshotInt(m, v186, l1, l2)
													mBase = m.M
													v200 = m.ExcPending
													if v200 != 0 {
														return
													} else {
														v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
														v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
														v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
														v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
														v225 = int32(2)
														v227 = int32(72)
														v232 = v223<<(uint(v225)%32) + v227
														if int32(0) < v222 {
															v235 = (v222+v223)<<(uint(v225)%32) + v227
														} else {
															v235 = v232
														}
														v236 = F_MemoryContextAlloc(m, v219, v235)
														mBase = m.M
														v237 = m.ExcPending
														if v237 != 0 {
															return
														} else {
															v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
															*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
															v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
															*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
															v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
															*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
															v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
															*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
															v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
															*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
															v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
															*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
															v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
															*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
															v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
															*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
															*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
															v256 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
															*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
															v260 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
															v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
															if v262 != 0 {
																v264 = v236 + int32(72)
																*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
																v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
																v268 = v266 << (uint(int32(2)) % 32)
																if v268 == int32(0) {
																} else {
																	v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
																	base.MemoryCopy(m, v264, v271, v268)
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
															}
															v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
															if v278 <= int32(0) {
																*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
																v300 = int32(1)
															} else {
																v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
																if v281 == int32(1) {
																	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
																	if v284 != int32(1) {
																		*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
																		v300 = int32(1)
																	} else {
																		v287 = v236 + v232
																		*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
																		v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
																		v291 = v289 << (uint(int32(2)) % 32)
																		if v291 != 0 {
																			v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																			base.MemoryCopy(m, v287, v292, v291)
																		} else {
																		}
																		v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
																		v300 = v294 + int32(1)
																	}
																} else {
																	v287 = v236 + v232
																	*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
																	v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
																	v291 = v289 << (uint(int32(2)) % 32)
																	if v291 != 0 {
																		v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																		base.MemoryCopy(m, v287, v292, v291)
																	} else {
																	}
																	v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
																	v300 = v294 + int32(1)
																}
															}
															*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
															*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
															*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
															F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
															mBase = m.M
															v310 = m.ExcPending
															if v310 != 0 {
																return
															} else {
																v317 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
																m.G0 = v11 + int32(16)
																return
															}
														}
													}
												}
											} else {
												v199 = F_GetSerializableTransactionSnapshotInt(m, v186, l1, l2)
												mBase = m.M
												v200 = m.ExcPending
												if v200 != 0 {
													return
												} else {
													v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
													v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
													v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
													v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
													v225 = int32(2)
													v227 = int32(72)
													v232 = v223<<(uint(v225)%32) + v227
													if int32(0) < v222 {
														v235 = (v222+v223)<<(uint(v225)%32) + v227
													} else {
														v235 = v232
													}
													v236 = F_MemoryContextAlloc(m, v219, v235)
													mBase = m.M
													v237 = m.ExcPending
													if v237 != 0 {
														return
													} else {
														v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
														v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
														v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
														v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
														v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
														v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
														v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
														v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
														*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
														*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
														v256 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
														*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
														v260 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
														v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
														if v262 != 0 {
															v264 = v236 + int32(72)
															*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
															v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
															v268 = v266 << (uint(int32(2)) % 32)
															if v268 == int32(0) {
															} else {
																v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
																base.MemoryCopy(m, v264, v271, v268)
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
														}
														v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
														if v278 <= int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
															v300 = int32(1)
														} else {
															v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
															if v281 == int32(1) {
																v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
																if v284 != int32(1) {
																	*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
																	v300 = int32(1)
																} else {
																	v287 = v236 + v232
																	*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
																	v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
																	v291 = v289 << (uint(int32(2)) % 32)
																	if v291 != 0 {
																		v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																		base.MemoryCopy(m, v287, v292, v291)
																	} else {
																	}
																	v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
																	v300 = v294 + int32(1)
																}
															} else {
																v287 = v236 + v232
																*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
																v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
																v291 = v289 << (uint(int32(2)) % 32)
																if v291 != 0 {
																	v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																	base.MemoryCopy(m, v287, v292, v291)
																} else {
																}
																v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
																v300 = v294 + int32(1)
															}
														}
														*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
														*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
														*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
														F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
														mBase = m.M
														v310 = m.ExcPending
														if v310 != 0 {
															return
														} else {
															v317 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
															m.G0 = v11 + int32(16)
															return
														}
													}
												}
											}
										} else {
											v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
											v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
											v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
											v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
											v225 = int32(2)
											v227 = int32(72)
											v232 = v223<<(uint(v225)%32) + v227
											if int32(0) < v222 {
												v235 = (v222+v223)<<(uint(v225)%32) + v227
											} else {
												v235 = v232
											}
											v236 = F_MemoryContextAlloc(m, v219, v235)
											mBase = m.M
											v237 = m.ExcPending
											if v237 != 0 {
												return
											} else {
												v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
												*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
												v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
												*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
												v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
												*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
												v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
												*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
												v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
												*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
												v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
												*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
												v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
												*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
												v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
												*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
												*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
												v256 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
												*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
												v260 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
												v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
												if v262 != 0 {
													v264 = v236 + int32(72)
													*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
													v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
													v268 = v266 << (uint(int32(2)) % 32)
													if v268 == int32(0) {
													} else {
														v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
														base.MemoryCopy(m, v264, v271, v268)
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
												}
												v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
												if v278 <= int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
													v300 = int32(1)
												} else {
													v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
													if v281 == int32(1) {
														v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
														if v284 != int32(1) {
															*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
															v300 = int32(1)
														} else {
															v287 = v236 + v232
															*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
															v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
															v291 = v289 << (uint(int32(2)) % 32)
															if v291 != 0 {
																v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																base.MemoryCopy(m, v287, v292, v291)
															} else {
															}
															v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
															v300 = v294 + int32(1)
														}
													} else {
														v287 = v236 + v232
														*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
														v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
														v291 = v289 << (uint(int32(2)) % 32)
														if v291 != 0 {
															v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
															base.MemoryCopy(m, v287, v292, v291)
														} else {
														}
														v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
														v300 = v294 + int32(1)
													}
												}
												*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
												*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
												*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
												F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
												mBase = m.M
												v310 = m.ExcPending
												if v310 != 0 {
													return
												} else {
													v317 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
													m.G0 = v11 + int32(16)
													return
												}
											}
										}
									} else {
										v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
										v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
										v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
										v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
										v225 = int32(2)
										v227 = int32(72)
										v232 = v223<<(uint(v225)%32) + v227
										if int32(0) < v222 {
											v235 = (v222+v223)<<(uint(v225)%32) + v227
										} else {
											v235 = v232
										}
										v236 = F_MemoryContextAlloc(m, v219, v235)
										mBase = m.M
										v237 = m.ExcPending
										if v237 != 0 {
											return
										} else {
											v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
											v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
											v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
											v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
											v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
											v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
											v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
											v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
											*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
											*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
											v256 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
											*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
											v260 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
											v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
											if v262 != 0 {
												v264 = v236 + int32(72)
												*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
												v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
												v268 = v266 << (uint(int32(2)) % 32)
												if v268 == int32(0) {
												} else {
													v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
													base.MemoryCopy(m, v264, v271, v268)
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
											}
											v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
											if v278 <= int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
												v300 = int32(1)
											} else {
												v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
												if v281 == int32(1) {
													v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
													if v284 != int32(1) {
														*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
														v300 = int32(1)
													} else {
														v287 = v236 + v232
														*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
														v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
														v291 = v289 << (uint(int32(2)) % 32)
														if v291 != 0 {
															v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
															base.MemoryCopy(m, v287, v292, v291)
														} else {
														}
														v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
														v300 = v294 + int32(1)
													}
												} else {
													v287 = v236 + v232
													*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
													v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
													v291 = v289 << (uint(int32(2)) % 32)
													if v291 != 0 {
														v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
														base.MemoryCopy(m, v287, v292, v291)
													} else {
													}
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
													v300 = v294 + int32(1)
												}
											}
											*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
											*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
											*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
											F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
											mBase = m.M
											v310 = m.ExcPending
											if v310 != 0 {
												return
											} else {
												v317 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
												m.G0 = v11 + int32(16)
												return
											}
										}
									}
								} else {
									v317 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
									m.G0 = v11 + int32(16)
									return
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v155 = m.ExcPending
								if v155 != 0 {
									return
								} else {
									F_errcode(m, int32(325))
									mBase = m.M
									v158 = m.ExcPending
									if v158 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_SetTransactionSnapshot_5), int32(0))
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return
										} else {
											v165 = F_errdetail(m, int32(_a_F_SetTransactionSnapshot_6), int32(0))
											mBase = m.M
											v166 = m.ExcPending
											if v166 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_SetTransactionSnapshot_7), int32(570), int32(_a_F_SetTransactionSnapshot_8))
												mBase = m.M
												v171 = m.ExcPending
												if v171 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v172 = F_ProcArrayInstallImportedXmin(m, v102, l1)
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
						return
					} else {
						if v172 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v325 = m.ExcPending
							if v325 != 0 {
								return
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v328 = m.ExcPending
								if v328 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_SetTransactionSnapshot_5), int32(0))
									mBase = m.M
									v332 = m.ExcPending
									if v332 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
										v335 = F_errdetail(m, int32(_a_F_SetTransactionSnapshot_9), v11)
										mBase = m.M
										v336 = m.ExcPending
										if v336 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_SetTransactionSnapshot_7), int32(577), int32(_a_F_SetTransactionSnapshot_8))
											mBase = m.M
											v341 = m.ExcPending
											if v341 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						} else {
							v180 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[7]))
							if int32(2) <= v180 {
								if v180 == int32(3) {
									v186 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
									v188 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[8]))
									if v188 < int32(0) {
										v192 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[9])))
										if v192 == int32(1) {
											v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[10])))
											if v196&int32(1) != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v204 = m.ExcPending
												if v204 != 0 {
													return
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v207 = m.ExcPending
													if v207 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_SetTransactionSnapshot_1), int32(0))
														mBase = m.M
														v211 = m.ExcPending
														if v211 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_SetTransactionSnapshot_2), int32(1677), int32(_a_F_SetTransactionSnapshot_3))
															mBase = m.M
															v216 = m.ExcPending
															if v216 != 0 {
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
												v199 = F_GetSerializableTransactionSnapshotInt(m, v186, l1, l2)
												mBase = m.M
												v200 = m.ExcPending
												if v200 != 0 {
													return
												} else {
													v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
													v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
													v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
													v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
													v225 = int32(2)
													v227 = int32(72)
													v232 = v223<<(uint(v225)%32) + v227
													if int32(0) < v222 {
														v235 = (v222+v223)<<(uint(v225)%32) + v227
													} else {
														v235 = v232
													}
													v236 = F_MemoryContextAlloc(m, v219, v235)
													mBase = m.M
													v237 = m.ExcPending
													if v237 != 0 {
														return
													} else {
														v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
														v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
														v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
														v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
														v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
														v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
														v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
														v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
														*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
														*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
														v256 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
														*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
														v260 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
														v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
														if v262 != 0 {
															v264 = v236 + int32(72)
															*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
															v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
															v268 = v266 << (uint(int32(2)) % 32)
															if v268 == int32(0) {
															} else {
																v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
																base.MemoryCopy(m, v264, v271, v268)
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
														}
														v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
														if v278 <= int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
															v300 = int32(1)
														} else {
															v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
															if v281 == int32(1) {
																v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
																if v284 != int32(1) {
																	*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
																	v300 = int32(1)
																} else {
																	v287 = v236 + v232
																	*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
																	v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
																	v291 = v289 << (uint(int32(2)) % 32)
																	if v291 != 0 {
																		v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																		base.MemoryCopy(m, v287, v292, v291)
																	} else {
																	}
																	v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
																	v300 = v294 + int32(1)
																}
															} else {
																v287 = v236 + v232
																*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
																v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
																v291 = v289 << (uint(int32(2)) % 32)
																if v291 != 0 {
																	v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																	base.MemoryCopy(m, v287, v292, v291)
																} else {
																}
																v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
																v300 = v294 + int32(1)
															}
														}
														*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
														*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
														*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
														F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
														mBase = m.M
														v310 = m.ExcPending
														if v310 != 0 {
															return
														} else {
															v317 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
															m.G0 = v11 + int32(16)
															return
														}
													}
												}
											}
										} else {
											v199 = F_GetSerializableTransactionSnapshotInt(m, v186, l1, l2)
											mBase = m.M
											v200 = m.ExcPending
											if v200 != 0 {
												return
											} else {
												v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
												v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
												v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
												v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
												v225 = int32(2)
												v227 = int32(72)
												v232 = v223<<(uint(v225)%32) + v227
												if int32(0) < v222 {
													v235 = (v222+v223)<<(uint(v225)%32) + v227
												} else {
													v235 = v232
												}
												v236 = F_MemoryContextAlloc(m, v219, v235)
												mBase = m.M
												v237 = m.ExcPending
												if v237 != 0 {
													return
												} else {
													v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
													v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
													v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
													v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
													v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
													v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
													v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
													v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
													*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
													*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
													v256 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
													*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
													v260 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
													v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
													if v262 != 0 {
														v264 = v236 + int32(72)
														*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
														v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
														v268 = v266 << (uint(int32(2)) % 32)
														if v268 == int32(0) {
														} else {
															v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
															base.MemoryCopy(m, v264, v271, v268)
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
													}
													v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
													if v278 <= int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
														v300 = int32(1)
													} else {
														v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
														if v281 == int32(1) {
															v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
															if v284 != int32(1) {
																*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
																v300 = int32(1)
															} else {
																v287 = v236 + v232
																*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
																v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
																v291 = v289 << (uint(int32(2)) % 32)
																if v291 != 0 {
																	v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																	base.MemoryCopy(m, v287, v292, v291)
																} else {
																}
																v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
																v300 = v294 + int32(1)
															}
														} else {
															v287 = v236 + v232
															*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
															v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
															v291 = v289 << (uint(int32(2)) % 32)
															if v291 != 0 {
																v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
																base.MemoryCopy(m, v287, v292, v291)
															} else {
															}
															v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
															v300 = v294 + int32(1)
														}
													}
													*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
													*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
													*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
													F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
													mBase = m.M
													v310 = m.ExcPending
													if v310 != 0 {
														return
													} else {
														v317 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
														m.G0 = v11 + int32(16)
														return
													}
												}
											}
										}
									} else {
										v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
										v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
										v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
										v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
										v225 = int32(2)
										v227 = int32(72)
										v232 = v223<<(uint(v225)%32) + v227
										if int32(0) < v222 {
											v235 = (v222+v223)<<(uint(v225)%32) + v227
										} else {
											v235 = v232
										}
										v236 = F_MemoryContextAlloc(m, v219, v235)
										mBase = m.M
										v237 = m.ExcPending
										if v237 != 0 {
											return
										} else {
											v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
											v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
											v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
											v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
											v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
											v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
											v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
											v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
											*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
											*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
											v256 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
											*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
											v260 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
											v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
											if v262 != 0 {
												v264 = v236 + int32(72)
												*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
												v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
												v268 = v266 << (uint(int32(2)) % 32)
												if v268 == int32(0) {
												} else {
													v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
													base.MemoryCopy(m, v264, v271, v268)
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
											}
											v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
											if v278 <= int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
												v300 = int32(1)
											} else {
												v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
												if v281 == int32(1) {
													v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
													if v284 != int32(1) {
														*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
														v300 = int32(1)
													} else {
														v287 = v236 + v232
														*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
														v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
														v291 = v289 << (uint(int32(2)) % 32)
														if v291 != 0 {
															v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
															base.MemoryCopy(m, v287, v292, v291)
														} else {
														}
														v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
														v300 = v294 + int32(1)
													}
												} else {
													v287 = v236 + v232
													*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
													v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
													v291 = v289 << (uint(int32(2)) % 32)
													if v291 != 0 {
														v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
														base.MemoryCopy(m, v287, v292, v291)
													} else {
													}
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
													v300 = v294 + int32(1)
												}
											}
											*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
											*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
											*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
											F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
											mBase = m.M
											v310 = m.ExcPending
											if v310 != 0 {
												return
											} else {
												v317 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
												m.G0 = v11 + int32(16)
												return
											}
										}
									}
								} else {
									v219 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
									v221 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
									v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
									v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
									v225 = int32(2)
									v227 = int32(72)
									v232 = v223<<(uint(v225)%32) + v227
									if int32(0) < v222 {
										v235 = (v222+v223)<<(uint(v225)%32) + v227
									} else {
										v235 = v232
									}
									v236 = F_MemoryContextAlloc(m, v219, v235)
									mBase = m.M
									v237 = m.ExcPending
									if v237 != 0 {
										return
									} else {
										v238 = *(*int64)(unsafe.Add(mBase, uint32(v221)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+48)) = v238
										v240 = *(*int64)(unsafe.Add(mBase, uint32(v221)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+40)) = v240
										v242 = *(*int64)(unsafe.Add(mBase, uint32(v221)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v242
										v244 = *(*int64)(unsafe.Add(mBase, uint32(v221)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+56)) = v244
										v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v246
										v248 = *(*int64)(unsafe.Add(mBase, uint32(v221)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v248
										v250 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v250
										v252 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
										*(*int64)(unsafe.Add(mBase, uint32(v236))) = v252
										*(*int64)(unsafe.Add(mBase, uint32(v236)+64)) = int64(0)
										v256 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v256
										*(*int32)(unsafe.Add(mBase, uint32(v236)+44)) = v256
										v260 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v236)+30)) = uint8(v260)
										v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
										if v262 != 0 {
											v264 = v236 + int32(72)
											*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v264
											v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
											v268 = v266 << (uint(int32(2)) % 32)
											if v268 == int32(0) {
											} else {
												v271 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
												base.MemoryCopy(m, v264, v271, v268)
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = int32(0)
										}
										v278 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
										if v278 <= int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
											v300 = int32(1)
										} else {
											v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
											if v281 == int32(1) {
												v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+29)))
												if v284 != int32(1) {
													*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = int32(0)
													v300 = int32(1)
												} else {
													v287 = v236 + v232
													*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
													v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
													v291 = v289 << (uint(int32(2)) % 32)
													if v291 != 0 {
														v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
														base.MemoryCopy(m, v287, v292, v291)
													} else {
													}
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
													v300 = v294 + int32(1)
												}
											} else {
												v287 = v236 + v232
												*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v287
												v289 = *(*int32)(unsafe.Add(mBase, uint32(v221)+24))
												v291 = v289 << (uint(int32(2)) % 32)
												if v291 != 0 {
													v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
													base.MemoryCopy(m, v287, v292, v291)
												} else {
												}
												v294 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
												v300 = v294 + int32(1)
											}
										}
										*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v236
										*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v236
										*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v300
										F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v236+int32(52))
										mBase = m.M
										v310 = m.ExcPending
										if v310 != 0 {
											return
										} else {
											v317 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
											m.G0 = v11 + int32(16)
											return
										}
									}
								}
							} else {
								v317 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v317)
								m.G0 = v11 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_TransactionIdCommitTree(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v7 int32
	_ = v7
	F_TransactionIdSetTreeStatus(m, l0, l1, l2, int32(1), int64(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_TransactionIdDidAbort(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidAbort[0]))
	if v10 == l0 {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidAbort[1]))
		v30 = v13
		v31 = int32(1)
		switch v30 - int32(2) {
		case 0:
			v72 = v31
			m.G0 = v7 + int32(16)
			return v72
		case 1:
			v34 = int32(3)
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidAbort[2]))
			if base.B2i32(base.Ui32(l0) < base.Ui32(v34))|base.B2i32(base.Ui32(v37) < base.Ui32(v34)) == int32(0) {
				if int32(0) <= l0-v37 {
					v47 = F_SubTransGetParent(m, l0)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						if v47 == int32(0) {
							v53 = F_errstart(m, int32(19), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								if v53 == int32(0) {
									v72 = v31
									m.G0 = v7 + int32(16)
									return v72
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
									F_errmsg_internal(m, int32(_a_F_TransactionIdDidAbort_0), v7)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_TransactionIdDidAbort_1), int32(217), int32(_a_F_TransactionIdDidAbort_2))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v72 = v31
											m.G0 = v7 + int32(16)
											return v72
										}
									}
								}
							}
						} else {
							v66 = F_TransactionIdDidAbort(m, v47)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v72 = v66
								m.G0 = v7 + int32(16)
								return v72
							}
						}
					}
				} else {
					v72 = v31
					m.G0 = v7 + int32(16)
					return v72
				}
			} else {
				if base.Ui32(l0) < base.Ui32(v37) {
					v72 = v31
					m.G0 = v7 + int32(16)
					return v72
				} else {
					v47 = F_SubTransGetParent(m, l0)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						if v47 == int32(0) {
							v53 = F_errstart(m, int32(19), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								if v53 == int32(0) {
									v72 = v31
									m.G0 = v7 + int32(16)
									return v72
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
									F_errmsg_internal(m, int32(_a_F_TransactionIdDidAbort_0), v7)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_TransactionIdDidAbort_1), int32(217), int32(_a_F_TransactionIdDidAbort_2))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v72 = v31
											m.G0 = v7 + int32(16)
											return v72
										}
									}
								}
							}
						} else {
							v66 = F_TransactionIdDidAbort(m, v47)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v72 = v66
								m.G0 = v7 + int32(16)
								return v72
							}
						}
					}
				}
			}
		default:
			v72 = int32(0)
			m.G0 = v7 + int32(16)
			return v72
		}
	} else {
		if base.Ui32(l0) <= base.Ui32(int32(2)) {
			if l0 != 0 {
				v72 = int32(0)
			} else {
				v72 = int32(1)
			}
			m.G0 = v7 + int32(16)
			return v72
		} else {
			v19 = F_TransactionIdGetStatus(m, l0, v7+int32(8))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				switch v19 {
				case 0, 3:
					v30 = v19
				default:
					*(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidAbort[1])) = v19
					*(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidAbort[0])) = l0
					v28 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
					*(*int64)(unsafe.Add(mBase, _c_F_TransactionIdDidAbort[3])) = v28
					v30 = v19
				}
				v31 = int32(1)
				switch v30 - int32(2) {
				case 0:
					v72 = v31
					m.G0 = v7 + int32(16)
					return v72
				case 1:
					v34 = int32(3)
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidAbort[2]))
					if base.B2i32(base.Ui32(l0) < base.Ui32(v34))|base.B2i32(base.Ui32(v37) < base.Ui32(v34)) == int32(0) {
						if int32(0) <= l0-v37 {
							v47 = F_SubTransGetParent(m, l0)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								if v47 == int32(0) {
									v53 = F_errstart(m, int32(19), int32(0))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										if v53 == int32(0) {
											v72 = v31
											m.G0 = v7 + int32(16)
											return v72
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
											F_errmsg_internal(m, int32(_a_F_TransactionIdDidAbort_0), v7)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_TransactionIdDidAbort_1), int32(217), int32(_a_F_TransactionIdDidAbort_2))
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													v72 = v31
													m.G0 = v7 + int32(16)
													return v72
												}
											}
										}
									}
								} else {
									v66 = F_TransactionIdDidAbort(m, v47)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										v72 = v66
										m.G0 = v7 + int32(16)
										return v72
									}
								}
							}
						} else {
							v72 = v31
							m.G0 = v7 + int32(16)
							return v72
						}
					} else {
						if base.Ui32(l0) < base.Ui32(v37) {
							v72 = v31
							m.G0 = v7 + int32(16)
							return v72
						} else {
							v47 = F_SubTransGetParent(m, l0)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								if v47 == int32(0) {
									v53 = F_errstart(m, int32(19), int32(0))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										if v53 == int32(0) {
											v72 = v31
											m.G0 = v7 + int32(16)
											return v72
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
											F_errmsg_internal(m, int32(_a_F_TransactionIdDidAbort_0), v7)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_TransactionIdDidAbort_1), int32(217), int32(_a_F_TransactionIdDidAbort_2))
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													v72 = v31
													m.G0 = v7 + int32(16)
													return v72
												}
											}
										}
									}
								} else {
									v66 = F_TransactionIdDidAbort(m, v47)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										v72 = v66
										m.G0 = v7 + int32(16)
										return v72
									}
								}
							}
						}
					}
				default:
					v72 = int32(0)
					m.G0 = v7 + int32(16)
					return v72
				}
			}
		}
	}
}
func F_assign_transaction_timeout(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_assign_transaction_timeout[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	if base.B2i32(v5 == int32(2)) == int32(0) {
		return
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_assign_transaction_timeout[1])))
		if int32(0) < l0 {
			if v13 != 0 {
				return
			} else {
				F_enable_timeout_after(m, int32(8), l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			if v13 == int32(0) {
				return
			} else {
				F_disable_timeout(m, int32(8))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
