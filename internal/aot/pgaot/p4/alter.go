package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AlterObjectTypeCommandTag(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v4 = l0 - int32(1)
	if base.Ui32(v4) <= base.Ui32(int32(50)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v4<<(uint(int32(2))%32))+uint32(_consts[1224])))
		v12 = v11
	} else {
		v12 = int32(0)
	}
	return v12
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
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
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L42
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
		v69 = v21
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L39
	}
L7:
	;
	if l2 == l3 {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	goto L7
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v41 = v34
	goto L10
L10:
	;
	v47 = v38 + v41*int32(12)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v37 != v48 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v69 = v63
	goto L8
L12:
	;
	v63 = int32(0)
	if v63 < v41 {
		v41 = v41 - int32(1)
		goto L10
	} else {
		goto L17
	}
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v50 != v51 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v53 = int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v54 == v55 {
		v69 = v53
		goto L8
	} else {
		goto L15
	}
L15:
	;
	if v54 == int32(0) {
		v69 = v53
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	goto L11
L18:
	;
	if v69 != 0 {
		goto L33
	} else {
		goto L34
	}
L19:
	;
	F_UnlockTuple(m, l0, v17+int32(4), int32(7))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L2
	} else {
		goto L32
	}
L20:
	;
	if v69 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+48)) = uint16(v73)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v75
	v77 = v20 + v19
	v79 = v77 + int32(4)
	v80 = F_get_relname_relid(m, v79, l3)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	if v80 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+68)) = l3
	F_CatalogTupleUpdate(m, l0, v12+int32(-20), v17)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	F_UnlockTuple(m, l0, v12+int32(-20), int32(7))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	if l4 == int32(0) {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v96 = F_changeDependencyFor(m, int32(1259), l1, int32(2615), l2, l3)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	if v96 == int32(1) {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v79
	F_errmsg_internal(m, int32(704007), v12+int32(-48))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(494212), int32(19110), int32(312688))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	goto L18
L33:
	;
	F_pfree(m, v17)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L38
	}
L34:
	;
	F_add_exact_object_address(m, v12+int32(-12), l5)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _consts[389]))
	if v127 == int32(0) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v131 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), l1, v131, v131, v131)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	m.G0 = v14 - int32(-64)
	return
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l1
	F_errmsg_internal(m, int32(46249), v14)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(494212), int32(19067), int32(312688))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_errcode(m, int32(117571716))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v161 = F_get_namespace_name(m, l3)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v79
	F_errmsg(m, int32(723435), v12+int32(-32))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(494212), int32(19093), int32(312688))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
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
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L31
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
		v101 = v55
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L8
L14:
	;
	if v101 == int32(0) {
		goto L25
	} else {
		goto L26
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
	v101 = v95
	goto L15
L19:
	;
	v95 = int32(0)
	if v95 < v73 {
		v73 = v73 - int32(1)
		goto L17
	} else {
		goto L24
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
	v85 = int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v86 == v87 {
		v101 = v85
		goto L15
	} else {
		goto L22
	}
L22:
	;
	if v86 == int32(0) {
		v101 = v85
		goto L15
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	goto L18
L25:
	;
	F_AlterRelationNamespaceInternal(m, v18, v54, l1, l2, int32(0), l3)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v112 = v43 + int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v112 < v113 {
		v43 = v112
		goto L12
	} else {
		goto L30
	}
L28:
	;
	F_add_exact_object_address(m, v14, l3)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L13
L31:
	;
	v130 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_ScanKeyInit(m, v14, int32(4), int32(3), int32(184), int32(1259))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v14+int32(48), int32(5), int32(3), int32(184), v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v150 = F_systable_beginscan(m, v130, int32(2674), int32(1), int32(0), int32(2), v14)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v152 = F_systable_getnext(m, v150)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v152 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v158 = v152
	goto L40
L38:
	;
	goto L39
L39:
	;
	F_systable_endscan(m, v150)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L55
	}
L40:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+22)))
	v167 = v165 + v166
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	if v168 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L39
L42:
	;
	v197 = F_systable_getnext(m, v150)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L53
	}
L43:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v171 != int32(1259) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	if v174 != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+24)))
	switch v175 - int32(97) {
	case 0, 8:
		goto L46
	default:
		goto L42
	}
L46:
	;
	v178 = int32(8)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v181 = F_relation_open(m, v179, v178)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v181)+48))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+119)))
	if v184 == int32(83) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	F_AlterRelationNamespaceInternal(m, v18, v187, l1, l2, int32(1), l3)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	v192 = v178
	goto L50
L50:
	;
	F_relation_close(m, v181, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L52
	}
L51:
	;
	v192 = int32(0)
	goto L50
L52:
	;
	goto L42
L53:
	;
	if v197 != 0 {
		v158 = v197
		goto L40
	} else {
		goto L54
	}
L54:
	;
	goto L41
L55:
	;
	F_relation_close(m, v130, int32(1))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_AlterConstraintNamespaces(m, v215, l1, l2, int32(0), l3)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_sequence_close(m, v18, int32(3))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
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
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
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
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
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
		v68 = v7
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L18
	} else {
		goto L85
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L18
	} else {
		goto L79
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L18
	} else {
		goto L74
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L18
	} else {
		goto L71
	}
L5:
	;
	m.G0 = v18 + int32(96)
	return v187
L6:
	;
	if v68 != 0 {
		v187 = v7
		goto L5
	} else {
		goto L17
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
	v68 = v62
	goto L7
L11:
	;
	v62 = int32(0)
	if v62 < v40 {
		v40 = v40 - int32(1)
		goto L9
	} else {
		goto L16
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
	v52 = int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v53 == v54 {
		v68 = v52
		goto L7
	} else {
		goto L14
	}
L14:
	;
	if v53 == int32(0) {
		v68 = v52
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	goto L10
L17:
	;
	v73 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v79 = F_SearchSysCacheCopy(m, int32(82), l0, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if v79 == int32(0) {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+22)))
	v85 = v83 + v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+96))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)+68))
	v88 = base.B2i32(v87 == l1)
	if v88 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_CheckSetNamespace(m, v87, l1)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L18
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+79)))
	if v102 != int32(99) {
		v119 = int32(0)
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v95 = v85 + int32(4)
	v96 = int32(0)
	v98 = F_SearchSysCacheExists(m, int32(81), v95, l1, v96, v96)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	if v98 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	if v88 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v85)+84))
	v106 = F_get_rel_relkind(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	v108 = int32(99)
	v109 = base.B2i32(v106 == v108)
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+79)))
	if v110 != v108 {
		v119 = v109
		goto L28
	} else {
		goto L31
	}
L31:
	;
	if v106 == int32(99) {
		v119 = v109
		goto L28
	} else {
		goto L32
	}
L32:
	;
	if l3 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_sequence_close(m, v73, int32(3))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L18
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if l4 != 0 {
		goto L2
	} else {
		goto L37
	}
L36:
	;
	v187 = int32(0)
	goto L5
L37:
	;
	v119 = v109
	goto L28
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+68)) = l1
	F_CatalogTupleUpdate(m, v73, v79+int32(4), v79)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L18
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v119 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	goto L40
L42:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _consts[389]))
	if v163 != 0 {
		goto L62
	} else {
		goto L63
	}
L43:
	;
	v157 = F_changeDependencyFor(m, int32(1247), l0, int32(2615), v87, l1)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L18
	} else {
		goto L60
	}
L44:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+79)))
	if v130 == int32(100) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v141 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L18
	} else {
		goto L54
	}
L47:
	;
	F_AlterConstraintNamespaces(m, l0, v87, l1, int32(1), l5)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L18
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v87 == l1 {
		goto L42
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	if l2 != 0 {
		goto L42
	} else {
		goto L52
	}
L52:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+79)))
	if v136 != int32(99) {
		goto L43
	} else {
		goto L53
	}
L53:
	;
	goto L42
L54:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v85)+84))
	F_AlterRelationNamespaceInternal(m, v141, v143, v87, l1, int32(0), l5)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L18
	} else {
		goto L55
	}
L55:
	;
	F_sequence_close(m, v141, int32(3))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L18
	} else {
		goto L56
	}
L56:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v85)+84))
	F_AlterConstraintNamespaces(m, v150, v87, l1, int32(0), l5)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L18
	} else {
		goto L57
	}
L57:
	;
	if l2 != 0 {
		goto L42
	} else {
		goto L58
	}
L58:
	;
	if v87 == l1 {
		goto L42
	} else {
		goto L59
	}
L59:
	;
	goto L43
L60:
	;
	if v157 != int32(1) {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L42
L62:
	;
	v165 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1247), l0, v165, v165, v165)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L18
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	F_pfree(m, v79)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L18
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	F_sequence_close(m, v73, int32(3))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L18
	} else {
		goto L67
	}
L67:
	;
	F_add_exact_object_address(m, v18+int32(84), l5)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L18
	} else {
		goto L68
	}
L68:
	;
	if v86 == int32(0) {
		v187 = v87
		goto L5
	} else {
		goto L69
	}
L69:
	;
	v181 = int32(1)
	v184 = F_AlterTypeNamespaceInternal(m, v86, l1, v181, int32(0), v181, l5)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L18
	} else {
		goto L70
	}
L70:
	;
	v187 = v87
	goto L5
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
	F_errmsg_internal(m, int32(50314), v18)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L18
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(494201), int32(4183), int32(312719))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L18
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L18
	} else {
		goto L75
	}
L75:
	;
	v218 = F_get_namespace_name(m, l1)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L18
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v95
	F_errmsg(m, int32(723644), v18-int32(-64))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L18
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(494201), int32(4203), int32(312719))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L18
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L18
	} else {
		goto L80
	}
L80:
	;
	v239 = F_format_type_be(m, l0)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L18
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v239
	F_errmsg(m, int32(367007), v18+int32(48))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L18
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(540466)
	F_errhint(m, int32(651391), v18+int32(32))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L18
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(494201), int32(4225), int32(312719))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L18
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v263 = F_format_type_be(m, l0)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L18
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v263
	F_errmsg_internal(m, int32(713603), v18+int32(16))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L18
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(494201), int32(4281), int32(312719))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L18
	} else {
		goto L88
	}
L88:
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
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
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
	return v51
L2:
	;
	return int32(0)
L3:
	;
	if v8 == int32(0) {
		v51 = v3
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
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v48 = F_equal(m, v46, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L20
	}
L6:
	;
	if v14 == int32(0) {
		v51 = v3
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
		v51 = v3
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v21 == int32(0) {
		v40 = v20
		v41 = v21
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v41-v40 == int32(0) {
		goto L5
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	if v20 != v21 {
		v40 = v20
		v41 = v21
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v25 = v15
	v26 = v14
	goto L14
L14:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v30 == int32(0) {
		v40 = v29
		v41 = v30
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v40 = v29
	v41 = v30
	goto L11
L16:
	;
	v33 = int32(1)
	if v29 == v30 {
		v25 = v25 + v33
		v26 = v26 + v33
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v51 = v3
	goto L1
L19:
	;
	goto L5
L20:
	;
	v51 = v48
	goto L1
}
