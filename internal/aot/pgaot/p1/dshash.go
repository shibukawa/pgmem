package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dshash_find_or_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
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
	var v99 int32
	_ = v99
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
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
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
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
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v19 = m.T0[v18].(func(*base.Module, int32, int32, int32) int32)(m, l1, v16, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v27 = int32(base.Ui32(v19)>>(uint(int32(25))%32)) * int32(20)
	v28 = v23 + v27
	v32 = v23
	goto L5
L3:
	;
	return v337 + int32(8)
L4:
	;
	v332 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v332)
	v337 = v88
	goto L3
L5:
	;
	v48 = F_LWLockAcquire(m, v32+v27+int32(8), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v306 = F_dsa_allocate_extended(m, v301, v302+int32(8), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L52
	}
L7:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+2572))
	if v50 == v52 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v64+int32(base.Ui32(v19)>>(uint(int32(32)-v63)%32))<<(uint(int32(2))%32))))
	if v71 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v63 = v50
	v64 = v54
	goto L8
L10:
	;
	goto L11
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+2576))
	v57 = F_dsa_get_address(m, v55, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v57
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+2572))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v61
	v63 = v61
	v64 = v57
	goto L8
L13:
	;
	v75 = v71
	goto L16
L14:
	;
	goto L15
L15:
	;
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v115)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v118 = int32(1)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v122 = v118 << (uint(v119-int32(7)) % 32)
	if base.Ui32(int32(base.Ui32(v122)>>(uint(v118)%32))+int32(base.Ui32(v122)>>(uint(int32(2))%32))) < base.Ui32(v117) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v88 = F_dsa_get_address(m, v87, v75)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v95 = m.T0[v94].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v88+int32(8), v92, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v95 == int32(0) {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v99 != 0 {
		v75 = v99
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_LWLockRelease(m, v129+v27+int32(8))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L6
L25:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v140 = F_LWLockAcquire(m, v136+int32(8), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v142 = int32(1)
	v144 = v135 + v142
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+2572))
	if base.Ui32(v144) <= base.Ui32(v146) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_LWLockRelease(m, v145+int32(8))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v156 = v142
	goto L31
L30:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v32 = v152
	goto L5
L31:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v175 = F_LWLockAcquire(m, v168+v156*int32(20)+int32(8), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v187 = F_dsa_allocate_extended(m, v183, int32(4)<<(uint(v144)%32), int32(5))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	v178 = v156 + int32(1)
	if v178 != int32(128) {
		v156 = v178
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v190 = F_dsa_get_address(m, v189, v187)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+2572))
	v203 = int32(0)
	goto L37
L37:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210+v203<<(uint(int32(2))%32))))
	if v214 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+2576))
	*(*int32)(unsafe.Add(mBase, uint32(v262)+2576)) = v187
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v265)+2572)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v190
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_dsa_free(m, v268, v263)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L47
	}
L39:
	;
	v218 = v214
	goto L42
L40:
	;
	goto L41
L41:
	;
	v258 = v203 + int32(1)
	if int32(base.Ui32(v258)>>(uint(v193)%32)) == int32(0) {
		v203 = v258
		goto L37
	} else {
		goto L46
	}
L42:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v231 = F_dsa_get_address(m, v230, v218)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	goto L41
L44:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v238 = v190 + int32(base.Ui32(v234)>>(uint(int32(31)-v135)%32))<<(uint(int32(2))%32)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	*(*int32)(unsafe.Add(mBase, uint32(v231))) = v239
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v218
	if v233 != 0 {
		v218 = v233
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	goto L38
L47:
	;
	v275 = int32(0)
	goto L48
L48:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_LWLockRelease(m, v287+v275*int32(20)+int32(8))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v32 = v299
	goto L5
L50:
	;
	v296 = v275 + int32(1)
	if v296 != int32(128) {
		v275 = v296
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v309 = F_dsa_get_address(m, v308, v306)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	m.T0[v315].(func(*base.Module, int32, int32, int32, int32))(m, v309+int32(8), l1, v313, v314)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v323 = v300 + int32(base.Ui32(v19)>>(uint(int32(32)-v119)%32))<<(uint(int32(2))%32)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	*(*int32)(unsafe.Add(mBase, uint32(v309))) = v324
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v309)+4)) = v19
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v328 + int32(1)
	v337 = v309
	goto L3
}
func F_dshash_memcmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	if base.Ui32(int32(4)) <= base.Ui32(l2) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v66
L2:
	;
	v66 = int32(0)
	goto L1
L3:
	;
	v40 = v35
	v41 = v36
	v42 = v37
	goto L13
L4:
	;
	if (l0|l1)&int32(3) != 0 {
		v35 = l0
		v36 = l1
		v37 = l2
		goto L3
	} else {
		goto L7
	}
L5:
	;
	v28 = l0
	v29 = l1
	v30 = l2
	goto L6
L6:
	;
	if v30 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v12 = l0
	v13 = l1
	v14 = l2
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v17 != v18 {
		v35 = v12
		v36 = v13
		v37 = v14
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v28 = v23
	v29 = v21
	v30 = v25
	goto L6
L10:
	;
	v20 = int32(4)
	v21 = v13 + v20
	v23 = v12 + v20
	v25 = v14 - v20
	if base.Ui32(int32(3)) < base.Ui32(v25) {
		v12 = v23
		v13 = v21
		v14 = v25
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v35 = v28
	v36 = v29
	v37 = v30
	goto L3
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v45 == v46 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v66 = v45 - v46
	goto L1
L15:
	;
	v48 = int32(1)
	v53 = v42 - v48
	if v53 != 0 {
		v40 = v40 + v48
		v41 = v41 + v48
		v42 = v53
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	goto L2
}
func F_dshash_release_lock(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1-int32(4))))
	F_LWLockRelease(m, v3+int32(base.Ui32(v6)>>(uint(int32(25))%32))*int32(20)+int32(8))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		return
	}
}
func F_dshash_seq_next(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v33 int32
	_ = v33
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
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
	var v94 int32
	_ = v94
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v5 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v112 = F_dsa_get_address(m, v111, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L26
	}
L2:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v53 != 0 {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v17 = F_LWLockAcquire(m, v11+int32(8), v14^int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v52 = l0 + int32(16)
	goto L2
L6:
	;
	return int32(0)
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+2572))
	if v23 != v25 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+2576))
	v29 = F_dsa_get_address(m, v27, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	v38 = v22
	v39 = v23
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1) << (uint(v39) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v52 = v42 + v43<<(uint(int32(2))%32)
	goto L2
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v29
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+2572))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+32))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+2572))
	v38 = v35
	v39 = v37
	goto L10
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v109 = v54
	v110 = v53
	goto L1
L13:
	;
	goto L14
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v57 = v55
	goto L15
L15:
	;
	v61 = v57 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v63 <= v61 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v109 = v99
	v110 = v104
	goto L1
L17:
	;
	return int32(0)
L18:
	;
	goto L19
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+40))
	v71 = v61 >> (uint(v68-int32(7)) % 32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v71 != v72 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v83 = F_LWLockAcquire(m, v74+v71*int32(20)+int32(8), v80^int32(1))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L23
	}
L21:
	;
	v98 = v61
	v99 = v67
	goto L22
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+36))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v98<<(uint(int32(2))%32))))
	if v104 == int32(0) {
		v57 = v98
		goto L15
	} else {
		goto L25
	}
L23:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+32))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_LWLockRelease(m, v86+v87*int32(20)+int32(8))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v71
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v98 = v97
	v99 = v96
	goto L22
L25:
	;
	goto L16
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v112
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v115
	return v112 + int32(8)
}
