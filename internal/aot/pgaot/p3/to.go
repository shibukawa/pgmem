package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyToBinaryStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v7, int32(1592480), int32(11))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		F_appendBinaryStringInfo(m, v14, v5+int32(8), int32(4))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(0)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_appendBinaryStringInfo(m, v22, v5+int32(12), int32(4))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				m.G0 = v5 + int32(16)
				return
			}
		}
	}
}
func F_CopyToTextOneRow(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v243 int32
	_ = v243
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_CopySendTextLikeEndOfRow(m, l0)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L62
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v15 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = v20 - int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v23))))
	if v25 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v102 = int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v103 <= v102 {
		goto L1
	} else {
		goto L29
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v22<<(uint(int32(2))%32))))
	v36 = F_OutputFunctionCall(m, v18+v22*int32(28), v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v41&int32(3) == int32(0) {
		v65 = v41
		goto L13
	} else {
		goto L14
	}
L8:
	;
	return
L9:
	;
	F_CopyAttributeOutText(m, l0, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	F_appendBinaryStringInfo(m, v40, v41, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L8
	} else {
		goto L28
	}
L12:
	;
	v98 = v90 - v41
	goto L11
L13:
	;
	v69 = v65
	goto L22
L14:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v49 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v98 = int32(0)
	goto L11
L16:
	;
	goto L17
L17:
	;
	v54 = v41
	goto L18
L18:
	;
	v58 = v54 + int32(1)
	if v58&int32(3) == int32(0) {
		v65 = v58
		goto L13
	} else {
		goto L20
	}
L19:
	;
	v90 = v58
	goto L12
L20:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v63 != 0 {
		v54 = v58
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v78 = int32(-2139062144)
	if (int32(16843008)-v75|v75)&v78 == v78 {
		v69 = v69 + int32(4)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v84 = v69
	goto L25
L24:
	;
	goto L23
L25:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v88 != 0 {
		v84 = v84 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v90 = v84
	goto L12
L27:
	;
	goto L26
L28:
	;
	goto L4
L29:
	;
	v108 = v102
	goto L30
L30:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v118 = int32(2)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117+v108<<(uint(v118)%32))))
	v122 = int32(1)
	v123 = v121 - v122
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+v124))))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v123<<(uint(v118)%32))))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v133 = int32(*(*int8)(unsafe.Add(mBase, uint32(v132))))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	if v138 <= v135+v122 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L1
L32:
	;
	if v126&int32(1) != 0 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	F_appendStringInfoChar(m, v134, v133)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L8
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	*(*uint8)(unsafe.Add(mBase, uint32(v142+v135))) = uint8(v133)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	v148 = v146 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+4)) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v150+v148))) = uint8(v152)
	goto L32
L36:
	;
	goto L32
L37:
	;
	v228 = v108 + int32(1)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v228 < v229 {
		v108 = v228
		goto L30
	} else {
		goto L61
	}
L38:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v159&int32(3) == int32(0) {
		v183 = v159
		goto L43
	} else {
		goto L44
	}
L39:
	;
	goto L40
L40:
	;
	v222 = F_OutputFunctionCall(m, v18+v123*int32(28), v131)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L8
	} else {
		goto L59
	}
L41:
	;
	F_appendBinaryStringInfo(m, v158, v159, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L8
	} else {
		goto L58
	}
L42:
	;
	v216 = v208 - v159
	goto L41
L43:
	;
	v187 = v183
	goto L52
L44:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v167 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v216 = int32(0)
	goto L41
L46:
	;
	goto L47
L47:
	;
	v172 = v159
	goto L48
L48:
	;
	v176 = v172 + int32(1)
	if v176&int32(3) == int32(0) {
		v183 = v176
		goto L43
	} else {
		goto L50
	}
L49:
	;
	v208 = v176
	goto L42
L50:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	if v181 != 0 {
		v172 = v176
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v196 = int32(-2139062144)
	if (int32(16843008)-v193|v193)&v196 == v196 {
		v187 = v187 + int32(4)
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v202 = v187
	goto L55
L54:
	;
	goto L53
L55:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v206 != 0 {
		v202 = v202 + int32(1)
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v208 = v202
	goto L42
L57:
	;
	goto L56
L58:
	;
	goto L37
L59:
	;
	F_CopyAttributeOutText(m, l0, v222)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	goto L37
L61:
	;
	goto L31
L62:
	;
	return
}
func F_to_tsquery(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v9 = F_getTSCurrentConfig(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_DirectFunctionCall2Coll(m, int32(1174), int32(0), v9, v3)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
