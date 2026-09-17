package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ArrayGetNItemsSafe(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int64
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 <= v4 {
		v86 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v86
L2:
	;
	v17 = int64(1)
	v19 = v4
	goto L6
L3:
	;
	v86 = int32(-1)
	goto L1
L4:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L11
	} else {
		goto L23
	}
L5:
	;
	if base.Ui64(v37) < base.Ui64(int64(268435456)) {
		v86 = base.I32_wrap_i64(v37)
		goto L1
	} else {
		goto L20
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1+v19<<(uint(int32(2))%32))))
	if v24 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v46 = F_errsave_start(m, int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L11
	} else {
		goto L18
	}
L8:
	;
	v28 = F_errsave_start(m, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v37 = base.I64_extend_i32_u(v24) * base.I64_extend32_s(v17)
	if base.Ui64(v37+int64(2147483648)) < base.Ui64(int64(4294967296)) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	return int32(0)
L12:
	;
	if v28 == int32(0) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v62 = int32(84)
	goto L4
L14:
	;
	v43 = v19 + int32(1)
	if v43 == l0 {
		goto L5
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	goto L7
L17:
	;
	v17 = v37
	v19 = v43
	goto L6
L18:
	;
	if v46 == int32(0) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v62 = int32(93)
	goto L4
L20:
	;
	v55 = F_errsave_start(m, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	if v55 == int32(0) {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v62 = int32(100)
	goto L4
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(268435455)
	F_errmsg(m, int32(_a_F_ArrayGetNItemsSafe_0), v9)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	F_errsave_finish(m, int32(0), int32(_a_F_ArrayGetNItemsSafe_1), v62, int32(_a_F_ArrayGetNItemsSafe_2))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	goto L3
}
func F_array_agg_array_finalfn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v3 == int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v6 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_finalfn[0]))
			v15 = F_makeArrayResultArr(m, v6, v13, int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v15
			}
		} else {
			v8 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v8)
			return int32(0)
		}
	} else {
		v8 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v8)
		return int32(0)
	}
}
func F_array_agg_finalfn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v11 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v17 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v16
			v26 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_finalfn[0]))
			v28 = F_makeMdArrayResult(m, v11, v17, v6+int32(12), v6+int32(8), v26, int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v34 = v28
				m.G0 = v6 + int32(16)
				return v34
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			v34 = int32(0)
			m.G0 = v6 + int32(16)
			return v34
		}
	} else {
		v13 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
		v34 = int32(0)
		m.G0 = v6 + int32(16)
		return v34
	}
}
func F_array_agg_serialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = v12 + int32(16)
	F_pq_begintypsend(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	F_enlargeStringInfo(m, v16, int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v30 = int32(16711935)
	v34 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v25+v26))) = base.I32_rotr(v21, int32(24))&v30 | base.I32_rotr(v21&v30, v34)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v25 + int32(4)
	v41 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+16)))
	F_enlargeStringInfo(m, v16, v34)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v48 = int64(56)
	v50 = int64(65280)
	v52 = int64(40)
	v55 = int64(16711680)
	v57 = int64(24)
	v59 = int64(4278190080)
	v61 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v45+v46))) = v41<<(uint(v48)%64) | v41&v50<<(uint(v52)%64) | (v41&v55<<(uint(v57)%64) | v41&v59<<(uint(v61)%64)) | (int64(base.Ui64(v41)>>(uint(v61)%64))&v59 | int64(base.Ui64(v41)>>(uint(v57)%64))&v55 | (int64(base.Ui64(v41)>>(uint(v52)%64))&v50 | int64(base.Ui64(v41)>>(uint(v48)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v45 + int32(8)
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+24)))
	F_enlargeStringInfo(m, v16, int32(2))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v94 = int32(8)
	v98 = v87<<(uint(v94)%32) | int32(base.Ui32(v87)>>(uint(v94)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v91+v92))) = uint16(v98)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v91 + int32(2)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+26)))
	F_enlargeStringInfo(m, v16, int32(1))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v107+v108))) = uint8(v103)
	v111 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v107 + v111
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+27)))
	F_enlargeStringInfo(m, v16, v111)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v118+v119))) = uint8(v114)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v118 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	F_appendBinaryStringInfo(m, v16, v125, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+26)))
	if v129 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v239 = v12 + int32(16)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v241))) = v242 << (uint(int32(2)) % 32)
	goto L30
L10:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	F_appendBinaryStringInfo(m, v16, v132, v133<<(uint(int32(2))%32))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	if v139 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L9
L14:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
	v144 = F_MemoryContextAlloc(m, v142, int32(28))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v160 = v139
	goto L16
L16:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v161 <= int32(0) {
		goto L9
	} else {
		goto L20
	}
L17:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	F_getTypeBinaryOutputInfo(m, v146, v12+int32(12), v12+int32(11))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	F_fmgr_info_cxt(m, v153, v144, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v158)+16)) = v144
	v160 = v144
	goto L16
L20:
	;
	v165 = v161
	v169 = int32(0)
	goto L21
L21:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v169))))
	if v176 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L9
L23:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179+v169<<(uint(int32(2))%32))))
	v184 = F_SendFunctionCall(m, v160, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v222 = v165
	goto L25
L25:
	;
	v227 = v169 + int32(1)
	if v227 < v222 {
		v165 = v222
		v169 = v227
		goto L21
	} else {
		goto L29
	}
L26:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v188 = v12 + int32(16)
	F_enlargeStringInfo(m, v188, int32(4))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v195 = int32(2)
	v197 = int32(4)
	v198 = int32(base.Ui32(v186)>>(uint(v195)%32)) - v197
	v199 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v192+v193))) = base.I32_rotr(v198&v199, int32(8)) | base.I32_rotr(v198, int32(24))&v199
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v192 + v197
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	F_appendBinaryStringInfo(m, v188, v184+v197, int32(base.Ui32(v214)>>(uint(v195)%32))-v197)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v222 = v221
	goto L25
L29:
	;
	goto L22
L30:
	;
	m.G0 = v12 + int32(32)
	return v241
}
func F_array_agg_transfn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = F_get_fn_expr_argtype(m, v9, int32(1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != 0 {
			v16 = v7 + int32(12)
			v17 = int32(0)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v18 == v17 {
				v35 = int32(0)
				if v16 == v35 {
					v43 = v35
				} else {
					v38 = v35
					v39 = v17
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = v38
					v43 = v39
				}
				v46 = v43
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				switch v21 - int32(429) {
				case 0:
					if v16 == int32(0) {
						v46 = int32(1)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+168))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
						v38 = v28
						v39 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = v38
						v43 = v39
						v46 = v43
					}
				case 1:
					if v16 == int32(0) {
						v46 = int32(2)
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+368))
						v38 = v33
						v39 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = v38
						v43 = v39
						v46 = v43
					}
				default:
					v35 = int32(0)
					if v16 == v35 {
						v43 = v35
					} else {
						v38 = v35
						v39 = v17
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = v38
						v43 = v39
					}
					v46 = v43
				}
			}
			if v46 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_array_agg_transfn_0), int32(0))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_array_agg_transfn_1), int32(576), int32(_a_F_array_agg_transfn_2))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
				if v49 == int32(1) {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v54 = F_initArrayResult(m, v11, v52, int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v57 = v54
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						if v58 != 0 {
							v61 = int32(0)
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v61 = v60
						}
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						v63 = F_accumArrayResult(m, v57, v61, v58, v11, v62)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(16)
							return v63
						}
					}
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v57 = v56
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v58 != 0 {
						v61 = int32(0)
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v61 = v60
					}
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v63 = F_accumArrayResult(m, v57, v61, v58, v11, v62)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v63
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_array_agg_transfn_3), int32(0))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_array_agg_transfn_1), int32(565), int32(_a_F_array_agg_transfn_2))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
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
func F_array_cat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v304 int32
	_ = v304
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
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v833 int32
	_ = v833
	v25 = m.G0
	v27 = v25 - int32(32)
	m.G0 = v27
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v30 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v27 + int32(32)
	return v833
L2:
	;
	if v29&int32(1) != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	if v29&int32(1) != 0 {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v35 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
	v833 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v39 = F_pg_detoast_datum(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v833 = v39
	goto L1
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v46 = F_pg_detoast_datum(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v49 = F_pg_detoast_datum(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	v833 = v46
	goto L1
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v52 = F_pg_detoast_datum(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	if v54 == v55 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v458 = F_ArrayGetNItemsSafe(m, v457, v436)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L8
	} else {
		goto L115
	}
L17:
	;
	v436 = v412
	v439 = v415
	v457 = v57
	goto L16
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L8
	} else {
		goto L110
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L8
	} else {
		goto L105
	}
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v59 = int32(0)
	if v57|base.B2i32(v58 <= v59) == v59 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L8
	} else {
		goto L98
	}
L23:
	;
	v833 = v52
	goto L1
L24:
	;
	goto L25
L25:
	;
	if v58 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v833 = v49
	goto L1
L27:
	;
	goto L28
L28:
	;
	v67 = int32(1)
	v68 = v58 - v67
	if base.B2i32(base.B2i32(v57 == v58)|base.B2i32(v68 == v57) == int32(0))&base.B2i32(v57 != v58+v67) != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v80 = v49 + int32(16)
	v81 = F_ArrayGetNItemsSafe(m, v57, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	v84 = v52 + int32(16)
	v85 = F_ArrayGetNItemsSafe(m, v58, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	if v87 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v97 = (v90<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L34
L33:
	;
	v97 = v87
	goto L34
L34:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v98 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v108 = (v101<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L37
L36:
	;
	v108 = v98
	goto L37
L37:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v111 = int32(2)
	v112 = v58 << (uint(v111) % 32)
	v113 = v112 + v84
	v115 = v57 << (uint(v111) % 32)
	v116 = v115 + v80
	if v57 == v58 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v118 = F_palloc(m, v115)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v68 == v57 {
		goto L49
	} else {
		goto L50
	}
L41:
	;
	v120 = F_palloc(m, v115)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v122 + v123
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v126
	if v57 < int32(2) {
		v412 = v118
		v415 = v120
		goto L17
	} else {
		goto L43
	}
L43:
	;
	v132 = int32(1)
	goto L44
L44:
	;
	v156 = v132 << (uint(int32(2)) % 32)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v80+v156)))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156+v84)))
	if v158 != v160 {
		goto L18
	} else {
		goto L46
	}
L45:
	;
	v412 = v118
	v415 = v120
	goto L17
L46:
	;
	v162 = v156 + v116
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v156+v113)))
	if v163 != v165 {
		goto L18
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156+v118))) = v158
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	*(*int32)(unsafe.Add(mBase, uint32(v156+v120))) = v170
	v173 = v132 + int32(1)
	if v173 != v57 {
		v132 = v173
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v176 = F_palloc(m, v112)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L8
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v258 = F_palloc(m, v115)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L75
	}
L52:
	;
	v178 = F_palloc(m, v112)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v180 = int32(0)
	v181 = base.B2i32(v112 == v180)
	if v181 == v180 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	base.MemoryCopy(m, v176, v84, v112)
	goto L56
L55:
	;
	goto L56
L56:
	;
	if v181 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	base.MemoryCopy(m, v178, v113, v112)
	goto L59
L58:
	;
	goto L59
L59:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = v188 + int32(1)
	v192 = int32(0)
	if v192 < v57 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v196 = v57
	goto L62
L61:
	;
	v196 = v192
	goto L62
L62:
	;
	v197 = v192
	goto L63
L63:
	;
	if v197 == v196 {
		v436 = v176
		v439 = v178
		v457 = v58
		goto L16
	} else {
		goto L65
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L8
	} else {
		goto L70
	}
L65:
	;
	v222 = int32(2)
	v223 = v197 << (uint(v222) % 32)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v80+v223)))
	v227 = v197 + int32(1)
	v229 = v227 << (uint(v222) % 32)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v176+v229)))
	if v225 == v231 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v116+v223)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v178+v229)))
	if v234 == v236 {
		v197 = v227
		goto L63
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	goto L64
L69:
	;
	goto L68
L70:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(_a_F_array_cat_0), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	F_errdetail(m, int32(_a_F_array_cat_1), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_array_cat_2), int32(478), int32(_a_F_array_cat_3))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	v260 = F_palloc(m, v115)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	v262 = int32(0)
	v263 = base.B2i32(v115 == v262)
	if v263 == v262 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	base.MemoryCopy(m, v258, v80, v115)
	goto L79
L78:
	;
	goto L79
L79:
	;
	if v263 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	base.MemoryCopy(m, v260, v116, v115)
	goto L82
L81:
	;
	goto L82
L82:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v270 + int32(1)
	v274 = int32(0)
	if v274 < v58 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v278 = v58
	goto L85
L84:
	;
	v278 = v274
	goto L85
L85:
	;
	v279 = v274
	goto L86
L86:
	;
	if v279 == v278 {
		v412 = v258
		v415 = v260
		goto L17
	} else {
		goto L88
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L8
	} else {
		goto L93
	}
L88:
	;
	v304 = int32(2)
	v305 = v279 << (uint(v304) % 32)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v84+v305)))
	v309 = v279 + int32(1)
	v311 = v309 << (uint(v304) % 32)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v258+v311)))
	if v307 == v313 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v305+v113)))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v260+v311)))
	if v316 == v318 {
		v279 = v309
		goto L86
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	goto L87
L92:
	;
	goto L91
L93:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(_a_F_array_cat_0), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L8
	} else {
		goto L95
	}
L95:
	;
	F_errdetail(m, int32(_a_F_array_cat_1), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_array_cat_2), int32(506), int32(_a_F_array_cat_3))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L8
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	F_errmsg(m, int32(_a_F_array_cat_0), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L8
	} else {
		goto L100
	}
L100:
	;
	v351 = F_format_type_be(m, v54)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	v353 = F_format_type_be(m, v55)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v351
	F_errdetail(m, int32(_a_F_array_cat_4), v27+int32(16))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_array_cat_2), int32(375), int32(_a_F_array_cat_3))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L8
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(_a_F_array_cat_0), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v57
	F_errdetail(m, int32(_a_F_array_cat_5), v27)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_array_cat_2), int32(413), int32(_a_F_array_cat_3))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L8
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_array_cat_0), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L8
	} else {
		goto L112
	}
L112:
	;
	F_errdetail(m, int32(_a_F_array_cat_6), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L8
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_array_cat_2), int32(449), int32(_a_F_array_cat_3))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L8
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	F_ArrayCheckBounds(m, v457, v436, v439)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L8
	} else {
		goto L116
	}
L116:
	;
	v462 = int32(2)
	v464 = int32(base.Ui32(v110)>>(uint(v462)%32)) - v108
	v467 = int32(base.Ui32(v109)>>(uint(v462)%32)) - v97
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	if v469 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	v495 = v493 + (v464 + v467)
	v496 = F_palloc0(m, v495)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L8
	} else {
		goto L123
	}
L118:
	;
	v493 = (v457<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v494 = int32(0)
	goto L117
L119:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v472 == int32(0) {
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v478 = base.I32_div_s(v458+int32(7), int32(8))
	v485 = (v478 + v457<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v493 = v485
	v494 = v485
	goto L117
L122:
	;
	goto L121
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v496)+12)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v496)+8)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v496)+4)) = v457
	v501 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v496))) = v495 << (uint(v501) % 32)
	v505 = v496 + int32(16)
	v507 = v457 << (uint(v501) % 32)
	v508 = int32(0)
	v509 = base.B2i32(v507 == v508)
	if v509 == v508 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	base.MemoryCopy(m, v505, v436, v507)
	goto L126
L125:
	;
	goto L126
L126:
	;
	if v509 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	base.MemoryCopy(m, v507+v505, v439, v507)
	goto L129
L128:
	;
	goto L129
L129:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v496)+8))
	if v517 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v496)+4))
	v527 = (v520<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L132
L131:
	;
	v527 = v517
	goto L132
L132:
	;
	v529 = v57 << (uint(int32(3)) % 32)
	if v467 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	if v78 != 0 {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	goto L135
L135:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v496)+8))
	if v538 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	v535 = v78
	goto L138
L137:
	;
	v535 = (v529 + int32(23)) & int32(-8)
	goto L138
L138:
	;
	base.MemoryCopy(m, v496+v527, v49+v535, v467)
	goto L135
L139:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v496)+4))
	v548 = (v541<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L141
L140:
	;
	v548 = v538
	goto L141
L141:
	;
	v550 = v58 << (uint(int32(3)) % 32)
	if v464 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	if v77 != 0 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	goto L144
L144:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v496)+8))
	if v560 == int32(0) {
		v833 = v496
		goto L1
	} else {
		goto L148
	}
L145:
	;
	v557 = v77
	goto L147
L146:
	;
	v557 = (v550 + int32(23)) & int32(-8)
	goto L147
L147:
	;
	base.MemoryCopy(m, v496+v548+v467, v52+v557, v464)
	goto L144
L148:
	;
	if v77 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v565 = v550 + v84
	goto L151
L150:
	;
	v565 = int32(0)
	goto L151
L151:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v496)+4))
	v570 = int32(0)
	if v78 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v573 = v529 + v80
	goto L154
L153:
	;
	v573 = v570
	goto L154
L154:
	;
	v574 = int32(0)
	if v81 <= v574 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v496)+8))
	if v700 != 0 {
		goto L186
	} else {
		goto L187
	}
L156:
	;
	goto L155
L157:
	;
	v581 = int32(1)
	v586 = base.I32_div_s(v570, int32(8))
	v587 = v505 + v566<<(uint(int32(3))%32) + v586
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587))))
	if v573 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v681))) = uint8(v682)
	goto L156
L159:
	;
	v591 = v587
	v592 = v588
	v595 = v81
	v596 = v581
	goto L162
L160:
	;
	goto L161
L161:
	;
	v626 = base.I32_div_s(v574, int32(8))
	v627 = v573 + v626
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627))))
	v629 = v587
	v630 = v588
	v632 = v627
	v633 = v81
	v634 = v581
	v635 = int32(1)
	v636 = v628
	goto L170
L162:
	;
	v600 = v592 | v596
	v601 = int32(1)
	v602 = v595 - v601
	v604 = v596 << (uint(v601) % 32)
	if v604 == int32(256) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if v616 != int32(1) {
		v681 = v614
		v682 = v615
		goto L158
	} else {
		goto L169
	}
L164:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v591))) = uint8(v600)
	if v602 == int32(0) {
		goto L156
	} else {
		goto L167
	}
L165:
	;
	v614 = v591
	v615 = v600
	v616 = v604
	goto L166
L166:
	;
	if base.Ui32(int32(1)) < base.Ui32(v595) {
		v591 = v614
		v592 = v615
		v595 = v602
		v596 = v616
		goto L162
	} else {
		goto L168
	}
L167:
	;
	v610 = int32(1)
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591)+1)))
	v614 = v591 + v610
	v615 = v611
	v616 = v610
	goto L166
L168:
	;
	goto L163
L169:
	;
	goto L156
L170:
	;
	if v635&v636 != 0 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if v659 == int32(1) {
		goto L156
	} else {
		goto L185
	}
L172:
	;
	v643 = v630 | v634
	goto L174
L173:
	;
	v643 = v630 & (v634 ^ int32(-1))
	goto L174
L174:
	;
	v644 = int32(1)
	v645 = v633 - v644
	v647 = v634 << (uint(v644) % 32)
	if v647 == int32(256) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v629))) = uint8(v643)
	if v645 == int32(0) {
		goto L156
	} else {
		goto L178
	}
L176:
	;
	v657 = v629
	v658 = v643
	v659 = v647
	goto L177
L177:
	;
	v661 = v635 << (uint(int32(1)) % 32)
	if v661 == int32(256) {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	v653 = int32(1)
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629)+1)))
	v657 = v629 + v653
	v658 = v654
	v659 = v653
	goto L177
L179:
	;
	goto L171
L180:
	;
	if v645 == int32(0) {
		goto L179
	} else {
		goto L183
	}
L181:
	;
	v670 = v632
	v671 = v661
	v672 = v636
	goto L182
L182:
	;
	if base.Ui32(int32(1)) < base.Ui32(v633) {
		v629 = v657
		v630 = v658
		v632 = v670
		v633 = v645
		v634 = v659
		v635 = v671
		v636 = v672
		goto L170
	} else {
		goto L184
	}
L183:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632)+1)))
	v667 = int32(1)
	v670 = v632 + v667
	v671 = v667
	v672 = v666
	goto L182
L184:
	;
	goto L179
L185:
	;
	v681 = v657
	v682 = v658
	goto L158
L186:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v496)+4))
	v706 = v505 + v701<<(uint(int32(3))%32)
	goto L188
L187:
	;
	v706 = int32(0)
	goto L188
L188:
	;
	v707 = int32(0)
	if v85 <= v707 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v833 = v496
	goto L1
L190:
	;
	goto L189
L191:
	;
	v717 = int32(1) << (uint(v81&int32(7)) % 32)
	v719 = base.I32_div_s(v81, int32(8))
	v720 = v706 + v719
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v720))))
	if v565 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v814))) = uint8(v815)
	goto L190
L193:
	;
	v724 = v720
	v725 = v721
	v728 = v85
	v729 = v717
	goto L196
L194:
	;
	goto L195
L195:
	;
	v759 = base.I32_div_s(v707, int32(8))
	v760 = v565 + v759
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v760))))
	v762 = v720
	v763 = v721
	v765 = v760
	v766 = v85
	v767 = v717
	v768 = int32(1)
	v769 = v761
	goto L204
L196:
	;
	v733 = v725 | v729
	v734 = int32(1)
	v735 = v728 - v734
	v737 = v729 << (uint(v734) % 32)
	if v737 == int32(256) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	if v749 != int32(1) {
		v814 = v747
		v815 = v748
		goto L192
	} else {
		goto L203
	}
L198:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v724))) = uint8(v733)
	if v735 == int32(0) {
		goto L190
	} else {
		goto L201
	}
L199:
	;
	v747 = v724
	v748 = v733
	v749 = v737
	goto L200
L200:
	;
	if base.Ui32(int32(1)) < base.Ui32(v728) {
		v724 = v747
		v725 = v748
		v728 = v735
		v729 = v749
		goto L196
	} else {
		goto L202
	}
L201:
	;
	v743 = int32(1)
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724)+1)))
	v747 = v724 + v743
	v748 = v744
	v749 = v743
	goto L200
L202:
	;
	goto L197
L203:
	;
	goto L190
L204:
	;
	if v768&v769 != 0 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	if v792 == int32(1) {
		goto L190
	} else {
		goto L219
	}
L206:
	;
	v776 = v763 | v767
	goto L208
L207:
	;
	v776 = v763 & (v767 ^ int32(-1))
	goto L208
L208:
	;
	v777 = int32(1)
	v778 = v766 - v777
	v780 = v767 << (uint(v777) % 32)
	if v780 == int32(256) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v762))) = uint8(v776)
	if v778 == int32(0) {
		goto L190
	} else {
		goto L212
	}
L210:
	;
	v790 = v762
	v791 = v776
	v792 = v780
	goto L211
L211:
	;
	v794 = v768 << (uint(int32(1)) % 32)
	if v794 == int32(256) {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	v786 = int32(1)
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762)+1)))
	v790 = v762 + v786
	v791 = v787
	v792 = v786
	goto L211
L213:
	;
	goto L205
L214:
	;
	if v778 == int32(0) {
		goto L213
	} else {
		goto L217
	}
L215:
	;
	v803 = v765
	v804 = v794
	v805 = v769
	goto L216
L216:
	;
	if base.Ui32(int32(1)) < base.Ui32(v766) {
		v762 = v790
		v763 = v791
		v765 = v803
		v766 = v778
		v767 = v792
		v768 = v804
		v769 = v805
		goto L204
	} else {
		goto L218
	}
L217:
	;
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765)+1)))
	v800 = int32(1)
	v803 = v765 + v800
	v804 = v800
	v805 = v799
	goto L216
L218:
	;
	goto L213
L219:
	;
	v814 = v790
	v815 = v791
	goto L192
}
func F_array_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(96)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_DatumGetAnyArrayP(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v24 = F_DatumGetAnyArrayP(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v27 == int32(-1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v26 == int32(-1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v33 = v30
	goto L4
L6:
	;
	goto L7
L7:
	;
	v33 = v19 + int32(16)
	goto L4
L8:
	;
	if v27 == int32(-1) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v39 = v36
	goto L8
L10:
	;
	goto L11
L11:
	;
	v39 = v24 + int32(16)
	goto L8
L12:
	;
	if v27 == int32(-1) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v49 = v42
	goto L12
L14:
	;
	goto L15
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v49 = v19 + v43<<(uint(int32(2))%32) + int32(16)
	goto L12
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L132
	}
L17:
	;
	v54 = int32(40)
	goto L19
L18:
	;
	v54 = int32(12)
	goto L19
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v19+v54)))
	if v26 == int32(-1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69+v24)))
	if v56 == v71 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v68 = v59
	v69 = int32(40)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v68 = v24 + v61<<(uint(int32(2))%32) + int32(16)
	v69 = int32(12)
	goto L20
L24:
	;
	if v27 == int32(-1) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L128
	}
L27:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v443 == int32(-1) {
		goto L120
	} else {
		goto L121
	}
L28:
	;
	v77 = int32(28)
	goto L30
L29:
	;
	v77 = int32(4)
	goto L30
L30:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v19+v77)))
	if v26 == int32(-1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v84 = int32(28)
	goto L33
L32:
	;
	v84 = int32(4)
	goto L33
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v24+v84)))
	if v79 != v86 {
		v435 = v2
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v90 = v79 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v90) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	if v152 != 0 {
		v435 = v2
		goto L27
	} else {
		goto L53
	}
L36:
	;
	v152 = int32(0)
	goto L35
L37:
	;
	v126 = v121
	v127 = v122
	v128 = v123
	goto L47
L38:
	;
	if (v33|v39)&int32(3) != 0 {
		v121 = v33
		v122 = v39
		v123 = v90
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v114 = v33
	v115 = v39
	v116 = v90
	goto L40
L40:
	;
	if v116 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v98 = v33
	v99 = v39
	v100 = v90
	goto L42
L42:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v103 != v104 {
		v121 = v98
		v122 = v99
		v123 = v100
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v114 = v109
	v115 = v107
	v116 = v111
	goto L40
L44:
	;
	v106 = int32(4)
	v107 = v99 + v106
	v109 = v98 + v106
	v111 = v100 - v106
	if base.Ui32(int32(3)) < base.Ui32(v111) {
		v98 = v109
		v99 = v107
		v100 = v111
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v121 = v114
	v122 = v115
	v123 = v116
	goto L37
L47:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v131 == v132 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v152 = v131 - v132
	goto L35
L49:
	;
	v134 = int32(1)
	v139 = v128 - v134
	if v139 != 0 {
		v126 = v126 + v134
		v127 = v127 + v134
		v128 = v139
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v90) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	if v214 != 0 {
		v435 = v2
		goto L27
	} else {
		goto L72
	}
L55:
	;
	v214 = int32(0)
	goto L54
L56:
	;
	v188 = v183
	v189 = v184
	v190 = v185
	goto L66
L57:
	;
	if (v49|v68)&int32(3) != 0 {
		v183 = v49
		v184 = v68
		v185 = v90
		goto L56
	} else {
		goto L60
	}
L58:
	;
	v176 = v49
	v177 = v68
	v178 = v90
	goto L59
L59:
	;
	if v178 == int32(0) {
		goto L55
	} else {
		goto L65
	}
L60:
	;
	v160 = v49
	v161 = v68
	v162 = v90
	goto L61
L61:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v165 != v166 {
		v183 = v160
		v184 = v161
		v185 = v162
		goto L56
	} else {
		goto L63
	}
L62:
	;
	v176 = v171
	v177 = v169
	v178 = v173
	goto L59
L63:
	;
	v168 = int32(4)
	v169 = v161 + v168
	v171 = v160 + v168
	v173 = v162 - v168
	if base.Ui32(int32(3)) < base.Ui32(v173) {
		v160 = v171
		v161 = v169
		v162 = v173
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v183 = v176
	v184 = v177
	v185 = v178
	goto L56
L66:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v193 == v194 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v214 = v193 - v194
	goto L54
L68:
	;
	v196 = int32(1)
	v201 = v190 - v196
	if v201 != 0 {
		v188 = v188 + v196
		v189 = v189 + v196
		v190 = v201
		goto L66
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	goto L55
L72:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
	if v216 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v228 = int32(*(*int8)(unsafe.Add(mBase, uint32(v227)+11)))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+10)))
	v230 = int32(*(*int16)(unsafe.Add(mBase, uint32(v227)+8)))
	v231 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+78)) = uint16(v231)
	v233 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+76)) = uint8(v233)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v227 + int32(76)
	v241 = F_ArrayGetNItemsSafe(m, v79, v33)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L80
	}
L74:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	if v217 == v56 {
		v227 = v216
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v220 = F_lookup_type_cache(m, v56, int32(32))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v220)+80))
	if v222 == int32(0) {
		goto L16
	} else {
		goto L79
	}
L79:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v225)+16)) = v220
	v227 = v220
	goto L73
L80:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v243 == int32(-1) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v302
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v306 == int32(-1) {
		goto L95
	} else {
		goto L96
	}
L82:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	if v246 != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = int64(0)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v279 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v249 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v248
	v302 = v249
	goto L81
L86:
	;
	goto L87
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = int64(0)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
	if v256 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v255 + (v259<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v302 = int32(0)
	goto L81
L89:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v255 + v256
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v302 = v255 + v271<<(uint(int32(3))%32) + int32(16)
	goto L81
L91:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v19 + (v282<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v302 = int32(0)
	goto L81
L92:
	;
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v279 + v19
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v302 = v19 + v294<<(uint(int32(3))%32) + int32(16)
	goto L81
L94:
	;
	v366 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v365
	if v241 <= int32(0) {
		v435 = v366
		goto L27
	} else {
		goto L107
	}
L95:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	if v309 != 0 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+20)) = int64(0)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v342 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v309
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	v312 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v311
	v365 = v312
	goto L94
L99:
	;
	goto L100
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+20)) = int64(0)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	if v319 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v318 + (v322<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v365 = int32(0)
	goto L94
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v318 + v319
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	v365 = v318 + v334<<(uint(int32(3))%32) + int32(16)
	goto L94
L104:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v24 + (v345<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v365 = int32(0)
	goto L94
L105:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v342 + v24
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v365 = v24 + v357<<(uint(int32(3))%32) + int32(16)
	goto L94
L107:
	;
	v373 = v229 & int32(1)
	v376 = int32(0)
	goto L108
L108:
	;
	v392 = F_array_iter_next(m, v16+int32(40), v16+int32(19), v376, v230, v373, v228)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L110
	}
L109:
	;
	v435 = v426
	goto L27
L110:
	;
	v398 = F_array_iter_next(m, v16+int32(20), v16+int32(18), v376, v230, v373, v228)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+18)))
	v401 = int32(1)
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+19)))
	if v400&v401&base.B2i32(v403 == v401) != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v426 = int32(1)
	v428 = v376 + v426
	if v428 != v241 {
		v376 = v428
		goto L108
	} else {
		goto L119
	}
L113:
	;
	if v403|v400&int32(1) != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v435 = int32(0)
	goto L27
L115:
	;
	v410 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+92)) = uint8(v410)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v398
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+84)) = uint8(v410)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v392
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+76)) = uint8(v410)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v16)+60))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)))
	v422 = m.T0[v421].(func(*base.Module, int32) int32)(m, v16+int32(60))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+76)))
	if v424 != 0 {
		goto L114
	} else {
		goto L117
	}
L117:
	;
	if v422 != 0 {
		goto L112
	} else {
		goto L118
	}
L118:
	;
	goto L114
L119:
	;
	goto L109
L120:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v450 == int32(-1) {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v19 == v446 {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	F_pfree(m, v19)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	goto L120
L124:
	;
	m.G0 = v16 + int32(96)
	return v435
L125:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v24 == v453 {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	F_pfree(m, v24)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	goto L124
L128:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_array_eq_0), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_array_eq_1), int32(3846), int32(_a_F_array_eq_2))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v484 = F_format_type_be(m, v56)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v484
	F_errmsg(m, int32(_a_F_array_eq_3), v16)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_array_eq_1), int32(3871), int32(_a_F_array_eq_2))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_ge(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2^int32(-1)) >> (uint(int32(31)) % 32))
	}
}
func F_array_lt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2) >> (uint(int32(31)) % 32))
	}
}
func F_array_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v359 int32
	_ = v359
	var v371 int32
	_ = v371
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v576 int32
	_ = v576
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v768 int32
	_ = v768
	var v778 int32
	_ = v778
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v972 int32
	_ = v972
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v997 int32
	_ = v997
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1072 int32
	_ = v1072
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1206 int32
	_ = v1206
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(272)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = F_DatumGetAnyArrayP(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v32 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v35 = int32(40)
	goto L5
L4:
	;
	v35 = int32(12)
	goto L5
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26+v35)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v39 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v82 == int32(-1) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	F_get_type_io_data(m, v37, int32(1), v55+int32(4), v55+int32(6), v55+int32(7), v55+int32(8), v55+int32(12), v55+int32(16))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v44 = F_MemoryContextAlloc(m, v42, int32(48))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v53 == v37 {
		v79 = v39
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v44
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v37 ^ int32(-1)
	v55 = v49
	goto L7
L12:
	;
	v55 = v39
	goto L7
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	F_fmgr_info_cxt(m, v71, v55+int32(20), v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v37
	v79 = v55
	goto L6
L15:
	;
	v85 = int32(28)
	goto L17
L16:
	;
	v85 = int32(4)
	goto L17
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v26+v85)))
	if v82 == int32(-1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+8)))
	v101 = int32(*(*int8)(unsafe.Add(mBase, uint32(v79)+7)))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+6)))
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v79)+4)))
	v104 = F_ArrayGetNItemsSafe(m, v87, v98)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L25
	}
L19:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v98 = v90
	v99 = v91
	goto L18
L20:
	;
	goto L21
L21:
	;
	v93 = v26 + int32(16)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v98 = v93
	v99 = v93 + v94<<(uint(int32(2))%32)
	goto L18
L22:
	;
	m.G0 = v23 + int32(272)
	return v1206
L23:
	;
	v163 = F_palloc(m, v104<<(uint(int32(2))%32))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L35
	}
L24:
	;
	v112 = v106
	goto L31
L25:
	;
	if v104 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v106 = int32(0)
	if v106 < v87 {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v110 = F_pstrdup(m, int32(_a_F_array_out_0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	v151 = v2
	goto L23
L30:
	;
	v1206 = v110
	goto L22
L31:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v99+v112<<(uint(int32(2))%32))))
	v137 = base.B2i32(v135 != int32(1))
	if v135 != int32(1) {
		v151 = v137
		goto L23
	} else {
		goto L33
	}
L32:
	;
	v151 = v137
	goto L23
L33:
	;
	v139 = v112 + int32(1)
	if v139 != v87 {
		v112 = v139
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v165 = F_palloc(m, v104)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v167 == int32(-1) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = int32(1)
	if v104 <= int32(0) {
		goto L51
	} else {
		goto L52
	}
L38:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	if v170 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+12)) = int64(0)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v203 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v170
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	v173 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v172
	v226 = v173
	goto L37
L42:
	;
	goto L43
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+12)) = int64(0)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	if v180 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v179 + (v183<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v226 = int32(0)
	goto L37
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v179 + v180
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v226 = v179 + v195<<(uint(int32(3))%32) + int32(16)
	goto L37
L47:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v26 + (v206<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v226 = int32(0)
	goto L37
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v203 + v26
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v226 = v26 + v218<<(uint(int32(3))%32) + int32(16)
	goto L37
L50:
	;
	if int32(0) < v87 {
		goto L101
	} else {
		goto L102
	}
L51:
	;
	v445 = int32(0)
	goto L50
L52:
	;
	goto L53
L53:
	;
	v240 = int32(0)
	v243 = v2
	goto L54
L54:
	;
	v260 = v163 + v243<<(uint(int32(2))%32)
	v265 = F_array_iter_next(m, v23+int32(12), v23-int32(-64), v243, v103, v102&int32(1), v101)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	v445 = v439
	goto L50
L56:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+64)))
	if v267 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v438 = int32(1)
	v439 = v437 + v438
	v441 = v243 + v438
	if v441 != v104 {
		v240 = v439
		v243 = v441
		goto L54
	} else {
		goto L97
	}
L58:
	;
	v271 = F_pstrdup(m, int32(_a_F_array_out_1))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v279 = F_OutputFunctionCall(m, v79+int32(20), v265)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v271
	v275 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v243+v165))) = uint8(v275)
	v437 = v240 + int32(4)
	goto L57
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v279
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	if v282 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v286 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v243+v165))) = uint8(v286)
	v437 = v240 + int32(2)
	goto L57
L64:
	;
	goto L65
L65:
	;
	v293 = v279
	v294 = int32(_a_F_array_out_1)
	goto L67
L66:
	;
	v333 = base.B2i32(v331 == int32(0))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	if v335 != 0 {
		goto L79
	} else {
		goto L80
	}
L67:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v297 == v298 {
		v320 = v297
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v331 = int32(0)
	goto L66
L69:
	;
	v322 = int32(1)
	if v320 != 0 {
		v293 = v293 + v322
		v294 = v294 + v322
		goto L67
	} else {
		goto L78
	}
L70:
	;
	if base.Ui32((v297-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v308 = v297 | int32(32)
	goto L73
L72:
	;
	v308 = v297
	goto L73
L73:
	;
	if base.Ui32((v298-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v317 = v298 | int32(32)
	goto L76
L75:
	;
	v317 = v298
	goto L76
L76:
	;
	if v308 == v317 {
		v320 = v308
		goto L69
	} else {
		goto L77
	}
L77:
	;
	v331 = v308 - v317
	goto L66
L78:
	;
	goto L68
L79:
	;
	v336 = v335
	v337 = v334
	v338 = v240
	v340 = v333
	goto L82
L80:
	;
	v392 = v240
	v394 = v333
	goto L81
L81:
	;
	v412 = v394 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v243+v165))) = uint8(v412)
	if v412 != 0 {
		goto L94
	} else {
		goto L95
	}
L82:
	;
	v359 = v336 & int32(255)
	switch v359 - int32(123) {
	case 0, 2:
		goto L86
	case 1:
		goto L87
	default:
		goto L88
	}
L83:
	;
	v392 = v386
	v394 = v385
	goto L81
L84:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+1)))
	if v387 != 0 {
		v336 = v387
		v337 = v337 + int32(1)
		v338 = v386
		v340 = v385
		goto L82
	} else {
		goto L93
	}
L85:
	;
	v385 = v384
	v386 = v338 + int32(1)
	goto L84
L86:
	;
	v384 = int32(1)
	goto L85
L87:
	;
	if v359 == v100 {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	if base.B2i32(v359 != int32(92))&base.B2i32(v359 != int32(34)) != 0 {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v385 = int32(1)
	v386 = v338 + int32(2)
	goto L84
L90:
	;
	v371 = base.I32_extend8_s(v336)
	goto L91
L91:
	;
	if base.B2i32(v371 == int32(32))|base.B2i32(base.Ui32((v371-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		v384 = v340
		goto L85
	} else {
		goto L92
	}
L92:
	;
	goto L86
L93:
	;
	goto L83
L94:
	;
	v416 = v392 + int32(2)
	goto L96
L95:
	;
	v416 = v392
	goto L96
L96:
	;
	v437 = v416
	goto L57
L97:
	;
	goto L55
L98:
	;
	v791 = int32(123)
	*(*uint16)(unsafe.Add(mBase, uint32(v790))) = uint16(v791)
	if v87 <= int32(0) {
		goto L144
	} else {
		goto L145
	}
L99:
	;
	v684 = int32(61)
	*(*uint16)(unsafe.Add(mBase, uint32(v664))) = uint16(v684)
	v690 = F_palloc(m, v664+(v666-v23)-int32(63))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L122
	}
L100:
	;
	v662 = F_palloc(m, v644)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L121
	}
L101:
	;
	v466 = v87 & int32(3)
	v467 = int32(0)
	v468 = int32(1)
	if base.Ui32(int32(4)) <= base.Ui32(v87) {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	goto L103
L103:
	;
	v638 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+64)) = uint8(v638)
	if v151 != 0 {
		v664 = v23 - int32(-64)
		v666 = v445
		goto L99
	} else {
		goto L120
	}
L104:
	;
	v589 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+64)) = uint8(v589)
	v594 = v576<<(uint(int32(1))%32) + v445
	if v151 == v589 {
		v644 = v594
		goto L100
	} else {
		goto L115
	}
L105:
	;
	v476 = v468
	v477 = v467
	v481 = int32(0)
	v483 = v467
	goto L108
L106:
	;
	v518 = v468
	v519 = v467
	v525 = v467
	goto L107
L107:
	;
	v538 = v518
	v539 = v519
	v542 = v467
	v545 = v525
	goto L112
L108:
	;
	v498 = v98 + v477<<(uint(int32(2))%32)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+8))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v498)))
	v501 = v500 * v476
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	v503 = v501 * v502
	v504 = v499 * v503
	v508 = v504 + (v503 + (v501 + (v476 + v483)))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v498)+12))
	v510 = v509 * v504
	v511 = int32(4)
	v512 = v477 + v511
	v514 = v481 + v511
	if v514 != v87&int32(2147483644) {
		v476 = v510
		v477 = v512
		v481 = v514
		v483 = v508
		goto L108
	} else {
		goto L110
	}
L109:
	;
	if v466 == int32(0) {
		v576 = v508
		goto L104
	} else {
		goto L111
	}
L110:
	;
	goto L109
L111:
	;
	v518 = v510
	v519 = v512
	v525 = v508
	goto L107
L112:
	;
	v558 = v538 + v545
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v98+v539<<(uint(int32(2))%32))))
	v564 = int32(1)
	v567 = v542 + v564
	if v567 != v466 {
		v538 = v562 * v538
		v539 = v539 + v564
		v542 = v567
		v545 = v558
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v576 = v558
	goto L104
L114:
	;
	goto L113
L115:
	;
	v599 = v23 - int32(-64)
	v600 = v589
	goto L116
L116:
	;
	v620 = v600 << (uint(int32(2)) % 32)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v98+v620)))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v620+v99)))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v624
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v624 + v622 - int32(1)
	v631 = F_pg_sprintf(m, v599, int32(_a_F_array_out_2), v23)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L118
	}
L117:
	;
	v664 = v634
	v666 = v594
	goto L99
L118:
	;
	v633 = F_strlen(m, v599)
	mBase = m.M
	v634 = v633 + v599
	v636 = v600 + int32(1)
	if v636 != v87 {
		v599 = v634
		v600 = v636
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v644 = v445
	goto L100
L121:
	;
	v778 = v662
	v790 = v662
	goto L98
L122:
	;
	v693 = v23 - int32(-64)
	if (v693^v690)&int32(3) != 0 {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	v768 = F_strlen(m, v690)
	mBase = m.M
	v778 = v690
	v790 = v768 + v690
	goto L98
L124:
	;
	goto L123
L125:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v748))) = uint8(v747)
	if v747&int32(255) == int32(0) {
		goto L124
	} else {
		goto L140
	}
L126:
	;
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v693))))
	v746 = v693
	v747 = v699
	v748 = v690
	goto L125
L127:
	;
	goto L128
L128:
	;
	if v693&int32(3) != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v703 = v693
	v705 = v690
	goto L132
L130:
	;
	v717 = v693
	v719 = v690
	goto L131
L131:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v717)))
	v724 = int32(-2139062144)
	if (int32(16843008)-v721|v721)&v724 != v724 {
		v746 = v717
		v747 = v721
		v748 = v719
		goto L125
	} else {
		goto L136
	}
L132:
	;
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703))))
	*(*uint8)(unsafe.Add(mBase, uint32(v705))) = uint8(v706)
	if v706 == int32(0) {
		goto L124
	} else {
		goto L134
	}
L133:
	;
	v717 = v713
	v719 = v711
	goto L131
L134:
	;
	v710 = int32(1)
	v711 = v705 + v710
	v713 = v703 + v710
	if v713&int32(3) != 0 {
		v703 = v713
		v705 = v711
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v729 = v717
	v730 = v721
	v731 = v719
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v731))) = v730
	v733 = int32(4)
	v734 = v731 + v733
	v736 = v729 + v733
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v729)+4))
	v741 = int32(-2139062144)
	if (int32(16843008)-v738|v738)&v741 == v741 {
		v729 = v736
		v730 = v738
		v731 = v734
		goto L137
	} else {
		goto L139
	}
L138:
	;
	v746 = v736
	v747 = v738
	v748 = v734
	goto L125
L139:
	;
	goto L138
L140:
	;
	v755 = v746
	v757 = v748
	goto L141
L141:
	;
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v757)+1)) = uint8(v758)
	v760 = int32(1)
	if v758 != 0 {
		v755 = v755 + v760
		v757 = v757 + v760
		goto L141
	} else {
		goto L143
	}
L142:
	;
	goto L124
L143:
	;
	goto L142
L144:
	;
	v804 = int32(1)
	v809 = v87 - v804
	v810 = int32(0)
	v812 = v790 + v804
	v813 = v810
	v817 = v810
	goto L147
L145:
	;
	v796 = v87 << (uint(int32(2)) % 32)
	if v796 == int32(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	base.MemoryFill(m, v23+int32(32), int32(0), v796)
	goto L144
L147:
	;
	if v809 <= v813 {
		v920 = v812
		goto L149
	} else {
		goto L150
	}
L148:
	;
	F_pfree(m, v163)
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L1
	} else {
		goto L204
	}
L149:
	;
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v817+v165))))
	if v941 == int32(1) {
		goto L162
	} else {
		goto L163
	}
L150:
	;
	v838 = (v87 + (v813 ^ int32(-1))) & int32(7)
	if v838 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v839 = v812
	v841 = v813
	v846 = int32(0)
	goto L154
L152:
	;
	v868 = v812
	v870 = v813
	goto L153
L153:
	;
	if base.Ui32(v87-int32(2)-v813) < base.Ui32(int32(7)) {
		v920 = v868
		goto L149
	} else {
		goto L157
	}
L154:
	;
	v859 = int32(123)
	*(*uint16)(unsafe.Add(mBase, uint32(v839))) = uint16(v859)
	v861 = int32(1)
	v862 = v841 + v861
	v864 = v839 + v861
	v866 = v846 + v861
	if v866 != v838 {
		v839 = v864
		v841 = v862
		v846 = v866
		goto L154
	} else {
		goto L156
	}
L155:
	;
	v868 = v864
	v870 = v862
	goto L153
L156:
	;
	goto L155
L157:
	;
	v891 = v868
	v893 = v870
	goto L158
L158:
	;
	v911 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v891)+8)) = uint8(v911)
	*(*int64)(unsafe.Add(mBase, uint32(v891))) = int64(8897841259083430779)
	v915 = int32(8)
	v916 = v891 + v915
	v918 = v893 + v915
	if v918 != v809 {
		v891 = v916
		v893 = v918
		goto L158
	} else {
		goto L160
	}
L159:
	;
	v920 = v916
	goto L149
L160:
	;
	goto L159
L161:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v163+v817<<(uint(int32(2))%32))))
	F_pfree(m, v1098)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L1
	} else {
		goto L193
	}
L162:
	;
	v944 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v920))) = uint16(v944)
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v163+v817<<(uint(int32(2))%32))))
	v952 = v920 + int32(1)
	v954 = v951
	goto L165
L163:
	;
	goto L164
L164:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v163+v817<<(uint(int32(2))%32))))
	if (v997^v920)&int32(3) != 0 {
		goto L175
	} else {
		goto L176
	}
L165:
	;
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v954))))
	if base.B2i32(v972 == int32(34))|base.B2i32(v972 == int32(92)) == int32(0) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v988))) = uint8(v972)
	v990 = int32(1)
	v952 = v988 + v990
	v954 = v954 + v990
	goto L165
L168:
	;
	if v972 != 0 {
		v988 = v952
		goto L167
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v984 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v952))) = uint8(v984)
	v988 = v952 + int32(1)
	goto L167
L171:
	;
	v980 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v952))) = uint16(v980)
	v1094 = v952 + int32(1)
	goto L161
L172:
	;
	v1072 = F_strlen(m, v920)
	mBase = m.M
	v1094 = v1072 + v920
	goto L161
L173:
	;
	goto L172
L174:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1052))) = uint8(v1051)
	if v1051&int32(255) == int32(0) {
		goto L173
	} else {
		goto L189
	}
L175:
	;
	v1003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v997))))
	v1050 = v997
	v1051 = v1003
	v1052 = v920
	goto L174
L176:
	;
	goto L177
L177:
	;
	if v997&int32(3) != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v1007 = v997
	v1009 = v920
	goto L181
L179:
	;
	v1021 = v997
	v1023 = v920
	goto L180
L180:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1021)))
	v1028 = int32(-2139062144)
	if (int32(16843008)-v1025|v1025)&v1028 != v1028 {
		v1050 = v1021
		v1051 = v1025
		v1052 = v1023
		goto L174
	} else {
		goto L185
	}
L181:
	;
	v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1009))) = uint8(v1010)
	if v1010 == int32(0) {
		goto L173
	} else {
		goto L183
	}
L182:
	;
	v1021 = v1017
	v1023 = v1015
	goto L180
L183:
	;
	v1014 = int32(1)
	v1015 = v1009 + v1014
	v1017 = v1007 + v1014
	if v1017&int32(3) != 0 {
		v1007 = v1017
		v1009 = v1015
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	v1033 = v1021
	v1034 = v1025
	v1035 = v1023
	goto L186
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1035))) = v1034
	v1037 = int32(4)
	v1038 = v1035 + v1037
	v1040 = v1033 + v1037
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1033)+4))
	v1045 = int32(-2139062144)
	if (int32(16843008)-v1042|v1042)&v1045 == v1045 {
		v1033 = v1040
		v1034 = v1042
		v1035 = v1038
		goto L186
	} else {
		goto L188
	}
L187:
	;
	v1050 = v1040
	v1051 = v1042
	v1052 = v1038
	goto L174
L188:
	;
	goto L187
L189:
	;
	v1059 = v1050
	v1061 = v1052
	goto L190
L190:
	;
	v1062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1059)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1061)+1)) = uint8(v1062)
	v1064 = int32(1)
	if v1062 != 0 {
		v1059 = v1059 + v1064
		v1061 = v1061 + v1064
		goto L190
	} else {
		goto L192
	}
L191:
	;
	goto L173
L192:
	;
	goto L191
L193:
	;
	if v809 < int32(0) {
		v1150 = v1094
		v1151 = v809
		goto L195
	} else {
		goto L196
	}
L194:
	;
	goto L148
L195:
	;
	if v1151 != int32(-1) {
		v812 = v1150
		v813 = v1151
		v817 = v817 + int32(1)
		goto L147
	} else {
		goto L203
	}
L196:
	;
	v1103 = v1094
	v1104 = v809
	goto L197
L197:
	;
	v1124 = v1104 << (uint(int32(2)) % 32)
	v1127 = v1124 + (v23 + int32(32))
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1127)))
	v1130 = v1128 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1127))) = v1130
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1124+v98)))
	if v1130 < v1133 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	goto L194
L199:
	;
	v1135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1103)+1)) = uint8(v1135)
	*(*uint8)(unsafe.Add(mBase, uint32(v1103))) = uint8(v100)
	v1150 = v1103 + int32(1)
	v1151 = v1104
	goto L195
L200:
	;
	goto L201
L201:
	;
	v1140 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1127))) = v1140
	v1142 = int32(125)
	*(*uint16)(unsafe.Add(mBase, uint32(v1103))) = uint16(v1142)
	v1144 = int32(1)
	if v1140 < v1104 {
		v1103 = v1103 + v1144
		v1104 = v1104 - v1144
		goto L197
	} else {
		goto L202
	}
L202:
	;
	goto L198
L203:
	;
	goto L194
L204:
	;
	F_pfree(m, v165)
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v1206 = v778
	goto L22
}
func F_array_prepend(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v12 == v2 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v16 = v15
	} else {
		v16 = v2
	}
	v17 = int32(1)
	v19 = F_fetch_array_arg_replace_nulls(m, l0, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
		switch v23 {
		case 0:
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(1)
			v64 = v17
			v66 = int32(12)
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
			v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+4)))
			v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+6)))
			v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73)+7)))
			v77 = F_array_set_element(m, v19+v66, int32(1), v10+v66, v16, v12, int32(-1), v74, v75, v76)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
				if v79 == int32(1) {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v82))) = v64
				} else {
				}
				m.G0 = v10 + int32(16)
				return v77
			}
		case 1:
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
			v27 = v25 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v27
			if v27 < v25 {
				v64 = v25
				v66 = int32(12)
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
				v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+4)))
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+6)))
				v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73)+7)))
				v77 = F_array_set_element(m, v19+v66, int32(1), v10+v66, v16, v12, int32(-1), v74, v75, v76)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
					if v79 == int32(1) {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v82))) = v64
					} else {
					}
					m.G0 = v10 + int32(16)
					return v77
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_array_prepend_0), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_array_prepend_1), int32(249), int32(_a_F_array_prepend_2))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(130))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_array_prepend_3), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_array_prepend_1), int32(259), int32(_a_F_array_prepend_2))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
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
func F_array_remove(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v4 == int32(1) {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v7)
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v19 = int32(1)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v22 = F_array_replace_internal(m, v14, v11, v12, int32(0), v19, v19, v21, l0)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return v22
			}
		}
	}
}
func F_array_reverse(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	v18 = m.G0
	v20 = v18 - int32(80)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
		if v27 <= int32(0) {
			v206 = v23
			m.G0 = v20 + int32(80)
			return v206
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
			if v30 < int32(2) {
				v206 = v23
				m.G0 = v20 + int32(80)
				return v206
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
				if v35 != 0 {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
					if v36 == v33 {
						v44 = v35
						v45 = v27
						v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v44)+8)))
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+10)))
						v48 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+11)))
						F_deconstruct_array(m, v23, v46, v47, v48, v20+int32(12), v20+int32(8), v20+int32(76))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
							v59 = base.I32_div_s(v57, v58)
							*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v59
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
							if int32(1) < v58 {
								v66 = base.I32_div_s(v58, int32(2))
								v71 = v61
								v72 = v62
								v73 = v59
								v77 = int32(0)
								for {
									if int32(0) < v73 {
										v89 = (v58 + (v77 ^ int32(-1))) * v73
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
										v97 = v89 + v90
										v99 = v92 + v89<<(uint(int32(2))%32)
										v101 = v71
										v102 = v72
										v108 = int32(0)
										for {
											v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
											v116 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
											*(*int32)(unsafe.Add(mBase, uint32(v102))) = v116
											v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
											*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v118)
											*(*int32)(unsafe.Add(mBase, uint32(v99))) = v115
											*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v114)
											v122 = int32(1)
											v124 = int32(4)
											v127 = v101 + v122
											v129 = v102 + v124
											v131 = v108 + v122
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
											if v131 < v132 {
												v97 = v97 + v122
												v99 = v99 + v124
												v101 = v127
												v102 = v129
												v108 = v131
												continue
											} else {
												break
											}
											break
										}
										v138 = v127
										v139 = v129
										v140 = v132
									} else {
										v138 = v71
										v139 = v72
										v140 = v73
									}
									v152 = v77 + int32(1)
									if v152 != v66 {
										v71 = v138
										v72 = v139
										v73 = v140
										v77 = v152
										continue
									} else {
										break
									}
									break
								}
								v154 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
								v155 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
								v160 = v155
								v161 = v154
							} else {
								v160 = v61
								v161 = v62
							}
							v174 = v23 + int32(16)
							v176 = v45 << (uint(int32(2)) % 32)
							v177 = int32(0)
							v178 = base.B2i32(v176 == v177)
							if v178 == v177 {
								base.MemoryCopy(m, v20+int32(48), v174, v176)
							} else {
							}
							if v178 == int32(0) {
								base.MemoryCopy(m, v20+int32(16), v176+v174, v176)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v58
							v195 = F_construct_md_array(m, v161, v160, v45, v20+int32(48), v20+int32(16), v33, v46, v47, v48)
							mBase = m.M
							v196 = m.ExcPending
							if v196 != 0 {
								return int32(0)
							} else {
								v197 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
								F_pfree(m, v197)
								mBase = m.M
								v199 = m.ExcPending
								if v199 != 0 {
									return int32(0)
								} else {
									v200 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
									F_pfree(m, v200)
									mBase = m.M
									v202 = m.ExcPending
									if v202 != 0 {
										return int32(0)
									} else {
										v206 = v195
										m.G0 = v20 + int32(80)
										return v206
									}
								}
							}
						}
					} else {
						v39 = F_lookup_type_cache(m, v33, int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v39
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
							v44 = v39
							v45 = v43
							v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v44)+8)))
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+10)))
							v48 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+11)))
							F_deconstruct_array(m, v23, v46, v47, v48, v20+int32(12), v20+int32(8), v20+int32(76))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
								v59 = base.I32_div_s(v57, v58)
								*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v59
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
								if int32(1) < v58 {
									v66 = base.I32_div_s(v58, int32(2))
									v71 = v61
									v72 = v62
									v73 = v59
									v77 = int32(0)
									for {
										if int32(0) < v73 {
											v89 = (v58 + (v77 ^ int32(-1))) * v73
											v90 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
											v97 = v89 + v90
											v99 = v92 + v89<<(uint(int32(2))%32)
											v101 = v71
											v102 = v72
											v108 = int32(0)
											for {
												v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
												*(*int32)(unsafe.Add(mBase, uint32(v102))) = v116
												v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
												*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v118)
												*(*int32)(unsafe.Add(mBase, uint32(v99))) = v115
												*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v114)
												v122 = int32(1)
												v124 = int32(4)
												v127 = v101 + v122
												v129 = v102 + v124
												v131 = v108 + v122
												v132 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
												if v131 < v132 {
													v97 = v97 + v122
													v99 = v99 + v124
													v101 = v127
													v102 = v129
													v108 = v131
													continue
												} else {
													break
												}
												break
											}
											v138 = v127
											v139 = v129
											v140 = v132
										} else {
											v138 = v71
											v139 = v72
											v140 = v73
										}
										v152 = v77 + int32(1)
										if v152 != v66 {
											v71 = v138
											v72 = v139
											v73 = v140
											v77 = v152
											continue
										} else {
											break
										}
										break
									}
									v154 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
									v155 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
									v160 = v155
									v161 = v154
								} else {
									v160 = v61
									v161 = v62
								}
								v174 = v23 + int32(16)
								v176 = v45 << (uint(int32(2)) % 32)
								v177 = int32(0)
								v178 = base.B2i32(v176 == v177)
								if v178 == v177 {
									base.MemoryCopy(m, v20+int32(48), v174, v176)
								} else {
								}
								if v178 == int32(0) {
									base.MemoryCopy(m, v20+int32(16), v176+v174, v176)
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v58
								v195 = F_construct_md_array(m, v161, v160, v45, v20+int32(48), v20+int32(16), v33, v46, v47, v48)
								mBase = m.M
								v196 = m.ExcPending
								if v196 != 0 {
									return int32(0)
								} else {
									v197 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
									F_pfree(m, v197)
									mBase = m.M
									v199 = m.ExcPending
									if v199 != 0 {
										return int32(0)
									} else {
										v200 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
										F_pfree(m, v200)
										mBase = m.M
										v202 = m.ExcPending
										if v202 != 0 {
											return int32(0)
										} else {
											v206 = v195
											m.G0 = v20 + int32(80)
											return v206
										}
									}
								}
							}
						}
					}
				} else {
					v39 = F_lookup_type_cache(m, v33, int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v39
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
						v44 = v39
						v45 = v43
						v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v44)+8)))
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+10)))
						v48 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+11)))
						F_deconstruct_array(m, v23, v46, v47, v48, v20+int32(12), v20+int32(8), v20+int32(76))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
							v59 = base.I32_div_s(v57, v58)
							*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v59
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
							if int32(1) < v58 {
								v66 = base.I32_div_s(v58, int32(2))
								v71 = v61
								v72 = v62
								v73 = v59
								v77 = int32(0)
								for {
									if int32(0) < v73 {
										v89 = (v58 + (v77 ^ int32(-1))) * v73
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
										v97 = v89 + v90
										v99 = v92 + v89<<(uint(int32(2))%32)
										v101 = v71
										v102 = v72
										v108 = int32(0)
										for {
											v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
											v116 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
											*(*int32)(unsafe.Add(mBase, uint32(v102))) = v116
											v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
											*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v118)
											*(*int32)(unsafe.Add(mBase, uint32(v99))) = v115
											*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v114)
											v122 = int32(1)
											v124 = int32(4)
											v127 = v101 + v122
											v129 = v102 + v124
											v131 = v108 + v122
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
											if v131 < v132 {
												v97 = v97 + v122
												v99 = v99 + v124
												v101 = v127
												v102 = v129
												v108 = v131
												continue
											} else {
												break
											}
											break
										}
										v138 = v127
										v139 = v129
										v140 = v132
									} else {
										v138 = v71
										v139 = v72
										v140 = v73
									}
									v152 = v77 + int32(1)
									if v152 != v66 {
										v71 = v138
										v72 = v139
										v73 = v140
										v77 = v152
										continue
									} else {
										break
									}
									break
								}
								v154 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
								v155 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
								v160 = v155
								v161 = v154
							} else {
								v160 = v61
								v161 = v62
							}
							v174 = v23 + int32(16)
							v176 = v45 << (uint(int32(2)) % 32)
							v177 = int32(0)
							v178 = base.B2i32(v176 == v177)
							if v178 == v177 {
								base.MemoryCopy(m, v20+int32(48), v174, v176)
							} else {
							}
							if v178 == int32(0) {
								base.MemoryCopy(m, v20+int32(16), v176+v174, v176)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v58
							v195 = F_construct_md_array(m, v161, v160, v45, v20+int32(48), v20+int32(16), v33, v46, v47, v48)
							mBase = m.M
							v196 = m.ExcPending
							if v196 != 0 {
								return int32(0)
							} else {
								v197 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
								F_pfree(m, v197)
								mBase = m.M
								v199 = m.ExcPending
								if v199 != 0 {
									return int32(0)
								} else {
									v200 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
									F_pfree(m, v200)
									mBase = m.M
									v202 = m.ExcPending
									if v202 != 0 {
										return int32(0)
									} else {
										v206 = v195
										m.G0 = v20 + int32(80)
										return v206
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_array_sample(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		if int32(0) < v16 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v20 = v19
		} else {
			v20 = int32(0)
		}
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v22 = int32(0)
		if base.B2i32(v21 < v22)|base.B2i32(v20 < v21) == v22 {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
			if v30 != 0 {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				if v31 == v28 {
					v38 = v30
					v40 = F_array_shuffle_n(m, v12, v21, int32(0), v28, v38)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v40
					}
				} else {
					v34 = F_lookup_type_cache(m, v28, int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v34
						v38 = v34
						v40 = F_array_shuffle_n(m, v12, v21, int32(0), v28, v38)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v40
						}
					}
				}
			} else {
				v34 = F_lookup_type_cache(m, v28, int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v34
					v38 = v34
					v40 = F_array_shuffle_n(m, v12, v21, int32(0), v28, v38)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v40
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v20
					F_errmsg(m, int32(_a_F_array_sample_0), v9)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_array_sample_1), int32(1750), int32(_a_F_array_sample_2))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
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
func F_array_slice_size(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v395 int32
	_ = v395
	var v412 int32
	_ = v412
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v563 int32
	_ = v563
	var v576 int32
	_ = v576
	v10 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(128)
	m.G0 = v15
	v18 = v15 + int32(96)
	if l2 <= v10 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v95 = int32(0)
	if l1|base.B2i32(l7 <= v95) == v95 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	goto L1
L3:
	;
	if l2 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v35 = v10
	v38 = v10
	goto L7
L5:
	;
	v72 = v10
	goto L6
L6:
	;
	v77 = v72 << (uint(int32(2)) % 32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77+l6)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77+l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v18+v77))) = v80 - v82 + int32(1)
	goto L2
L7:
	;
	v39 = int32(2)
	v40 = v35 << (uint(v39) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40+l6)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40+l5)))
	v47 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18+v40))) = v43 - v45 + v47
	v51 = v40 | int32(4)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51+l6)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51+l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v18+v51))) = v54 - v56 + v47
	v62 = v35 + v39
	v64 = v38 + v39
	if v64 != l2&int32(2147483646) {
		v35 = v62
		v38 = v64
		goto L7
	} else {
		goto L9
	}
L8:
	;
	if l2&int32(1) == int32(0) {
		goto L2
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v72 = v62
	goto L6
L11:
	;
	m.G0 = v15 + int32(128)
	return v576
L12:
	;
	v100 = F_ArrayGetNItemsSafe(m, l2, v18)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v122 = int32(0)
	v132 = l2 - int32(1)
	if v132 < v122 {
		v210 = v122
		goto L22
	} else {
		goto L23
	}
L15:
	;
	return int32(0)
L16:
	;
	switch l8 - int32(99) {
	case 0:
		v120 = l7
		goto L17
	case 1:
		goto L19
	default:
		goto L18
	case 6:
		goto L20
	}
L17:
	;
	v576 = v100 * v120
	goto L11
L18:
	;
	v120 = (l7 + int32(1)) & int32(-2)
	goto L17
L19:
	;
	v576 = (l7 + int32(7)) & int32(-8) * v100
	goto L11
L20:
	;
	v576 = (l7 + int32(3)) & int32(-4) * v100
	goto L11
L21:
	;
	v216 = F_array_seek(m, l0, v122, l1, v210, l7, l8)
	mBase = m.M
	v218 = v15 - int32(-64)
	v222 = int32(2)
	v223 = l2 << (uint(v222) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v218+v223-int32(4)))) = int32(1)
	v230 = l2 - v222
	if v230 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	goto L21
L23:
	;
	v135 = int32(1)
	if v132 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v144 = v132
	v145 = v135
	v146 = v122
	v151 = v122
	goto L27
L25:
	;
	v187 = v132
	v188 = v135
	v189 = v122
	goto L26
L26:
	;
	v196 = v187 << (uint(int32(2)) % 32)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l5+v196)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v196+l4)))
	v210 = (v198-v200)*v188 + v189
	goto L22
L27:
	;
	v152 = int32(2)
	v153 = v144 << (uint(v152) % 32)
	v155 = v153 - int32(4)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l5+v155)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l4+v155)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v153+l3)))
	v163 = v162 * v145
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v153+l5)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v153+l4)))
	v172 = (v157-v159)*v163 + ((v166-v168)*v145 + v146)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l3+v155)))
	v175 = v174 * v163
	v177 = v144 - v152
	v179 = v151 + v152
	if v179 != l2&int32(-2) {
		v144 = v177
		v145 = v175
		v146 = v172
		v151 = v179
		goto L27
	} else {
		goto L29
	}
L28:
	;
	if l2&int32(1) == int32(0) {
		v210 = v172
		goto L22
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v187 = v177
	v188 = v175
	v189 = v172
	goto L26
L31:
	;
	v288 = v15 + int32(32)
	v290 = v15 + int32(96)
	v291 = int32(0)
	v297 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v288+l2<<(uint(v297)%32)-int32(4)))) = v291
	v305 = l2 - v297
	if v291 <= v305 {
		goto L42
	} else {
		goto L43
	}
L32:
	;
	goto L31
L33:
	;
	if l2&int32(1) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v241 = v223 - int32(4)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l3+v241)))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v218+v241)))
	*(*int32)(unsafe.Add(mBase, uint32(v218+v230<<(uint(int32(2))%32)))) = v243 * v245
	v250 = l2 - int32(3)
	goto L36
L35:
	;
	v250 = v230
	goto L36
L36:
	;
	if v230 == int32(0) {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v256 = v250
	goto L38
L38:
	;
	v259 = int32(2)
	v260 = v256 << (uint(v259) % 32)
	v263 = v260 + int32(4)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l3+v263)))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v263+v218)))
	v268 = v265 * v267
	*(*int32)(unsafe.Add(mBase, uint32(v218+v260))) = v268
	v271 = v256 - int32(1)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l3+v260)))
	*(*int32)(unsafe.Add(mBase, uint32(v218+v271<<(uint(v259)%32)))) = v276 * v268
	if v271 != 0 {
		v256 = v256 - v259
		goto L38
	} else {
		goto L40
	}
L39:
	;
	goto L32
L40:
	;
	goto L39
L41:
	;
	v412 = l2 << (uint(int32(2)) % 32)
	if v412 != 0 {
		goto L57
	} else {
		goto L58
	}
L42:
	;
	v312 = v305
	v316 = v291
	goto L45
L43:
	;
	goto L44
L44:
	;
	goto L41
L45:
	;
	v319 = v312 << (uint(int32(2)) % 32)
	v320 = v288 + v319
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v218+v319)))
	v323 = int32(1)
	v324 = v322 - v323
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = v324
	v327 = v312 + v323
	if l2 <= v327 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L44
L47:
	;
	v395 = int32(1)
	if int32(0) < v312 {
		v312 = v312 - v395
		v316 = v316 + v395
		goto L45
	} else {
		goto L56
	}
L48:
	;
	if v316&int32(1) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v333 = int32(2)
	v334 = v327 << (uint(v333) % 32)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v290+v334)))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v218+v334)))
	v342 = v324 - (v336-int32(1))*v340
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = v342
	v346 = v312 + v333
	v347 = v342
	goto L51
L50:
	;
	v346 = v327
	v347 = v324
	goto L51
L51:
	;
	if v316 == int32(0) {
		goto L47
	} else {
		goto L52
	}
L52:
	;
	v354 = v346
	v355 = v347
	goto L53
L53:
	;
	v360 = int32(2)
	v361 = v354 << (uint(v360) % 32)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v290+v361)))
	v364 = int32(1)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v218+v361)))
	v369 = v355 - (v363-v364)*v367
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = v369
	v372 = v361 + int32(4)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v290+v372)))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v218+v372)))
	v380 = v369 - (v374-v364)*v378
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = v380
	v383 = v354 + v360
	if v383 != l2 {
		v354 = v383
		v355 = v380
		goto L53
	} else {
		goto L55
	}
L54:
	;
	goto L47
L55:
	;
	goto L54
L56:
	;
	goto L46
L57:
	;
	base.MemoryFill(m, v15, int32(0), v412)
	goto L59
L58:
	;
	goto L59
L59:
	;
	v425 = v216
	v426 = v210
	v427 = l2 - int32(1)
	v431 = v10
	goto L60
L60:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(32)+v427<<(uint(int32(2))%32))))
	if v438 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v576 = v505
	goto L11
L62:
	;
	if l1 != 0 {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	v443 = v425
	v444 = v426
	goto L62
L64:
	;
	goto L65
L65:
	;
	v442 = F_array_seek(m, v425, v426, l1, v438, l7, l8)
	mBase = m.M
	v443 = v442
	v444 = v438 + v426
	goto L62
L66:
	;
	v509 = v15 + int32(96)
	if l2 <= int32(0) {
		goto L91
	} else {
		goto L92
	}
L67:
	;
	v446 = base.I32_div_s(v444, int32(8))
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v446))))
	if int32(base.Ui32(v448)>>(uint(v444&int32(7))%32))&int32(1) == int32(0) {
		v503 = v443
		v505 = v431
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if int32(0) < l7 {
		v486 = l7
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	switch l8 - int32(99) {
	case 0:
		v499 = v486
		goto L86
	case 1:
		goto L88
	default:
		goto L87
	case 6:
		goto L89
	}
L72:
	;
	if l7 == int32(-1) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v458 == int32(1) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v482 = F_strlen(m, v443)
	mBase = m.M
	v486 = v482 + int32(1)
	goto L71
L76:
	;
	v462 = int32(18)
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443)+1)))
	if v464 == v462 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	v475 = int32(1)
	if v458&v475 != 0 {
		v486 = int32(base.Ui32(v458) >> (uint(v475) % 32))
		goto L71
	} else {
		goto L85
	}
L79:
	;
	v467 = v462
	goto L81
L80:
	;
	v467 = int32(2)
	goto L81
L81:
	;
	if base.Ui32((v464-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v474 = int32(6)
	goto L84
L83:
	;
	v474 = v467
	goto L84
L84:
	;
	v486 = v474
	goto L71
L85:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v443)))
	v486 = int32(base.Ui32(v479) >> (uint(int32(2)) % 32))
	goto L71
L86:
	;
	v503 = v443 + v499
	v505 = v499 + v431
	goto L66
L87:
	;
	v499 = (v486 + int32(1)) & int32(-2)
	goto L86
L88:
	;
	v499 = (v486 + int32(7)) & int32(-8)
	goto L86
L89:
	;
	v499 = (v486 + int32(3)) & int32(-4)
	goto L86
L90:
	;
	if v563 != int32(-1) {
		v425 = v503
		v426 = v444 + int32(1)
		v427 = v563
		v431 = v505
		goto L60
	} else {
		goto L105
	}
L91:
	;
	v563 = int32(-1)
	goto L90
L92:
	;
	goto L93
L93:
	;
	v515 = int32(1)
	v516 = l2 - v515
	v518 = v516 << (uint(int32(2)) % 32)
	v519 = v15 + v518
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v509+v518)))
	v525 = base.I32_rem_s(v520+v515, v524)
	*(*int32)(unsafe.Add(mBase, uint32(v519))) = v525
	if v516 != 0 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v563 = v553
	goto L90
L95:
	;
	v527 = v516
	v530 = v525
	goto L98
L96:
	;
	goto L97
L97:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v551 != 0 {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	if v530 != 0 {
		v553 = v527
		goto L94
	} else {
		goto L100
	}
L99:
	;
	goto L97
L100:
	;
	v532 = int32(1)
	v533 = v527 - v532
	v535 = v533 << (uint(int32(2)) % 32)
	v536 = v15 + v535
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v509+v535)))
	v542 = base.I32_rem_s(v537+v532, v541)
	*(*int32)(unsafe.Add(mBase, uint32(v536))) = v542
	if v533 != 0 {
		v527 = v533
		v530 = v542
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v552 = int32(0)
	goto L104
L103:
	;
	v552 = int32(-1)
	goto L104
L104:
	;
	v553 = v552
	goto L94
L105:
	;
	goto L61
}
func F_array_subscript_assign_slice(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v495 int32
	_ = v495
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v883 int32
	_ = v883
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1136 int32
	_ = v1136
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1172 int32
	_ = v1172
	var v1178 int32
	_ = v1178
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1342 int32
	_ = v1342
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1427 int32
	_ = v1427
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1447 int32
	_ = v1447
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1460 int32
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1473 int32
	_ = v1473
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1518 int32
	_ = v1518
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1539 int32
	_ = v1539
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1570 int32
	_ = v1570
	var v1578 int32
	_ = v1578
	var v1585 int32
	_ = v1585
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1600 int32
	_ = v1600
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1609 int32
	_ = v1609
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1651 int32
	_ = v1651
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1668 int32
	_ = v1668
	var v1684 int32
	_ = v1684
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1742 int32
	_ = v1742
	var v1748 int32
	_ = v1748
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1766 int32
	_ = v1766
	var v1771 int32
	_ = v1771
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1804 int32
	_ = v1804
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1893 int32
	_ = v1893
	var v1958 int32
	_ = v1958
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1974 int32
	_ = v1974
	var v1977 int32
	_ = v1977
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2012 int32
	_ = v2012
	var v2014 int32
	_ = v2014
	var v2020 int32
	_ = v2020
	var v2022 int32
	_ = v2022
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2036 int32
	_ = v2036
	var v2038 int32
	_ = v2038
	var v2040 int32
	_ = v2040
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2058 int32
	_ = v2058
	var v2061 int32
	_ = v2061
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2094 int32
	_ = v2094
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2109 int32
	_ = v2109
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2119 int32
	_ = v2119
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2150 int32
	_ = v2150
	var v2152 int32
	_ = v2152
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2167 int32
	_ = v2167
	var v2170 int32
	_ = v2170
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2189 int32
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2202 int32
	_ = v2202
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2215 int32
	_ = v2215
	var v2218 int32
	_ = v2218
	var v2221 int32
	_ = v2221
	var v2226 int32
	_ = v2226
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2253 int32
	_ = v2253
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2265 int32
	_ = v2265
	var v2267 int32
	_ = v2267
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2273 int32
	_ = v2273
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2285 int32
	_ = v2285
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2310 int32
	_ = v2310
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2324 int32
	_ = v2324
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2335 int32
	_ = v2335
	var v2350 int32
	_ = v2350
	var v2355 int32
	_ = v2355
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2419 int32
	_ = v2419
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2428 int32
	_ = v2428
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2445 int32
	_ = v2445
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2485 int32
	_ = v2485
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2493 int32
	_ = v2493
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2500 int32
	_ = v2500
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2530 int32
	_ = v2530
	var v2536 int32
	_ = v2536
	var v2537 int32
	_ = v2537
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2544 int32
	_ = v2544
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2555 int32
	_ = v2555
	var v2568 int32
	_ = v2568
	var v2574 int32
	_ = v2574
	var v2636 int32
	_ = v2636
	var v2637 int32
	_ = v2637
	var v2640 int32
	_ = v2640
	var v2642 int32
	_ = v2642
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2647 int32
	_ = v2647
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2652 int32
	_ = v2652
	var v2681 int32
	_ = v2681
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2685 int32
	_ = v2685
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2702 int32
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2704 int32
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2707 int32
	_ = v2707
	var v2708 int32
	_ = v2708
	var v2710 int32
	_ = v2710
	var v2711 int32
	_ = v2711
	var v2715 int32
	_ = v2715
	var v2717 int32
	_ = v2717
	var v2744 int32
	_ = v2744
	var v2745 int32
	_ = v2745
	var v2746 int32
	_ = v2746
	var v2748 int32
	_ = v2748
	var v2754 int32
	_ = v2754
	var v2755 int32
	_ = v2755
	var v2758 int32
	_ = v2758
	var v2759 int32
	_ = v2759
	var v2760 int32
	_ = v2760
	var v2762 int32
	_ = v2762
	var v2767 int32
	_ = v2767
	var v2768 int32
	_ = v2768
	var v2771 int32
	_ = v2771
	var v2772 int32
	_ = v2772
	var v2773 int32
	_ = v2773
	var v2782 int32
	_ = v2782
	var v2784 int32
	_ = v2784
	var v2822 int32
	_ = v2822
	var v2823 int32
	_ = v2823
	var v2824 int32
	_ = v2824
	var v2825 int32
	_ = v2825
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2830 int32
	_ = v2830
	var v2831 int32
	_ = v2831
	var v2835 int32
	_ = v2835
	var v2837 int32
	_ = v2837
	var v2864 int32
	_ = v2864
	var v2865 int32
	_ = v2865
	var v2866 int32
	_ = v2866
	var v2868 int32
	_ = v2868
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2878 int32
	_ = v2878
	var v2879 int32
	_ = v2879
	var v2880 int32
	_ = v2880
	var v2882 int32
	_ = v2882
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2891 int32
	_ = v2891
	var v2892 int32
	_ = v2892
	var v2893 int32
	_ = v2893
	var v2902 int32
	_ = v2902
	var v2907 int32
	_ = v2907
	var v2956 int32
	_ = v2956
	var v2974 int32
	_ = v2974
	v4 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40)+4)))
	if v4 < v41 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v64 = v40 + int32(12)
	v66 = v40 + int32(36)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+44)))
	v71 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40)+6)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+8)))
	v73 = int32(*(*int8)(unsafe.Add(mBase, uint32(v40)+9)))
	v75 = m.G0
	v77 = v75 - int32(272)
	m.G0 = v77
	if v70 != 0 {
		v2956 = v60
		goto L11
	} else {
		goto L12
	}
L3:
	;
	if v36&int32(1) != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v36&int32(1) == int32(0) {
		v60 = v38
		v61 = v41
		goto L2
	} else {
		goto L8
	}
L6:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+44)))
	if v46 == int32(0) {
		v60 = v38
		v61 = v41
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v54 = F_construct_empty_array(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v57)
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40)+4)))
	v60 = v54
	v61 = v59
	goto L2
L11:
	;
	m.G0 = v77 + int32(272)
	v2974 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2974))) = v2956
	goto L1
L12:
	;
	if v61 <= int32(0) {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	v974 = v789 + v772 + v961 - v958
	v975 = F_palloc0(m, v974)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L9
	} else {
		goto L203
	}
L14:
	;
	v909 = v905 << (uint(int32(3)) % 32)
	if v790 != 0 {
		goto L182
	} else {
		goto L183
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L9
	} else {
		goto L178
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L9
	} else {
		goto L174
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L9
	} else {
		goto L170
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L9
	} else {
		goto L166
	}
L19:
	;
	v661 = v77 + int32(112)
	v662 = F_ArrayGetNItemsSafe(m, v85, v661)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L9
	} else {
		goto L132
	}
L20:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v565 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L21:
	;
	if v85 <= v460 {
		v633 = v244
		v640 = v4
		goto L19
	} else {
		goto L104
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L9
	} else {
		goto L100
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L9
	} else {
		goto L96
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L9
	} else {
		goto L92
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L9
	} else {
		goto L88
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L9
	} else {
		goto L84
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L9
	} else {
		goto L79
	}
L28:
	;
	v81 = F_pg_detoast_datum(m, v60)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L9
	} else {
		goto L75
	}
L31:
	;
	v83 = F_pg_detoast_datum(m, v69)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v85 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	F_deconstruct_array(m, v83, v71, v72, v73, v77+int32(240), v77+int32(208), v77+int32(176))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if base.B2i32(v85 < v62)|base.B2i32(base.Ui32(int32(7)) <= base.Ui32(v85)) != 0 {
		goto L24
	} else {
		goto L50
	}
L36:
	;
	if int32(0) < v62 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v102 = v4
	goto L40
L38:
	;
	goto L39
L39:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v77)+176))
	v203 = v77 + int32(112)
	v204 = F_ArrayGetNItemsSafe(m, v62, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L9
	} else {
		goto L47
	}
L40:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v67))))
	if v134 != int32(1) {
		goto L27
	} else {
		goto L42
	}
L41:
	;
	goto L39
L42:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v68))))
	if v138 == int32(0) {
		goto L27
	} else {
		goto L43
	}
L43:
	;
	v142 = v102 << (uint(int32(2)) % 32)
	v145 = v142 + (v77 + int32(112))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v142+v64)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v142+v66)))
	v150 = v147 - v149
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v150
	if base.B2i32(v150 < v147)^base.B2i32(int32(0) < v149) != 0 {
		goto L26
	} else {
		goto L44
	}
L44:
	;
	v157 = v150 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v157
	if v157 < v150 {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77+int32(80)+v142))) = v149
	v165 = v102 + int32(1)
	if v165 != v62 {
		v102 = v165
		goto L40
	} else {
		goto L46
	}
L46:
	;
	goto L41
L47:
	;
	if v201 < v204 {
		goto L25
	} else {
		goto L48
	}
L48:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v77)+240))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v77)+208))
	v211 = F_construct_md_array(m, v207, v208, v62, v203, v77+int32(80), v88, v71, v72, v73)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	v2956 = v211
	goto L11
L50:
	;
	v218 = v81 + int32(16)
	v220 = v85 << (uint(int32(2)) % 32)
	v221 = int32(0)
	v222 = base.B2i32(v220 == v221)
	if v222 == v221 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	base.MemoryCopy(m, v77+int32(112), v218, v220)
	goto L53
L52:
	;
	goto L53
L53:
	;
	if v222 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	base.MemoryCopy(m, v77+int32(80), v218+v232<<(uint(int32(2))%32), v220)
	goto L56
L55:
	;
	goto L56
L56:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v238 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v244 = base.B2i32(v241 != int32(0))
	goto L59
L58:
	;
	v244 = int32(1)
	goto L59
L59:
	;
	if v85 == int32(1) {
		goto L20
	} else {
		goto L60
	}
L60:
	;
	v247 = int32(0)
	if v62 <= v247 {
		v460 = v247
		goto L21
	} else {
		goto L61
	}
L61:
	;
	v253 = v4
	goto L62
L62:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253+v68))))
	if v285 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v460 = v62
	goto L21
L64:
	;
	v289 = v253 << (uint(int32(2)) % 32)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v77+int32(80)+v289)))
	*(*int32)(unsafe.Add(mBase, uint32(v66+v289))) = v294
	goto L66
L65:
	;
	goto L66
L66:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253+v67))))
	if v298 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v322 = v253 << (uint(int32(2)) % 32)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v66+v322)))
	if v320 < v324 {
		goto L23
	} else {
		goto L71
	}
L68:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v64+v253<<(uint(int32(2))%32))))
	v320 = v304
	goto L67
L69:
	;
	goto L70
L70:
	;
	v306 = v253 << (uint(int32(2)) % 32)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v77+int32(112)+v306)))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v77+int32(80)+v306)))
	v318 = v311 + v315 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64+v306))) = v318
	v320 = v318
	goto L67
L71:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v77+int32(80)+v322)))
	if v324 < v329 {
		goto L22
	} else {
		goto L72
	}
L72:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v77+int32(112)+v322)))
	if v334+v329 <= v320 {
		goto L22
	} else {
		goto L73
	}
L73:
	;
	v338 = v253 + int32(1)
	if v338 != v62 {
		v253 = v338
		goto L62
	} else {
		goto L74
	}
L74:
	;
	goto L63
L75:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_0), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2855), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_3), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	F_errdetail(m, int32(_a_F_array_subscript_assign_slice_4), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L9
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2888), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L9
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L9
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(268435455)
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_5), v77)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L9
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2896), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L9
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L9
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_6), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L9
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2905), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L9
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L9
	} else {
		goto L93
	}
L93:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_7), int32(0))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L9
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2915), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L9
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L9
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_8), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L9
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2985), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L9
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L9
	} else {
		goto L101
	}
L101:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_9), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L9
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2990), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L9
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	v495 = v460
	goto L105
L105:
	;
	v528 = v495 << (uint(int32(2)) % 32)
	v529 = v66 + v528
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v77+int32(80)+v528)))
	*(*int32)(unsafe.Add(mBase, uint32(v529))) = v533
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v77+int32(112)+v528)))
	v542 = v533 + v539 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v528+v64))) = v542
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v529)))
	if v542 < v544 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L9
	} else {
		goto L110
	}
L107:
	;
	goto L106
L108:
	;
	v547 = v495 + int32(1)
	if v85 != v547 {
		v495 = v547
		goto L105
	} else {
		goto L109
	}
L109:
	;
	v633 = v244
	v640 = v4
	goto L19
L110:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L9
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_8), int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L9
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(3000), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L9
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v77)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v568
	goto L116
L115:
	;
	goto L116
L116:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v570 == int32(1) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v580 < v581 {
		goto L18
	} else {
		goto L121
	}
L118:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v580 = v573
	goto L117
L119:
	;
	goto L120
L120:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v77)+112))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v77)+80))
	v578 = v574 + v575 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v578
	v580 = v578
	goto L117
L121:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v77)+80))
	if v583 <= v581 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v607 = v602 + v604
	if v580 < v607 {
		v633 = v605
		v640 = v606
		goto L19
	} else {
		goto L128
	}
L123:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v77)+112))
	v602 = v585
	v604 = v583
	v605 = v244
	v606 = v4
	goto L122
L124:
	;
	goto L125
L125:
	;
	v586 = v583 - v581
	if base.B2i32(v586 < v583)^base.B2i32(int32(0) < v581) != 0 {
		goto L17
	} else {
		goto L126
	}
L126:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v77)+112))
	v592 = v591 + v586
	*(*int32)(unsafe.Add(mBase, uint32(v77)+112)) = v592
	if base.B2i32(v586 < int32(0)) != base.B2i32(v592 < v591) {
		goto L17
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+80)) = v581
	v602 = v592
	v604 = v581
	v605 = base.B2i32(int32(1) < v586) | v244
	v606 = v586
	goto L122
L128:
	;
	v611 = v580 - v607
	if base.B2i32(int32(0) < v607)^base.B2i32(v611 < v580) != 0 {
		goto L16
	} else {
		goto L129
	}
L129:
	;
	v615 = v611 + int32(1)
	if v615 < v611 {
		goto L16
	} else {
		goto L130
	}
L130:
	;
	v617 = v602 + v615
	*(*int32)(unsafe.Add(mBase, uint32(v77)+112)) = v617
	if base.B2i32(v615 < int32(0)) != base.B2i32(v617 < v602) {
		goto L16
	} else {
		goto L131
	}
L131:
	;
	v633 = base.B2i32(int32(1) < v615) | v605
	v640 = v606
	goto L19
L132:
	;
	F_ArrayCheckBounds(m, v85, v661, v77+int32(80))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L9
	} else {
		goto L133
	}
L133:
	;
	v669 = v77 + int32(48)
	v670 = int32(0)
	if v85 <= v670 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v746 = F_ArrayGetNItemsSafe(m, v85, v669)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L9
	} else {
		goto L144
	}
L135:
	;
	goto L134
L136:
	;
	if v85 != int32(1) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v686 = v670
	v689 = v670
	goto L140
L138:
	;
	v723 = v670
	goto L139
L139:
	;
	v728 = v723 << (uint(int32(2)) % 32)
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v728+v64)))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v728+v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v669+v728))) = v731 - v733 + int32(1)
	goto L135
L140:
	;
	v690 = int32(2)
	v691 = v686 << (uint(v690) % 32)
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v691+v64)))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v691+v66)))
	v698 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v669+v691))) = v694 - v696 + v698
	v702 = v691 | int32(4)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v702+v64)))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v702+v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v669+v702))) = v705 - v707 + v698
	v713 = v686 + v690
	v715 = v689 + v690
	if v715 != v85&int32(2147483646) {
		v686 = v713
		v689 = v715
		goto L140
	} else {
		goto L142
	}
L141:
	;
	if v85&int32(1) == int32(0) {
		goto L135
	} else {
		goto L143
	}
L142:
	;
	goto L141
L143:
	;
	v723 = v713
	goto L139
L144:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v750 = v83 + int32(16)
	v751 = F_ArrayGetNItemsSafe(m, v748, v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L9
	} else {
		goto L145
	}
L145:
	;
	if v751 < v746 {
		goto L15
	} else {
		goto L146
	}
L146:
	;
	v755 = v85 << (uint(int32(3)) % 32)
	if v633&int32(1) != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v777 = v775 << (uint(int32(3)) % 32)
	if v774 != 0 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v761 = base.I32_div_s(v662+int32(7), int32(8))
	v766 = (v755 + v761 + int32(23)) & int32(-8)
	v771 = v766
	v772 = v766
	goto L147
L149:
	;
	goto L150
L150:
	;
	v771 = v4
	v772 = (v755 + int32(23)) & int32(120)
	goto L147
L151:
	;
	v782 = v774
	goto L153
L152:
	;
	v782 = (v777 + int32(23)) & int32(-8)
	goto L153
L153:
	;
	v783 = v83 + v782
	v784 = int32(0)
	if v774 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v787 = v777 + v750
	goto L156
L155:
	;
	v787 = v784
	goto L156
L156:
	;
	v788 = F_array_seek(m, v783, v784, v787, v746, v71, v73)
	mBase = m.M
	v789 = v788 - v783
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v790 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v826 = F_array_slice_size(m, v81+v819, v816, v85, v77+int32(112), v77+int32(80), v66, v64, v71, v73)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L9
	} else {
		goto L165
	}
L158:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v802 = (v796<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v803 = int32(base.Ui32(v793)>>(uint(int32(2))%32)) - v802
	if int32(1) < v85 {
		v816 = int32(0)
		v818 = v803
		v819 = v802
		goto L157
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v807 = int32(2)
	v809 = int32(base.Ui32(v806)>>(uint(v807)%32)) - v790
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v85 < v807 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v905 = v796
	v906 = v803
	v907 = v802
	goto L14
L162:
	;
	v905 = v810
	v906 = v809
	v907 = v790
	goto L14
L163:
	;
	goto L164
L164:
	;
	v816 = v218 + v810<<(uint(int32(3))%32)
	v818 = v809
	v819 = v790
	goto L157
L165:
	;
	v828 = int32(0)
	v957 = v828
	v958 = v826
	v959 = v828
	v961 = v818
	v962 = v819
	v963 = int32(0)
	v965 = int32(1)
	v967 = v4
	v971 = v828
	goto L13
L166:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L9
	} else {
		goto L167
	}
L167:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_8), int32(0))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L9
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2940), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L9
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L9
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = int32(268435455)
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_5), v77+int32(16))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L9
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2950), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L9
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L174:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L9
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+32)) = int32(268435455)
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_5), v77+int32(32))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L9
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2965), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L9
	} else {
		goto L177
	}
L177:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L178:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L9
	} else {
		goto L179
	}
L179:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_6), int32(0))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L9
	} else {
		goto L180
	}
L180:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(3017), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L9
	} else {
		goto L181
	}
L181:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L182:
	;
	v914 = v790
	goto L184
L183:
	;
	v914 = (v909 + int32(23)) & int32(-8)
	goto L184
L184:
	;
	v915 = v81 + v914
	v916 = int32(0)
	if v790 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v919 = v909 + v218
	goto L187
L186:
	;
	v919 = v916
	goto L187
L187:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v218+v905<<(uint(int32(2))%32))))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v924 < v923 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v926 = v923
	goto L190
L189:
	;
	v926 = v924
	goto L190
L190:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	v928 = v927 + v923
	if v926 < v928 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v930 = v926
	goto L193
L192:
	;
	v930 = v928
	goto L193
L193:
	;
	v931 = v930 - v923
	v932 = F_array_seek(m, v915, v916, v919, v931, v71, v73)
	mBase = m.M
	v933 = v932 - v915
	v936 = v928 - int32(1)
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v936 < v937 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v939 = v936
	goto L196
L195:
	;
	v939 = v937
	goto L196
L196:
	;
	if v926 <= v939 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v944 = v939 - v926 + int32(1)
	v945 = F_array_seek(m, v933+v915, v931, v919, v944, v71, v73)
	mBase = m.M
	v947 = v945 - v932
	v948 = v944
	goto L199
L198:
	;
	v947 = int32(0)
	v948 = v4
	goto L199
L199:
	;
	v950 = v939 + int32(1)
	if v923 < v950 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v952 = v950
	goto L202
L201:
	;
	v952 = v923
	goto L202
L202:
	;
	v957 = v933
	v958 = v947
	v959 = v928 - v952
	v961 = v906
	v962 = v907
	v963 = v931
	v965 = v4
	v967 = v948
	v971 = v906 - (v933 + v947)
	goto L13
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v975)+8)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v975)+4)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v975))) = v974 << (uint(int32(2)) % 32)
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v975)+12)) = v982
	v985 = v975 + int32(16)
	if v222 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	base.MemoryCopy(m, v985, v77+int32(112), v220)
	goto L206
L205:
	;
	goto L206
L206:
	;
	if v222 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	base.MemoryCopy(m, v985+v220, v77+int32(80), v220)
	goto L209
L208:
	;
	goto L209
L209:
	;
	if v965 != 0 {
		goto L212
	} else {
		goto L213
	}
L210:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2902))) = uint8(v2907)
	v2956 = v975
	goto L11
L211:
	;
	v2822 = base.I32_div_s(v2028, int32(8))
	v2823 = v1024 + v2822
	v2824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2823))))
	v2825 = v2112
	v2827 = v2097
	v2828 = v2109
	v2830 = v2113
	v2831 = v2823
	v2835 = int32(1) << (uint(v2028&int32(7)) % 32)
	v2837 = v2824
	goto L516
L212:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v997 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L213:
	;
	goto L214
L214:
	;
	v2164 = v81 + v962
	v2165 = v975 + v772
	if v957 != 0 {
		goto L408
	} else {
		goto L409
	}
L215:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v1007 = (v1000<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L217
L216:
	;
	v1007 = v997
	goto L217
L217:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v1008 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v1018 = (v1011<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L220
L219:
	;
	v1018 = v1008
	goto L220
L220:
	;
	if v997 != 0 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v1024 = v218 + v1020<<(uint(int32(3))%32)
	goto L223
L222:
	;
	v1024 = int32(0)
	goto L223
L223:
	;
	if v1008 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v1030 = v750 + v1026<<(uint(int32(3))%32)
	goto L226
L225:
	;
	v1030 = int32(0)
	goto L226
L226:
	;
	v1031 = v1007 + v81
	v1033 = v85 << (uint(int32(3)) % 32)
	v1034 = v1033 + v985
	if v771 != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v1039 = v771
	goto L229
L228:
	;
	v1039 = (v1033 + int32(23)) & int32(120)
	goto L229
L229:
	;
	v1040 = v1039 + v975
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v1042 = F_ArrayGetNItemsSafe(m, v1041, v218)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L9
	} else {
		goto L230
	}
L230:
	;
	v1044 = int32(0)
	v1046 = v77 + int32(112)
	v1048 = v77 + int32(80)
	v1058 = v85 - int32(1)
	if v1058 < v1044 {
		v1136 = v1044
		goto L232
	} else {
		goto L233
	}
L231:
	;
	v1142 = F_array_seek(m, v1031, v1044, v1024, v1136, v71, v73)
	mBase = m.M
	v1143 = v1142 - v1031
	if v1143 != 0 {
		goto L241
	} else {
		goto L242
	}
L232:
	;
	goto L231
L233:
	;
	v1061 = int32(1)
	if v1058 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1070 = v1058
	v1071 = v1061
	v1072 = v1044
	v1077 = v1044
	goto L237
L235:
	;
	v1113 = v1058
	v1114 = v1061
	v1115 = v1044
	goto L236
L236:
	;
	v1122 = v1113 << (uint(int32(2)) % 32)
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v66+v1122)))
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1122+v1048)))
	v1136 = (v1124-v1126)*v1114 + v1115
	goto L232
L237:
	;
	v1078 = int32(2)
	v1079 = v1070 << (uint(v1078) % 32)
	v1081 = v1079 - int32(4)
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v66+v1081)))
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1048+v1081)))
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1079+v1046)))
	v1089 = v1088 * v1071
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1079+v66)))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1079+v1048)))
	v1098 = (v1083-v1085)*v1089 + ((v1092-v1094)*v1071 + v1072)
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1046+v1081)))
	v1101 = v1100 * v1089
	v1103 = v1070 - v1078
	v1105 = v1077 + v1078
	if v1105 != v85&int32(-2) {
		v1070 = v1103
		v1071 = v1101
		v1072 = v1098
		v1077 = v1105
		goto L237
	} else {
		goto L239
	}
L238:
	;
	if v85&int32(1) == int32(0) {
		v1136 = v1098
		goto L232
	} else {
		goto L240
	}
L239:
	;
	goto L238
L240:
	;
	v1113 = v1103
	v1114 = v1101
	v1115 = v1098
	goto L236
L241:
	;
	base.MemoryCopy(m, v1040, v1031, v1143)
	goto L243
L242:
	;
	goto L243
L243:
	;
	if v771 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1146 = v1034
	goto L246
L245:
	;
	v1146 = int32(0)
	goto L246
L246:
	;
	v1147 = int32(0)
	if base.B2i32(v771 == v1147)|base.B2i32(v1136 <= v1147) != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1413 = v77 + int32(112)
	v1415 = v77 + int32(240)
	v1419 = int32(2)
	v1420 = v85 << (uint(v1419) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v1415+v1420-int32(4)))) = int32(1)
	v1427 = v85 - v1419
	if v1427 < int32(0) {
		goto L290
	} else {
		goto L291
	}
L248:
	;
	v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034))))
	if v1024 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L249:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1349))) = uint8(v1348)
	goto L247
L250:
	;
	v1348 = v1342
	v1349 = v1237
	goto L249
L251:
	;
	v1342 = v1242 | int32(3)
	goto L250
L252:
	;
	v1342 = v1242 | int32(7)
	goto L250
L253:
	;
	v1342 = v1242 | int32(15)
	goto L250
L254:
	;
	v1342 = v1242 | int32(31)
	goto L250
L255:
	;
	v1342 = v1242 | int32(63)
	goto L250
L256:
	;
	v1342 = v1242 | int32(127)
	goto L250
L257:
	;
	v1348 = v1299 | int32(1)
	v1349 = v1300
	goto L249
L258:
	;
	v1237 = v1146
	v1239 = v1136
	v1242 = v1152
	goto L279
L259:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v1136) {
		goto L258
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1024))))
	v1158 = int32(1)
	v1160 = v1158
	v1162 = v1158
	v1164 = v1136
	v1165 = v1152
	v1166 = v1146
	v1172 = v1024
	v1178 = v1157
	goto L263
L262:
	;
	v1299 = v1152
	v1300 = v1146
	goto L257
L263:
	;
	if v1160&v1178 != 0 {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	if v1213 != int32(1) {
		v1348 = v1214
		v1349 = v1215
		goto L249
	} else {
		goto L278
	}
L265:
	;
	v1199 = v1162 | v1165
	goto L267
L266:
	;
	v1199 = v1165 & (v1162 ^ int32(-1))
	goto L267
L267:
	;
	v1200 = int32(1)
	v1201 = v1164 - v1200
	v1203 = v1162 << (uint(v1200) % 32)
	if v1203 == int32(256) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1166))) = uint8(v1199)
	if v1201 == int32(0) {
		goto L247
	} else {
		goto L271
	}
L269:
	;
	v1213 = v1203
	v1214 = v1199
	v1215 = v1166
	goto L270
L270:
	;
	v1217 = v1160 << (uint(int32(1)) % 32)
	if v1217 == int32(256) {
		goto L273
	} else {
		goto L274
	}
L271:
	;
	v1209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1166)+1)))
	v1210 = int32(1)
	v1213 = v1210
	v1214 = v1209
	v1215 = v1166 + v1210
	goto L270
L272:
	;
	goto L264
L273:
	;
	if v1201 == int32(0) {
		goto L272
	} else {
		goto L276
	}
L274:
	;
	v1226 = v1217
	v1227 = v1172
	v1228 = v1178
	goto L275
L275:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1164) {
		v1160 = v1226
		v1162 = v1213
		v1164 = v1201
		v1165 = v1214
		v1166 = v1215
		v1172 = v1227
		v1178 = v1228
		goto L263
	} else {
		goto L277
	}
L276:
	;
	v1222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1172)+1)))
	v1223 = int32(1)
	v1226 = v1223
	v1227 = v1172 + v1223
	v1228 = v1222
	goto L275
L277:
	;
	goto L272
L278:
	;
	goto L247
L279:
	;
	if v1239 < int32(3) {
		goto L251
	} else {
		goto L281
	}
L280:
	;
	v1299 = v1289
	v1300 = v1291
	goto L257
L281:
	;
	if v1239 == int32(3) {
		goto L252
	} else {
		goto L282
	}
L282:
	;
	if base.Ui32(v1239) < base.Ui32(int32(5)) {
		goto L253
	} else {
		goto L283
	}
L283:
	;
	if v1239 == int32(5) {
		goto L254
	} else {
		goto L284
	}
L284:
	;
	if base.Ui32(v1239) < base.Ui32(int32(7)) {
		goto L255
	} else {
		goto L285
	}
L285:
	;
	if v1239 == int32(7) {
		goto L256
	} else {
		goto L286
	}
L286:
	;
	v1283 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v1237))) = uint8(v1283)
	v1286 = v1239 - int32(8)
	if v1286 == int32(0) {
		goto L247
	} else {
		goto L287
	}
L287:
	;
	v1289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1237)+1)))
	v1290 = int32(1)
	v1291 = v1237 + v1290
	if v1286 != v1290 {
		v1237 = v1291
		v1239 = v1286
		v1242 = v1289
		goto L279
	} else {
		goto L288
	}
L288:
	;
	goto L280
L289:
	;
	v1485 = v77 + int32(208)
	v1486 = int32(0)
	if v85 <= v1486 {
		goto L300
	} else {
		goto L301
	}
L290:
	;
	goto L289
L291:
	;
	if v85&int32(1) == int32(0) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1438 = v1420 - int32(4)
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1413+v1438)))
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1415+v1438)))
	*(*int32)(unsafe.Add(mBase, uint32(v1415+v1427<<(uint(int32(2))%32)))) = v1440 * v1442
	v1447 = v85 - int32(3)
	goto L294
L293:
	;
	v1447 = v1427
	goto L294
L294:
	;
	if v1427 == int32(0) {
		goto L290
	} else {
		goto L295
	}
L295:
	;
	v1453 = v1447
	goto L296
L296:
	;
	v1456 = int32(2)
	v1457 = v1453 << (uint(v1456) % 32)
	v1460 = v1457 + int32(4)
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v1413+v1460)))
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1460+v1415)))
	v1465 = v1462 * v1464
	*(*int32)(unsafe.Add(mBase, uint32(v1415+v1457))) = v1465
	v1468 = v1453 - int32(1)
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v1413+v1457)))
	*(*int32)(unsafe.Add(mBase, uint32(v1415+v1468<<(uint(v1456)%32)))) = v1473 * v1465
	if v1468 != 0 {
		v1453 = v1453 - v1456
		goto L296
	} else {
		goto L298
	}
L297:
	;
	goto L290
L298:
	;
	goto L297
L299:
	;
	v1563 = v77 + int32(176)
	v1564 = int32(0)
	v1570 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1563+v85<<(uint(v1570)%32)-int32(4)))) = v1564
	v1578 = v85 - v1570
	if v1564 <= v1578 {
		goto L310
	} else {
		goto L311
	}
L300:
	;
	goto L299
L301:
	;
	if v85 != int32(1) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1502 = v1486
	v1505 = v1486
	goto L305
L303:
	;
	v1539 = v1486
	goto L304
L304:
	;
	v1544 = v1539 << (uint(int32(2)) % 32)
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1544+v64)))
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v1544+v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v1485+v1544))) = v1547 - v1549 + int32(1)
	goto L300
L305:
	;
	v1506 = int32(2)
	v1507 = v1502 << (uint(v1506) % 32)
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v1507+v64)))
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1507+v66)))
	v1514 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1485+v1507))) = v1510 - v1512 + v1514
	v1518 = v1507 | int32(4)
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v1518+v64)))
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1518+v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v1485+v1518))) = v1521 - v1523 + v1514
	v1529 = v1502 + v1506
	v1531 = v1505 + v1506
	if v1531 != v85&int32(2147483646) {
		v1502 = v1529
		v1505 = v1531
		goto L305
	} else {
		goto L307
	}
L306:
	;
	if v85&int32(1) == int32(0) {
		goto L300
	} else {
		goto L308
	}
L307:
	;
	goto L306
L308:
	;
	v1539 = v1529
	goto L304
L309:
	;
	v1684 = int32(0)
	if v222 == v1684 {
		goto L325
	} else {
		goto L326
	}
L310:
	;
	v1585 = v1578
	v1589 = v1564
	goto L313
L311:
	;
	goto L312
L312:
	;
	goto L309
L313:
	;
	v1592 = v1585 << (uint(int32(2)) % 32)
	v1593 = v1563 + v1592
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1415+v1592)))
	v1596 = int32(1)
	v1597 = v1595 - v1596
	*(*int32)(unsafe.Add(mBase, uint32(v1593))) = v1597
	v1600 = v1585 + v1596
	if v85 <= v1600 {
		goto L315
	} else {
		goto L316
	}
L314:
	;
	goto L312
L315:
	;
	v1668 = int32(1)
	if int32(0) < v1585 {
		v1585 = v1585 - v1668
		v1589 = v1589 + v1668
		goto L313
	} else {
		goto L324
	}
L316:
	;
	if v1589&int32(1) == int32(0) {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v1606 = int32(2)
	v1607 = v1600 << (uint(v1606) % 32)
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1485+v1607)))
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1415+v1607)))
	v1615 = v1597 - (v1609-int32(1))*v1613
	*(*int32)(unsafe.Add(mBase, uint32(v1593))) = v1615
	v1619 = v1585 + v1606
	v1620 = v1615
	goto L319
L318:
	;
	v1619 = v1600
	v1620 = v1597
	goto L319
L319:
	;
	if v1589 == int32(0) {
		goto L315
	} else {
		goto L320
	}
L320:
	;
	v1627 = v1619
	v1628 = v1620
	goto L321
L321:
	;
	v1633 = int32(2)
	v1634 = v1627 << (uint(v1633) % 32)
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1485+v1634)))
	v1637 = int32(1)
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1415+v1634)))
	v1642 = v1628 - (v1636-v1637)*v1640
	*(*int32)(unsafe.Add(mBase, uint32(v1593))) = v1642
	v1645 = v1634 + int32(4)
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v1485+v1645)))
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1415+v1645)))
	v1653 = v1642 - (v1647-v1637)*v1651
	*(*int32)(unsafe.Add(mBase, uint32(v1593))) = v1653
	v1656 = v1627 + v1633
	if v1656 != v85 {
		v1627 = v1656
		v1628 = v1653
		goto L321
	} else {
		goto L323
	}
L322:
	;
	goto L315
L323:
	;
	goto L322
L324:
	;
	goto L314
L325:
	;
	base.MemoryFill(m, v77+int32(144), int32(0), v220)
	goto L327
L326:
	;
	goto L327
L327:
	;
	v1697 = v85 - int32(1)
	v1698 = v1142
	v1699 = v1136
	v1701 = v83 + v1018
	v1702 = v1136
	v1707 = v1684
	v1708 = v1040 + v1143
	goto L328
L328:
	;
	v1733 = v77 + int32(176) + v1697<<(uint(int32(2))%32)
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1733)))
	if v1734 == int32(0) {
		goto L331
	} else {
		goto L332
	}
L329:
	;
	v2097 = v1042 - v2028
	v2098 = F_array_seek(m, v2036, v2028, v1024, v2097, v71, v73)
	mBase = m.M
	v2099 = v2098 - v2036
	if v2099 != 0 {
		goto L396
	} else {
		goto L397
	}
L330:
	;
	v1996 = F_array_seek(m, v1701, v1707, v1030, int32(1), v71, v73)
	mBase = m.M
	v1997 = v1996 - v1701
	if v1997 != 0 {
		goto L367
	} else {
		goto L368
	}
L331:
	;
	v1963 = v1702
	v1965 = v1699
	v1974 = v1708
	v1977 = v1698
	goto L330
L332:
	;
	goto L333
L333:
	;
	v1737 = F_array_seek(m, v1698, v1702, v1024, v1734, v71, v73)
	mBase = m.M
	v1738 = v1737 - v1698
	if v1738 != 0 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	base.MemoryCopy(m, v1708, v1698, v1738)
	goto L336
L335:
	;
	goto L336
L336:
	;
	if v771 == int32(0) {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v1733)))
	v1963 = v1958 + v1702
	v1965 = v1958 + v1699
	v1974 = v1708 + v1738
	v1977 = v1737
	goto L330
L338:
	;
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v1733)))
	if v1742 <= int32(0) {
		goto L337
	} else {
		goto L339
	}
L339:
	;
	v1748 = int32(1) << (uint(v1699&int32(7)) % 32)
	v1750 = base.I32_div_s(v1699, int32(8))
	v1751 = v1146 + v1750
	v1752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1751))))
	if v1024 != 0 {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1888))) = uint8(v1893)
	goto L337
L341:
	;
	v1758 = base.I32_div_s(v1702, int32(8))
	v1759 = v1024 + v1758
	v1760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1759))))
	v1761 = v1751
	v1763 = v1742
	v1764 = v1748
	v1766 = v1752
	v1771 = int32(1) << (uint(v1702&int32(7)) % 32)
	v1779 = v1760
	v1781 = v1759
	goto L344
L342:
	;
	goto L343
L343:
	;
	v1838 = v1751
	v1840 = v1742
	v1841 = v1748
	v1843 = v1752
	goto L360
L344:
	;
	if v1771&v1779 != 0 {
		goto L346
	} else {
		goto L347
	}
L345:
	;
	if v1815 != int32(1) {
		v1888 = v1814
		v1893 = v1816
		goto L340
	} else {
		goto L359
	}
L346:
	;
	v1800 = v1764 | v1766
	goto L348
L347:
	;
	v1800 = v1766 & (v1764 ^ int32(-1))
	goto L348
L348:
	;
	v1801 = int32(1)
	v1802 = v1763 - v1801
	v1804 = v1764 << (uint(v1801) % 32)
	if v1804 == int32(256) {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1761))) = uint8(v1800)
	if v1802 == int32(0) {
		goto L337
	} else {
		goto L352
	}
L350:
	;
	v1814 = v1761
	v1815 = v1804
	v1816 = v1800
	goto L351
L351:
	;
	v1818 = v1771 << (uint(int32(1)) % 32)
	if v1818 == int32(256) {
		goto L354
	} else {
		goto L355
	}
L352:
	;
	v1810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1761)+1)))
	v1811 = int32(1)
	v1814 = v1761 + v1811
	v1815 = v1811
	v1816 = v1810
	goto L351
L353:
	;
	goto L345
L354:
	;
	if v1802 == int32(0) {
		goto L353
	} else {
		goto L357
	}
L355:
	;
	v1827 = v1818
	v1828 = v1779
	v1829 = v1781
	goto L356
L356:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1763) {
		v1761 = v1814
		v1763 = v1802
		v1764 = v1815
		v1766 = v1816
		v1771 = v1827
		v1779 = v1828
		v1781 = v1829
		goto L344
	} else {
		goto L358
	}
L357:
	;
	v1823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1781)+1)))
	v1824 = int32(1)
	v1827 = v1824
	v1828 = v1823
	v1829 = v1781 + v1824
	goto L356
L358:
	;
	goto L353
L359:
	;
	goto L337
L360:
	;
	v1872 = v1841 | v1843
	v1873 = int32(1)
	v1874 = v1840 - v1873
	v1876 = v1841 << (uint(v1873) % 32)
	if v1876 == int32(256) {
		goto L362
	} else {
		goto L363
	}
L361:
	;
	v1888 = v1838
	v1893 = v1872
	goto L340
L362:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1838))) = uint8(v1872)
	if v1874 == int32(0) {
		goto L337
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1840) {
		v1840 = v1874
		v1841 = v1876
		v1843 = v1872
		goto L360
	} else {
		goto L366
	}
L365:
	;
	v1882 = int32(1)
	v1883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1838)+1)))
	v1838 = v1838 + v1882
	v1840 = v1874
	v1841 = v1882
	v1843 = v1883
	goto L360
L366:
	;
	goto L361
L367:
	;
	base.MemoryCopy(m, v1974, v1701, v1997)
	goto L369
L368:
	;
	goto L369
L369:
	;
	if v771 != 0 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v2002 = int32(1) << (uint(v1965&int32(7)) % 32)
	v2004 = base.I32_div_s(v1965, int32(8))
	v2005 = v1146 + v2004
	v2006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2005))))
	if v1030 != 0 {
		goto L374
	} else {
		goto L375
	}
L371:
	;
	goto L372
L372:
	;
	v2027 = int32(1)
	v2028 = v1963 + v2027
	v2032 = v1965 + v2027
	v2034 = v1997 + v1974
	v2036 = F_array_seek(m, v1977, v1963, v1024, v2027, v71, v73)
	mBase = m.M
	v2038 = v77 + int32(144)
	v2040 = v77 + int32(208)
	if v85 <= int32(0) {
		goto L381
	} else {
		goto L382
	}
L373:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2005))) = uint8(v2022)
	goto L372
L374:
	;
	v2012 = base.I32_div_s(v1707, int32(8))
	v2014 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1030+v2012))))
	if int32(base.Ui32(v2014)>>(uint(v1707&int32(7))%32))&int32(1) != 0 {
		goto L377
	} else {
		goto L378
	}
L375:
	;
	goto L376
L376:
	;
	v2022 = v2006 | v2002
	goto L373
L377:
	;
	v2020 = v2006 | v2002
	goto L379
L378:
	;
	v2020 = v2006 & (v2002 ^ int32(-1))
	goto L379
L379:
	;
	v2022 = v2020
	goto L373
L380:
	;
	if v2094 != int32(-1) {
		v1697 = v2094
		v1698 = v2036
		v1699 = v2032
		v1701 = v1997 + v1701
		v1702 = v2028
		v1707 = v1707 + v2027
		v1708 = v2034
		goto L328
	} else {
		goto L395
	}
L381:
	;
	v2094 = int32(-1)
	goto L380
L382:
	;
	goto L383
L383:
	;
	v2046 = int32(1)
	v2047 = v85 - v2046
	v2049 = v2047 << (uint(int32(2)) % 32)
	v2050 = v2038 + v2049
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v2050)))
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v2040+v2049)))
	v2056 = base.I32_rem_s(v2051+v2046, v2055)
	*(*int32)(unsafe.Add(mBase, uint32(v2050))) = v2056
	if v2047 != 0 {
		goto L385
	} else {
		goto L386
	}
L384:
	;
	v2094 = v2084
	goto L380
L385:
	;
	v2058 = v2047
	v2061 = v2056
	goto L388
L386:
	;
	goto L387
L387:
	;
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v2038)))
	if v2082 != 0 {
		goto L392
	} else {
		goto L393
	}
L388:
	;
	if v2061 != 0 {
		v2084 = v2058
		goto L384
	} else {
		goto L390
	}
L389:
	;
	goto L387
L390:
	;
	v2063 = int32(1)
	v2064 = v2058 - v2063
	v2066 = v2064 << (uint(int32(2)) % 32)
	v2067 = v2038 + v2066
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v2067)))
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v2040+v2066)))
	v2073 = base.I32_rem_s(v2068+v2063, v2072)
	*(*int32)(unsafe.Add(mBase, uint32(v2067))) = v2073
	if v2064 != 0 {
		v2058 = v2064
		v2061 = v2073
		goto L388
	} else {
		goto L391
	}
L391:
	;
	goto L389
L392:
	;
	v2083 = int32(0)
	goto L394
L393:
	;
	v2083 = int32(-1)
	goto L394
L394:
	;
	v2084 = v2083
	goto L384
L395:
	;
	goto L329
L396:
	;
	base.MemoryCopy(m, v2034, v2036, v2099)
	goto L398
L397:
	;
	goto L398
L398:
	;
	v2101 = int32(0)
	if base.B2i32(v771 == v2101)|base.B2i32(v2097 <= v2101) != 0 {
		v2956 = v975
		goto L11
	} else {
		goto L399
	}
L399:
	;
	v2109 = int32(1) << (uint(v2032&int32(7)) % 32)
	v2111 = base.I32_div_s(v2032, int32(8))
	v2112 = v1146 + v2111
	v2113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2112))))
	if v1024 != 0 {
		goto L211
	} else {
		goto L400
	}
L400:
	;
	v2114 = v2112
	v2116 = v2097
	v2117 = v2109
	v2119 = v2113
	goto L401
L401:
	;
	v2148 = v2117 | v2119
	v2149 = int32(1)
	v2150 = v2116 - v2149
	v2152 = v2117 << (uint(v2149) % 32)
	if v2152 == int32(256) {
		goto L403
	} else {
		goto L404
	}
L402:
	;
	v2902 = v2114
	v2907 = v2148
	goto L210
L403:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2114))) = uint8(v2148)
	if v2150 == int32(0) {
		v2956 = v975
		goto L11
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2116) {
		v2116 = v2150
		v2117 = v2152
		v2119 = v2148
		goto L401
	} else {
		goto L407
	}
L406:
	;
	v2158 = int32(1)
	v2159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2114)+1)))
	v2114 = v2114 + v2158
	v2116 = v2150
	v2117 = v2158
	v2119 = v2159
	goto L401
L407:
	;
	goto L402
L408:
	;
	base.MemoryCopy(m, v2165, v2164, v957)
	goto L410
L409:
	;
	goto L410
L410:
	;
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v2167 == int32(0) {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v2177 = (v2170<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L413
L412:
	;
	v2177 = v2167
	goto L413
L413:
	;
	v2178 = v957 + v2165
	if v789 != 0 {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	base.MemoryCopy(m, v2178, v2177+v83, v789)
	goto L416
L415:
	;
	goto L416
L416:
	;
	if v971 != 0 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	base.MemoryCopy(m, v2178+v789, v957+v2164+v958, v971)
	goto L419
L418:
	;
	goto L419
L419:
	;
	if v633&int32(1) == int32(0) {
		v2956 = v975
		goto L11
	} else {
		goto L420
	}
L420:
	;
	v2189 = int32(0)
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v975)+8))
	if v2191 != 0 {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v975)+4))
	v2196 = v985 + v2192<<(uint(int32(3))%32)
	goto L423
L422:
	;
	v2196 = v2189
	goto L423
L423:
	;
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v2197 != 0 {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v2198 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v2202 = v218 + v2198<<(uint(int32(3))%32)
	goto L426
L425:
	;
	v2202 = v2189
	goto L426
L426:
	;
	if v963 <= int32(0) {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v2413 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v2413 != 0 {
		goto L456
	} else {
		goto L457
	}
L428:
	;
	v2208 = int32(1) << (uint(v640&int32(7)) % 32)
	v2210 = base.I32_div_s(v640, int32(8))
	v2211 = v2196 + v2210
	v2212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2211))))
	if v2202 == int32(0) {
		goto L430
	} else {
		goto L431
	}
L429:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2350))) = uint8(v2355)
	goto L427
L430:
	;
	v2215 = v963
	v2218 = v2208
	v2221 = v2211
	v2226 = v2212
	goto L433
L431:
	;
	goto L432
L432:
	;
	v2265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2202))))
	v2267 = int32(1)
	v2270 = v2208
	v2271 = v963
	v2273 = v2211
	v2278 = v2212
	v2279 = v2202
	v2285 = v2265
	goto L440
L433:
	;
	v2249 = v2218 | v2226
	v2250 = int32(1)
	v2251 = v2215 - v2250
	v2253 = v2218 << (uint(v2250) % 32)
	if v2253 == int32(256) {
		goto L435
	} else {
		goto L436
	}
L434:
	;
	v2350 = v2221
	v2355 = v2249
	goto L429
L435:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2221))) = uint8(v2249)
	if v2251 == int32(0) {
		goto L427
	} else {
		goto L438
	}
L436:
	;
	goto L437
L437:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2215) {
		v2215 = v2251
		v2218 = v2253
		v2226 = v2249
		goto L433
	} else {
		goto L439
	}
L438:
	;
	v2259 = int32(1)
	v2260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2221)+1)))
	v2215 = v2251
	v2218 = v2259
	v2221 = v2221 + v2259
	v2226 = v2260
	goto L433
L439:
	;
	goto L434
L440:
	;
	if v2267&v2285 != 0 {
		goto L442
	} else {
		goto L443
	}
L441:
	;
	if v2320 == int32(1) {
		goto L427
	} else {
		goto L455
	}
L442:
	;
	v2306 = v2270 | v2278
	goto L444
L443:
	;
	v2306 = v2278 & (v2270 ^ int32(-1))
	goto L444
L444:
	;
	v2307 = int32(1)
	v2308 = v2271 - v2307
	v2310 = v2270 << (uint(v2307) % 32)
	if v2310 == int32(256) {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2273))) = uint8(v2306)
	if v2308 == int32(0) {
		goto L427
	} else {
		goto L448
	}
L446:
	;
	v2320 = v2310
	v2321 = v2273
	v2322 = v2306
	goto L447
L447:
	;
	v2324 = v2267 << (uint(int32(1)) % 32)
	if v2324 == int32(256) {
		goto L450
	} else {
		goto L451
	}
L448:
	;
	v2316 = int32(1)
	v2317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2273)+1)))
	v2320 = v2316
	v2321 = v2273 + v2316
	v2322 = v2317
	goto L447
L449:
	;
	goto L441
L450:
	;
	if v2308 == int32(0) {
		goto L449
	} else {
		goto L453
	}
L451:
	;
	v2333 = v2324
	v2334 = v2279
	v2335 = v2285
	goto L452
L452:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2271) {
		v2267 = v2333
		v2270 = v2320
		v2271 = v2308
		v2273 = v2321
		v2278 = v2322
		v2279 = v2334
		v2285 = v2335
		goto L440
	} else {
		goto L454
	}
L453:
	;
	v2329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2279)+1)))
	v2330 = int32(1)
	v2333 = v2330
	v2334 = v2279 + v2330
	v2335 = v2329
	goto L452
L454:
	;
	goto L449
L455:
	;
	v2350 = v2321
	v2355 = v2322
	goto L429
L456:
	;
	v2414 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v2419 = v750 + v2414<<(uint(int32(3))%32)
	goto L458
L457:
	;
	v2419 = int32(0)
	goto L458
L458:
	;
	if v746 <= int32(0) {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	if v959 <= int32(0) {
		v2956 = v975
		goto L11
	} else {
		goto L488
	}
L460:
	;
	v2423 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(v77)+80))
	v2425 = v2423 - v2424
	v2428 = int32(1) << (uint(v2425&int32(7)) % 32)
	v2430 = base.I32_div_s(v2425, int32(8))
	v2431 = v2196 + v2430
	v2432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2431))))
	if v2419 == int32(0) {
		goto L462
	} else {
		goto L463
	}
L461:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2568))) = uint8(v2574)
	goto L459
L462:
	;
	v2437 = v746
	v2438 = v2428
	v2439 = v2431
	v2445 = v2432
	goto L465
L463:
	;
	goto L464
L464:
	;
	v2485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2419))))
	v2489 = v746
	v2490 = v2428
	v2491 = v2431
	v2493 = v2419
	v2497 = v2432
	v2498 = int32(1)
	v2500 = v2485
	goto L472
L465:
	;
	v2469 = v2438 | v2445
	v2470 = int32(1)
	v2471 = v2437 - v2470
	v2473 = v2438 << (uint(v2470) % 32)
	if v2473 == int32(256) {
		goto L467
	} else {
		goto L468
	}
L467:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2439))) = uint8(v2469)
	if v2471 == int32(0) {
		goto L459
	} else {
		goto L470
	}
L468:
	;
	goto L469
L469:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2437) {
		v2437 = v2471
		v2438 = v2473
		v2445 = v2469
		goto L465
	} else {
		goto L471
	}
L470:
	;
	v2479 = int32(1)
	v2480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2439)+1)))
	v2437 = v2471
	v2438 = v2479
	v2439 = v2439 + v2479
	v2445 = v2480
	goto L465
L471:
	;
	v2568 = v2439
	v2574 = v2469
	goto L461
L472:
	;
	if v2498&v2500 != 0 {
		goto L474
	} else {
		goto L475
	}
L473:
	;
	if v2540 == int32(1) {
		goto L459
	} else {
		goto L487
	}
L474:
	;
	v2526 = v2490 | v2497
	goto L476
L475:
	;
	v2526 = v2497 & (v2490 ^ int32(-1))
	goto L476
L476:
	;
	v2527 = int32(1)
	v2528 = v2489 - v2527
	v2530 = v2490 << (uint(v2527) % 32)
	if v2530 == int32(256) {
		goto L477
	} else {
		goto L478
	}
L477:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2491))) = uint8(v2526)
	if v2528 == int32(0) {
		goto L459
	} else {
		goto L480
	}
L478:
	;
	v2540 = v2530
	v2541 = v2491
	v2542 = v2526
	goto L479
L479:
	;
	v2544 = v2498 << (uint(int32(1)) % 32)
	if v2544 == int32(256) {
		goto L482
	} else {
		goto L483
	}
L480:
	;
	v2536 = int32(1)
	v2537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2491)+1)))
	v2540 = v2536
	v2541 = v2491 + v2536
	v2542 = v2537
	goto L479
L481:
	;
	goto L473
L482:
	;
	if v2528 == int32(0) {
		goto L481
	} else {
		goto L485
	}
L483:
	;
	v2553 = v2493
	v2554 = v2544
	v2555 = v2500
	goto L484
L484:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2489) {
		v2489 = v2528
		v2490 = v2540
		v2491 = v2541
		v2493 = v2553
		v2497 = v2542
		v2498 = v2554
		v2500 = v2555
		goto L472
	} else {
		goto L486
	}
L485:
	;
	v2549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2493)+1)))
	v2550 = int32(1)
	v2553 = v2493 + v2550
	v2554 = v2550
	v2555 = v2549
	goto L484
L486:
	;
	goto L481
L487:
	;
	v2568 = v2541
	v2574 = v2542
	goto L461
L488:
	;
	v2636 = v963 + v967
	v2637 = v2636 + v640
	v2640 = int32(1) << (uint(v2637&int32(7)) % 32)
	v2642 = base.I32_div_s(v2637, int32(8))
	v2643 = v2196 + v2642
	v2644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2643))))
	if v2202 == int32(0) {
		goto L490
	} else {
		goto L491
	}
L489:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2782))) = uint8(v2784)
	v2956 = v975
	goto L11
L490:
	;
	v2647 = v2643
	v2649 = v2644
	v2650 = v2640
	v2652 = v959
	goto L493
L491:
	;
	goto L492
L492:
	;
	v2702 = base.I32_div_s(v2636, int32(8))
	v2703 = v2202 + v2702
	v2704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2703))))
	v2705 = v2643
	v2707 = v2644
	v2708 = v2640
	v2710 = v959
	v2711 = v2703
	v2715 = int32(1) << (uint(v2636&int32(7)) % 32)
	v2717 = v2704
	goto L500
L493:
	;
	v2681 = v2649 | v2650
	v2682 = int32(1)
	v2683 = v2652 - v2682
	v2685 = v2650 << (uint(v2682) % 32)
	if v2685 == int32(256) {
		goto L495
	} else {
		goto L496
	}
L495:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2647))) = uint8(v2681)
	if v2683 == int32(0) {
		v2956 = v975
		goto L11
	} else {
		goto L498
	}
L496:
	;
	goto L497
L497:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2652) {
		v2649 = v2681
		v2650 = v2685
		v2652 = v2683
		goto L493
	} else {
		goto L499
	}
L498:
	;
	v2691 = int32(1)
	v2692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2647)+1)))
	v2647 = v2647 + v2691
	v2649 = v2692
	v2650 = v2691
	v2652 = v2683
	goto L493
L499:
	;
	v2782 = v2647
	v2784 = v2681
	goto L489
L500:
	;
	if v2715&v2717 != 0 {
		goto L502
	} else {
		goto L503
	}
L501:
	;
	if v2760 == int32(1) {
		v2956 = v975
		goto L11
	} else {
		goto L515
	}
L502:
	;
	v2744 = v2707 | v2708
	goto L504
L503:
	;
	v2744 = v2707 & (v2708 ^ int32(-1))
	goto L504
L504:
	;
	v2745 = int32(1)
	v2746 = v2710 - v2745
	v2748 = v2708 << (uint(v2745) % 32)
	if v2748 == int32(256) {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2705))) = uint8(v2744)
	if v2746 == int32(0) {
		v2956 = v975
		goto L11
	} else {
		goto L508
	}
L506:
	;
	v2758 = v2705
	v2759 = v2744
	v2760 = v2748
	goto L507
L507:
	;
	v2762 = v2715 << (uint(int32(1)) % 32)
	if v2762 == int32(256) {
		goto L510
	} else {
		goto L511
	}
L508:
	;
	v2754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2705)+1)))
	v2755 = int32(1)
	v2758 = v2705 + v2755
	v2759 = v2754
	v2760 = v2755
	goto L507
L509:
	;
	goto L501
L510:
	;
	if v2746 == int32(0) {
		goto L509
	} else {
		goto L513
	}
L511:
	;
	v2771 = v2711
	v2772 = v2762
	v2773 = v2717
	goto L512
L512:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2710) {
		v2705 = v2758
		v2707 = v2759
		v2708 = v2760
		v2710 = v2746
		v2711 = v2771
		v2715 = v2772
		v2717 = v2773
		goto L500
	} else {
		goto L514
	}
L513:
	;
	v2767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2711)+1)))
	v2768 = int32(1)
	v2771 = v2711 + v2768
	v2772 = v2768
	v2773 = v2767
	goto L512
L514:
	;
	goto L509
L515:
	;
	v2782 = v2758
	v2784 = v2759
	goto L489
L516:
	;
	if v2835&v2837 != 0 {
		goto L518
	} else {
		goto L519
	}
L517:
	;
	if v2879 == int32(1) {
		v2956 = v975
		goto L11
	} else {
		goto L531
	}
L518:
	;
	v2864 = v2828 | v2830
	goto L520
L519:
	;
	v2864 = v2830 & (v2828 ^ int32(-1))
	goto L520
L520:
	;
	v2865 = int32(1)
	v2866 = v2827 - v2865
	v2868 = v2828 << (uint(v2865) % 32)
	if v2868 == int32(256) {
		goto L521
	} else {
		goto L522
	}
L521:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2825))) = uint8(v2864)
	if v2866 == int32(0) {
		v2956 = v975
		goto L11
	} else {
		goto L524
	}
L522:
	;
	v2878 = v2825
	v2879 = v2868
	v2880 = v2864
	goto L523
L523:
	;
	v2882 = v2835 << (uint(int32(1)) % 32)
	if v2882 == int32(256) {
		goto L526
	} else {
		goto L527
	}
L524:
	;
	v2874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2825)+1)))
	v2875 = int32(1)
	v2878 = v2825 + v2875
	v2879 = v2875
	v2880 = v2874
	goto L523
L525:
	;
	goto L517
L526:
	;
	if v2866 == int32(0) {
		goto L525
	} else {
		goto L529
	}
L527:
	;
	v2891 = v2831
	v2892 = v2882
	v2893 = v2837
	goto L528
L528:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2827) {
		v2825 = v2878
		v2827 = v2866
		v2828 = v2879
		v2830 = v2880
		v2831 = v2891
		v2835 = v2892
		v2837 = v2893
		goto L516
	} else {
		goto L530
	}
L529:
	;
	v2887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2831)+1)))
	v2888 = int32(1)
	v2891 = v2831 + v2888
	v2892 = v2888
	v2893 = v2887
	goto L528
L530:
	;
	goto L525
L531:
	;
	v2902 = v2878
	v2907 = v2880
	goto L210
}
func F_array_subscript_fetch_old(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v6 == int32(1) {
		v9 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+52)) = uint8(v9)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+48)) = int32(0)
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
		v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+6)))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)))
		v22 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16)+9)))
		v25 = F_array_get_element(m, v14, v15, v16+int32(12), v19, v20, v21, v22, v4+int32(52))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+48)) = v25
			return
		}
	}
}
func F_array_subscript_handler(m *base.Module, l0 int32) int32 {
	return int32(_a_F_array_subscript_handler_0)
}
func F_array_subscript_transform(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
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
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l1 == v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l3 != 0 {
		goto L51
	} else {
		goto L52
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v19 <= int32(0) {
		v137 = v6
		v139 = v6
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v137
	if v137 == int32(0) {
		goto L1
	} else {
		goto L45
	}
L6:
	;
	v28 = v6
	v30 = v6
	v31 = v6
	goto L7
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v31<<(uint(int32(2))%32))))
	if l3 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L16
	} else {
		goto L40
	}
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v85 = v30
	goto L11
L11:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	if v86 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L12:
	;
	v82 = F_lappend(m, v30, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L16
	} else {
		goto L28
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
	v39 = F_transformExpr(m, l2, v37, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
	if v71 != 0 {
		v81 = int32(0)
		goto L12
	} else {
		goto L26
	}
L16:
	;
	return
L17:
	;
	v41 = F_exprType(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v44 = int32(-1)
	v48 = F_coerce_to_target_type(m, l2, v39, v41, int32(23), v44, int32(1), int32(2), v44)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	if v48 != 0 {
		v81 = v48
		goto L12
	} else {
		goto L20
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(_a_F_array_subscript_transform_0), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v62 = F_exprLocation(m, v61)
	mBase = m.M
	F_parser_errposition(m, l2, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_array_subscript_transform_1), int32(95), int32(_a_F_array_subscript_transform_2))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v74 = int32(0)
	v76 = int32(1)
	v79 = F_makeConst(m, int32(23), int32(-1), v74, int32(4), v76, v74, v76)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	v81 = v79
	goto L12
L28:
	;
	v85 = v82
	goto L11
L29:
	;
	goto L8
L30:
	;
	v105 = F_lappend(m, v28, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L16
	} else {
		goto L38
	}
L31:
	;
	v104 = int32(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
	v91 = F_transformExpr(m, l2, v86, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L16
	} else {
		goto L34
	}
L34:
	;
	v93 = F_exprType(m, v91)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L16
	} else {
		goto L35
	}
L35:
	;
	v96 = int32(-1)
	v100 = F_coerce_to_target_type(m, l2, v91, v93, int32(23), v96, int32(1), int32(2), v96)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L16
	} else {
		goto L36
	}
L36:
	;
	if v100 == int32(0) {
		goto L29
	} else {
		goto L37
	}
L37:
	;
	v104 = v100
	goto L30
L38:
	;
	v108 = v31 + int32(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v109 <= v108 {
		v137 = v105
		v139 = v85
		goto L5
	} else {
		goto L39
	}
L39:
	;
	v28 = v105
	v30 = v85
	v31 = v108
	goto L7
L40:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L16
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(_a_F_array_subscript_transform_0), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L16
	} else {
		goto L42
	}
L42:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v123 = F_exprLocation(m, v122)
	mBase = m.M
	F_parser_errposition(m, l2, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L16
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_array_subscript_transform_1), int32(132), int32(_a_F_array_subscript_transform_2))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v145 <= int32(6) {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L16
	} else {
		goto L47
	}
L47:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L16
	} else {
		goto L48
	}
L48:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v155
	F_errmsg(m, int32(_a_F_array_subscript_transform_3), v13)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L16
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_array_subscript_transform_1), int32(152), int32(_a_F_array_subscript_transform_2))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L16
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v179 = int32(4)
	goto L53
L52:
	;
	v179 = int32(8)
	goto L53
L53:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0+v179)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v181
	m.G0 = v13 + int32(16)
	return
}
func F_array_to_json(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_makeStringInfo(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_array_to_json_internal(m, v2, v3, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
			v12 = F_cstring_to_text_with_len(m, v10, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v12
			}
		}
	}
}
func F_array_to_json_internal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(0)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v20 = v13 + int32(16)
		v21 = F_ArrayGetNItemsSafe(m, v18, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v21
			if v21 <= int32(0) {
				F_appendStringInfoString(m, l1, int32(_a_F_array_to_json_internal_0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					m.G0 = v11 + int32(32)
					return
				}
			} else {
				F_get_typlenbyvalalign(m, v15, v11+int32(14), v11+int32(13), v11+int32(12))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					F_json_categorize_type(m, v15, int32(0), v11+int32(8), v11+int32(4))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+14)))
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+13)))
						v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+12)))
						F_deconstruct_array(m, v13, v44, v45, v46, v11+int32(20), v11+int32(16), v11+int32(28))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							F_array_dim_to_json(m, l1, int32(0), v18, v20, v56, v57, v11+int32(24), v60, v61, l2)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
								F_pfree(m, v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
									F_pfree(m, v67)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										m.G0 = v11 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_array_to_sparsevec(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 float32
	_ = v123
	var v124 float32
	_ = v124
	var v127 float32
	_ = v127
	var v131 float32
	_ = v131
	var v135 float32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v173 float32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 float64
	_ = v209
	var v211 float32
	_ = v211
	var v214 int32
	_ = v214
	var v215 float64
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v242 int32
	_ = v242
	var v243 float64
	_ = v243
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v444 float32
	_ = v444
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v484 int32
	_ = v484
	var v485 float64
	_ = v485
	var v486 float32
	_ = v486
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v579 float32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L9
	} else {
		goto L159
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L9
	} else {
		goto L156
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L9
	} else {
		goto L153
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L9
	} else {
		goto L150
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L9
	} else {
		goto L147
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L9
	} else {
		goto L143
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L9
	} else {
		goto L139
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L9
	} else {
		goto L135
	}
L9:
	;
	return int32(0)
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v22 < int32(2) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v26 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L9
	} else {
		goto L131
	}
L14:
	;
	v27 = F_array_contains_nulls(m, v18)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	F_get_typlenbyvalalign(m, v29, v15+int32(30), v15+int32(29), v15+int32(28))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L19
	}
L17:
	;
	if v27 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+30)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+29)))
	v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+28)))
	F_deconstruct_array(m, v18, v39, v40, v41, v15+int32(24), int32(0), v15+int32(20))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	F_CheckDim_2(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if base.B2i32(v25 != int32(-1))&base.B2i32(v54 != v25) != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	switch v57 - int32(700) {
	case 0:
		goto L26
	case 1:
		goto L25
	default:
		goto L27
	}
L23:
	;
	F_CheckNnz(m, v334, v336)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L9
	} else {
		goto L71
	}
L24:
	;
	if v54 <= int32(0) {
		goto L58
	} else {
		goto L59
	}
L25:
	;
	if v54 <= int32(0) {
		goto L48
	} else {
		goto L49
	}
L26:
	;
	if v54 <= int32(0) {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	if v57 == int32(23) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v57 != int32(1700) {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v64 = int32(0)
	if v54 <= v64 {
		v334 = v64
		v336 = v54
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v67 = v64
	v68 = v2
	goto L31
L31:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+v68<<(uint(int32(2))%32))))
	v86 = F_DirectFunctionCall1Coll(m, int32(1319), int32(0), v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L9
	} else {
		goto L33
	}
L32:
	;
	v334 = v90
	v336 = v93
	goto L23
L33:
	;
	v90 = v67 + base.B2i32(v86 != int32(0))
	v92 = v68 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v92 < v93 {
		v67 = v90
		v68 = v92
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v334 = int32(0)
	v336 = v54
	goto L23
L36:
	;
	goto L37
L37:
	;
	v98 = int32(3)
	v99 = v54 & v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v101 = int32(0)
	if base.Ui32(v98) <= base.Ui32(v54-int32(1)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v108 = v101
	v109 = v2
	v117 = v2
	goto L41
L39:
	;
	v146 = v101
	v147 = v2
	goto L40
L40:
	;
	v158 = v146
	v159 = v147
	v168 = v2
	goto L45
L41:
	;
	v122 = v100 + v109<<(uint(int32(2))%32)
	v123 = *(*float32)(unsafe.Add(mBase, uint32(v122)))
	v124 = float32(0)
	v127 = *(*float32)(unsafe.Add(mBase, uint32(v122)+4))
	v131 = *(*float32)(unsafe.Add(mBase, uint32(v122)+8))
	v135 = *(*float32)(unsafe.Add(mBase, uint32(v122)+12))
	v138 = v108 + base.F32_ne(v123, v124) + base.F32_ne(v127, v124) + base.F32_ne(v131, v124) + base.F32_ne(v135, v124)
	v139 = int32(4)
	v140 = v109 + v139
	v142 = v117 + v139
	if v142 != v54&int32(2147483644) {
		v108 = v138
		v109 = v140
		v117 = v142
		goto L41
	} else {
		goto L43
	}
L42:
	;
	if v99 == int32(0) {
		v334 = v138
		v336 = v54
		goto L23
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v146 = v138
	v147 = v140
	goto L40
L45:
	;
	v173 = *(*float32)(unsafe.Add(mBase, uint32(v100+v159<<(uint(int32(2))%32))))
	v176 = v158 + base.F32_ne(v173, float32(0))
	v177 = int32(1)
	v180 = v168 + v177
	if v180 != v99 {
		v158 = v176
		v159 = v159 + v177
		v168 = v180
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v334 = v176
	v336 = v54
	goto L23
L47:
	;
	goto L46
L48:
	;
	v334 = int32(0)
	v336 = v54
	goto L23
L49:
	;
	goto L50
L50:
	;
	v185 = int32(0)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	if v54 != int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v193 = v185
	v194 = v2
	v197 = v2
	goto L54
L52:
	;
	v227 = v185
	v228 = v2
	goto L53
L53:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v186+v228<<(uint(int32(2))%32))))
	v243 = *(*float64)(unsafe.Add(mBase, uint32(v242)))
	v334 = v227 + base.F32_ne(base.F32_demote_f64(v243), float32(0))
	v336 = v54
	goto L23
L54:
	;
	v205 = int32(2)
	v207 = v186 + v194<<(uint(v205)%32)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	v209 = *(*float64)(unsafe.Add(mBase, uint32(v208)))
	v211 = float32(0)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v215 = *(*float64)(unsafe.Add(mBase, uint32(v214)))
	v219 = v193 + base.F32_ne(base.F32_demote_f64(v209), v211) + base.F32_ne(base.F32_demote_f64(v215), v211)
	v221 = v194 + v205
	v223 = v197 + v205
	if v223 != v54&int32(2147483646) {
		v193 = v219
		v194 = v221
		v197 = v223
		goto L54
	} else {
		goto L56
	}
L55:
	;
	if v54&int32(1) == int32(0) {
		v334 = v219
		v336 = v54
		goto L23
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v227 = v219
	v228 = v221
	goto L53
L58:
	;
	v334 = int32(0)
	v336 = v54
	goto L23
L59:
	;
	goto L60
L60:
	;
	v252 = v54 & int32(3)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v254 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v54) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v259 = v254
	v260 = v2
	v268 = v2
	goto L64
L62:
	;
	v297 = v254
	v298 = v2
	goto L63
L63:
	;
	v310 = v297
	v311 = v298
	v314 = int32(0)
	goto L68
L64:
	;
	v273 = v253 + v260<<(uint(int32(2))%32)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v275 = int32(0)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v273)+8))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	v289 = v259 + base.B2i32(v274 != v275) + base.B2i32(v278 != v275) + base.B2i32(v282 != v275) + base.B2i32(v286 != v275)
	v290 = int32(4)
	v291 = v260 + v290
	v293 = v268 + v290
	if v293 != v54&int32(2147483644) {
		v259 = v289
		v260 = v291
		v268 = v293
		goto L64
	} else {
		goto L66
	}
L65:
	;
	if v252 == int32(0) {
		v334 = v289
		v336 = v54
		goto L23
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	v297 = v289
	v298 = v291
	goto L63
L68:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v253+v311<<(uint(int32(2))%32))))
	v328 = v310 + base.B2i32(v325 != int32(0))
	v329 = int32(1)
	v332 = v314 + v329
	if v332 != v252 {
		v310 = v328
		v311 = v311 + v329
		v314 = v332
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v334 = v328
	v336 = v54
	goto L23
L70:
	;
	goto L69
L71:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v351 = F_mul_size(m, int32(4), v334)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	v353 = F_add_size(m, int32(16), v351)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	v356 = F_mul_size(m, int32(4), v334)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	v358 = F_add_size(m, v353, v356)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L9
	} else {
		goto L75
	}
L75:
	;
	v360 = F_palloc0(m, v358)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v360)+8)) = v334
	*(*int32)(unsafe.Add(mBase, uint32(v360)+4)) = v348
	v364 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v360))) = v358 << (uint(v364) % 32)
	v368 = v360 + int32(16)
	v371 = v368 + v334<<(uint(v364)%32)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	switch v372 - int32(700) {
	case 0:
		goto L80
	case 1:
		goto L79
	default:
		goto L81
	}
L77:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	F_pfree(m, v556)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L9
	} else {
		goto L117
	}
L78:
	;
	v505 = int32(0)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v506 <= v505 {
		v545 = v505
		goto L77
	} else {
		goto L109
	}
L79:
	;
	v463 = int32(0)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v464 <= v463 {
		v545 = v463
		goto L77
	} else {
		goto L101
	}
L80:
	;
	v423 = int32(0)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v424 <= v423 {
		v545 = v423
		goto L77
	} else {
		goto L93
	}
L81:
	;
	if v372 == int32(23) {
		goto L78
	} else {
		goto L82
	}
L82:
	;
	if v372 != int32(1700) {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v379 = int32(0)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v380 <= v379 {
		v545 = v379
		goto L77
	} else {
		goto L84
	}
L84:
	;
	v384 = int32(0)
	v385 = v379
	goto L85
L85:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v398+v384<<(uint(int32(2))%32))))
	v403 = F_DirectFunctionCall1Coll(m, int32(1319), int32(0), v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L9
	} else {
		goto L87
	}
L86:
	;
	v545 = v417
	goto L77
L87:
	;
	if v403&int32(2147483647) != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v360)+8))
	if v407 <= v385 {
		goto L2
	} else {
		goto L91
	}
L89:
	;
	v417 = v385
	goto L90
L90:
	;
	v420 = v384 + int32(1)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v420 < v421 {
		v384 = v420
		v385 = v417
		goto L85
	} else {
		goto L92
	}
L91:
	;
	v410 = v385 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v368+v410))) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v410+v371))) = v403
	v417 = v385 + int32(1)
	goto L90
L92:
	;
	goto L86
L93:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v429 = int32(0)
	v430 = v423
	v431 = v424
	goto L94
L94:
	;
	v444 = *(*float32)(unsafe.Add(mBase, uint32(v427+v429<<(uint(int32(2))%32))))
	if base.F32_ne(v444, float32(0)) != 0 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v545 = v458
	goto L77
L96:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v360)+8))
	if v447 <= v430 {
		goto L3
	} else {
		goto L99
	}
L97:
	;
	v458 = v430
	v459 = v431
	goto L98
L98:
	;
	v461 = v429 + int32(1)
	if v461 < v459 {
		v429 = v461
		v430 = v458
		v431 = v459
		goto L94
	} else {
		goto L100
	}
L99:
	;
	v450 = v430 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v368+v450))) = v429
	*(*float32)(unsafe.Add(mBase, uint32(v450+v371))) = v444
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v458 = v430 + int32(1)
	v459 = v455
	goto L98
L100:
	;
	goto L95
L101:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v469 = int32(0)
	v470 = v463
	v471 = v464
	goto L102
L102:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v467+v469<<(uint(int32(2))%32))))
	v485 = *(*float64)(unsafe.Add(mBase, uint32(v484)))
	v486 = base.F32_demote_f64(v485)
	if base.F32_ne(v486, float32(0)) != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v545 = v500
	goto L77
L104:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v360)+8))
	if v489 <= v470 {
		goto L4
	} else {
		goto L107
	}
L105:
	;
	v500 = v470
	v501 = v471
	goto L106
L106:
	;
	v503 = v469 + int32(1)
	if v503 < v501 {
		v469 = v503
		v470 = v500
		v471 = v501
		goto L102
	} else {
		goto L108
	}
L107:
	;
	v492 = v470 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v368+v492))) = v469
	*(*float32)(unsafe.Add(mBase, uint32(v492+v371))) = v486
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v500 = v470 + int32(1)
	v501 = v497
	goto L106
L108:
	;
	goto L103
L109:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v511 = int32(0)
	v512 = v505
	v513 = v506
	goto L110
L110:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v509+v511<<(uint(int32(2))%32))))
	if v526 != 0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v545 = v539
	goto L77
L112:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v360)+8))
	if v527 <= v512 {
		goto L5
	} else {
		goto L115
	}
L113:
	;
	v539 = v512
	v540 = v513
	goto L114
L114:
	;
	v542 = v511 + int32(1)
	if v542 < v540 {
		v511 = v542
		v512 = v539
		v513 = v540
		goto L110
	} else {
		goto L116
	}
L115:
	;
	v530 = v512 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v368+v530))) = v511
	*(*float32)(unsafe.Add(mBase, uint32(v530+v371))) = base.F32_convert_i32_s(v526)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v539 = v512 + int32(1)
	v540 = v536
	goto L114
L116:
	;
	goto L111
L117:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v360)+8))
	if v559 == v545 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v561 = int32(0)
	if v561 < v545 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L9
	} else {
		goto L128
	}
L121:
	;
	v564 = v561
	goto L124
L122:
	;
	goto L123
L123:
	;
	m.G0 = v15 + int32(32)
	return v360
L124:
	;
	v579 = *(*float32)(unsafe.Add(mBase, uint32(v371+v564<<(uint(int32(2))%32))))
	F_CheckElement_2(m, v579)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L9
	} else {
		goto L126
	}
L125:
	;
	goto L123
L126:
	;
	v583 = v564 + int32(1)
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v360)+8))
	if v583 < v584 {
		v564 = v583
		goto L124
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	F_errmsg_internal(m, int32(_a_F_array_to_sparsevec_0), int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L9
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(810), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L9
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L9
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(_a_F_array_to_sparsevec_3), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L9
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(709), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L9
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L9
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_array_to_sparsevec_4), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L9
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(714), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L9
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L9
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v25
	F_errmsg(m, int32(_a_F_array_to_sparsevec_5), v15)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L9
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(62), int32(_a_F_array_to_sparsevec_6))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L9
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L9
	} else {
		goto L144
	}
L144:
	;
	F_errmsg(m, int32(_a_F_array_to_sparsevec_7), int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L9
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(753), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L9
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	F_errmsg_internal(m, int32(_a_F_array_to_sparsevec_8), int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L9
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(776), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L9
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errmsg_internal(m, int32(_a_F_array_to_sparsevec_8), int32(0))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L9
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(781), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L9
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errmsg_internal(m, int32(_a_F_array_to_sparsevec_8), int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L9
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(786), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L9
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errmsg_internal(m, int32(_a_F_array_to_sparsevec_8), int32(0))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L9
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(791), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L9
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L9
	} else {
		goto L160
	}
L160:
	;
	F_errmsg(m, int32(_a_F_array_to_sparsevec_7), int32(0))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L9
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(797), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L9
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_upper(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_DatumGetAnyArrayP(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		if v13 == int32(-1) {
			v16 = int32(28)
		} else {
			v16 = int32(4)
		}
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v7+v16)))
		if base.Ui32(v18-int32(7)) <= base.Ui32(int32(-7)) {
			v23 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
			return int32(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v28 = int32(0)
			if base.B2i32(v28 < v27)&base.B2i32(base.Ui32(v27) <= base.Ui32(v18)) == v28 {
				v34 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
				return int32(0)
			} else {
				if v13 == int32(-1) {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
					v48 = v40
					v49 = v41
				} else {
					v43 = v7 + int32(16)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					v48 = v43
					v49 = v43 + v44<<(uint(int32(2))%32)
				}
				v53 = v27<<(uint(int32(2))%32) - int32(4)
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v48+v53)))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v49+v53)))
				return v55 + v57 - int32(1)
			}
		}
	}
}
func F_initArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	if l2 != 0 {
		v6 = int32(64)
	} else {
		v6 = int32(8)
	}
	v7 = F_initArrayResultWithSize(m, l0, l1, l2, v6)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_makeArrayResultAny(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 != 0 {
		v11 = int32(_a_F_makeArrayResultAny_0)
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_makeArrayResultAny[0]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
		*(*int32)(unsafe.Add(mBase, _c_F_makeArrayResultAny[0])) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v13
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
		v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+24)))
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+26)))
		v30 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10)+27)))
		v31 = F_construct_md_array(m, v19, v20, base.B2i32(int32(0) < v13), v8+int32(12), v8+int32(8), v27, v28, v29, v30)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_makeArrayResultAny[0])) = v12
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			F_MemoryContextDelete(m, v37)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v44 = v31
				m.G0 = v8 + int32(16)
				return v44
			}
		}
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v42 = F_makeArrayResultArr(m, v40, l1, int32(1))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			v44 = v42
			m.G0 = v8 + int32(16)
			return v44
		}
	}
}
func F_transformArrayExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v20 = F_palloc0(m, int32(36))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+20)) = uint8(v24)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(35)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L85
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v247
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v262
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v264
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v266
	m.G0 = v17 + int32(48)
	return v20
L5:
	;
	v247 = l2
	v248 = l3
	v257 = v6
	goto L4
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if int32(0) < v29 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if l2 == int32(0) {
		goto L3
	} else {
		goto L84
	}
L9:
	;
	v42 = v6
	v43 = v6
	goto L12
L10:
	;
	v91 = v6
	goto L11
L11:
	;
	if l2 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v43<<(uint(int32(2))%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v51 == int32(80) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v91 = v75
	goto L11
L14:
	;
	v75 = F_lappend(m, v42, v73)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L26
	}
L15:
	;
	v71 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+20)) = uint8(v71)
	v73 = v69
	goto L14
L16:
	;
	v54 = F_transformArrayExpr(m, l0, v50, l2, l3, l4)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v56 = F_transformExprRecurse(m, l0, v50)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v69 = v54
	goto L15
L20:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+20)))
	if v58 != 0 {
		v73 = v56
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v59 = F_exprType(m, v56)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v59&int32(-9) == int32(22) {
		v73 = v56
		goto L14
	} else {
		goto L23
	}
L23:
	;
	v65 = F_get_element_type(m, v59)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v65 == int32(0) {
		v73 = v56
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v69 = v56
	goto L15
L26:
	;
	v78 = v43 + int32(1)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v78 < v79 {
		v42 = v75
		v43 = v78
		goto L12
	} else {
		goto L27
	}
L27:
	;
	goto L13
L28:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v161 <= int32(0) {
		goto L61
	} else {
		goto L62
	}
L29:
	;
	if v91 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v91 == int32(0) {
		goto L5
	} else {
		goto L57
	}
L32:
	;
	v101 = F_select_common_type(m, l0, v91, int32(_a_F_transformArrayExpr_0), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+20)))
	if v103 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v106 = F_get_element_type(m, v101)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v129 = F_get_array_type(m, v101)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L47
	}
L37:
	;
	if v106 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v158 = v106
	v159 = v101
	v160 = v101
	goto L28
L39:
	;
	goto L40
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v115 = F_format_type_be(m, v101)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v115
	F_errmsg(m, int32(_a_F_transformArrayExpr_1), v17)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_parser_errposition(m, l0, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_transformArrayExpr_2), int32(2121), int32(_a_F_transformArrayExpr_3))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	if v129 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v158 = v101
	v159 = v101
	v160 = v129
	goto L28
L49:
	;
	goto L50
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v138 = F_format_type_be(m, v101)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v138
	F_errmsg(m, int32(_a_F_transformArrayExpr_4), v17+int32(32))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_parser_errposition(m, l0, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_transformArrayExpr_2), int32(2132), int32(_a_F_transformArrayExpr_3))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+20)))
	if v156 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v157 = l2
	goto L60
L59:
	;
	v157 = l3
	goto L60
L60:
	;
	v158 = l3
	v159 = v157
	v160 = l2
	goto L28
L61:
	;
	v247 = v160
	v248 = v158
	v257 = v6
	goto L4
L62:
	;
	goto L63
L63:
	;
	v176 = int32(0)
	v177 = v6
	goto L64
L64:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179+v176<<(uint(int32(2))%32))))
	if l2 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v247 = v160
	v248 = v158
	v257 = v223
	goto L4
L66:
	;
	v223 = F_lappend(m, v177, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L82
	}
L67:
	;
	v184 = F_exprType(m, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v220 = F_coerce_to_common_type(m, l0, v183, v159, int32(_a_F_transformArrayExpr_0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L81
	}
L70:
	;
	v189 = F_coerce_to_target_type(m, l0, v183, v184, v159, l4, int32(3), int32(1), int32(-1))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v189 != 0 {
		v222 = v189
		goto L66
	} else {
		goto L72
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errcode(m, int32(101744772))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v198 = F_exprType(m, v183)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v200 = F_format_type_be(m, v198)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v202 = F_format_type_be(m, v159)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v200
	F_errmsg(m, int32(_a_F_transformArrayExpr_5), v17+int32(16))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v211 = F_exprLocation(m, v183)
	mBase = m.M
	F_parser_errposition(m, l0, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_transformArrayExpr_2), int32(2167), int32(_a_F_transformArrayExpr_3))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	v222 = v220
	goto L66
L82:
	;
	v226 = v176 + int32(1)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v226 < v227 {
		v176 = v226
		v177 = v223
		goto L64
	} else {
		goto L83
	}
L83:
	;
	goto L65
L84:
	;
	goto L5
L85:
	;
	F_errcode(m, int32(134611076))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(_a_F_transformArrayExpr_6), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errhint(m, int32(_a_F_transformArrayExpr_7), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_parser_errposition(m, l0, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_transformArrayExpr_2), int32(2107), int32(_a_F_transformArrayExpr_3))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_trim_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if int32(0) < v15 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v19 = v18
		} else {
			v19 = int32(0)
		}
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = int32(0)
		if base.B2i32(v20 < v21)|base.B2i32(v19 < v20) == v21 {
			v27 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v8)+28)) = uint16(v27)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v27
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v27
			*(*uint16)(unsafe.Add(mBase, uint32(v8)+20)) = uint16(v27)
			if v27 < v15 {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v11+v15<<(uint(int32(2))%32))+16))
				v41 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)) = uint8(v41)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v40 + (v19 + (v20 ^ int32(-1)))
			} else {
			}
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
			F_get_typlenbyvalalign(m, v49, v8+int32(94), v8+int32(93), v8+int32(92))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8)+94)))
				v69 = int32(*(*int8)(unsafe.Add(mBase, uint32(v8)+92)))
				v70 = F_array_get_slice(m, v11, int32(1), v8+int32(32), v8-int32(-64), v8+int32(16), v8+int32(24), int32(-1), v68, v69)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(96)
					return v70
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v19
					F_errmsg(m, int32(_a_F_trim_array_0), v8)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_trim_array_1), int32(_a_F_trim_array_2), int32(_a_F_trim_array_3))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
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
