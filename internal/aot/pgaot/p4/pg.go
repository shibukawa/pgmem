package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PGSharedMemoryAttach(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	v5 = m.G0
	v7 = v5 - int32(192)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v11 = int32(2)
	v15 = F_pgl_shmctl(m, l0, v11, v7+int32(104))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(192)
	return v80
L2:
	;
	return int32(0)
L3:
	;
	if v15 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_PGSharedMemoryAttach[0]))
	switch v22 - int32(2) {
	case 0:
		goto L8
	default:
		goto L7
	case 22, 26:
		v80 = v11
		goto L1
	}
L5:
	;
	goto L6
L6:
	;
	v27 = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_PGSharedMemoryAttach[1]))
	v34 = F___fstatat(m, int32(-100), v29, v7+int32(8), v27)
	mBase = m.M
	goto L9
L7:
	;
	v80 = int32(0)
	goto L1
L8:
	;
	v80 = int32(3)
	goto L1
L9:
	;
	if v34 < int32(0) {
		v80 = v27
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_PGSharedMemoryAttach[2]))
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v54 == int32(-1) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	v41 = v39
	goto L15
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PGSharedMemoryAttach[0])) = int32(28)
	v54 = int32(-1)
	goto L11
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if l0 == v42 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v54 = v44
	goto L11
L18:
	;
	goto L19
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	if v45 != 0 {
		v41 = v45
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	v57 = int32(2)
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_PGSharedMemoryAttach[0]))
	switch v59 - v57 {
	case 0:
		goto L25
	default:
		goto L24
	case 22, 26:
		v80 = v57
		goto L1
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v54
	v65 = int32(3)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v66 != int32(679834894) {
		v80 = v65
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v80 = int32(0)
	goto L1
L25:
	;
	v80 = int32(3)
	goto L1
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v69 != v70 {
		v80 = v65
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v54)+32))
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v7)+96))
	if v72 != v73 {
		v80 = v65
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v7)+176))
	if v77 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v78 = int32(1)
	goto L31
L30:
	;
	v78 = int32(4)
	goto L31
L31:
	;
	v80 = v78
	goto L1
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
	F_errfinish(m, int32(_a_F_ProcessPgArchInterrupts_1), int32(900), int32(_a_F_ProcessPgArchInterrupts_2))
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
	F_errdetail(m, int32(_a_F_ProcessPgArchInterrupts_4), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_ProcessPgArchInterrupts_1), int32(882), int32(_a_F_ProcessPgArchInterrupts_2))
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
	F_errfinish(m, int32(_a_F_RemovePgTempRelationFiles_7), int32(3511), int32(_a_F_RemovePgTempRelationFiles_8))
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v244 int64
	_ = v244
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v422 int64
	_ = v422
	var v424 int64
	_ = v424
	var v426 int64
	_ = v426
	var v438 int32
	_ = v438
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v604 int64
	_ = v604
	var v606 int64
	_ = v606
	var v608 int64
	_ = v608
	var v620 int32
	_ = v620
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
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
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
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
	v669 = F_MemoryContextAllocZero(m, l1, int32(20))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L9
	} else {
		goto L173
	}
L2:
	;
	F_report_newlocale_failure(m, v47)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L9
	} else {
		goto L172
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L9
	} else {
		goto L169
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L9
	} else {
		goto L166
	}
L5:
	;
	v50 = F_text_to_cstring(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L20
	}
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[0]))
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
	v33 = F_SearchSysCache1(m, int32(16), l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
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
	v26 = F_text_to_cstring(m, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v30 = F_SysCacheGetAttrNotNull(m, int32(21), v16, int32(14))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v47 = v26
	v48 = v16
	v49 = v30
	goto L5
L15:
	;
	if v33 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v39 = F_SysCacheGetAttrNotNull(m, int32(16), v33, int32(8))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v41 = F_text_to_cstring(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v45 = F_SysCacheGetAttrNotNull(m, int32(16), v33, int32(9))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v47 = v41
	v48 = v33
	v49 = v45
	goto L5
L20:
	;
	F_ReleaseCatCache(m, v48)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if base.B2i32(v56 == int32(0))|base.B2i32(v56 != v59) != 0 {
		v77 = v56
		v78 = v59
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v77-v78 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	v62 = v47
	v63 = v50
	goto L25
L25:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v67 == int32(0) {
		v77 = v67
		v78 = v66
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v77 = v67
	v78 = v66
	goto L23
L27:
	;
	v70 = int32(1)
	if v67 == v66 {
		v62 = v62 + v70
		v63 = v63 + v70
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v82 == int32(67) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v264 != int32(67) {
		goto L74
	} else {
		goto L75
	}
L32:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v85 == int32(0) {
		v667 = v3
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v88 = int32(_a_F_create_pg_locale_libc_0)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[1])))
	if base.B2i32(v91 == int32(0))|base.B2i32(v91 != v94) != 0 {
		v112 = v91
		v113 = v94
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L34
L36:
	;
	if v112-v113 == int32(0) {
		v667 = v3
		goto L1
	} else {
		goto L43
	}
L37:
	;
	goto L36
L38:
	;
	v97 = v50
	v98 = v88
	goto L39
L39:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v102 == int32(0) {
		v112 = v102
		v113 = v101
		goto L37
	} else {
		goto L41
	}
L40:
	;
	v112 = v102
	v113 = v101
	goto L37
L41:
	;
	v105 = int32(1)
	if v102 == v101 {
		v97 = v97 + v105
		v98 = v98 + v105
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v118 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[2])) = v118
	v126 = m.G0
	v128 = v126 - int32(32)
	m.G0 = v128
	v133 = v118
	goto L47
L44:
	;
	if v256 != 0 {
		v667 = v256
		goto L1
	} else {
		goto L72
	}
L45:
	;
	m.G0 = v128 + int32(32)
	goto L44
L46:
	;
	v256 = int32(0)
	goto L45
L47:
	;
	v138 = v133 << (uint(int32(2)) % 32)
	v144 = int32(1) << (uint(v133) % 32) & int32(9)
	if v144|int32(1) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v163 = F___loc_is_allocated(m, v118)
	mBase = m.M
	if v163 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138+(v128+int32(8))))) = v155
	if v155 == int32(-1) {
		goto L46
	} else {
		goto L56
	}
L50:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v118+v138)))
	v155 = v151
	goto L49
L51:
	;
	goto L52
L52:
	;
	if v144 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v153 = v47
	goto L55
L54:
	;
	v153 = int32(_a_F_create_pg_locale_libc_1)
	goto L55
L55:
	;
	v154 = F___get_locale(m, v133, v153)
	mBase = m.M
	v155 = v154
	goto L49
L56:
	;
	v160 = v133 + int32(1)
	if v160 != int32(6) {
		v133 = v160
		goto L47
	} else {
		goto L57
	}
L57:
	;
	goto L48
L58:
	;
	v166 = int32(_a_F_create_pg_locale_libc_2)
	v168 = v128 + int32(8)
	v171 = F_memcmp(m, v168, v166, int32(24))
	mBase = m.M
	if v171 == int32(0) {
		v256 = v166
		goto L45
	} else {
		goto L61
	}
L59:
	;
	v235 = v118
	goto L60
L60:
	;
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v128)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v235)+16)) = v240
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v128)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v235)+8)) = v242
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v128)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v235))) = v244
	v256 = v235
	goto L45
L61:
	;
	v174 = int32(_a_F_create_pg_locale_libc_3)
	v177 = F_memcmp(m, v168, v174, int32(24))
	mBase = m.M
	if v177 == int32(0) {
		v256 = v174
		goto L45
	} else {
		goto L62
	}
L62:
	;
	v180 = int32(0)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[3])))
	if v182 == v180 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v188 = v180
	goto L66
L64:
	;
	goto L65
L65:
	;
	v215 = int32(_a_F_create_pg_locale_libc_4)
	v217 = v128 + int32(8)
	v220 = F_memcmp(m, v217, v215, int32(24))
	mBase = m.M
	if v220 == int32(0) {
		v256 = v215
		goto L45
	} else {
		goto L69
	}
L66:
	;
	v195 = F___get_locale(m, v188, int32(_a_F_create_pg_locale_libc_1))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v188<<(uint(int32(2))%32))+uint32(_c_F_create_pg_locale_libc[4]))) = v195
	v198 = v188 + int32(1)
	if v198 != int32(6) {
		v188 = v198
		goto L66
	} else {
		goto L68
	}
L67:
	;
	v202 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[3])) = uint8(v202)
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[5])) = v206
	goto L65
L68:
	;
	goto L67
L69:
	;
	v223 = int32(_a_F_create_pg_locale_libc_5)
	v226 = F_memcmp(m, v217, v223, int32(24))
	mBase = m.M
	if v226 == int32(0) {
		v256 = v223
		goto L45
	} else {
		goto L70
	}
L70:
	;
	v230 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v230 == int32(0) {
		goto L46
	} else {
		goto L71
	}
L71:
	;
	v235 = v230
	goto L60
L72:
	;
	goto L2
L73:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v449 != int32(67) {
		goto L116
	} else {
		goto L117
	}
L74:
	;
	v269 = int32(_a_F_create_pg_locale_libc_0)
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v275 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[1])))
	if base.B2i32(v272 == int32(0))|base.B2i32(v272 != v275) != 0 {
		v293 = v272
		v294 = v275
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	if v267 != 0 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v448 = int32(0)
	goto L73
L77:
	;
	if v293-v294 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L78:
	;
	goto L77
L79:
	;
	v278 = v47
	v279 = v269
	goto L80
L80:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+1)))
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+1)))
	if v283 == int32(0) {
		v293 = v283
		v294 = v282
		goto L78
	} else {
		goto L82
	}
L81:
	;
	v293 = v283
	v294 = v282
	goto L78
L82:
	;
	v286 = int32(1)
	if v283 == v282 {
		v278 = v278 + v286
		v279 = v279 + v286
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v448 = int32(0)
	goto L73
L85:
	;
	goto L86
L86:
	;
	v300 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[2])) = v300
	v308 = m.G0
	v310 = v308 - int32(32)
	m.G0 = v310
	v315 = v300
	goto L90
L87:
	;
	if v438 == int32(0) {
		goto L2
	} else {
		goto L115
	}
L88:
	;
	m.G0 = v310 + int32(32)
	goto L87
L89:
	;
	v438 = int32(0)
	goto L88
L90:
	;
	v320 = v315 << (uint(int32(2)) % 32)
	v326 = int32(1) << (uint(v315) % 32) & int32(8)
	if v326|int32(1) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v345 = F___loc_is_allocated(m, v300)
	mBase = m.M
	if v345 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320+(v310+int32(8))))) = v337
	if v337 == int32(-1) {
		goto L89
	} else {
		goto L99
	}
L93:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v300+v320)))
	v337 = v333
	goto L92
L94:
	;
	goto L95
L95:
	;
	if v326 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v335 = v47
	goto L98
L97:
	;
	v335 = int32(_a_F_create_pg_locale_libc_1)
	goto L98
L98:
	;
	v336 = F___get_locale(m, v315, v335)
	mBase = m.M
	v337 = v336
	goto L92
L99:
	;
	v342 = v315 + int32(1)
	if v342 != int32(6) {
		v315 = v342
		goto L90
	} else {
		goto L100
	}
L100:
	;
	goto L91
L101:
	;
	v348 = int32(_a_F_create_pg_locale_libc_2)
	v350 = v310 + int32(8)
	v353 = F_memcmp(m, v350, v348, int32(24))
	mBase = m.M
	if v353 == int32(0) {
		v438 = v348
		goto L88
	} else {
		goto L104
	}
L102:
	;
	v417 = v300
	goto L103
L103:
	;
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v310)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v417)+16)) = v422
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v310)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v417)+8)) = v424
	v426 = *(*int64)(unsafe.Add(mBase, uint32(v310)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v417))) = v426
	v438 = v417
	goto L88
L104:
	;
	v356 = int32(_a_F_create_pg_locale_libc_3)
	v359 = F_memcmp(m, v350, v356, int32(24))
	mBase = m.M
	if v359 == int32(0) {
		v438 = v356
		goto L88
	} else {
		goto L105
	}
L105:
	;
	v362 = int32(0)
	v364 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[3])))
	if v364 == v362 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v370 = v362
	goto L109
L107:
	;
	goto L108
L108:
	;
	v397 = int32(_a_F_create_pg_locale_libc_4)
	v399 = v310 + int32(8)
	v402 = F_memcmp(m, v399, v397, int32(24))
	mBase = m.M
	if v402 == int32(0) {
		v438 = v397
		goto L88
	} else {
		goto L112
	}
L109:
	;
	v377 = F___get_locale(m, v370, int32(_a_F_create_pg_locale_libc_1))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v370<<(uint(int32(2))%32))+uint32(_c_F_create_pg_locale_libc[4]))) = v377
	v380 = v370 + int32(1)
	if v380 != int32(6) {
		v370 = v380
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v384 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[3])) = uint8(v384)
	v388 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[5])) = v388
	goto L108
L111:
	;
	goto L110
L112:
	;
	v405 = int32(_a_F_create_pg_locale_libc_5)
	v408 = F_memcmp(m, v399, v405, int32(24))
	mBase = m.M
	if v408 == int32(0) {
		v438 = v405
		goto L88
	} else {
		goto L113
	}
L113:
	;
	v412 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v412 == int32(0) {
		goto L89
	} else {
		goto L114
	}
L114:
	;
	v417 = v412
	goto L103
L115:
	;
	v448 = v438
	goto L73
L116:
	;
	v453 = int32(_a_F_create_pg_locale_libc_0)
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v459 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[1])))
	if base.B2i32(v456 == int32(0))|base.B2i32(v456 != v459) != 0 {
		v477 = v456
		v478 = v459
		goto L120
	} else {
		goto L121
	}
L117:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v452 != 0 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v667 = v448
	goto L1
L119:
	;
	if v477-v478 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L120:
	;
	goto L119
L121:
	;
	v462 = v50
	v463 = v453
	goto L122
L122:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463)+1)))
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+1)))
	if v467 == int32(0) {
		v477 = v467
		v478 = v466
		goto L120
	} else {
		goto L124
	}
L123:
	;
	v477 = v467
	v478 = v466
	goto L120
L124:
	;
	v470 = int32(1)
	if v467 == v466 {
		v462 = v462 + v470
		v463 = v463 + v470
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v667 = v448
	goto L1
L127:
	;
	goto L128
L128:
	;
	v483 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[2])) = v483
	v490 = m.G0
	v492 = v490 - int32(32)
	m.G0 = v492
	v497 = v483
	goto L132
L129:
	;
	if v620 != 0 {
		v667 = v620
		goto L1
	} else {
		goto L157
	}
L130:
	;
	m.G0 = v492 + int32(32)
	goto L129
L131:
	;
	v620 = int32(0)
	goto L130
L132:
	;
	v502 = v497 << (uint(int32(2)) % 32)
	v508 = int32(1) << (uint(v497) % 32) & int32(1)
	v509 = int32(0)
	if v508|base.B2i32(v448 == v509) == v509 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	v527 = F___loc_is_allocated(m, v448)
	mBase = m.M
	if v527 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v502+(v492+int32(8))))) = v519
	if v519 == int32(-1) {
		goto L131
	} else {
		goto L141
	}
L135:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v448+v502)))
	v519 = v515
	goto L134
L136:
	;
	goto L137
L137:
	;
	if v508 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v517 = v50
	goto L140
L139:
	;
	v517 = int32(_a_F_create_pg_locale_libc_1)
	goto L140
L140:
	;
	v518 = F___get_locale(m, v497, v517)
	mBase = m.M
	v519 = v518
	goto L134
L141:
	;
	v524 = v497 + int32(1)
	if v524 != int32(6) {
		v497 = v524
		goto L132
	} else {
		goto L142
	}
L142:
	;
	goto L133
L143:
	;
	v530 = int32(_a_F_create_pg_locale_libc_2)
	v532 = v492 + int32(8)
	v535 = F_memcmp(m, v532, v530, int32(24))
	mBase = m.M
	if v535 == int32(0) {
		v620 = v530
		goto L130
	} else {
		goto L146
	}
L144:
	;
	v599 = v448
	goto L145
L145:
	;
	v604 = *(*int64)(unsafe.Add(mBase, uint32(v492)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v599)+16)) = v604
	v606 = *(*int64)(unsafe.Add(mBase, uint32(v492)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v599)+8)) = v606
	v608 = *(*int64)(unsafe.Add(mBase, uint32(v492)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v599))) = v608
	v620 = v599
	goto L130
L146:
	;
	v538 = int32(_a_F_create_pg_locale_libc_3)
	v541 = F_memcmp(m, v532, v538, int32(24))
	mBase = m.M
	if v541 == int32(0) {
		v620 = v538
		goto L130
	} else {
		goto L147
	}
L147:
	;
	v544 = int32(0)
	v546 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[3])))
	if v546 == v544 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v552 = v544
	goto L151
L149:
	;
	goto L150
L150:
	;
	v579 = int32(_a_F_create_pg_locale_libc_4)
	v581 = v492 + int32(8)
	v584 = F_memcmp(m, v581, v579, int32(24))
	mBase = m.M
	if v584 == int32(0) {
		v620 = v579
		goto L130
	} else {
		goto L154
	}
L151:
	;
	v559 = F___get_locale(m, v552, int32(_a_F_create_pg_locale_libc_1))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v552<<(uint(int32(2))%32))+uint32(_c_F_create_pg_locale_libc[4]))) = v559
	v562 = v552 + int32(1)
	if v562 != int32(6) {
		v552 = v562
		goto L151
	} else {
		goto L153
	}
L152:
	;
	v566 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[3])) = uint8(v566)
	v570 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[5])) = v570
	goto L150
L153:
	;
	goto L152
L154:
	;
	v587 = int32(_a_F_create_pg_locale_libc_5)
	v590 = F_memcmp(m, v581, v587, int32(24))
	mBase = m.M
	if v590 == int32(0) {
		v620 = v587
		goto L130
	} else {
		goto L155
	}
L155:
	;
	v594 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v594 == int32(0) {
		goto L131
	} else {
		goto L156
	}
L156:
	;
	v599 = v594
	goto L145
L157:
	;
	if v448 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v628 = F___loc_is_allocated(m, v448)
	mBase = m.M
	if v628 != 0 {
		goto L162
	} else {
		goto L163
	}
L159:
	;
	goto L160
L160:
	;
	F_report_newlocale_failure(m, v50)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L9
	} else {
		goto L165
	}
L161:
	;
	goto L160
L162:
	;
	F_emscripten_builtin_free(m, v448)
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
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v637
	F_errmsg_internal(m, int32(_a_F_create_pg_locale_libc_6), v9)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L9
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_create_pg_locale_libc_7), int32(435), int32(_a_F_create_pg_locale_libc_8))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
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
	v656 = m.ExcPending
	if v656 != 0 {
		goto L9
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(_a_F_create_pg_locale_libc_7), int32(452), int32(_a_F_create_pg_locale_libc_8))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
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
	v671 = int32(355)
	*(*uint16)(unsafe.Add(mBase, uint32(v669))) = uint16(v671)
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v673 != int32(67) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v669)+2)) = uint8(v707)
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v709 != int32(67) {
		goto L186
	} else {
		goto L187
	}
L175:
	;
	v678 = int32(_a_F_create_pg_locale_libc_0)
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v684 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[1])))
	if base.B2i32(v681 == int32(0))|base.B2i32(v681 != v684) != 0 {
		v702 = v681
		v703 = v684
		goto L179
	} else {
		goto L180
	}
L176:
	;
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	if v676 != 0 {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v707 = int32(1)
	goto L174
L178:
	;
	v707 = base.B2i32(v702-v703 == int32(0))
	goto L174
L179:
	;
	goto L178
L180:
	;
	v687 = v47
	v688 = v678
	goto L181
L181:
	;
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688)+1)))
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687)+1)))
	if v692 == int32(0) {
		v702 = v692
		v703 = v691
		goto L179
	} else {
		goto L183
	}
L182:
	;
	v702 = v692
	v703 = v691
	goto L179
L183:
	;
	v695 = int32(1)
	if v692 == v691 {
		v687 = v687 + v695
		v688 = v688 + v695
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v669)+12)) = v667
	*(*uint8)(unsafe.Add(mBase, uint32(v669)+3)) = uint8(v743)
	if v707 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L186:
	;
	v714 = int32(_a_F_create_pg_locale_libc_0)
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v720 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_pg_locale_libc[1])))
	if base.B2i32(v717 == int32(0))|base.B2i32(v717 != v720) != 0 {
		v738 = v717
		v739 = v720
		goto L190
	} else {
		goto L191
	}
L187:
	;
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v712 != 0 {
		goto L186
	} else {
		goto L188
	}
L188:
	;
	v743 = int32(1)
	goto L185
L189:
	;
	v743 = base.B2i32(v738-v739 == int32(0))
	goto L185
L190:
	;
	goto L189
L191:
	;
	v723 = v50
	v724 = v714
	goto L192
L192:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724)+1)))
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723)+1)))
	if v728 == int32(0) {
		v738 = v728
		v739 = v727
		goto L190
	} else {
		goto L194
	}
L193:
	;
	v738 = v728
	v739 = v727
	goto L190
L194:
	;
	v731 = int32(1)
	if v728 == v727 {
		v723 = v723 + v731
		v724 = v724 + v731
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v669)+8)) = int32(_a_F_create_pg_locale_libc_10)
	goto L198
L197:
	;
	goto L198
L198:
	;
	m.G0 = v9 + int32(32)
	return v669
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
func F_pg_available_extensions(m *base.Module, l0 int32) int32 {
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
	var v32 int32
	_ = v32
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
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
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
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
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
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
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
	return int32(0)
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
	m.G0 = v12 + int32(32)
	return int32(0)
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
	v32 = v2
	v35 = v2
	goto L7
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v32<<(uint(int32(2))%32))))
	v41 = F_AllocateDir(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L3
L9:
	;
	v217 = v32 + int32(1)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v217 < v218 {
		v32 = v217
		v35 = v215
		goto L7
	} else {
		goto L65
	}
L10:
	;
	if v41 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_pg_available_extensions[0]))
	if v46 == int32(44) {
		v215 = v35
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v57 = v35
	goto L15
L14:
	;
	goto L13
L15:
	;
	v58 = F_ReadDir(m, v41, v40)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	F_FreeDir(m, v41)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L64
	}
L17:
	;
	if v58 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v61 = v58 + int32(19)
	v65 = F_strlen(m, v61)
	mBase = m.M
	v72 = v65 + int32(1)
	goto L23
L19:
	;
	goto L20
L20:
	;
	goto L16
L21:
	;
	if v84 == int32(0) {
		goto L15
	} else {
		goto L27
	}
L22:
	;
	goto L21
L23:
	;
	v74 = int32(0)
	if v72 == v74 {
		v84 = v74
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v84 = v79
	goto L22
L25:
	;
	v78 = v72 - int32(1)
	v79 = v61 + v78
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v80 != int32(46) {
		v72 = v78
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v87 = int32(_a_F_pg_available_extensions_0)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_available_extensions[1])))
	if base.B2i32(v90 == int32(0))|base.B2i32(v90 != v93) != 0 {
		v111 = v90
		v112 = v93
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v111-v112 != 0 {
		goto L15
	} else {
		goto L35
	}
L29:
	;
	goto L28
L30:
	;
	v96 = v84
	v97 = v87
	goto L31
L31:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
	if v101 == int32(0) {
		v111 = v101
		v112 = v100
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v111 = v101
	v112 = v100
	goto L29
L33:
	;
	v104 = int32(1)
	if v101 == v100 {
		v96 = v96 + v104
		v97 = v97 + v104
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v114 = F_pstrdup(m, v61)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v119 = F_strlen(m, v114)
	mBase = m.M
	v126 = v119 + int32(1)
	goto L39
L37:
	;
	v139 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v139)
	v142 = F_strstr(m, v114, int32(_a_F_pg_available_extensions_1))
	mBase = m.M
	if v142 != 0 {
		goto L15
	} else {
		goto L43
	}
L38:
	;
	goto L37
L39:
	;
	v128 = int32(0)
	if v126 == v128 {
		v138 = v128
		goto L38
	} else {
		goto L41
	}
L40:
	;
	v138 = v133
	goto L38
L41:
	;
	v132 = v126 - int32(1)
	v133 = v114 + v132
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v134 != int32(46) {
		v126 = v132
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v143 = F_makeString(m, v114)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v145 = F_list_member(m, v57, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	if v145 != 0 {
		goto L15
	} else {
		goto L46
	}
L46:
	;
	v147 = F_lappend(m, v57, v143)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v150 = F_palloc0(m, int32(48))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v152 = F_pstrdup(m, v114)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+36)) = int32(-1)
	v156 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+34)) = uint8(v156)
	v158 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v150)+32)) = uint16(v158)
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v152
	v161 = F_pstrdup(m, v40)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+8)) = v161
	F_parse_extension_control_file(m, v150, int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v167 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v167
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)) = uint16(v167)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v167)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v178 = F_DirectFunctionCall1Coll(m, int32(500), v167, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v178
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v150)+16))
	if v181 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v150)+24))
	if v189 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L54:
	;
	v184 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v184)
	goto L53
L55:
	;
	goto L56
L56:
	;
	v186 = F_cstring_to_text(m, v181)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v186
	goto L53
L58:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	F_tuplestore_putvalues(m, v197, v198, v12+int32(16), v12+int32(12))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L63
	}
L59:
	;
	v192 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v192)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v194 = F_cstring_to_text(m, v189)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v194
	goto L58
L63:
	;
	v57 = v147
	goto L15
L64:
	;
	v215 = v57
	goto L9
L65:
	;
	goto L8
}
func F_pg_base64_decode_1(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L25
	} else {
		goto L48
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L25
	} else {
		goto L43
	}
L3:
	;
	v16 = l0 + l1
	v17 = l0
	v20 = l2
	v21 = v4
	v23 = v4
	v24 = v4
	goto L6
L4:
	;
	v155 = l2
	goto L5
L5:
	;
	m.G0 = v14 + int32(16)
	return base.I64_extend_i32_s(v155 - l2)
L6:
	;
	v29 = v17
	goto L9
L7:
	;
	if v151 != 0 {
		goto L1
	} else {
		goto L42
	}
L8:
	;
	goto L7
L9:
	;
	v39 = int32(1)
	v40 = v29 + v39
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v43 = v41 - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v43))|base.B2i32(v39<<(uint(v43)%32)&int32(_a_F_pg_base64_decode_1_0) == int32(0)) != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v41 == int32(61) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	goto L10
L12:
	;
	if base.Ui32(v40) < base.Ui32(v16) {
		v29 = v40
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v147 = v20
	v151 = v24
	goto L8
L14:
	;
	if base.Ui32(v40) < base.Ui32(v16) {
		v17 = v40
		v20 = v140
		v21 = v141
		v23 = v143
		v24 = v144
		goto L6
	} else {
		goto L41
	}
L15:
	;
	v140 = v133
	v141 = v134
	v143 = v136
	v144 = int32(0)
	goto L14
L16:
	;
	v133 = v130
	v134 = int32(0)
	v136 = v128
	goto L15
L17:
	;
	v102 = v99 + v21<<(uint(int32(6))%32)
	v104 = v24 + int32(1)
	if v104 != int32(4) {
		goto L32
	} else {
		goto L33
	}
L18:
	;
	if v23 != 0 {
		v99 = int32(0)
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if base.Ui32(int32(125)) < base.Ui32((v41-int32(1))&int32(255)) {
		goto L2
	} else {
		goto L30
	}
L21:
	;
	switch v24 - int32(2) {
	case 0:
		goto L24
	case 1:
		goto L23
	default:
		goto L22
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v63 = int32(2)
	v65 = int32(base.Ui32(v21) >> (uint(v63) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)) = uint8(v65)
	v68 = int32(base.Ui32(v21) >> (uint(int32(10)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v68)
	v128 = v63
	v130 = v20 + v63
	goto L16
L24:
	;
	v140 = v20
	v141 = v21 << (uint(int32(6)) % 32)
	v143 = int32(1)
	v144 = int32(3)
	goto L14
L25:
	;
	return int64(0)
L26:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(_a_F_pg_base64_decode_1_1), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_pg_base64_decode_1_2), int32(365), int32(_a_F_pg_base64_decode_1_3))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L25
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
	v96 = int32(*(*int8)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_pg_base64_decode_1[0]))))
	if v96 < int32(0) {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v99 = v96
	goto L17
L32:
	;
	v140 = v20
	v141 = v102
	v143 = v23
	v144 = v104
	goto L14
L33:
	;
	goto L34
L34:
	;
	v108 = int32(base.Ui32(v102) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v108)
	v110 = int32(0)
	if v23 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v113 = int32(1)
	v133 = v20 + v113
	v134 = v110
	v136 = v113
	goto L15
L36:
	;
	goto L37
L37:
	;
	v117 = int32(base.Ui32(v102) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)) = uint8(v117)
	if v23 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v133 = v20 + int32(2)
	v134 = v110
	v136 = v23
	goto L15
L39:
	;
	goto L40
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)) = uint8(v102)
	v128 = int32(0)
	v130 = v20 + int32(3)
	goto L16
L41:
	;
	v147 = v140
	v151 = v144
	goto L8
L42:
	;
	v155 = v147
	goto L5
L43:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L25
	} else {
		goto L44
	}
L44:
	;
	v177 = F_pg_mblen_range(m, v29, v16)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v177
	F_errmsg(m, int32(_a_F_pg_base64_decode_1_4), v14)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_pg_base64_decode_1_2), int32(378), int32(_a_F_pg_base64_decode_1_3))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L25
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L25
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(_a_F_pg_base64_decode_1_5), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L25
	} else {
		goto L50
	}
L50:
	;
	F_errhint(m, int32(_a_F_pg_base64_decode_1_6), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L25
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_pg_base64_decode_1_2), int32(399), int32(_a_F_pg_base64_decode_1_3))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L25
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_base64_encode(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	v4 = int32(0)
	if l1 == v4 {
		v98 = l2
	} else {
		v14 = l0
		v15 = int32(2)
		v17 = l2
		v18 = v4
		v19 = l2 + int32(76)
		for {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v25 = v21<<(uint(v15<<(uint(int32(3))%32))%32) | v18
			if int32(0) < v15 {
				v54 = v17
				v55 = v25
				v56 = v15 - int32(1)
			} else {
				v30 = int32(63)
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25&v30)+uint32(_c_F_pg_base64_encode[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v17)+3)) = uint8(v32)
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v25)>>(uint(int32(18))%32)))+uint32(_c_F_pg_base64_encode[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v36)
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v25)>>(uint(int32(6))%32))&v30)+uint32(_c_F_pg_base64_encode[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)) = uint8(v42)
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v25)>>(uint(int32(12))%32))&v30)+uint32(_c_F_pg_base64_encode[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)) = uint8(v48)
				v54 = v17 + int32(4)
				v55 = int32(0)
				v56 = int32(2)
			}
			if base.Ui32(v19) <= base.Ui32(v54) {
				v58 = int32(10)
				*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v58)
				v64 = v54 + int32(1)
				v65 = v54 + int32(77)
			} else {
				v64 = v54
				v65 = v19
			}
			v67 = v14 + int32(1)
			if base.Ui32(v67) < base.Ui32(l0+l1) {
				v14 = v67
				v15 = v56
				v17 = v64
				v18 = v55
				v19 = v65
				continue
			} else {
				break
			}
			break
		}
		if v56 == int32(2) {
			v98 = v64
		} else {
			v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v55)>>(uint(int32(18))%32)))+uint32(_c_F_pg_base64_encode[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v73)
			v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v55)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_pg_base64_encode[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)) = uint8(v79)
			if v56 == int32(0) {
				v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v55)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_pg_base64_encode[0]))))
				v89 = v88
			} else {
				v89 = int32(61)
			}
			v90 = int32(61)
			*(*uint8)(unsafe.Add(mBase, uint32(v64)+3)) = uint8(v90)
			*(*uint8)(unsafe.Add(mBase, uint32(v64)+2)) = uint8(v89)
			v98 = v64 + int32(4)
		}
	}
	return base.I64_extend_i32_s(v98 - l2)
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
func F_pg_buffercache_evict_all(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+30)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+28)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v2
	v27 = F_get_call_result_type(m, l0, v2, v12+int32(44))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L2
	} else {
		goto L51
	}
L2:
	;
	return int32(0)
L3:
	;
	if v27 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v33 = F_superuser(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
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
	v221 = m.ExcPending
	if v221 != 0 {
		goto L2
	} else {
		goto L48
	}
L7:
	;
	if v33 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v37 = m.G0
	v39 = v37 - int32(32)
	m.G0 = v39
	v42 = v12 + int32(24)
	v43 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v43
	v46 = v12 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v43
	v50 = v12 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v43
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_all[0]))
	if v43 < v54 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v62 = int32(1)
	goto L12
L10:
	;
	goto L11
L11:
	;
	v195 = int32(32)
	m.G0 = v39 + v195
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v200
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v209 = F_heap_form_tuple(m, v204, v12+v195, v12+int32(28))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L46
	}
L12:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_all[1]))
	v71 = v68 + v62<<(uint(int32(6))%32)
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_all[2]))
	if v73 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v77 = v71 - int32(40)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v78&int32(16777216) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v182 = v62 + int32(1)
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_all[0]))
	if v182 <= v184 {
		v62 = v182
		goto L12
	} else {
		goto L45
	}
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_all[3]))
	F_ResourceOwnerEnlarge(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+28)) = int32(_a_F_pg_buffercache_evict_all_0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = int32(_a_F_pg_buffercache_evict_all_1)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = int32(_a_F_pg_buffercache_evict_all_2)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+8)) = int64(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v102 = int32(_a_F_pg_buffercache_evict_all_3)
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v101 | v102
	if v101&v102 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	goto L25
L23:
	;
	goto L24
L24:
	;
	v136 = v39 + int32(8)
	v138 = int32(_a_F_pg_buffercache_evict_all_4)
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_all[4]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	if v141 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	F_perform_spin_delay(m, v39+int32(8))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L2
	} else {
		goto L27
	}
L26:
	;
	goto L24
L27:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v121 = int32(_a_F_pg_buffercache_evict_all_3)
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v120 | v121
	if v120&v121 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v158 = F_EvictUnpinnedBufferInternal(m, v71+int32(-64), v136)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L2
	} else {
		goto L40
	}
L30:
	;
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict_all[4])) = v156
	goto L30
L32:
	;
	if int32(999) < v139 {
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v139 < int32(11) {
		goto L30
	} else {
		goto L39
	}
L35:
	;
	v146 = int32(900)
	if v146 <= v139 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v149 = v146
	goto L38
L37:
	;
	v149 = v139
	goto L38
L38:
	;
	v156 = v149 + int32(100)
	goto L31
L39:
	;
	v156 = v139 - int32(1)
	goto L31
L40:
	;
	if v158 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v160 = v42
	goto L43
L42:
	;
	v160 = v46
	goto L43
L43:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v162 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v161 + v162
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+8)))
	if v165 != v162 {
		goto L18
	} else {
		goto L44
	}
L44:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v168 + int32(1)
	goto L18
L45:
	;
	goto L13
L46:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v209)+16))
	v212 = F_HeapTupleHeaderGetDatum(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	m.G0 = v12 + int32(48)
	return v212
L48:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_evict_all_5), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_all_6), int32(774), int32(_a_F_pg_buffercache_evict_all_7))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_pg_buffercache_evict_all_7)
	F_errmsg(m, int32(_a_F_pg_buffercache_evict_all_8), v12)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_all_6), int32(672), int32(_a_F_pg_buffercache_evict_all_9))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_column_is_updatable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.I32_extend16_s(v3) <= int32(0) {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = int32(0)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v14 = int32(16)
		v20 = F_bms_make_singleton(m, (v3<<(uint(v14)%32)+int32(_a_F_pg_column_is_updatable_0))>>(uint(v14)%32))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = F_relation_is_updatable(m, v9, v10, base.B2i32(v11 != v10), v20)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = int32(20)
				return base.B2i32(v24&v26 == v26)
			}
		}
	}
}
func F_pg_column_toast_chunk_id(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v10 == int32(0) {
		v14 = F_get_fn_expr_argtype(m, v9, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = F_get_typlen(m, v14)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v14
						F_errmsg_internal(m, int32(_a_F_pg_column_toast_chunk_id_0), v7)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_column_toast_chunk_id_1), int32(_a_F_pg_column_toast_chunk_id_2), int32(_a_F_pg_column_toast_chunk_id_3))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
					v25 = F_MemoryContextAlloc(m, v23, int32(4))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v25
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v30))) = v18
						v33 = v18
						if v33 != int32(-1) {
							v37 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v37)
							v52 = int32(0)
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
							if v41 == int32(1) {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
								if v44 == int32(18) {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v40)+10))
									v52 = v50
								} else {
									v47 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v47)
									v52 = int32(0)
								}
							} else {
								v47 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v47)
								v52 = int32(0)
							}
						}
						m.G0 = v7 + int32(16)
						return v52
					}
				}
			}
		}
	} else {
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v33 = v32
		if v33 != int32(-1) {
			v37 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v37)
			v52 = int32(0)
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
			if v41 == int32(1) {
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
				if v44 == int32(18) {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v40)+10))
					v52 = v50
				} else {
					v47 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v47)
					v52 = int32(0)
				}
			} else {
				v47 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v47)
				v52 = int32(0)
			}
		}
		m.G0 = v7 + int32(16)
		return v52
	}
}
func F_pg_control_init(m *base.Module, l0 int32) int32 {
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
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
	v6 = v4 - int32(80)
	m.G0 = v6
	v11 = F_get_call_result_type(m, l0, int32(0), v6+int32(16))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_init[0]))
			v22 = F_LWLockAcquire(m, v18+int32(1152), int32(1))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_init[1]))
				v28 = F_get_controlfile(m, v25, v6+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_init[0]))
					F_LWLockRelease(m, v31+int32(1152))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
						if v36 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_pg_control_init_0), int32(0))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_control_init_1), int32(222), int32(_a_F_pg_control_init_2))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+204))
							v40 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+20)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v39
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)+216))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+21)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v43
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v28)+220))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+22)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v47
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v28)+224))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+23)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v51
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v28)+228))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+24)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v55
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v28)+232))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+25)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v59
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v28)+236))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+26)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = v63
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v28)+240))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+27)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = v67
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v28)+244))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+28)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v71
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+248)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+29)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+68)) = v75
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v28)+252))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+30)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+72)) = v79
							v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+256)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+31)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+76)) = v83
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
							v92 = F_heap_form_tuple(m, v87, v6+int32(32), v6+int32(20))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
								v95 = F_HeapTupleHeaderGetDatum(m, v94)
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									m.G0 = v6 + int32(80)
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
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_control_init_3), int32(0))
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_control_init_1), int32(214), int32(_a_F_pg_control_init_2))
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return int32(0)
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
func F_pg_convert_from(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_convert_from[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = F_DirectFunctionCall1Coll(m, int32(500), v3, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = F_DirectFunctionCall3Coll(m, int32(1636), v3, v4, v5, v11)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func F_pg_create_restore_point(m *base.Module, l0 int32) int32 {
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
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L67
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L62
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+316))
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
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L57
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
	v213 = F_Int64GetDatum(m, v191)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L56
	}
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
	m.G0 = v7 + int32(16)
	return v213
L57:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(_a_F_pg_create_restore_point_4), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errhint(m, int32(_a_F_pg_create_restore_point_5), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_pg_create_restore_point_6), int32(243), int32(_a_F_pg_create_restore_point_7))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
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
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_pg_create_restore_point_8), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errhint(m, int32(_a_F_pg_create_restore_point_9), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_pg_create_restore_point_6), int32(249), int32(_a_F_pg_create_restore_point_7))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(63)
	F_errmsg(m, int32(_a_F_pg_create_restore_point_10), v7)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_pg_create_restore_point_6), int32(256), int32(_a_F_pg_create_restore_point_7))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
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
func F_pg_ddl_command_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13852(m, l0, int32(_a_F_pg_ddl_command_in_0), int32(358), int32(_a_F_pg_ddl_command_in_1), int32(_a_F_pg_ddl_command_in_2), int32(_a_F_pg_ddl_command_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_pg_dearmor(m *base.Module, l0 int32) int32 {
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
	var v86 int32
	_ = v86
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
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
			return int32(0)
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
				return int32(0)
			} else {
				if int32(0) <= v55 {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					v62 = F_palloc(m, v59+int32(4))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
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
							return int32(0)
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v77 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									return v62
								}
							} else {
								m.G0 = v9 + int32(16)
								return v62
							}
						}
					}
				} else {
					F_px_THROW_ERROR(m, v55)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
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
func F_pg_dependencies_out(m *base.Module, l0 int32) int32 {
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v121 float64
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	F_appendStringInfoChar(m, v22, int32(123))
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
	F_appendStringInfoChar(m, v10+int32(-16), int32(125))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L34
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
	v52 = v10 + int32(-16)
	F_appendStringInfoChar(m, v52, int32(34))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	if v56 <= int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v121 = *(*float64)(unsafe.Add(mBase, uint32(v43)))
	*(*float64)(unsafe.Add(mBase, uint32(v12))) = v121
	F_appendStringInfo(m, v10+int32(-16), int32(_a_F_pg_dependencies_out_1), v12)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L32
	}
L17:
	;
	if v56 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_appendStringInfoString(m, v52, int32(_a_F_pg_dependencies_out_2))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v64
	F_appendStringInfo(m, v10+int32(-16), int32(_a_F_pg_dependencies_out_3), v10+int32(-32))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	if v73 < int32(2) {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v79 = int32(1)
	v82 = v73
	goto L24
L24:
	;
	v89 = v10 + int32(-16)
	if v79 == v82-int32(1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L16
L26:
	;
	v95 = int32(_a_F_pg_dependencies_out_2)
	goto L28
L27:
	;
	v95 = int32(_a_F_pg_dependencies_out_0)
	goto L28
L28:
	;
	F_appendStringInfoString(m, v89, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43+int32(10)+v79<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v101
	F_appendStringInfo(m, v89, int32(_a_F_pg_dependencies_out_3), v10+int32(-48))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v109 = v79 + int32(1)
	v110 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	if v109 < v110 {
		v79 = v109
		v82 = v110
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L25
L32:
	;
	v129 = v36 + int32(1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if base.Ui32(v129) < base.Ui32(v130) {
		v36 = v129
		goto L9
	} else {
		goto L33
	}
L33:
	;
	goto L10
L34:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	m.G0 = v12 - int32(-64)
	return v146
}
func F_pg_dependencies_recv(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13852(m, l0, int32(_a_F_pg_dependencies_recv_0), int32(714), int32(_a_F_pg_dependencies_recv_1), int32(_a_F_pg_dependencies_recv_2), int32(_a_F_pg_dependencies_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
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
func F_pg_digest(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
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
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	return int32(0)
L2:
	;
	v18 = int32(1)
	v19 = v14 + v18
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v24 = v22 & v18
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v25 = v19
	goto L5
L4:
	;
	v25 = v14 + int32(4)
	goto L5
L5:
	;
	if v22 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v54 = F_downcase_truncate_identifier(m, v25, v52, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v31 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v42 = int32(1)
	if v24 != 0 {
		v52 = int32(base.Ui32(v22)>>(uint(v42)%32)) - v42
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v34 = int32(16)
	goto L12
L11:
	;
	v34 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v41 = int32(4)
	goto L15
L14:
	;
	v41 = v34
	goto L15
L15:
	;
	v52 = v41
	goto L6
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v58 = F_px_find_digest(m, v54, v11+int32(12))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v58 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_pfree(m, v54)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
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
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = m.T0[v65].(func(*base.Module, int32) int32)(m, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v69 = v66 + int32(4)
	v70 = F_palloc(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v69 << (uint(int32(2)) % 32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v76 = F_pg_detoast_datum_packed(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v108 = int32(1)
	if v78&v108 != 0 {
		goto L37
	} else {
		goto L38
	}
L26:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v78 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	if v84 == int32(18) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v95 = int32(1)
	if v78&v95 != 0 {
		v107 = int32(base.Ui32(v78)>>(uint(v95)%32)) - v95
		goto L25
	} else {
		goto L36
	}
L30:
	;
	v87 = int32(16)
	goto L32
L31:
	;
	v87 = int32(0)
	goto L32
L32:
	;
	if base.Ui32((v84-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v94 = int32(4)
	goto L35
L34:
	;
	v94 = v87
	goto L35
L35:
	;
	v107 = v94
	goto L25
L36:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
	goto L25
L37:
	;
	v112 = v108
	goto L39
L38:
	;
	v112 = int32(4)
	goto L39
L39:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	m.T0[v114].(func(*base.Module, int32, int32, int32))(m, v64, v76+v112, v107)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	m.T0[v119].(func(*base.Module, int32, int32))(m, v64, v70+int32(4))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	m.T0[v122].(func(*base.Module, int32))(m, v64)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v125 != v76 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_pfree(m, v76)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v129 != v14 {
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
	v132 = m.ExcPending
	if v132 != 0 {
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
	return v70
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
	if v58 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v54
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
	if v58 != v155 {
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
	if v58 != v161 {
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
func F_pg_event_trigger_table_rewrite_oid(m *base.Module, l0 int32) int32 {
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
			return v9
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50463299))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_pg_event_trigger_table_rewrite_oid_0)
					F_errmsg(m, int32(_a_F_pg_event_trigger_table_rewrite_oid_1), v5)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_event_trigger_table_rewrite_oid_2), int32(1630), int32(_a_F_pg_event_trigger_table_rewrite_oid_3))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
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
			return int32(0)
		} else {
			F_errcode(m, int32(50463299))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_pg_event_trigger_table_rewrite_oid_0)
				F_errmsg(m, int32(_a_F_pg_event_trigger_table_rewrite_oid_1), v5)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_event_trigger_table_rewrite_oid_2), int32(1630), int32(_a_F_pg_event_trigger_table_rewrite_oid_3))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
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
func F_pg_extension_update_paths(m *base.Module, l0 int32) int32 {
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
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
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
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v201 int32
	_ = v201
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	m.G0 = v13 + int32(48)
	return int32(0)
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
	v190 = v46
	goto L15
L15:
	;
	v201 = v54 + int32(1)
	if v201 < v190 {
		v46 = v190
		v54 = v201
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
	v190 = v177
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
	v177 = v64
	goto L20
L20:
	;
	v188 = v70 + int32(1)
	if v188 < v177 {
		v64 = v177
		v70 = v188
		goto L16
	} else {
		goto L41
	}
L21:
	;
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v84
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+28)) = uint16(v84)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+30)) = uint8(v84)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v93 = F_cstring_to_text(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v93
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v97 = F_cstring_to_text(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v97
	if v82 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	F_tuplestore_putvalues(m, v168, v169, v13+int32(32), v13+int32(28))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L40
	}
L25:
	;
	v102 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+30)) = uint8(v102)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v105 = v13 + int32(12)
	F_initStringInfo(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	F_appendStringInfoString(m, v105, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if int32(0) < v111 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v115 = int32(0)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v152 = F_cstring_to_text(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L38
	}
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125+v115<<(uint(int32(2))%32))))
	v131 = v13 + int32(12)
	F_appendStringInfoString(m, v131, int32(_a_F_pg_extension_update_paths_0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	F_appendStringInfoString(m, v131, v129)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v138 = v115 + int32(1)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v138 < v139 {
		v115 = v138
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v152
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_pfree(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L24
L40:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v177 = v176
	goto L20
L41:
	;
	goto L17
L42:
	;
	goto L12
}
func F_pg_function_is_visible(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_FunctionIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
	}
}
func F_pg_get_backend_memory_contexts(m *base.Module, l0 int32) int32 {
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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v162 int64
	_ = v162
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
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
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int64
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int64
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int64
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int64
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v401 int32
	_ = v401
	v16 = m.G0
	v18 = v16 - int32(1168)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = int64(34359738372)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_backend_memory_contexts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v24
	v31 = F_hash_create(m, int32(_a_F_pg_get_backend_memory_contexts_0), int32(256), v18+int32(16), int32(1064))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
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
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L87
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
	v61 = v42
	v63 = v44
	v64 = int32(0)
	goto L8
L8:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v64<<(uint(int32(2))%32))))
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
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v61
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
	F_errfinish(m, int32(_a_F_pg_get_backend_memory_contexts_2), int32(98), int32(_a_F_pg_get_backend_memory_contexts_3))
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
	v162 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1152)) = v162
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1144)) = v162
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1136)) = v162
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1128)) = v162
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1120)) = v162
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1104)) = v162
	v174 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+1112)) = uint16(v174)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v84)+32))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v84)+36))
	if v177 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v178 = int32(0)
	v179 = int32(_a_F_pg_get_backend_memory_contexts_4)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_backend_memory_contexts[2])))
	if base.B2i32(v182 == v178)|base.B2i32(v182 != v185) != 0 {
		v203 = v182
		v204 = v185
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v208 = v176
	v209 = v152
	goto L26
L26:
	;
	if v208 != 0 {
		goto L41
	} else {
		goto L42
	}
L27:
	;
	if v205 != 0 {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v205 = v203 - v204
	goto L27
L29:
	;
	v188 = v176
	v189 = v179
	goto L30
L30:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
	if v193 == int32(0) {
		v203 = v193
		v204 = v192
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v203 = v193
	v204 = v192
	goto L28
L32:
	;
	v196 = int32(1)
	if v193 == v192 {
		v188 = v188 + v196
		v189 = v189 + v196
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v206 = v177
	goto L36
L35:
	;
	v206 = v178
	goto L36
L36:
	;
	if v205 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v207 = v176
	goto L39
L38:
	;
	v207 = v177
	goto L39
L39:
	;
	v208 = v207
	v209 = v206
	goto L26
L40:
	;
	if v209 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v211 = F_cstring_to_text(m, v208)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v214 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+1104)) = uint8(v214)
	goto L40
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1120)) = v211
	goto L40
L45:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v239 = v237 - int32(474)
	if base.Ui32(v239) <= base.Ui32(int32(3)) {
		goto L57
	} else {
		goto L58
	}
L46:
	;
	v216 = F_strlen(m, v209)
	mBase = m.M
	if int32(1024) <= v216 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v234 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+1105)) = uint8(v234)
	goto L45
L49:
	;
	v220 = F_pg_mbcliplen(m, v209, v216, int32(1023))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	v222 = v216
	goto L51
L51:
	;
	if v222 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v222 = v220
	goto L51
L53:
	;
	base.MemoryCopy(m, v18-int32(-64), v209, v222)
	goto L55
L54:
	;
	goto L55
L55:
	;
	v227 = v18 - int32(-64)
	v229 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v222+v227))) = uint8(v229)
	v231 = F_cstring_to_text(m, v227)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1124)) = v231
	goto L45
L57:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v239<<(uint(int32(2))%32))+uint32(_c_F_pg_get_backend_memory_contexts[3])))
	v246 = v244
	goto L59
L58:
	;
	v246 = int32(_a_F_pg_get_backend_memory_contexts_5)
	goto L59
L59:
	;
	v247 = F_cstring_to_text(m, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1128)) = v247
	if v133 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v310 = F_construct_array_builtin(m, v298, v297, int32(23))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L71
	}
L62:
	;
	v252 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1132)) = v252
	v256 = F_palloc(m, v252)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1132)) = v258
	v262 = F_palloc(m, v258<<(uint(int32(2))%32))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L66
	}
L65:
	;
	v297 = v252
	v298 = v256
	goto L61
L66:
	;
	v264 = int32(0)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v265 <= v264 {
		v297 = v258
		v298 = v262
		goto L61
	} else {
		goto L67
	}
L67:
	;
	v270 = v264
	goto L68
L68:
	;
	v284 = v270 << (uint(int32(2)) % 32)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v286+v284)))
	*(*int32)(unsafe.Add(mBase, uint32(v262+v284))) = v288
	v291 = v270 + int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v291 < v292 {
		v270 = v291
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v297 = v258
	v298 = v262
	goto L61
L70:
	;
	goto L69
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1136)) = v310
	v313 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+1096)))
	v314 = F_Int64GetDatum(m, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1140)) = v314
	v317 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+1088)))
	v318 = F_Int64GetDatum(m, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1144)) = v318
	v321 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+1100)))
	v322 = F_Int64GetDatum(m, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1148)) = v322
	v325 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+1092)))
	v326 = F_Int64GetDatum(m, v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1152)) = v326
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1096))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1100))
	v333 = F_Int64GetDatum(m, base.I64_extend_i32_u(v329-v330))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1156)) = v333
	F_tuplestore_putvalues(m, v83, v82, v18+int32(1120), v18+int32(1104))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_list_free(m, v133)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+20))
	if v345 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v346 = v345
	v356 = v63
	goto L82
L80:
	;
	v374 = v63
	goto L81
L81:
	;
	v379 = int32(1)
	v382 = v64 + v379
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v382 < v383 {
		v61 = v61 + v379
		v63 = v374
		v64 = v382
		goto L8
	} else {
		goto L86
	}
L82:
	;
	v361 = F_lappend(m, v356, v346)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L84
	}
L83:
	;
	v374 = v361
	goto L81
L84:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v346)+28))
	if v363 != 0 {
		v346 = v363
		v356 = v361
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	goto L9
L87:
	;
	m.G0 = v18 + int32(1168)
	return int32(0)
}
func F_pg_get_constraintdef_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
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
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v412 int32
	_ = v412
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int64
	_ = v535
	var v542 int32
	_ = v542
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
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
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v746 int32
	_ = v746
	var v754 int32
	_ = v754
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	v12 = m.G0
	v14 = v12 - int32(432)
	m.G0 = v14
	v16 = F_GetTransactionSnapshot(m)
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
	v20 = F_RegisterSnapshot(m, v16)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = v14 + int32(304)
	F_ScanKeyInit(m, v27, int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v34 = int32(1)
	v36 = F_systable_beginscan(m, v24, int32(2667), v34, v20, v34, v27)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v38 = F_systable_getnext(m, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_UnregisterSnapshot(m, v20)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v38 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L219
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L216
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L1
	} else {
		goto L213
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L210
	}
L13:
	;
	m.G0 = v14 + int32(432)
	return v746
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
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+22)))
	v65 = v63 + v64
	v67 = v14 + int32(360)
	F_initStringInfo(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L25
	}
L17:
	;
	F_systable_endscan(m, v36)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
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
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v746 = int32(0)
	goto L13
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_0), v14)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(2232), int32(_a_F_pg_get_constraintdef_worker_2))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
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
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+72)))
	switch v139 - int32(99) {
	case 0:
		goto L51
	default:
		goto L48
	case 3:
		goto L54
	case 11:
		goto L50
	case 13:
		v289 = int32(_a_F_pg_get_constraintdef_worker_3)
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
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v65)+80))
	if v72 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v73 = F_generate_qualified_relation_name(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v65)+84))
	v88 = F_SearchSysCache1(m, int32(82), v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L34
	}
L31:
	;
	v77 = F_quote_identifier(m, v65+int32(4))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+292)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+288)) = v73
	F_appendStringInfo(m, v67, int32(_a_F_pg_get_constraintdef_worker_4), v14+int32(288))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L26
L34:
	;
	if v88 == int32(0) {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+22)))
	v94 = v92 + v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+68))
	v96 = F_get_namespace_name_or_temp(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v96 == int32(0) {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	v101 = v14 + int32(376)
	F_initStringInfo(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v104 = F_quote_identifier(m, v96)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+272)) = v104
	F_appendStringInfo(m, v101, int32(_a_F_pg_get_constraintdef_worker_5), v14+int32(272))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v114 = F_quote_identifier(m, v94+int32(4))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_appendStringInfoString(m, v101, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+376))
	F_ReleaseCatCache(m, v88)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v123 = F_quote_identifier(m, v65+int32(4))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+260)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v14)+256)) = v118
	F_appendStringInfo(m, v14+int32(360), int32(_a_F_pg_get_constraintdef_worker_6), v14+int32(256))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L26
L46:
	;
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+73)))
	if v702 == int32(1) {
		goto L194
	} else {
		goto L195
	}
L47:
	;
	F_appendStringInfoString(m, v14+int32(360), int32(_a_F_pg_get_constraintdef_worker_7))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L193
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L190
	}
L49:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v65)+88))
	v605 = F_SysCacheGetAttrNotNull(m, int32(19), v38, int32(27))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L178
	}
L50:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v65)+80))
	if v567 != 0 {
		goto L167
	} else {
		goto L168
	}
L51:
	;
	v461 = F_SysCacheGetAttrNotNull(m, int32(19), v38, int32(28))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L146
	}
L52:
	;
	v291 = v14 + int32(360)
	F_appendStringInfoString(m, v291, v289)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L96
	}
L53:
	;
	v289 = int32(_a_F_pg_get_constraintdef_worker_8)
	goto L52
L54:
	;
	v143 = v14 + int32(360)
	F_appendStringInfoString(m, v143, int32(_a_F_pg_get_constraintdef_worker_9))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v149 = F_SysCacheGetAttrNotNull(m, int32(19), v38, int32(21))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v65)+80))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+107)))
	v153 = F_decompile_column_index_array(m, v149, v151, v152, v143)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v65)+96))
	v157 = F_generate_relation_name(m, v155, int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v157
	F_appendStringInfo(m, v143, int32(_a_F_pg_get_constraintdef_worker_10), v14+int32(144))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v167 = F_SysCacheGetAttrNotNull(m, int32(19), v38, int32(22))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v65)+96))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+107)))
	v171 = F_decompile_column_index_array(m, v167, v169, v170, v143)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_appendStringInfoChar(m, v143, int32(41))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+102)))
	switch v177 - int32(102) {
	case 0:
		v198 = int32(_a_F_pg_get_constraintdef_worker_11)
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
	F_appendStringInfoString(m, v14+int32(360), v198)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L70
	}
L64:
	;
	v198 = int32(_a_F_pg_get_constraintdef_worker_12)
	goto L63
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	v198 = int32(_a_F_pg_get_constraintdef_worker_13)
	goto L63
L67:
	;
	v185 = int32(*(*int8)(unsafe.Add(mBase, uint32(v65)+102)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v185
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_14), v14-int32(-64))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(2309), int32(_a_F_pg_get_constraintdef_worker_2))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
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
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+100)))
	switch v204 - int32(97) {
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
		v226 = int32(_a_F_pg_get_constraintdef_worker_15)
		goto L72
	}
L71:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+101)))
	switch v237 - int32(97) {
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
		v259 = int32(_a_F_pg_get_constraintdef_worker_15)
		goto L82
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v226
	F_appendStringInfo(m, v14+int32(360), int32(_a_F_pg_get_constraintdef_worker_16), v14+int32(128))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L80
	}
L73:
	;
	v226 = int32(_a_F_pg_get_constraintdef_worker_17)
	goto L72
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	v226 = int32(_a_F_pg_get_constraintdef_worker_18)
	goto L72
L76:
	;
	v226 = int32(_a_F_pg_get_constraintdef_worker_19)
	goto L72
L77:
	;
	v213 = int32(*(*int8)(unsafe.Add(mBase, uint32(v65)+100)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v213
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_20), v14+int32(80))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(2335), int32(_a_F_pg_get_constraintdef_worker_2))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
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
	v273 = F_SysCacheGetAttr(m, int32(19), v38, int32(26), v14+int32(376))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L91
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v259
	F_appendStringInfo(m, v14+int32(360), int32(_a_F_pg_get_constraintdef_worker_21), v14+int32(112))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L90
	}
L83:
	;
	v259 = int32(_a_F_pg_get_constraintdef_worker_17)
	goto L82
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	v259 = int32(_a_F_pg_get_constraintdef_worker_18)
	goto L82
L86:
	;
	v259 = int32(_a_F_pg_get_constraintdef_worker_19)
	goto L82
L87:
	;
	v246 = int32(*(*int8)(unsafe.Add(mBase, uint32(v65)+101)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v246
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_22), v14+int32(96))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(2361), int32(_a_F_pg_get_constraintdef_worker_2))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
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
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+376)))
	if v275 != 0 {
		goto L46
	} else {
		goto L92
	}
L92:
	;
	v277 = v14 + int32(360)
	F_appendStringInfoString(m, v277, int32(_a_F_pg_get_constraintdef_worker_23))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v65)+80))
	v283 = F_decompile_column_index_array(m, v273, v281, int32(0), v277)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_appendStringInfoChar(m, v277, int32(41))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	goto L46
L96:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v65)+88))
	v296 = F_SearchSysCache1(m, int32(34), v295)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v296 == int32(0) {
		goto L10
	} else {
		goto L98
	}
L98:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+72)))
	if v300 != int32(117) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v314 = v14 + int32(360)
	F_appendStringInfoChar(m, v314, int32(40))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L103
	}
L100:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v296)+16))
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+22)))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303+v304)+13)))
	if v306 != int32(1) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	F_appendStringInfoString(m, v291, int32(_a_F_pg_get_constraintdef_worker_24))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	goto L99
L103:
	;
	v320 = F_SysCacheGetAttrNotNull(m, int32(19), v38, int32(21))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v65)+80))
	v324 = F_decompile_column_index_array(m, v320, v322, int32(0), v314)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+107)))
	if v326 == int32(1) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	F_appendStringInfoString(m, v314, int32(_a_F_pg_get_constraintdef_worker_25))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v333 = v14 + int32(360)
	F_appendStringInfoChar(m, v333, int32(41))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	v339 = F_SysCacheGetAttrNotNull(m, int32(34), v296, int32(3))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v324 < v339 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_appendStringInfoString(m, v333, int32(_a_F_pg_get_constraintdef_worker_26))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	F_ReleaseCatCache(m, v296)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L133
	}
L115:
	;
	v347 = F_SysCacheGetAttrNotNull(m, int32(34), v296, int32(16))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v349 = F_pg_detoast_datum(m, v347)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_deconstruct_array_builtin(m, v349, int32(21), v14+int32(376), int32(0), v14+int32(416))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v14)+416))
	if v324 < v359 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v361 = v324
	goto L122
L120:
	;
	goto L121
L121:
	;
	F_appendStringInfoChar(m, v14+int32(360), int32(41))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L132
	}
L122:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v65)+80))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v14)+376))
	v377 = int32(*(*int16)(unsafe.Add(mBase, uint32(v373+v361<<(uint(int32(2))%32)))))
	v379 = F_get_attname(m, v372, v377, int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	goto L121
L124:
	;
	if v324 < v361 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	F_appendStringInfoString(m, v14+int32(360), int32(_a_F_pg_get_constraintdef_worker_27))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v389 = F_quote_identifier(m, v379)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L129
	}
L128:
	;
	goto L127
L129:
	;
	F_appendStringInfoString(m, v14+int32(360), v389)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v394 = v361 + int32(1)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v14)+416))
	if v394 < v395 {
		v361 = v394
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
	v426 = int32(0)
	if base.B2i32(l1 == v426)|base.B2i32(v295 == v426) != 0 {
		goto L46
	} else {
		goto L134
	}
L134:
	;
	v431 = F_flatten_reloptions(m, v295)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	if v431 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v431
	F_appendStringInfo(m, v14+int32(360), int32(_a_F_pg_get_constraintdef_worker_28), v14+int32(240))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v443 = F_get_rel_tablespace(m, v295)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	F_pfree(m, v431)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	if v443 == int32(0) {
		goto L46
	} else {
		goto L142
	}
L142:
	;
	v447 = F_get_tablespace_name(m, v443)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v449 = F_quote_identifier(m, v447)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+224)) = v449
	F_appendStringInfo(m, v14+int32(360), int32(_a_F_pg_get_constraintdef_worker_29), v14+int32(224))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	goto L46
L146:
	;
	v463 = F_text_to_cstring(m, v461)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v465 = F_stringToNode(m, v463)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v65)+80))
	if v468 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v469 = F_get_rel_name(m, v468)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L152
	}
L150:
	;
	v523 = int32(0)
	goto L151
L151:
	;
	v528 = v14 + int32(416)
	F_initStringInfo(m, v528)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L161
	}
L152:
	;
	if v469 == int32(0) {
		goto L9
	} else {
		goto L153
	}
L153:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v65)+80))
	v475 = F_palloc0(m, int32(80))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v478 = F_palloc0(m, int32(136))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v478)+24)) = int32(1)
	v482 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v478)+21)) = uint8(v482)
	*(*int32)(unsafe.Add(mBase, uint32(v478)+16)) = v473
	v485 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v478)+12)) = v485
	*(*int32)(unsafe.Add(mBase, uint32(v478))) = int32(101)
	v490 = F_makeAlias(m, v469, v485)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v478)+8)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v478)+4)) = v490
	v494 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v478)+124)) = uint16(v494)
	v496 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v478)+20)) = uint8(v496)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+188)) = v478
	*(*int32)(unsafe.Add(mBase, uint32(v14)+376)) = v478
	v503 = F_list_make1_impl(m, int32(1), v14+int32(188))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v505 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v475)+20)) = v505
	*(*int64)(unsafe.Add(mBase, uint32(v475)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v475))) = v503
	F_set_rtable_names(m, v475, v505, v505)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_set_simple_column_names(m, v475)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+184)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v14)+416)) = v475
	v521 = F_list_make1_impl(m, int32(1), v14+int32(184))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v523 = v521
	goto L151
L161:
	;
	v531 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+408)) = uint8(v531)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+392)) = v531
	v535 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+384)) = v535
	*(*int32)(unsafe.Add(mBase, uint32(v14)+380)) = v523
	*(*int32)(unsafe.Add(mBase, uint32(v14)+412)) = v531
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+411)) = uint8(v531)
	v542 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+409)) = uint16(v542)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+400)) = v535
	*(*int32)(unsafe.Add(mBase, uint32(v14)+396)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v14)+376)) = v528
	F_get_rule_expr(m, v465, v14+int32(376), v531)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+106)))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v14)+416))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v554
	if v553 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v558 = int32(_a_F_pg_get_constraintdef_worker_30)
	goto L165
L164:
	;
	v558 = int32(_a_F_pg_get_constraintdef_worker_13)
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v558
	F_appendStringInfo(m, v14+int32(360), int32(_a_F_pg_get_constraintdef_worker_31), v14+int32(160))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	goto L46
L167:
	;
	v568 = F_extractNotNullColumn(m, v38)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v65)+84))
	if v593 == int32(0) {
		goto L46
	} else {
		goto L176
	}
L170:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v65)+80))
	v572 = F_get_attname(m, v570, v568, int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v574 = F_quote_identifier(m, v572)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = v574
	v578 = v14 + int32(360)
	F_appendStringInfo(m, v578, int32(_a_F_pg_get_constraintdef_worker_32), v14+int32(192))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584)+22)))
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584+v585)+106)))
	if v587 != int32(1) {
		goto L46
	} else {
		goto L174
	}
L174:
	;
	F_appendStringInfoString(m, v578, int32(_a_F_pg_get_constraintdef_worker_30))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	goto L46
L176:
	;
	F_appendStringInfoString(m, v14+int32(360), int32(_a_F_pg_get_constraintdef_worker_33))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	goto L46
L178:
	;
	v607 = F_pg_detoast_datum(m, v605)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	F_deconstruct_array_builtin(m, v607, int32(26), v14+int32(376), int32(0), v14+int32(416))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v14)+416))
	v620 = F_palloc(m, v617<<(uint(int32(2))%32))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v14)+416))
	if int32(0) < v622 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v625 = int32(0)
	goto L185
L183:
	;
	goto L184
L184:
	;
	v660 = int32(0)
	v666 = F_pg_get_indexdef_worker(m, v601, v660, v620, v660, v660, v660, v660, l2, v660)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L188
	}
L185:
	;
	v637 = v625 << (uint(int32(2)) % 32)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v14)+376))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v639+v637)))
	*(*int32)(unsafe.Add(mBase, uint32(v620+v637))) = v641
	v644 = v625 + int32(1)
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v14)+416))
	if v644 < v645 {
		v625 = v644
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
	F_appendStringInfoString(m, v14+int32(360), v666)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	goto L46
L190:
	;
	v674 = int32(*(*int8)(unsafe.Add(mBase, uint32(v65)+72)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v674
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_34), v14+int32(48))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(2592), int32(_a_F_pg_get_constraintdef_worker_2))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
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
	F_appendStringInfoString(m, v14+int32(360), int32(_a_F_pg_get_constraintdef_worker_35))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+74)))
	if v710 == int32(1) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	goto L196
L198:
	;
	F_appendStringInfoString(m, v14+int32(360), int32(_a_F_pg_get_constraintdef_worker_36))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+75)))
	if v720 != int32(1) {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	goto L200
L202:
	;
	F_systable_endscan(m, v36)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L208
	}
L203:
	;
	v726 = int32(_a_F_pg_get_constraintdef_worker_37)
	goto L205
L204:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+76)))
	if v724 != 0 {
		goto L202
	} else {
		goto L206
	}
L205:
	;
	F_appendStringInfoString(m, v14+int32(360), v726)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L207
	}
L206:
	;
	v726 = int32(_a_F_pg_get_constraintdef_worker_38)
	goto L205
L207:
	;
	goto L202
L208:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v14)+360))
	v746 = v734
	goto L13
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v87
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_39), v14+int32(16))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(_a_F_pg_get_constraintdef_worker_40), int32(_a_F_pg_get_constraintdef_worker_41))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
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
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v94)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v770
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_42), v14+int32(32))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(_a_F_pg_get_constraintdef_worker_43), int32(_a_F_pg_get_constraintdef_worker_41))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v295
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_44), v14+int32(208))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(2401), int32(_a_F_pg_get_constraintdef_worker_2))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = v468
	F_errmsg_internal(m, int32(_a_F_pg_get_constraintdef_worker_45), v14+int32(176))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(_a_F_pg_get_constraintdef_worker_1), int32(_a_F_pg_get_constraintdef_worker_46), int32(_a_F_pg_get_constraintdef_worker_47))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
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
func F_pg_get_expr(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_pg_get_expr_worker(m, v4, v8, int32(2))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			if v10 != 0 {
				v15 = v10
			} else {
				v12 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
				v15 = int32(0)
			}
			return v15
		}
	}
}
func F_pg_get_function_arg_default(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int64
	_ = v270
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_SearchSysCache1(m, int32(47), v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 == int32(0) {
			v25 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v25)
			v298 = v2
			m.G0 = v14 + int32(80)
			return v298
		} else {
			v33 = F_get_func_arg_info(m, v19, v14+int32(20), v14+int32(16), v14+int32(12))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				if base.B2i32(v33 < v16)|base.B2i32(v16 <= int32(0)) != 0 {
					F_ReleaseCatCache(m, v19)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						v58 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v58)
						v298 = int32(0)
						m.G0 = v14 + int32(80)
						return v298
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
					if v39 == int32(0) {
						v62 = int32(3)
						v63 = v16 & v62
						v64 = int32(0)
						if base.Ui32(v62) <= base.Ui32(v16-int32(1)) {
							v75 = v64
							v76 = int32(0)
							v77 = v2
							for {
								if v39 != 0 {
									v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v77))))
									v86 = v84 - int32(98)
									if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v86))|base.B2i32(int32(1)<<(uint(v86)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
										v99 = v75
									} else {
										v99 = v75 + int32(1)
									}
								} else {
									v99 = v75 + int32(1)
								}
								if v39 != 0 {
									v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v77)+1)))
									v104 = v102 - int32(98)
									if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v104))|base.B2i32(int32(1)<<(uint(v104)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
										v117 = v99
									} else {
										v117 = v99 + int32(1)
									}
								} else {
									v117 = v99 + int32(1)
								}
								if v39 != 0 {
									v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v77)+2)))
									v122 = v120 - int32(98)
									if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v122))|base.B2i32(int32(1)<<(uint(v122)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
										v135 = v117
									} else {
										v135 = v117 + int32(1)
									}
								} else {
									v135 = v117 + int32(1)
								}
								if v39 != 0 {
									v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v77)+3)))
									v140 = v138 - int32(98)
									if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v140))|base.B2i32(int32(1)<<(uint(v140)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
										v153 = v135
									} else {
										v153 = v135 + int32(1)
									}
								} else {
									v153 = v135 + int32(1)
								}
								v155 = int32(4)
								v156 = v77 + v155
								v158 = v76 + v155
								if v158 != v16&int32(2147483644) {
									v75 = v153
									v76 = v158
									v77 = v156
									continue
								} else {
									break
								}
								break
							}
							if v63 == int32(0) {
								v210 = v153
							} else {
								v165 = v153
								v167 = v156
								v176 = v165
								v178 = v167
								v183 = v2
								for {
									if v39 != 0 {
										v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v178))))
										v187 = v185 - int32(98)
										if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v187))|base.B2i32(int32(1)<<(uint(v187)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
											v200 = v176
										} else {
											v200 = v176 + int32(1)
										}
									} else {
										v200 = v176 + int32(1)
									}
									v202 = int32(1)
									v205 = v183 + v202
									if v205 != v63 {
										v176 = v200
										v178 = v178 + v202
										v183 = v205
										continue
									} else {
										break
									}
									break
								}
								v210 = v200
							}
						} else {
							v165 = v64
							v167 = v2
							v176 = v165
							v178 = v167
							v183 = v2
							for {
								if v39 != 0 {
									v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v178))))
									v187 = v185 - int32(98)
									if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v187))|base.B2i32(int32(1)<<(uint(v187)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
										v200 = v176
									} else {
										v200 = v176 + int32(1)
									}
								} else {
									v200 = v176 + int32(1)
								}
								v202 = int32(1)
								v205 = v183 + v202
								if v205 != v63 {
									v176 = v200
									v178 = v178 + v202
									v183 = v205
									continue
								} else {
									break
								}
								break
							}
							v210 = v200
						}
						v222 = F_SysCacheGetAttr(m, int32(47), v19, int32(24), v14+int32(11))
						mBase = m.M
						v223 = m.ExcPending
						if v223 != 0 {
							return int32(0)
						} else {
							v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
							if v224 == int32(1) {
								F_ReleaseCatCache(m, v19)
								mBase = m.M
								v228 = m.ExcPending
								if v228 != 0 {
									return int32(0)
								} else {
									v229 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v229)
									v298 = int32(0)
									m.G0 = v14 + int32(80)
									return v298
								}
							} else {
								v232 = F_text_to_cstring(m, v222)
								mBase = m.M
								v233 = m.ExcPending
								if v233 != 0 {
									return int32(0)
								} else {
									v234 = F_stringToNode(m, v232)
									mBase = m.M
									v235 = m.ExcPending
									if v235 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v232)
										mBase = m.M
										v237 = m.ExcPending
										if v237 != 0 {
											return int32(0)
										} else {
											v238 = int32(0)
											v241 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
											v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+22)))
											v243 = v241 + v242
											v244 = int32(*(*int16)(unsafe.Add(mBase, uint32(v243)+106)))
											v245 = int32(*(*int16)(unsafe.Add(mBase, uint32(v243)+104)))
											v249 = v210 + (v244 - v245) - int32(1)
											if base.B2i32(v234 == v238)|base.B2i32(v249 < v238) == v238 {
												v255 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
												if v249 < v255 {
													v261 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
													v265 = *(*int32)(unsafe.Add(mBase, uint32(v261+v249<<(uint(int32(2))%32))))
													v267 = v14 - int32(-64)
													F_initStringInfo(m, v267)
													mBase = m.M
													v269 = m.ExcPending
													if v269 != 0 {
														return int32(0)
													} else {
														v270 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v14)+28)) = v270
														*(*int64)(unsafe.Add(mBase, uint32(v14)+36)) = v270
														*(*int64)(unsafe.Add(mBase, uint32(v14)+44)) = v270
														*(*int64)(unsafe.Add(mBase, uint32(v14)+49)) = v270
														v278 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v278
														*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)) = uint8(v278)
														v282 = int32(1)
														*(*uint16)(unsafe.Add(mBase, uint32(v14)+57)) = uint16(v282)
														*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v267
														F_get_rule_expr(m, v265, v14+int32(24), v278)
														mBase = m.M
														v289 = m.ExcPending
														if v289 != 0 {
															return int32(0)
														} else {
															v290 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
															F_ReleaseCatCache(m, v19)
															mBase = m.M
															v292 = m.ExcPending
															if v292 != 0 {
																return int32(0)
															} else {
																v293 = F_cstring_to_text(m, v290)
																mBase = m.M
																v294 = m.ExcPending
																if v294 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v290)
																	mBase = m.M
																	v296 = m.ExcPending
																	if v296 != 0 {
																		return int32(0)
																	} else {
																		v298 = v293
																		m.G0 = v14 + int32(80)
																		return v298
																	}
																}
															}
														}
													}
												} else {
													F_ReleaseCatCache(m, v19)
													mBase = m.M
													v258 = m.ExcPending
													if v258 != 0 {
														return int32(0)
													} else {
														v259 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v259)
														v298 = v238
														m.G0 = v14 + int32(80)
														return v298
													}
												}
											} else {
												F_ReleaseCatCache(m, v19)
												mBase = m.M
												v258 = m.ExcPending
												if v258 != 0 {
													return int32(0)
												} else {
													v259 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v259)
													v298 = v238
													m.G0 = v14 + int32(80)
													return v298
												}
											}
										}
									}
								}
							}
						}
					} else {
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v16-int32(1)))))
						v47 = v45 - int32(98)
						if base.Ui32(int32(20)) < base.Ui32(v47) {
							F_ReleaseCatCache(m, v19)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								v58 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v58)
								v298 = int32(0)
								m.G0 = v14 + int32(80)
								return v298
							}
						} else {
							if int32(1)<<(uint(v47)%32)&int32(_a_F_pg_get_function_arg_default_0) != 0 {
								v62 = int32(3)
								v63 = v16 & v62
								v64 = int32(0)
								if base.Ui32(v62) <= base.Ui32(v16-int32(1)) {
									v75 = v64
									v76 = int32(0)
									v77 = v2
									for {
										if v39 != 0 {
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v77))))
											v86 = v84 - int32(98)
											if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v86))|base.B2i32(int32(1)<<(uint(v86)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
												v99 = v75
											} else {
												v99 = v75 + int32(1)
											}
										} else {
											v99 = v75 + int32(1)
										}
										if v39 != 0 {
											v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v77)+1)))
											v104 = v102 - int32(98)
											if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v104))|base.B2i32(int32(1)<<(uint(v104)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
												v117 = v99
											} else {
												v117 = v99 + int32(1)
											}
										} else {
											v117 = v99 + int32(1)
										}
										if v39 != 0 {
											v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v77)+2)))
											v122 = v120 - int32(98)
											if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v122))|base.B2i32(int32(1)<<(uint(v122)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
												v135 = v117
											} else {
												v135 = v117 + int32(1)
											}
										} else {
											v135 = v117 + int32(1)
										}
										if v39 != 0 {
											v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v77)+3)))
											v140 = v138 - int32(98)
											if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v140))|base.B2i32(int32(1)<<(uint(v140)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
												v153 = v135
											} else {
												v153 = v135 + int32(1)
											}
										} else {
											v153 = v135 + int32(1)
										}
										v155 = int32(4)
										v156 = v77 + v155
										v158 = v76 + v155
										if v158 != v16&int32(2147483644) {
											v75 = v153
											v76 = v158
											v77 = v156
											continue
										} else {
											break
										}
										break
									}
									if v63 == int32(0) {
										v210 = v153
									} else {
										v165 = v153
										v167 = v156
										v176 = v165
										v178 = v167
										v183 = v2
										for {
											if v39 != 0 {
												v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v178))))
												v187 = v185 - int32(98)
												if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v187))|base.B2i32(int32(1)<<(uint(v187)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
													v200 = v176
												} else {
													v200 = v176 + int32(1)
												}
											} else {
												v200 = v176 + int32(1)
											}
											v202 = int32(1)
											v205 = v183 + v202
											if v205 != v63 {
												v176 = v200
												v178 = v178 + v202
												v183 = v205
												continue
											} else {
												break
											}
											break
										}
										v210 = v200
									}
								} else {
									v165 = v64
									v167 = v2
									v176 = v165
									v178 = v167
									v183 = v2
									for {
										if v39 != 0 {
											v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v178))))
											v187 = v185 - int32(98)
											if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v187))|base.B2i32(int32(1)<<(uint(v187)%32)&int32(_a_F_pg_get_function_arg_default_0) == int32(0)) != 0 {
												v200 = v176
											} else {
												v200 = v176 + int32(1)
											}
										} else {
											v200 = v176 + int32(1)
										}
										v202 = int32(1)
										v205 = v183 + v202
										if v205 != v63 {
											v176 = v200
											v178 = v178 + v202
											v183 = v205
											continue
										} else {
											break
										}
										break
									}
									v210 = v200
								}
								v222 = F_SysCacheGetAttr(m, int32(47), v19, int32(24), v14+int32(11))
								mBase = m.M
								v223 = m.ExcPending
								if v223 != 0 {
									return int32(0)
								} else {
									v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
									if v224 == int32(1) {
										F_ReleaseCatCache(m, v19)
										mBase = m.M
										v228 = m.ExcPending
										if v228 != 0 {
											return int32(0)
										} else {
											v229 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v229)
											v298 = int32(0)
											m.G0 = v14 + int32(80)
											return v298
										}
									} else {
										v232 = F_text_to_cstring(m, v222)
										mBase = m.M
										v233 = m.ExcPending
										if v233 != 0 {
											return int32(0)
										} else {
											v234 = F_stringToNode(m, v232)
											mBase = m.M
											v235 = m.ExcPending
											if v235 != 0 {
												return int32(0)
											} else {
												F_pfree(m, v232)
												mBase = m.M
												v237 = m.ExcPending
												if v237 != 0 {
													return int32(0)
												} else {
													v238 = int32(0)
													v241 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
													v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+22)))
													v243 = v241 + v242
													v244 = int32(*(*int16)(unsafe.Add(mBase, uint32(v243)+106)))
													v245 = int32(*(*int16)(unsafe.Add(mBase, uint32(v243)+104)))
													v249 = v210 + (v244 - v245) - int32(1)
													if base.B2i32(v234 == v238)|base.B2i32(v249 < v238) == v238 {
														v255 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
														if v249 < v255 {
															v261 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
															v265 = *(*int32)(unsafe.Add(mBase, uint32(v261+v249<<(uint(int32(2))%32))))
															v267 = v14 - int32(-64)
															F_initStringInfo(m, v267)
															mBase = m.M
															v269 = m.ExcPending
															if v269 != 0 {
																return int32(0)
															} else {
																v270 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(v14)+28)) = v270
																*(*int64)(unsafe.Add(mBase, uint32(v14)+36)) = v270
																*(*int64)(unsafe.Add(mBase, uint32(v14)+44)) = v270
																*(*int64)(unsafe.Add(mBase, uint32(v14)+49)) = v270
																v278 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v278
																*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)) = uint8(v278)
																v282 = int32(1)
																*(*uint16)(unsafe.Add(mBase, uint32(v14)+57)) = uint16(v282)
																*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v267
																F_get_rule_expr(m, v265, v14+int32(24), v278)
																mBase = m.M
																v289 = m.ExcPending
																if v289 != 0 {
																	return int32(0)
																} else {
																	v290 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
																	F_ReleaseCatCache(m, v19)
																	mBase = m.M
																	v292 = m.ExcPending
																	if v292 != 0 {
																		return int32(0)
																	} else {
																		v293 = F_cstring_to_text(m, v290)
																		mBase = m.M
																		v294 = m.ExcPending
																		if v294 != 0 {
																			return int32(0)
																		} else {
																			F_pfree(m, v290)
																			mBase = m.M
																			v296 = m.ExcPending
																			if v296 != 0 {
																				return int32(0)
																			} else {
																				v298 = v293
																				m.G0 = v14 + int32(80)
																				return v298
																			}
																		}
																	}
																}
															}
														} else {
															F_ReleaseCatCache(m, v19)
															mBase = m.M
															v258 = m.ExcPending
															if v258 != 0 {
																return int32(0)
															} else {
																v259 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v259)
																v298 = v238
																m.G0 = v14 + int32(80)
																return v298
															}
														}
													} else {
														F_ReleaseCatCache(m, v19)
														mBase = m.M
														v258 = m.ExcPending
														if v258 != 0 {
															return int32(0)
														} else {
															v259 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v259)
															v298 = v238
															m.G0 = v14 + int32(80)
															return v298
														}
													}
												}
											}
										}
									}
								}
							} else {
								F_ReleaseCatCache(m, v19)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									v58 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v58)
									v298 = int32(0)
									m.G0 = v14 + int32(80)
									return v298
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_get_function_arguments(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13961(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pg_get_function_sqlbody(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
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
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_initStringInfo(m, v6+int32(16))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v16 = F_SearchSysCache1(m, int32(47), v8)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v16 == int32(0) {
				v20 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
				v47 = int32(0)
				m.G0 = v6 + int32(32)
				return v47
			} else {
				v27 = F_SysCacheGetAttr(m, int32(47), v16, int32(28), v6+int32(15))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
					if v29 == int32(1) {
						F_ReleaseCatCache(m, v16)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
							v47 = int32(0)
							m.G0 = v6 + int32(32)
							return v47
						}
					} else {
						F_print_function_sqlbody(m, v6+int32(16), v16)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							F_ReleaseCatCache(m, v16)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
								v45 = F_cstring_to_text_with_len(m, v43, v44)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									v47 = v45
									m.G0 = v6 + int32(32)
									return v47
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_get_functiondef(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
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
	var v104 int32
	_ = v104
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
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
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
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v230 int32
	_ = v230
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v310 float32
	_ = v310
	var v313 int32
	_ = v313
	var v318 float32
	_ = v318
	var v328 int32
	_ = v328
	var v329 float32
	_ = v329
	var v332 int32
	_ = v332
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
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
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v589 int32
	_ = v589
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v676 int32
	_ = v676
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	v13 = m.G0
	v15 = v13 - int32(192)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v25 = F_SearchSysCache1(m, int32(47), v17)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L231
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L227
	}
L5:
	;
	m.G0 = v15 + int32(192)
	return v803
L6:
	;
	if v25 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
	v803 = int32(0)
	goto L5
L8:
	;
	goto L9
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
	v34 = v32 + v33
	v36 = v34 + int32(4)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+96)))
	if v37 == int32(97) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v43 = base.B2i32(v37 == int32(112))
	if v37 == int32(112) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v44 = int32(_a_F_pg_get_functiondef_0)
	goto L13
L12:
	;
	v44 = int32(_a_F_pg_get_functiondef_1)
	goto L13
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	v46 = F_get_namespace_name_or_temp(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v49 = v15 + int32(176)
	F_initStringInfo(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v46 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v52 = F_quote_identifier(m, v46)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v62 = F_quote_identifier(m, v36)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v52
	F_appendStringInfo(m, v49, int32(_a_F_pg_get_functiondef_2), v15+int32(144))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	F_appendStringInfoString(m, v15+int32(176), v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
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
		goto L23
	}
L23:
	;
	v78 = F_print_function_arguments(m, v70, v25, int32(0), int32(1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_appendStringInfoString(m, v70, int32(_a_F_pg_get_functiondef_4))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v43 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_appendStringInfoString(m, v70, int32(_a_F_pg_get_functiondef_5))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
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
		goto L33
	}
L29:
	;
	F_print_function_rettype(m, v70, v25)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_appendStringInfoChar(m, v70, int32(10))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	if int32(0) < v141 {
		goto L50
	} else {
		goto L51
	}
L33:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+15)))
	if v106 != 0 {
		v141 = int32(0)
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L47
	}
L35:
	;
	m.G0 = v98 + int32(16)
	goto L32
L36:
	;
	v107 = F_pg_detoast_datum(m, v104)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v109 != int32(1) {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
	if v112 < int32(0) {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	if v115 != 0 {
		goto L34
	} else {
		goto L40
	}
L40:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	if v116 != int32(26) {
		goto L34
	} else {
		goto L41
	}
L41:
	;
	v120 = v112 << (uint(int32(2)) % 32)
	v121 = F_palloc(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(176)))) = v121
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	if v124 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v134 = (v127<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L45
L44:
	;
	v134 = v124
	goto L45
L45:
	;
	if v120 == int32(0) {
		v141 = v112
		goto L35
	} else {
		goto L46
	}
L46:
	;
	base.MemoryCopy(m, v121, v134+v107, v120)
	v141 = v112
	goto L35
L47:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_functiondef_6), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_pg_get_functiondef_7), int32(1500), int32(_a_F_pg_get_functiondef_8))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v164 = v15 + int32(160)
	F_appendStringInfoString(m, v164, int32(_a_F_pg_get_functiondef_9))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	v245 = F_get_language_name(m, v243, int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L66
	}
L53:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v170 = F_format_type_be(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v170
	F_appendStringInfo(m, v164, int32(_a_F_pg_get_functiondef_10), v15+int32(112))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v178 = int32(1)
	if v141 != v178 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v181 = v178
	goto L59
L57:
	;
	goto L58
L58:
	;
	F_appendStringInfoChar(m, v15+int32(160), int32(10))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L65
	}
L59:
	;
	v194 = v15 + int32(160)
	F_appendStringInfoString(m, v194, int32(_a_F_pg_get_functiondef_11))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	goto L58
L61:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198+v181<<(uint(int32(2))%32))))
	v203 = F_format_type_be(m, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v203
	F_appendStringInfo(m, v194, int32(_a_F_pg_get_functiondef_10), v15+int32(96))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v212 = v181 + int32(1)
	if v212 != v141 {
		v181 = v212
		goto L59
	} else {
		goto L64
	}
L64:
	;
	goto L60
L65:
	;
	goto L52
L66:
	;
	v247 = F_quote_identifier(m, v245)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v247
	v251 = v15 + int32(160)
	F_appendStringInfo(m, v251, int32(_a_F_pg_get_functiondef_12), v15+int32(80))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+96)))
	if v258 == int32(119) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_appendStringInfoString(m, v251, int32(_a_F_pg_get_functiondef_13))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+101)))
	switch v265 - int32(105) {
	case 0:
		v269 = int32(_a_F_pg_get_functiondef_14)
		goto L74
	default:
		goto L73
	case 10:
		goto L75
	}
L72:
	;
	goto L71
L73:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+102)))
	switch v276 - int32(114) {
	case 0:
		goto L79
	case 1:
		v280 = int32(_a_F_pg_get_functiondef_15)
		goto L78
	default:
		goto L77
	}
L74:
	;
	F_appendStringInfoString(m, v15+int32(160), v269)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L76
	}
L75:
	;
	v269 = int32(_a_F_pg_get_functiondef_16)
	goto L74
L76:
	;
	goto L73
L77:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+99)))
	if v286 == int32(1) {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	F_appendStringInfoString(m, v15+int32(160), v280)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L80
	}
L79:
	;
	v280 = int32(_a_F_pg_get_functiondef_17)
	goto L78
L80:
	;
	goto L77
L81:
	;
	F_appendStringInfoString(m, v15+int32(160), int32(_a_F_pg_get_functiondef_18))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+97)))
	if v294 == int32(1) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L83
L85:
	;
	F_appendStringInfoString(m, v15+int32(160), int32(_a_F_pg_get_functiondef_19))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+98)))
	if v302 == int32(1) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L87
L89:
	;
	F_appendStringInfoString(m, v15+int32(160), int32(_a_F_pg_get_functiondef_20))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v310 = *(*float32)(unsafe.Add(mBase, uint32(v34)+80))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	if v313&int32(-2) == int32(12) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L91
L93:
	;
	v318 = float32(1)
	goto L95
L94:
	;
	v318 = float32(100)
	goto L95
L95:
	;
	if base.F32_ne(v310, v318) != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v15)+64)) = base.F64_promote_f32(v310)
	F_appendStringInfo(m, v15+int32(160), int32(_a_F_pg_get_functiondef_21), v15-int32(-64))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v329 = *(*float32)(unsafe.Add(mBase, uint32(v34)+84))
	v332 = int32(0)
	if base.B2i32(base.F32_gt(v329, float32(0)) == v332)|base.F32_eq(v329, float32(1000)) == v332 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L98
L100:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v15)+48)) = base.F64_promote_f32(v329)
	F_appendStringInfo(m, v15+int32(160), int32(_a_F_pg_get_functiondef_22), v15+int32(48))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v34)+92))
	if v348 != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L102
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = int32(2281)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v34)+92))
	v353 = int32(0)
	v359 = F_generate_function_name(m, v351, int32(1), v353, v15+int32(176), v353, v353, v353)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v15)+164))
	if v369 != v257 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v359
	F_appendStringInfo(m, v15+int32(160), int32(_a_F_pg_get_functiondef_23), v15+int32(32))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	F_appendStringInfoChar(m, v15+int32(160), int32(10))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v380 = F_SysCacheGetAttr(m, int32(47), v25, int32(29), v15+int32(159))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+159)))
	if v382 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v624 = F_SysCacheGetAttr(m, int32(47), v25, int32(28), v15+int32(159))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L177
	}
L115:
	;
	v383 = F_pg_detoast_datum(m, v380)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = int32(1)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v383)+16))
	if v387 <= int32(0) {
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L118
L118:
	;
	v406 = F_array_ref(m, v383, v15+int32(176), v15+int32(159))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L120
	}
L119:
	;
	goto L114
L120:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+159)))
	if v408 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	v604 = v602 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v604
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v383)+16))
	if v604 <= v606 {
		goto L118
	} else {
		goto L176
	}
L122:
	;
	v409 = F_text_to_cstring(m, v406)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v411 = int32(61)
	v412 = F___strchrnul(m, v409, v411)
	mBase = m.M
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412))))
	if v414 == v411 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	if v418 == int32(0) {
		goto L121
	} else {
		goto L128
	}
L125:
	;
	v418 = v412
	goto L127
L126:
	;
	v418 = int32(0)
	goto L127
L127:
	;
	goto L124
L128:
	;
	v421 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v418))) = uint8(v421)
	v423 = F_quote_identifier(m, v409)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v423
	F_appendStringInfo(m, v15+int32(160), int32(_a_F_pg_get_functiondef_24), v15+int32(16))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v434 = v418 + int32(1)
	v435 = F_GetConfigOptionFlags(m, v409)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L132
	}
L131:
	;
	F_appendStringInfoChar(m, v15+int32(160), int32(10))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L175
	}
L132:
	;
	if v435&int32(2) != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v441 = F_SplitGUCList(m, v434, v15+int32(152))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	F_appendStringInfoChar(m, v15+int32(160), int32(39))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L161
	}
L136:
	;
	if v441 == int32(0) {
		goto L3
	} else {
		goto L137
	}
L137:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v15)+152))
	if v445 == int32(0) {
		goto L131
	} else {
		goto L138
	}
L138:
	;
	v448 = int32(0)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	if v449 <= v448 {
		goto L131
	} else {
		goto L139
	}
L139:
	;
	v457 = v448
	goto L140
L140:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v445)+12))
	v467 = v464 + v457<<(uint(int32(2))%32)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	F_appendStringInfoChar(m, v15+int32(160), int32(39))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L142
	}
L141:
	;
	goto L131
L142:
	;
	v474 = v468
	goto L143
L143:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	v487 = base.I32_extend8_s(v486)
	if v486 != int32(39) {
		goto L147
	} else {
		goto L148
	}
L144:
	;
	v509 = v15 + int32(160)
	F_appendStringInfoChar(m, v509, int32(39))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L155
	}
L145:
	;
	goto L144
L146:
	;
	F_appendStringInfoChar(m, v15+int32(160), v487)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L154
	}
L147:
	;
	if v486 == int32(0) {
		goto L145
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	F_appendStringInfoChar(m, v15+int32(160), v487)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L153
	}
L150:
	;
	if v487 != int32(92) {
		goto L146
	} else {
		goto L151
	}
L151:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_functiondef[0])))
	if v495&int32(1) != 0 {
		goto L146
	} else {
		goto L152
	}
L152:
	;
	goto L149
L153:
	;
	goto L146
L154:
	;
	v474 = v474 + int32(1)
	goto L143
L155:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v15)+152))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v515)+12))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v515)+4))
	if base.Ui32(v467+int32(4)) < base.Ui32(v516+v517<<(uint(int32(2))%32)) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	F_appendStringInfoString(m, v509, int32(_a_F_pg_get_functiondef_11))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v526 = v457 + int32(1)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	if v526 < v527 {
		v457 = v526
		goto L140
	} else {
		goto L160
	}
L159:
	;
	goto L158
L160:
	;
	goto L141
L161:
	;
	v534 = v434
	goto L162
L162:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534))))
	v547 = base.I32_extend8_s(v546)
	if v546 != int32(39) {
		goto L166
	} else {
		goto L167
	}
L163:
	;
	F_appendStringInfoChar(m, v15+int32(160), int32(39))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L174
	}
L164:
	;
	goto L163
L165:
	;
	F_appendStringInfoChar(m, v15+int32(160), v547)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L173
	}
L166:
	;
	if v546 == int32(0) {
		goto L164
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	F_appendStringInfoChar(m, v15+int32(160), v547)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L172
	}
L169:
	;
	if v547 != int32(92) {
		goto L165
	} else {
		goto L170
	}
L170:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_functiondef[0])))
	if v555&int32(1) != 0 {
		goto L165
	} else {
		goto L171
	}
L171:
	;
	goto L168
L172:
	;
	goto L165
L173:
	;
	v534 = v534 + int32(1)
	goto L162
L174:
	;
	goto L131
L175:
	;
	goto L121
L176:
	;
	goto L119
L177:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	if v626 != int32(14) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	F_appendStringInfoChar(m, v15+int32(160), int32(10))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L223
	}
L179:
	;
	v637 = v15 + int32(160)
	F_appendStringInfoString(m, v637, int32(_a_F_pg_get_functiondef_25))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L183
	}
L180:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+159)))
	if v629&int32(1) != 0 {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	F_print_function_sqlbody(m, v15+int32(160), v25)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	goto L178
L183:
	;
	v645 = F_SysCacheGetAttr(m, int32(47), v25, int32(27), v15+int32(159))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+159)))
	if v647 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v650 = F_text_to_cstring(m, v645)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v711 = F_SysCacheGetAttrNotNull(m, int32(47), v25, int32(26))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L204
	}
L188:
	;
	F_appendStringInfoChar(m, v637, int32(39))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v655 = v650
	goto L190
L190:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655))))
	v668 = base.I32_extend8_s(v667)
	if v667 != int32(39) {
		goto L194
	} else {
		goto L195
	}
L191:
	;
	v690 = v15 + int32(160)
	F_appendStringInfoChar(m, v690, int32(39))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L202
	}
L192:
	;
	goto L191
L193:
	;
	F_appendStringInfoChar(m, v15+int32(160), v668)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L201
	}
L194:
	;
	if v667 == int32(0) {
		goto L192
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	F_appendStringInfoChar(m, v15+int32(160), v668)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L200
	}
L197:
	;
	if v668 != int32(92) {
		goto L193
	} else {
		goto L198
	}
L198:
	;
	v676 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_functiondef[0])))
	if v676&int32(1) != 0 {
		goto L193
	} else {
		goto L199
	}
L199:
	;
	goto L196
L200:
	;
	goto L193
L201:
	;
	v655 = v655 + int32(1)
	goto L190
L202:
	;
	F_appendStringInfoString(m, v690, int32(_a_F_pg_get_functiondef_11))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	goto L187
L204:
	;
	v713 = F_text_to_cstring(m, v711)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v716 = v15 + int32(176)
	F_initStringInfo(m, v716)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	F_appendStringInfoChar(m, v716, int32(36))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	if v37 == int32(112) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v726 = int32(_a_F_pg_get_functiondef_26)
	goto L210
L209:
	;
	v726 = int32(_a_F_pg_get_functiondef_27)
	goto L210
L210:
	;
	F_appendStringInfoString(m, v716, v726)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	v730 = F_strstr(m, v713, v729)
	mBase = m.M
	if v730 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	goto L215
L213:
	;
	goto L214
L214:
	;
	F_appendStringInfoChar(m, v15+int32(176), int32(36))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L219
	}
L215:
	;
	F_appendStringInfoChar(m, v15+int32(176), int32(120))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L217
	}
L216:
	;
	goto L214
L217:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	v749 = F_strstr(m, v713, v748)
	mBase = m.M
	if v749 != 0 {
		goto L215
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	v768 = v15 + int32(160)
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v15)+180))
	F_appendBinaryStringInfo(m, v768, v769, v770)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	F_appendStringInfoString(m, v768, v713)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v15)+176))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v15)+180))
	F_appendBinaryStringInfo(m, v768, v775, v776)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	goto L178
L223:
	;
	F_ReleaseCatCache(m, v25)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
	v799 = F_cstring_to_text(m, v798)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	F_pfree(m, v798)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	v803 = v799
	goto L5
L227:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v36
	F_errmsg(m, int32(_a_F_pg_get_functiondef_28), v15)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(_a_F_pg_get_functiondef_29), int32(2956), int32(_a_F_pg_get_functiondef_30))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L231:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_functiondef_31), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_pg_get_functiondef_29), int32(3107), int32(_a_F_pg_get_functiondef_30))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_indexdef(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_get_indexdef_worker(m, v3, v2, v2, v2, v2, v2, v2, int32(2), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			return int32(0)
		} else {
			v22 = F_cstring_to_text(m, v12)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v12)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					return v22
				}
			}
		}
	}
}
func F_pg_get_loaded_modules(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v5 = m.G0
	v7 = v5 - int32(32)
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
	return int32(0)
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
	m.G0 = v7 + int32(32)
	return int32(0)
L6:
	;
	v21 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+6)) = uint8(v21)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)) = uint16(v21)
	v31 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v7+int32(28)))) = v17 + v31
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v7+v31))) = v37
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v7+int32(20)))) = v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	if v44 == v21 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	if v52 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v47 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v47)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v49 = F_cstring_to_text(m, v44)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v49
	goto L8
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	v63 = v60
	v65 = int32(0)
	goto L20
L14:
	;
	v55 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+5)) = uint8(v55)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v57 = F_cstring_to_text(m, v52)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v57
	goto L13
L18:
	;
	v79 = F_cstring_to_text(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L30
	}
L19:
	;
	if v65 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v66 != int32(47) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v63 = v63 + int32(1)
	v65 = v69
	goto L20
L23:
	;
	if v66 != 0 {
		v69 = v65
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v69 = v63
	goto L22
L26:
	;
	goto L19
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	v78 = v74
	goto L18
L28:
	;
	goto L29
L29:
	;
	v76 = v65 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v76
	v78 = v76
	goto L18
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v79
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	F_tuplestore_putvalues(m, v82, v83, v7+int32(8), v7+int32(4))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v90 != 0 {
		v17 = v90
		goto L6
	} else {
		goto L32
	}
L32:
	;
	goto L7
}
func F_pg_get_partition_constraintdef(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
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
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_get_partition_qual_relid(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v106 = int32(0)
			m.G0 = v9 + int32(80)
			return v106
		} else {
			v21 = F_get_rel_name(m, v11)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if v21 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
						F_errmsg_internal(m, int32(_a_F_pg_get_partition_constraintdef_0), v9)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_get_partition_constraintdef_1), int32(_a_F_pg_get_partition_constraintdef_2), int32(_a_F_pg_get_partition_constraintdef_3))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v26 = F_palloc0(m, int32(80))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v29 = F_palloc0(m, int32(136))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = int32(1)
							v33 = int32(114)
							*(*uint8)(unsafe.Add(mBase, uint32(v29)+21)) = uint8(v33)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v11
							v36 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v36
							*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(101)
							v41 = F_makeAlias(m, v21, v36)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v41
								*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v41
								v45 = int32(256)
								*(*uint16)(unsafe.Add(mBase, uint32(v29)+124)) = uint16(v45)
								v47 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v29)+20)) = uint8(v47)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v29
								*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v29
								v54 = F_list_make1_impl(m, int32(1), v9+int32(20))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v56 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v56
									*(*int64)(unsafe.Add(mBase, uint32(v26)+12)) = int64(0)
									*(*int32)(unsafe.Add(mBase, uint32(v26))) = v54
									F_set_rtable_names(m, v26, v56, v56)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										F_set_simple_column_names(m, v26)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v26
											*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v26
											v72 = F_list_make1_impl(m, int32(1), v9+int32(16))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												v75 = v9 - int32(-64)
												F_initStringInfo(m, v75)
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return int32(0)
												} else {
													v78 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+56)) = uint8(v78)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v78
													*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v72
													*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = v78
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+59)) = uint8(v78)
													v89 = int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(v9)+57)) = uint16(v89)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v78
													*(*int64)(unsafe.Add(mBase, uint32(v9)+44)) = int64(2)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v75
													F_get_rule_expr(m, v12, v9+int32(24), v78)
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return int32(0)
													} else {
														v101 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
														v102 = F_cstring_to_text(m, v101)
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v101)
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return int32(0)
															} else {
																v106 = v102
																m.G0 = v9 + int32(80)
																return v106
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
func F_pg_get_statisticsobjdef_expressions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_SearchSysCache1(m, int32(64), v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L3
	} else {
		goto L38
	}
L2:
	;
	m.G0 = v11 + int32(80)
	return v175
L3:
	;
	return int32(0)
L4:
	;
	if v15 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
	v175 = v2
	goto L2
L6:
	;
	goto L7
L7:
	;
	v25 = F_heap_attisnull(m, v15, int32(9), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_ReleaseCatCache(m, v15)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
	v35 = F_SysCacheGetAttrNotNull(m, int32(64), v15, int32(9))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L13
	}
L12:
	;
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
	v175 = v2
	goto L2
L13:
	;
	v37 = F_text_to_cstring(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v39 = F_stringToNode(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	F_pfree(m, v37)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v43 = v31 + v32
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v45 = F_get_rel_name(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	if v45 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v51 = F_palloc0(m, int32(80))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v54 = F_palloc0(m, int32(136))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = int32(1)
	v58 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+21)) = uint8(v58)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v49
	v61 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(101)
	v66 = F_makeAlias(m, v45, v61)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v66
	v70 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+124)) = uint16(v70)
	v72 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+20)) = uint8(v72)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v54
	v79 = F_list_make1_impl(m, int32(1), v11+int32(20))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v81 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+20)) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v51)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v79
	F_set_rtable_names(m, v51, v81, v81)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	F_set_simple_column_names(m, v51)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v51
	v97 = F_list_make1_impl(m, int32(1), v11+int32(16))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v99 = int32(0)
	if v39 == v99 {
		v161 = v99
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_ReleaseCatCache(m, v15)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L3
	} else {
		goto L36
	}
L27:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v102 <= int32(0) {
		v161 = v99
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v106 = int32(0)
	v108 = v99
	goto L29
L29:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v106<<(uint(int32(2))%32))))
	v120 = v11 - int32(-64)
	F_initStringInfo(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L31
	}
L30:
	;
	v161 = v153
	goto L26
L31:
	;
	v123 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+56)) = uint8(v123)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v123
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v123
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+59)) = uint8(v123)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v123
	*(*int64)(unsafe.Add(mBase, uint32(v11)+44)) = int64(2)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v120
	v139 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+57)) = uint16(v139)
	F_get_rule_expr(m, v118, v11+int32(24), v123)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v147 = F_cstring_to_text(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_statisticsobjdef_expressions[0]))
	v153 = F_accumArrayResult(m, v108, v147, int32(0), int32(25), v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v156 = v106 + int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v156 < v157 {
		v106 = v156
		v108 = v153
		goto L29
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_statisticsobjdef_expressions[0]))
	v171 = F_makeArrayResult(m, v161, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v175 = v171
	goto L2
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v44
	F_errmsg_internal(m, int32(_a_F_pg_get_statisticsobjdef_expressions_0), v11)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_pg_get_statisticsobjdef_expressions_1), int32(_a_F_pg_get_statisticsobjdef_expressions_2), int32(_a_F_pg_get_statisticsobjdef_expressions_3))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
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
func F_pg_identify_object(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
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
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v434 int32
	_ = v434
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13)+100)) = v15
	v24 = F_get_call_result_type(m, l0, v2, v13+int32(72))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errfinish(m, int32(_a_F_pg_identify_object_0), int32(2777), int32(_a_F_pg_identify_object_1))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L7
	} else {
		goto L123
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L7
	} else {
		goto L120
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L7
	} else {
		goto L118
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L7
	} else {
		goto L115
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L7
	} else {
		goto L113
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L7
	} else {
		goto L111
	}
L7:
	;
	return int32(0)
L8:
	;
	if v24 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v32 = v2
	goto L12
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L7
	} else {
		goto L108
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32*int32(40))+uint32(_c_F_pg_identify_object[0])))
	if v15 != v42 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v42 == v15 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v45 = v32 + int32(1)
	if v45 != int32(37) {
		v32 = v45
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
	v52 = F_table_open(m, v15, int32(1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	v268 = v2
	v273 = v2
	goto L20
L20:
	;
	v276 = v13 + int32(100)
	v278 = F_getObjectTypeDescription(m, v276, int32(1))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L7
	} else {
		goto L89
	}
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_pg_identify_object[1]))
	if v55 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v106 = int32(0)
	v107 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+20)))
	v109 = F_get_catalog_object_by_oid_extended(m, v52, v107, v16, v106)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L41
	}
L23:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v56 == v15 {
		v98 = v55
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v59 = int32(0)
	goto L31
L26:
	;
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_identify_object[1])) = v94
	v98 = v94
	goto L22
L28:
	;
	v94 = v70 + int32(_a_F_pg_identify_object_2)
	goto L27
L29:
	;
	v94 = v70 + int32(_a_F_pg_identify_object_3)
	goto L27
L30:
	;
	v94 = v70 + int32(_a_F_pg_identify_object_4)
	goto L27
L31:
	;
	v70 = v59 * int32(40)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_pg_identify_object[0])))
	if v15 != v71 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v94 = v70 + int32(_a_F_pg_identify_object_5)
	goto L27
L33:
	;
	if v59 == int32(36) {
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
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_pg_identify_object[2])))
	if v77 == v15 {
		goto L28
	} else {
		goto L37
	}
L37:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_pg_identify_object[3])))
	if v79 == v15 {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_pg_identify_object[4])))
	if v81 == v15 {
		goto L30
	} else {
		goto L39
	}
L39:
	;
	v59 = v59 + int32(4)
	goto L31
L40:
	;
	F_relation_close(m, v52, int32(1))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L7
	} else {
		goto L88
	}
L41:
	;
	if v109 == int32(0) {
		v255 = v106
		v260 = v2
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_pg_identify_object[1]))
	if v115 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v166 = int32(*(*int16)(unsafe.Add(mBase, uint32(v158)+24)))
	if v166 != 0 {
		goto L63
	} else {
		goto L64
	}
L44:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v116 == v15 {
		v158 = v115
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v119 = int32(0)
	goto L52
L47:
	;
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_identify_object[1])) = v154
	v158 = v154
	goto L43
L49:
	;
	v154 = v130 + int32(_a_F_pg_identify_object_2)
	goto L48
L50:
	;
	v154 = v130 + int32(_a_F_pg_identify_object_3)
	goto L48
L51:
	;
	v154 = v130 + int32(_a_F_pg_identify_object_4)
	goto L48
L52:
	;
	v130 = v119 * int32(40)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+uint32(_c_F_pg_identify_object[0])))
	if v15 != v131 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v154 = v130 + int32(_a_F_pg_identify_object_5)
	goto L48
L54:
	;
	if v119 == int32(36) {
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
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v130)+uint32(_c_F_pg_identify_object[2])))
	if v137 == v15 {
		goto L49
	} else {
		goto L58
	}
L58:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v130)+uint32(_c_F_pg_identify_object[3])))
	if v139 == v15 {
		goto L50
	} else {
		goto L59
	}
L59:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v130)+uint32(_c_F_pg_identify_object[4])))
	if v141 == v15 {
		goto L51
	} else {
		goto L60
	}
L60:
	;
	v119 = v119 + int32(4)
	goto L52
L61:
	;
	v234 = int32(0)
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+36)))
	if v235 != int32(1) {
		v255 = v234
		v260 = v233
		goto L40
	} else {
		goto L83
	}
L62:
	;
	v187 = int32(0)
	goto L74
L63:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
	v170 = F_heap_getattr_2(m, v109, v166, v167, v13+int32(80))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L7
	} else {
		goto L66
	}
L64:
	;
	v179 = v158
	v180 = v2
	goto L65
L65:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	if v181 == v15 {
		v226 = v179
		v233 = v180
		goto L61
	} else {
		goto L69
	}
L66:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+80)))
	if v172 == int32(1) {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_pg_identify_object[1]))
	if v176 == int32(0) {
		v184 = v170
		goto L62
	} else {
		goto L68
	}
L68:
	;
	v179 = v176
	v180 = v170
	goto L65
L69:
	;
	v184 = v180
	goto L62
L70:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_identify_object[1])) = v222
	v226 = v222
	v233 = v184
	goto L61
L71:
	;
	v222 = v198 + int32(_a_F_pg_identify_object_2)
	goto L70
L72:
	;
	v222 = v198 + int32(_a_F_pg_identify_object_3)
	goto L70
L73:
	;
	v222 = v198 + int32(_a_F_pg_identify_object_4)
	goto L70
L74:
	;
	v198 = v187 * int32(40)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+uint32(_c_F_pg_identify_object[0])))
	if v15 != v199 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v222 = v198 + int32(_a_F_pg_identify_object_5)
	goto L70
L76:
	;
	if v187 == int32(36) {
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
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v198)+uint32(_c_F_pg_identify_object[2])))
	if v205 == v15 {
		goto L71
	} else {
		goto L80
	}
L80:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v198)+uint32(_c_F_pg_identify_object[3])))
	if v207 == v15 {
		goto L72
	} else {
		goto L81
	}
L81:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v198)+uint32(_c_F_pg_identify_object[4])))
	if v209 == v15 {
		goto L73
	} else {
		goto L82
	}
L82:
	;
	v187 = v187 + int32(4)
	goto L74
L83:
	;
	v238 = int32(*(*int16)(unsafe.Add(mBase, uint32(v226)+22)))
	if v238 == int32(0) {
		v255 = v234
		v260 = v233
		goto L40
	} else {
		goto L84
	}
L84:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
	v244 = F_heap_getattr_2(m, v109, v238, v241, v13+int32(80))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+80)))
	if v246 == int32(1) {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	v249 = F_quote_identifier(m, v244)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	v255 = v249
	v260 = v233
	goto L40
L88:
	;
	v268 = v255
	v273 = v260
	goto L20
L89:
	;
	v280 = F_cstring_to_text(m, v278)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	v282 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)) = uint8(v282)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v280
	v290 = F_getObjectIdentityParts(m, v276, v282, v282, int32(1))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L7
	} else {
		goto L91
	}
L91:
	;
	if base.B2i32(v273 == v282)|base.B2i32(v290 == int32(0)) != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v304 = int32(1)
	goto L94
L93:
	;
	v296 = F_get_namespace_name(m, v273)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L95
	}
L94:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+77)) = uint8(v304)
	v306 = int32(0)
	if base.B2i32(v268 == v306)|base.B2i32(v290 == v306) == v306 {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v298 = F_quote_identifier(m, v296)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	v300 = F_cstring_to_text(m, v298)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v300
	v304 = int32(0)
	goto L94
L98:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)) = uint8(v329)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v336 = F_heap_form_tuple(m, v331, v13+int32(80), v13+int32(76))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L7
	} else {
		goto L106
	}
L99:
	;
	v324 = F_cstring_to_text(m, v290)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L7
	} else {
		goto L105
	}
L100:
	;
	v313 = F_cstring_to_text(m, v268)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L7
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v318 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+78)) = uint8(v318)
	if v290 == int32(0) {
		v329 = v318
		goto L98
	} else {
		goto L104
	}
L103:
	;
	v315 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+78)) = uint8(v315)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v313
	goto L99
L104:
	;
	goto L99
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+92)) = v324
	v329 = int32(0)
	goto L98
L106:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v336)+16))
	v339 = F_HeapTupleHeaderGetDatum(m, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	m.G0 = v13 + int32(112)
	return v339
L108:
	;
	F_errmsg_internal(m, int32(_a_F_pg_identify_object_6), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(_a_F_pg_identify_object_0), int32(_a_F_pg_identify_object_7), int32(_a_F_pg_identify_object_8))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v15
	F_errmsg_internal(m, int32(_a_F_pg_identify_object_9), v13-int32(-64))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	goto L1
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v15
	F_errmsg_internal(m, int32(_a_F_pg_identify_object_9), v13+int32(48))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	goto L1
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v15
	F_errmsg_internal(m, int32(_a_F_pg_identify_object_10), v13+int32(32))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L7
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_pg_identify_object_0), int32(_a_F_pg_identify_object_11), int32(_a_F_pg_identify_object_8))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v15
	F_errmsg_internal(m, int32(_a_F_pg_identify_object_9), v13+int32(16))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	goto L1
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v15
	F_errmsg_internal(m, int32(_a_F_pg_identify_object_12), v13)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_pg_identify_object_0), int32(_a_F_pg_identify_object_13), int32(_a_F_pg_identify_object_8))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
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
func F_pg_indexes_size(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_try_relation_open(m, v4, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v12 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
			return int32(0)
		} else {
			v16 = F_calculate_indexes_size(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_relation_close(m, v6, int32(1))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = F_Int64GetDatum(m, v16)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						return v21
					}
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
func F_pg_input_is_valid(m *base.Module, l0 int32) int32 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_input_is_valid[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v18
			v21 = *(*int64)(unsafe.Add(mBase, _c_F_pg_input_is_valid[1]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v21
			v23 = F_pg_input_is_valid_common(m, l0, v10, v15, v7)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v23
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
func F_pg_is_other_temp_schema(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
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
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pg_is_other_temp_schema[0]))
	if v6 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v118
L2:
	;
	if v4 == v6 {
		v118 = v2
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
		v118 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	return int32(0)
L8:
	;
	if v11 == int32(0) {
		v118 = v2
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v17 = int32(_a_F_pg_is_other_temp_schema_0)
	goto L12
L10:
	;
	if v55-v56 != 0 {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	goto L13
L13:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v24 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v25 = v11
	v26 = v17
	v27 = int32(8)
	v28 = v24
	goto L18
L15:
	;
	v51 = v17
	v55 = int32(0)
	goto L16
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	goto L10
L17:
	;
	v51 = v46
	v55 = v48
	goto L16
L18:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if base.B2i32(v28 != v30)|base.B2i32(v30 == int32(0)) != 0 {
		v46 = v26
		v48 = v28
		goto L17
	} else {
		goto L20
	}
L19:
	;
	v46 = v40
	v48 = int32(0)
	goto L17
L20:
	;
	v36 = v27 - int32(1)
	if v36 == int32(0) {
		v46 = v26
		v48 = v28
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v39 = int32(1)
	v40 = v26 + v39
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v41 != 0 {
		v25 = v25 + v39
		v26 = v40
		v27 = v36
		v28 = v41
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v64 = int32(_a_F_pg_is_other_temp_schema_1)
	goto L28
L24:
	;
	v112 = int32(0)
	goto L25
L25:
	;
	F_pfree(m, v11)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L7
	} else {
		goto L39
	}
L26:
	;
	v112 = v102 - v103
	goto L25
L28:
	;
	goto L29
L29:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v71 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v72 = v11
	v73 = v64
	v74 = int32(14)
	v75 = v71
	goto L34
L31:
	;
	v98 = v64
	v102 = int32(0)
	goto L32
L32:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	goto L26
L33:
	;
	v98 = v93
	v102 = v95
	goto L32
L34:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if base.B2i32(v75 != v77)|base.B2i32(v77 == int32(0)) != 0 {
		v93 = v73
		v95 = v75
		goto L33
	} else {
		goto L36
	}
L35:
	;
	v93 = v87
	v95 = int32(0)
	goto L33
L36:
	;
	v83 = v74 - int32(1)
	if v83 == int32(0) {
		v93 = v73
		v95 = v75
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v86 = int32(1)
	v87 = v73 + v86
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	if v88 != 0 {
		v72 = v72 + v86
		v73 = v87
		v74 = v83
		v75 = v88
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	v118 = base.B2i32(v112 == int32(0))
	goto L1
}
func F_pg_isolation_test_session_is_blocked(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v158 int32
	_ = v158
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = F_pg_detoast_datum(m, v16)
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
	v22 = F_BackendPidGetProc(m, v15)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L51
	}
L4:
	;
	m.G0 = v13 + int32(16)
	return v187
L5:
	;
	if v22 == int32(0) {
		v187 = int32(0)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+548))
	if v26 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v72 = F_array_contains_nulls(m, v17)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L24
	}
L8:
	;
	if v41 == int32(0) {
		goto L7
	} else {
		goto L15
	}
L9:
	;
	v41 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v31 = v26 - int32(16777216)
	if base.Ui32(int32(184549375)) < base.Ui32(v31) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v41 = int32(_a_F_pg_isolation_test_session_is_blocked_0)
	goto L8
L13:
	;
	goto L14
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v31)>>(uint(int32(22))%32))&int32(1020))+uint32(_c_F_pg_isolation_test_session_is_blocked[0])))
	v41 = v39
	goto L8
L15:
	;
	v44 = int32(_a_F_pg_isolation_test_session_is_blocked_1)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_isolation_test_session_is_blocked[1])))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v47 == int32(0))|base.B2i32(v47 != v50) != 0 {
		v68 = v47
		v69 = v50
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v68-v69 != 0 {
		goto L7
	} else {
		goto L23
	}
L17:
	;
	goto L16
L18:
	;
	v53 = v44
	v54 = v41
	goto L19
L19:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	if v58 == int32(0) {
		v68 = v58
		v69 = v57
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v68 = v58
	v69 = v57
	goto L17
L21:
	;
	v61 = int32(1)
	if v58 == v57 {
		v53 = v53 + v61
		v54 = v54 + v61
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v187 = int32(1)
	goto L4
L24:
	;
	if v72 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v78 = F_ArrayGetNItemsSafe(m, v75, v17+int32(16))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v82 = F_DirectFunctionCall1Coll(m, int32(1545), int32(0), v15)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v84 = F_pg_detoast_datum(m, v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v90 = F_ArrayGetNItemsSafe(m, v87, v84+int32(16))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if int32(0) < v90 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v86 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v173 = F_GetSafeSnapshotBlockingPids(m, v15, v13+int32(12), int32(1))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L50
	}
L33:
	;
	v100 = v86
	goto L35
L34:
	;
	v100 = (v87<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L35
L35:
	;
	if v74 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v108 = v74
	goto L38
L37:
	;
	v108 = (v75<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L38
L38:
	;
	v110 = int32(0)
	v114 = v110
	goto L39
L39:
	;
	if v78 <= v110 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L32
L41:
	;
	v158 = v114 + int32(1)
	if v158 != v90 {
		v114 = v158
		goto L39
	} else {
		goto L49
	}
L42:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v84+v100+v114<<(uint(int32(2))%32))))
	v128 = int32(0)
	goto L43
L43:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v17+v108+v128<<(uint(int32(2))%32))))
	if v141 != v126 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v187 = int32(1)
	goto L4
L45:
	;
	v144 = v128 + int32(1)
	if v78 != v144 {
		v128 = v144
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
	v187 = base.B2i32(int32(0) < v173)
	goto L4
L51:
	;
	F_errmsg_internal(m, int32(_a_F_pg_isolation_test_session_is_blocked_2), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_pg_isolation_test_session_is_blocked_3), int32(66), int32(_a_F_pg_isolation_test_session_is_blocked_4))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v77 int64
	_ = v77
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = F_superuser_arg(m, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		if v13 == int32(0) {
			v21 = F_table_open(m, int32(2995), int32(1))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				v24 = v9 + int32(-48)
				F_ScanKeyInit(m, v24, int32(1), int32(3), int32(184), l0)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v31 = int32(1)
					v33 = F_systable_beginscan(m, v21, int32(2996), v31, l3, v31, v24)
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
								v89 = m.ExcPending
								if v89 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(67137668))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
										F_errmsg(m, int32(_a_F_pg_largeobject_aclmask_snapshot_0), v11)
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_pg_largeobject_aclmask_snapshot_1), int32(3573), int32(_a_F_pg_largeobject_aclmask_snapshot_2))
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
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
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
								v47 = F_heap_getattr_2(m, v35, int32(3), v44, v9+int32(-49))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
									if v49 == int32(1) {
										v54 = F_acldefault(m, int32(22), v42)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int64(0)
										} else {
											v58 = int32(0)
											v59 = v54
											v61 = F_aclmask(m, v59, l1, v42, l2, int32(1))
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int64(0)
											} else {
												v63 = int32(0)
												if base.B2i32(v59 == v63)|base.B2i32(v58 == v59) == v63 {
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
															F_relation_close(m, v21, int32(1))
															mBase = m.M
															v75 = m.ExcPending
															if v75 != 0 {
																return int64(0)
															} else {
																v77 = v61
																m.G0 = v11 - int32(-64)
																return v77
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
														F_relation_close(m, v21, int32(1))
														mBase = m.M
														v75 = m.ExcPending
														if v75 != 0 {
															return int64(0)
														} else {
															v77 = v61
															m.G0 = v11 - int32(-64)
															return v77
														}
													}
												}
											}
										}
									} else {
										v56 = F_pg_detoast_datum(m, v47)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int64(0)
										} else {
											v58 = v47
											v59 = v56
											v61 = F_aclmask(m, v59, l1, v42, l2, int32(1))
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int64(0)
											} else {
												v63 = int32(0)
												if base.B2i32(v59 == v63)|base.B2i32(v58 == v59) == v63 {
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
															F_relation_close(m, v21, int32(1))
															mBase = m.M
															v75 = m.ExcPending
															if v75 != 0 {
																return int64(0)
															} else {
																v77 = v61
																m.G0 = v11 - int32(-64)
																return v77
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
														F_relation_close(m, v21, int32(1))
														mBase = m.M
														v75 = m.ExcPending
														if v75 != 0 {
															return int64(0)
														} else {
															v77 = v61
															m.G0 = v11 - int32(-64)
															return v77
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
		} else {
			v77 = l2
			m.G0 = v11 - int32(-64)
			return v77
		}
	}
}
func F_pg_last_xact_replay_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = F_GetLatestXTime(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int64(0) {
			v9 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v9)
			return int32(0)
		} else {
			v13 = F_Int64GetDatum(m, v3)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				return v13
			}
		}
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
func F_pg_mbcliplen(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	v4 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mbcliplen[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if base.Ui32(v10) <= base.Ui32(int32(41)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v71
L2:
	;
	if v17 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10*int32(28))+uint32(_c_F_pg_mbcliplen[1])))
	v17 = v15
	goto L5
L4:
	;
	v17 = int32(1)
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
		v71 = v4
		goto L1
	} else {
		goto L17
	}
L9:
	;
	v21 = l1
	goto L11
L10:
	;
	v21 = l2
	goto L11
L11:
	;
	if v21 <= int32(0) {
		v71 = v4
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v27 = v4
	goto L13
L13:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v27))))
	if v32 == int32(0) {
		v71 = v27
		goto L1
	} else {
		goto L15
	}
L14:
	;
	return v21
L15:
	;
	v36 = v27 + int32(1)
	if v36 != v21 {
		v27 = v36
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v10*int32(28))+uint32(_c_F_pg_mbcliplen[2])))
	v46 = l0
	v47 = l1
	v49 = v4
	goto L18
L18:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v53 == int32(0) {
		v71 = v49
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v71 = v60
	goto L1
L20:
	;
	v56 = m.T0[v45].(func(*base.Module, int32) int32)(m, v46)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	v60 = v56 + v49
	if l2 < v60 {
		v71 = v49
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if l2 == v60 {
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
	v65 = v47 - v56
	if int32(0) < v65 {
		v46 = v46 + v56
		v47 = v65
		v49 = v60
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v13 = F_umask(m, int32(0))
	mBase = m.M
	v16 = F_umask(m, v13&int32(-193))
	mBase = m.M
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v23 = l0 + base.B2i32(v17 == int32(47))
	goto L2
L1:
	;
	v76 = F_umask(m, v13)
	mBase = m.M
	m.G0 = v10 + int32(96)
	return v75
L2:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v28 != int32(47) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v75 = v59 >> (uint(int32(31)) % 32)
	goto L1
L4:
	;
	goto L3
L5:
	;
	v23 = v23 + int32(1)
	goto L2
L6:
	;
	v42 = F___fstatat(m, int32(-100), l0, v10, int32(0))
	mBase = m.M
	goto L14
L7:
	;
	v37 = F_umask(m, v13)
	mBase = m.M
	v39 = int32(0)
	goto L6
L8:
	;
	if v28 != 0 {
		goto L5
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
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v36 != 0 {
		v39 = int32(1)
		goto L6
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
	goto L7
L13:
	;
	v66 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v66)
	goto L5
L14:
	;
	if v42 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v45&int32(_a_F_pg_mkdir_p_0) != int32(_a_F_pg_mkdir_p_1) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	if v39 != 0 {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	if v39 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if v39 != 0 {
		goto L13
	} else {
		goto L24
	}
L21:
	;
	v53 = int32(54)
	goto L23
L22:
	;
	v53 = int32(20)
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_mkdir_p[0])) = v53
	v75 = int32(-1)
	goto L1
L24:
	;
	v75 = int32(0)
	goto L1
L25:
	;
	v58 = int32(511)
	goto L27
L26:
	;
	v58 = l1
	goto L27
L27:
	;
	v59 = F_mkdir(m, l0, v58)
	mBase = m.M
	v60 = int32(0)
	if v39&base.B2i32(v60 <= v59) == v60 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	goto L13
}
func F_pg_mule_dsplen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v4 = int32(1)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32((v5+int32(127))&int32(255)) < base.Ui32(int32(13)) {
		v30 = v4
	} else {
		v13 = v5 & int32(-2)
		if v13 == int32(154) {
			v30 = v4
		} else {
			if base.Ui32((v5+int32(112))&int32(255)) < base.Ui32(int32(10)) {
				v30 = int32(2)
			} else {
				if v13&int32(255) == int32(156) {
					v29 = int32(2)
				} else {
					v29 = int32(1)
				}
				v30 = v29
			}
		}
	}
	return v30
}
func F_pg_mule_mblen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v25 int32
	_ = v25
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32((v2+int32(127))&int32(255)) < base.Ui32(int32(13)) {
		return int32(2)
	} else {
		if base.Ui32((v2+int32(112))&int32(255)) < base.Ui32(int32(12)) {
			return int32(3)
		} else {
			if v2&int32(254) == int32(156) {
				v25 = int32(4)
			} else {
				v25 = int32(1)
			}
			return v25
		}
	}
}
func F_pg_node_tree_send(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_json_send(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_pg_num_nonnulls(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
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
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v11 = F_count_nulls(m, l0, v5+int32(12), v5+int32(8))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
			v23 = int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			v23 = v20 - v21
		}
		m.G0 = v5 + int32(16)
		return v23
	}
}
func F_pg_operator_is_visible(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_OperatorIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
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
				return v20
			}
		} else {
			return v20
		}
	}
}
func F_pg_plan_queries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v13 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = v5
	v23 = v5
	goto L7
L5:
	;
	v98 = v5
	goto L6
L6:
	;
	return v98
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v22<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v29 == int32(6) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v98 = v85
	goto L6
L9:
	;
	v85 = F_lappend(m, v23, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L13
	} else {
		goto L26
	}
L10:
	;
	v33 = F_palloc0(m, int32(104))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_queries[0])))
	if v50 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	return int32(0)
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = int64(25769804106)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+26)) = uint8(v39)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v28)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v45
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v47
	v84 = v33
	goto L9
L15:
	;
	v53 = int32(_a_F_pg_plan_queries_0)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_queries[1])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_queries[2])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_queries[3])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_queries[4])) = int64(1)
	v63 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L18
L16:
	;
	goto L17
L17:
	;
	v66 = F_planner(m, v28, l1, l2, l3)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
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
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_queries[0])))
	if v69 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_ShowUsage(m, int32(_a_F_pg_plan_queries_2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L13
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_queries[5])))
	if v76 != int32(1) {
		v84 = v66
		goto L9
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_queries[6])))
	F_elog_node_display(m, int32(_a_F_pg_plan_queries_3), v66, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L13
	} else {
		goto L25
	}
L25:
	;
	v84 = v66
	goto L9
L26:
	;
	v88 = v22 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v88 < v89 {
		v22 = v88
		v23 = v85
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
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pg_re_throw[0]))
	if v3 != 0 {
		F_pgl_longjmp(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_pg_re_throw[1]))
		v9 = v7 * int32(100)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_pg_re_throw[2]))) = int32(22)
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_pg_re_throw[3]))
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_pg_re_throw[4]))) = uint8(base.B2i32(v17 < int32(23)))
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_pg_re_throw[5]))
		if v24 == int32(2) {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_re_throw[6])))
			if v28 == int32(1) {
				v36 = int32(1)
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_pg_re_throw[7]))
				v36 = base.B2i32(v33 <= int32(22))
			}
			v38 = v36
		} else {
			v38 = int32(0)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_pg_re_throw[8]))) = uint8(v38)
		*(*int32)(unsafe.Add(mBase, _c_F_pg_re_throw[9])) = int32(0)
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_pg_re_throw[10])))
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_pg_re_throw[11])))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_pg_re_throw[12])))
		F_errfinish(m, v45, v48, v51)
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return
		} else {
			v54 = m.G0
			v56 = v54 - int32(32)
			m.G0 = v56
			*(*int32)(unsafe.Add(mBase, uint32(v56)+12)) = int32(42)
			*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = int32(2048)
			*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = int32(_a_F_pg_re_throw_0)
			*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(_a_F_pg_re_throw_1)
			F_write_stderr(m, int32(_a_F_pg_re_throw_2), v56)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, _c_F_pg_re_throw[13]))
				v71 = F_fflush(m, v70)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
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
func F_pg_read_file_off_len_missing(m *base.Module, l0 int32) int32 {
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
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v15 = F_pg_read_file_common(m, v4, v9, v11, base.B2i32(v12 != int32(0)))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				v19 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
				return int32(0)
			} else {
				return v15
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
						F_errfinish(m, int32(_a_F_pg_relpages_impl_1), int32(462), int32(_a_F_pg_relpages_impl_2))
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
func F_pg_set_regex_collation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v2 = int32(0)
	if l0 == int32(950) {
		v30 = v2
		v31 = v2
		*(*int32)(unsafe.Add(mBase, _c_F_pg_set_regex_collation[0])) = v31
		*(*int32)(unsafe.Add(mBase, _c_F_pg_set_regex_collation[1])) = v30
		return
	} else {
		if l0 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				F_errcode(m, int32(34209924))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_pg_set_regex_collation_0), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errhint(m, int32(_a_F_pg_set_regex_collation_1), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_pg_set_regex_collation_2), int32(242), int32(_a_F_pg_set_regex_collation_3))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
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
		} else {
			v9 = F_pg_newlocale_from_collation(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
				if v11 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_pg_set_regex_collation_4), int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_pg_set_regex_collation_2), int32(262), int32(_a_F_pg_set_regex_collation_3))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
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
					v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+3)))
					if v15 != 0 {
						v30 = v2
						v31 = int32(0)
					} else {
						v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
						if v17 != int32(98) {
							v23 = *(*int32)(unsafe.Add(mBase, _c_F_pg_set_regex_collation[2]))
							v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
							if v24 == int32(6) {
								v27 = int32(2)
							} else {
								v27 = int32(3)
							}
							v28 = v27
						} else {
							v28 = int32(1)
						}
						v30 = v28
						v31 = v9
					}
					*(*int32)(unsafe.Add(mBase, _c_F_pg_set_regex_collation[0])) = v31
					*(*int32)(unsafe.Add(mBase, _c_F_pg_set_regex_collation[1])) = v30
					return
				}
			}
		}
	}
}
func F_pg_settings_get_flags(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_text_to_cstring(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return v110
L2:
	;
	return int32(0)
L3:
	;
	v18 = F_find_option(m, v11, int32(0), int32(1), int32(21))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v18 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
	v110 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	if v25&int32(32) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v31 = F_cstring_to_text(m, int32(_a_F_pg_settings_get_flags_0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	v36 = v25
	v37 = int32(0)
	v38 = v8
	goto L10
L10:
	;
	if v36&int32(8) != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v31
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v36 = v35
	v37 = int32(1)
	v38 = v8 | int32(4)
	goto L10
L12:
	;
	v42 = F_cstring_to_text(m, int32(_a_F_pg_settings_get_flags_1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L15
	}
L13:
	;
	v48 = v36
	v49 = v37
	goto L14
L14:
	;
	if v48&int32(16) != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v42
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v48 = v47
	v49 = v37 + int32(1)
	goto L14
L16:
	;
	v56 = F_cstring_to_text(m, int32(_a_F_pg_settings_get_flags_2))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L19
	}
L17:
	;
	v62 = v48
	v63 = v49
	goto L18
L18:
	;
	if v62&int32(4) != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8|v49<<(uint(int32(2))%32)))) = v56
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v62 = v61
	v63 = v49 + int32(1)
	goto L18
L20:
	;
	v70 = F_cstring_to_text(m, int32(_a_F_pg_settings_get_flags_3))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	v76 = v62
	v77 = v63
	goto L22
L22:
	;
	if v76&int32(128) != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+v63<<(uint(int32(2))%32)))) = v70
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v76 = v75
	v77 = v63 + int32(1)
	goto L22
L24:
	;
	v84 = F_cstring_to_text(m, int32(_a_F_pg_settings_get_flags_4))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L27
	}
L25:
	;
	v90 = v77
	v91 = v76
	goto L26
L26:
	;
	if v91&int32(_a_F_pg_settings_get_flags_5) != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+v77<<(uint(int32(2))%32)))) = v84
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v90 = v77 + int32(1)
	v91 = v89
	goto L26
L28:
	;
	v98 = F_cstring_to_text(m, int32(_a_F_pg_settings_get_flags_6))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L31
	}
L29:
	;
	v103 = v90
	goto L30
L30:
	;
	v105 = F_construct_array_builtin(m, v8, v103, int32(25))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+v90<<(uint(int32(2))%32)))) = v98
	v103 = v90 + int32(1)
	goto L30
L32:
	;
	v110 = v105
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
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn13965(m, l0, l1, l2, l3, l4, int32(3), int32(_a_F_pg_strfold_0), int32(1342))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_pg_strnxfrm(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = m.T0[v7].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_pg_tablespace_size_name(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_get_tablespace_oid(m, v3, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_calculate_tablespace_size(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			if v9 < int64(0) {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				return int32(0)
			} else {
				v17 = F_Int64GetDatum(m, v9)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
	}
}
func F_pg_terminate_backend(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int64
	_ = v104
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int64
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int64
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v15 = base.I32_wrap_i64(v14)
	if int32(0) <= v15 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_signal_backend(m, v18, int32(15))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
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
	v194 = m.ExcPending
	if v194 != 0 {
		goto L8
	} else {
		goto L52
	}
L4:
	;
	v97 = int32(0)
	if base.B2i32(v15 == v97)|v20 != 0 {
		v181 = base.B2i32(v20 == v97)
		goto L25
	} else {
		goto L26
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L20
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L8
	} else {
		goto L15
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return int32(0)
L9:
	;
	switch v20 - int32(2) {
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
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	F_errmsg(m, int32(_a_F_pg_terminate_backend_0), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v37 = int32(_a_F_pg_terminate_backend_1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v37
	F_errdetail(m, int32(_a_F_pg_terminate_backend_2), v11+int32(32))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_pg_terminate_backend_3), int32(258), int32(_a_F_pg_terminate_backend_4))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
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
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	F_errmsg(m, int32(_a_F_pg_terminate_backend_0), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(_a_F_pg_terminate_backend_5)
	F_errdetail(m, int32(_a_F_pg_terminate_backend_6), v11+int32(48))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_pg_terminate_backend_3), int32(265), int32(_a_F_pg_terminate_backend_4))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
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
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(_a_F_pg_terminate_backend_0), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = int32(_a_F_pg_terminate_backend_7)
	F_errdetail(m, int32(_a_F_pg_terminate_backend_8), v11-int32(-64))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_pg_terminate_backend_3), int32(272), int32(_a_F_pg_terminate_backend_4))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
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
	return v181
L26:
	;
	v104 = v14 & int64(2147483647)
	v110 = v104
	v111 = int64(100)
	goto L27
L27:
	;
	v114 = F_kill(m, v18, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L8
	} else {
		goto L29
	}
L28:
	;
	v159 = int32(0)
	v162 = F_errstart(m, int32(19), v159)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L48
	}
L29:
	;
	if v114 == int32(-1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_pg_terminate_backend[0]))
	if v120 == int32(71) {
		v181 = int32(1)
		goto L25
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v110 < v111 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v18
	F_errmsg(m, int32(_a_F_pg_terminate_backend_9), v11)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_pg_terminate_backend_3), int32(201), int32(_a_F_pg_terminate_backend_10))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v140 = v110
	goto L40
L39:
	;
	v140 = v111
	goto L40
L40:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_pg_terminate_backend[1]))
	if v142 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L8
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_pg_terminate_backend[2]))
	v150 = F_WaitLatch(m, v146, int32(41), base.I32_wrap_i64(v140), int32(134217731))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L8
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_pg_terminate_backend[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = int32(0)
	goto L46
L46:
	;
	v156 = v110 - v140
	if int64(0) < v156 {
		v110 = v156
		v111 = v140
		goto L27
	} else {
		goto L47
	}
L47:
	;
	goto L28
L48:
	;
	if v162 == int32(0) {
		v181 = v159
		goto L25
	} else {
		goto L49
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v18
	F_errmsg_plural(m, int32(_a_F_pg_terminate_backend_11), int32(_a_F_pg_terminate_backend_12), v15, v11+int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_pg_terminate_backend_3), int32(221), int32(_a_F_pg_terminate_backend_10))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v181 = v159
	goto L25
L52:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(_a_F_pg_terminate_backend_13), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_pg_terminate_backend_3), int32(249), int32(_a_F_pg_terminate_backend_4))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
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
func F_pg_truncate_visibility_map(m *base.Module, l0 int32) int32 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v102)+120)) = v103 | int32(3)
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
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+120)) = v191 & int32(-4)
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
	return int32(0)
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
	F_errfinish(m, int32(_a_F_pg_truncate_visibility_map_2), int32(951), int32(_a_F_pg_truncate_visibility_map_3))
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
func F_pg_type_is_visible(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_TypeIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	v2 = int32(-1)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	if l0 == int32(0) {
		v89 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if base.Ui32(int32(35)) <= base.Ui32(v89) {
		goto L27
	} else {
		goto L28
	}
L2:
	;
	m.G0 = v12 - int32(-64)
	goto L1
L3:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v17 == int32(0) {
		v89 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v20 = F_strlen(m, l0)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v20) {
		v89 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v23 = l0
	v24 = v17
	v25 = v12
	goto L6
L6:
	;
	v33 = F_isalnum(m, v24&int32(255))
	mBase = m.M
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v50)
	v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12))))
	v55 = int32(_a_F_pg_valid_server_encoding_private_0)
	v56 = int32(_a_F_pg_valid_server_encoding_private_1)
	goto L15
L8:
	;
	if base.Ui32((v24-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v46 = v25
	goto L10
L10:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v47 != 0 {
		v23 = v23 + int32(1)
		v24 = v47
		v25 = v46
		goto L6
	} else {
		goto L14
	}
L11:
	;
	v42 = v24 | int32(32)
	goto L13
L12:
	;
	v42 = v24
	goto L13
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v42)
	v46 = v25 + int32(1)
	goto L10
L14:
	;
	goto L7
L15:
	;
	v68 = v56 + (v55-v56)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = int32(*(*int8)(unsafe.Add(mBase, uint32(v69))))
	v71 = v54 - v70
	if v71 != 0 {
		v74 = v71
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v89 = v2
	goto L2
L17:
	;
	v78 = base.B2i32(v74 < int32(0))
	if v74 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v72 = F_strcmp(m, v12, v69)
	mBase = m.M
	if v72 != 0 {
		v74 = v72
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v89 = v73
	goto L2
L20:
	;
	v79 = v68 - int32(8)
	goto L22
L21:
	;
	v79 = v55
	goto L22
L22:
	;
	if v74 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v82 = v56
	goto L25
L24:
	;
	v82 = v68 + int32(8)
	goto L25
L25:
	;
	if base.Ui32(v82) <= base.Ui32(v79) {
		v55 = v79
		v56 = v82
		goto L15
	} else {
		goto L26
	}
L26:
	;
	goto L16
L27:
	;
	v97 = v2
	goto L29
L28:
	;
	v97 = v89
	goto L29
L29:
	;
	return v97
}
func F_pg_visibility(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v2
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)) = uint16(v2)
	v22 = F_relation_open(m, v14, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+119)))
		v29 = v27 - int32(109)
		v36 = int32(0)
		if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v29))|base.B2i32(int32(1)<<(uint(v29)%32)&int32(161) == v36) == v36 {
			if base.Ui64(int64(4294967295)) <= base.Ui64(v13) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v166 = m.ExcPending
				if v166 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v169 = m.ExcPending
					if v169 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_pg_visibility_0), int32(0))
						mBase = m.M
						v173 = m.ExcPending
						if v173 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_visibility_1), int32(144), int32(_a_F_pg_visibility_2))
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v44 = F_CreateTemplateTupleDesc(m, int32(3))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_TupleDescInitEntry(m, v44, int32(1), int32(_a_F_pg_visibility_3), int32(16), int32(-1), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_TupleDescInitEntry(m, v44, int32(2), int32(_a_F_pg_visibility_4), int32(16), int32(-1), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_TupleDescInitEntry(m, v44, int32(3), int32(_a_F_pg_visibility_5), int32(16), int32(-1), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v67 = F_BlessTupleDesc(m, v44)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v69 = base.I32_wrap_i64(v13)
									v72 = F_visibilitymap_get_status(m, v22, v69, v10+int32(28))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
										if v74 != 0 {
											F_ReleaseBuffer(m, v74)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												v77 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v72 & v77
												*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(base.Ui32(v72)>>(uint(v77)%32)) & v77
												v86 = F_RelationGetNumberOfBlocksInFork(m, v22, int32(0))
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int32(0)
												} else {
													if base.Ui64(v13) < base.Ui64(base.I64_extend_i32_u(v86)) {
														v90 = F_ReadBuffer(m, v22, v69)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int32(0)
														} else {
															F_LockBuffer(m, v90, int32(1))
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int32(0)
															} else {
																if v90 < int32(0) {
																	v98 = *(*int32)(unsafe.Add(mBase, _c_F_pg_visibility[0]))
																	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98+(v90^int32(-1))<<(uint(int32(2))%32))))
																	v112 = v104
																} else {
																	v106 = *(*int32)(unsafe.Add(mBase, _c_F_pg_visibility[1]))
																	v112 = v106 + v90<<(uint(int32(13))%32) + int32(-8192)
																}
																v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+10)))
																*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(base.Ui32(v113)>>(uint(int32(2))%32)) & int32(1)
																F_UnlockReleaseBuffer(m, v90)
																mBase = m.M
																v120 = m.ExcPending
																if v120 != 0 {
																	return int32(0)
																} else {
																	F_relation_close(m, v22, int32(1))
																	mBase = m.M
																	v126 = m.ExcPending
																	if v126 != 0 {
																		return int32(0)
																	} else {
																		v131 = F_heap_form_tuple(m, v67, v10+int32(16), v10+int32(12))
																		mBase = m.M
																		v132 = m.ExcPending
																		if v132 != 0 {
																			return int32(0)
																		} else {
																			v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
																			v134 = F_HeapTupleHeaderGetDatum(m, v133)
																			mBase = m.M
																			v135 = m.ExcPending
																			if v135 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v10 + int32(32)
																				return v134
																			}
																		}
																	}
																}
															}
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(0)
														F_relation_close(m, v22, int32(1))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return int32(0)
														} else {
															v131 = F_heap_form_tuple(m, v67, v10+int32(16), v10+int32(12))
															mBase = m.M
															v132 = m.ExcPending
															if v132 != 0 {
																return int32(0)
															} else {
																v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
																v134 = F_HeapTupleHeaderGetDatum(m, v133)
																mBase = m.M
																v135 = m.ExcPending
																if v135 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v10 + int32(32)
																	return v134
																}
															}
														}
													}
												}
											}
										} else {
											v77 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v72 & v77
											*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(base.Ui32(v72)>>(uint(v77)%32)) & v77
											v86 = F_RelationGetNumberOfBlocksInFork(m, v22, int32(0))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if base.Ui64(v13) < base.Ui64(base.I64_extend_i32_u(v86)) {
													v90 = F_ReadBuffer(m, v22, v69)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return int32(0)
													} else {
														F_LockBuffer(m, v90, int32(1))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int32(0)
														} else {
															if v90 < int32(0) {
																v98 = *(*int32)(unsafe.Add(mBase, _c_F_pg_visibility[0]))
																v104 = *(*int32)(unsafe.Add(mBase, uint32(v98+(v90^int32(-1))<<(uint(int32(2))%32))))
																v112 = v104
															} else {
																v106 = *(*int32)(unsafe.Add(mBase, _c_F_pg_visibility[1]))
																v112 = v106 + v90<<(uint(int32(13))%32) + int32(-8192)
															}
															v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+10)))
															*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(base.Ui32(v113)>>(uint(int32(2))%32)) & int32(1)
															F_UnlockReleaseBuffer(m, v90)
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v22, int32(1))
																mBase = m.M
																v126 = m.ExcPending
																if v126 != 0 {
																	return int32(0)
																} else {
																	v131 = F_heap_form_tuple(m, v67, v10+int32(16), v10+int32(12))
																	mBase = m.M
																	v132 = m.ExcPending
																	if v132 != 0 {
																		return int32(0)
																	} else {
																		v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
																		v134 = F_HeapTupleHeaderGetDatum(m, v133)
																		mBase = m.M
																		v135 = m.ExcPending
																		if v135 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v10 + int32(32)
																			return v134
																		}
																	}
																}
															}
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(0)
													F_relation_close(m, v22, int32(1))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														v131 = F_heap_form_tuple(m, v67, v10+int32(16), v10+int32(12))
														mBase = m.M
														v132 = m.ExcPending
														if v132 != 0 {
															return int32(0)
														} else {
															v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
															v134 = F_HeapTupleHeaderGetDatum(m, v133)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return int32(0)
															} else {
																m.G0 = v10 + int32(32)
																return v134
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v143 = m.ExcPending
			if v143 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v146 = m.ExcPending
				if v146 != 0 {
					return int32(0)
				} else {
					v147 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v147 + int32(4)
					F_errmsg(m, int32(_a_F_pg_visibility_6), v10)
					mBase = m.M
					v153 = m.ExcPending
					if v153 != 0 {
						return int32(0)
					} else {
						v154 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
						v155 = int32(*(*int8)(unsafe.Add(mBase, uint32(v154)+119)))
						F_errdetail_relkind_not_supported(m, v155)
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_visibility_1), int32(951), int32(_a_F_pg_visibility_7))
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
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
func F_pg_visibility_map_summary(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int64
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
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
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	v2 = int32(0)
	v7 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v2
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+14)) = uint16(v2)
	v19 = F_relation_open(m, v13, int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L32
	}
L2:
	;
	return int32(0)
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+119)))
	v26 = v24 - int32(109)
	v33 = int32(0)
	if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v26))|base.B2i32(int32(1)<<(uint(v26)%32)&int32(161) == v33) == v33 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v38 = int32(0)
	v40 = F_RelationGetNumberOfBlocksInFork(m, v19, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L2
	} else {
		goto L27
	}
L7:
	;
	F_relation_close(m, v19, int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L20
	}
L8:
	;
	if v40 == int32(0) {
		v84 = v7
		v85 = v7
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v47 = v38
	v50 = v7
	v51 = v7
	goto L10
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_pg_visibility_map_summary[0]))
	if v53 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v73 == int32(0) {
		v84 = v63
		v85 = v69
		goto L7
	} else {
		goto L18
	}
L12:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v58 = F_visibilitymap_get_status(m, v19, v47, v11+int32(28))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v60 = int32(1)
	v63 = v50 + base.I64_extend_i32_u(v58&v60)
	v69 = v51 + base.I64_extend_i32_u(int32(base.Ui32(v58)>>(uint(v60)%32))&v60)
	v71 = v47 + v60
	if v71 != v40 {
		v47 = v71
		v50 = v63
		v51 = v69
		goto L10
	} else {
		goto L17
	}
L17:
	;
	goto L11
L18:
	;
	F_ReleaseBuffer(m, v73)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v84 = v63
	v85 = v69
	goto L7
L20:
	;
	v92 = F_get_call_result_type(m, l0, int32(0), v11+int32(24))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	if v92 != int32(1) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v96 = F_Int64GetDatum(m, v84)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v96
	v99 = F_Int64GetDatum(m, v85)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v99
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v107 = F_heap_form_tuple(m, v102, v11+int32(16), v11+int32(14))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
	v110 = F_HeapTupleHeaderGetDatum(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	m.G0 = v11 + int32(32)
	return v110
L27:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v123 + int32(4)
	F_errmsg(m, int32(_a_F_pg_visibility_map_summary_0), v11)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v131 = int32(*(*int8)(unsafe.Add(mBase, uint32(v130)+119)))
	F_errdetail_relkind_not_supported(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_pg_visibility_map_summary_1), int32(951), int32(_a_F_pg_visibility_map_summary_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_pg_visibility_map_summary_3), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_pg_visibility_map_summary_1), int32(310), int32(_a_F_pg_visibility_map_summary_4))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_visibility_rel(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v12 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v16 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(_a_F_pg_visibility_rel_0)
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_pg_visibility_rel[0]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_pg_visibility_rel[0])) = v23
			v26 = F_CreateTemplateTupleDesc(m, int32(4))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_TupleDescInitEntry(m, v26, int32(1), int32(_a_F_pg_visibility_rel_1), int32(20), int32(-1), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_TupleDescInitEntry(m, v26, int32(2), int32(_a_F_pg_visibility_rel_2), int32(16), int32(-1), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_TupleDescInitEntry(m, v26, int32(3), int32(_a_F_pg_visibility_rel_3), int32(16), int32(-1), int32(0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							F_TupleDescInitEntry(m, v26, int32(4), int32(_a_F_pg_visibility_rel_4), int32(16), int32(-1), int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v56 = F_BlessTupleDesc(m, v26)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v56
									v60 = F_collect_visibility_data(m, v15, int32(1))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v60
										*(*int32)(unsafe.Add(mBase, _c_F_pg_visibility_rel[0])) = v21
										v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
										v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
										if base.Ui32(v72) < base.Ui32(v73) {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
											v78 = F_Int64GetDatum(m, base.I64_extend_i32_u(v72))
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v78
												v81 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
												v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v81)+8)))
												v84 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v83 & v84
												*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = int32(base.Ui32(v83)>>(uint(int32(2))%32)) & v84
												*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(base.Ui32(v83)>>(uint(v84)%32)) & v84
												*(*int32)(unsafe.Add(mBase, uint32(v71))) = v81 + v84
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
												v105 = F_heap_form_tuple(m, v100, v9+int32(16), v9+int32(12))
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													v107 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
													*(*int64)(unsafe.Add(mBase, uint32(v70))) = v107 + int64(1)
													v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v111)+20)) = int32(1)
													v114 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
													v115 = F_HeapTupleHeaderGetDatum(m, v114)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return int32(0)
													} else {
														v127 = v115
														m.G0 = v9 + int32(32)
														return v127
													}
												}
											}
										} else {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v119)+20)) = int32(2)
												v122 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v122)
												v127 = int32(0)
												m.G0 = v9 + int32(32)
												return v127
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
	} else {
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
		v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
		v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
		if base.Ui32(v72) < base.Ui32(v73) {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
			v78 = F_Int64GetDatum(m, base.I64_extend_i32_u(v72))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v78
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
				v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v81)+8)))
				v84 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v83 & v84
				*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = int32(base.Ui32(v83)>>(uint(int32(2))%32)) & v84
				*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(base.Ui32(v83)>>(uint(v84)%32)) & v84
				*(*int32)(unsafe.Add(mBase, uint32(v71))) = v81 + v84
				v100 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
				v105 = F_heap_form_tuple(m, v100, v9+int32(16), v9+int32(12))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					v107 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
					*(*int64)(unsafe.Add(mBase, uint32(v70))) = v107 + int64(1)
					v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v111)+20)) = int32(1)
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
					v115 = F_HeapTupleHeaderGetDatum(m, v114)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						v127 = v115
						m.G0 = v9 + int32(32)
						return v127
					}
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v118 = m.ExcPending
			if v118 != 0 {
				return int32(0)
			} else {
				v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v119)+20)) = int32(2)
				v122 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v122)
				v127 = int32(0)
				m.G0 = v9 + int32(32)
				return v127
			}
		}
	}
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
func F_pg_xact_commit_timestamp_origin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_TransactionIdGetCommitTsData(m, v8, v6+int32(16), v6+int32(30))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = F_get_call_result_type(m, l0, int32(0), v6)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v18 == int32(1) {
				if v13 == int32(0) {
					v24 = int32(257)
					*(*uint16)(unsafe.Add(mBase, uint32(v6)+6)) = uint16(v24)
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v39 = F_heap_form_tuple(m, v34, v6+int32(8), v6+int32(6))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
						v42 = F_HeapTupleHeaderGetDatum(m, v41)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							m.G0 = v6 + int32(32)
							return v42
						}
					}
				} else {
					v26 = *(*int64)(unsafe.Add(mBase, uint32(v6)+16))
					v27 = F_Int64GetDatum(m, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v27
						v30 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v6)+6)) = uint16(v30)
						v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+30)))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v32
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						v39 = F_heap_form_tuple(m, v34, v6+int32(8), v6+int32(6))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
							v42 = F_HeapTupleHeaderGetDatum(m, v41)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								m.G0 = v6 + int32(32)
								return v42
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_pg_xact_commit_timestamp_origin_0), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_xact_commit_timestamp_origin_1), int32(478), int32(_a_F_pg_xact_commit_timestamp_origin_2))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
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
