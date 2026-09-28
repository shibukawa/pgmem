package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PGSharedMemoryAttach(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(192)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	v12 = int32(2)
	v16 = F_pgmem_shmctl(m, l0, v12, v8+int32(104))
	mBase = m.M
	if v16 < v4 {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_PGSharedMemoryAttach[0]))
		switch v20 - int32(2) {
		case 0:
			v61 = int32(3)
		default:
			v61 = int32(0)
		case 22, 26:
			v61 = v12
		}
	} else {
		v25 = int32(0)
		v27 = *(*int32)(unsafe.Add(mBase, _c_F_PGSharedMemoryAttach[1]))
		v32 = F___fstatat(m, int32(-100), v27, v8+int32(8), v25)
		mBase = m.M
		if v32 < int32(0) {
			v61 = v25
		} else {
			v35 = F_pgmem_shmat(m, l0, l1)
			mBase = m.M
			if v35 == int32(-1) {
				v38 = int32(2)
				v40 = *(*int32)(unsafe.Add(mBase, _c_F_PGSharedMemoryAttach[0]))
				switch v40 - v38 {
				case 0:
					v61 = int32(3)
				default:
					v61 = int32(0)
				case 22, 26:
					v61 = v38
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v35
				v46 = int32(3)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
				if v47 != int32(679834894) {
					v61 = v46
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
					if v50 != v51 {
						v61 = v46
					} else {
						v53 = *(*int64)(unsafe.Add(mBase, uint32(v35)+24))
						v54 = *(*int64)(unsafe.Add(mBase, uint32(v8)+96))
						if v53 != v54 {
							v61 = v46
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v8)+176))
							if v58 != 0 {
								v59 = int32(1)
							} else {
								v59 = int32(4)
							}
							v61 = v59
						}
					}
				}
			}
		}
	}
	m.G0 = v8 + int32(192)
	return v61
}
func F_PgArchShmemRequest(m *base.Module, l0 int32) {
	var v6 int32
	_ = v6
	Fn14222(m, l0, int32(_a_F_PgArchShmemRequest_0), int64(8), int32(_a_F_PgArchShmemRequest_1))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_ProcessPgArchInterrupts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPgArchInterrupts[0]))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPgArchInterrupts[1]))
	if v8 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	F_ProcessLogMemoryContextInterrupt(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPgArchInterrupts[2]))
	if v12 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L8
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L35
	}
L11:
	;
	return
L12:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPgArchInterrupts[3]))
	v17 = F_pstrdup(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessPgArchInterrupts[2])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPgArchInterrupts[3]))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v27 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPgArchInterrupts[4]))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v30 != 0 {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if base.B2i32(v33 == int32(0))|base.B2i32(v33 != v36) != 0 {
		v54 = v33
		v55 = v36
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L17
L19:
	;
	F_pfree(m, v17)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v39 = v26
	v40 = v17
	goto L22
L22:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if v44 == int32(0) {
		v54 = v44
		v55 = v43
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v54 = v44
	v55 = v43
	goto L20
L24:
	;
	v47 = int32(1)
	if v44 == v43 {
		v39 = v39 + v47
		v40 = v40 + v47
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	if v54-v55 == int32(0) {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v63 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	if v63 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_errmsg(m, int32(_a_F_ProcessPgArchInterrupts_0), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	F_errfinish(m, int32(_a_F_ProcessPgArchInterrupts_1), int32(903), int32(_a_F_ProcessPgArchInterrupts_2))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(_a_F_ProcessPgArchInterrupts_3), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v92 = F_errdetail(m, int32(_a_F_ProcessPgArchInterrupts_4), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_ProcessPgArchInterrupts_1), int32(885), int32(_a_F_ProcessPgArchInterrupts_2))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RemovePgTempRelationFiles(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v260 int32
	_ = v260
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v337 int32
	_ = v337
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v362 int32
	_ = v362
	v11 = m.G0
	v13 = v11 - int32(_a_F_RemovePgTempRelationFiles_0)
	m.G0 = v13
	v15 = F_AllocateDir(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = F_ReadDirExtended(m, v15, l0, int32(15))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = v18
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_FreeDir(m, v15)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L88
	}
L7:
	;
	v31 = v21 + int32(19)
	v32 = int32(_a_F_RemovePgTempRelationFiles_1)
	v36 = m.G0
	v38 = v36 - int32(32)
	v39 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = v39
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RemovePgTempRelationFiles[0])))
	if v47 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v116 = F_strlen(m, v31)
	mBase = m.M
	if v115 == v116 {
		goto L28
	} else {
		goto L29
	}
L10:
	;
	v115 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RemovePgTempRelationFiles[1])))
	if v51 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v55 = v31
	goto L16
L14:
	;
	goto L15
L15:
	;
	v65 = v32
	v66 = v47
	goto L19
L16:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v61 == v47 {
		v55 = v55 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v115 = v55 - v31
	goto L9
L18:
	;
	goto L17
L19:
	;
	v73 = v38 + int32(base.Ui32(v66)>>(uint(int32(3))%32))&int32(28)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v75 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v74 | v75<<(uint(v66)%32)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	if v79 != 0 {
		v65 = v65 + v75
		v66 = v79
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v82 == int32(0) {
		v105 = v31
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v115 = v105 - v31
	goto L9
L23:
	;
	v86 = v31
	v87 = v82
	goto L24
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(base.Ui32(v87)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v95)>>(uint(v87)%32))&int32(1) == int32(0) {
		v105 = v86
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v105 = v103
	goto L22
L26:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	v103 = v86 + int32(1)
	if v101 != 0 {
		v86 = v103
		v87 = v101
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l0
	v121 = v13 + int32(48)
	v126 = F_pg_snprintf(m, v121, int32(2048), int32(_a_F_RemovePgTempRelationFiles_2), v13+int32(32))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v349 = F_ReadDirExtended(m, v15, l0, int32(15))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L86
	}
L31:
	;
	v128 = F_AllocateDir(m, v121)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v131 = F_ReadDirExtended(m, v128, v121, int32(15))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v131 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v134 = v131
	goto L37
L35:
	;
	goto L36
L36:
	;
	F_FreeDir(m, v128)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L85
	}
L37:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+19)))
	if v143 != int32(116) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L36
L39:
	;
	v324 = F_ReadDirExtended(m, v128, v13+int32(48), int32(15))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L83
	}
L40:
	;
	v147 = v134 + int32(19)
	v150 = int32(1)
	goto L41
L41:
	;
	v160 = v150 + int32(1)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v147))))
	if base.Ui32((v162-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v150 = v160
		goto L41
	} else {
		goto L43
	}
L42:
	;
	if v150 == int32(1) {
		goto L39
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	if v162 != int32(95) {
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v175 = v160
	goto L46
L46:
	;
	v184 = v175 + int32(1)
	v185 = v175 + v147
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if base.Ui32((v186-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v175 = v184
		goto L46
	} else {
		goto L48
	}
L47:
	;
	if v160 == v175 {
		goto L39
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	if v186 == int32(95) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v197 = v185 + int32(1)
	v201 = int32(3)
	v204 = F_strncmp(m, int32(_a_F_RemovePgTempRelationFiles_3), v197, v201)
	mBase = m.M
	if v204 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	v241 = v175
	v242 = v186
	goto L52
L52:
	;
	if v242 == int32(46) {
		goto L68
	} else {
		goto L69
	}
L53:
	;
	if v233 <= int32(0) {
		goto L39
	} else {
		goto L67
	}
L54:
	;
	goto L53
L56:
	;
	v233 = v226
	goto L54
L57:
	;
	v226 = v201
	goto L56
L58:
	;
	goto L59
L59:
	;
	v208 = int32(2)
	v212 = F_strncmp(m, int32(_a_F_RemovePgTempRelationFiles_4), v197, v208)
	mBase = m.M
	if v212 == int32(0) {
		v226 = v208
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v215 = int32(4)
	v218 = F_strncmp(m, int32(_a_F_RemovePgTempRelationFiles_5), v197, v215)
	mBase = m.M
	if v218 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	goto L64
L62:
	;
	goto L63
L63:
	;
	v233 = int32(0)
	goto L54
L64:
	;
	v233 = v215
	goto L54
L67:
	;
	v237 = v233 + v184
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v237))))
	v241 = v237
	v242 = v239
	goto L52
L68:
	;
	v248 = int32(1)
	goto L71
L69:
	;
	v275 = v242
	goto L70
L70:
	;
	if v275 != 0 {
		goto L39
	} else {
		goto L75
	}
L71:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248+(v241+v147)))))
	if base.Ui32((v260-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v248 = v248 + int32(1)
		goto L71
	} else {
		goto L73
	}
L72:
	;
	if v248 < int32(2) {
		goto L39
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v275 = v260
	goto L70
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v13 + int32(48)
	v284 = v13 + int32(2096)
	v289 = F_pg_snprintf(m, v284, int32(2048), int32(_a_F_RemovePgTempRelationFiles_2), v13+int32(16))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v291 = F_unlink(m, v284)
	mBase = m.M
	if int32(0) <= v291 {
		goto L39
	} else {
		goto L77
	}
L77:
	;
	v296 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v296 == int32(0) {
		goto L39
	} else {
		goto L79
	}
L79:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v284
	F_errmsg(m, int32(_a_F_RemovePgTempRelationFiles_6), v13)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_RemovePgTempRelationFiles_7), int32(3496), int32(_a_F_RemovePgTempRelationFiles_8))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	goto L39
L83:
	;
	if v324 != 0 {
		v134 = v324
		goto L37
	} else {
		goto L84
	}
L84:
	;
	goto L38
L85:
	;
	goto L30
L86:
	;
	if v349 != 0 {
		v21 = v349
		goto L7
	} else {
		goto L87
	}
L87:
	;
	goto L8
L88:
	;
	m.G0 = v13 + int32(_a_F_RemovePgTempRelationFiles_0)
	return
}
func F__PG_init_isn(m *base.Module) {
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v3 = int32(0)
	F_DefineCustomBoolVariable(m, int32(_a_F__PG_init_isn_0), int32(_a_F__PG_init_isn_1), v3, int32(_a_F__PG_init_isn_2), v3, int32(6))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		F_MarkGUCPrefixReserved(m, int32(_a_F__PG_init_isn_3))
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F__PG_init_pg_prewarm(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	F_DefineCustomIntVariable(m, int32(_a_F__PG_init_pg_prewarm_0), int32(_a_F__PG_init_pg_prewarm_1), int32(_a_F__PG_init_pg_prewarm_2), int32(_a_F__PG_init_pg_prewarm_3), int32(300), int32(0), int32(_a_F__PG_init_pg_prewarm_4), int32(2), int32(536870912))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__PG_init_pg_prewarm[0])))
		if v13 != int32(1) {
			return
		} else {
			v20 = int32(1)
			F_DefineCustomBoolVariable(m, int32(_a_F__PG_init_pg_prewarm_5), int32(_a_F__PG_init_pg_prewarm_6), int32(0), int32(_a_F__PG_init_pg_prewarm_7), v20, v20)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				F_MarkGUCPrefixReserved(m, int32(_a_F__PG_init_pg_prewarm_8))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v28 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__PG_init_pg_prewarm[1])))
					if v28 != int32(1) {
						return
					} else {
						F_apw_start_leader_worker(m)
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
func F_create_pg_locale_libc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
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
	var v239 int32
	_ = v239
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v248 int64
	_ = v248
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v426 int64
	_ = v426
	var v428 int64
	_ = v428
	var v430 int64
	_ = v430
	var v442 int32
	_ = v442
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v608 int64
	_ = v608
	var v610 int64
	_ = v610
	var v612 int64
	_ = v612
	var v624 int32
	_ = v624
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l0 == int32(100) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	v674 = F_MemoryContextAllocZero(m, l1, int32(20))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L9
	} else {
		goto L173
	}
L2:
	;
	F_report_newlocale_failure(m, v50)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L9
	} else {
		goto L172
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L9
	} else {
		goto L169
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L9
	} else {
		goto L166
	}
L5:
	;
	v54 = F_text_to_cstring(m, base.I32_wrap_i64(v52))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L20
	}
L6:
	;
	v15 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[0])))
	v16 = F_SearchSysCache1(m, int32(21), v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v35 = F_SearchSysCache1(m, int32(16), base.I64_extend_i32_u(l0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L15
	}
L9:
	;
	return int32(0)
L10:
	;
	if v16 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v24 = F_SysCacheGetAttrNotNull(m, int32(21), v16, int32(13))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v27 = F_text_to_cstring(m, base.I32_wrap_i64(v24))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v31 = F_SysCacheGetAttrNotNull(m, int32(21), v16, int32(14))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v50 = v27
	v51 = v16
	v52 = v31
	goto L5
L15:
	;
	if v35 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v41 = F_SysCacheGetAttrNotNull(m, int32(16), v35, int32(8))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v44 = F_text_to_cstring(m, base.I32_wrap_i64(v41))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v48 = F_SysCacheGetAttrNotNull(m, int32(16), v35, int32(9))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v50 = v44
	v51 = v35
	v52 = v48
	goto L5
L20:
	;
	F_ReleaseCatCache(m, v51)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if base.B2i32(v60 == int32(0))|base.B2i32(v60 != v63) != 0 {
		v81 = v60
		v82 = v63
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v81-v82 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	v66 = v50
	v67 = v54
	goto L25
L25:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v71 == int32(0) {
		v81 = v71
		v82 = v70
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v81 = v71
	v82 = v70
	goto L23
L27:
	;
	v74 = int32(1)
	if v71 == v70 {
		v66 = v66 + v74
		v67 = v67 + v74
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v86 == int32(67) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v268 != int32(67) {
		goto L74
	} else {
		goto L75
	}
L32:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	if v89 == int32(0) {
		v671 = v3
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v92 = int32(_a_F_create_pg_locale_libc_0)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[1])))
	if base.B2i32(v95 == int32(0))|base.B2i32(v95 != v98) != 0 {
		v116 = v95
		v117 = v98
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L34
L36:
	;
	if v116-v117 == int32(0) {
		v671 = v3
		goto L1
	} else {
		goto L43
	}
L37:
	;
	goto L36
L38:
	;
	v101 = v54
	v102 = v92
	goto L39
L39:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	if v106 == int32(0) {
		v116 = v106
		v117 = v105
		goto L37
	} else {
		goto L41
	}
L40:
	;
	v116 = v106
	v117 = v105
	goto L37
L41:
	;
	v109 = int32(1)
	if v106 == v105 {
		v101 = v101 + v109
		v102 = v102 + v109
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v122 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[2])) = v122
	v130 = m.G0
	v132 = v130 - int32(32)
	m.G0 = v132
	v137 = v122
	goto L47
L44:
	;
	if v260 != 0 {
		v671 = v260
		goto L1
	} else {
		goto L72
	}
L45:
	;
	m.G0 = v132 + int32(32)
	goto L44
L46:
	;
	v260 = int32(0)
	goto L45
L47:
	;
	v142 = v137 << (uint(int32(2)) % 32)
	v148 = int32(1) << (uint(v137) % 32) & int32(9)
	if v148|int32(1) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v167 = F___loc_is_allocated(m, v122)
	mBase = m.M
	if v167 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142+(v132+int32(8))))) = v159
	if v159 == int32(-1) {
		goto L46
	} else {
		goto L56
	}
L50:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v122+v142)))
	v159 = v155
	goto L49
L51:
	;
	goto L52
L52:
	;
	if v148 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v157 = v50
	goto L55
L54:
	;
	v157 = int32(_a_F_create_pg_locale_libc_1)
	goto L55
L55:
	;
	v158 = F___get_locale(m, v137, v157)
	mBase = m.M
	v159 = v158
	goto L49
L56:
	;
	v164 = v137 + int32(1)
	if v164 != int32(6) {
		v137 = v164
		goto L47
	} else {
		goto L57
	}
L57:
	;
	goto L48
L58:
	;
	v170 = int32(_a_F_create_pg_locale_libc_2)
	v172 = v132 + int32(8)
	v175 = F_memcmp(m, v172, v170, int32(24))
	mBase = m.M
	if v175 == int32(0) {
		v260 = v170
		goto L45
	} else {
		goto L61
	}
L59:
	;
	v239 = v122
	goto L60
L60:
	;
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v132)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v239)+16)) = v244
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v132)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v239)+8)) = v246
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v132)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v239))) = v248
	v260 = v239
	goto L45
L61:
	;
	v178 = int32(_a_F_create_pg_locale_libc_3)
	v181 = F_memcmp(m, v172, v178, int32(24))
	mBase = m.M
	if v181 == int32(0) {
		v260 = v178
		goto L45
	} else {
		goto L62
	}
L62:
	;
	v184 = int32(0)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[3])))
	if v186 == v184 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v192 = v184
	goto L66
L64:
	;
	goto L65
L65:
	;
	v219 = int32(_a_F_create_pg_locale_libc_4)
	v221 = v132 + int32(8)
	v224 = F_memcmp(m, v221, v219, int32(24))
	mBase = m.M
	if v224 == int32(0) {
		v260 = v219
		goto L45
	} else {
		goto L69
	}
L66:
	;
	v199 = F___get_locale(m, v192, int32(_a_F_create_pg_locale_libc_1))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v192<<(uint(int32(2))%32))+uint32(_c_F_create_pg_locale_libc[4]))) = v199
	v202 = v192 + int32(1)
	if v202 != int32(6) {
		v192 = v202
		goto L66
	} else {
		goto L68
	}
L67:
	;
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[3])) = uint8(v206)
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[5])) = v210
	goto L65
L68:
	;
	goto L67
L69:
	;
	v227 = int32(_a_F_create_pg_locale_libc_5)
	v230 = F_memcmp(m, v221, v227, int32(24))
	mBase = m.M
	if v230 == int32(0) {
		v260 = v227
		goto L45
	} else {
		goto L70
	}
L70:
	;
	v234 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v234 == int32(0) {
		goto L46
	} else {
		goto L71
	}
L71:
	;
	v239 = v234
	goto L60
L72:
	;
	goto L2
L73:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v453 != int32(67) {
		goto L116
	} else {
		goto L117
	}
L74:
	;
	v273 = int32(_a_F_create_pg_locale_libc_0)
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[1])))
	if base.B2i32(v276 == int32(0))|base.B2i32(v276 != v279) != 0 {
		v297 = v276
		v298 = v279
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v271 != 0 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v452 = int32(0)
	goto L73
L77:
	;
	if v297-v298 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L78:
	;
	goto L77
L79:
	;
	v282 = v50
	v283 = v273
	goto L80
L80:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+1)))
	if v287 == int32(0) {
		v297 = v287
		v298 = v286
		goto L78
	} else {
		goto L82
	}
L81:
	;
	v297 = v287
	v298 = v286
	goto L78
L82:
	;
	v290 = int32(1)
	if v287 == v286 {
		v282 = v282 + v290
		v283 = v283 + v290
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v452 = int32(0)
	goto L73
L85:
	;
	goto L86
L86:
	;
	v304 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[2])) = v304
	v312 = m.G0
	v314 = v312 - int32(32)
	m.G0 = v314
	v319 = v304
	goto L90
L87:
	;
	if v442 == int32(0) {
		goto L2
	} else {
		goto L115
	}
L88:
	;
	m.G0 = v314 + int32(32)
	goto L87
L89:
	;
	v442 = int32(0)
	goto L88
L90:
	;
	v324 = v319 << (uint(int32(2)) % 32)
	v330 = int32(1) << (uint(v319) % 32) & int32(8)
	if v330|int32(1) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v349 = F___loc_is_allocated(m, v304)
	mBase = m.M
	if v349 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v324+(v314+int32(8))))) = v341
	if v341 == int32(-1) {
		goto L89
	} else {
		goto L99
	}
L93:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v304+v324)))
	v341 = v337
	goto L92
L94:
	;
	goto L95
L95:
	;
	if v330 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v339 = v50
	goto L98
L97:
	;
	v339 = int32(_a_F_create_pg_locale_libc_1)
	goto L98
L98:
	;
	v340 = F___get_locale(m, v319, v339)
	mBase = m.M
	v341 = v340
	goto L92
L99:
	;
	v346 = v319 + int32(1)
	if v346 != int32(6) {
		v319 = v346
		goto L90
	} else {
		goto L100
	}
L100:
	;
	goto L91
L101:
	;
	v352 = int32(_a_F_create_pg_locale_libc_2)
	v354 = v314 + int32(8)
	v357 = F_memcmp(m, v354, v352, int32(24))
	mBase = m.M
	if v357 == int32(0) {
		v442 = v352
		goto L88
	} else {
		goto L104
	}
L102:
	;
	v421 = v304
	goto L103
L103:
	;
	v426 = *(*int64)(unsafe.Add(mBase, uint32(v314)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v421)+16)) = v426
	v428 = *(*int64)(unsafe.Add(mBase, uint32(v314)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v421)+8)) = v428
	v430 = *(*int64)(unsafe.Add(mBase, uint32(v314)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v421))) = v430
	v442 = v421
	goto L88
L104:
	;
	v360 = int32(_a_F_create_pg_locale_libc_3)
	v363 = F_memcmp(m, v354, v360, int32(24))
	mBase = m.M
	if v363 == int32(0) {
		v442 = v360
		goto L88
	} else {
		goto L105
	}
L105:
	;
	v366 = int32(0)
	v368 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[3])))
	if v368 == v366 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v374 = v366
	goto L109
L107:
	;
	goto L108
L108:
	;
	v401 = int32(_a_F_create_pg_locale_libc_4)
	v403 = v314 + int32(8)
	v406 = F_memcmp(m, v403, v401, int32(24))
	mBase = m.M
	if v406 == int32(0) {
		v442 = v401
		goto L88
	} else {
		goto L112
	}
L109:
	;
	v381 = F___get_locale(m, v374, int32(_a_F_create_pg_locale_libc_1))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v374<<(uint(int32(2))%32))+uint32(_c_F_create_pg_locale_libc[4]))) = v381
	v384 = v374 + int32(1)
	if v384 != int32(6) {
		v374 = v384
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v388 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[3])) = uint8(v388)
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[5])) = v392
	goto L108
L111:
	;
	goto L110
L112:
	;
	v409 = int32(_a_F_create_pg_locale_libc_5)
	v412 = F_memcmp(m, v403, v409, int32(24))
	mBase = m.M
	if v412 == int32(0) {
		v442 = v409
		goto L88
	} else {
		goto L113
	}
L113:
	;
	v416 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v416 == int32(0) {
		goto L89
	} else {
		goto L114
	}
L114:
	;
	v421 = v416
	goto L103
L115:
	;
	v452 = v442
	goto L73
L116:
	;
	v457 = int32(_a_F_create_pg_locale_libc_0)
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v463 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[1])))
	if base.B2i32(v460 == int32(0))|base.B2i32(v460 != v463) != 0 {
		v481 = v460
		v482 = v463
		goto L120
	} else {
		goto L121
	}
L117:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	if v456 != 0 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v671 = v452
	goto L1
L119:
	;
	if v481-v482 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L120:
	;
	goto L119
L121:
	;
	v466 = v54
	v467 = v457
	goto L122
L122:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467)+1)))
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466)+1)))
	if v471 == int32(0) {
		v481 = v471
		v482 = v470
		goto L120
	} else {
		goto L124
	}
L123:
	;
	v481 = v471
	v482 = v470
	goto L120
L124:
	;
	v474 = int32(1)
	if v471 == v470 {
		v466 = v466 + v474
		v467 = v467 + v474
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v671 = v452
	goto L1
L127:
	;
	goto L128
L128:
	;
	v487 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[2])) = v487
	v494 = m.G0
	v496 = v494 - int32(32)
	m.G0 = v496
	v501 = v487
	goto L132
L129:
	;
	if v624 != 0 {
		v671 = v624
		goto L1
	} else {
		goto L157
	}
L130:
	;
	m.G0 = v496 + int32(32)
	goto L129
L131:
	;
	v624 = int32(0)
	goto L130
L132:
	;
	v506 = v501 << (uint(int32(2)) % 32)
	v512 = int32(1) << (uint(v501) % 32) & int32(1)
	v513 = int32(0)
	if v512|base.B2i32(v452 == v513) == v513 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	v531 = F___loc_is_allocated(m, v452)
	mBase = m.M
	if v531 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506+(v496+int32(8))))) = v523
	if v523 == int32(-1) {
		goto L131
	} else {
		goto L141
	}
L135:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v452+v506)))
	v523 = v519
	goto L134
L136:
	;
	goto L137
L137:
	;
	if v512 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v521 = v54
	goto L140
L139:
	;
	v521 = int32(_a_F_create_pg_locale_libc_1)
	goto L140
L140:
	;
	v522 = F___get_locale(m, v501, v521)
	mBase = m.M
	v523 = v522
	goto L134
L141:
	;
	v528 = v501 + int32(1)
	if v528 != int32(6) {
		v501 = v528
		goto L132
	} else {
		goto L142
	}
L142:
	;
	goto L133
L143:
	;
	v534 = int32(_a_F_create_pg_locale_libc_2)
	v536 = v496 + int32(8)
	v539 = F_memcmp(m, v536, v534, int32(24))
	mBase = m.M
	if v539 == int32(0) {
		v624 = v534
		goto L130
	} else {
		goto L146
	}
L144:
	;
	v603 = v452
	goto L145
L145:
	;
	v608 = *(*int64)(unsafe.Add(mBase, uint32(v496)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v603)+16)) = v608
	v610 = *(*int64)(unsafe.Add(mBase, uint32(v496)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v603)+8)) = v610
	v612 = *(*int64)(unsafe.Add(mBase, uint32(v496)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v603))) = v612
	v624 = v603
	goto L130
L146:
	;
	v542 = int32(_a_F_create_pg_locale_libc_3)
	v545 = F_memcmp(m, v536, v542, int32(24))
	mBase = m.M
	if v545 == int32(0) {
		v624 = v542
		goto L130
	} else {
		goto L147
	}
L147:
	;
	v548 = int32(0)
	v550 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[3])))
	if v550 == v548 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v556 = v548
	goto L151
L149:
	;
	goto L150
L150:
	;
	v583 = int32(_a_F_create_pg_locale_libc_4)
	v585 = v496 + int32(8)
	v588 = F_memcmp(m, v585, v583, int32(24))
	mBase = m.M
	if v588 == int32(0) {
		v624 = v583
		goto L130
	} else {
		goto L154
	}
L151:
	;
	v563 = F___get_locale(m, v556, int32(_a_F_create_pg_locale_libc_1))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v556<<(uint(int32(2))%32))+uint32(_c_F_create_pg_locale_libc[4]))) = v563
	v566 = v556 + int32(1)
	if v566 != int32(6) {
		v556 = v566
		goto L151
	} else {
		goto L153
	}
L152:
	;
	v570 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[3])) = uint8(v570)
	v574 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[5])) = v574
	goto L150
L153:
	;
	goto L152
L154:
	;
	v591 = int32(_a_F_create_pg_locale_libc_5)
	v594 = F_memcmp(m, v585, v591, int32(24))
	mBase = m.M
	if v594 == int32(0) {
		v624 = v591
		goto L130
	} else {
		goto L155
	}
L155:
	;
	v598 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v598 == int32(0) {
		goto L131
	} else {
		goto L156
	}
L156:
	;
	v603 = v598
	goto L145
L157:
	;
	if v452 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v632 = F___loc_is_allocated(m, v452)
	mBase = m.M
	if v632 != 0 {
		goto L162
	} else {
		goto L163
	}
L159:
	;
	goto L160
L160:
	;
	F_report_newlocale_failure(m, v54)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L9
	} else {
		goto L165
	}
L161:
	;
	goto L160
L162:
	;
	F_emscripten_builtin_free(m, v452)
	mBase = m.M
	goto L164
L163:
	;
	goto L164
L164:
	;
	goto L161
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	v641 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v641
	F_errmsg_internal(m, int32(_a_F_create_pg_locale_libc_6), v9)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L9
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_create_pg_locale_libc_7), int32(782), int32(_a_F_create_pg_locale_libc_8))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L9
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_create_pg_locale_libc_9), v9+int32(16))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L9
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(_a_F_create_pg_locale_libc_7), int32(799), int32(_a_F_create_pg_locale_libc_8))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L9
	} else {
		goto L171
	}
L171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	v676 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v674))) = uint8(v676)
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v678 == int32(67) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v674)+1)) = uint8(v713)
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v715 != int32(67) {
		goto L187
	} else {
		goto L188
	}
L175:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v681 == int32(0) {
		v713 = int32(1)
		goto L174
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v684 = int32(_a_F_create_pg_locale_libc_0)
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v690 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[1])))
	if base.B2i32(v687 == int32(0))|base.B2i32(v687 != v690) != 0 {
		v708 = v687
		v709 = v690
		goto L180
	} else {
		goto L181
	}
L178:
	;
	goto L177
L179:
	;
	v713 = base.B2i32(v708-v709 == int32(0))
	goto L174
L180:
	;
	goto L179
L181:
	;
	v693 = v50
	v694 = v684
	goto L182
L182:
	;
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694)+1)))
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v693)+1)))
	if v698 == int32(0) {
		v708 = v698
		v709 = v697
		goto L180
	} else {
		goto L184
	}
L183:
	;
	v708 = v698
	v709 = v697
	goto L180
L184:
	;
	v701 = int32(1)
	if v698 == v697 {
		v693 = v693 + v701
		v694 = v694 + v701
		goto L182
	} else {
		goto L185
	}
L185:
	;
	goto L183
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v674)+12)) = v671
	*(*uint8)(unsafe.Add(mBase, uint32(v674)+2)) = uint8(v749)
	if v713 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L187:
	;
	v720 = int32(_a_F_create_pg_locale_libc_0)
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v726 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[1])))
	if base.B2i32(v723 == int32(0))|base.B2i32(v723 != v726) != 0 {
		v744 = v723
		v745 = v726
		goto L191
	} else {
		goto L192
	}
L188:
	;
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	if v718 != 0 {
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v749 = int32(1)
	goto L186
L190:
	;
	v749 = base.B2i32(v744-v745 == int32(0))
	goto L186
L191:
	;
	goto L190
L192:
	;
	v729 = v54
	v730 = v720
	goto L193
L193:
	;
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730)+1)))
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v729)+1)))
	if v734 == int32(0) {
		v744 = v734
		v745 = v733
		goto L191
	} else {
		goto L195
	}
L194:
	;
	v744 = v734
	v745 = v733
	goto L191
L195:
	;
	v737 = int32(1)
	if v734 == v733 {
		v729 = v729 + v737
		v730 = v730 + v737
		goto L193
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v674)+4)) = int32(_a_F_create_pg_locale_libc_10)
	goto L199
L198:
	;
	goto L199
L199:
	;
	if v749 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	m.G0 = v9 + int32(32)
	return v674
L201:
	;
	v757 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[6]))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v757)+4))
	goto L202
L202:
	;
	if v758 == int32(6) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v674)+8)) = int32(_a_F_create_pg_locale_libc_11)
	goto L200
L204:
	;
	goto L205
L205:
	;
	v764 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[6]))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v764)+4))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v765*int32(28))+uint32(_c_F_create_pg_locale_libc[7])))
	goto L206
L206:
	;
	if int32(2) <= v770 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v674)+8)) = int32(_a_F_create_pg_locale_libc_12)
	goto L200
L208:
	;
	goto L209
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v674)+8)) = int32(_a_F_create_pg_locale_libc_13)
	goto L200
}
func F_pg_analyze_and_rewrite_fixedparams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_fixedparams[0])))
	if v7 == int32(1) {
		v10 = int32(_a_F_pg_analyze_and_rewrite_fixedparams_0)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_fixedparams[1])) = int64(4)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_fixedparams[2])) = int64(3)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_fixedparams[3])) = int64(2)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_fixedparams[4])) = int64(1)
		v20 = F___syscall_ret(m, int32(0))
		mBase = m.M
		F_gettimeofday(m, int32(_a_F_pg_analyze_and_rewrite_fixedparams_1))
		mBase = m.M
	} else {
	}
	v23 = F_parse_analyze_fixedparams(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		v28 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_analyze_and_rewrite_fixedparams[0])))
		if v28 == int32(1) {
			F_ShowUsage(m, int32(_a_F_pg_analyze_and_rewrite_fixedparams_2))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = F_pg_rewrite_query(m, v23)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					return v34
				}
			}
		} else {
			v34 = F_pg_rewrite_query(m, v23)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				return v34
			}
		}
	}
}
func F_pg_ascii_dsplen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v2 = int32(-1)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v5 == int32(127) {
		v8 = v2
	} else {
		v8 = int32(1)
	}
	if base.Ui32(v5) < base.Ui32(int32(32)) {
		v11 = v2
	} else {
		v11 = v8
	}
	if v5 != 0 {
		v13 = v11
	} else {
		v13 = int32(0)
	}
	return v13
}
func F_pg_available_extensions(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v178 int32
	_ = v178
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v20 = F_get_extension_control_directories(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v12 + int32(48)
	return int64(0)
L4:
	;
	if v20 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v24 <= int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v34 = v2
	v35 = v2
	goto L7
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v34<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v42 = F_AllocateDir(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L3
L9:
	;
	v251 = v34 + int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v251 < v252 {
		v34 = v251
		v35 = v249
		goto L7
	} else {
		goto L73
	}
L10:
	;
	if v42 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_pg_available_extensions[0]))
	if v47 == int32(44) {
		v249 = v35
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v51 = F_ReadDir(m, v42, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	if v51 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v53 = v51
	v61 = v35
	goto L19
L17:
	;
	v238 = v35
	goto L18
L18:
	;
	F_FreeDir(m, v42)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L72
	}
L19:
	;
	v63 = v53 + int32(19)
	v67 = F_strlen(m, v63)
	mBase = m.M
	v74 = v67 + int32(1)
	goto L24
L20:
	;
	v238 = v226
	goto L18
L21:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v228 = F_ReadDir(m, v42, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L70
	}
L22:
	;
	if v86 == int32(0) {
		v226 = v61
		goto L21
	} else {
		goto L28
	}
L23:
	;
	goto L22
L24:
	;
	v76 = int32(0)
	if v74 == v76 {
		v86 = v76
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v86 = v81
	goto L23
L26:
	;
	v80 = v74 - int32(1)
	v81 = v63 + v80
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v82 != int32(46) {
		v74 = v80
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v89 = int32(_a_F_pg_available_extensions_0)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_available_extensions[1])))
	if base.B2i32(v92 == int32(0))|base.B2i32(v92 != v95) != 0 {
		v113 = v92
		v114 = v95
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v113-v114 != 0 {
		v226 = v61
		goto L21
	} else {
		goto L36
	}
L30:
	;
	goto L29
L31:
	;
	v98 = v86
	v99 = v89
	goto L32
L32:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	if v103 == int32(0) {
		v113 = v103
		v114 = v102
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v113 = v103
	v114 = v102
	goto L30
L34:
	;
	v106 = int32(1)
	if v103 == v102 {
		v98 = v98 + v106
		v99 = v99 + v106
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v116 = F_pstrdup(m, v63)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v121 = F_strlen(m, v116)
	mBase = m.M
	v128 = v121 + int32(1)
	goto L40
L38:
	;
	v141 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v141)
	v144 = F_strstr(m, v116, int32(_a_F_pg_available_extensions_1))
	mBase = m.M
	if v144 != 0 {
		v226 = v61
		goto L21
	} else {
		goto L44
	}
L39:
	;
	goto L38
L40:
	;
	v130 = int32(0)
	if v128 == v130 {
		v140 = v130
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v140 = v135
	goto L39
L42:
	;
	v134 = v128 - int32(1)
	v135 = v116 + v134
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v136 != int32(46) {
		v128 = v134
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v145 = F_makeString(m, v116)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v147 = F_list_member(m, v61, v145)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v147 != 0 {
		v226 = v61
		goto L21
	} else {
		goto L47
	}
L47:
	;
	v149 = F_lappend(m, v61, v145)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v152 = F_palloc0(m, int32(48))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v154 = F_pstrdup(m, v116)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+36)) = int32(-1)
	v158 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v152)+34)) = uint8(v158)
	v160 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v152)+32)) = uint16(v160)
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v154
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v164 = F_pstrdup(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+8)) = v164
	F_parse_extension_control_file(m, v152, int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v170 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v170
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v170
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v170
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v170
	v178 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v178
	v182 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v152))))
	v183 = F_DirectFunctionCall1Coll(m, int32(534), v178, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v183
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
	if v186 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v196 = F_superuser(m)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L60
	}
L55:
	;
	v189 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v189)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v191 = F_cstring_to_text(m, v186)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = base.I64_extend_i32_u(v191)
	goto L54
L59:
	;
	v203 = F_cstring_to_text(m, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L63
	}
L60:
	;
	if v196 == int32(0) {
		v202 = int32(_a_F_pg_available_extensions_2)
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v200 != 0 {
		v202 = v200
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v202 = v201
	goto L59
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = base.I64_extend_i32_u(v203)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v152)+24))
	if v207 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	F_tuplestore_putvalues(m, v216, v217, v12+int32(16), v12+int32(12))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L69
	}
L65:
	;
	v210 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v210)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v212 = F_cstring_to_text(m, v207)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = base.I64_extend_i32_u(v212)
	goto L64
L69:
	;
	v226 = v149
	goto L21
L70:
	;
	if v228 != 0 {
		v53 = v228
		v61 = v226
		goto L19
	} else {
		goto L71
	}
L71:
	;
	goto L20
L72:
	;
	v249 = v238
	goto L9
L73:
	;
	goto L8
}
func F_pg_base64_decode_1(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_pg_base64_decode_internal(m, l0, l1, l2, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_pg_base64_encode(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	v4 = int32(0)
	if l1 == v4 {
		v100 = l2
	} else {
		v15 = l0
		v16 = int32(2)
		v18 = l2
		v19 = v4
		v20 = l2 + int32(76)
		for {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			v27 = v23<<(uint(v16<<(uint(int32(3))%32))%32) | v19
			if int32(0) < v16 {
				v63 = v16 - int32(1)
				v64 = v18
				v65 = v27
				v66 = v20
			} else {
				v32 = int32(63)
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27&v32)+uint32(_c_F_pg_base64_encode[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v18)+3)) = uint8(v34)
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v27)>>(uint(int32(18))%32)))+uint32(_c_F_pg_base64_encode[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v38)
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v27)>>(uint(int32(6))%32))&v32)+uint32(_c_F_pg_base64_encode[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v18)+2)) = uint8(v44)
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v27)>>(uint(int32(12))%32))&v32)+uint32(_c_F_pg_base64_encode[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)) = uint8(v50)
				v52 = int32(0)
				v53 = int32(2)
				v55 = v18 + int32(4)
				if base.Ui32(v55) < base.Ui32(v20) {
					v63 = v53
					v64 = v55
					v65 = v52
					v66 = v20
				} else {
					v57 = int32(10)
					*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)) = uint8(v57)
					v63 = v53
					v64 = v18 + int32(5)
					v65 = v52
					v66 = v18 + int32(81)
				}
			}
			v69 = v15 + int32(1)
			if base.Ui32(v69) < base.Ui32(l0+l1) {
				v15 = v69
				v16 = v63
				v18 = v64
				v19 = v65
				v20 = v66
				continue
			} else {
				break
			}
			break
		}
		if v63 == int32(2) {
			v100 = v64
		} else {
			v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v65)>>(uint(int32(18))%32)))+uint32(_c_F_pg_base64_encode[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v75)
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v65)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_pg_base64_encode[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)) = uint8(v81)
			if v63 == int32(0) {
				v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v65)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_pg_base64_encode[0]))))
				v91 = v90
			} else {
				v91 = int32(61)
			}
			v92 = int32(61)
			*(*uint8)(unsafe.Add(mBase, uint32(v64)+3)) = uint8(v92)
			*(*uint8)(unsafe.Add(mBase, uint32(v64)+2)) = uint8(v91)
			v100 = v64 + int32(4)
		}
	}
	return base.I64_extend_i32_s(v100 - l2)
}
func F_pg_big5_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	if l1 <= int32(0) {
		v38 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v38 - l0
L2:
	;
	v8 = l1
	v10 = l0
	goto L3
L3:
	;
	v11 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
	if int32(0) <= v11 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v38 = v31
	goto L1
L5:
	;
	v31 = v30 + v10
	v32 = v8 - v30
	if int32(0) < v32 {
		v8 = v32
		v10 = v31
		goto L3
	} else {
		goto L12
	}
L6:
	;
	if v11 == int32(0) {
		v38 = v10
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if v8 == int32(1) {
		v38 = v10
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v30 = int32(1)
	goto L5
L10:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if base.B2i32(v11 == int32(-115))&base.B2i32(v21 == int32(32))|base.B2i32(v21 == int32(0)) != 0 {
		v38 = v10
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v30 = int32(2)
	goto L5
L12:
	;
	goto L4
}
func F_pg_column_is_updatable(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	if v3 <= int32(0) {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		v16 = F_bms_make_singleton(m, base.I32_extend16_s(v3+int32(7)))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = F_relation_is_updatable(m, v8, int32(0), base.B2i32(v10 != int64(0)), v16)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				v22 = int32(20)
				return base.I64_extend_i32_u(base.B2i32(v20&v22 == v22))
			}
		}
	}
}
func F_pg_column_toast_chunk_id(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v11 == int32(0) {
		v15 = F_get_fn_expr_argtype(m, v10, int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = F_get_typlen(m, v15)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				if v19 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
						F_errmsg_internal(m, int32(_a_F_pg_column_toast_chunk_id_0), v8)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_column_toast_chunk_id_1), int32(_a_F_pg_column_toast_chunk_id_2), int32(_a_F_pg_column_toast_chunk_id_3))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
					v26 = F_MemoryContextAlloc(m, v24, int32(4))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v26
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = v19
						v34 = v19
						if v34 != int32(-1) {
							v38 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
							v53 = int64(0)
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
							if v42 == int32(1) {
								v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
								if v45 == int32(18) {
									v51 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v41)+10)))
									v53 = v51
								} else {
									v48 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v48)
									v53 = int64(0)
								}
							} else {
								v48 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v48)
								v53 = int64(0)
							}
						}
						m.G0 = v8 + int32(16)
						return v53
					}
				}
			}
		}
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v34 = v33
		if v34 != int32(-1) {
			v38 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
			v53 = int64(0)
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
			if v42 == int32(1) {
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
				if v45 == int32(18) {
					v51 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v41)+10)))
					v53 = v51
				} else {
					v48 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v48)
					v53 = int64(0)
				}
			} else {
				v48 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v48)
				v53 = int64(0)
			}
		}
		m.G0 = v8 + int32(16)
		return v53
	}
}
func F_pg_control_init(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v47 int64
	_ = v47
	var v51 int64
	_ = v51
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v63 int64
	_ = v63
	var v67 int64
	_ = v67
	var v71 int64
	_ = v71
	var v75 int64
	_ = v75
	var v79 int64
	_ = v79
	var v83 int64
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v4 = m.G0
	v6 = v4 - int32(128)
	m.G0 = v6
	v11 = F_get_call_result_type(m, l0, int32(0), v6+int32(16))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		if v11 == int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_init[0]))
			v22 = F_LWLockAcquire(m, v18+int32(1152), int32(1))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_init[1]))
				v28 = F_get_controlfile(m, v25, v6+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_init[0]))
					F_LWLockRelease(m, v31+int32(1152))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
						if v36 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_pg_control_init_0), int32(0))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_control_init_1), int32(225), int32(_a_F_pg_control_init_2))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v39 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+212)))
							v40 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+20)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v39
							v43 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+224)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+21)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+40)) = v43
							v47 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+228)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+22)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+48)) = v47
							v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+236)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+23)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = v51
							v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+240)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+24)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+64)) = v55
							v59 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+244)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+25)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+72)) = v59
							v63 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+248)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+26)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+80)) = v63
							v67 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+252)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+27)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+88)) = v67
							v71 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+256)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+28)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+96)) = v71
							v75 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v28)+260)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+29)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+104)) = v75
							v79 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+264)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+30)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+112)) = v79
							v83 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+31)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+120)) = v83
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
							v92 = F_heap_form_tuple(m, v87, v6+int32(32), v6+int32(20))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int64(0)
							} else {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
								v95 = F_HeapTupleHeaderGetDatum(m, v94)
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int64(0)
								} else {
									m.G0 = v6 + int32(128)
									return v95
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v104 = m.ExcPending
			if v104 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_control_init_3), int32(0))
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_control_init_1), int32(217), int32(_a_F_pg_control_init_2))
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return int64(0)
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
func F_pg_convert_from(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	v3 = int32(0)
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_convert_from[0]))
	v10 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9))))
	v11 = F_DirectFunctionCall1Coll(m, int32(534), v3, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = F_DirectFunctionCall3Coll(m, int32(1851), v3, v4, v5, v11)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			return v15
		}
	}
}
func F_pg_create_restore_point(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int64
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int64
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_create_restore_point[0])))
	if v16 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L66
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L61
	}
L5:
	;
	if v26 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_pg_create_restore_point[1]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+308))
	v24 = base.B2i32(v22 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_create_restore_point[0])) = uint8(v24)
	v26 = v24
	goto L8
L7:
	;
	v26 = int32(0)
	goto L8
L8:
	;
	goto L5
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_pg_create_restore_point[2]))
	if v30 <= int32(0) {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L56
	}
L12:
	;
	v33 = F_text_to_cstring(m, v10)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v35 = F_strlen(m, v33)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v35) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v38 = m.G0
	v40 = v38 - int32(96)
	m.G0 = v40
	v45 = m.G0
	v46 = int32(16)
	v47 = v45 - v46
	m.G0 = v47
	F_gettimeofday(m, v47)
	mBase = m.M
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
	v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v47)+8)))
	m.G0 = v47 + v46
	goto L15
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = v51 + v50*int64(1000000) - int64(946684800000000)
	v62 = v40 + int32(32)
	goto L19
L16:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L47
	}
L17:
	;
	v179 = F_strlen(m, v168)
	mBase = m.M
	goto L16
L19:
	;
	goto L20
L20:
	;
	v69 = int32(63)
	if (v62^v33)&int32(3) != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v169))) = uint8(v172)
	goto L17
L22:
	;
	v153 = v148
	v154 = v149
	v155 = v150
	goto L43
L23:
	;
	if v143 == int32(0) {
		v168 = v141
		v169 = v142
		goto L21
	} else {
		goto L42
	}
L24:
	;
	v141 = v33
	v142 = v62
	v143 = v69
	goto L23
L25:
	;
	goto L26
L26:
	;
	v73 = int32(0)
	if base.B2i32(v33&int32(3) == v73)|int32(0) == v73 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v109 == int32(0) {
		v168 = v106
		v169 = v107
		goto L21
	} else {
		goto L36
	}
L28:
	;
	v85 = v33
	v86 = v62
	v87 = v69
	goto L31
L29:
	;
	goto L30
L30:
	;
	v106 = v33
	v107 = v62
	v108 = v69
	v109 = int32(1)
	goto L27
L31:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v89)
	if v89 == int32(0) {
		v148 = v85
		v149 = v86
		v150 = v87
		goto L22
	} else {
		goto L33
	}
L32:
	;
	v106 = v100
	v107 = v94
	v108 = v96
	v109 = v98
	goto L27
L33:
	;
	v93 = int32(1)
	v94 = v86 + v93
	v96 = v87 - v93
	v97 = int32(0)
	v98 = base.B2i32(v96 != v97)
	v100 = v85 + v93
	if v100&int32(3) == v97 {
		v106 = v100
		v107 = v94
		v108 = v96
		v109 = v98
		goto L27
	} else {
		goto L34
	}
L34:
	;
	if v96 != 0 {
		v85 = v100
		v86 = v94
		v87 = v96
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if base.B2i32(v112 == int32(0))|base.B2i32(base.Ui32(v108) < base.Ui32(int32(4))) != 0 {
		v141 = v106
		v142 = v107
		v143 = v108
		goto L23
	} else {
		goto L37
	}
L37:
	;
	v119 = v106
	v120 = v107
	v121 = v108
	goto L38
L38:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v127 = int32(-2139062144)
	if (int32(16843008)-v124|v124)&v127 != v127 {
		v148 = v119
		v149 = v120
		v150 = v121
		goto L22
	} else {
		goto L40
	}
L39:
	;
	v141 = v135
	v142 = v133
	v143 = v137
	goto L23
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v124
	v132 = int32(4)
	v133 = v120 + v132
	v135 = v119 + v132
	v137 = v121 - v132
	if base.Ui32(int32(3)) < base.Ui32(v137) {
		v119 = v135
		v120 = v133
		v121 = v137
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v148 = v141
	v149 = v142
	v150 = v143
	goto L22
L43:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v157)
	if v157 == int32(0) {
		v168 = v153
		v169 = v154
		goto L21
	} else {
		goto L45
	}
L44:
	;
	v168 = v164
	v169 = v162
	goto L21
L45:
	;
	v161 = int32(1)
	v162 = v154 + v161
	v164 = v153 + v161
	v166 = v155 - v161
	if v166 != 0 {
		v153 = v164
		v154 = v162
		v155 = v166
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	F_XLogRegisterData(m, v40+int32(24), int32(72))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v191 = F_XLogInsert(m, int32(0), int32(112))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v195 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	if v195 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v33
	*(*uint32)(unsafe.Add(mBase, uint32(v40)+8)) = uint32(v191)
	v200 = int64(base.Ui64(v191) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v40)+4)) = uint32(v200)
	F_errmsg(m, int32(_a_F_pg_create_restore_point_0), v40)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	m.G0 = v40 + int32(96)
	m.G0 = v7 + int32(16)
	return v191
L54:
	;
	F_errfinish(m, int32(_a_F_pg_create_restore_point_1), int32(_a_F_pg_create_restore_point_2), int32(_a_F_pg_create_restore_point_3))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errmsg(m, int32(_a_F_pg_create_restore_point_4), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errhint(m, int32(_a_F_pg_create_restore_point_5), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_pg_create_restore_point_6), int32(273), int32(_a_F_pg_create_restore_point_7))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(_a_F_pg_create_restore_point_8), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errhint(m, int32(_a_F_pg_create_restore_point_9), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_pg_create_restore_point_6), int32(279), int32(_a_F_pg_create_restore_point_7))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(63)
	F_errmsg(m, int32(_a_F_pg_create_restore_point_10), v7)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_pg_create_restore_point_6), int32(286), int32(_a_F_pg_create_restore_point_7))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
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
func F_pg_cryptohash_create(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	v4 = F_palloc(m, int32(12))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = l0
			*(*int64)(unsafe.Add(mBase, uint32(v4)+4)) = int64(0)
			v15 = m.Env.Pgmem_hash_create(m, l0)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v15
			if int32(0) < v15 {
				return v4
			} else {
				F___memset(m, v4, int32(0), int32(12))
				mBase = m.M
				F_pfree(m, v4)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_pg_ddl_command_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_pg_ddl_command_in_0), int32(358), int32(_a_F_pg_ddl_command_in_1), int32(_a_F_pg_ddl_command_in_2), int32(_a_F_pg_ddl_command_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_pg_dearmor(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		if v16 == int32(1) {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
			if v22 == int32(18) {
				v25 = int32(16)
			} else {
				v25 = int32(0)
			}
			if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v32 = int32(4)
			} else {
				v32 = v25
			}
			v45 = v32
		} else {
			v33 = int32(1)
			if v16&v33 != 0 {
				v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		F_initStringInfo(m, v9)
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int64(0)
		} else {
			v48 = int32(1)
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v50&v48 != 0 {
				v53 = v48
			} else {
				v53 = int32(4)
			}
			v55 = F_pgp_armor_decode(m, v12+v53, v45, v9)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int64(0)
			} else {
				if int32(0) <= v55 {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					v62 = F_palloc(m, v59+int32(4))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int64(0)
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v62))) = v64<<(uint(int32(2))%32) + int32(16)
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						if v71 != 0 {
							base.MemoryCopy(m, v62+int32(4), v70, v71)
						} else {
						}
						F_pfree(m, v70)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int64(0)
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v77 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int64(0)
								} else {
									m.G0 = v9 + int32(16)
									return base.I64_extend_i32_u(v62)
								}
							} else {
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_u(v62)
							}
						}
					}
				} else {
					F_px_THROW_ERROR(m, v55)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int64(0)
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
func F_pg_dependencies_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v124 float64
	_ = v124
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v19 = F_statext_dependencies_deserialize(m, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = v10 + int32(-16)
	F_initStringInfo(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_appendStringInfoChar(m, v22, int32(91))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v36 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	F_appendStringInfoChar(m, v10+int32(-16), int32(93))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L32
	}
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(12)+v36<<(uint(int32(2))%32))))
	if int32(0) < v36 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	F_appendStringInfoString(m, v10+int32(-16), int32(_a_F_pg_dependencies_out_0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	if int32(1) < v51 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v64+v114<<(uint(int32(1))%32)))))
	v124 = *(*float64)(unsafe.Add(mBase, uint32(v43)))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v123
	F_appendStringInfo(m, v10+int32(-16), int32(_a_F_pg_dependencies_out_1), v12)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L30
	}
L16:
	;
	v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v54
	F_appendStringInfo(m, v10+int32(-16), int32(_a_F_pg_dependencies_out_2), v10+int32(-32))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L27
	}
L19:
	;
	v64 = v43 + int32(10)
	v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	if v66 <= int32(2) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v114 = v66 - int32(1)
	goto L15
L21:
	;
	goto L22
L22:
	;
	v71 = int32(1)
	goto L23
L23:
	;
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v64+v71<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v83
	F_appendStringInfo(m, v10+int32(-16), int32(_a_F_pg_dependencies_out_3), v10+int32(-48))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v114 = v96
	goto L15
L25:
	;
	v92 = int32(1)
	v93 = v71 + v92
	v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	v96 = v94 - v92
	if v93 < v96 {
		v71 = v93
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	F_errmsg_internal(m, int32(_a_F_pg_dependencies_out_4), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_pg_dependencies_out_5), int32(830), int32(_a_F_pg_dependencies_out_6))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v133 = v36 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if base.Ui32(v133) < base.Ui32(v134) {
		v36 = v133
		goto L9
	} else {
		goto L31
	}
L31:
	;
	goto L10
L32:
	;
	v150 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+48)))
	m.G0 = v12 - int32(-64)
	return v150
}
func F_pg_dependencies_recv(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_pg_dependencies_recv_0), int32(857), int32(_a_F_pg_dependencies_recv_1), int32(_a_F_pg_dependencies_recv_2), int32(_a_F_pg_dependencies_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_pg_detoast_datum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v2&int32(3) != 0 {
		v5 = F_detoast_attr(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			v9 = v5
			return v9
		}
	} else {
		v9 = l0
		return v9
	}
}
func F_pg_digest(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v18 = int32(1)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v22 = v20 & v18
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = v18
	goto L5
L4:
	;
	v23 = int32(4)
	goto L5
L5:
	;
	if v20 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v53 = F_downcase_truncate_identifier(m, v14+v23, v51, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v30 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v41 = int32(1)
	if v22 != 0 {
		v51 = int32(base.Ui32(v20)>>(uint(v41)%32)) - v41
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v33 = int32(16)
	goto L12
L11:
	;
	v33 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v40 = int32(4)
	goto L15
L14:
	;
	v40 = v33
	goto L15
L15:
	;
	v51 = v40
	goto L6
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v57 = F_px_find_digest(m, v53, v11+int32(12))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v57 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_pfree(m, v53)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L51
	}
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = m.T0[v64].(func(*base.Module, int32) int32)(m, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v68 = v65 + int32(4)
	v69 = F_palloc(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v68 << (uint(int32(2)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v75 = F_pg_detoast_datum_packed(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v107 = int32(1)
	if v77&v107 != 0 {
		goto L37
	} else {
		goto L38
	}
L26:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v77 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	if v83 == int32(18) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v94 = int32(1)
	if v77&v94 != 0 {
		v106 = int32(base.Ui32(v77)>>(uint(v94)%32)) - v94
		goto L25
	} else {
		goto L36
	}
L30:
	;
	v86 = int32(16)
	goto L32
L31:
	;
	v86 = int32(0)
	goto L32
L32:
	;
	if base.Ui32((v83-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v93 = int32(4)
	goto L35
L34:
	;
	v93 = v86
	goto L35
L35:
	;
	v106 = v93
	goto L25
L36:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v106 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
	goto L25
L37:
	;
	v111 = v107
	goto L39
L38:
	;
	v111 = int32(4)
	goto L39
L39:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	m.T0[v113].(func(*base.Module, int32, int32, int32))(m, v63, v75+v111, v106)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	m.T0[v118].(func(*base.Module, int32, int32))(m, v63, v69+int32(4))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	m.T0[v121].(func(*base.Module, int32))(m, v63)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v124 != v75 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_pfree(m, v75)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v128 != v14 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	F_pfree(m, v14)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	m.G0 = v11 + int32(16)
	return base.I64_extend_i32_u(v69)
L50:
	;
	goto L49
L51:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v57 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v53
	F_errmsg(m, int32(_a_F_pg_digest_0), v11)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L67
	}
L54:
	;
	v173 = int32(_a_F_pg_digest_1)
	goto L53
L55:
	;
	goto L56
L56:
	;
	v152 = int32(_a_F_pg_digest_2)
	goto L58
L57:
	;
	v173 = v167
	goto L53
L58:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
	if v57 != v155 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v167 = v165
	goto L57
L60:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v152)+20))
	if v157 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	v173 = int32(_a_F_pg_digest_3)
	goto L53
L64:
	;
	goto L65
L65:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
	if v57 != v161 {
		v152 = v152 + int32(16)
		goto L58
	} else {
		goto L66
	}
L66:
	;
	v167 = v157
	goto L57
L67:
	;
	F_errfinish(m, int32(_a_F_pg_digest_4), int32(513), int32(_a_F_pg_digest_5))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_error_on_null(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v2 == int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(67108994))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_error_on_null_0), int32(0))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_error_on_null_1), int32(201), int32(_a_F_pg_error_on_null_2))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		return v23
	}
}
func F_pg_eucjp_dsplen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v4 - int32(142) {
	case 0:
		v25 = int32(1)
		return v25
	case 1:
		return int32(2)
	default:
		if base.I32_extend8_s(v4) < int32(0) {
			return int32(2)
		} else {
			v14 = int32(-1)
			if v4 == int32(127) {
				v19 = v14
			} else {
				v19 = int32(1)
			}
			if base.Ui32(v4) < base.Ui32(int32(32)) {
				v22 = v14
			} else {
				v22 = v19
			}
			if v4 != 0 {
				v24 = v22
			} else {
				v24 = int32(0)
			}
			v25 = v24
			return v25
		}
	}
}
func F_pg_eucjp_mblen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v4 - int32(142) {
	case 0:
		v14 = int32(2)
		v16 = v14
	case 1:
		v16 = int32(3)
	default:
		if int32(0) <= base.I32_extend8_s(v4) {
			v13 = int32(1)
		} else {
			v13 = int32(2)
		}
		v14 = v13
		v16 = v14
	}
	return v16
}
func F_pg_eucjp_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	if l1 <= int32(0) {
		v72 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v72 - l0
L2:
	;
	v9 = l1
	v10 = l0
	goto L3
L3:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v14 = base.I32_extend8_s(v13)
	if int32(0) <= v14 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v72 = v66
	goto L1
L5:
	;
	v66 = v65 + v10
	v67 = v9 - v65
	if int32(0) < v67 {
		v9 = v67
		v10 = v66
		goto L3
	} else {
		goto L21
	}
L6:
	;
	if v14 == int32(0) {
		v72 = v10
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	switch v13 - int32(142) {
	case 0:
		goto L13
	case 1:
		goto L12
	default:
		goto L11
	}
L9:
	;
	v65 = int32(1)
	goto L5
L10:
	;
	v65 = int32(2)
	goto L5
L11:
	;
	if base.B2i32(v9 == int32(1))|base.B2i32(base.Ui32(int32(93)) < base.Ui32((v14+int32(95))&int32(255))) != 0 {
		v72 = v10
		goto L1
	} else {
		goto L19
	}
L12:
	;
	if base.Ui32(v9) < base.Ui32(int32(3)) {
		v72 = v10
		goto L1
	} else {
		goto L16
	}
L13:
	;
	if v9 == int32(1) {
		v72 = v10
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if base.Ui32(int32(193)) <= base.Ui32((v24+int32(32))&int32(255)) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v72 = v10
	goto L1
L16:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if base.Ui32(int32(93)) < base.Ui32((v33+int32(95))&int32(255)) {
		v72 = v10
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+2)))
	if base.Ui32(int32(94)) <= base.Ui32((v40+int32(95))&int32(255)) {
		v72 = v10
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v65 = int32(3)
	goto L5
L19:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if base.Ui32(int32(93)) < base.Ui32((v57+int32(95))&int32(255)) {
		v72 = v10
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L10
L21:
	;
	goto L4
}
func F_pg_euckr_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v5 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if int32(0) <= v5 {
		v29 = int32(1)
		return v29
	} else {
		if l1 < int32(2) {
			return int32(-1)
		} else {
			if base.Ui32(int32(93)) < base.Ui32((v5+int32(95))&int32(255)) {
				v29 = int32(-1)
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if base.Ui32((v21+int32(95))&int32(255)) < base.Ui32(int32(94)) {
					v28 = int32(2)
				} else {
					v28 = int32(-1)
				}
				v29 = v28
			}
			return v29
		}
	}
}
func F_pg_event_trigger_table_rewrite_oid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pg_event_trigger_table_rewrite_oid[0]))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		if v9 != 0 {
			m.G0 = v5 + int32(16)
			return base.I64_extend_i32_u(v9)
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50463299))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_pg_event_trigger_table_rewrite_oid_0)
					F_errmsg(m, int32(_a_F_pg_event_trigger_table_rewrite_oid_1), v5)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_event_trigger_table_rewrite_oid_2), int32(1640), int32(_a_F_pg_event_trigger_table_rewrite_oid_3))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int64(0)
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
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50463299))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_pg_event_trigger_table_rewrite_oid_0)
				F_errmsg(m, int32(_a_F_pg_event_trigger_table_rewrite_oid_1), v5)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_event_trigger_table_rewrite_oid_2), int32(1640), int32(_a_F_pg_event_trigger_table_rewrite_oid_3))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
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
func F_pg_extension_update_paths(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v206 int32
	_ = v206
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_check_valid_extension_name(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = F_palloc0(m, int32(48))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = F_pstrdup(m, v16)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = int32(-1)
	v31 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+34)) = uint8(v31)
	v33 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+32)) = uint16(v33)
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v27
	F_parse_extension_control_file(m, v25, v31)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v39 = F_get_ext_ver_list(m, v25)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	m.G0 = v13 - int32(-64)
	return int64(0)
L8:
	;
	if v39 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v43 <= int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v46 = v43
	v54 = int32(0)
	goto L11
L11:
	;
	if int32(0) < v46 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L7
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v54<<(uint(int32(2))%32))))
	v64 = v46
	v70 = int32(0)
	goto L16
L14:
	;
	v195 = v46
	goto L15
L15:
	;
	v206 = v54 + int32(1)
	if v206 < v195 {
		v46 = v195
		v54 = v206
		goto L11
	} else {
		goto L42
	}
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v70<<(uint(int32(2))%32))))
	if v78 != v62 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v195 = v182
	goto L15
L18:
	;
	v82 = F_find_update_path(m, v39, v62, v78, int32(0), int32(1))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v182 = v64
	goto L20
L20:
	;
	v193 = v70 + int32(1)
	if v193 < v182 {
		v64 = v182
		v70 = v193
		goto L16
	} else {
		goto L41
	}
L21:
	;
	v84 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v84
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v84
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v84
	v90 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+28)) = uint16(v90)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+30)) = uint8(v90)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v95 = F_cstring_to_text(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = base.I64_extend_i32_u(v95)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v100 = F_cstring_to_text(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = base.I64_extend_i32_u(v100)
	if v82 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	F_tuplestore_putvalues(m, v173, v174, v11+int32(-32), v11+int32(-36))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L40
	}
L25:
	;
	v106 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+30)) = uint8(v106)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v109 = v11 + int32(-52)
	F_initStringInfo(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	F_appendStringInfoString(m, v109, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if int32(0) < v115 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v119 = int32(0)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v156 = F_cstring_to_text(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L38
	}
L33:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129+v119<<(uint(int32(2))%32))))
	v135 = v11 + int32(-52)
	F_appendStringInfoString(m, v135, int32(_a_F_pg_extension_update_paths_0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	F_appendStringInfoString(m, v135, v133)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v142 = v119 + int32(1)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v142 < v143 {
		v119 = v142
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = base.I64_extend_i32_u(v156)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_pfree(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L24
L40:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v182 = v181
	goto L20
L41:
	;
	goto L17
L42:
	;
	goto L12
}
func F_pg_function_is_visible(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int64
	_ = v25
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_FunctionIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v25 = int64(0)
		} else {
			v25 = base.I64_extend_i32_u(v14)
		}
		m.G0 = v7 + int32(16)
		return v25
	}
}
func F_pg_get_backend_memory_contexts(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v148 int64
	_ = v148
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v285 int32
	_ = v285
	var v289 int64
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int64
	_ = v315
	var v317 int64
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
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
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v393 int32
	_ = v393
	v16 = m.G0
	v18 = v16 - int32(1200)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = int64(34359738372)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_backend_memory_contexts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v24
	v31 = F_hash_create(m, int32(_a_F_pg_get_backend_memory_contexts_0), int64(256), v18+int32(16), int32(1064))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_backend_memory_contexts[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v39
	v42 = int32(1)
	v44 = F_list_make1_impl(m, v42, v18)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v46 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v46
	if v44 == v46 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_hash_destroy(m, v31)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L82
	}
L6:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v50 <= int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v60 = v42
	v62 = v44
	v63 = int32(0)
	goto L8
L8:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v63<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v72
	v79 = F_hash_search(m, v31, v18+int32(8), int32(1), v18+int32(7))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L5
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v60
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v84
	v86 = int32(0)
	if v84 == v86 {
		v133 = v86
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v148 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1088)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1096)) = v148
	v152 = int32(0)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+32))
	m.T0[v159].(func(*base.Module, int32, int32, int32, int32, int32))(m, v84, v152, v152, v18+int32(1088), int32(1))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L23
	}
L12:
	;
	v89 = v86
	goto L13
L13:
	;
	v109 = F_hash_search(m, v31, v18-int32(-64), int32(0), v18+int32(1120))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L20
	}
L15:
	;
	goto L14
L16:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1120)))
	if v111 == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v115 = F_lcons_int(m, v114, v89)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v118
	if v118 != 0 {
		v89 = v115
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v133 = v115
	goto L11
L20:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_backend_memory_contexts_1), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_pg_get_backend_memory_contexts_2), int32(101), int32(_a_F_pg_get_backend_memory_contexts_3))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v164 = int32(0)
	base.MemoryFill(m, v18+int32(1120), v164, int32(80))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+1112)) = uint16(v164)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1104)) = int64(0)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v84)+32))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v84)+36))
	if v172 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v173 = int32(0)
	v174 = int32(_a_F_pg_get_backend_memory_contexts_4)
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_backend_memory_contexts[2])))
	if base.B2i32(v177 == v173)|base.B2i32(v177 != v180) != 0 {
		v198 = v177
		v199 = v180
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v203 = v171
	v204 = v152
	goto L26
L26:
	;
	if v203 != 0 {
		goto L41
	} else {
		goto L42
	}
L27:
	;
	if v200 != 0 {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v200 = v198 - v199
	goto L27
L29:
	;
	v183 = v171
	v184 = v174
	goto L30
L30:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+1)))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
	if v188 == int32(0) {
		v198 = v188
		v199 = v187
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v198 = v188
	v199 = v187
	goto L28
L32:
	;
	v191 = int32(1)
	if v188 == v187 {
		v183 = v183 + v191
		v184 = v184 + v191
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v201 = v172
	goto L36
L35:
	;
	v201 = v173
	goto L36
L36:
	;
	if v200 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v202 = v171
	goto L39
L38:
	;
	v202 = v172
	goto L39
L39:
	;
	v203 = v202
	v204 = v201
	goto L26
L40:
	;
	if v204 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v206 = F_cstring_to_text(m, v203)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v210 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+1104)) = uint8(v210)
	goto L40
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1120)) = base.I64_extend_i32_u(v206)
	goto L40
L45:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v236 = v234 - int32(482)
	if base.Ui32(v236) <= base.Ui32(int32(3)) {
		goto L57
	} else {
		goto L58
	}
L46:
	;
	v212 = F_strlen(m, v204)
	mBase = m.M
	if int32(1024) <= v212 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v231 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+1105)) = uint8(v231)
	goto L45
L49:
	;
	v216 = F_pg_mbcliplen(m, v204, v212, int32(1023))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	v218 = v212
	goto L51
L51:
	;
	if v218 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v218 = v216
	goto L51
L53:
	;
	base.MemoryCopy(m, v18-int32(-64), v204, v218)
	goto L55
L54:
	;
	goto L55
L55:
	;
	v223 = v18 - int32(-64)
	v225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v218+v223))) = uint8(v225)
	v227 = F_cstring_to_text(m, v223)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1128)) = base.I64_extend_i32_u(v227)
	goto L45
L57:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v236<<(uint(int32(2))%32))+uint32(_c_F_pg_get_backend_memory_contexts[3])))
	v243 = v241
	goto L59
L58:
	;
	v243 = int32(_a_F_pg_get_backend_memory_contexts_5)
	goto L59
L59:
	;
	v244 = F_cstring_to_text(m, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1136)) = base.I64_extend_i32_u(v244)
	if v133 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v311 = F_construct_array_builtin(m, v299, v298, int32(23))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L71
	}
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1144)) = int64(0)
	v252 = int32(0)
	v255 = F_palloc_mul(m, int32(8), v252)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1144)) = base.I64_extend_i32_s(v257)
	v261 = F_palloc_mul(m, int32(8), v257)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L66
	}
L65:
	;
	v298 = v252
	v299 = v255
	goto L61
L66:
	;
	v263 = int32(0)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v264 <= v263 {
		v298 = v257
		v299 = v261
		goto L61
	} else {
		goto L67
	}
L67:
	;
	v269 = v263
	goto L68
L68:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v289 = int64(*(*int32)(unsafe.Add(mBase, uint32(v285+v269<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v261+v269<<(uint(int32(3))%32)))) = v289
	v292 = v269 + int32(1)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v292 < v293 {
		v269 = v292
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v298 = v257
	v299 = v261
	goto L61
L70:
	;
	goto L69
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1152)) = base.I64_extend_i32_u(v311)
	v315 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+1088)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1168)) = v315
	v317 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+1092)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1184)) = v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1096))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1160)) = base.I64_extend_i32_u(v319)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1100))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1176)) = base.I64_extend_i32_u(v322)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1192)) = base.I64_extend_i32_u(v319 - v322)
	F_tuplestore_putvalues(m, v83, v82, v18+int32(1120), v18+int32(1104))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_list_free(m, v133)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+20))
	if v337 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v338 = v337
	v347 = v62
	goto L77
L75:
	;
	v365 = v62
	goto L76
L76:
	;
	v371 = int32(1)
	v374 = v63 + v371
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v374 < v375 {
		v60 = v60 + v371
		v62 = v365
		v63 = v374
		goto L8
	} else {
		goto L81
	}
L77:
	;
	v353 = F_lappend(m, v347, v338)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	v365 = v353
	goto L76
L79:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v338)+28))
	if v355 != 0 {
		v338 = v355
		v347 = v353
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	goto L9
L82:
	;
	m.G0 = v18 + int32(1200)
	return int64(0)
}
func F_pg_get_constraintdef_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
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
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v276 int64
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int64
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int64
	_ = v343
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v352 int64
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v420 int32
	_ = v420
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v467 int32
	_ = v467
	var v470 int64
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int64
	_ = v545
	var v552 int32
	_ = v552
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
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
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int64
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v651 int32
	_ = v651
	var v655 int64
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v718 int32
	_ = v718
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v763 int32
	_ = v763
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	v13 = m.G0
	v15 = v13 - int32(432)
	m.G0 = v15
	v17 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = F_RegisterSnapshot(m, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = v15 + int32(304)
	F_ScanKeyInit(m, v28, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v36 = int32(1)
	v38 = F_systable_beginscan(m, v25, int32(2667), v36, v21, v36, v28)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v40 = F_systable_getnext(m, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_UnregisterSnapshot(m, v21)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v40 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L219
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L216
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L213
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L210
	}
L13:
	;
	m.G0 = v15 + int32(432)
	return v763
L14:
	;
	if l3 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+22)))
	v67 = v65 + v66
	v69 = v15 + int32(360)
	F_initStringInfo(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L25
	}
L17:
	;
	F_systable_endscan(m, v38)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	F_relation_close(m, v25, int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v763 = int32(0)
	goto L13
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_0), v15)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(2234), int32(_a_F_pg_get_constraintdef_worker_2))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	if l1 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+72)))
	switch v142 - int32(99) {
	case 0:
		goto L51
	default:
		goto L48
	case 3:
		goto L54
	case 11:
		goto L50
	case 13:
		v292 = int32(_a_F_pg_get_constraintdef_worker_3)
		goto L52
	case 17:
		goto L47
	case 18:
		goto L53
	case 21:
		goto L49
	}
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
	if v74 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v75 = F_generate_qualified_relation_name(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v67)+84))
	v91 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v89))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L34
	}
L31:
	;
	v79 = F_quote_identifier(m, v67+int32(4))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+292)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v15)+288)) = v75
	F_appendStringInfo(m, v69, int32(_a_F_pg_get_constraintdef_worker_4), v15+int32(288))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L26
L34:
	;
	if v91 == int32(0) {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+22)))
	v97 = v95 + v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+68))
	v99 = F_get_namespace_name_or_temp(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v99 == int32(0) {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	v104 = v15 + int32(376)
	F_initStringInfo(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v107 = F_quote_identifier(m, v99)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+272)) = v107
	F_appendStringInfo(m, v104, int32(_a_F_pg_get_constraintdef_worker_5), v15+int32(272))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v117 = F_quote_identifier(m, v97+int32(4))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_appendStringInfoString(m, v104, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v15)+376))
	F_ReleaseCatCache(m, v91)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v126 = F_quote_identifier(m, v67+int32(4))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v121
	F_appendStringInfo(m, v15+int32(360), int32(_a_F_pg_get_constraintdef_worker_6), v15+int32(256))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L26
L46:
	;
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+73)))
	if v718 == int32(1) {
		goto L194
	} else {
		goto L195
	}
L47:
	;
	F_appendStringInfoString(m, v15+int32(360), int32(_a_F_pg_get_constraintdef_worker_7))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L193
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L190
	}
L49:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v67)+88))
	v615 = F_SysCacheGetAttrNotNull(m, int32(19), v40, int32(27))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L178
	}
L50:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
	if v577 != 0 {
		goto L167
	} else {
		goto L168
	}
L51:
	;
	v470 = F_SysCacheGetAttrNotNull(m, int32(19), v40, int32(28))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L146
	}
L52:
	;
	v294 = v15 + int32(360)
	F_appendStringInfoString(m, v294, v292)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L96
	}
L53:
	;
	v292 = int32(_a_F_pg_get_constraintdef_worker_8)
	goto L52
L54:
	;
	v146 = v15 + int32(360)
	F_appendStringInfoString(m, v146, int32(_a_F_pg_get_constraintdef_worker_9))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v152 = F_SysCacheGetAttrNotNull(m, int32(19), v40, int32(21))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+107)))
	v156 = F_decompile_column_index_array(m, v152, v154, v155, v146)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v67)+96))
	v160 = F_generate_relation_name(m, v158, int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v160
	F_appendStringInfo(m, v146, int32(_a_F_pg_get_constraintdef_worker_10), v15+int32(144))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v170 = F_SysCacheGetAttrNotNull(m, int32(19), v40, int32(22))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v67)+96))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+107)))
	v174 = F_decompile_column_index_array(m, v170, v172, v173, v146)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_appendStringInfoChar(m, v146, int32(41))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+102)))
	switch v180 - int32(102) {
	case 0:
		v201 = int32(_a_F_pg_get_constraintdef_worker_11)
		goto L63
	default:
		goto L65
	case 10:
		goto L64
	case 13:
		goto L66
	}
L63:
	;
	F_appendStringInfoString(m, v15+int32(360), v201)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L70
	}
L64:
	;
	v201 = int32(_a_F_pg_get_constraintdef_worker_12)
	goto L63
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	v201 = int32(_a_F_pg_get_constraintdef_worker_13)
	goto L63
L67:
	;
	v188 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67)+102)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v188
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_14), v15-int32(-64))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(2311), int32(_a_F_pg_get_constraintdef_worker_2))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
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
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+100)))
	switch v207 - int32(97) {
	case 0:
		goto L71
	default:
		goto L74
	case 2:
		goto L73
	case 3:
		goto L75
	case 13:
		goto L76
	case 17:
		v229 = int32(_a_F_pg_get_constraintdef_worker_15)
		goto L72
	}
L71:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+101)))
	switch v240 - int32(97) {
	case 0:
		goto L81
	default:
		goto L84
	case 2:
		goto L83
	case 3:
		goto L85
	case 13:
		goto L86
	case 17:
		v262 = int32(_a_F_pg_get_constraintdef_worker_15)
		goto L82
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v229
	F_appendStringInfo(m, v15+int32(360), int32(_a_F_pg_get_constraintdef_worker_16), v15+int32(128))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L80
	}
L73:
	;
	v229 = int32(_a_F_pg_get_constraintdef_worker_17)
	goto L72
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	v229 = int32(_a_F_pg_get_constraintdef_worker_18)
	goto L72
L76:
	;
	v229 = int32(_a_F_pg_get_constraintdef_worker_19)
	goto L72
L77:
	;
	v216 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67)+100)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v216
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_20), v15+int32(80))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(2337), int32(_a_F_pg_get_constraintdef_worker_2))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	goto L71
L81:
	;
	v276 = F_SysCacheGetAttr(m, int32(19), v40, int32(26), v15+int32(376))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L91
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v262
	F_appendStringInfo(m, v15+int32(360), int32(_a_F_pg_get_constraintdef_worker_21), v15+int32(112))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L90
	}
L83:
	;
	v262 = int32(_a_F_pg_get_constraintdef_worker_17)
	goto L82
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	v262 = int32(_a_F_pg_get_constraintdef_worker_18)
	goto L82
L86:
	;
	v262 = int32(_a_F_pg_get_constraintdef_worker_19)
	goto L82
L87:
	;
	v249 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67)+101)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v249
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_22), v15+int32(96))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(2363), int32(_a_F_pg_get_constraintdef_worker_2))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	goto L81
L91:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+376)))
	if v278 != 0 {
		goto L46
	} else {
		goto L92
	}
L92:
	;
	v280 = v15 + int32(360)
	F_appendStringInfoString(m, v280, int32(_a_F_pg_get_constraintdef_worker_23))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
	v286 = F_decompile_column_index_array(m, v276, v284, int32(0), v280)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_appendStringInfoChar(m, v280, int32(41))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	goto L46
L96:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v67)+88))
	v300 = F_SearchSysCache1(m, int32(34), base.I64_extend_i32_u(v298))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v300 == int32(0) {
		goto L10
	} else {
		goto L98
	}
L98:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+72)))
	if v304 != int32(117) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v318 = v15 + int32(360)
	F_appendStringInfoChar(m, v318, int32(40))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L103
	}
L100:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v300)+16))
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+22)))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307+v308)+13)))
	if v310 != int32(1) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	F_appendStringInfoString(m, v294, int32(_a_F_pg_get_constraintdef_worker_24))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	goto L99
L103:
	;
	v324 = F_SysCacheGetAttrNotNull(m, int32(19), v40, int32(21))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
	v328 = F_decompile_column_index_array(m, v324, v326, int32(0), v318)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+107)))
	if v330 == int32(1) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	F_appendStringInfoString(m, v318, int32(_a_F_pg_get_constraintdef_worker_25))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v337 = v15 + int32(360)
	F_appendStringInfoChar(m, v337, int32(41))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	v343 = F_SysCacheGetAttrNotNull(m, int32(34), v300, int32(3))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v328 < base.I32_wrap_i64(v343) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_appendStringInfoString(m, v337, int32(_a_F_pg_get_constraintdef_worker_26))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	F_ReleaseCatCache(m, v300)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L133
	}
L115:
	;
	v352 = F_SysCacheGetAttrNotNull(m, int32(34), v300, int32(16))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v355 = F_pg_detoast_datum(m, base.I32_wrap_i64(v352))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_deconstruct_array_builtin(m, v355, int32(21), v15+int32(376), int32(0), v15+int32(416))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v15)+416))
	if v328 < v365 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v367 = v328
	goto L122
L120:
	;
	goto L121
L121:
	;
	F_appendStringInfoChar(m, v15+int32(360), int32(41))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L132
	}
L122:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v15)+376))
	v384 = int32(*(*int16)(unsafe.Add(mBase, uint32(v380+v367<<(uint(int32(3))%32)))))
	v386 = F_get_attname(m, v379, v384, int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	goto L121
L124:
	;
	if v328 < v367 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	F_appendStringInfoString(m, v15+int32(360), int32(_a_F_pg_get_constraintdef_worker_27))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v396 = F_quote_identifier(m, v386)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L129
	}
L128:
	;
	goto L127
L129:
	;
	F_appendStringInfoString(m, v15+int32(360), v396)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v401 = v367 + int32(1)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v15)+416))
	if v401 < v402 {
		v367 = v401
		goto L122
	} else {
		goto L131
	}
L131:
	;
	goto L123
L132:
	;
	goto L114
L133:
	;
	v435 = int32(0)
	if base.B2i32(l1 == v435)|base.B2i32(v298 == v435) != 0 {
		goto L46
	} else {
		goto L134
	}
L134:
	;
	v440 = F_flatten_reloptions(m, v298)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	if v440 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v440
	F_appendStringInfo(m, v15+int32(360), int32(_a_F_pg_get_constraintdef_worker_28), v15+int32(240))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v452 = F_get_rel_tablespace(m, v298)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	F_pfree(m, v440)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	if v452 == int32(0) {
		goto L46
	} else {
		goto L142
	}
L142:
	;
	v456 = F_get_tablespace_name(m, v452)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v458 = F_quote_identifier(m, v456)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = v458
	F_appendStringInfo(m, v15+int32(360), int32(_a_F_pg_get_constraintdef_worker_29), v15+int32(224))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	goto L46
L146:
	;
	v473 = F_text_to_cstring(m, base.I32_wrap_i64(v470))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v475 = F_stringToNode(m, v473)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
	if v478 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v479 = F_get_rel_name(m, v478)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L152
	}
L150:
	;
	v533 = int32(0)
	goto L151
L151:
	;
	v538 = v15 + int32(416)
	F_initStringInfo(m, v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L161
	}
L152:
	;
	if v479 == int32(0) {
		goto L9
	} else {
		goto L153
	}
L153:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
	v485 = F_palloc0(m, int32(80))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v488 = F_palloc0(m, int32(136))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v488)+24)) = int32(1)
	v492 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v488)+21)) = uint8(v492)
	*(*int32)(unsafe.Add(mBase, uint32(v488)+16)) = v483
	v495 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v488)+12)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v488))) = int32(101)
	v500 = F_makeAlias(m, v479, v495)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v488)+8)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v488)+4)) = v500
	v504 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v488)+124)) = uint16(v504)
	v506 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v488)+20)) = uint8(v506)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+188)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v15)+376)) = v488
	v513 = F_list_make1_impl(m, int32(1), v15+int32(188))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v515 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v485)+20)) = v515
	*(*int64)(unsafe.Add(mBase, uint32(v485)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v485))) = v513
	F_set_rtable_names(m, v485, v515, v515)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_set_simple_column_names(m, v485)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+184)) = v485
	*(*int32)(unsafe.Add(mBase, uint32(v15)+416)) = v485
	v531 = F_list_make1_impl(m, int32(1), v15+int32(184))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v533 = v531
	goto L151
L161:
	;
	v541 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+408)) = uint8(v541)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+392)) = v541
	v545 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+384)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v15)+380)) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v15)+412)) = v541
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+411)) = uint8(v541)
	v552 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+409)) = uint16(v552)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+400)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v15)+396)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+376)) = v538
	F_get_rule_expr(m, v475, v15+int32(376), v541)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+106)))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v15)+416))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v564
	if v563 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v568 = int32(_a_F_pg_get_constraintdef_worker_30)
	goto L165
L164:
	;
	v568 = int32(_a_F_pg_get_constraintdef_worker_13)
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v568
	F_appendStringInfo(m, v15+int32(360), int32(_a_F_pg_get_constraintdef_worker_31), v15+int32(160))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	goto L46
L167:
	;
	v578 = F_extractNotNullColumn(m, v40)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v67)+84))
	if v603 == int32(0) {
		goto L46
	} else {
		goto L176
	}
L170:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v67)+80))
	v582 = F_get_attname(m, v580, v578, int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v584 = F_quote_identifier(m, v582)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+192)) = v584
	v588 = v15 + int32(360)
	F_appendStringInfo(m, v588, int32(_a_F_pg_get_constraintdef_worker_32), v15+int32(192))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594)+22)))
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594+v595)+106)))
	if v597 != int32(1) {
		goto L46
	} else {
		goto L174
	}
L174:
	;
	F_appendStringInfoString(m, v588, int32(_a_F_pg_get_constraintdef_worker_30))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	goto L46
L176:
	;
	F_appendStringInfoString(m, v15+int32(360), int32(_a_F_pg_get_constraintdef_worker_33))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	goto L46
L178:
	;
	v618 = F_pg_detoast_datum(m, base.I32_wrap_i64(v615))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	F_deconstruct_array_builtin(m, v618, int32(26), v15+int32(376), int32(0), v15+int32(416))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v15)+416))
	v631 = F_palloc(m, v628<<(uint(int32(2))%32))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v15)+416))
	if int32(0) < v633 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v636 = int32(0)
	goto L185
L183:
	;
	goto L184
L184:
	;
	v675 = int32(0)
	v681 = F_pg_get_indexdef_worker(m, v611, v675, v631, v675, v675, v675, v675, l2, v675)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L188
	}
L185:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v15)+376))
	v655 = *(*int64)(unsafe.Add(mBase, uint32(v651+v636<<(uint(int32(3))%32))))
	*(*uint32)(unsafe.Add(mBase, uint32(v631+v636<<(uint(int32(2))%32)))) = uint32(v655)
	v658 = v636 + int32(1)
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v15)+416))
	if v658 < v659 {
		v636 = v658
		goto L185
	} else {
		goto L187
	}
L186:
	;
	goto L184
L187:
	;
	goto L186
L188:
	;
	F_appendStringInfoString(m, v15+int32(360), v681)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	goto L46
L190:
	;
	v689 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67)+72)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v689
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_34), v15+int32(48))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(2594), int32(_a_F_pg_get_constraintdef_worker_2))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
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
	goto L46
L194:
	;
	F_appendStringInfoString(m, v15+int32(360), int32(_a_F_pg_get_constraintdef_worker_35))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L1
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+74)))
	if v726 == int32(1) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	goto L196
L198:
	;
	F_appendStringInfoString(m, v15+int32(360), int32(_a_F_pg_get_constraintdef_worker_36))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+75)))
	if v736 != int32(1) {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	goto L200
L202:
	;
	F_systable_endscan(m, v38)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L208
	}
L203:
	;
	v742 = int32(_a_F_pg_get_constraintdef_worker_37)
	goto L205
L204:
	;
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+76)))
	if v740 != 0 {
		goto L202
	} else {
		goto L206
	}
L205:
	;
	F_appendStringInfoString(m, v15+int32(360), v742)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L207
	}
L206:
	;
	v742 = int32(_a_F_pg_get_constraintdef_worker_38)
	goto L205
L207:
	;
	goto L202
L208:
	;
	F_relation_close(m, v25, int32(1))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v15)+360))
	v763 = v750
	goto L13
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v89
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_39), v15+int32(16))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(_a_F_pg_get_constraintdef_worker_40), int32(_a_F_pg_get_constraintdef_worker_41))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L213:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v97)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v787
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_42), v15+int32(32))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(_a_F_pg_get_constraintdef_worker_43), int32(_a_F_pg_get_constraintdef_worker_41))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v298
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_44), v15+int32(208))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(2403), int32(_a_F_pg_get_constraintdef_worker_2))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v478
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_45), v15+int32(176))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(_a_F_pg_get_constraintdef_worker_46), int32(_a_F_pg_get_constraintdef_worker_47))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_expr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v10 = F_pg_get_expr_worker(m, v4, v8, int32(2))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			if v10 != 0 {
				return base.I64_extend_i32_u(v10)
			} else {
				v14 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
				return int64(0)
			}
		}
	}
}
func F_pg_get_function_arg_default(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v228 int64
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int64
	_ = v277
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v317 int64
	_ = v317
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v20 = F_SearchSysCache1(m, int32(47), v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int64(0)
	} else {
		if v20 == int32(0) {
			v26 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
			v317 = int64(0)
			m.G0 = v15 + int32(80)
			return v317
		} else {
			v35 = F_get_func_arg_info(m, v20, v15+int32(20), v15+int32(16), v15+int32(12))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int64(0)
			} else {
				if base.B2i32(v35 < v17)|base.B2i32(v17 <= int32(0)) != 0 {
					F_ReleaseCatCache(m, v20)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int64(0)
					} else {
						v60 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v60)
						v317 = int64(0)
						m.G0 = v15 + int32(80)
						return v317
					}
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
					if v41 == int32(0) {
						v64 = int32(3)
						v65 = v17 & v64
						v66 = int32(0)
						if base.Ui32(v64) <= base.Ui32(v17-int32(1)) {
							v77 = v66
							v78 = v2
							v79 = int32(0)
							for {
								if v41 != 0 {
									v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v78))))
									v89 = v87 - int32(98)
									if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v89))|base.B2i32(int32(1)<<(uint(v89)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
										v102 = v77
									} else {
										v102 = v77 + int32(1)
									}
								} else {
									v102 = v77 + int32(1)
								}
								if v41 != 0 {
									v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v78)+1)))
									v107 = v105 - int32(98)
									if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v107))|base.B2i32(int32(1)<<(uint(v107)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
										v120 = v102
									} else {
										v120 = v102 + int32(1)
									}
								} else {
									v120 = v102 + int32(1)
								}
								if v41 != 0 {
									v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v78)+2)))
									v125 = v123 - int32(98)
									if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v125))|base.B2i32(int32(1)<<(uint(v125)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
										v138 = v120
									} else {
										v138 = v120 + int32(1)
									}
								} else {
									v138 = v120 + int32(1)
								}
								if v41 != 0 {
									v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v78)+3)))
									v143 = v141 - int32(98)
									if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v143))|base.B2i32(int32(1)<<(uint(v143)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
										v156 = v138
									} else {
										v156 = v138 + int32(1)
									}
								} else {
									v156 = v138 + int32(1)
								}
								v158 = int32(4)
								v159 = v78 + v158
								v161 = v79 + v158
								if v161 != v17&int32(2147483644) {
									v77 = v156
									v78 = v159
									v79 = v161
									continue
								} else {
									break
								}
								break
							}
							if v65 == int32(0) {
								v215 = v156
							} else {
								v168 = v156
								v169 = v159
								v180 = v168
								v181 = v169
								v187 = v2
								for {
									if v41 != 0 {
										v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v181))))
										v192 = v190 - int32(98)
										if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v192))|base.B2i32(int32(1)<<(uint(v192)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
											v205 = v180
										} else {
											v205 = v180 + int32(1)
										}
									} else {
										v205 = v180 + int32(1)
									}
									v207 = int32(1)
									v210 = v187 + v207
									if v210 != v65 {
										v180 = v205
										v181 = v181 + v207
										v187 = v210
										continue
									} else {
										break
									}
									break
								}
								v215 = v205
							}
						} else {
							v168 = v66
							v169 = v2
							v180 = v168
							v181 = v169
							v187 = v2
							for {
								if v41 != 0 {
									v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v181))))
									v192 = v190 - int32(98)
									if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v192))|base.B2i32(int32(1)<<(uint(v192)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
										v205 = v180
									} else {
										v205 = v180 + int32(1)
									}
								} else {
									v205 = v180 + int32(1)
								}
								v207 = int32(1)
								v210 = v187 + v207
								if v210 != v65 {
									v180 = v205
									v181 = v181 + v207
									v187 = v210
									continue
								} else {
									break
								}
								break
							}
							v215 = v205
						}
						v228 = F_SysCacheGetAttr(m, int32(47), v20, int32(24), v15+int32(11))
						mBase = m.M
						v229 = m.ExcPending
						if v229 != 0 {
							return int64(0)
						} else {
							v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+11)))
							if v230 == int32(1) {
								F_ReleaseCatCache(m, v20)
								mBase = m.M
								v234 = m.ExcPending
								if v234 != 0 {
									return int64(0)
								} else {
									v235 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v235)
									v317 = int64(0)
									m.G0 = v15 + int32(80)
									return v317
								}
							} else {
								v239 = F_text_to_cstring(m, base.I32_wrap_i64(v228))
								mBase = m.M
								v240 = m.ExcPending
								if v240 != 0 {
									return int64(0)
								} else {
									v241 = F_stringToNode(m, v239)
									mBase = m.M
									v242 = m.ExcPending
									if v242 != 0 {
										return int64(0)
									} else {
										F_pfree(m, v239)
										mBase = m.M
										v244 = m.ExcPending
										if v244 != 0 {
											return int64(0)
										} else {
											v245 = int32(0)
											v247 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
											v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+22)))
											v249 = v247 + v248
											v250 = int32(*(*int16)(unsafe.Add(mBase, uint32(v249)+106)))
											v251 = int32(*(*int16)(unsafe.Add(mBase, uint32(v249)+104)))
											v255 = v215 + (v250 - v251) - int32(1)
											if base.B2i32(v241 == v245)|base.B2i32(v255 < v245) == v245 {
												v261 = *(*int32)(unsafe.Add(mBase, uint32(v241)+4))
												if v255 < v261 {
													v268 = *(*int32)(unsafe.Add(mBase, uint32(v241)+12))
													v272 = *(*int32)(unsafe.Add(mBase, uint32(v268+v255<<(uint(int32(2))%32))))
													v274 = v15 - int32(-64)
													F_initStringInfo(m, v274)
													mBase = m.M
													v276 = m.ExcPending
													if v276 != 0 {
														return int64(0)
													} else {
														v277 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v15)+28)) = v277
														*(*int64)(unsafe.Add(mBase, uint32(v15)+36)) = v277
														*(*int64)(unsafe.Add(mBase, uint32(v15)+44)) = v277
														*(*int64)(unsafe.Add(mBase, uint32(v15)+49)) = v277
														v285 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v285
														*(*uint8)(unsafe.Add(mBase, uint32(v15)+59)) = uint8(v285)
														v289 = int32(1)
														*(*uint16)(unsafe.Add(mBase, uint32(v15)+57)) = uint16(v289)
														*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v274
														F_get_rule_expr(m, v272, v15+int32(24), v285)
														mBase = m.M
														v296 = m.ExcPending
														if v296 != 0 {
															return int64(0)
														} else {
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
															F_ReleaseCatCache(m, v20)
															mBase = m.M
															v299 = m.ExcPending
															if v299 != 0 {
																return int64(0)
															} else {
																v300 = F_cstring_to_text(m, v297)
																mBase = m.M
																v301 = m.ExcPending
																if v301 != 0 {
																	return int64(0)
																} else {
																	F_pfree(m, v297)
																	mBase = m.M
																	v303 = m.ExcPending
																	if v303 != 0 {
																		return int64(0)
																	} else {
																		v317 = base.I64_extend_i32_u(v300)
																		m.G0 = v15 + int32(80)
																		return v317
																	}
																}
															}
														}
													}
												} else {
													F_ReleaseCatCache(m, v20)
													mBase = m.M
													v264 = m.ExcPending
													if v264 != 0 {
														return int64(0)
													} else {
														v265 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v265)
														v317 = int64(0)
														m.G0 = v15 + int32(80)
														return v317
													}
												}
											} else {
												F_ReleaseCatCache(m, v20)
												mBase = m.M
												v264 = m.ExcPending
												if v264 != 0 {
													return int64(0)
												} else {
													v265 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v265)
													v317 = int64(0)
													m.G0 = v15 + int32(80)
													return v317
												}
											}
										}
									}
								}
							}
						}
					} else {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v17-int32(1)))))
						v49 = v47 - int32(98)
						if base.Ui32(int32(20)) < base.Ui32(v49) {
							F_ReleaseCatCache(m, v20)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int64(0)
							} else {
								v60 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v60)
								v317 = int64(0)
								m.G0 = v15 + int32(80)
								return v317
							}
						} else {
							if int32(1)<<(uint(v49)%32)&int32(_a_F_pg_get_function_arg_default_0) != 0 {
								v64 = int32(3)
								v65 = v17 & v64
								v66 = int32(0)
								if base.Ui32(v64) <= base.Ui32(v17-int32(1)) {
									v77 = v66
									v78 = v2
									v79 = int32(0)
									for {
										if v41 != 0 {
											v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v78))))
											v89 = v87 - int32(98)
											if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v89))|base.B2i32(int32(1)<<(uint(v89)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
												v102 = v77
											} else {
												v102 = v77 + int32(1)
											}
										} else {
											v102 = v77 + int32(1)
										}
										if v41 != 0 {
											v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v78)+1)))
											v107 = v105 - int32(98)
											if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v107))|base.B2i32(int32(1)<<(uint(v107)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
												v120 = v102
											} else {
												v120 = v102 + int32(1)
											}
										} else {
											v120 = v102 + int32(1)
										}
										if v41 != 0 {
											v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v78)+2)))
											v125 = v123 - int32(98)
											if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v125))|base.B2i32(int32(1)<<(uint(v125)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
												v138 = v120
											} else {
												v138 = v120 + int32(1)
											}
										} else {
											v138 = v120 + int32(1)
										}
										if v41 != 0 {
											v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v78)+3)))
											v143 = v141 - int32(98)
											if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v143))|base.B2i32(int32(1)<<(uint(v143)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
												v156 = v138
											} else {
												v156 = v138 + int32(1)
											}
										} else {
											v156 = v138 + int32(1)
										}
										v158 = int32(4)
										v159 = v78 + v158
										v161 = v79 + v158
										if v161 != v17&int32(2147483644) {
											v77 = v156
											v78 = v159
											v79 = v161
											continue
										} else {
											break
										}
										break
									}
									if v65 == int32(0) {
										v215 = v156
									} else {
										v168 = v156
										v169 = v159
										v180 = v168
										v181 = v169
										v187 = v2
										for {
											if v41 != 0 {
												v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v181))))
												v192 = v190 - int32(98)
												if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v192))|base.B2i32(int32(1)<<(uint(v192)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
													v205 = v180
												} else {
													v205 = v180 + int32(1)
												}
											} else {
												v205 = v180 + int32(1)
											}
											v207 = int32(1)
											v210 = v187 + v207
											if v210 != v65 {
												v180 = v205
												v181 = v181 + v207
												v187 = v210
												continue
											} else {
												break
											}
											break
										}
										v215 = v205
									}
								} else {
									v168 = v66
									v169 = v2
									v180 = v168
									v181 = v169
									v187 = v2
									for {
										if v41 != 0 {
											v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v181))))
											v192 = v190 - int32(98)
											if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v192))|base.B2i32(int32(1)<<(uint(v192)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
												v205 = v180
											} else {
												v205 = v180 + int32(1)
											}
										} else {
											v205 = v180 + int32(1)
										}
										v207 = int32(1)
										v210 = v187 + v207
										if v210 != v65 {
											v180 = v205
											v181 = v181 + v207
											v187 = v210
											continue
										} else {
											break
										}
										break
									}
									v215 = v205
								}
								v228 = F_SysCacheGetAttr(m, int32(47), v20, int32(24), v15+int32(11))
								mBase = m.M
								v229 = m.ExcPending
								if v229 != 0 {
									return int64(0)
								} else {
									v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+11)))
									if v230 == int32(1) {
										F_ReleaseCatCache(m, v20)
										mBase = m.M
										v234 = m.ExcPending
										if v234 != 0 {
											return int64(0)
										} else {
											v235 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v235)
											v317 = int64(0)
											m.G0 = v15 + int32(80)
											return v317
										}
									} else {
										v239 = F_text_to_cstring(m, base.I32_wrap_i64(v228))
										mBase = m.M
										v240 = m.ExcPending
										if v240 != 0 {
											return int64(0)
										} else {
											v241 = F_stringToNode(m, v239)
											mBase = m.M
											v242 = m.ExcPending
											if v242 != 0 {
												return int64(0)
											} else {
												F_pfree(m, v239)
												mBase = m.M
												v244 = m.ExcPending
												if v244 != 0 {
													return int64(0)
												} else {
													v245 = int32(0)
													v247 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
													v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+22)))
													v249 = v247 + v248
													v250 = int32(*(*int16)(unsafe.Add(mBase, uint32(v249)+106)))
													v251 = int32(*(*int16)(unsafe.Add(mBase, uint32(v249)+104)))
													v255 = v215 + (v250 - v251) - int32(1)
													if base.B2i32(v241 == v245)|base.B2i32(v255 < v245) == v245 {
														v261 = *(*int32)(unsafe.Add(mBase, uint32(v241)+4))
														if v255 < v261 {
															v268 = *(*int32)(unsafe.Add(mBase, uint32(v241)+12))
															v272 = *(*int32)(unsafe.Add(mBase, uint32(v268+v255<<(uint(int32(2))%32))))
															v274 = v15 - int32(-64)
															F_initStringInfo(m, v274)
															mBase = m.M
															v276 = m.ExcPending
															if v276 != 0 {
																return int64(0)
															} else {
																v277 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(v15)+28)) = v277
																*(*int64)(unsafe.Add(mBase, uint32(v15)+36)) = v277
																*(*int64)(unsafe.Add(mBase, uint32(v15)+44)) = v277
																*(*int64)(unsafe.Add(mBase, uint32(v15)+49)) = v277
																v285 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v285
																*(*uint8)(unsafe.Add(mBase, uint32(v15)+59)) = uint8(v285)
																v289 = int32(1)
																*(*uint16)(unsafe.Add(mBase, uint32(v15)+57)) = uint16(v289)
																*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v274
																F_get_rule_expr(m, v272, v15+int32(24), v285)
																mBase = m.M
																v296 = m.ExcPending
																if v296 != 0 {
																	return int64(0)
																} else {
																	v297 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
																	F_ReleaseCatCache(m, v20)
																	mBase = m.M
																	v299 = m.ExcPending
																	if v299 != 0 {
																		return int64(0)
																	} else {
																		v300 = F_cstring_to_text(m, v297)
																		mBase = m.M
																		v301 = m.ExcPending
																		if v301 != 0 {
																			return int64(0)
																		} else {
																			F_pfree(m, v297)
																			mBase = m.M
																			v303 = m.ExcPending
																			if v303 != 0 {
																				return int64(0)
																			} else {
																				v317 = base.I64_extend_i32_u(v300)
																				m.G0 = v15 + int32(80)
																				return v317
																			}
																		}
																	}
																}
															}
														} else {
															F_ReleaseCatCache(m, v20)
															mBase = m.M
															v264 = m.ExcPending
															if v264 != 0 {
																return int64(0)
															} else {
																v265 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v265)
																v317 = int64(0)
																m.G0 = v15 + int32(80)
																return v317
															}
														}
													} else {
														F_ReleaseCatCache(m, v20)
														mBase = m.M
														v264 = m.ExcPending
														if v264 != 0 {
															return int64(0)
														} else {
															v265 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v265)
															v317 = int64(0)
															m.G0 = v15 + int32(80)
															return v317
														}
													}
												}
											}
										}
									}
								}
							} else {
								F_ReleaseCatCache(m, v20)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int64(0)
								} else {
									v60 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v60)
									v317 = int64(0)
									m.G0 = v15 + int32(80)
									return v317
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_get_function_arguments(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14346(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_pg_get_function_sqlbody(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	F_initStringInfo(m, v7+int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v17 = F_SearchSysCache1(m, int32(47), v9)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			if v17 == int32(0) {
				v21 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
				v49 = int64(0)
				m.G0 = v7 + int32(32)
				return v49
			} else {
				v28 = F_SysCacheGetAttr(m, int32(47), v17, int32(28), v7+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
					if v30 == int32(1) {
						F_ReleaseCatCache(m, v17)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							v35 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
							v49 = int64(0)
							m.G0 = v7 + int32(32)
							return v49
						}
					} else {
						F_print_function_sqlbody(m, v7+int32(16), v17)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int64(0)
						} else {
							F_ReleaseCatCache(m, v17)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int64(0)
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
								v46 = F_cstring_to_text_with_len(m, v44, v45)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int64(0)
								} else {
									v49 = base.I64_extend_i32_u(v46)
									m.G0 = v7 + int32(32)
									return v49
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_get_functiondef(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v232 int32
	_ = v232
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v312 float32
	_ = v312
	var v315 int32
	_ = v315
	var v320 float32
	_ = v320
	var v330 int32
	_ = v330
	var v331 float32
	_ = v331
	var v334 int32
	_ = v334
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v382 int64
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v409 int64
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v599 int32
	_ = v599
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v634 int64
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v655 int64
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v729 int64
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v835 int64
	_ = v835
	v13 = m.G0
	v15 = v13 - int32(192)
	m.G0 = v15
	v17 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	F_initStringInfo(m, v15+int32(160))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v25 = F_SearchSysCache1(m, int32(47), v17)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v15 + int32(192)
	return v835
L4:
	;
	if v25 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
	v835 = int64(0)
	goto L3
L6:
	;
	goto L7
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
	v34 = v32 + v33
	v36 = v34 + int32(4)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+96)))
	if v37 != int32(97) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	F_appendStringInfoChar(m, v15+int32(160), int32(10))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L225
	}
L9:
	;
	v729 = F_SysCacheGetAttrNotNull(m, int32(47), v25, int32(26))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L206
	}
L10:
	;
	v43 = base.B2i32(v37 == int32(112))
	if v37 == int32(112) {
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
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L202
	}
L13:
	;
	v44 = int32(_a_F_pg_get_functiondef_0)
	goto L15
L14:
	;
	v44 = int32(_a_F_pg_get_functiondef_1)
	goto L15
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	v46 = F_get_namespace_name_or_temp(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v49 = v15 + int32(176)
	F_initStringInfo(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v46 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v52 = F_quote_identifier(m, v46)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v62 = F_quote_identifier(m, v36)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v52
	F_appendStringInfo(m, v49, int32(_a_F_pg_get_functiondef_2), v15+int32(144))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	F_appendStringInfoString(m, v15+int32(176), v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v44
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = v67
	v70 = v15 + int32(160)
	F_appendStringInfo(m, v70, int32(_a_F_pg_get_functiondef_3), v15+int32(128))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v78 = F_print_function_arguments(m, v70, v25, int32(0), int32(1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_appendStringInfoString(m, v70, int32(_a_F_pg_get_functiondef_4))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v43 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_appendStringInfoString(m, v70, int32(_a_F_pg_get_functiondef_5))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v96 = m.G0
	v98 = v96 - int32(16)
	m.G0 = v98
	v104 = F_SysCacheGetAttr(m, int32(47), v25, int32(25), v98+int32(15))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L35
	}
L31:
	;
	F_print_function_rettype(m, v70, v25)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_appendStringInfoChar(m, v70, int32(10))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	if int32(0) < v143 {
		goto L52
	} else {
		goto L53
	}
L35:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+15)))
	if v106 != 0 {
		v143 = int32(0)
		goto L37
	} else {
		goto L38
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L49
	}
L37:
	;
	m.G0 = v98 + int32(16)
	goto L34
L38:
	;
	v108 = F_pg_detoast_datum(m, base.I32_wrap_i64(v104))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v110 != int32(1) {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	if v113 < int32(0) {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	if v116 != 0 {
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	if v117 != int32(26) {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	v121 = F_palloc_mul(m, int32(4), v113)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(176)))) = v121
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	if v124 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v134 = (v127<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L47
L46:
	;
	v134 = v124
	goto L47
L47:
	;
	v136 = v113 << (uint(int32(2)) % 32)
	if v136 == int32(0) {
		v143 = v113
		goto L37
	} else {
		goto L48
	}
L48:
	;
	base.MemoryCopy(m, v121, v108+v134, v136)
	v143 = v113
	goto L37
L49:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_functiondef_6), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_pg_get_functiondef_7), int32(1503), int32(_a_F_pg_get_functiondef_8))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v166 = v15 + int32(160)
	F_appendStringInfoString(m, v166, int32(_a_F_pg_get_functiondef_9))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	v247 = F_get_language_name(m, v245, int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L68
	}
L55:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v172 = F_format_type_be(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v172
	F_appendStringInfo(m, v166, int32(_a_F_pg_get_functiondef_10), v15+int32(112))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v180 = int32(1)
	if v143 != v180 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v183 = v180
	goto L61
L59:
	;
	goto L60
L60:
	;
	F_appendStringInfoChar(m, v15+int32(160), int32(10))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L67
	}
L61:
	;
	v196 = v15 + int32(160)
	F_appendStringInfoString(m, v196, int32(_a_F_pg_get_functiondef_11))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	goto L60
L63:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200+v183<<(uint(int32(2))%32))))
	v205 = F_format_type_be(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v205
	F_appendStringInfo(m, v196, int32(_a_F_pg_get_functiondef_10), v15+int32(96))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v214 = v183 + int32(1)
	if v214 != v143 {
		v183 = v214
		goto L61
	} else {
		goto L66
	}
L66:
	;
	goto L62
L67:
	;
	goto L54
L68:
	;
	v249 = F_quote_identifier(m, v247)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v249
	v253 = v15 + int32(160)
	F_appendStringInfo(m, v253, int32(_a_F_pg_get_functiondef_12), v15+int32(80))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+96)))
	if v260 == int32(119) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_appendStringInfoString(m, v253, int32(_a_F_pg_get_functiondef_13))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+101)))
	switch v267 - int32(105) {
	case 0:
		v271 = int32(_a_F_pg_get_functiondef_14)
		goto L76
	default:
		goto L75
	case 10:
		goto L77
	}
L74:
	;
	goto L73
L75:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+102)))
	switch v278 - int32(114) {
	case 0:
		goto L81
	case 1:
		v282 = int32(_a_F_pg_get_functiondef_15)
		goto L80
	default:
		goto L79
	}
L76:
	;
	F_appendStringInfoString(m, v15+int32(160), v271)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	v271 = int32(_a_F_pg_get_functiondef_16)
	goto L76
L78:
	;
	goto L75
L79:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+99)))
	if v288 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	F_appendStringInfoString(m, v15+int32(160), v282)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L82
	}
L81:
	;
	v282 = int32(_a_F_pg_get_functiondef_17)
	goto L80
L82:
	;
	goto L79
L83:
	;
	F_appendStringInfoString(m, v15+int32(160), int32(_a_F_pg_get_functiondef_18))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+97)))
	if v296 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	F_appendStringInfoString(m, v15+int32(160), int32(_a_F_pg_get_functiondef_19))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+98)))
	if v304 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L89
L91:
	;
	F_appendStringInfoString(m, v15+int32(160), int32(_a_F_pg_get_functiondef_20))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v312 = *(*float32)(unsafe.Add(mBase, uint32(v34)+80))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	if v315&int32(-2) == int32(12) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L93
L95:
	;
	v320 = float32(1)
	goto L97
L96:
	;
	v320 = float32(100)
	goto L97
L97:
	;
	if base.F32_ne(v312, v320) != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v15)+64)) = base.F64_promote_f32(v312)
	F_appendStringInfo(m, v15+int32(160), int32(_a_F_pg_get_functiondef_21), v15-int32(-64))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v331 = *(*float32)(unsafe.Add(mBase, uint32(v34)+84))
	v334 = int32(0)
	if base.B2i32(base.F32_gt(v331, float32(0)) == v334)|base.F32_eq(v331, float32(1000)) == v334 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L100
L102:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = base.F64_promote_f32(v331)
	F_appendStringInfo(m, v15+int32(160), int32(_a_F_pg_get_functiondef_22), v15+int32(48))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v34)+92))
	if v350 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L104
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = int32(2281)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v34)+92))
	v355 = int32(0)
	v361 = F_generate_function_name(m, v353, int32(1), v355, v15+int32(176), v355, v355, v355)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
	if v371 != v259 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v361
	F_appendStringInfo(m, v15+int32(160), int32(_a_F_pg_get_functiondef_23), v15+int32(32))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	F_appendStringInfoChar(m, v15+int32(160), int32(10))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v382 = F_SysCacheGetAttr(m, int32(47), v25, int32(29), v15+int32(159))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+159)))
	if v384 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v634 = F_SysCacheGetAttr(m, int32(47), v25, int32(28), v15+int32(159))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L181
	}
L117:
	;
	v386 = F_pg_detoast_datum(m, base.I32_wrap_i64(v382))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = int32(1)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v386)+16))
	if v390 <= int32(0) {
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L120
L120:
	;
	v409 = F_array_ref(m, v386, v15+int32(176), v15+int32(159))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L122
	}
L121:
	;
	goto L116
L122:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+159)))
	if v411 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	v614 = v612 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v614
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v386)+16))
	if v614 <= v616 {
		goto L120
	} else {
		goto L180
	}
L124:
	;
	v413 = F_text_to_cstring(m, base.I32_wrap_i64(v409))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v415 = int32(61)
	v416 = F___strchrnul(m, v413, v415)
	mBase = m.M
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416))))
	if v418 == v415 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	if v422 == int32(0) {
		goto L123
	} else {
		goto L130
	}
L127:
	;
	v422 = v416
	goto L129
L128:
	;
	v422 = int32(0)
	goto L129
L129:
	;
	goto L126
L130:
	;
	v425 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v422))) = uint8(v425)
	v427 = F_quote_identifier(m, v413)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v427
	v431 = v15 + int32(160)
	F_appendStringInfo(m, v431, int32(_a_F_pg_get_functiondef_24), v15+int32(16))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v438 = v422 + int32(1)
	v439 = F_GetConfigOptionFlags(m, v413)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	F_appendStringInfoChar(m, v15+int32(160), int32(10))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L179
	}
L134:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L176
	}
L135:
	;
	if v439&int32(2) != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v445 = F_SplitGUCList(m, v438, v15+int32(152))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	F_appendStringInfoChar(m, v15+int32(160), int32(39))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L165
	}
L139:
	;
	if v445 == int32(0) {
		goto L134
	} else {
		goto L140
	}
L140:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v15)+152))
	if v449 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	F_appendStringInfoString(m, v431, int32(_a_F_pg_get_functiondef_25))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	v458 = v449
	goto L143
L143:
	;
	v459 = int32(0)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v458)+4))
	if v460 <= v459 {
		goto L133
	} else {
		goto L146
	}
L144:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v15)+152))
	if v455 == int32(0) {
		goto L133
	} else {
		goto L145
	}
L145:
	;
	v458 = v455
	goto L143
L146:
	;
	v465 = v459
	goto L147
L147:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v458)+12))
	v478 = v475 + v465<<(uint(int32(2))%32)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)))
	F_appendStringInfoChar(m, v15+int32(160), int32(39))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v485 = v479
	goto L150
L150:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	if v497 != int32(39) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	F_appendStringInfoChar(m, v15+int32(160), base.I32_extend8_s(v497))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L164
	}
L153:
	;
	if v497 != 0 {
		goto L152
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	F_appendStringInfoChar(m, v15+int32(160), int32(39))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L163
	}
L156:
	;
	v501 = v15 + int32(160)
	F_appendStringInfoChar(m, v501, int32(39))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v15)+152))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+12))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	if base.Ui32(v478+int32(4)) < base.Ui32(v508+v509<<(uint(int32(2))%32)) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	F_appendStringInfoString(m, v501, int32(_a_F_pg_get_functiondef_11))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v518 = v465 + int32(1)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v458)+4))
	if v518 < v519 {
		v465 = v518
		goto L147
	} else {
		goto L162
	}
L161:
	;
	goto L160
L162:
	;
	goto L133
L163:
	;
	goto L152
L164:
	;
	v485 = v485 + int32(1)
	goto L150
L165:
	;
	v538 = v438
	goto L166
L166:
	;
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538))))
	if v550 != int32(39) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	F_appendStringInfoChar(m, v15+int32(160), base.I32_extend8_s(v550))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L175
	}
L169:
	;
	if v550 != 0 {
		goto L168
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	F_appendStringInfoChar(m, v15+int32(160), int32(39))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L174
	}
L172:
	;
	F_appendStringInfoChar(m, v15+int32(160), int32(39))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	goto L133
L174:
	;
	goto L168
L175:
	;
	v538 = v538 + int32(1)
	goto L166
L176:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_functiondef_26), int32(0))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_pg_get_functiondef_27), int32(3110), int32(_a_F_pg_get_functiondef_28))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	goto L123
L180:
	;
	goto L121
L181:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	if v636 != int32(14) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v647 = v15 + int32(160)
	F_appendStringInfoString(m, v647, int32(_a_F_pg_get_functiondef_29))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L186
	}
L183:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+159)))
	if v639&int32(1) != 0 {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	F_print_function_sqlbody(m, v15+int32(160), v25)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	goto L8
L186:
	;
	v655 = F_SysCacheGetAttr(m, int32(47), v25, int32(27), v15+int32(159))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+159)))
	if v657 != 0 {
		goto L9
	} else {
		goto L188
	}
L188:
	;
	v659 = F_text_to_cstring(m, base.I32_wrap_i64(v655))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	F_appendStringInfoChar(m, v647, int32(39))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	v664 = v659
	goto L191
L191:
	;
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664))))
	if v676 != int32(39) {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	F_appendStringInfoChar(m, v15+int32(160), base.I32_extend8_s(v676))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L201
	}
L194:
	;
	if v676 != 0 {
		goto L193
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	F_appendStringInfoChar(m, v15+int32(160), int32(39))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L200
	}
L197:
	;
	v680 = v15 + int32(160)
	F_appendStringInfoChar(m, v680, int32(39))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_appendStringInfoString(m, v680, int32(_a_F_pg_get_functiondef_11))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	goto L9
L200:
	;
	goto L193
L201:
	;
	v664 = v664 + int32(1)
	goto L191
L202:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v36
	F_errmsg(m, int32(_a_F_pg_get_functiondef_30), v15)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_pg_get_functiondef_27), int32(2958), int32(_a_F_pg_get_functiondef_28))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L206:
	;
	v732 = F_text_to_cstring(m, base.I32_wrap_i64(v729))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	v735 = v15 + int32(176)
	F_initStringInfo(m, v735)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	F_appendStringInfoChar(m, v735, int32(36))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	if v37 == int32(112) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v745 = int32(_a_F_pg_get_functiondef_31)
	goto L212
L211:
	;
	v745 = int32(_a_F_pg_get_functiondef_32)
	goto L212
L212:
	;
	F_appendStringInfoString(m, v735, v745)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	v749 = F_strstr(m, v732, v748)
	mBase = m.M
	if v749 != 0 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	goto L217
L215:
	;
	goto L216
L216:
	;
	F_appendStringInfoChar(m, v15+int32(176), int32(36))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L221
	}
L217:
	;
	F_appendStringInfoChar(m, v15+int32(176), int32(120))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L219
	}
L218:
	;
	goto L216
L219:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	v768 = F_strstr(m, v732, v767)
	mBase = m.M
	if v768 != 0 {
		goto L217
	} else {
		goto L220
	}
L220:
	;
	goto L218
L221:
	;
	v787 = v15 + int32(160)
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v15)+180))
	F_appendBinaryStringInfo(m, v787, v788, v789)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	F_appendStringInfoString(m, v787, v732)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v15)+180))
	F_appendBinaryStringInfo(m, v787, v794, v795)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	goto L8
L225:
	;
	F_ReleaseCatCache(m, v25)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
	v818 = F_cstring_to_text(m, v817)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	F_pfree(m, v817)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v835 = base.I64_extend_i32_u(v818)
	goto L3
}
func F_pg_get_indexdef(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_get_indexdef_worker(m, v3, v2, v2, v2, v2, v2, v2, int32(2), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			return int64(0)
		} else {
			v22 = F_cstring_to_text(m, v12)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				F_pfree(m, v12)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v22)
				}
			}
		}
	}
}
func F_pg_get_loaded_modules(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_loaded_modules[0]))
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v17 = v16
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v7 - int32(-64)
	return int64(0)
L6:
	;
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v21
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v27)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v27)
	*(*int32)(unsafe.Add(mBase, uint32(v5+int32(-4)))) = v17 + int32(24)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v5+int32(-8)))) = v39
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v5+int32(-12)))) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
	if v46 == v27 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v7)+52))
	if v55 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v49 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v49)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v51 = F_cstring_to_text(m, v46)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = base.I64_extend_i32_u(v51)
	goto L8
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+60))
	v67 = v64
	v69 = int32(0)
	goto L20
L14:
	;
	v58 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v58)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v60 = F_cstring_to_text(m, v55)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = base.I64_extend_i32_u(v60)
	goto L13
L18:
	;
	v83 = F_cstring_to_text(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L30
	}
L19:
	;
	if v69 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v70 != int32(47) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v67 = v67 + int32(1)
	v69 = v73
	goto L20
L23:
	;
	if v70 != 0 {
		v73 = v69
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v73 = v67
	goto L22
L26:
	;
	goto L19
L27:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v7)+60))
	v82 = v78
	goto L18
L28:
	;
	goto L29
L29:
	;
	v80 = v69 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+60)) = v80
	v82 = v80
	goto L18
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = base.I64_extend_i32_u(v83)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	F_tuplestore_putvalues(m, v87, v88, v5+int32(-48), v5+int32(-52))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v95 != 0 {
		v17 = v95
		goto L6
	} else {
		goto L32
	}
L32:
	;
	goto L7
}
func F_pg_get_partition_constraintdef(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
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
	var v112 int64
	_ = v112
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_get_partition_qual_relid(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		if v13 == int32(0) {
			v19 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
			v112 = int64(0)
			m.G0 = v10 + int32(80)
			return v112
		} else {
			v22 = F_get_rel_name(m, v12)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				if v22 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v12
						F_errmsg_internal(m, int32(_a_F_pg_get_partition_constraintdef_0), v10)
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_get_partition_constraintdef_1), int32(_a_F_pg_get_partition_constraintdef_2), int32(_a_F_pg_get_partition_constraintdef_3))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v27 = F_palloc0(m, int32(80))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						v30 = F_palloc0(m, int32(136))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = int32(1)
							v34 = int32(114)
							*(*uint8)(unsafe.Add(mBase, uint32(v30)+21)) = uint8(v34)
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v12
							v37 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v37
							*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(101)
							v42 = F_makeAlias(m, v22, v37)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v42
								*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v42
								v46 = int32(256)
								*(*uint16)(unsafe.Add(mBase, uint32(v30)+124)) = uint16(v46)
								v48 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v30)+20)) = uint8(v48)
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v30
								*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v30
								v55 = F_list_make1_impl(m, int32(1), v10+int32(20))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int64(0)
								} else {
									v57 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v57
									*(*int64)(unsafe.Add(mBase, uint32(v27)+12)) = int64(0)
									*(*int32)(unsafe.Add(mBase, uint32(v27))) = v55
									F_set_rtable_names(m, v27, v57, v57)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int64(0)
									} else {
										F_set_simple_column_names(m, v27)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v27
											*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v27
											v73 = F_list_make1_impl(m, int32(1), v10+int32(16))
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int64(0)
											} else {
												v76 = v10 - int32(-64)
												F_initStringInfo(m, v76)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int64(0)
												} else {
													v79 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+56)) = uint8(v79)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v79
													*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v73
													*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v79
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+59)) = uint8(v79)
													v90 = int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(v10)+57)) = uint16(v90)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v79
													*(*int64)(unsafe.Add(mBase, uint32(v10)+44)) = int64(2)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v76
													F_get_rule_expr(m, v13, v10+int32(24), v79)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return int64(0)
													} else {
														v102 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
														v103 = F_cstring_to_text(m, v102)
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return int64(0)
														} else {
															F_pfree(m, v102)
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return int64(0)
															} else {
																v112 = base.I64_extend_i32_u(v103)
																m.G0 = v10 + int32(80)
																return v112
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
func F_pg_get_statisticsobjdef_expressions(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v144 int32
	_ = v144
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
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v189 int64
	_ = v189
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v15 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v16 = F_SearchSysCache1(m, int32(64), v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L3
	} else {
		goto L38
	}
L2:
	;
	m.G0 = v12 + int32(80)
	return v189
L3:
	;
	return int64(0)
L4:
	;
	if v16 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
	v189 = int64(0)
	goto L2
L6:
	;
	goto L7
L7:
	;
	v27 = F_heap_attisnull(m, v16, int32(9), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v27 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_ReleaseCatCache(m, v16)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
	v38 = F_SysCacheGetAttrNotNull(m, int32(64), v16, int32(9))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L13
	}
L12:
	;
	v31 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
	v189 = int64(0)
	goto L2
L13:
	;
	v41 = F_text_to_cstring(m, base.I32_wrap_i64(v38))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v43 = F_stringToNode(m, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	F_pfree(m, v41)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v47 = v34 + v35
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v49 = F_get_rel_name(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	if v49 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v55 = F_palloc0(m, int32(80))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v58 = F_palloc0(m, int32(136))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = int32(1)
	v62 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+21)) = uint8(v62)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v53
	v65 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(101)
	v70 = F_makeAlias(m, v49, v65)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v70
	v74 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v58)+124)) = uint16(v74)
	v76 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)) = uint8(v76)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v58
	v83 = F_list_make1_impl(m, int32(1), v12+int32(20))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v85 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+20)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v55)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v83
	F_set_rtable_names(m, v55, v85, v85)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	F_set_simple_column_names(m, v55)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v55
	v101 = F_list_make1_impl(m, int32(1), v12+int32(16))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v103 = int32(0)
	if v43 == v103 {
		v167 = v103
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_ReleaseCatCache(m, v16)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L3
	} else {
		goto L36
	}
L27:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v106 <= int32(0) {
		v167 = v103
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v110 = int32(0)
	v112 = v103
	goto L29
L29:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119+v110<<(uint(int32(2))%32))))
	v125 = v12 - int32(-64)
	F_initStringInfo(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L3
	} else {
		goto L31
	}
L30:
	;
	v167 = v159
	goto L26
L31:
	;
	v128 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+56)) = uint8(v128)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v128
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+59)) = uint8(v128)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v12)+44)) = int64(2)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v125
	v144 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+57)) = uint16(v144)
	F_get_rule_expr(m, v123, v12+int32(24), v128)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v152 = F_cstring_to_text(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_statisticsobjdef_expressions[0]))
	v159 = F_accumArrayResult(m, v112, base.I64_extend_i32_u(v152), int32(0), int32(25), v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v162 = v110 + int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v162 < v163 {
		v110 = v162
		v112 = v159
		goto L29
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_statisticsobjdef_expressions[0]))
	v178 = F_makeArrayResult(m, v167, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v189 = v178
	goto L2
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v48
	F_errmsg_internal(m, int32(_a_F_pg_get_statisticsobjdef_expressions_0), v12)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_pg_get_statisticsobjdef_expressions_1), int32(_a_F_pg_get_statisticsobjdef_expressions_2), int32(_a_F_pg_get_statisticsobjdef_expressions_3))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_identify_object(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int64
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int64
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int64
	_ = v355
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v451 int32
	_ = v451
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(128)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v14)+120)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v14)+116)) = v16
	v25 = F_get_call_result_type(m, l0, v2, v14+int32(72))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errfinish(m, int32(_a_F_pg_identify_object_0), int32(2827), int32(_a_F_pg_identify_object_1))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L7
	} else {
		goto L123
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L7
	} else {
		goto L120
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L7
	} else {
		goto L118
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L7
	} else {
		goto L115
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L7
	} else {
		goto L113
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L7
	} else {
		goto L111
	}
L7:
	;
	return int64(0)
L8:
	;
	if v25 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v33 = v2
	goto L12
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L7
	} else {
		goto L108
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v33*int32(40))+uint32(_c_F_pg_identify_object[0])))
	if v16 != v44 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v44 == v16 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v47 = v33 + int32(1)
	if v47 != int32(37) {
		v33 = v47
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	goto L13
L17:
	;
	goto L16
L18:
	;
	v54 = F_table_open(m, v16, int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	v279 = v2
	v284 = v2
	goto L20
L20:
	;
	v288 = v14 + int32(116)
	v290 = F_getObjectTypeDescription(m, v288, int32(1))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L7
	} else {
		goto L89
	}
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_pg_identify_object[1]))
	if v57 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v110 = int32(0)
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v101)+20)))
	v113 = F_get_catalog_object_by_oid_extended(m, v54, v111, v17, v110)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L7
	} else {
		goto L41
	}
L23:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v58 == v16 {
		v101 = v57
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v61 = int32(0)
	goto L31
L26:
	;
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_identify_object[1])) = v97
	v101 = v97
	goto L22
L28:
	;
	v97 = v73 + int32(_a_F_pg_identify_object_2)
	goto L27
L29:
	;
	v97 = v73 + int32(_a_F_pg_identify_object_3)
	goto L27
L30:
	;
	v97 = v73 + int32(_a_F_pg_identify_object_4)
	goto L27
L31:
	;
	v73 = v61 * int32(40)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_pg_identify_object[0])))
	if v16 != v74 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v97 = v73 + int32(_a_F_pg_identify_object_5)
	goto L27
L33:
	;
	if v61 == int32(36) {
		goto L6
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_pg_identify_object[2])))
	if v80 == v16 {
		goto L28
	} else {
		goto L37
	}
L37:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_pg_identify_object[3])))
	if v82 == v16 {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_pg_identify_object[4])))
	if v84 == v16 {
		goto L30
	} else {
		goto L39
	}
L39:
	;
	v61 = v61 + int32(4)
	goto L31
L40:
	;
	F_relation_close(m, v54, int32(1))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L7
	} else {
		goto L88
	}
L41:
	;
	if v113 == int32(0) {
		v265 = v110
		v270 = v2
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_pg_identify_object[1]))
	if v119 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v172 = int32(*(*int16)(unsafe.Add(mBase, uint32(v163)+24)))
	if v172 != 0 {
		goto L63
	} else {
		goto L64
	}
L44:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v120 == v16 {
		v163 = v119
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v123 = int32(0)
	goto L52
L47:
	;
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_identify_object[1])) = v159
	v163 = v159
	goto L43
L49:
	;
	v159 = v135 + int32(_a_F_pg_identify_object_2)
	goto L48
L50:
	;
	v159 = v135 + int32(_a_F_pg_identify_object_3)
	goto L48
L51:
	;
	v159 = v135 + int32(_a_F_pg_identify_object_4)
	goto L48
L52:
	;
	v135 = v123 * int32(40)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+uint32(_c_F_pg_identify_object[0])))
	if v16 != v136 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v159 = v135 + int32(_a_F_pg_identify_object_5)
	goto L48
L54:
	;
	if v123 == int32(36) {
		goto L5
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	goto L53
L57:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v135)+uint32(_c_F_pg_identify_object[2])))
	if v142 == v16 {
		goto L49
	} else {
		goto L58
	}
L58:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v135)+uint32(_c_F_pg_identify_object[3])))
	if v144 == v16 {
		goto L50
	} else {
		goto L59
	}
L59:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v135)+uint32(_c_F_pg_identify_object[4])))
	if v146 == v16 {
		goto L51
	} else {
		goto L60
	}
L60:
	;
	v123 = v123 + int32(4)
	goto L52
L61:
	;
	v243 = int32(0)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+36)))
	if v244 != int32(1) {
		v265 = v243
		v270 = v241
		goto L40
	} else {
		goto L83
	}
L62:
	;
	v194 = int32(0)
	goto L74
L63:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	v176 = F_heap_getattr_2(m, v113, v172, v173, v14+int32(80))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L7
	} else {
		goto L66
	}
L64:
	;
	v186 = v163
	v187 = v2
	goto L65
L65:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v188 == v16 {
		v234 = v186
		v241 = v187
		goto L61
	} else {
		goto L69
	}
L66:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+80)))
	if v178 == int32(1) {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	v181 = base.I32_wrap_i64(v176)
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_pg_identify_object[1]))
	if v183 == int32(0) {
		v191 = v181
		goto L62
	} else {
		goto L68
	}
L68:
	;
	v186 = v183
	v187 = v181
	goto L65
L69:
	;
	v191 = v187
	goto L62
L70:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_identify_object[1])) = v230
	v234 = v230
	v241 = v191
	goto L61
L71:
	;
	v230 = v206 + int32(_a_F_pg_identify_object_2)
	goto L70
L72:
	;
	v230 = v206 + int32(_a_F_pg_identify_object_3)
	goto L70
L73:
	;
	v230 = v206 + int32(_a_F_pg_identify_object_4)
	goto L70
L74:
	;
	v206 = v194 * int32(40)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+uint32(_c_F_pg_identify_object[0])))
	if v16 != v207 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v230 = v206 + int32(_a_F_pg_identify_object_5)
	goto L70
L76:
	;
	if v194 == int32(36) {
		goto L3
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	goto L75
L79:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v206)+uint32(_c_F_pg_identify_object[2])))
	if v213 == v16 {
		goto L71
	} else {
		goto L80
	}
L80:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v206)+uint32(_c_F_pg_identify_object[3])))
	if v215 == v16 {
		goto L72
	} else {
		goto L81
	}
L81:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v206)+uint32(_c_F_pg_identify_object[4])))
	if v217 == v16 {
		goto L73
	} else {
		goto L82
	}
L82:
	;
	v194 = v194 + int32(4)
	goto L74
L83:
	;
	v247 = int32(*(*int16)(unsafe.Add(mBase, uint32(v234)+22)))
	if v247 == int32(0) {
		v265 = v243
		v270 = v241
		goto L40
	} else {
		goto L84
	}
L84:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	v253 = F_heap_getattr_2(m, v113, v247, v250, v14+int32(80))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+80)))
	if v255 == int32(1) {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	v259 = F_quote_identifier(m, base.I32_wrap_i64(v253))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	v265 = v259
	v270 = v241
	goto L40
L88:
	;
	v279 = v265
	v284 = v270
	goto L20
L89:
	;
	v292 = F_cstring_to_text(m, v290)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	v294 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v294)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = base.I64_extend_i32_u(v292)
	v303 = F_getObjectIdentityParts(m, v288, v294, v294, int32(1))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L7
	} else {
		goto L91
	}
L91:
	;
	if base.B2i32(v284 == v294)|base.B2i32(v303 == int32(0)) != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v318 = int32(1)
	goto L94
L93:
	;
	v309 = F_get_namespace_name(m, v284)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L7
	} else {
		goto L95
	}
L94:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+77)) = uint8(v318)
	v320 = int32(0)
	if base.B2i32(v279 == v320)|base.B2i32(v303 == v320) == v320 {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v311 = F_quote_identifier(m, v309)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	v313 = F_cstring_to_text(m, v311)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = base.I64_extend_i32_u(v313)
	v318 = int32(0)
	goto L94
L98:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+79)) = uint8(v345)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v352 = F_heap_form_tuple(m, v347, v14+int32(80), v14+int32(76))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L7
	} else {
		goto L106
	}
L99:
	;
	v339 = F_cstring_to_text(m, v303)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L7
	} else {
		goto L105
	}
L100:
	;
	v327 = F_cstring_to_text(m, v279)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L7
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v333 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+78)) = uint8(v333)
	if v303 == int32(0) {
		v345 = v333
		goto L98
	} else {
		goto L104
	}
L103:
	;
	v329 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+78)) = uint8(v329)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = base.I64_extend_i32_u(v327)
	goto L99
L104:
	;
	goto L99
L105:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+104)) = base.I64_extend_i32_u(v339)
	v345 = int32(0)
	goto L98
L106:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v352)+16))
	v355 = F_HeapTupleHeaderGetDatum(m, v354)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	m.G0 = v14 + int32(128)
	return v355
L108:
	;
	F_errmsg_internal(m, int32(_a_F_pg_identify_object_6), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(_a_F_pg_identify_object_0), int32(_a_F_pg_identify_object_7), int32(_a_F_pg_identify_object_8))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L7
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v16
	F_errmsg_internal(m, int32(_a_F_pg_identify_object_9), v14-int32(-64))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	goto L1
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v16
	F_errmsg_internal(m, int32(_a_F_pg_identify_object_9), v14+int32(48))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	goto L1
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v16
	F_errmsg_internal(m, int32(_a_F_pg_identify_object_10), v14+int32(32))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L7
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_pg_identify_object_0), int32(_a_F_pg_identify_object_11), int32(_a_F_pg_identify_object_8))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L7
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v16
	F_errmsg_internal(m, int32(_a_F_pg_identify_object_9), v14+int32(16))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	goto L1
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v16
	F_errmsg_internal(m, int32(_a_F_pg_identify_object_12), v14)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_pg_identify_object_0), int32(_a_F_pg_identify_object_13), int32(_a_F_pg_identify_object_8))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L7
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_indexes_size(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_try_relation_open(m, v4, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		if v6 == int32(0) {
			v12 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
			return int64(0)
		} else {
			v16 = F_calculate_indexes_size(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				F_relation_close(m, v6, int32(1))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					return v16
				}
			}
		}
	}
}
func F_pg_inet_net_ntop(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
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
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v107 int64
	_ = v107
	var v116 int32
	_ = v116
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
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
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
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
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v630 int32
	_ = v630
	v14 = m.G0
	v16 = v14 - int32(288)
	m.G0 = v16
	switch l0 - int32(2) {
	case 0:
		goto L6
	case 1, 8:
		goto L5
	default:
		goto L4
	}
L1:
	;
	m.G0 = v16 + int32(288)
	return v630
L2:
	;
	v630 = int32(0)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_inet_net_ntop[0])) = int32(28)
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_inet_net_ntop[0])) = int32(5)
	goto L2
L5:
	;
	if base.Ui32(l2-int32(129)) <= base.Ui32(int32(-131)) {
		goto L29
	} else {
		goto L30
	}
L6:
	;
	if base.Ui32(int32(32)) < base.Ui32(l2) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v22
	v25 = l3 + int32(50)
	v29 = F_pg_sprintf(m, l3, int32(_a_F_pg_inet_net_ntop_0), v16-int32(-64))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_inet_net_ntop[0])) = int32(35)
	goto L2
L9:
	;
	return int32(0)
L10:
	;
	v33 = l3 + v29
	if base.Ui32(v25-v33) < base.Ui32(int32(6)) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v29 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v37 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v37)
	v41 = v33 + int32(1)
	goto L14
L13:
	;
	v41 = l3
	goto L14
L14:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v42
	v47 = F_pg_sprintf(m, v41, int32(_a_F_pg_inet_net_ntop_0), v16+int32(48))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v49 = v47 + v41
	if base.Ui32(v25-v49) < base.Ui32(int32(6)) {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	if v49 != l3 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v54 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v54)
	v58 = v49 + int32(1)
	goto L19
L18:
	;
	v58 = l3
	goto L19
L19:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v59
	v64 = F_pg_sprintf(m, v58, int32(_a_F_pg_inet_net_ntop_0), v16+int32(32))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v66 = v64 + v58
	if base.Ui32(v25-v66) < base.Ui32(int32(6)) {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	if v66 != l3 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v71 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v71)
	v75 = v66 + int32(1)
	goto L24
L23:
	;
	v75 = l3
	goto L24
L24:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v76
	v81 = F_pg_sprintf(m, v75, int32(_a_F_pg_inet_net_ntop_0), v16+int32(16))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	if l2 == int32(32) {
		v630 = l3
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v85 = v81 + v75
	if base.Ui32(v25-v85) < base.Ui32(int32(5)) {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
	v91 = F_pg_sprintf(m, v85, int32(_a_F_pg_inet_net_ntop_1), v16)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v630 = l3
	goto L1
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_inet_net_ntop[0])) = int32(28)
	goto L2
L30:
	;
	goto L31
L31:
	;
	v107 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+216)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v16)+208)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v16)+200)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v16)+192)) = v107
	v116 = int32(0)
	goto L32
L32:
	;
	v133 = v16 + int32(192) + v116<<(uint(int32(1))%32)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v135 = v116 + l1
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	v139 = v134 | v136<<(uint(int32(8))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v139
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v139 | v141
	v145 = v116 + int32(2)
	if v145 != int32(16) {
		v116 = v145
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v16)+192))
	v150 = base.B2i32(v148 != int32(0))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v16)+196))
	if v151 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L33
L35:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v16)+200))
	if v168 != 0 {
		goto L46
	} else {
		goto L47
	}
L36:
	;
	v152 = int32(-1)
	v153 = int32(0)
	v154 = base.B2i32(v148 == v153)
	if v148 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if v148 != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v157 = v152
	goto L41
L40:
	;
	v157 = v153
	goto L41
L41:
	;
	v163 = v152
	v164 = v154
	v165 = v154
	v166 = v150
	v167 = v157
	goto L35
L42:
	;
	v160 = int32(1)
	goto L44
L43:
	;
	v160 = int32(2)
	goto L44
L44:
	;
	v163 = v150
	v164 = v160
	v165 = int32(0)
	v166 = int32(1)
	v167 = int32(-1)
	goto L35
L45:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v16)+204))
	if v190 != 0 {
		goto L63
	} else {
		goto L64
	}
L46:
	;
	v169 = int32(-1)
	if v163 == v169 {
		v184 = v167
		v186 = v164
		v187 = v165
		v189 = v169
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v176 = int32(1)
	v180 = base.B2i32(v163 == int32(-1))
	if v163 == int32(-1) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	v173 = v166 | base.B2i32(base.Ui32(v165) < base.Ui32(v164))
	if v173 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v174 = v163
	goto L52
L51:
	;
	v174 = v167
	goto L52
L52:
	;
	if v173 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v175 = v164
	goto L55
L54:
	;
	v175 = v165
	goto L55
L55:
	;
	v184 = v174
	v186 = v164
	v187 = v175
	v189 = v169
	goto L45
L56:
	;
	v181 = v176
	goto L58
L57:
	;
	v181 = v164 + v176
	goto L58
L58:
	;
	if v163 == int32(-1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v183 = int32(2)
	goto L61
L60:
	;
	v183 = v163
	goto L61
L61:
	;
	v184 = v167
	v186 = v181
	v187 = v165
	v189 = v183
	goto L45
L62:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v16)+208))
	if v214 != 0 {
		goto L80
	} else {
		goto L81
	}
L63:
	;
	v191 = int32(-1)
	if v189 == v191 {
		v209 = v184
		v211 = v186
		v212 = v187
		v213 = v191
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v201 = int32(1)
	v205 = base.B2i32(v189 == int32(-1))
	if v189 == int32(-1) {
		goto L73
	} else {
		goto L74
	}
L66:
	;
	v197 = base.B2i32(v184 == int32(-1)) | base.B2i32(base.Ui32(v187) < base.Ui32(v186))
	if v197 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v198 = v189
	goto L69
L68:
	;
	v198 = v184
	goto L69
L69:
	;
	if v197 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v199 = v186
	goto L72
L71:
	;
	v199 = v187
	goto L72
L72:
	;
	v209 = v198
	v211 = v186
	v212 = v199
	v213 = int32(-1)
	goto L62
L73:
	;
	v206 = v201
	goto L75
L74:
	;
	v206 = v186 + v201
	goto L75
L75:
	;
	if v189 == int32(-1) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v208 = int32(3)
	goto L78
L77:
	;
	v208 = v189
	goto L78
L78:
	;
	v209 = v184
	v211 = v206
	v212 = v187
	v213 = v208
	goto L62
L79:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v16)+212))
	if v239 != 0 {
		goto L97
	} else {
		goto L98
	}
L80:
	;
	v215 = int32(-1)
	if v213 == v215 {
		v233 = v209
		v235 = v211
		v236 = v212
		v238 = v215
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v225 = int32(1)
	v229 = base.B2i32(v213 == int32(-1))
	if v213 == int32(-1) {
		goto L90
	} else {
		goto L91
	}
L83:
	;
	v221 = base.B2i32(v209 == int32(-1)) | base.B2i32(base.Ui32(v212) < base.Ui32(v211))
	if v221 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v222 = v213
	goto L86
L85:
	;
	v222 = v209
	goto L86
L86:
	;
	if v221 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v223 = v211
	goto L89
L88:
	;
	v223 = v212
	goto L89
L89:
	;
	v233 = v222
	v235 = v211
	v236 = v223
	v238 = int32(-1)
	goto L79
L90:
	;
	v230 = v225
	goto L92
L91:
	;
	v230 = v211 + v225
	goto L92
L92:
	;
	if v213 == int32(-1) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v232 = int32(4)
	goto L95
L94:
	;
	v232 = v213
	goto L95
L95:
	;
	v233 = v209
	v235 = v230
	v236 = v212
	v238 = v232
	goto L79
L96:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v16)+216))
	if v264 != 0 {
		goto L114
	} else {
		goto L115
	}
L97:
	;
	v240 = int32(-1)
	if v238 == v240 {
		v258 = v233
		v260 = v235
		v261 = v236
		v263 = v240
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v250 = int32(1)
	v254 = base.B2i32(v238 == int32(-1))
	if v238 == int32(-1) {
		goto L107
	} else {
		goto L108
	}
L100:
	;
	v246 = base.B2i32(v233 == int32(-1)) | base.B2i32(base.Ui32(v236) < base.Ui32(v235))
	if v246 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v247 = v238
	goto L103
L102:
	;
	v247 = v233
	goto L103
L103:
	;
	if v246 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v248 = v235
	goto L106
L105:
	;
	v248 = v236
	goto L106
L106:
	;
	v258 = v247
	v260 = v235
	v261 = v248
	v263 = int32(-1)
	goto L96
L107:
	;
	v255 = v250
	goto L109
L108:
	;
	v255 = v235 + v250
	goto L109
L109:
	;
	if v238 == int32(-1) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v257 = int32(5)
	goto L112
L111:
	;
	v257 = v238
	goto L112
L112:
	;
	v258 = v233
	v260 = v255
	v261 = v236
	v263 = v257
	goto L96
L113:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v16)+220))
	if v289 != 0 {
		goto L131
	} else {
		goto L132
	}
L114:
	;
	v265 = int32(-1)
	if v263 == v265 {
		v283 = v258
		v285 = v260
		v286 = v261
		v288 = v265
		goto L113
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v275 = int32(1)
	v279 = base.B2i32(v263 == int32(-1))
	if v263 == int32(-1) {
		goto L124
	} else {
		goto L125
	}
L117:
	;
	v271 = base.B2i32(v258 == int32(-1)) | base.B2i32(base.Ui32(v261) < base.Ui32(v260))
	if v271 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v272 = v263
	goto L120
L119:
	;
	v272 = v258
	goto L120
L120:
	;
	if v271 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v273 = v260
	goto L123
L122:
	;
	v273 = v261
	goto L123
L123:
	;
	v283 = v272
	v285 = v260
	v286 = v273
	v288 = int32(-1)
	goto L113
L124:
	;
	v280 = v275
	goto L126
L125:
	;
	v280 = v260 + v275
	goto L126
L126:
	;
	if v263 == int32(-1) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v282 = int32(6)
	goto L129
L128:
	;
	v282 = v263
	goto L129
L129:
	;
	v283 = v258
	v285 = v280
	v286 = v261
	v288 = v282
	goto L113
L130:
	;
	if base.Ui32(v315) < base.Ui32(int32(2)) {
		goto L153
	} else {
		goto L154
	}
L131:
	;
	if v288 == int32(-1) {
		v312 = v283
		v315 = v286
		goto L130
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v298 = int32(1)
	v302 = base.B2i32(v288 == int32(-1))
	if v288 == int32(-1) {
		goto L141
	} else {
		goto L142
	}
L134:
	;
	v295 = base.B2i32(v283 == int32(-1)) | base.B2i32(base.Ui32(v286) < base.Ui32(v285))
	if v295 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v296 = v288
	goto L137
L136:
	;
	v296 = v283
	goto L137
L137:
	;
	if v295 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v297 = v285
	goto L140
L139:
	;
	v297 = v286
	goto L140
L140:
	;
	v312 = v296
	v315 = v297
	goto L130
L141:
	;
	v303 = v298
	goto L143
L142:
	;
	v303 = v285 + v298
	goto L143
L143:
	;
	v307 = base.B2i32(v283 == int32(-1)) | base.B2i32(base.Ui32(v286) < base.Ui32(v303))
	if v307 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v308 = v303
	goto L146
L145:
	;
	v308 = v286
	goto L146
L146:
	;
	if v288 == int32(-1) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v310 = int32(7)
	goto L149
L148:
	;
	v310 = v288
	goto L149
L149:
	;
	if v307 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v311 = v310
	goto L152
L151:
	;
	v311 = v283
	goto L152
L152:
	;
	v312 = v311
	v315 = v308
	goto L130
L153:
	;
	v320 = int32(-1)
	goto L155
L154:
	;
	v320 = v312
	goto L155
L155:
	;
	if v312 != int32(-1) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v323 = v320
	goto L158
L157:
	;
	v323 = v312
	goto L158
L158:
	;
	v324 = v323 + v315
	v327 = int32(0)
	if base.B2i32(base.B2i32(v323 == int32(-1))|base.B2i32(v327 < v323) == v327)&base.B2i32(v327 < v324) == v327 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v361 = int32(1)
	v366 = v353
	goto L165
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+176)) = v148
	v339 = v16 + int32(224)
	v343 = F_pg_sprintf(m, v339, int32(_a_F_pg_inet_net_ntop_2), v16+int32(176))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L9
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v347 = v16 + int32(224)
	if v323 != 0 {
		v353 = v347
		goto L159
	} else {
		goto L164
	}
L163:
	;
	v353 = v343 + v339
	goto L159
L164:
	;
	v348 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+224)) = uint8(v348)
	v353 = v347 | int32(1)
	goto L159
L165:
	;
	if base.B2i32(v323 == int32(-1))|base.B2i32(v361 < v323)|base.B2i32(v324 <= v361) == int32(0) {
		goto L169
	} else {
		goto L170
	}
L166:
	;
	if base.B2i32(v323 == int32(-1))|base.B2i32(v324 != int32(8)) == int32(0) {
		goto L195
	} else {
		goto L196
	}
L167:
	;
	goto L166
L168:
	;
	v488 = v361 + int32(1)
	if v488 != int32(8) {
		v361 = v488
		v366 = v485
		goto L165
	} else {
		goto L194
	}
L169:
	;
	if v361 != v323 {
		v485 = v366
		goto L168
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v385 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v366))) = uint8(v385)
	v388 = v366 + int32(1)
	if v323|base.B2i32(v361 != int32(6)) != 0 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v381 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v366))) = uint8(v381)
	v485 = v366 + int32(1)
	goto L168
L173:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(192)+v361<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+160)) = v477
	v482 = F_pg_sprintf(m, v388, int32(_a_F_pg_inet_net_ntop_2), v16+int32(160))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L9
	} else {
		goto L193
	}
L174:
	;
	if v315 == int32(6) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v403 = v16 + int32(274)
	if base.Ui32(v403-v388) < base.Ui32(int32(6)) {
		goto L184
	} else {
		goto L185
	}
L176:
	;
	if base.B2i32(v315 != int32(7)) == int32(0) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v16)+220))
	if v396 != int32(1) {
		goto L175
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	if v315 != int32(5) {
		goto L173
	} else {
		goto L181
	}
L180:
	;
	goto L179
L181:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v16)+212))
	if v399 != int32(_a_F_pg_inet_net_ntop_3) {
		goto L173
	} else {
		goto L182
	}
L182:
	;
	goto L175
L183:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v463
	v468 = F_pg_sprintf(m, v452, int32(_a_F_pg_inet_net_ntop_0), v16+int32(96))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L9
	} else {
		goto L192
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_inet_net_ntop[0])) = int32(35)
	goto L2
L185:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v407
	v412 = F_pg_sprintf(m, v388, int32(_a_F_pg_inet_net_ntop_0), v16+int32(144))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L9
	} else {
		goto L186
	}
L186:
	;
	v415 = int32(46)
	*(*uint16)(unsafe.Add(mBase, uint32(v366+v412)+1)) = uint16(v415)
	v418 = v412 + int32(2)
	v419 = v366 + v418
	if base.Ui32(v403-v419) < base.Ui32(int32(6)) {
		goto L184
	} else {
		goto L187
	}
L187:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v423
	v428 = F_pg_sprintf(m, v419, int32(_a_F_pg_inet_net_ntop_0), v16+int32(128))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L9
	} else {
		goto L188
	}
L188:
	;
	v430 = v428 + v418
	v432 = int32(46)
	*(*uint16)(unsafe.Add(mBase, uint32(v366+v430))) = uint16(v432)
	v435 = v430 + int32(1)
	v436 = v366 + v435
	if base.Ui32(v403-v436) < base.Ui32(int32(6)) {
		goto L184
	} else {
		goto L189
	}
L189:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = v440
	v445 = F_pg_sprintf(m, v436, int32(_a_F_pg_inet_net_ntop_0), v16+int32(112))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L9
	} else {
		goto L190
	}
L190:
	;
	v448 = v445 + v366 + v435
	v449 = int32(46)
	*(*uint16)(unsafe.Add(mBase, uint32(v448))) = uint16(v449)
	v452 = v448 + int32(1)
	if base.Ui32(int32(5)) < base.Ui32(v403-v452) {
		goto L183
	} else {
		goto L191
	}
L191:
	;
	goto L184
L192:
	;
	v470 = F_strlen(m, v388)
	mBase = m.M
	v492 = v470 + v388
	goto L167
L193:
	;
	v485 = v482 + v388
	goto L168
L194:
	;
	v492 = v485
	goto L167
L195:
	;
	v503 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v492))) = uint8(v503)
	v507 = v492 + int32(1)
	goto L197
L196:
	;
	v507 = v492
	goto L197
L197:
	;
	v508 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v507))) = uint8(v508)
	if base.B2i32(l2 == int32(-1))|base.B2i32(l2 == int32(128)) != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v522 = v507
	goto L200
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = l2
	v519 = F_pg_sprintf(m, v507, int32(_a_F_pg_inet_net_ntop_1), v16+int32(80))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L9
	} else {
		goto L201
	}
L200:
	;
	if base.Ui32(int32(50)) < base.Ui32(v522-(v16+int32(224))) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v522 = v519 + v507
	goto L200
L202:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_inet_net_ntop[0])) = int32(35)
	goto L2
L203:
	;
	goto L204
L204:
	;
	v532 = v16 + int32(224)
	if (v532^l3)&int32(3) != 0 {
		goto L208
	} else {
		goto L209
	}
L205:
	;
	v630 = l3
	goto L1
L206:
	;
	goto L205
L207:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v587))) = uint8(v586)
	if v586&int32(255) == int32(0) {
		goto L206
	} else {
		goto L222
	}
L208:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532))))
	v585 = v532
	v586 = v538
	v587 = l3
	goto L207
L209:
	;
	goto L210
L210:
	;
	if v532&int32(3) != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v542 = v532
	v544 = l3
	goto L214
L212:
	;
	v556 = v532
	v558 = l3
	goto L213
L213:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	v563 = int32(-2139062144)
	if (int32(16843008)-v560|v560)&v563 != v563 {
		v585 = v556
		v586 = v560
		v587 = v558
		goto L207
	} else {
		goto L218
	}
L214:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542))))
	*(*uint8)(unsafe.Add(mBase, uint32(v544))) = uint8(v545)
	if v545 == int32(0) {
		goto L206
	} else {
		goto L216
	}
L215:
	;
	v556 = v552
	v558 = v550
	goto L213
L216:
	;
	v549 = int32(1)
	v550 = v544 + v549
	v552 = v542 + v549
	if v552&int32(3) != 0 {
		v542 = v552
		v544 = v550
		goto L214
	} else {
		goto L217
	}
L217:
	;
	goto L215
L218:
	;
	v568 = v556
	v569 = v560
	v570 = v558
	goto L219
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v570))) = v569
	v572 = int32(4)
	v573 = v570 + v572
	v575 = v568 + v572
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v568)+4))
	v580 = int32(-2139062144)
	if (int32(16843008)-v577|v577)&v580 == v580 {
		v568 = v575
		v569 = v577
		v570 = v573
		goto L219
	} else {
		goto L221
	}
L220:
	;
	v585 = v575
	v586 = v577
	v587 = v573
	goto L207
L221:
	;
	goto L220
L222:
	;
	v594 = v585
	v596 = v587
	goto L223
L223:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v596)+1)) = uint8(v597)
	v599 = int32(1)
	if v597 != 0 {
		v594 = v594 + v599
		v596 = v596 + v599
		goto L223
	} else {
		goto L225
	}
L224:
	;
	goto L206
L225:
	;
	goto L224
}
func F_pg_input_is_valid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_input_is_valid[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v18
			v21 = *(*int64)(unsafe.Add(mBase, _c_F_pg_input_is_valid[1]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v21
			v23 = F_pg_input_is_valid_common(m, l0, v10, v15, v7)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				m.G0 = v7 + int32(16)
				return base.I64_extend_i32_u(v23)
			}
		}
	}
}
func F_pg_interpret_timezone_abbrev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v191 int32
	_ = v191
	v6 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l4)+268))
	if v13 <= v6 {
		v191 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v191
L2:
	;
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v21 = int32(0)
	goto L3
L3:
	;
	v32 = v21 + (l4 + int32(_a_F_pg_interpret_timezone_abbrev_0))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if base.B2i32(v35 == int32(0))|base.B2i32(v35 != v38) != 0 {
		v56 = v35
		v57 = v38
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l4)+260))
	if v64 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	if v56-v57 != 0 {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	v41 = l0
	v42 = v32
	goto L8
L8:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v46 == int32(0) {
		v56 = v46
		v57 = v45
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v56 = v46
	v57 = v45
	goto L6
L10:
	;
	v49 = int32(1)
	if v46 == v45 {
		v41 = v41 + v49
		v42 = v42 + v49
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v59 = F_strlen(m, v32)
	mBase = m.M
	v62 = v59 + v21 + int32(1)
	if v62 < v13 {
		v21 = v62
		goto L3
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L4
L15:
	;
	v191 = v6
	goto L1
L16:
	;
	v109 = l4 + int32(_a_F_pg_interpret_timezone_abbrev_1)
	v111 = l4 + int32(_a_F_pg_interpret_timezone_abbrev_2)
	v112 = v101
	goto L30
L17:
	;
	v101 = int32(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v71 = v64
	v76 = int32(0)
	goto L20
L20:
	;
	v84 = int32(1)
	v85 = (v71 + v76) >> (uint(v84) % 32)
	v91 = *(*int64)(unsafe.Add(mBase, uint32(l4+int32(280)+v85<<(uint(int32(3))%32))))
	v92 = base.B2i32(v18 < v91)
	if v18 < v91 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v101 = v93
	goto L16
L22:
	;
	v93 = v76
	goto L24
L23:
	;
	v93 = v85 + v84
	goto L24
L24:
	;
	if v18 < v91 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v94 = v85
	goto L27
L26:
	;
	v94 = v71
	goto L27
L27:
	;
	if v93 < v94 {
		v71 = v94
		v76 = v93
		goto L20
	} else {
		goto L28
	}
L28:
	;
	goto L21
L29:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v176
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v178
	v191 = int32(1)
	goto L1
L30:
	;
	if int32(0) < v112 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l4)+uint32(_c_F_pg_interpret_timezone_abbrev[0])))
	v138 = v111 + v135<<(uint(int32(4))%32)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	if v139 == v21 {
		v170 = v138
		goto L29
	} else {
		goto L36
	}
L32:
	;
	v127 = v112 - int32(1)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v127))))
	v132 = v111 + v129<<(uint(int32(4))%32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	if v133 != v21 {
		v112 = v127
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L31
L35:
	;
	v170 = v132
	goto L29
L36:
	;
	if v64 <= v101 {
		v191 = v6
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v147 = v101
	goto L38
L38:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v109))))
	v158 = v111 + v155<<(uint(int32(4))%32)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	if v159 == v21 {
		v170 = v158
		goto L29
	} else {
		goto L40
	}
L39:
	;
	v191 = v6
	goto L1
L40:
	;
	v162 = v147 + int32(1)
	if v64 != v162 {
		v147 = v162
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
}
func F_pg_is_other_temp_schema(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
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
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	v3 = int64(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pg_is_other_temp_schema[0]))
	if v6 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v121
L2:
	;
	if v4 == v6 {
		v121 = v3
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v11 = F_get_namespace_name(m, v4)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_is_other_temp_schema[1]))
	if v9 == v4 {
		v121 = v3
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	return int64(0)
L8:
	;
	if v11 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int64(0)
L10:
	;
	goto L11
L11:
	;
	v19 = int32(_a_F_pg_is_other_temp_schema_0)
	goto L14
L12:
	;
	if v57-v58 != 0 {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	goto L15
L15:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v26 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v27 = v11
	v28 = v19
	v29 = int32(8)
	v30 = v26
	goto L20
L17:
	;
	v53 = v19
	v57 = int32(0)
	goto L18
L18:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	goto L12
L19:
	;
	v53 = v48
	v57 = v50
	goto L18
L20:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if base.B2i32(v30 != v32)|base.B2i32(v32 == int32(0)) != 0 {
		v48 = v28
		v50 = v30
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v48 = v42
	v50 = int32(0)
	goto L19
L22:
	;
	v38 = v29 - int32(1)
	if v38 == int32(0) {
		v48 = v28
		v50 = v30
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v41 = int32(1)
	v42 = v28 + v41
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	if v43 != 0 {
		v27 = v27 + v41
		v28 = v42
		v29 = v38
		v30 = v43
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v66 = int32(_a_F_pg_is_other_temp_schema_1)
	goto L30
L26:
	;
	v117 = int64(1)
	goto L27
L27:
	;
	F_pfree(m, v11)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L7
	} else {
		goto L41
	}
L28:
	;
	v117 = base.I64_extend_i32_u(base.B2i32(v104-v105 == int32(0)))
	goto L27
L30:
	;
	goto L31
L31:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v73 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v74 = v11
	v75 = v66
	v76 = int32(14)
	v77 = v73
	goto L36
L33:
	;
	v100 = v66
	v104 = int32(0)
	goto L34
L34:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	goto L28
L35:
	;
	v100 = v95
	v104 = v97
	goto L34
L36:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if base.B2i32(v77 != v79)|base.B2i32(v79 == int32(0)) != 0 {
		v95 = v75
		v97 = v77
		goto L35
	} else {
		goto L38
	}
L37:
	;
	v95 = v89
	v97 = int32(0)
	goto L35
L38:
	;
	v85 = v76 - int32(1)
	if v85 == int32(0) {
		v95 = v75
		v97 = v77
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v88 = int32(1)
	v89 = v75 + v88
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	if v90 != 0 {
		v74 = v74 + v88
		v75 = v89
		v76 = v85
		v77 = v90
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	v121 = v117
	goto L1
}
func F_pg_isolation_test_session_is_blocked(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v165 int32
	_ = v165
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v197 int64
	_ = v197
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v23 = base.I32_wrap_i64(v16)
	v24 = F_BackendPidGetProc(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L51
	}
L4:
	;
	m.G0 = v14 + int32(16)
	return v197
L5:
	;
	if v24 == int32(0) {
		v197 = int64(0)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+648))
	if v28 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v74 = F_array_contains_nulls(m, v18)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L24
	}
L8:
	;
	if v43 == int32(0) {
		goto L7
	} else {
		goto L15
	}
L9:
	;
	v43 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v33 = v28 - int32(16777216)
	if base.Ui32(int32(184549375)) < base.Ui32(v33) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = int32(_a_F_pg_isolation_test_session_is_blocked_0)
	goto L8
L13:
	;
	goto L14
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v33)>>(uint(int32(22))%32))&int32(1020))+uint32(_c_F_pg_isolation_test_session_is_blocked[0])))
	v43 = v41
	goto L8
L15:
	;
	v46 = int32(_a_F_pg_isolation_test_session_is_blocked_1)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_isolation_test_session_is_blocked[1])))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if base.B2i32(v49 == int32(0))|base.B2i32(v49 != v52) != 0 {
		v70 = v49
		v71 = v52
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v70-v71 != 0 {
		goto L7
	} else {
		goto L23
	}
L17:
	;
	goto L16
L18:
	;
	v55 = v46
	v56 = v43
	goto L19
L19:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v60 == int32(0) {
		v70 = v60
		v71 = v59
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v70 = v60
	v71 = v59
	goto L17
L21:
	;
	v63 = int32(1)
	if v60 == v59 {
		v55 = v55 + v63
		v56 = v56 + v63
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v197 = int64(1)
	goto L4
L24:
	;
	if v74 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v80 = F_ArrayGetNItemsSafe(m, v77, v18+int32(16))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v85 = F_DirectFunctionCall1Coll(m, int32(1758), int32(0), base.I64_extend32_s(v16))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v88 = F_pg_detoast_datum(m, base.I32_wrap_i64(v85))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v94 = F_ArrayGetNItemsSafe(m, v91, v88+int32(16))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if int32(0) < v94 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v90 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v181 = F_GetSafeSnapshotBlockingPids(m, v23, v14+int32(12), int32(1))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L50
	}
L33:
	;
	v104 = v90
	goto L35
L34:
	;
	v104 = (v91<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L35
L35:
	;
	if v76 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v112 = v76
	goto L38
L37:
	;
	v112 = (v77<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L38
L38:
	;
	v114 = int32(0)
	v118 = v114
	goto L39
L39:
	;
	if v80 <= v114 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L32
L41:
	;
	v165 = v118 + int32(1)
	if v165 != v94 {
		v118 = v165
		goto L39
	} else {
		goto L49
	}
L42:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v88+v104+v118<<(uint(int32(2))%32))))
	v133 = int32(0)
	goto L43
L43:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v18+v112+v133<<(uint(int32(2))%32))))
	if v147 != v131 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v197 = int64(1)
	goto L4
L45:
	;
	v150 = v133 + int32(1)
	if v80 != v150 {
		v133 = v150
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	goto L41
L49:
	;
	goto L40
L50:
	;
	v197 = base.I64_extend_i32_u(base.B2i32(int32(0) < v181))
	goto L4
L51:
	;
	F_errmsg_internal(m, int32(_a_F_pg_isolation_test_session_is_blocked_2), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_pg_isolation_test_session_is_blocked_3), int32(66), int32(_a_F_pg_isolation_test_session_is_blocked_4))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_iswalnum(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v4 == int32(0) {
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			v21 = int32(0)
			return v21
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_pg_iswalnum[0]))))
			return base.B2i32(v10&int32(3) != int32(0))
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
		v17 = m.T0[v16].(func(*base.Module, int32, int32) int32)(m, l0, l1)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = v17
			return v21
		}
	}
}
func F_pg_johab_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v7 = base.I32_extend8_s(v6)
	switch v6 - int32(142) {
	case 0:
		v17 = int32(2)
	case 1:
		v17 = int32(3)
	default:
		if int32(0) <= v7 {
			v16 = int32(1)
		} else {
			v16 = int32(2)
		}
		v17 = v16
	}
	if l1 < v17 {
		return int32(-1)
	} else {
		if base.B2i32(base.Ui32(v17) < base.Ui32(int32(2)))|base.B2i32(int32(0) <= v7) != 0 {
			v46 = v17
			return v46
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if base.Ui32(int32(93)) < base.Ui32((v26+int32(95))&int32(255)) {
				return int32(-1)
			} else {
				if v17 != int32(3) {
					v46 = v17
				} else {
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
					if base.Ui32(int32(94)) <= base.Ui32((v38+int32(95))&int32(255)) {
						v45 = int32(-1)
					} else {
						v45 = v17
					}
					v46 = v45
				}
				return v46
			}
		}
	}
}
func F_pg_largeobject_aclmask_snapshot(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = F_superuser_arg(m, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		if v14 != 0 {
			v109 = l2
			m.G0 = v12 + int32(80)
			return v109
		} else {
			v20 = F_table_open(m, int32(2995), int32(1))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				v23 = v12 + int32(16)
				F_ScanKeyInit(m, v23, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v31 = int32(1)
					v33 = F_systable_beginscan(m, v20, int32(2996), v31, l3, v31, v23)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						v35 = F_systable_getnext(m, v33)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
						} else {
							if v35 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(67137668))
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
										F_errmsg(m, int32(_a_F_pg_largeobject_aclmask_snapshot_0), v12)
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_pg_largeobject_aclmask_snapshot_1), int32(3564), int32(_a_F_pg_largeobject_aclmask_snapshot_2))
											mBase = m.M
											v134 = m.ExcPending
											if v134 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v39+v40)+4))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
								v47 = F_heap_getattr_2(m, v35, int32(3), v44, v12+int32(15))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
									if v49 == int32(1) {
										v53 = F_acldefault(m, int32(22), v42)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return int64(0)
										} else {
											v58 = int32(0)
											v59 = v53
											v61 = F_aclmask(m, v59, l1, v42, l2, int32(1))
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int64(0)
											} else {
												v63 = int32(0)
												if base.B2i32(v59 == v63)|base.B2i32(v59 == v58) == v63 {
													F_pfree(m, v59)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int64(0)
													} else {
														F_systable_endscan(m, v33)
														mBase = m.M
														v72 = m.ExcPending
														if v72 != 0 {
															return int64(0)
														} else {
															F_relation_close(m, v20, int32(1))
															mBase = m.M
															v75 = m.ExcPending
															if v75 != 0 {
																return int64(0)
															} else {
																v76 = int64(2)
																v78 = int64(0)
																if base.B2i32(l2&v76 == v78)|base.B2i32(v61&v76 != v78) == int32(0) {
																	v90 = F_has_privs_of_role(m, l1, int32(_a_F_pg_largeobject_aclmask_snapshot_3))
																	mBase = m.M
																	v91 = m.ExcPending
																	if v91 != 0 {
																		return int64(0)
																	} else {
																		if v90 != 0 {
																			v92 = v61 | int64(2)
																		} else {
																			v92 = v61
																		}
																		v93 = v92
																		if l2&int64(4) == int64(0) {
																			v109 = v93
																			m.G0 = v12 + int32(80)
																			return v109
																		} else {
																			if v93&int64(4) != int64(0) {
																				v109 = v93
																				m.G0 = v12 + int32(80)
																				return v109
																			} else {
																				v105 = F_has_privs_of_role(m, l1, int32(_a_F_pg_largeobject_aclmask_snapshot_4))
																				mBase = m.M
																				v106 = m.ExcPending
																				if v106 != 0 {
																					return int64(0)
																				} else {
																					if v105 != 0 {
																						v107 = v93 | int64(4)
																					} else {
																						v107 = v93
																					}
																					v109 = v107
																					m.G0 = v12 + int32(80)
																					return v109
																				}
																			}
																		}
																	}
																} else {
																	v93 = v61
																	if l2&int64(4) == int64(0) {
																		v109 = v93
																		m.G0 = v12 + int32(80)
																		return v109
																	} else {
																		if v93&int64(4) != int64(0) {
																			v109 = v93
																			m.G0 = v12 + int32(80)
																			return v109
																		} else {
																			v105 = F_has_privs_of_role(m, l1, int32(_a_F_pg_largeobject_aclmask_snapshot_4))
																			mBase = m.M
																			v106 = m.ExcPending
																			if v106 != 0 {
																				return int64(0)
																			} else {
																				if v105 != 0 {
																					v107 = v93 | int64(4)
																				} else {
																					v107 = v93
																				}
																				v109 = v107
																				m.G0 = v12 + int32(80)
																				return v109
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													F_systable_endscan(m, v33)
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return int64(0)
													} else {
														F_relation_close(m, v20, int32(1))
														mBase = m.M
														v75 = m.ExcPending
														if v75 != 0 {
															return int64(0)
														} else {
															v76 = int64(2)
															v78 = int64(0)
															if base.B2i32(l2&v76 == v78)|base.B2i32(v61&v76 != v78) == int32(0) {
																v90 = F_has_privs_of_role(m, l1, int32(_a_F_pg_largeobject_aclmask_snapshot_3))
																mBase = m.M
																v91 = m.ExcPending
																if v91 != 0 {
																	return int64(0)
																} else {
																	if v90 != 0 {
																		v92 = v61 | int64(2)
																	} else {
																		v92 = v61
																	}
																	v93 = v92
																	if l2&int64(4) == int64(0) {
																		v109 = v93
																		m.G0 = v12 + int32(80)
																		return v109
																	} else {
																		if v93&int64(4) != int64(0) {
																			v109 = v93
																			m.G0 = v12 + int32(80)
																			return v109
																		} else {
																			v105 = F_has_privs_of_role(m, l1, int32(_a_F_pg_largeobject_aclmask_snapshot_4))
																			mBase = m.M
																			v106 = m.ExcPending
																			if v106 != 0 {
																				return int64(0)
																			} else {
																				if v105 != 0 {
																					v107 = v93 | int64(4)
																				} else {
																					v107 = v93
																				}
																				v109 = v107
																				m.G0 = v12 + int32(80)
																				return v109
																			}
																		}
																	}
																}
															} else {
																v93 = v61
																if l2&int64(4) == int64(0) {
																	v109 = v93
																	m.G0 = v12 + int32(80)
																	return v109
																} else {
																	if v93&int64(4) != int64(0) {
																		v109 = v93
																		m.G0 = v12 + int32(80)
																		return v109
																	} else {
																		v105 = F_has_privs_of_role(m, l1, int32(_a_F_pg_largeobject_aclmask_snapshot_4))
																		mBase = m.M
																		v106 = m.ExcPending
																		if v106 != 0 {
																			return int64(0)
																		} else {
																			if v105 != 0 {
																				v107 = v93 | int64(4)
																			} else {
																				v107 = v93
																			}
																			v109 = v107
																			m.G0 = v12 + int32(80)
																			return v109
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
										v55 = base.I32_wrap_i64(v47)
										v56 = F_pg_detoast_datum(m, v55)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int64(0)
										} else {
											v58 = v55
											v59 = v56
											v61 = F_aclmask(m, v59, l1, v42, l2, int32(1))
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int64(0)
											} else {
												v63 = int32(0)
												if base.B2i32(v59 == v63)|base.B2i32(v59 == v58) == v63 {
													F_pfree(m, v59)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int64(0)
													} else {
														F_systable_endscan(m, v33)
														mBase = m.M
														v72 = m.ExcPending
														if v72 != 0 {
															return int64(0)
														} else {
															F_relation_close(m, v20, int32(1))
															mBase = m.M
															v75 = m.ExcPending
															if v75 != 0 {
																return int64(0)
															} else {
																v76 = int64(2)
																v78 = int64(0)
																if base.B2i32(l2&v76 == v78)|base.B2i32(v61&v76 != v78) == int32(0) {
																	v90 = F_has_privs_of_role(m, l1, int32(_a_F_pg_largeobject_aclmask_snapshot_3))
																	mBase = m.M
																	v91 = m.ExcPending
																	if v91 != 0 {
																		return int64(0)
																	} else {
																		if v90 != 0 {
																			v92 = v61 | int64(2)
																		} else {
																			v92 = v61
																		}
																		v93 = v92
																		if l2&int64(4) == int64(0) {
																			v109 = v93
																			m.G0 = v12 + int32(80)
																			return v109
																		} else {
																			if v93&int64(4) != int64(0) {
																				v109 = v93
																				m.G0 = v12 + int32(80)
																				return v109
																			} else {
																				v105 = F_has_privs_of_role(m, l1, int32(_a_F_pg_largeobject_aclmask_snapshot_4))
																				mBase = m.M
																				v106 = m.ExcPending
																				if v106 != 0 {
																					return int64(0)
																				} else {
																					if v105 != 0 {
																						v107 = v93 | int64(4)
																					} else {
																						v107 = v93
																					}
																					v109 = v107
																					m.G0 = v12 + int32(80)
																					return v109
																				}
																			}
																		}
																	}
																} else {
																	v93 = v61
																	if l2&int64(4) == int64(0) {
																		v109 = v93
																		m.G0 = v12 + int32(80)
																		return v109
																	} else {
																		if v93&int64(4) != int64(0) {
																			v109 = v93
																			m.G0 = v12 + int32(80)
																			return v109
																		} else {
																			v105 = F_has_privs_of_role(m, l1, int32(_a_F_pg_largeobject_aclmask_snapshot_4))
																			mBase = m.M
																			v106 = m.ExcPending
																			if v106 != 0 {
																				return int64(0)
																			} else {
																				if v105 != 0 {
																					v107 = v93 | int64(4)
																				} else {
																					v107 = v93
																				}
																				v109 = v107
																				m.G0 = v12 + int32(80)
																				return v109
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													F_systable_endscan(m, v33)
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return int64(0)
													} else {
														F_relation_close(m, v20, int32(1))
														mBase = m.M
														v75 = m.ExcPending
														if v75 != 0 {
															return int64(0)
														} else {
															v76 = int64(2)
															v78 = int64(0)
															if base.B2i32(l2&v76 == v78)|base.B2i32(v61&v76 != v78) == int32(0) {
																v90 = F_has_privs_of_role(m, l1, int32(_a_F_pg_largeobject_aclmask_snapshot_3))
																mBase = m.M
																v91 = m.ExcPending
																if v91 != 0 {
																	return int64(0)
																} else {
																	if v90 != 0 {
																		v92 = v61 | int64(2)
																	} else {
																		v92 = v61
																	}
																	v93 = v92
																	if l2&int64(4) == int64(0) {
																		v109 = v93
																		m.G0 = v12 + int32(80)
																		return v109
																	} else {
																		if v93&int64(4) != int64(0) {
																			v109 = v93
																			m.G0 = v12 + int32(80)
																			return v109
																		} else {
																			v105 = F_has_privs_of_role(m, l1, int32(_a_F_pg_largeobject_aclmask_snapshot_4))
																			mBase = m.M
																			v106 = m.ExcPending
																			if v106 != 0 {
																				return int64(0)
																			} else {
																				if v105 != 0 {
																					v107 = v93 | int64(4)
																				} else {
																					v107 = v93
																				}
																				v109 = v107
																				m.G0 = v12 + int32(80)
																				return v109
																			}
																		}
																	}
																}
															} else {
																v93 = v61
																if l2&int64(4) == int64(0) {
																	v109 = v93
																	m.G0 = v12 + int32(80)
																	return v109
																} else {
																	if v93&int64(4) != int64(0) {
																		v109 = v93
																		m.G0 = v12 + int32(80)
																		return v109
																	} else {
																		v105 = F_has_privs_of_role(m, l1, int32(_a_F_pg_largeobject_aclmask_snapshot_4))
																		mBase = m.M
																		v106 = m.ExcPending
																		if v106 != 0 {
																			return int64(0)
																		} else {
																			if v105 != 0 {
																				v107 = v93 | int64(4)
																			} else {
																				v107 = v93
																			}
																			v109 = v107
																			m.G0 = v12 + int32(80)
																			return v109
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
func F_pg_last_xact_replay_timestamp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = F_GetLatestXTime(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int64(0) {
			v9 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v9)
		} else {
		}
		return v3
	}
}
func F_pg_localtime(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_localsub(m, l1+int32(256), l0)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_pg_lsn_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*uint32)(unsafe.Add(mBase, uint32(v6)+4)) = uint32(v8)
	v11 = int64(base.Ui64(v8) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6))) = uint32(v11)
	v14 = v6 + int32(16)
	v17 = F_pg_snprintf(m, v14, int32(18), int32(_a_F_pg_lsn_out_0), v6)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = F_pstrdup(m, v14)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			m.G0 = v6 + int32(48)
			return base.I64_extend_i32_u(v21)
		}
	}
}
func F_pg_mbcliplen(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	v4 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mbcliplen[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if base.B2i32(v10 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(v10)) != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v76
L2:
	;
	if v22 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v22 = int32(1)
	goto L5
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v10*int32(28))+uint32(_c_F_pg_mbcliplen[1])))
	v22 = v21
	goto L5
L5:
	;
	goto L2
L6:
	;
	if l1 < l2 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if l1 <= int32(0) {
		v76 = v4
		goto L1
	} else {
		goto L17
	}
L9:
	;
	v26 = l1
	goto L11
L10:
	;
	v26 = l2
	goto L11
L11:
	;
	if v26 <= int32(0) {
		v76 = v4
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v32 = v4
	goto L13
L13:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v32))))
	if v37 == int32(0) {
		v76 = v32
		goto L1
	} else {
		goto L15
	}
L14:
	;
	return v26
L15:
	;
	v41 = v32 + int32(1)
	if v41 != v26 {
		v32 = v41
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v10*int32(28))+uint32(_c_F_pg_mbcliplen[2])))
	v51 = l0
	v52 = l1
	v54 = v4
	goto L18
L18:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v58 == int32(0) {
		v76 = v54
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v76 = v65
	goto L1
L20:
	;
	v61 = m.T0[v50].(func(*base.Module, int32) int32)(m, v51)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	v65 = v61 + v54
	if l2 < v65 {
		v76 = v54
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if l2 == v65 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return l2
L25:
	;
	goto L26
L26:
	;
	v70 = v52 - v61
	if int32(0) < v70 {
		v51 = v51 + v61
		v52 = v70
		v54 = v65
		goto L18
	} else {
		goto L27
	}
L27:
	;
	goto L19
}
func F_pg_mblen_range(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mblen_range[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6*int32(28))+uint32(_c_F_pg_mblen_range[1])))
	v12 = m.T0[v11].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if base.Ui32(l1) < base.Ui32(l0+v12) {
			F_report_invalid_encoding_db(m, l0, v12, l1-l0)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return v12
		}
	}
}
func F_pg_mbstrlen_with_len(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mbstrlen_with_len[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7*int32(28))+uint32(_c_F_pg_mbstrlen_with_len[1])))
	if v12 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l1
L2:
	;
	goto L3
L3:
	;
	if l1 <= int32(0) {
		v47 = v3
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_report_invalid_encoding_db(m, v18, v33, v19)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L14
	}
L5:
	;
	return v47
L6:
	;
	v18 = l0
	v19 = l1
	v21 = v3
	goto L7
L7:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v22 == int32(0) {
		v47 = v21
		goto L5
	} else {
		goto L9
	}
L8:
	;
	v47 = v39
	goto L5
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mbstrlen_with_len[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27*int32(28))+uint32(_c_F_pg_mbstrlen_with_len[2])))
	v33 = m.T0[v32].(func(*base.Module, int32) int32)(m, v18)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v19 < v33 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v39 = v21 + int32(1)
	v41 = v19 - v33
	if int32(0) < v41 {
		v18 = v18 + v33
		v19 = v41
		v21 = v39
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L8
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_mkdir_p(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v13 = F_umask(m, int32(0))
	mBase = m.M
	v16 = F_umask(m, v13&int32(-193))
	mBase = m.M
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v23 = l0 + base.B2i32(v17 == int32(47))
	goto L1
L1:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v28 != int32(47) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v74 = F_umask(m, v13)
	mBase = m.M
	m.G0 = v10 + int32(96)
	return v72
L3:
	;
	goto L2
L4:
	;
	v23 = v23 + int32(1)
	goto L1
L5:
	;
	if v42 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L6:
	;
	v44 = F_mkdir(m, l0, v43)
	mBase = m.M
	if int32(0) <= v44 {
		goto L5
	} else {
		goto L13
	}
L7:
	;
	v40 = F_umask(m, v13)
	mBase = m.M
	v42 = int32(0)
	v43 = l1
	goto L6
L8:
	;
	if v28 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v33)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v35 == v33 {
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v31 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v31)
	goto L7
L12:
	;
	v42 = int32(1)
	v43 = int32(511)
	goto L6
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mkdir_p[0]))
	if v48 != int32(20) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_mkdir_p[0])) = v48
	v72 = int32(-1)
	goto L3
L15:
	;
	v53 = F___fstatat(m, int32(-100), l0, v10, int32(0))
	mBase = m.M
	goto L16
L16:
	;
	if v53 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v54&int32(_a_F_pg_mkdir_p_0) == int32(_a_F_pg_mkdir_p_1) {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	v72 = int32(0)
	goto L3
L20:
	;
	goto L21
L21:
	;
	v66 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v66)
	goto L4
}
func F_pg_node_tree_send(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_json_send(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_pg_num_nonnulls(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v12 = F_count_nulls(m, l0, v6+int32(12), v6+int32(8))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v25 = int64(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			v25 = base.I64_extend_i32_s(v21 - v22)
		}
		m.G0 = v6 + int32(16)
		return v25
	}
}
func F_pg_operator_is_visible(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int64
	_ = v25
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_OperatorIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v25 = int64(0)
		} else {
			v25 = base.I64_extend_i32_u(v14)
		}
		m.G0 = v7 + int32(16)
		return v25
	}
}
func F_pg_parse_query(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_parse_query[0])))
	if v3 == int32(1) {
		v6 = int32(_a_F_pg_parse_query_0)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_parse_query[1])) = int64(4)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_parse_query[2])) = int64(3)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_parse_query[3])) = int64(2)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_parse_query[4])) = int64(1)
		v16 = F___syscall_ret(m, int32(0))
		mBase = m.M
		F_gettimeofday(m, int32(_a_F_pg_parse_query_1))
		mBase = m.M
	} else {
	}
	v20 = F_raw_parser(m, l0, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_parse_query[0])))
		if v25 == int32(1) {
			F_ShowUsage(m, int32(_a_F_pg_parse_query_2))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v32 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_parse_query[5])))
				if v32 == int32(1) {
					v37 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_parse_query[6])))
					F_elog_node_display(m, int32(_a_F_pg_parse_query_3), v20, v37)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						return v20
					}
				} else {
					return v20
				}
			}
		} else {
			v32 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_parse_query[5])))
			if v32 == int32(1) {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_parse_query[6])))
				F_elog_node_display(m, int32(_a_F_pg_parse_query_3), v20, v37)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					return v20
				}
			} else {
				return v20
			}
		}
	}
}
func F_pg_plan_queries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
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
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	v5 = int32(0)
	if l0 == v5 {
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v14 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v5
	v24 = v5
	goto L7
L5:
	;
	v104 = v5
	goto L6
L6:
	;
	return v104
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v23<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v31 == int32(6) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v104 = v91
	goto L6
L9:
	;
	v91 = F_lappend(m, v24, v89)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L13
	} else {
		goto L26
	}
L10:
	;
	v35 = F_palloc0(m, int32(120))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_queries[0])))
	if v54 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	return int32(0)
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = int64(25769804110)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+30)) = uint8(v41)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+100)) = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+112)) = v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v30)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+116)) = v47
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v30)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v49
	v89 = v35
	goto L9
L15:
	;
	v57 = int32(_a_F_pg_plan_queries_0)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_queries[1])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_queries[2])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_queries[3])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_queries[4])) = int64(1)
	v67 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L18
L16:
	;
	goto L17
L17:
	;
	v71 = F_planner(m, v30, l1, l2, l3, int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L13
	} else {
		goto L19
	}
L18:
	;
	F_gettimeofday(m, int32(_a_F_pg_plan_queries_1))
	mBase = m.M
	goto L17
L19:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_queries[0])))
	if v74 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_ShowUsage(m, int32(_a_F_pg_plan_queries_2))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L13
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_queries[5])))
	if v81 != int32(1) {
		v89 = v71
		goto L9
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_queries[6])))
	F_elog_node_display(m, int32(_a_F_pg_plan_queries_3), v71, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L13
	} else {
		goto L25
	}
L25:
	;
	v89 = v71
	goto L9
L26:
	;
	v94 = v23 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v94 < v95 {
		v23 = v94
		v24 = v91
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L8
}
func F_pg_re_throw(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pg_re_throw[0]))
	if v4 != 0 {
		F_pgl_longjmp(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_pg_re_throw[1]))
		v10 = v8 * int32(100)
		v13 = int32(22)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_pg_re_throw[2]))) = v13
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_re_throw[3]))
		v19 = int32(2)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v18<<(uint(v19)%32))+uint32(_c_F_pg_re_throw[4])))
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_pg_re_throw[5]))) = uint8(base.B2i32(base.B2i32(v21 != int32(15))&base.B2i32(v13 < v21) == int32(0)))
		v33 = *(*int32)(unsafe.Add(mBase, _c_F_pg_re_throw[6]))
		if v33 == v19 {
			v37 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_re_throw[7])))
			if v37 == int32(1) {
				v45 = int32(1)
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, _c_F_pg_re_throw[8]))
				v45 = base.B2i32(v42 <= int32(22))
			}
			v47 = v45
		} else {
			v47 = int32(0)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_pg_re_throw[9]))) = uint8(v47)
		*(*int32)(unsafe.Add(mBase, _c_F_pg_re_throw[10])) = int32(0)
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_pg_re_throw[11])))
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_pg_re_throw[12])))
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_pg_re_throw[13])))
		F_errfinish(m, v54, v57, v60)
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return
		} else {
			v63 = m.G0
			v65 = v63 - int32(32)
			m.G0 = v65
			v67 = m.Env.Pgmem_getpid(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v67
			*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = int32(2239)
			*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = int32(_a_F_pg_re_throw_0)
			*(*int32)(unsafe.Add(mBase, uint32(v65))) = int32(_a_F_pg_re_throw_1)
			F_write_stderr(m, int32(_a_F_pg_re_throw_2), v65)
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return
			} else {
				v79 = *(*int32)(unsafe.Add(mBase, _c_F_pg_re_throw[14]))
				v80 = F_fflush(m, v79)
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_pg_read_file_off_len_missing(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
		v13 = F_pg_read_file_common(m, v4, v8, v9, base.B2i32(v10 != int64(0)))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			if v13 == int32(0) {
				v17 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v13)
			}
		}
	}
}
func F_pg_reg_getnumoutarcs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 < v3 {
		v28 = v3
		m.G0 = v7 + int32(16)
		return v28
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v13 = v11 + int32(20)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		if v14 <= l1 {
			v28 = v3
			m.G0 = v7 + int32(16)
			return v28
		} else {
			v16 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v16
			F_traverse_lacons(m, v13, l1, v7+int32(12), v16, v16)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v28 = v26
				m.G0 = v7 + int32(16)
				return v28
			}
		}
	}
}
func F_pg_reg_getoutarcs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	v5 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if base.B2i32(l1 < v5)|base.B2i32(l3 <= v5) != 0 {
		m.G0 = v8 + int32(16)
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v17 = v15 + int32(20)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		if v18 <= l1 {
			m.G0 = v8 + int32(16)
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
			F_traverse_lacons(m, v17, l1, v8+int32(12), l2, l3)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_pg_relpages_impl(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+119)))
	switch v9 - int32(83) {
	case 0, 22, 26, 31, 33:
		v38 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int64(0)
		} else {
			F_relation_close(m, l0, int32(1))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int64(0)
			} else {
				m.G0 = v6 + int32(16)
				return base.I64_extend_i32_u(v38)
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v21 + int32(4)
				F_errmsg(m, int32(_a_F_pg_relpages_impl_0), v6)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v29 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28)+119)))
					F_errdetail_relkind_not_supported(m, v29)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_relpages_impl_1), int32(483), int32(_a_F_pg_relpages_impl_2))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
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
func F_pg_restore_extended_stats(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v35 int64
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
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
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v253 int64
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
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
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v383 int32
	_ = v383
	var v388 int64
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v656 int32
	_ = v656
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v765 int32
	_ = v765
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v862 int32
	_ = v862
	var v863 int64
	_ = v863
	var v867 int32
	_ = v867
	var v871 int64
	_ = v871
	var v878 int64
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v910 int32
	_ = v910
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v950 int32
	_ = v950
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v999 int32
	_ = v999
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1097 int32
	_ = v1097
	var v1102 int32
	_ = v1102
	var v1142 int32
	_ = v1142
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1192 int32
	_ = v1192
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1366 int64
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1399 int32
	_ = v1399
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1441 int32
	_ = v1441
	var v1476 int32
	_ = v1476
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1485 int32
	_ = v1485
	var v1524 int32
	_ = v1524
	var v1529 int32
	_ = v1529
	var v1572 int32
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1581 int32
	_ = v1581
	var v1587 int32
	_ = v1587
	var v1592 int32
	_ = v1592
	var v1632 int32
	_ = v1632
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1678 int32
	_ = v1678
	var v1717 int32
	_ = v1717
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1727 int32
	_ = v1727
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1814 int32
	_ = v1814
	var v1831 int32
	_ = v1831
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1877 int32
	_ = v1877
	var v1881 int32
	_ = v1881
	var v1887 int32
	_ = v1887
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1923 int32
	_ = v1923
	var v1928 int32
	_ = v1928
	var v1934 int32
	_ = v1934
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1975 int32
	_ = v1975
	var v1977 int32
	_ = v1977
	var v1979 int32
	_ = v1979
	var v1981 int32
	_ = v1981
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1997 int32
	_ = v1997
	var v2037 int32
	_ = v2037
	var v2039 int32
	_ = v2039
	var v2041 float64
	_ = v2041
	var v2044 float64
	_ = v2044
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2055 int32
	_ = v2055
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2110 int32
	_ = v2110
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2160 int32
	_ = v2160
	var v2162 int32
	_ = v2162
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2173 int32
	_ = v2173
	var v2210 int32
	_ = v2210
	var v2212 int32
	_ = v2212
	var v2215 int32
	_ = v2215
	var v2219 int32
	_ = v2219
	var v2221 int32
	_ = v2221
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2230 int32
	_ = v2230
	var v2233 int64
	_ = v2233
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2241 int32
	_ = v2241
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2253 int32
	_ = v2253
	var v2259 int32
	_ = v2259
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2268 int32
	_ = v2268
	var v2272 int32
	_ = v2272
	var v2313 int32
	_ = v2313
	var v2317 int32
	_ = v2317
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2359 int32
	_ = v2359
	var v2361 int32
	_ = v2361
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2368 int32
	_ = v2368
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2377 int32
	_ = v2377
	var v2379 int32
	_ = v2379
	var v2381 int32
	_ = v2381
	var v2384 int32
	_ = v2384
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2430 int32
	_ = v2430
	var v2432 int32
	_ = v2432
	var v2434 int32
	_ = v2434
	var v2439 int32
	_ = v2439
	var v2443 int32
	_ = v2443
	var v2448 int32
	_ = v2448
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2488 int32
	_ = v2488
	var v2490 int32
	_ = v2490
	var v2492 int32
	_ = v2492
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2501 int32
	_ = v2501
	var v2505 int32
	_ = v2505
	var v2510 int32
	_ = v2510
	var v2549 int32
	_ = v2549
	var v2553 int32
	_ = v2553
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2599 int32
	_ = v2599
	var v2601 int32
	_ = v2601
	var v2602 int32
	_ = v2602
	var v2643 int32
	_ = v2643
	var v2684 int64
	_ = v2684
	var v2691 int32
	_ = v2691
	var v2694 int32
	_ = v2694
	var v2699 int32
	_ = v2699
	var v2703 int32
	_ = v2703
	var v2708 int32
	_ = v2708
	var v2712 int32
	_ = v2712
	var v2718 int32
	_ = v2718
	var v2723 int32
	_ = v2723
	var v2727 int32
	_ = v2727
	var v2733 int32
	_ = v2733
	var v2738 int32
	_ = v2738
	var v2740 int32
	_ = v2740
	var v2745 int32
	_ = v2745
	var v2766 int32
	_ = v2766
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2796 int32
	_ = v2796
	var v2798 int32
	_ = v2798
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2802 int32
	_ = v2802
	var v2809 int32
	_ = v2809
	var v2810 int32
	_ = v2810
	var v2815 int32
	_ = v2815
	var v2821 int32
	_ = v2821
	var v2826 int32
	_ = v2826
	var v2832 int32
	_ = v2832
	var v2833 int32
	_ = v2833
	var v2838 int32
	_ = v2838
	var v2845 int32
	_ = v2845
	var v2850 int32
	_ = v2850
	var v2855 int32
	_ = v2855
	var v2860 int32
	_ = v2860
	var v2863 int32
	_ = v2863
	var v2877 int32
	_ = v2877
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2892 int32
	_ = v2892
	var v2920 int32
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2922 int64
	_ = v2922
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2932 int32
	_ = v2932
	var v2934 int32
	_ = v2934
	var v2936 int32
	_ = v2936
	var v2938 int32
	_ = v2938
	var v2939 int32
	_ = v2939
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2947 int64
	_ = v2947
	var v2956 int32
	_ = v2956
	var v2963 int32
	_ = v2963
	var v2964 int32
	_ = v2964
	var v2969 int32
	_ = v2969
	var v2976 int32
	_ = v2976
	var v2981 int32
	_ = v2981
	var v2985 int32
	_ = v2985
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3030 int32
	_ = v3030
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3035 int32
	_ = v3035
	var v3038 int32
	_ = v3038
	var v3039 int32
	_ = v3039
	var v3044 int32
	_ = v3044
	var v3051 int32
	_ = v3051
	var v3057 int32
	_ = v3057
	var v3062 int32
	_ = v3062
	var v3066 int32
	_ = v3066
	var v3069 int32
	_ = v3069
	var v3072 int32
	_ = v3072
	var v3073 int32
	_ = v3073
	var v3076 int32
	_ = v3076
	var v3078 int32
	_ = v3078
	var v3080 int32
	_ = v3080
	var v3081 int32
	_ = v3081
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3125 int32
	_ = v3125
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3133 int32
	_ = v3133
	var v3134 int32
	_ = v3134
	var v3138 int32
	_ = v3138
	var v3175 int32
	_ = v3175
	var v3176 int32
	_ = v3176
	var v3183 int32
	_ = v3183
	var v3184 int32
	_ = v3184
	var v3185 int32
	_ = v3185
	var v3186 int32
	_ = v3186
	var v3187 int32
	_ = v3187
	var v3189 int32
	_ = v3189
	var v3195 int32
	_ = v3195
	var v3198 int32
	_ = v3198
	var v3199 int32
	_ = v3199
	var v3200 int32
	_ = v3200
	var v3205 int32
	_ = v3205
	var v3207 int32
	_ = v3207
	var v3210 int32
	_ = v3210
	var v3214 int32
	_ = v3214
	var v3215 int32
	_ = v3215
	var v3222 int32
	_ = v3222
	var v3226 int32
	_ = v3226
	var v3231 int32
	_ = v3231
	var v3232 int32
	_ = v3232
	var v3233 int32
	_ = v3233
	var v3234 int32
	_ = v3234
	var v3238 int32
	_ = v3238
	var v3239 int32
	_ = v3239
	var v3242 int32
	_ = v3242
	var v3248 int32
	_ = v3248
	var v3253 int32
	_ = v3253
	var v3255 int32
	_ = v3255
	var v3295 int32
	_ = v3295
	var v3299 int32
	_ = v3299
	var v3300 int32
	_ = v3300
	var v3306 int32
	_ = v3306
	var v3307 int32
	_ = v3307
	var v3313 int32
	_ = v3313
	var v3314 int32
	_ = v3314
	var v3355 int32
	_ = v3355
	var v3356 int32
	_ = v3356
	var v3358 int32
	_ = v3358
	var v3359 int32
	_ = v3359
	var v3362 int32
	_ = v3362
	var v3363 int32
	_ = v3363
	var v3368 int32
	_ = v3368
	var v3375 int32
	_ = v3375
	var v3384 int32
	_ = v3384
	var v3389 int32
	_ = v3389
	var v3390 int32
	_ = v3390
	var v3391 int32
	_ = v3391
	var v3393 int32
	_ = v3393
	var v3394 int32
	_ = v3394
	var v3397 int32
	_ = v3397
	var v3398 int32
	_ = v3398
	var v3403 int32
	_ = v3403
	var v3410 int32
	_ = v3410
	var v3419 int32
	_ = v3419
	var v3424 int32
	_ = v3424
	var v3425 int32
	_ = v3425
	var v3426 int32
	_ = v3426
	var v3428 int32
	_ = v3428
	var v3430 int32
	_ = v3430
	var v3431 int32
	_ = v3431
	var v3434 int32
	_ = v3434
	var v3435 int32
	_ = v3435
	var v3440 int32
	_ = v3440
	var v3447 int32
	_ = v3447
	var v3458 int32
	_ = v3458
	var v3463 int32
	_ = v3463
	var v3465 int32
	_ = v3465
	var v3466 int32
	_ = v3466
	var v3467 int32
	_ = v3467
	var v3471 int32
	_ = v3471
	var v3473 int32
	_ = v3473
	var v3475 int32
	_ = v3475
	var v3476 int32
	_ = v3476
	var v3478 int64
	_ = v3478
	var v3488 int64
	_ = v3488
	var v3500 int64
	_ = v3500
	var v3572 int32
	_ = v3572
	var v3583 int32
	_ = v3583
	var v3584 int32
	_ = v3584
	var v3585 int32
	_ = v3585
	var v3586 int32
	_ = v3586
	var v3589 int32
	_ = v3589
	var v3590 int32
	_ = v3590
	var v3595 int32
	_ = v3595
	var v3602 int32
	_ = v3602
	var v3607 int32
	_ = v3607
	var v3610 int32
	_ = v3610
	var v3613 int32
	_ = v3613
	var v3614 int32
	_ = v3614
	var v3617 int32
	_ = v3617
	var v3618 int32
	_ = v3618
	var v3623 int32
	_ = v3623
	var v3630 int32
	_ = v3630
	var v3641 int32
	_ = v3641
	var v3646 int32
	_ = v3646
	var v3647 int32
	_ = v3647
	var v3656 int32
	_ = v3656
	var v3657 int32
	_ = v3657
	var v3660 int64
	_ = v3660
	var v3662 int32
	_ = v3662
	var v3669 int32
	_ = v3669
	var v3670 int32
	_ = v3670
	var v3673 int64
	_ = v3673
	var v3675 int32
	_ = v3675
	var v3682 int32
	_ = v3682
	var v3683 int32
	_ = v3683
	var v3686 int64
	_ = v3686
	var v3691 int32
	_ = v3691
	var v3692 int32
	_ = v3692
	var v3695 int32
	_ = v3695
	var v3696 int32
	_ = v3696
	var v3697 int32
	_ = v3697
	var v3698 int32
	_ = v3698
	var v3705 int64
	_ = v3705
	var v3706 int32
	_ = v3706
	var v3708 int32
	_ = v3708
	var v3709 int32
	_ = v3709
	var v3712 int32
	_ = v3712
	var v3713 int32
	_ = v3713
	var v3714 int32
	_ = v3714
	var v3715 int32
	_ = v3715
	var v3724 int64
	_ = v3724
	var v3725 int32
	_ = v3725
	var v3727 int32
	_ = v3727
	var v3728 int32
	_ = v3728
	var v3731 int32
	_ = v3731
	var v3737 int32
	_ = v3737
	var v3738 int32
	_ = v3738
	var v3740 int32
	_ = v3740
	var v3741 int32
	_ = v3741
	var v3742 int32
	_ = v3742
	var v3743 int32
	_ = v3743
	var v3747 int32
	_ = v3747
	var v3748 int32
	_ = v3748
	var v3753 int32
	_ = v3753
	var v3762 int32
	_ = v3762
	var v3767 int32
	_ = v3767
	var v3775 int32
	_ = v3775
	var v3776 int32
	_ = v3776
	var v3779 int32
	_ = v3779
	var v3784 int32
	_ = v3784
	var v3787 int32
	_ = v3787
	var v3790 int32
	_ = v3790
	var v3791 int32
	_ = v3791
	var v3792 int32
	_ = v3792
	var v3793 int32
	_ = v3793
	var v3800 int64
	_ = v3800
	var v3801 int32
	_ = v3801
	var v3803 int32
	_ = v3803
	var v3804 int32
	_ = v3804
	var v3814 int32
	_ = v3814
	var v3819 int32
	_ = v3819
	var v3823 int32
	_ = v3823
	var v3831 int32
	_ = v3831
	var v3832 int32
	_ = v3832
	var v3833 int32
	_ = v3833
	var v3838 int32
	_ = v3838
	var v3839 int32
	_ = v3839
	var v3847 int32
	_ = v3847
	var v3853 int32
	_ = v3853
	var v3855 int32
	_ = v3855
	var v3858 int32
	_ = v3858
	var v3859 int32
	_ = v3859
	var v3860 int32
	_ = v3860
	var v3861 int32
	_ = v3861
	var v3865 int32
	_ = v3865
	var v3869 int64
	_ = v3869
	var v3870 int32
	_ = v3870
	var v3872 int32
	_ = v3872
	var v3873 int32
	_ = v3873
	var v3876 int32
	_ = v3876
	var v3877 int32
	_ = v3877
	var v3878 int32
	_ = v3878
	var v3879 int32
	_ = v3879
	var v3888 int64
	_ = v3888
	var v3889 int32
	_ = v3889
	var v3891 int32
	_ = v3891
	var v3892 int32
	_ = v3892
	var v3895 int32
	_ = v3895
	var v3907 int32
	_ = v3907
	var v3908 int32
	_ = v3908
	var v3911 int32
	_ = v3911
	var v3918 int32
	_ = v3918
	var v3921 int32
	_ = v3921
	var v3922 int32
	_ = v3922
	var v3923 int32
	_ = v3923
	var v3924 int32
	_ = v3924
	var v3933 int64
	_ = v3933
	var v3934 int32
	_ = v3934
	var v3936 int32
	_ = v3936
	var v3937 int32
	_ = v3937
	var v3947 int32
	_ = v3947
	var v3952 int32
	_ = v3952
	var v3958 int32
	_ = v3958
	var v3959 int32
	_ = v3959
	var v3960 int32
	_ = v3960
	var v3961 int32
	_ = v3961
	var v3962 int32
	_ = v3962
	var v3963 int32
	_ = v3963
	var v3966 int32
	_ = v3966
	var v3967 int32
	_ = v3967
	var v3968 int32
	_ = v3968
	var v3969 int32
	_ = v3969
	var v3975 int32
	_ = v3975
	var v3976 int64
	_ = v3976
	var v3977 int32
	_ = v3977
	var v3978 int32
	_ = v3978
	var v3988 int32
	_ = v3988
	var v3994 int32
	_ = v3994
	var v3999 int32
	_ = v3999
	var v4000 int32
	_ = v4000
	var v4005 int32
	_ = v4005
	var v4006 int32
	_ = v4006
	var v4007 int32
	_ = v4007
	var v4010 int32
	_ = v4010
	var v4011 int32
	_ = v4011
	var v4012 int32
	_ = v4012
	var v4013 int32
	_ = v4013
	var v4022 int64
	_ = v4022
	var v4023 int32
	_ = v4023
	var v4024 int32
	_ = v4024
	var v4035 int32
	_ = v4035
	var v4040 int32
	_ = v4040
	var v4046 int32
	_ = v4046
	var v4051 int32
	_ = v4051
	var v4052 int32
	_ = v4052
	var v4053 int32
	_ = v4053
	var v4054 int64
	_ = v4054
	var v4055 int32
	_ = v4055
	var v4057 int32
	_ = v4057
	var v4062 int32
	_ = v4062
	var v4063 int32
	_ = v4063
	var v4066 int32
	_ = v4066
	var v4073 int32
	_ = v4073
	var v4078 int32
	_ = v4078
	var v4082 int32
	_ = v4082
	var v4126 int32
	_ = v4126
	var v4127 int32
	_ = v4127
	var v4158 int64
	_ = v4158
	var v4162 int32
	_ = v4162
	var v4164 int32
	_ = v4164
	var v4165 int32
	_ = v4165
	var v4166 int32
	_ = v4166
	var v4168 int32
	_ = v4168
	var v4171 int32
	_ = v4171
	var v4172 int64
	_ = v4172
	var v4173 int32
	_ = v4173
	var v4175 int64
	_ = v4175
	var v4178 int32
	_ = v4178
	var v4215 int64
	_ = v4215
	var v4219 int32
	_ = v4219
	var v4222 int32
	_ = v4222
	var v4302 int32
	_ = v4302
	var v4305 int32
	_ = v4305
	var v4319 int32
	_ = v4319
	var v4323 int32
	_ = v4323
	var v4324 int32
	_ = v4324
	var v4330 int32
	_ = v4330
	var v4332 int32
	_ = v4332
	var v4333 int32
	_ = v4333
	var v4342 int32
	_ = v4342
	var v4345 int32
	_ = v4345
	var v4346 int32
	_ = v4346
	var v4348 int32
	_ = v4348
	var v4349 int32
	_ = v4349
	var v4350 int32
	_ = v4350
	var v4357 int32
	_ = v4357
	var v4358 int32
	_ = v4358
	var v4362 int32
	_ = v4362
	var v4364 int32
	_ = v4364
	var v4369 int32
	_ = v4369
	var v4370 int32
	_ = v4370
	var v4372 int32
	_ = v4372
	var v4373 int32
	_ = v4373
	var v4375 int32
	_ = v4375
	var v4377 int32
	_ = v4377
	var v4380 int32
	_ = v4380
	var v4382 int32
	_ = v4382
	var v4384 int32
	_ = v4384
	var v4396 int32
	_ = v4396
	var v4400 int32
	_ = v4400
	var v4401 int32
	_ = v4401
	var v4407 int32
	_ = v4407
	var v4409 int32
	_ = v4409
	var v4410 int32
	_ = v4410
	var v4420 int32
	_ = v4420
	var v4422 int32
	_ = v4422
	var v4424 int32
	_ = v4424
	var v4436 int32
	_ = v4436
	var v4440 int32
	_ = v4440
	var v4441 int32
	_ = v4441
	var v4449 int32
	_ = v4449
	var v4450 int32
	_ = v4450
	var v4461 int32
	_ = v4461
	var v4463 int32
	_ = v4463
	var v4465 int32
	_ = v4465
	var v4469 int32
	_ = v4469
	var v4471 int32
	_ = v4471
	var v4473 int32
	_ = v4473
	var v4499 int32
	_ = v4499
	v2 = int32(0)
	v35 = int64(0)
	v39 = m.G0
	v41 = v39 - int32(1648)
	m.G0 = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+576)) = uint8(v2)
	*(*int64)(unsafe.Add(mBase, uint32(v41)+568)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v41)+560)) = v35
	v49 = int32(11)
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+578)) = uint16(v49)
	v54 = F_stats_fill_fcinfo_from_arg_pairs(m, l0, v41+int32(560), int32(_a_F_pg_restore_extended_stats_0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v58 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v41)+824)) = v58
	*(*int64)(unsafe.Add(mBase, uint32(v41)+816)) = v58
	*(*int64)(unsafe.Add(mBase, uint32(v41)+808)) = v58
	*(*int64)(unsafe.Add(mBase, uint32(v41)+800)) = v58
	*(*int64)(unsafe.Add(mBase, uint32(v41)+792)) = v58
	*(*int64)(unsafe.Add(mBase, uint32(v41)+784)) = v58
	v70 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+780)) = uint16(v70)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+776)) = v70
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+772)) = uint16(v70)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+768)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v41)+760)) = v70
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+752)))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+688)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+672)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+720)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+704)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+736)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[0])))
	if v88 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v4471 + int32(1648)
	return base.I64_extend_i32_u(v4473&v4499) & int64(1)
L4:
	;
	if v98 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[1]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+308))
	v96 = base.B2i32(v94 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[0])) = uint8(v96)
	v98 = v96
	goto L7
L6:
	;
	v98 = v70
	goto L7
L7:
	;
	goto L4
L8:
	;
	v101 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v122 = v41 + int32(560)
	F_stats_check_required_arg(m, v122, int32(_a_F_pg_restore_extended_stats_0), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	if v101 == int32(0) {
		v4471 = v41
		v4473 = v2
		v4499 = v54
		goto L3
	} else {
		goto L12
	}
L12:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_1), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errhint(m, int32(_a_F_pg_restore_extended_stats_2), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(371), int32(_a_F_pg_restore_extended_stats_4))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v4471 = v41
	v4473 = v2
	v4499 = v54
	goto L3
L17:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v41)+584))
	v128 = F_text_to_cstring(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_stats_check_required_arg(m, v122, int32(_a_F_pg_restore_extended_stats_0), int32(1))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v41)+600))
	v135 = F_text_to_cstring(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_stats_check_required_arg(m, v122, int32(_a_F_pg_restore_extended_stats_0), int32(2))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v41)+616))
	v142 = F_text_to_cstring(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_stats_check_required_arg(m, v122, int32(_a_F_pg_restore_extended_stats_0), int32(3))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v41)+632))
	v149 = F_text_to_cstring(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_stats_check_required_arg(m, v122, int32(_a_F_pg_restore_extended_stats_0), int32(4))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v41)+648))
	v157 = F_makeRangeVar(m, v128, v135, int32(-1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v164 = F_RangeVarGetRelidExtended(m, v157, int32(4), int32(0), int32(1138), v41+int32(760))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v167 = F_get_namespace_oid(m, v142, int32(1))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v167 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v171 = int32(0)
	v174 = F_errstart(m, int32(19), v171)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v192 = F_table_open(m, int32(3381), int32(3))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L38
	}
L32:
	;
	if v174 == int32(0) {
		v4471 = v41
		v4473 = v171
		v4499 = v54
		goto L3
	} else {
		goto L33
	}
L33:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v142
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_5), v41)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(404), int32(_a_F_pg_restore_extended_stats_4))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v4471 = v41
	v4473 = v171
	v4499 = v54
	goto L3
L37:
	;
	if v4449 != 0 {
		goto L708
	} else {
		goto L709
	}
L38:
	;
	v194 = F_get_pg_statistic_ext(m, v192, v167, v149)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v194 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v198 = int32(0)
	v202 = F_errstart(m, int32(19), v198)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v194)+16))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+22)))
	v223 = v221 + v222
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v164 != v224 {
		goto L49
	} else {
		goto L50
	}
L43:
	;
	if v202 == int32(0) {
		v4422 = v41
		v4424 = v198
		v4436 = v198
		v4440 = v2
		v4441 = v2
		v4449 = v192
		v4450 = v54
		goto L37
	} else {
		goto L44
	}
L44:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v142
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_6), v41+int32(16))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(417), int32(_a_F_pg_restore_extended_stats_4))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v4422 = v41
	v4424 = v198
	v4436 = v198
	v4440 = v2
	v4441 = v2
	v4449 = v192
	v4450 = v54
	goto L37
L48:
	;
	F_pfree(m, v4407)
	mBase = m.M
	v4420 = m.ExcPending
	if v4420 != 0 {
		goto L1
	} else {
		goto L707
	}
L49:
	;
	v226 = int32(0)
	v230 = F_errstart(m, int32(19), v226)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v253 = F_SysCacheGetAttrNotNull(m, int32(64), v194, int32(8))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L63
	}
L52:
	;
	if v230 == int32(0) {
		v4382 = v41
		v4384 = v226
		v4396 = v226
		v4400 = v2
		v4401 = v2
		v4407 = v194
		v4409 = v192
		v4410 = v54
		goto L48
	} else {
		goto L53
	}
L53:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+556)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v41)+552)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v41)+548)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v41)+544)) = v142
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_7), v41+int32(544))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(434), int32(_a_F_pg_restore_extended_stats_4))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v4382 = v41
	v4384 = v226
	v4396 = v226
	v4400 = v2
	v4401 = v2
	v4407 = v194
	v4409 = v192
	v4410 = v54
	goto L48
L57:
	;
	if v628&int32(1) != 0 {
		goto L397
	} else {
		goto L398
	}
L58:
	;
	v2766 = int32(0)
	goto L57
L59:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), v2740, int32(_a_F_pg_restore_extended_stats_8))
	mBase = m.M
	v2745 = m.ExcPending
	if v2745 != 0 {
		goto L1
	} else {
		goto L394
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2727 = m.ExcPending
	if v2727 != 0 {
		goto L1
	} else {
		goto L391
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2712 = m.ExcPending
	if v2712 != 0 {
		goto L1
	} else {
		goto L388
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2699 = m.ExcPending
	if v2699 != 0 {
		goto L1
	} else {
		goto L385
	}
L63:
	;
	v256 = F_pg_detoast_datum(m, base.I32_wrap_i64(v253))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v258 != int32(1) {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	if v261 != 0 {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	if v262 != int32(18) {
		goto L62
	} else {
		goto L67
	}
L67:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	if v265 <= int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v223)+96))
	v388 = F_SysCacheGetAttr(m, int32(64), v194, int32(9), v41+int32(767))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L84
	}
L69:
	;
	v345 = int32(0)
	v350 = v2
	v363 = v2
	v366 = v2
	goto L68
L70:
	;
	goto L71
L71:
	;
	v271 = int32(0)
	v275 = v271
	v276 = v271
	v278 = v2
	v291 = v2
	v294 = v2
	goto L72
L72:
	;
	v312 = v276 + (v256 + int32(24))
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312))))
	switch v313 - int32(100) {
	case 0:
		v338 = int32(1)
		v339 = v278
		v340 = v291
		v341 = v294
		goto L74
	case 1:
		goto L78
	case 2:
		goto L76
	default:
		goto L77
	case 9:
		goto L79
	}
L73:
	;
	v345 = v338
	v350 = v339
	v363 = v340
	v366 = v341
	goto L68
L74:
	;
	v343 = v276 + int32(1)
	if v343 != v265 {
		v275 = v338
		v276 = v343
		v278 = v339
		v291 = v340
		v294 = v341
		goto L72
	} else {
		goto L83
	}
L75:
	;
	v338 = v275
	v339 = v335
	v340 = v336
	v341 = v337
	goto L74
L76:
	;
	v335 = v278
	v336 = v291
	v337 = int32(1)
	goto L75
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	v335 = v278
	v336 = int32(1)
	v337 = v294
	goto L75
L79:
	;
	v335 = int32(1)
	v336 = v291
	v337 = v294
	goto L75
L80:
	;
	v322 = int32(*(*int8)(unsafe.Add(mBase, uint32(v312))))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v322
	F_errmsg_internal(m, int32(_a_F_pg_restore_extended_stats_9), v41+int32(32))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(258), int32(_a_F_pg_restore_extended_stats_10))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
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
	goto L73
L84:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+767)))
	if v390 != 0 {
		v409 = v2
		v410 = v2
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v411 = v345 | v82
	if v411&int32(1) != 0 {
		goto L96
	} else {
		goto L97
	}
L86:
	;
	v392 = F_text_to_cstring(m, base.I32_wrap_i64(v388))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v394 = F_stringToNode(m, v392)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_pfree(m, v392)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v399 = F_eval_const_expressions(m, int32(0), v394)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_fix_opfuncids(m, v399)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	if v399 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v409 = int32(0)
	v410 = v2
	goto L85
L93:
	;
	goto L94
L94:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	v409 = v399
	v410 = v406
	goto L85
L95:
	;
	if (v81|v366)&int32(1) != 0 {
		goto L106
	} else {
		goto L107
	}
L96:
	;
	v445 = v82 ^ int32(1)
	goto L95
L97:
	;
	goto L98
L98:
	;
	v418 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	if v418 == int32(0) {
		v445 = v2
		goto L95
	} else {
		goto L100
	}
L100:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+528)) = v426
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_11), v41+int32(528))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+516)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v41)+512)) = v142
	F_errhint(m, int32(_a_F_pg_restore_extended_stats_12), v41+int32(512))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(487), int32(_a_F_pg_restore_extended_stats_4))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v445 = v2
	goto L95
L105:
	;
	v485 = v84 | v83 | v85
	v486 = int32(1)
	if v350&v486 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L106:
	;
	v483 = v411
	v484 = v81 ^ int32(1)
	goto L105
L107:
	;
	goto L108
L108:
	;
	v452 = int32(0)
	v456 = F_errstart(m, int32(19), v452)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	if v456 == int32(0) {
		v483 = v452
		v484 = v452
		goto L105
	} else {
		goto L110
	}
L110:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+496)) = v464
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_11), v41+int32(496))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+484)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v41)+480)) = v142
	F_errhint(m, int32(_a_F_pg_restore_extended_stats_12), v41+int32(480))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(504), int32(_a_F_pg_restore_extended_stats_4))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v483 = v452
	v484 = v452
	goto L105
L115:
	;
	if (v80|v363)&int32(1) != 0 {
		goto L140
	} else {
		goto L141
	}
L116:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), v585, int32(_a_F_pg_restore_extended_stats_4))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L138
	}
L117:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+704)))
	if v492 != int32(1) {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	goto L119
L119:
	;
	if v485&int32(1) == int32(0) {
		v589 = v483
		v590 = v486
		goto L115
	} else {
		goto L129
	}
L120:
	;
	v507 = int32(0)
	v511 = F_errstart(m, int32(19), v507)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L124
	}
L121:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+720)))
	if v495&int32(1) == int32(0) {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+736)))
	if v500&int32(1) == int32(0) {
		goto L120
	} else {
		goto L123
	}
L123:
	;
	v589 = v483
	v590 = v485 ^ int32(1)
	goto L115
L124:
	;
	if v511 == int32(0) {
		v589 = v507
		v590 = v507
		goto L115
	} else {
		goto L125
	}
L125:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v519 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+464)) = v519
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+468)) = v522
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+472)) = v525
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_13), v41+int32(464))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+452)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v41)+448)) = v142
	F_errhint(m, int32(_a_F_pg_restore_extended_stats_12), v41+int32(448))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v583 = v507
	v584 = v507
	v585 = int32(528)
	goto L116
L129:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+704)))
	if v544 != int32(1) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v557 = int32(0)
	v561 = F_errstart(m, int32(19), v557)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L134
	}
L131:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+720)))
	if v547&int32(1) == int32(0) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+736)))
	if v553&int32(1) != 0 {
		v589 = v483
		v590 = int32(0)
		goto L115
	} else {
		goto L133
	}
L133:
	;
	goto L130
L134:
	;
	if v561 == int32(0) {
		v589 = v557
		v590 = v557
		goto L115
	} else {
		goto L135
	}
L135:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+432)) = v569
	v572 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+436)) = v572
	v575 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+440)) = v575
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_14), v41+int32(432))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v583 = v557
	v584 = v557
	v585 = int32(550)
	goto L116
L138:
	;
	v589 = v583
	v590 = v584
	goto L115
L139:
	;
	v629 = v383 + v410
	v630 = int32(0)
	if (v590|v628)&int32(1) == v630 {
		v838 = v630
		v842 = v630
		v843 = v630
		goto L149
	} else {
		goto L150
	}
L140:
	;
	v627 = v589
	v628 = v80 ^ int32(1)
	goto L139
L141:
	;
	goto L142
L142:
	;
	v596 = int32(0)
	v600 = F_errstart(m, int32(19), v596)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	if v600 == int32(0) {
		v627 = v596
		v628 = v596
		goto L139
	} else {
		goto L144
	}
L144:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v608 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+416)) = v608
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_11), v41+int32(416))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+404)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v41)+400)) = v142
	F_errhint(m, int32(_a_F_pg_restore_extended_stats_12), v41+int32(400))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(566), int32(_a_F_pg_restore_extended_stats_4))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v627 = v596
	v628 = v596
	goto L139
L149:
	;
	v862 = v223 + int32(80)
	v863 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v223))))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+784)) = v863
	*(*int32)(unsafe.Add(mBase, uint32(v41)+778)) = int32(16843009)
	v867 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+776)) = uint16(v867)
	v871 = base.I64_extend_i32_u(base.B2i32(v155 != int64(0)))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+792)) = v871
	if v445&int32(1) == v867 {
		goto L172
	} else {
		goto L173
	}
L150:
	;
	v639 = F_palloc0_mul(m, int32(4), v629)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v642 = F_palloc0_mul(m, int32(4), v629)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v645 = F_palloc0_mul(m, int32(4), v629)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	if int32(0) < v383 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v656 = int32(0)
	goto L157
L155:
	;
	goto L156
L156:
	;
	if v629 <= v383 {
		v838 = v645
		v842 = v639
		v843 = v642
		goto L149
	} else {
		goto L164
	}
L157:
	;
	v695 = int32(*(*int16)(unsafe.Add(mBase, uint32(v223+int32(104)+v656<<(uint(int32(1))%32)))))
	v697 = F_SearchSysCache2(m, int32(7), base.I64_extend_i32_u(v164), base.I64_extend_i32_s(v695))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L159
	}
L158:
	;
	goto L156
L159:
	;
	if v697 == int32(0) {
		goto L61
	} else {
		goto L160
	}
L160:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v697)+16))
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701)+22)))
	v703 = v701 + v702
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703)+91)))
	if v704 == int32(1) {
		goto L60
	} else {
		goto L161
	}
L161:
	;
	v708 = v656 << (uint(int32(2)) % 32)
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v703)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v639+v708))) = v710
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v703)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v708+v642))) = v713
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v703)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v708+v645))) = v716
	F_ReleaseCatCache(m, v697)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v721 = v656 + int32(1)
	if v721 != v383 {
		v656 = v721
		goto L157
	} else {
		goto L163
	}
L163:
	;
	goto L158
L164:
	;
	v765 = v383
	goto L165
L165:
	;
	v800 = int32(2)
	v801 = v765 << (uint(v800) % 32)
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v409)+12))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v803+(v765-v383)<<(uint(v800)%32))))
	v809 = F_exprType(m, v808)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L167
	}
L166:
	;
	v838 = v645
	v842 = v639
	v843 = v642
	goto L149
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v639+v801))) = v809
	v813 = F_exprTypmod(m, v808)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v801+v642))) = v813
	v817 = F_exprCollation(m, v808)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v801+v645))) = v817
	v821 = v765 + int32(1)
	if v821 != v629 {
		v765 = v821
		goto L165
	} else {
		goto L170
	}
L170:
	;
	goto L166
L171:
	;
	if v484&int32(1) == int32(0) {
		goto L219
	} else {
		goto L220
	}
L172:
	;
	v1323 = v627
	goto L171
L173:
	;
	goto L174
L174:
	;
	v878 = *(*int64)(unsafe.Add(mBase, uint32(v41)+664))
	v880 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v878))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v882 = F_statext_ndistinct_deserialize(m, v880)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v884 = int32(0)
	v885 = m.G0
	v887 = v885 - int32(16)
	m.G0 = v887
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v882)+8))
	if v889 == v884 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	m.G0 = v887 + int32(16)
	if v1192 != 0 {
		goto L207
	} else {
		goto L208
	}
L178:
	;
	v1192 = int32(1)
	goto L177
L179:
	;
	goto L180
L180:
	;
	v910 = v884
	goto L181
L181:
	;
	v939 = v882 + int32(16) + v910<<(uint(int32(4))%32)
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v939)+8))
	if int32(0) < v940 {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v1192 = v1182
	goto L177
L183:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v939)+12))
	v950 = int32(0)
	goto L186
L184:
	;
	goto L185
L185:
	;
	v1182 = int32(1)
	v1184 = v910 + v1182
	if v1184 != v889 {
		v910 = v1184
		goto L181
	} else {
		goto L206
	}
L186:
	;
	v986 = int32(*(*int16)(unsafe.Add(mBase, uint32(v943+v950<<(uint(int32(1))%32)))))
	if int32(0) < v986 {
		goto L190
	} else {
		goto L191
	}
L187:
	;
	goto L185
L188:
	;
	v1142 = v950 + int32(1)
	if v1142 != v940 {
		v950 = v1142
		goto L186
	} else {
		goto L205
	}
L189:
	;
	v1082 = int32(0)
	v1085 = F_errstart(m, int32(19), v1082)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L200
	}
L190:
	;
	v989 = int32(0)
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v862)+16))
	if v990 <= v989 {
		goto L189
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	if v986 == int32(0) {
		goto L189
	} else {
		goto L198
	}
L193:
	;
	v999 = v989
	goto L194
L194:
	;
	v1034 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v223+int32(104)+v999<<(uint(int32(1))%32)))))
	if v1034 == v986&int32(_a_F_pg_restore_extended_stats_15) {
		goto L188
	} else {
		goto L196
	}
L195:
	;
	goto L189
L196:
	;
	v1039 = v999 + int32(1)
	if v990 != v1039 {
		v999 = v1039
		goto L194
	} else {
		goto L197
	}
L197:
	;
	goto L195
L198:
	;
	if int32(0)-v410 <= v986 {
		goto L188
	} else {
		goto L199
	}
L199:
	;
	goto L189
L200:
	;
	if v1085 == int32(0) {
		v1192 = v1082
		goto L177
	} else {
		goto L201
	}
L201:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v887)+4)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v887))) = int32(_a_F_pg_restore_extended_stats_16)
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_17), v887)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_18), int32(395), int32(_a_F_pg_restore_extended_stats_19))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v1192 = v1082
	goto L177
L205:
	;
	goto L187
L206:
	;
	goto L182
L207:
	;
	v1227 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+778)) = uint8(v1227)
	*(*int64)(unsafe.Add(mBase, uint32(v41)+800)) = v878
	v1230 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+770)) = uint8(v1230)
	v1232 = v627
	goto L209
L208:
	;
	v1232 = int32(0)
	goto L209
L209:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v882)+8))
	if v1234 != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v1237 = int32(0)
	goto L213
L211:
	;
	goto L212
L212:
	;
	F_pfree(m, v882)
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L1
	} else {
		goto L217
	}
L213:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v882+v1237<<(uint(int32(4))%32))+28))
	F_pfree(m, v1276)
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L1
	} else {
		goto L215
	}
L214:
	;
	goto L212
L215:
	;
	v1280 = v1237 + int32(1)
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v882)+8))
	if base.Ui32(v1280) < base.Ui32(v1281) {
		v1237 = v1280
		goto L213
	} else {
		goto L216
	}
L216:
	;
	goto L214
L217:
	;
	v1323 = v1232
	goto L171
L218:
	;
	if v590&int32(1) == int32(0) {
		v2766 = v1831
		goto L57
	} else {
		goto L265
	}
L219:
	;
	v1831 = v1323
	goto L218
L220:
	;
	goto L221
L221:
	;
	v1366 = *(*int64)(unsafe.Add(mBase, uint32(v41)+680))
	v1368 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v1366))
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	v1370 = F_statext_dependencies_deserialize(m, v1368)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	v1372 = int32(0)
	v1373 = m.G0
	v1375 = v1373 - int32(16)
	m.G0 = v1375
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1370)+8))
	if v1377 == v1372 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	m.G0 = v1375 + int32(16)
	if v1678 != 0 {
		goto L254
	} else {
		goto L255
	}
L225:
	;
	v1678 = int32(1)
	goto L224
L226:
	;
	goto L227
L227:
	;
	v1399 = v1372
	goto L228
L228:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1370+int32(12)+v1399<<(uint(int32(2))%32))))
	v1429 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1428)+8)))
	if int32(0) < v1429 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v1678 = v1672
	goto L224
L230:
	;
	v1441 = int32(0)
	goto L233
L231:
	;
	goto L232
L232:
	;
	v1672 = int32(1)
	v1674 = v1399 + v1672
	if v1674 != v1377 {
		v1399 = v1674
		goto L228
	} else {
		goto L253
	}
L233:
	;
	v1476 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1428+int32(10)+v1441<<(uint(int32(1))%32)))))
	if int32(0) < v1476 {
		goto L237
	} else {
		goto L238
	}
L234:
	;
	goto L232
L235:
	;
	v1632 = v1441 + int32(1)
	if v1632 != v1429 {
		v1441 = v1632
		goto L233
	} else {
		goto L252
	}
L236:
	;
	v1572 = int32(0)
	v1575 = F_errstart(m, int32(19), v1572)
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		goto L1
	} else {
		goto L247
	}
L237:
	;
	v1479 = int32(0)
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v862)+16))
	if v1480 <= v1479 {
		goto L236
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	if v1476 == int32(0) {
		goto L236
	} else {
		goto L245
	}
L240:
	;
	v1485 = v1479
	goto L241
L241:
	;
	v1524 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v223+int32(104)+v1485<<(uint(int32(1))%32)))))
	if v1524 == v1476&int32(_a_F_pg_restore_extended_stats_15) {
		goto L235
	} else {
		goto L243
	}
L242:
	;
	goto L236
L243:
	;
	v1529 = v1485 + int32(1)
	if v1480 != v1529 {
		v1485 = v1529
		goto L241
	} else {
		goto L244
	}
L244:
	;
	goto L242
L245:
	;
	if int32(0)-v410 <= v1476 {
		goto L235
	} else {
		goto L246
	}
L246:
	;
	goto L236
L247:
	;
	if v1575 == int32(0) {
		v1678 = v1572
		goto L224
	} else {
		goto L248
	}
L248:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1375)+4)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v1375))) = int32(_a_F_pg_restore_extended_stats_20)
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_17), v1375)
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_21), int32(649), int32(_a_F_pg_restore_extended_stats_22))
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	v1678 = v1572
	goto L224
L252:
	;
	goto L234
L253:
	;
	goto L229
L254:
	;
	v1717 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+779)) = uint8(v1717)
	*(*int64)(unsafe.Add(mBase, uint32(v41)+808)) = v1366
	v1720 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+771)) = uint8(v1720)
	v1722 = v1323
	goto L256
L255:
	;
	v1722 = int32(0)
	goto L256
L256:
	;
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v1370)+8))
	if v1724 != 0 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1727 = int32(0)
	goto L260
L258:
	;
	goto L259
L259:
	;
	F_pfree(m, v1370)
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		goto L1
	} else {
		goto L264
	}
L260:
	;
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1370+int32(12)+v1727<<(uint(int32(2))%32))))
	F_pfree(m, v1768)
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L1
	} else {
		goto L262
	}
L261:
	;
	goto L259
L262:
	;
	v1772 = v1727 + int32(1)
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v1370)+8))
	if base.Ui32(v1772) < base.Ui32(v1773) {
		v1727 = v1772
		goto L260
	} else {
		goto L263
	}
L263:
	;
	goto L261
L264:
	;
	v1831 = v1722
	goto L218
L265:
	;
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(v41)+696))
	v1858 = F_pg_detoast_datum(m, v1857)
	mBase = m.M
	v1859 = m.ExcPending
	if v1859 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v41)+712))
	v1861 = F_pg_detoast_datum(m, v1860)
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(v41)+728))
	v1864 = F_pg_detoast_datum(m, v1863)
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(v1858)+4))
	if v1866 != int32(2) {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1871 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L1
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1858)+20))
	if v629 != v1889 {
		goto L276
	} else {
		goto L277
	}
L272:
	;
	if v1871 == int32(0) {
		goto L58
	} else {
		goto L273
	}
L273:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+388)) = int32(2)
	v1881 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+384)) = v1881
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_23), v41+int32(384))
	mBase = m.M
	v1887 = m.ExcPending
	if v1887 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	v2740 = int32(834)
	goto L59
L276:
	;
	v1893 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L1
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v1858)+16))
	if int32(_a_F_pg_restore_extended_stats_24) <= v1912 {
		goto L283
	} else {
		goto L284
	}
L279:
	;
	if v1893 == int32(0) {
		goto L58
	} else {
		goto L280
	}
L280:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(v1858)+20))
	v1902 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+368)) = v1902
	*(*int32)(unsafe.Add(mBase, uint32(v41)+372)) = v1900
	*(*int32)(unsafe.Add(mBase, uint32(v41)+376)) = v629
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_25), v41+int32(368))
	mBase = m.M
	v1910 = m.ExcPending
	if v1910 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	v2740 = int32(844)
	goto L59
L283:
	;
	v1917 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L1
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	v1937 = F_check_mcvlist_array(m, v1861, int32(8), v1912)
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L1
	} else {
		goto L290
	}
L286:
	;
	if v1917 == int32(0) {
		goto L58
	} else {
		goto L287
	}
L287:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+356)) = v1912
	*(*int32)(unsafe.Add(mBase, uint32(v41)+360)) = int32(_a_F_pg_restore_extended_stats_26)
	v1928 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+352)) = v1928
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_27), v41+int32(352))
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	v2740 = int32(865)
	goto L59
L290:
	;
	if v1937 == int32(0) {
		goto L58
	} else {
		goto L291
	}
L291:
	;
	v1942 = F_check_mcvlist_array(m, v1864, int32(9), v1912)
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	if v1942 == int32(0) {
		goto L58
	} else {
		goto L293
	}
L293:
	;
	F_deconstruct_array_builtin(m, v1858, int32(25), v41+int32(880), v41+int32(1360), v41+int32(1616))
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(v41)+880))
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1360))
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1861)+8))
	if v1957 != 0 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1965 = v1957
	goto L297
L296:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v1861)+4))
	v1965 = (v1958<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L297
L297:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v1864)+8))
	if v1967 != 0 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1975 = v1967
	goto L300
L299:
	;
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v1864)+4))
	v1975 = (v1968<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L300
L300:
	;
	v1977 = int32(0)
	v1979 = m.G0
	v1981 = v1979 - int32(80)
	m.G0 = v1981
	v1987 = F_palloc0(m, v1912*int32(24)+int32(48))
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1987)+12)) = uint16(v629)
	*(*int64)(unsafe.Add(mBase, uint32(v1987))) = int64(8080740802)
	*(*int32)(unsafe.Add(mBase, uint32(v1987)+8)) = v1912
	if int32(0) < v1912 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1997 = v1977
	goto L305
L303:
	;
	goto L304
L304:
	;
	if v629 <= int32(0) {
		goto L314
	} else {
		goto L315
	}
L305:
	;
	v2037 = v1987 + int32(48) + v1997*int32(24)
	v2039 = v1997 << (uint(int32(3)) % 32)
	v2041 = *(*float64)(unsafe.Add(mBase, uint32(v1965+v1861+v2039)))
	*(*float64)(unsafe.Add(mBase, uint32(v2037))) = v2041
	v2044 = *(*float64)(unsafe.Add(mBase, uint32(v1975+v1864+v2039)))
	*(*float64)(unsafe.Add(mBase, uint32(v2037)+8)) = v2044
	v2047 = F_palloc0_mul(m, int32(8), v629)
	mBase = m.M
	v2048 = m.ExcPending
	if v2048 != 0 {
		goto L1
	} else {
		goto L307
	}
L306:
	;
	goto L304
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2037)+20)) = v2047
	v2051 = F_palloc0_mul(m, int32(1), v629)
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2037)+16)) = v2051
	v2055 = v1997 + int32(1)
	if v2055 != v1912 {
		v1997 = v2055
		goto L305
	} else {
		goto L309
	}
L309:
	;
	goto L306
L310:
	;
	m.G0 = v1981 + int32(80)
	if v2684 == int64(0) {
		goto L382
	} else {
		goto L383
	}
L311:
	;
	v2684 = base.I64_extend_i32_u(v2452)
	goto L310
L312:
	;
	v2549 = *(*int32)(unsafe.Add(mBase, uint32(v1987)+8))
	if v2549 != 0 {
		goto L373
	} else {
		goto L374
	}
L313:
	;
	F_pfree(m, v2451)
	mBase = m.M
	v2488 = m.ExcPending
	if v2488 != 0 {
		goto L1
	} else {
		goto L364
	}
L314:
	;
	v2098 = F_palloc0_mul(m, int32(4), v629)
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L1
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	v2110 = v1977
	goto L319
L317:
	;
	v2100 = F_statext_mcv_serialize(m, v1987, v2098)
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	v2451 = v2098
	v2452 = v2100
	goto L313
L319:
	;
	v2147 = v2110 << (uint(int32(2)) % 32)
	v2148 = v842 + v2147
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(v2148)))
	F_getTypeInputInfo(m, v2149, v1981+int32(44), v1981+int32(48))
	mBase = m.M
	v2155 = m.ExcPending
	if v2155 != 0 {
		goto L1
	} else {
		goto L321
	}
L320:
	;
	v2317 = F_palloc0_mul(m, int32(4), v629)
	mBase = m.M
	v2318 = m.ExcPending
	if v2318 != 0 {
		goto L1
	} else {
		goto L348
	}
L321:
	;
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(v1981)+44))
	F_fmgr_info(m, v2156, v1981+int32(52))
	mBase = m.M
	v2160 = m.ExcPending
	if v2160 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(v2148)))
	*(*int32)(unsafe.Add(mBase, uint32(v2147+(v1987+int32(16))))) = v2162
	if base.B2i32(v1912 <= int32(0)) == int32(0) {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v2169 = v2110 << (uint(int32(3)) % 32)
	v2170 = v2110
	v2173 = int32(0)
	goto L326
L324:
	;
	goto L325
L325:
	;
	v2313 = v2110 + int32(1)
	if v2313 != v629 {
		v2110 = v2313
		goto L319
	} else {
		goto L347
	}
L326:
	;
	v2210 = v1987 + int32(48) + v2173*int32(24)
	v2212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2170+v1956))))
	if v2212 == int32(1) {
		goto L329
	} else {
		goto L330
	}
L327:
	;
	goto L325
L328:
	;
	v2272 = v2173 + int32(1)
	if v2272 != v1912 {
		v2170 = v2170 + v629
		v2173 = v2272
		goto L326
	} else {
		goto L346
	}
L329:
	;
	v2215 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v2215+v2169))) = int64(0)
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+16))
	v2221 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2219+v2110))) = uint8(v2221)
	goto L328
L330:
	;
	goto L331
L331:
	;
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v1955+v2170<<(uint(int32(3))%32))))
	v2227 = F_text_to_cstring(m, v2226)
	mBase = m.M
	v2228 = m.ExcPending
	if v2228 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	v2230 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v1981)+40)) = v2230
	v2233 = *(*int64)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v1981)+32)) = v2233
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v1981)+48))
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(v2147+v843)))
	v2241 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+20))
	v2243 = F_InputFunctionCallSafe(m, v1981+int32(52), v2227, v2237, v2238, v1981+int32(32), v2241+v2169)
	mBase = m.M
	v2244 = m.ExcPending
	if v2244 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	if v2243 == int32(0) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v2249 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v2250 = m.ExcPending
	if v2250 != 0 {
		goto L1
	} else {
		goto L337
	}
L335:
	;
	goto L336
L336:
	;
	F_pfree(m, v2227)
	mBase = m.M
	v2268 = m.ExcPending
	if v2268 != 0 {
		goto L1
	} else {
		goto L345
	}
L337:
	;
	if v2249 != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2253 = m.ExcPending
	if v2253 != 0 {
		goto L1
	} else {
		goto L341
	}
L339:
	;
	goto L340
L340:
	;
	F_pfree(m, v2227)
	mBase = m.M
	v2266 = m.ExcPending
	if v2266 != 0 {
		goto L1
	} else {
		goto L344
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1981)+16)) = v2227
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_28), v1981+int32(16))
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_29), int32(2274), int32(_a_F_pg_restore_extended_stats_30))
	mBase = m.M
	v2264 = m.ExcPending
	if v2264 != 0 {
		goto L1
	} else {
		goto L343
	}
L343:
	;
	goto L340
L344:
	;
	goto L312
L345:
	;
	goto L328
L346:
	;
	goto L327
L347:
	;
	goto L320
L348:
	;
	v2319 = int32(0)
	goto L350
L349:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2439 = m.ExcPending
	if v2439 != 0 {
		goto L1
	} else {
		goto L361
	}
L350:
	;
	v2359 = v2319 << (uint(int32(2)) % 32)
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(v842+v2359)))
	v2364 = F_SearchSysCacheCopy(m, int32(82), base.I64_extend_i32_u(v2361), int64(0))
	mBase = m.M
	v2365 = m.ExcPending
	if v2365 != 0 {
		goto L1
	} else {
		goto L352
	}
L351:
	;
	v2387 = F_statext_mcv_serialize(m, v1987, v2317)
	mBase = m.M
	v2388 = m.ExcPending
	if v2388 != 0 {
		goto L1
	} else {
		goto L356
	}
L352:
	;
	if v2364 == int32(0) {
		goto L349
	} else {
		goto L353
	}
L353:
	;
	v2368 = v2317 + v2359
	v2370 = F_palloc0(m, int32(248))
	mBase = m.M
	v2371 = m.ExcPending
	if v2371 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2368))) = v2370
	v2373 = *(*int32)(unsafe.Add(mBase, uint32(v2364)+16))
	v2374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2373)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v2370)+12)) = v2373 + v2374
	v2377 = *(*int32)(unsafe.Add(mBase, uint32(v2368)))
	*(*int32)(unsafe.Add(mBase, uint32(v2377)+4)) = v2361
	v2379 = *(*int32)(unsafe.Add(mBase, uint32(v2368)))
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(v2359+v838)))
	*(*int32)(unsafe.Add(mBase, uint32(v2379)+16)) = v2381
	v2384 = v2319 + int32(1)
	if v2384 != v629 {
		v2319 = v2384
		goto L350
	} else {
		goto L355
	}
L355:
	;
	goto L351
L356:
	;
	v2389 = int32(0)
	goto L357
L357:
	;
	v2430 = *(*int32)(unsafe.Add(mBase, uint32(v2317+v2389<<(uint(int32(2))%32))))
	F_pfree(m, v2430)
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L1
	} else {
		goto L359
	}
L358:
	;
	v2451 = v2317
	v2452 = v2387
	goto L313
L359:
	;
	v2434 = v2389 + int32(1)
	if v2434 != v629 {
		v2389 = v2434
		goto L357
	} else {
		goto L360
	}
L360:
	;
	goto L358
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1981))) = v2361
	F_errmsg_internal(m, int32(_a_F_pg_restore_extended_stats_31), v1981)
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_29), int32(2301), int32(_a_F_pg_restore_extended_stats_30))
	mBase = m.M
	v2448 = m.ExcPending
	if v2448 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L364:
	;
	F_pfree(m, v1955)
	mBase = m.M
	v2490 = m.ExcPending
	if v2490 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	F_pfree(m, v1956)
	mBase = m.M
	v2492 = m.ExcPending
	if v2492 != 0 {
		goto L1
	} else {
		goto L366
	}
L366:
	;
	if v2452 != 0 {
		goto L311
	} else {
		goto L367
	}
L367:
	;
	v2495 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v2496 = m.ExcPending
	if v2496 != 0 {
		goto L1
	} else {
		goto L368
	}
L368:
	;
	if v2495 == int32(0) {
		goto L312
	} else {
		goto L369
	}
L369:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2501 = m.ExcPending
	if v2501 != 0 {
		goto L1
	} else {
		goto L370
	}
L370:
	;
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_32), int32(0))
	mBase = m.M
	v2505 = m.ExcPending
	if v2505 != 0 {
		goto L1
	} else {
		goto L371
	}
L371:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_29), int32(2325), int32(_a_F_pg_restore_extended_stats_30))
	mBase = m.M
	v2510 = m.ExcPending
	if v2510 != 0 {
		goto L1
	} else {
		goto L372
	}
L372:
	;
	goto L312
L373:
	;
	v2553 = int32(0)
	goto L376
L374:
	;
	goto L375
L375:
	;
	F_pfree(m, v1987)
	mBase = m.M
	v2643 = m.ExcPending
	if v2643 != 0 {
		goto L1
	} else {
		goto L381
	}
L376:
	;
	v2593 = v1987 + int32(48) + v2553*int32(24)
	v2594 = *(*int32)(unsafe.Add(mBase, uint32(v2593)+20))
	F_pfree(m, v2594)
	mBase = m.M
	v2596 = m.ExcPending
	if v2596 != 0 {
		goto L1
	} else {
		goto L378
	}
L377:
	;
	goto L375
L378:
	;
	v2597 = *(*int32)(unsafe.Add(mBase, uint32(v2593)+16))
	F_pfree(m, v2597)
	mBase = m.M
	v2599 = m.ExcPending
	if v2599 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	v2601 = v2553 + int32(1)
	v2602 = *(*int32)(unsafe.Add(mBase, uint32(v1987)+8))
	if base.Ui32(v2601) < base.Ui32(v2602) {
		v2553 = v2601
		goto L376
	} else {
		goto L380
	}
L380:
	;
	goto L377
L381:
	;
	v2684 = int64(0)
	goto L310
L382:
	;
	v2766 = int32(0)
	goto L57
L383:
	;
	goto L384
L384:
	;
	v2691 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+780)) = uint8(v2691)
	*(*int64)(unsafe.Add(mBase, uint32(v41)+816)) = v2684
	v2694 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+772)) = uint8(v2694)
	v2766 = v1831
	goto L57
L385:
	;
	F_errmsg_internal(m, int32(_a_F_pg_restore_extended_stats_33), int32(0))
	mBase = m.M
	v2703 = m.ExcPending
	if v2703 != 0 {
		goto L1
	} else {
		goto L386
	}
L386:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(237), int32(_a_F_pg_restore_extended_stats_10))
	mBase = m.M
	v2708 = m.ExcPending
	if v2708 != 0 {
		goto L1
	} else {
		goto L387
	}
L387:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v695
	F_errmsg_internal(m, int32(_a_F_pg_restore_extended_stats_34), v41+int32(48))
	mBase = m.M
	v2718 = m.ExcPending
	if v2718 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(603), int32(_a_F_pg_restore_extended_stats_4))
	mBase = m.M
	v2723 = m.ExcPending
	if v2723 != 0 {
		goto L1
	} else {
		goto L390
	}
L390:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v695
	F_errmsg_internal(m, int32(_a_F_pg_restore_extended_stats_35), v41-int32(-64))
	mBase = m.M
	v2733 = m.ExcPending
	if v2733 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(608), int32(_a_F_pg_restore_extended_stats_4))
	mBase = m.M
	v2738 = m.ExcPending
	if v2738 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L394:
	;
	goto L58
L395:
	;
	v4345 = F_table_open(m, int32(3429), int32(3))
	mBase = m.M
	v4346 = m.ExcPending
	if v4346 != 0 {
		goto L1
	} else {
		goto L693
	}
L396:
	;
	F_relation_close(m, v2792, int32(3))
	mBase = m.M
	v4302 = m.ExcPending
	if v4302 != 0 {
		goto L1
	} else {
		goto L692
	}
L397:
	;
	v2792 = F_table_open(m, int32(2619), int32(3))
	mBase = m.M
	v2793 = m.ExcPending
	if v2793 != 0 {
		goto L1
	} else {
		goto L400
	}
L398:
	;
	goto L399
L399:
	;
	v4305 = v41
	v4319 = v838
	v4323 = v842
	v4324 = v843
	v4330 = v194
	v4332 = v192
	v4333 = v54
	v4342 = v2766
	goto L395
L400:
	;
	v2794 = *(*int32)(unsafe.Add(mBase, uint32(v41)+744))
	v2795 = F_pg_detoast_datum(m, v2794)
	mBase = m.M
	v2796 = m.ExcPending
	if v2796 != 0 {
		goto L1
	} else {
		goto L401
	}
L401:
	;
	v2798 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[7]))
	v2800 = F_get_rel_type_id(m, int32(2619))
	mBase = m.M
	v2801 = m.ExcPending
	if v2801 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	v2802 = *(*int32)(unsafe.Add(mBase, uint32(v2795)+4))
	if v2802&int32(1073741824) == int32(0) {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v2809 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v2810 = m.ExcPending
	if v2810 != 0 {
		goto L1
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	if v2802&int32(268435455) != v410 {
		goto L411
	} else {
		goto L412
	}
L406:
	;
	if v2809 == int32(0) {
		goto L396
	} else {
		goto L407
	}
L407:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2815 = m.ExcPending
	if v2815 != 0 {
		goto L1
	} else {
		goto L408
	}
L408:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v2798
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_36), v41+int32(80))
	mBase = m.M
	v2821 = m.ExcPending
	if v2821 != 0 {
		goto L1
	} else {
		goto L409
	}
L409:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(1584), int32(_a_F_pg_restore_extended_stats_37))
	mBase = m.M
	v2826 = m.ExcPending
	if v2826 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	goto L396
L411:
	;
	v2832 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v2833 = m.ExcPending
	if v2833 != 0 {
		goto L1
	} else {
		goto L414
	}
L412:
	;
	goto L413
L413:
	;
	F_fmgr_info(m, int32(750), v41+int32(840))
	mBase = m.M
	v2855 = m.ExcPending
	if v2855 != 0 {
		goto L1
	} else {
		goto L419
	}
L414:
	;
	if v2832 == int32(0) {
		goto L396
	} else {
		goto L415
	}
L415:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2838 = m.ExcPending
	if v2838 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+340)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v41)+336)) = v2798
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_38), v41+int32(336))
	mBase = m.M
	v2845 = m.ExcPending
	if v2845 != 0 {
		goto L1
	} else {
		goto L417
	}
L417:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(1600), int32(_a_F_pg_restore_extended_stats_37))
	mBase = m.M
	v2850 = m.ExcPending
	if v2850 != 0 {
		goto L1
	} else {
		goto L418
	}
L418:
	;
	goto L396
L419:
	;
	if v410 == int32(0) {
		goto L421
	} else {
		goto L422
	}
L420:
	;
	v4219 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+773)) = uint8(v4219)
	*(*int64)(unsafe.Add(mBase, uint32(v41)+824)) = v4215
	v4222 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+781)) = uint8(v4222)
	goto L399
L421:
	;
	F_relation_close(m, v2792, int32(3))
	mBase = m.M
	v2860 = m.ExcPending
	if v2860 != 0 {
		goto L1
	} else {
		goto L424
	}
L422:
	;
	goto L423
L423:
	;
	v2863 = v383 << (uint(int32(2)) % 32)
	v2877 = int32(0)
	v2887 = v2877
	v2888 = v2877
	v2892 = v2877
	goto L425
L424:
	;
	v4215 = int64(0)
	goto L420
L425:
	;
	v2920 = base.I32_extend16_s(v2887 ^ int32(-1))
	v2921 = int32(1)
	v2922 = int64(0)
	v2924 = F_getIthJsonbValueFromContainer(m, v2795+int32(4), v2887)
	mBase = m.M
	v2925 = m.ExcPending
	if v2925 != 0 {
		goto L1
	} else {
		goto L428
	}
L426:
	;
	if v4165 != 0 {
		goto L686
	} else {
		goto L687
	}
L427:
	;
	v4162 = v4127 + v2892
	v4164 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[10]))
	v4165 = F_accumArrayResult(m, v2888, v4158, v4126, v2800, v4164)
	mBase = m.M
	v4166 = m.ExcPending
	if v4166 != 0 {
		goto L1
	} else {
		goto L684
	}
L428:
	;
	v2926 = *(*int32)(unsafe.Add(mBase, uint32(v2924)))
	if v2926 == int32(0) {
		v4126 = v2921
		v4127 = v2921
		v4158 = v2922
		goto L427
	} else {
		goto L429
	}
L429:
	;
	if v2926 == int32(18) {
		goto L431
	} else {
		goto L432
	}
L430:
	;
	v4126 = int32(1)
	v4127 = int32(0)
	v4158 = int64(0)
	goto L427
L431:
	;
	v2932 = v2887 << (uint(int32(2)) % 32)
	v2934 = *(*int32)(unsafe.Add(mBase, uint32(v838+v2863+v2932)))
	v2936 = *(*int32)(unsafe.Add(mBase, uint32(v2932+(v2863+v843))))
	v2938 = *(*int32)(unsafe.Add(mBase, uint32(v2932+(v2863+v842))))
	v2939 = int32(0)
	v2941 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[7]))
	v2942 = *(*int32)(unsafe.Add(mBase, uint32(v2924)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+1324)) = v2939
	*(*int32)(unsafe.Add(mBase, uint32(v41)+1320)) = v2939
	v2947 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v41)+1309)) = v2947
	*(*int64)(unsafe.Add(mBase, uint32(v41)+1304)) = v2947
	base.MemoryFill(m, v41+int32(880), v2939, int32(416))
	v2956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2942)+3)))
	if v2956&int32(32) == v2939 {
		goto L434
	} else {
		goto L435
	}
L432:
	;
	goto L433
L433:
	;
	v4062 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4063 = m.ExcPending
	if v4063 != 0 {
		goto L1
	} else {
		goto L675
	}
L434:
	;
	v2963 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v2964 = m.ExcPending
	if v2964 != 0 {
		goto L1
	} else {
		goto L437
	}
L435:
	;
	goto L436
L436:
	;
	v2985 = v2939
	goto L442
L437:
	;
	if v2963 == int32(0) {
		v4126 = v2921
		v4127 = v2939
		v4158 = v2922
		goto L427
	} else {
		goto L438
	}
L438:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2969 = m.ExcPending
	if v2969 != 0 {
		goto L1
	} else {
		goto L439
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+116)) = v2920
	*(*int32)(unsafe.Add(mBase, uint32(v41)+112)) = v2941
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_39), v41+int32(112))
	mBase = m.M
	v2976 = m.ExcPending
	if v2976 != 0 {
		goto L1
	} else {
		goto L440
	}
L440:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(1137), int32(_a_F_pg_restore_extended_stats_40))
	mBase = m.M
	v2981 = m.ExcPending
	if v2981 != 0 {
		goto L1
	} else {
		goto L441
	}
L441:
	;
	v4126 = v2921
	v4127 = v2939
	v4158 = v2922
	goto L427
L442:
	;
	v3022 = *(*int32)(unsafe.Add(mBase, uint32(v2985<<(uint(int32(2))%32))+uint32(_c_F_pg_restore_extended_stats[11])))
	v3023 = F_strlen(m, v3022)
	mBase = m.M
	v3028 = v41 + int32(880) + v2985<<(uint(int32(5))%32)
	v3029 = F_getKeyJsonValueFromContainer(m, v2942, v3022, v3023, v3028)
	mBase = m.M
	v3030 = m.ExcPending
	if v3030 != 0 {
		goto L1
	} else {
		goto L445
	}
L443:
	;
	v3072 = F_JsonbIteratorInit(m, v2942)
	mBase = m.M
	v3073 = m.ExcPending
	if v3073 != 0 {
		goto L1
	} else {
		goto L456
	}
L444:
	;
	v3069 = v2985 + int32(1)
	if v3069 != int32(13) {
		v2985 = v3069
		goto L442
	} else {
		goto L455
	}
L445:
	;
	if v3029 == int32(0) {
		goto L444
	} else {
		goto L446
	}
L446:
	;
	v3033 = *(*int32)(unsafe.Add(mBase, uint32(v3028)))
	switch v3033 {
	case 0:
		goto L444
	case 1:
		goto L447
	default:
		goto L448
	}
L447:
	;
	v3066 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(1304)+v2985))) = uint8(v3066)
	goto L444
L448:
	;
	v3034 = int32(1)
	v3035 = int32(0)
	v3038 = F_errstart(m, int32(19), v3035)
	mBase = m.M
	v3039 = m.ExcPending
	if v3039 != 0 {
		goto L1
	} else {
		goto L449
	}
L449:
	;
	if v3038 == int32(0) {
		v4126 = v3034
		v4127 = v3035
		v4158 = v2922
		goto L427
	} else {
		goto L450
	}
L450:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v3044 = m.ExcPending
	if v3044 != 0 {
		goto L1
	} else {
		goto L451
	}
L451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+324)) = v2920
	*(*int32)(unsafe.Add(mBase, uint32(v41)+320)) = v2941
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_39), v41+int32(320))
	mBase = m.M
	v3051 = m.ExcPending
	if v3051 != 0 {
		goto L1
	} else {
		goto L452
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+304)) = v3022
	F_errhint(m, int32(_a_F_pg_restore_extended_stats_41), v41+int32(304))
	mBase = m.M
	v3057 = m.ExcPending
	if v3057 != 0 {
		goto L1
	} else {
		goto L453
	}
L453:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(1167), int32(_a_F_pg_restore_extended_stats_40))
	mBase = m.M
	v3062 = m.ExcPending
	if v3062 != 0 {
		goto L1
	} else {
		goto L454
	}
L454:
	;
	v4126 = v3034
	v4127 = v3035
	v4158 = v2922
	goto L427
L455:
	;
	goto L443
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+1328)) = v3072
	v3076 = v41 + int32(1328)
	v3078 = v41 + int32(1360)
	v3080 = F_JsonbIteratorNext(m, v3076, v3078, int32(0))
	mBase = m.M
	v3081 = m.ExcPending
	if v3081 != 0 {
		goto L1
	} else {
		goto L457
	}
L457:
	;
	v3083 = F_JsonbIteratorNext(m, v3076, v3078, int32(0))
	mBase = m.M
	v3084 = m.ExcPending
	if v3084 != 0 {
		goto L1
	} else {
		goto L458
	}
L458:
	;
	if v3083 != int32(7) {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	goto L462
L460:
	;
	goto L461
L461:
	;
	v3355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1307)))
	v3356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1308)))
	if v3355 != v3356 {
		goto L505
	} else {
		goto L506
	}
L462:
	;
	v3125 = int32(0)
	v3131 = F_JsonbIteratorNext(m, v41+int32(1328), v41+int32(1616), v3125)
	mBase = m.M
	v3132 = m.ExcPending
	if v3132 != 0 {
		goto L1
	} else {
		goto L464
	}
L463:
	;
	goto L461
L464:
	;
	v3133 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1372))
	v3134 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1368))
	v3138 = v3125
	goto L466
L465:
	;
	v3313 = F_JsonbIteratorNext(m, v41+int32(1328), v41+int32(1360), int32(0))
	mBase = m.M
	v3314 = m.ExcPending
	if v3314 != 0 {
		goto L1
	} else {
		goto L503
	}
L466:
	;
	v3175 = *(*int32)(unsafe.Add(mBase, uint32(v3138<<(uint(int32(2))%32))+uint32(_c_F_pg_restore_extended_stats[11])))
	v3176 = F_strlen(m, v3175)
	mBase = m.M
	if v3176 == v3134 {
		goto L468
	} else {
		goto L469
	}
L467:
	;
	v3231 = F_palloc0(m, v3134+int32(1))
	mBase = m.M
	v3232 = m.ExcPending
	if v3232 != 0 {
		goto L1
	} else {
		goto L486
	}
L468:
	;
	if v3134 == int32(0) {
		goto L472
	} else {
		goto L473
	}
L469:
	;
	goto L470
L470:
	;
	v3226 = v3138 + int32(1)
	if v3226 != int32(13) {
		v3138 = v3226
		goto L466
	} else {
		goto L485
	}
L471:
	;
	if v3222 == int32(0) {
		goto L465
	} else {
		goto L484
	}
L472:
	;
	v3222 = int32(0)
	goto L471
L473:
	;
	goto L474
L474:
	;
	v3183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3175))))
	if v3183 != 0 {
		goto L475
	} else {
		goto L476
	}
L475:
	;
	v3184 = v3175
	v3185 = v3133
	v3186 = v3134
	v3187 = v3183
	goto L479
L476:
	;
	v3210 = v3133
	v3214 = int32(0)
	goto L477
L477:
	;
	v3215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3210))))
	v3222 = v3214 - v3215
	goto L471
L478:
	;
	v3210 = v3205
	v3214 = v3207
	goto L477
L479:
	;
	v3189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3185))))
	if base.B2i32(v3187 != v3189)|base.B2i32(v3189 == int32(0)) != 0 {
		v3205 = v3185
		v3207 = v3187
		goto L478
	} else {
		goto L481
	}
L480:
	;
	v3205 = v3199
	v3207 = int32(0)
	goto L478
L481:
	;
	v3195 = v3186 - int32(1)
	if v3195 == int32(0) {
		v3205 = v3185
		v3207 = v3187
		goto L478
	} else {
		goto L482
	}
L482:
	;
	v3198 = int32(1)
	v3199 = v3185 + v3198
	v3200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3184)+1)))
	if v3200 != 0 {
		v3184 = v3184 + v3198
		v3185 = v3199
		v3186 = v3195
		v3187 = v3200
		goto L479
	} else {
		goto L483
	}
L483:
	;
	goto L480
L484:
	;
	goto L470
L485:
	;
	goto L467
L486:
	;
	v3233 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1368))
	if v3233 != 0 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v3234 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1372))
	base.MemoryCopy(m, v3231, v3234, v3233)
	goto L489
L488:
	;
	goto L489
L489:
	;
	v3238 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v3239 = m.ExcPending
	if v3239 != 0 {
		goto L1
	} else {
		goto L490
	}
L490:
	;
	if v3238 != 0 {
		goto L491
	} else {
		goto L492
	}
L491:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v3242 = m.ExcPending
	if v3242 != 0 {
		goto L1
	} else {
		goto L494
	}
L492:
	;
	goto L493
L493:
	;
	F_pfree(m, v3231)
	mBase = m.M
	v3255 = m.ExcPending
	if v3255 != 0 {
		goto L1
	} else {
		goto L497
	}
L494:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+288)) = v2920
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_42), v41+int32(288))
	mBase = m.M
	v3248 = m.ExcPending
	if v3248 != 0 {
		goto L1
	} else {
		goto L495
	}
L495:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(968), int32(_a_F_pg_restore_extended_stats_43))
	mBase = m.M
	v3253 = m.ExcPending
	if v3253 != 0 {
		goto L1
	} else {
		goto L496
	}
L496:
	;
	goto L493
L497:
	;
	goto L498
L498:
	;
	v3295 = v41 + int32(1328)
	v3299 = F_JsonbIteratorNext(m, v3295, v41+int32(1360), int32(0))
	mBase = m.M
	v3300 = m.ExcPending
	if v3300 != 0 {
		goto L1
	} else {
		goto L500
	}
L500:
	;
	if v3299 == int32(7) {
		goto L430
	} else {
		goto L501
	}
L501:
	;
	v3306 = F_JsonbIteratorNext(m, v3295, v41+int32(1616), int32(0))
	mBase = m.M
	v3307 = m.ExcPending
	if v3307 != 0 {
		goto L1
	} else {
		goto L502
	}
L502:
	;
	goto L498
L503:
	;
	if v3313 != int32(7) {
		goto L462
	} else {
		goto L504
	}
L504:
	;
	goto L463
L505:
	;
	v3358 = int32(1)
	v3359 = int32(0)
	v3362 = F_errstart(m, int32(19), v3359)
	mBase = m.M
	v3363 = m.ExcPending
	if v3363 != 0 {
		goto L1
	} else {
		goto L508
	}
L506:
	;
	goto L507
L507:
	;
	v3390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1311)))
	v3391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1312)))
	if v3390 != v3391 {
		goto L514
	} else {
		goto L515
	}
L508:
	;
	if v3362 == int32(0) {
		v4126 = v3358
		v4127 = v3359
		v4158 = v2922
		goto L427
	} else {
		goto L509
	}
L509:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v3368 = m.ExcPending
	if v3368 != 0 {
		goto L1
	} else {
		goto L510
	}
L510:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+276)) = v2920
	*(*int32)(unsafe.Add(mBase, uint32(v41)+272)) = v2941
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_39), v41+int32(272))
	mBase = m.M
	v3375 = m.ExcPending
	if v3375 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+260)) = int32(_a_F_pg_restore_extended_stats_44)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+256)) = int32(_a_F_pg_restore_extended_stats_45)
	F_errhint(m, int32(_a_F_pg_restore_extended_stats_46), v41+int32(256))
	mBase = m.M
	v3384 = m.ExcPending
	if v3384 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(1190), int32(_a_F_pg_restore_extended_stats_40))
	mBase = m.M
	v3389 = m.ExcPending
	if v3389 != 0 {
		goto L1
	} else {
		goto L513
	}
L513:
	;
	v4126 = v3358
	v4127 = v3359
	v4158 = v2922
	goto L427
L514:
	;
	v3393 = int32(1)
	v3394 = int32(0)
	v3397 = F_errstart(m, int32(19), v3394)
	mBase = m.M
	v3398 = m.ExcPending
	if v3398 != 0 {
		goto L1
	} else {
		goto L517
	}
L515:
	;
	goto L516
L516:
	;
	v3425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1314)))
	v3426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1315)))
	if v3425 == v3426 {
		goto L524
	} else {
		goto L525
	}
L517:
	;
	if v3397 == int32(0) {
		v4126 = v3393
		v4127 = v3394
		v4158 = v2922
		goto L427
	} else {
		goto L518
	}
L518:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v3403 = m.ExcPending
	if v3403 != 0 {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+244)) = v2920
	*(*int32)(unsafe.Add(mBase, uint32(v41)+240)) = v2941
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_39), v41+int32(240))
	mBase = m.M
	v3410 = m.ExcPending
	if v3410 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+228)) = int32(_a_F_pg_restore_extended_stats_47)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+224)) = int32(_a_F_pg_restore_extended_stats_48)
	F_errhint(m, int32(_a_F_pg_restore_extended_stats_46), v41+int32(224))
	mBase = m.M
	v3419 = m.ExcPending
	if v3419 != 0 {
		goto L1
	} else {
		goto L521
	}
L521:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(1201), int32(_a_F_pg_restore_extended_stats_40))
	mBase = m.M
	v3424 = m.ExcPending
	if v3424 != 0 {
		goto L1
	} else {
		goto L522
	}
L522:
	;
	v4126 = v3393
	v4127 = v3394
	v4158 = v2922
	goto L427
L523:
	;
	v3465 = F_lookup_type_cache(m, v2938, int32(3))
	mBase = m.M
	v3466 = m.ExcPending
	if v3466 != 0 {
		goto L1
	} else {
		goto L534
	}
L524:
	;
	v3428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1316)))
	if v3425 == v3428 {
		goto L523
	} else {
		goto L527
	}
L525:
	;
	goto L526
L526:
	;
	v3430 = int32(1)
	v3431 = int32(0)
	v3434 = F_errstart(m, int32(19), v3431)
	mBase = m.M
	v3435 = m.ExcPending
	if v3435 != 0 {
		goto L1
	} else {
		goto L528
	}
L527:
	;
	goto L526
L528:
	;
	if v3434 == int32(0) {
		v4126 = v3430
		v4127 = v3431
		v4158 = v2922
		goto L427
	} else {
		goto L529
	}
L529:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v3440 = m.ExcPending
	if v3440 != 0 {
		goto L1
	} else {
		goto L530
	}
L530:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+212)) = v2920
	*(*int32)(unsafe.Add(mBase, uint32(v41)+208)) = v2941
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_39), v41+int32(208))
	mBase = m.M
	v3447 = m.ExcPending
	if v3447 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+200)) = int32(_a_F_pg_restore_extended_stats_49)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+196)) = int32(_a_F_pg_restore_extended_stats_50)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+192)) = int32(_a_F_pg_restore_extended_stats_51)
	F_errhint(m, int32(_a_F_pg_restore_extended_stats_52), v41+int32(192))
	mBase = m.M
	v3458 = m.ExcPending
	if v3458 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(1219), int32(_a_F_pg_restore_extended_stats_40))
	mBase = m.M
	v3463 = m.ExcPending
	if v3463 != 0 {
		goto L1
	} else {
		goto L533
	}
L533:
	;
	v4126 = v3430
	v4127 = v3431
	v4158 = v2922
	goto L427
L534:
	;
	v3467 = int32(0)
	v3471 = v41 + int32(1360)
	v3473 = v41 + int32(1616)
	v3475 = v41 + int32(1328)
	v3476 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v3473)+29)) = uint16(v3476)
	v3478 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v3473)+21)) = v3478
	*(*int64)(unsafe.Add(mBase, uint32(v3475)+23)) = v3478
	*(*int64)(unsafe.Add(mBase, uint32(v3475)+16)) = v3478
	*(*int64)(unsafe.Add(mBase, uint32(v3475)+8)) = v3478
	*(*int64)(unsafe.Add(mBase, uint32(v3475))) = v3478
	v3488 = base.I64_extend_i32_u(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471))) = v3488
	*(*uint8)(unsafe.Add(mBase, uint32(v3473))) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+8)) = base.I64_extend_i32_s(v3467)
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+1)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+16)) = v3488
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+2)) = uint8(v3467)
	v3500 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+24)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+3)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+32)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+4)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+40)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+5)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+48)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+6)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+88)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+11)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+128)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+16)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+56)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+7)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+96)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+12)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+136)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+17)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+64)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+8)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+104)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+13)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+144)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+18)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+72)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+9)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+112)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+14)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+152)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+19)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+80)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+10)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+120)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+15)) = uint8(v3467)
	*(*int64)(unsafe.Add(mBase, uint32(v3471)+160)) = v3500
	*(*uint8)(unsafe.Add(mBase, uint32(v3473)+20)) = uint8(v3467)
	goto L535
L535:
	;
	v3572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1313)))
	if v3390|v3572&int32(1) == int32(0) {
		goto L536
	} else {
		goto L537
	}
L536:
	;
	if v3425 == int32(0) {
		goto L545
	} else {
		goto L546
	}
L537:
	;
	v3583 = F_statatt_get_elem_type(m, v2938, v41+int32(1324), v41+int32(1320))
	mBase = m.M
	v3584 = m.ExcPending
	if v3584 != 0 {
		goto L1
	} else {
		goto L538
	}
L538:
	;
	if v3583 != 0 {
		goto L536
	} else {
		goto L539
	}
L539:
	;
	v3585 = int32(1)
	v3586 = int32(0)
	v3589 = F_errstart(m, int32(19), v3586)
	mBase = m.M
	v3590 = m.ExcPending
	if v3590 != 0 {
		goto L1
	} else {
		goto L540
	}
L540:
	;
	if v3589 == int32(0) {
		v4126 = v3585
		v4127 = v3586
		v4158 = v2922
		goto L427
	} else {
		goto L541
	}
L541:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v3595 = m.ExcPending
	if v3595 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+180)) = v2920
	*(*int32)(unsafe.Add(mBase, uint32(v41)+176)) = v2941
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_53), v41+int32(176))
	mBase = m.M
	v3602 = m.ExcPending
	if v3602 != 0 {
		goto L1
	} else {
		goto L543
	}
L543:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(1249), int32(_a_F_pg_restore_extended_stats_40))
	mBase = m.M
	v3607 = m.ExcPending
	if v3607 != 0 {
		goto L1
	} else {
		goto L544
	}
L544:
	;
	v4126 = v3585
	v4127 = v3586
	v4158 = v2922
	goto L427
L545:
	;
	v3647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1304)))
	if v3647&int32(1) != 0 {
		goto L554
	} else {
		goto L555
	}
L546:
	;
	v3610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3465)+13)))
	switch v3610 - int32(109) {
	case 0, 5:
		goto L545
	default:
		goto L547
	}
L547:
	;
	v3613 = int32(1)
	v3614 = int32(0)
	v3617 = F_errstart(m, int32(19), v3614)
	mBase = m.M
	v3618 = m.ExcPending
	if v3618 != 0 {
		goto L1
	} else {
		goto L548
	}
L548:
	;
	if v3617 == int32(0) {
		v4126 = v3613
		v4127 = v3614
		v4158 = v2922
		goto L427
	} else {
		goto L549
	}
L549:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v3623 = m.ExcPending
	if v3623 != 0 {
		goto L1
	} else {
		goto L550
	}
L550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+164)) = v2920
	*(*int32)(unsafe.Add(mBase, uint32(v41)+160)) = v2941
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_54), v41+int32(160))
	mBase = m.M
	v3630 = m.ExcPending
	if v3630 != 0 {
		goto L1
	} else {
		goto L551
	}
L551:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+152)) = int32(_a_F_pg_restore_extended_stats_49)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+148)) = int32(_a_F_pg_restore_extended_stats_50)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+144)) = int32(_a_F_pg_restore_extended_stats_51)
	F_errhint(m, int32(_a_F_pg_restore_extended_stats_55), v41+int32(144))
	mBase = m.M
	v3641 = m.ExcPending
	if v3641 != 0 {
		goto L1
	} else {
		goto L552
	}
L552:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(1272), int32(_a_F_pg_restore_extended_stats_40))
	mBase = m.M
	v3646 = m.ExcPending
	if v3646 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	v4126 = v3613
	v4127 = v3614
	v4158 = v2922
	goto L427
L554:
	;
	v3656 = F_jbv_to_infunc_datum(m, v41+int32(880), int32(1141), v2920, int32(_a_F_pg_restore_extended_stats_56), v41+int32(872))
	mBase = m.M
	v3657 = m.ExcPending
	if v3657 != 0 {
		goto L1
	} else {
		goto L557
	}
L555:
	;
	goto L556
L556:
	;
	v3662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1305)))
	if v3662 == int32(1) {
		goto L559
	} else {
		goto L560
	}
L557:
	;
	if v3656 == int32(0) {
		goto L430
	} else {
		goto L558
	}
L558:
	;
	v3660 = *(*int64)(unsafe.Add(mBase, uint32(v41)+872))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+1384)) = v3660
	goto L556
L559:
	;
	v3669 = F_jbv_to_infunc_datum(m, v41+int32(912), int32(1142), v2920, int32(_a_F_pg_restore_extended_stats_57), v41+int32(872))
	mBase = m.M
	v3670 = m.ExcPending
	if v3670 != 0 {
		goto L1
	} else {
		goto L562
	}
L560:
	;
	goto L561
L561:
	;
	v3675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1306)))
	if v3675&int32(1) != 0 {
		goto L564
	} else {
		goto L565
	}
L562:
	;
	if v3669 == int32(0) {
		goto L430
	} else {
		goto L563
	}
L563:
	;
	v3673 = *(*int64)(unsafe.Add(mBase, uint32(v41)+872))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+1392)) = v3673
	goto L561
L564:
	;
	v3682 = F_jbv_to_infunc_datum(m, v41+int32(944), int32(1141), v2920, int32(_a_F_pg_restore_extended_stats_58), v41+int32(872))
	mBase = m.M
	v3683 = m.ExcPending
	if v3683 != 0 {
		goto L1
	} else {
		goto L567
	}
L565:
	;
	goto L566
L566:
	;
	if v2938 == int32(3614) {
		goto L569
	} else {
		goto L570
	}
L567:
	;
	if v3682 == int32(0) {
		goto L430
	} else {
		goto L568
	}
L568:
	;
	v3686 = *(*int64)(unsafe.Add(mBase, uint32(v41)+872))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+1400)) = v3686
	goto L566
L569:
	;
	v3691 = int32(100)
	goto L571
L570:
	;
	v3691 = v2934
	goto L571
L571:
	;
	if v3355 != 0 {
		goto L572
	} else {
		goto L573
	}
L572:
	;
	v3692 = *(*int32)(unsafe.Add(mBase, uint32(v41)+984))
	v3695 = F_palloc0(m, v3692+int32(1))
	mBase = m.M
	v3696 = m.ExcPending
	if v3696 != 0 {
		goto L1
	} else {
		goto L575
	}
L573:
	;
	goto L574
L574:
	;
	v3784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1309)))
	if v3784 == int32(1) {
		goto L600
	} else {
		goto L601
	}
L575:
	;
	v3697 = *(*int32)(unsafe.Add(mBase, uint32(v41)+984))
	if v3697 != 0 {
		goto L576
	} else {
		goto L577
	}
L576:
	;
	v3698 = *(*int32)(unsafe.Add(mBase, uint32(v41)+988))
	base.MemoryCopy(m, v3695, v3698, v3697)
	goto L578
L577:
	;
	goto L578
L578:
	;
	v3705 = F_array_in_safe(m, v41+int32(840), v3695, v2938, v2936, v2920, int32(_a_F_pg_restore_extended_stats_45), v41+int32(872))
	mBase = m.M
	v3706 = m.ExcPending
	if v3706 != 0 {
		goto L1
	} else {
		goto L579
	}
L579:
	;
	F_pfree(m, v3695)
	mBase = m.M
	v3708 = m.ExcPending
	if v3708 != 0 {
		goto L1
	} else {
		goto L580
	}
L580:
	;
	v3709 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1016))
	v3712 = F_palloc0(m, v3709+int32(1))
	mBase = m.M
	v3713 = m.ExcPending
	if v3713 != 0 {
		goto L1
	} else {
		goto L581
	}
L581:
	;
	v3714 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1016))
	if v3714 != 0 {
		goto L582
	} else {
		goto L583
	}
L582:
	;
	v3715 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1020))
	base.MemoryCopy(m, v3712, v3715, v3714)
	goto L584
L583:
	;
	goto L584
L584:
	;
	v3724 = F_array_in_safe(m, v41+int32(840), v3712, int32(700), int32(-1), v2920, int32(_a_F_pg_restore_extended_stats_44), v41+int32(871))
	mBase = m.M
	v3725 = m.ExcPending
	if v3725 != 0 {
		goto L1
	} else {
		goto L585
	}
L585:
	;
	F_pfree(m, v3712)
	mBase = m.M
	v3727 = m.ExcPending
	if v3727 != 0 {
		goto L1
	} else {
		goto L586
	}
L586:
	;
	v3728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+872)))
	if v3728 != int32(1) {
		goto L430
	} else {
		goto L587
	}
L587:
	;
	v3731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+871)))
	if v3731&int32(1) == int32(0) {
		goto L430
	} else {
		goto L588
	}
L588:
	;
	v3737 = F_pg_detoast_datum(m, base.I32_wrap_i64(v3705))
	mBase = m.M
	v3738 = m.ExcPending
	if v3738 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	v3740 = F_pg_detoast_datum(m, base.I32_wrap_i64(v3724))
	mBase = m.M
	v3741 = m.ExcPending
	if v3741 != 0 {
		goto L1
	} else {
		goto L590
	}
L590:
	;
	v3742 = *(*int32)(unsafe.Add(mBase, uint32(v3737)+16))
	v3743 = *(*int32)(unsafe.Add(mBase, uint32(v3740)+16))
	if v3742 != v3743 {
		goto L591
	} else {
		goto L592
	}
L591:
	;
	v3747 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v3748 = m.ExcPending
	if v3748 != 0 {
		goto L1
	} else {
		goto L594
	}
L592:
	;
	goto L593
L593:
	;
	v3775 = *(*int32)(unsafe.Add(mBase, uint32(v3465)+52))
	v3776 = int32(0)
	F_statatt_set_slot(m, v41+int32(1360), v41+int32(1616), v41+int32(1328), int32(1), v3775, v3691, v3724, v3776, v3705, v3776)
	mBase = m.M
	v3779 = m.ExcPending
	if v3779 != 0 {
		goto L1
	} else {
		goto L599
	}
L594:
	;
	if v3747 == int32(0) {
		goto L430
	} else {
		goto L595
	}
L595:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v3753 = m.ExcPending
	if v3753 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+132)) = int32(_a_F_pg_restore_extended_stats_44)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+128)) = int32(_a_F_pg_restore_extended_stats_45)
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_59), v41+int32(128))
	mBase = m.M
	v3762 = m.ExcPending
	if v3762 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(1361), int32(_a_F_pg_restore_extended_stats_40))
	mBase = m.M
	v3767 = m.ExcPending
	if v3767 != 0 {
		goto L1
	} else {
		goto L598
	}
L598:
	;
	goto L430
L599:
	;
	goto L574
L600:
	;
	v3787 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1048))
	v3790 = F_palloc0(m, v3787+int32(1))
	mBase = m.M
	v3791 = m.ExcPending
	if v3791 != 0 {
		goto L1
	} else {
		goto L603
	}
L601:
	;
	goto L602
L602:
	;
	v3823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1310)))
	if v3823&int32(1) != 0 {
		goto L611
	} else {
		goto L612
	}
L603:
	;
	v3792 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1048))
	if v3792 != 0 {
		goto L604
	} else {
		goto L605
	}
L604:
	;
	v3793 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1052))
	base.MemoryCopy(m, v3790, v3793, v3792)
	goto L606
L605:
	;
	goto L606
L606:
	;
	v3800 = F_array_in_safe(m, v41+int32(840), v3790, v2938, v2936, v2920, int32(_a_F_pg_restore_extended_stats_60), v41+int32(872))
	mBase = m.M
	v3801 = m.ExcPending
	if v3801 != 0 {
		goto L1
	} else {
		goto L607
	}
L607:
	;
	F_pfree(m, v3790)
	mBase = m.M
	v3803 = m.ExcPending
	if v3803 != 0 {
		goto L1
	} else {
		goto L608
	}
L608:
	;
	v3804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+872)))
	if v3804 != int32(1) {
		goto L430
	} else {
		goto L609
	}
L609:
	;
	v3814 = *(*int32)(unsafe.Add(mBase, uint32(v3465)+56))
	F_statatt_set_slot(m, v41+int32(1360), v41+int32(1616), v41+int32(1328), int32(2), v3814, v3691, int64(0), int32(1), v3800, int32(0))
	mBase = m.M
	v3819 = m.ExcPending
	if v3819 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	goto L602
L611:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41)+872)) = int64(0)
	v3831 = v41 + int32(872)
	v3832 = F_jbv_to_infunc_datum(m, v41+int32(1072), int32(1141), v2920, int32(_a_F_pg_restore_extended_stats_61), v3831)
	mBase = m.M
	v3833 = m.ExcPending
	if v3833 != 0 {
		goto L1
	} else {
		goto L614
	}
L612:
	;
	goto L613
L613:
	;
	if v3390 != 0 {
		goto L618
	} else {
		goto L619
	}
L614:
	;
	if v3832 == int32(0) {
		goto L430
	} else {
		goto L615
	}
L615:
	;
	v3838 = F_construct_array_builtin(m, v3831, int32(1), int32(700))
	mBase = m.M
	v3839 = m.ExcPending
	if v3839 != 0 {
		goto L1
	} else {
		goto L616
	}
L616:
	;
	v3847 = *(*int32)(unsafe.Add(mBase, uint32(v3465)+56))
	F_statatt_set_slot(m, v41+int32(1360), v41+int32(1616), v41+int32(1328), int32(3), v3847, v3691, base.I64_extend_i32_u(v3838), int32(0), int64(0), int32(1))
	mBase = m.M
	v3853 = m.ExcPending
	if v3853 != 0 {
		goto L1
	} else {
		goto L617
	}
L617:
	;
	goto L613
L618:
	;
	v3855 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1112))
	v3858 = F_palloc0(m, v3855+int32(1))
	mBase = m.M
	v3859 = m.ExcPending
	if v3859 != 0 {
		goto L1
	} else {
		goto L621
	}
L619:
	;
	goto L620
L620:
	;
	if v3572&int32(1) != 0 {
		goto L636
	} else {
		goto L637
	}
L621:
	;
	v3860 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1112))
	if v3860 != 0 {
		goto L622
	} else {
		goto L623
	}
L622:
	;
	v3861 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1116))
	base.MemoryCopy(m, v3858, v3861, v3860)
	goto L624
L623:
	;
	goto L624
L624:
	;
	v3865 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1324))
	v3869 = F_array_in_safe(m, v41+int32(840), v3858, v3865, v2936, v2920, int32(_a_F_pg_restore_extended_stats_48), v41+int32(872))
	mBase = m.M
	v3870 = m.ExcPending
	if v3870 != 0 {
		goto L1
	} else {
		goto L625
	}
L625:
	;
	F_pfree(m, v3858)
	mBase = m.M
	v3872 = m.ExcPending
	if v3872 != 0 {
		goto L1
	} else {
		goto L626
	}
L626:
	;
	v3873 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1144))
	v3876 = F_palloc0(m, v3873+int32(1))
	mBase = m.M
	v3877 = m.ExcPending
	if v3877 != 0 {
		goto L1
	} else {
		goto L627
	}
L627:
	;
	v3878 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1144))
	if v3878 != 0 {
		goto L628
	} else {
		goto L629
	}
L628:
	;
	v3879 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1148))
	base.MemoryCopy(m, v3876, v3879, v3878)
	goto L630
L629:
	;
	goto L630
L630:
	;
	v3888 = F_array_in_safe(m, v41+int32(840), v3876, int32(700), int32(-1), v2920, int32(_a_F_pg_restore_extended_stats_47), v41+int32(871))
	mBase = m.M
	v3889 = m.ExcPending
	if v3889 != 0 {
		goto L1
	} else {
		goto L631
	}
L631:
	;
	F_pfree(m, v3876)
	mBase = m.M
	v3891 = m.ExcPending
	if v3891 != 0 {
		goto L1
	} else {
		goto L632
	}
L632:
	;
	v3892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+872)))
	if v3892 != int32(1) {
		goto L430
	} else {
		goto L633
	}
L633:
	;
	v3895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+871)))
	if v3895&int32(1) == int32(0) {
		goto L430
	} else {
		goto L634
	}
L634:
	;
	v3907 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1320))
	v3908 = int32(0)
	F_statatt_set_slot(m, v41+int32(1360), v41+int32(1616), v41+int32(1328), int32(4), v3907, v3691, v3888, v3908, v3869, v3908)
	mBase = m.M
	v3911 = m.ExcPending
	if v3911 != 0 {
		goto L1
	} else {
		goto L635
	}
L635:
	;
	goto L620
L636:
	;
	v3918 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1176))
	v3921 = F_palloc0(m, v3918+int32(1))
	mBase = m.M
	v3922 = m.ExcPending
	if v3922 != 0 {
		goto L1
	} else {
		goto L639
	}
L637:
	;
	goto L638
L638:
	;
	if v3425&int32(1) != 0 {
		goto L647
	} else {
		goto L648
	}
L639:
	;
	v3923 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1176))
	if v3923 != 0 {
		goto L640
	} else {
		goto L641
	}
L640:
	;
	v3924 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1180))
	base.MemoryCopy(m, v3921, v3924, v3923)
	goto L642
L641:
	;
	goto L642
L642:
	;
	v3933 = F_array_in_safe(m, v41+int32(840), v3921, int32(700), int32(-1), v2920, int32(_a_F_pg_restore_extended_stats_62), v41+int32(872))
	mBase = m.M
	v3934 = m.ExcPending
	if v3934 != 0 {
		goto L1
	} else {
		goto L643
	}
L643:
	;
	F_pfree(m, v3921)
	mBase = m.M
	v3936 = m.ExcPending
	if v3936 != 0 {
		goto L1
	} else {
		goto L644
	}
L644:
	;
	v3937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+872)))
	if v3937 != int32(1) {
		goto L430
	} else {
		goto L645
	}
L645:
	;
	v3947 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1320))
	F_statatt_set_slot(m, v41+int32(1360), v41+int32(1616), v41+int32(1328), int32(5), v3947, v3691, v3933, int32(0), int64(0), int32(1))
	mBase = m.M
	v3952 = m.ExcPending
	if v3952 != 0 {
		goto L1
	} else {
		goto L646
	}
L646:
	;
	goto L638
L647:
	;
	v3958 = F_type_is_multirange(m, v2938)
	mBase = m.M
	v3959 = m.ExcPending
	if v3959 != 0 {
		goto L1
	} else {
		goto L650
	}
L648:
	;
	goto L649
L649:
	;
	v4046 = *(*int32)(unsafe.Add(mBase, uint32(v2792)+52))
	v4051 = F_heap_form_tuple(m, v4046, v41+int32(1360), v41+int32(1616))
	mBase = m.M
	v4052 = m.ExcPending
	if v4052 != 0 {
		goto L1
	} else {
		goto L672
	}
L650:
	;
	if v3958 != 0 {
		goto L651
	} else {
		goto L652
	}
L651:
	;
	v3960 = F_get_multirange_range(m, v2938)
	mBase = m.M
	v3961 = m.ExcPending
	if v3961 != 0 {
		goto L1
	} else {
		goto L654
	}
L652:
	;
	v3962 = v2938
	goto L653
L653:
	;
	v3963 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1272))
	v3966 = F_palloc0(m, v3963+int32(1))
	mBase = m.M
	v3967 = m.ExcPending
	if v3967 != 0 {
		goto L1
	} else {
		goto L655
	}
L654:
	;
	v3962 = v3960
	goto L653
L655:
	;
	v3968 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1272))
	if v3968 != 0 {
		goto L656
	} else {
		goto L657
	}
L656:
	;
	v3969 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1276))
	base.MemoryCopy(m, v3966, v3969, v3968)
	goto L658
L657:
	;
	goto L658
L658:
	;
	v3975 = v41 + int32(872)
	v3976 = F_array_in_safe(m, v41+int32(840), v3966, v3962, v2936, v2920, int32(_a_F_pg_restore_extended_stats_49), v3975)
	mBase = m.M
	v3977 = m.ExcPending
	if v3977 != 0 {
		goto L1
	} else {
		goto L659
	}
L659:
	;
	v3978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+872)))
	if v3978 == int32(0) {
		goto L430
	} else {
		goto L660
	}
L660:
	;
	v3988 = int32(0)
	F_statatt_set_slot(m, v41+int32(1360), v41+int32(1616), v41+int32(1328), int32(7), v3988, v3988, int64(0), int32(1), v3976, v3988)
	mBase = m.M
	v3994 = m.ExcPending
	if v3994 != 0 {
		goto L1
	} else {
		goto L661
	}
L661:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41)+872)) = int64(0)
	v3999 = F_jbv_to_infunc_datum(m, v41+int32(1232), int32(1141), v2920, int32(_a_F_pg_restore_extended_stats_50), v3975)
	mBase = m.M
	v4000 = m.ExcPending
	if v4000 != 0 {
		goto L1
	} else {
		goto L662
	}
L662:
	;
	if v3999 == int32(0) {
		goto L430
	} else {
		goto L663
	}
L663:
	;
	v4005 = F_construct_array_builtin(m, v3975, int32(1), int32(700))
	mBase = m.M
	v4006 = m.ExcPending
	if v4006 != 0 {
		goto L1
	} else {
		goto L664
	}
L664:
	;
	v4007 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1208))
	v4010 = F_palloc0(m, v4007+int32(1))
	mBase = m.M
	v4011 = m.ExcPending
	if v4011 != 0 {
		goto L1
	} else {
		goto L665
	}
L665:
	;
	v4012 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1208))
	if v4012 != 0 {
		goto L666
	} else {
		goto L667
	}
L666:
	;
	v4013 = *(*int32)(unsafe.Add(mBase, uint32(v41)+1212))
	base.MemoryCopy(m, v4010, v4013, v4012)
	goto L668
L667:
	;
	goto L668
L668:
	;
	v4022 = F_array_in_safe(m, v41+int32(840), v4010, int32(701), int32(-1), v2920, int32(_a_F_pg_restore_extended_stats_51), v41+int32(871))
	mBase = m.M
	v4023 = m.ExcPending
	if v4023 != 0 {
		goto L1
	} else {
		goto L669
	}
L669:
	;
	v4024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+871)))
	if v4024 == int32(0) {
		goto L430
	} else {
		goto L670
	}
L670:
	;
	v4035 = int32(0)
	F_statatt_set_slot(m, v41+int32(1360), v41+int32(1616), v41+int32(1328), int32(6), int32(672), v4035, base.I64_extend_i32_u(v4005), v4035, v4022, v4035)
	mBase = m.M
	v4040 = m.ExcPending
	if v4040 != 0 {
		goto L1
	} else {
		goto L671
	}
L671:
	;
	goto L649
L672:
	;
	v4053 = *(*int32)(unsafe.Add(mBase, uint32(v2792)+52))
	v4054 = F_heap_copy_tuple_as_datum(m, v4051, v4053)
	mBase = m.M
	v4055 = m.ExcPending
	if v4055 != 0 {
		goto L1
	} else {
		goto L673
	}
L673:
	;
	F_pfree(m, v4051)
	mBase = m.M
	v4057 = m.ExcPending
	if v4057 != 0 {
		goto L1
	} else {
		goto L674
	}
L674:
	;
	v4126 = int32(0)
	v4127 = int32(1)
	v4158 = v4054
	goto L427
L675:
	;
	if v4062 != 0 {
		goto L676
	} else {
		goto L677
	}
L676:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v4066 = m.ExcPending
	if v4066 != 0 {
		goto L1
	} else {
		goto L679
	}
L677:
	;
	goto L678
L678:
	;
	if v2888 == int32(0) {
		goto L396
	} else {
		goto L682
	}
L679:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+100)) = v2920
	*(*int32)(unsafe.Add(mBase, uint32(v41)+96)) = v2798
	F_errmsg(m, int32(_a_F_pg_restore_extended_stats_39), v41+int32(96))
	mBase = m.M
	v4073 = m.ExcPending
	if v4073 != 0 {
		goto L1
	} else {
		goto L680
	}
L680:
	;
	F_errfinish(m, int32(_a_F_pg_restore_extended_stats_3), int32(1660), int32(_a_F_pg_restore_extended_stats_37))
	mBase = m.M
	v4078 = m.ExcPending
	if v4078 != 0 {
		goto L1
	} else {
		goto L681
	}
L681:
	;
	goto L678
L682:
	;
	F_pfree(m, v2888)
	mBase = m.M
	v4082 = m.ExcPending
	if v4082 != 0 {
		goto L1
	} else {
		goto L683
	}
L683:
	;
	goto L396
L684:
	;
	v4168 = v2887 + int32(1)
	if v4168 != v410 {
		v2887 = v4168
		v2888 = v4165
		v2892 = v4162
		goto L425
	} else {
		goto L685
	}
L685:
	;
	goto L426
L686:
	;
	v4171 = *(*int32)(unsafe.Add(mBase, _c_F_pg_restore_extended_stats[10]))
	v4172 = F_makeArrayResult(m, v4165, v4171)
	mBase = m.M
	v4173 = m.ExcPending
	if v4173 != 0 {
		goto L1
	} else {
		goto L689
	}
L687:
	;
	v4175 = int64(0)
	goto L688
L688:
	;
	F_relation_close(m, v2792, int32(3))
	mBase = m.M
	v4178 = m.ExcPending
	if v4178 != 0 {
		goto L1
	} else {
		goto L690
	}
L689:
	;
	v4175 = v4172
	goto L688
L690:
	;
	if v4162 != v410 {
		v4305 = v41
		v4319 = v838
		v4323 = v842
		v4324 = v843
		v4330 = v194
		v4332 = v192
		v4333 = v54
		v4342 = int32(0)
		goto L395
	} else {
		goto L691
	}
L691:
	;
	v4215 = v4175
	goto L420
L692:
	;
	v4305 = v41
	v4319 = v838
	v4323 = v842
	v4324 = v843
	v4330 = v194
	v4332 = v192
	v4333 = v54
	v4342 = int32(0)
	goto L395
L693:
	;
	v4348 = F_SearchSysCache2(m, int32(62), v863, v871)
	mBase = m.M
	v4349 = m.ExcPending
	if v4349 != 0 {
		goto L1
	} else {
		goto L694
	}
L694:
	;
	v4350 = *(*int32)(unsafe.Add(mBase, uint32(v4345)+52))
	if v4348 != 0 {
		goto L696
	} else {
		goto L697
	}
L695:
	;
	F_pfree(m, v4373)
	mBase = m.M
	v4375 = m.ExcPending
	if v4375 != 0 {
		goto L1
	} else {
		goto L704
	}
L696:
	;
	v4357 = F_heap_modify_tuple(m, v4348, v4350, v4305+int32(784), v4305+int32(776), v4305+int32(768))
	mBase = m.M
	v4358 = m.ExcPending
	if v4358 != 0 {
		goto L1
	} else {
		goto L699
	}
L697:
	;
	goto L698
L698:
	;
	v4369 = F_heap_form_tuple(m, v4350, v4305+int32(784), v4305+int32(776))
	mBase = m.M
	v4370 = m.ExcPending
	if v4370 != 0 {
		goto L1
	} else {
		goto L702
	}
L699:
	;
	F_CatalogTupleUpdate(m, v4345, v4357+int32(4), v4357)
	mBase = m.M
	v4362 = m.ExcPending
	if v4362 != 0 {
		goto L1
	} else {
		goto L700
	}
L700:
	;
	F_ReleaseCatCache(m, v4348)
	mBase = m.M
	v4364 = m.ExcPending
	if v4364 != 0 {
		goto L1
	} else {
		goto L701
	}
L701:
	;
	v4373 = v4357
	goto L695
L702:
	;
	F_CatalogTupleInsert(m, v4345, v4369)
	mBase = m.M
	v4372 = m.ExcPending
	if v4372 != 0 {
		goto L1
	} else {
		goto L703
	}
L703:
	;
	v4373 = v4369
	goto L695
L704:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v4377 = m.ExcPending
	if v4377 != 0 {
		goto L1
	} else {
		goto L705
	}
L705:
	;
	F_relation_close(m, v4345, int32(3))
	mBase = m.M
	v4380 = m.ExcPending
	if v4380 != 0 {
		goto L1
	} else {
		goto L706
	}
L706:
	;
	v4382 = v4305
	v4384 = v4342
	v4396 = v4319
	v4400 = v4323
	v4401 = v4324
	v4407 = v4330
	v4409 = v4332
	v4410 = v4333
	goto L48
L707:
	;
	v4422 = v4382
	v4424 = v4384
	v4436 = v4396
	v4440 = v4400
	v4441 = v4401
	v4449 = v4409
	v4450 = v4410
	goto L37
L708:
	;
	F_relation_close(m, v4449, int32(3))
	mBase = m.M
	v4461 = m.ExcPending
	if v4461 != 0 {
		goto L1
	} else {
		goto L711
	}
L709:
	;
	goto L710
L710:
	;
	if v4440 != 0 {
		goto L712
	} else {
		goto L713
	}
L711:
	;
	goto L710
L712:
	;
	F_pfree(m, v4440)
	mBase = m.M
	v4463 = m.ExcPending
	if v4463 != 0 {
		goto L1
	} else {
		goto L715
	}
L713:
	;
	goto L714
L714:
	;
	if v4441 != 0 {
		goto L716
	} else {
		goto L717
	}
L715:
	;
	goto L714
L716:
	;
	F_pfree(m, v4441)
	mBase = m.M
	v4465 = m.ExcPending
	if v4465 != 0 {
		goto L1
	} else {
		goto L719
	}
L717:
	;
	goto L718
L718:
	;
	if v4436 == int32(0) {
		v4471 = v4422
		v4473 = v4424
		v4499 = v4450
		goto L3
	} else {
		goto L720
	}
L719:
	;
	goto L718
L720:
	;
	F_pfree(m, v4436)
	mBase = m.M
	v4469 = m.ExcPending
	if v4469 != 0 {
		goto L1
	} else {
		goto L721
	}
L721:
	;
	v4471 = v4422
	v4473 = v4424
	v4499 = v4450
	goto L3
}
func F_pg_set_regex_collation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	if l0 != 0 {
		v2 = F_pg_newlocale_from_collation(m, l0)
		mBase = m.M
		v3 = m.ExcPending
		if v3 != 0 {
			return
		} else {
			v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			if v4 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_pg_set_regex_collation_0), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_pg_set_regex_collation_1), int32(56), int32(_a_F_pg_set_regex_collation_2))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
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
				*(*int32)(unsafe.Add(mBase, _c_F_pg_set_regex_collation[0])) = v2
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_errcode(m, int32(34209924))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_pg_set_regex_collation_3), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_errhint(m, int32(_a_F_pg_set_regex_collation_4), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_pg_set_regex_collation_1), int32(48), int32(_a_F_pg_set_regex_collation_2))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
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
func F_pg_settings_get_flags(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int64
	_ = v118
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_text_to_cstring(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(48)
	return v118
L2:
	;
	return int64(0)
L3:
	;
	v19 = F_find_option(m, v12, int32(0), int32(1), int32(21))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v19 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
	v118 = int64(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v26&int32(32) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v32 = F_cstring_to_text(m, int32(_a_F_pg_settings_get_flags_0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	v38 = v26
	v39 = int32(0)
	v40 = v9
	goto L10
L10:
	;
	if v38&int32(8) != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = base.I64_extend_i32_u(v32)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v38 = v37
	v39 = int32(1)
	v40 = v9 | int32(8)
	goto L10
L12:
	;
	v44 = F_cstring_to_text(m, int32(_a_F_pg_settings_get_flags_1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L15
	}
L13:
	;
	v51 = v38
	v52 = v39
	goto L14
L14:
	;
	if v51&int32(16) != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = base.I64_extend_i32_u(v44)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v51 = v50
	v52 = v39 + int32(1)
	goto L14
L16:
	;
	v59 = F_cstring_to_text(m, int32(_a_F_pg_settings_get_flags_2))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L2
	} else {
		goto L19
	}
L17:
	;
	v66 = v51
	v67 = v52
	goto L18
L18:
	;
	if v66&int32(4) != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+v52<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v59)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v66 = v65
	v67 = v52 + int32(1)
	goto L18
L20:
	;
	v74 = F_cstring_to_text(m, int32(_a_F_pg_settings_get_flags_3))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	v81 = v66
	v82 = v67
	goto L22
L22:
	;
	if v81&int32(128) != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+v67<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v74)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v81 = v80
	v82 = v67 + int32(1)
	goto L22
L24:
	;
	v89 = F_cstring_to_text(m, int32(_a_F_pg_settings_get_flags_4))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L27
	}
L25:
	;
	v96 = v82
	v97 = v81
	goto L26
L26:
	;
	if v97&int32(_a_F_pg_settings_get_flags_5) != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+v82<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v89)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v96 = v82 + int32(1)
	v97 = v95
	goto L26
L28:
	;
	v104 = F_cstring_to_text(m, int32(_a_F_pg_settings_get_flags_6))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L31
	}
L29:
	;
	v110 = v96
	goto L30
L30:
	;
	v112 = F_construct_array_builtin(m, v9, v110, int32(25))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L32
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+v96<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v104)
	v110 = v96 + int32(1)
	goto L30
L32:
	;
	v118 = base.I64_extend_i32_u(v112)
	goto L1
}
func F_pg_sjis_dsplen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32((v2+int32(95))&int32(255)) < base.Ui32(int32(63)) {
		return int32(1)
	} else {
		if base.I32_extend8_s(v2) < int32(0) {
			return int32(2)
		} else {
			v16 = int32(-1)
			if v2 == int32(127) {
				v21 = v16
			} else {
				v21 = int32(1)
			}
			if base.Ui32(v2) < base.Ui32(int32(32)) {
				v24 = v16
			} else {
				v24 = v21
			}
			if v2 != 0 {
				v26 = v24
			} else {
				v26 = int32(0)
			}
			return v26
		}
	}
}
func F_pg_sjis_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v45 int32
	_ = v45
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	if l1 <= int32(0) {
		v69 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v69 - l0
L2:
	;
	v10 = l1
	v12 = l0
	goto L3
L3:
	;
	v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12))))
	if int32(0) <= v15 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v69 = v62
	goto L1
L5:
	;
	v62 = v12 + v61
	v63 = v10 - v61
	if int32(0) < v63 {
		v10 = v63
		v12 = v62
		goto L3
	} else {
		goto L20
	}
L6:
	;
	v61 = int32(1)
	goto L5
L7:
	;
	if v15 != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v25 = base.B2i32(base.Ui32((v15+int32(95))&int32(255)) < base.Ui32(int32(63)))
	if base.Ui32((v15+int32(95))&int32(255)) < base.Ui32(int32(63)) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v69 = v12
	goto L1
L11:
	;
	v26 = int32(1)
	goto L13
L12:
	;
	v26 = int32(2)
	goto L13
L13:
	;
	if base.B2i32(base.Ui32(v10) < base.Ui32(v26))|v25 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v33 = int32(255)
	if base.B2i32(base.Ui32(int32(28)) < base.Ui32((v15+int32(32))&v33))&base.B2i32(base.Ui32(int32(31)) <= base.Ui32((v15+int32(127))&v33)) != 0 {
		v69 = v12
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if base.Ui32(v10) < base.Ui32(v26) {
		v69 = v12
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+1)))
	if base.B2i32(v45 < int32(-3))|base.B2i32(base.Ui32((v45+int32(-64))&int32(255)) < base.Ui32(int32(63))) != 0 {
		v61 = int32(2)
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v69 = v12
	goto L1
L19:
	;
	goto L6
L20:
	;
	goto L4
}
func F_pg_sprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l2
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v8)+21)) = v11
	F_dopr(m, v8+int32(8), l1, l2)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		v24 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v24)
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
		m.G0 = v8 + int32(32)
		if v28 != 0 {
			v35 = int32(-1)
		} else {
			v35 = v27 + (v23 - v26)
		}
		return v35
	}
}
func F_pg_strfold(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	v6 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v11 == v6 {
		v14 = int32(0)
		if base.B2i32(l1 == v14)|base.B2i32(l3 == v14) == v14 {
			v21 = int32(1)
			v22 = l3 - v21
			v24 = l1 - v21
			if base.Ui32(v22) < base.Ui32(v24) {
				v26 = v22
			} else {
				v26 = v24
			}
			if v26 == int32(0) {
				v86 = int32(0)
				v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v86))))
				if base.Ui32((v94-int32(65))&int32(255)) < base.Ui32(int32(26)) {
					v103 = v94 | int32(32)
				} else {
					v103 = v94
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v86))) = uint8(v103)
				v111 = v86 + int32(1)
			} else {
				v30 = int32(1)
				v31 = v26 + v30
				v41 = int32(0)
				v46 = v6
				for {
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v41))))
					if base.Ui32((v49-int32(65))&int32(255)) < base.Ui32(int32(26)) {
						v58 = v49 | int32(32)
					} else {
						v58 = v49
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l0+v41))) = uint8(v58)
					v61 = v41 | int32(1)
					v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v61))))
					if base.Ui32((v64-int32(65))&int32(255)) < base.Ui32(int32(26)) {
						v73 = v64 | int32(32)
					} else {
						v73 = v64
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l0+v61))) = uint8(v73)
					v75 = int32(2)
					v76 = v41 + v75
					v78 = v46 + v75
					if v78 != v31&int32(-2) {
						v41 = v76
						v46 = v78
						continue
					} else {
						break
					}
					break
				}
				if v31&v30 == int32(0) {
					v111 = v76
				} else {
					v86 = v76
					v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v86))))
					if base.Ui32((v94-int32(65))&int32(255)) < base.Ui32(int32(26)) {
						v103 = v94 | int32(32)
					} else {
						v103 = v94
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l0+v86))) = uint8(v103)
					v111 = v86 + int32(1)
				}
			}
			if base.Ui32(l1) <= base.Ui32(v111) {
				v145 = l3
				return v145
			} else {
				v127 = v26 + int32(1)
				v134 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v127))) = uint8(v134)
				return l3
			}
		} else {
			v120 = int32(0)
			if l1 == v120 {
				v145 = l3
				return v145
			} else {
				v127 = v120
				v134 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v127))) = uint8(v134)
				return l3
			}
		}
	} else {
		v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		v138 = m.T0[v137].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3, l4)
		mBase = m.M
		v141 = m.ExcPending
		if v141 != 0 {
			return int32(0)
		} else {
			v145 = v138
			return v145
		}
	}
}
func F_pg_strnxfrm(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v7 == int32(0) {
		if base.Ui32(l1) <= base.Ui32(l3) {
			v21 = l3
			return v21
		} else {
			if l3 != 0 {
				base.MemoryCopy(m, l0, l2, l3)
			} else {
			}
			v13 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0+l3))) = uint8(v13)
			return l3
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		v17 = m.T0[v16].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3, l4)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = v17
			return v21
		}
	}
}
func F_pg_strtoint16_safe(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v17 = base.B2i32(v15 == int32(45))
	v18 = l0 + v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v23 = (v19 - int32(48)) & int32(255)
	if base.Ui32(int32(10)) <= base.Ui32(v23) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return base.I32_extend16_s(v406)
L2:
	;
	F_errsave_finish(m, l1, int32(_a_F_pg_strtoint16_safe_0), v399, int32(_a_F_pg_strtoint16_safe_1))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L81
	} else {
		goto L90
	}
L3:
	;
	v373 = int32(0)
	v374 = F_errsave_start(m, l1)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L81
	} else {
		goto L86
	}
L4:
	;
	v347 = int32(0)
	v348 = F_errsave_start(m, l1)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L81
	} else {
		goto L82
	}
L5:
	;
	v97 = l0
	v100 = v15
	goto L22
L6:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v28 = v26 - int32(48)
	if base.Ui32(v28&int32(255)) <= base.Ui32(int32(9)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v38 = v23
	v39 = v18 + int32(1)
	v40 = v28
	goto L10
L8:
	;
	v64 = v26
	v65 = v23
	goto L9
L9:
	;
	if v64&int32(255) != 0 {
		goto L5
	} else {
		goto L14
	}
L10:
	;
	if base.Ui32(int32(3276)) < base.Ui32(v38&int32(_a_F_pg_strtoint16_safe_2)) {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v64 = v53
	v65 = v52
	goto L9
L12:
	;
	v48 = int32(10)
	v50 = int32(255)
	v52 = v38*v48 + v40&v50
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	v57 = v53 - int32(48)
	if base.Ui32(v57&v50) < base.Ui32(v48) {
		v38 = v52
		v39 = v39 + int32(1)
		v40 = v57
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	if v15 == int32(45) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if base.Ui32(int32(_a_F_pg_strtoint16_safe_3)) < base.Ui32(v65&int32(_a_F_pg_strtoint16_safe_2)) {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if base.I32_extend16_s(v65) < int32(0) {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	v406 = int32(0) - v65
	goto L1
L19:
	;
	v406 = v65
	goto L1
L20:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v118 != int32(48) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v116 = v97 + int32(1)
	v117 = v17
	goto L20
L22:
	;
	if base.Ui32(v100-int32(9)) < base.Ui32(int32(5)) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v111 = int32(1)
	v116 = v97 + v111
	v117 = v111
	goto L20
L24:
	;
	goto L23
L25:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	v97 = v97 + int32(1)
	v100 = v108
	goto L22
L26:
	;
	switch v100 - int32(32) {
	case 0:
		goto L25
	default:
		v116 = v97
		v117 = v17
		goto L20
	case 11:
		goto L21
	case 13:
		goto L24
	}
L27:
	;
	if v300 == v302 {
		goto L3
	} else {
		goto L69
	}
L28:
	;
	v261 = v116
	v262 = int32(0)
	v264 = v118
	goto L60
L29:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	switch v121 - int32(66) {
	case 0, 32:
		goto L30
	default:
		goto L28
	case 13, 45:
		goto L31
	case 22, 54:
		goto L32
	}
L30:
	;
	v219 = v116 + int32(2)
	v222 = v219
	v223 = int32(0)
	goto L52
L31:
	;
	v178 = v116 + int32(2)
	v181 = v178
	v182 = int32(0)
	goto L44
L32:
	;
	v126 = v116 + int32(2)
	v129 = v126
	v130 = int32(0)
	goto L33
L33:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	goto L35
L34:
	;
	goto L3
L35:
	;
	if base.B2i32(base.Ui32(v136-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v136|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if base.Ui32(int32(2048)) < base.Ui32(v130&int32(_a_F_pg_strtoint16_safe_2)) {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if v136 != int32(95) {
		v300 = v129
		v301 = v130
		v302 = v126
		v303 = v136
		goto L27
	} else {
		goto L40
	}
L39:
	;
	v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_pg_strtoint16_safe[0]))))
	v129 = v129 + int32(1)
	v130 = v154 + v130<<(uint(int32(4))%32)
	goto L33
L40:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
	if v160 == int32(0) {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	goto L42
L42:
	;
	if base.B2i32(base.Ui32(v160-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v160|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		v129 = v129 + int32(1)
		goto L33
	} else {
		goto L43
	}
L43:
	;
	goto L34
L44:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	if v188&int32(248) == int32(48) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L3
L46:
	;
	if base.Ui32(int32(_a_F_pg_strtoint16_safe_4)) < base.Ui32(v182&int32(_a_F_pg_strtoint16_safe_2)) {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v188 != int32(95) {
		v300 = v181
		v301 = v182
		v302 = v178
		v303 = v188
		goto L27
	} else {
		goto L50
	}
L49:
	;
	v181 = v181 + int32(1)
	v182 = (v188-int32(48))&int32(255) | v182<<(uint(int32(3))%32)
	goto L44
L50:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+1)))
	if base.Ui32(int32(248)) <= base.Ui32((v208-int32(56))&int32(255)) {
		v181 = v181 + int32(1)
		goto L44
	} else {
		goto L51
	}
L51:
	;
	goto L45
L52:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v229&int32(254) == int32(48) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L3
L54:
	;
	if base.Ui32(int32(_a_F_pg_strtoint16_safe_5)) < base.Ui32(v223&int32(_a_F_pg_strtoint16_safe_2)) {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v229 != int32(95) {
		v300 = v222
		v301 = v223
		v302 = v219
		v303 = v229
		goto L27
	} else {
		goto L58
	}
L57:
	;
	v242 = int32(1)
	v222 = v222 + v242
	v223 = (v229-int32(48))&int32(255) | v223<<(uint(v242)%32)
	goto L52
L58:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+1)))
	if base.Ui32(int32(254)) <= base.Ui32((v249-int32(50))&int32(255)) {
		v222 = v222 + int32(1)
		goto L52
	} else {
		goto L59
	}
L59:
	;
	goto L53
L60:
	;
	v271 = (v264 - int32(48)) & int32(255)
	if base.Ui32(v271) <= base.Ui32(int32(9)) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L3
L62:
	;
	if base.Ui32(int32(3276)) < base.Ui32(v262&int32(_a_F_pg_strtoint16_safe_2)) {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if v264&int32(255) != int32(95) {
		v300 = v261
		v301 = v262
		v302 = v116
		v303 = v264
		goto L27
	} else {
		goto L66
	}
L65:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)))
	v261 = v261 + int32(1)
	v262 = v262*int32(10) + v271
	v264 = v281
	goto L60
L66:
	;
	if v261 == v116 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)))
	if base.Ui32((v289-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v261 = v261 + int32(1)
		v264 = v289
		goto L60
	} else {
		goto L68
	}
L68:
	;
	goto L61
L69:
	;
	v310 = v300
	v313 = v303
	goto L70
L70:
	;
	v318 = v313 & int32(255)
	if base.B2i32(base.Ui32(v318-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v318 == int32(32)) != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L4
L72:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
	v310 = v310 + int32(1)
	v313 = v326
	goto L70
L73:
	;
	if v318 != 0 {
		goto L3
	} else {
		goto L75
	}
L74:
	;
	goto L71
L75:
	;
	if v117 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if base.Ui32(int32(_a_F_pg_strtoint16_safe_3)) < base.Ui32(v301&int32(_a_F_pg_strtoint16_safe_2)) {
		goto L4
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if int32(0) <= base.I32_extend16_s(v301) {
		v406 = v301
		goto L1
	} else {
		goto L80
	}
L79:
	;
	v406 = int32(0) - v301
	goto L1
L80:
	;
	goto L74
L81:
	;
	return int32(0)
L82:
	;
	if v348 == int32(0) {
		v406 = v347
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(_a_F_pg_strtoint16_safe_6)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg(m, int32(_a_F_pg_strtoint16_safe_7), v12)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v393 = v347
	v399 = int32(350)
	goto L2
L86:
	;
	if v374 == int32(0) {
		v406 = v373
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L81
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_pg_strtoint16_safe_6)
	F_errmsg(m, int32(_a_F_pg_strtoint16_safe_8), v12+int32(16))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L81
	} else {
		goto L89
	}
L89:
	;
	v393 = v373
	v399 = int32(356)
	goto L2
L90:
	;
	v406 = v393
	goto L1
}
func F_pg_tablespace_size_name(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_get_tablespace_oid(m, v3, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = F_calculate_tablespace_size(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			if v9 < int64(0) {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				v16 = int64(0)
			} else {
				v16 = v9
			}
			return v16
		}
	}
}
func F_pg_terminate_backend(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int64
	_ = v103
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = base.I32_wrap_i64(v13)
	if int32(0) <= v14 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = F_pg_signal_backend(m, v17, int32(15))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L8
	} else {
		goto L51
	}
L4:
	;
	v96 = int32(0)
	if base.B2i32(v14 == v96)|v19 != 0 {
		v183 = base.B2i32(v19 == v96)
		goto L25
	} else {
		goto L26
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L8
	} else {
		goto L20
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L15
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return int64(0)
L9:
	;
	switch v19 - int32(2) {
	case 0:
		goto L5
	case 1:
		goto L7
	case 2:
		goto L6
	default:
		goto L4
	}
L10:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	F_errmsg(m, int32(_a_F_pg_terminate_backend_0), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v36 = int32(_a_F_pg_terminate_backend_1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v36
	v43 = F_errdetail(m, int32(_a_F_pg_terminate_backend_2), v11+int32(32))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_pg_terminate_backend_3), int32(256), int32(_a_F_pg_terminate_backend_4))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L15:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	F_errmsg(m, int32(_a_F_pg_terminate_backend_0), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(_a_F_pg_terminate_backend_5)
	v66 = F_errdetail(m, int32(_a_F_pg_terminate_backend_6), v11+int32(48))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_pg_terminate_backend_3), int32(263), int32(_a_F_pg_terminate_backend_4))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(_a_F_pg_terminate_backend_0), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = int32(_a_F_pg_terminate_backend_7)
	v89 = F_errdetail(m, int32(_a_F_pg_terminate_backend_8), v11-int32(-64))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_pg_terminate_backend_3), int32(270), int32(_a_F_pg_terminate_backend_4))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	m.G0 = v11 + int32(80)
	return base.I64_extend_i32_u(v183)
L26:
	;
	v103 = v13 & int64(2147483647)
	v109 = v103
	v110 = int64(100)
	goto L27
L27:
	;
	v113 = F_pgmem_kill(m, v17, int32(0))
	mBase = m.M
	if v113 == int32(-1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v161 = int32(0)
	v164 = F_errstart(m, int32(19), v161)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L8
	} else {
		goto L47
	}
L29:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_pg_terminate_backend[0]))
	if v118 == int32(71) {
		v183 = int32(1)
		goto L25
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v109 < v110 {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v17
	F_errmsg(m, int32(_a_F_pg_terminate_backend_9), v11)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_pg_terminate_backend_3), int32(199), int32(_a_F_pg_terminate_backend_10))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v138 = v109
	goto L39
L38:
	;
	v138 = v110
	goto L39
L39:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_pg_terminate_backend[1]))
	if v140 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L8
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_pg_terminate_backend[2]))
	v148 = F_WaitLatch(m, v144, int32(41), base.I32_wrap_i64(v138), int32(134217731))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L8
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_pg_terminate_backend[2]))
	v152 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v152
	v157 = base.AtomicRmwOr32(m, v152, int32(_a_F_pg_terminate_backend_11), v152)
	goto L45
L45:
	;
	v158 = v109 - v138
	if int64(0) < v158 {
		v109 = v158
		v110 = v138
		goto L27
	} else {
		goto L46
	}
L46:
	;
	goto L28
L47:
	;
	if v164 == int32(0) {
		v183 = v161
		goto L25
	} else {
		goto L48
	}
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v17
	F_errmsg_plural(m, int32(_a_F_pg_terminate_backend_12), int32(_a_F_pg_terminate_backend_13), v14, v11+int32(16))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_pg_terminate_backend_3), int32(219), int32(_a_F_pg_terminate_backend_10))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	v183 = v161
	goto L25
L51:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	F_errmsg(m, int32(_a_F_pg_terminate_backend_14), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_pg_terminate_backend_3), int32(247), int32(_a_F_pg_terminate_backend_4))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_tolower(m *base.Module, l0 int32) int32 {
	var v10 int32
	_ = v10
	if base.Ui32((l0-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		v10 = l0 | int32(32)
	} else {
		v10 = l0
	}
	return v10
}
func F_pg_truncate_visibility_map(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int64
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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int64
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int64
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int64
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	v5 = m.G0
	v7 = v5 - int32(96)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_relation_open(m, v9, int32(8))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+119)))
	v18 = v16 - int32(109)
	v25 = int32(0)
	if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v18))|base.B2i32(int32(1)<<(uint(v18)%32)&int32(161) == v25) == v25 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L51
	}
L6:
	;
	v56 = v30
	goto L8
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v32
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+48)) = v34
	v38 = F_smgropen(m, v7+int32(48), v31)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+28)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+92)) = int32(2)
	v62 = F_visibilitymap_prepare_truncate(m, v11, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v38
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v56 = v54
	goto L8
L11:
	;
	v50 = v42
	goto L13
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+76))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	v50 = v48
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v50 + int32(1)
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+88)) = v62
	if v62 != int32(-1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v67 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v99 = int32(0)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+84)) = v99
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_pg_truncate_visibility_map[0]))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v102)+336)) = v103 | int32(3)
	v107 = int32(_a_F_pg_truncate_visibility_map_0)
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_pg_truncate_visibility_map[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_truncate_visibility_map[1])) = v109 + int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+118)))
	if v114 != int32(112) {
		goto L27
	} else {
		goto L28
	}
L18:
	;
	v93 = v67
	goto L20
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v69
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v71
	v75 = F_smgropen(m, v7+int32(32), v68)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v95 = F_smgrnblocks(m, v93, int32(2))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L26
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v75
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75)+72))
	if v79 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v93 = v91
	goto L20
L23:
	;
	v87 = v79
	goto L25
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)+76))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)+72))
	v87 = v85
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+72)) = v87 + int32(1)
	goto L22
L26:
	;
	v99 = v95
	goto L17
L27:
	;
	if v62 != int32(-1) {
		goto L38
	} else {
		goto L39
	}
L28:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_pg_truncate_visibility_map[2]))
	if v118 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v121 != 0 {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = int32(0)
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+68)) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+76)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v7)+80)) = int32(2)
	F_XLogBeginInsert(m)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v122 != 0 {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	F_XLogRegisterData(m, v7-int32(-64), int32(20))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v140 = F_XLogInsert(m, int32(2), int32(33))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_XLogFlush(m, v140)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	goto L27
L38:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v146 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v183 = int32(_a_F_pg_truncate_visibility_map_0)
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_pg_truncate_visibility_map[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_truncate_visibility_map[1])) = v185 - int32(1)
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_pg_truncate_visibility_map[0]))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+336)) = v191 & int32(-4)
	F_relation_close(m, v11, int32(8))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L50
	}
L41:
	;
	v172 = v146
	goto L43
L42:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v148
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v150
	v154 = F_smgropen(m, v7+int32(16), v147)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	F_smgrtruncate(m, v172, v7+int32(92), int32(1), v7+int32(84), v7+int32(88))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L49
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v154
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154)+72))
	if v158 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v172 = v170
	goto L43
L46:
	;
	v166 = v158
	goto L48
L47:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v154)+76))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v154)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v154)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v154)+72))
	v166 = v164
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+72)) = v166 + int32(1)
	goto L45
L49:
	;
	goto L40
L50:
	;
	m.G0 = v7 + int32(96)
	return int64(0)
L51:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v210 + int32(4)
	F_errmsg(m, int32(_a_F_pg_truncate_visibility_map_1), v7)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v218 = int32(*(*int8)(unsafe.Add(mBase, uint32(v217)+119)))
	F_errdetail_relkind_not_supported(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_pg_truncate_visibility_map_2), int32(932), int32(_a_F_pg_truncate_visibility_map_3))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_type_is_visible(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int64
	_ = v25
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_TypeIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v25 = int64(0)
		} else {
			v25 = base.I64_extend_i32_u(v14)
		}
		m.G0 = v7 + int32(16)
		return v25
	}
}
func F_pg_utf_mblen_private(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v27 int32
	_ = v27
	v2 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if int32(0) <= v2 {
		return int32(1)
	} else {
		v8 = v2 & int32(255)
		if v8&int32(224) == int32(192) {
			return int32(2)
		} else {
			if v8&int32(240) == int32(224) {
				return int32(3)
			} else {
				if v8&int32(248) == int32(240) {
					v27 = int32(4)
				} else {
					v27 = int32(1)
				}
				return v27
			}
		}
	}
}
func F_pg_valid_server_encoding_private(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	v2 = int32(-1)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	if l0 == int32(0) {
		v91 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v91 == int32(7) {
		goto L27
	} else {
		goto L28
	}
L2:
	;
	m.G0 = v14 - int32(-64)
	goto L1
L3:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v19 == int32(0) {
		v91 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = F_strlen(m, l0)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v22) {
		v91 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v25 = l0
	v26 = v19
	v27 = v14
	goto L6
L6:
	;
	v35 = F_isalnum(m, v26&int32(255))
	mBase = m.M
	if v35 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v52 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v52)
	v56 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14))))
	v57 = int32(_a_F_pg_valid_server_encoding_private_0)
	v58 = int32(_a_F_pg_valid_server_encoding_private_1)
	goto L15
L8:
	;
	if base.Ui32((v26-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v48 = v27
	goto L10
L10:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v49 != 0 {
		v25 = v25 + int32(1)
		v26 = v49
		v27 = v48
		goto L6
	} else {
		goto L14
	}
L11:
	;
	v44 = v26 | int32(32)
	goto L13
L12:
	;
	v44 = v26
	goto L13
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v44)
	v48 = v27 + int32(1)
	goto L10
L14:
	;
	goto L7
L15:
	;
	v70 = v58 + (v57-v58)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v72 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71))))
	v73 = v56 - v72
	if v73 != 0 {
		v76 = v73
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v91 = v2
	goto L2
L17:
	;
	v80 = base.B2i32(v76 < int32(0))
	if v76 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v74 = F_strcmp(m, v14, v71)
	mBase = m.M
	if v74 != 0 {
		v76 = v74
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v91 = v75
	goto L2
L20:
	;
	v81 = v70 - int32(8)
	goto L22
L21:
	;
	v81 = v57
	goto L22
L22:
	;
	if v76 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v84 = v58
	goto L25
L24:
	;
	v84 = v70 + int32(8)
	goto L25
L25:
	;
	if base.Ui32(v84) <= base.Ui32(v81) {
		v57 = v81
		v58 = v84
		goto L15
	} else {
		goto L26
	}
L26:
	;
	goto L16
L27:
	;
	v99 = v2
	goto L29
L28:
	;
	v99 = v91
	goto L29
L29:
	;
	if int32(34) < v91 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v102 = v2
	goto L32
L31:
	;
	v102 = v99
	goto L32
L32:
	;
	if v91 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v105 = v2
	goto L35
L34:
	;
	v105 = v102
	goto L35
L35:
	;
	return v105
}
func F_pg_visibility(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int64
	_ = v201
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v2
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)) = uint16(v2)
	v21 = F_relation_open(m, v13, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L2
	} else {
		goto L57
	}
L2:
	;
	return int64(0)
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+119)))
	v28 = v26 - int32(109)
	v35 = int32(0)
	if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v28))|base.B2i32(int32(1)<<(uint(v28)%32)&int32(161) == v35) == v35 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if base.Ui64(int64(4294967295)) <= base.Ui64(v12) {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L2
	} else {
		goto L52
	}
L7:
	;
	v43 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v43, int32(1), int32(_a_F_pg_visibility_0), int32(16), int32(-1), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v43, int32(2), int32(_a_F_pg_visibility_1), int32(16), int32(-1), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	F_TupleDescInitEntry(m, v43, int32(3), int32(_a_F_pg_visibility_2), int32(16), int32(-1), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v66 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v66 < v75 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v153 = F_BlessTupleDesc(m, v43)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L2
	} else {
		goto L31
	}
L13:
	;
	v79 = v43 + int32(28)
	v86 = v66
	v87 = v75
	v89 = v66
	goto L17
L14:
	;
	v143 = v66
	v150 = v75
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v143
	goto L12
L16:
	;
	v143 = v137
	v150 = v116
	goto L15
L17:
	;
	v95 = v79 + v75<<(uint(int32(3))%32) + v86*int32(100)
	v98 = v79 + v86<<(uint(int32(3))%32)
	if v75 != v87 {
		v116 = v87
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v137 = v75
	goto L16
L19:
	;
	v117 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+2)))
	if v117 <= int32(0) {
		v137 = v86
		goto L16
	} else {
		goto L27
	}
L20:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+7)))
	if v100 != int32(118) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v116 = v86
	goto L19
L22:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+4)))
	if v103 != int32(1) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+6)))
	if v106&int32(6) != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+2)))
	if v109 <= int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+90)))
	if v112 != int32(118) {
		v116 = v75
		goto L19
	} else {
		goto L26
	}
L26:
	;
	goto L21
L27:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+90)))
	if v120 == int32(118) {
		v137 = v86
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+5)))
	v129 = (v89 + v123 - int32(1)) & (int32(0) - v123)
	if int32(_a_F_pg_visibility_3) < v129 {
		v137 = v86
		goto L16
	} else {
		goto L29
	}
L29:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v98))) = uint16(v129)
	v135 = v86 + int32(1)
	if v135 != v75 {
		v86 = v135
		v87 = v116
		v89 = v129 + v117
		goto L17
	} else {
		goto L30
	}
L30:
	;
	goto L18
L31:
	;
	v155 = base.I32_wrap_i64(v12)
	v158 = F_visibilitymap_get_status(m, v21, v155, v10+int32(44))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	if v160 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_ReleaseBuffer(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v163 = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = base.I64_extend_i32_u(v158 & v163)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = base.I64_extend_i32_u(int32(base.Ui32(v158)>>(uint(v163)%32)) & v163)
	v174 = F_RelationGetNumberOfBlocksInFork(m, v21, int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	F_relation_close(m, v21, int32(1))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L2
	} else {
		goto L49
	}
L38:
	;
	if base.Ui64(v12) < base.Ui64(base.I64_extend_i32_u(v174)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v178 = F_ReadBuffer(m, v21, v155)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L2
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = int64(0)
	goto L37
L42:
	;
	F_LockBufferInternal(m, v178, int32(1))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	if v178 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v201 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v200)+10)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = int64(base.Ui64(v201)>>(uint(int64(2))%64)) & int64(1)
	F_UnlockReleaseBuffer(m, v178)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L2
	} else {
		goto L48
	}
L45:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_pg_visibility[0]))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v186+(v178^int32(-1))<<(uint(int32(2))%32))))
	v200 = v192
	goto L44
L46:
	;
	goto L47
L47:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_pg_visibility[1]))
	v200 = v194 + v178<<(uint(int32(13))%32) + int32(-8192)
	goto L44
L48:
	;
	goto L37
L49:
	;
	v219 = F_heap_form_tuple(m, v153, v10+int32(16), v10+int32(12))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v222 = F_HeapTupleHeaderGetDatum(m, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	m.G0 = v10 + int32(48)
	return v222
L52:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v235 + int32(4)
	F_errmsg(m, int32(_a_F_pg_visibility_4), v10)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v243 = int32(*(*int8)(unsafe.Add(mBase, uint32(v242)+119)))
	F_errdetail_relkind_not_supported(m, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_pg_visibility_5), int32(932), int32(_a_F_pg_visibility_6))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(_a_F_pg_visibility_7), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_pg_visibility_5), int32(144), int32(_a_F_pg_visibility_8))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_visibility_map_summary(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v2
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v2)
	v18 = F_relation_open(m, v10, int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+119)))
		v25 = v23 - int32(109)
		v32 = int32(0)
		if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v25))|base.B2i32(int32(1)<<(uint(v25)%32)&int32(161) == v32) == v32 {
			F_visibilitymap_count(m, v18, v8+int32(44), v8+int32(40))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int64(0)
			} else {
				F_relation_close(m, v18, int32(1))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int64(0)
				} else {
					v49 = F_get_call_result_type(m, l0, int32(0), v8+int32(36))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
					} else {
						if v49 != int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int64(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_pg_visibility_map_summary_0), int32(0))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_visibility_map_summary_1), int32(289), int32(_a_F_pg_visibility_map_summary_2))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v53 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+44)))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v53
							v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+40)))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v55
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
							v62 = F_heap_form_tuple(m, v57, v8+int32(16), v8+int32(14))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int64(0)
							} else {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
								v65 = F_HeapTupleHeaderGetDatum(m, v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int64(0)
								} else {
									m.G0 = v8 + int32(48)
									return v65
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int64(0)
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v78 + int32(4)
					F_errmsg(m, int32(_a_F_pg_visibility_map_summary_3), v8)
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int64(0)
					} else {
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
						v86 = int32(*(*int8)(unsafe.Add(mBase, uint32(v85)+119)))
						F_errdetail_relkind_not_supported(m, v86)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_visibility_map_summary_1), int32(932), int32(_a_F_pg_visibility_map_summary_4))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int64(0)
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
func F_pg_visibility_rel(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int64
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int64
	_ = v214
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	goto L33
L4:
	;
	return int64(0)
L5:
	;
	v21 = int32(_a_F_pg_visibility_rel_0)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_visibility_rel[0]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_visibility_rel[0])) = v24
	v27 = F_CreateTemplateTupleDesc(m, int32(4))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitEntry(m, v27, int32(1), int32(_a_F_pg_visibility_rel_1), int32(20), int32(-1), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitEntry(m, v27, int32(2), int32(_a_F_pg_visibility_rel_2), int32(16), int32(-1), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v27, int32(3), int32(_a_F_pg_visibility_rel_3), int32(16), int32(-1), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v27, int32(4), int32(_a_F_pg_visibility_rel_4), int32(16), int32(-1), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v57 = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v57 < v66 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v144 = F_BlessTupleDesc(m, v27)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L30
	}
L12:
	;
	v70 = v27 + int32(28)
	v77 = v57
	v78 = v66
	v80 = v57
	goto L16
L13:
	;
	v134 = v57
	v141 = v66
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v134
	goto L11
L15:
	;
	v134 = v128
	v141 = v107
	goto L14
L16:
	;
	v86 = v70 + v66<<(uint(int32(3))%32) + v77*int32(100)
	v89 = v70 + v77<<(uint(int32(3))%32)
	if v66 != v78 {
		v107 = v78
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v128 = v66
	goto L15
L18:
	;
	v108 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+2)))
	if v108 <= int32(0) {
		v128 = v77
		goto L15
	} else {
		goto L26
	}
L19:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+7)))
	if v91 != int32(118) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v107 = v77
	goto L18
L21:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
	if v94 != int32(1) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+6)))
	if v97&int32(6) != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+2)))
	if v100 <= int32(0) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+90)))
	if v103 != int32(118) {
		v107 = v66
		goto L18
	} else {
		goto L25
	}
L25:
	;
	goto L20
L26:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+90)))
	if v111 == int32(118) {
		v128 = v77
		goto L15
	} else {
		goto L27
	}
L27:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+5)))
	v120 = (v80 + v114 - int32(1)) & (int32(0) - v114)
	if int32(_a_F_pg_visibility_rel_5) < v120 {
		v128 = v77
		goto L15
	} else {
		goto L28
	}
L28:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v89))) = uint16(v120)
	v126 = v77 + int32(1)
	if v126 != v66 {
		v77 = v126
		v78 = v107
		v80 = v120 + v108
		goto L16
	} else {
		goto L29
	}
L29:
	;
	goto L17
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v144
	v148 = F_collect_visibility_data(m, v16, int32(1))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v148
	*(*int32)(unsafe.Add(mBase, _c_F_pg_visibility_rel[0])) = v22
	goto L3
L32:
	;
	m.G0 = v10 + int32(48)
	return v214
L33:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if base.Ui32(v160) < base.Ui32(v161) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = base.I64_extend_i32_u(v160)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+v160)+8)))
	v169 = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = base.I64_extend_i32_u(v168 & v169)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = base.I64_extend_i32_u(int32(base.Ui32(v168)>>(uint(int32(2))%32)) & v169)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = base.I64_extend_i32_u(int32(base.Ui32(v168)>>(uint(v169)%32)) & v169)
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v160 + v169
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v158)+28))
	v193 = F_heap_form_tuple(m, v188, v10+int32(16), v10+int32(12))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v158)))
	*(*int64)(unsafe.Add(mBase, uint32(v158))) = v195 + int64(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+20)) = int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
	v203 = F_HeapTupleHeaderGetDatum(m, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v214 = v203
	goto L32
L39:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+20)) = int32(2)
	v210 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v210)
	v214 = int64(0)
	goto L32
}
func F_pg_vsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v45 int32
	_ = v45
	v5 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)) = uint8(v5)
	if l1 != 0 {
		v16 = l0
	} else {
		v16 = v8 + int32(7)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v16
	v19 = int32(1)
	if base.Ui32(l1) <= base.Ui32(v19) {
		v22 = v19
	} else {
		v22 = l1
	}
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16 + v22 - int32(1)
	F_dopr(m, v8+int32(8), l2, l3)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return int32(0)
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		v34 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v34)
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
		v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
		m.G0 = v8 + int32(32)
		if v38 != 0 {
			v45 = int32(-1)
		} else {
			v45 = v37 + (v33 - v36)
		}
		return v45
	}
}
func F_pg_wchar2single_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v8)
	return v8
L2:
	;
	goto L3
L3:
	;
	v12 = l0
	v13 = l1
	v15 = v4
	goto L5
L4:
	;
	v31 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v31)
	return v30
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v17 == int32(0) {
		v29 = v13
		v30 = v15
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v29 = v22
	v30 = l2
	goto L4
L7:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v17)
	v21 = int32(1)
	v22 = v13 + v21
	v26 = v15 + v21
	if v26 != l2 {
		v12 = v12 + int32(4)
		v13 = v22
		v15 = v26
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
func F_pg_xact_commit_timestamp_origin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_TransactionIdGetCommitTsData(m, v9, v7+int32(32), v7+int32(46))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v21 = F_get_call_result_type(m, l0, int32(0), v7+int32(8))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			if v21 == int32(1) {
				if v14 == int32(0) {
					v27 = int32(257)
					*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v27)
				} else {
					v29 = *(*int64)(unsafe.Add(mBase, uint32(v7)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v29
					v31 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v7)+46)))
					*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v31
					v33 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v33)
				}
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				v40 = F_heap_form_tuple(m, v35, v7+int32(16), v7+int32(14))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int64(0)
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
					v43 = F_HeapTupleHeaderGetDatum(m, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int64(0)
					} else {
						m.G0 = v7 + int32(48)
						return v43
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int64(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_pg_xact_commit_timestamp_origin_0), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_xact_commit_timestamp_origin_1), int32(487), int32(_a_F_pg_xact_commit_timestamp_origin_2))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int64(0)
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
