package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LogCheckpointStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v10 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if v10 != 0 {
			if l0&int32(16) != 0 {
				v16 = int32(_a_F_LogCheckpointStart_0)
			} else {
				v16 = int32(_a_F_LogCheckpointStart_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v16
			if l0&int32(256) != 0 {
				v22 = int32(_a_F_LogCheckpointStart_2)
			} else {
				v22 = int32(_a_F_LogCheckpointStart_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v22
			if l0&int32(128) != 0 {
				v28 = int32(_a_F_LogCheckpointStart_3)
			} else {
				v28 = int32(_a_F_LogCheckpointStart_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v28
			if l0&int32(32) != 0 {
				v34 = int32(_a_F_LogCheckpointStart_4)
			} else {
				v34 = int32(_a_F_LogCheckpointStart_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v34
			if l0&int32(8) != 0 {
				v40 = int32(_a_F_LogCheckpointStart_5)
			} else {
				v40 = int32(_a_F_LogCheckpointStart_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v40
			if l0&int32(4) != 0 {
				v46 = int32(_a_F_LogCheckpointStart_6)
			} else {
				v46 = int32(_a_F_LogCheckpointStart_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v46
			if l0&int32(2) != 0 {
				v52 = int32(_a_F_LogCheckpointStart_7)
			} else {
				v52 = int32(_a_F_LogCheckpointStart_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v52
			if l0&int32(1) != 0 {
				v58 = int32(_a_F_LogCheckpointStart_8)
			} else {
				v58 = int32(_a_F_LogCheckpointStart_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v58
			if l1 != 0 {
				v62 = int32(_a_F_LogCheckpointStart_9)
			} else {
				v62 = int32(_a_F_LogCheckpointStart_10)
			}
			F_errmsg(m, v62, v6)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				if l1 != 0 {
					v68 = int32(_a_F_LogCheckpointStart_11)
				} else {
					v68 = int32(_a_F_LogCheckpointStart_12)
				}
				F_errfinish(m, int32(_a_F_LogCheckpointStart_13), v68, int32(_a_F_LogCheckpointStart_14))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					m.G0 = v6 + int32(32)
					return
				}
			}
		} else {
			m.G0 = v6 + int32(32)
			return
		}
	}
}
func F_ProcessLogMemoryContextInterrupt(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	v1 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(176)
	m.G0 = v9
	v12 = v1
	v13 = int32(-1)
	v16 = v1
	v17 = v1
	goto L1
L1:
	;
	goto L3
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if v13 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	goto L2
L5:
	;
	v144 = int32(m.ExcTag)
	v145 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v144 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessLogMemoryContextInterrupt[0])) = v42
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessLogMemoryContextInterrupt[1])) = v41
	v137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessLogMemoryContextInterrupt[2])) = uint8(v137)
	F_pg_re_throw(m)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L34
	}
L7:
	;
	m.G0 = v9 + int32(176)
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessLogMemoryContextInterrupt[3])) = int32(0)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessLogMemoryContextInterrupt[2])))
	if v24 != 0 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v40 = v12
	v41 = v16
	v42 = v17
	goto L10
L10:
	;
	if v40 != 0 {
		goto L6
	} else {
		goto L16
	}
L11:
	;
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessLogMemoryContextInterrupt[2])) = uint8(v27)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessLogMemoryContextInterrupt[1]))
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessLogMemoryContextInterrupt[0]))
	goto L12
L12:
	;
	v34 = v9 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v9 + int32(12)
	goto L15
L13:
	;
	v40 = int32(0)
	v41 = v30
	v42 = v32
	goto L10
L15:
	;
	goto L13
L16:
	;
	v44 = int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessLogMemoryContextInterrupt[1])) = v9 + v44
	v49 = F_errstart(m, v44, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v49 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_errhidestmt(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessLogMemoryContextInterrupt[4]))
	v68 = m.G0
	v70 = v68 - int32(80)
	m.G0 = v70
	v72 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v70)+72)) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v70)+64)) = v72
	v77 = int32(100)
	F_MemoryContextStatsInternal(m, v67, int32(1), v77, v77, v70-int32(-64), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L25
	}
L21:
	;
	F_errhidecontext(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessLogMemoryContextInterrupt[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v56
	F_errmsg(m, int32(_a_F_ProcessLogMemoryContextInterrupt_0), v9)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_ProcessLogMemoryContextInterrupt_1), int32(1319), int32(_a_F_ProcessLogMemoryContextInterrupt_2))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v86 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	if v86 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_errhidestmt(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	m.G0 = v70 + int32(80)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessLogMemoryContextInterrupt[1])) = v41
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessLogMemoryContextInterrupt[0])) = v42
	v122 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessLogMemoryContextInterrupt[2])) = uint8(v122)
	goto L7
L30:
	;
	F_errhidecontext(m)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v70)+72))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v70)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+48)) = v92 - v93
	*(*int32)(unsafe.Add(mBase, uint32(v70)+32)) = v92
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v70)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+36)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v70)+40)) = v93
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v70)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+44)) = v100
	F_errmsg_internal(m, int32(_a_F_ProcessLogMemoryContextInterrupt_3), v70+int32(32))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_ProcessLogMemoryContextInterrupt_1), int32(867), int32(_a_F_ProcessLogMemoryContextInterrupt_4))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	goto L4
L35:
	;
	v149 = int32(v145)
	m.G0 = v9
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	if v9+int32(12) == v155 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	m.ExcPending = 1
	goto L44
L37:
	;
	if v159 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v159 = v157
	goto L40
L39:
	;
	v159 = int32(0)
	goto L40
L40:
	;
	goto L37
L41:
	;
	F___wasm_longjmp(m, v152, v151)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v12 = v151
	v13 = v159
	v16 = v41
	v17 = v42
	goto L1
L44:
	;
	return
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_log_destination(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = F_pstrdup(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v288
L2:
	;
	return int32(0)
L3:
	;
	v21 = F_SplitIdentifierString(m, v14, int32(44), v11+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v21 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_destination[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_destination[1])) = v27
	v32 = F_format_elog_string(m, int32(_a_F_check_log_destination_0), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v40 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v41 == v40 {
		v247 = v40
		goto L12
	} else {
		goto L13
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_destination[2])) = v32
	F_pfree(m, v14)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_list_free(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v288 = v4
	goto L1
L11:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_destination[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_destination[1])) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v61
	v273 = F_format_elog_string(m, int32(_a_F_check_log_destination_1), v11)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L2
	} else {
		goto L79
	}
L12:
	;
	F_pfree(m, v14)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L2
	} else {
		goto L75
	}
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v44 <= int32(0) {
		v247 = v40
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v48 = int32(0)
	v50 = v40
	goto L15
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v48<<(uint(int32(2))%32))))
	v65 = v61
	v66 = int32(_a_F_check_log_destination_2)
	goto L19
L16:
	;
	v247 = v240
	goto L12
L17:
	;
	v240 = v239 | v50
	v242 = v48 + int32(1)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v242 < v243 {
		v48 = v242
		v50 = v240
		goto L15
	} else {
		goto L74
	}
L18:
	;
	if v103 == int32(0) {
		v239 = int32(1)
		goto L17
	} else {
		goto L31
	}
L19:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v69 == v70 {
		v92 = v69
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v103 = int32(0)
	goto L18
L21:
	;
	v94 = int32(1)
	if v92 != 0 {
		v65 = v65 + v94
		v66 = v66 + v94
		goto L19
	} else {
		goto L30
	}
L22:
	;
	if base.Ui32((v69-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v80 = v69 | int32(32)
	goto L25
L24:
	;
	v80 = v69
	goto L25
L25:
	;
	if base.Ui32((v70-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v89 = v70 | int32(32)
	goto L28
L27:
	;
	v89 = v70
	goto L28
L28:
	;
	if v80 == v89 {
		v92 = v80
		goto L21
	} else {
		goto L29
	}
L29:
	;
	v103 = v80 - v89
	goto L18
L30:
	;
	goto L20
L31:
	;
	v110 = v61
	v111 = int32(_a_F_check_log_destination_3)
	goto L33
L32:
	;
	if v148 == int32(0) {
		v239 = int32(8)
		goto L17
	} else {
		goto L45
	}
L33:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v114 == v115 {
		v137 = v114
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v148 = int32(0)
	goto L32
L35:
	;
	v139 = int32(1)
	if v137 != 0 {
		v110 = v110 + v139
		v111 = v111 + v139
		goto L33
	} else {
		goto L44
	}
L36:
	;
	if base.Ui32((v114-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v125 = v114 | int32(32)
	goto L39
L38:
	;
	v125 = v114
	goto L39
L39:
	;
	if base.Ui32((v115-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v134 = v115 | int32(32)
	goto L42
L41:
	;
	v134 = v115
	goto L42
L42:
	;
	if v125 == v134 {
		v137 = v125
		goto L35
	} else {
		goto L43
	}
L43:
	;
	v148 = v125 - v134
	goto L32
L44:
	;
	goto L34
L45:
	;
	v155 = v61
	v156 = int32(_a_F_check_log_destination_4)
	goto L47
L46:
	;
	if v193 == int32(0) {
		v239 = int32(16)
		goto L17
	} else {
		goto L59
	}
L47:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v159 == v160 {
		v182 = v159
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v193 = int32(0)
	goto L46
L49:
	;
	v184 = int32(1)
	if v182 != 0 {
		v155 = v155 + v184
		v156 = v156 + v184
		goto L47
	} else {
		goto L58
	}
L50:
	;
	if base.Ui32((v159-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v170 = v159 | int32(32)
	goto L53
L52:
	;
	v170 = v159
	goto L53
L53:
	;
	if base.Ui32((v160-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v179 = v160 | int32(32)
	goto L56
L55:
	;
	v179 = v160
	goto L56
L56:
	;
	if v170 == v179 {
		v182 = v170
		goto L49
	} else {
		goto L57
	}
L57:
	;
	v193 = v170 - v179
	goto L46
L58:
	;
	goto L48
L59:
	;
	v199 = v61
	v200 = int32(_a_F_check_log_destination_5)
	goto L61
L60:
	;
	if v237 != 0 {
		goto L11
	} else {
		goto L73
	}
L61:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v203 == v204 {
		v226 = v203
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v237 = int32(0)
	goto L60
L63:
	;
	v228 = int32(1)
	if v226 != 0 {
		v199 = v199 + v228
		v200 = v200 + v228
		goto L61
	} else {
		goto L72
	}
L64:
	;
	if base.Ui32((v203-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v214 = v203 | int32(32)
	goto L67
L66:
	;
	v214 = v203
	goto L67
L67:
	;
	if base.Ui32((v204-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v223 = v204 | int32(32)
	goto L70
L69:
	;
	v223 = v204
	goto L70
L70:
	;
	if v214 == v223 {
		v226 = v214
		goto L63
	} else {
		goto L71
	}
L71:
	;
	v237 = v214 - v223
	goto L60
L72:
	;
	goto L62
L73:
	;
	v239 = int32(2)
	goto L17
L74:
	;
	goto L16
L75:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_list_free(m, v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v259 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	if v259 == int32(0) {
		v288 = v4
		goto L1
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v247
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v259
	v288 = int32(1)
	goto L1
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_destination[2])) = v273
	F_pfree(m, v14)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_list_free(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v288 = v4
	goto L1
}
func F_check_log_duration(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v57 int64
	_ = v57
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v114 float64
	_ = v114
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v146 float64
	_ = v146
	var v148 float64
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_log_duration[0])))
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v203
L2:
	;
	v32 = *(*int64)(unsafe.Add(mBase, _c_F_check_log_duration[1]))
	v36 = m.G0
	v37 = int32(16)
	v38 = v36 - v37
	m.G0 = v38
	F_gettimeofday(m, v38)
	mBase = m.M
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
	v42 = int64(*(*int32)(unsafe.Add(mBase, uint32(v38)+8)))
	m.G0 = v38 + v37
	goto L7
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_duration[3]))
	if int32(0) <= v18 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_duration[2]))
	if int32(0) <= v22 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_log_duration[5])))
	if v26&int32(1) == int32(0) {
		v203 = int32(0)
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L2
L7:
	;
	v57 = v42 + v41*int64(1000000) - int64(946684800000000) - v32
	if v57 <= int64(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v75 = base.I32_div_s(v73, int32(1000))
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_duration[2]))
	if v78 == int32(0) {
		v94 = int32(1)
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v69 = int32(0)
	v70 = int32(0)
	goto L11
L10:
	;
	v61 = int64(1000000)
	v62 = base.I64_div_u_s(v57, v61)
	v69 = base.I32_wrap_i64(v62)
	v70 = base.I32_wrap_i64(v57 - v62*v61)
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(12)))) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(8)))) = v70
	goto L8
L12:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_duration[3]))
	if v96 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v81 = int32(0)
	if v78 <= v81 {
		v94 = v81
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v87 = base.I32_div_u_s(v78, int32(1000))
	if v87 < v85 {
		v94 = int32(1)
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v94 = base.B2i32(v78 <= v85*int32(1000)+v75)
	goto L12
L16:
	;
	v175 = int32(1000)
	v176 = base.I32_rem_s(v172, v175)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v178*v175 + v75
	v185 = F_pg_snprintf(m, l0, int32(32), int32(_a_F_check_log_duration_0), v13)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L30
	} else {
		goto L31
	}
L17:
	;
	v154 = v151 | v94
	if v154&int32(1) != 0 {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	v114 = *(*float64)(unsafe.Add(mBase, _c_F_check_log_duration[4]))
	if base.F64_eq(v114, float64(0)) != 0 {
		v151 = int32(0)
		goto L17
	} else {
		goto L23
	}
L19:
	;
	v99 = int32(0)
	if v96 <= v99 {
		v151 = v99
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v104 = base.I32_div_u_s(v96, int32(1000))
	if v104 < v102 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	if v102*int32(1000)+v75 < v96 {
		v151 = v99
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	if base.F64_eq(v114, float64(1)) != 0 {
		v170 = int32(1)
		v172 = v73
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v120 = int32(_a_F_check_log_duration_1)
	v123 = *(*int64)(unsafe.Add(mBase, _c_F_check_log_duration[6]))
	v124 = *(*int64)(unsafe.Add(mBase, _c_F_check_log_duration[7]))
	v125 = v123 ^ v124
	*(*int64)(unsafe.Add(mBase, _c_F_check_log_duration[7])) = base.I64_rotl(v125, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_check_log_duration[6])) = v125<<(uint(int64(16))%64) ^ base.I64_rotl(v123, int64(24)) ^ v125
	v146 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v123*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L25
L25:
	;
	v148 = *(*float64)(unsafe.Add(mBase, _c_F_check_log_duration[4]))
	v151 = base.F64_le(v146, v148)
	goto L17
L26:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v170 = v154
	v172 = v169
	goto L16
L27:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_log_duration[0])))
	if v159&int32(1) != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_log_duration[5])))
	if v163&int32(1) == int32(0) {
		v203 = int32(0)
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	return int32(0)
L31:
	;
	v189 = int32(1)
	v190 = int32(2)
	v194 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_log_duration[5])))
	if v194&v189 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v197 = v190
	goto L34
L33:
	;
	v197 = v189
	goto L34
L34:
	;
	if v170&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v200 = v190
	goto L37
L36:
	;
	v200 = v197
	goto L37
L37:
	;
	if l1 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v201 = v189
	goto L40
L39:
	;
	v201 = v200
	goto L40
L40:
	;
	v203 = v201
	goto L1
}
func F_check_log_of_query(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_of_query[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v7-int32(15)) <= base.Ui32(int32(1)) {
		if v6 < int32(22) {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
			if v21 != 0 {
				v26 = v2
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_of_query[1]))
				v26 = base.B2i32(v23 != int32(0))
			}
		} else {
			v26 = v2
		}
	} else {
		if v7 == int32(20) {
			v26 = v2
		} else {
			if v6 == int32(15) {
				if int32(21) < v7 {
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
					if v21 != 0 {
						v26 = v2
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_of_query[1]))
						v26 = base.B2i32(v23 != int32(0))
					}
				} else {
					v26 = v2
				}
			} else {
				if v7 < v6 {
					v26 = v2
				} else {
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
					if v21 != 0 {
						v26 = v2
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_of_query[1]))
						v26 = base.B2i32(v23 != int32(0))
					}
				}
			}
		}
	}
	return v26
}
func F_show_log_timezone(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_show_log_timezone[0]))
	if v3 != 0 {
		v5 = v3
	} else {
		v5 = int32(_a_F_show_log_timezone_0)
	}
	return v5
}
