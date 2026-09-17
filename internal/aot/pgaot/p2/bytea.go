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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
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
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v45 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
	if v21 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v32 = int32(1)
	if v15&v32 != 0 {
		v44 = int32(base.Ui32(v15)>>(uint(v32)%32)) - v32
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v24 = int32(16)
	goto L10
L9:
	;
	v24 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v21-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = int32(4)
	goto L13
L12:
	;
	v31 = v24
	goto L13
L13:
	;
	v44 = v31
	goto L4
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	if v74 <= v44 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v51 == int32(18) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v62 = int32(1)
	if v45&v62 != 0 {
		v74 = int32(base.Ui32(v45)>>(uint(v62)%32)) - v62
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v54 = int32(16)
	goto L21
L20:
	;
	v54 = int32(0)
	goto L21
L21:
	;
	if base.Ui32((v51-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v61 = int32(4)
	goto L24
L23:
	;
	v61 = v54
	goto L24
L24:
	;
	v74 = v61
	goto L15
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v74 = int32(base.Ui32(v68)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v76 = v13
	goto L28
L27:
	;
	v76 = v8
	goto L28
L28:
	;
	v77 = int32(1)
	if v15&v77 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v81 = v77
	goto L31
L30:
	;
	v81 = int32(4)
	goto L31
L31:
	;
	v82 = v8 + v81
	v83 = int32(1)
	if v45&v83 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v87 = v83
	goto L34
L33:
	;
	v87 = int32(4)
	goto L34
L34:
	;
	v88 = v13 + v87
	if v44 < v74 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v90 = v44
	goto L37
L36:
	;
	v90 = v74
	goto L37
L37:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v90) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	if v152 != 0 {
		goto L56
	} else {
		goto L57
	}
L39:
	;
	v152 = int32(0)
	goto L38
L40:
	;
	v126 = v121
	v127 = v122
	v128 = v123
	goto L50
L41:
	;
	if (v82|v88)&int32(3) != 0 {
		v121 = v82
		v122 = v88
		v123 = v90
		goto L40
	} else {
		goto L44
	}
L42:
	;
	v114 = v82
	v115 = v88
	v116 = v90
	goto L43
L43:
	;
	if v116 == int32(0) {
		goto L39
	} else {
		goto L49
	}
L44:
	;
	v98 = v82
	v99 = v88
	v100 = v90
	goto L45
L45:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v103 != v104 {
		v121 = v98
		v122 = v99
		v123 = v100
		goto L40
	} else {
		goto L47
	}
L46:
	;
	v114 = v109
	v115 = v107
	v116 = v111
	goto L43
L47:
	;
	v106 = int32(4)
	v107 = v99 + v106
	v109 = v98 + v106
	v111 = v100 - v106
	if base.Ui32(int32(3)) < base.Ui32(v111) {
		v98 = v109
		v99 = v107
		v100 = v111
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v121 = v114
	v122 = v115
	v123 = v116
	goto L40
L50:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v131 == v132 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v152 = v131 - v132
	goto L38
L52:
	;
	v134 = int32(1)
	v139 = v128 - v134
	if v139 != 0 {
		v126 = v126 + v134
		v127 = v127 + v134
		v128 = v139
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
	v153 = v13
	goto L58
L57:
	;
	v153 = v76
	goto L58
L58:
	;
	if v152 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v156 = v8
	goto L61
L60:
	;
	v156 = v153
	goto L61
L61:
	;
	return v156
}
