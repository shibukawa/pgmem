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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
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
	var v64 int32
	_ = v64
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int64
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(320)
	m.G0 = v17
	v19 = int32(_a_F_gistFetchTuple_0)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_gistFetchTuple[0]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_gistFetchTuple[0])) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+10)))
	if v4 < v25 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v40 = v4
	goto L4
L2:
	;
	v126 = v4
	goto L3
L3:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if v126 < v136 {
		goto L23
	} else {
		goto L24
	}
L4:
	;
	v50 = v40 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v52 = v17 + v40
	v53 = F_index_getattr_2(m, l2, v50, v51, v52)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v126 = v50
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	v58 = v40 * int32(28)
	v59 = l0 + v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+uint32(_c_F_gistFetchTuple[1])))
	if v60 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(v118)+10)))
	if v50 < v119 {
		v40 = v50
		goto L4
	} else {
		goto L22
	}
L9:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v61 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v59)+1816))
	if v94 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v64 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+314)) = uint8(v64)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+312)) = uint16(v64)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+308)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v17)+304)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v17)+296)) = v53
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(_a_F_gistFetchTuple_1)+v40<<(uint(int32(2))%32))))
	v82 = F_FunctionCall1Coll(m, l0+int32(_a_F_gistFetchTuple_2)+v58, v81, base.I64_extend_i32_u(v17+int32(296)))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(32)+v40<<(uint(int32(3))%32)))) = int64(0)
	goto L8
L15:
	;
	v85 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v82))))
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(32)+v40<<(uint(int32(3))%32)))) = v85
	goto L8
L16:
	;
	v101 = v17 + int32(32) + v40<<(uint(int32(3))%32)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v102 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v108 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v108)
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(32)+v40<<(uint(int32(3))%32)))) = int64(0)
	goto L8
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v101))) = v53
	goto L8
L20:
	;
	goto L21
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v101))) = int64(0)
	goto L8
L22:
	;
	goto L5
L23:
	;
	v143 = v126
	goto L26
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistFetchTuple[0])) = v20
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v186 = F_heap_form_tuple(m, v183, v17+int32(32), v17)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L6
	} else {
		goto L30
	}
L26:
	;
	v158 = v143 + int32(1)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v161 = F_index_getattr_2(m, l2, v158, v159, v17+v143)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L6
	} else {
		goto L28
	}
L27:
	;
	goto L25
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17+int32(32)+v143<<(uint(int32(3))%32)))) = v161
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if v158 < v165 {
		v143 = v158
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	m.G0 = v17 + int32(320)
	return v186
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
								F_errfinish(m, int32(_a_F_gistLoadNodeBuffer_2), int32(749), int32(_a_F_gistLoadNodeBuffer_3))
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v56 int32
	_ = v56
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v158 int64
	_ = v158
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v331 int32
	_ = v331
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int64
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int64
	_ = v433
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v448 int64
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int64
	_ = v490
	var v492 int64
	_ = v492
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v534 int32
	_ = v534
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int64
	_ = v633
	var v636 int64
	_ = v636
	var v637 int64
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v650 int64
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v722 int64
	_ = v722
	var v723 int64
	_ = v723
	var v724 int32
	_ = v724
	var v740 int64
	_ = v740
	var v749 int64
	_ = v749
	var v758 int32
	_ = v758
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v770 float32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v776 float32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v787 float32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v795 float32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v800 float32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 float32
	_ = v804
	var v805 int32
	_ = v805
	var v814 int64
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v820 int64
	_ = v820
	var v821 int64
	_ = v821
	var v824 int32
	_ = v824
	var v846 int32
	_ = v846
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v878 int64
	_ = v878
	var v880 int64
	_ = v880
	var v882 int32
	_ = v882
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v893 int64
	_ = v893
	var v894 int64
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v911 int32
	_ = v911
	var v915 int64
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v928 int64
	_ = v928
	var v929 int32
	_ = v929
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v987 float32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1012 int32
	_ = v1012
	var v1033 int64
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1042 int32
	_ = v1042
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1089 float32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1264 int32
	_ = v1264
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1440 int32
	_ = v1440
	var v1463 int32
	_ = v1463
	var v1472 int32
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1497 int32
	_ = v1497
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1588 int32
	_ = v1588
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1613 int32
	_ = v1613
	var v1630 int32
	_ = v1630
	var v1632 int64
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1654 float32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1657 int64
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1669 float32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1724 int32
	_ = v1724
	var v1760 int32
	_ = v1760
	var v1764 int32
	_ = v1764
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1774 int32
	_ = v1774
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1835 int32
	_ = v1835
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1843 int32
	_ = v1843
	var v1849 int32
	_ = v1849
	var v1851 int32
	_ = v1851
	var v1857 int32
	_ = v1857
	var v1859 int32
	_ = v1859
	var v1872 int32
	_ = v1872
	var v1875 int32
	_ = v1875
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1898 int32
	_ = v1898
	var v1901 int32
	_ = v1901
	var v1907 int32
	_ = v1907
	var v1909 int32
	_ = v1909
	var v1927 int32
	_ = v1927
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1949 int64
	_ = v1949
	var v1951 int32
	_ = v1951
	var v1952 int64
	_ = v1952
	var v1954 int32
	_ = v1954
	var v1956 int64
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1989 int32
	_ = v1989
	var v1992 int32
	_ = v1992
	var v2011 int32
	_ = v2011
	var v2014 int32
	_ = v2014
	var v2018 int32
	_ = v2018
	var v2024 int32
	_ = v2024
	var v2027 int32
	_ = v2027
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2041 int32
	_ = v2041
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2077 int32
	_ = v2077
	var v2079 int32
	_ = v2079
	var v2098 int32
	_ = v2098
	var v2101 int32
	_ = v2101
	var v2105 int32
	_ = v2105
	var v2111 int32
	_ = v2111
	var v2114 int32
	_ = v2114
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2129 int32
	_ = v2129
	var v2152 int32
	_ = v2152
	var v2154 int64
	_ = v2154
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2162 int64
	_ = v2162
	var v2164 int64
	_ = v2164
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2235 int32
	_ = v2235
	v32 = m.G0
	v34 = v32 - int32(880)
	m.G0 = v34
	v40 = F_palloc(m, l3*int32(24)+int32(32))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v42 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = l3 + v42
	v46 = l3 << (uint(v42) % 32)
	v47 = F_palloc(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
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
	m.G0 = v34 + int32(880)
	return
L5:
	;
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(v2228)))
	if v2229 < int32(2) {
		goto L4
	} else {
		goto L286
	}
L6:
	;
	if l6 != 0 {
		goto L4
	} else {
		goto L285
	}
L7:
	;
	v410 = l5 + int32(304)
	v411 = v410 + l6
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	v413 = int32(1)
	v414 = v412 ^ v413
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+16)) = uint8(v414)
	v417 = l5 + int32(592)
	v418 = v417 + l6
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418))))
	v421 = v419 ^ v413
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+40)) = uint8(v421)
	v424 = l5 + int32(48)
	v426 = l6 << (uint(int32(3)) % 32)
	v427 = v424 + v426
	v428 = *(*int64)(unsafe.Add(mBase, uint32(v427)))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = v428
	v431 = l5 + int32(336)
	v432 = v431 + v426
	v433 = *(*int64)(unsafe.Add(mBase, uint32(v432)))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+32)) = v433
	v442 = l4 + l6<<(uint(int32(2))%32)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v442)+uint32(_c_F_gistSplitByKey[0])))
	v448 = F_FunctionCall2Coll(m, l4+l6*int32(28)+int32(_a_F_gistSplitByKey_0), v445, base.I64_extend_i32_u(v40), base.I64_extend_i32_u(l5))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L67
	}
L8:
	;
	v295 = l5 + l6
	v296 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v295)+592)) = uint8(v296)
	*(*uint8)(unsafe.Add(mBase, uint32(v295)+304)) = uint8(v296)
	v301 = l6 + v296
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	if v301 < v303 {
		goto L53
	} else {
		goto L54
	}
L9:
	;
	v52 = v40 + int32(8)
	v55 = int32(1)
	v56 = l6 + v55
	v71 = v55
	v76 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	if l3 != 0 {
		goto L7
	} else {
		goto L52
	}
L12:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l2+v71<<(uint(int32(2))%32)-int32(4))))
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v102)+6)))
	if int32(0) <= v103 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	if l3 == v181 {
		goto L8
	} else {
		goto L39
	}
L14:
	;
	v185 = v71 + int32(1)
	if v185 <= l3 {
		v71 = v185
		v76 = v181
		goto L12
	} else {
		goto L38
	}
L15:
	;
	F_gistdentryinit(m, l4, l6, v52+v71*int32(24), int64(0), l0, l1, v71&int32(_a_F_gistSplitByKey_1), int32(1))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L37
	}
L16:
	;
	F_gistdentryinit(m, l4, l6, v52+v71*int32(24), v158, l0, l1, v71&int32(_a_F_gistSplitByKey_1), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L36
	}
L17:
	;
	v152 = F_nocache_index_getattr(m, v102, v56, v93)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L35
	}
L18:
	;
	v110 = v93 + v56<<(uint(int32(3))%32) + int32(20)
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v110))))
	if v111 < int32(0) {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+l6>>(uint(int32(3))%32))+8)))
	if v55<<(uint(l6&int32(7))%32)&v146 == int32(0) {
		goto L15
	} else {
		goto L34
	}
L21:
	;
	v116 = v102 + v111 + int32(8)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+4)))
	if v117 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v110)+2)))
	if base.I32_popcnt(v120) != int32(1) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v158 = base.I64_extend_i32_u(v116)
	goto L16
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L31
	}
L26:
	;
	switch base.I32_ctz(v120) {
	case 0:
		goto L30
	case 1:
		goto L29
	case 2:
		goto L28
	case 3:
		goto L27
	default:
		goto L25
	}
L27:
	;
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v116)))
	v158 = v128
	goto L16
L28:
	;
	v127 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116))))
	v158 = v127
	goto L16
L29:
	;
	v126 = int64(*(*int16)(unsafe.Add(mBase, uint32(v116))))
	v158 = v126
	goto L16
L30:
	;
	v125 = int64(*(*int8)(unsafe.Add(mBase, uint32(v116))))
	v158 = v125
	goto L16
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v120
	F_errmsg_internal(m, int32(_a_F_gistSplitByKey_2), v34+int32(16))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_gistSplitByKey_3), int32(123), int32(_a_F_gistSplitByKey_4))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
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
	goto L17
L35:
	;
	v158 = v152
	goto L16
L36:
	;
	v181 = v76
	goto L14
L37:
	;
	v173 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v47+v76<<(uint(v173)%32)))) = uint16(v71)
	v181 = v76 + v173
	goto L14
L38:
	;
	goto L13
L39:
	;
	if v181 <= int32(0) {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v47
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5+l6)+592)) = uint8(v193)
	v196 = F_palloc(m, v46)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v196
	v211 = int32(1)
	v212 = int32(0)
	goto L42
L42:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	if v233 <= v212 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	if l6 != 0 {
		goto L4
	} else {
		goto L49
	}
L44:
	;
	if l3 != v211 {
		v211 = v211 + int32(1)
		v212 = v252
		goto L42
	} else {
		goto L48
	}
L45:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v243 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v242 + v243
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*uint16)(unsafe.Add(mBase, uint32(v246+v242<<(uint(v243)%32)))) = uint16(v211)
	v252 = v212
	goto L44
L46:
	;
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+v212<<(uint(int32(1))%32)))))
	if v211 != v238 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v252 = v212 + int32(1)
	goto L44
L48:
	;
	goto L43
L49:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	if v257 != int32(1) {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+624)) = int32(0)
	F_gistunionsubkey(m, l4, l2, l5)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L5
L52:
	;
	goto L8
L53:
	;
	F_gistSplitByKey(m, l0, l1, l2, l3, l4, l5, v301)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v307 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v307
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v307
	v311 = F_palloc(m, v46)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	goto L6
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v311
	v314 = F_palloc(m, v46)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v314
	if l3 <= int32(0) {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v319 = int32(1)
	v331 = v319
	goto L60
L60:
	;
	if base.Ui32(v331) < base.Ui32(int32(base.Ui32(l3)>>(uint(v319)%32))) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L6
L62:
	;
	if base.B2i32(l3 == v331) == int32(0) {
		v331 = v331 + int32(1)
		goto L60
	} else {
		goto L66
	}
L63:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v355 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v354 + v355
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v358+v354<<(uint(v355)%32)))) = uint16(v331)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v364 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v363 + v364
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*uint16)(unsafe.Add(mBase, uint32(v367+v363<<(uint(v364)%32)))) = uint16(v331)
	goto L62
L66:
	;
	goto L61
L67:
	;
	v451 = l5 + int32(16)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v452 != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v711 = l5 + int32(32)
	v713 = l5 + int32(8)
	v715 = l5 + int32(40)
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451))))
	if v716 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L69:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v658 = v653 + v452<<(uint(int32(1))%32) - int32(2)
	v659 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v658))))
	if v659 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L70:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	if v453 != 0 {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v457 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	if v457 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	v483 = int32(1)
	v484 = v482 ^ v483
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+16)) = uint8(v484)
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418))))
	v488 = v486 ^ v483
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+40)) = uint8(v488)
	v490 = *(*int64)(unsafe.Add(mBase, uint32(v427)))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = v490
	v492 = *(*int64)(unsafe.Add(mBase, uint32(v432)))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+32)) = v492
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v498 = (v494 - v483) & int32(_a_F_gistSplitByKey_1)
	v502 = v498<<(uint(v483)%32) + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+112)) = v502
	v504 = F_palloc(m, v502)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L82
	}
L78:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = l6 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v462 + int32(4)
	F_errmsg(m, int32(_a_F_gistSplitByKey_5), v34)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errhint(m, int32(_a_F_gistSplitByKey_6), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_gistSplitByKey_7), int32(448), int32(_a_F_gistSplitByKey_8))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	goto L77
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v504
	v507 = F_palloc(m, v502)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v509 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v509
	v515 = v494 & int32(_a_F_gistSplitByKey_1)
	if v515 != int32(1) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v518 = int32(2)
	if base.Ui32(v515) <= base.Ui32(v518) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v616 = F_palloc(m, v611*int32(24)+int32(8))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L97
	}
L87:
	;
	v521 = v518
	goto L89
L88:
	;
	v521 = v515
	goto L89
L89:
	;
	v522 = int32(1)
	v534 = v522
	goto L90
L90:
	;
	if base.Ui32(v534) <= base.Ui32(int32(base.Ui32(v498)>>(uint(v522)%32))) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L86
L92:
	;
	v578 = v534 + int32(1)
	if v578 != v521 {
		v534 = v578
		goto L90
	} else {
		goto L96
	}
L93:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v559 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v557+v558<<(uint(v559)%32)))) = uint16(v534)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v563 + v559
	goto L92
L94:
	;
	goto L95
L95:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v569 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v567+v568<<(uint(v569)%32)))) = uint16(v534)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v573 + v569
	goto L92
L96:
	;
	goto L91
L97:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v616))) = v618
	v621 = v40 + int32(32)
	v623 = v616 + int32(8)
	v625 = v618 * int32(24)
	if v625 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	base.MemoryCopy(m, v623, v621, v625)
	goto L100
L99:
	;
	goto L100
L100:
	;
	v631 = l4 + l6*int32(28) + int32(916)
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v442)+uint32(_c_F_gistSplitByKey[0])))
	v633 = base.I64_extend_i32_u(v616)
	v636 = base.I64_extend_i32_u(v34 + int32(112))
	v637 = F_FunctionCall2Coll(m, v631, v632, v633, v636)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = v637
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v616))) = v640
	v643 = v640 * int32(24)
	if v643 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	base.MemoryCopy(m, v623, v621+v644*int32(24), v643)
	goto L104
L103:
	;
	goto L104
L104:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v442)+uint32(_c_F_gistSplitByKey[0])))
	v650 = F_FunctionCall2Coll(m, v631, v649, v633, v636)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+32)) = v650
	goto L68
L106:
	;
	v662 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40))))
	v664 = v662 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v658))) = uint16(v664)
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v667 = v666
	goto L108
L107:
	;
	v667 = v453
	goto L108
L108:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v673 = v668 + v667<<(uint(int32(1))%32) - int32(2)
	v674 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v673))))
	if v674 != 0 {
		goto L68
	} else {
		goto L109
	}
L109:
	;
	v675 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40))))
	v677 = v675 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v673))) = uint16(v677)
	goto L68
L110:
	;
	v878 = *(*int64)(unsafe.Add(mBase, uint32(l5)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v427))) = v878
	v880 = *(*int64)(unsafe.Add(mBase, uint32(l5)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v432))) = v880
	v882 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v411))) = uint8(v882)
	*(*uint8)(unsafe.Add(mBase, uint32(v418))) = uint8(v882)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+624)) = v882
	v889 = l6 + int32(1)
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v890)))
	if v891 <= v889 {
		goto L6
	} else {
		goto L141
	}
L111:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v715))))
	if v719 != int32(1) {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v722 = *(*int64)(unsafe.Add(mBase, uint32(v432)))
	v723 = *(*int64)(unsafe.Add(mBase, uint32(v427)))
	v724 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+130)) = uint8(v724)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+128)) = uint16(v724)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+124)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v34)+120)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v34)+112)) = v723
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+98)) = uint8(v724)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+96)) = uint16(v724)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+92)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v34)+88)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v34)+80)) = v722
	v740 = *(*int64)(unsafe.Add(mBase, uint32(v713)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+74)) = uint8(v724)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+72)) = uint16(v724)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+68)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v34)+56)) = v740
	v749 = *(*int64)(unsafe.Add(mBase, uint32(v711)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+42)) = uint8(v724)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+40)) = uint16(v724)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v749
	if v716 != 0 {
		goto L118
	} else {
		goto L119
	}
L114:
	;
	goto L113
L115:
	;
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451))))
	if v846 == int32(1) {
		goto L133
	} else {
		goto L134
	}
L116:
	;
	v814 = *(*int64)(unsafe.Add(mBase, uint32(l5)+20))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v815
	v817 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v814
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v817
	v820 = *(*int64)(unsafe.Add(mBase, uint32(l5)+8))
	v821 = *(*int64)(unsafe.Add(mBase, uint32(l5)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = v821
	*(*int64)(unsafe.Add(mBase, uint32(l5)+32)) = v820
	v824 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+74)) = uint8(v824)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+72)) = uint16(v824)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+68)) = v824
	*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v34)+56)) = v821
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+42)) = uint8(v824)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+40)) = uint16(v824)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v824
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v820
	goto L115
L117:
	;
	v782 = v34 + int32(112)
	v783 = int32(0)
	v785 = v34 + int32(56)
	v787 = F_gistpenalty(m, l4, l6, v782, v783, v785, v783)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L1
	} else {
		goto L128
	}
L118:
	;
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v715))))
	if v758 == int32(1) {
		goto L117
	} else {
		goto L121
	}
L119:
	;
	v765 = v34 + int32(80)
	goto L120
L120:
	;
	v766 = int32(0)
	v770 = F_gistpenalty(m, l4, l6, v765, v766, v34+int32(56), v766)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L122
	}
L121:
	;
	v765 = v34 + int32(112)
	goto L120
L122:
	;
	v772 = int32(0)
	v776 = F_gistpenalty(m, l4, l6, v765, v772, v34+int32(24), v772)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	if base.F32_lt(v770, v776) != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v779 = v451
	goto L126
L125:
	;
	v779 = v715
	goto L126
L126:
	;
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779))))
	if v780 != 0 {
		goto L115
	} else {
		goto L127
	}
L127:
	;
	goto L116
L128:
	;
	v790 = v34 + int32(80)
	v791 = int32(0)
	v793 = v34 + int32(24)
	v795 = F_gistpenalty(m, l4, l6, v790, v791, v793, v791)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v798 = int32(0)
	v800 = F_gistpenalty(m, l4, l6, v782, v798, v793, v798)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v802 = int32(0)
	v804 = F_gistpenalty(m, l4, l6, v790, v802, v785, v802)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	if base.F32_gt(base.F32_add(v787, v795), base.F32_add(v800, v804)) == int32(0) {
		goto L115
	} else {
		goto L132
	}
L132:
	;
	goto L116
L133:
	;
	F_gistMakeUnionKey(m, l4, l6, v34+int32(112), v34+int32(56), v713, v34+int32(55))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L1
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v715))))
	if v857 == int32(1) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	goto L135
L137:
	;
	F_gistMakeUnionKey(m, l4, l6, v34+int32(80), v34+int32(24), v711, v34+int32(55))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v868 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v715))) = uint8(v868)
	*(*uint8)(unsafe.Add(mBase, uint32(v451))) = uint8(v868)
	goto L110
L140:
	;
	goto L139
L141:
	;
	v893 = *(*int64)(unsafe.Add(mBase, uint32(v713)))
	v894 = *(*int64)(unsafe.Add(mBase, uint32(v711)))
	v895 = m.G0
	v897 = v895 - int32(16)
	m.G0 = v897
	v899 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v897)+15)) = uint8(v899)
	v911 = *(*int32)(unsafe.Add(mBase, uint32(l4+l6<<(uint(int32(2))%32))+uint32(_c_F_gistSplitByKey[0])))
	v915 = F_FunctionCall3Coll(m, l4+l6*int32(28)+int32(_a_F_gistSplitByKey_9), v911, v893, v894, base.I64_extend_i32_u(v897+int32(15)))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v897)+15)))
	m.G0 = v897 + int32(16)
	if v917 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(l5)+624))
	if v1760 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L144:
	;
	v921 = int32(1)
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v925 = F_palloc0_mul(m, v921, v922+v921)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+624)) = v925
	v928 = *(*int64)(unsafe.Add(mBase, uint32(l5)+32))
	v929 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+130)) = uint8(v929)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+128)) = uint16(v929)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+124)) = v929
	*(*int32)(unsafe.Add(mBase, uint32(v34)+120)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v34)+112)) = v928
	v938 = v40 + int32(8)
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v929 < v940 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v953 = int32(0)
	v954 = v929
	goto L149
L147:
	;
	v1012 = v929
	goto L148
L148:
	;
	v1033 = *(*int64)(unsafe.Add(mBase, uint32(l5)+8))
	v1034 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+130)) = uint8(v1034)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+128)) = uint16(v1034)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+124)) = v1034
	*(*int32)(unsafe.Add(mBase, uint32(v34)+120)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v34)+112)) = v1033
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	if v1034 < v1042 {
		goto L156
	} else {
		goto L157
	}
L149:
	;
	v977 = int32(0)
	v978 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v982 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v978+v953<<(uint(int32(1))%32)))))
	v987 = F_gistpenalty(m, l4, l6, v34+int32(112), v977, v938+v982*int32(24), v977)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L1
	} else {
		goto L151
	}
L150:
	;
	v1012 = v997
	goto L148
L151:
	;
	if base.F32_eq(v987, float32(0)) != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(l5)+624))
	v993 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v991+v982))) = uint8(v993)
	v997 = v954 + v993
	goto L154
L153:
	;
	v997 = v954
	goto L154
L154:
	;
	v999 = v953 + int32(1)
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v999 < v1000 {
		v953 = v999
		v954 = v997
		goto L149
	} else {
		goto L155
	}
L155:
	;
	goto L150
L156:
	;
	v1055 = int32(0)
	v1056 = v1012
	goto L159
L157:
	;
	v1112 = v1042
	v1114 = v1012
	goto L158
L158:
	;
	if v1114 <= int32(0) {
		goto L6
	} else {
		goto L166
	}
L159:
	;
	v1079 = int32(0)
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v1084 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1080+v1055<<(uint(int32(1))%32)))))
	v1089 = F_gistpenalty(m, l4, l6, v34+int32(112), v1079, v938+v1084*int32(24), v1079)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L1
	} else {
		goto L161
	}
L160:
	;
	v1112 = v1102
	v1114 = v1099
	goto L158
L161:
	;
	if base.F32_eq(v1089, float32(0)) != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(l5)+624))
	v1095 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1093+v1084))) = uint8(v1095)
	v1099 = v1056 + v1095
	goto L164
L163:
	;
	v1099 = v1056
	goto L164
L164:
	;
	v1101 = v1055 + int32(1)
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	if v1101 < v1102 {
		v1055 = v1101
		v1056 = v1099
		goto L159
	} else {
		goto L165
	}
L165:
	;
	goto L160
L166:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(l5)+624))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if int32(0) < v1138 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v1142 = int32(0)
	if v1138 == int32(1) {
		goto L172
	} else {
		goto L173
	}
L168:
	;
	v1294 = v1112
	v1295 = v1137
	v1297 = v1138
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v1297
	if int32(0) < v1294 {
		goto L190
	} else {
		goto L191
	}
L170:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(l5)+624))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v1294 = v1285
	v1295 = v1284
	v1297 = v1264
	goto L169
L171:
	;
	v1247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1141+v1225<<(uint(int32(1))%32)))))
	v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137+v1247))))
	if v1249 != 0 {
		goto L187
	} else {
		goto L188
	}
L172:
	;
	v1221 = v1141
	v1224 = v1138
	v1225 = v1142
	goto L171
L173:
	;
	goto L174
L174:
	;
	v1158 = v1141
	v1161 = v1138
	v1162 = v1142
	v1164 = int32(0)
	goto L175
L175:
	;
	v1183 = v1141 + v1162<<(uint(int32(1))%32)
	v1184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1183))))
	v1186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137+v1184))))
	if v1186 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	if v1138&int32(1) == int32(0) {
		v1264 = v1205
		goto L170
	} else {
		goto L186
	}
L177:
	;
	v1196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1183)+2)))
	v1198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137+v1196))))
	if v1198 != 0 {
		goto L182
	} else {
		goto L183
	}
L178:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1158))) = uint16(v1184)
	v1194 = v1158 + int32(2)
	v1195 = v1161
	goto L177
L179:
	;
	goto L180
L180:
	;
	v1194 = v1158
	v1195 = v1161 - int32(1)
	goto L177
L181:
	;
	v1206 = int32(2)
	v1207 = v1162 + v1206
	v1209 = v1164 + v1206
	if v1209 != v1138&int32(2147483646) {
		v1158 = v1204
		v1161 = v1205
		v1162 = v1207
		v1164 = v1209
		goto L175
	} else {
		goto L185
	}
L182:
	;
	v1204 = v1194
	v1205 = v1195 - int32(1)
	goto L181
L183:
	;
	goto L184
L184:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1194))) = uint16(v1196)
	v1204 = v1194 + int32(2)
	v1205 = v1195
	goto L181
L185:
	;
	goto L176
L186:
	;
	v1221 = v1204
	v1224 = v1205
	v1225 = v1207
	goto L171
L187:
	;
	v1264 = v1224 - int32(1)
	goto L170
L188:
	;
	goto L189
L189:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1221))) = uint16(v1247)
	v1264 = v1224
	goto L170
L190:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v1321 = int32(0)
	if v1294 == int32(1) {
		goto L195
	} else {
		goto L196
	}
L191:
	;
	v1472 = v1294
	v1475 = v1297
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v1472
	if v1472 != 0 {
		goto L213
	} else {
		goto L214
	}
L193:
	;
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v1472 = v1440
	v1475 = v1463
	goto L192
L194:
	;
	v1426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1320+v1404<<(uint(int32(1))%32)))))
	v1428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1295+v1426))))
	if v1428 != 0 {
		goto L210
	} else {
		goto L211
	}
L195:
	;
	v1400 = v1294
	v1403 = v1320
	v1404 = v1321
	goto L194
L196:
	;
	goto L197
L197:
	;
	v1337 = v1294
	v1340 = v1320
	v1341 = v1321
	v1343 = int32(0)
	goto L198
L198:
	;
	v1362 = v1320 + v1341<<(uint(int32(1))%32)
	v1363 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1362))))
	v1365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1295+v1363))))
	if v1365 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	if v1294&int32(1) == int32(0) {
		v1440 = v1383
		goto L193
	} else {
		goto L209
	}
L200:
	;
	v1375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1362)+2)))
	v1377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1295+v1375))))
	if v1377 != 0 {
		goto L205
	} else {
		goto L206
	}
L201:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1340))) = uint16(v1363)
	v1373 = v1337
	v1374 = v1340 + int32(2)
	goto L200
L202:
	;
	goto L203
L203:
	;
	v1373 = v1337 - int32(1)
	v1374 = v1340
	goto L200
L204:
	;
	v1385 = int32(2)
	v1386 = v1341 + v1385
	v1388 = v1343 + v1385
	if v1388 != v1294&int32(2147483646) {
		v1337 = v1383
		v1340 = v1384
		v1341 = v1386
		v1343 = v1388
		goto L198
	} else {
		goto L208
	}
L205:
	;
	v1383 = v1373 - int32(1)
	v1384 = v1374
	goto L204
L206:
	;
	goto L207
L207:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1374))) = uint16(v1375)
	v1383 = v1373
	v1384 = v1374 + int32(2)
	goto L204
L208:
	;
	goto L199
L209:
	;
	v1400 = v1383
	v1403 = v1384
	v1404 = v1386
	goto L194
L210:
	;
	v1440 = v1400 - int32(1)
	goto L193
L211:
	;
	goto L212
L212:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1403))) = uint16(v1426)
	v1440 = v1400
	goto L193
L213:
	;
	v1497 = v1475
	goto L215
L214:
	;
	v1497 = int32(0)
	goto L215
L215:
	;
	if v1497 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+624)) = int32(0)
	F_gistSplitByKey(m, l0, l1, l2, l3, l4, l5, v889)
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L1
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	F_gistunionsubkey(m, l4, l2, l5)
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L1
	} else {
		goto L220
	}
L219:
	;
	goto L6
L220:
	;
	v1506 = int32(1)
	if v1114 != v1506 {
		goto L143
	} else {
		goto L221
	}
L221:
	;
	v1509 = int32(1)
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v1510 < int32(2) {
		v1561 = v1506
		v1562 = v1509
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1562<<(uint(int32(2))%32)-int32(4))))
	F_gistDeCompressAtt(m, l4, l0, v1588, v34+int32(112), v34+int32(80))
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L1
	} else {
		goto L228
	}
L223:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(l5)+624))
	v1523 = v1506
	v1524 = v1509
	goto L224
L224:
	;
	v1546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1513+v1524))))
	if v1546 != 0 {
		v1561 = v1523
		v1562 = v1524
		goto L222
	} else {
		goto L226
	}
L225:
	;
	v1561 = v1548
	v1562 = v1550
	goto L222
L226:
	;
	v1548 = v1523 + int32(1)
	v1550 = v1548 & int32(_a_F_gistSplitByKey_1)
	if base.Ui32(v1550) < base.Ui32(v1510) {
		v1523 = v1548
		v1524 = v1550
		goto L224
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1595)))
	if v1596 <= v889 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v1721 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v1720 + v1721
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1724+v1720<<(uint(v1721)%32)))) = uint16(v1561)
	goto L6
L230:
	;
	v1613 = v889
	goto L231
L231:
	;
	v1630 = v1613 << (uint(int32(3)) % 32)
	v1632 = *(*int64)(unsafe.Add(mBase, uint32(v424+v1630)))
	v1633 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+74)) = uint8(v1633)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+72)) = uint16(v1633)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+68)) = v1633
	*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v34)+56)) = v1632
	v1642 = v34 + int32(56)
	v1644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1613+v410))))
	v1649 = v34 + int32(112) + v1613*int32(24)
	v1652 = v34 + int32(80) + v1613
	v1653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1652))))
	v1654 = F_gistpenalty(m, l4, v1613, v1642, v1644, v1649, v1653)
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L1
	} else {
		goto L233
	}
L232:
	;
	goto L229
L233:
	;
	v1657 = *(*int64)(unsafe.Add(mBase, uint32(v1630+v431)))
	v1658 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+74)) = uint8(v1658)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+72)) = uint16(v1658)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+68)) = v1658
	*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v34)+56)) = v1657
	v1667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1613+v417))))
	v1668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1652))))
	v1669 = F_gistpenalty(m, l4, v1613, v1642, v1667, v1649, v1668)
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	if base.F32_ne(v1669, v1654) != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	if base.F32_gt(v1654, v1669) == int32(0) {
		goto L229
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v1685 = v1613 + int32(1)
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1686)))
	if v1685 < v1687 {
		v1613 = v1685
		goto L231
	} else {
		goto L239
	}
L238:
	;
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v1676 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v1675 + v1676
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v1679+v1675<<(uint(v1676)%32)))) = uint16(v1561)
	goto L6
L239:
	;
	goto L232
L240:
	;
	F_gistSplitByKey(m, l0, l1, l2, l3, l4, l5, v889)
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L1
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v1767 = F_palloc(m, l3<<(uint(int32(2))%32))
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L1
	} else {
		goto L244
	}
L243:
	;
	goto L6
L244:
	;
	v1769 = F_palloc(m, v46)
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	if l3 <= int32(0) {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v713)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+88)) = v1947
	v1949 = *(*int64)(unsafe.Add(mBase, uint32(v713)))
	*(*int64)(unsafe.Add(mBase, uint32(v34)+80)) = v1949
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v1952 = *(*int64)(unsafe.Add(mBase, uint32(l5)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v34)+120)) = v1952
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+128)) = v1954
	v1956 = *(*int64)(unsafe.Add(mBase, uint32(l5)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v34)+112)) = v1956
	v1959 = F_palloc_mul(m, int32(2), l3)
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L1
	} else {
		goto L264
	}
L247:
	;
	v1927 = int32(0)
	goto L246
L248:
	;
	goto L249
L249:
	;
	v1774 = int32(0)
	if l3 != int32(1) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v1792 = v1774
	v1793 = int32(0)
	v1795 = v1774
	goto L253
L251:
	;
	v1872 = v1774
	v1875 = v1774
	goto L252
L252:
	;
	v1894 = int32(1)
	v1895 = v1872 + v1894
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(l5)+624))
	v1898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1895+v1896))))
	if v1898 != v1894 {
		v1927 = v1875
		goto L246
	} else {
		goto L263
	}
L253:
	;
	v1814 = int32(1)
	v1815 = v1792 | v1814
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(l5)+624))
	v1818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1815+v1816))))
	if v1818 == v1814 {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	if l3&int32(1) == int32(0) {
		v1927 = v1857
		goto L246
	} else {
		goto L262
	}
L255:
	;
	v1821 = int32(2)
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1792<<(uint(v1821)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1767+v1795<<(uint(v1821)%32)))) = v1827
	v1829 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1769+v1795<<(uint(v1829)%32)))) = uint16(v1815)
	v1835 = v1795 + v1829
	goto L257
L256:
	;
	v1835 = v1795
	goto L257
L257:
	;
	v1837 = v1792 + int32(2)
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(l5)+624))
	v1840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1837+v1838))))
	if v1840 == int32(1) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1843 = int32(2)
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1815<<(uint(v1843)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1767+v1835<<(uint(v1843)%32)))) = v1849
	v1851 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1769+v1835<<(uint(v1851)%32)))) = uint16(v1837)
	v1857 = v1835 + v1851
	goto L260
L259:
	;
	v1857 = v1835
	goto L260
L260:
	;
	v1859 = v1793 + int32(2)
	if v1859 != l3&int32(2147483646) {
		v1792 = v1837
		v1793 = v1859
		v1795 = v1857
		goto L253
	} else {
		goto L261
	}
L261:
	;
	goto L254
L262:
	;
	v1872 = v1837
	v1875 = v1857
	goto L252
L263:
	;
	v1901 = int32(2)
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1872<<(uint(v1901)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1767+v1875<<(uint(v1901)%32)))) = v1907
	v1909 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1769+v1875<<(uint(v1909)%32)))) = uint16(v1895)
	v1927 = v1875 + v1909
	goto L246
L264:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v1963 = v1961 << (uint(int32(1)) % 32)
	if v1963 != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	base.MemoryCopy(m, v1959, v1964, v1963)
	goto L267
L266:
	;
	goto L267
L267:
	;
	v1967 = F_palloc_mul(m, int32(2), l3)
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v1971 = v1969 << (uint(int32(1)) % 32)
	if v1971 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	base.MemoryCopy(m, v1967, v1972, v1971)
	goto L271
L270:
	;
	goto L271
L271:
	;
	F_gistSplitByKey(m, l0, l1, v1767, v1927, l4, l5, v889)
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	v1976 = int32(0)
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v1976 < v1977 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1989 = v1946
	v1992 = v1976
	goto L276
L274:
	;
	v2041 = v1946
	goto L275
L275:
	;
	v2063 = int32(0)
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	if v2063 < v2064 {
		goto L279
	} else {
		goto L280
	}
L276:
	;
	v2011 = int32(1)
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v2018 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2014+v1992<<(uint(v2011)%32)))))
	v2024 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1769+v2018<<(uint(v2011)%32)-int32(2)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1959+v1989<<(uint(v2011)%32)))) = uint16(v2024)
	v2027 = v1989 + v2011
	v2029 = v1992 + v2011
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v2029 < v2030 {
		v1989 = v2027
		v1992 = v2029
		goto L276
	} else {
		goto L278
	}
L277:
	;
	v2041 = v2027
	goto L275
L278:
	;
	goto L277
L279:
	;
	v2077 = v1951
	v2079 = v2063
	goto L282
L280:
	;
	v2129 = v1951
	goto L281
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v2041
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v1959
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v34)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v713)+8)) = v2152
	v2154 = *(*int64)(unsafe.Add(mBase, uint32(v34)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v713))) = v2154
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v2129
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v1967
	v2159 = l5 + int32(28)
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v34)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v2159)+16)) = v2160
	v2162 = *(*int64)(unsafe.Add(mBase, uint32(v34)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v2159)+8)) = v2162
	v2164 = *(*int64)(unsafe.Add(mBase, uint32(v34)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v2159))) = v2164
	goto L6
L282:
	;
	v2098 = int32(1)
	v2101 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v2105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2101+v2079<<(uint(v2098)%32)))))
	v2111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1769+v2105<<(uint(v2098)%32)-int32(2)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1967+v2077<<(uint(v2098)%32)))) = uint16(v2111)
	v2114 = v2077 + v2098
	v2116 = v2079 + v2098
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	if v2116 < v2117 {
		v2077 = v2114
		v2079 = v2116
		goto L282
	} else {
		goto L284
	}
L283:
	;
	v2129 = v2114
	goto L281
L284:
	;
	goto L283
L285:
	;
	goto L5
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+624)) = int32(0)
	F_gistunionsubkey(m, l4, l2, l5)
	mBase = m.M
	v2235 = m.ExcPending
	if v2235 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	goto L4
}
func F_gist_bbox_zorder_abbrev_convert(m *base.Module, l0 int64, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	v5 = base.I32_wrap_i64(l0)
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)+24))
	v8 = int64(4294967295)
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v5)+16))
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
	return (v57<<(uint(v53)%64)|v57<<(uint(v60)%64))&int64(-6148914691236517206) | (v84<<(uint(v60)%64)|v84)&int64(6148914691236517205)
}
func F_gist_box_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v12 int32
	_ = v12
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v9 = F_box_penalty(m, v6, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v3)))) = base.F32_demote_f64(v9)
		return v3 & int64(4294967295)
	}
}
func F_gist_box_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v92 int32
	_ = v92
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v103 float64
	_ = v103
	var v109 float64
	_ = v109
	var v121 float64
	_ = v121
	var v127 float64
	_ = v127
	var v133 float64
	_ = v133
	var v135 float64
	_ = v135
	var v138 float64
	_ = v138
	var v144 float64
	_ = v144
	var v156 float64
	_ = v156
	var v162 float64
	_ = v162
	var v168 float64
	_ = v168
	var v170 float64
	_ = v170
	var v177 int32
	_ = v177
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 float64
	_ = v281
	var v283 float64
	_ = v283
	var v287 int32
	_ = v287
	var v290 float64
	_ = v290
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 float64
	_ = v306
	var v307 float64
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 float64
	_ = v311
	var v314 float64
	_ = v314
	var v316 int32
	_ = v316
	var v339 int32
	_ = v339
	var v343 float64
	_ = v343
	var v367 int32
	_ = v367
	var v368 float64
	_ = v368
	var v371 int64
	_ = v371
	var v385 float64
	_ = v385
	var v391 float64
	_ = v391
	var v393 float64
	_ = v393
	var v395 float64
	_ = v395
	var v397 int32
	_ = v397
	var v409 int32
	_ = v409
	var v434 float64
	_ = v434
	var v446 int32
	_ = v446
	var v454 int32
	_ = v454
	var v477 int32
	_ = v477
	var v505 float64
	_ = v505
	var v506 float64
	_ = v506
	var v507 int32
	_ = v507
	var v508 float64
	_ = v508
	var v511 float64
	_ = v511
	var v513 int32
	_ = v513
	var v536 int32
	_ = v536
	var v537 float64
	_ = v537
	var v564 int32
	_ = v564
	var v565 float64
	_ = v565
	var v568 int64
	_ = v568
	var v577 float64
	_ = v577
	var v584 float64
	_ = v584
	var v590 float64
	_ = v590
	var v591 float64
	_ = v591
	var v611 int32
	_ = v611
	var v630 float64
	_ = v630
	var v654 int32
	_ = v654
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v749 int32
	_ = v749
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int64
	_ = v796
	var v798 int64
	_ = v798
	var v800 int64
	_ = v800
	var v802 int64
	_ = v802
	var v804 float64
	_ = v804
	var v810 float64
	_ = v810
	var v822 float64
	_ = v822
	var v828 float64
	_ = v828
	var v840 float64
	_ = v840
	var v846 float64
	_ = v846
	var v858 float64
	_ = v858
	var v864 float64
	_ = v864
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int64
	_ = v893
	var v895 int64
	_ = v895
	var v897 int64
	_ = v897
	var v899 int64
	_ = v899
	var v901 float64
	_ = v901
	var v907 float64
	_ = v907
	var v919 float64
	_ = v919
	var v925 float64
	_ = v925
	var v937 float64
	_ = v937
	var v943 float64
	_ = v943
	var v955 float64
	_ = v955
	var v961 float64
	_ = v961
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v1014 int64
	_ = v1014
	var v1017 int64
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1042 float64
	_ = v1042
	var v1044 int64
	_ = v1044
	var v1046 float64
	_ = v1046
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1065 int32
	_ = v1065
	var v1071 int32
	_ = v1071
	var v1088 int32
	_ = v1088
	var v1090 float64
	_ = v1090
	var v1094 float64
	_ = v1094
	var v1109 int32
	_ = v1109
	var v1125 int32
	_ = v1125
	var v1128 float64
	_ = v1128
	var v1134 float64
	_ = v1134
	var v1146 float64
	_ = v1146
	var v1152 float64
	_ = v1152
	var v1164 float64
	_ = v1164
	var v1170 float64
	_ = v1170
	var v1182 float64
	_ = v1182
	var v1188 float64
	_ = v1188
	var v1199 int64
	_ = v1199
	var v1201 int64
	_ = v1201
	var v1203 int64
	_ = v1203
	var v1205 int64
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1221 float64
	_ = v1221
	var v1227 float64
	_ = v1227
	var v1239 float64
	_ = v1239
	var v1245 float64
	_ = v1245
	var v1257 float64
	_ = v1257
	var v1263 float64
	_ = v1263
	var v1275 float64
	_ = v1275
	var v1281 float64
	_ = v1281
	var v1292 int64
	_ = v1292
	var v1294 int64
	_ = v1294
	var v1296 int64
	_ = v1296
	var v1298 int64
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1324 int32
	_ = v1324
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1359 int32
	_ = v1359
	var v1360 float64
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 float64
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1365 float64
	_ = v1365
	var v1368 float64
	_ = v1368
	var v1377 float64
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1380 float64
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1390 int32
	_ = v1390
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1404 int32
	_ = v1404
	var v1411 int32
	_ = v1411
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1438 float64
	_ = v1438
	var v1444 float64
	_ = v1444
	var v1456 float64
	_ = v1456
	var v1462 float64
	_ = v1462
	var v1474 float64
	_ = v1474
	var v1480 float64
	_ = v1480
	var v1492 float64
	_ = v1492
	var v1498 float64
	_ = v1498
	var v1509 int64
	_ = v1509
	var v1511 int64
	_ = v1511
	var v1513 int64
	_ = v1513
	var v1515 int64
	_ = v1515
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1524 int32
	_ = v1524
	var v1529 int32
	_ = v1529
	var v1534 float64
	_ = v1534
	var v1540 float64
	_ = v1540
	var v1552 float64
	_ = v1552
	var v1558 float64
	_ = v1558
	var v1570 float64
	_ = v1570
	var v1576 float64
	_ = v1576
	var v1588 float64
	_ = v1588
	var v1594 float64
	_ = v1594
	var v1605 int64
	_ = v1605
	var v1607 int64
	_ = v1607
	var v1609 int64
	_ = v1609
	var v1611 int64
	_ = v1611
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1625 float64
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 float64
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1633 float64
	_ = v1633
	var v1639 float64
	_ = v1639
	var v1651 float64
	_ = v1651
	var v1657 float64
	_ = v1657
	var v1669 float64
	_ = v1669
	var v1675 float64
	_ = v1675
	var v1687 float64
	_ = v1687
	var v1693 float64
	_ = v1693
	var v1704 int64
	_ = v1704
	var v1706 int64
	_ = v1706
	var v1708 int64
	_ = v1708
	var v1710 int64
	_ = v1710
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1719 int32
	_ = v1719
	var v1724 int32
	_ = v1724
	var v1727 float64
	_ = v1727
	var v1733 float64
	_ = v1733
	var v1745 float64
	_ = v1745
	var v1751 float64
	_ = v1751
	var v1763 float64
	_ = v1763
	var v1769 float64
	_ = v1769
	var v1781 float64
	_ = v1781
	var v1787 float64
	_ = v1787
	var v1798 int64
	_ = v1798
	var v1800 int64
	_ = v1800
	var v1802 int64
	_ = v1802
	var v1804 int64
	_ = v1804
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1813 int32
	_ = v1813
	var v1823 int32
	_ = v1823
	var v1825 int32
	_ = v1825
	var v1879 int64
	_ = v1879
	v6 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(96)
	m.G0 = v29
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	base.MemoryFill(m, v29+int32(8), v6, int32(88))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v39 = int32(_a_F_gist_box_picksplit_0)
	v40 = v38 + v39
	v42 = v40 & v39
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v42
	v45 = v42 - int32(1)
	v47 = v42 << (uint(int32(4)) % 32)
	v48 = F_palloc(m, v47)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v52 = F_palloc(m, v47)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v55 = v38 & int32(_a_F_gist_box_picksplit_0)
	if v55 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v61 = v29 + int32(16)
	v73 = int32(1)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v207 = base.I32_wrap_i64(v31)
	v208 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+48)) = uint8(v208)
	v211 = v32 + int32(8)
	v213 = v45 << (uint(int32(4)) % 32)
	v222 = v208
	v226 = v6
	goto L36
L7:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(8)+v73*int32(24))))
	if v73 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v177 = (v73 + int32(1)) & int32(_a_F_gist_box_picksplit_0)
	if base.Ui32(v177) <= base.Ui32(v42) {
		v73 = v177
		goto L7
	} else {
		goto L35
	}
L10:
	;
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v92)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v61)+24)) = v95
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v92)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v61)+16)) = v97
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v61)+8)) = v99
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v92)))
	*(*int64)(unsafe.Add(mBase, uint32(v61))) = v101
	goto L9
L11:
	;
	goto L12
L12:
	;
	v103 = *(*float64)(unsafe.Add(mBase, uint32(v29)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v103)&int64(9223372036854775807)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v121 = *(*float64)(unsafe.Add(mBase, uint32(v92)+16))
	if base.Ui64(base.I64_reinterpret_f64(v121)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v109 = *(*float64)(unsafe.Add(mBase, uint32(v92)))
	if base.B2i32(base.F64_lt(v103, v109) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v109)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+16)) = v109
	goto L13
L16:
	;
	v127 = *(*float64)(unsafe.Add(mBase, uint32(v29)+32))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v127)&int64(9223372036854775807)) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v138 = *(*float64)(unsafe.Add(mBase, uint32(v29)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v138)&int64(9223372036854775807)) {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v133 = v121
	goto L21
L20:
	;
	v133 = v127
	goto L21
L21:
	;
	if base.F64_lt(v121, v127) != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v135 = v121
	goto L24
L23:
	;
	v135 = v133
	goto L24
L24:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+32)) = v135
	goto L18
L25:
	;
	v156 = *(*float64)(unsafe.Add(mBase, uint32(v92)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v156)&int64(9223372036854775807)) {
		goto L9
	} else {
		goto L28
	}
L26:
	;
	v144 = *(*float64)(unsafe.Add(mBase, uint32(v92)+8))
	if base.B2i32(base.F64_lt(v138, v144) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v144)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+24)) = v144
	goto L25
L28:
	;
	v162 = *(*float64)(unsafe.Add(mBase, uint32(v29)+40))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v162)&int64(9223372036854775807)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v168 = v156
	goto L31
L30:
	;
	v168 = v162
	goto L31
L31:
	;
	if base.F64_lt(v156, v162) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v170 = v156
	goto L34
L33:
	;
	v170 = v168
	goto L34
L34:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+40)) = v170
	goto L9
L35:
	;
	goto L8
L36:
	;
	v243 = int32(1)
	if v55 != v243 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+48)))
	if v721 == int32(1) {
		goto L125
	} else {
		goto L126
	}
L38:
	;
	if v222 != 0 {
		v222 = int32(0)
		v226 = int32(1)
		goto L36
	} else {
		goto L123
	}
L39:
	;
	v246 = v243
	goto L42
L40:
	;
	goto L41
L41:
	;
	if v47 != 0 {
		goto L118
	} else {
		goto L119
	}
L42:
	;
	v274 = v48 + v246<<(uint(int32(4))%32)
	v276 = v274 - int32(16)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v211+v246*int32(24))))
	if v222 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	if v47 != 0 {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	v290 = *(*float64)(unsafe.Add(mBase, uint32(v287)))
	*(*float64)(unsafe.Add(mBase, uint32(v274-int32(8)))) = v290
	v295 = (v246 + int32(1)) & int32(_a_F_gist_box_picksplit_0)
	if base.Ui32(v295) <= base.Ui32(v42) {
		v246 = v295
		goto L42
	} else {
		goto L48
	}
L45:
	;
	v281 = *(*float64)(unsafe.Add(mBase, uint32(v280)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v276))) = v281
	v287 = v280
	goto L44
L46:
	;
	goto L47
L47:
	;
	v283 = *(*float64)(unsafe.Add(mBase, uint32(v280)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v276))) = v283
	v287 = v280 + int32(8)
	goto L44
L48:
	;
	goto L43
L49:
	;
	base.MemoryCopy(m, v52, v48, v47)
	goto L51
L50:
	;
	goto L51
L51:
	;
	F_pg_qsort(m, v48, v42, int32(16), int32(107))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_pg_qsort(m, v52, v42, int32(16), int32(108))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v306 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	v307 = *(*float64)(unsafe.Add(mBase, uint32(v52)))
	v308 = int32(0)
	v310 = v308
	v311 = v306
	v314 = v307
	v316 = v308
	goto L54
L54:
	;
	v339 = v310
	v343 = v314
	goto L57
L55:
	;
	v505 = *(*float64)(unsafe.Add(mBase, uint32(v213+v48)+8))
	v506 = *(*float64)(unsafe.Add(mBase, uint32(v52+v213)+8))
	v507 = v45
	v508 = v505
	v511 = v506
	v513 = v45
	goto L87
L56:
	;
	goto L55
L57:
	;
	v367 = v48 + v339<<(uint(int32(4))%32)
	v368 = *(*float64)(unsafe.Add(mBase, uint32(v367)))
	v371 = base.I64_reinterpret_f64(v368) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v311)&int64(9223372036854775807)) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	if v42 <= v316 {
		v454 = v316
		goto L76
	} else {
		goto L77
	}
L59:
	;
	goto L58
L60:
	;
	if base.Ui64(base.I64_reinterpret_f64(v343)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L66
	} else {
		goto L67
	}
L61:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v371) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v371))|base.F64_ne(v368, v311) != 0 {
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
	v385 = *(*float64)(unsafe.Add(mBase, uint32(v367)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v385)&int64(9223372036854775807)) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v395 = v343
	goto L68
L68:
	;
	v397 = v339 + int32(1)
	if v397 < v42 {
		v339 = v397
		v343 = v395
		goto L57
	} else {
		goto L75
	}
L69:
	;
	v391 = v385
	goto L71
L70:
	;
	v391 = v343
	goto L71
L71:
	;
	if base.F64_gt(v385, v343) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v393 = v385
	goto L74
L73:
	;
	v393 = v391
	goto L74
L74:
	;
	v395 = v393
	goto L68
L75:
	;
	goto L56
L76:
	;
	F_g_box_consider_split(m, v29+int32(8), v226, v368, v339, v343, v454)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L85
	}
L77:
	;
	v409 = v316
	goto L78
L78:
	;
	if base.Ui64(base.I64_reinterpret_f64(v343)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v454 = v42
	goto L76
L80:
	;
	v434 = *(*float64)(unsafe.Add(mBase, uint32(v52+v409<<(uint(int32(4))%32))+8))
	if base.B2i32(base.F64_le(v434, v343) == int32(0))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v434)&int64(9223372036854775807))) != 0 {
		v454 = v409
		goto L76
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v446 = v409 + int32(1)
	if v446 != v42 {
		v409 = v446
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
	if v339 < v42 {
		v310 = v339
		v311 = v368
		v314 = v343
		v316 = v454
		goto L54
	} else {
		goto L86
	}
L86:
	;
	goto L56
L87:
	;
	v536 = v507
	v537 = v508
	goto L89
L88:
	;
	goto L38
L89:
	;
	v564 = v52 + v536<<(uint(int32(4))%32)
	v565 = *(*float64)(unsafe.Add(mBase, uint32(v564)+8))
	v568 = base.I64_reinterpret_f64(v565) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v511)&int64(9223372036854775807)) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	if v513 < int32(0) {
		v654 = v513
		goto L108
	} else {
		goto L109
	}
L91:
	;
	goto L90
L92:
	;
	v577 = *(*float64)(unsafe.Add(mBase, uint32(v564)))
	if base.Ui64(base.I64_reinterpret_f64(v577)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L98
	} else {
		goto L99
	}
L93:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v568) {
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v568))|base.F64_ne(v511, v565) != 0 {
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
	if base.F64_gt(v537, v577) != 0 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v591 = v537
	goto L100
L100:
	;
	if int32(0) < v536 {
		v536 = v536 - int32(1)
		v537 = v591
		goto L89
	} else {
		goto L107
	}
L101:
	;
	v584 = v577
	goto L103
L102:
	;
	v584 = v537
	goto L103
L103:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v537)&int64(9223372036854775807)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v590 = v577
	goto L106
L105:
	;
	v590 = v584
	goto L106
L106:
	;
	v591 = v590
	goto L100
L107:
	;
	goto L38
L108:
	;
	v676 = int32(1)
	F_g_box_consider_split(m, v29+int32(8), v226, v537, v654+v676, v565, v536+v676)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L116
	}
L109:
	;
	v611 = v513
	goto L110
L110:
	;
	v630 = *(*float64)(unsafe.Add(mBase, uint32(v48+v611<<(uint(int32(4))%32))))
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v630)&int64(9223372036854775807)))|base.B2i32(base.Ui64(base.I64_reinterpret_f64(v537)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))&base.F64_le(v537, v630) == int32(0) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v654 = int32(-1)
	goto L108
L112:
	;
	v654 = v611
	goto L108
L113:
	;
	goto L114
L114:
	;
	if int32(0) < v611 {
		v611 = v611 - int32(1)
		goto L110
	} else {
		goto L115
	}
L115:
	;
	goto L111
L116:
	;
	if int32(0) <= v536 {
		v507 = v536
		v508 = v537
		v511 = v565
		v513 = v654
		goto L87
	} else {
		goto L117
	}
L117:
	;
	goto L88
L118:
	;
	base.MemoryCopy(m, v52, v48, v47)
	goto L120
L119:
	;
	goto L120
L120:
	;
	F_pg_qsort(m, v48, v42, int32(16), int32(107))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_pg_qsort(m, v52, v42, int32(16), int32(108))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v207)+32)) = v1879
	m.G0 = v29 + int32(96)
	return v31 & int64(4294967295)
L125:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v725 = int32(_a_F_gist_box_picksplit_0)
	v728 = (v724 + v725) & v725
	v732 = v728<<(uint(int32(1))%32) + int32(4)
	v733 = F_palloc(m, v732)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v1020 = v42 << (uint(int32(1)) % 32)
	v1021 = F_palloc(m, v1020)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L172
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207))) = v733
	v737 = F_palloc(m, v732)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v739 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+24)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v207)+20)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v739
	if v724&int32(_a_F_gist_box_picksplit_0) != int32(1) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v749 = int32(1)
	v760 = int32(0)
	v764 = v749
	v766 = int32(0)
	v767 = v749
	goto L133
L131:
	;
	v1014 = int64(0)
	v1017 = int64(0)
	goto L132
L132:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v207)+8)) = v1017
	v1879 = v1014
	goto L124
L133:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v211+v764*int32(24))))
	if base.Ui32(v764) <= base.Ui32(int32(base.Ui32(v728)>>(uint(v749)%32))) {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	v1014 = base.I64_extend_i32_u(v981)
	v1017 = base.I64_extend_i32_u(v982)
	goto L132
L135:
	;
	v984 = v767 + int32(1)
	v986 = v984 & int32(_a_F_gist_box_picksplit_0)
	if base.Ui32(v986) <= base.Ui32(v728) {
		v760 = v981
		v764 = v986
		v766 = v982
		v767 = v984
		goto L133
	} else {
		goto L171
	}
L136:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v785+v786<<(uint(int32(1))%32)))) = uint16(v767)
	if v766 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	goto L138
L138:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v207)+20))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v207)+24))
	*(*uint16)(unsafe.Add(mBase, uint32(v882+v883<<(uint(int32(1))%32)))) = uint16(v767)
	if v760 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L139:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v878 + int32(1)
	v981 = v760
	v982 = v877
	goto L135
L140:
	;
	v794 = F_palloc(m, int32(32))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v804 = *(*float64)(unsafe.Add(mBase, uint32(v766)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v804)&int64(9223372036854775807)) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v796 = *(*int64)(unsafe.Add(mBase, uint32(v783)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v794)+24)) = v796
	v798 = *(*int64)(unsafe.Add(mBase, uint32(v783)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v794)+16)) = v798
	v800 = *(*int64)(unsafe.Add(mBase, uint32(v783)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v794)+8)) = v800
	v802 = *(*int64)(unsafe.Add(mBase, uint32(v783)))
	*(*int64)(unsafe.Add(mBase, uint32(v794))) = v802
	v877 = v794
	goto L139
L144:
	;
	v822 = *(*float64)(unsafe.Add(mBase, uint32(v783)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v822)&int64(9223372036854775807)) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v810 = *(*float64)(unsafe.Add(mBase, uint32(v783)))
	if base.B2i32(base.F64_lt(v804, v810) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v810)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v766))) = v810
	goto L144
L147:
	;
	v840 = *(*float64)(unsafe.Add(mBase, uint32(v766)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v840)&int64(9223372036854775807)) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v828 = *(*float64)(unsafe.Add(mBase, uint32(v766)+16))
	if base.B2i32(base.F64_gt(v828, v822) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v828)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v766)+16)) = v822
	goto L147
L150:
	;
	v858 = *(*float64)(unsafe.Add(mBase, uint32(v783)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v858)&int64(9223372036854775807)) {
		v877 = v766
		goto L139
	} else {
		goto L153
	}
L151:
	;
	v846 = *(*float64)(unsafe.Add(mBase, uint32(v783)+8))
	if base.B2i32(base.F64_lt(v840, v846) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v846)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v766)+8)) = v846
	goto L150
L153:
	;
	v864 = *(*float64)(unsafe.Add(mBase, uint32(v766)+24))
	if base.B2i32(base.F64_gt(v864, v858) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v864)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		v877 = v766
		goto L139
	} else {
		goto L154
	}
L154:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v766)+24)) = v858
	v877 = v766
	goto L139
L155:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v207)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+24)) = v975 + int32(1)
	v981 = v974
	v982 = v766
	goto L135
L156:
	;
	v891 = F_palloc(m, int32(32))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L1
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v901 = *(*float64)(unsafe.Add(mBase, uint32(v760)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v901)&int64(9223372036854775807)) {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v893 = *(*int64)(unsafe.Add(mBase, uint32(v783)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v891)+24)) = v893
	v895 = *(*int64)(unsafe.Add(mBase, uint32(v783)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v891)+16)) = v895
	v897 = *(*int64)(unsafe.Add(mBase, uint32(v783)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v891)+8)) = v897
	v899 = *(*int64)(unsafe.Add(mBase, uint32(v783)))
	*(*int64)(unsafe.Add(mBase, uint32(v891))) = v899
	v974 = v891
	goto L155
L160:
	;
	v919 = *(*float64)(unsafe.Add(mBase, uint32(v783)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v919)&int64(9223372036854775807)) {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	v907 = *(*float64)(unsafe.Add(mBase, uint32(v783)))
	if base.B2i32(base.F64_lt(v901, v907) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v907)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L160
	} else {
		goto L162
	}
L162:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v760))) = v907
	goto L160
L163:
	;
	v937 = *(*float64)(unsafe.Add(mBase, uint32(v760)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v937)&int64(9223372036854775807)) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	v925 = *(*float64)(unsafe.Add(mBase, uint32(v760)+16))
	if base.B2i32(base.F64_gt(v925, v919) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v925)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v760)+16)) = v919
	goto L163
L166:
	;
	v955 = *(*float64)(unsafe.Add(mBase, uint32(v783)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v955)&int64(9223372036854775807)) {
		v974 = v760
		goto L155
	} else {
		goto L169
	}
L167:
	;
	v943 = *(*float64)(unsafe.Add(mBase, uint32(v783)+8))
	if base.B2i32(base.F64_lt(v937, v943) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v943)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v760)+8)) = v943
	goto L166
L169:
	;
	v961 = *(*float64)(unsafe.Add(mBase, uint32(v760)+24))
	if base.B2i32(base.F64_gt(v961, v955) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v961)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		v974 = v760
		goto L155
	} else {
		goto L170
	}
L170:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v760)+24)) = v955
	v974 = v760
	goto L155
L171:
	;
	goto L134
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207))) = v1021
	v1025 = F_palloc(m, v1020)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v1027 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+24)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v207)+20)) = v1025
	v1033 = F_palloc0(m, int32(32))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v1036 = F_palloc0(m, int32(32))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v1038 = F_palloc(m, v47)
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	if v55 == int32(1) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v207)+8)) = base.I64_extend_i32_u(v1033)
	v1879 = base.I64_extend_i32_u(v1036)
	goto L124
L178:
	;
	v1042 = *(*float64)(unsafe.Add(mBase, uint32(v29)+64))
	v1044 = int64(9223372036854775807)
	v1046 = *(*float64)(unsafe.Add(mBase, uint32(v29)+56))
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	if v1052 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v1053 = int32(24)
	goto L181
L180:
	;
	v1053 = int32(16)
	goto L181
L181:
	;
	if v1052 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v1056 = int32(8)
	goto L184
L183:
	;
	v1056 = int32(0)
	goto L184
L184:
	;
	v1057 = int32(1)
	v1059 = v1057
	v1065 = v1057
	v1071 = int32(0)
	goto L185
L185:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v211+v1059*int32(24))))
	v1090 = *(*float64)(unsafe.Add(mBase, uint32(v1088+v1053)))
	if base.Ui64(base.I64_reinterpret_f64(v1046)&v1044) <= base.Ui64(int64(9218868437227405312)) {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	if v1314 <= int32(0) {
		goto L177
	} else {
		goto L227
	}
L187:
	;
	v1316 = v1065 + int32(1)
	v1317 = int32(_a_F_gist_box_picksplit_0)
	v1318 = v1316 & v1317
	if base.Ui32(v1318) <= base.Ui32(v40&v1317) {
		v1059 = v1318
		v1065 = v1316
		v1071 = v1314
		goto L185
	} else {
		goto L226
	}
L188:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v207)+24))
	if int32(0) < v1218 {
		goto L212
	} else {
		goto L213
	}
L189:
	;
	v1094 = *(*float64)(unsafe.Add(mBase, uint32(v1088+v1056)))
	if base.B2i32(base.F64_le(v1094, v1046) == int32(0))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1094)&int64(9223372036854775807))) != 0 {
		goto L188
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v1109 = int32(0)
	if base.B2i32(base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1042)&v1044) < base.Ui64(int64(9218868437227405313)))&base.F64_ge(v1090, v1042) == v1109)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1090)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) == v1109 {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	goto L191
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1038+v1071<<(uint(int32(4))%32)))) = v1059
	v1314 = v1071 + int32(1)
	goto L187
L194:
	;
	goto L195
L195:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	if int32(0) < v1125 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v1210 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v1209 + v1210
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1213+v1209<<(uint(v1210)%32)))) = uint16(v1065)
	v1314 = v1071
	goto L187
L197:
	;
	v1128 = *(*float64)(unsafe.Add(mBase, uint32(v1033)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1128)&int64(9223372036854775807)) {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	goto L199
L199:
	;
	v1199 = *(*int64)(unsafe.Add(mBase, uint32(v1088)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1033)+24)) = v1199
	v1201 = *(*int64)(unsafe.Add(mBase, uint32(v1088)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1033)+16)) = v1201
	v1203 = *(*int64)(unsafe.Add(mBase, uint32(v1088)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1033)+8)) = v1203
	v1205 = *(*int64)(unsafe.Add(mBase, uint32(v1088)))
	*(*int64)(unsafe.Add(mBase, uint32(v1033))) = v1205
	goto L196
L200:
	;
	v1146 = *(*float64)(unsafe.Add(mBase, uint32(v1088)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1146)&int64(9223372036854775807)) {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	v1134 = *(*float64)(unsafe.Add(mBase, uint32(v1088)))
	if base.B2i32(base.F64_lt(v1128, v1134) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1134)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1033))) = v1134
	goto L200
L203:
	;
	v1164 = *(*float64)(unsafe.Add(mBase, uint32(v1033)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1164)&int64(9223372036854775807)) {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	v1152 = *(*float64)(unsafe.Add(mBase, uint32(v1033)+16))
	if base.B2i32(base.F64_gt(v1152, v1146) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1152)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L203
	} else {
		goto L205
	}
L205:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1033)+16)) = v1146
	goto L203
L206:
	;
	v1182 = *(*float64)(unsafe.Add(mBase, uint32(v1088)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1182)&int64(9223372036854775807)) {
		goto L196
	} else {
		goto L209
	}
L207:
	;
	v1170 = *(*float64)(unsafe.Add(mBase, uint32(v1088)+8))
	if base.B2i32(base.F64_lt(v1164, v1170) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1170)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L206
	} else {
		goto L208
	}
L208:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1033)+8)) = v1170
	goto L206
L209:
	;
	v1188 = *(*float64)(unsafe.Add(mBase, uint32(v1033)+24))
	if base.B2i32(base.F64_gt(v1188, v1182) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1188)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L196
	} else {
		goto L210
	}
L210:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1033)+24)) = v1182
	goto L196
L211:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v207)+24))
	v1303 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+24)) = v1302 + v1303
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v207)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v1306+v1302<<(uint(v1303)%32)))) = uint16(v1065)
	v1314 = v1071
	goto L187
L212:
	;
	v1221 = *(*float64)(unsafe.Add(mBase, uint32(v1036)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1221)&int64(9223372036854775807)) {
		goto L215
	} else {
		goto L216
	}
L213:
	;
	goto L214
L214:
	;
	v1292 = *(*int64)(unsafe.Add(mBase, uint32(v1088)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1036)+24)) = v1292
	v1294 = *(*int64)(unsafe.Add(mBase, uint32(v1088)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1036)+16)) = v1294
	v1296 = *(*int64)(unsafe.Add(mBase, uint32(v1088)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1036)+8)) = v1296
	v1298 = *(*int64)(unsafe.Add(mBase, uint32(v1088)))
	*(*int64)(unsafe.Add(mBase, uint32(v1036))) = v1298
	goto L211
L215:
	;
	v1239 = *(*float64)(unsafe.Add(mBase, uint32(v1088)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1239)&int64(9223372036854775807)) {
		goto L218
	} else {
		goto L219
	}
L216:
	;
	v1227 = *(*float64)(unsafe.Add(mBase, uint32(v1088)))
	if base.B2i32(base.F64_lt(v1221, v1227) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1227)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1036))) = v1227
	goto L215
L218:
	;
	v1257 = *(*float64)(unsafe.Add(mBase, uint32(v1036)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1257)&int64(9223372036854775807)) {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	v1245 = *(*float64)(unsafe.Add(mBase, uint32(v1036)+16))
	if base.B2i32(base.F64_gt(v1245, v1239) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1245)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L218
	} else {
		goto L220
	}
L220:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1036)+16)) = v1239
	goto L218
L221:
	;
	v1275 = *(*float64)(unsafe.Add(mBase, uint32(v1088)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1275)&int64(9223372036854775807)) {
		goto L211
	} else {
		goto L224
	}
L222:
	;
	v1263 = *(*float64)(unsafe.Add(mBase, uint32(v1088)+8))
	if base.B2i32(base.F64_lt(v1257, v1263) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1263)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L221
	} else {
		goto L223
	}
L223:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1036)+8)) = v1263
	goto L221
L224:
	;
	v1281 = *(*float64)(unsafe.Add(mBase, uint32(v1036)+24))
	if base.B2i32(base.F64_gt(v1281, v1275) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1281)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L211
	} else {
		goto L225
	}
L225:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1036)+24)) = v1275
	goto L211
L226:
	;
	goto L186
L227:
	;
	v1324 = int32(0)
	v1332 = v1324
	v1336 = v1324
	goto L228
L228:
	;
	v1354 = v1038 + v1332<<(uint(int32(4))%32)
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1354)))
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v211+v1355*int32(24))))
	v1360 = F_box_penalty(m, v1033, v1359)
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L1
	} else {
		goto L231
	}
L229:
	;
	F_pg_qsort(m, v1038, v1314, int32(16), int32(109))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L1
	} else {
		goto L237
	}
L230:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1354)+8)) = v1380
	v1383 = v1336 + int32(1)
	v1385 = v1383 & int32(_a_F_gist_box_picksplit_0)
	if base.Ui32(v1385) < base.Ui32(v1314) {
		v1332 = v1385
		v1336 = v1383
		goto L228
	} else {
		goto L236
	}
L231:
	;
	v1362 = F_box_penalty(m, v1036, v1359)
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v1365 = base.F64_abs(base.F64_sub(v1360, v1362))
	if base.F64_ne(v1365, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v1380 = v1365
		goto L230
	} else {
		goto L233
	}
L233:
	;
	v1368 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v1360), v1368)|base.F64_eq(base.F64_abs(v1362), v1368) != 0 {
		v1380 = v1368
		goto L230
	} else {
		goto L234
	}
L234:
	;
	v1377 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	v1380 = base.F64_abs(v1377)
	goto L230
L236:
	;
	goto L229
L237:
	;
	v1395 = base.I32_trunc_sat_f64_s(base.F64_ceil(base.F64_mul(base.F64_convert_i32_u(v42), float64(0.3))))
	v1396 = int32(0)
	v1404 = v1396
	v1411 = v1396
	goto L238
L238:
	;
	v1426 = v1038 + v1411<<(uint(int32(4))%32)
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1426)))
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v211+v1427*int32(24))))
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v1433 = v1314 - v1411
	if v1432+v1433 <= v1395 {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	goto L177
L240:
	;
	v1823 = v1404 + int32(1)
	v1825 = v1823 & int32(_a_F_gist_box_picksplit_0)
	if base.Ui32(v1825) < base.Ui32(v1314) {
		v1404 = v1823
		v1411 = v1825
		goto L238
	} else {
		goto L312
	}
L241:
	;
	if int32(0) < v1432 {
		goto L245
	} else {
		goto L246
	}
L242:
	;
	goto L243
L243:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v207)+24))
	if v1433+v1529 <= v1395 {
		goto L259
	} else {
		goto L260
	}
L244:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1426)))
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v1521 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v1520 + v1521
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1524+v1520<<(uint(v1521)%32)))) = uint16(v1519)
	goto L240
L245:
	;
	v1438 = *(*float64)(unsafe.Add(mBase, uint32(v1033)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1438)&int64(9223372036854775807)) {
		goto L248
	} else {
		goto L249
	}
L246:
	;
	goto L247
L247:
	;
	v1509 = *(*int64)(unsafe.Add(mBase, uint32(v1431)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1033)+24)) = v1509
	v1511 = *(*int64)(unsafe.Add(mBase, uint32(v1431)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1033)+16)) = v1511
	v1513 = *(*int64)(unsafe.Add(mBase, uint32(v1431)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1033)+8)) = v1513
	v1515 = *(*int64)(unsafe.Add(mBase, uint32(v1431)))
	*(*int64)(unsafe.Add(mBase, uint32(v1033))) = v1515
	goto L244
L248:
	;
	v1456 = *(*float64)(unsafe.Add(mBase, uint32(v1431)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1456)&int64(9223372036854775807)) {
		goto L251
	} else {
		goto L252
	}
L249:
	;
	v1444 = *(*float64)(unsafe.Add(mBase, uint32(v1431)))
	if base.B2i32(base.F64_lt(v1438, v1444) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1444)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L248
	} else {
		goto L250
	}
L250:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1033))) = v1444
	goto L248
L251:
	;
	v1474 = *(*float64)(unsafe.Add(mBase, uint32(v1033)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1474)&int64(9223372036854775807)) {
		goto L254
	} else {
		goto L255
	}
L252:
	;
	v1462 = *(*float64)(unsafe.Add(mBase, uint32(v1033)+16))
	if base.B2i32(base.F64_gt(v1462, v1456) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1462)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L251
	} else {
		goto L253
	}
L253:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1033)+16)) = v1456
	goto L251
L254:
	;
	v1492 = *(*float64)(unsafe.Add(mBase, uint32(v1431)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1492)&int64(9223372036854775807)) {
		goto L244
	} else {
		goto L257
	}
L255:
	;
	v1480 = *(*float64)(unsafe.Add(mBase, uint32(v1431)+8))
	if base.B2i32(base.F64_lt(v1474, v1480) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1480)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L254
	} else {
		goto L256
	}
L256:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1033)+8)) = v1480
	goto L254
L257:
	;
	v1498 = *(*float64)(unsafe.Add(mBase, uint32(v1033)+24))
	if base.B2i32(base.F64_gt(v1498, v1492) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1498)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L244
	} else {
		goto L258
	}
L258:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1033)+24)) = v1492
	goto L244
L259:
	;
	if int32(0) < v1529 {
		goto L263
	} else {
		goto L264
	}
L260:
	;
	goto L261
L261:
	;
	v1625 = F_box_penalty(m, v1033, v1431)
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L1
	} else {
		goto L277
	}
L262:
	;
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v1426)))
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v207)+24))
	v1617 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+24)) = v1616 + v1617
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v207)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v1620+v1616<<(uint(v1617)%32)))) = uint16(v1615)
	goto L240
L263:
	;
	v1534 = *(*float64)(unsafe.Add(mBase, uint32(v1036)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1534)&int64(9223372036854775807)) {
		goto L266
	} else {
		goto L267
	}
L264:
	;
	goto L265
L265:
	;
	v1605 = *(*int64)(unsafe.Add(mBase, uint32(v1431)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1036)+24)) = v1605
	v1607 = *(*int64)(unsafe.Add(mBase, uint32(v1431)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1036)+16)) = v1607
	v1609 = *(*int64)(unsafe.Add(mBase, uint32(v1431)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1036)+8)) = v1609
	v1611 = *(*int64)(unsafe.Add(mBase, uint32(v1431)))
	*(*int64)(unsafe.Add(mBase, uint32(v1036))) = v1611
	goto L262
L266:
	;
	v1552 = *(*float64)(unsafe.Add(mBase, uint32(v1431)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1552)&int64(9223372036854775807)) {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	v1540 = *(*float64)(unsafe.Add(mBase, uint32(v1431)))
	if base.B2i32(base.F64_lt(v1534, v1540) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1540)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L266
	} else {
		goto L268
	}
L268:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1036))) = v1540
	goto L266
L269:
	;
	v1570 = *(*float64)(unsafe.Add(mBase, uint32(v1036)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1570)&int64(9223372036854775807)) {
		goto L272
	} else {
		goto L273
	}
L270:
	;
	v1558 = *(*float64)(unsafe.Add(mBase, uint32(v1036)+16))
	if base.B2i32(base.F64_gt(v1558, v1552) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1558)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L269
	} else {
		goto L271
	}
L271:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1036)+16)) = v1552
	goto L269
L272:
	;
	v1588 = *(*float64)(unsafe.Add(mBase, uint32(v1431)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1588)&int64(9223372036854775807)) {
		goto L262
	} else {
		goto L275
	}
L273:
	;
	v1576 = *(*float64)(unsafe.Add(mBase, uint32(v1431)+8))
	if base.B2i32(base.F64_lt(v1570, v1576) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1576)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L272
	} else {
		goto L274
	}
L274:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1036)+8)) = v1576
	goto L272
L275:
	;
	v1594 = *(*float64)(unsafe.Add(mBase, uint32(v1036)+24))
	if base.B2i32(base.F64_gt(v1594, v1588) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1594)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L262
	} else {
		goto L276
	}
L276:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1036)+24)) = v1588
	goto L262
L277:
	;
	v1627 = F_box_penalty(m, v1036, v1431)
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	if base.F64_lt(v1625, v1627) != 0 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	if int32(0) < v1630 {
		goto L283
	} else {
		goto L284
	}
L280:
	;
	goto L281
L281:
	;
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v207)+24))
	if int32(0) < v1724 {
		goto L298
	} else {
		goto L299
	}
L282:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1426)))
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v1716 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v1715 + v1716
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1719+v1715<<(uint(v1716)%32)))) = uint16(v1714)
	goto L240
L283:
	;
	v1633 = *(*float64)(unsafe.Add(mBase, uint32(v1033)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1633)&int64(9223372036854775807)) {
		goto L286
	} else {
		goto L287
	}
L284:
	;
	goto L285
L285:
	;
	v1704 = *(*int64)(unsafe.Add(mBase, uint32(v1431)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1033)+24)) = v1704
	v1706 = *(*int64)(unsafe.Add(mBase, uint32(v1431)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1033)+16)) = v1706
	v1708 = *(*int64)(unsafe.Add(mBase, uint32(v1431)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1033)+8)) = v1708
	v1710 = *(*int64)(unsafe.Add(mBase, uint32(v1431)))
	*(*int64)(unsafe.Add(mBase, uint32(v1033))) = v1710
	goto L282
L286:
	;
	v1651 = *(*float64)(unsafe.Add(mBase, uint32(v1431)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1651)&int64(9223372036854775807)) {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	v1639 = *(*float64)(unsafe.Add(mBase, uint32(v1431)))
	if base.B2i32(base.F64_lt(v1633, v1639) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1639)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L286
	} else {
		goto L288
	}
L288:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1033))) = v1639
	goto L286
L289:
	;
	v1669 = *(*float64)(unsafe.Add(mBase, uint32(v1033)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1669)&int64(9223372036854775807)) {
		goto L292
	} else {
		goto L293
	}
L290:
	;
	v1657 = *(*float64)(unsafe.Add(mBase, uint32(v1033)+16))
	if base.B2i32(base.F64_gt(v1657, v1651) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1657)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L289
	} else {
		goto L291
	}
L291:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1033)+16)) = v1651
	goto L289
L292:
	;
	v1687 = *(*float64)(unsafe.Add(mBase, uint32(v1431)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1687)&int64(9223372036854775807)) {
		goto L282
	} else {
		goto L295
	}
L293:
	;
	v1675 = *(*float64)(unsafe.Add(mBase, uint32(v1431)+8))
	if base.B2i32(base.F64_lt(v1669, v1675) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1675)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L292
	} else {
		goto L294
	}
L294:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1033)+8)) = v1675
	goto L292
L295:
	;
	v1693 = *(*float64)(unsafe.Add(mBase, uint32(v1033)+24))
	if base.B2i32(base.F64_gt(v1693, v1687) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1693)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L282
	} else {
		goto L296
	}
L296:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1033)+24)) = v1687
	goto L282
L297:
	;
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1426)))
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v207)+24))
	v1810 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+24)) = v1809 + v1810
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v207)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v1813+v1809<<(uint(v1810)%32)))) = uint16(v1808)
	goto L240
L298:
	;
	v1727 = *(*float64)(unsafe.Add(mBase, uint32(v1036)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1727)&int64(9223372036854775807)) {
		goto L301
	} else {
		goto L302
	}
L299:
	;
	goto L300
L300:
	;
	v1798 = *(*int64)(unsafe.Add(mBase, uint32(v1431)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1036)+24)) = v1798
	v1800 = *(*int64)(unsafe.Add(mBase, uint32(v1431)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1036)+16)) = v1800
	v1802 = *(*int64)(unsafe.Add(mBase, uint32(v1431)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1036)+8)) = v1802
	v1804 = *(*int64)(unsafe.Add(mBase, uint32(v1431)))
	*(*int64)(unsafe.Add(mBase, uint32(v1036))) = v1804
	goto L297
L301:
	;
	v1745 = *(*float64)(unsafe.Add(mBase, uint32(v1431)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1745)&int64(9223372036854775807)) {
		goto L304
	} else {
		goto L305
	}
L302:
	;
	v1733 = *(*float64)(unsafe.Add(mBase, uint32(v1431)))
	if base.B2i32(base.F64_lt(v1727, v1733) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1733)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L301
	} else {
		goto L303
	}
L303:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1036))) = v1733
	goto L301
L304:
	;
	v1763 = *(*float64)(unsafe.Add(mBase, uint32(v1036)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1763)&int64(9223372036854775807)) {
		goto L307
	} else {
		goto L308
	}
L305:
	;
	v1751 = *(*float64)(unsafe.Add(mBase, uint32(v1036)+16))
	if base.B2i32(base.F64_gt(v1751, v1745) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1751)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L304
	} else {
		goto L306
	}
L306:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1036)+16)) = v1745
	goto L304
L307:
	;
	v1781 = *(*float64)(unsafe.Add(mBase, uint32(v1431)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1781)&int64(9223372036854775807)) {
		goto L297
	} else {
		goto L310
	}
L308:
	;
	v1769 = *(*float64)(unsafe.Add(mBase, uint32(v1431)+8))
	if base.B2i32(base.F64_lt(v1763, v1769) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1769)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L307
	} else {
		goto L309
	}
L309:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1036)+8)) = v1769
	goto L307
L310:
	;
	v1787 = *(*float64)(unsafe.Add(mBase, uint32(v1036)+24))
	if base.B2i32(base.F64_gt(v1787, v1781) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1787)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L297
	} else {
		goto L311
	}
L311:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1036)+24)) = v1781
	goto L297
L312:
	;
	goto L239
}
func F_gist_circle_consistent(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v32 float64
	_ = v32
	var v45 float64
	_ = v45
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v56 float64
	_ = v56
	var v58 float64
	_ = v58
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v79 float64
	_ = v79
	var v92 float64
	_ = v92
	var v93 int32
	_ = v93
	var v94 float64
	_ = v94
	var v95 float64
	_ = v95
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v101 float64
	_ = v101
	var v103 float64
	_ = v103
	var v113 float64
	_ = v113
	var v114 int32
	_ = v114
	var v115 float64
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int64
	_ = v124
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v18 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v18)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if base.B2i32(v20 == v5)|base.B2i32(v15 == v5) == v5 {
		v28 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
		v29 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
		v30 = base.F64_add(v28, v29)
		v32 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v30), v32)|base.F64_eq(base.F64_abs(v28), v32)|base.F64_eq(base.F64_abs(v29), v32) == int32(0) {
			v45 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int64(0)
			} else {
				v49 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
				v50 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
				v51 = v50
				v52 = v45
				v53 = v49
				*(*float64)(unsafe.Add(mBase, uint32(v12))) = v52
				v56 = math.Float64frombits(uint64(0x7ff0000000000000))
				v58 = base.F64_sub(v53, v51)
				if base.F64_eq(base.F64_abs(v53), v56)|base.F64_ne(base.F64_abs(v58), v56)|base.F64_eq(base.F64_abs(v51), v56) == int32(0) {
					v70 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int64(0)
					} else {
						v72 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
						v73 = v72
						v74 = v70
						*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v74
						v76 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
						v77 = base.F64_add(v76, v73)
						v79 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v76), v79)|base.F64_eq(base.F64_abs(v73), v79) == int32(0) {
							v92 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int64(0)
							} else {
								v94 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
								v95 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
								v96 = v95
								v97 = v92
								v98 = v94
								*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v97
								v101 = math.Float64frombits(uint64(0x7ff0000000000000))
								v103 = base.F64_sub(v98, v96)
								if base.F64_eq(base.F64_abs(v98), v101)|base.F64_ne(base.F64_abs(v103), v101)|base.F64_eq(base.F64_abs(v96), v101) != 0 {
									v115 = v103
									*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
									v118 = F_rtree_internal_consistent(m, v117, v12, v14)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int64(0)
									} else {
										v124 = base.I64_extend_i32_u(v118)
										m.G0 = v12 + int32(32)
										return v124
									}
								} else {
									v113 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int64(0)
									} else {
										v115 = v113
										*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
										v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
										v118 = F_rtree_internal_consistent(m, v117, v12, v14)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int64(0)
										} else {
											v124 = base.I64_extend_i32_u(v118)
											m.G0 = v12 + int32(32)
											return v124
										}
									}
								}
							}
						} else {
							v96 = v73
							v97 = v77
							v98 = v76
							*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v97
							v101 = math.Float64frombits(uint64(0x7ff0000000000000))
							v103 = base.F64_sub(v98, v96)
							if base.F64_eq(base.F64_abs(v98), v101)|base.F64_ne(base.F64_abs(v103), v101)|base.F64_eq(base.F64_abs(v96), v101) != 0 {
								v115 = v103
								*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
								v118 = F_rtree_internal_consistent(m, v117, v12, v14)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int64(0)
								} else {
									v124 = base.I64_extend_i32_u(v118)
									m.G0 = v12 + int32(32)
									return v124
								}
							} else {
								v113 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int64(0)
								} else {
									v115 = v113
									*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
									v118 = F_rtree_internal_consistent(m, v117, v12, v14)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int64(0)
									} else {
										v124 = base.I64_extend_i32_u(v118)
										m.G0 = v12 + int32(32)
										return v124
									}
								}
							}
						}
					}
				} else {
					v73 = v51
					v74 = v58
					*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v74
					v76 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
					v77 = base.F64_add(v76, v73)
					v79 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v76), v79)|base.F64_eq(base.F64_abs(v73), v79) == int32(0) {
						v92 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int64(0)
						} else {
							v94 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
							v95 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
							v96 = v95
							v97 = v92
							v98 = v94
							*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v97
							v101 = math.Float64frombits(uint64(0x7ff0000000000000))
							v103 = base.F64_sub(v98, v96)
							if base.F64_eq(base.F64_abs(v98), v101)|base.F64_ne(base.F64_abs(v103), v101)|base.F64_eq(base.F64_abs(v96), v101) != 0 {
								v115 = v103
								*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
								v118 = F_rtree_internal_consistent(m, v117, v12, v14)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int64(0)
								} else {
									v124 = base.I64_extend_i32_u(v118)
									m.G0 = v12 + int32(32)
									return v124
								}
							} else {
								v113 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int64(0)
								} else {
									v115 = v113
									*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
									v118 = F_rtree_internal_consistent(m, v117, v12, v14)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int64(0)
									} else {
										v124 = base.I64_extend_i32_u(v118)
										m.G0 = v12 + int32(32)
										return v124
									}
								}
							}
						}
					} else {
						v96 = v73
						v97 = v77
						v98 = v76
						*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v97
						v101 = math.Float64frombits(uint64(0x7ff0000000000000))
						v103 = base.F64_sub(v98, v96)
						if base.F64_eq(base.F64_abs(v98), v101)|base.F64_ne(base.F64_abs(v103), v101)|base.F64_eq(base.F64_abs(v96), v101) != 0 {
							v115 = v103
							*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
							v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
							v118 = F_rtree_internal_consistent(m, v117, v12, v14)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int64(0)
							} else {
								v124 = base.I64_extend_i32_u(v118)
								m.G0 = v12 + int32(32)
								return v124
							}
						} else {
							v113 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int64(0)
							} else {
								v115 = v113
								*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
								v118 = F_rtree_internal_consistent(m, v117, v12, v14)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int64(0)
								} else {
									v124 = base.I64_extend_i32_u(v118)
									m.G0 = v12 + int32(32)
									return v124
								}
							}
						}
					}
				}
			}
		} else {
			v51 = v29
			v52 = v30
			v53 = v28
			*(*float64)(unsafe.Add(mBase, uint32(v12))) = v52
			v56 = math.Float64frombits(uint64(0x7ff0000000000000))
			v58 = base.F64_sub(v53, v51)
			if base.F64_eq(base.F64_abs(v53), v56)|base.F64_ne(base.F64_abs(v58), v56)|base.F64_eq(base.F64_abs(v51), v56) == int32(0) {
				v70 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int64(0)
				} else {
					v72 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
					v73 = v72
					v74 = v70
					*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v74
					v76 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
					v77 = base.F64_add(v76, v73)
					v79 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v76), v79)|base.F64_eq(base.F64_abs(v73), v79) == int32(0) {
						v92 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int64(0)
						} else {
							v94 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
							v95 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
							v96 = v95
							v97 = v92
							v98 = v94
							*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v97
							v101 = math.Float64frombits(uint64(0x7ff0000000000000))
							v103 = base.F64_sub(v98, v96)
							if base.F64_eq(base.F64_abs(v98), v101)|base.F64_ne(base.F64_abs(v103), v101)|base.F64_eq(base.F64_abs(v96), v101) != 0 {
								v115 = v103
								*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
								v118 = F_rtree_internal_consistent(m, v117, v12, v14)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int64(0)
								} else {
									v124 = base.I64_extend_i32_u(v118)
									m.G0 = v12 + int32(32)
									return v124
								}
							} else {
								v113 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int64(0)
								} else {
									v115 = v113
									*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
									v118 = F_rtree_internal_consistent(m, v117, v12, v14)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int64(0)
									} else {
										v124 = base.I64_extend_i32_u(v118)
										m.G0 = v12 + int32(32)
										return v124
									}
								}
							}
						}
					} else {
						v96 = v73
						v97 = v77
						v98 = v76
						*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v97
						v101 = math.Float64frombits(uint64(0x7ff0000000000000))
						v103 = base.F64_sub(v98, v96)
						if base.F64_eq(base.F64_abs(v98), v101)|base.F64_ne(base.F64_abs(v103), v101)|base.F64_eq(base.F64_abs(v96), v101) != 0 {
							v115 = v103
							*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
							v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
							v118 = F_rtree_internal_consistent(m, v117, v12, v14)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int64(0)
							} else {
								v124 = base.I64_extend_i32_u(v118)
								m.G0 = v12 + int32(32)
								return v124
							}
						} else {
							v113 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int64(0)
							} else {
								v115 = v113
								*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
								v118 = F_rtree_internal_consistent(m, v117, v12, v14)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int64(0)
								} else {
									v124 = base.I64_extend_i32_u(v118)
									m.G0 = v12 + int32(32)
									return v124
								}
							}
						}
					}
				}
			} else {
				v73 = v51
				v74 = v58
				*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v74
				v76 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
				v77 = base.F64_add(v76, v73)
				v79 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v76), v79)|base.F64_eq(base.F64_abs(v73), v79) == int32(0) {
					v92 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int64(0)
					} else {
						v94 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
						v95 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
						v96 = v95
						v97 = v92
						v98 = v94
						*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v97
						v101 = math.Float64frombits(uint64(0x7ff0000000000000))
						v103 = base.F64_sub(v98, v96)
						if base.F64_eq(base.F64_abs(v98), v101)|base.F64_ne(base.F64_abs(v103), v101)|base.F64_eq(base.F64_abs(v96), v101) != 0 {
							v115 = v103
							*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
							v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
							v118 = F_rtree_internal_consistent(m, v117, v12, v14)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int64(0)
							} else {
								v124 = base.I64_extend_i32_u(v118)
								m.G0 = v12 + int32(32)
								return v124
							}
						} else {
							v113 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int64(0)
							} else {
								v115 = v113
								*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
								v118 = F_rtree_internal_consistent(m, v117, v12, v14)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int64(0)
								} else {
									v124 = base.I64_extend_i32_u(v118)
									m.G0 = v12 + int32(32)
									return v124
								}
							}
						}
					}
				} else {
					v96 = v73
					v97 = v77
					v98 = v76
					*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v97
					v101 = math.Float64frombits(uint64(0x7ff0000000000000))
					v103 = base.F64_sub(v98, v96)
					if base.F64_eq(base.F64_abs(v98), v101)|base.F64_ne(base.F64_abs(v103), v101)|base.F64_eq(base.F64_abs(v96), v101) != 0 {
						v115 = v103
						*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
						v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
						v118 = F_rtree_internal_consistent(m, v117, v12, v14)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int64(0)
						} else {
							v124 = base.I64_extend_i32_u(v118)
							m.G0 = v12 + int32(32)
							return v124
						}
					} else {
						v113 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return int64(0)
						} else {
							v115 = v113
							*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v115
							v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
							v118 = F_rtree_internal_consistent(m, v117, v12, v14)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int64(0)
							} else {
								v124 = base.I64_extend_i32_u(v118)
								m.G0 = v12 + int32(32)
								return v124
							}
						}
					}
				}
			}
		}
	} else {
		v124 = int64(0)
		m.G0 = v12 + int32(32)
		return v124
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
func F_gist_point_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+18)))
	if v7 != int32(1) {
		return base.I64_extend_i32_u(v6)
	} else {
		v13 = F_palloc(m, int32(32))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v19 = F_palloc(m, int32(24))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v21 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v21
				v23 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v23
				v25 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				v26 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v26
				*(*int64)(unsafe.Add(mBase, uint32(v13))) = v25
				*(*int64)(unsafe.Add(mBase, uint32(v19))) = base.I64_extend_i32_u(v13)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v31
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v33
				v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+16)))
				v36 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v19)+18)) = uint8(v36)
				*(*uint16)(unsafe.Add(mBase, uint32(v19)+16)) = uint16(v35)
				return base.I64_extend_i32_u(v19)
			}
		}
	}
}
func F_gist_point_fetch(m *base.Module, l0 int32) int64 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v8 = F_palloc(m, int32(24))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v13 = F_palloc(m, int32(16))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
			*(*float64)(unsafe.Add(mBase, uint32(v13))) = v15
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = v17
			*(*int64)(unsafe.Add(mBase, uint32(v8))) = base.I64_extend_i32_u(v13)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v21
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v23
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+16)))
			v26 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)) = uint8(v26)
			*(*uint16)(unsafe.Add(mBase, uint32(v8)+16)) = uint16(v25)
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_gist_poly_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+18)))
	if v5 != int32(1) {
		return base.I64_extend_i32_u(v4)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v16 = F_palloc(m, int32(32))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				v18 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v18
				v20 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v20
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v22
				v24 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v16))) = v24
				v27 = F_palloc(m, int32(24))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v27))) = base.I64_extend_i32_u(v16)
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v31
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v33
					v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+16)))
					v36 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v27)+18)) = uint8(v36)
					*(*uint16)(unsafe.Add(mBase, uint32(v27)+16)) = uint16(v35)
					return base.I64_extend_i32_u(v27)
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
	var v58 int32
	_ = v58
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
	var v84 int32
	_ = v84
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
							v58 = F_errdetail(m, int32(_a_F_verify_gist_page_3), v4+int32(-32))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_verify_gist_page_4), int32(59), int32(_a_F_verify_gist_page_5))
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
								v84 = F_errdetail(m, int32(_a_F_verify_gist_page_7), v6)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_verify_gist_page_4), int32(68), int32(_a_F_verify_gist_page_5))
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
