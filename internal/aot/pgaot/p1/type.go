package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TypeNameListToString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_initStringInfo(m, v7)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l0 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	m.G0 = v7 + int32(16)
	return v42
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v19 = int32(0)
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v19<<(uint(int32(2))%32))))
	if v19&int32(1073741823) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L3
L8:
	;
	F_appendStringInfoChar(m, v7, int32(44))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_appendTypeNameToBuffer(m, v26, v7)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	v35 = v19 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v35 < v36 {
		v19 = v35
		goto L6
	} else {
		goto L13
	}
L13:
	;
	goto L7
}
func F_TypeNameToString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_initStringInfo(m, v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		F_appendTypeNameToBuffer(m, l0, v5)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			m.G0 = v5 + int32(16)
			return v13
		}
	}
}
func F_record_type_typmod_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v9 = F_hash_bytes_uint32(m, v8)
	mBase = m.M
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v11 = F_hash_bytes_uint32(m, v10)
	mBase = m.M
	v12 = int32(1640531527)
	v13 = v9 - v12
	v22 = v11 + v13<<(uint(int32(6))%32) + int32(base.Ui32(v13)>>(uint(int32(2))%32)) - v12 ^ v13
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	if v4 < v23 {
		v29 = v22
		v30 = v4
		v31 = v23
		for {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v3+int32(88)+v31<<(uint(int32(4))%32)+v30*int32(100))))
			v40 = F_hash_bytes_uint32(m, v39)
			mBase = m.M
			v49 = v40 + (v29<<(uint(int32(6))%32) + int32(base.Ui32(v29)>>(uint(int32(2))%32))) - int32(1640531527) ^ v29
			v51 = v30 + int32(1)
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
			if v51 < v52 {
				v29 = v49
				v30 = v51
				v31 = v52
				continue
			} else {
				break
			}
			break
		}
		v55 = v49
	} else {
		v55 = v22
	}
	return v55
}
func F_typeStringToTypeName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(717150)
	v13 = m.G0
	v15 = v13 - int32(32)
	v16 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v16
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _consts[339])))
	if v24 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v192
L2:
	;
	if l0&int32(3) == int32(0) {
		v116 = l0
		goto L25
	} else {
		goto L26
	}
L3:
	;
	v92 = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[340])))
	if v28 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v32 = l0
	goto L9
L7:
	;
	goto L8
L8:
	;
	v42 = v9
	v43 = v24
	goto L12
L9:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v38 == v24 {
		v32 = v32 + int32(1)
		goto L9
	} else {
		goto L11
	}
L10:
	;
	v92 = v32 - l0
	goto L2
L11:
	;
	goto L10
L12:
	;
	v50 = v15 + int32(base.Ui32(v43)>>(uint(int32(3))%32))&int32(28)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v52 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v51 | v52<<(uint(v43)%32)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v56 != 0 {
		v42 = v42 + v52
		v43 = v56
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v59 == int32(0) {
		v84 = l0
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v92 = v84 - l0
	goto L2
L16:
	;
	v63 = l0
	v64 = v59
	goto L17
L17:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(base.Ui32(v64)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v72)>>(uint(v64)%32))&int32(1) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v84 = v80
	goto L15
L19:
	;
	v84 = v63
	goto L15
L20:
	;
	goto L21
L21:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	v80 = v63 + int32(1)
	if v78 != 0 {
		v63 = v80
		v64 = v78
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	if v92 != v149 {
		goto L40
	} else {
		goto L41
	}
L24:
	;
	v149 = v141 - l0
	goto L23
L25:
	;
	v120 = v116
	goto L34
L26:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v100 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v149 = int32(0)
	goto L23
L28:
	;
	goto L29
L29:
	;
	v105 = l0
	goto L30
L30:
	;
	v109 = v105 + int32(1)
	if v109&int32(3) == int32(0) {
		v116 = v109
		goto L25
	} else {
		goto L32
	}
L31:
	;
	v141 = v109
	goto L24
L32:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v114 != 0 {
		v105 = v109
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v129 = int32(-2139062144)
	if (int32(16843008)-v126|v126)&v129 == v129 {
		v120 = v120 + int32(4)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v135 = v120
	goto L37
L36:
	;
	goto L35
L37:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v139 != 0 {
		v135 = v135 + int32(1)
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v141 = v135
	goto L24
L39:
	;
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(498)
	v154 = int32(4457400)
	v155 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v7 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v155
	v162 = F_raw_parser(m, l0, int32(1))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v175 = int32(0)
	v176 = F_errsave_start(m, l1)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L43
	} else {
		goto L46
	}
L43:
	;
	return int32(0)
L44:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+12)))
	if v171 != int32(1) {
		v192 = v170
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	if v176 == int32(0) {
		v192 = v175
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L43
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg(m, int32(685904), v7)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L43
	} else {
		goto L49
	}
L49:
	;
	F_errsave_finish(m, l1, int32(487596), int32(773), int32(373465))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L43
	} else {
		goto L50
	}
L50:
	;
	v192 = v175
	goto L1
}
func F_typeTypeId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	if l0 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(105984), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(487596), int32(593), int32(455256))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v19+v20)))
		return v22
	}
}
func F_type_is_collatable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+144))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v15 != int32(0))
			}
		}
	}
}
