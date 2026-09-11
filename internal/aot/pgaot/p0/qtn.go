package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_QTNBinary(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	F_check_stack_depth(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v7 != int32(2) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v10 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v14 = int32(0)
	goto L6
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+v14<<(uint(int32(2))%32))))
	F_QTNBinary(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	if v25 < int32(3) {
		goto L3
	} else {
		goto L10
	}
L8:
	;
	v24 = v14 + int32(1)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v24 < v25 {
		v14 = v24
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	goto L11
L11:
	;
	v33 = F_palloc0(m, int32(24))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L3
L13:
	;
	v36 = F_palloc0(m, int32(12))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v36
	v40 = F_palloc0(m, int32(8))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v33)+4)) = int64(8589934593)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v54 | v56
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v61)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)) = uint8(v65)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v33
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v71 = int32(2)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69+v70<<(uint(v71)%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v80 = v78 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v80
	if v71 < v80 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
}
func F_QTNTernary(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	F_check_stack_depth(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v11 != int32(2) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(0) < v14 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v19 = int32(0)
	goto L8
L6:
	;
	v41 = v14
	v44 = v10
	goto L7
L7:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v45&int32(254) != int32(2) {
		goto L3
	} else {
		goto L12
	}
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v19<<(uint(int32(2))%32))))
	F_QTNTernary(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = v34
	v44 = v36
	goto L7
L10:
	;
	v33 = v19 + int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v33 < v34 {
		v19 = v33
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	if v41 <= int32(0) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v54 = int32(0)
	v57 = v41
	goto L14
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v61 = int32(2)
	v62 = v54 << (uint(v61) % 32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v62)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v66 != v61 {
		v265 = v54
		v267 = v57
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L3
L16:
	;
	v269 = v265 + int32(1)
	if v269 < v267 {
		v54 = v269
		v57 = v267
		goto L14
	} else {
		goto L78
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	if v70 != v71 {
		v265 = v54
		v267 = v57
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v76 = v57 + v73 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v76
	v80 = F_repalloc(m, v60, v76<<(uint(int32(2))%32))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v80
	if v57 != v54+int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v86 = v80 + v62
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v88 = int32(2)
	v90 = v86 + v87<<(uint(v88)%32)
	v92 = v86 + int32(4)
	v97 = (v57 + (v54 ^ int32(-1))) << (uint(v88) % 32)
	if v90 == v92 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v244 = v80
	goto L22
L22:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v249 = v247 << (uint(int32(2)) % 32)
	if v249 != 0 {
		goto L70
	} else {
		goto L71
	}
L23:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v244 = v242
	goto L22
L24:
	;
	goto L23
L25:
	;
	v101 = v90 + v97
	if base.Ui32(v92-v101) <= base.Ui32(int32(0)-v97<<(uint(int32(1))%32)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v108 = F___memcpy(m, v90, v92, v97)
	mBase = m.M
	goto L23
L27:
	;
	goto L28
L28:
	;
	v111 = (v90 ^ v92) & int32(3)
	if base.Ui32(v90) < base.Ui32(v92) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if v213 == int32(0) {
		goto L24
	} else {
		goto L65
	}
L30:
	;
	if base.Ui32(v191) <= base.Ui32(int32(3)) {
		v212 = v190
		v213 = v191
		v214 = v192
		goto L29
	} else {
		goto L61
	}
L31:
	;
	if v111 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v111 != 0 {
		v173 = v97
		goto L44
	} else {
		goto L45
	}
L34:
	;
	v212 = v92
	v213 = v97
	v214 = v90
	goto L29
L35:
	;
	goto L36
L36:
	;
	if v90&int32(3) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v190 = v92
	v191 = v97
	v192 = v90
	goto L30
L38:
	;
	goto L39
L39:
	;
	v118 = v92
	v119 = v97
	v120 = v90
	goto L40
L40:
	;
	if v119 == int32(0) {
		goto L24
	} else {
		goto L42
	}
L41:
	;
	v190 = v127
	v191 = v129
	v192 = v131
	goto L30
L42:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v124)
	v126 = int32(1)
	v127 = v118 + v126
	v129 = v119 - v126
	v131 = v120 + v126
	if v131&int32(3) != 0 {
		v118 = v127
		v119 = v129
		v120 = v131
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	if v173 == int32(0) {
		goto L24
	} else {
		goto L57
	}
L45:
	;
	if v101&int32(3) != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v138 = v97
	goto L49
L47:
	;
	v153 = v97
	goto L48
L48:
	;
	if base.Ui32(v153) <= base.Ui32(int32(3)) {
		v173 = v153
		goto L44
	} else {
		goto L53
	}
L49:
	;
	if v138 == int32(0) {
		goto L24
	} else {
		goto L51
	}
L50:
	;
	v153 = v144
	goto L48
L51:
	;
	v144 = v138 - int32(1)
	v145 = v90 + v144
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v144))))
	*(*uint8)(unsafe.Add(mBase, uint32(v145))) = uint8(v147)
	if v145&int32(3) != 0 {
		v138 = v144
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v160 = v153
	goto L54
L54:
	;
	v164 = v160 - int32(4)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v92+v164)))
	*(*int32)(unsafe.Add(mBase, uint32(v90+v164))) = v167
	if base.Ui32(int32(3)) < base.Ui32(v164) {
		v160 = v164
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v173 = v164
	goto L44
L56:
	;
	goto L55
L57:
	;
	v180 = v173
	goto L58
L58:
	;
	v184 = v180 - int32(1)
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v184))))
	*(*uint8)(unsafe.Add(mBase, uint32(v90+v184))) = uint8(v187)
	if v184 != 0 {
		v180 = v184
		goto L58
	} else {
		goto L60
	}
L59:
	;
	goto L24
L60:
	;
	goto L59
L61:
	;
	v197 = v190
	v198 = v191
	v199 = v192
	goto L62
L62:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = v201
	v203 = int32(4)
	v204 = v197 + v203
	v206 = v199 + v203
	v208 = v198 - v203
	if base.Ui32(int32(3)) < base.Ui32(v208) {
		v197 = v204
		v198 = v208
		v199 = v206
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v212 = v204
	v213 = v208
	v214 = v206
	goto L29
L64:
	;
	goto L63
L65:
	;
	v219 = v212
	v220 = v213
	v221 = v214
	goto L66
L66:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	*(*uint8)(unsafe.Add(mBase, uint32(v221))) = uint8(v223)
	v225 = int32(1)
	v230 = v220 - v225
	if v230 != 0 {
		v219 = v219 + v225
		v220 = v230
		v221 = v221 + v225
		goto L66
	} else {
		goto L68
	}
L67:
	;
	goto L24
L68:
	;
	goto L67
L69:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)))
	if v254&int32(1) != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v250 = F__emscripten_memcpy_bulkmem(m, v244+v62, v246, v249)
	mBase = m.M
	goto L72
L71:
	;
	goto L72
L72:
	;
	goto L69
L73:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	F_pfree(m, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_pfree(m, v64)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v265 = v54 + v252 - int32(1)
	v267 = v264
	goto L16
L78:
	;
	goto L15
}
