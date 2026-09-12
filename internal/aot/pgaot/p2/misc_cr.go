package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(176)
	m.G0 = v11
	v13 = int32(1)
	if l3 == v5 {
		v33 = v13
		v34 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ScanKeyInit(m, v11+int32(32), int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L6
	}
L2:
	;
	v17 = int32(0)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v18 == v17 {
		v33 = v13
		v34 = v17
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(16843009)
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
	v29 = F_cstring_to_text(m, l3)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v29
	v33 = v23
	v34 = int32(1)
	goto L1
L6:
	;
	F_ScanKeyInit(m, v11+int32(80), int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v51 = int32(3)
	F_ScanKeyInit(m, v11+int32(128), v51, v51, int32(65), l2)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v58 = F_table_open(m, int32(2609), int32(3))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	F_systable_endscan(m, v66)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L20
	}
L10:
	;
	v66 = F_systable_beginscan(m, v58, int32(2675), int32(1), int32(0), int32(3), v11+int32(32))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v68 = F_systable_getnext(m, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v68 == int32(0) {
		v89 = v5
		goto L9
	} else {
		goto L13
	}
L13:
	;
	if v33 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_CatalogTupleDelete(m, v58, v68+int32(4))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v58)+52))
	v85 = F_heap_modify_tuple(m, v68, v78, v11+int32(16), v11+int32(12), v11+int32(8))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	v89 = v5
	goto L9
L18:
	;
	F_CatalogTupleUpdate(m, v58, v68+int32(4), v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v89 = v85
	goto L9
L20:
	;
	v92 = int32(0)
	if base.B2i32(v34 == v92)|base.B2i32(v89 != v92) == v92 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v58)+52))
	v104 = F_heap_form_tuple(m, v99, v11+int32(16), v11+int32(12))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	v108 = v89
	goto L23
L23:
	;
	if v108 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	F_CatalogTupleInsert(m, v58, v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v108 = v104
	goto L23
L26:
	;
	F_pfree(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_sequence_close(m, v58, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	m.G0 = v11 + int32(176)
	return
}
func F_CreateFakeRelcacheEntry(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = F_palloc0(m, int32(420))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v12 + int32(276)
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v12))) = v19
		v22 = l0 + int32(8)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v23
		v25 = int32(112)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+394)) = uint8(v25)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(-1)
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v29
		v36 = F_pg_sprintf(m, v12+int32(280), int32(_a_F_CreateFakeRelcacheEntry_0), v9+int32(16))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v29
			*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v38
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v41
			v43 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = v43
			v46 = F_smgropen(m, v9, int32(-1))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v46
				m.G0 = v9 + int32(32)
				return v12
			}
		}
	}
}
func F_CreateInitDecodingContext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	F_CheckLogicalDecodingRequirements(m)
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
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[0]))
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L66
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L62
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L58
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
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
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L52
	}
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	if v23 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[1]))
	if v23 != v27 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[2]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	goto L13
L13:
	;
	if v31 == int32(2) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[3]))
	if v35 != 0 {
		goto L3
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v39 = F_strncpy(m, v13+int32(4), l0, int32(64))
	mBase = m.M
	v40 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+63)) = uint8(v40)
	goto L18
L17:
	;
	goto L16
L18:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1)
	if v42 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_s_lock(m, v20, int32(_a_F_CreateInitDecodingContext_1), int32(387), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v13)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+137)) = v50
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v13)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+193)) = v52
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v13)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+185)) = v54
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v13)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+177)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v13)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+169)) = v58
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+161)) = v60
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v13)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+153)) = v62
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+145)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(0)
	if l2 == int64(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[4]))
	v89 = F_LWLockAcquire(m, v85+int32(_a_F_CreateInitDecodingContext_6), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
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
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1)
	if v72 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L23
L28:
	;
	F_s_lock(m, v20, int32(_a_F_CreateInitDecodingContext_1), int32(395), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+104)) = l2
	goto L23
L31:
	;
	goto L30
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[4]))
	v96 = F_LWLockAcquire(m, v92+int32(512), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v100 = F_GetOldestSafeDecodingTransactionId(m, l1^int32(1))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1)
	if v102 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_s_lock(m, v20, int32(_a_F_CreateInitDecodingContext_1), int32(430), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v100
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v100
	goto L41
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(0)
	F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[4]))
	F_LWLockRelease(m, v119+int32(512))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[4]))
	F_LWLockRelease(m, v125+int32(_a_F_CreateInitDecodingContext_6))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v134 = int32(0)
	v137 = F_StartupDecodingContext(m, v134, l2, v100, l1, v134, int32(1), l3, l4, l5, l6)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v139 = int32(_a_F_CreateInitDecodingContext_7)
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[5]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[5])) = v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v137)+24))
	if v144 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = int32(_a_F_CreateInitDecodingContext_8)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = int32(992)
	v152 = int32(_a_F_CreateInitDecodingContext_9)
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[6])) = v13 + int32(68)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v13 + int32(80)
	v162 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+164)) = uint8(v162)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+147)) = uint8(v162)
	m.T0[v144].(func(*base.Module, int32, int32, int32))(m, v137, v137+int32(108), int32(1))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[5])) = v140
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+145)))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+136)))
	v179 = v177 & v178
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+145)) = uint8(v179)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+112)))
	*(*uint8)(unsafe.Add(mBase, uint32(v181)+116)) = uint8(v182)
	m.G0 = v13 + int32(96)
	return v137
L51:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateInitDecodingContext[6])) = v172
	goto L50
L52:
	;
	F_errmsg_internal(m, int32(_a_F_CreateInitDecodingContext_10), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_CreateInitDecodingContext_1), int32(358), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
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
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_CreateInitDecodingContext_1), int32(361), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
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
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_CreateInitDecodingContext_3), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_CreateInitDecodingContext_1), int32(367), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
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
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v20 + int32(24)
	F_errmsg(m, int32(_a_F_CreateInitDecodingContext_4), v13)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_CreateInitDecodingContext_1), int32(373), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
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
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(_a_F_CreateInitDecodingContext_5), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_CreateInitDecodingContext_1), int32(379), int32(_a_F_CreateInitDecodingContext_2))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
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
func F_crc32c_bytea(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(1)
		v12 = v7 + v11
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v17 = v15 & v11
		if v17 != 0 {
			v18 = v12
		} else {
			v18 = v7 + int32(4)
		}
		if v15 == int32(1) {
			v21 = int32(4)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v23&int32(254) == int32(2) {
				v32 = v21
			} else {
				v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
			}
			if v23 == int32(1) {
				v35 = v21
			} else {
				v35 = v32
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v17 != 0 {
				v46 = int32(base.Ui32(v15)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = m.Env.Pgmem_crc32c(m, int32(-1), v18, v46)
		mBase = m.M
		v51 = F_Int64GetDatum(m, base.I64_extend_i32_u(v47^int32(-1)))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			return v51
		}
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
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
		*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(1559073128752)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v17
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
		if v24 == int32(1) {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
			v28 = v27
		} else {
			v28 = v18
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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
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
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
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
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
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
	var v283 int32
	_ = v283
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
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v441 int32
	_ = v441
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v463 int32
	_ = v463
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v513 int32
	_ = v513
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v576 int32
	_ = v576
	var v591 int32
	_ = v591
	var v593 float64
	_ = v593
	var v595 float64
	_ = v595
	var v597 float64
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	v6 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(32)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v30 == v6 {
		v142 = v29
		v145 = v6
		v149 = v6
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
	v171 = v156
	goto L19
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v33 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v142 = v29
	v145 = v6
	v149 = v6
	v154 = v25
	goto L1
L4:
	;
	goto L5
L5:
	;
	v43 = v33
	v45 = v6
	v47 = v6
	v51 = v6
	goto L6
L6:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v45<<(uint(int32(2))%32))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v61 == int32(0) {
		v116 = v43
		v120 = v47
		v124 = v51
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v142 = v132
	v145 = v120
	v149 = v124
	v154 = v133
	goto L1
L8:
	;
	v130 = v45 + int32(1)
	if v130 < v116 {
		v43 = v116
		v45 = v130
		v47 = v120
		v51 = v124
		goto L6
	} else {
		goto L18
	}
L9:
	;
	v64 = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v65 <= v64 {
		v116 = v43
		v120 = v47
		v124 = v51
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+14)))
	v74 = v64
	v80 = v47
	v84 = v51
	goto L11
L11:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v74<<(uint(int32(2))%32))))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v95 = F_lappend(m, v80, v94)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v116 = v108
	v120 = v95
	v124 = v102
	goto L8
L13:
	;
	return int32(0)
L14:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v100 = F_fix_indexqual_clause(m, l0, v25, v68, v94, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v102 = F_lappend(m, v84, v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v105 = v74 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v105 < v106 {
		v74 = v105
		v80 = v95
		v84 = v102
		goto L11
	} else {
		goto L17
	}
L17:
	;
	goto L12
L18:
	;
	goto L7
L19:
	;
	v178 = int32(0)
	if v142 == v178 {
		v189 = v178
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v341 = F_order_qual_clauses(m, l0, v328)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L13
	} else {
		goto L70
	}
L21:
	;
	if v155 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v183 <= v163 {
		v189 = int32(0)
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	v189 = v185 + v163<<(uint(int32(2))%32)
	goto L21
L24:
	;
	goto L20
L25:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v317 = F_fix_indexqual_clause(m, l0, v154, v314, v315, int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L13
	} else {
		goto L68
	}
L26:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v190 <= v163 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v199 = v178
	goto L28
L28:
	;
	v200 = int32(0)
	if l3 == v200 {
		v328 = v200
		goto L24
	} else {
		goto L33
	}
L29:
	;
	v199 = v171
	goto L28
L30:
	;
	if v189 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	v197 = v194 + v163<<(uint(int32(2))%32)
	if v197 != 0 {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v203 = int32(0)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v204 <= v203 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v328 = v200
	goto L24
L35:
	;
	goto L36
L36:
	;
	v212 = v203
	v214 = v200
	goto L37
L37:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227+v212<<(uint(int32(2))%32))))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+10)))
	if v232 != 0 {
		v306 = v214
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v328 = v306
	goto L24
L39:
	;
	v309 = v212 + int32(1)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v309 < v310 {
		v212 = v309
		v214 = v306
		goto L37
	} else {
		goto L67
	}
L40:
	;
	if v30 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	if v283 != 0 {
		v306 = v214
		goto L39
	} else {
		goto L58
	}
L42:
	;
	goto L41
L43:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v238 <= int32(0) {
		v283 = int32(0)
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v283 = int32(0)
	goto L42
L46:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v231)+60))
	v242 = int32(0)
	if v242 < v238 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v245 = v238
	goto L49
L48:
	;
	v245 = v242
	goto L49
L49:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v249 = int32(0)
	goto L50
L50:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v246+v249<<(uint(int32(2))%32))))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+12)))
	if v259 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L45
L52:
	;
	v270 = v249 + int32(1)
	if v270 != v245 {
		v249 = v270
		goto L50
	} else {
		goto L57
	}
L53:
	;
	v260 = int32(1)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v231 == v261 {
		v283 = v260
		goto L42
	} else {
		goto L54
	}
L54:
	;
	if v241 == int32(0) {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v261)+60))
	if v265 == v241 {
		v283 = v260
		goto L42
	} else {
		goto L56
	}
L56:
	;
	goto L52
L57:
	;
	goto L51
L58:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v288 = F_contain_mutable_functions(m, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	if v288 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v292
	v298 = F_list_make1_impl(m, int32(1), v23+int32(24))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L13
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v304 = F_lappend(m, v214, v231)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L13
	} else {
		goto L66
	}
L63:
	;
	v301 = F_predicate_implied_by(m, v298, v145, int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L13
	} else {
		goto L64
	}
L64:
	;
	if v301 != 0 {
		v306 = v214
		goto L39
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	v306 = v304
	goto L39
L67:
	;
	goto L38
L68:
	;
	v319 = F_lappend(m, v171, v317)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L13
	} else {
		goto L69
	}
L69:
	;
	v163 = v163 + int32(1)
	v171 = v319
	goto L19
L70:
	;
	v344 = F_extract_actual_clauses(m, v341, int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v346 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v347 = F_replace_nestloop_params_mutator(m, v145, l0)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L13
	} else {
		goto L75
	}
L73:
	;
	v353 = v344
	v354 = v145
	v355 = v29
	goto L74
L74:
	;
	if v355 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L75:
	;
	v349 = F_replace_nestloop_params_mutator(m, v344, l0)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L13
	} else {
		goto L76
	}
L76:
	;
	v351 = F_replace_nestloop_params_mutator(m, v29, l0)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	v353 = v349
	v354 = v347
	v355 = v351
	goto L74
L78:
	;
	if l4 != 0 {
		goto L100
	} else {
		goto L101
	}
L79:
	;
	v441 = int32(0)
	goto L78
L80:
	;
	goto L81
L81:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v360 = int32(0)
	v367 = v360
	v370 = v360
	goto L82
L82:
	;
	v382 = int32(0)
	if v359 == v382 {
		v392 = v382
		goto L84
	} else {
		goto L85
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L13
	} else {
		goto L96
	}
L84:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v355)+4))
	if v393 <= v367 {
		v441 = v370
		goto L78
	} else {
		goto L87
	}
L85:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	if v386 <= v367 {
		v392 = int32(0)
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	v392 = v388 + v367<<(uint(int32(2))%32)
	goto L84
L87:
	;
	if v392 == int32(0) {
		v441 = v370
		goto L78
	} else {
		goto L88
	}
L88:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v355)+12))
	v400 = v397 + v367<<(uint(int32(2))%32)
	if v400 == int32(0) {
		v441 = v370
		goto L78
	} else {
		goto L89
	}
L89:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	v405 = F_exprType(m, v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L13
	} else {
		goto L90
	}
L90:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v403)+8))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v403)+12))
	v409 = F_get_opfamily_member_for_cmptype(m, v407, v405, v405, v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L13
	} else {
		goto L91
	}
L91:
	;
	if v409 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v413 = F_lappend_oid(m, v370, v409)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L13
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	goto L83
L95:
	;
	v367 = v367 + int32(1)
	v370 = v413
	goto L82
L96:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v403)+12))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v403)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v420
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v419
	F_errmsg_internal(m, int32(_a_F_create_indexscan_plan_0), v23)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L13
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_create_indexscan_plan_1), int32(3138), int32(_a_F_create_indexscan_plan_2))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v576)+4)) = v591
	v593 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v576)+8)) = v593
	v595 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v576)+16)) = v595
	v597 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v576)+24)) = v597
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v599)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v576)+32)) = v600
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v576)+36)) = uint8(v602)
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v576)+37)) = uint8(v604)
	m.G0 = v23 + int32(32)
	return v576
L100:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	if v453 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v555 = F_palloc0(m, int32(112))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L13
	} else {
		goto L113
	}
L103:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	if int32(0) < v454 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v535 = int32(0)
	goto L105
L105:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v538 = F_palloc0(m, int32(104))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L13
	} else {
		goto L112
	}
L106:
	;
	v463 = int32(0)
	goto L109
L107:
	;
	goto L108
L108:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	v535 = v513
	goto L105
L109:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v453)+12))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v478+v463<<(uint(int32(2))%32))))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483+v463))))
	v486 = int32(1)
	v487 = v485 ^ v486
	*(*uint8)(unsafe.Add(mBase, uint32(v482)+26)) = uint8(v487)
	v490 = v463 + v486
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	if v490 < v491 {
		v463 = v490
		goto L109
	} else {
		goto L111
	}
L110:
	;
	goto L108
L111:
	;
	goto L110
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538)+100)) = v536
	*(*int32)(unsafe.Add(mBase, uint32(v538)+96)) = v535
	*(*int32)(unsafe.Add(mBase, uint32(v538)+92)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v538)+88)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v538)+84)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v538)+80)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v538)+72)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v538)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v538)+48)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v538)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = int32(342)
	v576 = v538
	goto L99
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v555)+104)) = v553
	*(*int32)(unsafe.Add(mBase, uint32(v555)+100)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v555)+96)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v555)+92)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v555)+88)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v555)+84)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v555)+80)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v555)+72)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v555)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v555)+48)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v555)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v555))) = int32(341)
	v576 = v555
	goto L99
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
	var v71 float64
	_ = v71
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
	var v105 float64
	_ = v105
	var v113 float64
	_ = v113
	var v115 float64
	_ = v115
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
		*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(1602022801725)
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
				v71 = float64(1e+100)
				if base.F64_gt(v69, v71) != 0 {
					v83 = v71
				} else {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v69)&int64(9223372036854775807)) {
						v83 = v71
					} else {
						v79 = float64(1)
						if base.F64_le(v69, v79) != 0 {
							v83 = v79
						} else {
							v83 = base.F64_nearest(v69)
						}
					}
				}
				v84 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
				v85 = v83
				v86 = v84
			}
			if base.F64_lt(v86, v85) != 0 {
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
			v105 = v103
		} else {
			v105 = v62
		}
		if l6 != int64(0) {
			if int64(0) < l6 {
				v129 = v105
				v130 = base.F64_convert_i64_u(l6)
			} else {
				v113 = base.F64_mul(v62, float64(0.1))
				v115 = float64(1e+100)
				if base.F64_gt(v113, v115) != 0 {
					v127 = v115
				} else {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v113)&int64(9223372036854775807)) {
						v127 = v115
					} else {
						v123 = float64(1)
						if base.F64_le(v113, v123) != 0 {
							v127 = v123
						} else {
							v127 = base.F64_nearest(v113)
						}
					}
				}
				v128 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
				v129 = v128
				v130 = v127
			}
			if base.F64_lt(v129, v130) != 0 {
				v132 = v129
			} else {
				v132 = v130
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
func F_create_setop_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 float64, l7 float64) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v76 int32
	_ = v76
	var v79 float64
	_ = v79
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v95 int32
	_ = v95
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v124 float64
	_ = v124
	var v126 int32
	_ = v126
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v134 float64
	_ = v134
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	v15 = F_palloc0(m, int32(104))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(1593432867129)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v23 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)) = uint8(v23)
		*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v22
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
		if v29 != int32(1) {
			v37 = v23
		} else {
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
			if v33 != int32(1) {
				v37 = int32(0)
			} else {
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
				v37 = v36
			}
		}
		v39 = v37 & int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v15)+21)) = uint8(v39)
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v41 + v42
		if l4 == int32(0) {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
			v48 = v47
		} else {
			v48 = int32(0)
		}
		*(*float64)(unsafe.Add(mBase, uint32(v15)+96)) = l6
		*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v48
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		v58 = v56 + v57
		*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v58
		if l4 == int32(0) {
			v62 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
			v63 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
			*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = base.F64_add(v62, v63)
			v66 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
			v67 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
			v70 = *(*float64)(unsafe.Add(mBase, _c_F_create_setop_path[0]))
			v71 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
			v72 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
			if l5 != 0 {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
				v79 = base.F64_convert_i32_s(v76)
			} else {
				v79 = float64(0)
			}
			*(*float64)(unsafe.Add(mBase, uint32(v15)+56)) = base.F64_add(base.F64_mul(v70, l7), base.F64_add(base.F64_mul(base.F64_mul(v70, base.F64_add(v71, v72)), v79), base.F64_add(v66, v67)))
			*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = l7
			return v15
		} else {
			v86 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
			v87 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
			v90 = *(*float64)(unsafe.Add(mBase, _c_F_create_setop_path[0]))
			v91 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
			v92 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
			if l5 != 0 {
				v95 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
				v98 = base.F64_convert_i32_s(v95)
			} else {
				v98 = float64(0)
			}
			v100 = base.F64_add(base.F64_mul(base.F64_mul(v90, base.F64_add(v91, v92)), v98), base.F64_add(v86, v87))
			*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = v100
			*(*float64)(unsafe.Add(mBase, uint32(v15)+56)) = base.F64_add(base.F64_mul(v90, l7), v100)
			v106 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_setop_path[1])))
			if v106 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v58 + int32(1)
			} else {
			}
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+32))
			v124 = *(*float64)(unsafe.Add(mBase, _c_F_create_setop_path[2]))
			v126 = *(*int32)(unsafe.Add(mBase, _c_F_create_setop_path[3]))
			v130 = base.F64_mul(base.F64_mul(v124, base.F64_convert_i32_s(v126)), float64(1024))
			v131 = float64(4.294967295e+09)
			if base.F64_lt(v130, v131) != 0 {
				v134 = v130
			} else {
				v134 = v131
			}
			if base.F64_lt(v134, float64(4.294967296e+09))&base.F64_ge(v134, float64(0)) != 0 {
				v140 = base.I32_trunc_f64_u(v134)
				v142 = v140
			} else {
				v142 = int32(0)
			}
			if base.F64_gt(base.F64_mul(l6, base.F64_convert_i32_u((v113+int32(7))&int32(-8)+int32(16))), base.F64_convert_i32_u(v142)) != 0 {
				v145 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v145 + int32(1)
			} else {
			}
			*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = l7
			return v15
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	v10 = F_palloc0(m, int32(80))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(1490353651999)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
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
				v28 = v27
			} else {
				v28 = int32(0)
			}
			v30 = v28 & int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v30)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v32
			F_cost_subqueryscan(m, v10, l0, l1, v19, l3)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
