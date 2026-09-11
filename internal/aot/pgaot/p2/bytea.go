package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bytea_smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
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
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v15 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v46 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v18 = int32(4)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
	if v20&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if v15&v33 != 0 {
		v45 = int32(base.Ui32(v15)>>(uint(v33)%32)) - v33
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v29 = v18
	goto L10
L9:
	;
	v29 = base.B2i32(v20 == int32(18)) << (uint(v18) % 32)
	goto L10
L10:
	;
	if v20 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = v18
	goto L13
L12:
	;
	v32 = v29
	goto L13
L13:
	;
	v45 = v32
	goto L4
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	if v76 <= v45 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v49 = int32(4)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v51&int32(254) == int32(2) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v64 = int32(1)
	if v46&v64 != 0 {
		v76 = int32(base.Ui32(v46)>>(uint(v64)%32)) - v64
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v60 = v49
	goto L21
L20:
	;
	v60 = base.B2i32(v51 == int32(18)) << (uint(v49) % 32)
	goto L21
L21:
	;
	if v51 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v63 = v49
	goto L24
L23:
	;
	v63 = v60
	goto L24
L24:
	;
	v76 = v63
	goto L15
L25:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v78 = v13
	goto L28
L27:
	;
	v78 = v8
	goto L28
L28:
	;
	v79 = int32(1)
	if v15&v79 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v83 = v79
	goto L31
L30:
	;
	v83 = int32(4)
	goto L31
L31:
	;
	v84 = v8 + v83
	v85 = int32(1)
	if v46&v85 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v89 = v85
	goto L34
L33:
	;
	v89 = int32(4)
	goto L34
L34:
	;
	v90 = v13 + v89
	if v45 < v76 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v92 = v45
	goto L37
L36:
	;
	v92 = v76
	goto L37
L37:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v92) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	if v154 != 0 {
		goto L56
	} else {
		goto L57
	}
L39:
	;
	v154 = int32(0)
	goto L38
L40:
	;
	v128 = v123
	v129 = v124
	v130 = v125
	goto L50
L41:
	;
	if (v84|v90)&int32(3) != 0 {
		v123 = v84
		v124 = v90
		v125 = v92
		goto L40
	} else {
		goto L44
	}
L42:
	;
	v116 = v84
	v117 = v90
	v118 = v92
	goto L43
L43:
	;
	if v118 == int32(0) {
		goto L39
	} else {
		goto L49
	}
L44:
	;
	v100 = v84
	v101 = v90
	v102 = v92
	goto L45
L45:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if v105 != v106 {
		v123 = v100
		v124 = v101
		v125 = v102
		goto L40
	} else {
		goto L47
	}
L46:
	;
	v116 = v111
	v117 = v109
	v118 = v113
	goto L43
L47:
	;
	v108 = int32(4)
	v109 = v101 + v108
	v111 = v100 + v108
	v113 = v102 - v108
	if base.Ui32(int32(3)) < base.Ui32(v113) {
		v100 = v111
		v101 = v109
		v102 = v113
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v123 = v116
	v124 = v117
	v125 = v118
	goto L40
L50:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	if v133 == v134 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v154 = v133 - v134
	goto L38
L52:
	;
	v136 = int32(1)
	v141 = v130 - v136
	if v141 != 0 {
		v128 = v128 + v136
		v129 = v129 + v136
		v130 = v141
		goto L50
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	goto L51
L55:
	;
	goto L39
L56:
	;
	v155 = v13
	goto L58
L57:
	;
	v155 = v78
	goto L58
L58:
	;
	if v154 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v158 = v8
	goto L61
L60:
	;
	v158 = v155
	goto L61
L61:
	;
	return v158
}
