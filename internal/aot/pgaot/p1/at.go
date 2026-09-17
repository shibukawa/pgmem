package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ATExecAlterConstrEnforceability(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
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
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	F_check_stack_depth(m)
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
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l6)+16))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
	v30 = v28 + v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+80))
	v33 = F_table_open(m, v32, l7)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+75)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
	if v35 != v36 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	F_relation_close(m, v33, int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L54
	}
L5:
	;
	v334 = int32(1)
	goto L4
L6:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	F_AlterConstrEnforceabilityRecurse(m, l0, l1, l2, l3, l4, l5, l6, l7, v295, v296, v297, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L53
	}
L7:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v30)+96))
	v272 = F_get_rel_relkind(m, v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L51
	}
L8:
	;
	if v237&int32(255) == int32(112) {
		goto L6
	} else {
		goto L50
	}
L9:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v212)+64))
	v224 = F_lappend(m, v223, v92)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L49
	}
L10:
	;
	v182 = F_palloc0(m, int32(140))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L46
	}
L11:
	;
	v147 = int32(0)
	goto L42
L12:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+119)))
	if v115 != int32(112) {
		goto L35
	} else {
		goto L36
	}
L13:
	;
	F_AlterConstrUpdateConstraintEntry(m, l1, l2, l6)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if v35 != 0 {
		v334 = int32(0)
		goto L4
	} else {
		goto L33
	}
L16:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
	if v40 == int32(0) {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v43 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v43
	v52 = F_palloc0(m, int32(108))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = int32(161)
	v58 = F_pstrdup(m, v30+int32(4))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v58
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+102)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+86)) = uint8(v61)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+87)) = uint8(v63)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+101)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+88)) = uint8(v65)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v30)+80))
	if l4 == v67 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v30)+96))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v30)+88))
	F_createForeignKeyActionTriggers(m, l4, v69, v52, v31, v70, l8, l9, v22+int32(12), v22+int32(8))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v30)+96))
	if l5 == v77 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v30)+80))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v30)+88))
	F_createForeignKeyCheckTriggers(m, v79, l5, v52, v31, v80, l10, l11, v22+int32(4), v22)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+119)))
	if v86 != int32(114) {
		v237 = v86
		goto L8
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v30)+96))
	if v89 != l5 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v92 = F_palloc0(m, int32(32))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v94
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v30)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v30)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v102
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v33)+56))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v106 == int32(0) {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v109 <= int32(0) {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	goto L11
L33:
	;
	goto L12
L34:
	;
	v130 = int32(0)
	F_DropForeignKeyConstraintTriggers(m, l3, v31, v130, v130)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L41
	}
L35:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v30)+96))
	v119 = F_get_rel_relkind(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v123 = int32(0)
	F_AlterConstrEnforceabilityRecurse(m, l0, l1, l2, l3, l4, l5, l6, l7, v123, v123, v123, v123)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	if v119 != int32(112) {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	goto L34
L41:
	;
	v334 = base.B2i32(v35 != v36)
	goto L4
L42:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v112+v147<<(uint(int32(2))%32))))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v157 == v105 {
		v212 = v156
		goto L9
	} else {
		goto L44
	}
L43:
	;
	goto L10
L44:
	;
	v160 = v147 + int32(1)
	if v109 != v160 {
		v147 = v160
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v105
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+4)) = uint8(v188)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v191 = F_CreateTupleDescCopyConstr(m, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+8)) = v191
	*(*int64)(unsafe.Add(mBase, uint32(v182)+88)) = int64(0)
	v196 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+84)) = uint8(v196)
	v198 = int32(_a_F_ATExecAlterConstrEnforceability_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v182)+96)) = uint16(v198)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v201 = F_lappend(m, v200, v182)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v201
	v212 = v182
	goto L9
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+64)) = v224
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+119)))
	v237 = v228
	goto L8
L50:
	;
	goto L7
L51:
	;
	if v272 != int32(112) {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	goto L6
L53:
	;
	goto L5
L54:
	;
	m.G0 = v22 + int32(16)
	return v334
}
func F_ATTypedTableRecursion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v13 = F_find_typed_table_dependencies(m, v9, v8+int32(4), v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L18
	}
L2:
	;
	return
L3:
	;
	return
L4:
	;
	if v13 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v17 <= int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v26 = int32(0)
	goto L7
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v26<<(uint(int32(2))%32))))
	v32 = F_relation_open(m, v31, l3)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L2
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+118)))
	if v35 == int32(116) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)))
	if v38 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_CheckTableNotInUse(m, v32, int32(_a_F_ATTypedTableRecursion_0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	v44 = int32(1)
	F_ATPrepCmd(m, l0, v32, l2, v44, v44, l3, l4)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	F_relation_close(m, v32, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v52 = v26 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v52 < v53 {
		v26 = v52
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L8
L18:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	F_errmsg(m, int32(_a_F_ATTypedTableRecursion_1), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_ATTypedTableRecursion_2), int32(_a_F_ATTypedTableRecursion_3), int32(_a_F_ATTypedTableRecursion_4))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
