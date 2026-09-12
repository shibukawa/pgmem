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
	F_sequence_close(m, v47, int32(0))
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
	F_sequence_close(m, v89, int32(0))
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
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
	v21 = int32(49152)
	if v16&v21 == v21 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L14
	} else {
		goto L70
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L14
	} else {
		goto L67
	}
L6:
	;
	m.G0 = v14 + int32(16)
	return v203
L7:
	;
	if v16 == int32(49152) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v39 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v32 = F_palloc(m, int32(6))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	if v16 == int32(53248) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if v16 != int32(61440) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	return int32(0)
L15:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+4)) = uint16(v16)
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(24)
	v203 = v32
	goto L6
L16:
	;
	if v98 != 0 {
		goto L30
	} else {
		goto L31
	}
L17:
	;
	v98 = v39
	v100 = v17
	v102 = v18
	v103 = v3
	goto L16
L18:
	;
	goto L19
L19:
	;
	v48 = v18
	v49 = v17
	v51 = v39
	goto L21
L20:
	;
	v74 = v51
	goto L26
L21:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48))))
	if v57 != 0 {
		goto L20
	} else {
		goto L23
	}
L22:
	;
	v98 = int32(0)
	v100 = v17 - v39
	v102 = v18 + v39<<(uint(int32(1))%32)
	v103 = v3
	goto L16
L23:
	;
	v58 = int32(1)
	if v58 < v51 {
		v48 = v48 + int32(2)
		v49 = v49 - v58
		v51 = v51 - v58
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v98 = v91
	v100 = v49
	v102 = v48
	v103 = base.B2i32(v83 != int32(0))
	goto L16
L26:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48-int32(2)+v74<<(uint(int32(1))%32)))))
	if v83 != 0 {
		v91 = v74
		goto L25
	} else {
		goto L28
	}
L27:
	;
	v91 = int32(0)
	goto L25
L28:
	;
	v86 = int32(1)
	if v86 < v74 {
		v74 = v74 - v86
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v105 = v16
	goto L32
L31:
	;
	v105 = int32(0)
	goto L32
L32:
	;
	if v98 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v107 = v100
	goto L35
L34:
	;
	v107 = int32(0)
	goto L35
L35:
	;
	v109 = v98 << (uint(int32(1)) % 32)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(63) < v110 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v153)+4)) = uint16(v155)
	if v103 != 0 {
		goto L46
	} else {
		goto L47
	}
L37:
	;
	v142 = v109 + int32(8)
	v143 = F_palloc(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L14
	} else {
		goto L44
	}
L38:
	;
	if base.Ui32(int32(127)) < base.Ui32(v107-int32(-64)) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v118 = v109 + int32(6)
	v119 = F_palloc(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v118 << (uint(int32(2)) % 32)
	if v105 == int32(16384) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v134 = int32(40960)
	goto L43
L42:
	;
	v134 = int32(32768)
	goto L43
L43:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v153 = v119
	v155 = int32(base.Ui32(v107)>>(uint(int32(25))%32))&int32(64) | (v107&int32(63) | v134) | v137<<(uint(int32(7))%32)
	goto L36
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v142 << (uint(int32(2)) % 32)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*uint16)(unsafe.Add(mBase, uint32(v143)+6)) = uint16(v107)
	v153 = v143
	v155 = v148&int32(16383) | v105
	goto L36
L45:
	;
	if v183 == v107 {
		goto L59
	} else {
		goto L60
	}
L46:
	;
	if base.I32_extend16_s(v155) < int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v169 = v155
	goto L48
L48:
	;
	v170 = base.I32_extend16_s(v169)
	if v170 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	v162 = int32(6)
	goto L51
L50:
	;
	v162 = int32(8)
	goto L51
L51:
	;
	v165 = v98 << (uint(int32(1)) % 32)
	if v165 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+4)))
	v169 = v168
	goto L48
L53:
	;
	v166 = F__emscripten_memcpy_bulkmem(m, v153+v162, v102, v165)
	mBase = m.M
	goto L55
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	v183 = v169<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v169&int32(63)
	goto L45
L57:
	;
	goto L58
L58:
	;
	v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(v153)+6)))
	v183 = v182
	goto L45
L59:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) <= v170 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L66
	}
L62:
	;
	v194 = v169 & int32(16383)
	goto L64
L63:
	;
	v194 = int32(base.Ui32(v169)>>(uint(int32(7))%32)) & int32(63)
	goto L64
L64:
	;
	if v185 == v194 {
		v203 = v153
		goto L6
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	v198 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v198)
	v203 = int32(0)
	goto L6
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v16
	F_errmsg_internal(m, int32(29782), v14)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L14
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(499737), int32(7923), int32(211507))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L14
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L14
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(112152), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L14
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(499737), int32(7995), int32(211507))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L14
	} else {
		goto L73
	}
L73:
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
	var v44 int32
	_ = v44
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
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
	v44 = v4
	v46 = int32(-1)
	goto L8
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v44<<(uint(int32(2))%32))))
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
	if v46 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v60 = v46
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
	v58 = v46
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
	v76 = v44 + int32(1)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v76 < v77 {
		v44 = v76
		v46 = v60
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
