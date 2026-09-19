package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_AdvanceXLInsertBuffer(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v92 int32
	_ = v92
	var v96 int64
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v120 int64
	_ = v120
	var v124 int32
	_ = v124
	var v128 int64
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int64
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v175 int32
	_ = v175
	var v176 int64
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int64
	_ = v204
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int64
	_ = v227
	var v229 int32
	_ = v229
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	v23 = F_LWLockAcquire(m, v19+int32(896), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+288))
	if base.B2i32(l2 == int32(0))&base.B2i32(base.Ui64(l0) < base.Ui64(v29)) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	F_LWLockRelease(m, v245+int32(896))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L37
	}
L4:
	;
	v35 = v29
	v37 = v28
	goto L5
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v37)+304))
	v49 = base.I64_rem_u_s(int64(base.Ui64(v35)>>(uint(int64(13))%64)), base.I64_extend_i32_s(v45+int32(1)))
	v50 = base.I32_wrap_i64(v49)
	v52 = v50 << (uint(int32(3)) % 32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)+300))
	v55 = int64(0)
	v58 = base.AtomicRmwCmpxchg64(m, v52+v53, int32(0), v55, v55)
	v60 = *(*int64)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[2]))
	if base.Ui64(v58) <= base.Ui64(v60) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v175)+288))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v175)+296))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175)+300))
	v181 = int32(0)
	v182 = base.AtomicRmwXchg64(m, v178+v52, v181, int64(0))
	v183 = int32(2)
	v186 = v177 + v50<<(uint(int32(13))%32)
	base.MemoryFill(m, v186+v183, v181, int32(_a_F_AdvanceXLInsertBuffer_0))
	*(*int64)(unsafe.Add(mBase, uint32(v186)+8)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v186)+4)) = l1
	v194 = int32(_a_F_AdvanceXLInsertBuffer_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v186))) = uint16(v194)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v17)+164))
	if v196 == v181 {
		goto L30
	} else {
		goto L31
	}
L8:
	;
	if l2 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	v66 = base.AtomicRmwXchg32(m, v63, int32(440), int32(1))
	if v66 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	F_s_lock(m, v68+int32(440), int32(_a_F_AdvanceXLInsertBuffer_2), int32(2025), int32(_a_F_AdvanceXLInsertBuffer_3))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+184))
	if base.Ui64(v78) < base.Ui64(v58) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77)+184)) = v58
	goto L16
L15:
	;
	goto L16
L16:
	;
	v81 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v77)+440)), uint32(v81))
	v85 = int64(0)
	v88 = base.AtomicRmwCmpxchg64(m, v77, int32(280), v85, v85)
	*(*int64)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[3])) = v88
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	v96 = base.AtomicRmwCmpxchg64(m, v92, int32(272), v85, v85)
	*(*int64)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[2])) = v96
	if base.Ui64(v58) <= base.Ui64(v96) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	F_LWLockRelease(m, v100+int32(896))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v105 = F_WaitXLogInsertionsToFinish(m, v58)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	v112 = F_LWLockAcquire(m, v108+int32(1024), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v115 = int32(_a_F_AdvanceXLInsertBuffer_4)
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	v117 = int64(0)
	v120 = base.AtomicRmwCmpxchg64(m, v116, int32(280), v117, v117)
	*(*int64)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[3])) = v120
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	v128 = base.AtomicRmwCmpxchg64(m, v124, int32(272), v117, v117)
	*(*int64)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[2])) = v128
	if base.Ui64(v58) <= base.Ui64(v128) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	v166 = F_LWLockAcquire(m, v162+int32(896), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L28
	}
L22:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	F_LWLockRelease(m, v132+int32(1024))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v137 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v58
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v58
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v137
	F_XLogWrite(m, v14, l1, int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L21
L26:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	F_LWLockRelease(m, v147+int32(1024))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v153 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[4])) = uint8(v153)
	v155 = int32(_a_F_AdvanceXLInsertBuffer_5)
	v157 = *(*int64)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[5]))
	*(*int64)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[5])) = v157 + int64(1)
	goto L21
L28:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v169)+288))
	if base.Ui64(v170) <= base.Ui64(l0) {
		v35 = v170
		v37 = v169
		goto L5
	} else {
		goto L29
	}
L29:
	;
	goto L3
L30:
	;
	v199 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v186)+2)) = uint16(v199)
	v202 = int32(6)
	goto L32
L31:
	;
	v202 = v183
	goto L32
L32:
	;
	v204 = v176 - int64(-8192)
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[6]))
	if v176&base.I64_extend_i32_s(v206-int32(1)) == int64(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[7]))
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v214)))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+36)) = int32(_a_F_AdvanceXLInsertBuffer_6)
	*(*int32)(unsafe.Add(mBase, uint32(v186)+32)) = v206
	*(*int64)(unsafe.Add(mBase, uint32(v186)+24)) = v215
	*(*uint16)(unsafe.Add(mBase, uint32(v186)+2)) = uint16(v202)
	goto L35
L34:
	;
	goto L35
L35:
	;
	v222 = int32(_a_F_AdvanceXLInsertBuffer_4)
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+300))
	v227 = base.AtomicRmwXchg64(m, v224+v52, int32(0), v204)
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v229)+288)) = v204
	if l2|base.B2i32(base.Ui64(v204) <= base.Ui64(l0)) != 0 {
		v35 = v204
		v37 = v229
		goto L5
	} else {
		goto L36
	}
L36:
	;
	goto L6
L37:
	;
	m.G0 = v14 + int32(32)
	return
}
func F_AllocSetAllocLarge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	v4 = int32(0)
	if (base.B2i32(l2&int32(1) == v4)|base.B2i32(l1 < v4))&base.B2i32(base.Ui32(int32(1073741824)) <= base.Ui32(l1)) == v4 {
		v23 = (l1+int32(7))&int32(-8) + int32(32)
		v24 = F_emscripten_builtin_malloc(m, v23)
		mBase = m.M
		if v24 == int32(0) {
			v27 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				return v27
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32 + v23
			v35 = v24 + v23
			*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
			*(*int32)(unsafe.Add(mBase, uint32(v24))) = l0
			*(*int64)(unsafe.Add(mBase, uint32(v24)+24)) = int64(-5645020766237429837)
			*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v35
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v41 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v41
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v43
				if v43 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v24
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v47 = v46
				} else {
					v47 = v41
				}
				*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v24
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v24)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v24
			}
			return v24 + int32(32)
		}
	} else {
		F_MemoryContextSizeFailure(m, l1)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_AllocSetContextCreateInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int64
	_ = v45
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v144 int32
	_ = v144
	var v159 int32
	_ = v159
	var v162 int64
	_ = v162
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l2|base.B2i32(l3 != int32(_a_F_AllocSetContextCreateInternal_0)) != 0 {
		if l2|base.B2i32(l3 != int32(1024)) != 0 {
			v72 = int32(-1)
			if l2 != 0 {
				v74 = l2
			} else {
				v74 = l3
			}
			if base.Ui32(v74) <= base.Ui32(int32(144)) {
				v77 = int32(144)
			} else {
				v77 = v74
			}
			v78 = F_emscripten_builtin_malloc(m, v77)
			mBase = m.M
			if v78 == int32(0) {
				v82 = *(*int32)(unsafe.Add(mBase, _c_F_AllocSetContextCreateInternal[0]))
				if v82 != 0 {
					F_MemoryContextStats(m, v82)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(_a_F_AllocSetContextCreateInternal_1))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_AllocSetContextCreateInternal_2), int32(0))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
									F_errdetail(m, int32(_a_F_AllocSetContextCreateInternal_3), v12)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_AllocSetContextCreateInternal_4), int32(453), int32(_a_F_AllocSetContextCreateInternal_5))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
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
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(_a_F_AllocSetContextCreateInternal_1))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_AllocSetContextCreateInternal_2), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
								F_errdetail(m, int32(_a_F_AllocSetContextCreateInternal_3), v12)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_AllocSetContextCreateInternal_4), int32(453), int32(_a_F_AllocSetContextCreateInternal_5))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
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
			} else {
				v107 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v78)+116)) = v107
				*(*int32)(unsafe.Add(mBase, uint32(v78)+108)) = v72
				*(*int32)(unsafe.Add(mBase, uint32(v78)+100)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v78)+96)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v78)+92)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v78)+128)) = v78 + v77
				*(*int32)(unsafe.Add(mBase, uint32(v78)+124)) = v78 + int32(136)
				*(*int64)(unsafe.Add(mBase, uint32(v78)+48)) = v107
				*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v78 + int32(112)
				*(*int64)(unsafe.Add(mBase, uint32(v78)+56)) = v107
				*(*int64)(unsafe.Add(mBase, uint32(v78-int32(-64)))) = v107
				*(*int64)(unsafe.Add(mBase, uint32(v78)+72)) = v107
				*(*int64)(unsafe.Add(mBase, uint32(v78)+80)) = v107
				*(*int32)(unsafe.Add(mBase, uint32(v78)+88)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v78)+112)) = v78
				v144 = int32(_a_F_AllocSetContextCreateInternal_0)
				for {
					if base.Ui32(int32(base.Ui32(l4-int32(24))>>(uint(int32(2))%32))) < base.Ui32(v144+int32(8)) {
						v144 = int32(base.Ui32(v144) >> (uint(int32(1)) % 32))
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(v78)+104)) = v144
				*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = l0
				v159 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v78)+4)) = uint8(v159)
				*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(474)
				v162 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v162
				*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = l1
				*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v162
				*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = int32(_a_F_AllocSetContextCreateInternal_6)
				if l0 != 0 {
					v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v174
					if v174 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v174)+24)) = v78
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v78
					v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
					*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)) = uint8(v178)
				} else {
					v180 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v180
					*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)) = uint8(v180)
				}
				v190 = v78
				v191 = v77
				*(*int32)(unsafe.Add(mBase, uint32(v190)+8)) = v191
				m.G0 = v12 + int32(16)
				return v190
			}
		} else {
			v24 = int32(1)
			v26 = v24 << (uint(int32(3)) % 32)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_AllocSetContextCreateInternal[1])))
			if v27 == int32(0) {
				v72 = v24
				if l2 != 0 {
					v74 = l2
				} else {
					v74 = l3
				}
				if base.Ui32(v74) <= base.Ui32(int32(144)) {
					v77 = int32(144)
				} else {
					v77 = v74
				}
				v78 = F_emscripten_builtin_malloc(m, v77)
				mBase = m.M
				if v78 == int32(0) {
					v82 = *(*int32)(unsafe.Add(mBase, _c_F_AllocSetContextCreateInternal[0]))
					if v82 != 0 {
						F_MemoryContextStats(m, v82)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(_a_F_AllocSetContextCreateInternal_1))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_AllocSetContextCreateInternal_2), int32(0))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
										F_errdetail(m, int32(_a_F_AllocSetContextCreateInternal_3), v12)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_AllocSetContextCreateInternal_4), int32(453), int32(_a_F_AllocSetContextCreateInternal_5))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
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
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(_a_F_AllocSetContextCreateInternal_1))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_AllocSetContextCreateInternal_2), int32(0))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
									F_errdetail(m, int32(_a_F_AllocSetContextCreateInternal_3), v12)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_AllocSetContextCreateInternal_4), int32(453), int32(_a_F_AllocSetContextCreateInternal_5))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
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
				} else {
					v107 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v78)+116)) = v107
					*(*int32)(unsafe.Add(mBase, uint32(v78)+108)) = v72
					*(*int32)(unsafe.Add(mBase, uint32(v78)+100)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v78)+96)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v78)+92)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v78)+128)) = v78 + v77
					*(*int32)(unsafe.Add(mBase, uint32(v78)+124)) = v78 + int32(136)
					*(*int64)(unsafe.Add(mBase, uint32(v78)+48)) = v107
					*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v78 + int32(112)
					*(*int64)(unsafe.Add(mBase, uint32(v78)+56)) = v107
					*(*int64)(unsafe.Add(mBase, uint32(v78-int32(-64)))) = v107
					*(*int64)(unsafe.Add(mBase, uint32(v78)+72)) = v107
					*(*int64)(unsafe.Add(mBase, uint32(v78)+80)) = v107
					*(*int32)(unsafe.Add(mBase, uint32(v78)+88)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v78)+112)) = v78
					v144 = int32(_a_F_AllocSetContextCreateInternal_0)
					for {
						if base.Ui32(int32(base.Ui32(l4-int32(24))>>(uint(int32(2))%32))) < base.Ui32(v144+int32(8)) {
							v144 = int32(base.Ui32(v144) >> (uint(int32(1)) % 32))
							continue
						} else {
							break
						}
						break
					}
					*(*int32)(unsafe.Add(mBase, uint32(v78)+104)) = v144
					*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = l0
					v159 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v78)+4)) = uint8(v159)
					*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(474)
					v162 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v162
					*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = l1
					*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v162
					*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = int32(_a_F_AllocSetContextCreateInternal_6)
					if l0 != 0 {
						v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v174
						if v174 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v174)+24)) = v78
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v78
						v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
						*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)) = uint8(v178)
					} else {
						v180 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v180
						*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)) = uint8(v180)
					}
					v190 = v78
					v191 = v77
					*(*int32)(unsafe.Add(mBase, uint32(v190)+8)) = v191
					m.G0 = v12 + int32(16)
					return v190
				}
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_AllocSetContextCreateInternal[1]))) = v32
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_AllocSetContextCreateInternal[2])))
				v35 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_AllocSetContextCreateInternal[2]))) = v34 - v35
				*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = l0
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+4)) = uint8(v35)
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(474)
				v45 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v27)+36)) = v45
				*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = l1
				*(*int64)(unsafe.Add(mBase, uint32(v27)+20)) = v45
				*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = int32(_a_F_AllocSetContextCreateInternal_6)
				if l0 != 0 {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v57
					if v57 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v27
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v27
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
					*(*uint8)(unsafe.Add(mBase, uint32(v27)+5)) = uint8(v61)
				} else {
					v63 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v63
					*(*uint8)(unsafe.Add(mBase, uint32(v27)+5)) = uint8(v63)
				}
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v27)+128))
				v190 = v27
				v191 = v68 - v27
				*(*int32)(unsafe.Add(mBase, uint32(v190)+8)) = v191
				m.G0 = v12 + int32(16)
				return v190
			}
		}
	} else {
		v24 = int32(0)
		v26 = v24 << (uint(int32(3)) % 32)
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_AllocSetContextCreateInternal[1])))
		if v27 == int32(0) {
			v72 = v24
			if l2 != 0 {
				v74 = l2
			} else {
				v74 = l3
			}
			if base.Ui32(v74) <= base.Ui32(int32(144)) {
				v77 = int32(144)
			} else {
				v77 = v74
			}
			v78 = F_emscripten_builtin_malloc(m, v77)
			mBase = m.M
			if v78 == int32(0) {
				v82 = *(*int32)(unsafe.Add(mBase, _c_F_AllocSetContextCreateInternal[0]))
				if v82 != 0 {
					F_MemoryContextStats(m, v82)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(_a_F_AllocSetContextCreateInternal_1))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_AllocSetContextCreateInternal_2), int32(0))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
									F_errdetail(m, int32(_a_F_AllocSetContextCreateInternal_3), v12)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_AllocSetContextCreateInternal_4), int32(453), int32(_a_F_AllocSetContextCreateInternal_5))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
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
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(_a_F_AllocSetContextCreateInternal_1))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_AllocSetContextCreateInternal_2), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
								F_errdetail(m, int32(_a_F_AllocSetContextCreateInternal_3), v12)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_AllocSetContextCreateInternal_4), int32(453), int32(_a_F_AllocSetContextCreateInternal_5))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
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
			} else {
				v107 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v78)+116)) = v107
				*(*int32)(unsafe.Add(mBase, uint32(v78)+108)) = v72
				*(*int32)(unsafe.Add(mBase, uint32(v78)+100)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v78)+96)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v78)+92)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v78)+128)) = v78 + v77
				*(*int32)(unsafe.Add(mBase, uint32(v78)+124)) = v78 + int32(136)
				*(*int64)(unsafe.Add(mBase, uint32(v78)+48)) = v107
				*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v78 + int32(112)
				*(*int64)(unsafe.Add(mBase, uint32(v78)+56)) = v107
				*(*int64)(unsafe.Add(mBase, uint32(v78-int32(-64)))) = v107
				*(*int64)(unsafe.Add(mBase, uint32(v78)+72)) = v107
				*(*int64)(unsafe.Add(mBase, uint32(v78)+80)) = v107
				*(*int32)(unsafe.Add(mBase, uint32(v78)+88)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v78)+112)) = v78
				v144 = int32(_a_F_AllocSetContextCreateInternal_0)
				for {
					if base.Ui32(int32(base.Ui32(l4-int32(24))>>(uint(int32(2))%32))) < base.Ui32(v144+int32(8)) {
						v144 = int32(base.Ui32(v144) >> (uint(int32(1)) % 32))
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(v78)+104)) = v144
				*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = l0
				v159 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v78)+4)) = uint8(v159)
				*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(474)
				v162 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v78)+36)) = v162
				*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = l1
				*(*int64)(unsafe.Add(mBase, uint32(v78)+20)) = v162
				*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = int32(_a_F_AllocSetContextCreateInternal_6)
				if l0 != 0 {
					v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v174
					if v174 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v174)+24)) = v78
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v78
					v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
					*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)) = uint8(v178)
				} else {
					v180 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v180
					*(*uint8)(unsafe.Add(mBase, uint32(v78)+5)) = uint8(v180)
				}
				v190 = v78
				v191 = v77
				*(*int32)(unsafe.Add(mBase, uint32(v190)+8)) = v191
				m.G0 = v12 + int32(16)
				return v190
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
			*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_AllocSetContextCreateInternal[1]))) = v32
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_AllocSetContextCreateInternal[2])))
			v35 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_AllocSetContextCreateInternal[2]))) = v34 - v35
			*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = l0
			*(*uint8)(unsafe.Add(mBase, uint32(v27)+4)) = uint8(v35)
			*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(474)
			v45 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v27)+36)) = v45
			*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = l1
			*(*int64)(unsafe.Add(mBase, uint32(v27)+20)) = v45
			*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = int32(_a_F_AllocSetContextCreateInternal_6)
			if l0 != 0 {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v57
				if v57 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v27
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v27
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+5)) = uint8(v61)
			} else {
				v63 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v63
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+5)) = uint8(v63)
			}
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v27)+128))
			v190 = v27
			v191 = v68 - v27
			*(*int32)(unsafe.Add(mBase, uint32(v190)+8)) = v191
			m.G0 = v12 + int32(16)
			return v190
		}
	}
}
func F_AllocSetFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = l0 - int32(8)
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	if v15&int64(16) != int64(0) {
		v21 = l0 - int32(32)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		if v22 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
				F_errmsg_internal(m, int32(_a_F_AllocSetFree_0), v11)
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_AllocSetFree_1), int32(1080), int32(_a_F_AllocSetFree_2))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			if v25 != int32(474) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
					F_errmsg_internal(m, int32(_a_F_AllocSetFree_0), v11)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_AllocSetFree_1), int32(1080), int32(_a_F_AllocSetFree_2))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(20))))
				v32 = l0 - int32(16)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
				if v30 != v33 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
						F_errmsg_internal(m, int32(_a_F_AllocSetFree_0), v11)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_AllocSetFree_1), int32(1080), int32(_a_F_AllocSetFree_2))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(24))))
					v39 = l0 - int32(28)
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					if v40 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v37
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v37
					}
					if v37 != 0 {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
						*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v43
					} else {
					}
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
					*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v45 + (v21 - v46)
					F_emscripten_builtin_free(m, v21)
					mBase = m.M
					m.G0 = v11 + int32(16)
					return
				}
			}
		}
	} else {
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v14-base.I32_wrap_i64(int64(base.Ui64(v15)>>(uint(int64(34))%64)))&int32(1073741822))))
		v65 = v57 + base.I32_wrap_i64(int64(base.Ui64(v15)>>(uint(int64(5))%64)))<<(uint(int32(2))%32) + int32(48)
		v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v66
		*(*int32)(unsafe.Add(mBase, uint32(v65))) = v14
		m.G0 = v11 + int32(16)
		return
	}
}
func F_AlterEventTriggerOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
	v12 = v10 + v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+132))
	if v13 == l2 {
		m.G0 = v8 + int32(16)
		return
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_AlterEventTriggerOwner_internal[0]))
		v19 = F_object_ownercheck(m, int32(3466), v16, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			if v19 == int32(0) {
				F_aclcheck_error(m, int32(2), int32(14), v12+int32(4))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = F_superuser_arg(m, l2)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						if v29 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								F_errcode(m, int32(16797828))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12 + int32(4)
									F_errmsg(m, int32(_a_F_AlterEventTriggerOwner_internal_0), v8)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										F_errhint(m, int32(_a_F_AlterEventTriggerOwner_internal_1), int32(0))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_AlterEventTriggerOwner_internal_2), int32(558), int32(_a_F_AlterEventTriggerOwner_internal_3))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
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
							*(*int32)(unsafe.Add(mBase, uint32(v12)+132)) = l2
							F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								F_changeDependencyOnOwner(m, int32(3466), v39, l2)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, _c_F_AlterEventTriggerOwner_internal[1]))
									if v43 == int32(0) {
										m.G0 = v8 + int32(16)
										return
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
										v48 = int32(0)
										F_RunObjectPostAlterHook(m, int32(3466), v47, v48, v48, v48)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				v29 = F_superuser_arg(m, l2)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					if v29 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12 + int32(4)
								F_errmsg(m, int32(_a_F_AlterEventTriggerOwner_internal_0), v8)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									F_errhint(m, int32(_a_F_AlterEventTriggerOwner_internal_1), int32(0))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_AlterEventTriggerOwner_internal_2), int32(558), int32(_a_F_AlterEventTriggerOwner_internal_3))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
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
						*(*int32)(unsafe.Add(mBase, uint32(v12)+132)) = l2
						F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							F_changeDependencyOnOwner(m, int32(3466), v39, l2)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, _c_F_AlterEventTriggerOwner_internal[1]))
								if v43 == int32(0) {
									m.G0 = v8 + int32(16)
									return
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									v48 = int32(0)
									F_RunObjectPostAlterHook(m, int32(3466), v47, v48, v48, v48)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return
									} else {
										m.G0 = v8 + int32(16)
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
func F_AppendJumble8(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
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
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int64
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int64
	_ = v713
	var v715 int32
	_ = v715
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v379 != int32(1024) {
		goto L66
	} else {
		goto L67
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v379 = v11
	goto L1
L3:
	;
	goto L4
L4:
	;
	v12 = int32(4)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v14-int32(1021)) < base.Ui32(v12) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v369
	v379 = v369
	goto L1
L6:
	;
	v23 = v14
	v24 = v12
	v27 = l0 + int32(28)
	goto L9
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+v13))) = v8
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v369 = v364 + int32(4)
	goto L5
L9:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v23) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v369 = v360
	goto L5
L11:
	;
	v30 = int32(1024)
	v37 = int32(-1636607408)
	goto L16
L12:
	;
	v352 = v23
	goto L13
L13:
	;
	v354 = int32(1024) - v352
	if base.Ui32(v24) < base.Ui32(v354) {
		goto L59
	} else {
		goto L60
	}
L14:
	;
	v347 = F_Int64GetDatum(m, base.I64_extend_i32_u(v337)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v337^v329-base.I32_rotl(v337, int32(24))))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L57
	} else {
		goto L58
	}
L15:
	;
	if v13&int32(3) != 0 {
		goto L31
	} else {
		goto L32
	}
L16:
	;
	goto L15
L19:
	;
	v315 = int32(14)
	v317 = v311 ^ v312 - base.I32_rotl(v311, v315)
	v321 = v317 ^ v310 - base.I32_rotl(v317, int32(11))
	v325 = v321 ^ v311 - base.I32_rotl(v321, int32(25))
	v329 = v325 ^ v317 - base.I32_rotl(v325, int32(16))
	v333 = v329 ^ v321 - base.I32_rotl(v329, int32(4))
	v337 = v333 ^ v325 - base.I32_rotl(v333, v315)
	goto L14
L20:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v310 = v302 + v305
	v311 = v303
	v312 = v304
	goto L19
L21:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	v302 = v298<<(uint(int32(8))%32) + v295
	v303 = v296
	v304 = v297
	goto L20
L22:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
	v295 = v291<<(uint(int32(16))%32) + v288
	v296 = v289
	v297 = v290
	goto L21
L23:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
	v288 = v284<<(uint(int32(24))%32) + v120
	v289 = v282
	v290 = v283
	goto L22
L24:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
	v282 = v278 + v280
	v283 = v279
	goto L23
L25:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
	v278 = v274<<(uint(int32(8))%32) + v272
	v279 = v273
	goto L24
L26:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
	v272 = v268<<(uint(int32(16))%32) + v266
	v273 = v267
	goto L25
L27:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+7)))
	v266 = v262<<(uint(int32(24))%32) + v121
	v267 = v261
	goto L26
L28:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+8)))
	v261 = v257<<(uint(int32(8))%32) + v256
	goto L27
L29:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+9)))
	v256 = v252<<(uint(int32(16))%32) + v251
	goto L28
L30:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+10)))
	v251 = v247<<(uint(int32(24))%32) + v125
	goto L29
L31:
	;
	goto L34
L32:
	;
	goto L33
L33:
	;
	goto L40
L34:
	;
	v83 = v13
	v84 = v30
	v86 = v37
	v87 = v37
	v88 = v37
	goto L37
L36:
	;
	switch v129 - int32(1) {
	case 0:
		v302 = v120
		v303 = v121
		v304 = v125
		goto L20
	case 1:
		v295 = v120
		v296 = v121
		v297 = v125
		goto L21
	case 2:
		v288 = v120
		v289 = v121
		v290 = v125
		goto L22
	case 3:
		v282 = v121
		v283 = v125
		goto L23
	case 4:
		v278 = v121
		v279 = v125
		goto L24
	case 5:
		v272 = v121
		v273 = v125
		goto L25
	case 6:
		v266 = v121
		v267 = v125
		goto L26
	case 7:
		v261 = v125
		goto L27
	case 8:
		v256 = v125
		goto L28
	case 9:
		v251 = v125
		goto L29
	case 10:
		goto L30
	default:
		v310 = v120
		v311 = v121
		v312 = v125
		goto L19
	}
L37:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v91 = v90 + v87
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v95 = v94 + v88
	v97 = int32(4)
	v99 = v92 + v86 - v95 ^ base.I32_rotl(v95, v97)
	v103 = v91 - v99 ^ base.I32_rotl(v99, int32(6))
	v104 = v95 + v91
	v105 = v99 + v104
	v106 = v103 + v105
	v110 = v104 - v103 ^ base.I32_rotl(v103, int32(8))
	v114 = v105 - v110 ^ base.I32_rotl(v110, int32(16))
	v118 = v106 - v114 ^ base.I32_rotl(v114, int32(19))
	v119 = v110 + v106
	v120 = v114 + v119
	v121 = v118 + v120
	v125 = v119 - v118 ^ base.I32_rotl(v118, v97)
	v126 = int32(12)
	v127 = v83 + v126
	v129 = v84 - v126
	if base.Ui32(int32(11)) < base.Ui32(v129) {
		v83 = v127
		v84 = v129
		v86 = v120
		v87 = v121
		v88 = v125
		goto L37
	} else {
		goto L39
	}
L38:
	;
	goto L36
L39:
	;
	goto L38
L40:
	;
	v143 = v13
	v144 = v30
	v146 = v37
	v147 = v37
	v148 = v37
	goto L43
L42:
	;
	switch v189 - int32(1) {
	case 0:
		v244 = v180
		goto L46
	case 1:
		v239 = v180
		goto L47
	case 2:
		goto L48
	case 3:
		v232 = v181
		goto L49
	case 4:
		v229 = v181
		goto L50
	case 5:
		v224 = v181
		goto L51
	case 6:
		goto L52
	case 7:
		v215 = v185
		goto L53
	case 8:
		v210 = v185
		goto L54
	case 9:
		v205 = v185
		goto L55
	case 10:
		goto L56
	default:
		v310 = v180
		v311 = v181
		v312 = v185
		goto L19
	}
L43:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v151 = v150 + v147
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	v155 = v154 + v148
	v157 = int32(4)
	v159 = v152 + v146 - v155 ^ base.I32_rotl(v155, v157)
	v163 = v151 - v159 ^ base.I32_rotl(v159, int32(6))
	v164 = v155 + v151
	v165 = v159 + v164
	v166 = v163 + v165
	v170 = v164 - v163 ^ base.I32_rotl(v163, int32(8))
	v174 = v165 - v170 ^ base.I32_rotl(v170, int32(16))
	v178 = v166 - v174 ^ base.I32_rotl(v174, int32(19))
	v179 = v170 + v166
	v180 = v174 + v179
	v181 = v178 + v180
	v185 = v179 - v178 ^ base.I32_rotl(v178, v157)
	v186 = int32(12)
	v187 = v143 + v186
	v189 = v144 - v186
	if base.Ui32(int32(11)) < base.Ui32(v189) {
		v143 = v187
		v144 = v189
		v146 = v180
		v147 = v181
		v148 = v185
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	goto L44
L46:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v310 = v244 + v245
	v311 = v181
	v312 = v185
	goto L19
L47:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	v244 = v240<<(uint(int32(8))%32) + v239
	goto L46
L48:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
	v239 = v235<<(uint(int32(16))%32) + v180
	goto L47
L49:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v310 = v233 + v180
	v311 = v232
	v312 = v185
	goto L19
L50:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
	v232 = v229 + v230
	goto L49
L51:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
	v229 = v225<<(uint(int32(8))%32) + v224
	goto L50
L52:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+6)))
	v224 = v220<<(uint(int32(16))%32) + v181
	goto L51
L53:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v310 = v216 + v180
	v311 = v218 + v181
	v312 = v215
	goto L19
L54:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
	v215 = v211<<(uint(int32(8))%32) + v210
	goto L53
L55:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+9)))
	v210 = v206<<(uint(int32(16))%32) + v205
	goto L54
L56:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+10)))
	v205 = v201<<(uint(int32(24))%32) + v185
	goto L55
L57:
	;
	return
L58:
	;
	v349 = *(*int64)(unsafe.Add(mBase, uint32(v347)))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v349
	v352 = int32(8)
	goto L13
L59:
	;
	v356 = v24
	goto L61
L60:
	;
	v356 = v354
	goto L61
L61:
	;
	if v356 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	base.MemoryCopy(m, v352+v13, v27, v356)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v360 = v352 + v356
	v361 = v24 - v356
	if v361 != 0 {
		v23 = v360
		v24 = v361
		v27 = v356 + v27
		goto L9
	} else {
		goto L65
	}
L65:
	;
	goto L10
L66:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	*(*uint8)(unsafe.Add(mBase, uint32(v379+v384))) = uint8(v388)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v390 + int32(1)
	return
L67:
	;
	goto L68
L68:
	;
	v394 = int32(1024)
	v401 = int32(-1636607408)
	goto L71
L69:
	;
	v711 = F_Int64GetDatum(m, base.I64_extend_i32_u(v701)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v701^v693-base.I32_rotl(v701, int32(24))))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L57
	} else {
		goto L112
	}
L70:
	;
	if v384&int32(3) != 0 {
		goto L86
	} else {
		goto L87
	}
L71:
	;
	goto L70
L74:
	;
	v679 = int32(14)
	v681 = v675 ^ v676 - base.I32_rotl(v675, v679)
	v685 = v681 ^ v674 - base.I32_rotl(v681, int32(11))
	v689 = v685 ^ v675 - base.I32_rotl(v685, int32(25))
	v693 = v689 ^ v681 - base.I32_rotl(v689, int32(16))
	v697 = v693 ^ v685 - base.I32_rotl(v693, int32(4))
	v701 = v697 ^ v689 - base.I32_rotl(v697, v679)
	goto L69
L75:
	;
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491))))
	v674 = v666 + v669
	v675 = v667
	v676 = v668
	goto L74
L76:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+1)))
	v666 = v662<<(uint(int32(8))%32) + v659
	v667 = v660
	v668 = v661
	goto L75
L77:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+2)))
	v659 = v655<<(uint(int32(16))%32) + v652
	v660 = v653
	v661 = v654
	goto L76
L78:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+3)))
	v652 = v648<<(uint(int32(24))%32) + v484
	v653 = v646
	v654 = v647
	goto L77
L79:
	;
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+4)))
	v646 = v642 + v644
	v647 = v643
	goto L78
L80:
	;
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+5)))
	v642 = v638<<(uint(int32(8))%32) + v636
	v643 = v637
	goto L79
L81:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+6)))
	v636 = v632<<(uint(int32(16))%32) + v630
	v637 = v631
	goto L80
L82:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+7)))
	v630 = v626<<(uint(int32(24))%32) + v485
	v631 = v625
	goto L81
L83:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+8)))
	v625 = v621<<(uint(int32(8))%32) + v620
	goto L82
L84:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+9)))
	v620 = v616<<(uint(int32(16))%32) + v615
	goto L83
L85:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+10)))
	v615 = v611<<(uint(int32(24))%32) + v489
	goto L84
L86:
	;
	goto L89
L87:
	;
	goto L88
L88:
	;
	goto L95
L89:
	;
	v447 = v384
	v448 = v394
	v450 = v401
	v451 = v401
	v452 = v401
	goto L92
L91:
	;
	switch v493 - int32(1) {
	case 0:
		v666 = v484
		v667 = v485
		v668 = v489
		goto L75
	case 1:
		v659 = v484
		v660 = v485
		v661 = v489
		goto L76
	case 2:
		v652 = v484
		v653 = v485
		v654 = v489
		goto L77
	case 3:
		v646 = v485
		v647 = v489
		goto L78
	case 4:
		v642 = v485
		v643 = v489
		goto L79
	case 5:
		v636 = v485
		v637 = v489
		goto L80
	case 6:
		v630 = v485
		v631 = v489
		goto L81
	case 7:
		v625 = v489
		goto L82
	case 8:
		v620 = v489
		goto L83
	case 9:
		v615 = v489
		goto L84
	case 10:
		goto L85
	default:
		v674 = v484
		v675 = v485
		v676 = v489
		goto L74
	}
L92:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v447)+4))
	v455 = v454 + v451
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v447)+8))
	v459 = v458 + v452
	v461 = int32(4)
	v463 = v456 + v450 - v459 ^ base.I32_rotl(v459, v461)
	v467 = v455 - v463 ^ base.I32_rotl(v463, int32(6))
	v468 = v459 + v455
	v469 = v463 + v468
	v470 = v467 + v469
	v474 = v468 - v467 ^ base.I32_rotl(v467, int32(8))
	v478 = v469 - v474 ^ base.I32_rotl(v474, int32(16))
	v482 = v470 - v478 ^ base.I32_rotl(v478, int32(19))
	v483 = v474 + v470
	v484 = v478 + v483
	v485 = v482 + v484
	v489 = v483 - v482 ^ base.I32_rotl(v482, v461)
	v490 = int32(12)
	v491 = v447 + v490
	v493 = v448 - v490
	if base.Ui32(int32(11)) < base.Ui32(v493) {
		v447 = v491
		v448 = v493
		v450 = v484
		v451 = v485
		v452 = v489
		goto L92
	} else {
		goto L94
	}
L93:
	;
	goto L91
L94:
	;
	goto L93
L95:
	;
	v507 = v384
	v508 = v394
	v510 = v401
	v511 = v401
	v512 = v401
	goto L98
L97:
	;
	switch v553 - int32(1) {
	case 0:
		v608 = v544
		goto L101
	case 1:
		v603 = v544
		goto L102
	case 2:
		goto L103
	case 3:
		v596 = v545
		goto L104
	case 4:
		v593 = v545
		goto L105
	case 5:
		v588 = v545
		goto L106
	case 6:
		goto L107
	case 7:
		v579 = v549
		goto L108
	case 8:
		v574 = v549
		goto L109
	case 9:
		v569 = v549
		goto L110
	case 10:
		goto L111
	default:
		v674 = v544
		v675 = v545
		v676 = v549
		goto L74
	}
L98:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	v515 = v514 + v511
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v507)+8))
	v519 = v518 + v512
	v521 = int32(4)
	v523 = v516 + v510 - v519 ^ base.I32_rotl(v519, v521)
	v527 = v515 - v523 ^ base.I32_rotl(v523, int32(6))
	v528 = v519 + v515
	v529 = v523 + v528
	v530 = v527 + v529
	v534 = v528 - v527 ^ base.I32_rotl(v527, int32(8))
	v538 = v529 - v534 ^ base.I32_rotl(v534, int32(16))
	v542 = v530 - v538 ^ base.I32_rotl(v538, int32(19))
	v543 = v534 + v530
	v544 = v538 + v543
	v545 = v542 + v544
	v549 = v543 - v542 ^ base.I32_rotl(v542, v521)
	v550 = int32(12)
	v551 = v507 + v550
	v553 = v508 - v550
	if base.Ui32(int32(11)) < base.Ui32(v553) {
		v507 = v551
		v508 = v553
		v510 = v544
		v511 = v545
		v512 = v549
		goto L98
	} else {
		goto L100
	}
L99:
	;
	goto L97
L100:
	;
	goto L99
L101:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551))))
	v674 = v608 + v609
	v675 = v545
	v676 = v549
	goto L74
L102:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+1)))
	v608 = v604<<(uint(int32(8))%32) + v603
	goto L101
L103:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+2)))
	v603 = v599<<(uint(int32(16))%32) + v544
	goto L102
L104:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v551)))
	v674 = v597 + v544
	v675 = v596
	v676 = v549
	goto L74
L105:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+4)))
	v596 = v593 + v594
	goto L104
L106:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+5)))
	v593 = v589<<(uint(int32(8))%32) + v588
	goto L105
L107:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+6)))
	v588 = v584<<(uint(int32(16))%32) + v545
	goto L106
L108:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v551)))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v551)+4))
	v674 = v580 + v544
	v675 = v582 + v545
	v676 = v579
	goto L74
L109:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+8)))
	v579 = v575<<(uint(int32(8))%32) + v574
	goto L108
L110:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+9)))
	v574 = v570<<(uint(int32(16))%32) + v569
	goto L109
L111:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+10)))
	v569 = v565<<(uint(int32(24))%32) + v549
	goto L110
L112:
	;
	v713 = *(*int64)(unsafe.Add(mBase, uint32(v711)))
	*(*int64)(unsafe.Add(mBase, uint32(v384))) = v713
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	*(*uint8)(unsafe.Add(mBase, uint32(v384)+8)) = uint8(v715)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
	return
}
func F_ApplyLauncherWakeupAtCommit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ApplyLauncherWakeupAtCommit[0])))
	if v2 == int32(0) {
		v6 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_ApplyLauncherWakeupAtCommit[0])) = uint8(v6)
	} else {
	}
	return
}
func F_ApplyWorkerMain(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
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
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
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
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	v8 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[0])) = uint8(v8)
	F_SetupApplyOrSyncWorker(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v13 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[0])) = uint8(v13)
		v15 = m.G0
		v17 = v15 - int32(160)
		m.G0 = v17
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
		if v21 != 0 {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v22
			v25 = v17 + int32(96)
			v30 = F_pg_snprintf(m, v25, int32(64), int32(_a_F_ApplyWorkerMain_0), v17+int32(32))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				F_StartTransactionCommand(m)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v35 = F_replorigin_by_name(m, v25, int32(1))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						if v35 == int32(0) {
							v39 = F_replorigin_create(m, v25)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								v41 = v39
								F_replorigin_session_setup(m, v41, int32(0))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									*(*uint16)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[2])) = uint16(v41)
									v47 = F_replorigin_session_get_progress(m)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										F_CommitTransactionCommand(m)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return
										} else {
											v52 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
											v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+30)))
											if v53 == int32(1) {
												v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+24)))
												v59 = v56 ^ int32(1)
											} else {
												v59 = int32(0)
											}
											v61 = *(*int32)(unsafe.Add(mBase, uint32(v52)+36))
											v62 = int32(1)
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
											v70 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[3]))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
											v72 = m.T0[v71].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v61, v62, v62, v59&v62, v66, v17+int32(48))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4])) = v72
												if v72 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v264 = m.ExcPending
													if v264 != 0 {
														return
													} else {
														F_errcode(m, int32(100663808))
														mBase = m.M
														v267 = m.ExcPending
														if v267 != 0 {
															return
														} else {
															v269 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
															v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v17))) = v270
															v272 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v272
															F_errmsg(m, int32(_a_F_ApplyWorkerMain_1), v17)
															mBase = m.M
															v276 = m.ExcPending
															if v276 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_3), int32(_a_F_ApplyWorkerMain_4))
																mBase = m.M
																v281 = m.ExcPending
																if v281 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v81 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[3]))
													v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
													v83 = m.T0[v82].(func(*base.Module, int32, int32) int32)(m, v72, v17+int32(52))
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return
													} else {
														v87 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
														v90 = F_MemoryContextStrdup(m, v87, v17+int32(96))
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[6])) = v90
															*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v47
															v94 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v17)+56)) = uint8(v94)
															*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v21
															v98 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
															v100 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[3]))
															v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+24))
															v102 = m.T0[v101].(func(*base.Module, int32) int32)(m, v98)
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return
															} else {
																if v102 <= int32(_a_F_ApplyWorkerMain_5) {
																	if int32(_a_F_ApplyWorkerMain_6) < v102 {
																		v111 = int32(2)
																	} else {
																		v111 = int32(1)
																	}
																	if int32(_a_F_ApplyWorkerMain_7) < v102 {
																		v114 = int32(3)
																	} else {
																		v114 = v111
																	}
																	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v114
																	v117 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v118
																	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+26)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v17)+80)) = uint8(v120)
																	if int32(_a_F_ApplyWorkerMain_8) <= v102 {
																		v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+27)))
																		v139 = v117
																		v140 = v138
																		if v140&int32(255) != int32(102) {
																			v147 = int32(_a_F_ApplyWorkerMain_9)
																		} else {
																			v147 = int32(0)
																		}
																		v149 = v139
																		v150 = v147
																		v151 = int32(0)
																	} else {
																		v149 = v117
																		v150 = int32(0)
																		v151 = int32(0)
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(4)
																	v128 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v129
																	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+26)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v17)+80)) = uint8(v131)
																	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+27)))
																	if v133 != int32(112) {
																		v139 = v128
																		v140 = v133
																		if v140&int32(255) != int32(102) {
																			v147 = int32(_a_F_ApplyWorkerMain_9)
																		} else {
																			v147 = int32(0)
																		}
																		v149 = v139
																		v150 = v147
																		v151 = int32(0)
																	} else {
																		v149 = v128
																		v150 = int32(_a_F_ApplyWorkerMain_10)
																		v151 = int32(1)
																	}
																}
																*(*int32)(unsafe.Add(mBase, uint32(v17)+84)) = v150
																v154 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[7]))
																*(*uint8)(unsafe.Add(mBase, uint32(v154)+68)) = uint8(v151)
																v156 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v17)+88)) = uint8(v156)
																v158 = *(*int32)(unsafe.Add(mBase, uint32(v149)+52))
																v159 = F_pstrdup(m, v158)
																mBase = m.M
																v160 = m.ExcPending
																if v160 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = v159
																	v163 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+28)))
																	if v164 != int32(112) {
																		v202 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																		v206 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[3]))
																		v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+32))
																		v208 = m.T0[v207].(func(*base.Module, int32, int32) int32)(m, v202, v17+int32(56))
																		mBase = m.M
																		v209 = m.ExcPending
																		if v209 != 0 {
																			return
																		} else {
																			v212 = F_errstart(m, int32(14), int32(0))
																			mBase = m.M
																			v213 = m.ExcPending
																			if v213 != 0 {
																				return
																			} else {
																				if v212 != 0 {
																					v215 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																					v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
																					v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)))
																					switch v218 - int32(100) {
																					case 0:
																						v224 = int32(_a_F_ApplyWorkerMain_11)
																					case 1:
																						v224 = int32(_a_F_ApplyWorkerMain_12)
																					default:
																						v224 = int32(_a_F_ApplyWorkerMain_13)
																					case 12:
																						v224 = int32(_a_F_ApplyWorkerMain_14)
																					}
																					*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v224
																					*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v216
																					F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v17+int32(16))
																					mBase = m.M
																					v231 = m.ExcPending
																					if v231 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																						mBase = m.M
																						v236 = m.ExcPending
																						if v236 != 0 {
																							return
																						} else {
																							F_start_apply(m, v47)
																							mBase = m.M
																							v241 = m.ExcPending
																							if v241 != 0 {
																								return
																							} else {
																								m.G0 = v17 + int32(160)
																								F_proc_exit(m, int32(0))
																								mBase = m.M
																								v284 = m.ExcPending
																								if v284 != 0 {
																									return
																								} else {
																									base.Wasm_trap_unreachable()
																									for {
																									}
																								}
																							}
																						}
																					}
																				} else {
																					F_start_apply(m, v47)
																					mBase = m.M
																					v241 = m.ExcPending
																					if v241 != 0 {
																						return
																					} else {
																						m.G0 = v17 + int32(160)
																						F_proc_exit(m, int32(0))
																						mBase = m.M
																						v284 = m.ExcPending
																						if v284 != 0 {
																							return
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
																		v167 = F_AllTablesyncsReady(m)
																		mBase = m.M
																		v168 = m.ExcPending
																		if v168 != 0 {
																			return
																		} else {
																			if v167 == int32(0) {
																				v202 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																				v206 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[3]))
																				v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+32))
																				v208 = m.T0[v207].(func(*base.Module, int32, int32) int32)(m, v202, v17+int32(56))
																				mBase = m.M
																				v209 = m.ExcPending
																				if v209 != 0 {
																					return
																				} else {
																					v212 = F_errstart(m, int32(14), int32(0))
																					mBase = m.M
																					v213 = m.ExcPending
																					if v213 != 0 {
																						return
																					} else {
																						if v212 != 0 {
																							v215 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																							v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
																							v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)))
																							switch v218 - int32(100) {
																							case 0:
																								v224 = int32(_a_F_ApplyWorkerMain_11)
																							case 1:
																								v224 = int32(_a_F_ApplyWorkerMain_12)
																							default:
																								v224 = int32(_a_F_ApplyWorkerMain_13)
																							case 12:
																								v224 = int32(_a_F_ApplyWorkerMain_14)
																							}
																							*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v224
																							*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v216
																							F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v17+int32(16))
																							mBase = m.M
																							v231 = m.ExcPending
																							if v231 != 0 {
																								return
																							} else {
																								F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																								mBase = m.M
																								v236 = m.ExcPending
																								if v236 != 0 {
																									return
																								} else {
																									F_start_apply(m, v47)
																									mBase = m.M
																									v241 = m.ExcPending
																									if v241 != 0 {
																										return
																									} else {
																										m.G0 = v17 + int32(160)
																										F_proc_exit(m, int32(0))
																										mBase = m.M
																										v284 = m.ExcPending
																										if v284 != 0 {
																											return
																										} else {
																											base.Wasm_trap_unreachable()
																											for {
																											}
																										}
																									}
																								}
																							}
																						} else {
																							F_start_apply(m, v47)
																							mBase = m.M
																							v241 = m.ExcPending
																							if v241 != 0 {
																								return
																							} else {
																								m.G0 = v17 + int32(160)
																								F_proc_exit(m, int32(0))
																								mBase = m.M
																								v284 = m.ExcPending
																								if v284 != 0 {
																									return
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
																				v171 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v17)+88)) = uint8(v171)
																				v174 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																				v178 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[3]))
																				v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+32))
																				v180 = m.T0[v179].(func(*base.Module, int32, int32) int32)(m, v174, v17+int32(56))
																				mBase = m.M
																				v181 = m.ExcPending
																				if v181 != 0 {
																					return
																				} else {
																					F_StartTransactionCommand(m)
																					mBase = m.M
																					v183 = m.ExcPending
																					if v183 != 0 {
																						return
																					} else {
																						v184 = F_GetTransactionSnapshot(m)
																						mBase = m.M
																						v185 = m.ExcPending
																						if v185 != 0 {
																							return
																						} else {
																							F_PushActiveSnapshot(m, v184)
																							mBase = m.M
																							v187 = m.ExcPending
																							if v187 != 0 {
																								return
																							} else {
																								v189 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																								v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
																								F_UpdateTwoPhaseState(m, v190)
																								mBase = m.M
																								v192 = m.ExcPending
																								if v192 != 0 {
																									return
																								} else {
																									v194 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																									v195 = int32(101)
																									*(*uint8)(unsafe.Add(mBase, uint32(v194)+28)) = uint8(v195)
																									F_PopActiveSnapshot(m)
																									mBase = m.M
																									v198 = m.ExcPending
																									if v198 != 0 {
																										return
																									} else {
																										F_CommitTransactionCommand(m)
																										mBase = m.M
																										v200 = m.ExcPending
																										if v200 != 0 {
																											return
																										} else {
																											v212 = F_errstart(m, int32(14), int32(0))
																											mBase = m.M
																											v213 = m.ExcPending
																											if v213 != 0 {
																												return
																											} else {
																												if v212 != 0 {
																													v215 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																													v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
																													v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)))
																													switch v218 - int32(100) {
																													case 0:
																														v224 = int32(_a_F_ApplyWorkerMain_11)
																													case 1:
																														v224 = int32(_a_F_ApplyWorkerMain_12)
																													default:
																														v224 = int32(_a_F_ApplyWorkerMain_13)
																													case 12:
																														v224 = int32(_a_F_ApplyWorkerMain_14)
																													}
																													*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v224
																													*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v216
																													F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v17+int32(16))
																													mBase = m.M
																													v231 = m.ExcPending
																													if v231 != 0 {
																														return
																													} else {
																														F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																														mBase = m.M
																														v236 = m.ExcPending
																														if v236 != 0 {
																															return
																														} else {
																															F_start_apply(m, v47)
																															mBase = m.M
																															v241 = m.ExcPending
																															if v241 != 0 {
																																return
																															} else {
																																m.G0 = v17 + int32(160)
																																F_proc_exit(m, int32(0))
																																mBase = m.M
																																v284 = m.ExcPending
																																if v284 != 0 {
																																	return
																																} else {
																																	base.Wasm_trap_unreachable()
																																	for {
																																	}
																																}
																															}
																														}
																													}
																												} else {
																													F_start_apply(m, v47)
																													mBase = m.M
																													v241 = m.ExcPending
																													if v241 != 0 {
																														return
																													} else {
																														m.G0 = v17 + int32(160)
																														F_proc_exit(m, int32(0))
																														mBase = m.M
																														v284 = m.ExcPending
																														if v284 != 0 {
																															return
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
												}
											}
										}
									}
								}
							}
						} else {
							v41 = v35
							F_replorigin_session_setup(m, v41, int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								*(*uint16)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[2])) = uint16(v41)
								v47 = F_replorigin_session_get_progress(m)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									F_CommitTransactionCommand(m)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
										v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+30)))
										if v53 == int32(1) {
											v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+24)))
											v59 = v56 ^ int32(1)
										} else {
											v59 = int32(0)
										}
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v52)+36))
										v62 = int32(1)
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
										v70 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[3]))
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
										v72 = m.T0[v71].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v61, v62, v62, v59&v62, v66, v17+int32(48))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4])) = v72
											if v72 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v264 = m.ExcPending
												if v264 != 0 {
													return
												} else {
													F_errcode(m, int32(100663808))
													mBase = m.M
													v267 = m.ExcPending
													if v267 != 0 {
														return
													} else {
														v269 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
														v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+16))
														*(*int32)(unsafe.Add(mBase, uint32(v17))) = v270
														v272 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
														*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v272
														F_errmsg(m, int32(_a_F_ApplyWorkerMain_1), v17)
														mBase = m.M
														v276 = m.ExcPending
														if v276 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_3), int32(_a_F_ApplyWorkerMain_4))
															mBase = m.M
															v281 = m.ExcPending
															if v281 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v81 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[3]))
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
												v83 = m.T0[v82].(func(*base.Module, int32, int32) int32)(m, v72, v17+int32(52))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
													v90 = F_MemoryContextStrdup(m, v87, v17+int32(96))
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[6])) = v90
														*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v47
														v94 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v17)+56)) = uint8(v94)
														*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v21
														v98 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
														v100 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[3]))
														v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+24))
														v102 = m.T0[v101].(func(*base.Module, int32) int32)(m, v98)
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return
														} else {
															if v102 <= int32(_a_F_ApplyWorkerMain_5) {
																if int32(_a_F_ApplyWorkerMain_6) < v102 {
																	v111 = int32(2)
																} else {
																	v111 = int32(1)
																}
																if int32(_a_F_ApplyWorkerMain_7) < v102 {
																	v114 = int32(3)
																} else {
																	v114 = v111
																}
																*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v114
																v117 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+48))
																*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v118
																v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+26)))
																*(*uint8)(unsafe.Add(mBase, uint32(v17)+80)) = uint8(v120)
																if int32(_a_F_ApplyWorkerMain_8) <= v102 {
																	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+27)))
																	v139 = v117
																	v140 = v138
																	if v140&int32(255) != int32(102) {
																		v147 = int32(_a_F_ApplyWorkerMain_9)
																	} else {
																		v147 = int32(0)
																	}
																	v149 = v139
																	v150 = v147
																	v151 = int32(0)
																} else {
																	v149 = v117
																	v150 = int32(0)
																	v151 = int32(0)
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(4)
																v128 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+48))
																*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v129
																v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+26)))
																*(*uint8)(unsafe.Add(mBase, uint32(v17)+80)) = uint8(v131)
																v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+27)))
																if v133 != int32(112) {
																	v139 = v128
																	v140 = v133
																	if v140&int32(255) != int32(102) {
																		v147 = int32(_a_F_ApplyWorkerMain_9)
																	} else {
																		v147 = int32(0)
																	}
																	v149 = v139
																	v150 = v147
																	v151 = int32(0)
																} else {
																	v149 = v128
																	v150 = int32(_a_F_ApplyWorkerMain_10)
																	v151 = int32(1)
																}
															}
															*(*int32)(unsafe.Add(mBase, uint32(v17)+84)) = v150
															v154 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[7]))
															*(*uint8)(unsafe.Add(mBase, uint32(v154)+68)) = uint8(v151)
															v156 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v17)+88)) = uint8(v156)
															v158 = *(*int32)(unsafe.Add(mBase, uint32(v149)+52))
															v159 = F_pstrdup(m, v158)
															mBase = m.M
															v160 = m.ExcPending
															if v160 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = v159
																v163 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+28)))
																if v164 != int32(112) {
																	v202 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																	v206 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[3]))
																	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+32))
																	v208 = m.T0[v207].(func(*base.Module, int32, int32) int32)(m, v202, v17+int32(56))
																	mBase = m.M
																	v209 = m.ExcPending
																	if v209 != 0 {
																		return
																	} else {
																		v212 = F_errstart(m, int32(14), int32(0))
																		mBase = m.M
																		v213 = m.ExcPending
																		if v213 != 0 {
																			return
																		} else {
																			if v212 != 0 {
																				v215 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																				v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
																				v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)))
																				switch v218 - int32(100) {
																				case 0:
																					v224 = int32(_a_F_ApplyWorkerMain_11)
																				case 1:
																					v224 = int32(_a_F_ApplyWorkerMain_12)
																				default:
																					v224 = int32(_a_F_ApplyWorkerMain_13)
																				case 12:
																					v224 = int32(_a_F_ApplyWorkerMain_14)
																				}
																				*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v224
																				*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v216
																				F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v17+int32(16))
																				mBase = m.M
																				v231 = m.ExcPending
																				if v231 != 0 {
																					return
																				} else {
																					F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																					mBase = m.M
																					v236 = m.ExcPending
																					if v236 != 0 {
																						return
																					} else {
																						F_start_apply(m, v47)
																						mBase = m.M
																						v241 = m.ExcPending
																						if v241 != 0 {
																							return
																						} else {
																							m.G0 = v17 + int32(160)
																							F_proc_exit(m, int32(0))
																							mBase = m.M
																							v284 = m.ExcPending
																							if v284 != 0 {
																								return
																							} else {
																								base.Wasm_trap_unreachable()
																								for {
																								}
																							}
																						}
																					}
																				}
																			} else {
																				F_start_apply(m, v47)
																				mBase = m.M
																				v241 = m.ExcPending
																				if v241 != 0 {
																					return
																				} else {
																					m.G0 = v17 + int32(160)
																					F_proc_exit(m, int32(0))
																					mBase = m.M
																					v284 = m.ExcPending
																					if v284 != 0 {
																						return
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
																	v167 = F_AllTablesyncsReady(m)
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return
																	} else {
																		if v167 == int32(0) {
																			v202 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																			v206 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[3]))
																			v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+32))
																			v208 = m.T0[v207].(func(*base.Module, int32, int32) int32)(m, v202, v17+int32(56))
																			mBase = m.M
																			v209 = m.ExcPending
																			if v209 != 0 {
																				return
																			} else {
																				v212 = F_errstart(m, int32(14), int32(0))
																				mBase = m.M
																				v213 = m.ExcPending
																				if v213 != 0 {
																					return
																				} else {
																					if v212 != 0 {
																						v215 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																						v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
																						v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)))
																						switch v218 - int32(100) {
																						case 0:
																							v224 = int32(_a_F_ApplyWorkerMain_11)
																						case 1:
																							v224 = int32(_a_F_ApplyWorkerMain_12)
																						default:
																							v224 = int32(_a_F_ApplyWorkerMain_13)
																						case 12:
																							v224 = int32(_a_F_ApplyWorkerMain_14)
																						}
																						*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v224
																						*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v216
																						F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v17+int32(16))
																						mBase = m.M
																						v231 = m.ExcPending
																						if v231 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																							mBase = m.M
																							v236 = m.ExcPending
																							if v236 != 0 {
																								return
																							} else {
																								F_start_apply(m, v47)
																								mBase = m.M
																								v241 = m.ExcPending
																								if v241 != 0 {
																									return
																								} else {
																									m.G0 = v17 + int32(160)
																									F_proc_exit(m, int32(0))
																									mBase = m.M
																									v284 = m.ExcPending
																									if v284 != 0 {
																										return
																									} else {
																										base.Wasm_trap_unreachable()
																										for {
																										}
																									}
																								}
																							}
																						}
																					} else {
																						F_start_apply(m, v47)
																						mBase = m.M
																						v241 = m.ExcPending
																						if v241 != 0 {
																							return
																						} else {
																							m.G0 = v17 + int32(160)
																							F_proc_exit(m, int32(0))
																							mBase = m.M
																							v284 = m.ExcPending
																							if v284 != 0 {
																								return
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
																			v171 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v17)+88)) = uint8(v171)
																			v174 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																			v178 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[3]))
																			v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+32))
																			v180 = m.T0[v179].(func(*base.Module, int32, int32) int32)(m, v174, v17+int32(56))
																			mBase = m.M
																			v181 = m.ExcPending
																			if v181 != 0 {
																				return
																			} else {
																				F_StartTransactionCommand(m)
																				mBase = m.M
																				v183 = m.ExcPending
																				if v183 != 0 {
																					return
																				} else {
																					v184 = F_GetTransactionSnapshot(m)
																					mBase = m.M
																					v185 = m.ExcPending
																					if v185 != 0 {
																						return
																					} else {
																						F_PushActiveSnapshot(m, v184)
																						mBase = m.M
																						v187 = m.ExcPending
																						if v187 != 0 {
																							return
																						} else {
																							v189 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																							v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
																							F_UpdateTwoPhaseState(m, v190)
																							mBase = m.M
																							v192 = m.ExcPending
																							if v192 != 0 {
																								return
																							} else {
																								v194 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																								v195 = int32(101)
																								*(*uint8)(unsafe.Add(mBase, uint32(v194)+28)) = uint8(v195)
																								F_PopActiveSnapshot(m)
																								mBase = m.M
																								v198 = m.ExcPending
																								if v198 != 0 {
																									return
																								} else {
																									F_CommitTransactionCommand(m)
																									mBase = m.M
																									v200 = m.ExcPending
																									if v200 != 0 {
																										return
																									} else {
																										v212 = F_errstart(m, int32(14), int32(0))
																										mBase = m.M
																										v213 = m.ExcPending
																										if v213 != 0 {
																											return
																										} else {
																											if v212 != 0 {
																												v215 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																												v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
																												v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+28)))
																												switch v218 - int32(100) {
																												case 0:
																													v224 = int32(_a_F_ApplyWorkerMain_11)
																												case 1:
																													v224 = int32(_a_F_ApplyWorkerMain_12)
																												default:
																													v224 = int32(_a_F_ApplyWorkerMain_13)
																												case 12:
																													v224 = int32(_a_F_ApplyWorkerMain_14)
																												}
																												*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v224
																												*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v216
																												F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v17+int32(16))
																												mBase = m.M
																												v231 = m.ExcPending
																												if v231 != 0 {
																													return
																												} else {
																													F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																													mBase = m.M
																													v236 = m.ExcPending
																													if v236 != 0 {
																														return
																													} else {
																														F_start_apply(m, v47)
																														mBase = m.M
																														v241 = m.ExcPending
																														if v241 != 0 {
																															return
																														} else {
																															m.G0 = v17 + int32(160)
																															F_proc_exit(m, int32(0))
																															mBase = m.M
																															v284 = m.ExcPending
																															if v284 != 0 {
																																return
																															} else {
																																base.Wasm_trap_unreachable()
																																for {
																																}
																															}
																														}
																													}
																												}
																											} else {
																												F_start_apply(m, v47)
																												mBase = m.M
																												v241 = m.ExcPending
																												if v241 != 0 {
																													return
																												} else {
																													m.G0 = v17 + int32(160)
																													F_proc_exit(m, int32(0))
																													mBase = m.M
																													v284 = m.ExcPending
																													if v284 != 0 {
																														return
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
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v248 = m.ExcPending
			if v248 != 0 {
				return
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v251 = m.ExcPending
				if v251 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_ApplyWorkerMain_17), int32(0))
					mBase = m.M
					v255 = m.ExcPending
					if v255 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_18), int32(_a_F_ApplyWorkerMain_4))
						mBase = m.M
						v260 = m.ExcPending
						if v260 != 0 {
							return
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
func F_AtSubAbort_Portals(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = v9 + int32(12)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_AtSubAbort_Portals[0]))
	F_hash_seq_init(m, v12, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = F_hash_seq_search(m, v12)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = v17
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v9 + int32(32)
	return
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	if l0 != v26 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v111 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L47
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	if v28 != l0 {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v25)+80))
	if v85&int32(-2) == int32(2) {
		goto L35
	} else {
		goto L36
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = l1
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+80))
	if v31 == int32(3) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v49 == int32(0) {
		goto L9
	} else {
		goto L21
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+80)) = int32(5)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v36 == int32(0) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v45 = v31
	goto L17
L17:
	;
	if v45 != int32(5) {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	m.T0[v36].(func(*base.Module, int32))(m, v25)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v25)+80))
	v45 = v43
	goto L17
L20:
	;
	goto L14
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v54 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = int32(0)
	goto L9
L23:
	;
	if l2 != 0 {
		goto L32
	} else {
		goto L33
	}
L24:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v57 == v49 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v59
	goto L23
L26:
	;
	goto L27
L27:
	;
	v64 = v57
	goto L28
L28:
	;
	if v64 == int32(0) {
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v69
	goto L23
L30:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if v49 != v67 {
		v64 = v67
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = l2
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v49
	goto L22
L33:
	;
	goto L34
L34:
	;
	v79 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v79
	goto L22
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+80)) = int32(5)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v92 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	m.T0[v92].(func(*base.Module, int32))(m, v25)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v25)+60))
	if v97 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = int32(0)
	goto L40
L42:
	;
	F_ReleaseCachedPlan(m, v97, int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = int32(0)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	F_MemoryContextDeleteChildren(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+56)) = int64(0)
	goto L44
L46:
	;
	goto L9
L47:
	;
	if v111 != 0 {
		v22 = v111
		goto L7
	} else {
		goto L48
	}
L48:
	;
	goto L8
}
func F_AutoVacLauncherShutdown(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	v3 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		if v3 != 0 {
			F_errmsg_internal(m, int32(_a_F_AutoVacLauncherShutdown_0), int32(0))
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_AutoVacLauncherShutdown_1), int32(795), int32(_a_F_AutoVacLauncherShutdown_2))
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherShutdown[0]))
					v16 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v16
					F_proc_exit(m, v16)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherShutdown[0]))
			v16 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v16
			F_proc_exit(m, v16)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_a_swap(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return
}
func F_accumArrayResultAny(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v6 = int32(0)
	if l0 == v6 {
		v10 = F_get_array_type(m, l3)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if v10 == int32(0) {
				v18 = F_initArrayResultArr(m, l3, int32(0), l4, int32(1))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v53 = v18
					v54 = v6
					v55 = v18
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
					v58 = F_MemoryContextAlloc(m, v56, int32(8))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v58))) = v54
						v63 = v58
						v64 = v54
						if v64 != 0 {
							v66 = F_accumArrayResult(m, v64, l1, l2, l3, l4)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								return v63
							}
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
							v70 = F_accumArrayResultArr(m, v69, l1, l2, l3, l4)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								return v63
							}
						}
					}
				}
			} else {
				v24 = F_AllocSetContextCreateInternal(m, l4, int32(_a_F_accumArrayResultAny_0), int32(0), int32(_a_F_accumArrayResultAny_1), int32(_a_F_accumArrayResultAny_2))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v27 = F_MemoryContextAlloc(m, v24, int32(32))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v27)+28)) = uint8(v29)
						*(*int32)(unsafe.Add(mBase, uint32(v27))) = v24
						*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = int32(64)
						v35 = F_MemoryContextAlloc(m, v24, int32(256))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v35
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
							v39 = F_MemoryContextAlloc(m, v24, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = l3
								*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v39
								F_get_typlenbyvalalign(m, l3, v27+int32(24), v27+int32(26), v27+int32(27))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v53 = v27
									v54 = v27
									v55 = v6
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
									v58 = F_MemoryContextAlloc(m, v56, int32(8))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v55
										*(*int32)(unsafe.Add(mBase, uint32(v58))) = v54
										v63 = v58
										v64 = v54
										if v64 != 0 {
											v66 = F_accumArrayResult(m, v64, l1, l2, l3, l4)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												return v63
											}
										} else {
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
											v70 = F_accumArrayResultArr(m, v69, l1, l2, l3, l4)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												return v63
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
	} else {
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v63 = l0
		v64 = v62
		if v64 != 0 {
			v66 = F_accumArrayResult(m, v64, l1, l2, l3, l4)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				return v63
			}
		} else {
			v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v70 = F_accumArrayResultArr(m, v69, l1, l2, l3, l4)
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int32(0)
			} else {
				return v63
			}
		}
	}
}
func F_accum_sum_add(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v218 int32
	_ = v218
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v423 int32
	_ = v423
	v3 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v15 == int32(_a_F_accum_sum_add_0) {
		v19 = v14 - int32(1)
		if v19 < int32(0) {
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v19 != 0 {
				v31 = v19
				v38 = v3
				v39 = v3
				for {
					v42 = v22 + v31<<(uint(int32(2))%32)
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v44 = v43 + v38
					if v44 < int32(_a_F_accum_sum_add_1) {
						v53 = v44
						v54 = int32(0)
					} else {
						v49 = base.I32_div_u_s(v44, int32(_a_F_accum_sum_add_1))
						v53 = v49*int32(-10000) + v44
						v54 = v49
					}
					*(*int32)(unsafe.Add(mBase, uint32(v42))) = v53
					v58 = v42 - int32(4)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
					v60 = v54 + v59
					if int32(_a_F_accum_sum_add_1) <= v60 {
						v64 = base.I32_div_u_s(v60, int32(_a_F_accum_sum_add_1))
						v68 = v64*int32(-10000) + v60
						v69 = v64
					} else {
						v68 = v60
						v69 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v58))) = v68
					v71 = int32(2)
					v72 = v31 - v71
					v74 = v39 + v71
					if v74 != v14&int32(-2) {
						v31 = v72
						v38 = v69
						v39 = v74
						continue
					} else {
						break
					}
					break
				}
				if v14&int32(1) == int32(0) {
					v104 = v68
				} else {
					v82 = v72
					v89 = v69
					v93 = v22 + v82<<(uint(int32(2))%32)
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
					v95 = v94 + v89
					v97 = base.I32_rem_u_s(v95, int32(_a_F_accum_sum_add_1))
					if int32(_a_F_accum_sum_add_0) < v95 {
						v100 = v97
					} else {
						v100 = v95
					}
					*(*int32)(unsafe.Add(mBase, uint32(v93))) = v100
					v104 = v100
				}
			} else {
				v82 = v19
				v89 = v3
				v93 = v22 + v82<<(uint(int32(2))%32)
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
				v95 = v94 + v89
				v97 = base.I32_rem_u_s(v95, int32(_a_F_accum_sum_add_1))
				if int32(_a_F_accum_sum_add_0) < v95 {
					v100 = v97
				} else {
					v100 = v95
				}
				*(*int32)(unsafe.Add(mBase, uint32(v93))) = v100
				v104 = v100
			}
			if int32(0) < v104 {
				v117 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v117)
			} else {
			}
			v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v120 = int32(0)
			if v14 != int32(1) {
				v131 = v19
				v135 = v120
				v139 = int32(0)
				for {
					v143 = v119 + v131<<(uint(int32(2))%32)
					v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
					v145 = v144 + v135
					if v145 < int32(_a_F_accum_sum_add_1) {
						v154 = v145
						v155 = int32(0)
					} else {
						v150 = base.I32_div_u_s(v145, int32(_a_F_accum_sum_add_1))
						v154 = v150*int32(-10000) + v145
						v155 = v150
					}
					*(*int32)(unsafe.Add(mBase, uint32(v143))) = v154
					v159 = v143 - int32(4)
					v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
					v161 = v160 + v155
					if int32(_a_F_accum_sum_add_1) <= v161 {
						v165 = base.I32_div_u_s(v161, int32(_a_F_accum_sum_add_1))
						v169 = v165*int32(-10000) + v161
						v170 = v165
					} else {
						v169 = v161
						v170 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v159))) = v169
					v172 = int32(2)
					v173 = v131 - v172
					v175 = v139 + v172
					if v175 != v14&int32(-2) {
						v131 = v173
						v135 = v170
						v139 = v175
						continue
					} else {
						break
					}
					break
				}
				if v14&int32(1) == int32(0) {
					v205 = v169
				} else {
					v182 = v173
					v186 = v170
					v194 = v119 + v182<<(uint(int32(2))%32)
					v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
					v196 = v195 + v186
					v198 = base.I32_rem_u_s(v196, int32(_a_F_accum_sum_add_1))
					if int32(_a_F_accum_sum_add_0) < v196 {
						v201 = v198
					} else {
						v201 = v196
					}
					*(*int32)(unsafe.Add(mBase, uint32(v194))) = v201
					v205 = v201
				}
			} else {
				v182 = v19
				v186 = v120
				v194 = v119 + v182<<(uint(int32(2))%32)
				v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
				v196 = v195 + v186
				v198 = base.I32_rem_u_s(v196, int32(_a_F_accum_sum_add_1))
				if int32(_a_F_accum_sum_add_0) < v196 {
					v201 = v198
				} else {
					v201 = v196
				}
				*(*int32)(unsafe.Add(mBase, uint32(v194))) = v201
				v205 = v201
			}
			if v205 <= int32(0) {
			} else {
				v218 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v218)
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
		v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v242 = v235
	} else {
		v242 = v14
	}
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v250 <= v249 {
		v253 = v249 + int32(1)
		v261 = v253
		v262 = v253 + (v242 - v250)
	} else {
		v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if v256 != 0 {
			v261 = v250
			v262 = v242
		} else {
			v257 = int32(1)
			v261 = v250 + v257
			v262 = v242 + v257
		}
	}
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v264 = int32(-1)
	v266 = v263 + (v249 ^ v264)
	v269 = v262 + (v261 ^ v264)
	if v269 < v266 {
		v273 = v266 - v269
	} else {
		v273 = int32(0)
	}
	v274 = v262 + v273
	if base.B2i32(v261 != v250)|base.B2i32(v242 != v274) == int32(0) {
		v318 = v250
		v325 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v326 < v325 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v325
		} else {
		}
		v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v329 <= int32(0) {
		} else {
			v334 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v334 != 0 {
				v335 = int32(24)
			} else {
				v335 = int32(20)
			}
			v337 = *(*int32)(unsafe.Add(mBase, uint32(l0+v335)))
			v338 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v339 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v340 = v318 - v339
			v341 = int32(0)
			if v329 != int32(1) {
				v351 = v341
				v352 = int32(0)
				v353 = v340
				for {
					v362 = int32(2)
					v364 = v337 + v353<<(uint(v362)%32)
					v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
					v368 = v338 + v351<<(uint(int32(1))%32)
					v369 = int32(*(*int16)(unsafe.Add(mBase, uint32(v368))))
					*(*int32)(unsafe.Add(mBase, uint32(v364))) = v365 + v369
					v373 = v364 + int32(4)
					v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
					v375 = int32(*(*int16)(unsafe.Add(mBase, uint32(v368)+2)))
					*(*int32)(unsafe.Add(mBase, uint32(v373))) = v374 + v375
					v379 = v351 + v362
					v381 = v353 + v362
					v383 = v352 + v362
					if v383 != v329&int32(2147483646) {
						v351 = v379
						v352 = v383
						v353 = v381
						continue
					} else {
						break
					}
					break
				}
				if v329&int32(1) == int32(0) {
				} else {
					v389 = v379
					v391 = v381
					v402 = v337 + v391<<(uint(int32(2))%32)
					v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
					v407 = int32(*(*int16)(unsafe.Add(mBase, uint32(v338+v389<<(uint(int32(1))%32)))))
					*(*int32)(unsafe.Add(mBase, uint32(v402))) = v403 + v407
				}
			} else {
				v389 = v341
				v391 = v340
				v402 = v337 + v391<<(uint(int32(2))%32)
				v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
				v407 = int32(*(*int16)(unsafe.Add(mBase, uint32(v338+v389<<(uint(int32(1))%32)))))
				*(*int32)(unsafe.Add(mBase, uint32(v402))) = v403 + v407
			}
		}
		v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v423 + int32(1)
		return
	} else {
		v281 = v274 << (uint(int32(2)) % 32)
		v282 = F_palloc0(m, v281)
		mBase = m.M
		v283 = m.ExcPending
		if v283 != 0 {
			return
		} else {
			v284 = F_palloc0(m, v281)
			mBase = m.M
			v285 = m.ExcPending
			if v285 != 0 {
				return
			} else {
				v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v286 != 0 {
					v288 = int32(2)
					v289 = (v261 - v250) << (uint(v288) % 32)
					v291 = v242 << (uint(v288) % 32)
					v292 = int32(0)
					v293 = base.B2i32(v291 == v292)
					if v293 == v292 {
						base.MemoryCopy(m, v282+v289, v286, v291)
					} else {
					}
					v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					F_pfree(m, v298)
					mBase = m.M
					v300 = m.ExcPending
					if v300 != 0 {
						return
					} else {
						if v293 == int32(0) {
							v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							base.MemoryCopy(m, v284+v289, v304, v291)
						} else {
						}
						v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						F_pfree(m, v306)
						mBase = m.M
						v308 = m.ExcPending
						if v308 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v284
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v282
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v261
							v315 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v315)
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v274
							v318 = v261
							v325 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v326 < v325 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v325
							} else {
							}
							v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if v329 <= int32(0) {
							} else {
								v334 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								if v334 != 0 {
									v335 = int32(24)
								} else {
									v335 = int32(20)
								}
								v337 = *(*int32)(unsafe.Add(mBase, uint32(l0+v335)))
								v338 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
								v339 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v340 = v318 - v339
								v341 = int32(0)
								if v329 != int32(1) {
									v351 = v341
									v352 = int32(0)
									v353 = v340
									for {
										v362 = int32(2)
										v364 = v337 + v353<<(uint(v362)%32)
										v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
										v368 = v338 + v351<<(uint(int32(1))%32)
										v369 = int32(*(*int16)(unsafe.Add(mBase, uint32(v368))))
										*(*int32)(unsafe.Add(mBase, uint32(v364))) = v365 + v369
										v373 = v364 + int32(4)
										v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
										v375 = int32(*(*int16)(unsafe.Add(mBase, uint32(v368)+2)))
										*(*int32)(unsafe.Add(mBase, uint32(v373))) = v374 + v375
										v379 = v351 + v362
										v381 = v353 + v362
										v383 = v352 + v362
										if v383 != v329&int32(2147483646) {
											v351 = v379
											v352 = v383
											v353 = v381
											continue
										} else {
											break
										}
										break
									}
									if v329&int32(1) == int32(0) {
									} else {
										v389 = v379
										v391 = v381
										v402 = v337 + v391<<(uint(int32(2))%32)
										v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
										v407 = int32(*(*int16)(unsafe.Add(mBase, uint32(v338+v389<<(uint(int32(1))%32)))))
										*(*int32)(unsafe.Add(mBase, uint32(v402))) = v403 + v407
									}
								} else {
									v389 = v341
									v391 = v340
									v402 = v337 + v391<<(uint(int32(2))%32)
									v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
									v407 = int32(*(*int16)(unsafe.Add(mBase, uint32(v338+v389<<(uint(int32(1))%32)))))
									*(*int32)(unsafe.Add(mBase, uint32(v402))) = v403 + v407
								}
							}
							v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v423 + int32(1)
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v284
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v282
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v261
					v315 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v315)
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v274
					v318 = v261
					v325 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v326 < v325 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v325
					} else {
					}
					v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v329 <= int32(0) {
					} else {
						v334 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						if v334 != 0 {
							v335 = int32(24)
						} else {
							v335 = int32(20)
						}
						v337 = *(*int32)(unsafe.Add(mBase, uint32(l0+v335)))
						v338 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v339 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v340 = v318 - v339
						v341 = int32(0)
						if v329 != int32(1) {
							v351 = v341
							v352 = int32(0)
							v353 = v340
							for {
								v362 = int32(2)
								v364 = v337 + v353<<(uint(v362)%32)
								v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
								v368 = v338 + v351<<(uint(int32(1))%32)
								v369 = int32(*(*int16)(unsafe.Add(mBase, uint32(v368))))
								*(*int32)(unsafe.Add(mBase, uint32(v364))) = v365 + v369
								v373 = v364 + int32(4)
								v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
								v375 = int32(*(*int16)(unsafe.Add(mBase, uint32(v368)+2)))
								*(*int32)(unsafe.Add(mBase, uint32(v373))) = v374 + v375
								v379 = v351 + v362
								v381 = v353 + v362
								v383 = v352 + v362
								if v383 != v329&int32(2147483646) {
									v351 = v379
									v352 = v383
									v353 = v381
									continue
								} else {
									break
								}
								break
							}
							if v329&int32(1) == int32(0) {
							} else {
								v389 = v379
								v391 = v381
								v402 = v337 + v391<<(uint(int32(2))%32)
								v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
								v407 = int32(*(*int16)(unsafe.Add(mBase, uint32(v338+v389<<(uint(int32(1))%32)))))
								*(*int32)(unsafe.Add(mBase, uint32(v402))) = v403 + v407
							}
						} else {
							v389 = v341
							v391 = v340
							v402 = v337 + v391<<(uint(int32(2))%32)
							v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
							v407 = int32(*(*int16)(unsafe.Add(mBase, uint32(v338+v389<<(uint(int32(1))%32)))))
							*(*int32)(unsafe.Add(mBase, uint32(v402))) = v403 + v407
						}
					}
					v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v423 + int32(1)
					return
				}
			}
		}
	}
}
func F_aclcheck_error(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	if l0 != 0 {
		switch l0 - int32(1) {
		case 0:
			switch l1 {
			case 0, 2, 3, 4, 5, 10, 11, 13, 31, 32, 33, 35, 40, 43, 44, 47, 48, 50:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l1
					F_errmsg_internal(m, int32(_a_F_aclcheck_error_0), v7+int32(32))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2790), int32(_a_F_aclcheck_error_2))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			case 1:
				v61 = int32(_a_F_aclcheck_error_3)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 6:
				v61 = int32(_a_F_aclcheck_error_4)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 7:
				v61 = int32(_a_F_aclcheck_error_5)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 8:
				v61 = int32(_a_F_aclcheck_error_6)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 9:
				v61 = int32(_a_F_aclcheck_error_7)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 12:
				v61 = int32(_a_F_aclcheck_error_8)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 14:
				v61 = int32(_a_F_aclcheck_error_9)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 15:
				v61 = int32(_a_F_aclcheck_error_10)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 16:
				v61 = int32(_a_F_aclcheck_error_11)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 17:
				v61 = int32(_a_F_aclcheck_error_12)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 18:
				v61 = int32(_a_F_aclcheck_error_13)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 19:
				v61 = int32(_a_F_aclcheck_error_14)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 20:
				v61 = int32(_a_F_aclcheck_error_15)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 21:
				v61 = int32(_a_F_aclcheck_error_16)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 22:
				v61 = int32(_a_F_aclcheck_error_17)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 23:
				v61 = int32(_a_F_aclcheck_error_18)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 24:
				v61 = int32(_a_F_aclcheck_error_19)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 25:
				v61 = int32(_a_F_aclcheck_error_20)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 26:
				v61 = int32(_a_F_aclcheck_error_21)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 27:
				v61 = int32(_a_F_aclcheck_error_22)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 28:
				v61 = int32(_a_F_aclcheck_error_23)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 29:
				v61 = int32(_a_F_aclcheck_error_24)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 30:
				v61 = int32(_a_F_aclcheck_error_25)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 34:
				v61 = int32(_a_F_aclcheck_error_26)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 36:
				v61 = int32(_a_F_aclcheck_error_27)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 37:
				v61 = int32(_a_F_aclcheck_error_28)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 38:
				v61 = int32(_a_F_aclcheck_error_29)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 39:
				v61 = int32(_a_F_aclcheck_error_30)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 41:
				v61 = int32(_a_F_aclcheck_error_31)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 42:
				v61 = int32(_a_F_aclcheck_error_32)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 45:
				v61 = int32(_a_F_aclcheck_error_33)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 46:
				v61 = int32(_a_F_aclcheck_error_34)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 49:
				v61 = int32(_a_F_aclcheck_error_35)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 51:
				v61 = int32(_a_F_aclcheck_error_36)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			default:
				v61 = int32(_a_F_aclcheck_error_37)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
						F_errmsg(m, v61, v7+int32(16))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2795), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		case 1:
			switch l1 {
			case 0, 2, 3, 4, 5, 10, 11, 13, 27, 31, 32, 33, 43, 47, 48, 50:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v114 = m.ExcPending
				if v114 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = l1
					F_errmsg_internal(m, int32(_a_F_aclcheck_error_0), v7-int32(-64))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2928), int32(_a_F_aclcheck_error_2))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			case 1:
				v127 = int32(_a_F_aclcheck_error_38)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 6, 28, 35, 40, 44:
				v127 = int32(_a_F_aclcheck_error_39)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 7:
				v127 = int32(_a_F_aclcheck_error_40)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 8:
				v127 = int32(_a_F_aclcheck_error_41)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 9:
				v127 = int32(_a_F_aclcheck_error_42)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 12:
				v127 = int32(_a_F_aclcheck_error_43)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 14:
				v127 = int32(_a_F_aclcheck_error_44)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 15:
				v127 = int32(_a_F_aclcheck_error_45)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 16:
				v127 = int32(_a_F_aclcheck_error_46)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 17:
				v127 = int32(_a_F_aclcheck_error_47)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 18:
				v127 = int32(_a_F_aclcheck_error_48)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 19:
				v127 = int32(_a_F_aclcheck_error_49)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 20:
				v127 = int32(_a_F_aclcheck_error_50)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 21:
				v127 = int32(_a_F_aclcheck_error_51)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 22:
				v127 = int32(_a_F_aclcheck_error_52)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 23:
				v127 = int32(_a_F_aclcheck_error_53)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 24:
				v127 = int32(_a_F_aclcheck_error_54)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 25:
				v127 = int32(_a_F_aclcheck_error_55)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 26:
				v127 = int32(_a_F_aclcheck_error_56)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 29:
				v127 = int32(_a_F_aclcheck_error_57)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 30:
				v127 = int32(_a_F_aclcheck_error_58)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 34:
				v127 = int32(_a_F_aclcheck_error_59)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 36:
				v127 = int32(_a_F_aclcheck_error_60)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 37:
				v127 = int32(_a_F_aclcheck_error_61)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 38:
				v127 = int32(_a_F_aclcheck_error_62)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 39:
				v127 = int32(_a_F_aclcheck_error_63)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 41:
				v127 = int32(_a_F_aclcheck_error_64)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 42:
				v127 = int32(_a_F_aclcheck_error_65)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 45:
				v127 = int32(_a_F_aclcheck_error_66)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 46:
				v127 = int32(_a_F_aclcheck_error_67)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 49:
				v127 = int32(_a_F_aclcheck_error_68)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 51:
				v127 = int32(_a_F_aclcheck_error_69)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			default:
				v127 = int32(_a_F_aclcheck_error_37)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l2
						F_errmsg(m, v127, v7+int32(48))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2933), int32(_a_F_aclcheck_error_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v148 = m.ExcPending
			if v148 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_errmsg_internal(m, int32(_a_F_aclcheck_error_70), v7)
				mBase = m.M
				v152 = m.ExcPending
				if v152 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2937), int32(_a_F_aclcheck_error_2))
					mBase = m.M
					v157 = m.ExcPending
					if v157 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v7 + int32(80)
		return
	}
}
func F_aclcontains(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v56 int32
	_ = v56
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_check_acl(m, v9)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v17 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v27 = (v20<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L6
L5:
	;
	v27 = v17
	goto L6
L6:
	;
	if int32(0) < v16 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v33 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	return int32(0)
L10:
	;
	v42 = v27 + v9 + v33<<(uint(int32(4))%32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v31 != v43 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v56 = v33 + int32(1)
	if v56 != v16 {
		v33 = v56
		goto L10
	} else {
		goto L16
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v45 != v46 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
	if v48&v49 != v48 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	return int32(1)
L16:
	;
	goto L11
}
func F_aclitemComparator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v9) < base.Ui32(v8) {
		return int32(1)
	} else {
		v13 = int32(-1)
		if base.Ui32(v8) < base.Ui32(v9) {
			v31 = v13
			return v31
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if base.Ui32(v16) < base.Ui32(v15) {
				return int32(1)
			} else {
				if base.Ui32(v15) < base.Ui32(v16) {
					v31 = v13
				} else {
					v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					v23 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
					if base.Ui64(v23) < base.Ui64(v22) {
						v31 = int32(1)
					} else {
						if base.Ui64(v22) < base.Ui64(v23) {
							v28 = int32(-1)
						} else {
							v28 = int32(0)
						}
						v31 = v28
					}
				}
				return v31
			}
		}
	}
}
func F_aclitemsort(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	if l0 == int32(0) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v6 < int32(2) {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v9 != 0 {
				v17 = v9
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v17 = (v10<<(uint(int32(3))%32) + int32(23)) & int32(-8)
			}
			F_pg_qsort(m, v17+l0, v6, int32(16), int32(1236))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_aclupdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v136 int64
	_ = v136
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v145 int32
	_ = v145
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v334 int64
	_ = v334
	var v337 int64
	_ = v337
	var v339 int64
	_ = v339
	var v343 int64
	_ = v343
	var v344 int64
	_ = v344
	var v346 int64
	_ = v346
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v369 int64
	_ = v369
	var v371 int64
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v381 int64
	_ = v381
	var v382 int32
	_ = v382
	var v387 int64
	_ = v387
	var v397 int32
	_ = v397
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v429 int32
	_ = v429
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int64
	_ = v441
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	F_check_acl(m, l0)
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
	if l2 == int32(2) {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L103
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L100
	}
L5:
	;
	v333 = v325 + v324<<(uint(int32(4))%32)
	v334 = *(*int64)(unsafe.Add(mBase, uint32(v333)+8))
	switch l2 - int32(1) {
	case 0:
		goto L72
	case 1:
		goto L71
	case 2:
		goto L70
	default:
		v346 = v334
		goto L68
	}
L6:
	;
	v286 = v160 + int32(1)
	if v286 < int32(0) {
		goto L4
	} else {
		goto L63
	}
L7:
	;
	if v160 == v263 {
		v278 = v263
		goto L6
	} else {
		goto L62
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L58
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L55
	}
L10:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v161 != 0 {
		goto L39
	} else {
		goto L40
	}
L11:
	;
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui64(v25) < base.Ui64(int64(4294967296)) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_check_acl(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v30 == l3 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v32 < int32(0) {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v38 = v32<<(uint(int32(4))%32) + int32(24)
	v39 = F_palloc0(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = int32(1033)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+4)) = int64(1)
	v47 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v38 << (uint(v47) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v32
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v53 = int32(base.Ui32(v51) >> (uint(v47) % 32))
	if v53 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	base.MemoryCopy(m, v39, l0, v53)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v61 = v39
	goto L20
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v70 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v132 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+12)))
	v136 = F_aclmask(m, v61, v131, l3, v132<<(uint(int64(32))%64), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L36
	}
L22:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v80 = (v73<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L24
L23:
	;
	v80 = v70
	goto L24
L24:
	;
	if int32(0) < v69 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v93 = int32(0)
	goto L28
L26:
	;
	goto L27
L27:
	;
	goto L21
L28:
	;
	v102 = v61 + v80 + v93<<(uint(int32(4))%32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if v103 != v84 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L27
L30:
	;
	v115 = v93 + int32(1)
	if v115 != v69 {
		v93 = v115
		goto L28
	} else {
		goto L35
	}
L31:
	;
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v102)+8))
	if base.Ui64(v105) < base.Ui64(int64(4294967296)) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v110 = F_aclupdate(m, v61, v102, int32(2), l3, int32(1))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_pfree(m, v61)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v61 = v110
	goto L20
L35:
	;
	goto L29
L36:
	;
	v138 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui64(int64(4294967296)) <= base.Ui64(v138&(v136^int64(-1))) {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	F_pfree(m, v61)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L10
L39:
	;
	v169 = v161
	goto L41
L40:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v169 = (v162<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L41
L41:
	;
	v170 = v169 + l0
	if v160 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v173 = int32(0)
	v261 = v173
	v263 = v173
	v264 = v173
	goto L7
L43:
	;
	goto L44
L44:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v185 = int32(0)
	goto L45
L45:
	;
	v194 = v170 + v185<<(uint(int32(4))%32)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	if v176 != v195 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v278 = v160
	goto L6
L47:
	;
	v223 = v185 + int32(1)
	if v223 != v160 {
		v185 = v223
		goto L45
	} else {
		goto L54
	}
L48:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	if v197 != v198 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v203 = v160<<(uint(int32(4))%32) + int32(24)
	v204 = F_palloc0(m, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v204)+12)) = int32(1033)
	*(*int64)(unsafe.Add(mBase, uint32(v204)+4)) = int64(1)
	v212 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v203 << (uint(v212) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v204)+16)) = v160
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v218 = int32(base.Ui32(v216) >> (uint(v212) % 32))
	if v218 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	base.MemoryCopy(m, v204, l0, v218)
	goto L53
L52:
	;
	goto L53
L53:
	;
	v261 = v204
	v263 = v185
	v264 = v204 + int32(24)
	goto L7
L54:
	;
	goto L46
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v32
	F_errmsg_internal(m, int32(_a_F_aclupdate_0), v17+int32(16))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_aclupdate_1), int32(433), int32(_a_F_aclupdate_2))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_aclupdate_3), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_aclupdate_1), int32(1281), int32(_a_F_aclupdate_4))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v317 = v160
	v322 = v261
	v324 = v263
	v325 = v264
	goto L5
L63:
	;
	v292 = v286<<(uint(int32(4))%32) + int32(24)
	v293 = F_palloc0(m, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = int32(1033)
	*(*int64)(unsafe.Add(mBase, uint32(v293)+4)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v293))) = v292 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v293)+16)) = v286
	v306 = v293 + int32(24)
	v308 = v160 << (uint(int32(4)) % 32)
	if v308 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	base.MemoryCopy(m, v306, v170, v308)
	goto L67
L66:
	;
	goto L67
L67:
	;
	v310 = v308 + v306
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v310))) = v311
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v310)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v310)+4)) = v313
	v317 = v286
	v322 = v293
	v324 = v278
	v325 = v306
	goto L5
L68:
	;
	if v346 == int64(0) {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v333)+8)) = v344
	v346 = v344
	goto L68
L70:
	;
	v343 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v344 = v343
	goto L69
L71:
	;
	v339 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v344 = v334 & (v339 ^ int64(-1))
	goto L69
L72:
	;
	v337 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v344 = v337 | v334
	goto L69
L73:
	;
	v353 = (v317 + (v324 ^ int32(-1))) << (uint(int32(4)) % 32)
	if v353 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v369 = v334 & (v346 ^ int64(-1))
	v371 = int64(base.Ui64(v369) >> (uint(int64(32)) % 64))
	if v371 == int64(0) {
		v466 = v322
		goto L79
	} else {
		goto L80
	}
L76:
	;
	base.MemoryCopy(m, v333, v333+int32(16), v353)
	goto L78
L77:
	;
	goto L78
L78:
	;
	v358 = v317 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v322))) = v358<<(uint(int32(6))%32) + int32(96)
	*(*int32)(unsafe.Add(mBase, uint32(v322)+16)) = v358
	goto L75
L79:
	;
	m.G0 = v17 + int32(48)
	return v466
L80:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_check_acl(m, v322)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	if v374 == l3 {
		v466 = v322
		goto L79
	} else {
		goto L82
	}
L82:
	;
	v381 = F_aclmask(m, v322, v374, l3, v369&int64(-4294967296), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v387 = v371 & (int64(base.Ui64(v381)>>(uint(int64(32))%64)) ^ int64(-1))
	if v387 == int64(0) {
		v466 = v322
		goto L79
	} else {
		goto L84
	}
L84:
	;
	v397 = v322
	goto L85
L85:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v397)+16))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v397)+8))
	if v407 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v466 = v397
	goto L79
L87:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	v417 = (v410<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L89
L88:
	;
	v417 = v407
	goto L89
L89:
	;
	if v406 <= int32(0) {
		v466 = v397
		goto L79
	} else {
		goto L90
	}
L90:
	;
	v429 = int32(0)
	goto L91
L91:
	;
	v438 = v397 + v417 + v429<<(uint(int32(4))%32)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	if v439 != v374 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L86
L93:
	;
	v459 = v429 + int32(1)
	if v459 != v406 {
		v429 = v459
		goto L91
	} else {
		goto L99
	}
L94:
	;
	v441 = *(*int64)(unsafe.Add(mBase, uint32(v438)+8))
	if v441&v387 == int64(0) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	if l4 == int32(0) {
		goto L3
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v374
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v438)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v387 * int64(4294967297)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v448
	v454 = F_aclupdate(m, v397, v17+int32(32), int32(2), l3, l4)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_pfree(m, v397)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v397 = v454
	goto L85
L99:
	;
	goto L92
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v286
	F_errmsg_internal(m, int32(_a_F_aclupdate_0), v17)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_aclupdate_1), int32(433), int32(_a_F_aclupdate_2))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	F_errcode(m, int32(16909442))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errmsg(m, int32(_a_F_aclupdate_5), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errhint(m, int32(_a_F_aclupdate_6), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_aclupdate_1), int32(1343), int32(_a_F_aclupdate_7))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_action_abort(m *base.Module, l0 int32) {
	F_abort(m)
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_add_nulling_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13866(m, l0, l1, l2, int32(1052))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_add_paths_to_append_rel(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 float64
	_ = v116
	var v124 int32
	_ = v124
	var v131 float64
	_ = v131
	var v137 float64
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v271 int32
	_ = v271
	var v279 float64
	_ = v279
	var v280 float64
	_ = v280
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v553 int32
	_ = v553
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v966 int32
	_ = v966
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 float64
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1066 float64
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1107 int32
	_ = v1107
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1189 int32
	_ = v1189
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1212 int32
	_ = v1212
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1325 int32
	_ = v1325
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1348 int32
	_ = v1348
	var v1355 int32
	_ = v1355
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1418 int32
	_ = v1418
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1454 int32
	_ = v1454
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1474 int32
	_ = v1474
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1510 int32
	_ = v1510
	var v1518 int32
	_ = v1518
	var v1527 int32
	_ = v1527
	var v1531 int32
	_ = v1531
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1571 int32
	_ = v1571
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1604 int32
	_ = v1604
	var v1607 int32
	_ = v1607
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1620 int32
	_ = v1620
	var v1628 int32
	_ = v1628
	var v1632 int32
	_ = v1632
	var v1636 int32
	_ = v1636
	var v1639 int32
	_ = v1639
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1671 int32
	_ = v1671
	var v1675 int32
	_ = v1675
	var v1683 int32
	_ = v1683
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1703 int32
	_ = v1703
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1719 int32
	_ = v1719
	var v1723 int32
	_ = v1723
	var v1727 int32
	_ = v1727
	var v1730 int32
	_ = v1730
	var v1738 int32
	_ = v1738
	var v1746 int32
	_ = v1746
	var v1750 int32
	_ = v1750
	var v1752 int32
	_ = v1752
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1763 int32
	_ = v1763
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1801 int32
	_ = v1801
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1818 int32
	_ = v1818
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1854 int32
	_ = v1854
	var v1864 int32
	_ = v1864
	var v1867 int32
	_ = v1867
	var v1870 int32
	_ = v1870
	var v1874 int32
	_ = v1874
	var v1878 int32
	_ = v1878
	var v1881 int32
	_ = v1881
	var v1889 int32
	_ = v1889
	var v1897 int32
	_ = v1897
	var v1901 int32
	_ = v1901
	var v1903 int32
	_ = v1903
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1914 int32
	_ = v1914
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1952 int32
	_ = v1952
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1969 int32
	_ = v1969
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1996 float64
	_ = v1996
	var v1999 int32
	_ = v1999
	var v2002 float64
	_ = v2002
	var v2004 float64
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2016 int32
	_ = v2016
	var v2024 int32
	_ = v2024
	var v2027 int32
	_ = v2027
	var v2030 int32
	_ = v2030
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2038 int32
	_ = v2038
	var v2044 int32
	_ = v2044
	var v2052 int32
	_ = v2052
	var v2056 int32
	_ = v2056
	var v2058 int32
	_ = v2058
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2069 int32
	_ = v2069
	var v2076 int32
	_ = v2076
	var v2078 int32
	_ = v2078
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2104 int32
	_ = v2104
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2119 int32
	_ = v2119
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2139 int32
	_ = v2139
	var v2141 int32
	_ = v2141
	var v2144 int32
	_ = v2144
	var v2147 int32
	_ = v2147
	var v2150 int32
	_ = v2150
	var v2153 int32
	_ = v2153
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2161 int32
	_ = v2161
	var v2164 int32
	_ = v2164
	var v2167 int32
	_ = v2167
	var v2170 int32
	_ = v2170
	var v2173 int32
	_ = v2173
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2187 int32
	_ = v2187
	var v2190 int32
	_ = v2190
	var v2193 int32
	_ = v2193
	var v2196 int32
	_ = v2196
	var v2199 int32
	_ = v2199
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2209 int32
	_ = v2209
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2225 int32
	_ = v2225
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2240 int32
	_ = v2240
	var v2243 int32
	_ = v2243
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2267 int32
	_ = v2267
	var v2269 int32
	_ = v2269
	var v2281 int32
	_ = v2281
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2287 int32
	_ = v2287
	var v2309 int32
	_ = v2309
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2315 int32
	_ = v2315
	var v2327 int32
	_ = v2327
	var v2332 int32
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2335 int32
	_ = v2335
	var v2338 int32
	_ = v2338
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2346 int32
	_ = v2346
	var v2349 int32
	_ = v2349
	var v2354 int32
	_ = v2354
	var v2355 int32
	_ = v2355
	var v2364 int32
	_ = v2364
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2370 int32
	_ = v2370
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2385 int32
	_ = v2385
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2391 int32
	_ = v2391
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2422 int32
	_ = v2422
	var v2424 int32
	_ = v2424
	var v2452 int32
	_ = v2452
	var v2453 int32
	_ = v2453
	var v2483 int32
	_ = v2483
	var v2501 int32
	_ = v2501
	var v2513 int32
	_ = v2513
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2533 int32
	_ = v2533
	var v2536 int32
	_ = v2536
	var v2551 int32
	_ = v2551
	var v2555 int32
	_ = v2555
	var v2556 int32
	_ = v2556
	var v2559 int32
	_ = v2559
	var v2574 int32
	_ = v2574
	var v2584 int32
	_ = v2584
	var v2587 int32
	_ = v2587
	var v2590 int32
	_ = v2590
	var v2594 int32
	_ = v2594
	var v2598 int32
	_ = v2598
	var v2601 int32
	_ = v2601
	var v2628 int32
	_ = v2628
	var v2659 int32
	_ = v2659
	var v2660 int32
	_ = v2660
	var v2662 int32
	_ = v2662
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2672 int32
	_ = v2672
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2689 int32
	_ = v2689
	var v2708 int32
	_ = v2708
	var v2709 int32
	_ = v2709
	var v2710 int32
	_ = v2710
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2727 int32
	_ = v2727
	var v2728 int32
	_ = v2728
	var v2730 int32
	_ = v2730
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2739 int32
	_ = v2739
	var v2747 int32
	_ = v2747
	var v2749 int32
	_ = v2749
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2755 int32
	_ = v2755
	var v2759 int32
	_ = v2759
	var v2766 int32
	_ = v2766
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2776 int32
	_ = v2776
	var v2792 int32
	_ = v2792
	var v2799 int32
	_ = v2799
	var v2803 int32
	_ = v2803
	var v2804 int32
	_ = v2804
	var v2805 int32
	_ = v2805
	var v2807 int32
	_ = v2807
	var v2808 int32
	_ = v2808
	var v2817 int32
	_ = v2817
	var v2818 int32
	_ = v2818
	var v2820 int32
	_ = v2820
	var v2823 int32
	_ = v2823
	var v2824 int32
	_ = v2824
	var v2829 int32
	_ = v2829
	var v2836 int32
	_ = v2836
	var v2838 int32
	_ = v2838
	var v2840 int32
	_ = v2840
	var v2843 int32
	_ = v2843
	var v2845 int32
	_ = v2845
	var v2847 int32
	_ = v2847
	var v2854 int32
	_ = v2854
	var v2861 int32
	_ = v2861
	var v2871 int32
	_ = v2871
	var v2872 int32
	_ = v2872
	var v2877 int32
	_ = v2877
	var v2893 int32
	_ = v2893
	var v2894 float64
	_ = v2894
	var v2895 float64
	_ = v2895
	var v2899 float64
	_ = v2899
	var v2900 float64
	_ = v2900
	var v2908 int32
	_ = v2908
	var v2914 int32
	_ = v2914
	var v2917 int32
	_ = v2917
	var v2918 int32
	_ = v2918
	var v2920 int32
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2935 int32
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2938 int32
	_ = v2938
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2947 int32
	_ = v2947
	var v2955 int32
	_ = v2955
	var v2957 int32
	_ = v2957
	var v2959 int32
	_ = v2959
	var v2960 int32
	_ = v2960
	var v2963 int32
	_ = v2963
	var v2967 int32
	_ = v2967
	var v2973 int32
	_ = v2973
	var v2974 int32
	_ = v2974
	var v2975 int32
	_ = v2975
	var v2976 int32
	_ = v2976
	var v2986 int32
	_ = v2986
	var v2987 int32
	_ = v2987
	var v2992 int32
	_ = v2992
	var v3008 int32
	_ = v3008
	var v3009 float64
	_ = v3009
	var v3010 float64
	_ = v3010
	var v3014 float64
	_ = v3014
	var v3015 float64
	_ = v3015
	var v3023 int32
	_ = v3023
	var v3029 int32
	_ = v3029
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3036 int32
	_ = v3036
	var v3037 int32
	_ = v3037
	var v3044 int32
	_ = v3044
	var v3067 int32
	_ = v3067
	var v3070 int32
	_ = v3070
	var v3073 int32
	_ = v3073
	var v3074 int32
	_ = v3074
	var v3075 int32
	_ = v3075
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3078 int32
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3080 int32
	_ = v3080
	var v3081 int32
	_ = v3081
	var v3082 int32
	_ = v3082
	var v3084 int32
	_ = v3084
	var v3085 int32
	_ = v3085
	var v3095 int32
	_ = v3095
	var v3113 int32
	_ = v3113
	var v3118 int32
	_ = v3118
	var v3119 int32
	_ = v3119
	var v3121 int32
	_ = v3121
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3180 int32
	_ = v3180
	var v3183 int32
	_ = v3183
	var v3184 int32
	_ = v3184
	var v3185 int32
	_ = v3185
	var v3188 int32
	_ = v3188
	var v3195 int32
	_ = v3195
	var v3218 int32
	_ = v3218
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3232 int32
	_ = v3232
	var v3234 int32
	_ = v3234
	var v3236 int32
	_ = v3236
	var v3237 int32
	_ = v3237
	var v3239 int32
	_ = v3239
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	v4 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(32)
	m.G0 = v29
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v4
	v35 = int32(1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_add_paths_to_append_rel[0])))
	if v37 == v35 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	v41 = v40
	goto L3
L2:
	;
	v41 = v4
	goto L3
L3:
	;
	if l2 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v822&int32(1) != 0 {
		goto L208
	} else {
		goto L209
	}
L5:
	;
	v797 = int32(0)
	v803 = F_create_append_path(m, l0, l1, v788, v797, v797, v797, v797, v797, float64(-1))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L29
	} else {
		goto L206
	}
L6:
	;
	v779 = v4
	v782 = v41
	v783 = v4
	v784 = int32(1)
	v785 = v35
	v788 = v4
	v790 = v4
	v791 = v4
	goto L5
L7:
	;
	goto L8
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v45 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v758&int32(1) != 0 {
		v779 = v750
		v782 = v753
		v783 = v754
		v784 = v755
		v785 = v756
		v788 = v759
		v790 = v761
		v791 = v762
		goto L5
	} else {
		goto L205
	}
L10:
	;
	v48 = int32(1)
	v750 = v4
	v753 = v41
	v754 = v4
	v755 = v48
	v756 = v35
	v758 = v48
	v759 = v4
	v761 = v4
	v762 = v4
	goto L9
L11:
	;
	goto L12
L12:
	;
	v50 = int32(1)
	v59 = v4
	v60 = v4
	v63 = v41
	v64 = v4
	v65 = v50
	v66 = v35
	v68 = v50
	v69 = v4
	v71 = v4
	v72 = v4
	goto L13
L13:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v59<<(uint(int32(2))%32))))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+32))
	if v83 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v750 = v219
	v753 = v357
	v754 = v193
	v755 = v220
	v756 = v194
	v758 = v106
	v759 = v107
	v761 = v731
	v762 = v732
	goto L9
L15:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v108 != int32(1) {
		goto L34
	} else {
		goto L35
	}
L16:
	;
	v106 = int32(0)
	v107 = v69
	goto L15
L17:
	;
	goto L18
L18:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)+48))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	if v88 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v106 = int32(0)
	v107 = v69
	goto L15
L20:
	;
	goto L21
L21:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	switch v90 - int32(290) {
	case 0:
		goto L24
	case 1:
		goto L23
	default:
		goto L22
	}
L22:
	;
	v103 = F_lappend(m, v69, v87)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L29
	} else {
		goto L32
	}
L23:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v87)+72))
	v101 = F_list_concat(m, v69, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L29
	} else {
		goto L31
	}
L24:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+20)))
	if v93 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+76))
	if v96 != 0 {
		goto L22
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v87)+72))
	v98 = F_list_concat(m, v69, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	return
L30:
	;
	v106 = v68
	v107 = v98
	goto L15
L31:
	;
	v106 = v68
	v107 = v101
	goto L15
L32:
	;
	v106 = v68
	v107 = v103
	goto L15
L33:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	if v196 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L34:
	;
	v193 = v64
	v194 = int32(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v82)+44))
	if v112 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v193 = v64
	v194 = int32(0)
	goto L33
L38:
	;
	goto L39
L39:
	;
	v116 = *(*float64)(unsafe.Add(mBase, uint32(l0)+296))
	if base.F64_gt(v116, float64(0)) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v82)+48))
	if base.F64_le(v116, float64(0)) != 0 {
		v171 = v124
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v176 = v112
	goto L42
L42:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	switch v177 - int32(290) {
	case 0:
		goto L62
	case 1:
		goto L61
	default:
		goto L60
	}
L43:
	;
	v176 = v171
	goto L42
L44:
	;
	goto L43
L45:
	;
	if base.F64_ge(v116, float64(1)) == int32(0) {
		v137 = v116
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v82)+32))
	if v139 == int32(0) {
		v171 = v124
		goto L44
	} else {
		goto L49
	}
L47:
	;
	v131 = *(*float64)(unsafe.Add(mBase, uint32(v124)+32))
	if base.F64_gt(v131, float64(0)) == int32(0) {
		v137 = v116
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v137 = base.F64_div(v116, v131)
	goto L46
L49:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v142 <= int32(0) {
		v171 = v124
		goto L44
	} else {
		goto L50
	}
L50:
	;
	v147 = v124
	v150 = int32(0)
	goto L51
L51:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152+v150<<(uint(int32(2))%32))))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	if v157 != 0 {
		v164 = v147
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v171 = v164
	goto L44
L53:
	;
	v166 = v150 + int32(1)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v166 < v167 {
		v147 = v164
		v150 = v166
		goto L51
	} else {
		goto L59
	}
L54:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v82)+48))
	if v156 == v158 {
		v164 = v147
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v160 = F_compare_fractional_path_costs(m, v147, v156, v137)
	mBase = m.M
	if v160 <= int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v163 = v147
	goto L58
L57:
	;
	v163 = v156
	goto L58
L58:
	;
	v164 = v163
	goto L53
L59:
	;
	goto L52
L60:
	;
	v190 = F_lappend(m, v64, v176)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L29
	} else {
		goto L69
	}
L61:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v176)+72))
	v188 = F_list_concat(m, v64, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L29
	} else {
		goto L68
	}
L62:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+20)))
	if v180 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v176)+76))
	if v183 != 0 {
		goto L60
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v176)+72))
	v185 = F_list_concat(m, v64, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L29
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	v193 = v185
	v194 = v66
	goto L33
L68:
	;
	v193 = v188
	v194 = v66
	goto L33
L69:
	;
	v193 = v190
	v194 = v66
	goto L33
L70:
	;
	if v63&int32(1) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L71:
	;
	v199 = int32(0)
	v218 = v199
	v219 = v60
	v220 = v199
	goto L70
L72:
	;
	goto L73
L73:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v196)+12))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	switch v203 - int32(290) {
	case 0:
		goto L76
	case 1:
		goto L75
	default:
		goto L74
	}
L74:
	;
	v216 = F_lappend(m, v60, v202)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L29
	} else {
		goto L83
	}
L75:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v202)+72))
	v214 = F_list_concat(m, v60, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L29
	} else {
		goto L82
	}
L76:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+20)))
	if v206 == int32(1) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v202)+76))
	if v209 != 0 {
		goto L74
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v202)+72))
	v211 = F_list_concat(m, v60, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L29
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	v218 = v202
	v219 = v211
	v220 = v65
	goto L70
L82:
	;
	v218 = v202
	v219 = v214
	v220 = v65
	goto L70
L83:
	;
	v218 = v202
	v219 = v216
	v220 = v65
	goto L70
L84:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v82)+32))
	if v358 == int32(0) {
		v731 = v71
		v732 = v72
		goto L140
	} else {
		goto L141
	}
L85:
	;
	v357 = int32(0)
	goto L84
L86:
	;
	goto L87
L87:
	;
	v226 = int32(0)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v82)+32))
	if v227 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	if v218|v271 == int32(0) {
		v357 = v226
		goto L84
	} else {
		goto L105
	}
L89:
	;
	goto L88
L90:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v232 <= int32(0) {
		v271 = v226
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v271 = int32(0)
	goto L89
L93:
	;
	v235 = int32(0)
	if v235 < v232 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v238 = v232
	goto L96
L95:
	;
	v238 = v235
	goto L96
L96:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	v241 = int32(0)
	goto L97
L97:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v239+v241<<(uint(int32(2))%32))))
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+21)))
	if v250 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L92
L99:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v249)+16))
	if v253 == int32(0) {
		v271 = v249
		goto L89
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v261 = v241 + int32(1)
	if v261 != v238 {
		v241 = v261
		goto L97
	} else {
		goto L104
	}
L102:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	if v256 == int32(0) {
		v271 = v249
		goto L89
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	goto L98
L105:
	;
	if v271 != 0 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v357 = int32(1)
	goto L84
L107:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	switch v329 - int32(290) {
	case 0:
		goto L132
	case 1:
		goto L131
	default:
		goto L130
	}
L108:
	;
	if v218 == int32(0) {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v285 = v29 + int32(24)
	v287 = v29 + int32(20)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	switch v288 - int32(290) {
	case 0:
		goto L116
	case 1:
		goto L115
	default:
		goto L114
	}
L111:
	;
	v279 = *(*float64)(unsafe.Add(mBase, uint32(v218)+56))
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v271)+56))
	if base.F64_lt(v279, v280) == int32(0) {
		goto L107
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	goto L106
L114:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v325 = F_lappend(m, v324, v218)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L29
	} else {
		goto L129
	}
L115:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v218)+72))
	v320 = F_list_concat(m, v318, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L29
	} else {
		goto L128
	}
L116:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+20)))
	if v291 == int32(1) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	if v287 == int32(0) {
		goto L114
	} else {
		goto L123
	}
L118:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v218)+76))
	if v294 != 0 {
		goto L117
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v218)+72))
	v298 = F_list_concat(m, v296, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L29
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v298
	goto L113
L123:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v218)+72))
	v305 = F_list_copy_tail(m, v304, v294)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L29
	} else {
		goto L124
	}
L124:
	;
	v307 = F_list_concat(m, v303, v305)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L29
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v307
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v218)+72))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v218)+76))
	v312 = F_list_copy_head(m, v310, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L29
	} else {
		goto L126
	}
L126:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	v315 = F_list_concat(m, v314, v312)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L29
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = v315
	goto L113
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v320
	goto L113
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v325
	goto L113
L130:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v347 = F_lappend(m, v346, v271)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L29
	} else {
		goto L139
	}
L131:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v271)+72))
	v343 = F_list_concat(m, v341, v342)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L29
	} else {
		goto L138
	}
L132:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+20)))
	if v332 == int32(1) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v271)+76))
	if v335 != 0 {
		goto L130
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v271)+72))
	v338 = F_list_concat(m, v336, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L29
	} else {
		goto L137
	}
L136:
	;
	goto L135
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v338
	goto L106
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v343
	goto L106
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v347
	goto L106
L140:
	;
	v739 = v59 + int32(1)
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v739 < v740 {
		v59 = v739
		v60 = v219
		v63 = v357
		v64 = v193
		v65 = v220
		v66 = v194
		v68 = v106
		v69 = v107
		v71 = v731
		v72 = v732
		goto L13
	} else {
		goto L204
	}
L141:
	;
	v361 = int32(0)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	if v362 <= v361 {
		v731 = v71
		v732 = v72
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v383 = v361
	v384 = v71
	v385 = v72
	goto L143
L143:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v358)+12))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v391+v383<<(uint(int32(2))%32))))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+64))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v395)+16))
	if v398 != 0 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v731 = v553
	v732 = v702
	goto L140
L145:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)+4))
	v400 = v399
	goto L147
L146:
	;
	v400 = int32(0)
	goto L147
L147:
	;
	if v396 == int32(0) {
		v553 = v384
		goto L148
	} else {
		goto L149
	}
L148:
	;
	if v400 == int32(0) {
		v702 = v385
		goto L182
	} else {
		goto L183
	}
L149:
	;
	if v384 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v532 = F_lappend(m, v384, v396)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L29
	} else {
		goto L181
	}
L151:
	;
	v405 = int32(0)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v384)+4))
	if v406 <= v405 {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v412 = v405
	goto L153
L153:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v384)+12))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v435+v412<<(uint(int32(2))%32))))
	if v439 == v396 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	goto L150
L155:
	;
	if v499 == int32(0) {
		v553 = v384
		goto L148
	} else {
		goto L179
	}
L156:
	;
	v499 = int32(0)
	goto L155
L157:
	;
	goto L158
L158:
	;
	v448 = int32(0)
	goto L161
L159:
	;
	if v488 != 0 {
		goto L176
	} else {
		goto L177
	}
L160:
	;
	v483 = int32(0)
	if v470 != 0 {
		goto L173
	} else {
		goto L174
	}
L161:
	;
	v452 = int32(0)
	if v439 == v452 {
		v462 = v452
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v499 = int32(3)
	goto L155
L163:
	;
	if v396 != 0 {
		goto L167
	} else {
		goto L168
	}
L164:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	if v456 <= v448 {
		v462 = int32(0)
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v439)+12))
	v462 = v458 + v448<<(uint(int32(2))%32)
	goto L163
L166:
	;
	v468 = int32(0)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v396)+12))
	if base.B2i32(v462 == v468)|base.B2i32(v470 == v468) != 0 {
		goto L160
	} else {
		goto L171
	}
L167:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	if v448 < v463 {
		goto L166
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v465 = int32(0)
	v488 = base.B2i32(v462 == v465)
	v490 = v465
	goto L159
L170:
	;
	goto L169
L171:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v470+v448<<(uint(int32(2))%32))))
	if v478 == v480 {
		v448 = v448 + int32(1)
		goto L161
	} else {
		goto L172
	}
L172:
	;
	goto L162
L173:
	;
	v487 = int32(2)
	goto L175
L174:
	;
	v487 = v483
	goto L175
L175:
	;
	v488 = base.B2i32(v462 == v483)
	v490 = v487
	goto L159
L176:
	;
	v492 = v490
	goto L178
L177:
	;
	v492 = int32(1)
	goto L178
L178:
	;
	v499 = v492
	goto L155
L179:
	;
	v503 = v412 + int32(1)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v384)+4))
	if v503 < v504 {
		v412 = v503
		goto L153
	} else {
		goto L180
	}
L180:
	;
	goto L154
L181:
	;
	v553 = v532
	goto L148
L182:
	;
	v709 = v383 + int32(1)
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	if v709 < v710 {
		v383 = v709
		v384 = v553
		v385 = v702
		goto L143
	} else {
		goto L203
	}
L183:
	;
	if v385 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v680 = F_lappend(m, v385, v400)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L29
	} else {
		goto L202
	}
L185:
	;
	v564 = int32(0)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v385)+4))
	if v565 <= v564 {
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v571 = v564
	goto L187
L187:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v385)+12))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v594+v571<<(uint(int32(2))%32))))
	v599 = int32(0)
	if base.B2i32(v598 == v599)|base.B2i32(v400 == v599) != 0 {
		v645 = base.B2i32(v598|v400 == v599)
		goto L190
	} else {
		goto L191
	}
L188:
	;
	goto L184
L189:
	;
	if v645 != 0 {
		v702 = v385
		goto L182
	} else {
		goto L200
	}
L190:
	;
	goto L189
L191:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v598)+4))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	if v613 != v614 {
		v645 = int32(0)
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v616 = int32(1)
	if v613 <= v616 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v619 = v616
	goto L195
L194:
	;
	v619 = v613
	goto L195
L195:
	;
	v620 = int32(8)
	v625 = int32(0)
	goto L196
L196:
	;
	v633 = v625 << (uint(int32(2)) % 32)
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v598+v620+v633)))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v400+v620+v633)))
	v638 = base.B2i32(v635 == v637)
	if v635 != v637 {
		v645 = v638
		goto L190
	} else {
		goto L198
	}
L197:
	;
	v645 = v638
	goto L190
L198:
	;
	v641 = v625 + int32(1)
	if v641 != v619 {
		v625 = v641
		goto L196
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	v651 = v571 + int32(1)
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v385)+4))
	if v651 < v652 {
		v571 = v651
		goto L187
	} else {
		goto L201
	}
L201:
	;
	goto L188
L202:
	;
	v702 = v680
	goto L182
L203:
	;
	goto L144
L204:
	;
	goto L14
L205:
	;
	v816 = v750
	v819 = v753
	v820 = v754
	v821 = v755
	v822 = v756
	v825 = int32(0)
	v827 = v761
	v828 = v762
	goto L4
L206:
	;
	F_add_path(m, l1, v803)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L29
	} else {
		goto L207
	}
L207:
	;
	v816 = v779
	v819 = v782
	v820 = v783
	v821 = v784
	v822 = v785
	v825 = int32(1)
	v827 = v790
	v828 = v791
	goto L4
L208:
	;
	v836 = int32(0)
	v842 = F_create_append_path(m, l0, l1, v820, v836, v836, v836, v836, v836, float64(-1))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L29
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v849 = int32(0)
	if base.B2i32(v821&int32(1) == v849)|base.B2i32(v816 == v849) == v849 {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	F_add_path(m, l1, v842)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L29
	} else {
		goto L212
	}
L212:
	;
	goto L210
L213:
	;
	v856 = int32(0)
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v816)+4))
	if v857 <= v856 {
		v989 = v856
		goto L216
	} else {
		goto L217
	}
L214:
	;
	v1066 = float64(-1)
	goto L215
L215:
	;
	v1067 = int32(0)
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	if base.B2i32(v1068 != v1067)&v819 != 0 {
		goto L257
	} else {
		goto L258
	}
L216:
	;
	v1012 = int32(0)
	v1016 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_add_paths_to_append_rel[0])))
	if v1016 == int32(1) {
		goto L243
	} else {
		goto L244
	}
L217:
	;
	v861 = v857 & int32(3)
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v816)+12))
	v863 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v857) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v873 = v856
	v875 = v863
	v877 = int32(0)
	goto L221
L219:
	;
	v925 = v856
	v927 = v863
	goto L220
L220:
	;
	v951 = v925
	v953 = v927
	v966 = v863
	goto L237
L221:
	;
	v898 = v862 + v875<<(uint(int32(2))%32)
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v898)))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v899)+24))
	if v900 < v873 {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	if v861 == int32(0) {
		v989 = v914
		goto L216
	} else {
		goto L236
	}
L223:
	;
	v902 = v873
	goto L225
L224:
	;
	v902 = v900
	goto L225
L225:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v898)+4))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v903)+24))
	if v904 < v902 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v906 = v902
	goto L228
L227:
	;
	v906 = v904
	goto L228
L228:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v898)+8))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v907)+24))
	if v908 < v906 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v910 = v906
	goto L231
L230:
	;
	v910 = v908
	goto L231
L231:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v898)+12))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v911)+24))
	if v912 < v910 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v914 = v910
	goto L234
L233:
	;
	v914 = v912
	goto L234
L234:
	;
	v915 = int32(4)
	v916 = v875 + v915
	v918 = v877 + v915
	if v918 != v857&int32(2147483644) {
		v873 = v914
		v875 = v916
		v877 = v918
		goto L221
	} else {
		goto L235
	}
L235:
	;
	goto L222
L236:
	;
	v925 = v914
	v927 = v916
	goto L220
L237:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v862+v953<<(uint(int32(2))%32))))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v977)+24))
	if v978 < v951 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v989 = v980
	goto L216
L239:
	;
	v980 = v951
	goto L241
L240:
	;
	v980 = v978
	goto L241
L241:
	;
	v981 = int32(1)
	v984 = v966 + v981
	if v984 != v861 {
		v951 = v980
		v953 = v953 + v981
		v966 = v984
		goto L237
	} else {
		goto L242
	}
L242:
	;
	goto L238
L243:
	;
	if l2 != 0 {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	v1034 = v989
	goto L245
L245:
	;
	v1036 = F_create_append_path(m, l0, l1, v1012, v816, v1012, v1012, v1034, v1016, float64(-1))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L29
	} else {
		goto L255
	}
L246:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1024 = int32(32) - base.I32_clz(v1021)
	goto L248
L247:
	;
	v1024 = int32(0)
	goto L248
L248:
	;
	if v1024 < v989 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1026 = v989
	goto L251
L250:
	;
	v1026 = v1024
	goto L251
L251:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, _c_F_add_paths_to_append_rel[1]))
	if v1026 < v1028 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1030 = v1026
	goto L254
L253:
	;
	v1030 = v1028
	goto L254
L254:
	;
	v1034 = v1030
	goto L245
L255:
	;
	v1038 = *(*float64)(unsafe.Add(mBase, uint32(v1036)+32))
	F_add_partial_path(m, l1, v1036)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L29
	} else {
		goto L256
	}
L256:
	;
	v1066 = v1038
	goto L215
L257:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	if v1072 == int32(0) {
		v1212 = v1067
		goto L260
	} else {
		goto L261
	}
L258:
	;
	goto L259
L259:
	;
	if v825 == int32(0) {
		goto L302
	} else {
		goto L303
	}
L260:
	;
	if l2 != 0 {
		goto L291
	} else {
		goto L292
	}
L261:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+4))
	if v1075 <= int32(0) {
		v1212 = v1067
		goto L260
	} else {
		goto L262
	}
L262:
	;
	v1078 = int32(0)
	if v1078 < v1075 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1081 = v1075
	goto L265
L264:
	;
	v1081 = v1078
	goto L265
L265:
	;
	v1083 = v1081 & int32(3)
	v1084 = int32(0)
	if int32(4) <= v1075 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+12))
	v1095 = v1067
	v1097 = v1084
	v1107 = int32(0)
	goto L269
L267:
	;
	v1147 = v1067
	v1149 = v1084
	goto L268
L268:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+12))
	v1174 = v1147
	v1176 = v1149
	v1189 = v1084
	goto L285
L269:
	;
	v1120 = v1090 + v1097<<(uint(int32(2))%32)
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1120)))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+24))
	if v1122 < v1095 {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	if v1083 == int32(0) {
		v1212 = v1136
		goto L260
	} else {
		goto L284
	}
L271:
	;
	v1124 = v1095
	goto L273
L272:
	;
	v1124 = v1122
	goto L273
L273:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+4))
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+24))
	if v1126 < v1124 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v1128 = v1124
	goto L276
L275:
	;
	v1128 = v1126
	goto L276
L276:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+8))
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+24))
	if v1130 < v1128 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1132 = v1128
	goto L279
L278:
	;
	v1132 = v1130
	goto L279
L279:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+12))
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+24))
	if v1134 < v1132 {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1136 = v1132
	goto L282
L281:
	;
	v1136 = v1134
	goto L282
L282:
	;
	v1137 = int32(4)
	v1138 = v1097 + v1137
	v1140 = v1107 + v1137
	if v1140 != v1081&int32(2147483644) {
		v1095 = v1136
		v1097 = v1138
		v1107 = v1140
		goto L269
	} else {
		goto L283
	}
L283:
	;
	goto L270
L284:
	;
	v1147 = v1136
	v1149 = v1138
	goto L268
L285:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1170+v1176<<(uint(int32(2))%32))))
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1200)+24))
	if v1201 < v1174 {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	v1212 = v1203
	goto L260
L287:
	;
	v1203 = v1174
	goto L289
L288:
	;
	v1203 = v1201
	goto L289
L289:
	;
	v1204 = int32(1)
	v1207 = v1189 + v1204
	if v1207 != v1083 {
		v1174 = v1203
		v1176 = v1176 + v1204
		v1189 = v1207
		goto L285
	} else {
		goto L290
	}
L290:
	;
	goto L286
L291:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1240 = int32(32) - base.I32_clz(v1237)
	goto L293
L292:
	;
	v1240 = int32(0)
	goto L293
L293:
	;
	v1241 = int32(0)
	if v1240 < v1212 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1244 = v1212
	goto L296
L295:
	;
	v1244 = v1240
	goto L296
L296:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, _c_F_add_paths_to_append_rel[1]))
	if v1244 < v1246 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1248 = v1244
	goto L299
L298:
	;
	v1248 = v1246
	goto L299
L299:
	;
	v1250 = F_create_append_path(m, l0, l1, v1068, v1072, v1241, v1241, v1248, int32(1), v1066)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L29
	} else {
		goto L300
	}
L300:
	;
	F_add_partial_path(m, l1, v1250)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L29
	} else {
		goto L301
	}
L301:
	;
	goto L259
L302:
	;
	if v828 == int32(0) {
		goto L658
	} else {
		goto L659
	}
L303:
	;
	v1282 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+31)) = uint8(v1282)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+30)) = uint8(v1282)
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	if v1286 == int32(0) {
		v1371 = v4
		v1372 = v4
		goto L304
	} else {
		goto L305
	}
L304:
	;
	if v827 == int32(0) {
		goto L302
	} else {
		goto L331
	}
L305:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v1289 {
	case 0, 2:
		goto L306
	default:
		v1371 = v4
		v1372 = v4
		goto L304
	}
L306:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(l1)+240))
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	switch v1292 - int32(108) {
	case 0:
		goto L309
	default:
		goto L308
	case 6:
		goto L310
	}
L307:
	;
	if v1355 == int32(0) {
		v1371 = v4
		v1372 = v4
		goto L304
	} else {
		goto L328
	}
L308:
	;
	v1355 = int32(0)
	goto L307
L309:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+16))
	v1303 = int32(0)
	if base.B2i32(v1290 == v1303)|base.B2i32(v1302 == v1303) != 0 {
		v1348 = v1303
		goto L315
	} else {
		goto L316
	}
L310:
	;
	v1295 = int32(1)
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+32))
	if v1296 == int32(-1) {
		v1355 = v1295
		goto L307
	} else {
		goto L311
	}
L311:
	;
	v1299 = F_bms_is_member(m, v1296, v1290)
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L29
	} else {
		goto L312
	}
L312:
	;
	if v1299 != 0 {
		goto L308
	} else {
		goto L313
	}
L313:
	;
	v1355 = v1295
	goto L307
L314:
	;
	if v1348 == int32(0) {
		v1355 = int32(1)
		goto L307
	} else {
		goto L327
	}
L315:
	;
	goto L314
L316:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v1290)+4))
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1302)+4))
	if v1313 < v1314 {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v1316 = v1313
	goto L319
L318:
	;
	v1316 = v1314
	goto L319
L319:
	;
	if v1316 <= int32(1) {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1319 = int32(1)
	goto L322
L321:
	;
	v1319 = v1316
	goto L322
L322:
	;
	v1320 = int32(8)
	v1325 = int32(0)
	goto L323
L323:
	;
	v1332 = v1325 << (uint(int32(2)) % 32)
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1302+v1320+v1332)))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1290+v1320+v1332)))
	v1337 = v1334 & v1336
	v1339 = base.B2i32(v1337 != int32(0))
	if v1337 != 0 {
		v1348 = v1339
		goto L315
	} else {
		goto L325
	}
L324:
	;
	v1348 = v1339
	goto L315
L325:
	;
	v1341 = v1325 + int32(1)
	if v1341 != v1319 {
		v1325 = v1341
		goto L323
	} else {
		goto L326
	}
L326:
	;
	goto L324
L327:
	;
	goto L308
L328:
	;
	v1361 = F_build_partition_pathkeys(m, l0, l1, int32(1), v29+int32(31))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L29
	} else {
		goto L329
	}
L329:
	;
	v1366 = F_build_partition_pathkeys(m, l0, l1, int32(-1), v29+int32(30))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L29
	} else {
		goto L330
	}
L330:
	;
	v1371 = v1361
	v1372 = v1366
	goto L304
L331:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v827)+4))
	if v1375 <= int32(0) {
		goto L302
	} else {
		goto L332
	}
L332:
	;
	v1401 = v4
	goto L333
L333:
	;
	v1404 = int32(1)
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v827)+12))
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1405+v1401<<(uint(int32(2))%32))))
	if v1409 == v1371 {
		goto L345
	} else {
		goto L346
	}
L334:
	;
	goto L302
L335:
	;
	v2452 = v1401 + int32(1)
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(v827)+4))
	if v2452 < v2453 {
		v1401 = v2452
		goto L333
	} else {
		goto L657
	}
L336:
	;
	F_add_path(m, l1, v2422)
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		goto L29
	} else {
		goto L656
	}
L337:
	;
	v2382 = F_create_merge_append_path(m, l0, l1, v2364, v1409)
	mBase = m.M
	v2383 = m.ExcPending
	if v2383 != 0 {
		goto L29
	} else {
		goto L647
	}
L338:
	;
	v2327 = int32(0)
	v2332 = F_create_append_path(m, l0, l1, v2309, v2327, v1409, v2327, v2327, v2327, float64(-1))
	mBase = m.M
	v2333 = m.ExcPending
	if v2333 != 0 {
		goto L29
	} else {
		goto L638
	}
L339:
	;
	if v1649 == int32(0) {
		v2364 = v2281
		v2367 = v2284
		v2368 = v2285
		v2370 = v2287
		goto L337
	} else {
		goto L637
	}
L340:
	;
	v2269 = int32(0)
	if v1636 != 0 {
		v2309 = v2269
		v2312 = v2269
		v2313 = v2269
		v2315 = v2269
		goto L338
	} else {
		goto L636
	}
L341:
	;
	v1652 = int32(0)
	if v1648 == v1650 {
		v2281 = v1652
		v2284 = v1652
		v2285 = v1652
		v2287 = v1652
		goto L339
	} else {
		goto L428
	}
L342:
	;
	v1642 = int32(-1)
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1644 = int32(1)
	v1648 = v1642
	v1649 = v1644
	v1650 = v1643 - v1644
	v1651 = v1642
	goto L341
L343:
	;
	if l2 == int32(0) {
		goto L340
	} else {
		goto L427
	}
L344:
	;
	if v1462 != 0 {
		v1636 = v1404
		goto L343
	} else {
		goto L362
	}
L345:
	;
	v1462 = int32(1)
	goto L344
L346:
	;
	goto L347
L347:
	;
	v1418 = int32(0)
	goto L349
L348:
	;
	v1462 = v1454
	goto L344
L349:
	;
	v1422 = int32(0)
	if v1409 == v1422 {
		v1432 = v1422
		goto L351
	} else {
		goto L352
	}
L350:
	;
	v1454 = int32(0)
	goto L348
L351:
	;
	if v1371 != 0 {
		goto L355
	} else {
		goto L356
	}
L352:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+4))
	if v1426 <= v1418 {
		v1432 = int32(0)
		goto L351
	} else {
		goto L353
	}
L353:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+12))
	v1432 = v1428 + v1418<<(uint(int32(2))%32)
	goto L351
L354:
	;
	v1438 = base.B2i32(v1432 == int32(0))
	if v1432 == int32(0) {
		v1454 = v1438
		goto L348
	} else {
		goto L359
	}
L355:
	;
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v1371)+4))
	if v1418 < v1433 {
		goto L354
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	v1462 = base.B2i32(v1432 == int32(0))
	goto L344
L358:
	;
	goto L357
L359:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1371)+12))
	if v1441 == int32(0) {
		v1454 = v1438
		goto L348
	} else {
		goto L360
	}
L360:
	;
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1432)))
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v1418<<(uint(int32(2))%32)+v1441)))
	if v1448 == v1450 {
		v1418 = v1418 + int32(1)
		goto L349
	} else {
		goto L361
	}
L361:
	;
	goto L350
L362:
	;
	v1463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+31)))
	if v1463 == int32(0) {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	if v1371 == v1409 {
		goto L367
	} else {
		goto L368
	}
L364:
	;
	goto L365
L365:
	;
	if v1409 == v1372 {
		goto L386
	} else {
		goto L387
	}
L366:
	;
	if v1518 != 0 {
		v1636 = v1404
		goto L343
	} else {
		goto L384
	}
L367:
	;
	v1518 = int32(1)
	goto L366
L368:
	;
	goto L369
L369:
	;
	v1474 = int32(0)
	goto L371
L370:
	;
	v1518 = v1510
	goto L366
L371:
	;
	v1478 = int32(0)
	if v1371 == v1478 {
		v1488 = v1478
		goto L373
	} else {
		goto L374
	}
L372:
	;
	v1510 = int32(0)
	goto L370
L373:
	;
	if v1409 != 0 {
		goto L377
	} else {
		goto L378
	}
L374:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v1371)+4))
	if v1482 <= v1474 {
		v1488 = int32(0)
		goto L373
	} else {
		goto L375
	}
L375:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1371)+12))
	v1488 = v1484 + v1474<<(uint(int32(2))%32)
	goto L373
L376:
	;
	v1494 = base.B2i32(v1488 == int32(0))
	if v1488 == int32(0) {
		v1510 = v1494
		goto L370
	} else {
		goto L381
	}
L377:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+4))
	if v1474 < v1489 {
		goto L376
	} else {
		goto L380
	}
L378:
	;
	goto L379
L379:
	;
	v1518 = base.B2i32(v1488 == int32(0))
	goto L366
L380:
	;
	goto L379
L381:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+12))
	if v1497 == int32(0) {
		v1510 = v1494
		goto L370
	} else {
		goto L382
	}
L382:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1488)))
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1474<<(uint(int32(2))%32)+v1497)))
	if v1504 == v1506 {
		v1474 = v1474 + int32(1)
		goto L371
	} else {
		goto L383
	}
L383:
	;
	goto L372
L384:
	;
	goto L365
L385:
	;
	if v1571 == int32(0) {
		goto L403
	} else {
		goto L404
	}
L386:
	;
	v1571 = int32(1)
	goto L385
L387:
	;
	goto L388
L388:
	;
	v1527 = int32(0)
	goto L390
L389:
	;
	v1571 = v1563
	goto L385
L390:
	;
	v1531 = int32(0)
	if v1409 == v1531 {
		v1541 = v1531
		goto L392
	} else {
		goto L393
	}
L391:
	;
	v1563 = int32(0)
	goto L389
L392:
	;
	if v1372 != 0 {
		goto L396
	} else {
		goto L397
	}
L393:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+4))
	if v1535 <= v1527 {
		v1541 = int32(0)
		goto L392
	} else {
		goto L394
	}
L394:
	;
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+12))
	v1541 = v1537 + v1527<<(uint(int32(2))%32)
	goto L392
L395:
	;
	v1547 = base.B2i32(v1541 == int32(0))
	if v1541 == int32(0) {
		v1563 = v1547
		goto L389
	} else {
		goto L400
	}
L396:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v1372)+4))
	if v1527 < v1542 {
		goto L395
	} else {
		goto L399
	}
L397:
	;
	goto L398
L398:
	;
	v1571 = base.B2i32(v1541 == int32(0))
	goto L385
L399:
	;
	goto L398
L400:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1372)+12))
	if v1550 == int32(0) {
		v1563 = v1547
		goto L389
	} else {
		goto L401
	}
L401:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1541)))
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1527<<(uint(int32(2))%32)+v1550)))
	if v1557 == v1559 {
		v1527 = v1527 + int32(1)
		goto L390
	} else {
		goto L402
	}
L402:
	;
	goto L391
L403:
	;
	v1574 = int32(0)
	v1575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+30)))
	if v1575 != 0 {
		v1636 = v1574
		goto L343
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	if l2 != 0 {
		goto L342
	} else {
		goto L426
	}
L406:
	;
	if v1372 == v1409 {
		goto L408
	} else {
		goto L409
	}
L407:
	;
	if v1628 == int32(0) {
		v1636 = v1574
		goto L343
	} else {
		goto L425
	}
L408:
	;
	v1628 = int32(1)
	goto L407
L409:
	;
	goto L410
L410:
	;
	v1584 = int32(0)
	goto L412
L411:
	;
	v1628 = v1620
	goto L407
L412:
	;
	v1588 = int32(0)
	if v1372 == v1588 {
		v1598 = v1588
		goto L414
	} else {
		goto L415
	}
L413:
	;
	v1620 = int32(0)
	goto L411
L414:
	;
	if v1409 != 0 {
		goto L418
	} else {
		goto L419
	}
L415:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v1372)+4))
	if v1592 <= v1584 {
		v1598 = int32(0)
		goto L414
	} else {
		goto L416
	}
L416:
	;
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v1372)+12))
	v1598 = v1594 + v1584<<(uint(int32(2))%32)
	goto L414
L417:
	;
	v1604 = base.B2i32(v1598 == int32(0))
	if v1598 == int32(0) {
		v1620 = v1604
		goto L411
	} else {
		goto L422
	}
L418:
	;
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+4))
	if v1584 < v1599 {
		goto L417
	} else {
		goto L421
	}
L419:
	;
	goto L420
L420:
	;
	v1628 = base.B2i32(v1598 == int32(0))
	goto L407
L421:
	;
	goto L420
L422:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+12))
	if v1607 == int32(0) {
		v1620 = v1604
		goto L411
	} else {
		goto L423
	}
L423:
	;
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1598)))
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v1584<<(uint(int32(2))%32)+v1607)))
	if v1614 == v1616 {
		v1584 = v1584 + int32(1)
		goto L412
	} else {
		goto L424
	}
L424:
	;
	goto L413
L425:
	;
	goto L405
L426:
	;
	v1632 = int32(0)
	v2309 = v1632
	v2312 = v1632
	v2313 = v1632
	v2315 = v1632
	goto L338
L427:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1648 = v1639
	v1649 = v1636
	v1650 = int32(0)
	v1651 = int32(1)
	goto L341
L428:
	;
	v1665 = v1652
	v1668 = v1652
	v1669 = v1652
	v1671 = v1652
	v1675 = v1650
	goto L429
L429:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1683+v1675<<(uint(int32(2))%32))))
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v1687)+32))
	v1689 = int32(0)
	if v1688 == v1689 {
		goto L432
	} else {
		goto L433
	}
L430:
	;
	v2281 = v2262
	v2284 = v2266
	v2285 = v2263
	v2287 = v2264
	goto L339
L431:
	;
	v1838 = int32(0)
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1687)+32))
	if v1839 == v1838 {
		goto L475
	} else {
		goto L476
	}
L432:
	;
	v1837 = int32(0)
	goto L431
L433:
	;
	goto L434
L434:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v1688)+4))
	if int32(0) < v1703 {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v1713 = v1689
	v1716 = v1689
	goto L438
L436:
	;
	v1818 = v1689
	goto L437
L437:
	;
	v1837 = v1818
	goto L431
L438:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1688)+12))
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1719+v1716<<(uint(int32(2))%32))))
	goto L442
L439:
	;
	v1818 = v1801
	goto L437
L440:
	;
	v1808 = v1716 + int32(1)
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v1688)+4))
	if v1808 < v1809 {
		v1713 = v1801
		v1716 = v1808
		goto L438
	} else {
		goto L473
	}
L442:
	;
	goto L443
L443:
	;
	if v1713 != 0 {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	v1727 = F_compare_path_costs(m, v1713, v1723, v1689)
	mBase = m.M
	if v1727 <= int32(0) {
		v1801 = v1713
		goto L440
	} else {
		goto L448
	}
L446:
	;
	goto L447
L447:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1723)+64))
	if v1409 == v1730 {
		goto L449
	} else {
		goto L450
	}
L448:
	;
	goto L447
L449:
	;
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v1723)+16))
	if v1788 != 0 {
		goto L467
	} else {
		goto L468
	}
L450:
	;
	v1738 = int32(0)
	goto L451
L451:
	;
	v1746 = int32(0)
	if v1409 == v1746 {
		v1756 = v1746
		goto L453
	} else {
		goto L454
	}
L452:
	;
	if v1756 != 0 {
		v1801 = v1713
		goto L440
	} else {
		goto L466
	}
L453:
	;
	if v1730 != 0 {
		goto L457
	} else {
		goto L458
	}
L454:
	;
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+4))
	if v1750 <= v1738 {
		v1756 = int32(0)
		goto L453
	} else {
		goto L455
	}
L455:
	;
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+12))
	v1756 = v1752 + v1738<<(uint(int32(2))%32)
	goto L453
L456:
	;
	if v1756 == int32(0) {
		goto L462
	} else {
		goto L463
	}
L457:
	;
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1730)+4))
	if v1738 < v1757 {
		goto L456
	} else {
		goto L460
	}
L458:
	;
	goto L459
L459:
	;
	if v1756 == int32(0) {
		goto L449
	} else {
		goto L461
	}
L460:
	;
	goto L459
L461:
	;
	v1801 = v1713
	goto L440
L462:
	;
	goto L452
L463:
	;
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1730)+12))
	if v1763 == int32(0) {
		goto L462
	} else {
		goto L464
	}
L464:
	;
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(v1756)))
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1763+v1738<<(uint(int32(2))%32))))
	if v1770 == v1772 {
		v1738 = v1738 + int32(1)
		goto L451
	} else {
		goto L465
	}
L465:
	;
	v1801 = v1713
	goto L440
L466:
	;
	goto L449
L467:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v1788)+4))
	v1791 = v1789
	goto L469
L468:
	;
	v1791 = int32(0)
	goto L469
L469:
	;
	v1792 = F_bms_is_subset(m, v1791, v1689)
	mBase = m.M
	if v1792 != 0 {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v1793 = v1723
	goto L472
L471:
	;
	v1793 = v1713
	goto L472
L472:
	;
	v1801 = v1793
	goto L440
L473:
	;
	goto L439
L474:
	;
	if v1988 != 0 {
		goto L517
	} else {
		goto L518
	}
L475:
	;
	v1988 = int32(0)
	goto L474
L476:
	;
	goto L477
L477:
	;
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(v1839)+4))
	if int32(0) < v1854 {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v1864 = v1838
	v1867 = v1838
	goto L481
L479:
	;
	v1969 = v1838
	goto L480
L480:
	;
	v1988 = v1969
	goto L474
L481:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(v1839)+12))
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1870+v1867<<(uint(int32(2))%32))))
	goto L485
L482:
	;
	v1969 = v1952
	goto L480
L483:
	;
	v1959 = v1867 + int32(1)
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(v1839)+4))
	if v1959 < v1960 {
		v1864 = v1952
		v1867 = v1959
		goto L481
	} else {
		goto L516
	}
L485:
	;
	goto L486
L486:
	;
	if v1864 != 0 {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v1878 = F_compare_path_costs(m, v1864, v1874, int32(1))
	mBase = m.M
	if v1878 <= int32(0) {
		v1952 = v1864
		goto L483
	} else {
		goto L491
	}
L489:
	;
	goto L490
L490:
	;
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v1874)+64))
	if v1409 == v1881 {
		goto L492
	} else {
		goto L493
	}
L491:
	;
	goto L490
L492:
	;
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v1874)+16))
	if v1939 != 0 {
		goto L510
	} else {
		goto L511
	}
L493:
	;
	v1889 = int32(0)
	goto L494
L494:
	;
	v1897 = int32(0)
	if v1409 == v1897 {
		v1907 = v1897
		goto L496
	} else {
		goto L497
	}
L495:
	;
	if v1907 != 0 {
		v1952 = v1864
		goto L483
	} else {
		goto L509
	}
L496:
	;
	if v1881 != 0 {
		goto L500
	} else {
		goto L501
	}
L497:
	;
	v1901 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+4))
	if v1901 <= v1889 {
		v1907 = int32(0)
		goto L496
	} else {
		goto L498
	}
L498:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+12))
	v1907 = v1903 + v1889<<(uint(int32(2))%32)
	goto L496
L499:
	;
	if v1907 == int32(0) {
		goto L505
	} else {
		goto L506
	}
L500:
	;
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v1881)+4))
	if v1889 < v1908 {
		goto L499
	} else {
		goto L503
	}
L501:
	;
	goto L502
L502:
	;
	if v1907 == int32(0) {
		goto L492
	} else {
		goto L504
	}
L503:
	;
	goto L502
L504:
	;
	v1952 = v1864
	goto L483
L505:
	;
	goto L495
L506:
	;
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v1881)+12))
	if v1914 == int32(0) {
		goto L505
	} else {
		goto L507
	}
L507:
	;
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v1907)))
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v1914+v1889<<(uint(int32(2))%32))))
	if v1921 == v1923 {
		v1889 = v1889 + int32(1)
		goto L494
	} else {
		goto L508
	}
L508:
	;
	v1952 = v1864
	goto L483
L509:
	;
	goto L492
L510:
	;
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v1939)+4))
	v1942 = v1940
	goto L512
L511:
	;
	v1942 = int32(0)
	goto L512
L512:
	;
	v1943 = F_bms_is_subset(m, v1942, v1838)
	mBase = m.M
	if v1943 != 0 {
		goto L513
	} else {
		goto L514
	}
L513:
	;
	v1944 = v1874
	goto L515
L514:
	;
	v1944 = v1864
	goto L515
L515:
	;
	v1952 = v1944
	goto L483
L516:
	;
	goto L482
L517:
	;
	v1989 = v1837
	goto L519
L518:
	;
	v1989 = v1838
	goto L519
L519:
	;
	if v1989 == int32(0) {
		goto L520
	} else {
		goto L521
	}
L520:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v1687)+48))
	v1993 = v1992
	v1994 = v1992
	goto L522
L521:
	;
	v1993 = v1988
	v1994 = v1837
	goto L522
L522:
	;
	v1996 = *(*float64)(unsafe.Add(mBase, uint32(l0)+296))
	if base.F64_gt(v1996, float64(0)) != 0 {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v1687)+32))
	if base.F64_ge(v1996, float64(1)) != 0 {
		goto L526
	} else {
		goto L527
	}
L524:
	;
	v2139 = int32(0)
	goto L525
L525:
	;
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(v1994)))
	if v1649 != 0 {
		goto L573
	} else {
		goto L574
	}
L526:
	;
	v2002 = *(*float64)(unsafe.Add(mBase, uint32(v1993)+32))
	v2004 = base.F64_div(v1996, v2002)
	goto L528
L527:
	;
	v2004 = v1996
	goto L528
L528:
	;
	v2005 = int32(0)
	if v1999 == v2005 {
		goto L530
	} else {
		goto L531
	}
L529:
	;
	if v2136 != 0 {
		goto L568
	} else {
		goto L569
	}
L530:
	;
	v2136 = int32(0)
	goto L529
L531:
	;
	goto L532
L532:
	;
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+4))
	if int32(0) < v2016 {
		goto L533
	} else {
		goto L534
	}
L533:
	;
	v2024 = v2005
	v2027 = v2005
	goto L536
L534:
	;
	v2119 = v2005
	goto L535
L535:
	;
	v2136 = v2119
	goto L529
L536:
	;
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+12))
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v2030+v2027<<(uint(int32(2))%32))))
	if v2024 != 0 {
		goto L539
	} else {
		goto L540
	}
L537:
	;
	v2119 = v2104
	goto L535
L538:
	;
	v2111 = v2027 + int32(1)
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(v1999)+4))
	if v2111 < v2112 {
		v2024 = v2104
		v2027 = v2111
		goto L536
	} else {
		goto L567
	}
L539:
	;
	v2035 = F_compare_fractional_path_costs(m, v2024, v2034, v2004)
	mBase = m.M
	if v2035 <= int32(0) {
		v2104 = v2024
		goto L538
	} else {
		goto L542
	}
L540:
	;
	goto L541
L541:
	;
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v2034)+64))
	if v1409 == v2038 {
		goto L543
	} else {
		goto L544
	}
L542:
	;
	goto L541
L543:
	;
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v2034)+16))
	if v2092 != 0 {
		goto L561
	} else {
		goto L562
	}
L544:
	;
	v2044 = int32(0)
	goto L545
L545:
	;
	v2052 = int32(0)
	if v1409 == v2052 {
		v2062 = v2052
		goto L547
	} else {
		goto L548
	}
L546:
	;
	if v2062 != 0 {
		v2104 = v2024
		goto L538
	} else {
		goto L560
	}
L547:
	;
	if v2038 != 0 {
		goto L551
	} else {
		goto L552
	}
L548:
	;
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+4))
	if v2056 <= v2044 {
		v2062 = int32(0)
		goto L547
	} else {
		goto L549
	}
L549:
	;
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+12))
	v2062 = v2058 + v2044<<(uint(int32(2))%32)
	goto L547
L550:
	;
	if v2062 == int32(0) {
		goto L556
	} else {
		goto L557
	}
L551:
	;
	v2063 = *(*int32)(unsafe.Add(mBase, uint32(v2038)+4))
	if v2044 < v2063 {
		goto L550
	} else {
		goto L554
	}
L552:
	;
	goto L553
L553:
	;
	if v2062 == int32(0) {
		goto L543
	} else {
		goto L555
	}
L554:
	;
	goto L553
L555:
	;
	v2104 = v2024
	goto L538
L556:
	;
	goto L546
L557:
	;
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v2038)+12))
	if v2069 == int32(0) {
		goto L556
	} else {
		goto L558
	}
L558:
	;
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(v2062)))
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v2069+v2044<<(uint(int32(2))%32))))
	if v2076 == v2078 {
		v2044 = v2044 + int32(1)
		goto L545
	} else {
		goto L559
	}
L559:
	;
	v2104 = v2024
	goto L538
L560:
	;
	goto L543
L561:
	;
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v2092)+4))
	v2095 = v2093
	goto L563
L562:
	;
	v2095 = int32(0)
	goto L563
L563:
	;
	v2097 = F_bms_is_subset(m, v2095, int32(0))
	mBase = m.M
	if v2097 != 0 {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v2098 = v2034
	goto L566
L565:
	;
	v2098 = v2024
	goto L566
L566:
	;
	v2104 = v2098
	goto L538
L567:
	;
	goto L537
L568:
	;
	v2137 = v2136
	goto L570
L569:
	;
	v2137 = v1993
	goto L570
L570:
	;
	v2139 = v2137
	goto L525
L571:
	;
	v2266 = base.B2i32(v1993 != v1994) | v1668
	v2267 = v1651 + v1675
	if v1648 != v2267 {
		v1665 = v2262
		v1668 = v2266
		v1669 = v2263
		v1671 = v2264
		v1675 = v2267
		goto L429
	} else {
		goto L635
	}
L572:
	;
	v2258 = F_lappend(m, v1669, v2139)
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
		goto L29
	} else {
		goto L634
	}
L573:
	;
	switch v2141 - int32(290) {
	case 0:
		goto L579
	case 1:
		goto L578
	default:
		v2159 = v1994
		goto L576
	}
L574:
	;
	goto L575
L575:
	;
	switch v2141 - int32(290) {
	case 0:
		goto L606
	case 1:
		goto L605
	default:
		goto L604
	}
L576:
	;
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(v1993)))
	switch v2161 - int32(290) {
	case 0:
		goto L587
	case 1:
		goto L586
	default:
		v2179 = v1993
		goto L584
	}
L577:
	;
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v2156)+12))
	v2158 = *(*int32)(unsafe.Add(mBase, uint32(v2157)))
	v2159 = v2158
	goto L576
L578:
	;
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(v1994)+72))
	if v2150 == int32(0) {
		v2159 = v1994
		goto L576
	} else {
		goto L582
	}
L579:
	;
	v2144 = *(*int32)(unsafe.Add(mBase, uint32(v1994)+72))
	if v2144 == int32(0) {
		v2159 = v1994
		goto L576
	} else {
		goto L580
	}
L580:
	;
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v2144)+4))
	if v2147 == int32(1) {
		v2156 = v2144
		goto L577
	} else {
		goto L581
	}
L581:
	;
	v2159 = v1994
	goto L576
L582:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v2150)+4))
	if v2153 != int32(1) {
		v2159 = v1994
		goto L576
	} else {
		goto L583
	}
L583:
	;
	v2156 = v2150
	goto L577
L584:
	;
	v2181 = F_lappend(m, v1665, v2159)
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L29
	} else {
		goto L592
	}
L585:
	;
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(v2176)+12))
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(v2177)))
	v2179 = v2178
	goto L584
L586:
	;
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v1993)+72))
	if v2170 == int32(0) {
		v2179 = v1993
		goto L584
	} else {
		goto L590
	}
L587:
	;
	v2164 = *(*int32)(unsafe.Add(mBase, uint32(v1993)+72))
	if v2164 == int32(0) {
		v2179 = v1993
		goto L584
	} else {
		goto L588
	}
L588:
	;
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(v2164)+4))
	if v2167 == int32(1) {
		v2176 = v2164
		goto L585
	} else {
		goto L589
	}
L589:
	;
	v2179 = v1993
	goto L584
L590:
	;
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v2170)+4))
	if v2173 != int32(1) {
		v2179 = v1993
		goto L584
	} else {
		goto L591
	}
L591:
	;
	v2176 = v2170
	goto L585
L592:
	;
	v2183 = F_lappend(m, v1671, v2179)
	mBase = m.M
	v2184 = m.ExcPending
	if v2184 != 0 {
		goto L29
	} else {
		goto L593
	}
L593:
	;
	if v2139 == int32(0) {
		v2262 = v2181
		v2263 = v1669
		v2264 = v2183
		goto L571
	} else {
		goto L594
	}
L594:
	;
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(v2139)))
	switch v2187 - int32(290) {
	case 0:
		goto L597
	case 1:
		goto L596
	default:
		v2255 = v2181
		v2256 = v2183
		goto L572
	}
L595:
	;
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(v2202)+12))
	v2204 = *(*int32)(unsafe.Add(mBase, uint32(v2203)))
	v2205 = F_lappend(m, v1669, v2204)
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		goto L29
	} else {
		goto L602
	}
L596:
	;
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(v2139)+72))
	if v2196 == int32(0) {
		v2255 = v2181
		v2256 = v2183
		goto L572
	} else {
		goto L600
	}
L597:
	;
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(v2139)+72))
	if v2190 == int32(0) {
		v2255 = v2181
		v2256 = v2183
		goto L572
	} else {
		goto L598
	}
L598:
	;
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(v2190)+4))
	if v2193 == int32(1) {
		v2202 = v2190
		goto L595
	} else {
		goto L599
	}
L599:
	;
	v2255 = v2181
	v2256 = v2183
	goto L572
L600:
	;
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(v2196)+4))
	if v2199 != int32(1) {
		v2255 = v2181
		v2256 = v2183
		goto L572
	} else {
		goto L601
	}
L601:
	;
	v2202 = v2196
	goto L595
L602:
	;
	v2262 = v2181
	v2263 = v2205
	v2264 = v2183
	goto L571
L603:
	;
	v2222 = *(*int32)(unsafe.Add(mBase, uint32(v1993)))
	switch v2222 - int32(290) {
	case 0:
		goto L617
	case 1:
		goto L616
	default:
		goto L615
	}
L604:
	;
	v2219 = F_lappend(m, v1665, v1994)
	mBase = m.M
	v2220 = m.ExcPending
	if v2220 != 0 {
		goto L29
	} else {
		goto L613
	}
L605:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v1994)+72))
	v2217 = F_list_concat(m, v1665, v2216)
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L29
	} else {
		goto L612
	}
L606:
	;
	v2209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1994)+20)))
	if v2209 == int32(1) {
		goto L607
	} else {
		goto L608
	}
L607:
	;
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(v1994)+76))
	if v2212 != 0 {
		goto L604
	} else {
		goto L610
	}
L608:
	;
	goto L609
L609:
	;
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v1994)+72))
	v2214 = F_list_concat(m, v1665, v2213)
	mBase = m.M
	v2215 = m.ExcPending
	if v2215 != 0 {
		goto L29
	} else {
		goto L611
	}
L610:
	;
	goto L609
L611:
	;
	v2221 = v2214
	goto L603
L612:
	;
	v2221 = v2217
	goto L603
L613:
	;
	v2221 = v2219
	goto L603
L614:
	;
	if v2139 == int32(0) {
		v2262 = v2221
		v2263 = v1669
		v2264 = v2237
		goto L571
	} else {
		goto L625
	}
L615:
	;
	v2235 = F_lappend(m, v1671, v1993)
	mBase = m.M
	v2236 = m.ExcPending
	if v2236 != 0 {
		goto L29
	} else {
		goto L624
	}
L616:
	;
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(v1993)+72))
	v2233 = F_list_concat(m, v1671, v2232)
	mBase = m.M
	v2234 = m.ExcPending
	if v2234 != 0 {
		goto L29
	} else {
		goto L623
	}
L617:
	;
	v2225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1993)+20)))
	if v2225 == int32(1) {
		goto L618
	} else {
		goto L619
	}
L618:
	;
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v1993)+76))
	if v2228 != 0 {
		goto L615
	} else {
		goto L621
	}
L619:
	;
	goto L620
L620:
	;
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(v1993)+72))
	v2230 = F_list_concat(m, v1671, v2229)
	mBase = m.M
	v2231 = m.ExcPending
	if v2231 != 0 {
		goto L29
	} else {
		goto L622
	}
L621:
	;
	goto L620
L622:
	;
	v2237 = v2230
	goto L614
L623:
	;
	v2237 = v2233
	goto L614
L624:
	;
	v2237 = v2235
	goto L614
L625:
	;
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v2139)))
	switch v2240 - int32(290) {
	case 0:
		goto L627
	case 1:
		goto L626
	default:
		v2255 = v2221
		v2256 = v2237
		goto L572
	}
L626:
	;
	v2250 = *(*int32)(unsafe.Add(mBase, uint32(v2139)+72))
	v2251 = F_list_concat(m, v1669, v2250)
	mBase = m.M
	v2252 = m.ExcPending
	if v2252 != 0 {
		goto L29
	} else {
		goto L633
	}
L627:
	;
	v2243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2139)+20)))
	if v2243 == int32(1) {
		goto L628
	} else {
		goto L629
	}
L628:
	;
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v2139)+76))
	if v2246 != 0 {
		v2255 = v2221
		v2256 = v2237
		goto L572
	} else {
		goto L631
	}
L629:
	;
	goto L630
L630:
	;
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(v2139)+72))
	v2248 = F_list_concat(m, v1669, v2247)
	mBase = m.M
	v2249 = m.ExcPending
	if v2249 != 0 {
		goto L29
	} else {
		goto L632
	}
L631:
	;
	goto L630
L632:
	;
	v2262 = v2221
	v2263 = v2248
	v2264 = v2237
	goto L571
L633:
	;
	v2262 = v2221
	v2263 = v2251
	v2264 = v2237
	goto L571
L634:
	;
	v2262 = v2255
	v2263 = v2258
	v2264 = v2256
	goto L571
L635:
	;
	goto L430
L636:
	;
	v2364 = v2269
	v2367 = v2269
	v2368 = v2269
	v2370 = v2269
	goto L337
L637:
	;
	v2309 = v2281
	v2312 = v2284
	v2313 = v2285
	v2315 = v2287
	goto L338
L638:
	;
	F_add_path(m, l1, v2332)
	mBase = m.M
	v2335 = m.ExcPending
	if v2335 != 0 {
		goto L29
	} else {
		goto L639
	}
L639:
	;
	if v2312&int32(1) != 0 {
		goto L640
	} else {
		goto L641
	}
L640:
	;
	v2338 = int32(0)
	v2343 = F_create_append_path(m, l0, l1, v2315, v2338, v1409, v2338, v2338, v2338, float64(-1))
	mBase = m.M
	v2344 = m.ExcPending
	if v2344 != 0 {
		goto L29
	} else {
		goto L643
	}
L641:
	;
	goto L642
L642:
	;
	if v2313 == int32(0) {
		goto L335
	} else {
		goto L645
	}
L643:
	;
	F_add_path(m, l1, v2343)
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		goto L29
	} else {
		goto L644
	}
L644:
	;
	goto L642
L645:
	;
	v2349 = int32(0)
	v2354 = F_create_append_path(m, l0, l1, v2313, v2349, v1409, v2349, v2349, v2349, float64(-1))
	mBase = m.M
	v2355 = m.ExcPending
	if v2355 != 0 {
		goto L29
	} else {
		goto L646
	}
L646:
	;
	v2422 = v2354
	goto L336
L647:
	;
	F_add_path(m, l1, v2382)
	mBase = m.M
	v2385 = m.ExcPending
	if v2385 != 0 {
		goto L29
	} else {
		goto L648
	}
L648:
	;
	if v2367&int32(1) != 0 {
		goto L649
	} else {
		goto L650
	}
L649:
	;
	v2388 = F_create_merge_append_path(m, l0, l1, v2370, v1409)
	mBase = m.M
	v2389 = m.ExcPending
	if v2389 != 0 {
		goto L29
	} else {
		goto L652
	}
L650:
	;
	goto L651
L651:
	;
	if v2368 == int32(0) {
		goto L335
	} else {
		goto L654
	}
L652:
	;
	F_add_path(m, l1, v2388)
	mBase = m.M
	v2391 = m.ExcPending
	if v2391 != 0 {
		goto L29
	} else {
		goto L653
	}
L653:
	;
	goto L651
L654:
	;
	v2394 = F_create_merge_append_path(m, l0, l1, v2368, v1409)
	mBase = m.M
	v2395 = m.ExcPending
	if v2395 != 0 {
		goto L29
	} else {
		goto L655
	}
L655:
	;
	v2422 = v2394
	goto L336
L656:
	;
	goto L335
L657:
	;
	goto L334
L658:
	;
	if l2 == int32(0) {
		goto L846
	} else {
		goto L847
	}
L659:
	;
	v2483 = *(*int32)(unsafe.Add(mBase, uint32(v828)+4))
	if v2483 <= int32(0) {
		goto L658
	} else {
		goto L660
	}
L660:
	;
	v2501 = int32(0)
	goto L661
L661:
	;
	v2513 = *(*int32)(unsafe.Add(mBase, uint32(v828)+12))
	v2517 = *(*int32)(unsafe.Add(mBase, uint32(v2513+v2501<<(uint(int32(2))%32))))
	v2518 = int32(0)
	if l2 == v2518 {
		v3095 = v2518
		goto L664
	} else {
		goto L665
	}
L662:
	;
	goto L658
L663:
	;
	v3149 = v2501 + int32(1)
	v3150 = *(*int32)(unsafe.Add(mBase, uint32(v828)+4))
	if v3149 < v3150 {
		v2501 = v3149
		goto L661
	} else {
		goto L845
	}
L664:
	;
	v3113 = int32(0)
	v3118 = F_create_append_path(m, l0, l1, v3095, v3113, v3113, v2517, v3113, v3113, float64(-1))
	mBase = m.M
	v3119 = m.ExcPending
	if v3119 != 0 {
		goto L29
	} else {
		goto L843
	}
L665:
	;
	v2521 = int32(0)
	v2522 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v2522 <= v2521 {
		v3095 = v2518
		goto L664
	} else {
		goto L666
	}
L666:
	;
	v2533 = v2518
	v2536 = v2521
	goto L667
L667:
	;
	v2551 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v2555 = *(*int32)(unsafe.Add(mBase, uint32(v2551+v2536<<(uint(int32(2))%32))))
	v2556 = *(*int32)(unsafe.Add(mBase, uint32(v2555)+32))
	if v2556 == int32(0) {
		goto L663
	} else {
		goto L669
	}
L668:
	;
	v3095 = v3082
	goto L664
L669:
	;
	v2559 = int32(0)
	if v2556 == v2559 {
		goto L671
	} else {
		goto L672
	}
L670:
	;
	v2709 = *(*int32)(unsafe.Add(mBase, uint32(v2708)+16))
	if v2709 != 0 {
		goto L713
	} else {
		goto L714
	}
L671:
	;
	v2708 = int32(0)
	goto L670
L672:
	;
	goto L673
L673:
	;
	v2574 = *(*int32)(unsafe.Add(mBase, uint32(v2556)+4))
	if int32(0) < v2574 {
		goto L674
	} else {
		goto L675
	}
L674:
	;
	v2584 = v2559
	v2587 = v2559
	goto L677
L675:
	;
	v2689 = v2559
	goto L676
L676:
	;
	v2708 = v2689
	goto L670
L677:
	;
	v2590 = *(*int32)(unsafe.Add(mBase, uint32(v2556)+12))
	v2594 = *(*int32)(unsafe.Add(mBase, uint32(v2590+v2587<<(uint(int32(2))%32))))
	goto L681
L678:
	;
	v2689 = v2672
	goto L676
L679:
	;
	v2679 = v2587 + int32(1)
	v2680 = *(*int32)(unsafe.Add(mBase, uint32(v2556)+4))
	if v2679 < v2680 {
		v2584 = v2672
		v2587 = v2679
		goto L677
	} else {
		goto L712
	}
L681:
	;
	goto L682
L682:
	;
	if v2584 != 0 {
		goto L684
	} else {
		goto L685
	}
L684:
	;
	v2598 = F_compare_path_costs(m, v2584, v2594, int32(1))
	mBase = m.M
	if v2598 <= int32(0) {
		v2672 = v2584
		goto L679
	} else {
		goto L687
	}
L685:
	;
	goto L686
L686:
	;
	v2601 = *(*int32)(unsafe.Add(mBase, uint32(v2594)+64))
	if v2559 == v2601 {
		goto L688
	} else {
		goto L689
	}
L687:
	;
	goto L686
L688:
	;
	v2659 = *(*int32)(unsafe.Add(mBase, uint32(v2594)+16))
	if v2659 != 0 {
		goto L706
	} else {
		goto L707
	}
L689:
	;
	goto L690
L690:
	;
	goto L692
L691:
	;
	goto L705
L692:
	;
	if v2601 != 0 {
		goto L696
	} else {
		goto L697
	}
L695:
	;
	goto L701
L696:
	;
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v2601)+4))
	if int32(0) < v2628 {
		goto L695
	} else {
		goto L699
	}
L697:
	;
	goto L698
L698:
	;
	goto L688
L699:
	;
	goto L698
L701:
	;
	goto L691
L705:
	;
	goto L688
L706:
	;
	v2660 = *(*int32)(unsafe.Add(mBase, uint32(v2659)+4))
	v2662 = v2660
	goto L708
L707:
	;
	v2662 = int32(0)
	goto L708
L708:
	;
	v2663 = F_bms_is_subset(m, v2662, v2517)
	mBase = m.M
	if v2663 != 0 {
		goto L709
	} else {
		goto L710
	}
L709:
	;
	v2664 = v2594
	goto L711
L710:
	;
	v2664 = v2584
	goto L711
L711:
	;
	v2672 = v2664
	goto L679
L712:
	;
	goto L678
L713:
	;
	v2710 = *(*int32)(unsafe.Add(mBase, uint32(v2709)+4))
	v2712 = v2710
	goto L715
L714:
	;
	v2712 = int32(0)
	goto L715
L715:
	;
	v2713 = int32(0)
	if base.B2i32(v2712 == v2713)|base.B2i32(v2517 == v2713) != 0 {
		v2759 = base.B2i32(v2712|v2517 == v2713)
		goto L717
	} else {
		goto L718
	}
L716:
	;
	if v2759 == int32(0) {
		goto L727
	} else {
		goto L728
	}
L717:
	;
	goto L716
L718:
	;
	v2727 = *(*int32)(unsafe.Add(mBase, uint32(v2712)+4))
	v2728 = *(*int32)(unsafe.Add(mBase, uint32(v2517)+4))
	if v2727 != v2728 {
		v2759 = int32(0)
		goto L717
	} else {
		goto L719
	}
L719:
	;
	v2730 = int32(1)
	if v2727 <= v2730 {
		goto L720
	} else {
		goto L721
	}
L720:
	;
	v2733 = v2730
	goto L722
L721:
	;
	v2733 = v2727
	goto L722
L722:
	;
	v2734 = int32(8)
	v2739 = int32(0)
	goto L723
L723:
	;
	v2747 = v2739 << (uint(int32(2)) % 32)
	v2749 = *(*int32)(unsafe.Add(mBase, uint32(v2712+v2734+v2747)))
	v2751 = *(*int32)(unsafe.Add(mBase, uint32(v2517+v2734+v2747)))
	v2752 = base.B2i32(v2749 == v2751)
	if v2749 != v2751 {
		v2759 = v2752
		goto L717
	} else {
		goto L725
	}
L724:
	;
	v2759 = v2752
	goto L717
L725:
	;
	v2755 = v2739 + int32(1)
	if v2755 != v2733 {
		v2739 = v2755
		goto L723
	} else {
		goto L726
	}
L726:
	;
	goto L724
L727:
	;
	v2766 = *(*int32)(unsafe.Add(mBase, uint32(v2555)+32))
	if v2766 == int32(0) {
		goto L663
	} else {
		goto L730
	}
L728:
	;
	v3044 = v2708
	goto L729
L729:
	;
	v3067 = *(*int32)(unsafe.Add(mBase, uint32(v3044)))
	switch v3067 - int32(290) {
	case 0:
		goto L834
	case 1:
		goto L833
	default:
		goto L832
	}
L730:
	;
	v2769 = int32(0)
	v2770 = *(*int32)(unsafe.Add(mBase, uint32(v2766)+4))
	if v2770 <= v2769 {
		goto L663
	} else {
		goto L731
	}
L731:
	;
	v2776 = v2769
	v2792 = v2559
	goto L732
L732:
	;
	v2799 = *(*int32)(unsafe.Add(mBase, uint32(v2766)+12))
	v2803 = *(*int32)(unsafe.Add(mBase, uint32(v2799+v2792<<(uint(int32(2))%32))))
	v2804 = *(*int32)(unsafe.Add(mBase, uint32(v2803)+16))
	if v2804 != 0 {
		goto L735
	} else {
		goto L736
	}
L733:
	;
	if v3033 == int32(0) {
		goto L663
	} else {
		goto L830
	}
L734:
	;
	v3036 = v2792 + int32(1)
	v3037 = *(*int32)(unsafe.Add(mBase, uint32(v2766)+4))
	if v3036 < v3037 {
		v2776 = v3033
		v2792 = v3036
		goto L732
	} else {
		goto L829
	}
L735:
	;
	v2805 = *(*int32)(unsafe.Add(mBase, uint32(v2804)+4))
	v2807 = v2805
	goto L737
L736:
	;
	v2807 = int32(0)
	goto L737
L737:
	;
	v2808 = int32(0)
	if v2807 == v2808 {
		goto L739
	} else {
		goto L740
	}
L738:
	;
	if v2861 == int32(0) {
		goto L752
	} else {
		goto L753
	}
L739:
	;
	v2861 = int32(1)
	goto L738
L740:
	;
	goto L741
L741:
	;
	if v2517 == int32(0) {
		v2854 = v2808
		goto L742
	} else {
		goto L743
	}
L742:
	;
	v2861 = v2854
	goto L738
L743:
	;
	v2817 = *(*int32)(unsafe.Add(mBase, uint32(v2807)+4))
	v2818 = *(*int32)(unsafe.Add(mBase, uint32(v2517)+4))
	if v2818 < v2817 {
		v2854 = v2808
		goto L742
	} else {
		goto L744
	}
L744:
	;
	v2820 = int32(1)
	if v2817 <= v2820 {
		goto L745
	} else {
		goto L746
	}
L745:
	;
	v2823 = v2820
	goto L747
L746:
	;
	v2823 = v2817
	goto L747
L747:
	;
	v2824 = int32(8)
	v2829 = int32(0)
	goto L748
L748:
	;
	v2836 = v2829 << (uint(int32(2)) % 32)
	v2838 = *(*int32)(unsafe.Add(mBase, uint32(v2807+v2824+v2836)))
	v2840 = *(*int32)(unsafe.Add(mBase, uint32(v2517+v2824+v2836)))
	v2843 = v2838 & (v2840 ^ int32(-1))
	v2845 = base.B2i32(v2843 == int32(0))
	if v2843 != 0 {
		v2854 = v2845
		goto L742
	} else {
		goto L750
	}
L749:
	;
	v2854 = v2845
	goto L742
L750:
	;
	v2847 = v2829 + int32(1)
	if v2847 != v2823 {
		v2829 = v2847
		goto L748
	} else {
		goto L751
	}
L751:
	;
	goto L749
L752:
	;
	v3033 = v2776
	goto L734
L753:
	;
	goto L754
L754:
	;
	if v2776 == int32(0) {
		goto L755
	} else {
		goto L756
	}
L755:
	;
	v2917 = *(*int32)(unsafe.Add(mBase, uint32(v2803)+16))
	if v2917 != 0 {
		goto L782
	} else {
		goto L783
	}
L756:
	;
	v2871 = *(*int32)(unsafe.Add(mBase, uint32(v2776)+40))
	v2872 = *(*int32)(unsafe.Add(mBase, uint32(v2803)+40))
	if v2871 != v2872 {
		goto L758
	} else {
		goto L759
	}
L757:
	;
	if int32(0) < v2914 {
		goto L755
	} else {
		goto L781
	}
L758:
	;
	if v2871 < v2872 {
		goto L761
	} else {
		goto L762
	}
L759:
	;
	goto L760
L760:
	;
	goto L767
L761:
	;
	v2877 = int32(-1)
	goto L763
L762:
	;
	v2877 = int32(1)
	goto L763
L763:
	;
	v2914 = v2877
	goto L757
L764:
	;
	v2914 = v2908
	goto L757
L765:
	;
	v2908 = int32(0)
	goto L764
L767:
	;
	goto L768
L768:
	;
	v2893 = int32(-1)
	v2894 = *(*float64)(unsafe.Add(mBase, uint32(v2776)+56))
	v2895 = *(*float64)(unsafe.Add(mBase, uint32(v2803)+56))
	if base.F64_lt(v2894, v2895) != 0 {
		v2908 = v2893
		goto L764
	} else {
		goto L775
	}
L775:
	;
	if base.F64_gt(v2894, v2895) != 0 {
		goto L776
	} else {
		goto L777
	}
L776:
	;
	v2914 = int32(1)
	goto L757
L777:
	;
	goto L778
L778:
	;
	v2899 = *(*float64)(unsafe.Add(mBase, uint32(v2776)+48))
	v2900 = *(*float64)(unsafe.Add(mBase, uint32(v2803)+48))
	if base.F64_lt(v2899, v2900) != 0 {
		v2908 = v2893
		goto L764
	} else {
		goto L779
	}
L779:
	;
	if base.F64_gt(v2899, v2900) != 0 {
		v2908 = int32(1)
		goto L764
	} else {
		goto L780
	}
L780:
	;
	goto L765
L781:
	;
	v3033 = v2776
	goto L734
L782:
	;
	v2918 = *(*int32)(unsafe.Add(mBase, uint32(v2917)+4))
	v2920 = v2918
	goto L784
L783:
	;
	v2920 = int32(0)
	goto L784
L784:
	;
	v2921 = int32(0)
	if base.B2i32(v2920 == v2921)|base.B2i32(v2517 == v2921) != 0 {
		v2967 = base.B2i32(v2920|v2517 == v2921)
		goto L786
	} else {
		goto L787
	}
L785:
	;
	if v2967 != 0 {
		v3033 = v2803
		goto L734
	} else {
		goto L796
	}
L786:
	;
	goto L785
L787:
	;
	v2935 = *(*int32)(unsafe.Add(mBase, uint32(v2920)+4))
	v2936 = *(*int32)(unsafe.Add(mBase, uint32(v2517)+4))
	if v2935 != v2936 {
		v2967 = int32(0)
		goto L786
	} else {
		goto L788
	}
L788:
	;
	v2938 = int32(1)
	if v2935 <= v2938 {
		goto L789
	} else {
		goto L790
	}
L789:
	;
	v2941 = v2938
	goto L791
L790:
	;
	v2941 = v2935
	goto L791
L791:
	;
	v2942 = int32(8)
	v2947 = int32(0)
	goto L792
L792:
	;
	v2955 = v2947 << (uint(int32(2)) % 32)
	v2957 = *(*int32)(unsafe.Add(mBase, uint32(v2920+v2942+v2955)))
	v2959 = *(*int32)(unsafe.Add(mBase, uint32(v2517+v2942+v2955)))
	v2960 = base.B2i32(v2957 == v2959)
	if v2957 != v2959 {
		v2967 = v2960
		goto L786
	} else {
		goto L794
	}
L793:
	;
	v2967 = v2960
	goto L786
L794:
	;
	v2963 = v2947 + int32(1)
	if v2963 != v2941 {
		v2947 = v2963
		goto L792
	} else {
		goto L795
	}
L795:
	;
	goto L793
L796:
	;
	v2973 = F_reparameterize_path(m, l0, v2803, v2517, float64(1))
	mBase = m.M
	v2974 = m.ExcPending
	if v2974 != 0 {
		goto L29
	} else {
		goto L797
	}
L797:
	;
	if v2973 != 0 {
		goto L798
	} else {
		goto L799
	}
L798:
	;
	v2975 = v2973
	goto L800
L799:
	;
	v2975 = v2776
	goto L800
L800:
	;
	v2976 = int32(0)
	if base.B2i32(v2776 == v2976)|base.B2i32(v2973 == v2976) != 0 {
		v3033 = v2975
		goto L734
	} else {
		goto L801
	}
L801:
	;
	v2986 = *(*int32)(unsafe.Add(mBase, uint32(v2776)+40))
	v2987 = *(*int32)(unsafe.Add(mBase, uint32(v2973)+40))
	if v2986 != v2987 {
		goto L803
	} else {
		goto L804
	}
L802:
	;
	if v3029 <= int32(0) {
		goto L826
	} else {
		goto L827
	}
L803:
	;
	if v2986 < v2987 {
		goto L806
	} else {
		goto L807
	}
L804:
	;
	goto L805
L805:
	;
	goto L812
L806:
	;
	v2992 = int32(-1)
	goto L808
L807:
	;
	v2992 = int32(1)
	goto L808
L808:
	;
	v3029 = v2992
	goto L802
L809:
	;
	v3029 = v3023
	goto L802
L810:
	;
	v3023 = int32(0)
	goto L809
L812:
	;
	goto L813
L813:
	;
	v3008 = int32(-1)
	v3009 = *(*float64)(unsafe.Add(mBase, uint32(v2776)+56))
	v3010 = *(*float64)(unsafe.Add(mBase, uint32(v2973)+56))
	if base.F64_lt(v3009, v3010) != 0 {
		v3023 = v3008
		goto L809
	} else {
		goto L820
	}
L820:
	;
	if base.F64_gt(v3009, v3010) != 0 {
		goto L821
	} else {
		goto L822
	}
L821:
	;
	v3029 = int32(1)
	goto L802
L822:
	;
	goto L823
L823:
	;
	v3014 = *(*float64)(unsafe.Add(mBase, uint32(v2776)+48))
	v3015 = *(*float64)(unsafe.Add(mBase, uint32(v2973)+48))
	if base.F64_lt(v3014, v3015) != 0 {
		v3023 = v3008
		goto L809
	} else {
		goto L824
	}
L824:
	;
	if base.F64_gt(v3014, v3015) != 0 {
		v3023 = int32(1)
		goto L809
	} else {
		goto L825
	}
L825:
	;
	goto L810
L826:
	;
	v3032 = v2776
	goto L828
L827:
	;
	v3032 = v2973
	goto L828
L828:
	;
	v3033 = v3032
	goto L734
L829:
	;
	goto L733
L830:
	;
	v3044 = v3033
	goto L729
L831:
	;
	v3084 = v2536 + int32(1)
	v3085 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v3084 < v3085 {
		v2533 = v3082
		v2536 = v3084
		goto L667
	} else {
		goto L842
	}
L832:
	;
	v3080 = F_lappend(m, v2533, v3044)
	mBase = m.M
	v3081 = m.ExcPending
	if v3081 != 0 {
		goto L29
	} else {
		goto L841
	}
L833:
	;
	v3077 = *(*int32)(unsafe.Add(mBase, uint32(v3044)+72))
	v3078 = F_list_concat(m, v2533, v3077)
	mBase = m.M
	v3079 = m.ExcPending
	if v3079 != 0 {
		goto L29
	} else {
		goto L840
	}
L834:
	;
	v3070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3044)+20)))
	if v3070 == int32(1) {
		goto L835
	} else {
		goto L836
	}
L835:
	;
	v3073 = *(*int32)(unsafe.Add(mBase, uint32(v3044)+76))
	if v3073 != 0 {
		goto L832
	} else {
		goto L838
	}
L836:
	;
	goto L837
L837:
	;
	v3074 = *(*int32)(unsafe.Add(mBase, uint32(v3044)+72))
	v3075 = F_list_concat(m, v2533, v3074)
	mBase = m.M
	v3076 = m.ExcPending
	if v3076 != 0 {
		goto L29
	} else {
		goto L839
	}
L838:
	;
	goto L837
L839:
	;
	v3082 = v3075
	goto L831
L840:
	;
	v3082 = v3078
	goto L831
L841:
	;
	v3082 = v3080
	goto L831
L842:
	;
	goto L668
L843:
	;
	F_add_path(m, l1, v3118)
	mBase = m.M
	v3121 = m.ExcPending
	if v3121 != 0 {
		goto L29
	} else {
		goto L844
	}
L844:
	;
	goto L663
L845:
	;
	goto L662
L846:
	;
	m.G0 = v29 + int32(32)
	return
L847:
	;
	v3180 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v3180 != int32(1) {
		goto L846
	} else {
		goto L848
	}
L848:
	;
	v3183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v3184 = *(*int32)(unsafe.Add(mBase, uint32(v3183)))
	v3185 = *(*int32)(unsafe.Add(mBase, uint32(v3184)+40))
	if v3185 == int32(0) {
		goto L846
	} else {
		goto L849
	}
L849:
	;
	v3188 = *(*int32)(unsafe.Add(mBase, uint32(v3185)+4))
	if v3188 < int32(2) {
		goto L846
	} else {
		goto L850
	}
L850:
	;
	v3195 = int32(1)
	goto L851
L851:
	;
	v3218 = *(*int32)(unsafe.Add(mBase, uint32(v3185)+12))
	v3222 = *(*int32)(unsafe.Add(mBase, uint32(v3218+v3195<<(uint(int32(2))%32))))
	v3223 = *(*int32)(unsafe.Add(mBase, uint32(v3222)+64))
	if v3223 != 0 {
		goto L853
	} else {
		goto L854
	}
L852:
	;
	goto L846
L853:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v3222
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v3222
	v3230 = F_list_make1_impl(m, int32(1), v29+int32(12))
	mBase = m.M
	v3231 = m.ExcPending
	if v3231 != 0 {
		goto L29
	} else {
		goto L856
	}
L854:
	;
	goto L855
L855:
	;
	v3241 = v3195 + int32(1)
	v3242 = *(*int32)(unsafe.Add(mBase, uint32(v3185)+4))
	if v3241 < v3242 {
		v3195 = v3241
		goto L851
	} else {
		goto L859
	}
L856:
	;
	v3232 = int32(0)
	v3234 = *(*int32)(unsafe.Add(mBase, uint32(v3222)+24))
	v3236 = F_create_append_path(m, l0, l1, int32(0), v3230, v3232, v3232, v3234, int32(1), v1066)
	mBase = m.M
	v3237 = m.ExcPending
	if v3237 != 0 {
		goto L29
	} else {
		goto L857
	}
L857:
	;
	F_add_partial_path(m, l1, v3236)
	mBase = m.M
	v3239 = m.ExcPending
	if v3239 != 0 {
		goto L29
	} else {
		goto L858
	}
L858:
	;
	goto L855
L859:
	;
	goto L852
}
func F_add_reloption_kind(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_add_reloption_kind[0]))
	if base.Ui32(int32(1073741824)) <= base.Ui32(v3) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_add_reloption_kind_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_add_reloption_kind_1), int32(700), int32(_a_F_add_reloption_kind_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v26 = v3 << (uint(int32(1)) % 32)
		*(*int32)(unsafe.Add(mBase, _c_F_add_reloption_kind[0])) = v26
		return v26
	}
}
func F_add_security_quals(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	v6 = int32(0)
	if l1 == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v131 = int32(0)
	v132 = int32(1)
	v136 = F_makeConst(m, int32(16), int32(-1), v131, v132, v131, v131, v132)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L11
	} else {
		goto L35
	}
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v12 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = v6
	v22 = v6
	goto L6
L4:
	;
	v50 = v6
	goto L5
L5:
	;
	if v50 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v20<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v29 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v50 = v38
	goto L5
L8:
	;
	v30 = F_copyObjectImpl(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v38 = v22
	goto L10
L10:
	;
	v40 = v20 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v40 < v41 {
		v20 = v40
		v22 = v38
		goto L6
	} else {
		goto L14
	}
L11:
	;
	return
L12:
	;
	v32 = F_lappend(m, v22, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+24)))
	v36 = v34 | v35
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v36)
	v38 = v32
	goto L10
L14:
	;
	goto L7
L15:
	;
	if l2 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v102 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L17:
	;
	v56 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v57 <= v56 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v65 = v56
	goto L19
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v65<<(uint(int32(2))%32))))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	if v74 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L16
L21:
	;
	v75 = F_copyObjectImpl(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L11
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v90 = v65 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v90 < v91 {
		v65 = v90
		goto L19
	} else {
		goto L27
	}
L24:
	;
	F_ChangeVarNodes(m, v75, int32(1), l0)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v81 = F_list_append_unique(m, v80, v75)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v81
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+24)))
	v86 = v84 | v85
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v86)
	goto L23
L27:
	;
	goto L20
L28:
	;
	F_ChangeVarNodes(m, v111, int32(1), l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L11
	} else {
		goto L33
	}
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v111 = v106
	goto L28
L30:
	;
	goto L31
L31:
	;
	v109 = F_makeBoolExpr(m, int32(1), v50, int32(-1))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v111 = v109
	goto L28
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v116 = F_list_append_unique(m, v115, v111)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v116
	return
L35:
	;
	v138 = F_lappend(m, v128, v136)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v138
	return
}
func F_adjust_relid_set(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	if l1 < int32(0) {
		v20 = l0
		return v20
	} else {
		v6 = F_bms_is_member(m, l1, l0)
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			if v6 == int32(0) {
				v20 = l0
				return v20
			} else {
				v12 = F_bms_copy(m, l0)
				v13 = m.ExcPending
				if v13 != 0 {
					return int32(0)
				} else {
					v14 = F_bms_del_member(m, v12, l1)
					v15 = m.ExcPending
					if v15 != 0 {
						return int32(0)
					} else {
						if l2 < int32(0) {
							v20 = v14
							return v20
						} else {
							v18 = F_bms_add_member(m, v14, l2)
							v19 = m.ExcPending
							if v19 != 0 {
								return int32(0)
							} else {
								v20 = v18
								return v20
							}
						}
					}
				}
			}
		}
	}
}
func F_alen_object_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+32))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_alen_object_start_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_alen_object_start_1), int32(1908), int32(_a_F_alen_object_start_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		return int32(0)
	}
}
func F_anychar_typmodin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v10 = F_ArrayGetIntegerTypmods(m, l0, v6+int32(28))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
		if v14 == int32(1) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v17 <= int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
						F_errmsg(m, int32(_a_F_anychar_typmodin_0), v6)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_anychar_typmodin_1), int32(53), int32(_a_F_anychar_typmodin_2))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if base.Ui32(int32(_a_F_anychar_typmodin_3)) <= base.Ui32(v17) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(_a_F_anychar_typmodin_4)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l1
							F_errmsg(m, int32(_a_F_anychar_typmodin_5), v6+int32(16))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_anychar_typmodin_1), int32(58), int32(_a_F_anychar_typmodin_2))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					m.G0 = v6 + int32(32)
					return v17 + int32(4)
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_anychar_typmodin_6), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_anychar_typmodin_1), int32(48), int32(_a_F_anychar_typmodin_2))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
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
func F_anyrange_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_anyrange_in_0), int32(207), int32(_a_F_anyrange_in_1), int32(_a_F_anyrange_in_2), int32(_a_F_anyrange_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_appendCSVLiteral(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 <= v6+int32(1) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v27 = l1
	goto L10
L5:
	;
	F_appendStringInfoChar(m, l0, int32(34))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v13+v6))) = uint8(v15)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = v17 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21+v19))) = uint8(v23)
	goto L4
L8:
	;
	return
L9:
	;
	goto L4
L10:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v30 != int32(34) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v86 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v84+v34))) = uint8(v86)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v90 = v88 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v94 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v92+v90))) = uint8(v94)
	goto L3
L12:
	;
	goto L11
L13:
	;
	v62 = int32(1)
	v64 = base.I32_extend8_s(v30)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v65 <= v66+v62 {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	if v30 != 0 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v41 <= v42+int32(1) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v34+int32(1) < v33 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	F_appendStringInfoChar(m, l0, int32(34))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	return
L20:
	;
	F_appendStringInfoChar(m, l0, int32(34))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v49+v42))) = uint8(v51)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v55 = v53 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v57+v55))) = uint8(v59)
	goto L13
L23:
	;
	goto L13
L24:
	;
	F_appendStringInfoChar(m, l0, v64)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L27
	}
L25:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v72+v66))) = uint8(v64)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v77 = v75 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v81 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v79+v77))) = uint8(v81)
	goto L26
L26:
	;
	v27 = v27 + v62
	goto L10
L27:
	;
	goto L26
}
func F_applyRelabelType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	v8 = int32(0)
	if l0 == v8 {
		v45 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v47 = F_exprType(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L14
	} else {
		goto L17
	}
L2:
	;
	v12 = l0
	goto L3
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v21 != int32(27) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if l6 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	goto L4
L6:
	;
	if v21 == int32(7) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v26 != 0 {
		v12 = v26
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v45 = v12
	goto L1
L10:
	;
	v45 = v8
	goto L1
L11:
	;
	v29 = F_copyObjectImpl(m, v12)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v33 = v12
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = l1
	return v33
L14:
	;
	return int32(0)
L15:
	;
	v33 = v29
	goto L13
L16:
	;
	v58 = F_palloc0(m, int32(28))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L14
	} else {
		goto L23
	}
L17:
	;
	if v47 != l1 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v50 = F_exprTypmod(m, v45)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	if v50 != l2 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v53 = F_exprCollation(m, v45)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	if v53 != l3 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	return v45
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(27)
	return v58
}
func F_apply_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v69 int64
	_ = v69
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v108 int32
	_ = v108
	var v111 int64
	_ = v111
	var v114 int64
	_ = v114
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	var v150 int64
	_ = v150
	var v153 int64
	_ = v153
	var v163 int32
	_ = v163
	v9 = m.G0
	v11 = v9 - int32(192)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_apply_error_callback[0]))
	if v14 == int32(0) {
		m.G0 = v11 + int32(192)
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_apply_error_callback[1]))
		if v18 == int32(0) {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_apply_error_callback[2]))
			if v22 == int32(0) {
				F_set_errcontext_domain(m, int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, _c_F_apply_error_callback[3]))
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_apply_error_callback[0]))
					v32 = F_logicalrep_message_type(m, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v29
						F_errcontext_msg(m, int32(_a_F_apply_error_callback_0), v11)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							m.G0 = v11 + int32(192)
							return
						}
					}
				}
			} else {
				v40 = *(*int64)(unsafe.Add(mBase, _c_F_apply_error_callback[4]))
				F_set_errcontext_domain(m, int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, _c_F_apply_error_callback[3]))
					v47 = *(*int32)(unsafe.Add(mBase, _c_F_apply_error_callback[0]))
					v48 = F_logicalrep_message_type(m, v47)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, _c_F_apply_error_callback[2]))
						if v40 == int64(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v51
							*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v48
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v45
							F_errcontext_msg(m, int32(_a_F_apply_error_callback_1), v11+int32(16))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								m.G0 = v11 + int32(192)
								return
							}
						} else {
							v63 = *(*int64)(unsafe.Add(mBase, _c_F_apply_error_callback[4]))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+48)) = uint32(v63)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v45
							*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v48
							*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v51
							v69 = int64(base.Ui64(v63) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+44)) = uint32(v69)
							F_errcontext_msg(m, int32(_a_F_apply_error_callback_2), v11+int32(32))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								m.G0 = v11 + int32(192)
								return
							}
						}
					}
				}
			}
		} else {
			v77 = *(*int64)(unsafe.Add(mBase, _c_F_apply_error_callback[4]))
			v79 = *(*int32)(unsafe.Add(mBase, _c_F_apply_error_callback[5]))
			F_set_errcontext_domain(m, int32(0))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return
			} else {
				v84 = *(*int32)(unsafe.Add(mBase, _c_F_apply_error_callback[3]))
				v86 = *(*int32)(unsafe.Add(mBase, _c_F_apply_error_callback[0]))
				v87 = F_logicalrep_message_type(m, v86)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return
				} else {
					v90 = *(*int32)(unsafe.Add(mBase, _c_F_apply_error_callback[1]))
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
					if v79 < int32(0) {
						v96 = *(*int32)(unsafe.Add(mBase, _c_F_apply_error_callback[2]))
						if v77 == int64(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v96
							*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v91
							*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v87
							*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v84
							F_errcontext_msg(m, int32(_a_F_apply_error_callback_3), v11-int32(-64))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return
							} else {
								m.G0 = v11 + int32(192)
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v96
							v111 = *(*int64)(unsafe.Add(mBase, _c_F_apply_error_callback[4]))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+120)) = uint32(v111)
							v114 = int64(base.Ui64(v111) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+116)) = uint32(v114)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v84
							*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = v87
							*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = v91
							F_errcontext_msg(m, int32(_a_F_apply_error_callback_4), v11+int32(96))
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return
							} else {
								m.G0 = v11 + int32(192)
								return
							}
						}
					} else {
						v125 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
						v127 = *(*int32)(unsafe.Add(mBase, _c_F_apply_error_callback[5]))
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v125+v127<<(uint(int32(2))%32))))
						v133 = *(*int32)(unsafe.Add(mBase, _c_F_apply_error_callback[2]))
						if v77 == int64(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+148)) = v133
							*(*int32)(unsafe.Add(mBase, uint32(v11)+144)) = v131
							*(*int32)(unsafe.Add(mBase, uint32(v11)+140)) = v91
							*(*int32)(unsafe.Add(mBase, uint32(v11)+136)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v87
							*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v84
							F_errcontext_msg(m, int32(_a_F_apply_error_callback_5), v11+int32(128))
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return
							} else {
								m.G0 = v11 + int32(192)
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+176)) = v131
							*(*int32)(unsafe.Add(mBase, uint32(v11)+180)) = v133
							v150 = *(*int64)(unsafe.Add(mBase, _c_F_apply_error_callback[4]))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+188)) = uint32(v150)
							v153 = int64(base.Ui64(v150) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+184)) = uint32(v153)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v84
							*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = v87
							*(*int32)(unsafe.Add(mBase, uint32(v11)+168)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(v11)+172)) = v91
							F_errcontext_msg(m, int32(_a_F_apply_error_callback_6), v11+int32(160))
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return
							} else {
								m.G0 = v11 + int32(192)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_apply_pathtarget_labeling_to_tlist(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	v3 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v13 == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L32
	} else {
		goto L47
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L32
	} else {
		goto L44
	}
L3:
	;
	return
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v16 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 <= int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v31 = v3
	goto L7
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v36 = v31 << (uint(int32(2)) % 32)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v36)))
	if v38 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L3
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39+v36)))
	if v41 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v157 = v31 + int32(1)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v157 < v158 {
		v31 = v157
		goto L7
	} else {
		goto L43
	}
L12:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132+v36)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
	if v138 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L13:
	;
	if l0 == int32(0) {
		goto L1
	} else {
		goto L28
	}
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v44 != int32(6) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if l0 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v49 <= int32(0) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v56 = int32(0)
	goto L18
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v52+v56<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v70 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L1
L20:
	;
	v89 = v56 + int32(1)
	if v49 != v89 {
		v56 = v89
		goto L18
	} else {
		goto L27
	}
L21:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v73 != int32(6) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v76 != v77 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+8)))
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)))
	if v79 != v80 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
	if v82 != v83 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	if v85 == v86 {
		v131 = v69
		v132 = v34
		goto L12
	} else {
		goto L26
	}
L26:
	;
	goto L20
L27:
	;
	goto L19
L28:
	;
	v93 = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v94 <= v93 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v99 = v93
	goto L30
L30:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109+v99<<(uint(int32(2))%32))))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	v115 = F_equal(m, v41, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v131 = v113
	v132 = v123
	goto L12
L32:
	;
	return
L33:
	;
	if v115 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v120 = v99 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v120 < v121 {
		v99 = v120
		goto L30
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L31
L37:
	;
	goto L1
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v142
	goto L11
L39:
	;
	v142 = v137
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v138 != v137 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v142 = v138
	goto L38
L43:
	;
	goto L8
L44:
	;
	F_errmsg_internal(m, int32(_a_F_apply_pathtarget_labeling_to_tlist_0), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L32
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_apply_pathtarget_labeling_to_tlist_1), int32(824), int32(_a_F_apply_pathtarget_labeling_to_tlist_2))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L32
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	F_errmsg_internal(m, int32(_a_F_apply_pathtarget_labeling_to_tlist_3), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L32
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_apply_pathtarget_labeling_to_tlist_1), int32(821), int32(_a_F_apply_pathtarget_labeling_to_tlist_2))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L32
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_arrayconst_next_fn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+68))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)+72))
	if v4 < v5 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+76))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7+v4<<(uint(int32(2))%32))))
		*(*int32)(unsafe.Add(mBase, uint32(v3)+56)) = v11
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v3)+80))
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v4))))
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+60)) = uint8(v15)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+68)) = v4 + int32(1)
		v21 = v3
	} else {
		v21 = int32(0)
	}
	return v21
}
func F_arrayconst_startup_fn(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_palloc(m, int32(84))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v12
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
		v19 = F_pg_detoast_datum(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
			F_get_typlenbyvalalign(m, v21, v9+int32(14), v9+int32(13), v9+int32(12))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v31 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+14)))
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+13)))
				v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9)+12)))
				F_deconstruct_array(m, v19, v31, v32, v33, v12+int32(76), v12+int32(80), v12+int32(72))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(17)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v44
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v47 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v47
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v47)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(16)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v46
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v54
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v57 = F_list_copy(m, v56)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(7)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v57
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v62
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v66
						v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+14)))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v68
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+13)))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+61)) = uint8(v70)
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v12 + int32(36)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = int32(0)
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_arraycontjoinsel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 float64
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v4 == int32(2750) {
		v7 = float64(0.01)
	} else {
		v7 = float64(0.005)
	}
	v8 = F_Float8GetDatum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_arrayexpr_next_fn(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	if v5 == int32(0) {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
		v16 = v14 + int32(4)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
		if base.Ui32(v16) < base.Ui32(v19+v20<<(uint(int32(2))%32)) {
			v25 = v16
		} else {
			v25 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v4)+36)) = v25
		return v4
	}
}
func F_ascii(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_ascii[0]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		if v14 == int32(1) {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
			if base.Ui32((v17-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v45 = int32(1)
				if v14&v45 != 0 {
					v49 = v45
				} else {
					v49 = int32(4)
				}
				v50 = v7 + v49
				if v13 != int32(6) {
					if base.Ui32(v13) <= base.Ui32(int32(41)) {
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v13*int32(28))+uint32(_c_F_ascii[1])))
						v99 = v97
					} else {
						v99 = int32(1)
					}
					v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
					if base.B2i32(v99 < int32(2))|base.B2i32(int32(0) <= v100) != 0 {
						v125 = v100 & int32(255)
						return v125
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_ascii_0), int32(0))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ascii_1), int32(1001), int32(_a_F_ascii_2))
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
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
					v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
					if int32(0) <= v53 {
						if base.Ui32(v13) <= base.Ui32(int32(41)) {
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v13*int32(28))+uint32(_c_F_ascii[1])))
							v99 = v97
						} else {
							v99 = int32(1)
						}
						v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
						if base.B2i32(v99 < int32(2))|base.B2i32(int32(0) <= v100) != 0 {
							v125 = v100 & int32(255)
							return v125
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_ascii_0), int32(0))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ascii_1), int32(1001), int32(_a_F_ascii_2))
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
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
						v60 = base.B2i32(base.Ui32(int32(-33)) < base.Ui32(v53))
						if base.Ui32(int32(-33)) < base.Ui32(v53) {
							v61 = int32(2)
						} else {
							v61 = int32(1)
						}
						v63 = base.B2i32(base.Ui32(int32(-17)) < base.Ui32(v53))
						if base.Ui32(int32(-17)) < base.Ui32(v53) {
							v64 = int32(3)
						} else {
							v64 = v61
						}
						if base.Ui32(int32(-33)) < base.Ui32(v53) {
							v70 = int32(15)
						} else {
							v70 = int32(31)
						}
						if base.Ui32(int32(-17)) < base.Ui32(v53) {
							v71 = int32(7)
						} else {
							v71 = v70
						}
						v75 = int32(1)
						v76 = v53 & int32(255) & v71
						v77 = int32(0)
						for {
							v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v50))))
							v86 = v81&int32(63) | v76<<(uint(int32(6))%32)
							v87 = int32(1)
							v90 = v77 + v87
							if v90 != v64 {
								v75 = v75 + v87
								v76 = v86
								v77 = v90
								continue
							} else {
								break
							}
							break
						}
						v125 = v86
						return v125
					}
				}
			} else {
				if v17 == int32(18) {
					v28 = int32(16)
				} else {
					v28 = int32(0)
				}
				v41 = v28
				if v41 != 0 {
					v45 = int32(1)
					if v14&v45 != 0 {
						v49 = v45
					} else {
						v49 = int32(4)
					}
					v50 = v7 + v49
					if v13 != int32(6) {
						if base.Ui32(v13) <= base.Ui32(int32(41)) {
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v13*int32(28))+uint32(_c_F_ascii[1])))
							v99 = v97
						} else {
							v99 = int32(1)
						}
						v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
						if base.B2i32(v99 < int32(2))|base.B2i32(int32(0) <= v100) != 0 {
							v125 = v100 & int32(255)
							return v125
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_ascii_0), int32(0))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ascii_1), int32(1001), int32(_a_F_ascii_2))
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
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
						v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
						if int32(0) <= v53 {
							if base.Ui32(v13) <= base.Ui32(int32(41)) {
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v13*int32(28))+uint32(_c_F_ascii[1])))
								v99 = v97
							} else {
								v99 = int32(1)
							}
							v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
							if base.B2i32(v99 < int32(2))|base.B2i32(int32(0) <= v100) != 0 {
								v125 = v100 & int32(255)
								return v125
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(261))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_ascii_0), int32(0))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_ascii_1), int32(1001), int32(_a_F_ascii_2))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
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
							v60 = base.B2i32(base.Ui32(int32(-33)) < base.Ui32(v53))
							if base.Ui32(int32(-33)) < base.Ui32(v53) {
								v61 = int32(2)
							} else {
								v61 = int32(1)
							}
							v63 = base.B2i32(base.Ui32(int32(-17)) < base.Ui32(v53))
							if base.Ui32(int32(-17)) < base.Ui32(v53) {
								v64 = int32(3)
							} else {
								v64 = v61
							}
							if base.Ui32(int32(-33)) < base.Ui32(v53) {
								v70 = int32(15)
							} else {
								v70 = int32(31)
							}
							if base.Ui32(int32(-17)) < base.Ui32(v53) {
								v71 = int32(7)
							} else {
								v71 = v70
							}
							v75 = int32(1)
							v76 = v53 & int32(255) & v71
							v77 = int32(0)
							for {
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v50))))
								v86 = v81&int32(63) | v76<<(uint(int32(6))%32)
								v87 = int32(1)
								v90 = v77 + v87
								if v90 != v64 {
									v75 = v75 + v87
									v76 = v86
									v77 = v90
									continue
								} else {
									break
								}
								break
							}
							v125 = v86
							return v125
						}
					}
				} else {
					return int32(0)
				}
			}
		} else {
			v29 = int32(1)
			if v14&v29 != 0 {
				v41 = int32(base.Ui32(v14)>>(uint(v29)%32)) - v29
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v41 = int32(base.Ui32(v35)>>(uint(int32(2))%32)) - int32(4)
			}
			if v41 != 0 {
				v45 = int32(1)
				if v14&v45 != 0 {
					v49 = v45
				} else {
					v49 = int32(4)
				}
				v50 = v7 + v49
				if v13 != int32(6) {
					if base.Ui32(v13) <= base.Ui32(int32(41)) {
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v13*int32(28))+uint32(_c_F_ascii[1])))
						v99 = v97
					} else {
						v99 = int32(1)
					}
					v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
					if base.B2i32(v99 < int32(2))|base.B2i32(int32(0) <= v100) != 0 {
						v125 = v100 & int32(255)
						return v125
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_ascii_0), int32(0))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ascii_1), int32(1001), int32(_a_F_ascii_2))
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
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
					v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
					if int32(0) <= v53 {
						if base.Ui32(v13) <= base.Ui32(int32(41)) {
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v13*int32(28))+uint32(_c_F_ascii[1])))
							v99 = v97
						} else {
							v99 = int32(1)
						}
						v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
						if base.B2i32(v99 < int32(2))|base.B2i32(int32(0) <= v100) != 0 {
							v125 = v100 & int32(255)
							return v125
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_ascii_0), int32(0))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ascii_1), int32(1001), int32(_a_F_ascii_2))
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
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
						v60 = base.B2i32(base.Ui32(int32(-33)) < base.Ui32(v53))
						if base.Ui32(int32(-33)) < base.Ui32(v53) {
							v61 = int32(2)
						} else {
							v61 = int32(1)
						}
						v63 = base.B2i32(base.Ui32(int32(-17)) < base.Ui32(v53))
						if base.Ui32(int32(-17)) < base.Ui32(v53) {
							v64 = int32(3)
						} else {
							v64 = v61
						}
						if base.Ui32(int32(-33)) < base.Ui32(v53) {
							v70 = int32(15)
						} else {
							v70 = int32(31)
						}
						if base.Ui32(int32(-17)) < base.Ui32(v53) {
							v71 = int32(7)
						} else {
							v71 = v70
						}
						v75 = int32(1)
						v76 = v53 & int32(255) & v71
						v77 = int32(0)
						for {
							v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v50))))
							v86 = v81&int32(63) | v76<<(uint(int32(6))%32)
							v87 = int32(1)
							v90 = v77 + v87
							if v90 != v64 {
								v75 = v75 + v87
								v76 = v86
								v77 = v90
								continue
							} else {
								break
							}
							break
						}
						v125 = v86
						return v125
					}
				}
			} else {
				return int32(0)
			}
		}
	}
}
func F_assign_io_max_combine_limit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_assign_io_max_combine_limit[0]))
	if l0 < v5 {
		v7 = l0
	} else {
		v7 = v5
	}
	*(*int32)(unsafe.Add(mBase, _c_F_assign_io_max_combine_limit[1])) = v7
	return
}
func F_assign_random_seed(m *base.Module, l0 float64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v3 != 0 {
		v6 = F_Float8GetDatum(m, l0)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = F_DirectFunctionCall1Coll(m, int32(586), int32(0), v6)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
				return
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
		return
	}
}
func F_autoprewarm_database_main(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int64
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v18 = int32(295)
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	switch int32(297) {
	case 0, 2:
		v32 = v18
		goto L2
	default:
		goto L3
	}
L1:
	;
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v32
	F_sigemptyset(m, v22+int32(16))
	mBase = m.M
	goto L5
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[0])) = v18
	v32 = int32(_a_F_autoprewarm_database_main_0)
	goto L2
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = int32(268435456)
	v46 = F___sigaction(m, int32(15), v22+int32(12), int32(0))
	mBase = m.M
	m.G0 = v22 + int32(32)
	goto L1
L7:
	;
	return
L8:
	;
	v55 = F_GetNamedDSMSegment(m, v15+int32(24))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1])) = v55
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55))))
	F_LWLockRegisterTranche(m, v58, int32(_a_F_autoprewarm_database_main_1))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
	v65 = F_dsm_attach(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	if v65 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
	v70 = int32(0)
	F_BackgroundWorkerInitializeConnectionByOid(m, v69, v70, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L7
	} else {
		goto L93
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+32))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v76)+36))
	if v78 <= v77 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_dsm_detach(m, v65)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L7
	} else {
		goto L92
	}
L17:
	;
	v82 = v74 + v77*int32(20)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v86 = v84
	v87 = v83
	v88 = v85
	v89 = v77
	goto L18
L18:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[2]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	goto L20
L19:
	;
	goto L16
L20:
	;
	if int32(base.Ui32(v100^int32(-1))>>(uint(int32(31))%32)) == int32(0) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v109 = F_RelidByRelfilenumber(m, v87, v86)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L30
	}
L23:
	;
	if v336 < v341 {
		v86 = v333
		v87 = v334
		v88 = v335
		v89 = v336
		goto L18
	} else {
		goto L91
	}
L24:
	;
	F_relation_close(m, v114, int32(1))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L7
	} else {
		goto L89
	}
L25:
	;
	v313 = v86
	v314 = v87
	v315 = v303
	v316 = v304
	goto L24
L26:
	;
	v313 = v86
	v314 = v87
	v315 = v154
	v316 = v207
	goto L24
L27:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	v333 = v144
	v334 = v145
	v335 = v300
	v336 = v132
	v341 = v127
	goto L23
L28:
	;
	v154 = v88
	v155 = v89
	goto L43
L29:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L35
	}
L30:
	;
	if v109 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v114 = F_try_relation_open(m, v109, int32(1))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	if v114 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+36))
	if v89 < v120 {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	v303 = v88
	v304 = v89
	goto L25
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+36))
	if v127 <= v89 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v333 = v86
	v334 = v87
	v335 = v88
	v336 = v89
	v341 = v127
	goto L23
L37:
	;
	goto L38
L38:
	;
	v132 = v89
	goto L39
L39:
	;
	v143 = v74 + v132*int32(20)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if base.B2i32(v145 != v87)|base.B2i32(v144 != v86) != 0 {
		goto L27
	} else {
		goto L41
	}
L40:
	;
	goto L16
L41:
	;
	v150 = v132 + int32(1)
	if v150 != v127 {
		v132 = v150
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[2]))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	goto L45
L44:
	;
	v313 = v284
	v314 = v285
	v315 = v286
	v316 = v287
	goto L24
L45:
	;
	if int32(base.Ui32(v166^int32(-1))>>(uint(int32(31))%32)) == int32(0) {
		v303 = v154
		v304 = v155
		goto L25
	} else {
		goto L46
	}
L46:
	;
	if base.Ui32(v154) <= base.Ui32(int32(3)) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	if base.B2i32(v285 != v87)|base.B2i32(v292 <= v287) != 0 {
		v313 = v284
		v314 = v285
		v315 = v286
		v316 = v287
		goto L24
	} else {
		goto L87
	}
L48:
	;
	v233 = F_RelationGetNumberOfBlocksInFork(m, v114, v154)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L7
	} else {
		goto L77
	}
L49:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	if v175 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+36))
	if v207 <= v155 {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v201 = v175
	goto L54
L53:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v114)+20))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v177
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v114)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v179
	v183 = F_smgropen(m, v15+int32(8), v176)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L7
	} else {
		goto L55
	}
L54:
	;
	v202 = F_smgrexists(m, v201, v154)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L60
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+12)) = v183
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183)+72))
	if v187 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v201 = v199
	goto L54
L57:
	;
	v195 = v187
	goto L59
L58:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183)+76))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v183)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v183)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v191
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v183)+72))
	v195 = v193
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+72)) = v195 + int32(1)
	goto L56
L60:
	;
	if v202 != 0 {
		goto L48
	} else {
		goto L61
	}
L61:
	;
	goto L51
L62:
	;
	v284 = v86
	v285 = v87
	v286 = v154
	v287 = v155
	v292 = v207
	goto L47
L63:
	;
	goto L64
L64:
	;
	v212 = v155
	goto L65
L65:
	;
	v223 = v74 + v212*int32(20)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v87 != v226 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v284 = v225
	v285 = v226
	v286 = v224
	v287 = v212
	v292 = v207
	goto L47
L67:
	;
	v284 = v225
	v285 = v226
	v286 = v224
	v287 = v212
	v292 = v207
	goto L47
L68:
	;
	goto L69
L69:
	;
	if v225 != v86 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v284 = v225
	v285 = v226
	v286 = v224
	v287 = v212
	v292 = v207
	goto L47
L71:
	;
	goto L72
L72:
	;
	if v154 == v224 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v231 = v212 + int32(1)
	if v231 == v207 {
		goto L26
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	goto L66
L76:
	;
	v212 = v231
	goto L65
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v74
	v242 = int32(0)
	v247 = F_read_stream_begin_relation(m, int32(9), v242, v114, v154, int32(_a_F_autoprewarm_database_main_2), v15+int32(24), v242)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L7
	} else {
		goto L78
	}
L78:
	;
	goto L79
L79:
	;
	v262 = F_read_stream_next_buffer(m, v247, int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L7
	} else {
		goto L81
	}
L80:
	;
	F_read_stream_end(m, v247)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L7
	} else {
		goto L86
	}
L81:
	;
	if v262 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v265)+40)) = v266 + int32(1)
	F_ReleaseBuffer(m, v262)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L7
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	goto L80
L85:
	;
	goto L79
L86:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+36))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v280 = v74 + v277*int32(20)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)+12))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	v284 = v282
	v285 = v283
	v286 = v281
	v287 = v277
	v292 = v276
	goto L47
L87:
	;
	if v284 == v86 {
		v154 = v286
		v155 = v287
		goto L43
	} else {
		goto L88
	}
L88:
	;
	goto L44
L89:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)+36))
	v333 = v313
	v334 = v314
	v335 = v315
	v336 = v316
	v341 = v332
	goto L23
L91:
	;
	goto L19
L92:
	;
	m.G0 = v15 + int32(48)
	return
L93:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L7
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(_a_F_autoprewarm_database_main_3), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_database_main_4), int32(518), int32(_a_F_autoprewarm_database_main_5))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_autoprewarm_main(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v285 int32
	_ = v285
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v448 int32
	_ = v448
	var v451 int64
	_ = v451
	var v454 int32
	_ = v454
	var v457 int64
	_ = v457
	var v460 int64
	_ = v460
	var v463 int64
	_ = v463
	var v466 int32
	_ = v466
	var v469 int64
	_ = v469
	var v472 int64
	_ = v472
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int64
	_ = v574
	var v575 int64
	_ = v575
	var v588 int32
	_ = v588
	var v596 int64
	_ = v596
	var v599 int32
	_ = v599
	var v612 int64
	_ = v612
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int64
	_ = v641
	var v642 int64
	_ = v642
	var v650 int64
	_ = v650
	var v656 int64
	_ = v656
	var v662 int64
	_ = v662
	var v671 int64
	_ = v671
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int64
	_ = v700
	var v701 int64
	_ = v701
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	v15 = m.G0
	v17 = v15 - int32(1632)
	m.G0 = v17
	v20 = int32(916)
	v22 = m.G0
	v24 = v22 - int32(32)
	m.G0 = v24
	switch int32(918) {
	case 0, 2:
		v34 = v20
		goto L2
	default:
		goto L3
	}
L1:
	;
	v52 = int32(1)
	v54 = int32(914)
	v56 = m.G0
	v58 = v56 - int32(32)
	m.G0 = v58
	switch int32(916) {
	case 0, 2:
		v68 = v54
		goto L8
	default:
		goto L9
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
	F_sigemptyset(m, v24+int32(16))
	mBase = m.M
	goto L5
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[0])) = v20
	v34 = int32(_a_F_autoprewarm_main_0)
	goto L2
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = int32(268435456)
	v48 = F___sigaction(m, int32(15), v24+int32(12), int32(0))
	mBase = m.M
	m.G0 = v24 + int32(32)
	goto L1
L7:
	;
	v87 = int32(917)
	v89 = m.G0
	v91 = v89 - int32(32)
	m.G0 = v91
	switch int32(919) {
	case 0, 2:
		v101 = v87
		goto L14
	default:
		goto L15
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v68
	F_sigemptyset(m, v58+int32(16))
	mBase = m.M
	goto L11
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[1])) = v54
	v68 = int32(_a_F_autoprewarm_main_0)
	goto L8
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = int32(268435456)
	v82 = F___sigaction(m, v52, v58+int32(12), int32(0))
	mBase = m.M
	m.G0 = v58 + int32(32)
	goto L7
L13:
	;
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v101
	F_sigemptyset(m, v91+int32(16))
	mBase = m.M
	goto L17
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[2])) = v87
	v101 = int32(_a_F_autoprewarm_main_0)
	goto L14
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+24)) = int32(268435456)
	v115 = F___sigaction(m, int32(10), v91+int32(12), int32(0))
	mBase = m.M
	m.G0 = v91 + int32(32)
	goto L13
L19:
	;
	return
L20:
	;
	v124 = F_GetNamedDSMSegment(m, v17+int32(168))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3])) = v124
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	F_LWLockRegisterTranche(m, v127, int32(_a_F_autoprewarm_main_1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+168)))
	F_before_shmem_exit(m, int32(_a_F_autoprewarm_main_2), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v139 = F_LWLockAcquire(m, v137, int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
	if v143 != int32(-1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	m.G0 = v17 + int32(1632)
	return
L26:
	;
	F_LWLockRelease(m, v142)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L19
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v142)+16)) = v169
	F_LWLockRelease(m, v142)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L19
	} else {
		goto L34
	}
L29:
	;
	v150 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	if v150 == int32(0) {
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+144)) = v156
	F_errmsg(m, int32(_a_F_autoprewarm_main_3), v17+int32(144))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_main_4), int32(202), int32(_a_F_autoprewarm_main_5))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	goto L25
L34:
	;
	if v131&int32(1) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	if v588 == int32(0) {
		goto L25
	} else {
		goto L147
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L19
	} else {
		goto L143
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L19
	} else {
		goto L140
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L19
	} else {
		goto L135
	}
L39:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v180 = F_LWLockAcquire(m, v178, int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L19
	} else {
		goto L42
	}
L40:
	;
	v588 = v52
	v596 = int64(0)
	goto L41
L41:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[5]))
	if v599 != 0 {
		goto L35
	} else {
		goto L110
	}
L42:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+20))
	if v184 == int32(-1) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[5]))
	v569 = m.G0
	v570 = int32(16)
	v571 = v569 - v570
	m.G0 = v571
	F_gettimeofday(m, v571)
	mBase = m.M
	v574 = *(*int64)(unsafe.Add(mBase, uint32(v571)))
	v575 = int64(*(*int32)(unsafe.Add(mBase, uint32(v571)+8)))
	m.G0 = v571 + v570
	goto L109
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v17 + int32(160)
	v255 = F_fscanf(m, v194, int32(_a_F_autoprewarm_main_6), v17+int32(112))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L19
	} else {
		goto L64
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L19
	} else {
		goto L60
	}
L46:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+20)) = v188
	F_LWLockRelease(m, v183)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L19
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	F_LWLockRelease(m, v183)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L19
	} else {
		goto L55
	}
L49:
	;
	v194 = F_AllocateFile(m, int32(_a_F_autoprewarm_main_7), int32(_a_F_autoprewarm_main_8))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L19
	} else {
		goto L50
	}
L50:
	;
	if v194 != 0 {
		goto L44
	} else {
		goto L51
	}
L51:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[6]))
	if v197 != int32(44) {
		goto L45
	} else {
		goto L52
	}
L52:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v203 = F_LWLockAcquire(m, v201, int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L19
	} else {
		goto L53
	}
L53:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v206)+20)) = int32(-1)
	F_LWLockRelease(m, v206)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L19
	} else {
		goto L54
	}
L54:
	;
	goto L43
L55:
	;
	v215 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L19
	} else {
		goto L56
	}
L56:
	;
	if v215 == int32(0) {
		goto L43
	} else {
		goto L57
	}
L57:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = v221
	F_errmsg(m, int32(_a_F_autoprewarm_main_9), v17+int32(128))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L19
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_main_4), int32(312), int32(_a_F_autoprewarm_main_10))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L19
	} else {
		goto L59
	}
L59:
	;
	goto L43
L60:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L19
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(_a_F_autoprewarm_main_7)
	F_errmsg(m, int32(_a_F_autoprewarm_main_11), v17)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L19
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_main_4), int32(334), int32(_a_F_autoprewarm_main_10))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L19
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	if v255 != int32(1) {
		goto L36
	} else {
		goto L65
	}
L65:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v17)+160))
	v263 = F_dsm_create(m, v259*int32(20), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L19
	} else {
		goto L66
	}
L66:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v263)+24))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v17)+160))
	if int32(0) < v266 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v271 = int32(0)
	goto L70
L68:
	;
	goto L69
L69:
	;
	v326 = F_FreeFile(m, v194)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L19
	} else {
		goto L75
	}
L70:
	;
	v285 = v265 + v271*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v285 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v285 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v285 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v17 + int32(168)
	v302 = F_fscanf(m, v194, int32(_a_F_autoprewarm_main_12), v17-int32(-64))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L19
	} else {
		goto L72
	}
L71:
	;
	goto L69
L72:
	;
	if v302 != int32(5) {
		goto L37
	} else {
		goto L73
	}
L73:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v285)+12)) = v306
	v309 = v271 + int32(1)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v17)+160))
	if v309 < v310 {
		v271 = v309
		goto L70
	} else {
		goto L74
	}
L74:
	;
	goto L71
L75:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v17)+160))
	F_pg_qsort(m, v265, v328, int32(20), int32(_a_F_autoprewarm_main_13))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L19
	} else {
		goto L76
	}
L76:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v334)+24)) = v335
	v337 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v334)+40)) = v337
	*(*int64)(unsafe.Add(mBase, uint32(v334)+32)) = int64(0)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v17)+160))
	if v341 <= v337 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_dsm_detach(m, v263)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L19
	} else {
		goto L101
	}
L78:
	;
	v345 = v17 + int32(264)
	v347 = v17 + int32(1396)
	v349 = v17 + int32(372)
	v355 = int32(0)
	v357 = v341
	v358 = v334
	goto L79
L79:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v265+v355*int32(20))))
	v372 = v355 + int32(1)
	if v372 < v357 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L77
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+28)) = v413
	*(*int32)(unsafe.Add(mBase, uint32(v358)+36)) = v415
	v430 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[7]))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+8))
	goto L94
L82:
	;
	v374 = v370
	v376 = v372
	goto L85
L83:
	;
	v397 = v370
	v399 = v372
	goto L84
L84:
	;
	if v397 == int32(0) {
		goto L77
	} else {
		goto L93
	}
L85:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v265+v376*int32(20))))
	if v391 == v374 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v397 = v393
	v399 = v357
	goto L84
L87:
	;
	v395 = v376 + int32(1)
	if v395 != v357 {
		v374 = v393
		v376 = v395
		goto L85
	} else {
		goto L92
	}
L88:
	;
	v393 = v374
	goto L87
L89:
	;
	goto L90
L90:
	;
	if v374 != 0 {
		v413 = v374
		v415 = v376
		goto L81
	} else {
		goto L91
	}
L91:
	;
	v393 = v391
	goto L87
L92:
	;
	goto L86
L93:
	;
	v413 = v397
	v415 = v399
	goto L81
L94:
	;
	if int32(base.Ui32(v431^int32(-1))>>(uint(int32(31))%32)) == int32(0) {
		goto L77
	} else {
		goto L95
	}
L95:
	;
	v439 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[5]))
	if v439 != 0 {
		goto L77
	} else {
		goto L96
	}
L96:
	;
	base.MemoryFill(m, v17+int32(184), int32(0), int32(1440))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+368)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+360)) = int64(4294967299)
	v448 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+7)) = v448
	v451 = *(*int64)(unsafe.Add(mBase, _c_F_autoprewarm_main[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v349))) = v451
	v454 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_autoprewarm_main[10])))
	*(*uint16)(unsafe.Add(mBase, uint32(v347)+24)) = uint16(v454)
	v457 = *(*int64)(unsafe.Add(mBase, _c_F_autoprewarm_main[11]))
	*(*int64)(unsafe.Add(mBase, uint32(v347)+16)) = v457
	v460 = *(*int64)(unsafe.Add(mBase, _c_F_autoprewarm_main[12]))
	*(*int64)(unsafe.Add(mBase, uint32(v347)+8)) = v460
	v463 = *(*int64)(unsafe.Add(mBase, _c_F_autoprewarm_main[13]))
	*(*int64)(unsafe.Add(mBase, uint32(v347))) = v463
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+183)) = v466
	v469 = *(*int64)(unsafe.Add(mBase, _c_F_autoprewarm_main[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+176)) = v469
	v472 = *(*int64)(unsafe.Add(mBase, _c_F_autoprewarm_main[16]))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+168)) = v472
	*(*int32)(unsafe.Add(mBase, uint32(v345)+15)) = v466
	*(*int64)(unsafe.Add(mBase, uint32(v345)+8)) = v469
	*(*int64)(unsafe.Add(mBase, uint32(v345))) = v472
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+1624)) = v478
	v484 = F_RegisterDynamicBackgroundWorker(m, v17+int32(168), v17+int32(164))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L19
	} else {
		goto L97
	}
L97:
	;
	if v484 == int32(0) {
		goto L38
	} else {
		goto L98
	}
L98:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v17)+164))
	v489 = F_WaitForBackgroundWorkerShutdown(m, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L19
	} else {
		goto L99
	}
L99:
	;
	v492 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v492)+32)) = v493
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v17)+160))
	if v493 < v495 {
		v355 = v493
		v357 = v495
		v358 = v492
		goto L79
	} else {
		goto L100
	}
L100:
	;
	goto L80
L101:
	;
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v516 = F_LWLockAcquire(m, v514, int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L19
	} else {
		goto L102
	}
L102:
	;
	v519 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v519)+20)) = int64(4294967295)
	F_LWLockRelease(m, v519)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L19
	} else {
		goto L103
	}
L103:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[5]))
	if v525 != 0 {
		goto L43
	} else {
		goto L104
	}
L104:
	;
	v528 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L19
	} else {
		goto L105
	}
L105:
	;
	if v528 == int32(0) {
		goto L43
	} else {
		goto L106
	}
L106:
	;
	v533 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v533)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v534
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v17)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v536
	F_errmsg(m, int32(_a_F_autoprewarm_main_14), v17+int32(16))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L19
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_main_4), int32(445), int32(_a_F_autoprewarm_main_10))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L19
	} else {
		goto L108
	}
L108:
	;
	goto L43
L109:
	;
	v588 = base.B2i32(v563 == int32(0))
	v596 = v575 + v574*int64(1000000) - int64(946684800000000)
	goto L41
L110:
	;
	v612 = v596
	goto L111
L111:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[17]))
	if v615 != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	goto L35
L113:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[17])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L19
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v623 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[18]))
	if v623 <= int32(0) {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	goto L115
L117:
	;
	v695 = m.G0
	v696 = int32(16)
	v697 = v695 - v696
	m.G0 = v697
	F_gettimeofday(m, v697)
	mBase = m.M
	v700 = *(*int64)(unsafe.Add(mBase, uint32(v697)))
	v701 = int64(*(*int32)(unsafe.Add(mBase, uint32(v697)+8)))
	m.G0 = v697 + v696
	goto L132
L118:
	;
	v685 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v685))) = int32(0)
	goto L130
L119:
	;
	v627 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[19]))
	v631 = F_WaitLatch(m, v627, int32(33), int32(-1), int32(117440512))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L19
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v636 = m.G0
	v637 = int32(16)
	v638 = v636 - v637
	m.G0 = v638
	F_gettimeofday(m, v638)
	mBase = m.M
	v641 = *(*int64)(unsafe.Add(mBase, uint32(v638)))
	v642 = int64(*(*int32)(unsafe.Add(mBase, uint32(v638)+8)))
	m.G0 = v638 + v637
	v650 = v642 + v641*int64(1000000) - int64(946684800000000)
	goto L123
L122:
	;
	goto L118
L123:
	;
	v656 = base.I64_extend_i32_u(v623*int32(1000))*int64(1000) + v612
	if v656 <= v650 {
		v674 = int32(0)
		goto L125
	} else {
		goto L126
	}
L124:
	;
	if v674 <= int32(0) {
		goto L117
	} else {
		goto L128
	}
L125:
	;
	goto L124
L126:
	;
	v662 = v656 - v650
	if base.B2i32(int64(0) < v650)^base.B2i32(v662 < v656)|base.B2i32(int64(2147483646000) < v662) != 0 {
		v674 = int32(2147483647)
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v671 = base.I64_div_s(v662+int64(999), int64(1000))
	v674 = base.I32_wrap_i64(v671)
	goto L125
L128:
	;
	v678 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[19]))
	v681 = F_WaitLatch(m, v678, int32(41), v674, int32(117440512))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L19
	} else {
		goto L129
	}
L129:
	;
	goto L118
L130:
	;
	v689 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[5]))
	if v689 == int32(0) {
		goto L111
	} else {
		goto L131
	}
L131:
	;
	goto L35
L132:
	;
	v712 = F_apw_dump_now(m, int32(1), int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L19
	} else {
		goto L133
	}
L133:
	;
	v715 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[5]))
	if v715 == int32(0) {
		v612 = v701 + v700*int64(1000000) - int64(946684800000000)
		goto L111
	} else {
		goto L134
	}
L134:
	;
	goto L112
L135:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L19
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_autoprewarm_main_15), int32(0))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L19
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(_a_F_autoprewarm_main_16)
	F_errhint(m, int32(_a_F_autoprewarm_main_17), v17+int32(32))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L19
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_main_4), int32(971), int32(_a_F_autoprewarm_main_18))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L19
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
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v271 + int32(1)
	F_errmsg(m, int32(_a_F_autoprewarm_main_19), v17+int32(48))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L19
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_main_4), int32(358), int32(_a_F_autoprewarm_main_10))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L19
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L19
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = int32(_a_F_autoprewarm_main_7)
	F_errmsg(m, int32(_a_F_autoprewarm_main_20), v17+int32(96))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L19
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_main_4), int32(342), int32(_a_F_autoprewarm_main_10))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L19
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	v792 = int32(1)
	v794 = F_apw_dump_now(m, v792, v792)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L19
	} else {
		goto L148
	}
L148:
	;
	goto L25
}
func F_autoprewarm_start_worker(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_autoprewarm_start_worker[0])))
	if v9 != 0 {
		v13 = F_GetNamedDSMSegment(m, v6+int32(15))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_start_worker[1])) = v13
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
			F_LWLockRegisterTranche(m, v18, int32(_a_F_autoprewarm_start_worker_0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_start_worker[1]))
				v25 = F_LWLockAcquire(m, v23, int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_start_worker[1]))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
					F_LWLockRelease(m, v28)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v29 != int32(-1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v29
									F_errmsg(m, int32(_a_F_autoprewarm_start_worker_1), v6)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_autoprewarm_start_worker_2), int32(833), int32(_a_F_autoprewarm_start_worker_3))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							F_apw_start_leader_worker(m)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v6 + int32(16)
								return int32(0)
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_autoprewarm_start_worker_4), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_autoprewarm_start_worker_2), int32(822), int32(_a_F_autoprewarm_start_worker_3))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
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
func F_avals(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_avals(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
