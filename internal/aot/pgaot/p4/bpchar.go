package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bpchar(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v18 <= int32(3) {
		v132 = v14
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v11 + int32(16)
	return base.I64_extend_i32_u(v132)
L4:
	;
	v22 = v18 - int32(4)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v23 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v54 = int32(1)
	if v23&v54 != 0 {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v29 == int32(18) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v40 = int32(1)
	if v23&v40 != 0 {
		v52 = int32(base.Ui32(v23)>>(uint(v40)%32)) - v40
		goto L5
	} else {
		goto L15
	}
L9:
	;
	v32 = int32(16)
	goto L11
L10:
	;
	v32 = int32(0)
	goto L11
L11:
	;
	if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v39 = int32(4)
	goto L14
L13:
	;
	v39 = v32
	goto L14
L14:
	;
	v52 = v39
	goto L5
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
	goto L5
L16:
	;
	v58 = v54
	goto L18
L17:
	;
	v58 = int32(4)
	goto L18
L18:
	;
	v59 = v14 + v58
	v60 = F_pg_mbstrlen_with_len(m, v59, v52)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v60 == v22 {
		v132 = v14
		goto L3
	} else {
		goto L20
	}
L20:
	;
	if v22 < v60 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v113 = v105 + int32(4)
	v114 = F_palloc(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L42
	}
L22:
	;
	v64 = F_pg_mbcharcliplen(m, v59, v52, v22)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v105 = v52 + v22 - v60
	v106 = v52
	goto L21
L25:
	;
	if v53 != int64(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v105 = v64
	v106 = v64
	goto L21
L27:
	;
	goto L28
L28:
	;
	if v52 <= v64 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v105 = v64
	v106 = v64
	goto L21
L30:
	;
	goto L31
L31:
	;
	v72 = v64
	goto L33
L32:
	;
	v84 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v86 = F_errsave_start(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L37
	}
L33:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v59))))
	if v78 != int32(32) {
		goto L32
	} else {
		goto L35
	}
L34:
	;
	v105 = v64
	v106 = v64
	goto L21
L35:
	;
	v82 = v72 + int32(1)
	if v82 != v52 {
		v72 = v82
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	if v86 == int32(0) {
		v132 = v84
		goto L3
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v22
	F_errmsg(m, int32(_a_F_bpchar_0), v11)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errsave_finish(m, v85, int32(_a_F_bpchar_1), int32(313), int32(_a_F_bpchar_2))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v132 = v84
	goto L3
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v113 << (uint(int32(2)) % 32)
	v120 = v114 + int32(4)
	if v106 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	base.MemoryCopy(m, v120, v59, v106)
	goto L45
L44:
	;
	goto L45
L45:
	;
	if v105 <= v106 {
		v132 = v114
		goto L3
	} else {
		goto L46
	}
L46:
	;
	v123 = v105 - v106
	if v123 == int32(0) {
		v132 = v114
		goto L3
	} else {
		goto L47
	}
L47:
	;
	base.MemoryFill(m, v120+v106, int32(32), v123)
	v132 = v114
	goto L3
}
func F_bpchar_name(m *base.Module, l0 int32) int64 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	return base.I64_extend_i32_u(v79)
}
func F_bpchar_pattern_lt(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
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
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v17 = int32(1)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v21 = v19 & v17
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v129 != v6 {
		goto L53
	} else {
		goto L54
	}
L5:
	;
	v22 = v17
	goto L7
L6:
	;
	v22 = int32(4)
	goto L7
L7:
	;
	v23 = v22 + v6
	if v19 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v56 = v50
	goto L19
L9:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	if v29 == int32(18) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v40 = int32(1)
	if v21 != 0 {
		v50 = int32(base.Ui32(v19)>>(uint(v40)%32)) - v40
		goto L8
	} else {
		goto L18
	}
L12:
	;
	v32 = int32(16)
	goto L14
L13:
	;
	v32 = int32(0)
	goto L14
L14:
	;
	if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v39 = int32(4)
	goto L17
L16:
	;
	v39 = v32
	goto L17
L17:
	;
	v50 = v39
	goto L8
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L19:
	;
	if v56 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v70 = int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v74 = v72 & v70
	if v74 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	goto L20
L22:
	;
	v68 = v50 & (v50 >> (uint(int32(31)) % 32))
	goto L21
L23:
	;
	goto L24
L24:
	;
	v63 = v56 - int32(1)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v63))))
	if v65 == int32(32) {
		v56 = v63
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v68 = v56
	goto L21
L26:
	;
	v75 = v70
	goto L28
L27:
	;
	v75 = int32(4)
	goto L28
L28:
	;
	v76 = v75 + v11
	if v72 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v109 = v103
	goto L40
L30:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v82 == int32(18) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v93 = int32(1)
	if v74 != 0 {
		v103 = int32(base.Ui32(v72)>>(uint(v93)%32)) - v93
		goto L29
	} else {
		goto L39
	}
L33:
	;
	v85 = int32(16)
	goto L35
L34:
	;
	v85 = int32(0)
	goto L35
L35:
	;
	if base.Ui32((v82-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v92 = int32(4)
	goto L38
L37:
	;
	v92 = v85
	goto L38
L38:
	;
	v103 = v92
	goto L29
L39:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v103 = int32(base.Ui32(v97)>>(uint(int32(2))%32)) - int32(4)
	goto L29
L40:
	;
	if v109 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v123 = base.B2i32(v68 < v121)
	if v68 < v121 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	v121 = v103 & (v103 >> (uint(int32(31)) % 32))
	goto L42
L44:
	;
	goto L45
L45:
	;
	v116 = v109 - int32(1)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v116))))
	if v118 == int32(32) {
		v109 = v116
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v121 = v109
	goto L42
L47:
	;
	goto L4
L48:
	;
	v124 = v68
	goto L50
L49:
	;
	v124 = v121
	goto L50
L50:
	;
	v125 = F_memcmp(m, v23, v76, v124)
	mBase = m.M
	if v125 != 0 {
		v128 = v125
		goto L47
	} else {
		goto L51
	}
L51:
	;
	if v68 < v121 {
		v128 = int32(-1)
		goto L47
	} else {
		goto L52
	}
L52:
	;
	v128 = base.B2i32(v121 < v68)
	goto L47
L53:
	;
	F_pfree(m, v6)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v133 != v11 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v11)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	return base.I64_extend_i32_u(int32(base.Ui32(v128) >> (uint(int32(31)) % 32)))
L60:
	;
	goto L59
}
func F_bpchar_sortsupport(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14238(m, l0, int32(1042))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
