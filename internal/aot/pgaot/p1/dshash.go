package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dshash_find_or_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
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
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
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
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v20 = m.T0[v19].(func(*base.Module, int32, int32, int32) int32)(m, l1, v17, v18)
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
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v28 = int32(base.Ui32(v20)>>(uint(int32(25))%32)) * int32(20)
	v31 = v24 + v28 + int32(24)
	v35 = v24
	goto L4
L3:
	;
	return v350 + int32(8)
L4:
	;
	v52 = F_LWLockAcquire(m, v35+v28+int32(8), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v321 = F_dsa_allocate_extended(m, v316, v317+int32(8), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L53
	}
L6:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+2572))
	if v54 == v56 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v68+int32(base.Ui32(v20)>>(uint(int32(32)-v67)%32))<<(uint(int32(2))%32))))
	if v75 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v67 = v54
	v68 = v58
	goto L7
L9:
	;
	goto L10
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)+2576))
	v61 = F_dsa_get_address(m, v59, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v61
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+2572))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v65
	v67 = v65
	v68 = v61
	goto L7
L12:
	;
	v125 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v125)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v128 = int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v132 = v128 << (uint(v129-int32(7)) % 32)
	if base.Ui32(int32(base.Ui32(v132)>>(uint(v128)%32))+int32(base.Ui32(v132)>>(uint(int32(2))%32))) < base.Ui32(v127) {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v81 = v75
	goto L14
L14:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v95 = F_dsa_get_address(m, v94, v81)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	if v95 == int32(0) {
		goto L12
	} else {
		goto L22
	}
L16:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v102 = m.T0[v101].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v95+int32(8), v99, v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v102 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v104 != 0 {
		v81 = v104
		goto L14
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L15
L21:
	;
	goto L12
L22:
	;
	v107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v107)
	v350 = v95
	goto L3
L23:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_LWLockRelease(m, v139+v28+int32(8))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L5
L26:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v150 = F_LWLockAcquire(m, v146+int32(8), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v152 = int32(1)
	v154 = v145 + v152
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+2572))
	if base.Ui32(v154) <= base.Ui32(v156) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_LWLockRelease(m, v155+int32(8))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v166 = v152
	goto L32
L31:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v35 = v162
	goto L4
L32:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v186 = F_LWLockAcquire(m, v179+v166*int32(20)+int32(8), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v198 = F_dsa_allocate_extended(m, v194, int32(4)<<(uint(v154)%32), int32(5))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	v189 = v166 + int32(1)
	if v189 != int32(128) {
		v166 = v189
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v201 = F_dsa_get_address(m, v200, v198)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+2572))
	v214 = int32(0)
	goto L38
L38:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v222+v214<<(uint(int32(2))%32))))
	if v226 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+2576))
	*(*int32)(unsafe.Add(mBase, uint32(v276)+2576)) = v198
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v279)+2572)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v201
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_dsa_free(m, v282, v277)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L48
	}
L40:
	;
	v230 = v226
	goto L43
L41:
	;
	goto L42
L42:
	;
	v272 = v214 + int32(1)
	if int32(base.Ui32(v272)>>(uint(v204)%32)) == int32(0) {
		v214 = v272
		goto L38
	} else {
		goto L47
	}
L43:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v244 = F_dsa_get_address(m, v243, v230)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v251 = v201 + int32(base.Ui32(v247)>>(uint(int32(31)-v145)%32))<<(uint(int32(2))%32)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = v252
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v230
	if v246 != 0 {
		v230 = v246
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L39
L48:
	;
	v289 = int32(0)
	goto L49
L49:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_LWLockRelease(m, v302+v289*int32(20)+int32(8))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v35 = v314
	goto L4
L51:
	;
	v311 = v289 + int32(1)
	if v311 != int32(128) {
		v289 = v311
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v324 = F_dsa_get_address(m, v323, v321)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	m.T0[v330].(func(*base.Module, int32, int32, int32, int32))(m, v324+int32(8), l1, v328, v329)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v338 = v315 + int32(base.Ui32(v20)>>(uint(int32(32)-v129)%32))<<(uint(int32(2))%32)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	*(*int32)(unsafe.Add(mBase, uint32(v324))) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v338))) = v321
	*(*int32)(unsafe.Add(mBase, uint32(v324)+4)) = v20
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v343 + int32(1)
	v350 = v324
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
