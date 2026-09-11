package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bpchar_pattern_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
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
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = l0 + int32(28)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(1)
	v24 = v7 + v23
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v27 = v25 & v23
	if v25 == v23 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v157 != v7 {
		goto L59
	} else {
		goto L60
	}
L5:
	;
	if v27 != 0 {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v30 = int32(4)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v32&int32(254) == int32(2) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v45 = int32(1)
	if v27 != 0 {
		v55 = int32(base.Ui32(v25)>>(uint(v45)%32)) - v45
		goto L5
	} else {
		goto L15
	}
L9:
	;
	v41 = v30
	goto L11
L10:
	;
	v41 = base.B2i32(v32 == int32(18)) << (uint(v30) % 32)
	goto L11
L11:
	;
	if v32 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = v30
	goto L14
L13:
	;
	v44 = v41
	goto L14
L14:
	;
	v55 = v44
	goto L5
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
	goto L5
L16:
	;
	v58 = v24
	goto L18
L17:
	;
	v58 = v7 + int32(4)
	goto L18
L18:
	;
	v64 = v55
	goto L19
L19:
	;
	if v64 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v81 = int32(1)
	v82 = v14 + v81
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v85 = v83 & v81
	if v83 == v81 {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	goto L20
L22:
	;
	v80 = v55 >> (uint(int32(31)) % 32) & v55
	goto L21
L23:
	;
	goto L24
L24:
	;
	v74 = v64 - int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v74))))
	if v76 == int32(32) {
		v64 = v74
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v80 = v64
	goto L21
L26:
	;
	if v85 != 0 {
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v88 = int32(4)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v90&int32(254) == int32(2) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v103 = int32(1)
	if v85 != 0 {
		v113 = int32(base.Ui32(v83)>>(uint(v103)%32)) - v103
		goto L26
	} else {
		goto L36
	}
L30:
	;
	v99 = v88
	goto L32
L31:
	;
	v99 = base.B2i32(v90 == int32(18)) << (uint(v88) % 32)
	goto L32
L32:
	;
	if v90 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v102 = v88
	goto L35
L34:
	;
	v102 = v99
	goto L35
L35:
	;
	v113 = v102
	goto L26
L36:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v113 = int32(base.Ui32(v107)>>(uint(int32(2))%32)) - int32(4)
	goto L26
L37:
	;
	v116 = v82
	goto L39
L38:
	;
	v116 = v14 + int32(4)
	goto L39
L39:
	;
	v122 = v113
	goto L40
L40:
	;
	if v122 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v139 = int32(1)
	if v25&v139 != 0 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	v138 = v113 >> (uint(int32(31)) % 32) & v113
	goto L42
L44:
	;
	goto L45
L45:
	;
	v132 = v122 - int32(1)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+v132))))
	if v134 == int32(32) {
		v122 = v132
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v138 = v122
	goto L42
L47:
	;
	goto L4
L48:
	;
	v143 = v139
	goto L50
L49:
	;
	v143 = int32(4)
	goto L50
L50:
	;
	v145 = int32(1)
	if v83&v145 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v149 = v145
	goto L53
L52:
	;
	v149 = int32(4)
	goto L53
L53:
	;
	v151 = base.B2i32(v80 < v138)
	if v80 < v138 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v152 = v80
	goto L56
L55:
	;
	v152 = v138
	goto L56
L56:
	;
	v153 = F_memcmp(m, v7+v143, v14+v149, v152)
	mBase = m.M
	if v153 != 0 {
		v156 = v153
		goto L47
	} else {
		goto L57
	}
L57:
	;
	if v80 < v138 {
		v156 = int32(-1)
		goto L47
	} else {
		goto L58
	}
L58:
	;
	v156 = base.B2i32(v138 < v80)
	goto L47
L59:
	;
	F_pfree(m, v7)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v161 != v14 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	F_pfree(m, v14)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	return int32(base.Ui32(v156^int32(-1)) >> (uint(int32(31)) % 32))
L66:
	;
	goto L65
}
