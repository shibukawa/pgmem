package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_QTNodeCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	F_check_stack_depth(m)
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v18 != v20 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v11 + int32(16)
	return v152
L4:
	;
	if base.I32_extend8_s(v20) < base.I32_extend8_s(v18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	switch v18 - int32(1) {
	case 0:
		goto L12
	case 1:
		goto L13
	default:
		goto L11
	}
L7:
	;
	v27 = int32(-1)
	goto L9
L8:
	;
	v27 = int32(1)
	goto L9
L9:
	;
	v152 = v27
	goto L3
L10:
	;
	if v40 < v39 {
		goto L71
	} else {
		goto L72
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L68
	}
L12:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v90 != v91 {
		goto L34
	} else {
		goto L35
	}
L13:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	v31 = base.I32_extend8_s(v30)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v30 != v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if base.I32_extend8_s(v32) < v31 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v39 != v40 {
		goto L10
	} else {
		goto L20
	}
L17:
	;
	v38 = int32(-1)
	goto L19
L18:
	;
	v38 = int32(1)
	goto L19
L19:
	;
	v152 = v38
	goto L3
L20:
	;
	if int32(0) < v39 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v48 = int32(0)
	goto L24
L22:
	;
	v71 = v31
	goto L23
L23:
	;
	v76 = int32(0)
	if v71&int32(255) != int32(4) {
		v152 = v76
		goto L3
	} else {
		goto L29
	}
L24:
	;
	v54 = v48 << (uint(int32(2)) % 32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54+v55)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v58+v54)))
	v61 = F_QTNodeCompare(m, v57, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	v71 = v67
	goto L23
L26:
	;
	if v61 != 0 {
		v152 = v61
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v64 = v48 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v64 < v65 {
		v48 = v64
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+2)))
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+2)))
	if v81 == v82 {
		v152 = v76
		goto L3
	} else {
		goto L30
	}
L30:
	;
	if base.I32_extend16_s(v82) < base.I32_extend16_s(v81) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v89 = int32(-1)
	goto L33
L32:
	;
	v89 = int32(1)
	goto L33
L33:
	;
	v152 = v89
	goto L3
L34:
	;
	if v91 < v90 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v99 = int32(4095)
	v100 = v98 & v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v104 = v102 & v99
	if v100 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v96 = int32(-1)
	goto L39
L38:
	;
	v96 = int32(1)
	goto L39
L39:
	;
	v152 = v96
	goto L3
L40:
	;
	v152 = v130
	goto L3
L41:
	;
	goto L45
L42:
	;
	goto L43
L43:
	;
	if v104 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	goto L46
L46:
	;
	v110 = int32(0)
	if v110 < v104 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v113 = int32(-1)
	goto L49
L48:
	;
	v113 = v110
	goto L49
L49:
	;
	v130 = v113
	goto L40
L50:
	;
	v130 = base.B2i32(int32(0) < v100)
	goto L40
L51:
	;
	goto L52
L52:
	;
	if base.Ui32(v100) < base.Ui32(v104) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v119 = v100
	goto L55
L54:
	;
	v119 = v104
	goto L55
L55:
	;
	v120 = F_memcmp(m, v97, v101, v119)
	mBase = m.M
	goto L58
L56:
	;
	v130 = v128
	goto L40
L58:
	;
	goto L59
L59:
	;
	if v120 != 0 {
		v128 = v120
		goto L56
	} else {
		goto L61
	}
L61:
	;
	if v100 == v104 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v130 = int32(0)
	goto L40
L63:
	;
	goto L64
L64:
	;
	if v100 < v104 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v127 = int32(-1)
	goto L67
L66:
	;
	v127 = int32(1)
	goto L67
L67:
	;
	v128 = v127
	goto L56
L68:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v136 = int32(*(*int8)(unsafe.Add(mBase, uint32(v135))))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v136
	F_errmsg_internal(m, int32(_a_F_QTNodeCompare_0), v11)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_QTNodeCompare_1), int32(144), int32(_a_F_QTNodeCompare_2))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v149 = int32(-1)
	goto L73
L72:
	;
	v149 = int32(1)
	goto L73
L73:
	;
	v152 = v149
	goto L3
}
func F_qtext_load_file(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v118 int64
	_ = v118
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(176)
	m.G0 = v11
	v15 = F_OpenTransientFile(m, int32(_a_F_qtext_load_file_0), v2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(176)
	return v199
L2:
	;
	return int32(0)
L3:
	;
	if v15 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_qtext_load_file[0]))
	if v22 == int32(44) {
		v199 = v2
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v15 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v27 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	if v27 == int32(0) {
		v199 = v2
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_qtext_load_file_0)
	F_errmsg(m, int32(_a_F_qtext_load_file_1), v11)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_qtext_load_file_2), int32(2324), int32(_a_F_qtext_load_file_3))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v199 = v2
	goto L1
L13:
	;
	if v52 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v48 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v52 = v48
	goto L13
L15:
	;
	goto L16
L16:
	;
	v51 = F___fstatat(m, v15, int32(_a_F_qtext_load_file_4), v11+int32(80), int32(_a_F_qtext_load_file_5))
	mBase = m.M
	v52 = v51
	goto L13
L17:
	;
	v55 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v11)+104))
	if int64(2147483647) < v73 {
		goto L30
	} else {
		goto L31
	}
L20:
	;
	if v55 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v71 = F_CloseTransientFile(m, v15)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L27
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = int32(_a_F_qtext_load_file_0)
	F_errmsg(m, int32(_a_F_qtext_load_file_6), v11-int32(-64))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_qtext_load_file_2), int32(2334), int32(_a_F_qtext_load_file_3))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	v199 = v2
	goto L1
L28:
	;
	v171 = F_CloseTransientFile(m, v15)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L2
	} else {
		goto L62
	}
L29:
	;
	v115 = v2
	v118 = int64(0)
	goto L44
L30:
	;
	v85 = int32(0)
	v88 = F_errstart(m, int32(15), v85)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L35
	}
L31:
	;
	v78 = F_palloc_extended(m, base.I32_wrap_i64(v73), int32(3))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	if v78 == int32(0) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	if v73 <= int64(0) {
		v167 = v2
		goto L28
	} else {
		goto L34
	}
L34:
	;
	goto L29
L35:
	;
	if v88 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_errcode(m, int32(_a_F_qtext_load_file_7))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v109 = F_CloseTransientFile(m, v15)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L43
	}
L39:
	;
	F_errmsg(m, int32(_a_F_qtext_load_file_8), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_qtext_load_file_0)
	v102 = F_errdetail(m, int32(_a_F_qtext_load_file_9), v11+int32(16))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_qtext_load_file_2), int32(2350), int32(_a_F_qtext_load_file_3))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	goto L38
L43:
	;
	v199 = v85
	goto L1
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_qtext_load_file[0])) = int32(0)
	v123 = int64(1073741824)
	v124 = v73 - v118
	if v123 <= v124 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_qtext_load_file[0]))
	if v135 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L46:
	;
	goto L45
L47:
	;
	v127 = v123
	goto L49
L48:
	;
	v127 = v124
	goto L49
L49:
	;
	v128 = base.I32_wrap_i64(v127)
	v129 = F_read(m, v15, v78+v115, v128)
	mBase = m.M
	if v129 != v128 {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v131 = v115 + v128
	v132 = base.I64_extend_i32_u(v131)
	if base.Ui64(v132) < base.Ui64(v73) {
		v115 = v131
		v118 = v132
		goto L44
	} else {
		goto L51
	}
L51:
	;
	v167 = v131
	goto L28
L52:
	;
	F_pfree(m, v78)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L2
	} else {
		goto L59
	}
L53:
	;
	v140 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	if v140 == int32(0) {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(_a_F_qtext_load_file_0)
	F_errmsg(m, int32(_a_F_qtext_load_file_1), v11+int32(48))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_qtext_load_file_2), int32(2379), int32(_a_F_qtext_load_file_3))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	goto L52
L59:
	;
	v160 = F_CloseTransientFile(m, v15)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v199 = int32(0)
	goto L1
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v167
	v199 = v78
	goto L1
L62:
	;
	if v171 == int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v177 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	if v177 == int32(0) {
		goto L61
	} else {
		goto L65
	}
L65:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(_a_F_qtext_load_file_0)
	F_errmsg(m, int32(_a_F_qtext_load_file_10), v11+int32(32))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_qtext_load_file_2), int32(2390), int32(_a_F_qtext_load_file_3))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	goto L61
}
