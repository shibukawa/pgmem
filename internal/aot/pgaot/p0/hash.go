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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = F_ExecFetchSlotMinimalTuple(m, l1, v10+int32(15))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if base.Ui32(int32(2)) <= base.Ui32(v16) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v24 = (v16 - int32(1)) & base.I32_rotr(l2, v21)
		} else {
			v24 = int32(0)
		}
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		if v25 == v24 {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			v33 = v31 + int32(8)
			v34 = F_dense_alloc(m, l0, v33)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = l2
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				if v37 != 0 {
					base.MemoryCopy(m, v34+int32(8), v14, v37)
				} else {
				}
				v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)))
				v43 = v41 & int32(_a_F_ExecHashTableInsert_0)
				*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v43)
				v46 = (v27 - int32(1)) & l2 << (uint(int32(2)) % 32)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v46+v47)))
				*(*int32)(unsafe.Add(mBase, uint32(v34))) = v49
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v51+v46))) = v34
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v54 != int32(1) {
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if int32(1073741823) < v57 {
					} else {
						v61 = *(*float64)(unsafe.Add(mBase, uint32(l0)+64))
						v62 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
						if base.F64_lt(base.F64_convert_i32_s(v57), base.F64_sub(v61, v62)) == int32(0) {
						} else {
							v68 = v57 << (uint(int32(1)) % 32)
							if base.Ui32(int32(268435455)) < base.Ui32(v68) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v68
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v72 + int32(1)
							}
						}
					}
				}
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				v78 = v77 + v33
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v78
				v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				if base.Ui32(v80) < base.Ui32(v78) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v78
				} else {
				}
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if base.Ui32(v84<<(uint(int32(2))%32)+v78) <= base.Ui32(v83) {
					v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
					if v101 == int32(1) {
						F_pfree(m, v14)
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return
						} else {
							m.G0 = v10 + int32(16)
							return
						}
					} else {
						m.G0 = v10 + int32(16)
						return
					}
				} else {
					F_ExecHashIncreaseNumBatches(m, l0)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						if v101 == int32(1) {
							F_pfree(m, v14)
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						} else {
							m.G0 = v10 + int32(16)
							return
						}
					}
				}
			}
		} else {
			v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
			F_ExecHashJoinSaveTuple(m, v14, l2, v91+v24<<(uint(int32(2))%32), l0)
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return
			} else {
				v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
				if v101 == int32(1) {
					F_pfree(m, v14)
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return
					} else {
						m.G0 = v10 + int32(16)
						return
					}
				} else {
					m.G0 = v10 + int32(16)
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
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v7 == int32(0) {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		v11 = int32(1)
		v13 = F_index_getprocinfo(m, l0, v11, v11)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v19 = F_FunctionCall1Coll(m, v13, v18, v10)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l3))) = v19 & int64(4294967295)
				v24 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v24)
				return v7 ^ int32(1)
			}
		}
	} else {
		return v7 ^ int32(1)
	}
}
func F__hash_getbucketbuf_from_hashkey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
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
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v5
	v18 = F__hash_getcachedmetap(m, l0, v11+int32(12), v5)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v181 != 0 {
		goto L49
	} else {
		goto L50
	}
L2:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v83)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	if base.Ui32(v86) <= base.Ui32(v87) {
		v177 = v64
		v178 = v18
		goto L1
	} else {
		goto L23
	}
L3:
	;
	return int32(0)
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v26 = l1 & v23
	if base.Ui32(v26) <= base.Ui32(v22) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v28 = int32(-1)
	goto L8
L7:
	;
	v28 = v24
	goto L8
L8:
	;
	v29 = v28 & v26
	goto L5
L9:
	;
	v30 = int32(1)
	v31 = v29 + v30
	v35 = v31 - v30
	if base.Ui32(int32(2)) <= base.Ui32(v31) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v62 = int32(1)
	goto L11
L11:
	;
	v64 = F__hash_getbuf(m, l0, v62, l2, int32(2))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L19
	}
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v18+v54<<(uint(int32(2))%32))+48))
	v62 = v58 + v31
	goto L11
L13:
	;
	v41 = int32(32) - base.I32_clz(v35)
	goto L15
L14:
	;
	v41 = int32(0)
	goto L15
L15:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v41) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v44 = int32(3)
	v54 = int32(base.Ui32(v35)>>(uint(v41-v44)%32))&v44 | v41<<(uint(int32(2))%32) - int32(30)
	goto L18
L17:
	;
	v54 = v41
	goto L18
L18:
	;
	goto L12
L19:
	;
	if int32(0) <= v64 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getbucketbuf_from_hashkey[0]))
	v83 = v69 + v64<<(uint(int32(13))%32) + int32(-8192)
	goto L2
L21:
	;
	goto L22
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getbucketbuf_from_hashkey[1]))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v64^int32(-1))<<(uint(int32(2))%32))))
	v83 = v82
	goto L2
L23:
	;
	F_UnlockReleaseBuffer(m, v64)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L25
L25:
	;
	v102 = F__hash_getcachedmetap(m, l0, v11+int32(12), int32(1))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L3
	} else {
		goto L28
	}
L27:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+16)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v166+v165)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v102)+24))
	if base.Ui32(v168) <= base.Ui32(v169) {
		v177 = v146
		v178 = v102
		goto L1
	} else {
		goto L47
	}
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)+24))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v102)+28))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102)+32))
	v108 = l1 & v105
	if base.Ui32(v108) <= base.Ui32(v104) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v111 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v110 = int32(-1)
	goto L32
L31:
	;
	v110 = v106
	goto L32
L32:
	;
	v111 = v110 & v108
	goto L29
L33:
	;
	v112 = int32(1)
	v113 = v111 + v112
	v117 = v113 - v112
	if base.Ui32(int32(2)) <= base.Ui32(v113) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v144 = int32(1)
	goto L35
L35:
	;
	v146 = F__hash_getbuf(m, l0, v144, l2, int32(2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L3
	} else {
		goto L43
	}
L36:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v102+v136<<(uint(int32(2))%32))+48))
	v144 = v140 + v113
	goto L35
L37:
	;
	v123 = int32(32) - base.I32_clz(v117)
	goto L39
L38:
	;
	v123 = int32(0)
	goto L39
L39:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v123) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v126 = int32(3)
	v136 = int32(base.Ui32(v117)>>(uint(v123-v126)%32))&v126 | v123<<(uint(int32(2))%32) - int32(30)
	goto L42
L41:
	;
	v136 = v123
	goto L42
L42:
	;
	goto L36
L43:
	;
	if v146 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getbucketbuf_from_hashkey[1]))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v151+(v146^int32(-1))<<(uint(int32(2))%32))))
	v165 = v157
	goto L27
L45:
	;
	goto L46
L46:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F__hash_getbucketbuf_from_hashkey[0]))
	v165 = v159 + v146<<(uint(int32(13))%32) + int32(-8192)
	goto L27
L47:
	;
	F_UnlockReleaseBuffer(m, v146)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	goto L25
L49:
	;
	F_ReleaseBuffer(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L3
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if l3 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v178
	goto L55
L54:
	;
	goto L55
L55:
	;
	m.G0 = v11 + int32(16)
	return v177
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
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
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
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
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
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v521 int64
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v542 int64
	_ = v542
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
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
	v593 = m.ExcPending
	if v593 != 0 {
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
	v575 = m.ExcPending
	if v575 != 0 {
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
	F_UnlockBuffer(m, v77)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v244))) = base.I64_rotl(v225, int64(32))
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
	v253 = v249 + int32(1)
	if v253 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_LockBufferInternal(m, v77, int32(3))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L2
	} else {
		goto L103
	}
L53:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[3]))
	if v257 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L2
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v261 = F__hash_getnewbuf(m, l0, int32(1), l2)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L2
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v206)+48))
	if int32(0) <= v261 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v281)+16)))
	v283 = v282 + v281
	*(*int32)(unsafe.Add(mBase, uint32(v283)+12)) = int32(-8388606)
	*(*int64)(unsafe.Add(mBase, uint32(v283)+4)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v263
	F_MarkBufferDirty(m, v261)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L2
	} else {
		goto L63
	}
L60:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v281 = v267 + v261<<(uint(int32(13))%32) + int32(-8192)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v274+(v261^int32(-1))<<(uint(int32(2))%32))))
	v281 = v280
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
	if int32(0) <= v261 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	goto L66
L66:
	;
	F_UnlockReleaseBuffer(m, v261)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L2
	} else {
		goto L72
	}
L67:
	;
	F_log_newpage(m, l0, l2, int32(1), v309, int32(1))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L2
	} else {
		goto L71
	}
L68:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v309 = v295 + v261<<(uint(int32(13))%32) + int32(-8192)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v302+(v261^int32(-1))<<(uint(int32(2))%32))))
	v309 = v308
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
	v323 = int32(1)
	goto L74
L74:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[3]))
	if v339 != 0 {
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
	v341 = m.ExcPending
	if v341 != 0 {
		goto L2
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v342 = int32(1)
	v343 = v323 + v342
	v347 = v343 - v342
	if base.Ui32(int32(2)) <= base.Ui32(v343) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L78
L80:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v206+int32(72)+v366<<(uint(int32(2))%32))))
	v371 = v370 + v343
	v372 = F__hash_getnewbuf(m, l0, v371, l2)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L2
	} else {
		goto L87
	}
L81:
	;
	v353 = int32(32) - base.I32_clz(v347)
	goto L83
L82:
	;
	v353 = int32(0)
	goto L83
L83:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v353) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v356 = int32(3)
	v366 = int32(base.Ui32(v347)>>(uint(v353-v356)%32))&v356 | v353<<(uint(int32(2))%32) - int32(30)
	goto L86
L85:
	;
	v366 = v353
	goto L86
L86:
	;
	goto L80
L87:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v206)+48))
	v375 = int32(0)
	v376 = base.B2i32(v375 <= v372)
	if v376 == v375 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v395 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v394)+16)))
	v396 = v395 + v394
	*(*int32)(unsafe.Add(mBase, uint32(v396)+12)) = int32(-8388606)
	*(*int32)(unsafe.Add(mBase, uint32(v396)+8)) = v323
	*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v396))) = v374
	F_MarkBufferDirty(m, v372)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L2
	} else {
		goto L92
	}
L89:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v380+(v372^int32(-1))<<(uint(int32(2))%32))))
	v394 = v386
	goto L88
L90:
	;
	goto L91
L91:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v394 = v388 + v372<<(uint(int32(13))%32) + int32(-8192)
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
	if v376 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L95
L95:
	;
	F_UnlockReleaseBuffer(m, v372)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L2
	} else {
		goto L101
	}
L96:
	;
	F_log_newpage(m, l0, l2, v371, v422, int32(1))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L2
	} else {
		goto L100
	}
L97:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v408+(v372^int32(-1))<<(uint(int32(2))%32))))
	v422 = v414
	goto L96
L98:
	;
	goto L99
L99:
	;
	v416 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v422 = v416 + v372<<(uint(int32(13))%32) + int32(-8192)
	goto L96
L100:
	;
	goto L95
L101:
	;
	if v323 != v249 {
		v323 = v343
		goto L74
	} else {
		goto L102
	}
L102:
	;
	goto L75
L103:
	;
	v451 = v249 + int32(2)
	v452 = F__hash_getnewbuf(m, l0, v451, l2)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	v454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206)+44)))
	if v452 < int32(0) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	F_MarkBufferDirty(m, v452)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L2
	} else {
		goto L116
	}
L106:
	;
	goto L111
L107:
	;
	v459 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v459+(v452^int32(-1))<<(uint(int32(2))%32))))
	v473 = v465
	goto L106
L108:
	;
	goto L109
L109:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v473 = v467 + v452<<(uint(int32(13))%32) + int32(-8192)
	goto L106
L111:
	;
	goto L112
L112:
	;
	v475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v473)+16)))
	v476 = v473 + v475
	*(*int64)(unsafe.Add(mBase, uint32(v476)+8)) = int64(-36028775544127489)
	*(*int64)(unsafe.Add(mBase, uint32(v476))) = int64(-1)
	if v454 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	base.MemoryFill(m, v473+int32(24), int32(255), v454)
	goto L115
L114:
	;
	goto L115
L115:
	;
	v486 = v454 + int32(24)
	*(*uint16)(unsafe.Add(mBase, uint32(v473)+12)) = uint16(v486)
	goto L105
L116:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v206)+68))
	if base.Ui32(int32(1024)) <= base.Ui32(v490) {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206+v490<<(uint(int32(2))%32))+468)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v206)+68)) = v490 + int32(1)
	F_MarkBufferDirty(m, v77)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
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
	v502 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206)+44)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+32)) = uint16(v502)
	F_XLogBeginInsert(m)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L2
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	F_UnlockReleaseBuffer(m, v452)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L2
	} else {
		goto L135
	}
L122:
	;
	F_XLogRegisterData(m, v21+int32(32), int32(2))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	F_XLogRegisterBuffer(m, int32(0), v452, int32(6))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	F_XLogRegisterBuffer(m, int32(1), v77, int32(8))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	v521 = F_XLogInsert(m, int32(12), int32(16))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L2
	} else {
		goto L126
	}
L126:
	;
	if v452 < int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v542 = base.I64_rotl(v521, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v540))) = v542
	if v77 < int32(0) {
		goto L132
	} else {
		goto L133
	}
L128:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v526+(v452^int32(-1))<<(uint(int32(2))%32))))
	v540 = v532
	goto L127
L129:
	;
	goto L130
L130:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v540 = v534 + v452<<(uint(int32(13))%32) + int32(-8192)
	goto L127
L131:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v561))) = v542
	goto L121
L132:
	;
	v547 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[1]))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v547+(v77^int32(-1))<<(uint(int32(2))%32))))
	v561 = v553
	goto L131
L133:
	;
	goto L134
L134:
	;
	v555 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init[2]))
	v561 = v555 + v77<<(uint(int32(13))%32) + int32(-8192)
	goto L131
L135:
	;
	F_UnlockReleaseBuffer(m, v77)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L2
	} else {
		goto L136
	}
L136:
	;
	m.G0 = v21 + int32(48)
	return v253
L137:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v576 + int32(4)
	F_errmsg_internal(m, int32(_a_F__hash_init_4), v21+int32(16))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L2
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F__hash_init_5), int32(345), int32(_a_F__hash_init_6))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
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
	v596 = m.ExcPending
	if v596 != 0 {
		goto L2
	} else {
		goto L141
	}
L141:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v597 + int32(4)
	F_errmsg(m, int32(_a_F__hash_init_7), v21)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L2
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F__hash_init_5), int32(455), int32(_a_F__hash_init_6))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
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
	var v81 int32
	_ = v81
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
	var v100 int32
	_ = v100
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
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v272 int32
	_ = v272
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v492 int32
	_ = v492
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v672 int64
	_ = v672
	var v673 int32
	_ = v673
	var v674 int64
	_ = v674
	var v675 int32
	_ = v675
	var v676 int64
	_ = v676
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v696 int64
	_ = v696
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
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
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72+(l4^int32(-1))*int32(56))+16))
	v87 = v78
	goto L9
L11:
	;
	goto L12
L12:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[3]))
	v81 = int32(56)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v80+l4*v81-v81)+16))
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
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91+(l5^int32(-1))*int32(56))+16))
	v106 = v97
	goto L13
L15:
	;
	goto L16
L16:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[3]))
	v100 = int32(56)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99+l5*v100-v100)+16))
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
	v125 = v11
	v127 = v67
	v128 = v49 + v48
	goto L20
L19:
	;
	v594 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v593)+16)))
	F_LockBufferInternal(m, l5, int32(3))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L17
	} else {
		goto L103
	}
L20:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+12)))
	if base.Ui32(v135) < base.Ui32(int32(25)) {
		v427 = v120
		v431 = v124
		v432 = v125
		v434 = v127
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v587 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v593 = v587 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L19
L22:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if l4 == v122 {
		goto L69
	} else {
		goto L70
	}
L23:
	;
	v139 = v135 + int32(_a_F__hash_splitbucket_0)
	if v139&int32(_a_F__hash_splitbucket_1) == int32(0) {
		v427 = v120
		v431 = v124
		v432 = v125
		v434 = v127
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
		v398 = v162
		v402 = v166
		v403 = v167
		v405 = v169
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v427 = v398
	v431 = v402
	v432 = v403
	v434 = v405
	goto L22
L27:
	;
	if v161 != int32(base.Ui32(v139)>>(uint(int32(2))%32))&int32(_a_F__hash_splitbucket_3) {
		v161 = v161 + int32(1)
		v162 = v398
		v166 = v402
		v167 = v403
		v169 = v405
		goto L25
	} else {
		goto L67
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
		v398 = v162
		v402 = v166
		v403 = v167
		v405 = v169
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
		v398 = v162
		v402 = v166
		v403 = v167
		v405 = v169
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
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+14)))
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+12)))
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
	if base.Ui32(v228) < base.Ui32(v167+v234) {
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
	v238 = int32(_a_F__hash_splitbucket_8)
	v240 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4])) = v240 + int32(1)
	F__hash_pgaddmultitup(m, l0, v162, v29+int32(16), v29+int32(1648), v217)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L17
	} else {
		goto L51
	}
L49:
	;
	v361 = v162
	v365 = v166
	v366 = v167
	v368 = v169
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29+int32(16)+v365&int32(_a_F__hash_splitbucket_3)<<(uint(int32(2))%32)))) = v210
	v398 = v361
	v402 = v365 + int32(1)
	v403 = v366 + v234
	v405 = v368
	goto L27
L51:
	;
	F_MarkBufferDirty(m, v162)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	F_log_split_page(m, l0, v162)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L17
	} else {
		goto L53
	}
L53:
	;
	v254 = int32(_a_F__hash_splitbucket_8)
	v256 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4])) = v256 - int32(1)
	F_UnlockBuffer(m, v162)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	if v217 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v272 = int32(0)
	goto L58
L56:
	;
	goto L57
L57:
	;
	v326 = F__hash_addovflpage(m, l0, l1, v162, base.B2i32(l5 == v162))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L17
	} else {
		goto L63
	}
L58:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(16)+v272<<(uint(int32(2))%32))))
	F_pfree(m, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L17
	} else {
		goto L60
	}
L59:
	;
	goto L57
L60:
	;
	v297 = v272 + int32(1)
	if v297 != v217 {
		v272 = v297
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v361 = v326
	v365 = int32(0)
	v366 = v347
	v368 = v348
	goto L50
L63:
	;
	if v326 < int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v332+(v326^int32(-1))<<(uint(int32(2))%32))))
	v347 = int32(0)
	v348 = v338
	goto L62
L65:
	;
	goto L66
L66:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v347 = int32(0)
	v348 = v341 + v326<<(uint(int32(13))%32) + int32(-8192)
	goto L62
L67:
	;
	goto L26
L68:
	;
	if v442 == int32(-1) {
		goto L75
	} else {
		goto L76
	}
L69:
	;
	F_UnlockBuffer(m, l4)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L17
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	F_UnlockReleaseBuffer(m, v122)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L17
	} else {
		goto L73
	}
L72:
	;
	goto L68
L73:
	;
	goto L68
L74:
	;
	goto L21
L75:
	;
	v450 = int32(_a_F__hash_splitbucket_8)
	v452 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4])) = v452 + int32(1)
	F__hash_pgaddmultitup(m, l0, v427, v29+int32(16), v29+int32(1648), v431&int32(_a_F__hash_splitbucket_3))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L17
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v558 = F_ReadBuffer(m, l0, v442)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L17
	} else {
		goto L96
	}
L78:
	;
	F_MarkBufferDirty(m, v427)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L17
	} else {
		goto L79
	}
L79:
	;
	F_log_split_page(m, l0, v427)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L17
	} else {
		goto L80
	}
L80:
	;
	v468 = int32(_a_F__hash_splitbucket_8)
	v470 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4])) = v470 - int32(1)
	if l5 == v427 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v480 = v431 & int32(_a_F__hash_splitbucket_3)
	if v480 != 0 {
		goto L87
	} else {
		goto L88
	}
L82:
	;
	F_UnlockBuffer(m, l5)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L17
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	F_UnlockReleaseBuffer(m, v427)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
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
	v492 = int32(0)
	goto L90
L88:
	;
	goto L89
L89:
	;
	F_LockBufferInternal(m, l4, int32(3))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L17
	} else {
		goto L94
	}
L90:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(16)+v492<<(uint(int32(2))%32))))
	F_pfree(m, v513)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L17
	} else {
		goto L92
	}
L91:
	;
	goto L89
L92:
	;
	v517 = v492 + int32(1)
	if v517 != v480 {
		v492 = v517
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	if int32(0) <= l4 {
		goto L74
	} else {
		goto L95
	}
L95:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v551+(l4^int32(-1))<<(uint(int32(2))%32))))
	v593 = v557
	goto L19
L96:
	;
	F_LockBufferInternal(m, v558, int32(1))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L17
	} else {
		goto L97
	}
L97:
	;
	F__hash_checkpage(m, l0, v558, int32(1))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L17
	} else {
		goto L98
	}
L98:
	;
	if v558 < int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v584 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583)+16)))
	v120 = v427
	v122 = v558
	v123 = v583
	v124 = v431
	v125 = v432
	v127 = v434
	v128 = v584 + v583
	goto L20
L100:
	;
	v569 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v569+(v558^int32(-1))<<(uint(int32(2))%32))))
	v583 = v575
	goto L99
L101:
	;
	goto L102
L102:
	;
	v577 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v583 = v577 + v558<<(uint(int32(13))%32) + int32(-8192)
	goto L99
L103:
	;
	v598 = v593 + v594
	if l5 < int32(0) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v617 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v616)+16)))
	v618 = int32(_a_F__hash_splitbucket_8)
	v620 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4])) = v620 + int32(1)
	v624 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v598)+12)))
	v626 = v624 & int32(_a_F__hash_splitbucket_9)
	*(*uint16)(unsafe.Add(mBase, uint32(v598)+12)) = uint16(v626)
	v628 = v616 + v617
	v629 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v628)+12)))
	v631 = v629 & int32(_a_F__hash_splitbucket_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v628)+12)) = uint16(v631)
	v633 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v598)+12)))
	v635 = v633 | int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v598)+12)) = uint16(v635)
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L17
	} else {
		goto L108
	}
L105:
	;
	v602 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v602+(l5^int32(-1))<<(uint(int32(2))%32))))
	v616 = v608
	goto L104
L106:
	;
	goto L107
L107:
	;
	v610 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v616 = v610 + l5<<(uint(int32(13))%32) + int32(-8192)
	goto L104
L108:
	;
	F_MarkBufferDirty(m, l5)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L17
	} else {
		goto L109
	}
L109:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+118)))
	if v642 != int32(112) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if l4 < int32(0) {
		goto L125
	} else {
		goto L126
	}
L111:
	;
	v674 = F_XLogGetFakeLSN(m, l0)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L17
	} else {
		goto L123
	}
L112:
	;
	v646 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[5]))
	if v646 <= int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v649 != 0 {
		goto L111
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v651 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v598)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)) = uint16(v651)
	v653 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v628)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)) = uint16(v653)
	F_XLogBeginInsert(m)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L17
	} else {
		goto L118
	}
L116:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v650 != 0 {
		goto L111
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	F_XLogRegisterData(m, v29+int32(12), int32(4))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L17
	} else {
		goto L119
	}
L119:
	;
	F_XLogRegisterBuffer(m, int32(0), l4, int32(8))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L17
	} else {
		goto L120
	}
L120:
	;
	F_XLogRegisterBuffer(m, int32(1), l5, int32(8))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L17
	} else {
		goto L121
	}
L121:
	;
	v672 = F_XLogInsert(m, int32(12), int32(96))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L17
	} else {
		goto L122
	}
L122:
	;
	v676 = v672
	goto L110
L123:
	;
	v676 = v674
	goto L110
L124:
	;
	v696 = base.I64_rotl(v676, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v694))) = v696
	if l5 < int32(0) {
		goto L129
	} else {
		goto L130
	}
L125:
	;
	v680 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v680+(l4^int32(-1))<<(uint(int32(2))%32))))
	v694 = v686
	goto L124
L126:
	;
	goto L127
L127:
	;
	v688 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v694 = v688 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L124
L128:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v715))) = v696
	v717 = int32(_a_F__hash_splitbucket_8)
	v719 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[4])) = v719 - int32(1)
	v723 = F_IsBufferCleanupOK(m, l4)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L17
	} else {
		goto L132
	}
L129:
	;
	v701 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[0]))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v701+(l5^int32(-1))<<(uint(int32(2))%32))))
	v715 = v707
	goto L128
L130:
	;
	goto L131
L131:
	;
	v709 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[1]))
	v715 = v709 + l5<<(uint(int32(13))%32) + int32(-8192)
	goto L128
L132:
	;
	F_UnlockBuffer(m, l5)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L17
	} else {
		goto L133
	}
L133:
	;
	if v723 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	m.G0 = v29 + int32(2464)
	return
L135:
	;
	if l4 < int32(0) {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	goto L137
L137:
	;
	F_UnlockBuffer(m, l4)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L17
	} else {
		goto L143
	}
L138:
	;
	v746 = int32(0)
	F_hashbucketcleanup(m, l0, l2, l4, v745, v746, l7, l8, l9, v746, v746, int32(1), v746, v746)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L17
	} else {
		goto L142
	}
L139:
	;
	v730 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[2]))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v730+(l4^int32(-1))*int32(56))+16))
	v745 = v736
	goto L138
L140:
	;
	goto L141
L141:
	;
	v738 = *(*int32)(unsafe.Add(mBase, _c_F__hash_splitbucket[3]))
	v739 = int32(56)
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v738+l4*v739-v739)+16))
	v745 = v744
	goto L138
L142:
	;
	goto L134
L143:
	;
	goto L134
}
func F_hash_aclitem(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	return base.I64_extend_i32_u(v3 + v4 + v6)
}
func F_hash_aclitem_extended(m *base.Module, l0 int32) int64 {
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
	var v10 int64
	_ = v10
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
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
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v9 = v5 + v6 + v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if v10 == int64(0) {
		return base.I64_extend_i32_u(v9)
	} else {
		if v10 == int64(0) {
			v21 = int32(-1636608428)
			v60 = v21
			v61 = v21
			v64 = int32(0)
		} else {
			v24 = base.I32_wrap_i64(v10)
			v26 = v24 + int32(1021750440)
			v31 = base.I32_wrap_i64(int64(base.Ui64(v10)>>(uint(int64(32))%64))) ^ int32(-415931063)
			v37 = v24 - v31 - int32(1636608428) ^ base.I32_rotl(v31, int32(6))
			v41 = v26 - v37 ^ base.I32_rotl(v37, int32(8))
			v42 = v31 + v26
			v43 = v37 + v42
			v44 = v41 + v43
			v48 = v42 - v41 ^ base.I32_rotl(v41, int32(16))
			v52 = v43 - v48 ^ base.I32_rotl(v48, int32(19))
			v57 = v48 + v44
			v58 = v52 + v57
			v60 = v58
			v61 = v57
			v64 = v44 - v52 ^ base.I32_rotl(v52, int32(4)) ^ v58
		}
		v65 = int32(14)
		v67 = v64 - base.I32_rotl(v60, v65)
		v72 = v67 ^ (v9 + v61) - base.I32_rotl(v67, int32(11))
		v76 = v60 ^ v72 - base.I32_rotl(v72, int32(25))
		v80 = v76 ^ v67 - base.I32_rotl(v76, int32(16))
		v84 = v80 ^ v72 - base.I32_rotl(v80, int32(4))
		v88 = v84 ^ v76 - base.I32_rotl(v84, v65)
		return base.I64_extend_i32_u(v88)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v88^v80-base.I32_rotl(v88, int32(24)))
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
	var v35 int32
	_ = v35
	var v38 float64
	_ = v38
	var v42 float64
	_ = v42
	var v44 int32
	_ = v44
	var v48 float64
	_ = v48
	var v49 float64
	_ = v49
	var v52 float64
	_ = v52
	var v54 float64
	_ = v54
	var v60 float64
	_ = v60
	var v66 float64
	_ = v66
	var v68 float64
	_ = v68
	var v71 float64
	_ = v71
	var v74 float64
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v104 float64
	_ = v104
	var v109 int64
	_ = v109
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
		v35 = int32(32)
		v38 = float64(1024)
		v42 = *(*float64)(unsafe.Add(mBase, _c_F_hash_agg_set_limits[0]))
		v44 = *(*int32)(unsafe.Add(mBase, _c_F_hash_agg_set_limits[1]))
		v48 = base.F64_mul(base.F64_mul(v42, base.F64_convert_i32_s(v44)), v38)
		v49 = float64(4.294967295e+09)
		if base.F64_lt(v48, v49) != 0 {
			v52 = v48
		} else {
			v52 = v49
		}
		v54 = base.F64_convert_i32_u(base.I32_trunc_sat_f64_u(v52))
		v60 = base.F64_mul(base.F64_add(base.F64_mul(v54, float64(0.25)), float64(-8192)), float64(0.0001220703125))
		v66 = base.F64_add(base.F64_div(base.F64_mul(l0, base.F64_mul(l1, float64(1.5))), v54), float64(1))
		if base.F64_gt(v66, v60) != 0 {
			v68 = v60
		} else {
			v68 = v66
		}
		if base.F64_lt(v68, float64(4)) != 0 {
			v71 = float64(4)
		} else {
			v71 = v68
		}
		if base.F64_gt(v71, float64(1024)) != 0 {
			v74 = v38
		} else {
			v74 = v71
		}
		v75 = base.I32_trunc_sat_f64_s(v74)
		if base.Ui32(int32(2)) <= base.Ui32(v75) {
			v83 = v35 - base.I32_clz(v75-int32(1))
		} else {
			v83 = int32(0)
		}
		if int32(31) < l2+v83 {
			v87 = v35 - l2
		} else {
			v87 = v83
		}
		if l5 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(1) << (uint(v87) % 32)
		} else {
		}
		v94 = int32(_a_F_hash_agg_set_limits_0)<<(uint(v87)%32) - int32(-8192)
		if base.Ui32(v94<<(uint(int32(2))%32)) < base.Ui32(v25) {
			v102 = v25 - v94
		} else {
			v102 = base.I32_trunc_sat_f64_u(base.F64_mul(v26, float64(0.75)))
		}
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
		v104 = base.F64_convert_i32_u(v102)
		if base.F64_gt(v104, l0) != 0 {
			v109 = base.I64_trunc_sat_f64_u(base.F64_div(v104, l0))
		} else {
			v109 = int64(1)
		}
		*(*int64)(unsafe.Add(mBase, uint32(l4))) = v109
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
func F_hash_array(m *base.Module, l0 int32) int64 {
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v157 int64
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L41
	}
L11:
	;
	v84 = int32(*(*int8)(unsafe.Add(mBase, uint32(v81)+11)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+10)))
	v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v81)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v81 + int32(132)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v93 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+56)) = uint8(v93)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v92
	v96 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+58)) = uint16(v96)
	v98 = F_ArrayGetNItemsSafe(m, v28, v36)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L24
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v42 == v39 {
		v81 = v41
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
	v51 = base.B2i32(v39 != int32(2249))
	if base.B2i32(v47 == int32(0))&v51 != 0 {
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
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v77
	v81 = v77
	goto L11
L19:
	;
	v77 = v45
	goto L18
L20:
	;
	goto L21
L21:
	;
	v53 = int32(_a_F_hash_array_0)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_hash_array[0]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_hash_array[0])) = v57
	v60 = F_palloc0(m, int32(328))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(2249)
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+8)) = uint16(v64)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+10)) = uint8(v66)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+11)))
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+11)) = uint8(v68)
	F_fmgr_info(m, int32(_a_F_hash_array_1), v60+int32(132))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hash_array[0])) = v54
	v77 = v60
	goto L18
L24:
	;
	F_array_iter_setup(m, v14+int32(12), v17, v86, v85, v84)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v98 <= int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v157 = int64(1)
	goto L28
L27:
	;
	v110 = int32(0)
	v112 = int32(1)
	goto L29
L28:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v158 == int32(-1) {
		goto L37
	} else {
		goto L38
	}
L29:
	;
	v124 = F_array_iter_next(m, v14+int32(12), v14+int32(11), v110)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	v157 = base.I64_extend_i32_u(v141)
	goto L28
L31:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
	if v126 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v138 = int32(0)
	goto L34
L33:
	;
	v128 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+72)) = uint8(v128)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v124
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v135 = m.T0[v134].(func(*base.Module, int32) int64)(m, v14+int32(40))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	v141 = v138 + v112*int32(31)
	v143 = v110 + int32(1)
	if v143 != v98 {
		v110 = v143
		v112 = v141
		goto L29
	} else {
		goto L36
	}
L35:
	;
	v138 = base.I32_wrap_i64(v135)
	goto L34
L36:
	;
	goto L30
L37:
	;
	m.G0 = v14 + int32(80)
	return v157
L38:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v17 == v161 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	F_pfree(m, v17)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v176 = F_format_type_be(m, v39)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v176
	F_errmsg(m, int32(_a_F_hash_array_2), v14)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_hash_array_3), int32(_a_F_hash_array_4), int32(_a_F_hash_array_5))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
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
func F_hash_array_extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
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
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v82 int32
	_ = v82
	var v90 int64
	_ = v90
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v118 int32
	_ = v118
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	v14 = m.G0
	v16 = v14 - int32(96)
	m.G0 = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v25 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v28 = int32(28)
	goto L5
L4:
	;
	v28 = int32(4)
	goto L5
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v19+v28)))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if v25 == int32(-1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40+v19)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	if v44 != 0 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v39 = v34
	v40 = int32(40)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v39 = v19 + int32(16)
	v40 = int32(12)
	goto L6
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L35
	}
L11:
	;
	v56 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55)+11)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+10)))
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v55 + int32(160)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v65 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+58)) = uint16(v65)
	v67 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+56)) = uint8(v67)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v64
	v70 = F_ArrayGetNItemsSafe(m, v30, v39)
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
	F_array_iter_setup(m, v16+int32(12), v19, v58, v57, v56)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v76 = int64(1)
	if int32(0) < v70 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v82 = int32(0)
	v90 = v76
	goto L23
L21:
	;
	v130 = v76
	goto L22
L22:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v133 == int32(-1) {
		goto L31
	} else {
		goto L32
	}
L23:
	;
	v97 = F_array_iter_next(m, v16+int32(12), v16+int32(11), v82)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v130 = v116
	goto L22
L25:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+11)))
	if v99 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v113 = int64(0)
	goto L28
L27:
	;
	v101 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+88)) = uint8(v101)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v31
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+72)) = uint8(v101)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = v97
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v111 = m.T0[v110].(func(*base.Module, int32) int64)(m, v16+int32(40))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v116 = v113 + v90*int64(31)
	v118 = v82 + int32(1)
	if v118 != v70 {
		v82 = v118
		v90 = v116
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v113 = v111
	goto L28
L30:
	;
	goto L24
L31:
	;
	m.G0 = v16 + int32(96)
	return v130
L32:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v19 == v136 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	F_pfree(m, v19)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v151 = F_format_type_be(m, v42)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v151
	F_errmsg(m, int32(_a_F_hash_array_extended_1), v16)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_hash_array_extended_2), int32(_a_F_hash_array_extended_3), int32(_a_F_hash_array_extended_4))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hash_bulkdelete_read_stream_cb(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v4) <= base.Ui32(v5) {
		v8 = v4 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v8
		if v4 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v14 = v8 - int32(1)
			if base.Ui32(int32(2)) <= base.Ui32(v8) {
				v20 = int32(32) - base.I32_clz(v14)
			} else {
				v20 = int32(0)
			}
			if base.Ui32(int32(10)) <= base.Ui32(v20) {
				v23 = int32(3)
				v33 = int32(base.Ui32(v14)>>(uint(v20-v23)%32))&v23 | v20<<(uint(int32(2))%32) - int32(30)
			} else {
				v33 = v20
			}
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v10+v33<<(uint(int32(2))%32))+48))
			v39 = v37
		} else {
			v39 = int32(0)
		}
		v43 = v39 + v8
	} else {
		v43 = int32(-1)
	}
	return v43
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
func F_hash_record_extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int64
	_ = v14
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
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
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
	var v36 int32
	_ = v36
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v141 int64
	_ = v141
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v193 int64
	_ = v193
	var v202 int32
	_ = v202
	var v203 int64
	_ = v203
	var v204 int32
	_ = v204
	var v207 int64
	_ = v207
	var v215 int64
	_ = v215
	var v218 int32
	_ = v218
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
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v252 int64
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	v14 = int64(0)
	v17 = m.G0
	v19 = v17 - int32(96)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	F_check_stack_depth(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v31 = F_lookup_rowtype_tupdesc(m, v29, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v22
	v36 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v36
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+84)) = uint16(v36)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = int32(base.Ui32(v34) >> (uint(int32(2)) % 32))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	if v46 == v36 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v29 == v67 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	v57 = F_MemoryContextAlloc(m, v52, v33<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v49 < v33 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v66 = v46
	v67 = v51
	goto L5
L9:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v57
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v62)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v33
	v66 = v62
	v67 = int32(0)
	goto L5
L10:
	;
	v116 = F_palloc_mul(m, int32(8), v33)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L25
	}
L11:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	if v69 == v30 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v72 = v66 + int32(20)
	v76 = v33 << (uint(int32(2)) % 32)
	if v72&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v76)) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v29
	goto L10
L16:
	;
	if v76 == int32(0) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v76 == int32(0) {
		goto L15
	} else {
		goto L24
	}
L19:
	;
	v86 = v66 + v76 + int32(20)
	v88 = v66 + int32(24)
	if base.Ui32(v88) < base.Ui32(v86) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v90 = v86
	goto L22
L21:
	;
	v90 = v88
	goto L22
L22:
	;
	v97 = (v90-v66-int32(21))&int32(-4) + int32(4)
	if v97 == int32(0) {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	base.MemoryFill(m, v72, int32(0), v97)
	goto L15
L24:
	;
	base.MemoryFill(m, v72, int32(0), v76)
	goto L15
L25:
	;
	v119 = F_palloc_mul(m, int32(1), v33)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_heap_deform_tuple(m, v19+int32(76), v31, v116, v119)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v33 <= int32(0) {
		v252 = v14
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_pfree(m, v116)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L54
	}
L29:
	;
	v129 = int32(0)
	v141 = v14
	goto L30
L30:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v150 = v31 + v144<<(uint(int32(3))%32) + v129*int32(100)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+119)))
	if v151 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L49
	}
L32:
	;
	goto L31
L33:
	;
	v155 = v150 + int32(28)
	v158 = v66 + int32(20) + v129<<(uint(int32(2))%32)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v159 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v215 = v141
	goto L35
L35:
	;
	v218 = v129 + int32(1)
	if v33 != v218 {
		v129 = v218
		v141 = v215
		goto L30
	} else {
		goto L48
	}
L36:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v119))))
	if v177 != 0 {
		goto L44
	} else {
		goto L45
	}
L37:
	;
	v168 = F_lookup_type_cache(m, v166, int32(_a_F_hash_record_extended_0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L42
	}
L38:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v155)+68))
	v166 = v162
	goto L37
L39:
	;
	goto L40
L40:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v155)+68))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v163 == v164 {
		v174 = v159
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v166 = v163
	goto L37
L42:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v168)+164))
	if v170 == int32(0) {
		goto L32
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v168
	v174 = v168
	goto L36
L44:
	;
	v207 = int64(0)
	goto L46
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+20)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v174 + int32(160)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v155)+96))
	v185 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+34)) = uint16(v185)
	v187 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+32)) = uint8(v187)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v184
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v116+v129<<(uint(int32(3))%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+64)) = uint8(v187)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v26
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)) = uint8(v187)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v193
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v174)+160))
	v203 = m.T0[v202].(func(*base.Module, int32) int64)(m, v19+int32(16))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v215 = v207 + v141*int64(31)
	goto L35
L47:
	;
	v207 = v203
	goto L46
L48:
	;
	v252 = v215
	goto L28
L49:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v228 = F_format_type_be(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v228
	F_errmsg(m, int32(_a_F_hash_record_extended_1), v19)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_hash_record_extended_2), int32(2015), int32(_a_F_hash_record_extended_3))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
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
	F_pfree(m, v119)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if int32(0) <= v259 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_DecrTupleDescRefCount(m, v31)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v264 != v22 {
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
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	m.G0 = v19 + int32(96)
	return v252
L63:
	;
	goto L62
}
