package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecHashGetBucketAndBatch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = (v7 - int32(1)) & l1
	if base.Ui32(int32(2)) <= base.Ui32(v6) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v20 = (v6 - int32(1)) & base.I32_rotr(l1, v16)
	} else {
		v20 = int32(0)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v20
	return
}
func F_ExecHashTableInsert(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v35 int32
	_ = v35
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
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = F_ExecFetchSlotMinimalTuple(m, l1, v12+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if base.Ui32(int32(2)) <= base.Ui32(v18) {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v26 = (v18 - int32(1)) & base.I32_rotr(l2, v23)
		} else {
			v26 = int32(0)
		}
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		if v27 == v26 {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v33 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
			v34 = *(*float64)(unsafe.Add(mBase, uint32(l0)+64))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v37 = v35 + int32(8)
			v38 = F_dense_alloc(m, l0, v37)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = l2
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				if v41 != 0 {
					base.MemoryCopy(m, v38+int32(8), v16, v41)
				} else {
				}
				v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+18)))
				v47 = v45 & int32(_a_F_ExecHashTableInsert_0)
				*(*uint16)(unsafe.Add(mBase, uint32(v38)+18)) = uint16(v47)
				v50 = (v29 - int32(1)) & l2 << (uint(int32(2)) % 32)
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v50+v51)))
				*(*int32)(unsafe.Add(mBase, uint32(v38))) = v53
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v55+v50))) = v38
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v58 != int32(1) {
				} else {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if base.B2i32(int32(1073741823) < v61)|base.B2i32(base.F64_lt(base.F64_convert_i32_s(v61), base.F64_sub(v34, v33)) == int32(0)) != 0 {
					} else {
						v71 = v61 << (uint(int32(1)) % 32)
						if base.Ui32(int32(268435455)) < base.Ui32(v71) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v71
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v75 + int32(1)
						}
					}
				}
				v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				v81 = v80 + v37
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v81
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				if base.Ui32(v83) < base.Ui32(v81) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v81
				} else {
				}
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if base.Ui32(v87<<(uint(int32(2))%32)+v81) <= base.Ui32(v86) {
					v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
					if v106 == int32(1) {
						F_pfree(m, v16)
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return
						} else {
							m.G0 = v12 + int32(16)
							return
						}
					} else {
						m.G0 = v12 + int32(16)
						return
					}
				} else {
					F_ExecHashIncreaseNumBatches(m, l0)
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
						if v106 == int32(1) {
							F_pfree(m, v16)
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return
							} else {
								m.G0 = v12 + int32(16)
								return
							}
						} else {
							m.G0 = v12 + int32(16)
							return
						}
					}
				}
			}
		} else {
			v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
			F_ExecHashJoinSaveTuple(m, v16, l2, v94+v26<<(uint(int32(2))%32), l0)
			mBase = m.M
			v99 = m.ExcPending
			if v99 != 0 {
				return
			} else {
				v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
				if v106 == int32(1) {
					F_pfree(m, v16)
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return
					} else {
						m.G0 = v12 + int32(16)
						return
					}
				} else {
					m.G0 = v12 + int32(16)
					return
				}
			}
		}
	}
}
func F__hash_binsearch_last(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	v3 = int32(0)
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v8) < base.Ui32(int32(25)) {
		v65 = v3
	} else {
		v14 = int32(base.Ui32(v8+int32(_a_F__hash_binsearch_last_0)) >> (uint(int32(2)) % 32))
		if v14&int32(_a_F__hash_binsearch_last_1) == int32(0) {
			v65 = v3
		} else {
			v23 = v14
			v24 = v3
			for {
				v28 = int32(_a_F__hash_binsearch_last_1)
				v33 = int32(1)
				v36 = int32(base.Ui32(v24&v28+v23&v28+v33) >> (uint(v33) % 32))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(20)+v36<<(uint(int32(2))%32))))
				v45 = l0 + v42&int32(_a_F__hash_binsearch_last_2)
				v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+6)))
				if int32(0) <= v48 {
					v51 = int32(8)
				} else {
					v51 = int32(16)
				}
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v45+v51)))
				v54 = base.B2i32(base.Ui32(l1) < base.Ui32(v53))
				if base.Ui32(l1) < base.Ui32(v53) {
					v55 = v36 - v33
				} else {
					v55 = v23
				}
				if base.Ui32(l1) < base.Ui32(v53) {
					v58 = v24
				} else {
					v58 = v36
				}
				if base.Ui32(v58&int32(_a_F__hash_binsearch_last_1)) < base.Ui32(v55&int32(_a_F__hash_binsearch_last_1)) {
					v23 = v55
					v24 = v58
					continue
				} else {
					break
				}
				break
			}
			v65 = v58
		}
	}
	return v65 & int32(_a_F__hash_binsearch_last_1)
}
func F__hash_convert_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v6 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v10 = int32(1)
		v12 = F_index_getprocinfo(m, l0, v10, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v18 = F_FunctionCall1Coll(m, v12, v17, v9)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v18
				v21 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v21)
				return v6 ^ int32(1)
			}
		}
	} else {
		return v6 ^ int32(1)
	}
}
func F__hash_getbucketbuf_from_hashkey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v48 int32
	_ = v48
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v5
	v19 = F__hash_getcachedmetap(m, l0, v12+int32(12), v5)
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
	v30 = v19
	goto L4
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L43
	}
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	v39 = l1 & v36
	if base.Ui32(v39) <= base.Ui32(v35) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v115 != 0 {
		goto L36
	} else {
		goto L37
	}
L6:
	;
	if v42 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v41 = int32(-1)
	goto L9
L8:
	;
	v41 = v37
	goto L9
L9:
	;
	v42 = v41 & v39
	goto L6
L10:
	;
	v43 = int32(1)
	v44 = v42 + v43
	v48 = v44 - v43
	if base.Ui32(int32(2)) <= base.Ui32(v44) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v75 = int32(1)
	goto L12
L12:
	;
	v76 = F_ReadBuffer(m, l0, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L21
	}
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v30+v67<<(uint(int32(2))%32))+48))
	v72 = v71 + v44
	if v72 == int32(-1) {
		goto L3
	} else {
		goto L20
	}
L14:
	;
	v54 = int32(32) - base.I32_clz(v48)
	goto L16
L15:
	;
	v54 = int32(0)
	goto L16
L16:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v54) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v57 = int32(3)
	v67 = int32(base.Ui32(v48)>>(uint(v54-v57)%32))&v57 | v54<<(uint(int32(2))%32) - int32(30)
	goto L19
L18:
	;
	v67 = v54
	goto L19
L19:
	;
	goto L13
L20:
	;
	v75 = v72
	goto L12
L21:
	;
	if base.B2i32(l2 == int32(-1)) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_LockBuffer(m, v76, l2)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F__hash_checkpage(m, l0, v76, int32(2))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	if v76 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+16)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v103+v102)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	if base.Ui32(v106) < base.Ui32(v105) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getbucketbuf_from_hashkey[0]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88+(v76^int32(-1))<<(uint(int32(2))%32))))
	v102 = v94
	goto L27
L29:
	;
	goto L30
L30:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getbucketbuf_from_hashkey[1]))
	v102 = v96 + v76<<(uint(int32(13))%32) + int32(-8192)
	goto L27
L31:
	;
	F_UnlockReleaseBuffer(m, v76)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	goto L5
L34:
	;
	v113 = F__hash_getcachedmetap(m, l0, v12+int32(12), int32(1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v30 = v113
	goto L4
L36:
	;
	F_ReleaseBuffer(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if l3 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v30
	goto L42
L41:
	;
	goto L42
L42:
	;
	m.G0 = v12 + int32(16)
	return v76
L43:
	;
	F_errmsg_internal(m, int32(_a_F__hash_getbucketbuf_from_hashkey_0), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F__hash_getbucketbuf_from_hashkey_1), int32(75), int32(_a_F__hash_getbucketbuf_from_hashkey_2))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__hash_init(m *base.Module, l0 int32, l1 float64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 float64
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int64
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v522 int64
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v542 int64
	_ = v542
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	v23 = F_RelationGetNumberOfBlocksInFork(m, l0, l2)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L2
	} else {
		goto L140
	}
L2:
	;
	return int32(0)
L3:
	;
	if v23 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+118)))
	if v30 != int32(112) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L2
	} else {
		goto L137
	}
L7:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v46 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v45 = base.B2i32(l2 == int32(3))
	goto L7
L9:
	;
	v33 = int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[0]))
	if int32(0) < v35 {
		v45 = v33
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v38 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v39 == int32(0) {
		v45 = v33
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	v47 = int32(10)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v52 = base.I32_div_s(v48<<(uint(int32(13))%32), int32(2000))
	if v52 <= v47 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v58 = int32(307)
	goto L15
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+6)))
	v71 = int32(4)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v61+v63*int32(0)<<(uint(int32(2))%32)+v71-v71)))
	goto L19
L16:
	;
	v55 = v47
	goto L18
L17:
	;
	v55 = v52
	goto L18
L18:
	;
	v58 = v55
	goto L15
L19:
	;
	v77 = F__hash_getnewbuf(m, l0, int32(0), l2)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v80 = v58 & int32(_a_F__hash_init_0)
	v86 = base.F64_div(l1, base.F64_convert_i32_u(v80))
	if base.F64_le(v86, float64(2)) != 0 {
		v95 = int32(2)
		goto L22
	} else {
		goto L23
	}
L21:
	;
	F_MarkBufferDirty(m, v77)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L2
	} else {
		goto L35
	}
L22:
	;
	v96 = F__hash_spareindex(m, v95)
	mBase = m.M
	if v77 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	if base.F64_ge(v86, float64(1.073741824e+09)) != 0 {
		v95 = int32(1073741824)
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v93 = F__hash_spareindex(m, base.I32_trunc_sat_f64_u(v86))
	mBase = m.M
	v94 = F__hash_get_totalbuckets(m, v93)
	mBase = m.M
	v95 = v94
	goto L22
L25:
	;
	goto L30
L26:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v100+(v77^int32(-1))<<(uint(int32(2))%32))))
	v114 = v106
	goto L25
L27:
	;
	goto L28
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v114 = v108 + v77<<(uint(int32(13))%32) + int32(-8192)
	goto L25
L30:
	;
	goto L31
L31:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+16)))
	v119 = v114 + v118
	*(*int64)(unsafe.Add(mBase, uint32(v119)+8)) = int64(-36028758364258305)
	*(*int64)(unsafe.Add(mBase, uint32(v119))) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v114)+68)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v114)+32)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v114)+24)) = int64(17284990528)
	*(*uint16)(unsafe.Add(mBase, uint32(v114)+40)) = uint16(v80)
	*(*int32)(unsafe.Add(mBase, uint32(v114)+72)) = v75
	v132 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v114)+48)) = v95 - v132
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+19)))
	v139 = v135<<(uint(int32(8))%32) - int32(40)
	*(*uint16)(unsafe.Add(mBase, uint32(v114)+42)) = uint16(v139)
	v141 = int32(-1)
	v144 = v95 + v132
	if v144&v95 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v151 = v141<<(uint(int32(32)-base.I32_clz(v144))%32) ^ v141
	goto L34
L33:
	;
	v151 = v95
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+52)) = v151
	v153 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v114)+56)) = int32(base.Ui32(v151) >> (uint(v153) % 32))
	v160 = base.I32_clz(v139&int32(_a_F__hash_init_1)) ^ int32(31)
	v162 = v160 + int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v114)+46)) = uint16(v162)
	v165 = v153 << (uint(v160) % 32)
	*(*uint16)(unsafe.Add(mBase, uint32(v114)+44)) = uint16(v165)
	v168 = v114 + int32(76)
	v169 = int32(0)
	base.MemoryFill(m, v168, v169, int32(392))
	base.MemoryFill(m, v114+int32(468), v169, int32(_a_F__hash_init_2))
	*(*int32)(unsafe.Add(mBase, uint32(v168+v96<<(uint(int32(2))%32)))) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v114)+64)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v114)+60)) = v96
	v185 = int32(_a_F__hash_init_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v114)+12)) = uint16(v185)
	goto L21
L35:
	;
	if v77 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v45 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192+(v77^int32(-1))<<(uint(int32(2))%32))))
	v206 = v198
	goto L36
L38:
	;
	goto L39
L39:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v206 = v200 + v77<<(uint(int32(13))%32) + int32(-8192)
	goto L36
L40:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21)+32)) = l1
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v206)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v208
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206)+40)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+44)) = uint16(v210)
	F_XLogBeginInsert(m)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v206)+48))
	F_LockBuffer(m, v77, int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L2
	} else {
		goto L51
	}
L43:
	;
	F_XLogRegisterData(m, v21+int32(32), int32(14))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	F_XLogRegisterBuffer(m, int32(0), v77, int32(14))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v225 = F_XLogInsert(m, int32(12), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	if v77 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v244))) = base.I64_rotr(v225, int64(32))
	goto L42
L48:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v230+(v77^int32(-1))<<(uint(int32(2))%32))))
	v244 = v236
	goto L47
L49:
	;
	goto L50
L50:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v244 = v238 + v77<<(uint(int32(13))%32) + int32(-8192)
	goto L47
L51:
	;
	v254 = v249 + int32(1)
	if v254 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_LockBuffer(m, v77, int32(2))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L2
	} else {
		goto L103
	}
L53:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[3]))
	if v258 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L2
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v262 = F__hash_getnewbuf(m, l0, int32(1), l2)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L2
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v206)+48))
	if int32(0) <= v262 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v282)+16)))
	v284 = v283 + v282
	*(*int32)(unsafe.Add(mBase, uint32(v284)+12)) = int32(-8388606)
	*(*int64)(unsafe.Add(mBase, uint32(v284)+4)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v284))) = v264
	F_MarkBufferDirty(m, v262)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L2
	} else {
		goto L63
	}
L60:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v282 = v268 + v262<<(uint(int32(13))%32) + int32(-8192)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v275+(v262^int32(-1))<<(uint(int32(2))%32))))
	v282 = v281
	goto L59
L63:
	;
	if v45 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if int32(0) <= v262 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	goto L66
L66:
	;
	F_UnlockReleaseBuffer(m, v262)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L2
	} else {
		goto L72
	}
L67:
	;
	F_log_newpage(m, l0, l2, int32(1), v310, int32(1))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L2
	} else {
		goto L71
	}
L68:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v310 = v296 + v262<<(uint(int32(13))%32) + int32(-8192)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v303+(v262^int32(-1))<<(uint(int32(2))%32))))
	v310 = v309
	goto L67
L71:
	;
	goto L66
L72:
	;
	if v249 == int32(0) {
		goto L52
	} else {
		goto L73
	}
L73:
	;
	v324 = int32(1)
	goto L74
L74:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[3]))
	if v340 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L52
L76:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L2
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v343 = int32(1)
	v344 = v324 + v343
	v348 = v344 - v343
	if base.Ui32(int32(2)) <= base.Ui32(v344) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L78
L80:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v206+int32(72)+v367<<(uint(int32(2))%32))))
	v372 = v371 + v344
	v373 = F__hash_getnewbuf(m, l0, v372, l2)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L2
	} else {
		goto L87
	}
L81:
	;
	v354 = int32(32) - base.I32_clz(v348)
	goto L83
L82:
	;
	v354 = int32(0)
	goto L83
L83:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v354) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v357 = int32(3)
	v367 = int32(base.Ui32(v348)>>(uint(v354-v357)%32))&v357 | v354<<(uint(int32(2))%32) - int32(30)
	goto L86
L85:
	;
	v367 = v354
	goto L86
L86:
	;
	goto L80
L87:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v206)+48))
	v376 = int32(0)
	v377 = base.B2i32(v376 <= v373)
	if v377 == v376 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v395)+16)))
	v397 = v396 + v395
	*(*int32)(unsafe.Add(mBase, uint32(v397)+12)) = int32(-8388606)
	*(*int32)(unsafe.Add(mBase, uint32(v397)+8)) = v324
	*(*int32)(unsafe.Add(mBase, uint32(v397)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v397))) = v375
	F_MarkBufferDirty(m, v373)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L2
	} else {
		goto L92
	}
L89:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v381+(v373^int32(-1))<<(uint(int32(2))%32))))
	v395 = v387
	goto L88
L90:
	;
	goto L91
L91:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v395 = v389 + v373<<(uint(int32(13))%32) + int32(-8192)
	goto L88
L92:
	;
	if v45 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v377 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L95
L95:
	;
	F_UnlockReleaseBuffer(m, v373)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L2
	} else {
		goto L101
	}
L96:
	;
	F_log_newpage(m, l0, l2, v372, v423, int32(1))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L2
	} else {
		goto L100
	}
L97:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v409+(v373^int32(-1))<<(uint(int32(2))%32))))
	v423 = v415
	goto L96
L98:
	;
	goto L99
L99:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v423 = v417 + v373<<(uint(int32(13))%32) + int32(-8192)
	goto L96
L100:
	;
	goto L95
L101:
	;
	if v324 != v249 {
		v324 = v344
		goto L74
	} else {
		goto L102
	}
L102:
	;
	goto L75
L103:
	;
	v452 = v249 + int32(2)
	v453 = F__hash_getnewbuf(m, l0, v452, l2)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	v455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206)+44)))
	if v453 < int32(0) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	F_MarkBufferDirty(m, v453)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L2
	} else {
		goto L116
	}
L106:
	;
	goto L111
L107:
	;
	v460 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v460+(v453^int32(-1))<<(uint(int32(2))%32))))
	v474 = v466
	goto L106
L108:
	;
	goto L109
L109:
	;
	v468 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v474 = v468 + v453<<(uint(int32(13))%32) + int32(-8192)
	goto L106
L111:
	;
	goto L112
L112:
	;
	v476 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v474)+16)))
	v477 = v474 + v476
	*(*int64)(unsafe.Add(mBase, uint32(v477)+8)) = int64(-36028775544127489)
	*(*int64)(unsafe.Add(mBase, uint32(v477))) = int64(-1)
	if v455 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	base.MemoryFill(m, v474+int32(24), int32(255), v455)
	goto L115
L114:
	;
	goto L115
L115:
	;
	v487 = v455 + int32(24)
	*(*uint16)(unsafe.Add(mBase, uint32(v474)+12)) = uint16(v487)
	goto L105
L116:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v206)+68))
	if base.Ui32(int32(1024)) <= base.Ui32(v491) {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206+v491<<(uint(int32(2))%32))+468)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v206)+68)) = v491 + int32(1)
	F_MarkBufferDirty(m, v77)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	if v45 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206)+44)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+32)) = uint16(v503)
	F_XLogBeginInsert(m)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L2
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	F_UnlockReleaseBuffer(m, v453)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L2
	} else {
		goto L135
	}
L122:
	;
	F_XLogRegisterData(m, v21+int32(32), int32(2))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	F_XLogRegisterBuffer(m, int32(0), v453, int32(6))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	F_XLogRegisterBuffer(m, int32(1), v77, int32(8))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	v522 = F_XLogInsert(m, int32(12), int32(16))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L2
	} else {
		goto L126
	}
L126:
	;
	if v453 < int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v542 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v541))) = base.I64_rotr(v522, v542)
	if v77 < int32(0) {
		goto L132
	} else {
		goto L133
	}
L128:
	;
	v527 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v527+(v453^int32(-1))<<(uint(int32(2))%32))))
	v541 = v533
	goto L127
L129:
	;
	goto L130
L130:
	;
	v535 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v541 = v535 + v453<<(uint(int32(13))%32) + int32(-8192)
	goto L127
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v566)+4)) = base.I32_wrap_i64(v522)
	*(*int32)(unsafe.Add(mBase, uint32(v566))) = base.I32_wrap_i64(int64(base.Ui64(v522) >> (uint(v542) % 64)))
	goto L121
L132:
	;
	v552 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v552+(v77^int32(-1))<<(uint(int32(2))%32))))
	v566 = v558
	goto L131
L133:
	;
	goto L134
L134:
	;
	v560 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v566 = v560 + v77<<(uint(int32(13))%32) + int32(-8192)
	goto L131
L135:
	;
	F_UnlockReleaseBuffer(m, v77)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L2
	} else {
		goto L136
	}
L136:
	;
	m.G0 = v21 + int32(48)
	return v254
L137:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v585 + int32(4)
	F_errmsg_internal(m, int32(_a_F__hash_init_4), v21+int32(16))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L2
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F__hash_init_5), int32(345), int32(_a_F__hash_init_6))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L2
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L2
	} else {
		goto L141
	}
L141:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v606 + int32(4)
	F_errmsg(m, int32(_a_F__hash_init_7), v21)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L2
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F__hash_init_5), int32(455), int32(_a_F__hash_init_6))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L2
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__hash_load_qualified_items(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	v5 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if l3 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v201
L2:
	;
	v120 = l2
	v125 = int32(408)
	goto L31
L3:
	;
	if l2 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v24) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	return int32(408)
L7:
	;
	goto L8
L8:
	;
	goto L2
L9:
	;
	v32 = int32(base.Ui32(v24+int32(_a_F__hash_load_qualified_items_0)) >> (uint(int32(2)) % 32))
	goto L11
L10:
	;
	v32 = int32(0)
	goto L11
L11:
	;
	v34 = v32 & int32(_a_F__hash_load_qualified_items_1)
	if base.Ui32(v34) < base.Ui32(l2) {
		v201 = v5
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v42 = l2
	v47 = v5
	goto L13
L13:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)))
	v56 = v42
	goto L15
L14:
	;
	v201 = v112
	goto L1
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(20)+v56&int32(_a_F__hash_load_qualified_items_1)<<(uint(int32(2))%32))))
	v73 = l1 + v70&int32(_a_F__hash_load_qualified_items_2)
	if v51&int32(1) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v96 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+6)))
	if int32(0) <= v96 {
		goto L26
	} else {
		goto L27
	}
L17:
	;
	goto L16
L18:
	;
	v89 = v56 + int32(1)
	if base.Ui32(v89&int32(_a_F__hash_load_qualified_items_1)) <= base.Ui32(v34) {
		v56 = v89
		goto L15
	} else {
		goto L24
	}
L19:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	v83 = int32(_a_F__hash_load_qualified_items_3)
	if base.B2i32(v80 != int32(1))|base.B2i32(v70&v83 != v83) != 0 {
		goto L17
	} else {
		goto L23
	}
L20:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)))
	if v76 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+7)))
	if v77&int32(32) != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	goto L18
L24:
	;
	v201 = v47
	goto L1
L25:
	;
	if v93 != v101 {
		v201 = v47
		goto L1
	} else {
		goto L29
	}
L26:
	;
	v99 = int32(8)
	goto L28
L27:
	;
	v99 = int32(16)
	goto L28
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v73+v99)))
	goto L25
L29:
	;
	v105 = v12 + int32(52) + v47<<(uint(int32(3))%32)
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v105)+4)) = uint16(v106)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v108
	*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)) = uint16(v56)
	v111 = int32(1)
	v112 = v47 + v111
	v114 = v56 + v111
	if base.Ui32(v114&int32(_a_F__hash_load_qualified_items_1)) <= base.Ui32(v34) {
		v42 = v114
		v47 = v112
		goto L13
	} else {
		goto L30
	}
L30:
	;
	goto L14
L31:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)))
	v134 = v120
	goto L33
L32:
	;
	v201 = v181
	goto L1
L33:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(20)+v134&int32(_a_F__hash_load_qualified_items_1)<<(uint(int32(2))%32))))
	v151 = l1 + v148&int32(_a_F__hash_load_qualified_items_2)
	if v129&int32(1) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v173 = int32(*(*int16)(unsafe.Add(mBase, uint32(v151)+6)))
	if int32(0) <= v173 {
		goto L44
	} else {
		goto L45
	}
L35:
	;
	goto L34
L36:
	;
	v167 = v134 - int32(1)
	if v167&int32(_a_F__hash_load_qualified_items_1) != 0 {
		v134 = v167
		goto L33
	} else {
		goto L42
	}
L37:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	v161 = int32(_a_F__hash_load_qualified_items_3)
	if base.B2i32(v158 != int32(1))|base.B2i32(v148&v161 != v161) != 0 {
		goto L35
	} else {
		goto L41
	}
L38:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)))
	if v154 != 0 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+7)))
	if v155&int32(32) != 0 {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	goto L36
L42:
	;
	v201 = v125
	goto L1
L43:
	;
	if v170 != v178 {
		v201 = v125
		goto L1
	} else {
		goto L47
	}
L44:
	;
	v176 = int32(8)
	goto L46
L45:
	;
	v176 = int32(16)
	goto L46
L46:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v151+v176)))
	goto L43
L47:
	;
	v180 = int32(1)
	v181 = v125 - v180
	v184 = v12 + int32(52) + v181<<(uint(int32(3))%32)
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v184)+4)) = uint16(v185)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v187
	*(*uint16)(unsafe.Add(mBase, uint32(v184)+6)) = uint16(v134)
	v191 = v134 - v180
	if v191&int32(_a_F__hash_load_qualified_items_1) != 0 {
		v120 = v191
		v125 = v181
		goto L31
	} else {
		goto L48
	}
L48:
	;
	goto L32
}
func F__hash_splitbucket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int64
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v313 int32
	_ = v313
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v524 int64
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v573 int32
	_ = v573
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v753 int64
	_ = v753
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v773 int64
	_ = v773
	var v783 int32
	_ = v783
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v797 int32
	_ = v797
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	v11 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(2464)
	m.G0 = v29
	if l4 < v11 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+16)))
	if l5 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34+(l4^int32(-1))<<(uint(int32(2))%32))))
	v48 = v40
	goto L1
L3:
	;
	goto L4
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v48 = v42 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	if l4 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53+(l5^int32(-1))<<(uint(int32(2))%32))))
	v67 = v59
	goto L5
L7:
	;
	goto L8
L8:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v67 = v61 + l5<<(uint(int32(13))%32) + int32(-8192)
	goto L5
L9:
	;
	if l5 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[2]))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72+(l4^int32(-1))<<(uint(int32(6))%32))+16))
	v87 = v78
	goto L9
L11:
	;
	goto L12
L12:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[3]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v80+l4<<(uint(int32(6))%32)+int32(-64))+16))
	v87 = v86
	goto L9
L13:
	;
	F_PredicateLockPageSplit(m, l0, v87, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[2]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91+(l5^int32(-1))<<(uint(int32(6))%32))+16))
	v106 = v97
	goto L13
L15:
	;
	goto L16
L16:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[3]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99+l5<<(uint(int32(6))%32)+int32(-64))+16))
	v106 = v105
	goto L13
L17:
	;
	return
L18:
	;
	v120 = l5
	v122 = l4
	v123 = v48
	v124 = v11
	v125 = v67
	v127 = v11
	v128 = v49 + v48
	goto L20
L19:
	;
	v675 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v674)+16)))
	F_LockBuffer(m, l5, int32(2))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L17
	} else {
		goto L129
	}
L20:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+12)))
	if base.Ui32(v135) < base.Ui32(int32(25)) {
		v466 = v120
		v470 = v124
		v471 = v125
		v473 = v127
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v668 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v674 = v668 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L19
L22:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if l4 == v122 {
		goto L82
	} else {
		goto L83
	}
L23:
	;
	v139 = v135 + int32(_a_F__hash_splitbucket_0)
	if v139&int32(_a_F__hash_splitbucket_1) == int32(0) {
		v466 = v120
		v470 = v124
		v471 = v125
		v473 = v127
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v161 = int32(1)
	v162 = v120
	v166 = v124
	v167 = v125
	v169 = v127
	goto L25
L25:
	;
	v177 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+12)) = uint8(v177)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v123+int32(20)+v161<<(uint(int32(2))%32))))
	v183 = int32(_a_F__hash_splitbucket_2)
	if v182&v183 == v183 {
		v437 = v162
		v441 = v166
		v442 = v167
		v444 = v169
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v466 = v437
	v470 = v441
	v471 = v442
	v473 = v444
	goto L22
L27:
	;
	if v161 != int32(base.Ui32(v139)>>(uint(int32(2))%32))&int32(_a_F__hash_splitbucket_3) {
		v161 = v161 + int32(1)
		v162 = v437
		v166 = v441
		v167 = v442
		v169 = v444
		goto L25
	} else {
		goto L80
	}
L28:
	;
	v189 = v123 + v182&int32(_a_F__hash_splitbucket_4)
	if l6 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v193 = F_hash_search(m, l6, v189, int32(0), v29+int32(12))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L17
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v198 = int32(*(*int16)(unsafe.Add(mBase, uint32(v189)+6)))
	if int32(0) <= v198 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+12)))
	if v195 != 0 {
		v437 = v162
		v441 = v166
		v442 = v167
		v444 = v169
		goto L27
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v205 = v203 & l8
	if base.Ui32(v205) <= base.Ui32(l7) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v201 = int32(8)
	goto L37
L36:
	;
	v201 = int32(16)
	goto L37
L37:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v189+v201)))
	goto L34
L38:
	;
	if v207&v205 != l3 {
		v437 = v162
		v441 = v166
		v442 = v167
		v444 = v169
		goto L27
	} else {
		goto L42
	}
L39:
	;
	v207 = int32(-1)
	goto L41
L40:
	;
	v207 = l9
	goto L41
L41:
	;
	goto L38
L42:
	;
	v210 = F_CopyIndexTuple(m, v189)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L17
	} else {
		goto L43
	}
L43:
	;
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v210)+6)))
	v214 = v212 | int32(_a_F__hash_splitbucket_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v210)+6)) = uint16(v214)
	v217 = v166 & int32(_a_F__hash_splitbucket_3)
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+14)))
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+12)))
	v222 = v220 - v221
	v224 = (v217 + int32(1)) << (uint(int32(2)) % 32)
	if v224 <= v222 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v234 = (v212&int32(_a_F__hash_splitbucket_6) + int32(7)) & int32(_a_F__hash_splitbucket_7)
	if base.Ui32(v228) < base.Ui32(v169+v234) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v228 = v222 - v224
	goto L47
L46:
	;
	v228 = int32(0)
	goto L47
L47:
	;
	goto L44
L48:
	;
	v237 = int32(_a_F__hash_splitbucket_8)
	v239 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4])) = v239 + int32(1)
	F__hash_pgaddmultitup(m, l0, v162, v29+int32(16), v29+int32(1648), v217)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L17
	} else {
		goto L51
	}
L49:
	;
	v400 = v162
	v404 = v166
	v405 = v167
	v407 = v169
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29+int32(16)+v404&int32(_a_F__hash_splitbucket_3)<<(uint(int32(2))%32)))) = v210
	v437 = v400
	v441 = v404 + int32(1)
	v442 = v405
	v444 = v407 + v234
	goto L27
L51:
	;
	F_MarkBufferDirty(m, v162)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+118)))
	if v252 != int32(112) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v293 = int32(0)
	v294 = int32(_a_F__hash_splitbucket_8)
	v296 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4])) = v296 - int32(1)
	F_LockBuffer(m, v162, v293)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L17
	} else {
		goto L67
	}
L54:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[5]))
	if v256 <= int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v259 != 0 {
		goto L53
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L17
	} else {
		goto L60
	}
L58:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v260 != 0 {
		goto L53
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	F_XLogRegisterBuffer(m, int32(0), v162, int32(9))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L17
	} else {
		goto L61
	}
L61:
	;
	v269 = F_XLogInsert(m, int32(12), int32(80))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L17
	} else {
		goto L62
	}
L62:
	;
	if v162 < int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v288))) = base.I64_rotr(v269, int64(32))
	goto L53
L64:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v274+(v162^int32(-1))<<(uint(int32(2))%32))))
	v288 = v280
	goto L63
L65:
	;
	goto L66
L66:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v288 = v282 + v162<<(uint(int32(13))%32) + int32(-8192)
	goto L63
L67:
	;
	if v217 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v313 = v293
	goto L71
L69:
	;
	goto L70
L70:
	;
	v367 = F__hash_addovflpage(m, l0, l1, v162, base.B2i32(l5 == v162))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L17
	} else {
		goto L76
	}
L71:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(16)+v313<<(uint(int32(2))%32))))
	F_pfree(m, v334)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L17
	} else {
		goto L73
	}
L72:
	;
	goto L70
L73:
	;
	v338 = v313 + int32(1)
	if v338 != v217 {
		v313 = v338
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v387 = int32(0)
	v400 = v367
	v404 = v387
	v405 = v386
	v407 = v387
	goto L50
L76:
	;
	if v367 < int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v372+(v367^int32(-1))<<(uint(int32(2))%32))))
	v386 = v378
	goto L75
L78:
	;
	goto L79
L79:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v386 = v380 + v367<<(uint(int32(13))%32) + int32(-8192)
	goto L75
L80:
	;
	goto L26
L81:
	;
	if v481 == int32(-1) {
		goto L88
	} else {
		goto L89
	}
L82:
	;
	F_LockBuffer(m, l4, int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L17
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	F_UnlockReleaseBuffer(m, v122)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L17
	} else {
		goto L86
	}
L85:
	;
	goto L81
L86:
	;
	goto L81
L87:
	;
	goto L21
L88:
	;
	v490 = int32(_a_F__hash_splitbucket_8)
	v492 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4])) = v492 + int32(1)
	F__hash_pgaddmultitup(m, l0, v466, v29+int32(16), v29+int32(1648), v470&int32(_a_F__hash_splitbucket_3))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L17
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v639 = F_ReadBuffer(m, l0, v481)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L17
	} else {
		goto L122
	}
L91:
	;
	F_MarkBufferDirty(m, v466)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L17
	} else {
		goto L92
	}
L92:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+118)))
	if v507 != int32(112) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v548 = int32(_a_F__hash_splitbucket_8)
	v550 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4])) = v550 - int32(1)
	if l5 == v466 {
		goto L108
	} else {
		goto L109
	}
L94:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[5]))
	if v511 <= int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v514 != 0 {
		goto L93
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L17
	} else {
		goto L100
	}
L98:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v515 != 0 {
		goto L93
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	F_XLogRegisterBuffer(m, int32(0), v466, int32(9))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L17
	} else {
		goto L101
	}
L101:
	;
	v524 = F_XLogInsert(m, int32(12), int32(80))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L17
	} else {
		goto L102
	}
L102:
	;
	if v466 < int32(0) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v543))) = base.I64_rotr(v524, int64(32))
	goto L93
L104:
	;
	v529 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v529+(v466^int32(-1))<<(uint(int32(2))%32))))
	v543 = v535
	goto L103
L105:
	;
	goto L106
L106:
	;
	v537 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v543 = v537 + v466<<(uint(int32(13))%32) + int32(-8192)
	goto L103
L107:
	;
	v561 = v470 & int32(_a_F__hash_splitbucket_3)
	if v561 != 0 {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	F_LockBuffer(m, l5, int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L17
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	F_UnlockReleaseBuffer(m, v466)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L17
	} else {
		goto L112
	}
L111:
	;
	goto L107
L112:
	;
	goto L107
L113:
	;
	v573 = int32(0)
	goto L116
L114:
	;
	goto L115
L115:
	;
	F_LockBuffer(m, l4, int32(2))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L17
	} else {
		goto L120
	}
L116:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(16)+v573<<(uint(int32(2))%32))))
	F_pfree(m, v594)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L17
	} else {
		goto L118
	}
L117:
	;
	goto L115
L118:
	;
	v598 = v573 + int32(1)
	if v598 != v561 {
		v573 = v598
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	if int32(0) <= l4 {
		goto L87
	} else {
		goto L121
	}
L121:
	;
	v632 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v632+(l4^int32(-1))<<(uint(int32(2))%32))))
	v674 = v638
	goto L19
L122:
	;
	F_LockBuffer(m, v639, int32(1))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L17
	} else {
		goto L123
	}
L123:
	;
	F__hash_checkpage(m, l0, v639, int32(1))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L17
	} else {
		goto L124
	}
L124:
	;
	if v639 < int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v665 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v664)+16)))
	v120 = v466
	v122 = v639
	v123 = v664
	v124 = v470
	v125 = v471
	v127 = v473
	v128 = v665 + v664
	goto L20
L126:
	;
	v650 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v650+(v639^int32(-1))<<(uint(int32(2))%32))))
	v664 = v656
	goto L125
L127:
	;
	goto L128
L128:
	;
	v658 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v664 = v658 + v639<<(uint(int32(13))%32) + int32(-8192)
	goto L125
L129:
	;
	v679 = v674 + v675
	if l5 < int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v698 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v697)+16)))
	v699 = int32(_a_F__hash_splitbucket_8)
	v701 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4])) = v701 + int32(1)
	v705 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v679)+12)))
	v707 = v705 & int32(_a_F__hash_splitbucket_9)
	*(*uint16)(unsafe.Add(mBase, uint32(v679)+12)) = uint16(v707)
	v709 = v697 + v698
	v710 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v709)+12)))
	v712 = v710 & int32(_a_F__hash_splitbucket_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v709)+12)) = uint16(v712)
	v714 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v679)+12)))
	v716 = v714 | int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v679)+12)) = uint16(v716)
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L17
	} else {
		goto L134
	}
L131:
	;
	v683 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v683+(l5^int32(-1))<<(uint(int32(2))%32))))
	v697 = v689
	goto L130
L132:
	;
	goto L133
L133:
	;
	v691 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v697 = v691 + l5<<(uint(int32(13))%32) + int32(-8192)
	goto L130
L134:
	;
	F_MarkBufferDirty(m, l5)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L17
	} else {
		goto L135
	}
L135:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+118)))
	if v723 != int32(112) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v804 = int32(_a_F__hash_splitbucket_8)
	v806 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4])) = v806 - int32(1)
	v810 = F_IsBufferCleanupOK(m, l4)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L17
	} else {
		goto L156
	}
L137:
	;
	v727 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[5]))
	if v727 <= int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v730 != 0 {
		goto L136
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v732 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v679)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)) = uint16(v732)
	v734 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v709)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)) = uint16(v734)
	F_XLogBeginInsert(m)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L17
	} else {
		goto L143
	}
L141:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v731 != 0 {
		goto L136
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	F_XLogRegisterData(m, v29+int32(12), int32(4))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L17
	} else {
		goto L144
	}
L144:
	;
	F_XLogRegisterBuffer(m, int32(0), l4, int32(8))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L17
	} else {
		goto L145
	}
L145:
	;
	F_XLogRegisterBuffer(m, int32(1), l5, int32(8))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L17
	} else {
		goto L146
	}
L146:
	;
	v753 = F_XLogInsert(m, int32(12), int32(96))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L17
	} else {
		goto L147
	}
L147:
	;
	if l4 < int32(0) {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v773 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v772))) = base.I64_rotr(v753, v773)
	if l5 < int32(0) {
		goto L153
	} else {
		goto L154
	}
L149:
	;
	v758 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v758+(l4^int32(-1))<<(uint(int32(2))%32))))
	v772 = v764
	goto L148
L150:
	;
	goto L151
L151:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v772 = v766 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L148
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v797)+4)) = base.I32_wrap_i64(v753)
	*(*int32)(unsafe.Add(mBase, uint32(v797))) = base.I32_wrap_i64(int64(base.Ui64(v753) >> (uint(v773) % 64)))
	goto L136
L153:
	;
	v783 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v783+(l5^int32(-1))<<(uint(int32(2))%32))))
	v797 = v789
	goto L152
L154:
	;
	goto L155
L155:
	;
	v791 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v797 = v791 + l5<<(uint(int32(13))%32) + int32(-8192)
	goto L152
L156:
	;
	F_LockBuffer(m, l5, int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L17
	} else {
		goto L157
	}
L157:
	;
	if v810 != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	m.G0 = v29 + int32(2464)
	return
L159:
	;
	if l4 < int32(0) {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	goto L161
L161:
	;
	F_LockBuffer(m, l4, int32(0))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L17
	} else {
		goto L167
	}
L162:
	;
	v834 = int32(0)
	F_hashbucketcleanup(m, l0, l2, l4, v833, v834, l7, l8, l9, v834, v834, int32(1), v834, v834)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L17
	} else {
		goto L166
	}
L163:
	;
	v818 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[2]))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v818+(l4^int32(-1))<<(uint(int32(6))%32))+16))
	v833 = v824
	goto L162
L164:
	;
	goto L165
L165:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[3]))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v826+l4<<(uint(int32(6))%32)+int32(-64))+16))
	v833 = v832
	goto L162
L166:
	;
	goto L158
L167:
	;
	goto L158
}
func F_hash_aclitem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	return v3 + v4 + v6
}
func F_hash_aclitem_extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v9 = v5 + v6 + v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	if v11 == int64(0) {
		v15 = F_Int64GetDatum(m, base.I64_extend_i32_u(v9))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			return v15
		}
	} else {
		if v11 == int64(0) {
			v26 = int32(-1636608428)
			v65 = v26
			v66 = v26
			v69 = int32(0)
		} else {
			v29 = base.I32_wrap_i64(v11)
			v31 = v29 + int32(1021750440)
			v36 = base.I32_wrap_i64(int64(base.Ui64(v11)>>(uint(int64(32))%64))) ^ int32(-415931063)
			v42 = v29 - v36 - int32(1636608428) ^ base.I32_rotl(v36, int32(6))
			v46 = v31 - v42 ^ base.I32_rotl(v42, int32(8))
			v47 = v36 + v31
			v48 = v42 + v47
			v49 = v46 + v48
			v53 = v47 - v46 ^ base.I32_rotl(v46, int32(16))
			v57 = v48 - v53 ^ base.I32_rotl(v53, int32(19))
			v62 = v53 + v49
			v63 = v57 + v62
			v65 = v63
			v66 = v62
			v69 = v49 - v57 ^ base.I32_rotl(v57, int32(4)) ^ v63
		}
		v70 = int32(14)
		v72 = v69 - base.I32_rotl(v65, v70)
		v77 = v72 ^ (v9 + v66) - base.I32_rotl(v72, int32(11))
		v81 = v65 ^ v77 - base.I32_rotl(v77, int32(25))
		v85 = v81 ^ v72 - base.I32_rotl(v81, int32(16))
		v89 = v85 ^ v77 - base.I32_rotl(v85, int32(4))
		v93 = v89 ^ v81 - base.I32_rotl(v89, v70)
		v103 = F_Int64GetDatum(m, base.I64_extend_i32_u(v93)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v93^v85-base.I32_rotl(v93, int32(24))))
		mBase = m.M
		v104 = m.ExcPending
		if v104 != 0 {
			return int32(0)
		} else {
			return v103
		}
	}
}
func F_hash_agg_set_limits(m *base.Module, l0 float64, l1 float64, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v14 float64
	_ = v14
	var v16 int32
	_ = v16
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v24 float64
	_ = v24
	var v25 int32
	_ = v25
	var v26 float64
	_ = v26
	var v37 float64
	_ = v37
	var v41 float64
	_ = v41
	var v43 int32
	_ = v43
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v51 float64
	_ = v51
	var v53 float64
	_ = v53
	var v59 float64
	_ = v59
	var v65 float64
	_ = v65
	var v67 float64
	_ = v67
	var v70 float64
	_ = v70
	var v73 float64
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v108 float64
	_ = v108
	var v113 int64
	_ = v113
	v14 = *(*float64)(unsafe.Add(mBase, _c_F_hash_agg_set_limits[0]))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_hash_agg_set_limits[1]))
	v20 = base.F64_mul(base.F64_mul(v14, base.F64_convert_i32_s(v16)), float64(1024))
	v21 = float64(4.294967295e+09)
	if base.F64_lt(v20, v21) != 0 {
		v24 = v20
	} else {
		v24 = v21
	}
	v25 = base.I32_trunc_sat_f64_u(v24)
	v26 = base.F64_convert_i32_u(v25)
	if base.F64_ge(v26, base.F64_mul(l0, l1)) != 0 {
		if l5 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v25
		*(*int64)(unsafe.Add(mBase, uint32(l4))) = base.I64_trunc_sat_f64_u(base.F64_div(v26, l0))
		return
	} else {
		v37 = float64(1024)
		v41 = *(*float64)(unsafe.Add(mBase, _c_F_hash_agg_set_limits[0]))
		v43 = *(*int32)(unsafe.Add(mBase, _c_F_hash_agg_set_limits[1]))
		v47 = base.F64_mul(base.F64_mul(v41, base.F64_convert_i32_s(v43)), v37)
		v48 = float64(4.294967295e+09)
		if base.F64_lt(v47, v48) != 0 {
			v51 = v47
		} else {
			v51 = v48
		}
		v53 = base.F64_convert_i32_u(base.I32_trunc_sat_f64_u(v51))
		v59 = base.F64_mul(base.F64_add(base.F64_mul(v53, float64(0.25)), float64(-8192)), float64(0.0001220703125))
		v65 = base.F64_add(base.F64_div(base.F64_mul(l0, base.F64_mul(l1, float64(1.5))), v53), float64(1))
		if base.F64_gt(v65, v59) != 0 {
			v67 = v59
		} else {
			v67 = v65
		}
		if base.F64_lt(v67, float64(4)) != 0 {
			v70 = float64(4)
		} else {
			v70 = v67
		}
		if base.F64_gt(v70, float64(1024)) != 0 {
			v73 = v37
		} else {
			v73 = v70
		}
		v74 = base.I32_trunc_sat_f64_s(v73)
		v76 = int32(1073741823)
		if v76 <= v74 {
			v79 = v76
		} else {
			v79 = v74
		}
		if base.Ui32(int32(2)) <= base.Ui32(v79) {
			v87 = int32(32) - base.I32_clz(v79-int32(1))
		} else {
			v87 = int32(0)
		}
		if int32(31) < l2+v87 {
			v91 = int32(32) - l2
		} else {
			v91 = v87
		}
		if l5 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(1) << (uint(v91) % 32)
		} else {
		}
		v98 = int32(_a_F_hash_agg_set_limits_0)<<(uint(v91)%32) - int32(-8192)
		if base.Ui32(v98<<(uint(int32(2))%32)) < base.Ui32(v25) {
			v106 = v25 - v98
		} else {
			v106 = base.I32_trunc_sat_f64_u(base.F64_mul(v26, float64(0.75)))
		}
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v106
		v108 = base.F64_convert_i32_u(v106)
		if base.F64_gt(v108, l0) != 0 {
			v113 = base.I64_trunc_sat_f64_u(base.F64_div(v108, l0))
		} else {
			v113 = int64(1)
		}
		*(*int64)(unsafe.Add(mBase, uint32(l4))) = v113
		return
	}
}
func F_hash_agg_update_metrics(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
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
	var v62 int32
	_ = v62
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
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v118 int64
	_ = v118
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v5&int32(-2) != int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	goto L5
L3:
	;
	v39 = l2 << (uint(int32(13)) % 32)
	if l1 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	goto L3
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v17 == int32(0) {
		v36 = v14
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v22 = v14
	v23 = v17
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v25 = v24 + v22
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	if v26 != 0 {
		v22 = v25
		v23 = v26
		goto L7
	} else {
		goto L9
	}
L8:
	;
	v36 = v25
	goto L4
L9:
	;
	v28 = v23
	goto L10
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	if v31 != 0 {
		v22 = v25
		v23 = v31
		goto L7
	} else {
		goto L12
	}
L11:
	;
	goto L8
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v32 != v10 {
		v28 = v32
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v42 = v39 - int32(-8192)
	goto L16
L15:
	;
	v42 = v39
	goto L16
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	goto L19
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	goto L30
L18:
	;
	goto L17
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	if v51 == int32(0) {
		v70 = v48
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v56 = v48
	v57 = v51
	goto L21
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v59 = v58 + v56
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	if v60 != 0 {
		v56 = v59
		v57 = v60
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v70 = v59
	goto L18
L23:
	;
	v62 = v57
	goto L24
L24:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+28))
	if v65 != 0 {
		v56 = v59
		v57 = v65
		goto L21
	} else {
		goto L26
	}
L25:
	;
	goto L22
L26:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	if v66 != v44 {
		v62 = v66
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v102 = v36 + v42 + v70 + v100
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	if base.Ui32(v103) < base.Ui32(v102) {
		goto L39
	} else {
		goto L40
	}
L29:
	;
	goto L28
L30:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	if v81 == int32(0) {
		v100 = v78
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v86 = v78
	v87 = v81
	goto L32
L32:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v89 = v88 + v86
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+20))
	if v90 != 0 {
		v86 = v89
		v87 = v90
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v100 = v89
	goto L29
L34:
	;
	v92 = v87
	goto L35
L35:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+28))
	if v95 != 0 {
		v86 = v89
		v87 = v95
		goto L32
	} else {
		goto L37
	}
L36:
	;
	goto L33
L37:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
	if v96 != v74 {
		v92 = v96
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v102
	goto L41
L40:
	;
	goto L41
L41:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v106 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v118 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	if v118 == int64(0) {
		goto L1
	} else {
		goto L46
	}
L43:
	;
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v106)+24))
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v106)+32))
	goto L44
L44:
	;
	v113 = (v109 - v110) << (uint(int64(3)) % 64)
	v114 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	if base.Ui64(v113) <= base.Ui64(v114) {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v113
	goto L42
L46:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+304)) = base.F64_add(base.F64_div(base.F64_convert_i32_u(v100), base.F64_convert_i64_u(v118)), float64(12))
	goto L1
}
func F_hash_array(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
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
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
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
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
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
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_DatumGetAnyArrayP(m, v15)
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v22 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v25 = int32(28)
	goto L5
L4:
	;
	v25 = int32(4)
	goto L5
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v16+v25)))
	if v22 == int32(-1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36+v16)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	v35 = v30
	v36 = int32(40)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v35 = v16 + int32(16)
	v36 = int32(12)
	goto L6
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L53
	}
L11:
	;
	v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(v81)+11)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+10)))
	v85 = int32(*(*int16)(unsafe.Add(mBase, uint32(v81)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v81 + int32(132)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v92 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+54)) = uint16(v92)
	v94 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+52)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v91
	v97 = F_ArrayGetNItemsSafe(m, v27, v35)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L24
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v41 == v38 {
		v81 = v40
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v44 = F_lookup_type_cache(m, v38, int32(128))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+136))
	v50 = base.B2i32(v38 != int32(2249))
	if base.B2i32(v46 == int32(0))&v50 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	if v38 != int32(2249) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v76
	v81 = v76
	goto L11
L19:
	;
	v76 = v44
	goto L18
L20:
	;
	goto L21
L21:
	;
	v52 = int32(_a_F_hash_array_0)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_hash_array[0]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_hash_array[0])) = v56
	v59 = F_palloc0(m, int32(328))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = int32(2249)
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v59)+8)) = uint16(v63)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+10)) = uint8(v65)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+11)))
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+11)) = uint8(v67)
	F_fmgr_info(m, int32(_a_F_hash_array_1), v59+int32(132))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hash_array[0])) = v53
	v76 = v59
	goto L18
L24:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v99 == int32(-1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v159 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v158
	if int32(0) < v97 {
		goto L39
	} else {
		goto L40
	}
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
	if v102 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(0)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v135 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
	v105 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v104
	v158 = v105
	goto L25
L30:
	;
	goto L31
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	if v112 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v111 + (v115<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v158 = int32(0)
	goto L25
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v112 + v111
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v158 = v111 + v127<<(uint(int32(3))%32) + int32(16)
	goto L25
L35:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v16 + (v138<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v158 = int32(0)
	goto L25
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v135 + v16
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v158 = v16 + v150<<(uint(int32(3))%32) + int32(16)
	goto L25
L38:
	;
	m.G0 = v13 - int32(-64)
	return v208
L39:
	;
	v171 = int32(0)
	v173 = v159
	goto L42
L40:
	;
	v208 = v159
	v213 = v99
	goto L41
L41:
	;
	if v213 == int32(-1) {
		goto L38
	} else {
		goto L50
	}
L42:
	;
	v182 = F_array_iter_next(m, v11+int32(-48), v11+int32(-49), v171, v85, v84&int32(1), v83)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v208 = v198
	v213 = v202
	goto L41
L44:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v184 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v195 = int32(0)
	goto L47
L46:
	;
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v186)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v182
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	v193 = m.T0[v192].(func(*base.Module, int32) int32)(m, v11+int32(-28))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v198 = v195 + v173*int32(31)
	v200 = v171 + int32(1)
	if v200 != v97 {
		v171 = v200
		v173 = v198
		goto L42
	} else {
		goto L49
	}
L48:
	;
	v195 = v193
	goto L47
L49:
	;
	goto L43
L50:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v16 == v216 {
		goto L38
	} else {
		goto L51
	}
L51:
	;
	F_pfree(m, v16)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L38
L53:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v231 = F_format_type_be(m, v38)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v231
	F_errmsg(m, int32(_a_F_hash_array_2), v13)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_hash_array_3), int32(_a_F_hash_array_4), int32(_a_F_hash_array_5))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hash_array_extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v144 int32
	_ = v144
	var v152 int64
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v181 int64
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v196 int64
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_DatumGetAnyArrayP(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v24 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = int32(28)
	goto L5
L4:
	;
	v27 = int32(4)
	goto L5
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18+v27)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	if v24 == int32(-1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40+v18)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	if v44 != 0 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v39 = v34
	v40 = int32(40)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v39 = v18 + int32(16)
	v40 = int32(12)
	goto L6
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L50
	}
L11:
	;
	v56 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55)+11)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+10)))
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v55 + int32(160)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v65 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+46)) = uint16(v65)
	v67 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+44)) = uint8(v67)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v64
	v70 = F_ArrayGetNItemsSafe(m, v29, v39)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v45 == v42 {
		v55 = v44
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v48 = F_lookup_type_cache(m, v42, int32(_a_F_hash_array_extended_0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+164))
	if v50 == int32(0) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v48
	v55 = v48
	goto L11
L18:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v72 == int32(-1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(1)
	if v70 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	if v75 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = int64(0)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v108 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v78 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v77
	v131 = v78
	goto L19
L24:
	;
	goto L25
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = int64(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v18)+68))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	if v85 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v84 + (v88<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v131 = int32(0)
	goto L19
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v84 + v85
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v131 = v84 + v100<<(uint(int32(3))%32) + int32(16)
	goto L19
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v18 + (v111<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v131 = int32(0)
	goto L19
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v18 + v108
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v131 = v18 + v123<<(uint(int32(3))%32) + int32(16)
	goto L19
L32:
	;
	if v188 == int32(-1) {
		goto L45
	} else {
		goto L46
	}
L33:
	;
	v188 = v72
	v196 = int64(1)
	goto L32
L34:
	;
	goto L35
L35:
	;
	v144 = int32(0)
	v152 = int64(1)
	goto L36
L36:
	;
	v158 = F_array_iter_next(m, v13+int32(-56), v13+int32(-57), v144, v58, v57&int32(1), v56)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v188 = v185
	v196 = v181
	goto L32
L38:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+7)))
	if v160 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v178 = int64(0)
	goto L41
L40:
	;
	v162 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+52)) = uint8(v162)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v158
	v165 = F_Int64GetDatum(m, v31)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	v181 = v178 + v152*int64(31)
	v183 = v144 + int32(1)
	if v183 != v70 {
		v144 = v183
		v152 = v181
		goto L36
	} else {
		goto L44
	}
L42:
	;
	v167 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+60)) = uint8(v167)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v165
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v174 = m.T0[v173].(func(*base.Module, int32) int32)(m, v13+int32(-36))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v174)))
	v178 = v176
	goto L41
L44:
	;
	goto L37
L45:
	;
	v204 = F_Int64GetDatum(m, v196)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L49
	}
L46:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v18 == v200 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	F_pfree(m, v18)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	m.G0 = v15 - int32(-64)
	return v204
L50:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v217 = F_format_type_be(m, v42)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v217
	F_errmsg(m, int32(_a_F_hash_array_extended_1), v15)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_hash_array_extended_2), int32(_a_F_hash_array_extended_3), int32(_a_F_hash_array_extended_4))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
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
func F_hash_bytes_uint32(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	v6 = int32(711645284)
	v9 = l0 - int32(1636608428) ^ v6 - int32(1455628627)
	v14 = v9 ^ int32(-1636608428) - base.I32_rotl(v9, int32(25))
	v19 = v14 ^ v6 - base.I32_rotl(v14, int32(16))
	v23 = v19 ^ v9 - base.I32_rotl(v19, int32(4))
	v27 = v23 ^ v14 - base.I32_rotl(v23, int32(14))
	return v27 ^ v19 - base.I32_rotl(v27, int32(24))
}
func F_hash_bytes_uint32_extended(m *base.Module, l0 int32, l1 int64) int64 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	if l1 == int64(0) {
		v9 = int32(-1636608428)
		v48 = v9
		v49 = v9
		v52 = int32(0)
	} else {
		v12 = base.I32_wrap_i64(l1)
		v14 = v12 + int32(1021750440)
		v19 = base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v25 = v12 - v19 - int32(1636608428) ^ base.I32_rotl(v19, int32(6))
		v29 = v14 - v25 ^ base.I32_rotl(v25, int32(8))
		v30 = v19 + v14
		v31 = v25 + v30
		v32 = v29 + v31
		v36 = v30 - v29 ^ base.I32_rotl(v29, int32(16))
		v40 = v31 - v36 ^ base.I32_rotl(v36, int32(19))
		v45 = v36 + v32
		v46 = v40 + v45
		v48 = v46
		v49 = v45
		v52 = v32 - v40 ^ base.I32_rotl(v40, int32(4)) ^ v46
	}
	v53 = int32(14)
	v55 = v52 - base.I32_rotl(v48, v53)
	v60 = v55 ^ (l0 + v49) - base.I32_rotl(v55, int32(11))
	v64 = v48 ^ v60 - base.I32_rotl(v60, int32(25))
	v68 = v64 ^ v55 - base.I32_rotl(v64, int32(16))
	v72 = v68 ^ v60 - base.I32_rotl(v68, int32(4))
	v76 = v72 ^ v64 - base.I32_rotl(v72, v53)
	return base.I64_extend_i32_u(v76)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v76^v68-base.I32_rotl(v76, int32(24)))
}
func F_hash_record_extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v144 int64
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
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
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int64
	_ = v208
	var v210 int64
	_ = v210
	var v219 int64
	_ = v219
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
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
	var v241 int32
	_ = v241
	var v256 int64
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	v15 = int64(0)
	v17 = m.G0
	v19 = v17 + int32(-64)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
	F_check_stack_depth(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v32 = F_lookup_rowtype_tupdesc(m, v30, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v22
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v37
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+52)) = uint16(v37)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = int32(base.Ui32(v35) >> (uint(int32(2)) % 32))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	if v47 == v37 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v68 != v30 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v58 = F_MemoryContextAlloc(m, v53, v34<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v50 < v34 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v67 = v47
	v68 = v52
	goto L5
L9:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v58
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v63)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v34
	v67 = v63
	v68 = int32(0)
	goto L5
L10:
	;
	v119 = F_palloc(m, v113)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L24
	}
L11:
	;
	v75 = v67 + int32(20)
	v79 = v34 << (uint(int32(2)) % 32)
	if v75&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v79)) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	if v70 != v31 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v113 = v34 << (uint(int32(2)) % 32)
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+8)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v30
	v113 = v79
	goto L10
L15:
	;
	if v79 == int32(0) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v79 == int32(0) {
		goto L14
	} else {
		goto L23
	}
L18:
	;
	v89 = v67 + v79 + int32(20)
	v91 = v67 + int32(24)
	if base.Ui32(v91) < base.Ui32(v89) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v93 = v89
	goto L21
L20:
	;
	v93 = v91
	goto L21
L21:
	;
	v100 = (v93-v67-int32(21))&int32(-4) + int32(4)
	if v100 == int32(0) {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	base.MemoryFill(m, v75, int32(0), v100)
	goto L14
L23:
	;
	base.MemoryFill(m, v75, int32(0), v79)
	goto L14
L24:
	;
	v121 = F_palloc(m, v34)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_heap_deform_tuple(m, v17+int32(-20), v32, v119, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v34 <= int32(0) {
		v256 = v15
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_pfree(m, v119)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L54
	}
L28:
	;
	v131 = int32(0)
	v144 = v15
	goto L29
L29:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v152 = v32 + v146<<(uint(int32(4))%32) + v131*int32(100)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+111)))
	if v153 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L49
	}
L31:
	;
	goto L30
L32:
	;
	v157 = v152 + int32(20)
	v159 = v131 << (uint(int32(2)) % 32)
	v160 = v67 + int32(20) + v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if v161 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v219 = v144
	goto L34
L34:
	;
	v221 = v131 + int32(1)
	if v34 != v221 {
		v131 = v221
		v144 = v219
		goto L29
	} else {
		goto L48
	}
L35:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v121))))
	if v179 != 0 {
		goto L43
	} else {
		goto L44
	}
L36:
	;
	v170 = F_lookup_type_cache(m, v168, int32(_a_F_hash_record_extended_0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L41
	}
L37:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v157)+68))
	v168 = v164
	goto L36
L38:
	;
	goto L39
L39:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v157)+68))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v165 == v166 {
		v176 = v161
		goto L35
	} else {
		goto L40
	}
L40:
	;
	v168 = v165
	goto L36
L41:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v170)+164))
	if v172 == int32(0) {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v170
	v176 = v170
	goto L35
L43:
	;
	v210 = int64(0)
	goto L45
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v176 + int32(160)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v157)+96))
	v187 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)) = uint8(v187)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v186
	v190 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+26)) = uint16(v190)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v119+v159)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+32)) = uint8(v187)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v193
	v197 = F_Int64GetDatum(m, v27)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	v219 = v210 + v144*int64(31)
	goto L34
L46:
	;
	v199 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+32)) = uint8(v199)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v197
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v206 = m.T0[v205].(func(*base.Module, int32) int32)(m, v17+int32(-56))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v206)))
	v210 = v208
	goto L45
L48:
	;
	v256 = v219
	goto L27
L49:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v231 = F_format_type_be(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v231
	F_errmsg(m, int32(_a_F_hash_record_extended_1), v19)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_hash_record_extended_2), int32(2015), int32(_a_F_hash_record_extended_3))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_pfree(m, v121)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if int32(0) <= v262 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_DecrTupleDescRefCount(m, v32)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v267 != v22 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	F_pfree(m, v22)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v271 = F_Int64GetDatum(m, v256)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	m.G0 = v19 - int32(-64)
	return v271
}
