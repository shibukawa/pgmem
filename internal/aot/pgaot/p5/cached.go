package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BuildCachedPlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	v5 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCachedPlan[0]))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	if v11 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = F_RevalidateCachedQuery(m, l0, l3)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v18 = l1
	goto L3
L3:
	;
	if v18 != 0 {
		v23 = v18
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v18 = v14
	goto L3
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCachedPlan[1]))
	goto L13
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v20 != 0 {
		v23 = v19
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v21 = F_copyObjectImpl(m, v19)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v23 = v21
	goto L6
L10:
	;
	v83 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCachedPlan[0]))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v86 == v83 {
		goto L33
	} else {
		goto L34
	}
L11:
	;
	v69 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L29
	}
L12:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v66 = F_pg_plan_queries(m, v23, v64, v65, l2)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L28
	}
L13:
	;
	if v25 != int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v28 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	switch v32 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v36 = int32(1)
		goto L19
	default:
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v39 == int32(0) {
		goto L12
	} else {
		goto L22
	}
L18:
	;
	if v36 == int32(0) {
		goto L12
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v36 = int32(0)
	goto L19
L21:
	;
	goto L11
L22:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v43 != int32(6) {
		v58 = int32(1)
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v58&int32(1) != 0 {
		goto L11
	} else {
		goto L27
	}
L24:
	;
	goto L23
L25:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v50 = v48 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v50) {
		v58 = int32(0)
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v58 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v50)) % 64)))
	goto L24
L27:
	;
	goto L12
L28:
	;
	v81 = v66
	v82 = l0 + int32(12)
	goto L10
L29:
	;
	F_PushActiveSnapshot(m, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v77 = F_pg_plan_queries(m, v23, v75, v76, l2)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v81 = v77
	v82 = l0 + int32(12)
	goto L10
L33:
	;
	v93 = F_AllocSetContextCreateInternal(m, v85, int32(_a_F_BuildCachedPlan_0), int32(0), int32(1024), int32(_a_F_BuildCachedPlan_1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	v103 = v81
	v104 = v85
	goto L35
L35:
	;
	v106 = F_palloc(m, int32(36))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L39
	}
L36:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v96 = F_MemoryContextStrdup(m, v93, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+36)) = v96
	*(*int32)(unsafe.Add(mBase, _c_F_BuildCachedPlan[0])) = v93
	v101 = F_copyObjectImpl(m, v81)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v103 = v101
	v104 = v93
	goto L35
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(953717834)
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCachedPlan[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = v112
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+85)))
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+16)) = uint8(v114)
	if v103 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if int32(0) < v116 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v162 = v83
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+32)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v106)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = v162
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	v174 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v106)+9)) = uint16(v174)
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+8)) = uint8(v173)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v179 = v177 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v106)+24)) = v179
	*(*int32)(unsafe.Add(mBase, _c_F_BuildCachedPlan[0])) = v10
	return v106
L43:
	;
	v120 = v83
	v125 = v5
	goto L46
L44:
	;
	v153 = v5
	goto L45
L45:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCachedPlan[3]))
	if v153&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L46:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v120<<(uint(int32(2))%32))))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	if v132 == int32(6) {
		v142 = v125
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v153 = v142
	goto L45
L48:
	;
	v144 = v120 + int32(1)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v144 < v145 {
		v120 = v144
		v125 = v142
		goto L46
	} else {
		goto L51
	}
L49:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+27)))
	v136 = v135 | v125
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+28)))
	if v137 != int32(1) {
		v142 = v136
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v140 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+16)) = uint8(v140)
	v142 = v136
	goto L48
L51:
	;
	goto L47
L52:
	;
	v160 = v156
	goto L54
L53:
	;
	v160 = int32(0)
	goto L54
L54:
	;
	v162 = v160
	goto L42
}
func F_CachedPlanAllowsSimpleValidityCheck(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	v4 = int32(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v10 != 0 {
		v159 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v159
L2:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+85)))
	if v11 != 0 {
		v159 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v12 != 0 {
		v159 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v13 != 0 {
		v159 = v4
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v14 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v57 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v17 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v20 = int32(0)
	if v20 < v17 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v23 = v17
	goto L11
L10:
	;
	v23 = v20
	goto L11
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v29 = v4
	goto L12
L12:
	;
	v34 = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v24+v29<<(uint(int32(2))%32))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v39 == int32(6) {
		v159 = v34
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L6
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
	if v42 != 0 {
		v159 = v34
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	if v43 != 0 {
		v159 = v34
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+39)))
	if v44 != 0 {
		v159 = v34
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v46 = v29 + int32(1)
	if v46 != v23 {
		v29 = v46
		goto L12
	} else {
		goto L18
	}
L18:
	;
	goto L13
L19:
	;
	if l2 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v60 <= int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v63 = int32(0)
	if v63 < v60 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v66 = v60
	goto L24
L23:
	;
	v66 = v63
	goto L24
L24:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v73 = int32(0)
	goto L25
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v67+v73<<(uint(int32(2))%32))))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v82 == int32(6) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L19
L27:
	;
	return int32(0)
L28:
	;
	goto L29
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+44))
	if v87 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v129 = v73 + int32(1)
	if v129 != v66 {
		v73 = v129
		goto L25
	} else {
		goto L40
	}
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v90 <= int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v93 = int32(0)
	if v93 < v90 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = v90
	goto L35
L34:
	;
	v96 = v93
	goto L35
L35:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v98 = int32(0)
	v100 = v98
	goto L36
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v97+v100<<(uint(int32(2))%32))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	if v113 == int32(0) {
		v159 = v98
		goto L1
	} else {
		goto L38
	}
L37:
	;
	goto L30
L38:
	;
	v117 = v100 + int32(1)
	if v117 != v96 {
		v100 = v117
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	goto L26
L41:
	;
	return int32(1)
L42:
	;
	goto L43
L43:
	;
	F_ResourceOwnerEnlarge(m, l2)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	return int32(0)
L45:
	;
	v148 = int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v149 + v148
	F_ResourceOwnerRemember(m, l2, l1, int32(_a_F_CachedPlanAllowsSimpleValidityCheck_0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v159 = v148
	goto L1
}
