package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GenerationAlloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	v12 = (l1 + int32(7)) & int32(-8)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v13) < base.Ui32(v12) {
		if base.Ui32(int32(1073741824)) <= base.Ui32(l1) {
			if l1 < int32(0) {
				F_MemoryContextSizeFailure(m, l1)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if l2&int32(1) == int32(0) {
					F_MemoryContextSizeFailure(m, l1)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v28 = (l1+int32(7))&int32(-8) + int32(40)
					v29 = F_emscripten_builtin_malloc(m, v28)
					mBase = m.M
					if v29 == int32(0) {
						v32 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v65 = v32
							return v65
						}
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v36 + v28
						v39 = v29 + v28
						*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v39
						*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = int64(1)
						*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v28
						*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = l0
						*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = int64(-5645020766237429836)
						*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
						v49 = l0 + int32(68)
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
						if v50 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v49))) = v49
							v54 = v49
						} else {
							v54 = v50
						}
						*(*int32)(unsafe.Add(mBase, uint32(v29))) = v49
						*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v54
						*(*int32)(unsafe.Add(mBase, uint32(v54))) = v29
						*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v29
						v65 = v29 + int32(40)
						return v65
					}
				}
			}
		} else {
			v28 = (l1+int32(7))&int32(-8) + int32(40)
			v29 = F_emscripten_builtin_malloc(m, v28)
			mBase = m.M
			if v29 == int32(0) {
				v32 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v65 = v32
					return v65
				}
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v36 + v28
				v39 = v29 + v28
				*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v39
				*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = int64(1)
				*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v28
				*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = l0
				*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = int64(-5645020766237429836)
				*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v39
				v49 = l0 + int32(68)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				if v50 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v49))) = v49
					v54 = v49
				} else {
					v54 = v50
				}
				*(*int32)(unsafe.Add(mBase, uint32(v29))) = v49
				*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v54
				*(*int32)(unsafe.Add(mBase, uint32(v54))) = v29
				*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v29
				v65 = v29 + int32(40)
				return v65
			}
		}
	} else {
		v68 = v12 + int32(8)
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+28))
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+24))
		if base.Ui32(v70-v71) < base.Ui32(v68) {
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			if v74 == int32(0) {
				v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v109 = v107 << (uint(int32(1)) % 32)
				v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if base.Ui32(v109) < base.Ui32(v110) {
					v112 = v109
				} else {
					v112 = v110
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v112
				v115 = v12 + int32(40)
				if base.Ui32(v107) < base.Ui32(v115) {
					v117 = int32(1)
					if v115&(v115-v117) != 0 {
						v125 = v117 << (uint(int32(32)-base.I32_clz(v115)) % 32)
					} else {
						v125 = v115
					}
					v126 = v125
				} else {
					v126 = v107
				}
				v127 = F_emscripten_builtin_malloc(m, v126)
				mBase = m.M
				if v127 == int32(0) {
					v130 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return int32(0)
					} else {
						v181 = v130
						return v181
					}
				} else {
					v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v132 + v126
					*(*int32)(unsafe.Add(mBase, uint32(v127)+28)) = v127 + v126
					*(*int32)(unsafe.Add(mBase, uint32(v127)+24)) = v127 + int32(32)
					*(*int64)(unsafe.Add(mBase, uint32(v127)+16)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v127)+12)) = v126
					*(*int32)(unsafe.Add(mBase, uint32(v127)+8)) = l0
					v145 = l0 + int32(68)
					v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					if v146 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v145))) = v145
						v150 = v145
					} else {
						v150 = v146
					}
					*(*int32)(unsafe.Add(mBase, uint32(v127))) = v145
					*(*int32)(unsafe.Add(mBase, uint32(v127)+4)) = v150
					*(*int32)(unsafe.Add(mBase, uint32(v150))) = v127
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v127
					*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v127
					v156 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v127)+16)) = v156 + int32(1)
					v160 = *(*int32)(unsafe.Add(mBase, uint32(v127)+24))
					v162 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v127)+24)) = v160 + v12 + v162
					*(*int64)(unsafe.Add(mBase, uint32(v160))) = base.I64_extend_i32_u(v12)<<(uint(int64(5))%64) | base.I64_extend_i32_u(v160-v127)<<(uint(int64(34))%64) | int64(4)
					v181 = v160 + v162
					return v181
				}
			} else {
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
				if base.Ui32(v77-v78) < base.Ui32(v68) {
					v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v109 = v107 << (uint(int32(1)) % 32)
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if base.Ui32(v109) < base.Ui32(v110) {
						v112 = v109
					} else {
						v112 = v110
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v112
					v115 = v12 + int32(40)
					if base.Ui32(v107) < base.Ui32(v115) {
						v117 = int32(1)
						if v115&(v115-v117) != 0 {
							v125 = v117 << (uint(int32(32)-base.I32_clz(v115)) % 32)
						} else {
							v125 = v115
						}
						v126 = v125
					} else {
						v126 = v107
					}
					v127 = F_emscripten_builtin_malloc(m, v126)
					mBase = m.M
					if v127 == int32(0) {
						v130 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return int32(0)
						} else {
							v181 = v130
							return v181
						}
					} else {
						v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v132 + v126
						*(*int32)(unsafe.Add(mBase, uint32(v127)+28)) = v127 + v126
						*(*int32)(unsafe.Add(mBase, uint32(v127)+24)) = v127 + int32(32)
						*(*int64)(unsafe.Add(mBase, uint32(v127)+16)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v127)+12)) = v126
						*(*int32)(unsafe.Add(mBase, uint32(v127)+8)) = l0
						v145 = l0 + int32(68)
						v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
						if v146 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v145))) = v145
							v150 = v145
						} else {
							v150 = v146
						}
						*(*int32)(unsafe.Add(mBase, uint32(v127))) = v145
						*(*int32)(unsafe.Add(mBase, uint32(v127)+4)) = v150
						*(*int32)(unsafe.Add(mBase, uint32(v150))) = v127
						*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v127
						*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v127
						v156 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v127)+16)) = v156 + int32(1)
						v160 = *(*int32)(unsafe.Add(mBase, uint32(v127)+24))
						v162 = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v127)+24)) = v160 + v12 + v162
						*(*int64)(unsafe.Add(mBase, uint32(v160))) = base.I64_extend_i32_u(v12)<<(uint(int64(5))%64) | base.I64_extend_i32_u(v160-v127)<<(uint(int64(34))%64) | int64(4)
						v181 = v160 + v162
						return v181
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v74
					*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v74)+16)) = v84 + int32(1)
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
					v90 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v74)+24)) = v88 + v12 + v90
					*(*int64)(unsafe.Add(mBase, uint32(v88))) = base.I64_extend_i32_u(v12)<<(uint(int64(5))%64) | base.I64_extend_i32_u(v88-v74)<<(uint(int64(34))%64) | int64(4)
					return v88 + v90
				}
			}
		} else {
			v184 = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = v12 + v71 + v184
			v187 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v69)+16)) = v187 + int32(1)
			*(*int64)(unsafe.Add(mBase, uint32(v71))) = base.I64_extend_i32_u(v12)<<(uint(int64(5))%64) | base.I64_extend_i32_u(v71-v69)<<(uint(int64(34))%64) | int64(4)
			return v71 + v184
		}
	}
}
func F_GenerationContextCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	var v72 int64
	_ = v72
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l2 != 0 {
		v14 = l2
	} else {
		v14 = l3
	}
	if base.Ui32(v14) <= base.Ui32(int32(120)) {
		v17 = int32(120)
	} else {
		v17 = v14
	}
	v18 = F_emscripten_builtin_malloc(m, v17)
	mBase = m.M
	if v18 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v17 + v18
		*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v18 + int32(112)
		*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = int64(0)
		v31 = int32(80)
		v32 = v17 - v31
		*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v32
		*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v18
		v36 = v18 + int32(68)
		*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v36
		v39 = v18 + v31
		*(*int32)(unsafe.Add(mBase, uint32(v39))) = v36
		*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v39
		*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v39
		*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
		v44 = int32(1073741823)
		if base.Ui32(v44) <= base.Ui32(l4) {
			v47 = v44
		} else {
			v47 = l4
		}
		v55 = v47
		for {
			if base.Ui32(int32(base.Ui32(l4-int32(32))>>(uint(int32(3))%32))) < base.Ui32(v55+int32(8)) {
				v55 = int32(base.Ui32(v55) >> (uint(int32(1)) % 32))
				continue
			} else {
				break
			}
			break
		}
		*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v55
		*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l0
		v69 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)) = uint8(v69)
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(475)
		v72 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v18)+36)) = v72
		*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v18)+20)) = v72
		*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(1752432)
		if l0 != 0 {
			v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v84
			if v84 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v84)+24)) = v18
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v18
			v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+5)) = uint8(v88)
		} else {
			v90 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v90
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+5)) = uint8(v90)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v32
		m.G0 = v11 + int32(16)
		return v18
	} else {
		v101 = *(*int32)(unsafe.Add(mBase, _consts[12]))
		F_MemoryContextStats(m, v101)
		mBase = m.M
		v105 = m.ExcPending
		if v105 != 0 {
			return int32(0)
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v109 = m.ExcPending
			if v109 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(8389))
				mBase = m.M
				v112 = m.ExcPending
				if v112 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(13869), int32(0))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
						F_errdetail(m, int32(646394), v11)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(493130), int32(217), int32(353341))
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
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
		}
	}
}
func F_GenerationIsEmpty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
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
	var v25 int32
	_ = v25
	v5 = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v6 == int32(0) {
		v25 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v25
L2:
	;
	v10 = l0 + int32(68)
	if v6 == v10 {
		v25 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = v6
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v17 = int32(0)
	v18 = base.B2i32(v16 <= v17)
	if v17 < v16 {
		v25 = v18
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v25 = v18
	goto L1
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v21 != v10 {
		v13 = v21
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
}
func F_GenerationStats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(240)
	m.G0 = v16
	v18 = int32(80)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v19 == v6 {
		v60 = v6
		v61 = v18
		v62 = v6
		v63 = v6
		v64 = v6
	} else {
		v23 = l0 + int32(68)
		if v19 == v23 {
			v60 = v6
			v61 = v18
			v62 = v6
			v63 = v6
			v64 = v6
		} else {
			v31 = v19
			v32 = v6
			v33 = v18
			v34 = v6
			v35 = v6
			v36 = v6
			for {
				v39 = v35 + int32(1)
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
				v41 = v40 + v33
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
				v43 = v42 + v34
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
				v45 = v44 + v36
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
				v49 = v32 + v46 - v48
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
				if v50 != v23 {
					v31 = v50
					v32 = v49
					v33 = v41
					v34 = v43
					v35 = v39
					v36 = v45
					continue
				} else {
					break
				}
				break
			}
			v60 = v49
			v61 = v41
			v62 = v43
			v63 = v39
			v64 = v45
		}
	}
	if l1 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v62
		*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v60
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = v61
		*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v61 - v60
		*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v63
		*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v64
		v77 = F_pg_snprintf(m, v16+int32(32), int32(200), int32(446587), v16)
		mBase = m.M
		v78 = m.ExcPending
		if v78 != 0 {
			return
		} else {
			m.T0[l1].(func(*base.Module, int32, int32, int32, int32))(m, l0, l2, v16+int32(32), l4)
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return
			} else {
				if l3 != 0 {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83 + v63
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v86 + v62
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v89 + v61
					v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v92 + v60
				} else {
				}
				m.G0 = v16 + int32(240)
				return
			}
		}
	} else {
		if l3 != 0 {
			v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v83 + v63
			v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v86 + v62
			v89 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v89 + v61
			v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v92 + v60
		} else {
		}
		m.G0 = v16 + int32(240)
		return
	}
}
