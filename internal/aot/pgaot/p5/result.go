package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecCloseResultRelations(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v7 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v71 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v10 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = v2
	goto L4
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19+v17<<(uint(int32(2))%32))))
	F_ExecCloseIndices(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+212))
	if v26 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v62 = v17 + int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v62 < v63 {
		v17 = v62
		goto L4
	} else {
		goto L18
	}
L9:
	;
	v29 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v30 <= v29 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v34 = v29
	goto L11
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v34<<(uint(int32(2))%32))))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v44 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L8
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	F_relation_close(m, v47, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v52 = v34 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v52 < v53 {
		v34 = v52
		goto L11
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	goto L12
L18:
	;
	goto L5
L19:
	;
	return
L20:
	;
	v74 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v75 <= v74 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v79 = v74
	goto L22
L22:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84+v79<<(uint(int32(2))%32))))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	F_relation_close(m, v89, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L24
	}
L23:
	;
	goto L19
L24:
	;
	v94 = v79 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v94 < v95 {
		v79 = v94
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
}
func F_make_result_opt_error(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v19)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v21 = int32(_a_F_make_result_opt_error_0)
	if v16&v21 == v21 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L11
	} else {
		goto L65
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L11
	} else {
		goto L62
	}
L6:
	;
	m.G0 = v14 + int32(16)
	return v209
L7:
	;
	if base.B2i32(base.B2i32(v16 == int32(_a_F_make_result_opt_error_0))|base.B2i32(v16 == int32(_a_F_make_result_opt_error_1)) == int32(0))&base.B2i32(v16 != int32(_a_F_make_result_opt_error_2)) != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v43 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v36 = F_palloc(m, int32(6))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)) = uint16(v16)
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(24)
	v209 = v36
	goto L6
L13:
	;
	if v100 != 0 {
		goto L28
	} else {
		goto L29
	}
L14:
	;
	v100 = v43
	v102 = v17
	v103 = v18
	v105 = v3
	goto L13
L15:
	;
	goto L16
L16:
	;
	v52 = v18
	v53 = v17
	v54 = v43
	goto L17
L17:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52))))
	if v61 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v100 = int32(0)
	v102 = v17 - v43
	v103 = v18 + v43<<(uint(int32(1))%32)
	v105 = v3
	goto L13
L19:
	;
	v66 = v54
	goto L23
L20:
	;
	goto L21
L21:
	;
	v87 = int32(1)
	if v87 < v54 {
		v52 = v52 + int32(2)
		v53 = v53 - v87
		v54 = v54 - v87
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v100 = v86
	v102 = v53
	v103 = v52
	v105 = base.B2i32(v78 != int32(0))
	goto L13
L23:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+v66<<(uint(int32(1))%32)-int32(2)))))
	if v78 != 0 {
		v86 = v66
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v86 = int32(0)
	goto L22
L25:
	;
	v81 = int32(1)
	if v81 < v66 {
		v66 = v66 - v81
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	goto L18
L28:
	;
	v108 = v16
	goto L30
L29:
	;
	v108 = int32(0)
	goto L30
L30:
	;
	v110 = v100 << (uint(int32(1)) % 32)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v100 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v159)+4)) = uint16(v161)
	if v105 != 0 {
		goto L44
	} else {
		goto L45
	}
L32:
	;
	v115 = v102
	goto L34
L33:
	;
	v115 = int32(0)
	goto L34
L34:
	;
	if base.B2i32(int32(63) < v111)|base.B2i32(base.Ui32(int32(127)) < base.Ui32(v115-int32(-64))) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v124 = v110 + int32(6)
	v125 = F_palloc(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L11
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v148 = v110 + int32(8)
	v149 = F_palloc(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L11
	} else {
		goto L42
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v124 << (uint(int32(2)) % 32)
	if v108 == int32(_a_F_make_result_opt_error_3) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v140 = int32(_a_F_make_result_opt_error_4)
	goto L41
L40:
	;
	v140 = int32(_a_F_make_result_opt_error_5)
	goto L41
L41:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v159 = v125
	v161 = int32(base.Ui32(v115)>>(uint(int32(25))%32))&int32(64) | (v115&int32(63) | v140) | v143<<(uint(int32(7))%32)
	goto L31
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v148 << (uint(int32(2)) % 32)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*uint16)(unsafe.Add(mBase, uint32(v149)+6)) = uint16(v115)
	v159 = v149
	v161 = v154&int32(_a_F_make_result_opt_error_6) | v108
	goto L31
L43:
	;
	if v189 == v115 {
		goto L54
	} else {
		goto L55
	}
L44:
	;
	v164 = v100 << (uint(int32(1)) % 32)
	if v164 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v174 = v161
	goto L46
L46:
	;
	v185 = base.I32_extend16_s(v174)
	if v185 < int32(0) {
		v189 = v174<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v174&int32(63)
		goto L43
	} else {
		goto L53
	}
L47:
	;
	if base.I32_extend16_s(v161) < int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v159)+4)))
	v174 = v173
	goto L46
L50:
	;
	v170 = int32(6)
	goto L52
L51:
	;
	v170 = int32(8)
	goto L52
L52:
	;
	base.MemoryCopy(m, v159+v170, v103, v164)
	goto L49
L53:
	;
	v188 = int32(*(*int16)(unsafe.Add(mBase, uint32(v159)+6)))
	v189 = v188
	goto L43
L54:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) <= v185 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L61
	}
L57:
	;
	v200 = v174 & int32(_a_F_make_result_opt_error_6)
	goto L59
L58:
	;
	v200 = int32(base.Ui32(v174)>>(uint(int32(7))%32)) & int32(63)
	goto L59
L59:
	;
	if v191 == v200 {
		v209 = v159
		goto L6
	} else {
		goto L60
	}
L60:
	;
	goto L56
L61:
	;
	v204 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v204)
	v209 = int32(0)
	goto L6
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v16
	F_errmsg_internal(m, int32(_a_F_make_result_opt_error_7), v14)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_make_result_opt_error_8), int32(_a_F_make_result_opt_error_9), int32(_a_F_make_result_opt_error_10))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(_a_F_make_result_opt_error_11), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_make_result_opt_error_8), int32(_a_F_make_result_opt_error_12), int32(_a_F_make_result_opt_error_10))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_remove_result_refs(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	if v14 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	v19 = F_get_relids_in_jointree(m, l2, int32(1), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v19
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
	v30 = F_query_or_expression_tree_walker_impl(m, v21, int32(852), v11+int32(4), v23)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v32 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v35 <= int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v45 = int32(-1)
	v46 = v4
	goto L8
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v46<<(uint(int32(2))%32))))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if l1 == v52 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L1
L10:
	;
	if v45 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v60 = v45
	goto L12
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+68))
	if v62 != 0 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v56 = F_bms_singleton_member(m, v19)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L3
	} else {
		goto L16
	}
L14:
	;
	v58 = v45
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v58
	v60 = v58
	goto L12
L16:
	;
	v58 = v56
	goto L15
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v19
	v65 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
	v72 = F_query_or_expression_tree_walker_impl(m, v63, int32(852), v11+int32(4), v65)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v75 = v46 + int32(1)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v75 < v76 {
		v45 = v60
		v46 = v75
		goto L8
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	goto L9
}
