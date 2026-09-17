package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AlterObjectTypeCommandTag(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = l0 - int32(1)
	if base.Ui32(v3) <= base.Ui32(int32(50)) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(int32(2))%32))+uint32(_c_F_AlterObjectTypeCommandTag[0])))
		v10 = v8
	} else {
		v10 = int32(0)
	}
	return v10
}
func F_AlterRelationNamespaceInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v17 = F_SearchSysCacheLockedCopy1(m, int32(57), l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L2
	} else {
		goto L41
	}
L2:
	;
	return
L3:
	;
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
	v21 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = int32(1259)
	v27 = v12 + int32(-12)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v34 = v32 - int32(1)
	if v34 < v21 {
		v70 = v21
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L38
	}
L7:
	;
	if v70 != 0 {
		goto L32
	} else {
		goto L33
	}
L8:
	;
	if v70|base.B2i32(l2 == l3) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L9:
	;
	goto L8
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v41 = v34
	goto L11
L11:
	;
	v47 = v38 + v41*int32(12)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v37 != v48 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v70 = v64
	goto L9
L13:
	;
	v64 = int32(0)
	if v64 < v41 {
		v41 = v41 - int32(1)
		goto L11
	} else {
		goto L17
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v50 != v51 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if base.B2i32(v54 == v55)|base.B2i32(v54 == int32(0)) != 0 {
		v70 = int32(1)
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	goto L12
L18:
	;
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+48)) = uint16(v77)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v79
	v81 = v20 + v19
	v83 = v81 + int32(4)
	v84 = F_get_relname_relid(m, v83, l3)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	F_UnlockTuple(m, l0, v17+int32(4), int32(7))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L31
	}
L21:
	;
	if v84 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+68)) = l3
	v88 = v12 + int32(-20)
	F_CatalogTupleUpdate(m, l0, v88, v17)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	F_UnlockTuple(m, l0, v88, int32(7))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	if l4 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v98 = F_changeDependencyFor(m, int32(1259), l1, int32(2615), l2, l3)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	if v98 == int32(1) {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v83
	F_errmsg_internal(m, int32(_a_F_AlterRelationNamespaceInternal_0), v12+int32(-48))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_AlterRelationNamespaceInternal_1), int32(_a_F_AlterRelationNamespaceInternal_2), int32(_a_F_AlterRelationNamespaceInternal_3))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	goto L7
L32:
	;
	F_pfree(m, v17)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L37
	}
L33:
	;
	F_add_exact_object_address(m, v12+int32(-12), l5)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_AlterRelationNamespaceInternal[0]))
	if v129 == int32(0) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v133 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), l1, v133, v133, v133)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	m.G0 = v14 - int32(-64)
	return
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l1
	F_errmsg_internal(m, int32(_a_F_AlterRelationNamespaceInternal_4), v14)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_AlterRelationNamespaceInternal_1), int32(_a_F_AlterRelationNamespaceInternal_5), int32(_a_F_AlterRelationNamespaceInternal_3))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
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
	F_errcode(m, int32(117571716))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v163 = F_get_namespace_name(m, l3)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v83
	F_errmsg(m, int32(_a_F_AlterRelationNamespaceInternal_6), v12+int32(-32))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_AlterRelationNamespaceInternal_1), int32(_a_F_AlterRelationNamespaceInternal_7), int32(_a_F_AlterRelationNamespaceInternal_3))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_AlterTableNamespaceInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v18 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_AlterRelationNamespaceInternal(m, v18, v20, l1, l2, int32(1), l3)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+72))
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = int32(0)
	v29 = F_AlterTypeNamespaceInternal(m, v25, l2, v26, v26, v26, l3)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v31 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	F_list_free(m, v31)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L30
	}
L9:
	;
	if v31 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v35 <= int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v43 = int32(0)
	goto L12
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v43<<(uint(int32(2))%32))))
	v55 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(1259)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v66 = v64 - int32(1)
	if v66 < v55 {
		v102 = v55
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L8
L14:
	;
	if v102 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	goto L14
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v73 = v66
	goto L17
L17:
	;
	v79 = v70 + v73*int32(12)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v69 != v80 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v102 = v96
	goto L15
L19:
	;
	v96 = int32(0)
	if v96 < v73 {
		v73 = v73 - int32(1)
		goto L17
	} else {
		goto L23
	}
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v82 != v83 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if base.B2i32(v86 == v87)|base.B2i32(v86 == int32(0)) != 0 {
		v102 = int32(1)
		goto L15
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	goto L18
L24:
	;
	F_AlterRelationNamespaceInternal(m, v18, v54, l1, l2, int32(0), l3)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v113 = v43 + int32(1)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v113 < v114 {
		v43 = v113
		goto L12
	} else {
		goto L29
	}
L27:
	;
	F_add_exact_object_address(m, v14, l3)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	goto L13
L30:
	;
	v131 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_ScanKeyInit(m, v14, int32(4), int32(3), int32(184), int32(1259))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v14+int32(48), int32(5), int32(3), int32(184), v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v151 = F_systable_beginscan(m, v131, int32(2674), int32(1), int32(0), int32(2), v14)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v153 = F_systable_getnext(m, v151)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v153 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v159 = v153
	goto L39
L37:
	;
	goto L38
L38:
	;
	F_systable_endscan(m, v151)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L54
	}
L39:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+22)))
	v168 = v166 + v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+20))
	if v169 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L38
L41:
	;
	v198 = F_systable_getnext(m, v151)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L52
	}
L42:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	if v172 != int32(1259) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v168)+8))
	if v175 != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+24)))
	switch v176 - int32(97) {
	case 0, 8:
		goto L45
	default:
		goto L41
	}
L45:
	;
	v179 = int32(8)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v182 = F_relation_open(m, v180, v179)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)+48))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+119)))
	if v185 == int32(83) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	F_AlterRelationNamespaceInternal(m, v18, v188, l1, l2, int32(1), l3)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	v193 = v179
	goto L49
L49:
	;
	F_relation_close(m, v182, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	v193 = int32(0)
	goto L49
L51:
	;
	goto L41
L52:
	;
	if v198 != 0 {
		v159 = v198
		goto L39
	} else {
		goto L53
	}
L53:
	;
	goto L40
L54:
	;
	F_relation_close(m, v131, int32(1))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_AlterConstraintNamespaces(m, v216, l1, l2, int32(0), l3)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_relation_close(m, v18, int32(3))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	m.G0 = v14 + int32(96)
	return
}
func F_AlterTypeNamespaceInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
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
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	v7 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(96)
	m.G0 = v18
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = int32(1247)
	v26 = v18 + int32(84)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v33 = v31 - int32(1)
	if v33 < v7 {
		v69 = v7
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L17
	} else {
		goto L81
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L17
	} else {
		goto L75
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L17
	} else {
		goto L70
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L17
	} else {
		goto L67
	}
L5:
	;
	m.G0 = v18 + int32(96)
	return v192
L6:
	;
	if v69 != 0 {
		v192 = v7
		goto L5
	} else {
		goto L16
	}
L7:
	;
	goto L6
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v40 = v33
	goto L9
L9:
	;
	v46 = v37 + v40*int32(12)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v36 != v47 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v69 = v63
	goto L7
L11:
	;
	v63 = int32(0)
	if v63 < v40 {
		v40 = v40 - int32(1)
		goto L9
	} else {
		goto L15
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v49 != v50 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if base.B2i32(v53 == v54)|base.B2i32(v53 == int32(0)) != 0 {
		v69 = int32(1)
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	goto L10
L16:
	;
	v74 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	v80 = F_SearchSysCacheCopy(m, int32(82), l0, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	if v80 == int32(0) {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+22)))
	v86 = v84 + v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+96))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+68))
	v89 = base.B2i32(v88 == l1)
	if v89 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_CheckSetNamespace(m, v88, l1)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L17
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+79)))
	if v103 != int32(99) {
		v121 = int32(0)
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v96 = v86 + int32(4)
	v97 = int32(0)
	v99 = F_SearchSysCacheExists(m, int32(81), v96, l1, v97, v97)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	if v99 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	if v89 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L28:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v86)+84))
	v107 = F_get_rel_relkind(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	v109 = int32(99)
	v110 = base.B2i32(v107 == v109)
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+79)))
	if v110|base.B2i32(v113 != v109) != 0 {
		v121 = v110
		goto L27
	} else {
		goto L30
	}
L30:
	;
	if l3 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_relation_close(m, v74, int32(3))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L17
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if l4 != 0 {
		goto L2
	} else {
		goto L35
	}
L34:
	;
	v192 = int32(0)
	goto L5
L35:
	;
	v121 = v110
	goto L27
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = l1
	F_CatalogTupleUpdate(m, v74, v80+int32(4), v80)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L17
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if v121 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L38
L40:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_AlterTypeNamespaceInternal[0]))
	if v167 != 0 {
		goto L58
	} else {
		goto L59
	}
L41:
	;
	v161 = F_changeDependencyFor(m, int32(1247), l0, int32(2615), v88, l1)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L17
	} else {
		goto L56
	}
L42:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+79)))
	if v132 == int32(100) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v144 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L17
	} else {
		goto L51
	}
L45:
	;
	F_AlterConstraintNamespaces(m, l0, v88, l1, int32(1), l5)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L17
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if l2|v89 != 0 {
		goto L40
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+79)))
	if v139 != int32(99) {
		goto L41
	} else {
		goto L50
	}
L50:
	;
	goto L40
L51:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v86)+84))
	F_AlterRelationNamespaceInternal(m, v144, v146, v88, l1, int32(0), l5)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	F_relation_close(m, v144, int32(3))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L17
	} else {
		goto L53
	}
L53:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v86)+84))
	F_AlterConstraintNamespaces(m, v153, v88, l1, int32(0), l5)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	if l2|v89 != 0 {
		goto L40
	} else {
		goto L55
	}
L55:
	;
	goto L41
L56:
	;
	if v161 != int32(1) {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L40
L58:
	;
	v169 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1247), l0, v169, v169, v169)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L17
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	F_pfree(m, v80)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L17
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	F_relation_close(m, v74, int32(3))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L17
	} else {
		goto L63
	}
L63:
	;
	F_add_exact_object_address(m, v18+int32(84), l5)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L17
	} else {
		goto L64
	}
L64:
	;
	if v87 == int32(0) {
		v192 = v88
		goto L5
	} else {
		goto L65
	}
L65:
	;
	v185 = int32(1)
	v188 = F_AlterTypeNamespaceInternal(m, v87, l1, v185, int32(0), v185, l5)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L17
	} else {
		goto L66
	}
L66:
	;
	v192 = v88
	goto L5
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
	F_errmsg_internal(m, int32(_a_F_AlterTypeNamespaceInternal_0), v18)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L17
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_AlterTypeNamespaceInternal_1), int32(_a_F_AlterTypeNamespaceInternal_2), int32(_a_F_AlterTypeNamespaceInternal_3))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L17
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
	F_errcode(m, int32(_a_F_AlterTypeNamespaceInternal_4))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L17
	} else {
		goto L71
	}
L71:
	;
	v223 = F_get_namespace_name(m, l1)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L17
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v96
	F_errmsg(m, int32(_a_F_AlterTypeNamespaceInternal_5), v18-int32(-64))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L17
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_AlterTypeNamespaceInternal_1), int32(_a_F_AlterTypeNamespaceInternal_6), int32(_a_F_AlterTypeNamespaceInternal_3))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L17
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L17
	} else {
		goto L76
	}
L76:
	;
	v244 = F_format_type_be(m, l0)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L17
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v244
	F_errmsg(m, int32(_a_F_AlterTypeNamespaceInternal_7), v18+int32(48))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L17
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(_a_F_AlterTypeNamespaceInternal_8)
	F_errhint(m, int32(_a_F_AlterTypeNamespaceInternal_9), v18+int32(32))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L17
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_AlterTypeNamespaceInternal_1), int32(_a_F_AlterTypeNamespaceInternal_10), int32(_a_F_AlterTypeNamespaceInternal_3))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L17
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	v268 = F_format_type_be(m, l0)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L17
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v268
	F_errmsg_internal(m, int32(_a_F_AlterTypeNamespaceInternal_11), v18+int32(16))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L17
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_AlterTypeNamespaceInternal_1), int32(_a_F_AlterTypeNamespaceInternal_12), int32(_a_F_AlterTypeNamespaceInternal_3))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L17
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__equalAlterUserMappingStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
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
	var v53 int32
	_ = v53
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v8 = F_equal(m, v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v53
L2:
	;
	return int32(0)
L3:
	;
	if v8 == int32(0) {
		v53 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v49 = F_equal(m, v47, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L19
	}
L6:
	;
	if v14 == int32(0) {
		v53 = v3
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if v14 != v15 {
		v53 = v3
		goto L1
	} else {
		goto L18
	}
L9:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if base.B2i32(v20 == int32(0))|base.B2i32(v20 != v23) != 0 {
		v41 = v20
		v42 = v23
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v41-v42 == int32(0) {
		goto L5
	} else {
		goto L17
	}
L11:
	;
	goto L10
L12:
	;
	v26 = v15
	v27 = v14
	goto L13
L13:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	if v31 == int32(0) {
		v41 = v31
		v42 = v30
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v41 = v31
	v42 = v30
	goto L11
L15:
	;
	v34 = int32(1)
	if v31 == v30 {
		v26 = v26 + v34
		v27 = v27 + v34
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v53 = v3
	goto L1
L18:
	;
	goto L5
L19:
	;
	v53 = v49
	goto L1
}
