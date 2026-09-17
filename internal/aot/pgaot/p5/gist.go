package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_gistFetchTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
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
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(176)
	m.G0 = v16
	v18 = int32(_a_F_gistFetchTuple_0)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_gistFetchTuple[0]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_gistFetchTuple[0])) = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+10)))
	if v4 < v24 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v36 = v4
	goto L4
L2:
	;
	v120 = v4
	goto L3
L3:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if v120 < v129 {
		goto L23
	} else {
		goto L24
	}
L4:
	;
	v45 = v36 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v47 = v16 + v36
	v48 = F_index_getattr_2(m, l2, v45, v46, v47)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v120 = v45
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	v53 = v36 * int32(28)
	v54 = l0 + v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+uint32(_c_F_gistFetchTuple[1])))
	if v55 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(v112)+10)))
	if v45 < v113 {
		v36 = v45
		goto L4
	} else {
		goto L22
	}
L9:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v56 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v54)+1816))
	if v88 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v59 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+174)) = uint8(v59)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+172)) = uint16(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+168)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v16)+164)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+160)) = v48
	v68 = v36 << (uint(int32(2)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v68+(l0+int32(_a_F_gistFetchTuple_1)))))
	v77 = F_FunctionCall1Coll(m, l0+int32(_a_F_gistFetchTuple_2)+v53, v74, v16+int32(160))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(32)+v36<<(uint(int32(2))%32)))) = int32(0)
	goto L8
L15:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	*(*int32)(unsafe.Add(mBase, uint32(v68+(v16+int32(32))))) = v79
	goto L8
L16:
	;
	v95 = v16 + int32(32) + v36<<(uint(int32(2))%32)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v96 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v102 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(32)+v36<<(uint(int32(2))%32)))) = int32(0)
	goto L8
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v48
	goto L8
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = int32(0)
	goto L8
L22:
	;
	goto L5
L23:
	;
	v136 = v120
	goto L26
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistFetchTuple[0])) = v19
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v177 = F_heap_form_tuple(m, v174, v16+int32(32), v16)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L6
	} else {
		goto L30
	}
L26:
	;
	v150 = v136 + int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v153 = F_index_getattr_2(m, l2, v150, v151, v16+v136)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L6
	} else {
		goto L28
	}
L27:
	;
	goto L25
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(32)+v136<<(uint(int32(2))%32)))) = v153
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v150 < v157 {
		v136 = v150
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	m.G0 = v16 + int32(176)
	return v177
}
func F_gistGetNodeBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = F_hash_search(m, v13, v10+int32(12), int32(1), v10+int32(11))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
		if v23 == int32(0) {
			v26 = int32(_a_F_gistGetNodeBuffer_0)
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_gistGetNodeBuffer[0]))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, _c_F_gistGetNodeBuffer[0])) = v29
			*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = l2
			v32 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v19)+16)) = uint16(v32)
			*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v32
			*(*int64)(unsafe.Add(mBase, uint32(v19)+4)) = int64(-4294967296)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v38 <= l2 {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v42 = l2 + int32(1)
				v45 = F_repalloc(m, v40, v42<<(uint(int32(2))%32))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v45
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					if v48 <= l2 {
						v51 = v48
						for {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v57+v51<<(uint(int32(2))%32)))) = int32(0)
							v64 = v51 + int32(1)
							if v64 <= l2 {
								v51 = v64
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v42
					v82 = l2 << (uint(int32(2)) % 32)
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v82+v83)))
					v86 = F_lcons(m, v19, v85)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v88+v82))) = v86
						*(*int32)(unsafe.Add(mBase, _c_F_gistGetNodeBuffer[0])) = v27
						m.G0 = v10 + int32(16)
						return v19
					}
				}
			} else {
				v82 = l2 << (uint(int32(2)) % 32)
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v82+v83)))
				v86 = F_lcons(m, v19, v85)
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v88+v82))) = v86
					*(*int32)(unsafe.Add(mBase, _c_F_gistGetNodeBuffer[0])) = v27
					m.G0 = v10 + int32(16)
					return v19
				}
			}
		} else {
			m.G0 = v10 + int32(16)
			return v19
		}
	}
}
func F_gistLoadNodeBuffer(m *base.Module, l0 int32, l1 int32) {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
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
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
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
	var v77 int32
	_ = v77
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v11 != 0 {
		m.G0 = v9 + int32(16)
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v12 <= int32(0) {
			m.G0 = v9 + int32(16)
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v17 = F_MemoryContextAllocZero(m, v15, int32(_a_F_gistLoadNodeBuffer_0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v17))) = int64(35154307317759)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v17
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v25 = F_BufFileSeekBlock(m, v22, base.I64_extend_i32_s(v23))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					if v25 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v23
							F_errmsg_internal(m, int32(_a_F_gistLoadNodeBuffer_1), v9)
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_gistLoadNodeBuffer_2), int32(753), int32(_a_F_gistLoadNodeBuffer_3))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_BufFileReadExact(m, v22, v17, int32(_a_F_gistLoadNodeBuffer_0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v31 < v32 {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v45 = v34
								v46 = v31
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v46 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v45+v46<<(uint(int32(2))%32)))) = v30
								v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
								if v54 == int32(0) {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if v57 < v58 {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v71 = v60
										v72 = v57
										*(*int32)(unsafe.Add(mBase, uint32(v71+v72<<(uint(int32(2))%32)))) = l1
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v77 + int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(-1)
										m.G0 = v9 + int32(16)
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v58 << (uint(int32(1)) % 32)
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v67 = F_repalloc(m, v64, v58<<(uint(int32(3))%32))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v67
											v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
											v71 = v67
											v72 = v70
											*(*int32)(unsafe.Add(mBase, uint32(v71+v72<<(uint(int32(2))%32)))) = l1
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v77 + int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(-1)
											m.G0 = v9 + int32(16)
											return
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(-1)
									m.G0 = v9 + int32(16)
									return
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v32 << (uint(int32(1)) % 32)
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v41 = F_repalloc(m, v38, v32<<(uint(int32(3))%32))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v41
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v45 = v41
									v46 = v44
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v46 + int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v45+v46<<(uint(int32(2))%32)))) = v30
									v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
									if v54 == int32(0) {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v57 < v58 {
											v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v71 = v60
											v72 = v57
											*(*int32)(unsafe.Add(mBase, uint32(v71+v72<<(uint(int32(2))%32)))) = l1
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v77 + int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(-1)
											m.G0 = v9 + int32(16)
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v58 << (uint(int32(1)) % 32)
											v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v67 = F_repalloc(m, v64, v58<<(uint(int32(3))%32))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v67
												v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
												v71 = v67
												v72 = v70
												*(*int32)(unsafe.Add(mBase, uint32(v71+v72<<(uint(int32(2))%32)))) = l1
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v77 + int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(-1)
												m.G0 = v9 + int32(16)
												return
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(-1)
										m.G0 = v9 + int32(16)
										return
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
func F_gistSplitByKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v324 int32
	_ = v324
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v521 int32
	_ = v521
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v724 int32
	_ = v724
	var v733 int32
	_ = v733
	var v742 int32
	_ = v742
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v754 float32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v760 float32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v771 float32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 float32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v784 float32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 float32
	_ = v788
	var v789 int32
	_ = v789
	var v798 int64
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v829 int32
	_ = v829
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v966 float32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v992 int32
	_ = v992
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1020 int32
	_ = v1020
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1066 float32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1145 int32
	_ = v1145
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1237 int32
	_ = v1237
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1320 int32
	_ = v1320
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1409 int32
	_ = v1409
	var v1432 int32
	_ = v1432
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1465 int32
	_ = v1465
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1554 int32
	_ = v1554
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1580 int32
	_ = v1580
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1609 int32
	_ = v1609
	var v1611 int32
	_ = v1611
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1621 float32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1636 float32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1678 int32
	_ = v1678
	var v1682 int32
	_ = v1682
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1692 int32
	_ = v1692
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1735 int32
	_ = v1735
	var v1738 int32
	_ = v1738
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1757 int32
	_ = v1757
	var v1760 int32
	_ = v1760
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1774 int32
	_ = v1774
	var v1776 int32
	_ = v1776
	var v1789 int32
	_ = v1789
	var v1792 int32
	_ = v1792
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1814 int32
	_ = v1814
	var v1817 int32
	_ = v1817
	var v1823 int32
	_ = v1823
	var v1825 int32
	_ = v1825
	var v1843 int32
	_ = v1843
	var v1861 int64
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1863 int64
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1915 int32
	_ = v1915
	var v1918 int32
	_ = v1918
	var v1922 int32
	_ = v1922
	var v1928 int32
	_ = v1928
	var v1931 int32
	_ = v1931
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1947 int32
	_ = v1947
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1979 int32
	_ = v1979
	var v1982 int32
	_ = v1982
	var v2000 int32
	_ = v2000
	var v2003 int32
	_ = v2003
	var v2007 int32
	_ = v2007
	var v2013 int32
	_ = v2013
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2030 int32
	_ = v2030
	var v2095 int32
	_ = v2095
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2196 int32
	_ = v2196
	v31 = m.G0
	v33 = v31 - int32(608)
	m.G0 = v33
	v36 = l3 + int32(1)
	v37 = int32(4)
	v41 = F_palloc(m, v36<<(uint(v37)%32)|v37)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v36
	v45 = l3 << (uint(int32(1)) % 32)
	v46 = F_palloc(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(0) < l3 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	m.G0 = v33 + int32(608)
	return
L5:
	;
	v2189 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(v2189)))
	if v2190 < int32(2) {
		goto L4
	} else {
		goto L283
	}
L6:
	;
	if l6 != 0 {
		goto L4
	} else {
		goto L282
	}
L7:
	;
	v401 = l5 + int32(160)
	v402 = v401 + l6
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402))))
	v404 = int32(1)
	v405 = v403 ^ v404
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+12)) = uint8(v405)
	v408 = l5 + int32(320)
	v409 = v408 + l6
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	v412 = v410 ^ v404
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+28)) = uint8(v412)
	v415 = l5 + int32(32)
	v417 = l6 << (uint(int32(2)) % 32)
	v418 = v415 + v417
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v419
	v422 = l5 + int32(192)
	v423 = v422 + v417
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v424
	v431 = l4 + v417
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v431)+uint32(_c_F_gistSplitByKey[0])))
	v435 = F_FunctionCall2Coll(m, l4+l6*int32(28)+int32(_a_F_gistSplitByKey_0), v434, v41, l5)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L63
	}
L8:
	;
	v288 = l5 + l6
	v289 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v288)+320)) = uint8(v289)
	*(*uint8)(unsafe.Add(mBase, uint32(v288)+160)) = uint8(v289)
	v294 = l6 + v289
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if v294 < v296 {
		goto L49
	} else {
		goto L50
	}
L9:
	;
	v51 = v41 + int32(4)
	v54 = int32(1)
	v55 = l6 + v54
	v68 = int32(0)
	v70 = v54
	goto L12
L10:
	;
	goto L11
L11:
	;
	if l3 != 0 {
		goto L7
	} else {
		goto L48
	}
L12:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v92 = int32(4)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l2+v70<<(uint(int32(2))%32)-v92)))
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+6)))
	if int32(0) <= v101 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	if l3 == v174 {
		goto L8
	} else {
		goto L35
	}
L14:
	;
	v180 = v70 + int32(1)
	if v180 <= l3 {
		v68 = v174
		v70 = v180
		goto L12
	} else {
		goto L34
	}
L15:
	;
	F_gistdentryinit(m, l4, l6, v51+v70<<(uint(int32(4))%32), int32(0), l0, l1, v70&int32(_a_F_gistSplitByKey_1), int32(1))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L33
	}
L16:
	;
	F_gistdentryinit(m, l4, l6, v51+v70<<(uint(v92)%32), v153, l0, l1, v70&int32(_a_F_gistSplitByKey_1), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L32
	}
L17:
	;
	v147 = F_nocache_index_getattr(m, v100, v55, v91)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L31
	}
L18:
	;
	v104 = int32(4)
	v108 = v91 + v55<<(uint(v104)%32) + v104
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v109 < int32(0) {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+l6>>(uint(int32(3))%32))+8)))
	if v54<<(uint(l6&int32(7))%32)&v141 == int32(0) {
		goto L15
	} else {
		goto L30
	}
L21:
	;
	v114 = v100 + v109 + int32(8)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+6)))
	if v115 != int32(1) {
		v153 = v114
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+4)))
	switch v118 - int32(1) {
	case 0:
		goto L26
	case 1:
		goto L25
	default:
		goto L23
	case 3:
		goto L24
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L27
	}
L24:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v153 = v123
	goto L16
L25:
	;
	v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(v114))))
	v153 = v122
	goto L16
L26:
	;
	v121 = int32(*(*int8)(unsafe.Add(mBase, uint32(v114))))
	v153 = v121
	goto L16
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = base.I32_extend16_s(v118)
	F_errmsg_internal(m, int32(_a_F_gistSplitByKey_2), v33+int32(16))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_gistSplitByKey_3), int32(70), int32(_a_F_gistSplitByKey_4))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
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
	goto L17
L31:
	;
	v153 = v147
	goto L16
L32:
	;
	v174 = v68
	goto L14
L33:
	;
	v168 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v46+v68<<(uint(v168)%32)))) = uint16(v70)
	v174 = v68 + v168
	goto L14
L34:
	;
	goto L13
L35:
	;
	if v174 <= int32(0) {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = v46
	v188 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5+l6)+320)) = uint8(v188)
	v191 = F_palloc(m, v45)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v191
	v206 = int32(1)
	v208 = int32(0)
	goto L38
L38:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	if v227 <= v208 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if l6 != 0 {
		goto L4
	} else {
		goto L45
	}
L40:
	;
	if l3 != v206 {
		v206 = v206 + int32(1)
		v208 = v246
		goto L38
	} else {
		goto L44
	}
L41:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v237 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v236 + v237
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*uint16)(unsafe.Add(mBase, uint32(v240+v236<<(uint(v237)%32)))) = uint16(v206)
	v246 = v208
	goto L40
L42:
	;
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46+v208<<(uint(int32(1))%32)))))
	if v206 != v232 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v246 = v208 + int32(1)
	goto L40
L44:
	;
	goto L39
L45:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v251 != int32(1) {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+352)) = int32(0)
	F_gistunionsubkey(m, l4, l2, l5)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L5
L48:
	;
	goto L8
L49:
	;
	F_gistSplitByKey(m, l0, l1, l2, l3, l4, l5, v294)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v300 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v300
	v304 = F_palloc(m, v45)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L6
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v304
	v307 = F_palloc(m, v45)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = v307
	if l3 <= int32(0) {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v312 = int32(1)
	v324 = v312
	goto L56
L56:
	;
	if base.Ui32(v324) < base.Ui32(int32(base.Ui32(l3)>>(uint(v312)%32))) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L6
L58:
	;
	if base.B2i32(l3 == v324) == int32(0) {
		v324 = v324 + int32(1)
		goto L56
	} else {
		goto L62
	}
L59:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v347 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v346 + v347
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v350+v346<<(uint(v347)%32)))) = uint16(v324)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v356 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v355 + v356
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*uint16)(unsafe.Add(mBase, uint32(v359+v355<<(uint(v356)%32)))) = uint16(v324)
	goto L58
L62:
	;
	goto L57
L63:
	;
	v438 = l5 + int32(12)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v439 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v695 = l5 + int32(24)
	v697 = l5 + int32(8)
	v699 = l5 + int32(28)
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	if v700 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L65:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v643 = v638 + v439<<(uint(int32(1))%32) - int32(2)
	v644 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v643))))
	if v644 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L66:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	if v440 != 0 {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v444 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	if v444 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402))))
	v470 = int32(1)
	v471 = v469 ^ v470
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+12)) = uint8(v471)
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	v475 = v473 ^ v470
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+28)) = uint8(v475)
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v477
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v479
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v485 = (v481 - v470) & int32(_a_F_gistSplitByKey_1)
	v489 = v485<<(uint(v470)%32) + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v489
	v491 = F_palloc(m, v489)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L78
	}
L74:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = l6 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v449 + int32(4)
	F_errmsg(m, int32(_a_F_gistSplitByKey_5), v33)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errhint(m, int32(_a_F_gistSplitByKey_6), int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_gistSplitByKey_7), int32(448), int32(_a_F_gistSplitByKey_8))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L73
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v491
	v494 = F_palloc(m, v489)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v496 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v496
	v502 = v481 & int32(_a_F_gistSplitByKey_1)
	if v502 != int32(1) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v505 = int32(2)
	if base.Ui32(v502) <= base.Ui32(v505) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v597 = int32(4)
	v601 = F_palloc(m, v596<<(uint(v597)%32)|v597)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L93
	}
L83:
	;
	v508 = v505
	goto L85
L84:
	;
	v508 = v502
	goto L85
L85:
	;
	v509 = int32(1)
	v521 = v509
	goto L86
L86:
	;
	if base.Ui32(v521) <= base.Ui32(int32(base.Ui32(v485)>>(uint(v509)%32))) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L82
L88:
	;
	v564 = v521 + int32(1)
	if v564 != v508 {
		v521 = v564
		goto L86
	} else {
		goto L92
	}
L89:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v545 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v543+v544<<(uint(v545)%32)))) = uint16(v521)
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v549 + v545
	goto L88
L90:
	;
	goto L91
L91:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v555 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v553+v554<<(uint(v555)%32)))) = uint16(v521)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v559 + v555
	goto L88
L92:
	;
	goto L87
L93:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v601))) = v603
	v606 = v41 + int32(20)
	v607 = int32(4)
	v608 = v601 + v607
	v610 = v603 << (uint(v607) % 32)
	if v610 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	base.MemoryCopy(m, v608, v606, v610)
	goto L96
L95:
	;
	goto L96
L96:
	;
	v616 = l4 + l6*int32(28) + int32(916)
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v431)+uint32(_c_F_gistSplitByKey[0])))
	v620 = F_FunctionCall2Coll(m, v616, v617, v601, v33+int32(96))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v620
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v601))) = v623
	v626 = v623 << (uint(int32(4)) % 32)
	if v626 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	base.MemoryCopy(m, v608, v606+v627<<(uint(int32(4))%32), v626)
	goto L100
L99:
	;
	goto L100
L100:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v431)+uint32(_c_F_gistSplitByKey[0])))
	v635 = F_FunctionCall2Coll(m, v616, v632, v601, v33+int32(96))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v635
	goto L64
L102:
	;
	v647 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41))))
	v649 = v647 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v643))) = uint16(v649)
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v652 = v651
	goto L104
L103:
	;
	v652 = v440
	goto L104
L104:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v658 = v653 + v652<<(uint(int32(1))%32) - int32(2)
	v659 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v658))))
	if v659 != 0 {
		goto L64
	} else {
		goto L105
	}
L105:
	;
	v660 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41))))
	v662 = v660 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v658))) = uint16(v662)
	goto L64
L106:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v418))) = v860
	v862 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v423))) = v862
	v864 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v402))) = uint8(v864)
	*(*uint8)(unsafe.Add(mBase, uint32(v409))) = uint8(v864)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+352)) = v864
	v871 = l6 + int32(1)
	v872 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v872)))
	if v873 <= v871 {
		goto L6
	} else {
		goto L137
	}
L107:
	;
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699))))
	if v703 != int32(1) {
		goto L106
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v708 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+110)) = uint8(v708)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+108)) = uint16(v708)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v707
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+78)) = uint8(v708)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+76)) = uint16(v708)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+72)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v33)+68)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = v706
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v697)))
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+62)) = uint8(v708)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+60)) = uint16(v708)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+56)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v724
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v695)))
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+42)) = uint8(v708)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+40)) = uint16(v708)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+36)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = v733
	if v700 != 0 {
		goto L114
	} else {
		goto L115
	}
L110:
	;
	goto L109
L111:
	;
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	if v829 == int32(1) {
		goto L129
	} else {
		goto L130
	}
L112:
	;
	v798 = *(*int64)(unsafe.Add(mBase, uint32(l5)+16))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = v799
	v801 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v798
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v801
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v805 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v804
	v808 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+62)) = uint8(v808)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+60)) = uint16(v808)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+56)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v805
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+42)) = uint8(v808)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+40)) = uint16(v808)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+36)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = v804
	goto L111
L113:
	;
	v766 = v33 + int32(96)
	v767 = int32(0)
	v769 = v33 + int32(48)
	v771 = F_gistpenalty(m, l4, l6, v766, v767, v769, v767)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L124
	}
L114:
	;
	v742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699))))
	if v742 == int32(1) {
		goto L113
	} else {
		goto L117
	}
L115:
	;
	v749 = v33 - int32(-64)
	goto L116
L116:
	;
	v750 = int32(0)
	v754 = F_gistpenalty(m, l4, l6, v749, v750, v33+int32(48), v750)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L118
	}
L117:
	;
	v749 = v33 + int32(96)
	goto L116
L118:
	;
	v756 = int32(0)
	v760 = F_gistpenalty(m, l4, l6, v749, v756, v33+int32(28), v756)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if base.F32_lt(v754, v760) != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v763 = v438
	goto L122
L121:
	;
	v763 = v699
	goto L122
L122:
	;
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v763))))
	if v764 != 0 {
		goto L111
	} else {
		goto L123
	}
L123:
	;
	goto L112
L124:
	;
	v774 = v33 - int32(-64)
	v775 = int32(0)
	v777 = v33 + int32(28)
	v779 = F_gistpenalty(m, l4, l6, v774, v775, v777, v775)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v782 = int32(0)
	v784 = F_gistpenalty(m, l4, l6, v766, v782, v777, v782)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v786 = int32(0)
	v788 = F_gistpenalty(m, l4, l6, v774, v786, v769, v786)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	if base.F32_gt(base.F32_add(v771, v779), base.F32_add(v784, v788)) == int32(0) {
		goto L111
	} else {
		goto L128
	}
L128:
	;
	goto L112
L129:
	;
	F_gistMakeUnionKey(m, l4, l6, v33+int32(96), v33+int32(48), v697, v33+int32(47))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699))))
	if v840 == int32(1) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	goto L131
L133:
	;
	F_gistMakeUnionKey(m, l4, l6, v33-int32(-64), v33+int32(28), v695, v33+int32(47))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L1
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v851 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v699))) = uint8(v851)
	*(*uint8)(unsafe.Add(mBase, uint32(v438))) = uint8(v851)
	goto L106
L136:
	;
	goto L135
L137:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v697)))
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v695)))
	v877 = m.G0
	v879 = v877 - int32(16)
	m.G0 = v879
	v881 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v879)+15)) = uint8(v881)
	v893 = *(*int32)(unsafe.Add(mBase, uint32(l4+l6<<(uint(int32(2))%32))+uint32(_c_F_gistSplitByKey[0])))
	v896 = F_FunctionCall3Coll(m, l4+l6*int32(28)+int32(_a_F_gistSplitByKey_9), v893, v875, v876, v879+int32(15))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L1
	} else {
		goto L139
	}
L138:
	;
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v2095)))
	v2120 = v2118 + l5
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v2120)))
	v2122 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2120))) = v2121 + v2122
	*(*uint16)(unsafe.Add(mBase, uint32(v2119+v2121<<(uint(v2122)%32)))) = uint16(v1528)
	goto L6
L139:
	;
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v879)+15)))
	m.G0 = v879 + int32(16)
	if v898 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v2095 = l5
	v2118 = int32(4)
	goto L138
L141:
	;
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	if v1678 == int32(0) {
		goto L237
	} else {
		goto L238
	}
L142:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v905 = F_palloc0(m, v902+int32(1))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+352)) = v905
	v908 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v909 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+110)) = uint8(v909)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+108)) = uint16(v909)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v909
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v908
	v918 = v41 + int32(4)
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v909 < v920 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v933 = int32(0)
	v935 = v909
	goto L147
L145:
	;
	v992 = v909
	goto L146
L146:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v1012 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+110)) = uint8(v1012)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+108)) = uint16(v1012)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1012
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1011
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	if v1012 < v1020 {
		goto L154
	} else {
		goto L155
	}
L147:
	;
	v956 = int32(0)
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v961 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v957+v933<<(uint(int32(1))%32)))))
	v966 = F_gistpenalty(m, l4, l6, v33+int32(96), v956, v918+v961<<(uint(int32(4))%32), v956)
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L1
	} else {
		goto L149
	}
L148:
	;
	v992 = v976
	goto L146
L149:
	;
	if base.F32_eq(v966, float32(0)) != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v972 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v970+v961))) = uint8(v972)
	v976 = v935 + v972
	goto L152
L151:
	;
	v976 = v935
	goto L152
L152:
	;
	v978 = v933 + int32(1)
	v979 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v978 < v979 {
		v933 = v978
		v935 = v976
		goto L147
	} else {
		goto L153
	}
L153:
	;
	goto L148
L154:
	;
	v1033 = int32(0)
	v1035 = v992
	goto L157
L155:
	;
	v1088 = v1020
	v1092 = v992
	goto L156
L156:
	;
	if v1092 <= int32(0) {
		goto L6
	} else {
		goto L164
	}
L157:
	;
	v1056 = int32(0)
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v1061 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1057+v1033<<(uint(int32(1))%32)))))
	v1066 = F_gistpenalty(m, l4, l6, v33+int32(96), v1056, v918+v1061<<(uint(int32(4))%32), v1056)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L1
	} else {
		goto L159
	}
L158:
	;
	v1088 = v1079
	v1092 = v1076
	goto L156
L159:
	;
	if base.F32_eq(v1066, float32(0)) != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v1072 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1070+v1061))) = uint8(v1072)
	v1076 = v1035 + v1072
	goto L162
L161:
	;
	v1076 = v1035
	goto L162
L162:
	;
	v1078 = v1033 + int32(1)
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	if v1078 < v1079 {
		v1033 = v1078
		v1035 = v1076
		goto L157
	} else {
		goto L163
	}
L163:
	;
	goto L158
L164:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if int32(0) < v1114 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v1118 = int32(0)
	if v1114 == int32(1) {
		goto L170
	} else {
		goto L171
	}
L166:
	;
	v1266 = v1088
	v1268 = v1113
	v1269 = v1114
	goto L167
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v1269
	if int32(0) < v1266 {
		goto L188
	} else {
		goto L189
	}
L168:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v1266 = v1258
	v1268 = v1257
	v1269 = v1237
	goto L167
L169:
	;
	v1221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1117+v1200<<(uint(int32(1))%32)))))
	v1223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1113+v1221))))
	if v1223 != 0 {
		goto L185
	} else {
		goto L186
	}
L170:
	;
	v1195 = v1117
	v1198 = v1114
	v1200 = v1118
	goto L169
L171:
	;
	goto L172
L172:
	;
	v1133 = v1117
	v1136 = v1114
	v1138 = v1118
	v1145 = int32(0)
	goto L173
L173:
	;
	v1158 = v1117 + v1138<<(uint(int32(1))%32)
	v1159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1158))))
	v1161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1113+v1159))))
	if v1161 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	if v1114&int32(1) == int32(0) {
		v1237 = v1180
		goto L168
	} else {
		goto L184
	}
L175:
	;
	v1171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1158)+2)))
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1113+v1171))))
	if v1173 != 0 {
		goto L180
	} else {
		goto L181
	}
L176:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1133))) = uint16(v1159)
	v1169 = v1133 + int32(2)
	v1170 = v1136
	goto L175
L177:
	;
	goto L178
L178:
	;
	v1169 = v1133
	v1170 = v1136 - int32(1)
	goto L175
L179:
	;
	v1181 = int32(2)
	v1182 = v1138 + v1181
	v1184 = v1145 + v1181
	if v1184 != v1114&int32(2147483646) {
		v1133 = v1179
		v1136 = v1180
		v1138 = v1182
		v1145 = v1184
		goto L173
	} else {
		goto L183
	}
L180:
	;
	v1179 = v1169
	v1180 = v1170 - int32(1)
	goto L179
L181:
	;
	goto L182
L182:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1169))) = uint16(v1171)
	v1179 = v1169 + int32(2)
	v1180 = v1170
	goto L179
L183:
	;
	goto L174
L184:
	;
	v1195 = v1179
	v1198 = v1180
	v1200 = v1182
	goto L169
L185:
	;
	v1237 = v1198 - int32(1)
	goto L168
L186:
	;
	goto L187
L187:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1195))) = uint16(v1221)
	v1237 = v1198
	goto L168
L188:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v1293 = int32(0)
	if v1266 == int32(1) {
		goto L193
	} else {
		goto L194
	}
L189:
	;
	v1440 = v1266
	v1443 = v1269
	goto L190
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v1440
	if v1440 != 0 {
		goto L211
	} else {
		goto L212
	}
L191:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v1440 = v1409
	v1443 = v1432
	goto L190
L192:
	;
	v1396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1292+v1375<<(uint(int32(1))%32)))))
	v1398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1268+v1396))))
	if v1398 != 0 {
		goto L208
	} else {
		goto L209
	}
L193:
	;
	v1370 = v1266
	v1373 = v1292
	v1375 = v1293
	goto L192
L194:
	;
	goto L195
L195:
	;
	v1308 = v1266
	v1311 = v1292
	v1313 = v1293
	v1320 = int32(0)
	goto L196
L196:
	;
	v1333 = v1292 + v1313<<(uint(int32(1))%32)
	v1334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1333))))
	v1336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1268+v1334))))
	if v1336 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	if v1266&int32(1) == int32(0) {
		v1409 = v1354
		goto L191
	} else {
		goto L207
	}
L198:
	;
	v1346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1333)+2)))
	v1348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1268+v1346))))
	if v1348 != 0 {
		goto L203
	} else {
		goto L204
	}
L199:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1311))) = uint16(v1334)
	v1344 = v1308
	v1345 = v1311 + int32(2)
	goto L198
L200:
	;
	goto L201
L201:
	;
	v1344 = v1308 - int32(1)
	v1345 = v1311
	goto L198
L202:
	;
	v1356 = int32(2)
	v1357 = v1313 + v1356
	v1359 = v1320 + v1356
	if v1359 != v1266&int32(2147483646) {
		v1308 = v1354
		v1311 = v1355
		v1313 = v1357
		v1320 = v1359
		goto L196
	} else {
		goto L206
	}
L203:
	;
	v1354 = v1344 - int32(1)
	v1355 = v1345
	goto L202
L204:
	;
	goto L205
L205:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1345))) = uint16(v1346)
	v1354 = v1344
	v1355 = v1345 + int32(2)
	goto L202
L206:
	;
	goto L197
L207:
	;
	v1370 = v1354
	v1373 = v1355
	v1375 = v1357
	goto L192
L208:
	;
	v1409 = v1370 - int32(1)
	goto L191
L209:
	;
	goto L210
L210:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1373))) = uint16(v1396)
	v1409 = v1370
	goto L191
L211:
	;
	v1465 = v1443
	goto L213
L212:
	;
	v1465 = int32(0)
	goto L213
L213:
	;
	if v1465 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+352)) = int32(0)
	F_gistSplitByKey(m, l0, l1, l2, l3, l4, l5, v871)
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L1
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	F_gistunionsubkey(m, l4, l2, l5)
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L1
	} else {
		goto L218
	}
L217:
	;
	goto L6
L218:
	;
	v1474 = int32(1)
	if v1092 != v1474 {
		goto L141
	} else {
		goto L219
	}
L219:
	;
	v1477 = int32(1)
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v1478 < int32(2) {
		v1528 = v1474
		v1530 = v1477
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1530<<(uint(int32(2))%32)-int32(4))))
	F_gistDeCompressAtt(m, l4, l0, v1554, v33+int32(96), v33-int32(-64))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L1
	} else {
		goto L226
	}
L221:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v1491 = v1474
	v1493 = v1477
	goto L222
L222:
	;
	v1513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1481+v1493))))
	if v1513 != 0 {
		v1528 = v1491
		v1530 = v1493
		goto L220
	} else {
		goto L224
	}
L223:
	;
	v1528 = v1515
	v1530 = v1517
	goto L220
L224:
	;
	v1515 = v1491 + int32(1)
	v1517 = v1515 & int32(_a_F_gistSplitByKey_1)
	if base.Ui32(v1517) < base.Ui32(v1478) {
		v1491 = v1515
		v1493 = v1517
		goto L222
	} else {
		goto L225
	}
L225:
	;
	goto L223
L226:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1561)))
	if v1562 <= v871 {
		goto L140
	} else {
		goto L227
	}
L227:
	;
	v1580 = v871
	goto L228
L228:
	;
	v1597 = v1580 << (uint(int32(2)) % 32)
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(v415+v1597)))
	v1600 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+62)) = uint8(v1600)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+60)) = uint16(v1600)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+56)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v1599
	v1609 = v33 + int32(48)
	v1611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1580+v401))))
	v1616 = v33 + int32(96) + v1580<<(uint(int32(4))%32)
	v1619 = v33 - int32(-64) + v1580
	v1620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1619))))
	v1621 = F_gistpenalty(m, l4, v1580, v1609, v1611, v1616, v1620)
	mBase = m.M
	v1622 = m.ExcPending
	if v1622 != 0 {
		goto L1
	} else {
		goto L230
	}
L229:
	;
	goto L140
L230:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1597+v422)))
	v1625 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+62)) = uint8(v1625)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+60)) = uint16(v1625)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+56)) = v1625
	*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v1624
	v1634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1580+v408))))
	v1635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1619))))
	v1636 = F_gistpenalty(m, l4, v1580, v1609, v1634, v1616, v1635)
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	if base.F32_ne(v1636, v1621) != 0 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	if base.F32_gt(v1621, v1636) == int32(0) {
		goto L140
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v1644 = v1580 + int32(1)
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1645)))
	if v1644 < v1646 {
		v1580 = v1644
		goto L228
	} else {
		goto L236
	}
L235:
	;
	v2095 = l5 + int32(16)
	v2118 = int32(20)
	goto L138
L236:
	;
	goto L229
L237:
	;
	F_gistSplitByKey(m, l0, l1, l2, l3, l4, l5, v871)
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L1
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v1685 = F_palloc(m, l3<<(uint(int32(2))%32))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L1
	} else {
		goto L241
	}
L240:
	;
	goto L6
L241:
	;
	v1687 = F_palloc(m, v45)
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	if l3 <= int32(0) {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	v1861 = *(*int64)(unsafe.Add(mBase, uint32(l5)+24))
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v1863 = *(*int64)(unsafe.Add(mBase, uint32(l5)+8))
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v1865 = F_palloc(m, v45)
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L1
	} else {
		goto L261
	}
L244:
	;
	v1843 = int32(0)
	goto L243
L245:
	;
	goto L246
L246:
	;
	v1692 = int32(0)
	if l3 != int32(1) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1710 = v1692
	v1712 = int32(0)
	v1713 = v1692
	goto L250
L248:
	;
	v1789 = v1692
	v1792 = v1692
	goto L249
L249:
	;
	v1810 = int32(1)
	v1811 = v1789 + v1810
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v1814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1811+v1812))))
	if v1814 != v1810 {
		v1843 = v1792
		goto L243
	} else {
		goto L260
	}
L250:
	;
	v1731 = int32(1)
	v1732 = v1710 | v1731
	v1733 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v1735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1732+v1733))))
	if v1735 == v1731 {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	if l3&int32(1) == int32(0) {
		v1843 = v1774
		goto L243
	} else {
		goto L259
	}
L252:
	;
	v1738 = int32(2)
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1710<<(uint(v1738)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1685+v1713<<(uint(v1738)%32)))) = v1744
	v1746 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1687+v1713<<(uint(v1746)%32)))) = uint16(v1732)
	v1752 = v1713 + v1746
	goto L254
L253:
	;
	v1752 = v1713
	goto L254
L254:
	;
	v1754 = v1710 + int32(2)
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v1757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1754+v1755))))
	if v1757 == int32(1) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1760 = int32(2)
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1732<<(uint(v1760)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1685+v1752<<(uint(v1760)%32)))) = v1766
	v1768 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1687+v1752<<(uint(v1768)%32)))) = uint16(v1754)
	v1774 = v1752 + v1768
	goto L257
L256:
	;
	v1774 = v1752
	goto L257
L257:
	;
	v1776 = v1712 + int32(2)
	if v1776 != l3&int32(2147483646) {
		v1710 = v1754
		v1712 = v1776
		v1713 = v1774
		goto L250
	} else {
		goto L258
	}
L258:
	;
	goto L251
L259:
	;
	v1789 = v1754
	v1792 = v1774
	goto L249
L260:
	;
	v1817 = int32(2)
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1789<<(uint(v1817)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1685+v1792<<(uint(v1817)%32)))) = v1823
	v1825 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1687+v1792<<(uint(v1825)%32)))) = uint16(v1811)
	v1843 = v1792 + v1825
	goto L243
L261:
	;
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v1869 = v1867 << (uint(int32(1)) % 32)
	if v1869 != 0 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	base.MemoryCopy(m, v1865, v1870, v1869)
	goto L264
L263:
	;
	goto L264
L264:
	;
	v1872 = F_palloc(m, v45)
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v1876 = v1874 << (uint(int32(1)) % 32)
	if v1876 != 0 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	base.MemoryCopy(m, v1872, v1877, v1876)
	goto L268
L267:
	;
	goto L268
L268:
	;
	F_gistSplitByKey(m, l0, l1, v1685, v1843, l4, l5, v871)
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	v1881 = int32(0)
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v1881 < v1882 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1896 = v1864
	v1897 = v1881
	goto L273
L271:
	;
	v1947 = v1864
	goto L272
L272:
	;
	v1966 = int32(0)
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	if v1966 < v1967 {
		goto L276
	} else {
		goto L277
	}
L273:
	;
	v1915 = int32(1)
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v1922 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1918+v1897<<(uint(v1915)%32)))))
	v1928 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1687+v1922<<(uint(v1915)%32)-int32(2)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1865+v1896<<(uint(v1915)%32)))) = uint16(v1928)
	v1931 = v1896 + v1915
	v1933 = v1897 + v1915
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v1933 < v1934 {
		v1896 = v1931
		v1897 = v1933
		goto L273
	} else {
		goto L275
	}
L274:
	;
	v1947 = v1931
	goto L272
L275:
	;
	goto L274
L276:
	;
	v1979 = v1862
	v1982 = v1966
	goto L279
L277:
	;
	v2030 = v1862
	goto L278
L278:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+24)) = v1861
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v2030
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = v1872
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = v1863
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v1947
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v1865
	goto L6
L279:
	;
	v2000 = int32(1)
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v2007 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2003+v1982<<(uint(v2000)%32)))))
	v2013 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1687+v2007<<(uint(v2000)%32)-int32(2)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1872+v1979<<(uint(v2000)%32)))) = uint16(v2013)
	v2016 = v1979 + v2000
	v2018 = v1982 + v2000
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	if v2018 < v2019 {
		v1979 = v2016
		v1982 = v2018
		goto L279
	} else {
		goto L281
	}
L280:
	;
	v2030 = v2016
	goto L278
L281:
	;
	goto L280
L282:
	;
	goto L5
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+352)) = int32(0)
	F_gistunionsubkey(m, l4, l2, l5)
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	goto L4
}
func F_gist_bbox_zorder_abbrev_convert(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v8 int64
	_ = v8
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
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
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v69 int64
	_ = v69
	var v74 int64
	_ = v74
	var v79 int64
	_ = v79
	var v84 int64
	_ = v84
	v6 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = int64(4294967295)
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = base.I32_reinterpret_f32(base.F32_demote_f64(v10))
	if base.Ui32(v12&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
		if v12 < int32(0) {
			v21 = int32(-1)
		} else {
			v21 = int32(-2147483648)
		}
		v24 = base.I64_extend_i32_u(v12 ^ v21)
	} else {
		v24 = v8
	}
	v25 = base.I32_reinterpret_f32(base.F32_demote_f64(v6))
	if base.Ui32(v25&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
		if v25 < int32(0) {
			v34 = int32(-1)
		} else {
			v34 = int32(-2147483648)
		}
		v37 = base.I64_extend_i32_u(v25 ^ v34)
	} else {
		v37 = v8
	}
	v38 = int64(16)
	v41 = int64(281470681808895)
	v42 = (v37<<(uint(v38)%64) | v37) & v41
	v43 = int64(8)
	v46 = int64(71777214294589695)
	v47 = (v42<<(uint(v43)%64) | v42) & v46
	v48 = int64(4)
	v51 = int64(1085102592571150095)
	v52 = (v47<<(uint(v48)%64) | v47) & v51
	v53 = int64(2)
	v56 = int64(3689348814741910323)
	v57 = (v52<<(uint(v53)%64) | v52) & v56
	v60 = int64(1)
	v69 = (v24<<(uint(v38)%64) | v24) & v41
	v74 = (v69<<(uint(v43)%64) | v69) & v46
	v79 = (v74<<(uint(v48)%64) | v74) & v51
	v84 = (v79<<(uint(v53)%64) | v79) & v56
	return base.I32_wrap_i64(int64(base.Ui64((v57<<(uint(v53)%64)|v57<<(uint(v60)%64))&int64(-6148914694099828736)|(v84<<(uint(v60)%64)|v84)&int64(6148914689804861440)) >> (uint(int64(32)) % 64)))
}
func F_gist_box_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = F_box_penalty(m, v5, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*float32)(unsafe.Add(mBase, uint32(v3))) = base.F32_demote_f64(v8)
		return v3
	}
}
func F_gist_box_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v90 int32
	_ = v90
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v101 float64
	_ = v101
	var v107 float64
	_ = v107
	var v119 float64
	_ = v119
	var v125 float64
	_ = v125
	var v131 float64
	_ = v131
	var v133 float64
	_ = v133
	var v136 float64
	_ = v136
	var v142 float64
	_ = v142
	var v154 float64
	_ = v154
	var v160 float64
	_ = v160
	var v166 float64
	_ = v166
	var v168 float64
	_ = v168
	var v175 int32
	_ = v175
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v219 int32
	_ = v219
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 float64
	_ = v271
	var v273 float64
	_ = v273
	var v277 int32
	_ = v277
	var v280 float64
	_ = v280
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 float64
	_ = v296
	var v297 float64
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 float64
	_ = v301
	var v302 float64
	_ = v302
	var v306 int32
	_ = v306
	var v328 int32
	_ = v328
	var v330 float64
	_ = v330
	var v355 int32
	_ = v355
	var v356 float64
	_ = v356
	var v359 int64
	_ = v359
	var v373 float64
	_ = v373
	var v379 float64
	_ = v379
	var v381 float64
	_ = v381
	var v383 float64
	_ = v383
	var v385 int32
	_ = v385
	var v397 int32
	_ = v397
	var v421 float64
	_ = v421
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v463 int32
	_ = v463
	var v490 float64
	_ = v490
	var v491 float64
	_ = v491
	var v492 int32
	_ = v492
	var v493 float64
	_ = v493
	var v494 float64
	_ = v494
	var v501 int32
	_ = v501
	var v520 int32
	_ = v520
	var v521 float64
	_ = v521
	var v547 int32
	_ = v547
	var v548 float64
	_ = v548
	var v551 int64
	_ = v551
	var v560 float64
	_ = v560
	var v567 float64
	_ = v567
	var v573 float64
	_ = v573
	var v574 float64
	_ = v574
	var v590 int32
	_ = v590
	var v612 float64
	_ = v612
	var v639 int32
	_ = v639
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int64
	_ = v774
	var v776 int64
	_ = v776
	var v778 int64
	_ = v778
	var v780 int64
	_ = v780
	var v782 float64
	_ = v782
	var v788 float64
	_ = v788
	var v800 float64
	_ = v800
	var v806 float64
	_ = v806
	var v818 float64
	_ = v818
	var v824 float64
	_ = v824
	var v836 float64
	_ = v836
	var v842 float64
	_ = v842
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int64
	_ = v871
	var v873 int64
	_ = v873
	var v875 int64
	_ = v875
	var v877 int64
	_ = v877
	var v879 float64
	_ = v879
	var v885 float64
	_ = v885
	var v897 float64
	_ = v897
	var v903 float64
	_ = v903
	var v915 float64
	_ = v915
	var v921 float64
	_ = v921
	var v933 float64
	_ = v933
	var v939 float64
	_ = v939
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v972 int32
	_ = v972
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1015 float64
	_ = v1015
	var v1017 int64
	_ = v1017
	var v1019 float64
	_ = v1019
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1060 int32
	_ = v1060
	var v1062 float64
	_ = v1062
	var v1066 float64
	_ = v1066
	var v1081 int32
	_ = v1081
	var v1097 int32
	_ = v1097
	var v1100 float64
	_ = v1100
	var v1106 float64
	_ = v1106
	var v1118 float64
	_ = v1118
	var v1124 float64
	_ = v1124
	var v1136 float64
	_ = v1136
	var v1142 float64
	_ = v1142
	var v1154 float64
	_ = v1154
	var v1160 float64
	_ = v1160
	var v1171 int64
	_ = v1171
	var v1173 int64
	_ = v1173
	var v1175 int64
	_ = v1175
	var v1177 int64
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1190 int32
	_ = v1190
	var v1193 float64
	_ = v1193
	var v1199 float64
	_ = v1199
	var v1211 float64
	_ = v1211
	var v1217 float64
	_ = v1217
	var v1229 float64
	_ = v1229
	var v1235 float64
	_ = v1235
	var v1247 float64
	_ = v1247
	var v1253 float64
	_ = v1253
	var v1264 int64
	_ = v1264
	var v1266 int64
	_ = v1266
	var v1268 int64
	_ = v1268
	var v1270 int64
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1336 float64
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 float64
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1341 float64
	_ = v1341
	var v1342 float64
	_ = v1342
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1405 float64
	_ = v1405
	var v1411 float64
	_ = v1411
	var v1423 float64
	_ = v1423
	var v1429 float64
	_ = v1429
	var v1441 float64
	_ = v1441
	var v1447 float64
	_ = v1447
	var v1459 float64
	_ = v1459
	var v1465 float64
	_ = v1465
	var v1476 int64
	_ = v1476
	var v1478 int64
	_ = v1478
	var v1480 int64
	_ = v1480
	var v1482 int64
	_ = v1482
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1496 int32
	_ = v1496
	var v1501 float64
	_ = v1501
	var v1507 float64
	_ = v1507
	var v1519 float64
	_ = v1519
	var v1525 float64
	_ = v1525
	var v1537 float64
	_ = v1537
	var v1543 float64
	_ = v1543
	var v1555 float64
	_ = v1555
	var v1561 float64
	_ = v1561
	var v1572 int64
	_ = v1572
	var v1574 int64
	_ = v1574
	var v1576 int64
	_ = v1576
	var v1578 int64
	_ = v1578
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1592 float64
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1594 float64
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1597 int32
	_ = v1597
	var v1600 float64
	_ = v1600
	var v1606 float64
	_ = v1606
	var v1618 float64
	_ = v1618
	var v1624 float64
	_ = v1624
	var v1636 float64
	_ = v1636
	var v1642 float64
	_ = v1642
	var v1654 float64
	_ = v1654
	var v1660 float64
	_ = v1660
	var v1671 int64
	_ = v1671
	var v1673 int64
	_ = v1673
	var v1675 int64
	_ = v1675
	var v1677 int64
	_ = v1677
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1686 int32
	_ = v1686
	var v1691 int32
	_ = v1691
	var v1694 float64
	_ = v1694
	var v1700 float64
	_ = v1700
	var v1712 float64
	_ = v1712
	var v1718 float64
	_ = v1718
	var v1730 float64
	_ = v1730
	var v1736 float64
	_ = v1736
	var v1748 float64
	_ = v1748
	var v1754 float64
	_ = v1754
	var v1765 int64
	_ = v1765
	var v1767 int64
	_ = v1767
	var v1769 int64
	_ = v1769
	var v1771 int64
	_ = v1771
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1780 int32
	_ = v1780
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1820 int32
	_ = v1820
	var v1851 int32
	_ = v1851
	v6 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(96)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryFill(m, v28+int32(8), v6, int32(88))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v38 = int32(_a_F_gist_box_picksplit_0)
	v39 = v37 + v38
	v41 = v39 & v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v41
	v44 = v41 - int32(1)
	v46 = v41 << (uint(int32(4)) % 32)
	v47 = F_palloc(m, v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v51 = F_palloc(m, v46)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v54 = v37 & int32(_a_F_gist_box_picksplit_0)
	if v54 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v60 = v28 + int32(16)
	v68 = int32(1)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v202 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+48)) = uint8(v202)
	v204 = int32(4)
	v205 = v31 + v204
	v207 = v44 << (uint(v204) % 32)
	v219 = v202
	v230 = v6
	goto L36
L7:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(4)+v68<<(uint(int32(4))%32))))
	if v68 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v175 = (v68 + int32(1)) & int32(_a_F_gist_box_picksplit_0)
	if base.Ui32(v175) <= base.Ui32(v41) {
		v68 = v175
		goto L7
	} else {
		goto L35
	}
L10:
	;
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v90)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v93
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v90)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v60)+16)) = v95
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v90)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v60)+8)) = v97
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v90)))
	*(*int64)(unsafe.Add(mBase, uint32(v60))) = v99
	goto L9
L11:
	;
	goto L12
L12:
	;
	v101 = *(*float64)(unsafe.Add(mBase, uint32(v28)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v101)&int64(9223372036854775807)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v119 = *(*float64)(unsafe.Add(mBase, uint32(v90)+16))
	if base.Ui64(base.I64_reinterpret_f64(v119)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v107 = *(*float64)(unsafe.Add(mBase, uint32(v90)))
	if base.B2i32(base.F64_lt(v101, v107) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v107)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v28)+16)) = v107
	goto L13
L16:
	;
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v28)+32))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v125)&int64(9223372036854775807)) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v136 = *(*float64)(unsafe.Add(mBase, uint32(v28)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v136)&int64(9223372036854775807)) {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v131 = v119
	goto L21
L20:
	;
	v131 = v125
	goto L21
L21:
	;
	if base.F64_lt(v119, v125) != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v133 = v119
	goto L24
L23:
	;
	v133 = v131
	goto L24
L24:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v28)+32)) = v133
	goto L18
L25:
	;
	v154 = *(*float64)(unsafe.Add(mBase, uint32(v90)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v154)&int64(9223372036854775807)) {
		goto L9
	} else {
		goto L28
	}
L26:
	;
	v142 = *(*float64)(unsafe.Add(mBase, uint32(v90)+8))
	if base.B2i32(base.F64_lt(v136, v142) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v142)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v28)+24)) = v142
	goto L25
L28:
	;
	v160 = *(*float64)(unsafe.Add(mBase, uint32(v28)+40))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v160)&int64(9223372036854775807)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v166 = v154
	goto L31
L30:
	;
	v166 = v160
	goto L31
L31:
	;
	if base.F64_lt(v154, v160) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v168 = v154
	goto L34
L33:
	;
	v168 = v166
	goto L34
L34:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v28)+40)) = v168
	goto L9
L35:
	;
	goto L8
L36:
	;
	v236 = int32(1)
	if v54 != v236 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+48)))
	if v701 == int32(1) {
		goto L126
	} else {
		goto L127
	}
L38:
	;
	if v219 != 0 {
		v219 = int32(0)
		v230 = int32(1)
		goto L36
	} else {
		goto L123
	}
L39:
	;
	v245 = v236
	goto L42
L40:
	;
	goto L41
L41:
	;
	if v46 != 0 {
		goto L118
	} else {
		goto L119
	}
L42:
	;
	v265 = v245 << (uint(int32(4)) % 32)
	v266 = v47 + v265
	v268 = v266 - int32(16)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v265+v205)))
	if v219 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	if v46 != 0 {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v277)))
	*(*float64)(unsafe.Add(mBase, uint32(v266-int32(8)))) = v280
	v285 = (v245 + int32(1)) & int32(_a_F_gist_box_picksplit_0)
	if base.Ui32(v285) <= base.Ui32(v41) {
		v245 = v285
		goto L42
	} else {
		goto L48
	}
L45:
	;
	v271 = *(*float64)(unsafe.Add(mBase, uint32(v270)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v268))) = v271
	v277 = v270
	goto L44
L46:
	;
	goto L47
L47:
	;
	v273 = *(*float64)(unsafe.Add(mBase, uint32(v270)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v268))) = v273
	v277 = v270 + int32(8)
	goto L44
L48:
	;
	goto L43
L49:
	;
	base.MemoryCopy(m, v51, v47, v46)
	goto L51
L50:
	;
	goto L51
L51:
	;
	F_pg_qsort(m, v47, v41, int32(16), int32(105))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_pg_qsort(m, v51, v41, int32(16), int32(106))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v296 = *(*float64)(unsafe.Add(mBase, uint32(v47)))
	v297 = *(*float64)(unsafe.Add(mBase, uint32(v51)))
	v298 = int32(0)
	v300 = v298
	v301 = v296
	v302 = v297
	v306 = v298
	goto L54
L54:
	;
	v328 = v300
	v330 = v302
	goto L57
L55:
	;
	v490 = *(*float64)(unsafe.Add(mBase, uint32(v207+v47)+8))
	v491 = *(*float64)(unsafe.Add(mBase, uint32(v51+v207)+8))
	v492 = v44
	v493 = v490
	v494 = v491
	v501 = v44
	goto L87
L56:
	;
	goto L55
L57:
	;
	v355 = v47 + v328<<(uint(int32(4))%32)
	v356 = *(*float64)(unsafe.Add(mBase, uint32(v355)))
	v359 = base.I64_reinterpret_f64(v356) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v301)&int64(9223372036854775807)) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	if v41 <= v306 {
		v441 = v306
		goto L76
	} else {
		goto L77
	}
L59:
	;
	goto L58
L60:
	;
	if base.Ui64(base.I64_reinterpret_f64(v330)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L66
	} else {
		goto L67
	}
L61:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v359) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v359))|base.F64_ne(v356, v301) != 0 {
		goto L59
	} else {
		goto L65
	}
L64:
	;
	goto L59
L65:
	;
	goto L60
L66:
	;
	v373 = *(*float64)(unsafe.Add(mBase, uint32(v355)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v373)&int64(9223372036854775807)) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v383 = v330
	goto L68
L68:
	;
	v385 = v328 + int32(1)
	if v385 < v41 {
		v328 = v385
		v330 = v383
		goto L57
	} else {
		goto L75
	}
L69:
	;
	v379 = v373
	goto L71
L70:
	;
	v379 = v330
	goto L71
L71:
	;
	if base.F64_gt(v373, v330) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v381 = v373
	goto L74
L73:
	;
	v381 = v379
	goto L74
L74:
	;
	v383 = v381
	goto L68
L75:
	;
	goto L56
L76:
	;
	F_g_box_consider_split(m, v28+int32(8), v230, v356, v328, v330, v441)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L85
	}
L77:
	;
	v397 = v306
	goto L78
L78:
	;
	if base.Ui64(base.I64_reinterpret_f64(v330)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v441 = v41
	goto L76
L80:
	;
	v421 = *(*float64)(unsafe.Add(mBase, uint32(v51+v397<<(uint(int32(4))%32))+8))
	if base.B2i32(base.F64_le(v421, v330) == int32(0))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v421)&int64(9223372036854775807))) != 0 {
		v441 = v397
		goto L76
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v433 = v397 + int32(1)
	if v433 != v41 {
		v397 = v433
		goto L78
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	goto L79
L85:
	;
	if v328 < v41 {
		v300 = v328
		v301 = v356
		v302 = v330
		v306 = v441
		goto L54
	} else {
		goto L86
	}
L86:
	;
	goto L56
L87:
	;
	v520 = v492
	v521 = v493
	goto L89
L88:
	;
	goto L38
L89:
	;
	v547 = v51 + v520<<(uint(int32(4))%32)
	v548 = *(*float64)(unsafe.Add(mBase, uint32(v547)+8))
	v551 = base.I64_reinterpret_f64(v548) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v494)&int64(9223372036854775807)) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	if v501 < int32(0) {
		v639 = v501
		goto L108
	} else {
		goto L109
	}
L91:
	;
	goto L90
L92:
	;
	v560 = *(*float64)(unsafe.Add(mBase, uint32(v547)))
	if base.Ui64(base.I64_reinterpret_f64(v560)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L98
	} else {
		goto L99
	}
L93:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v551) {
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v551))|base.F64_ne(v548, v494) != 0 {
		goto L91
	} else {
		goto L97
	}
L96:
	;
	goto L91
L97:
	;
	goto L92
L98:
	;
	if base.F64_gt(v521, v560) != 0 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v574 = v521
	goto L100
L100:
	;
	if int32(0) < v520 {
		v520 = v520 - int32(1)
		v521 = v574
		goto L89
	} else {
		goto L107
	}
L101:
	;
	v567 = v560
	goto L103
L102:
	;
	v567 = v521
	goto L103
L103:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v521)&int64(9223372036854775807)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v573 = v560
	goto L106
L105:
	;
	v573 = v567
	goto L106
L106:
	;
	v574 = v573
	goto L100
L107:
	;
	goto L38
L108:
	;
	v657 = int32(1)
	F_g_box_consider_split(m, v28+int32(8), v230, v521, v639+v657, v548, v520+v657)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L116
	}
L109:
	;
	v590 = v501
	goto L110
L110:
	;
	v612 = *(*float64)(unsafe.Add(mBase, uint32(v47+v590<<(uint(int32(4))%32))))
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v612)&int64(9223372036854775807)))|base.B2i32(base.Ui64(base.I64_reinterpret_f64(v521)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))&base.F64_le(v521, v612) == int32(0) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v639 = int32(-1)
	goto L108
L112:
	;
	v639 = v590
	goto L108
L113:
	;
	goto L114
L114:
	;
	if int32(0) < v590 {
		v590 = v590 - int32(1)
		goto L110
	} else {
		goto L115
	}
L115:
	;
	goto L111
L116:
	;
	if int32(0) <= v520 {
		v492 = v520
		v493 = v521
		v494 = v548
		v501 = v639
		goto L87
	} else {
		goto L117
	}
L117:
	;
	goto L88
L118:
	;
	base.MemoryCopy(m, v51, v47, v46)
	goto L120
L119:
	;
	goto L120
L120:
	;
	F_pg_qsort(m, v47, v41, int32(16), int32(105))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_pg_qsort(m, v51, v41, int32(16), int32(106))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	goto L38
L123:
	;
	goto L37
L124:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L1
	} else {
		goto L311
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v1820
	m.G0 = v28 + int32(96)
	return v30
L126:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v705 = int32(_a_F_gist_box_picksplit_0)
	v708 = (v704 + v705) & v705
	v712 = v708<<(uint(int32(1))%32) + int32(4)
	v713 = F_palloc(m, v712)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v993 = v41 << (uint(int32(1)) % 32)
	v994 = F_palloc(m, v993)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L173
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v713
	v716 = int32(0)
	v717 = F_palloc(m, v712)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v719 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v719
	if v704&int32(_a_F_gist_box_picksplit_0) != int32(1) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v729 = int32(1)
	v733 = v716
	v738 = v729
	v739 = v719
	v740 = v729
	goto L134
L132:
	;
	v966 = v716
	v972 = v719
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v972
	v1820 = v966
	goto L125
L134:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v205+v738<<(uint(int32(4))%32))))
	if base.Ui32(v738) <= base.Ui32(int32(base.Ui32(v708)>>(uint(v729)%32))) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v966 = v957
	v972 = v960
	goto L133
L136:
	;
	v962 = v740 + int32(1)
	v964 = v962 & int32(_a_F_gist_box_picksplit_0)
	if base.Ui32(v964) <= base.Ui32(v708) {
		v733 = v957
		v738 = v964
		v739 = v960
		v740 = v962
		goto L134
	} else {
		goto L172
	}
L137:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v763+v764<<(uint(int32(1))%32)))) = uint16(v740)
	if v739 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	goto L139
L139:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v860+v861<<(uint(int32(1))%32)))) = uint16(v740)
	if v733 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L140:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v856 + int32(1)
	v957 = v733
	v960 = v855
	goto L136
L141:
	;
	v772 = F_palloc(m, int32(32))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v782 = *(*float64)(unsafe.Add(mBase, uint32(v739)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v782)&int64(9223372036854775807)) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v774 = *(*int64)(unsafe.Add(mBase, uint32(v761)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v772)+24)) = v774
	v776 = *(*int64)(unsafe.Add(mBase, uint32(v761)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v772)+16)) = v776
	v778 = *(*int64)(unsafe.Add(mBase, uint32(v761)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v772)+8)) = v778
	v780 = *(*int64)(unsafe.Add(mBase, uint32(v761)))
	*(*int64)(unsafe.Add(mBase, uint32(v772))) = v780
	v855 = v772
	goto L140
L145:
	;
	v800 = *(*float64)(unsafe.Add(mBase, uint32(v761)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v800)&int64(9223372036854775807)) {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v788 = *(*float64)(unsafe.Add(mBase, uint32(v761)))
	if base.B2i32(base.F64_lt(v782, v788) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v788)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v739))) = v788
	goto L145
L148:
	;
	v818 = *(*float64)(unsafe.Add(mBase, uint32(v739)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v818)&int64(9223372036854775807)) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v806 = *(*float64)(unsafe.Add(mBase, uint32(v739)+16))
	if base.B2i32(base.F64_gt(v806, v800) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v806)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v739)+16)) = v800
	goto L148
L151:
	;
	v836 = *(*float64)(unsafe.Add(mBase, uint32(v761)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v836)&int64(9223372036854775807)) {
		v855 = v739
		goto L140
	} else {
		goto L154
	}
L152:
	;
	v824 = *(*float64)(unsafe.Add(mBase, uint32(v761)+8))
	if base.B2i32(base.F64_lt(v818, v824) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v824)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v739)+8)) = v824
	goto L151
L154:
	;
	v842 = *(*float64)(unsafe.Add(mBase, uint32(v739)+24))
	if base.B2i32(base.F64_gt(v842, v836) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v842)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		v855 = v739
		goto L140
	} else {
		goto L155
	}
L155:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v739)+24)) = v836
	v855 = v739
	goto L140
L156:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v953 + int32(1)
	v957 = v950
	v960 = v739
	goto L136
L157:
	;
	v869 = F_palloc(m, int32(32))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v879 = *(*float64)(unsafe.Add(mBase, uint32(v733)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v879)&int64(9223372036854775807)) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v871 = *(*int64)(unsafe.Add(mBase, uint32(v761)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v869)+24)) = v871
	v873 = *(*int64)(unsafe.Add(mBase, uint32(v761)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v869)+16)) = v873
	v875 = *(*int64)(unsafe.Add(mBase, uint32(v761)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v869)+8)) = v875
	v877 = *(*int64)(unsafe.Add(mBase, uint32(v761)))
	*(*int64)(unsafe.Add(mBase, uint32(v869))) = v877
	v950 = v869
	goto L156
L161:
	;
	v897 = *(*float64)(unsafe.Add(mBase, uint32(v761)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v897)&int64(9223372036854775807)) {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v885 = *(*float64)(unsafe.Add(mBase, uint32(v761)))
	if base.B2i32(base.F64_lt(v879, v885) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v885)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v733))) = v885
	goto L161
L164:
	;
	v915 = *(*float64)(unsafe.Add(mBase, uint32(v733)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v915)&int64(9223372036854775807)) {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v903 = *(*float64)(unsafe.Add(mBase, uint32(v733)+16))
	if base.B2i32(base.F64_gt(v903, v897) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v903)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v733)+16)) = v897
	goto L164
L167:
	;
	v933 = *(*float64)(unsafe.Add(mBase, uint32(v761)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v933)&int64(9223372036854775807)) {
		v950 = v733
		goto L156
	} else {
		goto L170
	}
L168:
	;
	v921 = *(*float64)(unsafe.Add(mBase, uint32(v761)+8))
	if base.B2i32(base.F64_lt(v915, v921) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v921)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v733)+8)) = v921
	goto L167
L170:
	;
	v939 = *(*float64)(unsafe.Add(mBase, uint32(v733)+24))
	if base.B2i32(base.F64_gt(v939, v933) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v939)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		v950 = v733
		goto L156
	} else {
		goto L171
	}
L171:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v733)+24)) = v933
	v950 = v733
	goto L156
L172:
	;
	goto L135
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v994
	v998 = F_palloc(m, v993)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v1000 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v998
	v1006 = F_palloc0(m, int32(32))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v1009 = F_palloc0(m, int32(32))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v1011 = F_palloc(m, v46)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	if v54 == int32(1) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v1006
	v1820 = v1009
	goto L125
L179:
	;
	v1015 = *(*float64)(unsafe.Add(mBase, uint32(v28)+64))
	v1017 = int64(9223372036854775807)
	v1019 = *(*float64)(unsafe.Add(mBase, uint32(v28)+56))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v28)+80))
	if v1025 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v1026 = int32(24)
	goto L182
L181:
	;
	v1026 = int32(16)
	goto L182
L182:
	;
	if v1025 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v1029 = int32(8)
	goto L185
L184:
	;
	v1029 = int32(0)
	goto L185
L185:
	;
	v1030 = int32(1)
	v1038 = int32(0)
	v1041 = v1030
	v1044 = v1030
	goto L186
L186:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v205+v1044<<(uint(int32(4))%32))))
	v1062 = *(*float64)(unsafe.Add(mBase, uint32(v1060+v1026)))
	if base.Ui64(base.I64_reinterpret_f64(v1019)&v1017) <= base.Ui64(int64(9218868437227405312)) {
		goto L190
	} else {
		goto L191
	}
L187:
	;
	if v1285 <= int32(0) {
		goto L178
	} else {
		goto L228
	}
L188:
	;
	v1288 = v1041 + int32(1)
	v1289 = int32(_a_F_gist_box_picksplit_0)
	v1290 = v1288 & v1289
	if base.Ui32(v1290) <= base.Ui32(v39&v1289) {
		v1038 = v1285
		v1041 = v1288
		v1044 = v1290
		goto L186
	} else {
		goto L227
	}
L189:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	if int32(0) < v1190 {
		goto L213
	} else {
		goto L214
	}
L190:
	;
	v1066 = *(*float64)(unsafe.Add(mBase, uint32(v1060+v1029)))
	if base.B2i32(base.F64_le(v1066, v1019) == int32(0))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1066)&int64(9223372036854775807))) != 0 {
		goto L189
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v1081 = int32(0)
	if base.B2i32(base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1015)&v1017) < base.Ui64(int64(9218868437227405313)))&base.F64_ge(v1062, v1015) == v1081)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1062)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) == v1081 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	goto L192
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1011+v1038<<(uint(int32(4))%32)))) = v1044
	v1285 = v1038 + int32(1)
	goto L188
L195:
	;
	goto L196
L196:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if int32(0) < v1097 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v1182 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v1181 + v1182
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1185+v1181<<(uint(v1182)%32)))) = uint16(v1041)
	v1285 = v1038
	goto L188
L198:
	;
	v1100 = *(*float64)(unsafe.Add(mBase, uint32(v1006)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1100)&int64(9223372036854775807)) {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	goto L200
L200:
	;
	v1171 = *(*int64)(unsafe.Add(mBase, uint32(v1060)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1006)+24)) = v1171
	v1173 = *(*int64)(unsafe.Add(mBase, uint32(v1060)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1006)+16)) = v1173
	v1175 = *(*int64)(unsafe.Add(mBase, uint32(v1060)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1006)+8)) = v1175
	v1177 = *(*int64)(unsafe.Add(mBase, uint32(v1060)))
	*(*int64)(unsafe.Add(mBase, uint32(v1006))) = v1177
	goto L197
L201:
	;
	v1118 = *(*float64)(unsafe.Add(mBase, uint32(v1060)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1118)&int64(9223372036854775807)) {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	v1106 = *(*float64)(unsafe.Add(mBase, uint32(v1060)))
	if base.B2i32(base.F64_lt(v1100, v1106) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1106)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1006))) = v1106
	goto L201
L204:
	;
	v1136 = *(*float64)(unsafe.Add(mBase, uint32(v1006)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1136)&int64(9223372036854775807)) {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	v1124 = *(*float64)(unsafe.Add(mBase, uint32(v1006)+16))
	if base.B2i32(base.F64_gt(v1124, v1118) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1124)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L204
	} else {
		goto L206
	}
L206:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1006)+16)) = v1118
	goto L204
L207:
	;
	v1154 = *(*float64)(unsafe.Add(mBase, uint32(v1060)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1154)&int64(9223372036854775807)) {
		goto L197
	} else {
		goto L210
	}
L208:
	;
	v1142 = *(*float64)(unsafe.Add(mBase, uint32(v1060)+8))
	if base.B2i32(base.F64_lt(v1136, v1142) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1142)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1006)+8)) = v1142
	goto L207
L210:
	;
	v1160 = *(*float64)(unsafe.Add(mBase, uint32(v1006)+24))
	if base.B2i32(base.F64_gt(v1160, v1154) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1160)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L197
	} else {
		goto L211
	}
L211:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1006)+24)) = v1154
	goto L197
L212:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v1275 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v1274 + v1275
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1278+v1274<<(uint(v1275)%32)))) = uint16(v1041)
	v1285 = v1038
	goto L188
L213:
	;
	v1193 = *(*float64)(unsafe.Add(mBase, uint32(v1009)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1193)&int64(9223372036854775807)) {
		goto L216
	} else {
		goto L217
	}
L214:
	;
	goto L215
L215:
	;
	v1264 = *(*int64)(unsafe.Add(mBase, uint32(v1060)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1009)+24)) = v1264
	v1266 = *(*int64)(unsafe.Add(mBase, uint32(v1060)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1009)+16)) = v1266
	v1268 = *(*int64)(unsafe.Add(mBase, uint32(v1060)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1009)+8)) = v1268
	v1270 = *(*int64)(unsafe.Add(mBase, uint32(v1060)))
	*(*int64)(unsafe.Add(mBase, uint32(v1009))) = v1270
	goto L212
L216:
	;
	v1211 = *(*float64)(unsafe.Add(mBase, uint32(v1060)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1211)&int64(9223372036854775807)) {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	v1199 = *(*float64)(unsafe.Add(mBase, uint32(v1060)))
	if base.B2i32(base.F64_lt(v1193, v1199) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1199)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L216
	} else {
		goto L218
	}
L218:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1009))) = v1199
	goto L216
L219:
	;
	v1229 = *(*float64)(unsafe.Add(mBase, uint32(v1009)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1229)&int64(9223372036854775807)) {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	v1217 = *(*float64)(unsafe.Add(mBase, uint32(v1009)+16))
	if base.B2i32(base.F64_gt(v1217, v1211) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1217)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L219
	} else {
		goto L221
	}
L221:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1009)+16)) = v1211
	goto L219
L222:
	;
	v1247 = *(*float64)(unsafe.Add(mBase, uint32(v1060)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1247)&int64(9223372036854775807)) {
		goto L212
	} else {
		goto L225
	}
L223:
	;
	v1235 = *(*float64)(unsafe.Add(mBase, uint32(v1060)+8))
	if base.B2i32(base.F64_lt(v1229, v1235) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1235)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L222
	} else {
		goto L224
	}
L224:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1009)+8)) = v1235
	goto L222
L225:
	;
	v1253 = *(*float64)(unsafe.Add(mBase, uint32(v1009)+24))
	if base.B2i32(base.F64_gt(v1253, v1247) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1253)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L212
	} else {
		goto L226
	}
L226:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1009)+24)) = v1247
	goto L212
L227:
	;
	goto L187
L228:
	;
	v1300 = base.I32_trunc_sat_f64_s(base.F64_ceil(base.F64_mul(base.F64_convert_i32_u(v41), float64(0.3))))
	v1301 = int32(0)
	v1308 = v1301
	v1310 = v1301
	goto L229
L229:
	;
	v1328 = int32(4)
	v1330 = v1011 + v1310<<(uint(v1328)%32)
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1330)))
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v205+v1331<<(uint(v1328)%32))))
	v1336 = F_box_penalty(m, v1006, v1335)
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L1
	} else {
		goto L231
	}
L230:
	;
	F_pg_qsort(m, v1011, v1285, int32(16), int32(107))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L1
	} else {
		goto L235
	}
L231:
	;
	v1338 = F_box_penalty(m, v1009, v1335)
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v1341 = base.F64_abs(base.F64_sub(v1336, v1338))
	v1342 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(v1341, v1342)|base.F64_eq(base.F64_abs(v1336), v1342) == int32(0))&base.F64_ne(base.F64_abs(v1338), v1342) != 0 {
		goto L124
	} else {
		goto L233
	}
L233:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1330)+8)) = v1341
	v1356 = v1308 + int32(1)
	v1358 = v1356 & int32(_a_F_gist_box_picksplit_0)
	if base.Ui32(v1358) < base.Ui32(v1285) {
		v1308 = v1356
		v1310 = v1358
		goto L229
	} else {
		goto L234
	}
L234:
	;
	goto L230
L235:
	;
	v1364 = int32(0)
	v1373 = v1364
	v1375 = v1364
	goto L236
L236:
	;
	v1391 = int32(4)
	v1393 = v1011 + v1373<<(uint(v1391)%32)
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1393)))
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v205+v1394<<(uint(v1391)%32))))
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v1400 = v1285 - v1373
	if v1399+v1400 <= v1300 {
		goto L239
	} else {
		goto L240
	}
L237:
	;
	goto L178
L238:
	;
	v1790 = v1375 + int32(1)
	v1792 = v1790 & int32(_a_F_gist_box_picksplit_0)
	if base.Ui32(v1792) < base.Ui32(v1285) {
		v1373 = v1792
		v1375 = v1790
		goto L236
	} else {
		goto L310
	}
L239:
	;
	if int32(0) < v1399 {
		goto L243
	} else {
		goto L244
	}
L240:
	;
	goto L241
L241:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	if v1400+v1496 <= v1300 {
		goto L257
	} else {
		goto L258
	}
L242:
	;
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v1393)))
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v1488 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v1487 + v1488
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1491+v1487<<(uint(v1488)%32)))) = uint16(v1486)
	goto L238
L243:
	;
	v1405 = *(*float64)(unsafe.Add(mBase, uint32(v1006)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1405)&int64(9223372036854775807)) {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	goto L245
L245:
	;
	v1476 = *(*int64)(unsafe.Add(mBase, uint32(v1398)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1006)+24)) = v1476
	v1478 = *(*int64)(unsafe.Add(mBase, uint32(v1398)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1006)+16)) = v1478
	v1480 = *(*int64)(unsafe.Add(mBase, uint32(v1398)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1006)+8)) = v1480
	v1482 = *(*int64)(unsafe.Add(mBase, uint32(v1398)))
	*(*int64)(unsafe.Add(mBase, uint32(v1006))) = v1482
	goto L242
L246:
	;
	v1423 = *(*float64)(unsafe.Add(mBase, uint32(v1398)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1423)&int64(9223372036854775807)) {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	v1411 = *(*float64)(unsafe.Add(mBase, uint32(v1398)))
	if base.B2i32(base.F64_lt(v1405, v1411) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1411)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L246
	} else {
		goto L248
	}
L248:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1006))) = v1411
	goto L246
L249:
	;
	v1441 = *(*float64)(unsafe.Add(mBase, uint32(v1006)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1441)&int64(9223372036854775807)) {
		goto L252
	} else {
		goto L253
	}
L250:
	;
	v1429 = *(*float64)(unsafe.Add(mBase, uint32(v1006)+16))
	if base.B2i32(base.F64_gt(v1429, v1423) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1429)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L249
	} else {
		goto L251
	}
L251:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1006)+16)) = v1423
	goto L249
L252:
	;
	v1459 = *(*float64)(unsafe.Add(mBase, uint32(v1398)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1459)&int64(9223372036854775807)) {
		goto L242
	} else {
		goto L255
	}
L253:
	;
	v1447 = *(*float64)(unsafe.Add(mBase, uint32(v1398)+8))
	if base.B2i32(base.F64_lt(v1441, v1447) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1447)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L252
	} else {
		goto L254
	}
L254:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1006)+8)) = v1447
	goto L252
L255:
	;
	v1465 = *(*float64)(unsafe.Add(mBase, uint32(v1006)+24))
	if base.B2i32(base.F64_gt(v1465, v1459) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1465)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L242
	} else {
		goto L256
	}
L256:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1006)+24)) = v1459
	goto L242
L257:
	;
	if int32(0) < v1496 {
		goto L261
	} else {
		goto L262
	}
L258:
	;
	goto L259
L259:
	;
	v1592 = F_box_penalty(m, v1006, v1398)
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L1
	} else {
		goto L275
	}
L260:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1393)))
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v1584 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v1583 + v1584
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1587+v1583<<(uint(v1584)%32)))) = uint16(v1582)
	goto L238
L261:
	;
	v1501 = *(*float64)(unsafe.Add(mBase, uint32(v1009)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1501)&int64(9223372036854775807)) {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	goto L263
L263:
	;
	v1572 = *(*int64)(unsafe.Add(mBase, uint32(v1398)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1009)+24)) = v1572
	v1574 = *(*int64)(unsafe.Add(mBase, uint32(v1398)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1009)+16)) = v1574
	v1576 = *(*int64)(unsafe.Add(mBase, uint32(v1398)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1009)+8)) = v1576
	v1578 = *(*int64)(unsafe.Add(mBase, uint32(v1398)))
	*(*int64)(unsafe.Add(mBase, uint32(v1009))) = v1578
	goto L260
L264:
	;
	v1519 = *(*float64)(unsafe.Add(mBase, uint32(v1398)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1519)&int64(9223372036854775807)) {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	v1507 = *(*float64)(unsafe.Add(mBase, uint32(v1398)))
	if base.B2i32(base.F64_lt(v1501, v1507) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1507)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L264
	} else {
		goto L266
	}
L266:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1009))) = v1507
	goto L264
L267:
	;
	v1537 = *(*float64)(unsafe.Add(mBase, uint32(v1009)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1537)&int64(9223372036854775807)) {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	v1525 = *(*float64)(unsafe.Add(mBase, uint32(v1009)+16))
	if base.B2i32(base.F64_gt(v1525, v1519) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1525)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L267
	} else {
		goto L269
	}
L269:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1009)+16)) = v1519
	goto L267
L270:
	;
	v1555 = *(*float64)(unsafe.Add(mBase, uint32(v1398)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1555)&int64(9223372036854775807)) {
		goto L260
	} else {
		goto L273
	}
L271:
	;
	v1543 = *(*float64)(unsafe.Add(mBase, uint32(v1398)+8))
	if base.B2i32(base.F64_lt(v1537, v1543) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1543)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L270
	} else {
		goto L272
	}
L272:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1009)+8)) = v1543
	goto L270
L273:
	;
	v1561 = *(*float64)(unsafe.Add(mBase, uint32(v1009)+24))
	if base.B2i32(base.F64_gt(v1561, v1555) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1561)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L260
	} else {
		goto L274
	}
L274:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1009)+24)) = v1555
	goto L260
L275:
	;
	v1594 = F_box_penalty(m, v1009, v1398)
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	if base.F64_lt(v1592, v1594) != 0 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if int32(0) < v1597 {
		goto L281
	} else {
		goto L282
	}
L278:
	;
	goto L279
L279:
	;
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	if int32(0) < v1691 {
		goto L296
	} else {
		goto L297
	}
L280:
	;
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1393)))
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v1683 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v1682 + v1683
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1686+v1682<<(uint(v1683)%32)))) = uint16(v1681)
	goto L238
L281:
	;
	v1600 = *(*float64)(unsafe.Add(mBase, uint32(v1006)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1600)&int64(9223372036854775807)) {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	goto L283
L283:
	;
	v1671 = *(*int64)(unsafe.Add(mBase, uint32(v1398)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1006)+24)) = v1671
	v1673 = *(*int64)(unsafe.Add(mBase, uint32(v1398)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1006)+16)) = v1673
	v1675 = *(*int64)(unsafe.Add(mBase, uint32(v1398)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1006)+8)) = v1675
	v1677 = *(*int64)(unsafe.Add(mBase, uint32(v1398)))
	*(*int64)(unsafe.Add(mBase, uint32(v1006))) = v1677
	goto L280
L284:
	;
	v1618 = *(*float64)(unsafe.Add(mBase, uint32(v1398)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1618)&int64(9223372036854775807)) {
		goto L287
	} else {
		goto L288
	}
L285:
	;
	v1606 = *(*float64)(unsafe.Add(mBase, uint32(v1398)))
	if base.B2i32(base.F64_lt(v1600, v1606) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1606)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L284
	} else {
		goto L286
	}
L286:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1006))) = v1606
	goto L284
L287:
	;
	v1636 = *(*float64)(unsafe.Add(mBase, uint32(v1006)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1636)&int64(9223372036854775807)) {
		goto L290
	} else {
		goto L291
	}
L288:
	;
	v1624 = *(*float64)(unsafe.Add(mBase, uint32(v1006)+16))
	if base.B2i32(base.F64_gt(v1624, v1618) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1624)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L287
	} else {
		goto L289
	}
L289:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1006)+16)) = v1618
	goto L287
L290:
	;
	v1654 = *(*float64)(unsafe.Add(mBase, uint32(v1398)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1654)&int64(9223372036854775807)) {
		goto L280
	} else {
		goto L293
	}
L291:
	;
	v1642 = *(*float64)(unsafe.Add(mBase, uint32(v1398)+8))
	if base.B2i32(base.F64_lt(v1636, v1642) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1642)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L290
	} else {
		goto L292
	}
L292:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1006)+8)) = v1642
	goto L290
L293:
	;
	v1660 = *(*float64)(unsafe.Add(mBase, uint32(v1006)+24))
	if base.B2i32(base.F64_gt(v1660, v1654) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1660)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L280
	} else {
		goto L294
	}
L294:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1006)+24)) = v1654
	goto L280
L295:
	;
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(v1393)))
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v1777 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v1776 + v1777
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1780+v1776<<(uint(v1777)%32)))) = uint16(v1775)
	goto L238
L296:
	;
	v1694 = *(*float64)(unsafe.Add(mBase, uint32(v1009)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1694)&int64(9223372036854775807)) {
		goto L299
	} else {
		goto L300
	}
L297:
	;
	goto L298
L298:
	;
	v1765 = *(*int64)(unsafe.Add(mBase, uint32(v1398)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1009)+24)) = v1765
	v1767 = *(*int64)(unsafe.Add(mBase, uint32(v1398)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1009)+16)) = v1767
	v1769 = *(*int64)(unsafe.Add(mBase, uint32(v1398)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1009)+8)) = v1769
	v1771 = *(*int64)(unsafe.Add(mBase, uint32(v1398)))
	*(*int64)(unsafe.Add(mBase, uint32(v1009))) = v1771
	goto L295
L299:
	;
	v1712 = *(*float64)(unsafe.Add(mBase, uint32(v1398)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1712)&int64(9223372036854775807)) {
		goto L302
	} else {
		goto L303
	}
L300:
	;
	v1700 = *(*float64)(unsafe.Add(mBase, uint32(v1398)))
	if base.B2i32(base.F64_lt(v1694, v1700) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1700)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L299
	} else {
		goto L301
	}
L301:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1009))) = v1700
	goto L299
L302:
	;
	v1730 = *(*float64)(unsafe.Add(mBase, uint32(v1009)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1730)&int64(9223372036854775807)) {
		goto L305
	} else {
		goto L306
	}
L303:
	;
	v1718 = *(*float64)(unsafe.Add(mBase, uint32(v1009)+16))
	if base.B2i32(base.F64_gt(v1718, v1712) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1718)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1009)+16)) = v1712
	goto L302
L305:
	;
	v1748 = *(*float64)(unsafe.Add(mBase, uint32(v1398)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1748)&int64(9223372036854775807)) {
		goto L295
	} else {
		goto L308
	}
L306:
	;
	v1736 = *(*float64)(unsafe.Add(mBase, uint32(v1398)+8))
	if base.B2i32(base.F64_lt(v1730, v1736) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1736)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L305
	} else {
		goto L307
	}
L307:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1009)+8)) = v1736
	goto L305
L308:
	;
	v1754 = *(*float64)(unsafe.Add(mBase, uint32(v1009)+24))
	if base.B2i32(base.F64_gt(v1754, v1748) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1754)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L295
	} else {
		goto L309
	}
L309:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1009)+24)) = v1748
	goto L295
L310:
	;
	goto L237
L311:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gist_circle_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v33 float64
	_ = v33
	var v43 int32
	_ = v43
	var v47 float64
	_ = v47
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v64 float64
	_ = v64
	var v71 int32
	_ = v71
	var v75 float64
	_ = v75
	var v77 float64
	_ = v77
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v18 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v18)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if base.B2i32(v21 == v5)|base.B2i32(v15 == v5) == v5 {
		v29 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
		v30 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
		v31 = base.F64_add(v29, v30)
		v33 = math.Float64frombits(uint64(0x7ff0000000000000))
		v43 = base.F64_ne(base.F64_abs(v30), v33)
		if base.B2i32(base.F64_ne(base.F64_abs(v31), v33)|base.F64_eq(base.F64_abs(v29), v33) == int32(0))&v43 != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v107 = m.ExcPending
			if v107 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v12))) = v31
			v47 = math.Float64frombits(uint64(0x7ff0000000000000))
			v48 = base.F64_eq(base.F64_abs(v30), v47)
			v49 = base.F64_sub(v29, v30)
			if base.B2i32(v48|base.F64_ne(base.F64_abs(v49), v47) == int32(0))&base.F64_ne(base.F64_abs(v29), v47) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v49
				v61 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
				v62 = base.F64_add(v61, v30)
				v64 = math.Float64frombits(uint64(0x7ff0000000000000))
				if v43 != 0 {
					v71 = base.F64_ne(base.F64_abs(v62), v64) | base.F64_eq(base.F64_abs(v61), v64)
				} else {
					v71 = int32(1)
				}
				if v71 == int32(0) {
					F_float_overflow_error(m)
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v62
					v75 = base.F64_sub(v61, v30)
					v77 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.B2i32(base.F64_ne(base.F64_abs(v75), v77)|v48 == int32(0))&base.F64_ne(base.F64_abs(v61), v77) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v75
						v89 = F_rtree_internal_consistent(m, v21, v12, v14&int32(_a_F_gist_circle_consistent_0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							v93 = v89
							m.G0 = v12 + int32(32)
							return v93
						}
					}
				}
			}
		}
	} else {
		v93 = v5
		m.G0 = v12 + int32(32)
		return v93
	}
}
func F_gist_desc(m *base.Module, l0 int32, l1 int32) {
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
	var v15 int32
	_ = v15
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v65 int64
	_ = v65
	var v69 int32
	_ = v69
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+48)))
	switch int32(base.Ui32(v15)>>(uint(int32(4))%32)) - int32(1) {
	case 0:
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+6)))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v41
		*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v40
		if v39 != 0 {
			v46 = int32(84)
		} else {
			v46 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v46
		F_appendStringInfo(m, l0, int32(_a_F_gist_desc_0), v11+int32(48))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return
		} else {
			m.G0 = v11 + int32(80)
			return
		}
	case 1:
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)))
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
		*(*uint32)(unsafe.Add(mBase, uint32(v11)+36)) = uint32(v23)
		if v22 != 0 {
			v27 = int32(84)
		} else {
			v27 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v27
		v30 = int64(base.Ui64(v23) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v11)+32)) = uint32(v30)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v20
		F_appendStringInfo(m, l0, int32(_a_F_gist_desc_1), v11+int32(16))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			m.G0 = v11 + int32(80)
			return
		}
	case 2:
		v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+18)))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v53
		F_appendStringInfo(m, l0, int32(_a_F_gist_desc_2), v11-int32(-64))
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return
		} else {
			m.G0 = v11 + int32(80)
			return
		}
	default:
		m.G0 = v11 + int32(80)
		return
	case 5:
		v60 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
		v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+8)))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v61
		*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v60)
		v65 = int64(base.Ui64(v60) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v65)
		F_appendStringInfo(m, l0, int32(_a_F_gist_desc_3), v11)
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return
		} else {
			m.G0 = v11 + int32(80)
			return
		}
	}
}
func F_gist_point_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)))
	if v7 != int32(1) {
		return v6
	} else {
		v12 = F_palloc(m, int32(32))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v18 = F_palloc(m, int32(16))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v20
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v22
				v24 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
				v25 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v25
				*(*int64)(unsafe.Add(mBase, uint32(v12))) = v24
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v12
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v29
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v31
				v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)))
				v34 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v18)+14)) = uint8(v34)
				*(*uint16)(unsafe.Add(mBase, uint32(v18)+12)) = uint16(v33)
				return v18
			}
		}
	}
}
func F_gist_point_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v17 float64
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = F_palloc(m, int32(16))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
			*(*float64)(unsafe.Add(mBase, uint32(v13))) = v15
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v13
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v20
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v22
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+12)))
			v25 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v25)
			*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v24)
			return v8
		}
	}
}
func F_gist_poly_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+14)))
	if v5 != int32(1) {
		return v4
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v10 = F_pg_detoast_datum(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = F_palloc(m, int32(32))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = *(*int64)(unsafe.Add(mBase, uint32(v10)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v17
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v19
				v21 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v21
				v23 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v15))) = v23
				v26 = F_palloc(m, int32(16))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v15
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v29
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v31
					v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+12)))
					v34 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+14)) = uint8(v34)
					*(*uint16)(unsafe.Add(mBase, uint32(v26)+12)) = uint16(v33)
					return v26
				}
			}
		}
	}
}
func F_verify_gist_page(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v8 = F_get_page_from_raw(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)))
		if v12 != 0 {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+19)))
			v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+16)))
			if (v13<<(uint(int32(8))%32)-v16)&int32(_a_F_verify_gist_page_0) != int32(16) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(_a_F_verify_gist_page_1)
						F_errmsg(m, int32(_a_F_verify_gist_page_2), v4+int32(-16))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+16)))
							v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+19)))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = (v46<<(uint(int32(8))%32) - v45) & int32(_a_F_verify_gist_page_0)
							F_errdetail(m, int32(_a_F_verify_gist_page_3), v4+int32(-32))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_verify_gist_page_4), int32(56), int32(_a_F_verify_gist_page_5))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
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
				v22 = v8 + v16
				v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+14)))
				if v23 != int32(_a_F_verify_gist_page_6) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a_F_verify_gist_page_1)
							F_errmsg(m, int32(_a_F_verify_gist_page_2), v4+int32(-48))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+14)))
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v79
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_verify_gist_page_6)
								F_errdetail(m, int32(_a_F_verify_gist_page_7), v6)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_verify_gist_page_4), int32(65), int32(_a_F_verify_gist_page_5))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
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
					m.G0 = v6 - int32(-64)
					return v8
				}
			}
		} else {
			m.G0 = v6 - int32(-64)
			return v8
		}
	}
}
