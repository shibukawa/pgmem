package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bpchar(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v17 < int32(4) {
		v124 = v13
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v10 + int32(16)
	return v124
L4:
	;
	v21 = v17 - int32(4)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v22 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v53 = int32(1)
	if v22&v53 != 0 {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v28 == int32(18) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v39 = int32(1)
	if v22&v39 != 0 {
		v51 = int32(base.Ui32(v22)>>(uint(v39)%32)) - v39
		goto L5
	} else {
		goto L15
	}
L9:
	;
	v31 = int32(16)
	goto L11
L10:
	;
	v31 = int32(0)
	goto L11
L11:
	;
	if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v38 = int32(4)
	goto L14
L13:
	;
	v38 = v31
	goto L14
L14:
	;
	v51 = v38
	goto L5
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L5
L16:
	;
	v57 = v53
	goto L18
L17:
	;
	v57 = int32(4)
	goto L18
L18:
	;
	v58 = v13 + v57
	v59 = F_pg_mbstrlen_with_len(m, v58, v51)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v59 == v21 {
		v124 = v13
		goto L3
	} else {
		goto L20
	}
L20:
	;
	if v21 < v59 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v106 = v98 + int32(4)
	v107 = F_palloc(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L41
	}
L22:
	;
	v63 = F_pg_mbcharcliplen(m, v58, v51, v21)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v98 = v51 + v21 - v59
	v99 = v51
	goto L21
L25:
	;
	if v52 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v98 = v63
	v99 = v63
	goto L21
L27:
	;
	goto L28
L28:
	;
	if v51 <= v63 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v98 = v63
	v99 = v63
	goto L21
L30:
	;
	goto L31
L31:
	;
	v68 = v63
	goto L33
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L37
	}
L33:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v58))))
	if v74 != int32(32) {
		goto L32
	} else {
		goto L35
	}
L34:
	;
	v98 = v63
	v99 = v63
	goto L21
L35:
	;
	v78 = v68 + int32(1)
	if v78 != v51 {
		v68 = v78
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v21
	F_errmsg(m, int32(_a_F_bpchar_0), v10)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_bpchar_1), int32(313), int32(_a_F_bpchar_2))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v106 << (uint(int32(2)) % 32)
	v113 = v107 + int32(4)
	if v99 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	base.MemoryCopy(m, v113, v58, v99)
	goto L44
L43:
	;
	goto L44
L44:
	;
	if v98 <= v99 {
		v124 = v107
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v116 = v98 - v99
	if v116 == int32(0) {
		v124 = v107
		goto L3
	} else {
		goto L46
	}
L46:
	;
	base.MemoryFill(m, v99+v113, int32(32), v116)
	v124 = v107
	goto L3
}
func F_bpchar_name(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v79 = F_palloc0(m, int32(64))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L26
	}
L2:
	;
	v59 = v55
	goto L22
L3:
	;
	if v49 <= int32(0) {
		v74 = v49
		v76 = v51
		goto L1
	} else {
		goto L21
	}
L4:
	;
	return int32(0)
L5:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v10 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v13 = int32(1)
	v14 = v6 + v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	if base.Ui32((v16-v13)&int32(255)) < base.Ui32(int32(3)) {
		v55 = int32(4)
		v57 = v14
		goto L2
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v28 = int32(1)
	v31 = v10 & v28
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	if v16 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v27 = int32(16)
	goto L12
L11:
	;
	v27 = int32(0)
	goto L12
L12:
	;
	v49 = v27
	v51 = v14
	goto L3
L13:
	;
	v32 = v28
	goto L15
L14:
	;
	v32 = int32(4)
	goto L15
L15:
	;
	v33 = v6 + v32
	if v31 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v34 = int32(1)
	v43 = int32(base.Ui32(v10)>>(uint(v34)%32)) - v34
	goto L18
L17:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v43 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
	goto L18
L18:
	;
	if v43 < int32(64) {
		v49 = v43
		v51 = v33
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v47 = F_pg_mbcliplen(m, v33, v43, int32(63))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v49 = v47
	v51 = v33
	goto L3
L21:
	;
	v55 = v49
	v57 = v51
	goto L2
L22:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v57-int32(1)))))
	if v66 != int32(32) {
		v74 = v59
		v76 = v57
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v74 = int32(0)
	v76 = v57
	goto L1
L24:
	;
	v69 = int32(1)
	if v69 < v59 {
		v59 = v59 - v69
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	if v74 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	base.MemoryCopy(m, v79, v76, v74)
	goto L29
L28:
	;
	goto L29
L29:
	;
	return v79
}
func F_bpchar_pattern_lt(m *base.Module, l0 int32) int32 {
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
	return int32(base.Ui32(v151) >> (uint(int32(31)) % 32))
L66:
	;
	goto L65
}
func F_bpchar_sortsupport(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13876(m, l0, int32(1042))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
