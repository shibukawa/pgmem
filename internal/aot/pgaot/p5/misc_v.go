package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ValuesNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	v12 = int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v16 == v12 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	m.T0[v30].(func(*base.Module, int32))(m, v14)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v13 + v24
	goto L1
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v13 < v19 {
		v24 = v12
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v13 < int32(0) {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L1
L7:
	;
	v24 = int32(-1)
	goto L2
L8:
	;
	return int32(0)
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v35 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return v14
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v38 <= v35 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v41 = v35 << (uint(int32(2)) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41+v42)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v45+v41)))
	F_ReScanExprContext(m, v28)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v50 = int32(0)
	v51 = int32(4486928)
	v52 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v54
	if v44 == v50 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v52
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	v130 = v128 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)) = uint16(v130)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+6)) = uint16(v133)
	goto L32
L15:
	;
	v59 = F_ExecInitExprList(m, v47, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L8
	} else {
		goto L18
	}
L16:
	;
	v63 = v44
	goto L17
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v64 <= int32(0) {
		goto L14
	} else {
		goto L20
	}
L18:
	;
	if v59 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v63 = v59
	goto L17
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v69 = v50
	goto L21
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v82 = v69 << (uint(int32(2)) % 32)
	v83 = v68 + v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v82)))
	v87 = v69 + v67
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v89 = m.T0[v88].(func(*base.Module, int32, int32, int32) int32)(m, v86, v28, v87)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L8
	} else {
		goto L23
	}
L22:
	;
	goto L14
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v89
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v92 != 0 {
		v109 = v89
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v109
	v112 = v69 + int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v112 < v113 {
		v69 = v112
		goto L21
	} else {
		goto L31
	}
L25:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80+v69<<(uint(int32(4))%32))+24)))
	if v96 != int32(65535) {
		v109 = v89
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v99 != int32(1) {
		v108 = v89
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v109 = v108
	goto L24
L28:
	;
	goto L27
L29:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	if v102 != int32(3) {
		v108 = v89
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v89)+2))
	v108 = v105 + int32(18)
	goto L28
L31:
	;
	goto L22
L32:
	;
	goto L10
}
func F_VirtualXactLockTableInsert(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, _consts[296]))
	v8 = F_LWLockAcquire(m, v4+int32(584), int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[296]))
		v12 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+608)) = uint8(v12)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+612)) = v14
		F_LWLockRelease(m, v11+int32(584))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			return
		}
	}
}
func F_vac_cleanup_one_index(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = F_index_vacuum_cleanup(m, l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			m.G0 = v8 - int32(-64)
			return v10
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = F_errstart(m, v16, int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 == int32(0) {
					m.G0 = v8 - int32(-64)
					return v10
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
					v24 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v25
					*(*float64)(unsafe.Add(mBase, uint32(v8)+40)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v23 + int32(4)
					F_errmsg(m, int32(169827), v6+int32(-32))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
						v37 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v38
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v37
						*(*float64)(unsafe.Add(mBase, uint32(v8))) = v36
						F_errdetail(m, int32(621823), v8)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(493871), int32(2687), int32(28152))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 - int32(-64)
								return v10
							}
						}
					}
				}
			}
		}
	}
}
func F_varchar(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v18 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v50 = v17 - int32(4)
	if v50 < int32(0) {
		v104 = v13
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v21 = int32(4)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v23&int32(254) == int32(2) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v36 = int32(1)
	if v18&v36 != 0 {
		v48 = int32(base.Ui32(v18)>>(uint(v36)%32)) - v36
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v32 = v21
	goto L9
L8:
	;
	v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
	goto L9
L9:
	;
	if v23 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v35 = v21
	goto L12
L11:
	;
	v35 = v32
	goto L12
L12:
	;
	v48 = v35
	goto L3
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	m.G0 = v10 + int32(16)
	return v104
L15:
	;
	if v48 <= v50 {
		v104 = v13
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v55 = int32(1)
	if v18&v55 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v59 = v55
	goto L19
L18:
	;
	v59 = int32(4)
	goto L19
L19:
	;
	v60 = v13 + v59
	v61 = F_pg_mbcharcliplen(m, v60, v48, v50)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v54 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v101 = F_cstring_to_text_with_len(m, v60, v61)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L34
	}
L22:
	;
	if v48 <= v61 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v65 = v61
	goto L24
L24:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v65))))
	if v72 == int32(32) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L30
	}
L26:
	;
	v76 = v65 + int32(1)
	if v48 != v76 {
		v65 = v76
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	goto L21
L30:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v50
	F_errmsg(m, int32(660242), v10)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(492418), int32(640), int32(228714))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v104 = v101
	goto L14
}
func F_vfprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F___vfprintf_internal(m, l0, l1, l2, int32(6955), int32(6956))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_vfscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int64
	_ = v68
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int64
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int64
	_ = v187
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v298 int32
	_ = v298
	var v300 int64
	_ = v300
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int64
	_ = v313
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v565 int32
	_ = v565
	var v566 int64
	_ = v566
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int64
	_ = v576
	var v583 int32
	_ = v583
	var v601 int64
	_ = v601
	var v604 int64
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int64
	_ = v633
	var v636 int32
	_ = v636
	var v644 int32
	_ = v644
	var v658 int32
	_ = v658
	var v659 int64
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v669 int64
	_ = v669
	var v670 int64
	_ = v670
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v753 int32
	_ = v753
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v864 int32
	_ = v864
	var v865 int64
	_ = v865
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v981 int64
	_ = v981
	var v984 int32
	_ = v984
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1014 int64
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1102 int32
	_ = v1102
	var v1118 int64
	_ = v1118
	var v1149 int64
	_ = v1149
	var v1151 int64
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1167 int64
	_ = v1167
	var v1176 int64
	_ = v1176
	var v1177 int64
	_ = v1177
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1191 int32
	_ = v1191
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1234 int32
	_ = v1234
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1258 int64
	_ = v1258
	var v1264 int64
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1285 int64
	_ = v1285
	var v1290 int64
	_ = v1290
	var v1293 int64
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int64
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1312 int64
	_ = v1312
	var v1318 int64
	_ = v1318
	var v1319 int64
	_ = v1319
	var v1321 int64
	_ = v1321
	var v1324 int64
	_ = v1324
	var v1325 int64
	_ = v1325
	var v1327 int64
	_ = v1327
	var v1328 int64
	_ = v1328
	var v1332 int64
	_ = v1332
	var v1339 int64
	_ = v1339
	var v1350 int64
	_ = v1350
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1407 int32
	_ = v1407
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1431 int64
	_ = v1431
	var v1438 int64
	_ = v1438
	var v1439 int64
	_ = v1439
	var v1445 int32
	_ = v1445
	var v1461 int64
	_ = v1461
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1480 int64
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1506 int64
	_ = v1506
	var v1513 int32
	_ = v1513
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1552 int32
	_ = v1552
	var v1571 int32
	_ = v1571
	var v1579 int64
	_ = v1579
	var v1584 int64
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1593 int64
	_ = v1593
	var v1617 int64
	_ = v1617
	var v1624 int64
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1652 int32
	_ = v1652
	var v1654 int32
	_ = v1654
	var v1657 int64
	_ = v1657
	var v1661 int64
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1669 int32
	_ = v1669
	var v1673 int64
	_ = v1673
	var v1674 int64
	_ = v1674
	var v1678 int32
	_ = v1678
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1728 int64
	_ = v1728
	var v1733 int64
	_ = v1733
	var v1736 int32
	_ = v1736
	var v1737 int64
	_ = v1737
	var v1739 int64
	_ = v1739
	var v1740 int64
	_ = v1740
	var v1742 int64
	_ = v1742
	var v1746 int64
	_ = v1746
	var v1750 int64
	_ = v1750
	var v1751 int64
	_ = v1751
	var v1755 int32
	_ = v1755
	var v1768 int32
	_ = v1768
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1801 int32
	_ = v1801
	var v1803 int32
	_ = v1803
	var v1806 int64
	_ = v1806
	var v1810 int64
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1820 int64
	_ = v1820
	var v1825 int64
	_ = v1825
	var v1835 int64
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1839 int64
	_ = v1839
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1867 int64
	_ = v1867
	var v1873 int64
	_ = v1873
	var v1878 int64
	_ = v1878
	var v1881 int64
	_ = v1881
	var v1884 int64
	_ = v1884
	var v1886 int64
	_ = v1886
	var v1887 int64
	_ = v1887
	var v1895 int64
	_ = v1895
	var v1905 int64
	_ = v1905
	var v1909 int32
	_ = v1909
	var v1910 int64
	_ = v1910
	var v1912 int64
	_ = v1912
	var v1919 int64
	_ = v1919
	var v1938 int32
	_ = v1938
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1968 int32
	_ = v1968
	var v1971 int32
	_ = v1971
	var v1976 int32
	_ = v1976
	var v1979 int32
	_ = v1979
	var v1981 int32
	_ = v1981
	var v2001 int32
	_ = v2001
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2036 int32
	_ = v2036
	var v2041 int32
	_ = v2041
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2048 int32
	_ = v2048
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2073 int32
	_ = v2073
	var v2080 int32
	_ = v2080
	var v2089 int32
	_ = v2089
	var v2091 int32
	_ = v2091
	var v2104 int32
	_ = v2104
	var v2108 int32
	_ = v2108
	var v2111 int32
	_ = v2111
	var v2115 int32
	_ = v2115
	var v2126 int32
	_ = v2126
	var v2133 int32
	_ = v2133
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2204 int32
	_ = v2204
	var v2209 int32
	_ = v2209
	var v2257 int32
	_ = v2257
	var v2266 int32
	_ = v2266
	var v2270 int32
	_ = v2270
	var v2274 int32
	_ = v2274
	var v2277 int32
	_ = v2277
	var v2280 int32
	_ = v2280
	var v2283 int32
	_ = v2283
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2300 int32
	_ = v2300
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2321 int32
	_ = v2321
	var v2324 int32
	_ = v2324
	var v2326 int32
	_ = v2326
	var v2346 int32
	_ = v2346
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2381 int32
	_ = v2381
	var v2388 int32
	_ = v2388
	var v2390 int32
	_ = v2390
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2399 int32
	_ = v2399
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2434 int32
	_ = v2434
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2482 int32
	_ = v2482
	var v2484 int32
	_ = v2484
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2506 int32
	_ = v2506
	var v2507 int64
	_ = v2507
	var v2511 int32
	_ = v2511
	var v2513 int32
	_ = v2513
	var v2514 int64
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2518 int64
	_ = v2518
	var v2537 int32
	_ = v2537
	var v2544 int32
	_ = v2544
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2564 int32
	_ = v2564
	var v2565 int32
	_ = v2565
	var v2568 int64
	_ = v2568
	var v2574 int32
	_ = v2574
	var v2576 int32
	_ = v2576
	var v2579 int32
	_ = v2579
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2583 int32
	_ = v2583
	var v2590 int32
	_ = v2590
	var v2593 int32
	_ = v2593
	var v2596 int64
	_ = v2596
	var v2601 int32
	_ = v2601
	var v2605 int32
	_ = v2605
	var v2615 int32
	_ = v2615
	var v2616 int32
	_ = v2616
	var v2620 int32
	_ = v2620
	var v2633 int32
	_ = v2633
	var v2642 int32
	_ = v2642
	var v2643 int32
	_ = v2643
	var v2647 int32
	_ = v2647
	var v2650 int32
	_ = v2650
	var v2670 int32
	_ = v2670
	var v2696 int32
	_ = v2696
	var v2705 int32
	_ = v2705
	v4 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(304)
	m.G0 = v28
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v31 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v2696 + int32(304)
	return v2705
L2:
	;
	v2696 = v2670
	v2705 = int32(-1)
	goto L1
L3:
	;
	v34 = F___toread(m, l0)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v41 == int32(0) {
		v2696 = v28
		v2705 = v4
		goto L1
	} else {
		goto L9
	}
L6:
	;
	return int32(0)
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v38 == int32(0) {
		v2670 = v28
		goto L2
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v46 = l0
	v47 = l1
	v48 = l2
	v51 = v41
	v53 = v28
	v54 = v4
	v55 = v4
	v62 = v4
	v65 = v28 + int32(16)
	v68 = int64(0)
	goto L12
L10:
	;
	if v2647 == int32(0) {
		v2696 = v53
		v2705 = v2650
		goto L1
	} else {
		goto L590
	}
L11:
	;
	if v62 != 0 {
		goto L587
	} else {
		goto L588
	}
L12:
	;
	v72 = v51 & int32(255)
	goto L16
L13:
	;
	v2605 = int32(0)
	v2615 = v2605
	v2616 = v2605
	v2620 = int32(1)
	goto L11
L14:
	;
	goto L13
L15:
	;
	v2601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2579)+1)))
	if v2601 != 0 {
		v46 = v2574
		v47 = v2579 + int32(1)
		v48 = v2576
		v51 = v2601
		v53 = v2581
		v54 = v2582
		v55 = v2583
		v62 = v2590
		v65 = v2593
		v68 = v2596
		goto L12
	} else {
		goto L586
	}
L16:
	;
	if base.B2i32(v72 == int32(32))|base.B2i32(base.Ui32(v72-int32(9)) < base.Ui32(int32(5))) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v81 = v47
	goto L20
L18:
	;
	goto L19
L19:
	;
	if v72 == int32(37) {
		goto L43
	} else {
		goto L44
	}
L20:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	goto L22
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+112)) = int64(0)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v46)+120)) = base.I64_extend_i32_s(v119 - v120)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	goto L25
L22:
	;
	if base.B2i32(v107 == int32(32))|base.B2i32(base.Ui32(v107-int32(9)) < base.Ui32(int32(5))) != 0 {
		v81 = v81 + int32(1)
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L28
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+104)) = v124
	goto L24
L28:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v159 != v160 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v46)+112))
	if int64(0) <= v177 {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	goto L35
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v159 + int32(1)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v168 = v165
	goto L30
L32:
	;
	goto L33
L33:
	;
	v166 = F___shgetc(m, v46)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v168 = v166
	goto L30
L35:
	;
	if base.B2i32(v168 == int32(32))|base.B2i32(base.Ui32(v168-int32(9)) < base.Ui32(int32(5))) != 0 {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	goto L29
L37:
	;
	v181 = v176 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v181
	v183 = v181
	goto L39
L38:
	;
	v183 = v176
	goto L39
L39:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v46)+120))
	v2574 = v46
	v2576 = v48
	v2579 = v81
	v2581 = v53
	v2582 = v54
	v2583 = v55
	v2590 = v62
	v2593 = v65
	v2596 = base.I64_extend_i32_s(v183-v184) + (v187 + v68)
	goto L15
L40:
	;
	v354 = int32(0)
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	if base.Ui32((v356-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L83
	} else {
		goto L84
	}
L41:
	;
	v320 = v192 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v320) {
		goto L77
	} else {
		goto L78
	}
L42:
	;
	v350 = v48
	v352 = v47 + int32(2)
	v353 = int32(0)
	goto L40
L43:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	if v192 == int32(42) {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+112)) = int64(0)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v46)+120)) = base.I64_extend_i32_s(v202 - v203)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	goto L49
L46:
	;
	if v192 != int32(37) {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v217 == int32(37) {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+104)) = v207
	goto L48
L52:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	if v298 != v278 {
		goto L69
	} else {
		goto L70
	}
L53:
	;
	goto L56
L54:
	;
	goto L55
L55:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v264 != v265 {
		goto L65
	} else {
		goto L66
	}
L56:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v245 != v246 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v274 = v47 + int32(1)
	v278 = v254
	goto L52
L58:
	;
	goto L63
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v245 + int32(1)
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	v254 = v251
	goto L58
L60:
	;
	goto L61
L61:
	;
	v252 = F___shgetc(m, v46)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v254 = v252
	goto L58
L63:
	;
	if base.B2i32(v254 == int32(32))|base.B2i32(base.Ui32(v254-int32(9)) < base.Ui32(int32(5))) != 0 {
		goto L56
	} else {
		goto L64
	}
L64:
	;
	goto L57
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v264 + int32(1)
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	v274 = v47
	v278 = v270
	goto L52
L66:
	;
	goto L67
L67:
	;
	v271 = F___shgetc(m, v46)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	v274 = v47
	v278 = v271
	goto L52
L69:
	;
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v46)+112))
	if int64(0) <= v300 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v46)+120))
	v2574 = v46
	v2576 = v48
	v2579 = v274
	v2581 = v53
	v2582 = v54
	v2583 = v55
	v2590 = v62
	v2593 = v65
	v2596 = base.I64_extend_i32_s(v309-v310) + (v313 + v68)
	goto L15
L72:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v303 - int32(1)
	goto L74
L73:
	;
	goto L74
L74:
	;
	if int32(0) <= v278 {
		v2696 = v53
		v2705 = v62
		goto L1
	} else {
		goto L75
	}
L75:
	;
	if v62 != 0 {
		v2696 = v53
		v2705 = v62
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v2670 = v53
	goto L2
L77:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v350 = v48 + int32(4)
	v352 = v47 + int32(1)
	v353 = v346
	goto L40
L78:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+2)))
	if v323 != int32(36) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v328 = m.G0
	v330 = v328 - int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v330)+12)) = v48
	if base.Ui32(int32(1)) < base.Ui32(v320) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v339 = v48 + v320<<(uint(int32(2))%32) - int32(4)
	goto L82
L81:
	;
	v339 = v48
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+8)) = v339 + int32(4)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v350 = v48
	v352 = v47 + int32(3)
	v353 = v343
	goto L40
L83:
	;
	v364 = v356
	v367 = v354
	v368 = v352
	goto L86
L84:
	;
	v405 = v356
	v408 = v354
	v409 = v352
	goto L85
L85:
	;
	if v405&int32(255) != int32(109) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v388 = int32(10)
	v390 = int32(255)
	v393 = int32(48)
	v394 = v367*v388 + v364&v390 - v393
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+1)))
	v397 = v368 + int32(1)
	if base.Ui32((v395-v393)&v390) < base.Ui32(v388) {
		v364 = v395
		v367 = v394
		v368 = v397
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v405 = v395
	v408 = v394
	v409 = v397
	goto L85
L88:
	;
	goto L87
L89:
	;
	v440 = v405
	v441 = v54
	v442 = v55
	v443 = v354
	v444 = v409
	goto L91
L90:
	;
	v433 = int32(0)
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+1)))
	v440 = v436
	v441 = v433
	v442 = v433
	v443 = base.B2i32(v353 != v433)
	v444 = v409 + int32(1)
	goto L91
L91:
	;
	v446 = v444 + int32(1)
	switch v440&int32(255) - int32(65) {
	case 0, 2, 4, 5, 6, 18, 23, 26, 32, 34, 35, 36, 37, 38, 40, 45, 46, 47, 50, 52, 55:
		goto L93
	default:
		v2615 = v441
		v2616 = v442
		v2620 = v443
		goto L11
	case 11:
		goto L94
	case 39:
		goto L97
	case 41:
		v474 = int32(3)
		v475 = v446
		goto L92
	case 43:
		goto L96
	case 51, 57:
		goto L95
	}
L92:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
	v481 = base.B2i32(v477&int32(47) == int32(3))
	if v477&int32(47) == int32(3) {
		goto L110
	} else {
		goto L111
	}
L93:
	;
	v474 = int32(0)
	v475 = v444
	goto L92
L94:
	;
	v474 = int32(2)
	v475 = v446
	goto L92
L95:
	;
	v474 = int32(1)
	v475 = v446
	goto L92
L96:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+1)))
	v465 = base.B2i32(v463 == int32(108))
	if v463 == int32(108) {
		goto L104
	} else {
		goto L105
	}
L97:
	;
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+1)))
	v456 = base.B2i32(v454 == int32(104))
	if v454 == int32(104) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v457 = v444 + int32(2)
	goto L100
L99:
	;
	v457 = v446
	goto L100
L100:
	;
	if v454 == int32(104) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v460 = int32(-2)
	goto L103
L102:
	;
	v460 = int32(-1)
	goto L103
L103:
	;
	v474 = v460
	v475 = v457
	goto L92
L104:
	;
	v466 = v444 + int32(2)
	goto L106
L105:
	;
	v466 = v446
	goto L106
L106:
	;
	if v463 == int32(108) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v469 = int32(3)
	goto L109
L108:
	;
	v469 = int32(1)
	goto L109
L109:
	;
	v474 = v469
	v475 = v466
	goto L92
L110:
	;
	v482 = int32(1)
	goto L112
L111:
	;
	v482 = v474
	goto L112
L112:
	;
	if v477&int32(47) == int32(3) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v604 = base.I64_extend_i32_s(v583)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+112)) = v604
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v46)+120)) = base.I64_extend_i32_s(v608 - v609)
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if v604 == int64(0) {
		v621 = v613
		goto L150
	} else {
		goto L151
	}
L114:
	;
	v485 = v477 | int32(32)
	goto L116
L115:
	;
	v485 = v477
	goto L116
L116:
	;
	if v485 == int32(91) {
		v583 = v408
		v601 = v68
		goto L113
	} else {
		goto L117
	}
L117:
	;
	if v485 != int32(110) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+112)) = int64(0)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v46)+120)) = base.I64_extend_i32_s(v508 - v509)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	goto L134
L119:
	;
	if v485 != int32(99) {
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	if v353 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L122:
	;
	v492 = int32(1)
	if v408 <= v492 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v495 = v492
	goto L125
L124:
	;
	v495 = v408
	goto L125
L125:
	;
	v583 = v495
	v601 = v68
	goto L113
L126:
	;
	v2574 = v46
	v2576 = v350
	v2579 = v475
	v2581 = v53
	v2582 = v441
	v2583 = v442
	v2590 = v62
	v2593 = v65
	v2596 = v68
	goto L15
L127:
	;
	goto L126
L128:
	;
	switch v482 + int32(2) {
	case 0:
		goto L132
	case 1:
		goto L131
	case 2, 3:
		goto L130
	default:
		goto L127
	case 5:
		goto L129
	}
L129:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v353))) = v68
	goto L127
L130:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v353))) = uint32(v68)
	goto L126
L131:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v353))) = uint16(v68)
	goto L126
L132:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v353))) = uint8(v68)
	goto L126
L133:
	;
	goto L137
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+104)) = v513
	goto L133
L137:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v548 != v549 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v46)+112))
	if int64(0) <= v566 {
		goto L146
	} else {
		goto L147
	}
L139:
	;
	goto L144
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v548 + int32(1)
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548))))
	v557 = v554
	goto L139
L141:
	;
	goto L142
L142:
	;
	v555 = F___shgetc(m, v46)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L6
	} else {
		goto L143
	}
L143:
	;
	v557 = v555
	goto L139
L144:
	;
	if base.B2i32(v557 == int32(32))|base.B2i32(base.Ui32(v557-int32(9)) < base.Ui32(int32(5))) != 0 {
		goto L137
	} else {
		goto L145
	}
L145:
	;
	goto L138
L146:
	;
	v570 = v565 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v570
	v572 = v570
	goto L148
L147:
	;
	v572 = v565
	goto L148
L148:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	v576 = *(*int64)(unsafe.Add(mBase, uint32(v46)+120))
	v583 = v408
	v601 = base.I64_extend_i32_s(v572-v573) + (v576 + v68)
	goto L113
L149:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v623 != v624 {
		goto L154
	} else {
		goto L155
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+104)) = v621
	goto L149
L151:
	;
	if base.I64_extend_i32_s(v613-v609) <= v604 {
		v621 = v613
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v621 = v609 + base.I32_wrap_i64(v604)
	goto L150
L153:
	;
	v633 = *(*int64)(unsafe.Add(mBase, uint32(v46)+112))
	if int64(0) <= v633 {
		goto L159
	} else {
		goto L160
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v623 + int32(1)
	goto L153
L155:
	;
	goto L156
L156:
	;
	v629 = F___shgetc(m, v46)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L6
	} else {
		goto L157
	}
L157:
	;
	if v629 < int32(0) {
		v2615 = v441
		v2616 = v442
		v2620 = v443
		goto L11
	} else {
		goto L158
	}
L158:
	;
	goto L153
L159:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v636 - int32(1)
	goto L161
L160:
	;
	goto L161
L161:
	;
	switch v485 - int32(88) {
	case 0, 24, 32:
		v864 = int32(16)
		goto L167
	case 1, 2, 4, 5, 6, 7, 8, 10, 16, 18, 19, 20, 21, 22, 25, 26, 28, 30, 31:
		v2544 = v475
		v2547 = v441
		v2548 = v442
		goto L162
	case 3, 11, 27:
		goto L171
	case 9, 13, 14, 15:
		goto L172
	case 12, 29:
		goto L169
	case 17:
		goto L168
	case 23:
		goto L170
	default:
		goto L173
	}
L162:
	;
	v2564 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v2565 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	v2568 = *(*int64)(unsafe.Add(mBase, uint32(v46)+120))
	v2574 = v46
	v2576 = v350
	v2579 = v2544
	v2581 = v53
	v2582 = v2547
	v2583 = v2548
	v2590 = v62 + base.B2i32(v353 != int32(0))
	v2593 = v65
	v2596 = base.I64_extend_i32_s(v2564-v2565) + (v2568 + v601)
	goto L15
L163:
	;
	v1962 = base.B2i32(v485 != int32(99))
	if v485 != int32(99) {
		goto L457
	} else {
		goto L458
	}
L164:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v353))) = v670
	*(*int64)(unsafe.Add(mBase, uint32(v353)+8)) = v669
	v2544 = v475
	v2547 = v441
	v2548 = v442
	goto L162
L165:
	;
	v1801 = m.G0
	v1803 = v1801 - int32(32)
	m.G0 = v1803
	v1806 = v669 & int64(281474976710655)
	v1810 = int64(base.Ui64(v669)>>(uint(int64(48))%64)) & int64(32767)
	v1811 = base.I32_wrap_i64(v1810)
	if base.Ui32(v1811-int32(15361)) <= base.Ui32(int32(2045)) {
		goto L423
	} else {
		goto L424
	}
L166:
	;
	v1652 = m.G0
	v1654 = v1652 - int32(32)
	m.G0 = v1654
	v1657 = v669 & int64(281474976710655)
	v1661 = int64(base.Ui64(v669)>>(uint(int64(48))%64)) & int64(32767)
	v1662 = base.I32_wrap_i64(v1661)
	if base.Ui32(v1662-int32(16257)) <= base.Ui32(int32(253)) {
		goto L378
	} else {
		goto L379
	}
L167:
	;
	v865 = int64(0)
	v866 = int32(0)
	v870 = m.G0
	v872 = v870 - int32(16)
	m.G0 = v872
	if base.B2i32(v864 != int32(1))&base.B2i32(base.Ui32(v864) <= base.Ui32(int32(36))) == v866 {
		goto L216
	} else {
		goto L217
	}
L168:
	;
	v864 = int32(0)
	goto L167
L169:
	;
	v864 = int32(10)
	goto L167
L170:
	;
	v864 = int32(8)
	goto L167
L171:
	;
	if v485|int32(16) == int32(115) {
		goto L179
	} else {
		goto L180
	}
L172:
	;
	F___floatscan(m, v53+int32(8), v46, v482, int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L6
	} else {
		goto L176
	}
L173:
	;
	v644 = v485 - int32(65)
	if base.Ui32(int32(6)) < base.Ui32(v644) {
		v2544 = v475
		v2547 = v441
		v2548 = v442
		goto L162
	} else {
		goto L174
	}
L174:
	;
	if int32(1)<<(uint(v644)%32)&int32(113) == int32(0) {
		v2544 = v475
		v2547 = v441
		v2548 = v442
		goto L162
	} else {
		goto L175
	}
L175:
	;
	goto L172
L176:
	;
	v659 = *(*int64)(unsafe.Add(mBase, uint32(v46)+120))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v659 == int64(0)-base.I64_extend_i32_s(v661-v662) {
		v2642 = v441
		v2643 = v442
		v2647 = v443
		v2650 = v62
		goto L10
	} else {
		goto L177
	}
L177:
	;
	if v353 == int32(0) {
		v2544 = v475
		v2547 = v441
		v2548 = v442
		goto L162
	} else {
		goto L178
	}
L178:
	;
	v669 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	v670 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
	switch v482 {
	case 0:
		goto L166
	case 1:
		goto L165
	case 2:
		goto L164
	default:
		v2544 = v475
		v2547 = v441
		v2548 = v442
		goto L162
	}
L179:
	;
	v680 = F__emscripten_memset_bulkmem(m, v53+int32(32), base.I32_extend8_s(int32(-1)), int32(257))
	mBase = m.M
	goto L182
L180:
	;
	goto L181
L181:
	;
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+1)))
	v695 = base.B2i32(v693 == int32(94))
	v698 = F__emscripten_memset_bulkmem(m, v53+int32(32), base.I32_extend8_s(v695), int32(257))
	mBase = m.M
	goto L184
L182:
	;
	v681 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+32)) = uint8(v681)
	if v485 != int32(115) {
		v1938 = v475
		goto L163
	} else {
		goto L183
	}
L183:
	;
	v685 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+65)) = uint8(v685)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+46)) = uint8(v685)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+42)) = v685
	v1938 = v475
	goto L163
L184:
	;
	v699 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+32)) = uint8(v699)
	if v693 == int32(94) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v705 = v475 + int32(2)
	goto L187
L186:
	;
	v705 = v475 + int32(1)
	goto L187
L187:
	;
	if v693 == int32(94) {
		goto L191
	} else {
		goto L192
	}
L188:
	;
	v733 = v727
	goto L198
L189:
	;
	v726 = v723
	v727 = v705 + int32(1)
	goto L188
L190:
	;
	v721 = base.B2i32(v693 != int32(94))
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+126)) = uint8(v721)
	v723 = v721
	goto L189
L191:
	;
	v708 = int32(2)
	goto L193
L192:
	;
	v708 = int32(1)
	goto L193
L193:
	;
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475+v708))))
	if v710 != int32(45) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	if v710 == int32(93) {
		goto L190
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v718 = base.B2i32(v693 != int32(94))
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+78)) = uint8(v718)
	v723 = v718
	goto L189
L197:
	;
	v726 = base.B2i32(v693 != int32(94))
	v727 = v705
	goto L188
L198:
	;
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v733))))
	if v753 != int32(45) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v833+(v53+int32(32)))+1)) = uint8(v726)
	v733 = v835 + int32(1)
	goto L198
L201:
	;
	if v753 == int32(0) {
		v2615 = v441
		v2616 = v442
		v2620 = v443
		goto L11
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v760 = int32(45)
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v733)+1)))
	if v761 == int32(0) {
		v833 = v760
		v835 = v733
		goto L200
	} else {
		goto L206
	}
L204:
	;
	if v753 == int32(93) {
		v1938 = v733
		goto L163
	} else {
		goto L205
	}
L205:
	;
	v833 = v753
	v835 = v733
	goto L200
L206:
	;
	if v761 == int32(93) {
		v833 = v760
		v835 = v733
		goto L200
	} else {
		goto L207
	}
L207:
	;
	v766 = int32(1)
	v767 = v733 + v766
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v733-v766))))
	if base.Ui32(v761) <= base.Ui32(v770) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v833 = v808
	v835 = v767
	goto L200
L209:
	;
	v808 = v761
	goto L208
L210:
	;
	goto L211
L211:
	;
	v773 = v770
	goto L212
L212:
	;
	v798 = v773 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v798+(v53+int32(32))))) = uint8(v726)
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767))))
	if base.Ui32(v798) < base.Ui32(v803) {
		v773 = v798
		goto L212
	} else {
		goto L214
	}
L213:
	;
	v808 = v803
	goto L208
L214:
	;
	goto L213
L215:
	;
	m.G0 = v872 + int32(16)
	v1624 = *(*int64)(unsafe.Add(mBase, uint32(v46)+120))
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v1624 == int64(0)-base.I64_extend_i32_s(v1626-v1627) {
		v2642 = v441
		v2643 = v442
		v2647 = v443
		v2650 = v62
		goto L10
	} else {
		goto L365
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(28)
	v1617 = v865
	goto L215
L217:
	;
	goto L218
L218:
	;
	goto L219
L219:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v909 != v910 {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	switch v918 - int32(43) {
	case 0, 2:
		goto L229
	default:
		v942 = v918
		v943 = v866
		goto L228
	}
L221:
	;
	goto L226
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v909 + int32(1)
	v915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v909))))
	v918 = v915
	goto L221
L223:
	;
	goto L224
L224:
	;
	v916 = F___shgetc(m, v46)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L6
	} else {
		goto L225
	}
L225:
	;
	v918 = v916
	goto L221
L226:
	;
	if base.B2i32(v918 == int32(32))|base.B2i32(base.Ui32(v918-int32(9)) < base.Ui32(int32(5))) != 0 {
		goto L219
	} else {
		goto L227
	}
L227:
	;
	goto L220
L228:
	;
	if base.B2i32(v864 != int32(0))&base.B2i32(v864 != int32(16)) != 0 {
		goto L241
	} else {
		goto L242
	}
L229:
	;
	if v918 == int32(45) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v932 = int32(-1)
	goto L232
L231:
	;
	v932 = int32(0)
	goto L232
L232:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v933 != v934 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v933 + int32(1)
	v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v933))))
	v942 = v939
	v943 = v932
	goto L228
L234:
	;
	goto L235
L235:
	;
	v940 = F___shgetc(m, v46)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L6
	} else {
		goto L236
	}
L236:
	;
	v942 = v940
	v943 = v932
	goto L228
L237:
	;
	v1584 = *(*int64)(unsafe.Add(mBase, uint32(v46)+112))
	if int64(0) <= v1584 {
		goto L360
	} else {
		goto L361
	}
L238:
	;
	v1513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1489)+uint32(_consts[1480]))))
	if base.Ui32(v1487) <= base.Ui32(v1513) {
		v1571 = v943
		v1579 = v1506
		goto L237
	} else {
		goto L351
	}
L239:
	;
	if v1184&(v1184-int32(1)) != 0 {
		goto L303
	} else {
		goto L304
	}
L240:
	;
	if v1043 != int32(10) {
		v1184 = v1043
		v1185 = v1044
		goto L239
	} else {
		goto L277
	}
L241:
	;
	if v864 != 0 {
		goto L266
	} else {
		goto L267
	}
L242:
	;
	if v942 != int32(48) {
		goto L241
	} else {
		goto L243
	}
L243:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v951 != v952 {
		goto L245
	} else {
		goto L246
	}
L244:
	;
	if v960&int32(-33) == int32(88) {
		goto L249
	} else {
		goto L250
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v951 + int32(1)
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951))))
	v960 = v957
	goto L244
L246:
	;
	goto L247
L247:
	;
	v958 = F___shgetc(m, v46)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L6
	} else {
		goto L248
	}
L248:
	;
	v960 = v958
	goto L244
L249:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v965 != v966 {
		goto L253
	} else {
		goto L254
	}
L250:
	;
	goto L251
L251:
	;
	if v864 != 0 {
		v1043 = v864
		v1044 = v960
		goto L240
	} else {
		goto L265
	}
L252:
	;
	v975 = int32(16)
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v974)+uint32(_consts[1480]))))
	if base.Ui32(v978) < base.Ui32(v975) {
		v1184 = v975
		v1185 = v974
		goto L239
	} else {
		goto L257
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v965 + int32(1)
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965))))
	v974 = v971
	goto L252
L254:
	;
	goto L255
L255:
	;
	v972 = F___shgetc(m, v46)
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L6
	} else {
		goto L256
	}
L256:
	;
	v974 = v972
	goto L252
L257:
	;
	v981 = *(*int64)(unsafe.Add(mBase, uint32(v46)+112))
	if int64(0) <= v981 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v984 - int32(1)
	goto L260
L259:
	;
	goto L260
L260:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+112)) = int64(0)
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v46)+120)) = base.I64_extend_i32_s(v992 - v993)
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	goto L262
L261:
	;
	v1617 = v865
	goto L215
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+104)) = v997
	goto L261
L265:
	;
	v1184 = int32(8)
	v1185 = v960
	goto L239
L266:
	;
	v1009 = v864
	goto L268
L267:
	;
	v1009 = int32(10)
	goto L268
L268:
	;
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v942)+uint32(_consts[1480]))))
	if base.Ui32(v1012) < base.Ui32(v1009) {
		v1043 = v1009
		v1044 = v942
		goto L240
	} else {
		goto L269
	}
L269:
	;
	v1014 = *(*int64)(unsafe.Add(mBase, uint32(v46)+112))
	if int64(0) <= v1014 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v1017 - int32(1)
	goto L272
L271:
	;
	goto L272
L272:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+112)) = int64(0)
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v46)+120)) = base.I64_extend_i32_s(v1025 - v1026)
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	goto L274
L273:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(28)
	v1617 = v865
	goto L215
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+104)) = v1030
	goto L273
L277:
	;
	v1048 = v1044 - int32(48)
	if base.Ui32(v1048) <= base.Ui32(int32(9)) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1055 = int32(0)
	v1056 = v1048
	goto L281
L279:
	;
	v1102 = v1048
	v1118 = v865
	goto L280
L280:
	;
	if base.Ui32(int32(9)) < base.Ui32(v1102) {
		v1571 = v943
		v1579 = v1118
		goto L237
	} else {
		goto L289
	}
L281:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v1077 != v1078 {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	v1102 = v1093
	v1118 = base.I64_extend_i32_u(v1089)
	goto L280
L283:
	;
	v1089 = v1055*int32(10) + v1056
	v1093 = v1086 - int32(48)
	if base.B2i32(base.Ui32(v1089) < base.Ui32(int32(429496729)))&base.B2i32(base.Ui32(v1093) <= base.Ui32(int32(9))) != 0 {
		v1055 = v1089
		v1056 = v1093
		goto L281
	} else {
		goto L288
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v1077 + int32(1)
	v1083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1077))))
	v1086 = v1083
	goto L283
L285:
	;
	goto L286
L286:
	;
	v1084 = F___shgetc(m, v46)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L6
	} else {
		goto L287
	}
L287:
	;
	v1086 = v1084
	goto L283
L288:
	;
	goto L282
L289:
	;
	v1149 = v1118 * int64(10)
	v1151 = base.I64_extend_i32_u(v1102)
	goto L290
L290:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v1153 != v1154 {
		goto L294
	} else {
		goto L295
	}
L291:
	;
	v1487 = int32(10)
	v1489 = v1162
	v1506 = v1167
	goto L238
L292:
	;
	goto L291
L293:
	;
	v1164 = v1162 - int32(48)
	v1167 = v1149 + v1151
	if base.B2i32(base.Ui32(v1164) <= base.Ui32(int32(9)))&base.B2i32(base.Ui64(v1167) < base.Ui64(int64(1844674407370955162))) == int32(0) {
		goto L298
	} else {
		goto L299
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v1153 + int32(1)
	v1159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1153))))
	v1162 = v1159
	goto L293
L295:
	;
	goto L296
L296:
	;
	v1160 = F___shgetc(m, v46)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L6
	} else {
		goto L297
	}
L297:
	;
	v1162 = v1160
	goto L293
L298:
	;
	if base.Ui32(v1164) <= base.Ui32(int32(9)) {
		goto L292
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	v1176 = v1167 * int64(10)
	v1177 = base.I64_extend_i32_u(v1164)
	if base.Ui64(v1176) <= base.Ui64(v1177^int64(-1)) {
		v1149 = v1176
		v1151 = v1177
		goto L290
	} else {
		goto L302
	}
L301:
	;
	v1571 = v943
	v1579 = v1167
	goto L237
L302:
	;
	goto L292
L303:
	;
	v1191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1185)+uint32(_consts[1480]))))
	if base.Ui32(v1191) < base.Ui32(v1184) {
		goto L306
	} else {
		goto L307
	}
L304:
	;
	goto L305
L305:
	;
	v1361 = int32(*(*int8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1184*int32(23))>>(uint(int32(5))%32))&int32(7))+uint32(_consts[1481]))))
	v1364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1185)+uint32(_consts[1480]))))
	if base.Ui32(v1364) < base.Ui32(v1184) {
		goto L329
	} else {
		goto L330
	}
L306:
	;
	v1197 = v866
	v1199 = v1191
	goto L309
L307:
	;
	v1241 = v1185
	v1244 = v1191
	v1258 = v865
	goto L308
L308:
	;
	if base.Ui32(v1184) <= base.Ui32(v1244) {
		v1487 = v1184
		v1489 = v1241
		v1506 = v1258
		goto L238
	} else {
		goto L317
	}
L309:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v1218 != v1219 {
		goto L312
	} else {
		goto L313
	}
L310:
	;
	v1241 = v1227
	v1244 = v1234
	v1258 = base.I64_extend_i32_u(v1229)
	goto L308
L311:
	;
	v1229 = v1199 + v1184*v1197
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227)+uint32(_consts[1480]))))
	if base.B2i32(base.Ui32(v1229) < base.Ui32(int32(119304647)))&base.B2i32(base.Ui32(v1234) < base.Ui32(v1184)) != 0 {
		v1197 = v1229
		v1199 = v1234
		goto L309
	} else {
		goto L316
	}
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v1218 + int32(1)
	v1224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1218))))
	v1227 = v1224
	goto L311
L313:
	;
	goto L314
L314:
	;
	v1225 = F___shgetc(m, v46)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L6
	} else {
		goto L315
	}
L315:
	;
	v1227 = v1225
	goto L311
L316:
	;
	goto L310
L317:
	;
	v1264 = base.I64_extend_i32_u(v1184)
	v1268 = v1241
	v1271 = v1244
	v1285 = v1258
	goto L318
L318:
	;
	v1290 = v1285 * v1264
	v1293 = base.I64_extend_i32_u(v1271) & int64(255)
	if base.Ui64(v1293^int64(-1)) < base.Ui64(v1290) {
		v1487 = v1184
		v1489 = v1268
		v1506 = v1285
		goto L238
	} else {
		goto L320
	}
L319:
	;
	v1487 = v1184
	v1489 = v1306
	v1506 = v1307
	goto L238
L320:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v1297 != v1298 {
		goto L322
	} else {
		goto L323
	}
L321:
	;
	v1307 = v1290 + v1293
	v1310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1306)+uint32(_consts[1480]))))
	if base.Ui32(v1184) <= base.Ui32(v1310) {
		v1487 = v1184
		v1489 = v1306
		v1506 = v1307
		goto L238
	} else {
		goto L326
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v1297 + int32(1)
	v1303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1297))))
	v1306 = v1303
	goto L321
L323:
	;
	goto L324
L324:
	;
	v1304 = F___shgetc(m, v46)
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L6
	} else {
		goto L325
	}
L325:
	;
	v1306 = v1304
	goto L321
L326:
	;
	v1312 = int64(0)
	v1318 = int64(32)
	v1319 = int64(base.Ui64(v1307) >> (uint(v1318) % 64))
	v1321 = int64(base.Ui64(v1264) >> (uint(v1318) % 64))
	v1324 = int64(4294967295)
	v1325 = v1307 & v1324
	v1327 = v1264 & v1324
	v1328 = v1325 * v1327
	v1332 = int64(base.Ui64(v1328)>>(uint(v1318)%64)) + v1325*v1321
	v1339 = v1327*v1319 + v1332&v1324
	*(*int64)(unsafe.Add(mBase, uint32(v872)+8)) = v1264*v1312 + v1312*v1307 + v1319*v1321 + int64(base.Ui64(v1332)>>(uint(v1318)%64)) + int64(base.Ui64(v1339)>>(uint(v1318)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v872))) = v1328&v1324 | v1339<<(uint(v1318)%64)
	goto L327
L327:
	;
	v1350 = *(*int64)(unsafe.Add(mBase, uint32(v872)+8))
	if v1350 == int64(0) {
		v1268 = v1306
		v1271 = v1310
		v1285 = v1307
		goto L318
	} else {
		goto L328
	}
L328:
	;
	goto L319
L329:
	;
	v1370 = v1364
	v1372 = v866
	goto L332
L330:
	;
	v1414 = v1185
	v1415 = v1364
	v1431 = v865
	goto L331
L331:
	;
	if base.Ui32(v1184) <= base.Ui32(v1415) {
		v1487 = v1184
		v1489 = v1414
		v1506 = v1431
		goto L238
	} else {
		goto L340
	}
L332:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v1391 != v1392 {
		goto L335
	} else {
		goto L336
	}
L333:
	;
	v1414 = v1400
	v1415 = v1407
	v1431 = base.I64_extend_i32_u(v1402)
	goto L331
L334:
	;
	v1401 = v1372 << (uint(v1361) % 32)
	v1402 = v1370 | v1401
	v1407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1400)+uint32(_consts[1480]))))
	if base.B2i32(base.Ui32(v1401) < base.Ui32(int32(134217728)))&base.B2i32(base.Ui32(v1407) < base.Ui32(v1184)) != 0 {
		v1370 = v1407
		v1372 = v1402
		goto L332
	} else {
		goto L339
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v1391 + int32(1)
	v1397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391))))
	v1400 = v1397
	goto L334
L336:
	;
	goto L337
L337:
	;
	v1398 = F___shgetc(m, v46)
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L6
	} else {
		goto L338
	}
L338:
	;
	v1400 = v1398
	goto L334
L339:
	;
	goto L333
L340:
	;
	v1438 = base.I64_extend_i32_u(v1361)
	v1439 = int64(base.Ui64(int64(-1)) >> (uint(v1438) % 64))
	if base.Ui64(v1439) < base.Ui64(v1431) {
		v1487 = v1184
		v1489 = v1414
		v1506 = v1431
		goto L238
	} else {
		goto L341
	}
L341:
	;
	v1445 = v1415
	v1461 = v1431
	goto L342
L342:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v1469 != v1470 {
		goto L345
	} else {
		goto L346
	}
L343:
	;
	v1487 = v1184
	v1489 = v1478
	v1506 = v1480
	goto L238
L344:
	;
	v1480 = v1461<<(uint(v1438)%64) | base.I64_extend_i32_u(v1445)&int64(255)
	v1483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1478)+uint32(_consts[1480]))))
	if base.Ui32(v1184) <= base.Ui32(v1483) {
		v1487 = v1184
		v1489 = v1478
		v1506 = v1480
		goto L238
	} else {
		goto L349
	}
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v1469 + int32(1)
	v1475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1469))))
	v1478 = v1475
	goto L344
L346:
	;
	goto L347
L347:
	;
	v1476 = F___shgetc(m, v46)
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L6
	} else {
		goto L348
	}
L348:
	;
	v1478 = v1476
	goto L344
L349:
	;
	if base.Ui64(v1480) <= base.Ui64(v1439) {
		v1445 = v1483
		v1461 = v1480
		goto L342
	} else {
		goto L350
	}
L350:
	;
	goto L343
L351:
	;
	goto L352
L352:
	;
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v1540 != v1541 {
		goto L355
	} else {
		goto L356
	}
L353:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(68)
	v1571 = int32(0)
	v1579 = int64(-1)
	goto L237
L354:
	;
	v1552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1549)+uint32(_consts[1480]))))
	if base.Ui32(v1552) < base.Ui32(v1487) {
		goto L352
	} else {
		goto L359
	}
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v1540 + int32(1)
	v1546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1540))))
	v1549 = v1546
	goto L354
L356:
	;
	goto L357
L357:
	;
	v1547 = F___shgetc(m, v46)
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L6
	} else {
		goto L358
	}
L358:
	;
	v1549 = v1547
	goto L354
L359:
	;
	goto L353
L360:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v1587 - int32(1)
	goto L362
L361:
	;
	goto L362
L362:
	;
	goto L363
L363:
	;
	v1593 = base.I64_extend_i32_s(v1571)
	v1617 = v1579 ^ v1593 - v1593
	goto L215
L365:
	;
	if v485 != int32(112) {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	if v353 == int32(0) {
		goto L370
	} else {
		goto L371
	}
L367:
	;
	if v353 == int32(0) {
		goto L366
	} else {
		goto L368
	}
L368:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v353))) = uint32(v1617)
	v2544 = v475
	v2547 = v441
	v2548 = v442
	goto L162
L369:
	;
	v2544 = v475
	v2547 = v441
	v2548 = v442
	goto L162
L370:
	;
	goto L369
L371:
	;
	switch v482 + int32(2) {
	case 0:
		goto L375
	case 1:
		goto L374
	case 2, 3:
		goto L373
	default:
		goto L370
	case 5:
		goto L372
	}
L372:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v353))) = v1617
	goto L370
L373:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v353))) = uint32(v1617)
	goto L369
L374:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v353))) = uint16(v1617)
	goto L369
L375:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v353))) = uint8(v1617)
	goto L369
L376:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v353))) = base.F32_reinterpret_i32(base.I32_wrap_i64(int64(base.Ui64(v669)>>(uint(int64(32))%64)))&int32(-2147483648) | v1776<<(uint(int32(23))%32) | v1775)
	v2544 = v475
	v2547 = v441
	v2548 = v442
	goto L162
L377:
	;
	m.G0 = v1654 + int32(32)
	goto L376
L378:
	;
	v1669 = base.I32_wrap_i64(int64(base.Ui64(v1657) >> (uint(int64(25)) % 64)))
	v1673 = v669 & int64(33554431)
	v1674 = int64(16777216)
	if v1673 == v1674 {
		goto L382
	} else {
		goto L383
	}
L379:
	;
	goto L380
L380:
	;
	if v670|v1657 == int64(0) {
		goto L395
	} else {
		goto L396
	}
L381:
	;
	v1694 = base.B2i32(base.Ui32(int32(8388607)) < base.Ui32(v1691))
	if base.Ui32(int32(8388607)) < base.Ui32(v1691) {
		goto L389
	} else {
		goto L390
	}
L382:
	;
	v1678 = base.B2i32(v670 == int64(0))
	goto L384
L383:
	;
	v1678 = base.B2i32(base.Ui64(v1673) < base.Ui64(v1674))
	goto L384
L384:
	;
	if v1678 == int32(0) {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1691 = v1669 + int32(1)
	goto L381
L386:
	;
	goto L387
L387:
	;
	if v670|(v1673^int64(16777216)) != int64(0) {
		v1691 = v1669
		goto L381
	} else {
		goto L388
	}
L388:
	;
	v1691 = v1669&int32(1) + v1669
	goto L381
L389:
	;
	v1695 = int32(0)
	goto L391
L390:
	;
	v1695 = v1691
	goto L391
L391:
	;
	if base.Ui32(int32(8388607)) < base.Ui32(v1691) {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v1698 = int32(-16255)
	goto L394
L393:
	;
	v1698 = int32(-16256)
	goto L394
L394:
	;
	v1775 = v1695
	v1776 = v1698 + v1662
	goto L377
L395:
	;
	if base.Ui32(int32(16510)) < base.Ui32(v1662) {
		goto L398
	} else {
		goto L399
	}
L396:
	;
	if v1661 != int64(32767) {
		goto L395
	} else {
		goto L397
	}
L397:
	;
	v1775 = base.I32_wrap_i64(int64(base.Ui64(v1657)>>(uint(int64(25))%64))) | int32(4194304)
	v1776 = int32(255)
	goto L377
L398:
	;
	v1775 = int32(0)
	v1776 = int32(255)
	goto L377
L399:
	;
	goto L400
L400:
	;
	v1717 = base.B2i32(v1661 == int64(0))
	if v1661 == int64(0) {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v1718 = int32(16256)
	goto L403
L402:
	;
	v1718 = int32(16257)
	goto L403
L403:
	;
	v1719 = v1718 - v1662
	if int32(112) < v1719 {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1722 = int32(0)
	v1775 = v1722
	v1776 = v1722
	goto L377
L405:
	;
	goto L406
L406:
	;
	if v1661 == int64(0) {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v1728 = v1657
	goto L409
L408:
	;
	v1728 = v1657 | int64(281474976710656)
	goto L409
L409:
	;
	F___ashlti3(m, v1654+int32(16), v670, v1728, int32(128)-v1719)
	mBase = m.M
	F___lshrti3(m, v1654, v670, v1728, v1719)
	mBase = m.M
	v1733 = *(*int64)(unsafe.Add(mBase, uint32(v1654)+8))
	v1736 = base.I32_wrap_i64(int64(base.Ui64(v1733) >> (uint(int64(25)) % 64)))
	v1737 = *(*int64)(unsafe.Add(mBase, uint32(v1654)))
	v1739 = *(*int64)(unsafe.Add(mBase, uint32(v1654)+16))
	v1740 = *(*int64)(unsafe.Add(mBase, uint32(v1654)+24))
	v1742 = int64(0)
	v1746 = v1737 | base.I64_extend_i32_u(base.B2i32(v1662 != v1718)&base.B2i32(v1739|v1740 != v1742))
	v1750 = v1733 & int64(33554431)
	v1751 = int64(16777216)
	if v1750 == v1751 {
		goto L411
	} else {
		goto L412
	}
L410:
	;
	v1772 = base.B2i32(base.Ui32(int32(8388607)) < base.Ui32(v1768))
	if base.Ui32(int32(8388607)) < base.Ui32(v1768) {
		goto L418
	} else {
		goto L419
	}
L411:
	;
	v1755 = base.B2i32(v1746 == v1742)
	goto L413
L412:
	;
	v1755 = base.B2i32(base.Ui64(v1750) < base.Ui64(v1751))
	goto L413
L413:
	;
	if v1755 == int32(0) {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v1768 = v1736 + int32(1)
	goto L410
L415:
	;
	goto L416
L416:
	;
	if v1746|(v1750^int64(16777216)) != int64(0) {
		v1768 = v1736
		goto L410
	} else {
		goto L417
	}
L417:
	;
	v1768 = v1736&int32(1) + v1736
	goto L410
L418:
	;
	v1773 = v1768 ^ int32(8388608)
	goto L420
L419:
	;
	v1773 = v1768
	goto L420
L420:
	;
	v1775 = v1773
	v1776 = v1772
	goto L377
L421:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v353))) = base.F64_reinterpret_i64(v669&int64(-9223372036854775807-1) | v1919<<(uint(int64(52))%64) | v1912)
	v2544 = v475
	v2547 = v441
	v2548 = v442
	goto L162
L422:
	;
	m.G0 = v1803 + int32(32)
	goto L421
L423:
	;
	v1820 = v1806<<(uint(int64(4))%64) | int64(base.Ui64(v670)>>(uint(int64(60))%64))
	v1825 = v670 & int64(1152921504606846975)
	if base.Ui64(int64(576460752303423489)) <= base.Ui64(v1825) {
		goto L427
	} else {
		goto L428
	}
L424:
	;
	goto L425
L425:
	;
	if v670|v1806 == int64(0) {
		goto L434
	} else {
		goto L435
	}
L426:
	;
	v1838 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v1835))
	if base.Ui64(int64(4503599627370495)) < base.Ui64(v1835) {
		goto L431
	} else {
		goto L432
	}
L427:
	;
	v1835 = v1820 + int64(1)
	goto L426
L428:
	;
	goto L429
L429:
	;
	if v1825 != int64(576460752303423488) {
		v1835 = v1820
		goto L426
	} else {
		goto L430
	}
L430:
	;
	v1835 = v1820&int64(1) + v1820
	goto L426
L431:
	;
	v1839 = int64(0)
	goto L433
L432:
	;
	v1839 = v1835
	goto L433
L433:
	;
	v1912 = v1839
	v1919 = base.I64_extend_i32_u(v1838) + base.I64_extend_i32_u(v1811-int32(15360))
	goto L422
L434:
	;
	if base.Ui32(int32(17406)) < base.Ui32(v1811) {
		goto L437
	} else {
		goto L438
	}
L435:
	;
	if v1810 != int64(32767) {
		goto L434
	} else {
		goto L436
	}
L436:
	;
	v1912 = v1806<<(uint(int64(4))%64) | int64(base.Ui64(v670)>>(uint(int64(60))%64)) | int64(2251799813685248)
	v1919 = int64(2047)
	goto L422
L437:
	;
	v1912 = int64(0)
	v1919 = int64(2047)
	goto L422
L438:
	;
	goto L439
L439:
	;
	v1862 = base.B2i32(v1810 == int64(0))
	if v1810 == int64(0) {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v1863 = int32(15360)
	goto L442
L441:
	;
	v1863 = int32(15361)
	goto L442
L442:
	;
	v1864 = v1863 - v1811
	if int32(112) < v1864 {
		goto L443
	} else {
		goto L444
	}
L443:
	;
	v1867 = int64(0)
	v1912 = v1867
	v1919 = v1867
	goto L422
L444:
	;
	goto L445
L445:
	;
	if v1810 == int64(0) {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v1873 = v1806
	goto L448
L447:
	;
	v1873 = v1806 | int64(281474976710656)
	goto L448
L448:
	;
	F___ashlti3(m, v1803+int32(16), v670, v1873, int32(128)-v1864)
	mBase = m.M
	F___lshrti3(m, v1803, v670, v1873, v1864)
	mBase = m.M
	v1878 = *(*int64)(unsafe.Add(mBase, uint32(v1803)+8))
	v1881 = *(*int64)(unsafe.Add(mBase, uint32(v1803)))
	v1884 = v1878<<(uint(int64(4))%64) | int64(base.Ui64(v1881)>>(uint(int64(60))%64))
	v1886 = *(*int64)(unsafe.Add(mBase, uint32(v1803)+16))
	v1887 = *(*int64)(unsafe.Add(mBase, uint32(v1803)+24))
	v1895 = base.I64_extend_i32_u(base.B2i32(v1811 != v1863)&base.B2i32(v1886|v1887 != int64(0))) | v1881&int64(1152921504606846975)
	if base.Ui64(int64(576460752303423489)) <= base.Ui64(v1895) {
		goto L450
	} else {
		goto L451
	}
L449:
	;
	v1909 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v1905))
	if base.Ui64(int64(4503599627370495)) < base.Ui64(v1905) {
		goto L454
	} else {
		goto L455
	}
L450:
	;
	v1905 = v1884 + int64(1)
	goto L449
L451:
	;
	goto L452
L452:
	;
	if v1895 != int64(576460752303423488) {
		v1905 = v1884
		goto L449
	} else {
		goto L453
	}
L453:
	;
	v1905 = v1884&int64(1) + v1884
	goto L449
L454:
	;
	v1910 = v1905 ^ int64(4503599627370496)
	goto L456
L455:
	;
	v1910 = v1905
	goto L456
L456:
	;
	v1912 = v1910
	v1919 = base.I64_extend_i32_u(v1909)
	goto L422
L457:
	;
	v1963 = int32(31)
	goto L459
L458:
	;
	v1963 = v583 + int32(1)
	goto L459
L459:
	;
	if v482 == int32(1) {
		goto L461
	} else {
		goto L462
	}
L460:
	;
	v2506 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v2507 = *(*int64)(unsafe.Add(mBase, uint32(v46)+112))
	if int64(0) <= v2507 {
		goto L571
	} else {
		goto L572
	}
L461:
	;
	if v443 != 0 {
		goto L464
	} else {
		goto L465
	}
L462:
	;
	goto L463
L463:
	;
	if v443 != 0 {
		goto L532
	} else {
		goto L533
	}
L464:
	;
	v1968 = F_emscripten_builtin_malloc(m, v1963<<(uint(int32(2))%32))
	mBase = m.M
	if v1968 == int32(0) {
		goto L14
	} else {
		goto L467
	}
L465:
	;
	v1971 = v353
	goto L466
L466:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v53)+296)) = int64(0)
	v1976 = int32(0)
	v1979 = v1971
	v1981 = v1963
	goto L470
L467:
	;
	v1971 = v1968
	goto L466
L468:
	;
	v2615 = v1979
	v2616 = v2300
	v2620 = v443
	goto L11
L469:
	;
	v2283 = int32(0)
	v2285 = v53 + int32(296)
	if v2285 != 0 {
		goto L528
	} else {
		goto L529
	}
L470:
	;
	v2001 = v1976
	goto L472
L471:
	;
	v2615 = v1979
	v2616 = int32(0)
	v2620 = int32(1)
	goto L11
L472:
	;
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v2025 != v2026 {
		goto L475
	} else {
		goto L476
	}
L473:
	;
	v2274 = int32(1)
	v2277 = v1981<<(uint(v2274)%32) | v2274
	v2280 = F_emscripten_builtin_realloc(m, v1979, v2277<<(uint(int32(2))%32))
	mBase = m.M
	if v2280 != 0 {
		v1976 = v2270
		v1979 = v2280
		v1981 = v2277
		goto L470
	} else {
		goto L527
	}
L474:
	;
	v2036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2034+v53)+33)))
	if v2036 == int32(0) {
		goto L469
	} else {
		goto L479
	}
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v2025 + int32(1)
	v2031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2025))))
	v2034 = v2031
	goto L474
L476:
	;
	goto L477
L477:
	;
	v2032 = F___shgetc(m, v46)
	mBase = m.M
	v2033 = m.ExcPending
	if v2033 != 0 {
		goto L6
	} else {
		goto L478
	}
L478:
	;
	v2034 = v2032
	goto L474
L479:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+27)) = uint8(v2034)
	v2041 = v53 + int32(28)
	v2043 = v53 + int32(296)
	if v2043 != 0 {
		goto L481
	} else {
		goto L482
	}
L480:
	;
	if v2257 == int32(-2) {
		goto L472
	} else {
		goto L518
	}
L481:
	;
	v2045 = v2043
	goto L483
L482:
	;
	v2045 = int32(4663076)
	goto L483
L483:
	;
	v2046 = *(*int32)(unsafe.Add(mBase, uint32(v2045)))
	v2048 = v53 + int32(27)
	if v2048 == int32(0) {
		goto L487
	} else {
		goto L488
	}
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2045))) = v2209
	v2257 = int32(-2)
	goto L480
L485:
	;
	v2257 = v2204
	goto L480
L486:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2045))) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(25)
	v2204 = int32(-1)
	goto L485
L487:
	;
	if v2046 != 0 {
		goto L486
	} else {
		goto L490
	}
L488:
	;
	goto L489
L489:
	;
	if v2046 != 0 {
		goto L492
	} else {
		goto L493
	}
L490:
	;
	v2257 = int32(0)
	goto L480
L491:
	;
	v2089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2048))))
	v2091 = int32(base.Ui32(v2089) >> (uint(int32(3)) % 32))
	if base.Ui32(int32(7)) < base.Ui32(v2091-int32(16)|(v2046>>(uint(int32(26))%32)+v2091)) {
		goto L486
	} else {
		goto L507
	}
L492:
	;
	goto L491
L493:
	;
	goto L494
L494:
	;
	v2053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2048))))
	v2054 = base.I32_extend8_s(v2053)
	if int32(0) <= v2054 {
		goto L495
	} else {
		goto L496
	}
L495:
	;
	if v2041 != 0 {
		goto L498
	} else {
		goto L499
	}
L496:
	;
	goto L497
L497:
	;
	v2061 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v2062 = *(*int32)(unsafe.Add(mBase, uint32(v2061)))
	if v2062 == int32(0) {
		goto L501
	} else {
		goto L502
	}
L498:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2041))) = v2053
	goto L500
L499:
	;
	goto L500
L500:
	;
	v2257 = base.B2i32(v2054 != int32(0))
	goto L480
L501:
	;
	if v2041 == int32(0) {
		v2204 = int32(1)
		goto L485
	} else {
		goto L504
	}
L502:
	;
	goto L503
L503:
	;
	v2073 = v2053 - int32(194)
	if base.Ui32(int32(50)) < base.Ui32(v2073) {
		goto L486
	} else {
		goto L505
	}
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2041))) = v2054 & int32(57343)
	v2257 = int32(1)
	goto L480
L505:
	;
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v2073<<(uint(int32(2))%32))+uint32(_consts[1482])))
	v2209 = v2080
	goto L484
L507:
	;
	v2104 = v2046
	v2108 = int32(1)
	v2111 = v2048
	v2115 = v2089
	goto L508
L508:
	;
	v2126 = v2108 - int32(1)
	v2133 = v2115&int32(255) - int32(128) | v2104<<(uint(int32(6))%32)
	if int32(0) <= v2133 {
		goto L510
	} else {
		goto L511
	}
L509:
	;
	goto L486
L510:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2045))) = int32(0)
	if v2041 != 0 {
		goto L513
	} else {
		goto L514
	}
L511:
	;
	goto L512
L512:
	;
	if v2126 == int32(0) {
		v2209 = v2133
		goto L484
	} else {
		goto L516
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2041))) = v2133
	goto L515
L514:
	;
	goto L515
L515:
	;
	v2257 = int32(1) - v2126
	goto L480
L516:
	;
	v2144 = v2111 + int32(1)
	v2145 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2144))))
	if v2145 < int32(-64) {
		v2104 = v2133
		v2108 = v2126
		v2111 = v2144
		v2115 = v2145
		goto L508
	} else {
		goto L517
	}
L517:
	;
	goto L509
L518:
	;
	if v2257 == int32(-1) {
		goto L519
	} else {
		goto L520
	}
L519:
	;
	v2300 = int32(0)
	goto L468
L520:
	;
	goto L521
L521:
	;
	if v1979 != 0 {
		goto L522
	} else {
		goto L523
	}
L522:
	;
	v2266 = *(*int32)(unsafe.Add(mBase, uint32(v53)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1979+v2001<<(uint(int32(2))%32)))) = v2266
	v2270 = v2001 + int32(1)
	goto L524
L523:
	;
	v2270 = v2001
	goto L524
L524:
	;
	if v443 == int32(0) {
		v2001 = v2270
		goto L472
	} else {
		goto L525
	}
L525:
	;
	if v2270 != v1981 {
		v2001 = v2270
		goto L472
	} else {
		goto L526
	}
L526:
	;
	goto L473
L527:
	;
	goto L471
L528:
	;
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(v2285)))
	v2288 = v2286
	goto L530
L529:
	;
	v2288 = int32(0)
	goto L530
L530:
	;
	if v2288 == int32(0) {
		v2482 = v2001
		v2484 = v1979
		v2489 = v1979
		v2490 = v2283
		goto L460
	} else {
		goto L531
	}
L531:
	;
	v2300 = v2283
	goto L468
L532:
	;
	v2316 = int32(0)
	v2317 = F_emscripten_builtin_malloc(m, v1963)
	mBase = m.M
	if v2317 == v2316 {
		goto L14
	} else {
		goto L535
	}
L533:
	;
	goto L534
L534:
	;
	if v353 != 0 {
		goto L550
	} else {
		goto L551
	}
L535:
	;
	v2321 = v2316
	v2324 = v2317
	v2326 = v1963
	goto L536
L536:
	;
	v2346 = v2321
	goto L538
L537:
	;
	v2615 = int32(0)
	v2616 = v2324
	v2620 = int32(1)
	goto L11
L538:
	;
	v2370 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v2370 != v2371 {
		goto L541
	} else {
		goto L542
	}
L539:
	;
	v2390 = int32(1)
	v2393 = v2326<<(uint(v2390)%32) | v2390
	v2394 = F_emscripten_builtin_realloc(m, v2324, v2393)
	mBase = m.M
	if v2394 != 0 {
		v2321 = v2388
		v2324 = v2394
		v2326 = v2393
		goto L536
	} else {
		goto L549
	}
L540:
	;
	v2381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2379+v53)+33)))
	if v2381 == int32(0) {
		goto L545
	} else {
		goto L546
	}
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v2370 + int32(1)
	v2376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2370))))
	v2379 = v2376
	goto L540
L542:
	;
	goto L543
L543:
	;
	v2377 = F___shgetc(m, v46)
	mBase = m.M
	v2378 = m.ExcPending
	if v2378 != 0 {
		goto L6
	} else {
		goto L544
	}
L544:
	;
	v2379 = v2377
	goto L540
L545:
	;
	v2482 = v2346
	v2484 = v2324
	v2489 = int32(0)
	v2490 = v2324
	goto L460
L546:
	;
	goto L547
L547:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2346+v2324))) = uint8(v2379)
	v2388 = v2346 + int32(1)
	if v2388 != v2326 {
		v2346 = v2388
		goto L538
	} else {
		goto L548
	}
L548:
	;
	goto L539
L549:
	;
	goto L537
L550:
	;
	v2399 = int32(0)
	goto L553
L551:
	;
	goto L552
L552:
	;
	goto L563
L553:
	;
	v2423 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v2423 != v2424 {
		goto L556
	} else {
		goto L557
	}
L555:
	;
	v2434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2432+v53)+33)))
	if v2434 != 0 {
		goto L560
	} else {
		goto L561
	}
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v2423 + int32(1)
	v2429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2423))))
	v2432 = v2429
	goto L555
L557:
	;
	goto L558
L558:
	;
	v2430 = F___shgetc(m, v46)
	mBase = m.M
	v2431 = m.ExcPending
	if v2431 != 0 {
		goto L6
	} else {
		goto L559
	}
L559:
	;
	v2432 = v2430
	goto L555
L560:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2399+v353))) = uint8(v2432)
	v2399 = v2399 + int32(1)
	goto L553
L561:
	;
	v2482 = v2399
	v2484 = v353
	v2489 = int32(0)
	v2490 = v353
	goto L460
L563:
	;
	v2465 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v2466 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	if v2465 != v2466 {
		goto L566
	} else {
		goto L567
	}
L564:
	;
	v2477 = int32(0)
	v2482 = v2477
	v2484 = v2477
	v2489 = v2477
	v2490 = v2477
	goto L460
L565:
	;
	v2476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2474+v53)+33)))
	if v2476 != 0 {
		goto L563
	} else {
		goto L570
	}
L566:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v2465 + int32(1)
	v2471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2465))))
	v2474 = v2471
	goto L565
L567:
	;
	goto L568
L568:
	;
	v2472 = F___shgetc(m, v46)
	mBase = m.M
	v2473 = m.ExcPending
	if v2473 != 0 {
		goto L6
	} else {
		goto L569
	}
L569:
	;
	v2474 = v2472
	goto L565
L570:
	;
	goto L564
L571:
	;
	v2511 = v2506 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v2511
	v2513 = v2511
	goto L573
L572:
	;
	v2513 = v2506
	goto L573
L573:
	;
	v2514 = *(*int64)(unsafe.Add(mBase, uint32(v46)+120))
	v2515 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	v2518 = v2514 + base.I64_extend_i32_s(v2513-v2515)
	if v2518 == int64(0) {
		v2642 = v2489
		v2643 = v2490
		v2647 = v443
		v2650 = v62
		goto L10
	} else {
		goto L574
	}
L574:
	;
	if v1962|base.B2i32(v604 == v2518) == int32(0) {
		v2642 = v2489
		v2643 = v2490
		v2647 = v443
		v2650 = v62
		goto L10
	} else {
		goto L575
	}
L575:
	;
	if v443 != 0 {
		goto L576
	} else {
		goto L577
	}
L576:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = v2484
	goto L578
L577:
	;
	goto L578
L578:
	;
	if v485 == int32(99) {
		v2544 = v1938
		v2547 = v2489
		v2548 = v2490
		goto L162
	} else {
		goto L579
	}
L579:
	;
	if v2489 != 0 {
		goto L580
	} else {
		goto L581
	}
L580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2489+v2482<<(uint(int32(2))%32)))) = int32(0)
	goto L582
L581:
	;
	goto L582
L582:
	;
	if v2490 == int32(0) {
		goto L583
	} else {
		goto L584
	}
L583:
	;
	v2544 = v1938
	v2547 = v2489
	v2548 = int32(0)
	goto L162
L584:
	;
	goto L585
L585:
	;
	v2537 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2482+v2490))) = uint8(v2537)
	v2544 = v1938
	v2547 = v2489
	v2548 = v2490
	goto L162
L586:
	;
	v2696 = v2581
	v2705 = v2590
	goto L1
L587:
	;
	v2633 = v62
	goto L589
L588:
	;
	v2633 = int32(-1)
	goto L589
L589:
	;
	v2642 = v2615
	v2643 = v2616
	v2647 = v2620
	v2650 = v2633
	goto L10
L590:
	;
	F_emscripten_builtin_free(m, v2643)
	mBase = m.M
	F_emscripten_builtin_free(m, v2642)
	mBase = m.M
	v2696 = v53
	v2705 = v2650
	goto L1
}
func F_view_reloptions(m *base.Module, l0 int32) {
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v7 = F_build_reloptions(m, l0, int32(1), int32(512), int32(12), int32(739104), int32(3))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_visibilitymap_pin_ok(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	if l1 == int32(0) {
		return int32(0)
	} else {
		if l1 < int32(0) {
			v10 = *(*int32)(unsafe.Add(mBase, _consts[15]))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v10+(l1^int32(-1))<<(uint(int32(6))%32))+16))
			v25 = v16
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[16]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+l1<<(uint(int32(6))%32)+int32(-64))+16))
			v25 = v24
		}
		v27 = base.I32_div_u_s(l0, int32(32672))
		return base.B2i32(v25 == v27)
	}
}
func F_void_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_pq_begintypsend(m, v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v13 << (uint(int32(2)) % 32)
		m.G0 = v5 + int32(16)
		return v12
	}
}
