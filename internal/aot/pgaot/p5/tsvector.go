package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tsvector_cmp(m *base.Module, l0 int32) int32 {
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
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
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v29 = int32(2)
	v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v33 = int32(base.Ui32(v31) >> (uint(v29) % 32))
	if base.Ui32(v30) < base.Ui32(v33) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v224 != v7 {
		goto L67
	} else {
		goto L68
	}
L5:
	;
	v223 = int32(-1)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v36 = int32(1)
	if base.Ui32(v33) < base.Ui32(v30) {
		v199 = v36
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v223 = v199
	goto L4
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v38 < v39 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v223 = int32(-1)
	goto L4
L11:
	;
	goto L12
L12:
	;
	if v39 < v38 {
		v199 = v36
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v38 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v223 = int32(0)
	goto L4
L15:
	;
	goto L16
L16:
	;
	v46 = int32(8)
	v47 = v14 + v46
	v48 = int32(2)
	v50 = v47 + v39<<(uint(v48)%32)
	v52 = v7 + v46
	v55 = v52 + v38<<(uint(v48)%32)
	v63 = v47
	v64 = v52
	v68 = int32(0)
	goto L17
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v71 = int32(1)
	v72 = v70 & v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v75 = v73 & v71
	if v72 != v75 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v199 = int32(0)
	goto L8
L19:
	;
	if base.Ui32(v75) < base.Ui32(v72) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v81 = int32(1)
	v83 = int32(2047)
	v84 = int32(base.Ui32(v73)>>(uint(v81)%32)) & v83
	v85 = int32(12)
	v86 = int32(base.Ui32(v73) >> (uint(v85) % 32))
	v88 = int32(base.Ui32(v70) >> (uint(v85) % 32))
	v92 = int32(base.Ui32(v70)>>(uint(v81)%32)) & v83
	if v92 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v80 = int32(-1)
	goto L24
L23:
	;
	v80 = int32(1)
	goto L24
L24:
	;
	v223 = v80
	goto L4
L25:
	;
	if v72 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L26:
	;
	if v84 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v84 == int32(0) {
		goto L25
	} else {
		goto L40
	}
L29:
	;
	v223 = int32(1)
	goto L4
L30:
	;
	goto L31
L31:
	;
	v98 = base.B2i32(base.Ui32(v92) < base.Ui32(v84))
	if base.Ui32(v92) < base.Ui32(v84) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v92
	goto L34
L33:
	;
	v99 = v84
	goto L34
L34:
	;
	v100 = F_memcmp(m, v88+v55, v86+v50, v99)
	mBase = m.M
	if v100 != 0 {
		v199 = v100
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v92 == v84 {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v92) < base.Ui32(v84) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v104 = int32(-1)
	goto L39
L38:
	;
	v104 = int32(1)
	goto L39
L39:
	;
	v223 = v104
	goto L4
L40:
	;
	v223 = int32(-1)
	goto L4
L41:
	;
	v187 = int32(4)
	v193 = v68 + int32(1)
	if v193 != v38 {
		v63 = v63 + v187
		v64 = v64 + v187
		v68 = v193
		goto L17
	} else {
		goto L66
	}
L42:
	;
	v113 = int32(1)
	v115 = int32(4194302)
	v117 = v50 + (v84+v86+v113)&v115
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117))))
	v124 = v55 + (v92+v88+v113)&v115
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	if v118 == v125 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v132 = v124
	v133 = v117
	v134 = int32(0)
	goto L51
L44:
	;
	if v125 != 0 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v118) < base.Ui32(v125) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L41
L48:
	;
	v131 = int32(-1)
	goto L50
L49:
	;
	v131 = int32(1)
	goto L50
L50:
	;
	v223 = v131
	goto L4
L51:
	;
	v146 = int32(2)
	v147 = v132 + v146
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147))))
	v149 = int32(16383)
	v150 = v148 & v149
	v152 = v133 + v146
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	v155 = v153 & v149
	if v150 != v155 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(v164) < base.Ui32(v162) {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	if base.Ui32(v155) < base.Ui32(v150) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v161 = int32(14)
	v162 = int32(base.Ui32(v148) >> (uint(v161) % 32))
	v164 = int32(base.Ui32(v153) >> (uint(v161) % 32))
	if v162 == v164 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v160 = int32(-1)
	goto L58
L57:
	;
	v160 = int32(1)
	goto L58
L58:
	;
	v223 = v160
	goto L4
L59:
	;
	v167 = v134 + int32(1)
	if v167 == v125 {
		goto L41
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L52
L62:
	;
	v132 = v147
	v133 = v152
	v134 = v167
	goto L51
L63:
	;
	v172 = int32(-1)
	goto L65
L64:
	;
	v172 = int32(1)
	goto L65
L65:
	;
	v199 = v172
	goto L8
L66:
	;
	goto L18
L67:
	;
	F_pfree(m, v7)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v228 != v14 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v14)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	return v223
L74:
	;
	goto L73
}
func F_tsvector_ge(m *base.Module, l0 int32) int32 {
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
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
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v29 = int32(2)
	v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v33 = int32(base.Ui32(v31) >> (uint(v29) % 32))
	if base.Ui32(v30) < base.Ui32(v33) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v224 != v7 {
		goto L67
	} else {
		goto L68
	}
L5:
	;
	v223 = int32(-1)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v36 = int32(1)
	if base.Ui32(v33) < base.Ui32(v30) {
		v199 = v36
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v223 = v199
	goto L4
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v38 < v39 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v223 = int32(-1)
	goto L4
L11:
	;
	goto L12
L12:
	;
	if v39 < v38 {
		v199 = v36
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v38 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v223 = int32(0)
	goto L4
L15:
	;
	goto L16
L16:
	;
	v46 = int32(8)
	v47 = v14 + v46
	v48 = int32(2)
	v50 = v47 + v39<<(uint(v48)%32)
	v52 = v7 + v46
	v55 = v52 + v38<<(uint(v48)%32)
	v63 = v47
	v64 = v52
	v68 = int32(0)
	goto L17
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v71 = int32(1)
	v72 = v70 & v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v75 = v73 & v71
	if v72 != v75 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v199 = int32(0)
	goto L8
L19:
	;
	if base.Ui32(v75) < base.Ui32(v72) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v81 = int32(1)
	v83 = int32(2047)
	v84 = int32(base.Ui32(v73)>>(uint(v81)%32)) & v83
	v85 = int32(12)
	v86 = int32(base.Ui32(v73) >> (uint(v85) % 32))
	v88 = int32(base.Ui32(v70) >> (uint(v85) % 32))
	v92 = int32(base.Ui32(v70)>>(uint(v81)%32)) & v83
	if v92 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v80 = int32(-1)
	goto L24
L23:
	;
	v80 = int32(1)
	goto L24
L24:
	;
	v223 = v80
	goto L4
L25:
	;
	if v72 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L26:
	;
	if v84 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v84 == int32(0) {
		goto L25
	} else {
		goto L40
	}
L29:
	;
	v223 = int32(1)
	goto L4
L30:
	;
	goto L31
L31:
	;
	v98 = base.B2i32(base.Ui32(v92) < base.Ui32(v84))
	if base.Ui32(v92) < base.Ui32(v84) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v92
	goto L34
L33:
	;
	v99 = v84
	goto L34
L34:
	;
	v100 = F_memcmp(m, v88+v55, v86+v50, v99)
	mBase = m.M
	if v100 != 0 {
		v199 = v100
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v92 == v84 {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v92) < base.Ui32(v84) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v104 = int32(-1)
	goto L39
L38:
	;
	v104 = int32(1)
	goto L39
L39:
	;
	v223 = v104
	goto L4
L40:
	;
	v223 = int32(-1)
	goto L4
L41:
	;
	v187 = int32(4)
	v193 = v68 + int32(1)
	if v193 != v38 {
		v63 = v63 + v187
		v64 = v64 + v187
		v68 = v193
		goto L17
	} else {
		goto L66
	}
L42:
	;
	v113 = int32(1)
	v115 = int32(4194302)
	v117 = v50 + (v84+v86+v113)&v115
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117))))
	v124 = v55 + (v92+v88+v113)&v115
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	if v118 == v125 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v132 = v124
	v133 = v117
	v134 = int32(0)
	goto L51
L44:
	;
	if v125 != 0 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v118) < base.Ui32(v125) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L41
L48:
	;
	v131 = int32(-1)
	goto L50
L49:
	;
	v131 = int32(1)
	goto L50
L50:
	;
	v223 = v131
	goto L4
L51:
	;
	v146 = int32(2)
	v147 = v132 + v146
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147))))
	v149 = int32(16383)
	v150 = v148 & v149
	v152 = v133 + v146
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	v155 = v153 & v149
	if v150 != v155 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(v164) < base.Ui32(v162) {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	if base.Ui32(v155) < base.Ui32(v150) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v161 = int32(14)
	v162 = int32(base.Ui32(v148) >> (uint(v161) % 32))
	v164 = int32(base.Ui32(v153) >> (uint(v161) % 32))
	if v162 == v164 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v160 = int32(-1)
	goto L58
L57:
	;
	v160 = int32(1)
	goto L58
L58:
	;
	v223 = v160
	goto L4
L59:
	;
	v167 = v134 + int32(1)
	if v167 == v125 {
		goto L41
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L52
L62:
	;
	v132 = v147
	v133 = v152
	v134 = v167
	goto L51
L63:
	;
	v172 = int32(-1)
	goto L65
L64:
	;
	v172 = int32(1)
	goto L65
L65:
	;
	v199 = v172
	goto L8
L66:
	;
	goto L18
L67:
	;
	F_pfree(m, v7)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v228 != v14 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v14)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	return int32(base.Ui32(v223^int32(-1)) >> (uint(int32(31)) % 32))
L74:
	;
	goto L73
}
func F_tsvector_lt(m *base.Module, l0 int32) int32 {
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
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
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v29 = int32(2)
	v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v33 = int32(base.Ui32(v31) >> (uint(v29) % 32))
	if base.Ui32(v30) < base.Ui32(v33) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v224 != v7 {
		goto L67
	} else {
		goto L68
	}
L5:
	;
	v223 = int32(-1)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v36 = int32(1)
	if base.Ui32(v33) < base.Ui32(v30) {
		v199 = v36
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v223 = v199
	goto L4
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v38 < v39 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v223 = int32(-1)
	goto L4
L11:
	;
	goto L12
L12:
	;
	if v39 < v38 {
		v199 = v36
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v38 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v223 = int32(0)
	goto L4
L15:
	;
	goto L16
L16:
	;
	v46 = int32(8)
	v47 = v14 + v46
	v48 = int32(2)
	v50 = v47 + v39<<(uint(v48)%32)
	v52 = v7 + v46
	v55 = v52 + v38<<(uint(v48)%32)
	v63 = v47
	v64 = v52
	v68 = int32(0)
	goto L17
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v71 = int32(1)
	v72 = v70 & v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v75 = v73 & v71
	if v72 != v75 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v199 = int32(0)
	goto L8
L19:
	;
	if base.Ui32(v75) < base.Ui32(v72) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v81 = int32(1)
	v83 = int32(2047)
	v84 = int32(base.Ui32(v73)>>(uint(v81)%32)) & v83
	v85 = int32(12)
	v86 = int32(base.Ui32(v73) >> (uint(v85) % 32))
	v88 = int32(base.Ui32(v70) >> (uint(v85) % 32))
	v92 = int32(base.Ui32(v70)>>(uint(v81)%32)) & v83
	if v92 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v80 = int32(-1)
	goto L24
L23:
	;
	v80 = int32(1)
	goto L24
L24:
	;
	v223 = v80
	goto L4
L25:
	;
	if v72 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L26:
	;
	if v84 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v84 == int32(0) {
		goto L25
	} else {
		goto L40
	}
L29:
	;
	v223 = int32(1)
	goto L4
L30:
	;
	goto L31
L31:
	;
	v98 = base.B2i32(base.Ui32(v92) < base.Ui32(v84))
	if base.Ui32(v92) < base.Ui32(v84) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v92
	goto L34
L33:
	;
	v99 = v84
	goto L34
L34:
	;
	v100 = F_memcmp(m, v88+v55, v86+v50, v99)
	mBase = m.M
	if v100 != 0 {
		v199 = v100
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v92 == v84 {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v92) < base.Ui32(v84) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v104 = int32(-1)
	goto L39
L38:
	;
	v104 = int32(1)
	goto L39
L39:
	;
	v223 = v104
	goto L4
L40:
	;
	v223 = int32(-1)
	goto L4
L41:
	;
	v187 = int32(4)
	v193 = v68 + int32(1)
	if v193 != v38 {
		v63 = v63 + v187
		v64 = v64 + v187
		v68 = v193
		goto L17
	} else {
		goto L66
	}
L42:
	;
	v113 = int32(1)
	v115 = int32(4194302)
	v117 = v50 + (v84+v86+v113)&v115
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117))))
	v124 = v55 + (v92+v88+v113)&v115
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	if v118 == v125 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v132 = v124
	v133 = v117
	v134 = int32(0)
	goto L51
L44:
	;
	if v125 != 0 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v118) < base.Ui32(v125) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L41
L48:
	;
	v131 = int32(-1)
	goto L50
L49:
	;
	v131 = int32(1)
	goto L50
L50:
	;
	v223 = v131
	goto L4
L51:
	;
	v146 = int32(2)
	v147 = v132 + v146
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147))))
	v149 = int32(16383)
	v150 = v148 & v149
	v152 = v133 + v146
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	v155 = v153 & v149
	if v150 != v155 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(v164) < base.Ui32(v162) {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	if base.Ui32(v155) < base.Ui32(v150) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v161 = int32(14)
	v162 = int32(base.Ui32(v148) >> (uint(v161) % 32))
	v164 = int32(base.Ui32(v153) >> (uint(v161) % 32))
	if v162 == v164 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v160 = int32(-1)
	goto L58
L57:
	;
	v160 = int32(1)
	goto L58
L58:
	;
	v223 = v160
	goto L4
L59:
	;
	v167 = v134 + int32(1)
	if v167 == v125 {
		goto L41
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L52
L62:
	;
	v132 = v147
	v133 = v152
	v134 = v167
	goto L51
L63:
	;
	v172 = int32(-1)
	goto L65
L64:
	;
	v172 = int32(1)
	goto L65
L65:
	;
	v199 = v172
	goto L8
L66:
	;
	goto L18
L67:
	;
	F_pfree(m, v7)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v228 != v14 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v14)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	return int32(base.Ui32(v223) >> (uint(int32(31)) % 32))
L74:
	;
	goto L73
}
func F_tsvector_update_trigger(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v397 int32
	_ = v397
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	v15 = m.G0
	v17 = v15 - int32(144)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L19
	} else {
		goto L166
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L19
	} else {
		goto L162
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L19
	} else {
		goto L158
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L19
	} else {
		goto L154
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L19
	} else {
		goto L150
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L19
	} else {
		goto L146
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L19
	} else {
		goto L142
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L19
	} else {
		goto L139
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L19
	} else {
		goto L136
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L19
	} else {
		goto L133
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L19
	} else {
		goto L130
	}
L12:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v22 != int32(442) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v25&int32(4) == int32(0) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	if v25&int32(24) != int32(8) {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	switch v25 & int32(3) {
	case 0:
		v55 = int32(12)
		v56 = int32(1)
		goto L16
	default:
		goto L18
	case 2:
		goto L17
	}
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v57)+34)))
	if v58 <= int32(2) {
		goto L8
	} else {
		goto L23
	}
L17:
	;
	v55 = int32(16)
	v56 = int32(0)
	goto L16
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	F_errmsg_internal(m, int32(526767), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(484666), int32(2779), int32(219505))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v55+v19)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v67 < v70 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+140)) = v115
	if v115 == int32(-9) {
		goto L7
	} else {
		goto L39
	}
L25:
	;
	v115 = v77 + int32(1)
	goto L24
L26:
	;
	v77 = v67
	v78 = v70
	goto L29
L27:
	;
	goto L28
L28:
	;
	v103 = F_SystemAttributeByName(m, v66)
	mBase = m.M
	if v103 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v80 = int32(4)
	v85 = v64 + int32(20) + v78<<(uint(v80)%32) + v77*int32(100)
	v88 = F_namestrcmp(m, v85+v80, v66)
	mBase = m.M
	if v88 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+91)))
	if v91 != int32(1) {
		goto L25
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v95 = v77 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v95 < v96 {
		v77 = v95
		v78 = v96
		goto L29
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	goto L30
L36:
	;
	v115 = int32(-9)
	goto L24
L37:
	;
	goto L38
L38:
	;
	v107 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+74)))
	v115 = v107
	goto L24
L39:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v120 = F_SPI_gettypeid(m, v119, v115)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	v123 = F_IsBinaryCoercible(m, v120, int32(3614))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L19
	} else {
		goto L41
	}
L41:
	;
	if v123 == int32(0) {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	if l1 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+136)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = int64(32)
	v237 = F_palloc(m, int32(512))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L19
	} else {
		goto L76
	}
L44:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v130 = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	if v130 < v133 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	goto L46
L46:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	v220 = F_stringToQualifiedNameList(m, v218, int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L19
	} else {
		goto L72
	}
L47:
	;
	if v178 == int32(-9) {
		goto L5
	} else {
		goto L62
	}
L48:
	;
	v178 = v140 + int32(1)
	goto L47
L49:
	;
	v140 = v130
	v141 = v133
	goto L52
L50:
	;
	goto L51
L51:
	;
	v166 = F_SystemAttributeByName(m, v129)
	mBase = m.M
	if v166 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L52:
	;
	v143 = int32(4)
	v148 = v127 + int32(20) + v141<<(uint(v143)%32) + v140*int32(100)
	v151 = F_namestrcmp(m, v148+v143, v129)
	mBase = m.M
	if v151 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L51
L54:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+91)))
	if v154 != int32(1) {
		goto L48
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v158 = v140 + int32(1)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	if v158 < v159 {
		v140 = v158
		v141 = v159
		goto L52
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	goto L53
L59:
	;
	v178 = int32(-9)
	goto L47
L60:
	;
	goto L61
L61:
	;
	v170 = int32(*(*int16)(unsafe.Add(mBase, uint32(v166)+74)))
	v178 = v170
	goto L47
L62:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v182 = F_SPI_gettypeid(m, v181, v178)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L19
	} else {
		goto L63
	}
L63:
	;
	v185 = F_IsBinaryCoercible(m, v182, int32(3734))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L19
	} else {
		goto L64
	}
L64:
	;
	if v185 == int32(0) {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v192 = F_SPI_getbinval(m, v62, v189, v178, v17+int32(119))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L19
	} else {
		goto L66
	}
L66:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)))
	if v194 != int32(1) {
		v231 = v192
		goto L43
	} else {
		goto L67
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L19
	} else {
		goto L68
	}
L68:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L19
	} else {
		goto L69
	}
L69:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v205
	F_errmsg(m, int32(296762), v17+int32(32))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L19
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(484666), int32(2825), int32(219505))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L19
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	if v220 == int32(0) {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v224 <= int32(1) {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	v228 = F_get_ts_config_oid(m, v220, int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L19
	} else {
		goto L75
	}
L75:
	;
	v231 = v228
	goto L43
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+124)) = v237
	v240 = int32(*(*int16)(unsafe.Add(mBase, uint32(v57)+34)))
	if int32(3) <= v240 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v245 = int32(2)
	v252 = v56
	goto L80
L78:
	;
	v397 = v56
	goto L79
L79:
	;
	if v397&int32(1) != 0 {
		goto L124
	} else {
		goto L125
	}
L80:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v259+v245<<(uint(int32(2))%32))))
	v264 = int32(0)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	if v264 < v267 {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v397 = v384
	goto L79
L82:
	;
	if v312 == int32(-9) {
		goto L2
	} else {
		goto L97
	}
L83:
	;
	v312 = v274 + int32(1)
	goto L82
L84:
	;
	v274 = v264
	v275 = v267
	goto L87
L85:
	;
	goto L86
L86:
	;
	v300 = F_SystemAttributeByName(m, v263)
	mBase = m.M
	if v300 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L87:
	;
	v277 = int32(4)
	v282 = v258 + int32(20) + v275<<(uint(v277)%32) + v274*int32(100)
	v285 = F_namestrcmp(m, v282+v277, v263)
	mBase = m.M
	if v285 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L86
L89:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+91)))
	if v288 != int32(1) {
		goto L83
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v292 = v274 + int32(1)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	if v292 < v293 {
		v274 = v292
		v275 = v293
		goto L87
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	goto L88
L94:
	;
	v312 = int32(-9)
	goto L82
L95:
	;
	goto L96
L96:
	;
	v304 = int32(*(*int16)(unsafe.Add(mBase, uint32(v300)+74)))
	v312 = v304
	goto L82
L97:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v316 = F_SPI_gettypeid(m, v315, v312)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L19
	} else {
		goto L98
	}
L98:
	;
	v319 = F_IsBinaryCoercible(m, v316, int32(25))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L19
	} else {
		goto L99
	}
L99:
	;
	if v319 == int32(0) {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v326 = F_bms_is_member(m, v312+int32(7), v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L19
	} else {
		goto L101
	}
L101:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v331 = F_SPI_getbinval(m, v62, v328, v312, v17+int32(119))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L19
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+120)) = v331
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)))
	if v334 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v384 = v252 | v326
	v386 = v245 + int32(1)
	v387 = int32(*(*int16)(unsafe.Add(mBase, uint32(v57)+34)))
	if v386 < v387 {
		v245 = v386
		v252 = v384
		goto L80
	} else {
		goto L123
	}
L104:
	;
	v337 = F_pg_detoast_datum_packed(m, v331)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L19
	} else {
		goto L105
	}
L105:
	;
	v339 = int32(1)
	v340 = v337 + v339
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	v345 = v343 & v339
	if v345 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v346 = v340
	goto L108
L107:
	;
	v346 = v337 + int32(4)
	goto L108
L108:
	;
	if v343 == int32(1) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	F_parsetext(m, v231, v17+int32(124), v346, v374)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L19
	} else {
		goto L120
	}
L110:
	;
	v349 = int32(4)
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	if v351&int32(254) == int32(2) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	v364 = int32(1)
	if v345 != 0 {
		v374 = int32(base.Ui32(v343)>>(uint(v364)%32)) - v364
		goto L109
	} else {
		goto L119
	}
L113:
	;
	v360 = v349
	goto L115
L114:
	;
	v360 = base.B2i32(v351 == int32(18)) << (uint(v349) % 32)
	goto L115
L115:
	;
	if v351 == int32(1) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v363 = v349
	goto L118
L117:
	;
	v363 = v360
	goto L118
L118:
	;
	v374 = v363
	goto L109
L119:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	v374 = int32(base.Ui32(v368)>>(uint(int32(2))%32)) - int32(4)
	goto L109
L120:
	;
	if v337 == v331 {
		goto L103
	} else {
		goto L121
	}
L121:
	;
	F_pfree(m, v337)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L19
	} else {
		goto L122
	}
L122:
	;
	goto L103
L123:
	;
	goto L81
L124:
	;
	v407 = F_make_tsvector(m, v17+int32(124))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L19
	} else {
		goto L127
	}
L125:
	;
	v426 = v62
	goto L126
L126:
	;
	m.G0 = v17 + int32(144)
	return v426
L127:
	;
	v409 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)) = uint8(v409)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+120)) = v407
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v420 = F_heap_modify_tuple_by_cols(m, v62, v412, int32(1), v17+int32(140), v17+int32(120), v17+int32(119))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L19
	} else {
		goto L128
	}
L128:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v17)+120))
	F_pfree(m, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L19
	} else {
		goto L129
	}
L129:
	;
	v426 = v420
	goto L126
L130:
	;
	F_errmsg_internal(m, int32(220944), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L19
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(484666), int32(2760), int32(219505))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L19
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errmsg_internal(m, int32(29383), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L19
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(484666), int32(2764), int32(219505))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L19
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errmsg_internal(m, int32(89612), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L19
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(484666), int32(2766), int32(219505))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L19
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
	F_errmsg_internal(m, int32(653760), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L19
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(484666), int32(2785), int32(219505))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L19
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
	F_errcode(m, int32(50360452))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L19
	} else {
		goto L143
	}
L143:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v491
	F_errmsg(m, int32(70301), v17)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L19
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(484666), int32(2793), int32(219505))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L19
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L19
	} else {
		goto L147
	}
L147:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v508)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v509
	F_errmsg(m, int32(360011), v17+int32(112))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L19
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(484666), int32(2800), int32(219505))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L19
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
	F_errcode(m, int32(50360452))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L19
	} else {
		goto L151
	}
L151:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v529
	F_errmsg(m, int32(70337), v17+int32(16))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L19
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(484666), int32(2812), int32(219505))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L19
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L19
	} else {
		goto L155
	}
L155:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v549
	F_errmsg(m, int32(361176), v17+int32(80))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L19
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(484666), int32(2818), int32(219505))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L19
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L19
	} else {
		goto L159
	}
L159:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v569
	F_errmsg(m, int32(446712), v17+int32(96))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L19
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(484666), int32(2838), int32(219505))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L19
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
	F_errcode(m, int32(50360452))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L19
	} else {
		goto L163
	}
L163:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v588+v245<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v592
	F_errmsg(m, int32(70351), v17+int32(48))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L19
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(484666), int32(2858), int32(219505))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L19
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L19
	} else {
		goto L167
	}
L167:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v57)+44))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v611+v245<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v615
	F_errmsg(m, int32(360082), v17-int32(-64))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L19
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(484666), int32(2863), int32(219505))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L19
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsvector_update_trigger_bycolumn(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_tsvector_update_trigger(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
