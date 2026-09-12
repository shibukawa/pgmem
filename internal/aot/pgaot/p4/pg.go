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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
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
	return v84
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
	v22 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	switch v22 - int32(2) {
	case 0:
		goto L8
	default:
		goto L7
	case 22, 26:
		v84 = v11
		goto L1
	}
L5:
	;
	goto L6
L6:
	;
	v27 = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, _consts[644]))
	v34 = F___fstatat(m, int32(-100), v29, v7+int32(8), v27)
	mBase = m.M
	goto L9
L7:
	;
	v84 = int32(0)
	goto L1
L8:
	;
	v84 = int32(3)
	goto L1
L9:
	;
	if v34 < int32(0) {
		v84 = v27
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[645]))
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v58 == int32(-1) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	v42 = v40
	goto L15
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(28)
	v58 = int32(-1)
	goto L11
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if l0 == v44 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v58 = v46
	goto L11
L18:
	;
	goto L19
L19:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	if v47 != 0 {
		v42 = v47
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	v61 = int32(2)
	v63 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	switch v63 - v61 {
	case 0:
		goto L25
	default:
		goto L24
	case 22, 26:
		v84 = v61
		goto L1
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v58
	v69 = int32(3)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v70 != int32(679834894) {
		v84 = v69
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v84 = int32(0)
	goto L1
L25:
	;
	v84 = int32(3)
	goto L1
L26:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v58)+24))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v73 != v74 {
		v84 = v69
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v58)+32))
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v7)+96))
	if v76 != v77 {
		v84 = v69
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+176))
	if v81 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v82 = int32(1)
	goto L31
L30:
	;
	v82 = int32(4)
	goto L31
L31:
	;
	v84 = v82
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
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v4 = *(*int32)(unsafe.Add(mBase, _consts[717]))
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
	v8 = *(*int32)(unsafe.Add(mBase, _consts[718]))
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
	v12 = *(*int32)(unsafe.Add(mBase, _consts[716]))
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
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L36
	}
L11:
	;
	return
L12:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[777]))
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
	*(*int32)(unsafe.Add(mBase, _consts[716])) = int32(0)
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
	v26 = *(*int32)(unsafe.Add(mBase, _consts[777]))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v27 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[778]))
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
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v34 == int32(0) {
		v53 = v33
		v54 = v34
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
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	if v33 != v34 {
		v53 = v33
		v54 = v34
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v38 = v26
	v39 = v17
	goto L23
L23:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v43 == int32(0) {
		v53 = v42
		v54 = v43
		goto L20
	} else {
		goto L25
	}
L24:
	;
	v53 = v42
	v54 = v43
	goto L20
L25:
	;
	v46 = int32(1)
	if v42 == v43 {
		v38 = v38 + v46
		v39 = v39 + v46
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	if v54-v53 == int32(0) {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	v62 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	if v62 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_errmsg(m, int32(460884), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	F_errfinish(m, int32(499468), int32(900), int32(118556))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(107687), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	F_errdetail(m, int32(583939), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(499468), int32(882), int32(118556))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
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
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v266 int32
	_ = v266
	var v281 int32
	_ = v281
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v347 int32
	_ = v347
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v372 int32
	_ = v372
	v11 = m.G0
	v13 = v11 - int32(4144)
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
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L90
	}
L7:
	;
	v31 = v21 + int32(19)
	v32 = int32(553500)
	v36 = m.G0
	v38 = v36 - int32(32)
	v39 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = v39
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[957])))
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
		goto L30
	} else {
		goto L31
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
	v51 = int32(*(*uint8)(unsafe.Add(mBase, _consts[958])))
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
		v107 = v31
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v115 = v107 - v31
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
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v107 = v103
	goto L22
L26:
	;
	v107 = v86
	goto L22
L27:
	;
	goto L28
L28:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	v103 = v86 + int32(1)
	if v101 != 0 {
		v86 = v103
		v87 = v101
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l0
	v126 = F_pg_snprintf(m, v13+int32(48), int32(2048), int32(177577), v13+int32(32))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v359 = F_ReadDirExtended(m, v15, l0, int32(15))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L88
	}
L33:
	;
	v130 = F_AllocateDir(m, v13+int32(48))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v135 = F_ReadDirExtended(m, v130, v13+int32(48), int32(15))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v135 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v138 = v135
	goto L39
L37:
	;
	goto L38
L38:
	;
	F_FreeDir(m, v130)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L87
	}
L39:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+19)))
	if v147 != int32(116) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L38
L41:
	;
	v334 = F_ReadDirExtended(m, v130, v13+int32(48), int32(15))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L85
	}
L42:
	;
	v151 = v138 + int32(19)
	v154 = int32(1)
	goto L43
L43:
	;
	v164 = v154 + int32(1)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+v151))))
	if base.Ui32((v166-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v154 = v164
		goto L43
	} else {
		goto L45
	}
L44:
	;
	if v154 == int32(1) {
		goto L41
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	if v166&int32(255) != int32(95) {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v182 = v164
	goto L48
L48:
	;
	v190 = v182 + int32(1)
	v191 = v182 + v151
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if base.Ui32((v192-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v182 = v190
		goto L48
	} else {
		goto L50
	}
L49:
	;
	if v164 == v182 {
		goto L41
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	if v192 == int32(95) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v203 = v191 + int32(1)
	v207 = int32(3)
	v210 = F_strncmp(m, int32(288482), v203, v207)
	mBase = m.M
	if v210 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L53:
	;
	v247 = v182
	v248 = v192
	goto L54
L54:
	;
	if v248 == int32(46) {
		goto L70
	} else {
		goto L71
	}
L55:
	;
	if v239 <= int32(0) {
		goto L41
	} else {
		goto L69
	}
L56:
	;
	goto L55
L58:
	;
	v239 = v232
	goto L56
L59:
	;
	v232 = v207
	goto L58
L60:
	;
	goto L61
L61:
	;
	v214 = int32(2)
	v218 = F_strncmp(m, int32(286496), v203, v214)
	mBase = m.M
	if v218 == int32(0) {
		v232 = v214
		goto L58
	} else {
		goto L62
	}
L62:
	;
	v221 = int32(4)
	v224 = F_strncmp(m, int32(100816), v203, v221)
	mBase = m.M
	if v224 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	goto L66
L64:
	;
	goto L65
L65:
	;
	v239 = int32(0)
	goto L56
L66:
	;
	v239 = v221
	goto L56
L69:
	;
	v243 = v239 + v190
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v243))))
	v247 = v243
	v248 = v245
	goto L54
L70:
	;
	v254 = int32(1)
	goto L73
L71:
	;
	v281 = v248
	goto L72
L72:
	;
	if v281 != 0 {
		goto L41
	} else {
		goto L77
	}
L73:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254+(v247+v151)))))
	if base.Ui32((v266-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v254 = v254 + int32(1)
		goto L73
	} else {
		goto L75
	}
L74:
	;
	if v254 < int32(2) {
		goto L41
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	v281 = v266
	goto L72
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v13 + int32(48)
	v295 = F_pg_snprintf(m, v13+int32(2096), int32(2048), int32(177577), v13+int32(16))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v299 = F_unlink(m, v13+int32(2096))
	mBase = m.M
	if int32(0) <= v299 {
		goto L41
	} else {
		goto L79
	}
L79:
	;
	v304 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v304 == int32(0) {
		goto L41
	} else {
		goto L81
	}
L81:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v13 + int32(2096)
	F_errmsg(m, int32(300038), v13)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(501242), int32(3511), int32(420621))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	goto L41
L85:
	;
	if v334 != 0 {
		v138 = v334
		goto L39
	} else {
		goto L86
	}
L86:
	;
	goto L40
L87:
	;
	goto L32
L88:
	;
	if v359 != 0 {
		v21 = v359
		goto L7
	} else {
		goto L89
	}
L89:
	;
	goto L8
L90:
	;
	m.G0 = v13 + int32(4144)
	return
}
func F__PG_init_isn(m *base.Module) {
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	F_DefineCustomBoolVariable(m, int32(320534), int32(588595), int32(4685140), int32(0))
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		F_MarkGUCPrefixReserved(m, int32(245501))
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			return
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
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v417 int64
	_ = v417
	var v419 int64
	_ = v419
	var v421 int64
	_ = v421
	var v432 int32
	_ = v432
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v596 int64
	_ = v596
	var v598 int64
	_ = v598
	var v600 int64
	_ = v600
	var v611 int32
	_ = v611
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
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
	v659 = F_MemoryContextAllocZero(m, l1, int32(20))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L9
	} else {
		goto L177
	}
L2:
	;
	F_report_newlocale_failure(m, v48)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L9
	} else {
		goto L176
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L9
	} else {
		goto L173
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L9
	} else {
		goto L170
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
	v15 = *(*int32)(unsafe.Add(mBase, _consts[100]))
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
	v47 = v16
	v48 = v26
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
	v47 = v33
	v48 = v41
	v49 = v45
	goto L5
L20:
	;
	F_ReleaseCatCache(m, v47)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v57 == int32(0) {
		v76 = v56
		v77 = v57
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v77-v76 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	goto L22
L24:
	;
	if v56 != v57 {
		v76 = v56
		v77 = v57
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v61 = v48
	v62 = v50
	goto L26
L26:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v66 == int32(0) {
		v76 = v65
		v77 = v66
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v76 = v65
	v77 = v66
	goto L23
L28:
	;
	v69 = int32(1)
	if v65 == v66 {
		v61 = v61 + v69
		v62 = v62 + v69
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v81 == int32(67) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v260 != int32(67) {
		goto L76
	} else {
		goto L77
	}
L33:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v84 == int32(0) {
		v657 = v3
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v87 = int32(511453)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1317])))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v91 == int32(0) {
		v110 = v90
		v111 = v91
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L35
L37:
	;
	if v111-v110 == int32(0) {
		v657 = v3
		goto L1
	} else {
		goto L45
	}
L38:
	;
	goto L37
L39:
	;
	if v90 != v91 {
		v110 = v90
		v111 = v91
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v95 = v50
	v96 = v87
	goto L41
L41:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	if v100 == int32(0) {
		v110 = v99
		v111 = v100
		goto L38
	} else {
		goto L43
	}
L42:
	;
	v110 = v99
	v111 = v100
	goto L38
L43:
	;
	v103 = int32(1)
	if v99 == v100 {
		v95 = v95 + v103
		v96 = v96 + v103
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v116 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v116
	v123 = m.G0
	v125 = v123 - int32(32)
	m.G0 = v125
	v130 = v116
	goto L49
L46:
	;
	if v253 != 0 {
		v657 = v253
		goto L1
	} else {
		goto L74
	}
L47:
	;
	m.G0 = v125 + int32(32)
	goto L46
L48:
	;
	v253 = int32(0)
	goto L47
L49:
	;
	goto L52
L50:
	;
	v158 = F___loc_is_allocated(m, v116)
	mBase = m.M
	if v158 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125+int32(8)+v130<<(uint(int32(2))%32)))) = v149
	if v149 == int32(-1) {
		goto L48
	} else {
		goto L58
	}
L52:
	;
	if int32(1)<<(uint(v130)%32)&int32(9) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v148 = v48
	goto L57
L56:
	;
	v148 = int32(759461)
	goto L57
L57:
	;
	v149 = F___get_locale(m, v130, v148)
	mBase = m.M
	goto L51
L58:
	;
	v155 = v130 + int32(1)
	if v155 != int32(6) {
		v130 = v155
		goto L49
	} else {
		goto L59
	}
L59:
	;
	goto L50
L60:
	;
	v161 = int32(4101880)
	v166 = F_memcmp(m, v125+int32(8), v161, int32(24))
	mBase = m.M
	if v166 == int32(0) {
		v253 = v161
		goto L47
	} else {
		goto L63
	}
L61:
	;
	v234 = v116
	goto L62
L62:
	;
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v125)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v234))) = v238
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v125)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v234)+16)) = v240
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v125)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v234)+8)) = v242
	v253 = v234
	goto L47
L63:
	;
	v169 = int32(4101904)
	v174 = F_memcmp(m, v125+int32(8), v169, int32(24))
	mBase = m.M
	if v174 == int32(0) {
		v253 = v169
		goto L47
	} else {
		goto L64
	}
L64:
	;
	v177 = int32(0)
	v179 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1318])))
	if v179 == v177 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v185 = v177
	goto L68
L66:
	;
	goto L67
L67:
	;
	v212 = int32(4685380)
	v217 = F_memcmp(m, v125+int32(8), v212, int32(24))
	mBase = m.M
	if v217 == int32(0) {
		v253 = v212
		goto L47
	} else {
		goto L71
	}
L68:
	;
	v193 = F___get_locale(m, v185, int32(759461))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v185<<(uint(int32(2))%32))+uint32(_consts[1319]))) = v193
	v196 = v185 + int32(1)
	if v196 != int32(6) {
		v185 = v196
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v200 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1318])) = uint8(v200)
	v204 = *(*int32)(unsafe.Add(mBase, _consts[1319]))
	*(*int32)(unsafe.Add(mBase, _consts[1320])) = v204
	goto L67
L70:
	;
	goto L69
L71:
	;
	v220 = int32(4685404)
	v225 = F_memcmp(m, v125+int32(8), v220, int32(24))
	mBase = m.M
	if v225 == int32(0) {
		v253 = v220
		goto L47
	} else {
		goto L72
	}
L72:
	;
	v229 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v229 == int32(0) {
		goto L48
	} else {
		goto L73
	}
L73:
	;
	v234 = v229
	goto L62
L74:
	;
	goto L2
L75:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v442 != int32(67) {
		goto L119
	} else {
		goto L120
	}
L76:
	;
	v265 = int32(511453)
	v268 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1317])))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v269 == int32(0) {
		v288 = v268
		v289 = v269
		goto L80
	} else {
		goto L81
	}
L77:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	if v263 != 0 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v441 = int32(0)
	goto L75
L79:
	;
	if v289-v288 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L80:
	;
	goto L79
L81:
	;
	if v268 != v269 {
		v288 = v268
		v289 = v269
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v273 = v48
	v274 = v265
	goto L83
L83:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+1)))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+1)))
	if v278 == int32(0) {
		v288 = v277
		v289 = v278
		goto L80
	} else {
		goto L85
	}
L84:
	;
	v288 = v277
	v289 = v278
	goto L80
L85:
	;
	v281 = int32(1)
	if v277 == v278 {
		v273 = v273 + v281
		v274 = v274 + v281
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v441 = int32(0)
	goto L75
L88:
	;
	goto L89
L89:
	;
	v295 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v295
	v302 = m.G0
	v304 = v302 - int32(32)
	m.G0 = v304
	v309 = v295
	goto L93
L90:
	;
	if v432 == int32(0) {
		goto L2
	} else {
		goto L118
	}
L91:
	;
	m.G0 = v304 + int32(32)
	goto L90
L92:
	;
	v432 = int32(0)
	goto L91
L93:
	;
	goto L96
L94:
	;
	v337 = F___loc_is_allocated(m, v295)
	mBase = m.M
	if v337 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v304+int32(8)+v309<<(uint(int32(2))%32)))) = v328
	if v328 == int32(-1) {
		goto L92
	} else {
		goto L102
	}
L96:
	;
	if int32(1)<<(uint(v309)%32)&int32(8) != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v327 = v48
	goto L101
L100:
	;
	v327 = int32(759461)
	goto L101
L101:
	;
	v328 = F___get_locale(m, v309, v327)
	mBase = m.M
	goto L95
L102:
	;
	v334 = v309 + int32(1)
	if v334 != int32(6) {
		v309 = v334
		goto L93
	} else {
		goto L103
	}
L103:
	;
	goto L94
L104:
	;
	v340 = int32(4101880)
	v345 = F_memcmp(m, v304+int32(8), v340, int32(24))
	mBase = m.M
	if v345 == int32(0) {
		v432 = v340
		goto L91
	} else {
		goto L107
	}
L105:
	;
	v413 = v295
	goto L106
L106:
	;
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v304)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v413))) = v417
	v419 = *(*int64)(unsafe.Add(mBase, uint32(v304)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v413)+16)) = v419
	v421 = *(*int64)(unsafe.Add(mBase, uint32(v304)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v413)+8)) = v421
	v432 = v413
	goto L91
L107:
	;
	v348 = int32(4101904)
	v353 = F_memcmp(m, v304+int32(8), v348, int32(24))
	mBase = m.M
	if v353 == int32(0) {
		v432 = v348
		goto L91
	} else {
		goto L108
	}
L108:
	;
	v356 = int32(0)
	v358 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1318])))
	if v358 == v356 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v364 = v356
	goto L112
L110:
	;
	goto L111
L111:
	;
	v391 = int32(4685380)
	v396 = F_memcmp(m, v304+int32(8), v391, int32(24))
	mBase = m.M
	if v396 == int32(0) {
		v432 = v391
		goto L91
	} else {
		goto L115
	}
L112:
	;
	v372 = F___get_locale(m, v364, int32(759461))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v364<<(uint(int32(2))%32))+uint32(_consts[1319]))) = v372
	v375 = v364 + int32(1)
	if v375 != int32(6) {
		v364 = v375
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v379 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1318])) = uint8(v379)
	v383 = *(*int32)(unsafe.Add(mBase, _consts[1319]))
	*(*int32)(unsafe.Add(mBase, _consts[1320])) = v383
	goto L111
L114:
	;
	goto L113
L115:
	;
	v399 = int32(4685404)
	v404 = F_memcmp(m, v304+int32(8), v399, int32(24))
	mBase = m.M
	if v404 == int32(0) {
		v432 = v399
		goto L91
	} else {
		goto L116
	}
L116:
	;
	v408 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v408 == int32(0) {
		goto L92
	} else {
		goto L117
	}
L117:
	;
	v413 = v408
	goto L106
L118:
	;
	v441 = v432
	goto L75
L119:
	;
	v446 = int32(511453)
	v449 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1317])))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v450 == int32(0) {
		v469 = v449
		v470 = v450
		goto L123
	} else {
		goto L124
	}
L120:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v445 != 0 {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v657 = v441
	goto L1
L122:
	;
	if v470-v469 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L123:
	;
	goto L122
L124:
	;
	if v449 != v450 {
		v469 = v449
		v470 = v450
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v454 = v50
	v455 = v446
	goto L126
L126:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+1)))
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454)+1)))
	if v459 == int32(0) {
		v469 = v458
		v470 = v459
		goto L123
	} else {
		goto L128
	}
L127:
	;
	v469 = v458
	v470 = v459
	goto L123
L128:
	;
	v462 = int32(1)
	if v458 == v459 {
		v454 = v454 + v462
		v455 = v455 + v462
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	v657 = v441
	goto L1
L131:
	;
	goto L132
L132:
	;
	v475 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v475
	v481 = m.G0
	v483 = v481 - int32(32)
	m.G0 = v483
	v488 = v475
	goto L136
L133:
	;
	if v611 != 0 {
		v657 = v611
		goto L1
	} else {
		goto L161
	}
L134:
	;
	m.G0 = v483 + int32(32)
	goto L133
L135:
	;
	v611 = int32(0)
	goto L134
L136:
	;
	v493 = int32(1) << (uint(v488) % 32) & int32(1)
	if v441 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	v516 = F___loc_is_allocated(m, v441)
	mBase = m.M
	if v516 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v483+int32(8)+v488<<(uint(int32(2))%32)))) = v508
	if v508 == int32(-1) {
		goto L135
	} else {
		goto L145
	}
L139:
	;
	if v493 != 0 {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	if v493 != 0 {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v441+v488<<(uint(int32(2))%32))))
	v508 = v504
	goto L138
L142:
	;
	v506 = v50
	goto L144
L143:
	;
	v506 = int32(759461)
	goto L144
L144:
	;
	v507 = F___get_locale(m, v488, v506)
	mBase = m.M
	v508 = v507
	goto L138
L145:
	;
	v513 = v488 + int32(1)
	if v513 != int32(6) {
		v488 = v513
		goto L136
	} else {
		goto L146
	}
L146:
	;
	goto L137
L147:
	;
	v519 = int32(4101880)
	v524 = F_memcmp(m, v483+int32(8), v519, int32(24))
	mBase = m.M
	if v524 == int32(0) {
		v611 = v519
		goto L134
	} else {
		goto L150
	}
L148:
	;
	v592 = v441
	goto L149
L149:
	;
	v596 = *(*int64)(unsafe.Add(mBase, uint32(v483)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v592))) = v596
	v598 = *(*int64)(unsafe.Add(mBase, uint32(v483)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v592)+16)) = v598
	v600 = *(*int64)(unsafe.Add(mBase, uint32(v483)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v592)+8)) = v600
	v611 = v592
	goto L134
L150:
	;
	v527 = int32(4101904)
	v532 = F_memcmp(m, v483+int32(8), v527, int32(24))
	mBase = m.M
	if v532 == int32(0) {
		v611 = v527
		goto L134
	} else {
		goto L151
	}
L151:
	;
	v535 = int32(0)
	v537 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1318])))
	if v537 == v535 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v543 = v535
	goto L155
L153:
	;
	goto L154
L154:
	;
	v570 = int32(4685380)
	v575 = F_memcmp(m, v483+int32(8), v570, int32(24))
	mBase = m.M
	if v575 == int32(0) {
		v611 = v570
		goto L134
	} else {
		goto L158
	}
L155:
	;
	v551 = F___get_locale(m, v543, int32(759461))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v543<<(uint(int32(2))%32))+uint32(_consts[1319]))) = v551
	v554 = v543 + int32(1)
	if v554 != int32(6) {
		v543 = v554
		goto L155
	} else {
		goto L157
	}
L156:
	;
	v558 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1318])) = uint8(v558)
	v562 = *(*int32)(unsafe.Add(mBase, _consts[1319]))
	*(*int32)(unsafe.Add(mBase, _consts[1320])) = v562
	goto L154
L157:
	;
	goto L156
L158:
	;
	v578 = int32(4685404)
	v583 = F_memcmp(m, v483+int32(8), v578, int32(24))
	mBase = m.M
	if v583 == int32(0) {
		v611 = v578
		goto L134
	} else {
		goto L159
	}
L159:
	;
	v587 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v587 == int32(0) {
		goto L135
	} else {
		goto L160
	}
L160:
	;
	v592 = v587
	goto L149
L161:
	;
	if v441 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v618 = F___loc_is_allocated(m, v441)
	mBase = m.M
	if v618 != 0 {
		goto L166
	} else {
		goto L167
	}
L163:
	;
	goto L164
L164:
	;
	F_report_newlocale_failure(m, v50)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L9
	} else {
		goto L169
	}
L165:
	;
	goto L164
L166:
	;
	F_emscripten_builtin_free(m, v441)
	mBase = m.M
	goto L168
L167:
	;
	goto L168
L168:
	;
	goto L165
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	v627 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v627
	F_errmsg_internal(m, int32(49766), v9)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L9
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(501558), int32(435), int32(493475))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L9
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
	F_errmsg_internal(m, int32(46256), v9+int32(16))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L9
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(501558), int32(452), int32(493475))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L9
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	v661 = int32(355)
	*(*uint16)(unsafe.Add(mBase, uint32(v659))) = uint16(v661)
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v663 != int32(67) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v659)+2)) = uint8(v696)
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v698 != int32(67) {
		goto L191
	} else {
		goto L192
	}
L179:
	;
	v668 = int32(511453)
	v671 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1317])))
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v672 == int32(0) {
		v691 = v671
		v692 = v672
		goto L183
	} else {
		goto L184
	}
L180:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	if v666 != 0 {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v696 = int32(1)
	goto L178
L182:
	;
	v696 = base.B2i32(v692-v691 == int32(0))
	goto L178
L183:
	;
	goto L182
L184:
	;
	if v671 != v672 {
		v691 = v671
		v692 = v672
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v676 = v48
	v677 = v668
	goto L186
L186:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+1)))
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676)+1)))
	if v681 == int32(0) {
		v691 = v680
		v692 = v681
		goto L183
	} else {
		goto L188
	}
L187:
	;
	v691 = v680
	v692 = v681
	goto L183
L188:
	;
	v684 = int32(1)
	if v680 == v681 {
		v676 = v676 + v684
		v677 = v677 + v684
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v659)+12)) = v657
	*(*uint8)(unsafe.Add(mBase, uint32(v659)+3)) = uint8(v731)
	if v696 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L191:
	;
	v703 = int32(511453)
	v706 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1317])))
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v707 == int32(0) {
		v726 = v706
		v727 = v707
		goto L195
	} else {
		goto L196
	}
L192:
	;
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v701 != 0 {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v731 = int32(1)
	goto L190
L194:
	;
	v731 = base.B2i32(v727-v726 == int32(0))
	goto L190
L195:
	;
	goto L194
L196:
	;
	if v706 != v707 {
		v726 = v706
		v727 = v707
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v711 = v50
	v712 = v703
	goto L198
L198:
	;
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v712)+1)))
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711)+1)))
	if v716 == int32(0) {
		v726 = v715
		v727 = v716
		goto L195
	} else {
		goto L200
	}
L199:
	;
	v726 = v715
	v727 = v716
	goto L195
L200:
	;
	v719 = int32(1)
	if v715 == v716 {
		v711 = v711 + v719
		v712 = v712 + v719
		goto L198
	} else {
		goto L201
	}
L201:
	;
	goto L199
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v659)+8)) = int32(1742576)
	goto L204
L203:
	;
	goto L204
L204:
	;
	m.G0 = v9 + int32(32)
	return v659
}
func F_pg_analyze_and_rewrite_fixedparams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1209])))
	if v7 == int32(1) {
		v10 = int32(4444200)
		v11 = int32(0)
		v15 = m.G0
		v17 = v15 - int32(16)
		m.G0 = v17
		v20 = int32(4444216)
		v25 = F___memset(m, int32(4444224), v11, int32(144))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, _consts[1210])) = int32(4)
		*(*int64)(unsafe.Add(mBase, _consts[1211])) = int64(3)
		*(*int32)(unsafe.Add(mBase, _consts[1212])) = int32(2)
		*(*int64)(unsafe.Add(mBase, _consts[1213])) = int64(1)
		v38 = F___memcpy(m, v17, v20, int32(16))
		mBase = m.M
		v39 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17))))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
		v41 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[1214])) = v41
		*(*int32)(unsafe.Add(mBase, _consts[1215])) = v40
		*(*int64)(unsafe.Add(mBase, _consts[1216])) = v39
		v45 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17)+8)))
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
		*(*int32)(unsafe.Add(mBase, _consts[1217])) = v41
		*(*int32)(unsafe.Add(mBase, _consts[1212])) = v46
		*(*int64)(unsafe.Add(mBase, _consts[1213])) = v45
		v53 = F___syscall_ret(m, v11)
		mBase = m.M
		m.G0 = v17 + int32(16)
		F___gettimeofday(m, int32(4444352))
		mBase = m.M
	} else {
	}
	v59 = F_parse_analyze_fixedparams(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		return int32(0)
	} else {
		v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1209])))
		if v64 == int32(1) {
			F_ShowUsage(m, int32(525382))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v70 = F_pg_rewrite_query(m, v59)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					return v70
				}
			}
		} else {
			v70 = F_pg_rewrite_query(m, v59)
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int32(0)
			} else {
				return v70
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
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
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
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
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
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
	v21 = F_get_extension_control_directories(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v13 + int32(32)
	return int32(0)
L4:
	;
	if v21 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v25 <= int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v35 = v2
	v38 = v2
	goto L7
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v35<<(uint(int32(2))%32))))
	v45 = F_AllocateDir(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L3
L9:
	;
	v222 = v35 + int32(1)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v222 < v223 {
		v35 = v222
		v38 = v219
		goto L7
	} else {
		goto L66
	}
L10:
	;
	if v45 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v50 == int32(44) {
		v219 = v38
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v61 = v38
	goto L15
L14:
	;
	goto L13
L15:
	;
	v63 = F_ReadDir(m, v45, v44)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	F_FreeDir(m, v45)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L65
	}
L17:
	;
	if v63 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v66 = v63 + int32(19)
	v70 = F_strlen(m, v66)
	mBase = m.M
	v77 = v70 + int32(1)
	goto L23
L19:
	;
	goto L20
L20:
	;
	goto L16
L21:
	;
	if v89 == int32(0) {
		goto L15
	} else {
		goto L27
	}
L22:
	;
	goto L21
L23:
	;
	v79 = int32(0)
	if v77 == v79 {
		v89 = v79
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v89 = v84
	goto L22
L25:
	;
	v83 = v77 - int32(1)
	v84 = v66 + v83
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v85 != int32(46) {
		v77 = v83
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v92 = int32(302303)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, _consts[452])))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v96 == int32(0) {
		v115 = v95
		v116 = v96
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v116-v115 != 0 {
		goto L15
	} else {
		goto L36
	}
L29:
	;
	goto L28
L30:
	;
	if v95 != v96 {
		v115 = v95
		v116 = v96
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v100 = v89
	v101 = v92
	goto L32
L32:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	if v105 == int32(0) {
		v115 = v104
		v116 = v105
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v115 = v104
	v116 = v105
	goto L29
L34:
	;
	v108 = int32(1)
	if v104 == v105 {
		v100 = v100 + v108
		v101 = v101 + v108
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v118 = F_pstrdup(m, v66)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v123 = F_strlen(m, v118)
	mBase = m.M
	v130 = v123 + int32(1)
	goto L40
L38:
	;
	v143 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v143)
	v146 = F_strstr(m, v118, int32(671928))
	mBase = m.M
	if v146 != 0 {
		goto L15
	} else {
		goto L44
	}
L39:
	;
	goto L38
L40:
	;
	v132 = int32(0)
	if v130 == v132 {
		v142 = v132
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v142 = v137
	goto L39
L42:
	;
	v136 = v130 - int32(1)
	v137 = v118 + v136
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v138 != int32(46) {
		v130 = v136
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v147 = F_makeString(m, v118)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v149 = F_list_member(m, v61, v147)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v149 != 0 {
		goto L15
	} else {
		goto L47
	}
L47:
	;
	v151 = F_lappend(m, v61, v147)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v154 = F_palloc0(m, int32(48))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v156 = F_pstrdup(m, v118)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+36)) = int32(-1)
	v160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+34)) = uint8(v160)
	v162 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v154)+32)) = uint16(v162)
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v156
	v165 = F_pstrdup(m, v44)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v165
	F_parse_extension_control_file(m, v154, int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v171 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(24)))) = v171
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v171)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)) = uint16(v171)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v182 = F_DirectFunctionCall1Coll(m, int32(500), v171, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v182
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	if v185 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v154)+24))
	if v193 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	v188 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v188)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v190 = F_cstring_to_text(m, v185)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v190
	goto L54
L59:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	F_tuplestore_putvalues(m, v201, v202, v13+int32(16), v13+int32(12))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L64
	}
L60:
	;
	v196 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v196)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v198 = F_cstring_to_text(m, v193)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v198
	goto L59
L64:
	;
	v61 = v151
	goto L15
L65:
	;
	v219 = v61
	goto L9
L66:
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
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
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
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = l0 + l1
	if base.Ui32(l0) < base.Ui32(v16) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L26
	} else {
		goto L48
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L26
	} else {
		goto L43
	}
L3:
	;
	v18 = l0
	v21 = l2
	v23 = v4
	v24 = v4
	v25 = v4
	goto L6
L4:
	;
	v161 = l2
	goto L5
L5:
	;
	m.G0 = v14 + int32(16)
	return base.I64_extend_i32_s(v161 - l2)
L6:
	;
	v30 = v18
	goto L9
L7:
	;
	if v157 != 0 {
		goto L1
	} else {
		goto L42
	}
L8:
	;
	goto L7
L9:
	;
	v41 = v30 + int32(1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v44 = v42 - int32(9)
	if base.Ui32(int32(23)) < base.Ui32(v44) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v42 == int32(61) {
		goto L20
	} else {
		goto L21
	}
L11:
	;
	goto L10
L12:
	;
	if int32(1)<<(uint(v44)%32)&int32(8388627) == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if base.Ui32(v41) < base.Ui32(v16) {
		v30 = v41
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v153 = v21
	v157 = v25
	goto L8
L15:
	;
	if base.Ui32(v41) < base.Ui32(v16) {
		v18 = v41
		v21 = v146
		v23 = v148
		v24 = v149
		v25 = v150
		goto L6
	} else {
		goto L41
	}
L16:
	;
	v132 = int32(0)
	if base.B2i32(v128 == v132)&base.B2i32(base.Ui32(v129) < base.Ui32(int32(3))) == v132 {
		goto L38
	} else {
		goto L39
	}
L17:
	;
	v123 = int32(base.Ui32(v118) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v123)
	v127 = v118
	v128 = v119
	v129 = v120
	v131 = v21 + int32(2)
	goto L16
L18:
	;
	v113 = int32(base.Ui32(v24) >> (uint(int32(10)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v113)
	v118 = v24 << (uint(int32(6)) % 32)
	v119 = v56
	v120 = int32(2)
	goto L17
L19:
	;
	v95 = v92 + v24<<(uint(int32(6))%32)
	v97 = v25 + int32(1)
	if v97 != int32(4) {
		goto L33
	} else {
		goto L34
	}
L20:
	;
	v56 = int32(0)
	if v23 != 0 {
		v92 = v56
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if base.Ui32(int32(125)) < base.Ui32((v42-int32(1))&int32(255)) {
		goto L2
	} else {
		goto L31
	}
L23:
	;
	switch v25 - int32(2) {
	case 0:
		goto L25
	case 1:
		goto L18
	default:
		goto L24
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v146 = v21
	v148 = int32(1)
	v149 = v24 << (uint(int32(6)) % 32)
	v150 = int32(3)
	goto L15
L26:
	;
	return int64(0)
L27:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(417496), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(500984), int32(365), int32(415042))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+uint32(_consts[1268]))))
	if v89 < int32(0) {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v92 = v89
	goto L19
L33:
	;
	v146 = v21
	v148 = v23
	v149 = v95
	v150 = v97
	goto L15
L34:
	;
	goto L35
L35:
	;
	v101 = int32(base.Ui32(v95) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v101)
	v104 = base.B2i32(v23 == int32(0))
	if v23 == int32(0) {
		v118 = v95
		v119 = v104
		v120 = v23
		goto L17
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v23) {
		v118 = v95
		v119 = v104
		v120 = v23
		goto L17
	} else {
		goto L37
	}
L37:
	;
	v127 = v95
	v128 = int32(0)
	v129 = v23
	v131 = v21 + int32(1)
	goto L16
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v127)
	v143 = v131 + int32(1)
	goto L40
L39:
	;
	v143 = v131
	goto L40
L40:
	;
	v146 = v143
	v148 = v129
	v149 = v132
	v150 = int32(0)
	goto L15
L41:
	;
	v153 = v146
	v157 = v150
	goto L8
L42:
	;
	v161 = v153
	goto L5
L43:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L26
	} else {
		goto L44
	}
L44:
	;
	v183 = F_pg_mblen_range(m, v30, v16)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v183
	F_errmsg(m, int32(417437), v14)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L26
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(500984), int32(378), int32(415042))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L26
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
	v201 = m.ExcPending
	if v201 != 0 {
		goto L26
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(417360), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L26
	} else {
		goto L50
	}
L50:
	;
	F_errhint(m, int32(649797), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L26
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(500984), int32(399), int32(415042))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L26
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
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	v8 = l0 + l1
	if base.Ui32(v8) <= base.Ui32(l0) {
		v115 = l2
	} else {
		v13 = l0
		v14 = int32(2)
		v16 = l2
		v17 = int32(0)
		v18 = l2 + int32(76)
		for {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			v24 = v20<<(uint(v14<<(uint(int32(3))%32))%32) | v17
			if int32(0) < v14 {
				v63 = v16
				v64 = v24
				v65 = v14 - int32(1)
			} else {
				v29 = int32(63)
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24&v29)+uint32(_consts[1267]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v16)+3)) = uint8(v33)
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v24)>>(uint(int32(6))%32))&v29)+uint32(_consts[1267]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v16)+2)) = uint8(v41)
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v24)>>(uint(int32(12))%32))&v29)+uint32(_consts[1267]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)) = uint8(v49)
				v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v24)>>(uint(int32(18))%32))&v29)+uint32(_consts[1267]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v57)
				v63 = v16 + int32(4)
				v64 = int32(0)
				v65 = int32(2)
			}
			if base.Ui32(v18) <= base.Ui32(v63) {
				v67 = int32(10)
				*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v67)
				v73 = v63 + int32(1)
				v74 = v63 + int32(77)
			} else {
				v73 = v63
				v74 = v18
			}
			v76 = v13 + int32(1)
			if v76 != v8 {
				v13 = v76
				v14 = v65
				v16 = v73
				v17 = v64
				v18 = v74
				continue
			} else {
				break
			}
			break
		}
		if v65 == int32(2) {
			v115 = v73
		} else {
			v82 = int32(63)
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v64)>>(uint(int32(12))%32))&v82)+uint32(_consts[1267]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)) = uint8(v86)
			v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v64)>>(uint(int32(18))%32))&v82)+uint32(_consts[1267]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v94)
			if v65 == int32(0) {
				v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v64)>>(uint(int32(6))%32))&int32(63))+uint32(_consts[1267]))))
				v106 = v105
			} else {
				v106 = int32(61)
			}
			v107 = int32(61)
			*(*uint8)(unsafe.Add(mBase, uint32(v73)+3)) = uint8(v107)
			*(*uint8)(unsafe.Add(mBase, uint32(v73)+2)) = uint8(v106)
			v115 = v73 + int32(4)
		}
	}
	return base.I64_extend_i32_s(v115 - l2)
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	if l1 <= int32(0) {
		v37 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v37 - l0
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
	v37 = v30
	goto L1
L5:
	;
	v30 = v29 + v10
	v31 = v8 - v29
	if int32(0) < v31 {
		v8 = v31
		v10 = v30
		goto L3
	} else {
		goto L13
	}
L6:
	;
	if v11 == int32(0) {
		v37 = v10
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
		v37 = v10
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v29 = int32(1)
	goto L5
L10:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if base.B2i32(v11 == int32(-115))&base.B2i32(v21 == int32(32)) != 0 {
		v37 = v10
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v21 == int32(0) {
		v37 = v10
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v29 = int32(2)
	goto L5
L13:
	;
	goto L4
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
		v20 = F_bms_make_singleton(m, (v3<<(uint(v14)%32)+int32(458752))>>(uint(v14)%32))
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
						F_errmsg_internal(m, int32(50705), v7)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(501619), int32(5389), int32(437570))
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
			v18 = *(*int32)(unsafe.Add(mBase, _consts[86]))
			v22 = F_LWLockAcquire(m, v18+int32(1152), int32(1))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _consts[644]))
				v28 = F_get_controlfile(m, v25, v6+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _consts[86]))
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
								F_errmsg(m, int32(388849), int32(0))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(501596), int32(222), int32(100515))
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
				F_errmsg_internal(m, int32(368493), int32(0))
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(501596), int32(214), int32(100515))
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
	v9 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = F_DirectFunctionCall1Coll(m, int32(500), v3, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = F_DirectFunctionCall3Coll(m, int32(1652), v3, v4, v5, v11)
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v187 int64
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int64
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
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
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
	if v16 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L68
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L63
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
	v21 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+316))
	v24 = base.B2i32(v22 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v24)
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
	v30 = *(*int32)(unsafe.Add(mBase, _consts[27]))
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
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L58
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
	F___gettimeofday(m, v47)
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
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L48
	}
L17:
	;
	v175 = F_strlen(m, v164)
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
	v168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v165))) = uint8(v168)
	goto L17
L22:
	;
	v149 = v144
	v150 = v145
	v151 = v146
	goto L44
L23:
	;
	if v139 == int32(0) {
		v164 = v137
		v165 = v138
		goto L21
	} else {
		goto L43
	}
L24:
	;
	v137 = v33
	v138 = v62
	v139 = v69
	goto L23
L25:
	;
	goto L26
L26:
	;
	if v33&int32(3) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v106 == int32(0) {
		v164 = v103
		v165 = v104
		goto L21
	} else {
		goto L36
	}
L28:
	;
	v103 = v33
	v104 = v62
	v105 = v69
	v106 = int32(1)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v82 = v33
	v83 = v62
	v84 = v69
	goto L31
L31:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	*(*uint8)(unsafe.Add(mBase, uint32(v83))) = uint8(v86)
	if v86 == int32(0) {
		v144 = v82
		v145 = v83
		v146 = v84
		goto L22
	} else {
		goto L33
	}
L32:
	;
	v103 = v97
	v104 = v91
	v105 = v93
	v106 = v95
	goto L27
L33:
	;
	v90 = int32(1)
	v91 = v83 + v90
	v93 = v84 - v90
	v94 = int32(0)
	v95 = base.B2i32(v93 != v94)
	v97 = v82 + v90
	if v97&int32(3) == v94 {
		v103 = v97
		v104 = v91
		v105 = v93
		v106 = v95
		goto L27
	} else {
		goto L34
	}
L34:
	;
	if v93 != 0 {
		v82 = v97
		v83 = v91
		v84 = v93
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v109 == int32(0) {
		v137 = v103
		v138 = v104
		v139 = v105
		goto L23
	} else {
		goto L37
	}
L37:
	;
	if base.Ui32(v105) < base.Ui32(int32(4)) {
		v137 = v103
		v138 = v104
		v139 = v105
		goto L23
	} else {
		goto L38
	}
L38:
	;
	v115 = v103
	v116 = v104
	v117 = v105
	goto L39
L39:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v123 = int32(-2139062144)
	if (int32(16843008)-v120|v120)&v123 != v123 {
		v144 = v115
		v145 = v116
		v146 = v117
		goto L22
	} else {
		goto L41
	}
L40:
	;
	v137 = v131
	v138 = v129
	v139 = v133
	goto L23
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v120
	v128 = int32(4)
	v129 = v116 + v128
	v131 = v115 + v128
	v133 = v117 - v128
	if base.Ui32(int32(3)) < base.Ui32(v133) {
		v115 = v131
		v116 = v129
		v117 = v133
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v144 = v137
	v145 = v138
	v146 = v139
	goto L22
L44:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	*(*uint8)(unsafe.Add(mBase, uint32(v150))) = uint8(v153)
	if v153 == int32(0) {
		v164 = v149
		v165 = v150
		goto L21
	} else {
		goto L46
	}
L45:
	;
	v164 = v160
	v165 = v158
	goto L21
L46:
	;
	v157 = int32(1)
	v158 = v150 + v157
	v160 = v149 + v157
	v162 = v151 - v157
	if v162 != 0 {
		v149 = v160
		v150 = v158
		v151 = v162
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	F_XLogRegisterData(m, v40+int32(24), int32(72))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v187 = F_XLogInsert(m, int32(0), int32(112))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v191 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v191 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v33
	*(*uint32)(unsafe.Add(mBase, uint32(v40)+8)) = uint32(v187)
	v196 = int64(base.Ui64(v187) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v40)+4)) = uint32(v196)
	F_errmsg(m, int32(514671), v40)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	m.G0 = v40 + int32(96)
	v209 = F_Int64GetDatum(m, v187)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	F_errfinish(m, int32(499589), int32(8140), int32(89494))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	m.G0 = v7 + int32(16)
	return v209
L58:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(128519), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errhint(m, int32(575095), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(496318), int32(243), int32(88947))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(89251), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errhint(m, int32(581022), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(496318), int32(249), int32(88947))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(63)
	F_errmsg(m, int32(675404), v7)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(496318), int32(256), int32(88947))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
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
	var v22 int32
	_ = v22
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
				v22 = F___memset(m, v4, int32(0), int32(12))
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
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(428892)
			F_errmsg(m, int32(193003), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(495589), int32(358), int32(280690))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_pg_dearmor(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
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
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
		if v15 == int32(1) {
			v18 = int32(4)
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
			if v20&int32(254) == int32(2) {
				v29 = v18
			} else {
				v29 = base.B2i32(v20 == int32(18)) << (uint(v18) % 32)
			}
			if v20 == int32(1) {
				v32 = v18
			} else {
				v32 = v29
			}
			v45 = v32
		} else {
			v33 = int32(1)
			if v15&v33 != 0 {
				v45 = int32(base.Ui32(v15)>>(uint(v33)%32)) - v33
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		F_initStringInfo(m, v8)
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int32(0)
		} else {
			v48 = int32(1)
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			if v50&v48 != 0 {
				v53 = v48
			} else {
				v53 = int32(4)
			}
			v55 = F_pgp_armor_decode(m, v11+v53, v45, v8)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				if int32(0) <= v55 {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					v62 = F_palloc(m, v59+int32(4))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v62))) = v64<<(uint(int32(2))%32) + int32(16)
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						if v73 != 0 {
							v74 = F__emscripten_memcpy_bulkmem(m, v62+int32(4), v72, v73)
							mBase = m.M
						} else {
						}
						F_pfree(m, v72)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v78 != v11 {
								F_pfree(m, v11)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(16)
									return v62
								}
							} else {
								m.G0 = v8 + int32(16)
								return v62
							}
						}
					}
				} else {
					F_px_THROW_ERROR(m, v55)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
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
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v123 float64
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	v18 = F_statext_dependencies_deserialize(m, v14)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_initStringInfo(m, v9+int32(-16))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_appendStringInfoChar(m, v9+int32(-16), int32(123))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v35 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	F_appendStringInfoChar(m, v9+int32(-16), int32(125))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L34
	}
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(12)+v35<<(uint(int32(2))%32))))
	if int32(0) < v35 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	F_appendStringInfoString(m, v9+int32(-16), int32(748219))
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
	F_appendStringInfoChar(m, v9+int32(-16), int32(34))
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
	v123 = *(*float64)(unsafe.Add(mBase, uint32(v43)))
	*(*float64)(unsafe.Add(mBase, uint32(v11))) = v123
	F_appendStringInfo(m, v9+int32(-16), int32(341143), v11)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
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
	F_appendStringInfoString(m, v9+int32(-16), int32(747237))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v66
	F_appendStringInfo(m, v9+int32(-16), int32(489630), v9+int32(-32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	if v75 < int32(2) {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v81 = int32(1)
	v86 = v75
	goto L24
L24:
	;
	if v81 == v86-int32(1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L16
L26:
	;
	v96 = int32(747237)
	goto L28
L27:
	;
	v96 = int32(748219)
	goto L28
L28:
	;
	F_appendStringInfoString(m, v9+int32(-16), v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v102 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43+int32(10)+v81<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v102
	F_appendStringInfo(m, v9+int32(-16), int32(489630), v9+int32(-48))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v112 = v81 + int32(1)
	v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	if v112 < v113 {
		v81 = v112
		v86 = v113
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L25
L32:
	;
	v131 = v35 + int32(1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if base.Ui32(v131) < base.Ui32(v132) {
		v35 = v131
		goto L9
	} else {
		goto L33
	}
L33:
	;
	goto L10
L34:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	m.G0 = v11 - int32(-64)
	return v147
}
func F_pg_dependencies_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(168990)
			F_errmsg(m, int32(193003), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(495646), int32(714), int32(36671))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
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
	v55 = F_downcase_truncate_identifier(m, v25, v53, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v28 = int32(4)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v30&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v43 = int32(1)
	if v24 != 0 {
		v53 = int32(base.Ui32(v22)>>(uint(v43)%32)) - v43
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v39 = v28
	goto L12
L11:
	;
	v39 = base.B2i32(v30 == int32(18)) << (uint(v28) % 32)
	goto L12
L12:
	;
	if v30 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v42 = v28
	goto L15
L14:
	;
	v42 = v39
	goto L15
L15:
	;
	v53 = v42
	goto L6
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v59 = F_px_find_digest(m, v55, v11+int32(12))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v59 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_pfree(m, v55)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
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
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L51
	}
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = m.T0[v66].(func(*base.Module, int32) int32)(m, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v70 = v67 + int32(4)
	v71 = F_palloc(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v70 << (uint(int32(2)) % 32)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v77 = F_pg_detoast_datum_packed(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v110 = int32(1)
	if v79&v110 != 0 {
		goto L37
	} else {
		goto L38
	}
L26:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v79 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v82 = int32(4)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	if v84&int32(254) == int32(2) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v97 = int32(1)
	if v79&v97 != 0 {
		v109 = int32(base.Ui32(v79)>>(uint(v97)%32)) - v97
		goto L25
	} else {
		goto L36
	}
L30:
	;
	v93 = v82
	goto L32
L31:
	;
	v93 = base.B2i32(v84 == int32(18)) << (uint(v82) % 32)
	goto L32
L32:
	;
	if v84 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = v82
	goto L35
L34:
	;
	v96 = v93
	goto L35
L35:
	;
	v109 = v96
	goto L25
L36:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
	goto L25
L37:
	;
	v114 = v110
	goto L39
L38:
	;
	v114 = int32(4)
	goto L39
L39:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	m.T0[v116].(func(*base.Module, int32, int32, int32))(m, v65, v77+v114, v109)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	m.T0[v121].(func(*base.Module, int32, int32))(m, v65, v71+int32(4))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	m.T0[v124].(func(*base.Module, int32))(m, v65)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v127 != v77 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_pfree(m, v77)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v131 != v14 {
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
	v134 = m.ExcPending
	if v134 != 0 {
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
	return v71
L50:
	;
	goto L49
L51:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v59 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v55
	F_errmsg(m, int32(206270), v11)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L67
	}
L54:
	;
	v179 = int32(316408)
	goto L53
L55:
	;
	goto L56
L56:
	;
	v156 = int32(4400400)
	goto L58
L57:
	;
	v179 = v174
	goto L53
L58:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)+8))
	if v59 != v159 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v174 = v171
	goto L57
L60:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v156)+20))
	if v162 == int32(0) {
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
	v179 = int32(415110)
	goto L53
L64:
	;
	goto L65
L65:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	if v59 != v167 {
		v156 = v156 + int32(16)
		goto L58
	} else {
		goto L66
	}
L66:
	;
	v174 = v162
	goto L57
L67:
	;
	F_errfinish(m, int32(497529), int32(513), int32(228080))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v4 - int32(142) {
	case 0:
		v15 = int32(2)
		return v15
	case 1:
		return int32(3)
	default:
		if int32(0) <= base.I32_extend8_s(v4) {
			v14 = int32(1)
		} else {
			v14 = int32(2)
		}
		v15 = v14
		return v15
	}
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
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	if l1 <= int32(0) {
		v71 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v71 - l0
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
	v71 = v65
	goto L1
L5:
	;
	v65 = v64 + v10
	v66 = v9 - v64
	if int32(0) < v66 {
		v9 = v66
		v10 = v65
		goto L3
	} else {
		goto L22
	}
L6:
	;
	if v14 == int32(0) {
		v71 = v10
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
	v64 = int32(1)
	goto L5
L10:
	;
	v64 = int32(2)
	goto L5
L11:
	;
	if v9 == int32(1) {
		v71 = v10
		goto L1
	} else {
		goto L19
	}
L12:
	;
	if base.Ui32(v9) < base.Ui32(int32(3)) {
		v71 = v10
		goto L1
	} else {
		goto L16
	}
L13:
	;
	if v9 == int32(1) {
		v71 = v10
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
	v71 = v10
	goto L1
L16:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if base.Ui32(int32(93)) < base.Ui32((v33+int32(95))&int32(255)) {
		v71 = v10
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+2)))
	if base.Ui32(int32(94)) <= base.Ui32((v40+int32(95))&int32(255)) {
		v71 = v10
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v64 = int32(3)
	goto L5
L19:
	;
	if base.Ui32(int32(93)) < base.Ui32((v14+int32(95))&int32(255)) {
		v71 = v10
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if base.Ui32(int32(93)) < base.Ui32((v56+int32(95))&int32(255)) {
		v71 = v10
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L10
L22:
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
	v8 = *(*int32)(unsafe.Add(mBase, _consts[427]))
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
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(686369)
					F_errmsg(m, int32(253159), v5)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(496942), int32(1630), int32(435977))
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
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(686369)
				F_errmsg(m, int32(253159), v5)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(496942), int32(1630), int32(435977))
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
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
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
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
	v50 = v43
	v53 = int32(0)
	goto L11
L11:
	;
	if int32(0) < v50 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L7
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v53<<(uint(int32(2))%32))))
	v68 = v50
	v72 = int32(0)
	goto L16
L14:
	;
	v198 = v50
	goto L15
L15:
	;
	v207 = v53 + int32(1)
	if v207 < v198 {
		v50 = v198
		v53 = v207
		goto L11
	} else {
		goto L42
	}
L16:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v72<<(uint(int32(2))%32))))
	if v80 != v64 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v198 = v185
	goto L15
L18:
	;
	v84 = F_find_update_path(m, v39, v64, v80, int32(0), int32(1))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v185 = v68
	goto L20
L20:
	;
	v194 = v72 + int32(1)
	if v194 < v185 {
		v68 = v185
		v72 = v194
		goto L16
	} else {
		goto L41
	}
L21:
	;
	v86 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(40)))) = v86
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+30)) = uint8(v86)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+28)) = uint16(v86)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v95
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v99 = F_cstring_to_text(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v99
	if v84 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	F_tuplestore_putvalues(m, v174, v175, v13+int32(32), v13+int32(28))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L40
	}
L25:
	;
	v104 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+30)) = uint8(v104)
	goto L24
L26:
	;
	goto L27
L27:
	;
	F_initStringInfo(m, v13+int32(12))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	F_appendStringInfoString(m, v13+int32(12), v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
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
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v158 = F_cstring_to_text(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L38
	}
L33:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129+v119<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v13+int32(12), int32(671928))
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
	F_appendStringInfoString(m, v13+int32(12), v133)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v144 = v119 + int32(1)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v144 < v145 {
		v119 = v144
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v158
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_pfree(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L24
L40:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v185 = v182
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v164 int64
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int64
	_ = v177
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int64
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int64
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int64
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int64
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v435 int32
	_ = v435
	v18 = m.G0
	v20 = v18 - int32(1168)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = int64(34359738372)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v26
	v33 = F_hash_create(m, int32(115654), int32(256), v20+int32(16), int32(1064))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
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
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v41
	v44 = int32(1)
	v46 = F_list_make1_impl(m, v44, v20)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v48
	if v46 == v48 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_hash_destroy(m, v33)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L83
	}
L6:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v52 <= int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v68 = int32(0)
	v70 = v44
	v72 = v46
	goto L8
L8:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v68<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v82
	v89 = F_hash_search(m, v33, v20+int32(8), int32(1), v20+int32(7))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L5
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v70
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v94
	v96 = int32(0)
	if v94 == v96 {
		v145 = v96
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v164 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(1096)))) = v164
	*(*int64)(unsafe.Add(mBase, uint32(v20)+1088)) = v164
	v168 = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+32))
	m.T0[v174].(func(*base.Module, int32, int32, int32, int32, int32))(m, v94, v168, v168, v20+int32(1088), int32(1))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L23
	}
L12:
	;
	v99 = v96
	goto L13
L13:
	;
	v121 = F_hash_search(m, v33, v20-int32(-64), int32(0), v20+int32(1120))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L20
	}
L15:
	;
	goto L14
L16:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1120)))
	if v123 == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v127 = F_lcons_int(m, v126, v99)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v20)+64))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v130
	if v130 != 0 {
		v99 = v127
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v145 = v127
	goto L11
L20:
	;
	F_errmsg_internal(m, int32(446401), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(496018), int32(98), int32(365535))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
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
	v177 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(1152)))) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(1144)))) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(1136)))) = v177
	v185 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v20+int32(1112)))) = uint16(v185)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+1128)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v20)+1120)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v20)+1104)) = v177
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v94)+32))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v94)+36))
	if v194 == v185 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	if v233 != 0 {
		goto L41
	} else {
		goto L42
	}
L25:
	;
	v231 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+1104)) = uint8(v231)
	v233 = v194
	goto L24
L26:
	;
	v228 = F_cstring_to_text(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L39
	}
L27:
	;
	if v193 == int32(0) {
		goto L25
	} else {
		goto L38
	}
L28:
	;
	v197 = int32(323855)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1306])))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	if v201 == int32(0) {
		v220 = v200
		v221 = v201
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v221-v220 != 0 {
		goto L27
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	if v200 != v201 {
		v220 = v200
		v221 = v201
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v205 = v193
	v206 = v197
	goto L33
L33:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
	if v210 == int32(0) {
		v220 = v209
		v221 = v210
		goto L30
	} else {
		goto L35
	}
L34:
	;
	v220 = v209
	v221 = v210
	goto L30
L35:
	;
	v213 = int32(1)
	if v209 == v210 {
		v205 = v205 + v213
		v206 = v206 + v213
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v226 = int32(0)
	v227 = v194
	goto L26
L38:
	;
	v226 = v194
	v227 = v193
	goto L26
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1120)) = v228
	v233 = v226
	goto L24
L40:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v262 = v260 - int32(474)
	if base.Ui32(v262) <= base.Ui32(int32(3)) {
		goto L53
	} else {
		goto L54
	}
L41:
	;
	v235 = F_strlen(m, v233)
	mBase = m.M
	if int32(1024) <= v235 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v256 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+1105)) = uint8(v256)
	goto L40
L44:
	;
	v239 = F_pg_mbcliplen(m, v233, v235, int32(1023))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	v241 = v235
	goto L46
L46:
	;
	if v241 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v241 = v239
	goto L46
L48:
	;
	v247 = v20 - int32(-64)
	v249 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v247+v241))) = uint8(v249)
	v253 = F_cstring_to_text(m, v247)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L52
	}
L49:
	;
	v244 = F__emscripten_memcpy_bulkmem(m, v20-int32(-64), v233, v241)
	mBase = m.M
	goto L51
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1124)) = v253
	goto L40
L53:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v262<<(uint(int32(2))%32))+uint32(_consts[1307])))
	v270 = v269
	goto L55
L54:
	;
	v270 = int32(547298)
	goto L55
L55:
	;
	v271 = F_cstring_to_text(m, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1128)) = v271
	if v145 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v338 = F_construct_array_builtin(m, v323, v325, int32(23))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L67
	}
L58:
	;
	v276 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1132)) = v276
	v280 = F_palloc(m, v276)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1132)) = v282
	v286 = F_palloc(m, v282<<(uint(int32(2))%32))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	v323 = v280
	v325 = v276
	goto L57
L62:
	;
	v288 = int32(0)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	if v289 <= v288 {
		v323 = v286
		v325 = v282
		goto L57
	} else {
		goto L63
	}
L63:
	;
	v294 = v288
	goto L64
L64:
	;
	v310 = v294 << (uint(int32(2)) % 32)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v145)+12))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v312+v310)))
	*(*int32)(unsafe.Add(mBase, uint32(v286+v310))) = v314
	v317 = v294 + int32(1)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	if v317 < v318 {
		v294 = v317
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v323 = v286
	v325 = v282
	goto L57
L66:
	;
	goto L65
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1136)) = v338
	v341 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v20)+1096)))
	v342 = F_Int64GetDatum(m, v341)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1140)) = v342
	v345 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v20)+1088)))
	v346 = F_Int64GetDatum(m, v345)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1144)) = v346
	v349 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v20)+1100)))
	v350 = F_Int64GetDatum(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1148)) = v350
	v353 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v20)+1092)))
	v354 = F_Int64GetDatum(m, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1152)) = v354
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v20)+1096))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v20)+1100))
	v361 = F_Int64GetDatum(m, base.I64_extend_i32_u(v357-v358))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1156)) = v361
	F_tuplestore_putvalues(m, v93, v92, v20+int32(1120), v20+int32(1104))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_list_free(m, v145)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+20))
	if v373 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v374 = v373
	v385 = v72
	goto L78
L76:
	;
	v405 = v72
	goto L77
L77:
	;
	v411 = int32(1)
	v414 = v68 + v411
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v414 < v415 {
		v68 = v414
		v70 = v70 + v411
		v72 = v405
		goto L8
	} else {
		goto L82
	}
L78:
	;
	v391 = F_lappend(m, v385, v374)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L80
	}
L79:
	;
	v405 = v391
	goto L77
L80:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v374)+28))
	if v393 != 0 {
		v374 = v393
		v385 = v391
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	goto L9
L83:
	;
	m.G0 = v20 + int32(1168)
	return int32(0)
}
func F_pg_get_constraintdef_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v31 int32
	_ = v31
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
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
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
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v436 int32
	_ = v436
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int64
	_ = v557
	var v564 int32
	_ = v564
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v725 int32
	_ = v725
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v768 int32
	_ = v768
	var v776 int32
	_ = v776
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	v11 = m.G0
	v13 = v11 - int32(432)
	m.G0 = v13
	v15 = F_GetTransactionSnapshot(m)
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
	v19 = F_RegisterSnapshot(m, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_ScanKeyInit(m, v13+int32(304), int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v33 = int32(1)
	v37 = F_systable_beginscan(m, v23, int32(2667), v33, v19, v33, v13+int32(304))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v39 = F_systable_getnext(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_UnregisterSnapshot(m, v19)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v39 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L220
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L217
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L214
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L211
	}
L13:
	;
	m.G0 = v13 + int32(432)
	return v768
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
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+22)))
	v66 = v64 + v65
	F_initStringInfo(m, v13+int32(360))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L25
	}
L17:
	;
	F_systable_endscan(m, v37)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
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
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	F_sequence_close(m, v23, int32(1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v768 = int32(0)
	goto L13
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
	F_errmsg_internal(m, int32(41309), v13)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(495391), int32(2232), int32(220861))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
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
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+72)))
	switch v145 - int32(99) {
	case 0:
		goto L51
	default:
		goto L48
	case 3:
		goto L54
	case 11:
		goto L50
	case 13:
		v307 = int32(745913)
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
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v66)+80))
	if v73 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v74 = F_generate_qualified_relation_name(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v66)+84))
	v91 = F_SearchSysCache1(m, int32(82), v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L34
	}
L31:
	;
	v78 = F_quote_identifier(m, v66+int32(4))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+292)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v13)+288)) = v74
	F_appendStringInfo(m, v13+int32(360), int32(740257), v13+int32(288))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
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
	F_initStringInfo(m, v13+int32(376))
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+272)) = v107
	F_appendStringInfo(m, v13+int32(376), int32(608647), v13+int32(272))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v121 = F_quote_identifier(m, v97+int32(4))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_appendStringInfoString(m, v13+int32(376), v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+376))
	F_ReleaseCatCache(m, v91)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v130 = F_quote_identifier(m, v66+int32(4))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+260)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v13)+256)) = v125
	F_appendStringInfo(m, v13+int32(360), int32(740222), v13+int32(256))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L26
L46:
	;
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+73)))
	if v725 == int32(1) {
		goto L195
	} else {
		goto L196
	}
L47:
	;
	F_appendStringInfoString(m, v13+int32(360), int32(527455))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L194
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L191
	}
L49:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v66)+88))
	v631 = F_SysCacheGetAttrNotNull(m, int32(19), v39, int32(27))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L179
	}
L50:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v66)+80))
	if v591 != 0 {
		goto L168
	} else {
		goto L169
	}
L51:
	;
	v483 = F_SysCacheGetAttrNotNull(m, int32(19), v39, int32(28))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L147
	}
L52:
	;
	F_appendStringInfoString(m, v13+int32(360), v307)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L96
	}
L53:
	;
	v307 = int32(746928)
	goto L52
L54:
	;
	F_appendStringInfoString(m, v13+int32(360), int32(687974))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v155 = F_SysCacheGetAttrNotNull(m, int32(19), v39, int32(21))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v66)+80))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+107)))
	v161 = F_decompile_column_index_array(m, v155, v157, v158, v13+int32(360))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v66)+96))
	v165 = F_generate_relation_name(m, v163, int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+144)) = v165
	F_appendStringInfo(m, v13+int32(360), int32(687422), v13+int32(144))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v177 = F_SysCacheGetAttrNotNull(m, int32(19), v39, int32(22))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v66)+96))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+107)))
	v183 = F_decompile_column_index_array(m, v177, v179, v180, v13+int32(360))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_appendStringInfoChar(m, v13+int32(360), int32(41))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+102)))
	switch v191 - int32(102) {
	case 0:
		v212 = int32(534951)
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
	F_appendStringInfoString(m, v13+int32(360), v212)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L70
	}
L64:
	;
	v212 = int32(535468)
	goto L63
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	v212 = int32(759461)
	goto L63
L67:
	;
	v199 = int32(*(*int8)(unsafe.Add(mBase, uint32(v66)+102)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v199
	F_errmsg_internal(m, int32(485562), v13-int32(-64))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(495391), int32(2309), int32(220861))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
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
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+100)))
	switch v218 - int32(97) {
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
		v240 = int32(522742)
		goto L72
	}
L71:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+101)))
	switch v251 - int32(97) {
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
		v273 = int32(522742)
		goto L82
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = v240
	F_appendStringInfo(m, v13+int32(360), int32(198812), v13+int32(128))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L80
	}
L73:
	;
	v240 = int32(543605)
	goto L72
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	v240 = int32(522060)
	goto L72
L76:
	;
	v240 = int32(534472)
	goto L72
L77:
	;
	v227 = int32(*(*int8)(unsafe.Add(mBase, uint32(v66)+100)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v227
	F_errmsg_internal(m, int32(485629), v13+int32(80))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(495391), int32(2335), int32(220861))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
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
	v287 = F_SysCacheGetAttr(m, int32(19), v39, int32(26), v13+int32(376))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L91
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v273
	F_appendStringInfo(m, v13+int32(360), int32(198765), v13+int32(112))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L90
	}
L83:
	;
	v273 = int32(543605)
	goto L82
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	v273 = int32(522060)
	goto L82
L86:
	;
	v273 = int32(534472)
	goto L82
L87:
	;
	v260 = int32(*(*int8)(unsafe.Add(mBase, uint32(v66)+101)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v260
	F_errmsg_internal(m, int32(485498), v13+int32(96))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(495391), int32(2361), int32(220861))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
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
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+376)))
	if v289 != 0 {
		goto L46
	} else {
		goto L92
	}
L92:
	;
	F_appendStringInfoString(m, v13+int32(360), int32(688852))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v66)+80))
	v299 = F_decompile_column_index_array(m, v287, v295, int32(0), v13+int32(360))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_appendStringInfoChar(m, v13+int32(360), int32(41))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	goto L46
L96:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v66)+88))
	v314 = F_SearchSysCache1(m, int32(34), v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v314 == int32(0) {
		goto L10
	} else {
		goto L98
	}
L98:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+72)))
	if v318 != int32(117) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	F_appendStringInfoChar(m, v13+int32(360), int32(40))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L103
	}
L100:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v314)+16))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+22)))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321+v322)+13)))
	if v324 != int32(1) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	F_appendStringInfoString(m, v13+int32(360), int32(746261))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	goto L99
L103:
	;
	v340 = F_SysCacheGetAttrNotNull(m, int32(19), v39, int32(21))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v66)+80))
	v346 = F_decompile_column_index_array(m, v340, v342, int32(0), v13+int32(360))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+107)))
	if v348 == int32(1) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	F_appendStringInfoString(m, v13+int32(360), int32(524566))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	F_appendStringInfoChar(m, v13+int32(360), int32(41))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	v363 = F_SysCacheGetAttrNotNull(m, int32(34), v314, int32(3))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v346 < v363 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_appendStringInfoString(m, v13+int32(360), int32(688790))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	F_ReleaseCatCache(m, v314)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L133
	}
L115:
	;
	v373 = F_SysCacheGetAttrNotNull(m, int32(34), v314, int32(16))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v375 = F_pg_detoast_datum(m, v373)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_deconstruct_array_builtin(m, v375, int32(21), v13+int32(376), int32(0), v13+int32(416))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v13)+416))
	if v346 < v385 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v387 = v346
	goto L122
L120:
	;
	goto L121
L121:
	;
	F_appendStringInfoChar(m, v13+int32(360), int32(41))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L132
	}
L122:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v66)+80))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v13)+376))
	v402 = int32(*(*int16)(unsafe.Add(mBase, uint32(v398+v387<<(uint(int32(2))%32)))))
	v404 = F_get_attname(m, v397, v402, int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	goto L121
L124:
	;
	if v346 < v387 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	F_appendStringInfoString(m, v13+int32(360), int32(748219))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v414 = F_quote_identifier(m, v404)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L129
	}
L128:
	;
	goto L127
L129:
	;
	F_appendStringInfoString(m, v13+int32(360), v414)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v419 = v387 + int32(1)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v13)+416))
	if v419 < v420 {
		v387 = v419
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
	if l1 == int32(0) {
		goto L46
	} else {
		goto L134
	}
L134:
	;
	if v313 == int32(0) {
		goto L46
	} else {
		goto L135
	}
L135:
	;
	v453 = F_flatten_reloptions(m, v313)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	if v453 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+240)) = v453
	F_appendStringInfo(m, v13+int32(360), int32(677605), v13+int32(240))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v465 = F_get_rel_tablespace(m, v313)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	F_pfree(m, v453)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	if v465 == int32(0) {
		goto L46
	} else {
		goto L143
	}
L143:
	;
	v469 = F_get_tablespace_name(m, v465)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v471 = F_quote_identifier(m, v469)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+224)) = v471
	F_appendStringInfo(m, v13+int32(360), int32(198856), v13+int32(224))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	goto L46
L147:
	;
	v485 = F_text_to_cstring(m, v483)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v487 = F_stringToNode(m, v485)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v66)+80))
	if v490 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v491 = F_get_rel_name(m, v490)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L153
	}
L151:
	;
	v545 = int32(0)
	goto L152
L152:
	;
	F_initStringInfo(m, v13+int32(416))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L162
	}
L153:
	;
	if v491 == int32(0) {
		goto L9
	} else {
		goto L154
	}
L154:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v66)+80))
	v497 = F_palloc0(m, int32(80))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v500 = F_palloc0(m, int32(136))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v500)+24)) = int32(1)
	v504 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v500)+21)) = uint8(v504)
	*(*int32)(unsafe.Add(mBase, uint32(v500)+16)) = v495
	v507 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v500)+12)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v500))) = int32(101)
	v512 = F_makeAlias(m, v491, v507)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v500)+8)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v500)+4)) = v512
	v516 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v500)+124)) = uint16(v516)
	v518 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v500)+20)) = uint8(v518)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v13)+376)) = v500
	v525 = F_list_make1_impl(m, int32(1), v13+int32(188))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v527 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v497)+20)) = v527
	*(*int64)(unsafe.Add(mBase, uint32(v497)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v497))) = v525
	F_set_rtable_names(m, v497, v527, v527)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_set_simple_column_names(m, v497)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v13)+416)) = v497
	v543 = F_list_make1_impl(m, int32(1), v13+int32(184))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v545 = v543
	goto L152
L162:
	;
	v553 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+408)) = uint8(v553)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+392)) = v553
	v557 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+384)) = v557
	*(*int32)(unsafe.Add(mBase, uint32(v13)+380)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v13)+412)) = v553
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+411)) = uint8(v553)
	v564 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+409)) = uint16(v564)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+400)) = v557
	*(*int32)(unsafe.Add(mBase, uint32(v13)+396)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+376)) = v13 + int32(416)
	F_get_rule_expr(m, v487, v13+int32(376), v553)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+106)))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v13)+416))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+160)) = v578
	if v577 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v582 = int32(522299)
	goto L166
L165:
	;
	v582 = int32(759461)
	goto L166
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+164)) = v582
	F_appendStringInfo(m, v13+int32(360), int32(177809), v13+int32(160))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	goto L46
L168:
	;
	v592 = F_extractNotNullColumn(m, v39)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v66)+84))
	if v619 == int32(0) {
		goto L46
	} else {
		goto L177
	}
L171:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v66)+80))
	v596 = F_get_attname(m, v594, v592, int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v598 = F_quote_identifier(m, v596)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+192)) = v598
	F_appendStringInfo(m, v13+int32(360), int32(198631), v13+int32(192))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608)+22)))
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608+v609)+106)))
	if v611 != int32(1) {
		goto L46
	} else {
		goto L175
	}
L175:
	;
	F_appendStringInfoString(m, v13+int32(360), int32(522299))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	goto L46
L177:
	;
	F_appendStringInfoString(m, v13+int32(360), int32(534463))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	goto L46
L179:
	;
	v633 = F_pg_detoast_datum(m, v631)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	F_deconstruct_array_builtin(m, v633, int32(26), v13+int32(376), int32(0), v13+int32(416))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v13)+416))
	v646 = F_palloc(m, v643<<(uint(int32(2))%32))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v13)+416))
	if int32(0) < v648 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v651 = int32(0)
	goto L186
L184:
	;
	goto L185
L185:
	;
	v684 = int32(0)
	v690 = F_pg_get_indexdef_worker(m, v627, v684, v646, v684, v684, v684, v684, l2, v684)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L189
	}
L186:
	;
	v662 = v651 << (uint(int32(2)) % 32)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v13)+376))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v664+v662)))
	*(*int32)(unsafe.Add(mBase, uint32(v646+v662))) = v666
	v669 = v651 + int32(1)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v13)+416))
	if v669 < v670 {
		v651 = v669
		goto L186
	} else {
		goto L188
	}
L187:
	;
	goto L185
L188:
	;
	goto L187
L189:
	;
	F_appendStringInfoString(m, v13+int32(360), v690)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	goto L46
L191:
	;
	v698 = int32(*(*int8)(unsafe.Add(mBase, uint32(v66)+72)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v698
	F_errmsg_internal(m, int32(731726), v13+int32(48))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(495391), int32(2592), int32(220861))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	goto L46
L195:
	;
	F_appendStringInfoString(m, v13+int32(360), int32(542595))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+74)))
	if v733 == int32(1) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	goto L197
L199:
	;
	F_appendStringInfoString(m, v13+int32(360), int32(545179))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+75)))
	if v743 != int32(1) {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	goto L201
L203:
	;
	F_systable_endscan(m, v37)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L209
	}
L204:
	;
	v749 = int32(545484)
	goto L206
L205:
	;
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+76)))
	if v747 != 0 {
		goto L203
	} else {
		goto L207
	}
L206:
	;
	F_appendStringInfoString(m, v13+int32(360), v749)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L208
	}
L207:
	;
	v749 = int32(544776)
	goto L206
L208:
	;
	goto L203
L209:
	;
	F_sequence_close(m, v23, int32(1))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v13)+360))
	v768 = v757
	goto L13
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v90
	F_errmsg_internal(m, int32(50705), v13+int32(16))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(495391), int32(13520), int32(380548))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v97)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v792
	F_errmsg_internal(m, int32(53155), v13+int32(32))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(495391), int32(13527), int32(380548))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+208)) = v313
	F_errmsg_internal(m, int32(40436), v13+int32(208))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(495391), int32(2401), int32(220861))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+176)) = v490
	F_errmsg_internal(m, int32(46584), v13+int32(176))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	F_errfinish(m, int32(495391), int32(13138), int32(380053))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
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
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
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
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int64
	_ = v259
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
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
			v290 = v2
			m.G0 = v14 + int32(80)
			return v290
		} else {
			v33 = F_get_func_arg_info(m, v19, v14+int32(20), v14+int32(16), v14+int32(12))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				if v16 <= int32(0) {
					F_ReleaseCatCache(m, v19)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						v57 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v57)
						v290 = int32(0)
						m.G0 = v14 + int32(80)
						return v290
					}
				} else {
					if v33 < v16 {
						F_ReleaseCatCache(m, v19)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							v57 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v57)
							v290 = int32(0)
							m.G0 = v14 + int32(80)
							return v290
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
						if v38 == int32(0) {
							v61 = int32(3)
							v62 = v16 & v61
							v63 = int32(0)
							if base.Ui32(v61) <= base.Ui32(v16-int32(1)) {
								v72 = v63
								v75 = v2
								v76 = int32(0)
								for {
									if v38 != 0 {
										v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v75))))
										v85 = v83 - int32(98)
										if base.Ui32(int32(20)) < base.Ui32(v85) {
											v97 = v72
										} else {
											if int32(1)<<(uint(v85)%32)&int32(1048705) == int32(0) {
												v97 = v72
											} else {
												v97 = v72 + int32(1)
											}
										}
									} else {
										v97 = v72 + int32(1)
									}
									if v38 != 0 {
										v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v75)+1)))
										v102 = v100 - int32(98)
										if base.Ui32(int32(20)) < base.Ui32(v102) {
											v114 = v97
										} else {
											if int32(1)<<(uint(v102)%32)&int32(1048705) == int32(0) {
												v114 = v97
											} else {
												v114 = v97 + int32(1)
											}
										}
									} else {
										v114 = v97 + int32(1)
									}
									if v38 != 0 {
										v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v75)+2)))
										v119 = v117 - int32(98)
										if base.Ui32(int32(20)) < base.Ui32(v119) {
											v131 = v114
										} else {
											if int32(1)<<(uint(v119)%32)&int32(1048705) == int32(0) {
												v131 = v114
											} else {
												v131 = v114 + int32(1)
											}
										}
									} else {
										v131 = v114 + int32(1)
									}
									if v38 != 0 {
										v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v75)+3)))
										v136 = v134 - int32(98)
										if base.Ui32(int32(20)) < base.Ui32(v136) {
											v148 = v131
										} else {
											if int32(1)<<(uint(v136)%32)&int32(1048705) == int32(0) {
												v148 = v131
											} else {
												v148 = v131 + int32(1)
											}
										}
									} else {
										v148 = v131 + int32(1)
									}
									v150 = int32(4)
									v151 = v75 + v150
									v153 = v76 + v150
									if v153 != v16&int32(2147483644) {
										v72 = v148
										v75 = v151
										v76 = v153
										continue
									} else {
										break
									}
									break
								}
								v156 = v148
								v159 = v151
							} else {
								v156 = v63
								v159 = v2
							}
							if v62 != 0 {
								v167 = v156
								v170 = v159
								v174 = v2
								for {
									if v38 != 0 {
										v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v170))))
										v180 = v178 - int32(98)
										if base.Ui32(int32(20)) < base.Ui32(v180) {
											v192 = v167
										} else {
											if int32(1)<<(uint(v180)%32)&int32(1048705) == int32(0) {
												v192 = v167
											} else {
												v192 = v167 + int32(1)
											}
										}
									} else {
										v192 = v167 + int32(1)
									}
									v194 = int32(1)
									v197 = v174 + v194
									if v197 != v62 {
										v167 = v192
										v170 = v170 + v194
										v174 = v197
										continue
									} else {
										break
									}
									break
								}
								v200 = v192
							} else {
								v200 = v156
							}
							v214 = F_SysCacheGetAttr(m, int32(47), v19, int32(24), v14+int32(11))
							mBase = m.M
							v215 = m.ExcPending
							if v215 != 0 {
								return int32(0)
							} else {
								v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
								if v216 == int32(1) {
									F_ReleaseCatCache(m, v19)
									mBase = m.M
									v220 = m.ExcPending
									if v220 != 0 {
										return int32(0)
									} else {
										v221 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v221)
										v290 = int32(0)
										m.G0 = v14 + int32(80)
										return v290
									}
								} else {
									v224 = F_text_to_cstring(m, v214)
									mBase = m.M
									v225 = m.ExcPending
									if v225 != 0 {
										return int32(0)
									} else {
										v226 = F_stringToNode(m, v224)
										mBase = m.M
										v227 = m.ExcPending
										if v227 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v224)
											mBase = m.M
											v229 = m.ExcPending
											if v229 != 0 {
												return int32(0)
											} else {
												v230 = int32(0)
												v231 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
												v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+22)))
												v233 = v231 + v232
												v234 = int32(*(*int16)(unsafe.Add(mBase, uint32(v233)+106)))
												v235 = int32(*(*int16)(unsafe.Add(mBase, uint32(v233)+104)))
												v239 = v200 + (v234 - v235) - int32(1)
												if v239 < v230 {
													F_ReleaseCatCache(m, v19)
													mBase = m.M
													v247 = m.ExcPending
													if v247 != 0 {
														return int32(0)
													} else {
														v248 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v248)
														v290 = v230
														m.G0 = v14 + int32(80)
														return v290
													}
												} else {
													if v226 == int32(0) {
														F_ReleaseCatCache(m, v19)
														mBase = m.M
														v247 = m.ExcPending
														if v247 != 0 {
															return int32(0)
														} else {
															v248 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v248)
															v290 = v230
															m.G0 = v14 + int32(80)
															return v290
														}
													} else {
														v244 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
														if v239 < v244 {
															v250 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
															v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+v239<<(uint(int32(2))%32))))
															F_initStringInfo(m, v14-int32(-64))
															mBase = m.M
															v258 = m.ExcPending
															if v258 != 0 {
																return int32(0)
															} else {
																v259 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(v14)+36)) = v259
																*(*int64)(unsafe.Add(mBase, uint32(v14)+44)) = v259
																*(*int64)(unsafe.Add(mBase, uint32(v14)+49)) = v259
																*(*int64)(unsafe.Add(mBase, uint32(v14)+28)) = v259
																v267 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v267
																*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)) = uint8(v267)
																v271 = int32(1)
																*(*uint16)(unsafe.Add(mBase, uint32(v14)+57)) = uint16(v271)
																*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v14 - int32(-64)
																F_get_rule_expr(m, v254, v14+int32(24), v267)
																mBase = m.M
																v280 = m.ExcPending
																if v280 != 0 {
																	return int32(0)
																} else {
																	v281 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
																	F_ReleaseCatCache(m, v19)
																	mBase = m.M
																	v283 = m.ExcPending
																	if v283 != 0 {
																		return int32(0)
																	} else {
																		v284 = F_cstring_to_text(m, v281)
																		mBase = m.M
																		v285 = m.ExcPending
																		if v285 != 0 {
																			return int32(0)
																		} else {
																			F_pfree(m, v281)
																			mBase = m.M
																			v287 = m.ExcPending
																			if v287 != 0 {
																				return int32(0)
																			} else {
																				v290 = v284
																				m.G0 = v14 + int32(80)
																				return v290
																			}
																		}
																	}
																}
															}
														} else {
															F_ReleaseCatCache(m, v19)
															mBase = m.M
															v247 = m.ExcPending
															if v247 != 0 {
																return int32(0)
															} else {
																v248 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v248)
																v290 = v230
																m.G0 = v14 + int32(80)
																return v290
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
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v16-int32(1)))))
							v46 = v44 - int32(98)
							if base.Ui32(int32(20)) < base.Ui32(v46) {
								F_ReleaseCatCache(m, v19)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v57 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v57)
									v290 = int32(0)
									m.G0 = v14 + int32(80)
									return v290
								}
							} else {
								if int32(1)<<(uint(v46)%32)&int32(1048705) != 0 {
									v61 = int32(3)
									v62 = v16 & v61
									v63 = int32(0)
									if base.Ui32(v61) <= base.Ui32(v16-int32(1)) {
										v72 = v63
										v75 = v2
										v76 = int32(0)
										for {
											if v38 != 0 {
												v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v75))))
												v85 = v83 - int32(98)
												if base.Ui32(int32(20)) < base.Ui32(v85) {
													v97 = v72
												} else {
													if int32(1)<<(uint(v85)%32)&int32(1048705) == int32(0) {
														v97 = v72
													} else {
														v97 = v72 + int32(1)
													}
												}
											} else {
												v97 = v72 + int32(1)
											}
											if v38 != 0 {
												v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v75)+1)))
												v102 = v100 - int32(98)
												if base.Ui32(int32(20)) < base.Ui32(v102) {
													v114 = v97
												} else {
													if int32(1)<<(uint(v102)%32)&int32(1048705) == int32(0) {
														v114 = v97
													} else {
														v114 = v97 + int32(1)
													}
												}
											} else {
												v114 = v97 + int32(1)
											}
											if v38 != 0 {
												v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v75)+2)))
												v119 = v117 - int32(98)
												if base.Ui32(int32(20)) < base.Ui32(v119) {
													v131 = v114
												} else {
													if int32(1)<<(uint(v119)%32)&int32(1048705) == int32(0) {
														v131 = v114
													} else {
														v131 = v114 + int32(1)
													}
												}
											} else {
												v131 = v114 + int32(1)
											}
											if v38 != 0 {
												v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v75)+3)))
												v136 = v134 - int32(98)
												if base.Ui32(int32(20)) < base.Ui32(v136) {
													v148 = v131
												} else {
													if int32(1)<<(uint(v136)%32)&int32(1048705) == int32(0) {
														v148 = v131
													} else {
														v148 = v131 + int32(1)
													}
												}
											} else {
												v148 = v131 + int32(1)
											}
											v150 = int32(4)
											v151 = v75 + v150
											v153 = v76 + v150
											if v153 != v16&int32(2147483644) {
												v72 = v148
												v75 = v151
												v76 = v153
												continue
											} else {
												break
											}
											break
										}
										v156 = v148
										v159 = v151
									} else {
										v156 = v63
										v159 = v2
									}
									if v62 != 0 {
										v167 = v156
										v170 = v159
										v174 = v2
										for {
											if v38 != 0 {
												v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v170))))
												v180 = v178 - int32(98)
												if base.Ui32(int32(20)) < base.Ui32(v180) {
													v192 = v167
												} else {
													if int32(1)<<(uint(v180)%32)&int32(1048705) == int32(0) {
														v192 = v167
													} else {
														v192 = v167 + int32(1)
													}
												}
											} else {
												v192 = v167 + int32(1)
											}
											v194 = int32(1)
											v197 = v174 + v194
											if v197 != v62 {
												v167 = v192
												v170 = v170 + v194
												v174 = v197
												continue
											} else {
												break
											}
											break
										}
										v200 = v192
									} else {
										v200 = v156
									}
									v214 = F_SysCacheGetAttr(m, int32(47), v19, int32(24), v14+int32(11))
									mBase = m.M
									v215 = m.ExcPending
									if v215 != 0 {
										return int32(0)
									} else {
										v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
										if v216 == int32(1) {
											F_ReleaseCatCache(m, v19)
											mBase = m.M
											v220 = m.ExcPending
											if v220 != 0 {
												return int32(0)
											} else {
												v221 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v221)
												v290 = int32(0)
												m.G0 = v14 + int32(80)
												return v290
											}
										} else {
											v224 = F_text_to_cstring(m, v214)
											mBase = m.M
											v225 = m.ExcPending
											if v225 != 0 {
												return int32(0)
											} else {
												v226 = F_stringToNode(m, v224)
												mBase = m.M
												v227 = m.ExcPending
												if v227 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v224)
													mBase = m.M
													v229 = m.ExcPending
													if v229 != 0 {
														return int32(0)
													} else {
														v230 = int32(0)
														v231 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
														v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+22)))
														v233 = v231 + v232
														v234 = int32(*(*int16)(unsafe.Add(mBase, uint32(v233)+106)))
														v235 = int32(*(*int16)(unsafe.Add(mBase, uint32(v233)+104)))
														v239 = v200 + (v234 - v235) - int32(1)
														if v239 < v230 {
															F_ReleaseCatCache(m, v19)
															mBase = m.M
															v247 = m.ExcPending
															if v247 != 0 {
																return int32(0)
															} else {
																v248 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v248)
																v290 = v230
																m.G0 = v14 + int32(80)
																return v290
															}
														} else {
															if v226 == int32(0) {
																F_ReleaseCatCache(m, v19)
																mBase = m.M
																v247 = m.ExcPending
																if v247 != 0 {
																	return int32(0)
																} else {
																	v248 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v248)
																	v290 = v230
																	m.G0 = v14 + int32(80)
																	return v290
																}
															} else {
																v244 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
																if v239 < v244 {
																	v250 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
																	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+v239<<(uint(int32(2))%32))))
																	F_initStringInfo(m, v14-int32(-64))
																	mBase = m.M
																	v258 = m.ExcPending
																	if v258 != 0 {
																		return int32(0)
																	} else {
																		v259 = int64(0)
																		*(*int64)(unsafe.Add(mBase, uint32(v14)+36)) = v259
																		*(*int64)(unsafe.Add(mBase, uint32(v14)+44)) = v259
																		*(*int64)(unsafe.Add(mBase, uint32(v14)+49)) = v259
																		*(*int64)(unsafe.Add(mBase, uint32(v14)+28)) = v259
																		v267 = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v267
																		*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)) = uint8(v267)
																		v271 = int32(1)
																		*(*uint16)(unsafe.Add(mBase, uint32(v14)+57)) = uint16(v271)
																		*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v14 - int32(-64)
																		F_get_rule_expr(m, v254, v14+int32(24), v267)
																		mBase = m.M
																		v280 = m.ExcPending
																		if v280 != 0 {
																			return int32(0)
																		} else {
																			v281 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
																			F_ReleaseCatCache(m, v19)
																			mBase = m.M
																			v283 = m.ExcPending
																			if v283 != 0 {
																				return int32(0)
																			} else {
																				v284 = F_cstring_to_text(m, v281)
																				mBase = m.M
																				v285 = m.ExcPending
																				if v285 != 0 {
																					return int32(0)
																				} else {
																					F_pfree(m, v281)
																					mBase = m.M
																					v287 = m.ExcPending
																					if v287 != 0 {
																						return int32(0)
																					} else {
																						v290 = v284
																						m.G0 = v14 + int32(80)
																						return v290
																					}
																				}
																			}
																		}
																	}
																} else {
																	F_ReleaseCatCache(m, v19)
																	mBase = m.M
																	v247 = m.ExcPending
																	if v247 != 0 {
																		return int32(0)
																	} else {
																		v248 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v248)
																		v290 = v230
																		m.G0 = v14 + int32(80)
																		return v290
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
									F_ReleaseCatCache(m, v19)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										v57 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v57)
										v290 = int32(0)
										m.G0 = v14 + int32(80)
										return v290
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
func F_pg_get_function_arguments(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_SearchSysCache1(m, int32(47), v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			v16 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
			v32 = int32(0)
			m.G0 = v6 + int32(16)
			return v32
		} else {
			F_initStringInfo(m, v6)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v23 = F_print_function_arguments(m, v6, v10, int32(0), int32(1))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_ReleaseCatCache(m, v10)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						v28 = F_cstring_to_text(m, v27)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v27)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v32 = v28
								m.G0 = v6 + int32(16)
								return v32
							}
						}
					}
				}
			}
		}
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v247 int32
	_ = v247
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v330 float32
	_ = v330
	var v333 int32
	_ = v333
	var v338 float32
	_ = v338
	var v348 int32
	_ = v348
	var v349 float32
	_ = v349
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
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
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v623 int32
	_ = v623
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	v14 = m.G0
	v16 = v14 - int32(192)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_initStringInfo(m, v16+int32(160))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = F_SearchSysCache1(m, int32(47), v18)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v16 + int32(192)
	return v864
L4:
	;
	if v26 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v30 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
	v864 = int32(0)
	goto L3
L6:
	;
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
	v35 = v33 + v34
	v37 = v35 + int32(4)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+96)))
	if v38 != int32(97) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	F_appendStringInfoChar(m, v16+int32(160), int32(10))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L1
	} else {
		goto L237
	}
L9:
	;
	v761 = F_SysCacheGetAttrNotNull(m, int32(47), v26, int32(26))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L218
	}
L10:
	;
	v44 = base.B2i32(v38 == int32(112))
	if v38 == int32(112) {
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
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L214
	}
L13:
	;
	v45 = int32(540978)
	goto L15
L14:
	;
	v45 = int32(530797)
	goto L15
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
	v47 = F_get_namespace_name_or_temp(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_initStringInfo(m, v16+int32(176))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v47 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v53 = F_quote_identifier(m, v47)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v65 = F_quote_identifier(m, v37)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v53
	F_appendStringInfo(m, v16+int32(176), int32(608647), v16+int32(144))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	F_appendStringInfoString(m, v16+int32(176), v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v45
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v16)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+132)) = v70
	F_appendStringInfo(m, v16+int32(160), int32(687397), v16+int32(128))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v83 = F_print_function_arguments(m, v16+int32(160), v26, int32(0), int32(1))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_appendStringInfoString(m, v16+int32(160), int32(758724))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v44 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_appendStringInfoString(m, v16+int32(160), int32(746328))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v109 = m.G0
	v111 = v109 - int32(16)
	m.G0 = v111
	v117 = F_SysCacheGetAttr(m, int32(47), v26, int32(25), v111+int32(15))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L35
	}
L31:
	;
	F_print_function_rettype(m, v16+int32(160), v26)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_appendStringInfoChar(m, v16+int32(160), int32(10))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	if int32(0) < v153 {
		goto L56
	} else {
		goto L57
	}
L35:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+15)))
	if v119 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L53
	}
L37:
	;
	v122 = F_pg_detoast_datum(m, v117)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v153 = int32(0)
	goto L39
L39:
	;
	m.G0 = v111 + int32(16)
	goto L34
L40:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v124 != int32(1) {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	if v127 < int32(0) {
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	if v130 != 0 {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	if v131 != int32(26) {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v135 = v127 << (uint(int32(2)) % 32)
	v136 = F_palloc(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(176)))) = v136
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	if v139 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v147 = v139
	goto L48
L47:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v147 = (v140<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L48
L48:
	;
	if v135 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v153 = v127
	goto L39
L50:
	;
	v149 = F__emscripten_memcpy_bulkmem(m, v136, v147+v122, v135)
	mBase = m.M
	goto L52
L51:
	;
	goto L52
L52:
	;
	goto L49
L53:
	;
	F_errmsg_internal(m, int32(152741), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(499202), int32(1500), int32(162341))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_appendStringInfoString(m, v16+int32(160), int32(746649))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v35)+76))
	v263 = F_get_language_name(m, v261, int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L72
	}
L59:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v16)+176))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v181 = F_format_type_be(m, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = v181
	F_appendStringInfo(m, v16+int32(160), int32(198836), v16+int32(112))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v191 = int32(1)
	if v153 != v191 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v194 = v191
	goto L65
L63:
	;
	goto L64
L64:
	;
	F_appendStringInfoChar(m, v16+int32(160), int32(10))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L71
	}
L65:
	;
	F_appendStringInfoString(m, v16+int32(160), int32(748219))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	goto L64
L67:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v16)+176))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v212+v194<<(uint(int32(2))%32))))
	v217 = F_format_type_be(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v217
	F_appendStringInfo(m, v16+int32(160), int32(198836), v16+int32(96))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v228 = v194 + int32(1)
	if v228 != v153 {
		v194 = v228
		goto L65
	} else {
		goto L70
	}
L70:
	;
	goto L66
L71:
	;
	goto L58
L72:
	;
	v265 = F_quote_identifier(m, v263)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v265
	F_appendStringInfo(m, v16+int32(160), int32(750232), v16+int32(80))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v16)+164))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+96)))
	if v276 == int32(119) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_appendStringInfoString(m, v16+int32(160), int32(517674))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+101)))
	switch v285 - int32(105) {
	case 0:
		v289 = int32(542101)
		goto L80
	default:
		goto L79
	case 10:
		goto L81
	}
L78:
	;
	goto L77
L79:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+102)))
	switch v296 - int32(114) {
	case 0:
		goto L85
	case 1:
		v300 = int32(543420)
		goto L84
	default:
		goto L83
	}
L80:
	;
	F_appendStringInfoString(m, v16+int32(160), v289)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L82
	}
L81:
	;
	v289 = int32(542112)
	goto L80
L82:
	;
	goto L79
L83:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+99)))
	if v306 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	F_appendStringInfoString(m, v16+int32(160), v300)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L86
	}
L85:
	;
	v300 = int32(545110)
	goto L84
L86:
	;
	goto L83
L87:
	;
	F_appendStringInfoString(m, v16+int32(160), int32(522876))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+97)))
	if v314 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L89
L91:
	;
	F_appendStringInfoString(m, v16+int32(160), int32(527225))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+98)))
	if v322 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L93
L95:
	;
	F_appendStringInfoString(m, v16+int32(160), int32(538411))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v330 = *(*float32)(unsafe.Add(mBase, uint32(v35)+80))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v35)+76))
	if v333&int32(-2) == int32(12) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L97
L99:
	;
	v338 = float32(1)
	goto L101
L100:
	;
	v338 = float32(100)
	goto L101
L101:
	;
	if base.F32_ne(v330, v338) != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16)+64)) = base.F64_promote_f32(v330)
	F_appendStringInfo(m, v16+int32(160), int32(339099), v16-int32(-64))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v349 = *(*float32)(unsafe.Add(mBase, uint32(v35)+84))
	if base.F32_gt(v349, float32(0)) == int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L104
L106:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v35)+92))
	if v365 != 0 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	if base.F32_eq(v349, float32(1000)) != 0 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16)+48)) = base.F64_promote_f32(v349)
	F_appendStringInfo(m, v16+int32(160), int32(339108), v16+int32(48))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	goto L106
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+176)) = int32(2281)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v35)+92))
	v370 = int32(0)
	v376 = F_generate_function_name(m, v368, int32(1), v370, v16+int32(176), v370, v370, v370)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v16)+164))
	if v386 != v275 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v376
	F_appendStringInfo(m, v16+int32(160), int32(198443), v16+int32(32))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	F_appendStringInfoChar(m, v16+int32(160), int32(10))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v397 = F_SysCacheGetAttr(m, int32(47), v26, int32(29), v16+int32(159))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L119
	}
L118:
	;
	goto L117
L119:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+159)))
	if v399 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v660 = F_SysCacheGetAttr(m, int32(47), v26, int32(28), v16+int32(159))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L189
	}
L121:
	;
	v400 = F_pg_detoast_datum(m, v397)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+176)) = int32(1)
	v405 = v400 + int32(16)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	if v406 <= int32(0) {
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L124
L124:
	;
	v426 = F_array_ref(m, v400, v16+int32(176), v16+int32(159))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L126
	}
L125:
	;
	goto L120
L126:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+159)))
	if v428 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v16)+176))
	v639 = v637 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+176)) = v639
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	if v639 <= v641 {
		goto L124
	} else {
		goto L188
	}
L128:
	;
	v429 = F_text_to_cstring(m, v426)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v431 = int32(61)
	v432 = F___strchrnul(m, v429, v431)
	mBase = m.M
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
	if v434 == v431 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v438 == int32(0) {
		goto L127
	} else {
		goto L134
	}
L131:
	;
	v438 = v432
	goto L133
L132:
	;
	v438 = int32(0)
	goto L133
L133:
	;
	goto L130
L134:
	;
	v441 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v438))) = uint8(v441)
	v443 = F_quote_identifier(m, v429)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v443
	F_appendStringInfo(m, v16+int32(160), int32(746491), v16+int32(16))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v454 = v438 + int32(1)
	v455 = F_GetConfigOptionFlags(m, v429)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	F_appendStringInfoChar(m, v16+int32(160), int32(10))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L187
	}
L138:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L184
	}
L139:
	;
	if v455&int32(2) != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v461 = F_SplitGUCList(m, v454, v16+int32(152))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	F_appendStringInfoChar(m, v16+int32(160), int32(39))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L169
	}
L143:
	;
	if v461 == int32(0) {
		goto L138
	} else {
		goto L144
	}
L144:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v16)+152))
	if v465 == int32(0) {
		goto L137
	} else {
		goto L145
	}
L145:
	;
	v468 = int32(0)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v465)+4))
	if v469 <= v468 {
		goto L137
	} else {
		goto L146
	}
L146:
	;
	v478 = v468
	goto L147
L147:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v465)+12))
	v488 = v485 + v478<<(uint(int32(2))%32)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)))
	F_appendStringInfoChar(m, v16+int32(160), int32(39))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v497 = v489
	goto L150
L150:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
	v509 = base.I32_extend8_s(v508)
	if v508 != int32(39) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	F_appendStringInfoChar(m, v16+int32(160), v509)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L168
	}
L153:
	;
	if v508 != int32(92) {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	goto L155
L155:
	;
	F_appendStringInfoChar(m, v16+int32(160), v509)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L167
	}
L156:
	;
	if v508 != 0 {
		goto L152
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1338])))
	if v541 != 0 {
		goto L152
	} else {
		goto L166
	}
L159:
	;
	F_appendStringInfoChar(m, v16+int32(160), int32(39))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v520 = v488 + int32(4)
	if v520 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v537 = v478 + int32(1)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v465)+4))
	if v537 < v538 {
		v478 = v537
		goto L147
	} else {
		goto L165
	}
L162:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v16)+152))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+12))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	if base.Ui32(v524+v525<<(uint(int32(2))%32)) <= base.Ui32(v520) {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	F_appendStringInfoString(m, v16+int32(160), int32(748219))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	goto L161
L165:
	;
	goto L137
L166:
	;
	goto L155
L167:
	;
	goto L152
L168:
	;
	v497 = v497 + int32(1)
	goto L150
L169:
	;
	v557 = v454
	goto L170
L170:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	v571 = base.I32_extend8_s(v570)
	if v570 != int32(39) {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	F_appendStringInfoChar(m, v16+int32(160), v571)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L183
	}
L173:
	;
	if v570 != int32(92) {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	goto L175
L175:
	;
	F_appendStringInfoChar(m, v16+int32(160), v571)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L182
	}
L176:
	;
	if v570 != 0 {
		goto L172
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v582 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1338])))
	if v582 != 0 {
		goto L172
	} else {
		goto L181
	}
L179:
	;
	F_appendStringInfoChar(m, v16+int32(160), int32(39))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	goto L137
L181:
	;
	goto L175
L182:
	;
	goto L172
L183:
	;
	v557 = v557 + int32(1)
	goto L170
L184:
	;
	F_errmsg_internal(m, int32(291943), int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(495391), int32(3107), int32(340509))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L187:
	;
	goto L127
L188:
	;
	goto L125
L189:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v35)+76))
	if v662 != int32(14) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	F_appendStringInfoString(m, v16+int32(160), int32(746438))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L194
	}
L191:
	;
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+159)))
	if v665 != 0 {
		goto L190
	} else {
		goto L192
	}
L192:
	;
	F_print_function_sqlbody(m, v16+int32(160), v26)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	goto L8
L194:
	;
	v679 = F_SysCacheGetAttr(m, int32(47), v26, int32(27), v16+int32(159))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+159)))
	if v681 != 0 {
		goto L9
	} else {
		goto L196
	}
L196:
	;
	v682 = F_text_to_cstring(m, v679)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_appendStringInfoChar(m, v16+int32(160), int32(39))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v691 = v682
	goto L199
L199:
	;
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v691))))
	v703 = base.I32_extend8_s(v702)
	if v702 != int32(39) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	F_appendStringInfoChar(m, v16+int32(160), v703)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L213
	}
L202:
	;
	if v702 != int32(92) {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	goto L204
L204:
	;
	F_appendStringInfoChar(m, v16+int32(160), v703)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L212
	}
L205:
	;
	if v702 != 0 {
		goto L201
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1338])))
	if v719 != 0 {
		goto L201
	} else {
		goto L211
	}
L208:
	;
	F_appendStringInfoChar(m, v16+int32(160), int32(39))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	F_appendStringInfoString(m, v16+int32(160), int32(748219))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	goto L9
L211:
	;
	goto L204
L212:
	;
	goto L201
L213:
	;
	v691 = v691 + int32(1)
	goto L199
L214:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v37
	F_errmsg(m, int32(254125), v16)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(495391), int32(2956), int32(340509))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L218:
	;
	v763 = F_text_to_cstring(m, v761)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	F_initStringInfo(m, v16+int32(176))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	F_appendStringInfoChar(m, v16+int32(176), int32(36))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	if v38 == int32(112) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v780 = int32(365134)
	goto L224
L223:
	;
	v780 = int32(255053)
	goto L224
L224:
	;
	F_appendStringInfoString(m, v16+int32(176), v780)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v16)+176))
	v784 = F_strstr(m, v763, v783)
	mBase = m.M
	if v784 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	goto L229
L227:
	;
	goto L228
L228:
	;
	F_appendStringInfoChar(m, v16+int32(176), int32(36))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L233
	}
L229:
	;
	F_appendStringInfoChar(m, v16+int32(176), int32(120))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L231
	}
L230:
	;
	goto L228
L231:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v16)+176))
	v804 = F_strstr(m, v763, v803)
	mBase = m.M
	if v804 != 0 {
		goto L229
	} else {
		goto L232
	}
L232:
	;
	goto L230
L233:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v16)+176))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v16)+180))
	F_appendBinaryStringInfo(m, v16+int32(160), v825, v826)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	F_appendStringInfoString(m, v16+int32(160), v763)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v16)+176))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v16)+180))
	F_appendBinaryStringInfo(m, v16+int32(160), v835, v836)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	goto L8
L237:
	;
	F_ReleaseCatCache(m, v26)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v16)+160))
	v860 = F_cstring_to_text(m, v859)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	F_pfree(m, v859)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	v864 = v860
	goto L3
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[455]))
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = v18
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v9 + int32(32)
	return int32(0)
L6:
	;
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(6)))) = uint8(v29)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)) = uint16(v29)
	v39 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(28)))) = v23 + v39
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v9+v39))) = v45
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(20)))) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if v52 == v29 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if v60 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v55 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)) = uint8(v55)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v57 = F_cstring_to_text(m, v52)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v57
	goto L8
L13:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v71 = v68
	v72 = int32(0)
	goto L20
L14:
	;
	v63 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)) = uint8(v63)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v65 = F_cstring_to_text(m, v60)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v65
	goto L13
L18:
	;
	v87 = F_cstring_to_text(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L30
	}
L19:
	;
	if v72 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v74 != int32(47) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v71 = v71 + int32(1)
	v72 = v77
	goto L20
L23:
	;
	if v74 != 0 {
		v77 = v72
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v77 = v71
	goto L22
L26:
	;
	goto L19
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v86 = v82
	goto L18
L28:
	;
	goto L29
L29:
	;
	v84 = v72 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v84
	v86 = v84
	goto L18
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v87
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_tuplestore_putvalues(m, v90, v91, v9+int32(8), v9+int32(4))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v98 != 0 {
		v23 = v98
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
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
			v108 = int32(0)
			m.G0 = v9 + int32(80)
			return v108
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
					v119 = m.ExcPending
					if v119 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
						F_errmsg_internal(m, int32(46584), v9)
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495391), int32(13138), int32(380053))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
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
												F_initStringInfo(m, v9-int32(-64))
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
													*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v9 - int32(-64)
													F_get_rule_expr(m, v12, v9+int32(24), v78)
													mBase = m.M
													v102 = m.ExcPending
													if v102 != 0 {
														return int32(0)
													} else {
														v103 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
														v104 = F_cstring_to_text(m, v103)
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v103)
															mBase = m.M
															v107 = m.ExcPending
															if v107 != 0 {
																return int32(0)
															} else {
																v108 = v104
																m.G0 = v9 + int32(80)
																return v108
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v48 int32
	_ = v48
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
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
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_SearchSysCache1(m, int32(64), v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(80)
	return v187
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
	v187 = v2
	goto L1
L5:
	;
	goto L6
L6:
	;
	v24 = F_heap_attisnull(m, v14, int32(9), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v24 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
	v34 = F_SysCacheGetAttrNotNull(m, int32(64), v14, int32(9))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L12
	}
L11:
	;
	v28 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
	v187 = v2
	goto L1
L12:
	;
	v36 = F_text_to_cstring(m, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v38 = F_stringToNode(m, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	F_pfree(m, v36)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v42 = v30 + v31
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v44 = F_get_rel_name(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L17
	}
L16:
	;
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L40
	}
L17:
	;
	if v44 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v48 = F_palloc0(m, int32(80))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L2
	} else {
		goto L37
	}
L21:
	;
	v51 = F_palloc0(m, int32(136))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+24)) = int32(1)
	v55 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+21)) = uint8(v55)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = v46
	v58 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(101)
	v63 = F_makeAlias(m, v44, v58)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v63
	v67 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v51)+124)) = uint16(v67)
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+20)) = uint8(v69)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v51
	v76 = F_list_make1_impl(m, int32(1), v10+int32(20))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v78 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v78
	*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v76
	F_set_rtable_names(m, v48, v78, v78)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	F_set_simple_column_names(m, v48)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v48
	v94 = F_list_make1_impl(m, int32(1), v10+int32(16))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v96 = int32(0)
	if v38 == v96 {
		v172 = v96
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v99 = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v100 <= v99 {
		v172 = v96
		goto L16
	} else {
		goto L29
	}
L29:
	;
	v103 = v99
	v105 = v96
	goto L30
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110+v103<<(uint(int32(2))%32))))
	F_initStringInfo(m, v10-int32(-64))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L32
	}
L31:
	;
	v172 = v151
	goto L16
L32:
	;
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+56)) = uint8(v119)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v119
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v119
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+59)) = uint8(v119)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v119
	*(*int64)(unsafe.Add(mBase, uint32(v10)+44)) = int64(2)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v10 - int32(-64)
	v137 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+57)) = uint16(v137)
	F_get_rule_expr(m, v114, v10+int32(24), v119)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v145 = F_cstring_to_text(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v151 = F_accumArrayResult(m, v105, v145, int32(0), int32(25), v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v154 = v103 + int32(1)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v154 < v155 {
		v103 = v154
		v105 = v151
		goto L30
	} else {
		goto L36
	}
L36:
	;
	goto L31
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v43
	F_errmsg_internal(m, int32(46584), v10)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(495391), int32(13138), int32(380053))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v181 = F_makeArrayResult(m, v172, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v187 = v181
	goto L1
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
	var v31 int32
	_ = v31
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
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v13)+100)) = v16
	v24 = F_get_call_result_type(m, l0, v2, v13+int32(72))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L6
	} else {
		goto L134
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L6
	} else {
		goto L131
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L6
	} else {
		goto L128
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L6
	} else {
		goto L125
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L122
	}
L6:
	;
	return int32(0)
L7:
	;
	if v24 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v31 = v2
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L6
	} else {
		goto L119
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v31*int32(40))+uint32(_consts[391])))
	if v16 != v44 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v44 == v16 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v47 = v31 + int32(1)
	if v47 != int32(37) {
		v31 = v47
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	goto L15
L17:
	;
	v54 = F_table_open(m, v16, int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L20
	}
L18:
	;
	v324 = v2
	v325 = v2
	goto L19
L19:
	;
	v330 = int32(1)
	v334 = F_getObjectTypeDescription(m, v13+int32(100), v330)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L6
	} else {
		goto L100
	}
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[390]))
	if v57 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v116)+20)))
	v128 = F_get_catalog_object_by_oid_extended(m, v54, v126, v15, int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L43
	}
L22:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v58 == v16 {
		v116 = v57
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v62 = int32(0)
	goto L27
L25:
	;
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, _consts[390])) = v114
	v116 = v114
	goto L21
L27:
	;
	v72 = v62 * int32(40)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+uint32(_consts[391])))
	if v16 != v75 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v114 = v72 + int32(768656)
	goto L26
L29:
	;
	if v62 == int32(36) {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	v82 = (v62 | int32(1)) * int32(40)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[391])))
	if v16 == v85 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v114 = v82 + int32(768656)
	goto L26
L34:
	;
	goto L35
L35:
	;
	v92 = (v62 | int32(2)) * int32(40)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[391])))
	if v16 == v95 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v114 = v92 + int32(768656)
	goto L26
L37:
	;
	goto L38
L38:
	;
	v102 = (v62 | int32(3)) * int32(40)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v102)+uint32(_consts[391])))
	if v16 == v105 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v114 = v102 + int32(768656)
	goto L26
L40:
	;
	v62 = v62 + int32(4)
	goto L27
L42:
	;
	F_sequence_close(m, v54, int32(1))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L6
	} else {
		goto L99
	}
L43:
	;
	if v128 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v311 = v2
	v312 = v2
	goto L42
L45:
	;
	goto L46
L46:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[390]))
	if v134 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v203 = int32(*(*int16)(unsafe.Add(mBase, uint32(v193)+24)))
	if v203 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L48:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	if v135 == v16 {
		v193 = v134
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v139 = int32(0)
	goto L53
L51:
	;
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, _consts[390])) = v191
	v193 = v191
	goto L47
L53:
	;
	v149 = v139 * int32(40)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v149)+uint32(_consts[391])))
	if v16 != v152 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v191 = v149 + int32(768656)
	goto L52
L55:
	;
	if v139 == int32(36) {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	goto L54
L58:
	;
	v159 = (v139 | int32(1)) * int32(40)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v159)+uint32(_consts[391])))
	if v16 == v162 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v191 = v159 + int32(768656)
	goto L52
L60:
	;
	goto L61
L61:
	;
	v169 = (v139 | int32(2)) * int32(40)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v169)+uint32(_consts[391])))
	if v16 == v172 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v191 = v169 + int32(768656)
	goto L52
L63:
	;
	goto L64
L64:
	;
	v179 = (v139 | int32(3)) * int32(40)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v179)+uint32(_consts[391])))
	if v16 == v182 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v191 = v179 + int32(768656)
	goto L52
L66:
	;
	v139 = v139 + int32(4)
	goto L53
L68:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+36)))
	if v291 != int32(1) {
		v311 = v2
		v312 = v286
		goto L42
	} else {
		goto L94
	}
L69:
	;
	v227 = int32(0)
	goto L79
L70:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	if v220 == v16 {
		v281 = v218
		v286 = v219
		goto L68
	} else {
		goto L77
	}
L71:
	;
	v218 = v193
	v219 = v2
	goto L70
L72:
	;
	goto L73
L73:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	v209 = F_heap_getattr_2(m, v128, v203, v206, v13+int32(80))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+80)))
	if v211 == int32(1) {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _consts[390]))
	if v215 == int32(0) {
		v223 = v209
		goto L69
	} else {
		goto L76
	}
L76:
	;
	v218 = v215
	v219 = v209
	goto L70
L77:
	;
	v223 = v219
	goto L69
L78:
	;
	*(*int32)(unsafe.Add(mBase, _consts[390])) = v279
	v281 = v279
	v286 = v223
	goto L68
L79:
	;
	v237 = v227 * int32(40)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v237)+uint32(_consts[391])))
	if v16 != v240 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v279 = v237 + int32(768656)
	goto L78
L81:
	;
	if v227 == int32(36) {
		goto L2
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	goto L80
L84:
	;
	v247 = (v227 | int32(1)) * int32(40)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v247)+uint32(_consts[391])))
	if v16 == v250 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v279 = v247 + int32(768656)
	goto L78
L86:
	;
	goto L87
L87:
	;
	v257 = (v227 | int32(2)) * int32(40)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v257)+uint32(_consts[391])))
	if v16 == v260 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v279 = v257 + int32(768656)
	goto L78
L89:
	;
	goto L90
L90:
	;
	v267 = (v227 | int32(3)) * int32(40)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v267)+uint32(_consts[391])))
	if v16 == v270 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v279 = v267 + int32(768656)
	goto L78
L92:
	;
	v227 = v227 + int32(4)
	goto L79
L94:
	;
	v294 = int32(*(*int16)(unsafe.Add(mBase, uint32(v281)+22)))
	if v294 == int32(0) {
		v311 = v2
		v312 = v286
		goto L42
	} else {
		goto L95
	}
L95:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	v300 = F_heap_getattr_2(m, v128, v294, v297, v13+int32(80))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+80)))
	if v302 == int32(1) {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v305 = F_quote_identifier(m, v300)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	v311 = v305
	v312 = v286
	goto L42
L99:
	;
	v324 = v311
	v325 = v312
	goto L19
L100:
	;
	v336 = F_cstring_to_text(m, v334)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	v338 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)) = uint8(v338)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v336
	v346 = F_getObjectIdentityParts(m, v13+int32(100), v338, v338, int32(1))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	if v325 == int32(0) {
		v360 = v330
		goto L103
	} else {
		goto L104
	}
L103:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+77)) = uint8(v360)
	if v324 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L104:
	;
	if v346 == int32(0) {
		v360 = v330
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v352 = F_get_namespace_name(m, v325)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	v354 = F_quote_identifier(m, v352)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	v356 = F_cstring_to_text(m, v354)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v356
	v360 = int32(0)
	goto L103
L109:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)) = uint8(v382)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v389 = F_heap_form_tuple(m, v384, v13+int32(80), v13+int32(76))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L6
	} else {
		goto L117
	}
L110:
	;
	v377 = F_cstring_to_text(m, v346)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L6
	} else {
		goto L116
	}
L111:
	;
	v371 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+78)) = uint8(v371)
	if v346 == int32(0) {
		v382 = v371
		goto L109
	} else {
		goto L115
	}
L112:
	;
	if v346 == int32(0) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v366 = F_cstring_to_text(m, v324)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	v368 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+78)) = uint8(v368)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v366
	goto L110
L115:
	;
	goto L110
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+92)) = v377
	v382 = int32(0)
	goto L109
L117:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v389)+16))
	v392 = F_HeapTupleHeaderGetDatum(m, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L6
	} else {
		goto L118
	}
L118:
	;
	m.G0 = v13 + int32(112)
	return v392
L119:
	;
	F_errmsg_internal(m, int32(368493), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(495158), int32(4267), int32(110794))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L6
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v16
	F_errmsg_internal(m, int32(59738), v13-int32(-64))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L6
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(495158), int32(2777), int32(505477))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v16
	F_errmsg_internal(m, int32(59738), v13+int32(48))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(495158), int32(2777), int32(505477))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L6
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v16
	F_errmsg_internal(m, int32(467946), v13+int32(32))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L6
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(495158), int32(4290), int32(110794))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L6
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v16
	F_errmsg_internal(m, int32(59738), v13+int32(16))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L6
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(495158), int32(2777), int32(505477))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v16
	F_errmsg_internal(m, int32(467909), v13)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L6
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(495158), int32(4308), int32(110794))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v104 int64
	_ = v104
	var v113 int32
	_ = v113
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
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
	var v151 int32
	_ = v151
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
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v605 int32
	_ = v605
	v13 = m.G0
	v15 = v13 - int32(288)
	m.G0 = v15
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
	m.G0 = v15 + int32(288)
	return v605
L2:
	;
	v605 = int32(0)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(28)
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(5)
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
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v21
	v24 = l3 + int32(50)
	v28 = F_pg_sprintf(m, l3, int32(59790), v15-int32(-64))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(35)
	goto L2
L9:
	;
	return int32(0)
L10:
	;
	v32 = l3 + v28
	if base.Ui32(v24-v32) < base.Ui32(int32(6)) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v28 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v36 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v36)
	v40 = v32 + int32(1)
	goto L14
L13:
	;
	v40 = l3
	goto L14
L14:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v41
	v46 = F_pg_sprintf(m, v40, int32(59790), v15+int32(48))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v48 = v46 + v40
	if base.Ui32(v24-v48) < base.Ui32(int32(6)) {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	if v48 != l3 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v53 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v53)
	v57 = v48 + int32(1)
	goto L19
L18:
	;
	v57 = l3
	goto L19
L19:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v58
	v63 = F_pg_sprintf(m, v57, int32(59790), v15+int32(32))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v65 = v63 + v57
	if base.Ui32(v24-v65) < base.Ui32(int32(6)) {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	if v65 != l3 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v70 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v65))) = uint8(v70)
	v74 = v65 + int32(1)
	goto L24
L23:
	;
	v74 = l3
	goto L24
L24:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v75
	v80 = F_pg_sprintf(m, v74, int32(59790), v15+int32(16))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	if l2 == int32(32) {
		v605 = l3
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v84 = v80 + v74
	if base.Ui32(v24-v84) < base.Ui32(int32(5)) {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l2
	v90 = F_pg_sprintf(m, v84, int32(39502), v15)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v605 = l3
	goto L1
L29:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(28)
	goto L2
L30:
	;
	goto L31
L31:
	;
	v104 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+216)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v15)+208)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v15)+200)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v15)+192)) = v104
	v113 = int32(0)
	goto L32
L32:
	;
	v129 = v15 + int32(192) + v113<<(uint(int32(1))%32)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v131 = v113 + l1
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	v135 = v130 | v132<<(uint(int32(8))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v135
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v135 | v137
	v141 = v113 + int32(2)
	if v141 != int32(16) {
		v113 = v141
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v15)+192))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v15)+196))
	if v145 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L33
L35:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v15)+200))
	if v162 != 0 {
		goto L46
	} else {
		goto L47
	}
L36:
	;
	v146 = int32(-1)
	v147 = int32(0)
	v148 = base.B2i32(v144 == v147)
	if v144 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if v144 != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v151 = v146
	goto L41
L40:
	;
	v151 = v147
	goto L41
L41:
	;
	v158 = v148
	v159 = v146
	v160 = v148
	v161 = v151
	goto L35
L42:
	;
	v156 = int32(1)
	goto L44
L43:
	;
	v156 = int32(2)
	goto L44
L44:
	;
	v158 = v156
	v159 = base.B2i32(v144 != int32(0))
	v160 = int32(0)
	v161 = int32(-1)
	goto L35
L45:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v15)+204))
	if v187 != 0 {
		goto L63
	} else {
		goto L64
	}
L46:
	;
	v163 = int32(-1)
	if v159 == v163 {
		v181 = v161
		v182 = v158
		v184 = v160
		v186 = v163
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v173 = int32(1)
	v177 = base.B2i32(v159 == int32(-1))
	if v159 == int32(-1) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	v169 = base.B2i32(v161 == int32(-1)) | base.B2i32(base.Ui32(v160) < base.Ui32(v158))
	if v169 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v170 = v159
	goto L52
L51:
	;
	v170 = v161
	goto L52
L52:
	;
	if v169 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v171 = v158
	goto L55
L54:
	;
	v171 = v160
	goto L55
L55:
	;
	v181 = v170
	v182 = v158
	v184 = v171
	v186 = int32(-1)
	goto L45
L56:
	;
	v178 = v173
	goto L58
L57:
	;
	v178 = v158 + v173
	goto L58
L58:
	;
	if v159 == int32(-1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v180 = int32(2)
	goto L61
L60:
	;
	v180 = v159
	goto L61
L61:
	;
	v181 = v161
	v182 = v178
	v184 = v160
	v186 = v180
	goto L45
L62:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v15)+208))
	if v212 != 0 {
		goto L80
	} else {
		goto L81
	}
L63:
	;
	v188 = int32(-1)
	if v186 == v188 {
		v206 = v181
		v207 = v182
		v209 = v184
		v211 = v188
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v198 = int32(1)
	v202 = base.B2i32(v186 == int32(-1))
	if v186 == int32(-1) {
		goto L73
	} else {
		goto L74
	}
L66:
	;
	v194 = base.B2i32(v181 == int32(-1)) | base.B2i32(base.Ui32(v184) < base.Ui32(v182))
	if v194 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v195 = v186
	goto L69
L68:
	;
	v195 = v181
	goto L69
L69:
	;
	if v194 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v196 = v182
	goto L72
L71:
	;
	v196 = v184
	goto L72
L72:
	;
	v206 = v195
	v207 = v182
	v209 = v196
	v211 = int32(-1)
	goto L62
L73:
	;
	v203 = v198
	goto L75
L74:
	;
	v203 = v182 + v198
	goto L75
L75:
	;
	if v186 == int32(-1) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v205 = int32(3)
	goto L78
L77:
	;
	v205 = v186
	goto L78
L78:
	;
	v206 = v181
	v207 = v203
	v209 = v184
	v211 = v205
	goto L62
L79:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v15)+212))
	if v237 != 0 {
		goto L97
	} else {
		goto L98
	}
L80:
	;
	v213 = int32(-1)
	if v211 == v213 {
		v231 = v206
		v232 = v207
		v234 = v209
		v236 = v213
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v223 = int32(1)
	v227 = base.B2i32(v211 == int32(-1))
	if v211 == int32(-1) {
		goto L90
	} else {
		goto L91
	}
L83:
	;
	v219 = base.B2i32(v206 == int32(-1)) | base.B2i32(base.Ui32(v209) < base.Ui32(v207))
	if v219 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v220 = v211
	goto L86
L85:
	;
	v220 = v206
	goto L86
L86:
	;
	if v219 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v221 = v207
	goto L89
L88:
	;
	v221 = v209
	goto L89
L89:
	;
	v231 = v220
	v232 = v207
	v234 = v221
	v236 = int32(-1)
	goto L79
L90:
	;
	v228 = v223
	goto L92
L91:
	;
	v228 = v207 + v223
	goto L92
L92:
	;
	if v211 == int32(-1) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v230 = int32(4)
	goto L95
L94:
	;
	v230 = v211
	goto L95
L95:
	;
	v231 = v206
	v232 = v228
	v234 = v209
	v236 = v230
	goto L79
L96:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v15)+216))
	if v262 != 0 {
		goto L114
	} else {
		goto L115
	}
L97:
	;
	v238 = int32(-1)
	if v236 == v238 {
		v256 = v231
		v257 = v232
		v259 = v234
		v261 = v238
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v248 = int32(1)
	v252 = base.B2i32(v236 == int32(-1))
	if v236 == int32(-1) {
		goto L107
	} else {
		goto L108
	}
L100:
	;
	v244 = base.B2i32(v231 == int32(-1)) | base.B2i32(base.Ui32(v234) < base.Ui32(v232))
	if v244 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v245 = v236
	goto L103
L102:
	;
	v245 = v231
	goto L103
L103:
	;
	if v244 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v246 = v232
	goto L106
L105:
	;
	v246 = v234
	goto L106
L106:
	;
	v256 = v245
	v257 = v232
	v259 = v246
	v261 = int32(-1)
	goto L96
L107:
	;
	v253 = v248
	goto L109
L108:
	;
	v253 = v232 + v248
	goto L109
L109:
	;
	if v236 == int32(-1) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v255 = int32(5)
	goto L112
L111:
	;
	v255 = v236
	goto L112
L112:
	;
	v256 = v231
	v257 = v253
	v259 = v234
	v261 = v255
	goto L96
L113:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v15)+220))
	if v287 != 0 {
		goto L131
	} else {
		goto L132
	}
L114:
	;
	v263 = int32(-1)
	if v261 == v263 {
		v281 = v256
		v282 = v257
		v284 = v259
		v286 = v263
		goto L113
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v273 = int32(1)
	v277 = base.B2i32(v261 == int32(-1))
	if v261 == int32(-1) {
		goto L124
	} else {
		goto L125
	}
L117:
	;
	v269 = base.B2i32(v256 == int32(-1)) | base.B2i32(base.Ui32(v259) < base.Ui32(v257))
	if v269 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v270 = v261
	goto L120
L119:
	;
	v270 = v256
	goto L120
L120:
	;
	if v269 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v271 = v257
	goto L123
L122:
	;
	v271 = v259
	goto L123
L123:
	;
	v281 = v270
	v282 = v257
	v284 = v271
	v286 = int32(-1)
	goto L113
L124:
	;
	v278 = v273
	goto L126
L125:
	;
	v278 = v257 + v273
	goto L126
L126:
	;
	if v261 == int32(-1) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v280 = int32(6)
	goto L129
L128:
	;
	v280 = v261
	goto L129
L129:
	;
	v281 = v256
	v282 = v278
	v284 = v259
	v286 = v280
	goto L113
L130:
	;
	if base.Ui32(v313) < base.Ui32(int32(2)) {
		goto L153
	} else {
		goto L154
	}
L131:
	;
	if v286 == int32(-1) {
		v310 = v281
		v313 = v284
		goto L130
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v296 = int32(1)
	v300 = base.B2i32(v286 == int32(-1))
	if v286 == int32(-1) {
		goto L141
	} else {
		goto L142
	}
L134:
	;
	v293 = base.B2i32(v281 == int32(-1)) | base.B2i32(base.Ui32(v284) < base.Ui32(v282))
	if v293 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v294 = v286
	goto L137
L136:
	;
	v294 = v281
	goto L137
L137:
	;
	if v293 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v295 = v282
	goto L140
L139:
	;
	v295 = v284
	goto L140
L140:
	;
	v310 = v294
	v313 = v295
	goto L130
L141:
	;
	v301 = v296
	goto L143
L142:
	;
	v301 = v282 + v296
	goto L143
L143:
	;
	v305 = base.B2i32(v281 == int32(-1)) | base.B2i32(base.Ui32(v284) < base.Ui32(v301))
	if v305 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v306 = v301
	goto L146
L145:
	;
	v306 = v284
	goto L146
L146:
	;
	if v286 == int32(-1) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v308 = int32(7)
	goto L149
L148:
	;
	v308 = v286
	goto L149
L149:
	;
	if v305 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v309 = v308
	goto L152
L151:
	;
	v309 = v281
	goto L152
L152:
	;
	v310 = v309
	v313 = v306
	goto L130
L153:
	;
	v317 = int32(-1)
	goto L155
L154:
	;
	v317 = v310
	goto L155
L155:
	;
	if v310 != int32(-1) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v320 = v317
	goto L158
L157:
	;
	v320 = v310
	goto L158
L158:
	;
	v321 = v320 + v313
	if v320 == int32(-1) {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	v355 = int32(1)
	v359 = v347
	goto L167
L160:
	;
	if v320 != 0 {
		v347 = v15 + int32(224)
		goto L159
	} else {
		goto L166
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v144
	v334 = F_pg_sprintf(m, v15+int32(224), int32(30221), v15+int32(176))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L9
	} else {
		goto L165
	}
L162:
	;
	if int32(0) < v320 {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	if int32(0) < v321 {
		goto L160
	} else {
		goto L164
	}
L164:
	;
	goto L161
L165:
	;
	v347 = v334 + (v15 + int32(224))
	goto L159
L166:
	;
	v341 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+224)) = uint8(v341)
	v347 = v15 + int32(224) | int32(1)
	goto L159
L167:
	;
	if v320 == int32(-1) {
		goto L171
	} else {
		goto L172
	}
L168:
	;
	if v320 == int32(-1) {
		v484 = v474
		goto L200
	} else {
		goto L201
	}
L169:
	;
	goto L168
L170:
	;
	v470 = v355 + int32(1)
	if v470 != int32(8) {
		v355 = v470
		v359 = v468
		goto L167
	} else {
		goto L199
	}
L171:
	;
	v374 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v359))) = uint8(v374)
	v377 = v359 + int32(1)
	if v320 != 0 {
		goto L176
	} else {
		goto L177
	}
L172:
	;
	if v355 < v320 {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	if v321 <= v355 {
		goto L171
	} else {
		goto L174
	}
L174:
	;
	if v355 != v320 {
		v468 = v359
		goto L170
	} else {
		goto L175
	}
L175:
	;
	v370 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v359))) = uint8(v370)
	v468 = v359 + int32(1)
	goto L170
L176:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(192)+v355<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v460
	v465 = F_pg_sprintf(m, v377, int32(30221), v15+int32(160))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L9
	} else {
		goto L198
	}
L177:
	;
	if v355 != int32(6) {
		goto L176
	} else {
		goto L178
	}
L178:
	;
	if v313 == int32(6) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v391 = v15 + int32(274)
	if base.Ui32(v391-v377) < base.Ui32(int32(6)) {
		goto L188
	} else {
		goto L189
	}
L180:
	;
	if base.B2i32(v313 != int32(7)) == int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v15)+220))
	if v384 != int32(1) {
		goto L179
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	if v313 != int32(5) {
		goto L176
	} else {
		goto L185
	}
L184:
	;
	goto L183
L185:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v15)+212))
	if v387 != int32(65535) {
		goto L176
	} else {
		goto L186
	}
L186:
	;
	goto L179
L187:
	;
	v453 = F_strlen(m, v377)
	mBase = m.M
	v474 = v453 + v377
	goto L169
L188:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(35)
	goto L2
L189:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v395
	v400 = F_pg_sprintf(m, v377, int32(59790), v15+int32(144))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L9
	} else {
		goto L190
	}
L190:
	;
	v402 = v400 + v377
	v403 = int32(46)
	*(*uint16)(unsafe.Add(mBase, uint32(v402))) = uint16(v403)
	v406 = v402 + int32(1)
	if base.Ui32(v391-v406) < base.Ui32(int32(6)) {
		goto L188
	} else {
		goto L191
	}
L191:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v410
	v415 = F_pg_sprintf(m, v406, int32(59790), v15+int32(128))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L9
	} else {
		goto L192
	}
L192:
	;
	v417 = v415 + v406
	v418 = int32(46)
	*(*uint16)(unsafe.Add(mBase, uint32(v417))) = uint16(v418)
	v421 = v417 + int32(1)
	if base.Ui32(v391-v421) < base.Ui32(int32(6)) {
		goto L188
	} else {
		goto L193
	}
L193:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v425
	v430 = F_pg_sprintf(m, v421, int32(59790), v15+int32(112))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L9
	} else {
		goto L194
	}
L194:
	;
	v432 = v430 + v421
	v433 = int32(46)
	*(*uint16)(unsafe.Add(mBase, uint32(v432))) = uint16(v433)
	v436 = v432 + int32(1)
	if base.Ui32(v391-v436) < base.Ui32(int32(6)) {
		goto L188
	} else {
		goto L195
	}
L195:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v440
	v445 = F_pg_sprintf(m, v436, int32(59790), v15+int32(96))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L9
	} else {
		goto L196
	}
L196:
	;
	if v445+v436 != v377 {
		goto L187
	} else {
		goto L197
	}
L197:
	;
	goto L188
L198:
	;
	v468 = v465 + v377
	goto L170
L199:
	;
	v474 = v468
	goto L169
L200:
	;
	v485 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v484))) = uint8(v485)
	if l2 == int32(-1) {
		v498 = v484
		goto L203
	} else {
		goto L204
	}
L201:
	;
	if v321 != int32(8) {
		v484 = v474
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v480 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v474))) = uint8(v480)
	v484 = v474 + int32(1)
	goto L200
L203:
	;
	if base.Ui32(int32(50)) < base.Ui32(v498-(v15+int32(224))) {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	if l2 == int32(128) {
		v498 = v484
		goto L203
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = l2
	v495 = F_pg_sprintf(m, v484, int32(39502), v15+int32(80))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L9
	} else {
		goto L206
	}
L206:
	;
	v498 = v495 + v484
	goto L203
L207:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(35)
	goto L2
L208:
	;
	goto L209
L209:
	;
	v508 = v15 + int32(224)
	if (v508^l3)&int32(3) != 0 {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	v605 = l3
	goto L1
L211:
	;
	goto L210
L212:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v563))) = uint8(v562)
	if v562&int32(255) == int32(0) {
		goto L211
	} else {
		goto L227
	}
L213:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508))))
	v561 = v508
	v562 = v514
	v563 = l3
	goto L212
L214:
	;
	goto L215
L215:
	;
	if v508&int32(3) != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v518 = v508
	v520 = l3
	goto L219
L217:
	;
	v532 = v508
	v534 = l3
	goto L218
L218:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v539 = int32(-2139062144)
	if (int32(16843008)-v536|v536)&v539 != v539 {
		v561 = v532
		v562 = v536
		v563 = v534
		goto L212
	} else {
		goto L223
	}
L219:
	;
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518))))
	*(*uint8)(unsafe.Add(mBase, uint32(v520))) = uint8(v521)
	if v521 == int32(0) {
		goto L211
	} else {
		goto L221
	}
L220:
	;
	v532 = v528
	v534 = v526
	goto L218
L221:
	;
	v525 = int32(1)
	v526 = v520 + v525
	v528 = v518 + v525
	if v528&int32(3) != 0 {
		v518 = v528
		v520 = v526
		goto L219
	} else {
		goto L222
	}
L222:
	;
	goto L220
L223:
	;
	v544 = v532
	v545 = v536
	v546 = v534
	goto L224
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v546))) = v545
	v548 = int32(4)
	v549 = v546 + v548
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v544)+4))
	v552 = v544 + v548
	v556 = int32(-2139062144)
	if (v550|(int32(16843008)-v550))&v556 == v556 {
		v544 = v552
		v545 = v550
		v546 = v549
		goto L224
	} else {
		goto L226
	}
L225:
	;
	v561 = v552
	v562 = v550
	v563 = v549
	goto L212
L226:
	;
	goto L225
L227:
	;
	v570 = v561
	v572 = v563
	goto L228
L228:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v572)+1)) = uint8(v573)
	v575 = int32(1)
	if v573 != 0 {
		v570 = v570 + v575
		v572 = v572 + v575
		goto L228
	} else {
		goto L230
	}
L229:
	;
	goto L211
L230:
	;
	goto L229
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
			v18 = *(*int32)(unsafe.Add(mBase, _consts[1308]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v18
			v21 = *(*int64)(unsafe.Add(mBase, _consts[1309]))
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
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v200 int32
	_ = v200
	v6 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l4)+268))
	if v13 <= v6 {
		v200 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v200
L2:
	;
	v17 = l4 + int32(22376)
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v26 = v6
	goto L3
L3:
	;
	v31 = v17 + v26
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v35 == int32(0) {
		v54 = v34
		v55 = v35
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l4)+260))
	if v74 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L5:
	;
	if v55-v54 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	goto L5
L7:
	;
	if v34 != v35 {
		v54 = v34
		v55 = v35
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v39 = l0
	v40 = v31
	goto L9
L9:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if v44 == int32(0) {
		v54 = v43
		v55 = v44
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v54 = v43
	v55 = v44
	goto L6
L11:
	;
	v47 = int32(1)
	if v43 == v44 {
		v39 = v39 + v47
		v40 = v40 + v47
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v58 = v26
	goto L16
L14:
	;
	goto L15
L15:
	;
	goto L4
L16:
	;
	v71 = v58 + int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v17))))
	if v72 != 0 {
		v58 = v71
		goto L16
	} else {
		goto L18
	}
L17:
	;
	if v71 < v13 {
		v26 = v71
		goto L3
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v200 = v6
	goto L1
L20:
	;
	v119 = l4 + int32(16280)
	v121 = l4 + int32(18280)
	v128 = v107
	goto L34
L21:
	;
	v107 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v82 = int32(0)
	v87 = v74
	goto L24
L24:
	;
	v94 = int32(1)
	v95 = (v82 + v87) >> (uint(v94) % 32)
	v101 = *(*int64)(unsafe.Add(mBase, uint32(l4+int32(280)+v95<<(uint(int32(3))%32))))
	v102 = base.B2i32(v18 < v101)
	if v18 < v101 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v107 = v103
	goto L20
L26:
	;
	v103 = v82
	goto L28
L27:
	;
	v103 = v95 + v94
	goto L28
L28:
	;
	if v18 < v101 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v104 = v95
	goto L31
L30:
	;
	v104 = v87
	goto L31
L31:
	;
	if v103 < v104 {
		v82 = v103
		v87 = v104
		goto L24
	} else {
		goto L32
	}
L32:
	;
	goto L25
L33:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v186
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v188
	v200 = int32(1)
	goto L1
L34:
	;
	if int32(0) < v128 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l4)+uint32(_consts[991])))
	v148 = v121 + v145<<(uint(int32(4))%32)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	if v149 == v26 {
		v179 = v148
		goto L33
	} else {
		goto L40
	}
L36:
	;
	v137 = v128 - int32(1)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119+v137))))
	v142 = v121 + v139<<(uint(int32(4))%32)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+8))
	if v143 != v26 {
		v128 = v137
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L35
L39:
	;
	v179 = v142
	goto L33
L40:
	;
	if v74 <= v107 {
		v200 = v6
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v153 = v107
	goto L42
L42:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153+v119))))
	v168 = v121 + v165<<(uint(int32(4))%32)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+8))
	if v169 == v26 {
		v179 = v168
		goto L33
	} else {
		goto L44
	}
L43:
	;
	v200 = v6
	goto L1
L44:
	;
	v172 = v153 + int32(1)
	if v74 != v172 {
		v153 = v172
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
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
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[260]))
	if v6 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v116
L2:
	;
	if v4 == v6 {
		v116 = v2
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
	v9 = *(*int32)(unsafe.Add(mBase, _consts[261]))
	if v9 == v4 {
		v116 = v2
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
		v116 = v2
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v17 = int32(509017)
	goto L12
L10:
	;
	if v54-v55 != 0 {
		goto L24
	} else {
		goto L25
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
	v50 = v17
	v54 = int32(0)
	goto L16
L16:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	goto L10
L17:
	;
	v50 = v45
	v54 = v47
	goto L16
L18:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v28 != v30 {
		v45 = v26
		v47 = v28
		goto L17
	} else {
		goto L20
	}
L19:
	;
	v45 = v39
	v47 = int32(0)
	goto L17
L20:
	;
	if v30 == int32(0) {
		v45 = v26
		v47 = v28
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v35 = v27 - int32(1)
	if v35 == int32(0) {
		v45 = v26
		v47 = v28
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v38 = int32(1)
	v39 = v26 + v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v40 != 0 {
		v25 = v25 + v38
		v26 = v39
		v27 = v35
		v28 = v40
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v63 = int32(509002)
	goto L29
L25:
	;
	v112 = int32(1)
	goto L26
L26:
	;
	F_pfree(m, v11)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L7
	} else {
		goto L41
	}
L27:
	;
	v112 = base.B2i32(v100-v101 == int32(0))
	goto L26
L29:
	;
	goto L30
L30:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v70 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v71 = v11
	v72 = v63
	v73 = int32(14)
	v74 = v70
	goto L35
L32:
	;
	v96 = v63
	v100 = int32(0)
	goto L33
L33:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	goto L27
L34:
	;
	v96 = v91
	v100 = v93
	goto L33
L35:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v74 != v76 {
		v91 = v72
		v93 = v74
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v91 = v85
	v93 = int32(0)
	goto L34
L37:
	;
	if v76 == int32(0) {
		v91 = v72
		v93 = v74
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v81 = v73 - int32(1)
	if v81 == int32(0) {
		v91 = v72
		v93 = v74
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v84 = int32(1)
	v85 = v72 + v84
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v86 != 0 {
		v71 = v71 + v84
		v72 = v85
		v73 = v81
		v74 = v86
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	v116 = v112
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v159 int32
	_ = v159
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
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
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L52
	}
L4:
	;
	m.G0 = v13 + int32(16)
	return v188
L5:
	;
	if v22 == int32(0) {
		v188 = int32(0)
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
	v73 = F_array_contains_nulls(m, v17)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L25
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
	v31 = v26 - int32(16777216)
	if base.Ui32(int32(184549375)) < base.Ui32(v31) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = int32(547298)
	goto L8
L13:
	;
	goto L14
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v31)>>(uint(int32(22))%32))&int32(1020))+uint32(_consts[1256])))
	v43 = v41
	goto L8
L15:
	;
	v46 = int32(89447)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1362])))
	if v50 == int32(0) {
		v69 = v49
		v70 = v50
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v70-v69 != 0 {
		goto L7
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	if v49 != v50 {
		v69 = v49
		v70 = v50
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v54 = v46
	v55 = v43
	goto L20
L20:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	if v59 == int32(0) {
		v69 = v58
		v70 = v59
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v69 = v58
	v70 = v59
	goto L17
L22:
	;
	v62 = int32(1)
	if v58 == v59 {
		v54 = v54 + v62
		v55 = v55 + v62
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v188 = int32(1)
	goto L4
L25:
	;
	if v73 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v79 = F_ArrayGetNItems(m, v76, v17+int32(16))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v83 = F_DirectFunctionCall1Coll(m, int32(1561), int32(0), v15)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v85 = F_pg_detoast_datum(m, v83)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v91 = F_ArrayGetNItems(m, v88, v85+int32(16))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if int32(0) < v91 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v87 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v174 = F_GetSafeSnapshotBlockingPids(m, v15, v13+int32(12), int32(1))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L51
	}
L34:
	;
	v101 = v87
	goto L36
L35:
	;
	v101 = (v88<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L36
L36:
	;
	if v75 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v109 = v75
	goto L39
L38:
	;
	v109 = (v76<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L39
L39:
	;
	v111 = int32(0)
	v116 = v111
	goto L40
L40:
	;
	if v79 <= v111 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L33
L42:
	;
	v159 = v116 + int32(1)
	if v159 != v91 {
		v116 = v159
		goto L40
	} else {
		goto L50
	}
L43:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v85+v101+v116<<(uint(int32(2))%32))))
	v129 = int32(0)
	goto L44
L44:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v17+v109+v129<<(uint(int32(2))%32))))
	if v142 != v127 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v188 = int32(1)
	goto L4
L46:
	;
	v145 = v129 + int32(1)
	if v79 != v145 {
		v129 = v145
		goto L44
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	goto L45
L49:
	;
	goto L42
L50:
	;
	goto L41
L51:
	;
	v188 = base.B2i32(int32(0) < v174)
	goto L4
L52:
	;
	F_errmsg_internal(m, int32(153131), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(496053), int32(66), int32(457573))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
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
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
		if int32(0) <= v7 {
			v45 = v17
			return v45
		} else {
			if base.Ui32(v17) < base.Ui32(int32(2)) {
				v45 = v17
				return v45
			} else {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if base.Ui32(int32(93)) < base.Ui32((v25+int32(95))&int32(255)) {
					return int32(-1)
				} else {
					if v17 != int32(3) {
						v45 = v17
					} else {
						v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
						if base.Ui32(int32(94)) <= base.Ui32((v37+int32(95))&int32(255)) {
							v44 = int32(-1)
						} else {
							v44 = v17
						}
						v45 = v44
					}
					return v45
				}
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
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
				F_ScanKeyInit(m, v9+int32(-48), int32(1), int32(3), int32(184), l0)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v31 = int32(1)
					v35 = F_systable_beginscan(m, v21, int32(2996), v31, l3, v31, v9+int32(-48))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int64(0)
					} else {
						v37 = F_systable_getnext(m, v35)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							if v37 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(67137668))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
										F_errmsg(m, int32(69115), v11)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(499110), int32(3573), int32(86898))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
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
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+22)))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v41+v42)+4))
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
								v49 = F_heap_getattr_2(m, v37, int32(3), v46, v9+int32(-49))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int64(0)
								} else {
									v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
									if v51 == int32(1) {
										v56 = F_acldefault(m, int32(22), v44)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int64(0)
										} else {
											v60 = int32(0)
											v61 = v56
											v63 = F_aclmask(m, v61, l1, v44, l2, int32(1))
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int64(0)
											} else {
												if v61 == int32(0) {
													F_systable_endscan(m, v35)
													mBase = m.M
													v71 = m.ExcPending
													if v71 != 0 {
														return int64(0)
													} else {
														F_sequence_close(m, v21, int32(1))
														mBase = m.M
														v74 = m.ExcPending
														if v74 != 0 {
															return int64(0)
														} else {
															v76 = v63
															m.G0 = v11 - int32(-64)
															return v76
														}
													}
												} else {
													if v60 == v61 {
														F_systable_endscan(m, v35)
														mBase = m.M
														v71 = m.ExcPending
														if v71 != 0 {
															return int64(0)
														} else {
															F_sequence_close(m, v21, int32(1))
															mBase = m.M
															v74 = m.ExcPending
															if v74 != 0 {
																return int64(0)
															} else {
																v76 = v63
																m.G0 = v11 - int32(-64)
																return v76
															}
														}
													} else {
														F_pfree(m, v61)
														mBase = m.M
														v69 = m.ExcPending
														if v69 != 0 {
															return int64(0)
														} else {
															F_systable_endscan(m, v35)
															mBase = m.M
															v71 = m.ExcPending
															if v71 != 0 {
																return int64(0)
															} else {
																F_sequence_close(m, v21, int32(1))
																mBase = m.M
																v74 = m.ExcPending
																if v74 != 0 {
																	return int64(0)
																} else {
																	v76 = v63
																	m.G0 = v11 - int32(-64)
																	return v76
																}
															}
														}
													}
												}
											}
										}
									} else {
										v58 = F_pg_detoast_datum(m, v49)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int64(0)
										} else {
											v60 = v49
											v61 = v58
											v63 = F_aclmask(m, v61, l1, v44, l2, int32(1))
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int64(0)
											} else {
												if v61 == int32(0) {
													F_systable_endscan(m, v35)
													mBase = m.M
													v71 = m.ExcPending
													if v71 != 0 {
														return int64(0)
													} else {
														F_sequence_close(m, v21, int32(1))
														mBase = m.M
														v74 = m.ExcPending
														if v74 != 0 {
															return int64(0)
														} else {
															v76 = v63
															m.G0 = v11 - int32(-64)
															return v76
														}
													}
												} else {
													if v60 == v61 {
														F_systable_endscan(m, v35)
														mBase = m.M
														v71 = m.ExcPending
														if v71 != 0 {
															return int64(0)
														} else {
															F_sequence_close(m, v21, int32(1))
															mBase = m.M
															v74 = m.ExcPending
															if v74 != 0 {
																return int64(0)
															} else {
																v76 = v63
																m.G0 = v11 - int32(-64)
																return v76
															}
														}
													} else {
														F_pfree(m, v61)
														mBase = m.M
														v69 = m.ExcPending
														if v69 != 0 {
															return int64(0)
														} else {
															F_systable_endscan(m, v35)
															mBase = m.M
															v71 = m.ExcPending
															if v71 != 0 {
																return int64(0)
															} else {
																F_sequence_close(m, v21, int32(1))
																mBase = m.M
																v74 = m.ExcPending
																if v74 != 0 {
																	return int64(0)
																} else {
																	v76 = v63
																	m.G0 = v11 - int32(-64)
																	return v76
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
		} else {
			v76 = l2
			m.G0 = v11 - int32(-64)
			return v76
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	v4 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if base.Ui32(v10) <= base.Ui32(int32(41)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v74
L2:
	;
	if v20 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v10*int32(28))+uint32(_consts[1305])))
	v20 = v19
	goto L5
L4:
	;
	v20 = int32(1)
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
		v74 = v4
		goto L1
	} else {
		goto L17
	}
L9:
	;
	v24 = l1
	goto L11
L10:
	;
	v24 = l2
	goto L11
L11:
	;
	if v24 <= int32(0) {
		v74 = v4
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v30 = v4
	goto L13
L13:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v30))))
	if v35 == int32(0) {
		v74 = v30
		goto L1
	} else {
		goto L15
	}
L14:
	;
	return v24
L15:
	;
	v39 = v30 + int32(1)
	if v39 != v24 {
		v30 = v39
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v10*int32(28))+uint32(_consts[1436])))
	v49 = l0
	v50 = l1
	v52 = v4
	goto L18
L18:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v56 == int32(0) {
		v74 = v52
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v74 = v63
	goto L1
L20:
	;
	v59 = m.T0[v48].(func(*base.Module, int32) int32)(m, v49)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	v63 = v59 + v52
	if l2 < v63 {
		v74 = v52
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if l2 == v63 {
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
	v68 = v50 - v59
	if int32(0) < v68 {
		v49 = v49 + v59
		v50 = v68
		v52 = v63
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6*int32(28))+uint32(_consts[1436])))
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
	var v20 int32
	_ = v20
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
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7*int32(28))+uint32(_consts[1305])))
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
		v46 = v3
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
	return v46
L6:
	;
	v18 = l0
	v19 = l1
	v20 = v3
	goto L7
L7:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v22 == int32(0) {
		v46 = v20
		goto L5
	} else {
		goto L9
	}
L8:
	;
	v46 = v39
	goto L5
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27*int32(28))+uint32(_consts[1436])))
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
	v39 = v20 + int32(1)
	v41 = v19 - v33
	if int32(0) < v41 {
		v18 = v18 + v33
		v19 = v41
		v20 = v39
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v13 = int32(4414316)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	*(*int32)(unsafe.Add(mBase, _consts[521])) = int32(0)
	v17 = F___syscall_ret(m, v14)
	mBase = m.M
	goto L1
L1:
	;
	v21 = int32(4414316)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v17 & int32(-193)
	v25 = F___syscall_ret(m, v22)
	mBase = m.M
	goto L2
L2:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v32 = l0 + base.B2i32(v26 == int32(47))
	goto L4
L3:
	;
	v89 = int32(4414316)
	v90 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v17
	v93 = F___syscall_ret(m, v90)
	mBase = m.M
	goto L32
L4:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v36 != int32(47) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v87 = v72 >> (uint(int32(31)) % 32)
	goto L3
L6:
	;
	goto L5
L7:
	;
	v32 = v32 + int32(1)
	goto L4
L8:
	;
	v55 = F___fstatat(m, int32(-100), l0, v9, int32(0))
	mBase = m.M
	goto L17
L9:
	;
	v46 = int32(4414316)
	v47 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v17
	v50 = F___syscall_ret(m, v47)
	mBase = m.M
	goto L15
L10:
	;
	if v36 != 0 {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v41 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v41)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	if v44 != 0 {
		v52 = int32(1)
		goto L8
	} else {
		goto L14
	}
L13:
	;
	v39 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v39)
	goto L9
L14:
	;
	goto L9
L15:
	;
	v52 = int32(0)
	goto L8
L16:
	;
	v79 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v79)
	goto L7
L17:
	;
	if v55 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v58&int32(61440) != int32(16384) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if v52 != 0 {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	if v52 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	if v52 != 0 {
		goto L16
	} else {
		goto L27
	}
L24:
	;
	v66 = int32(54)
	goto L26
L25:
	;
	v66 = int32(20)
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v66
	v87 = int32(-1)
	goto L3
L27:
	;
	v87 = int32(0)
	goto L3
L28:
	;
	v71 = int32(511)
	goto L30
L29:
	;
	v71 = l1
	goto L30
L30:
	;
	v72 = F_mkdir(m, l0, v71)
	mBase = m.M
	v73 = int32(0)
	if v52&base.B2i32(v73 <= v72) == v73 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	goto L16
L32:
	;
	m.G0 = v9 + int32(96)
	return v87
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v4 = int32(1)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32((v5+int32(127))&int32(255)) < base.Ui32(int32(13)) {
		v32 = v4
	} else {
		v13 = v5 & int32(-2)
		if v13&int32(255) == int32(154) {
			v32 = v4
		} else {
			if base.Ui32((v5+int32(112))&int32(255)) < base.Ui32(int32(10)) {
				v32 = int32(2)
			} else {
				if v13&int32(255) == int32(156) {
					v31 = int32(2)
				} else {
					v31 = int32(1)
				}
				v32 = v31
			}
		}
	}
	return v32
}
func F_pg_mule_mblen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32((v5+int32(127))&int32(255)) < base.Ui32(int32(13)) {
		v28 = int32(2)
	} else {
		v12 = int32(3)
		v14 = v5 & int32(254)
		if v14 == int32(154) {
			v28 = v12
		} else {
			if base.Ui32((v5+int32(112))&int32(255)) < base.Ui32(int32(10)) {
				v28 = v12
			} else {
				if v14 == int32(156) {
					v27 = int32(4)
				} else {
					v27 = int32(1)
				}
				v28 = v27
			}
		}
	}
	return v28
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
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1209])))
	if v3 == int32(1) {
		v6 = int32(4444200)
		v7 = int32(0)
		v11 = m.G0
		v13 = v11 - int32(16)
		m.G0 = v13
		v16 = int32(4444216)
		v21 = F___memset(m, int32(4444224), v7, int32(144))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, _consts[1210])) = int32(4)
		*(*int64)(unsafe.Add(mBase, _consts[1211])) = int64(3)
		*(*int32)(unsafe.Add(mBase, _consts[1212])) = int32(2)
		*(*int64)(unsafe.Add(mBase, _consts[1213])) = int64(1)
		v34 = F___memcpy(m, v13, v16, int32(16))
		mBase = m.M
		v35 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13))))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v37 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[1214])) = v37
		*(*int32)(unsafe.Add(mBase, _consts[1215])) = v36
		*(*int64)(unsafe.Add(mBase, _consts[1216])) = v35
		v41 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13)+8)))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
		*(*int32)(unsafe.Add(mBase, _consts[1217])) = v37
		*(*int32)(unsafe.Add(mBase, _consts[1212])) = v42
		*(*int64)(unsafe.Add(mBase, _consts[1213])) = v41
		v49 = F___syscall_ret(m, v7)
		mBase = m.M
		m.G0 = v13 + int32(16)
		F___gettimeofday(m, int32(4444352))
		mBase = m.M
	} else {
	}
	v56 = F_raw_parser(m, l0, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		return int32(0)
	} else {
		v61 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1209])))
		if v61 == int32(1) {
			F_ShowUsage(m, int32(525465))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				return v56
			}
		} else {
			return v56
		}
	}
}
func F_pg_plan_queries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	v5 = int32(0)
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v129
L2:
	;
	v18 = v5
	v19 = v5
	goto L7
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v9 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v129 = v5
	goto L1
L6:
	;
	goto L5
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v19<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v25 == int32(6) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v129 = v117
	goto L1
L9:
	;
	v117 = F_lappend(m, v18, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L13
	} else {
		goto L29
	}
L10:
	;
	v29 = F_palloc0(m, int32(104))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1218])))
	if v46 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	return int32(0)
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(25769804106)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+26)) = uint8(v35)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+88)) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v24)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+92)) = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+96)) = v41
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v43
	v116 = v29
	goto L9
L15:
	;
	v49 = int32(4444200)
	v50 = int32(0)
	v54 = m.G0
	v56 = v54 - int32(16)
	m.G0 = v56
	v59 = int32(4444216)
	v64 = F___memset(m, int32(4444224), v50, int32(144))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[1210])) = int32(4)
	*(*int64)(unsafe.Add(mBase, _consts[1211])) = int64(3)
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = int32(2)
	*(*int64)(unsafe.Add(mBase, _consts[1213])) = int64(1)
	goto L19
L16:
	;
	goto L17
L17:
	;
	v98 = F_planner(m, v24, l1, l2, l3)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L13
	} else {
		goto L22
	}
L18:
	;
	F___gettimeofday(m, int32(4444352))
	mBase = m.M
	goto L17
L19:
	;
	v77 = F___memcpy(m, v56, v59, int32(16))
	mBase = m.M
	v78 = int64(*(*int32)(unsafe.Add(mBase, uint32(v56))))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v80 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1214])) = v80
	*(*int32)(unsafe.Add(mBase, _consts[1215])) = v79
	*(*int64)(unsafe.Add(mBase, _consts[1216])) = v78
	v84 = int64(*(*int32)(unsafe.Add(mBase, uint32(v56)+8)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	*(*int32)(unsafe.Add(mBase, _consts[1217])) = v80
	*(*int32)(unsafe.Add(mBase, _consts[1212])) = v85
	*(*int64)(unsafe.Add(mBase, _consts[1213])) = v84
	goto L21
L21:
	;
	v92 = F___syscall_ret(m, v50)
	mBase = m.M
	m.G0 = v56 + int32(16)
	goto L18
L22:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1218])))
	if v101 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_ShowUsage(m, int32(525483))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L13
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1219])))
	if v108 != int32(1) {
		v116 = v98
		goto L9
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1220])))
	F_elog_node_display(m, int32(284207), v98, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v116 = v98
	goto L9
L29:
	;
	v120 = v19 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v120 < v121 {
		v18 = v117
		v19 = v120
		goto L7
	} else {
		goto L30
	}
L30:
	;
	goto L8
}
func F_pg_re_throw(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v3 = *(*int32)(unsafe.Add(mBase, _consts[426]))
	if v3 != 0 {
		F_pgl_longjmp(m, v3, int32(1))
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
		v8 = *(*int32)(unsafe.Add(mBase, _consts[1410]))
		v10 = v8 * int32(100)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[1415]))) = int32(22)
		v18 = *(*int32)(unsafe.Add(mBase, _consts[224]))
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[1416]))) = uint8(base.B2i32(v18 < int32(23)))
		v25 = *(*int32)(unsafe.Add(mBase, _consts[225]))
		if v25 == int32(2) {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, _consts[226])))
			if v29 == int32(1) {
				v37 = int32(1)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, _consts[227]))
				v37 = base.B2i32(v34 <= int32(22))
			}
			v39 = v37
		} else {
			v39 = int32(0)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[1417]))) = uint8(v39)
		*(*int32)(unsafe.Add(mBase, _consts[394])) = int32(0)
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[1418])))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[1419])))
		v52 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[1420])))
		F_errfinish(m, v46, v49, v52)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return
		} else {
			v55 = m.G0
			v57 = v55 - int32(32)
			m.G0 = v57
			*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = int32(42)
			*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(2048)
			*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = int32(499610)
			*(*int32)(unsafe.Add(mBase, uint32(v57))) = int32(245555)
			F_write_stderr(m, int32(753861), v57)
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, _consts[1421]))
				v72 = F_fflush(m, v71)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
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
	var v24 int32
	_ = v24
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 < int32(0) {
		m.G0 = v8 + int32(16)
		return
	} else {
		if l3 <= int32(0) {
			m.G0 = v8 + int32(16)
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v16 = v14 + int32(20)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if v17 <= l1 {
				m.G0 = v8 + int32(16)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
				F_traverse_lacons(m, v16, l1, v8+int32(12), l2, l3)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+119)))
	switch v9 - int32(83) {
	case 0, 22, 26, 31, 33:
		v42 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int64(0)
		} else {
			F_relation_close(m, l0, int32(1))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int64(0)
			} else {
				m.G0 = v6 + int32(16)
				return base.I64_extend_i32_u(v42)
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
				F_errmsg(m, int32(709654), v6)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30)+119)))
					F_errdetail_relkind_not_supported(m, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(493888), int32(462), int32(301919))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
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
		*(*int32)(unsafe.Add(mBase, _consts[818])) = v31
		*(*int32)(unsafe.Add(mBase, _consts[819])) = v30
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
					F_errmsg(m, int32(270074), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errhint(m, int32(576864), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_errfinish(m, int32(500622), int32(242), int32(263185))
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
							F_errmsg(m, int32(146169), int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_errfinish(m, int32(500622), int32(262), int32(263185))
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
							v23 = *(*int32)(unsafe.Add(mBase, _consts[496]))
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
					*(*int32)(unsafe.Add(mBase, _consts[818])) = v31
					*(*int32)(unsafe.Add(mBase, _consts[819])) = v30
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
	v31 = F_cstring_to_text(m, int32(531981))
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
	v42 = F_cstring_to_text(m, int32(522538))
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
	v56 = F_cstring_to_text(m, int32(534993))
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
	v70 = F_cstring_to_text(m, int32(534981))
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
	v84 = F_cstring_to_text(m, int32(541635))
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
	if v91&int32(16384) != 0 {
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
	v98 = F_cstring_to_text(m, int32(545062))
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32((v2+int32(95))&int32(255)) < base.Ui32(int32(63)) {
		return int32(1)
	} else {
		if base.I32_extend8_s(v2) < int32(0) {
			return int32(2)
		} else {
			v16 = int32(-1)
			v20 = v2 & int32(255)
			if v20 == int32(127) {
				v23 = v16
			} else {
				v23 = int32(1)
			}
			if base.Ui32(v20) < base.Ui32(int32(32)) {
				v26 = v16
			} else {
				v26 = v23
			}
			if v20 != 0 {
				v28 = v26
			} else {
				v28 = int32(0)
			}
			return v28
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	if l1 <= int32(0) {
		v64 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v64 - l0
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
	v64 = v57
	goto L1
L5:
	;
	v57 = v54 + v12
	v58 = v10 - v54
	if int32(0) < v58 {
		v10 = v58
		v12 = v57
		goto L3
	} else {
		goto L24
	}
L6:
	;
	v54 = int32(1)
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
	v64 = v12
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
	if base.Ui32((v15+int32(95))&int32(255)) < base.Ui32(int32(63)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if base.Ui32(v10) < base.Ui32(v26) {
		v64 = v12
		goto L1
	} else {
		goto L23
	}
L15:
	;
	if base.Ui32(v10) < base.Ui32(v26) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(int32(31)) <= base.Ui32((v15+int32(127))&int32(255)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if base.Ui32(int32(28)) < base.Ui32((v15+int32(32))&int32(255)) {
		v64 = v12
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v40 = int32(2)
	v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v41 < int32(-3) {
		v54 = v40
		goto L5
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	if base.Ui32((v41+int32(-64))&int32(255)) < base.Ui32(int32(63)) {
		v54 = v40
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v64 = v12
	goto L1
L23:
	;
	goto L6
L24:
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
	*(*int64)(unsafe.Add(mBase, uint32(v8)+21)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l0
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	switch v11 - int32(98) {
	case 0:
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+16)))
		v36 = int32(0)
		v38 = F_convert_case(m, l0, l1, l2, l3, int32(3), v35, v36, v36)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			v40 = v38
			m.G0 = v9 + int32(16)
			return v40
		}
	case 1:
		v14 = F_strlower_libc(m, l0, l1, l2, l3, l4)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v40 = v14
			m.G0 = v9 + int32(16)
			return v40
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(*(*int8)(unsafe.Add(mBase, uint32(l4))))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v22
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(431007)
			F_errmsg_internal(m, int32(503659), v9)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(500629), int32(1342), int32(431007))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
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
	var v98 int32
	_ = v98
	var v103 int64
	_ = v103
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
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
	v193 = m.ExcPending
	if v193 != 0 {
		goto L8
	} else {
		goto L53
	}
L4:
	;
	v97 = int32(0)
	v98 = base.B2i32(v20 == v97)
	if v15 == v97 {
		v180 = v98
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
	F_errmsg(m, int32(130349), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v37 = int32(526922)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v37
	F_errdetail(m, int32(631541), v11+int32(32))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(496271), int32(258), int32(428238))
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
	F_errmsg(m, int32(130349), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(220735)
	F_errdetail(m, int32(590395), v11+int32(48))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(496271), int32(265), int32(428238))
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
	F_errmsg(m, int32(130349), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = int32(428220)
	F_errdetail(m, int32(589090), v11-int32(-64))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(496271), int32(272), int32(428238))
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
	return v180
L26:
	;
	if v20 != 0 {
		v180 = v98
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v103 = v14 & int64(2147483647)
	v109 = v103
	v110 = int64(100)
	goto L28
L28:
	;
	v113 = F_kill(m, v18, int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L8
	} else {
		goto L30
	}
L29:
	;
	v158 = int32(0)
	v161 = F_errstart(m, int32(19), v158)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L8
	} else {
		goto L49
	}
L30:
	;
	if v113 == int32(-1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v119 == int32(71) {
		v180 = int32(1)
		goto L25
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v109 < v110 {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v18
	F_errmsg(m, int32(296797), v11)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(496271), int32(201), int32(262084))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v139 = v109
	goto L41
L40:
	;
	v139 = v110
	goto L41
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v141 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	v149 = F_WaitLatch(m, v145, int32(41), base.I32_wrap_i64(v139), int32(134217731))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L8
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[507]))
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = int32(0)
	goto L47
L47:
	;
	v155 = v109 - v139
	if int64(0) < v155 {
		v109 = v155
		v110 = v139
		goto L28
	} else {
		goto L48
	}
L48:
	;
	goto L29
L49:
	;
	if v161 == int32(0) {
		v180 = v158
		goto L25
	} else {
		goto L50
	}
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v18
	F_errmsg_plural(m, int32(426261), int32(173065), v15, v11+int32(16))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(496271), int32(221), int32(262084))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	v180 = v158
	goto L25
L53:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(344630), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(496271), int32(249), int32(428238))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	v55 = int32(1847568)
	v56 = int32(1846928)
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
	v48 = v23 + int32(1)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v49 != 0 {
		v23 = v48
		v24 = v49
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
func F_pg_vsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
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
	var v48 int32
	_ = v48
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = v9 + int32(28)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+20)) = int64(0)
	if l1 != 0 {
		v19 = l0
	} else {
		v19 = v9 + int32(7)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v19
	v22 = int32(1)
	if base.Ui32(l1) <= base.Ui32(v22) {
		v25 = v22
	} else {
		v25 = l1
	}
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v19 + v25 - int32(1)
	F_dopr(m, v9+int32(8), l2, l3)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
		v37 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v37)
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
		v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		m.G0 = v9 + int32(32)
		if v41 != 0 {
			v48 = int32(-1)
		} else {
			v48 = v40 + (v36 - v39)
		}
		return v48
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
					F_errmsg_internal(m, int32(368493), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495146), int32(478), int32(277292))
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
