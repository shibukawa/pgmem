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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = (v7 - int32(1)) & l1
	if base.Ui32(int32(2)) <= base.Ui32(v6) {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v20 = (v6 - int32(1)) & base.I32_rotr(l1, v17)
	} else {
		v20 = int32(0)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v20
	return
}
func F_ExecHashTableInsert(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v17 = F_ExecFetchSlotMinimalTuple(m, l1, v13+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if base.Ui32(int32(2)) <= base.Ui32(v19) {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v27 = (v19 - int32(1)) & base.I32_rotr(l2, v24)
		} else {
			v27 = int32(0)
		}
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		if v28 == v27 {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v31 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
			v32 = *(*float64)(unsafe.Add(mBase, uint32(l0)+64))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v35 = v33 + int32(8)
			v36 = F_dense_alloc(m, l0, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = l2
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				if v41 != 0 {
					v42 = F__emscripten_memcpy_bulkmem(m, v36+int32(8), v17, v41)
					mBase = m.M
				} else {
				}
				v45 = v36 + int32(18)
				v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45))))
				v48 = v46 & int32(32767)
				*(*uint16)(unsafe.Add(mBase, uint32(v45))) = uint16(v48)
				v50 = int32(1)
				v54 = (v30 - v50) & l2 << (uint(int32(2)) % 32)
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v54+v55)))
				*(*int32)(unsafe.Add(mBase, uint32(v36))) = v57
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v59+v54))) = v36
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v62 != v50 {
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if int32(1073741823) < v65 {
					} else {
						if base.F64_lt(base.F64_convert_i32_s(v65), base.F64_sub(v32, v31)) == int32(0) {
						} else {
							v74 = v65 << (uint(int32(1)) % 32)
							if base.Ui32(int32(268435455)) < base.Ui32(v74) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v74
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v78 + int32(1)
							}
						}
					}
				}
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				v84 = v83 + v35
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v84
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				if base.Ui32(v86) < base.Ui32(v84) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v84
				} else {
				}
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if base.Ui32(v90<<(uint(int32(2))%32)+v84) <= base.Ui32(v89) {
					v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
					if v110 == int32(1) {
						F_pfree(m, v17)
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return
						} else {
							m.G0 = v13 + int32(16)
							return
						}
					} else {
						m.G0 = v13 + int32(16)
						return
					}
				} else {
					F_ExecHashIncreaseNumBatches(m, l0)
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return
					} else {
						v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
						if v110 == int32(1) {
							F_pfree(m, v17)
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return
							} else {
								m.G0 = v13 + int32(16)
								return
							}
						} else {
							m.G0 = v13 + int32(16)
							return
						}
					}
				}
			}
		} else {
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
			F_ExecHashJoinSaveTuple(m, v17, l2, v97+v27<<(uint(int32(2))%32), l0)
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return
			} else {
				v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
				if v110 == int32(1) {
					F_pfree(m, v17)
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return
					} else {
						m.G0 = v13 + int32(16)
						return
					}
				} else {
					m.G0 = v13 + int32(16)
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
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	v3 = int32(0)
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v8) < base.Ui32(int32(25)) {
		v66 = v3
	} else {
		v14 = int32(base.Ui32(v8+int32(262120)) >> (uint(int32(2)) % 32))
		if v14&int32(65535) == int32(0) {
			v66 = v3
		} else {
			v23 = v3
			v24 = v14
			for {
				v28 = int32(65535)
				v33 = int32(1)
				v36 = int32(base.Ui32(v23&v28+v24&v28+v33) >> (uint(v33) % 32))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v36<<(uint(int32(2))%32)+(l0+int32(24))-int32(4))))
				v47 = l0 + v44&int32(32767)
				v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+6)))
				if int32(0) <= v50 {
					v53 = int32(8)
				} else {
					v53 = int32(16)
				}
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v47+v53)))
				v56 = base.B2i32(base.Ui32(l1) < base.Ui32(v55))
				if base.Ui32(l1) < base.Ui32(v55) {
					v57 = v36 - v33
				} else {
					v57 = v24
				}
				if base.Ui32(l1) < base.Ui32(v55) {
					v60 = v23
				} else {
					v60 = v36
				}
				if base.Ui32(v60&int32(65535)) < base.Ui32(v57&int32(65535)) {
					v23 = v60
					v24 = v57
					continue
				} else {
					break
				}
				break
			}
			v66 = v60
		}
	}
	return v66 & int32(65535)
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
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67<<(uint(int32(2))%32)+v30)+48))
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
	v88 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88+(v76^int32(-1))<<(uint(int32(2))%32))))
	v102 = v94
	goto L27
L29:
	;
	goto L30
L30:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
	F_errmsg_internal(m, int32(492350), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(475953), int32(75), int32(321435))
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 float64
	_ = v86
	var v97 int32
	_ = v97
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
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int64
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v358 int32
	_ = v358
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v401 int32
	_ = v401
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
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v562 int64
	_ = v562
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v582 int64
	_ = v582
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	v20 = m.G0
	v22 = v20 - int32(48)
	m.G0 = v22
	v24 = F_RelationGetNumberOfBlocksInFork(m, l0, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L2
	} else {
		goto L149
	}
L2:
	;
	return int32(0)
L3:
	;
	if v24 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+118)))
	if v31 != int32(112) {
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
	v624 = m.ExcPending
	if v624 != 0 {
		goto L2
	} else {
		goto L146
	}
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v46 = base.B2i32(l2 == int32(3))
	goto L7
L9:
	;
	v34 = int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if int32(0) < v36 {
		v46 = v34
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v39 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v40 == int32(0) {
		v46 = v34
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v52 = base.I32_div_s(v48<<(uint(int32(13))%32), int32(2000))
	v54 = v52
	goto L15
L14:
	;
	v54 = int32(307)
	goto L15
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)))
	v67 = int32(4)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v57+v59*int32(0)<<(uint(int32(2))%32)+v67-v67)))
	goto L16
L16:
	;
	v73 = F__hash_getnewbuf(m, l0, int32(0), l2)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v75 = int32(10)
	if v54 <= v75 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v78 = v75
	goto L20
L19:
	;
	v78 = v54
	goto L20
L20:
	;
	v80 = v78 & int32(65535)
	v86 = base.F64_div(l1, base.F64_convert_i32_u(v80))
	if base.F64_le(v86, float64(2)) != 0 {
		v102 = int32(2)
		goto L22
	} else {
		goto L23
	}
L21:
	;
	F_MarkBufferDirty(m, v73)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L2
	} else {
		goto L47
	}
L22:
	;
	v103 = F__hash_spareindex(m, v102)
	mBase = m.M
	if v73 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	if base.F64_ge(v86, float64(1.073741824e+09)) != 0 {
		v102 = int32(1073741824)
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if base.F64_lt(v86, float64(4.294967296e+09))&base.F64_ge(v86, float64(0)) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v100 = F__hash_spareindex(m, v99)
	mBase = m.M
	v101 = F__hash_get_totalbuckets(m, v100)
	mBase = m.M
	v102 = v101
	goto L22
L26:
	;
	v97 = base.I32_trunc_f64_u(v86)
	v99 = v97
	goto L25
L27:
	;
	goto L28
L28:
	;
	v99 = int32(0)
	goto L25
L29:
	;
	goto L34
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v107+(v73^int32(-1))<<(uint(int32(2))%32))))
	v121 = v113
	goto L29
L31:
	;
	goto L32
L32:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v121 = v115 + v73<<(uint(int32(13))%32) + int32(-8192)
	goto L29
L34:
	;
	goto L35
L35:
	;
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+16)))
	v126 = v121 + v125
	*(*int64)(unsafe.Add(mBase, uint32(v126)+8)) = int64(-36028758364258305)
	*(*int64)(unsafe.Add(mBase, uint32(v126))) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+68)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v121)+32)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v121)+24)) = int64(17284990528)
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+40)) = uint16(v80)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+72)) = v71
	v139 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+48)) = v102 - v139
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+19)))
	v146 = v142<<(uint(int32(8))%32) - int32(40)
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+42)) = uint16(v146)
	v148 = int32(-1)
	v151 = v102 + v139
	if v151&v102 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v158 = v148<<(uint(int32(32)-base.I32_clz(v151))%32) ^ v148
	goto L38
L37:
	;
	v158 = v102
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+52)) = v158
	v160 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+56)) = int32(base.Ui32(v158) >> (uint(v160) % 32))
	v167 = base.I32_clz(v146&int32(65496)) ^ int32(31)
	v168 = int32(3)
	v169 = v167 + v168
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+46)) = uint16(v169)
	v172 = v160 << (uint(v167) % 32)
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+44)) = uint16(v172)
	v175 = v121 + int32(76)
	if v175&v168 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v203 = int32(0)
	v205 = F___memset(m, v121+int32(468), v203, int32(4096))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v175+v103<<(uint(int32(2))%32)))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v121-int32(-64)))) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v121)+60)) = v103
	v216 = int32(4568)
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+12)) = uint16(v216)
	goto L21
L40:
	;
	v181 = v121 + int32(468)
	if base.Ui32(v181) <= base.Ui32(v175) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v198 = F___memset(m, v175, int32(0), int32(392))
	mBase = m.M
	goto L39
L43:
	;
	v185 = v121 + int32(80)
	if base.Ui32(v185) < base.Ui32(v181) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v187 = v181
	goto L46
L45:
	;
	v187 = v185
	goto L46
L46:
	;
	v195 = F___memset(m, v175, int32(0), (v187-v121-int32(77))&int32(-4)+int32(4))
	mBase = m.M
	goto L39
L47:
	;
	if v73 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v46 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v223+(v73^int32(-1))<<(uint(int32(2))%32))))
	v237 = v229
	goto L48
L50:
	;
	goto L51
L51:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v237 = v231 + v73<<(uint(int32(13))%32) + int32(-8192)
	goto L48
L52:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v22)+32)) = l1
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v237)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v239
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v237)+40)))
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+44)) = uint16(v241)
	F_XLogBeginInsert(m)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L2
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v237)+48))
	F_LockBuffer(m, v73, int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L2
	} else {
		goto L63
	}
L55:
	;
	F_XLogRegisterData(m, v22+int32(32), int32(14))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	F_XLogRegisterBuffer(m, int32(0), v73, int32(14))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v256 = F_XLogInsert(m, int32(12), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	if v73 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v275))) = base.I64_rotr(v256, int64(32))
	goto L54
L60:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v261+(v73^int32(-1))<<(uint(int32(2))%32))))
	v275 = v267
	goto L59
L61:
	;
	goto L62
L62:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v275 = v269 + v73<<(uint(int32(13))%32) + int32(-8192)
	goto L59
L63:
	;
	v285 = v280 + int32(1)
	if v285 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_LockBuffer(m, v73, int32(2))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L2
	} else {
		goto L115
	}
L65:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v289 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L2
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v293 = F__hash_getnewbuf(m, l0, int32(1), l2)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L2
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v237)+48))
	if int32(0) <= v293 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v313)+16)))
	v315 = v314 + v313
	*(*int32)(unsafe.Add(mBase, uint32(v315)+12)) = int32(-8388606)
	*(*int64)(unsafe.Add(mBase, uint32(v315)+4)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v315))) = v295
	F_MarkBufferDirty(m, v293)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L2
	} else {
		goto L75
	}
L72:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v313 = v299 + v293<<(uint(int32(13))%32) + int32(-8192)
	goto L71
L73:
	;
	goto L74
L74:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v306+(v293^int32(-1))<<(uint(int32(2))%32))))
	v313 = v312
	goto L71
L75:
	;
	if v46 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if int32(0) <= v293 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	goto L78
L78:
	;
	F_UnlockReleaseBuffer(m, v293)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L2
	} else {
		goto L84
	}
L79:
	;
	F_log_newpage(m, l0, l2, int32(1), v341, int32(1))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L2
	} else {
		goto L83
	}
L80:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v341 = v327 + v293<<(uint(int32(13))%32) + int32(-8192)
	goto L79
L81:
	;
	goto L82
L82:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v334+(v293^int32(-1))<<(uint(int32(2))%32))))
	v341 = v340
	goto L79
L83:
	;
	goto L78
L84:
	;
	if v280 == int32(0) {
		goto L64
	} else {
		goto L85
	}
L85:
	;
	v358 = int32(1)
	goto L86
L86:
	;
	v374 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v374 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L64
L88:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L2
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v377 = int32(1)
	v378 = v358 + v377
	v382 = v378 - v377
	if base.Ui32(int32(2)) <= base.Ui32(v378) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L90
L92:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v401<<(uint(int32(2))%32)+(v237+int32(76))-int32(4))))
	v408 = v407 + v378
	v409 = F__hash_getnewbuf(m, l0, v408, l2)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L2
	} else {
		goto L99
	}
L93:
	;
	v388 = int32(32) - base.I32_clz(v382)
	goto L95
L94:
	;
	v388 = int32(0)
	goto L95
L95:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v388) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v391 = int32(3)
	v401 = int32(base.Ui32(v382)>>(uint(v388-v391)%32))&v391 | v388<<(uint(int32(2))%32) - int32(30)
	goto L98
L97:
	;
	v401 = v388
	goto L98
L98:
	;
	goto L92
L99:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v237+int32(48))))
	v412 = int32(0)
	v413 = base.B2i32(v412 <= v409)
	if v413 == v412 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v432 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v431)+16)))
	v433 = v432 + v431
	*(*int32)(unsafe.Add(mBase, uint32(v433)+12)) = int32(-8388606)
	*(*int32)(unsafe.Add(mBase, uint32(v433)+8)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v433)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v433))) = v411
	F_MarkBufferDirty(m, v409)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L2
	} else {
		goto L104
	}
L101:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v417+(v409^int32(-1))<<(uint(int32(2))%32))))
	v431 = v423
	goto L100
L102:
	;
	goto L103
L103:
	;
	v425 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v431 = v425 + v409<<(uint(int32(13))%32) + int32(-8192)
	goto L100
L104:
	;
	if v46 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	if v413 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	goto L107
L107:
	;
	F_UnlockReleaseBuffer(m, v409)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L2
	} else {
		goto L113
	}
L108:
	;
	F_log_newpage(m, l0, l2, v408, v459, int32(1))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L2
	} else {
		goto L112
	}
L109:
	;
	v445 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v445+(v409^int32(-1))<<(uint(int32(2))%32))))
	v459 = v451
	goto L108
L110:
	;
	goto L111
L111:
	;
	v453 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v459 = v453 + v409<<(uint(int32(13))%32) + int32(-8192)
	goto L108
L112:
	;
	goto L107
L113:
	;
	if v358 != v280 {
		v358 = v378
		goto L86
	} else {
		goto L114
	}
L114:
	;
	goto L87
L115:
	;
	v489 = v280 + int32(2)
	v490 = F__hash_getnewbuf(m, l0, v489, l2)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L2
	} else {
		goto L116
	}
L116:
	;
	v492 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v237)+44)))
	if v490 < int32(0) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	F_MarkBufferDirty(m, v490)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L2
	} else {
		goto L125
	}
L118:
	;
	goto L123
L119:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v497+(v490^int32(-1))<<(uint(int32(2))%32))))
	v511 = v503
	goto L118
L120:
	;
	goto L121
L121:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v511 = v505 + v490<<(uint(int32(13))%32) + int32(-8192)
	goto L118
L123:
	;
	goto L124
L124:
	;
	v513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v511)+16)))
	v514 = v511 + v513
	*(*int64)(unsafe.Add(mBase, uint32(v514)+8)) = int64(-36028775544127489)
	*(*int64)(unsafe.Add(mBase, uint32(v514))) = int64(-1)
	v519 = int32(24)
	v522 = F___memset(m, v511+v519, int32(255), v492)
	mBase = m.M
	v524 = v492 + v519
	*(*uint16)(unsafe.Add(mBase, uint32(v511)+12)) = uint16(v524)
	goto L117
L125:
	;
	v529 = v237 + int32(68)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)))
	if base.Ui32(int32(1024)) <= base.Ui32(v530) {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237+v530<<(uint(int32(2))%32))+468)) = v489
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v529)))
	*(*int32)(unsafe.Add(mBase, uint32(v529))) = v537 + int32(1)
	F_MarkBufferDirty(m, v73)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L2
	} else {
		goto L127
	}
L127:
	;
	if v46 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v237)+44)))
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+32)) = uint16(v543)
	F_XLogBeginInsert(m)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L2
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	F_UnlockReleaseBuffer(m, v490)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L2
	} else {
		goto L144
	}
L131:
	;
	F_XLogRegisterData(m, v22+int32(32), int32(2))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L2
	} else {
		goto L132
	}
L132:
	;
	F_XLogRegisterBuffer(m, int32(0), v490, int32(6))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L2
	} else {
		goto L133
	}
L133:
	;
	F_XLogRegisterBuffer(m, int32(1), v73, int32(8))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L2
	} else {
		goto L134
	}
L134:
	;
	v562 = F_XLogInsert(m, int32(12), int32(16))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L2
	} else {
		goto L135
	}
L135:
	;
	if v490 < int32(0) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v582 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v581))) = base.I64_rotr(v562, v582)
	if v73 < int32(0) {
		goto L141
	} else {
		goto L142
	}
L137:
	;
	v567 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v567+(v490^int32(-1))<<(uint(int32(2))%32))))
	v581 = v573
	goto L136
L138:
	;
	goto L139
L139:
	;
	v575 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v581 = v575 + v490<<(uint(int32(13))%32) + int32(-8192)
	goto L136
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v606)+4)) = base.I32_wrap_i64(v562)
	*(*int32)(unsafe.Add(mBase, uint32(v606))) = base.I32_wrap_i64(int64(base.Ui64(v562) >> (uint(v582) % 64)))
	goto L130
L141:
	;
	v592 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v592+(v73^int32(-1))<<(uint(int32(2))%32))))
	v606 = v598
	goto L140
L142:
	;
	goto L143
L143:
	;
	v600 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v606 = v600 + v73<<(uint(int32(13))%32) + int32(-8192)
	goto L140
L144:
	;
	F_UnlockReleaseBuffer(m, v73)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L2
	} else {
		goto L145
	}
L145:
	;
	m.G0 = v22 + int32(48)
	return v285
L146:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v625 + int32(4)
	F_errmsg_internal(m, int32(652756), v22+int32(16))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L2
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(475953), int32(345), int32(94020))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L2
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L2
	} else {
		goto L150
	}
L150:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v646 + int32(4)
	F_errmsg(m, int32(652800), v22)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L2
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(475953), int32(455), int32(94020))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L2
	} else {
		goto L152
	}
L152:
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
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	v5 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if l3 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v202
L2:
	;
	v121 = l2
	v125 = int32(408)
	goto L32
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
	v32 = int32(base.Ui32(v24+int32(262120)) >> (uint(int32(2)) % 32))
	goto L11
L10:
	;
	v32 = int32(0)
	goto L11
L11:
	;
	v34 = v32 & int32(65535)
	if base.Ui32(v34) < base.Ui32(l2) {
		v202 = v5
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v42 = l2
	v46 = v5
	goto L13
L13:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)))
	v56 = v42
	goto L15
L14:
	;
	v202 = v113
	goto L1
L15:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v56&int32(65535)<<(uint(int32(2))%32)+(l1+int32(24))-int32(4))))
	v75 = l1 + v72&int32(32767)
	if v51&int32(1) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v75)+6)))
	if int32(0) <= v97 {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	goto L16
L18:
	;
	v90 = v56 + int32(1)
	if base.Ui32(v90&int32(65535)) <= base.Ui32(v34) {
		v56 = v90
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	if v82 != int32(1) {
		goto L17
	} else {
		goto L23
	}
L20:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)))
	if v78 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+7)))
	if v79&int32(32) != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v85 = int32(98304)
	if v72&v85 != v85 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	v202 = v46
	goto L1
L26:
	;
	if v94 != v102 {
		v202 = v46
		goto L1
	} else {
		goto L30
	}
L27:
	;
	v100 = int32(8)
	goto L29
L28:
	;
	v100 = int32(16)
	goto L29
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v75+v100)))
	goto L26
L30:
	;
	v106 = v12 + int32(52) + v46<<(uint(int32(3))%32)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v107
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v106)+6)) = uint16(v56)
	*(*uint16)(unsafe.Add(mBase, uint32(v106)+4)) = uint16(v109)
	v112 = int32(1)
	v113 = v46 + v112
	v115 = v56 + v112
	if base.Ui32(v115&int32(65535)) <= base.Ui32(v34) {
		v42 = v115
		v46 = v113
		goto L13
	} else {
		goto L31
	}
L31:
	;
	goto L14
L32:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)))
	v135 = v121
	goto L34
L33:
	;
	v202 = v183
	goto L1
L34:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v135&int32(65535)<<(uint(int32(2))%32)+(l1+int32(24))-int32(4))))
	v154 = l1 + v151&int32(32767)
	if v130&int32(1) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v175 = int32(*(*int16)(unsafe.Add(mBase, uint32(v154)+6)))
	if int32(0) <= v175 {
		goto L46
	} else {
		goto L47
	}
L36:
	;
	goto L35
L37:
	;
	v169 = v135 - int32(1)
	if v169&int32(65535) != 0 {
		v135 = v169
		goto L34
	} else {
		goto L44
	}
L38:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	if v161 != int32(1) {
		goto L36
	} else {
		goto L42
	}
L39:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)))
	if v157 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+7)))
	if v158&int32(32) != 0 {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	goto L38
L42:
	;
	v164 = int32(98304)
	if v151&v164 != v164 {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	goto L37
L44:
	;
	v202 = v125
	goto L1
L45:
	;
	if v172 != v180 {
		v202 = v125
		goto L1
	} else {
		goto L49
	}
L46:
	;
	v178 = int32(8)
	goto L48
L47:
	;
	v178 = int32(16)
	goto L48
L48:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v154+v178)))
	goto L45
L49:
	;
	v182 = int32(1)
	v183 = v125 - v182
	v186 = v12 + int32(52) + v183<<(uint(int32(3))%32)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v187
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v186)+6)) = uint16(v135)
	*(*uint16)(unsafe.Add(mBase, uint32(v186)+4)) = uint16(v189)
	v193 = v135 - v182
	if v193&int32(65535) != 0 {
		v121 = v193
		v125 = v183
		goto L32
	} else {
		goto L50
	}
L50:
	;
	goto L33
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int64
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v317 int32
	_ = v317
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v530 int64
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v579 int32
	_ = v579
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v759 int64
	_ = v759
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v778 int32
	_ = v778
	var v779 int64
	_ = v779
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
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
	v34 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34+(l4^int32(-1))<<(uint(int32(2))%32))))
	v48 = v40
	goto L1
L3:
	;
	goto L4
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
	v53 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53+(l5^int32(-1))<<(uint(int32(2))%32))))
	v67 = v59
	goto L5
L7:
	;
	goto L8
L8:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
	v72 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72+(l4^int32(-1))<<(uint(int32(6))%32))+16))
	v87 = v78
	goto L9
L11:
	;
	goto L12
L12:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _consts[10]))
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
	v91 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91+(l5^int32(-1))<<(uint(int32(6))%32))+16))
	v106 = v97
	goto L13
L15:
	;
	goto L16
L16:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99+l5<<(uint(int32(6))%32)+int32(-64))+16))
	v106 = v105
	goto L13
L17:
	;
	return
L18:
	;
	v120 = l5
	v123 = v11
	v124 = l4
	v126 = v48
	v127 = v11
	v128 = v67
	v129 = v49 + v48
	goto L20
L19:
	;
	v681 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v680)+16)))
	F_LockBuffer(m, l5, int32(2))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L17
	} else {
		goto L129
	}
L20:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+12)))
	if base.Ui32(v135) < base.Ui32(int32(25)) {
		v472 = v120
		v475 = v123
		v479 = v127
		v480 = v128
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v674 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v680 = v674 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L19
L22:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if l4 == v124 {
		goto L82
	} else {
		goto L83
	}
L23:
	;
	v139 = v135 + int32(262120)
	if v139&int32(262140) == int32(0) {
		v472 = v120
		v475 = v123
		v479 = v127
		v480 = v128
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v161 = int32(1)
	v162 = v120
	v165 = v123
	v169 = v127
	v170 = v128
	goto L25
L25:
	;
	v177 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+12)) = uint8(v177)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v161<<(uint(int32(2))%32)+(v126+int32(24))-int32(4))))
	v185 = int32(98304)
	if v184&v185 == v185 {
		v443 = v162
		v446 = v165
		v450 = v169
		v451 = v170
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v472 = v443
	v475 = v446
	v479 = v450
	v480 = v451
	goto L22
L27:
	;
	if v161 != int32(base.Ui32(v139)>>(uint(int32(2))%32))&int32(65535) {
		v161 = v161 + int32(1)
		v162 = v443
		v165 = v446
		v169 = v450
		v170 = v451
		goto L25
	} else {
		goto L80
	}
L28:
	;
	v191 = v126 + v184&int32(32767)
	if l6 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v195 = F_hash_search(m, l6, v191, int32(0), v29+int32(12))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L17
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v202 = int32(*(*int16)(unsafe.Add(mBase, uint32(v191)+6)))
	if int32(0) <= v202 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+12)))
	if v197&int32(1) != 0 {
		v443 = v162
		v446 = v165
		v450 = v169
		v451 = v170
		goto L27
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v209 = v207 & l8
	if base.Ui32(v209) <= base.Ui32(l7) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v205 = int32(8)
	goto L37
L36:
	;
	v205 = int32(16)
	goto L37
L37:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v191+v205)))
	goto L34
L38:
	;
	if v211&v209 != l3 {
		v443 = v162
		v446 = v165
		v450 = v169
		v451 = v170
		goto L27
	} else {
		goto L42
	}
L39:
	;
	v211 = int32(-1)
	goto L41
L40:
	;
	v211 = l9
	goto L41
L41:
	;
	goto L38
L42:
	;
	v214 = F_CopyIndexTuple(m, v191)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L17
	} else {
		goto L43
	}
L43:
	;
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214)+6)))
	v218 = v216 | int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v214)+6)) = uint16(v218)
	v221 = v165 & int32(65535)
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+14)))
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+12)))
	v226 = v224 - v225
	v228 = (v221 + int32(1)) << (uint(int32(2)) % 32)
	if v228 <= v226 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v238 = (v216&int32(8191) + int32(7)) & int32(16376)
	if base.Ui32(v232) < base.Ui32(v169+v238) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v232 = v226 - v228
	goto L47
L46:
	;
	v232 = int32(0)
	goto L47
L47:
	;
	goto L44
L48:
	;
	v241 = int32(4437236)
	v243 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v243 + int32(1)
	F__hash_pgaddmultitup(m, l0, v162, v29+int32(16), v29+int32(1648), v221)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L17
	} else {
		goto L51
	}
L49:
	;
	v406 = v162
	v409 = v165
	v413 = v169
	v414 = v170
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29+int32(16)+v409&int32(65535)<<(uint(int32(2))%32)))) = v214
	v443 = v406
	v446 = v409 + int32(1)
	v450 = v413 + v238
	v451 = v414
	goto L27
L51:
	;
	F_MarkBufferDirty(m, v162)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+118)))
	if v256 != int32(112) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v297 = int32(0)
	v298 = int32(4437236)
	v300 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v300 - int32(1)
	F_LockBuffer(m, v162, v297)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L17
	} else {
		goto L67
	}
L54:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v260 <= int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v263 != 0 {
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
	v266 = m.ExcPending
	if v266 != 0 {
		goto L17
	} else {
		goto L60
	}
L58:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v264 != 0 {
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
	v270 = m.ExcPending
	if v270 != 0 {
		goto L17
	} else {
		goto L61
	}
L61:
	;
	v273 = F_XLogInsert(m, int32(12), int32(80))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v292))) = base.I64_rotr(v273, int64(32))
	goto L53
L64:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v278+(v162^int32(-1))<<(uint(int32(2))%32))))
	v292 = v284
	goto L63
L65:
	;
	goto L66
L66:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v292 = v286 + v162<<(uint(int32(13))%32) + int32(-8192)
	goto L63
L67:
	;
	if v221 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v317 = v297
	goto L71
L69:
	;
	goto L70
L70:
	;
	v371 = F__hash_addovflpage(m, l0, l1, v162, base.B2i32(l5 == v162))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L17
	} else {
		goto L76
	}
L71:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(16)+v317<<(uint(int32(2))%32))))
	F_pfree(m, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L17
	} else {
		goto L73
	}
L72:
	;
	goto L70
L73:
	;
	v342 = v317 + int32(1)
	if v342 != v221 {
		v317 = v342
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v406 = v371
	v409 = int32(0)
	v413 = v393
	v414 = v392
	goto L50
L76:
	;
	if v371 < int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v376+(v371^int32(-1))<<(uint(int32(2))%32))))
	v392 = v382
	v393 = int32(0)
	goto L75
L78:
	;
	goto L79
L79:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v392 = v385 + v371<<(uint(int32(13))%32) + int32(-8192)
	v393 = int32(0)
	goto L75
L80:
	;
	goto L26
L81:
	;
	if v487 == int32(-1) {
		goto L88
	} else {
		goto L89
	}
L82:
	;
	F_LockBuffer(m, l4, int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L17
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	F_UnlockReleaseBuffer(m, v124)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
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
	v496 = int32(4437236)
	v498 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v498 + int32(1)
	F__hash_pgaddmultitup(m, l0, v472, v29+int32(16), v29+int32(1648), v475&int32(65535))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L17
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v645 = F_ReadBuffer(m, l0, v487)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L17
	} else {
		goto L122
	}
L91:
	;
	F_MarkBufferDirty(m, v472)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L17
	} else {
		goto L92
	}
L92:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+118)))
	if v513 != int32(112) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v554 = int32(4437236)
	v556 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v556 - int32(1)
	if l5 == v472 {
		goto L108
	} else {
		goto L109
	}
L94:
	;
	v517 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v517 <= int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v520 != 0 {
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
	v523 = m.ExcPending
	if v523 != 0 {
		goto L17
	} else {
		goto L100
	}
L98:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v521 != 0 {
		goto L93
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	F_XLogRegisterBuffer(m, int32(0), v472, int32(9))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L17
	} else {
		goto L101
	}
L101:
	;
	v530 = F_XLogInsert(m, int32(12), int32(80))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L17
	} else {
		goto L102
	}
L102:
	;
	if v472 < int32(0) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v549))) = base.I64_rotr(v530, int64(32))
	goto L93
L104:
	;
	v535 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v535+(v472^int32(-1))<<(uint(int32(2))%32))))
	v549 = v541
	goto L103
L105:
	;
	goto L106
L106:
	;
	v543 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v549 = v543 + v472<<(uint(int32(13))%32) + int32(-8192)
	goto L103
L107:
	;
	v567 = v475 & int32(65535)
	if v567 != 0 {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	F_LockBuffer(m, l5, int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L17
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	F_UnlockReleaseBuffer(m, v472)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
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
	v579 = int32(0)
	goto L116
L114:
	;
	goto L115
L115:
	;
	F_LockBuffer(m, l4, int32(2))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L17
	} else {
		goto L120
	}
L116:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(16)+v579<<(uint(int32(2))%32))))
	F_pfree(m, v600)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L17
	} else {
		goto L118
	}
L117:
	;
	goto L115
L118:
	;
	v604 = v579 + int32(1)
	if v604 != v567 {
		v579 = v604
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
	v638 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v638+(l4^int32(-1))<<(uint(int32(2))%32))))
	v680 = v644
	goto L19
L122:
	;
	F_LockBuffer(m, v645, int32(1))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L17
	} else {
		goto L123
	}
L123:
	;
	F__hash_checkpage(m, l0, v645, int32(1))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L17
	} else {
		goto L124
	}
L124:
	;
	if v645 < int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v671 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v670)+16)))
	v120 = v472
	v123 = v475
	v124 = v645
	v126 = v670
	v127 = v479
	v128 = v480
	v129 = v671 + v670
	goto L20
L126:
	;
	v656 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v656+(v645^int32(-1))<<(uint(int32(2))%32))))
	v670 = v662
	goto L125
L127:
	;
	goto L128
L128:
	;
	v664 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v670 = v664 + v645<<(uint(int32(13))%32) + int32(-8192)
	goto L125
L129:
	;
	v685 = v680 + v681
	if l5 < int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v704 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v703)+16)))
	v705 = int32(4437236)
	v707 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v707 + int32(1)
	v711 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v685)+12)))
	v713 = v711 & int32(65503)
	*(*uint16)(unsafe.Add(mBase, uint32(v685)+12)) = uint16(v713)
	v715 = v703 + v704
	v716 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v715)+12)))
	v718 = v716 & int32(65519)
	*(*uint16)(unsafe.Add(mBase, uint32(v715)+12)) = uint16(v718)
	v720 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v685)+12)))
	v722 = v720 | int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v685)+12)) = uint16(v722)
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L17
	} else {
		goto L134
	}
L131:
	;
	v689 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v689+(l5^int32(-1))<<(uint(int32(2))%32))))
	v703 = v695
	goto L130
L132:
	;
	goto L133
L133:
	;
	v697 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v703 = v697 + l5<<(uint(int32(13))%32) + int32(-8192)
	goto L130
L134:
	;
	F_MarkBufferDirty(m, l5)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L17
	} else {
		goto L135
	}
L135:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728)+118)))
	if v729 != int32(112) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v810 = int32(4437236)
	v812 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v812 - int32(1)
	v816 = F_IsBufferCleanupOK(m, l4)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L17
	} else {
		goto L156
	}
L137:
	;
	v733 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v733 <= int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v736 != 0 {
		goto L136
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v738 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v685)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)) = uint16(v738)
	v740 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v715)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)) = uint16(v740)
	F_XLogBeginInsert(m)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L17
	} else {
		goto L143
	}
L141:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v737 != 0 {
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
	v748 = m.ExcPending
	if v748 != 0 {
		goto L17
	} else {
		goto L144
	}
L144:
	;
	F_XLogRegisterBuffer(m, int32(0), l4, int32(8))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L17
	} else {
		goto L145
	}
L145:
	;
	F_XLogRegisterBuffer(m, int32(1), l5, int32(8))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L17
	} else {
		goto L146
	}
L146:
	;
	v759 = F_XLogInsert(m, int32(12), int32(96))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
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
	v779 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v778))) = base.I64_rotr(v759, v779)
	if l5 < int32(0) {
		goto L153
	} else {
		goto L154
	}
L149:
	;
	v764 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v764+(l4^int32(-1))<<(uint(int32(2))%32))))
	v778 = v770
	goto L148
L150:
	;
	goto L151
L151:
	;
	v772 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v778 = v772 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L148
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v803)+4)) = base.I32_wrap_i64(v759)
	*(*int32)(unsafe.Add(mBase, uint32(v803))) = base.I32_wrap_i64(int64(base.Ui64(v759) >> (uint(v779) % 64)))
	goto L136
L153:
	;
	v789 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v789+(l5^int32(-1))<<(uint(int32(2))%32))))
	v803 = v795
	goto L152
L154:
	;
	goto L155
L155:
	;
	v797 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v803 = v797 + l5<<(uint(int32(13))%32) + int32(-8192)
	goto L152
L156:
	;
	F_LockBuffer(m, l5, int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L17
	} else {
		goto L157
	}
L157:
	;
	if v816 != 0 {
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
	v850 = m.ExcPending
	if v850 != 0 {
		goto L17
	} else {
		goto L167
	}
L162:
	;
	v840 = int32(0)
	F_hashbucketcleanup(m, l0, l2, l4, v839, v840, l7, l8, l9, v840, v840, int32(1), v840, v840)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L17
	} else {
		goto L166
	}
L163:
	;
	v824 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v824+(l4^int32(-1))<<(uint(int32(6))%32))+16))
	v839 = v830
	goto L162
L164:
	;
	goto L165
L165:
	;
	v832 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v832+l4<<(uint(int32(6))%32)+int32(-64))+16))
	v839 = v838
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
			v27 = int32(-1636608428)
			v65 = v27
			v67 = v27
			v70 = v27
		} else {
			v30 = base.I32_wrap_i64(v11)
			v35 = base.I32_wrap_i64(int64(base.Ui64(v11)>>(uint(int64(32))%64))) ^ int32(-415931063)
			v41 = v30 - v35 - int32(1636608428) ^ base.I32_rotl(v35, int32(6))
			v43 = v30 + int32(1021750440)
			v44 = v35 + v43
			v45 = v41 + v44
			v49 = v43 - v41 ^ base.I32_rotl(v41, int32(8))
			v53 = v44 - v49 ^ base.I32_rotl(v49, int32(16))
			v57 = v45 - v53 ^ base.I32_rotl(v53, int32(19))
			v58 = v49 + v45
			v59 = v53 + v58
			v65 = v57 + v59
			v67 = v59
			v70 = v58 - v57 ^ base.I32_rotl(v57, int32(4))
		}
		v72 = int32(14)
		v74 = v65 ^ v70 - base.I32_rotl(v65, v72)
		v79 = v74 ^ (v9 + v67) - base.I32_rotl(v74, int32(11))
		v83 = v79 ^ v65 - base.I32_rotl(v79, int32(25))
		v87 = v83 ^ v74 - base.I32_rotl(v83, int32(16))
		v91 = v87 ^ v79 - base.I32_rotl(v87, int32(4))
		v95 = v91 ^ v83 - base.I32_rotl(v91, v72)
		v105 = F_Int64GetDatum(m, base.I64_extend_i32_u(v95)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v95^v87-base.I32_rotl(v95, int32(24))))
		mBase = m.M
		v106 = m.ExcPending
		if v106 != 0 {
			return int32(0)
		} else {
			return v105
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 float64
	_ = v33
	var v39 float64
	_ = v39
	var v47 int64
	_ = v47
	var v51 float64
	_ = v51
	var v55 float64
	_ = v55
	var v57 int32
	_ = v57
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v65 float64
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 float64
	_ = v74
	var v80 float64
	_ = v80
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v91 float64
	_ = v91
	var v94 float64
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v130 float64
	_ = v130
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 float64
	_ = v142
	var v146 float64
	_ = v146
	var v154 int64
	_ = v154
	var v171 int64
	_ = v171
	v14 = *(*float64)(unsafe.Add(mBase, _consts[345]))
	v16 = *(*int32)(unsafe.Add(mBase, _consts[326]))
	v20 = base.F64_mul(base.F64_mul(v14, base.F64_convert_i32_s(v16)), float64(1024))
	v21 = float64(4.294967295e+09)
	if base.F64_lt(v20, v21) != 0 {
		v24 = v20
	} else {
		v24 = v21
	}
	if base.F64_lt(v24, float64(4.294967296e+09))&base.F64_ge(v24, float64(0)) != 0 {
		v30 = base.I32_trunc_f64_u(v24)
		v32 = v30
	} else {
		v32 = int32(0)
	}
	v33 = base.F64_convert_i32_u(v32)
	if base.F64_ge(v33, base.F64_mul(l0, l1)) != 0 {
		if l5 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v32
		v39 = base.F64_div(v33, l0)
		if base.F64_lt(v39, float64(1.8446744073709552e+19))&base.F64_ge(v39, float64(0)) == int32(0) {
			v171 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(l4))) = v171
			return
		} else {
			v47 = base.I64_trunc_f64_u(v39)
			*(*int64)(unsafe.Add(mBase, uint32(l4))) = v47
			return
		}
	} else {
		v51 = float64(1024)
		v55 = *(*float64)(unsafe.Add(mBase, _consts[345]))
		v57 = *(*int32)(unsafe.Add(mBase, _consts[326]))
		v61 = base.F64_mul(base.F64_mul(v55, base.F64_convert_i32_s(v57)), v51)
		v62 = float64(4.294967295e+09)
		if base.F64_lt(v61, v62) != 0 {
			v65 = v61
		} else {
			v65 = v62
		}
		if base.F64_lt(v65, float64(4.294967296e+09))&base.F64_ge(v65, float64(0)) != 0 {
			v71 = base.I32_trunc_f64_u(v65)
			v73 = v71
		} else {
			v73 = int32(0)
		}
		v74 = base.F64_convert_i32_u(v73)
		v80 = base.F64_mul(base.F64_add(base.F64_mul(v74, float64(0.25)), float64(-8192)), float64(0.0001220703125))
		v86 = base.F64_add(base.F64_div(base.F64_mul(l0, base.F64_mul(l1, float64(1.5))), v74), float64(1))
		if base.F64_gt(v86, v80) != 0 {
			v88 = v80
		} else {
			v88 = v86
		}
		if base.F64_lt(v88, float64(4)) != 0 {
			v91 = float64(4)
		} else {
			v91 = v88
		}
		if base.F64_gt(v91, float64(1024)) != 0 {
			v94 = v51
		} else {
			v94 = v91
		}
		if base.F64_lt(base.F64_abs(v94), float64(2.147483648e+09)) != 0 {
			v98 = base.I32_trunc_f64_s(v94)
			v100 = v98
		} else {
			v100 = int32(-2147483648)
		}
		v102 = int32(1073741823)
		if v102 <= v100 {
			v105 = v102
		} else {
			v105 = v100
		}
		if base.Ui32(int32(2)) <= base.Ui32(v105) {
			v113 = int32(32) - base.I32_clz(v105-int32(1))
		} else {
			v113 = int32(0)
		}
		if int32(31) < l2+v113 {
			v117 = int32(32) - l2
		} else {
			v117 = v113
		}
		if l5 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(1) << (uint(v117) % 32)
		} else {
		}
		v124 = int32(8192)<<(uint(v117)%32) - int32(-8192)
		v130 = base.F64_mul(v33, float64(0.75))
		if base.F64_lt(v130, float64(4.294967296e+09))&base.F64_ge(v130, float64(0)) != 0 {
			v136 = base.I32_trunc_f64_u(v130)
			v138 = v136
		} else {
			v138 = int32(0)
		}
		if base.Ui32(v124<<(uint(int32(2))%32)) < base.Ui32(v32) {
			v139 = v32 - v124
		} else {
			v139 = v138
		}
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v139
		v142 = base.F64_convert_i32_u(v139)
		if base.F64_lt(l0, v142) == int32(0) {
			v171 = int64(1)
			*(*int64)(unsafe.Add(mBase, uint32(l4))) = v171
			return
		} else {
			v146 = base.F64_div(v142, l0)
			if base.F64_lt(v146, float64(1.8446744073709552e+19))&base.F64_ge(v146, float64(0)) == int32(0) {
				v171 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(l4))) = v171
				return
			} else {
				v154 = base.I64_trunc_f64_u(v146)
				*(*int64)(unsafe.Add(mBase, uint32(l4))) = v154
				return
			}
		}
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v32 int32
	_ = v32
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v106 int64
	_ = v106
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v35 = l2 << (uint(int32(13)) % 32)
	if l1 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v18 = v14
	v19 = v15
	goto L7
L5:
	;
	v32 = v14
	goto L6
L6:
	;
	goto L3
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v21 = v20 + v18
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v22 != 0 {
		v18 = v21
		v19 = v22
		goto L7
	} else {
		goto L9
	}
L8:
	;
	v32 = v21
	goto L6
L9:
	;
	v24 = v19
	goto L10
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v27 != 0 {
		v18 = v21
		v19 = v27
		goto L7
	} else {
		goto L12
	}
L11:
	;
	goto L8
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if v28 != v10 {
		v24 = v28
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v38 = v35 - int32(-8192)
	goto L16
L15:
	;
	v38 = v35
	goto L16
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	if v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	if v71 != 0 {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	v48 = v44
	v49 = v45
	goto L21
L19:
	;
	v62 = v44
	goto L20
L20:
	;
	goto L17
L21:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v51 = v50 + v48
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	if v52 != 0 {
		v48 = v51
		v49 = v52
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v62 = v51
	goto L20
L23:
	;
	v54 = v49
	goto L24
L24:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
	if v57 != 0 {
		v48 = v51
		v49 = v57
		goto L21
	} else {
		goto L26
	}
L25:
	;
	goto L22
L26:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if v58 != v40 {
		v54 = v58
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v90 = v32 + v38 + v62 + v88
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	if base.Ui32(v91) < base.Ui32(v90) {
		goto L39
	} else {
		goto L40
	}
L29:
	;
	v74 = v70
	v75 = v71
	goto L32
L30:
	;
	v88 = v70
	goto L31
L31:
	;
	goto L28
L32:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	v77 = v76 + v74
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	if v78 != 0 {
		v74 = v77
		v75 = v78
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v88 = v77
	goto L31
L34:
	;
	v80 = v75
	goto L35
L35:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
	if v83 != 0 {
		v74 = v77
		v75 = v83
		goto L32
	} else {
		goto L37
	}
L36:
	;
	goto L33
L37:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	if v84 != v66 {
		v80 = v84
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v90
	goto L41
L40:
	;
	goto L41
L41:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v94 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v106 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	if v106 == int64(0) {
		goto L1
	} else {
		goto L46
	}
L43:
	;
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v94)+24))
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v94)+32))
	goto L44
L44:
	;
	v101 = (v97 - v98) << (uint(int64(3)) % 64)
	v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	if base.Ui64(v101) <= base.Ui64(v102) {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v101
	goto L42
L46:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+304)) = base.F64_add(base.F64_div(base.F64_convert_i32_u(v88), base.F64_convert_i64_u(v106)), float64(12))
	goto L1
}
func F_hash_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_DatumGetAnyArrayP(m, v16)
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
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v23 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = int32(28)
	goto L5
L4:
	;
	v26 = int32(4)
	goto L5
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v17+v26)))
	if v23 == int32(-1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37+v17)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v36 = v31
	v37 = int32(40)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v36 = v17 + int32(16)
	v37 = int32(12)
	goto L6
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L53
	}
L11:
	;
	v86 = int32(*(*int8)(unsafe.Add(mBase, uint32(v83)+11)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+10)))
	v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v83 + int32(132)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v95 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+54)) = uint16(v95)
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+52)) = uint8(v97)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v94
	v100 = F_ArrayGetNItems(m, v28, v36)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L24
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v42 == v39 {
		v83 = v41
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v45 = F_lookup_type_cache(m, v39, int32(128))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+136))
	if base.B2i32(v47 == int32(0))&base.B2i32(v39 != int32(2249)) != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	if v39 != int32(2249) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = v79
	v83 = v79
	goto L11
L19:
	;
	v79 = v45
	goto L18
L20:
	;
	goto L21
L21:
	;
	v55 = int32(4442576)
	v56 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v59
	v62 = F_palloc0(m, int32(328))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = int32(2249)
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v62)+8)) = uint16(v66)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+10)) = uint8(v68)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+11)))
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+11)) = uint8(v70)
	F_fmgr_info(m, int32(6192), v62+int32(132))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v56
	v79 = v62
	goto L18
L24:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v102 == int32(-1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v162 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v161
	if int32(0) < v100 {
		goto L39
	} else {
		goto L40
	}
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	if v105 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = int64(0)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v138 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	v108 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v107
	v161 = v108
	goto L25
L30:
	;
	goto L31
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = int64(0)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	if v115 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v114 + (v118<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v161 = int32(0)
	goto L25
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v114 + v115
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v161 = v114 + v130<<(uint(int32(3))%32) + int32(16)
	goto L25
L35:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v17 + (v141<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v161 = int32(0)
	goto L25
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v138 + v17
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v161 = v17 + v153<<(uint(int32(3))%32) + int32(16)
	goto L25
L38:
	;
	m.G0 = v14 - int32(-64)
	return v212
L39:
	;
	v173 = int32(0)
	v176 = v162
	goto L42
L40:
	;
	v212 = v162
	v218 = v102
	goto L41
L41:
	;
	if v218 == int32(-1) {
		goto L38
	} else {
		goto L50
	}
L42:
	;
	v186 = F_array_iter_next(m, v12+int32(-48), v12+int32(-49), v173, v88, v87&int32(1), v86)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v212 = v202
	v218 = v206
	goto L41
L44:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
	if v189 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v199 = int32(0)
	goto L47
L46:
	;
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+60)) = uint8(v190)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v186
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v197 = m.T0[v196].(func(*base.Module, int32) int32)(m, v12+int32(-28))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v202 = v199 + v176*int32(31)
	v204 = v173 + int32(1)
	if v204 != v100 {
		v173 = v204
		v176 = v202
		goto L42
	} else {
		goto L49
	}
L48:
	;
	v199 = v197
	goto L47
L49:
	;
	goto L43
L50:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v17 == v221 {
		goto L38
	} else {
		goto L51
	}
L51:
	;
	F_pfree(m, v17)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
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
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v236 = F_format_type_be(m, v39)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v236
	F_errmsg(m, int32(179494), v14)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(471433), int32(4196), int32(22683))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
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
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v147 int32
	_ = v147
	var v153 int64
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int64
	_ = v177
	var v179 int64
	_ = v179
	var v182 int64
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v197 int64
	_ = v197
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_DatumGetAnyArrayP(m, v18)
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
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v27 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v30 = int32(28)
	goto L5
L4:
	;
	v30 = int32(4)
	goto L5
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v19+v30)))
	if v27 == int32(-1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41+v19)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	if v45 != 0 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v40 = v35
	v41 = int32(40)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v40 = v19 + int32(16)
	v41 = int32(12)
	goto L6
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L50
	}
L11:
	;
	v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v56)+11)))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+10)))
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v56)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v56 + int32(160)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v66 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+46)) = uint16(v66)
	v68 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+44)) = uint8(v68)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v65
	v71 = F_ArrayGetNItems(m, v32, v40)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L18
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v46 == v43 {
		v56 = v45
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v49 = F_lookup_type_cache(m, v43, int32(32768))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+164))
	if v51 == int32(0) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v49
	v56 = v49
	goto L11
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v73 == int32(-1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(1)
	if v71 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	if v76 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(0)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v109 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v79 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v78
	v132 = v79
	goto L19
L24:
	;
	goto L25
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if v86 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v85 + (v89<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v132 = int32(0)
	goto L19
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v85 + v86
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v132 = v85 + v101<<(uint(int32(3))%32) + int32(16)
	goto L19
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v19 + (v112<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v132 = int32(0)
	goto L19
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v109 + v19
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v132 = v19 + v124<<(uint(int32(3))%32) + int32(16)
	goto L19
L32:
	;
	if v191 == int32(-1) {
		goto L45
	} else {
		goto L46
	}
L33:
	;
	v191 = v73
	v197 = int64(1)
	goto L32
L34:
	;
	goto L35
L35:
	;
	v147 = int32(0)
	v153 = int64(1)
	goto L36
L36:
	;
	v160 = F_array_iter_next(m, v14+int32(-56), v14+int32(-57), v147, v59, v58&int32(1), v57)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v191 = v186
	v197 = v182
	goto L32
L38:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+7)))
	if v162 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v179 = int64(0)
	goto L41
L40:
	;
	v163 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+52)) = uint8(v163)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v160
	v166 = F_Int64GetDatum(m, v24)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	v182 = v179 + v153*int64(31)
	v184 = v147 + int32(1)
	if v184 != v71 {
		v147 = v184
		v153 = v182
		goto L36
	} else {
		goto L44
	}
L42:
	;
	v168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+60)) = uint8(v168)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v166
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v175 = m.T0[v174].(func(*base.Module, int32) int32)(m, v14+int32(-36))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v175)))
	v179 = v177
	goto L41
L44:
	;
	goto L37
L45:
	;
	v206 = F_Int64GetDatum(m, v197)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L49
	}
L46:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v19 == v202 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	F_pfree(m, v19)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	m.G0 = v16 - int32(-64)
	return v206
L50:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v219 = F_format_type_be(m, v43)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v219
	F_errmsg(m, int32(179437), v16)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(471433), int32(4324), int32(439337))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	if l1 == int64(0) {
		v10 = int32(-1636608428)
		v48 = v10
		v50 = v10
		v53 = v10
	} else {
		v13 = base.I32_wrap_i64(l1)
		v18 = base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v24 = v13 - v18 - int32(1636608428) ^ base.I32_rotl(v18, int32(6))
		v26 = v13 + int32(1021750440)
		v27 = v18 + v26
		v28 = v24 + v27
		v32 = v26 - v24 ^ base.I32_rotl(v24, int32(8))
		v36 = v27 - v32 ^ base.I32_rotl(v32, int32(16))
		v40 = v28 - v36 ^ base.I32_rotl(v36, int32(19))
		v41 = v32 + v28
		v42 = v36 + v41
		v48 = v40 + v42
		v50 = v42
		v53 = v41 - v40 ^ base.I32_rotl(v40, int32(4))
	}
	v55 = int32(14)
	v57 = v48 ^ v53 - base.I32_rotl(v48, v55)
	v62 = v57 ^ (l0 + v50) - base.I32_rotl(v57, int32(11))
	v66 = v62 ^ v48 - base.I32_rotl(v62, int32(25))
	v70 = v66 ^ v57 - base.I32_rotl(v66, int32(16))
	v74 = v70 ^ v62 - base.I32_rotl(v70, int32(4))
	v78 = v74 ^ v66 - base.I32_rotl(v74, v55)
	return base.I64_extend_i32_u(v78)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v78^v70-base.I32_rotl(v78, int32(24)))
}
func F_hash_record_extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v144 int64
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v217 int64
	_ = v217
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v255 int64
	_ = v255
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
	v16 = int64(0)
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	F_check_stack_depth(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v34 = F_lookup_rowtype_tupdesc(m, v32, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = v24
	v39 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v39
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+52)) = uint16(v39)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = int32(base.Ui32(v37) >> (uint(int32(2)) % 32))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	if v49 == v39 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v70 != v32 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v60 = F_MemoryContextAlloc(m, v55, v36<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v52 < v36 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v69 = v49
	v70 = v54
	goto L5
L9:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v60
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v36
	v69 = v65
	v70 = int32(0)
	goto L5
L10:
	;
	v116 = F_palloc(m, v110)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L24
	}
L11:
	;
	v77 = v36 << (uint(int32(2)) % 32)
	v79 = v69 + int32(20)
	if v79&int32(3) != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	if v72 != v33 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v110 = v36 << (uint(int32(2)) % 32)
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v32
	v110 = v77
	goto L10
L15:
	;
	v105 = F__emscripten_memset_bulkmem(m, v79, base.I32_extend8_s(int32(0)), v77)
	mBase = m.M
	goto L23
L16:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v77) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(v77+v79) <= base.Ui32(v79) {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v89 = v69 + v77 + int32(20)
	v91 = v69 + int32(24)
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
	v102 = F__emscripten_memset_bulkmem(m, v79, base.I32_extend8_s(int32(0)), (v93-v69-int32(21))&int32(-4)+int32(4))
	mBase = m.M
	goto L22
L22:
	;
	goto L14
L23:
	;
	goto L14
L24:
	;
	v118 = F_palloc(m, v36)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_heap_deform_tuple(m, v19+int32(-20), v34, v116, v118)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v36 <= int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_pfree(m, v116)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L56
	}
L28:
	;
	v255 = v16
	goto L27
L29:
	;
	goto L30
L30:
	;
	v124 = int32(20)
	v130 = int32(0)
	v144 = v16
	goto L31
L31:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v153 = v34 + v124 + v147<<(uint(int32(4))%32) + v130*int32(100)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+91)))
	if v154 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L51
	}
L33:
	;
	goto L32
L34:
	;
	v158 = v130 << (uint(int32(2)) % 32)
	v159 = v69 + v124 + v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v160 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v217 = v144
	goto L36
L36:
	;
	v219 = v130 + int32(1)
	if v36 != v219 {
		v130 = v219
		v144 = v217
		goto L31
	} else {
		goto L50
	}
L37:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+v118))))
	if v178 != 0 {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v169 = F_lookup_type_cache(m, v167, int32(32768))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L43
	}
L39:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v153)+68))
	v167 = v163
	goto L38
L40:
	;
	goto L41
L41:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v153)+68))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if v164 == v165 {
		v175 = v160
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v167 = v164
	goto L38
L43:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v169)+164))
	if v171 == int32(0) {
		goto L33
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v169
	v175 = v169
	goto L37
L45:
	;
	v208 = v16
	goto L47
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v175 + int32(160)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v153)+96))
	v185 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+24)) = uint8(v185)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v184
	v188 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+26)) = uint16(v188)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v158+v116)))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)) = uint8(v185)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v191
	v195 = F_Int64GetDatum(m, v29)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v217 = v208 + v144*int64(31)
	goto L36
L48:
	;
	v197 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)) = uint8(v197)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v195
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v204 = m.T0[v203].(func(*base.Module, int32) int32)(m, v19+int32(-56))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v204)))
	v208 = v206
	goto L47
L50:
	;
	v255 = v217
	goto L27
L51:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v229 = F_format_type_be(m, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v229
	F_errmsg(m, int32(179437), v21)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(471020), int32(2015), int32(439701))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
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
	F_pfree(m, v118)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	if int32(0) <= v262 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_DecrTupleDescRefCount(m, v34)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v267 != v24 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L60
L62:
	;
	F_pfree(m, v24)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v271 = F_Int64GetDatum(m, v255)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	m.G0 = v21 - int32(-64)
	return v271
}
