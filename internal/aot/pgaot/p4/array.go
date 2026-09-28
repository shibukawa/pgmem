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
	if base.Ui64(v37) < base.Ui64(int64(134217728)) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(134217727)
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
func F_array_agg_array_finalfn(m *base.Module, l0 int32) int64 {
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
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v3 == int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v6 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_array_finalfn[0]))
			v15 = F_makeArrayResultArr(m, v6, v13, int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				return v15
			}
		} else {
			v8 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v8)
			return int64(0)
		}
	} else {
		v8 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v8)
		return int64(0)
	}
}
func F_array_agg_finalfn(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v9 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v12 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v18 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v17
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_array_agg_finalfn[0]))
			v29 = F_makeMdArrayResult(m, v12, v18, v7+int32(12), v7+int32(8), v27, int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int64(0)
			} else {
				v35 = v29
				m.G0 = v7 + int32(16)
				return v35
			}
		} else {
			v14 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
			v35 = int64(0)
			m.G0 = v7 + int32(16)
			return v35
		}
	} else {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		v35 = int64(0)
		m.G0 = v7 + int32(16)
		return v35
	}
}
func F_array_agg_serialize(m *base.Module, l0 int32) int64 {
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
	var v183 int64
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	F_appendBinaryStringInfo(m, v16, v132, v133<<(uint(int32(3))%32))
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
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v179+v169<<(uint(int32(3))%32))))
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
	return base.I64_extend_i32_u(v241)
}
func F_array_agg_transfn(m *base.Module, l0 int32) int64 {
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
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = F_get_fn_expr_argtype(m, v9, int32(1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
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
				switch v21 - int32(435) {
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
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+376))
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
				v89 = m.ExcPending
				if v89 != 0 {
					return int64(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_array_agg_transfn_0), int32(0))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_array_agg_transfn_1), int32(577), int32(_a_F_array_agg_transfn_2))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v49 == int32(1) {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v54 = F_initArrayResult(m, v11, v52, int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int64(0)
					} else {
						v57 = v54
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
						if v58 != 0 {
							v61 = int64(0)
						} else {
							v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
							v61 = v60
						}
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						v63 = F_accumArrayResult(m, v57, v61, v58, v11, v62)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int64(0)
						} else {
							m.G0 = v7 + int32(16)
							return base.I64_extend_i32_u(v63)
						}
					}
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v57 = v56
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
					if v58 != 0 {
						v61 = int64(0)
					} else {
						v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
						v61 = v60
					}
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v63 = F_accumArrayResult(m, v57, v61, v58, v11, v62)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int64(0)
					} else {
						m.G0 = v7 + int32(16)
						return base.I64_extend_i32_u(v63)
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_array_agg_transfn_3), int32(0))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_array_agg_transfn_1), int32(566), int32(_a_F_array_agg_transfn_2))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int64(0)
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
func F_array_cat(m *base.Module, l0 int32) int64 {
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
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
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
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
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
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
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
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
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
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v839 int32
	_ = v839
	v25 = m.G0
	v27 = v25 - int32(32)
	m.G0 = v27
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v30 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v27 + int32(32)
	return base.I64_extend_i32_u(v839)
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
	v839 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return int64(0)
L9:
	;
	v839 = v39
	goto L1
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v839 = v46
	goto L1
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v464 = F_ArrayGetNItemsSafe(m, v463, v443)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L8
	} else {
		goto L115
	}
L17:
	;
	v443 = v419
	v446 = v422
	v463 = v57
	goto L16
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L8
	} else {
		goto L110
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
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
	v349 = m.ExcPending
	if v349 != 0 {
		goto L8
	} else {
		goto L98
	}
L23:
	;
	v839 = v52
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
	v839 = v49
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
	v119 = F_palloc_mul(m, int32(4), v57)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L8
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v57 == v68 {
		goto L49
	} else {
		goto L50
	}
L41:
	;
	v122 = F_palloc_mul(m, int32(4), v57)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v124 + v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v128
	if v57 < int32(2) {
		v419 = v119
		v422 = v122
		goto L17
	} else {
		goto L43
	}
L43:
	;
	v134 = int32(1)
	goto L44
L44:
	;
	v158 = v134 << (uint(int32(2)) % 32)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v80+v158)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158+v84)))
	if v160 != v162 {
		goto L18
	} else {
		goto L46
	}
L45:
	;
	v419 = v119
	v422 = v122
	goto L17
L46:
	;
	v164 = v158 + v116
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v158+v113)))
	if v165 != v167 {
		goto L18
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158+v119))) = v160
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	*(*int32)(unsafe.Add(mBase, uint32(v158+v122))) = v172
	v175 = v134 + int32(1)
	if v175 != v57 {
		v134 = v175
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v179 = F_palloc_mul(m, int32(4), v58)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L8
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v263 = F_palloc_mul(m, int32(4), v57)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L8
	} else {
		goto L75
	}
L52:
	;
	v182 = F_palloc_mul(m, int32(4), v58)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v184 = int32(0)
	v185 = base.B2i32(v112 == v184)
	if v185 == v184 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	base.MemoryCopy(m, v179, v84, v112)
	goto L56
L55:
	;
	goto L56
L56:
	;
	if v185 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	base.MemoryCopy(m, v182, v113, v112)
	goto L59
L58:
	;
	goto L59
L59:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	*(*int32)(unsafe.Add(mBase, uint32(v179))) = v192 + int32(1)
	v196 = int32(0)
	if v196 < v57 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v200 = v57
	goto L62
L61:
	;
	v200 = v196
	goto L62
L62:
	;
	v201 = v196
	goto L63
L63:
	;
	if v201 == v200 {
		v443 = v179
		v446 = v182
		v463 = v58
		goto L16
	} else {
		goto L65
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L8
	} else {
		goto L70
	}
L65:
	;
	v226 = int32(2)
	v227 = v201 << (uint(v226) % 32)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v80+v227)))
	v231 = v201 + int32(1)
	v233 = v231 << (uint(v226) % 32)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v179+v233)))
	if v229 == v235 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v116+v227)))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v182+v233)))
	if v238 == v240 {
		v201 = v231
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
	v248 = m.ExcPending
	if v248 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(_a_F_array_cat_0), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	v255 = F_errdetail(m, int32(_a_F_array_cat_1), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_array_cat_2), int32(479), int32(_a_F_array_cat_3))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
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
	v266 = F_palloc_mul(m, int32(4), v57)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	v268 = int32(0)
	v269 = base.B2i32(v115 == v268)
	if v269 == v268 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	base.MemoryCopy(m, v263, v80, v115)
	goto L79
L78:
	;
	goto L79
L79:
	;
	if v269 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	base.MemoryCopy(m, v266, v116, v115)
	goto L82
L81:
	;
	goto L82
L82:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = v276 + int32(1)
	v280 = int32(0)
	if v280 < v58 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v284 = v58
	goto L85
L84:
	;
	v284 = v280
	goto L85
L85:
	;
	v285 = v280
	goto L86
L86:
	;
	if v285 == v284 {
		v419 = v263
		v422 = v266
		goto L17
	} else {
		goto L88
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L8
	} else {
		goto L93
	}
L88:
	;
	v310 = int32(2)
	v311 = v285 << (uint(v310) % 32)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v84+v311)))
	v315 = v285 + int32(1)
	v317 = v315 << (uint(v310) % 32)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v263+v317)))
	if v313 == v319 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v311+v113)))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v266+v317)))
	if v322 == v324 {
		v285 = v315
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
	v332 = m.ExcPending
	if v332 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(_a_F_array_cat_0), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L8
	} else {
		goto L95
	}
L95:
	;
	v339 = F_errdetail(m, int32(_a_F_array_cat_1), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_array_cat_2), int32(507), int32(_a_F_array_cat_3))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
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
	v352 = m.ExcPending
	if v352 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	F_errmsg(m, int32(_a_F_array_cat_0), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L8
	} else {
		goto L100
	}
L100:
	;
	v357 = F_format_type_be(m, v54)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	v359 = F_format_type_be(m, v55)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v357
	v366 = F_errdetail(m, int32(_a_F_array_cat_4), v27+int32(16))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_array_cat_2), int32(376), int32(_a_F_array_cat_3))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
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
	v379 = m.ExcPending
	if v379 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(_a_F_array_cat_0), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v57
	v387 = F_errdetail(m, int32(_a_F_array_cat_5), v27)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_array_cat_2), int32(414), int32(_a_F_array_cat_3))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
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
	v401 = m.ExcPending
	if v401 != 0 {
		goto L8
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_array_cat_0), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L8
	} else {
		goto L112
	}
L112:
	;
	v408 = F_errdetail(m, int32(_a_F_array_cat_6), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L8
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_array_cat_2), int32(450), int32(_a_F_array_cat_3))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
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
	F_ArrayCheckBounds(m, v463, v443, v446)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L8
	} else {
		goto L116
	}
L116:
	;
	v468 = int32(2)
	v470 = int32(base.Ui32(v110)>>(uint(v468)%32)) - v108
	v473 = int32(base.Ui32(v109)>>(uint(v468)%32)) - v97
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	if v475 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	v501 = v499 + (v470 + v473)
	v502 = F_palloc0(m, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L8
	} else {
		goto L123
	}
L118:
	;
	v499 = (v463<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v500 = int32(0)
	goto L117
L119:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v478 == int32(0) {
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v484 = base.I32_div_s(v464+int32(7), int32(8))
	v491 = (v484 + v463<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v499 = v491
	v500 = v491
	goto L117
L122:
	;
	goto L121
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v502)+12)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v502)+8)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v502)+4)) = v463
	v507 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v502))) = v501 << (uint(v507) % 32)
	v511 = v502 + int32(16)
	v513 = v463 << (uint(v507) % 32)
	v514 = int32(0)
	v515 = base.B2i32(v513 == v514)
	if v515 == v514 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	base.MemoryCopy(m, v511, v443, v513)
	goto L126
L125:
	;
	goto L126
L126:
	;
	if v515 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	base.MemoryCopy(m, v513+v511, v446, v513)
	goto L129
L128:
	;
	goto L129
L129:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v502)+8))
	if v523 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v502)+4))
	v533 = (v526<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L132
L131:
	;
	v533 = v523
	goto L132
L132:
	;
	v535 = v57 << (uint(int32(3)) % 32)
	if v473 != 0 {
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
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v502)+8))
	if v544 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	v541 = v78
	goto L138
L137:
	;
	v541 = (v535 + int32(23)) & int32(-8)
	goto L138
L138:
	;
	base.MemoryCopy(m, v502+v533, v49+v541, v473)
	goto L135
L139:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v502)+4))
	v554 = (v547<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L141
L140:
	;
	v554 = v544
	goto L141
L141:
	;
	v556 = v58 << (uint(int32(3)) % 32)
	if v470 != 0 {
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
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v502)+8))
	if v566 == int32(0) {
		v839 = v502
		goto L1
	} else {
		goto L148
	}
L145:
	;
	v563 = v77
	goto L147
L146:
	;
	v563 = (v556 + int32(23)) & int32(-8)
	goto L147
L147:
	;
	base.MemoryCopy(m, v502+v554+v473, v52+v563, v470)
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
	v571 = v556 + v84
	goto L151
L150:
	;
	v571 = int32(0)
	goto L151
L151:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v502)+4))
	v576 = int32(0)
	if v78 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v579 = v535 + v80
	goto L154
L153:
	;
	v579 = v576
	goto L154
L154:
	;
	v580 = int32(0)
	if v81 <= v580 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v502)+8))
	if v706 != 0 {
		goto L186
	} else {
		goto L187
	}
L156:
	;
	goto L155
L157:
	;
	v587 = int32(1)
	v592 = base.I32_div_s(v576, int32(8))
	v593 = v511 + v572<<(uint(int32(3))%32) + v592
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593))))
	if v579 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v687))) = uint8(v688)
	goto L156
L159:
	;
	v597 = v593
	v598 = v594
	v601 = v81
	v602 = v587
	goto L162
L160:
	;
	goto L161
L161:
	;
	v632 = base.I32_div_s(v580, int32(8))
	v633 = v579 + v632
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633))))
	v635 = v593
	v636 = v594
	v638 = v633
	v639 = v81
	v640 = v587
	v641 = int32(1)
	v642 = v634
	goto L170
L162:
	;
	v606 = v598 | v602
	v607 = int32(1)
	v608 = v601 - v607
	v610 = v602 << (uint(v607) % 32)
	if v610 == int32(256) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if v622 != int32(1) {
		v687 = v620
		v688 = v621
		goto L158
	} else {
		goto L169
	}
L164:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v597))) = uint8(v606)
	if v608 == int32(0) {
		goto L156
	} else {
		goto L167
	}
L165:
	;
	v620 = v597
	v621 = v606
	v622 = v610
	goto L166
L166:
	;
	if base.Ui32(int32(1)) < base.Ui32(v601) {
		v597 = v620
		v598 = v621
		v601 = v608
		v602 = v622
		goto L162
	} else {
		goto L168
	}
L167:
	;
	v616 = int32(1)
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597)+1)))
	v620 = v597 + v616
	v621 = v617
	v622 = v616
	goto L166
L168:
	;
	goto L163
L169:
	;
	goto L156
L170:
	;
	if v641&v642 != 0 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if v665 == int32(1) {
		goto L156
	} else {
		goto L185
	}
L172:
	;
	v649 = v636 | v640
	goto L174
L173:
	;
	v649 = v636 & (v640 ^ int32(-1))
	goto L174
L174:
	;
	v650 = int32(1)
	v651 = v639 - v650
	v653 = v640 << (uint(v650) % 32)
	if v653 == int32(256) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v635))) = uint8(v649)
	if v651 == int32(0) {
		goto L156
	} else {
		goto L178
	}
L176:
	;
	v663 = v635
	v664 = v649
	v665 = v653
	goto L177
L177:
	;
	v667 = v641 << (uint(int32(1)) % 32)
	if v667 == int32(256) {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	v659 = int32(1)
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v635)+1)))
	v663 = v635 + v659
	v664 = v660
	v665 = v659
	goto L177
L179:
	;
	goto L171
L180:
	;
	if v651 == int32(0) {
		goto L179
	} else {
		goto L183
	}
L181:
	;
	v676 = v638
	v677 = v667
	v678 = v642
	goto L182
L182:
	;
	if base.Ui32(int32(1)) < base.Ui32(v639) {
		v635 = v663
		v636 = v664
		v638 = v676
		v639 = v651
		v640 = v665
		v641 = v677
		v642 = v678
		goto L170
	} else {
		goto L184
	}
L183:
	;
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638)+1)))
	v673 = int32(1)
	v676 = v638 + v673
	v677 = v673
	v678 = v672
	goto L182
L184:
	;
	goto L179
L185:
	;
	v687 = v663
	v688 = v664
	goto L158
L186:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v502)+4))
	v712 = v511 + v707<<(uint(int32(3))%32)
	goto L188
L187:
	;
	v712 = int32(0)
	goto L188
L188:
	;
	v713 = int32(0)
	if v85 <= v713 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v839 = v502
	goto L1
L190:
	;
	goto L189
L191:
	;
	v723 = int32(1) << (uint(v81&int32(7)) % 32)
	v725 = base.I32_div_s(v81, int32(8))
	v726 = v712 + v725
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726))))
	if v571 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v820))) = uint8(v821)
	goto L190
L193:
	;
	v730 = v726
	v731 = v727
	v734 = v85
	v735 = v723
	goto L196
L194:
	;
	goto L195
L195:
	;
	v765 = base.I32_div_s(v713, int32(8))
	v766 = v571 + v765
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766))))
	v768 = v726
	v769 = v727
	v771 = v766
	v772 = v85
	v773 = v723
	v774 = int32(1)
	v775 = v767
	goto L204
L196:
	;
	v739 = v731 | v735
	v740 = int32(1)
	v741 = v734 - v740
	v743 = v735 << (uint(v740) % 32)
	if v743 == int32(256) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	if v755 != int32(1) {
		v820 = v753
		v821 = v754
		goto L192
	} else {
		goto L203
	}
L198:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v730))) = uint8(v739)
	if v741 == int32(0) {
		goto L190
	} else {
		goto L201
	}
L199:
	;
	v753 = v730
	v754 = v739
	v755 = v743
	goto L200
L200:
	;
	if base.Ui32(int32(1)) < base.Ui32(v734) {
		v730 = v753
		v731 = v754
		v734 = v741
		v735 = v755
		goto L196
	} else {
		goto L202
	}
L201:
	;
	v749 = int32(1)
	v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730)+1)))
	v753 = v730 + v749
	v754 = v750
	v755 = v749
	goto L200
L202:
	;
	goto L197
L203:
	;
	goto L190
L204:
	;
	if v774&v775 != 0 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	if v798 == int32(1) {
		goto L190
	} else {
		goto L219
	}
L206:
	;
	v782 = v769 | v773
	goto L208
L207:
	;
	v782 = v769 & (v773 ^ int32(-1))
	goto L208
L208:
	;
	v783 = int32(1)
	v784 = v772 - v783
	v786 = v773 << (uint(v783) % 32)
	if v786 == int32(256) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v768))) = uint8(v782)
	if v784 == int32(0) {
		goto L190
	} else {
		goto L212
	}
L210:
	;
	v796 = v768
	v797 = v782
	v798 = v786
	goto L211
L211:
	;
	v800 = v774 << (uint(int32(1)) % 32)
	if v800 == int32(256) {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	v792 = int32(1)
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v768)+1)))
	v796 = v768 + v792
	v797 = v793
	v798 = v792
	goto L211
L213:
	;
	goto L205
L214:
	;
	if v784 == int32(0) {
		goto L213
	} else {
		goto L217
	}
L215:
	;
	v809 = v771
	v810 = v800
	v811 = v775
	goto L216
L216:
	;
	if base.Ui32(int32(1)) < base.Ui32(v772) {
		v768 = v796
		v769 = v797
		v771 = v809
		v772 = v784
		v773 = v798
		v774 = v810
		v775 = v811
		goto L204
	} else {
		goto L218
	}
L217:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v771)+1)))
	v806 = int32(1)
	v809 = v771 + v806
	v810 = v806
	v811 = v805
	goto L216
L218:
	;
	goto L213
L219:
	;
	v820 = v796
	v821 = v797
	goto L192
}
func F_array_eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v259 int32
	_ = v259
	var v276 int64
	_ = v276
	var v277 int32
	_ = v277
	var v282 int64
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int64
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v328 int64
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	v13 = int64(0)
	v16 = m.G0
	v18 = v16 - int32(128)
	m.G0 = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = F_DatumGetAnyArrayP(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v26 = F_DatumGetAnyArrayP(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v29 == int32(-1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v28 == int32(-1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	v35 = v32
	goto L4
L6:
	;
	goto L7
L7:
	;
	v35 = v21 + int32(16)
	goto L4
L8:
	;
	if v29 == int32(-1) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v41 = v38
	goto L8
L10:
	;
	goto L11
L11:
	;
	v41 = v26 + int32(16)
	goto L8
L12:
	;
	if v29 == int32(-1) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v51 = v44
	goto L12
L14:
	;
	goto L15
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v51 = v21 + v45<<(uint(int32(2))%32) + int32(16)
	goto L12
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L108
	}
L17:
	;
	v56 = int32(40)
	goto L19
L18:
	;
	v56 = int32(12)
	goto L19
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v21+v56)))
	if v28 == int32(-1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71+v26)))
	if v58 == v73 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v70 = v61
	v71 = int32(40)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v70 = v26 + v63<<(uint(int32(2))%32) + int32(16)
	v71 = int32(12)
	goto L20
L24:
	;
	if v29 == int32(-1) {
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
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L104
	}
L27:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v330 == int32(-1) {
		goto L96
	} else {
		goto L97
	}
L28:
	;
	v79 = int32(28)
	goto L30
L29:
	;
	v79 = int32(4)
	goto L30
L30:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v21+v79)))
	if v28 == int32(-1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v86 = int32(28)
	goto L33
L32:
	;
	v86 = int32(4)
	goto L33
L33:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v26+v86)))
	if v81 != v88 {
		v328 = v13
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v92 = v81 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v92) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	if v154 != 0 {
		v328 = v13
		goto L27
	} else {
		goto L53
	}
L36:
	;
	v154 = int32(0)
	goto L35
L37:
	;
	v128 = v123
	v129 = v124
	v130 = v125
	goto L47
L38:
	;
	if (v35|v41)&int32(3) != 0 {
		v123 = v35
		v124 = v41
		v125 = v92
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v116 = v35
	v117 = v41
	v118 = v92
	goto L40
L40:
	;
	if v118 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v100 = v35
	v101 = v41
	v102 = v92
	goto L42
L42:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if v105 != v106 {
		v123 = v100
		v124 = v101
		v125 = v102
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v116 = v111
	v117 = v109
	v118 = v113
	goto L40
L44:
	;
	v108 = int32(4)
	v109 = v101 + v108
	v111 = v100 + v108
	v113 = v102 - v108
	if base.Ui32(int32(3)) < base.Ui32(v113) {
		v100 = v111
		v101 = v109
		v102 = v113
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v123 = v116
	v124 = v117
	v125 = v118
	goto L37
L47:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	if v133 == v134 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v154 = v133 - v134
	goto L35
L49:
	;
	v136 = int32(1)
	v141 = v130 - v136
	if v141 != 0 {
		v128 = v128 + v136
		v129 = v129 + v136
		v130 = v141
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
	if base.Ui32(int32(4)) <= base.Ui32(v92) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	if v216 != 0 {
		v328 = v13
		goto L27
	} else {
		goto L72
	}
L55:
	;
	v216 = int32(0)
	goto L54
L56:
	;
	v190 = v185
	v191 = v186
	v192 = v187
	goto L66
L57:
	;
	if (v51|v70)&int32(3) != 0 {
		v185 = v51
		v186 = v70
		v187 = v92
		goto L56
	} else {
		goto L60
	}
L58:
	;
	v178 = v51
	v179 = v70
	v180 = v92
	goto L59
L59:
	;
	if v180 == int32(0) {
		goto L55
	} else {
		goto L65
	}
L60:
	;
	v162 = v51
	v163 = v70
	v164 = v92
	goto L61
L61:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	if v167 != v168 {
		v185 = v162
		v186 = v163
		v187 = v164
		goto L56
	} else {
		goto L63
	}
L62:
	;
	v178 = v173
	v179 = v171
	v180 = v175
	goto L59
L63:
	;
	v170 = int32(4)
	v171 = v163 + v170
	v173 = v162 + v170
	v175 = v164 - v170
	if base.Ui32(int32(3)) < base.Ui32(v175) {
		v162 = v173
		v163 = v171
		v164 = v175
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v185 = v178
	v186 = v179
	v187 = v180
	goto L56
L66:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v195 == v196 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v216 = v195 - v196
	goto L54
L68:
	;
	v198 = int32(1)
	v203 = v192 - v198
	if v203 != 0 {
		v190 = v190 + v198
		v191 = v191 + v198
		v192 = v203
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
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+16))
	if v218 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v230 = int32(*(*int8)(unsafe.Add(mBase, uint32(v229)+11)))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+10)))
	v232 = int32(*(*int16)(unsafe.Add(mBase, uint32(v229)+8)))
	v233 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+90)) = uint16(v233)
	v235 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+88)) = uint8(v235)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v90
	*(*int64)(unsafe.Add(mBase, uint32(v18)+76)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v229 + int32(76)
	v243 = F_ArrayGetNItemsSafe(m, v81, v35)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L80
	}
L74:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	if v219 == v58 {
		v229 = v218
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v222 = F_lookup_type_cache(m, v58, int32(32))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v222)+80))
	if v224 == int32(0) {
		goto L16
	} else {
		goto L79
	}
L79:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v227)+16)) = v222
	v229 = v222
	goto L73
L80:
	;
	F_array_iter_setup(m, v18+int32(44), v21, v232, v231, v230)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_array_iter_setup(m, v18+int32(16), v26, v232, v231, v230)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v253 = int64(1)
	if v243 <= int32(0) {
		v328 = v253
		goto L27
	} else {
		goto L83
	}
L83:
	;
	v259 = int32(0)
	goto L84
L84:
	;
	v276 = F_array_iter_next(m, v18+int32(44), v18+int32(15), v259)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L86
	}
L85:
	;
	v328 = v253
	goto L27
L86:
	;
	v282 = F_array_iter_next(m, v18+int32(16), v18+int32(14), v259)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+14)))
	v285 = int32(1)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)))
	if v284&v285&base.B2i32(v287 == v285) != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v313 = v259 + int32(1)
	if v313 != v243 {
		v259 = v313
		goto L84
	} else {
		goto L95
	}
L89:
	;
	if v287|v284&int32(1) != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v328 = int64(0)
	goto L27
L91:
	;
	v294 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+120)) = uint8(v294)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v282
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+104)) = uint8(v294)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v276
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+88)) = uint8(v294)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	v306 = m.T0[v305].(func(*base.Module, int32) int64)(m, v18+int32(72))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+88)))
	if v308 != 0 {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	if v306 != int64(0) {
		goto L88
	} else {
		goto L94
	}
L94:
	;
	goto L90
L95:
	;
	goto L85
L96:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v337 == int32(-1) {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v21 == v333 {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	F_pfree(m, v21)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	goto L96
L100:
	;
	m.G0 = v18 + int32(128)
	return v328
L101:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v26 == v340 {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	F_pfree(m, v26)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	goto L100
L104:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errmsg(m, int32(_a_F_array_eq_0), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_array_eq_1), int32(3854), int32(_a_F_array_eq_2))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v371 = F_format_type_be(m, v58)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v371
	F_errmsg(m, int32(_a_F_array_eq_3), v18)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_array_eq_1), int32(3879), int32(_a_F_array_eq_2))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_ge(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(int32(0) <= v2))
	}
}
func F_array_lt(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(int32(base.Ui32(v2) >> (uint(int32(31)) % 32)))
	}
}
func F_array_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v202 int32
	_ = v202
	var v207 int64
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v302 int32
	_ = v302
	var v314 int32
	_ = v314
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v527 int32
	_ = v527
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v724 int32
	_ = v724
	var v733 int32
	_ = v733
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v935 int32
	_ = v935
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1035 int32
	_ = v1035
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1172 int32
	_ = v1172
	v22 = m.G0
	v24 = v22 - int32(288)
	m.G0 = v24
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v27 = F_DatumGetAnyArrayP(m, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v33 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v36 = int32(40)
	goto L5
L4:
	;
	v36 = int32(12)
	goto L5
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27+v36)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	if v40 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v83 == int32(-1) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	F_get_type_io_data(m, v38, int32(1), v56+int32(4), v56+int32(6), v56+int32(7), v56+int32(8), v56+int32(12), v56+int32(16))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v45 = F_MemoryContextAlloc(m, v43, int32(48))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v54 == v38 {
		v80 = v40
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v45
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v38 ^ int32(-1)
	v56 = v50
	goto L7
L12:
	;
	v56 = v40
	goto L7
L13:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	F_fmgr_info_cxt(m, v72, v56+int32(20), v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v38
	v80 = v56
	goto L6
L15:
	;
	v86 = int32(28)
	goto L17
L16:
	;
	v86 = int32(4)
	goto L17
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v27+v86)))
	if v83 == int32(-1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+8)))
	v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v80)+7)))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+6)))
	v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80)+4)))
	v105 = F_ArrayGetNItemsSafe(m, v88, v99)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L24
	}
L19:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	v99 = v91
	v100 = v92
	goto L18
L20:
	;
	goto L21
L21:
	;
	v94 = v27 + int32(16)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v99 = v94
	v100 = v94 + v95<<(uint(int32(2))%32)
	goto L18
L22:
	;
	m.G0 = v24 + int32(288)
	return base.I64_extend_i32_u(v1172)
L23:
	;
	if int32(0) < v88 {
		goto L86
	} else {
		goto L87
	}
L24:
	;
	if v105 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v107 = int32(0)
	if v88 <= v107 {
		v152 = int32(0)
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v389 = F_pstrdup(m, int32(_a_F_array_out_0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L82
	}
L28:
	;
	v164 = F_palloc(m, v105<<(uint(int32(2))%32))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L34
	}
L29:
	;
	v111 = int32(0)
	goto L30
L30:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v100+v111<<(uint(int32(2))%32))))
	v137 = base.B2i32(v135 != int32(1))
	if v135 != int32(1) {
		v152 = v137
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v152 = v137
	goto L28
L32:
	;
	v139 = v111 + int32(1)
	if v139 != v88 {
		v111 = v139
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v166 = F_palloc(m, v105)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_array_iter_setup(m, v24+int32(20), v27, v104, v103&int32(1), v102)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v105 <= int32(0) {
		v393 = v107
		goto L23
	} else {
		goto L37
	}
L37:
	;
	v181 = v107
	v183 = int32(0)
	goto L38
L38:
	;
	v202 = v164 + v183<<(uint(int32(2))%32)
	v207 = F_array_iter_next(m, v24+int32(20), v24+int32(80), v183)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	v393 = v384
	goto L23
L40:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+80)))
	if v209 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v383 = int32(1)
	v384 = v382 + v383
	v386 = v183 + v383
	if v386 != v105 {
		v181 = v384
		v183 = v386
		goto L38
	} else {
		goto L81
	}
L42:
	;
	v213 = F_pstrdup(m, int32(_a_F_array_out_1))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v221 = F_OutputFunctionCall(m, v80+int32(20), v207)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v213
	v217 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v183+v166))) = uint8(v217)
	v382 = v181 + int32(4)
	goto L41
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v221
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	if v224 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v228 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v183+v166))) = uint8(v228)
	v382 = v181 + int32(2)
	goto L41
L48:
	;
	goto L49
L49:
	;
	v235 = v221
	v236 = int32(_a_F_array_out_1)
	goto L51
L50:
	;
	v275 = base.B2i32(v273 == int32(0))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	if v277 != 0 {
		goto L63
	} else {
		goto L64
	}
L51:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
	if v239 == v240 {
		v262 = v239
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v273 = int32(0)
	goto L50
L53:
	;
	v264 = int32(1)
	if v262 != 0 {
		v235 = v235 + v264
		v236 = v236 + v264
		goto L51
	} else {
		goto L62
	}
L54:
	;
	if base.Ui32((v239-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v250 = v239 | int32(32)
	goto L57
L56:
	;
	v250 = v239
	goto L57
L57:
	;
	if base.Ui32((v240-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v259 = v240 | int32(32)
	goto L60
L59:
	;
	v259 = v240
	goto L60
L60:
	;
	if v250 == v259 {
		v262 = v250
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v273 = v250 - v259
	goto L50
L62:
	;
	goto L52
L63:
	;
	v278 = v277
	v279 = v276
	v280 = v181
	v281 = v275
	goto L66
L64:
	;
	v335 = v181
	v336 = v275
	goto L65
L65:
	;
	v356 = v336 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v183+v166))) = uint8(v356)
	if v356 != 0 {
		goto L78
	} else {
		goto L79
	}
L66:
	;
	v302 = v278 & int32(255)
	switch v302 - int32(123) {
	case 0, 2:
		goto L70
	case 1:
		goto L71
	default:
		goto L72
	}
L67:
	;
	v335 = v329
	v336 = v328
	goto L65
L68:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+1)))
	if v330 != 0 {
		v278 = v330
		v279 = v279 + int32(1)
		v280 = v329
		v281 = v328
		goto L66
	} else {
		goto L77
	}
L69:
	;
	v328 = v327
	v329 = v280 + int32(1)
	goto L68
L70:
	;
	v327 = int32(1)
	goto L69
L71:
	;
	if v302 == v101 {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	if base.B2i32(v302 != int32(92))&base.B2i32(v302 != int32(34)) != 0 {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v328 = int32(1)
	v329 = v280 + int32(2)
	goto L68
L74:
	;
	v314 = base.I32_extend8_s(v278)
	goto L75
L75:
	;
	if base.B2i32(v314 == int32(32))|base.B2i32(base.Ui32((v314-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		v327 = v281
		goto L69
	} else {
		goto L76
	}
L76:
	;
	goto L70
L77:
	;
	goto L67
L78:
	;
	v360 = v335 + int32(2)
	goto L80
L79:
	;
	v360 = v335
	goto L80
L80:
	;
	v382 = v360
	goto L41
L81:
	;
	goto L39
L82:
	;
	v1172 = v389
	goto L22
L83:
	;
	v748 = int32(123)
	*(*uint16)(unsafe.Add(mBase, uint32(v747))) = uint16(v748)
	if v88 <= int32(0) {
		goto L129
	} else {
		goto L130
	}
L84:
	;
	v640 = int32(61)
	*(*uint16)(unsafe.Add(mBase, uint32(v619))) = uint16(v640)
	v646 = F_palloc(m, v619+(v621-v24)-int32(79))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L107
	}
L85:
	;
	v617 = F_palloc(m, v598)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L106
	}
L86:
	;
	v415 = v88 & int32(3)
	v416 = int32(0)
	v417 = int32(1)
	if base.Ui32(int32(4)) <= base.Ui32(v88) {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	goto L88
L88:
	;
	v592 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+80)) = uint8(v592)
	if v152 != 0 {
		v619 = v24 + int32(80)
		v621 = v393
		goto L84
	} else {
		goto L105
	}
L89:
	;
	v542 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+80)) = uint8(v542)
	v547 = v527<<(uint(int32(1))%32) + v393
	if v152 == v542 {
		v598 = v547
		goto L85
	} else {
		goto L100
	}
L90:
	;
	v425 = v417
	v426 = v416
	v429 = int32(0)
	v431 = v416
	goto L93
L91:
	;
	v468 = v417
	v469 = v416
	v474 = v416
	goto L92
L92:
	;
	v489 = v468
	v490 = v469
	v492 = v416
	v495 = v474
	goto L97
L93:
	;
	v448 = v99 + v426<<(uint(int32(2))%32)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	v451 = v450 * v425
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	v453 = v451 * v452
	v454 = v449 * v453
	v458 = v454 + (v453 + (v451 + (v425 + v431)))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v448)+12))
	v460 = v459 * v454
	v461 = int32(4)
	v462 = v426 + v461
	v464 = v429 + v461
	if v464 != v88&int32(2147483644) {
		v425 = v460
		v426 = v462
		v429 = v464
		v431 = v458
		goto L93
	} else {
		goto L95
	}
L94:
	;
	if v415 == int32(0) {
		v527 = v458
		goto L89
	} else {
		goto L96
	}
L95:
	;
	goto L94
L96:
	;
	v468 = v460
	v469 = v462
	v474 = v458
	goto L92
L97:
	;
	v510 = v489 + v495
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v99+v490<<(uint(int32(2))%32))))
	v516 = int32(1)
	v519 = v492 + v516
	if v519 != v415 {
		v489 = v514 * v489
		v490 = v490 + v516
		v492 = v519
		v495 = v510
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v527 = v510
	goto L89
L99:
	;
	goto L98
L100:
	;
	v552 = v24 + int32(80)
	v553 = v542
	goto L101
L101:
	;
	v574 = v553 << (uint(int32(2)) % 32)
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v99+v574)))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v574+v100)))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v578
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v578 + v576 - int32(1)
	v585 = F_pg_sprintf(m, v552, int32(_a_F_array_out_2), v24)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L103
	}
L102:
	;
	v619 = v588
	v621 = v547
	goto L84
L103:
	;
	v587 = F_strlen(m, v552)
	mBase = m.M
	v588 = v587 + v552
	v590 = v553 + int32(1)
	if v590 != v88 {
		v552 = v588
		v553 = v590
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v598 = v393
	goto L85
L106:
	;
	v733 = v617
	v747 = v617
	goto L83
L107:
	;
	v649 = v24 + int32(80)
	if (v649^v646)&int32(3) != 0 {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v724 = F_strlen(m, v646)
	mBase = m.M
	v733 = v646
	v747 = v724 + v646
	goto L83
L109:
	;
	goto L108
L110:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v704))) = uint8(v703)
	if v703&int32(255) == int32(0) {
		goto L109
	} else {
		goto L125
	}
L111:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649))))
	v702 = v649
	v703 = v655
	v704 = v646
	goto L110
L112:
	;
	goto L113
L113:
	;
	if v649&int32(3) != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v659 = v649
	v661 = v646
	goto L117
L115:
	;
	v673 = v649
	v675 = v646
	goto L116
L116:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v673)))
	v680 = int32(-2139062144)
	if (int32(16843008)-v677|v677)&v680 != v680 {
		v702 = v673
		v703 = v677
		v704 = v675
		goto L110
	} else {
		goto L121
	}
L117:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v659))))
	*(*uint8)(unsafe.Add(mBase, uint32(v661))) = uint8(v662)
	if v662 == int32(0) {
		goto L109
	} else {
		goto L119
	}
L118:
	;
	v673 = v669
	v675 = v667
	goto L116
L119:
	;
	v666 = int32(1)
	v667 = v661 + v666
	v669 = v659 + v666
	if v669&int32(3) != 0 {
		v659 = v669
		v661 = v667
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v685 = v673
	v686 = v677
	v687 = v675
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v687))) = v686
	v689 = int32(4)
	v690 = v687 + v689
	v692 = v685 + v689
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v685)+4))
	v697 = int32(-2139062144)
	if (int32(16843008)-v694|v694)&v697 == v697 {
		v685 = v692
		v686 = v694
		v687 = v690
		goto L122
	} else {
		goto L124
	}
L123:
	;
	v702 = v692
	v703 = v694
	v704 = v690
	goto L110
L124:
	;
	goto L123
L125:
	;
	v711 = v702
	v713 = v704
	goto L126
L126:
	;
	v714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v713)+1)) = uint8(v714)
	v716 = int32(1)
	if v714 != 0 {
		v711 = v711 + v716
		v713 = v713 + v716
		goto L126
	} else {
		goto L128
	}
L127:
	;
	goto L109
L128:
	;
	goto L127
L129:
	;
	v761 = int32(1)
	v766 = v88 - v761
	v767 = int32(0)
	v769 = v747 + v761
	v770 = v767
	v773 = v767
	goto L132
L130:
	;
	v753 = v88 << (uint(int32(2)) % 32)
	if v753 == int32(0) {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	base.MemoryFill(m, v24+int32(48), int32(0), v753)
	goto L129
L132:
	;
	if v766 <= v770 {
		v881 = v769
		goto L134
	} else {
		goto L135
	}
L133:
	;
	F_pfree(m, v164)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L1
	} else {
		goto L189
	}
L134:
	;
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773+v166))))
	if v903 == int32(1) {
		goto L147
	} else {
		goto L148
	}
L135:
	;
	v796 = (v88 + (v770 ^ int32(-1))) & int32(7)
	if v796 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v797 = v769
	v799 = v770
	v803 = int32(0)
	goto L139
L137:
	;
	v827 = v769
	v829 = v770
	goto L138
L138:
	;
	if base.Ui32(v88-int32(2)-v770) < base.Ui32(int32(7)) {
		v881 = v827
		goto L134
	} else {
		goto L142
	}
L139:
	;
	v818 = int32(123)
	*(*uint16)(unsafe.Add(mBase, uint32(v797))) = uint16(v818)
	v820 = int32(1)
	v821 = v799 + v820
	v823 = v797 + v820
	v825 = v803 + v820
	if v825 != v796 {
		v797 = v823
		v799 = v821
		v803 = v825
		goto L139
	} else {
		goto L141
	}
L140:
	;
	v827 = v823
	v829 = v821
	goto L138
L141:
	;
	goto L140
L142:
	;
	v851 = v827
	v853 = v829
	goto L143
L143:
	;
	v872 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v851)+8)) = uint8(v872)
	*(*int64)(unsafe.Add(mBase, uint32(v851))) = int64(8897841259083430779)
	v876 = int32(8)
	v877 = v851 + v876
	v879 = v853 + v876
	if v879 != v766 {
		v851 = v877
		v853 = v879
		goto L143
	} else {
		goto L145
	}
L144:
	;
	v881 = v877
	goto L134
L145:
	;
	goto L144
L146:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v164+v773<<(uint(int32(2))%32))))
	F_pfree(m, v1062)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L1
	} else {
		goto L178
	}
L147:
	;
	v906 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v881))) = uint16(v906)
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v164+v773<<(uint(int32(2))%32))))
	v914 = v881 + int32(1)
	v916 = v913
	goto L150
L148:
	;
	goto L149
L149:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v164+v773<<(uint(int32(2))%32))))
	if (v960^v881)&int32(3) != 0 {
		goto L160
	} else {
		goto L161
	}
L150:
	;
	v935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v916))))
	if base.B2i32(v935 == int32(34))|base.B2i32(v935 == int32(92)) == int32(0) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v951))) = uint8(v935)
	v953 = int32(1)
	v914 = v951 + v953
	v916 = v916 + v953
	goto L150
L153:
	;
	if v935 != 0 {
		v951 = v914
		goto L152
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v947 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v914))) = uint8(v947)
	v951 = v914 + int32(1)
	goto L152
L156:
	;
	v943 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v914))) = uint16(v943)
	v1058 = v914 + int32(1)
	goto L146
L157:
	;
	v1035 = F_strlen(m, v881)
	mBase = m.M
	v1058 = v1035 + v881
	goto L146
L158:
	;
	goto L157
L159:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1015))) = uint8(v1014)
	if v1014&int32(255) == int32(0) {
		goto L158
	} else {
		goto L174
	}
L160:
	;
	v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v960))))
	v1013 = v960
	v1014 = v966
	v1015 = v881
	goto L159
L161:
	;
	goto L162
L162:
	;
	if v960&int32(3) != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v970 = v960
	v972 = v881
	goto L166
L164:
	;
	v984 = v960
	v986 = v881
	goto L165
L165:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v984)))
	v991 = int32(-2139062144)
	if (int32(16843008)-v988|v988)&v991 != v991 {
		v1013 = v984
		v1014 = v988
		v1015 = v986
		goto L159
	} else {
		goto L170
	}
L166:
	;
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970))))
	*(*uint8)(unsafe.Add(mBase, uint32(v972))) = uint8(v973)
	if v973 == int32(0) {
		goto L158
	} else {
		goto L168
	}
L167:
	;
	v984 = v980
	v986 = v978
	goto L165
L168:
	;
	v977 = int32(1)
	v978 = v972 + v977
	v980 = v970 + v977
	if v980&int32(3) != 0 {
		v970 = v980
		v972 = v978
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v996 = v984
	v997 = v988
	v998 = v986
	goto L171
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v998))) = v997
	v1000 = int32(4)
	v1001 = v998 + v1000
	v1003 = v996 + v1000
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v996)+4))
	v1008 = int32(-2139062144)
	if (int32(16843008)-v1005|v1005)&v1008 == v1008 {
		v996 = v1003
		v997 = v1005
		v998 = v1001
		goto L171
	} else {
		goto L173
	}
L172:
	;
	v1013 = v1003
	v1014 = v1005
	v1015 = v1001
	goto L159
L173:
	;
	goto L172
L174:
	;
	v1022 = v1013
	v1024 = v1015
	goto L175
L175:
	;
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1022)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1024)+1)) = uint8(v1025)
	v1027 = int32(1)
	if v1025 != 0 {
		v1022 = v1022 + v1027
		v1024 = v1024 + v1027
		goto L175
	} else {
		goto L177
	}
L176:
	;
	goto L158
L177:
	;
	goto L176
L178:
	;
	if v766 < int32(0) {
		v1115 = v1058
		v1116 = v766
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L133
L180:
	;
	if v1116 != int32(-1) {
		v769 = v1115
		v770 = v1116
		v773 = v773 + int32(1)
		goto L132
	} else {
		goto L188
	}
L181:
	;
	v1067 = v1058
	v1068 = v766
	goto L182
L182:
	;
	v1089 = v1068 << (uint(int32(2)) % 32)
	v1092 = v1089 + (v24 + int32(48))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1092)))
	v1095 = v1093 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1092))) = v1095
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1089+v99)))
	if v1095 < v1098 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	goto L179
L184:
	;
	v1100 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1067)+1)) = uint8(v1100)
	*(*uint8)(unsafe.Add(mBase, uint32(v1067))) = uint8(v101)
	v1115 = v1067 + int32(1)
	v1116 = v1068
	goto L180
L185:
	;
	goto L186
L186:
	;
	v1105 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1092))) = v1105
	v1107 = int32(125)
	*(*uint16)(unsafe.Add(mBase, uint32(v1067))) = uint16(v1107)
	v1109 = int32(1)
	if v1105 < v1068 {
		v1067 = v1067 + v1109
		v1068 = v1068 - v1109
		goto L182
	} else {
		goto L187
	}
L187:
	;
	goto L183
L188:
	;
	goto L179
L189:
	;
	F_pfree(m, v166)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	v1172 = v733
	goto L22
}
func F_array_prepend(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
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
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v12 == int32(0) {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v16 = v15
	} else {
		v16 = int64(0)
	}
	v17 = int32(1)
	v19 = F_fetch_array_arg_replace_nulls(m, l0, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int64(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
		switch v23 {
		case 0:
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(1)
			v64 = v17
			v66 = int32(12)
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
			v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v74)+4)))
			v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+6)))
			v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v74)+7)))
			v78 = F_array_set_element(m, base.I64_extend_i32_u(v19+v66), int32(1), v10+v66, v16, v12, int32(-1), v75, v76, v77)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int64(0)
			} else {
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
				if v80 == int32(1) {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v83))) = v64
				} else {
				}
				m.G0 = v10 + int32(16)
				return v78
			}
		case 1:
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
			v27 = v25 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v27
			if v27 < v25 {
				v64 = v25
				v66 = int32(12)
				v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
				v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v74)+4)))
				v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+6)))
				v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v74)+7)))
				v78 = F_array_set_element(m, base.I64_extend_i32_u(v19+v66), int32(1), v10+v66, v16, v12, int32(-1), v75, v76, v77)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int64(0)
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
					if v80 == int32(1) {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v83))) = v64
					} else {
					}
					m.G0 = v10 + int32(16)
					return v78
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_array_prepend_0), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_array_prepend_1), int32(250), int32(_a_F_array_prepend_2))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int64(0)
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
				return int64(0)
			} else {
				F_errcode(m, int32(130))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_array_prepend_3), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_array_prepend_1), int32(260), int32(_a_F_array_prepend_2))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int64(0)
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
func F_array_remove(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int64
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
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v4 == int32(1) {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v7)
		return int64(0)
	} else {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v19 = int32(1)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v22 = F_array_replace_internal(m, v14, v11, v12, int64(0), v19, v19, v21, l0)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v22)
			}
		}
	}
}
func F_array_reverse(m *base.Module, l0 int32) int64 {
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
	var v115 int64
	_ = v115
	var v116 int64
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int64(0)
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
		if v27 <= int32(0) {
			v206 = v23
			m.G0 = v20 + int32(80)
			return base.I64_extend_i32_u(v206)
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
			if v30 < int32(2) {
				v206 = v23
				m.G0 = v20 + int32(80)
				return base.I64_extend_i32_u(v206)
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
							return int64(0)
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
										v99 = v92 + v89<<(uint(int32(3))%32)
										v101 = v71
										v102 = v72
										v108 = int32(0)
										for {
											v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
											v115 = *(*int64)(unsafe.Add(mBase, uint32(v102)))
											v116 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
											*(*int64)(unsafe.Add(mBase, uint32(v102))) = v116
											v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
											*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v118)
											*(*int64)(unsafe.Add(mBase, uint32(v99))) = v115
											*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v114)
											v122 = int32(1)
											v124 = int32(8)
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
								return int64(0)
							} else {
								v197 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
								F_pfree(m, v197)
								mBase = m.M
								v199 = m.ExcPending
								if v199 != 0 {
									return int64(0)
								} else {
									v200 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
									F_pfree(m, v200)
									mBase = m.M
									v202 = m.ExcPending
									if v202 != 0 {
										return int64(0)
									} else {
										v206 = v195
										m.G0 = v20 + int32(80)
										return base.I64_extend_i32_u(v206)
									}
								}
							}
						}
					} else {
						v39 = F_lookup_type_cache(m, v33, int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
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
								return int64(0)
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
											v99 = v92 + v89<<(uint(int32(3))%32)
											v101 = v71
											v102 = v72
											v108 = int32(0)
											for {
												v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
												v115 = *(*int64)(unsafe.Add(mBase, uint32(v102)))
												v116 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
												*(*int64)(unsafe.Add(mBase, uint32(v102))) = v116
												v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
												*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v118)
												*(*int64)(unsafe.Add(mBase, uint32(v99))) = v115
												*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v114)
												v122 = int32(1)
												v124 = int32(8)
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
									return int64(0)
								} else {
									v197 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
									F_pfree(m, v197)
									mBase = m.M
									v199 = m.ExcPending
									if v199 != 0 {
										return int64(0)
									} else {
										v200 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
										F_pfree(m, v200)
										mBase = m.M
										v202 = m.ExcPending
										if v202 != 0 {
											return int64(0)
										} else {
											v206 = v195
											m.G0 = v20 + int32(80)
											return base.I64_extend_i32_u(v206)
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
						return int64(0)
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
							return int64(0)
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
										v99 = v92 + v89<<(uint(int32(3))%32)
										v101 = v71
										v102 = v72
										v108 = int32(0)
										for {
											v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
											v115 = *(*int64)(unsafe.Add(mBase, uint32(v102)))
											v116 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
											*(*int64)(unsafe.Add(mBase, uint32(v102))) = v116
											v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
											*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v118)
											*(*int64)(unsafe.Add(mBase, uint32(v99))) = v115
											*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v114)
											v122 = int32(1)
											v124 = int32(8)
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
								return int64(0)
							} else {
								v197 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
								F_pfree(m, v197)
								mBase = m.M
								v199 = m.ExcPending
								if v199 != 0 {
									return int64(0)
								} else {
									v200 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
									F_pfree(m, v200)
									mBase = m.M
									v202 = m.ExcPending
									if v202 != 0 {
										return int64(0)
									} else {
										v206 = v195
										m.G0 = v20 + int32(80)
										return base.I64_extend_i32_u(v206)
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
func F_array_sample(m *base.Module, l0 int32) int64 {
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		if int32(0) < v16 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v20 = v19
		} else {
			v20 = int32(0)
		}
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v40)
					}
				} else {
					v34 = F_lookup_type_cache(m, v28, int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v34
						v38 = v34
						v40 = F_array_shuffle_n(m, v12, v21, int32(0), v28, v38)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int64(0)
						} else {
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v40)
						}
					}
				}
			} else {
				v34 = F_lookup_type_cache(m, v28, int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int64(0)
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v34
					v38 = v34
					v40 = F_array_shuffle_n(m, v12, v21, int32(0), v28, v38)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v40)
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v20
					F_errmsg(m, int32(_a_F_array_sample_0), v9)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_array_sample_1), int32(1759), int32(_a_F_array_sample_2))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int64(0)
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v406 int32
	_ = v406
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v569 int32
	_ = v569
	var v584 int32
	_ = v584
	v14 = m.G0
	v16 = v14 - int32(144)
	m.G0 = v16
	switch l8 - int32(99) {
	case 0:
		v39 = int32(1)
		goto L1
	case 1:
		goto L4
	default:
		goto L3
	case 6:
		goto L5
	case 16:
		goto L2
	}
L1:
	;
	v41 = v16 + int32(112)
	v42 = int32(0)
	if l2 <= v42 {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v39 = int32(2)
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v39 = int32(8)
	goto L1
L5:
	;
	v39 = int32(4)
	goto L1
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l8
	F_errmsg_internal(m, int32(_a_F_array_slice_size_0), v16)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_errfinish(m, int32(_a_F_array_slice_size_1), int32(322), int32(_a_F_array_slice_size_2))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	v118 = int32(0)
	if l1|base.B2i32(l7 <= v118) == v118 {
		goto L21
	} else {
		goto L22
	}
L11:
	;
	goto L10
L12:
	;
	if l2 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v58 = v42
	v61 = v42
	goto L16
L14:
	;
	v95 = v42
	goto L15
L15:
	;
	v100 = v95 << (uint(int32(2)) % 32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100+l6)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100+l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v41+v100))) = v103 - v105 + int32(1)
	goto L11
L16:
	;
	v62 = int32(2)
	v63 = v58 << (uint(v62) % 32)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63+l6)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63+l5)))
	v70 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v41+v63))) = v66 - v68 + v70
	v74 = v63 | int32(4)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74+l6)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74+l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v41+v74))) = v77 - v79 + v70
	v85 = v58 + v62
	v87 = v61 + v62
	if v87 != l2&int32(2147483646) {
		v58 = v85
		v61 = v87
		goto L16
	} else {
		goto L18
	}
L17:
	;
	if l2&int32(1) == int32(0) {
		goto L11
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v95 = v85
	goto L15
L20:
	;
	m.G0 = v16 + int32(144)
	return v584
L21:
	;
	v123 = F_ArrayGetNItemsSafe(m, l2, v41)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v132 = int32(0)
	v142 = l2 - int32(1)
	if v142 < v132 {
		v220 = v132
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v584 = v123 * ((l7 + v39 - int32(1)) & (int32(0) - v39))
	goto L20
L25:
	;
	v226 = F_array_seek(m, l0, v132, l1, v220, l7, l8)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L6
	} else {
		goto L35
	}
L26:
	;
	goto L25
L27:
	;
	v145 = int32(1)
	if v142 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v154 = v142
	v155 = v145
	v156 = v132
	v161 = v132
	goto L31
L29:
	;
	v197 = v142
	v198 = v145
	v199 = v132
	goto L30
L30:
	;
	v206 = v197 << (uint(int32(2)) % 32)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l5+v206)))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v206+l4)))
	v220 = (v208-v210)*v198 + v199
	goto L26
L31:
	;
	v162 = int32(2)
	v163 = v154 << (uint(v162) % 32)
	v165 = v163 - int32(4)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l5+v165)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l4+v165)))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v163+l3)))
	v173 = v172 * v155
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v163+l5)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v163+l4)))
	v182 = (v167-v169)*v173 + ((v176-v178)*v155 + v156)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l3+v165)))
	v185 = v184 * v173
	v187 = v154 - v162
	v189 = v161 + v162
	if v189 != l2&int32(-2) {
		v154 = v187
		v155 = v185
		v156 = v182
		v161 = v189
		goto L31
	} else {
		goto L33
	}
L32:
	;
	if l2&int32(1) == int32(0) {
		v220 = v182
		goto L26
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	v197 = v187
	v198 = v185
	v199 = v182
	goto L30
L35:
	;
	v229 = v16 + int32(80)
	v233 = int32(2)
	v234 = l2 << (uint(v233) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v229+v234-int32(4)))) = int32(1)
	v241 = l2 - v233
	if v241 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v299 = v16 + int32(48)
	v301 = v16 + int32(112)
	v302 = int32(0)
	v308 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v299+l2<<(uint(v308)%32)-int32(4)))) = v302
	v316 = l2 - v308
	if v302 <= v316 {
		goto L47
	} else {
		goto L48
	}
L37:
	;
	goto L36
L38:
	;
	if l2&int32(1) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v252 = v234 - int32(4)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l3+v252)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v229+v252)))
	*(*int32)(unsafe.Add(mBase, uint32(v229+v241<<(uint(int32(2))%32)))) = v254 * v256
	v261 = l2 - int32(3)
	goto L41
L40:
	;
	v261 = v241
	goto L41
L41:
	;
	if v241 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v267 = v261
	goto L43
L43:
	;
	v270 = int32(2)
	v271 = v267 << (uint(v270) % 32)
	v274 = v271 + int32(4)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l3+v274)))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v274+v229)))
	v279 = v276 * v278
	*(*int32)(unsafe.Add(mBase, uint32(v229+v271))) = v279
	v282 = v267 - int32(1)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l3+v271)))
	*(*int32)(unsafe.Add(mBase, uint32(v229+v282<<(uint(v270)%32)))) = v287 * v279
	if v282 != 0 {
		v267 = v267 - v270
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L37
L45:
	;
	goto L44
L46:
	;
	v423 = l2 << (uint(int32(2)) % 32)
	if v423 != 0 {
		goto L62
	} else {
		goto L63
	}
L47:
	;
	v323 = v316
	v327 = v302
	goto L50
L48:
	;
	goto L49
L49:
	;
	goto L46
L50:
	;
	v330 = v323 << (uint(int32(2)) % 32)
	v331 = v299 + v330
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v229+v330)))
	v334 = int32(1)
	v335 = v333 - v334
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = v335
	v338 = v323 + v334
	if l2 <= v338 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L49
L52:
	;
	v406 = int32(1)
	if int32(0) < v323 {
		v323 = v323 - v406
		v327 = v327 + v406
		goto L50
	} else {
		goto L61
	}
L53:
	;
	if v327&int32(1) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v344 = int32(2)
	v345 = v338 << (uint(v344) % 32)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v301+v345)))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v229+v345)))
	v353 = v335 - (v347-int32(1))*v351
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = v353
	v357 = v323 + v344
	v358 = v353
	goto L56
L55:
	;
	v357 = v338
	v358 = v335
	goto L56
L56:
	;
	if v327 == int32(0) {
		goto L52
	} else {
		goto L57
	}
L57:
	;
	v365 = v357
	v366 = v358
	goto L58
L58:
	;
	v371 = int32(2)
	v372 = v365 << (uint(v371) % 32)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v301+v372)))
	v375 = int32(1)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v229+v372)))
	v380 = v366 - (v374-v375)*v378
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = v380
	v383 = v372 + int32(4)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v301+v383)))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v229+v383)))
	v391 = v380 - (v385-v375)*v389
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = v391
	v394 = v365 + v371
	if v394 != l2 {
		v365 = v394
		v366 = v391
		goto L58
	} else {
		goto L60
	}
L59:
	;
	goto L52
L60:
	;
	goto L59
L61:
	;
	goto L51
L62:
	;
	base.MemoryFill(m, v16+int32(16), int32(0), v423)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v428 = int32(0)
	v430 = int32(1)
	v441 = v220
	v442 = v226
	v445 = l2 - v430
	v448 = int32(0)
	goto L65
L65:
	;
	v453 = v16 + int32(48) + v445<<(uint(int32(2))%32)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	if v454 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v584 = v509
	goto L20
L67:
	;
	v455 = F_array_seek(m, v442, v441, l1, v454, l7, l8)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L6
	} else {
		goto L70
	}
L68:
	;
	v459 = v441
	v460 = v442
	goto L69
L69:
	;
	if l1 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v459 = v457 + v441
	v460 = v455
	goto L69
L71:
	;
	v513 = v16 + int32(16)
	v515 = v16 + int32(112)
	if l2 <= int32(0) {
		goto L92
	} else {
		goto L93
	}
L72:
	;
	v462 = base.I32_div_s(v459, int32(8))
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v462))))
	if int32(base.Ui32(v464)>>(uint(v459&int32(7))%32))&int32(1) == int32(0) {
		v508 = v460
		v509 = v448
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v428 < l7 {
		v502 = l7
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L74
L76:
	;
	v504 = (v502 + (v39 - v430)) & (v428 - v39)
	v508 = v504 + v460
	v509 = v504 + v448
	goto L71
L77:
	;
	if l7 == int32(-1) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	if v474 == int32(1) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v498 = F_strlen(m, v460)
	mBase = m.M
	v502 = v498 + int32(1)
	goto L76
L81:
	;
	v478 = int32(18)
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+1)))
	if v480 == v478 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v491 = int32(1)
	if v474&v491 != 0 {
		v502 = int32(base.Ui32(v474) >> (uint(v491) % 32))
		goto L76
	} else {
		goto L90
	}
L84:
	;
	v483 = v478
	goto L86
L85:
	;
	v483 = int32(2)
	goto L86
L86:
	;
	if base.Ui32((v480-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v490 = int32(6)
	goto L89
L88:
	;
	v490 = v483
	goto L89
L89:
	;
	v502 = v490
	goto L76
L90:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	v502 = int32(base.Ui32(v495) >> (uint(int32(2)) % 32))
	goto L76
L91:
	;
	if v569 != int32(-1) {
		v441 = v459 + int32(1)
		v442 = v508
		v445 = v569
		v448 = v509
		goto L65
	} else {
		goto L106
	}
L92:
	;
	v569 = int32(-1)
	goto L91
L93:
	;
	goto L94
L94:
	;
	v521 = int32(1)
	v522 = l2 - v521
	v524 = v522 << (uint(int32(2)) % 32)
	v525 = v513 + v524
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v515+v524)))
	v531 = base.I32_rem_s(v526+v521, v530)
	*(*int32)(unsafe.Add(mBase, uint32(v525))) = v531
	if v522 != 0 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v569 = v559
	goto L91
L96:
	;
	v533 = v522
	v536 = v531
	goto L99
L97:
	;
	goto L98
L98:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	if v557 != 0 {
		goto L103
	} else {
		goto L104
	}
L99:
	;
	if v536 != 0 {
		v559 = v533
		goto L95
	} else {
		goto L101
	}
L100:
	;
	goto L98
L101:
	;
	v538 = int32(1)
	v539 = v533 - v538
	v541 = v539 << (uint(int32(2)) % 32)
	v542 = v513 + v541
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v515+v541)))
	v548 = base.I32_rem_s(v543+v538, v547)
	*(*int32)(unsafe.Add(mBase, uint32(v542))) = v548
	if v539 != 0 {
		v533 = v539
		v536 = v548
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v558 = int32(0)
	goto L105
L104:
	;
	v558 = int32(-1)
	goto L105
L105:
	;
	v559 = v558
	goto L95
L106:
	;
	goto L66
}
func F_array_subscript_assign_slice(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int64
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
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v513 int32
	_ = v513
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v653 int32
	_ = v653
	var v663 int32
	_ = v663
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1163 int32
	_ = v1163
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1180 int32
	_ = v1180
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1201 int32
	_ = v1201
	var v1208 int32
	_ = v1208
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1379 int32
	_ = v1379
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1470 int32
	_ = v1470
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1490 int32
	_ = v1490
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1516 int32
	_ = v1516
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1545 int32
	_ = v1545
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1582 int32
	_ = v1582
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1613 int32
	_ = v1613
	var v1621 int32
	_ = v1621
	var v1628 int32
	_ = v1628
	var v1632 int32
	_ = v1632
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1643 int32
	_ = v1643
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1699 int32
	_ = v1699
	var v1711 int32
	_ = v1711
	var v1727 int32
	_ = v1727
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1789 int32
	_ = v1789
	var v1795 int32
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1818 int32
	_ = v1818
	var v1828 int32
	_ = v1828
	var v1834 int32
	_ = v1834
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1854 int32
	_ = v1854
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1868 int32
	_ = v1868
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1929 int32
	_ = v1929
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1941 int32
	_ = v1941
	var v1946 int32
	_ = v1946
	var v2017 int32
	_ = v2017
	var v2022 int32
	_ = v2022
	var v2024 int32
	_ = v2024
	var v2032 int32
	_ = v2032
	var v2052 int32
	_ = v2052
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2065 int32
	_ = v2065
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2075 int32
	_ = v2075
	var v2077 int32
	_ = v2077
	var v2083 int32
	_ = v2083
	var v2085 int32
	_ = v2085
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2122 int32
	_ = v2122
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2158 int32
	_ = v2158
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2166 int32
	_ = v2166
	var v2174 int32
	_ = v2174
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2184 int32
	_ = v2184
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2235 int32
	_ = v2235
	var v2238 int32
	_ = v2238
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2257 int32
	_ = v2257
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2270 int32
	_ = v2270
	var v2276 int32
	_ = v2276
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2283 int32
	_ = v2283
	var v2286 int32
	_ = v2286
	var v2290 int32
	_ = v2290
	var v2294 int32
	_ = v2294
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2324 int32
	_ = v2324
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2336 int32
	_ = v2336
	var v2338 int32
	_ = v2338
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2345 int32
	_ = v2345
	var v2349 int32
	_ = v2349
	var v2351 int32
	_ = v2351
	var v2358 int32
	_ = v2358
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2384 int32
	_ = v2384
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2398 int32
	_ = v2398
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2425 int32
	_ = v2425
	var v2429 int32
	_ = v2429
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2499 int32
	_ = v2499
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2508 int32
	_ = v2508
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2525 int32
	_ = v2525
	var v2552 int32
	_ = v2552
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2556 int32
	_ = v2556
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2568 int32
	_ = v2568
	var v2572 int32
	_ = v2572
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2577 int32
	_ = v2577
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2614 int32
	_ = v2614
	var v2616 int32
	_ = v2616
	var v2622 int32
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2630 int32
	_ = v2630
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2639 int32
	_ = v2639
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2654 int32
	_ = v2654
	var v2660 int32
	_ = v2660
	var v2728 int32
	_ = v2728
	var v2729 int32
	_ = v2729
	var v2732 int32
	_ = v2732
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2739 int32
	_ = v2739
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2744 int32
	_ = v2744
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2780 int32
	_ = v2780
	var v2786 int32
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2797 int32
	_ = v2797
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2802 int32
	_ = v2802
	var v2803 int32
	_ = v2803
	var v2805 int32
	_ = v2805
	var v2807 int32
	_ = v2807
	var v2810 int32
	_ = v2810
	var v2827 int32
	_ = v2827
	var v2842 int32
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2844 int32
	_ = v2844
	var v2846 int32
	_ = v2846
	var v2852 int32
	_ = v2852
	var v2853 int32
	_ = v2853
	var v2856 int32
	_ = v2856
	var v2857 int32
	_ = v2857
	var v2858 int32
	_ = v2858
	var v2860 int32
	_ = v2860
	var v2865 int32
	_ = v2865
	var v2866 int32
	_ = v2866
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2871 int32
	_ = v2871
	var v2880 int32
	_ = v2880
	var v2882 int32
	_ = v2882
	var v2923 int32
	_ = v2923
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2928 int32
	_ = v2928
	var v2929 int32
	_ = v2929
	var v2931 int32
	_ = v2931
	var v2933 int32
	_ = v2933
	var v2936 int32
	_ = v2936
	var v2953 int32
	_ = v2953
	var v2968 int32
	_ = v2968
	var v2969 int32
	_ = v2969
	var v2970 int32
	_ = v2970
	var v2972 int32
	_ = v2972
	var v2978 int32
	_ = v2978
	var v2979 int32
	_ = v2979
	var v2982 int32
	_ = v2982
	var v2983 int32
	_ = v2983
	var v2984 int32
	_ = v2984
	var v2986 int32
	_ = v2986
	var v2991 int32
	_ = v2991
	var v2992 int32
	_ = v2992
	var v2995 int32
	_ = v2995
	var v2996 int32
	_ = v2996
	var v2997 int32
	_ = v2997
	var v3006 int32
	_ = v3006
	var v3011 int32
	_ = v3011
	var v3117 int64
	_ = v3117
	var v3122 int32
	_ = v3122
	v4 = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	if v4 < v44 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v68 = v43 + int32(12)
	v70 = v43 + int32(36)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v42)+40))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+48)))
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+6)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+8)))
	v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v43)+9)))
	v78 = m.G0
	v80 = v78 - int32(272)
	m.G0 = v80
	if v74 != 0 {
		v3117 = v65
		goto L11
	} else {
		goto L12
	}
L3:
	;
	if v39&int32(1) != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v39&int32(1) == int32(0) {
		v64 = v44
		v65 = v41
		goto L2
	} else {
		goto L8
	}
L6:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+48)))
	if v49 == int32(0) {
		v64 = v44
		v65 = v41
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v57 = F_construct_empty_array(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v60 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v60)
	v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+4)))
	v64 = v63
	v65 = base.I64_extend_i32_u(v57)
	goto L2
L11:
	;
	m.G0 = v80 + int32(272)
	v3122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v3122))) = v3117
	goto L1
L12:
	;
	if v64 <= int32(0) {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	v1001 = v796 + v814 + v995 - v985
	v1002 = F_palloc0(m, v1001)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L9
	} else {
		goto L206
	}
L14:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v935 = v930 << (uint(int32(3)) % 32)
	if v815 != 0 {
		goto L183
	} else {
		goto L184
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L9
	} else {
		goto L179
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L9
	} else {
		goto L175
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L9
	} else {
		goto L171
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L9
	} else {
		goto L167
	}
L19:
	;
	v685 = v80 + int32(112)
	v686 = F_ArrayGetNItemsSafe(m, v90, v685)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L9
	} else {
		goto L132
	}
L20:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v586 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L21:
	;
	if v90 <= v475 {
		v653 = v256
		v663 = v4
		goto L19
	} else {
		goto L104
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L9
	} else {
		goto L100
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L9
	} else {
		goto L96
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L9
	} else {
		goto L92
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L9
	} else {
		goto L88
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L9
	} else {
		goto L84
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L9
	} else {
		goto L79
	}
L28:
	;
	v85 = F_pg_detoast_datum(m, base.I32_wrap_i64(v65))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
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
	v358 = m.ExcPending
	if v358 != 0 {
		goto L9
	} else {
		goto L75
	}
L31:
	;
	v88 = F_pg_detoast_datum(m, base.I32_wrap_i64(v73))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v90 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	F_deconstruct_array(m, v88, v75, v76, v77, v80+int32(240), v80+int32(208), v80+int32(176))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L9
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if base.B2i32(v90 < v66)|base.B2i32(base.Ui32(int32(7)) <= base.Ui32(v90)) != 0 {
		goto L24
	} else {
		goto L50
	}
L36:
	;
	if int32(0) < v66 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v107 = v4
	goto L40
L38:
	;
	goto L39
L39:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v80)+176))
	v214 = v80 + int32(112)
	v215 = F_ArrayGetNItemsSafe(m, v66, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L9
	} else {
		goto L47
	}
L40:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v71))))
	if v142 != int32(1) {
		goto L27
	} else {
		goto L42
	}
L41:
	;
	goto L39
L42:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v72))))
	if v146 == int32(0) {
		goto L27
	} else {
		goto L43
	}
L43:
	;
	v150 = v107 << (uint(int32(2)) % 32)
	v153 = v150 + (v80 + int32(112))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v150+v68)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v150+v70)))
	v158 = v155 - v157
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v158
	if base.B2i32(v158 < v155)^base.B2i32(int32(0) < v157) != 0 {
		goto L26
	} else {
		goto L44
	}
L44:
	;
	v165 = v158 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v165
	if v165 < v158 {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80+int32(80)+v150))) = v157
	v173 = v107 + int32(1)
	if v173 != v66 {
		v107 = v173
		goto L40
	} else {
		goto L46
	}
L46:
	;
	goto L41
L47:
	;
	if v212 < v215 {
		goto L25
	} else {
		goto L48
	}
L48:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v80)+240))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v80)+208))
	v222 = F_construct_md_array(m, v218, v219, v66, v214, v80+int32(80), v93, v75, v76, v77)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	v3117 = base.I64_extend_i32_u(v222)
	goto L11
L50:
	;
	v230 = v85 + int32(16)
	v232 = v90 << (uint(int32(2)) % 32)
	v233 = int32(0)
	v234 = base.B2i32(v232 == v233)
	if v234 == v233 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	base.MemoryCopy(m, v80+int32(112), v230, v232)
	goto L53
L52:
	;
	goto L53
L53:
	;
	if v234 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	base.MemoryCopy(m, v80+int32(80), v230+v244<<(uint(int32(2))%32), v232)
	goto L56
L55:
	;
	goto L56
L56:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if v250 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v256 = base.B2i32(v253 != int32(0))
	goto L59
L58:
	;
	v256 = int32(1)
	goto L59
L59:
	;
	if v90 == int32(1) {
		goto L20
	} else {
		goto L60
	}
L60:
	;
	v259 = int32(0)
	if v66 <= v259 {
		v475 = v259
		goto L21
	} else {
		goto L61
	}
L61:
	;
	v265 = v4
	goto L62
L62:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+v72))))
	if v300 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v475 = v66
	goto L21
L64:
	;
	v304 = v265 << (uint(int32(2)) % 32)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(80)+v304)))
	*(*int32)(unsafe.Add(mBase, uint32(v70+v304))) = v309
	goto L66
L65:
	;
	goto L66
L66:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+v71))))
	if v313 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v337 = v265 << (uint(int32(2)) % 32)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v70+v337)))
	if v335 < v339 {
		goto L23
	} else {
		goto L71
	}
L68:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v68+v265<<(uint(int32(2))%32))))
	v335 = v319
	goto L67
L69:
	;
	goto L70
L70:
	;
	v321 = v265 << (uint(int32(2)) % 32)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(112)+v321)))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(80)+v321)))
	v333 = v326 + v330 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v68+v321))) = v333
	v335 = v333
	goto L67
L71:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(80)+v337)))
	if v339 < v344 {
		goto L22
	} else {
		goto L72
	}
L72:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(112)+v337)))
	if v349+v344 <= v335 {
		goto L22
	} else {
		goto L73
	}
L73:
	;
	v353 = v265 + int32(1)
	if v353 != v66 {
		v265 = v353
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
	v361 = m.ExcPending
	if v361 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_0), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2860), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
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
	v377 = m.ExcPending
	if v377 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_3), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	v384 = F_errdetail(m, int32(_a_F_array_subscript_assign_slice_4), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L9
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2893), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
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
	v398 = m.ExcPending
	if v398 != 0 {
		goto L9
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = int32(134217727)
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_5), v80)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L9
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2901), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
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
	v415 = m.ExcPending
	if v415 != 0 {
		goto L9
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_6), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L9
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2910), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
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
	v431 = m.ExcPending
	if v431 != 0 {
		goto L9
	} else {
		goto L93
	}
L93:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_7), int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L9
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2920), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
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
	v447 = m.ExcPending
	if v447 != 0 {
		goto L9
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_8), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L9
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2990), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
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
	v463 = m.ExcPending
	if v463 != 0 {
		goto L9
	} else {
		goto L101
	}
L101:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_9), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L9
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2995), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
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
	v513 = v475
	goto L105
L105:
	;
	v549 = v513 << (uint(int32(2)) % 32)
	v550 = v70 + v549
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(80)+v549)))
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = v554
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v80+int32(112)+v549)))
	v563 = v554 + v560 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v549+v68))) = v563
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	if v563 < v565 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L9
	} else {
		goto L110
	}
L107:
	;
	goto L106
L108:
	;
	v568 = v513 + int32(1)
	if v90 != v568 {
		v513 = v568
		goto L105
	} else {
		goto L109
	}
L109:
	;
	v653 = v256
	v663 = v4
	goto L19
L110:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L9
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_8), int32(0))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L9
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(3005), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
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
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v80)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v589
	goto L116
L115:
	;
	goto L116
L116:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v591 == int32(1) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v601 < v602 {
		goto L18
	} else {
		goto L121
	}
L118:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v601 = v594
	goto L117
L119:
	;
	goto L120
L120:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v80)+112))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v80)+80))
	v599 = v595 + v596 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v599
	v601 = v599
	goto L117
L121:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v80)+80))
	if v604 <= v602 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v628 = v623 + v625
	if v601 < v628 {
		v653 = v626
		v663 = v627
		goto L19
	} else {
		goto L128
	}
L123:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v80)+112))
	v623 = v606
	v625 = v604
	v626 = v256
	v627 = v4
	goto L122
L124:
	;
	goto L125
L125:
	;
	v607 = v604 - v602
	if base.B2i32(v607 < v604)^base.B2i32(int32(0) < v602) != 0 {
		goto L17
	} else {
		goto L126
	}
L126:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v80)+112))
	v613 = v612 + v607
	*(*int32)(unsafe.Add(mBase, uint32(v80)+112)) = v613
	if base.B2i32(v607 < int32(0)) != base.B2i32(v613 < v612) {
		goto L17
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+80)) = v602
	v623 = v613
	v625 = v602
	v626 = base.B2i32(int32(1) < v607) | v256
	v627 = v607
	goto L122
L128:
	;
	v632 = v601 - v628
	if base.B2i32(int32(0) < v628)^base.B2i32(v632 < v601) != 0 {
		goto L16
	} else {
		goto L129
	}
L129:
	;
	v636 = v632 + int32(1)
	if v636 < v632 {
		goto L16
	} else {
		goto L130
	}
L130:
	;
	v638 = v623 + v636
	*(*int32)(unsafe.Add(mBase, uint32(v80)+112)) = v638
	if base.B2i32(v636 < int32(0)) != base.B2i32(v638 < v623) {
		goto L16
	} else {
		goto L131
	}
L131:
	;
	v653 = base.B2i32(int32(1) < v636) | v626
	v663 = v627
	goto L19
L132:
	;
	F_ArrayCheckBounds(m, v90, v685, v80+int32(80))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L9
	} else {
		goto L133
	}
L133:
	;
	v693 = v80 + int32(48)
	v694 = int32(0)
	if v90 <= v694 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v770 = F_ArrayGetNItemsSafe(m, v90, v693)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L9
	} else {
		goto L144
	}
L135:
	;
	goto L134
L136:
	;
	if v90 != int32(1) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v710 = v694
	v713 = v694
	goto L140
L138:
	;
	v747 = v694
	goto L139
L139:
	;
	v752 = v747 << (uint(int32(2)) % 32)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v752+v68)))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v752+v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v693+v752))) = v755 - v757 + int32(1)
	goto L135
L140:
	;
	v714 = int32(2)
	v715 = v710 << (uint(v714) % 32)
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v715+v68)))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v715+v70)))
	v722 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v693+v715))) = v718 - v720 + v722
	v726 = v715 | int32(4)
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v726+v68)))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v726+v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v693+v726))) = v729 - v731 + v722
	v737 = v710 + v714
	v739 = v713 + v714
	if v739 != v90&int32(2147483646) {
		v710 = v737
		v713 = v739
		goto L140
	} else {
		goto L142
	}
L141:
	;
	if v90&int32(1) == int32(0) {
		goto L135
	} else {
		goto L143
	}
L142:
	;
	goto L141
L143:
	;
	v747 = v737
	goto L139
L144:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v774 = v88 + int32(16)
	v775 = F_ArrayGetNItemsSafe(m, v772, v774)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L9
	} else {
		goto L145
	}
L145:
	;
	if v775 < v770 {
		goto L15
	} else {
		goto L146
	}
L146:
	;
	v779 = v90 << (uint(int32(3)) % 32)
	if v653&int32(1) != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v801 = v799 << (uint(int32(3)) % 32)
	if v798 != 0 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v785 = base.I32_div_s(v686+int32(7), int32(8))
	v790 = (v779 + v785 + int32(23)) & int32(-8)
	v795 = v790
	v796 = v790
	goto L147
L149:
	;
	goto L150
L150:
	;
	v795 = v4
	v796 = (v779 + int32(23)) & int32(120)
	goto L147
L151:
	;
	v806 = v798
	goto L153
L152:
	;
	v806 = (v801 + int32(23)) & int32(-8)
	goto L153
L153:
	;
	v807 = v88 + v806
	v808 = int32(0)
	if v798 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v811 = v801 + v774
	goto L156
L155:
	;
	v811 = v808
	goto L156
L156:
	;
	v812 = F_array_seek(m, v807, v808, v811, v770, v75, v77)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L9
	} else {
		goto L157
	}
L157:
	;
	v814 = v812 - v807
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if v815 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v851 = F_array_slice_size(m, v843+v85, v841, v90, v80+int32(112), v80+int32(80), v70, v68, v75, v77)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L9
	} else {
		goto L166
	}
L159:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v827 = (v821<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v828 = int32(base.Ui32(v818)>>(uint(int32(2))%32)) - v827
	if int32(1) < v90 {
		v841 = int32(0)
		v843 = v827
		v844 = v828
		goto L158
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v832 = int32(2)
	v834 = int32(base.Ui32(v831)>>(uint(v832)%32)) - v815
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v90 < v832 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v930 = v821
	v931 = v827
	v932 = v828
	goto L14
L163:
	;
	v930 = v835
	v931 = v815
	v932 = v834
	goto L14
L164:
	;
	goto L165
L165:
	;
	v841 = v230 + v835<<(uint(int32(3))%32)
	v843 = v815
	v844 = v834
	goto L158
L166:
	;
	v853 = int32(0)
	v984 = v853
	v985 = v851
	v986 = v853
	v990 = v843
	v992 = v4
	v995 = v844
	v996 = int32(1)
	v997 = v4
	v998 = v853
	goto L13
L167:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L9
	} else {
		goto L168
	}
L168:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_8), int32(0))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L9
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2945), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L9
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L9
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = int32(134217727)
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_5), v80+int32(16))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L9
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2955), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L9
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L9
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+32)) = int32(134217727)
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_5), v80+int32(32))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L9
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(2970), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L9
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L9
	} else {
		goto L180
	}
L180:
	;
	F_errmsg(m, int32(_a_F_array_subscript_assign_slice_6), int32(0))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L9
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_array_subscript_assign_slice_1), int32(3022), int32(_a_F_array_subscript_assign_slice_2))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L9
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	v940 = v815
	goto L185
L184:
	;
	v940 = (v935 + int32(23)) & int32(-8)
	goto L185
L185:
	;
	v941 = v85 + v940
	v942 = int32(0)
	if v815 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v945 = v935 + v230
	goto L188
L187:
	;
	v945 = v942
	goto L188
L188:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v230+v930<<(uint(int32(2))%32))))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v950 < v949 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v952 = v949
	goto L191
L190:
	;
	v952 = v950
	goto L191
L191:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	v954 = v953 + v949
	if v952 < v954 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v956 = v952
	goto L194
L193:
	;
	v956 = v954
	goto L194
L194:
	;
	v957 = v956 - v949
	v958 = F_array_seek(m, v941, v942, v945, v957, v75, v77)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L9
	} else {
		goto L195
	}
L195:
	;
	v960 = v958 - v941
	v963 = v954 - int32(1)
	if v963 < v933 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v965 = v963
	goto L198
L197:
	;
	v965 = v933
	goto L198
L198:
	;
	if v952 <= v965 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v970 = v965 - v952 + int32(1)
	v971 = F_array_seek(m, v960+v941, v957, v945, v970, v75, v77)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L9
	} else {
		goto L202
	}
L200:
	;
	v974 = int32(0)
	v975 = v4
	goto L201
L201:
	;
	v977 = v965 + int32(1)
	if v949 < v977 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v974 = v971 - v958
	v975 = v970
	goto L201
L203:
	;
	v979 = v977
	goto L205
L204:
	;
	v979 = v949
	goto L205
L205:
	;
	v984 = v960
	v985 = v974
	v986 = v954 - v979
	v990 = v931
	v992 = v957
	v995 = v932
	v996 = v4
	v997 = v975
	v998 = v932 - (v960 + v974)
	goto L13
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1002)+8)) = v795
	*(*int32)(unsafe.Add(mBase, uint32(v1002)+4)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v1002))) = v1001 << (uint(int32(2)) % 32)
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1002)+12)) = v1009
	v1012 = v1002 + int32(16)
	if v234 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	base.MemoryCopy(m, v1012, v80+int32(112), v232)
	goto L209
L208:
	;
	goto L209
L209:
	;
	if v234 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	base.MemoryCopy(m, v1012+v232, v80+int32(80), v232)
	goto L212
L211:
	;
	goto L212
L212:
	;
	if v996 != 0 {
		goto L216
	} else {
		goto L217
	}
L213:
	;
	v3117 = base.I64_extend_i32_u(v1002)
	goto L11
L214:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3006))) = uint8(v3011)
	goto L213
L215:
	;
	v2923 = base.I32_div_s(v2091, int32(8))
	v2924 = v1051 + v2923
	v2925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2924))))
	v2926 = v2177
	v2928 = v2161
	v2929 = v2174
	v2931 = v2178
	v2933 = v2924
	v2936 = int32(1) << (uint(v2091&int32(7)) % 32)
	v2953 = v2925
	goto L525
L216:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if v1024 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	goto L218
L218:
	;
	v2232 = v990 + v85
	v2233 = v796 + v1002
	if v984 != 0 {
		goto L417
	} else {
		goto L418
	}
L219:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v1034 = (v1027<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L221
L220:
	;
	v1034 = v1024
	goto L221
L221:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	if v1035 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v1045 = (v1038<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L224
L223:
	;
	v1045 = v1035
	goto L224
L224:
	;
	if v1024 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v1051 = v230 + v1047<<(uint(int32(3))%32)
	goto L227
L226:
	;
	v1051 = int32(0)
	goto L227
L227:
	;
	if v1035 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v1057 = v774 + v1053<<(uint(int32(3))%32)
	goto L230
L229:
	;
	v1057 = int32(0)
	goto L230
L230:
	;
	v1058 = v1034 + v85
	v1060 = v90 << (uint(int32(3)) % 32)
	v1061 = v1060 + v1012
	if v795 != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v1066 = v795
	goto L233
L232:
	;
	v1066 = (v1060 + int32(23)) & int32(120)
	goto L233
L233:
	;
	v1067 = v1066 + v1002
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v1069 = F_ArrayGetNItemsSafe(m, v1068, v230)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L9
	} else {
		goto L234
	}
L234:
	;
	v1071 = int32(0)
	v1073 = v80 + int32(112)
	v1075 = v80 + int32(80)
	v1085 = v90 - int32(1)
	if v1085 < v1071 {
		v1163 = v1071
		goto L236
	} else {
		goto L237
	}
L235:
	;
	v1169 = F_array_seek(m, v1058, v1071, v1051, v1163, v75, v77)
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L9
	} else {
		goto L245
	}
L236:
	;
	goto L235
L237:
	;
	v1088 = int32(1)
	if v1085 != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v1097 = v1085
	v1098 = v1088
	v1099 = v1071
	v1104 = v1071
	goto L241
L239:
	;
	v1140 = v1085
	v1141 = v1088
	v1142 = v1071
	goto L240
L240:
	;
	v1149 = v1140 << (uint(int32(2)) % 32)
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v70+v1149)))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1149+v1075)))
	v1163 = (v1151-v1153)*v1141 + v1142
	goto L236
L241:
	;
	v1105 = int32(2)
	v1106 = v1097 << (uint(v1105) % 32)
	v1108 = v1106 - int32(4)
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v70+v1108)))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1075+v1108)))
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1106+v1073)))
	v1116 = v1115 * v1098
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1106+v70)))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1106+v1075)))
	v1125 = (v1110-v1112)*v1116 + ((v1119-v1121)*v1098 + v1099)
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1073+v1108)))
	v1128 = v1127 * v1116
	v1130 = v1097 - v1105
	v1132 = v1104 + v1105
	if v1132 != v90&int32(-2) {
		v1097 = v1130
		v1098 = v1128
		v1099 = v1125
		v1104 = v1132
		goto L241
	} else {
		goto L243
	}
L242:
	;
	if v90&int32(1) == int32(0) {
		v1163 = v1125
		goto L236
	} else {
		goto L244
	}
L243:
	;
	goto L242
L244:
	;
	v1140 = v1130
	v1141 = v1128
	v1142 = v1125
	goto L240
L245:
	;
	v1171 = v1169 - v1058
	if v1171 != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	base.MemoryCopy(m, v1067, v1058, v1171)
	goto L248
L247:
	;
	goto L248
L248:
	;
	if v795 != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1174 = v1061
	goto L251
L250:
	;
	v1174 = int32(0)
	goto L251
L251:
	;
	v1175 = int32(0)
	if base.B2i32(v795 == v1175)|base.B2i32(v1163 <= v1175) != 0 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1456 = v80 + int32(112)
	v1458 = v80 + int32(240)
	v1462 = int32(2)
	v1463 = v90 << (uint(v1462) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v1458+v1463-int32(4)))) = int32(1)
	v1470 = v90 - v1462
	if v1470 < int32(0) {
		goto L295
	} else {
		goto L296
	}
L253:
	;
	v1180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1061))))
	if v1051 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L254:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1387))) = uint8(v1385)
	goto L252
L255:
	;
	v1385 = v1379
	v1387 = v1268
	goto L254
L256:
	;
	v1379 = v1273 | int32(3)
	goto L255
L257:
	;
	v1379 = v1273 | int32(7)
	goto L255
L258:
	;
	v1379 = v1273 | int32(15)
	goto L255
L259:
	;
	v1379 = v1273 | int32(31)
	goto L255
L260:
	;
	v1379 = v1273 | int32(63)
	goto L255
L261:
	;
	v1379 = v1273 | int32(127)
	goto L255
L262:
	;
	v1385 = v1333 | int32(1)
	v1387 = v1335
	goto L254
L263:
	;
	v1268 = v1174
	v1270 = v1163
	v1273 = v1180
	goto L284
L264:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v1163) {
		goto L263
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v1185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1051))))
	v1186 = int32(1)
	v1188 = v1186
	v1190 = v1186
	v1192 = v1163
	v1193 = v1180
	v1195 = v1174
	v1201 = v1051
	v1208 = v1185
	goto L268
L267:
	;
	v1333 = v1180
	v1335 = v1174
	goto L262
L268:
	;
	if v1188&v1208 != 0 {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	if v1244 != int32(1) {
		v1385 = v1245
		v1387 = v1246
		goto L254
	} else {
		goto L283
	}
L270:
	;
	v1230 = v1190 | v1193
	goto L272
L271:
	;
	v1230 = v1193 & (v1190 ^ int32(-1))
	goto L272
L272:
	;
	v1231 = int32(1)
	v1232 = v1192 - v1231
	v1234 = v1190 << (uint(v1231) % 32)
	if v1234 == int32(256) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1195))) = uint8(v1230)
	if v1232 == int32(0) {
		goto L252
	} else {
		goto L276
	}
L274:
	;
	v1244 = v1234
	v1245 = v1230
	v1246 = v1195
	goto L275
L275:
	;
	v1248 = v1188 << (uint(int32(1)) % 32)
	if v1248 == int32(256) {
		goto L278
	} else {
		goto L279
	}
L276:
	;
	v1240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1195)+1)))
	v1241 = int32(1)
	v1244 = v1241
	v1245 = v1240
	v1246 = v1195 + v1241
	goto L275
L277:
	;
	goto L269
L278:
	;
	if v1232 == int32(0) {
		goto L277
	} else {
		goto L281
	}
L279:
	;
	v1257 = v1248
	v1258 = v1201
	v1259 = v1208
	goto L280
L280:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1192) {
		v1188 = v1257
		v1190 = v1244
		v1192 = v1232
		v1193 = v1245
		v1195 = v1246
		v1201 = v1258
		v1208 = v1259
		goto L268
	} else {
		goto L282
	}
L281:
	;
	v1253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1201)+1)))
	v1254 = int32(1)
	v1257 = v1254
	v1258 = v1201 + v1254
	v1259 = v1253
	goto L280
L282:
	;
	goto L277
L283:
	;
	goto L252
L284:
	;
	if v1270 < int32(3) {
		goto L256
	} else {
		goto L286
	}
L285:
	;
	v1333 = v1323
	v1335 = v1325
	goto L262
L286:
	;
	if v1270 == int32(3) {
		goto L257
	} else {
		goto L287
	}
L287:
	;
	if base.Ui32(v1270) < base.Ui32(int32(5)) {
		goto L258
	} else {
		goto L288
	}
L288:
	;
	if v1270 == int32(5) {
		goto L259
	} else {
		goto L289
	}
L289:
	;
	if base.Ui32(v1270) < base.Ui32(int32(7)) {
		goto L260
	} else {
		goto L290
	}
L290:
	;
	if v1270 == int32(7) {
		goto L261
	} else {
		goto L291
	}
L291:
	;
	v1317 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v1268))) = uint8(v1317)
	v1320 = v1270 - int32(8)
	if v1320 == int32(0) {
		goto L252
	} else {
		goto L292
	}
L292:
	;
	v1323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1268)+1)))
	v1324 = int32(1)
	v1325 = v1268 + v1324
	if v1320 != v1324 {
		v1268 = v1325
		v1270 = v1320
		v1273 = v1323
		goto L284
	} else {
		goto L293
	}
L293:
	;
	goto L285
L294:
	;
	v1528 = v80 + int32(208)
	v1529 = int32(0)
	if v90 <= v1529 {
		goto L305
	} else {
		goto L306
	}
L295:
	;
	goto L294
L296:
	;
	if v90&int32(1) == int32(0) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1481 = v1463 - int32(4)
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1456+v1481)))
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v1458+v1481)))
	*(*int32)(unsafe.Add(mBase, uint32(v1458+v1470<<(uint(int32(2))%32)))) = v1483 * v1485
	v1490 = v90 - int32(3)
	goto L299
L298:
	;
	v1490 = v1470
	goto L299
L299:
	;
	if v1470 == int32(0) {
		goto L295
	} else {
		goto L300
	}
L300:
	;
	v1496 = v1490
	goto L301
L301:
	;
	v1499 = int32(2)
	v1500 = v1496 << (uint(v1499) % 32)
	v1503 = v1500 + int32(4)
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1456+v1503)))
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1503+v1458)))
	v1508 = v1505 * v1507
	*(*int32)(unsafe.Add(mBase, uint32(v1458+v1500))) = v1508
	v1511 = v1496 - int32(1)
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1456+v1500)))
	*(*int32)(unsafe.Add(mBase, uint32(v1458+v1511<<(uint(v1499)%32)))) = v1516 * v1508
	if v1511 != 0 {
		v1496 = v1496 - v1499
		goto L301
	} else {
		goto L303
	}
L302:
	;
	goto L295
L303:
	;
	goto L302
L304:
	;
	v1606 = v80 + int32(176)
	v1607 = int32(0)
	v1613 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1606+v90<<(uint(v1613)%32)-int32(4)))) = v1607
	v1621 = v90 - v1613
	if v1607 <= v1621 {
		goto L315
	} else {
		goto L316
	}
L305:
	;
	goto L304
L306:
	;
	if v90 != int32(1) {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1545 = v1529
	v1548 = v1529
	goto L310
L308:
	;
	v1582 = v1529
	goto L309
L309:
	;
	v1587 = v1582 << (uint(int32(2)) % 32)
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v1587+v68)))
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v1587+v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v1528+v1587))) = v1590 - v1592 + int32(1)
	goto L305
L310:
	;
	v1549 = int32(2)
	v1550 = v1545 << (uint(v1549) % 32)
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v1550+v68)))
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(v1550+v70)))
	v1557 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1528+v1550))) = v1553 - v1555 + v1557
	v1561 = v1550 | int32(4)
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1561+v68)))
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1561+v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v1528+v1561))) = v1564 - v1566 + v1557
	v1572 = v1545 + v1549
	v1574 = v1548 + v1549
	if v1574 != v90&int32(2147483646) {
		v1545 = v1572
		v1548 = v1574
		goto L310
	} else {
		goto L312
	}
L311:
	;
	if v90&int32(1) == int32(0) {
		goto L305
	} else {
		goto L313
	}
L312:
	;
	goto L311
L313:
	;
	v1582 = v1572
	goto L309
L314:
	;
	v1727 = int32(0)
	if v234 == v1727 {
		goto L330
	} else {
		goto L331
	}
L315:
	;
	v1628 = v1621
	v1632 = v1607
	goto L318
L316:
	;
	goto L317
L317:
	;
	goto L314
L318:
	;
	v1635 = v1628 << (uint(int32(2)) % 32)
	v1636 = v1606 + v1635
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1458+v1635)))
	v1639 = int32(1)
	v1640 = v1638 - v1639
	*(*int32)(unsafe.Add(mBase, uint32(v1636))) = v1640
	v1643 = v1628 + v1639
	if v90 <= v1643 {
		goto L320
	} else {
		goto L321
	}
L319:
	;
	goto L317
L320:
	;
	v1711 = int32(1)
	if int32(0) < v1628 {
		v1628 = v1628 - v1711
		v1632 = v1632 + v1711
		goto L318
	} else {
		goto L329
	}
L321:
	;
	if v1632&int32(1) == int32(0) {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v1649 = int32(2)
	v1650 = v1643 << (uint(v1649) % 32)
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v1528+v1650)))
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1458+v1650)))
	v1658 = v1640 - (v1652-int32(1))*v1656
	*(*int32)(unsafe.Add(mBase, uint32(v1636))) = v1658
	v1662 = v1628 + v1649
	v1663 = v1658
	goto L324
L323:
	;
	v1662 = v1643
	v1663 = v1640
	goto L324
L324:
	;
	if v1632 == int32(0) {
		goto L320
	} else {
		goto L325
	}
L325:
	;
	v1670 = v1662
	v1671 = v1663
	goto L326
L326:
	;
	v1676 = int32(2)
	v1677 = v1670 << (uint(v1676) % 32)
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1528+v1677)))
	v1680 = int32(1)
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1458+v1677)))
	v1685 = v1671 - (v1679-v1680)*v1683
	*(*int32)(unsafe.Add(mBase, uint32(v1636))) = v1685
	v1688 = v1677 + int32(4)
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1528+v1688)))
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1458+v1688)))
	v1696 = v1685 - (v1690-v1680)*v1694
	*(*int32)(unsafe.Add(mBase, uint32(v1636))) = v1696
	v1699 = v1670 + v1676
	if v1699 != v90 {
		v1670 = v1699
		v1671 = v1696
		goto L326
	} else {
		goto L328
	}
L327:
	;
	goto L320
L328:
	;
	goto L327
L329:
	;
	goto L319
L330:
	;
	base.MemoryFill(m, v80+int32(144), int32(0), v232)
	goto L332
L331:
	;
	goto L332
L332:
	;
	v1740 = v90 - int32(1)
	v1741 = v1169
	v1742 = v1163
	v1744 = v1163
	v1745 = v88 + v1045
	v1750 = v1067 + v1171
	v1751 = v1727
	goto L333
L333:
	;
	v1779 = v80 + int32(176) + v1740<<(uint(int32(2))%32)
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v1779)))
	if v1780 == int32(0) {
		goto L336
	} else {
		goto L337
	}
L334:
	;
	v2161 = v1069 - v2091
	v2162 = F_array_seek(m, v2099, v2091, v1051, v2161, v75, v77)
	mBase = m.M
	v2163 = m.ExcPending
	if v2163 != 0 {
		goto L9
	} else {
		goto L404
	}
L335:
	;
	v2058 = F_array_seek(m, v1745, v1751, v1057, int32(1), v75, v77)
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L9
	} else {
		goto L373
	}
L336:
	;
	v2022 = v1744
	v2024 = v1742
	v2032 = v1750
	v2052 = v1741
	goto L335
L337:
	;
	goto L338
L338:
	;
	v1783 = F_array_seek(m, v1741, v1744, v1051, v1780, v75, v77)
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L9
	} else {
		goto L339
	}
L339:
	;
	v1785 = v1783 - v1741
	if v1785 != 0 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	base.MemoryCopy(m, v1750, v1741, v1785)
	goto L342
L341:
	;
	goto L342
L342:
	;
	if v795 == int32(0) {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(v1779)))
	v2022 = v2017 + v1744
	v2024 = v2017 + v1742
	v2032 = v1750 + v1785
	v2052 = v1783
	goto L335
L344:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v1779)))
	if v1789 <= int32(0) {
		goto L343
	} else {
		goto L345
	}
L345:
	;
	v1795 = int32(1) << (uint(v1742&int32(7)) % 32)
	v1797 = base.I32_div_s(v1742, int32(8))
	v1798 = v1174 + v1797
	v1799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1798))))
	if v1051 != 0 {
		goto L347
	} else {
		goto L348
	}
L346:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1941))) = uint8(v1946)
	goto L343
L347:
	;
	v1805 = base.I32_div_s(v1744, int32(8))
	v1806 = v1051 + v1805
	v1807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1806))))
	v1808 = v1798
	v1810 = v1789
	v1811 = v1795
	v1813 = v1799
	v1818 = int32(1) << (uint(v1744&int32(7)) % 32)
	v1828 = v1807
	v1834 = v1806
	goto L350
L348:
	;
	goto L349
L349:
	;
	v1888 = v1798
	v1890 = v1789
	v1891 = v1795
	v1893 = v1799
	goto L366
L350:
	;
	if v1818&v1828 != 0 {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	if v1865 != int32(1) {
		v1941 = v1864
		v1946 = v1866
		goto L346
	} else {
		goto L365
	}
L352:
	;
	v1850 = v1811 | v1813
	goto L354
L353:
	;
	v1850 = v1813 & (v1811 ^ int32(-1))
	goto L354
L354:
	;
	v1851 = int32(1)
	v1852 = v1810 - v1851
	v1854 = v1811 << (uint(v1851) % 32)
	if v1854 == int32(256) {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1808))) = uint8(v1850)
	if v1852 == int32(0) {
		goto L343
	} else {
		goto L358
	}
L356:
	;
	v1864 = v1808
	v1865 = v1854
	v1866 = v1850
	goto L357
L357:
	;
	v1868 = v1818 << (uint(int32(1)) % 32)
	if v1868 == int32(256) {
		goto L360
	} else {
		goto L361
	}
L358:
	;
	v1860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1808)+1)))
	v1861 = int32(1)
	v1864 = v1808 + v1861
	v1865 = v1861
	v1866 = v1860
	goto L357
L359:
	;
	goto L351
L360:
	;
	if v1852 == int32(0) {
		goto L359
	} else {
		goto L363
	}
L361:
	;
	v1877 = v1868
	v1878 = v1828
	v1879 = v1834
	goto L362
L362:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1810) {
		v1808 = v1864
		v1810 = v1852
		v1811 = v1865
		v1813 = v1866
		v1818 = v1877
		v1828 = v1878
		v1834 = v1879
		goto L350
	} else {
		goto L364
	}
L363:
	;
	v1873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1834)+1)))
	v1874 = int32(1)
	v1877 = v1874
	v1878 = v1873
	v1879 = v1834 + v1874
	goto L362
L364:
	;
	goto L359
L365:
	;
	goto L343
L366:
	;
	v1925 = v1891 | v1893
	v1926 = int32(1)
	v1927 = v1890 - v1926
	v1929 = v1891 << (uint(v1926) % 32)
	if v1929 == int32(256) {
		goto L368
	} else {
		goto L369
	}
L367:
	;
	v1941 = v1888
	v1946 = v1925
	goto L346
L368:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1888))) = uint8(v1925)
	if v1927 == int32(0) {
		goto L343
	} else {
		goto L371
	}
L369:
	;
	goto L370
L370:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1890) {
		v1890 = v1927
		v1891 = v1929
		v1893 = v1925
		goto L366
	} else {
		goto L372
	}
L371:
	;
	v1935 = int32(1)
	v1936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1888)+1)))
	v1888 = v1888 + v1935
	v1890 = v1927
	v1891 = v1935
	v1893 = v1936
	goto L366
L372:
	;
	goto L367
L373:
	;
	v2060 = v2058 - v1745
	if v2060 != 0 {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	base.MemoryCopy(m, v2032, v1745, v2060)
	goto L376
L375:
	;
	goto L376
L376:
	;
	if v795 != 0 {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v2065 = int32(1) << (uint(v2024&int32(7)) % 32)
	v2067 = base.I32_div_s(v2024, int32(8))
	v2068 = v1174 + v2067
	v2069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2068))))
	if v1057 != 0 {
		goto L381
	} else {
		goto L382
	}
L378:
	;
	goto L379
L379:
	;
	v2090 = int32(1)
	v2091 = v2022 + v2090
	v2095 = v2024 + v2090
	v2097 = v2060 + v2032
	v2099 = F_array_seek(m, v2052, v2022, v1051, v2090, v75, v77)
	mBase = m.M
	v2100 = m.ExcPending
	if v2100 != 0 {
		goto L9
	} else {
		goto L387
	}
L380:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2068))) = uint8(v2085)
	goto L379
L381:
	;
	v2075 = base.I32_div_s(v1751, int32(8))
	v2077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1057+v2075))))
	if int32(base.Ui32(v2077)>>(uint(v1751&int32(7))%32))&int32(1) != 0 {
		goto L384
	} else {
		goto L385
	}
L382:
	;
	goto L383
L383:
	;
	v2085 = v2065 | v2069
	goto L380
L384:
	;
	v2083 = v2065 | v2069
	goto L386
L385:
	;
	v2083 = v2069 & (v2065 ^ int32(-1))
	goto L386
L386:
	;
	v2085 = v2083
	goto L380
L387:
	;
	v2102 = v80 + int32(144)
	v2104 = v80 + int32(208)
	if v90 <= int32(0) {
		goto L389
	} else {
		goto L390
	}
L388:
	;
	if v2158 != int32(-1) {
		v1740 = v2158
		v1741 = v2099
		v1742 = v2095
		v1744 = v2091
		v1745 = v1745 + v2060
		v1750 = v2097
		v1751 = v1751 + v2090
		goto L333
	} else {
		goto L403
	}
L389:
	;
	v2158 = int32(-1)
	goto L388
L390:
	;
	goto L391
L391:
	;
	v2110 = int32(1)
	v2111 = v90 - v2110
	v2113 = v2111 << (uint(int32(2)) % 32)
	v2114 = v2102 + v2113
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v2114)))
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v2104+v2113)))
	v2120 = base.I32_rem_s(v2115+v2110, v2119)
	*(*int32)(unsafe.Add(mBase, uint32(v2114))) = v2120
	if v2111 != 0 {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	v2158 = v2148
	goto L388
L393:
	;
	v2122 = v2111
	v2125 = v2120
	goto L396
L394:
	;
	goto L395
L395:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v2102)))
	if v2146 != 0 {
		goto L400
	} else {
		goto L401
	}
L396:
	;
	if v2125 != 0 {
		v2148 = v2122
		goto L392
	} else {
		goto L398
	}
L397:
	;
	goto L395
L398:
	;
	v2127 = int32(1)
	v2128 = v2122 - v2127
	v2130 = v2128 << (uint(int32(2)) % 32)
	v2131 = v2102 + v2130
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(v2131)))
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(v2104+v2130)))
	v2137 = base.I32_rem_s(v2132+v2127, v2136)
	*(*int32)(unsafe.Add(mBase, uint32(v2131))) = v2137
	if v2128 != 0 {
		v2122 = v2128
		v2125 = v2137
		goto L396
	} else {
		goto L399
	}
L399:
	;
	goto L397
L400:
	;
	v2147 = int32(0)
	goto L402
L401:
	;
	v2147 = int32(-1)
	goto L402
L402:
	;
	v2148 = v2147
	goto L392
L403:
	;
	goto L334
L404:
	;
	v2164 = v2162 - v2099
	if v2164 != 0 {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	base.MemoryCopy(m, v2097, v2099, v2164)
	goto L407
L406:
	;
	goto L407
L407:
	;
	v2166 = int32(0)
	if base.B2i32(v795 == v2166)|base.B2i32(v2161 <= v2166) != 0 {
		goto L213
	} else {
		goto L408
	}
L408:
	;
	v2174 = int32(1) << (uint(v2095&int32(7)) % 32)
	v2176 = base.I32_div_s(v2095, int32(8))
	v2177 = v1174 + v2176
	v2178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2177))))
	if v1051 != 0 {
		goto L215
	} else {
		goto L409
	}
L409:
	;
	v2179 = v2177
	v2181 = v2161
	v2182 = v2174
	v2184 = v2178
	goto L410
L410:
	;
	v2216 = v2182 | v2184
	v2217 = int32(1)
	v2218 = v2181 - v2217
	v2220 = v2182 << (uint(v2217) % 32)
	if v2220 == int32(256) {
		goto L412
	} else {
		goto L413
	}
L411:
	;
	v3006 = v2179
	v3011 = v2216
	goto L214
L412:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2179))) = uint8(v2216)
	if v2218 == int32(0) {
		goto L213
	} else {
		goto L415
	}
L413:
	;
	goto L414
L414:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2181) {
		v2181 = v2218
		v2182 = v2220
		v2184 = v2216
		goto L410
	} else {
		goto L416
	}
L415:
	;
	v2226 = int32(1)
	v2227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2179)+1)))
	v2179 = v2179 + v2226
	v2181 = v2218
	v2182 = v2226
	v2184 = v2227
	goto L410
L416:
	;
	goto L411
L417:
	;
	base.MemoryCopy(m, v2233, v2232, v984)
	goto L419
L418:
	;
	goto L419
L419:
	;
	v2235 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	if v2235 == int32(0) {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v2245 = (v2238<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L422
L421:
	;
	v2245 = v2235
	goto L422
L422:
	;
	v2246 = v2233 + v984
	if v814 != 0 {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	base.MemoryCopy(m, v2246, v2245+v88, v814)
	goto L425
L424:
	;
	goto L425
L425:
	;
	if v998 != 0 {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	base.MemoryCopy(m, v2246+v814, v984+v2232+v985, v998)
	goto L428
L427:
	;
	goto L428
L428:
	;
	if v653&int32(1) == int32(0) {
		goto L213
	} else {
		goto L429
	}
L429:
	;
	v2257 = int32(0)
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v1002)+8))
	if v2259 != 0 {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v2260 = *(*int32)(unsafe.Add(mBase, uint32(v1002)+4))
	v2264 = v1012 + v2260<<(uint(int32(3))%32)
	goto L432
L431:
	;
	v2264 = v2257
	goto L432
L432:
	;
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if v2265 != 0 {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v2266 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v2270 = v230 + v2266<<(uint(int32(3))%32)
	goto L435
L434:
	;
	v2270 = v2257
	goto L435
L435:
	;
	if v992 <= int32(0) {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v2493 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	if v2493 != 0 {
		goto L465
	} else {
		goto L466
	}
L437:
	;
	v2276 = int32(1) << (uint(v663&int32(7)) % 32)
	v2278 = base.I32_div_s(v663, int32(8))
	v2279 = v2264 + v2278
	v2280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2279))))
	if v2270 == int32(0) {
		goto L439
	} else {
		goto L440
	}
L438:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2425))) = uint8(v2429)
	goto L436
L439:
	;
	v2283 = v992
	v2286 = v2276
	v2290 = v2279
	v2294 = v2280
	goto L442
L440:
	;
	goto L441
L441:
	;
	v2336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2270))))
	v2338 = int32(1)
	v2341 = v2276
	v2342 = v992
	v2345 = v2279
	v2349 = v2280
	v2351 = v2270
	v2358 = v2336
	goto L449
L442:
	;
	v2320 = v2286 | v2294
	v2321 = int32(1)
	v2322 = v2283 - v2321
	v2324 = v2286 << (uint(v2321) % 32)
	if v2324 == int32(256) {
		goto L444
	} else {
		goto L445
	}
L443:
	;
	v2425 = v2290
	v2429 = v2320
	goto L438
L444:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2290))) = uint8(v2320)
	if v2322 == int32(0) {
		goto L436
	} else {
		goto L447
	}
L445:
	;
	goto L446
L446:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2283) {
		v2283 = v2322
		v2286 = v2324
		v2294 = v2320
		goto L442
	} else {
		goto L448
	}
L447:
	;
	v2330 = int32(1)
	v2331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2290)+1)))
	v2283 = v2322
	v2286 = v2330
	v2290 = v2290 + v2330
	v2294 = v2331
	goto L442
L448:
	;
	goto L443
L449:
	;
	if v2338&v2358 != 0 {
		goto L451
	} else {
		goto L452
	}
L450:
	;
	if v2394 == int32(1) {
		goto L436
	} else {
		goto L464
	}
L451:
	;
	v2380 = v2341 | v2349
	goto L453
L452:
	;
	v2380 = v2349 & (v2341 ^ int32(-1))
	goto L453
L453:
	;
	v2381 = int32(1)
	v2382 = v2342 - v2381
	v2384 = v2341 << (uint(v2381) % 32)
	if v2384 == int32(256) {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2345))) = uint8(v2380)
	if v2382 == int32(0) {
		goto L436
	} else {
		goto L457
	}
L455:
	;
	v2394 = v2384
	v2395 = v2345
	v2396 = v2380
	goto L456
L456:
	;
	v2398 = v2338 << (uint(int32(1)) % 32)
	if v2398 == int32(256) {
		goto L459
	} else {
		goto L460
	}
L457:
	;
	v2390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2345)+1)))
	v2391 = int32(1)
	v2394 = v2391
	v2395 = v2345 + v2391
	v2396 = v2390
	goto L456
L458:
	;
	goto L450
L459:
	;
	if v2382 == int32(0) {
		goto L458
	} else {
		goto L462
	}
L460:
	;
	v2407 = v2398
	v2408 = v2351
	v2409 = v2358
	goto L461
L461:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2342) {
		v2338 = v2407
		v2341 = v2394
		v2342 = v2382
		v2345 = v2395
		v2349 = v2396
		v2351 = v2408
		v2358 = v2409
		goto L449
	} else {
		goto L463
	}
L462:
	;
	v2403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2351)+1)))
	v2404 = int32(1)
	v2407 = v2404
	v2408 = v2351 + v2404
	v2409 = v2403
	goto L461
L463:
	;
	goto L458
L464:
	;
	v2425 = v2395
	v2429 = v2396
	goto L438
L465:
	;
	v2494 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v2499 = v774 + v2494<<(uint(int32(3))%32)
	goto L467
L466:
	;
	v2499 = int32(0)
	goto L467
L467:
	;
	if v770 <= int32(0) {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	if v986 <= int32(0) {
		goto L213
	} else {
		goto L497
	}
L469:
	;
	v2503 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v2504 = *(*int32)(unsafe.Add(mBase, uint32(v80)+80))
	v2505 = v2503 - v2504
	v2508 = int32(1) << (uint(v2505&int32(7)) % 32)
	v2510 = base.I32_div_s(v2505, int32(8))
	v2511 = v2264 + v2510
	v2512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2511))))
	if v2499 == int32(0) {
		goto L471
	} else {
		goto L472
	}
L470:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2654))) = uint8(v2660)
	goto L468
L471:
	;
	v2517 = v770
	v2518 = v2508
	v2519 = v2511
	v2525 = v2512
	goto L474
L472:
	;
	goto L473
L473:
	;
	v2568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2499))))
	v2572 = v770
	v2573 = v2508
	v2574 = v2511
	v2577 = v2499
	v2580 = v2512
	v2581 = int32(1)
	v2582 = v2568
	goto L481
L474:
	;
	v2552 = v2518 | v2525
	v2553 = int32(1)
	v2554 = v2517 - v2553
	v2556 = v2518 << (uint(v2553) % 32)
	if v2556 == int32(256) {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2519))) = uint8(v2552)
	if v2554 == int32(0) {
		goto L468
	} else {
		goto L479
	}
L477:
	;
	goto L478
L478:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2517) {
		v2517 = v2554
		v2518 = v2556
		v2525 = v2552
		goto L474
	} else {
		goto L480
	}
L479:
	;
	v2562 = int32(1)
	v2563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2519)+1)))
	v2517 = v2554
	v2518 = v2562
	v2519 = v2519 + v2562
	v2525 = v2563
	goto L474
L480:
	;
	v2654 = v2519
	v2660 = v2552
	goto L470
L481:
	;
	if v2581&v2582 != 0 {
		goto L483
	} else {
		goto L484
	}
L482:
	;
	if v2626 == int32(1) {
		goto L468
	} else {
		goto L496
	}
L483:
	;
	v2612 = v2573 | v2580
	goto L485
L484:
	;
	v2612 = v2580 & (v2573 ^ int32(-1))
	goto L485
L485:
	;
	v2613 = int32(1)
	v2614 = v2572 - v2613
	v2616 = v2573 << (uint(v2613) % 32)
	if v2616 == int32(256) {
		goto L486
	} else {
		goto L487
	}
L486:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2574))) = uint8(v2612)
	if v2614 == int32(0) {
		goto L468
	} else {
		goto L489
	}
L487:
	;
	v2626 = v2616
	v2627 = v2574
	v2628 = v2612
	goto L488
L488:
	;
	v2630 = v2581 << (uint(int32(1)) % 32)
	if v2630 == int32(256) {
		goto L491
	} else {
		goto L492
	}
L489:
	;
	v2622 = int32(1)
	v2623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2574)+1)))
	v2626 = v2622
	v2627 = v2574 + v2622
	v2628 = v2623
	goto L488
L490:
	;
	goto L482
L491:
	;
	if v2614 == int32(0) {
		goto L490
	} else {
		goto L494
	}
L492:
	;
	v2639 = v2577
	v2640 = v2630
	v2641 = v2582
	goto L493
L493:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2572) {
		v2572 = v2614
		v2573 = v2626
		v2574 = v2627
		v2577 = v2639
		v2580 = v2628
		v2581 = v2640
		v2582 = v2641
		goto L481
	} else {
		goto L495
	}
L494:
	;
	v2635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2577)+1)))
	v2636 = int32(1)
	v2639 = v2577 + v2636
	v2640 = v2636
	v2641 = v2635
	goto L493
L495:
	;
	goto L490
L496:
	;
	v2654 = v2627
	v2660 = v2628
	goto L470
L497:
	;
	v2728 = v992 + v997
	v2729 = v2728 + v663
	v2732 = int32(1) << (uint(v2729&int32(7)) % 32)
	v2734 = base.I32_div_s(v2729, int32(8))
	v2735 = v2264 + v2734
	v2736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2735))))
	if v2270 == int32(0) {
		goto L499
	} else {
		goto L500
	}
L498:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2880))) = uint8(v2882)
	goto L213
L499:
	;
	v2739 = v2735
	v2741 = v2736
	v2742 = v2732
	v2744 = v986
	goto L502
L500:
	;
	goto L501
L501:
	;
	v2797 = base.I32_div_s(v2728, int32(8))
	v2798 = v2270 + v2797
	v2799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2798))))
	v2800 = v2735
	v2802 = v2736
	v2803 = v2732
	v2805 = v986
	v2807 = v2798
	v2810 = int32(1) << (uint(v2728&int32(7)) % 32)
	v2827 = v2799
	goto L509
L502:
	;
	v2776 = v2741 | v2742
	v2777 = int32(1)
	v2778 = v2744 - v2777
	v2780 = v2742 << (uint(v2777) % 32)
	if v2780 == int32(256) {
		goto L504
	} else {
		goto L505
	}
L504:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2739))) = uint8(v2776)
	if v2778 == int32(0) {
		goto L213
	} else {
		goto L507
	}
L505:
	;
	goto L506
L506:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2744) {
		v2741 = v2776
		v2742 = v2780
		v2744 = v2778
		goto L502
	} else {
		goto L508
	}
L507:
	;
	v2786 = int32(1)
	v2787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2739)+1)))
	v2739 = v2739 + v2786
	v2741 = v2787
	v2742 = v2786
	v2744 = v2778
	goto L502
L508:
	;
	v2880 = v2739
	v2882 = v2776
	goto L498
L509:
	;
	if v2810&v2827 != 0 {
		goto L511
	} else {
		goto L512
	}
L510:
	;
	if v2858 == int32(1) {
		goto L213
	} else {
		goto L524
	}
L511:
	;
	v2842 = v2802 | v2803
	goto L513
L512:
	;
	v2842 = v2802 & (v2803 ^ int32(-1))
	goto L513
L513:
	;
	v2843 = int32(1)
	v2844 = v2805 - v2843
	v2846 = v2803 << (uint(v2843) % 32)
	if v2846 == int32(256) {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2800))) = uint8(v2842)
	if v2844 == int32(0) {
		goto L213
	} else {
		goto L517
	}
L515:
	;
	v2856 = v2800
	v2857 = v2842
	v2858 = v2846
	goto L516
L516:
	;
	v2860 = v2810 << (uint(int32(1)) % 32)
	if v2860 == int32(256) {
		goto L519
	} else {
		goto L520
	}
L517:
	;
	v2852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2800)+1)))
	v2853 = int32(1)
	v2856 = v2800 + v2853
	v2857 = v2852
	v2858 = v2853
	goto L516
L518:
	;
	goto L510
L519:
	;
	if v2844 == int32(0) {
		goto L518
	} else {
		goto L522
	}
L520:
	;
	v2869 = v2807
	v2870 = v2860
	v2871 = v2827
	goto L521
L521:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2805) {
		v2800 = v2856
		v2802 = v2857
		v2803 = v2858
		v2805 = v2844
		v2807 = v2869
		v2810 = v2870
		v2827 = v2871
		goto L509
	} else {
		goto L523
	}
L522:
	;
	v2865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2807)+1)))
	v2866 = int32(1)
	v2869 = v2807 + v2866
	v2870 = v2866
	v2871 = v2865
	goto L521
L523:
	;
	goto L518
L524:
	;
	v2880 = v2856
	v2882 = v2857
	goto L498
L525:
	;
	if v2936&v2953 != 0 {
		goto L527
	} else {
		goto L528
	}
L526:
	;
	if v2983 == int32(1) {
		goto L213
	} else {
		goto L540
	}
L527:
	;
	v2968 = v2929 | v2931
	goto L529
L528:
	;
	v2968 = v2931 & (v2929 ^ int32(-1))
	goto L529
L529:
	;
	v2969 = int32(1)
	v2970 = v2928 - v2969
	v2972 = v2929 << (uint(v2969) % 32)
	if v2972 == int32(256) {
		goto L530
	} else {
		goto L531
	}
L530:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2926))) = uint8(v2968)
	if v2970 == int32(0) {
		goto L213
	} else {
		goto L533
	}
L531:
	;
	v2982 = v2926
	v2983 = v2972
	v2984 = v2968
	goto L532
L532:
	;
	v2986 = v2936 << (uint(int32(1)) % 32)
	if v2986 == int32(256) {
		goto L535
	} else {
		goto L536
	}
L533:
	;
	v2978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2926)+1)))
	v2979 = int32(1)
	v2982 = v2926 + v2979
	v2983 = v2979
	v2984 = v2978
	goto L532
L534:
	;
	goto L526
L535:
	;
	if v2970 == int32(0) {
		goto L534
	} else {
		goto L538
	}
L536:
	;
	v2995 = v2933
	v2996 = v2986
	v2997 = v2953
	goto L537
L537:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2928) {
		v2926 = v2982
		v2928 = v2970
		v2929 = v2983
		v2931 = v2984
		v2933 = v2995
		v2936 = v2996
		v2953 = v2997
		goto L525
	} else {
		goto L539
	}
L538:
	;
	v2991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2933)+1)))
	v2992 = int32(1)
	v2995 = v2933 + v2992
	v2996 = v2992
	v2997 = v2991
	goto L537
L539:
	;
	goto L534
L540:
	;
	v3006 = v2982
	v3011 = v2984
	goto L214
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
	var v14 int64
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
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v6 == int32(1) {
		v9 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+64)) = uint8(v9)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+56)) = int64(0)
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
		v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+6)))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)))
		v22 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16)+9)))
		v25 = F_array_get_element(m, v14, v15, v16+int32(12), v19, v20, v21, v22, v4-int32(-64))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v4)+56)) = v25
			return
		}
	}
}
func F_array_subscript_handler(m *base.Module, l0 int32) int64 {
	return int64(1736264)
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
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
	F_errfinish(m, int32(_a_F_array_subscript_transform_1), int32(96), int32(_a_F_array_subscript_transform_2))
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
	v79 = F_makeConst(m, int32(23), int32(-1), v74, int32(4), int64(1), v74, int32(1))
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
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
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
	F_errfinish(m, int32(_a_F_array_subscript_transform_1), int32(133), int32(_a_F_array_subscript_transform_2))
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
	F_errfinish(m, int32(_a_F_array_subscript_transform_1), int32(153), int32(_a_F_array_subscript_transform_2))
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
func F_array_to_json(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v12 int32
	_ = v12
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
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_initStringInfo(m, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		F_array_to_json_internal(m, v8, v6, int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v18 = F_cstring_to_text_with_len(m, v16, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				m.G0 = v6 + int32(16)
				return base.I64_extend_i32_u(v18)
			}
		}
	}
}
func F_array_to_json_internal(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = F_pg_detoast_datum(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(0)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		v22 = v15 + int32(16)
		v23 = F_ArrayGetNItemsSafe(m, v20, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v23
			if v23 <= int32(0) {
				F_appendStringInfoString(m, l1, int32(_a_F_array_to_json_internal_0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					m.G0 = v12 + int32(32)
					return
				}
			} else {
				F_get_typlenbyvalalign(m, v17, v12+int32(14), v12+int32(13), v12+int32(12))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_json_categorize_type(m, v17, int32(0), v12+int32(8), v12+int32(4))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+14)))
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)))
						v48 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+12)))
						F_deconstruct_array(m, v15, v46, v47, v48, v12+int32(20), v12+int32(16), v12+int32(28))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							F_array_dim_to_json(m, l1, int32(0), v20, v22, v58, v59, v12+int32(24), v62, v63, l2)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
								F_pfree(m, v66)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
									F_pfree(m, v69)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										m.G0 = v12 + int32(32)
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
func F_array_to_sparsevec(m *base.Module, l0 int32) int64 {
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
	var v85 int64
	_ = v85
	var v86 int64
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
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 float64
	_ = v219
	var v221 float32
	_ = v221
	var v224 float64
	_ = v224
	var v229 float64
	_ = v229
	var v234 float64
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v273 float64
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v433 int32
	_ = v433
	var v437 int64
	_ = v437
	var v438 int64
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v520 float64
	_ = v520
	var v521 float32
	_ = v521
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v614 float32
	_ = v614
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v771 = m.ExcPending
	if v771 != 0 {
		goto L9
	} else {
		goto L162
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L9
	} else {
		goto L159
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L9
	} else {
		goto L156
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L9
	} else {
		goto L153
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L9
	} else {
		goto L150
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L9
	} else {
		goto L146
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L9
	} else {
		goto L142
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L9
	} else {
		goto L138
	}
L9:
	;
	return int64(0)
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v654 = m.ExcPending
	if v654 != 0 {
		goto L9
	} else {
		goto L134
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
	F_CheckNnz(m, v369, v371)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L9
	} else {
		goto L74
	}
L24:
	;
	if v54 <= int32(0) {
		goto L61
	} else {
		goto L62
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
		v369 = v64
		v371 = v54
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
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v81+v68<<(uint(int32(3))%32))))
	v86 = F_DirectFunctionCall1Coll(m, int32(1459), int32(0), v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L9
	} else {
		goto L33
	}
L32:
	;
	v369 = v90
	v371 = v93
	goto L23
L33:
	;
	v90 = v67 + base.B2i32(v86 != int64(0))
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
	v369 = int32(0)
	v371 = v54
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
	v115 = v2
	goto L41
L39:
	;
	v154 = v101
	v155 = v2
	goto L40
L40:
	;
	v167 = v154
	v168 = v155
	v171 = int32(0)
	goto L45
L41:
	;
	v122 = v100 + v109<<(uint(int32(3))%32)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v124 = int32(2147483647)
	v126 = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v122)+24))
	v146 = v108 + base.B2i32(v123&v124 != v126) + base.B2i32(v129&v124 != v126) + base.B2i32(v135&v124 != v126) + base.B2i32(v141&v124 != v126)
	v147 = int32(4)
	v148 = v109 + v147
	v150 = v115 + v147
	if v150 != v54&int32(2147483644) {
		v108 = v146
		v109 = v148
		v115 = v150
		goto L41
	} else {
		goto L43
	}
L42:
	;
	if v99 == int32(0) {
		v369 = v146
		v371 = v54
		goto L23
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v154 = v146
	v155 = v148
	goto L40
L45:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v100+v168<<(uint(int32(3))%32))))
	v187 = v167 + base.B2i32(v182&int32(2147483647) != int32(0))
	v188 = int32(1)
	v191 = v171 + v188
	if v191 != v99 {
		v167 = v187
		v168 = v168 + v188
		v171 = v191
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v369 = v187
	v371 = v54
	goto L23
L47:
	;
	goto L46
L48:
	;
	v369 = int32(0)
	v371 = v54
	goto L23
L49:
	;
	goto L50
L50:
	;
	v197 = v54 & int32(3)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v199 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v54) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v204 = v199
	v205 = v2
	v213 = v2
	goto L54
L52:
	;
	v246 = v199
	v247 = v2
	goto L53
L53:
	;
	v258 = v246
	v259 = v247
	v268 = v2
	goto L58
L54:
	;
	v218 = v198 + v205<<(uint(int32(3))%32)
	v219 = *(*float64)(unsafe.Add(mBase, uint32(v218)))
	v221 = float32(0)
	v224 = *(*float64)(unsafe.Add(mBase, uint32(v218)+8))
	v229 = *(*float64)(unsafe.Add(mBase, uint32(v218)+16))
	v234 = *(*float64)(unsafe.Add(mBase, uint32(v218)+24))
	v238 = v204 + base.F32_ne(base.F32_demote_f64(v219), v221) + base.F32_ne(base.F32_demote_f64(v224), v221) + base.F32_ne(base.F32_demote_f64(v229), v221) + base.F32_ne(base.F32_demote_f64(v234), v221)
	v239 = int32(4)
	v240 = v205 + v239
	v242 = v213 + v239
	if v242 != v54&int32(2147483644) {
		v204 = v238
		v205 = v240
		v213 = v242
		goto L54
	} else {
		goto L56
	}
L55:
	;
	if v197 == int32(0) {
		v369 = v238
		v371 = v54
		goto L23
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v246 = v238
	v247 = v240
	goto L53
L58:
	;
	v273 = *(*float64)(unsafe.Add(mBase, uint32(v198+v259<<(uint(int32(3))%32))))
	v277 = v258 + base.F32_ne(base.F32_demote_f64(v273), float32(0))
	v278 = int32(1)
	v281 = v268 + v278
	if v281 != v197 {
		v258 = v277
		v259 = v259 + v278
		v268 = v281
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v369 = v277
	v371 = v54
	goto L23
L60:
	;
	goto L59
L61:
	;
	v369 = int32(0)
	v371 = v54
	goto L23
L62:
	;
	goto L63
L63:
	;
	v287 = v54 & int32(3)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v289 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v54) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v294 = v289
	v295 = v2
	v301 = v2
	goto L67
L65:
	;
	v332 = v289
	v333 = v2
	goto L66
L66:
	;
	v345 = v332
	v346 = v333
	v349 = int32(0)
	goto L71
L67:
	;
	v308 = v288 + v295<<(uint(int32(3))%32)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
	v310 = int32(0)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v308)+8))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v308)+16))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v308)+24))
	v324 = v294 + base.B2i32(v309 != v310) + base.B2i32(v313 != v310) + base.B2i32(v317 != v310) + base.B2i32(v321 != v310)
	v325 = int32(4)
	v326 = v295 + v325
	v328 = v301 + v325
	if v328 != v54&int32(2147483644) {
		v294 = v324
		v295 = v326
		v301 = v328
		goto L67
	} else {
		goto L69
	}
L68:
	;
	if v287 == int32(0) {
		v369 = v324
		v371 = v54
		goto L23
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	v332 = v324
	v333 = v326
	goto L66
L71:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v288+v346<<(uint(int32(3))%32))))
	v363 = v345 + base.B2i32(v360 != int32(0))
	v364 = int32(1)
	v367 = v349 + v364
	if v367 != v287 {
		v345 = v363
		v346 = v346 + v364
		v349 = v367
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v369 = v363
	v371 = v54
	goto L23
L73:
	;
	goto L72
L74:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v386 = F_mul_size(m, int32(4), v369)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L9
	} else {
		goto L75
	}
L75:
	;
	v388 = F_add_size(m, int32(16), v386)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	v391 = F_mul_size(m, int32(4), v369)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	v393 = F_add_size(m, v388, v391)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	v395 = F_palloc0(m, v393)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L9
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395)+8)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v395)+4)) = v383
	v399 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v395))) = v393 << (uint(v399) % 32)
	v403 = v395 + int32(16)
	v406 = v403 + v369<<(uint(v399)%32)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	switch v407 - int32(700) {
	case 0:
		goto L83
	case 1:
		goto L82
	default:
		goto L84
	}
L80:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	F_pfree(m, v591)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L9
	} else {
		goto L120
	}
L81:
	;
	v540 = int32(0)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v541 <= v540 {
		v580 = v540
		goto L80
	} else {
		goto L112
	}
L82:
	;
	v499 = int32(0)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v500 <= v499 {
		v580 = v499
		goto L80
	} else {
		goto L104
	}
L83:
	;
	v459 = int32(0)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v460 <= v459 {
		v580 = v459
		goto L80
	} else {
		goto L96
	}
L84:
	;
	if v407 == int32(23) {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	if v407 != int32(1700) {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v414 = int32(0)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v415 <= v414 {
		v580 = v414
		goto L80
	} else {
		goto L87
	}
L87:
	;
	v419 = int32(0)
	v420 = v414
	goto L88
L88:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v437 = *(*int64)(unsafe.Add(mBase, uint32(v433+v419<<(uint(int32(3))%32))))
	v438 = F_DirectFunctionCall1Coll(m, int32(1459), int32(0), v437)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L9
	} else {
		goto L90
	}
L89:
	;
	v580 = v453
	goto L80
L90:
	;
	v440 = base.I32_wrap_i64(v438)
	if v440&int32(2147483647) != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v395)+8))
	if v443 <= v420 {
		goto L2
	} else {
		goto L94
	}
L92:
	;
	v453 = v420
	goto L93
L93:
	;
	v456 = v419 + int32(1)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v456 < v457 {
		v419 = v456
		v420 = v453
		goto L88
	} else {
		goto L95
	}
L94:
	;
	v446 = v420 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v403+v446))) = v419
	*(*int32)(unsafe.Add(mBase, uint32(v446+v406))) = v440
	v453 = v420 + int32(1)
	goto L93
L95:
	;
	goto L89
L96:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v465 = int32(0)
	v466 = v459
	v467 = v460
	goto L97
L97:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v463+v465<<(uint(int32(3))%32))))
	if v480&int32(2147483647) != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v580 = v494
	goto L80
L99:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v395)+8))
	if v483 <= v466 {
		goto L3
	} else {
		goto L102
	}
L100:
	;
	v494 = v466
	v495 = v467
	goto L101
L101:
	;
	v497 = v465 + int32(1)
	if v497 < v495 {
		v465 = v497
		v466 = v494
		v467 = v495
		goto L97
	} else {
		goto L103
	}
L102:
	;
	v486 = v466 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v403+v486))) = v465
	*(*int32)(unsafe.Add(mBase, uint32(v486+v406))) = v480
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v494 = v466 + int32(1)
	v495 = v491
	goto L101
L103:
	;
	goto L98
L104:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v505 = int32(0)
	v506 = v499
	v507 = v500
	goto L105
L105:
	;
	v520 = *(*float64)(unsafe.Add(mBase, uint32(v503+v505<<(uint(int32(3))%32))))
	v521 = base.F32_demote_f64(v520)
	if base.F32_ne(v521, float32(0)) != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v580 = v535
	goto L80
L107:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v395)+8))
	if v524 <= v506 {
		goto L4
	} else {
		goto L110
	}
L108:
	;
	v535 = v506
	v536 = v507
	goto L109
L109:
	;
	v538 = v505 + int32(1)
	if v538 < v536 {
		v505 = v538
		v506 = v535
		v507 = v536
		goto L105
	} else {
		goto L111
	}
L110:
	;
	v527 = v506 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v403+v527))) = v505
	*(*float32)(unsafe.Add(mBase, uint32(v527+v406))) = v521
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v535 = v506 + int32(1)
	v536 = v532
	goto L109
L111:
	;
	goto L106
L112:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v546 = int32(0)
	v547 = v540
	v548 = v541
	goto L113
L113:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v544+v546<<(uint(int32(3))%32))))
	if v561 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v580 = v574
	goto L80
L115:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v395)+8))
	if v562 <= v547 {
		goto L5
	} else {
		goto L118
	}
L116:
	;
	v574 = v547
	v575 = v548
	goto L117
L117:
	;
	v577 = v546 + int32(1)
	if v577 < v575 {
		v546 = v577
		v547 = v574
		v548 = v575
		goto L113
	} else {
		goto L119
	}
L118:
	;
	v565 = v547 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v403+v565))) = v546
	*(*float32)(unsafe.Add(mBase, uint32(v565+v406))) = base.F32_convert_i32_s(v561)
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v574 = v547 + int32(1)
	v575 = v571
	goto L117
L119:
	;
	goto L114
L120:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v395)+8))
	if v594 == v580 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v596 = int32(0)
	if v596 < v580 {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L9
	} else {
		goto L131
	}
L124:
	;
	v599 = v596
	goto L127
L125:
	;
	goto L126
L126:
	;
	m.G0 = v15 + int32(32)
	return base.I64_extend_i32_u(v395)
L127:
	;
	v614 = *(*float32)(unsafe.Add(mBase, uint32(v406+v599<<(uint(int32(2))%32))))
	F_CheckElement_2(m, v614)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L9
	} else {
		goto L129
	}
L128:
	;
	goto L126
L129:
	;
	v618 = v599 + int32(1)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v395)+8))
	if v618 < v619 {
		v599 = v618
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	F_errmsg_internal(m, int32(_a_F_array_to_sparsevec_0), int32(0))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L9
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(810), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L9
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L9
	} else {
		goto L135
	}
L135:
	;
	F_errmsg(m, int32(_a_F_array_to_sparsevec_3), int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L9
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(709), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L9
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L9
	} else {
		goto L139
	}
L139:
	;
	F_errmsg(m, int32(_a_F_array_to_sparsevec_4), int32(0))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L9
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(714), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L9
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L9
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v25
	F_errmsg(m, int32(_a_F_array_to_sparsevec_5), v15)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L9
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(62), int32(_a_F_array_to_sparsevec_6))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L9
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L9
	} else {
		goto L147
	}
L147:
	;
	F_errmsg(m, int32(_a_F_array_to_sparsevec_7), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L9
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(753), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
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
	v723 = m.ExcPending
	if v723 != 0 {
		goto L9
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(776), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
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
	v736 = m.ExcPending
	if v736 != 0 {
		goto L9
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(781), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
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
	v749 = m.ExcPending
	if v749 != 0 {
		goto L9
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(786), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_array_to_sparsevec_8), int32(0))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L9
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(791), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L9
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L9
	} else {
		goto L163
	}
L163:
	;
	F_errmsg(m, int32(_a_F_array_to_sparsevec_7), int32(0))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L9
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_array_to_sparsevec_1), int32(797), int32(_a_F_array_to_sparsevec_2))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L9
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_upper(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
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
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_DatumGetAnyArrayP(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
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
			return int64(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v28 = int32(0)
			if base.B2i32(v28 < v27)&base.B2i32(base.Ui32(v27) <= base.Ui32(v18)) == v28 {
				v34 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
				return int64(0)
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
				return base.I64_extend_i32_s(v55 + v57 - int32(1))
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
func F_makeArrayResultAny(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11 != 0 {
		v12 = int32(_a_F_makeArrayResultAny_0)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_makeArrayResultAny[0]))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		*(*int32)(unsafe.Add(mBase, _c_F_makeArrayResultAny[0])) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v14
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
		v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+24)))
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)))
		v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+27)))
		v32 = F_construct_md_array(m, v20, v21, base.B2i32(int32(0) < v14), v9+int32(12), v9+int32(8), v28, v29, v30, v31)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_makeArrayResultAny[0])) = v13
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			F_MemoryContextDelete(m, v38)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int64(0)
			} else {
				v48 = base.I64_extend_i32_u(v32)
				m.G0 = v9 + int32(16)
				return v48
			}
		}
	} else {
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v44 = F_makeArrayResultArr(m, v42, l1, int32(1))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			v48 = v44
			m.G0 = v9 + int32(16)
			return v48
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
	F_errfinish(m, int32(_a_F_transformArrayExpr_2), int32(2122), int32(_a_F_transformArrayExpr_3))
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
	F_errfinish(m, int32(_a_F_transformArrayExpr_2), int32(2133), int32(_a_F_transformArrayExpr_3))
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
	F_errfinish(m, int32(_a_F_transformArrayExpr_2), int32(2168), int32(_a_F_transformArrayExpr_3))
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
	F_errfinish(m, int32(_a_F_transformArrayExpr_2), int32(2108), int32(_a_F_transformArrayExpr_3))
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
func F_trim_array(m *base.Module, l0 int32) int64 {
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		if int32(0) < v16 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v20 = v19
		} else {
			v20 = int32(0)
		}
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v22 = int32(0)
		if base.B2i32(v21 < v22)|base.B2i32(v20 < v21) == v22 {
			v28 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v9)+28)) = uint16(v28)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v28
			*(*uint16)(unsafe.Add(mBase, uint32(v9)+20)) = uint16(v28)
			if v28 < v16 {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v12+v16<<(uint(int32(2))%32))+16))
				v42 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)) = uint8(v42)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v41 + (v20 + (v21 ^ int32(-1)))
			} else {
			}
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
			F_get_typlenbyvalalign(m, v50, v9+int32(94), v9+int32(93), v9+int32(92))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int64(0)
			} else {
				v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+94)))
				v71 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9)+92)))
				v72 = F_array_get_slice(m, base.I64_extend_i32_u(v12), int32(1), v9+int32(32), v9-int32(-64), v9+int32(16), v9+int32(24), int32(-1), v70, v71)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int64(0)
				} else {
					m.G0 = v9 + int32(96)
					return v72
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v20
					F_errmsg(m, int32(_a_F_trim_array_0), v9)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_trim_array_1), int32(_a_F_trim_array_2), int32(_a_F_trim_array_3))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int64(0)
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
