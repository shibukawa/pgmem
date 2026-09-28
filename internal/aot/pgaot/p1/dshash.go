package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dshash_find_or_insert_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
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
	var v102 int32
	_ = v102
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
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
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v408 int32
	_ = v408
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
	v26 = int32(base.Ui32(v20) >> (uint(int32(25)) % 32))
	v28 = v26 * int32(20)
	v29 = v24 + v28
	v33 = v24
	goto L6
L3:
	;
	return v408
L4:
	;
	v408 = v390 + int32(8)
	goto L3
L5:
	;
	v385 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v385)
	v390 = v91
	goto L4
L6:
	;
	v50 = F_LWLockAcquire(m, v33+v28+int32(8), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v347 = F_dsa_allocate_extended(m, v342, v343+int32(8), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L60
	}
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+2572))
	if v52 == v54 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v66+int32(base.Ui32(v20)>>(uint(int32(32)-v65)%32))<<(uint(int32(2))%32))))
	if v73 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v65 = v52
	v66 = v56
	goto L9
L11:
	;
	goto L12
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+2576))
	v59 = F_dsa_get_address(m, v57, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v59
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+2572))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v63
	v65 = v63
	v66 = v59
	goto L9
L14:
	;
	v77 = v73
	goto L17
L15:
	;
	goto L16
L16:
	;
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v119)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v122 = int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v126 = v122 << (uint(v123-int32(7)) % 32)
	if base.Ui32(int32(base.Ui32(v126)>>(uint(v122)%32))+int32(base.Ui32(v126)>>(uint(int32(2))%32))) < base.Ui32(v121) {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v91 = F_dsa_get_address(m, v90, v77)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v98 = m.T0[v97].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v91+int32(8), v95, v96)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v98 == int32(0) {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v102 != 0 {
		v77 = v102
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_LWLockRelease(m, v133+v28+int32(8))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L7
L26:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v144 = F_LWLockAcquire(m, v140+int32(8), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v146 = int32(1)
	v148 = v139 + v146
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+2572))
	if base.Ui32(v148) <= base.Ui32(v150) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_LWLockRelease(m, v149+int32(8))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v160 = v146
	goto L32
L31:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v33 = v156
	goto L6
L32:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v180 = F_LWLockAcquire(m, v173+v160*int32(20)+int32(8), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v190 = F_dsa_allocate_extended(m, v186, int32(4)<<(uint(v148)%32), int32(5))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	v183 = v160 + int32(1)
	if v183 != int32(128) {
		v160 = v183
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	if v190 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v199 = int32(0)
	goto L40
L38:
	;
	goto L39
L39:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v227 = F_dsa_get_address(m, v226, v190)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L44
	}
L40:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_LWLockRelease(m, v211+v199*int32(20)+int32(8))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	v408 = int32(0)
	goto L3
L42:
	;
	v221 = v199 + int32(1)
	if v221 != int32(128) {
		v199 = v221
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+2572))
	v240 = int32(0)
	goto L45
L45:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v248+v240<<(uint(int32(2))%32))))
	if v252 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+2576))
	*(*int32)(unsafe.Add(mBase, uint32(v302)+2576)) = v190
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v305)+2572)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v227
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_dsa_free(m, v308, v303)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L55
	}
L47:
	;
	v256 = v252
	goto L50
L48:
	;
	goto L49
L49:
	;
	v298 = v240 + int32(1)
	if int32(base.Ui32(v298)>>(uint(v230)%32)) == int32(0) {
		v240 = v298
		goto L45
	} else {
		goto L54
	}
L50:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v270 = F_dsa_get_address(m, v269, v256)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L52
	}
L51:
	;
	goto L49
L52:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	v277 = v227 + int32(base.Ui32(v273)>>(uint(int32(31)-v139)%32))<<(uint(int32(2))%32)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	*(*int32)(unsafe.Add(mBase, uint32(v270))) = v278
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = v256
	if v272 != 0 {
		v256 = v272
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	goto L46
L55:
	;
	v315 = int32(0)
	goto L56
L56:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_LWLockRelease(m, v328+v315*int32(20)+int32(8))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v33 = v340
	goto L6
L58:
	;
	v337 = v315 + int32(1)
	if v337 != int32(128) {
		v315 = v337
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	if v347 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_LWLockRelease(m, v351+v26*int32(20)+int32(8))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v362 = F_dsa_get_address(m, v361, v347)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	return int32(0)
L65:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	m.T0[v368].(func(*base.Module, int32, int32, int32, int32))(m, v362+int32(8), l1, v366, v367)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v376 = v341 + int32(base.Ui32(v20)>>(uint(int32(32)-v123)%32))<<(uint(int32(2))%32)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	*(*int32)(unsafe.Add(mBase, uint32(v362))) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v376))) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v362)+4)) = v20
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v381 + int32(1)
	v390 = v362
	goto L4
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
