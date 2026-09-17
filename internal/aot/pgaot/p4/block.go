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
	var v8 int32
	_ = v8
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v270 int32
	_ = v270
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = v12
	goto L3
L2:
	;
	v14 = int32(_a_F_dump_block_0)
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_dump_block[0]))
	if int32(0) < v16 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = v2
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v14
	F_pg_printf(m, int32(_a_F_dump_block_1), v10+int32(32))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L12
	}
L7:
	;
	F_pg_printf(m, int32(_a_F_dump_block_2), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
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
	v31 = v20 + int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_dump_block[0]))
	if v31 < v33 {
		v20 = v31
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v49 = int32(_a_F_dump_block_3)
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_dump_block[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_dump_block[0])) = v51 + int32(2)
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if int32(0) < v55 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v90 = v51
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dump_block[0])) = v90
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v97 == int32(0) {
		v236 = v90
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v60 = int32(0)
	goto L19
L17:
	;
	goto L18
L18:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_dump_block[0]))
	v90 = v85 - int32(2)
	goto L15
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+v60<<(uint(int32(2))%32))))
	F_dump_stmt(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L9
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	v74 = v60 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v74 < v75 {
		v60 = v74
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if int32(0) < v236 {
		goto L60
	} else {
		goto L61
	}
L24:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	if v100 == int32(0) {
		v236 = v90
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v103 <= int32(0) {
		v236 = v90
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v108 = v90
	v112 = v2
	goto L27
L27:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v112<<(uint(int32(2))%32))))
	v118 = int32(0)
	if v118 < v108 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v236 = v223
	goto L23
L29:
	;
	v122 = v118
	goto L32
L30:
	;
	goto L31
L31:
	;
	F_pg_printf(m, int32(_a_F_dump_block_4), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L9
	} else {
		goto L36
	}
L32:
	;
	F_pg_printf(m, int32(_a_F_dump_block_2), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L9
	} else {
		goto L34
	}
L33:
	;
	goto L31
L34:
	;
	v133 = v122 + int32(1)
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_dump_block[0]))
	if v133 < v135 {
		v122 = v133
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v148 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v150 = v148
	goto L40
L38:
	;
	goto L39
L39:
	;
	F_pg_printf(m, int32(_a_F_dump_block_5), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L9
	} else {
		goto L48
	}
L40:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v156 != v150 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L39
L42:
	;
	F_pg_printf(m, int32(_a_F_dump_block_6), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L9
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v162
	F_pg_printf(m, int32(_a_F_dump_block_7), v10+int32(16))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L9
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	if v169 != 0 {
		v150 = v169
		goto L40
	} else {
		goto L47
	}
L47:
	;
	goto L41
L48:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	v182 = int32(_a_F_dump_block_3)
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_dump_block[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_dump_block[0])) = v184 + int32(2)
	if v181 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v188 = int32(0)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v188 < v189 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v223 = v184
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dump_block[0])) = v223
	v231 = v112 + int32(1)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v231 < v232 {
		v108 = v223
		v112 = v231
		goto L27
	} else {
		goto L59
	}
L52:
	;
	v193 = v188
	goto L55
L53:
	;
	goto L54
L54:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_dump_block[0]))
	v223 = v218 - int32(2)
	goto L51
L55:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v199+v193<<(uint(int32(2))%32))))
	F_dump_stmt(m, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L9
	} else {
		goto L57
	}
L56:
	;
	goto L54
L57:
	;
	v207 = v193 + int32(1)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v207 < v208 {
		v193 = v207
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
	v245 = int32(0)
	goto L63
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v14
	F_pg_printf(m, int32(_a_F_dump_block_8), v10)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L9
	} else {
		goto L67
	}
L63:
	;
	F_pg_printf(m, int32(_a_F_dump_block_2), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L9
	} else {
		goto L65
	}
L64:
	;
	goto L62
L65:
	;
	v256 = v245 + int32(1)
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_dump_block[0]))
	if v256 < v258 {
		v245 = v256
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	m.G0 = v10 + int32(48)
	return
}
