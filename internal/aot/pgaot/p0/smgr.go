package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_smgrGetPendingDeletes(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_smgrGetPendingDeletes[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_smgrGetPendingDeletes[1]))
	if v10 != 0 {
		v13 = v10
		v14 = int32(0)
		for {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
			if v16 < v8 {
				v24 = v14
			} else {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+16)))
				if v18 != l0 {
					v24 = v14
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
					v24 = v14 + base.B2i32(v20 == int32(-1))
				}
			}
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
			if v25 != 0 {
				v13 = v25
				v14 = v24
				continue
			} else {
				break
			}
			break
		}
		if v24 != 0 {
			v37 = F_palloc(m, v24*int32(12))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v37
				v43 = *(*int32)(unsafe.Add(mBase, _c_F_smgrGetPendingDeletes[1]))
				if v43 != 0 {
					v45 = v37
					v46 = v43
					for {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
						if v49 < v8 {
							v62 = v45
						} else {
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+16)))
							if v51 != l0 {
								v62 = v45
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
								if v53 != int32(-1) {
									v62 = v45
								} else {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v56
									v58 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
									*(*int64)(unsafe.Add(mBase, uint32(v45))) = v58
									v62 = v45 + int32(12)
								}
							}
						}
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
						if v63 != 0 {
							v45 = v62
							v46 = v63
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				return v24
			}
		} else {
			v31 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31
			return v31
		}
	} else {
		v31 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31
		return v31
	}
}
func F_smgr_bulk_write(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	v4 = l3
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6 + int32(1)
	v12 = l0 + v6*int32(12)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l2
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v16 == int32(32) {
		F_smgr_bulk_flush(m, l0)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_smgr_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int64
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int64
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+48)))
	v21 = v19 & int32(240)
	v23 = v21 - int32(16)
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(112)
	return
L2:
	;
	if v23 == int32(16) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v220
	v222 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v222
	v227 = F_smgropen(m, v16+int32(16), int32(-1))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L8
	} else {
		goto L45
	}
L5:
	;
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v27)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v30
	v36 = F_smgropen(m, v16-int32(-64), int32(-1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L8
	} else {
		goto L42
	}
L8:
	;
	return
L9:
	;
	F_smgrcreate(m, v36, int32(0), int32(1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_XLogFlush(m, v26)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v45 = v27 + int32(4)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+16)))
	if v46&int32(1) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v99
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v101
	v105 = F_CreateFakeRelcacheEntry(m, v16+int32(32))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L18
	}
L13:
	;
	v92 = int32(0)
	v94 = v16 + int32(100)
	v95 = v16 + int32(88)
	v96 = v16 + int32(76)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v64 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = v64
	v67 = F_smgrnblocks(m, v36, v64)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v70
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v74
	v76 = m.G0
	v78 = v76 - int32(16)
	m.G0 = v78
	v81 = v16 + int32(48)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v82
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v84
	F_forget_invalid_pages(m, v78, int32(0), v70)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	m.G0 = v78 + int32(16)
	v92 = int32(1)
	v94 = v16 + int32(104)
	v95 = v16 + int32(92)
	v96 = v16 + int32(80)
	goto L12
L18:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+16)))
	if v107&int32(4) == int32(0) {
		v132 = v92
		v134 = v2
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+16)))
	if v135&int32(2) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v113 = F_smgrexists(m, v36, int32(1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	if v113 == int32(0) {
		v132 = v92
		v134 = v2
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v118 = F_FreeSpaceMapPrepareTruncateRel(m, v105, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v118
	if v118 == int32(-1) {
		v132 = v92
		v134 = v2
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v123 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v123
	v127 = F_smgrnblocks(m, v36, v123)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v127
	v132 = v92 + int32(1)
	v134 = v123
	goto L19
L26:
	;
	if v134 != 0 {
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v177 = int32(_a_F_smgr_redo_0)
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_redo[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgr_redo[0])) = v179 + int32(1)
	F_smgrtruncate(m, v36, v16+int32(100), v174, v16+int32(76), v16+int32(88))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L8
	} else {
		goto L36
	}
L28:
	;
	if v132 == int32(0) {
		goto L26
	} else {
		goto L35
	}
L29:
	;
	v141 = F_smgrexists(m, v36, int32(2))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	if v141 == int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v146 = v132 << (uint(int32(2)) % 32)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v151 = F_visibilitymap_prepare_truncate(m, v105, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146+(v16+int32(88))))) = v151
	if v151 == int32(-1) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v159 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(100)+v146))) = v159
	v165 = F_smgrnblocks(m, v36, v159)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(76)+v146))) = v165
	v174 = v132 + int32(1)
	goto L27
L35:
	;
	v174 = v132
	goto L27
L36:
	;
	v191 = int32(_a_F_smgr_redo_0)
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_redo[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgr_redo[0])) = v193 - int32(1)
	goto L26
L37:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	F_FreeSpaceMapVacuumRange(m, v105, v200, int32(-1))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L8
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_pfree(m, v105)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L8
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	goto L1
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v21
	F_errmsg_internal(m, int32(_a_F_smgr_redo_1), v16)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_smgr_redo_2), int32(1094), int32(_a_F_smgr_redo_3))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
	F_smgrcreate(m, v227, v229, int32(1))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	goto L1
}
