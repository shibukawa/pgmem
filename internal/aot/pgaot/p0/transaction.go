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
				F_errfinish(m, int32(_a_F_BeginTransactionBlock_2), int32(3974), int32(_a_F_BeginTransactionBlock_3))
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
						F_errfinish(m, int32(_a_F_BeginTransactionBlock_2), int32(3956), int32(_a_F_BeginTransactionBlock_3))
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
	F_errfinish(m, int32(_a_F_CommitTransactionCommand_2), int32(3194), int32(_a_F_CommitTransactionCommand_3))
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
	F_errfinish(m, int32(_a_F_CommitTransactionCommand_2), int32(3360), int32(_a_F_CommitTransactionCommand_3))
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int64
	_ = v283
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v293 int64
	_ = v293
	var v295 int64
	_ = v295
	var v297 int64
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
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
	v436 = m.ExcPending
	if v436 != 0 {
		goto L10
	} else {
		goto L120
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
	v431 = v7
	goto L4
L4:
	;
	return v431
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
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[3]))
	if int32(2) <= v373 {
		goto L104
	} else {
		goto L105
	}
L8:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[4]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+72))
	if v66 != 0 {
		goto L22
	} else {
		goto L23
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
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[6]))
	if v32 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[7]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[6]))
	v39 = v37 - int32(48)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v40))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v35)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v56 = int32(0)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[8])) = v56
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+40)) = v56
	goto L8
L16:
	;
	if v52 == int32(0) {
		goto L8
	} else {
		goto L20
	}
L17:
	;
	v52 = base.B2i32(base.Ui32(v35) < base.Ui32(v40))
	goto L16
L18:
	;
	goto L19
L19:
	;
	v52 = int32(base.Ui32(v35-v40) >> (uint(int32(31)) % 32))
	goto L16
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v56 = v55
	goto L15
L21:
	;
	if v69&int32(1) != 0 {
		goto L1
	} else {
		goto L25
	}
L22:
	;
	v69 = int32(1)
	goto L24
L23:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+76)))
	v69 = v68
	goto L24
L24:
	;
	goto L21
L25:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[3]))
	if int32(2) <= v73 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v369 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[1])) = uint8(v369)
	return v363
L27:
	;
	if v73 == int32(3) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L29
L29:
	;
	v360 = F_GetSnapshotData(m, int32(_a_F_GetTransactionSnapshot_1))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L10
	} else {
		goto L103
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[9])) = v263
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[10]))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v263)+24))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v263)+16))
	v270 = int32(2)
	v272 = int32(72)
	v277 = v268<<(uint(v270)%32) + v272
	if int32(0) < v267 {
		goto L83
	} else {
		goto L84
	}
L31:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[11])))
	if v81 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	goto L33
L33:
	;
	v256 = F_GetSnapshotData(m, int32(_a_F_GetTransactionSnapshot_1))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L10
	} else {
		goto L82
	}
L34:
	;
	v263 = v252
	goto L30
L35:
	;
	if v91 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[12]))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+316))
	v89 = base.B2i32(v87 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[11])) = uint8(v89)
	v91 = v89
	goto L38
L37:
	;
	v91 = int32(0)
	goto L38
L38:
	;
	goto L35
L39:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[13])))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[14])))
	v101 = F_GetSerializableTransactionSnapshotInt(m, int32(_a_F_GetTransactionSnapshot_1), int32(0), int32(-1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L10
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L10
	} else {
		goto L76
	}
L42:
	;
	v103 = int32(1)
	if base.B2i32(v97&v103 == int32(0))|base.B2i32(v95 != v103) != 0 {
		v223 = v101
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v252 = v223
	goto L34
L44:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[15]))
	if v111 == int32(0) {
		v223 = v101
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v116 = v101
	goto L46
L46:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[16]))
	v124 = F_LWLockAcquire(m, v120+int32(3584), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L10
	} else {
		goto L48
	}
L47:
	;
	v223 = v217
	goto L43
L48:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[15]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+108))
	v130 = v128 | int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(v127)+108)) = v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)+92))
	if base.B2i32(v132 == int32(0))|base.B2i32(v132 == v127+int32(88)) != 0 {
		v172 = v127
		v173 = v130
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+108)) = v173 & int32(-65)
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[16]))
	F_LWLockRelease(m, v181+int32(3584))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L10
	} else {
		goto L61
	}
L50:
	;
	v139 = v127
	goto L51
L51:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v139)+108))
	if v144&int32(256) != 0 {
		v172 = v139
		v173 = v144
		goto L49
	} else {
		goto L53
	}
L52:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v164)+108))
	v172 = v164
	v173 = v171
	goto L49
L53:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[16]))
	F_LWLockRelease(m, v148+int32(3584))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L10
	} else {
		goto L54
	}
L54:
	;
	F_ProcWaitForSignal(m, int32(134217779))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[16]))
	v161 = F_LWLockAcquire(m, v157+int32(3584), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[15]))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+92))
	if v165 != v164+int32(88) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v170 = v165
	goto L59
L58:
	;
	v170 = int32(0)
	goto L59
L59:
	;
	if v170 != 0 {
		v139 = v164
		goto L51
	} else {
		goto L60
	}
L60:
	;
	goto L52
L61:
	;
	if v173&int32(256) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_ReleasePredicateLocks(m, int32(0), int32(1))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L10
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v196 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L10
	} else {
		goto L66
	}
L65:
	;
	v252 = v116
	goto L34
L66:
	;
	if v196 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L10
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v210 = int32(0)
	F_ReleasePredicateLocks(m, v210, v210)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L10
	} else {
		goto L73
	}
L70:
	;
	F_errmsg_internal(m, int32(_a_F_GetTransactionSnapshot_2), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L10
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_GetTransactionSnapshot_3), int32(1605), int32(_a_F_GetTransactionSnapshot_4))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	v217 = F_GetSerializableTransactionSnapshotInt(m, int32(_a_F_GetTransactionSnapshot_1), int32(0), int32(-1))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[15]))
	if v220 != 0 {
		v116 = v217
		goto L46
	} else {
		goto L75
	}
L75:
	;
	goto L47
L76:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L10
	} else {
		goto L77
	}
L77:
	;
	F_errmsg(m, int32(_a_F_GetTransactionSnapshot_5), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L10
	} else {
		goto L78
	}
L78:
	;
	F_errdetail(m, int32(_a_F_GetTransactionSnapshot_6), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L10
	} else {
		goto L79
	}
L79:
	;
	F_errhint(m, int32(_a_F_GetTransactionSnapshot_7), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L10
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_GetTransactionSnapshot_3), int32(1697), int32(_a_F_GetTransactionSnapshot_8))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L10
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	v263 = v256
	goto L30
L83:
	;
	v280 = (v267+v268)<<(uint(v270)%32) + v272
	goto L85
L84:
	;
	v280 = v277
	goto L85
L85:
	;
	v281 = F_MemoryContextAlloc(m, v266, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v263)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v281)+48)) = v283
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v263)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v281)+40)) = v285
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v263)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v281)+24)) = v287
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v263)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v281)+56)) = v289
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v263)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v281)+32)) = v291
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v263)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v281)+16)) = v293
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v263)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v281)+8)) = v295
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v263)))
	*(*int64)(unsafe.Add(mBase, uint32(v281))) = v297
	*(*int64)(unsafe.Add(mBase, uint32(v281)+64)) = int64(0)
	v301 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v281)+48)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(v281)+44)) = v301
	v305 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v281)+30)) = uint8(v305)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v263)+16))
	if v307 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v263)+24))
	if v323 <= int32(0) {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	v309 = v281 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v281)+12)) = v309
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v263)+16))
	v313 = v311 << (uint(int32(2)) % 32)
	if v313 == int32(0) {
		goto L87
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281)+12)) = int32(0)
	goto L87
L91:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	base.MemoryCopy(m, v309, v316, v313)
	goto L87
L92:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[17])) = v281
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[9])) = v281
	*(*int32)(unsafe.Add(mBase, uint32(v281)+48)) = v344
	F_pairingheap_add(m, int32(_a_F_GetTransactionSnapshot_0), v281+int32(52))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L10
	} else {
		goto L102
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281)+20)) = int32(0)
	v344 = int32(1)
	goto L92
L94:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+28)))
	if v326 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+29)))
	if v329 != int32(1) {
		goto L93
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v332 = v281 + v277
	*(*int32)(unsafe.Add(mBase, uint32(v281)+20)) = v332
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v263)+24))
	v336 = v334 << (uint(int32(2)) % 32)
	if v336 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L97
L99:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v263)+20))
	base.MemoryCopy(m, v332, v337, v336)
	goto L101
L100:
	;
	goto L101
L101:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v281)+48))
	v344 = v339 + int32(1)
	goto L92
L102:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[9]))
	v363 = v357
	goto L26
L103:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[9])) = v360
	v363 = v360
	goto L26
L104:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[9]))
	return v377
L105:
	;
	goto L106
L106:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[2]))
	if v380 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v428 = F_GetSnapshotData(m, int32(_a_F_GetTransactionSnapshot_1))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L10
	} else {
		goto L119
	}
L108:
	;
	F_pairingheap_remove(m, int32(_a_F_GetTransactionSnapshot_0), v380+int32(52))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	v388 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[2])) = v388
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[5]))
	if v393 != 0 {
		goto L107
	} else {
		goto L110
	}
L110:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[6]))
	if v395 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[7]))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)+40))
	v400 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[6]))
	v402 = v400 - int32(48)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v403))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v398)) == int32(0) {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v419 = v388
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[8])) = v419
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v423)+40)) = v419
	goto L107
L114:
	;
	if v415 == int32(0) {
		goto L107
	} else {
		goto L118
	}
L115:
	;
	v415 = base.B2i32(base.Ui32(v398) < base.Ui32(v403))
	goto L114
L116:
	;
	goto L117
L117:
	;
	v415 = int32(base.Ui32(v398-v403) >> (uint(int32(31)) % 32))
	goto L114
L118:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	v419 = v418
	goto L113
L119:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetTransactionSnapshot[9])) = v428
	v431 = v428
	goto L4
L120:
	;
	F_errmsg_internal(m, int32(_a_F_GetTransactionSnapshot_9), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L10
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_GetTransactionSnapshot_10), int32(306), int32(_a_F_GetTransactionSnapshot_11))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L10
	} else {
		goto L122
	}
L122:
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
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
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int64
	_ = v243
	var v245 int64
	_ = v245
	var v247 int64
	_ = v247
	var v249 int64
	_ = v249
	var v251 int64
	_ = v251
	var v253 int64
	_ = v253
	var v255 int64
	_ = v255
	var v257 int64
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[0]))
	if v13 == int32(0) {
		v61 = F_GetSnapshotData(m, int32(_a_F_SetTransactionSnapshot_0))
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v61
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v64
			v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v66
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v68
			if v68 == int32(0) {
			} else {
				v73 = v68 << (uint(int32(2)) % 32)
				if v73 == int32(0) {
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					base.MemoryCopy(m, v76, v77, v73)
				}
			}
			v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = v80
			if v80 <= int32(0) {
			} else {
				v85 = v80 << (uint(int32(2)) % 32)
				if v85 == int32(0) {
				} else {
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					base.MemoryCopy(m, v88, v89, v85)
				}
			}
			v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
			*(*uint8)(unsafe.Add(mBase, uint32(v61)+28)) = uint8(v92)
			v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
			*(*int64)(unsafe.Add(mBase, uint32(v61)+64)) = int64(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v61)+29)) = uint8(v94)
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
			if l3 != 0 {
				v99 = int32(0)
				v101 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[2]))
				v105 = F_LWLockAcquire(m, v101+int32(512), v99)
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return
				} else {
					v107 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
					v109 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[3]))
					v111 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
					if base.B2i32(v107 != v109)|base.B2i32(base.Ui32(v111) < base.Ui32(int32(3))) != 0 {
						v149 = v99
					} else {
						if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v98))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v111)) == int32(0) {
							v126 = base.B2i32(base.Ui32(v111) <= base.Ui32(v98))
						} else {
							v126 = base.B2i32(v111-v98 <= int32(0))
						}
						if v126 == int32(0) {
							v149 = v99
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[4])) = v98
							v132 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[5]))
							*(*int32)(unsafe.Add(mBase, uint32(v132)+40)) = v98
							v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+124)))
							v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+124)))
							v140 = v134&int32(6) | v137&int32(-7)
							*(*uint8)(unsafe.Add(mBase, uint32(v132)+124)) = uint8(v140)
							v143 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[6]))
							v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
							v145 = *(*int32)(unsafe.Add(mBase, uint32(v132)+48))
							*(*uint8)(unsafe.Add(mBase, uint32(v144+v145))) = uint8(v140)
							v149 = int32(1)
						}
					}
					v152 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[2]))
					F_LWLockRelease(m, v152+int32(512))
					mBase = m.M
					v156 = m.ExcPending
					if v156 != 0 {
						return
					} else {
						if v149 != 0 {
							v185 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[7]))
							if int32(2) <= v185 {
								if v185 == int32(3) {
									v191 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
									v193 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[8]))
									if v193 < int32(0) {
										v197 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[9])))
										if v197 == int32(1) {
											v201 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[10])))
											if v201&int32(1) != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v209 = m.ExcPending
												if v209 != 0 {
													return
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_SetTransactionSnapshot_1), int32(0))
														mBase = m.M
														v216 = m.ExcPending
														if v216 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_SetTransactionSnapshot_2), int32(1748), int32(_a_F_SetTransactionSnapshot_3))
															mBase = m.M
															v221 = m.ExcPending
															if v221 != 0 {
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
												v204 = F_GetSerializableTransactionSnapshotInt(m, v191, l1, l2)
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return
												} else {
													v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
													v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
													v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
													v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
													v230 = int32(2)
													v232 = int32(72)
													v237 = v228<<(uint(v230)%32) + v232
													if int32(0) < v227 {
														v240 = (v227+v228)<<(uint(v230)%32) + v232
													} else {
														v240 = v237
													}
													v241 = F_MemoryContextAlloc(m, v224, v240)
													mBase = m.M
													v242 = m.ExcPending
													if v242 != 0 {
														return
													} else {
														v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
														v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
														v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
														v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
														v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
														v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
														v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
														v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
														*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
														*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
														v261 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
														*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
														v265 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
														v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
														if v267 != 0 {
															v269 = v241 + int32(72)
															*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
															v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
															v273 = v271 << (uint(int32(2)) % 32)
															if v273 == int32(0) {
															} else {
																v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
																base.MemoryCopy(m, v269, v276, v273)
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
														}
														v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
														if v283 <= int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
															v305 = int32(1)
														} else {
															v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
															if v286 == int32(1) {
																v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
																if v289 != int32(1) {
																	*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
																	v305 = int32(1)
																} else {
																	v292 = v241 + v237
																	*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
																	v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
																	v296 = v294 << (uint(int32(2)) % 32)
																	if v296 != 0 {
																		v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																		base.MemoryCopy(m, v292, v297, v296)
																	} else {
																	}
																	v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
																	v305 = v299 + int32(1)
																}
															} else {
																v292 = v241 + v237
																*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
																v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
																v296 = v294 << (uint(int32(2)) % 32)
																if v296 != 0 {
																	v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																	base.MemoryCopy(m, v292, v297, v296)
																} else {
																}
																v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
																v305 = v299 + int32(1)
															}
														}
														*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
														*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
														*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
														F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
														mBase = m.M
														v315 = m.ExcPending
														if v315 != 0 {
															return
														} else {
															v322 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
															m.G0 = v10 + int32(16)
															return
														}
													}
												}
											}
										} else {
											v204 = F_GetSerializableTransactionSnapshotInt(m, v191, l1, l2)
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
												return
											} else {
												v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
												v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
												v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
												v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
												v230 = int32(2)
												v232 = int32(72)
												v237 = v228<<(uint(v230)%32) + v232
												if int32(0) < v227 {
													v240 = (v227+v228)<<(uint(v230)%32) + v232
												} else {
													v240 = v237
												}
												v241 = F_MemoryContextAlloc(m, v224, v240)
												mBase = m.M
												v242 = m.ExcPending
												if v242 != 0 {
													return
												} else {
													v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
													v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
													v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
													v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
													v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
													v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
													v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
													v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
													*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
													*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
													v261 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
													*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
													v265 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
													v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
													if v267 != 0 {
														v269 = v241 + int32(72)
														*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
														v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
														v273 = v271 << (uint(int32(2)) % 32)
														if v273 == int32(0) {
														} else {
															v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
															base.MemoryCopy(m, v269, v276, v273)
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
													}
													v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
													if v283 <= int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
														v305 = int32(1)
													} else {
														v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
														if v286 == int32(1) {
															v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
															if v289 != int32(1) {
																*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
																v305 = int32(1)
															} else {
																v292 = v241 + v237
																*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
																v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
																v296 = v294 << (uint(int32(2)) % 32)
																if v296 != 0 {
																	v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																	base.MemoryCopy(m, v292, v297, v296)
																} else {
																}
																v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
																v305 = v299 + int32(1)
															}
														} else {
															v292 = v241 + v237
															*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
															v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
															v296 = v294 << (uint(int32(2)) % 32)
															if v296 != 0 {
																v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																base.MemoryCopy(m, v292, v297, v296)
															} else {
															}
															v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
															v305 = v299 + int32(1)
														}
													}
													*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
													*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
													*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
													F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
													mBase = m.M
													v315 = m.ExcPending
													if v315 != 0 {
														return
													} else {
														v322 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
														m.G0 = v10 + int32(16)
														return
													}
												}
											}
										}
									} else {
										v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
										v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
										v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
										v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
										v230 = int32(2)
										v232 = int32(72)
										v237 = v228<<(uint(v230)%32) + v232
										if int32(0) < v227 {
											v240 = (v227+v228)<<(uint(v230)%32) + v232
										} else {
											v240 = v237
										}
										v241 = F_MemoryContextAlloc(m, v224, v240)
										mBase = m.M
										v242 = m.ExcPending
										if v242 != 0 {
											return
										} else {
											v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
											v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
											v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
											v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
											v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
											v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
											v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
											v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
											*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
											*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
											v261 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
											*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
											v265 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
											v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
											if v267 != 0 {
												v269 = v241 + int32(72)
												*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
												v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
												v273 = v271 << (uint(int32(2)) % 32)
												if v273 == int32(0) {
												} else {
													v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
													base.MemoryCopy(m, v269, v276, v273)
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
											}
											v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
											if v283 <= int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
												v305 = int32(1)
											} else {
												v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
												if v286 == int32(1) {
													v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
													if v289 != int32(1) {
														*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
														v305 = int32(1)
													} else {
														v292 = v241 + v237
														*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
														v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
														v296 = v294 << (uint(int32(2)) % 32)
														if v296 != 0 {
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
															base.MemoryCopy(m, v292, v297, v296)
														} else {
														}
														v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
														v305 = v299 + int32(1)
													}
												} else {
													v292 = v241 + v237
													*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
													v296 = v294 << (uint(int32(2)) % 32)
													if v296 != 0 {
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
														base.MemoryCopy(m, v292, v297, v296)
													} else {
													}
													v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
													v305 = v299 + int32(1)
												}
											}
											*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
											*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
											*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
											F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
											mBase = m.M
											v315 = m.ExcPending
											if v315 != 0 {
												return
											} else {
												v322 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
												m.G0 = v10 + int32(16)
												return
											}
										}
									}
								} else {
									v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
									v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
									v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
									v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
									v230 = int32(2)
									v232 = int32(72)
									v237 = v228<<(uint(v230)%32) + v232
									if int32(0) < v227 {
										v240 = (v227+v228)<<(uint(v230)%32) + v232
									} else {
										v240 = v237
									}
									v241 = F_MemoryContextAlloc(m, v224, v240)
									mBase = m.M
									v242 = m.ExcPending
									if v242 != 0 {
										return
									} else {
										v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
										v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
										v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
										v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
										v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
										v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
										v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
										v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
										*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
										*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
										v261 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
										*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
										v265 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
										v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
										if v267 != 0 {
											v269 = v241 + int32(72)
											*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
											v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
											v273 = v271 << (uint(int32(2)) % 32)
											if v273 == int32(0) {
											} else {
												v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
												base.MemoryCopy(m, v269, v276, v273)
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
										}
										v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
										if v283 <= int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
											v305 = int32(1)
										} else {
											v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
											if v286 == int32(1) {
												v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
												if v289 != int32(1) {
													*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
													v305 = int32(1)
												} else {
													v292 = v241 + v237
													*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
													v296 = v294 << (uint(int32(2)) % 32)
													if v296 != 0 {
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
														base.MemoryCopy(m, v292, v297, v296)
													} else {
													}
													v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
													v305 = v299 + int32(1)
												}
											} else {
												v292 = v241 + v237
												*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
												v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
												v296 = v294 << (uint(int32(2)) % 32)
												if v296 != 0 {
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
													base.MemoryCopy(m, v292, v297, v296)
												} else {
												}
												v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
												v305 = v299 + int32(1)
											}
										}
										*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
										*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
										*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
										F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
										mBase = m.M
										v315 = m.ExcPending
										if v315 != 0 {
											return
										} else {
											v322 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
											m.G0 = v10 + int32(16)
											return
										}
									}
								}
							} else {
								v322 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
								m.G0 = v10 + int32(16)
								return
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v160 = m.ExcPending
							if v160 != 0 {
								return
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_SetTransactionSnapshot_5), int32(0))
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return
									} else {
										F_errdetail(m, int32(_a_F_SetTransactionSnapshot_6), int32(0))
										mBase = m.M
										v171 = m.ExcPending
										if v171 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_SetTransactionSnapshot_7), int32(568), int32(_a_F_SetTransactionSnapshot_8))
											mBase = m.M
											v176 = m.ExcPending
											if v176 != 0 {
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
				v177 = F_ProcArrayInstallImportedXmin(m, v98, l1)
				mBase = m.M
				v178 = m.ExcPending
				if v178 != 0 {
					return
				} else {
					if v177 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v330 = m.ExcPending
						if v330 != 0 {
							return
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v333 = m.ExcPending
							if v333 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_SetTransactionSnapshot_5), int32(0))
								mBase = m.M
								v337 = m.ExcPending
								if v337 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
									F_errdetail(m, int32(_a_F_SetTransactionSnapshot_9), v10)
									mBase = m.M
									v341 = m.ExcPending
									if v341 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_SetTransactionSnapshot_7), int32(575), int32(_a_F_SetTransactionSnapshot_8))
										mBase = m.M
										v346 = m.ExcPending
										if v346 != 0 {
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
						v185 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[7]))
						if int32(2) <= v185 {
							if v185 == int32(3) {
								v191 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
								v193 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[8]))
								if v193 < int32(0) {
									v197 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[9])))
									if v197 == int32(1) {
										v201 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[10])))
										if v201&int32(1) != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v209 = m.ExcPending
											if v209 != 0 {
												return
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_SetTransactionSnapshot_1), int32(0))
													mBase = m.M
													v216 = m.ExcPending
													if v216 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_SetTransactionSnapshot_2), int32(1748), int32(_a_F_SetTransactionSnapshot_3))
														mBase = m.M
														v221 = m.ExcPending
														if v221 != 0 {
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
											v204 = F_GetSerializableTransactionSnapshotInt(m, v191, l1, l2)
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
												return
											} else {
												v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
												v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
												v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
												v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
												v230 = int32(2)
												v232 = int32(72)
												v237 = v228<<(uint(v230)%32) + v232
												if int32(0) < v227 {
													v240 = (v227+v228)<<(uint(v230)%32) + v232
												} else {
													v240 = v237
												}
												v241 = F_MemoryContextAlloc(m, v224, v240)
												mBase = m.M
												v242 = m.ExcPending
												if v242 != 0 {
													return
												} else {
													v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
													v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
													v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
													v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
													v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
													v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
													v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
													v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
													*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
													*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
													v261 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
													*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
													v265 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
													v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
													if v267 != 0 {
														v269 = v241 + int32(72)
														*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
														v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
														v273 = v271 << (uint(int32(2)) % 32)
														if v273 == int32(0) {
														} else {
															v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
															base.MemoryCopy(m, v269, v276, v273)
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
													}
													v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
													if v283 <= int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
														v305 = int32(1)
													} else {
														v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
														if v286 == int32(1) {
															v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
															if v289 != int32(1) {
																*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
																v305 = int32(1)
															} else {
																v292 = v241 + v237
																*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
																v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
																v296 = v294 << (uint(int32(2)) % 32)
																if v296 != 0 {
																	v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																	base.MemoryCopy(m, v292, v297, v296)
																} else {
																}
																v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
																v305 = v299 + int32(1)
															}
														} else {
															v292 = v241 + v237
															*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
															v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
															v296 = v294 << (uint(int32(2)) % 32)
															if v296 != 0 {
																v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																base.MemoryCopy(m, v292, v297, v296)
															} else {
															}
															v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
															v305 = v299 + int32(1)
														}
													}
													*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
													*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
													*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
													F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
													mBase = m.M
													v315 = m.ExcPending
													if v315 != 0 {
														return
													} else {
														v322 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
														m.G0 = v10 + int32(16)
														return
													}
												}
											}
										}
									} else {
										v204 = F_GetSerializableTransactionSnapshotInt(m, v191, l1, l2)
										mBase = m.M
										v205 = m.ExcPending
										if v205 != 0 {
											return
										} else {
											v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
											v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
											v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
											v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
											v230 = int32(2)
											v232 = int32(72)
											v237 = v228<<(uint(v230)%32) + v232
											if int32(0) < v227 {
												v240 = (v227+v228)<<(uint(v230)%32) + v232
											} else {
												v240 = v237
											}
											v241 = F_MemoryContextAlloc(m, v224, v240)
											mBase = m.M
											v242 = m.ExcPending
											if v242 != 0 {
												return
											} else {
												v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
												*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
												v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
												*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
												v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
												*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
												v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
												*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
												v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
												*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
												v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
												*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
												v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
												*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
												v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
												*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
												*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
												v261 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
												*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
												v265 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
												v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
												if v267 != 0 {
													v269 = v241 + int32(72)
													*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
													v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
													v273 = v271 << (uint(int32(2)) % 32)
													if v273 == int32(0) {
													} else {
														v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
														base.MemoryCopy(m, v269, v276, v273)
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
												}
												v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
												if v283 <= int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
													v305 = int32(1)
												} else {
													v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
													if v286 == int32(1) {
														v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
														if v289 != int32(1) {
															*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
															v305 = int32(1)
														} else {
															v292 = v241 + v237
															*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
															v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
															v296 = v294 << (uint(int32(2)) % 32)
															if v296 != 0 {
																v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																base.MemoryCopy(m, v292, v297, v296)
															} else {
															}
															v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
															v305 = v299 + int32(1)
														}
													} else {
														v292 = v241 + v237
														*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
														v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
														v296 = v294 << (uint(int32(2)) % 32)
														if v296 != 0 {
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
															base.MemoryCopy(m, v292, v297, v296)
														} else {
														}
														v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
														v305 = v299 + int32(1)
													}
												}
												*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
												*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
												*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
												F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
												mBase = m.M
												v315 = m.ExcPending
												if v315 != 0 {
													return
												} else {
													v322 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
													m.G0 = v10 + int32(16)
													return
												}
											}
										}
									}
								} else {
									v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
									v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
									v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
									v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
									v230 = int32(2)
									v232 = int32(72)
									v237 = v228<<(uint(v230)%32) + v232
									if int32(0) < v227 {
										v240 = (v227+v228)<<(uint(v230)%32) + v232
									} else {
										v240 = v237
									}
									v241 = F_MemoryContextAlloc(m, v224, v240)
									mBase = m.M
									v242 = m.ExcPending
									if v242 != 0 {
										return
									} else {
										v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
										v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
										v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
										v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
										v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
										v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
										v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
										v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
										*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
										*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
										v261 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
										*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
										v265 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
										v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
										if v267 != 0 {
											v269 = v241 + int32(72)
											*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
											v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
											v273 = v271 << (uint(int32(2)) % 32)
											if v273 == int32(0) {
											} else {
												v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
												base.MemoryCopy(m, v269, v276, v273)
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
										}
										v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
										if v283 <= int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
											v305 = int32(1)
										} else {
											v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
											if v286 == int32(1) {
												v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
												if v289 != int32(1) {
													*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
													v305 = int32(1)
												} else {
													v292 = v241 + v237
													*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
													v296 = v294 << (uint(int32(2)) % 32)
													if v296 != 0 {
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
														base.MemoryCopy(m, v292, v297, v296)
													} else {
													}
													v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
													v305 = v299 + int32(1)
												}
											} else {
												v292 = v241 + v237
												*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
												v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
												v296 = v294 << (uint(int32(2)) % 32)
												if v296 != 0 {
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
													base.MemoryCopy(m, v292, v297, v296)
												} else {
												}
												v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
												v305 = v299 + int32(1)
											}
										}
										*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
										*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
										*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
										F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
										mBase = m.M
										v315 = m.ExcPending
										if v315 != 0 {
											return
										} else {
											v322 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
											m.G0 = v10 + int32(16)
											return
										}
									}
								}
							} else {
								v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
								v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
								v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
								v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
								v230 = int32(2)
								v232 = int32(72)
								v237 = v228<<(uint(v230)%32) + v232
								if int32(0) < v227 {
									v240 = (v227+v228)<<(uint(v230)%32) + v232
								} else {
									v240 = v237
								}
								v241 = F_MemoryContextAlloc(m, v224, v240)
								mBase = m.M
								v242 = m.ExcPending
								if v242 != 0 {
									return
								} else {
									v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
									*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
									v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
									*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
									v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
									*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
									v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
									*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
									v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
									*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
									v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
									*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
									v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
									v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
									*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
									*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
									v261 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
									*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
									v265 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
									v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
									if v267 != 0 {
										v269 = v241 + int32(72)
										*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
										v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
										v273 = v271 << (uint(int32(2)) % 32)
										if v273 == int32(0) {
										} else {
											v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
											base.MemoryCopy(m, v269, v276, v273)
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
									}
									v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
									if v283 <= int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
										v305 = int32(1)
									} else {
										v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
										if v286 == int32(1) {
											v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
											if v289 != int32(1) {
												*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
												v305 = int32(1)
											} else {
												v292 = v241 + v237
												*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
												v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
												v296 = v294 << (uint(int32(2)) % 32)
												if v296 != 0 {
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
													base.MemoryCopy(m, v292, v297, v296)
												} else {
												}
												v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
												v305 = v299 + int32(1)
											}
										} else {
											v292 = v241 + v237
											*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
											v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
											v296 = v294 << (uint(int32(2)) % 32)
											if v296 != 0 {
												v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
												base.MemoryCopy(m, v292, v297, v296)
											} else {
											}
											v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
											v305 = v299 + int32(1)
										}
									}
									*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
									*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
									*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
									F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
									mBase = m.M
									v315 = m.ExcPending
									if v315 != 0 {
										return
									} else {
										v322 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
										m.G0 = v10 + int32(16)
										return
									}
								}
							}
						} else {
							v322 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
							m.G0 = v10 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		F_pairingheap_remove(m, int32(_a_F_SetTransactionSnapshot_4), v13+int32(52))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[0])) = int32(0)
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[14]))
			if v25 != 0 {
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[15]))
				if v28 != 0 {
					v30 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[5]))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+40))
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[15]))
					v35 = v33 - int32(48)
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v36))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v31)) == int32(0) {
						v48 = base.B2i32(base.Ui32(v31) < base.Ui32(v36))
					} else {
						v48 = int32(base.Ui32(v31-v36) >> (uint(int32(31)) % 32))
					}
					if v48 == int32(0) {
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
						v52 = v51
						*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[4])) = v52
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[5]))
						*(*int32)(unsafe.Add(mBase, uint32(v56)+40)) = v52
					}
				} else {
					v52 = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[4])) = v52
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[5]))
					*(*int32)(unsafe.Add(mBase, uint32(v56)+40)) = v52
				}
			}
			v61 = F_GetSnapshotData(m, int32(_a_F_SetTransactionSnapshot_0))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v61
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v64
				v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v66
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v68
				if v68 == int32(0) {
				} else {
					v73 = v68 << (uint(int32(2)) % 32)
					if v73 == int32(0) {
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						base.MemoryCopy(m, v76, v77, v73)
					}
				}
				v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = v80
				if v80 <= int32(0) {
				} else {
					v85 = v80 << (uint(int32(2)) % 32)
					if v85 == int32(0) {
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						base.MemoryCopy(m, v88, v89, v85)
					}
				}
				v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
				*(*uint8)(unsafe.Add(mBase, uint32(v61)+28)) = uint8(v92)
				v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
				*(*int64)(unsafe.Add(mBase, uint32(v61)+64)) = int64(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v61)+29)) = uint8(v94)
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
				if l3 != 0 {
					v99 = int32(0)
					v101 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[2]))
					v105 = F_LWLockAcquire(m, v101+int32(512), v99)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return
					} else {
						v107 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
						v109 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[3]))
						v111 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
						if base.B2i32(v107 != v109)|base.B2i32(base.Ui32(v111) < base.Ui32(int32(3))) != 0 {
							v149 = v99
						} else {
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v98))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v111)) == int32(0) {
								v126 = base.B2i32(base.Ui32(v111) <= base.Ui32(v98))
							} else {
								v126 = base.B2i32(v111-v98 <= int32(0))
							}
							if v126 == int32(0) {
								v149 = v99
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[4])) = v98
								v132 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[5]))
								*(*int32)(unsafe.Add(mBase, uint32(v132)+40)) = v98
								v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+124)))
								v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+124)))
								v140 = v134&int32(6) | v137&int32(-7)
								*(*uint8)(unsafe.Add(mBase, uint32(v132)+124)) = uint8(v140)
								v143 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[6]))
								v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
								v145 = *(*int32)(unsafe.Add(mBase, uint32(v132)+48))
								*(*uint8)(unsafe.Add(mBase, uint32(v144+v145))) = uint8(v140)
								v149 = int32(1)
							}
						}
						v152 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[2]))
						F_LWLockRelease(m, v152+int32(512))
						mBase = m.M
						v156 = m.ExcPending
						if v156 != 0 {
							return
						} else {
							if v149 != 0 {
								v185 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[7]))
								if int32(2) <= v185 {
									if v185 == int32(3) {
										v191 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
										v193 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[8]))
										if v193 < int32(0) {
											v197 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[9])))
											if v197 == int32(1) {
												v201 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[10])))
												if v201&int32(1) != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v212 = m.ExcPending
														if v212 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_SetTransactionSnapshot_1), int32(0))
															mBase = m.M
															v216 = m.ExcPending
															if v216 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_SetTransactionSnapshot_2), int32(1748), int32(_a_F_SetTransactionSnapshot_3))
																mBase = m.M
																v221 = m.ExcPending
																if v221 != 0 {
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
													v204 = F_GetSerializableTransactionSnapshotInt(m, v191, l1, l2)
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return
													} else {
														v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
														v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
														v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
														v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
														v230 = int32(2)
														v232 = int32(72)
														v237 = v228<<(uint(v230)%32) + v232
														if int32(0) < v227 {
															v240 = (v227+v228)<<(uint(v230)%32) + v232
														} else {
															v240 = v237
														}
														v241 = F_MemoryContextAlloc(m, v224, v240)
														mBase = m.M
														v242 = m.ExcPending
														if v242 != 0 {
															return
														} else {
															v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
															*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
															v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
															*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
															v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
															*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
															v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
															*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
															v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
															*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
															v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
															*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
															v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
															*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
															v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
															*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
															*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
															v261 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
															*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
															v265 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
															v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
															if v267 != 0 {
																v269 = v241 + int32(72)
																*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
																v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
																v273 = v271 << (uint(int32(2)) % 32)
																if v273 == int32(0) {
																} else {
																	v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
																	base.MemoryCopy(m, v269, v276, v273)
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
															}
															v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
															if v283 <= int32(0) {
																*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
																v305 = int32(1)
															} else {
																v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
																if v286 == int32(1) {
																	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
																	if v289 != int32(1) {
																		*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
																		v305 = int32(1)
																	} else {
																		v292 = v241 + v237
																		*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
																		v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
																		v296 = v294 << (uint(int32(2)) % 32)
																		if v296 != 0 {
																			v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																			base.MemoryCopy(m, v292, v297, v296)
																		} else {
																		}
																		v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
																		v305 = v299 + int32(1)
																	}
																} else {
																	v292 = v241 + v237
																	*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
																	v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
																	v296 = v294 << (uint(int32(2)) % 32)
																	if v296 != 0 {
																		v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																		base.MemoryCopy(m, v292, v297, v296)
																	} else {
																	}
																	v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
																	v305 = v299 + int32(1)
																}
															}
															*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
															*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
															*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
															F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
															mBase = m.M
															v315 = m.ExcPending
															if v315 != 0 {
																return
															} else {
																v322 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
																m.G0 = v10 + int32(16)
																return
															}
														}
													}
												}
											} else {
												v204 = F_GetSerializableTransactionSnapshotInt(m, v191, l1, l2)
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return
												} else {
													v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
													v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
													v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
													v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
													v230 = int32(2)
													v232 = int32(72)
													v237 = v228<<(uint(v230)%32) + v232
													if int32(0) < v227 {
														v240 = (v227+v228)<<(uint(v230)%32) + v232
													} else {
														v240 = v237
													}
													v241 = F_MemoryContextAlloc(m, v224, v240)
													mBase = m.M
													v242 = m.ExcPending
													if v242 != 0 {
														return
													} else {
														v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
														v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
														v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
														v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
														v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
														v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
														v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
														v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
														*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
														*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
														v261 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
														*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
														v265 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
														v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
														if v267 != 0 {
															v269 = v241 + int32(72)
															*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
															v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
															v273 = v271 << (uint(int32(2)) % 32)
															if v273 == int32(0) {
															} else {
																v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
																base.MemoryCopy(m, v269, v276, v273)
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
														}
														v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
														if v283 <= int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
															v305 = int32(1)
														} else {
															v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
															if v286 == int32(1) {
																v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
																if v289 != int32(1) {
																	*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
																	v305 = int32(1)
																} else {
																	v292 = v241 + v237
																	*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
																	v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
																	v296 = v294 << (uint(int32(2)) % 32)
																	if v296 != 0 {
																		v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																		base.MemoryCopy(m, v292, v297, v296)
																	} else {
																	}
																	v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
																	v305 = v299 + int32(1)
																}
															} else {
																v292 = v241 + v237
																*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
																v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
																v296 = v294 << (uint(int32(2)) % 32)
																if v296 != 0 {
																	v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																	base.MemoryCopy(m, v292, v297, v296)
																} else {
																}
																v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
																v305 = v299 + int32(1)
															}
														}
														*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
														*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
														*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
														F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
														mBase = m.M
														v315 = m.ExcPending
														if v315 != 0 {
															return
														} else {
															v322 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
															m.G0 = v10 + int32(16)
															return
														}
													}
												}
											}
										} else {
											v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
											v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
											v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
											v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
											v230 = int32(2)
											v232 = int32(72)
											v237 = v228<<(uint(v230)%32) + v232
											if int32(0) < v227 {
												v240 = (v227+v228)<<(uint(v230)%32) + v232
											} else {
												v240 = v237
											}
											v241 = F_MemoryContextAlloc(m, v224, v240)
											mBase = m.M
											v242 = m.ExcPending
											if v242 != 0 {
												return
											} else {
												v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
												*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
												v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
												*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
												v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
												*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
												v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
												*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
												v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
												*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
												v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
												*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
												v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
												*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
												v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
												*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
												*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
												v261 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
												*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
												v265 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
												v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
												if v267 != 0 {
													v269 = v241 + int32(72)
													*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
													v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
													v273 = v271 << (uint(int32(2)) % 32)
													if v273 == int32(0) {
													} else {
														v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
														base.MemoryCopy(m, v269, v276, v273)
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
												}
												v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
												if v283 <= int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
													v305 = int32(1)
												} else {
													v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
													if v286 == int32(1) {
														v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
														if v289 != int32(1) {
															*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
															v305 = int32(1)
														} else {
															v292 = v241 + v237
															*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
															v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
															v296 = v294 << (uint(int32(2)) % 32)
															if v296 != 0 {
																v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																base.MemoryCopy(m, v292, v297, v296)
															} else {
															}
															v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
															v305 = v299 + int32(1)
														}
													} else {
														v292 = v241 + v237
														*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
														v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
														v296 = v294 << (uint(int32(2)) % 32)
														if v296 != 0 {
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
															base.MemoryCopy(m, v292, v297, v296)
														} else {
														}
														v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
														v305 = v299 + int32(1)
													}
												}
												*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
												*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
												*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
												F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
												mBase = m.M
												v315 = m.ExcPending
												if v315 != 0 {
													return
												} else {
													v322 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
													m.G0 = v10 + int32(16)
													return
												}
											}
										}
									} else {
										v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
										v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
										v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
										v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
										v230 = int32(2)
										v232 = int32(72)
										v237 = v228<<(uint(v230)%32) + v232
										if int32(0) < v227 {
											v240 = (v227+v228)<<(uint(v230)%32) + v232
										} else {
											v240 = v237
										}
										v241 = F_MemoryContextAlloc(m, v224, v240)
										mBase = m.M
										v242 = m.ExcPending
										if v242 != 0 {
											return
										} else {
											v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
											v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
											v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
											v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
											v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
											v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
											v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
											v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
											*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
											*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
											v261 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
											*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
											v265 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
											v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
											if v267 != 0 {
												v269 = v241 + int32(72)
												*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
												v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
												v273 = v271 << (uint(int32(2)) % 32)
												if v273 == int32(0) {
												} else {
													v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
													base.MemoryCopy(m, v269, v276, v273)
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
											}
											v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
											if v283 <= int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
												v305 = int32(1)
											} else {
												v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
												if v286 == int32(1) {
													v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
													if v289 != int32(1) {
														*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
														v305 = int32(1)
													} else {
														v292 = v241 + v237
														*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
														v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
														v296 = v294 << (uint(int32(2)) % 32)
														if v296 != 0 {
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
															base.MemoryCopy(m, v292, v297, v296)
														} else {
														}
														v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
														v305 = v299 + int32(1)
													}
												} else {
													v292 = v241 + v237
													*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
													v296 = v294 << (uint(int32(2)) % 32)
													if v296 != 0 {
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
														base.MemoryCopy(m, v292, v297, v296)
													} else {
													}
													v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
													v305 = v299 + int32(1)
												}
											}
											*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
											*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
											*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
											F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
											mBase = m.M
											v315 = m.ExcPending
											if v315 != 0 {
												return
											} else {
												v322 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
												m.G0 = v10 + int32(16)
												return
											}
										}
									}
								} else {
									v322 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
									m.G0 = v10 + int32(16)
									return
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
									return
								} else {
									F_errcode(m, int32(325))
									mBase = m.M
									v163 = m.ExcPending
									if v163 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_SetTransactionSnapshot_5), int32(0))
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return
										} else {
											F_errdetail(m, int32(_a_F_SetTransactionSnapshot_6), int32(0))
											mBase = m.M
											v171 = m.ExcPending
											if v171 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_SetTransactionSnapshot_7), int32(568), int32(_a_F_SetTransactionSnapshot_8))
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
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
					v177 = F_ProcArrayInstallImportedXmin(m, v98, l1)
					mBase = m.M
					v178 = m.ExcPending
					if v178 != 0 {
						return
					} else {
						if v177 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v330 = m.ExcPending
							if v330 != 0 {
								return
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v333 = m.ExcPending
								if v333 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_SetTransactionSnapshot_5), int32(0))
									mBase = m.M
									v337 = m.ExcPending
									if v337 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
										F_errdetail(m, int32(_a_F_SetTransactionSnapshot_9), v10)
										mBase = m.M
										v341 = m.ExcPending
										if v341 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_SetTransactionSnapshot_7), int32(575), int32(_a_F_SetTransactionSnapshot_8))
											mBase = m.M
											v346 = m.ExcPending
											if v346 != 0 {
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
							v185 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[7]))
							if int32(2) <= v185 {
								if v185 == int32(3) {
									v191 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
									v193 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[8]))
									if v193 < int32(0) {
										v197 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[9])))
										if v197 == int32(1) {
											v201 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[10])))
											if v201&int32(1) != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v209 = m.ExcPending
												if v209 != 0 {
													return
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_SetTransactionSnapshot_1), int32(0))
														mBase = m.M
														v216 = m.ExcPending
														if v216 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_SetTransactionSnapshot_2), int32(1748), int32(_a_F_SetTransactionSnapshot_3))
															mBase = m.M
															v221 = m.ExcPending
															if v221 != 0 {
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
												v204 = F_GetSerializableTransactionSnapshotInt(m, v191, l1, l2)
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return
												} else {
													v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
													v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
													v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
													v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
													v230 = int32(2)
													v232 = int32(72)
													v237 = v228<<(uint(v230)%32) + v232
													if int32(0) < v227 {
														v240 = (v227+v228)<<(uint(v230)%32) + v232
													} else {
														v240 = v237
													}
													v241 = F_MemoryContextAlloc(m, v224, v240)
													mBase = m.M
													v242 = m.ExcPending
													if v242 != 0 {
														return
													} else {
														v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
														v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
														v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
														v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
														v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
														v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
														v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
														v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
														*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
														*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
														v261 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
														*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
														v265 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
														v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
														if v267 != 0 {
															v269 = v241 + int32(72)
															*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
															v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
															v273 = v271 << (uint(int32(2)) % 32)
															if v273 == int32(0) {
															} else {
																v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
																base.MemoryCopy(m, v269, v276, v273)
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
														}
														v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
														if v283 <= int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
															v305 = int32(1)
														} else {
															v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
															if v286 == int32(1) {
																v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
																if v289 != int32(1) {
																	*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
																	v305 = int32(1)
																} else {
																	v292 = v241 + v237
																	*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
																	v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
																	v296 = v294 << (uint(int32(2)) % 32)
																	if v296 != 0 {
																		v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																		base.MemoryCopy(m, v292, v297, v296)
																	} else {
																	}
																	v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
																	v305 = v299 + int32(1)
																}
															} else {
																v292 = v241 + v237
																*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
																v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
																v296 = v294 << (uint(int32(2)) % 32)
																if v296 != 0 {
																	v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																	base.MemoryCopy(m, v292, v297, v296)
																} else {
																}
																v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
																v305 = v299 + int32(1)
															}
														}
														*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
														*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
														*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
														F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
														mBase = m.M
														v315 = m.ExcPending
														if v315 != 0 {
															return
														} else {
															v322 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
															m.G0 = v10 + int32(16)
															return
														}
													}
												}
											}
										} else {
											v204 = F_GetSerializableTransactionSnapshotInt(m, v191, l1, l2)
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
												return
											} else {
												v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
												v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
												v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
												v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
												v230 = int32(2)
												v232 = int32(72)
												v237 = v228<<(uint(v230)%32) + v232
												if int32(0) < v227 {
													v240 = (v227+v228)<<(uint(v230)%32) + v232
												} else {
													v240 = v237
												}
												v241 = F_MemoryContextAlloc(m, v224, v240)
												mBase = m.M
												v242 = m.ExcPending
												if v242 != 0 {
													return
												} else {
													v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
													v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
													v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
													v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
													v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
													v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
													v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
													v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
													*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
													*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
													v261 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
													*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
													v265 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
													v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
													if v267 != 0 {
														v269 = v241 + int32(72)
														*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
														v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
														v273 = v271 << (uint(int32(2)) % 32)
														if v273 == int32(0) {
														} else {
															v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
															base.MemoryCopy(m, v269, v276, v273)
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
													}
													v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
													if v283 <= int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
														v305 = int32(1)
													} else {
														v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
														if v286 == int32(1) {
															v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
															if v289 != int32(1) {
																*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
																v305 = int32(1)
															} else {
																v292 = v241 + v237
																*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
																v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
																v296 = v294 << (uint(int32(2)) % 32)
																if v296 != 0 {
																	v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																	base.MemoryCopy(m, v292, v297, v296)
																} else {
																}
																v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
																v305 = v299 + int32(1)
															}
														} else {
															v292 = v241 + v237
															*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
															v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
															v296 = v294 << (uint(int32(2)) % 32)
															if v296 != 0 {
																v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
																base.MemoryCopy(m, v292, v297, v296)
															} else {
															}
															v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
															v305 = v299 + int32(1)
														}
													}
													*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
													*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
													*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
													F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
													mBase = m.M
													v315 = m.ExcPending
													if v315 != 0 {
														return
													} else {
														v322 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
														m.G0 = v10 + int32(16)
														return
													}
												}
											}
										}
									} else {
										v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
										v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
										v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
										v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
										v230 = int32(2)
										v232 = int32(72)
										v237 = v228<<(uint(v230)%32) + v232
										if int32(0) < v227 {
											v240 = (v227+v228)<<(uint(v230)%32) + v232
										} else {
											v240 = v237
										}
										v241 = F_MemoryContextAlloc(m, v224, v240)
										mBase = m.M
										v242 = m.ExcPending
										if v242 != 0 {
											return
										} else {
											v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
											v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
											v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
											v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
											v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
											v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
											v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
											v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
											*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
											*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
											v261 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
											*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
											v265 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
											v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
											if v267 != 0 {
												v269 = v241 + int32(72)
												*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
												v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
												v273 = v271 << (uint(int32(2)) % 32)
												if v273 == int32(0) {
												} else {
													v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
													base.MemoryCopy(m, v269, v276, v273)
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
											}
											v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
											if v283 <= int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
												v305 = int32(1)
											} else {
												v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
												if v286 == int32(1) {
													v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
													if v289 != int32(1) {
														*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
														v305 = int32(1)
													} else {
														v292 = v241 + v237
														*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
														v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
														v296 = v294 << (uint(int32(2)) % 32)
														if v296 != 0 {
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
															base.MemoryCopy(m, v292, v297, v296)
														} else {
														}
														v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
														v305 = v299 + int32(1)
													}
												} else {
													v292 = v241 + v237
													*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
													v296 = v294 << (uint(int32(2)) % 32)
													if v296 != 0 {
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
														base.MemoryCopy(m, v292, v297, v296)
													} else {
													}
													v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
													v305 = v299 + int32(1)
												}
											}
											*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
											*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
											*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
											F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
											mBase = m.M
											v315 = m.ExcPending
											if v315 != 0 {
												return
											} else {
												v322 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
												m.G0 = v10 + int32(16)
												return
											}
										}
									}
								} else {
									v224 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[11]))
									v226 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1]))
									v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
									v228 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
									v230 = int32(2)
									v232 = int32(72)
									v237 = v228<<(uint(v230)%32) + v232
									if int32(0) < v227 {
										v240 = (v227+v228)<<(uint(v230)%32) + v232
									} else {
										v240 = v237
									}
									v241 = F_MemoryContextAlloc(m, v224, v240)
									mBase = m.M
									v242 = m.ExcPending
									if v242 != 0 {
										return
									} else {
										v243 = *(*int64)(unsafe.Add(mBase, uint32(v226)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+48)) = v243
										v245 = *(*int64)(unsafe.Add(mBase, uint32(v226)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+40)) = v245
										v247 = *(*int64)(unsafe.Add(mBase, uint32(v226)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+24)) = v247
										v249 = *(*int64)(unsafe.Add(mBase, uint32(v226)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+56)) = v249
										v251 = *(*int64)(unsafe.Add(mBase, uint32(v226)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+32)) = v251
										v253 = *(*int64)(unsafe.Add(mBase, uint32(v226)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+16)) = v253
										v255 = *(*int64)(unsafe.Add(mBase, uint32(v226)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v255
										v257 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
										*(*int64)(unsafe.Add(mBase, uint32(v241))) = v257
										*(*int64)(unsafe.Add(mBase, uint32(v241)+64)) = int64(0)
										v261 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v261
										*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v261
										v265 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v241)+30)) = uint8(v265)
										v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
										if v267 != 0 {
											v269 = v241 + int32(72)
											*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v269
											v271 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
											v273 = v271 << (uint(int32(2)) % 32)
											if v273 == int32(0) {
											} else {
												v276 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
												base.MemoryCopy(m, v269, v276, v273)
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(0)
										}
										v283 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
										if v283 <= int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
											v305 = int32(1)
										} else {
											v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+28)))
											if v286 == int32(1) {
												v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+29)))
												if v289 != int32(1) {
													*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(0)
													v305 = int32(1)
												} else {
													v292 = v241 + v237
													*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
													v296 = v294 << (uint(int32(2)) % 32)
													if v296 != 0 {
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
														base.MemoryCopy(m, v292, v297, v296)
													} else {
													}
													v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
													v305 = v299 + int32(1)
												}
											} else {
												v292 = v241 + v237
												*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v292
												v294 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
												v296 = v294 << (uint(int32(2)) % 32)
												if v296 != 0 {
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
													base.MemoryCopy(m, v292, v297, v296)
												} else {
												}
												v299 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
												v305 = v299 + int32(1)
											}
										}
										*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[12])) = v241
										*(*int32)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[1])) = v241
										*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v305
										F_pairingheap_add(m, int32(_a_F_SetTransactionSnapshot_4), v241+int32(52))
										mBase = m.M
										v315 = m.ExcPending
										if v315 != 0 {
											return
										} else {
											v322 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
											m.G0 = v10 + int32(16)
											return
										}
									}
								}
							} else {
								v322 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionSnapshot[13])) = uint8(v322)
								m.G0 = v10 + int32(16)
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
func F_TransactionIdPrecedesOrEquals(m *base.Module, l0 int32, l1 int32) int32 {
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		return base.B2i32(base.Ui32(l0) <= base.Ui32(l1))
	} else {
		return base.B2i32(l0-l1 <= int32(0))
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
