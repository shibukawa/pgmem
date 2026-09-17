package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GenerationAlloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	v11 = (l1 + int32(7)) & int32(-8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v12) < base.Ui32(v11) {
		v16 = int32(0)
		if (base.B2i32(l2&int32(1) == v16)|base.B2i32(l1 < v16))&base.B2i32(base.Ui32(int32(1073741824)) <= base.Ui32(l1)) == v16 {
			v27 = v11 + int32(40)
			v28 = F_emscripten_builtin_malloc(m, v27)
			mBase = m.M
			if v28 == int32(0) {
				v31 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v64 = v31
					return v64
				}
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v35 + v27
				v38 = v28 + v27
				*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v38
				*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = int64(1)
				*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v27
				*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = l0
				*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = int64(-5645020766237429836)
				*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v38
				v48 = l0 + int32(68)
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				if v49 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v48))) = v48
					v53 = v48
				} else {
					v53 = v49
				}
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = v48
				*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v53
				*(*int32)(unsafe.Add(mBase, uint32(v53))) = v28
				*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v28
				v64 = v28 + int32(40)
				return v64
			}
		} else {
			F_MemoryContextSizeFailure(m, l1)
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		v67 = v11 + int32(8)
		v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
		if base.Ui32(v69-v70) < base.Ui32(v67) {
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			if v73 == int32(0) {
				v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v108 = v106 << (uint(int32(1)) % 32)
				v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if base.Ui32(v108) < base.Ui32(v109) {
					v111 = v108
				} else {
					v111 = v109
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v111
				v114 = v11 + int32(40)
				if base.Ui32(v106) < base.Ui32(v114) {
					v116 = int32(1)
					if v114&(v114-v116) != 0 {
						v124 = v116 << (uint(int32(32)-base.I32_clz(v114)) % 32)
					} else {
						v124 = v114
					}
					v125 = v124
				} else {
					v125 = v106
				}
				v126 = F_emscripten_builtin_malloc(m, v125)
				mBase = m.M
				if v126 == int32(0) {
					v129 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return int32(0)
					} else {
						v179 = v129
						return v179
					}
				} else {
					v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v131 + v125
					*(*int32)(unsafe.Add(mBase, uint32(v126)+28)) = v125 + v126
					*(*int32)(unsafe.Add(mBase, uint32(v126)+24)) = v126 + int32(32)
					*(*int64)(unsafe.Add(mBase, uint32(v126)+16)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v126)+12)) = v125
					*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = l0
					v144 = l0 + int32(68)
					v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					if v145 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v144))) = v144
						v149 = v144
					} else {
						v149 = v145
					}
					*(*int32)(unsafe.Add(mBase, uint32(v126))) = v144
					*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v149
					*(*int32)(unsafe.Add(mBase, uint32(v149))) = v126
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v126
					*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v126
					v155 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v126)+16)) = v155 + int32(1)
					v159 = *(*int32)(unsafe.Add(mBase, uint32(v126)+24))
					v161 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v126)+24)) = v159 + v11 + v161
					*(*int64)(unsafe.Add(mBase, uint32(v159))) = base.I64_extend_i32_u(v11)<<(uint(int64(5))%64) | base.I64_extend_i32_u(v159-v126)<<(uint(int64(34))%64) | int64(4)
					v179 = v159 + v161
					return v179
				}
			} else {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)+28))
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)+24))
				if base.Ui32(v76-v77) < base.Ui32(v67) {
					v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v108 = v106 << (uint(int32(1)) % 32)
					v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if base.Ui32(v108) < base.Ui32(v109) {
						v111 = v108
					} else {
						v111 = v109
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v111
					v114 = v11 + int32(40)
					if base.Ui32(v106) < base.Ui32(v114) {
						v116 = int32(1)
						if v114&(v114-v116) != 0 {
							v124 = v116 << (uint(int32(32)-base.I32_clz(v114)) % 32)
						} else {
							v124 = v114
						}
						v125 = v124
					} else {
						v125 = v106
					}
					v126 = F_emscripten_builtin_malloc(m, v125)
					mBase = m.M
					if v126 == int32(0) {
						v129 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int32(0)
						} else {
							v179 = v129
							return v179
						}
					} else {
						v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v131 + v125
						*(*int32)(unsafe.Add(mBase, uint32(v126)+28)) = v125 + v126
						*(*int32)(unsafe.Add(mBase, uint32(v126)+24)) = v126 + int32(32)
						*(*int64)(unsafe.Add(mBase, uint32(v126)+16)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v126)+12)) = v125
						*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = l0
						v144 = l0 + int32(68)
						v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
						if v145 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v144))) = v144
							v149 = v144
						} else {
							v149 = v145
						}
						*(*int32)(unsafe.Add(mBase, uint32(v126))) = v144
						*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v149
						*(*int32)(unsafe.Add(mBase, uint32(v149))) = v126
						*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v126
						*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v126
						v155 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v126)+16)) = v155 + int32(1)
						v159 = *(*int32)(unsafe.Add(mBase, uint32(v126)+24))
						v161 = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v126)+24)) = v159 + v11 + v161
						*(*int64)(unsafe.Add(mBase, uint32(v159))) = base.I64_extend_i32_u(v11)<<(uint(int64(5))%64) | base.I64_extend_i32_u(v159-v126)<<(uint(int64(34))%64) | int64(4)
						v179 = v159 + v161
						return v179
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v73
					*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = v83 + int32(1)
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v73)+24))
					v89 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = v87 + v11 + v89
					*(*int64)(unsafe.Add(mBase, uint32(v87))) = base.I64_extend_i32_u(v11)<<(uint(int64(5))%64) | base.I64_extend_i32_u(v87-v73)<<(uint(int64(34))%64) | int64(4)
					return v87 + v89
				}
			}
		} else {
			v182 = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = v70 + v11 + v182
			v185 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v68)+16)) = v185 + int32(1)
			*(*int64)(unsafe.Add(mBase, uint32(v70))) = base.I64_extend_i32_u(v11)<<(uint(int64(5))%64) | base.I64_extend_i32_u(v70-v68)<<(uint(int64(34))%64) | int64(4)
			return v70 + v182
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
	var v35 int32
	_ = v35
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
		*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v17 + v18
		*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v18 + int32(112)
		v31 = int32(80)
		v32 = v17 - v31
		*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v32
		v35 = v18 + int32(68)
		*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v35
		*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v35
		v39 = v18 + v31
		*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v39
		*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v39
		*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
		*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v18
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
		*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(_a_F_GenerationContextCreate_0)
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
		v101 = *(*int32)(unsafe.Add(mBase, _c_F_GenerationContextCreate[0]))
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
				F_errcode(m, int32(_a_F_GenerationContextCreate_1))
				mBase = m.M
				v112 = m.ExcPending
				if v112 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_GenerationContextCreate_2), int32(0))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
						F_errdetail(m, int32(_a_F_GenerationContextCreate_3), v11)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_GenerationContextCreate_4), int32(217), int32(_a_F_GenerationContextCreate_5))
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
	var v30 int32
	_ = v30
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(240)
	m.G0 = v16
	v18 = int32(80)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v19 == v6 {
		v59 = v18
		v60 = v6
		v61 = v6
		v62 = v6
		v63 = v6
	} else {
		v23 = l0 + int32(68)
		if v19 == v23 {
			v59 = v18
			v60 = v6
			v61 = v6
			v62 = v6
			v63 = v6
		} else {
			v30 = v19
			v32 = v18
			v33 = v6
			v34 = v6
			v35 = v6
			v36 = v6
			for {
				v39 = v34 + int32(1)
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				v41 = v40 + v32
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
				v43 = v42 + v35
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
				v45 = v44 + v36
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
				v49 = v33 + v46 - v48
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
				if v50 != v23 {
					v30 = v50
					v32 = v41
					v33 = v49
					v34 = v39
					v35 = v43
					v36 = v45
					continue
				} else {
					break
				}
				break
			}
			v59 = v41
			v60 = v49
			v61 = v39
			v62 = v43
			v63 = v45
		}
	}
	if l1 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v62
		*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v60
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = v59
		*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v59 - v60
		*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v61
		*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v63
		v73 = v16 + int32(32)
		v76 = F_pg_snprintf(m, v73, int32(200), int32(_a_F_GenerationStats_0), v16)
		mBase = m.M
		v77 = m.ExcPending
		if v77 != 0 {
			return
		} else {
			m.T0[l1].(func(*base.Module, int32, int32, int32, int32))(m, l0, l2, v73, l4)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return
			} else {
				if l3 != 0 {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v81 + v61
					v84 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v84 + v62
					v87 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v87 + v59
					v90 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v90 + v60
				} else {
				}
				m.G0 = v16 + int32(240)
				return
			}
		}
	} else {
		if l3 != 0 {
			v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v81 + v61
			v84 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v84 + v62
			v87 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v87 + v59
			v90 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v90 + v60
		} else {
		}
		m.G0 = v16 + int32(240)
		return
	}
}
