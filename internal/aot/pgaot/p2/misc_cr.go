package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_CreateComments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(224)
	m.G0 = v11
	v13 = int32(1)
	if l3 == v5 {
		v37 = v13
		v38 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v40 = v11 + int32(48)
	F_ScanKeyInit(m, v40, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L6
	}
L2:
	;
	v17 = int32(0)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v18 == v17 {
		v37 = v13
		v38 = v17
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(16843009)
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = base.I64_extend_i32_s(l2)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = base.I64_extend_i32_u(l1)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = base.I64_extend_i32_u(l0)
	v32 = F_cstring_to_text(m, l3)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = base.I64_extend_i32_u(v32)
	v37 = v23
	v38 = int32(1)
	goto L1
L6:
	;
	F_ScanKeyInit(m, v11+int32(104), int32(2), int32(3), int32(184), base.I64_extend_i32_u(l1))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v57 = int32(3)
	F_ScanKeyInit(m, v11+int32(160), v57, v57, int32(65), base.I64_extend_i32_s(l2))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v65 = F_table_open(m, int32(2609), int32(3))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	F_systable_endscan(m, v71)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L20
	}
L10:
	;
	v71 = F_systable_beginscan(m, v65, int32(2675), int32(1), int32(0), int32(3), v40)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v73 = F_systable_getnext(m, v71)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v73 == int32(0) {
		v94 = v5
		goto L9
	} else {
		goto L13
	}
L13:
	;
	if v37 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_simple_heap_delete(m, v65, v73+int32(4))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v65)+52))
	v90 = F_heap_modify_tuple(m, v73, v83, v11+int32(16), v11+int32(12), v11+int32(8))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	v94 = v5
	goto L9
L18:
	;
	F_CatalogTupleUpdate(m, v65, v73+int32(4), v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v94 = v90
	goto L9
L20:
	;
	v97 = int32(0)
	if base.B2i32(v38 == v97)|base.B2i32(v94 != v97) == v97 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v65)+52))
	v109 = F_heap_form_tuple(m, v104, v11+int32(16), v11+int32(12))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	v113 = v94
	goto L23
L23:
	;
	if v113 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	F_CatalogTupleInsert(m, v65, v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v113 = v109
	goto L23
L26:
	;
	F_pfree(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_relation_close(m, v65, int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	m.G0 = v11 + int32(224)
	return
}
func F_CreateFakeRelcacheEntry(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = F_palloc0(m, int32(420))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v11 + int32(276)
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v11))) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v20
		v22 = int32(112)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+394)) = uint8(v22)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(-1)
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v26
		v33 = F_pg_sprintf(m, v11+int32(280), int32(_a_F_CreateFakeRelcacheEntry_0), v8+int32(16))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v26
			*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v35
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v38
			v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v8))) = v40
			v43 = F_smgropen(m, v8, int32(-1))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v43
				m.G0 = v8 + int32(32)
				return v11
			}
		}
	}
}
func F_CreateInitDecodingContext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
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
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	F_CheckLogicalDecodingRequirements(m, l2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[0]))
	if v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L66
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L62
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L58
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L55
	}
L7:
	;
	if l0 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L52
	}
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+88))
	if v24 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[1]))
	if v24 != v28 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[2]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	goto L13
L13:
	;
	if v32 == int32(2) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[3]))
	if v36 != 0 {
		goto L3
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v40 = F_strncpy(m, v14+int32(4), l0, int32(64))
	mBase = m.M
	v41 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+63)) = uint8(v41)
	goto L18
L17:
	;
	goto L16
L18:
	;
	v45 = base.AtomicRmwXchg32(m, v21, int32(0), int32(1))
	if v45 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_s_lock(m, v21, int32(_a_F_CreateInitDecodingContext_6))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v14)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+193)) = v49
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v14)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+185)) = v51
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v14)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+177)) = v53
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v14)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+169)) = v55
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+161)) = v57
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v14)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+153)) = v59
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+145)) = v61
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+137)) = v63
	v65 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v21))), uint32(v65))
	if l3 == int64(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[4]))
	v87 = F_LWLockAcquire(m, v83+int32(_a_F_CreateInitDecodingContext_7), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L32
	}
L24:
	;
	F_ReplicationSlotReserveWal(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v74 = base.AtomicRmwXchg32(m, v21, int32(0), int32(1))
	if v74 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L23
L28:
	;
	F_s_lock(m, v21, int32(_a_F_CreateInitDecodingContext_6))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+104)) = l3
	v79 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v21))), uint32(v79))
	goto L23
L31:
	;
	goto L30
L32:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[4]))
	v94 = F_LWLockAcquire(m, v90+int32(512), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v98 = F_GetOldestSafeDecodingTransactionId(m, l1^int32(1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v102 = base.AtomicRmwXchg32(m, v21, int32(0), int32(1))
	if v102 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_s_lock(m, v21, int32(_a_F_CreateInitDecodingContext_6))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+100)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v98
	if l1 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v98
	goto L41
L40:
	;
	goto L41
L41:
	;
	v109 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v21))), uint32(v109))
	F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[4]))
	F_LWLockRelease(m, v116+int32(512))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[4]))
	F_LWLockRelease(m, v122+int32(_a_F_CreateInitDecodingContext_7))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v131 = int32(0)
	v134 = F_StartupDecodingContext(m, v131, l3, v98, l1, v131, int32(1), l2, l4, l5, l6, l7)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v136 = int32(_a_F_CreateInitDecodingContext_8)
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[5]))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[5])) = v139
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
	if v141 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = int32(_a_F_CreateInitDecodingContext_9)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = int32(1058)
	v149 = int32(_a_F_CreateInitDecodingContext_10)
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[6])) = v14 + int32(68)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v14 + int32(80)
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v134)+164)) = uint8(v159)
	*(*uint8)(unsafe.Add(mBase, uint32(v134)+147)) = uint8(v159)
	m.T0[v141].(func(*base.Module, int32, int32, int32))(m, v134, v134+int32(108), int32(1))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[5])) = v137
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+145)))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+136)))
	v176 = v174 & v175
	*(*uint8)(unsafe.Add(mBase, uint32(v134)+145)) = uint8(v176)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+112)))
	*(*uint8)(unsafe.Add(mBase, uint32(v178)+116)) = uint8(v179)
	m.G0 = v14 + int32(96)
	return v134
L51:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[6])) = v169
	goto L50
L52:
	;
	F_errmsg_internal(m, int32(_a_F_CreateInitDecodingContext_11), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_CreateInitDecodingContext_1), int32(435), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errmsg_internal(m, int32(_a_F_CreateInitDecodingContext_0), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_CreateInitDecodingContext_1), int32(438), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_CreateInitDecodingContext_3), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_CreateInitDecodingContext_1), int32(444), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v21 + int32(24)
	F_errmsg(m, int32(_a_F_CreateInitDecodingContext_4), v14)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_CreateInitDecodingContext_1), int32(450), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(16777538))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(_a_F_CreateInitDecodingContext_5), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_CreateInitDecodingContext_1), int32(456), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_crc32c_bytea(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
		if v9 == int32(1) {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
			if v15 == int32(18) {
				v18 = int32(16)
			} else {
				v18 = int32(0)
			}
			if base.Ui32((v15-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v25 = int32(4)
			} else {
				v25 = v18
			}
			v38 = v25
		} else {
			v26 = int32(1)
			if v9&v26 != 0 {
				v38 = int32(base.Ui32(v9)>>(uint(v26)%32)) - v26
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v38 = int32(base.Ui32(v32)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v40 = int32(1)
		if v9&v40 != 0 {
			v44 = v40
		} else {
			v44 = int32(4)
		}
		v46 = m.Env.Pgmem_crc32c(m, int32(-1), v5+v44, v38)
		mBase = m.M
		return base.I64_extend_i32_u(v46 ^ int32(-1))
	}
}
func F_create_incremental_sort_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v39 float64
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v10 = F_palloc0(m, int32(88))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(1576252997938)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v17
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
		v20 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)) = uint8(v20)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v19
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
		if v23 == int32(1) {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
			v28 = v26
		} else {
			v28 = int32(0)
		}
		v30 = v28 & int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v30)
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v32
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
		v37 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
		v38 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
		v39 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
		v43 = *(*int32)(unsafe.Add(mBase, _c_F_create_incremental_sort_path[0]))
		F_cost_incremental_sort(m, v10, l0, l3, l4, v36, v37, v38, v39, v41, v43, l5)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = l4
			return v10
		}
	}
}
func F_create_indexscan_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v446 int32
	_ = v446
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v469 int32
	_ = v469
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v519 int32
	_ = v519
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v582 int32
	_ = v582
	var v597 int32
	_ = v597
	var v599 float64
	_ = v599
	var v601 float64
	_ = v601
	var v603 float64
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	v6 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(32)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v30 == v6 {
		v141 = v29
		v148 = v6
		v150 = v6
		v154 = v25
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v156 = int32(0)
	v163 = v156
	v164 = v156
	goto L18
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v33 <= int32(0) {
		v141 = v29
		v148 = v6
		v150 = v6
		v154 = v25
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v41 = v33
	v45 = v6
	v50 = v6
	v52 = v6
	goto L4
L4:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v45<<(uint(int32(2))%32))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v61 == int32(0) {
		v114 = v41
		v123 = v50
		v125 = v52
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v141 = v132
	v148 = v123
	v150 = v125
	v154 = v133
	goto L1
L6:
	;
	v130 = v45 + int32(1)
	if v130 < v114 {
		v41 = v114
		v45 = v130
		v50 = v123
		v52 = v125
		goto L4
	} else {
		goto L16
	}
L7:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v64 <= int32(0) {
		v114 = v41
		v123 = v50
		v125 = v52
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+14)))
	v74 = int32(0)
	v83 = v50
	v85 = v52
	goto L9
L9:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v74<<(uint(int32(2))%32))))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v95 = F_lappend(m, v83, v94)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v114 = v108
	v123 = v95
	v125 = v102
	goto L6
L11:
	;
	return int32(0)
L12:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v100 = F_fix_indexqual_clause(m, l0, v25, v67, v94, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v102 = F_lappend(m, v85, v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v105 = v74 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v105 < v106 {
		v74 = v105
		v83 = v95
		v85 = v102
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	goto L5
L17:
	;
	if l4 != 0 {
		goto L94
	} else {
		goto L95
	}
L18:
	;
	v178 = int32(0)
	if v141 == v178 {
		v189 = v178
		goto L20
	} else {
		goto L21
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L11
	} else {
		goto L90
	}
L20:
	;
	if v155 != 0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v183 <= v163 {
		v189 = int32(0)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v189 = v185 + v163<<(uint(int32(2))%32)
	goto L20
L23:
	;
	goto L19
L24:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v197+v163<<(uint(int32(2))%32))))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v415 = F_fix_indexqual_clause(m, l0, v154, v412, v413, int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L11
	} else {
		goto L88
	}
L25:
	;
	v190 = int32(0)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if base.B2i32(v189 == v190)|base.B2i32(v192 <= v163) == v190 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v199 = v178
	goto L27
L27:
	;
	v200 = int32(0)
	if l3 == v200 {
		v319 = v200
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	if v197 != 0 {
		goto L24
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v199 = v164
	goto L27
L31:
	;
	goto L30
L32:
	;
	v332 = F_order_qual_clauses(m, l0, v319)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L11
	} else {
		goto L66
	}
L33:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v203 <= int32(0) {
		v319 = v200
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v212 = int32(0)
	v214 = v200
	goto L35
L35:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227+v212<<(uint(int32(2))%32))))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+10)))
	if v232 != 0 {
		v306 = v214
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v319 = v306
	goto L32
L37:
	;
	v309 = v212 + int32(1)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v309 < v310 {
		v212 = v309
		v214 = v306
		goto L35
	} else {
		goto L65
	}
L38:
	;
	if v30 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if v285 != 0 {
		v306 = v214
		goto L37
	} else {
		goto L56
	}
L40:
	;
	goto L39
L41:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v238 <= int32(0) {
		v285 = int32(0)
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v285 = int32(0)
	goto L40
L44:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v231)+60))
	v242 = int32(0)
	if v242 < v238 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v245 = v238
	goto L47
L46:
	;
	v245 = v242
	goto L47
L47:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v249 = int32(0)
	goto L48
L48:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v246+v249<<(uint(int32(2))%32))))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+12)))
	if v259 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L43
L50:
	;
	v270 = v249 + int32(1)
	if v270 != v245 {
		v249 = v270
		goto L48
	} else {
		goto L55
	}
L51:
	;
	v260 = int32(1)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v231 == v261 {
		v285 = v260
		goto L40
	} else {
		goto L52
	}
L52:
	;
	if v241 == int32(0) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v261)+60))
	if v265 == v241 {
		v285 = v260
		goto L40
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	goto L49
L56:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v288 = F_contain_mutable_functions(m, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	if v288 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v292
	v298 = F_list_make1_impl(m, int32(1), v23+int32(24))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L11
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v304 = F_lappend(m, v214, v231)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L11
	} else {
		goto L64
	}
L61:
	;
	v301 = F_predicate_implied_by(m, v298, v148, int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	if v301 != 0 {
		v306 = v214
		goto L37
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	v306 = v304
	goto L37
L65:
	;
	goto L36
L66:
	;
	v335 = F_extract_actual_clauses(m, v332, int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v337 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v338 = F_replace_nestloop_params_mutator(m, v148, l0)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L11
	} else {
		goto L71
	}
L69:
	;
	v344 = v335
	v345 = v29
	v346 = v148
	goto L70
L70:
	;
	if v345 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v340 = F_replace_nestloop_params_mutator(m, v335, l0)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	v342 = F_replace_nestloop_params_mutator(m, v29, l0)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	v344 = v340
	v345 = v342
	v346 = v338
	goto L70
L74:
	;
	v446 = int32(0)
	goto L17
L75:
	;
	goto L76
L76:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v351 = int32(0)
	v358 = v351
	v360 = v351
	goto L77
L77:
	;
	v373 = int32(0)
	if v350 == v373 {
		v383 = v373
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	if base.B2i32(v383 == int32(0))|base.B2i32(v386 <= v358) != 0 {
		v446 = v360
		goto L17
	} else {
		goto L82
	}
L80:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	if v377 <= v358 {
		v383 = int32(0)
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v350)+12))
	v383 = v379 + v358<<(uint(int32(2))%32)
	goto L79
L82:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
	if v389 == int32(0) {
		v446 = v360
		goto L17
	} else {
		goto L83
	}
L83:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v389+v358<<(uint(int32(2))%32))))
	v397 = F_exprType(m, v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L11
	} else {
		goto L84
	}
L84:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v392)+8))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v392)+12))
	v401 = F_get_opfamily_member_for_cmptype(m, v399, v397, v397, v400)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L11
	} else {
		goto L85
	}
L85:
	;
	if v401 == int32(0) {
		goto L23
	} else {
		goto L86
	}
L86:
	;
	v407 = F_lappend_oid(m, v360, v401)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L11
	} else {
		goto L87
	}
L87:
	;
	v358 = v358 + int32(1)
	v360 = v407
	goto L77
L88:
	;
	v417 = F_lappend(m, v164, v415)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L11
	} else {
		goto L89
	}
L89:
	;
	v163 = v163 + int32(1)
	v164 = v417
	goto L18
L90:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v392)+12))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v392)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v426
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v425
	F_errmsg_internal(m, int32(_a_F_create_indexscan_plan_0), v23)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L11
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_create_indexscan_plan_1), int32(2983), int32(_a_F_create_indexscan_plan_2))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v582)+4)) = v597
	v599 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v582)+8)) = v599
	v601 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v582)+16)) = v601
	v603 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v582)+24)) = v603
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v582)+32)) = v606
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v582)+36)) = uint8(v608)
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v582)+37)) = uint8(v610)
	m.G0 = v23 + int32(32)
	return v582
L94:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	if v459 != 0 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v561 = F_palloc0(m, int32(112))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L11
	} else {
		goto L107
	}
L97:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+4))
	if int32(0) < v460 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v541 = int32(0)
	goto L99
L99:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v544 = F_palloc0(m, int32(104))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L11
	} else {
		goto L106
	}
L100:
	;
	v469 = int32(0)
	goto L103
L101:
	;
	goto L102
L102:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	v541 = v519
	goto L99
L103:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v459)+12))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v484+v469<<(uint(int32(2))%32))))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489+v469))))
	v492 = int32(1)
	v493 = v491 ^ v492
	*(*uint8)(unsafe.Add(mBase, uint32(v488)+26)) = uint8(v493)
	v496 = v469 + v492
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v459)+4))
	if v496 < v497 {
		v469 = v496
		goto L103
	} else {
		goto L105
	}
L104:
	;
	goto L102
L105:
	;
	goto L104
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v544)+100)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v544)+96)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v544)+92)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v544)+88)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v544)+84)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v544)+80)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v544)+72)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v544)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v544)+48)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v544)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v544))) = int32(346)
	v582 = v544
	goto L93
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v561)+104)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v561)+100)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v561)+96)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v561)+92)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v561)+88)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v561)+84)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v561)+80)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v561)+72)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v561)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v561)+48)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v561)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v561))) = int32(345)
	v582 = v561
	goto L93
}
func F_create_limit_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64, l6 int64) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 float64
	_ = v40
	var v42 int32
	_ = v42
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 float64
	_ = v56
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v61 int32
	_ = v61
	var v62 float64
	_ = v62
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v79 float64
	_ = v79
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v95 float64
	_ = v95
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v103 float64
	_ = v103
	var v106 float64
	_ = v106
	var v113 float64
	_ = v113
	var v114 float64
	_ = v114
	var v123 float64
	_ = v123
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v129 float64
	_ = v129
	var v130 float64
	_ = v130
	var v132 float64
	_ = v132
	var v138 float64
	_ = v138
	var v141 float64
	_ = v141
	var v144 float64
	_ = v144
	v16 = F_palloc0(m, int32(88))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(1619202670911)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v24 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)) = uint8(v24)
		*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v23
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
		if v29 == int32(1) {
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
			v34 = v32
		} else {
			v34 = int32(0)
		}
		v36 = v34 & int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v16)+21)) = uint8(v36)
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v38
		v40 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v40
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v42
		v44 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
		*(*float64)(unsafe.Add(mBase, uint32(v16)+48)) = v44
		v46 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
		*(*float64)(unsafe.Add(mBase, uint32(v16)+56)) = v46
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
		*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v48
		v55 = v16 + int32(56)
		v56 = *(*float64)(unsafe.Add(mBase, uint32(v55)))
		v58 = v16 + int32(48)
		v59 = *(*float64)(unsafe.Add(mBase, uint32(v58)))
		v61 = v16 + int32(32)
		v62 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
		if l5 != int64(0) {
			if int64(0) < l5 {
				v85 = base.F64_convert_i64_u(l5)
				v86 = v62
			} else {
				v69 = base.F64_mul(v62, float64(0.1))
				v70 = float64(1e+100)
				if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v69)&int64(9223372036854775807)))|base.F64_gt(v69, v70) != 0 {
					v83 = v70
				} else {
					v79 = float64(1)
					if base.F64_le(v69, v79) != 0 {
						v83 = v79
					} else {
						v83 = base.F64_nearest(v69)
					}
				}
				v84 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
				v85 = v83
				v86 = v84
			}
			if base.F64_gt(v85, v86) != 0 {
				v88 = v86
			} else {
				v88 = v85
			}
			if base.F64_gt(v62, float64(0)) != 0 {
				v95 = *(*float64)(unsafe.Add(mBase, uint32(v58)))
				*(*float64)(unsafe.Add(mBase, uint32(v58))) = base.F64_add(base.F64_div(base.F64_mul(base.F64_sub(v56, v59), v88), v62), v95)
				v98 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
				v99 = v98
			} else {
				v99 = v86
			}
			v100 = base.F64_sub(v99, v88)
			if base.F64_lt(v100, float64(1)) != 0 {
				v103 = float64(1)
			} else {
				v103 = v100
			}
			*(*float64)(unsafe.Add(mBase, uint32(v61))) = v103
			v106 = v103
		} else {
			v106 = v62
		}
		if l6 != int64(0) {
			if int64(0) < l6 {
				v129 = base.F64_convert_i64_u(l6)
				v130 = v106
			} else {
				v113 = base.F64_mul(v62, float64(0.1))
				v114 = float64(1e+100)
				if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v113)&int64(9223372036854775807)))|base.F64_gt(v113, v114) != 0 {
					v127 = v114
				} else {
					v123 = float64(1)
					if base.F64_le(v113, v123) != 0 {
						v127 = v123
					} else {
						v127 = base.F64_nearest(v113)
					}
				}
				v128 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
				v129 = v127
				v130 = v128
			}
			if base.F64_gt(v129, v130) != 0 {
				v132 = v130
			} else {
				v132 = v129
			}
			if base.F64_gt(v62, float64(0)) != 0 {
				v138 = *(*float64)(unsafe.Add(mBase, uint32(v58)))
				*(*float64)(unsafe.Add(mBase, uint32(v55))) = base.F64_add(base.F64_div(base.F64_mul(base.F64_sub(v56, v59), v132), v62), v138)
			} else {
			}
			v141 = float64(1)
			if base.F64_lt(v132, v141) != 0 {
				v144 = v141
			} else {
				v144 = v132
			}
			*(*float64)(unsafe.Add(mBase, uint32(v61))) = v144
		} else {
		}
		return v16
	}
}
func F_create_s(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_palloc(m, int32(10))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(1)
			return v3 + int32(8)
		}
	}
}
func F_create_subqueryscan_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v10 = F_palloc0(m, int32(80))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(1507533521186)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v17
		v19 = F_get_baserel_parampathinfo(m, l0, l1, l5)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)) = uint8(v21)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v19
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
			if v24 == int32(1) {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
				v29 = v27
			} else {
				v29 = int32(0)
			}
			v31 = v29 & int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v31)
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v33
			F_cost_subqueryscan(m, v10, l0, l1, v19, l3)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
