package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_ChooseRelationName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	v16 = m.G0
	v18 = v16 - int32(272)
	m.G0 = v18
	*(*int32)(unsafe.Add(mBase, uint32(v18)+136)) = int32(4)
	v24 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = v18 + int32(208)
	goto L6
L3:
	;
	v165 = int32(0)
	goto L34
L4:
	;
	v146 = F_strlen(m, v135)
	mBase = m.M
	goto L3
L6:
	;
	goto L7
L7:
	;
	v36 = int32(63)
	if (v29^l2)&int32(3) != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v139 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v139)
	goto L4
L9:
	;
	v120 = v115
	v121 = v116
	v122 = v117
	goto L30
L10:
	;
	if v110 == int32(0) {
		v135 = v108
		v136 = v109
		goto L8
	} else {
		goto L29
	}
L11:
	;
	v108 = l2
	v109 = v29
	v110 = v36
	goto L10
L12:
	;
	goto L13
L13:
	;
	v40 = int32(0)
	if base.B2i32(l2&int32(3) == v40)|int32(0) == v40 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v76 == int32(0) {
		v135 = v73
		v136 = v74
		goto L8
	} else {
		goto L23
	}
L15:
	;
	v52 = l2
	v53 = v29
	v54 = v36
	goto L18
L16:
	;
	goto L17
L17:
	;
	v73 = l2
	v74 = v29
	v75 = v36
	v76 = int32(1)
	goto L14
L18:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v56)
	if v56 == int32(0) {
		v115 = v52
		v116 = v53
		v117 = v54
		goto L9
	} else {
		goto L20
	}
L19:
	;
	v73 = v67
	v74 = v61
	v75 = v63
	v76 = v65
	goto L14
L20:
	;
	v60 = int32(1)
	v61 = v53 + v60
	v63 = v54 - v60
	v64 = int32(0)
	v65 = base.B2i32(v63 != v64)
	v67 = v52 + v60
	if v67&int32(3) == v64 {
		v73 = v67
		v74 = v61
		v75 = v63
		v76 = v65
		goto L14
	} else {
		goto L21
	}
L21:
	;
	if v63 != 0 {
		v52 = v67
		v53 = v61
		v54 = v63
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if base.B2i32(v79 == int32(0))|base.B2i32(base.Ui32(v75) < base.Ui32(int32(4))) != 0 {
		v108 = v73
		v109 = v74
		v110 = v75
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v86 = v73
	v87 = v74
	v88 = v75
	goto L25
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v94 = int32(-2139062144)
	if (int32(16843008)-v91|v91)&v94 != v94 {
		v115 = v86
		v116 = v87
		v117 = v88
		goto L9
	} else {
		goto L27
	}
L26:
	;
	v108 = v102
	v109 = v100
	v110 = v104
	goto L10
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v91
	v99 = int32(4)
	v100 = v87 + v99
	v102 = v86 + v99
	v104 = v88 - v99
	if base.Ui32(int32(3)) < base.Ui32(v104) {
		v86 = v102
		v87 = v100
		v88 = v104
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v115 = v108
	v116 = v109
	v117 = v110
	goto L9
L30:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
	if v124 == int32(0) {
		v135 = v120
		v136 = v121
		goto L8
	} else {
		goto L32
	}
L31:
	;
	v135 = v131
	v136 = v129
	goto L8
L32:
	;
	v128 = int32(1)
	v129 = v121 + v128
	v131 = v120 + v128
	v133 = v122 - v128
	if v133 != 0 {
		v120 = v131
		v121 = v129
		v122 = v133
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v168 = v18 + int32(16)
	v174 = F_makeObjectName(m, l0, l1, v18+int32(208))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L57
	}
L36:
	;
	goto L35
L37:
	;
	F_ScanKeyInit(m, v168, int32(2), int32(3), int32(62), base.I64_extend_i32_u(v174))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v179 = int32(3)
	F_ScanKeyInit(m, v18+int32(72), v179, v179, int32(184), base.I64_extend_i32_u(l3))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v189 = F_systable_beginscan(m, v24, int32(2663), int32(1), v18+int32(136), int32(2), v168)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v191 = F_systable_getnext(m, v189)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_systable_endscan(m, v189)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v191 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if l4 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_pfree(m, v174)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L55
	}
L46:
	;
	v199 = m.G0
	v201 = v199 - int32(112)
	m.G0 = v201
	v205 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_ScanKeyInit(m, v201, int32(2), int32(3), int32(62), base.I64_extend_i32_u(v174))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v215 = int32(3)
	F_ScanKeyInit(m, v201+int32(56), v215, v215, int32(184), base.I64_extend_i32_u(l3))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v225 = F_systable_beginscan(m, v205, int32(2664), int32(1), int32(0), int32(2), v201)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v227 = F_systable_getnext(m, v225)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_systable_endscan(m, v225)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_relation_close(m, v205, int32(1))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	m.G0 = v201 + int32(112)
	if v227 == int32(0) {
		goto L36
	} else {
		goto L54
	}
L54:
	;
	goto L45
L55:
	;
	v245 = v165 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l2
	v252 = F_pg_snprintf(m, v18+int32(208), int32(64), int32(_a_F_ChooseRelationName_0), v18)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v165 = v245
	goto L34
L57:
	;
	m.G0 = v18 + int32(272)
	return v174
}
func F_CopyRelationTo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	v5 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_CopyRelationTo[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	goto L1
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_CopyRelationTo[1]))
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L7
	} else {
		goto L54
	}
L3:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopyRelationTo[2])))
	if v17&int32(1) == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v22 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+188))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v28 = m.T0[v27].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l1, v13, v22, v22, v22, int32(449))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	return
L8:
	;
	v31 = F_table_slot_create(m, l1, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if l2 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v34 = F_table_slot_create(m, l2, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v41 = v5
	v42 = v5
	goto L12
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+40)) = v44
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+188))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v50 = m.T0[v49].(func(*base.Module, int32, int32, int32) int32)(m, v28, int32(1), v31)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v39 = F_build_attrmap_by_name_if_req(m, v36, v37, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v41 = v39
	v42 = v34
	goto L12
L15:
	;
	if v50 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	goto L19
L17:
	;
	goto L18
L18:
	;
	F_ExecDropSingleTupleTableSlot(m, v31)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L7
	} else {
		goto L44
	}
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_CopyRelationTo[3]))
	if v63 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v41 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L23
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	F_MemoryContextReset(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L7
	} else {
		goto L32
	}
L26:
	;
	v66 = F_execute_attr_map_slot(m, v41, v31, v42)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31)+6)))
	if v69 <= v70 {
		v77 = v31
		goto L25
	} else {
		goto L30
	}
L29:
	;
	v77 = v66
	goto L25
L30:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	m.T0[v73].(func(*base.Module, int32, int32))(m, v31, v69)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	v77 = v31
	goto L25
L32:
	;
	v81 = int32(_a_F_CopyRelationTo_0)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_CopyRelationTo[4]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyRelationTo[4])) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v77)+6)))
	if v88 < v87 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	m.T0[v91].(func(*base.Module, int32, int32))(m, v77, v87)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	m.T0[v95].(func(*base.Module, int32, int32))(m, l0, v77)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CopyRelationTo[4])) = v82
	v100 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	v102 = v100 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v102
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_CopyRelationTo[5]))
	if v107 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+40)) = v149
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+188))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+20))
	v155 = m.T0[v154].(func(*base.Module, int32, int32, int32) int32)(m, v28, int32(1), v31)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L7
	} else {
		goto L42
	}
L39:
	;
	goto L38
L40:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopyRelationTo[6])))
	if v111&int32(1) == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v116 = int32(_a_F_CopyRelationTo_1)
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_CopyRelationTo[7]))
	v119 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CopyRelationTo[7])) = v118 + v119
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v122 + v119
	v126 = int32(0)
	v128 = int32(_a_F_CopyRelationTo_2)
	v129 = base.AtomicRmwOr32(m, v126, v128, v126)
	*(*int64)(unsafe.Add(mBase, uint32(v107+int32(16))+232)) = v102
	v137 = base.AtomicRmwOr32(m, v126, v128, v126)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v138 + v119
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_CopyRelationTo[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyRelationTo[7])) = v144 - v119
	goto L39
L42:
	;
	if v155 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	goto L20
L44:
	;
	if v42 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_ExecDropSingleTupleTableSlot(m, v42)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L7
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if v41 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	F_free_attrmap(m, v41)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L7
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+188))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	m.T0[v175].(func(*base.Module, int32))(m, v28)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L7
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	return
L54:
	;
	F_errmsg_internal(m, int32(_a_F_CopyRelationTo_3), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_CopyRelationTo_4), int32(931), int32(_a_F_CopyRelationTo_5))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_DropRelationFiles(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v14 = F_palloc_mul(m, int32(4), l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if int32(0) < l1 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_pfree(m, v14)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L24
	}
L4:
	;
	v24 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_smgrdounlinkall(m, v14, l1, l2)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L23
	}
L7:
	;
	v28 = l0 + v24*int32(12)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v29
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = v31
	v36 = F_smgropen(m, v11-int32(-64), int32(-1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	F_smgrdounlinkall(m, v14, l1, l2)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L18
	}
L9:
	;
	if l2 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v38
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v40
	F_XLogDropRelation(m, v11+int32(48), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+v24<<(uint(int32(2))%32)))) = v36
	v77 = v24 + int32(1)
	if v77 != l1 {
		v24 = v77
		goto L7
	} else {
		goto L17
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v47
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v49
	F_XLogDropRelation(m, v11+int32(32), int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v58
	F_XLogDropRelation(m, v11+int32(16), int32(2))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v65
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v67
	F_XLogDropRelation(m, v11, int32(3))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	goto L8
L18:
	;
	v82 = int32(0)
	goto L19
L19:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v14+v82<<(uint(int32(2))%32))))
	F_smgrclose(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L3
L21:
	;
	v97 = v82 + int32(1)
	if v97 != l1 {
		v82 = v97
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L3
L24:
	;
	m.G0 = v11 + int32(80)
	return
}
func F_LockRelationId(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = v9
	v21 = F_LockAcquireExtended(m, v5+int32(16), int32(1), v2, v2, v5+int32(12), v2)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		if v21 != int32(3) {
			F_ReceiveSharedInvalidMessages(m)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				v28 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+53)) = uint8(v28)
				m.G0 = v5 + int32(32)
				return
			}
		} else {
			m.G0 = v5 + int32(32)
			return
		}
	}
}
func F_RelationCacheInitFilePreInvalidate(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v3 = m.G0
	v5 = v3 - int32(2080)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitFilePreInvalidate[0]))
	if v8 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = int32(_a_F_RelationCacheInitFilePreInvalidate_0)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v8
		v18 = F_pg_snprintf(m, v5+int32(1056), int32(1024), int32(_a_F_RelationCacheInitFilePreInvalidate_1), v5+int32(16))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_RelationCacheInitFilePreInvalidate_0)
			v26 = F_pg_snprintf(m, v5+int32(32), int32(1024), int32(_a_F_RelationCacheInitFilePreInvalidate_2), v5)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitFilePreInvalidate[1]))
				v33 = F_LWLockAcquire(m, v29+int32(2048), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitFilePreInvalidate[0]))
					if v36 != 0 {
						F_unlink_initfile(m, v5+int32(1056), int32(21))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							F_unlink_initfile(m, v5+int32(32), int32(21))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								m.G0 = v5 + int32(2080)
								return
							}
						}
					} else {
						F_unlink_initfile(m, v5+int32(32), int32(21))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v5 + int32(2080)
							return
						}
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_RelationCacheInitFilePreInvalidate_0)
		v26 = F_pg_snprintf(m, v5+int32(32), int32(1024), int32(_a_F_RelationCacheInitFilePreInvalidate_2), v5)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitFilePreInvalidate[1]))
			v33 = F_LWLockAcquire(m, v29+int32(2048), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitFilePreInvalidate[0]))
				if v36 != 0 {
					F_unlink_initfile(m, v5+int32(1056), int32(21))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_unlink_initfile(m, v5+int32(32), int32(21))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v5 + int32(2080)
							return
						}
					}
				} else {
					F_unlink_initfile(m, v5+int32(32), int32(21))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						m.G0 = v5 + int32(2080)
						return
					}
				}
			}
		}
	}
}
func F_RelationClose(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = v3 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_RelationClose[0]))
	if v8 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_RelationClose[1]))
		F_ResourceOwnerForget(m, v10, base.I64_extend_i32_u(l0), int32(_a_F_RelationClose_0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v16 = v15
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				if v17 == int32(0) {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
					if v25 == int32(0) {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
						if v28 == int32(0) {
							return
						} else {
							F_MemoryContextDeleteChildren(m, v25)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
					if v20 == int32(0) {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
						if v25 == int32(0) {
							return
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
							if v28 == int32(0) {
								return
							} else {
								F_MemoryContextDeleteChildren(m, v25)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						F_MemoryContextDeleteChildren(m, v17)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
							if v25 == int32(0) {
								return
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
								if v28 == int32(0) {
									return
								} else {
									F_MemoryContextDeleteChildren(m, v25)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v16 = v5
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v17 == int32(0) {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
				if v25 == int32(0) {
					return
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
					if v28 == int32(0) {
						return
					} else {
						F_MemoryContextDeleteChildren(m, v25)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
				if v20 == int32(0) {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
					if v25 == int32(0) {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
						if v28 == int32(0) {
							return
						} else {
							F_MemoryContextDeleteChildren(m, v25)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					F_MemoryContextDeleteChildren(m, v17)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
						if v25 == int32(0) {
							return
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
							if v28 == int32(0) {
								return
							} else {
								F_MemoryContextDeleteChildren(m, v25)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_RelationGetIndexAttrBitmap(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v300 int64
	_ = v300
	var v301 int32
	_ = v301
	var v305 int64
	_ = v305
	var v306 int32
	_ = v306
	var v307 int64
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v467 int32
	_ = v467
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v494 int64
	_ = v494
	var v495 int32
	_ = v495
	var v499 int64
	_ = v499
	var v500 int32
	_ = v500
	var v501 int64
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	v3 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(32)
	m.G0 = v23
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)))
	if v25 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v23 + int32(32)
	return v761
L2:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v754 = F_bms_copy(m, v753)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L11
	} else {
		goto L168
	}
L3:
	;
	switch l1 {
	case 0:
		goto L2
	case 1:
		goto L10
	case 2:
		goto L9
	case 3:
		goto L8
	case 4:
		goto L7
	default:
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+116)))
	if v56 != int32(1) {
		v761 = v3
		goto L1
	} else {
		goto L19
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L16
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v40 = F_bms_copy(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L11
	} else {
		goto L15
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v37 = F_bms_copy(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L11
	} else {
		goto L14
	}
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v34 = F_bms_copy(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L11
	} else {
		goto L13
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v29 = F_bms_copy(m, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v761 = v29
	goto L1
L13:
	;
	v761 = v34
	goto L1
L14:
	;
	v761 = v37
	goto L1
L15:
	;
	v761 = v40
	goto L1
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = l1
	F_errmsg_internal(m, int32(_a_F_RelationGetIndexAttrBitmap_0), v23)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_RelationGetIndexAttrBitmap_1), int32(_a_F_RelationGetIndexAttrBitmap_2), int32(_a_F_RelationGetIndexAttrBitmap_3))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v59 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	if v59 == int32(0) {
		v761 = v3
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v72 = v59
	goto L27
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L11
	} else {
		goto L165
	}
L23:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	v761 = v737
	goto L1
L24:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v761 = v736
	goto L1
L25:
	;
	v761 = v641
	goto L1
L26:
	;
	v761 = v642
	goto L1
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v85 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v85
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v85 < v92 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v761 = int32(0)
	goto L1
L29:
	;
	v101 = v85
	v108 = v85
	v109 = v85
	v110 = int32(0)
	goto L32
L30:
	;
	v634 = v85
	v641 = v85
	v642 = v85
	goto L31
L31:
	;
	v649 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L11
	} else {
		goto L139
	}
L32:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116+v110<<(uint(int32(2))%32))))
	v122 = F_index_open(m, v120, int32(1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L11
	} else {
		goto L34
	}
L33:
	;
	v634 = v601
	v641 = v608
	v642 = v609
	goto L31
L34:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122)+196))
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[0]))
	if v126 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v130 = int32(_a_F_RelationGetIndexAttrBitmap_4)
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1]))
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1])) = v134
	v137 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L11
	} else {
		goto L38
	}
L36:
	;
	v273 = v126
	goto L37
L37:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v291)+18)))
	if base.Ui32(v292&int32(2044)) <= base.Ui32(int32(19)) {
		goto L63
	} else {
		goto L64
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v137)+4)) = int64(-4294965047)
	v144 = int32(0)
	goto L39
L39:
	;
	v161 = int32(100)
	v162 = v144 * v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	base.MemoryCopy(m, v162+(v137+v163<<(uint(int32(3))%32))+int32(28), v162+int32(_a_F_RelationGetIndexAttrBitmap_5), v161)
	F_populate_compact_attribute(m, v137, v144)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L11
	} else {
		goto L41
	}
L40:
	;
	v180 = int32(0)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	if v180 < v189 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v177 = v144 + int32(1)
	if v177 != int32(21) {
		v144 = v177
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[0])) = v137
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1])) = v131
	v273 = v137
	goto L37
L44:
	;
	v193 = v137 + int32(28)
	v200 = v180
	v201 = v189
	v203 = v180
	goto L48
L45:
	;
	v257 = v180
	v264 = v189
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137)+20)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v137)+16)) = v257
	goto L43
L47:
	;
	v257 = v251
	v264 = v230
	goto L46
L48:
	;
	v209 = v193 + v189<<(uint(int32(3))%32) + v200*int32(100)
	v212 = v193 + v200<<(uint(int32(3))%32)
	if v189 != v201 {
		v230 = v201
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v251 = v189
	goto L47
L50:
	;
	v231 = int32(*(*int16)(unsafe.Add(mBase, uint32(v212)+2)))
	if v231 <= int32(0) {
		v251 = v200
		goto L47
	} else {
		goto L58
	}
L51:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+7)))
	if v214 != int32(118) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v230 = v200
	goto L50
L53:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+4)))
	if v217 != int32(1) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+6)))
	if v220&int32(6) != 0 {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v223 = int32(*(*int16)(unsafe.Add(mBase, uint32(v212)+2)))
	if v223 <= int32(0) {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+90)))
	if v226 != int32(118) {
		v230 = v189
		goto L50
	} else {
		goto L57
	}
L57:
	;
	goto L52
L58:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+90)))
	if v234 == int32(118) {
		v251 = v200
		goto L47
	} else {
		goto L59
	}
L59:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+5)))
	v243 = (v203 + v237 - int32(1)) & (int32(0) - v237)
	if int32(_a_F_RelationGetIndexAttrBitmap_6) < v243 {
		v251 = v200
		goto L47
	} else {
		goto L60
	}
L60:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v212))) = uint16(v243)
	v249 = v200 + int32(1)
	if v249 != v189 {
		v200 = v249
		v201 = v230
		v203 = v243 + v231
		goto L48
	} else {
		goto L61
	}
L61:
	;
	goto L49
L62:
	;
	v308 = int32(0)
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+23)))
	if v309 == v308 {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	v300 = F_getmissingattr(m, v273, int32(20), v23+int32(23))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L11
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v305 = F_fastgetattr_3(m, v124, int32(20), v273, v23+int32(23))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L11
	} else {
		goto L67
	}
L66:
	;
	v307 = v300
	goto L62
L67:
	;
	v307 = v305
	goto L62
L68:
	;
	v313 = F_text_to_cstring(m, base.I32_wrap_i64(v307))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L11
	} else {
		goto L71
	}
L69:
	;
	v317 = v308
	goto L70
L70:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v122)+196))
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[0]))
	if v320 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v315 = F_stringToNode(m, v313)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	v317 = v315
	goto L70
L73:
	;
	v324 = int32(_a_F_RelationGetIndexAttrBitmap_4)
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1]))
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1])) = v328
	v331 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L11
	} else {
		goto L76
	}
L74:
	;
	v467 = v320
	goto L75
L75:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v318)+16))
	v486 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v485)+18)))
	if base.Ui32(v486&int32(2047)) <= base.Ui32(int32(20)) {
		goto L101
	} else {
		goto L102
	}
L76:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v331)+4)) = int64(-4294965047)
	v338 = int32(0)
	goto L77
L77:
	;
	v355 = int32(100)
	v356 = v338 * v355
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	base.MemoryCopy(m, v356+(v331+v357<<(uint(int32(3))%32))+int32(28), v356+int32(_a_F_RelationGetIndexAttrBitmap_5), v355)
	F_populate_compact_attribute(m, v331, v338)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L11
	} else {
		goto L79
	}
L78:
	;
	v374 = int32(0)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	if v374 < v383 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v371 = v338 + int32(1)
	if v371 != int32(21) {
		v338 = v371
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[0])) = v331
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1])) = v325
	v467 = v331
	goto L75
L82:
	;
	v387 = v331 + int32(28)
	v394 = v374
	v395 = v383
	v397 = v374
	goto L86
L83:
	;
	v451 = v374
	v458 = v383
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v331)+20)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v331)+16)) = v451
	goto L81
L85:
	;
	v451 = v445
	v458 = v424
	goto L84
L86:
	;
	v403 = v387 + v383<<(uint(int32(3))%32) + v394*int32(100)
	v406 = v387 + v394<<(uint(int32(3))%32)
	if v383 != v395 {
		v424 = v395
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v445 = v383
	goto L85
L88:
	;
	v425 = int32(*(*int16)(unsafe.Add(mBase, uint32(v406)+2)))
	if v425 <= int32(0) {
		v445 = v394
		goto L85
	} else {
		goto L96
	}
L89:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+7)))
	if v408 != int32(118) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v424 = v394
	goto L88
L91:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+4)))
	if v411 != int32(1) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+6)))
	if v414&int32(6) != 0 {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v417 = int32(*(*int16)(unsafe.Add(mBase, uint32(v406)+2)))
	if v417 <= int32(0) {
		goto L90
	} else {
		goto L94
	}
L94:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403)+90)))
	if v420 != int32(118) {
		v424 = v383
		goto L88
	} else {
		goto L95
	}
L95:
	;
	goto L90
L96:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403)+90)))
	if v428 == int32(118) {
		v445 = v394
		goto L85
	} else {
		goto L97
	}
L97:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+5)))
	v437 = (v397 + v431 - int32(1)) & (int32(0) - v431)
	if int32(_a_F_RelationGetIndexAttrBitmap_6) < v437 {
		v445 = v394
		goto L85
	} else {
		goto L98
	}
L98:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v406))) = uint16(v437)
	v443 = v394 + int32(1)
	if v443 != v383 {
		v394 = v443
		v395 = v424
		v397 = v437 + v425
		goto L86
	} else {
		goto L99
	}
L99:
	;
	goto L87
L100:
	;
	v502 = int32(0)
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+23)))
	if v504 == v502 {
		goto L106
	} else {
		goto L107
	}
L101:
	;
	v494 = F_getmissingattr(m, v467, int32(21), v23+int32(23))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L11
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v499 = F_fastgetattr_3(m, v318, int32(21), v467, v23+int32(23))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L11
	} else {
		goto L105
	}
L104:
	;
	v501 = v494
	goto L100
L105:
	;
	v501 = v499
	goto L100
L106:
	;
	v508 = F_text_to_cstring(m, base.I32_wrap_i64(v501))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L11
	} else {
		goto L109
	}
L107:
	;
	v512 = v502
	goto L108
L108:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v122)+204))
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+28)))
	if v518 != 0 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v510 = F_stringToNode(m, v508)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L11
	} else {
		goto L110
	}
L110:
	;
	v512 = v510
	goto L108
L111:
	;
	v519 = v23 + int32(24)
	goto L113
L112:
	;
	v519 = v23 + int32(28)
	goto L113
L113:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v122)+192))
	v521 = int32(*(*int16)(unsafe.Add(mBase, uint32(v520)+8)))
	if int32(0) < v521 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v524 = int32(0)
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520)+12)))
	v533 = v502
	v534 = v520
	v536 = v101
	v543 = v108
	v544 = v109
	goto L117
L115:
	;
	v601 = v101
	v608 = v108
	v609 = v109
	goto L116
L116:
	;
	F_pull_varattnos(m, v317, int32(1), v519)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L11
	} else {
		goto L134
	}
L117:
	;
	v554 = int32(*(*int16)(unsafe.Add(mBase, uint32(v534+v533<<(uint(int32(1))%32))+48)))
	if v554 == int32(0) {
		v587 = v534
		v588 = v536
		v590 = v543
		v591 = v544
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v601 = v588
	v608 = v590
	v609 = v591
	goto L116
L119:
	;
	v593 = v533 + int32(1)
	v594 = int32(*(*int16)(unsafe.Add(mBase, uint32(v587)+8)))
	if v593 < v594 {
		v533 = v593
		v534 = v587
		v536 = v588
		v543 = v590
		v544 = v591
		goto L117
	} else {
		goto L133
	}
L120:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	v559 = v554 + int32(7)
	v560 = F_bms_add_member(m, v557, v559)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L11
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519))) = v560
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v122)+192))
	if base.B2i32(v512 == v524)&(v526&base.B2i32(v317 == v524)) == int32(0) {
		v571 = v563
		v572 = v536
		goto L122
	} else {
		goto L123
	}
L122:
	;
	if v84 != v120 {
		v579 = v571
		v580 = v544
		goto L126
	} else {
		goto L127
	}
L123:
	;
	v566 = int32(*(*int16)(unsafe.Add(mBase, uint32(v563)+10)))
	if v566 <= v533 {
		v571 = v563
		v572 = v536
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v568 = F_bms_add_member(m, v536, v559)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L11
	} else {
		goto L125
	}
L125:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v122)+192))
	v571 = v570
	v572 = v568
	goto L122
L126:
	;
	if v83 != v120 {
		v587 = v579
		v588 = v572
		v590 = v543
		v591 = v580
		goto L119
	} else {
		goto L130
	}
L127:
	;
	v574 = int32(*(*int16)(unsafe.Add(mBase, uint32(v571)+10)))
	if v574 <= v533 {
		v579 = v571
		v580 = v544
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v576 = F_bms_add_member(m, v544, v559)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L11
	} else {
		goto L129
	}
L129:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v122)+192))
	v579 = v578
	v580 = v576
	goto L126
L130:
	;
	v582 = int32(*(*int16)(unsafe.Add(mBase, uint32(v579)+10)))
	if v582 <= v533 {
		v587 = v579
		v588 = v572
		v590 = v543
		v591 = v580
		goto L119
	} else {
		goto L131
	}
L131:
	;
	v584 = F_bms_add_member(m, v543, v559)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L11
	} else {
		goto L132
	}
L132:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v122)+192))
	v587 = v586
	v588 = v572
	v590 = v584
	v591 = v580
	goto L119
L133:
	;
	goto L118
L134:
	;
	F_pull_varattnos(m, v512, int32(1), v519)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L11
	} else {
		goto L135
	}
L135:
	;
	F_relation_close(m, v122, int32(1))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L11
	} else {
		goto L136
	}
L136:
	;
	v626 = v110 + int32(1)
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v626 < v627 {
		v101 = v601
		v108 = v608
		v109 = v609
		v110 = v626
		goto L32
	} else {
		goto L137
	}
L137:
	;
	goto L33
L138:
	;
	F_list_free(m, v649)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L11
	} else {
		goto L156
	}
L139:
	;
	v651 = F_equal(m, v72, v649)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L11
	} else {
		goto L140
	}
L140:
	;
	if v651 == int32(0) {
		goto L138
	} else {
		goto L141
	}
L141:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v84 != v655 {
		goto L138
	} else {
		goto L142
	}
L142:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v83 != v657 {
		goto L138
	} else {
		goto L143
	}
L143:
	;
	F_list_free(m, v649)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L11
	} else {
		goto L144
	}
L144:
	;
	F_list_free(m, v72)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L11
	} else {
		goto L145
	}
L145:
	;
	v663 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)) = uint8(v663)
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_bms_free(m, v665)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L11
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = int32(0)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	F_bms_free(m, v670)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L11
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = int32(0)
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	F_bms_free(m, v675)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L11
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(0)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	F_bms_free(m, v680)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L11
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(0)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	F_bms_free(m, v685)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L11
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	v690 = int32(_a_F_RelationGetIndexAttrBitmap_4)
	v691 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1]))
	v694 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1])) = v694
	v696 = F_bms_copy(m, v634)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L11
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v696
	v699 = F_bms_copy(m, v642)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L11
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v699
	v702 = F_bms_copy(m, v641)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L11
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v702
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v706 = F_bms_copy(m, v705)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L11
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v706
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	v710 = F_bms_copy(m, v709)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L11
	} else {
		goto L155
	}
L155:
	;
	v712 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)) = uint8(v712)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v710
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1])) = v691
	switch l1 {
	case 0:
		v761 = v634
		goto L1
	case 1:
		goto L26
	case 2:
		goto L25
	case 3:
		goto L24
	case 4:
		goto L23
	default:
		goto L22
	}
L156:
	;
	F_list_free(m, v72)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L11
	} else {
		goto L157
	}
L157:
	;
	F_bms_free(m, v634)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L11
	} else {
		goto L158
	}
L158:
	;
	F_bms_free(m, v642)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L11
	} else {
		goto L159
	}
L159:
	;
	F_bms_free(m, v641)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L11
	} else {
		goto L160
	}
L160:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	F_bms_free(m, v727)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L11
	} else {
		goto L161
	}
L161:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	F_bms_free(m, v730)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L11
	} else {
		goto L162
	}
L162:
	;
	v733 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L11
	} else {
		goto L163
	}
L163:
	;
	if v733 != 0 {
		v72 = v733
		goto L27
	} else {
		goto L164
	}
L164:
	;
	goto L28
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_RelationGetIndexAttrBitmap_0), v23+int32(16))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L11
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_RelationGetIndexAttrBitmap_1), int32(_a_F_RelationGetIndexAttrBitmap_7), int32(_a_F_RelationGetIndexAttrBitmap_3))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L11
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	v761 = v754
	goto L1
}
func F_RelationInitIndexAccessInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
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
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v299 int64
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v455 int32
	_ = v455
	var v478 int64
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v501 int32
	_ = v501
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v625 int64
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int64
	_ = v631
	var v633 int32
	_ = v633
	var v637 int64
	_ = v637
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v715 int32
	_ = v715
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v800 int32
	_ = v800
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v901 int32
	_ = v901
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v915 int32
	_ = v915
	var v922 int32
	_ = v922
	var v930 int32
	_ = v930
	var v953 int64
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int64
	_ = v963
	var v965 int32
	_ = v965
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v999 int32
	_ = v999
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	v22 = m.G0
	v24 = v22 - int32(304)
	m.G0 = v24
	v27 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v28 = F_SearchSysCache1(m, int32(34), v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L6
	} else {
		goto L193
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L6
	} else {
		goto L190
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L6
	} else {
		goto L187
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L6
	} else {
		goto L184
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L6
	} else {
		goto L181
	}
L6:
	;
	return
L7:
	;
	if v28 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v30 = int32(_a_F_RelationInitIndexAccessInfo_0)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0]))
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v34
	v36 = F_heap_copytuple(m, v28)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L6
	} else {
		goto L178
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v36
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v39 + v40
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v31
	F_ReleaseCatCache(m, v28)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v48)+84)))
	v50 = F_SearchSysCache1(m, int32(2), v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	if v50 == int32(0) {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+22)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54+v55)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v57
	F_ReleaseCatCache(m, v50)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+120)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+8)))
	if v62 != v64 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v63)+10)))
	v67 = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[1]))
	v74 = F_AllocSetContextCreateInternal(m, v69, int32(_a_F_RelationInitIndexAccessInfo_1), v67, int32(1024), int32(_a_F_RelationInitIndexAccessInfo_2))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v80 = F_MemoryContextStrdup(m, v74, v77+int32(4))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+36)) = v80
	v83 = int32(_a_F_RelationInitIndexAccessInfo_0)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v89 = F_GetIndexAmRoutine(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v89
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v84
	v95 = v66 << (uint(int32(2)) % 32)
	v96 = F_MemoryContextAllocZero(m, v74, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v96
	v99 = F_MemoryContextAllocZero(m, v74, v95)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+212)) = v99
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+6)))
	if v103 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+220)) = v117
	v119 = F_MemoryContextAllocZero(m, v74, v95)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L28
	}
L23:
	;
	v105 = v103 * base.I32_extend16_s(v62)
	v108 = F_MemoryContextAllocZero(m, v74, v105<<(uint(int32(2))%32))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = int32(0)
	v117 = v67
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v108
	v113 = F_MemoryContextAllocZero(m, v74, v105*int32(28))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v117 = v113
	goto L22
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+248)) = v119
	v123 = v66 << (uint(int32(1)) % 32)
	v124 = F_MemoryContextAllocZero(m, v74, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = v124
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v128 = int32(0)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[2]))
	if v130 == v128 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v133 = int32(_a_F_RelationInitIndexAccessInfo_0)
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0]))
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v137
	v140 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L33
	}
L31:
	;
	v276 = v130
	goto L32
L32:
	;
	v299 = F_fastgetattr_3(m, v127, int32(17), v276, v24+int32(79))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L6
	} else {
		goto L57
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v140)+4)) = int64(-4294965047)
	v146 = v128
	goto L34
L34:
	;
	v165 = int32(100)
	v166 = v146 * v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	base.MemoryCopy(m, v166+(v140+v167<<(uint(int32(3))%32))+int32(28), v166+int32(_a_F_RelationInitIndexAccessInfo_3), v165)
	F_populate_compact_attribute(m, v140, v146)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L36
	}
L35:
	;
	v184 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	if v184 < v193 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v181 = v146 + int32(1)
	if v181 != int32(21) {
		v146 = v181
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[2])) = v140
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v134
	v276 = v140
	goto L32
L39:
	;
	v197 = v140 + int32(28)
	v204 = v184
	v205 = v193
	v207 = v184
	goto L43
L40:
	;
	v261 = v184
	v268 = v193
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+20)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v140)+16)) = v261
	goto L38
L42:
	;
	v261 = v255
	v268 = v234
	goto L41
L43:
	;
	v213 = v197 + v193<<(uint(int32(3))%32) + v204*int32(100)
	v216 = v197 + v204<<(uint(int32(3))%32)
	if v193 != v205 {
		v234 = v205
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v255 = v193
	goto L42
L45:
	;
	v235 = int32(*(*int16)(unsafe.Add(mBase, uint32(v216)+2)))
	if v235 <= int32(0) {
		v255 = v204
		goto L42
	} else {
		goto L53
	}
L46:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+7)))
	if v218 != int32(118) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v234 = v204
	goto L45
L48:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+4)))
	if v221 != int32(1) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+6)))
	if v224&int32(6) != 0 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v227 = int32(*(*int16)(unsafe.Add(mBase, uint32(v216)+2)))
	if v227 <= int32(0) {
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+90)))
	if v230 != int32(118) {
		v234 = v193
		goto L45
	} else {
		goto L52
	}
L52:
	;
	goto L47
L53:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+90)))
	if v238 == int32(118) {
		v255 = v204
		goto L42
	} else {
		goto L54
	}
L54:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+5)))
	v247 = (v207 + v241 - int32(1)) & (int32(0) - v241)
	if int32(_a_F_RelationInitIndexAccessInfo_4) < v247 {
		v255 = v204
		goto L42
	} else {
		goto L55
	}
L55:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v216))) = uint16(v247)
	v253 = v204 + int32(1)
	if v253 != v193 {
		v204 = v253
		v205 = v234
		v207 = v247 + v235
		goto L43
	} else {
		goto L56
	}
L56:
	;
	goto L44
L57:
	;
	if v95 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	base.MemoryCopy(m, v301, base.I32_wrap_i64(v299)+int32(24), v95)
	goto L60
L59:
	;
	goto L60
L60:
	;
	v306 = int32(0)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[2]))
	if v309 == v306 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v312 = int32(_a_F_RelationInitIndexAccessInfo_0)
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0]))
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v316
	v319 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L6
	} else {
		goto L64
	}
L62:
	;
	v455 = v309
	goto L63
L63:
	;
	v478 = F_fastgetattr_3(m, v307, int32(18), v455, v24+int32(79))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L6
	} else {
		goto L88
	}
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v319)+4)) = int64(-4294965047)
	v325 = v306
	goto L65
L65:
	;
	v344 = int32(100)
	v345 = v325 * v344
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	base.MemoryCopy(m, v345+(v319+v346<<(uint(int32(3))%32))+int32(28), v345+int32(_a_F_RelationInitIndexAccessInfo_3), v344)
	F_populate_compact_attribute(m, v319, v325)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L6
	} else {
		goto L67
	}
L66:
	;
	v363 = int32(0)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	if v363 < v372 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v360 = v325 + int32(1)
	if v360 != int32(21) {
		v325 = v360
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[2])) = v319
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v313
	v455 = v319
	goto L63
L70:
	;
	v376 = v319 + int32(28)
	v383 = v363
	v384 = v372
	v386 = v363
	goto L74
L71:
	;
	v440 = v363
	v447 = v372
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319)+20)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v319)+16)) = v440
	goto L69
L73:
	;
	v440 = v434
	v447 = v413
	goto L72
L74:
	;
	v392 = v376 + v372<<(uint(int32(3))%32) + v383*int32(100)
	v395 = v376 + v383<<(uint(int32(3))%32)
	if v372 != v384 {
		v413 = v384
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v434 = v372
	goto L73
L76:
	;
	v414 = int32(*(*int16)(unsafe.Add(mBase, uint32(v395)+2)))
	if v414 <= int32(0) {
		v434 = v383
		goto L73
	} else {
		goto L84
	}
L77:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+7)))
	if v397 != int32(118) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v413 = v383
	goto L76
L79:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+4)))
	if v400 != int32(1) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+6)))
	if v403&int32(6) != 0 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v406 = int32(*(*int16)(unsafe.Add(mBase, uint32(v395)+2)))
	if v406 <= int32(0) {
		goto L78
	} else {
		goto L82
	}
L82:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+90)))
	if v409 != int32(118) {
		v413 = v372
		goto L76
	} else {
		goto L83
	}
L83:
	;
	goto L78
L84:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+90)))
	if v417 == int32(118) {
		v434 = v383
		goto L73
	} else {
		goto L85
	}
L85:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+5)))
	v426 = (v386 + v420 - int32(1)) & (int32(0) - v420)
	if int32(_a_F_RelationInitIndexAccessInfo_4) < v426 {
		v434 = v383
		goto L73
	} else {
		goto L86
	}
L86:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v395))) = uint16(v426)
	v432 = v383 + int32(1)
	if v432 != v372 {
		v383 = v432
		v384 = v413
		v386 = v426 + v414
		goto L74
	} else {
		goto L87
	}
L87:
	;
	goto L75
L88:
	;
	if int32(0) < v66 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v486 = v103 << (uint(int32(2)) % 32)
	v501 = int32(0)
	goto L92
L90:
	;
	goto L91
L91:
	;
	v781 = int32(0)
	v782 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v784 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[2]))
	if v784 == v781 {
		goto L146
	} else {
		goto L147
	}
L92:
	;
	v517 = v501 << (uint(int32(2)) % 32)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v478)+int32(24)+v517)))
	if v519 == int32(0) {
		goto L3
	} else {
		goto L94
	}
L93:
	;
	goto L91
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+300)) = v519
	v524 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[3]))
	if v524 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v528 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[1]))
	if v528 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v544 = v524
	goto L97
L97:
	;
	v550 = F_hash_search(m, v544, v24+int32(300), int32(1), v24+int32(80))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L6
	} else {
		goto L103
	}
L98:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L6
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+88)) = int64(85899345924)
	v541 = F_hash_create(m, int32(_a_F_RelationInitIndexAccessInfo_5), int64(64), v24+int32(80), int32(40))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L6
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[3])) = v541
	v544 = v541
	goto L97
L103:
	;
	v552 = int32(0)
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+80)))
	if v554 == v552 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v550)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v517+v483))) = v739
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v550)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v517+v482))) = v742
	v744 = int32(0)
	if base.B2i32(v103 == v744)|base.B2i32(v486 == v744) == v744 {
		goto L142
	} else {
		goto L143
	}
L105:
	;
	v568 = int32(0)
	if base.B2i32(v103 == v552)|base.B2i32(v567 == v568) == v568 {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	v557 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v550)+16)) = v557
	*(*uint16)(unsafe.Add(mBase, uint32(v550)+6)) = uint16(v103)
	*(*uint8)(unsafe.Add(mBase, uint32(v550)+4)) = uint8(v557)
	v567 = int32(1)
	goto L105
L107:
	;
	goto L108
L108:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+4)))
	if v563 != 0 {
		goto L104
	} else {
		goto L109
	}
L109:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v550)+16))
	v567 = base.B2i32(v564 == int32(0))
	goto L105
L110:
	;
	v574 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[1]))
	v575 = F_MemoryContextAllocZero(m, v574, v486)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L6
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v24)+300))
	v587 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[4])))
	if v587 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v550)+16)) = v575
	goto L112
L114:
	;
	v588 = int32(1)
	goto L116
L115:
	;
	v588 = base.B2i32(v579 != int32(1981)) & base.B2i32(v579 != int32(1979))
	goto L116
L116:
	;
	v590 = v24 + int32(128)
	F_ScanKeyInit(m, v590, int32(1), int32(3), int32(184), base.I64_extend_i32_u(v579))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	v599 = F_table_open(m, int32(2616), int32(1))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L6
	} else {
		goto L118
	}
L118:
	;
	v604 = F_systable_beginscan(m, v599, int32(2687), v588, int32(0), int32(1), v590)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	v606 = F_systable_getnext(m, v604)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	if v606 == int32(0) {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v606)+16))
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610)+22)))
	v612 = v610 + v611
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v612)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v550)+8)) = v613
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v612)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v550)+12)) = v615
	F_systable_endscan(m, v604)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	F_relation_close(m, v599, int32(1))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L6
	} else {
		goto L123
	}
L123:
	;
	if v103 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v625 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v550)+8)))
	F_ScanKeyInit(m, v590, int32(2), int32(3), int32(184), v625)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L6
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v715 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v550)+4)) = uint8(v715)
	goto L104
L127:
	;
	v628 = int32(3)
	v631 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v550)+12)))
	F_ScanKeyInit(m, v24+int32(184), v628, v628, int32(184), v631)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	v637 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v550)+12)))
	F_ScanKeyInit(m, v24+int32(240), int32(4), int32(3), int32(184), v637)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L6
	} else {
		goto L129
	}
L129:
	;
	v642 = F_table_open(m, int32(2603), int32(1))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L6
	} else {
		goto L130
	}
L130:
	;
	v647 = F_systable_beginscan(m, v642, int32(2655), v588, int32(0), int32(3), v590)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	goto L132
L132:
	;
	v670 = F_systable_getnext(m, v647)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L6
	} else {
		goto L134
	}
L133:
	;
	F_systable_endscan(m, v647)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L6
	} else {
		goto L140
	}
L134:
	;
	if v670 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v670)+16))
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+22)))
	v674 = v672 + v673
	v675 = int32(*(*int16)(unsafe.Add(mBase, uint32(v674)+16)))
	if v675 <= int32(0) {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	goto L133
L138:
	;
	v679 = v675 & int32(_a_F_RelationInitIndexAccessInfo_6)
	if base.Ui32(v103) < base.Ui32(v679) {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v550)+16))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v674)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v681+v679<<(uint(int32(2))%32)-int32(4)))) = v687
	goto L132
L140:
	;
	F_relation_close(m, v642, int32(1))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	goto L126
L142:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v550)+16))
	base.MemoryCopy(m, v484+v103*v501<<(uint(int32(2))%32), v755, v486)
	goto L144
L143:
	;
	goto L144
L144:
	;
	v758 = v501 + int32(1)
	if v758 != v66 {
		v501 = v758
		goto L92
	} else {
		goto L145
	}
L145:
	;
	goto L93
L146:
	;
	v787 = int32(_a_F_RelationInitIndexAccessInfo_0)
	v788 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0]))
	v791 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v791
	v794 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L6
	} else {
		goto L149
	}
L147:
	;
	v930 = v784
	goto L148
L148:
	;
	v953 = F_fastgetattr_3(m, v782, int32(19), v930, v24+int32(79))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L6
	} else {
		goto L173
	}
L149:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v794)+4)) = int64(-4294965047)
	v800 = v781
	goto L150
L150:
	;
	v819 = int32(100)
	v820 = v800 * v819
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v794)))
	base.MemoryCopy(m, v820+(v794+v821<<(uint(int32(3))%32))+int32(28), v820+int32(_a_F_RelationInitIndexAccessInfo_3), v819)
	F_populate_compact_attribute(m, v794, v800)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L6
	} else {
		goto L152
	}
L151:
	;
	v838 = int32(0)
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v794)))
	if v838 < v847 {
		goto L155
	} else {
		goto L156
	}
L152:
	;
	v835 = v800 + int32(1)
	if v835 != int32(21) {
		v800 = v835
		goto L150
	} else {
		goto L153
	}
L153:
	;
	goto L151
L154:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[2])) = v794
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v788
	v930 = v794
	goto L148
L155:
	;
	v851 = v794 + int32(28)
	v858 = v838
	v859 = v847
	v861 = v838
	goto L159
L156:
	;
	v915 = v838
	v922 = v847
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v794)+20)) = v922
	*(*int32)(unsafe.Add(mBase, uint32(v794)+16)) = v915
	goto L154
L158:
	;
	v915 = v909
	v922 = v888
	goto L157
L159:
	;
	v867 = v851 + v847<<(uint(int32(3))%32) + v858*int32(100)
	v870 = v851 + v858<<(uint(int32(3))%32)
	if v847 != v859 {
		v888 = v859
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v909 = v847
	goto L158
L161:
	;
	v889 = int32(*(*int16)(unsafe.Add(mBase, uint32(v870)+2)))
	if v889 <= int32(0) {
		v909 = v858
		goto L158
	} else {
		goto L169
	}
L162:
	;
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v870)+7)))
	if v872 != int32(118) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v888 = v858
	goto L161
L164:
	;
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v870)+4)))
	if v875 != int32(1) {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v870)+6)))
	if v878&int32(6) != 0 {
		goto L163
	} else {
		goto L166
	}
L166:
	;
	v881 = int32(*(*int16)(unsafe.Add(mBase, uint32(v870)+2)))
	if v881 <= int32(0) {
		goto L163
	} else {
		goto L167
	}
L167:
	;
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867)+90)))
	if v884 != int32(118) {
		v888 = v847
		goto L161
	} else {
		goto L168
	}
L168:
	;
	goto L163
L169:
	;
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867)+90)))
	if v892 == int32(118) {
		v909 = v858
		goto L158
	} else {
		goto L170
	}
L170:
	;
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v870)+5)))
	v901 = (v861 + v895 - int32(1)) & (int32(0) - v895)
	if int32(_a_F_RelationInitIndexAccessInfo_4) < v901 {
		v909 = v858
		goto L158
	} else {
		goto L171
	}
L171:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v870))) = uint16(v901)
	v907 = v858 + int32(1)
	if v907 != v847 {
		v858 = v907
		v859 = v888
		v861 = v901 + v889
		goto L159
	} else {
		goto L172
	}
L172:
	;
	goto L160
L173:
	;
	if v123 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	base.MemoryCopy(m, v955, base.I32_wrap_i64(v953)+int32(24), v123)
	goto L176
L175:
	;
	goto L176
L176:
	;
	v961 = F_RelationGetIndexAttOptions(m, l0, int32(0))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	v963 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+228)) = v963
	v965 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v965
	*(*int64)(unsafe.Add(mBase, uint32(l0)+236)) = v963
	*(*int32)(unsafe.Add(mBase, uint32(l0)+244)) = v965
	m.G0 = v24 + int32(304)
	return
L178:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v978
	F_errmsg_internal(m, int32(_a_F_RelationInitIndexAccessInfo_7), v24)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L6
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_RelationInitIndexAccessInfo_8), int32(1466), int32(_a_F_RelationInitIndexAccessInfo_9))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L6
	} else {
		goto L180
	}
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v992)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v993
	F_errmsg_internal(m, int32(_a_F_RelationInitIndexAccessInfo_10), v24+int32(16))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L6
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_RelationInitIndexAccessInfo_8), int32(1480), int32(_a_F_RelationInitIndexAccessInfo_9))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L6
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v1009
	F_errmsg_internal(m, int32(_a_F_RelationInitIndexAccessInfo_11), v24-int32(-64))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L6
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(_a_F_RelationInitIndexAccessInfo_8), int32(1488), int32(_a_F_RelationInitIndexAccessInfo_9))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L6
	} else {
		goto L186
	}
L186:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L187:
	;
	F_errmsg_internal(m, int32(_a_F_RelationInitIndexAccessInfo_12), int32(0))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L6
	} else {
		goto L188
	}
L188:
	;
	F_errfinish(m, int32(_a_F_RelationInitIndexAccessInfo_8), int32(1625), int32(_a_F_RelationInitIndexAccessInfo_13))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L190:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v24)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v1038
	F_errmsg_internal(m, int32(_a_F_RelationInitIndexAccessInfo_14), v24+int32(32))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L6
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_RelationInitIndexAccessInfo_8), int32(1761), int32(_a_F_RelationInitIndexAccessInfo_15))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L6
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	v1055 = int32(*(*int16)(unsafe.Add(mBase, uint32(v674)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v1055
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v24)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+52)) = v1057
	F_errmsg_internal(m, int32(_a_F_RelationInitIndexAccessInfo_16), v24+int32(48))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L6
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_RelationInitIndexAccessInfo_8), int32(1795), int32(_a_F_RelationInitIndexAccessInfo_15))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L6
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationMapFilenumberToOid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v3 = int32(0)
	if l1 == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v105
L2:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v105 = v100
	goto L1
L3:
	;
	v8 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapFilenumberToOid[0]))
	if v8 < v10 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v51 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapFilenumberToOid[1]))
	if v51 < v53 {
		goto L24
	} else {
		goto L25
	}
L6:
	;
	v14 = v8
	goto L9
L7:
	;
	goto L8
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapFilenumberToOid[2]))
	if v33 <= int32(0) {
		v105 = v3
		goto L1
	} else {
		goto L15
	}
L9:
	;
	v19 = v14 << (uint(int32(3)) % 32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_RelationMapFilenumberToOid[3])))
	if v20 == l0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v99 = v19 + int32(_a_F_RelationMapFilenumberToOid_0)
	goto L2
L12:
	;
	goto L13
L13:
	;
	v25 = v14 + int32(1)
	if v25 != v10 {
		v14 = v25
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v38 = int32(0)
	goto L16
L16:
	;
	v43 = v38 << (uint(int32(3)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_RelationMapFilenumberToOid[4])))
	if v44 != l0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v99 = v43 + int32(_a_F_RelationMapFilenumberToOid_1)
	goto L2
L18:
	;
	v47 = v38 + int32(1)
	if v33 != v47 {
		v38 = v47
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v105 = v3
	goto L1
L22:
	;
	v81 = int32(0)
	goto L32
L23:
	;
	v99 = v62 + int32(_a_F_RelationMapFilenumberToOid_2)
	goto L2
L24:
	;
	v57 = v51
	goto L27
L25:
	;
	goto L26
L26:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapFilenumberToOid[5]))
	if v74 <= int32(0) {
		v105 = v3
		goto L1
	} else {
		goto L31
	}
L27:
	;
	v62 = v57 << (uint(int32(3)) % 32)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+uint32(_c_F_RelationMapFilenumberToOid[6])))
	if l0 == v63 {
		goto L23
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v66 = v57 + int32(1)
	if v66 != v53 {
		v57 = v66
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L22
L32:
	;
	v86 = v81 << (uint(int32(3)) % 32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+uint32(_c_F_RelationMapFilenumberToOid[7])))
	if v87 != l0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v99 = v86 + int32(_a_F_RelationMapFilenumberToOid_3)
	goto L2
L34:
	;
	v90 = v81 + int32(1)
	if v74 != v90 {
		v81 = v90
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	v105 = v3
	goto L1
}
func F_RelationMapOidToFilenumber(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v3 = int32(0)
	if l1 == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v105
L2:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v105 = v100
	goto L1
L3:
	;
	v8 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapOidToFilenumber[0]))
	if v8 < v10 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v51 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapOidToFilenumber[1]))
	if v51 < v53 {
		goto L24
	} else {
		goto L25
	}
L6:
	;
	v14 = v8
	goto L9
L7:
	;
	goto L8
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapOidToFilenumber[2]))
	if v33 <= int32(0) {
		v105 = v3
		goto L1
	} else {
		goto L15
	}
L9:
	;
	v19 = v14 << (uint(int32(3)) % 32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_RelationMapOidToFilenumber[3])))
	if v20 == l0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v99 = v19 + int32(_a_F_RelationMapOidToFilenumber_0)
	goto L2
L12:
	;
	goto L13
L13:
	;
	v25 = v14 + int32(1)
	if v25 != v10 {
		v14 = v25
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v38 = int32(0)
	goto L16
L16:
	;
	v43 = v38 << (uint(int32(3)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_RelationMapOidToFilenumber[4])))
	if v44 != l0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v99 = v43 + int32(_a_F_RelationMapOidToFilenumber_1)
	goto L2
L18:
	;
	v47 = v38 + int32(1)
	if v33 != v47 {
		v38 = v47
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v105 = v3
	goto L1
L22:
	;
	v81 = int32(0)
	goto L32
L23:
	;
	v99 = v62 + int32(_a_F_RelationMapOidToFilenumber_2)
	goto L2
L24:
	;
	v57 = v51
	goto L27
L25:
	;
	goto L26
L26:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapOidToFilenumber[5]))
	if v74 <= int32(0) {
		v105 = v3
		goto L1
	} else {
		goto L31
	}
L27:
	;
	v62 = v57 << (uint(int32(3)) % 32)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+uint32(_c_F_RelationMapOidToFilenumber[6])))
	if l0 == v63 {
		goto L23
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v66 = v57 + int32(1)
	if v66 != v53 {
		v57 = v66
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L22
L32:
	;
	v86 = v81 << (uint(int32(3)) % 32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+uint32(_c_F_RelationMapOidToFilenumber[7])))
	if v87 != l0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v99 = v86 + int32(_a_F_RelationMapOidToFilenumber_3)
	goto L2
L34:
	;
	v90 = v81 + int32(1)
	if v74 != v90 {
		v81 = v90
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	v105 = v3
	goto L1
}
func F_RelationMapOidToFilenumberForDatabase(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(528)
	m.G0 = v9
	F_read_relmap_file(m, v9+int32(4), l0, v3, int32(21))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v19 <= int32(0) {
		v42 = v3
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v9 + int32(528)
	return v42
L4:
	;
	v26 = v3
	goto L6
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v42 = v39
	goto L3
L6:
	;
	v32 = v9 + int32(12) + v26<<(uint(int32(3))%32)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if l1 == v33 {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	v42 = int32(0)
	goto L3
L8:
	;
	v36 = v26 + int32(1)
	if v36 != v19 {
		v26 = v36
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
}
func F_RelationTruncateIndexes(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v247 int64
	_ = v247
	var v248 int32
	_ = v248
	var v252 int64
	_ = v252
	var v253 int32
	_ = v253
	var v254 int64
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v351 int32
	_ = v351
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v400 int32
	_ = v400
	var v418 int32
	_ = v418
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	v17 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	return
L3:
	;
	if v17 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v21 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v37 = int32(0)
	goto L6
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v37<<(uint(int32(2))%32))))
	v46 = F_index_open(m, v44, int32(8))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	F_RelationTruncate(m, v46, int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L2
	} else {
		goto L82
	}
L9:
	;
	v48 = int32(0)
	v50 = m.G0
	v52 = v50 - int32(16)
	m.G0 = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)+192))
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+8)))
	if base.Ui32(int32(_a_F_RelationTruncateIndexes_0)) < base.Ui32((v55-int32(33))&int32(_a_F_RelationTruncateIndexes_1)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v63 = v54 + int32(48)
	v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+10)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+84))
	v67 = int32(0)
	v68 = m.G0
	v70 = v68 - int32(16)
	m.G0 = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v46)+196))
	if v72 == v67 {
		v314 = v67
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L2
	} else {
		goto L79
	}
L13:
	;
	m.G0 = v70 + int32(16)
	v329 = int32(0)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+12)))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+13)))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+20)))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v46)+204))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+28)))
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+15)))
	v338 = F_makeIndexInfo(m, v55, v64, v66, v314, v329, v330, v331, v332, v329, v335, v330&v336)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L2
	} else {
		goto L67
	}
L14:
	;
	v77 = F_heap_attisnull(m, v72, int32(20), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v77 != 0 {
		v314 = v67
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v79 = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v46)+196))
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncateIndexes[0]))
	if v82 == v79 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v85 = int32(_a_F_RelationTruncateIndexes_2)
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncateIndexes[1]))
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncateIndexes[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationTruncateIndexes[1])) = v89
	v92 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	v224 = v82
	goto L19
L19:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v238)+18)))
	if base.Ui32(v239&int32(2044)) <= base.Ui32(int32(19)) {
		goto L45
	} else {
		goto L46
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v92)+4)) = int64(-4294965047)
	v104 = v79
	goto L21
L21:
	;
	v112 = int32(100)
	v113 = v104 * v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	base.MemoryCopy(m, v113+(v92+v114<<(uint(int32(3))%32))+int32(28), v113+int32(_a_F_RelationTruncateIndexes_3), v112)
	F_populate_compact_attribute(m, v92, v104)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L23
	}
L22:
	;
	v131 = int32(0)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v131 < v140 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v128 = v104 + int32(1)
	if v128 != int32(21) {
		v104 = v128
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationTruncateIndexes[0])) = v92
	*(*int32)(unsafe.Add(mBase, _c_F_RelationTruncateIndexes[1])) = v86
	v224 = v92
	goto L19
L26:
	;
	v144 = v92 + int32(28)
	v151 = v131
	v152 = v140
	v154 = v131
	goto L30
L27:
	;
	v208 = v131
	v215 = v140
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v208
	goto L25
L29:
	;
	v208 = v202
	v215 = v181
	goto L28
L30:
	;
	v160 = v144 + v140<<(uint(int32(3))%32) + v151*int32(100)
	v163 = v144 + v151<<(uint(int32(3))%32)
	if v140 != v152 {
		v181 = v152
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v202 = v140
	goto L29
L32:
	;
	v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(v163)+2)))
	if v182 <= int32(0) {
		v202 = v151
		goto L29
	} else {
		goto L40
	}
L33:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+7)))
	if v165 != int32(118) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v181 = v151
	goto L32
L35:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+4)))
	if v168 != int32(1) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+6)))
	if v171&int32(6) != 0 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v174 = int32(*(*int16)(unsafe.Add(mBase, uint32(v163)+2)))
	if v174 <= int32(0) {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+90)))
	if v177 != int32(118) {
		v181 = v140
		goto L32
	} else {
		goto L39
	}
L39:
	;
	goto L34
L40:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+90)))
	if v185 == int32(118) {
		v202 = v151
		goto L29
	} else {
		goto L41
	}
L41:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+5)))
	v194 = (v154 + v188 - int32(1)) & (int32(0) - v188)
	if int32(_a_F_RelationTruncateIndexes_4) < v194 {
		v202 = v151
		goto L29
	} else {
		goto L42
	}
L42:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v163))) = uint16(v194)
	v200 = v151 + int32(1)
	if v200 != v140 {
		v151 = v200
		v152 = v181
		v154 = v194 + v182
		goto L30
	} else {
		goto L43
	}
L43:
	;
	goto L31
L44:
	;
	v256 = F_text_to_cstring(m, base.I32_wrap_i64(v254))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L2
	} else {
		goto L50
	}
L45:
	;
	v247 = F_getmissingattr(m, v224, int32(20), v70+int32(15))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L2
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v252 = F_fastgetattr_3(m, v80, int32(20), v224, v70+int32(15))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L2
	} else {
		goto L49
	}
L48:
	;
	v254 = v247
	goto L44
L49:
	;
	v254 = v252
	goto L44
L50:
	;
	v258 = F_stringToNode(m, v256)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	F_pfree(m, v256)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	if v258 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v314 = int32(0)
	goto L13
L54:
	;
	goto L55
L55:
	;
	v265 = int32(0)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v266 <= v265 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v314 = int32(0)
	goto L13
L57:
	;
	goto L58
L58:
	;
	v275 = int32(0)
	v279 = v265
	goto L59
L59:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v258)+12))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v287+v279<<(uint(int32(2))%32))))
	v292 = F_exprType(m, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L2
	} else {
		goto L61
	}
L60:
	;
	v314 = v304
	goto L13
L61:
	;
	v294 = F_exprTypmod(m, v291)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	v296 = F_exprCollation(m, v291)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	v298 = int32(1)
	v302 = F_makeConst(m, v292, v294, v296, v298, int64(0), v298, v298)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	v304 = F_lappend(m, v275, v302)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	v307 = v279 + int32(1)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v307 < v308 {
		v275 = v304
		v279 = v307
		goto L59
	} else {
		goto L66
	}
L66:
	;
	goto L60
L67:
	;
	v341 = v338 + int32(12)
	v342 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v55-int32(1)) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	m.G0 = v52 + int32(16)
	goto L8
L69:
	;
	v351 = v342
	v363 = v48
	goto L72
L70:
	;
	v400 = v342
	goto L71
L71:
	;
	v418 = v400
	v431 = v48
	goto L76
L72:
	;
	v366 = v351 << (uint(int32(1)) % 32)
	v369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v366+v63))))
	*(*uint16)(unsafe.Add(mBase, uint32(v341+v366))) = uint16(v369)
	v372 = v366 | int32(2)
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v372+v63))))
	*(*uint16)(unsafe.Add(mBase, uint32(v341+v372))) = uint16(v375)
	v377 = int32(4)
	v378 = v366 | v377
	v381 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v378+v63))))
	*(*uint16)(unsafe.Add(mBase, uint32(v341+v378))) = uint16(v381)
	v384 = v366 | int32(6)
	v387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v384+v63))))
	*(*uint16)(unsafe.Add(mBase, uint32(v341+v384))) = uint16(v387)
	v390 = v351 + v377
	v392 = v363 + v377
	if v392 != v55&int32(60) {
		v351 = v390
		v363 = v392
		goto L72
	} else {
		goto L74
	}
L73:
	;
	if v55&int32(3) == int32(0) {
		goto L68
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	v400 = v390
	goto L71
L76:
	;
	v432 = int32(1)
	v433 = v418 << (uint(v432) % 32)
	v436 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v433+v63))))
	*(*uint16)(unsafe.Add(mBase, uint32(v341+v433))) = uint16(v436)
	v441 = v431 + v432
	if v441 != v55&int32(3) {
		v418 = v418 + v432
		v431 = v441
		goto L76
	} else {
		goto L78
	}
L77:
	;
	goto L68
L78:
	;
	goto L77
L79:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v46)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v55
	F_errmsg_internal(m, int32(_a_F_RelationTruncateIndexes_5), v52)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_RelationTruncateIndexes_6), int32(2532), int32(_a_F_RelationTruncateIndexes_7))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	v480 = int32(1)
	F_index_build(m, l0, v46, v338, v480, int32(0), v480)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	F_relation_close(m, v46, int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v489 = v37 + int32(1)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v489 < v490 {
		v37 = v489
		goto L6
	} else {
		goto L85
	}
L85:
	;
	goto L7
}
func F_SetRelationTableSpace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v20 = F_SearchSysCacheLockedCopy1(m, int32(57), base.I64_extend_i32_u(v13))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if v20 != 0 {
				v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+8)))
				*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v22)
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v24
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+22)))
				v28 = v26 + v27
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_SetRelationTableSpace[0]))
				if l1 != v31 {
					v33 = l1
				} else {
					v33 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = v33
				if l2 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v28)+88)) = l2
				} else {
				}
				v37 = v11 + int32(8)
				F_CatalogTupleUpdate(m, v16, v37, v20)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					F_UnlockTuple(m, v16, v37, int32(7))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+119)))
						switch v44 - int32(83) {
						case 0, 22, 26, 31, 33:
							F_pfree(m, v20)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								F_relation_close(m, v16, int32(3))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									m.G0 = v11 + int32(16)
									return
								}
							}
						default:
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v28)+92))
							v50 = F_table_open(m, int32(1214), int32(3))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								v52 = int32(0)
								if base.B2i32(v47 == v52)|base.B2i32(v47 == int32(1663)) == v52 {
									F_shdepChangeDep(m, v50, int32(1259), v13, int32(1213), v47, int32(116))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										F_relation_close(m, v50, int32(3))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												F_relation_close(m, v16, int32(3))
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16)
													return
												}
											}
										}
									}
								} else {
									v65 = int32(0)
									F_shdepDropDependency(m, v50, int32(1259), v13, v65, int32(1), v65, v65, v65)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										F_relation_close(m, v50, int32(3))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											F_pfree(m, v20)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												F_relation_close(m, v16, int32(3))
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
					F_errmsg_internal(m, int32(_a_F_SetRelationTableSpace_0), v11)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_SetRelationTableSpace_1), int32(3800), int32(_a_F_SetRelationTableSpace_2))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_generate_relation_name(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v18 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v198 = F_quote_identifier(m, v26)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L3
	} else {
		goto L45
	}
L2:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	v169 = F_get_namespace_name_or_temp(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L3
	} else {
		goto L40
	}
L3:
	;
	return int32(0)
L4:
	;
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
	v24 = v22 + v23
	v26 = v24 + int32(4)
	if l1 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L3
	} else {
		goto L37
	}
L8:
	;
	v136 = F_RelationIsVisible(m, l0)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L3
	} else {
		goto L34
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v29 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v32 = int32(0)
	if v32 < v29 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v35 = v29
	goto L13
L12:
	;
	v35 = v32
	goto L13
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v43 = int32(0)
	goto L14
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v36+v43<<(uint(int32(2))%32))))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	if v52 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L8
L16:
	;
	v123 = v43 + int32(1)
	if v123 != v35 {
		v43 = v123
		goto L14
	} else {
		goto L33
	}
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v55 <= int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v58 = int32(0)
	if v58 < v55 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v61 = v55
	goto L21
L20:
	;
	v61 = v58
	goto L21
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v65 = int32(0)
	goto L22
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v62+v65<<(uint(int32(2))%32))))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if base.B2i32(v82 == int32(0))|base.B2i32(v82 != v85) != 0 {
		v103 = v82
		v104 = v85
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L16
L24:
	;
	if v103-v104 == int32(0) {
		goto L2
	} else {
		goto L31
	}
L25:
	;
	goto L24
L26:
	;
	v88 = v79
	v89 = v26
	goto L27
L27:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	if v93 == int32(0) {
		v103 = v93
		v104 = v92
		goto L25
	} else {
		goto L29
	}
L28:
	;
	v103 = v93
	v104 = v92
	goto L25
L29:
	;
	v96 = int32(1)
	if v93 == v92 {
		v88 = v88 + v96
		v89 = v89 + v96
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v109 = v65 + int32(1)
	if v109 != v61 {
		v65 = v109
		goto L22
	} else {
		goto L32
	}
L32:
	;
	goto L23
L33:
	;
	goto L15
L34:
	;
	if v136 == int32(0) {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	F_initStringInfo(m, v14+int32(32))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	goto L1
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg_internal(m, int32(_a_F_generate_relation_name_0), v14)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_generate_relation_name_1), int32(_a_F_generate_relation_name_2), int32(_a_F_generate_relation_name_3))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v172 = v14 + int32(32)
	F_initStringInfo(m, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	if v169 == int32(0) {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v177 = F_quote_identifier(m, v169)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v177
	F_appendStringInfo(m, v172, int32(_a_F_generate_relation_name_4), v14+int32(16))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	goto L1
L45:
	;
	F_appendStringInfoString(m, v14+int32(32), v198)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	F_ReleaseCatCache(m, v18)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	m.G0 = v14 + int32(48)
	return v202
}
func F_getRelationDescription(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v13 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 == int32(0) {
			if l2 != 0 {
				m.G0 = v9 + int32(32)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
					F_errmsg_internal(m, int32(_a_F_getRelationDescription_0), v9)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_getRelationDescription_1), int32(_a_F_getRelationDescription_2), int32(_a_F_getRelationDescription_3))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
			v32 = v30 + v31
			v33 = F_RelationIsVisible(m, l1)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				if v33 != 0 {
					v39 = int32(0)
					v42 = F_quote_qualified_identifier(m, v39, v32+int32(4))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+119)))
						switch v45 - int32(73) {
						case 0, 32:
							v56 = int32(_a_F_getRelationDescription_4)
						default:
							v56 = int32(_a_F_getRelationDescription_5)
						case 10:
							v56 = int32(_a_F_getRelationDescription_6)
						case 26:
							v56 = int32(_a_F_getRelationDescription_7)
						case 29:
							v56 = int32(_a_F_getRelationDescription_8)
						case 36:
							v56 = int32(_a_F_getRelationDescription_9)
						case 39, 41:
							v56 = int32(_a_F_getRelationDescription_10)
						case 43:
							v56 = int32(_a_F_getRelationDescription_11)
						case 45:
							v56 = int32(_a_F_getRelationDescription_12)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v42
						F_appendStringInfo(m, l0, v56, v9+int32(16))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							F_ReleaseCatCache(m, v13)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+68))
					v37 = F_get_namespace_name(m, v36)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v39 = v37
						v42 = F_quote_qualified_identifier(m, v39, v32+int32(4))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+119)))
							switch v45 - int32(73) {
							case 0, 32:
								v56 = int32(_a_F_getRelationDescription_4)
							default:
								v56 = int32(_a_F_getRelationDescription_5)
							case 10:
								v56 = int32(_a_F_getRelationDescription_6)
							case 26:
								v56 = int32(_a_F_getRelationDescription_7)
							case 29:
								v56 = int32(_a_F_getRelationDescription_8)
							case 36:
								v56 = int32(_a_F_getRelationDescription_9)
							case 39, 41:
								v56 = int32(_a_F_getRelationDescription_10)
							case 43:
								v56 = int32(_a_F_getRelationDescription_11)
							case 45:
								v56 = int32(_a_F_getRelationDescription_12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v42
							F_appendStringInfo(m, l0, v56, v9+int32(16))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								F_ReleaseCatCache(m, v13)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_getRelationIdentity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v13 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 == int32(0) {
			if l3 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
					F_errmsg_internal(m, int32(_a_F_getRelationIdentity_0), v9)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_getRelationIdentity_1), int32(_a_F_getRelationIdentity_2), int32(_a_F_getRelationIdentity_3))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				if l2 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
				}
				m.G0 = v9 + int32(32)
				return
			}
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
			v25 = v23 + v24
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
			v27 = F_get_namespace_name_or_temp(m, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v30 = v25 + int32(4)
				v31 = F_quote_qualified_identifier(m, v27, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_appendStringInfoString(m, l0, v31)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						if l2 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v27
							v36 = F_pstrdup(m, v30)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v36
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v36
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v40
								v46 = F_list_make2_impl(m, v9+int32(20), v9+int32(16))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v46
									F_ReleaseCatCache(m, v13)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						} else {
							F_ReleaseCatCache(m, v13)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_get_relation_constraint_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v9 = m.G0
	v11 = v9 - int32(192)
	m.G0 = v11
	v15 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v20 = v11 + int32(16)
		F_ScanKeyInit(m, v20, int32(9), int32(3), int32(184), base.I64_extend_i32_u(l0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			F_ScanKeyInit(m, v11+int32(72), int32(10), int32(3), int32(184), int64(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_ScanKeyInit(m, v11+int32(128), int32(2), int32(3), int32(62), base.I64_extend_i32_u(l1))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v47 = F_systable_beginscan(m, v15, int32(2665), int32(1), int32(0), int32(3), v20)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = F_systable_getnext(m, v47)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 != 0 {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
								v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+22)))
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v51+v52)))
								v55 = v54
							} else {
								v55 = int32(0)
							}
							F_systable_endscan(m, v47)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if l2|v55 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67137668))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											v68 = F_get_rel_name(m, l0)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v68
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
												F_errmsg(m, int32(_a_F_get_relation_constraint_oid_0), v11)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_get_relation_constraint_oid_1), int32(1240), int32(_a_F_get_relation_constraint_oid_2))
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									}
								} else {
									F_relation_close(m, v15, int32(1))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(192)
										return v55
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_preprocess_relation_rtes(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v221 int32
	_ = v221
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v18 == int32(0) {
		v221 = v17
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(48)
	return v221
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v21 <= int32(0) {
		v221 = v17
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = int32(0)
	v29 = v17
	goto L4
L4:
	;
	v38 = v26 + int32(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v26<<(uint(int32(2))%32))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	if v45 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v221 = v208
	goto L1
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v50 = F_table_open(m, v48, int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v208 = v29
	goto L8
L8:
	;
	if v38 != v21 {
		v26 = v38
		v29 = v208
		goto L4
	} else {
		goto L49
	}
L9:
	;
	return int32(0)
L10:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+20)))
	if v54 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+126)))
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+20)) = uint8(v58)
	goto L13
L12:
	;
	goto L13
L13:
	;
	F_get_relation_notnullatts(m, l0, v50)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+24))
	if v63 == int32(0) {
		v193 = v29
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_relation_close(m, v50, int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L9
	} else {
		goto L48
	}
L16:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+18)))
	if v66 != int32(1) {
		v193 = v29
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v69 = int32(0)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v69 < v71 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v75 = v69
	v77 = v69
	v79 = v71
	goto L21
L19:
	;
	v136 = v69
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l0
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	v149 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v149
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v148
	if v136 != 0 {
		goto L35
	} else {
		goto L36
	}
L21:
	;
	v91 = v62 + v79<<(uint(int32(3))%32) + v75*int32(100)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+118)))
	if v92 == int32(118) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v136 = v130
	goto L20
L23:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v126 < v131 {
		v75 = v126
		v77 = v130
		v79 = v131
		goto L21
	} else {
		goto L34
	}
L24:
	;
	v96 = v75 + int32(1)
	v97 = F_build_generation_expression(m, v50, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v110 = v75 + int32(1)
	v111 = base.I32_extend16_s(v110)
	v113 = v91 + int32(28)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+68))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113)+76))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v113)+96))
	v118 = F_makeVar(m, v38, v111, v114, v115, v116, int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L9
	} else {
		goto L31
	}
L27:
	;
	F_ChangeVarNodes(m, v97, int32(1), v38)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v103 = int32(0)
	v105 = F_makeTargetEntry(m, v97, base.I32_extend16_s(v96), v103, v103)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v107 = F_lappend(m, v77, v105)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v126 = v96
	v130 = v107
	goto L23
L31:
	;
	v120 = int32(0)
	v122 = F_makeTargetEntry(m, v118, v111, v120, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	v124 = F_lappend(m, v77, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	v126 = v110
	v130 = v124
	goto L23
L34:
	;
	goto L22
L35:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v163 = v157<<(uint(int32(2))%32) + int32(4)
	goto L37
L36:
	;
	v163 = int32(4)
	goto L37
L37:
	;
	v164 = F_palloc0(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v164
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v29)+108))
	if v167 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = int32(1)
	goto L41
L40:
	;
	goto L41
L41:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v29)+84))
	if v170 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v178 = int32(0)
	v183 = F_replace_rte_variables(m, v29, v38, v178, int32(899), v15+int32(8), v178)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L9
	} else {
		goto L46
	}
L43:
	;
	v177 = int32(0)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v170)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v170)+36)) = int32(0)
	v177 = v174
	goto L42
L46:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v183)+84))
	if v185 == int32(0) {
		v193 = v183
		goto L15
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+36)) = v177
	v193 = v183
	goto L15
L48:
	;
	v208 = v193
	goto L8
L49:
	;
	goto L5
}
func F_rebuild_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int64
	_ = v103
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v442 int32
	_ = v442
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int64
	_ = v498
	var v504 int32
	_ = v504
	var v524 int32
	_ = v524
	var v528 int64
	_ = v528
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int64
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v662 int64
	_ = v662
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v700 int64
	_ = v700
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v882 float64
	_ = v882
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 float64
	_ = v897
	var v898 float64
	_ = v898
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 float64
	_ = v906
	var v907 float64
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 float64
	_ = v910
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v923 float64
	_ = v923
	var v926 int32
	_ = v926
	var v927 float64
	_ = v927
	var v933 float64
	_ = v933
	var v939 int32
	_ = v939
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 float64
	_ = v950
	var v951 float64
	_ = v951
	var v968 int32
	_ = v968
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1061 int32
	_ = v1061
	var v1070 int32
	_ = v1070
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1099 int32
	_ = v1099
	var v1119 int32
	_ = v1119
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 float64
	_ = v1145
	var v1146 float64
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1161 int32
	_ = v1161
	var v1162 float64
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1187 int64
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1197 float64
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1231 int32
	_ = v1231
	var v1235 int32
	_ = v1235
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1268 int32
	_ = v1268
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1344 int64
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1359 int32
	_ = v1359
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1392 int64
	_ = v1392
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1473 int32
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1515 int32
	_ = v1515
	var v1516 int64
	_ = v1516
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1529 int32
	_ = v1529
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1540 int32
	_ = v1540
	var v1543 int32
	_ = v1543
	var v1570 int32
	_ = v1570
	var v1572 int32
	_ = v1572
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1586 int32
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1594 int32
	_ = v1594
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1626 int32
	_ = v1626
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1653 int32
	_ = v1653
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1666 int32
	_ = v1666
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1676 int32
	_ = v1676
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1718 int32
	_ = v1718
	var v1721 int32
	_ = v1721
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1728 int32
	_ = v1728
	var v1735 int32
	_ = v1735
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1781 int32
	_ = v1781
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1837 int32
	_ = v1837
	var v1839 int32
	_ = v1839
	var v1866 int32
	_ = v1866
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1883 int32
	_ = v1883
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1907 int32
	_ = v1907
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1926 int32
	_ = v1926
	var v1933 int32
	_ = v1933
	var v1940 int32
	_ = v1940
	var v1960 int32
	_ = v1960
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1984 int32
	_ = v1984
	var v1988 int32
	_ = v1988
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1994 int32
	_ = v1994
	var v2004 int32
	_ = v2004
	var v2049 int64
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2057 int64
	_ = v2057
	var v2060 int64
	_ = v2060
	var v2065 int32
	_ = v2065
	var v2068 int32
	_ = v2068
	var v2072 int64
	_ = v2072
	var v2079 int64
	_ = v2079
	var v2084 int32
	_ = v2084
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2094 int32
	_ = v2094
	var v2098 int32
	_ = v2098
	var v2118 int32
	_ = v2118
	var v2122 int32
	_ = v2122
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2133 int32
	_ = v2133
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2161 int32
	_ = v2161
	var v2163 int32
	_ = v2163
	var v2166 int32
	_ = v2166
	var v2175 int32
	_ = v2175
	var v2195 int32
	_ = v2195
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2204 int32
	_ = v2204
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2235 int32
	_ = v2235
	var v2236 int64
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int64
	_ = v2244
	var v2247 int64
	_ = v2247
	var v2252 int32
	_ = v2252
	var v2255 int32
	_ = v2255
	var v2259 int64
	_ = v2259
	var v2266 int64
	_ = v2266
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2283 int32
	_ = v2283
	var v2285 int32
	_ = v2285
	var v2290 int32
	_ = v2290
	var v2294 int32
	_ = v2294
	var v2299 int32
	_ = v2299
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2305 int32
	_ = v2305
	var v2309 int32
	_ = v2309
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2327 int32
	_ = v2327
	var v2335 int32
	_ = v2335
	var v2360 int32
	_ = v2360
	var v2362 int32
	_ = v2362
	var v2366 int32
	_ = v2366
	var v2369 int32
	_ = v2369
	var v2372 int32
	_ = v2372
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2380 int64
	_ = v2380
	var v2384 int32
	_ = v2384
	var v2391 int32
	_ = v2391
	var v2396 int32
	_ = v2396
	var v2399 int32
	_ = v2399
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2408 int32
	_ = v2408
	var v2410 int32
	_ = v2410
	var v2411 int32
	_ = v2411
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2419 int32
	_ = v2419
	var v2424 int32
	_ = v2424
	var v2428 int32
	_ = v2428
	var v2433 int32
	_ = v2433
	var v2435 int32
	_ = v2435
	var v2436 int32
	_ = v2436
	var v2439 int32
	_ = v2439
	var v2443 int32
	_ = v2443
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2461 int32
	_ = v2461
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2474 int32
	_ = v2474
	var v2476 int32
	_ = v2476
	var v2479 int32
	_ = v2479
	var v2482 int32
	_ = v2482
	var v2485 int32
	_ = v2485
	var v2487 int32
	_ = v2487
	var v2490 int32
	_ = v2490
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2530 int32
	_ = v2530
	var v2535 int32
	_ = v2535
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2544 int32
	_ = v2544
	var v2549 int32
	_ = v2549
	var v2578 int32
	_ = v2578
	var v2582 int32
	_ = v2582
	var v2587 int32
	_ = v2587
	var v2591 int32
	_ = v2591
	var v2598 int32
	_ = v2598
	var v2603 int32
	_ = v2603
	var v2607 int32
	_ = v2607
	var v2616 int32
	_ = v2616
	var v2621 int32
	_ = v2621
	var v2625 int32
	_ = v2625
	var v2631 int32
	_ = v2631
	var v2636 int32
	_ = v2636
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2673 int32
	_ = v2673
	var v2678 int32
	_ = v2678
	v5 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(1920)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+92))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+84))
	if l3 == v5 {
		v442 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 != 0 {
		goto L83
	} else {
		goto L84
	}
L2:
	;
	F_BecomeLockGroupLeader(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v40 = F_palloc0(m, int32(12))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[0])) = v40
	v45 = F_dsm_create(m, int32(_a_F_rebuild_relation_0), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v47 = int32(_a_F_rebuild_relation_1)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v45
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
	v51 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+16)) = uint8(v51)
	*(*int64)(unsafe.Add(mBase, uint32(v50)+8)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v50))) = uint8(v51)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[0]))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	F_SharedFileSetInit(m, v50+int32(20), v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+72)) = int32(-1)
	v66 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v50)+76)), uint32(v66))
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+80)) = v70
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+96)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v50)+84)) = v73
	v77 = v50 + int32(100)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v77))), uint32(v66))
	*(*int64)(unsafe.Add(mBase, uint32(v77)+4)) = int64(-1)
	goto L8
L8:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+112)) = v84
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+116)) = v87
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+88)) = v90
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+92)) = v93
	v98 = (v50 + int32(151)) & int32(-32)
	v100 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v98))), uint32(v100))
	v103 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v98)+16)) = v103
	*(*int64)(unsafe.Add(mBase, uint32(v98)+4)) = v103
	v107 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v98)+36)) = uint16(v107)
	*(*int64)(unsafe.Add(mBase, uint32(v98)+24)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v98)+32)) = int32(_a_F_rebuild_relation_2)
	goto L9
L9:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[7]))
	F_shm_mq_set_receiver(m, v98, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[0]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v123 = F_shm_mq_attach(m, v98, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = v123
	v129 = v28 + int32(240)
	base.MemoryFill(m, v129, int32(0), int32(1472))
	v133 = F_get_rel_name(m, v30)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+224)) = v133
	v140 = F_pg_snprintf(m, v129, int32(96), int32(_a_F_rebuild_relation_3), v28+int32(224))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v147 = F_pg_snprintf(m, v28+int32(336), int32(96), int32(_a_F_rebuild_relation_4), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+440)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+432)) = int64(8589934595)
	v158 = F_pg_snprintf(m, v28+int32(444), int32(1024), int32(_a_F_rebuild_relation_5), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v165 = F_pg_snprintf(m, v28+int32(1468), int32(96), int32(_a_F_rebuild_relation_6), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v167 = int32(_a_F_rebuild_relation_1)
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[0]))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v170 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v169)+12)))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1568)) = v170
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1704)) = v173
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[0]))
	v177 = F_RegisterDynamicBackgroundWorker(m, v129, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	if v177 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v180 = v50 + int32(76)
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[0]))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+8))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+8)) = v184
	goto L21
L19:
	;
	goto L20
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L3
	} else {
		goto L78
	}
L21:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[8]))
	if v212 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L3
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[0]))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	if v217 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L25
L27:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[9]))
	v395 = F_WaitLatch(m, v391, int32(33), int32(-1), int32(134217734))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L3
	} else {
		goto L75
	}
L28:
	;
	F_ConditionVariablePrepareToSleep(m, v77)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L3
	} else {
		goto L42
	}
L29:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v223 = F_GetBackgroundWorkerPid(m, v220, v28+int32(1808))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[0]))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+8))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v255 = F_shm_mq_get_sender(m, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L3
	} else {
		goto L40
	}
L31:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[0]))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+8))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v229 = F_shm_mq_get_sender(m, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L3
	} else {
		goto L33
	}
L32:
	;
	switch v223 {
	case 0:
		goto L30
	default:
		goto L27
	case 2:
		goto L31
	}
L33:
	;
	if v229 != 0 {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(_a_F_rebuild_relation_7), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	F_errhint(m, int32(_a_F_rebuild_relation_8), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_rebuild_relation_9), int32(3895), int32(_a_F_rebuild_relation_10))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	if v255 == int32(0) {
		goto L27
	} else {
		goto L41
	}
L41:
	;
	goto L28
L42:
	;
	goto L43
L43:
	;
	v288 = base.AtomicRmwXchg32(m, v180, int32(0), int32(1))
	if v288 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L3
	} else {
		goto L53
	}
L45:
	;
	F_s_lock(m, v180, int32(_a_F_rebuild_relation_11))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L3
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v293 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v50)+76)), uint32(v293))
	if v292 == v293 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	F_ConditionVariableSleep(m, v77, int32(134217776))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L3
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L44
L52:
	;
	goto L43
L53:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[0]))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+24))
	v308 = v306 + int32(100)
	F_ConditionVariablePrepareToSleep(m, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v312 = v306 + int32(76)
	goto L55
L55:
	;
	v340 = base.AtomicRmwXchg32(m, v312, int32(0), int32(1))
	if v340 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L3
	} else {
		goto L65
	}
L57:
	;
	F_s_lock(m, v312, int32(_a_F_rebuild_relation_11))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L3
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v306)+72))
	v345 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v306)+76)), uint32(v345))
	if v344 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	F_ConditionVariableSleep(m, v308, int32(134217776))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L3
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	goto L56
L64:
	;
	goto L55
L65:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v306)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+196)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+192)) = v353
	v358 = v28 + int32(240)
	v363 = F_pg_snprintf(m, v358, int32(1024), int32(_a_F_rebuild_relation_12), v28+int32(192))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	v367 = int32(0)
	v369 = F_BufFileOpenFileSet(m, v306+int32(20), v358, v367, v367)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	F_BufFileReadExact(m, v369, v28+int32(1808), int32(4))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v28)+1808))
	v377 = F_palloc(m, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v28)+1808))
	F_BufFileReadExact(m, v369, v377, v379)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	F_BufFileClose(m, v369)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	v384 = F_RestoreSnapshot(m, v377)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	F_pfree(m, v377)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	F_PushActiveSnapshot(m, v384)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	v442 = v384
	goto L1
L75:
	;
	if v395&int32(1) == int32(0) {
		goto L21
	} else {
		goto L76
	}
L76:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[9]))
	v403 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v402))) = v403
	v408 = base.AtomicRmwOr32(m, v403, int32(_a_F_rebuild_relation_13), v403)
	goto L77
L77:
	;
	goto L21
L78:
	;
	F_errcode(m, int32(_a_F_rebuild_relation_14))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(_a_F_rebuild_relation_15), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L3
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+208)) = int32(_a_F_rebuild_relation_16)
	F_errhint(m, int32(_a_F_rebuild_relation_17), v28+int32(208))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L3
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_rebuild_relation_9), int32(3803), int32(_a_F_rebuild_relation_18))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L3
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_mark_index_clustered(m, l0, v457, int32(1))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L3
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v462 = int32(*(*int8)(unsafe.Add(mBase, uint32(v461)+118)))
	v464 = F_make_new_heap(m, v30, v32, v33, v462, int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L3
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	v467 = F_table_open(m, v464, int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L3
	} else {
		goto L88
	}
L88:
	;
	if l3 != 0 {
		goto L96
	} else {
		goto L97
	}
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2665 = m.ExcPending
	if v2665 != 0 {
		goto L3
	} else {
		goto L457
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2625 = m.ExcPending
	if v2625 != 0 {
		goto L3
	} else {
		goto L454
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2607 = m.ExcPending
	if v2607 != 0 {
		goto L3
	} else {
		goto L451
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2591 = m.ExcPending
	if v2591 != 0 {
		goto L3
	} else {
		goto L448
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2578 = m.ExcPending
	if v2578 != 0 {
		goto L3
	} else {
		goto L445
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2539 = m.ExcPending
	if v2539 != 0 {
		goto L3
	} else {
		goto L442
	}
L95:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2522 = m.ExcPending
	if v2522 != 0 {
		goto L3
	} else {
		goto L439
	}
L96:
	;
	v471 = F_table_open(m, int32(2604), int32(3))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L3
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v662 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1888)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1728)) = v662
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1872)) = v662
	F_getrusage(m, v28+int32(256))
	mBase = m.M
	F_gettimeofday(m, v28+int32(240))
	mBase = m.M
	goto L124
L99:
	;
	v475 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1880)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1876)) = v464
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1872)) = int32(1259)
	v483 = v28 + int32(1808)
	F_ScanKeyInit(m, v483, int32(2), int32(3), int32(184), base.I64_extend_i32_u(v30))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L3
	} else {
		goto L101
	}
L101:
	;
	v491 = int32(1)
	v494 = F_systable_beginscan(m, v471, int32(2656), v491, int32(0), v491, v483)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L3
	} else {
		goto L102
	}
L102:
	;
	v496 = F_systable_getnext(m, v494)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L3
	} else {
		goto L103
	}
L103:
	;
	if v496 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v498 = base.I64_extend_i32_u(v464)
	v504 = v496
	goto L107
L105:
	;
	goto L106
L106:
	;
	F_systable_endscan(m, v494)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L3
	} else {
		goto L120
	}
L107:
	;
	v524 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1768)) = v524
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1752)) = uint8(v524)
	v528 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1744)) = v528
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1736)) = v528
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1728)) = v528
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v504)+16))
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+22)))
	v538 = F_GetNewOidWithIndex(m, v471, int32(2657), int32(1))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L3
	} else {
		goto L109
	}
L108:
	;
	goto L106
L109:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1784)) = v498
	v541 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+1772)) = uint16(v541)
	v543 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+1768)) = uint16(v543)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1776)) = base.I64_extend_i32_u(v538)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v471)+52))
	v554 = F_heap_modify_tuple(m, v504, v547, v28+int32(1776), v28+int32(1772), v28+int32(1768))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L3
	} else {
		goto L110
	}
L110:
	;
	F_CatalogTupleInsert(m, v471, v554)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L3
	} else {
		goto L111
	}
L111:
	;
	v559 = v534 + v535
	v560 = int64(*(*int16)(unsafe.Add(mBase, uint32(v559)+8)))
	v563 = F_SearchSysCache2(m, int32(7), v498, v560&int64(4294967295))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L3
	} else {
		goto L112
	}
L112:
	;
	if v563 == int32(0) {
		goto L95
	} else {
		goto L113
	}
L113:
	;
	v567 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1900)) = uint8(v567)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+336)) = int64(1)
	v571 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1740)) = uint8(v571)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v475)+52))
	v580 = F_heap_modify_tuple(m, v563, v573, v28+int32(240), v28+int32(1888), v28+int32(1728))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L3
	} else {
		goto L114
	}
L114:
	;
	F_CatalogTupleUpdate(m, v475, v580+int32(4), v580)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L3
	} else {
		goto L115
	}
L115:
	;
	F_ReleaseCatCache(m, v563)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1724)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1720)) = v538
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1716)) = int32(2604)
	F_recordDependencyOn(m, v28+int32(1716), v28+int32(1872), int32(97))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L3
	} else {
		goto L117
	}
L117:
	;
	v600 = F_systable_getnext(m, v494)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L3
	} else {
		goto L118
	}
L118:
	;
	if v600 != 0 {
		v504 = v600
		goto L107
	} else {
		goto L119
	}
L119:
	;
	goto L108
L120:
	;
	F_relation_close(m, v471, int32(3))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L3
	} else {
		goto L121
	}
L121:
	;
	F_relation_close(m, v475, int32(3))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L3
	} else {
		goto L122
	}
L122:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L3
	} else {
		goto L123
	}
L123:
	;
	goto L98
L124:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v674)+68))
	v676 = F_get_namespace_name(m, v675)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L3
	} else {
		goto L125
	}
L125:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v678)+112))
	if v679 == int32(0) {
		v699 = v5
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v700 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1864)) = v700
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1856)) = v700
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1848)) = v700
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1840)) = v700
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1832)) = v700
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1824)) = v700
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1816)) = v700
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1808)) = v700
	v720 = F_vacuum_get_cutoffs(m, l0, v28+int32(1808), v28+int32(1776))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L3
	} else {
		goto L134
	}
L127:
	;
	if v442 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v684 = int32(4)
	goto L130
L129:
	;
	v684 = int32(8)
	goto L130
L130:
	;
	F_LockRelationOid(m, v679, v684)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v687)+112))
	if base.B2i32(v688 == int32(0))|v442 != 0 {
		v699 = v5
		goto L126
	} else {
		goto L132
	}
L132:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v467)+48))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v692)+112))
	if v693 == int32(0) {
		v699 = v5
		goto L126
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v467)+264)) = v688
	v699 = int32(1)
	goto L126
L134:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)+136))
	if v723 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v722)+140))
	if v740 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L136:
	;
	v726 = int32(3)
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v28)+1792))
	if base.B2i32(base.Ui32(v723) < base.Ui32(v726))|base.B2i32(base.Ui32(v728) < base.Ui32(v726)) == int32(0) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1792)) = v723
	goto L135
L138:
	;
	if v728-v723 < int32(0) {
		goto L137
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	if base.Ui32(v723) <= base.Ui32(v728) {
		goto L135
	} else {
		goto L142
	}
L141:
	;
	goto L135
L142:
	;
	goto L137
L143:
	;
	if l2 != 0 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v28)+1796))
	if int32(0) <= v743-v740 {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1796)) = v740
	goto L143
L146:
	;
	v750 = int32(17)
	goto L148
L147:
	;
	v750 = int32(13)
	goto L148
L148:
	;
	if l1 != 0 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v28)+1784))
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+124))
	m.T0[v1131].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, v467, l1, v1099, v1119, v442, v28+int32(1792), v28+int32(1796), v28+int32(1888), v28+int32(1728), v28+int32(1872))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L3
	} else {
		goto L197
	}
L150:
	;
	F_errfinish(m, int32(_a_F_rebuild_relation_9), v1070, int32(_a_F_rebuild_relation_19))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L3
	} else {
		goto L196
	}
L151:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v751)+84))
	if v752 == int32(403) {
		goto L155
	} else {
		goto L156
	}
L152:
	;
	goto L153
L153:
	;
	v1046 = int32(0)
	v1048 = F_errstart(m, v750, v1046)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L3
	} else {
		goto L193
	}
L154:
	;
	v1030 = F_errstart(m, v750, int32(0))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L3
	} else {
		goto L190
	}
L155:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v757 = m.G0
	v759 = v757 - int32(112)
	m.G0 = v759
	v761 = int32(1)
	v763 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_rebuild_relation[10])))
	if v763 != v761 {
		v968 = v761
		goto L158
	} else {
		goto L159
	}
L156:
	;
	goto L157
L157:
	;
	v1006 = int32(0)
	v1008 = F_errstart(m, v750, v1006)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L3
	} else {
		goto L187
	}
L158:
	;
	m.G0 = v759 + int32(112)
	if v968 != 0 {
		goto L154
	} else {
		goto L186
	}
L159:
	;
	v767 = F_palloc0(m, int32(168))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L3
	} else {
		goto L160
	}
L160:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v767))) = int64(4294967363)
	v772 = F_palloc0(m, int32(128))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L3
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v772))) = int32(268)
	v777 = F_palloc0(m, int32(400))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L3
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v777)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v777)+8)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v777)+4)) = v767
	*(*int32)(unsafe.Add(mBase, uint32(v777))) = int32(269)
	v786 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v777)+360)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v777)+300)) = v786
	v791 = F_palloc0(m, int32(8))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L3
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v791))) = int32(275)
	*(*int32)(unsafe.Add(mBase, uint32(v759)+12)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v759)+20)) = v791
	v800 = F_list_make1_impl(m, int32(1), v759+int32(12))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L3
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v777)+92)) = v800
	v804 = F_palloc0(m, int32(136))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L3
	} else {
		goto L165
	}
L165:
	;
	v806 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v804)+24)) = v806
	*(*int32)(unsafe.Add(mBase, uint32(v804)+16)) = v755
	*(*int32)(unsafe.Add(mBase, uint32(v804)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v804))) = int32(101)
	v813 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v804)+124)) = uint16(v813)
	v815 = int32(_a_F_rebuild_relation_20)
	*(*uint16)(unsafe.Add(mBase, uint32(v804)+20)) = uint16(v815)
	*(*int32)(unsafe.Add(mBase, uint32(v759)+8)) = v804
	*(*int32)(unsafe.Add(mBase, uint32(v759)+16)) = v804
	v822 = F_list_make1_impl(m, v806, v759+int32(8))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L3
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v767)+52)) = v822
	v827 = F_addRTEPermissionInfo(m, v767+int32(56), v804)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L3
	} else {
		goto L167
	}
L167:
	;
	F_setup_simple_rel_arrays(m, v777)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L3
	} else {
		goto L168
	}
L168:
	;
	v833 = F_build_simple_rel(m, v777, int32(1), int32(0))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L3
	} else {
		goto L169
	}
L169:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v833)+116))
	if v835 == int32(0) {
		v968 = v761
		goto L158
	} else {
		goto L170
	}
L170:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v835)+4))
	if v838 <= int32(0) {
		v968 = v761
		goto L158
	} else {
		goto L171
	}
L171:
	;
	v841 = int32(0)
	if v841 < v838 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v845 = v838
	goto L174
L173:
	;
	v845 = v841
	goto L174
L174:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v835)+12))
	v849 = v841
	goto L175
L175:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v846+v849<<(uint(int32(2))%32))))
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v875)+4))
	if v756 != v876 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v882 = *(*float64)(unsafe.Add(mBase, uint32(v833)+128))
	*(*float64)(unsafe.Add(mBase, uint32(v833)+16)) = v882
	v885 = F_get_relation_data_width(m, v755, int32(0))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L3
	} else {
		goto L181
	}
L177:
	;
	v878 = int32(1)
	v880 = v849 + v878
	if v845 != v880 {
		v849 = v880
		goto L175
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	goto L176
L180:
	;
	v968 = v878
	goto L158
L181:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v833)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v887)+32)) = v885
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v833)+124))
	*(*float64)(unsafe.Add(mBase, uint32(v777)+304)) = base.F64_convert_i32_u(v889)
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v875)+84))
	F_cost_qual_eval(m, v759+int32(96), v894, v777)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L3
	} else {
		goto L182
	}
L182:
	;
	v897 = *(*float64)(unsafe.Add(mBase, uint32(v759)+104))
	v898 = *(*float64)(unsafe.Add(mBase, uint32(v759)+96))
	v900 = v759 + int32(24)
	v901 = int32(0)
	v903 = F_create_seqscan_path(m, v777, v833, v901, v901)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L3
	} else {
		goto L183
	}
L183:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v903)+40))
	v906 = *(*float64)(unsafe.Add(mBase, uint32(v903)+56))
	v907 = *(*float64)(unsafe.Add(mBase, uint32(v833)+128))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v833)+40))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v908)+32))
	v910 = base.F64_add(v898, v897)
	v913 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[12]))
	v916 = m.G0
	v917 = int32(16)
	v918 = v916 - v917
	m.G0 = v918
	F_cost_tuplesort(m, v918+int32(8), v918, v907, v909, base.F64_add(v910, v910), v913, float64(-1))
	mBase = m.M
	v923 = *(*float64)(unsafe.Add(mBase, uint32(v918)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v900)+32)) = v907
	v926 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_rebuild_relation[13])))
	v927 = base.F64_add(v906, v923)
	*(*float64)(unsafe.Add(mBase, uint32(v900)+48)) = v927
	*(*int32)(unsafe.Add(mBase, uint32(v900)+40)) = v905 + (v926 ^ int32(1))
	v933 = *(*float64)(unsafe.Add(mBase, uint32(v918)))
	*(*float64)(unsafe.Add(mBase, uint32(v900)+56)) = base.F64_add(v927, v933)
	m.G0 = v918 + v917
	goto L184
L184:
	;
	v939 = int32(0)
	v948 = F_create_index_path(m, v777, v875, v939, v939, v939, v939, int32(1), v939, v939, float64(1), v939)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L3
	} else {
		goto L185
	}
L185:
	;
	v950 = *(*float64)(unsafe.Add(mBase, uint32(v759)+80))
	v951 = *(*float64)(unsafe.Add(mBase, uint32(v948)+56))
	v968 = base.F64_lt(v950, v951)
	goto L158
L186:
	;
	goto L157
L187:
	;
	if v1008 == int32(0) {
		v1099 = v1006
		goto L149
	} else {
		goto L188
	}
L188:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+160)) = v676
	v1015 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+168)) = v1013 + v1015
	*(*int32)(unsafe.Add(mBase, uint32(v28)+164)) = v1012 + v1015
	F_errmsg(m, int32(_a_F_rebuild_relation_21), v28+int32(160))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L3
	} else {
		goto L189
	}
L189:
	;
	v1070 = int32(1522)
	v1089 = int32(0)
	goto L150
L190:
	;
	if v1030 == int32(0) {
		v1099 = int32(1)
		goto L149
	} else {
		goto L191
	}
L191:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+144)) = v676
	*(*int32)(unsafe.Add(mBase, uint32(v28)+148)) = v1034 + int32(4)
	F_errmsg(m, int32(_a_F_rebuild_relation_22), v28+int32(144))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L3
	} else {
		goto L192
	}
L192:
	;
	v1070 = int32(1527)
	v1089 = int32(1)
	goto L150
L193:
	;
	if v1048 == int32(0) {
		v1099 = v1046
		goto L149
	} else {
		goto L194
	}
L194:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+128)) = v676
	*(*int32)(unsafe.Add(mBase, uint32(v28)+132)) = v1052 + int32(4)
	F_errmsg(m, int32(_a_F_rebuild_relation_23), v28+int32(128))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L3
	} else {
		goto L195
	}
L195:
	;
	v1070 = int32(1532)
	v1089 = int32(0)
	goto L150
L196:
	;
	v1099 = v1089
	goto L149
L197:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v28)+1796))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v28)+1792))
	v1136 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v467)+264)) = v1136
	v1139 = F_RelationGetNumberOfBlocksInFork(m, v467, v1136)
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L3
	} else {
		goto L198
	}
L198:
	;
	v1142 = F_errstart(m, v750, int32(0))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L3
	} else {
		goto L199
	}
L199:
	;
	if v1142 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1145 = *(*float64)(unsafe.Add(mBase, uint32(v28)+1728))
	v1146 = *(*float64)(unsafe.Add(mBase, uint32(v28)+1888))
	v1148 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L3
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	v1184 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L3
	} else {
		goto L208
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+120)) = v1148
	*(*float64)(unsafe.Add(mBase, uint32(v28)+112)) = v1146
	*(*float64)(unsafe.Add(mBase, uint32(v28)+104)) = v1145
	*(*int32)(unsafe.Add(mBase, uint32(v28)+100)) = v1144 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v676
	F_errmsg(m, int32(_a_F_rebuild_relation_24), v28+int32(96))
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L3
	} else {
		goto L204
	}
L204:
	;
	v1162 = *(*float64)(unsafe.Add(mBase, uint32(v28)+1872))
	v1165 = F_pg_rusage_show(m, v28+int32(240))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L3
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+88)) = v1165
	*(*float64)(unsafe.Add(mBase, uint32(v28)+80)) = v1162
	v1172 = F_errdetail(m, int32(_a_F_rebuild_relation_25), v28+int32(80))
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L3
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(_a_F_rebuild_relation_9), int32(1570), int32(_a_F_rebuild_relation_19))
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L3
	} else {
		goto L207
	}
L207:
	;
	goto L202
L208:
	;
	v1187 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v467)+56)))
	v1189 = F_SearchSysCacheCopy(m, int32(57), v1187, int64(0))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L3
	} else {
		goto L209
	}
L209:
	;
	if v1189 == int32(0) {
		goto L94
	} else {
		goto L210
	}
L210:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1189)+16))
	v1194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1193)+22)))
	v1195 = v1193 + v1194
	*(*int32)(unsafe.Add(mBase, uint32(v1195)+96)) = v1139
	v1197 = *(*float64)(unsafe.Add(mBase, uint32(v28)+1888))
	*(*float32)(unsafe.Add(mBase, uint32(v1195)+100)) = base.F32_demote_f64(v1197)
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v1200 != int32(1259) {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	F_pfree(m, v1189)
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L3
	} else {
		goto L217
	}
L212:
	;
	F_CatalogTupleUpdate(m, v1184, v1189+int32(4), v1189)
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L3
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	F_CacheInvalidateRelcacheByTuple(m, v1189)
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L3
	} else {
		goto L216
	}
L215:
	;
	goto L211
L216:
	;
	goto L211
L217:
	;
	F_relation_close(m, v1184, int32(3))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L3
	} else {
		goto L218
	}
L218:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L3
	} else {
		goto L219
	}
L219:
	;
	if v442 != 0 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L3
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	if l3 != 0 {
		goto L226
	} else {
		goto L227
	}
L223:
	;
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L3
	} else {
		goto L224
	}
L224:
	;
	goto L222
L225:
	;
	m.G0 = v28 + int32(1920)
	return
L226:
	;
	if l1 != 0 {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	goto L228
L228:
	;
	v2466 = int32(1)
	v2467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v2467) < base.Ui32(int32(_a_F_rebuild_relation_26)) {
		v2476 = v2466
		goto L429
	} else {
		goto L430
	}
L229:
	;
	F_relation_close(m, l1, int32(0))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L3
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v467)+56))
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v1225 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L3
	} else {
		goto L233
	}
L232:
	;
	goto L231
L233:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[14]))
	if v1231 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	if v1225 == int32(0) {
		goto L89
	} else {
		goto L238
	}
L235:
	;
	goto L234
L236:
	;
	v1235 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_rebuild_relation[15])))
	if v1235&int32(1) == int32(0) {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v1240 = int32(_a_F_rebuild_relation_27)
	v1242 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16]))
	v1243 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16])) = v1242 + v1243
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1231)))
	*(*int32)(unsafe.Add(mBase, uint32(v1231))) = v1246 + v1243
	v1250 = int32(0)
	v1252 = int32(_a_F_rebuild_relation_28)
	v1253 = base.AtomicRmwOr32(m, v1250, v1252, v1250)
	*(*int64)(unsafe.Add(mBase, uint32(v1231+int32(8))+232)) = int64(7)
	v1261 = base.AtomicRmwOr32(m, v1250, v1252, v1250)
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1231)))
	*(*int32)(unsafe.Add(mBase, uint32(v1231))) = v1262 + v1243
	v1268 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16]))
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16])) = v1268 - v1243
	goto L235
L238:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+4))
	if v1274 <= int32(0) {
		goto L89
	} else {
		goto L239
	}
L239:
	;
	v1277 = int32(0)
	v1286 = v1277
	v1289 = v1277
	goto L240
L240:
	;
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+12))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1304+v1286<<(uint(int32(2))%32))))
	v1310 = F_index_open(m, v1308, int32(4))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L3
	} else {
		goto L242
	}
L241:
	;
	if v1536 <= int32(0) {
		goto L89
	} else {
		goto L275
	}
L242:
	;
	v1312 = F_get_rel_name(m, v1308)
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L3
	} else {
		goto L243
	}
L243:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1310)+192))
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1316)+4))
	v1318 = F_get_rel_namespace(m, v1317)
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L3
	} else {
		goto L244
	}
L244:
	;
	v1321 = F_ChooseRelationName(m, v1312, int32(0), int32(_a_F_rebuild_relation_29), v1318, int32(0))
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L3
	} else {
		goto L245
	}
L245:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1310)+48))
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1324)+92))
	v1326 = F_index_create_copy(m, v467, int32(128), v1308, v1325, v1321)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L3
	} else {
		goto L246
	}
L246:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v467)+56))
	v1331 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L3
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1736)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1732)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1728)) = int32(1259)
	v1339 = v28 + int32(1808)
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1310)+192))
	v1344 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1343)+4)))
	F_ScanKeyInit(m, v1339, int32(9), int32(3), int32(184), v1344)
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L3
	} else {
		goto L248
	}
L248:
	;
	v1348 = int32(1)
	v1351 = F_systable_beginscan(m, v1331, int32(2665), v1348, int32(0), v1348, v1339)
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L3
	} else {
		goto L249
	}
L249:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+52))
	v1354 = F_systable_getnext(m, v1351)
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L3
	} else {
		goto L250
	}
L250:
	;
	if v1354 != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1359 = v1354
	goto L254
L252:
	;
	goto L253
L253:
	;
	F_systable_endscan(m, v1351)
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L3
	} else {
		goto L265
	}
L254:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1359)+16))
	v1384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1383)+22)))
	v1386 = v28 + int32(240)
	v1387 = int32(0)
	base.MemoryFill(m, v1386, v1387, int32(224))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1800)) = v1387
	v1392 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1792)) = v1392
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1784)) = v1392
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1776)) = v1392
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1912)) = v1387
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1904)) = v1392
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1896)) = v1392
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1888)) = v1392
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1383+v1384)+88))
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1310)+56))
	if v1407 == v1408 {
		goto L256
	} else {
		goto L257
	}
L255:
	;
	goto L253
L256:
	;
	v1412 = F_GetNewOidWithIndex(m, v1331, int32(2667), int32(1))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L3
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	v1445 = F_systable_getnext(m, v1351)
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L3
	} else {
		goto L263
	}
L259:
	;
	v1414 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1888)) = uint8(v1414)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+304)) = base.I64_extend_i32_u(v1328)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1896)) = uint8(v1414)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+320)) = base.I64_extend_i32_u(v1326)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1898)) = uint8(v1414)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+240)) = base.I64_extend_i32_u(v1412)
	v1428 = F_heap_modify_tuple(m, v1359, v1353, v1386, v28+int32(1776), v28+int32(1888))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L3
	} else {
		goto L260
	}
L260:
	;
	F_CatalogTupleInsert(m, v1331, v1428)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L3
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1880)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1876)) = v1412
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1872)) = int32(2606)
	F_recordDependencyOn(m, v28+int32(1872), v28+int32(1728), int32(97))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L3
	} else {
		goto L262
	}
L262:
	;
	goto L258
L263:
	;
	if v1445 != 0 {
		v1359 = v1445
		goto L254
	} else {
		goto L264
	}
L264:
	;
	goto L255
L265:
	;
	F_relation_close(m, v1331, int32(3))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L3
	} else {
		goto L266
	}
L266:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L3
	} else {
		goto L267
	}
L267:
	;
	v1479 = F_lappend_oid(m, v1289, v1326)
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L3
	} else {
		goto L268
	}
L268:
	;
	F_relation_close(m, v1310, int32(0))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L3
	} else {
		goto L269
	}
L269:
	;
	v1488 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[14]))
	if v1488 == int32(0) {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1535 = v1286 + int32(1)
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+4))
	if v1535 < v1536 {
		v1286 = v1535
		v1289 = v1479
		goto L240
	} else {
		goto L274
	}
L271:
	;
	goto L270
L272:
	;
	v1492 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_rebuild_relation[15])))
	if v1492&int32(1) == int32(0) {
		goto L271
	} else {
		goto L273
	}
L273:
	;
	v1497 = int32(_a_F_rebuild_relation_27)
	v1499 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16]))
	v1500 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16])) = v1499 + v1500
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1488)))
	*(*int32)(unsafe.Add(mBase, uint32(v1488))) = v1503 + v1500
	v1507 = int32(0)
	v1509 = int32(_a_F_rebuild_relation_28)
	v1510 = base.AtomicRmwOr32(m, v1507, v1509, v1507)
	v1515 = v1488 + int32(304)
	v1516 = *(*int64)(unsafe.Add(mBase, uint32(v1515)))
	*(*int64)(unsafe.Add(mBase, uint32(v1515))) = v1516 + int64(1)
	v1522 = base.AtomicRmwOr32(m, v1507, v1509, v1507)
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1488)))
	*(*int32)(unsafe.Add(mBase, uint32(v1488))) = v1523 + v1500
	v1529 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16]))
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16])) = v1529 - v1500
	goto L271
L274:
	;
	goto L241
L275:
	;
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+12))
	v1543 = int32(0)
	goto L276
L276:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v1540+v1543<<(uint(int32(2))%32))))
	if v1570 == l3 {
		goto L279
	} else {
		goto L280
	}
L277:
	;
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v1479)+12))
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v1590+v1543<<(uint(int32(2))%32))))
	if v1594 == int32(0) {
		goto L89
	} else {
		goto L290
	}
L278:
	;
	goto L277
L279:
	;
	if v1479 != 0 {
		goto L282
	} else {
		goto L283
	}
L280:
	;
	goto L281
L281:
	;
	v1588 = v1543 + int32(1)
	if v1536 != v1588 {
		v1543 = v1588
		goto L276
	} else {
		goto L289
	}
L282:
	;
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1479)+4))
	if v1543 < v1572 {
		goto L278
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L3
	} else {
		goto L286
	}
L285:
	;
	goto L284
L286:
	;
	F_errmsg_internal(m, int32(_a_F_rebuild_relation_30), int32(0))
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L3
	} else {
		goto L287
	}
L287:
	;
	F_errfinish(m, int32(_a_F_rebuild_relation_9), int32(3370), int32(_a_F_rebuild_relation_31))
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L3
	} else {
		goto L288
	}
L288:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L289:
	;
	goto L89
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+240)) = v467
	v1598 = F_CreateExecutorState(m)
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L3
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+248)) = v1598
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v467)+52))
	v1602 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1808)) = v1602
	v1606 = F_palloc0(m, int32(136))
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L3
	} else {
		goto L292
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1606)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1606))) = int32(101)
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v467)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1606)+16)) = v1612
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v467)+48))
	v1615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1614)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1606)+21)) = uint8(v1615)
	v1619 = F_addRTEPermissionInfo(m, v28+int32(1808), v1606)
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L3
	} else {
		goto L293
	}
L293:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v1601)))
	if int32(0) < v1621 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1626 = int32(0)
	v1631 = v1621
	v1632 = v1602
	goto L297
L295:
	;
	v1670 = v1598
	v1676 = v1602
	goto L296
L296:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v28)+1808))
	v1695 = F_getRTEPermissionInfo(m, v1694, v1606)
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L3
	} else {
		goto L304
	}
L297:
	;
	v1653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1601+v1626<<(uint(int32(3))%32))+34)))
	if v1653&int32(4) == int32(0) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	v1670 = v1668
	v1676 = v1664
	goto L296
L299:
	;
	v1660 = F_bms_add_member(m, v1632, v1626+int32(8))
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L3
	} else {
		goto L302
	}
L300:
	;
	v1663 = v1631
	v1664 = v1632
	goto L301
L301:
	;
	v1666 = v1626 + int32(1)
	if v1666 < v1663 {
		v1626 = v1666
		v1631 = v1663
		v1632 = v1664
		goto L297
	} else {
		goto L303
	}
L302:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v1601)))
	v1663 = v1662
	v1664 = v1660
	goto L301
L303:
	;
	goto L298
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1695)+36)) = v1676
	*(*int32)(unsafe.Add(mBase, uint32(v28)+76)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v28)+1776)) = v1606
	v1703 = F_list_make1_impl(m, int32(1), v28+int32(76))
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L3
	} else {
		goto L305
	}
L305:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v28)+1808))
	v1707 = F_bms_make_singleton(m, int32(1))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L3
	} else {
		goto L306
	}
L306:
	;
	F_ExecInitRangeTable(m, v1670, v1703, v1705, v1707)
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L3
	} else {
		goto L307
	}
L307:
	;
	v1712 = F_palloc0(m, int32(216))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L3
	} else {
		goto L308
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1712))) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+244)) = v1712
	v1718 = int32(0)
	F_InitResultRelInfo(m, v1712, v467, int32(1), v1718, v1718)
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		goto L3
	} else {
		goto L309
	}
L309:
	;
	F_ExecOpenIndices(m, v1712, int32(0))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L3
	} else {
		goto L310
	}
L310:
	;
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v1712)+12))
	if v1725 <= int32(0) {
		goto L93
	} else {
		goto L311
	}
L311:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1712)+16))
	v1735 = int32(0)
	goto L312
L312:
	;
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v1728+v1735<<(uint(int32(2))%32))))
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+56))
	if v1594 != v1759 {
		goto L314
	} else {
		goto L315
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+252)) = v1758
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+192))
	v1766 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1765)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+260)) = v1766
	v1769 = F_palloc_mul(m, int32(56), v1766)
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L3
	} else {
		goto L318
	}
L314:
	;
	v1762 = v1735 + int32(1)
	if v1725 != v1762 {
		v1735 = v1762
		goto L312
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	goto L313
L317:
	;
	goto L93
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+256)) = v1769
	v1772 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1765)+10)))
	if int32(0) < v1772 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1781 = int32(0)
	goto L322
L320:
	;
	goto L321
L321:
	;
	v1866 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+264)) = uint16(v1866)
	if v1866 < v1766 {
		goto L332
	} else {
		goto L333
	}
L322:
	;
	v1802 = v1781 << (uint(int32(2)) % 32)
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+212))
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(v1802+v1803)))
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+48))
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1807)+84))
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+208))
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v1809+v1802)))
	v1813 = F_IndexAmTranslateCompareType(m, int32(3), v1808, v1811, int32(0))
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		goto L3
	} else {
		goto L324
	}
L323:
	;
	goto L321
L324:
	;
	if v1813 == int32(0) {
		goto L92
	} else {
		goto L325
	}
L325:
	;
	v1818 = F_get_opfamily_member(m, v1811, v1805, v1805, base.I32_extend16_s(v1813))
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L3
	} else {
		goto L326
	}
L326:
	;
	if v1818 == int32(0) {
		goto L91
	} else {
		goto L327
	}
L327:
	;
	v1822 = F_get_opcode(m, v1818)
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L3
	} else {
		goto L328
	}
L328:
	;
	if v1822 == int32(0) {
		goto L90
	} else {
		goto L329
	}
L329:
	;
	v1828 = v1769 + v1781*int32(56)
	v1830 = v1781 + int32(1)
	F_ScanKeyInit(m, v1828, base.I32_extend16_s(v1830), v1813, v1822, int64(0))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L3
	} else {
		goto L330
	}
L330:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+248))
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1835+v1802)))
	*(*int32)(unsafe.Add(mBase, uint32(v1828)+12)) = v1837
	v1839 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1765)+10)))
	if v1830 < v1839 {
		v1781 = v1830
		goto L322
	} else {
		goto L331
	}
L331:
	;
	goto L323
L332:
	;
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+192))
	v1873 = v1871 + int32(48)
	v1874 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1766) {
		goto L336
	} else {
		goto L337
	}
L333:
	;
	goto L334
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+268)) = int32(1)
	v2049 = F_GetXLogInsertEndRecPtr(m)
	mBase = m.M
	v2050 = m.ExcPending
	if v2050 != 0 {
		goto L3
	} else {
		goto L361
	}
L335:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+264)) = uint16(v2004)
	goto L334
L336:
	;
	v1883 = v1874
	v1889 = int32(0)
	v1890 = v1874
	goto L339
L337:
	;
	v1933 = v1874
	v1940 = v1874
	goto L338
L338:
	;
	v1960 = v1933
	v1965 = v1874
	v1967 = v1940
	goto L355
L339:
	;
	v1907 = base.I32_extend16_s(v1890)
	v1910 = v1873 + v1883<<(uint(int32(1))%32)
	v1911 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1910))))
	if v1911 < v1907 {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	if v1766&int32(3) == int32(0) {
		v2004 = v1922
		goto L335
	} else {
		goto L354
	}
L341:
	;
	v1913 = v1907
	goto L343
L342:
	;
	v1913 = v1911
	goto L343
L343:
	;
	v1914 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1910)+2)))
	if v1914 < v1913 {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1916 = v1913
	goto L346
L345:
	;
	v1916 = v1914
	goto L346
L346:
	;
	v1917 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1910)+4)))
	if v1917 < v1916 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v1919 = v1916
	goto L349
L348:
	;
	v1919 = v1917
	goto L349
L349:
	;
	v1920 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1910)+6)))
	if v1920 < v1919 {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v1922 = v1919
	goto L352
L351:
	;
	v1922 = v1920
	goto L352
L352:
	;
	v1923 = int32(4)
	v1924 = v1883 + v1923
	v1926 = v1889 + v1923
	if v1926 != v1766&int32(_a_F_rebuild_relation_32) {
		v1883 = v1924
		v1889 = v1926
		v1890 = v1922
		goto L339
	} else {
		goto L353
	}
L353:
	;
	goto L340
L354:
	;
	v1933 = v1924
	v1940 = v1922
	goto L338
L355:
	;
	v1984 = base.I32_extend16_s(v1967)
	v1988 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1873+v1960<<(uint(int32(1))%32)))))
	if v1988 < v1984 {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	v2004 = v1990
	goto L335
L357:
	;
	v1990 = v1984
	goto L359
L358:
	;
	v1990 = v1988
	goto L359
L359:
	;
	v1991 = int32(1)
	v1994 = v1965 + v1991
	if v1994 != v1766&int32(3) {
		v1960 = v1960 + v1991
		v1965 = v1994
		v1967 = v1990
		goto L355
	} else {
		goto L360
	}
L360:
	;
	goto L356
L361:
	;
	F_XLogFlush(m, v2049)
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L3
	} else {
		goto L362
	}
L362:
	;
	v2053 = int32(0)
	v2055 = int32(_a_F_rebuild_relation_33)
	v2056 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[17]))
	v2057 = int64(0)
	v2060 = base.AtomicRmwCmpxchg64(m, v2056, int32(272), v2057, v2057)
	*(*int64)(unsafe.Add(mBase, _c_F_rebuild_relation[18])) = v2060
	v2065 = base.AtomicRmwOr32(m, v2053, int32(_a_F_rebuild_relation_34), v2053)
	v2068 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[17]))
	v2072 = base.AtomicRmwCmpxchg64(m, v2068, int32(264), v2057, v2057)
	*(*int64)(unsafe.Add(mBase, _c_F_rebuild_relation[19])) = v2072
	goto L365
L363:
	;
	F_process_concurrent_changes(m, v2079, v28+int32(240), int32(0))
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L3
	} else {
		goto L367
	}
L365:
	;
	goto L366
L366:
	;
	v2079 = *(*int64)(unsafe.Add(mBase, _c_F_rebuild_relation[18]))
	goto L363
L367:
	;
	F_LockRelationOid(m, v1224, int32(8))
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L3
	} else {
		goto L368
	}
L368:
	;
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+4))
	if v2088 <= int32(0) {
		goto L370
	} else {
		goto L371
	}
L369:
	;
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2158 = *(*int32)(unsafe.Add(mBase, uint32(v2157)+112))
	if v2158 != 0 {
		goto L378
	} else {
		goto L379
	}
L370:
	;
	v2133 = int32(0)
	goto L369
L371:
	;
	goto L372
L372:
	;
	v2094 = int32(0)
	v2098 = v1866
	goto L373
L373:
	;
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+12))
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v2118+v2098<<(uint(int32(2))%32))))
	v2124 = F_index_open(m, v2122, int32(8))
	mBase = m.M
	v2125 = m.ExcPending
	if v2125 != 0 {
		goto L3
	} else {
		goto L375
	}
L374:
	;
	v2133 = v2126
	goto L369
L375:
	;
	v2126 = F_lappend(m, v2094, v2124)
	mBase = m.M
	v2127 = m.ExcPending
	if v2127 != 0 {
		goto L3
	} else {
		goto L376
	}
L376:
	;
	v2129 = v2098 + int32(1)
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+4))
	if v2129 < v2130 {
		v2094 = v2126
		v2098 = v2129
		goto L373
	} else {
		goto L377
	}
L377:
	;
	goto L374
L378:
	;
	F_LockRelationOid(m, v2158, int32(8))
	mBase = m.M
	v2161 = m.ExcPending
	if v2161 != 0 {
		goto L3
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	F_TransferPredicateLocksToHeapRelation(m, l0)
	mBase = m.M
	v2163 = m.ExcPending
	if v2163 != 0 {
		goto L3
	} else {
		goto L382
	}
L381:
	;
	goto L380
L382:
	;
	if v2133 == int32(0) {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	F_list_free(m, v2133)
	mBase = m.M
	v2235 = m.ExcPending
	if v2235 != 0 {
		goto L3
	} else {
		goto L391
	}
L384:
	;
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(v2133)+4))
	if v2166 <= int32(0) {
		goto L383
	} else {
		goto L385
	}
L385:
	;
	v2175 = int32(0)
	goto L386
L386:
	;
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(v2133)+12))
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(v2195+v2175<<(uint(int32(2))%32))))
	F_TransferPredicateLocksToHeapRelation(m, v2199)
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L3
	} else {
		goto L388
	}
L387:
	;
	goto L383
L388:
	;
	F_relation_close(m, v2199, int32(0))
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L3
	} else {
		goto L389
	}
L389:
	;
	v2206 = v2175 + int32(1)
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(v2133)+4))
	if v2206 < v2207 {
		v2175 = v2206
		goto L386
	} else {
		goto L390
	}
L390:
	;
	goto L387
L391:
	;
	v2236 = F_GetXLogInsertEndRecPtr(m)
	mBase = m.M
	v2237 = m.ExcPending
	if v2237 != 0 {
		goto L3
	} else {
		goto L392
	}
L392:
	;
	F_XLogFlush(m, v2236)
	mBase = m.M
	v2239 = m.ExcPending
	if v2239 != 0 {
		goto L3
	} else {
		goto L393
	}
L393:
	;
	v2240 = int32(0)
	v2242 = int32(_a_F_rebuild_relation_33)
	v2243 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[17]))
	v2244 = int64(0)
	v2247 = base.AtomicRmwCmpxchg64(m, v2243, int32(272), v2244, v2244)
	*(*int64)(unsafe.Add(mBase, _c_F_rebuild_relation[18])) = v2247
	v2252 = base.AtomicRmwOr32(m, v2240, int32(_a_F_rebuild_relation_34), v2240)
	v2255 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[17]))
	v2259 = base.AtomicRmwCmpxchg64(m, v2255, int32(264), v2244, v2244)
	*(*int64)(unsafe.Add(mBase, _c_F_rebuild_relation[19])) = v2259
	goto L396
L394:
	;
	F_process_concurrent_changes(m, v2266, v28+int32(240), int32(1))
	mBase = m.M
	v2271 = m.ExcPending
	if v2271 != 0 {
		goto L3
	} else {
		goto L398
	}
L396:
	;
	goto L397
L397:
	;
	v2266 = *(*int64)(unsafe.Add(mBase, _c_F_rebuild_relation[18]))
	goto L394
L398:
	;
	v2272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2273 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2272)+118)))
	v2275 = int32(1)
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v2276) < base.Ui32(int32(_a_F_rebuild_relation_26)) {
		v2285 = v2275
		goto L400
	} else {
		goto L401
	}
L399:
	;
	v2290 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[14]))
	if v2290 == int32(0) {
		goto L404
	} else {
		goto L405
	}
L400:
	;
	goto L399
L401:
	;
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2280 = *(*int32)(unsafe.Add(mBase, uint32(v2279)+68))
	if v2280 == int32(99) {
		v2285 = v2275
		goto L400
	} else {
		goto L402
	}
L402:
	;
	v2283 = F_isTempToastNamespace(m, v2280)
	mBase = m.M
	v2285 = v2283
	goto L400
L403:
	;
	v2335 = int32(0)
	goto L407
L404:
	;
	goto L403
L405:
	;
	v2294 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_rebuild_relation[15])))
	if v2294&int32(1) == int32(0) {
		goto L404
	} else {
		goto L406
	}
L406:
	;
	v2299 = int32(_a_F_rebuild_relation_27)
	v2301 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16]))
	v2302 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16])) = v2301 + v2302
	v2305 = *(*int32)(unsafe.Add(mBase, uint32(v2290)))
	*(*int32)(unsafe.Add(mBase, uint32(v2290))) = v2305 + v2302
	v2309 = int32(0)
	v2311 = int32(_a_F_rebuild_relation_28)
	v2312 = base.AtomicRmwOr32(m, v2309, v2311, v2309)
	*(*int64)(unsafe.Add(mBase, uint32(v2290+int32(8))+232)) = int64(6)
	v2320 = base.AtomicRmwOr32(m, v2309, v2311, v2309)
	v2321 = *(*int32)(unsafe.Add(mBase, uint32(v2290)))
	*(*int32)(unsafe.Add(mBase, uint32(v2290))) = v2321 + v2302
	v2327 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16]))
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16])) = v2327 - v2302
	goto L404
L407:
	;
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+4))
	if v2335 < v2360 {
		goto L409
	} else {
		goto L410
	}
L408:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		goto L3
	} else {
		goto L416
	}
L409:
	;
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+12))
	v2366 = v2362 + v2335<<(uint(int32(2))%32)
	goto L411
L410:
	;
	v2366 = int32(0)
	goto L411
L411:
	;
	v2369 = *(*int32)(unsafe.Add(mBase, uint32(v1479)+4))
	if base.B2i32(v2366 == int32(0))|base.B2i32(v2369 <= v2335) != 0 {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	goto L408
L413:
	;
	v2372 = *(*int32)(unsafe.Add(mBase, uint32(v1479)+12))
	if v2372 == int32(0) {
		goto L412
	} else {
		goto L414
	}
L414:
	;
	v2378 = *(*int32)(unsafe.Add(mBase, uint32(v2372+v2335<<(uint(int32(2))%32))))
	v2379 = *(*int32)(unsafe.Add(mBase, uint32(v2366)))
	v2380 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1816)) = v2380
	*(*int64)(unsafe.Add(mBase, uint32(v28)+1808)) = v2380
	v2384 = int32(0)
	F_swap_relation_files(m, v2379, v2378, base.B2i32(v1224 == int32(1259)), v2384, int32(1), v2384, v2384, v28+int32(1808))
	mBase = m.M
	v2391 = m.ExcPending
	if v2391 != 0 {
		goto L3
	} else {
		goto L415
	}
L415:
	;
	v2335 = v2335 + int32(1)
	goto L407
L416:
	;
	F_relation_close(m, l0, int32(0))
	mBase = m.M
	v2399 = m.ExcPending
	if v2399 != 0 {
		goto L3
	} else {
		goto L417
	}
L417:
	;
	F_relation_close(m, v467, int32(0))
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		goto L3
	} else {
		goto L418
	}
L418:
	;
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v28)+244))
	F_ExecCloseIndices(m, v2403)
	mBase = m.M
	v2405 = m.ExcPending
	if v2405 != 0 {
		goto L3
	} else {
		goto L419
	}
L419:
	;
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	F_FreeExecutorState(m, v2406)
	mBase = m.M
	v2408 = m.ExcPending
	if v2408 != 0 {
		goto L3
	} else {
		goto L420
	}
L420:
	;
	F_pfree(m, v2403)
	mBase = m.M
	v2410 = m.ExcPending
	if v2410 != 0 {
		goto L3
	} else {
		goto L421
	}
L421:
	;
	v2411 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	F_pfree(m, v2411)
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L3
	} else {
		goto L422
	}
L422:
	;
	v2414 = int32(0)
	F_finish_heap_swap(m, v1224, v1223, v2285, v2414, v2414, int32(1), v2414, v1135, v1134, v2273)
	mBase = m.M
	v2419 = m.ExcPending
	if v2419 != 0 {
		goto L3
	} else {
		goto L423
	}
L423:
	;
	v2424 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[14]))
	if v2424 == int32(0) {
		goto L425
	} else {
		goto L426
	}
L424:
	;
	goto L225
L425:
	;
	goto L424
L426:
	;
	v2428 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_rebuild_relation[15])))
	if v2428&int32(1) == int32(0) {
		goto L425
	} else {
		goto L427
	}
L427:
	;
	v2433 = int32(_a_F_rebuild_relation_27)
	v2435 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16]))
	v2436 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16])) = v2435 + v2436
	v2439 = *(*int32)(unsafe.Add(mBase, uint32(v2424)))
	*(*int32)(unsafe.Add(mBase, uint32(v2424))) = v2439 + v2436
	v2443 = int32(0)
	v2445 = int32(_a_F_rebuild_relation_28)
	v2446 = base.AtomicRmwOr32(m, v2443, v2445, v2443)
	*(*int64)(unsafe.Add(mBase, uint32(v2424+int32(8))+232)) = int64(8)
	v2454 = base.AtomicRmwOr32(m, v2443, v2445, v2443)
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(v2424)))
	*(*int32)(unsafe.Add(mBase, uint32(v2424))) = v2455 + v2436
	v2461 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16]))
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_relation[16])) = v2461 - v2436
	goto L425
L428:
	;
	F_relation_close(m, l0, int32(0))
	mBase = m.M
	v2479 = m.ExcPending
	if v2479 != 0 {
		goto L3
	} else {
		goto L432
	}
L429:
	;
	goto L428
L430:
	;
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2471 = *(*int32)(unsafe.Add(mBase, uint32(v2470)+68))
	if v2471 == int32(99) {
		v2476 = v2466
		goto L429
	} else {
		goto L431
	}
L431:
	;
	v2474 = F_isTempToastNamespace(m, v2471)
	mBase = m.M
	v2476 = v2474
	goto L429
L432:
	;
	if l1 != 0 {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	F_relation_close(m, l1, int32(0))
	mBase = m.M
	v2482 = m.ExcPending
	if v2482 != 0 {
		goto L3
	} else {
		goto L436
	}
L434:
	;
	goto L435
L435:
	;
	F_relation_close(m, v467, int32(0))
	mBase = m.M
	v2485 = m.ExcPending
	if v2485 != 0 {
		goto L3
	} else {
		goto L437
	}
L436:
	;
	goto L435
L437:
	;
	v2487 = int32(1)
	F_finish_heap_swap(m, v30, v464, v2476, v699, int32(0), v2487, v2487, v1135, v1134, v462)
	mBase = m.M
	v2490 = m.ExcPending
	if v2490 != 0 {
		goto L3
	} else {
		goto L438
	}
L438:
	;
	goto L225
L439:
	;
	v2523 = int32(*(*int16)(unsafe.Add(mBase, uint32(v559)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+180)) = v464
	*(*int32)(unsafe.Add(mBase, uint32(v28)+176)) = v2523
	F_errmsg_internal(m, int32(_a_F_rebuild_relation_35), v28+int32(176))
	mBase = m.M
	v2530 = m.ExcPending
	if v2530 != 0 {
		goto L3
	} else {
		goto L440
	}
L440:
	;
	F_errfinish(m, int32(_a_F_rebuild_relation_9), int32(3715), int32(_a_F_rebuild_relation_36))
	mBase = m.M
	v2535 = m.ExcPending
	if v2535 != 0 {
		goto L3
	} else {
		goto L441
	}
L441:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L442:
	;
	v2540 = *(*int32)(unsafe.Add(mBase, uint32(v467)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v2540
	F_errmsg_internal(m, int32(_a_F_rebuild_relation_37), v28)
	mBase = m.M
	v2544 = m.ExcPending
	if v2544 != 0 {
		goto L3
	} else {
		goto L443
	}
L443:
	;
	F_errfinish(m, int32(_a_F_rebuild_relation_9), int32(1579), int32(_a_F_rebuild_relation_19))
	mBase = m.M
	v2549 = m.ExcPending
	if v2549 != 0 {
		goto L3
	} else {
		goto L444
	}
L444:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L445:
	;
	F_errmsg_internal(m, int32(_a_F_rebuild_relation_38), int32(0))
	mBase = m.M
	v2582 = m.ExcPending
	if v2582 != 0 {
		goto L3
	} else {
		goto L446
	}
L446:
	;
	F_errfinish(m, int32(_a_F_rebuild_relation_9), int32(3239), int32(_a_F_rebuild_relation_39))
	mBase = m.M
	v2587 = m.ExcPending
	if v2587 != 0 {
		goto L3
	} else {
		goto L447
	}
L447:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v1805
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v1811
	F_errmsg_internal(m, int32(_a_F_rebuild_relation_40), v28+int32(32))
	mBase = m.M
	v2598 = m.ExcPending
	if v2598 != 0 {
		goto L3
	} else {
		goto L449
	}
L449:
	;
	F_errfinish(m, int32(_a_F_rebuild_relation_9), int32(3266), int32(_a_F_rebuild_relation_39))
	mBase = m.M
	v2603 = m.ExcPending
	if v2603 != 0 {
		goto L3
	} else {
		goto L450
	}
L450:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+60)) = v1811
	*(*int32)(unsafe.Add(mBase, uint32(v28)+56)) = v1805
	*(*int32)(unsafe.Add(mBase, uint32(v28)+52)) = v1805
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v1813
	F_errmsg_internal(m, int32(_a_F_rebuild_relation_41), v28+int32(48))
	mBase = m.M
	v2616 = m.ExcPending
	if v2616 != 0 {
		goto L3
	} else {
		goto L452
	}
L452:
	;
	F_errfinish(m, int32(_a_F_rebuild_relation_9), int32(3271), int32(_a_F_rebuild_relation_39))
	mBase = m.M
	v2621 = m.ExcPending
	if v2621 != 0 {
		goto L3
	} else {
		goto L453
	}
L453:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v1818
	F_errmsg_internal(m, int32(_a_F_rebuild_relation_42), v28-int32(-64))
	mBase = m.M
	v2631 = m.ExcPending
	if v2631 != 0 {
		goto L3
	} else {
		goto L455
	}
L455:
	;
	F_errfinish(m, int32(_a_F_rebuild_relation_9), int32(3274), int32(_a_F_rebuild_relation_39))
	mBase = m.M
	v2636 = m.ExcPending
	if v2636 != 0 {
		goto L3
	} else {
		goto L456
	}
L456:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L457:
	;
	v2666 = F_get_rel_name(m, l3)
	mBase = m.M
	v2667 = m.ExcPending
	if v2667 != 0 {
		goto L3
	} else {
		goto L458
	}
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v2666
	F_errmsg_internal(m, int32(_a_F_rebuild_relation_43), v28+int32(16))
	mBase = m.M
	v2673 = m.ExcPending
	if v2673 != 0 {
		goto L3
	} else {
		goto L459
	}
L459:
	;
	F_errfinish(m, int32(_a_F_rebuild_relation_9), int32(3377), int32(_a_F_rebuild_relation_31))
	mBase = m.M
	v2678 = m.ExcPending
	if v2678 != 0 {
		goto L3
	} else {
		goto L460
	}
L460:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_relation_excluded_by_constraints(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v51 int32
	_ = v51
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	v4 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	if v13 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if int32(0) < v18 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v384
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v25 = v4
	goto L8
L6:
	;
	goto L7
L7:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_relation_excluded_by_constraints[0]))
	switch v67 {
	case 0:
		v384 = int32(0)
		goto L4
	case 1:
		goto L17
	case 2:
		goto L18
	default:
		v76 = v4
		goto L16
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v21+v25<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v51 = v25 + int32(1)
	if v51 != v18 {
		v25 = v51
		goto L8
	} else {
		goto L15
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v41 != int32(7) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v44 = int32(1)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+32)))
	if v45 != 0 {
		v384 = v44
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
	if v46 == int64(0) {
		v384 = v44
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	goto L9
L16:
	;
	v78 = int32(0)
	if v78 < v18 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v76 = base.B2i32(v73 == int32(0))
	goto L16
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v68 == int32(2) {
		v76 = v4
		goto L16
	} else {
		goto L19
	}
L19:
	;
	return int32(0)
L20:
	;
	v85 = int32(0)
	v86 = v78
	goto L23
L21:
	;
	v118 = v78
	goto L22
L22:
	;
	v127 = F_predicate_refuted_by(m, v118, v118, int32(1))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L25
	} else {
		goto L32
	}
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v85<<(uint(int32(2))%32))))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v100 = F_contain_mutable_functions(m, v99)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v118 = v109
	goto L22
L25:
	;
	return int32(0)
L26:
	;
	if v100 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v107 = F_lappend(m, v86, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L25
	} else {
		goto L30
	}
L28:
	;
	v109 = v86
	goto L29
L29:
	;
	v111 = v85 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v111 < v112 {
		v85 = v111
		v86 = v109
		goto L23
	} else {
		goto L31
	}
L30:
	;
	v109 = v107
	goto L29
L31:
	;
	goto L24
L32:
	;
	if v127 != 0 {
		v384 = int32(1)
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v130 != 0 {
		v384 = int32(0)
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v131 = int32(1)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v132 == v131 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
	v138 = base.B2i32(v135 == int32(112))
	goto L37
L36:
	;
	v138 = v131
	goto L37
L37:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v140 = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v143 = F_table_open(m, v141, v140)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L25
	} else {
		goto L39
	}
L38:
	;
	if v76 == int32(0) {
		v314 = v279
		goto L71
	} else {
		goto L72
	}
L39:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v143)+52))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+24))
	if v146 == int32(0) {
		v279 = v140
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+14)))
	if v149 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v154 = int32(0)
	v155 = v140
	goto L44
L42:
	;
	v198 = v140
	goto L43
L43:
	;
	if v138 == int32(0) {
		v279 = v198
		goto L38
	} else {
		goto L59
	}
L44:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	v166 = v163 + v154*int32(12)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+9)))
	if v167 != int32(1) {
		v189 = v155
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v198 = v189
	goto L43
L46:
	;
	v192 = v154 + int32(1)
	if v192 != v149 {
		v154 = v192
		v155 = v189
		goto L44
	} else {
		goto L58
	}
L47:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+10)))
	if v170&v132 != 0 {
		v189 = v155
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v173 = F_stringToNode(m, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L25
	} else {
		goto L49
	}
L49:
	;
	if v139 != int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	F_ChangeVarNodes(m, v173, int32(1), v139)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L25
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v180 = F_eval_const_expressions(m, l0, v173)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L25
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v183 = F_canonicalize_qual(m, v180, int32(1))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L25
	} else {
		goto L55
	}
L55:
	;
	v185 = F_make_ands_implicit(m, v183)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L25
	} else {
		goto L56
	}
L56:
	;
	v187 = F_list_concat(m, v155, v185)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L25
	} else {
		goto L57
	}
L57:
	;
	v189 = v187
	goto L46
L58:
	;
	goto L45
L59:
	;
	v208 = int32(1)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+16)))
	if v209 != v208 {
		v279 = v198
		goto L38
	} else {
		goto L60
	}
L60:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v143)+52))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v213 <= int32(0) {
		v279 = v198
		goto L38
	} else {
		goto L61
	}
L61:
	;
	v219 = v208
	v220 = v198
	goto L62
L62:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v143)+52))
	v230 = v219 - int32(1)
	v233 = v228 + v230<<(uint(int32(3))%32)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+35)))
	if v234 != int32(118) {
		v269 = v220
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v279 = v269
	goto L38
L64:
	;
	v273 = v219 + int32(1)
	if v273 <= v213 {
		v219 = v273
		v220 = v269
		goto L62
	} else {
		goto L70
	}
L65:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+34)))
	if v237&int32(4) != 0 {
		v269 = v220
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v242 = F_palloc0(m, int32(20))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L25
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = int32(52)
	v252 = v228 + v240<<(uint(int32(3))%32) + v230*int32(100)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+96))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v252)+104))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v252)+124))
	v257 = F_makeVar(m, v139, base.I32_extend16_s(v219), v253, v254, v255, int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L25
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242)+16)) = int32(-1)
	v261 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v242)+12)) = uint8(v261)
	*(*int32)(unsafe.Add(mBase, uint32(v242)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v242)+4)) = v257
	v266 = F_lappend(m, v220, v242)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L25
	} else {
		goto L69
	}
L69:
	;
	v269 = v266
	goto L64
L70:
	;
	goto L63
L71:
	;
	if v314 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L72:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v143)+48))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+131)))
	if v290 != int32(1) {
		v314 = v279
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l1)+272))
	if v293 != 0 {
		v309 = v293
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v310 = F_list_concat(m, v279, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L25
	} else {
		goto L85
	}
L75:
	;
	v294 = F_RelationGetPartitionQual(m, v143)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L25
	} else {
		goto L76
	}
L76:
	;
	if v294 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l1)+272))
	v309 = v298
	goto L74
L78:
	;
	goto L79
L79:
	;
	v299 = F_expression_planner(m, v294)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L25
	} else {
		goto L80
	}
L80:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v301 != int32(1) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	F_ChangeVarNodes(m, v299, int32(1), v301)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L25
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+272)) = v299
	v309 = v299
	goto L74
L84:
	;
	goto L83
L85:
	;
	v314 = v310
	goto L71
L86:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v375 = F_predicate_refuted_by(m, v361, v373, int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L25
	} else {
		goto L103
	}
L87:
	;
	v317 = int32(0)
	F_relation_close(m, v143, v317)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L25
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v321 = int32(0)
	v322 = F_expand_generated_columns_in_expr(m, v314, v143, v139)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L25
	} else {
		goto L91
	}
L90:
	;
	v361 = v317
	goto L86
L91:
	;
	F_relation_close(m, v143, int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L25
	} else {
		goto L92
	}
L92:
	;
	if v322 == int32(0) {
		v361 = v321
		goto L86
	} else {
		goto L93
	}
L93:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	if v329 <= int32(0) {
		v361 = v321
		goto L86
	} else {
		goto L94
	}
L94:
	;
	v333 = v321
	v336 = int32(0)
	goto L95
L95:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v322)+12))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v345+v336<<(uint(int32(2))%32))))
	v350 = F_contain_mutable_functions(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L25
	} else {
		goto L97
	}
L96:
	;
	v361 = v356
	goto L86
L97:
	;
	if v350 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v354 = F_lappend(m, v333, v349)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L25
	} else {
		goto L101
	}
L99:
	;
	v356 = v333
	goto L100
L100:
	;
	v358 = v336 + int32(1)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	if v358 < v359 {
		v333 = v356
		v336 = v358
		goto L95
	} else {
		goto L102
	}
L101:
	;
	v356 = v354
	goto L100
L102:
	;
	goto L96
L103:
	;
	v384 = v375
	goto L4
}
func F_relation_is_updatable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	F_check_stack_depth(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = F_try_relation_open(m, l0, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v17 + int32(16)
	return v542
L4:
	;
	if v24 == int32(0) {
		v542 = v5
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v29 = int32(0)
	if l1 == v29 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v67 != 0 {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v67 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v35 <= int32(0) {
		v61 = v29
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v67 = v61
	goto L6
L11:
	;
	v38 = int32(0)
	if v38 < v35 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v41 = v35
	goto L14
L13:
	;
	v41 = v38
	goto L14
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v44 = int32(0)
	goto L15
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42+v44<<(uint(int32(2))%32))))
	v53 = base.B2i32(v52 == v28)
	if v52 == v28 {
		v61 = v53
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v61 = v53
	goto L10
L17:
	;
	v55 = v44 + int32(1)
	if v55 != v41 {
		v44 = v55
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+119)))
	switch v72 - int32(112) {
	case 0, 2:
		goto L24
	default:
		goto L23
	}
L22:
	;
	v542 = v5
	goto L3
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	if v79 == int32(0) {
		v194 = v5
		goto L26
	} else {
		goto L27
	}
L24:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v542 = int32(28)
	goto L3
L26:
	;
	if l2 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v82 <= int32(0) {
		v194 = v5
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v82 != int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v182 = int32(28)
	if v174 != v182 {
		v194 = v174
		goto L26
	} else {
		goto L45
	}
L30:
	;
	v96 = v5
	v98 = v5
	v100 = v5
	goto L33
L31:
	;
	v144 = v5
	v146 = v5
	goto L32
L32:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v85+v144<<(uint(int32(2))%32))))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+17)))
	if v158 != int32(1) {
		v174 = v146
		goto L29
	} else {
		goto L43
	}
L33:
	;
	v108 = v85 + v96<<(uint(int32(2))%32)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+17)))
	if v110 != int32(1) {
		v120 = v98
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v82&int32(1) == int32(0) {
		v174 = v132
		goto L29
	} else {
		goto L42
	}
L35:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+17)))
	if v122 != int32(1) {
		v132 = v120
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	if v113 != 0 {
		v120 = v98
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v120 = int32(1)<<(uint(v115)%32)&int32(28) | v98
	goto L35
L38:
	;
	v133 = int32(2)
	v134 = v96 + v133
	v136 = v100 + v133
	if v136 != v82&int32(2147483646) {
		v96 = v134
		v98 = v132
		v100 = v136
		goto L33
	} else {
		goto L41
	}
L39:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	if v125 != 0 {
		v132 = v120
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v132 = int32(1)<<(uint(v127)%32)&int32(28) | v120
	goto L38
L41:
	;
	goto L34
L42:
	;
	v144 = v134
	v146 = v132
	goto L32
L43:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	if v161 != 0 {
		v174 = v146
		goto L29
	} else {
		goto L44
	}
L44:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v174 = int32(1)<<(uint(v163)%32)&int32(28) | v146
	goto L29
L45:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v542 = v182
	goto L3
L47:
	;
	v231 = v72 - int32(102)
	if v231 != 0 {
		goto L55
	} else {
		goto L56
	}
L48:
	;
	v229 = v194
	goto L47
L49:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
	if v204 == int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+15)))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+10)))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+20)))
	v218 = int32(28)
	v220 = v194 | (v207<<(uint(int32(2))%32)|v210<<(uint(int32(3))%32)|v214<<(uint(int32(4))%32))&v218
	if v220 != v218 {
		v229 = v220
		goto L47
	} else {
		goto L51
	}
L51:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v542 = int32(28)
	goto L3
L53:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L169
	}
L54:
	;
	v258 = F_get_view_query(m, v24)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L77
	}
L55:
	;
	if v231 == int32(16) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	v235 = F_GetFdwRoutineForRelation(m, v24, int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L62
	}
L58:
	;
	goto L54
L59:
	;
	v525 = v229
	goto L53
L61:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L76
	}
L62:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v235)+84))
	if v237 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v238 = m.T0[v237].(func(*base.Module, int32) int32)(m, v24)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v235)+52))
	if v243 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v254 = v238 | v229
	goto L61
L67:
	;
	v244 = v229 | int32(8)
	goto L69
L68:
	;
	v244 = v229
	goto L69
L69:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v235)+64))
	if v247 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v248 = v244 | int32(4)
	goto L72
L71:
	;
	v248 = v244
	goto L72
L72:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v235)+68))
	if v251 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v252 = v248 | int32(16)
	goto L75
L74:
	;
	v252 = v248
	goto L75
L75:
	;
	v254 = v252
	goto L61
L76:
	;
	v542 = v254
	goto L3
L77:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v258)+120))
	if v266 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v402 != 0 {
		v525 = v229
		goto L53
	} else {
		goto L135
	}
L79:
	;
	v402 = int32(_a_F_relation_is_updatable_0)
	goto L78
L80:
	;
	goto L81
L81:
	;
	v268 = int32(_a_F_relation_is_updatable_1)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v258)+100))
	if v269 != 0 {
		v391 = v268
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v402 = v391
	goto L78
L83:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v258)+108))
	if v270 != 0 {
		v391 = v268
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v258)+112))
	if v271 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v402 = int32(_a_F_relation_is_updatable_2)
	goto L78
L86:
	;
	goto L87
L87:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v258)+144))
	if v273 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v402 = int32(_a_F_relation_is_updatable_3)
	goto L78
L89:
	;
	goto L90
L90:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v258)+48))
	if v275 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v402 = int32(_a_F_relation_is_updatable_4)
	goto L78
L92:
	;
	goto L93
L93:
	;
	v277 = int32(_a_F_relation_is_updatable_5)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v258)+128))
	if v278 != 0 {
		v391 = v277
		goto L82
	} else {
		goto L94
	}
L94:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v258)+132))
	if v279 != 0 {
		v391 = v277
		goto L82
	} else {
		goto L95
	}
L95:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+36)))
	if v280 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v402 = int32(_a_F_relation_is_updatable_6)
	goto L78
L97:
	;
	goto L98
L98:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+37)))
	if v282 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v402 = int32(_a_F_relation_is_updatable_7)
	goto L78
L100:
	;
	goto L101
L101:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+38)))
	if v284 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v402 = int32(_a_F_relation_is_updatable_8)
	goto L78
L103:
	;
	goto L104
L104:
	;
	v286 = int32(_a_F_relation_is_updatable_9)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v258)+60))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	if v288 == int32(0) {
		v391 = v286
		goto L82
	} else {
		goto L105
	}
L105:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	if v291 != int32(1) {
		v391 = v286
		goto L82
	} else {
		goto L106
	}
L106:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v288)+12))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if v296 != int32(63) {
		v391 = v286
		goto L82
	} else {
		goto L107
	}
L107:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v258)+52))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+12))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v300+v301<<(uint(int32(2))%32)-int32(4))))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+12))
	if v308 != 0 {
		v391 = v286
		goto L82
	} else {
		goto L108
	}
L108:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+21)))
	v311 = v309 - int32(102)
	v316 = int32(1)
	v320 = (v311<<(uint(int32(7))%32) | int32(base.Ui32(v311&int32(254))>>(uint(v316)%32))) & int32(255)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v320))|base.B2i32(v316<<(uint(v320)%32)&int32(353) == int32(0)) != 0 {
		v391 = v286
		goto L82
	} else {
		goto L109
	}
L109:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v307)+32))
	if v332 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v333 = int32(_a_F_relation_is_updatable_10)
	goto L112
L111:
	;
	v333 = int32(0)
	goto L112
L112:
	;
	if int32(1)|v332 != 0 {
		v391 = v333
		goto L82
	} else {
		goto L113
	}
L113:
	;
	v337 = int32(_a_F_relation_is_updatable_11)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v258)+76))
	if v338 == int32(0) {
		v391 = v337
		goto L82
	} else {
		goto L114
	}
L114:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	if v341 <= int32(0) {
		v391 = v337
		goto L82
	} else {
		goto L115
	}
L115:
	;
	v344 = int32(0)
	if v344 < v341 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v348 = v341
	goto L118
L117:
	;
	v348 = v344
	goto L118
L118:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	v350 = v344
	goto L119
L119:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v349+v350<<(uint(int32(2))%32))))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+26)))
	if v362 != 0 {
		v383 = int32(_a_F_relation_is_updatable_12)
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v391 = int32(0)
	goto L82
L121:
	;
	if v383 != 0 {
		goto L131
	} else {
		goto L132
	}
L122:
	;
	v363 = int32(_a_F_relation_is_updatable_13)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	if v365 != int32(6) {
		v380 = v363
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v383 = v380
	goto L121
L124:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	if v368 != v369 {
		v380 = v363
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v364)+28))
	if v371 != 0 {
		v380 = v363
		goto L123
	} else {
		goto L126
	}
L126:
	;
	v373 = int32(*(*int16)(unsafe.Add(mBase, uint32(v364)+8)))
	if v373 < int32(0) {
		v383 = int32(_a_F_relation_is_updatable_14)
		goto L121
	} else {
		goto L127
	}
L127:
	;
	if v373 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v378 = int32(0)
	goto L130
L129:
	;
	v378 = int32(_a_F_relation_is_updatable_15)
	goto L130
L130:
	;
	v380 = v378
	goto L123
L131:
	;
	v385 = v350 + int32(1)
	if v348 != v385 {
		v350 = v385
		goto L119
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	goto L120
L134:
	;
	v391 = v337
	goto L82
L135:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v258)+60))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)+12))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	v409 = v17 + int32(12)
	if v409 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409))) = int32(0)
	goto L138
L137:
	;
	goto L138
L138:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v258)+76))
	if v412 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if l3 != 0 {
		goto L156
	} else {
		goto L157
	}
L140:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	if v415 <= int32(0) {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v425 = int32(7)
	v428 = int32(0)
	goto L142
L142:
	;
	v434 = v425 + int32(1)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v412)+12))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v435+v428<<(uint(int32(2))%32))))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439)+26)))
	if v440 != 0 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	goto L139
L144:
	;
	v466 = v428 + int32(1)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	if v466 < v467 {
		v425 = v434
		v428 = v466
		goto L142
	} else {
		goto L155
	}
L145:
	;
	v462 = F_bms_is_member(m, base.I32_extend16_s(v434), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L153
	}
L146:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	if v442 != int32(6) {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	if v445 != v446 {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v441)+28))
	if v448 != 0 {
		goto L145
	} else {
		goto L149
	}
L149:
	;
	v449 = int32(*(*int16)(unsafe.Add(mBase, uint32(v441)+8)))
	if v449 <= int32(0) {
		goto L145
	} else {
		goto L150
	}
L150:
	;
	if v409 == int32(0) {
		goto L144
	} else {
		goto L151
	}
L151:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	v456 = F_bms_add_member(m, v454, base.I32_extend16_s(v434))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409))) = v456
	goto L144
L153:
	;
	if v462 != 0 {
		goto L139
	} else {
		goto L154
	}
L154:
	;
	goto L144
L155:
	;
	goto L143
L156:
	;
	v484 = F_bms_int_members(m, v483, l3)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L159
	}
L157:
	;
	v486 = v483
	goto L158
L158:
	;
	if v486 != 0 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v486 = v484
	goto L158
L160:
	;
	v489 = int32(28)
	goto L162
L161:
	;
	v489 = int32(16)
	goto L162
L162:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v258)+52))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)+12))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v258)+60))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+12))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v491+v496<<(uint(int32(2))%32)-int32(4))))
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502)+21)))
	switch v503 - int32(112) {
	case 0, 2:
		v519 = v489
		goto L163
	default:
		goto L164
	}
L163:
	;
	v525 = v519 | v229
	goto L53
L164:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v502)+16))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v508 = F_lappend_oid(m, l1, v507)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v258)+76))
	v511 = F_adjust_view_column_set(m, v486, v510)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v513 = F_relation_is_updatable(m, v506, v508, l2, v511)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v515 = F_list_delete_last(m, v508)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v519 = v513 & v489
	goto L163
L169:
	;
	v542 = v525
	goto L3
}
func F_relation_statistics_update(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 float32
	_ = v109
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 float32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 float32
	_ = v207
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v2
	F_stats_check_required_arg(m, l0, int32(_a_F_relation_statistics_update_0), v2)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		F_stats_check_required_arg(m, l0, int32(_a_F_relation_statistics_update_0), int32(1))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v35 = F_text_to_cstring(m, v34)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v38 = F_text_to_cstring(m, v37)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_statistics_update[0])))
					if v42 == int32(1) {
						v47 = *(*int32)(unsafe.Add(mBase, _c_F_relation_statistics_update[1]))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+308))
						v50 = base.B2i32(v48 != int32(2))
						*(*uint8)(unsafe.Add(mBase, _c_F_relation_statistics_update[0])) = uint8(v50)
						v52 = v50
					} else {
						v52 = int32(0)
					}
					if v52 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_relation_statistics_update_1), int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(_a_F_relation_statistics_update_2), int32(0))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_relation_statistics_update_3), int32(87), int32(_a_F_relation_statistics_update_4))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					} else {
						v74 = F_makeRangeVar(m, v35, v38, int32(-1))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							v81 = F_RangeVarGetRelidExtended(m, v74, int32(4), int32(0), int32(1138), v20+int32(12))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								v83 = m.G0
								v85 = v83 - int32(80)
								m.G0 = v85
								v87 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v85)+72)) = v87
								*(*int64)(unsafe.Add(mBase, uint32(v85)+64)) = v87
								*(*int64)(unsafe.Add(mBase, uint32(v85)+56)) = v87
								*(*int64)(unsafe.Add(mBase, uint32(v85)+48)) = v87
								*(*int64)(unsafe.Add(mBase, uint32(v85)+40)) = v87
								*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v87
								v99 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v99
								v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
								if v101 == v99 {
									v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									v105 = v104
								} else {
									v105 = v2
								}
								v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
								if v107 != 0 {
									v162 = v2
									v163 = float32(0)
									v164 = int32(1)
									v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
									if v165 == int32(0) {
										v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
										v169 = v168
									} else {
										v169 = v2
									}
									v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+112)))
									if v170 == int32(0) {
										v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
										v174 = v173
									} else {
										v174 = v2
									}
									v177 = F_table_open(m, int32(1259), int32(3))
									mBase = m.M
									v178 = m.ExcPending
									if v178 != 0 {
										return int32(0)
									} else {
										v181 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v81))
										mBase = m.M
										v182 = m.ExcPending
										if v182 != 0 {
											return int32(0)
										} else {
											if v181 != 0 {
												v183 = *(*int32)(unsafe.Add(mBase, uint32(v181)+16))
												v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+22)))
												v185 = v183 + v184
												v186 = int32(0)
												v188 = v85 - int32(-64)
												v190 = v85 + int32(32)
												if v101 != 0 {
													v202 = v186
													v203 = v188
													v204 = v190
												} else {
													v191 = *(*int32)(unsafe.Add(mBase, uint32(v185)+96))
													if v105 == v191 {
														v202 = v186
														v203 = v188
														v204 = v190
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v85)+64)) = int32(10)
														*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = base.I64_extend_i32_u(v105)
														v202 = int32(1)
														v203 = v188 | int32(4)
														v204 = v190 | int32(8)
													}
												}
												if v162 == int32(0) {
													v216 = v202
												} else {
													v207 = *(*float32)(unsafe.Add(mBase, uint32(v185)+100))
													if base.F32_eq(v163, v207) != 0 {
														v216 = v202
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v203))) = int32(11)
														*(*int64)(unsafe.Add(mBase, uint32(v204))) = base.I64_extend_i32_s(base.I32_reinterpret_f32(v163))
														v216 = v202 + int32(1)
													}
												}
												if v165 != 0 {
													v235 = v216
												} else {
													v217 = *(*int32)(unsafe.Add(mBase, uint32(v185)+104))
													if v169 == v217 {
														v235 = v216
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v85+int32(32)+v216<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v169)
														*(*int32)(unsafe.Add(mBase, uint32(v85-int32(-64)|v216<<(uint(int32(2))%32)))) = int32(12)
														v235 = v216 + int32(1)
													}
												}
												if v170 != 0 {
													if v235 == int32(0) {
														F_ReleaseCatCache(m, v181)
														mBase = m.M
														v274 = m.ExcPending
														if v274 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v177, int32(3))
															mBase = m.M
															v277 = m.ExcPending
															if v277 != 0 {
																return int32(0)
															} else {
																F_CommandCounterIncrement(m)
																mBase = m.M
																v279 = m.ExcPending
																if v279 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v85 + int32(80)
																	m.G0 = v20 + int32(16)
																	return v164
																}
															}
														}
													} else {
														v256 = v235
														v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
														v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
														mBase = m.M
														v265 = m.ExcPending
														if v265 != 0 {
															return int32(0)
														} else {
															F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
															mBase = m.M
															v269 = m.ExcPending
															if v269 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v264)
																mBase = m.M
																v271 = m.ExcPending
																if v271 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v181)
																	mBase = m.M
																	v274 = m.ExcPending
																	if v274 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v177, int32(3))
																		mBase = m.M
																		v277 = m.ExcPending
																		if v277 != 0 {
																			return int32(0)
																		} else {
																			F_CommandCounterIncrement(m)
																			mBase = m.M
																			v279 = m.ExcPending
																			if v279 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v85 + int32(80)
																				m.G0 = v20 + int32(16)
																				return v164
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v236 = *(*int32)(unsafe.Add(mBase, uint32(v185)+108))
													if v174 == v236 {
														if v235 == int32(0) {
															F_ReleaseCatCache(m, v181)
															mBase = m.M
															v274 = m.ExcPending
															if v274 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v177, int32(3))
																mBase = m.M
																v277 = m.ExcPending
																if v277 != 0 {
																	return int32(0)
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v279 = m.ExcPending
																	if v279 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v85 + int32(80)
																		m.G0 = v20 + int32(16)
																		return v164
																	}
																}
															}
														} else {
															v256 = v235
															v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
															v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
															mBase = m.M
															v265 = m.ExcPending
															if v265 != 0 {
																return int32(0)
															} else {
																F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																mBase = m.M
																v269 = m.ExcPending
																if v269 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v264)
																	mBase = m.M
																	v271 = m.ExcPending
																	if v271 != 0 {
																		return int32(0)
																	} else {
																		F_ReleaseCatCache(m, v181)
																		mBase = m.M
																		v274 = m.ExcPending
																		if v274 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v177, int32(3))
																			mBase = m.M
																			v277 = m.ExcPending
																			if v277 != 0 {
																				return int32(0)
																			} else {
																				F_CommandCounterIncrement(m)
																				mBase = m.M
																				v279 = m.ExcPending
																				if v279 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v85 + int32(80)
																					m.G0 = v20 + int32(16)
																					return v164
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v85+int32(32)+v235<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v174)
														*(*int32)(unsafe.Add(mBase, uint32(v85-int32(-64)+v235<<(uint(int32(2))%32)))) = int32(13)
														v256 = v235 + int32(1)
														v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
														v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
														mBase = m.M
														v265 = m.ExcPending
														if v265 != 0 {
															return int32(0)
														} else {
															F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
															mBase = m.M
															v269 = m.ExcPending
															if v269 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v264)
																mBase = m.M
																v271 = m.ExcPending
																if v271 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v181)
																	mBase = m.M
																	v274 = m.ExcPending
																	if v274 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v177, int32(3))
																		mBase = m.M
																		v277 = m.ExcPending
																		if v277 != 0 {
																			return int32(0)
																		} else {
																			F_CommandCounterIncrement(m)
																			mBase = m.M
																			v279 = m.ExcPending
																			if v279 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v85 + int32(80)
																				m.G0 = v20 + int32(16)
																				return v164
																			}
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v286 = m.ExcPending
												if v286 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v85))) = v81
													F_errmsg_internal(m, int32(_a_F_relation_statistics_update_5), v85)
													mBase = m.M
													v290 = m.ExcPending
													if v290 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_relation_statistics_update_3), int32(166), int32(_a_F_relation_statistics_update_6))
														mBase = m.M
														v295 = m.ExcPending
														if v295 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										}
									}
								} else {
									v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
									v109 = base.F32_reinterpret_i32(v108)
									if base.B2i32(base.Ui32(v108&int32(2147483647)) <= base.Ui32(int32(2139095040)))&base.F32_ne(base.F32_abs(v109), math.Float32frombits(uint32(0x7f800000))) == int32(0) {
										v122 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return int32(0)
										} else {
											if v122 == int32(0) {
												v162 = v2
												v163 = v109
												v164 = int32(0)
												v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
												if v165 == int32(0) {
													v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
													v169 = v168
												} else {
													v169 = v2
												}
												v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+112)))
												if v170 == int32(0) {
													v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
													v174 = v173
												} else {
													v174 = v2
												}
												v177 = F_table_open(m, int32(1259), int32(3))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return int32(0)
												} else {
													v181 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v81))
													mBase = m.M
													v182 = m.ExcPending
													if v182 != 0 {
														return int32(0)
													} else {
														if v181 != 0 {
															v183 = *(*int32)(unsafe.Add(mBase, uint32(v181)+16))
															v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+22)))
															v185 = v183 + v184
															v186 = int32(0)
															v188 = v85 - int32(-64)
															v190 = v85 + int32(32)
															if v101 != 0 {
																v202 = v186
																v203 = v188
																v204 = v190
															} else {
																v191 = *(*int32)(unsafe.Add(mBase, uint32(v185)+96))
																if v105 == v191 {
																	v202 = v186
																	v203 = v188
																	v204 = v190
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v85)+64)) = int32(10)
																	*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = base.I64_extend_i32_u(v105)
																	v202 = int32(1)
																	v203 = v188 | int32(4)
																	v204 = v190 | int32(8)
																}
															}
															if v162 == int32(0) {
																v216 = v202
															} else {
																v207 = *(*float32)(unsafe.Add(mBase, uint32(v185)+100))
																if base.F32_eq(v163, v207) != 0 {
																	v216 = v202
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v203))) = int32(11)
																	*(*int64)(unsafe.Add(mBase, uint32(v204))) = base.I64_extend_i32_s(base.I32_reinterpret_f32(v163))
																	v216 = v202 + int32(1)
																}
															}
															if v165 != 0 {
																v235 = v216
															} else {
																v217 = *(*int32)(unsafe.Add(mBase, uint32(v185)+104))
																if v169 == v217 {
																	v235 = v216
																} else {
																	*(*int64)(unsafe.Add(mBase, uint32(v85+int32(32)+v216<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v169)
																	*(*int32)(unsafe.Add(mBase, uint32(v85-int32(-64)|v216<<(uint(int32(2))%32)))) = int32(12)
																	v235 = v216 + int32(1)
																}
															}
															if v170 != 0 {
																if v235 == int32(0) {
																	F_ReleaseCatCache(m, v181)
																	mBase = m.M
																	v274 = m.ExcPending
																	if v274 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v177, int32(3))
																		mBase = m.M
																		v277 = m.ExcPending
																		if v277 != 0 {
																			return int32(0)
																		} else {
																			F_CommandCounterIncrement(m)
																			mBase = m.M
																			v279 = m.ExcPending
																			if v279 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v85 + int32(80)
																				m.G0 = v20 + int32(16)
																				return v164
																			}
																		}
																	}
																} else {
																	v256 = v235
																	v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
																	v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
																	mBase = m.M
																	v265 = m.ExcPending
																	if v265 != 0 {
																		return int32(0)
																	} else {
																		F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																		mBase = m.M
																		v269 = m.ExcPending
																		if v269 != 0 {
																			return int32(0)
																		} else {
																			F_pfree(m, v264)
																			mBase = m.M
																			v271 = m.ExcPending
																			if v271 != 0 {
																				return int32(0)
																			} else {
																				F_ReleaseCatCache(m, v181)
																				mBase = m.M
																				v274 = m.ExcPending
																				if v274 != 0 {
																					return int32(0)
																				} else {
																					F_relation_close(m, v177, int32(3))
																					mBase = m.M
																					v277 = m.ExcPending
																					if v277 != 0 {
																						return int32(0)
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v279 = m.ExcPending
																						if v279 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v85 + int32(80)
																							m.G0 = v20 + int32(16)
																							return v164
																						}
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																v236 = *(*int32)(unsafe.Add(mBase, uint32(v185)+108))
																if v174 == v236 {
																	if v235 == int32(0) {
																		F_ReleaseCatCache(m, v181)
																		mBase = m.M
																		v274 = m.ExcPending
																		if v274 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v177, int32(3))
																			mBase = m.M
																			v277 = m.ExcPending
																			if v277 != 0 {
																				return int32(0)
																			} else {
																				F_CommandCounterIncrement(m)
																				mBase = m.M
																				v279 = m.ExcPending
																				if v279 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v85 + int32(80)
																					m.G0 = v20 + int32(16)
																					return v164
																				}
																			}
																		}
																	} else {
																		v256 = v235
																		v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
																		v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
																		mBase = m.M
																		v265 = m.ExcPending
																		if v265 != 0 {
																			return int32(0)
																		} else {
																			F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																			mBase = m.M
																			v269 = m.ExcPending
																			if v269 != 0 {
																				return int32(0)
																			} else {
																				F_pfree(m, v264)
																				mBase = m.M
																				v271 = m.ExcPending
																				if v271 != 0 {
																					return int32(0)
																				} else {
																					F_ReleaseCatCache(m, v181)
																					mBase = m.M
																					v274 = m.ExcPending
																					if v274 != 0 {
																						return int32(0)
																					} else {
																						F_relation_close(m, v177, int32(3))
																						mBase = m.M
																						v277 = m.ExcPending
																						if v277 != 0 {
																							return int32(0)
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v279 = m.ExcPending
																							if v279 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v85 + int32(80)
																								m.G0 = v20 + int32(16)
																								return v164
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	*(*int64)(unsafe.Add(mBase, uint32(v85+int32(32)+v235<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v174)
																	*(*int32)(unsafe.Add(mBase, uint32(v85-int32(-64)+v235<<(uint(int32(2))%32)))) = int32(13)
																	v256 = v235 + int32(1)
																	v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
																	v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
																	mBase = m.M
																	v265 = m.ExcPending
																	if v265 != 0 {
																		return int32(0)
																	} else {
																		F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																		mBase = m.M
																		v269 = m.ExcPending
																		if v269 != 0 {
																			return int32(0)
																		} else {
																			F_pfree(m, v264)
																			mBase = m.M
																			v271 = m.ExcPending
																			if v271 != 0 {
																				return int32(0)
																			} else {
																				F_ReleaseCatCache(m, v181)
																				mBase = m.M
																				v274 = m.ExcPending
																				if v274 != 0 {
																					return int32(0)
																				} else {
																					F_relation_close(m, v177, int32(3))
																					mBase = m.M
																					v277 = m.ExcPending
																					if v277 != 0 {
																						return int32(0)
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v279 = m.ExcPending
																						if v279 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v85 + int32(80)
																							m.G0 = v20 + int32(16)
																							return v164
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v286 = m.ExcPending
															if v286 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v85))) = v81
																F_errmsg_internal(m, int32(_a_F_relation_statistics_update_5), v85)
																mBase = m.M
																v290 = m.ExcPending
																if v290 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_relation_statistics_update_3), int32(166), int32(_a_F_relation_statistics_update_6))
																	mBase = m.M
																	v295 = m.ExcPending
																	if v295 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v142 = int32(_a_F_relation_statistics_update_7)
												v143 = int32(132)
												F_errcode(m, int32(50856066))
												mBase = m.M
												v146 = m.ExcPending
												if v146 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v85)+16)) = int32(_a_F_relation_statistics_update_8)
													F_errmsg(m, v142, v85+int32(16))
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_relation_statistics_update_3), v143, int32(_a_F_relation_statistics_update_6))
														mBase = m.M
														v156 = m.ExcPending
														if v156 != 0 {
															return int32(0)
														} else {
															v162 = v2
															v163 = v109
															v164 = int32(0)
															v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
															if v165 == int32(0) {
																v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																v169 = v168
															} else {
																v169 = v2
															}
															v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+112)))
															if v170 == int32(0) {
																v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
																v174 = v173
															} else {
																v174 = v2
															}
															v177 = F_table_open(m, int32(1259), int32(3))
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return int32(0)
															} else {
																v181 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v81))
																mBase = m.M
																v182 = m.ExcPending
																if v182 != 0 {
																	return int32(0)
																} else {
																	if v181 != 0 {
																		v183 = *(*int32)(unsafe.Add(mBase, uint32(v181)+16))
																		v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+22)))
																		v185 = v183 + v184
																		v186 = int32(0)
																		v188 = v85 - int32(-64)
																		v190 = v85 + int32(32)
																		if v101 != 0 {
																			v202 = v186
																			v203 = v188
																			v204 = v190
																		} else {
																			v191 = *(*int32)(unsafe.Add(mBase, uint32(v185)+96))
																			if v105 == v191 {
																				v202 = v186
																				v203 = v188
																				v204 = v190
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v85)+64)) = int32(10)
																				*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = base.I64_extend_i32_u(v105)
																				v202 = int32(1)
																				v203 = v188 | int32(4)
																				v204 = v190 | int32(8)
																			}
																		}
																		if v162 == int32(0) {
																			v216 = v202
																		} else {
																			v207 = *(*float32)(unsafe.Add(mBase, uint32(v185)+100))
																			if base.F32_eq(v163, v207) != 0 {
																				v216 = v202
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v203))) = int32(11)
																				*(*int64)(unsafe.Add(mBase, uint32(v204))) = base.I64_extend_i32_s(base.I32_reinterpret_f32(v163))
																				v216 = v202 + int32(1)
																			}
																		}
																		if v165 != 0 {
																			v235 = v216
																		} else {
																			v217 = *(*int32)(unsafe.Add(mBase, uint32(v185)+104))
																			if v169 == v217 {
																				v235 = v216
																			} else {
																				*(*int64)(unsafe.Add(mBase, uint32(v85+int32(32)+v216<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v169)
																				*(*int32)(unsafe.Add(mBase, uint32(v85-int32(-64)|v216<<(uint(int32(2))%32)))) = int32(12)
																				v235 = v216 + int32(1)
																			}
																		}
																		if v170 != 0 {
																			if v235 == int32(0) {
																				F_ReleaseCatCache(m, v181)
																				mBase = m.M
																				v274 = m.ExcPending
																				if v274 != 0 {
																					return int32(0)
																				} else {
																					F_relation_close(m, v177, int32(3))
																					mBase = m.M
																					v277 = m.ExcPending
																					if v277 != 0 {
																						return int32(0)
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v279 = m.ExcPending
																						if v279 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v85 + int32(80)
																							m.G0 = v20 + int32(16)
																							return v164
																						}
																					}
																				}
																			} else {
																				v256 = v235
																				v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
																				v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
																				mBase = m.M
																				v265 = m.ExcPending
																				if v265 != 0 {
																					return int32(0)
																				} else {
																					F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																					mBase = m.M
																					v269 = m.ExcPending
																					if v269 != 0 {
																						return int32(0)
																					} else {
																						F_pfree(m, v264)
																						mBase = m.M
																						v271 = m.ExcPending
																						if v271 != 0 {
																							return int32(0)
																						} else {
																							F_ReleaseCatCache(m, v181)
																							mBase = m.M
																							v274 = m.ExcPending
																							if v274 != 0 {
																								return int32(0)
																							} else {
																								F_relation_close(m, v177, int32(3))
																								mBase = m.M
																								v277 = m.ExcPending
																								if v277 != 0 {
																									return int32(0)
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v279 = m.ExcPending
																									if v279 != 0 {
																										return int32(0)
																									} else {
																										m.G0 = v85 + int32(80)
																										m.G0 = v20 + int32(16)
																										return v164
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v236 = *(*int32)(unsafe.Add(mBase, uint32(v185)+108))
																			if v174 == v236 {
																				if v235 == int32(0) {
																					F_ReleaseCatCache(m, v181)
																					mBase = m.M
																					v274 = m.ExcPending
																					if v274 != 0 {
																						return int32(0)
																					} else {
																						F_relation_close(m, v177, int32(3))
																						mBase = m.M
																						v277 = m.ExcPending
																						if v277 != 0 {
																							return int32(0)
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v279 = m.ExcPending
																							if v279 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v85 + int32(80)
																								m.G0 = v20 + int32(16)
																								return v164
																							}
																						}
																					}
																				} else {
																					v256 = v235
																					v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
																					v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
																					mBase = m.M
																					v265 = m.ExcPending
																					if v265 != 0 {
																						return int32(0)
																					} else {
																						F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																						mBase = m.M
																						v269 = m.ExcPending
																						if v269 != 0 {
																							return int32(0)
																						} else {
																							F_pfree(m, v264)
																							mBase = m.M
																							v271 = m.ExcPending
																							if v271 != 0 {
																								return int32(0)
																							} else {
																								F_ReleaseCatCache(m, v181)
																								mBase = m.M
																								v274 = m.ExcPending
																								if v274 != 0 {
																									return int32(0)
																								} else {
																									F_relation_close(m, v177, int32(3))
																									mBase = m.M
																									v277 = m.ExcPending
																									if v277 != 0 {
																										return int32(0)
																									} else {
																										F_CommandCounterIncrement(m)
																										mBase = m.M
																										v279 = m.ExcPending
																										if v279 != 0 {
																											return int32(0)
																										} else {
																											m.G0 = v85 + int32(80)
																											m.G0 = v20 + int32(16)
																											return v164
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			} else {
																				*(*int64)(unsafe.Add(mBase, uint32(v85+int32(32)+v235<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v174)
																				*(*int32)(unsafe.Add(mBase, uint32(v85-int32(-64)+v235<<(uint(int32(2))%32)))) = int32(13)
																				v256 = v235 + int32(1)
																				v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
																				v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
																				mBase = m.M
																				v265 = m.ExcPending
																				if v265 != 0 {
																					return int32(0)
																				} else {
																					F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																					mBase = m.M
																					v269 = m.ExcPending
																					if v269 != 0 {
																						return int32(0)
																					} else {
																						F_pfree(m, v264)
																						mBase = m.M
																						v271 = m.ExcPending
																						if v271 != 0 {
																							return int32(0)
																						} else {
																							F_ReleaseCatCache(m, v181)
																							mBase = m.M
																							v274 = m.ExcPending
																							if v274 != 0 {
																								return int32(0)
																							} else {
																								F_relation_close(m, v177, int32(3))
																								mBase = m.M
																								v277 = m.ExcPending
																								if v277 != 0 {
																									return int32(0)
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v279 = m.ExcPending
																									if v279 != 0 {
																										return int32(0)
																									} else {
																										m.G0 = v85 + int32(80)
																										m.G0 = v20 + int32(16)
																										return v164
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	} else {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v286 = m.ExcPending
																		if v286 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v85))) = v81
																			F_errmsg_internal(m, int32(_a_F_relation_statistics_update_5), v85)
																			mBase = m.M
																			v290 = m.ExcPending
																			if v290 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_relation_statistics_update_3), int32(166), int32(_a_F_relation_statistics_update_6))
																				mBase = m.M
																				v295 = m.ExcPending
																				if v295 != 0 {
																					return int32(0)
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										if base.F32_lt(v109, float32(-1)) == int32(0) {
											v132 = int32(1)
											v162 = v132
											v163 = v109
											v164 = v132
											v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
											if v165 == int32(0) {
												v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
												v169 = v168
											} else {
												v169 = v2
											}
											v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+112)))
											if v170 == int32(0) {
												v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
												v174 = v173
											} else {
												v174 = v2
											}
											v177 = F_table_open(m, int32(1259), int32(3))
											mBase = m.M
											v178 = m.ExcPending
											if v178 != 0 {
												return int32(0)
											} else {
												v181 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v81))
												mBase = m.M
												v182 = m.ExcPending
												if v182 != 0 {
													return int32(0)
												} else {
													if v181 != 0 {
														v183 = *(*int32)(unsafe.Add(mBase, uint32(v181)+16))
														v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+22)))
														v185 = v183 + v184
														v186 = int32(0)
														v188 = v85 - int32(-64)
														v190 = v85 + int32(32)
														if v101 != 0 {
															v202 = v186
															v203 = v188
															v204 = v190
														} else {
															v191 = *(*int32)(unsafe.Add(mBase, uint32(v185)+96))
															if v105 == v191 {
																v202 = v186
																v203 = v188
																v204 = v190
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v85)+64)) = int32(10)
																*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = base.I64_extend_i32_u(v105)
																v202 = int32(1)
																v203 = v188 | int32(4)
																v204 = v190 | int32(8)
															}
														}
														if v162 == int32(0) {
															v216 = v202
														} else {
															v207 = *(*float32)(unsafe.Add(mBase, uint32(v185)+100))
															if base.F32_eq(v163, v207) != 0 {
																v216 = v202
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v203))) = int32(11)
																*(*int64)(unsafe.Add(mBase, uint32(v204))) = base.I64_extend_i32_s(base.I32_reinterpret_f32(v163))
																v216 = v202 + int32(1)
															}
														}
														if v165 != 0 {
															v235 = v216
														} else {
															v217 = *(*int32)(unsafe.Add(mBase, uint32(v185)+104))
															if v169 == v217 {
																v235 = v216
															} else {
																*(*int64)(unsafe.Add(mBase, uint32(v85+int32(32)+v216<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v169)
																*(*int32)(unsafe.Add(mBase, uint32(v85-int32(-64)|v216<<(uint(int32(2))%32)))) = int32(12)
																v235 = v216 + int32(1)
															}
														}
														if v170 != 0 {
															if v235 == int32(0) {
																F_ReleaseCatCache(m, v181)
																mBase = m.M
																v274 = m.ExcPending
																if v274 != 0 {
																	return int32(0)
																} else {
																	F_relation_close(m, v177, int32(3))
																	mBase = m.M
																	v277 = m.ExcPending
																	if v277 != 0 {
																		return int32(0)
																	} else {
																		F_CommandCounterIncrement(m)
																		mBase = m.M
																		v279 = m.ExcPending
																		if v279 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v85 + int32(80)
																			m.G0 = v20 + int32(16)
																			return v164
																		}
																	}
																}
															} else {
																v256 = v235
																v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
																v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
																mBase = m.M
																v265 = m.ExcPending
																if v265 != 0 {
																	return int32(0)
																} else {
																	F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																	mBase = m.M
																	v269 = m.ExcPending
																	if v269 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v264)
																		mBase = m.M
																		v271 = m.ExcPending
																		if v271 != 0 {
																			return int32(0)
																		} else {
																			F_ReleaseCatCache(m, v181)
																			mBase = m.M
																			v274 = m.ExcPending
																			if v274 != 0 {
																				return int32(0)
																			} else {
																				F_relation_close(m, v177, int32(3))
																				mBase = m.M
																				v277 = m.ExcPending
																				if v277 != 0 {
																					return int32(0)
																				} else {
																					F_CommandCounterIncrement(m)
																					mBase = m.M
																					v279 = m.ExcPending
																					if v279 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v85 + int32(80)
																						m.G0 = v20 + int32(16)
																						return v164
																					}
																				}
																			}
																		}
																	}
																}
															}
														} else {
															v236 = *(*int32)(unsafe.Add(mBase, uint32(v185)+108))
															if v174 == v236 {
																if v235 == int32(0) {
																	F_ReleaseCatCache(m, v181)
																	mBase = m.M
																	v274 = m.ExcPending
																	if v274 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v177, int32(3))
																		mBase = m.M
																		v277 = m.ExcPending
																		if v277 != 0 {
																			return int32(0)
																		} else {
																			F_CommandCounterIncrement(m)
																			mBase = m.M
																			v279 = m.ExcPending
																			if v279 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v85 + int32(80)
																				m.G0 = v20 + int32(16)
																				return v164
																			}
																		}
																	}
																} else {
																	v256 = v235
																	v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
																	v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
																	mBase = m.M
																	v265 = m.ExcPending
																	if v265 != 0 {
																		return int32(0)
																	} else {
																		F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																		mBase = m.M
																		v269 = m.ExcPending
																		if v269 != 0 {
																			return int32(0)
																		} else {
																			F_pfree(m, v264)
																			mBase = m.M
																			v271 = m.ExcPending
																			if v271 != 0 {
																				return int32(0)
																			} else {
																				F_ReleaseCatCache(m, v181)
																				mBase = m.M
																				v274 = m.ExcPending
																				if v274 != 0 {
																					return int32(0)
																				} else {
																					F_relation_close(m, v177, int32(3))
																					mBase = m.M
																					v277 = m.ExcPending
																					if v277 != 0 {
																						return int32(0)
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v279 = m.ExcPending
																						if v279 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v85 + int32(80)
																							m.G0 = v20 + int32(16)
																							return v164
																						}
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																*(*int64)(unsafe.Add(mBase, uint32(v85+int32(32)+v235<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v174)
																*(*int32)(unsafe.Add(mBase, uint32(v85-int32(-64)+v235<<(uint(int32(2))%32)))) = int32(13)
																v256 = v235 + int32(1)
																v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
																v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
																mBase = m.M
																v265 = m.ExcPending
																if v265 != 0 {
																	return int32(0)
																} else {
																	F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																	mBase = m.M
																	v269 = m.ExcPending
																	if v269 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v264)
																		mBase = m.M
																		v271 = m.ExcPending
																		if v271 != 0 {
																			return int32(0)
																		} else {
																			F_ReleaseCatCache(m, v181)
																			mBase = m.M
																			v274 = m.ExcPending
																			if v274 != 0 {
																				return int32(0)
																			} else {
																				F_relation_close(m, v177, int32(3))
																				mBase = m.M
																				v277 = m.ExcPending
																				if v277 != 0 {
																					return int32(0)
																				} else {
																					F_CommandCounterIncrement(m)
																					mBase = m.M
																					v279 = m.ExcPending
																					if v279 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v85 + int32(80)
																						m.G0 = v20 + int32(16)
																						return v164
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v286 = m.ExcPending
														if v286 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v85))) = v81
															F_errmsg_internal(m, int32(_a_F_relation_statistics_update_5), v85)
															mBase = m.M
															v290 = m.ExcPending
															if v290 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_relation_statistics_update_3), int32(166), int32(_a_F_relation_statistics_update_6))
																mBase = m.M
																v295 = m.ExcPending
																if v295 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v136 = F_errstart(m, int32(19), int32(0))
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return int32(0)
											} else {
												if v136 == int32(0) {
													v162 = v2
													v163 = v109
													v164 = int32(0)
													v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
													if v165 == int32(0) {
														v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
														v169 = v168
													} else {
														v169 = v2
													}
													v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+112)))
													if v170 == int32(0) {
														v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
														v174 = v173
													} else {
														v174 = v2
													}
													v177 = F_table_open(m, int32(1259), int32(3))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return int32(0)
													} else {
														v181 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v81))
														mBase = m.M
														v182 = m.ExcPending
														if v182 != 0 {
															return int32(0)
														} else {
															if v181 != 0 {
																v183 = *(*int32)(unsafe.Add(mBase, uint32(v181)+16))
																v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+22)))
																v185 = v183 + v184
																v186 = int32(0)
																v188 = v85 - int32(-64)
																v190 = v85 + int32(32)
																if v101 != 0 {
																	v202 = v186
																	v203 = v188
																	v204 = v190
																} else {
																	v191 = *(*int32)(unsafe.Add(mBase, uint32(v185)+96))
																	if v105 == v191 {
																		v202 = v186
																		v203 = v188
																		v204 = v190
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v85)+64)) = int32(10)
																		*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = base.I64_extend_i32_u(v105)
																		v202 = int32(1)
																		v203 = v188 | int32(4)
																		v204 = v190 | int32(8)
																	}
																}
																if v162 == int32(0) {
																	v216 = v202
																} else {
																	v207 = *(*float32)(unsafe.Add(mBase, uint32(v185)+100))
																	if base.F32_eq(v163, v207) != 0 {
																		v216 = v202
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v203))) = int32(11)
																		*(*int64)(unsafe.Add(mBase, uint32(v204))) = base.I64_extend_i32_s(base.I32_reinterpret_f32(v163))
																		v216 = v202 + int32(1)
																	}
																}
																if v165 != 0 {
																	v235 = v216
																} else {
																	v217 = *(*int32)(unsafe.Add(mBase, uint32(v185)+104))
																	if v169 == v217 {
																		v235 = v216
																	} else {
																		*(*int64)(unsafe.Add(mBase, uint32(v85+int32(32)+v216<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v169)
																		*(*int32)(unsafe.Add(mBase, uint32(v85-int32(-64)|v216<<(uint(int32(2))%32)))) = int32(12)
																		v235 = v216 + int32(1)
																	}
																}
																if v170 != 0 {
																	if v235 == int32(0) {
																		F_ReleaseCatCache(m, v181)
																		mBase = m.M
																		v274 = m.ExcPending
																		if v274 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v177, int32(3))
																			mBase = m.M
																			v277 = m.ExcPending
																			if v277 != 0 {
																				return int32(0)
																			} else {
																				F_CommandCounterIncrement(m)
																				mBase = m.M
																				v279 = m.ExcPending
																				if v279 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v85 + int32(80)
																					m.G0 = v20 + int32(16)
																					return v164
																				}
																			}
																		}
																	} else {
																		v256 = v235
																		v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
																		v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
																		mBase = m.M
																		v265 = m.ExcPending
																		if v265 != 0 {
																			return int32(0)
																		} else {
																			F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																			mBase = m.M
																			v269 = m.ExcPending
																			if v269 != 0 {
																				return int32(0)
																			} else {
																				F_pfree(m, v264)
																				mBase = m.M
																				v271 = m.ExcPending
																				if v271 != 0 {
																					return int32(0)
																				} else {
																					F_ReleaseCatCache(m, v181)
																					mBase = m.M
																					v274 = m.ExcPending
																					if v274 != 0 {
																						return int32(0)
																					} else {
																						F_relation_close(m, v177, int32(3))
																						mBase = m.M
																						v277 = m.ExcPending
																						if v277 != 0 {
																							return int32(0)
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v279 = m.ExcPending
																							if v279 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v85 + int32(80)
																								m.G0 = v20 + int32(16)
																								return v164
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	v236 = *(*int32)(unsafe.Add(mBase, uint32(v185)+108))
																	if v174 == v236 {
																		if v235 == int32(0) {
																			F_ReleaseCatCache(m, v181)
																			mBase = m.M
																			v274 = m.ExcPending
																			if v274 != 0 {
																				return int32(0)
																			} else {
																				F_relation_close(m, v177, int32(3))
																				mBase = m.M
																				v277 = m.ExcPending
																				if v277 != 0 {
																					return int32(0)
																				} else {
																					F_CommandCounterIncrement(m)
																					mBase = m.M
																					v279 = m.ExcPending
																					if v279 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v85 + int32(80)
																						m.G0 = v20 + int32(16)
																						return v164
																					}
																				}
																			}
																		} else {
																			v256 = v235
																			v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
																			v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
																			mBase = m.M
																			v265 = m.ExcPending
																			if v265 != 0 {
																				return int32(0)
																			} else {
																				F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																				mBase = m.M
																				v269 = m.ExcPending
																				if v269 != 0 {
																					return int32(0)
																				} else {
																					F_pfree(m, v264)
																					mBase = m.M
																					v271 = m.ExcPending
																					if v271 != 0 {
																						return int32(0)
																					} else {
																						F_ReleaseCatCache(m, v181)
																						mBase = m.M
																						v274 = m.ExcPending
																						if v274 != 0 {
																							return int32(0)
																						} else {
																							F_relation_close(m, v177, int32(3))
																							mBase = m.M
																							v277 = m.ExcPending
																							if v277 != 0 {
																								return int32(0)
																							} else {
																								F_CommandCounterIncrement(m)
																								mBase = m.M
																								v279 = m.ExcPending
																								if v279 != 0 {
																									return int32(0)
																								} else {
																									m.G0 = v85 + int32(80)
																									m.G0 = v20 + int32(16)
																									return v164
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	} else {
																		*(*int64)(unsafe.Add(mBase, uint32(v85+int32(32)+v235<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v174)
																		*(*int32)(unsafe.Add(mBase, uint32(v85-int32(-64)+v235<<(uint(int32(2))%32)))) = int32(13)
																		v256 = v235 + int32(1)
																		v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
																		v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
																		mBase = m.M
																		v265 = m.ExcPending
																		if v265 != 0 {
																			return int32(0)
																		} else {
																			F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																			mBase = m.M
																			v269 = m.ExcPending
																			if v269 != 0 {
																				return int32(0)
																			} else {
																				F_pfree(m, v264)
																				mBase = m.M
																				v271 = m.ExcPending
																				if v271 != 0 {
																					return int32(0)
																				} else {
																					F_ReleaseCatCache(m, v181)
																					mBase = m.M
																					v274 = m.ExcPending
																					if v274 != 0 {
																						return int32(0)
																					} else {
																						F_relation_close(m, v177, int32(3))
																						mBase = m.M
																						v277 = m.ExcPending
																						if v277 != 0 {
																							return int32(0)
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v279 = m.ExcPending
																							if v279 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v85 + int32(80)
																								m.G0 = v20 + int32(16)
																								return v164
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v286 = m.ExcPending
																if v286 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v81
																	F_errmsg_internal(m, int32(_a_F_relation_statistics_update_5), v85)
																	mBase = m.M
																	v290 = m.ExcPending
																	if v290 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_relation_statistics_update_3), int32(166), int32(_a_F_relation_statistics_update_6))
																		mBase = m.M
																		v295 = m.ExcPending
																		if v295 != 0 {
																			return int32(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v142 = int32(_a_F_relation_statistics_update_9)
													v143 = int32(139)
													F_errcode(m, int32(50856066))
													mBase = m.M
													v146 = m.ExcPending
													if v146 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v85)+16)) = int32(_a_F_relation_statistics_update_8)
														F_errmsg(m, v142, v85+int32(16))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_relation_statistics_update_3), v143, int32(_a_F_relation_statistics_update_6))
															mBase = m.M
															v156 = m.ExcPending
															if v156 != 0 {
																return int32(0)
															} else {
																v162 = v2
																v163 = v109
																v164 = int32(0)
																v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
																if v165 == int32(0) {
																	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
																	v169 = v168
																} else {
																	v169 = v2
																}
																v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+112)))
																if v170 == int32(0) {
																	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
																	v174 = v173
																} else {
																	v174 = v2
																}
																v177 = F_table_open(m, int32(1259), int32(3))
																mBase = m.M
																v178 = m.ExcPending
																if v178 != 0 {
																	return int32(0)
																} else {
																	v181 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v81))
																	mBase = m.M
																	v182 = m.ExcPending
																	if v182 != 0 {
																		return int32(0)
																	} else {
																		if v181 != 0 {
																			v183 = *(*int32)(unsafe.Add(mBase, uint32(v181)+16))
																			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+22)))
																			v185 = v183 + v184
																			v186 = int32(0)
																			v188 = v85 - int32(-64)
																			v190 = v85 + int32(32)
																			if v101 != 0 {
																				v202 = v186
																				v203 = v188
																				v204 = v190
																			} else {
																				v191 = *(*int32)(unsafe.Add(mBase, uint32(v185)+96))
																				if v105 == v191 {
																					v202 = v186
																					v203 = v188
																					v204 = v190
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v85)+64)) = int32(10)
																					*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = base.I64_extend_i32_u(v105)
																					v202 = int32(1)
																					v203 = v188 | int32(4)
																					v204 = v190 | int32(8)
																				}
																			}
																			if v162 == int32(0) {
																				v216 = v202
																			} else {
																				v207 = *(*float32)(unsafe.Add(mBase, uint32(v185)+100))
																				if base.F32_eq(v163, v207) != 0 {
																					v216 = v202
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v203))) = int32(11)
																					*(*int64)(unsafe.Add(mBase, uint32(v204))) = base.I64_extend_i32_s(base.I32_reinterpret_f32(v163))
																					v216 = v202 + int32(1)
																				}
																			}
																			if v165 != 0 {
																				v235 = v216
																			} else {
																				v217 = *(*int32)(unsafe.Add(mBase, uint32(v185)+104))
																				if v169 == v217 {
																					v235 = v216
																				} else {
																					*(*int64)(unsafe.Add(mBase, uint32(v85+int32(32)+v216<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v169)
																					*(*int32)(unsafe.Add(mBase, uint32(v85-int32(-64)|v216<<(uint(int32(2))%32)))) = int32(12)
																					v235 = v216 + int32(1)
																				}
																			}
																			if v170 != 0 {
																				if v235 == int32(0) {
																					F_ReleaseCatCache(m, v181)
																					mBase = m.M
																					v274 = m.ExcPending
																					if v274 != 0 {
																						return int32(0)
																					} else {
																						F_relation_close(m, v177, int32(3))
																						mBase = m.M
																						v277 = m.ExcPending
																						if v277 != 0 {
																							return int32(0)
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v279 = m.ExcPending
																							if v279 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v85 + int32(80)
																								m.G0 = v20 + int32(16)
																								return v164
																							}
																						}
																					}
																				} else {
																					v256 = v235
																					v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
																					v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
																					mBase = m.M
																					v265 = m.ExcPending
																					if v265 != 0 {
																						return int32(0)
																					} else {
																						F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																						mBase = m.M
																						v269 = m.ExcPending
																						if v269 != 0 {
																							return int32(0)
																						} else {
																							F_pfree(m, v264)
																							mBase = m.M
																							v271 = m.ExcPending
																							if v271 != 0 {
																								return int32(0)
																							} else {
																								F_ReleaseCatCache(m, v181)
																								mBase = m.M
																								v274 = m.ExcPending
																								if v274 != 0 {
																									return int32(0)
																								} else {
																									F_relation_close(m, v177, int32(3))
																									mBase = m.M
																									v277 = m.ExcPending
																									if v277 != 0 {
																										return int32(0)
																									} else {
																										F_CommandCounterIncrement(m)
																										mBase = m.M
																										v279 = m.ExcPending
																										if v279 != 0 {
																											return int32(0)
																										} else {
																											m.G0 = v85 + int32(80)
																											m.G0 = v20 + int32(16)
																											return v164
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			} else {
																				v236 = *(*int32)(unsafe.Add(mBase, uint32(v185)+108))
																				if v174 == v236 {
																					if v235 == int32(0) {
																						F_ReleaseCatCache(m, v181)
																						mBase = m.M
																						v274 = m.ExcPending
																						if v274 != 0 {
																							return int32(0)
																						} else {
																							F_relation_close(m, v177, int32(3))
																							mBase = m.M
																							v277 = m.ExcPending
																							if v277 != 0 {
																								return int32(0)
																							} else {
																								F_CommandCounterIncrement(m)
																								mBase = m.M
																								v279 = m.ExcPending
																								if v279 != 0 {
																									return int32(0)
																								} else {
																									m.G0 = v85 + int32(80)
																									m.G0 = v20 + int32(16)
																									return v164
																								}
																							}
																						}
																					} else {
																						v256 = v235
																						v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
																						v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
																						mBase = m.M
																						v265 = m.ExcPending
																						if v265 != 0 {
																							return int32(0)
																						} else {
																							F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																							mBase = m.M
																							v269 = m.ExcPending
																							if v269 != 0 {
																								return int32(0)
																							} else {
																								F_pfree(m, v264)
																								mBase = m.M
																								v271 = m.ExcPending
																								if v271 != 0 {
																									return int32(0)
																								} else {
																									F_ReleaseCatCache(m, v181)
																									mBase = m.M
																									v274 = m.ExcPending
																									if v274 != 0 {
																										return int32(0)
																									} else {
																										F_relation_close(m, v177, int32(3))
																										mBase = m.M
																										v277 = m.ExcPending
																										if v277 != 0 {
																											return int32(0)
																										} else {
																											F_CommandCounterIncrement(m)
																											mBase = m.M
																											v279 = m.ExcPending
																											if v279 != 0 {
																												return int32(0)
																											} else {
																												m.G0 = v85 + int32(80)
																												m.G0 = v20 + int32(16)
																												return v164
																											}
																										}
																									}
																								}
																							}
																						}
																					}
																				} else {
																					*(*int64)(unsafe.Add(mBase, uint32(v85+int32(32)+v235<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v174)
																					*(*int32)(unsafe.Add(mBase, uint32(v85-int32(-64)+v235<<(uint(int32(2))%32)))) = int32(13)
																					v256 = v235 + int32(1)
																					v257 = *(*int32)(unsafe.Add(mBase, uint32(v177)+52))
																					v264 = F_heap_modify_tuple_by_cols(m, v181, v257, v256, v85-int32(-64), v85+int32(32), v85+int32(28))
																					mBase = m.M
																					v265 = m.ExcPending
																					if v265 != 0 {
																						return int32(0)
																					} else {
																						F_CatalogTupleUpdate(m, v177, v264+int32(4), v264)
																						mBase = m.M
																						v269 = m.ExcPending
																						if v269 != 0 {
																							return int32(0)
																						} else {
																							F_pfree(m, v264)
																							mBase = m.M
																							v271 = m.ExcPending
																							if v271 != 0 {
																								return int32(0)
																							} else {
																								F_ReleaseCatCache(m, v181)
																								mBase = m.M
																								v274 = m.ExcPending
																								if v274 != 0 {
																									return int32(0)
																								} else {
																									F_relation_close(m, v177, int32(3))
																									mBase = m.M
																									v277 = m.ExcPending
																									if v277 != 0 {
																										return int32(0)
																									} else {
																										F_CommandCounterIncrement(m)
																										mBase = m.M
																										v279 = m.ExcPending
																										if v279 != 0 {
																											return int32(0)
																										} else {
																											m.G0 = v85 + int32(80)
																											m.G0 = v20 + int32(16)
																											return v164
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		} else {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v286 = m.ExcPending
																			if v286 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v85))) = v81
																				F_errmsg_internal(m, int32(_a_F_relation_statistics_update_5), v85)
																				mBase = m.M
																				v290 = m.ExcPending
																				if v290 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_relation_statistics_update_3), int32(166), int32(_a_F_relation_statistics_update_6))
																					mBase = m.M
																					v295 = m.ExcPending
																					if v295 != 0 {
																						return int32(0)
																					} else {
																						base.Wasm_trap_unreachable()
																						for {
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
