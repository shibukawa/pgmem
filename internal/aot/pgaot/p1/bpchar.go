package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bpchar_pattern_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(1)
	v21 = v6 + v20
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v26 = v24 & v20
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v152 != v6 {
		goto L59
	} else {
		goto L60
	}
L5:
	;
	v27 = v21
	goto L7
L6:
	;
	v27 = v6 + int32(4)
	goto L7
L7:
	;
	if v24 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v61 = v54
	goto L19
L9:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v33 == int32(18) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v44 = int32(1)
	if v26 != 0 {
		v54 = int32(base.Ui32(v24)>>(uint(v44)%32)) - v44
		goto L8
	} else {
		goto L18
	}
L12:
	;
	v36 = int32(16)
	goto L14
L13:
	;
	v36 = int32(0)
	goto L14
L14:
	;
	if base.Ui32((v33-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v43 = int32(4)
	goto L17
L16:
	;
	v43 = v36
	goto L17
L17:
	;
	v54 = v43
	goto L8
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L19:
	;
	if v61 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v77 = int32(1)
	v78 = v11 + v77
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v83 = v81 & v77
	if v83 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	goto L20
L22:
	;
	v76 = v54 & (v54 >> (uint(int32(31)) % 32))
	goto L21
L23:
	;
	goto L24
L24:
	;
	v70 = v61 - int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v70))))
	if v72 == int32(32) {
		v61 = v70
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v76 = v61
	goto L21
L26:
	;
	v84 = v78
	goto L28
L27:
	;
	v84 = v11 + int32(4)
	goto L28
L28:
	;
	if v81 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v118 = v111
	goto L40
L30:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v90 == int32(18) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v101 = int32(1)
	if v83 != 0 {
		v111 = int32(base.Ui32(v81)>>(uint(v101)%32)) - v101
		goto L29
	} else {
		goto L39
	}
L33:
	;
	v93 = int32(16)
	goto L35
L34:
	;
	v93 = int32(0)
	goto L35
L35:
	;
	if base.Ui32((v90-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v100 = int32(4)
	goto L38
L37:
	;
	v100 = v93
	goto L38
L38:
	;
	v111 = v100
	goto L29
L39:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v111 = int32(base.Ui32(v105)>>(uint(int32(2))%32)) - int32(4)
	goto L29
L40:
	;
	if v118 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v134 = int32(1)
	if v24&v134 != 0 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	v132 = v111 & (v111 >> (uint(int32(31)) % 32))
	goto L42
L44:
	;
	goto L45
L45:
	;
	v127 = v118 - int32(1)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v127))))
	if v129 == int32(32) {
		v118 = v127
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v132 = v118
	goto L42
L47:
	;
	goto L4
L48:
	;
	v138 = v134
	goto L50
L49:
	;
	v138 = int32(4)
	goto L50
L50:
	;
	v140 = int32(1)
	if v81&v140 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v144 = v140
	goto L53
L52:
	;
	v144 = int32(4)
	goto L53
L53:
	;
	v146 = base.B2i32(v76 < v132)
	if v76 < v132 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v147 = v76
	goto L56
L55:
	;
	v147 = v132
	goto L56
L56:
	;
	v148 = F_memcmp(m, v6+v138, v11+v144, v147)
	mBase = m.M
	if v148 != 0 {
		v151 = v148
		goto L47
	} else {
		goto L57
	}
L57:
	;
	if v76 < v132 {
		v151 = int32(-1)
		goto L47
	} else {
		goto L58
	}
L58:
	;
	v151 = base.B2i32(v132 < v76)
	goto L47
L59:
	;
	F_pfree(m, v6)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v156 != v11 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	F_pfree(m, v11)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	return base.B2i32(int32(0) < v151)
L66:
	;
	goto L65
}
func F_bpchar_pattern_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(1)
	v21 = v6 + v20
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v26 = v24 & v20
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v152 != v6 {
		goto L59
	} else {
		goto L60
	}
L5:
	;
	v27 = v21
	goto L7
L6:
	;
	v27 = v6 + int32(4)
	goto L7
L7:
	;
	if v24 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v61 = v54
	goto L19
L9:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v33 == int32(18) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v44 = int32(1)
	if v26 != 0 {
		v54 = int32(base.Ui32(v24)>>(uint(v44)%32)) - v44
		goto L8
	} else {
		goto L18
	}
L12:
	;
	v36 = int32(16)
	goto L14
L13:
	;
	v36 = int32(0)
	goto L14
L14:
	;
	if base.Ui32((v33-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v43 = int32(4)
	goto L17
L16:
	;
	v43 = v36
	goto L17
L17:
	;
	v54 = v43
	goto L8
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L19:
	;
	if v61 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v77 = int32(1)
	v78 = v11 + v77
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v83 = v81 & v77
	if v83 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	goto L20
L22:
	;
	v76 = v54 & (v54 >> (uint(int32(31)) % 32))
	goto L21
L23:
	;
	goto L24
L24:
	;
	v70 = v61 - int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v70))))
	if v72 == int32(32) {
		v61 = v70
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v76 = v61
	goto L21
L26:
	;
	v84 = v78
	goto L28
L27:
	;
	v84 = v11 + int32(4)
	goto L28
L28:
	;
	if v81 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v118 = v111
	goto L40
L30:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v90 == int32(18) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v101 = int32(1)
	if v83 != 0 {
		v111 = int32(base.Ui32(v81)>>(uint(v101)%32)) - v101
		goto L29
	} else {
		goto L39
	}
L33:
	;
	v93 = int32(16)
	goto L35
L34:
	;
	v93 = int32(0)
	goto L35
L35:
	;
	if base.Ui32((v90-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v100 = int32(4)
	goto L38
L37:
	;
	v100 = v93
	goto L38
L38:
	;
	v111 = v100
	goto L29
L39:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v111 = int32(base.Ui32(v105)>>(uint(int32(2))%32)) - int32(4)
	goto L29
L40:
	;
	if v118 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v134 = int32(1)
	if v24&v134 != 0 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	v132 = v111 & (v111 >> (uint(int32(31)) % 32))
	goto L42
L44:
	;
	goto L45
L45:
	;
	v127 = v118 - int32(1)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v127))))
	if v129 == int32(32) {
		v118 = v127
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v132 = v118
	goto L42
L47:
	;
	goto L4
L48:
	;
	v138 = v134
	goto L50
L49:
	;
	v138 = int32(4)
	goto L50
L50:
	;
	v140 = int32(1)
	if v81&v140 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v144 = v140
	goto L53
L52:
	;
	v144 = int32(4)
	goto L53
L53:
	;
	v146 = base.B2i32(v76 < v132)
	if v76 < v132 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v147 = v76
	goto L56
L55:
	;
	v147 = v132
	goto L56
L56:
	;
	v148 = F_memcmp(m, v6+v138, v11+v144, v147)
	mBase = m.M
	if v148 != 0 {
		v151 = v148
		goto L47
	} else {
		goto L57
	}
L57:
	;
	if v76 < v132 {
		v151 = int32(-1)
		goto L47
	} else {
		goto L58
	}
L58:
	;
	v151 = base.B2i32(v132 < v76)
	goto L47
L59:
	;
	F_pfree(m, v6)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v156 != v11 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	F_pfree(m, v11)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	return base.B2i32(v151 <= int32(0))
L66:
	;
	goto L65
}
