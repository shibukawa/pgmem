package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgstat_backend_flush_cb(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_pgstat_flush_backend(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pgstat_beinit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, _consts[912]))
	v5 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	*(*int32)(unsafe.Add(mBase, _consts[6])) = v3 + v5*int32(408)
	F_on_shmem_exit(m, int32(1210), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		return
	}
}
func F_pgstat_bgwriter_snapshot_cb(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	v15 = *(*int32)(unsafe.Add(mBase, _consts[936]))
	goto L1
L1:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+336))
	v40 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v40 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v57 = v15 + int32(320)
	v59 = F_LWLockAcquire(m, v57, int32(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L10
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(344))))
	*(*int64)(unsafe.Add(mBase, _consts[937])) = v44
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(368))))
	*(*int64)(unsafe.Add(mBase, _consts[938])) = v46
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(360))))
	*(*int64)(unsafe.Add(mBase, _consts[939])) = v48
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(352))))
	*(*int64)(unsafe.Add(mBase, _consts[940])) = v50
	if v38&int32(1) != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	return
L7:
	;
	goto L5
L8:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v15)+336))
	if v38 != v54 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L2
L10:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v15)+392))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v15)+384))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v15)+376))
	F_LWLockRelease(m, v57)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v66 = int32(4349352)
	v68 = *(*int64)(unsafe.Add(mBase, _consts[937]))
	*(*int64)(unsafe.Add(mBase, _consts[937])) = v68 - v63
	v71 = int32(4349360)
	v73 = *(*int64)(unsafe.Add(mBase, _consts[940]))
	*(*int64)(unsafe.Add(mBase, _consts[940])) = v73 - v62
	v76 = int32(4349368)
	v78 = *(*int64)(unsafe.Add(mBase, _consts[939]))
	*(*int64)(unsafe.Add(mBase, _consts[939])) = v78 - v61
	return
}
func F_pgstat_build_snapshot(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[926]))
	if v12 != int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_pgstat_prep_snapshot(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v9 + int32(48)
	return
L4:
	;
	return
L5:
	;
	v21 = m.G0
	v22 = int32(16)
	v23 = v21 - v22
	m.G0 = v23
	F___gettimeofday(m, v23)
	mBase = m.M
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	v27 = int64(*(*int32)(unsafe.Add(mBase, uint32(v23)+8)))
	m.G0 = v23 + v22
	goto L6
L6:
	;
	*(*int64)(unsafe.Add(mBase, _consts[928])) = v27 + v26*int64(1000000) - int64(946684800000000)
	v38 = v9 + int32(20)
	v40 = *(*int32)(unsafe.Add(mBase, _consts[929]))
	v41 = int32(0)
	v42 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+4)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v40
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+24)) = uint8(v41)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+12)) = v42
	goto L7
L7:
	;
	v52 = F_dshash_seq_next(m, v9+int32(20))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	if v52 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v55 = v52
	goto L12
L10:
	;
	goto L11
L11:
	;
	F_dshash_seq_term(m, v9+int32(20))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L44
	}
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v64 = base.B2i32(base.Ui32(int32(11)) < base.Ui32(v60-int32(1)))
	if v64 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L11
L14:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v89 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v88 = v60*int32(72) + int32(1601888)
	goto L14
L16:
	;
	goto L17
L17:
	;
	if base.Ui32(int32(8)) < base.Ui32(v60-int32(24)) {
		v88 = int32(0)
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v76 = int32(0)
	v78 = *(*int32)(unsafe.Add(mBase, _consts[923]))
	if v78 == v76 {
		v88 = v76
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v78+v60<<(uint(int32(2))%32)-int32(96))))
	v88 = v86
	goto L14
L20:
	;
	v176 = F_dshash_seq_next(m, v9+int32(20))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L42
	}
L21:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+16)))
	if v100 != 0 {
		goto L20
	} else {
		goto L25
	}
L22:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v89 == v93 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v95&int32(2) == int32(0) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _consts[930]))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v55)+28))
	v104 = F_dsa_get_address(m, v102, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v55)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v108
	v111 = *(*int32)(unsafe.Add(mBase, _consts[925]))
	v114 = F_pgstat_snapshot_insert(m, v111, v9, v9+int32(19))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[927]))
	if base.Ui32(int32(11)) < base.Ui32(v60-int32(1)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[923]))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v119+v60<<(uint(int32(2))%32)-int32(96))))
	v130 = v125
	goto L30
L29:
	;
	v130 = v60*int32(72) + int32(1601888)
	goto L30
L30:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+20))
	v132 = F_MemoryContextAlloc(m, v117, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+20)) = v132
	v136 = v104 + int32(4)
	v138 = F_LWLockAcquire(m, v136, int32(1))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v114)+20))
	if v64 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)+20))
	if v164 != 0 {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v144 = v60 * int32(72)
	v159 = v144 + int32(1601888)
	v161 = v144 + int32(1601904)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _consts[923]))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v150+v60<<(uint(int32(2))%32)-int32(96))))
	v159 = v156
	v161 = v156 + int32(16)
	goto L33
L37:
	;
	F_LWLockRelease(m, v136)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L41
	}
L38:
	;
	v165 = F__emscripten_memcpy_bulkmem(m, v140, v104+v162, v164)
	mBase = m.M
	goto L40
L39:
	;
	goto L40
L40:
	;
	goto L37
L41:
	;
	goto L20
L42:
	;
	if v176 != 0 {
		v55 = v176
		goto L12
	} else {
		goto L43
	}
L43:
	;
	goto L13
L44:
	;
	v190 = int32(1)
	goto L45
L45:
	;
	v196 = base.B2i32(base.Ui32(int32(12)) < base.Ui32(v190))
	if base.Ui32(int32(12)) < base.Ui32(v190) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, _consts[926])) = int32(2)
	goto L3
L47:
	;
	v270 = v190 + int32(1)
	if v270 != int32(33) {
		v190 = v270
		goto L45
	} else {
		goto L67
	}
L48:
	;
	if base.Ui32(int32(8)) < base.Ui32(v190-int32(24)) {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	v216 = v190*int32(72) + int32(1601888)
	goto L50
L50:
	;
	if v216 == int32(0) {
		goto L47
	} else {
		goto L53
	}
L51:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _consts[923]))
	if v202 == int32(0) {
		goto L47
	} else {
		goto L52
	}
L52:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v202+v190<<(uint(int32(2))%32)-int32(96))))
	v216 = v210
	goto L50
L53:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	if v219&int32(1) == int32(0) {
		goto L47
	} else {
		goto L54
	}
L54:
	;
	if v196 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v252 = v248 + v249
	v254 = *(*int32)(unsafe.Add(mBase, _consts[924]))
	if v254 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	v248 = int32(4349200)
	v249 = v190
	v251 = v190*int32(72) + int32(1601888)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v231 = int32(4401872)
	v234 = v190 - int32(24)
	if base.Ui32(int32(8)) < base.Ui32(v234) {
		v248 = v231
		v249 = v234
		v251 = int32(0)
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v237 = int32(0)
	v239 = *(*int32)(unsafe.Add(mBase, _consts[923]))
	if v239 == v237 {
		v248 = v231
		v249 = v234
		v251 = v237
		goto L55
	} else {
		goto L60
	}
L60:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v239+v190<<(uint(int32(2))%32)-int32(96))))
	v248 = v231
	v249 = v234
	v251 = v247
	goto L55
L61:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v251)+64))
	m.T0[v260].(func(*base.Module))(m)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L66
	}
L62:
	;
	v257 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v252))) = uint8(v257)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	if v259 != 0 {
		goto L47
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	v263 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v252))) = uint8(v263)
	goto L47
L67:
	;
	goto L46
}
func F_pgstat_checkpointer_reset_all_cb(m *base.Module, l0 int64) {
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	v8 = *(*int32)(unsafe.Add(mBase, _consts[936]))
	v10 = v8 + int32(408)
	v12 = F_LWLockAcquire(m, v10, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+424))
	v26 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+512)) = l0
	F_LWLockRelease(m, v10)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L15
	}
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	goto L10
L8:
	;
	goto L7
L9:
	;
	if v24&int32(1) != 0 {
		goto L3
	} else {
		goto L13
	}
L10:
	;
	v30 = F__emscripten_memcpy_bulkmem(m, v8+int32(520), v8+int32(432), int32(88))
	mBase = m.M
	goto L12
L12:
	;
	goto L9
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+424))
	if v24 != v34 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	goto L4
L15:
	;
	return
}
func F_pgstat_fetch_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
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
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v77 int64
	_ = v77
	var v83 int64
	_ = v83
	var v89 int64
	_ = v89
	var v94 int64
	_ = v94
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
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
	var v236 int32
	_ = v236
	var v239 int64
	_ = v239
	var v241 int64
	_ = v241
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int64
	_ = v299
	var v301 int64
	_ = v301
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v320 int32
	_ = v320
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v23 = l0 - int32(1)
	if base.Ui32(v23) <= base.Ui32(int32(11)) {
		v43 = l0*int32(72) + int32(1601888)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_pgstat_prep_snapshot(m)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
		v43 = int32(0)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[923]))
	if v33 == v31 {
		v43 = v31
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33+l0<<(uint(int32(2))%32)-int32(96))))
	v43 = v41
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+56)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = l0
	v52 = *(*int32)(unsafe.Add(mBase, _consts[924]))
	if v52 == int32(2) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_pgstat_build_snapshot(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	v59 = v52
	goto L9
L9:
	;
	if int32(0) < v59 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[924]))
	v59 = v58
	goto L9
L11:
	;
	m.G0 = v16 - int32(-64)
	return v320
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[925]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
	v67 = int64(23)
	v70 = int64(2388976653695081527)
	v71 = (int64(base.Ui64(v66)>>(uint(v67)%64)) ^ v66) * v70
	v72 = int64(47)
	v77 = int64(-8645972361240307355)
	v83 = (int64(base.Ui64(l2)>>(uint(v67)%64)) ^ l2) * v70
	v89 = ((v71^int64(base.Ui64(v71)>>(uint(v72)%64))^int64(-9208349263878056368))*v77 ^ int64(base.Ui64(v83)>>(uint(v72)%64)) ^ v83) * v77
	v94 = (int64(base.Ui64(v89)>>(uint(v67)%64)) ^ v89) * v70
	v102 = v65 & base.I32_wrap_i64(int64(base.Ui64(v94)>>(uint(v72)%64))^v94-int64(base.Ui64(v94)>>(uint(int64(32))%64)))
	v105 = v64 + v102*int32(24)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+16)))
	if v106 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[926])) = v59
	v227 = int32(0)
	v229 = F_pgstat_get_entry_ref(m, l0, l1, l2, v227, v227)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L5
	} else {
		goto L44
	}
L15:
	;
	if v59 == int32(2) {
		v320 = v4
		goto L11
	} else {
		goto L42
	}
L16:
	;
	v113 = v105
	v114 = v102
	goto L17
L17:
	;
	v123 = v14 + int32(-16)
	v124 = int32(16)
	goto L22
L18:
	;
	if v113 == int32(0) {
		goto L15
	} else {
		goto L41
	}
L19:
	;
	if v186 != 0 {
		goto L37
	} else {
		goto L38
	}
L20:
	;
	v186 = int32(0)
	goto L19
L21:
	;
	v160 = v155
	v161 = v156
	v162 = v157
	goto L31
L22:
	;
	if (v113|v123)&int32(3) != 0 {
		v155 = v113
		v156 = v123
		v157 = v124
		goto L21
	} else {
		goto L25
	}
L24:
	;
	if v145 == int32(0) {
		goto L20
	} else {
		goto L30
	}
L25:
	;
	v132 = v113
	v133 = v123
	v134 = v124
	goto L26
L26:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	if v137 != v138 {
		v155 = v132
		v156 = v133
		v157 = v134
		goto L21
	} else {
		goto L28
	}
L27:
	;
	goto L24
L28:
	;
	v140 = int32(4)
	v141 = v133 + v140
	v143 = v132 + v140
	v145 = v134 - v140
	if base.Ui32(int32(3)) < base.Ui32(v145) {
		v132 = v143
		v133 = v141
		v134 = v145
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v155 = v143
	v156 = v141
	v157 = v145
	goto L21
L31:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v165 == v166 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v186 = v165 - v166
	goto L19
L33:
	;
	v168 = int32(1)
	v173 = v162 - v168
	if v173 != 0 {
		v160 = v160 + v168
		v161 = v161 + v168
		v162 = v173
		goto L31
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
	goto L20
L37:
	;
	v189 = (v114 + int32(1)) & v65
	v192 = v64 + v189*int32(24)
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+16)))
	if v193 != 0 {
		v113 = v192
		v114 = v189
		goto L17
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	goto L18
L40:
	;
	goto L15
L41:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v113)+20))
	v320 = v196
	goto L11
L42:
	;
	goto L14
L43:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _consts[924]))
	if v254 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L44:
	;
	if v229 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+16)))
	if v232 != int32(1) {
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _consts[924]))
	if v236 != int32(1) {
		v320 = v4
		goto L11
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v16)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v239
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v241
	v244 = *(*int32)(unsafe.Add(mBase, _consts[925]))
	v249 = F_pgstat_snapshot_insert(m, v244, v14+int32(-56), v14+int32(-17))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v249)+20)) = int32(0)
	v320 = v4
	goto L11
L51:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	v270 = F_LWLockAcquire(m, v266+int32(4), int32(1))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L5
	} else {
		goto L57
	}
L52:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	v258 = F_palloc(m, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _consts[927]))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	v263 = F_MemoryContextAlloc(m, v261, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L5
	} else {
		goto L56
	}
L55:
	;
	v265 = v258
	goto L51
L56:
	;
	v265 = v263
	goto L51
L57:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if base.Ui32(v23) <= base.Ui32(int32(11)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v287 = l0*int32(72) + int32(1601888)
	goto L60
L59:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _consts[923]))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v280+l0<<(uint(int32(2))%32)-int32(96))))
	v287 = v286
	goto L60
L60:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+16))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v290 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	F_pgstat_unlock_entry(m, v229)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L5
	} else {
		goto L65
	}
L62:
	;
	v291 = F__emscripten_memcpy_bulkmem(m, v265, v272+v288, v290)
	mBase = m.M
	v292 = v291
	goto L64
L63:
	;
	v292 = v265
	goto L64
L64:
	;
	goto L61
L65:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _consts[924]))
	if v296 <= int32(0) {
		v320 = v265
		goto L11
	} else {
		goto L66
	}
L66:
	;
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v16)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v299
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v301
	v304 = *(*int32)(unsafe.Add(mBase, _consts[925]))
	v309 = F_pgstat_snapshot_insert(m, v304, v14+int32(-40), v14+int32(-17))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309)+20)) = v292
	v320 = v265
	goto L11
}
func F_pgstat_fetch_stat_numbackends(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	F_pgstat_read_current_status(m)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return int32(0)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[913]))
		return v6
	}
}
func F_pgstat_fetch_stat_tabentry(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v45 int32
	_ = v45
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	v5 = int32(1)
	if l0 <= int32(3591) {
		if l0 <= int32(2670) {
			switch l0 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v73 = v5
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v73 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
					v73 = int32(0)
				} else {
					v73 = v5
				}
			}
		} else {
			v17 = l0 - int32(2671)
			if base.Ui32(int32(27)) < base.Ui32(v17) {
				if base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)) {
					v73 = v5
				} else {
					if base.Ui32(l0-int32(2846)) < base.Ui32(int32(2)) {
						v73 = v5
					} else {
						v73 = int32(0)
					}
				}
			} else {
				if int32(1)<<(uint(v17)%32)&int32(226492515) == int32(0) {
					if base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)) {
						v73 = v5
					} else {
						if base.Ui32(l0-int32(2846)) < base.Ui32(int32(2)) {
							v73 = v5
						} else {
							v73 = int32(0)
						}
					}
				} else {
					v73 = v5
				}
			}
		}
	} else {
		if l0 <= int32(5999) {
			v29 = l0 - int32(4177)
			if base.Ui32(int32(9)) < base.Ui32(v29) {
				if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
					v73 = v5
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
						v73 = int32(0)
					} else {
						v73 = v5
					}
				}
			} else {
				if int32(1)<<(uint(v29)%32)&int32(963) == int32(0) {
					if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
						v73 = v5
					} else {
						if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
							v73 = int32(0)
						} else {
							v73 = v5
						}
					}
				} else {
					v73 = v5
				}
			}
		} else {
			switch l0 - int32(6243) {
			case 0, 1, 2, 3, 4, 59, 60:
				v73 = v5
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v73 = int32(0)
			default:
				if base.Ui32(l0-int32(6000)) < base.Ui32(int32(3)) {
					v73 = v5
				} else {
					v45 = l0 - int32(6100)
					if base.Ui32(int32(15)) < base.Ui32(v45) {
						v73 = int32(0)
					} else {
						if int32(1)<<(uint(v45)%32)&int32(49153) != 0 {
							v73 = v5
						} else {
							v73 = int32(0)
						}
					}
				}
			}
		}
	}
	v77 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v73 != 0 {
		v78 = int32(0)
	} else {
		v78 = v77
	}
	v80 = F_pgstat_fetch_entry(m, int32(2), v78, base.I64_extend_i32_u(l0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		return int32(0)
	} else {
		return v80
	}
}
func F_pgstat_flush_backend(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v75 int64
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v124 int32
	_ = v124
	var v157 int32
	_ = v157
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v242 int32
	_ = v242
	var v243 int64
	_ = v243
	var v244 int32
	_ = v244
	var v246 int64
	_ = v246
	var v249 int32
	_ = v249
	var v250 int64
	_ = v250
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int64
	_ = v262
	var v263 int32
	_ = v263
	var v264 int64
	_ = v264
	var v267 int32
	_ = v267
	var v268 int64
	_ = v268
	var v271 int64
	_ = v271
	var v274 int32
	_ = v274
	var v275 int64
	_ = v275
	var v276 int32
	_ = v276
	var v277 int64
	_ = v277
	var v278 int64
	_ = v278
	var v279 int64
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int64
	_ = v284
	var v285 int64
	_ = v285
	var v289 int32
	_ = v289
	var v290 int64
	_ = v290
	var v291 int64
	_ = v291
	var v295 int32
	_ = v295
	var v296 int64
	_ = v296
	var v297 int64
	_ = v297
	var v299 int64
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int64
	_ = v304
	var v305 int64
	_ = v305
	var v309 int32
	_ = v309
	var v310 int64
	_ = v310
	var v311 int64
	_ = v311
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v317 int64
	_ = v317
	var v319 int64
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int64
	_ = v324
	var v325 int64
	_ = v325
	var v329 int32
	_ = v329
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v335 int32
	_ = v335
	var v336 int64
	_ = v336
	var v337 int64
	_ = v337
	var v339 int64
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int64
	_ = v344
	var v345 int64
	_ = v345
	var v349 int32
	_ = v349
	var v350 int64
	_ = v350
	var v351 int64
	_ = v351
	var v355 int32
	_ = v355
	var v356 int64
	_ = v356
	var v357 int64
	_ = v357
	var v359 int64
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int64
	_ = v364
	var v365 int64
	_ = v365
	var v369 int32
	_ = v369
	var v370 int64
	_ = v370
	var v371 int64
	_ = v371
	var v375 int32
	_ = v375
	var v376 int64
	_ = v376
	var v377 int64
	_ = v377
	var v379 int64
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int64
	_ = v384
	var v385 int64
	_ = v385
	var v389 int32
	_ = v389
	var v390 int64
	_ = v390
	var v391 int64
	_ = v391
	var v395 int32
	_ = v395
	var v396 int64
	_ = v396
	var v397 int64
	_ = v397
	var v399 int64
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int64
	_ = v404
	var v405 int64
	_ = v405
	var v409 int32
	_ = v409
	var v410 int64
	_ = v410
	var v411 int64
	_ = v411
	var v415 int32
	_ = v415
	var v416 int64
	_ = v416
	var v417 int64
	_ = v417
	var v419 int64
	_ = v419
	var v422 int32
	_ = v422
	var v423 int64
	_ = v423
	var v424 int32
	_ = v424
	var v425 int64
	_ = v425
	var v428 int32
	_ = v428
	var v429 int64
	_ = v429
	var v430 int64
	_ = v430
	var v433 int32
	_ = v433
	var v434 int64
	_ = v434
	var v435 int32
	_ = v435
	var v436 int64
	_ = v436
	var v438 int64
	_ = v438
	var v442 int32
	_ = v442
	var v443 int64
	_ = v443
	var v444 int64
	_ = v444
	var v448 int32
	_ = v448
	var v449 int64
	_ = v449
	var v450 int64
	_ = v450
	var v454 int32
	_ = v454
	var v455 int64
	_ = v455
	var v456 int64
	_ = v456
	var v458 int64
	_ = v458
	var v462 int32
	_ = v462
	var v463 int64
	_ = v463
	var v464 int64
	_ = v464
	var v468 int32
	_ = v468
	var v469 int64
	_ = v469
	var v470 int64
	_ = v470
	var v474 int32
	_ = v474
	var v475 int64
	_ = v475
	var v476 int64
	_ = v476
	var v478 int64
	_ = v478
	var v482 int32
	_ = v482
	var v483 int64
	_ = v483
	var v484 int64
	_ = v484
	var v488 int32
	_ = v488
	var v489 int64
	_ = v489
	var v490 int64
	_ = v490
	var v494 int32
	_ = v494
	var v495 int64
	_ = v495
	var v496 int64
	_ = v496
	var v498 int64
	_ = v498
	var v502 int32
	_ = v502
	var v503 int64
	_ = v503
	var v504 int64
	_ = v504
	var v508 int32
	_ = v508
	var v509 int64
	_ = v509
	var v510 int64
	_ = v510
	var v514 int32
	_ = v514
	var v515 int64
	_ = v515
	var v516 int64
	_ = v516
	var v518 int64
	_ = v518
	var v522 int32
	_ = v522
	var v523 int64
	_ = v523
	var v524 int64
	_ = v524
	var v528 int32
	_ = v528
	var v529 int64
	_ = v529
	var v530 int64
	_ = v530
	var v534 int32
	_ = v534
	var v535 int64
	_ = v535
	var v536 int64
	_ = v536
	var v538 int64
	_ = v538
	var v542 int32
	_ = v542
	var v543 int64
	_ = v543
	var v544 int64
	_ = v544
	var v548 int32
	_ = v548
	var v549 int64
	_ = v549
	var v550 int64
	_ = v550
	var v554 int32
	_ = v554
	var v555 int64
	_ = v555
	var v556 int64
	_ = v556
	var v558 int64
	_ = v558
	var v562 int32
	_ = v562
	var v563 int64
	_ = v563
	var v564 int64
	_ = v564
	var v568 int32
	_ = v568
	var v569 int64
	_ = v569
	var v570 int64
	_ = v570
	var v574 int32
	_ = v574
	var v575 int64
	_ = v575
	var v576 int64
	_ = v576
	var v578 int64
	_ = v578
	var v581 int32
	_ = v581
	var v582 int64
	_ = v582
	var v583 int32
	_ = v583
	var v584 int64
	_ = v584
	var v587 int32
	_ = v587
	var v588 int64
	_ = v588
	var v589 int64
	_ = v589
	var v592 int32
	_ = v592
	var v593 int64
	_ = v593
	var v594 int32
	_ = v594
	var v595 int64
	_ = v595
	var v597 int64
	_ = v597
	var v601 int32
	_ = v601
	var v602 int64
	_ = v602
	var v603 int64
	_ = v603
	var v607 int32
	_ = v607
	var v608 int64
	_ = v608
	var v609 int64
	_ = v609
	var v613 int32
	_ = v613
	var v614 int64
	_ = v614
	var v615 int64
	_ = v615
	var v617 int64
	_ = v617
	var v621 int32
	_ = v621
	var v622 int64
	_ = v622
	var v623 int64
	_ = v623
	var v627 int32
	_ = v627
	var v628 int64
	_ = v628
	var v629 int64
	_ = v629
	var v633 int32
	_ = v633
	var v634 int64
	_ = v634
	var v635 int64
	_ = v635
	var v637 int64
	_ = v637
	var v641 int32
	_ = v641
	var v642 int64
	_ = v642
	var v643 int64
	_ = v643
	var v647 int32
	_ = v647
	var v648 int64
	_ = v648
	var v649 int64
	_ = v649
	var v653 int32
	_ = v653
	var v654 int64
	_ = v654
	var v655 int64
	_ = v655
	var v657 int64
	_ = v657
	var v661 int32
	_ = v661
	var v662 int64
	_ = v662
	var v663 int64
	_ = v663
	var v667 int32
	_ = v667
	var v668 int64
	_ = v668
	var v669 int64
	_ = v669
	var v673 int32
	_ = v673
	var v674 int64
	_ = v674
	var v675 int64
	_ = v675
	var v677 int64
	_ = v677
	var v681 int32
	_ = v681
	var v682 int64
	_ = v682
	var v683 int64
	_ = v683
	var v687 int32
	_ = v687
	var v688 int64
	_ = v688
	var v689 int64
	_ = v689
	var v693 int32
	_ = v693
	var v694 int64
	_ = v694
	var v695 int64
	_ = v695
	var v697 int64
	_ = v697
	var v701 int32
	_ = v701
	var v702 int64
	_ = v702
	var v703 int64
	_ = v703
	var v707 int32
	_ = v707
	var v708 int64
	_ = v708
	var v709 int64
	_ = v709
	var v713 int32
	_ = v713
	var v714 int64
	_ = v714
	var v715 int64
	_ = v715
	var v717 int64
	_ = v717
	var v721 int32
	_ = v721
	var v722 int64
	_ = v722
	var v723 int64
	_ = v723
	var v727 int32
	_ = v727
	var v728 int64
	_ = v728
	var v729 int64
	_ = v729
	var v733 int32
	_ = v733
	var v734 int64
	_ = v734
	var v735 int64
	_ = v735
	var v737 int64
	_ = v737
	var v740 int32
	_ = v740
	var v741 int64
	_ = v741
	var v742 int32
	_ = v742
	var v743 int64
	_ = v743
	var v746 int32
	_ = v746
	var v747 int64
	_ = v747
	var v748 int64
	_ = v748
	var v751 int32
	_ = v751
	var v752 int64
	_ = v752
	var v753 int32
	_ = v753
	var v754 int64
	_ = v754
	var v756 int64
	_ = v756
	var v760 int32
	_ = v760
	var v761 int64
	_ = v761
	var v762 int64
	_ = v762
	var v766 int32
	_ = v766
	var v767 int64
	_ = v767
	var v768 int64
	_ = v768
	var v772 int32
	_ = v772
	var v773 int64
	_ = v773
	var v774 int64
	_ = v774
	var v776 int64
	_ = v776
	var v780 int32
	_ = v780
	var v781 int64
	_ = v781
	var v782 int64
	_ = v782
	var v786 int32
	_ = v786
	var v787 int64
	_ = v787
	var v788 int64
	_ = v788
	var v792 int32
	_ = v792
	var v793 int64
	_ = v793
	var v794 int64
	_ = v794
	var v796 int64
	_ = v796
	var v800 int32
	_ = v800
	var v801 int64
	_ = v801
	var v802 int64
	_ = v802
	var v806 int32
	_ = v806
	var v807 int64
	_ = v807
	var v808 int64
	_ = v808
	var v812 int32
	_ = v812
	var v813 int64
	_ = v813
	var v814 int64
	_ = v814
	var v816 int64
	_ = v816
	var v820 int32
	_ = v820
	var v821 int64
	_ = v821
	var v822 int64
	_ = v822
	var v826 int32
	_ = v826
	var v827 int64
	_ = v827
	var v828 int64
	_ = v828
	var v832 int32
	_ = v832
	var v833 int64
	_ = v833
	var v834 int64
	_ = v834
	var v836 int64
	_ = v836
	var v840 int32
	_ = v840
	var v841 int64
	_ = v841
	var v842 int64
	_ = v842
	var v846 int32
	_ = v846
	var v847 int64
	_ = v847
	var v848 int64
	_ = v848
	var v852 int32
	_ = v852
	var v853 int64
	_ = v853
	var v854 int64
	_ = v854
	var v856 int64
	_ = v856
	var v860 int32
	_ = v860
	var v861 int64
	_ = v861
	var v862 int64
	_ = v862
	var v866 int32
	_ = v866
	var v867 int64
	_ = v867
	var v868 int64
	_ = v868
	var v872 int32
	_ = v872
	var v873 int64
	_ = v873
	var v874 int64
	_ = v874
	var v876 int64
	_ = v876
	var v880 int32
	_ = v880
	var v881 int64
	_ = v881
	var v882 int64
	_ = v882
	var v886 int32
	_ = v886
	var v887 int64
	_ = v887
	var v888 int64
	_ = v888
	var v892 int32
	_ = v892
	var v893 int64
	_ = v893
	var v894 int64
	_ = v894
	var v896 int64
	_ = v896
	var v900 int32
	_ = v900
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v953 int64
	_ = v953
	var v962 int64
	_ = v962
	var v964 int64
	_ = v964
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int64
	_ = v968
	var v970 int64
	_ = v970
	var v971 int64
	_ = v971
	var v975 int64
	_ = v975
	var v977 int64
	_ = v977
	var v978 int64
	_ = v978
	var v982 int64
	_ = v982
	var v984 int64
	_ = v984
	var v985 int64
	_ = v985
	var v989 int64
	_ = v989
	var v991 int64
	_ = v991
	var v992 int64
	_ = v992
	var v996 int64
	_ = v996
	var v997 int64
	_ = v997
	var v1000 int64
	_ = v1000
	var v1001 int64
	_ = v1001
	var v1004 int64
	_ = v1004
	var v1005 int64
	_ = v1005
	var v1008 int64
	_ = v1008
	var v1009 int64
	_ = v1009
	var v1014 int64
	_ = v1014
	var v1018 int64
	_ = v1018
	var v1022 int64
	_ = v1022
	var v1026 int64
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1035 int32
	_ = v1035
	v3 = int32(0)
	v41 = m.G0
	v43 = v41 - int32(2880)
	m.G0 = v43
	v46 = *(*int32)(unsafe.Add(mBase, _consts[407]))
	if base.Ui32(int32(16)) < base.Ui32(v46) {
		v1035 = v3
		m.G0 = v43 + int32(2880)
		return v1035
	} else {
		if int32(1)<<(uint(v46)%32)&int32(115186) == int32(0) {
			v1035 = v3
			m.G0 = v43 + int32(2880)
			return v1035
		} else {
			v56 = int32(*(*uint8)(unsafe.Add(mBase, _consts[738])))
			v57 = l1 & v56
			v59 = l1 & int32(2)
			if v59 != 0 {
				v61 = *(*int64)(unsafe.Add(mBase, _consts[18]))
				v63 = *(*int64)(unsafe.Add(mBase, _consts[932]))
				if (base.B2i32(v61 != v63)|v57)&int32(1) != 0 {
					v75 = int64(*(*int32)(unsafe.Add(mBase, _consts[126])))
					v76 = F_pgstat_get_entry_ref_locked(m, int32(6), int32(0), v75, l0)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						if v76 == int32(0) {
							v1035 = int32(1)
							m.G0 = v43 + int32(2880)
							return v1035
						} else {
							if l1&int32(1) == int32(0) {
							} else {
								v88 = int32(*(*uint8)(unsafe.Add(mBase, _consts[738])))
								if v88 == int32(0) {
								} else {
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
									v124 = F__emscripten_memcpy_bulkmem(m, v43, int32(4401960), int32(2880))
									mBase = m.M
									v157 = v3
									for {
										v188 = v157 * int32(320)
										v195 = int32(0)
										for {
											v235 = v195 << (uint(int32(3)) % 32)
											v236 = v188 + (v91 + int32(992)) + v235
											v237 = *(*int64)(unsafe.Add(mBase, uint32(v236)))
											v239 = *(*int64)(unsafe.Add(mBase, uint32(v235+(v124+int32(960)+v188))))
											*(*int64)(unsafe.Add(mBase, uint32(v236))) = v237 + v239
											v242 = v188 + (v91 + int32(32)) + v235
											v243 = *(*int64)(unsafe.Add(mBase, uint32(v242)))
											v244 = v188 + v124
											v246 = *(*int64)(unsafe.Add(mBase, uint32(v244+v235)))
											*(*int64)(unsafe.Add(mBase, uint32(v242))) = v243 + v246
											v249 = v235 + (v188 + (v91 + int32(1952)))
											v250 = *(*int64)(unsafe.Add(mBase, uint32(v249)))
											v252 = *(*int64)(unsafe.Add(mBase, uint32(v235+(v188+(v124+int32(1920))))))
											v254 = base.I64_div_s(v252, int64(1000))
											*(*int64)(unsafe.Add(mBase, uint32(v249))) = v250 + v254
											v258 = v195 + int32(1)
											if v258 != int32(8) {
												v195 = v258
												continue
											} else {
												break
											}
											break
										}
										v261 = v188 + (v91 + int32(1056))
										v262 = *(*int64)(unsafe.Add(mBase, uint32(v261)))
										v263 = v188 + (v124 + int32(1024))
										v264 = *(*int64)(unsafe.Add(mBase, uint32(v263)))
										*(*int64)(unsafe.Add(mBase, uint32(v261))) = v262 + v264
										v267 = v188 + (v91 + int32(96))
										v268 = *(*int64)(unsafe.Add(mBase, uint32(v267)))
										v271 = *(*int64)(unsafe.Add(mBase, uint32(v244-int32(-64))))
										*(*int64)(unsafe.Add(mBase, uint32(v267))) = v268 + v271
										v274 = v188 + (v91 + int32(2016))
										v275 = *(*int64)(unsafe.Add(mBase, uint32(v274)))
										v276 = v188 + (v124 + int32(1984))
										v277 = *(*int64)(unsafe.Add(mBase, uint32(v276)))
										v278 = int64(1000)
										v279 = base.I64_div_s(v277, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v274))) = v275 + v279
										v282 = int32(8)
										v283 = v261 + v282
										v284 = *(*int64)(unsafe.Add(mBase, uint32(v283)))
										v285 = *(*int64)(unsafe.Add(mBase, uint32(v263)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v283))) = v284 + v285
										v289 = v267 + v282
										v290 = *(*int64)(unsafe.Add(mBase, uint32(v289)))
										v291 = *(*int64)(unsafe.Add(mBase, uint32(v244)+72))
										*(*int64)(unsafe.Add(mBase, uint32(v289))) = v290 + v291
										v295 = v274 + v282
										v296 = *(*int64)(unsafe.Add(mBase, uint32(v295)))
										v297 = *(*int64)(unsafe.Add(mBase, uint32(v276)+8))
										v299 = base.I64_div_s(v297, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v295))) = v296 + v299
										v302 = int32(16)
										v303 = v261 + v302
										v304 = *(*int64)(unsafe.Add(mBase, uint32(v303)))
										v305 = *(*int64)(unsafe.Add(mBase, uint32(v263)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v303))) = v304 + v305
										v309 = v267 + v302
										v310 = *(*int64)(unsafe.Add(mBase, uint32(v309)))
										v311 = *(*int64)(unsafe.Add(mBase, uint32(v244)+80))
										*(*int64)(unsafe.Add(mBase, uint32(v309))) = v310 + v311
										v315 = v274 + v302
										v316 = *(*int64)(unsafe.Add(mBase, uint32(v315)))
										v317 = *(*int64)(unsafe.Add(mBase, uint32(v276)+16))
										v319 = base.I64_div_s(v317, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v315))) = v316 + v319
										v322 = int32(24)
										v323 = v261 + v322
										v324 = *(*int64)(unsafe.Add(mBase, uint32(v323)))
										v325 = *(*int64)(unsafe.Add(mBase, uint32(v263)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v323))) = v324 + v325
										v329 = v267 + v322
										v330 = *(*int64)(unsafe.Add(mBase, uint32(v329)))
										v331 = *(*int64)(unsafe.Add(mBase, uint32(v244)+88))
										*(*int64)(unsafe.Add(mBase, uint32(v329))) = v330 + v331
										v335 = v274 + v322
										v336 = *(*int64)(unsafe.Add(mBase, uint32(v335)))
										v337 = *(*int64)(unsafe.Add(mBase, uint32(v276)+24))
										v339 = base.I64_div_s(v337, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v335))) = v336 + v339
										v342 = int32(32)
										v343 = v261 + v342
										v344 = *(*int64)(unsafe.Add(mBase, uint32(v343)))
										v345 = *(*int64)(unsafe.Add(mBase, uint32(v263)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v343))) = v344 + v345
										v349 = v267 + v342
										v350 = *(*int64)(unsafe.Add(mBase, uint32(v349)))
										v351 = *(*int64)(unsafe.Add(mBase, uint32(v244)+96))
										*(*int64)(unsafe.Add(mBase, uint32(v349))) = v350 + v351
										v355 = v274 + v342
										v356 = *(*int64)(unsafe.Add(mBase, uint32(v355)))
										v357 = *(*int64)(unsafe.Add(mBase, uint32(v276)+32))
										v359 = base.I64_div_s(v357, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v355))) = v356 + v359
										v362 = int32(40)
										v363 = v261 + v362
										v364 = *(*int64)(unsafe.Add(mBase, uint32(v363)))
										v365 = *(*int64)(unsafe.Add(mBase, uint32(v263)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v363))) = v364 + v365
										v369 = v267 + v362
										v370 = *(*int64)(unsafe.Add(mBase, uint32(v369)))
										v371 = *(*int64)(unsafe.Add(mBase, uint32(v244)+104))
										*(*int64)(unsafe.Add(mBase, uint32(v369))) = v370 + v371
										v375 = v274 + v362
										v376 = *(*int64)(unsafe.Add(mBase, uint32(v375)))
										v377 = *(*int64)(unsafe.Add(mBase, uint32(v276)+40))
										v379 = base.I64_div_s(v377, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v375))) = v376 + v379
										v382 = int32(48)
										v383 = v261 + v382
										v384 = *(*int64)(unsafe.Add(mBase, uint32(v383)))
										v385 = *(*int64)(unsafe.Add(mBase, uint32(v263)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v383))) = v384 + v385
										v389 = v267 + v382
										v390 = *(*int64)(unsafe.Add(mBase, uint32(v389)))
										v391 = *(*int64)(unsafe.Add(mBase, uint32(v244)+112))
										*(*int64)(unsafe.Add(mBase, uint32(v389))) = v390 + v391
										v395 = v274 + v382
										v396 = *(*int64)(unsafe.Add(mBase, uint32(v395)))
										v397 = *(*int64)(unsafe.Add(mBase, uint32(v276)+48))
										v399 = base.I64_div_s(v397, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v395))) = v396 + v399
										v402 = int32(56)
										v403 = v261 + v402
										v404 = *(*int64)(unsafe.Add(mBase, uint32(v403)))
										v405 = *(*int64)(unsafe.Add(mBase, uint32(v263)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v403))) = v404 + v405
										v409 = v267 + v402
										v410 = *(*int64)(unsafe.Add(mBase, uint32(v409)))
										v411 = *(*int64)(unsafe.Add(mBase, uint32(v244)+120))
										*(*int64)(unsafe.Add(mBase, uint32(v409))) = v410 + v411
										v415 = v274 + v402
										v416 = *(*int64)(unsafe.Add(mBase, uint32(v415)))
										v417 = *(*int64)(unsafe.Add(mBase, uint32(v276)+56))
										v419 = base.I64_div_s(v417, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v415))) = v416 + v419
										v422 = v188 + (v91 + int32(1120))
										v423 = *(*int64)(unsafe.Add(mBase, uint32(v422)))
										v424 = v188 + (v124 + int32(1088))
										v425 = *(*int64)(unsafe.Add(mBase, uint32(v424)))
										*(*int64)(unsafe.Add(mBase, uint32(v422))) = v423 + v425
										v428 = v188 + (v91 + int32(160))
										v429 = *(*int64)(unsafe.Add(mBase, uint32(v428)))
										v430 = *(*int64)(unsafe.Add(mBase, uint32(v244)+128))
										*(*int64)(unsafe.Add(mBase, uint32(v428))) = v429 + v430
										v433 = v188 + (v91 + int32(2080))
										v434 = *(*int64)(unsafe.Add(mBase, uint32(v433)))
										v435 = v188 + (v124 + int32(2048))
										v436 = *(*int64)(unsafe.Add(mBase, uint32(v435)))
										v438 = base.I64_div_s(v436, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v433))) = v434 + v438
										v442 = v422 + v282
										v443 = *(*int64)(unsafe.Add(mBase, uint32(v442)))
										v444 = *(*int64)(unsafe.Add(mBase, uint32(v424)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v442))) = v443 + v444
										v448 = v428 + v282
										v449 = *(*int64)(unsafe.Add(mBase, uint32(v448)))
										v450 = *(*int64)(unsafe.Add(mBase, uint32(v244)+136))
										*(*int64)(unsafe.Add(mBase, uint32(v448))) = v449 + v450
										v454 = v433 + v282
										v455 = *(*int64)(unsafe.Add(mBase, uint32(v454)))
										v456 = *(*int64)(unsafe.Add(mBase, uint32(v435)+8))
										v458 = base.I64_div_s(v456, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v454))) = v455 + v458
										v462 = v422 + v302
										v463 = *(*int64)(unsafe.Add(mBase, uint32(v462)))
										v464 = *(*int64)(unsafe.Add(mBase, uint32(v424)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v462))) = v463 + v464
										v468 = v428 + v302
										v469 = *(*int64)(unsafe.Add(mBase, uint32(v468)))
										v470 = *(*int64)(unsafe.Add(mBase, uint32(v244)+144))
										*(*int64)(unsafe.Add(mBase, uint32(v468))) = v469 + v470
										v474 = v433 + v302
										v475 = *(*int64)(unsafe.Add(mBase, uint32(v474)))
										v476 = *(*int64)(unsafe.Add(mBase, uint32(v435)+16))
										v478 = base.I64_div_s(v476, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v474))) = v475 + v478
										v482 = v422 + v322
										v483 = *(*int64)(unsafe.Add(mBase, uint32(v482)))
										v484 = *(*int64)(unsafe.Add(mBase, uint32(v424)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v482))) = v483 + v484
										v488 = v428 + v322
										v489 = *(*int64)(unsafe.Add(mBase, uint32(v488)))
										v490 = *(*int64)(unsafe.Add(mBase, uint32(v244)+152))
										*(*int64)(unsafe.Add(mBase, uint32(v488))) = v489 + v490
										v494 = v433 + v322
										v495 = *(*int64)(unsafe.Add(mBase, uint32(v494)))
										v496 = *(*int64)(unsafe.Add(mBase, uint32(v435)+24))
										v498 = base.I64_div_s(v496, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v494))) = v495 + v498
										v502 = v422 + v342
										v503 = *(*int64)(unsafe.Add(mBase, uint32(v502)))
										v504 = *(*int64)(unsafe.Add(mBase, uint32(v424)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v502))) = v503 + v504
										v508 = v428 + v342
										v509 = *(*int64)(unsafe.Add(mBase, uint32(v508)))
										v510 = *(*int64)(unsafe.Add(mBase, uint32(v244)+160))
										*(*int64)(unsafe.Add(mBase, uint32(v508))) = v509 + v510
										v514 = v433 + v342
										v515 = *(*int64)(unsafe.Add(mBase, uint32(v514)))
										v516 = *(*int64)(unsafe.Add(mBase, uint32(v435)+32))
										v518 = base.I64_div_s(v516, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v514))) = v515 + v518
										v522 = v422 + v362
										v523 = *(*int64)(unsafe.Add(mBase, uint32(v522)))
										v524 = *(*int64)(unsafe.Add(mBase, uint32(v424)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v522))) = v523 + v524
										v528 = v428 + v362
										v529 = *(*int64)(unsafe.Add(mBase, uint32(v528)))
										v530 = *(*int64)(unsafe.Add(mBase, uint32(v244)+168))
										*(*int64)(unsafe.Add(mBase, uint32(v528))) = v529 + v530
										v534 = v433 + v362
										v535 = *(*int64)(unsafe.Add(mBase, uint32(v534)))
										v536 = *(*int64)(unsafe.Add(mBase, uint32(v435)+40))
										v538 = base.I64_div_s(v536, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v534))) = v535 + v538
										v542 = v422 + v382
										v543 = *(*int64)(unsafe.Add(mBase, uint32(v542)))
										v544 = *(*int64)(unsafe.Add(mBase, uint32(v424)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v542))) = v543 + v544
										v548 = v428 + v382
										v549 = *(*int64)(unsafe.Add(mBase, uint32(v548)))
										v550 = *(*int64)(unsafe.Add(mBase, uint32(v244)+176))
										*(*int64)(unsafe.Add(mBase, uint32(v548))) = v549 + v550
										v554 = v433 + v382
										v555 = *(*int64)(unsafe.Add(mBase, uint32(v554)))
										v556 = *(*int64)(unsafe.Add(mBase, uint32(v435)+48))
										v558 = base.I64_div_s(v556, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v554))) = v555 + v558
										v562 = v422 + v402
										v563 = *(*int64)(unsafe.Add(mBase, uint32(v562)))
										v564 = *(*int64)(unsafe.Add(mBase, uint32(v424)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v562))) = v563 + v564
										v568 = v428 + v402
										v569 = *(*int64)(unsafe.Add(mBase, uint32(v568)))
										v570 = *(*int64)(unsafe.Add(mBase, uint32(v244)+184))
										*(*int64)(unsafe.Add(mBase, uint32(v568))) = v569 + v570
										v574 = v433 + v402
										v575 = *(*int64)(unsafe.Add(mBase, uint32(v574)))
										v576 = *(*int64)(unsafe.Add(mBase, uint32(v435)+56))
										v578 = base.I64_div_s(v576, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v574))) = v575 + v578
										v581 = v188 + (v91 + int32(1184))
										v582 = *(*int64)(unsafe.Add(mBase, uint32(v581)))
										v583 = v188 + (v124 + int32(1152))
										v584 = *(*int64)(unsafe.Add(mBase, uint32(v583)))
										*(*int64)(unsafe.Add(mBase, uint32(v581))) = v582 + v584
										v587 = v188 + (v91 + int32(224))
										v588 = *(*int64)(unsafe.Add(mBase, uint32(v587)))
										v589 = *(*int64)(unsafe.Add(mBase, uint32(v244)+192))
										*(*int64)(unsafe.Add(mBase, uint32(v587))) = v588 + v589
										v592 = v188 + (v91 + int32(2144))
										v593 = *(*int64)(unsafe.Add(mBase, uint32(v592)))
										v594 = v188 + (v124 + int32(2112))
										v595 = *(*int64)(unsafe.Add(mBase, uint32(v594)))
										v597 = base.I64_div_s(v595, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v592))) = v593 + v597
										v601 = v581 + v282
										v602 = *(*int64)(unsafe.Add(mBase, uint32(v601)))
										v603 = *(*int64)(unsafe.Add(mBase, uint32(v583)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v601))) = v602 + v603
										v607 = v587 + v282
										v608 = *(*int64)(unsafe.Add(mBase, uint32(v607)))
										v609 = *(*int64)(unsafe.Add(mBase, uint32(v244)+200))
										*(*int64)(unsafe.Add(mBase, uint32(v607))) = v608 + v609
										v613 = v592 + v282
										v614 = *(*int64)(unsafe.Add(mBase, uint32(v613)))
										v615 = *(*int64)(unsafe.Add(mBase, uint32(v594)+8))
										v617 = base.I64_div_s(v615, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v613))) = v614 + v617
										v621 = v581 + v302
										v622 = *(*int64)(unsafe.Add(mBase, uint32(v621)))
										v623 = *(*int64)(unsafe.Add(mBase, uint32(v583)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v621))) = v622 + v623
										v627 = v587 + v302
										v628 = *(*int64)(unsafe.Add(mBase, uint32(v627)))
										v629 = *(*int64)(unsafe.Add(mBase, uint32(v244)+208))
										*(*int64)(unsafe.Add(mBase, uint32(v627))) = v628 + v629
										v633 = v592 + v302
										v634 = *(*int64)(unsafe.Add(mBase, uint32(v633)))
										v635 = *(*int64)(unsafe.Add(mBase, uint32(v594)+16))
										v637 = base.I64_div_s(v635, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v633))) = v634 + v637
										v641 = v581 + v322
										v642 = *(*int64)(unsafe.Add(mBase, uint32(v641)))
										v643 = *(*int64)(unsafe.Add(mBase, uint32(v583)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v641))) = v642 + v643
										v647 = v587 + v322
										v648 = *(*int64)(unsafe.Add(mBase, uint32(v647)))
										v649 = *(*int64)(unsafe.Add(mBase, uint32(v244)+216))
										*(*int64)(unsafe.Add(mBase, uint32(v647))) = v648 + v649
										v653 = v592 + v322
										v654 = *(*int64)(unsafe.Add(mBase, uint32(v653)))
										v655 = *(*int64)(unsafe.Add(mBase, uint32(v594)+24))
										v657 = base.I64_div_s(v655, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v653))) = v654 + v657
										v661 = v581 + v342
										v662 = *(*int64)(unsafe.Add(mBase, uint32(v661)))
										v663 = *(*int64)(unsafe.Add(mBase, uint32(v583)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v661))) = v662 + v663
										v667 = v587 + v342
										v668 = *(*int64)(unsafe.Add(mBase, uint32(v667)))
										v669 = *(*int64)(unsafe.Add(mBase, uint32(v244)+224))
										*(*int64)(unsafe.Add(mBase, uint32(v667))) = v668 + v669
										v673 = v592 + v342
										v674 = *(*int64)(unsafe.Add(mBase, uint32(v673)))
										v675 = *(*int64)(unsafe.Add(mBase, uint32(v594)+32))
										v677 = base.I64_div_s(v675, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v673))) = v674 + v677
										v681 = v581 + v362
										v682 = *(*int64)(unsafe.Add(mBase, uint32(v681)))
										v683 = *(*int64)(unsafe.Add(mBase, uint32(v583)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v681))) = v682 + v683
										v687 = v587 + v362
										v688 = *(*int64)(unsafe.Add(mBase, uint32(v687)))
										v689 = *(*int64)(unsafe.Add(mBase, uint32(v244)+232))
										*(*int64)(unsafe.Add(mBase, uint32(v687))) = v688 + v689
										v693 = v592 + v362
										v694 = *(*int64)(unsafe.Add(mBase, uint32(v693)))
										v695 = *(*int64)(unsafe.Add(mBase, uint32(v594)+40))
										v697 = base.I64_div_s(v695, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v693))) = v694 + v697
										v701 = v581 + v382
										v702 = *(*int64)(unsafe.Add(mBase, uint32(v701)))
										v703 = *(*int64)(unsafe.Add(mBase, uint32(v583)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v701))) = v702 + v703
										v707 = v587 + v382
										v708 = *(*int64)(unsafe.Add(mBase, uint32(v707)))
										v709 = *(*int64)(unsafe.Add(mBase, uint32(v244)+240))
										*(*int64)(unsafe.Add(mBase, uint32(v707))) = v708 + v709
										v713 = v592 + v382
										v714 = *(*int64)(unsafe.Add(mBase, uint32(v713)))
										v715 = *(*int64)(unsafe.Add(mBase, uint32(v594)+48))
										v717 = base.I64_div_s(v715, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v713))) = v714 + v717
										v721 = v581 + v402
										v722 = *(*int64)(unsafe.Add(mBase, uint32(v721)))
										v723 = *(*int64)(unsafe.Add(mBase, uint32(v583)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v721))) = v722 + v723
										v727 = v587 + v402
										v728 = *(*int64)(unsafe.Add(mBase, uint32(v727)))
										v729 = *(*int64)(unsafe.Add(mBase, uint32(v244)+248))
										*(*int64)(unsafe.Add(mBase, uint32(v727))) = v728 + v729
										v733 = v592 + v402
										v734 = *(*int64)(unsafe.Add(mBase, uint32(v733)))
										v735 = *(*int64)(unsafe.Add(mBase, uint32(v594)+56))
										v737 = base.I64_div_s(v735, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v733))) = v734 + v737
										v740 = v188 + (v91 + int32(1248))
										v741 = *(*int64)(unsafe.Add(mBase, uint32(v740)))
										v742 = v188 + (v124 + int32(1216))
										v743 = *(*int64)(unsafe.Add(mBase, uint32(v742)))
										*(*int64)(unsafe.Add(mBase, uint32(v740))) = v741 + v743
										v746 = v188 + (v91 + int32(288))
										v747 = *(*int64)(unsafe.Add(mBase, uint32(v746)))
										v748 = *(*int64)(unsafe.Add(mBase, uint32(v244)+256))
										*(*int64)(unsafe.Add(mBase, uint32(v746))) = v747 + v748
										v751 = v188 + (v91 + int32(2208))
										v752 = *(*int64)(unsafe.Add(mBase, uint32(v751)))
										v753 = v188 + (v124 + int32(2176))
										v754 = *(*int64)(unsafe.Add(mBase, uint32(v753)))
										v756 = base.I64_div_s(v754, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v751))) = v752 + v756
										v760 = v740 + v282
										v761 = *(*int64)(unsafe.Add(mBase, uint32(v760)))
										v762 = *(*int64)(unsafe.Add(mBase, uint32(v742)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v760))) = v761 + v762
										v766 = v746 + v282
										v767 = *(*int64)(unsafe.Add(mBase, uint32(v766)))
										v768 = *(*int64)(unsafe.Add(mBase, uint32(v244)+264))
										*(*int64)(unsafe.Add(mBase, uint32(v766))) = v767 + v768
										v772 = v751 + v282
										v773 = *(*int64)(unsafe.Add(mBase, uint32(v772)))
										v774 = *(*int64)(unsafe.Add(mBase, uint32(v753)+8))
										v776 = base.I64_div_s(v774, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v772))) = v773 + v776
										v780 = v740 + v302
										v781 = *(*int64)(unsafe.Add(mBase, uint32(v780)))
										v782 = *(*int64)(unsafe.Add(mBase, uint32(v742)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v780))) = v781 + v782
										v786 = v746 + v302
										v787 = *(*int64)(unsafe.Add(mBase, uint32(v786)))
										v788 = *(*int64)(unsafe.Add(mBase, uint32(v244)+272))
										*(*int64)(unsafe.Add(mBase, uint32(v786))) = v787 + v788
										v792 = v751 + v302
										v793 = *(*int64)(unsafe.Add(mBase, uint32(v792)))
										v794 = *(*int64)(unsafe.Add(mBase, uint32(v753)+16))
										v796 = base.I64_div_s(v794, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v792))) = v793 + v796
										v800 = v740 + v322
										v801 = *(*int64)(unsafe.Add(mBase, uint32(v800)))
										v802 = *(*int64)(unsafe.Add(mBase, uint32(v742)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v800))) = v801 + v802
										v806 = v746 + v322
										v807 = *(*int64)(unsafe.Add(mBase, uint32(v806)))
										v808 = *(*int64)(unsafe.Add(mBase, uint32(v244)+280))
										*(*int64)(unsafe.Add(mBase, uint32(v806))) = v807 + v808
										v812 = v751 + v322
										v813 = *(*int64)(unsafe.Add(mBase, uint32(v812)))
										v814 = *(*int64)(unsafe.Add(mBase, uint32(v753)+24))
										v816 = base.I64_div_s(v814, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v812))) = v813 + v816
										v820 = v740 + v342
										v821 = *(*int64)(unsafe.Add(mBase, uint32(v820)))
										v822 = *(*int64)(unsafe.Add(mBase, uint32(v742)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v820))) = v821 + v822
										v826 = v746 + v342
										v827 = *(*int64)(unsafe.Add(mBase, uint32(v826)))
										v828 = *(*int64)(unsafe.Add(mBase, uint32(v244)+288))
										*(*int64)(unsafe.Add(mBase, uint32(v826))) = v827 + v828
										v832 = v751 + v342
										v833 = *(*int64)(unsafe.Add(mBase, uint32(v832)))
										v834 = *(*int64)(unsafe.Add(mBase, uint32(v753)+32))
										v836 = base.I64_div_s(v834, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v832))) = v833 + v836
										v840 = v740 + v362
										v841 = *(*int64)(unsafe.Add(mBase, uint32(v840)))
										v842 = *(*int64)(unsafe.Add(mBase, uint32(v742)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v840))) = v841 + v842
										v846 = v746 + v362
										v847 = *(*int64)(unsafe.Add(mBase, uint32(v846)))
										v848 = *(*int64)(unsafe.Add(mBase, uint32(v244)+296))
										*(*int64)(unsafe.Add(mBase, uint32(v846))) = v847 + v848
										v852 = v751 + v362
										v853 = *(*int64)(unsafe.Add(mBase, uint32(v852)))
										v854 = *(*int64)(unsafe.Add(mBase, uint32(v753)+40))
										v856 = base.I64_div_s(v854, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v852))) = v853 + v856
										v860 = v740 + v382
										v861 = *(*int64)(unsafe.Add(mBase, uint32(v860)))
										v862 = *(*int64)(unsafe.Add(mBase, uint32(v742)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v860))) = v861 + v862
										v866 = v746 + v382
										v867 = *(*int64)(unsafe.Add(mBase, uint32(v866)))
										v868 = *(*int64)(unsafe.Add(mBase, uint32(v244)+304))
										*(*int64)(unsafe.Add(mBase, uint32(v866))) = v867 + v868
										v872 = v751 + v382
										v873 = *(*int64)(unsafe.Add(mBase, uint32(v872)))
										v874 = *(*int64)(unsafe.Add(mBase, uint32(v753)+48))
										v876 = base.I64_div_s(v874, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v872))) = v873 + v876
										v880 = v740 + v402
										v881 = *(*int64)(unsafe.Add(mBase, uint32(v880)))
										v882 = *(*int64)(unsafe.Add(mBase, uint32(v742)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v880))) = v881 + v882
										v886 = v746 + v402
										v887 = *(*int64)(unsafe.Add(mBase, uint32(v886)))
										v888 = *(*int64)(unsafe.Add(mBase, uint32(v244)+312))
										*(*int64)(unsafe.Add(mBase, uint32(v886))) = v887 + v888
										v892 = v751 + v402
										v893 = *(*int64)(unsafe.Add(mBase, uint32(v892)))
										v894 = *(*int64)(unsafe.Add(mBase, uint32(v753)+56))
										v896 = base.I64_div_s(v894, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v892))) = v893 + v896
										v900 = v157 + int32(1)
										if v900 != int32(3) {
											v157 = v900
											continue
										} else {
											break
										}
										break
									}
									v907 = F__emscripten_memset_bulkmem(m, int32(4401960), base.I32_extend8_s(int32(0)), int32(2880))
									mBase = m.M
									v909 = int32(0)
									*(*uint8)(unsafe.Add(mBase, _consts[738])) = uint8(v909)
								}
							}
							if v59 == int32(0) {
							} else {
								v953 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v43)+24)) = v953
								*(*int64)(unsafe.Add(mBase, uint32(v43)+16)) = v953
								*(*int64)(unsafe.Add(mBase, uint32(v43)+8)) = v953
								*(*int64)(unsafe.Add(mBase, uint32(v43))) = v953
								v962 = *(*int64)(unsafe.Add(mBase, _consts[18]))
								v964 = *(*int64)(unsafe.Add(mBase, _consts[932]))
								if v962 == v964 {
								} else {
									v966 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
									v967 = int32(4404848)
									v968 = *(*int64)(unsafe.Add(mBase, uint32(v43)+16))
									v970 = *(*int64)(unsafe.Add(mBase, _consts[14]))
									v971 = *(*int64)(unsafe.Add(mBase, _consts[933]))
									*(*int64)(unsafe.Add(mBase, uint32(v43)+16)) = v968 + (v970 - v971)
									v975 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
									v977 = *(*int64)(unsafe.Add(mBase, _consts[18]))
									v978 = *(*int64)(unsafe.Add(mBase, _consts[932]))
									*(*int64)(unsafe.Add(mBase, uint32(v43))) = v975 + (v977 - v978)
									v982 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
									v984 = *(*int64)(unsafe.Add(mBase, _consts[16]))
									v985 = *(*int64)(unsafe.Add(mBase, _consts[934]))
									*(*int64)(unsafe.Add(mBase, uint32(v43)+8)) = v982 + (v984 - v985)
									v989 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
									v991 = *(*int64)(unsafe.Add(mBase, _consts[12]))
									v992 = *(*int64)(unsafe.Add(mBase, _consts[935]))
									*(*int64)(unsafe.Add(mBase, uint32(v43)+24)) = v989 + (v991 - v992)
									v996 = *(*int64)(unsafe.Add(mBase, uint32(v966)+2936))
									v997 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
									*(*int64)(unsafe.Add(mBase, uint32(v966)+2936)) = v996 + v997
									v1000 = *(*int64)(unsafe.Add(mBase, uint32(v966)+2912))
									v1001 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
									*(*int64)(unsafe.Add(mBase, uint32(v966)+2912)) = v1000 + v1001
									v1004 = *(*int64)(unsafe.Add(mBase, uint32(v966)+2920))
									v1005 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v966)+2920)) = v1004 + v1005
									v1008 = *(*int64)(unsafe.Add(mBase, uint32(v966)+2928))
									v1009 = *(*int64)(unsafe.Add(mBase, uint32(v43)+16))
									*(*int64)(unsafe.Add(mBase, uint32(v966)+2928)) = v1008 + v1009
									v1014 = *(*int64)(unsafe.Add(mBase, _consts[12]))
									*(*int64)(unsafe.Add(mBase, _consts[935])) = v1014
									v1018 = *(*int64)(unsafe.Add(mBase, _consts[14]))
									*(*int64)(unsafe.Add(mBase, _consts[933])) = v1018
									v1022 = *(*int64)(unsafe.Add(mBase, _consts[16]))
									*(*int64)(unsafe.Add(mBase, _consts[934])) = v1022
									v1026 = *(*int64)(unsafe.Add(mBase, _consts[18]))
									*(*int64)(unsafe.Add(mBase, _consts[932])) = v1026
								}
							}
							F_pgstat_unlock_entry(m, v76)
							mBase = m.M
							v1030 = m.ExcPending
							if v1030 != 0 {
								return int32(0)
							} else {
								v1035 = int32(0)
								m.G0 = v43 + int32(2880)
								return v1035
							}
						}
					}
				} else {
					v1035 = v3
					m.G0 = v43 + int32(2880)
					return v1035
				}
			} else {
				if v57&int32(1) == int32(0) {
					v1035 = v3
					m.G0 = v43 + int32(2880)
					return v1035
				} else {
					v75 = int64(*(*int32)(unsafe.Add(mBase, _consts[126])))
					v76 = F_pgstat_get_entry_ref_locked(m, int32(6), int32(0), v75, l0)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						if v76 == int32(0) {
							v1035 = int32(1)
							m.G0 = v43 + int32(2880)
							return v1035
						} else {
							if l1&int32(1) == int32(0) {
							} else {
								v88 = int32(*(*uint8)(unsafe.Add(mBase, _consts[738])))
								if v88 == int32(0) {
								} else {
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
									v124 = F__emscripten_memcpy_bulkmem(m, v43, int32(4401960), int32(2880))
									mBase = m.M
									v157 = v3
									for {
										v188 = v157 * int32(320)
										v195 = int32(0)
										for {
											v235 = v195 << (uint(int32(3)) % 32)
											v236 = v188 + (v91 + int32(992)) + v235
											v237 = *(*int64)(unsafe.Add(mBase, uint32(v236)))
											v239 = *(*int64)(unsafe.Add(mBase, uint32(v235+(v124+int32(960)+v188))))
											*(*int64)(unsafe.Add(mBase, uint32(v236))) = v237 + v239
											v242 = v188 + (v91 + int32(32)) + v235
											v243 = *(*int64)(unsafe.Add(mBase, uint32(v242)))
											v244 = v188 + v124
											v246 = *(*int64)(unsafe.Add(mBase, uint32(v244+v235)))
											*(*int64)(unsafe.Add(mBase, uint32(v242))) = v243 + v246
											v249 = v235 + (v188 + (v91 + int32(1952)))
											v250 = *(*int64)(unsafe.Add(mBase, uint32(v249)))
											v252 = *(*int64)(unsafe.Add(mBase, uint32(v235+(v188+(v124+int32(1920))))))
											v254 = base.I64_div_s(v252, int64(1000))
											*(*int64)(unsafe.Add(mBase, uint32(v249))) = v250 + v254
											v258 = v195 + int32(1)
											if v258 != int32(8) {
												v195 = v258
												continue
											} else {
												break
											}
											break
										}
										v261 = v188 + (v91 + int32(1056))
										v262 = *(*int64)(unsafe.Add(mBase, uint32(v261)))
										v263 = v188 + (v124 + int32(1024))
										v264 = *(*int64)(unsafe.Add(mBase, uint32(v263)))
										*(*int64)(unsafe.Add(mBase, uint32(v261))) = v262 + v264
										v267 = v188 + (v91 + int32(96))
										v268 = *(*int64)(unsafe.Add(mBase, uint32(v267)))
										v271 = *(*int64)(unsafe.Add(mBase, uint32(v244-int32(-64))))
										*(*int64)(unsafe.Add(mBase, uint32(v267))) = v268 + v271
										v274 = v188 + (v91 + int32(2016))
										v275 = *(*int64)(unsafe.Add(mBase, uint32(v274)))
										v276 = v188 + (v124 + int32(1984))
										v277 = *(*int64)(unsafe.Add(mBase, uint32(v276)))
										v278 = int64(1000)
										v279 = base.I64_div_s(v277, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v274))) = v275 + v279
										v282 = int32(8)
										v283 = v261 + v282
										v284 = *(*int64)(unsafe.Add(mBase, uint32(v283)))
										v285 = *(*int64)(unsafe.Add(mBase, uint32(v263)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v283))) = v284 + v285
										v289 = v267 + v282
										v290 = *(*int64)(unsafe.Add(mBase, uint32(v289)))
										v291 = *(*int64)(unsafe.Add(mBase, uint32(v244)+72))
										*(*int64)(unsafe.Add(mBase, uint32(v289))) = v290 + v291
										v295 = v274 + v282
										v296 = *(*int64)(unsafe.Add(mBase, uint32(v295)))
										v297 = *(*int64)(unsafe.Add(mBase, uint32(v276)+8))
										v299 = base.I64_div_s(v297, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v295))) = v296 + v299
										v302 = int32(16)
										v303 = v261 + v302
										v304 = *(*int64)(unsafe.Add(mBase, uint32(v303)))
										v305 = *(*int64)(unsafe.Add(mBase, uint32(v263)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v303))) = v304 + v305
										v309 = v267 + v302
										v310 = *(*int64)(unsafe.Add(mBase, uint32(v309)))
										v311 = *(*int64)(unsafe.Add(mBase, uint32(v244)+80))
										*(*int64)(unsafe.Add(mBase, uint32(v309))) = v310 + v311
										v315 = v274 + v302
										v316 = *(*int64)(unsafe.Add(mBase, uint32(v315)))
										v317 = *(*int64)(unsafe.Add(mBase, uint32(v276)+16))
										v319 = base.I64_div_s(v317, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v315))) = v316 + v319
										v322 = int32(24)
										v323 = v261 + v322
										v324 = *(*int64)(unsafe.Add(mBase, uint32(v323)))
										v325 = *(*int64)(unsafe.Add(mBase, uint32(v263)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v323))) = v324 + v325
										v329 = v267 + v322
										v330 = *(*int64)(unsafe.Add(mBase, uint32(v329)))
										v331 = *(*int64)(unsafe.Add(mBase, uint32(v244)+88))
										*(*int64)(unsafe.Add(mBase, uint32(v329))) = v330 + v331
										v335 = v274 + v322
										v336 = *(*int64)(unsafe.Add(mBase, uint32(v335)))
										v337 = *(*int64)(unsafe.Add(mBase, uint32(v276)+24))
										v339 = base.I64_div_s(v337, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v335))) = v336 + v339
										v342 = int32(32)
										v343 = v261 + v342
										v344 = *(*int64)(unsafe.Add(mBase, uint32(v343)))
										v345 = *(*int64)(unsafe.Add(mBase, uint32(v263)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v343))) = v344 + v345
										v349 = v267 + v342
										v350 = *(*int64)(unsafe.Add(mBase, uint32(v349)))
										v351 = *(*int64)(unsafe.Add(mBase, uint32(v244)+96))
										*(*int64)(unsafe.Add(mBase, uint32(v349))) = v350 + v351
										v355 = v274 + v342
										v356 = *(*int64)(unsafe.Add(mBase, uint32(v355)))
										v357 = *(*int64)(unsafe.Add(mBase, uint32(v276)+32))
										v359 = base.I64_div_s(v357, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v355))) = v356 + v359
										v362 = int32(40)
										v363 = v261 + v362
										v364 = *(*int64)(unsafe.Add(mBase, uint32(v363)))
										v365 = *(*int64)(unsafe.Add(mBase, uint32(v263)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v363))) = v364 + v365
										v369 = v267 + v362
										v370 = *(*int64)(unsafe.Add(mBase, uint32(v369)))
										v371 = *(*int64)(unsafe.Add(mBase, uint32(v244)+104))
										*(*int64)(unsafe.Add(mBase, uint32(v369))) = v370 + v371
										v375 = v274 + v362
										v376 = *(*int64)(unsafe.Add(mBase, uint32(v375)))
										v377 = *(*int64)(unsafe.Add(mBase, uint32(v276)+40))
										v379 = base.I64_div_s(v377, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v375))) = v376 + v379
										v382 = int32(48)
										v383 = v261 + v382
										v384 = *(*int64)(unsafe.Add(mBase, uint32(v383)))
										v385 = *(*int64)(unsafe.Add(mBase, uint32(v263)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v383))) = v384 + v385
										v389 = v267 + v382
										v390 = *(*int64)(unsafe.Add(mBase, uint32(v389)))
										v391 = *(*int64)(unsafe.Add(mBase, uint32(v244)+112))
										*(*int64)(unsafe.Add(mBase, uint32(v389))) = v390 + v391
										v395 = v274 + v382
										v396 = *(*int64)(unsafe.Add(mBase, uint32(v395)))
										v397 = *(*int64)(unsafe.Add(mBase, uint32(v276)+48))
										v399 = base.I64_div_s(v397, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v395))) = v396 + v399
										v402 = int32(56)
										v403 = v261 + v402
										v404 = *(*int64)(unsafe.Add(mBase, uint32(v403)))
										v405 = *(*int64)(unsafe.Add(mBase, uint32(v263)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v403))) = v404 + v405
										v409 = v267 + v402
										v410 = *(*int64)(unsafe.Add(mBase, uint32(v409)))
										v411 = *(*int64)(unsafe.Add(mBase, uint32(v244)+120))
										*(*int64)(unsafe.Add(mBase, uint32(v409))) = v410 + v411
										v415 = v274 + v402
										v416 = *(*int64)(unsafe.Add(mBase, uint32(v415)))
										v417 = *(*int64)(unsafe.Add(mBase, uint32(v276)+56))
										v419 = base.I64_div_s(v417, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v415))) = v416 + v419
										v422 = v188 + (v91 + int32(1120))
										v423 = *(*int64)(unsafe.Add(mBase, uint32(v422)))
										v424 = v188 + (v124 + int32(1088))
										v425 = *(*int64)(unsafe.Add(mBase, uint32(v424)))
										*(*int64)(unsafe.Add(mBase, uint32(v422))) = v423 + v425
										v428 = v188 + (v91 + int32(160))
										v429 = *(*int64)(unsafe.Add(mBase, uint32(v428)))
										v430 = *(*int64)(unsafe.Add(mBase, uint32(v244)+128))
										*(*int64)(unsafe.Add(mBase, uint32(v428))) = v429 + v430
										v433 = v188 + (v91 + int32(2080))
										v434 = *(*int64)(unsafe.Add(mBase, uint32(v433)))
										v435 = v188 + (v124 + int32(2048))
										v436 = *(*int64)(unsafe.Add(mBase, uint32(v435)))
										v438 = base.I64_div_s(v436, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v433))) = v434 + v438
										v442 = v422 + v282
										v443 = *(*int64)(unsafe.Add(mBase, uint32(v442)))
										v444 = *(*int64)(unsafe.Add(mBase, uint32(v424)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v442))) = v443 + v444
										v448 = v428 + v282
										v449 = *(*int64)(unsafe.Add(mBase, uint32(v448)))
										v450 = *(*int64)(unsafe.Add(mBase, uint32(v244)+136))
										*(*int64)(unsafe.Add(mBase, uint32(v448))) = v449 + v450
										v454 = v433 + v282
										v455 = *(*int64)(unsafe.Add(mBase, uint32(v454)))
										v456 = *(*int64)(unsafe.Add(mBase, uint32(v435)+8))
										v458 = base.I64_div_s(v456, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v454))) = v455 + v458
										v462 = v422 + v302
										v463 = *(*int64)(unsafe.Add(mBase, uint32(v462)))
										v464 = *(*int64)(unsafe.Add(mBase, uint32(v424)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v462))) = v463 + v464
										v468 = v428 + v302
										v469 = *(*int64)(unsafe.Add(mBase, uint32(v468)))
										v470 = *(*int64)(unsafe.Add(mBase, uint32(v244)+144))
										*(*int64)(unsafe.Add(mBase, uint32(v468))) = v469 + v470
										v474 = v433 + v302
										v475 = *(*int64)(unsafe.Add(mBase, uint32(v474)))
										v476 = *(*int64)(unsafe.Add(mBase, uint32(v435)+16))
										v478 = base.I64_div_s(v476, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v474))) = v475 + v478
										v482 = v422 + v322
										v483 = *(*int64)(unsafe.Add(mBase, uint32(v482)))
										v484 = *(*int64)(unsafe.Add(mBase, uint32(v424)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v482))) = v483 + v484
										v488 = v428 + v322
										v489 = *(*int64)(unsafe.Add(mBase, uint32(v488)))
										v490 = *(*int64)(unsafe.Add(mBase, uint32(v244)+152))
										*(*int64)(unsafe.Add(mBase, uint32(v488))) = v489 + v490
										v494 = v433 + v322
										v495 = *(*int64)(unsafe.Add(mBase, uint32(v494)))
										v496 = *(*int64)(unsafe.Add(mBase, uint32(v435)+24))
										v498 = base.I64_div_s(v496, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v494))) = v495 + v498
										v502 = v422 + v342
										v503 = *(*int64)(unsafe.Add(mBase, uint32(v502)))
										v504 = *(*int64)(unsafe.Add(mBase, uint32(v424)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v502))) = v503 + v504
										v508 = v428 + v342
										v509 = *(*int64)(unsafe.Add(mBase, uint32(v508)))
										v510 = *(*int64)(unsafe.Add(mBase, uint32(v244)+160))
										*(*int64)(unsafe.Add(mBase, uint32(v508))) = v509 + v510
										v514 = v433 + v342
										v515 = *(*int64)(unsafe.Add(mBase, uint32(v514)))
										v516 = *(*int64)(unsafe.Add(mBase, uint32(v435)+32))
										v518 = base.I64_div_s(v516, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v514))) = v515 + v518
										v522 = v422 + v362
										v523 = *(*int64)(unsafe.Add(mBase, uint32(v522)))
										v524 = *(*int64)(unsafe.Add(mBase, uint32(v424)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v522))) = v523 + v524
										v528 = v428 + v362
										v529 = *(*int64)(unsafe.Add(mBase, uint32(v528)))
										v530 = *(*int64)(unsafe.Add(mBase, uint32(v244)+168))
										*(*int64)(unsafe.Add(mBase, uint32(v528))) = v529 + v530
										v534 = v433 + v362
										v535 = *(*int64)(unsafe.Add(mBase, uint32(v534)))
										v536 = *(*int64)(unsafe.Add(mBase, uint32(v435)+40))
										v538 = base.I64_div_s(v536, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v534))) = v535 + v538
										v542 = v422 + v382
										v543 = *(*int64)(unsafe.Add(mBase, uint32(v542)))
										v544 = *(*int64)(unsafe.Add(mBase, uint32(v424)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v542))) = v543 + v544
										v548 = v428 + v382
										v549 = *(*int64)(unsafe.Add(mBase, uint32(v548)))
										v550 = *(*int64)(unsafe.Add(mBase, uint32(v244)+176))
										*(*int64)(unsafe.Add(mBase, uint32(v548))) = v549 + v550
										v554 = v433 + v382
										v555 = *(*int64)(unsafe.Add(mBase, uint32(v554)))
										v556 = *(*int64)(unsafe.Add(mBase, uint32(v435)+48))
										v558 = base.I64_div_s(v556, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v554))) = v555 + v558
										v562 = v422 + v402
										v563 = *(*int64)(unsafe.Add(mBase, uint32(v562)))
										v564 = *(*int64)(unsafe.Add(mBase, uint32(v424)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v562))) = v563 + v564
										v568 = v428 + v402
										v569 = *(*int64)(unsafe.Add(mBase, uint32(v568)))
										v570 = *(*int64)(unsafe.Add(mBase, uint32(v244)+184))
										*(*int64)(unsafe.Add(mBase, uint32(v568))) = v569 + v570
										v574 = v433 + v402
										v575 = *(*int64)(unsafe.Add(mBase, uint32(v574)))
										v576 = *(*int64)(unsafe.Add(mBase, uint32(v435)+56))
										v578 = base.I64_div_s(v576, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v574))) = v575 + v578
										v581 = v188 + (v91 + int32(1184))
										v582 = *(*int64)(unsafe.Add(mBase, uint32(v581)))
										v583 = v188 + (v124 + int32(1152))
										v584 = *(*int64)(unsafe.Add(mBase, uint32(v583)))
										*(*int64)(unsafe.Add(mBase, uint32(v581))) = v582 + v584
										v587 = v188 + (v91 + int32(224))
										v588 = *(*int64)(unsafe.Add(mBase, uint32(v587)))
										v589 = *(*int64)(unsafe.Add(mBase, uint32(v244)+192))
										*(*int64)(unsafe.Add(mBase, uint32(v587))) = v588 + v589
										v592 = v188 + (v91 + int32(2144))
										v593 = *(*int64)(unsafe.Add(mBase, uint32(v592)))
										v594 = v188 + (v124 + int32(2112))
										v595 = *(*int64)(unsafe.Add(mBase, uint32(v594)))
										v597 = base.I64_div_s(v595, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v592))) = v593 + v597
										v601 = v581 + v282
										v602 = *(*int64)(unsafe.Add(mBase, uint32(v601)))
										v603 = *(*int64)(unsafe.Add(mBase, uint32(v583)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v601))) = v602 + v603
										v607 = v587 + v282
										v608 = *(*int64)(unsafe.Add(mBase, uint32(v607)))
										v609 = *(*int64)(unsafe.Add(mBase, uint32(v244)+200))
										*(*int64)(unsafe.Add(mBase, uint32(v607))) = v608 + v609
										v613 = v592 + v282
										v614 = *(*int64)(unsafe.Add(mBase, uint32(v613)))
										v615 = *(*int64)(unsafe.Add(mBase, uint32(v594)+8))
										v617 = base.I64_div_s(v615, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v613))) = v614 + v617
										v621 = v581 + v302
										v622 = *(*int64)(unsafe.Add(mBase, uint32(v621)))
										v623 = *(*int64)(unsafe.Add(mBase, uint32(v583)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v621))) = v622 + v623
										v627 = v587 + v302
										v628 = *(*int64)(unsafe.Add(mBase, uint32(v627)))
										v629 = *(*int64)(unsafe.Add(mBase, uint32(v244)+208))
										*(*int64)(unsafe.Add(mBase, uint32(v627))) = v628 + v629
										v633 = v592 + v302
										v634 = *(*int64)(unsafe.Add(mBase, uint32(v633)))
										v635 = *(*int64)(unsafe.Add(mBase, uint32(v594)+16))
										v637 = base.I64_div_s(v635, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v633))) = v634 + v637
										v641 = v581 + v322
										v642 = *(*int64)(unsafe.Add(mBase, uint32(v641)))
										v643 = *(*int64)(unsafe.Add(mBase, uint32(v583)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v641))) = v642 + v643
										v647 = v587 + v322
										v648 = *(*int64)(unsafe.Add(mBase, uint32(v647)))
										v649 = *(*int64)(unsafe.Add(mBase, uint32(v244)+216))
										*(*int64)(unsafe.Add(mBase, uint32(v647))) = v648 + v649
										v653 = v592 + v322
										v654 = *(*int64)(unsafe.Add(mBase, uint32(v653)))
										v655 = *(*int64)(unsafe.Add(mBase, uint32(v594)+24))
										v657 = base.I64_div_s(v655, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v653))) = v654 + v657
										v661 = v581 + v342
										v662 = *(*int64)(unsafe.Add(mBase, uint32(v661)))
										v663 = *(*int64)(unsafe.Add(mBase, uint32(v583)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v661))) = v662 + v663
										v667 = v587 + v342
										v668 = *(*int64)(unsafe.Add(mBase, uint32(v667)))
										v669 = *(*int64)(unsafe.Add(mBase, uint32(v244)+224))
										*(*int64)(unsafe.Add(mBase, uint32(v667))) = v668 + v669
										v673 = v592 + v342
										v674 = *(*int64)(unsafe.Add(mBase, uint32(v673)))
										v675 = *(*int64)(unsafe.Add(mBase, uint32(v594)+32))
										v677 = base.I64_div_s(v675, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v673))) = v674 + v677
										v681 = v581 + v362
										v682 = *(*int64)(unsafe.Add(mBase, uint32(v681)))
										v683 = *(*int64)(unsafe.Add(mBase, uint32(v583)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v681))) = v682 + v683
										v687 = v587 + v362
										v688 = *(*int64)(unsafe.Add(mBase, uint32(v687)))
										v689 = *(*int64)(unsafe.Add(mBase, uint32(v244)+232))
										*(*int64)(unsafe.Add(mBase, uint32(v687))) = v688 + v689
										v693 = v592 + v362
										v694 = *(*int64)(unsafe.Add(mBase, uint32(v693)))
										v695 = *(*int64)(unsafe.Add(mBase, uint32(v594)+40))
										v697 = base.I64_div_s(v695, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v693))) = v694 + v697
										v701 = v581 + v382
										v702 = *(*int64)(unsafe.Add(mBase, uint32(v701)))
										v703 = *(*int64)(unsafe.Add(mBase, uint32(v583)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v701))) = v702 + v703
										v707 = v587 + v382
										v708 = *(*int64)(unsafe.Add(mBase, uint32(v707)))
										v709 = *(*int64)(unsafe.Add(mBase, uint32(v244)+240))
										*(*int64)(unsafe.Add(mBase, uint32(v707))) = v708 + v709
										v713 = v592 + v382
										v714 = *(*int64)(unsafe.Add(mBase, uint32(v713)))
										v715 = *(*int64)(unsafe.Add(mBase, uint32(v594)+48))
										v717 = base.I64_div_s(v715, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v713))) = v714 + v717
										v721 = v581 + v402
										v722 = *(*int64)(unsafe.Add(mBase, uint32(v721)))
										v723 = *(*int64)(unsafe.Add(mBase, uint32(v583)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v721))) = v722 + v723
										v727 = v587 + v402
										v728 = *(*int64)(unsafe.Add(mBase, uint32(v727)))
										v729 = *(*int64)(unsafe.Add(mBase, uint32(v244)+248))
										*(*int64)(unsafe.Add(mBase, uint32(v727))) = v728 + v729
										v733 = v592 + v402
										v734 = *(*int64)(unsafe.Add(mBase, uint32(v733)))
										v735 = *(*int64)(unsafe.Add(mBase, uint32(v594)+56))
										v737 = base.I64_div_s(v735, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v733))) = v734 + v737
										v740 = v188 + (v91 + int32(1248))
										v741 = *(*int64)(unsafe.Add(mBase, uint32(v740)))
										v742 = v188 + (v124 + int32(1216))
										v743 = *(*int64)(unsafe.Add(mBase, uint32(v742)))
										*(*int64)(unsafe.Add(mBase, uint32(v740))) = v741 + v743
										v746 = v188 + (v91 + int32(288))
										v747 = *(*int64)(unsafe.Add(mBase, uint32(v746)))
										v748 = *(*int64)(unsafe.Add(mBase, uint32(v244)+256))
										*(*int64)(unsafe.Add(mBase, uint32(v746))) = v747 + v748
										v751 = v188 + (v91 + int32(2208))
										v752 = *(*int64)(unsafe.Add(mBase, uint32(v751)))
										v753 = v188 + (v124 + int32(2176))
										v754 = *(*int64)(unsafe.Add(mBase, uint32(v753)))
										v756 = base.I64_div_s(v754, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v751))) = v752 + v756
										v760 = v740 + v282
										v761 = *(*int64)(unsafe.Add(mBase, uint32(v760)))
										v762 = *(*int64)(unsafe.Add(mBase, uint32(v742)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v760))) = v761 + v762
										v766 = v746 + v282
										v767 = *(*int64)(unsafe.Add(mBase, uint32(v766)))
										v768 = *(*int64)(unsafe.Add(mBase, uint32(v244)+264))
										*(*int64)(unsafe.Add(mBase, uint32(v766))) = v767 + v768
										v772 = v751 + v282
										v773 = *(*int64)(unsafe.Add(mBase, uint32(v772)))
										v774 = *(*int64)(unsafe.Add(mBase, uint32(v753)+8))
										v776 = base.I64_div_s(v774, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v772))) = v773 + v776
										v780 = v740 + v302
										v781 = *(*int64)(unsafe.Add(mBase, uint32(v780)))
										v782 = *(*int64)(unsafe.Add(mBase, uint32(v742)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v780))) = v781 + v782
										v786 = v746 + v302
										v787 = *(*int64)(unsafe.Add(mBase, uint32(v786)))
										v788 = *(*int64)(unsafe.Add(mBase, uint32(v244)+272))
										*(*int64)(unsafe.Add(mBase, uint32(v786))) = v787 + v788
										v792 = v751 + v302
										v793 = *(*int64)(unsafe.Add(mBase, uint32(v792)))
										v794 = *(*int64)(unsafe.Add(mBase, uint32(v753)+16))
										v796 = base.I64_div_s(v794, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v792))) = v793 + v796
										v800 = v740 + v322
										v801 = *(*int64)(unsafe.Add(mBase, uint32(v800)))
										v802 = *(*int64)(unsafe.Add(mBase, uint32(v742)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v800))) = v801 + v802
										v806 = v746 + v322
										v807 = *(*int64)(unsafe.Add(mBase, uint32(v806)))
										v808 = *(*int64)(unsafe.Add(mBase, uint32(v244)+280))
										*(*int64)(unsafe.Add(mBase, uint32(v806))) = v807 + v808
										v812 = v751 + v322
										v813 = *(*int64)(unsafe.Add(mBase, uint32(v812)))
										v814 = *(*int64)(unsafe.Add(mBase, uint32(v753)+24))
										v816 = base.I64_div_s(v814, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v812))) = v813 + v816
										v820 = v740 + v342
										v821 = *(*int64)(unsafe.Add(mBase, uint32(v820)))
										v822 = *(*int64)(unsafe.Add(mBase, uint32(v742)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v820))) = v821 + v822
										v826 = v746 + v342
										v827 = *(*int64)(unsafe.Add(mBase, uint32(v826)))
										v828 = *(*int64)(unsafe.Add(mBase, uint32(v244)+288))
										*(*int64)(unsafe.Add(mBase, uint32(v826))) = v827 + v828
										v832 = v751 + v342
										v833 = *(*int64)(unsafe.Add(mBase, uint32(v832)))
										v834 = *(*int64)(unsafe.Add(mBase, uint32(v753)+32))
										v836 = base.I64_div_s(v834, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v832))) = v833 + v836
										v840 = v740 + v362
										v841 = *(*int64)(unsafe.Add(mBase, uint32(v840)))
										v842 = *(*int64)(unsafe.Add(mBase, uint32(v742)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v840))) = v841 + v842
										v846 = v746 + v362
										v847 = *(*int64)(unsafe.Add(mBase, uint32(v846)))
										v848 = *(*int64)(unsafe.Add(mBase, uint32(v244)+296))
										*(*int64)(unsafe.Add(mBase, uint32(v846))) = v847 + v848
										v852 = v751 + v362
										v853 = *(*int64)(unsafe.Add(mBase, uint32(v852)))
										v854 = *(*int64)(unsafe.Add(mBase, uint32(v753)+40))
										v856 = base.I64_div_s(v854, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v852))) = v853 + v856
										v860 = v740 + v382
										v861 = *(*int64)(unsafe.Add(mBase, uint32(v860)))
										v862 = *(*int64)(unsafe.Add(mBase, uint32(v742)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v860))) = v861 + v862
										v866 = v746 + v382
										v867 = *(*int64)(unsafe.Add(mBase, uint32(v866)))
										v868 = *(*int64)(unsafe.Add(mBase, uint32(v244)+304))
										*(*int64)(unsafe.Add(mBase, uint32(v866))) = v867 + v868
										v872 = v751 + v382
										v873 = *(*int64)(unsafe.Add(mBase, uint32(v872)))
										v874 = *(*int64)(unsafe.Add(mBase, uint32(v753)+48))
										v876 = base.I64_div_s(v874, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v872))) = v873 + v876
										v880 = v740 + v402
										v881 = *(*int64)(unsafe.Add(mBase, uint32(v880)))
										v882 = *(*int64)(unsafe.Add(mBase, uint32(v742)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v880))) = v881 + v882
										v886 = v746 + v402
										v887 = *(*int64)(unsafe.Add(mBase, uint32(v886)))
										v888 = *(*int64)(unsafe.Add(mBase, uint32(v244)+312))
										*(*int64)(unsafe.Add(mBase, uint32(v886))) = v887 + v888
										v892 = v751 + v402
										v893 = *(*int64)(unsafe.Add(mBase, uint32(v892)))
										v894 = *(*int64)(unsafe.Add(mBase, uint32(v753)+56))
										v896 = base.I64_div_s(v894, v278)
										*(*int64)(unsafe.Add(mBase, uint32(v892))) = v893 + v896
										v900 = v157 + int32(1)
										if v900 != int32(3) {
											v157 = v900
											continue
										} else {
											break
										}
										break
									}
									v907 = F__emscripten_memset_bulkmem(m, int32(4401960), base.I32_extend8_s(int32(0)), int32(2880))
									mBase = m.M
									v909 = int32(0)
									*(*uint8)(unsafe.Add(mBase, _consts[738])) = uint8(v909)
								}
							}
							if v59 == int32(0) {
							} else {
								v953 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v43)+24)) = v953
								*(*int64)(unsafe.Add(mBase, uint32(v43)+16)) = v953
								*(*int64)(unsafe.Add(mBase, uint32(v43)+8)) = v953
								*(*int64)(unsafe.Add(mBase, uint32(v43))) = v953
								v962 = *(*int64)(unsafe.Add(mBase, _consts[18]))
								v964 = *(*int64)(unsafe.Add(mBase, _consts[932]))
								if v962 == v964 {
								} else {
									v966 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
									v967 = int32(4404848)
									v968 = *(*int64)(unsafe.Add(mBase, uint32(v43)+16))
									v970 = *(*int64)(unsafe.Add(mBase, _consts[14]))
									v971 = *(*int64)(unsafe.Add(mBase, _consts[933]))
									*(*int64)(unsafe.Add(mBase, uint32(v43)+16)) = v968 + (v970 - v971)
									v975 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
									v977 = *(*int64)(unsafe.Add(mBase, _consts[18]))
									v978 = *(*int64)(unsafe.Add(mBase, _consts[932]))
									*(*int64)(unsafe.Add(mBase, uint32(v43))) = v975 + (v977 - v978)
									v982 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
									v984 = *(*int64)(unsafe.Add(mBase, _consts[16]))
									v985 = *(*int64)(unsafe.Add(mBase, _consts[934]))
									*(*int64)(unsafe.Add(mBase, uint32(v43)+8)) = v982 + (v984 - v985)
									v989 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
									v991 = *(*int64)(unsafe.Add(mBase, _consts[12]))
									v992 = *(*int64)(unsafe.Add(mBase, _consts[935]))
									*(*int64)(unsafe.Add(mBase, uint32(v43)+24)) = v989 + (v991 - v992)
									v996 = *(*int64)(unsafe.Add(mBase, uint32(v966)+2936))
									v997 = *(*int64)(unsafe.Add(mBase, uint32(v43)+24))
									*(*int64)(unsafe.Add(mBase, uint32(v966)+2936)) = v996 + v997
									v1000 = *(*int64)(unsafe.Add(mBase, uint32(v966)+2912))
									v1001 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
									*(*int64)(unsafe.Add(mBase, uint32(v966)+2912)) = v1000 + v1001
									v1004 = *(*int64)(unsafe.Add(mBase, uint32(v966)+2920))
									v1005 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v966)+2920)) = v1004 + v1005
									v1008 = *(*int64)(unsafe.Add(mBase, uint32(v966)+2928))
									v1009 = *(*int64)(unsafe.Add(mBase, uint32(v43)+16))
									*(*int64)(unsafe.Add(mBase, uint32(v966)+2928)) = v1008 + v1009
									v1014 = *(*int64)(unsafe.Add(mBase, _consts[12]))
									*(*int64)(unsafe.Add(mBase, _consts[935])) = v1014
									v1018 = *(*int64)(unsafe.Add(mBase, _consts[14]))
									*(*int64)(unsafe.Add(mBase, _consts[933])) = v1018
									v1022 = *(*int64)(unsafe.Add(mBase, _consts[16]))
									*(*int64)(unsafe.Add(mBase, _consts[934])) = v1022
									v1026 = *(*int64)(unsafe.Add(mBase, _consts[18]))
									*(*int64)(unsafe.Add(mBase, _consts[932])) = v1026
								}
							}
							F_pgstat_unlock_entry(m, v76)
							mBase = m.M
							v1030 = m.ExcPending
							if v1030 != 0 {
								return int32(0)
							} else {
								v1035 = int32(0)
								m.G0 = v43 + int32(2880)
								return v1035
							}
						}
					}
				}
			}
		}
	}
}
func F_pgstat_hash_hash_key(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int64
	_ = v17
	var v22 int64
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v59 int64
	_ = v59
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v108 int64
	_ = v108
	var v113 int64
	_ = v113
	var v121 int64
	_ = v121
	var v126 int64
	_ = v126
	v8 = base.I64_extend_i32_u(l1) * int64(-8645972361240307355)
	if base.Ui32(l1) < base.Ui32(int32(8)) {
		v72 = l0
		v74 = l1
		v76 = v8
	} else {
		v11 = int32(8)
		v12 = l1 - v11
		if v12&v11 == int32(0) {
			v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			v22 = (int64(base.Ui64(v17)>>(uint(int64(23))%64)) ^ v17) * int64(2388976653695081527)
			v31 = l0 + int32(8)
			v32 = v12
			v34 = (v8 ^ int64(base.Ui64(v22)>>(uint(int64(47))%64)) ^ v22) * int64(-8645972361240307355)
		} else {
			v31 = l0
			v32 = l1
			v34 = v8
		}
		if base.Ui32(v12) < base.Ui32(int32(8)) {
			v72 = v31
			v74 = v12
			v76 = v34
		} else {
			v37 = v31
			v39 = v32
			v41 = v34
			for {
				v42 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
				v43 = int64(23)
				v46 = int64(2388976653695081527)
				v47 = (int64(base.Ui64(v42)>>(uint(v43)%64)) ^ v42) * v46
				v48 = int64(47)
				v52 = int64(-8645972361240307355)
				v54 = *(*int64)(unsafe.Add(mBase, uint32(v37)+8))
				v59 = (int64(base.Ui64(v54)>>(uint(v43)%64)) ^ v54) * v46
				v65 = ((v41^int64(base.Ui64(v47)>>(uint(v48)%64))^v47)*v52 ^ int64(base.Ui64(v59)>>(uint(v48)%64)) ^ v59) * v52
				v66 = int32(16)
				v67 = v37 + v66
				v69 = v39 - v66
				if base.Ui32(int32(7)) < base.Ui32(v69) {
					v37 = v67
					v39 = v69
					v41 = v65
					continue
				} else {
					break
				}
				break
			}
			v72 = v67
			v74 = v69
			v76 = v65
		}
	}
	v77 = int64(0)
	switch v74 - int32(1) {
	case 0:
		v104 = v77
		v105 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72))))
		v108 = v104 | v105
		v113 = (v108 ^ int64(base.Ui64(v108)>>(uint(int64(23))%64))) * int64(2388976653695081527)
		v121 = (v76 ^ int64(base.Ui64(v113)>>(uint(int64(47))%64)) ^ v113) * int64(-8645972361240307355)
	case 1:
		v99 = v77
		v100 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72))))
		v108 = v104 | v105
		v113 = (v108 ^ int64(base.Ui64(v108)>>(uint(int64(23))%64))) * int64(2388976653695081527)
		v121 = (v76 ^ int64(base.Ui64(v113)>>(uint(int64(47))%64)) ^ v113) * int64(-8645972361240307355)
	case 2:
		v96 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+2)))
		v99 = v96 << (uint(int64(16)) % 64)
		v100 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+1)))
		v104 = v100<<(uint(int64(8))%64) | v99
		v105 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72))))
		v108 = v104 | v105
		v113 = (v108 ^ int64(base.Ui64(v108)>>(uint(int64(23))%64))) * int64(2388976653695081527)
		v121 = (v76 ^ int64(base.Ui64(v113)>>(uint(int64(47))%64)) ^ v113) * int64(-8645972361240307355)
	case 3:
		v93 = v77
		v94 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
		v108 = v93 | v94
		v113 = (v108 ^ int64(base.Ui64(v108)>>(uint(int64(23))%64))) * int64(2388976653695081527)
		v121 = (v76 ^ int64(base.Ui64(v113)>>(uint(int64(47))%64)) ^ v113) * int64(-8645972361240307355)
	case 4:
		v88 = v77
		v89 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+4)))
		v93 = v89<<(uint(int64(32))%64) | v88
		v94 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
		v108 = v93 | v94
		v113 = (v108 ^ int64(base.Ui64(v108)>>(uint(int64(23))%64))) * int64(2388976653695081527)
		v121 = (v76 ^ int64(base.Ui64(v113)>>(uint(int64(47))%64)) ^ v113) * int64(-8645972361240307355)
	case 5:
		v83 = v77
		v84 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+5)))
		v88 = v84<<(uint(int64(40))%64) | v83
		v89 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+4)))
		v93 = v89<<(uint(int64(32))%64) | v88
		v94 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
		v108 = v93 | v94
		v113 = (v108 ^ int64(base.Ui64(v108)>>(uint(int64(23))%64))) * int64(2388976653695081527)
		v121 = (v76 ^ int64(base.Ui64(v113)>>(uint(int64(47))%64)) ^ v113) * int64(-8645972361240307355)
	case 6:
		v80 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+6)))
		v83 = v80 << (uint(int64(48)) % 64)
		v84 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+5)))
		v88 = v84<<(uint(int64(40))%64) | v83
		v89 = int64(*(*int8)(unsafe.Add(mBase, uint32(v72)+4)))
		v93 = v89<<(uint(int64(32))%64) | v88
		v94 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v72))))
		v108 = v93 | v94
		v113 = (v108 ^ int64(base.Ui64(v108)>>(uint(int64(23))%64))) * int64(2388976653695081527)
		v121 = (v76 ^ int64(base.Ui64(v113)>>(uint(int64(47))%64)) ^ v113) * int64(-8645972361240307355)
	default:
		v121 = v76
	}
	v126 = (int64(base.Ui64(v121)>>(uint(int64(23))%64)) ^ v121) * int64(2388976653695081527)
	return base.I32_wrap_i64(int64(base.Ui64(v126)>>(uint(int64(47))%64)) ^ v126 - int64(base.Ui64(v126)>>(uint(int64(32))%64)))
}
func F_pgstat_init_function_usage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[941]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
	if v10 <= v12 {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
		m.G0 = v7 + int32(32)
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _consts[158]))
		v19 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)))
		v22 = F_pgstat_prep_pending_entry(m, int32(3), v18, v19, v7+int32(15))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
			if v24 == int32(1) {
				F_ReceiveSharedInvalidMessages(m)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					v32 = int32(0)
					v35 = F_SearchSysCacheExists(m, int32(47), v31, v32, v32, v32)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						if v35 == int32(0) {
							v63 = *(*int32)(unsafe.Add(mBase, _consts[158]))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v65 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v64)+4)))
							v66 = F_pgstat_drop_entry(m, int32(3), v63, v65)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									F_errcode(m, int32(52461700))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										F_errmsg(m, int32(238281), int32(0))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											F_errfinish(m, int32(465793), int32(118), int32(379754))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
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
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v39
							v41 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v41
							v44 = *(*int64)(unsafe.Add(mBase, _consts[942]))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v44
							F___clock_gettime(m, int32(1), v7+int32(16))
							mBase = m.M
							v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7)+24)))
							v51 = *(*int64)(unsafe.Add(mBase, uint32(v7)+16))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v50 + v51*int64(1000000000)
							m.G0 = v7 + int32(32)
							return
						}
					}
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v39
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v41
				v44 = *(*int64)(unsafe.Add(mBase, _consts[942]))
				*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v44
				F___clock_gettime(m, int32(1), v7+int32(16))
				mBase = m.M
				v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7)+24)))
				v51 = *(*int64)(unsafe.Add(mBase, uint32(v7)+16))
				*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v50 + v51*int64(1000000000)
				m.G0 = v7 + int32(32)
				return
			}
		}
	}
}
func F_pgstat_init_relation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+119)))
	switch v4 - int32(83) {
	case 0, 22, 26, 29, 31, 33:
		v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[23])))
		if v12 == int32(0) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
			if v15 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = int32(0)
			} else {
			}
			v18 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v18
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)) = uint8(v18)
			return
		} else {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)) = uint8(v22)
			return
		}
	default:
		v7 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v7
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)) = uint8(v7)
		return
	}
}
func F_pgstat_io_flush_cb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v124 int64
	_ = v124
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int64
	_ = v141
	var v144 int64
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v153 int64
	_ = v153
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v170 int64
	_ = v170
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v178 int64
	_ = v178
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v186 int64
	_ = v186
	var v188 int64
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int64
	_ = v193
	var v196 int64
	_ = v196
	var v200 int32
	_ = v200
	var v201 int64
	_ = v201
	var v204 int64
	_ = v204
	var v208 int32
	_ = v208
	var v209 int64
	_ = v209
	var v212 int64
	_ = v212
	var v214 int64
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int64
	_ = v219
	var v222 int64
	_ = v222
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v230 int64
	_ = v230
	var v234 int32
	_ = v234
	var v235 int64
	_ = v235
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	var v248 int64
	_ = v248
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v256 int64
	_ = v256
	var v260 int32
	_ = v260
	var v261 int64
	_ = v261
	var v264 int64
	_ = v264
	var v266 int64
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int64
	_ = v271
	var v274 int64
	_ = v274
	var v278 int32
	_ = v278
	var v279 int64
	_ = v279
	var v282 int64
	_ = v282
	var v286 int32
	_ = v286
	var v287 int64
	_ = v287
	var v290 int64
	_ = v290
	var v292 int64
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int64
	_ = v297
	var v300 int64
	_ = v300
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v308 int64
	_ = v308
	var v312 int32
	_ = v312
	var v313 int64
	_ = v313
	var v316 int64
	_ = v316
	var v318 int64
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int64
	_ = v323
	var v326 int64
	_ = v326
	var v330 int32
	_ = v330
	var v331 int64
	_ = v331
	var v334 int64
	_ = v334
	var v338 int32
	_ = v338
	var v339 int64
	_ = v339
	var v342 int64
	_ = v342
	var v344 int64
	_ = v344
	var v347 int32
	_ = v347
	var v348 int64
	_ = v348
	var v351 int64
	_ = v351
	var v355 int32
	_ = v355
	var v356 int64
	_ = v356
	var v359 int64
	_ = v359
	var v362 int32
	_ = v362
	var v363 int64
	_ = v363
	var v366 int64
	_ = v366
	var v368 int64
	_ = v368
	var v372 int32
	_ = v372
	var v373 int64
	_ = v373
	var v376 int64
	_ = v376
	var v380 int32
	_ = v380
	var v381 int64
	_ = v381
	var v384 int64
	_ = v384
	var v388 int32
	_ = v388
	var v389 int64
	_ = v389
	var v392 int64
	_ = v392
	var v394 int64
	_ = v394
	var v398 int32
	_ = v398
	var v399 int64
	_ = v399
	var v402 int64
	_ = v402
	var v406 int32
	_ = v406
	var v407 int64
	_ = v407
	var v410 int64
	_ = v410
	var v414 int32
	_ = v414
	var v415 int64
	_ = v415
	var v418 int64
	_ = v418
	var v420 int64
	_ = v420
	var v424 int32
	_ = v424
	var v425 int64
	_ = v425
	var v428 int64
	_ = v428
	var v432 int32
	_ = v432
	var v433 int64
	_ = v433
	var v436 int64
	_ = v436
	var v440 int32
	_ = v440
	var v441 int64
	_ = v441
	var v444 int64
	_ = v444
	var v446 int64
	_ = v446
	var v450 int32
	_ = v450
	var v451 int64
	_ = v451
	var v454 int64
	_ = v454
	var v458 int32
	_ = v458
	var v459 int64
	_ = v459
	var v462 int64
	_ = v462
	var v466 int32
	_ = v466
	var v467 int64
	_ = v467
	var v470 int64
	_ = v470
	var v472 int64
	_ = v472
	var v476 int32
	_ = v476
	var v477 int64
	_ = v477
	var v480 int64
	_ = v480
	var v484 int32
	_ = v484
	var v485 int64
	_ = v485
	var v488 int64
	_ = v488
	var v492 int32
	_ = v492
	var v493 int64
	_ = v493
	var v496 int64
	_ = v496
	var v498 int64
	_ = v498
	var v502 int32
	_ = v502
	var v503 int64
	_ = v503
	var v506 int64
	_ = v506
	var v510 int32
	_ = v510
	var v511 int64
	_ = v511
	var v514 int64
	_ = v514
	var v518 int32
	_ = v518
	var v519 int64
	_ = v519
	var v522 int64
	_ = v522
	var v524 int64
	_ = v524
	var v528 int32
	_ = v528
	var v529 int64
	_ = v529
	var v532 int64
	_ = v532
	var v536 int32
	_ = v536
	var v537 int64
	_ = v537
	var v540 int64
	_ = v540
	var v544 int32
	_ = v544
	var v545 int64
	_ = v545
	var v548 int64
	_ = v548
	var v550 int64
	_ = v550
	var v553 int32
	_ = v553
	var v554 int64
	_ = v554
	var v557 int64
	_ = v557
	var v561 int32
	_ = v561
	var v562 int64
	_ = v562
	var v565 int64
	_ = v565
	var v568 int32
	_ = v568
	var v569 int64
	_ = v569
	var v572 int64
	_ = v572
	var v574 int64
	_ = v574
	var v578 int32
	_ = v578
	var v579 int64
	_ = v579
	var v582 int64
	_ = v582
	var v586 int32
	_ = v586
	var v587 int64
	_ = v587
	var v590 int64
	_ = v590
	var v594 int32
	_ = v594
	var v595 int64
	_ = v595
	var v598 int64
	_ = v598
	var v600 int64
	_ = v600
	var v604 int32
	_ = v604
	var v605 int64
	_ = v605
	var v608 int64
	_ = v608
	var v612 int32
	_ = v612
	var v613 int64
	_ = v613
	var v616 int64
	_ = v616
	var v620 int32
	_ = v620
	var v621 int64
	_ = v621
	var v624 int64
	_ = v624
	var v626 int64
	_ = v626
	var v630 int32
	_ = v630
	var v631 int64
	_ = v631
	var v634 int64
	_ = v634
	var v638 int32
	_ = v638
	var v639 int64
	_ = v639
	var v642 int64
	_ = v642
	var v646 int32
	_ = v646
	var v647 int64
	_ = v647
	var v650 int64
	_ = v650
	var v652 int64
	_ = v652
	var v656 int32
	_ = v656
	var v657 int64
	_ = v657
	var v660 int64
	_ = v660
	var v664 int32
	_ = v664
	var v665 int64
	_ = v665
	var v668 int64
	_ = v668
	var v672 int32
	_ = v672
	var v673 int64
	_ = v673
	var v676 int64
	_ = v676
	var v678 int64
	_ = v678
	var v682 int32
	_ = v682
	var v683 int64
	_ = v683
	var v686 int64
	_ = v686
	var v690 int32
	_ = v690
	var v691 int64
	_ = v691
	var v694 int64
	_ = v694
	var v698 int32
	_ = v698
	var v699 int64
	_ = v699
	var v702 int64
	_ = v702
	var v704 int64
	_ = v704
	var v708 int32
	_ = v708
	var v709 int64
	_ = v709
	var v712 int64
	_ = v712
	var v716 int32
	_ = v716
	var v717 int64
	_ = v717
	var v720 int64
	_ = v720
	var v724 int32
	_ = v724
	var v725 int64
	_ = v725
	var v728 int64
	_ = v728
	var v730 int64
	_ = v730
	var v734 int32
	_ = v734
	var v735 int64
	_ = v735
	var v738 int64
	_ = v738
	var v742 int32
	_ = v742
	var v743 int64
	_ = v743
	var v746 int64
	_ = v746
	var v750 int32
	_ = v750
	var v751 int64
	_ = v751
	var v754 int64
	_ = v754
	var v756 int64
	_ = v756
	var v759 int32
	_ = v759
	var v760 int64
	_ = v760
	var v763 int64
	_ = v763
	var v767 int32
	_ = v767
	var v768 int64
	_ = v768
	var v771 int64
	_ = v771
	var v774 int32
	_ = v774
	var v775 int64
	_ = v775
	var v778 int64
	_ = v778
	var v780 int64
	_ = v780
	var v784 int32
	_ = v784
	var v785 int64
	_ = v785
	var v788 int64
	_ = v788
	var v792 int32
	_ = v792
	var v793 int64
	_ = v793
	var v796 int64
	_ = v796
	var v800 int32
	_ = v800
	var v801 int64
	_ = v801
	var v804 int64
	_ = v804
	var v806 int64
	_ = v806
	var v810 int32
	_ = v810
	var v811 int64
	_ = v811
	var v814 int64
	_ = v814
	var v818 int32
	_ = v818
	var v819 int64
	_ = v819
	var v822 int64
	_ = v822
	var v826 int32
	_ = v826
	var v827 int64
	_ = v827
	var v830 int64
	_ = v830
	var v832 int64
	_ = v832
	var v836 int32
	_ = v836
	var v837 int64
	_ = v837
	var v840 int64
	_ = v840
	var v844 int32
	_ = v844
	var v845 int64
	_ = v845
	var v848 int64
	_ = v848
	var v852 int32
	_ = v852
	var v853 int64
	_ = v853
	var v856 int64
	_ = v856
	var v858 int64
	_ = v858
	var v862 int32
	_ = v862
	var v863 int64
	_ = v863
	var v866 int64
	_ = v866
	var v870 int32
	_ = v870
	var v871 int64
	_ = v871
	var v874 int64
	_ = v874
	var v878 int32
	_ = v878
	var v879 int64
	_ = v879
	var v882 int64
	_ = v882
	var v884 int64
	_ = v884
	var v888 int32
	_ = v888
	var v889 int64
	_ = v889
	var v892 int64
	_ = v892
	var v896 int32
	_ = v896
	var v897 int64
	_ = v897
	var v900 int64
	_ = v900
	var v904 int32
	_ = v904
	var v905 int64
	_ = v905
	var v908 int64
	_ = v908
	var v910 int64
	_ = v910
	var v914 int32
	_ = v914
	var v915 int64
	_ = v915
	var v918 int64
	_ = v918
	var v922 int32
	_ = v922
	var v923 int64
	_ = v923
	var v926 int64
	_ = v926
	var v930 int32
	_ = v930
	var v931 int64
	_ = v931
	var v934 int64
	_ = v934
	var v936 int64
	_ = v936
	var v940 int32
	_ = v940
	var v941 int64
	_ = v941
	var v944 int64
	_ = v944
	var v948 int32
	_ = v948
	var v949 int64
	_ = v949
	var v952 int64
	_ = v952
	var v956 int32
	_ = v956
	var v957 int64
	_ = v957
	var v960 int64
	_ = v960
	var v962 int64
	_ = v962
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v984 int32
	_ = v984
	v2 = int32(0)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _consts[726])))
	if v20 == v2 {
		v984 = v2
		return v984
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, _consts[936]))
		v26 = *(*int32)(unsafe.Add(mBase, _consts[407]))
		v31 = v24 + v26<<(uint(int32(4))%32) + int32(608)
		if l0 == int32(0) {
			v35 = F_LWLockAcquire(m, v31, int32(0))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v48 = v24 + v26*int32(2880)
				v50 = v48 + int32(904)
				v78 = v2
				for {
					v90 = v78 * int32(320)
					v94 = int32(0)
					for {
						v112 = v90 + v94<<(uint(int32(3))%32)
						v113 = v112 + (v48 + int32(1864))
						v114 = *(*int64)(unsafe.Add(mBase, uint32(v113)))
						v117 = *(*int64)(unsafe.Add(mBase, uint32(v112)+uint32(_consts[943])))
						*(*int64)(unsafe.Add(mBase, uint32(v113))) = v114 + v117
						v120 = v112 + v50
						v121 = *(*int64)(unsafe.Add(mBase, uint32(v120)))
						v124 = *(*int64)(unsafe.Add(mBase, uint32(v112)+uint32(_consts[944])))
						*(*int64)(unsafe.Add(mBase, uint32(v120))) = v121 + v124
						v127 = v112 + (v48 + int32(2824))
						v128 = *(*int64)(unsafe.Add(mBase, uint32(v127)))
						v131 = *(*int64)(unsafe.Add(mBase, uint32(v112)+uint32(_consts[945])))
						v133 = base.I64_div_s(v131, int64(1000))
						*(*int64)(unsafe.Add(mBase, uint32(v127))) = v128 + v133
						v137 = v94 + int32(1)
						if v137 != int32(8) {
							v94 = v137
							continue
						} else {
							break
						}
						break
					}
					v140 = v90 + (v48 + int32(1928))
					v141 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
					v144 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[946])))
					*(*int64)(unsafe.Add(mBase, uint32(v140))) = v141 + v144
					v147 = v90 + v50
					v149 = v147 - int32(-64)
					v150 = *(*int64)(unsafe.Add(mBase, uint32(v149)))
					v153 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[947])))
					*(*int64)(unsafe.Add(mBase, uint32(v149))) = v150 + v153
					v156 = v90 + (v48 + int32(2888))
					v157 = *(*int64)(unsafe.Add(mBase, uint32(v156)))
					v160 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[948])))
					v161 = int64(1000)
					v162 = base.I64_div_s(v160, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v156))) = v157 + v162
					v165 = int32(8)
					v166 = v140 + v165
					v167 = *(*int64)(unsafe.Add(mBase, uint32(v166)))
					v170 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[949])))
					*(*int64)(unsafe.Add(mBase, uint32(v166))) = v167 + v170
					v174 = v147 + int32(72)
					v175 = *(*int64)(unsafe.Add(mBase, uint32(v174)))
					v178 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[950])))
					*(*int64)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
					v182 = v156 + v165
					v183 = *(*int64)(unsafe.Add(mBase, uint32(v182)))
					v186 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[951])))
					v188 = base.I64_div_s(v186, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v182))) = v183 + v188
					v191 = int32(16)
					v192 = v140 + v191
					v193 = *(*int64)(unsafe.Add(mBase, uint32(v192)))
					v196 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[952])))
					*(*int64)(unsafe.Add(mBase, uint32(v192))) = v193 + v196
					v200 = v147 + int32(80)
					v201 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
					v204 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[953])))
					*(*int64)(unsafe.Add(mBase, uint32(v200))) = v201 + v204
					v208 = v156 + v191
					v209 = *(*int64)(unsafe.Add(mBase, uint32(v208)))
					v212 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[954])))
					v214 = base.I64_div_s(v212, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v208))) = v209 + v214
					v217 = int32(24)
					v218 = v140 + v217
					v219 = *(*int64)(unsafe.Add(mBase, uint32(v218)))
					v222 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[955])))
					*(*int64)(unsafe.Add(mBase, uint32(v218))) = v219 + v222
					v226 = v147 + int32(88)
					v227 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
					v230 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[956])))
					*(*int64)(unsafe.Add(mBase, uint32(v226))) = v227 + v230
					v234 = v156 + v217
					v235 = *(*int64)(unsafe.Add(mBase, uint32(v234)))
					v238 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[957])))
					v240 = base.I64_div_s(v238, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v234))) = v235 + v240
					v243 = int32(32)
					v244 = v140 + v243
					v245 = *(*int64)(unsafe.Add(mBase, uint32(v244)))
					v248 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[958])))
					*(*int64)(unsafe.Add(mBase, uint32(v244))) = v245 + v248
					v252 = v147 + int32(96)
					v253 = *(*int64)(unsafe.Add(mBase, uint32(v252)))
					v256 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[959])))
					*(*int64)(unsafe.Add(mBase, uint32(v252))) = v253 + v256
					v260 = v156 + v243
					v261 = *(*int64)(unsafe.Add(mBase, uint32(v260)))
					v264 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[960])))
					v266 = base.I64_div_s(v264, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v260))) = v261 + v266
					v269 = int32(40)
					v270 = v140 + v269
					v271 = *(*int64)(unsafe.Add(mBase, uint32(v270)))
					v274 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[961])))
					*(*int64)(unsafe.Add(mBase, uint32(v270))) = v271 + v274
					v278 = v147 + int32(104)
					v279 = *(*int64)(unsafe.Add(mBase, uint32(v278)))
					v282 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[962])))
					*(*int64)(unsafe.Add(mBase, uint32(v278))) = v279 + v282
					v286 = v156 + v269
					v287 = *(*int64)(unsafe.Add(mBase, uint32(v286)))
					v290 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[963])))
					v292 = base.I64_div_s(v290, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v286))) = v287 + v292
					v295 = int32(48)
					v296 = v140 + v295
					v297 = *(*int64)(unsafe.Add(mBase, uint32(v296)))
					v300 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[964])))
					*(*int64)(unsafe.Add(mBase, uint32(v296))) = v297 + v300
					v304 = v147 + int32(112)
					v305 = *(*int64)(unsafe.Add(mBase, uint32(v304)))
					v308 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[965])))
					*(*int64)(unsafe.Add(mBase, uint32(v304))) = v305 + v308
					v312 = v156 + v295
					v313 = *(*int64)(unsafe.Add(mBase, uint32(v312)))
					v316 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[966])))
					v318 = base.I64_div_s(v316, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v312))) = v313 + v318
					v321 = int32(56)
					v322 = v140 + v321
					v323 = *(*int64)(unsafe.Add(mBase, uint32(v322)))
					v326 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[967])))
					*(*int64)(unsafe.Add(mBase, uint32(v322))) = v323 + v326
					v330 = v147 + int32(120)
					v331 = *(*int64)(unsafe.Add(mBase, uint32(v330)))
					v334 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[968])))
					*(*int64)(unsafe.Add(mBase, uint32(v330))) = v331 + v334
					v338 = v156 + v321
					v339 = *(*int64)(unsafe.Add(mBase, uint32(v338)))
					v342 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[969])))
					v344 = base.I64_div_s(v342, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v338))) = v339 + v344
					v347 = v90 + (v48 + int32(1992))
					v348 = *(*int64)(unsafe.Add(mBase, uint32(v347)))
					v351 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[970])))
					*(*int64)(unsafe.Add(mBase, uint32(v347))) = v348 + v351
					v355 = v147 + int32(128)
					v356 = *(*int64)(unsafe.Add(mBase, uint32(v355)))
					v359 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[971])))
					*(*int64)(unsafe.Add(mBase, uint32(v355))) = v356 + v359
					v362 = v90 + (v48 + int32(2952))
					v363 = *(*int64)(unsafe.Add(mBase, uint32(v362)))
					v366 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[972])))
					v368 = base.I64_div_s(v366, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v362))) = v363 + v368
					v372 = v347 + v165
					v373 = *(*int64)(unsafe.Add(mBase, uint32(v372)))
					v376 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[973])))
					*(*int64)(unsafe.Add(mBase, uint32(v372))) = v373 + v376
					v380 = v147 + int32(136)
					v381 = *(*int64)(unsafe.Add(mBase, uint32(v380)))
					v384 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[974])))
					*(*int64)(unsafe.Add(mBase, uint32(v380))) = v381 + v384
					v388 = v362 + v165
					v389 = *(*int64)(unsafe.Add(mBase, uint32(v388)))
					v392 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[975])))
					v394 = base.I64_div_s(v392, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v388))) = v389 + v394
					v398 = v347 + v191
					v399 = *(*int64)(unsafe.Add(mBase, uint32(v398)))
					v402 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[976])))
					*(*int64)(unsafe.Add(mBase, uint32(v398))) = v399 + v402
					v406 = v147 + int32(144)
					v407 = *(*int64)(unsafe.Add(mBase, uint32(v406)))
					v410 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[977])))
					*(*int64)(unsafe.Add(mBase, uint32(v406))) = v407 + v410
					v414 = v362 + v191
					v415 = *(*int64)(unsafe.Add(mBase, uint32(v414)))
					v418 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[978])))
					v420 = base.I64_div_s(v418, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v414))) = v415 + v420
					v424 = v347 + v217
					v425 = *(*int64)(unsafe.Add(mBase, uint32(v424)))
					v428 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[979])))
					*(*int64)(unsafe.Add(mBase, uint32(v424))) = v425 + v428
					v432 = v147 + int32(152)
					v433 = *(*int64)(unsafe.Add(mBase, uint32(v432)))
					v436 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[980])))
					*(*int64)(unsafe.Add(mBase, uint32(v432))) = v433 + v436
					v440 = v362 + v217
					v441 = *(*int64)(unsafe.Add(mBase, uint32(v440)))
					v444 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[981])))
					v446 = base.I64_div_s(v444, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v440))) = v441 + v446
					v450 = v347 + v243
					v451 = *(*int64)(unsafe.Add(mBase, uint32(v450)))
					v454 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[982])))
					*(*int64)(unsafe.Add(mBase, uint32(v450))) = v451 + v454
					v458 = v147 + int32(160)
					v459 = *(*int64)(unsafe.Add(mBase, uint32(v458)))
					v462 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[983])))
					*(*int64)(unsafe.Add(mBase, uint32(v458))) = v459 + v462
					v466 = v362 + v243
					v467 = *(*int64)(unsafe.Add(mBase, uint32(v466)))
					v470 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[984])))
					v472 = base.I64_div_s(v470, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v466))) = v467 + v472
					v476 = v347 + v269
					v477 = *(*int64)(unsafe.Add(mBase, uint32(v476)))
					v480 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[985])))
					*(*int64)(unsafe.Add(mBase, uint32(v476))) = v477 + v480
					v484 = v147 + int32(168)
					v485 = *(*int64)(unsafe.Add(mBase, uint32(v484)))
					v488 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[986])))
					*(*int64)(unsafe.Add(mBase, uint32(v484))) = v485 + v488
					v492 = v362 + v269
					v493 = *(*int64)(unsafe.Add(mBase, uint32(v492)))
					v496 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[987])))
					v498 = base.I64_div_s(v496, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v492))) = v493 + v498
					v502 = v347 + v295
					v503 = *(*int64)(unsafe.Add(mBase, uint32(v502)))
					v506 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[988])))
					*(*int64)(unsafe.Add(mBase, uint32(v502))) = v503 + v506
					v510 = v147 + int32(176)
					v511 = *(*int64)(unsafe.Add(mBase, uint32(v510)))
					v514 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[989])))
					*(*int64)(unsafe.Add(mBase, uint32(v510))) = v511 + v514
					v518 = v362 + v295
					v519 = *(*int64)(unsafe.Add(mBase, uint32(v518)))
					v522 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[990])))
					v524 = base.I64_div_s(v522, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v518))) = v519 + v524
					v528 = v347 + v321
					v529 = *(*int64)(unsafe.Add(mBase, uint32(v528)))
					v532 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[991])))
					*(*int64)(unsafe.Add(mBase, uint32(v528))) = v529 + v532
					v536 = v147 + int32(184)
					v537 = *(*int64)(unsafe.Add(mBase, uint32(v536)))
					v540 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[992])))
					*(*int64)(unsafe.Add(mBase, uint32(v536))) = v537 + v540
					v544 = v362 + v321
					v545 = *(*int64)(unsafe.Add(mBase, uint32(v544)))
					v548 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[993])))
					v550 = base.I64_div_s(v548, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v544))) = v545 + v550
					v553 = v90 + (v48 + int32(2056))
					v554 = *(*int64)(unsafe.Add(mBase, uint32(v553)))
					v557 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[994])))
					*(*int64)(unsafe.Add(mBase, uint32(v553))) = v554 + v557
					v561 = v147 + int32(192)
					v562 = *(*int64)(unsafe.Add(mBase, uint32(v561)))
					v565 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[995])))
					*(*int64)(unsafe.Add(mBase, uint32(v561))) = v562 + v565
					v568 = v90 + (v48 + int32(3016))
					v569 = *(*int64)(unsafe.Add(mBase, uint32(v568)))
					v572 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[996])))
					v574 = base.I64_div_s(v572, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v568))) = v569 + v574
					v578 = v553 + v165
					v579 = *(*int64)(unsafe.Add(mBase, uint32(v578)))
					v582 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[997])))
					*(*int64)(unsafe.Add(mBase, uint32(v578))) = v579 + v582
					v586 = v147 + int32(200)
					v587 = *(*int64)(unsafe.Add(mBase, uint32(v586)))
					v590 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[998])))
					*(*int64)(unsafe.Add(mBase, uint32(v586))) = v587 + v590
					v594 = v568 + v165
					v595 = *(*int64)(unsafe.Add(mBase, uint32(v594)))
					v598 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[999])))
					v600 = base.I64_div_s(v598, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v594))) = v595 + v600
					v604 = v553 + v191
					v605 = *(*int64)(unsafe.Add(mBase, uint32(v604)))
					v608 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1000])))
					*(*int64)(unsafe.Add(mBase, uint32(v604))) = v605 + v608
					v612 = v147 + int32(208)
					v613 = *(*int64)(unsafe.Add(mBase, uint32(v612)))
					v616 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1001])))
					*(*int64)(unsafe.Add(mBase, uint32(v612))) = v613 + v616
					v620 = v568 + v191
					v621 = *(*int64)(unsafe.Add(mBase, uint32(v620)))
					v624 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1002])))
					v626 = base.I64_div_s(v624, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v620))) = v621 + v626
					v630 = v553 + v217
					v631 = *(*int64)(unsafe.Add(mBase, uint32(v630)))
					v634 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1003])))
					*(*int64)(unsafe.Add(mBase, uint32(v630))) = v631 + v634
					v638 = v147 + int32(216)
					v639 = *(*int64)(unsafe.Add(mBase, uint32(v638)))
					v642 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1004])))
					*(*int64)(unsafe.Add(mBase, uint32(v638))) = v639 + v642
					v646 = v568 + v217
					v647 = *(*int64)(unsafe.Add(mBase, uint32(v646)))
					v650 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1005])))
					v652 = base.I64_div_s(v650, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v646))) = v647 + v652
					v656 = v553 + v243
					v657 = *(*int64)(unsafe.Add(mBase, uint32(v656)))
					v660 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1006])))
					*(*int64)(unsafe.Add(mBase, uint32(v656))) = v657 + v660
					v664 = v147 + int32(224)
					v665 = *(*int64)(unsafe.Add(mBase, uint32(v664)))
					v668 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1007])))
					*(*int64)(unsafe.Add(mBase, uint32(v664))) = v665 + v668
					v672 = v568 + v243
					v673 = *(*int64)(unsafe.Add(mBase, uint32(v672)))
					v676 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1008])))
					v678 = base.I64_div_s(v676, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v672))) = v673 + v678
					v682 = v553 + v269
					v683 = *(*int64)(unsafe.Add(mBase, uint32(v682)))
					v686 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1009])))
					*(*int64)(unsafe.Add(mBase, uint32(v682))) = v683 + v686
					v690 = v147 + int32(232)
					v691 = *(*int64)(unsafe.Add(mBase, uint32(v690)))
					v694 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1010])))
					*(*int64)(unsafe.Add(mBase, uint32(v690))) = v691 + v694
					v698 = v568 + v269
					v699 = *(*int64)(unsafe.Add(mBase, uint32(v698)))
					v702 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1011])))
					v704 = base.I64_div_s(v702, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v698))) = v699 + v704
					v708 = v553 + v295
					v709 = *(*int64)(unsafe.Add(mBase, uint32(v708)))
					v712 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1012])))
					*(*int64)(unsafe.Add(mBase, uint32(v708))) = v709 + v712
					v716 = v147 + int32(240)
					v717 = *(*int64)(unsafe.Add(mBase, uint32(v716)))
					v720 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1013])))
					*(*int64)(unsafe.Add(mBase, uint32(v716))) = v717 + v720
					v724 = v568 + v295
					v725 = *(*int64)(unsafe.Add(mBase, uint32(v724)))
					v728 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1014])))
					v730 = base.I64_div_s(v728, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v724))) = v725 + v730
					v734 = v553 + v321
					v735 = *(*int64)(unsafe.Add(mBase, uint32(v734)))
					v738 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1015])))
					*(*int64)(unsafe.Add(mBase, uint32(v734))) = v735 + v738
					v742 = v147 + int32(248)
					v743 = *(*int64)(unsafe.Add(mBase, uint32(v742)))
					v746 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1016])))
					*(*int64)(unsafe.Add(mBase, uint32(v742))) = v743 + v746
					v750 = v568 + v321
					v751 = *(*int64)(unsafe.Add(mBase, uint32(v750)))
					v754 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1017])))
					v756 = base.I64_div_s(v754, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v750))) = v751 + v756
					v759 = v90 + (v48 + int32(2120))
					v760 = *(*int64)(unsafe.Add(mBase, uint32(v759)))
					v763 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1018])))
					*(*int64)(unsafe.Add(mBase, uint32(v759))) = v760 + v763
					v767 = v147 + int32(256)
					v768 = *(*int64)(unsafe.Add(mBase, uint32(v767)))
					v771 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1019])))
					*(*int64)(unsafe.Add(mBase, uint32(v767))) = v768 + v771
					v774 = v90 + (v48 + int32(3080))
					v775 = *(*int64)(unsafe.Add(mBase, uint32(v774)))
					v778 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1020])))
					v780 = base.I64_div_s(v778, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v774))) = v775 + v780
					v784 = v759 + v165
					v785 = *(*int64)(unsafe.Add(mBase, uint32(v784)))
					v788 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1021])))
					*(*int64)(unsafe.Add(mBase, uint32(v784))) = v785 + v788
					v792 = v147 + int32(264)
					v793 = *(*int64)(unsafe.Add(mBase, uint32(v792)))
					v796 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1022])))
					*(*int64)(unsafe.Add(mBase, uint32(v792))) = v793 + v796
					v800 = v774 + v165
					v801 = *(*int64)(unsafe.Add(mBase, uint32(v800)))
					v804 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1023])))
					v806 = base.I64_div_s(v804, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v800))) = v801 + v806
					v810 = v759 + v191
					v811 = *(*int64)(unsafe.Add(mBase, uint32(v810)))
					v814 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1024])))
					*(*int64)(unsafe.Add(mBase, uint32(v810))) = v811 + v814
					v818 = v147 + int32(272)
					v819 = *(*int64)(unsafe.Add(mBase, uint32(v818)))
					v822 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1025])))
					*(*int64)(unsafe.Add(mBase, uint32(v818))) = v819 + v822
					v826 = v774 + v191
					v827 = *(*int64)(unsafe.Add(mBase, uint32(v826)))
					v830 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1026])))
					v832 = base.I64_div_s(v830, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v826))) = v827 + v832
					v836 = v759 + v217
					v837 = *(*int64)(unsafe.Add(mBase, uint32(v836)))
					v840 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1027])))
					*(*int64)(unsafe.Add(mBase, uint32(v836))) = v837 + v840
					v844 = v147 + int32(280)
					v845 = *(*int64)(unsafe.Add(mBase, uint32(v844)))
					v848 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1028])))
					*(*int64)(unsafe.Add(mBase, uint32(v844))) = v845 + v848
					v852 = v774 + v217
					v853 = *(*int64)(unsafe.Add(mBase, uint32(v852)))
					v856 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1029])))
					v858 = base.I64_div_s(v856, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v852))) = v853 + v858
					v862 = v759 + v243
					v863 = *(*int64)(unsafe.Add(mBase, uint32(v862)))
					v866 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1030])))
					*(*int64)(unsafe.Add(mBase, uint32(v862))) = v863 + v866
					v870 = v147 + int32(288)
					v871 = *(*int64)(unsafe.Add(mBase, uint32(v870)))
					v874 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1031])))
					*(*int64)(unsafe.Add(mBase, uint32(v870))) = v871 + v874
					v878 = v774 + v243
					v879 = *(*int64)(unsafe.Add(mBase, uint32(v878)))
					v882 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1032])))
					v884 = base.I64_div_s(v882, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v878))) = v879 + v884
					v888 = v759 + v269
					v889 = *(*int64)(unsafe.Add(mBase, uint32(v888)))
					v892 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1033])))
					*(*int64)(unsafe.Add(mBase, uint32(v888))) = v889 + v892
					v896 = v147 + int32(296)
					v897 = *(*int64)(unsafe.Add(mBase, uint32(v896)))
					v900 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1034])))
					*(*int64)(unsafe.Add(mBase, uint32(v896))) = v897 + v900
					v904 = v774 + v269
					v905 = *(*int64)(unsafe.Add(mBase, uint32(v904)))
					v908 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1035])))
					v910 = base.I64_div_s(v908, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v904))) = v905 + v910
					v914 = v759 + v295
					v915 = *(*int64)(unsafe.Add(mBase, uint32(v914)))
					v918 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1036])))
					*(*int64)(unsafe.Add(mBase, uint32(v914))) = v915 + v918
					v922 = v147 + int32(304)
					v923 = *(*int64)(unsafe.Add(mBase, uint32(v922)))
					v926 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1037])))
					*(*int64)(unsafe.Add(mBase, uint32(v922))) = v923 + v926
					v930 = v774 + v295
					v931 = *(*int64)(unsafe.Add(mBase, uint32(v930)))
					v934 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1038])))
					v936 = base.I64_div_s(v934, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v930))) = v931 + v936
					v940 = v759 + v321
					v941 = *(*int64)(unsafe.Add(mBase, uint32(v940)))
					v944 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1039])))
					*(*int64)(unsafe.Add(mBase, uint32(v940))) = v941 + v944
					v948 = v147 + int32(312)
					v949 = *(*int64)(unsafe.Add(mBase, uint32(v948)))
					v952 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1040])))
					*(*int64)(unsafe.Add(mBase, uint32(v948))) = v949 + v952
					v956 = v774 + v321
					v957 = *(*int64)(unsafe.Add(mBase, uint32(v956)))
					v960 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1041])))
					v962 = base.I64_div_s(v960, v161)
					*(*int64)(unsafe.Add(mBase, uint32(v956))) = v957 + v962
					v966 = v78 + int32(1)
					if v966 != int32(3) {
						v78 = v966
						continue
					} else {
						break
					}
					break
				}
				F_LWLockRelease(m, v31)
				mBase = m.M
				v970 = m.ExcPending
				if v970 != 0 {
					return int32(0)
				} else {
					v971 = int32(0)
					v976 = F__emscripten_memset_bulkmem(m, int32(4405064), base.I32_extend8_s(v971), int32(2880))
					mBase = m.M
					v978 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v978)
					v984 = v971
					return v984
				}
			}
		} else {
			v41 = F_LWLockConditionalAcquire(m, v31, int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				if v41 == int32(0) {
					v984 = int32(1)
					return v984
				} else {
					v48 = v24 + v26*int32(2880)
					v50 = v48 + int32(904)
					v78 = v2
					for {
						v90 = v78 * int32(320)
						v94 = int32(0)
						for {
							v112 = v90 + v94<<(uint(int32(3))%32)
							v113 = v112 + (v48 + int32(1864))
							v114 = *(*int64)(unsafe.Add(mBase, uint32(v113)))
							v117 = *(*int64)(unsafe.Add(mBase, uint32(v112)+uint32(_consts[943])))
							*(*int64)(unsafe.Add(mBase, uint32(v113))) = v114 + v117
							v120 = v112 + v50
							v121 = *(*int64)(unsafe.Add(mBase, uint32(v120)))
							v124 = *(*int64)(unsafe.Add(mBase, uint32(v112)+uint32(_consts[944])))
							*(*int64)(unsafe.Add(mBase, uint32(v120))) = v121 + v124
							v127 = v112 + (v48 + int32(2824))
							v128 = *(*int64)(unsafe.Add(mBase, uint32(v127)))
							v131 = *(*int64)(unsafe.Add(mBase, uint32(v112)+uint32(_consts[945])))
							v133 = base.I64_div_s(v131, int64(1000))
							*(*int64)(unsafe.Add(mBase, uint32(v127))) = v128 + v133
							v137 = v94 + int32(1)
							if v137 != int32(8) {
								v94 = v137
								continue
							} else {
								break
							}
							break
						}
						v140 = v90 + (v48 + int32(1928))
						v141 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
						v144 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[946])))
						*(*int64)(unsafe.Add(mBase, uint32(v140))) = v141 + v144
						v147 = v90 + v50
						v149 = v147 - int32(-64)
						v150 = *(*int64)(unsafe.Add(mBase, uint32(v149)))
						v153 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[947])))
						*(*int64)(unsafe.Add(mBase, uint32(v149))) = v150 + v153
						v156 = v90 + (v48 + int32(2888))
						v157 = *(*int64)(unsafe.Add(mBase, uint32(v156)))
						v160 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[948])))
						v161 = int64(1000)
						v162 = base.I64_div_s(v160, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v156))) = v157 + v162
						v165 = int32(8)
						v166 = v140 + v165
						v167 = *(*int64)(unsafe.Add(mBase, uint32(v166)))
						v170 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[949])))
						*(*int64)(unsafe.Add(mBase, uint32(v166))) = v167 + v170
						v174 = v147 + int32(72)
						v175 = *(*int64)(unsafe.Add(mBase, uint32(v174)))
						v178 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[950])))
						*(*int64)(unsafe.Add(mBase, uint32(v174))) = v175 + v178
						v182 = v156 + v165
						v183 = *(*int64)(unsafe.Add(mBase, uint32(v182)))
						v186 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[951])))
						v188 = base.I64_div_s(v186, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v182))) = v183 + v188
						v191 = int32(16)
						v192 = v140 + v191
						v193 = *(*int64)(unsafe.Add(mBase, uint32(v192)))
						v196 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[952])))
						*(*int64)(unsafe.Add(mBase, uint32(v192))) = v193 + v196
						v200 = v147 + int32(80)
						v201 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
						v204 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[953])))
						*(*int64)(unsafe.Add(mBase, uint32(v200))) = v201 + v204
						v208 = v156 + v191
						v209 = *(*int64)(unsafe.Add(mBase, uint32(v208)))
						v212 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[954])))
						v214 = base.I64_div_s(v212, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v208))) = v209 + v214
						v217 = int32(24)
						v218 = v140 + v217
						v219 = *(*int64)(unsafe.Add(mBase, uint32(v218)))
						v222 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[955])))
						*(*int64)(unsafe.Add(mBase, uint32(v218))) = v219 + v222
						v226 = v147 + int32(88)
						v227 = *(*int64)(unsafe.Add(mBase, uint32(v226)))
						v230 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[956])))
						*(*int64)(unsafe.Add(mBase, uint32(v226))) = v227 + v230
						v234 = v156 + v217
						v235 = *(*int64)(unsafe.Add(mBase, uint32(v234)))
						v238 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[957])))
						v240 = base.I64_div_s(v238, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v234))) = v235 + v240
						v243 = int32(32)
						v244 = v140 + v243
						v245 = *(*int64)(unsafe.Add(mBase, uint32(v244)))
						v248 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[958])))
						*(*int64)(unsafe.Add(mBase, uint32(v244))) = v245 + v248
						v252 = v147 + int32(96)
						v253 = *(*int64)(unsafe.Add(mBase, uint32(v252)))
						v256 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[959])))
						*(*int64)(unsafe.Add(mBase, uint32(v252))) = v253 + v256
						v260 = v156 + v243
						v261 = *(*int64)(unsafe.Add(mBase, uint32(v260)))
						v264 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[960])))
						v266 = base.I64_div_s(v264, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v260))) = v261 + v266
						v269 = int32(40)
						v270 = v140 + v269
						v271 = *(*int64)(unsafe.Add(mBase, uint32(v270)))
						v274 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[961])))
						*(*int64)(unsafe.Add(mBase, uint32(v270))) = v271 + v274
						v278 = v147 + int32(104)
						v279 = *(*int64)(unsafe.Add(mBase, uint32(v278)))
						v282 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[962])))
						*(*int64)(unsafe.Add(mBase, uint32(v278))) = v279 + v282
						v286 = v156 + v269
						v287 = *(*int64)(unsafe.Add(mBase, uint32(v286)))
						v290 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[963])))
						v292 = base.I64_div_s(v290, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v286))) = v287 + v292
						v295 = int32(48)
						v296 = v140 + v295
						v297 = *(*int64)(unsafe.Add(mBase, uint32(v296)))
						v300 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[964])))
						*(*int64)(unsafe.Add(mBase, uint32(v296))) = v297 + v300
						v304 = v147 + int32(112)
						v305 = *(*int64)(unsafe.Add(mBase, uint32(v304)))
						v308 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[965])))
						*(*int64)(unsafe.Add(mBase, uint32(v304))) = v305 + v308
						v312 = v156 + v295
						v313 = *(*int64)(unsafe.Add(mBase, uint32(v312)))
						v316 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[966])))
						v318 = base.I64_div_s(v316, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v312))) = v313 + v318
						v321 = int32(56)
						v322 = v140 + v321
						v323 = *(*int64)(unsafe.Add(mBase, uint32(v322)))
						v326 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[967])))
						*(*int64)(unsafe.Add(mBase, uint32(v322))) = v323 + v326
						v330 = v147 + int32(120)
						v331 = *(*int64)(unsafe.Add(mBase, uint32(v330)))
						v334 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[968])))
						*(*int64)(unsafe.Add(mBase, uint32(v330))) = v331 + v334
						v338 = v156 + v321
						v339 = *(*int64)(unsafe.Add(mBase, uint32(v338)))
						v342 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[969])))
						v344 = base.I64_div_s(v342, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v338))) = v339 + v344
						v347 = v90 + (v48 + int32(1992))
						v348 = *(*int64)(unsafe.Add(mBase, uint32(v347)))
						v351 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[970])))
						*(*int64)(unsafe.Add(mBase, uint32(v347))) = v348 + v351
						v355 = v147 + int32(128)
						v356 = *(*int64)(unsafe.Add(mBase, uint32(v355)))
						v359 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[971])))
						*(*int64)(unsafe.Add(mBase, uint32(v355))) = v356 + v359
						v362 = v90 + (v48 + int32(2952))
						v363 = *(*int64)(unsafe.Add(mBase, uint32(v362)))
						v366 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[972])))
						v368 = base.I64_div_s(v366, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v362))) = v363 + v368
						v372 = v347 + v165
						v373 = *(*int64)(unsafe.Add(mBase, uint32(v372)))
						v376 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[973])))
						*(*int64)(unsafe.Add(mBase, uint32(v372))) = v373 + v376
						v380 = v147 + int32(136)
						v381 = *(*int64)(unsafe.Add(mBase, uint32(v380)))
						v384 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[974])))
						*(*int64)(unsafe.Add(mBase, uint32(v380))) = v381 + v384
						v388 = v362 + v165
						v389 = *(*int64)(unsafe.Add(mBase, uint32(v388)))
						v392 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[975])))
						v394 = base.I64_div_s(v392, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v388))) = v389 + v394
						v398 = v347 + v191
						v399 = *(*int64)(unsafe.Add(mBase, uint32(v398)))
						v402 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[976])))
						*(*int64)(unsafe.Add(mBase, uint32(v398))) = v399 + v402
						v406 = v147 + int32(144)
						v407 = *(*int64)(unsafe.Add(mBase, uint32(v406)))
						v410 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[977])))
						*(*int64)(unsafe.Add(mBase, uint32(v406))) = v407 + v410
						v414 = v362 + v191
						v415 = *(*int64)(unsafe.Add(mBase, uint32(v414)))
						v418 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[978])))
						v420 = base.I64_div_s(v418, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v414))) = v415 + v420
						v424 = v347 + v217
						v425 = *(*int64)(unsafe.Add(mBase, uint32(v424)))
						v428 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[979])))
						*(*int64)(unsafe.Add(mBase, uint32(v424))) = v425 + v428
						v432 = v147 + int32(152)
						v433 = *(*int64)(unsafe.Add(mBase, uint32(v432)))
						v436 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[980])))
						*(*int64)(unsafe.Add(mBase, uint32(v432))) = v433 + v436
						v440 = v362 + v217
						v441 = *(*int64)(unsafe.Add(mBase, uint32(v440)))
						v444 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[981])))
						v446 = base.I64_div_s(v444, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v440))) = v441 + v446
						v450 = v347 + v243
						v451 = *(*int64)(unsafe.Add(mBase, uint32(v450)))
						v454 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[982])))
						*(*int64)(unsafe.Add(mBase, uint32(v450))) = v451 + v454
						v458 = v147 + int32(160)
						v459 = *(*int64)(unsafe.Add(mBase, uint32(v458)))
						v462 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[983])))
						*(*int64)(unsafe.Add(mBase, uint32(v458))) = v459 + v462
						v466 = v362 + v243
						v467 = *(*int64)(unsafe.Add(mBase, uint32(v466)))
						v470 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[984])))
						v472 = base.I64_div_s(v470, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v466))) = v467 + v472
						v476 = v347 + v269
						v477 = *(*int64)(unsafe.Add(mBase, uint32(v476)))
						v480 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[985])))
						*(*int64)(unsafe.Add(mBase, uint32(v476))) = v477 + v480
						v484 = v147 + int32(168)
						v485 = *(*int64)(unsafe.Add(mBase, uint32(v484)))
						v488 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[986])))
						*(*int64)(unsafe.Add(mBase, uint32(v484))) = v485 + v488
						v492 = v362 + v269
						v493 = *(*int64)(unsafe.Add(mBase, uint32(v492)))
						v496 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[987])))
						v498 = base.I64_div_s(v496, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v492))) = v493 + v498
						v502 = v347 + v295
						v503 = *(*int64)(unsafe.Add(mBase, uint32(v502)))
						v506 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[988])))
						*(*int64)(unsafe.Add(mBase, uint32(v502))) = v503 + v506
						v510 = v147 + int32(176)
						v511 = *(*int64)(unsafe.Add(mBase, uint32(v510)))
						v514 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[989])))
						*(*int64)(unsafe.Add(mBase, uint32(v510))) = v511 + v514
						v518 = v362 + v295
						v519 = *(*int64)(unsafe.Add(mBase, uint32(v518)))
						v522 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[990])))
						v524 = base.I64_div_s(v522, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v518))) = v519 + v524
						v528 = v347 + v321
						v529 = *(*int64)(unsafe.Add(mBase, uint32(v528)))
						v532 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[991])))
						*(*int64)(unsafe.Add(mBase, uint32(v528))) = v529 + v532
						v536 = v147 + int32(184)
						v537 = *(*int64)(unsafe.Add(mBase, uint32(v536)))
						v540 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[992])))
						*(*int64)(unsafe.Add(mBase, uint32(v536))) = v537 + v540
						v544 = v362 + v321
						v545 = *(*int64)(unsafe.Add(mBase, uint32(v544)))
						v548 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[993])))
						v550 = base.I64_div_s(v548, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v544))) = v545 + v550
						v553 = v90 + (v48 + int32(2056))
						v554 = *(*int64)(unsafe.Add(mBase, uint32(v553)))
						v557 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[994])))
						*(*int64)(unsafe.Add(mBase, uint32(v553))) = v554 + v557
						v561 = v147 + int32(192)
						v562 = *(*int64)(unsafe.Add(mBase, uint32(v561)))
						v565 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[995])))
						*(*int64)(unsafe.Add(mBase, uint32(v561))) = v562 + v565
						v568 = v90 + (v48 + int32(3016))
						v569 = *(*int64)(unsafe.Add(mBase, uint32(v568)))
						v572 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[996])))
						v574 = base.I64_div_s(v572, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v568))) = v569 + v574
						v578 = v553 + v165
						v579 = *(*int64)(unsafe.Add(mBase, uint32(v578)))
						v582 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[997])))
						*(*int64)(unsafe.Add(mBase, uint32(v578))) = v579 + v582
						v586 = v147 + int32(200)
						v587 = *(*int64)(unsafe.Add(mBase, uint32(v586)))
						v590 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[998])))
						*(*int64)(unsafe.Add(mBase, uint32(v586))) = v587 + v590
						v594 = v568 + v165
						v595 = *(*int64)(unsafe.Add(mBase, uint32(v594)))
						v598 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[999])))
						v600 = base.I64_div_s(v598, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v594))) = v595 + v600
						v604 = v553 + v191
						v605 = *(*int64)(unsafe.Add(mBase, uint32(v604)))
						v608 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1000])))
						*(*int64)(unsafe.Add(mBase, uint32(v604))) = v605 + v608
						v612 = v147 + int32(208)
						v613 = *(*int64)(unsafe.Add(mBase, uint32(v612)))
						v616 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1001])))
						*(*int64)(unsafe.Add(mBase, uint32(v612))) = v613 + v616
						v620 = v568 + v191
						v621 = *(*int64)(unsafe.Add(mBase, uint32(v620)))
						v624 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1002])))
						v626 = base.I64_div_s(v624, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v620))) = v621 + v626
						v630 = v553 + v217
						v631 = *(*int64)(unsafe.Add(mBase, uint32(v630)))
						v634 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1003])))
						*(*int64)(unsafe.Add(mBase, uint32(v630))) = v631 + v634
						v638 = v147 + int32(216)
						v639 = *(*int64)(unsafe.Add(mBase, uint32(v638)))
						v642 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1004])))
						*(*int64)(unsafe.Add(mBase, uint32(v638))) = v639 + v642
						v646 = v568 + v217
						v647 = *(*int64)(unsafe.Add(mBase, uint32(v646)))
						v650 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1005])))
						v652 = base.I64_div_s(v650, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v646))) = v647 + v652
						v656 = v553 + v243
						v657 = *(*int64)(unsafe.Add(mBase, uint32(v656)))
						v660 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1006])))
						*(*int64)(unsafe.Add(mBase, uint32(v656))) = v657 + v660
						v664 = v147 + int32(224)
						v665 = *(*int64)(unsafe.Add(mBase, uint32(v664)))
						v668 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1007])))
						*(*int64)(unsafe.Add(mBase, uint32(v664))) = v665 + v668
						v672 = v568 + v243
						v673 = *(*int64)(unsafe.Add(mBase, uint32(v672)))
						v676 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1008])))
						v678 = base.I64_div_s(v676, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v672))) = v673 + v678
						v682 = v553 + v269
						v683 = *(*int64)(unsafe.Add(mBase, uint32(v682)))
						v686 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1009])))
						*(*int64)(unsafe.Add(mBase, uint32(v682))) = v683 + v686
						v690 = v147 + int32(232)
						v691 = *(*int64)(unsafe.Add(mBase, uint32(v690)))
						v694 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1010])))
						*(*int64)(unsafe.Add(mBase, uint32(v690))) = v691 + v694
						v698 = v568 + v269
						v699 = *(*int64)(unsafe.Add(mBase, uint32(v698)))
						v702 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1011])))
						v704 = base.I64_div_s(v702, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v698))) = v699 + v704
						v708 = v553 + v295
						v709 = *(*int64)(unsafe.Add(mBase, uint32(v708)))
						v712 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1012])))
						*(*int64)(unsafe.Add(mBase, uint32(v708))) = v709 + v712
						v716 = v147 + int32(240)
						v717 = *(*int64)(unsafe.Add(mBase, uint32(v716)))
						v720 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1013])))
						*(*int64)(unsafe.Add(mBase, uint32(v716))) = v717 + v720
						v724 = v568 + v295
						v725 = *(*int64)(unsafe.Add(mBase, uint32(v724)))
						v728 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1014])))
						v730 = base.I64_div_s(v728, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v724))) = v725 + v730
						v734 = v553 + v321
						v735 = *(*int64)(unsafe.Add(mBase, uint32(v734)))
						v738 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1015])))
						*(*int64)(unsafe.Add(mBase, uint32(v734))) = v735 + v738
						v742 = v147 + int32(248)
						v743 = *(*int64)(unsafe.Add(mBase, uint32(v742)))
						v746 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1016])))
						*(*int64)(unsafe.Add(mBase, uint32(v742))) = v743 + v746
						v750 = v568 + v321
						v751 = *(*int64)(unsafe.Add(mBase, uint32(v750)))
						v754 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1017])))
						v756 = base.I64_div_s(v754, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v750))) = v751 + v756
						v759 = v90 + (v48 + int32(2120))
						v760 = *(*int64)(unsafe.Add(mBase, uint32(v759)))
						v763 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1018])))
						*(*int64)(unsafe.Add(mBase, uint32(v759))) = v760 + v763
						v767 = v147 + int32(256)
						v768 = *(*int64)(unsafe.Add(mBase, uint32(v767)))
						v771 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1019])))
						*(*int64)(unsafe.Add(mBase, uint32(v767))) = v768 + v771
						v774 = v90 + (v48 + int32(3080))
						v775 = *(*int64)(unsafe.Add(mBase, uint32(v774)))
						v778 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1020])))
						v780 = base.I64_div_s(v778, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v774))) = v775 + v780
						v784 = v759 + v165
						v785 = *(*int64)(unsafe.Add(mBase, uint32(v784)))
						v788 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1021])))
						*(*int64)(unsafe.Add(mBase, uint32(v784))) = v785 + v788
						v792 = v147 + int32(264)
						v793 = *(*int64)(unsafe.Add(mBase, uint32(v792)))
						v796 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1022])))
						*(*int64)(unsafe.Add(mBase, uint32(v792))) = v793 + v796
						v800 = v774 + v165
						v801 = *(*int64)(unsafe.Add(mBase, uint32(v800)))
						v804 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1023])))
						v806 = base.I64_div_s(v804, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v800))) = v801 + v806
						v810 = v759 + v191
						v811 = *(*int64)(unsafe.Add(mBase, uint32(v810)))
						v814 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1024])))
						*(*int64)(unsafe.Add(mBase, uint32(v810))) = v811 + v814
						v818 = v147 + int32(272)
						v819 = *(*int64)(unsafe.Add(mBase, uint32(v818)))
						v822 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1025])))
						*(*int64)(unsafe.Add(mBase, uint32(v818))) = v819 + v822
						v826 = v774 + v191
						v827 = *(*int64)(unsafe.Add(mBase, uint32(v826)))
						v830 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1026])))
						v832 = base.I64_div_s(v830, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v826))) = v827 + v832
						v836 = v759 + v217
						v837 = *(*int64)(unsafe.Add(mBase, uint32(v836)))
						v840 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1027])))
						*(*int64)(unsafe.Add(mBase, uint32(v836))) = v837 + v840
						v844 = v147 + int32(280)
						v845 = *(*int64)(unsafe.Add(mBase, uint32(v844)))
						v848 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1028])))
						*(*int64)(unsafe.Add(mBase, uint32(v844))) = v845 + v848
						v852 = v774 + v217
						v853 = *(*int64)(unsafe.Add(mBase, uint32(v852)))
						v856 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1029])))
						v858 = base.I64_div_s(v856, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v852))) = v853 + v858
						v862 = v759 + v243
						v863 = *(*int64)(unsafe.Add(mBase, uint32(v862)))
						v866 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1030])))
						*(*int64)(unsafe.Add(mBase, uint32(v862))) = v863 + v866
						v870 = v147 + int32(288)
						v871 = *(*int64)(unsafe.Add(mBase, uint32(v870)))
						v874 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1031])))
						*(*int64)(unsafe.Add(mBase, uint32(v870))) = v871 + v874
						v878 = v774 + v243
						v879 = *(*int64)(unsafe.Add(mBase, uint32(v878)))
						v882 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1032])))
						v884 = base.I64_div_s(v882, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v878))) = v879 + v884
						v888 = v759 + v269
						v889 = *(*int64)(unsafe.Add(mBase, uint32(v888)))
						v892 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1033])))
						*(*int64)(unsafe.Add(mBase, uint32(v888))) = v889 + v892
						v896 = v147 + int32(296)
						v897 = *(*int64)(unsafe.Add(mBase, uint32(v896)))
						v900 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1034])))
						*(*int64)(unsafe.Add(mBase, uint32(v896))) = v897 + v900
						v904 = v774 + v269
						v905 = *(*int64)(unsafe.Add(mBase, uint32(v904)))
						v908 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1035])))
						v910 = base.I64_div_s(v908, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v904))) = v905 + v910
						v914 = v759 + v295
						v915 = *(*int64)(unsafe.Add(mBase, uint32(v914)))
						v918 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1036])))
						*(*int64)(unsafe.Add(mBase, uint32(v914))) = v915 + v918
						v922 = v147 + int32(304)
						v923 = *(*int64)(unsafe.Add(mBase, uint32(v922)))
						v926 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1037])))
						*(*int64)(unsafe.Add(mBase, uint32(v922))) = v923 + v926
						v930 = v774 + v295
						v931 = *(*int64)(unsafe.Add(mBase, uint32(v930)))
						v934 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1038])))
						v936 = base.I64_div_s(v934, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v930))) = v931 + v936
						v940 = v759 + v321
						v941 = *(*int64)(unsafe.Add(mBase, uint32(v940)))
						v944 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1039])))
						*(*int64)(unsafe.Add(mBase, uint32(v940))) = v941 + v944
						v948 = v147 + int32(312)
						v949 = *(*int64)(unsafe.Add(mBase, uint32(v948)))
						v952 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1040])))
						*(*int64)(unsafe.Add(mBase, uint32(v948))) = v949 + v952
						v956 = v774 + v321
						v957 = *(*int64)(unsafe.Add(mBase, uint32(v956)))
						v960 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1041])))
						v962 = base.I64_div_s(v960, v161)
						*(*int64)(unsafe.Add(mBase, uint32(v956))) = v957 + v962
						v966 = v78 + int32(1)
						if v966 != int32(3) {
							v78 = v966
							continue
						} else {
							break
						}
						break
					}
					F_LWLockRelease(m, v31)
					mBase = m.M
					v970 = m.ExcPending
					if v970 != 0 {
						return int32(0)
					} else {
						v971 = int32(0)
						v976 = F__emscripten_memset_bulkmem(m, int32(4405064), base.I32_extend8_s(v971), int32(2880))
						mBase = m.M
						v978 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v978)
						v984 = v971
						return v984
					}
				}
			}
		}
	}
}
func F_pgstat_progress_end_command(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	v3 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v3 == int32(0) {
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
		if v7 != int32(1) {
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+220))
			if v10 == int32(0) {
			} else {
				v13 = int32(4419940)
				v15 = *(*int32)(unsafe.Add(mBase, _consts[11]))
				v16 = int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[11])) = v15 + v16
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
				*(*int32)(unsafe.Add(mBase, uint32(v3))) = v19 + v16
				v23 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v3)+220)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v3)+224)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v3))) = v19 + int32(2)
				v33 = *(*int32)(unsafe.Add(mBase, _consts[11]))
				*(*int32)(unsafe.Add(mBase, _consts[11])) = v33 - v16
			}
		}
	}
	return
}
func F_pgstat_relation_flush_cb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
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
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v174 int64
	_ = v174
	var v179 int64
	_ = v179
	var v183 int64
	_ = v183
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v194 int64
	_ = v194
	var v195 int64
	_ = v195
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v210 int64
	_ = v210
	var v211 int64
	_ = v211
	var v214 int64
	_ = v214
	var v215 int64
	_ = v215
	var v218 int32
	_ = v218
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v224 int64
	_ = v224
	var v227 int32
	_ = v227
	var v234 int64
	_ = v234
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v241 int64
	_ = v241
	var v243 int64
	_ = v243
	var v244 int64
	_ = v244
	var v247 int64
	_ = v247
	var v250 int64
	_ = v250
	var v251 int64
	_ = v251
	var v254 int64
	_ = v254
	var v255 int64
	_ = v255
	var v258 int64
	_ = v258
	var v260 int64
	_ = v260
	var v263 int64
	_ = v263
	var v265 int64
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int64
	_ = v276
	var v277 int64
	_ = v277
	var v280 int64
	_ = v280
	var v281 int64
	_ = v281
	var v284 int64
	_ = v284
	var v285 int64
	_ = v285
	var v288 int64
	_ = v288
	var v289 int64
	_ = v289
	var v292 int64
	_ = v292
	var v293 int64
	_ = v293
	var v296 int64
	_ = v296
	var v297 int64
	_ = v297
	var v300 int64
	_ = v300
	var v301 int64
	_ = v301
	var v310 int32
	_ = v310
	v11 = int64(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = v17 + int32(16)
	if v19&int32(3) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v310
L2:
	;
	v164 = F_pgstat_lock_entry(m, l0, l1)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L29
	} else {
		goto L30
	}
L3:
	;
	v47 = v19 + (int32(-16)-v17)&int32(3)
	v49 = v17 + int32(128)
	v51 = v49 & int32(-4)
	v53 = v51 - int32(28)
	if base.Ui32(v47) < base.Ui32(v53) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+16)))
	if v24 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if (v17+int32(17))&int32(3) == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+17)))
	if v31 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if (v17+int32(18))&int32(3) == int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+18)))
	if v38 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if (v17-int32(1))&int32(3) != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	goto L3
L11:
	;
	v58 = v47
	goto L14
L12:
	;
	v89 = v47
	goto L13
L13:
	;
	if base.Ui32(v89) < base.Ui32(v51) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)+28))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v58)+24))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v68|(v69|(v70|(v71|(v72|(v73|(v74|v75)))))) != 0 {
		goto L2
	} else {
		goto L16
	}
L15:
	;
	v89 = v84
	goto L13
L16:
	;
	v84 = v58 + int32(32)
	if base.Ui32(v84) < base.Ui32(v53) {
		v58 = v84
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v103 = v89
	goto L21
L19:
	;
	v120 = v89
	goto L20
L20:
	;
	v134 = v120
	goto L25
L21:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v113 != 0 {
		goto L2
	} else {
		goto L23
	}
L22:
	;
	v120 = v115
	goto L20
L23:
	;
	v115 = v103 + int32(4)
	if base.Ui32(v115) < base.Ui32(v51) {
		v103 = v115
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	if base.Ui32(v49) <= base.Ui32(v134) {
		v310 = int32(1)
		goto L1
	} else {
		goto L27
	}
L26:
	;
	goto L2
L27:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v145 == int32(0) {
		v134 = v134 + int32(1)
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	return int32(0)
L30:
	;
	if v164 == int32(0) {
		v310 = int32(0)
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v170 + v171
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	if v174 == int64(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v190 + v191
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v14)+48))
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+48)) = v194 + v195
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = v198 + v199
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v14)+64))
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v202 + v203
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v14)+72))
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v17)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v206 + v207
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v14)+80))
	v211 = *(*int64)(unsafe.Add(mBase, uint32(v17)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v210 + v211
	v214 = *(*int64)(unsafe.Add(mBase, uint32(v14)+88))
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v17)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v214 + v215
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+80)))
	if v218 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	v179 = *(*int64)(unsafe.Add(mBase, _consts[170]))
	if v179 == int64(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
	if v185 <= v186 {
		goto L32
	} else {
		goto L38
	}
L35:
	;
	v183 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[170])) = v183
	v185 = v183
	goto L37
L36:
	;
	v185 = v179
	goto L37
L37:
	;
	goto L34
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v185
	goto L32
L39:
	;
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v17)+88))
	v238 = v236 + v237
	*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = v238
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v17)+96))
	v241 = v234 + v240
	*(*int64)(unsafe.Add(mBase, uint32(v14)+104)) = v241
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v14)+112))
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v17)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+112)) = v243 + v244
	v247 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+120)) = v235 + v247
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v14)+128))
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v17)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+128)) = v250 + v251
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v17)+120))
	v255 = int64(0)
	if v255 < v241 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v14)+120))
	v222 = *(*int64)(unsafe.Add(mBase, uint32(v14)+104))
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v14)+96))
	v234 = v222
	v235 = v221
	v236 = v223
	goto L39
L41:
	;
	goto L42
L42:
	;
	v224 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+120)) = v224
	v227 = v14 + int32(96)
	*(*int64)(unsafe.Add(mBase, uint32(v227)+8)) = v224
	*(*int64)(unsafe.Add(mBase, uint32(v227))) = v224
	v234 = v11
	v235 = v11
	v236 = v224
	goto L39
L43:
	;
	v258 = v241
	goto L45
L44:
	;
	v258 = v255
	goto L45
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+104)) = v258
	v260 = int64(0)
	if v260 < v238 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v263 = v238
	goto L48
L47:
	;
	v263 = v260
	goto L48
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = v263
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v14)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+136)) = v254 + v265
	F_pgstat_unlock_entry(m, l0)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	v273 = F_pgstat_prep_pending_entry(m, int32(1), v16, int64(0), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v275)+32))
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+32)) = v276 + v277
	v280 = *(*int64)(unsafe.Add(mBase, uint32(v275)+40))
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+40)) = v280 + v281
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v275)+48))
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+48)) = v284 + v285
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v275)+56))
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+56)) = v288 + v289
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v275)+64))
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v17)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+64)) = v292 + v293
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v275)+16))
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v17)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+16)) = v296 + v297
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v275)+24))
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v17)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+24)) = v300 + v301
	v310 = int32(1)
	goto L1
}
func F_pgstat_report_analyze(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int64) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
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
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v95 int64
	_ = v95
	var v98 int32
	_ = v98
	var v101 int64
	_ = v101
	var v109 int64
	_ = v109
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _consts[23])))
	if v11 == int32(1) {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[158]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+117)))
		if v18 != 0 {
			v19 = int32(0)
		} else {
			v19 = v16
		}
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
		if v20 != 0 {
			v27 = v17
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+119)))
			if v28 == int32(112) {
				v70 = l1
				v71 = l2
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
				if v32 != 0 {
					v34 = l1
					v35 = l2
					v38 = v32
					for {
						v42 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
						v43 = *(*int64)(unsafe.Add(mBase, uint32(v38)+8))
						v45 = v35 - (v42 + v43)
						v46 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
						v48 = v34 - v46 + v42
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
						if v49 != 0 {
							v34 = v48
							v35 = v45
							v38 = v49
							continue
						} else {
							break
						}
						break
					}
					v51 = v48
					v52 = v45
				} else {
					v51 = l1
					v52 = l2
				}
				v59 = *(*int64)(unsafe.Add(mBase, uint32(v31)+96))
				v60 = v52 - v59
				v61 = int64(0)
				if v61 < v60 {
					v64 = v60
				} else {
					v64 = v61
				}
				v65 = int64(0)
				if v65 < v51 {
					v68 = v51
				} else {
					v68 = v65
				}
				v70 = v68
				v71 = v64
			}
			v81 = m.G0
			v82 = int32(16)
			v83 = v81 - v82
			m.G0 = v83
			F___gettimeofday(m, v83)
			mBase = m.M
			v86 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
			v87 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
			m.G0 = v83 + v82
			v95 = v87 + v86*int64(1000000) - int64(946684800000000)
			if v95 <= l4 {
				v112 = int32(0)
			} else {
				v98 = int32(2147483647)
				v101 = v95 - l4
				if base.B2i32(int64(0) < l4)^base.B2i32(v101 < v95) != 0 {
					v112 = v98
				} else {
					if int64(2147483646000) < v101 {
						v112 = v98
					} else {
						v109 = base.I64_div_s(v101+int64(999), int64(1000))
						v112 = base.I32_wrap_i64(v109)
					}
				}
			}
			v114 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
			v116 = F_pgstat_get_entry_ref_locked(m, int32(2), v19, v114, int32(0))
			mBase = m.M
			v117 = m.ExcPending
			if v117 != 0 {
				return
			} else {
				v118 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
				*(*int64)(unsafe.Add(mBase, uint32(v118)+104)) = v71
				*(*int64)(unsafe.Add(mBase, uint32(v118)+96)) = v70
				if l3 != 0 {
					*(*int64)(unsafe.Add(mBase, uint32(v118)+112)) = int64(0)
				} else {
				}
				v127 = *(*int32)(unsafe.Add(mBase, _consts[407]))
				v129 = base.B2i32(v127 == int32(4))
				if v127 == int32(4) {
					v130 = int32(192)
				} else {
					v130 = int32(176)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v118+v130))) = v95
				if v127 == int32(4) {
					v135 = int32(200)
				} else {
					v135 = int32(184)
				}
				v136 = v118 + v135
				v137 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
				*(*int64)(unsafe.Add(mBase, uint32(v136))) = v137 + int64(1)
				if v127 == int32(4) {
					v143 = int32(232)
				} else {
					v143 = int32(224)
				}
				v144 = v118 + v143
				v145 = *(*int64)(unsafe.Add(mBase, uint32(v144)))
				*(*int64)(unsafe.Add(mBase, uint32(v144))) = v145 + base.I64_extend_i32_s(v112)
				F_pgstat_unlock_entry(m, v116)
				mBase = m.M
				v149 = m.ExcPending
				if v149 != 0 {
					return
				} else {
					F_pgstat_flush_io(m, int32(0))
					mBase = m.M
					v152 = m.ExcPending
					if v152 != 0 {
						return
					} else {
						v155 = F_pgstat_flush_backend(m, int32(0), int32(1))
						mBase = m.M
						v156 = m.ExcPending
						if v156 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
			if v21 != int32(1) {
				v70 = l1
				v71 = l2
				v81 = m.G0
				v82 = int32(16)
				v83 = v81 - v82
				m.G0 = v83
				F___gettimeofday(m, v83)
				mBase = m.M
				v86 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
				v87 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
				m.G0 = v83 + v82
				v95 = v87 + v86*int64(1000000) - int64(946684800000000)
				if v95 <= l4 {
					v112 = int32(0)
				} else {
					v98 = int32(2147483647)
					v101 = v95 - l4
					if base.B2i32(int64(0) < l4)^base.B2i32(v101 < v95) != 0 {
						v112 = v98
					} else {
						if int64(2147483646000) < v101 {
							v112 = v98
						} else {
							v109 = base.I64_div_s(v101+int64(999), int64(1000))
							v112 = base.I32_wrap_i64(v109)
						}
					}
				}
				v114 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
				v116 = F_pgstat_get_entry_ref_locked(m, int32(2), v19, v114, int32(0))
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return
				} else {
					v118 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
					*(*int64)(unsafe.Add(mBase, uint32(v118)+104)) = v71
					*(*int64)(unsafe.Add(mBase, uint32(v118)+96)) = v70
					if l3 != 0 {
						*(*int64)(unsafe.Add(mBase, uint32(v118)+112)) = int64(0)
					} else {
					}
					v127 = *(*int32)(unsafe.Add(mBase, _consts[407]))
					v129 = base.B2i32(v127 == int32(4))
					if v127 == int32(4) {
						v130 = int32(192)
					} else {
						v130 = int32(176)
					}
					*(*int64)(unsafe.Add(mBase, uint32(v118+v130))) = v95
					if v127 == int32(4) {
						v135 = int32(200)
					} else {
						v135 = int32(184)
					}
					v136 = v118 + v135
					v137 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
					*(*int64)(unsafe.Add(mBase, uint32(v136))) = v137 + int64(1)
					if v127 == int32(4) {
						v143 = int32(232)
					} else {
						v143 = int32(224)
					}
					v144 = v118 + v143
					v145 = *(*int64)(unsafe.Add(mBase, uint32(v144)))
					*(*int64)(unsafe.Add(mBase, uint32(v144))) = v145 + base.I64_extend_i32_s(v112)
					F_pgstat_unlock_entry(m, v116)
					mBase = m.M
					v149 = m.ExcPending
					if v149 != 0 {
						return
					} else {
						F_pgstat_flush_io(m, int32(0))
						mBase = m.M
						v152 = m.ExcPending
						if v152 != 0 {
							return
						} else {
							v155 = F_pgstat_flush_backend(m, int32(0), int32(1))
							mBase = m.M
							v156 = m.ExcPending
							if v156 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				F_pgstat_assoc_relation(m, l0)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v27 = v26
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+119)))
					if v28 == int32(112) {
						v70 = l1
						v71 = l2
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
						if v32 != 0 {
							v34 = l1
							v35 = l2
							v38 = v32
							for {
								v42 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
								v43 = *(*int64)(unsafe.Add(mBase, uint32(v38)+8))
								v45 = v35 - (v42 + v43)
								v46 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
								v48 = v34 - v46 + v42
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
								if v49 != 0 {
									v34 = v48
									v35 = v45
									v38 = v49
									continue
								} else {
									break
								}
								break
							}
							v51 = v48
							v52 = v45
						} else {
							v51 = l1
							v52 = l2
						}
						v59 = *(*int64)(unsafe.Add(mBase, uint32(v31)+96))
						v60 = v52 - v59
						v61 = int64(0)
						if v61 < v60 {
							v64 = v60
						} else {
							v64 = v61
						}
						v65 = int64(0)
						if v65 < v51 {
							v68 = v51
						} else {
							v68 = v65
						}
						v70 = v68
						v71 = v64
					}
					v81 = m.G0
					v82 = int32(16)
					v83 = v81 - v82
					m.G0 = v83
					F___gettimeofday(m, v83)
					mBase = m.M
					v86 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
					v87 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
					m.G0 = v83 + v82
					v95 = v87 + v86*int64(1000000) - int64(946684800000000)
					if v95 <= l4 {
						v112 = int32(0)
					} else {
						v98 = int32(2147483647)
						v101 = v95 - l4
						if base.B2i32(int64(0) < l4)^base.B2i32(v101 < v95) != 0 {
							v112 = v98
						} else {
							if int64(2147483646000) < v101 {
								v112 = v98
							} else {
								v109 = base.I64_div_s(v101+int64(999), int64(1000))
								v112 = base.I32_wrap_i64(v109)
							}
						}
					}
					v114 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
					v116 = F_pgstat_get_entry_ref_locked(m, int32(2), v19, v114, int32(0))
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return
					} else {
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
						*(*int64)(unsafe.Add(mBase, uint32(v118)+104)) = v71
						*(*int64)(unsafe.Add(mBase, uint32(v118)+96)) = v70
						if l3 != 0 {
							*(*int64)(unsafe.Add(mBase, uint32(v118)+112)) = int64(0)
						} else {
						}
						v127 = *(*int32)(unsafe.Add(mBase, _consts[407]))
						v129 = base.B2i32(v127 == int32(4))
						if v127 == int32(4) {
							v130 = int32(192)
						} else {
							v130 = int32(176)
						}
						*(*int64)(unsafe.Add(mBase, uint32(v118+v130))) = v95
						if v127 == int32(4) {
							v135 = int32(200)
						} else {
							v135 = int32(184)
						}
						v136 = v118 + v135
						v137 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
						*(*int64)(unsafe.Add(mBase, uint32(v136))) = v137 + int64(1)
						if v127 == int32(4) {
							v143 = int32(232)
						} else {
							v143 = int32(224)
						}
						v144 = v118 + v143
						v145 = *(*int64)(unsafe.Add(mBase, uint32(v144)))
						*(*int64)(unsafe.Add(mBase, uint32(v144))) = v145 + base.I64_extend_i32_s(v112)
						F_pgstat_unlock_entry(m, v116)
						mBase = m.M
						v149 = m.ExcPending
						if v149 != 0 {
							return
						} else {
							F_pgstat_flush_io(m, int32(0))
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return
							} else {
								v155 = F_pgstat_flush_backend(m, int32(0), int32(1))
								mBase = m.M
								v156 = m.ExcPending
								if v156 != 0 {
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
	} else {
		return
	}
}
func F_pgstat_report_query_id(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	v5 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v5 == int32(0) {
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
		if v9 != int32(1) {
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v5)+392))
			if base.B2i32(l1 == int32(0))&base.B2i32(v14 != int64(0)) != 0 {
			} else {
				v18 = int32(4419940)
				v20 = *(*int32)(unsafe.Add(mBase, _consts[11]))
				v21 = int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[11])) = v20 + v21
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v24 + v21
				*(*int64)(unsafe.Add(mBase, uint32(v5)+392)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v24 + int32(2)
				v35 = *(*int32)(unsafe.Add(mBase, _consts[11]))
				*(*int32)(unsafe.Add(mBase, _consts[11])) = v35 - v21
			}
		}
	}
	return
}
func F_pgstat_report_stat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
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
	var v53 int64
	_ = v53
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v83 int64
	_ = v83
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v128 int32
	_ = v128
	var v132 int64
	_ = v132
	var v139 int64
	_ = v139
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v170 int64
	_ = v170
	var v172 int64
	_ = v172
	var v176 int64
	_ = v176
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v362 int64
	_ = v362
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	v2 = int32(0)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _consts[914])))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[914])) = uint8(v15)
	v18 = int32(1)
	goto L3
L2:
	;
	v18 = l0
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[915]))
	if v20 != int32(4047312) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v376
L5:
	;
	v24 = v20
	goto L7
L6:
	;
	v24 = int32(0)
	goto L7
L7:
	;
	if v24 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, _consts[725])))
	if v29 != int32(1) {
		v376 = int32(0)
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	if v18 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L10
L12:
	;
	v93 = m.G0
	v95 = v93 - int32(16)
	m.G0 = v95
	v98 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v98 != 0 {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	v36 = m.G0
	v37 = int32(16)
	v38 = v36 - v37
	m.G0 = v38
	F___gettimeofday(m, v38)
	mBase = m.M
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
	v42 = int64(*(*int32)(unsafe.Add(mBase, uint32(v38)+8)))
	m.G0 = v38 + v37
	goto L16
L14:
	;
	goto L15
L15:
	;
	v53 = *(*int64)(unsafe.Add(mBase, _consts[170]))
	if v53 == int64(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v90 = v2
	v91 = v42 + v41*int64(1000000) - int64(946684800000000)
	goto L12
L17:
	;
	v61 = *(*int64)(unsafe.Add(mBase, _consts[916]))
	if int64(0) < v61 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v57 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[170])) = v57
	v59 = v57
	goto L20
L19:
	;
	v59 = v53
	goto L20
L20:
	;
	goto L17
L21:
	;
	goto L24
L22:
	;
	goto L23
L23:
	;
	v70 = int32(1)
	v72 = *(*int64)(unsafe.Add(mBase, _consts[917]))
	if v72 <= int64(0) {
		v90 = v70
		v91 = v59
		goto L12
	} else {
		goto L26
	}
L24:
	;
	if base.I64_extend_i32_s(int32(60000))*int64(1000) <= v59-v61 {
		v90 = v2
		v91 = v59
		goto L12
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	goto L27
L27:
	;
	if base.I64_extend_i32_s(int32(1000))*int64(1000) <= v59-v72 {
		v90 = v70
		v91 = v59
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v83 = *(*int64)(unsafe.Add(mBase, _consts[916]))
	if v83 != int64(0) {
		v376 = int32(10000)
		goto L4
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, _consts[916])) = v59
	return int32(10000)
L30:
	;
	v102 = F_pgstat_prep_pending_entry(m, int32(1), v98, int64(0), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	m.G0 = v95 + int32(16)
	v198 = *(*int32)(unsafe.Add(mBase, _consts[915]))
	if v198 == int32(4047312) {
		v283 = v2
		goto L43
	} else {
		goto L44
	}
L33:
	;
	return int32(0)
L34:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v106)))
	v109 = int64(*(*int32)(unsafe.Add(mBase, _consts[918])))
	*(*int64)(unsafe.Add(mBase, uint32(v106))) = v107 + v109
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v106)+8))
	v114 = int64(*(*int32)(unsafe.Add(mBase, _consts[919])))
	*(*int64)(unsafe.Add(mBase, uint32(v106)+8)) = v112 + v114
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v106)+168))
	v119 = *(*int64)(unsafe.Add(mBase, _consts[736]))
	*(*int64)(unsafe.Add(mBase, uint32(v106)+168)) = v117 + v119
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v106)+176))
	v124 = *(*int64)(unsafe.Add(mBase, _consts[744]))
	*(*int64)(unsafe.Add(mBase, uint32(v106)+176)) = v122 + v124
	v128 = *(*int32)(unsafe.Add(mBase, _consts[407]))
	if v128 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v132 = *(*int64)(unsafe.Add(mBase, _consts[920]))
	v139 = v91 - v132
	if v139 <= int64(0) {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	goto L37
L37:
	;
	v176 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[736])) = v176
	*(*int64)(unsafe.Add(mBase, _consts[744])) = v176
	*(*int64)(unsafe.Add(mBase, _consts[921])) = v176
	*(*int64)(unsafe.Add(mBase, _consts[922])) = v176
	v188 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[919])) = v188
	*(*int32)(unsafe.Add(mBase, _consts[918])) = v188
	goto L32
L38:
	;
	*(*int64)(unsafe.Add(mBase, _consts[920])) = v91
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v106)+192))
	v158 = int64(*(*int32)(unsafe.Add(mBase, uint32(v95)+8)))
	v159 = int64(*(*int32)(unsafe.Add(mBase, uint32(v95)+12)))
	*(*int64)(unsafe.Add(mBase, uint32(v106)+192)) = v157 + (v158 + v159*int64(1000000))
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v106)+200))
	v167 = *(*int64)(unsafe.Add(mBase, _consts[921]))
	*(*int64)(unsafe.Add(mBase, uint32(v106)+200)) = v165 + v167
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v106)+208))
	v172 = *(*int64)(unsafe.Add(mBase, _consts[922]))
	*(*int64)(unsafe.Add(mBase, uint32(v106)+208)) = v170 + v172
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95+int32(12)))) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v95+int32(8)))) = v152
	goto L38
L40:
	;
	v151 = int32(0)
	v152 = int32(0)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v143 = int64(1000000)
	v144 = base.I64_div_u_s(v139, v143)
	v151 = base.I32_wrap_i64(v144)
	v152 = base.I32_wrap_i64(v139 - v144*v143)
	goto L39
L43:
	;
	v292 = int32(1)
	v294 = int32(*(*uint8)(unsafe.Add(mBase, _consts[725])))
	if v294 == v292 {
		goto L67
	} else {
		goto L68
	}
L44:
	;
	if v198 == int32(0) {
		v283 = v2
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v204 = v198
	v211 = v2
	goto L46
L46:
	;
	v215 = v204 - int32(16)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	if base.Ui32(v217-int32(1)) <= base.Ui32(int32(11)) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v283 = v276
	goto L43
L48:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+32))
	v236 = m.T0[v235].(func(*base.Module, int32, int32) int32)(m, v215, v90)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L33
	} else {
		goto L52
	}
L49:
	;
	v234 = v217*int32(72) + int32(1601888)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _consts[923]))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v227+v217<<(uint(int32(2))%32)-int32(96))))
	v234 = v233
	goto L48
L52:
	;
	v239 = v204 + int32(4)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	if v236 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v243 = v204 - int32(4)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	if base.Ui32(v246-int32(1)) <= base.Ui32(int32(11)) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v276 = int32(1)
	goto L55
L55:
	;
	if v240 == int32(4047312) {
		v283 = v276
		goto L43
	} else {
		goto L65
	}
L56:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+36))
	if v264 != 0 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v263 = v246*int32(72) + int32(1601888)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _consts[923]))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v256+v246<<(uint(int32(2))%32)-int32(96))))
	v263 = v262
	goto L56
L60:
	;
	m.T0[v264].(func(*base.Module, int32))(m, v215)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L33
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_pfree(m, v244)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L33
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = int32(0)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	*(*int32)(unsafe.Add(mBase, uint32(v271)+4)) = v272
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	*(*int32)(unsafe.Add(mBase, uint32(v272))) = v274
	v276 = v211
	goto L55
L65:
	;
	if v240 != 0 {
		v204 = v240
		v211 = v276
		goto L46
	} else {
		goto L66
	}
L66:
	;
	goto L47
L67:
	;
	v298 = v292
	v299 = v283
	goto L70
L68:
	;
	v346 = v283
	goto L69
L69:
	;
	*(*int64)(unsafe.Add(mBase, _consts[917])) = v91
	if v346&int32(1) != 0 {
		goto L82
	} else {
		goto L83
	}
L70:
	;
	if base.Ui32(v298) <= base.Ui32(int32(12)) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v346 = v339
	goto L69
L72:
	;
	v341 = v298 + int32(1)
	if v341 != int32(33) {
		v298 = v341
		v299 = v339
		goto L70
	} else {
		goto L81
	}
L73:
	;
	v329 = v298*int32(72) + int32(1601888)
	goto L75
L74:
	;
	if base.Ui32(int32(8)) < base.Ui32(v298-int32(24)) {
		v339 = v299
		goto L72
	} else {
		goto L76
	}
L75:
	;
	if v329 == int32(0) {
		v339 = v299
		goto L72
	} else {
		goto L78
	}
L76:
	;
	v319 = *(*int32)(unsafe.Add(mBase, _consts[923]))
	if v319 == int32(0) {
		v339 = v299
		goto L72
	} else {
		goto L77
	}
L77:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v319+v298<<(uint(int32(2))%32)-int32(96))))
	v329 = v327
	goto L75
L78:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v329)+56))
	if v332 == int32(0) {
		v339 = v299
		goto L72
	} else {
		goto L79
	}
L79:
	;
	v335 = m.T0[v332].(func(*base.Module, int32) int32)(m, v90)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L33
	} else {
		goto L80
	}
L80:
	;
	v339 = v335 | v299
	goto L72
L81:
	;
	goto L71
L82:
	;
	v362 = *(*int64)(unsafe.Add(mBase, _consts[916]))
	if v362 != int64(0) {
		v376 = int32(10000)
		goto L4
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	*(*int64)(unsafe.Add(mBase, _consts[916])) = int64(0)
	v373 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[725])) = uint8(v373)
	v376 = int32(0)
	goto L4
L85:
	;
	*(*int64)(unsafe.Add(mBase, _consts[916])) = v91
	return int32(10000)
}
func F_pgstat_reset(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
		v30 = l0*int32(72) + int32(1601888)
	} else {
		if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
			v30 = int32(0)
		} else {
			v18 = int32(0)
			v20 = *(*int32)(unsafe.Add(mBase, _consts[923]))
			if v20 == v18 {
				v30 = v18
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v20+l0<<(uint(int32(2))%32)-int32(96))))
				v30 = v28
			}
		}
	}
	v34 = m.G0
	v35 = int32(16)
	v36 = v34 - v35
	m.G0 = v36
	F___gettimeofday(m, v36)
	mBase = m.M
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
	v40 = int64(*(*int32)(unsafe.Add(mBase, uint32(v36)+8)))
	m.G0 = v36 + v35
	v48 = v40 + v39*int64(1000000) - int64(946684800000000)
	F_pgstat_reset_entry(m, l0, l1, l2, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		return
	} else {
		v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
		if v51&int32(2) == int32(0) {
			v58 = *(*int32)(unsafe.Add(mBase, _consts[158]))
			v61 = F_pgstat_get_entry_ref_locked(m, int32(1), v58, int64(0), int32(0))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
				*(*int64)(unsafe.Add(mBase, uint32(v63)+280)) = v48
				F_pgstat_unlock_entry(m, v61)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			return
		}
	}
}
func F_pgstat_unlock_entry(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_LWLockRelease(m, v2+int32(4))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_pgstat_wal_flush_cb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v87 int32
	_ = v87
	var v91 int64
	_ = v91
	var v95 int64
	_ = v95
	var v99 int64
	_ = v99
	var v103 int64
	_ = v103
	var v105 int32
	_ = v105
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[936]))
	v12 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v12
	v21 = *(*int64)(unsafe.Add(mBase, _consts[18]))
	v23 = *(*int64)(unsafe.Add(mBase, _consts[1042]))
	if v21 == v23 {
		v105 = int32(0)
		m.G0 = v8 + int32(32)
		return v105
	} else {
		v26 = v11 + int32(53272)
		v27 = int32(4408480)
		v28 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
		v30 = *(*int64)(unsafe.Add(mBase, _consts[14]))
		v31 = *(*int64)(unsafe.Add(mBase, _consts[1043]))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v28 + (v30 - v31)
		v35 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
		v37 = *(*int64)(unsafe.Add(mBase, _consts[18]))
		v38 = *(*int64)(unsafe.Add(mBase, _consts[1042]))
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v35 + (v37 - v38)
		v42 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
		v44 = *(*int64)(unsafe.Add(mBase, _consts[16]))
		v45 = *(*int64)(unsafe.Add(mBase, _consts[1044]))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v42 + (v44 - v45)
		v49 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
		v51 = *(*int64)(unsafe.Add(mBase, _consts[12]))
		v52 = *(*int64)(unsafe.Add(mBase, _consts[1045]))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v49 + (v51 - v52)
		if l0 == int32(0) {
			v59 = F_LWLockAcquire(m, v26, int32(0))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				v70 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1046])))
				v71 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1046]))) = v70 + v71
				v74 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1047])))
				v75 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1047]))) = v74 + v75
				v78 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1048])))
				v79 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1048]))) = v78 + v79
				v82 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1049])))
				v83 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1049]))) = v82 + v83
				F_LWLockRelease(m, v26)
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					v91 = *(*int64)(unsafe.Add(mBase, _consts[12]))
					*(*int64)(unsafe.Add(mBase, _consts[1045])) = v91
					v95 = *(*int64)(unsafe.Add(mBase, _consts[14]))
					*(*int64)(unsafe.Add(mBase, _consts[1043])) = v95
					v99 = *(*int64)(unsafe.Add(mBase, _consts[16]))
					*(*int64)(unsafe.Add(mBase, _consts[1044])) = v99
					v103 = *(*int64)(unsafe.Add(mBase, _consts[18]))
					*(*int64)(unsafe.Add(mBase, _consts[1042])) = v103
					v105 = int32(0)
					m.G0 = v8 + int32(32)
					return v105
				}
			}
		} else {
			v65 = F_LWLockConditionalAcquire(m, v26, int32(0))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				if v65 == int32(0) {
					v105 = int32(1)
					m.G0 = v8 + int32(32)
					return v105
				} else {
					v70 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1046])))
					v71 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
					*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1046]))) = v70 + v71
					v74 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1047])))
					v75 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1047]))) = v74 + v75
					v78 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1048])))
					v79 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1048]))) = v78 + v79
					v82 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1049])))
					v83 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[1049]))) = v82 + v83
					F_LWLockRelease(m, v26)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						v91 = *(*int64)(unsafe.Add(mBase, _consts[12]))
						*(*int64)(unsafe.Add(mBase, _consts[1045])) = v91
						v95 = *(*int64)(unsafe.Add(mBase, _consts[14]))
						*(*int64)(unsafe.Add(mBase, _consts[1043])) = v95
						v99 = *(*int64)(unsafe.Add(mBase, _consts[16]))
						*(*int64)(unsafe.Add(mBase, _consts[1044])) = v99
						v103 = *(*int64)(unsafe.Add(mBase, _consts[18]))
						*(*int64)(unsafe.Add(mBase, _consts[1042])) = v103
						v105 = int32(0)
						m.G0 = v8 + int32(32)
						return v105
					}
				}
			}
		}
	}
}
