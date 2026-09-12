package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dump_block(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v303 int32
	_ = v303
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = v13
	goto L3
L2:
	;
	v17 = int32(694817)
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[1469]))
	if int32(0) < v20 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = v2
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v17
	F_pg_printf(m, int32(782200), v11+int32(32))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L9
	} else {
		goto L12
	}
L7:
	;
	F_pg_printf(m, int32(774629), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	return
L10:
	;
	v38 = v24 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, _consts[1469]))
	if v38 < v41 {
		v24 = v38
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v60 = int32(4642264)
	v62 = *(*int32)(unsafe.Add(mBase, _consts[1469]))
	*(*int32)(unsafe.Add(mBase, _consts[1469])) = v62 + int32(2)
	if v59 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if int32(0) < v66 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v105 = v62
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1469])) = v105
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v111 == int32(0) {
		v265 = v105
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v71 = int32(0)
	goto L19
L17:
	;
	goto L18
L18:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[1469]))
	v105 = v98 - int32(2)
	goto L15
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v71<<(uint(int32(2))%32))))
	F_dump_stmt(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	v86 = v71 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v86 < v87 {
		v71 = v86
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if int32(0) < v265 {
		goto L60
	} else {
		goto L61
	}
L24:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	if v114 == int32(0) {
		v265 = v105
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v117 <= int32(0) {
		v265 = v105
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v124 = v105
	v126 = v2
	goto L27
L27:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128+v126<<(uint(int32(2))%32))))
	v133 = int32(0)
	if v133 < v124 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v265 = v251
	goto L23
L29:
	;
	v137 = v133
	goto L32
L30:
	;
	goto L31
L31:
	;
	F_pg_printf(m, int32(772832), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L9
	} else {
		goto L36
	}
L32:
	;
	F_pg_printf(m, int32(774629), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L9
	} else {
		goto L34
	}
L33:
	;
	goto L31
L34:
	;
	v151 = v137 + int32(1)
	v154 = *(*int32)(unsafe.Add(mBase, _consts[1469]))
	if v151 < v154 {
		v137 = v151
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v168 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v170 = v168
	goto L40
L38:
	;
	goto L39
L39:
	;
	F_pg_printf(m, int32(780448), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L48
	}
L40:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v177 != v170 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L39
L42:
	;
	F_pg_printf(m, int32(772677), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L9
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v183
	F_pg_printf(m, int32(215163), v11+int32(16))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L9
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
	if v190 != 0 {
		v170 = v190
		goto L40
	} else {
		goto L47
	}
L47:
	;
	goto L41
L48:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	v206 = int32(4642264)
	v208 = *(*int32)(unsafe.Add(mBase, _consts[1469]))
	*(*int32)(unsafe.Add(mBase, _consts[1469])) = v208 + int32(2)
	if v205 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v212 = int32(0)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	if v212 < v213 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v251 = v208
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1469])) = v251
	v258 = v126 + int32(1)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v258 < v259 {
		v124 = v251
		v126 = v258
		goto L27
	} else {
		goto L59
	}
L52:
	;
	v217 = v212
	goto L55
L53:
	;
	goto L54
L54:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _consts[1469]))
	v251 = v244 - int32(2)
	goto L51
L55:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v205)+12))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v224+v217<<(uint(int32(2))%32))))
	F_dump_stmt(m, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L9
	} else {
		goto L57
	}
L56:
	;
	goto L54
L57:
	;
	v232 = v217 + int32(1)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	if v232 < v233 {
		v217 = v232
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	goto L28
L60:
	;
	v273 = int32(0)
	goto L63
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v17
	F_pg_printf(m, int32(776557), v11)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L9
	} else {
		goto L67
	}
L63:
	;
	F_pg_printf(m, int32(774629), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L9
	} else {
		goto L65
	}
L64:
	;
	goto L62
L65:
	;
	v287 = v273 + int32(1)
	v290 = *(*int32)(unsafe.Add(mBase, _consts[1469]))
	if v287 < v290 {
		v273 = v287
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	m.G0 = v11 + int32(48)
	return
}
