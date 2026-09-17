package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_appendStringInfoSpaces(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	if int32(0) < l1 {
		F_enlargeStringInfo(m, l0, l1)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			if l1 != 0 {
				v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				base.MemoryFill(m, v7+v8, int32(32), l1)
			} else {
			}
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v13 = v12 + l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v13
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v17 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v15+v13))) = uint8(v17)
			return
		}
	} else {
		return
	}
}
func F_appendStringInfoString(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v4 = F_strlen(m, l1)
	mBase = m.M
	F_enlargeStringInfo(m, l0, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if v4 != 0 {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			base.MemoryCopy(m, v7+v8, l1, v4)
		} else {
		}
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v12 = v11 + v4
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14+v12))) = uint8(v16)
		return
	}
}
func F_appendStringInfoVA(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = v8 - v9
	if int32(16) <= v10 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = F_pvsnprintf(m, v13+v9, v10, l1, l2)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if base.Ui32(v15) < base.Ui32(v10) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 + v15
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v28 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v25+v26))) = uint8(v28)
				v30 = v15
				return v30
			}
		}
	} else {
		v30 = int32(32)
		return v30
	}
}
func F_makeStringConst(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_palloc0(m, int32(20))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = int64(2010044694600)
		return v5
	}
}
func F_stringToNode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v2 = int32(0)
	v3 = int32(_a_F_stringToNode_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_stringToNode[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_stringToNode[0])) = l0
	v9 = F_nodeRead(m, v2, v2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_stringToNode[0])) = v4
		return v9
	}
}
func F_string_field_used(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v4 = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if l1 == v6 {
		v28 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v28
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if l1 == v8 {
		v28 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if l1 == v10 {
		v28 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v14 = l0 + int32(56)
	goto L5
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v18 = int32(0)
	v19 = base.B2i32(v17 != v18)
	if v17 == v18 {
		v28 = v19
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v28 = v19
	goto L1
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if l1 == v22 {
		v28 = v19
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	if l1 != v24 {
		v14 = v17
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L6
}
func F_string_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v288 int32
	_ = v288
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v9 = l2 + int32(256)
	if base.B2i32(v6&int32(3) == v4)|base.B2i32(v9 == v4) != 0 {
		v40 = v6
		v42 = v9
		v43 = base.B2i32(v9 != v4)
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if v114 != 0 {
		goto L26
	} else {
		goto L27
	}
L2:
	;
	v114 = int32(0)
	goto L1
L3:
	;
	v92 = v85
	v94 = v87
	goto L20
L4:
	;
	if v43 == int32(0) {
		goto L2
	} else {
		goto L11
	}
L5:
	;
	v23 = v6
	v25 = v9
	goto L6
L6:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v28 == int32(0) {
		v85 = v23
		v87 = v25
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v40 = v35
	v42 = v31
	v43 = v33
	goto L4
L8:
	;
	v30 = int32(1)
	v31 = v25 - v30
	v32 = int32(0)
	v33 = base.B2i32(v31 != v32)
	v35 = v23 + v30
	if v35&int32(3) == v32 {
		v40 = v35
		v42 = v31
		v43 = v33
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v31 != 0 {
		v23 = v35
		v25 = v31
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v48 = int32(0)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if base.B2i32(v48 == v49)|base.B2i32(base.Ui32(v42) < base.Ui32(int32(4))) == v48 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v58 = v40
	v60 = v42
	goto L15
L13:
	;
	v78 = v40
	v80 = v42
	goto L14
L14:
	;
	if v80 == int32(0) {
		goto L2
	} else {
		goto L19
	}
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v65 = v64 ^ int32(0)
	v68 = int32(-2139062144)
	if (int32(16843008)-v65|v65)&v68 != v68 {
		v85 = v58
		v87 = v60
		goto L3
	} else {
		goto L17
	}
L16:
	;
	v78 = v73
	v80 = v75
	goto L14
L17:
	;
	v72 = int32(4)
	v73 = v58 + v72
	v75 = v60 - v72
	if base.Ui32(int32(3)) < base.Ui32(v75) {
		v58 = v73
		v60 = v75
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v85 = v78
	v87 = v80
	goto L3
L20:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if int32(0) == v97 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L2
L22:
	;
	v114 = v92
	goto L1
L23:
	;
	goto L24
L24:
	;
	v99 = int32(1)
	v102 = v94 - v99
	if v102 != 0 {
		v92 = v92 + v99
		v94 = v102
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v116 = v114 - v6
	goto L28
L27:
	;
	v116 = v9
	goto L28
L28:
	;
	if base.Ui32(v116) < base.Ui32(l2) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v118 = v116
	goto L31
L30:
	;
	v118 = l2
	goto L31
L31:
	;
	if base.Ui32(int32(512)) <= base.Ui32(v118) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v288 = v6 + v116
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 + v118
	return v118
L33:
	;
	if v118 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v125 = l1 + v118
	if (l1^v6)&int32(3) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	base.MemoryCopy(m, l1, v6, v118)
	goto L38
L37:
	;
	goto L38
L38:
	;
	goto L32
L39:
	;
	if base.Ui32(v257) < base.Ui32(v125) {
		goto L73
	} else {
		goto L74
	}
L40:
	;
	if l1&int32(3) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	goto L42
L42:
	;
	if base.Ui32(v125) < base.Ui32(int32(4)) {
		goto L64
	} else {
		goto L65
	}
L43:
	;
	v161 = v125 & int32(-4)
	if base.Ui32(v125) < base.Ui32(int32(64)) {
		v211 = v155
		v212 = v156
		goto L54
	} else {
		goto L55
	}
L44:
	;
	v155 = v6
	v156 = l1
	goto L43
L45:
	;
	goto L46
L46:
	;
	if v118 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v155 = v6
	v156 = l1
	goto L43
L48:
	;
	goto L49
L49:
	;
	v138 = v6
	v139 = l1
	goto L50
L50:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	*(*uint8)(unsafe.Add(mBase, uint32(v139))) = uint8(v143)
	v145 = int32(1)
	v146 = v138 + v145
	v148 = v139 + v145
	if v148&int32(3) == int32(0) {
		v155 = v146
		v156 = v148
		goto L43
	} else {
		goto L52
	}
L51:
	;
	v155 = v146
	v156 = v148
	goto L43
L52:
	;
	if base.Ui32(v148) < base.Ui32(v125) {
		v138 = v146
		v139 = v148
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	if base.Ui32(v161) <= base.Ui32(v212) {
		v256 = v211
		v257 = v212
		goto L39
	} else {
		goto L60
	}
L55:
	;
	v165 = v161 + int32(-64)
	if base.Ui32(v165) < base.Ui32(v156) {
		v211 = v155
		v212 = v156
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v168 = v155
	v169 = v156
	goto L57
L57:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v173
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+4)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v168)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+8)) = v177
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+12)) = v179
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+16)) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v168)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+20)) = v183
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v168)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+24)) = v185
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v168)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+28)) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v168)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+32)) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v168)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+36)) = v191
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v168)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+40)) = v193
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v168)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+44)) = v195
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v168)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+48)) = v197
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v168)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+52)) = v199
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v168)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+56)) = v201
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v168)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+60)) = v203
	v205 = int32(-64)
	v206 = v168 - v205
	v208 = v169 - v205
	if base.Ui32(v208) <= base.Ui32(v165) {
		v168 = v206
		v169 = v208
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v211 = v206
	v212 = v208
	goto L54
L59:
	;
	goto L58
L60:
	;
	v218 = v211
	v219 = v212
	goto L61
L61:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = v223
	v225 = int32(4)
	v226 = v218 + v225
	v228 = v219 + v225
	if base.Ui32(v228) < base.Ui32(v161) {
		v218 = v226
		v219 = v228
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v256 = v226
	v257 = v228
	goto L39
L63:
	;
	goto L62
L64:
	;
	v256 = v6
	v257 = l1
	goto L39
L65:
	;
	goto L66
L66:
	;
	if base.Ui32(v118) < base.Ui32(int32(4)) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v256 = v6
	v257 = l1
	goto L39
L68:
	;
	goto L69
L69:
	;
	v237 = v6
	v238 = l1
	goto L70
L70:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	*(*uint8)(unsafe.Add(mBase, uint32(v238))) = uint8(v242)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v238)+1)) = uint8(v244)
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v238)+2)) = uint8(v246)
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v238)+3)) = uint8(v248)
	v250 = int32(4)
	v251 = v237 + v250
	v253 = v238 + v250
	if base.Ui32(v253) <= base.Ui32(v125-int32(4)) {
		v237 = v251
		v238 = v253
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v256 = v251
	v257 = v253
	goto L39
L72:
	;
	goto L71
L73:
	;
	v263 = v256
	v264 = v257
	goto L76
L74:
	;
	goto L75
L75:
	;
	goto L32
L76:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	*(*uint8)(unsafe.Add(mBase, uint32(v264))) = uint8(v268)
	v270 = int32(1)
	v273 = v264 + v270
	if v273 != v125 {
		v263 = v263 + v270
		v264 = v273
		goto L76
	} else {
		goto L78
	}
L77:
	;
	goto L75
L78:
	;
	goto L77
}
