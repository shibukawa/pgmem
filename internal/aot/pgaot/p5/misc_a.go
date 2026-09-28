package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_AdvanceXLInsertBuffer(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v73 int32
	_ = v73
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int64
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
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
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v200 int64
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v211 int64
	_ = v211
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int64
	_ = v228
	var v230 int32
	_ = v230
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	v20 = F_LWLockAcquire(m, v16+int32(896), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)+280))
	if base.B2i32(l2 == int32(0))&base.B2i32(base.Ui64(l0) < base.Ui64(v26)) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	F_LWLockRelease(m, v245+int32(896))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L34
	}
L4:
	;
	v32 = v25
	v37 = v26
	goto L5
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
	v45 = base.I64_rem_u_s(int64(base.Ui64(v37)>>(uint(int64(13))%64)), base.I64_extend_i32_s(v41+int32(1)))
	v46 = base.I32_wrap_i64(v45)
	v48 = v46 << (uint(int32(3)) % 32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v32)+292))
	v51 = int64(0)
	v54 = base.AtomicRmwCmpxchg64(m, v48+v49, int32(0), v51, v51)
	v56 = *(*int64)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[2]))
	if base.Ui64(v54) <= base.Ui64(v56) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v175)+280))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v175)+288))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175)+292))
	v180 = int64(0)
	v181 = int32(0)
	v182 = base.AtomicRmwXchg64(m, v178+v48, v181, v180)
	v186 = base.AtomicRmwOr32(m, v181, int32(_a_F_AdvanceXLInsertBuffer_0), v181)
	v189 = v177 + v46<<(uint(int32(13))%32)
	base.MemoryFill(m, v189+int32(2), v181, int32(_a_F_AdvanceXLInsertBuffer_1))
	*(*int64)(unsafe.Add(mBase, uint32(v189)+8)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = l1
	v197 = int32(_a_F_AdvanceXLInsertBuffer_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v189))) = uint16(v197)
	v200 = v176 - int64(-8192)
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[3]))
	if v176&base.I64_extend_i32_s(v202-int32(1)) == v180 {
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
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	v62 = base.AtomicRmwXchg32(m, v59, int32(440), int32(1))
	if v62 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_s_lock(m, v59+int32(440), int32(_a_F_AdvanceXLInsertBuffer_3))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v69)+184))
	if base.Ui64(v70) < base.Ui64(v54) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v69)+184)) = v54
	goto L16
L15:
	;
	goto L16
L16:
	;
	v73 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v69)+440)), uint32(v73))
	v77 = int64(0)
	v80 = base.AtomicRmwCmpxchg64(m, v69, int32(272), v77, v77)
	*(*int64)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[4])) = v80
	v85 = base.AtomicRmwOr32(m, v73, int32(_a_F_AdvanceXLInsertBuffer_0), v73)
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	v92 = base.AtomicRmwCmpxchg64(m, v88, int32(264), v77, v77)
	*(*int64)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[2])) = v92
	if base.Ui64(v54) <= base.Ui64(v92) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	F_LWLockRelease(m, v96+int32(896))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v101 = F_WaitXLogInsertionsToFinish(m, v54)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
	v108 = F_LWLockAcquire(m, v104+int32(1024), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v111 = int32(_a_F_AdvanceXLInsertBuffer_4)
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	v113 = int64(0)
	v116 = base.AtomicRmwCmpxchg64(m, v112, int32(272), v113, v113)
	*(*int64)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[4])) = v116
	v118 = int32(0)
	v121 = base.AtomicRmwOr32(m, v118, int32(_a_F_AdvanceXLInsertBuffer_0), v118)
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	v128 = base.AtomicRmwCmpxchg64(m, v124, int32(264), v113, v113)
	*(*int64)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[2])) = v128
	if base.Ui64(v54) <= base.Ui64(v128) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
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
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
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
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v137
	F_XLogWrite(m, v13, l1, int32(0))
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
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[0]))
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
	*(*uint8)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[5])) = uint8(v153)
	v155 = int32(_a_F_AdvanceXLInsertBuffer_5)
	v157 = *(*int64)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[6]))
	*(*int64)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[6])) = v157 + int64(1)
	goto L21
L28:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v169)+280))
	if base.Ui64(v170) <= base.Ui64(l0) {
		v32 = v169
		v37 = v170
		goto L5
	} else {
		goto L29
	}
L29:
	;
	goto L3
L30:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[7]))
	v211 = *(*int64)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+36)) = int32(_a_F_AdvanceXLInsertBuffer_6)
	*(*int32)(unsafe.Add(mBase, uint32(v189)+32)) = v202
	*(*int64)(unsafe.Add(mBase, uint32(v189)+24)) = v211
	v216 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v189)+2)) = uint16(v216)
	goto L32
L31:
	;
	goto L32
L32:
	;
	v219 = int32(0)
	v222 = base.AtomicRmwOr32(m, v219, int32(_a_F_AdvanceXLInsertBuffer_0), v219)
	v223 = int32(_a_F_AdvanceXLInsertBuffer_4)
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+292))
	v228 = base.AtomicRmwXchg64(m, v225+v48, v219, v200)
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceXLInsertBuffer[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v230)+280)) = v200
	if l2|base.B2i32(base.Ui64(v200) <= base.Ui64(l0)) != 0 {
		v32 = v230
		v37 = v200
		goto L5
	} else {
		goto L33
	}
L33:
	;
	goto L6
L34:
	;
	m.G0 = v13 + int32(32)
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
	var v100 int32
	_ = v100
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
									v100 = F_errdetail(m, int32(_a_F_AllocSetContextCreateInternal_3), v12)
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
								v100 = F_errdetail(m, int32(_a_F_AllocSetContextCreateInternal_3), v12)
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
				*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(482)
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
										v100 = F_errdetail(m, int32(_a_F_AllocSetContextCreateInternal_3), v12)
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
									v100 = F_errdetail(m, int32(_a_F_AllocSetContextCreateInternal_3), v12)
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
					*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(482)
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
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(482)
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
									v100 = F_errdetail(m, int32(_a_F_AllocSetContextCreateInternal_3), v12)
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
								v100 = F_errdetail(m, int32(_a_F_AllocSetContextCreateInternal_3), v12)
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
				*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(482)
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
			*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(482)
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
					F_errfinish(m, int32(_a_F_AllocSetFree_1), int32(1125), int32(_a_F_AllocSetFree_2))
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
			if v25 != int32(482) {
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
						F_errfinish(m, int32(_a_F_AllocSetFree_1), int32(1125), int32(_a_F_AllocSetFree_2))
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
							F_errfinish(m, int32(_a_F_AllocSetFree_1), int32(1125), int32(_a_F_AllocSetFree_2))
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
											F_errfinish(m, int32(_a_F_AlterEventTriggerOwner_internal_2), int32(560), int32(_a_F_AlterEventTriggerOwner_internal_3))
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
										F_errfinish(m, int32(_a_F_AlterEventTriggerOwner_internal_2), int32(560), int32(_a_F_AlterEventTriggerOwner_internal_3))
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
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
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
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v709 int32
	_ = v709
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v376 = v11
	} else {
		v12 = int32(4)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(v14-int32(1021)) < base.Ui32(v12) {
			v23 = v14
			v24 = v12
			v27 = l0 + int32(28)
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v23) {
					v30 = int32(1024)
					v37 = int32(-1636607408)
					if v13&int32(3) != 0 {
						v83 = v13
						v84 = v30
						v86 = v37
						v87 = v37
						v88 = v37
						for {
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
								continue
							} else {
								break
							}
							break
						}
						switch v129 - int32(1) {
						case 0:
							v302 = v120
							v303 = v121
							v304 = v125
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 1:
							v295 = v120
							v296 = v121
							v297 = v125
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 2:
							v288 = v120
							v289 = v121
							v290 = v125
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 3:
							v282 = v121
							v283 = v125
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 4:
							v278 = v121
							v279 = v125
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 5:
							v272 = v121
							v273 = v125
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 6:
							v266 = v121
							v267 = v125
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
							v272 = v268<<(uint(int32(16))%32) + v266
							v273 = v267
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 7:
							v261 = v125
							v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+7)))
							v266 = v262<<(uint(int32(24))%32) + v121
							v267 = v261
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
							v272 = v268<<(uint(int32(16))%32) + v266
							v273 = v267
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 8:
							v256 = v125
							v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+8)))
							v261 = v257<<(uint(int32(8))%32) + v256
							v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+7)))
							v266 = v262<<(uint(int32(24))%32) + v121
							v267 = v261
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
							v272 = v268<<(uint(int32(16))%32) + v266
							v273 = v267
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 9:
							v251 = v125
							v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+9)))
							v256 = v252<<(uint(int32(16))%32) + v251
							v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+8)))
							v261 = v257<<(uint(int32(8))%32) + v256
							v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+7)))
							v266 = v262<<(uint(int32(24))%32) + v121
							v267 = v261
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
							v272 = v268<<(uint(int32(16))%32) + v266
							v273 = v267
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 10:
							v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+10)))
							v251 = v247<<(uint(int32(24))%32) + v125
							v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+9)))
							v256 = v252<<(uint(int32(16))%32) + v251
							v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+8)))
							v261 = v257<<(uint(int32(8))%32) + v256
							v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+7)))
							v266 = v262<<(uint(int32(24))%32) + v121
							v267 = v261
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
							v272 = v268<<(uint(int32(16))%32) + v266
							v273 = v267
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						default:
							v310 = v120
							v311 = v121
							v312 = v125
						}
					} else {
						v143 = v13
						v144 = v30
						v146 = v37
						v147 = v37
						v148 = v37
						for {
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
								continue
							} else {
								break
							}
							break
						}
						switch v189 - int32(1) {
						case 0:
							v244 = v180
							v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
							v310 = v244 + v245
							v311 = v181
							v312 = v185
						case 1:
							v239 = v180
							v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
							v244 = v240<<(uint(int32(8))%32) + v239
							v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
							v310 = v244 + v245
							v311 = v181
							v312 = v185
						case 2:
							v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
							v239 = v235<<(uint(int32(16))%32) + v180
							v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
							v244 = v240<<(uint(int32(8))%32) + v239
							v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
							v310 = v244 + v245
							v311 = v181
							v312 = v185
						case 3:
							v232 = v181
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v310 = v233 + v180
							v311 = v232
							v312 = v185
						case 4:
							v229 = v181
							v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
							v232 = v229 + v230
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v310 = v233 + v180
							v311 = v232
							v312 = v185
						case 5:
							v224 = v181
							v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
							v229 = v225<<(uint(int32(8))%32) + v224
							v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
							v232 = v229 + v230
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v310 = v233 + v180
							v311 = v232
							v312 = v185
						case 6:
							v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+6)))
							v224 = v220<<(uint(int32(16))%32) + v181
							v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
							v229 = v225<<(uint(int32(8))%32) + v224
							v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
							v232 = v229 + v230
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v310 = v233 + v180
							v311 = v232
							v312 = v185
						case 7:
							v215 = v185
							v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
							v310 = v216 + v180
							v311 = v218 + v181
							v312 = v215
						case 8:
							v210 = v185
							v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
							v215 = v211<<(uint(int32(8))%32) + v210
							v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
							v310 = v216 + v180
							v311 = v218 + v181
							v312 = v215
						case 9:
							v205 = v185
							v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+9)))
							v210 = v206<<(uint(int32(16))%32) + v205
							v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
							v215 = v211<<(uint(int32(8))%32) + v210
							v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
							v310 = v216 + v180
							v311 = v218 + v181
							v312 = v215
						case 10:
							v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+10)))
							v205 = v201<<(uint(int32(24))%32) + v185
							v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+9)))
							v210 = v206<<(uint(int32(16))%32) + v205
							v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
							v215 = v211<<(uint(int32(8))%32) + v210
							v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
							v310 = v216 + v180
							v311 = v218 + v181
							v312 = v215
						default:
							v310 = v180
							v311 = v181
							v312 = v185
						}
					}
					v315 = int32(14)
					v317 = v311 ^ v312 - base.I32_rotl(v311, v315)
					v321 = v317 ^ v310 - base.I32_rotl(v317, int32(11))
					v325 = v321 ^ v311 - base.I32_rotl(v321, int32(25))
					v329 = v325 ^ v317 - base.I32_rotl(v325, int32(16))
					v333 = v329 ^ v321 - base.I32_rotl(v329, int32(4))
					v337 = v333 ^ v325 - base.I32_rotl(v333, v315)
					*(*int64)(unsafe.Add(mBase, uint32(v13))) = base.I64_extend_i32_u(v337)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v337^v329-base.I32_rotl(v337, int32(24)))
					v349 = int32(8)
				} else {
					v349 = v23
				}
				v351 = int32(1024) - v349
				if base.Ui32(v24) < base.Ui32(v351) {
					v353 = v24
				} else {
					v353 = v351
				}
				if v353 != 0 {
					base.MemoryCopy(m, v349+v13, v27, v353)
				} else {
				}
				v357 = v349 + v353
				v358 = v24 - v353
				if v358 != 0 {
					v23 = v357
					v24 = v358
					v27 = v353 + v27
					continue
				} else {
					break
				}
				break
			}
			v366 = v357
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14+v13))) = v8
			v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v366 = v361 + int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v366
		v376 = v366
	}
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v376 != int32(1024) {
		v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		*(*uint8)(unsafe.Add(mBase, uint32(v376+v381))) = uint8(v385)
		v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v387 + int32(1)
		return
	} else {
		v391 = int32(1024)
		v398 = int32(-1636607408)
		if v381&int32(3) != 0 {
			v444 = v381
			v445 = v391
			v447 = v398
			v448 = v398
			v449 = v398
			for {
				v451 = *(*int32)(unsafe.Add(mBase, uint32(v444)+4))
				v452 = v451 + v448
				v453 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
				v455 = *(*int32)(unsafe.Add(mBase, uint32(v444)+8))
				v456 = v455 + v449
				v458 = int32(4)
				v460 = v453 + v447 - v456 ^ base.I32_rotl(v456, v458)
				v464 = v452 - v460 ^ base.I32_rotl(v460, int32(6))
				v465 = v456 + v452
				v466 = v460 + v465
				v467 = v464 + v466
				v471 = v465 - v464 ^ base.I32_rotl(v464, int32(8))
				v475 = v466 - v471 ^ base.I32_rotl(v471, int32(16))
				v479 = v467 - v475 ^ base.I32_rotl(v475, int32(19))
				v480 = v471 + v467
				v481 = v475 + v480
				v482 = v479 + v481
				v486 = v480 - v479 ^ base.I32_rotl(v479, v458)
				v487 = int32(12)
				v488 = v444 + v487
				v490 = v445 - v487
				if base.Ui32(int32(11)) < base.Ui32(v490) {
					v444 = v488
					v445 = v490
					v447 = v481
					v448 = v482
					v449 = v486
					continue
				} else {
					break
				}
				break
			}
			switch v490 - int32(1) {
			case 0:
				v663 = v481
				v664 = v482
				v665 = v486
				v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
				v671 = v663 + v666
				v672 = v664
				v673 = v665
			case 1:
				v656 = v481
				v657 = v482
				v658 = v486
				v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
				v663 = v659<<(uint(int32(8))%32) + v656
				v664 = v657
				v665 = v658
				v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
				v671 = v663 + v666
				v672 = v664
				v673 = v665
			case 2:
				v649 = v481
				v650 = v482
				v651 = v486
				v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+2)))
				v656 = v652<<(uint(int32(16))%32) + v649
				v657 = v650
				v658 = v651
				v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
				v663 = v659<<(uint(int32(8))%32) + v656
				v664 = v657
				v665 = v658
				v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
				v671 = v663 + v666
				v672 = v664
				v673 = v665
			case 3:
				v643 = v482
				v644 = v486
				v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+3)))
				v649 = v645<<(uint(int32(24))%32) + v481
				v650 = v643
				v651 = v644
				v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+2)))
				v656 = v652<<(uint(int32(16))%32) + v649
				v657 = v650
				v658 = v651
				v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
				v663 = v659<<(uint(int32(8))%32) + v656
				v664 = v657
				v665 = v658
				v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
				v671 = v663 + v666
				v672 = v664
				v673 = v665
			case 4:
				v639 = v482
				v640 = v486
				v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+4)))
				v643 = v639 + v641
				v644 = v640
				v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+3)))
				v649 = v645<<(uint(int32(24))%32) + v481
				v650 = v643
				v651 = v644
				v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+2)))
				v656 = v652<<(uint(int32(16))%32) + v649
				v657 = v650
				v658 = v651
				v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
				v663 = v659<<(uint(int32(8))%32) + v656
				v664 = v657
				v665 = v658
				v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
				v671 = v663 + v666
				v672 = v664
				v673 = v665
			case 5:
				v633 = v482
				v634 = v486
				v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+5)))
				v639 = v635<<(uint(int32(8))%32) + v633
				v640 = v634
				v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+4)))
				v643 = v639 + v641
				v644 = v640
				v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+3)))
				v649 = v645<<(uint(int32(24))%32) + v481
				v650 = v643
				v651 = v644
				v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+2)))
				v656 = v652<<(uint(int32(16))%32) + v649
				v657 = v650
				v658 = v651
				v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
				v663 = v659<<(uint(int32(8))%32) + v656
				v664 = v657
				v665 = v658
				v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
				v671 = v663 + v666
				v672 = v664
				v673 = v665
			case 6:
				v627 = v482
				v628 = v486
				v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+6)))
				v633 = v629<<(uint(int32(16))%32) + v627
				v634 = v628
				v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+5)))
				v639 = v635<<(uint(int32(8))%32) + v633
				v640 = v634
				v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+4)))
				v643 = v639 + v641
				v644 = v640
				v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+3)))
				v649 = v645<<(uint(int32(24))%32) + v481
				v650 = v643
				v651 = v644
				v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+2)))
				v656 = v652<<(uint(int32(16))%32) + v649
				v657 = v650
				v658 = v651
				v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
				v663 = v659<<(uint(int32(8))%32) + v656
				v664 = v657
				v665 = v658
				v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
				v671 = v663 + v666
				v672 = v664
				v673 = v665
			case 7:
				v622 = v486
				v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+7)))
				v627 = v623<<(uint(int32(24))%32) + v482
				v628 = v622
				v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+6)))
				v633 = v629<<(uint(int32(16))%32) + v627
				v634 = v628
				v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+5)))
				v639 = v635<<(uint(int32(8))%32) + v633
				v640 = v634
				v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+4)))
				v643 = v639 + v641
				v644 = v640
				v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+3)))
				v649 = v645<<(uint(int32(24))%32) + v481
				v650 = v643
				v651 = v644
				v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+2)))
				v656 = v652<<(uint(int32(16))%32) + v649
				v657 = v650
				v658 = v651
				v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
				v663 = v659<<(uint(int32(8))%32) + v656
				v664 = v657
				v665 = v658
				v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
				v671 = v663 + v666
				v672 = v664
				v673 = v665
			case 8:
				v617 = v486
				v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+8)))
				v622 = v618<<(uint(int32(8))%32) + v617
				v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+7)))
				v627 = v623<<(uint(int32(24))%32) + v482
				v628 = v622
				v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+6)))
				v633 = v629<<(uint(int32(16))%32) + v627
				v634 = v628
				v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+5)))
				v639 = v635<<(uint(int32(8))%32) + v633
				v640 = v634
				v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+4)))
				v643 = v639 + v641
				v644 = v640
				v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+3)))
				v649 = v645<<(uint(int32(24))%32) + v481
				v650 = v643
				v651 = v644
				v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+2)))
				v656 = v652<<(uint(int32(16))%32) + v649
				v657 = v650
				v658 = v651
				v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
				v663 = v659<<(uint(int32(8))%32) + v656
				v664 = v657
				v665 = v658
				v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
				v671 = v663 + v666
				v672 = v664
				v673 = v665
			case 9:
				v612 = v486
				v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+9)))
				v617 = v613<<(uint(int32(16))%32) + v612
				v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+8)))
				v622 = v618<<(uint(int32(8))%32) + v617
				v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+7)))
				v627 = v623<<(uint(int32(24))%32) + v482
				v628 = v622
				v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+6)))
				v633 = v629<<(uint(int32(16))%32) + v627
				v634 = v628
				v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+5)))
				v639 = v635<<(uint(int32(8))%32) + v633
				v640 = v634
				v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+4)))
				v643 = v639 + v641
				v644 = v640
				v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+3)))
				v649 = v645<<(uint(int32(24))%32) + v481
				v650 = v643
				v651 = v644
				v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+2)))
				v656 = v652<<(uint(int32(16))%32) + v649
				v657 = v650
				v658 = v651
				v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
				v663 = v659<<(uint(int32(8))%32) + v656
				v664 = v657
				v665 = v658
				v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
				v671 = v663 + v666
				v672 = v664
				v673 = v665
			case 10:
				v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+10)))
				v612 = v608<<(uint(int32(24))%32) + v486
				v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+9)))
				v617 = v613<<(uint(int32(16))%32) + v612
				v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+8)))
				v622 = v618<<(uint(int32(8))%32) + v617
				v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+7)))
				v627 = v623<<(uint(int32(24))%32) + v482
				v628 = v622
				v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+6)))
				v633 = v629<<(uint(int32(16))%32) + v627
				v634 = v628
				v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+5)))
				v639 = v635<<(uint(int32(8))%32) + v633
				v640 = v634
				v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+4)))
				v643 = v639 + v641
				v644 = v640
				v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+3)))
				v649 = v645<<(uint(int32(24))%32) + v481
				v650 = v643
				v651 = v644
				v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+2)))
				v656 = v652<<(uint(int32(16))%32) + v649
				v657 = v650
				v658 = v651
				v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
				v663 = v659<<(uint(int32(8))%32) + v656
				v664 = v657
				v665 = v658
				v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
				v671 = v663 + v666
				v672 = v664
				v673 = v665
			default:
				v671 = v481
				v672 = v482
				v673 = v486
			}
		} else {
			v504 = v381
			v505 = v391
			v507 = v398
			v508 = v398
			v509 = v398
			for {
				v511 = *(*int32)(unsafe.Add(mBase, uint32(v504)+4))
				v512 = v511 + v508
				v513 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
				v515 = *(*int32)(unsafe.Add(mBase, uint32(v504)+8))
				v516 = v515 + v509
				v518 = int32(4)
				v520 = v513 + v507 - v516 ^ base.I32_rotl(v516, v518)
				v524 = v512 - v520 ^ base.I32_rotl(v520, int32(6))
				v525 = v516 + v512
				v526 = v520 + v525
				v527 = v524 + v526
				v531 = v525 - v524 ^ base.I32_rotl(v524, int32(8))
				v535 = v526 - v531 ^ base.I32_rotl(v531, int32(16))
				v539 = v527 - v535 ^ base.I32_rotl(v535, int32(19))
				v540 = v531 + v527
				v541 = v535 + v540
				v542 = v539 + v541
				v546 = v540 - v539 ^ base.I32_rotl(v539, v518)
				v547 = int32(12)
				v548 = v504 + v547
				v550 = v505 - v547
				if base.Ui32(int32(11)) < base.Ui32(v550) {
					v504 = v548
					v505 = v550
					v507 = v541
					v508 = v542
					v509 = v546
					continue
				} else {
					break
				}
				break
			}
			switch v550 - int32(1) {
			case 0:
				v605 = v541
				v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548))))
				v671 = v605 + v606
				v672 = v542
				v673 = v546
			case 1:
				v600 = v541
				v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+1)))
				v605 = v601<<(uint(int32(8))%32) + v600
				v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548))))
				v671 = v605 + v606
				v672 = v542
				v673 = v546
			case 2:
				v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+2)))
				v600 = v596<<(uint(int32(16))%32) + v541
				v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+1)))
				v605 = v601<<(uint(int32(8))%32) + v600
				v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548))))
				v671 = v605 + v606
				v672 = v542
				v673 = v546
			case 3:
				v593 = v542
				v594 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
				v671 = v594 + v541
				v672 = v593
				v673 = v546
			case 4:
				v590 = v542
				v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+4)))
				v593 = v590 + v591
				v594 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
				v671 = v594 + v541
				v672 = v593
				v673 = v546
			case 5:
				v585 = v542
				v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+5)))
				v590 = v586<<(uint(int32(8))%32) + v585
				v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+4)))
				v593 = v590 + v591
				v594 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
				v671 = v594 + v541
				v672 = v593
				v673 = v546
			case 6:
				v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+6)))
				v585 = v581<<(uint(int32(16))%32) + v542
				v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+5)))
				v590 = v586<<(uint(int32(8))%32) + v585
				v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+4)))
				v593 = v590 + v591
				v594 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
				v671 = v594 + v541
				v672 = v593
				v673 = v546
			case 7:
				v576 = v546
				v577 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
				v579 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
				v671 = v577 + v541
				v672 = v579 + v542
				v673 = v576
			case 8:
				v571 = v546
				v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+8)))
				v576 = v572<<(uint(int32(8))%32) + v571
				v577 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
				v579 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
				v671 = v577 + v541
				v672 = v579 + v542
				v673 = v576
			case 9:
				v566 = v546
				v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+9)))
				v571 = v567<<(uint(int32(16))%32) + v566
				v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+8)))
				v576 = v572<<(uint(int32(8))%32) + v571
				v577 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
				v579 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
				v671 = v577 + v541
				v672 = v579 + v542
				v673 = v576
			case 10:
				v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+10)))
				v566 = v562<<(uint(int32(24))%32) + v546
				v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+9)))
				v571 = v567<<(uint(int32(16))%32) + v566
				v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+8)))
				v576 = v572<<(uint(int32(8))%32) + v571
				v577 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
				v579 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
				v671 = v577 + v541
				v672 = v579 + v542
				v673 = v576
			default:
				v671 = v541
				v672 = v542
				v673 = v546
			}
		}
		v676 = int32(14)
		v678 = v672 ^ v673 - base.I32_rotl(v672, v676)
		v682 = v678 ^ v671 - base.I32_rotl(v678, int32(11))
		v686 = v682 ^ v672 - base.I32_rotl(v682, int32(25))
		v690 = v686 ^ v678 - base.I32_rotl(v686, int32(16))
		v694 = v690 ^ v682 - base.I32_rotl(v690, int32(4))
		v698 = v694 ^ v686 - base.I32_rotl(v694, v676)
		*(*int64)(unsafe.Add(mBase, uint32(v381))) = base.I64_extend_i32_u(v698)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v698^v690-base.I32_rotl(v698, int32(24)))
		v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		*(*uint8)(unsafe.Add(mBase, uint32(v381)+8)) = uint8(v709)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
		return
	}
}
func F_ApplyLauncherShmemRequest(m *base.Module, l0 int32) {
	var v8 int32
	_ = v8
	Fn14205(m, l0, int32(_a_F_ApplyLauncherShmemRequest_0), int32(_a_F_ApplyLauncherShmemRequest_1), int32(128), int32(_a_F_ApplyLauncherShmemRequest_2), int32(16))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_ApplyLauncherWakeup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherWakeup[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	if v4 != 0 {
		v6 = F_pgmem_kill(m, v4, int32(10))
		mBase = m.M
	} else {
	}
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
func F_ApplyPendingListenActions(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyPendingListenActions[0]))
	if v17 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L5
	} else {
		goto L95
	}
L2:
	;
	m.G0 = v14 + int32(96)
	return
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyPendingListenActions[1]))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = v14 + int32(76)
	F_hash_seq_init(m, v25, v17)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v28 = F_hash_seq_search(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v28 == int32(0) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v33 = v14 + int32(12)
	v41 = v28
	goto L9
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyPendingListenActions[2]))
	v47 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v33)+56)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v33)+48)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v33)+40)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v33)+32)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v33)+24)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v46
	goto L14
L10:
	;
	goto L2
L11:
	;
	v183 = int32(1)
	v184 = int32(0)
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyPendingListenActions[1]))
	v190 = F_dshash_find(m, v186, v14+int32(8), v183)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L43
	}
L12:
	;
	v180 = F_strlen(m, v169)
	mBase = m.M
	goto L11
L14:
	;
	goto L15
L15:
	;
	v70 = int32(63)
	if (v33^v41)&int32(3) != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v173 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v173)
	goto L12
L17:
	;
	v154 = v149
	v155 = v150
	v156 = v151
	goto L38
L18:
	;
	if v144 == int32(0) {
		v169 = v142
		v170 = v143
		goto L16
	} else {
		goto L37
	}
L19:
	;
	v142 = v41
	v143 = v33
	v144 = v70
	goto L18
L20:
	;
	goto L21
L21:
	;
	v74 = int32(0)
	if base.B2i32(v41&int32(3) == v74)|int32(0) == v74 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v110 == int32(0) {
		v169 = v107
		v170 = v108
		goto L16
	} else {
		goto L31
	}
L23:
	;
	v86 = v41
	v87 = v33
	v88 = v70
	goto L26
L24:
	;
	goto L25
L25:
	;
	v107 = v41
	v108 = v33
	v109 = v70
	v110 = int32(1)
	goto L22
L26:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v90)
	if v90 == int32(0) {
		v149 = v86
		v150 = v87
		v151 = v88
		goto L17
	} else {
		goto L28
	}
L27:
	;
	v107 = v101
	v108 = v95
	v109 = v97
	v110 = v99
	goto L22
L28:
	;
	v94 = int32(1)
	v95 = v87 + v94
	v97 = v88 - v94
	v98 = int32(0)
	v99 = base.B2i32(v97 != v98)
	v101 = v86 + v94
	if v101&int32(3) == v98 {
		v107 = v101
		v108 = v95
		v109 = v97
		v110 = v99
		goto L22
	} else {
		goto L29
	}
L29:
	;
	if v97 != 0 {
		v86 = v101
		v87 = v95
		v88 = v97
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if base.B2i32(v113 == int32(0))|base.B2i32(base.Ui32(v109) < base.Ui32(int32(4))) != 0 {
		v142 = v107
		v143 = v108
		v144 = v109
		goto L18
	} else {
		goto L32
	}
L32:
	;
	v120 = v107
	v121 = v108
	v122 = v109
	goto L33
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v128 = int32(-2139062144)
	if (int32(16843008)-v125|v125)&v128 != v128 {
		v149 = v120
		v150 = v121
		v151 = v122
		goto L17
	} else {
		goto L35
	}
L34:
	;
	v142 = v136
	v143 = v134
	v144 = v138
	goto L18
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v125
	v133 = int32(4)
	v134 = v121 + v133
	v136 = v120 + v133
	v138 = v122 - v133
	if base.Ui32(int32(3)) < base.Ui32(v138) {
		v120 = v136
		v121 = v134
		v122 = v138
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v149 = v142
	v150 = v143
	v151 = v144
	goto L17
L38:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v158)
	if v158 == int32(0) {
		v169 = v154
		v170 = v155
		goto L16
	} else {
		goto L40
	}
L39:
	;
	v169 = v165
	v170 = v163
	goto L16
L40:
	;
	v162 = int32(1)
	v163 = v155 + v162
	v165 = v154 + v162
	v167 = v156 - v162
	if v167 != 0 {
		v154 = v165
		v155 = v163
		v156 = v167
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	if v329 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L43:
	;
	if v190 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyPendingListenActions[3]))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190)+68))
	v195 = F_dsa_get_address(m, v193, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L5
	} else {
		goto L47
	}
L45:
	;
	v298 = v183
	v302 = v184
	goto L46
L46:
	;
	if l0 == int32(0) {
		v329 = v298
		goto L42
	} else {
		goto L84
	}
L47:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v190)+72))
	if v197 <= int32(0) {
		v283 = v183
		v287 = v184
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyPendingListenActions[1]))
	F_dshash_release_lock(m, v293, v190)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L5
	} else {
		goto L83
	}
L49:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyPendingListenActions[4]))
	v207 = int32(0)
	goto L51
L50:
	;
	v279 = int32(1)
	v283 = v279
	v287 = v279
	goto L48
L51:
	;
	v216 = v195 + v207<<(uint(int32(3))%32)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	if v202 == v217 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v283 = v183
	v287 = int32(0)
	goto L48
L53:
	;
	if l0 != 0 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	goto L55
L55:
	;
	v273 = v207 + int32(1)
	if v273 != v197 {
		v207 = v273
		goto L51
	} else {
		goto L82
	}
L56:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyPendingListenActions[3]))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v190)+68))
	F_dsa_free(m, v263, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L5
	} else {
		goto L80
	}
L57:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v41)+64))
	if v219 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v241 = int32(1)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+4)))
	if v242 != v241 {
		goto L70
	} else {
		goto L71
	}
L60:
	;
	v222 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+4)) = uint8(v222)
	v283 = v222
	v287 = int32(1)
	goto L48
L61:
	;
	goto L62
L62:
	;
	v227 = v197 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v190)+72)) = v227
	if v207 < v227 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v232 = (v227 - v207) << (uint(int32(3)) % 32)
	if v232 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v238 = v227
	goto L65
L65:
	;
	if v238 == int32(0) {
		goto L56
	} else {
		goto L69
	}
L66:
	;
	base.MemoryCopy(m, v216, v216+int32(8), v232)
	goto L68
L67:
	;
	goto L68
L68:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v190)+72))
	v238 = v236
	goto L65
L69:
	;
	goto L50
L70:
	;
	v283 = int32(0)
	v287 = v241
	goto L48
L71:
	;
	goto L72
L72:
	;
	v247 = v197 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v190)+72)) = v247
	if v207 < v247 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v252 = (v247 - v207) << (uint(int32(3)) % 32)
	if v252 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v258 = v247
	goto L75
L75:
	;
	if v258 != 0 {
		goto L50
	} else {
		goto L79
	}
L76:
	;
	base.MemoryCopy(m, v216, v216+int32(8), v252)
	goto L78
L77:
	;
	goto L78
L78:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v190)+72))
	v258 = v256
	goto L75
L79:
	;
	goto L56
L80:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyPendingListenActions[1]))
	F_dshash_delete_entry(m, v268, v190)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	v329 = int32(1)
	goto L42
L82:
	;
	goto L52
L83:
	;
	v298 = v283
	v302 = v287
	goto L46
L84:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v41)+64))
	if v302|v309 != 0 {
		v329 = v298
		goto L42
	} else {
		goto L85
	}
L85:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v41
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyPendingListenActions[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v317
	F_errmsg_internal(m, int32(_a_F_ApplyPendingListenActions_0), v14)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_ApplyPendingListenActions_1), int32(1840), int32(_a_F_ApplyPendingListenActions_2))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v351 = F_hash_seq_search(m, v14+int32(76))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L5
	} else {
		goto L93
	}
L90:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyPendingListenActions[5]))
	if v341 == int32(0) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v346 = F_hash_search(m, v341, v41, int32(2), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	goto L89
L93:
	;
	if v351 != 0 {
		v41 = v351
		goto L9
	} else {
		goto L94
	}
L94:
	;
	goto L10
L95:
	;
	F_errmsg_internal(m, int32(_a_F_ApplyPendingListenActions_3), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_ApplyPendingListenActions_1), int32(1748), int32(_a_F_ApplyPendingListenActions_2))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ApplyWorkerMain(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
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
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	v8 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[0])) = uint8(v8)
	F_SetupApplyOrSyncWorker(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v14 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[0])) = uint8(v14)
		v16 = m.G0
		v18 = v16 - int32(160)
		m.G0 = v18
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
		if v22 != 0 {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v23
			v26 = v18 + int32(96)
			v31 = F_pg_snprintf(m, v26, int32(64), int32(_a_F_ApplyWorkerMain_0), v18+int32(32))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				F_StartTransactionCommand(m)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v36 = F_replorigin_by_name(m, v26, int32(1))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						if v36 == int32(0) {
							v40 = F_replorigin_create(m, v26)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								v42 = v40
								F_replorigin_session_setup(m, v42, int32(0))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									*(*uint16)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[2])) = uint16(v42)
									v48 = F_replorigin_session_get_progress(m)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return
									} else {
										F_CommitTransactionCommand(m)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return
										} else {
											v54 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[3]))
											v55 = int32(1)
											v58 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
											v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+38)))
											if v59 == v55 {
												v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+32)))
												v66 = v62 ^ int32(1)
											} else {
												v66 = int32(0)
											}
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v58)+24))
											v73 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
											v75 = m.T0[v74].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v54, v55, v55, v66&int32(1), v69, v18+int32(48))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5])) = v75
												if v75 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v280 = m.ExcPending
													if v280 != 0 {
														return
													} else {
														F_errcode(m, int32(100663808))
														mBase = m.M
														v283 = m.ExcPending
														if v283 != 0 {
															return
														} else {
															v285 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
															v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+24))
															*(*int32)(unsafe.Add(mBase, uint32(v18))) = v286
															v288 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v288
															F_errmsg(m, int32(_a_F_ApplyWorkerMain_1), v18)
															mBase = m.M
															v292 = m.ExcPending
															if v292 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_3), int32(_a_F_ApplyWorkerMain_4))
																mBase = m.M
																v297 = m.ExcPending
																if v297 != 0 {
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
													v84 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
													v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
													v86 = m.T0[v85].(func(*base.Module, int32, int32, int32) int32)(m, v75, v18+int32(52), int32(0))
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return
													} else {
														v89 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
														v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+41)))
														if v90 == int32(1) {
															F_StartTransactionCommand(m)
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return
															} else {
																v96 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																F_CheckPubDeadTupleRetention(m, v96)
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return
																} else {
																	F_CommitTransactionCommand(m)
																	mBase = m.M
																	v100 = m.ExcPending
																	if v100 != 0 {
																		return
																	} else {
																		v103 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[6]))
																		v106 = F_MemoryContextStrdup(m, v103, v18+int32(96))
																		mBase = m.M
																		v107 = m.ExcPending
																		if v107 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[7])) = v106
																			*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v48
																			v110 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v18)+56)) = uint8(v110)
																			*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v22
																			v115 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																			v117 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																			v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+24))
																			v119 = m.T0[v118].(func(*base.Module, int32) int32)(m, v115)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return
																			} else {
																				if v119 <= int32(_a_F_ApplyWorkerMain_5) {
																					if int32(_a_F_ApplyWorkerMain_6) < v119 {
																						v128 = int32(2)
																					} else {
																						v128 = int32(1)
																					}
																					if int32(_a_F_ApplyWorkerMain_7) < v119 {
																						v131 = int32(3)
																					} else {
																						v131 = v128
																					}
																					*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v131
																					v134 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																					v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+64))
																					*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v135
																					v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+34)))
																					*(*uint8)(unsafe.Add(mBase, uint32(v18)+80)) = uint8(v137)
																					if v119 < int32(_a_F_ApplyWorkerMain_8) {
																						v164 = v134
																						v166 = int32(0)
																						v167 = int32(0)
																					} else {
																						v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+35)))
																						v155 = v134
																						v156 = v142
																						v158 = int32(0)
																						if v156 != int32(102) {
																							v163 = int32(_a_F_ApplyWorkerMain_9)
																						} else {
																							v163 = v158
																						}
																						v164 = v155
																						v166 = v158
																						v167 = v163
																					}
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = int32(4)
																					v146 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																					v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+64))
																					*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v147
																					v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+34)))
																					*(*uint8)(unsafe.Add(mBase, uint32(v18)+80)) = uint8(v149)
																					v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+35)))
																					if v152 == int32(112) {
																						v164 = v146
																						v166 = v110
																						v167 = int32(_a_F_ApplyWorkerMain_10)
																					} else {
																						v155 = v146
																						v156 = v152
																						v158 = int32(0)
																						if v156 != int32(102) {
																							v163 = int32(_a_F_ApplyWorkerMain_9)
																						} else {
																							v163 = v158
																						}
																						v164 = v155
																						v166 = v158
																						v167 = v163
																					}
																				}
																				*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v167
																				v170 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[8]))
																				*(*uint8)(unsafe.Add(mBase, uint32(v170)+68)) = uint8(v166)
																				v172 = int32(0)
																				*(*uint8)(unsafe.Add(mBase, uint32(v18)+88)) = uint8(v172)
																				v174 = *(*int32)(unsafe.Add(mBase, uint32(v164)+68))
																				v175 = F_pstrdup(m, v174)
																				mBase = m.M
																				v176 = m.ExcPending
																				if v176 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v175
																					v179 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																					v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+36)))
																					if v180 != int32(112) {
																						v218 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																						v222 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																						v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+32))
																						v224 = m.T0[v223].(func(*base.Module, int32, int32) int32)(m, v218, v18+int32(56))
																						mBase = m.M
																						v225 = m.ExcPending
																						if v225 != 0 {
																							return
																						} else {
																							v228 = F_errstart(m, int32(14), int32(0))
																							mBase = m.M
																							v229 = m.ExcPending
																							if v229 != 0 {
																								return
																							} else {
																								if v228 != 0 {
																									v231 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																									v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+24))
																									v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+36)))
																									switch v234 - int32(100) {
																									case 0:
																										v240 = int32(_a_F_ApplyWorkerMain_11)
																									case 1:
																										v240 = int32(_a_F_ApplyWorkerMain_12)
																									default:
																										v240 = int32(_a_F_ApplyWorkerMain_13)
																									case 12:
																										v240 = int32(_a_F_ApplyWorkerMain_14)
																									}
																									*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v240
																									*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v232
																									F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v18+int32(16))
																									mBase = m.M
																									v247 = m.ExcPending
																									if v247 != 0 {
																										return
																									} else {
																										F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																										mBase = m.M
																										v252 = m.ExcPending
																										if v252 != 0 {
																											return
																										} else {
																											F_start_apply(m, v48)
																											mBase = m.M
																											v257 = m.ExcPending
																											if v257 != 0 {
																												return
																											} else {
																												m.G0 = v18 + int32(160)
																												F_proc_exit(m, int32(0))
																												mBase = m.M
																												v300 = m.ExcPending
																												if v300 != 0 {
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
																									F_start_apply(m, v48)
																									mBase = m.M
																									v257 = m.ExcPending
																									if v257 != 0 {
																										return
																									} else {
																										m.G0 = v18 + int32(160)
																										F_proc_exit(m, int32(0))
																										mBase = m.M
																										v300 = m.ExcPending
																										if v300 != 0 {
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
																						v183 = F_AllTablesyncsReady(m)
																						mBase = m.M
																						v184 = m.ExcPending
																						if v184 != 0 {
																							return
																						} else {
																							if v183 == int32(0) {
																								v218 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																								v222 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																								v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+32))
																								v224 = m.T0[v223].(func(*base.Module, int32, int32) int32)(m, v218, v18+int32(56))
																								mBase = m.M
																								v225 = m.ExcPending
																								if v225 != 0 {
																									return
																								} else {
																									v228 = F_errstart(m, int32(14), int32(0))
																									mBase = m.M
																									v229 = m.ExcPending
																									if v229 != 0 {
																										return
																									} else {
																										if v228 != 0 {
																											v231 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																											v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+24))
																											v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+36)))
																											switch v234 - int32(100) {
																											case 0:
																												v240 = int32(_a_F_ApplyWorkerMain_11)
																											case 1:
																												v240 = int32(_a_F_ApplyWorkerMain_12)
																											default:
																												v240 = int32(_a_F_ApplyWorkerMain_13)
																											case 12:
																												v240 = int32(_a_F_ApplyWorkerMain_14)
																											}
																											*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v240
																											*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v232
																											F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v18+int32(16))
																											mBase = m.M
																											v247 = m.ExcPending
																											if v247 != 0 {
																												return
																											} else {
																												F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																												mBase = m.M
																												v252 = m.ExcPending
																												if v252 != 0 {
																													return
																												} else {
																													F_start_apply(m, v48)
																													mBase = m.M
																													v257 = m.ExcPending
																													if v257 != 0 {
																														return
																													} else {
																														m.G0 = v18 + int32(160)
																														F_proc_exit(m, int32(0))
																														mBase = m.M
																														v300 = m.ExcPending
																														if v300 != 0 {
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
																											F_start_apply(m, v48)
																											mBase = m.M
																											v257 = m.ExcPending
																											if v257 != 0 {
																												return
																											} else {
																												m.G0 = v18 + int32(160)
																												F_proc_exit(m, int32(0))
																												mBase = m.M
																												v300 = m.ExcPending
																												if v300 != 0 {
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
																								v187 = int32(1)
																								*(*uint8)(unsafe.Add(mBase, uint32(v18)+88)) = uint8(v187)
																								v190 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																								v194 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																								v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+32))
																								v196 = m.T0[v195].(func(*base.Module, int32, int32) int32)(m, v190, v18+int32(56))
																								mBase = m.M
																								v197 = m.ExcPending
																								if v197 != 0 {
																									return
																								} else {
																									F_StartTransactionCommand(m)
																									mBase = m.M
																									v199 = m.ExcPending
																									if v199 != 0 {
																										return
																									} else {
																										v200 = F_GetTransactionSnapshot(m)
																										mBase = m.M
																										v201 = m.ExcPending
																										if v201 != 0 {
																											return
																										} else {
																											F_PushActiveSnapshot(m, v200)
																											mBase = m.M
																											v203 = m.ExcPending
																											if v203 != 0 {
																												return
																											} else {
																												v205 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																												v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
																												F_UpdateTwoPhaseState(m, v206)
																												mBase = m.M
																												v208 = m.ExcPending
																												if v208 != 0 {
																													return
																												} else {
																													v210 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																													v211 = int32(101)
																													*(*uint8)(unsafe.Add(mBase, uint32(v210)+36)) = uint8(v211)
																													F_PopActiveSnapshot(m)
																													mBase = m.M
																													v214 = m.ExcPending
																													if v214 != 0 {
																														return
																													} else {
																														F_CommitTransactionCommand(m)
																														mBase = m.M
																														v216 = m.ExcPending
																														if v216 != 0 {
																															return
																														} else {
																															v228 = F_errstart(m, int32(14), int32(0))
																															mBase = m.M
																															v229 = m.ExcPending
																															if v229 != 0 {
																																return
																															} else {
																																if v228 != 0 {
																																	v231 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																																	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+24))
																																	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+36)))
																																	switch v234 - int32(100) {
																																	case 0:
																																		v240 = int32(_a_F_ApplyWorkerMain_11)
																																	case 1:
																																		v240 = int32(_a_F_ApplyWorkerMain_12)
																																	default:
																																		v240 = int32(_a_F_ApplyWorkerMain_13)
																																	case 12:
																																		v240 = int32(_a_F_ApplyWorkerMain_14)
																																	}
																																	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v240
																																	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v232
																																	F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v18+int32(16))
																																	mBase = m.M
																																	v247 = m.ExcPending
																																	if v247 != 0 {
																																		return
																																	} else {
																																		F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																																		mBase = m.M
																																		v252 = m.ExcPending
																																		if v252 != 0 {
																																			return
																																		} else {
																																			F_start_apply(m, v48)
																																			mBase = m.M
																																			v257 = m.ExcPending
																																			if v257 != 0 {
																																				return
																																			} else {
																																				m.G0 = v18 + int32(160)
																																				F_proc_exit(m, int32(0))
																																				mBase = m.M
																																				v300 = m.ExcPending
																																				if v300 != 0 {
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
																																	F_start_apply(m, v48)
																																	mBase = m.M
																																	v257 = m.ExcPending
																																	if v257 != 0 {
																																		return
																																	} else {
																																		m.G0 = v18 + int32(160)
																																		F_proc_exit(m, int32(0))
																																		mBase = m.M
																																		v300 = m.ExcPending
																																		if v300 != 0 {
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
														} else {
															v103 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[6]))
															v106 = F_MemoryContextStrdup(m, v103, v18+int32(96))
															mBase = m.M
															v107 = m.ExcPending
															if v107 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[7])) = v106
																*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v48
																v110 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v18)+56)) = uint8(v110)
																*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v22
																v115 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																v117 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+24))
																v119 = m.T0[v118].(func(*base.Module, int32) int32)(m, v115)
																mBase = m.M
																v120 = m.ExcPending
																if v120 != 0 {
																	return
																} else {
																	if v119 <= int32(_a_F_ApplyWorkerMain_5) {
																		if int32(_a_F_ApplyWorkerMain_6) < v119 {
																			v128 = int32(2)
																		} else {
																			v128 = int32(1)
																		}
																		if int32(_a_F_ApplyWorkerMain_7) < v119 {
																			v131 = int32(3)
																		} else {
																			v131 = v128
																		}
																		*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v131
																		v134 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																		v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+64))
																		*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v135
																		v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+34)))
																		*(*uint8)(unsafe.Add(mBase, uint32(v18)+80)) = uint8(v137)
																		if v119 < int32(_a_F_ApplyWorkerMain_8) {
																			v164 = v134
																			v166 = int32(0)
																			v167 = int32(0)
																		} else {
																			v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+35)))
																			v155 = v134
																			v156 = v142
																			v158 = int32(0)
																			if v156 != int32(102) {
																				v163 = int32(_a_F_ApplyWorkerMain_9)
																			} else {
																				v163 = v158
																			}
																			v164 = v155
																			v166 = v158
																			v167 = v163
																		}
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = int32(4)
																		v146 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																		v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+64))
																		*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v147
																		v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+34)))
																		*(*uint8)(unsafe.Add(mBase, uint32(v18)+80)) = uint8(v149)
																		v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+35)))
																		if v152 == int32(112) {
																			v164 = v146
																			v166 = v110
																			v167 = int32(_a_F_ApplyWorkerMain_10)
																		} else {
																			v155 = v146
																			v156 = v152
																			v158 = int32(0)
																			if v156 != int32(102) {
																				v163 = int32(_a_F_ApplyWorkerMain_9)
																			} else {
																				v163 = v158
																			}
																			v164 = v155
																			v166 = v158
																			v167 = v163
																		}
																	}
																	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v167
																	v170 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[8]))
																	*(*uint8)(unsafe.Add(mBase, uint32(v170)+68)) = uint8(v166)
																	v172 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(v18)+88)) = uint8(v172)
																	v174 = *(*int32)(unsafe.Add(mBase, uint32(v164)+68))
																	v175 = F_pstrdup(m, v174)
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v175
																		v179 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																		v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+36)))
																		if v180 != int32(112) {
																			v218 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																			v222 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																			v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+32))
																			v224 = m.T0[v223].(func(*base.Module, int32, int32) int32)(m, v218, v18+int32(56))
																			mBase = m.M
																			v225 = m.ExcPending
																			if v225 != 0 {
																				return
																			} else {
																				v228 = F_errstart(m, int32(14), int32(0))
																				mBase = m.M
																				v229 = m.ExcPending
																				if v229 != 0 {
																					return
																				} else {
																					if v228 != 0 {
																						v231 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																						v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+24))
																						v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+36)))
																						switch v234 - int32(100) {
																						case 0:
																							v240 = int32(_a_F_ApplyWorkerMain_11)
																						case 1:
																							v240 = int32(_a_F_ApplyWorkerMain_12)
																						default:
																							v240 = int32(_a_F_ApplyWorkerMain_13)
																						case 12:
																							v240 = int32(_a_F_ApplyWorkerMain_14)
																						}
																						*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v240
																						*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v232
																						F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v18+int32(16))
																						mBase = m.M
																						v247 = m.ExcPending
																						if v247 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																							mBase = m.M
																							v252 = m.ExcPending
																							if v252 != 0 {
																								return
																							} else {
																								F_start_apply(m, v48)
																								mBase = m.M
																								v257 = m.ExcPending
																								if v257 != 0 {
																									return
																								} else {
																									m.G0 = v18 + int32(160)
																									F_proc_exit(m, int32(0))
																									mBase = m.M
																									v300 = m.ExcPending
																									if v300 != 0 {
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
																						F_start_apply(m, v48)
																						mBase = m.M
																						v257 = m.ExcPending
																						if v257 != 0 {
																							return
																						} else {
																							m.G0 = v18 + int32(160)
																							F_proc_exit(m, int32(0))
																							mBase = m.M
																							v300 = m.ExcPending
																							if v300 != 0 {
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
																			v183 = F_AllTablesyncsReady(m)
																			mBase = m.M
																			v184 = m.ExcPending
																			if v184 != 0 {
																				return
																			} else {
																				if v183 == int32(0) {
																					v218 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																					v222 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																					v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+32))
																					v224 = m.T0[v223].(func(*base.Module, int32, int32) int32)(m, v218, v18+int32(56))
																					mBase = m.M
																					v225 = m.ExcPending
																					if v225 != 0 {
																						return
																					} else {
																						v228 = F_errstart(m, int32(14), int32(0))
																						mBase = m.M
																						v229 = m.ExcPending
																						if v229 != 0 {
																							return
																						} else {
																							if v228 != 0 {
																								v231 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																								v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+24))
																								v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+36)))
																								switch v234 - int32(100) {
																								case 0:
																									v240 = int32(_a_F_ApplyWorkerMain_11)
																								case 1:
																									v240 = int32(_a_F_ApplyWorkerMain_12)
																								default:
																									v240 = int32(_a_F_ApplyWorkerMain_13)
																								case 12:
																									v240 = int32(_a_F_ApplyWorkerMain_14)
																								}
																								*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v240
																								*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v232
																								F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v18+int32(16))
																								mBase = m.M
																								v247 = m.ExcPending
																								if v247 != 0 {
																									return
																								} else {
																									F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																									mBase = m.M
																									v252 = m.ExcPending
																									if v252 != 0 {
																										return
																									} else {
																										F_start_apply(m, v48)
																										mBase = m.M
																										v257 = m.ExcPending
																										if v257 != 0 {
																											return
																										} else {
																											m.G0 = v18 + int32(160)
																											F_proc_exit(m, int32(0))
																											mBase = m.M
																											v300 = m.ExcPending
																											if v300 != 0 {
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
																								F_start_apply(m, v48)
																								mBase = m.M
																								v257 = m.ExcPending
																								if v257 != 0 {
																									return
																								} else {
																									m.G0 = v18 + int32(160)
																									F_proc_exit(m, int32(0))
																									mBase = m.M
																									v300 = m.ExcPending
																									if v300 != 0 {
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
																					v187 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v18)+88)) = uint8(v187)
																					v190 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																					v194 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																					v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+32))
																					v196 = m.T0[v195].(func(*base.Module, int32, int32) int32)(m, v190, v18+int32(56))
																					mBase = m.M
																					v197 = m.ExcPending
																					if v197 != 0 {
																						return
																					} else {
																						F_StartTransactionCommand(m)
																						mBase = m.M
																						v199 = m.ExcPending
																						if v199 != 0 {
																							return
																						} else {
																							v200 = F_GetTransactionSnapshot(m)
																							mBase = m.M
																							v201 = m.ExcPending
																							if v201 != 0 {
																								return
																							} else {
																								F_PushActiveSnapshot(m, v200)
																								mBase = m.M
																								v203 = m.ExcPending
																								if v203 != 0 {
																									return
																								} else {
																									v205 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																									v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
																									F_UpdateTwoPhaseState(m, v206)
																									mBase = m.M
																									v208 = m.ExcPending
																									if v208 != 0 {
																										return
																									} else {
																										v210 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																										v211 = int32(101)
																										*(*uint8)(unsafe.Add(mBase, uint32(v210)+36)) = uint8(v211)
																										F_PopActiveSnapshot(m)
																										mBase = m.M
																										v214 = m.ExcPending
																										if v214 != 0 {
																											return
																										} else {
																											F_CommitTransactionCommand(m)
																											mBase = m.M
																											v216 = m.ExcPending
																											if v216 != 0 {
																												return
																											} else {
																												v228 = F_errstart(m, int32(14), int32(0))
																												mBase = m.M
																												v229 = m.ExcPending
																												if v229 != 0 {
																													return
																												} else {
																													if v228 != 0 {
																														v231 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																														v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+24))
																														v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+36)))
																														switch v234 - int32(100) {
																														case 0:
																															v240 = int32(_a_F_ApplyWorkerMain_11)
																														case 1:
																															v240 = int32(_a_F_ApplyWorkerMain_12)
																														default:
																															v240 = int32(_a_F_ApplyWorkerMain_13)
																														case 12:
																															v240 = int32(_a_F_ApplyWorkerMain_14)
																														}
																														*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v240
																														*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v232
																														F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v18+int32(16))
																														mBase = m.M
																														v247 = m.ExcPending
																														if v247 != 0 {
																															return
																														} else {
																															F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																															mBase = m.M
																															v252 = m.ExcPending
																															if v252 != 0 {
																																return
																															} else {
																																F_start_apply(m, v48)
																																mBase = m.M
																																v257 = m.ExcPending
																																if v257 != 0 {
																																	return
																																} else {
																																	m.G0 = v18 + int32(160)
																																	F_proc_exit(m, int32(0))
																																	mBase = m.M
																																	v300 = m.ExcPending
																																	if v300 != 0 {
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
																														F_start_apply(m, v48)
																														mBase = m.M
																														v257 = m.ExcPending
																														if v257 != 0 {
																															return
																														} else {
																															m.G0 = v18 + int32(160)
																															F_proc_exit(m, int32(0))
																															mBase = m.M
																															v300 = m.ExcPending
																															if v300 != 0 {
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
						} else {
							v42 = v36
							F_replorigin_session_setup(m, v42, int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								*(*uint16)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[2])) = uint16(v42)
								v48 = F_replorigin_session_get_progress(m)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									F_CommitTransactionCommand(m)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										v54 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[3]))
										v55 = int32(1)
										v58 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
										v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+38)))
										if v59 == v55 {
											v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+32)))
											v66 = v62 ^ int32(1)
										} else {
											v66 = int32(0)
										}
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v58)+24))
										v73 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
										v75 = m.T0[v74].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v54, v55, v55, v66&int32(1), v69, v18+int32(48))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5])) = v75
											if v75 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v280 = m.ExcPending
												if v280 != 0 {
													return
												} else {
													F_errcode(m, int32(100663808))
													mBase = m.M
													v283 = m.ExcPending
													if v283 != 0 {
														return
													} else {
														v285 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
														v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+24))
														*(*int32)(unsafe.Add(mBase, uint32(v18))) = v286
														v288 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
														*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v288
														F_errmsg(m, int32(_a_F_ApplyWorkerMain_1), v18)
														mBase = m.M
														v292 = m.ExcPending
														if v292 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_3), int32(_a_F_ApplyWorkerMain_4))
															mBase = m.M
															v297 = m.ExcPending
															if v297 != 0 {
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
												v84 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
												v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
												v86 = m.T0[v85].(func(*base.Module, int32, int32, int32) int32)(m, v75, v18+int32(52), int32(0))
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
													v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+41)))
													if v90 == int32(1) {
														F_StartTransactionCommand(m)
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return
														} else {
															v96 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
															F_CheckPubDeadTupleRetention(m, v96)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																F_CommitTransactionCommand(m)
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return
																} else {
																	v103 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[6]))
																	v106 = F_MemoryContextStrdup(m, v103, v18+int32(96))
																	mBase = m.M
																	v107 = m.ExcPending
																	if v107 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[7])) = v106
																		*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v48
																		v110 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v18)+56)) = uint8(v110)
																		*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v22
																		v115 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																		v117 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																		v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+24))
																		v119 = m.T0[v118].(func(*base.Module, int32) int32)(m, v115)
																		mBase = m.M
																		v120 = m.ExcPending
																		if v120 != 0 {
																			return
																		} else {
																			if v119 <= int32(_a_F_ApplyWorkerMain_5) {
																				if int32(_a_F_ApplyWorkerMain_6) < v119 {
																					v128 = int32(2)
																				} else {
																					v128 = int32(1)
																				}
																				if int32(_a_F_ApplyWorkerMain_7) < v119 {
																					v131 = int32(3)
																				} else {
																					v131 = v128
																				}
																				*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v131
																				v134 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																				v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+64))
																				*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v135
																				v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+34)))
																				*(*uint8)(unsafe.Add(mBase, uint32(v18)+80)) = uint8(v137)
																				if v119 < int32(_a_F_ApplyWorkerMain_8) {
																					v164 = v134
																					v166 = int32(0)
																					v167 = int32(0)
																				} else {
																					v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+35)))
																					v155 = v134
																					v156 = v142
																					v158 = int32(0)
																					if v156 != int32(102) {
																						v163 = int32(_a_F_ApplyWorkerMain_9)
																					} else {
																						v163 = v158
																					}
																					v164 = v155
																					v166 = v158
																					v167 = v163
																				}
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = int32(4)
																				v146 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																				v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+64))
																				*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v147
																				v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+34)))
																				*(*uint8)(unsafe.Add(mBase, uint32(v18)+80)) = uint8(v149)
																				v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+35)))
																				if v152 == int32(112) {
																					v164 = v146
																					v166 = v110
																					v167 = int32(_a_F_ApplyWorkerMain_10)
																				} else {
																					v155 = v146
																					v156 = v152
																					v158 = int32(0)
																					if v156 != int32(102) {
																						v163 = int32(_a_F_ApplyWorkerMain_9)
																					} else {
																						v163 = v158
																					}
																					v164 = v155
																					v166 = v158
																					v167 = v163
																				}
																			}
																			*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v167
																			v170 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[8]))
																			*(*uint8)(unsafe.Add(mBase, uint32(v170)+68)) = uint8(v166)
																			v172 = int32(0)
																			*(*uint8)(unsafe.Add(mBase, uint32(v18)+88)) = uint8(v172)
																			v174 = *(*int32)(unsafe.Add(mBase, uint32(v164)+68))
																			v175 = F_pstrdup(m, v174)
																			mBase = m.M
																			v176 = m.ExcPending
																			if v176 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v175
																				v179 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																				v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+36)))
																				if v180 != int32(112) {
																					v218 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																					v222 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																					v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+32))
																					v224 = m.T0[v223].(func(*base.Module, int32, int32) int32)(m, v218, v18+int32(56))
																					mBase = m.M
																					v225 = m.ExcPending
																					if v225 != 0 {
																						return
																					} else {
																						v228 = F_errstart(m, int32(14), int32(0))
																						mBase = m.M
																						v229 = m.ExcPending
																						if v229 != 0 {
																							return
																						} else {
																							if v228 != 0 {
																								v231 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																								v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+24))
																								v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+36)))
																								switch v234 - int32(100) {
																								case 0:
																									v240 = int32(_a_F_ApplyWorkerMain_11)
																								case 1:
																									v240 = int32(_a_F_ApplyWorkerMain_12)
																								default:
																									v240 = int32(_a_F_ApplyWorkerMain_13)
																								case 12:
																									v240 = int32(_a_F_ApplyWorkerMain_14)
																								}
																								*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v240
																								*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v232
																								F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v18+int32(16))
																								mBase = m.M
																								v247 = m.ExcPending
																								if v247 != 0 {
																									return
																								} else {
																									F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																									mBase = m.M
																									v252 = m.ExcPending
																									if v252 != 0 {
																										return
																									} else {
																										F_start_apply(m, v48)
																										mBase = m.M
																										v257 = m.ExcPending
																										if v257 != 0 {
																											return
																										} else {
																											m.G0 = v18 + int32(160)
																											F_proc_exit(m, int32(0))
																											mBase = m.M
																											v300 = m.ExcPending
																											if v300 != 0 {
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
																								F_start_apply(m, v48)
																								mBase = m.M
																								v257 = m.ExcPending
																								if v257 != 0 {
																									return
																								} else {
																									m.G0 = v18 + int32(160)
																									F_proc_exit(m, int32(0))
																									mBase = m.M
																									v300 = m.ExcPending
																									if v300 != 0 {
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
																					v183 = F_AllTablesyncsReady(m)
																					mBase = m.M
																					v184 = m.ExcPending
																					if v184 != 0 {
																						return
																					} else {
																						if v183 == int32(0) {
																							v218 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																							v222 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																							v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+32))
																							v224 = m.T0[v223].(func(*base.Module, int32, int32) int32)(m, v218, v18+int32(56))
																							mBase = m.M
																							v225 = m.ExcPending
																							if v225 != 0 {
																								return
																							} else {
																								v228 = F_errstart(m, int32(14), int32(0))
																								mBase = m.M
																								v229 = m.ExcPending
																								if v229 != 0 {
																									return
																								} else {
																									if v228 != 0 {
																										v231 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																										v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+24))
																										v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+36)))
																										switch v234 - int32(100) {
																										case 0:
																											v240 = int32(_a_F_ApplyWorkerMain_11)
																										case 1:
																											v240 = int32(_a_F_ApplyWorkerMain_12)
																										default:
																											v240 = int32(_a_F_ApplyWorkerMain_13)
																										case 12:
																											v240 = int32(_a_F_ApplyWorkerMain_14)
																										}
																										*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v240
																										*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v232
																										F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v18+int32(16))
																										mBase = m.M
																										v247 = m.ExcPending
																										if v247 != 0 {
																											return
																										} else {
																											F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																											mBase = m.M
																											v252 = m.ExcPending
																											if v252 != 0 {
																												return
																											} else {
																												F_start_apply(m, v48)
																												mBase = m.M
																												v257 = m.ExcPending
																												if v257 != 0 {
																													return
																												} else {
																													m.G0 = v18 + int32(160)
																													F_proc_exit(m, int32(0))
																													mBase = m.M
																													v300 = m.ExcPending
																													if v300 != 0 {
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
																										F_start_apply(m, v48)
																										mBase = m.M
																										v257 = m.ExcPending
																										if v257 != 0 {
																											return
																										} else {
																											m.G0 = v18 + int32(160)
																											F_proc_exit(m, int32(0))
																											mBase = m.M
																											v300 = m.ExcPending
																											if v300 != 0 {
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
																							v187 = int32(1)
																							*(*uint8)(unsafe.Add(mBase, uint32(v18)+88)) = uint8(v187)
																							v190 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																							v194 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																							v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+32))
																							v196 = m.T0[v195].(func(*base.Module, int32, int32) int32)(m, v190, v18+int32(56))
																							mBase = m.M
																							v197 = m.ExcPending
																							if v197 != 0 {
																								return
																							} else {
																								F_StartTransactionCommand(m)
																								mBase = m.M
																								v199 = m.ExcPending
																								if v199 != 0 {
																									return
																								} else {
																									v200 = F_GetTransactionSnapshot(m)
																									mBase = m.M
																									v201 = m.ExcPending
																									if v201 != 0 {
																										return
																									} else {
																										F_PushActiveSnapshot(m, v200)
																										mBase = m.M
																										v203 = m.ExcPending
																										if v203 != 0 {
																											return
																										} else {
																											v205 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																											v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
																											F_UpdateTwoPhaseState(m, v206)
																											mBase = m.M
																											v208 = m.ExcPending
																											if v208 != 0 {
																												return
																											} else {
																												v210 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																												v211 = int32(101)
																												*(*uint8)(unsafe.Add(mBase, uint32(v210)+36)) = uint8(v211)
																												F_PopActiveSnapshot(m)
																												mBase = m.M
																												v214 = m.ExcPending
																												if v214 != 0 {
																													return
																												} else {
																													F_CommitTransactionCommand(m)
																													mBase = m.M
																													v216 = m.ExcPending
																													if v216 != 0 {
																														return
																													} else {
																														v228 = F_errstart(m, int32(14), int32(0))
																														mBase = m.M
																														v229 = m.ExcPending
																														if v229 != 0 {
																															return
																														} else {
																															if v228 != 0 {
																																v231 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																																v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+24))
																																v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+36)))
																																switch v234 - int32(100) {
																																case 0:
																																	v240 = int32(_a_F_ApplyWorkerMain_11)
																																case 1:
																																	v240 = int32(_a_F_ApplyWorkerMain_12)
																																default:
																																	v240 = int32(_a_F_ApplyWorkerMain_13)
																																case 12:
																																	v240 = int32(_a_F_ApplyWorkerMain_14)
																																}
																																*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v240
																																*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v232
																																F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v18+int32(16))
																																mBase = m.M
																																v247 = m.ExcPending
																																if v247 != 0 {
																																	return
																																} else {
																																	F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																																	mBase = m.M
																																	v252 = m.ExcPending
																																	if v252 != 0 {
																																		return
																																	} else {
																																		F_start_apply(m, v48)
																																		mBase = m.M
																																		v257 = m.ExcPending
																																		if v257 != 0 {
																																			return
																																		} else {
																																			m.G0 = v18 + int32(160)
																																			F_proc_exit(m, int32(0))
																																			mBase = m.M
																																			v300 = m.ExcPending
																																			if v300 != 0 {
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
																																F_start_apply(m, v48)
																																mBase = m.M
																																v257 = m.ExcPending
																																if v257 != 0 {
																																	return
																																} else {
																																	m.G0 = v18 + int32(160)
																																	F_proc_exit(m, int32(0))
																																	mBase = m.M
																																	v300 = m.ExcPending
																																	if v300 != 0 {
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
													} else {
														v103 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[6]))
														v106 = F_MemoryContextStrdup(m, v103, v18+int32(96))
														mBase = m.M
														v107 = m.ExcPending
														if v107 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[7])) = v106
															*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v48
															v110 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v18)+56)) = uint8(v110)
															*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v22
															v115 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
															v117 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+24))
															v119 = m.T0[v118].(func(*base.Module, int32) int32)(m, v115)
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return
															} else {
																if v119 <= int32(_a_F_ApplyWorkerMain_5) {
																	if int32(_a_F_ApplyWorkerMain_6) < v119 {
																		v128 = int32(2)
																	} else {
																		v128 = int32(1)
																	}
																	if int32(_a_F_ApplyWorkerMain_7) < v119 {
																		v131 = int32(3)
																	} else {
																		v131 = v128
																	}
																	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v131
																	v134 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+64))
																	*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v135
																	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+34)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v18)+80)) = uint8(v137)
																	if v119 < int32(_a_F_ApplyWorkerMain_8) {
																		v164 = v134
																		v166 = int32(0)
																		v167 = int32(0)
																	} else {
																		v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+35)))
																		v155 = v134
																		v156 = v142
																		v158 = int32(0)
																		if v156 != int32(102) {
																			v163 = int32(_a_F_ApplyWorkerMain_9)
																		} else {
																			v163 = v158
																		}
																		v164 = v155
																		v166 = v158
																		v167 = v163
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = int32(4)
																	v146 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+64))
																	*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v147
																	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+34)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v18)+80)) = uint8(v149)
																	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+35)))
																	if v152 == int32(112) {
																		v164 = v146
																		v166 = v110
																		v167 = int32(_a_F_ApplyWorkerMain_10)
																	} else {
																		v155 = v146
																		v156 = v152
																		v158 = int32(0)
																		if v156 != int32(102) {
																			v163 = int32(_a_F_ApplyWorkerMain_9)
																		} else {
																			v163 = v158
																		}
																		v164 = v155
																		v166 = v158
																		v167 = v163
																	}
																}
																*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v167
																v170 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[8]))
																*(*uint8)(unsafe.Add(mBase, uint32(v170)+68)) = uint8(v166)
																v172 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v18)+88)) = uint8(v172)
																v174 = *(*int32)(unsafe.Add(mBase, uint32(v164)+68))
																v175 = F_pstrdup(m, v174)
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v175
																	v179 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+36)))
																	if v180 != int32(112) {
																		v218 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																		v222 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																		v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+32))
																		v224 = m.T0[v223].(func(*base.Module, int32, int32) int32)(m, v218, v18+int32(56))
																		mBase = m.M
																		v225 = m.ExcPending
																		if v225 != 0 {
																			return
																		} else {
																			v228 = F_errstart(m, int32(14), int32(0))
																			mBase = m.M
																			v229 = m.ExcPending
																			if v229 != 0 {
																				return
																			} else {
																				if v228 != 0 {
																					v231 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																					v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+24))
																					v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+36)))
																					switch v234 - int32(100) {
																					case 0:
																						v240 = int32(_a_F_ApplyWorkerMain_11)
																					case 1:
																						v240 = int32(_a_F_ApplyWorkerMain_12)
																					default:
																						v240 = int32(_a_F_ApplyWorkerMain_13)
																					case 12:
																						v240 = int32(_a_F_ApplyWorkerMain_14)
																					}
																					*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v240
																					*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v232
																					F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v18+int32(16))
																					mBase = m.M
																					v247 = m.ExcPending
																					if v247 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																						mBase = m.M
																						v252 = m.ExcPending
																						if v252 != 0 {
																							return
																						} else {
																							F_start_apply(m, v48)
																							mBase = m.M
																							v257 = m.ExcPending
																							if v257 != 0 {
																								return
																							} else {
																								m.G0 = v18 + int32(160)
																								F_proc_exit(m, int32(0))
																								mBase = m.M
																								v300 = m.ExcPending
																								if v300 != 0 {
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
																					F_start_apply(m, v48)
																					mBase = m.M
																					v257 = m.ExcPending
																					if v257 != 0 {
																						return
																					} else {
																						m.G0 = v18 + int32(160)
																						F_proc_exit(m, int32(0))
																						mBase = m.M
																						v300 = m.ExcPending
																						if v300 != 0 {
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
																		v183 = F_AllTablesyncsReady(m)
																		mBase = m.M
																		v184 = m.ExcPending
																		if v184 != 0 {
																			return
																		} else {
																			if v183 == int32(0) {
																				v218 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																				v222 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																				v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+32))
																				v224 = m.T0[v223].(func(*base.Module, int32, int32) int32)(m, v218, v18+int32(56))
																				mBase = m.M
																				v225 = m.ExcPending
																				if v225 != 0 {
																					return
																				} else {
																					v228 = F_errstart(m, int32(14), int32(0))
																					mBase = m.M
																					v229 = m.ExcPending
																					if v229 != 0 {
																						return
																					} else {
																						if v228 != 0 {
																							v231 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																							v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+24))
																							v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+36)))
																							switch v234 - int32(100) {
																							case 0:
																								v240 = int32(_a_F_ApplyWorkerMain_11)
																							case 1:
																								v240 = int32(_a_F_ApplyWorkerMain_12)
																							default:
																								v240 = int32(_a_F_ApplyWorkerMain_13)
																							case 12:
																								v240 = int32(_a_F_ApplyWorkerMain_14)
																							}
																							*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v240
																							*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v232
																							F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v18+int32(16))
																							mBase = m.M
																							v247 = m.ExcPending
																							if v247 != 0 {
																								return
																							} else {
																								F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																								mBase = m.M
																								v252 = m.ExcPending
																								if v252 != 0 {
																									return
																								} else {
																									F_start_apply(m, v48)
																									mBase = m.M
																									v257 = m.ExcPending
																									if v257 != 0 {
																										return
																									} else {
																										m.G0 = v18 + int32(160)
																										F_proc_exit(m, int32(0))
																										mBase = m.M
																										v300 = m.ExcPending
																										if v300 != 0 {
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
																							F_start_apply(m, v48)
																							mBase = m.M
																							v257 = m.ExcPending
																							if v257 != 0 {
																								return
																							} else {
																								m.G0 = v18 + int32(160)
																								F_proc_exit(m, int32(0))
																								mBase = m.M
																								v300 = m.ExcPending
																								if v300 != 0 {
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
																				v187 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v18)+88)) = uint8(v187)
																				v190 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[5]))
																				v194 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[4]))
																				v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+32))
																				v196 = m.T0[v195].(func(*base.Module, int32, int32) int32)(m, v190, v18+int32(56))
																				mBase = m.M
																				v197 = m.ExcPending
																				if v197 != 0 {
																					return
																				} else {
																					F_StartTransactionCommand(m)
																					mBase = m.M
																					v199 = m.ExcPending
																					if v199 != 0 {
																						return
																					} else {
																						v200 = F_GetTransactionSnapshot(m)
																						mBase = m.M
																						v201 = m.ExcPending
																						if v201 != 0 {
																							return
																						} else {
																							F_PushActiveSnapshot(m, v200)
																							mBase = m.M
																							v203 = m.ExcPending
																							if v203 != 0 {
																								return
																							} else {
																								v205 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																								v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
																								F_UpdateTwoPhaseState(m, v206)
																								mBase = m.M
																								v208 = m.ExcPending
																								if v208 != 0 {
																									return
																								} else {
																									v210 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																									v211 = int32(101)
																									*(*uint8)(unsafe.Add(mBase, uint32(v210)+36)) = uint8(v211)
																									F_PopActiveSnapshot(m)
																									mBase = m.M
																									v214 = m.ExcPending
																									if v214 != 0 {
																										return
																									} else {
																										F_CommitTransactionCommand(m)
																										mBase = m.M
																										v216 = m.ExcPending
																										if v216 != 0 {
																											return
																										} else {
																											v228 = F_errstart(m, int32(14), int32(0))
																											mBase = m.M
																											v229 = m.ExcPending
																											if v229 != 0 {
																												return
																											} else {
																												if v228 != 0 {
																													v231 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyWorkerMain[1]))
																													v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+24))
																													v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+36)))
																													switch v234 - int32(100) {
																													case 0:
																														v240 = int32(_a_F_ApplyWorkerMain_11)
																													case 1:
																														v240 = int32(_a_F_ApplyWorkerMain_12)
																													default:
																														v240 = int32(_a_F_ApplyWorkerMain_13)
																													case 12:
																														v240 = int32(_a_F_ApplyWorkerMain_14)
																													}
																													*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v240
																													*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v232
																													F_errmsg_internal(m, int32(_a_F_ApplyWorkerMain_15), v18+int32(16))
																													mBase = m.M
																													v247 = m.ExcPending
																													if v247 != 0 {
																														return
																													} else {
																														F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_16), int32(_a_F_ApplyWorkerMain_4))
																														mBase = m.M
																														v252 = m.ExcPending
																														if v252 != 0 {
																															return
																														} else {
																															F_start_apply(m, v48)
																															mBase = m.M
																															v257 = m.ExcPending
																															if v257 != 0 {
																																return
																															} else {
																																m.G0 = v18 + int32(160)
																																F_proc_exit(m, int32(0))
																																mBase = m.M
																																v300 = m.ExcPending
																																if v300 != 0 {
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
																													F_start_apply(m, v48)
																													mBase = m.M
																													v257 = m.ExcPending
																													if v257 != 0 {
																														return
																													} else {
																														m.G0 = v18 + int32(160)
																														F_proc_exit(m, int32(0))
																														mBase = m.M
																														v300 = m.ExcPending
																														if v300 != 0 {
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
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v264 = m.ExcPending
			if v264 != 0 {
				return
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v267 = m.ExcPending
				if v267 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_ApplyWorkerMain_17), int32(0))
					mBase = m.M
					v271 = m.ExcPending
					if v271 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ApplyWorkerMain_2), int32(_a_F_ApplyWorkerMain_18), int32(_a_F_ApplyWorkerMain_4))
						mBase = m.M
						v276 = m.ExcPending
						if v276 != 0 {
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
				F_errfinish(m, int32(_a_F_AutoVacLauncherShutdown_1), int32(839), int32(_a_F_AutoVacLauncherShutdown_2))
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
func F_AutoVacuumShmemRequest(m *base.Module, l0 int32) {
	var v8 int32
	_ = v8
	Fn14205(m, l0, int32(_a_F_AutoVacuumShmemRequest_0), int32(_a_F_AutoVacuumShmemRequest_1), int32(40), int32(_a_F_AutoVacuumShmemRequest_2), int32(_a_F_AutoVacuumShmemRequest_3))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_a_swap(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return
}
func F_accumArrayResultAny(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32 {
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
						v35 = F_MemoryContextAlloc(m, v24, int32(512))
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
			case 0, 2, 3, 4, 5, 10, 11, 13, 31, 32, 33, 34, 36, 41, 44, 45, 48, 49, 51:
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
						F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2779), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
			case 35:
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
			case 40:
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
			case 43:
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
			case 47:
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
			case 50:
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
			case 52:
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2784), int32(_a_F_aclcheck_error_2))
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
			case 0, 2, 3, 4, 5, 10, 11, 13, 27, 31, 32, 33, 34, 44, 48, 49, 51:
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
						F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2918), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
			case 6, 28, 36, 41, 45:
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
			case 35:
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
			case 40:
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
			case 43:
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
			case 47:
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
			case 50:
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
			case 52:
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
							F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2923), int32(_a_F_aclcheck_error_2))
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
					F_errfinish(m, int32(_a_F_aclcheck_error_1), int32(2927), int32(_a_F_aclcheck_error_2))
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
func F_aclcontains(m *base.Module, l0 int32) int64 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return int64(0)
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
	return int64(1)
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
			F_pg_qsort(m, v17+l0, v6, int32(16), int32(1364))
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
	F_errfinish(m, int32(_a_F_aclupdate_1), int32(446), int32(_a_F_aclupdate_2))
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
	F_errfinish(m, int32(_a_F_aclupdate_1), int32(1304), int32(_a_F_aclupdate_4))
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
	F_errfinish(m, int32(_a_F_aclupdate_1), int32(446), int32(_a_F_aclupdate_2))
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
	F_errfinish(m, int32(_a_F_aclupdate_1), int32(1366), int32(_a_F_aclupdate_7))
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
	v5 = Fn14227(m, l0, l1, l2, int32(1130))
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int64
	_ = v32
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 float64
	_ = v132
	var v140 int32
	_ = v140
	var v147 float64
	_ = v147
	var v153 float64
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v269 float64
	_ = v269
	var v270 float64
	_ = v270
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v472 int32
	_ = v472
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v620 int32
	_ = v620
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v715 int32
	_ = v715
	var v717 int64
	_ = v717
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v758 int32
	_ = v758
	var v760 int64
	_ = v760
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 float64
	_ = v773
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v910 int32
	_ = v910
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v952 int64
	_ = v952
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 float64
	_ = v961
	var v963 int32
	_ = v963
	var v987 float64
	_ = v987
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1012 int32
	_ = v1012
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1025 int32
	_ = v1025
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1132 int32
	_ = v1132
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1159 int64
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1250 int32
	_ = v1250
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1273 int32
	_ = v1273
	var v1279 int32
	_ = v1279
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1322 int32
	_ = v1322
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1342 int64
	_ = v1342
	var v1352 int32
	_ = v1352
	var v1361 int32
	_ = v1361
	var v1365 int32
	_ = v1365
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1417 int32
	_ = v1417
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1447 int32
	_ = v1447
	var v1449 int32
	_ = v1449
	var v1453 int32
	_ = v1453
	var v1461 int32
	_ = v1461
	var v1470 int32
	_ = v1470
	var v1474 int32
	_ = v1474
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1506 int32
	_ = v1506
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
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
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1602 int32
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1621 int32
	_ = v1621
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1641 int32
	_ = v1641
	var v1651 int32
	_ = v1651
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1661 int32
	_ = v1661
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1676 int32
	_ = v1676
	var v1684 int32
	_ = v1684
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1701 int32
	_ = v1701
	var v1708 int32
	_ = v1708
	var v1710 int32
	_ = v1710
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1739 int32
	_ = v1739
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1756 int32
	_ = v1756
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1792 int32
	_ = v1792
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1808 int32
	_ = v1808
	var v1812 int32
	_ = v1812
	var v1816 int32
	_ = v1816
	var v1819 int32
	_ = v1819
	var v1827 int32
	_ = v1827
	var v1835 int32
	_ = v1835
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1852 int32
	_ = v1852
	var v1859 int32
	_ = v1859
	var v1861 int32
	_ = v1861
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1890 int32
	_ = v1890
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1907 int32
	_ = v1907
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1934 float64
	_ = v1934
	var v1939 int32
	_ = v1939
	var v1942 float64
	_ = v1942
	var v1944 float64
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1956 int32
	_ = v1956
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1970 int32
	_ = v1970
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1978 int32
	_ = v1978
	var v1984 int32
	_ = v1984
	var v1992 int32
	_ = v1992
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2009 int32
	_ = v2009
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2035 int32
	_ = v2035
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2044 int32
	_ = v2044
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2059 int32
	_ = v2059
	var v2076 int32
	_ = v2076
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2108 int32
	_ = v2108
	var v2113 int32
	_ = v2113
	var v2120 int32
	_ = v2120
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2126 int32
	_ = v2126
	var v2137 int32
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2162 int32
	_ = v2162
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2167 int32
	_ = v2167
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2229 int32
	_ = v2229
	var v2231 int64
	_ = v2231
	var v2235 int32
	_ = v2235
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2242 int32
	_ = v2242
	var v2245 int32
	_ = v2245
	var v2247 int64
	_ = v2247
	var v2251 int32
	_ = v2251
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2265 int32
	_ = v2265
	var v2267 int64
	_ = v2267
	var v2271 int32
	_ = v2271
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2302 int32
	_ = v2302
	var v2304 int32
	_ = v2304
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2361 int32
	_ = v2361
	var v2378 int32
	_ = v2378
	var v2392 int32
	_ = v2392
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2413 int32
	_ = v2413
	var v2432 int32
	_ = v2432
	var v2436 int32
	_ = v2436
	var v2437 int32
	_ = v2437
	var v2440 int32
	_ = v2440
	var v2455 int32
	_ = v2455
	var v2465 int32
	_ = v2465
	var v2468 int32
	_ = v2468
	var v2471 int32
	_ = v2471
	var v2475 int32
	_ = v2475
	var v2479 int32
	_ = v2479
	var v2482 int32
	_ = v2482
	var v2509 int32
	_ = v2509
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2545 int32
	_ = v2545
	var v2553 int32
	_ = v2553
	var v2560 int32
	_ = v2560
	var v2561 int32
	_ = v2561
	var v2570 int32
	_ = v2570
	var v2589 int32
	_ = v2589
	var v2590 int32
	_ = v2590
	var v2591 int32
	_ = v2591
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2608 int32
	_ = v2608
	var v2609 int32
	_ = v2609
	var v2611 int32
	_ = v2611
	var v2614 int32
	_ = v2614
	var v2615 int32
	_ = v2615
	var v2620 int32
	_ = v2620
	var v2628 int32
	_ = v2628
	var v2630 int32
	_ = v2630
	var v2632 int32
	_ = v2632
	var v2633 int32
	_ = v2633
	var v2636 int32
	_ = v2636
	var v2640 int32
	_ = v2640
	var v2647 int32
	_ = v2647
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2658 int32
	_ = v2658
	var v2667 int32
	_ = v2667
	var v2679 int32
	_ = v2679
	var v2683 int32
	_ = v2683
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2700 int32
	_ = v2700
	var v2703 int32
	_ = v2703
	var v2704 int32
	_ = v2704
	var v2709 int32
	_ = v2709
	var v2716 int32
	_ = v2716
	var v2718 int32
	_ = v2718
	var v2720 int32
	_ = v2720
	var v2723 int32
	_ = v2723
	var v2725 int32
	_ = v2725
	var v2727 int32
	_ = v2727
	var v2734 int32
	_ = v2734
	var v2741 int32
	_ = v2741
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2757 int32
	_ = v2757
	var v2773 int32
	_ = v2773
	var v2774 float64
	_ = v2774
	var v2775 float64
	_ = v2775
	var v2779 float64
	_ = v2779
	var v2780 float64
	_ = v2780
	var v2788 int32
	_ = v2788
	var v2794 int32
	_ = v2794
	var v2797 int32
	_ = v2797
	var v2798 int32
	_ = v2798
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2815 int32
	_ = v2815
	var v2816 int32
	_ = v2816
	var v2818 int32
	_ = v2818
	var v2821 int32
	_ = v2821
	var v2822 int32
	_ = v2822
	var v2827 int32
	_ = v2827
	var v2835 int32
	_ = v2835
	var v2837 int32
	_ = v2837
	var v2839 int32
	_ = v2839
	var v2840 int32
	_ = v2840
	var v2843 int32
	_ = v2843
	var v2847 int32
	_ = v2847
	var v2853 int32
	_ = v2853
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2866 int32
	_ = v2866
	var v2867 int32
	_ = v2867
	var v2872 int32
	_ = v2872
	var v2888 int32
	_ = v2888
	var v2889 float64
	_ = v2889
	var v2890 float64
	_ = v2890
	var v2894 float64
	_ = v2894
	var v2895 float64
	_ = v2895
	var v2903 int32
	_ = v2903
	var v2909 int32
	_ = v2909
	var v2912 int32
	_ = v2912
	var v2913 int32
	_ = v2913
	var v2916 int32
	_ = v2916
	var v2917 int32
	_ = v2917
	var v2925 int32
	_ = v2925
	var v2950 int32
	_ = v2950
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2980 int32
	_ = v2980
	var v2982 int64
	_ = v2982
	var v2986 int32
	_ = v2986
	var v2990 int32
	_ = v2990
	var v2991 int32
	_ = v2991
	var v2993 int32
	_ = v2993
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3050 int32
	_ = v3050
	var v3053 int32
	_ = v3053
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3058 int32
	_ = v3058
	var v3075 int32
	_ = v3075
	var v3087 int32
	_ = v3087
	var v3091 int32
	_ = v3091
	var v3096 int32
	_ = v3096
	var v3102 int32
	_ = v3102
	var v3103 int32
	_ = v3103
	var v3105 int32
	_ = v3105
	var v3106 int32
	_ = v3106
	var v3108 int64
	_ = v3108
	var v3110 int32
	_ = v3110
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3116 int32
	_ = v3116
	var v3119 int32
	_ = v3119
	var v3120 int32
	_ = v3120
	v4 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(256)
	m.G0 = v28
	*(*int32)(unsafe.Add(mBase, uint32(v28)+200)) = v4
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+192)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v28)+184)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(v28)+176)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v28)+168)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(v28)+160)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v28)+152)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(v28)+144)) = v32
	v46 = int32(1)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_add_paths_to_append_rel[0])))
	if v48 == v46 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	v52 = v51
	goto L3
L2:
	;
	v52 = v4
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
	if v737&int32(1) != 0 {
		goto L146
	} else {
		goto L147
	}
L5:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v28)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+136)) = v715
	v717 = *(*int64)(unsafe.Add(mBase, uint32(v28)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+128)) = v717
	v721 = int32(0)
	v726 = F_create_append_path(m, l0, l1, v28+int32(128), v721, v721, v721, v721, float64(-1))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L18
	} else {
		goto L144
	}
L6:
	;
	v696 = v46
	v699 = int32(1)
	v701 = v52
	v703 = v4
	v706 = v4
	goto L5
L7:
	;
	goto L8
L8:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v56 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v676&int32(1) != 0 {
		v696 = v668
		v699 = v671
		v701 = v673
		v703 = v675
		v706 = v678
		goto L5
	} else {
		goto L143
	}
L10:
	;
	v59 = int32(1)
	v668 = v46
	v671 = v59
	v673 = v52
	v675 = v4
	v676 = v59
	v678 = v4
	goto L9
L11:
	;
	goto L12
L12:
	;
	v62 = v28 + int32(152)
	v65 = int32(4)
	v77 = int32(1)
	v85 = v46
	v86 = v4
	v88 = v77
	v90 = v52
	v92 = v4
	v93 = v77
	v95 = v4
	goto L13
L13:
	;
	v104 = int32(0)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105+v86<<(uint(int32(2))%32))))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+44))
	if v110 == v104 {
		v122 = v104
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v668 = v200
	v671 = v211
	v673 = v285
	v675 = v646
	v676 = v122
	v678 = v649
	goto L9
L15:
	;
	v123 = int32(0)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v125 != int32(1) {
		v200 = v123
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v109)+60))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	if v115 != 0 {
		v122 = int32(0)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_accumulate_append_subpath(m, v114, v28+int32(192), int32(0), v28+int32(200))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v122 = v93
	goto L15
L20:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v109)+52))
	if v201 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L21:
	;
	v128 = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v109)+56))
	if v129 == v128 {
		v200 = v128
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v132 = *(*float64)(unsafe.Add(mBase, uint32(l0)+312))
	if base.F64_gt(v132, float64(0)) != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v109)+60))
	if base.F64_le(v132, float64(0)) != 0 {
		v187 = v140
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v192 = v129
	goto L25
L25:
	;
	F_accumulate_append_subpath(m, v192, v28+int32(176), int32(0), v28+int32(184))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L18
	} else {
		goto L43
	}
L26:
	;
	v192 = v187
	goto L25
L27:
	;
	goto L26
L28:
	;
	if base.F64_ge(v132, float64(1)) == int32(0) {
		v153 = v132
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v109)+44))
	if v155 == int32(0) {
		v187 = v140
		goto L27
	} else {
		goto L32
	}
L30:
	;
	v147 = *(*float64)(unsafe.Add(mBase, uint32(v140)+32))
	if base.F64_gt(v147, float64(0)) == int32(0) {
		v153 = v132
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v153 = base.F64_div(v132, v147)
	goto L29
L32:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v158 <= int32(0) {
		v187 = v140
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v163 = v140
	v166 = int32(0)
	goto L34
L34:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168+v166<<(uint(int32(2))%32))))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+16))
	if v173 != 0 {
		v180 = v163
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v187 = v180
	goto L27
L36:
	;
	v182 = v166 + int32(1)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v182 < v183 {
		v163 = v180
		v166 = v182
		goto L34
	} else {
		goto L42
	}
L37:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v109)+60))
	if v172 == v174 {
		v180 = v163
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v176 = F_compare_fractional_path_costs(m, v163, v172, v153)
	mBase = m.M
	if v176 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v179 = v163
	goto L41
L40:
	;
	v179 = v172
	goto L41
L41:
	;
	v180 = v179
	goto L36
L42:
	;
	goto L35
L43:
	;
	v200 = v85
	goto L20
L44:
	;
	v214 = int32(0)
	if v90&int32(1) == v214 {
		v285 = v214
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v210 = v123
	v211 = int32(0)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	F_accumulate_append_subpath(m, v206, v28+int32(160)|v65, int32(0), v28+int32(168))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L18
	} else {
		goto L48
	}
L48:
	;
	v210 = v206
	v211 = v88
	goto L44
L49:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v109)+44))
	if v286 == int32(0) {
		v646 = v92
		v649 = v95
		goto L78
	} else {
		goto L79
	}
L50:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v109)+44))
	if v217 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	if v210|v261 == int32(0) {
		v285 = v214
		goto L49
	} else {
		goto L68
	}
L52:
	;
	goto L51
L53:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if v222 <= int32(0) {
		v261 = int32(0)
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v261 = int32(0)
	goto L52
L56:
	;
	v225 = int32(0)
	if v225 < v222 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v228 = v222
	goto L59
L58:
	;
	v228 = v225
	goto L59
L59:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v231 = int32(0)
	goto L60
L60:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v229+v231<<(uint(int32(2))%32))))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+21)))
	if v240 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L55
L62:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v239)+16))
	if v243 == int32(0) {
		v261 = v239
		goto L52
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v251 = v231 + int32(1)
	if v251 != v228 {
		v231 = v251
		goto L60
	} else {
		goto L67
	}
L65:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	if v246 == int32(0) {
		v261 = v239
		goto L52
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	goto L61
L68:
	;
	if v261 != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v285 = int32(1)
	goto L49
L70:
	;
	F_accumulate_append_subpath(m, v261, v28+int32(144), int32(0), v62)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L18
	} else {
		goto L77
	}
L71:
	;
	if v210 == int32(0) {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	F_accumulate_append_subpath(m, v210, v28+int32(144)|v65, v28+int32(144), v62)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L18
	} else {
		goto L76
	}
L74:
	;
	v269 = *(*float64)(unsafe.Add(mBase, uint32(v210)+56))
	v270 = *(*float64)(unsafe.Add(mBase, uint32(v261)+56))
	if base.F64_lt(v269, v270) == int32(0) {
		goto L70
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	goto L69
L77:
	;
	goto L69
L78:
	;
	v659 = v86 + int32(1)
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v659 < v660 {
		v85 = v200
		v86 = v659
		v88 = v211
		v90 = v285
		v92 = v646
		v93 = v122
		v95 = v649
		goto L13
	} else {
		goto L142
	}
L79:
	;
	v289 = int32(0)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v290 <= v289 {
		v646 = v92
		v649 = v95
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v301 = v289
	v306 = v92
	v309 = v95
	goto L81
L81:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v318+v301<<(uint(int32(2))%32))))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+64))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v322)+16))
	if v325 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v646 = v472
	v649 = v620
	goto L78
L83:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	v327 = v326
	goto L85
L84:
	;
	v327 = int32(0)
	goto L85
L85:
	;
	if v323 == int32(0) {
		v472 = v306
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if v327 == int32(0) {
		v620 = v309
		goto L120
	} else {
		goto L121
	}
L87:
	;
	if v306 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v457 = F_lappend(m, v306, v323)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L18
	} else {
		goto L119
	}
L89:
	;
	v332 = int32(0)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	if v333 <= v332 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v340 = v332
	goto L91
L91:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v306)+12))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v361+v340<<(uint(int32(2))%32))))
	if v365 == v323 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L88
L93:
	;
	if v425 == int32(0) {
		v472 = v306
		goto L86
	} else {
		goto L117
	}
L94:
	;
	v425 = int32(0)
	goto L93
L95:
	;
	goto L96
L96:
	;
	v374 = int32(0)
	goto L99
L97:
	;
	if v414 != 0 {
		goto L114
	} else {
		goto L115
	}
L98:
	;
	v409 = int32(0)
	if v396 != 0 {
		goto L111
	} else {
		goto L112
	}
L99:
	;
	v378 = int32(0)
	if v365 == v378 {
		v388 = v378
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v425 = int32(3)
	goto L93
L101:
	;
	if v323 != 0 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	if v382 <= v374 {
		v388 = int32(0)
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v365)+12))
	v388 = v384 + v374<<(uint(int32(2))%32)
	goto L101
L104:
	;
	v394 = int32(0)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v323)+12))
	if base.B2i32(v388 == v394)|base.B2i32(v396 == v394) != 0 {
		goto L98
	} else {
		goto L109
	}
L105:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	if v374 < v389 {
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v391 = int32(0)
	v414 = base.B2i32(v388 == v391)
	v416 = v391
	goto L97
L108:
	;
	goto L107
L109:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v396+v374<<(uint(int32(2))%32))))
	if v404 == v406 {
		v374 = v374 + int32(1)
		goto L99
	} else {
		goto L110
	}
L110:
	;
	goto L100
L111:
	;
	v413 = int32(2)
	goto L113
L112:
	;
	v413 = v409
	goto L113
L113:
	;
	v414 = base.B2i32(v388 == v409)
	v416 = v413
	goto L97
L114:
	;
	v418 = v416
	goto L116
L115:
	;
	v418 = int32(1)
	goto L116
L116:
	;
	v425 = v418
	goto L93
L117:
	;
	v429 = v340 + int32(1)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	if v429 < v430 {
		v340 = v429
		goto L91
	} else {
		goto L118
	}
L118:
	;
	goto L92
L119:
	;
	v472 = v457
	goto L86
L120:
	;
	v630 = v301 + int32(1)
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v630 < v631 {
		v301 = v630
		v306 = v472
		v309 = v620
		goto L81
	} else {
		goto L141
	}
L121:
	;
	if v309 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v602 = F_lappend(m, v309, v327)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L18
	} else {
		goto L140
	}
L123:
	;
	v488 = int32(0)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	if v489 <= v488 {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v496 = v488
	goto L125
L125:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v309)+12))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v517+v496<<(uint(int32(2))%32))))
	v522 = int32(0)
	if base.B2i32(v521 == v522)|base.B2i32(v327 == v522) != 0 {
		v568 = base.B2i32(v521|v327 == v522)
		goto L128
	} else {
		goto L129
	}
L126:
	;
	goto L122
L127:
	;
	if v568 != 0 {
		v620 = v309
		goto L120
	} else {
		goto L138
	}
L128:
	;
	goto L127
L129:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v521)+4))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v536 != v537 {
		v568 = int32(0)
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v539 = int32(1)
	if v536 <= v539 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v542 = v539
	goto L133
L132:
	;
	v542 = v536
	goto L133
L133:
	;
	v543 = int32(8)
	v548 = int32(0)
	goto L134
L134:
	;
	v556 = v548 << (uint(int32(2)) % 32)
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v521+v543+v556)))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v327+v543+v556)))
	v561 = base.B2i32(v558 == v560)
	if v558 != v560 {
		v568 = v561
		goto L128
	} else {
		goto L136
	}
L135:
	;
	v568 = v561
	goto L128
L136:
	;
	v564 = v548 + int32(1)
	if v564 != v542 {
		v548 = v564
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v574 = v496 + int32(1)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	if v574 < v575 {
		v496 = v574
		goto L125
	} else {
		goto L139
	}
L139:
	;
	goto L126
L140:
	;
	v620 = v602
	goto L120
L141:
	;
	goto L82
L142:
	;
	goto L14
L143:
	;
	v737 = v668
	v740 = v671
	v741 = int32(0)
	v742 = v673
	v744 = v675
	v747 = v678
	goto L4
L144:
	;
	F_add_path(m, l1, v726)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L18
	} else {
		goto L145
	}
L145:
	;
	v737 = v696
	v740 = v699
	v741 = int32(1)
	v742 = v701
	v744 = v703
	v747 = v706
	goto L4
L146:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v28)+184))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+120)) = v758
	v760 = *(*int64)(unsafe.Add(mBase, uint32(v28)+176))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+112)) = v760
	v764 = int32(0)
	v769 = F_create_append_path(m, l0, l1, v28+int32(112), v764, v764, v764, v764, float64(-1))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L18
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v773 = float64(-1)
	if v740 == int32(0) {
		v987 = v773
		goto L151
	} else {
		goto L152
	}
L149:
	;
	F_add_path(m, l1, v769)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L18
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v989 = int32(0)
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v28)+144))
	if base.B2i32(v990 != v989)&v742 != 0 {
		goto L195
	} else {
		goto L196
	}
L152:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v28)+164))
	if v776 == int32(0) {
		v987 = v773
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v779 = int32(0)
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v776)+4))
	if v780 <= v779 {
		v910 = v779
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v932 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_add_paths_to_append_rel[0])))
	if v932 == int32(1) {
		goto L181
	} else {
		goto L182
	}
L155:
	;
	v784 = v780 & int32(3)
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v776)+12))
	v786 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v780) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v797 = v779
	v798 = v786
	v800 = int32(0)
	goto L159
L157:
	;
	v848 = v779
	v849 = v786
	goto L158
L158:
	;
	v873 = v848
	v874 = v849
	v877 = v786
	goto L175
L159:
	;
	v820 = v785 + v798<<(uint(int32(2))%32)
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v820)))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v821)+24))
	if v822 < v797 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	if v784 == int32(0) {
		v910 = v836
		goto L154
	} else {
		goto L174
	}
L161:
	;
	v824 = v797
	goto L163
L162:
	;
	v824 = v822
	goto L163
L163:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v820)+4))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v825)+24))
	if v826 < v824 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v828 = v824
	goto L166
L165:
	;
	v828 = v826
	goto L166
L166:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v820)+8))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v829)+24))
	if v830 < v828 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v832 = v828
	goto L169
L168:
	;
	v832 = v830
	goto L169
L169:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v820)+12))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v833)+24))
	if v834 < v832 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v836 = v832
	goto L172
L171:
	;
	v836 = v834
	goto L172
L172:
	;
	v837 = int32(4)
	v838 = v798 + v837
	v840 = v800 + v837
	if v840 != v780&int32(2147483644) {
		v797 = v836
		v798 = v838
		v800 = v840
		goto L159
	} else {
		goto L173
	}
L173:
	;
	goto L160
L174:
	;
	v848 = v836
	v849 = v838
	goto L158
L175:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v785+v874<<(uint(int32(2))%32))))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v897)+24))
	if v898 < v873 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v910 = v900
	goto L154
L177:
	;
	v900 = v873
	goto L179
L178:
	;
	v900 = v898
	goto L179
L179:
	;
	v901 = int32(1)
	v904 = v877 + v901
	if v904 != v784 {
		v873 = v900
		v874 = v874 + v901
		v877 = v904
		goto L175
	} else {
		goto L180
	}
L180:
	;
	goto L176
L181:
	;
	if l2 != 0 {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	v947 = v910
	goto L183
L183:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v28)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+104)) = v950
	v952 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+96)) = v952
	v956 = int32(0)
	v959 = F_create_append_path(m, l0, l1, v28+int32(96), v956, v956, v947, v932, float64(-1))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L18
	} else {
		goto L193
	}
L184:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v940 = int32(32) - base.I32_clz(v937)
	goto L186
L185:
	;
	v940 = int32(0)
	goto L186
L186:
	;
	if v940 < v910 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v942 = v910
	goto L189
L188:
	;
	v942 = v940
	goto L189
L189:
	;
	v944 = *(*int32)(unsafe.Add(mBase, _c_F_add_paths_to_append_rel[1]))
	if v942 < v944 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v946 = v942
	goto L192
L191:
	;
	v946 = v944
	goto L192
L192:
	;
	v947 = v946
	goto L183
L193:
	;
	v961 = *(*float64)(unsafe.Add(mBase, uint32(v959)+32))
	F_add_partial_path(m, l1, v959)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L18
	} else {
		goto L194
	}
L194:
	;
	v987 = v961
	goto L151
L195:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v28)+148))
	if v994 == int32(0) {
		v1132 = v989
		goto L198
	} else {
		goto L199
	}
L196:
	;
	goto L197
L197:
	;
	if v741 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L198:
	;
	if l2 != 0 {
		goto L229
	} else {
		goto L230
	}
L199:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v994)+4))
	if v997 <= int32(0) {
		v1132 = v989
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1000 = int32(0)
	if v1000 < v997 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1003 = v997
	goto L203
L202:
	;
	v1003 = v1000
	goto L203
L203:
	;
	v1005 = v1003 & int32(3)
	v1006 = int32(0)
	if int32(4) <= v997 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v994)+12))
	v1018 = v989
	v1019 = v1006
	v1025 = int32(0)
	goto L207
L205:
	;
	v1069 = v989
	v1070 = v1006
	goto L206
L206:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v994)+12))
	v1095 = v1069
	v1096 = v1070
	v1099 = v1006
	goto L223
L207:
	;
	v1041 = v1012 + v1019<<(uint(int32(2))%32)
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1041)))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1042)+24))
	if v1043 < v1018 {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	if v1005 == int32(0) {
		v1132 = v1057
		goto L198
	} else {
		goto L222
	}
L209:
	;
	v1045 = v1018
	goto L211
L210:
	;
	v1045 = v1043
	goto L211
L211:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+4))
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1046)+24))
	if v1047 < v1045 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1049 = v1045
	goto L214
L213:
	;
	v1049 = v1047
	goto L214
L214:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+8))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1050)+24))
	if v1051 < v1049 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1053 = v1049
	goto L217
L216:
	;
	v1053 = v1051
	goto L217
L217:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+12))
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+24))
	if v1055 < v1053 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1057 = v1053
	goto L220
L219:
	;
	v1057 = v1055
	goto L220
L220:
	;
	v1058 = int32(4)
	v1059 = v1019 + v1058
	v1061 = v1025 + v1058
	if v1061 != v1003&int32(2147483644) {
		v1018 = v1057
		v1019 = v1059
		v1025 = v1061
		goto L207
	} else {
		goto L221
	}
L221:
	;
	goto L208
L222:
	;
	v1069 = v1057
	v1070 = v1059
	goto L206
L223:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1090+v1096<<(uint(int32(2))%32))))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+24))
	if v1120 < v1095 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v1132 = v1122
	goto L198
L225:
	;
	v1122 = v1095
	goto L227
L226:
	;
	v1122 = v1120
	goto L227
L227:
	;
	v1123 = int32(1)
	v1126 = v1099 + v1123
	if v1126 != v1005 {
		v1095 = v1122
		v1096 = v1096 + v1123
		v1099 = v1126
		goto L223
	} else {
		goto L228
	}
L228:
	;
	goto L224
L229:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1158 = int32(32) - base.I32_clz(v1155)
	goto L231
L230:
	;
	v1158 = int32(0)
	goto L231
L231:
	;
	v1159 = *(*int64)(unsafe.Add(mBase, uint32(v28)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+80)) = v1159
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v28)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+88)) = v1161
	v1165 = int32(0)
	if v1158 < v1132 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v1168 = v1132
	goto L234
L233:
	;
	v1168 = v1158
	goto L234
L234:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, _c_F_add_paths_to_append_rel[1]))
	if v1168 < v1170 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v1172 = v1168
	goto L237
L236:
	;
	v1172 = v1170
	goto L237
L237:
	;
	v1174 = F_create_append_path(m, l0, l1, v28+int32(80), v1165, v1165, v1172, int32(1), v987)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L18
	} else {
		goto L238
	}
L238:
	;
	F_add_partial_path(m, l1, v1174)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L18
	} else {
		goto L239
	}
L239:
	;
	goto L197
L240:
	;
	if v747 == int32(0) {
		goto L550
	} else {
		goto L551
	}
L241:
	;
	v1205 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+255)) = uint8(v1205)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+254)) = uint8(v1205)
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	if v1209 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	if v744 == int32(0) {
		goto L240
	} else {
		goto L271
	}
L243:
	;
	v1296 = int32(0)
	v1297 = v4
	goto L242
L244:
	;
	goto L245
L245:
	;
	v1213 = int32(0)
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v1214 {
	case 0, 2:
		goto L246
	default:
		v1296 = v1213
		v1297 = v4
		goto L242
	}
L246:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+280))
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1216)))
	switch v1217 - int32(108) {
	case 0:
		goto L249
	default:
		goto L248
	case 6:
		goto L250
	}
L247:
	;
	if v1279 == int32(0) {
		v1296 = v1213
		v1297 = v4
		goto L242
	} else {
		goto L268
	}
L248:
	;
	v1279 = int32(0)
	goto L247
L249:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1216)+16))
	v1228 = int32(0)
	if base.B2i32(v1215 == v1228)|base.B2i32(v1227 == v1228) != 0 {
		v1273 = v1228
		goto L255
	} else {
		goto L256
	}
L250:
	;
	v1220 = int32(1)
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1216)+32))
	if v1221 == int32(-1) {
		v1279 = v1220
		goto L247
	} else {
		goto L251
	}
L251:
	;
	v1224 = F_bms_is_member(m, v1221, v1215)
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L18
	} else {
		goto L252
	}
L252:
	;
	if v1224 != 0 {
		goto L248
	} else {
		goto L253
	}
L253:
	;
	v1279 = v1220
	goto L247
L254:
	;
	if v1273 == int32(0) {
		v1279 = int32(1)
		goto L247
	} else {
		goto L267
	}
L255:
	;
	goto L254
L256:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1215)+4))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+4))
	if v1238 < v1239 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1241 = v1238
	goto L259
L258:
	;
	v1241 = v1239
	goto L259
L259:
	;
	if v1241 <= int32(1) {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1244 = int32(1)
	goto L262
L261:
	;
	v1244 = v1241
	goto L262
L262:
	;
	v1245 = int32(8)
	v1250 = int32(0)
	goto L263
L263:
	;
	v1257 = v1250 << (uint(int32(2)) % 32)
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1227+v1245+v1257)))
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1215+v1245+v1257)))
	v1262 = v1259 & v1261
	v1264 = base.B2i32(v1262 != int32(0))
	if v1262 != 0 {
		v1273 = v1264
		goto L255
	} else {
		goto L265
	}
L264:
	;
	v1273 = v1264
	goto L255
L265:
	;
	v1266 = v1250 + int32(1)
	if v1266 != v1244 {
		v1250 = v1266
		goto L263
	} else {
		goto L266
	}
L266:
	;
	goto L264
L267:
	;
	goto L248
L268:
	;
	v1286 = F_build_partition_pathkeys(m, l0, l1, int32(1), v28+int32(255))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L18
	} else {
		goto L269
	}
L269:
	;
	v1291 = F_build_partition_pathkeys(m, l0, l1, int32(-1), v28+int32(254))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L18
	} else {
		goto L270
	}
L270:
	;
	v1296 = v1286
	v1297 = v1291
	goto L242
L271:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v744)+4))
	if v1300 <= int32(0) {
		goto L240
	} else {
		goto L272
	}
L272:
	;
	v1304 = v28 + int32(216)
	v1306 = v28 + int32(232)
	v1308 = v28 + int32(248)
	v1322 = int32(0)
	goto L273
L273:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v744)+12))
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1335+v1322<<(uint(int32(2))%32))))
	v1340 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+248)) = v1340
	v1342 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+240)) = v1342
	*(*int32)(unsafe.Add(mBase, uint32(v28)+232)) = v1340
	*(*int64)(unsafe.Add(mBase, uint32(v28)+224)) = v1342
	*(*int32)(unsafe.Add(mBase, uint32(v28)+216)) = v1340
	*(*int64)(unsafe.Add(mBase, uint32(v28)+208)) = v1342
	v1352 = int32(1)
	if v1339 == v1296 {
		goto L285
	} else {
		goto L286
	}
L274:
	;
	goto L240
L275:
	;
	v2331 = v1322 + int32(1)
	v2332 = *(*int32)(unsafe.Add(mBase, uint32(v744)+4))
	if v2331 < v2332 {
		v1322 = v2331
		goto L273
	} else {
		goto L549
	}
L276:
	;
	F_add_path(m, l1, v2302)
	mBase = m.M
	v2304 = m.ExcPending
	if v2304 != 0 {
		goto L18
	} else {
		goto L548
	}
L277:
	;
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+72)) = v2229
	v2231 = *(*int64)(unsafe.Add(mBase, uint32(v28)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+64)) = v2231
	v2235 = int32(0)
	v2239 = F_create_append_path(m, l0, l1, v28-int32(-64), v1339, v2235, v2235, v2235, float64(-1))
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		goto L18
	} else {
		goto L539
	}
L278:
	;
	v2183 = F_create_merge_append_path(m, l0, l1, v2165, v2162, v1339)
	mBase = m.M
	v2184 = m.ExcPending
	if v2184 != 0 {
		goto L18
	} else {
		goto L530
	}
L279:
	;
	if v1591 != 0 {
		goto L527
	} else {
		goto L528
	}
L280:
	;
	v2126 = int32(0)
	if v1577 != 0 {
		v2208 = v2126
		v2209 = v2126
		goto L277
	} else {
		goto L526
	}
L281:
	;
	v1593 = int32(0)
	if v1589 == v1590 {
		v2137 = v1593
		v2140 = v1593
		goto L279
	} else {
		goto L368
	}
L282:
	;
	v1583 = int32(-1)
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1585 = int32(1)
	v1589 = v1584 - v1585
	v1590 = v1583
	v1591 = v1585
	v1592 = v1583
	goto L281
L283:
	;
	if l2 == int32(0) {
		goto L280
	} else {
		goto L367
	}
L284:
	;
	if v1405 != 0 {
		v1577 = v1352
		goto L283
	} else {
		goto L302
	}
L285:
	;
	v1405 = int32(1)
	goto L284
L286:
	;
	goto L287
L287:
	;
	v1361 = v1340
	goto L289
L288:
	;
	v1405 = v1397
	goto L284
L289:
	;
	v1365 = int32(0)
	if v1339 == v1365 {
		v1375 = v1365
		goto L291
	} else {
		goto L292
	}
L290:
	;
	v1397 = int32(0)
	goto L288
L291:
	;
	if v1296 != 0 {
		goto L295
	} else {
		goto L296
	}
L292:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+4))
	if v1369 <= v1361 {
		v1375 = int32(0)
		goto L291
	} else {
		goto L293
	}
L293:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+12))
	v1375 = v1371 + v1361<<(uint(int32(2))%32)
	goto L291
L294:
	;
	v1381 = base.B2i32(v1375 == int32(0))
	if v1375 == int32(0) {
		v1397 = v1381
		goto L288
	} else {
		goto L299
	}
L295:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+4))
	if v1361 < v1376 {
		goto L294
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	v1405 = base.B2i32(v1375 == int32(0))
	goto L284
L298:
	;
	goto L297
L299:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+12))
	if v1384 == int32(0) {
		v1397 = v1381
		goto L288
	} else {
		goto L300
	}
L300:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1375)))
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1361<<(uint(int32(2))%32)+v1384)))
	if v1391 == v1393 {
		v1361 = v1361 + int32(1)
		goto L289
	} else {
		goto L301
	}
L301:
	;
	goto L290
L302:
	;
	v1406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+255)))
	if v1406 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	if v1296 == v1339 {
		goto L307
	} else {
		goto L308
	}
L304:
	;
	goto L305
L305:
	;
	if v1339 == v1297 {
		goto L326
	} else {
		goto L327
	}
L306:
	;
	if v1461 != 0 {
		v1577 = v1352
		goto L283
	} else {
		goto L324
	}
L307:
	;
	v1461 = int32(1)
	goto L306
L308:
	;
	goto L309
L309:
	;
	v1417 = int32(0)
	goto L311
L310:
	;
	v1461 = v1453
	goto L306
L311:
	;
	v1421 = int32(0)
	if v1296 == v1421 {
		v1431 = v1421
		goto L313
	} else {
		goto L314
	}
L312:
	;
	v1453 = int32(0)
	goto L310
L313:
	;
	if v1339 != 0 {
		goto L317
	} else {
		goto L318
	}
L314:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+4))
	if v1425 <= v1417 {
		v1431 = int32(0)
		goto L313
	} else {
		goto L315
	}
L315:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+12))
	v1431 = v1427 + v1417<<(uint(int32(2))%32)
	goto L313
L316:
	;
	v1437 = base.B2i32(v1431 == int32(0))
	if v1431 == int32(0) {
		v1453 = v1437
		goto L310
	} else {
		goto L321
	}
L317:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+4))
	if v1417 < v1432 {
		goto L316
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	v1461 = base.B2i32(v1431 == int32(0))
	goto L306
L320:
	;
	goto L319
L321:
	;
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+12))
	if v1440 == int32(0) {
		v1453 = v1437
		goto L310
	} else {
		goto L322
	}
L322:
	;
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1431)))
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v1417<<(uint(int32(2))%32)+v1440)))
	if v1447 == v1449 {
		v1417 = v1417 + int32(1)
		goto L311
	} else {
		goto L323
	}
L323:
	;
	goto L312
L324:
	;
	goto L305
L325:
	;
	if v1514 == int32(0) {
		goto L343
	} else {
		goto L344
	}
L326:
	;
	v1514 = int32(1)
	goto L325
L327:
	;
	goto L328
L328:
	;
	v1470 = int32(0)
	goto L330
L329:
	;
	v1514 = v1506
	goto L325
L330:
	;
	v1474 = int32(0)
	if v1339 == v1474 {
		v1484 = v1474
		goto L332
	} else {
		goto L333
	}
L331:
	;
	v1506 = int32(0)
	goto L329
L332:
	;
	if v1297 != 0 {
		goto L336
	} else {
		goto L337
	}
L333:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+4))
	if v1478 <= v1470 {
		v1484 = int32(0)
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+12))
	v1484 = v1480 + v1470<<(uint(int32(2))%32)
	goto L332
L335:
	;
	v1490 = base.B2i32(v1484 == int32(0))
	if v1484 == int32(0) {
		v1506 = v1490
		goto L329
	} else {
		goto L340
	}
L336:
	;
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v1297)+4))
	if v1470 < v1485 {
		goto L335
	} else {
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	v1514 = base.B2i32(v1484 == int32(0))
	goto L325
L339:
	;
	goto L338
L340:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1297)+12))
	if v1493 == int32(0) {
		v1506 = v1490
		goto L329
	} else {
		goto L341
	}
L341:
	;
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1484)))
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1470<<(uint(int32(2))%32)+v1493)))
	if v1500 == v1502 {
		v1470 = v1470 + int32(1)
		goto L330
	} else {
		goto L342
	}
L342:
	;
	goto L331
L343:
	;
	v1517 = int32(0)
	v1518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+254)))
	if v1518 != 0 {
		v1577 = v1517
		goto L283
	} else {
		goto L346
	}
L344:
	;
	goto L345
L345:
	;
	if l2 != 0 {
		goto L282
	} else {
		goto L366
	}
L346:
	;
	if v1297 == v1339 {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	if v1571 == int32(0) {
		v1577 = v1517
		goto L283
	} else {
		goto L365
	}
L348:
	;
	v1571 = int32(1)
	goto L347
L349:
	;
	goto L350
L350:
	;
	v1527 = int32(0)
	goto L352
L351:
	;
	v1571 = v1563
	goto L347
L352:
	;
	v1531 = int32(0)
	if v1297 == v1531 {
		v1541 = v1531
		goto L354
	} else {
		goto L355
	}
L353:
	;
	v1563 = int32(0)
	goto L351
L354:
	;
	if v1339 != 0 {
		goto L358
	} else {
		goto L359
	}
L355:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1297)+4))
	if v1535 <= v1527 {
		v1541 = int32(0)
		goto L354
	} else {
		goto L356
	}
L356:
	;
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1297)+12))
	v1541 = v1537 + v1527<<(uint(int32(2))%32)
	goto L354
L357:
	;
	v1547 = base.B2i32(v1541 == int32(0))
	if v1541 == int32(0) {
		v1563 = v1547
		goto L351
	} else {
		goto L362
	}
L358:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+4))
	if v1527 < v1542 {
		goto L357
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	v1571 = base.B2i32(v1541 == int32(0))
	goto L347
L361:
	;
	goto L360
L362:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+12))
	if v1550 == int32(0) {
		v1563 = v1547
		goto L351
	} else {
		goto L363
	}
L363:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1541)))
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1527<<(uint(int32(2))%32)+v1550)))
	if v1557 == v1559 {
		v1527 = v1527 + int32(1)
		goto L352
	} else {
		goto L364
	}
L364:
	;
	goto L353
L365:
	;
	goto L345
L366:
	;
	v1575 = int32(0)
	v2208 = v1575
	v2209 = v1575
	goto L277
L367:
	;
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1589 = int32(0)
	v1590 = v1580
	v1591 = v1577
	v1592 = int32(1)
	goto L281
L368:
	;
	v1602 = v1593
	v1604 = v1589
	v1605 = v1593
	goto L369
L369:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1621+v1604<<(uint(int32(2))%32))))
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1625)+44))
	v1627 = int32(0)
	if v1626 == v1627 {
		goto L372
	} else {
		goto L373
	}
L370:
	;
	v2137 = v2123
	v2140 = v2082
	goto L279
L371:
	;
	v1776 = int32(0)
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v1625)+44))
	if v1777 == v1776 {
		goto L415
	} else {
		goto L416
	}
L372:
	;
	v1775 = int32(0)
	goto L371
L373:
	;
	goto L374
L374:
	;
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1626)+4))
	if int32(0) < v1641 {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v1651 = v1627
	v1654 = v1627
	goto L378
L376:
	;
	v1756 = v1627
	goto L377
L377:
	;
	v1775 = v1756
	goto L371
L378:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1626)+12))
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v1657+v1654<<(uint(int32(2))%32))))
	goto L382
L379:
	;
	v1756 = v1739
	goto L377
L380:
	;
	v1746 = v1654 + int32(1)
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v1626)+4))
	if v1746 < v1747 {
		v1651 = v1739
		v1654 = v1746
		goto L378
	} else {
		goto L413
	}
L382:
	;
	goto L383
L383:
	;
	if v1651 != 0 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1665 = F_compare_path_costs(m, v1651, v1661, v1627)
	mBase = m.M
	if v1665 <= int32(0) {
		v1739 = v1651
		goto L380
	} else {
		goto L388
	}
L386:
	;
	goto L387
L387:
	;
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+64))
	if v1339 == v1668 {
		goto L389
	} else {
		goto L390
	}
L388:
	;
	goto L387
L389:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+16))
	if v1726 != 0 {
		goto L407
	} else {
		goto L408
	}
L390:
	;
	v1676 = int32(0)
	goto L391
L391:
	;
	v1684 = int32(0)
	if v1339 == v1684 {
		v1694 = v1684
		goto L393
	} else {
		goto L394
	}
L392:
	;
	if v1694 != 0 {
		v1739 = v1651
		goto L380
	} else {
		goto L406
	}
L393:
	;
	if v1668 != 0 {
		goto L397
	} else {
		goto L398
	}
L394:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+4))
	if v1688 <= v1676 {
		v1694 = int32(0)
		goto L393
	} else {
		goto L395
	}
L395:
	;
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+12))
	v1694 = v1690 + v1676<<(uint(int32(2))%32)
	goto L393
L396:
	;
	if v1694 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L397:
	;
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v1668)+4))
	if v1676 < v1695 {
		goto L396
	} else {
		goto L400
	}
L398:
	;
	goto L399
L399:
	;
	if v1694 == int32(0) {
		goto L389
	} else {
		goto L401
	}
L400:
	;
	goto L399
L401:
	;
	v1739 = v1651
	goto L380
L402:
	;
	goto L392
L403:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v1668)+12))
	if v1701 == int32(0) {
		goto L402
	} else {
		goto L404
	}
L404:
	;
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v1694)))
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v1701+v1676<<(uint(int32(2))%32))))
	if v1708 == v1710 {
		v1676 = v1676 + int32(1)
		goto L391
	} else {
		goto L405
	}
L405:
	;
	v1739 = v1651
	goto L380
L406:
	;
	goto L389
L407:
	;
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v1726)+4))
	v1729 = v1727
	goto L409
L408:
	;
	v1729 = int32(0)
	goto L409
L409:
	;
	v1730 = F_bms_is_subset(m, v1729, v1627)
	mBase = m.M
	if v1730 != 0 {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	v1731 = v1661
	goto L412
L411:
	;
	v1731 = v1651
	goto L412
L412:
	;
	v1739 = v1731
	goto L380
L413:
	;
	goto L379
L414:
	;
	if v1926 != 0 {
		goto L457
	} else {
		goto L458
	}
L415:
	;
	v1926 = int32(0)
	goto L414
L416:
	;
	goto L417
L417:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v1777)+4))
	if int32(0) < v1792 {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	v1802 = v1776
	v1805 = v1776
	goto L421
L419:
	;
	v1907 = v1776
	goto L420
L420:
	;
	v1926 = v1907
	goto L414
L421:
	;
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1777)+12))
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1808+v1805<<(uint(int32(2))%32))))
	goto L425
L422:
	;
	v1907 = v1890
	goto L420
L423:
	;
	v1897 = v1805 + int32(1)
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1777)+4))
	if v1897 < v1898 {
		v1802 = v1890
		v1805 = v1897
		goto L421
	} else {
		goto L456
	}
L425:
	;
	goto L426
L426:
	;
	if v1802 != 0 {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1816 = F_compare_path_costs(m, v1802, v1812, int32(1))
	mBase = m.M
	if v1816 <= int32(0) {
		v1890 = v1802
		goto L423
	} else {
		goto L431
	}
L429:
	;
	goto L430
L430:
	;
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1812)+64))
	if v1339 == v1819 {
		goto L432
	} else {
		goto L433
	}
L431:
	;
	goto L430
L432:
	;
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1812)+16))
	if v1877 != 0 {
		goto L450
	} else {
		goto L451
	}
L433:
	;
	v1827 = int32(0)
	goto L434
L434:
	;
	v1835 = int32(0)
	if v1339 == v1835 {
		v1845 = v1835
		goto L436
	} else {
		goto L437
	}
L435:
	;
	if v1845 != 0 {
		v1890 = v1802
		goto L423
	} else {
		goto L449
	}
L436:
	;
	if v1819 != 0 {
		goto L440
	} else {
		goto L441
	}
L437:
	;
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+4))
	if v1839 <= v1827 {
		v1845 = int32(0)
		goto L436
	} else {
		goto L438
	}
L438:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+12))
	v1845 = v1841 + v1827<<(uint(int32(2))%32)
	goto L436
L439:
	;
	if v1845 == int32(0) {
		goto L445
	} else {
		goto L446
	}
L440:
	;
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1819)+4))
	if v1827 < v1846 {
		goto L439
	} else {
		goto L443
	}
L441:
	;
	goto L442
L442:
	;
	if v1845 == int32(0) {
		goto L432
	} else {
		goto L444
	}
L443:
	;
	goto L442
L444:
	;
	v1890 = v1802
	goto L423
L445:
	;
	goto L435
L446:
	;
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v1819)+12))
	if v1852 == int32(0) {
		goto L445
	} else {
		goto L447
	}
L447:
	;
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1845)))
	v1861 = *(*int32)(unsafe.Add(mBase, uint32(v1852+v1827<<(uint(int32(2))%32))))
	if v1859 == v1861 {
		v1827 = v1827 + int32(1)
		goto L434
	} else {
		goto L448
	}
L448:
	;
	v1890 = v1802
	goto L423
L449:
	;
	goto L432
L450:
	;
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1877)+4))
	v1880 = v1878
	goto L452
L451:
	;
	v1880 = int32(0)
	goto L452
L452:
	;
	v1881 = F_bms_is_subset(m, v1880, v1776)
	mBase = m.M
	if v1881 != 0 {
		goto L453
	} else {
		goto L454
	}
L453:
	;
	v1882 = v1812
	goto L455
L454:
	;
	v1882 = v1802
	goto L455
L455:
	;
	v1890 = v1882
	goto L423
L456:
	;
	goto L422
L457:
	;
	v1927 = v1775
	goto L459
L458:
	;
	v1927 = v1776
	goto L459
L459:
	;
	if v1927 == int32(0) {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v1625)+60))
	v1931 = v1930
	v1932 = v1930
	goto L462
L461:
	;
	v1931 = v1926
	v1932 = v1775
	goto L462
L462:
	;
	v1933 = int32(0)
	v1934 = *(*float64)(unsafe.Add(mBase, uint32(l0)+312))
	if base.F64_gt(v1934, float64(0)) == v1933 {
		v2081 = v1933
		v2082 = v1605
		goto L463
	} else {
		goto L464
	}
L463:
	;
	if v1591 != 0 {
		goto L511
	} else {
		goto L512
	}
L464:
	;
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v1625)+44))
	if base.F64_ge(v1934, float64(1)) != 0 {
		goto L465
	} else {
		goto L466
	}
L465:
	;
	v1942 = *(*float64)(unsafe.Add(mBase, uint32(v1931)+32))
	v1944 = base.F64_div(v1934, v1942)
	goto L467
L466:
	;
	v1944 = v1934
	goto L467
L467:
	;
	v1945 = int32(0)
	if v1939 == v1945 {
		goto L469
	} else {
		goto L470
	}
L468:
	;
	if v2076 == int32(0) {
		goto L507
	} else {
		goto L508
	}
L469:
	;
	v2076 = int32(0)
	goto L468
L470:
	;
	goto L471
L471:
	;
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(v1939)+4))
	if int32(0) < v1956 {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	v1964 = v1945
	v1967 = v1945
	goto L475
L473:
	;
	v2059 = v1945
	goto L474
L474:
	;
	v2076 = v2059
	goto L468
L475:
	;
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(v1939)+12))
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(v1970+v1967<<(uint(int32(2))%32))))
	if v1964 != 0 {
		goto L478
	} else {
		goto L479
	}
L476:
	;
	v2059 = v2044
	goto L474
L477:
	;
	v2051 = v1967 + int32(1)
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v1939)+4))
	if v2051 < v2052 {
		v1964 = v2044
		v1967 = v2051
		goto L475
	} else {
		goto L506
	}
L478:
	;
	v1975 = F_compare_fractional_path_costs(m, v1964, v1974, v1944)
	mBase = m.M
	if v1975 <= int32(0) {
		v2044 = v1964
		goto L477
	} else {
		goto L481
	}
L479:
	;
	goto L480
L480:
	;
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+64))
	if v1339 == v1978 {
		goto L482
	} else {
		goto L483
	}
L481:
	;
	goto L480
L482:
	;
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+16))
	if v2032 != 0 {
		goto L500
	} else {
		goto L501
	}
L483:
	;
	v1984 = int32(0)
	goto L484
L484:
	;
	v1992 = int32(0)
	if v1339 == v1992 {
		v2002 = v1992
		goto L486
	} else {
		goto L487
	}
L485:
	;
	if v2002 != 0 {
		v2044 = v1964
		goto L477
	} else {
		goto L499
	}
L486:
	;
	if v1978 != 0 {
		goto L490
	} else {
		goto L491
	}
L487:
	;
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+4))
	if v1996 <= v1984 {
		v2002 = int32(0)
		goto L486
	} else {
		goto L488
	}
L488:
	;
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+12))
	v2002 = v1998 + v1984<<(uint(int32(2))%32)
	goto L486
L489:
	;
	if v2002 == int32(0) {
		goto L495
	} else {
		goto L496
	}
L490:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v1978)+4))
	if v1984 < v2003 {
		goto L489
	} else {
		goto L493
	}
L491:
	;
	goto L492
L492:
	;
	if v2002 == int32(0) {
		goto L482
	} else {
		goto L494
	}
L493:
	;
	goto L492
L494:
	;
	v2044 = v1964
	goto L477
L495:
	;
	goto L485
L496:
	;
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v1978)+12))
	if v2009 == int32(0) {
		goto L495
	} else {
		goto L497
	}
L497:
	;
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v2002)))
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(v2009+v1984<<(uint(int32(2))%32))))
	if v2016 == v2018 {
		v1984 = v1984 + int32(1)
		goto L484
	} else {
		goto L498
	}
L498:
	;
	v2044 = v1964
	goto L477
L499:
	;
	goto L482
L500:
	;
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v2032)+4))
	v2035 = v2033
	goto L502
L501:
	;
	v2035 = int32(0)
	goto L502
L502:
	;
	v2037 = F_bms_is_subset(m, v2035, int32(0))
	mBase = m.M
	if v2037 != 0 {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	v2038 = v1974
	goto L505
L504:
	;
	v2038 = v1964
	goto L505
L505:
	;
	v2044 = v2038
	goto L477
L506:
	;
	goto L476
L507:
	;
	v2081 = v1931
	v2082 = v1605
	goto L463
L508:
	;
	goto L509
L509:
	;
	v2081 = v2076
	v2082 = base.B2i32(v1931 != v2076) | v1605
	goto L463
L510:
	;
	v2123 = v1602 | base.B2i32(v1931 != v1932)
	v2124 = v1604 + v1592
	if v1590 != v2124 {
		v1602 = v2123
		v1604 = v2124
		v1605 = v2082
		goto L369
	} else {
		goto L525
	}
L511:
	;
	v2084 = F_get_singleton_append_subpath(m, v1932, v1308)
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L18
	} else {
		goto L514
	}
L512:
	;
	goto L513
L513:
	;
	F_accumulate_append_subpath(m, v1932, v28+int32(240), int32(0), v1308)
	mBase = m.M
	v2108 = m.ExcPending
	if v2108 != 0 {
		goto L18
	} else {
		goto L521
	}
L514:
	;
	v2086 = F_get_singleton_append_subpath(m, v1931, v1306)
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L18
	} else {
		goto L515
	}
L515:
	;
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(v28)+240))
	v2089 = F_lappend(m, v2088, v2084)
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L18
	} else {
		goto L516
	}
L516:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+240)) = v2089
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v28)+224))
	v2093 = F_lappend(m, v2092, v2086)
	mBase = m.M
	v2094 = m.ExcPending
	if v2094 != 0 {
		goto L18
	} else {
		goto L517
	}
L517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+224)) = v2093
	if v2081 == int32(0) {
		goto L510
	} else {
		goto L518
	}
L518:
	;
	v2098 = F_get_singleton_append_subpath(m, v2081, v1304)
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L18
	} else {
		goto L519
	}
L519:
	;
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v28)+208))
	v2101 = F_lappend(m, v2100, v2098)
	mBase = m.M
	v2102 = m.ExcPending
	if v2102 != 0 {
		goto L18
	} else {
		goto L520
	}
L520:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+208)) = v2101
	goto L510
L521:
	;
	F_accumulate_append_subpath(m, v1931, v28+int32(224), int32(0), v1306)
	mBase = m.M
	v2113 = m.ExcPending
	if v2113 != 0 {
		goto L18
	} else {
		goto L522
	}
L522:
	;
	if v2081 == int32(0) {
		goto L510
	} else {
		goto L523
	}
L523:
	;
	F_accumulate_append_subpath(m, v2081, v28+int32(208), int32(0), v1304)
	mBase = m.M
	v2120 = m.ExcPending
	if v2120 != 0 {
		goto L18
	} else {
		goto L524
	}
L524:
	;
	goto L510
L525:
	;
	goto L370
L526:
	;
	v2162 = v2126
	v2164 = v2126
	v2165 = v2126
	v2167 = v2126
	goto L278
L527:
	;
	v2208 = v2140
	v2209 = v2137
	goto L277
L528:
	;
	goto L529
L529:
	;
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v28)+240))
	v2162 = v2156
	v2164 = v2137
	v2165 = v2157
	v2167 = v2140
	goto L278
L530:
	;
	F_add_path(m, l1, v2183)
	mBase = m.M
	v2186 = m.ExcPending
	if v2186 != 0 {
		goto L18
	} else {
		goto L531
	}
L531:
	;
	if v2164&int32(1) != 0 {
		goto L532
	} else {
		goto L533
	}
L532:
	;
	v2189 = *(*int32)(unsafe.Add(mBase, uint32(v28)+224))
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(v28)+232))
	v2191 = F_create_merge_append_path(m, l0, l1, v2189, v2190, v1339)
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L18
	} else {
		goto L535
	}
L533:
	;
	goto L534
L534:
	;
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(v28)+208))
	v2196 = int32(0)
	if base.B2i32(v2195 != v2196)&v2167 == v2196 {
		goto L275
	} else {
		goto L537
	}
L535:
	;
	F_add_path(m, l1, v2191)
	mBase = m.M
	v2194 = m.ExcPending
	if v2194 != 0 {
		goto L18
	} else {
		goto L536
	}
L536:
	;
	goto L534
L537:
	;
	v2201 = *(*int32)(unsafe.Add(mBase, uint32(v28)+216))
	v2202 = F_create_merge_append_path(m, l0, l1, v2195, v2201, v1339)
	mBase = m.M
	v2203 = m.ExcPending
	if v2203 != 0 {
		goto L18
	} else {
		goto L538
	}
L538:
	;
	v2302 = v2202
	goto L276
L539:
	;
	F_add_path(m, l1, v2239)
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L18
	} else {
		goto L540
	}
L540:
	;
	if v2209&int32(1) != 0 {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(v28)+232))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+56)) = v2245
	v2247 = *(*int64)(unsafe.Add(mBase, uint32(v28)+224))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+48)) = v2247
	v2251 = int32(0)
	v2255 = F_create_append_path(m, l0, l1, v28+int32(48), v1339, v2251, v2251, v2251, float64(-1))
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		goto L18
	} else {
		goto L544
	}
L542:
	;
	goto L543
L543:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v28)+208))
	v2260 = int32(0)
	if base.B2i32(v2259 != v2260)&v2208 == v2260 {
		goto L275
	} else {
		goto L546
	}
L544:
	;
	F_add_path(m, l1, v2255)
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L18
	} else {
		goto L545
	}
L545:
	;
	goto L543
L546:
	;
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(v28)+216))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+40)) = v2265
	v2267 = *(*int64)(unsafe.Add(mBase, uint32(v28)+208))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = v2267
	v2271 = int32(0)
	v2275 = F_create_append_path(m, l0, l1, v28+int32(32), v1339, v2271, v2271, v2271, float64(-1))
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L18
	} else {
		goto L547
	}
L547:
	;
	v2302 = v2275
	goto L276
L548:
	;
	goto L275
L549:
	;
	goto L274
L550:
	;
	if l2 == int32(0) {
		goto L728
	} else {
		goto L729
	}
L551:
	;
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(v747)+4))
	if v2361 <= int32(0) {
		goto L550
	} else {
		goto L552
	}
L552:
	;
	v2378 = int32(0)
	goto L553
L553:
	;
	v2392 = *(*int32)(unsafe.Add(mBase, uint32(v747)+12))
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(v2392+v2378<<(uint(int32(2))%32))))
	v2397 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+248)) = v2397
	*(*int64)(unsafe.Add(mBase, uint32(v28)+240)) = int64(0)
	if l2 == v2397 {
		goto L556
	} else {
		goto L557
	}
L554:
	;
	goto L550
L555:
	;
	v3020 = v2378 + int32(1)
	v3021 = *(*int32)(unsafe.Add(mBase, uint32(v747)+4))
	if v3020 < v3021 {
		v2378 = v3020
		goto L553
	} else {
		goto L727
	}
L556:
	;
	v2980 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v2980
	v2982 = *(*int64)(unsafe.Add(mBase, uint32(v28)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = v2982
	v2986 = int32(0)
	v2990 = F_create_append_path(m, l0, l1, v28+int32(16), v2986, v2396, v2986, v2986, float64(-1))
	mBase = m.M
	v2991 = m.ExcPending
	if v2991 != 0 {
		goto L18
	} else {
		goto L725
	}
L557:
	;
	v2403 = int32(0)
	v2404 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v2404 <= v2403 {
		goto L556
	} else {
		goto L558
	}
L558:
	;
	v2413 = v2403
	goto L559
L559:
	;
	v2432 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v2436 = *(*int32)(unsafe.Add(mBase, uint32(v2432+v2413<<(uint(int32(2))%32))))
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(v2436)+44))
	if v2437 == int32(0) {
		goto L555
	} else {
		goto L561
	}
L560:
	;
	goto L556
L561:
	;
	v2440 = int32(0)
	if v2437 == v2440 {
		goto L563
	} else {
		goto L564
	}
L562:
	;
	v2590 = *(*int32)(unsafe.Add(mBase, uint32(v2589)+16))
	if v2590 != 0 {
		goto L605
	} else {
		goto L606
	}
L563:
	;
	v2589 = int32(0)
	goto L562
L564:
	;
	goto L565
L565:
	;
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(v2437)+4))
	if int32(0) < v2455 {
		goto L566
	} else {
		goto L567
	}
L566:
	;
	v2465 = v2440
	v2468 = v2440
	goto L569
L567:
	;
	v2570 = v2440
	goto L568
L568:
	;
	v2589 = v2570
	goto L562
L569:
	;
	v2471 = *(*int32)(unsafe.Add(mBase, uint32(v2437)+12))
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(v2471+v2468<<(uint(int32(2))%32))))
	goto L573
L570:
	;
	v2570 = v2553
	goto L568
L571:
	;
	v2560 = v2468 + int32(1)
	v2561 = *(*int32)(unsafe.Add(mBase, uint32(v2437)+4))
	if v2560 < v2561 {
		v2465 = v2553
		v2468 = v2560
		goto L569
	} else {
		goto L604
	}
L573:
	;
	goto L574
L574:
	;
	if v2465 != 0 {
		goto L576
	} else {
		goto L577
	}
L576:
	;
	v2479 = F_compare_path_costs(m, v2465, v2475, int32(1))
	mBase = m.M
	if v2479 <= int32(0) {
		v2553 = v2465
		goto L571
	} else {
		goto L579
	}
L577:
	;
	goto L578
L578:
	;
	v2482 = *(*int32)(unsafe.Add(mBase, uint32(v2475)+64))
	if v2440 == v2482 {
		goto L580
	} else {
		goto L581
	}
L579:
	;
	goto L578
L580:
	;
	v2540 = *(*int32)(unsafe.Add(mBase, uint32(v2475)+16))
	if v2540 != 0 {
		goto L598
	} else {
		goto L599
	}
L581:
	;
	goto L582
L582:
	;
	goto L584
L583:
	;
	goto L597
L584:
	;
	if v2482 != 0 {
		goto L588
	} else {
		goto L589
	}
L587:
	;
	goto L593
L588:
	;
	v2509 = *(*int32)(unsafe.Add(mBase, uint32(v2482)+4))
	if int32(0) < v2509 {
		goto L587
	} else {
		goto L591
	}
L589:
	;
	goto L590
L590:
	;
	goto L580
L591:
	;
	goto L590
L593:
	;
	goto L583
L597:
	;
	goto L580
L598:
	;
	v2541 = *(*int32)(unsafe.Add(mBase, uint32(v2540)+4))
	v2543 = v2541
	goto L600
L599:
	;
	v2543 = int32(0)
	goto L600
L600:
	;
	v2544 = F_bms_is_subset(m, v2543, v2396)
	mBase = m.M
	if v2544 != 0 {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	v2545 = v2475
	goto L603
L602:
	;
	v2545 = v2465
	goto L603
L603:
	;
	v2553 = v2545
	goto L571
L604:
	;
	goto L570
L605:
	;
	v2591 = *(*int32)(unsafe.Add(mBase, uint32(v2590)+4))
	v2593 = v2591
	goto L607
L606:
	;
	v2593 = int32(0)
	goto L607
L607:
	;
	v2594 = int32(0)
	if base.B2i32(v2593 == v2594)|base.B2i32(v2396 == v2594) != 0 {
		v2640 = base.B2i32(v2593|v2396 == v2594)
		goto L609
	} else {
		goto L610
	}
L608:
	;
	if v2640 == int32(0) {
		goto L619
	} else {
		goto L620
	}
L609:
	;
	goto L608
L610:
	;
	v2608 = *(*int32)(unsafe.Add(mBase, uint32(v2593)+4))
	v2609 = *(*int32)(unsafe.Add(mBase, uint32(v2396)+4))
	if v2608 != v2609 {
		v2640 = int32(0)
		goto L609
	} else {
		goto L611
	}
L611:
	;
	v2611 = int32(1)
	if v2608 <= v2611 {
		goto L612
	} else {
		goto L613
	}
L612:
	;
	v2614 = v2611
	goto L614
L613:
	;
	v2614 = v2608
	goto L614
L614:
	;
	v2615 = int32(8)
	v2620 = int32(0)
	goto L615
L615:
	;
	v2628 = v2620 << (uint(int32(2)) % 32)
	v2630 = *(*int32)(unsafe.Add(mBase, uint32(v2593+v2615+v2628)))
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(v2396+v2615+v2628)))
	v2633 = base.B2i32(v2630 == v2632)
	if v2630 != v2632 {
		v2640 = v2633
		goto L609
	} else {
		goto L617
	}
L616:
	;
	v2640 = v2633
	goto L609
L617:
	;
	v2636 = v2620 + int32(1)
	if v2636 != v2614 {
		v2620 = v2636
		goto L615
	} else {
		goto L618
	}
L618:
	;
	goto L616
L619:
	;
	v2647 = *(*int32)(unsafe.Add(mBase, uint32(v2436)+44))
	if v2647 == int32(0) {
		goto L555
	} else {
		goto L622
	}
L620:
	;
	v2925 = v2589
	goto L621
L621:
	;
	F_accumulate_append_subpath(m, v2925, v28+int32(240), int32(0), v28+int32(248))
	mBase = m.M
	v2950 = m.ExcPending
	if v2950 != 0 {
		goto L18
	} else {
		goto L723
	}
L622:
	;
	v2650 = int32(0)
	v2651 = *(*int32)(unsafe.Add(mBase, uint32(v2647)+4))
	if v2651 <= v2650 {
		goto L555
	} else {
		goto L623
	}
L623:
	;
	v2658 = v2650
	v2667 = v2440
	goto L624
L624:
	;
	v2679 = *(*int32)(unsafe.Add(mBase, uint32(v2647)+12))
	v2683 = *(*int32)(unsafe.Add(mBase, uint32(v2679+v2667<<(uint(int32(2))%32))))
	v2684 = *(*int32)(unsafe.Add(mBase, uint32(v2683)+16))
	if v2684 != 0 {
		goto L627
	} else {
		goto L628
	}
L625:
	;
	if v2913 == int32(0) {
		goto L555
	} else {
		goto L722
	}
L626:
	;
	v2916 = v2667 + int32(1)
	v2917 = *(*int32)(unsafe.Add(mBase, uint32(v2647)+4))
	if v2916 < v2917 {
		v2658 = v2913
		v2667 = v2916
		goto L624
	} else {
		goto L721
	}
L627:
	;
	v2685 = *(*int32)(unsafe.Add(mBase, uint32(v2684)+4))
	v2687 = v2685
	goto L629
L628:
	;
	v2687 = int32(0)
	goto L629
L629:
	;
	v2688 = int32(0)
	if v2687 == v2688 {
		goto L631
	} else {
		goto L632
	}
L630:
	;
	if v2741 == int32(0) {
		goto L644
	} else {
		goto L645
	}
L631:
	;
	v2741 = int32(1)
	goto L630
L632:
	;
	goto L633
L633:
	;
	if v2396 == int32(0) {
		v2734 = v2688
		goto L634
	} else {
		goto L635
	}
L634:
	;
	v2741 = v2734
	goto L630
L635:
	;
	v2697 = *(*int32)(unsafe.Add(mBase, uint32(v2687)+4))
	v2698 = *(*int32)(unsafe.Add(mBase, uint32(v2396)+4))
	if v2698 < v2697 {
		v2734 = v2688
		goto L634
	} else {
		goto L636
	}
L636:
	;
	v2700 = int32(1)
	if v2697 <= v2700 {
		goto L637
	} else {
		goto L638
	}
L637:
	;
	v2703 = v2700
	goto L639
L638:
	;
	v2703 = v2697
	goto L639
L639:
	;
	v2704 = int32(8)
	v2709 = int32(0)
	goto L640
L640:
	;
	v2716 = v2709 << (uint(int32(2)) % 32)
	v2718 = *(*int32)(unsafe.Add(mBase, uint32(v2687+v2704+v2716)))
	v2720 = *(*int32)(unsafe.Add(mBase, uint32(v2396+v2704+v2716)))
	v2723 = v2718 & (v2720 ^ int32(-1))
	v2725 = base.B2i32(v2723 == int32(0))
	if v2723 != 0 {
		v2734 = v2725
		goto L634
	} else {
		goto L642
	}
L641:
	;
	v2734 = v2725
	goto L634
L642:
	;
	v2727 = v2709 + int32(1)
	if v2727 != v2703 {
		v2709 = v2727
		goto L640
	} else {
		goto L643
	}
L643:
	;
	goto L641
L644:
	;
	v2913 = v2658
	goto L626
L645:
	;
	goto L646
L646:
	;
	if v2658 == int32(0) {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	v2797 = *(*int32)(unsafe.Add(mBase, uint32(v2683)+16))
	if v2797 != 0 {
		goto L674
	} else {
		goto L675
	}
L648:
	;
	v2751 = *(*int32)(unsafe.Add(mBase, uint32(v2658)+40))
	v2752 = *(*int32)(unsafe.Add(mBase, uint32(v2683)+40))
	if v2751 != v2752 {
		goto L650
	} else {
		goto L651
	}
L649:
	;
	if int32(0) < v2794 {
		goto L647
	} else {
		goto L673
	}
L650:
	;
	if v2751 < v2752 {
		goto L653
	} else {
		goto L654
	}
L651:
	;
	goto L652
L652:
	;
	goto L659
L653:
	;
	v2757 = int32(-1)
	goto L655
L654:
	;
	v2757 = int32(1)
	goto L655
L655:
	;
	v2794 = v2757
	goto L649
L656:
	;
	v2794 = v2788
	goto L649
L657:
	;
	v2788 = int32(0)
	goto L656
L659:
	;
	goto L660
L660:
	;
	v2773 = int32(-1)
	v2774 = *(*float64)(unsafe.Add(mBase, uint32(v2658)+56))
	v2775 = *(*float64)(unsafe.Add(mBase, uint32(v2683)+56))
	if base.F64_lt(v2774, v2775) != 0 {
		v2788 = v2773
		goto L656
	} else {
		goto L667
	}
L667:
	;
	if base.F64_gt(v2774, v2775) != 0 {
		goto L668
	} else {
		goto L669
	}
L668:
	;
	v2794 = int32(1)
	goto L649
L669:
	;
	goto L670
L670:
	;
	v2779 = *(*float64)(unsafe.Add(mBase, uint32(v2658)+48))
	v2780 = *(*float64)(unsafe.Add(mBase, uint32(v2683)+48))
	if base.F64_lt(v2779, v2780) != 0 {
		v2788 = v2773
		goto L656
	} else {
		goto L671
	}
L671:
	;
	if base.F64_gt(v2779, v2780) != 0 {
		v2788 = int32(1)
		goto L656
	} else {
		goto L672
	}
L672:
	;
	goto L657
L673:
	;
	v2913 = v2658
	goto L626
L674:
	;
	v2798 = *(*int32)(unsafe.Add(mBase, uint32(v2797)+4))
	v2800 = v2798
	goto L676
L675:
	;
	v2800 = int32(0)
	goto L676
L676:
	;
	v2801 = int32(0)
	if base.B2i32(v2800 == v2801)|base.B2i32(v2396 == v2801) != 0 {
		v2847 = base.B2i32(v2800|v2396 == v2801)
		goto L678
	} else {
		goto L679
	}
L677:
	;
	if v2847 != 0 {
		v2913 = v2683
		goto L626
	} else {
		goto L688
	}
L678:
	;
	goto L677
L679:
	;
	v2815 = *(*int32)(unsafe.Add(mBase, uint32(v2800)+4))
	v2816 = *(*int32)(unsafe.Add(mBase, uint32(v2396)+4))
	if v2815 != v2816 {
		v2847 = int32(0)
		goto L678
	} else {
		goto L680
	}
L680:
	;
	v2818 = int32(1)
	if v2815 <= v2818 {
		goto L681
	} else {
		goto L682
	}
L681:
	;
	v2821 = v2818
	goto L683
L682:
	;
	v2821 = v2815
	goto L683
L683:
	;
	v2822 = int32(8)
	v2827 = int32(0)
	goto L684
L684:
	;
	v2835 = v2827 << (uint(int32(2)) % 32)
	v2837 = *(*int32)(unsafe.Add(mBase, uint32(v2800+v2822+v2835)))
	v2839 = *(*int32)(unsafe.Add(mBase, uint32(v2396+v2822+v2835)))
	v2840 = base.B2i32(v2837 == v2839)
	if v2837 != v2839 {
		v2847 = v2840
		goto L678
	} else {
		goto L686
	}
L685:
	;
	v2847 = v2840
	goto L678
L686:
	;
	v2843 = v2827 + int32(1)
	if v2843 != v2821 {
		v2827 = v2843
		goto L684
	} else {
		goto L687
	}
L687:
	;
	goto L685
L688:
	;
	v2853 = F_reparameterize_path(m, l0, v2683, v2396, float64(1))
	mBase = m.M
	v2854 = m.ExcPending
	if v2854 != 0 {
		goto L18
	} else {
		goto L689
	}
L689:
	;
	if v2853 != 0 {
		goto L690
	} else {
		goto L691
	}
L690:
	;
	v2855 = v2853
	goto L692
L691:
	;
	v2855 = v2658
	goto L692
L692:
	;
	v2856 = int32(0)
	if base.B2i32(v2658 == v2856)|base.B2i32(v2853 == v2856) != 0 {
		v2913 = v2855
		goto L626
	} else {
		goto L693
	}
L693:
	;
	v2866 = *(*int32)(unsafe.Add(mBase, uint32(v2658)+40))
	v2867 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+40))
	if v2866 != v2867 {
		goto L695
	} else {
		goto L696
	}
L694:
	;
	if v2909 <= int32(0) {
		goto L718
	} else {
		goto L719
	}
L695:
	;
	if v2866 < v2867 {
		goto L698
	} else {
		goto L699
	}
L696:
	;
	goto L697
L697:
	;
	goto L704
L698:
	;
	v2872 = int32(-1)
	goto L700
L699:
	;
	v2872 = int32(1)
	goto L700
L700:
	;
	v2909 = v2872
	goto L694
L701:
	;
	v2909 = v2903
	goto L694
L702:
	;
	v2903 = int32(0)
	goto L701
L704:
	;
	goto L705
L705:
	;
	v2888 = int32(-1)
	v2889 = *(*float64)(unsafe.Add(mBase, uint32(v2658)+56))
	v2890 = *(*float64)(unsafe.Add(mBase, uint32(v2853)+56))
	if base.F64_lt(v2889, v2890) != 0 {
		v2903 = v2888
		goto L701
	} else {
		goto L712
	}
L712:
	;
	if base.F64_gt(v2889, v2890) != 0 {
		goto L713
	} else {
		goto L714
	}
L713:
	;
	v2909 = int32(1)
	goto L694
L714:
	;
	goto L715
L715:
	;
	v2894 = *(*float64)(unsafe.Add(mBase, uint32(v2658)+48))
	v2895 = *(*float64)(unsafe.Add(mBase, uint32(v2853)+48))
	if base.F64_lt(v2894, v2895) != 0 {
		v2903 = v2888
		goto L701
	} else {
		goto L716
	}
L716:
	;
	if base.F64_gt(v2894, v2895) != 0 {
		v2903 = int32(1)
		goto L701
	} else {
		goto L717
	}
L717:
	;
	goto L702
L718:
	;
	v2912 = v2658
	goto L720
L719:
	;
	v2912 = v2853
	goto L720
L720:
	;
	v2913 = v2912
	goto L626
L721:
	;
	goto L625
L722:
	;
	v2925 = v2913
	goto L621
L723:
	;
	v2952 = v2413 + int32(1)
	v2953 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v2952 < v2953 {
		v2413 = v2952
		goto L559
	} else {
		goto L724
	}
L724:
	;
	goto L560
L725:
	;
	F_add_path(m, l1, v2990)
	mBase = m.M
	v2993 = m.ExcPending
	if v2993 != 0 {
		goto L18
	} else {
		goto L726
	}
L726:
	;
	goto L555
L727:
	;
	goto L554
L728:
	;
	m.G0 = v28 + int32(256)
	return
L729:
	;
	v3050 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v3050 != int32(1) {
		goto L728
	} else {
		goto L730
	}
L730:
	;
	v3053 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v3054 = *(*int32)(unsafe.Add(mBase, uint32(v3053)))
	v3055 = *(*int32)(unsafe.Add(mBase, uint32(v3054)+52))
	if v3055 == int32(0) {
		goto L728
	} else {
		goto L731
	}
L731:
	;
	v3058 = *(*int32)(unsafe.Add(mBase, uint32(v3055)+4))
	if v3058 < int32(2) {
		goto L728
	} else {
		goto L732
	}
L732:
	;
	v3075 = int32(1)
	goto L733
L733:
	;
	v3087 = *(*int32)(unsafe.Add(mBase, uint32(v3055)+12))
	v3091 = *(*int32)(unsafe.Add(mBase, uint32(v3087+v3075<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+248)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+240)) = int64(0)
	v3096 = *(*int32)(unsafe.Add(mBase, uint32(v3091)+64))
	if v3096 != 0 {
		goto L735
	} else {
		goto L736
	}
L734:
	;
	goto L728
L735:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v3091
	*(*int32)(unsafe.Add(mBase, uint32(v28)+140)) = v3091
	v3102 = F_list_make1_impl(m, int32(1), v28+int32(12))
	mBase = m.M
	v3103 = m.ExcPending
	if v3103 != 0 {
		goto L18
	} else {
		goto L738
	}
L736:
	;
	goto L737
L737:
	;
	v3119 = v3075 + int32(1)
	v3120 = *(*int32)(unsafe.Add(mBase, uint32(v3055)+4))
	if v3119 < v3120 {
		v3075 = v3119
		goto L733
	} else {
		goto L741
	}
L738:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+244)) = v3102
	v3105 = *(*int32)(unsafe.Add(mBase, uint32(v3091)+24))
	v3106 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v3106
	v3108 = *(*int64)(unsafe.Add(mBase, uint32(v28)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v3108
	v3110 = int32(0)
	v3113 = F_create_append_path(m, l0, l1, v28, v3110, v3110, v3105, int32(1), v987)
	mBase = m.M
	v3114 = m.ExcPending
	if v3114 != 0 {
		goto L18
	} else {
		goto L739
	}
L739:
	;
	F_add_partial_path(m, l1, v3113)
	mBase = m.M
	v3116 = m.ExcPending
	if v3116 != 0 {
		goto L18
	} else {
		goto L740
	}
L740:
	;
	goto L737
L741:
	;
	goto L734
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
					F_errfinish(m, int32(_a_F_add_reloption_kind_1), int32(747), int32(_a_F_add_reloption_kind_2))
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
	v136 = F_makeConst(m, int32(16), int32(-1), v131, v132, int64(0), v131, v132)
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
func F_anyrange_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anyrange_in_0), int32(207), int32(_a_F_anyrange_in_1), int32(_a_F_anyrange_in_2), int32(_a_F_anyrange_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
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
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+80))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)+84))
	if v4 < v5 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+88))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v7+v4<<(uint(int32(3))%32))))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+64)) = v11
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v3)+92))
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v4))))
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+72)) = uint8(v15)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+80)) = v4 + int32(1)
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
	v12 = F_palloc(m, int32(96))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v12
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
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
				F_deconstruct_array(m, v19, v31, v32, v33, v12+int32(88), v12+int32(92), v12+int32(84))
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
						*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = int32(7)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v57
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v62
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v66
						v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+14)))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v68
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+13)))
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+73)) = uint8(v70)
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v12 + int32(40)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = int32(0)
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_arraycontjoinsel(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int64
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4 == int32(2750) {
		v7 = int64(4576918229304087675)
	} else {
		v7 = int64(4572414629676717179)
	}
	return v7
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
func F_ascii(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v154 int64
	_ = v154
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_ascii[0]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	goto L3
L3:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v22 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L61
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L56
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L51
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L46
	}
L8:
	;
	m.G0 = v12 + int32(48)
	return v154
L9:
	;
	v57 = int32(1)
	if v22&v57 != 0 {
		goto L20
	} else {
		goto L21
	}
L10:
	;
	if int32(0) < v51 {
		v55 = v51
		goto L9
	} else {
		goto L19
	}
L11:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if base.Ui32((v26-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v55 = int32(4)
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v38 = int32(1)
	if v22&v38 != 0 {
		v51 = int32(base.Ui32(v22)>>(uint(v38)%32)) - v38
		goto L10
	} else {
		goto L18
	}
L14:
	;
	if v26 == int32(18) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v37 = int32(16)
	goto L17
L16:
	;
	v37 = int32(0)
	goto L17
L17:
	;
	v51 = v37
	goto L10
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v51 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L19:
	;
	v154 = int64(0)
	goto L8
L20:
	;
	v61 = v57
	goto L22
L21:
	;
	v61 = int32(4)
	goto L22
L22:
	;
	v62 = v15 + v61
	if v21 != int32(6) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if base.B2i32(v21 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(v21)) != 0 {
		goto L42
	} else {
		goto L43
	}
L24:
	;
	v65 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62))))
	if int32(0) <= v65 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	if base.Ui32(int32(-17)) < base.Ui32(v65) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if base.Ui32(v55) <= base.Ui32(v85) {
		goto L6
	} else {
		goto L34
	}
L27:
	;
	v82 = int32(7)
	v83 = v2
	v84 = v2
	v85 = int32(3)
	goto L26
L28:
	;
	goto L29
L29:
	;
	if base.Ui32(int32(-33)) < base.Ui32(v65) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v82 = int32(15)
	v83 = int32(1)
	v84 = v2
	v85 = int32(2)
	goto L26
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(v65) <= base.Ui32(int32(-64)) {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	v80 = int32(1)
	v82 = int32(31)
	v83 = v2
	v84 = v80
	v85 = v80
	goto L26
L34:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v87&int32(192) != int32(128) {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v99 = v87&int32(63) | v82&(v65&int32(255))<<(uint(int32(6))%32)
	if v84 != 0 {
		v123 = v99
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v154 = base.I64_extend_i32_u(v123)
	goto L8
L37:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+2)))
	if v100&int32(192) != int32(128) {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v109 = v100&int32(63) | v99<<(uint(int32(6))%32)
	if v83 != 0 {
		v123 = v109
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+3)))
	if v110&int32(192) != int32(128) {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v123 = v110&int32(63) | v109<<(uint(int32(6))%32)
	goto L36
L41:
	;
	v140 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62))))
	if base.B2i32(int32(2) <= v137)&base.B2i32(v140 < int32(0)) != 0 {
		goto L4
	} else {
		goto L45
	}
L42:
	;
	v137 = int32(1)
	goto L44
L43:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v21*int32(28))+uint32(_c_F_ascii[1])))
	v137 = v136
	goto L44
L44:
	;
	goto L41
L45:
	;
	v154 = base.I64_extend_i32_u(v140) & int64(255)
	goto L8
L46:
	;
	F_errcode(m, int32(17301634))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_ascii[0]))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v168
	F_errmsg(m, int32(_a_F_ascii_3), v12+int32(32))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_ascii_1), int32(990), int32(_a_F_ascii_2))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errcode(m, int32(17301634))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_ascii[0]))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v189
	F_errmsg(m, int32(_a_F_ascii_3), v12+int32(16))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_ascii_1), int32(997), int32(_a_F_ascii_2))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
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
	F_errcode(m, int32(17301634))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_ascii[0]))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v213
	F_errmsg(m, int32(_a_F_ascii_3), v12)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_ascii_1), int32(1005), int32(_a_F_ascii_2))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(_a_F_ascii_0), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_ascii_1), int32(1016), int32(_a_F_ascii_2))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_assign_random_seed(m *base.Module, l0 float64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v3 != 0 {
		v7 = F_DirectFunctionCall1Coll(m, int32(632), int32(0), base.I64_reinterpret_f64(l0))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
		return
	}
}
func F_autoprewarm_database_main(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int64
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v23 = m.G0
	v25 = v23 - int32(32)
	m.G0 = v25
	v28 = int32(974)
	switch v28 {
	case 0, 2:
		goto L2
	default:
		goto L3
	}
L1:
	;
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	F_sigemptyset(m, v25+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = int32(268435456)
	switch v28 {
	case 0:
		goto L7
	default:
		goto L5
	case 2:
		goto L6
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[0])) = int32(972)
	goto L2
L4:
	;
	goto L9
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = int32(_a_F_autoprewarm_database_main_0)
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = int32(0)
	goto L4
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = int32(-2)
	goto L4
L9:
	;
	goto L10
L10:
	;
	v57 = F___sigaction(m, int32(15), v25+int32(12), int32(0))
	mBase = m.M
	m.G0 = v25 + int32(32)
	goto L1
L11:
	;
	return
L12:
	;
	v66 = F_GetNamedDSMSegment(m, v17+int32(24))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1])) = v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+24))
	v70 = F_dsm_attach(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	if v70 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+28))
	v75 = int32(0)
	F_BackgroundWorkerInitializeConnectionByOid(m, v74, v75, v75)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L11
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L11
	} else {
		goto L81
	}
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)+24))
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+32))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)+36))
	if v83 <= v82 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_dsm_detach(m, v70)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L11
	} else {
		goto L80
	}
L20:
	;
	v87 = v79 + v82*int32(20)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v92 = v82
	v93 = v89
	v97 = v88
	goto L21
L21:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L11
	} else {
		goto L23
	}
L22:
	;
	goto L19
L23:
	;
	v106 = F_RelidByRelfilenumber(m, v97, v93)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L11
	} else {
		goto L30
	}
L24:
	;
	if v352 < v354 {
		v92 = v352
		v93 = v353
		v97 = v357
		goto L21
	} else {
		goto L79
	}
L25:
	;
	F_relation_close(m, v111, int32(1))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L11
	} else {
		goto L77
	}
L26:
	;
	v329 = v97
	v330 = v316
	v334 = v93
	goto L25
L27:
	;
	v352 = v128
	v353 = v143
	v354 = v124
	v357 = v144
	goto L24
L28:
	;
	v153 = v92
	v162 = v116
	goto L41
L29:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L11
	} else {
		goto L35
	}
L30:
	;
	if v106 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v111 = F_try_relation_open(m, v106, int32(1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	if v111 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+36))
	if v92 < v117 {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	v316 = v92
	goto L26
L35:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+36))
	if v124 <= v92 {
		v352 = v92
		v353 = v93
		v354 = v124
		v357 = v97
		goto L24
	} else {
		goto L36
	}
L36:
	;
	v128 = v92
	goto L37
L37:
	;
	v142 = v79 + v128*int32(20)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+8))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if base.B2i32(v144 != v97)|base.B2i32(v93 != v143) != 0 {
		goto L27
	} else {
		goto L39
	}
L38:
	;
	goto L19
L39:
	;
	v149 = v128 + int32(1)
	if v149 != v124 {
		v128 = v149
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v167 = v79 + v153*int32(20)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if base.B2i32(v169 != v97)|base.B2i32(v93 != v168) != 0 {
		v329 = v169
		v330 = v153
		v334 = v168
		goto L25
	} else {
		goto L43
	}
L42:
	;
	v329 = v300
	v330 = v301
	v334 = v305
	goto L25
L43:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	if base.Ui32(v173) <= base.Ui32(int32(3)) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	if v301 < v303 {
		v153 = v301
		v162 = v310
		goto L41
	} else {
		goto L76
	}
L45:
	;
	v300 = v97
	v301 = v287
	v303 = v289
	v305 = v93
	v310 = v296
	goto L44
L46:
	;
	v238 = F_RelationGetNumberOfBlocksInFork(m, v111, v173)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L11
	} else {
		goto L66
	}
L47:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	if v176 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v208 = v162
	goto L49
L49:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+36))
	if v209 <= v153 {
		v287 = v153
		v289 = v209
		v296 = v208
		goto L45
	} else {
		goto L60
	}
L50:
	;
	v202 = v176
	goto L52
L51:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v178
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v111)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v180
	v184 = F_smgropen(m, v17+int32(8), v177)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L11
	} else {
		goto L53
	}
L52:
	;
	v203 = F_smgrexists(m, v202, v173)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L11
	} else {
		goto L58
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = v184
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)+72))
	if v188 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v202 = v200
	goto L52
L55:
	;
	v196 = v188
	goto L57
L56:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v184)+76))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v184)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v184)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v184)+72))
	v196 = v194
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184)+72)) = v196 + int32(1)
	goto L54
L58:
	;
	if v203 != 0 {
		goto L46
	} else {
		goto L59
	}
L59:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v208 = v206
	goto L49
L60:
	;
	v213 = v153
	goto L61
L61:
	;
	v227 = v79 + v213*int32(20)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if base.B2i32(v229 != v97)|base.B2i32(v93 != v228) != 0 {
		v300 = v229
		v301 = v213
		v303 = v209
		v305 = v228
		v310 = v208
		goto L44
	} else {
		goto L63
	}
L62:
	;
	v316 = v209
	goto L26
L63:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	if v233 != v173 {
		v300 = v229
		v301 = v213
		v303 = v209
		v305 = v228
		v310 = v208
		goto L44
	} else {
		goto L64
	}
L64:
	;
	v236 = v213 + int32(1)
	if v236 != v209 {
		v213 = v236
		goto L61
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v79
	v247 = int32(0)
	v252 = F_read_stream_begin_relation(m, int32(9), v247, v111, v173, int32(_a_F_autoprewarm_database_main_1), v17+int32(24), v247)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	goto L68
L68:
	;
	v269 = F_read_stream_next_buffer(m, v252, int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L11
	} else {
		goto L70
	}
L69:
	;
	F_read_stream_end(m, v252)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L11
	} else {
		goto L75
	}
L70:
	;
	if v269 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v272)+40)) = v273 + int32(1)
	F_ReleaseBuffer(m, v269)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L11
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	goto L69
L74:
	;
	goto L68
L75:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+36))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v287 = v284
	v289 = v283
	v296 = v282
	goto L45
L76:
	;
	goto L42
L77:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_database_main[1]))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)+36))
	v352 = v330
	v353 = v334
	v354 = v349
	v357 = v329
	goto L24
L79:
	;
	goto L22
L80:
	;
	m.G0 = v17 + int32(48)
	return
L81:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(_a_F_autoprewarm_database_main_2), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L11
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_database_main_3), int32(518), int32(_a_F_autoprewarm_database_main_4))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L11
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_autoprewarm_main(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v308 int32
	_ = v308
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v478 int32
	_ = v478
	var v487 int32
	_ = v487
	var v490 int64
	_ = v490
	var v493 int32
	_ = v493
	var v496 int64
	_ = v496
	var v499 int64
	_ = v499
	var v502 int64
	_ = v502
	var v505 int32
	_ = v505
	var v508 int64
	_ = v508
	var v511 int64
	_ = v511
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int64
	_ = v613
	var v614 int64
	_ = v614
	var v627 int32
	_ = v627
	var v636 int64
	_ = v636
	var v638 int32
	_ = v638
	var v652 int64
	_ = v652
	var v654 int32
	_ = v654
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int64
	_ = v680
	var v681 int64
	_ = v681
	var v689 int64
	_ = v689
	var v695 int64
	_ = v695
	var v701 int64
	_ = v701
	var v710 int64
	_ = v710
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int64
	_ = v743
	var v744 int64
	_ = v744
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v795 int32
	_ = v795
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	v15 = m.G0
	v17 = v15 - int32(1648)
	m.G0 = v17
	v23 = m.G0
	v25 = v23 - int32(32)
	m.G0 = v25
	v28 = int32(969)
	switch v28 {
	case 0, 2:
		goto L2
	default:
		goto L3
	}
L1:
	;
	v61 = int32(1)
	v66 = m.G0
	v68 = v66 - int32(32)
	m.G0 = v68
	v71 = int32(967)
	switch v71 {
	case 0, 2:
		goto L12
	default:
		goto L13
	}
L2:
	;
	F_sigemptyset(m, v25+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = int32(268435456)
	switch v28 {
	case 0:
		goto L7
	default:
		goto L5
	case 2:
		goto L6
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[0])) = int32(967)
	goto L2
L4:
	;
	goto L9
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = int32(_a_F_autoprewarm_main_0)
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = int32(0)
	goto L4
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = int32(-2)
	goto L4
L9:
	;
	goto L10
L10:
	;
	v57 = F___sigaction(m, int32(15), v25+int32(12), int32(0))
	mBase = m.M
	m.G0 = v25 + int32(32)
	goto L1
L11:
	;
	v108 = m.G0
	v110 = v108 - int32(32)
	m.G0 = v110
	v113 = int32(970)
	switch v113 {
	case 0, 2:
		goto L22
	default:
		goto L23
	}
L12:
	;
	F_sigemptyset(m, v68+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(268435456)
	switch v71 {
	case 0:
		goto L17
	default:
		goto L15
	case 2:
		goto L16
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[1])) = int32(965)
	goto L12
L14:
	;
	goto L19
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = int32(_a_F_autoprewarm_main_0)
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = int32(0)
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = int32(-2)
	goto L14
L19:
	;
	goto L20
L20:
	;
	v100 = F___sigaction(m, v61, v68+int32(12), int32(0))
	mBase = m.M
	m.G0 = v68 + int32(32)
	goto L11
L21:
	;
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L31
	} else {
		goto L32
	}
L22:
	;
	F_sigemptyset(m, v110+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v110)+24)) = int32(268435456)
	switch v113 {
	case 0:
		goto L27
	default:
		goto L25
	case 2:
		goto L26
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[2])) = int32(968)
	goto L22
L24:
	;
	goto L29
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+12)) = int32(_a_F_autoprewarm_main_0)
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+12)) = int32(0)
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+12)) = int32(-2)
	goto L24
L29:
	;
	goto L30
L30:
	;
	v142 = F___sigaction(m, int32(10), v110+int32(12), int32(0))
	mBase = m.M
	m.G0 = v110 + int32(32)
	goto L21
L31:
	;
	return
L32:
	;
	v151 = F_GetNamedDSMSegment(m, v17+int32(176))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3])) = v151
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+176)))
	F_before_shmem_exit(m, int32(_a_F_autoprewarm_main_1), int64(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v162 = F_LWLockAcquire(m, v160, int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	if v166 != int32(-1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	m.G0 = v17 + int32(1648)
	return
L37:
	;
	F_LWLockRelease(m, v165)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L31
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v192
	F_LWLockRelease(m, v165)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L31
	} else {
		goto L45
	}
L40:
	;
	v173 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L31
	} else {
		goto L41
	}
L41:
	;
	if v173 == int32(0) {
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+160)) = v179
	F_errmsg(m, int32(_a_F_autoprewarm_main_2), v17+int32(160))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L31
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_main_3), int32(203), int32(_a_F_autoprewarm_main_4))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L31
	} else {
		goto L44
	}
L44:
	;
	goto L36
L45:
	;
	if v154&int32(1) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	if v627 == int32(0) {
		goto L36
	} else {
		goto L162
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L31
	} else {
		goto L158
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L31
	} else {
		goto L155
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L31
	} else {
		goto L150
	}
L50:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v203 = F_LWLockAcquire(m, v201, int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L31
	} else {
		goto L53
	}
L51:
	;
	v627 = v61
	v636 = int64(0)
	goto L52
L52:
	;
	v638 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[5]))
	if v638 != 0 {
		goto L46
	} else {
		goto L125
	}
L53:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+20))
	if v207 == int32(-1) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v602 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[5]))
	v608 = m.G0
	v609 = int32(16)
	v610 = v608 - v609
	m.G0 = v610
	F_gettimeofday(m, v610)
	mBase = m.M
	v613 = *(*int64)(unsafe.Add(mBase, uint32(v610)))
	v614 = int64(*(*int32)(unsafe.Add(mBase, uint32(v610)+8)))
	m.G0 = v610 + v609
	goto L124
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = v17 + int32(168)
	v278 = F_fscanf(m, v217, int32(_a_F_autoprewarm_main_5), v17+int32(128))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L31
	} else {
		goto L75
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L31
	} else {
		goto L71
	}
L57:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v206)+20)) = v211
	F_LWLockRelease(m, v206)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L31
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_LWLockRelease(m, v206)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L31
	} else {
		goto L66
	}
L60:
	;
	v217 = F_AllocateFile(m, int32(_a_F_autoprewarm_main_6), int32(_a_F_autoprewarm_main_7))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L31
	} else {
		goto L61
	}
L61:
	;
	if v217 != 0 {
		goto L55
	} else {
		goto L62
	}
L62:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[6]))
	if v220 != int32(44) {
		goto L56
	} else {
		goto L63
	}
L63:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v226 = F_LWLockAcquire(m, v224, int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L31
	} else {
		goto L64
	}
L64:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v229)+20)) = int32(-1)
	F_LWLockRelease(m, v229)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L31
	} else {
		goto L65
	}
L65:
	;
	goto L54
L66:
	;
	v238 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L31
	} else {
		goto L67
	}
L67:
	;
	if v238 == int32(0) {
		goto L54
	} else {
		goto L68
	}
L68:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+144)) = v244
	F_errmsg(m, int32(_a_F_autoprewarm_main_8), v17+int32(144))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L31
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_main_3), int32(313), int32(_a_F_autoprewarm_main_9))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L31
	} else {
		goto L70
	}
L70:
	;
	goto L54
L71:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L31
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(_a_F_autoprewarm_main_6)
	F_errmsg(m, int32(_a_F_autoprewarm_main_10), v17)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L31
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_main_3), int32(335), int32(_a_F_autoprewarm_main_9))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L31
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	if v278 != int32(1) {
		goto L47
	} else {
		goto L76
	}
L76:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
	v286 = F_dsm_create(m, v282*int32(20), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L31
	} else {
		goto L77
	}
L77:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v286)+24))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
	if int32(0) < v289 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v295 = int32(0)
	goto L81
L79:
	;
	goto L80
L80:
	;
	v349 = F_FreeFile(m, v217)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L31
	} else {
		goto L86
	}
L81:
	;
	v308 = v288 + v295*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v308 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v17)+84)) = v308 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v308 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = v17 + int32(176)
	v325 = F_fscanf(m, v217, int32(_a_F_autoprewarm_main_11), v17+int32(80))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L31
	} else {
		goto L83
	}
L82:
	;
	goto L80
L83:
	;
	if v325 != int32(5) {
		goto L48
	} else {
		goto L84
	}
L84:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v17)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v308)+12)) = v329
	v332 = v295 + int32(1)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
	if v332 < v333 {
		v295 = v332
		goto L81
	} else {
		goto L85
	}
L85:
	;
	goto L82
L86:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
	F_pg_qsort(m, v288, v351, int32(20), int32(_a_F_autoprewarm_main_12))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L31
	} else {
		goto L87
	}
L87:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v357)+24)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v357)+40)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v357)+32)) = int64(0)
	v365 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[7]))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
	if v366 <= v365 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)+32))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
	if v391 <= v390 {
		goto L94
	} else {
		goto L95
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+168)) = v365
	v371 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L31
	} else {
		goto L90
	}
L90:
	;
	if v371 == int32(0) {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v376
	F_errmsg(m, int32(_a_F_autoprewarm_main_13), v17+int32(48))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L31
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_main_3), int32(380), int32(_a_F_autoprewarm_main_9))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L31
	} else {
		goto L93
	}
L93:
	;
	goto L88
L94:
	;
	F_dsm_detach(m, v286)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L31
	} else {
		goto L116
	}
L95:
	;
	v394 = v17 + int32(272)
	v396 = v17 + int32(1404)
	v398 = v17 + int32(380)
	v404 = v390
	v406 = v391
	v408 = v389
	goto L96
L96:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v288+v404*int32(20))))
	v420 = v404 + int32(1)
	if v420 < v406 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L94
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v408)+28)) = v463
	*(*int32)(unsafe.Add(mBase, uint32(v408)+36)) = v464
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[5]))
	if v478 != 0 {
		goto L94
	} else {
		goto L111
	}
L99:
	;
	v424 = v418
	v425 = v420
	goto L102
L100:
	;
	v447 = v418
	v448 = v420
	goto L101
L101:
	;
	if v447 == int32(0) {
		goto L94
	} else {
		goto L110
	}
L102:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v288+v425*int32(20))))
	if v439 == v424 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v447 = v441
	v448 = v406
	goto L101
L104:
	;
	v443 = v425 + int32(1)
	if v443 != v406 {
		v424 = v441
		v425 = v443
		goto L102
	} else {
		goto L109
	}
L105:
	;
	v441 = v424
	goto L104
L106:
	;
	goto L107
L107:
	;
	if v424 != 0 {
		v463 = v424
		v464 = v425
		goto L98
	} else {
		goto L108
	}
L108:
	;
	v441 = v439
	goto L104
L109:
	;
	goto L103
L110:
	;
	v463 = v447
	v464 = v448
	goto L98
L111:
	;
	base.MemoryFill(m, v17+int32(192), int32(0), int32(1456))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+376)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+368)) = int64(4294967299)
	v487 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v398)+7)) = v487
	v490 = *(*int64)(unsafe.Add(mBase, _c_F_autoprewarm_main[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v398))) = v490
	v493 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_autoprewarm_main[10])))
	*(*uint16)(unsafe.Add(mBase, uint32(v396)+24)) = uint16(v493)
	v496 = *(*int64)(unsafe.Add(mBase, _c_F_autoprewarm_main[11]))
	*(*int64)(unsafe.Add(mBase, uint32(v396)+16)) = v496
	v499 = *(*int64)(unsafe.Add(mBase, _c_F_autoprewarm_main[12]))
	*(*int64)(unsafe.Add(mBase, uint32(v396)+8)) = v499
	v502 = *(*int64)(unsafe.Add(mBase, _c_F_autoprewarm_main[13]))
	*(*int64)(unsafe.Add(mBase, uint32(v396))) = v502
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+191)) = v505
	v508 = *(*int64)(unsafe.Add(mBase, _c_F_autoprewarm_main[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+184)) = v508
	v511 = *(*int64)(unsafe.Add(mBase, _c_F_autoprewarm_main[16]))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+176)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v394)+15)) = v505
	*(*int64)(unsafe.Add(mBase, uint32(v394)+8)) = v508
	*(*int64)(unsafe.Add(mBase, uint32(v394))) = v511
	v517 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+1640)) = v517
	v523 = F_RegisterDynamicBackgroundWorker(m, v17+int32(176), v17+int32(172))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L31
	} else {
		goto L112
	}
L112:
	;
	if v523 == int32(0) {
		goto L49
	} else {
		goto L113
	}
L113:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v17)+172))
	v528 = F_WaitForBackgroundWorkerShutdown(m, v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L31
	} else {
		goto L114
	}
L114:
	;
	v531 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v531)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v531)+32)) = v532
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
	if v532 < v534 {
		v404 = v532
		v406 = v534
		v408 = v531
		goto L96
	} else {
		goto L115
	}
L115:
	;
	goto L97
L116:
	;
	v553 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v555 = F_LWLockAcquire(m, v553, int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L31
	} else {
		goto L117
	}
L117:
	;
	v558 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v558)+20)) = int64(4294967295)
	F_LWLockRelease(m, v558)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L31
	} else {
		goto L118
	}
L118:
	;
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[5]))
	if v564 != 0 {
		goto L54
	} else {
		goto L119
	}
L119:
	;
	v567 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L31
	} else {
		goto L120
	}
L120:
	;
	if v567 == int32(0) {
		goto L54
	} else {
		goto L121
	}
L121:
	;
	v572 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[3]))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v573
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v575
	F_errmsg(m, int32(_a_F_autoprewarm_main_14), v17+int32(16))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L31
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_main_3), int32(451), int32(_a_F_autoprewarm_main_9))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L31
	} else {
		goto L123
	}
L123:
	;
	goto L54
L124:
	;
	v627 = base.B2i32(v602 == int32(0))
	v636 = v614 + v613*int64(1000000) - int64(946684800000000)
	goto L52
L125:
	;
	v652 = v636
	goto L126
L126:
	;
	v654 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[17]))
	if v654 != 0 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	goto L46
L128:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[17])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L31
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v662 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[18]))
	if v662 <= int32(0) {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	goto L130
L132:
	;
	v738 = m.G0
	v739 = int32(16)
	v740 = v738 - v739
	m.G0 = v740
	F_gettimeofday(m, v740)
	mBase = m.M
	v743 = *(*int64)(unsafe.Add(mBase, uint32(v740)))
	v744 = int64(*(*int32)(unsafe.Add(mBase, uint32(v740)+8)))
	m.G0 = v740 + v739
	goto L147
L133:
	;
	v724 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[19]))
	v725 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v724))) = v725
	v730 = base.AtomicRmwOr32(m, v725, int32(_a_F_autoprewarm_main_15), v725)
	goto L145
L134:
	;
	v666 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[19]))
	v670 = F_WaitLatch(m, v666, int32(33), int32(-1), int32(117440512))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L31
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v675 = m.G0
	v676 = int32(16)
	v677 = v675 - v676
	m.G0 = v677
	F_gettimeofday(m, v677)
	mBase = m.M
	v680 = *(*int64)(unsafe.Add(mBase, uint32(v677)))
	v681 = int64(*(*int32)(unsafe.Add(mBase, uint32(v677)+8)))
	m.G0 = v677 + v676
	v689 = v681 + v680*int64(1000000) - int64(946684800000000)
	goto L138
L137:
	;
	goto L133
L138:
	;
	v695 = base.I64_extend_i32_u(v662*int32(1000))*int64(1000) + v652
	if v695 <= v689 {
		v713 = int32(0)
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v713 <= int32(0) {
		goto L132
	} else {
		goto L143
	}
L140:
	;
	goto L139
L141:
	;
	v701 = v695 - v689
	if base.B2i32(int64(0) < v689)^base.B2i32(v701 < v695)|base.B2i32(int64(2147483646000) < v701) != 0 {
		v713 = int32(2147483647)
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v710 = base.I64_div_s(v701+int64(999), int64(1000))
	v713 = base.I32_wrap_i64(v710)
	goto L140
L143:
	;
	v717 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[19]))
	v720 = F_WaitLatch(m, v717, int32(41), v713, int32(117440512))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L31
	} else {
		goto L144
	}
L144:
	;
	goto L133
L145:
	;
	v732 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[5]))
	if v732 == int32(0) {
		goto L126
	} else {
		goto L146
	}
L146:
	;
	goto L46
L147:
	;
	v755 = F_apw_dump_now(m, int32(1), int32(0))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L31
	} else {
		goto L148
	}
L148:
	;
	v758 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_main[5]))
	if v758 == int32(0) {
		v652 = v744 + v743*int64(1000000) - int64(946684800000000)
		goto L126
	} else {
		goto L149
	}
L149:
	;
	goto L127
L150:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L31
	} else {
		goto L151
	}
L151:
	;
	F_errmsg(m, int32(_a_F_autoprewarm_main_16), int32(0))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L31
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(_a_F_autoprewarm_main_17)
	F_errhint(m, int32(_a_F_autoprewarm_main_18), v17+int32(32))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L31
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_main_3), int32(979), int32(_a_F_autoprewarm_main_19))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L31
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v295 + int32(1)
	F_errmsg(m, int32(_a_F_autoprewarm_main_20), v17-int32(-64))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L31
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_main_3), int32(359), int32(_a_F_autoprewarm_main_9))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L31
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L31
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = int32(_a_F_autoprewarm_main_6)
	F_errmsg(m, int32(_a_F_autoprewarm_main_21), v17+int32(112))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L31
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_autoprewarm_main_3), int32(343), int32(_a_F_autoprewarm_main_9))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L31
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	v835 = int32(1)
	v837 = F_apw_dump_now(m, v835, v835)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L31
	} else {
		goto L163
	}
L163:
	;
	goto L36
}
func F_autoprewarm_start_worker(m *base.Module, l0 int32) int64 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_autoprewarm_start_worker[0])))
	if v9 != 0 {
		v13 = F_GetNamedDSMSegment(m, v6+int32(15))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_start_worker[1])) = v13
			v19 = F_LWLockAcquire(m, v13, int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_start_worker[1]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
				F_LWLockRelease(m, v22)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					if v23 != int32(-1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v23
								F_errmsg(m, int32(_a_F_autoprewarm_start_worker_0), v6)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_autoprewarm_start_worker_1), int32(842), int32(_a_F_autoprewarm_start_worker_2))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int64(0)
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
						v29 = m.ExcPending
						if v29 != 0 {
							return int64(0)
						} else {
							m.G0 = v6 + int32(16)
							return int64(0)
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_autoprewarm_start_worker_3), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_autoprewarm_start_worker_1), int32(831), int32(_a_F_autoprewarm_start_worker_2))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
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
func F_avals(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_avals(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
