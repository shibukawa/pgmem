package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RI_FKey_setdefault_upd(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_setdefault_upd_0), int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_set(m, v8, int32(0), int32(2))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			return int64(0)
		}
	}
}
func F_RI_FKey_setnull_del(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_setnull_del_0), int32(3))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_set(m, v8, int32(1), int32(3))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			return int64(0)
		}
	}
}
func F_RI_FKey_trigger_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = l0 - int32(1644)
	if base.Ui32(v3) <= base.Ui32(int32(11)) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(int32(2))%32))+uint32(_c_F_RI_FKey_trigger_type[0])))
		v10 = v8
	} else {
		v10 = int32(0)
	}
	return v10
}
func F_ri_LoadConstraintInfo(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v132 int64
	_ = v132
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = l0
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ri_LoadConstraintInfo[0]))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v59 = v14
	goto L3
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = int64(3092376453124)
	v21 = v10 + int32(48)
	v23 = F_hash_create(m, int32(_a_F_ri_LoadConstraintInfo_0), int64(64), v21, int32(40))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v65 = F_hash_search(m, v59, v10+int32(44), int32(1), v10+int32(48))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L10
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ri_LoadConstraintInfo[0])) = v23
	F_CacheRegisterSyscacheCallback(m, int32(19), int32(1698), int64(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_CacheRegisterSyscacheCallback(m, int32(3), int32(1698), int64(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = int64(51539607560)
	v44 = F_hash_create(m, int32(_a_F_ri_LoadConstraintInfo_1), int64(256), v21, int32(40))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ri_LoadConstraintInfo[1])) = v44
	*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = int64(292057776136)
	v53 = F_hash_create(m, int32(_a_F_ri_LoadConstraintInfo_2), int64(256), v21, int32(40))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ri_LoadConstraintInfo[2])) = v53
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_ri_LoadConstraintInfo[0]))
	v59 = v57
	goto L3
L10:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	if v67 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L53
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L50
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L47
	}
L14:
	;
	m.G0 = v10 + int32(96)
	return v65
L15:
	;
	v74 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+44)))
	v75 = F_SearchSysCache1(m, int32(19), v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L20
	}
L16:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)) = uint8(v70)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)))
	if v72 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	if v75 == int32(0) {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+22)))
	v81 = v79 + v80
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+72)))
	if v82 != int32(102) {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+92))
	if v85 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v119 = F_GetSysCacheHashValue(m, int32(19), base.I64_extend_i32_u(v109), int64(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L33
	}
L24:
	;
	v86 = v85
	goto L27
L25:
	;
	goto L26
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v107
	v109 = v107
	goto L23
L27:
	;
	v95 = F_SearchSysCache1(m, int32(19), base.I64_extend_i32_u(v86))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v86
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	v109 = v106
	goto L23
L29:
	;
	if v95 == int32(0) {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+22)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99+v100)+92))
	F_ReleaseCatCache(m, v95)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if v102 != 0 {
		v86 = v102
		goto L27
	} else {
		goto L32
	}
L32:
	;
	goto L28
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v119
	v123 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v65)+8)))
	v125 = F_GetSysCacheHashValue(m, int32(19), v123, int64(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v125
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v81)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+20)) = v128
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v81)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+28)) = v130
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v81)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+36)) = v132
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v81)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+44)) = v134
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v81)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+52)) = v136
	v138 = *(*int64)(unsafe.Add(mBase, uint32(v81)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+60)) = v138
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v81)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+68)) = v140
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v81)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+76)) = v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v81)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+84)) = v144
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v81)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+88)) = v146
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v65)+92)) = uint8(v148)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+101)))
	*(*uint8)(unsafe.Add(mBase, uint32(v65)+93)) = uint8(v150)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+102)))
	*(*uint8)(unsafe.Add(mBase, uint32(v65)+164)) = uint8(v152)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+107)))
	*(*uint8)(unsafe.Add(mBase, uint32(v65)+165)) = uint8(v154)
	F_DeconstructFkConstraintRow(m, v75, v65+int32(168), v65+int32(236), v65+int32(172), v65+int32(300), v65+int32(428), v65+int32(556), v65+int32(96), v65+int32(100))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+165)))
	if v174 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v81)+88))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v65)+168))
	v179 = F_get_index_column_opclass(m, v177, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v81)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+704)) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v65)+84))
	v192 = F_get_rel_relkind(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	F_FindFKPeriodOpers(m, v179, v65+int32(684), v65+int32(688), v65+int32(692))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+712)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v65)+708)) = uint8(base.B2i32(v192 == int32(112)))
	F_ReleaseCatCache(m, v75)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v202 = v65 + int32(696)
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_ri_LoadConstraintInfo[3]))
	if v204 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+696)) = v214
	v216 = int32(_a_F_ri_LoadConstraintInfo_3)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+700)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v214)+4)) = v202
	*(*int32)(unsafe.Add(mBase, _c_F_ri_LoadConstraintInfo[4])) = v202
	v221 = int32(_a_F_ri_LoadConstraintInfo_4)
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_ri_LoadConstraintInfo[5]))
	v224 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ri_LoadConstraintInfo[5])) = v223 + v224
	*(*int32)(unsafe.Add(mBase, uint32(v65)+716)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)) = uint8(v224)
	goto L14
L44:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_ri_LoadConstraintInfo[4]))
	v214 = v206
	goto L43
L45:
	;
	goto L46
L46:
	;
	v208 = int32(_a_F_ri_LoadConstraintInfo_3)
	*(*int32)(unsafe.Add(mBase, _c_F_ri_LoadConstraintInfo[3])) = v208
	*(*int32)(unsafe.Add(mBase, _c_F_ri_LoadConstraintInfo[5])) = int32(0)
	v214 = v208
	goto L43
L47:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v246
	F_errmsg_internal(m, int32(_a_F_ri_LoadConstraintInfo_5), v10)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_ri_LoadConstraintInfo_6), int32(2388), int32(_a_F_ri_LoadConstraintInfo_7))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v260
	F_errmsg_internal(m, int32(_a_F_ri_LoadConstraintInfo_8), v10+int32(32))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_ri_LoadConstraintInfo_6), int32(2393), int32(_a_F_ri_LoadConstraintInfo_7))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v86
	F_errmsg_internal(m, int32(_a_F_ri_LoadConstraintInfo_5), v10+int32(16))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_ri_LoadConstraintInfo_6), int32(2474), int32(_a_F_ri_LoadConstraintInfo_9))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
