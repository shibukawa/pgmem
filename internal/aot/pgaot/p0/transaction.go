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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[37]))
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
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v34<<(uint(int32(2))%32))+uint32(_consts[157])))
				v44 = v43
			} else {
				v44 = int32(542838)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v44
			F_errmsg_internal(m, int32(187100), v5)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				F_errfinish(m, int32(492874), int32(3974), int32(316660))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
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
					F_errmsg(m, int32(128002), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						F_errfinish(m, int32(492874), int32(3956), int32(316660))
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int64
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int64
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
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
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _consts[111])))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _consts[156])))
	v24 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	v26 = *(*int32)(unsafe.Add(mBase, _consts[37]))
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
	v239 = m.ExcPending
	if v239 != 0 {
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
	v237 = m.ExcPending
	if v237 != 0 {
		goto L19
	} else {
		goto L70
	}
L6:
	;
	F_AbortSubTransaction(m)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L19
	} else {
		goto L69
	}
L7:
	;
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v26)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = int32(0)
	F_CleanupSubTransaction(m)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L19
	} else {
		goto L66
	}
L8:
	;
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v26)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = int32(0)
	F_AbortSubTransaction(m)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
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
	v120 = m.ExcPending
	if v120 != 0 {
		goto L19
	} else {
		goto L39
	}
L12:
	;
	F_PrepareTransaction(m)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L19
	} else {
		goto L38
	}
L13:
	;
	F_AbortTransaction(m)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L19
	} else {
		goto L34
	}
L14:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L19
	} else {
		goto L31
	}
L15:
	;
	F_CommitTransaction(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L19
	} else {
		goto L28
	}
L16:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v42
	F_errmsg_internal(m, int32(187229), v10)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L19
	} else {
		goto L25
	}
L22:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_consts[157])))
	v42 = v41
	goto L24
L23:
	;
	v42 = int32(542838)
	goto L24
L24:
	;
	goto L21
L25:
	;
	F_errfinish(m, int32(492874), int32(3194), int32(312521))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
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
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+77)))
	if v60 != int32(1) {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_StartTransaction(m)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	v65 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+77)) = uint8(v65)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(3)
	*(*uint8)(unsafe.Add(mBase, _consts[156])) = uint8(v22)
	*(*int32)(unsafe.Add(mBase, _consts[110])) = v24
	*(*uint8)(unsafe.Add(mBase, _consts[111])) = uint8(v20)
	goto L1
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(0)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+77)))
	if v79 != int32(1) {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_StartTransaction(m)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	v84 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+77)) = uint8(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(3)
	*(*uint8)(unsafe.Add(mBase, _consts[156])) = uint8(v22)
	*(*int32)(unsafe.Add(mBase, _consts[110])) = v24
	*(*uint8)(unsafe.Add(mBase, _consts[111])) = uint8(v20)
	goto L1
L34:
	;
	F_CleanupTransaction(m)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(0)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+77)))
	if v100 != int32(1) {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_StartTransaction(m)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	v105 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+77)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(3)
	*(*uint8)(unsafe.Add(mBase, _consts[156])) = uint8(v22)
	*(*int32)(unsafe.Add(mBase, _consts[110])) = v24
	*(*uint8)(unsafe.Add(mBase, _consts[111])) = uint8(v20)
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
	v131 = m.ExcPending
	if v131 != 0 {
		goto L19
	} else {
		goto L42
	}
L41:
	;
	goto L1
L42:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+24))
	if v134 == int32(13) {
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
	v145 = m.ExcPending
	if v145 != 0 {
		goto L19
	} else {
		goto L46
	}
L45:
	;
	switch v148 - int32(6) {
	case 0:
		goto L50
	default:
		goto L48
	case 4:
		goto L49
	}
L46:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+24))
	if v148 == int32(14) {
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
	v179 = m.ExcPending
	if v179 != 0 {
		goto L19
	} else {
		goto L55
	}
L49:
	;
	F_PrepareTransaction(m)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L19
	} else {
		goto L54
	}
L50:
	;
	F_CommitTransaction(m)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L19
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+24)) = int32(0)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+77)))
	if v157 != int32(1) {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_StartTransaction(m)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L19
	} else {
		goto L53
	}
L53:
	;
	v162 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v147)+77)) = uint8(v162)
	*(*int32)(unsafe.Add(mBase, uint32(v147)+24)) = int32(3)
	*(*uint8)(unsafe.Add(mBase, _consts[156])) = uint8(v22)
	*(*int32)(unsafe.Add(mBase, _consts[110])) = v24
	*(*uint8)(unsafe.Add(mBase, _consts[111])) = uint8(v20)
	goto L1
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+24)) = int32(0)
	goto L1
L55:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v147)+24))
	if base.Ui32(v180) <= base.Ui32(int32(19)) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v190
	F_errmsg_internal(m, int32(187229), v10+int32(16))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L19
	} else {
		goto L60
	}
L57:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v180<<(uint(int32(2))%32))+uint32(_consts[157])))
	v190 = v189
	goto L59
L58:
	;
	v190 = int32(542838)
	goto L59
L59:
	;
	goto L56
L60:
	;
	F_errfinish(m, int32(492874), int32(3360), int32(312521))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
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
	v208 = m.ExcPending
	if v208 != 0 {
		goto L19
	} else {
		goto L63
	}
L63:
	;
	F_DefineSavepoint(m, int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L19
	} else {
		goto L64
	}
L64:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	*(*int64)(unsafe.Add(mBase, uint32(v213)+12)) = v202
	F_StartSubTransaction(m)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L19
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+24)) = int32(12)
	goto L1
L66:
	;
	F_DefineSavepoint(m, int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L19
	} else {
		goto L67
	}
L67:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	*(*int64)(unsafe.Add(mBase, uint32(v228)+12)) = v219
	F_StartSubTransaction(m)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L19
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228)+24)) = int32(12)
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
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
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
	var v138 int32
	_ = v138
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
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int64
	_ = v299
	var v301 int64
	_ = v301
	var v303 int64
	_ = v303
	var v305 int64
	_ = v305
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	v8 = *(*int32)(unsafe.Add(mBase, _consts[529]))
	if v8 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L10
	} else {
		goto L128
	}
L2:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[153])))
	if v12 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v446 = v8
	goto L4
L4:
	;
	return v446
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	if v16 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if int32(2) <= v388 {
		goto L112
	} else {
		goto L113
	}
L8:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+72))
	if v69 != 0 {
		goto L22
	} else {
		goto L23
	}
L9:
	;
	F_pairingheap_remove(m, int32(4156648), v16+int32(52))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[42])) = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v30 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[41]))
	if v33 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+40))
	v38 = *(*int32)(unsafe.Add(mBase, _consts[41]))
	v40 = v38 - int32(48)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v41))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v36)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v57 = int32(0)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[146])) = v57
	v61 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+40)) = v57
	goto L8
L16:
	;
	if v53 == int32(0) {
		goto L8
	} else {
		goto L20
	}
L17:
	;
	v53 = base.B2i32(base.Ui32(v36) < base.Ui32(v41))
	goto L16
L18:
	;
	goto L19
L19:
	;
	v53 = int32(base.Ui32(v36-v41) >> (uint(int32(31)) % 32))
	goto L16
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v57 = v56
	goto L15
L21:
	;
	if v71&int32(1) != 0 {
		goto L1
	} else {
		goto L25
	}
L22:
	;
	v71 = int32(1)
	goto L24
L23:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+76)))
	v71 = v70
	goto L24
L24:
	;
	goto L21
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if int32(2) <= v75 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v384 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v384)
	return v377
L27:
	;
	if v75 == int32(3) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L29
L29:
	;
	v374 = F_GetSnapshotData(m, int32(4489568))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L10
	} else {
		goto L111
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _consts[154])) = v273
	v276 = *(*int32)(unsafe.Add(mBase, _consts[76]))
	v278 = v273 + int32(24)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	v281 = v273 + int32(16)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	v284 = int32(2)
	v286 = int32(72)
	v291 = v282<<(uint(v284)%32) + v286
	if int32(0) < v279 {
		goto L87
	} else {
		goto L88
	}
L31:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
	if v83 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	goto L33
L33:
	;
	v265 = F_GetSnapshotData(m, int32(4489568))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L10
	} else {
		goto L86
	}
L34:
	;
	v273 = v263
	goto L30
L35:
	;
	if v93 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+316))
	v91 = base.B2i32(v89 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v91)
	v93 = v91
	goto L38
L37:
	;
	v93 = int32(0)
	goto L38
L38:
	;
	goto L35
L39:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, _consts[156])))
	if v97 != int32(1) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	goto L41
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L10
	} else {
		goto L80
	}
L42:
	;
	v263 = v229
	goto L34
L43:
	;
	v225 = F_GetSerializableTransactionSnapshotInt(m, int32(4489568), int32(0), int32(-1))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L10
	} else {
		goto L79
	}
L44:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, _consts[111])))
	if v101 != int32(1) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v107 = F_GetSerializableTransactionSnapshotInt(m, int32(4489568), int32(0), int32(-1))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	if v110 == int32(0) {
		v229 = v107
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v115 = v107
	goto L48
L48:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v124 = F_LWLockAcquire(m, v120+int32(3584), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L10
	} else {
		goto L50
	}
L49:
	;
	v229 = v218
	goto L42
L50:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+108))
	v130 = v128 | int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(v127)+108)) = v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)+92))
	if v132 == int32(0) {
		v172 = v127
		v173 = v130
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+108)) = v173 & int32(-65)
	v182 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v182+int32(3584))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L10
	} else {
		goto L64
	}
L52:
	;
	if v132 == v127+int32(88) {
		v172 = v127
		v173 = v130
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v138 = v127
	goto L54
L54:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v138)+108))
	if v144&int32(256) != 0 {
		v172 = v138
		v173 = v144
		goto L51
	} else {
		goto L56
	}
L55:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v164)+108))
	v172 = v164
	v173 = v171
	goto L51
L56:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v148+int32(3584))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L10
	} else {
		goto L57
	}
L57:
	;
	F_ProcWaitForSignal(m, int32(134217779))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v161 = F_LWLockAcquire(m, v157+int32(3584), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+92))
	if v165 != v164+int32(88) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v170 = v165
	goto L62
L61:
	;
	v170 = int32(0)
	goto L62
L62:
	;
	if v170 != 0 {
		v138 = v164
		goto L54
	} else {
		goto L63
	}
L63:
	;
	goto L55
L64:
	;
	if v173&int32(256) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	F_ReleasePredicateLocks(m, int32(0), int32(1))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L10
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v197 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L10
	} else {
		goto L69
	}
L68:
	;
	v263 = v115
	goto L34
L69:
	;
	if v197 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L10
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v211 = int32(0)
	F_ReleasePredicateLocks(m, v211, v211)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L10
	} else {
		goto L76
	}
L73:
	;
	F_errmsg_internal(m, int32(372350), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(498059), int32(1605), int32(87033))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	v218 = F_GetSerializableTransactionSnapshotInt(m, int32(4489568), int32(0), int32(-1))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L10
	} else {
		goto L77
	}
L77:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	if v221 != 0 {
		v115 = v218
		goto L48
	} else {
		goto L78
	}
L78:
	;
	goto L49
L79:
	;
	v229 = v225
	goto L42
L80:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L10
	} else {
		goto L81
	}
L81:
	;
	F_errmsg(m, int32(23463), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L10
	} else {
		goto L82
	}
L82:
	;
	F_errdetail(m, int32(652978), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	F_errhint(m, int32(565385), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(498059), int32(1697), int32(86948))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L10
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v273 = v265
	goto L30
L87:
	;
	v294 = (v279+v282)<<(uint(v284)%32) + v286
	goto L89
L88:
	;
	v294 = v291
	goto L89
L89:
	;
	v295 = F_MemoryContextAlloc(m, v276, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	v298 = v295 + int32(48)
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v273)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v298))) = v299
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v273)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v295)+40)) = v301
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v278)))
	*(*int64)(unsafe.Add(mBase, uint32(v295)+24)) = v303
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v273)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v295)+56)) = v305
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v273)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v295)+32)) = v307
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v281)))
	*(*int64)(unsafe.Add(mBase, uint32(v295)+16)) = v309
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v273)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v295)+8)) = v311
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v273)))
	*(*int64)(unsafe.Add(mBase, uint32(v295))) = v313
	*(*int64)(unsafe.Add(mBase, uint32(v295)+64)) = int64(0)
	v317 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v298))) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v295)+44)) = v317
	v321 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v295)+30)) = uint8(v321)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	if v323 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v273)+24))
	if v337 <= int32(0) {
		goto L100
	} else {
		goto L101
	}
L92:
	;
	v325 = v295 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = v325
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v273)+16))
	v330 = v328 << (uint(int32(2)) % 32)
	if v330 != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = int32(0)
	goto L91
L95:
	;
	goto L91
L96:
	;
	v331 = F__emscripten_memcpy_bulkmem(m, v325, v327, v330)
	mBase = m.M
	goto L98
L97:
	;
	goto L98
L98:
	;
	goto L95
L99:
	;
	*(*int32)(unsafe.Add(mBase, _consts[988])) = v295
	*(*int32)(unsafe.Add(mBase, _consts[154])) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v295)+48)) = v359
	F_pairingheap_add(m, int32(4156648), v295+int32(52))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L10
	} else {
		goto L110
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+20)) = int32(0)
	v359 = int32(1)
	goto L99
L101:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+28)))
	if v340 == int32(1) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+29)))
	if v343 != int32(1) {
		goto L100
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v346 = v295 + v291
	*(*int32)(unsafe.Add(mBase, uint32(v295)+20)) = v346
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v273)+20))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v273)+24))
	v351 = v349 << (uint(int32(2)) % 32)
	if v351 != 0 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L104
L106:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v295)+48))
	v359 = v354 + int32(1)
	goto L99
L107:
	;
	v352 = F__emscripten_memcpy_bulkmem(m, v346, v348, v351)
	mBase = m.M
	goto L109
L108:
	;
	goto L109
L109:
	;
	goto L106
L110:
	;
	v371 = *(*int32)(unsafe.Add(mBase, _consts[154]))
	v377 = v371
	goto L26
L111:
	;
	*(*int32)(unsafe.Add(mBase, _consts[154])) = v374
	v377 = v374
	goto L26
L112:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _consts[154]))
	return v392
L113:
	;
	goto L114
L114:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	if v395 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v443 = F_GetSnapshotData(m, int32(4489568))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L10
	} else {
		goto L127
	}
L116:
	;
	F_pairingheap_remove(m, int32(4156648), v395+int32(52))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L10
	} else {
		goto L117
	}
L117:
	;
	v403 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[42])) = v403
	v408 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v408 != 0 {
		goto L115
	} else {
		goto L118
	}
L118:
	;
	v410 = *(*int32)(unsafe.Add(mBase, _consts[41]))
	if v410 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v412 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)+40))
	v415 = *(*int32)(unsafe.Add(mBase, _consts[41]))
	v417 = v415 - int32(48)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v418))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v413)) == int32(0) {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	v434 = v403
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, _consts[146])) = v434
	v438 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v438)+40)) = v434
	goto L115
L122:
	;
	if v430 == int32(0) {
		goto L115
	} else {
		goto L126
	}
L123:
	;
	v430 = base.B2i32(base.Ui32(v413) < base.Ui32(v418))
	goto L122
L124:
	;
	goto L125
L125:
	;
	v430 = int32(base.Ui32(v413-v418) >> (uint(int32(31)) % 32))
	goto L122
L126:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	v434 = v433
	goto L121
L127:
	;
	*(*int32)(unsafe.Add(mBase, _consts[154])) = v443
	v446 = v443
	goto L4
L128:
	;
	F_errmsg_internal(m, int32(259407), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L10
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(494533), int32(306), int32(86890))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L10
	} else {
		goto L130
	}
L130:
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
	v2 = *(*int32)(unsafe.Add(mBase, _consts[37]))
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
	v7 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v9 = *(*int32)(unsafe.Add(mBase, _consts[76]))
	v11 = F_MemoryContextAllocZero(m, v9, int32(88))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = int32(1)
		v14 = int32(4384932)
		v16 = *(*int32)(unsafe.Add(mBase, _consts[158]))
		v18 = v16 + v13
		*(*int32)(unsafe.Add(mBase, _consts[158])) = v18
		if v18 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v7
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(0)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
			v25 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v24 + v25
			v29 = int32(4487480)
			v31 = *(*int32)(unsafe.Add(mBase, _consts[159]))
			v33 = v31 + v25
			*(*int32)(unsafe.Add(mBase, _consts[159])) = v33
			*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v33
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v11)+20)) = int64(47244640256)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v36
			v45 = *(*int32)(unsafe.Add(mBase, _consts[4]))
			*(*int32)(unsafe.Add(mBase, uint32(v11+int32(60)))) = v45
			v48 = *(*int32)(unsafe.Add(mBase, _consts[122]))
			*(*int32)(unsafe.Add(mBase, uint32(v11-int32(-64)))) = v48
			v51 = int32(*(*uint8)(unsafe.Add(mBase, _consts[156])))
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
			*(*int32)(unsafe.Add(mBase, _consts[37])) = v11
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[158])) = v16
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
						F_errmsg(m, int32(257254), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							F_errfinish(m, int32(492874), int32(5438), int32(257740))
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
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
	var v254 int64
	_ = v254
	var v256 int64
	_ = v256
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[42]))
	if v13 == int32(0) {
		v61 = F_GetSnapshotData(m, int32(4489568))
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[154])) = v61
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v64
			v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v66
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v68
			if v68 != 0 {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v73 = v68 << (uint(int32(2)) % 32)
				if v73 != 0 {
					v74 = F__emscripten_memcpy_bulkmem(m, v70, v71, v73)
					mBase = m.M
				} else {
				}
			} else {
			}
			v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = v76
			if int32(0) < v76 {
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
				v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v83 = v76 << (uint(int32(2)) % 32)
				if v83 != 0 {
					v84 = F__emscripten_memcpy_bulkmem(m, v80, v81, v83)
					mBase = m.M
				} else {
				}
			} else {
			}
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
			*(*uint8)(unsafe.Add(mBase, uint32(v61)+28)) = uint8(v86)
			v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
			*(*int64)(unsafe.Add(mBase, uint32(v61)+64)) = int64(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v61)+29)) = uint8(v88)
			v92 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
			if l3 != 0 {
				v93 = int32(0)
				v95 = *(*int32)(unsafe.Add(mBase, _consts[29]))
				v99 = F_LWLockAcquire(m, v95+int32(512), v93)
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return
				} else {
					v101 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
					v102 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
					v104 = *(*int32)(unsafe.Add(mBase, _consts[226]))
					if v102 != v104 {
						v142 = v93
					} else {
						if base.Ui32(v101) < base.Ui32(int32(3)) {
							v142 = v93
						} else {
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v92))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v101)) == int32(0) {
								v119 = base.B2i32(base.Ui32(v101) <= base.Ui32(v92))
							} else {
								v119 = base.B2i32(v101-v92 <= int32(0))
							}
							if v119 == int32(0) {
								v142 = v93
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[146])) = v92
								v125 = *(*int32)(unsafe.Add(mBase, _consts[91]))
								*(*int32)(unsafe.Add(mBase, uint32(v125)+40)) = v92
								v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+124)))
								v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+124)))
								v133 = v127&int32(6) | v130&int32(-7)
								*(*uint8)(unsafe.Add(mBase, uint32(v125)+124)) = uint8(v133)
								v136 = *(*int32)(unsafe.Add(mBase, _consts[94]))
								v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
								v138 = *(*int32)(unsafe.Add(mBase, uint32(v125)+48))
								*(*uint8)(unsafe.Add(mBase, uint32(v137+v138))) = uint8(v133)
								v142 = int32(1)
							}
						}
					}
					v145 = *(*int32)(unsafe.Add(mBase, _consts[29]))
					F_LWLockRelease(m, v145+int32(512))
					mBase = m.M
					v149 = m.ExcPending
					if v149 != 0 {
						return
					} else {
						if v142 != 0 {
							v178 = *(*int32)(unsafe.Add(mBase, _consts[110]))
							if int32(2) <= v178 {
								if v178 == int32(3) {
									v184 = *(*int32)(unsafe.Add(mBase, _consts[154]))
									v186 = *(*int32)(unsafe.Add(mBase, _consts[55]))
									if v186 < int32(0) {
										v190 = int32(*(*uint8)(unsafe.Add(mBase, _consts[156])))
										if v190 == int32(1) {
											v194 = int32(*(*uint8)(unsafe.Add(mBase, _consts[111])))
											if v194 == int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v202 = m.ExcPending
												if v202 != 0 {
													return
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return
													} else {
														F_errmsg(m, int32(540340), int32(0))
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return
														} else {
															F_errfinish(m, int32(498059), int32(1748), int32(86913))
															mBase = m.M
															v214 = m.ExcPending
															if v214 != 0 {
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
												v197 = F_GetSerializableTransactionSnapshotInt(m, v184, l1, l2)
												mBase = m.M
												v198 = m.ExcPending
												if v198 != 0 {
													return
												} else {
													v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
													v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
													v221 = v219 + int32(24)
													v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
													v224 = v219 + int32(16)
													v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
													v227 = int32(2)
													v229 = int32(72)
													v234 = v225<<(uint(v227)%32) + v229
													if int32(0) < v222 {
														v237 = (v222+v225)<<(uint(v227)%32) + v229
													} else {
														v237 = v234
													}
													v238 = F_MemoryContextAlloc(m, v217, v237)
													mBase = m.M
													v239 = m.ExcPending
													if v239 != 0 {
														return
													} else {
														v241 = v238 + int32(48)
														v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
														*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
														v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
														v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
														v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
														v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
														v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
														v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
														v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
														*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
														*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
														v260 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
														*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
														v264 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
														v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
														if v266 != 0 {
															v268 = v238 + int32(72)
															*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
															v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
															v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
															v273 = v271 << (uint(int32(2)) % 32)
															if v273 != 0 {
																v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
																mBase = m.M
															} else {
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
														}
														v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
														if v280 <= int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
															v302 = int32(1)
														} else {
															v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
															if v283 == int32(1) {
																v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
																if v286 != int32(1) {
																	*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
																	v302 = int32(1)
																} else {
																	v289 = v234 + v238
																	*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
																	v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
																	v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
																	v294 = v292 << (uint(int32(2)) % 32)
																	if v294 != 0 {
																		v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																		mBase = m.M
																	} else {
																	}
																	v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
																	v302 = v297 + int32(1)
																}
															} else {
																v289 = v234 + v238
																*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
																v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
																v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
																v294 = v292 << (uint(int32(2)) % 32)
																if v294 != 0 {
																	v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																	mBase = m.M
																} else {
																}
																v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
																v302 = v297 + int32(1)
															}
														}
														*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
														*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
														*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
														F_pairingheap_add(m, int32(4156648), v238+int32(52))
														mBase = m.M
														v312 = m.ExcPending
														if v312 != 0 {
															return
														} else {
															v320 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
															m.G0 = v10 + int32(16)
															return
														}
													}
												}
											}
										} else {
											v197 = F_GetSerializableTransactionSnapshotInt(m, v184, l1, l2)
											mBase = m.M
											v198 = m.ExcPending
											if v198 != 0 {
												return
											} else {
												v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
												v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
												v221 = v219 + int32(24)
												v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
												v224 = v219 + int32(16)
												v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
												v227 = int32(2)
												v229 = int32(72)
												v234 = v225<<(uint(v227)%32) + v229
												if int32(0) < v222 {
													v237 = (v222+v225)<<(uint(v227)%32) + v229
												} else {
													v237 = v234
												}
												v238 = F_MemoryContextAlloc(m, v217, v237)
												mBase = m.M
												v239 = m.ExcPending
												if v239 != 0 {
													return
												} else {
													v241 = v238 + int32(48)
													v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
													*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
													v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
													v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
													v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
													v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
													v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
													v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
													v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
													*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
													*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
													v260 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
													*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
													v264 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
													v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
													if v266 != 0 {
														v268 = v238 + int32(72)
														*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
														v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
														v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
														v273 = v271 << (uint(int32(2)) % 32)
														if v273 != 0 {
															v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
															mBase = m.M
														} else {
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
													}
													v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
													if v280 <= int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
														v302 = int32(1)
													} else {
														v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
														if v283 == int32(1) {
															v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
															if v286 != int32(1) {
																*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
																v302 = int32(1)
															} else {
																v289 = v234 + v238
																*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
																v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
																v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
																v294 = v292 << (uint(int32(2)) % 32)
																if v294 != 0 {
																	v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																	mBase = m.M
																} else {
																}
																v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
																v302 = v297 + int32(1)
															}
														} else {
															v289 = v234 + v238
															*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
															v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
															v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
															v294 = v292 << (uint(int32(2)) % 32)
															if v294 != 0 {
																v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																mBase = m.M
															} else {
															}
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
															v302 = v297 + int32(1)
														}
													}
													*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
													*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
													*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
													F_pairingheap_add(m, int32(4156648), v238+int32(52))
													mBase = m.M
													v312 = m.ExcPending
													if v312 != 0 {
														return
													} else {
														v320 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
														m.G0 = v10 + int32(16)
														return
													}
												}
											}
										}
									} else {
										v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
										v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
										v221 = v219 + int32(24)
										v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
										v224 = v219 + int32(16)
										v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
										v227 = int32(2)
										v229 = int32(72)
										v234 = v225<<(uint(v227)%32) + v229
										if int32(0) < v222 {
											v237 = (v222+v225)<<(uint(v227)%32) + v229
										} else {
											v237 = v234
										}
										v238 = F_MemoryContextAlloc(m, v217, v237)
										mBase = m.M
										v239 = m.ExcPending
										if v239 != 0 {
											return
										} else {
											v241 = v238 + int32(48)
											v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
											*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
											v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
											v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
											v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
											v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
											v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
											v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
											v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
											*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
											*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
											v260 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
											*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
											v264 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
											v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
											if v266 != 0 {
												v268 = v238 + int32(72)
												*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
												v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
												v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
												v273 = v271 << (uint(int32(2)) % 32)
												if v273 != 0 {
													v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
													mBase = m.M
												} else {
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
											}
											v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
											if v280 <= int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
												v302 = int32(1)
											} else {
												v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
												if v283 == int32(1) {
													v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
													if v286 != int32(1) {
														*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
														v302 = int32(1)
													} else {
														v289 = v234 + v238
														*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
														v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
														v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
														v294 = v292 << (uint(int32(2)) % 32)
														if v294 != 0 {
															v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
															mBase = m.M
														} else {
														}
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
														v302 = v297 + int32(1)
													}
												} else {
													v289 = v234 + v238
													*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
													v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
													v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
													v294 = v292 << (uint(int32(2)) % 32)
													if v294 != 0 {
														v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
														mBase = m.M
													} else {
													}
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
													v302 = v297 + int32(1)
												}
											}
											*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
											*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
											*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
											F_pairingheap_add(m, int32(4156648), v238+int32(52))
											mBase = m.M
											v312 = m.ExcPending
											if v312 != 0 {
												return
											} else {
												v320 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
												m.G0 = v10 + int32(16)
												return
											}
										}
									}
								} else {
									v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
									v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
									v221 = v219 + int32(24)
									v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
									v224 = v219 + int32(16)
									v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
									v227 = int32(2)
									v229 = int32(72)
									v234 = v225<<(uint(v227)%32) + v229
									if int32(0) < v222 {
										v237 = (v222+v225)<<(uint(v227)%32) + v229
									} else {
										v237 = v234
									}
									v238 = F_MemoryContextAlloc(m, v217, v237)
									mBase = m.M
									v239 = m.ExcPending
									if v239 != 0 {
										return
									} else {
										v241 = v238 + int32(48)
										v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
										v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
										v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
										v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
										v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
										v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
										v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
										v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
										*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
										*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
										v260 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
										*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
										v264 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
										v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
										if v266 != 0 {
											v268 = v238 + int32(72)
											*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
											v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
											v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
											v273 = v271 << (uint(int32(2)) % 32)
											if v273 != 0 {
												v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
												mBase = m.M
											} else {
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
										}
										v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
										if v280 <= int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
											v302 = int32(1)
										} else {
											v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
											if v283 == int32(1) {
												v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
												if v286 != int32(1) {
													*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
													v302 = int32(1)
												} else {
													v289 = v234 + v238
													*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
													v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
													v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
													v294 = v292 << (uint(int32(2)) % 32)
													if v294 != 0 {
														v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
														mBase = m.M
													} else {
													}
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
													v302 = v297 + int32(1)
												}
											} else {
												v289 = v234 + v238
												*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
												v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
												v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
												v294 = v292 << (uint(int32(2)) % 32)
												if v294 != 0 {
													v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
													mBase = m.M
												} else {
												}
												v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
												v302 = v297 + int32(1)
											}
										}
										*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
										*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
										*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
										F_pairingheap_add(m, int32(4156648), v238+int32(52))
										mBase = m.M
										v312 = m.ExcPending
										if v312 != 0 {
											return
										} else {
											v320 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
											m.G0 = v10 + int32(16)
											return
										}
									}
								}
							} else {
								v320 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
								m.G0 = v10 + int32(16)
								return
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v153 = m.ExcPending
							if v153 != 0 {
								return
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v156 = m.ExcPending
								if v156 != 0 {
									return
								} else {
									F_errmsg(m, int32(86693), int32(0))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										F_errdetail(m, int32(617453), int32(0))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return
										} else {
											F_errfinish(m, int32(494533), int32(568), int32(86867))
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
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
				v170 = F_ProcArrayInstallImportedXmin(m, v92, l1)
				mBase = m.M
				v171 = m.ExcPending
				if v171 != 0 {
					return
				} else {
					if v170 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v328 = m.ExcPending
						if v328 != 0 {
							return
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v331 = m.ExcPending
							if v331 != 0 {
								return
							} else {
								F_errmsg(m, int32(86693), int32(0))
								mBase = m.M
								v335 = m.ExcPending
								if v335 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
									F_errdetail(m, int32(617500), v10)
									mBase = m.M
									v339 = m.ExcPending
									if v339 != 0 {
										return
									} else {
										F_errfinish(m, int32(494533), int32(575), int32(86867))
										mBase = m.M
										v344 = m.ExcPending
										if v344 != 0 {
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
						v178 = *(*int32)(unsafe.Add(mBase, _consts[110]))
						if int32(2) <= v178 {
							if v178 == int32(3) {
								v184 = *(*int32)(unsafe.Add(mBase, _consts[154]))
								v186 = *(*int32)(unsafe.Add(mBase, _consts[55]))
								if v186 < int32(0) {
									v190 = int32(*(*uint8)(unsafe.Add(mBase, _consts[156])))
									if v190 == int32(1) {
										v194 = int32(*(*uint8)(unsafe.Add(mBase, _consts[111])))
										if v194 == int32(1) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v202 = m.ExcPending
											if v202 != 0 {
												return
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return
												} else {
													F_errmsg(m, int32(540340), int32(0))
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return
													} else {
														F_errfinish(m, int32(498059), int32(1748), int32(86913))
														mBase = m.M
														v214 = m.ExcPending
														if v214 != 0 {
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
											v197 = F_GetSerializableTransactionSnapshotInt(m, v184, l1, l2)
											mBase = m.M
											v198 = m.ExcPending
											if v198 != 0 {
												return
											} else {
												v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
												v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
												v221 = v219 + int32(24)
												v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
												v224 = v219 + int32(16)
												v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
												v227 = int32(2)
												v229 = int32(72)
												v234 = v225<<(uint(v227)%32) + v229
												if int32(0) < v222 {
													v237 = (v222+v225)<<(uint(v227)%32) + v229
												} else {
													v237 = v234
												}
												v238 = F_MemoryContextAlloc(m, v217, v237)
												mBase = m.M
												v239 = m.ExcPending
												if v239 != 0 {
													return
												} else {
													v241 = v238 + int32(48)
													v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
													*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
													v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
													v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
													v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
													v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
													v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
													v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
													v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
													*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
													*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
													v260 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
													*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
													v264 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
													v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
													if v266 != 0 {
														v268 = v238 + int32(72)
														*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
														v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
														v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
														v273 = v271 << (uint(int32(2)) % 32)
														if v273 != 0 {
															v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
															mBase = m.M
														} else {
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
													}
													v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
													if v280 <= int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
														v302 = int32(1)
													} else {
														v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
														if v283 == int32(1) {
															v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
															if v286 != int32(1) {
																*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
																v302 = int32(1)
															} else {
																v289 = v234 + v238
																*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
																v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
																v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
																v294 = v292 << (uint(int32(2)) % 32)
																if v294 != 0 {
																	v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																	mBase = m.M
																} else {
																}
																v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
																v302 = v297 + int32(1)
															}
														} else {
															v289 = v234 + v238
															*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
															v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
															v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
															v294 = v292 << (uint(int32(2)) % 32)
															if v294 != 0 {
																v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																mBase = m.M
															} else {
															}
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
															v302 = v297 + int32(1)
														}
													}
													*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
													*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
													*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
													F_pairingheap_add(m, int32(4156648), v238+int32(52))
													mBase = m.M
													v312 = m.ExcPending
													if v312 != 0 {
														return
													} else {
														v320 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
														m.G0 = v10 + int32(16)
														return
													}
												}
											}
										}
									} else {
										v197 = F_GetSerializableTransactionSnapshotInt(m, v184, l1, l2)
										mBase = m.M
										v198 = m.ExcPending
										if v198 != 0 {
											return
										} else {
											v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
											v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
											v221 = v219 + int32(24)
											v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
											v224 = v219 + int32(16)
											v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
											v227 = int32(2)
											v229 = int32(72)
											v234 = v225<<(uint(v227)%32) + v229
											if int32(0) < v222 {
												v237 = (v222+v225)<<(uint(v227)%32) + v229
											} else {
												v237 = v234
											}
											v238 = F_MemoryContextAlloc(m, v217, v237)
											mBase = m.M
											v239 = m.ExcPending
											if v239 != 0 {
												return
											} else {
												v241 = v238 + int32(48)
												v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
												*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
												v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
												*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
												v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
												*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
												v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
												*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
												v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
												*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
												v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
												*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
												v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
												*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
												v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
												*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
												*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
												v260 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
												*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
												v264 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
												v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
												if v266 != 0 {
													v268 = v238 + int32(72)
													*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
													v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
													v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
													v273 = v271 << (uint(int32(2)) % 32)
													if v273 != 0 {
														v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
														mBase = m.M
													} else {
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
												}
												v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
												if v280 <= int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
													v302 = int32(1)
												} else {
													v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
													if v283 == int32(1) {
														v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
														if v286 != int32(1) {
															*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
															v302 = int32(1)
														} else {
															v289 = v234 + v238
															*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
															v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
															v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
															v294 = v292 << (uint(int32(2)) % 32)
															if v294 != 0 {
																v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																mBase = m.M
															} else {
															}
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
															v302 = v297 + int32(1)
														}
													} else {
														v289 = v234 + v238
														*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
														v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
														v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
														v294 = v292 << (uint(int32(2)) % 32)
														if v294 != 0 {
															v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
															mBase = m.M
														} else {
														}
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
														v302 = v297 + int32(1)
													}
												}
												*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
												*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
												*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
												F_pairingheap_add(m, int32(4156648), v238+int32(52))
												mBase = m.M
												v312 = m.ExcPending
												if v312 != 0 {
													return
												} else {
													v320 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
													m.G0 = v10 + int32(16)
													return
												}
											}
										}
									}
								} else {
									v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
									v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
									v221 = v219 + int32(24)
									v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
									v224 = v219 + int32(16)
									v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
									v227 = int32(2)
									v229 = int32(72)
									v234 = v225<<(uint(v227)%32) + v229
									if int32(0) < v222 {
										v237 = (v222+v225)<<(uint(v227)%32) + v229
									} else {
										v237 = v234
									}
									v238 = F_MemoryContextAlloc(m, v217, v237)
									mBase = m.M
									v239 = m.ExcPending
									if v239 != 0 {
										return
									} else {
										v241 = v238 + int32(48)
										v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
										v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
										v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
										v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
										v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
										v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
										v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
										v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
										*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
										*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
										v260 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
										*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
										v264 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
										v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
										if v266 != 0 {
											v268 = v238 + int32(72)
											*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
											v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
											v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
											v273 = v271 << (uint(int32(2)) % 32)
											if v273 != 0 {
												v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
												mBase = m.M
											} else {
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
										}
										v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
										if v280 <= int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
											v302 = int32(1)
										} else {
											v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
											if v283 == int32(1) {
												v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
												if v286 != int32(1) {
													*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
													v302 = int32(1)
												} else {
													v289 = v234 + v238
													*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
													v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
													v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
													v294 = v292 << (uint(int32(2)) % 32)
													if v294 != 0 {
														v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
														mBase = m.M
													} else {
													}
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
													v302 = v297 + int32(1)
												}
											} else {
												v289 = v234 + v238
												*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
												v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
												v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
												v294 = v292 << (uint(int32(2)) % 32)
												if v294 != 0 {
													v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
													mBase = m.M
												} else {
												}
												v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
												v302 = v297 + int32(1)
											}
										}
										*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
										*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
										*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
										F_pairingheap_add(m, int32(4156648), v238+int32(52))
										mBase = m.M
										v312 = m.ExcPending
										if v312 != 0 {
											return
										} else {
											v320 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
											m.G0 = v10 + int32(16)
											return
										}
									}
								}
							} else {
								v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
								v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
								v221 = v219 + int32(24)
								v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
								v224 = v219 + int32(16)
								v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
								v227 = int32(2)
								v229 = int32(72)
								v234 = v225<<(uint(v227)%32) + v229
								if int32(0) < v222 {
									v237 = (v222+v225)<<(uint(v227)%32) + v229
								} else {
									v237 = v234
								}
								v238 = F_MemoryContextAlloc(m, v217, v237)
								mBase = m.M
								v239 = m.ExcPending
								if v239 != 0 {
									return
								} else {
									v241 = v238 + int32(48)
									v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
									*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
									v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
									*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
									v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
									*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
									v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
									*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
									v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
									*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
									v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
									*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
									v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
									v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
									*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
									*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
									v260 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
									*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
									v264 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
									v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
									if v266 != 0 {
										v268 = v238 + int32(72)
										*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
										v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
										v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
										v273 = v271 << (uint(int32(2)) % 32)
										if v273 != 0 {
											v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
											mBase = m.M
										} else {
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
									}
									v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
									if v280 <= int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
										v302 = int32(1)
									} else {
										v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
										if v283 == int32(1) {
											v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
											if v286 != int32(1) {
												*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
												v302 = int32(1)
											} else {
												v289 = v234 + v238
												*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
												v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
												v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
												v294 = v292 << (uint(int32(2)) % 32)
												if v294 != 0 {
													v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
													mBase = m.M
												} else {
												}
												v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
												v302 = v297 + int32(1)
											}
										} else {
											v289 = v234 + v238
											*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
											v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
											v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
											v294 = v292 << (uint(int32(2)) % 32)
											if v294 != 0 {
												v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
												mBase = m.M
											} else {
											}
											v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
											v302 = v297 + int32(1)
										}
									}
									*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
									*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
									*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
									F_pairingheap_add(m, int32(4156648), v238+int32(52))
									mBase = m.M
									v312 = m.ExcPending
									if v312 != 0 {
										return
									} else {
										v320 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
										m.G0 = v10 + int32(16)
										return
									}
								}
							}
						} else {
							v320 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
							m.G0 = v10 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		F_pairingheap_remove(m, int32(4156648), v13+int32(52))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[42])) = int32(0)
			v25 = *(*int32)(unsafe.Add(mBase, _consts[40]))
			if v25 != 0 {
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _consts[41]))
				if v28 != 0 {
					v30 = *(*int32)(unsafe.Add(mBase, _consts[91]))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+40))
					v33 = *(*int32)(unsafe.Add(mBase, _consts[41]))
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
						*(*int32)(unsafe.Add(mBase, _consts[146])) = v52
						v56 = *(*int32)(unsafe.Add(mBase, _consts[91]))
						*(*int32)(unsafe.Add(mBase, uint32(v56)+40)) = v52
					}
				} else {
					v52 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[146])) = v52
					v56 = *(*int32)(unsafe.Add(mBase, _consts[91]))
					*(*int32)(unsafe.Add(mBase, uint32(v56)+40)) = v52
				}
			}
			v61 = F_GetSnapshotData(m, int32(4489568))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[154])) = v61
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v64
				v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v66
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v68
				if v68 != 0 {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v73 = v68 << (uint(int32(2)) % 32)
					if v73 != 0 {
						v74 = F__emscripten_memcpy_bulkmem(m, v70, v71, v73)
						mBase = m.M
					} else {
					}
				} else {
				}
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = v76
				if int32(0) < v76 {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v83 = v76 << (uint(int32(2)) % 32)
					if v83 != 0 {
						v84 = F__emscripten_memcpy_bulkmem(m, v80, v81, v83)
						mBase = m.M
					} else {
					}
				} else {
				}
				v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
				*(*uint8)(unsafe.Add(mBase, uint32(v61)+28)) = uint8(v86)
				v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
				*(*int64)(unsafe.Add(mBase, uint32(v61)+64)) = int64(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v61)+29)) = uint8(v88)
				v92 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
				if l3 != 0 {
					v93 = int32(0)
					v95 = *(*int32)(unsafe.Add(mBase, _consts[29]))
					v99 = F_LWLockAcquire(m, v95+int32(512), v93)
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return
					} else {
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
						v102 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
						v104 = *(*int32)(unsafe.Add(mBase, _consts[226]))
						if v102 != v104 {
							v142 = v93
						} else {
							if base.Ui32(v101) < base.Ui32(int32(3)) {
								v142 = v93
							} else {
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v92))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v101)) == int32(0) {
									v119 = base.B2i32(base.Ui32(v101) <= base.Ui32(v92))
								} else {
									v119 = base.B2i32(v101-v92 <= int32(0))
								}
								if v119 == int32(0) {
									v142 = v93
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[146])) = v92
									v125 = *(*int32)(unsafe.Add(mBase, _consts[91]))
									*(*int32)(unsafe.Add(mBase, uint32(v125)+40)) = v92
									v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+124)))
									v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+124)))
									v133 = v127&int32(6) | v130&int32(-7)
									*(*uint8)(unsafe.Add(mBase, uint32(v125)+124)) = uint8(v133)
									v136 = *(*int32)(unsafe.Add(mBase, _consts[94]))
									v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
									v138 = *(*int32)(unsafe.Add(mBase, uint32(v125)+48))
									*(*uint8)(unsafe.Add(mBase, uint32(v137+v138))) = uint8(v133)
									v142 = int32(1)
								}
							}
						}
						v145 = *(*int32)(unsafe.Add(mBase, _consts[29]))
						F_LWLockRelease(m, v145+int32(512))
						mBase = m.M
						v149 = m.ExcPending
						if v149 != 0 {
							return
						} else {
							if v142 != 0 {
								v178 = *(*int32)(unsafe.Add(mBase, _consts[110]))
								if int32(2) <= v178 {
									if v178 == int32(3) {
										v184 = *(*int32)(unsafe.Add(mBase, _consts[154]))
										v186 = *(*int32)(unsafe.Add(mBase, _consts[55]))
										if v186 < int32(0) {
											v190 = int32(*(*uint8)(unsafe.Add(mBase, _consts[156])))
											if v190 == int32(1) {
												v194 = int32(*(*uint8)(unsafe.Add(mBase, _consts[111])))
												if v194 == int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v202 = m.ExcPending
													if v202 != 0 {
														return
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v205 = m.ExcPending
														if v205 != 0 {
															return
														} else {
															F_errmsg(m, int32(540340), int32(0))
															mBase = m.M
															v209 = m.ExcPending
															if v209 != 0 {
																return
															} else {
																F_errfinish(m, int32(498059), int32(1748), int32(86913))
																mBase = m.M
																v214 = m.ExcPending
																if v214 != 0 {
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
													v197 = F_GetSerializableTransactionSnapshotInt(m, v184, l1, l2)
													mBase = m.M
													v198 = m.ExcPending
													if v198 != 0 {
														return
													} else {
														v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
														v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
														v221 = v219 + int32(24)
														v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
														v224 = v219 + int32(16)
														v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
														v227 = int32(2)
														v229 = int32(72)
														v234 = v225<<(uint(v227)%32) + v229
														if int32(0) < v222 {
															v237 = (v222+v225)<<(uint(v227)%32) + v229
														} else {
															v237 = v234
														}
														v238 = F_MemoryContextAlloc(m, v217, v237)
														mBase = m.M
														v239 = m.ExcPending
														if v239 != 0 {
															return
														} else {
															v241 = v238 + int32(48)
															v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
															*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
															v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
															*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
															v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
															*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
															v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
															*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
															v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
															*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
															v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
															*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
															v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
															*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
															v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
															*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
															*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
															v260 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
															*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
															v264 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
															v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
															if v266 != 0 {
																v268 = v238 + int32(72)
																*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
																v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
																v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
																v273 = v271 << (uint(int32(2)) % 32)
																if v273 != 0 {
																	v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
																	mBase = m.M
																} else {
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
															}
															v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
															if v280 <= int32(0) {
																*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
																v302 = int32(1)
															} else {
																v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
																if v283 == int32(1) {
																	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
																	if v286 != int32(1) {
																		*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
																		v302 = int32(1)
																	} else {
																		v289 = v234 + v238
																		*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
																		v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
																		v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
																		v294 = v292 << (uint(int32(2)) % 32)
																		if v294 != 0 {
																			v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																			mBase = m.M
																		} else {
																		}
																		v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
																		v302 = v297 + int32(1)
																	}
																} else {
																	v289 = v234 + v238
																	*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
																	v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
																	v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
																	v294 = v292 << (uint(int32(2)) % 32)
																	if v294 != 0 {
																		v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																		mBase = m.M
																	} else {
																	}
																	v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
																	v302 = v297 + int32(1)
																}
															}
															*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
															*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
															*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
															F_pairingheap_add(m, int32(4156648), v238+int32(52))
															mBase = m.M
															v312 = m.ExcPending
															if v312 != 0 {
																return
															} else {
																v320 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
																m.G0 = v10 + int32(16)
																return
															}
														}
													}
												}
											} else {
												v197 = F_GetSerializableTransactionSnapshotInt(m, v184, l1, l2)
												mBase = m.M
												v198 = m.ExcPending
												if v198 != 0 {
													return
												} else {
													v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
													v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
													v221 = v219 + int32(24)
													v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
													v224 = v219 + int32(16)
													v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
													v227 = int32(2)
													v229 = int32(72)
													v234 = v225<<(uint(v227)%32) + v229
													if int32(0) < v222 {
														v237 = (v222+v225)<<(uint(v227)%32) + v229
													} else {
														v237 = v234
													}
													v238 = F_MemoryContextAlloc(m, v217, v237)
													mBase = m.M
													v239 = m.ExcPending
													if v239 != 0 {
														return
													} else {
														v241 = v238 + int32(48)
														v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
														*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
														v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
														v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
														v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
														v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
														v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
														v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
														v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
														*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
														*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
														v260 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
														*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
														v264 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
														v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
														if v266 != 0 {
															v268 = v238 + int32(72)
															*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
															v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
															v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
															v273 = v271 << (uint(int32(2)) % 32)
															if v273 != 0 {
																v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
																mBase = m.M
															} else {
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
														}
														v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
														if v280 <= int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
															v302 = int32(1)
														} else {
															v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
															if v283 == int32(1) {
																v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
																if v286 != int32(1) {
																	*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
																	v302 = int32(1)
																} else {
																	v289 = v234 + v238
																	*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
																	v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
																	v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
																	v294 = v292 << (uint(int32(2)) % 32)
																	if v294 != 0 {
																		v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																		mBase = m.M
																	} else {
																	}
																	v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
																	v302 = v297 + int32(1)
																}
															} else {
																v289 = v234 + v238
																*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
																v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
																v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
																v294 = v292 << (uint(int32(2)) % 32)
																if v294 != 0 {
																	v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																	mBase = m.M
																} else {
																}
																v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
																v302 = v297 + int32(1)
															}
														}
														*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
														*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
														*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
														F_pairingheap_add(m, int32(4156648), v238+int32(52))
														mBase = m.M
														v312 = m.ExcPending
														if v312 != 0 {
															return
														} else {
															v320 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
															m.G0 = v10 + int32(16)
															return
														}
													}
												}
											}
										} else {
											v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
											v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
											v221 = v219 + int32(24)
											v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
											v224 = v219 + int32(16)
											v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
											v227 = int32(2)
											v229 = int32(72)
											v234 = v225<<(uint(v227)%32) + v229
											if int32(0) < v222 {
												v237 = (v222+v225)<<(uint(v227)%32) + v229
											} else {
												v237 = v234
											}
											v238 = F_MemoryContextAlloc(m, v217, v237)
											mBase = m.M
											v239 = m.ExcPending
											if v239 != 0 {
												return
											} else {
												v241 = v238 + int32(48)
												v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
												*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
												v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
												*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
												v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
												*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
												v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
												*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
												v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
												*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
												v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
												*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
												v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
												*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
												v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
												*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
												*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
												v260 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
												*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
												v264 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
												v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
												if v266 != 0 {
													v268 = v238 + int32(72)
													*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
													v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
													v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
													v273 = v271 << (uint(int32(2)) % 32)
													if v273 != 0 {
														v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
														mBase = m.M
													} else {
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
												}
												v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
												if v280 <= int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
													v302 = int32(1)
												} else {
													v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
													if v283 == int32(1) {
														v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
														if v286 != int32(1) {
															*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
															v302 = int32(1)
														} else {
															v289 = v234 + v238
															*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
															v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
															v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
															v294 = v292 << (uint(int32(2)) % 32)
															if v294 != 0 {
																v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																mBase = m.M
															} else {
															}
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
															v302 = v297 + int32(1)
														}
													} else {
														v289 = v234 + v238
														*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
														v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
														v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
														v294 = v292 << (uint(int32(2)) % 32)
														if v294 != 0 {
															v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
															mBase = m.M
														} else {
														}
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
														v302 = v297 + int32(1)
													}
												}
												*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
												*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
												*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
												F_pairingheap_add(m, int32(4156648), v238+int32(52))
												mBase = m.M
												v312 = m.ExcPending
												if v312 != 0 {
													return
												} else {
													v320 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
													m.G0 = v10 + int32(16)
													return
												}
											}
										}
									} else {
										v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
										v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
										v221 = v219 + int32(24)
										v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
										v224 = v219 + int32(16)
										v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
										v227 = int32(2)
										v229 = int32(72)
										v234 = v225<<(uint(v227)%32) + v229
										if int32(0) < v222 {
											v237 = (v222+v225)<<(uint(v227)%32) + v229
										} else {
											v237 = v234
										}
										v238 = F_MemoryContextAlloc(m, v217, v237)
										mBase = m.M
										v239 = m.ExcPending
										if v239 != 0 {
											return
										} else {
											v241 = v238 + int32(48)
											v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
											*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
											v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
											v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
											v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
											v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
											v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
											v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
											v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
											*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
											*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
											v260 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
											*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
											v264 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
											v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
											if v266 != 0 {
												v268 = v238 + int32(72)
												*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
												v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
												v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
												v273 = v271 << (uint(int32(2)) % 32)
												if v273 != 0 {
													v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
													mBase = m.M
												} else {
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
											}
											v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
											if v280 <= int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
												v302 = int32(1)
											} else {
												v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
												if v283 == int32(1) {
													v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
													if v286 != int32(1) {
														*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
														v302 = int32(1)
													} else {
														v289 = v234 + v238
														*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
														v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
														v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
														v294 = v292 << (uint(int32(2)) % 32)
														if v294 != 0 {
															v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
															mBase = m.M
														} else {
														}
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
														v302 = v297 + int32(1)
													}
												} else {
													v289 = v234 + v238
													*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
													v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
													v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
													v294 = v292 << (uint(int32(2)) % 32)
													if v294 != 0 {
														v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
														mBase = m.M
													} else {
													}
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
													v302 = v297 + int32(1)
												}
											}
											*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
											*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
											*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
											F_pairingheap_add(m, int32(4156648), v238+int32(52))
											mBase = m.M
											v312 = m.ExcPending
											if v312 != 0 {
												return
											} else {
												v320 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
												m.G0 = v10 + int32(16)
												return
											}
										}
									}
								} else {
									v320 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
									m.G0 = v10 + int32(16)
									return
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									F_errcode(m, int32(325))
									mBase = m.M
									v156 = m.ExcPending
									if v156 != 0 {
										return
									} else {
										F_errmsg(m, int32(86693), int32(0))
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
											return
										} else {
											F_errdetail(m, int32(617453), int32(0))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return
											} else {
												F_errfinish(m, int32(494533), int32(568), int32(86867))
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
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
					v170 = F_ProcArrayInstallImportedXmin(m, v92, l1)
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return
					} else {
						if v170 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v328 = m.ExcPending
							if v328 != 0 {
								return
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v331 = m.ExcPending
								if v331 != 0 {
									return
								} else {
									F_errmsg(m, int32(86693), int32(0))
									mBase = m.M
									v335 = m.ExcPending
									if v335 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
										F_errdetail(m, int32(617500), v10)
										mBase = m.M
										v339 = m.ExcPending
										if v339 != 0 {
											return
										} else {
											F_errfinish(m, int32(494533), int32(575), int32(86867))
											mBase = m.M
											v344 = m.ExcPending
											if v344 != 0 {
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
							v178 = *(*int32)(unsafe.Add(mBase, _consts[110]))
							if int32(2) <= v178 {
								if v178 == int32(3) {
									v184 = *(*int32)(unsafe.Add(mBase, _consts[154]))
									v186 = *(*int32)(unsafe.Add(mBase, _consts[55]))
									if v186 < int32(0) {
										v190 = int32(*(*uint8)(unsafe.Add(mBase, _consts[156])))
										if v190 == int32(1) {
											v194 = int32(*(*uint8)(unsafe.Add(mBase, _consts[111])))
											if v194 == int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v202 = m.ExcPending
												if v202 != 0 {
													return
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return
													} else {
														F_errmsg(m, int32(540340), int32(0))
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return
														} else {
															F_errfinish(m, int32(498059), int32(1748), int32(86913))
															mBase = m.M
															v214 = m.ExcPending
															if v214 != 0 {
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
												v197 = F_GetSerializableTransactionSnapshotInt(m, v184, l1, l2)
												mBase = m.M
												v198 = m.ExcPending
												if v198 != 0 {
													return
												} else {
													v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
													v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
													v221 = v219 + int32(24)
													v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
													v224 = v219 + int32(16)
													v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
													v227 = int32(2)
													v229 = int32(72)
													v234 = v225<<(uint(v227)%32) + v229
													if int32(0) < v222 {
														v237 = (v222+v225)<<(uint(v227)%32) + v229
													} else {
														v237 = v234
													}
													v238 = F_MemoryContextAlloc(m, v217, v237)
													mBase = m.M
													v239 = m.ExcPending
													if v239 != 0 {
														return
													} else {
														v241 = v238 + int32(48)
														v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
														*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
														v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
														v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
														v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
														v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
														v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
														v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
														v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
														*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
														*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
														v260 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
														*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
														v264 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
														v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
														if v266 != 0 {
															v268 = v238 + int32(72)
															*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
															v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
															v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
															v273 = v271 << (uint(int32(2)) % 32)
															if v273 != 0 {
																v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
																mBase = m.M
															} else {
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
														}
														v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
														if v280 <= int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
															v302 = int32(1)
														} else {
															v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
															if v283 == int32(1) {
																v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
																if v286 != int32(1) {
																	*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
																	v302 = int32(1)
																} else {
																	v289 = v234 + v238
																	*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
																	v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
																	v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
																	v294 = v292 << (uint(int32(2)) % 32)
																	if v294 != 0 {
																		v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																		mBase = m.M
																	} else {
																	}
																	v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
																	v302 = v297 + int32(1)
																}
															} else {
																v289 = v234 + v238
																*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
																v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
																v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
																v294 = v292 << (uint(int32(2)) % 32)
																if v294 != 0 {
																	v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																	mBase = m.M
																} else {
																}
																v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
																v302 = v297 + int32(1)
															}
														}
														*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
														*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
														*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
														F_pairingheap_add(m, int32(4156648), v238+int32(52))
														mBase = m.M
														v312 = m.ExcPending
														if v312 != 0 {
															return
														} else {
															v320 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
															m.G0 = v10 + int32(16)
															return
														}
													}
												}
											}
										} else {
											v197 = F_GetSerializableTransactionSnapshotInt(m, v184, l1, l2)
											mBase = m.M
											v198 = m.ExcPending
											if v198 != 0 {
												return
											} else {
												v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
												v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
												v221 = v219 + int32(24)
												v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
												v224 = v219 + int32(16)
												v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
												v227 = int32(2)
												v229 = int32(72)
												v234 = v225<<(uint(v227)%32) + v229
												if int32(0) < v222 {
													v237 = (v222+v225)<<(uint(v227)%32) + v229
												} else {
													v237 = v234
												}
												v238 = F_MemoryContextAlloc(m, v217, v237)
												mBase = m.M
												v239 = m.ExcPending
												if v239 != 0 {
													return
												} else {
													v241 = v238 + int32(48)
													v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
													*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
													v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
													v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
													v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
													v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
													v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
													v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
													v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
													*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
													*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
													v260 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
													*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
													v264 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
													v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
													if v266 != 0 {
														v268 = v238 + int32(72)
														*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
														v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
														v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
														v273 = v271 << (uint(int32(2)) % 32)
														if v273 != 0 {
															v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
															mBase = m.M
														} else {
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
													}
													v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
													if v280 <= int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
														v302 = int32(1)
													} else {
														v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
														if v283 == int32(1) {
															v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
															if v286 != int32(1) {
																*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
																v302 = int32(1)
															} else {
																v289 = v234 + v238
																*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
																v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
																v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
																v294 = v292 << (uint(int32(2)) % 32)
																if v294 != 0 {
																	v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																	mBase = m.M
																} else {
																}
																v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
																v302 = v297 + int32(1)
															}
														} else {
															v289 = v234 + v238
															*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
															v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
															v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
															v294 = v292 << (uint(int32(2)) % 32)
															if v294 != 0 {
																v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
																mBase = m.M
															} else {
															}
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
															v302 = v297 + int32(1)
														}
													}
													*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
													*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
													*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
													F_pairingheap_add(m, int32(4156648), v238+int32(52))
													mBase = m.M
													v312 = m.ExcPending
													if v312 != 0 {
														return
													} else {
														v320 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
														m.G0 = v10 + int32(16)
														return
													}
												}
											}
										}
									} else {
										v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
										v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
										v221 = v219 + int32(24)
										v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
										v224 = v219 + int32(16)
										v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
										v227 = int32(2)
										v229 = int32(72)
										v234 = v225<<(uint(v227)%32) + v229
										if int32(0) < v222 {
											v237 = (v222+v225)<<(uint(v227)%32) + v229
										} else {
											v237 = v234
										}
										v238 = F_MemoryContextAlloc(m, v217, v237)
										mBase = m.M
										v239 = m.ExcPending
										if v239 != 0 {
											return
										} else {
											v241 = v238 + int32(48)
											v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
											*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
											v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
											v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
											v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
											v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
											v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
											v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
											v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
											*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
											*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
											v260 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
											*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
											v264 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
											v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
											if v266 != 0 {
												v268 = v238 + int32(72)
												*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
												v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
												v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
												v273 = v271 << (uint(int32(2)) % 32)
												if v273 != 0 {
													v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
													mBase = m.M
												} else {
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
											}
											v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
											if v280 <= int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
												v302 = int32(1)
											} else {
												v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
												if v283 == int32(1) {
													v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
													if v286 != int32(1) {
														*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
														v302 = int32(1)
													} else {
														v289 = v234 + v238
														*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
														v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
														v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
														v294 = v292 << (uint(int32(2)) % 32)
														if v294 != 0 {
															v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
															mBase = m.M
														} else {
														}
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
														v302 = v297 + int32(1)
													}
												} else {
													v289 = v234 + v238
													*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
													v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
													v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
													v294 = v292 << (uint(int32(2)) % 32)
													if v294 != 0 {
														v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
														mBase = m.M
													} else {
													}
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
													v302 = v297 + int32(1)
												}
											}
											*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
											*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
											*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
											F_pairingheap_add(m, int32(4156648), v238+int32(52))
											mBase = m.M
											v312 = m.ExcPending
											if v312 != 0 {
												return
											} else {
												v320 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
												m.G0 = v10 + int32(16)
												return
											}
										}
									}
								} else {
									v217 = *(*int32)(unsafe.Add(mBase, _consts[76]))
									v219 = *(*int32)(unsafe.Add(mBase, _consts[154]))
									v221 = v219 + int32(24)
									v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
									v224 = v219 + int32(16)
									v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
									v227 = int32(2)
									v229 = int32(72)
									v234 = v225<<(uint(v227)%32) + v229
									if int32(0) < v222 {
										v237 = (v222+v225)<<(uint(v227)%32) + v229
									} else {
										v237 = v234
									}
									v238 = F_MemoryContextAlloc(m, v217, v237)
									mBase = m.M
									v239 = m.ExcPending
									if v239 != 0 {
										return
									} else {
										v241 = v238 + int32(48)
										v242 = *(*int64)(unsafe.Add(mBase, uint32(v219)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v241))) = v242
										v244 = *(*int64)(unsafe.Add(mBase, uint32(v219)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+40)) = v244
										v246 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+24)) = v246
										v248 = *(*int64)(unsafe.Add(mBase, uint32(v219)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+56)) = v248
										v250 = *(*int64)(unsafe.Add(mBase, uint32(v219)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = v250
										v252 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v252
										v254 = *(*int64)(unsafe.Add(mBase, uint32(v219)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v254
										v256 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
										*(*int64)(unsafe.Add(mBase, uint32(v238))) = v256
										*(*int64)(unsafe.Add(mBase, uint32(v238)+64)) = int64(0)
										v260 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v241))) = v260
										*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v260
										v264 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v238)+30)) = uint8(v264)
										v266 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
										if v266 != 0 {
											v268 = v238 + int32(72)
											*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = v268
											v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
											v271 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
											v273 = v271 << (uint(int32(2)) % 32)
											if v273 != 0 {
												v274 = F__emscripten_memcpy_bulkmem(m, v268, v270, v273)
												mBase = m.M
											} else {
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v238)+12)) = int32(0)
										}
										v280 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
										if v280 <= int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
											v302 = int32(1)
										} else {
											v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+28)))
											if v283 == int32(1) {
												v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+29)))
												if v286 != int32(1) {
													*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = int32(0)
													v302 = int32(1)
												} else {
													v289 = v234 + v238
													*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
													v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
													v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
													v294 = v292 << (uint(int32(2)) % 32)
													if v294 != 0 {
														v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
														mBase = m.M
													} else {
													}
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
													v302 = v297 + int32(1)
												}
											} else {
												v289 = v234 + v238
												*(*int32)(unsafe.Add(mBase, uint32(v238)+20)) = v289
												v291 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
												v292 = *(*int32)(unsafe.Add(mBase, uint32(v219)+24))
												v294 = v292 << (uint(int32(2)) % 32)
												if v294 != 0 {
													v295 = F__emscripten_memcpy_bulkmem(m, v289, v291, v294)
													mBase = m.M
												} else {
												}
												v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)+48))
												v302 = v297 + int32(1)
											}
										}
										*(*int32)(unsafe.Add(mBase, _consts[988])) = v238
										*(*int32)(unsafe.Add(mBase, _consts[154])) = v238
										*(*int32)(unsafe.Add(mBase, uint32(v238)+48)) = v302
										F_pairingheap_add(m, int32(4156648), v238+int32(52))
										mBase = m.M
										v312 = m.ExcPending
										if v312 != 0 {
											return
										} else {
											v320 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
											m.G0 = v10 + int32(16)
											return
										}
									}
								}
							} else {
								v320 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[153])) = uint8(v320)
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
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[144]))
	if v10 == l0 {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[145]))
		v34 = v13
		v35 = int32(1)
		switch v34 - int32(2) {
		case 0:
			v73 = v35
			m.G0 = v7 + int32(16)
			return v73
		case 1:
			v39 = *(*int32)(unsafe.Add(mBase, _consts[146]))
			if base.Ui32(l0) < base.Ui32(int32(3)) {
				if base.Ui32(l0) < base.Ui32(v39) {
					v73 = v35
					m.G0 = v7 + int32(16)
					return v73
				} else {
					v48 = F_SubTransGetParent(m, l0)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						if v48 == int32(0) {
							v54 = F_errstart(m, int32(19), int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								if v54 == int32(0) {
									v73 = v35
									m.G0 = v7 + int32(16)
									return v73
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
									F_errmsg_internal(m, int32(55706), v7)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(496593), int32(217), int32(81185))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v73 = v35
											m.G0 = v7 + int32(16)
											return v73
										}
									}
								}
							}
						} else {
							v67 = F_TransactionIdDidAbort(m, v48)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v73 = v67
								m.G0 = v7 + int32(16)
								return v73
							}
						}
					}
				}
			} else {
				if base.Ui32(v39) < base.Ui32(int32(3)) {
					if base.Ui32(l0) < base.Ui32(v39) {
						v73 = v35
						m.G0 = v7 + int32(16)
						return v73
					} else {
						v48 = F_SubTransGetParent(m, l0)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							if v48 == int32(0) {
								v54 = F_errstart(m, int32(19), int32(0))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									if v54 == int32(0) {
										v73 = v35
										m.G0 = v7 + int32(16)
										return v73
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
										F_errmsg_internal(m, int32(55706), v7)
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496593), int32(217), int32(81185))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v73 = v35
												m.G0 = v7 + int32(16)
												return v73
											}
										}
									}
								}
							} else {
								v67 = F_TransactionIdDidAbort(m, v48)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v73 = v67
									m.G0 = v7 + int32(16)
									return v73
								}
							}
						}
					}
				} else {
					if int32(0) <= l0-v39 {
						v48 = F_SubTransGetParent(m, l0)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							if v48 == int32(0) {
								v54 = F_errstart(m, int32(19), int32(0))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									if v54 == int32(0) {
										v73 = v35
										m.G0 = v7 + int32(16)
										return v73
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
										F_errmsg_internal(m, int32(55706), v7)
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496593), int32(217), int32(81185))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v73 = v35
												m.G0 = v7 + int32(16)
												return v73
											}
										}
									}
								}
							} else {
								v67 = F_TransactionIdDidAbort(m, v48)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v73 = v67
									m.G0 = v7 + int32(16)
									return v73
								}
							}
						}
					} else {
						v73 = v35
						m.G0 = v7 + int32(16)
						return v73
					}
				}
			}
		default:
			v73 = int32(0)
			m.G0 = v7 + int32(16)
			return v73
		}
	} else {
		if base.Ui32(l0) <= base.Ui32(int32(2)) {
			v16 = int32(1)
			if base.Ui32(int32(2)) <= base.Ui32(l0-v16) {
				v73 = v16
			} else {
				v73 = int32(0)
			}
			m.G0 = v7 + int32(16)
			return v73
		} else {
			v23 = F_TransactionIdGetStatus(m, l0, v7+int32(8))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				switch v23 {
				case 0, 3:
					v34 = v23
				default:
					*(*int32)(unsafe.Add(mBase, _consts[145])) = v23
					*(*int32)(unsafe.Add(mBase, _consts[144])) = l0
					v32 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
					*(*int64)(unsafe.Add(mBase, _consts[147])) = v32
					v34 = v23
				}
				v35 = int32(1)
				switch v34 - int32(2) {
				case 0:
					v73 = v35
					m.G0 = v7 + int32(16)
					return v73
				case 1:
					v39 = *(*int32)(unsafe.Add(mBase, _consts[146]))
					if base.Ui32(l0) < base.Ui32(int32(3)) {
						if base.Ui32(l0) < base.Ui32(v39) {
							v73 = v35
							m.G0 = v7 + int32(16)
							return v73
						} else {
							v48 = F_SubTransGetParent(m, l0)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								if v48 == int32(0) {
									v54 = F_errstart(m, int32(19), int32(0))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										if v54 == int32(0) {
											v73 = v35
											m.G0 = v7 + int32(16)
											return v73
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
											F_errmsg_internal(m, int32(55706), v7)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496593), int32(217), int32(81185))
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													v73 = v35
													m.G0 = v7 + int32(16)
													return v73
												}
											}
										}
									}
								} else {
									v67 = F_TransactionIdDidAbort(m, v48)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v73 = v67
										m.G0 = v7 + int32(16)
										return v73
									}
								}
							}
						}
					} else {
						if base.Ui32(v39) < base.Ui32(int32(3)) {
							if base.Ui32(l0) < base.Ui32(v39) {
								v73 = v35
								m.G0 = v7 + int32(16)
								return v73
							} else {
								v48 = F_SubTransGetParent(m, l0)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									if v48 == int32(0) {
										v54 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											if v54 == int32(0) {
												v73 = v35
												m.G0 = v7 + int32(16)
												return v73
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
												F_errmsg_internal(m, int32(55706), v7)
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(496593), int32(217), int32(81185))
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return int32(0)
													} else {
														v73 = v35
														m.G0 = v7 + int32(16)
														return v73
													}
												}
											}
										}
									} else {
										v67 = F_TransactionIdDidAbort(m, v48)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return int32(0)
										} else {
											v73 = v67
											m.G0 = v7 + int32(16)
											return v73
										}
									}
								}
							}
						} else {
							if int32(0) <= l0-v39 {
								v48 = F_SubTransGetParent(m, l0)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									if v48 == int32(0) {
										v54 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											if v54 == int32(0) {
												v73 = v35
												m.G0 = v7 + int32(16)
												return v73
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
												F_errmsg_internal(m, int32(55706), v7)
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(496593), int32(217), int32(81185))
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return int32(0)
													} else {
														v73 = v35
														m.G0 = v7 + int32(16)
														return v73
													}
												}
											}
										}
									} else {
										v67 = F_TransactionIdDidAbort(m, v48)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return int32(0)
										} else {
											v73 = v67
											m.G0 = v7 + int32(16)
											return v73
										}
									}
								}
							} else {
								v73 = v35
								m.G0 = v7 + int32(16)
								return v73
							}
						}
					}
				default:
					v73 = int32(0)
					m.G0 = v7 + int32(16)
					return v73
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
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	if base.B2i32(v5 == int32(2)) == int32(0) {
		return
	} else {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[667])))
		if int32(0) < l0 {
			if v15 != 0 {
				return
			} else {
				F_enable_timeout_after(m, int32(8), l0)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			if v15 == int32(0) {
				return
			} else {
				F_disable_timeout(m, int32(8))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
