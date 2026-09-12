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
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v7 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	v10 = *(*int32)(unsafe.Add(mBase, _consts[228]))
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
				v43 = *(*int32)(unsafe.Add(mBase, _consts[228]))
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
									v56 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
									*(*int64)(unsafe.Add(mBase, uint32(v45))) = v56
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v58
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
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int64
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+48)))
	v21 = v19 & int32(240)
	switch v21 - int32(16) {
	case 0:
		goto L2
	default:
		goto L3
	case 16:
		goto L4
	}
L1:
	;
	m.G0 = v16 + int32(112)
	return
L2:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v218
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v217)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v220
	v225 = F_smgropen(m, v16+int32(16), int32(-1))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L5
	} else {
		goto L42
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L5
	} else {
		goto L39
	}
L4:
	;
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v26
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = v28
	v34 = F_smgropen(m, v16-int32(-64), int32(-1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	F_smgrcreate(m, v34, int32(0), int32(1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	F_XLogFlush(m, v24)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v43 = v25 + int32(4)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+16)))
	if v44&int32(1) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v97
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
	v103 = F_CreateFakeRelcacheEntry(m, v16+int32(32))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L15
	}
L10:
	;
	v90 = int32(0)
	v91 = v16 + int32(88)
	v92 = v16 + int32(100)
	v95 = v16 + int32(76)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v62 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = v62
	v65 = F_smgrnblocks(m, v34, v62)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v65
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v70
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v72
	v74 = m.G0
	v76 = v74 - int32(16)
	m.G0 = v76
	v79 = v16 + int32(48)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v80
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
	*(*int64)(unsafe.Add(mBase, uint32(v76))) = v82
	F_forget_invalid_pages(m, v76, int32(0), v68)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	m.G0 = v76 + int32(16)
	v90 = int32(1)
	v91 = v16 + int32(92)
	v92 = v16 + int32(104)
	v95 = v16 + int32(80)
	goto L9
L15:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+16)))
	if v105&int32(4) == int32(0) {
		v130 = v90
		v132 = v2
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+16)))
	if v133&int32(2) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v111 = F_smgrexists(m, v34, int32(1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	if v111 == int32(0) {
		v130 = v90
		v132 = v2
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v116 = F_FreeSpaceMapPrepareTruncateRel(m, v103, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v116
	if v116 == int32(-1) {
		v130 = v90
		v132 = v2
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v121 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v121
	v125 = F_smgrnblocks(m, v34, v121)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v125
	v130 = v90 + int32(1)
	v132 = v121
	goto L16
L23:
	;
	if v132 != 0 {
		goto L34
	} else {
		goto L35
	}
L24:
	;
	v175 = int32(4484100)
	v177 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v177 + int32(1)
	F_smgrtruncate(m, v34, v16+int32(100), v172, v16+int32(76), v16+int32(88))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L5
	} else {
		goto L33
	}
L25:
	;
	if v130 == int32(0) {
		goto L23
	} else {
		goto L32
	}
L26:
	;
	v139 = F_smgrexists(m, v34, int32(2))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	if v139 == int32(0) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v144 = v130 << (uint(int32(2)) % 32)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v149 = F_visibilitymap_prepare_truncate(m, v103, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144+(v16+int32(88))))) = v149
	if v149 == int32(-1) {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v157 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(100)+v144))) = v157
	v163 = F_smgrnblocks(m, v34, v157)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(76)+v144))) = v163
	v172 = v130 + int32(1)
	goto L24
L32:
	;
	v172 = v130
	goto L24
L33:
	;
	v189 = int32(4484100)
	v191 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v191 - int32(1)
	goto L23
L34:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	F_FreeSpaceMapVacuumRange(m, v103, v198, int32(-1))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_pfree(m, v103)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L5
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	goto L1
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v21
	F_errmsg_internal(m, int32(52352), v16)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(498687), int32(1094), int32(242342))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
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
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	F_smgrcreate(m, v225, v227, int32(1))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	goto L1
}
